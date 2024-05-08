// 版权所有 2015 The Go 作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

//---build---//go:build windows

// 包 registry 提供对 Windows 注册表的访问。
//
// 以下是一个简单的示例，打开一个注册表键并从中读取字符串值：
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
//	fmt.Printf("Windows 系统根目录是 %q\n", s)
//
// 注意：此包是 golang.org/x/sys/windows/registry 的副本，
// 删除了 KeyInfo.ModTime 以防止依赖循环。
package registry

import (
	"runtime"
	"syscall"
)

const (
// 注册表键的安全性和访问权限。
// 详情请参阅：https://learn.microsoft.com/en-us/windows/win32/sysinfo/registry-key-security-and-access-rights
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

// Key 是一个打开的 Windows 注册表键的句柄。键可以通过调用 OpenKey 获取；还有一些预定义的根键，如 CURRENT_USER。
// 键可以直接用于 Windows API。
type Key syscall.Handle

const (
// Windows 定义了一些始终处于打开状态的预定义根键。
// 应用程序可以将这些键作为访问注册表的入口点。
// 通常，这些键用于 OpenKey 中以打开新的键，
// 但也可以在需要 Key 的任何地方使用它们。
	CLASSES_ROOT   = Key(syscall.HKEY_CLASSES_ROOT)
	CURRENT_USER   = Key(syscall.HKEY_CURRENT_USER)
	LOCAL_MACHINE  = Key(syscall.HKEY_LOCAL_MACHINE)
	USERS          = Key(syscall.HKEY_USERS)
	CURRENT_CONFIG = Key(syscall.HKEY_CURRENT_CONFIG)
)

// Close 关闭打开的键 k。
func (k Key) Close() error {
	return syscall.RegCloseKey(syscall.Handle(k))
}

// OpenKey 打开一个新的密钥，其路径名相对于键 k。它接受任何打开的密钥，包括 CURRENT_USER 等，并返回新的密钥和一个错误。access 参数指定了要打开的密钥所需的访问权限。
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

// ReadSubKeyNames 返回键k的子键名称。
func (k Key) ReadSubKeyNames() ([]string, error) {
// 必须反复调用且直至完成RegEnumKeyEx。
// 在此期间，此goroutine不能从其当前线程迁移离开。参见#49320。
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	names := make([]string, 0)
// 注册表键的大小限制为255字节，具体描述如下：
// https://learn.microsoft.com/en-us/windows/win32/sysinfo/registry-element-size-limits
	buf := make([]uint16, 256) //额外预留一个终止零字节的空间
loopItems:
	for i := uint32(0); ; i++ {
		l := uint32(len(buf))
		for {
			err := syscall.RegEnumKeyEx(syscall.Handle(k), i, &buf[0], &l, nil, nil, nil, nil)
			if err == nil {
				break
			}
			if err == syscall.ERROR_MORE_DATA {
				// 将缓冲区大小翻倍，然后重试。
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
	return names, nil
}

// CreateKey 在已打开的键 k 下创建一个名为 path 的键。
// CreateKey 返回新创建的键以及一个布尔标志，该标志报告
// 键是否已存在。
// 参数 access 指定要创建的键的访问权限。
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
func DeleteKey(k Key, path string) error {
	return regDeleteKey(syscall.Handle(k), syscall.StringToUTF16Ptr(path))
}

// KeyInfo描述密钥的统计信息。它是通过Stat返回的。
type KeyInfo struct {
	SubKeyCount     uint32
	MaxSubKeyLen    uint32 // size of the key's subkey with the longest name, in Unicode characters, not including the terminating zero byte
	ValueCount      uint32
	MaxValueNameLen uint32 // size of the key's longest value name, in Unicode characters, not including the terminating zero byte
	MaxValueLen     uint32 // longest data component among the key's values, in bytes
	lastWriteTime   syscall.Filetime
}

// Stat获取打开的键k的信息。
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
