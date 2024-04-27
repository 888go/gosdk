// 版权所有 ? 2009 Go 作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 此许可证可在 LICENSE 文件中找到。

package os_test

import (
	"errors"
	"flag"
	"fmt"
	. "github.com/888go/gosdk/os"
	"github.com/888go/gosdk/os/internal/testenv"
	"io"
	"io/fs"
	"log"
	os2 "os"

	"path/filepath"
	"reflect"
	"runtime"
	"runtime/debug"
	"sort"
	"strings"
	"sync"
	"syscall"
	"testing"
	"testing/fstest"
	"time"
)

func TestMain(m *testing.M) {
	if Getenv("GO_OS_TEST_DRAIN_STDIN") == "1" {
		Stdout.Close()
		io.Copy(io.Discard, Stdin)
		Exit(0)
	}

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	Exit(m.Run())
}

var dot = []string{
	"dir_unix.go",
	"env.go",
	"error.go",
	"file.go",
	"os_test.go",
	"types.go",
	"stat_darwin.go",
	"stat_linux.go",
}

type sysDir struct {
	name  string
	files []string
}

var sysdir = func() *sysDir {
	switch runtime.GOOS {
	case "android":
		return &sysDir{
			"/system/lib",
			[]string{
				"libmedia.so",
				"libpowermanager.so",
			},
		}
	case "ios":
		wd, err := syscall.Getwd()
		if err != nil {
			wd = err.Error()
		}
		sd := &sysDir{
			filepath.Join(wd, "..", ".."),
			[]string{
				"ResourceRules.plist",
				"Info.plist",
			},
		}
		found := true
		for _, f := range sd.files {
			path := filepath.Join(sd.name, f)
			if _, err := Stat(path); err != nil {
				found = false
				break
			}
		}
		if found {
			return sd
		}
		// 在自托管的iOS构建中，上述文件可能不存在。请在下面查找系统文件。
	case "windows":
		return &sysDir{
			Getenv("SystemRoot") + "\\system32\\drivers\\etc",
			[]string{
				"networks",
				"protocol",
				"services",
			},
		}
	case "plan9":
		return &sysDir{
			"/lib/ndb",
			[]string{
				"common",
				"local",
			},
		}
	case "wasip1":
		// wasmtime在解析像`/etc/group`这样的目录中经常存在的符号链接时存在问题（例如，在OSX上的`private/etc/group`）。
		// 因此，我们使用Go源代码树中的文件代替。
		return &sysDir{
			runtime.GOROOT(),
			[]string{
				"go.env",
				"LICENSE",
				"CONTRIBUTING.md",
			},
		}
	}
	return &sysDir{
		"/etc",
		[]string{
			"group",
			"hosts",
			"passwd",
		},
	}
}()

func size(name string, t *testing.T) int64 {
	file, err := Open(name)
	if err != nil {
		t.Fatal("open failed:", err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			t.Error(err)
		}
	}()
	n, err := io.Copy(io.Discard, file)
	if err != nil {
		t.Fatal(err)
	}
	return n
}

func equal(name1, name2 string) (r bool) {
	switch runtime.GOOS {
	case "windows":
		r = strings.EqualFold(name1, name2)
	default:
		r = name1 == name2
	}
	return
}

// localTmp 返回一个不在NFS上的本地临时目录。
func localTmp() string {
	switch runtime.GOOS {
	case "android", "ios", "windows":
		return TempDir()
	}
	return "/tmp"
}

func newFile(testName string, t *testing.T) (f *File) {
	f, err := CreateTemp(localTmp(), "_Go_"+testName)
	if err != nil {
		t.Fatalf("TempFile %s: %s", testName, err)
	}
	return
}

func newDir(testName string, t *testing.T) (name string) {
	name, err := MkdirTemp(localTmp(), "_Go_"+testName)
	if err != nil {
		t.Fatalf("TempDir %s: %s", testName, err)
	}
	return
}

var sfdir = sysdir.name
var sfname = sysdir.files[0]

func TestStat(t *testing.T) {
	t.Parallel()

	path := sfdir + "/" + sfname
	dir, err := Stat(path)
	if err != nil {
		t.Fatal("stat failed:", err)
	}
	if !equal(sfname, dir.Name()) {
		t.Error("name should be ", sfname, "; is", dir.Name())
	}
	filesize := size(path, t)
	if dir.Size() != filesize {
		t.Error("size should be", filesize, "; is", dir.Size())
	}
}

func TestStatError(t *testing.T) {
	defer chtmpdir(t)()

	path := "no-such-file"

	fi, err := Stat(path)
	if err == nil {
		t.Fatal("got nil, want error")
	}
	if fi != nil {
		t.Errorf("got %v, want nil", fi)
	}
	if perr, ok := err.(*PathError); !ok {
		t.Errorf("got %T, want %T", err, perr)
	}

	testenv.MustHaveSymlink(t)

	link := "symlink"
	err = Symlink(path, link)
	if err != nil {
		t.Fatal(err)
	}

	fi, err = Stat(link)
	if err == nil {
		t.Fatal("got nil, want error")
	}
	if fi != nil {
		t.Errorf("got %v, want nil", fi)
	}
	if perr, ok := err.(*PathError); !ok {
		t.Errorf("got %T, want %T", err, perr)
	}
}

func TestStatSymlinkLoop(t *testing.T) {
	testenv.MustHaveSymlink(t)

	defer chtmpdir(t)()

	err := Symlink("x", "y")
	if err != nil {
		t.Fatal(err)
	}
	defer Remove("y")

	err = Symlink("y", "x")
	if err != nil {
		t.Fatal(err)
	}
	defer Remove("x")

	_, err = Stat("x")
	if _, ok := err.(*fs.PathError); !ok {
		t.Errorf("expected *PathError, got %T: %v\n", err, err)
	}
}

func TestFstat(t *testing.T) {
	t.Parallel()

	path := sfdir + "/" + sfname
	file, err1 := Open(path)
	if err1 != nil {
		t.Fatal("open failed:", err1)
	}
	defer file.Close()
	dir, err2 := file.Stat()
	if err2 != nil {
		t.Fatal("fstat failed:", err2)
	}
	if !equal(sfname, dir.Name()) {
		t.Error("name should be ", sfname, "; is", dir.Name())
	}
	filesize := size(path, t)
	if dir.Size() != filesize {
		t.Error("size should be", filesize, "; is", dir.Size())
	}
}

func TestLstat(t *testing.T) {
	t.Parallel()

	path := sfdir + "/" + sfname
	dir, err := Lstat(path)
	if err != nil {
		t.Fatal("lstat failed:", err)
	}
	if !equal(sfname, dir.Name()) {
		t.Error("name should be ", sfname, "; is", dir.Name())
	}
	if dir.Mode()&ModeSymlink == 0 {
		filesize := size(path, t)
		if dir.Size() != filesize {
			t.Error("size should be", filesize, "; is", dir.Size())
		}
	}
}

// 长度为0的读取操作不应返回EOF。
func TestRead0(t *testing.T) {
	t.Parallel()

	path := sfdir + "/" + sfname
	f, err := Open(path)
	if err != nil {
		t.Fatal("open failed:", err)
	}
	defer f.Close()

	b := make([]byte, 0)
	n, err := f.Read(b)
	if n != 0 || err != nil {
		t.Errorf("Read(0) = %d, %v, want 0, nil", n, err)
	}
	b = make([]byte, 100)
	n, err = f.Read(b)
	if n <= 0 || err != nil {
		t.Errorf("Read(100) = %d, %v, want >0, nil", n, err)
	}
}

// 读取一个已关闭的文件应该返回ErrClosed错误
func TestReadClosed(t *testing.T) {
	t.Parallel()

	path := sfdir + "/" + sfname
	file, err := Open(path)
	if err != nil {
		t.Fatal("open failed:", err)
	}
	file.Close() // close immediately

	b := make([]byte, 100)
	_, err = file.Read(b)

	e, ok := err.(*PathError)
	if !ok || e.Err != ErrClosed {
		t.Fatalf("Read: got %T(%v), want %T(%v)", err, err, e, ErrClosed)
	}
}

func testReaddirnames(dir string, contents []string) func(*testing.T) {
	return func(t *testing.T) {
		t.Parallel()

		file, err := Open(dir)
		if err != nil {
			t.Fatalf("open %q failed: %v", dir, err)
		}
		defer file.Close()
		s, err2 := file.Readdirnames(-1)
		if err2 != nil {
			t.Fatalf("Readdirnames %q failed: %v", dir, err2)
		}
		for _, m := range contents {
			found := false
			for _, n := range s {
				if n == "." || n == ".." {
					t.Errorf("got %q in directory", n)
				}
				if !equal(m, n) {
					continue
				}
				if found {
					t.Error("present twice:", m)
				}
				found = true
			}
			if !found {
				t.Error("could not find", m)
			}
		}
		if s == nil {
			t.Error("Readdirnames returned nil instead of empty slice")
		}
	}
}

func testReaddir(dir string, contents []string) func(*testing.T) {
	return func(t *testing.T) {
		t.Parallel()

		file, err := Open(dir)
		if err != nil {
			t.Fatalf("open %q failed: %v", dir, err)
		}
		defer file.Close()
		s, err2 := file.Readdir(-1)
		if err2 != nil {
			t.Fatalf("Readdir %q failed: %v", dir, err2)
		}
		for _, m := range contents {
			found := false
			for _, n := range s {
				if n.Name() == "." || n.Name() == ".." {
					t.Errorf("got %q in directory", n.Name())
				}
				if !equal(m, n.Name()) {
					continue
				}
				if found {
					t.Error("present twice:", m)
				}
				found = true
			}
			if !found {
				t.Error("could not find", m)
			}
		}
		if s == nil {
			t.Error("Readdir returned nil instead of empty slice")
		}
	}
}

func testReadDir(dir string, contents []string) func(*testing.T) {
	return func(t *testing.T) {
		t.Parallel()

		file, err := Open(dir)
		if err != nil {
			t.Fatalf("open %q failed: %v", dir, err)
		}
		defer file.Close()
		s, err2 := file.ReadDir(-1)
		if err2 != nil {
			t.Fatalf("ReadDir %q failed: %v", dir, err2)
		}
		for _, m := range contents {
			found := false
			for _, n := range s {
				if n.Name() == "." || n.Name() == ".." {
					t.Errorf("got %q in directory", n)
				}
				if !equal(m, n.Name()) {
					continue
				}
				if found {
					t.Error("present twice:", m)
				}
				found = true
				lstat, err := Lstat(dir + "/" + m)
				if err != nil {
					t.Fatal(err)
				}
				if n.IsDir() != lstat.IsDir() {
					t.Errorf("%s: IsDir=%v, want %v", m, n.IsDir(), lstat.IsDir())
				}
				if n.Type() != lstat.Mode().Type() {
					t.Errorf("%s: IsDir=%v, want %v", m, n.Type(), lstat.Mode().Type())
				}
				info, err := n.Info()
				if err != nil {
					t.Errorf("%s: Info: %v", m, err)
					continue
				}
				if !SameFile(info, lstat) {
					t.Errorf("%s: Info: SameFile(info, lstat) = false", m)
				}
			}
			if !found {
				t.Error("could not find", m)
			}
		}
		if s == nil {
			t.Error("ReadDir returned nil instead of empty slice")
		}
	}
}

func TestFileReaddirnames(t *testing.T) {
	t.Parallel()

	t.Run(".", testReaddirnames(".", dot))
	t.Run("sysdir", testReaddirnames(sysdir.name, sysdir.files))
	t.Run("TempDir", testReaddirnames(t.TempDir(), nil))
}

func TestFileReaddir(t *testing.T) {
	t.Parallel()

	t.Run(".", testReaddir(".", dot))
	t.Run("sysdir", testReaddir(sysdir.name, sysdir.files))
	t.Run("TempDir", testReaddir(t.TempDir(), nil))
}

func TestFileReadDir(t *testing.T) {
	t.Parallel()

	t.Run(".", testReadDir(".", dot))
	t.Run("sysdir", testReadDir(sysdir.name, sysdir.files))
	t.Run("TempDir", testReadDir(t.TempDir(), nil))
}

func benchmarkReaddirname(path string, b *testing.B) {
	var nentries int
	for i := 0; i < b.N; i++ {
		f, err := Open(path)
		if err != nil {
			b.Fatalf("open %q failed: %v", path, err)
		}
		ns, err := f.Readdirnames(-1)
		f.Close()
		if err != nil {
			b.Fatalf("readdirnames %q failed: %v", path, err)
		}
		nentries = len(ns)
	}
	b.Logf("benchmarkReaddirname %q: %d entries", path, nentries)
}

func benchmarkReaddir(path string, b *testing.B) {
	var nentries int
	for i := 0; i < b.N; i++ {
		f, err := Open(path)
		if err != nil {
			b.Fatalf("open %q failed: %v", path, err)
		}
		fs, err := f.Readdir(-1)
		f.Close()
		if err != nil {
			b.Fatalf("readdir %q failed: %v", path, err)
		}
		nentries = len(fs)
	}
	b.Logf("benchmarkReaddir %q: %d entries", path, nentries)
}

func benchmarkReadDir(path string, b *testing.B) {
	var nentries int
	for i := 0; i < b.N; i++ {
		f, err := Open(path)
		if err != nil {
			b.Fatalf("open %q failed: %v", path, err)
		}
		fs, err := f.ReadDir(-1)
		f.Close()
		if err != nil {
			b.Fatalf("readdir %q failed: %v", path, err)
		}
		nentries = len(fs)
	}
	b.Logf("benchmarkReadDir %q: %d entries", path, nentries)
}

func BenchmarkReaddirname(b *testing.B) {
	benchmarkReaddirname(".", b)
}

func BenchmarkReaddir(b *testing.B) {
	benchmarkReaddir(".", b)
}

func BenchmarkReadDir(b *testing.B) {
	benchmarkReadDir(".", b)
}

func benchmarkStat(b *testing.B, path string) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := Stat(path)
		if err != nil {
			b.Fatalf("Stat(%q) failed: %v", path, err)
		}
	}
}

func benchmarkLstat(b *testing.B, path string) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := Lstat(path)
		if err != nil {
			b.Fatalf("Lstat(%q) failed: %v", path, err)
		}
	}
}

func BenchmarkStatDot(b *testing.B) {
	benchmarkStat(b, ".")
}

func BenchmarkStatFile(b *testing.B) {
	benchmarkStat(b, filepath.Join(runtime.GOROOT(), "src/os/os_test.go"))
}

func BenchmarkStatDir(b *testing.B) {
	benchmarkStat(b, filepath.Join(runtime.GOROOT(), "src/os"))
}

func BenchmarkLstatDot(b *testing.B) {
	benchmarkLstat(b, ".")
}

func BenchmarkLstatFile(b *testing.B) {
	benchmarkLstat(b, filepath.Join(runtime.GOROOT(), "src/os/os_test.go"))
}

func BenchmarkLstatDir(b *testing.B) {
	benchmarkLstat(b, filepath.Join(runtime.GOROOT(), "src/os"))
}

// 一次读取目录中的一个条目。
func smallReaddirnames(file *File, length int, t *testing.T) []string {
	names := make([]string, length)
	count := 0
	for {
		d, err := file.Readdirnames(1)
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatalf("readdirnames %q failed: %v", file.Name(), err)
		}
		if len(d) == 0 {
			t.Fatalf("readdirnames %q returned empty slice and no error", file.Name())
		}
		names[count] = d[0]
		count++
	}
	return names[0:count]
}

// 检查一次读取目录一个条目是否与一次性读取整个目录的结果相同。
func TestReaddirnamesOneAtATime(t *testing.T) {
	t.Parallel()

	// 大型目录，不常变动
	dir := "/usr/bin"
	switch runtime.GOOS {
	case "android":
		dir = "/system/bin"
	case "ios", "wasip1":
		wd, err := Getwd()
		if err != nil {
			t.Fatal(err)
		}
		dir = wd
	case "plan9":
		dir = "/bin"
	case "windows":
		dir = Getenv("SystemRoot") + "\\system32"
	}
	file, err := Open(dir)
	if err != nil {
		t.Fatalf("open %q failed: %v", dir, err)
	}
	defer file.Close()
	all, err1 := file.Readdirnames(-1)
	if err1 != nil {
		t.Fatalf("readdirnames %q failed: %v", dir, err1)
	}
	file1, err2 := Open(dir)
	if err2 != nil {
		t.Fatalf("open %q failed: %v", dir, err2)
	}
	defer file1.Close()
	small := smallReaddirnames(file1, len(all)+100, t) // +100 以防我们出错
	if len(small) < len(all) {
		t.Fatalf("len(small) is %d, less than %d", len(small), len(all))
	}
	for i, n := range all {
		if small[i] != n {
			t.Errorf("small read %q mismatch: %v", small[i], n)
		}
	}
}

func TestReaddirNValues(t *testing.T) {
	if testing.Short() {
		t.Skip("test.short; skipping")
	}
	t.Parallel()

	dir := t.TempDir()
	for i := 1; i <= 105; i++ {
		f, err := Create(filepath.Join(dir, fmt.Sprintf("%d", i)))
		if err != nil {
			t.Fatalf("Create: %v", err)
		}
		f.Write([]byte(strings.Repeat("X", i)))
		f.Close()
	}

	var d *File
	openDir := func() {
		var err error
		d, err = Open(dir)
		if err != nil {
			t.Fatalf("Open directory: %v", err)
		}
	}

	readdirExpect := func(n, want int, wantErr error) {
		t.Helper()
		fi, err := d.Readdir(n)
		if err != wantErr {
			t.Fatalf("Readdir of %d got error %v, want %v", n, err, wantErr)
		}
		if g, e := len(fi), want; g != e {
			t.Errorf("Readdir of %d got %d files, want %d", n, g, e)
		}
	}

	readDirExpect := func(n, want int, wantErr error) {
		t.Helper()
		de, err := d.ReadDir(n)
		if err != wantErr {
			t.Fatalf("ReadDir of %d got error %v, want %v", n, err, wantErr)
		}
		if g, e := len(de), want; g != e {
			t.Errorf("ReadDir of %d got %d files, want %d", n, g, e)
		}
	}

	readdirnamesExpect := func(n, want int, wantErr error) {
		t.Helper()
		fi, err := d.Readdirnames(n)
		if err != wantErr {
			t.Fatalf("Readdirnames of %d got error %v, want %v", n, err, wantErr)
		}
		if g, e := len(fi), want; g != e {
			t.Errorf("Readdirnames of %d got %d files, want %d", n, g, e)
		}
	}

	for _, fn := range []func(int, int, error){readdirExpect, readdirnamesExpect, readDirExpect} {
		// Test the slurp case
		openDir()
		fn(0, 105, nil)
		fn(0, 0, nil)
		d.Close()

		// Slurp with -1 instead
		openDir()
		fn(-1, 105, nil)
		fn(-2, 0, nil)
		fn(0, 0, nil)
		d.Close()

		// Test the bounded case
		openDir()
		fn(1, 1, nil)
		fn(2, 2, nil)
		fn(105, 102, nil) // 并且测试缓冲区大小大于100的情况
		fn(3, 0, io.EOF)
		d.Close()
	}
}

func touch(t *testing.T, name string) {
	f, err := Create(name)
	if err != nil {
		t.Fatal(err)
	}
	if err := f.Close(); err != nil {
		t.Fatal(err)
	}
}

func TestReaddirStatFailures(t *testing.T) {
	switch runtime.GOOS {
	case "windows", "plan9":
		// Windows和Plan 9已经正确地完成了这个操作，
		// 但是它们的系统调用结构不同，
		// 所以它们不使用Lstat，因此下面用于测试的钩子将无法工作。
		t.Skipf("skipping test on %v", runtime.GOOS)
	}

	var xerr error // error to return for x
	*LstatP = func(path string) (FileInfo, error) {
		if xerr != nil && strings.HasSuffix(path, "x") {
			return nil, xerr
		}
		return Lstat(path)
	}
	defer func() { *LstatP = Lstat }()

	dir := t.TempDir()
	touch(t, filepath.Join(dir, "good1"))
	touch(t, filepath.Join(dir, "x")) // 将消失或出现错误
	touch(t, filepath.Join(dir, "good2"))
	readDir := func() ([]FileInfo, error) {
		d, err := Open(dir)
		if err != nil {
			t.Fatal(err)
		}
		defer d.Close()
		return d.Readdir(-1)
	}
	mustReadDir := func(testName string) []FileInfo {
		fis, err := readDir()
		if err != nil {
			t.Fatalf("%s: Readdir: %v", testName, err)
		}
		return fis
	}
	names := func(fis []FileInfo) []string {
		s := make([]string, len(fis))
		for i, fi := range fis {
			s[i] = fi.Name()
		}
		sort.Strings(s)
		return s
	}

	if got, want := names(mustReadDir("initial readdir")),
		[]string{"good1", "good2", "x"}; !reflect.DeepEqual(got, want) {
		t.Errorf("initial readdir got %q; want %q", got, want)
	}

	xerr = ErrNotExist
	if got, want := names(mustReadDir("with x disappearing")),
		[]string{"good1", "good2"}; !reflect.DeepEqual(got, want) {
		t.Errorf("with x disappearing, got %q; want %q", got, want)
	}

	xerr = errors.New("some real error")
	if _, err := readDir(); err != xerr {
		t.Errorf("with a non-ErrNotExist error, got error %v; want %v", err, xerr)
	}
}

// 对普通文件执行Readdir应该失败。
func TestReaddirOfFile(t *testing.T) {
	t.Parallel()

	f, err := CreateTemp(t.TempDir(), "_Go_ReaddirOfFile")
	if err != nil {
		t.Fatal(err)
	}
	f.Write([]byte("foo"))
	f.Close()
	reg, err := Open(f.Name())
	if err != nil {
		t.Fatal(err)
	}
	defer reg.Close()

	names, err := reg.Readdirnames(-1)
	if err == nil {
		t.Error("Readdirnames succeeded; want non-nil error")
	}
	var pe *PathError
	if !errors.As(err, &pe) || pe.Path != f.Name() {
		t.Errorf("Readdirnames returned %q; want a PathError with path %q", err, f.Name())
	}
	if len(names) > 0 {
		t.Errorf("unexpected dir names in regular file: %q", names)
	}
}

func TestHardLink(t *testing.T) {
	testenv.MustHaveLink(t)

	defer chtmpdir(t)()
	from, to := "hardlinktestfrom", "hardlinktestto"
	file, err := Create(to)
	if err != nil {
		t.Fatalf("open %q failed: %v", to, err)
	}
	if err = file.Close(); err != nil {
		t.Errorf("close %q failed: %v", to, err)
	}
	err = Link(to, from)
	if err != nil {
		t.Fatalf("link %q, %q failed: %v", to, from, err)
	}

	none := "hardlinktestnone"
	err = Link(none, none)
	// 检查返回的错误是否有效。
	if lerr, ok := err.(*LinkError); !ok || lerr.Error() == "" {
		t.Errorf("link %q, %q failed to return a valid error", none, none)
	}

	tostat, err := Stat(to)
	if err != nil {
		t.Fatalf("stat %q failed: %v", to, err)
	}
	fromstat, err := Stat(from)
	if err != nil {
		t.Fatalf("stat %q failed: %v", from, err)
	}
	if !SameFile(tostat, fromstat) {
		t.Errorf("link %q, %q did not create hard link", to, from)
	}
	// 我们应该无法再次执行Link()操作
	err = Link(to, from)
	switch err := err.(type) {
	case *LinkError:
		if err.F.Op != "link" {
			t.Errorf("Link(%q, %q) err.Op = %q; want %q", to, from, err.F.Op, "link")
		}
		if err.F.Old != to {
			t.Errorf("Link(%q, %q) err.Old = %q; want %q", to, from, err.F.Old, to)
		}
		if err.F.New != from {
			t.Errorf("Link(%q, %q) err.New = %q; want %q", to, from, err.F.New, from)
		}
		if !IsExist(err.F.Err) {
			t.Errorf("Link(%q, %q) err.Err = %q; want %q", to, from, err.F.Err, "file exists error")
		}
	case nil:
		t.Errorf("link %q, %q: expected error, got nil", from, to)
	default:
		t.Errorf("link %q, %q: expected %T, got %T %v", from, to, new(LinkError), err, err)
	}
}

// chtmpdir 将工作目录更改为新的临时目录并提供一个清理函数。
func chtmpdir(t *testing.T) func() {
	oldwd, err := Getwd()
	if err != nil {
		t.Fatalf("chtmpdir: %v", err)
	}
	d, err := MkdirTemp("", "test")
	if err != nil {
		t.Fatalf("chtmpdir: %v", err)
	}
	if err := Chdir(d); err != nil {
		t.Fatalf("chtmpdir: %v", err)
	}
	return func() {
		if err := Chdir(oldwd); err != nil {
			t.Fatalf("chtmpdir: %v", err)
		}
		RemoveAll(d)
	}
}

func TestSymlink(t *testing.T) {
	testenv.MustHaveSymlink(t)

	defer chtmpdir(t)()
	from, to := "symlinktestfrom", "symlinktestto"
	file, err := Create(to)
	if err != nil {
		t.Fatalf("Create(%q) failed: %v", to, err)
	}
	if err = file.Close(); err != nil {
		t.Errorf("Close(%q) failed: %v", to, err)
	}
	err = Symlink(to, from)
	if err != nil {
		t.Fatalf("Symlink(%q, %q) failed: %v", to, from, err)
	}
	tostat, err := Lstat(to)
	if err != nil {
		t.Fatalf("Lstat(%q) failed: %v", to, err)
	}
	if tostat.Mode()&ModeSymlink != 0 {
		t.Fatalf("Lstat(%q).Mode()&ModeSymlink = %v, want 0", to, tostat.Mode()&ModeSymlink)
	}
	fromstat, err := Stat(from)
	if err != nil {
		t.Fatalf("Stat(%q) failed: %v", from, err)
	}
	if !SameFile(tostat, fromstat) {
		t.Errorf("Symlink(%q, %q) did not create symlink", to, from)
	}
	fromstat, err = Lstat(from)
	if err != nil {
		t.Fatalf("Lstat(%q) failed: %v", from, err)
	}
	if fromstat.Mode()&ModeSymlink == 0 {
		t.Fatalf("Lstat(%q).Mode()&ModeSymlink = 0, want %v", from, ModeSymlink)
	}
	fromstat, err = Stat(from)
	if err != nil {
		t.Fatalf("Stat(%q) failed: %v", from, err)
	}
	if fromstat.Name() != from {
		t.Errorf("Stat(%q).Name() = %q, want %q", from, fromstat.Name(), from)
	}
	if fromstat.Mode()&ModeSymlink != 0 {
		t.Fatalf("Stat(%q).Mode()&ModeSymlink = %v, want 0", from, fromstat.Mode()&ModeSymlink)
	}
	s, err := Readlink(from)
	if err != nil {
		t.Fatalf("Readlink(%q) failed: %v", from, err)
	}
	if s != to {
		t.Fatalf("Readlink(%q) = %q, want %q", from, s, to)
	}
	file, err = Open(from)
	if err != nil {
		t.Fatalf("Open(%q) failed: %v", from, err)
	}
	file.Close()
}

func TestLongSymlink(t *testing.T) {
	testenv.MustHaveSymlink(t)

	defer chtmpdir(t)()
	s := "0123456789abcdef"
	// 长度适中，但不要过长：常见的限制是255个字符。
	s = s + s + s + s + s + s + s + s + s + s + s + s + s + s + s
	from := "longsymlinktestfrom"
	err := Symlink(s, from)
	if err != nil {
		t.Fatalf("symlink %q, %q failed: %v", s, from, err)
	}
	r, err := Readlink(from)
	if err != nil {
		t.Fatalf("readlink %q failed: %v", from, err)
	}
	if r != s {
		t.Fatalf("after symlink %q != %q", r, s)
	}
}

func TestRename(t *testing.T) {
	defer chtmpdir(t)()
	from, to := "renamefrom", "renameto"

	file, err := Create(from)
	if err != nil {
		t.Fatalf("open %q failed: %v", from, err)
	}
	if err = file.Close(); err != nil {
		t.Errorf("close %q failed: %v", from, err)
	}
	err = Rename(from, to)
	if err != nil {
		t.Fatalf("rename %q, %q failed: %v", to, from, err)
	}
	_, err = Stat(to)
	if err != nil {
		t.Errorf("stat %q failed: %v", to, err)
	}
}

func TestRenameOverwriteDest(t *testing.T) {
	defer chtmpdir(t)()
	from, to := "renamefrom", "renameto"

	toData := []byte("to")
	fromData := []byte("from")

	err := WriteFile(to, toData, 0777)
	if err != nil {
		t.Fatalf("write file %q failed: %v", to, err)
	}

	err = WriteFile(from, fromData, 0777)
	if err != nil {
		t.Fatalf("write file %q failed: %v", from, err)
	}
	err = Rename(from, to)
	if err != nil {
		t.Fatalf("rename %q, %q failed: %v", to, from, err)
	}

	_, err = Stat(from)
	if err == nil {
		t.Errorf("from file %q still exists", from)
	}
	if err != nil && !IsNotExist(err) {
		t.Fatalf("stat from: %v", err)
	}
	toFi, err := Stat(to)
	if err != nil {
		t.Fatalf("stat %q failed: %v", to, err)
	}
	if toFi.Size() != int64(len(fromData)) {
		t.Errorf(`"to" size = %d; want %d (old "from" size)`, toFi.Size(), len(fromData))
	}
}

func TestRenameFailed(t *testing.T) {
	defer chtmpdir(t)()
	from, to := "renamefrom", "renameto"

	err := Rename(from, to)
	switch err := err.(type) {
	case *LinkError:
		if err.F.Op != "rename" {
			t.Errorf("rename %q, %q: err.Op: want %q, got %q", from, to, "rename", err.F.Op)
		}
		if err.F.Old != from {
			t.Errorf("rename %q, %q: err.Old: want %q, got %q", from, to, from, err.F.Old)
		}
		if err.F.New != to {
			t.Errorf("rename %q, %q: err.New: want %q, got %q", from, to, to, err.F.New)
		}
	case nil:
		t.Errorf("rename %q, %q: expected error, got nil", from, to)
	default:
		t.Errorf("rename %q, %q: expected %T, got %T %v", from, to, new(LinkError), err, err)
	}
}

func TestRenameNotExisting(t *testing.T) {
	defer chtmpdir(t)()
	from, to := "doesnt-exist", "dest"

	Mkdir(to, 0777)

	if err := Rename(from, to); !IsNotExist(err) {
		t.Errorf("Rename(%q, %q) = %v; want an IsNotExist error", from, to, err)
	}
}

func TestRenameToDirFailed(t *testing.T) {
	defer chtmpdir(t)()
	from, to := "renamefrom", "renameto"

	Mkdir(from, 0777)
	Mkdir(to, 0777)

	err := Rename(from, to)
	switch err := err.(type) {
	case *LinkError:
		if err.F.Op != "rename" {
			t.Errorf("rename %q, %q: err.Op: want %q, got %q", from, to, "rename", err.F.Op)
		}
		if err.F.Old != from {
			t.Errorf("rename %q, %q: err.Old: want %q, got %q", from, to, from, err.F.Old)
		}
		if err.F.New != to {
			t.Errorf("rename %q, %q: err.New: want %q, got %q", from, to, to, err.F.New)
		}
	case nil:
		t.Errorf("rename %q, %q: expected error, got nil", from, to)
	default:
		t.Errorf("rename %q, %q: expected %T, got %T %v", from, to, new(LinkError), err, err)
	}
}

func TestRenameCaseDifference(pt *testing.T) {
	from, to := "renameFROM", "RENAMEfrom"
	tests := []struct {
		name   string
		create func() error
	}{
		{"dir", func() error {
			return Mkdir(from, 0777)
		}},
		{"file", func() error {
			fd, err := Create(from)
			if err != nil {
				return err
			}
			return fd.Close()
		}},
	}

	for _, test := range tests {
		pt.Run(test.name, func(t *testing.T) {
			defer chtmpdir(t)()

			if err := test.create(); err != nil {
				t.Fatalf("failed to create test file: %s", err)
			}

			if _, err := Stat(to); err != nil {
				// 对底层文件系统是否区分大小写进行合理性检查。
				if IsNotExist(err) {
					t.Skipf("case sensitive filesystem")
				}
				t.Fatalf("stat %q, got: %q", to, err)
			}

			if err := Rename(from, to); err != nil {
				t.Fatalf("unexpected error when renaming from %q to %q: %s", from, to, err)
			}

			fd, err := Open(".")
			if err != nil {
				t.Fatalf("Open .: %s", err)
			}

			// Stat方法不会返回文件的实际大小（它返回的是调用者所请求的内容），所以我们必须使用readdir来获取文件的真正名称。
			dirNames, err := fd.Readdirnames(-1)
			fd.Close()
			if err != nil {
				t.Fatalf("readdirnames: %s", err)
			}

			if dirNamesLen := len(dirNames); dirNamesLen != 1 {
				t.Fatalf("unexpected dirNames len, got %q, want %q", dirNamesLen, 1)
			}

			if dirNames[0] != to {
				t.Errorf("unexpected name, got %q, want %q", dirNames[0], to)
			}
		})
	}
}

//func testStartProcess(dir, cmd string, args []string, expect string) func(t *testing.T) {
//	return func(t *testing.T) {
//		t.Parallel()
//
//		r, w, err := Pipe()
//		if err != nil {
//			t.Fatalf("Pipe: %v", err)
//		}
//		defer r.Close()
//
//		attr := &ProcAttr{Dir: dir, Files: []*File{nil, w, Stderr}}
//		p, err := StartProcess(cmd, args, attr)
//		if err != nil {
//			t.Fatalf("StartProcess: %v", err)
//		}
//		w.Close()
//
//		var b strings.Builder
//		io.Copy(&b, r)
//		output := b.String()
//
//		fi1, _ := Stat(strings.TrimSpace(output))
//		fi2, _ := Stat(expect)
//		if !SameFile(fi1, fi2) {
//			t.Errorf("exec %q returned %q wanted %q",
//				strings.Join(append([]string{cmd}, args...), " "), output, expect)
//		}
//		p.Wait()
//	}
//}

// ```go
//func TestStartProcess(t *testing.T) {
// 测试环境必须支持执行命令
//testenv.MustHaveExec(t)
// 并行执行测试
//t.Parallel()
//
// 根据操作系统设置变量
//var dir, cmd string
//var args []string
// 判断当前操作系统
//switch runtime.GOOS {
// 如果是Android，跳过测试（因为没有/bin/pwd）
//case "android":
//	t.Skip("android doesn't have /bin/pwd")
// 如果是Windows
//case "windows":
// 设置cmd为环境变量COMSPEC的值
//	cmd = Getenv("COMSPEC")
// 设置dir为环境变量SystemRoot的值
//	dir = Getenv("SystemRoot")
// 设置args为/c和cd
//	args = []string{"/c", "cd"}
// 其他情况（默认为非Windows系统）
//default:
// 使用LookPath查找"pwd"命令
//	var err error
//	cmd, err = exec.LookPath("pwd")
// 如果找不到"pwd"，终止测试
//	if err != nil {
//		t.Fatalf("Can't find pwd: %v", err)
//	}
// 设置dir为"/"
//	dir = "/"
// 设置args为空
//	args = []string{}
// 打印正在使用的命令进行测试
//	t.Logf("Testing with %v", cmd)
//}
//
// 分离cmd的目录和基础命令
//cmddir, cmdbase := filepath.Split(cmd)
// 将基础命令添加到args列表的前面
//args = append([]string{cmdbase}, args...)
// 运行绝对路径测试
//t.Run("absolute", testStartProcess(dir, cmd, args, dir))
// 运行相对路径测试
//t.Run("relative", testStartProcess(cmddir, cmdbase, args, cmddir))
//}
// ```
//
// 这段代码是一个Go语言的测试函数，用于测试`StartProcess`功能。它根据不同的操作系统设置命令和参数，并分别测试绝对路径和相对路径的情况。

func checkMode(t *testing.T, path string, mode FileMode) {
	dir, err := Stat(path)
	if err != nil {
		t.Fatalf("Stat %q (looking for mode %#o): %s", path, mode, err)
	}
	if dir.Mode()&ModePerm != mode {
		t.Errorf("Stat %q: mode %#o want %#o", path, dir.Mode(), mode)
	}
}

func TestChmod(t *testing.T) {
	// Chmod 在 wasip1 上不受支持。
	if runtime.GOOS == "wasip1" {
		t.Skip("Chmod is not supported on " + runtime.GOOS)
	}
	t.Parallel()

	f := newFile("TestChmod", t)
	defer Remove(f.Name())
	defer f.Close()
	// 创建模式为读写

	fm := FileMode(0456)
	if runtime.GOOS == "windows" {
		fm = FileMode(0444) // read-only file
	}
	if err := Chmod(f.Name(), fm); err != nil {
		t.Fatalf("chmod %s %#o: %s", f.Name(), fm, err)
	}
	checkMode(t, f.Name(), fm)

	fm = FileMode(0123)
	if runtime.GOOS == "windows" {
		fm = FileMode(0666) // read-write file
	}
	if err := f.Chmod(fm); err != nil {
		t.Fatalf("chmod %s %#o: %s", f.Name(), fm, err)
	}
	checkMode(t, f.Name(), fm)
}

func checkSize(t *testing.T, f *File, size int64) {
	t.Helper()
	dir, err := f.Stat()
	if err != nil {
		t.Fatalf("Stat %q (looking for size %d): %s", f.Name(), size, err)
	}
	if dir.Size() != size {
		t.Errorf("Stat %q: size %d want %d", f.Name(), dir.Size(), size)
	}
}

func TestFTruncate(t *testing.T) {
	t.Parallel()

	f := newFile("TestFTruncate", t)
	defer Remove(f.Name())
	defer f.Close()

	checkSize(t, f, 0)
	f.Write([]byte("hello, world\n"))
	checkSize(t, f, 13)
	f.Truncate(10)
	checkSize(t, f, 10)
	f.Truncate(1024)
	checkSize(t, f, 1024)
	f.Truncate(0)
	checkSize(t, f, 0)
	_, err := f.Write([]byte("surprise!"))
	if err == nil {
		checkSize(t, f, 13+9) // 在“hello, world”之后的偏移处写入。
	}
}

func TestTruncate(t *testing.T) {
	t.Parallel()

	f := newFile("TestTruncate", t)
	defer Remove(f.Name())
	defer f.Close()

	checkSize(t, f, 0)
	f.Write([]byte("hello, world\n"))
	checkSize(t, f, 13)
	Truncate(f.Name(), 10)
	checkSize(t, f, 10)
	Truncate(f.Name(), 1024)
	checkSize(t, f, 1024)
	Truncate(f.Name(), 0)
	checkSize(t, f, 0)
	_, err := f.Write([]byte("surprise!"))
	if err == nil {
		checkSize(t, f, 13+9) // 在“hello, world”之后的偏移处写入。
	}
}

func TestTruncateNonexistentFile(t *testing.T) {
	t.Parallel()

	assertPathError := func(t testing.TB, path string, err error) {
		t.Helper()
		if pe, ok := err.(*PathError); !ok || !IsNotExist(err) || pe.Path != path {
			t.Errorf("got error: %v\nwant an ErrNotExist PathError with path %q", err, path)
		}
	}

	path := filepath.Join(t.TempDir(), "nonexistent")

	err := Truncate(path, 1)
	assertPathError(t, path, err)

	// Truncate不应该创建任何新的文件。
	_, err = Stat(path)
	assertPathError(t, path, err)
}

// 通过使用 TempDir（经由 newFile）确保我们位于本地文件系统上，
// 这样时间测量就不会因延迟和缓存而失真。
// 在 NFS（网络文件系统）上，由于 NFS 服务器上元数据的缓存问题（Issue 848），时间测量可能会出现偏差。
func TestChtimes(t *testing.T) {
	t.Parallel()

	f := newFile("TestChtimes", t)
	defer Remove(f.Name())

	f.Write([]byte("hello, world\n"))
	f.Close()

	testChtimes(t, f.Name())
}

func TestChtimesWithZeroTimes(t *testing.T) {
	file := newFile("chtimes-with-zero", t)
	_, err := file.Write([]byte("hello, world\n"))
	if err != nil {
		t.Fatalf("Write: %s", err)
	}
	fName := file.Name()
	defer Remove(file.Name())
	err = file.Close()
	if err != nil {
		t.Errorf("%v", err)
	}
	fs, err := Stat(fName)
	if err != nil {
		t.Fatal(err)
	}
	startAtime := Atime(fs)
	startMtime := fs.ModTime()
	switch runtime.GOOS {
	case "js":
		startAtime = startAtime.Truncate(time.Second)
		startMtime = startMtime.Truncate(time.Second)
	}
	at0 := startAtime
	mt0 := startMtime
	t0 := startMtime.Truncate(time.Second).Add(1 * time.Hour)

	tests := []struct {
		aTime     time.Time
		mTime     time.Time
		wantATime time.Time
		wantMTime time.Time
	}{
		{
			aTime:     time.Time{},
			mTime:     time.Time{},
			wantATime: startAtime,
			wantMTime: startMtime,
		},
		{
			aTime:     t0.Add(200 * time.Second),
			mTime:     time.Time{},
			wantATime: t0.Add(200 * time.Second),
			wantMTime: startMtime,
		},
		{
			aTime:     time.Time{},
			mTime:     t0.Add(100 * time.Second),
			wantATime: t0.Add(200 * time.Second),
			wantMTime: t0.Add(100 * time.Second),
		},
		{
			aTime:     t0.Add(300 * time.Second),
			mTime:     t0.Add(100 * time.Second),
			wantATime: t0.Add(300 * time.Second),
			wantMTime: t0.Add(100 * time.Second),
		},
	}

	for _, tt := range tests {
		// 现在根据需要修改时间。
		if err := Chtimes(fName, tt.aTime, tt.mTime); err != nil {
			t.Error(err)
		}

		// 最后验证预期。
		fs, err = Stat(fName)
		if err != nil {
			t.Error(err)
		}
		at0 = Atime(fs)
		mt0 = fs.ModTime()

		if got, want := at0, tt.wantATime; !got.Equal(want) {
			errormsg := fmt.Sprintf("AccessTime mismatch with values ATime:%q-MTime:%q\ngot:  %q\nwant: %q", tt.aTime, tt.mTime, got, want)
			switch runtime.GOOS {
			case "plan9":
				// Mtime是上次内容更改的时间。类似地，每当内容被访问时，atime也会被设置；同时，当mtime被设置时，它也会被设置。
			case "windows":
				t.Error(errormsg)
			default: // unix's
				if got, want := at0, tt.wantATime; !got.Equal(want) {
					mounts, err := ReadFile("/bin/mounts")
					if err != nil {
						mounts, err = ReadFile("/etc/mtab")
					}
					if strings.Contains(string(mounts), "noatime") {
						t.Log(errormsg)
						t.Log("A filesystem is mounted with noatime; ignoring.")
					} else {
						switch runtime.GOOS {
						case "netbsd", "dragonfly":
							// 在64位实现中，通常支持并不可更改出生时间（birth time）。
							// 若得到支持，访问时间（atime）的更新则受到限制，具体取决于文件系统以及操作系统配置。
							if strings.Contains(runtime.GOARCH, "64") {
								t.Log(errormsg)
								t.Log("Filesystem might not support atime changes; ignoring.")
							}
						default:
							t.Error(errormsg)
						}
					}
				}
			}
		}
		if got, want := mt0, tt.wantMTime; !got.Equal(want) {
			errormsg := fmt.Sprintf("ModTime mismatch with values ATime:%q-MTime:%q\ngot:  %q\nwant: %q", tt.aTime, tt.mTime, got, want)
			switch runtime.GOOS {
			case "dragonfly":
				t.Log(errormsg)
				t.Log("Mtime is always updated; ignoring.")
			default:
				t.Error(errormsg)
			}
		}
	}
}

// 使用TempDir（通过newDir）确保我们在本地文件系统上，
// 这样时间测量就不会因延迟和缓存而失真。
// 在NFS上，由于NFS服务器对元数据的缓存，时间可能会有偏差（问题848）。
func TestChtimesDir(t *testing.T) {
	t.Parallel()

	name := newDir("TestChtimes", t)
	defer RemoveAll(name)

	testChtimes(t, name)
}

func testChtimes(t *testing.T, name string) {
	st, err := Stat(name)
	if err != nil {
		t.Fatalf("Stat %s: %s", name, err)
	}
	preStat := st

	// 将访问和修改时间向后回退一秒
	at := Atime(preStat)
	mt := preStat.ModTime()
	err = Chtimes(name, at.Add(-time.Second), mt.Add(-time.Second))
	if err != nil {
		t.Fatalf("Chtimes %s: %s", name, err)
	}

	st, err = Stat(name)
	if err != nil {
		t.Fatalf("second Stat %s: %s", name, err)
	}
	postStat := st

	pat := Atime(postStat)
	pmt := postStat.ModTime()
	if !pat.Before(at) {
		switch runtime.GOOS {
		case "plan9":
			// Mtime is the time of the last change of
			// content.  Similarly, atime is set whenever
			// the contents are accessed; also, it is set
			// whenever mtime is set.
		case "netbsd":
			mounts, _ := ReadFile("/proc/mounts")
			if strings.Contains(string(mounts), "noatime") {
				t.Logf("AccessTime didn't go backwards, but see a filesystem mounted noatime; ignoring. Issue 19293.")
			} else {
				t.Logf("AccessTime didn't go backwards; was=%v, after=%v (Ignoring on NetBSD, assuming noatime, Issue 19293)", at, pat)
			}
		default:
			t.Errorf("AccessTime didn't go backwards; was=%v, after=%v", at, pat)
		}
	}

	if !pmt.Before(mt) {
		t.Errorf("ModTime didn't go backwards; was=%v, after=%v", mt, pmt)
	}
}

func TestChtimesToUnixZero(t *testing.T) {
	file := newFile("chtimes-to-unix-zero", t)
	fn := file.Name()
	defer Remove(fn)
	if _, err := file.Write([]byte("hi")); err != nil {
		t.Fatal(err)
	}
	if err := file.Close(); err != nil {
		t.Fatal(err)
	}

	unixZero := time.Unix(0, 0)
	if err := Chtimes(fn, unixZero, unixZero); err != nil {
		t.Fatalf("Chtimes failed: %v", err)
	}

	st, err := Stat(fn)
	if err != nil {
		t.Fatal(err)
	}

	if mt := st.ModTime(); mt != unixZero {
		t.Errorf("mtime is %v, want %v", mt, unixZero)
	}
}

func TestFileChdir(t *testing.T) {
	wd, err := Getwd()
	if err != nil {
		t.Fatalf("Getwd: %s", err)
	}
	defer Chdir(wd)

	fd, err := Open(".")
	if err != nil {
		t.Fatalf("Open .: %s", err)
	}
	defer fd.Close()

	if err := Chdir("/"); err != nil {
		t.Fatalf("Chdir /: %s", err)
	}

	if err := fd.Chdir(); err != nil {
		t.Fatalf("fd.Chdir: %s", err)
	}

	wdNew, err := Getwd()
	if err != nil {
		t.Fatalf("Getwd: %s", err)
	}

	wdInfo, err := fd.Stat()
	if err != nil {
		t.Fatal(err)
	}
	newInfo, err := Stat(wdNew)
	if err != nil {
		t.Fatal(err)
	}
	if !SameFile(wdInfo, newInfo) {
		t.Fatalf("fd.Chdir failed: got %s, want %s", wdNew, wd)
	}
}

func TestChdirAndGetwd(t *testing.T) {
	fd, err := Open(".")
	if err != nil {
		t.Fatalf("Open .: %s", err)
	}
	// 这些路径都是精心选择的，不会在Mac上成为符号链接（不像/var、/etc），除了/tmp，我们稍后会处理。
	dirs := []string{"/", "/usr/bin", "/tmp"}
	// 在 Plan 9 或 Android 系统中，/usr/bin 通常不存在。
	switch runtime.GOOS {
	case "android":
		dirs = []string{"/system/bin"}
	case "plan9":
		dirs = []string{"/", "/usr"}
	case "ios", "windows", "wasip1":
		dirs = nil
		for _, dir := range []string{t.TempDir(), t.TempDir()} {
			// 展开符号链接，以便进行路径相等性测试。
			dir, err = filepath.EvalSymlinks(dir)
			if err != nil {
				t.Fatalf("EvalSymlinks: %v", err)
			}
			dirs = append(dirs, dir)
		}
	}
	oldwd := Getenv("PWD")
	for mode := 0; mode < 2; mode++ {
		for _, d := range dirs {
			if mode == 0 {
				err = Chdir(d)
			} else {
				fd1, err1 := Open(d)
				if err1 != nil {
					t.Errorf("Open %s: %s", d, err1)
					continue
				}
				err = fd1.Chdir()
				fd1.Close()
			}
			if d == "/tmp" {
				Setenv("PWD", "/tmp")
			}
			pwd, err1 := Getwd()
			Setenv("PWD", oldwd)
			err2 := fd.Chdir()
			if err2 != nil {
				// 我们已经改变了当前目录，无法返回。
				// 不要让测试继续进行；它们会乱写
				// 到其他目录中。
				fmt.Fprintf(Stderr, "fchdir back to dot failed: %s\n", err2)
				Exit(1)
			}
			if err != nil {
				fd.Close()
				t.Fatalf("Chdir %s: %s", d, err)
			}
			if err1 != nil {
				fd.Close()
				t.Fatalf("Getwd in %s: %s", d, err1)
			}
			if !equal(pwd, d) {
				fd.Close()
				t.Fatalf("Getwd returned %q want %q", pwd, d)
			}
		}
	}
	fd.Close()
}

// 测试Chdir+Getwd在整个程序中的功能。
func TestProgWideChdir(t *testing.T) {
	const N = 10
	var wg sync.WaitGroup
	hold := make(chan struct{})
	done := make(chan struct{})

	d := t.TempDir()
	oldwd, err := Getwd()
	if err != nil {
		t.Fatalf("Getwd: %v", err)
	}
	defer func() {
		if err := Chdir(oldwd); err != nil {
			// 如果我们无法返回到原始工作目录，那么继续进行测试是不安全的。
			panic(err)
		}
	}()

	// 注意，延迟调用的Wait必须在延迟关闭(done)之后调用，
	// 以确保即使主goroutine调用了Fatalf，N个goroutine也已经被释放。它必须在
	// 恢复到原始目录之前以及TempDir隐含的延迟删除之前调用，
	// 以免在N个goroutine仍在运行时产生干扰。
	defer wg.Wait()
	defer close(done)

	for i := 0; i < N; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			// 将一半的goroutine锁定在自己的操作系统线程中，以锻炼更多的调度可能性。
			if i%2 == 1 {
				// 在Plan 9上，调用LockOSThread后，goroutine将在不同的进程中运行，这些进程不共享工作目录。这曾经是一个问题，因为Go期望工作目录是全局的。
				// 参见问题9428。
				runtime.LockOSThread()
			}
			select {
			case <-done:
				return
			case <-hold:
			}
			// Getwd might be wrong
			f0, err := Stat(".")
			if err != nil {
				t.Error(err)
				return
			}
			pwd, err := Getwd()
			if err != nil {
				t.Errorf("Getwd: %v", err)
				return
			}
			if pwd != d {
				t.Errorf("Getwd() = %q, want %q", pwd, d)
				return
			}
			f1, err := Stat(pwd)
			if err != nil {
				t.Error(err)
				return
			}
			if !SameFile(f0, f1) {
				t.Errorf(`Samefile(Stat("."), Getwd()) reports false (%s != %s)`, f0.Name(), f1.Name())
				return
			}
		}(i)
	}
	if err = Chdir(d); err != nil {
		t.Fatalf("Chdir: %v", err)
	}
	// 在OS X系统中，TMPDIR设置为一个符号链接。
	// 因此我们在测试之前再次解析我们的工作目录。
	d, err = Getwd()
	if err != nil {
		t.Fatalf("Getwd: %v", err)
	}
	close(hold)
	wg.Wait()
}

func TestSeek(t *testing.T) {
	t.Parallel()

	f := newFile("TestSeek", t)
	defer Remove(f.Name())
	defer f.Close()

	const data = "hello, world\n"
	io.WriteString(f, data)

	type test struct {
		in     int64
		whence int
		out    int64
	}
	var tests = []test{
		{0, io.SeekCurrent, int64(len(data))},
		{0, io.SeekStart, 0},
		{5, io.SeekStart, 5},
		{0, io.SeekEnd, int64(len(data))},
		{0, io.SeekStart, 0},
		{-1, io.SeekEnd, int64(len(data)) - 1},
		{1 << 33, io.SeekStart, 1 << 33},
		{1 << 33, io.SeekEnd, 1<<33 + int64(len(data))},

		// 问题21681，Windows 4G-1等：
		{1<<32 - 1, io.SeekStart, 1<<32 - 1},
		{0, io.SeekCurrent, 1<<32 - 1},
		{2<<32 - 1, io.SeekStart, 2<<32 - 1},
		{0, io.SeekCurrent, 2<<32 - 1},
	}
	for i, tt := range tests {
		off, err := f.Seek(tt.in, tt.whence)
		if off != tt.out || err != nil {
			if e, ok := err.(*PathError); ok && e.Err == syscall.EINVAL && tt.out > 1<<32 && runtime.GOOS == "linux" {
				mounts, _ := ReadFile("/proc/mounts")
				if strings.Contains(string(mounts), "reiserfs") {
					// Reiserfs拒绝大的 seek 操作。
					t.Skipf("skipping test known to fail on reiserfs; https://golang.org/issue/91")
				}
			}
			t.Errorf("#%d: Seek(%v, %v) = %v, %v want %v, nil", i, tt.in, tt.whence, off, err, tt.out)
		}
	}
}

func TestSeekError(t *testing.T) {
	switch runtime.GOOS {
	case "js", "plan9", "wasip1":
		t.Skipf("skipping test on %v", runtime.GOOS)
	}
	t.Parallel()

	r, w, err := Pipe()
	if err != nil {
		t.Fatal(err)
	}
	_, err = r.Seek(0, 0)
	if err == nil {
		t.Fatal("Seek on pipe should fail")
	}
	if perr, ok := err.(*PathError); !ok || perr.Err != syscall.ESPIPE {
		t.Errorf("Seek returned error %v, want &PathError{Err: syscall.ESPIPE}", err)
	}
	_, err = w.Seek(0, 0)
	if err == nil {
		t.Fatal("Seek on pipe should fail")
	}
	if perr, ok := err.(*PathError); !ok || perr.Err != syscall.ESPIPE {
		t.Errorf("Seek returned error %v, want &PathError{Err: syscall.ESPIPE}", err)
	}
}

type openErrorTest struct {
	path  string
	mode  int
	error error
}

var openErrorTests = []openErrorTest{
	{
		sfdir + "/no-such-file",
		O_RDONLY,
		syscall.ENOENT,
	},
	{
		sfdir,
		O_WRONLY,
		syscall.EISDIR,
	},
	{
		sfdir + "/" + sfname + "/no-such-file",
		O_WRONLY,
		syscall.ENOTDIR,
	},
}

func TestOpenError(t *testing.T) {
	t.Parallel()

	for _, tt := range openErrorTests {
		f, err := OpenFile(tt.path, tt.mode, 0)
		if err == nil {
			t.Errorf("Open(%q, %d) succeeded", tt.path, tt.mode)
			f.Close()
			continue
		}
		perr, ok := err.(*PathError)
		if !ok {
			t.Errorf("Open(%q, %d) returns error of %T type; want *PathError", tt.path, tt.mode, err)
		}
		if perr.Err != tt.error {
			if runtime.GOOS == "plan9" {
				syscallErrStr := perr.Err.Error()
				expectedErrStr := strings.Replace(tt.error.Error(), "file ", "", 1)
				if !strings.HasSuffix(syscallErrStr, expectedErrStr) {
					// 一些 Plan 9 文件服务器在打开目录进行写操作时，错误地返回 EACCES 而不是 EISDIR。
					if tt.error == syscall.EISDIR && strings.HasSuffix(syscallErrStr, syscall.EACCES.Error()) {
						continue
					}
					t.Errorf("Open(%q, %d) = _, %q; want suffix %q", tt.path, tt.mode, syscallErrStr, expectedErrStr)
				}
				continue
			}
			if runtime.GOOS == "dragonfly" {
				// DragonFly 系统在尝试以写入方式打开目录时，错误地返回 EACCES 而非 EISDIR。
				if tt.error == syscall.EISDIR && perr.Err == syscall.EACCES {
					continue
				}
			}
			t.Errorf("Open(%q, %d) = _, %q; want %q", tt.path, tt.mode, perr.Err.Error(), tt.error.Error())
		}
	}
}

func TestOpenNoName(t *testing.T) {
	f, err := Open("")
	if err == nil {
		f.Close()
		t.Fatal(`Open("") succeeded`)
	}
}

//func runBinHostname(t *testing.T) string {
//	// 运行/bin/hostname命令并收集输出。
//	r, w, err := Pipe()
//	if err != nil {
//		t.Fatal(err)
//	}
//	defer r.Close()
//
//	path, err := exec.LookPath("hostname")
//	if err != nil {
//		if errors.Is(err, exec.ErrNotFound) {
//			t.Skip("skipping test; test requires hostname but it does not exist")
//		}
//		t.Fatal(err)
//	}
//
//	argv := []string{"hostname"}
//	if runtime.GOOS == "aix" {
//		argv = []string{"hostname", "-s"}
//	}
//
//	p, err := StartProcess(path, argv, &ProcAttr{Files: []*File{nil, w, Stderr}})
//	if err != nil {
//		t.Fatal(err)
//	}
//	w.Close()
//
//	var b strings.Builder
//	io.Copy(&b, r)
//	_, err = p.Wait()
//	if err != nil {
//		t.Fatalf("run hostname Wait: %v", err)
//	}
//	err = p.Kill()
//	if err == nil {
//		t.Errorf("expected an error from Kill running 'hostname'")
//	}
//	output := b.String()
//	if n := len(output); n > 0 && output[n-1] == '\n' {
//		output = output[0 : n-1]
//	}
//	if output == "" {
//		t.Fatalf("/bin/hostname produced no output")
//	}
//
//	return output
//}

func testWindowsHostname(t *testing.T, hostname string) {
	cmd := testenv.Command(t, "hostname")
	out, err := cmd.Output()
	if err != nil {
		t.Fatalf("Failed to execute hostname command: %v %s", err, out)
	}
	want := strings.Trim(string(out), "\r\n")
	if hostname != want {
		t.Fatalf("Hostname() = %q != system hostname of %q", hostname, want)
	}
}

//
//func TestHostname(t *testing.T) {
//	t.Parallel()
//
//	hostname, err := Hostname()
//	if err != nil {
//		t.Fatal(err)
//	}
//	if hostname == "" {
//		t.Fatal("Hostname returned empty string and no error")
//	}
//	if strings.Contains(hostname, "\x00") {
//		t.Fatalf("unexpected zero byte in hostname: %q", hostname)
//	}
//
//	// 在 Windows 上，没有其他方法可以获取主机名，但可以通过 WinAPI 实现。
//	// 在 Plan 9 中，可以从 #c/sysname 获取，就像 Hostname() 函数所做的那样。
//	switch runtime.GOOS {
//	case "android", "plan9":
//		// 没有 /bin/hostname 可以进行校验。
//		return
//	case "windows":
//		testWindowsHostname(t, hostname)
//		return
//	}
//
//	testenv.MustHaveExec(t)
//
//	// 检查内部 Hostname() 函数的输出与 /bin/hostname 命令的输出是否一致。
//	// 允许内部 Hostname 返回一个完整的域名（Fully Qualified Domain Name，FQDN），
//	// 而 /bin/hostname 只返回第一个名称组件。
//	want := runBinHostname(t)
//	if hostname != want {
//		host, _, ok := strings.Cut(hostname, ".")
//		if !ok || host != want {
//			t.Errorf("Hostname() = %q, want %q", hostname, want)
//		}
//	}
//}

func TestReadAt(t *testing.T) {
	t.Parallel()

	f := newFile("TestReadAt", t)
	defer Remove(f.Name())
	defer f.Close()

	const data = "hello, world\n"
	io.WriteString(f, data)

	b := make([]byte, 5)
	n, err := f.ReadAt(b, 7)
	if err != nil || n != len(b) {
		t.Fatalf("ReadAt 7: %d, %v", n, err)
	}
	if string(b) != "world" {
		t.Fatalf("ReadAt 7: have %q want %q", string(b), "world")
	}
}

// 验证ReadAt方法不会影响 Seek 偏移量。
// 在Plan 9的内核中，pread系统调用的实现曾经存在一个错误，
// 在对文件调用pread后，会错误地更新通道偏移量。
func TestReadAtOffset(t *testing.T) {
	t.Parallel()

	f := newFile("TestReadAtOffset", t)
	defer Remove(f.Name())
	defer f.Close()

	const data = "hello, world\n"
	io.WriteString(f, data)

	f.Seek(0, 0)
	b := make([]byte, 5)

	n, err := f.ReadAt(b, 7)
	if err != nil || n != len(b) {
		t.Fatalf("ReadAt 7: %d, %v", n, err)
	}
	if string(b) != "world" {
		t.Fatalf("ReadAt 7: have %q want %q", string(b), "world")
	}

	n, err = f.Read(b)
	if err != nil || n != len(b) {
		t.Fatalf("Read: %d, %v", n, err)
	}
	if string(b) != "hello" {
		t.Fatalf("Read: have %q want %q", string(b), "hello")
	}
}

// 验证ReadAt是否不允许负偏移量。
func TestReadAtNegativeOffset(t *testing.T) {
	t.Parallel()

	f := newFile("TestReadAtNegativeOffset", t)
	defer Remove(f.Name())
	defer f.Close()

	const data = "hello, world\n"
	io.WriteString(f, data)

	f.Seek(0, 0)
	b := make([]byte, 5)

	n, err := f.ReadAt(b, -10)

	const wantsub = "negative offset"
	if !strings.Contains(fmt.Sprint(err), wantsub) || n != 0 {
		t.Errorf("ReadAt(-10) = %v, %v; want 0, ...%q...", n, err, wantsub)
	}
}

func TestWriteAt(t *testing.T) {
	t.Parallel()

	f := newFile("TestWriteAt", t)
	defer Remove(f.Name())
	defer f.Close()

	const data = "hello, world\n"
	io.WriteString(f, data)

	n, err := f.WriteAt([]byte("WORLD"), 7)
	if err != nil || n != 5 {
		t.Fatalf("WriteAt 7: %d, %v", n, err)
	}

	b, err := ReadFile(f.Name())
	if err != nil {
		t.Fatalf("ReadFile %s: %v", f.Name(), err)
	}
	if string(b) != "hello, WORLD\n" {
		t.Fatalf("after write: have %q want %q", string(b), "hello, WORLD\n")
	}
}

// 验证WriteAt是否不允许负偏移量。
func TestWriteAtNegativeOffset(t *testing.T) {
	t.Parallel()

	f := newFile("TestWriteAtNegativeOffset", t)
	defer Remove(f.Name())
	defer f.Close()

	n, err := f.WriteAt([]byte("WORLD"), -10)

	const wantsub = "negative offset"
	if !strings.Contains(fmt.Sprint(err), wantsub) || n != 0 {
		t.Errorf("WriteAt(-10) = %v, %v; want 0, ...%q...", n, err, wantsub)
	}
}

// 验证在追加模式下WriteAt方法不能正常工作
func TestWriteAtInAppendMode(t *testing.T) {
	defer chtmpdir(t)()
	f, err := OpenFile("write_at_in_append_mode.txt", O_APPEND|O_CREATE, 0666)
	if err != nil {
		t.Fatalf("OpenFile: %v", err)
	}
	defer f.Close()

	_, err = f.WriteAt([]byte(""), 1)
	if err != ErrWriteAtInAppendMode {
		t.Fatalf("f.WriteAt returned %v, expected %v", err, ErrWriteAtInAppendMode)
	}
}

func writeFile(t *testing.T, fname string, flag int, text string) string {
	f, err := OpenFile(fname, flag, 0666)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	n, err := io.WriteString(f, text)
	if err != nil {
		t.Fatalf("WriteString: %d, %v", n, err)
	}
	f.Close()
	data, err := ReadFile(fname)
	if err != nil {
		t.Fatalf("ReadFile: %v", err)
	}
	return string(data)
}

func TestAppend(t *testing.T) {
	defer chtmpdir(t)()
	const f = "append.txt"
	s := writeFile(t, f, O_CREATE|O_TRUNC|O_RDWR, "new")
	if s != "new" {
		t.Fatalf("writeFile: have %q want %q", s, "new")
	}
	s = writeFile(t, f, O_APPEND|O_RDWR, "|append")
	if s != "new|append" {
		t.Fatalf("writeFile: have %q want %q", s, "new|append")
	}
	s = writeFile(t, f, O_CREATE|O_APPEND|O_RDWR, "|append")
	if s != "new|append|append" {
		t.Fatalf("writeFile: have %q want %q", s, "new|append|append")
	}
	err := Remove(f)
	if err != nil {
		t.Fatalf("Remove: %v", err)
	}
	s = writeFile(t, f, O_CREATE|O_APPEND|O_RDWR, "new&append")
	if s != "new&append" {
		t.Fatalf("writeFile: after append have %q want %q", s, "new&append")
	}
	s = writeFile(t, f, O_CREATE|O_RDWR, "old")
	if s != "old&append" {
		t.Fatalf("writeFile: after create have %q want %q", s, "old&append")
	}
	s = writeFile(t, f, O_CREATE|O_TRUNC|O_RDWR, "new")
	if s != "new" {
		t.Fatalf("writeFile: after truncate have %q want %q", s, "new")
	}
}

func TestStatDirWithTrailingSlash(t *testing.T) {
	t.Parallel()

	// 创建一个新的临时目录，并安排在完成后进行清理。
	path := t.TempDir()

	// 路径的状态应该成功。
	if _, err := Stat(path); err != nil {
		t.Fatalf("stat %s failed: %s", path, err)
	}

	// 路径("/")的Stat也应该成功。
	path += "/"
	if _, err := Stat(path); err != nil {
		t.Fatalf("stat %s failed: %s", path, err)
	}
}

func TestNilProcessStateString(t *testing.T) {
	var ps *ProcessState
	s := ps.String()
	if s != "<nil>" {
		t.Errorf("(*ProcessState)(nil).String() = %q, want %q", s, "<nil>")
	}
}

func TestSameFile(t *testing.T) {
	defer chtmpdir(t)()
	fa, err := Create("a")
	if err != nil {
		t.Fatalf("Create(a): %v", err)
	}
	fa.Close()
	fb, err := Create("b")
	if err != nil {
		t.Fatalf("Create(b): %v", err)
	}
	fb.Close()

	ia1, err := Stat("a")
	if err != nil {
		t.Fatalf("Stat(a): %v", err)
	}
	ia2, err := Stat("a")
	if err != nil {
		t.Fatalf("Stat(a): %v", err)
	}
	if !SameFile(ia1, ia2) {
		t.Errorf("files should be same")
	}

	ib, err := Stat("b")
	if err != nil {
		t.Fatalf("Stat(b): %v", err)
	}
	if SameFile(ia1, ib) {
		t.Errorf("files should be different")
	}
}

func testDevNullFileInfo(t *testing.T, statname, devNullName string, fi FileInfo) {
	pre := fmt.Sprintf("%s(%q): ", statname, devNullName)
	if fi.Size() != 0 {
		t.Errorf(pre+"wrong file size have %d want 0", fi.Size())
	}
	if fi.Mode()&ModeDevice == 0 {
		t.Errorf(pre+"wrong file mode %q: ModeDevice is not set", fi.Mode())
	}
	if fi.Mode()&ModeCharDevice == 0 {
		t.Errorf(pre+"wrong file mode %q: ModeCharDevice is not set", fi.Mode())
	}
	if fi.Mode().IsRegular() {
		t.Errorf(pre+"wrong file mode %q: IsRegular returns true", fi.Mode())
	}
}

func testDevNullFile(t *testing.T, devNullName string) {
	f, err := Open(devNullName)
	if err != nil {
		t.Fatalf("Open(%s): %v", devNullName, err)
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		t.Fatalf("Stat(%s): %v", devNullName, err)
	}
	testDevNullFileInfo(t, "f.Stat", devNullName, fi)

	fi, err = Stat(devNullName)
	if err != nil {
		t.Fatalf("Stat(%s): %v", devNullName, err)
	}
	testDevNullFileInfo(t, "Stat", devNullName, fi)
}

func TestDevNullFile(t *testing.T) {
	t.Parallel()

	testDevNullFile(t, DevNull)
	if runtime.GOOS == "windows" {
		testDevNullFile(t, "./nul")
		testDevNullFile(t, "//./nul")
	}
}

var testLargeWrite = flag.Bool("large_write", false, "run TestLargeWriteToConsole test that floods console with output")

func TestLargeWriteToConsole(t *testing.T) {
	if !*testLargeWrite {
		t.Skip("skipping console-flooding test; enable with -large_write")
	}
	b := make([]byte, 32000)
	for i := range b {
		b[i] = '.'
	}
	b[len(b)-1] = '\n'
	n, err := Stdout.Write(b)
	if err != nil {
		t.Fatalf("Write to os.Stdout failed: %v", err)
	}
	if n != len(b) {
		t.Errorf("Write to os.Stdout should return %d; got %d", len(b), n)
	}
	n, err = Stderr.Write(b)
	if err != nil {
		t.Fatalf("Write to os.Stderr failed: %v", err)
	}
	if n != len(b) {
		t.Errorf("Write to os.Stderr should return %d; got %d", len(b), n)
	}
}

func TestStatDirModeExec(t *testing.T) {
	if runtime.GOOS == "wasip1" {
		t.Skip("Chmod is not supported on " + runtime.GOOS)
	}
	t.Parallel()

	const mode = 0111

	path := t.TempDir()
	if err := Chmod(path, 0777); err != nil {
		t.Fatalf("Chmod %q 0777: %v", path, err)
	}

	dir, err := Stat(path)
	if err != nil {
		t.Fatalf("Stat %q (looking for mode %#o): %s", path, mode, err)
	}
	if dir.Mode()&mode != mode {
		t.Errorf("Stat %q: mode %#o want %#o", path, dir.Mode()&mode, mode)
	}
}

func TestStatStdin(t *testing.T) {
	switch runtime.GOOS {
	case "android", "plan9":
		t.Skipf("%s doesn't have /bin/sh", runtime.GOOS)
	}

	if Getenv("GO_WANT_HELPER_PROCESS") == "1" {
		st, err := Stdin.Stat()
		if err != nil {
			t.Fatalf("Stat failed: %v", err)
		}
		fmt.Println(st.Mode() & ModeNamedPipe)
		Exit(0)
	}

	exe, err := Executable()
	if err != nil {
		t.Skipf("can't find executable: %v", err)
	}

	testenv.MustHaveExec(t)
	t.Parallel()

	fi, err := Stdin.Stat()
	if err != nil {
		t.Fatal(err)
	}
	switch mode := fi.Mode(); {
	case mode&ModeCharDevice != 0 && mode&ModeDevice != 0:
	case mode&ModeNamedPipe != 0:
	default:
		t.Fatalf("unexpected Stdin mode (%v), want ModeCharDevice or ModeNamedPipe", mode)
	}

	cmd := testenv.Command(t, exe, "-test.run=^TestStatStdin$")
	cmd = testenv.CleanCmdEnv(cmd)
	cmd.Env = append(cmd.Env, "GO_WANT_HELPER_PROCESS=1")
	// 这将会使标准输入变为管道。
	cmd.Stdin = strings.NewReader("output")

	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Failed to spawn child process: %v %q", err, string(output))
	}

	// result 将类似于 "prw-rw-rw"
	if len(output) < 1 || output[0] != 'p' {
		t.Fatalf("Child process reports stdin is not pipe '%v'", string(output))
	}
}

func TestStatRelativeSymlink(t *testing.T) {
	testenv.MustHaveSymlink(t)
	t.Parallel()

	tmpdir := t.TempDir()
	target := filepath.Join(tmpdir, "target")
	f, err := Create(target)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	st, err := f.Stat()
	if err != nil {
		t.Fatal(err)
	}

	link := filepath.Join(tmpdir, "link")
	err = Symlink(filepath.Base(target), link)
	if err != nil {
		t.Fatal(err)
	}

	st1, err := Stat(link)
	if err != nil {
		t.Fatal(err)
	}

	if !SameFile(st, st1) {
		t.Error("Stat doesn't follow relative symlink")
	}

	if runtime.GOOS == "windows" {
		Remove(link)
		err = Symlink(target[len(filepath.VolumeName(target)):], link)
		if err != nil {
			t.Fatal(err)
		}

		st1, err := Stat(link)
		if err != nil {
			t.Fatal(err)
		}

		if !SameFile(st, st1) {
			t.Error("Stat doesn't follow relative symlink")
		}
	}
}

func TestReadAtEOF(t *testing.T) {
	t.Parallel()

	f := newFile("TestReadAtEOF", t)
	defer Remove(f.Name())
	defer f.Close()

	_, err := f.ReadAt(make([]byte, 10), 0)
	switch err {
	case io.EOF:
		// all good
	case nil:
		t.Fatalf("ReadAt succeeded")
	default:
		t.Fatalf("ReadAt failed: %s", err)
	}
}

func TestLongPath(t *testing.T) {
	t.Parallel()

	tmpdir := newDir("TestLongPath", t)
	defer func(d string) {
		if err := RemoveAll(d); err != nil {
			t.Fatalf("RemoveAll failed: %v", err)
		}
	}(tmpdir)

	// 测试正常情况下（247字节及更少）和调整后情况下（248字节及以上）的边界。
	sizes := []int{247, 248, 249, 400}
	for len(tmpdir) < 400 {
		tmpdir += "/dir3456789"
	}
	for _, sz := range sizes {
		t.Run(fmt.Sprintf("length=%d", sz), func(t *testing.T) {
			sizedTempDir := tmpdir[:sz-1] + "x" // 确保它不以斜杠结束。

			// 各种大小的运行是为了触发此调用的边界条件。
			if err := MkdirAll(sizedTempDir, 0755); err != nil {
				t.Fatalf("MkdirAll failed: %v", err)
			}
			data := []byte("hello world\n")
			if err := WriteFile(sizedTempDir+"/foo.txt", data, 0644); err != nil {
				t.Fatalf("os.WriteFile() failed: %v", err)
			}
			if err := Rename(sizedTempDir+"/foo.txt", sizedTempDir+"/bar.txt"); err != nil {
				t.Fatalf("Rename failed: %v", err)
			}
			mtime := time.Now().Truncate(time.Minute)
			if err := Chtimes(sizedTempDir+"/bar.txt", mtime, mtime); err != nil {
				t.Fatalf("Chtimes failed: %v", err)
			}
			names := []string{"bar.txt"}
			if testenv.HasSymlink() {
				if err := Symlink(sizedTempDir+"/bar.txt", sizedTempDir+"/symlink.txt"); err != nil {
					t.Fatalf("Symlink failed: %v", err)
				}
				names = append(names, "symlink.txt")
			}
			if testenv.HasLink() {
				if err := Link(sizedTempDir+"/bar.txt", sizedTempDir+"/link.txt"); err != nil {
					t.Fatalf("Link failed: %v", err)
				}
				names = append(names, "link.txt")
			}
			for _, wantSize := range []int64{int64(len(data)), 0} {
				for _, name := range names {
					path := sizedTempDir + "/" + name
					dir, err := Stat(path)
					if err != nil {
						t.Fatalf("Stat(%q) failed: %v", path, err)
					}
					filesize := size(path, t)
					if dir.Size() != filesize || filesize != wantSize {
						t.Errorf("Size(%q) is %d, len(ReadFile()) is %d, want %d", path, dir.Size(), filesize, wantSize)
					}
					if runtime.GOOS != "wasip1" { // Chmod is not supported on wasip1
						err = Chmod(path, dir.Mode())
						if err != nil {
							t.Fatalf("Chmod(%q) failed: %v", path, err)
						}
					}
				}
				if err := Truncate(sizedTempDir+"/bar.txt", 0); err != nil {
					t.Fatalf("Truncate failed: %v", err)
				}
			}
		})
	}
}

func testKillProcess(t *testing.T, processKiller func(p *Process)) {
	testenv.MustHaveExec(t)
	t.Parallel()

	// 重新执行测试二进制文件，以启动一个进程，该进程将挂起，直到标准输入被关闭。
	cmd := testenv.Command(t, os2.Args[0])
	cmd.Env = append(cmd.Environ(), "GO_OS_TEST_DRAIN_STDIN=1")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		t.Fatal(err)
	}
	stdin, err := cmd.StdinPipe()
	if err != nil {
		t.Fatal(err)
	}
	err = cmd.Start()
	if err != nil {
		t.Fatalf("Failed to start test process: %v", err)
	}

	defer func() {
		if err := cmd.Wait(); err == nil {
			t.Errorf("Test process succeeded, but expected to fail")
		}
		stdin.Close() // 保持标准输入流畅通，直到该进程完全退出。
	}()

	// 等待进程启动。
	// （当它达到TestMain时，它将关闭其stdout。）
	io.Copy(io.Discard, stdout)
	processKiller(&Process{
		F: cmd.Process,
	})

}

func TestKillStartProcess(t *testing.T) {
	testKillProcess(t, func(p *Process) {
		err := p.Kill()
		if err != nil {
			t.Fatalf("Failed to kill test process: %v", err)
		}
	})
}

func TestGetppid(t *testing.T) {
	if runtime.GOOS == "plan9" {
		// 待办事项：golang.org/issue/8206
		t.Skipf("skipping test on plan9; see issue 8206")
	}

	if Getenv("GO_WANT_HELPER_PROCESS") == "1" {
		fmt.Print(Getppid())
		Exit(0)
	}

	testenv.MustHaveExec(t)
	t.Parallel()

	cmd := testenv.Command(t, os2.Args[0], "-test.run=^TestGetppid$")
	cmd.Env = append(Environ(), "GO_WANT_HELPER_PROCESS=1")

	// 验证从forked进程调用的Getppid()是否报告我们的进程ID
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Failed to spawn child process: %v %q", err, string(output))
	}

	childPpid := string(output)
	ourPid := fmt.Sprintf("%d", Getpid())
	if childPpid != ourPid {
		t.Fatalf("Child process reports parent process id '%v', expected '%v'", childPpid, ourPid)
	}
}

func TestKillFindProcess(t *testing.T) {
	testKillProcess(t, func(p *Process) {
		p2, err := FindProcess(p.F.Pid)
		if err != nil {
			t.Fatalf("Failed to find test process: %v", err)
		}
		err = p2.Kill()
		if err != nil {
			t.Fatalf("Failed to kill test process: %v", err)
		}
	})
}

var nilFileMethodTests = []struct {
	name string
	f    func(*File) error
}{
	{"Chdir", func(f *File) error { return f.Chdir() }},
	{"Close", func(f *File) error { return f.Close() }},
	{"Chmod", func(f *File) error { return f.Chmod(0) }},
	{"Chown", func(f *File) error { return f.Chown(0, 0) }},
	{"Read", func(f *File) error { _, err := f.Read(make([]byte, 0)); return err }},
	{"ReadAt", func(f *File) error { _, err := f.ReadAt(make([]byte, 0), 0); return err }},
	{"Readdir", func(f *File) error { _, err := f.Readdir(1); return err }},
	{"Readdirnames", func(f *File) error { _, err := f.Readdirnames(1); return err }},
	{"Seek", func(f *File) error { _, err := f.Seek(0, io.SeekStart); return err }},
	{"Stat", func(f *File) error { _, err := f.Stat(); return err }},
	{"Sync", func(f *File) error { return f.Sync() }},
	{"Truncate", func(f *File) error { return f.Truncate(0) }},
	{"Write", func(f *File) error { _, err := f.Write(make([]byte, 0)); return err }},
	{"WriteAt", func(f *File) error { _, err := f.WriteAt(make([]byte, 0), 0); return err }},
	{"WriteString", func(f *File) error { _, err := f.WriteString(""); return err }},
}

// 测试如果接收者为nil，所有File方法都会返回ErrInvalid。
func TestNilFileMethods(t *testing.T) {
	t.Parallel()

	for _, tt := range nilFileMethodTests {
		var file *File
		got := tt.f(file)
		if got != ErrInvalid {
			t.Errorf("%v should fail when f is nil; got %v", tt.name, got)
		}
	}
}

func mkdirTree(t *testing.T, root string, level, max int) {
	if level >= max {
		return
	}
	level++
	for i := 'a'; i < 'c'; i++ {
		dir := filepath.Join(root, string(i))
		if err := Mkdir(dir, 0700); err != nil {
			t.Fatal(err)
		}
		mkdirTree(t, dir, level, max)
	}
}

// 测试同时移除所有是否报告错误。
// 只要它被移除了，我们就应该感到满意。
func TestRemoveAllRace(t *testing.T) {
	if runtime.GOOS == "windows" {
		// Windows 对诸如在其他程序打开目录时删除目录等操作有着非常严格的规定。这种竞态条件在 Unix 系统中可以很好地处理，但在 Windows 上并不理想。
		t.Skip("skipping on windows")
	}
	if runtime.GOOS == "dragonfly" {
		testenv.SkipFlaky(t, 52301)
	}

	n := runtime.GOMAXPROCS(16)
	defer runtime.GOMAXPROCS(n)
	root, err := MkdirTemp("", "issue")
	if err != nil {
		t.Fatal(err)
	}
	mkdirTree(t, root, 1, 6)
	hold := make(chan struct{})
	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			<-hold
			err := RemoveAll(root)
			if err != nil {
				t.Errorf("unexpected error: %T, %q", err, err)
			}
		}()
	}
	close(hold) // 让工作者竞争移除根节点
	wg.Wait()
}

// 测试从管道读取不会占用线程。
func TestPipeThreads(t *testing.T) {
	switch runtime.GOOS {
	case "illumos", "solaris":
		t.Skip("skipping on Solaris and illumos; issue 19111")
	case "windows":
		t.Skip("skipping on Windows; issue 19098")
	case "plan9":
		t.Skip("skipping on Plan 9; does not support runtime poller")
	case "js":
		t.Skip("skipping on js; no support for os.Pipe")
	case "wasip1":
		t.Skip("skipping on wasip1; no support for os.Pipe")
	}

	threads := 100

	// OpenBSD 对最大文件数量的默认值设置较低。
	if runtime.GOOS == "openbsd" {
		threads = 50
	}

	r := make([]*File, threads)
	w := make([]*File, threads)
	for i := 0; i < threads; i++ {
		rp, wp, err := Pipe()
		if err != nil {
			for j := 0; j < i; j++ {
				r[j].Close()
				w[j].Close()
			}
			t.Fatal(err)
		}
		r[i] = rp
		w[i] = wp
	}

	defer debug.SetMaxThreads(debug.SetMaxThreads(threads / 2))

	creading := make(chan bool, threads)
	cdone := make(chan bool, threads)
	for i := 0; i < threads; i++ {
		go func(i int) {
			var b [1]byte
			creading <- true
			if _, err := r[i].Read(b[:]); err != nil {
				t.Error(err)
			}
			if err := r[i].Close(); err != nil {
				t.Error(err)
			}
			cdone <- true
		}(i)
	}

	for i := 0; i < threads; i++ {
		<-creading
	}

	// 如果我们仍然存活，那就意味着这100个goroutine并未要求使用100个线程。

	for i := 0; i < threads; i++ {
		if _, err := w[i].Write([]byte{0}); err != nil {
			t.Error(err)
		}
		if err := w[i].Close(); err != nil {
			t.Error(err)
		}
		<-cdone
	}
}

func testDoubleCloseError(path string) func(*testing.T) {
	return func(t *testing.T) {
		t.Parallel()

		file, err := Open(path)
		if err != nil {
			t.Fatal(err)
		}
		if err := file.Close(); err != nil {
			t.Fatalf("unexpected error from Close: %v", err)
		}
		if err := file.Close(); err == nil {
			t.Error("second Close did not fail")
		} else if pe, ok := err.(*PathError); !ok {
			t.Errorf("second Close: got %T, want %T", err, pe)
		} else if pe.Err != ErrClosed {
			t.Errorf("second Close: got %q, want %q", pe.Err, ErrClosed)
		} else {
			t.Logf("second close returned expected error %q", err)
		}
	}
}

func TestDoubleCloseError(t *testing.T) {
	t.Parallel()
	t.Run("file", testDoubleCloseError(filepath.Join(sfdir, sfname)))
	t.Run("dir", testDoubleCloseError(sfdir))
}

func TestUserCacheDir(t *testing.T) {
	t.Parallel()

	dir, err := UserCacheDir()
	if err != nil {
		t.Skipf("skipping: %v", err)
	}
	if dir == "" {
		t.Fatalf("UserCacheDir returned %q; want non-empty path or error", dir)
	}

	fi, err := Stat(dir)
	if err != nil {
		if IsNotExist(err) {
			t.Log(err)
			return
		}
		t.Fatal(err)
	}
	if !fi.IsDir() {
		t.Fatalf("dir %s is not directory; type = %v", dir, fi.Mode())
	}
}

func TestUserConfigDir(t *testing.T) {
	t.Parallel()

	dir, err := UserConfigDir()
	if err != nil {
		t.Skipf("skipping: %v", err)
	}
	if dir == "" {
		t.Fatalf("UserConfigDir returned %q; want non-empty path or error", dir)
	}

	fi, err := Stat(dir)
	if err != nil {
		if IsNotExist(err) {
			t.Log(err)
			return
		}
		t.Fatal(err)
	}
	if !fi.IsDir() {
		t.Fatalf("dir %s is not directory; type = %v", dir, fi.Mode())
	}
}

func TestUserHomeDir(t *testing.T) {
	t.Parallel()

	dir, err := UserHomeDir()
	if dir == "" && err == nil {
		t.Fatal("UserHomeDir returned an empty string but no error")
	}
	if err != nil {
		// UserHomeDir 如果环境变量指定的家目录为空或未设置，可能会返回一个非nil的错误。
		t.Skipf("skipping: %v", err)
	}

	fi, err := Stat(dir)
	if err != nil {
		if IsNotExist(err) {
			// 用户的主目录有一个明确的位置，但可能不存在。（也许还没有写入内容？这在用于CI测试的最小VM映像中可能发生。）
			t.Log(err)
			return
		}
		t.Fatal(err)
	}
	if !fi.IsDir() {
		t.Fatalf("dir %s is not directory; type = %v", dir, fi.Mode())
	}
}

func TestDirSeek(t *testing.T) {
	t.Parallel()

	wd, err := Getwd()
	if err != nil {
		t.Fatal(err)
	}
	f, err := Open(wd)
	if err != nil {
		t.Fatal(err)
	}
	dirnames1, err := f.Readdirnames(0)
	if err != nil {
		t.Fatal(err)
	}

	ret, err := f.Seek(0, 0)
	if err != nil {
		t.Fatal(err)
	}
	if ret != 0 {
		t.Fatalf("seek result not zero: %d", ret)
	}

	dirnames2, err := f.Readdirnames(0)
	if err != nil {
		t.Fatal(err)
	}

	if len(dirnames1) != len(dirnames2) {
		t.Fatalf("listings have different lengths: %d and %d\n", len(dirnames1), len(dirnames2))
	}
	for i, n1 := range dirnames1 {
		n2 := dirnames2[i]
		if n1 != n2 {
			t.Fatalf("different name i=%d n1=%s n2=%s\n", i, n1, n2)
		}
	}
}

func TestReaddirSmallSeek(t *testing.T) {
	// 参见问题37161。从目录中只读取一个条目，然后移动到开头再读取一次。我们不应该看到重复的条目。
	t.Parallel()

	wd, err := Getwd()
	if err != nil {
		t.Fatal(err)
	}
	df, err := Open(filepath.Join(wd, "testdata", "issue37161"))
	if err != nil {
		t.Fatal(err)
	}
	names1, err := df.Readdirnames(1)
	if err != nil {
		t.Fatal(err)
	}
	if _, err = df.Seek(0, 0); err != nil {
		t.Fatal(err)
	}
	names2, err := df.Readdirnames(0)
	if err != nil {
		t.Fatal(err)
	}
	if len(names2) != 3 {
		t.Fatalf("first names: %v, second names: %v", names1, names2)
	}
}

//// isDeadlineExceeded 判断 err 是否为 ErrDeadlineExceeded 错误或其封装错误。
//// 同时，我们还会检查该错误是否具有返回 true 的 Timeout 方法。
//func isDeadlineExceeded(err error) bool {
//	if !IsTimeout(err) {
//		return false
//	}
//	if !errors.Is(err, ErrDeadlineExceeded) {
//		return false
//	}
//	return true
//}

// 测试打开一个文件不会改变其权限。问题38225。
func TestOpenFileKeepsPermissions(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	name := filepath.Join(dir, "x")
	f, err := Create(name)
	if err != nil {
		t.Fatal(err)
	}
	if err := f.Close(); err != nil {
		t.Error(err)
	}
	f, err = OpenFile(name, O_WRONLY|O_CREATE|O_TRUNC, 0)
	if err != nil {
		t.Fatal(err)
	}
	if fi, err := f.Stat(); err != nil {
		t.Error(err)
	} else if fi.Mode()&0222 == 0 {
		t.Errorf("f.Stat.Mode after OpenFile is %v, should be writable", fi.Mode())
	}
	if err := f.Close(); err != nil {
		t.Error(err)
	}
	if fi, err := Stat(name); err != nil {
		t.Error(err)
	} else if fi.Mode()&0222 == 0 {
		t.Errorf("Stat after OpenFile is %v, should be writable", fi.Mode())
	}
}

func TestDirFS(t *testing.T) {
	t.Parallel()

	// 在Windows系统上，我们通过从GetFileInformationByHandle获取实际元数据并明确设置它来强制更新MFT。否则它可能会与FindFirstFile不一致。参见golang.org/issues/42637。
	if runtime.GOOS == "windows" {
		if err := filepath.WalkDir("./testdata/dirfs", func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				t.Fatal(err)
			}
			info, err := d.Info()
			if err != nil {
				t.Fatal(err)
			}
			stat, err := Stat(path) // 这内部使用了GetFileInformationByHandle。
			if err != nil {
				t.Fatal(err)
			}
			if stat.ModTime() == info.ModTime() {
				return nil
			}
			if err := Chtimes(path, stat.ModTime(), stat.ModTime()); err != nil {
				t.Log(err) // 仅在测试目录不可写时进行日志记录，而不强制退出。
			}
			return nil
		}); err != nil {
			t.Fatal(err)
		}
	}
	fsys := DirFS("./testdata/dirfs")
	if err := fstest.TestFS(fsys, "a", "b", "dir/x"); err != nil {
		t.Fatal(err)
	}

	rdfs, ok := fsys.(fs.ReadDirFS)
	if !ok {
		t.Error("expected DirFS result to implement fs.ReadDirFS")
	}
	if _, err := rdfs.ReadDir("nonexistent"); err == nil {
		t.Error("fs.ReadDir of nonexistent directory succeeded")
	}

	// 测试错误消息中不包含反斜杠，
	// 且不包含DirFS参数。
	const nonesuch = "dir/nonesuch"
	_, err := fsys.Open(nonesuch)
	if err == nil {
		t.Error("fs.Open of nonexistent file succeeded")
	} else {
		if !strings.Contains(err.Error(), nonesuch) {
			t.Errorf("error %q does not contain %q", err, nonesuch)
		}
		if strings.Contains(err.(*PathError).Path, "testdata") {
			t.Errorf("error %q contains %q", err, "testdata")
		}
	}

	// 测试Open方法是否接受反斜杠作为分隔符。
	d := DirFS(".")
	_, err = d.Open(`testdata\dirfs`)
	if err == nil {
		t.Fatalf(`Open testdata\dirfs succeeded`)
	}

	// 测试Open是否不会打开Windows设备文件。
	_, err = d.Open(`NUL`)
	if err == nil {
		t.Errorf(`Open NUL succeeded`)
	}
}

func TestDirFSRootDir(t *testing.T) {
	t.Parallel()

	cwd, err := Getwd()
	if err != nil {
		t.Fatal(err)
	}
	cwd = cwd[len(filepath.VolumeName(cwd)):] // 在Windows上移除卷标前缀（如C:)
	cwd = filepath.ToSlash(cwd)               // convert \ to /
	cwd = strings.TrimPrefix(cwd, "/")        // trim leading /

	// 测试Open函数是否可以打开以/开头的路径。
	d := DirFS("/")
	f, err := d.Open(cwd + "/testdata/dirfs/a")
	if err != nil {
		t.Fatal(err)
	}
	f.Close()
}

func TestDirFSEmptyDir(t *testing.T) {
	t.Parallel()

	d := DirFS("")
	cwd, _ := Getwd()
	for _, path := range []string{
		"testdata/dirfs/a",                          // not DirFS(".")
		filepath.ToSlash(cwd) + "/testdata/dirfs/a", // not DirFS("/")
	} {
		_, err := d.Open(path)
		if err == nil {
			t.Fatalf(`DirFS("").Open(%q) succeeded`, path)
		}
	}
}

func TestDirFSPathsValid(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skipf("skipping on Windows")
	}
	t.Parallel()

	d := t.TempDir()
	if err := WriteFile(filepath.Join(d, "control.txt"), []byte(string("Hello, world!")), 0644); err != nil {
		t.Fatal(err)
	}
	if err := WriteFile(filepath.Join(d, `e:xperi\ment.txt`), []byte(string("Hello, colon and backslash!")), 0644); err != nil {
		t.Fatal(err)
	}

	fsys := DirFS(d)
	err := fs.WalkDir(fsys, ".", func(path string, e fs.DirEntry, err error) error {
		if fs.ValidPath(e.Name()) {
			t.Logf("%q ok", e.Name())
		} else {
			t.Errorf("%q INVALID", e.Name())
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestReadFileProc(t *testing.T) {
	t.Parallel()

	// 在Linux中，/proc目录下的文件报告大小为0，
	// 但如果使用ReadFile在偏移量0处读取单个字节后，
	// 偏移量1处的读取会返回EOF（文件结束）而不是更多的数据。
	// ReadFile有一个最小读取大小为512的处理机制来解决这个问题，
	// 但我们需要明确测试它是否有效。
	name := "/proc/sys/fs/pipe-max-size"
	if _, err := Stat(name); err != nil {
		t.Skip(err)
	}
	data, err := ReadFile(name)
	if err != nil {
		t.Fatal(err)
	}
	if len(data) == 0 || data[len(data)-1] != '\n' {
		t.Fatalf("read %s: not newline-terminated: %q", name, data)
	}
}

func TestDirFSReadFileProc(t *testing.T) {
	t.Parallel()

	fsys := DirFS("/")
	name := "proc/sys/fs/pipe-max-size"
	if _, err := fs.Stat(fsys, name); err != nil {
		t.Skip()
	}
	data, err := fs.ReadFile(fsys, name)
	if err != nil {
		t.Fatal(err)
	}
	if len(data) == 0 || data[len(data)-1] != '\n' {
		t.Fatalf("read %s: not newline-terminated: %q", name, data)
	}
}

func TestWriteStringAlloc(t *testing.T) {
	if runtime.GOOS == "js" {
		t.Skip("js allocates a lot during File.WriteString")
	}
	d := t.TempDir()
	f, err := Create(filepath.Join(d, "whiteboard.txt"))
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	allocs := testing.AllocsPerRun(100, func() {
		f.WriteString("I will not allocate when passed a string longer than 32 bytes.\n")
	})
	if allocs != 0 {
		t.Errorf("expected 0 allocs for File.WriteString, got %v", allocs)
	}
}

// 测试并行I/O和管道Close操作是否正常。
func TestPipeIOCloseRace(t *testing.T) {
	// 在没有管道的wasm上跳过
	if runtime.GOOS == "js" || runtime.GOOS == "wasip1" {
		t.Skipf("skipping on %s: no pipes", runtime.GOOS)
	}
	t.Parallel()

	r, w, err := Pipe()
	if err != nil {
		t.Fatal(err)
	}

	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		for {
			n, err := w.Write([]byte("hi"))
			if err != nil {
				// 我们将错误字符串视为依赖操作系统的
				// 预期错误。
				switch {
				case errors.Is(err, ErrClosed),
					strings.Contains(err.Error(), "broken pipe"),
					strings.Contains(err.Error(), "pipe is being closed"),
					strings.Contains(err.Error(), "hungup channel"):
					// 忽略预期的错误。
				default:
					// Unexpected error.
					t.Error(err)
				}
				return
			}
			if n != 2 {
				t.Errorf("wrote %d bytes, expected 2", n)
				return
			}
		}
	}()

	go func() {
		defer wg.Done()
		for {
			var buf [2]byte
			n, err := r.Read(buf[:])
			if err != nil {
				if err != io.EOF && !errors.Is(err, ErrClosed) {
					t.Error(err)
				}
				return
			}
			if n != 2 {
				t.Errorf("read %d bytes, want 2", n)
			}
		}
	}()

	go func() {
		defer wg.Done()

		// 让其他goroutine开始。这只是为了让测试更好，即使它们不开始，测试仍然会通过。
		time.Sleep(time.Millisecond)

		if err := r.Close(); err != nil {
			t.Error(err)
		}
		if err := w.Close(); err != nil {
			t.Error(err)
		}
	}()

	wg.Wait()
}

// 测试对管道同时调用Close方法是正常的
func TestPipeCloseRace(t *testing.T) {
	// 在没有管道的wasm上跳过
	if runtime.GOOS == "js" || runtime.GOOS == "wasip1" {
		t.Skipf("skipping on %s: no pipes", runtime.GOOS)
	}
	t.Parallel()

	r, w, err := Pipe()
	if err != nil {
		t.Fatal(err)
	}
	var wg sync.WaitGroup
	c := make(chan error, 4)
	f := func() {
		defer wg.Done()
		c <- r.Close()
		c <- w.Close()
	}
	wg.Add(2)
	go f()
	go f()
	nils, errs := 0, 0
	for i := 0; i < 4; i++ {
		err := <-c
		if err == nil {
			nils++
		} else {
			errs++
		}
	}
	if nils != 2 || errs != 2 {
		t.Errorf("got nils %d errs %d, want 2 2", nils, errs)
	}
}

//
// // 测试RandomLen函数
// func TestRandomLen(t *testing.T) {
//     // 重复5次循环
//     for i := 0; i < 5; i++ {
//         dir, err := MkdirTemp(t.TempDir(), "*") // 在临时目录下创建子目录
//         if err != nil {
//             t.Fatal(err) // 如果出错，打印错误并终止测试
//         }
//         base := filepath.Base(dir) // 获取子目录的基名
//         if len(base) > 10 { // 如果基名长度超过10
//             t.Errorf("MkdirTemp返回的基名长度为%d：%s", len(base), base) // 打印错误信息
//         }
//     }
//     // 再次重复5次循环
//     for i := 0; i < 5; i++ {
//         f, err := CreateTemp(t.TempDir(), "*") // 在临时目录下创建临时文件
//         if err != nil {
//             t.Fatal(err) // 如果出错，打印错误并终止测试
//         }
//         base := filepath.Base(f.Name()) // 获取文件的基名
//         f.Close() // 关闭文件
//         if len(base) > 10 { // 如果基名长度超过10
//             t.Errorf("CreateTemp返回的基名长度为%d：%s", len(base), base) // 打印错误信息
//         }
//     }
// }
