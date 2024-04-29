// 版权所有 2014 Go 作者。保留所有权利。
// 本源代码的使用由bsd风格的许可证控制，该许可证可在LICENSE文件中找到。

package os_test

import (
	"errors"
	"fmt"
	"github.com/888go/gosdk/os"
	"github.com/888go/gosdk/os/internal/syscall/windows"
	"github.com/888go/gosdk/os/internal/syscall/windows/registry"
	"github.com/888go/gosdk/os/internal/testenv"
	"io/fs"
	"os/exec"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"syscall"
	"testing"
	"unsafe"
)

// 用于TestRawConnReadWrite测试。
type syscallDescriptor = syscall.Handle

// chdir 将当前工作目录更改为指定的目录，然后在测试结束时恢复原来的目录。
func chdir(t *testing.T, dir string) {
	olddir, err := os.Getwd()
	if err != nil {
		t.Fatalf("chdir: %v", err)
	}
	if err := os.Chdir(dir); err != nil {
		t.Fatalf("chdir %s: %v", dir, err)
	}

	t.Cleanup(func() {
		if err := os.Chdir(olddir); err != nil {
			t.Errorf("chdir to original working directory %s: %v", olddir, err)
			os.Exit(1)
		}
	})
}

func TestSameWindowsFile(t *testing.T) {
	temp := t.TempDir()
	chdir(t, temp)

	f, err := os.Create("a")
	if err != nil {
		t.Fatal(err)
	}
	f.Close()

	ia1, err := os.Stat("a")
	if err != nil {
		t.Fatal(err)
	}

	path, err := filepath.Abs("a")
	if err != nil {
		t.Fatal(err)
	}
	ia2, err := os.Stat(path)
	if err != nil {
		t.Fatal(err)
	}
	if !os.SameFile(ia1, ia2) {
		t.Errorf("files should be same")
	}

	p := filepath.VolumeName(path) + filepath.Base(path)
	if err != nil {
		t.Fatal(err)
	}
	ia3, err := os.Stat(p)
	if err != nil {
		t.Fatal(err)
	}
	if !os.SameFile(ia1, ia3) {
		t.Errorf("files should be same")
	}
}

type dirLinkTest struct {
	name    string
	mklink  func(link, target string) error
	issueNo int // 对应的问题编号（针对失效的测试）
}

func testDirLinks(t *testing.T, tests []dirLinkTest) {
	tmpdir := t.TempDir()
	chdir(t, tmpdir)

	dir := filepath.Join(tmpdir, "dir")
	err := os.Mkdir(dir, 0777)
	if err != nil {
		t.Fatal(err)
	}
	fi, err := os.Stat(dir)
	if err != nil {
		t.Fatal(err)
	}
	err = os.WriteFile(filepath.Join(dir, "abc"), []byte("abc"), 0644)
	if err != nil {
		t.Fatal(err)
	}
	for _, test := range tests {
		link := filepath.Join(tmpdir, test.name+"_link")
		err := test.mklink(link, dir)
		if err != nil {
			t.Errorf("creating link for %q test failed: %v", test.name, err)
			continue
		}

		data, err := os.ReadFile(filepath.Join(link, "abc"))
		if err != nil {
			t.Errorf("failed to read abc file: %v", err)
			continue
		}
		if string(data) != "abc" {
			t.Errorf(`abc file is expected to have "abc" in it, but has %v`, data)
			continue
		}

		if test.issueNo > 0 {
			t.Logf("skipping broken %q test: see issue %d", test.name, test.issueNo)
			continue
		}

		fi1, err := os.Stat(link)
		if err != nil {
			t.Errorf("failed to stat link %v: %v", link, err)
			continue
		}
		if !fi1.IsDir() {
			t.Errorf("%q should be a directory", link)
			continue
		}
		if fi1.Name() != filepath.Base(link) {
			t.Errorf("Stat(%q).Name() = %q, want %q", link, fi1.Name(), filepath.Base(link))
			continue
		}
		if !os.SameFile(fi, fi1) {
			t.Errorf("%q should point to %q", link, dir)
			continue
		}

		fi2, err := os.Lstat(link)
		if err != nil {
			t.Errorf("failed to lstat link %v: %v", link, err)
			continue
		}
		if m := fi2.Mode(); m&fs.ModeSymlink == 0 {
			t.Errorf("%q should be a link, but is not (mode=0x%x)", link, uint32(m))
			continue
		}
		if m := fi2.Mode(); m&fs.ModeDir != 0 {
			t.Errorf("%q should be a link, not a directory (mode=0x%x)", link, uint32(m))
			continue
		}
	}
}

// reparseData 用于构建测试所需的重新解析缓冲区数据。
type reparseData struct {
	substituteName namePosition
	printName      namePosition
	pathBuf        []uint16
}

type namePosition struct {
	offset uint16
	length uint16
}

func (rd *reparseData) addUTF16s(s []uint16) (offset uint16) {
	off := len(rd.pathBuf) * 2
	rd.pathBuf = append(rd.pathBuf, s...)
	return uint16(off)
}

func (rd *reparseData) addString(s string) (offset, length uint16) {
	p := syscall.StringToUTF16(s)
	return rd.addUTF16s(p), uint16(len(p)-1) * 2 // 不要在长度中包含终止 NULL（根据 PrintNameLength 和 SubstituteNameLength 文档）
}

func (rd *reparseData) addSubstituteName(name string) {
	rd.substituteName.offset, rd.substituteName.length = rd.addString(name)
}

func (rd *reparseData) addPrintName(name string) {
	rd.printName.offset, rd.printName.length = rd.addString(name)
}

func (rd *reparseData) addStringNoNUL(s string) (offset, length uint16) {
	p := syscall.StringToUTF16(s)
	p = p[:len(p)-1]
	return rd.addUTF16s(p), uint16(len(p)) * 2
}

func (rd *reparseData) addSubstituteNameNoNUL(name string) {
	rd.substituteName.offset, rd.substituteName.length = rd.addStringNoNUL(name)
}

func (rd *reparseData) addPrintNameNoNUL(name string) {
	rd.printName.offset, rd.printName.length = rd.addStringNoNUL(name)
}

// pathBuffeLen 返回rd路径缓冲区的字节长度。
func (rd *reparseData) pathBuffeLen() uint16 {
	return uint16(len(rd.pathBuf)) * 2
}

// Windows REPARSE_DATA_BUFFER 结构包含联合成员，无法直接转换为 Go 语言表示。_REPARSE_DATA_BUFFER 类型旨在帮助构建 Windows REPARSE_DATA_BUFFER 的替代版本，其中联合部分为 SymbolicLinkReparseBuffer 或 MountPointReparseBuffer 类型。
type _REPARSE_DATA_BUFFER struct {
	header windows.REPARSE_DATA_BUFFER_HEADER
	detail [syscall.MAXIMUM_REPARSE_DATA_BUFFER_SIZE]byte
}

func createDirLink(link string, rdb *_REPARSE_DATA_BUFFER) error {
	err := os.Mkdir(link, 0777)
	if err != nil {
		return err
	}

	linkp := syscall.StringToUTF16(link)
	fd, err := syscall.CreateFile(&linkp[0], syscall.GENERIC_WRITE, 0, nil, syscall.OPEN_EXISTING,
		syscall.FILE_FLAG_OPEN_REPARSE_POINT|syscall.FILE_FLAG_BACKUP_SEMANTICS, 0)
	if err != nil {
		return err
	}
	defer syscall.CloseHandle(fd)

	buflen := uint32(rdb.header.ReparseDataLength) + uint32(unsafe.Sizeof(rdb.header))
	var bytesReturned uint32
	return syscall.DeviceIoControl(fd, windows.FSCTL_SET_REPARSE_POINT,
		(*byte)(unsafe.Pointer(&rdb.header)), buflen, nil, 0, &bytesReturned, nil)
}

func createMountPoint(link string, target *reparseData) error {
	var buf *windows.MountPointReparseBuffer
	buflen := uint16(unsafe.Offsetof(buf.PathBuffer)) + target.pathBuffeLen() // 参考ReparseDataLength的文档说明
	byteblob := make([]byte, buflen)
	buf = (*windows.MountPointReparseBuffer)(unsafe.Pointer(&byteblob[0]))
	buf.SubstituteNameOffset = target.substituteName.offset
	buf.SubstituteNameLength = target.substituteName.length
	buf.PrintNameOffset = target.printName.offset
	buf.PrintNameLength = target.printName.length
	pbuflen := len(target.pathBuf)
	copy((*[2048]uint16)(unsafe.Pointer(&buf.PathBuffer[0]))[:pbuflen:pbuflen], target.pathBuf)

	var rdb _REPARSE_DATA_BUFFER
	rdb.header.ReparseTag = windows.IO_REPARSE_TAG_MOUNT_POINT
	rdb.header.ReparseDataLength = buflen
	copy(rdb.detail[:], byteblob)

	return createDirLink(link, &rdb)
}

func TestDirectoryJunction(t *testing.T) {
	var tests = []dirLinkTest{
		{
			// 创建类似于mklink所做的链接，通过在绝对目标的前面插入\??\。
			name: "standard",
			mklink: func(link, target string) error {
				var t reparseData
				t.addSubstituteName(`\??\` + target)
				t.addPrintName(target)
				return createMountPoint(link, &t)
			},
		},
		{
			// 像junction实用程序（https://learn.microsoft.com/en-us/sysinternals/downloads/junction）一样操作 - 将PrintNameLength设置为0。
			name: "have_blank_print_name",
			mklink: func(link, target string) error {
				var t reparseData
				t.addSubstituteName(`\??\` + target)
				t.addPrintName("")
				return createMountPoint(link, &t)
			},
		},
	}
	output, _ := testenv.Command(t, "cmd", "/c", "mklink", "/?").Output()
	mklinkSupportsJunctionLinks := strings.Contains(string(output), " /J ")
	if mklinkSupportsJunctionLinks {
		tests = append(tests,
			dirLinkTest{
				name: "use_mklink_cmd",
				mklink: func(link, target string) error {
					output, err := testenv.Command(t, "cmd", "/c", "mklink", "/J", link, target).CombinedOutput()
					if err != nil {
						t.Errorf("failed to run mklink %v %v: %v %q", link, target, err, output)
					}
					return nil
				},
			},
		)
	} else {
		t.Log(`skipping "use_mklink_cmd" test, mklink does not supports directory junctions`)
	}
	testDirLinks(t, tests)
}

func enableCurrentThreadPrivilege(privilegeName string) error {
	ct, err := windows.GetCurrentThread()
	if err != nil {
		return err
	}
	var t syscall.Token
	err = windows.OpenThreadToken(ct, syscall.TOKEN_QUERY|windows.TOKEN_ADJUST_PRIVILEGES, false, &t)
	if err != nil {
		return err
	}
	defer syscall.CloseHandle(syscall.Handle(t))

	var tp windows.TOKEN_PRIVILEGES

	privStr, err := syscall.UTF16PtrFromString(privilegeName)
	if err != nil {
		return err
	}
	err = windows.LookupPrivilegeValue(nil, privStr, &tp.Privileges[0].Luid)
	if err != nil {
		return err
	}
	tp.PrivilegeCount = 1
	tp.Privileges[0].Attributes = windows.SE_PRIVILEGE_ENABLED
	return windows.AdjustTokenPrivileges(t, false, &tp, 0, nil, nil)
}

func createSymbolicLink(link string, target *reparseData, isrelative bool) error {
	var buf *windows.SymbolicLinkReparseBuffer
	buflen := uint16(unsafe.Offsetof(buf.PathBuffer)) + target.pathBuffeLen() // 参考ReparseDataLength的文档说明
	byteblob := make([]byte, buflen)
	buf = (*windows.SymbolicLinkReparseBuffer)(unsafe.Pointer(&byteblob[0]))
	buf.SubstituteNameOffset = target.substituteName.offset
	buf.SubstituteNameLength = target.substituteName.length
	buf.PrintNameOffset = target.printName.offset
	buf.PrintNameLength = target.printName.length
	if isrelative {
		buf.Flags = windows.SYMLINK_FLAG_RELATIVE
	}
	pbuflen := len(target.pathBuf)
	copy((*[2048]uint16)(unsafe.Pointer(&buf.PathBuffer[0]))[:pbuflen:pbuflen], target.pathBuf)

	var rdb _REPARSE_DATA_BUFFER
	rdb.header.ReparseTag = syscall.IO_REPARSE_TAG_SYMLINK
	rdb.header.ReparseDataLength = buflen
	copy(rdb.detail[:], byteblob)

	return createDirLink(link, &rdb)
}

func TestDirectorySymbolicLink(t *testing.T) {
	var tests []dirLinkTest
	output, _ := testenv.Command(t, "cmd", "/c", "mklink", "/?").Output()
	mklinkSupportsDirectorySymbolicLinks := strings.Contains(string(output), " /D ")
	if mklinkSupportsDirectorySymbolicLinks {
		tests = append(tests,
			dirLinkTest{
				name: "use_mklink_cmd",
				mklink: func(link, target string) error {
					output, err := testenv.Command(t, "cmd", "/c", "mklink", "/D", link, target).CombinedOutput()
					if err != nil {
						t.Errorf("failed to run mklink %v %v: %v %q", link, target, err, output)
					}
					return nil
				},
			},
		)
	} else {
		t.Log(`skipping "use_mklink_cmd" test, mklink does not supports directory symbolic links`)
	}

	// 余下的这些测试要求拥有SeCreateSymbolicLinkPrivilege权限。
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	err := windows.ImpersonateSelf(windows.SecurityImpersonation)
	if err != nil {
		t.Fatal(err)
	}
	defer windows.RevertToSelf()

	err = enableCurrentThreadPrivilege("SeCreateSymbolicLinkPrivilege")
	if err != nil {
		t.Skipf(`skipping some tests, could not enable "SeCreateSymbolicLinkPrivilege": %v`, err)
	}
	tests = append(tests,
		dirLinkTest{
			name: "use_os_pkg",
			mklink: func(link, target string) error {
				return os.Symlink(target, link)
			},
		},
		dirLinkTest{
			// 创建类似于mklink所做的链接，通过在绝对目标的前面插入\??\。
			name: "standard",
			mklink: func(link, target string) error {
				var t reparseData
				t.addPrintName(target)
				t.addSubstituteName(`\??\` + target)
				return createSymbolicLink(link, &t, false)
			},
		},
		dirLinkTest{
			name: "relative",
			mklink: func(link, target string) error {
				var t reparseData
				t.addSubstituteNameNoNUL(filepath.Base(target))
				t.addPrintNameNoNUL(filepath.Base(target))
				return createSymbolicLink(link, &t, true)
			},
		},
	)
	testDirLinks(t, tests)
}

func mustHaveWorkstation(t *testing.T) {
	mar, err := windows.OpenSCManager(nil, nil, windows.SERVICE_QUERY_STATUS)
	if err != nil {
		return
	}
	defer syscall.CloseHandle(mar)
	//LanmanWorkstation是服务名称，而Workstation是显示名称。
	srv, err := windows.OpenService(mar, syscall.StringToUTF16Ptr("LanmanWorkstation"), windows.SERVICE_QUERY_STATUS)
	if err != nil {
		return
	}
	defer syscall.CloseHandle(srv)
	var state windows.SERVICE_STATUS
	err = windows.QueryServiceStatus(srv, &state)
	if err != nil {
		return
	}
	if state.CurrentState != windows.SERVICE_RUNNING {
		t.Skip("Requires the Windows service Workstation, but it is detected that it is not enabled.")
	}
}

func TestNetworkSymbolicLink(t *testing.T) {
	testenv.MustHaveSymlink(t)

	const _NERR_ServerNotStarted = syscall.Errno(2114)

	dir := t.TempDir()
	chdir(t, dir)

	pid := os.Getpid()
	shareName := fmt.Sprintf("GoSymbolicLinkTestShare%d", pid)
	sharePath := filepath.Join(dir, shareName)
	testDir := "TestDir"

	err := os.MkdirAll(filepath.Join(sharePath, testDir), 0777)
	if err != nil {
		t.Fatal(err)
	}

	wShareName, err := syscall.UTF16PtrFromString(shareName)
	if err != nil {
		t.Fatal(err)
	}
	wSharePath, err := syscall.UTF16PtrFromString(sharePath)
	if err != nil {
		t.Fatal(err)
	}

	// Per https://learn.microsoft.com/en-us/windows/win32/api/lmshare/ns-lmshare-share_info_2:
	//
	// “[The shi2_permissions field] indicates the shared resource's permissions
	// for servers running with share-level security. A server running user-level
	// security ignores this member.
	// …
	// Note that Windows does not support share-level security.”
	//
	// So it shouldn't matter what permissions we set here.
	const permissions = 0

	p := windows.SHARE_INFO_2{
		Netname:     wShareName,
		Type:        windows.STYPE_DISKTREE | windows.STYPE_TEMPORARY,
		Remark:      nil,
		Permissions: permissions,
		MaxUses:     1,
		CurrentUses: 0,
		Path:        wSharePath,
		Passwd:      nil,
	}

	err = windows.NetShareAdd(nil, 2, (*byte)(unsafe.Pointer(&p)), nil)
	if err != nil {
		if err == syscall.ERROR_ACCESS_DENIED || err == _NERR_ServerNotStarted {
			t.Skipf("skipping: NetShareAdd: %v", err)
		}
		t.Fatal(err)
	}
	defer func() {
		err := windows.NetShareDel(nil, wShareName, 0)
		if err != nil {
			t.Fatal(err)
		}
	}()

	UNCPath := `\\localhost\` + shareName + `\`

	fi1, err := os.Stat(sharePath)
	if err != nil {
		t.Fatal(err)
	}
	fi2, err := os.Stat(UNCPath)
	if err != nil {
		mustHaveWorkstation(t)
		t.Fatal(err)
	}
	if !os.SameFile(fi1, fi2) {
		t.Fatalf("%q and %q should be the same directory, but not", sharePath, UNCPath)
	}

	target := filepath.Join(UNCPath, testDir)
	link := "link"

	err = os.Symlink(target, link)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(link)

	got, err := os.Readlink(link)
	if err != nil {
		t.Fatal(err)
	}
	if got != target {
		t.Errorf(`os.Readlink(%#q): got %v, want %v`, link, got, target)
	}

	got, err = filepath.EvalSymlinks(link)
	if err != nil {
		t.Fatal(err)
	}
	if got != target {
		t.Errorf(`filepath.EvalSymlinks(%#q): got %v, want %v`, link, got, target)
	}
}

func TestStatLxSymLink(t *testing.T) {
	if _, err := exec.LookPath("wsl"); err != nil {
		t.Skip("skipping: WSL not detected")
	}

	temp := t.TempDir()
	chdir(t, temp)

	const target = "target"
	const link = "link"

	_, err := testenv.Command(t, "wsl", "/bin/mkdir", target).Output()
	if err != nil {
		// 这通常发生在WSL尚未安装运行的发行版时。
		t.Skipf("skipping: WSL is not correctly installed: %v", err)
	}

	_, err = testenv.Command(t, "wsl", "/bin/ln", "-s", target, link).Output()
	if err != nil {
		t.Fatal(err)
	}

	fi, err := os.Lstat(link)
	if err != nil {
		t.Fatal(err)
	}
	if m := fi.Mode(); m&fs.ModeSymlink != 0 {
		// 如果以管理员身份运行或在开发模式下运行，根据较新的WSL版本，可能会发生这种情况。
		t.Skip("skipping: WSL created reparse tag IO_REPARSE_TAG_SYMLINK instead of an IO_REPARSE_TAG_LX_SYMLINK")
	}
	// 对于一个IO_REPARSE_TAG_LX_SYMLINK类型的符号链接，若从WSL外部进行Stat操作，总会返回ERROR_CANT_ACCESS_FILE错误。
	// 我们检查这一条件，以验证os.Stat确实尝试了跟踪该链接。
	_, err = os.Stat(link)
	const ERROR_CANT_ACCESS_FILE = syscall.Errno(1920)
	if err == nil || !errors.Is(err, ERROR_CANT_ACCESS_FILE) {
		t.Fatalf("os.Stat(%q): got %v, want ERROR_CANT_ACCESS_FILE", link, err)
	}
}

func TestStartProcessAttr(t *testing.T) {
	t.Parallel()

	p, err := os.StartProcess(os.Getenv("COMSPEC"), []string{"/c", "cd"}, new(os.ProcAttr))
	if err != nil {
		return
	}
	defer p.Wait()
	t.Fatalf("StartProcess expected to fail, but succeeded.")
}

func TestShareNotExistError(t *testing.T) {
	if testing.Short() {
		t.Skip("slow test that uses network; skipping")
	}
	t.Parallel()

	_, err := os.Stat(`\\no_such_server\no_such_share\no_such_file`)
	if err == nil {
		t.Fatal("stat succeeded, but expected to fail")
	}
	if !os.IsNotExist(err) {
		t.Fatalf("os.Stat failed with %q, but os.IsNotExist(err) is false", err)
	}
}

func TestBadNetPathError(t *testing.T) {
	const ERROR_BAD_NETPATH = syscall.Errno(53)
	if !os.IsNotExist(ERROR_BAD_NETPATH) {
		t.Fatal("os.IsNotExist(syscall.Errno(53)) is false, but want true")
	}
}

func TestStatDir(t *testing.T) {
	defer chtmpdir(t)()

	f, err := os.Open(".")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		t.Fatal(err)
	}

	err = os.Chdir("..")
	if err != nil {
		t.Fatal(err)
	}

	fi2, err := f.Stat()
	if err != nil {
		t.Fatal(err)
	}

	if !os.SameFile(fi, fi2) {
		t.Fatal("race condition occurred")
	}
}

func TestOpenVolumeName(t *testing.T) {
	tmpdir := t.TempDir()
	chdir(t, tmpdir)

	want := []string{"file1", "file2", "file3", "gopher.txt"}
	sort.Strings(want)
	for _, name := range want {
		err := os.WriteFile(filepath.Join(tmpdir, name), nil, 0777)
		if err != nil {
			t.Fatal(err)
		}
	}

	f, err := os.Open(filepath.VolumeName(tmpdir))
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	have, err := f.Readdirnames(-1)
	if err != nil {
		t.Fatal(err)
	}
	sort.Strings(have)

	if strings.Join(want, "/") != strings.Join(have, "/") {
		t.Fatalf("unexpected file list %q, want %q", have, want)
	}
}

func TestDeleteReadOnly(t *testing.T) {
	t.Parallel()

	tmpdir := t.TempDir()
	p := filepath.Join(tmpdir, "a")
	// 这将设置 FILE_ATTRIBUTE_READONLY 属性。
	f, err := os.OpenFile(p, os.O_CREATE, 0400)
	if err != nil {
		t.Fatal(err)
	}
	f.Close()

	if err = os.Chmod(p, 0400); err != nil {
		t.Fatal(err)
	}
	if err = os.Remove(p); err != nil {
		t.Fatal(err)
	}
}

//func TestReadStdin(t *testing.T) {
//	old := poll.ReadConsole
//	defer func() {
//		poll.ReadConsole = old
//	}()
//
//	p, err := syscall.GetCurrentProcess()
//	if err != nil {
//		t.Fatalf("Unable to get handle to current process: %v", err)
//	}
//	var stdinDuplicate syscall.Handle
//	err = syscall.DuplicateHandle(p, syscall.Handle(syscall.Stdin), p, &stdinDuplicate, 0, false, syscall.DUPLICATE_SAME_ACCESS)
//	if err != nil {
//		t.Fatalf("Unable to duplicate stdin: %v", err)
//	}
//	testConsole := os.NewConsoleFile(stdinDuplicate, "test")
//
//	var tests = []string{
//		"abc",
//		"äöü",
//		"\u3042",
//		"“hi”™",
//		"hello\x1aworld",
//		"\U0001F648\U0001F649\U0001F64A",
//	}
//
//	for _, consoleSize := range []int{1, 2, 3, 10, 16, 100, 1000} {
//		for _, readSize := range []int{1, 2, 3, 4, 5, 8, 10, 16, 20, 50, 100} {
//			for _, s := range tests {
//				t.Run(fmt.Sprintf("c%d/r%d/%s", consoleSize, readSize, s), func(t *testing.T) {
//					s16 := utf16.Encode([]rune(s))
//					poll.ReadConsole = func(h syscall.Handle, buf *uint16, toread uint32, read *uint32, inputControl *byte) error {
//						if inputControl != nil {
//							t.Fatalf("inputControl not nil")
//						}
//						n := int(toread)
//						if n > consoleSize {
//							n = consoleSize
//						}
//						n = copy((*[10000]uint16)(unsafe.Pointer(buf))[:n:n], s16)
//						s16 = s16[n:]
//						*read = uint32(n)
//						t.Logf("read %d -> %d", toread, *read)
//						return nil
//					}
//
//					var all []string
//					var buf []byte
//					chunk := make([]byte, readSize)
//					for {
//						n, err := testConsole.Read(chunk)
//						buf = append(buf, chunk[:n]...)
//						if err == io.EOF {
//							all = append(all, string(buf))
//							if len(all) >= 5 {
//								break
//							}
//							buf = buf[:0]
//						} else if err != nil {
//							t.Fatalf("reading %q: error: %v", s, err)
//						}
//						if len(buf) >= 2000 {
//							t.Fatalf("reading %q: stuck in loop: %q", s, buf)
//						}
//					}
//
//					want := strings.Split(s, "\x1a")
//					for len(want) < 5 {
//						want = append(want, "")
//					}
//					if !reflect.DeepEqual(all, want) {
//						t.Errorf("reading %q:\nhave %x\nwant %x", s, all, want)
//					}
//				})
//			}
//		}
//	}
//}

func TestStatPagefile(t *testing.T) {
	t.Parallel()

	const path = `c:\pagefile.sys`
	fi, err := os.Stat(path)
	if err == nil {
		if fi.Name() == "" {
			t.Fatalf("Stat(%q).Name() is empty", path)
		}
		t.Logf("Stat(%q).Size() = %v", path, fi.Size())
		return
	}
	if os.IsNotExist(err) {
		t.Skip(`skipping because c:\pagefile.sys is not found`)
	}
	t.Fatal(err)
}

// syscallCommandLineToArgv 调用 syscall.CommandLineToArgv 函数，并将返回的结果转换为 []string 类型。
func syscallCommandLineToArgv(cmd string) ([]string, error) {
	var argc int32
	argv, err := syscall.CommandLineToArgv(&syscall.StringToUTF16(cmd)[0], &argc)
	if err != nil {
		return nil, err
	}
	defer syscall.LocalFree(syscall.Handle(uintptr(unsafe.Pointer(argv))))

	var args []string
	for _, v := range (*argv)[:argc] {
		args = append(args, syscall.UTF16ToString((*v)[:]))
	}
	return args, nil
}

//// compareCommandLineToArgvWithSyscall 确保
//// os.CommandLineToArgv(cmd) 和 syscall.CommandLineToArgv(cmd)
//// 返回相同的结果。
//func compareCommandLineToArgvWithSyscall(t *testing.T, cmd string) {
//	syscallArgs, err := syscallCommandLineToArgv(cmd)
//	if err != nil {
//		t.Fatal(err)
//	}
//	args := os.CommandLineToArgv(cmd)
//	if want, have := fmt.Sprintf("%q", syscallArgs), fmt.Sprintf("%q", args); want != have {
//		t.Errorf("testing os.commandLineToArgv(%q) failed: have %q want %q", cmd, args, syscallArgs)
//		return
//	}
//}

//func TestCmdArgs(t *testing.T) {
//	if testing.Short() {
//		t.Skipf("in short mode; skipping test that builds a binary")
//	}
//	t.Parallel()
//
//	tmpdir := t.TempDir()
//
//	const prog = `
//package main
//
//import (
//	"fmt"
//	"github.com/888go/gosdk/os"
//)
//
//func main() {
//	fmt.Printf("%q", os.Args)
//}
//`
//	src := filepath.Join(tmpdir, "main.go")
//	if err := os.WriteFile(src, []byte(prog), 0666); err != nil {
//		t.Fatal(err)
//	}
//
//	exe := filepath.Join(tmpdir, "main.exe")
//	cmd := testenv.Command(t, testenv.GoToolPath(t), "build", "-o", exe, src)
//	cmd.Dir = tmpdir
//	out, err := cmd.CombinedOutput()
//	if err != nil {
//		t.Fatalf("building main.exe failed: %v\n%s", err, out)
//	}
//
//	var cmds = []string{
//		``,
//		` a b c`,
//		` "`,
//		` ""`,
//		` """`,
//		` "" a`,
//		` "123"`,
//		` \"123\"`,
//		` \"123 456\"`,
//		` \\"`,
//		` \\\"`,
//		` \\\\\"`,
//		` \\\"x`,
//		` """"\""\\\"`,
//		` abc`,
//		` \\\\\""x"""y z`,
//		"\tb\t\"x\ty\"",
//		` "Брад" d e`,
//		// 示例来自 https://learn.microsoft.com/zh-cn/cpp/cpp/main-function-command-line-args
//		` "abc" d e`,
//		` a\\b d"e f"g h`,
//		` a\\\"b c d`,
//		` a\\\\"b c" d e`,
//		// 参考：http://daviddeley.com/autohotkey/parameters/parameters.htm#WINARGV
//		// 从5.4节：示例
//		` CallMeIshmael`,
//		` "Call Me Ishmael"`,
//		` Cal"l Me I"shmael`,
//		` CallMe\"Ishmael`,
//		` "CallMe\"Ishmael"`,
//		` "Call Me Ishmael\\"`,
//		` "CallMe\\\"Ishmael"`,
//		` a\\\b`,
//		` "a\\\b"`,
//		// 从5.5 开始 一些常见任务
//		` "\"Call Me Ishmael\""`,
//		` "C:\TEST A\\"`,
//		` "\"C:\TEST A\\\""`,
//		// 从 5.6 起，微软示例解释
//		` "a b c"  d  e`,
//		` "ab\"c"  "\\"  d`,
//		` a\\\b d"e f"g h`,
//		` a\\\"b c d`,
//		` a\\\\"b c" d e`,
//		// 自 5.7 双重双引号示例（2008年前）
//		` "a b c""`,
//		` """CallMeIshmael"""  b  c`,
//		` """Call Me Ishmael"""`,
//		` """"Call Me Ishmael"" b c`,
//	}
//	for _, cmd := range cmds {
//		compareCommandLineToArgvWithSyscall(t, "test"+cmd)
//		compareCommandLineToArgvWithSyscall(t, `"cmd line"`+cmd)
//		compareCommandLineToArgvWithSyscall(t, exe+cmd)
//
//		// 测试syscall.EscapeArg和os.commandLineToArgv两个函数
//		args := os.CommandLineToArgv(exe + cmd)
//		out, err := testenv.Command(t, args[0], args[1:]...).CombinedOutput()
//		if err != nil {
//			t.Fatalf("running %q failed: %v\n%v", args, err, string(out))
//		}
//		if want, have := fmt.Sprintf("%q", args), string(out); want != have {
//			t.Errorf("wrong output of executing %q: have %q want %q", args, have, want)
//			continue
//		}
//	}
//}

func findOneDriveDir() (string, error) {
	// 根据https://stackoverflow.com/questions/42519624/how-to-determine-location-of-onedrive-on-windows-7-and-8-in-c 的说明
	const onedrivekey = `SOFTWARE\Microsoft\OneDrive`
	k, err := registry.OpenKey(registry.CURRENT_USER, onedrivekey, registry.READ)
	if err != nil {
		return "", fmt.Errorf("OpenKey(%q) failed: %v", onedrivekey, err)
	}
	defer k.Close()

	path, valtype, err := k.GetStringValue("UserFolder")
	if err != nil {
		return "", fmt.Errorf("reading UserFolder failed: %v", err)
	}

	if valtype == registry.EXPAND_SZ {
		expanded, err := registry.ExpandString(path)
		if err != nil {
			return "", fmt.Errorf("expanding UserFolder failed: %v", err)
		}
		path = expanded
	}

	return path, nil
}

// TestOneDrive 验证OneDrive文件夹是一个目录，而不是符号链接。
func TestOneDrive(t *testing.T) {
	t.Parallel()

	dir, err := findOneDriveDir()
	if err != nil {
		t.Skipf("Skipping, because we did not find OneDrive directory: %v", err)
	}
	testDirStats(t, dir)
}

func TestWindowsDevNullFile(t *testing.T) {
	t.Parallel()

	f1, err := os.Open("NUL")
	if err != nil {
		t.Fatal(err)
	}
	defer f1.Close()

	fi1, err := f1.Stat()
	if err != nil {
		t.Fatal(err)
	}

	f2, err := os.Open("nul")
	if err != nil {
		t.Fatal(err)
	}
	defer f2.Close()

	fi2, err := f2.Stat()
	if err != nil {
		t.Fatal(err)
	}

	if !os.SameFile(fi1, fi2) {
		t.Errorf(`"NUL" and "nul" are not the same file`)
	}
}

func TestFileStatNUL(t *testing.T) {
	t.Parallel()

	f, err := os.Open("NUL")
	if err != nil {
		t.Fatal(err)
	}
	fi, err := f.Stat()
	if err != nil {
		t.Fatal(err)
	}
	if got, want := fi.Mode(), os.ModeDevice|os.ModeCharDevice|0666; got != want {
		t.Errorf("Open(%q).Stat().Mode() = %v, want %v", "NUL", got, want)
	}
}

func TestStatNUL(t *testing.T) {
	t.Parallel()

	fi, err := os.Stat("NUL")
	if err != nil {
		t.Fatal(err)
	}
	if got, want := fi.Mode(), os.ModeDevice|os.ModeCharDevice|0666; got != want {
		t.Errorf("Stat(%q).Mode() = %v, want %v", "NUL", got, want)
	}
}

// TestSymlinkCreation 验证在开发者模式激活时，创建符号链接在Windows系统上能够正常工作。
// 此功能自Windows 10（版本1703，内部版本号10.0.14972）起得到支持。
func TestSymlinkCreation(t *testing.T) {
	if !testenv.HasSymlink() && !isWindowsDeveloperModeActive() {
		t.Skip("Windows developer mode is not active")
	}
	t.Parallel()

	temp := t.TempDir()
	dummyFile := filepath.Join(temp, "file")
	if err := os.WriteFile(dummyFile, []byte(""), 0644); err != nil {
		t.Fatal(err)
	}

	linkFile := filepath.Join(temp, "link")
	if err := os.Symlink(dummyFile, linkFile); err != nil {
		t.Fatal(err)
	}
}

// isWindowsDeveloperModeActive 检查Windows 10上是否启用了开发者模式。
// 对于早期版本的Windows，返回false。
// 参考：https://docs.microsoft.com/en-us/windows/uwp/get-started/enable-your-device-for-development
func isWindowsDeveloperModeActive() bool {
	key, err := registry.OpenKey(registry.LOCAL_MACHINE, "SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\AppModelUnlock", registry.READ)
	if err != nil {
		return false
	}

	val, _, err := key.GetIntegerValue("AllowDevelopmentWithoutDevLicense")
	if err != nil {
		return false
	}

	return val != 0
}

// TestRootRelativeDirSymlink 验证对于从驱动器根目录（以 "\" 开头但没有卷名）开始的路径，symlink 是否被创建为正确的类型。
// （参见 https://golang.org/issue/39183#issuecomment-632175728。）
func TestRootRelativeDirSymlink(t *testing.T) {
	testenv.MustHaveSymlink(t)
	t.Parallel()

	temp := t.TempDir()
	dir := filepath.Join(temp, "dir")
	if err := os.Mkdir(dir, 0755); err != nil {
		t.Fatal(err)
	}

	volumeRelDir := strings.TrimPrefix(dir, filepath.VolumeName(dir)) // 保留开头的反斜杠

	link := filepath.Join(temp, "link")
	err := os.Symlink(volumeRelDir, link)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Symlink(%#q, %#q)", volumeRelDir, link)

	f, err := os.Open(link)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	if fi, err := f.Stat(); err != nil {
		t.Fatal(err)
	} else if !fi.IsDir() {
		t.Errorf("Open(%#q).Stat().IsDir() = false; want true", f.Name())
	}
}

// TestWorkingDirectoryRelativeSymlink 验证对于驱动器（如 "C:File.txt"）中相对于当前工作目录的符号链接，能够被正确转换为相应符号链接类型的绝对链接（参照 https://docs.microsoft.com/en-us/windows/win32/fileio/creating-symbolic-links）。
func TestWorkingDirectoryRelativeSymlink(t *testing.T) {
	testenv.MustHaveSymlink(t)

	// 创建一个用于建立软链接的目录。
	temp := t.TempDir()
	if v := filepath.VolumeName(temp); len(v) < 2 || v[1] != ':' {
		t.Skipf("Can't test relative symlinks: t.TempDir() (%#q) does not begin with a drive letter.", temp)
	}

	absDir := filepath.Join(temp, `dir\sub`)
	if err := os.MkdirAll(absDir, 0755); err != nil {
		t.Fatal(err)
	}

	// 改变到临时目录，并构造一个相对于工作目录的符号链接。
	oldwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := os.Chdir(oldwd); err != nil {
			t.Fatal(err)
		}
	}()
	if err := os.Chdir(temp); err != nil {
		t.Fatal(err)
	}
	t.Logf("Chdir(%#q)", temp)

	wdRelDir := filepath.VolumeName(temp) + `dir\sub` // 体积后不跟反斜杠。
	absLink := filepath.Join(temp, "link")
	err = os.Symlink(wdRelDir, absLink)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Symlink(%#q, %#q)", wdRelDir, absLink)

	// 现在切换回原始的工作目录，并验证该符号链接仍指向其原始路径，且正确标记为目录。
	if err := os.Chdir(oldwd); err != nil {
		t.Fatal(err)
	}
	t.Logf("Chdir(%#q)", oldwd)

	resolved, err := os.Readlink(absLink)
	if err != nil {
		t.Errorf("Readlink(%#q): %v", absLink, err)
	} else if resolved != absDir {
		t.Errorf("Readlink(%#q) = %#q; want %#q", absLink, resolved, absDir)
	}

	linkFile, err := os.Open(absLink)
	if err != nil {
		t.Fatal(err)
	}
	defer linkFile.Close()

	linkInfo, err := linkFile.Stat()
	if err != nil {
		t.Fatal(err)
	}
	if !linkInfo.IsDir() {
		t.Errorf("Open(%#q).Stat().IsDir() = false; want true", absLink)
	}

	absInfo, err := os.Stat(absDir)
	if err != nil {
		t.Fatal(err)
	}

	if !os.SameFile(absInfo, linkInfo) {
		t.Errorf("SameFile(Stat(%#q), Open(%#q).Stat()) = false; want true", absDir, absLink)
	}
}

// TestStatOfInvalidName 是针对问题 #24999 的回归测试。
func TestStatOfInvalidName(t *testing.T) {
	t.Parallel()

	_, err := os.Stat("*.go")
	if err == nil {
		t.Fatal(`os.Stat("*.go") unexpectedly succeeded`)
	}
}

// findUnusedDriveLetter 在系统上搜索已挂载的驱动器列表（从 Z: 开始，到 D: 结束），寻找未使用的驱动器字母。它返回找到的驱动器根目录的路径（如 Z:\）或错误。
func findUnusedDriveLetter() (string, error) {
	// 不要使用A:和B:，因为它们被保留给了软盘驱动。
	// 不要使用C:，因为它通常用于主驱动。
	for l := 'Z'; l >= 'D'; l-- {
		p := string(l) + `:\`
		_, err := os.Stat(p)
		if os.IsNotExist(err) {
			return p, nil
		}
	}
	return "", errors.New("Could not find unused drive letter.")
}

func TestRootDirAsTemp(t *testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS") == "1" {
		fmt.Print(os.TempDir())
		os.Exit(0)
	}

	testenv.MustHaveExec(t)
	t.Parallel()

	exe, err := os.Executable()
	if err != nil {
		t.Fatal(err)
	}

	newtmp, err := findUnusedDriveLetter()
	if err != nil {
		t.Skip(err)
	}

	cmd := testenv.Command(t, exe, "-test.run=^TestRootDirAsTemp$")
	cmd.Env = cmd.Environ()
	cmd.Env = append(cmd.Env, "GO_WANT_HELPER_PROCESS=1")
	cmd.Env = append(cmd.Env, "TMP="+newtmp)
	cmd.Env = append(cmd.Env, "TEMP="+newtmp)
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Failed to spawn child process: %v %q", err, string(output))
	}
	if want, have := newtmp, string(output); have != want {
		t.Fatalf("unexpected child process output %q, want %q", have, want)
	}
}

func testReadlink(t *testing.T, path, want string) {
	got, err := os.Readlink(path)
	if err != nil {
		t.Error(err)
		return
	}
	if got != want {
		t.Errorf(`Readlink(%q): got %q, want %q`, path, got, want)
	}
}

func mklink(t *testing.T, link, target string) {
	output, err := testenv.Command(t, "cmd", "/c", "mklink", link, target).CombinedOutput()
	if err != nil {
		t.Fatalf("failed to run mklink %v %v: %v %q", link, target, err, output)
	}
}

func mklinkj(t *testing.T, link, target string) {
	output, err := testenv.Command(t, "cmd", "/c", "mklink", "/J", link, target).CombinedOutput()
	if err != nil {
		t.Fatalf("failed to run mklink %v %v: %v %q", link, target, err, output)
	}
}

func mklinkd(t *testing.T, link, target string) {
	output, err := testenv.Command(t, "cmd", "/c", "mklink", "/D", link, target).CombinedOutput()
	if err != nil {
		t.Fatalf("failed to run mklink %v %v: %v %q", link, target, err, output)
	}
}

func TestWindowsReadlink(t *testing.T) {
	tmpdir, err := os.MkdirTemp("", "TestWindowsReadlink")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpdir)

	// 确保 tmpdir 不是符号链接，否则测试将会失败。
	tmpdir, err = filepath.EvalSymlinks(tmpdir)
	if err != nil {
		t.Fatal(err)
	}
	chdir(t, tmpdir)

	vol := filepath.VolumeName(tmpdir)
	output, err := testenv.Command(t, "cmd", "/c", "mountvol", vol, "/L").CombinedOutput()
	if err != nil {
		t.Fatalf("failed to run mountvol %v /L: %v %q", vol, err, output)
	}
	ntvol := strings.Trim(string(output), " \n\r")

	dir := filepath.Join(tmpdir, "dir")
	err = os.MkdirAll(dir, 0777)
	if err != nil {
		t.Fatal(err)
	}

	absdirjlink := filepath.Join(tmpdir, "absdirjlink")
	mklinkj(t, absdirjlink, dir)
	testReadlink(t, absdirjlink, dir)

	ntdirjlink := filepath.Join(tmpdir, "ntdirjlink")
	mklinkj(t, ntdirjlink, ntvol+absdirjlink[len(filepath.VolumeName(absdirjlink)):])
	testReadlink(t, ntdirjlink, absdirjlink)

	ntdirjlinktolink := filepath.Join(tmpdir, "ntdirjlinktolink")
	mklinkj(t, ntdirjlinktolink, ntvol+absdirjlink[len(filepath.VolumeName(absdirjlink)):])
	testReadlink(t, ntdirjlinktolink, absdirjlink)

	mklinkj(t, "reldirjlink", "dir")
	testReadlink(t, "reldirjlink", dir) // 相对目录连接解析为绝对路径

	// 确保我们有足够的权限来运行mklink命令。
	testenv.MustHaveSymlink(t)

	absdirlink := filepath.Join(tmpdir, "absdirlink")
	mklinkd(t, absdirlink, dir)
	testReadlink(t, absdirlink, dir)

	ntdirlink := filepath.Join(tmpdir, "ntdirlink")
	mklinkd(t, ntdirlink, ntvol+absdirlink[len(filepath.VolumeName(absdirlink)):])
	testReadlink(t, ntdirlink, absdirlink)

	mklinkd(t, "reldirlink", "dir")
	testReadlink(t, "reldirlink", "dir")

	file := filepath.Join(tmpdir, "file")
	err = os.WriteFile(file, []byte(""), 0666)
	if err != nil {
		t.Fatal(err)
	}

	filelink := filepath.Join(tmpdir, "filelink")
	mklink(t, filelink, file)
	testReadlink(t, filelink, file)

	linktofilelink := filepath.Join(tmpdir, "linktofilelink")
	mklink(t, linktofilelink, ntvol+filelink[len(filepath.VolumeName(filelink)):])
	testReadlink(t, linktofilelink, filelink)

	mklink(t, "relfilelink", "file")
	testReadlink(t, "relfilelink", "file")
}

func TestOpenDirTOCTOU(t *testing.T) {
	t.Parallel()

	// 检查打开的目录在句柄关闭之前不能重命名。参见问题52747。
	tmpdir := t.TempDir()
	dir := filepath.Join(tmpdir, "dir")
	if err := os.Mkdir(dir, 0777); err != nil {
		t.Fatal(err)
	}
	f, err := os.Open(dir)
	if err != nil {
		t.Fatal(err)
	}
	newpath := filepath.Join(tmpdir, "dir1")
	err = os.Rename(dir, newpath)
	if err == nil || !errors.Is(err, windows.ERROR_SHARING_VIOLATION) {
		f.Close()
		t.Fatalf("Rename(%q, %q) = %v; want windows.ERROR_SHARING_VIOLATION", dir, newpath, err)
	}
	f.Close()
	err = os.Rename(dir, newpath)
	if err != nil {
		t.Error(err)
	}
}

func TestAppExecLinkStat(t *testing.T) {
	// 我们期望安装在 %LOCALAPPDATA%\Microsoft\WindowsApps 中的可执行文件
	// 是带有 IO_REPARSE_TAG_APPEXECLINK 标签的重分析点。此处我们检查这些重分析点
	// 是否被视为非标准（但可执行）文件，而非损坏的符号链接。
	appdata := os.Getenv("LOCALAPPDATA")
	if appdata == "" {
		t.Skipf("skipping: LOCALAPPDATA not set")
	}

	pythonExeName := "python3.exe"
	pythonPath := filepath.Join(appdata, `Microsoft\WindowsApps`, pythonExeName)

	lfi, err := os.Lstat(pythonPath)
	if err != nil {
		t.Skip("skipping test, because Python 3 is not installed via the Windows App Store on this system; see https://golang.org/issue/42919")
	}

	// APPEXECLINK 重解析点不是符号链接，因此 os.Readlink 应该为其返回非空错误，而 Stat 应该返回与 Lstat 相同的结果。
	linkName, err := os.Readlink(pythonPath)
	if err == nil {
		t.Errorf("os.Readlink(%q) = %q, but expected an error\n(should be an APPEXECLINK reparse point, not a symlink)", pythonPath, linkName)
	}

	sfi, err := os.Stat(pythonPath)
	if err != nil {
		t.Fatalf("Stat %s: %v", pythonPath, err)
	}

	if lfi.Name() != sfi.Name() {
		t.Logf("os.Lstat(%q) = %+v", pythonPath, lfi)
		t.Logf("os.Stat(%q)  = %+v", pythonPath, sfi)
		t.Errorf("files should be same")
	}

	if lfi.Name() != pythonExeName {
		t.Errorf("Stat %s: got %q, but wanted %q", pythonPath, lfi.Name(), pythonExeName)
	}
	if m := lfi.Mode(); m&fs.ModeSymlink != 0 {
		t.Errorf("%q should be a file, not a link (mode=0x%x)", pythonPath, uint32(m))
	}
	if m := lfi.Mode(); m&fs.ModeDir != 0 {
		t.Errorf("%q should be a file, not a directory (mode=0x%x)", pythonPath, uint32(m))
	}
	if m := lfi.Mode(); m&fs.ModeIrregular == 0 {
		// 重定向点不是常规文件，但我们没有更合适的ModeType位来标记它，所以应该将其标记为不规则的。
		t.Errorf("%q should not be a regular file (mode=0x%x)", pythonPath, uint32(m))
	}

	if sfi.Name() != pythonExeName {
		t.Errorf("Stat %s: got %q, but wanted %q", pythonPath, sfi.Name(), pythonExeName)
	}
	if m := sfi.Mode(); m&fs.ModeSymlink != 0 {
		t.Errorf("%q should be a file, not a link (mode=0x%x)", pythonPath, uint32(m))
	}
	if m := sfi.Mode(); m&fs.ModeDir != 0 {
		t.Errorf("%q should be a file, not a directory (mode=0x%x)", pythonPath, uint32(m))
	}
	if m := sfi.Mode(); m&fs.ModeIrregular == 0 {
		// 重定向点不是常规文件，但我们没有更合适的ModeType位来标记它，所以应该将其标记为不规则的。
		t.Errorf("%q should not be a regular file (mode=0x%x)", pythonPath, uint32(m))
	}

	p, err := exec.LookPath(pythonPath)
	if err != nil {
		t.Errorf("exec.LookPath(%q): %v", pythonPath, err)
	}
	if p != pythonPath {
		t.Errorf("exec.LookPath(%q) = %q; want %q", pythonPath, p, pythonPath)
	}
}

//func TestIllformedUTF16FileName(t *testing.T) {
//	dir := t.TempDir()
//	const sep = string(os.PathSeparator)
//	if !strings.HasSuffix(dir, sep) {
//		dir += sep
//	}
//
//	// 这个UTF-16文件名格式不正确，因为它包含没有被高-surrogate（[1:5]）前导的低-surrogates。
//	namew := []uint16{0x2e, 0xdc6d, 0xdc73, 0xdc79, 0xdc73, 0x30, 0x30, 0x30, 0x31, 0}
//
//	// 创建一个文件名包含非配对代理项的文件。
//	// 使用syscall.CreateFile 而非 os.Create，以模拟由非Go程序创建的文件，
//	// 使得文件名未经 syscall.UTF16FromString 处理。
//	dirw := utf16.Encode([]rune(dir))
//	pathw := append(dirw, namew...)
//	fd, err := syscall.CreateFile(&pathw[0], syscall.GENERIC_ALL, 0, nil, syscall.CREATE_NEW, 0, 0)
//	if err != nil {
//		t.Fatal(err)
//	}
//	syscall.CloseHandle(fd)
//
//	name := syscall.UTF16ToString(namew)
//	path := filepath.Join(dir, name)
//	// 验证os.Lstat能够查询文件。
//	fi, err := os.Lstat(path)
//	if err != nil {
//		t.Fatal(err)
//	}
//	if got := fi.Name(); got != name {
//		t.Errorf("got %q, want %q", got, name)
//	}
//	// 验证File.Readdirnames列出的文件。
//	f, err := os.Open(dir)
//	if err != nil {
//		t.Fatal(err)
//	}
//	files, err := f.Readdirnames(0)
//	f.Close()
//	if err != nil {
//		t.Fatal(err)
//	}
//	if !slices.Contains(files, name) {
//		t.Error("file not listed")
//	}
//	// 验证os.RemoveAll是否可以删除目录，以及它是否会挂起。
//	err = os.RemoveAll(dir)
//	if err != nil {
//		t.Error(err)
//	}
//}

func TestUTF16Alloc(t *testing.T) {
	allowsPerRun := func(want int, f func()) {
		t.Helper()
		got := int(testing.AllocsPerRun(5, f))
		if got != want {
			t.Errorf("got %d allocs, want %d", got, want)
		}
	}
	allowsPerRun(1, func() {
		syscall.UTF16ToString([]uint16{'a', 'b', 'c'})
	})
	allowsPerRun(1, func() {
		syscall.UTF16FromString("abc")
	})
}

func TestNewFileInvalid(t *testing.T) {
	t.Parallel()
	if f := os.NewFile(uintptr(syscall.InvalidHandle), "invalid"); f != nil {
		t.Errorf("NewFile(InvalidHandle) got %v want nil", f)
	}
}

func TestReadDirPipe(t *testing.T) {
	dir := `\\.\pipe\`
	fi, err := os.Stat(dir)
	if err != nil || !fi.IsDir() {
		t.Skipf("%s is not a directory", dir)
	}
	_, err = os.ReadDir(dir)
	if err != nil {
		t.Errorf("ReadDir(%q) = %v", dir, err)
	}
}

func TestReadDirNoFileID(t *testing.T) {
	*os.AllowReadDirFileID = false
	defer func() { *os.AllowReadDirFileID = true }()

	dir := t.TempDir()
	pathA := filepath.Join(dir, "a")
	pathB := filepath.Join(dir, "b")
	if err := os.WriteFile(pathA, nil, 0666); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(pathB, nil, 0666); err != nil {
		t.Fatal(err)
	}

	files, err := os.ReadDir(dir)
	if err != nil {
		t.Fatal(err)
	}
	if len(files) != 2 {
		t.Fatalf("ReadDir(%q) = %v; want 2 files", dir, files)
	}

	// 检查 os.SameFile 函数是否适用于由 os.ReadDir 返回的文件。
	f1, err := files[0].Info()
	if err != nil {
		t.Fatal(err)
	}
	f2, err := files[1].Info()
	if err != nil {
		t.Fatal(err)
	}
	if !os.SameFile(f1, f1) {
		t.Errorf("SameFile(%v, %v) = false; want true", f1, f1)
	}
	if !os.SameFile(f2, f2) {
		t.Errorf("SameFile(%v, %v) = false; want true", f2, f2)
	}
	if os.SameFile(f1, f2) {
		t.Errorf("SameFile(%v, %v) = true; want false", f1, f2)
	}

	// 检查os.SameFile在处理os.ReadDir和os.Stat获取的文件时是否能正常工作。
	f1s, err := os.Stat(pathA)
	if err != nil {
		t.Fatal(err)
	}
	f2s, err := os.Stat(pathB)
	if err != nil {
		t.Fatal(err)
	}
	if !os.SameFile(f1, f1s) {
		t.Errorf("SameFile(%v, %v) = false; want true", f1, f1s)
	}
	if !os.SameFile(f2, f2s) {
		t.Errorf("SameFile(%v, %v) = false; want true", f2, f2s)
	}
}