
<原文开始>
// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// 版权所有 ? 2015 The Go Authors。保留所有权利。
// 本源代码的使用受 BSD 风格许可协议约束，
// 该协议可在 LICENSE 文件中找到。
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
// package registry 提供对 Windows 注册表的访问功能。
//
// 下面是一个简单的示例，演示如何打开一个注册表键并从中读取字符串值。
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
// 
// 译文：
// 
// package registry 提供对 Windows 注册表的访问接口。
//
// 以下是一个简单的示例，展示如何打开一个注册表键并从中读取一个字符串值。
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
// 详细信息参见 https://msdn.microsoft.com/en-us/library/windows/desktop/ms724878.aspx
# <翻译结束>


<原文开始>
// Key is a handle to an open Windows registry key.
// Keys can be obtained by calling OpenKey; there are
// also some predefined root keys such as CURRENT_USER.
// Keys can be used directly in the Windows API.
<原文结束>

# <翻译开始>
// Key 是对一个已打开的 Windows 注册表键的句柄。
// 可通过调用 OpenKey 获取 Key；同时存在一些预定义的根键，如 CURRENT_USER。
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
// 但也可以在需要 Key 的任何地方使用它们。
# <翻译结束>


<原文开始>
// Close closes open key k.
<原文结束>

# <翻译开始>
// Close 关闭已打开的密钥 k。
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
// 并返回新键和错误信息。
// 参数 access 指定了对将要打开的键所期望的访问权限。
# <翻译结束>


<原文开始>
// OpenRemoteKey opens a predefined registry key on another
// computer pcname. The key to be opened is specified by k, but
// can only be one of LOCAL_MACHINE, PERFORMANCE_DATA or USERS.
// If pcname is "", OpenRemoteKey returns local computer key.
<原文结束>

# <翻译开始>
// OpenRemoteKey 用于在另一台计算机（pcname）上打开一个预定义的注册表键。
// 待打开的键由参数 k 指定，但其值只能是 LOCAL_MACHINE、PERFORMANCE_DATA 或 USERS 中的一个。
// 若 pcname 为空字符串（""），则 OpenRemoteKey 将返回本地计算机的键。
# <翻译结束>


<原文开始>
// ReadSubKeyNames returns the names of subkeys of key k.
// The parameter n controls the number of returned names,
// analogous to the way os.File.Readdirnames works.
<原文结束>

# <翻译开始>
// ReadSubKeyNames 返回键 k 的子键名称。
// 参数 n 用于控制返回的名称数量，其作用方式与 os.File.Readdirnames 类似。
# <翻译结束>


<原文开始>
	// RegEnumKeyEx must be called repeatedly and to completion.
	// During this time, this goroutine cannot migrate away from
	// its current thread. See https://golang.org/issue/49320 and
	// https://golang.org/issue/49466.
<原文结束>

# <翻译开始>
// 必须反复调用 RegEnumKeyEx 直至完成。在此期间，此goroutine不得从其当前线程迁移。参考 https://golang.org/issue/49320 和 https://golang.org/issue/49466.
# <翻译结束>


<原文开始>
	// Registry key size limit is 255 bytes and described there:
	// https://msdn.microsoft.com/library/windows/desktop/ms724872.aspx
<原文结束>

# <翻译开始>
// 注册表键大小限制为255字节，详情参见：
// https://msdn.microsoft.com/library/windows/desktop/ms724872.aspx
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
// 双倍缓冲区大小并尝试再次执行
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
// CreateKey 返回新键以及一个布尔标志，该标志报告
// 该键是否已存在。
// 参数 access 指定了要创建的键的访问权限。
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

