
<原文开始>
// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
# <翻译结束>


<原文开始>
// NTStatus corresponds with NTSTATUS, error values returned by ntdll.dll and
// other native functions.
<原文结束>

# <翻译开始>
// NTStatus corresponds with NTSTATUS, error values returned by ntdll.dll and
// other native functions.
# <翻译结束>


<原文开始>
// Invented values to support what package os expects.
<原文结束>

# <翻译开始>
// Invented values to support what package os expects.
# <翻译结束>


<原文开始>
// More invented values for signals
<原文结束>

# <翻译开始>
// More invented values for signals
# <翻译结束>


<原文开始>
// Access rights for process.
<原文结束>

# <翻译开始>
// Access rights for process.
# <翻译结束>


<原文开始>
// Access rights for thread.
<原文结束>

# <翻译开始>
// Access rights for thread.
# <翻译结束>


<原文开始>
// Windows reserves errors >= 1<<29 for application use.
<原文结束>

# <翻译开始>
// Windows reserves errors >= 1<<29 for application use.
# <翻译结束>


<原文开始>
// Process creation flags.
<原文结束>

# <翻译开始>
// Process creation flags.
# <翻译结束>


<原文开始>
// attributes for ProcThreadAttributeList
<原文结束>

# <翻译开始>
// attributes for ProcThreadAttributeList
# <翻译结束>


<原文开始>
// flags for CreateToolhelp32Snapshot
<原文结束>

# <翻译开始>
// flags for CreateToolhelp32Snapshot
# <翻译结束>


<原文开始>
// flags for EnumProcessModulesEx
<原文结束>

# <翻译开始>
// flags for EnumProcessModulesEx
# <翻译结束>


<原文开始>
// filters for ReadDirectoryChangesW and FindFirstChangeNotificationW
<原文结束>

# <翻译开始>
// filters for ReadDirectoryChangesW and FindFirstChangeNotificationW
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
// Pointer represents a pointer to an arbitrary Windows type.
//
// Pointer-typed fields may point to one of many different types. It's
// up to the caller to provide a pointer to the appropriate type, cast
// to Pointer. The caller must obey the unsafe.Pointer rules while
// doing so.
# <翻译结束>


<原文开始>
// Nanoseconds returns Filetime ft in nanoseconds
// since Epoch (00:00:00 UTC, January 1, 1970).
<原文结束>

# <翻译开始>
// Nanoseconds returns Filetime ft in nanoseconds
// since Epoch (00:00:00 UTC, January 1, 1970).
# <翻译结束>


<原文开始>
// 100-nanosecond intervals since January 1, 1601
<原文结束>

# <翻译开始>
// 100-nanosecond intervals since January 1, 1601
# <翻译结束>


<原文开始>
// change starting time to the Epoch (00:00:00 UTC, January 1, 1970)
<原文结束>

# <翻译开始>
// change starting time to the Epoch (00:00:00 UTC, January 1, 1970)
# <翻译结束>


<原文开始>
// convert into nanoseconds
<原文结束>

# <翻译开始>
// convert into nanoseconds
# <翻译结束>


<原文开始>
// convert into 100-nanosecond
<原文结束>

# <翻译开始>
// convert into 100-nanosecond
# <翻译结束>


<原文开始>
// change starting time to January 1, 1601
<原文结束>

# <翻译开始>
// change starting time to January 1, 1601
# <翻译结束>


<原文开始>
// This is the actual system call structure.
// Win32finddata is what we committed to in Go 1.
<原文结束>

# <翻译开始>
// This is the actual system call structure.
// Win32finddata is what we committed to in Go 1.
# <翻译结束>


<原文开始>
// The src is 1 element bigger than dst, but it must be NUL.
<原文结束>

# <翻译开始>
// The src is 1 element bigger than dst, but it must be NUL.
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
// ProcThreadAttributeList is a placeholder type to represent a PROC_THREAD_ATTRIBUTE_LIST.
//
// To create a *ProcThreadAttributeList, use NewProcThreadAttributeList, update
// it with ProcThreadAttributeListContainer.Update, free its memory using
// ProcThreadAttributeListContainer.Delete, and access the list itself using
// ProcThreadAttributeListContainer.List.
# <翻译结束>


<原文开始>
// cf. http://support.microsoft.com/default.aspx?scid=kb;en-us;257460
<原文结束>

# <翻译开始>
// cf. http://support.microsoft.com/default.aspx?scid=kb;en-us;257460
# <翻译结束>


<原文开始>
// flags inside DNSRecord.Dw
<原文结束>

# <翻译开始>
// flags inside DNSRecord.Dw
# <翻译结束>


<原文开始>
// flags of WSALookupService
<原文结束>

# <翻译开始>
// flags of WSALookupService
# <翻译结束>


<原文开始>
// values of WSAQUERYSET's namespace
<原文结束>

# <翻译开始>
// values of WSAQUERYSET's namespace
# <翻译结束>


<原文开始>
// TODO(mattn): SockaddrGen is union of sockaddr/sockaddr_in/sockaddr_in6_old.
// will be fixed to change variable type as suitable.
<原文结束>

# <翻译开始>
// TODO(mattn): SockaddrGen is union of sockaddr/sockaddr_in/sockaddr_in6_old.
// will be fixed to change variable type as suitable.
# <翻译结束>


<原文开始>
// IP returns an IPv4 or IPv6 address, or nil if the underlying SocketAddress is neither.
<原文结束>

# <翻译开始>
// IP returns an IPv4 or IPv6 address, or nil if the underlying SocketAddress is neither.
# <翻译结束>


<原文开始>
// Console related constants used for the mode parameter to SetConsoleMode. See
// https://docs.microsoft.com/en-us/windows/console/setconsolemode for details.
<原文结束>

# <翻译开始>
// Console related constants used for the mode parameter to SetConsoleMode. See
// https://docs.microsoft.com/en-us/windows/console/setconsolemode for details.
# <翻译结束>


<原文开始>
// Pseudo console related constants used for the flags parameter to
// CreatePseudoConsole. See: https://learn.microsoft.com/en-us/windows/console/createpseudoconsole
<原文结束>

# <翻译开始>
// Pseudo console related constants used for the flags parameter to
// CreatePseudoConsole. See: https://learn.microsoft.com/en-us/windows/console/createpseudoconsole
# <翻译结束>


<原文开始>
// Used with GetConsoleScreenBuffer to retrieve information about a console
// screen buffer. See
// https://docs.microsoft.com/en-us/windows/console/console-screen-buffer-info-str
// for details.
<原文结束>

# <翻译开始>
// Used with GetConsoleScreenBuffer to retrieve information about a console
// screen buffer. See
// https://docs.microsoft.com/en-us/windows/console/console-screen-buffer-info-str
// for details.
# <翻译结束>


<原文开始>
// flags for JOBOBJECT_BASIC_LIMIT_INFORMATION.LimitFlags
<原文结束>

# <翻译开始>
// flags for JOBOBJECT_BASIC_LIMIT_INFORMATION.LimitFlags
# <翻译结束>


<原文开始>
// JobObjectInformationClass for QueryInformationJobObject and SetInformationJobObject
<原文结束>

# <翻译开始>
// JobObjectInformationClass for QueryInformationJobObject and SetInformationJobObject
# <翻译结束>


<原文开始>
// Flags used for GetModuleHandleEx
<原文结束>

# <翻译开始>
// Flags used for GetModuleHandleEx
# <翻译结束>


<原文开始>
// MUI function flag values
<原文结束>

# <翻译开始>
// MUI function flag values
# <翻译结束>


<原文开始>
// FILE_INFO_BY_HANDLE_CLASS constants for SetFileInformationByHandle/GetFileInformationByHandleEx
<原文结束>

# <翻译开始>
// FILE_INFO_BY_HANDLE_CLASS constants for SetFileInformationByHandle/GetFileInformationByHandleEx
# <翻译结束>


<原文开始>
// LoadLibrary flags for determining from where to search for a DLL
<原文结束>

# <翻译开始>
// LoadLibrary flags for determining from where to search for a DLL
# <翻译结束>


<原文开始>
// RegNotifyChangeKeyValue notifyFilter flags.
<原文结束>

# <翻译开始>
// RegNotifyChangeKeyValue notifyFilter flags.
# <翻译结束>


<原文开始>
// REG_NOTIFY_CHANGE_NAME notifies the caller if a subkey is added or deleted.
<原文结束>

# <翻译开始>
// REG_NOTIFY_CHANGE_NAME notifies the caller if a subkey is added or deleted.
# <翻译结束>


<原文开始>
// REG_NOTIFY_CHANGE_ATTRIBUTES notifies the caller of changes to the attributes of the key, such as the security descriptor information.
<原文结束>

# <翻译开始>
// REG_NOTIFY_CHANGE_ATTRIBUTES notifies the caller of changes to the attributes of the key, such as the security descriptor information.
# <翻译结束>


<原文开始>
// REG_NOTIFY_CHANGE_LAST_SET notifies the caller of changes to a value of the key. This can include adding or deleting a value, or changing an existing value.
<原文结束>

# <翻译开始>
// REG_NOTIFY_CHANGE_LAST_SET notifies the caller of changes to a value of the key. This can include adding or deleting a value, or changing an existing value.
# <翻译结束>


<原文开始>
// REG_NOTIFY_CHANGE_SECURITY notifies the caller of changes to the security descriptor of the key.
<原文结束>

# <翻译开始>
// REG_NOTIFY_CHANGE_SECURITY notifies the caller of changes to the security descriptor of the key.
# <翻译结束>


<原文开始>
// REG_NOTIFY_THREAD_AGNOSTIC indicates that the lifetime of the registration must not be tied to the lifetime of the thread issuing the RegNotifyChangeKeyValue call. Note: This flag value is only supported in Windows 8 and later.
<原文结束>

# <翻译开始>
// REG_NOTIFY_THREAD_AGNOSTIC indicates that the lifetime of the registration must not be tied to the lifetime of the thread issuing the RegNotifyChangeKeyValue call. Note: This flag value is only supported in Windows 8 and later.
# <翻译结束>


<原文开始>
// NTUnicodeString is a UTF-16 string for NT native APIs, corresponding to UNICODE_STRING.
<原文结束>

# <翻译开始>
// NTUnicodeString is a UTF-16 string for NT native APIs, corresponding to UNICODE_STRING.
# <翻译结束>


<原文开始>
// NTString is an ANSI string for NT native APIs, corresponding to STRING.
<原文结束>

# <翻译开始>
// NTString is an ANSI string for NT native APIs, corresponding to STRING.
# <翻译结束>


<原文开始>
// Values for the Attributes member of OBJECT_ATTRIBUTES.
<原文结束>

# <翻译开始>
// Values for the Attributes member of OBJECT_ATTRIBUTES.
# <翻译结束>


<原文开始>
// CreateDisposition flags for NtCreateFile and NtCreateNamedPipeFile.
<原文结束>

# <翻译开始>
// CreateDisposition flags for NtCreateFile and NtCreateNamedPipeFile.
# <翻译结束>


<原文开始>
// CreateOptions flags for NtCreateFile and NtCreateNamedPipeFile.
<原文结束>

# <翻译开始>
// CreateOptions flags for NtCreateFile and NtCreateNamedPipeFile.
# <翻译结束>


<原文开始>
// Parameter constants for NtCreateNamedPipeFile.
<原文结束>

# <翻译开始>
// Parameter constants for NtCreateNamedPipeFile.
# <翻译结束>


<原文开始>
// FileInformationClass for NtSetInformationFile
<原文结束>

# <翻译开始>
// FileInformationClass for NtSetInformationFile
# <翻译结束>


<原文开始>
// Flags for FILE_RENAME_INFORMATION
<原文结束>

# <翻译开始>
// Flags for FILE_RENAME_INFORMATION
# <翻译结束>


<原文开始>
// Flags for FILE_DISPOSITION_INFORMATION_EX
<原文结束>

# <翻译开始>
// Flags for FILE_DISPOSITION_INFORMATION_EX
# <翻译结束>


<原文开始>
// Flags for FILE_CASE_SENSITIVE_INFORMATION
<原文结束>

# <翻译开始>
// Flags for FILE_CASE_SENSITIVE_INFORMATION
# <翻译结束>


<原文开始>
// Flags for FILE_LINK_INFORMATION
<原文结束>

# <翻译开始>
// Flags for FILE_LINK_INFORMATION
# <翻译结束>


<原文开始>
// ProcessInformationClasses for NtQueryInformationProcess and NtSetInformationProcess.
<原文结束>

# <翻译开始>
// ProcessInformationClasses for NtQueryInformationProcess and NtSetInformationProcess.
# <翻译结束>


<原文开始>
// SystemInformationClasses for NtQuerySystemInformation and NtSetSystemInformation
<原文结束>

# <翻译开始>
// SystemInformationClasses for NtQuerySystemInformation and NtSetSystemInformation
# <翻译结束>


<原文开始>
// Constants for LocalAlloc flags.
<原文结束>

# <翻译开始>
// Constants for LocalAlloc flags.
# <翻译结束>


<原文开始>
// Constants for the CreateNamedPipe-family of functions.
<原文结束>

# <翻译开始>
// Constants for the CreateNamedPipe-family of functions.
# <翻译结束>


<原文开始>
// Constants for security attributes when opening named pipes.
<原文结束>

# <翻译开始>
// Constants for security attributes when opening named pipes.
# <翻译结束>


<原文开始>
// ResourceID represents a 16-bit resource identifier, traditionally created with the MAKEINTRESOURCE macro.
<原文结束>

# <翻译开始>
// ResourceID represents a 16-bit resource identifier, traditionally created with the MAKEINTRESOURCE macro.
# <翻译结束>


<原文开始>
// ResourceIDOrString must be either a ResourceID, to specify a resource or resource type by ID,
// or a string, to specify a resource or resource type by name.
<原文结束>

# <翻译开始>
// ResourceIDOrString must be either a ResourceID, to specify a resource or resource type by ID,
// or a string, to specify a resource or resource type by name.
# <翻译结束>


<原文开始>
// Predefined resource names and types.
<原文结束>

# <翻译开始>
// Predefined resource names and types.
# <翻译结束>


<原文开始>
// Flag for QueryFullProcessImageName.
<原文结束>

# <翻译开始>
// Flag for QueryFullProcessImageName.
# <翻译结束>

