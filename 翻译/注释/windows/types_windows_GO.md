
<原文开始>
// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// 版权所有 2011 Go 作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可协议约束，
// 该协议可在 LICENSE 文件中找到。
# <翻译结束>


<原文开始>
// NTStatus corresponds with NTSTATUS, error values returned by ntdll.dll and
// other native functions.
<原文结束>

# <翻译开始>
// NTStatus 与 NTSTATUS 相对应，是 ntdll.dll 及其他原生函数返回的错误值。
# <翻译结束>


<原文开始>
// Invented values to support what package os expects.
<原文结束>

# <翻译开始>
// 为满足package os所期望的内容而设定的虚构值
# <翻译结束>


<原文开始>
// More invented values for signals
<原文结束>

# <翻译开始>
// 更多用于信号的虚构值
# <翻译结束>


<原文开始>
// Access rights for process.
<原文结束>

# <翻译开始>
// 进程访问权限
# <翻译结束>


<原文开始>
// Access rights for thread.
<原文结束>

# <翻译开始>
// 线程访问权限
# <翻译结束>


<原文开始>
// Windows reserves errors >= 1<<29 for application use.
<原文结束>

# <翻译开始>
// Windows将错误值预留为>= 1<<29，供应用程序使用。
# <翻译结束>


<原文开始>
// Process creation flags.
<原文结束>

# <翻译开始>
// 进程创建标志
# <翻译结束>


<原文开始>
// attributes for ProcThreadAttributeList
<原文结束>

# <翻译开始>
// ProcThreadAttributeList的属性
# <翻译结束>


<原文开始>
// flags for CreateToolhelp32Snapshot
<原文结束>

# <翻译开始>
// CreateToolhelp32Snapshot函数的标志
# <翻译结束>


<原文开始>
// flags for EnumProcessModulesEx
<原文结束>

# <翻译开始>
// EnumProcessModulesEx的标志
# <翻译结束>


<原文开始>
// filters for ReadDirectoryChangesW and FindFirstChangeNotificationW
<原文结束>

# <翻译开始>
// 用于 ReadDirectoryChangesW 和 FindFirstChangeNotificationW 的过滤器
# <翻译结束>


<原文开始>
// Pointer represents a pointer to an arbitrary Windows type.
//
// Pointer-typed fields may point to one of many different types. It's
// up to the caller to provide a pointer to the appropriate type, cast
// to Pointer. The caller must obey the unsafe.Pointer rules while
// doing so.
<原文结束>

# <翻译开始>
// Pointer 代表指向任意 Windows 类型的指针。
// 
// Pointer 类型的字段可能指向多种不同类型的其中之一。调用者需提供已转换为 Pointer 类型的相应类型指针。
// 调用者在进行此操作时必须遵守 unsafe.Pointer 的规则。
# <翻译结束>


<原文开始>
// Nanoseconds returns Filetime ft in nanoseconds
// since Epoch (00:00:00 UTC, January 1, 1970).
<原文结束>

# <翻译开始>
// Nanoseconds 函数返回 Filetime 结构体 ft 自 Unix 纪元（UTC 时间 1970年1月1日 00:00:00）以来的纳秒数。
# <翻译结束>


<原文开始>
// 100-nanosecond intervals since January 1, 1601
<原文结束>

# <翻译开始>
// 自1601年1月1日起的100纳秒间隔
# <翻译结束>


<原文开始>
// change starting time to the Epoch (00:00:00 UTC, January 1, 1970)
<原文结束>

# <翻译开始>
// 将起始时间更改为纪元（UTC时间1970年1月1日00:00:00）
# <翻译结束>


<原文开始>
// convert into nanoseconds
<原文结束>

# <翻译开始>
// 转换为纳秒
# <翻译结束>


<原文开始>
// convert into 100-nanosecond
<原文结束>

# <翻译开始>
// 转换为100纳秒
# <翻译结束>


<原文开始>
// change starting time to January 1, 1601
<原文结束>

# <翻译开始>
// 将起始时间更改为1601年1月1日
# <翻译结束>


<原文开始>
// This is the actual system call structure.
// Win32finddata is what we committed to in Go 1.
<原文结束>

# <翻译开始>
// 这是实际的系统调用结构。
// Win32finddata 是我们在 Go 1 中承诺采用的。
# <翻译结束>


<原文开始>
// The src is 1 element bigger than dst, but it must be NUL.
<原文结束>

# <翻译开始>
// src比dst大1个元素，但该元素必须为NUL
# <翻译结束>


<原文开始>
// ProcThreadAttributeList is a placeholder type to represent a PROC_THREAD_ATTRIBUTE_LIST.
//
// To create a *ProcThreadAttributeList, use NewProcThreadAttributeList, update
// it with ProcThreadAttributeListContainer.Update, free its memory using
// ProcThreadAttributeListContainer.Delete, and access the list itself using
// ProcThreadAttributeListContainer.List.
<原文结束>

# <翻译开始>
// ProcThreadAttributeList 是一个占位符类型，用于表示一个 PROC_THREAD_ATTRIBUTE_LIST。
//
// 要创建一个 *ProcThreadAttributeList，可使用 NewProcThreadAttributeList，通过 ProcThreadAttributeListContainer.Update 更新它，使用 ProcThreadAttributeListContainer.Delete 释放其内存，并通过 ProcThreadAttributeListContainer.List 访问列表本身。
# <翻译结束>


<原文开始>
// cf. http://support.microsoft.com/default.aspx?scid=kb;en-us;257460
<原文结束>

# <翻译开始>
// 参见: http://support.microsoft.com/default.aspx?scid=kb;en-us;257460
# <翻译结束>


<原文开始>
// flags inside DNSRecord.Dw
<原文结束>

# <翻译开始>
// DNSRecord.Dw内的标志
# <翻译结束>


<原文开始>
// flags of WSALookupService
<原文结束>

# <翻译开始>
// WSALookupService的标志
# <翻译结束>


<原文开始>
// values of WSAQUERYSET's namespace
<原文结束>

# <翻译开始>
// WSAQUERYSET 结构体的命名空间值
# <翻译结束>


<原文开始>
// TODO(mattn): SockaddrGen is union of sockaddr/sockaddr_in/sockaddr_in6_old.
// will be fixed to change variable type as suitable.
<原文结束>

# <翻译开始>
// TODO(mattn): SockaddrGen 是 sockaddr、sockaddr_in 和 sockaddr_in6_old 的联合体。
// 将会被修复以适当地更改变量类型。
# <翻译结束>


<原文开始>
// IP returns an IPv4 or IPv6 address, or nil if the underlying SocketAddress is neither.
<原文结束>

# <翻译开始>
// IP 返回一个 IPv4 或 IPv6 地址，如果基础 SocketAddress 两者都不是，则返回 nil。
# <翻译结束>


<原文开始>
// Console related constants used for the mode parameter to SetConsoleMode. See
// https://docs.microsoft.com/en-us/windows/console/setconsolemode for details.
<原文结束>

# <翻译开始>
// 以下为与控制台相关的常量，用于`SetConsoleMode`函数的`mode`参数。具体信息请参阅
// https://docs.microsoft.com/en-us/windows/console/setconsolemode 。
# <翻译结束>


<原文开始>
// Pseudo console related constants used for the flags parameter to
// CreatePseudoConsole. See: https://learn.microsoft.com/en-us/windows/console/createpseudoconsole
<原文结束>

# <翻译开始>
// 用于作为参数传递给CreatePseudoConsole函数的标志位所涉及的伪控制台相关常量。参见：https://learn.microsoft.com/en-us/windows/console/createpseudoconsole
# <翻译结束>


<原文开始>
// Used with GetConsoleScreenBuffer to retrieve information about a console
// screen buffer. See
// https://docs.microsoft.com/en-us/windows/console/console-screen-buffer-info-str
// for details.
<原文结束>

# <翻译开始>
// 用于与 GetConsoleScreenBuffer 配合使用，以检索有关控制台屏幕缓冲区的信息。
// 详情请参阅 https://docs.microsoft.com/en-us/windows/console/console-screen-buffer-info-str
# <翻译结束>


<原文开始>
// flags for JOBOBJECT_BASIC_LIMIT_INFORMATION.LimitFlags
<原文结束>

# <翻译开始>
// JOBOBJECT_BASIC_LIMIT_INFORMATION结构体中LimitFlags字段的标志
# <翻译结束>


<原文开始>
// JobObjectInformationClass for QueryInformationJobObject and SetInformationJobObject
<原文结束>

# <翻译开始>
// JobObjectInformationClass 用于 QueryInformationJobObject 和 SetInformationJobObject
# <翻译结束>


<原文开始>
// Flags used for GetModuleHandleEx
<原文结束>

# <翻译开始>
// 用于 GetModuleHandleEx 函数的标志
# <翻译结束>


<原文开始>
// MUI function flag values
<原文结束>

# <翻译开始>
// MUI函数标志值
# <翻译结束>


<原文开始>
// FILE_INFO_BY_HANDLE_CLASS constants for SetFileInformationByHandle/GetFileInformationByHandleEx
<原文结束>

# <翻译开始>
// FILE_INFO_BY_HANDLE_CLASS 常量，用于 SetFileInformationByHandle/GetFileInformationByHandleEx 函数
# <翻译结束>


<原文开始>
// LoadLibrary flags for determining from where to search for a DLL
<原文结束>

# <翻译开始>
// LoadLibrary标志，用于确定从何处搜索DLL
# <翻译结束>


<原文开始>
// RegNotifyChangeKeyValue notifyFilter flags.
<原文结束>

# <翻译开始>
// RegNotifyChangeKeyValue 注册表更改通知过滤标志
# <翻译结束>


<原文开始>
// REG_NOTIFY_CHANGE_NAME notifies the caller if a subkey is added or deleted.
<原文结束>

# <翻译开始>
// REG_NOTIFY_CHANGE_NAME 通知调用者如果子键被添加或删除。
# <翻译结束>


<原文开始>
// REG_NOTIFY_CHANGE_ATTRIBUTES notifies the caller of changes to the attributes of the key, such as the security descriptor information.
<原文结束>

# <翻译开始>
// REG_NOTIFY_CHANGE_ATTRIBUTES 通知调用者键属性的更改，例如安全描述符信息。
# <翻译结束>


<原文开始>
// REG_NOTIFY_CHANGE_LAST_SET notifies the caller of changes to a value of the key. This can include adding or deleting a value, or changing an existing value.
<原文结束>

# <翻译开始>
// REG_NOTIFY_CHANGE_LAST_SET 通知调用者键值发生更改。这包括添加或删除某个值，以及更改现有值。
# <翻译结束>


<原文开始>
// REG_NOTIFY_CHANGE_SECURITY notifies the caller of changes to the security descriptor of the key.
<原文结束>

# <翻译开始>
// REG_NOTIFY_CHANGE_SECURITY 通知调用者密钥安全描述符发生更改。
# <翻译结束>


<原文开始>
// REG_NOTIFY_THREAD_AGNOSTIC indicates that the lifetime of the registration must not be tied to the lifetime of the thread issuing the RegNotifyChangeKeyValue call. Note: This flag value is only supported in Windows 8 and later.
<原文结束>

# <翻译开始>
// REG_NOTIFY_THREAD_AGNOSTIC 表示注册的生命周期不得与发出 RegNotifyChangeKeyValue 调用的线程的生命周期相关联。注意：此标志值仅在 Windows 8 及更高版本中受支持。
# <翻译结束>


<原文开始>
// NTUnicodeString is a UTF-16 string for NT native APIs, corresponding to UNICODE_STRING.
<原文结束>

# <翻译开始>
// NTUnicodeString 是一个用于 NT 本地 API 的 UTF-16 字符串，对应于 UNICODE_STRING。
# <翻译结束>


<原文开始>
// NTString is an ANSI string for NT native APIs, corresponding to STRING.
<原文结束>

# <翻译开始>
// NTString 是一个用于 NT 本地 API 的 ANSI 字符串，与 STRING 类型相对应。
# <翻译结束>


<原文开始>
// Values for the Attributes member of OBJECT_ATTRIBUTES.
<原文结束>

# <翻译开始>
// OBJECT_ATTRIBUTES成员的Attributes值。
# <翻译结束>


<原文开始>
// CreateDisposition flags for NtCreateFile and NtCreateNamedPipeFile.
<原文结束>

# <翻译开始>
// CreateDisposition 标志，用于 NtCreateFile 和 NtCreateNamedPipeFile。
# <翻译结束>


<原文开始>
// CreateOptions flags for NtCreateFile and NtCreateNamedPipeFile.
<原文结束>

# <翻译开始>
// CreateOptions 标志，用于 NtCreateFile 和 NtCreateNamedPipeFile。
# <翻译结束>


<原文开始>
// Parameter constants for NtCreateNamedPipeFile.
<原文结束>

# <翻译开始>
// NtCreateNamedPipeFile函数的参数常量
# <翻译结束>


<原文开始>
// FileInformationClass for NtSetInformationFile
<原文结束>

# <翻译开始>
// FileInformationClass 用于 NtSetInformationFile
# <翻译结束>


<原文开始>
// Flags for FILE_RENAME_INFORMATION
<原文结束>

# <翻译开始>
// FILE_RENAME_INFORMATION的标志
# <翻译结束>


<原文开始>
// Flags for FILE_DISPOSITION_INFORMATION_EX
<原文结束>

# <翻译开始>
// FILE_DISPOSITION_INFORMATION_EX的标志
# <翻译结束>


<原文开始>
// Flags for FILE_CASE_SENSITIVE_INFORMATION
<原文结束>

# <翻译开始>
// FILE_CASE_SENSITIVE_INFORMATION标志
# <翻译结束>


<原文开始>
// Flags for FILE_LINK_INFORMATION
<原文结束>

# <翻译开始>
// FILE_LINK_INFORMATION 的标志
# <翻译结束>


<原文开始>
// ProcessInformationClasses for NtQueryInformationProcess and NtSetInformationProcess.
<原文结束>

# <翻译开始>
// ProcessInformationClasses 用于 NtQueryInformationProcess 和 NtSetInformationProcess.
# <翻译结束>


<原文开始>
// SystemInformationClasses for NtQuerySystemInformation and NtSetSystemInformation
<原文结束>

# <翻译开始>
// SystemInformationClasses 用于 NtQuerySystemInformation 和 NtSetSystemInformation
# <翻译结束>


<原文开始>
// Constants for LocalAlloc flags.
<原文结束>

# <翻译开始>
// 用于LocalAlloc标志的常量
# <翻译结束>


<原文开始>
// Constants for the CreateNamedPipe-family of functions.
<原文结束>

# <翻译开始>
// 用于CreateNamedPipe系列函数的常量
# <翻译结束>


<原文开始>
// Constants for security attributes when opening named pipes.
<原文结束>

# <翻译开始>
// 常量，用于命名管道打开时的安全属性
# <翻译结束>


<原文开始>
// ResourceID represents a 16-bit resource identifier, traditionally created with the MAKEINTRESOURCE macro.
<原文结束>

# <翻译开始>
// ResourceID 表示一个16位的资源标识符，传统上通过使用 MAKEINTRESOURCE 宏来创建。
# <翻译结束>


<原文开始>
// ResourceIDOrString must be either a ResourceID, to specify a resource or resource type by ID,
// or a string, to specify a resource or resource type by name.
<原文结束>

# <翻译开始>
// ResourceIDOrString 必须为以下两者之一：
//   - ResourceID，用于通过ID指定资源或资源类型；
//   - 字符串，用于通过名称指定资源或资源类型。
# <翻译结束>


<原文开始>
// Predefined resource names and types.
<原文结束>

# <翻译开始>
// 预定义的资源名称和类型
# <翻译结束>


<原文开始>
// Flag for QueryFullProcessImageName.
<原文结束>

# <翻译开始>
// QueryFullProcessImageName的标志
# <翻译结束>

