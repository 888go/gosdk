// 版权所有 2015 The Go 作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

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

// MustHaveExec 检查当前系统是否能够使用 os.StartProcess 或（更常见地）exec.Command 启动新进程。
// 若不能，MustHaveExec 会调用 t.Skip 并附带解释原因。
//
// 在某些平台上，MustHaveExec 通过重新执行当前可执行文件来检查 exec 支持情况，
// 此可执行文件必须是由 'go test' 构建的二进制文件。我们有意不提供 HasExec 函数，
// 因为在 TestMain 函数中存在不合适递归的风险。
//
// 要在测试之外检查 exec 支持情况，只需尝试执行命令即可。若不支持 exec，
// 对于由此产生的错误，testenv.SyscallIsNotSupported 将返回 true。

// ff:
// t:
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
		// 假设exec总是在非移动平台和Android上正常工作。
		return nil
	}

// 在iOS中有一个exec系统调用，但在实际的iOS设备上可能会返回权限错误。在模拟环境中（如Corellium主机），它可能会成功。因此，如果我们需要执行，就只能尝试并找出结果。
// 
// 截至2023-04-19，wasip1和js根本没有exec系统调用，但我们最好使用相同的路径，这样即使没有iOS环境，这个分支也可以进行测试。

	if !testing.Testing() {
// 这不是一个标准的 "go test" 可执行文件，所以我们不知道如何以不会产生副作用的方式自我执行。就忽略它吧。
		return errors.New("can't probe for exec support with a non-test executable")
	}

// 我们知道这是一个测试可执行文件。我们应该能够使用一个无操作（no-op）标志来运行它，以检查总体上的执行支持情况。
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
// 如果不能，MustHaveExecPath 将调用 t.Skip 并给出解释。

// ff:
// path:
// t:
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

// CleanCmdEnv 将向 cmd.Env 中填充环境，但排除可能修改 Go 工具行为的某些变量，如 GODEBUG 和 GOTRACEBACK。
// 
// 如果调用者想要设置 cmd.Dir，请在调用此函数之前设置，以便 PWD 在环境中被正确设置。

// ff:
// cmd:
func CleanCmdEnv(cmd *exec.Cmd) *exec.Cmd {
	if cmd.Env != nil {
		panic("environment already set")
	}
	for _, env := range cmd.Environ() {
// 从环境中排除GODEBUG，防止其输出破坏试图解析其他命令输出的测试。
		if strings.HasPrefix(env, "GODEBUG=") {
			continue
		}
		// Exclude GOTRACEBACK for the same reason.
		if strings.HasPrefix(env, "GOTRACEBACK=") {
			continue
		}
		cmd.Env = append(cmd.Env, env)
	}
	return cmd
}

// CommandContext 类似于 exec.CommandContext，但有以下不同：
//   - 如果平台不支持 os/exec，则跳过 t
//   - 在 Cancel 函数中发送 SIGQUIT（如果平台支持）而不是 SIGKILL
//   - 如果测试有截止时间，会在测试截止时间前添加一个上下文超时和 WaitDelay 以提供一段任意的宽限期
//   - 如果命令在测试的截止时间之前未完成，会导致测试失败
//   - 设置一个 Cleanup 函数，用于检查测试是否泄漏了子进程

// ff:
// args:
// name:
// ctx:
// t:
func CommandContext(t testing.TB, ctx context.Context, name string, args ...string) *exec.Cmd {
	t.Helper()
	MustHaveExec(t)

	var (
		cancelCtx   context.CancelFunc
		gracePeriod time.Duration // 除非测试有截止日期（为了允许交互式调试），否则是无限的
	)

	if t, ok := t.(interface {
		testing.TB
		Deadline() (time.Time, bool)
	}); ok {
		if td, ok := t.Deadline(); ok {
// 从一个最小的超时时间开始，足够长的时间来消耗一个合理程序终止后的输出。
			gracePeriod = 100 * time.Millisecond
			if s := os.Getenv("GO_TEST_TIMEOUT_SCALE"); s != "" {
				scale, err := strconv.Atoi(s)
				if err != nil {
					t.Fatalf("invalid GO_TEST_TIMEOUT_SCALE: %v", err)
				}
				gracePeriod *= time.Duration(scale)
			}

// 如果时间允许，将终止宽限期增加到测试剩余时间的5%。
			testTimeout := time.Until(td)
			if gp := testTimeout / 20; gp > gracePeriod {
				gracePeriod = gp
			}

// 当我们运行产生子进程的命令时，我们需要预留两个宽限期来清理：
// 一个是从首次发送终止信号（通过Cancel回调在Context超时时）到
// 进程被强制终止（通过WaitDelay字段）之间的延迟，
// 另一个是进程被终止和测试记录其输出以供调试之间的延迟。
//
// （我们希望确保在测试进程本身被终止之前有足够的时间记录输出。）
			cmdTimeout := testTimeout - 2*gracePeriod

			if cd, ok := ctx.Deadline(); !ok || time.Until(cd) > cmdTimeout {
// 如果上下文没有截止日期，或者其截止日期在测试超时时（或接近）已经过期。
// 添加一个较短的超时时间，以便测试能够产生有用的输出。
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
// 命令因 ctx 被取消而终止，但似乎不是由于我们添加的明确测试截止时间。如果这对诊断失败有帮助，请记录该信息，但不要因此实际失败测试。
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

// Command 类似于 exec.Command，但会应用与 testenv.CommandContext（使用默认的 Context）相同的变化。

// ff:
// args:
// name:
// t:
func Command(t testing.TB, name string, args ...string) *exec.Cmd {
	t.Helper()
	return CommandContext(t, context.Background(), name, args...)
}
