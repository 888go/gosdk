// 版权所有 2018 The Go Authors。保留所有权利。
// 使用本源代码须遵循 BSD 风格许可证，
// 该许可证可在 LICENSE 文件中找到。

package os_test

import (
	"bytes"
	"fmt"
	"internal/testenv"
	. "os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"testing"
)

func TestRemoveAll(t *testing.T) {
	t.Parallel()

	tmpDir := t.TempDir()
	if err := RemoveAll(""); err != nil {
		t.Errorf("RemoveAll(\"\"): %v; want nil", err)
	}

	file := filepath.Join(tmpDir, "file")
	path := filepath.Join(tmpDir, "_TestRemoveAll_")
	fpath := filepath.Join(path, "file")
	dpath := filepath.Join(path, "dir")

	// 创建一个普通文件并删除
	fd, err := Create(file)
	if err != nil {
		t.Fatalf("create %q: %s", file, err)
	}
	fd.Close()
	if err = RemoveAll(file); err != nil {
		t.Fatalf("RemoveAll %q (first): %s", file, err)
	}
	if _, err = Lstat(file); err == nil {
		t.Fatalf("Lstat %q succeeded after RemoveAll (first)", file)
	}

	// 创建一个包含一个文件的目录并删除它。
	if err := MkdirAll(path, 0777); err != nil {
		t.Fatalf("MkdirAll %q: %s", path, err)
	}
	fd, err = Create(fpath)
	if err != nil {
		t.Fatalf("create %q: %s", fpath, err)
	}
	fd.Close()
	if err = RemoveAll(path); err != nil {
		t.Fatalf("RemoveAll %q (second): %s", path, err)
	}
	if _, err = Lstat(path); err == nil {
		t.Fatalf("Lstat %q succeeded after RemoveAll (second)", path)
	}

	// 创建带有文件和子目录的目录并将其删除
	if err = MkdirAll(dpath, 0777); err != nil {
		t.Fatalf("MkdirAll %q: %s", dpath, err)
	}
	fd, err = Create(fpath)
	if err != nil {
		t.Fatalf("create %q: %s", fpath, err)
	}
	fd.Close()
	fd, err = Create(dpath + "/file")
	if err != nil {
		t.Fatalf("create %q: %s", fpath, err)
	}
	fd.Close()
	if err = RemoveAll(path); err != nil {
		t.Fatalf("RemoveAll %q (third): %s", path, err)
	}
	if _, err := Lstat(path); err == nil {
		t.Fatalf("Lstat %q succeeded after RemoveAll (third)", path)
	}

	// Chmod 在Windows或wasip1下不被支持，而且以root用户运行测试会导致失败。
	if runtime.GOOS != "windows" && runtime.GOOS != "wasip1" && Getuid() != 0 {
		// 创建包含文件和子目录的目录，并触发错误。
		if err = MkdirAll(dpath, 0777); err != nil {
			t.Fatalf("MkdirAll %q: %s", dpath, err)
		}

		for _, s := range []string{fpath, dpath + "/file1", path + "/zzz"} {
			fd, err = Create(s)
			if err != nil {
				t.Fatalf("create %q: %s", s, err)
			}
			fd.Close()
		}
		if err = Chmod(dpath, 0); err != nil {
			t.Fatalf("Chmod %q 0: %s", dpath, err)
		}

// 在这里不进行错误检查：RemoveAll 或许能够移除 dpath，或许不能；无论如何，我们都想知道它是否能移除 fpath 和 path/zzz。RemoveAll 也可能会成功移除 dpath 的原因包括：
// * 以 root 身份运行
// * 在没有权限的文件系统（如 FAT）上运行
		RemoveAll(path)
		Chmod(dpath, 0777)

		for _, s := range []string{fpath, path + "/zzz"} {
			if _, err = Lstat(s); err == nil {
				t.Fatalf("Lstat %q succeeded after partial RemoveAll", s)
			}
		}
	}
	if err = RemoveAll(path); err != nil {
		t.Fatalf("RemoveAll %q after partial RemoveAll: %s", path, err)
	}
	if _, err = Lstat(path); err == nil {
		t.Fatalf("Lstat %q succeeded after RemoveAll (final)", path)
	}
}

// 测试在大型目录上使用RemoveAll
func TestRemoveAllLarge(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping in short mode")
	}
	t.Parallel()

	tmpDir := t.TempDir()
	path := filepath.Join(tmpDir, "_TestRemoveAllLarge_")

	// 创建一个包含1000个文件的目录，然后删除。
	if err := MkdirAll(path, 0777); err != nil {
		t.Fatalf("MkdirAll %q: %s", path, err)
	}
	for i := 0; i < 1000; i++ {
		fpath := fmt.Sprintf("%s/file%d", path, i)
		fd, err := Create(fpath)
		if err != nil {
			t.Fatalf("create %q: %s", fpath, err)
		}
		fd.Close()
	}
	if err := RemoveAll(path); err != nil {
		t.Fatalf("RemoveAll %q: %s", path, err)
	}
	if _, err := Lstat(path); err == nil {
		t.Fatalf("Lstat %q succeeded after RemoveAll", path)
	}
}

func TestRemoveAllLongPath(t *testing.T) {
	switch runtime.GOOS {
	case "aix", "darwin", "ios", "dragonfly", "freebsd", "linux", "netbsd", "openbsd", "illumos", "solaris":
		break
	default:
		t.Skip("skipping for not implemented platforms")
	}

	prevDir, err := Getwd()
	if err != nil {
		t.Fatalf("Could not get wd: %s", err)
	}

	startPath, err := MkdirTemp("", "TestRemoveAllLongPath-")
	if err != nil {
		t.Fatalf("Could not create TempDir: %s", err)
	}
	defer RemoveAll(startPath)

	err = Chdir(startPath)
	if err != nil {
		t.Fatalf("Could not chdir %s: %s", startPath, err)
	}

	// 删除超过4096个字符的路径通常会失败
	for i := 0; i < 41; i++ {
		name := strings.Repeat("a", 100)

		err = Mkdir(name, 0755)
		if err != nil {
			t.Fatalf("Could not mkdir %s: %s", name, err)
		}

		err = Chdir(name)
		if err != nil {
			t.Fatalf("Could not chdir %s: %s", name, err)
		}
	}

	err = Chdir(prevDir)
	if err != nil {
		t.Fatalf("Could not chdir %s: %s", prevDir, err)
	}

	err = RemoveAll(startPath)
	if err != nil {
		t.Errorf("RemoveAll could not remove long file path %s: %s", startPath, err)
	}
}

func TestRemoveAllDot(t *testing.T) {
	prevDir, err := Getwd()
	if err != nil {
		t.Fatalf("Could not get wd: %s", err)
	}
	tempDir, err := MkdirTemp("", "TestRemoveAllDot-")
	if err != nil {
		t.Fatalf("Could not create TempDir: %s", err)
	}
	defer RemoveAll(tempDir)

	err = Chdir(tempDir)
	if err != nil {
		t.Fatalf("Could not chdir to tempdir: %s", err)
	}

	err = RemoveAll(".")
	if err == nil {
		t.Errorf("RemoveAll succeed to remove .")
	}

	err = Chdir(prevDir)
	if err != nil {
		t.Fatalf("Could not chdir %s: %s", prevDir, err)
	}
}

func TestRemoveAllDotDot(t *testing.T) {
	t.Parallel()

	tempDir := t.TempDir()
	subdir := filepath.Join(tempDir, "x")
	subsubdir := filepath.Join(subdir, "y")
	if err := MkdirAll(subsubdir, 0777); err != nil {
		t.Fatal(err)
	}
	if err := RemoveAll(filepath.Join(subsubdir, "..")); err != nil {
		t.Error(err)
	}
	for _, dir := range []string{subsubdir, subdir} {
		if _, err := Stat(dir); err == nil {
			t.Errorf("%s: exists after RemoveAll", dir)
		}
	}
}

// Issue #29178.
func TestRemoveReadOnlyDir(t *testing.T) {
	t.Parallel()

	tempDir := t.TempDir()
	subdir := filepath.Join(tempDir, "x")
	if err := Mkdir(subdir, 0); err != nil {
		t.Fatal(err)
	}

// 如果发生错误，设法提高删除临时目录成功的可能性。
	defer Chmod(subdir, 0777)

	if err := RemoveAll(subdir); err != nil {
		t.Fatal(err)
	}

	if _, err := Stat(subdir); err == nil {
		t.Error("subdirectory was not removed")
	}
}

// Issue #29983.
func TestRemoveAllButReadOnlyAndPathError(t *testing.T) {
	switch runtime.GOOS {
	case "js", "wasip1", "windows":
		t.Skipf("skipping test on %s", runtime.GOOS)
	}

	if Getuid() == 0 {
		t.Skip("skipping test when running as root")
	}

	t.Parallel()

	tempDir := t.TempDir()
	dirs := []string{
		"a",
		"a/x",
		"a/x/1",
		"b",
		"b/y",
		"b/y/2",
		"c",
		"c/z",
		"c/z/3",
	}
	readonly := []string{
		"b",
	}
	inReadonly := func(d string) bool {
		for _, ro := range readonly {
			if d == ro {
				return true
			}
			dd, _ := filepath.Split(d)
			if filepath.Clean(dd) == ro {
				return true
			}
		}
		return false
	}

	for _, dir := range dirs {
		if err := Mkdir(filepath.Join(tempDir, dir), 0777); err != nil {
			t.Fatal(err)
		}
	}
	for _, dir := range readonly {
		d := filepath.Join(tempDir, dir)
		if err := Chmod(d, 0555); err != nil {
			t.Fatal(err)
		}

// 延迟恢复原模式，以便于后续执行的deferred RemoveAll(tempDir)能够成功。
		defer Chmod(d, 0777)
	}

	err := RemoveAll(tempDir)
	if err == nil {
		t.Fatal("RemoveAll succeeded unexpectedly")
	}

// 错误应为类型 *PathError。
// 有关详细信息，请参阅问题 30491。
	if pathErr, ok := err.(*PathError); ok {
		want := filepath.Join(tempDir, "b", "y")
		if pathErr.Path != want {
			t.Errorf("RemoveAll(%q): err.Path=%q, want %q", tempDir, pathErr.Path, want)
		}
	} else {
		t.Errorf("RemoveAll(%q): error has type %T, want *fs.PathError", tempDir, err)
	}

	for _, dir := range dirs {
		_, err := Stat(filepath.Join(tempDir, dir))
		if inReadonly(dir) {
			if err != nil {
				t.Errorf("file %q was deleted but should still exist", dir)
			}
		} else {
			if err == nil {
				t.Errorf("file %q still exists but should have been deleted", dir)
			}
		}
	}
}

func TestRemoveUnreadableDir(t *testing.T) {
	switch runtime.GOOS {
	case "js":
		t.Skipf("skipping test on %s", runtime.GOOS)
	}

	if Getuid() == 0 {
		t.Skip("skipping test when running as root")
	}

	t.Parallel()

	tempDir := t.TempDir()
	target := filepath.Join(tempDir, "d0", "d1", "d2")
	if err := MkdirAll(target, 0755); err != nil {
		t.Fatal(err)
	}
	if err := Chmod(target, 0300); err != nil {
		t.Fatal(err)
	}
	if err := RemoveAll(filepath.Join(tempDir, "d0")); err != nil {
		t.Fatal(err)
	}
}

// Issue 29921
func TestRemoveAllWithMoreErrorThanReqSize(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping in short mode")
	}
	t.Parallel()

	tmpDir := t.TempDir()
	path := filepath.Join(tmpDir, "_TestRemoveAllWithMoreErrorThanReqSize_")

	// 创建一个包含1025个只读文件的目录。
	if err := MkdirAll(path, 0777); err != nil {
		t.Fatalf("MkdirAll %q: %s", path, err)
	}
	for i := 0; i < 1025; i++ {
		fpath := filepath.Join(path, fmt.Sprintf("file%d", i))
		fd, err := Create(fpath)
		if err != nil {
			t.Fatalf("create %q: %s", fpath, err)
		}
		fd.Close()
	}

// 将父目录设置为只读。在某些平台上，这是阻止Remove删除该目录内文件的原因。
	if err := Chmod(path, 0555); err != nil {
		t.Fatal(err)
	}
	defer Chmod(path, 0755)

// 此调用不应阻塞，即使在不允许从只读目录中删除文件的平台上也是如此
	err := RemoveAll(path)

	if Getuid() == 0 {
		// 在许多平台上，超级用户可以删除只读目录中的文件。
		return
	}
	if err == nil {
		if runtime.GOOS == "windows" || runtime.GOOS == "wasip1" {
// 在Windows中，将目录标记为只读并不能阻止RemoveAll在其中创建或删除文件。
//
// 对于wasip1，没有文件权限支持，因此我们无法阻止RemoveAll删除文件。
			return
		}
		t.Fatal("RemoveAll(<read-only directory>) = nil; want error")
	}

	dir, err := Open(path)
	if err != nil {
		t.Fatal(err)
	}
	defer dir.Close()

	names, _ := dir.Readdirnames(1025)
	if len(names) < 1025 {
		t.Fatalf("RemoveAll(<read-only directory>) unexpectedly removed %d read-only files from that directory", 1025-len(names))
	}
}

func TestRemoveAllNoFcntl(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping in short mode")
	}

	const env = "GO_TEST_REMOVE_ALL_NO_FCNTL"
	if dir := Getenv(env); dir != "" {
		if err := RemoveAll(dir); err != nil {
			t.Fatal(err)
		}
		return
	}

// 只在Linux上进行测试，这样我们就可以假设我们有strace工具。
// 该代码是跨平台的，所以如果在Linux上通过了，那么在其他Unix系统上也应该通过。
	if runtime.GOOS != "linux" {
		t.Skipf("skipping test on %s", runtime.GOOS)
	}
	if _, err := Stat("/bin/strace"); err != nil {
		t.Skipf("skipping test because /bin/strace not found: %v", err)
	}
	me, err := Executable()
	if err != nil {
		t.Skipf("skipping because Executable failed: %v", err)
	}

// 创建100个目录。
// 测试目标是我们可以在不针对每个目录调用fcntl的情况下将其删除。
	tmpdir := t.TempDir()
	subdir := filepath.Join(tmpdir, "subdir")
	if err := Mkdir(subdir, 0o755); err != nil {
		t.Fatal(err)
	}
	for i := 0; i < 100; i++ {
		subsubdir := filepath.Join(subdir, strconv.Itoa(i))
		if err := Mkdir(filepath.Join(subdir, strconv.Itoa(i)), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := WriteFile(filepath.Join(subsubdir, "file"), nil, 0o644); err != nil {
			t.Fatal(err)
		}
	}

	cmd := testenv.Command(t, "/bin/strace", "-f", "-e", "fcntl", me, "-test.run=^TestRemoveAllNoFcntl$")
	cmd = testenv.CleanCmdEnv(cmd)
	cmd.Env = append(cmd.Env, env+"="+subdir)
	out, err := cmd.CombinedOutput()
	if len(out) > 0 {
		t.Logf("%s", out)
	}
	if err != nil {
		t.Fatal(err)
	}

	if got := bytes.Count(out, []byte("fcntl")); got >= 100 {
		t.Errorf("found %d fcntl calls, want < 100", got)
	}
}
