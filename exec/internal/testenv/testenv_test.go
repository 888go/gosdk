// 版权所有 ? 2022 Go作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

package testenv_test

import (
	"github.com/888go/gosdk/exec/internal/platform"
	"github.com/888go/gosdk/exec/internal/testenv"

	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func TestGoToolLocation(t *testing.T) {
	testenv.MustHaveGoBuild(t)

	var exeSuffix string
	if runtime.GOOS == "windows" {
		exeSuffix = ".exe"
	}

// 测试定义为在其包源目录中运行，
// 而此包的源目录是 $GOROOT/src/internal/testenv。
// 'go' 命令安装在 $GOROOT/bin/go，因此如果环境设置正确，
// 则 testenv.GoTool() 应该与 ../../../bin/go 完全相同。

	relWant := "../../../bin/go" + exeSuffix
	absWant, err := filepath.Abs(relWant)
	if err != nil {
		t.Fatal(err)
	}

	wantInfo, err := os.Stat(absWant)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("found go tool at %q (%q)", relWant, absWant)

	goTool, err := testenv.GoTool()
	if err != nil {
		t.Fatalf("testenv.GoTool(): %v", err)
	}
	t.Logf("testenv.GoTool() = %q", goTool)

	gotInfo, err := os.Stat(goTool)
	if err != nil {
		t.Fatal(err)
	}
	if !os.SameFile(wantInfo, gotInfo) {
		t.Fatalf("%q is not the same file as %q", absWant, goTool)
	}
}

func TestHasGoBuild(t *testing.T) {
	if !testenv.HasGoBuild() {
		switch runtime.GOOS {
		case "js", "wasip1":
			// 无 exec 系统调用，因此它们不应能通过 'go build' 构建。
			t.Logf("HasGoBuild is false on %s", runtime.GOOS)
			return
		}

		b := testenv.Builder()
		if b == "" {
// 我们不应对外部 Go 用户可能运行的沙盒或构建环境类型做出任何假设。
			t.Skipf("skipping: 'go build' unavailable")
		}

// 由于我们掌控着Go构建器，我们知道哪些构建器应当能够运行`go build`。检查它们是否确实可以。
// 
// （注意，我们并不验证任何构建器*不能*运行`go build`。如果某个构建器在不应运行`go build`测试的情况下开始运行这些测试，我们推测当那些测试失败时，我们会发现这个问题。）
		switch runtime.GOOS {
		case "ios":
			if strings.HasSuffix(b, "-corellium") {
// Corellium 环境是自承载的，因此即使真实的“ios”设备无法执行，也应该能够进行构建。
			} else {
// 在通常的iOS沙盒环境中，应用程序不允许启动另一个进程。如果我们要在原生iOS设备上添加构建器，它们大概率将无法执行exec操作，因此我们现在不妨允许这种情况发生。
				t.Logf("HasGoBuild is false on %s", b)
				return
			}
		case "android":
			if strings.HasSuffix(b, "-emu") && platform.MustLinkExternal(runtime.GOOS, runtime.GOARCH, false) {
// 截至2023年5月2日，模拟构建器上的测试环境中缺少C链接器。
				t.Logf("HasGoBuild is false on %s", b)
				return
			}
		}

		if strings.HasSuffix(b, "-noopt") {
// -noopt构建器设置了GO_GCFLAGS，导致对'go build'的测试被跳过。
			t.Logf("HasGoBuild is false on %s", b)
			return
		}

		t.Fatalf("HasGoBuild unexpectedly false on %s", b)
	}

	t.Logf("HasGoBuild is true; checking consistency with other functions")

	hasExec := false
	hasExecGo := false
	t.Run("MustHaveExec", func(t *testing.T) {
		testenv.MustHaveExec(t)
		hasExec = true
	})
	t.Run("MustHaveExecPath", func(t *testing.T) {
		testenv.MustHaveExecPath(t, "go")
		hasExecGo = true
	})
	if !hasExec {
		t.Errorf(`MustHaveExec(t) skipped unexpectedly`)
	}
	if !hasExecGo {
		t.Errorf(`MustHaveExecPath(t, "go") skipped unexpectedly`)
	}

	dir := t.TempDir()
	mainGo := filepath.Join(dir, "main.go")
	if err := os.WriteFile(mainGo, []byte("package main\nfunc main() {}\n"), 0644); err != nil {
		t.Fatal(err)
	}
	cmd := testenv.Command(t, "go", "build", "-o", os.DevNull, mainGo)
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("%v: %v\n%s", cmd, err, out)
	}
}

func TestMustHaveExec(t *testing.T) {
	hasExec := false
	t.Run("MustHaveExec", func(t *testing.T) {
		testenv.MustHaveExec(t)
		t.Logf("MustHaveExec did not skip")
		hasExec = true
	})

	switch runtime.GOOS {
	case "js", "wasip1":
		if hasExec {
			// js and wasip1 lack an “exec” syscall.
			t.Errorf("expected MustHaveExec to skip on %v", runtime.GOOS)
		}
	case "ios":
		if b := testenv.Builder(); strings.HasSuffix(b, "-corellium") && !hasExec {
			// 大多数iOS环境无法执行，但Corellium构建器可以。
			t.Errorf("expected MustHaveExec not to skip on %v", b)
		}
	default:
		if b := testenv.Builder(); b != "" && !hasExec {
			t.Errorf("expected MustHaveExec not to skip on %v", b)
		}
	}
}

func TestCleanCmdEnvPWD(t *testing.T) {
	// 测试当 cmd.Dir 被设置时，CleanCmdEnv 是否会正确设置 PWD。
	switch runtime.GOOS {
	case "plan9", "windows":
		t.Skipf("PWD is not used on %s", runtime.GOOS)
	}
	dir := t.TempDir()
	cmd := testenv.Command(t, testenv.GoToolPath(t), "help")
	cmd.Dir = dir
	cmd = testenv.CleanCmdEnv(cmd)

	for _, env := range cmd.Env {
		if strings.HasPrefix(env, "PWD=") {
			pwd := strings.TrimPrefix(env, "PWD=")
			if pwd != dir {
				t.Errorf("unexpected PWD: want %s, got %s", dir, pwd)
			}
			return
		}
	}
	t.Error("PWD not set in cmd.Env")
}
