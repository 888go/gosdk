
<原文开始>
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// 版权所有 ? 2009 Go 作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 此许可证可在 LICENSE 文件中找到。
# <翻译结束>


<原文开始>
// Flags for DefineDosDevice.
<原文结束>

# <翻译开始>
// DefineDosDevice 的标志
# <翻译结束>


<原文开始>
// Return values for GetDriveType.
<原文结束>

# <翻译开始>
// GetDriveType函数的返回值
# <翻译结束>


<原文开始>
// File system flags from GetVolumeInformation and GetVolumeInformationByHandle.
<原文结束>

# <翻译开始>
// 文件系统标志，来源于 GetVolumeInformation 和 GetVolumeInformationByHandle 函数。
# <翻译结束>


<原文开始>
// Return value of SleepEx and other APC functions
<原文结束>

# <翻译开始>
// SleepEx及其他APC函数的返回值
# <翻译结束>


<原文开始>
// StringToUTF16 is deprecated. Use UTF16FromString instead.
// If s contains a NUL byte this function panics instead of
// returning an error.
<原文结束>

# <翻译开始>
// StringToUTF16 已被弃用。请改用 UTF16FromString。
// 若 s 中包含 NUL 字节，此函数将触发 panic，而非返回错误。
# <翻译结束>


<原文开始>
// UTF16FromString returns the UTF-16 encoding of the UTF-8 string
// s, with a terminating NUL added. If s contains a NUL byte at any
// location, it returns (nil, syscall.EINVAL).
<原文结束>

# <翻译开始>
// UTF16FromString返回UTF-8字符串s的UTF-16编码，并在其末尾添加终止空字符（NUL）。如果s在任意位置包含空字节（NUL），则返回(nil, syscall.EINVAL)。
# <翻译结束>


<原文开始>
// UTF16ToString returns the UTF-8 encoding of the UTF-16 sequence s,
// with a terminating NUL and any bytes after the NUL removed.
<原文结束>

# <翻译开始>
// UTF16ToString 将UTF-16序列s转换为UTF-8编码形式，
// 并移除终止的空字符（NUL）及该字符之后的所有字节。
# <翻译结束>


<原文开始>
// StringToUTF16Ptr is deprecated. Use UTF16PtrFromString instead.
// If s contains a NUL byte this function panics instead of
// returning an error.
<原文结束>

# <翻译开始>
// StringToUTF16Ptr 已被弃用。请改用 UTF16PtrFromString。
// 若 s 中包含 NUL 字节，此函数将引发 panic，而非返回错误。
# <翻译结束>


<原文开始>
// UTF16PtrFromString returns pointer to the UTF-16 encoding of
// the UTF-8 string s, with a terminating NUL added. If s
// contains a NUL byte at any location, it returns (nil, syscall.EINVAL).
<原文结束>

# <翻译开始>
// UTF16PtrFromString返回对UTF-8字符串s的UTF-16编码的指针，同时在其末尾添加一个终止空字符（NUL）。如果s在任何位置包含空字节（NUL），则返回(nil, syscall.EINVAL)。
# <翻译结束>


<原文开始>
// UTF16PtrToString takes a pointer to a UTF-16 sequence and returns the corresponding UTF-8 encoded string.
// If the pointer is nil, it returns the empty string. It assumes that the UTF-16 sequence is terminated
// at a zero word; if the zero word is not present, the program may crash.
<原文结束>

# <翻译开始>
// UTF16PtrToString 接收一个指向UTF-16序列的指针，并返回对应的UTF-8编码字符串。
// 若该指针为nil，则返回空字符串。此函数假设UTF-16序列以零字词（zero word）结尾；
// 若未出现零字词，程序可能会崩溃。
# <翻译结束>


<原文开始>
// NewCallback converts a Go function to a function pointer conforming to the stdcall calling convention.
// This is useful when interoperating with Windows code requiring callbacks.
// The argument is expected to be a function with one uintptr-sized result. The function must not have arguments with size larger than the size of uintptr.
<原文结束>

# <翻译开始>
// NewCallback 将一个 Go 函数转换为遵循 stdcall 调用约定的函数指针。
// 这在与需要回调的 Windows 代码进行互操作时非常有用。
// 该参数应为具有一个 uintptr 大小结果的函数。该函数不得包含大小大于 uintptr 的参数。
# <翻译结束>


<原文开始>
// NewCallbackCDecl converts a Go function to a function pointer conforming to the cdecl calling convention.
// This is useful when interoperating with Windows code requiring callbacks.
// The argument is expected to be a function with one uintptr-sized result. The function must not have arguments with size larger than the size of uintptr.
<原文结束>

# <翻译开始>
// NewCallbackCDecl 将一个Go函数转换为遵循cdecl调用约定的函数指针。
// 这在与需要回调的Windows代码进行互操作时非常有用。
// 该参数应为具有一个uintptr大小结果的函数。该函数不得包含大小大于uintptr的参数。
# <翻译结束>


<原文开始>
// syscall interface implementation for other packages
<原文结束>

# <翻译开始>
// 为其他包实现的系统调用接口
# <翻译结束>


<原文开始>
// GetCurrentProcess returns the handle for the current process.
// It is a pseudo handle that does not need to be closed.
// The returned error is always nil.
//
// Deprecated: use CurrentProcess for the same Handle without the nil
// error.
<原文结束>

# <翻译开始>
// GetCurrentProcess 返回当前进程的句柄。
// 它是一个伪句柄，无需关闭。
// 返回的错误始终为 nil。
// 
// 已弃用：使用 CurrentProcess 直接获取相同的 Handle，且无 nil 错误。
# <翻译结束>


<原文开始>
// CurrentProcess returns the handle for the current process.
// It is a pseudo handle that does not need to be closed.
<原文结束>

# <翻译开始>
// CurrentProcess 返回当前进程的句柄。
// 它是一个无需关闭的伪句柄。
# <翻译结束>


<原文开始>
// GetCurrentThread returns the handle for the current thread.
// It is a pseudo handle that does not need to be closed.
// The returned error is always nil.
//
// Deprecated: use CurrentThread for the same Handle without the nil
// error.
<原文结束>

# <翻译开始>
// GetCurrentThread 返回当前线程的句柄。
// 它是一个无需关闭的伪句柄。
// 返回的错误始终为 nil。
//
// 已废弃: 请使用 CurrentThread 直接获取相同的 Handle，其不会带有 nil 错误。
# <翻译结束>


<原文开始>
// CurrentThread returns the handle for the current thread.
// It is a pseudo handle that does not need to be closed.
<原文结束>

# <翻译开始>
// CurrentThread 返回当前线程的句柄。
// 它是一个伪句柄，无需关闭。
# <翻译结束>


<原文开始>
// GetProcAddressByOrdinal retrieves the address of the exported
// function from module by ordinal.
<原文结束>

# <翻译开始>
// GetProcAddressByOrdinal 通过序数从模块中获取导出函数的地址
# <翻译结束>


<原文开始>
// NOTE(brainman): work around ERROR_BROKEN_PIPE is returned on reading EOF from stdin
<原文结束>

# <翻译开始>
// 注意(brainman): 当从stdin读取EOF时，为解决返回ERROR_BROKEN_PIPE的问题
# <翻译结束>


<原文开始>
// use GetFileType to check pipe, pipe can't do seek
<原文结束>

# <翻译开始>
// 使用GetFileType检查管道，管道不能进行seek操作
# <翻译结束>


<原文开始>
	// Every other win32 array API takes arguments as "pointer, count", except for this function. So we
	// can't declare it as a usual [] type, because mksyscall will use the opposite order. We therefore
	// trivially stub this ourselves.
<原文结束>

# <翻译开始>
// 除本函数外，所有其他 Win32 数组 API 都采用“指针，计数”作为参数。因此，我们无法将其声明为常规的 [] 类型，因为 mksyscall 会使用相反的顺序。因此，我们自己为此简单地创建了一个存根（stub）。
# <翻译结束>


<原文开始>
// For testing: clients can set this flag to force
// creation of IPv6 sockets to return EAFNOSUPPORT.
<原文结束>

# <翻译开始>
// 用于测试：客户端可以设置此标志，以强制IPv6套接字创建返回EAFNOSUPPORT错误。
# <翻译结束>


<原文开始>
// lowercase; only we can define Sockaddrs
<原文结束>

# <翻译开始>
// 小写形式；只有我们能定义Sockaddrs
# <翻译结束>


<原文开始>
// length is family (uint16), name, NUL.
<原文结束>

# <翻译开始>
// length 包含家族（uint16 类型）、名称以及 NUL 字符
# <翻译结束>


<原文开始>
// Check sl > 3 so we don't change unnamed socket behavior.
<原文结束>

# <翻译开始>
// 检查 sl 是否大于 3，以避免改变无名套接字的行为
# <翻译结束>


<原文开始>
// Don't count trailing NUL for abstract address.
<原文结束>

# <翻译开始>
// 对于抽象地址，不计算尾部的空字符（NUL）
# <翻译结束>


<原文开始>
			// "Abstract" Unix domain socket.
			// Rewrite leading NUL as @ for textual display.
			// (This is the standard convention.)
			// Not friendly to overwrite in place,
			// but the callers below don't care.
<原文结束>

# <翻译开始>
// “抽象”Unix域套接字。
// 将起始空字符（NUL）重写为@以供文本显示。
// （这是标准约定。）
// 不适合在原地进行覆盖，
// 但下面的调用者并不关心这一点。
# <翻译结束>


<原文开始>
		// Assume path ends at NUL.
		// This is not technically the Linux semantics for
		// abstract Unix domain sockets--they are supposed
		// to be uninterpreted fixed-size binary blobs--but
		// everyone uses this convention.
<原文结束>

# <翻译开始>
// 假设路径以 NUL 结尾。
// 这并非严格遵循 Linux 对于抽象 Unix 域套接字的语义——它们本应被视为未经解释的固定大小二进制块——但大家普遍采用这种约定。
# <翻译结束>


<原文开始>
// Invented structures to support what package os expects.
<原文结束>

# <翻译开始>
// 为满足package os所期望的内容，设计了如下结构体
# <翻译结束>


<原文开始>
// Timespec is an invented structure on Windows, but here for
// consistency with the corresponding package for other operating systems.
<原文结束>

# <翻译开始>
// Timespec 在 Windows 上是一个虚构的结构，但此处为了与其他操作系统对应的包保持一致性而存在。
# <翻译结束>


<原文开始>
// TODO(brainman): fix all needed for net
<原文结束>

# <翻译开始>
// TODO(brainman): 修复net所需的全部内容
# <翻译结束>


<原文开始>
// The Linger struct is wrong but we only noticed after Go 1.
// sysLinger is the real system call structure.
<原文结束>

# <翻译开始>
// Linger 结构体是错误的，但我们直到 Go 1 发布后才注意到。
// sysLinger 是实际的系统调用结构体。
# <翻译结束>


<原文开始>
// BUG(brainman): The definition of Linger is not appropriate for direct use
// with Setsockopt and Getsockopt.
// Use SetsockoptLinger instead.
<原文结束>

# <翻译开始>
// BUG(brainman)：Linger 的定义并不适合直接与 Setsockopt 和 Getsockopt 一起使用。
// 请改用 SetsockoptLinger。
# <翻译结束>


<原文开始>
	// EnumProcesses syscall expects the size parameter to be in bytes, but the code generated with mksyscall uses
	// the length of the processIds slice instead. Hence, this wrapper function is added to fix the discrepancy.
<原文结束>

# <翻译开始>
// 函数EnumProcesses期望其size参数以字节为单位，但使用mksyscall生成的代码中，
// 实际使用的是processIds切片的长度。因此，添加此封装函数以修正这一差异。
# <翻译结束>


<原文开始>
	// NOTE(rsc): The Win32finddata struct is wrong for the system call:
	// the two paths are each one uint16 short. Use the correct struct,
	// a win32finddata1, and then copy the results out.
	// There is no loss of expressivity here, because the final
	// uint16, if it is used, is supposed to be a NUL, and Go doesn't need that.
	// For Go 1.1, we might avoid the allocation of win32finddata1 here
	// by adding a final Bug [2]uint16 field to the struct and then
	// adjusting the fields in the result directly.
<原文结束>

# <翻译开始>
// 注意(rsc): 对于系统调用，Win32finddata 结构体是错误的：
// 两个路径各自缺少一个 uint16。使用正确的结构体 win32finddata1，
// 然后将结果复制出来。这里没有表达力的损失，因为如果使用了最后一个 uint16，
// 它应该是空字符（NUL），而 Go 并不需要这个。对于 Go 1.1，
// 我们可能通过在该结构体中添加一个最后的 Bug [2]uint16 字段来避免在此处分配 win32finddata1，
// 然后直接调整结果中的字段。
# <翻译结束>


<原文开始>
// TODO(brainman): fix all needed for os
<原文结束>

# <翻译开始>
// TODO(brainman): 修复os所需的全部内容
# <翻译结束>


<原文开始>
// Readlink returns the destination of the named symbolic link.
<原文结束>

# <翻译开始>
// Readlink返回指定符号链接的目标。
# <翻译结束>


<原文开始>
		// the path is not a symlink or junction but another type of reparse
		// point
<原文结束>

# <翻译开始>
// 路径并非符号链接或junction（硬链接），而是一种其他类型的重解析点
# <翻译结束>


<原文开始>
// GUIDFromString parses a string in the form of
// "{XXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX}" into a GUID.
<原文结束>

# <翻译开始>
// GUIDFromString 将形如
// "{XXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX}" 的字符串解析为一个 GUID。
# <翻译结束>


<原文开始>
// GenerateGUID creates a new random GUID.
<原文结束>

# <翻译开始>
// GenerateGUID 生成一个新的随机 GUID。
# <翻译结束>


<原文开始>
// String returns the canonical string form of the GUID,
// in the form of "{XXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX}".
<原文结束>

# <翻译开始>
// String 方法返回 GUID 的规范字符串形式，
// 其格式为 "{XXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX}"。
# <翻译结束>


<原文开始>
// KnownFolderPath returns a well-known folder path for the current user, specified by one of
// the FOLDERID_ constants, and chosen and optionally created based on a KF_ flag.
<原文结束>

# <翻译开始>
// KnownFolderPath 返回当前用户指定 FOLDERID_ 常量所表示的已知文件夹路径，该路径基于 KF_ 标志进行选择，并可选地根据该标志创建。
# <翻译结束>


<原文开始>
// KnownFolderPath returns a well-known folder path for the user token, specified by one of
// the FOLDERID_ constants, and chosen and optionally created based on a KF_ flag.
<原文结束>

# <翻译开始>
// KnownFolderPath 返回用户令牌中由 FOLDERID_ 常量指定的一个已知文件夹路径，该路径基于 KF_ 标志的选择和可选创建。
# <翻译结束>


<原文开始>
// RtlGetVersion returns the version of the underlying operating system, ignoring
// manifest semantics but is affected by the application compatibility layer.
<原文结束>

# <翻译开始>
// RtlGetVersion 返回底层操作系统版本，忽略清单语义但受应用程序兼容性层影响。
# <翻译结束>


<原文开始>
	// According to documentation, this function always succeeds.
	// The function doesn't even check the validity of the
	// osVersionInfoSize member. Disassembling ntdll.dll indicates
	// that the documentation is indeed correct about that.
<原文结束>

# <翻译开始>
// 根据文档，此函数总是能成功。
// 该函数甚至不检查osVersionInfoSize成员的有效性。
// 对ntdll.dll进行反汇编表明，
// 文档在这一点上的描述确实正确。
# <翻译结束>


<原文开始>
// RtlGetNtVersionNumbers returns the version of the underlying operating system,
// ignoring manifest semantics and the application compatibility layer.
<原文结束>

# <翻译开始>
// RtlGetNtVersionNumbers 返回底层操作系统版本，
// 忽略清单语义及应用程序兼容性层。
# <翻译结束>


<原文开始>
// GetProcessPreferredUILanguages retrieves the process preferred UI languages.
<原文结束>

# <翻译开始>
// 获取进程首选的UI语言
# <翻译结束>


<原文开始>
// GetThreadPreferredUILanguages retrieves the thread preferred UI languages for the current thread.
<原文结束>

# <翻译开始>
// GetThreadPreferredUILanguages 获取当前线程的线程首选 UI 语言。
# <翻译结束>


<原文开始>
// GetUserPreferredUILanguages retrieves information about the user preferred UI languages.
<原文结束>

# <翻译开始>
// GetUserPreferredUILanguages 获取用户首选的UI语言信息
# <翻译结束>


<原文开始>
// GetSystemPreferredUILanguages retrieves the system preferred UI languages.
<原文结束>

# <翻译开始>
// 获取系统首选的UI语言
# <翻译结束>


<原文开始>
// GetProcessPreferredUILanguages may return numLanguages==0 with "\0\0"
<原文结束>

# <翻译开始>
// GetProcessPreferredUILanguages 可能返回 numLanguages==0 且内容为 "\0\0" 的情况
# <翻译结束>


<原文开始>
// remove terminating null
<原文结束>

# <翻译开始>
// 移除终止空字符
# <翻译结束>


<原文开始>
// trim terminating \r and \n
<原文结束>

# <翻译开始>
// 去除结尾的\r和\n
# <翻译结束>


<原文开始>
// NewNTUnicodeString returns a new NTUnicodeString structure for use with native
// NT APIs that work over the NTUnicodeString type. Note that most Windows APIs
// do not use NTUnicodeString, and instead UTF16PtrFromString should be used for
// the more common *uint16 string type.
<原文结束>

# <翻译开始>
// NewNTUnicodeString 用于返回一个新创建的 NTUnicodeString 结构，该结构适用于直接操作 NTUnicodeString 类型的原生 NT API。需要注意的是，大多数 Windows API 并不使用 NTUnicodeString 类型；对于更为常见的 *uint16 字符串类型，应使用 UTF16PtrFromString 函数。
# <翻译结束>


<原文开始>
// Slice returns a uint16 slice that aliases the data in the NTUnicodeString.
<原文结束>

# <翻译开始>
// Slice 返回一个uint16切片，该切片别名引用NTUnicodeString中的数据。
# <翻译结束>


<原文开始>
// NewNTString returns a new NTString structure for use with native
// NT APIs that work over the NTString type. Note that most Windows APIs
// do not use NTString, and instead UTF16PtrFromString should be used for
// the more common *uint16 string type.
<原文结束>

# <翻译开始>
// NewNTString 函数用于创建一个新的 NTString 结构，以便与使用 NTString 类型的原生 NT API 一起工作。请注意，大多数 Windows API 并不使用 NTString，对于更为常见的 *uint16 字符串类型，应使用 UTF16PtrFromString 函数。
# <翻译结束>


<原文开始>
// Slice returns a byte slice that aliases the data in the NTString.
<原文结束>

# <翻译开始>
// Slice 返回一个字节切片，该切片对NTString中的数据进行别名引用。
# <翻译结束>


<原文开始>
// FindResource resolves a resource of the given name and resource type.
<原文结束>

# <翻译开始>
// FindResource 解析给定名称和资源类型的资源。
# <翻译结束>


<原文开始>
// PSAPI_WORKING_SET_EX_BLOCK contains extended working set information for a page.
<原文结束>

# <翻译开始>
// PSAPI_WORKING_SET_EX_BLOCK 结构包含一页的扩展工作集信息。
# <翻译结束>


<原文开始>
// Valid returns the validity of this page.
// If this bit is 1, the subsequent members are valid; otherwise they should be ignored.
<原文结束>

# <翻译开始>
// Valid 返回此页面的有效性。
// 若该位为1，则后续成员有效；否则应忽略它们。
# <翻译结束>


<原文开始>
// ShareCount is the number of processes that share this page. The maximum value of this member is 7.
<原文结束>

# <翻译开始>
// ShareCount 表示共享此页面的进程数量。该成员的最大值为 7。
# <翻译结束>


<原文开始>
// Win32Protection is the memory protection attributes of the page. For a list of values, see
// https://docs.microsoft.com/en-us/windows/win32/memory/memory-protection-constants
<原文结束>

# <翻译开始>
// Win32Protection 表示页面的内存保护属性。有关值列表，请参阅
// https://docs.microsoft.com/zh-cn/windows/win32/memory/memory-protection-constants
# <翻译结束>


<原文开始>
// Shared returns the shared status of this page.
// If this bit is 1, the page can be shared.
<原文结束>

# <翻译开始>
// Shared 返回此页面的共享状态。
// 若该位为1，则表示该页面可被共享。
# <翻译结束>


<原文开始>
// Node is the NUMA node. The maximum value of this member is 63.
<原文结束>

# <翻译开始>
// Node 代表 NUMA（非均匀内存访问）节点。该成员的最大值为 63。
# <翻译结束>


<原文开始>
// Locked returns the locked status of this page.
// If this bit is 1, the virtual page is locked in physical memory.
<原文结束>

# <翻译开始>
// Locked 返回此页面的锁定状态。
// 若该位为1，则表示虚拟页面被锁定在物理内存中。
# <翻译结束>


<原文开始>
// LargePage returns the large page status of this page.
// If this bit is 1, the page is a large page.
<原文结束>

# <翻译开始>
// LargePage 返回该页面的大页状态。
// 若此位为1，则表示该页面为大页。
# <翻译结束>


<原文开始>
// Bad returns the bad status of this page.
// If this bit is 1, the page is has been reported as bad.
<原文结束>

# <翻译开始>
// Bad 返回此页面的错误状态。
// 若该位为1，则表示该页面已被报告为错误状态。
# <翻译结束>


<原文开始>
// intField extracts an integer field in the PSAPI_WORKING_SET_EX_BLOCK union.
<原文结束>

# <翻译开始>
// intField 从 PSAPI_WORKING_SET_EX_BLOCK 联合体中提取一个整型字段
# <翻译结束>


<原文开始>
// PSAPI_WORKING_SET_EX_INFORMATION contains extended working set information for a process.
<原文结束>

# <翻译开始>
// PSAPI_WORKING_SET_EX_INFORMATION 结构包含进程的扩展工作集信息。
# <翻译结束>


<原文开始>
// A PSAPI_WORKING_SET_EX_BLOCK union that indicates the attributes of the page at VirtualAddress.
<原文结束>

# <翻译开始>
// 一个PSAPI_WORKING_SET_EX_BLOCK联合体，用于指示VirtualAddress处页面的属性。
# <翻译结束>


<原文开始>
// CreatePseudoConsole creates a windows pseudo console.
<原文结束>

# <翻译开始>
// 创建Windows伪控制台
# <翻译结束>


<原文开始>
	// We need this wrapper to manually cast Coord to uint32. The autogenerated wrappers only
	// accept arguments that can be casted to uintptr, and Coord can't.
<原文结束>

# <翻译开始>
// 我们需要这个包装器来手动将 Coord 转换为 uint32。自动生成的包装器仅接受可以转换为 uintptr 的参数，而 Coord 无法做到这一点。
# <翻译结束>


<原文开始>
// ResizePseudoConsole resizes the internal buffers of the pseudo console to the width and height specified in `size`.
<原文结束>

# <翻译开始>
// ResizePseudoConsole 将伪控制台的内部缓冲区调整为`size`中指定的宽度和高度。
# <翻译结束>


<原文开始>
// DCB constants. See https://learn.microsoft.com/en-us/windows/win32/api/winbase/ns-winbase-dcb.
<原文结束>

# <翻译开始>
// DCB 常量。参见 https://learn.microsoft.com/en-us/windows/win32/api/winbase/ns-winbase-dcb.
# <翻译结束>


<原文开始>
// EscapeCommFunction constants. See https://learn.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-escapecommfunction.
<原文结束>

# <翻译开始>
// EscapeCommFunction 常量。参见 https://learn.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-escapecommfunction.
# <翻译结束>


<原文开始>
// PurgeComm constants. See https://learn.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-purgecomm.
<原文结束>

# <翻译开始>
// PurgeComm 常量。参见 https://learn.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-purgecomm。
# <翻译结束>


<原文开始>
// SetCommMask constants. See https://learn.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-setcommmask.
<原文结束>

# <翻译开始>
// SetCommMask 常量。参见 https://learn.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-setcommmask.
# <翻译结束>

