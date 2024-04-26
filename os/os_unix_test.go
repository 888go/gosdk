// 版权所有 ? 2009 Go 作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 此许可证可在 LICENSE 文件中找到。

//go:build unix || (js && wasm) || wasip1

package os_test

import (
	. "github.com/888go/gosdk/os"
	"github.com/888go/gosdk/os/internal/testenv"
	"io"
	"path/filepath"
	"runtime"
	"strings"
	"syscall"
	"testing"
	"time"
)

func init() {
	isReadonlyError = func(err error) bool { return err == syscall.EROFS }
}

// 用于TestRawConnReadWrite测试。
type syscallDescriptor = int

func checkUidGid(t *testing.T, path string, uid, gid int) {
	dir, err := Lstat(path)
	if err != nil {
		t.Fatalf("Lstat %q (looking for uid/gid %d/%d): %s", path, uid, gid, err)
	}
	sys := dir.Sys().(*syscall.Stat_t)
	if int(sys.Uid) != uid {
		t.Errorf("Lstat %q: uid %d want %d", path, sys.Uid, uid)
	}
	if int(sys.Gid) != gid {
		t.Errorf("Lstat %q: gid %d want %d", path, sys.Gid, gid)
	}
}

func TestChown(t *testing.T) {
	if runtime.GOOS == "wasip1" {
		t.Skip("file ownership not supported on " + runtime.GOOS)
	}
	t.Parallel()

	// 使用TempDir()来确保我们在本地文件系统上，这样Getgroups返回的组ID将被允许应用于文件。在NFS上，Getgroups获取的组基本上没有用。
	f := newFile("TestChown", t)
	defer Remove(f.Name())
	defer f.Close()
	dir, err := f.Stat()
	if err != nil {
		t.Fatalf("stat %s: %s", f.Name(), err)
	}

	// 如果不是root用户，不能改变uid，但可以尝试改变gid。首先尝试使用当前组。
	gid := Getgid()
	t.Log("gid:", gid)
	if err = Chown(f.Name(), -1, gid); err != nil {
		t.Fatalf("chown %s -1 %d: %s", f.Name(), gid, err)
	}
	sys := dir.Sys().(*syscall.Stat_t)
	checkUidGid(t, f.Name(), int(sys.Uid), gid)

	// 接着尝试所有的辅助组
	groups, err := Getgroups()
	if err != nil {
		t.Fatalf("getgroups: %s", err)
	}
	t.Log("groups: ", groups)
	for _, g := range groups {
		if err = Chown(f.Name(), -1, g); err != nil {
			if testenv.SyscallIsNotSupported(err) {
				t.Logf("chown %s -1 %d: %s (error ignored)", f.Name(), g, err)
				// 由于Chown调用失败，文件应该保持未修改。
				checkUidGid(t, f.Name(), int(sys.Uid), gid)
				continue
			}
			t.Fatalf("chown %s -1 %d: %s", f.Name(), g, err)
		}
		checkUidGid(t, f.Name(), int(sys.Uid), g)

		// 改回gid以测试fd.Chown
		if err = f.Chown(-1, gid); err != nil {
			t.Fatalf("fchown %s -1 %d: %s", f.Name(), gid, err)
		}
		checkUidGid(t, f.Name(), int(sys.Uid), gid)
	}
}

func TestFileChown(t *testing.T) {
	if runtime.GOOS == "wasip1" {
		t.Skip("file ownership not supported on " + runtime.GOOS)
	}
	t.Parallel()

	// 使用TempDir()来确保我们在本地文件系统上，这样Getgroups返回的组ID将被允许应用于文件。在NFS上，Getgroups获取的组基本上没有用。
	f := newFile("TestFileChown", t)
	defer Remove(f.Name())
	defer f.Close()
	dir, err := f.Stat()
	if err != nil {
		t.Fatalf("stat %s: %s", f.Name(), err)
	}

	// 如果不是root用户，不能改变uid，但可以尝试改变gid。首先尝试使用当前组。
	gid := Getgid()
	t.Log("gid:", gid)
	if err = f.Chown(-1, gid); err != nil {
		t.Fatalf("fchown %s -1 %d: %s", f.Name(), gid, err)
	}
	sys := dir.Sys().(*syscall.Stat_t)
	checkUidGid(t, f.Name(), int(sys.Uid), gid)

	// 接着尝试所有的辅助组
	groups, err := Getgroups()
	if err != nil {
		t.Fatalf("getgroups: %s", err)
	}
	t.Log("groups: ", groups)
	for _, g := range groups {
		if err = f.Chown(-1, g); err != nil {
			if testenv.SyscallIsNotSupported(err) {
				t.Logf("chown %s -1 %d: %s (error ignored)", f.Name(), g, err)
				// 由于Chown调用失败，文件应该保持未修改。
				checkUidGid(t, f.Name(), int(sys.Uid), gid)
				continue
			}
			t.Fatalf("fchown %s -1 %d: %s", f.Name(), g, err)
		}
		checkUidGid(t, f.Name(), int(sys.Uid), g)

		// 改回gid以测试fd.Chown
		if err = f.Chown(-1, gid); err != nil {
			t.Fatalf("fchown %s -1 %d: %s", f.Name(), gid, err)
		}
		checkUidGid(t, f.Name(), int(sys.Uid), gid)
	}
}

func TestLchown(t *testing.T) {
	testenv.MustHaveSymlink(t)
	t.Parallel()

	// 使用TempDir()来确保我们在本地文件系统上，这样Getgroups返回的组ID将被允许应用于文件。在NFS上，Getgroups获取的组基本上没有用。
	f := newFile("TestLchown", t)
	defer Remove(f.Name())
	defer f.Close()
	dir, err := f.Stat()
	if err != nil {
		t.Fatalf("stat %s: %s", f.Name(), err)
	}

	linkname := f.Name() + "2"
	if err := Symlink(f.Name(), linkname); err != nil {
		if runtime.GOOS == "android" && IsPermission(err) {
			t.Skip("skipping test on Android; permission error creating symlink")
		}
		t.Fatalf("link %s -> %s: %v", f.Name(), linkname, err)
	}
	defer Remove(linkname)

	// 如果不是root用户，不能改变uid，但可以尝试改变gid。首先尝试使用当前组。
	gid := Getgid()
	t.Log("gid:", gid)
	if err = Lchown(linkname, -1, gid); err != nil {
		if err, ok := err.(*PathError); ok && err.Err == syscall.ENOSYS {
			t.Skip("lchown is unavailable")
		}
		t.Fatalf("lchown %s -1 %d: %s", linkname, gid, err)
	}
	sys := dir.Sys().(*syscall.Stat_t)
	checkUidGid(t, linkname, int(sys.Uid), gid)

	// 接着尝试所有的辅助组
	groups, err := Getgroups()
	if err != nil {
		t.Fatalf("getgroups: %s", err)
	}
	t.Log("groups: ", groups)
	for _, g := range groups {
		if err = Lchown(linkname, -1, g); err != nil {
			if testenv.SyscallIsNotSupported(err) {
				t.Logf("lchown %s -1 %d: %s (error ignored)", f.Name(), g, err)
				// 由于Lchown调用失败，文件应该未被修改。
				checkUidGid(t, f.Name(), int(sys.Uid), gid)
				continue
			}
			t.Fatalf("lchown %s -1 %d: %s", linkname, g, err)
		}
		checkUidGid(t, linkname, int(sys.Uid), g)

		// 检查链接目标的gid未发生改变
		checkUidGid(t, f.Name(), int(sys.Uid), int(sys.Gid))

		if err = Lchown(linkname, -1, gid); err != nil {
			t.Fatalf("lchown %s -1 %d: %s", f.Name(), gid, err)
		}
	}
}

// 问题16919：Readdir必须返回一个非空的切片或一个错误。
func TestReaddirRemoveRace(t *testing.T) {
	oldStat := *LstatP
	defer func() { *LstatP = oldStat }()
	*LstatP = func(name string) (FileInfo, error) {
		if strings.HasSuffix(name, "some-file") {
			// 像被删除一样表现。
			return nil, ErrNotExist
		}
		return oldStat(name)
	}
	dir := newDir("TestReaddirRemoveRace", t)
	defer RemoveAll(dir)
	if err := WriteFile(filepath.Join(dir, "some-file"), []byte("hello"), 0644); err != nil {
		t.Fatal(err)
	}
	d, err := Open(dir)
	if err != nil {
		t.Fatal(err)
	}
	defer d.Close()
	fis, err := d.Readdir(2) // 值得注意的是，大于零
	if len(fis) == 0 && err == nil {
		// 这是过去发生的情况（问题 16919）
		t.Fatal("Readdir = empty slice & err == nil")
	}
	if len(fis) != 0 || err != io.EOF {
		t.Errorf("Readdir = %d entries: %v; want 0, io.EOF", len(fis), err)
		for i, fi := range fis {
			t.Errorf("  entry[%d]: %q, %v", i, fi.Name(), fi.Mode())
		}
		t.FailNow()
	}
}

// 问题23120：在使用Mkdir时尊重umask设置，同时处理粘滞位
func TestMkdirStickyUmask(t *testing.T) {
	if runtime.GOOS == "wasip1" {
		t.Skip("file permissions not supported on " + runtime.GOOS)
	}
	t.Parallel()

	const umask = 0077
	dir := newDir("TestMkdirStickyUmask", t)
	defer RemoveAll(dir)

	oldUmask := syscall.Umask(umask)
	defer syscall.Umask(oldUmask)

	// We have set a umask, but if the parent directory happens to have a default
	// ACL, the umask may be ignored. To prevent spurious failures from an ACL,
	// we create a non-sticky directory as a “control case” to compare against our
	// sticky-bit “experiment”.
	control := filepath.Join(dir, "control")
	if err := Mkdir(control, 0755); err != nil {
		t.Fatal(err)
	}
	cfi, err := Stat(control)
	if err != nil {
		t.Fatal(err)
	}

	p := filepath.Join(dir, "dir1")
	if err := Mkdir(p, ModeSticky|0755); err != nil {
		t.Fatal(err)
	}
	fi, err := Stat(p)
	if err != nil {
		t.Fatal(err)
	}

	got := fi.Mode()
	want := cfi.Mode() | ModeSticky
	if got != want {
		t.Errorf("Mkdir(_, ModeSticky|0755) created dir with mode %v; want %v", got, want)
	}
}

// 参见问题：22939，24331
func newFileTest(t *testing.T, blocking bool) {
	if runtime.GOOS == "js" || runtime.GOOS == "wasip1" {
		t.Skipf("syscall.Pipe is not available on %s.", runtime.GOOS)
	}

	p := make([]int, 2)
	if err := syscall.Pipe(p); err != nil {
		t.Fatalf("pipe: %v", err)
	}
	defer syscall.Close(p[1])

	// 将读取侧设置为非阻塞。
	if !blocking {
		if err := syscall.SetNonblock(p[0], true); err != nil {
			syscall.Close(p[0])
			t.Fatalf("SetNonblock: %v", err)
		}
	}
	// Convert it to a file.
	file := NewFile(uintptr(p[0]), "notapipe")
	if file == nil {
		syscall.Close(p[0])
		t.Fatalf("failed to convert fd to file!")
	}
	defer file.Close()

	timeToWrite := 100 * time.Millisecond
	timeToDeadline := 1 * time.Millisecond
	if !blocking {
		// 使用较长的时间以避免出现异常。
		// 无论如何我们都不会等待这么长时间。
		timeToWrite = 1 * time.Second
	}

	// 尝试在截止时间前读取（但不要永远阻塞）。
	b := make([]byte, 1)
	timer := time.AfterFunc(timeToWrite, func() { syscall.Write(p[1], []byte("a")) })
	defer timer.Stop()
	file.SetReadDeadline(time.Now().Add(timeToDeadline))
	_, err := file.Read(b)
	if !blocking {
		// 我们希望它因为超时而失败。
		if !isDeadlineExceeded(err) {
			t.Fatalf("No timeout reading from file: %v", err)
		}
	} else {
		// 我们希望它在100毫秒后成功
		if err != nil {
			t.Fatalf("Error reading from file: %v", err)
		}
	}
}

func TestNewFileBlock(t *testing.T) {
	t.Parallel()
	newFileTest(t, true)
}

func TestNewFileNonBlock(t *testing.T) {
	t.Parallel()
	newFileTest(t, false)
}

func TestNewFileInvalid(t *testing.T) {
	t.Parallel()
	const negOne = ^uintptr(0)
	if f := NewFile(negOne, "invalid"); f != nil {
		t.Errorf("NewFile(-1) got %v want nil", f)
	}
}

func TestSplitPath(t *testing.T) {
	t.Parallel()
	for _, tt := range []struct{ path, wantDir, wantBase string }{
		{"a", ".", "a"},
		{"a/", ".", "a"},
		{"a//", ".", "a"},
		{"a/b", "a", "b"},
		{"a/b/", "a", "b"},
		{"a/b/c", "a/b", "c"},
		{"/a", "/", "a"},
		{"/a/", "/", "a"},
		{"/a/b", "/a", "b"},
		{"/a/b/", "/a", "b"},
		{"/a/b/c", "/a/b", "c"},
		{"//a", "/", "a"},
		{"//a/", "/", "a"},
		{"///a", "/", "a"},
		{"///a/", "/", "a"},
	} {
		if dir, base := SplitPath(tt.path); dir != tt.wantDir || base != tt.wantBase {
			t.Errorf("splitPath(%q) = %q, %q, want %q, %q", tt.path, dir, base, tt.wantDir, tt.wantBase)
		}
	}
}

// 测试将内容复制到以 O_APPEND 方式打开的文件中是否正常工作，以及在 Linux 系统上未使用 copy_file_range 系统调用。
//
// 作为对 go.dev/issue/60181 问题的回归测试
func TestIssue60181(t *testing.T) {
	defer chtmpdir(t)()

	want := "hello gopher"

	a, err := CreateTemp("", "a")
	if err != nil {
		t.Fatal(err)
	}
	a.WriteString(want[:5])
	a.Close()

	b, err := CreateTemp("", "b")
	if err != nil {
		t.Fatal(err)
	}
	b.WriteString(want[5:])
	b.Close()

	afd, err := syscall.Open(a.Name(), syscall.O_RDWR|syscall.O_APPEND, 0)
	if err != nil {
		t.Fatal(err)
	}

	bfd, err := syscall.Open(b.Name(), syscall.O_RDONLY, 0)
	if err != nil {
		t.Fatal(err)
	}

	aa := NewFile(uintptr(afd), a.Name())
	defer aa.Close()
	bb := NewFile(uintptr(bfd), b.Name())
	defer bb.Close()

	// 在Linux上，如果使用了copy_file_range系统调用，此操作将失败，因为它不支持以O_APPEND打开的目标文件，详情见
	// https://man7.org/linux/man-pages/man2/copy_file_range.2.html#ERRORS
	_, err = io.Copy(aa, bb)
	if err != nil {
		t.Fatal(err)
	}

	buf, err := ReadFile(aa.Name())
	if err != nil {
		t.Fatal(err)
	}

	if got := string(buf); got != want {
		t.Errorf("files not concatenated: got %q, want %q", got, want)
	}
}
