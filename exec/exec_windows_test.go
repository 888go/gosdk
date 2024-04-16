// 版权所有 ? 2021 Go作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

//go:build windows

package exec_test

import (
	"fmt"
	"github.com/888go/gosdk/exec"
	"github.com/888go/gosdk/exec/internal/testenv"

	"io"
	"os"
	"strconv"
	"strings"
	"syscall"
	"testing"
)

var (
	quitSignal os.Signal = nil
	pipeSignal os.Signal = syscall.SIGPIPE
)

func init() {
	registerHelperCommand("pipehandle", cmdPipeHandle)
}

func cmdPipeHandle(args ...string) {
	handle, _ := strconv.ParseUint(args[0], 16, 64)
	pipe := os.NewFile(uintptr(handle), "")
	_, err := fmt.Fprint(pipe, args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "writing to pipe failed: %v\n", err)
		os.Exit(1)
	}
	pipe.Close()
}

func TestPipePassing(t *testing.T) {
	t.Parallel()

	r, w, err := os.Pipe()
	if err != nil {
		t.Error(err)
	}
	const marker = "arrakis, dune, desert planet"
	childProc := helperCommand(t, "pipehandle", strconv.FormatUint(uint64(w.Fd()), 16), marker)
	childProc.F.SysProcAttr = &syscall.SysProcAttr{AdditionalInheritedHandles: []syscall.Handle{syscall.Handle(w.Fd())}}
	err = childProc.Start()
	if err != nil {
		t.Error(err)
	}
	w.Close()
	response, err := io.ReadAll(r)
	if err != nil {
		t.Error(err)
	}
	r.Close()
	if string(response) != marker {
		t.Errorf("got %q; want %q", string(response), marker)
	}
	err = childProc.Wait()
	if err != nil {
		t.Error(err)
	}
}

func TestNoInheritHandles(t *testing.T) {
	t.Parallel()

	cmd := testenv.Command(t, "cmd", "/c exit 88")
	cmd.SysProcAttr = &syscall.SysProcAttr{NoInheritHandles: true}
	err := cmd.Run()
	exitError, ok := err.(*exec.ExitError)
	if !ok {
		t.Fatalf("got error %v; want ExitError", err)
	}
	if exitError.F.ExitCode() != 88 {
		t.Fatalf("got exit code %d; want 88", exitError.F.ExitCode())
	}
}

// 在不需用户代码明确从父进程的SYSTEMROOT开始的情况下启动子进程
// （参考问题25210）
func TestChildCriticalEnv(t *testing.T) {
	t.Parallel()
	cmd := helperCommand(t, "echoenv", "SYSTEMROOT")

	// 显式地从命令的环境变量中移除 SYSTEMROOT
	var env []string
	for _, kv := range cmd.Environ() {
		k, _, ok := strings.Cut(kv, "=")
		if !ok || !strings.EqualFold(k, "SYSTEMROOT") {
			env = append(env, kv)
		}
	}
	cmd.F.Env = env

	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatal(err)
	}
	if strings.TrimSpace(string(out)) == "" {
		t.Error("no SYSTEMROOT found")
	}
}
