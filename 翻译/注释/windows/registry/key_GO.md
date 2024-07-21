
<原文开始>
// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// 版权所有 ? 2015 The Go 作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。
# <翻译结束>


<原文开始>
// Package registry provides access to the Windows registry.
//
// Here is a simple example, opening a registry key and reading a string value from it.
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
//	fmt.Printf("Windows system root is %q\n", s)
<原文结束>

# <翻译开始>
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
# <翻译结束>


<原文开始>
	// Registry key security and access rights.
	// See https://msdn.microsoft.com/en-us/library/windows/desktop/ms724878.aspx
	// for details.
<原文结束>

# <翻译开始>
	// 注册表键的安全性和访问权限。
	// 详情请参阅 https:	//msdn.microsoft.com/en-us/library/windows/desktop/ms724878.aspx
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
# <翻译结束>


<原文开始>
// Key is a handle to an open Windows registry key.
// Keys can be obtained by calling OpenKey; there are
// also some predefined root keys such as CURRENT_USER.
// Keys can be used directly in the Windows API.
<原文结束>

# <翻译开始>
// Key 是一个指向已打开的 Windows 注册表键的句柄。
// 可通过调用 OpenKey 方法获取 Key；同时存在一些预定义的根键，如 CURRENT_USER。
// Key 可直接用于 Windows API 中。
# <翻译结束>


<原文开始>
	// Windows defines some predefined root keys that are always open.
	// An application can use these keys as entry points to the registry.
	// Normally these keys are used in OpenKey to open new keys,
	// but they can also be used anywhere a Key is required.
<原文结束>

# <翻译开始>
	// Windows 定义了一些始终处于打开状态的预定义根键。
	// 应用程序可以将这些键作为访问注册表的入口点。
	// 通常，这些键用于 OpenKey 中以打开新的键，
	// 但它们也可用于任何需要 Key 的地方。
# <翻译结束>


<原文开始>
// Close closes open key k.
<原文结束>

# <翻译开始>
// Close 关闭已打开的键 k。
# <翻译结束>


<原文开始>
// OpenKey opens a new key with path name relative to key k.
// It accepts any open key, including CURRENT_USER and others,
// and returns the new key and an error.
// The access parameter specifies desired access rights to the
// key to be opened.
<原文结束>

# <翻译开始>
// OpenKey 以相对于键 k 的路径名称打开一个新键。
// 它接受任何已打开的键，包括 CURRENT_USER 等，
// 并返回新键及错误信息。
// 参数 access 指定对即将打开的键所期望的访问权限。
//
// 翻译备注:
// win64系统上运行32位软件,会被自动定位到32位的注册表, 设置access(访问权限)参数可以指定访问.具体参考精易"注册表操作Ex"类
// 如,访问64位注册表 WOW64_64KEY | ALL_ACCESS
// 如,访问32位注册表 WOW64_32KEY | ALL_ACCESS
# <翻译结束>


<原文开始>
// OpenRemoteKey opens a predefined registry key on another
// computer pcname. The key to be opened is specified by k, but
// can only be one of LOCAL_MACHINE, PERFORMANCE_DATA or USERS.
// If pcname is "", OpenRemoteKey returns local computer key.
<原文结束>

# <翻译开始>
// OpenRemoteKey 用于在另一台计算机（pcname）上打开一个预定义的注册表键。待打开的键由参数 k 指定，但其值只能是 LOCAL_MACHINE、PERFORMANCE_DATA 或 USERS 之一。若 pcname 为空字符串（""），则 OpenRemoteKey 返回本地计算机的键。
# <翻译结束>


<原文开始>
// ReadSubKeyNames returns the names of subkeys of key k.
// The parameter n controls the number of returned names,
// analogous to the way os.File.Readdirnames works.
<原文结束>

# <翻译开始>
// ReadSubKeyNames 返回键 k 的所有子键名称。
// 参数 n 用于控制返回的子键名称数量，其作用方式与 os.File.Readdirnames 类似。
# <翻译结束>


<原文开始>
	// RegEnumKeyEx must be called repeatedly and to completion.
	// During this time, this goroutine cannot migrate away from
	// its current thread. See https://golang.org/issue/49320 and
	// https://golang.org/issue/49466.
<原文结束>

# <翻译开始>
	// 必须反复调用并确保完成RegEnumKeyEx。
	// 在此期间，该goroutine不能从其当前线程迁移离开。
	// 参见https:	//golang.org/issue/49320和https:	//golang.org/issue/49466。
# <翻译结束>


<原文开始>
	// Registry key size limit is 255 bytes and described there:
	// https://msdn.microsoft.com/library/windows/desktop/ms724872.aspx
<原文结束>

# <翻译开始>
	// 注册表键大小限制为255字节，具体描述如下：
	// https:	//msdn.microsoft.com/library/windows/desktop/ms724872.aspx
# <翻译结束>


<原文开始>
//plus extra room for terminating zero byte
<原文结束>

# <翻译开始>
// 另外预留用于终止零字节的空间
# <翻译结束>


<原文开始>
// Double buffer size and try again.
<原文结束>

# <翻译开始>
// 双倍缓冲区大小并重新尝试
# <翻译结束>


<原文开始>
// CreateKey creates a key named path under open key k.
// CreateKey returns the new key and a boolean flag that reports
// whether the key already existed.
// The access parameter specifies the access rights for the key
// to be created.
<原文结束>

# <翻译开始>
// CreateKey 在已打开的键 k 下创建一个名为 path 的键。
// CreateKey 返回新创建的键以及一个布尔标志，该标志报告
// 该键是否已存在。
// 参数 access 指定了待创建键的访问权限。
//
// 翻译备注:
// win64系统上运行32位软件,会被自动定位到32位的注册表, 设置access(访问权限)参数可以指定访问.具体参考精易"注册表操作Ex"类
// 如,访问64位注册表 WOW64_64KEY | ALL_ACCESS
// 如,访问32位注册表 WOW64_32KEY | ALL_ACCESS
# <翻译结束>


<原文开始>
// DeleteKey deletes the subkey path of key k and its values.
<原文结束>

# <翻译开始>
// DeleteKey 删除键k的子键路径及其值。
# <翻译结束>


<原文开始>
// A KeyInfo describes the statistics of a key. It is returned by Stat.
<原文结束>

# <翻译开始>
// KeyInfo 描述了一个键的统计信息。它由 Stat 函数返回。
# <翻译结束>


<原文开始>
// ModTime returns the key's last write time.
<原文结束>

# <翻译开始>
// ModTime 返回键的最后写入时间。
# <翻译结束>


<原文开始>
// Stat retrieves information about the open key k.
<原文结束>

# <翻译开始>
// Stat 获取关于已打开键 k 的信息。
# <翻译结束>

