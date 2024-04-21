// 版权所有 2016 The Go 作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

package os_test

import (
	"fmt"
	"internal/testenv"
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

const executable_EnvVar = "OSTEST_OUTPUT_EXECPATH"

func TestExecutable(t *testing.T) {
	testenv.MustHaveExec(t)
	t.Parallel()

	ep, err := os.Executable()
	if err != nil {
		t.Fatalf("Executable failed: %v", err)
	}
	// 我们希望`fn`的形式为"dir/prog"
	dir := filepath.Dir(filepath.Dir(ep))
	fn, err := filepath.Rel(dir, ep)
	if err != nil {
		t.Fatalf("filepath.Rel: %v", err)
	}

	cmd := testenv.Command(t, fn, "-test.run=^$")
	// 让子进程以相对程序路径开始
	cmd.Dir = dir
	cmd.Path = fn
	if runtime.GOOS == "openbsd" || runtime.GOOS == "aix" {
		// OpenBSD 和 AIX 依赖于 argv[0]
	} else {
// 为子进程伪造argv[0]，以便我们可以验证在不受argv[0]影响的情况下正确获取可执行文件的真正路径。
		cmd.Args[0] = "-"
	}
	cmd.Env = append(cmd.Environ(), fmt.Sprintf("%s=1", executable_EnvVar))
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("exec(self) failed: %v", err)
	}
	outs := string(out)
	if !filepath.IsAbs(outs) {
		t.Fatalf("Child returned %q, want an absolute path", out)
	}
	if !sameFile(outs, ep) {
		t.Fatalf("Child returned %q, not the same file as %q", out, ep)
	}
}

func sameFile(fn1, fn2 string) bool {
	fi1, err := os.Stat(fn1)
	if err != nil {
		return false
	}
	fi2, err := os.Stat(fn2)
	if err != nil {
		return false
	}
	return os.SameFile(fi1, fi2)
}

func init() {
	if e := os.Getenv(executable_EnvVar); e != "" {
		// 首先切换到另一个路径
		dir := "/"
		if runtime.GOOS == "windows" {
			cwd, err := os.Getwd()
			if err != nil {
				panic(err)
			}
			dir = filepath.VolumeName(cwd)
		}
		os.Chdir(dir)
		if ep, err := os.Executable(); err != nil {
			fmt.Fprint(os.Stderr, "ERROR: ", err)
		} else {
			fmt.Fprint(os.Stderr, ep)
		}
		os.Exit(0)
	}
}

func TestExecutableDeleted(t *testing.T) {
	testenv.MustHaveGoBuild(t)
	switch runtime.GOOS {
	case "windows", "plan9":
		t.Skipf("%v does not support deleting running binary", runtime.GOOS)
	case "openbsd", "freebsd", "aix":
		t.Skipf("%v does not support reading deleted binary name", runtime.GOOS)
	}
	t.Parallel()

	dir := t.TempDir()

	src := filepath.Join(dir, "testdel.go")
	exe := filepath.Join(dir, "testdel.exe")

	err := os.WriteFile(src, []byte(testExecutableDeletion), 0666)
	if err != nil {
		t.Fatal(err)
	}

	out, err := testenv.Command(t, testenv.GoToolPath(t), "build", "-o", exe, src).CombinedOutput()
	t.Logf("build output:\n%s", out)
	if err != nil {
		t.Fatal(err)
	}

	out, err = testenv.Command(t, exe).CombinedOutput()
	t.Logf("exec output:\n%s", out)
	if err != nil {
		t.Fatal(err)
	}
}

const testExecutableDeletion = `package main

import (
	"fmt"
	"os"
)

func main() {
	before, err := os.Executable()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read executable name before deletion: %v\n", err)
		os.Exit(1)
	}

	err = os.Remove(before)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to remove executable: %v\n", err)
		os.Exit(1)
	}

	after, err := os.Executable()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read executable name after deletion: %v\n", err)
		os.Exit(1)
	}

	if before != after {
		fmt.Fprintf(os.Stderr, "before and after do not match: %v != %v\n", before, after)
		os.Exit(1)
	}
}
`
