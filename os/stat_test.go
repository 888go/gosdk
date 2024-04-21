// 版权所有 2018 The Go Authors。保留所有权利。
// 使用本源代码须遵循 BSD 风格许可证，
// 该许可证可在 LICENSE 文件中找到。

package os_test

import (
	"internal/testenv"
	"io/fs"
	"os"
	"path/filepath"
	"testing"
)

// testStatAndLstat 验证所有 os.Stat、os.Lstat、os.File.Stat 和 os.Readdir 的功能。
func testStatAndLstat(t *testing.T, path string, isLink bool, statCheck, lstatCheck func(*testing.T, string, fs.FileInfo)) {
	// test os.Stat
	sfi, err := os.Stat(path)
	if err != nil {
		t.Error(err)
		return
	}
	statCheck(t, path, sfi)

	// test os.Lstat
	lsfi, err := os.Lstat(path)
	if err != nil {
		t.Error(err)
		return
	}
	lstatCheck(t, path, lsfi)

	if isLink {
		if os.SameFile(sfi, lsfi) {
			t.Errorf("stat and lstat of %q should not be the same", path)
		}
	} else {
		if !os.SameFile(sfi, lsfi) {
			t.Errorf("stat and lstat of %q should be the same", path)
		}
	}

	// test os.File.Stat
	f, err := os.Open(path)
	if err != nil {
		t.Error(err)
		return
	}
	defer f.Close()

	sfi2, err := f.Stat()
	if err != nil {
		t.Error(err)
		return
	}
	statCheck(t, path, sfi2)

	if !os.SameFile(sfi, sfi2) {
		t.Errorf("stat of open %q file and stat of %q should be the same", path, path)
	}

	if isLink {
		if os.SameFile(sfi2, lsfi) {
			t.Errorf("stat of opened %q file and lstat of %q should not be the same", path, path)
		}
	} else {
		if !os.SameFile(sfi2, lsfi) {
			t.Errorf("stat of opened %q file and lstat of %q should be the same", path, path)
		}
	}

	// 测试由 os.Readdir 返回的 fs.FileInfo
	if len(path) > 0 && os.IsPathSeparator(path[len(path)-1]) {
		// 跳过以斜杠结尾的目录的os.Readdir测试
		return
	}
	parentdir := filepath.Dir(path)
	parent, err := os.Open(parentdir)
	if err != nil {
		t.Error(err)
		return
	}
	defer parent.Close()

	fis, err := parent.Readdir(-1)
	if err != nil {
		t.Error(err)
		return
	}
	var lsfi2 fs.FileInfo
	base := filepath.Base(path)
	for _, fi2 := range fis {
		if fi2.Name() == base {
			lsfi2 = fi2
			break
		}
	}
	if lsfi2 == nil {
		t.Errorf("failed to find %q in its parent", path)
		return
	}
	lstatCheck(t, path, lsfi2)

	if !os.SameFile(lsfi, lsfi2) {
		t.Errorf("lstat of %q file in %q directory and %q should be the same", lsfi2.Name(), parentdir, path)
	}
}

// testIsDir 验证fi是否指的是目录。
func testIsDir(t *testing.T, path string, fi fs.FileInfo) {
	t.Helper()
	if !fi.IsDir() {
		t.Errorf("%q should be a directory", path)
	}
	if fi.Mode()&fs.ModeSymlink != 0 {
		t.Errorf("%q should not be a symlink", path)
	}
}

// testIsSymlink 验证fi是否引用符号链接。
func testIsSymlink(t *testing.T, path string, fi fs.FileInfo) {
	t.Helper()
	if fi.IsDir() {
		t.Errorf("%q should not be a directory", path)
	}
	if fi.Mode()&fs.ModeSymlink == 0 {
		t.Errorf("%q should be a symlink", path)
	}
}

// testIsFile 验证 fi 指向的是文件。
func testIsFile(t *testing.T, path string, fi fs.FileInfo) {
	t.Helper()
	if fi.IsDir() {
		t.Errorf("%q should not be a directory", path)
	}
	if fi.Mode()&fs.ModeSymlink != 0 {
		t.Errorf("%q should not be a symlink", path)
	}
}

func testDirStats(t *testing.T, path string) {
	testStatAndLstat(t, path, false, testIsDir, testIsDir)
}

func testFileStats(t *testing.T, path string) {
	testStatAndLstat(t, path, false, testIsFile, testIsFile)
}

func testSymlinkStats(t *testing.T, path string, isdir bool) {
	if isdir {
		testStatAndLstat(t, path, true, testIsDir, testIsSymlink)
	} else {
		testStatAndLstat(t, path, true, testIsFile, testIsSymlink)
	}
}

func testSymlinkSameFile(t *testing.T, path, link string) {
	pathfi, err := os.Stat(path)
	if err != nil {
		t.Error(err)
		return
	}

	linkfi, err := os.Stat(link)
	if err != nil {
		t.Error(err)
		return
	}
	if !os.SameFile(pathfi, linkfi) {
		t.Errorf("os.Stat(%q) and os.Stat(%q) are not the same file", path, link)
	}

	linkfi, err = os.Lstat(link)
	if err != nil {
		t.Error(err)
		return
	}
	if os.SameFile(pathfi, linkfi) {
		t.Errorf("os.Stat(%q) and os.Lstat(%q) are the same file", path, link)
	}
}

func testSymlinkSameFileOpen(t *testing.T, link string) {
	f, err := os.Open(link)
	if err != nil {
		t.Error(err)
		return
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		t.Error(err)
		return
	}

	fi2, err := os.Stat(link)
	if err != nil {
		t.Error(err)
		return
	}

	if !os.SameFile(fi, fi2) {
		t.Errorf("os.Open(%q).Stat() and os.Stat(%q) are not the same file", link, link)
	}
}

func TestDirAndSymlinkStats(t *testing.T) {
	testenv.MustHaveSymlink(t)
	t.Parallel()

	tmpdir := t.TempDir()
	dir := filepath.Join(tmpdir, "dir")
	if err := os.Mkdir(dir, 0777); err != nil {
		t.Fatal(err)
	}
	testDirStats(t, dir)

	dirlink := filepath.Join(tmpdir, "link")
	if err := os.Symlink(dir, dirlink); err != nil {
		t.Fatal(err)
	}
	testSymlinkStats(t, dirlink, true)
	testSymlinkSameFile(t, dir, dirlink)
	testSymlinkSameFileOpen(t, dirlink)

	linklink := filepath.Join(tmpdir, "linklink")
	if err := os.Symlink(dirlink, linklink); err != nil {
		t.Fatal(err)
	}
	testSymlinkStats(t, linklink, true)
	testSymlinkSameFile(t, dir, linklink)
	testSymlinkSameFileOpen(t, linklink)
}

func TestFileAndSymlinkStats(t *testing.T) {
	testenv.MustHaveSymlink(t)
	t.Parallel()

	tmpdir := t.TempDir()
	file := filepath.Join(tmpdir, "file")
	if err := os.WriteFile(file, []byte(""), 0644); err != nil {
		t.Fatal(err)
	}
	testFileStats(t, file)

	filelink := filepath.Join(tmpdir, "link")
	if err := os.Symlink(file, filelink); err != nil {
		t.Fatal(err)
	}
	testSymlinkStats(t, filelink, false)
	testSymlinkSameFile(t, file, filelink)
	testSymlinkSameFileOpen(t, filelink)

	linklink := filepath.Join(tmpdir, "linklink")
	if err := os.Symlink(filelink, linklink); err != nil {
		t.Fatal(err)
	}
	testSymlinkStats(t, linklink, false)
	testSymlinkSameFile(t, file, linklink)
	testSymlinkSameFileOpen(t, linklink)
}

// 查看问题 27225 以获取详细信息
func TestSymlinkWithTrailingSlash(t *testing.T) {
	testenv.MustHaveSymlink(t)
	t.Parallel()

	tmpdir := t.TempDir()
	dir := filepath.Join(tmpdir, "dir")
	if err := os.Mkdir(dir, 0777); err != nil {
		t.Fatal(err)
	}
	dirlink := filepath.Join(tmpdir, "link")
	if err := os.Symlink(dir, dirlink); err != nil {
		t.Fatal(err)
	}
	dirlinkWithSlash := dirlink + string(os.PathSeparator)

	testDirStats(t, dirlinkWithSlash)

	fi1, err := os.Stat(dir)
	if err != nil {
		t.Error(err)
		return
	}
	fi2, err := os.Stat(dirlinkWithSlash)
	if err != nil {
		t.Error(err)
		return
	}
	if !os.SameFile(fi1, fi2) {
		t.Errorf("os.Stat(%q) and os.Stat(%q) are not the same file", dir, dirlinkWithSlash)
	}
}
