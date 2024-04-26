
<原文开始>
// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// 版权所有 2015 The Go 作者。保留所有权利。
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
//
// NOTE: This package is a copy of golang.org/x/sys/windows/registry
// with KeyInfo.ModTime removed to prevent dependency cycles.
<原文结束>

# <翻译开始>
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
# <翻译结束>


<原文开始>
	// Registry key security and access rights.
	// See https://learn.microsoft.com/en-us/windows/win32/sysinfo/registry-key-security-and-access-rights
	// for details.
<原文结束>

# <翻译开始>
// 注册表键的安全性和访问权限。
// 详情请参阅：https://learn.microsoft.com/en-us/windows/win32/sysinfo/registry-key-security-and-access-rights
# <翻译结束>


<原文开始>
// Key is a handle to an open Windows registry key.
// Keys can be obtained by calling OpenKey; there are
// also some predefined root keys such as CURRENT_USER.
// Keys can be used directly in the Windows API.
<原文结束>

# <翻译开始>
// Key 是一个打开的 Windows 注册表键的句柄。键可以通过调用 OpenKey 获取；还有一些预定义的根键，如 CURRENT_USER。
// 键可以直接用于 Windows API。
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
// Close 关闭打开的键 k。
# <翻译结束>


<原文开始>
// OpenKey opens a new key with path name relative to key k.
// It accepts any open key, including CURRENT_USER and others,
// and returns the new key and an error.
// The access parameter specifies desired access rights to the
// key to be opened.
<原文结束>

# <翻译开始>
// OpenKey 打开一个新的密钥，其路径名相对于键 k。它接受任何打开的密钥，包括 CURRENT_USER 等，并返回新的密钥和一个错误。access 参数指定了要打开的密钥所需的访问权限。
# <翻译结束>


<原文开始>
// ReadSubKeyNames returns the names of subkeys of key k.
<原文结束>

# <翻译开始>
// ReadSubKeyNames 返回键k的子键名称。
# <翻译结束>


<原文开始>
	// RegEnumKeyEx must be called repeatedly and to completion.
	// During this time, this goroutine cannot migrate away from
	// its current thread. See #49320.
<原文结束>

# <翻译开始>
// 必须反复调用且直至完成RegEnumKeyEx。
// 在此期间，此goroutine不能从其当前线程迁移离开。参见#49320。
# <翻译结束>


<原文开始>
	// Registry key size limit is 255 bytes and described there:
	// https://learn.microsoft.com/en-us/windows/win32/sysinfo/registry-element-size-limits
<原文结束>

# <翻译开始>
// 注册表键的大小限制为255字节，具体描述如下：
// https://learn.microsoft.com/en-us/windows/win32/sysinfo/registry-element-size-limits
# <翻译结束>


<原文开始>
//plus extra room for terminating zero byte
<原文结束>

# <翻译开始>
//额外预留一个终止零字节的空间
# <翻译结束>


<原文开始>
// Double buffer size and try again.
<原文结束>

# <翻译开始>
// 将缓冲区大小翻倍，然后重试。
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
// 键是否已存在。
// 参数 access 指定要创建的键的访问权限。
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
// KeyInfo描述密钥的统计信息。它是通过Stat返回的。
# <翻译结束>


<原文开始>
// Stat retrieves information about the open key k.
<原文结束>

# <翻译开始>
// Stat获取打开的键k的信息。
# <翻译结束>

