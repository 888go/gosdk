//go:build windows

package registry

import (
	"golang.org/x/sys/windows/registry"
	"runtime"
	"syscall"
	"time"
)

const (
	// ALL_ACCESS 注册表项安全性和'访问权限'。
	// 见 https://msdn.microsoft.com/en-us/library/windows/desktop/ms724878.aspx
	// 详细信息。

	// ALL_ACCESS 合并STANDARD_RIGHTS_REQUIRED、KEY_QUERY_VALUE、KEY_SET_VALUE、KEY_CREATE_SUB_KEY、
	//KEY_ENUMERATE_SUB_KEYS、KEY_NOTIFY和KEY_CREATE_LINK访问权限。
	ALL_ACCESS = 0xf003f

	CREATE_LINK        = 0x00020 //预留给系统使用。
	CREATE_SUB_KEY     = 0x00004 //创建注册表项的子项是必需的。
	ENUMERATE_SUB_KEYS = 0x00008 //枚举注册表项的子项所必需的。
	EXECUTE            = 0x20019 //等效于KEY_READ。
	NOTIFY             = 0x00010 //请求注册表项或注册表项子项的更改通知所必需的。
	QUERY_VALUE        = 0x00001 //查询注册表项的值所必需的。
	READ               = 0x20019 //合并STANDARD_RIGHTS_READ、KEY_QUERY_VALUE、KEY_ENUMERATE_SUB_KEYS和KEY_NOTIFY值。
	SET_VALUE          = 0x00002 //创建、删除或设置注册表值所必需的。

	// WOW64_32KEY 指示 64 位Windows上的应用程序应在 32 位注册表视图中运行。 此标志被 32 位Windows忽略。
	//有关详细信息，请参阅 访问备用注册表视图。
	//必须将此标志与此表中查询或访问注册表值的其他标志结合使用。Windows 2000：不支持此标志
	WOW64_32KEY = 0x00200 //

	// WOW64_64KEY 指示 64 位Windows上的应用程序应在 64 位注册表视图中运行。
	//此标志被 32 位Windows忽略。 有关详细信息，请参阅 访问备用注册表视图。
	//必须将此标志与此表中查询或访问注册表值的其他标志结合使用。
	//Windows 2000：不支持此标志。
	WOW64_64KEY = 0x00100

	WRITE = 0x20006 //合并STANDARD_RIGHTS_WRITE、KEY_SET_VALUE和KEY_CREATE_SUB_KEY访问权限。
)

// Key结构 是打开的Windows注册表项的句柄。
// 可以通过调用OpenKey获取注册表对象; 还有一些预定义的根注册表对象，例如CURRENT_USER。
// 注册表对象可以直接在Windows API中使用。
// type Key结构 syscall.Handle
type Key struct {
	F registry.Key
}

var (
	// CLASSES_ROOT Windows定义了一些始终打开的预定义根注册表对象。
	// 应用程序可以使用这些键作为注册表的入口点。
	// 通常在OpenKey中使用这些键来打开新的键，
	//但它们也可以在需要注册表对象的任何地方使用。
	CLASSES_ROOT     = Key{registry.Key(syscall.HKEY_CLASSES_ROOT)}
	CURRENT_USER     = Key{registry.Key(syscall.HKEY_CURRENT_USER)}
	LOCAL_MACHINE    = Key{registry.Key(syscall.HKEY_LOCAL_MACHINE)}
	USERS            = Key{registry.Key(syscall.HKEY_USERS)}
	CURRENT_CONFIG   = Key{registry.Key(syscall.HKEY_CURRENT_CONFIG)}
	PERFORMANCE_DATA = Key{registry.Key(syscall.HKEY_PERFORMANCE_DATA)}
)

// Close 关闭已打开的密钥 k。

// ff:关闭
func (k Key) Close() error { //md5:101e90b7cfafc7dfd868583867e9a4eb
	return k.F.Close()
}

// OpenKey 以相对于键 k 的路径名称打开一个新键。
// 它接受任何已打开的键，包括 CURRENT_USER 等，
// 并返回新键和错误信息。
// 参数 access 指定了对将要打开的键所期望的访问权限。

// ff:打开表项
// Key:
// access:访问权限
// path:路径
// k:
func OpenKey(k Key, path string, access uint32) (Key, error) { //md5:3a12c169bc95450438a63ebba1d45fb7
	//这里是单独增加的, 防止win64系统运行32位软件, 访问注册表被重定向到32位注册表, 具体参考精易"注册表操作Ex"类
	if access == 0 {
		if runtime.GOARCH == "amd64" {
			access = WOW64_64KEY | ALL_ACCESS //64位注册表 注意:这里的权限 采用的是  #ALL_ACCESS  全部权限
		} else {
			access = WOW64_32KEY | ALL_ACCESS //32位注册表 注意:这里的权限 采用的是  #ALL_ACCESS  全部权限
		}
	}
	new, err := registry.OpenKey(k.F, path, access)
	return Key{new}, err
}

// OpenRemoteKey 用于在另一台计算机（pcname）上打开一个预定义的注册表键。
// 待打开的键由参数 k 指定，但其值只能是 LOCAL_MACHINE、PERFORMANCE_DATA 或 USERS 中的一个。
// 若 pcname 为空字符串（""），则 OpenRemoteKey 将返回本地计算机的键。

// ff:打开远程表项
// Key:
// k:
// pcname:计算机名
func OpenRemoteKey(pcname string, k Key) (Key, error) { //md5:33eca2105a504ff567599c2b43300af9
	new, err := registry.OpenRemoteKey(pcname, k.F)
	return Key{new}, err
}

// ReadSubKeyNames 返回键 k 的子键名称。
// 参数 n 用于控制返回的名称数量，其作用方式与 os.File.Readdirnames 类似。

// ff:取所有子项名称
// n:
func (k Key) ReadSubKeyNames(n int) ([]string, error) { //md5:11a1c7236cd7d72ba501d833bfb6fd8c
	return k.F.ReadSubKeyNames(n)
}

// CreateKey 在已打开的键 k 下创建一个名为 path 的键。
// CreateKey 返回新键以及一个布尔标志，该标志报告
// 该键是否已存在。
// 参数 access 指定了要创建的键的访问权限。

// ff:创建表项
// err:错误
// openedExisting:是否已存在
// newk:
// access:访问权限
// path:路径
// k:
func CreateKey(k Key, path string, access uint32) (newk Key, openedExisting bool, err error) { //md5:1297b365aa944808d337c3215d9de370
	//这里是单独增加的, 防止win64系统运行32位软件, 访问注册表被重定向到32位注册表, 具体参考精易"注册表操作Ex"类
	if access == 0 {
		if runtime.GOARCH == "amd64" {
			access = WOW64_64KEY | ALL_ACCESS //64位注册表 注意:这里的权限 采用的是  #ALL_ACCESS  全部权限
		} else {
			access = WOW64_32KEY | ALL_ACCESS //32位注册表 注意:这里的权限 采用的是  #ALL_ACCESS  全部权限
		}
	}
	new, 是否已存在, err := registry.CreateKey(k.F, path, access)
	return Key{new}, 是否已存在, err
}

// DeleteKey 删除键k的子键路径及其值。

// ff:删除表项
// path:路径
// k:
func DeleteKey(k Key, path string) error { //md5:177479006aec43c6614f80eadef20888
	return registry.DeleteKey(k.F, path)
}

// A I对象信息 描述注册表对象的统计信息。由Stat.返回。
type KeyInfo struct {
	SubKeyCount     uint32
	MaxSubKeyLen    uint32 // 具有最长名称的键的子键的大小，以Unicode字符表示，不包括终止的零字节
	ValueCount      uint32
	MaxValueNameLen uint32 // 键的最长值名称的大小，以Unicode字符表示，不包括终止的零字节
	MaxValueLen     uint32 //键值中最长的数据组件，以字节为单位
	F               registry.KeyInfo
}

// ModTime 返回键的最后写入时间。

// ff:取写入时间
func (ki *KeyInfo) ModTime() time.Time { //md5:1f67aad4834fc70ab710d65b82aa6782
	return ki.F.ModTime()
}

// Stat 获取关于已打开键 k 的信息。

// ff:取对象信息
func (k Key) Stat() (*KeyInfo, error) { //md5:af6f18b48b40be37bba1092a582e79b3
	返回, err := k.F.Stat()

	对象信息 := KeyInfo{
		SubKeyCount:     返回.SubKeyCount,
		MaxSubKeyLen:    返回.MaxSubKeyLen, // 具有最长名称的键的子键的大小，以Unicode字符表示，不包括终止的零字节
		ValueCount:      返回.ValueCount,
		MaxValueNameLen: 返回.MaxValueNameLen, // 键的最长值名称的大小，以Unicode字符表示，不包括终止的零字节
		MaxValueLen:     返回.MaxValueLen,     //键值中最长的数据组件，以字节为单位
		F:               *返回,
	}
	return &对象信息, err
}
