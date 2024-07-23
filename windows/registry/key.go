// 版权所有 ? 2015 The Go 作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

//go:build windows

// Package registry 提供对 Windows 注册表的访问。
//
// 以下是一个简单示例，打开一个注册表键并从中读取字符串值。
//
//	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows NT\CurrentVersion`, registry.QUERY_VALUE)
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer k.Close()
//
//	s, _, err := k.GetStringValue("SystemRoot")
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Printf("Windows 系统根目录为 %q\n", s)
package registry//bm:win注册表类

import (
	"io"
	"runtime"
	"syscall"
	"time"
)

const (
// 注册表键的安全性和访问权限。
// 详情请参阅 https://msdn.microsoft.com/en-us/library/windows/desktop/ms724878.aspx
/*
CREATE_LINK : 预留给系统使用。
CREATE_SUB_KEY : 创建注册表项的子项是必需的。
ENUMERATE_SUB_KEYS : 枚举注册表项的子项所必需的。
EXECUTE : 等效于KEY_READ。
NOTIFY : 请求注册表项或注册表项子项的更改通知所必需的。
QUERY_VALUE : 查询注册表项的值所必需的。
READ : 合并STANDARD_RIGHTS_READ、KEY_QUERY_VALUE、KEY_ENUMERATE_SUB_KEYS和KEY_NOTIFY值。
SET_VALUE : 创建、删除或设置注册表值所必需的。
WOW64_32KEY : 指示64位Windows上的应用程序应在32位注册表视图中运行。 此标志被 32 位Windows忽略。必须将此标志与此表中查询或访问注册表值的其他标志结合使用。Windows 2000：不支持此标志
使用方法参考  WOW64_32KEY | ALL_ACCESS    
WOW64_64KEY : 指示64位Windows上的应用程序应在64位注册表视图中运行。此标志被 32 位Windows忽略。 必须将此标志与此表中查询或访问注册表值的其他标志结合使用。Windows 2000：不支持此标志。
使用方法参考  WOW64_64KEY | ALL_ACCESS

WRITE : 合并STANDARD_RIGHTS_WRITE、KEY_SET_VALUE和KEY_CREATE_SUB_KEY访问权限。

翻译备注:
win64系统上运行32位软件,会被自动定位到32位的注册表, 设置access(访问权限)参数可以指定访问.具体参考精易"注册表操作Ex"类
如,访问64位注册表 WOW64_64KEY | ALL_ACCESS
如,访问32位注册表 WOW64_32KEY | ALL_ACCESS    
*/
	ALL_ACCESS         = 0xf003f
	CREATE_LINK        = 0x00020
	CREATE_SUB_KEY     = 0x00004
	ENUMERATE_SUB_KEYS = 0x00008
	EXECUTE            = 0x20019
	NOTIFY             = 0x00010
	QUERY_VALUE        = 0x00001
	READ               = 0x20019
	SET_VALUE          = 0x00002
	WOW64_32KEY        = 0x00200
	WOW64_64KEY        = 0x00100
	WRITE              = 0x20006
)

// Key 是一个指向已打开的 Windows 注册表键的句柄。
// 可通过调用 OpenKey 方法获取 Key；同时存在一些预定义的根键，如 CURRENT_USER。
// Key 可直接用于 Windows API 中。
type Key syscall.Handle

const (
// Windows 定义了一些始终处于打开状态的预定义根键。
// 应用程序可以将这些键作为访问注册表的入口点。
// 通常，这些键用于 OpenKey 中以打开新的键，
// 但它们也可用于任何需要 Key 的地方。
	CLASSES_ROOT     = Key(syscall.HKEY_CLASSES_ROOT)
	CURRENT_USER     = Key(syscall.HKEY_CURRENT_USER)
	LOCAL_MACHINE    = Key(syscall.HKEY_LOCAL_MACHINE)
	USERS            = Key(syscall.HKEY_USERS)
	CURRENT_CONFIG   = Key(syscall.HKEY_CURRENT_CONFIG)
	PERFORMANCE_DATA = Key(syscall.HKEY_PERFORMANCE_DATA)
)

// Close 关闭已打开的键 k。
// ff:关闭
// k:
func (k Key) Close() error {
	return syscall.RegCloseKey(syscall.Handle(k))
}

// OpenKey 以相对于键 k 的路径名称打开一个新键。
// 它接受任何已打开的键，包括 CURRENT_USER 等，
// 并返回新键及错误信息。
// 参数 access 指定对即将打开的键所期望的访问权限。
//
// 翻译备注:
// win64系统上运行32位软件,会被自动定位到32位的注册表, 设置access(访问权限)参数可以指定访问.具体参考精易"注册表操作Ex"类
// 如,访问64位注册表 WOW64_64KEY | ALL_ACCESS
// 如,访问32位注册表 WOW64_32KEY | ALL_ACCESS
// ff:打开表项
// k:
// path:路径
// access:访问权限
// Key:
func OpenKey(k Key, path string, access uint32) (Key, error) {
	p, err := syscall.UTF16PtrFromString(path)
	if err != nil {
		return 0, err
	}
	var subkey syscall.Handle
	err = syscall.RegOpenKeyEx(syscall.Handle(k), p, 0, access, &subkey)
	if err != nil {
		return 0, err
	}
	return Key(subkey), nil
}

// OpenRemoteKey 用于在另一台计算机（pcname）上打开一个预定义的注册表键。待打开的键由参数 k 指定，但其值只能是 LOCAL_MACHINE、PERFORMANCE_DATA 或 USERS 之一。若 pcname 为空字符串（""），则 OpenRemoteKey 返回本地计算机的键。
// ff:打开远程表项
// pcname:计算机名
// k:
// Key:
func OpenRemoteKey(pcname string, k Key) (Key, error) {
	var err error
	var p *uint16
	if pcname != "" {
		p, err = syscall.UTF16PtrFromString(`\\` + pcname)
		if err != nil {
			return 0, err
		}
	}
	var remoteKey syscall.Handle
	err = regConnectRegistry(p, syscall.Handle(k), &remoteKey)
	if err != nil {
		return 0, err
	}
	return Key(remoteKey), nil
}

// ReadSubKeyNames 返回键 k 的所有子键名称。
// 参数 n 用于控制返回的子键名称数量，其作用方式与 os.File.Readdirnames 类似。
// ff:取所有子项名称
// k:
// n:
func (k Key) ReadSubKeyNames(n int) ([]string, error) {
// 必须反复调用并确保完成RegEnumKeyEx。
// 在此期间，该goroutine不能从其当前线程迁移离开。
// 参见https://golang.org/issue/49320和https://golang.org/issue/49466。
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	names := make([]string, 0)
// 注册表键大小限制为255字节，具体描述如下：
	buf := make([]uint16, 256) // 另外预留用于终止零字节的空间
loopItems:
	for i := uint32(0); ; i++ {
		if n > 0 {
			if len(names) == n {
				return names, nil
			}
		}
		l := uint32(len(buf))
		for {
			err := syscall.RegEnumKeyEx(syscall.Handle(k), i, &buf[0], &l, nil, nil, nil, nil)
			if err == nil {
				break
			}
			if err == syscall.ERROR_MORE_DATA {
				// 双倍缓冲区大小并重新尝试
				l = uint32(2 * len(buf))
				buf = make([]uint16, l)
				continue
			}
			if err == _ERROR_NO_MORE_ITEMS {
				break loopItems
			}
			return names, err
		}
		names = append(names, syscall.UTF16ToString(buf[:l]))
	}
	if n > len(names) {
		return names, io.EOF
	}
	return names, nil
}

// CreateKey 在已打开的键 k 下创建一个名为 path 的键。
// CreateKey 返回新创建的键以及一个布尔标志，该标志报告
// 该键是否已存在。
// 参数 access 指定了待创建键的访问权限。
//
// 翻译备注:
// win64系统上运行32位软件,会被自动定位到32位的注册表, 设置access(访问权限)参数可以指定访问.具体参考精易"注册表操作Ex"类
// 如,访问64位注册表 WOW64_64KEY | ALL_ACCESS
// 如,访问32位注册表 WOW64_32KEY | ALL_ACCESS
// ff:创建表项
// k:
// path:路径
// access:访问权限
// newk:
// openedExisting:是否已存在
// err:错误
func CreateKey(k Key, path string, access uint32) (newk Key, openedExisting bool, err error) {
	var h syscall.Handle
	var d uint32
	err = regCreateKeyEx(syscall.Handle(k), syscall.StringToUTF16Ptr(path),
		0, nil, _REG_OPTION_NON_VOLATILE, access, nil, &h, &d)
	if err != nil {
		return 0, false, err
	}
	return Key(h), d == _REG_OPENED_EXISTING_KEY, nil
}

// DeleteKey 删除键k的子键路径及其值。
// ff:删除表项
// k:
// path:路径
func DeleteKey(k Key, path string) error {
	return regDeleteKey(syscall.Handle(k), syscall.StringToUTF16Ptr(path))
}

// KeyInfo 描述了一个键的统计信息。它由 Stat 函数返回。
type KeyInfo struct {
	SubKeyCount     uint32
	MaxSubKeyLen    uint32 // size of the key's subkey with the longest name, in Unicode characters, not including the terminating zero byte
	ValueCount      uint32
	MaxValueNameLen uint32 // size of the key's longest value name, in Unicode characters, not including the terminating zero byte
	MaxValueLen     uint32 // longest data component among the key's values, in bytes
	lastWriteTime   syscall.Filetime
}

// ModTime 返回键的最后写入时间。
// ff:取写入时间
// ki:
func (ki *KeyInfo) ModTime() time.Time {
	return time.Unix(0, ki.lastWriteTime.Nanoseconds())
}

// Stat 获取关于已打开键 k 的信息。
// ff:取对象信息
// k:
func (k Key) Stat() (*KeyInfo, error) {
	var ki KeyInfo
	err := syscall.RegQueryInfoKey(syscall.Handle(k), nil, nil, nil,
		&ki.SubKeyCount, &ki.MaxSubKeyLen, nil, &ki.ValueCount,
		&ki.MaxValueNameLen, &ki.MaxValueLen, nil, &ki.lastWriteTime)
	if err != nil {
		return nil, err
	}
	return &ki, nil
}
