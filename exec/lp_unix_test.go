// 版权所有 2013 The Go Authors。保留所有权利。
// 使用本源代码受 BSD 风格许可协议约束，
// 该协议可在 LICENSE 文件中找到。

//go:build unix

package exec_test

import (
	"os"
	"os/exec"
	"testing"
)

func TestLookPathUnixEmptyPath(t *testing.T) {
	// 非并行：使用了 Chdir 和 Setenv。

	tmp := t.TempDir()
	chdir(t, tmp)

	f, err := os.OpenFile("exec_me", os.O_CREATE|os.O_EXCL, 0700)
	if err != nil {
		t.Fatal("OpenFile failed: ", err)
	}
	err = f.Close()
	if err != nil {
		t.Fatal("Close failed: ", err)
	}

	t.Setenv("PATH", "")

	path, err := exec.LookPath("exec_me")
	if err == nil {
		t.Fatal("LookPath found exec_me in empty $PATH")
	}
	if path != "" {
		t.Fatalf("LookPath path == %q when err != nil", path)
	}
}
