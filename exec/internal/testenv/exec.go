// 版权所有 ? 2015 The Go Authors。保留所有权利。
// 本源代码的使用受 BSD 风格许可协议约束，
// 该协议可在 LICENSE 文件中找到。

package testenv

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"
)

// MustHaveExec 检查当前系统是否能够使用 os.StartProcess 或（更常见地）exec.Command 启动新的进程。
// 如果不能，MustHaveExec 将调用 t.Skip 并附带解释原因。
//
// 在某些平台上，MustHaveExec 通过重新执行当前可执行文件来检查 exec 支持，
// 这个可执行文件必须是由 'go test' 构建的二进制文件。我们故意不提供 HasExec 函数，
// 因为 TestMain 函数中存在不合适递归的风险。
//
// 要在测试之外检查 exec 支持，只需尝试执行命令即可。如果 exec 不被支持，
// 对于由此产生的错误，testenv.SyscallIsNotSupported 将返回 true。
func MustHaveExec(t testing.TB) {
	tryExecOnce.Do(func() {
		tryExecErr = tryExec()
	})
	if tryExecErr != nil {
		t.Skipf("skipping test: cannot exec subprocess on %s/%s: %v", runtime.GOOS, runtime.GOARCH, tryExecErr)
	}
}

var (
	tryExecOnce sync.Once
	tryExecErr  error
)

func tryExec() error {
	switch runtime.GOOS {
	case "wasip1", "js", "ios":
	default:
		// 假设 exec 在非移动平台和 Android 上始终能正常工作
		return nil
	}

// 在 iOS 上存在 exec 系统调用，但在实际的 iOS 设备上，它可能会返回权限错误。在模拟环境中（如 Corellium 主机），exec 可能会成功执行。因此，如果我们需要使用 exec，就只能尝试一下并观察结果。
//
// 截至 2023-04-19，wasip1 和 js 完全没有 exec 系统调用，但我们也应采用相同的处理路径，以便在没有 iOS 环境的情况下对这一分支进行测试。

	if !testing.Testing() {
// 这并不是一个标准的“go test”二进制文件，因此我们不知道如何以一种没有副作用的方式使其自执行。
// 算了，放弃吧。
		return errors.New("can't probe for exec support with a non-test executable")
	}

// 我们知道这是一个测试可执行文件。我们应该能够使用一个无操作（no-op）标志来运行它，以检查整体执行支持情况。
	exe, err := os.Executable()
	if err != nil {
		return fmt.Errorf("can't probe for exec support: %w", err)
	}
	cmd := exec.Command(exe, "-test.list=^$")
	cmd.Env = origEnv
	return cmd.Run()
}

var execPaths sync.Map // path -> error

// MustHaveExecPath 检查当前系统是否能够使用 os.StartProcess 或（更常见地）exec.Command 启动指定的可执行文件。
// 若不能，MustHaveExecPath 将调用 t.Skip 并附带解释原因。
func MustHaveExecPath(t testing.TB, path string) {
	MustHaveExec(t)

	err, found := execPaths.Load(path)
	if !found {
		_, err = exec.LookPath(path)
		err, _ = execPaths.LoadOrStore(path, err)
	}
	if err != nil {
		t.Skipf("skipping test: %s: %s", path, err)
	}
}

// CleanCmdEnv 会将环境变量填充到 cmd.Env 中，但会排除某些可能修改 Go 工具行为的特定变量，
// 如 GODEBUG 和 GOTRACEBACK。
// 
// 如果调用者想要设置 cmd.Dir，请在调用本函数前完成设置，以便环境中PWD能够正确设定。
func CleanCmdEnv(cmd *exec.Cmd) *exec.Cmd {
	if cmd.Env != nil {
		panic("environment already set")
	}
	for _, env := range cmd.Environ() {
// 从环境中排除GODEBUG，以防止其输出干扰那些试图解析其他命令输出的测试。
		if strings.HasPrefix(env, "GODEBUG=") {
			continue
		}
		// 同理，排除GOTRACEBACK
		if strings.HasPrefix(env, "GOTRACEBACK=") {
			continue
		}
		cmd.Env = append(cmd.Env, env)
	}
	return cmd
}

// CommandContext 与 exec.CommandContext 类似，但具有以下特点：
//   - 若平台不支持 os/exec，则跳过 t；
//   - 在其 Cancel 函数中发送 SIGQUIT（若平台支持）而非 SIGKILL；
//   - 若测试具有截止时间，则在测试截止时间前设置一段任意的宽限期，添加 Context 超时及 WaitDelay；
//   - 若命令未能在测试截止时间前完成，导致测试失败；
//   - 设置一个 Cleanup 函数，用于验证测试未泄露子进程。
func CommandContext(t testing.TB, ctx context.Context, name string, args ...string) *exec.Cmd {
	t.Helper()
	MustHaveExec(t)

	var (
		cancelCtx   context.CancelFunc
		gracePeriod time.Duration // 除非测试具有截止时间（以允许进行交互式调试），否则不受限制
	)

	if t, ok := t.(interface {
		testing.TB
		Deadline() (time.Time, bool)
	}); ok {
		if td, ok := t.Deadline(); ok {
// 以一个最小的宽限期开始，这个时间长度仅足以在程序终止后接收其合理的输出。
			gracePeriod = 100 * time.Millisecond
			if s := os.Getenv("GO_TEST_TIMEOUT_SCALE"); s != "" {
				scale, err := strconv.Atoi(s)
				if err != nil {
					t.Fatalf("invalid GO_TEST_TIMEOUT_SCALE: %v", err)
				}
				gracePeriod *= time.Duration(scale)
			}

// 如果时间允许，将终止宽限期增加至测试剩余时间的5%。
			testTimeout := time.Until(td)
			if gp := testTimeout / 20; gp > gracePeriod {
				gracePeriod = gp
			}

// 当我们执行涉及子进程的命令时，我们希望预留两个清理阶段：
// 一是从发送第一个终止信号（通过Context过期时的Cancel回调）到强制终止进程（通过WaitDelay字段）之间的时间；
// 二是从进程被终止到测试为调试目的记录其输出之间的时间。
//
// （我们需要确保测试进程自身有足够时间在被终止前记录下输出。）
			cmdTimeout := testTimeout - 2*gracePeriod

			if cd, ok := ctx.Deadline(); !ok || time.Until(cd) > cmdTimeout {
// 或者ctx没有截止时间，或者其截止时间将在（或过早于）测试已超时时到期。
// 添加一个较短的超时时间，以便测试能产生有用输出。
				ctx, cancelCtx = context.WithTimeout(ctx, cmdTimeout)
			}
		}
	}

	cmd := exec.CommandContext(ctx, name, args...)
	cmd.Cancel = func() error {
		if cancelCtx != nil && ctx.Err() == context.DeadlineExceeded {
			// The command timed out due to running too close to the test's deadline.
			// There is no way the test did that intentionally — it's too close to the
			// wire! — so mark it as a test failure. That way, if the test expects the
			// command to fail for some other reason, it doesn't have to distinguish
			// between that reason and a timeout.
			t.Errorf("test timed out while running command: %v", cmd)
		} else {
// 由于ctx被取消，命令正在被终止，但
// 显然并非由于我们添加的显式测试截止时间导致。
// 将此信息记录下来，以便于诊断失败原因，
// 但切勿因此而令测试失败。
			t.Logf("%v: terminating command: %v", ctx.Err(), cmd)
		}
		return cmd.Process.Signal(Sigquit)
	}
	cmd.WaitDelay = gracePeriod

	t.Cleanup(func() {
		if cancelCtx != nil {
			cancelCtx()
		}
		if cmd.Process != nil && cmd.ProcessState == nil {
			t.Errorf("command was started, but test did not wait for it to complete: %v", cmd)
		}
	})

	return cmd
}

// Command类似于exec.Command，但会应用与testenv.CommandContext相同的变化（使用默认的Context）。
func Command(t testing.TB, name string, args ...string) *exec.Cmd {
	t.Helper()
	return CommandContext(t, context.Background(), name, args...)
}
