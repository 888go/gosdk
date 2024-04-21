// 版权所有 2020 Go作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

//go:build unix

package os_test

import (
	"internal/testenv"
	. "os"
	"syscall"
	"testing"
)

func TestErrProcessDone(t *testing.T) {
	testenv.MustHaveGoBuild(t)
	t.Parallel()

	p, err := StartProcess(testenv.GoToolPath(t), []string{"go"}, &ProcAttr{})
	if err != nil {
		t.Fatalf("starting test process: %v", err)
	}
	p.Wait()
	if got := p.Signal(Kill); got != ErrProcessDone {
		t.Errorf("got %v want %v", got, ErrProcessDone)
	}
}

func TestUNIXProcessAlive(t *testing.T) {
	testenv.MustHaveGoBuild(t)
	t.Parallel()

	p, err := StartProcess(testenv.GoToolPath(t), []string{"sleep", "1"}, &ProcAttr{})
	if err != nil {
		t.Skipf("starting test process: %v", err)
	}
	defer p.Kill()

	proc, _ := FindProcess(p.Pid)
	err = proc.Signal(syscall.Signal(0))
	if err != nil {
		t.Errorf("OS reported error for running process: %v", err)
	}
}
