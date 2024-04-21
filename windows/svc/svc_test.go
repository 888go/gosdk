// 版权所有 2012 The Go 作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

//go:build windows

package svc_test

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/888go/gosdk/windows/svc"
	"github.com/888go/gosdk/windows/svc/mgr"
)

func getState(t *testing.T, s *mgr.Service) svc.State {
	status, err := s.Query()
	if err != nil {
		t.Fatalf("Query(%s) failed: %s", s.Name, err)
	}
	return status.State
}

func testState(t *testing.T, s *mgr.Service, want svc.State) {
	have := getState(t, s)
	if have != want {
		t.Fatalf("%s state is=%d want=%d", s.Name, have, want)
	}
}

func waitState(t *testing.T, s *mgr.Service, want svc.State) {
	for i := 0; ; i++ {
		have := getState(t, s)
		if have == want {
			return
		}
		if i > 10 {
			t.Fatalf("%s state is=%d, waiting timeout", s.Name, have)
		}
		time.Sleep(300 * time.Millisecond)
	}
}

// stopAndDeleteIfInstalled 如果服务正在运行和/或已安装，则停止并删除名为name的服务
func stopAndDeleteIfInstalled(t *testing.T, m *mgr.Mgr, name string) {
	s, err := m.OpenService(name)
	if err != nil {
		// 服务未安装
		return

	}
	defer s.Close()

	// 确保服务未在运行，否则我们将无法删除它。
	if getState(t, s) == svc.Running {
		_, err = s.Control(svc.Stop)
		if err != nil {
			t.Fatalf("Control(%s) failed: %s", s.Name, err)
		}
		waitState(t, s, svc.Stopped)
	}

	err = s.Delete()
	if err != nil {
		t.Fatalf("Delete failed: %s", err)
	}
}

func TestExample(t *testing.T) {
	if os.Getenv("GO_BUILDER_NAME") == "" {
		// 不要在任意用户机器上安装服务
		t.Skip("skipping test that modifies system services: GO_BUILDER_NAME not set")
	}
	if testing.Short() {
		t.Skip("skipping test in short mode that modifies system services")
	}

	const name = "svctestservice"

	m, err := mgr.Connect()
	if err != nil {
		t.Fatalf("SCM connection failed: %s", err)
	}
	defer m.Disconnect()

	exepath := filepath.Join(t.TempDir(), "a.exe")
	o, err := exec.Command("go", "build", "-o", exepath, "golang.org/x/sys/windows/svc/example").CombinedOutput()
	if err != nil {
		t.Fatalf("failed to build service program: %v\n%v", err, string(o))
	}

	stopAndDeleteIfInstalled(t, m, name)

	s, err := m.CreateService(name, exepath, mgr.Config{DisplayName: "x-sys svc test service"}, "-name", name)
	if err != nil {
		t.Fatalf("CreateService(%s) failed: %v", name, err)
	}
	defer s.Close()

	args := []string{"is", "manual-started", fmt.Sprintf("%d", rand.Int())}

	testState(t, s, svc.Stopped)
	err = s.Start(args...)
	if err != nil {
		t.Fatalf("Start(%s) failed: %s", s.Name, err)
	}
	waitState(t, s, svc.Running)
	time.Sleep(1 * time.Second)

	// 测试来自问题4的死锁
	_, err = s.Control(svc.Interrogate)
	if err != nil {
		t.Fatalf("Control(%s) failed: %s", s.Name, err)
	}
	_, err = s.Control(svc.Interrogate)
	if err != nil {
		t.Fatalf("Control(%s) failed: %s", s.Name, err)
	}
	time.Sleep(1 * time.Second)

	_, err = s.Control(svc.Stop)
	if err != nil {
		t.Fatalf("Control(%s) failed: %s", s.Name, err)
	}
	waitState(t, s, svc.Stopped)

	err = s.Delete()
	if err != nil {
		t.Fatalf("Delete failed: %s", err)
	}

	out, err := exec.Command("wevtutil.exe", "qe", "Application", "/q:*[System[Provider[@Name='"+name+"']]]", "/rd:true", "/c:10").CombinedOutput()
	if err != nil {
		t.Fatalf("wevtutil failed: %v\n%v", err, string(out))
	}
	want := strings.Join(append([]string{name}, args...), "-")
	// 测试上下文传递（参见 sys_386.s 和 sys_amd64.s 中的 servicemain）。
	want += "-123456"
	if !strings.Contains(string(out), want) {
		t.Errorf("%q string does not contain %q", out, want)
	}
}

func TestIsAnInteractiveSession(t *testing.T) {
	isInteractive, err := svc.IsAnInteractiveSession()
	if err != nil {
		t.Fatal(err)
	}
	if !isInteractive {
		t.Error("IsAnInteractiveSession returns false when running interactively.")
	}
}

func TestIsWindowsService(t *testing.T) {
	isSvc, err := svc.IsWindowsService()
	if err != nil {
		t.Fatal(err)
	}
	if isSvc {
		t.Error("IsWindowsService returns true when not running in a service.")
	}
}

func TestIsWindowsServiceWhenParentExits(t *testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS") == "parent" {
		// in parent process

		// 启动子进程并快速退出
		child := exec.Command(os.Args[0], "-test.run=^TestIsWindowsServiceWhenParentExits$")
		child.Env = append(os.Environ(), "GO_WANT_HELPER_PROCESS=child")
		err := child.Start()
		if err != nil {
			fmt.Fprintf(os.Stderr, fmt.Sprintf("child start failed: %v", err))
			os.Exit(1)
		}
		os.Exit(0)
	}

	if os.Getenv("GO_WANT_HELPER_PROCESS") == "child" {
		// in child process
		dumpPath := os.Getenv("GO_WANT_HELPER_PROCESS_FILE")
		if dumpPath == "" {
// 我们无法报告这个错误。但主测试会注意到我们没有创建转储文件。
			os.Exit(1)
		}
		var msg string
		isSvc, err := svc.IsWindowsService()
		if err != nil {
			msg = err.Error()
		}
		if isSvc {
			msg = "IsWindowsService returns true when not running in a service."
		}
		err = os.WriteFile(dumpPath, []byte(msg), 0644)
		if err != nil {
// 我们无法报告这个错误。但主测试会注意到我们没有创建转储文件。
			os.Exit(2)
		}
		os.Exit(0)
	}

	// 以循环方式运行，直到出现失败为止。
	for i := 0; i < 10; i++ {
		childDumpPath := filepath.Join(t.TempDir(), "issvc.txt")

		parent := exec.Command(os.Args[0], "-test.run=^TestIsWindowsServiceWhenParentExits$")
		parent.Env = append(os.Environ(),
			"GO_WANT_HELPER_PROCESS=parent",
			"GO_WANT_HELPER_PROCESS_FILE="+childDumpPath)
		parentOutput, err := parent.CombinedOutput()
		if err != nil {
			t.Errorf("parent failed: %v: %v", err, string(parentOutput))
		}
		for i := 0; ; i++ {
			if _, err := os.Stat(childDumpPath); err == nil {
				break
			}
			time.Sleep(100 * time.Millisecond)
			if i > 10 {
				t.Fatal("timed out waiting for child output file to be created.")
			}
		}
		childOutput, err := os.ReadFile(childDumpPath)
		if err != nil {
			t.Fatalf("reading child output failed: %v", err)
		}
		if got, want := string(childOutput), ""; got != want {
			t.Fatalf("child output: want %q, got %q", want, got)
		}
	}
}
