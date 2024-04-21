// 版权所有 2023 The Go Authors。保留所有权利。
// 使用本源代码须遵循 BSD 风格的许可证，
// 该许可证可在 LICENSE 文件中找到。

//go:build windows

package os_test

import (
	"internal/testenv"
	"io"
	. "os"
	"path/filepath"
	"sync"
	"testing"
)

func TestRemoveAllWithExecutedProcess(t *testing.T) {
	// 为golang.org/issue/25965问题的回归测试。
	if testing.Short() {
		t.Skip("slow test; skipping")
	}
	testenv.MustHaveExec(t)

	name, err := Executable()
	if err != nil {
		t.Fatal(err)
	}
	r, err := Open(name)
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()
	const n = 100
	var execs [n]string
	// 首先创建n个可执行文件。
	for i := 0; i < n; i++ {
		// Rewind r.
		if _, err := r.Seek(0, io.SeekStart); err != nil {
			t.Fatal(err)
		}
		name := filepath.Join(t.TempDir(), "test.exe")
		execs[i] = name
		w, err := Create(name)
		if err != nil {
			t.Fatal(err)
		}
		if _, err = io.Copy(w, r); err != nil {
			w.Close()
			t.Fatal(err)
		}
		if err := w.Sync(); err != nil {
			w.Close()
			t.Fatal(err)
		}
		if err = w.Close(); err != nil {
			t.Fatal(err)
		}
	}
// 然后运行每个可执行文件并移除其所在的目录。
// 为增加负载并提高触发问题的可能性，以单独的goroutine运行每个可执行文件。
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(i int) {
			defer wg.Done()
			name := execs[i]
			dir := filepath.Dir(name)
			// 运行test.exe而不执行任何测试，只是为了使其执行一些操作。
			cmd := testenv.Command(t, name, "-test.run=^$")
			if err := cmd.Run(); err != nil {
				t.Errorf("exec failed: %v", err)
			}
			// 删除目录，并检查它不返回“ERROR_ACCESS_DENIED”错误。
			err = RemoveAll(dir)
			if err != nil {
				t.Errorf("RemoveAll failed: %v", err)
			}
		}(i)
	}
	wg.Wait()
}
