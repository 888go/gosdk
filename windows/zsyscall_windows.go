// 代码由'go generate'命令生成；请勿编辑。

package windows

import (
	"syscall"
	"unsafe"
)

var _ unsafe.Pointer

// 对于常见的Errno值，仅一次性完成接口分配。
const (
	errnoERROR_IO_PENDING = 997
)

// 翻译提示:var  (
// 	err异步操作正在进行  error  =  syscall.Errno(errno异步操作正在进行)
// 	err无效参数          error  =  syscall.EINVAL
// )
var (
	errERROR_IO_PENDING error = syscall.Errno(errnoERROR_IO_PENDING)
	errERROR_EINVAL     error = syscall.EINVAL
)

// errnoErr 函数返回常见的已封装 Errno 值，旨在避免运行时的内存分配。
func errnoErr(e syscall.Errno) error {
	switch e {
	case 0:
		return errERROR_EINVAL
	case errnoERROR_IO_PENDING:
		return errERROR_IO_PENDING
	}
// TODO: 在收集到Windows上常见的错误值数据后，此处应添加更多内容。（可能在运行all.bat时进行？）
	return e
}

var (
	modCfgMgr32 = NewLazySystemDLL("CfgMgr32.dll")
	modadvapi32 = NewLazySystemDLL("advapi32.dll")
	modcrypt32  = NewLazySystemDLL("crypt32.dll")
	moddnsapi   = NewLazySystemDLL("dnsapi.dll")
	moddwmapi   = NewLazySystemDLL("dwmapi.dll")
	modiphlpapi = NewLazySystemDLL("iphlpapi.dll")
	modkernel32 = NewLazySystemDLL("kernel32.dll")
	modmswsock  = NewLazySystemDLL("mswsock.dll")
	modnetapi32 = NewLazySystemDLL("netapi32.dll")
	modntdll    = NewLazySystemDLL("ntdll.dll")
	modole32    = NewLazySystemDLL("ole32.dll")
	modpsapi    = NewLazySystemDLL("psapi.dll")
	modsechost  = NewLazySystemDLL("sechost.dll")
	modsecur32  = NewLazySystemDLL("secur32.dll")
	modsetupapi = NewLazySystemDLL("setupapi.dll")
	modshell32  = NewLazySystemDLL("shell32.dll")
	moduser32   = NewLazySystemDLL("user32.dll")
	moduserenv  = NewLazySystemDLL("userenv.dll")
	modversion  = NewLazySystemDLL("version.dll")
	modwinmm    = NewLazySystemDLL("winmm.dll")
	modwintrust = NewLazySystemDLL("wintrust.dll")
	modws2_32   = NewLazySystemDLL("ws2_32.dll")
	modwtsapi32 = NewLazySystemDLL("wtsapi32.dll")

	procCM_Get_DevNode_Status                                = modCfgMgr32.NewProc("CM_Get_DevNode_Status")
	procCM_Get_Device_Interface_ListW                        = modCfgMgr32.NewProc("CM_Get_Device_Interface_ListW")
	procCM_Get_Device_Interface_List_SizeW                   = modCfgMgr32.NewProc("CM_Get_Device_Interface_List_SizeW")
	procCM_MapCrToWin32Err                                   = modCfgMgr32.NewProc("CM_MapCrToWin32Err")
	procAdjustTokenGroups                                    = modadvapi32.NewProc("AdjustTokenGroups")
	procAdjustTokenPrivileges                                = modadvapi32.NewProc("AdjustTokenPrivileges")
	procAllocateAndInitializeSid                             = modadvapi32.NewProc("AllocateAndInitializeSid")
	procBuildSecurityDescriptorW                             = modadvapi32.NewProc("BuildSecurityDescriptorW")
	procChangeServiceConfig2W                                = modadvapi32.NewProc("ChangeServiceConfig2W")
	procChangeServiceConfigW                                 = modadvapi32.NewProc("ChangeServiceConfigW")
	procCheckTokenMembership                                 = modadvapi32.NewProc("CheckTokenMembership")
	procCloseServiceHandle                                   = modadvapi32.NewProc("CloseServiceHandle")
	procControlService                                       = modadvapi32.NewProc("ControlService")
	procConvertSecurityDescriptorToStringSecurityDescriptorW = modadvapi32.NewProc("ConvertSecurityDescriptorToStringSecurityDescriptorW")
	procConvertSidToStringSidW                               = modadvapi32.NewProc("ConvertSidToStringSidW")
	procConvertStringSecurityDescriptorToSecurityDescriptorW = modadvapi32.NewProc("ConvertStringSecurityDescriptorToSecurityDescriptorW")
	procConvertStringSidToSidW                               = modadvapi32.NewProc("ConvertStringSidToSidW")
	procCopySid                                              = modadvapi32.NewProc("CopySid")
	procCreateProcessAsUserW                                 = modadvapi32.NewProc("CreateProcessAsUserW")
	procCreateServiceW                                       = modadvapi32.NewProc("CreateServiceW")
	procCreateWellKnownSid                                   = modadvapi32.NewProc("CreateWellKnownSid")
	procCryptAcquireContextW                                 = modadvapi32.NewProc("CryptAcquireContextW")
	procCryptGenRandom                                       = modadvapi32.NewProc("CryptGenRandom")
	procCryptReleaseContext                                  = modadvapi32.NewProc("CryptReleaseContext")
	procDeleteService                                        = modadvapi32.NewProc("DeleteService")
	procDeregisterEventSource                                = modadvapi32.NewProc("DeregisterEventSource")
	procDuplicateTokenEx                                     = modadvapi32.NewProc("DuplicateTokenEx")
	procEnumDependentServicesW                               = modadvapi32.NewProc("EnumDependentServicesW")
	procEnumServicesStatusExW                                = modadvapi32.NewProc("EnumServicesStatusExW")
	procEqualSid                                             = modadvapi32.NewProc("EqualSid")
	procFreeSid                                              = modadvapi32.NewProc("FreeSid")
	procGetLengthSid                                         = modadvapi32.NewProc("GetLengthSid")
	procGetNamedSecurityInfoW                                = modadvapi32.NewProc("GetNamedSecurityInfoW")
	procGetSecurityDescriptorControl                         = modadvapi32.NewProc("GetSecurityDescriptorControl")
	procGetSecurityDescriptorDacl                            = modadvapi32.NewProc("GetSecurityDescriptorDacl")
	procGetSecurityDescriptorGroup                           = modadvapi32.NewProc("GetSecurityDescriptorGroup")
	procGetSecurityDescriptorLength                          = modadvapi32.NewProc("GetSecurityDescriptorLength")
	procGetSecurityDescriptorOwner                           = modadvapi32.NewProc("GetSecurityDescriptorOwner")
	procGetSecurityDescriptorRMControl                       = modadvapi32.NewProc("GetSecurityDescriptorRMControl")
	procGetSecurityDescriptorSacl                            = modadvapi32.NewProc("GetSecurityDescriptorSacl")
	procGetSecurityInfo                                      = modadvapi32.NewProc("GetSecurityInfo")
	procGetSidIdentifierAuthority                            = modadvapi32.NewProc("GetSidIdentifierAuthority")
	procGetSidSubAuthority                                   = modadvapi32.NewProc("GetSidSubAuthority")
	procGetSidSubAuthorityCount                              = modadvapi32.NewProc("GetSidSubAuthorityCount")
	procGetTokenInformation                                  = modadvapi32.NewProc("GetTokenInformation")
	procImpersonateSelf                                      = modadvapi32.NewProc("ImpersonateSelf")
	procInitializeSecurityDescriptor                         = modadvapi32.NewProc("InitializeSecurityDescriptor")
	procInitiateSystemShutdownExW                            = modadvapi32.NewProc("InitiateSystemShutdownExW")
	procIsTokenRestricted                                    = modadvapi32.NewProc("IsTokenRestricted")
	procIsValidSecurityDescriptor                            = modadvapi32.NewProc("IsValidSecurityDescriptor")
	procIsValidSid                                           = modadvapi32.NewProc("IsValidSid")
	procIsWellKnownSid                                       = modadvapi32.NewProc("IsWellKnownSid")
	procLookupAccountNameW                                   = modadvapi32.NewProc("LookupAccountNameW")
	procLookupAccountSidW                                    = modadvapi32.NewProc("LookupAccountSidW")
	procLookupPrivilegeValueW                                = modadvapi32.NewProc("LookupPrivilegeValueW")
	procMakeAbsoluteSD                                       = modadvapi32.NewProc("MakeAbsoluteSD")
	procMakeSelfRelativeSD                                   = modadvapi32.NewProc("MakeSelfRelativeSD")
	procNotifyServiceStatusChangeW                           = modadvapi32.NewProc("NotifyServiceStatusChangeW")
	procOpenProcessToken                                     = modadvapi32.NewProc("OpenProcessToken")
	procOpenSCManagerW                                       = modadvapi32.NewProc("OpenSCManagerW")
	procOpenServiceW                                         = modadvapi32.NewProc("OpenServiceW")
	procOpenThreadToken                                      = modadvapi32.NewProc("OpenThreadToken")
	procQueryServiceConfig2W                                 = modadvapi32.NewProc("QueryServiceConfig2W")
	procQueryServiceConfigW                                  = modadvapi32.NewProc("QueryServiceConfigW")
	procQueryServiceDynamicInformation                       = modadvapi32.NewProc("QueryServiceDynamicInformation")
	procQueryServiceLockStatusW                              = modadvapi32.NewProc("QueryServiceLockStatusW")
	procQueryServiceStatus                                   = modadvapi32.NewProc("QueryServiceStatus")
	procQueryServiceStatusEx                                 = modadvapi32.NewProc("QueryServiceStatusEx")
	procRegCloseKey                                          = modadvapi32.NewProc("RegCloseKey")
	procRegEnumKeyExW                                        = modadvapi32.NewProc("RegEnumKeyExW")
	procRegNotifyChangeKeyValue                              = modadvapi32.NewProc("RegNotifyChangeKeyValue")
	procRegOpenKeyExW                                        = modadvapi32.NewProc("RegOpenKeyExW")
	procRegQueryInfoKeyW                                     = modadvapi32.NewProc("RegQueryInfoKeyW")
	procRegQueryValueExW                                     = modadvapi32.NewProc("RegQueryValueExW")
	procRegisterEventSourceW                                 = modadvapi32.NewProc("RegisterEventSourceW")
	procRegisterServiceCtrlHandlerExW                        = modadvapi32.NewProc("RegisterServiceCtrlHandlerExW")
	procReportEventW                                         = modadvapi32.NewProc("ReportEventW")
	procRevertToSelf                                         = modadvapi32.NewProc("RevertToSelf")
	procSetEntriesInAclW                                     = modadvapi32.NewProc("SetEntriesInAclW")
	procSetKernelObjectSecurity                              = modadvapi32.NewProc("SetKernelObjectSecurity")
	procSetNamedSecurityInfoW                                = modadvapi32.NewProc("SetNamedSecurityInfoW")
	procSetSecurityDescriptorControl                         = modadvapi32.NewProc("SetSecurityDescriptorControl")
	procSetSecurityDescriptorDacl                            = modadvapi32.NewProc("SetSecurityDescriptorDacl")
	procSetSecurityDescriptorGroup                           = modadvapi32.NewProc("SetSecurityDescriptorGroup")
	procSetSecurityDescriptorOwner                           = modadvapi32.NewProc("SetSecurityDescriptorOwner")
	procSetSecurityDescriptorRMControl                       = modadvapi32.NewProc("SetSecurityDescriptorRMControl")
	procSetSecurityDescriptorSacl                            = modadvapi32.NewProc("SetSecurityDescriptorSacl")
	procSetSecurityInfo                                      = modadvapi32.NewProc("SetSecurityInfo")
	procSetServiceStatus                                     = modadvapi32.NewProc("SetServiceStatus")
	procSetThreadToken                                       = modadvapi32.NewProc("SetThreadToken")
	procSetTokenInformation                                  = modadvapi32.NewProc("SetTokenInformation")
	procStartServiceCtrlDispatcherW                          = modadvapi32.NewProc("StartServiceCtrlDispatcherW")
	procStartServiceW                                        = modadvapi32.NewProc("StartServiceW")
	procCertAddCertificateContextToStore                     = modcrypt32.NewProc("CertAddCertificateContextToStore")
	procCertCloseStore                                       = modcrypt32.NewProc("CertCloseStore")
	procCertCreateCertificateContext                         = modcrypt32.NewProc("CertCreateCertificateContext")
	procCertDeleteCertificateFromStore                       = modcrypt32.NewProc("CertDeleteCertificateFromStore")
	procCertDuplicateCertificateContext                      = modcrypt32.NewProc("CertDuplicateCertificateContext")
	procCertEnumCertificatesInStore                          = modcrypt32.NewProc("CertEnumCertificatesInStore")
	procCertFindCertificateInStore                           = modcrypt32.NewProc("CertFindCertificateInStore")
	procCertFindChainInStore                                 = modcrypt32.NewProc("CertFindChainInStore")
	procCertFindExtension                                    = modcrypt32.NewProc("CertFindExtension")
	procCertFreeCertificateChain                             = modcrypt32.NewProc("CertFreeCertificateChain")
	procCertFreeCertificateContext                           = modcrypt32.NewProc("CertFreeCertificateContext")
	procCertGetCertificateChain                              = modcrypt32.NewProc("CertGetCertificateChain")
	procCertGetNameStringW                                   = modcrypt32.NewProc("CertGetNameStringW")
	procCertOpenStore                                        = modcrypt32.NewProc("CertOpenStore")
	procCertOpenSystemStoreW                                 = modcrypt32.NewProc("CertOpenSystemStoreW")
	procCertVerifyCertificateChainPolicy                     = modcrypt32.NewProc("CertVerifyCertificateChainPolicy")
	procCryptAcquireCertificatePrivateKey                    = modcrypt32.NewProc("CryptAcquireCertificatePrivateKey")
	procCryptDecodeObject                                    = modcrypt32.NewProc("CryptDecodeObject")
	procCryptProtectData                                     = modcrypt32.NewProc("CryptProtectData")
	procCryptQueryObject                                     = modcrypt32.NewProc("CryptQueryObject")
	procCryptUnprotectData                                   = modcrypt32.NewProc("CryptUnprotectData")
	procPFXImportCertStore                                   = modcrypt32.NewProc("PFXImportCertStore")
	procDnsNameCompare_W                                     = moddnsapi.NewProc("DnsNameCompare_W")
	procDnsQuery_W                                           = moddnsapi.NewProc("DnsQuery_W")
	procDnsRecordListFree                                    = moddnsapi.NewProc("DnsRecordListFree")
	procDwmGetWindowAttribute                                = moddwmapi.NewProc("DwmGetWindowAttribute")
	procDwmSetWindowAttribute                                = moddwmapi.NewProc("DwmSetWindowAttribute")
	procGetAdaptersAddresses                                 = modiphlpapi.NewProc("GetAdaptersAddresses")
	procGetAdaptersInfo                                      = modiphlpapi.NewProc("GetAdaptersInfo")
	procGetBestInterfaceEx                                   = modiphlpapi.NewProc("GetBestInterfaceEx")
	procGetIfEntry                                           = modiphlpapi.NewProc("GetIfEntry")
	procAddDllDirectory                                      = modkernel32.NewProc("AddDllDirectory")
	procAssignProcessToJobObject                             = modkernel32.NewProc("AssignProcessToJobObject")
	procCancelIo                                             = modkernel32.NewProc("CancelIo")
	procCancelIoEx                                           = modkernel32.NewProc("CancelIoEx")
	procClearCommBreak                                       = modkernel32.NewProc("ClearCommBreak")
	procClearCommError                                       = modkernel32.NewProc("ClearCommError")
	procCloseHandle                                          = modkernel32.NewProc("CloseHandle")
	procClosePseudoConsole                                   = modkernel32.NewProc("ClosePseudoConsole")
	procConnectNamedPipe                                     = modkernel32.NewProc("ConnectNamedPipe")
	procCreateDirectoryW                                     = modkernel32.NewProc("CreateDirectoryW")
	procCreateEventExW                                       = modkernel32.NewProc("CreateEventExW")
	procCreateEventW                                         = modkernel32.NewProc("CreateEventW")
	procCreateFileMappingW                                   = modkernel32.NewProc("CreateFileMappingW")
	procCreateFileW                                          = modkernel32.NewProc("CreateFileW")
	procCreateHardLinkW                                      = modkernel32.NewProc("CreateHardLinkW")
	procCreateIoCompletionPort                               = modkernel32.NewProc("CreateIoCompletionPort")
	procCreateJobObjectW                                     = modkernel32.NewProc("CreateJobObjectW")
	procCreateMutexExW                                       = modkernel32.NewProc("CreateMutexExW")
	procCreateMutexW                                         = modkernel32.NewProc("CreateMutexW")
	procCreateNamedPipeW                                     = modkernel32.NewProc("CreateNamedPipeW")
	procCreatePipe                                           = modkernel32.NewProc("CreatePipe")
	procCreateProcessW                                       = modkernel32.NewProc("CreateProcessW")
	procCreatePseudoConsole                                  = modkernel32.NewProc("CreatePseudoConsole")
	procCreateSymbolicLinkW                                  = modkernel32.NewProc("CreateSymbolicLinkW")
	procCreateToolhelp32Snapshot                             = modkernel32.NewProc("CreateToolhelp32Snapshot")
	procDefineDosDeviceW                                     = modkernel32.NewProc("DefineDosDeviceW")
	procDeleteFileW                                          = modkernel32.NewProc("DeleteFileW")
	procDeleteProcThreadAttributeList                        = modkernel32.NewProc("DeleteProcThreadAttributeList")
	procDeleteVolumeMountPointW                              = modkernel32.NewProc("DeleteVolumeMountPointW")
	procDeviceIoControl                                      = modkernel32.NewProc("DeviceIoControl")
	procDisconnectNamedPipe                                  = modkernel32.NewProc("DisconnectNamedPipe")
	procDuplicateHandle                                      = modkernel32.NewProc("DuplicateHandle")
	procEscapeCommFunction                                   = modkernel32.NewProc("EscapeCommFunction")
	procExitProcess                                          = modkernel32.NewProc("ExitProcess")
	procExpandEnvironmentStringsW                            = modkernel32.NewProc("ExpandEnvironmentStringsW")
	procFindClose                                            = modkernel32.NewProc("FindClose")
	procFindCloseChangeNotification                          = modkernel32.NewProc("FindCloseChangeNotification")
	procFindFirstChangeNotificationW                         = modkernel32.NewProc("FindFirstChangeNotificationW")
	procFindFirstFileW                                       = modkernel32.NewProc("FindFirstFileW")
	procFindFirstVolumeMountPointW                           = modkernel32.NewProc("FindFirstVolumeMountPointW")
	procFindFirstVolumeW                                     = modkernel32.NewProc("FindFirstVolumeW")
	procFindNextChangeNotification                           = modkernel32.NewProc("FindNextChangeNotification")
	procFindNextFileW                                        = modkernel32.NewProc("FindNextFileW")
	procFindNextVolumeMountPointW                            = modkernel32.NewProc("FindNextVolumeMountPointW")
	procFindNextVolumeW                                      = modkernel32.NewProc("FindNextVolumeW")
	procFindResourceW                                        = modkernel32.NewProc("FindResourceW")
	procFindVolumeClose                                      = modkernel32.NewProc("FindVolumeClose")
	procFindVolumeMountPointClose                            = modkernel32.NewProc("FindVolumeMountPointClose")
	procFlushFileBuffers                                     = modkernel32.NewProc("FlushFileBuffers")
	procFlushViewOfFile                                      = modkernel32.NewProc("FlushViewOfFile")
	procFormatMessageW                                       = modkernel32.NewProc("FormatMessageW")
	procFreeEnvironmentStringsW                              = modkernel32.NewProc("FreeEnvironmentStringsW")
	procFreeLibrary                                          = modkernel32.NewProc("FreeLibrary")
	procGenerateConsoleCtrlEvent                             = modkernel32.NewProc("GenerateConsoleCtrlEvent")
	procGetACP                                               = modkernel32.NewProc("GetACP")
	procGetActiveProcessorCount                              = modkernel32.NewProc("GetActiveProcessorCount")
	procGetCommModemStatus                                   = modkernel32.NewProc("GetCommModemStatus")
	procGetCommState                                         = modkernel32.NewProc("GetCommState")
	procGetCommTimeouts                                      = modkernel32.NewProc("GetCommTimeouts")
	procGetCommandLineW                                      = modkernel32.NewProc("GetCommandLineW")
	procGetComputerNameExW                                   = modkernel32.NewProc("GetComputerNameExW")
	procGetComputerNameW                                     = modkernel32.NewProc("GetComputerNameW")
	procGetConsoleMode                                       = modkernel32.NewProc("GetConsoleMode")
	procGetConsoleScreenBufferInfo                           = modkernel32.NewProc("GetConsoleScreenBufferInfo")
	procGetCurrentDirectoryW                                 = modkernel32.NewProc("GetCurrentDirectoryW")
	procGetCurrentProcessId                                  = modkernel32.NewProc("GetCurrentProcessId")
	procGetCurrentThreadId                                   = modkernel32.NewProc("GetCurrentThreadId")
	procGetDiskFreeSpaceExW                                  = modkernel32.NewProc("GetDiskFreeSpaceExW")
	procGetDriveTypeW                                        = modkernel32.NewProc("GetDriveTypeW")
	procGetEnvironmentStringsW                               = modkernel32.NewProc("GetEnvironmentStringsW")
	procGetEnvironmentVariableW                              = modkernel32.NewProc("GetEnvironmentVariableW")
	procGetExitCodeProcess                                   = modkernel32.NewProc("GetExitCodeProcess")
	procGetFileAttributesExW                                 = modkernel32.NewProc("GetFileAttributesExW")
	procGetFileAttributesW                                   = modkernel32.NewProc("GetFileAttributesW")
	procGetFileInformationByHandle                           = modkernel32.NewProc("GetFileInformationByHandle")
	procGetFileInformationByHandleEx                         = modkernel32.NewProc("GetFileInformationByHandleEx")
	procGetFileTime                                          = modkernel32.NewProc("GetFileTime")
	procGetFileType                                          = modkernel32.NewProc("GetFileType")
	procGetFinalPathNameByHandleW                            = modkernel32.NewProc("GetFinalPathNameByHandleW")
	procGetFullPathNameW                                     = modkernel32.NewProc("GetFullPathNameW")
	procGetLargePageMinimum                                  = modkernel32.NewProc("GetLargePageMinimum")
	procGetLastError                                         = modkernel32.NewProc("GetLastError")
	procGetLogicalDriveStringsW                              = modkernel32.NewProc("GetLogicalDriveStringsW")
	procGetLogicalDrives                                     = modkernel32.NewProc("GetLogicalDrives")
	procGetLongPathNameW                                     = modkernel32.NewProc("GetLongPathNameW")
	procGetMaximumProcessorCount                             = modkernel32.NewProc("GetMaximumProcessorCount")
	procGetModuleFileNameW                                   = modkernel32.NewProc("GetModuleFileNameW")
	procGetModuleHandleExW                                   = modkernel32.NewProc("GetModuleHandleExW")
	procGetNamedPipeHandleStateW                             = modkernel32.NewProc("GetNamedPipeHandleStateW")
	procGetNamedPipeInfo                                     = modkernel32.NewProc("GetNamedPipeInfo")
	procGetOverlappedResult                                  = modkernel32.NewProc("GetOverlappedResult")
	procGetPriorityClass                                     = modkernel32.NewProc("GetPriorityClass")
	procGetProcAddress                                       = modkernel32.NewProc("GetProcAddress")
	procGetProcessId                                         = modkernel32.NewProc("GetProcessId")
	procGetProcessPreferredUILanguages                       = modkernel32.NewProc("GetProcessPreferredUILanguages")
	procGetProcessShutdownParameters                         = modkernel32.NewProc("GetProcessShutdownParameters")
	procGetProcessTimes                                      = modkernel32.NewProc("GetProcessTimes")
	procGetProcessWorkingSetSizeEx                           = modkernel32.NewProc("GetProcessWorkingSetSizeEx")
	procGetQueuedCompletionStatus                            = modkernel32.NewProc("GetQueuedCompletionStatus")
	procGetShortPathNameW                                    = modkernel32.NewProc("GetShortPathNameW")
	procGetStartupInfoW                                      = modkernel32.NewProc("GetStartupInfoW")
	procGetStdHandle                                         = modkernel32.NewProc("GetStdHandle")
	procGetSystemDirectoryW                                  = modkernel32.NewProc("GetSystemDirectoryW")
	procGetSystemPreferredUILanguages                        = modkernel32.NewProc("GetSystemPreferredUILanguages")
	procGetSystemTimeAsFileTime                              = modkernel32.NewProc("GetSystemTimeAsFileTime")
	procGetSystemTimePreciseAsFileTime                       = modkernel32.NewProc("GetSystemTimePreciseAsFileTime")
	procGetSystemWindowsDirectoryW                           = modkernel32.NewProc("GetSystemWindowsDirectoryW")
	procGetTempPathW                                         = modkernel32.NewProc("GetTempPathW")
	procGetThreadPreferredUILanguages                        = modkernel32.NewProc("GetThreadPreferredUILanguages")
	procGetTickCount64                                       = modkernel32.NewProc("GetTickCount64")
	procGetTimeZoneInformation                               = modkernel32.NewProc("GetTimeZoneInformation")
	procGetUserPreferredUILanguages                          = modkernel32.NewProc("GetUserPreferredUILanguages")
	procGetVersion                                           = modkernel32.NewProc("GetVersion")
	procGetVolumeInformationByHandleW                        = modkernel32.NewProc("GetVolumeInformationByHandleW")
	procGetVolumeInformationW                                = modkernel32.NewProc("GetVolumeInformationW")
	procGetVolumeNameForVolumeMountPointW                    = modkernel32.NewProc("GetVolumeNameForVolumeMountPointW")
	procGetVolumePathNameW                                   = modkernel32.NewProc("GetVolumePathNameW")
	procGetVolumePathNamesForVolumeNameW                     = modkernel32.NewProc("GetVolumePathNamesForVolumeNameW")
	procGetWindowsDirectoryW                                 = modkernel32.NewProc("GetWindowsDirectoryW")
	procInitializeProcThreadAttributeList                    = modkernel32.NewProc("InitializeProcThreadAttributeList")
	procIsWow64Process                                       = modkernel32.NewProc("IsWow64Process")
	procIsWow64Process2                                      = modkernel32.NewProc("IsWow64Process2")
	procLoadLibraryExW                                       = modkernel32.NewProc("LoadLibraryExW")
	procLoadLibraryW                                         = modkernel32.NewProc("LoadLibraryW")
	procLoadResource                                         = modkernel32.NewProc("LoadResource")
	procLocalAlloc                                           = modkernel32.NewProc("LocalAlloc")
	procLocalFree                                            = modkernel32.NewProc("LocalFree")
	procLockFileEx                                           = modkernel32.NewProc("LockFileEx")
	procLockResource                                         = modkernel32.NewProc("LockResource")
	procMapViewOfFile                                        = modkernel32.NewProc("MapViewOfFile")
	procModule32FirstW                                       = modkernel32.NewProc("Module32FirstW")
	procModule32NextW                                        = modkernel32.NewProc("Module32NextW")
	procMoveFileExW                                          = modkernel32.NewProc("MoveFileExW")
	procMoveFileW                                            = modkernel32.NewProc("MoveFileW")
	procMultiByteToWideChar                                  = modkernel32.NewProc("MultiByteToWideChar")
	procOpenEventW                                           = modkernel32.NewProc("OpenEventW")
	procOpenMutexW                                           = modkernel32.NewProc("OpenMutexW")
	procOpenProcess                                          = modkernel32.NewProc("OpenProcess")
	procOpenThread                                           = modkernel32.NewProc("OpenThread")
	procPostQueuedCompletionStatus                           = modkernel32.NewProc("PostQueuedCompletionStatus")
	procProcess32FirstW                                      = modkernel32.NewProc("Process32FirstW")
	procProcess32NextW                                       = modkernel32.NewProc("Process32NextW")
	procProcessIdToSessionId                                 = modkernel32.NewProc("ProcessIdToSessionId")
	procPulseEvent                                           = modkernel32.NewProc("PulseEvent")
	procPurgeComm                                            = modkernel32.NewProc("PurgeComm")
	procQueryDosDeviceW                                      = modkernel32.NewProc("QueryDosDeviceW")
	procQueryFullProcessImageNameW                           = modkernel32.NewProc("QueryFullProcessImageNameW")
	procQueryInformationJobObject                            = modkernel32.NewProc("QueryInformationJobObject")
	procReadConsoleW                                         = modkernel32.NewProc("ReadConsoleW")
	procReadDirectoryChangesW                                = modkernel32.NewProc("ReadDirectoryChangesW")
	procReadFile                                             = modkernel32.NewProc("ReadFile")
	procReadProcessMemory                                    = modkernel32.NewProc("ReadProcessMemory")
	procReleaseMutex                                         = modkernel32.NewProc("ReleaseMutex")
	procRemoveDirectoryW                                     = modkernel32.NewProc("RemoveDirectoryW")
	procRemoveDllDirectory                                   = modkernel32.NewProc("RemoveDllDirectory")
	procResetEvent                                           = modkernel32.NewProc("ResetEvent")
	procResizePseudoConsole                                  = modkernel32.NewProc("ResizePseudoConsole")
	procResumeThread                                         = modkernel32.NewProc("ResumeThread")
	procSetCommBreak                                         = modkernel32.NewProc("SetCommBreak")
	procSetCommMask                                          = modkernel32.NewProc("SetCommMask")
	procSetCommState                                         = modkernel32.NewProc("SetCommState")
	procSetCommTimeouts                                      = modkernel32.NewProc("SetCommTimeouts")
	procSetConsoleCursorPosition                             = modkernel32.NewProc("SetConsoleCursorPosition")
	procSetConsoleMode                                       = modkernel32.NewProc("SetConsoleMode")
	procSetCurrentDirectoryW                                 = modkernel32.NewProc("SetCurrentDirectoryW")
	procSetDefaultDllDirectories                             = modkernel32.NewProc("SetDefaultDllDirectories")
	procSetDllDirectoryW                                     = modkernel32.NewProc("SetDllDirectoryW")
	procSetEndOfFile                                         = modkernel32.NewProc("SetEndOfFile")
	procSetEnvironmentVariableW                              = modkernel32.NewProc("SetEnvironmentVariableW")
	procSetErrorMode                                         = modkernel32.NewProc("SetErrorMode")
	procSetEvent                                             = modkernel32.NewProc("SetEvent")
	procSetFileAttributesW                                   = modkernel32.NewProc("SetFileAttributesW")
	procSetFileCompletionNotificationModes                   = modkernel32.NewProc("SetFileCompletionNotificationModes")
	procSetFileInformationByHandle                           = modkernel32.NewProc("SetFileInformationByHandle")
	procSetFilePointer                                       = modkernel32.NewProc("SetFilePointer")
	procSetFileTime                                          = modkernel32.NewProc("SetFileTime")
	procSetFileValidData                                     = modkernel32.NewProc("SetFileValidData")
	procSetHandleInformation                                 = modkernel32.NewProc("SetHandleInformation")
	procSetInformationJobObject                              = modkernel32.NewProc("SetInformationJobObject")
	procSetNamedPipeHandleState                              = modkernel32.NewProc("SetNamedPipeHandleState")
	procSetPriorityClass                                     = modkernel32.NewProc("SetPriorityClass")
	procSetProcessPriorityBoost                              = modkernel32.NewProc("SetProcessPriorityBoost")
	procSetProcessShutdownParameters                         = modkernel32.NewProc("SetProcessShutdownParameters")
	procSetProcessWorkingSetSizeEx                           = modkernel32.NewProc("SetProcessWorkingSetSizeEx")
	procSetStdHandle                                         = modkernel32.NewProc("SetStdHandle")
	procSetVolumeLabelW                                      = modkernel32.NewProc("SetVolumeLabelW")
	procSetVolumeMountPointW                                 = modkernel32.NewProc("SetVolumeMountPointW")
	procSetupComm                                            = modkernel32.NewProc("SetupComm")
	procSizeofResource                                       = modkernel32.NewProc("SizeofResource")
	procSleepEx                                              = modkernel32.NewProc("SleepEx")
	procTerminateJobObject                                   = modkernel32.NewProc("TerminateJobObject")
	procTerminateProcess                                     = modkernel32.NewProc("TerminateProcess")
	procThread32First                                        = modkernel32.NewProc("Thread32First")
	procThread32Next                                         = modkernel32.NewProc("Thread32Next")
	procUnlockFileEx                                         = modkernel32.NewProc("UnlockFileEx")
	procUnmapViewOfFile                                      = modkernel32.NewProc("UnmapViewOfFile")
	procUpdateProcThreadAttribute                            = modkernel32.NewProc("UpdateProcThreadAttribute")
	procVirtualAlloc                                         = modkernel32.NewProc("VirtualAlloc")
	procVirtualFree                                          = modkernel32.NewProc("VirtualFree")
	procVirtualLock                                          = modkernel32.NewProc("VirtualLock")
	procVirtualProtect                                       = modkernel32.NewProc("VirtualProtect")
	procVirtualProtectEx                                     = modkernel32.NewProc("VirtualProtectEx")
	procVirtualQuery                                         = modkernel32.NewProc("VirtualQuery")
	procVirtualQueryEx                                       = modkernel32.NewProc("VirtualQueryEx")
	procVirtualUnlock                                        = modkernel32.NewProc("VirtualUnlock")
	procWTSGetActiveConsoleSessionId                         = modkernel32.NewProc("WTSGetActiveConsoleSessionId")
	procWaitCommEvent                                        = modkernel32.NewProc("WaitCommEvent")
	procWaitForMultipleObjects                               = modkernel32.NewProc("WaitForMultipleObjects")
	procWaitForSingleObject                                  = modkernel32.NewProc("WaitForSingleObject")
	procWriteConsoleW                                        = modkernel32.NewProc("WriteConsoleW")
	procWriteFile                                            = modkernel32.NewProc("WriteFile")
	procWriteProcessMemory                                   = modkernel32.NewProc("WriteProcessMemory")
	procAcceptEx                                             = modmswsock.NewProc("AcceptEx")
	procGetAcceptExSockaddrs                                 = modmswsock.NewProc("GetAcceptExSockaddrs")
	procTransmitFile                                         = modmswsock.NewProc("TransmitFile")
	procNetApiBufferFree                                     = modnetapi32.NewProc("NetApiBufferFree")
	procNetGetJoinInformation                                = modnetapi32.NewProc("NetGetJoinInformation")
	procNetUserGetInfo                                       = modnetapi32.NewProc("NetUserGetInfo")
	procNtCreateFile                                         = modntdll.NewProc("NtCreateFile")
	procNtCreateNamedPipeFile                                = modntdll.NewProc("NtCreateNamedPipeFile")
	procNtQueryInformationProcess                            = modntdll.NewProc("NtQueryInformationProcess")
	procNtQuerySystemInformation                             = modntdll.NewProc("NtQuerySystemInformation")
	procNtSetInformationFile                                 = modntdll.NewProc("NtSetInformationFile")
	procNtSetInformationProcess                              = modntdll.NewProc("NtSetInformationProcess")
	procNtSetSystemInformation                               = modntdll.NewProc("NtSetSystemInformation")
	procRtlAddFunctionTable                                  = modntdll.NewProc("RtlAddFunctionTable")
	procRtlDefaultNpAcl                                      = modntdll.NewProc("RtlDefaultNpAcl")
	procRtlDeleteFunctionTable                               = modntdll.NewProc("RtlDeleteFunctionTable")
	procRtlDosPathNameToNtPathName_U_WithStatus              = modntdll.NewProc("RtlDosPathNameToNtPathName_U_WithStatus")
	procRtlDosPathNameToRelativeNtPathName_U_WithStatus      = modntdll.NewProc("RtlDosPathNameToRelativeNtPathName_U_WithStatus")
	procRtlGetCurrentPeb                                     = modntdll.NewProc("RtlGetCurrentPeb")
	procRtlGetNtVersionNumbers                               = modntdll.NewProc("RtlGetNtVersionNumbers")
	procRtlGetVersion                                        = modntdll.NewProc("RtlGetVersion")
	procRtlInitString                                        = modntdll.NewProc("RtlInitString")
	procRtlInitUnicodeString                                 = modntdll.NewProc("RtlInitUnicodeString")
	procRtlNtStatusToDosErrorNoTeb                           = modntdll.NewProc("RtlNtStatusToDosErrorNoTeb")
	procCLSIDFromString                                      = modole32.NewProc("CLSIDFromString")
	procCoCreateGuid                                         = modole32.NewProc("CoCreateGuid")
	procCoGetObject                                          = modole32.NewProc("CoGetObject")
	procCoInitializeEx                                       = modole32.NewProc("CoInitializeEx")
	procCoTaskMemFree                                        = modole32.NewProc("CoTaskMemFree")
	procCoUninitialize                                       = modole32.NewProc("CoUninitialize")
	procStringFromGUID2                                      = modole32.NewProc("StringFromGUID2")
	procEnumProcessModules                                   = modpsapi.NewProc("EnumProcessModules")
	procEnumProcessModulesEx                                 = modpsapi.NewProc("EnumProcessModulesEx")
	procEnumProcesses                                        = modpsapi.NewProc("EnumProcesses")
	procGetModuleBaseNameW                                   = modpsapi.NewProc("GetModuleBaseNameW")
	procGetModuleFileNameExW                                 = modpsapi.NewProc("GetModuleFileNameExW")
	procGetModuleInformation                                 = modpsapi.NewProc("GetModuleInformation")
	procQueryWorkingSetEx                                    = modpsapi.NewProc("QueryWorkingSetEx")
	procSubscribeServiceChangeNotifications                  = modsechost.NewProc("SubscribeServiceChangeNotifications")
	procUnsubscribeServiceChangeNotifications                = modsechost.NewProc("UnsubscribeServiceChangeNotifications")
	procGetUserNameExW                                       = modsecur32.NewProc("GetUserNameExW")
	procTranslateNameW                                       = modsecur32.NewProc("TranslateNameW")
	procSetupDiBuildDriverInfoList                           = modsetupapi.NewProc("SetupDiBuildDriverInfoList")
	procSetupDiCallClassInstaller                            = modsetupapi.NewProc("SetupDiCallClassInstaller")
	procSetupDiCancelDriverInfoSearch                        = modsetupapi.NewProc("SetupDiCancelDriverInfoSearch")
	procSetupDiClassGuidsFromNameExW                         = modsetupapi.NewProc("SetupDiClassGuidsFromNameExW")
	procSetupDiClassNameFromGuidExW                          = modsetupapi.NewProc("SetupDiClassNameFromGuidExW")
	procSetupDiCreateDeviceInfoListExW                       = modsetupapi.NewProc("SetupDiCreateDeviceInfoListExW")
	procSetupDiCreateDeviceInfoW                             = modsetupapi.NewProc("SetupDiCreateDeviceInfoW")
	procSetupDiDestroyDeviceInfoList                         = modsetupapi.NewProc("SetupDiDestroyDeviceInfoList")
	procSetupDiDestroyDriverInfoList                         = modsetupapi.NewProc("SetupDiDestroyDriverInfoList")
	procSetupDiEnumDeviceInfo                                = modsetupapi.NewProc("SetupDiEnumDeviceInfo")
	procSetupDiEnumDriverInfoW                               = modsetupapi.NewProc("SetupDiEnumDriverInfoW")
	procSetupDiGetClassDevsExW                               = modsetupapi.NewProc("SetupDiGetClassDevsExW")
	procSetupDiGetClassInstallParamsW                        = modsetupapi.NewProc("SetupDiGetClassInstallParamsW")
	procSetupDiGetDeviceInfoListDetailW                      = modsetupapi.NewProc("SetupDiGetDeviceInfoListDetailW")
	procSetupDiGetDeviceInstallParamsW                       = modsetupapi.NewProc("SetupDiGetDeviceInstallParamsW")
	procSetupDiGetDeviceInstanceIdW                          = modsetupapi.NewProc("SetupDiGetDeviceInstanceIdW")
	procSetupDiGetDevicePropertyW                            = modsetupapi.NewProc("SetupDiGetDevicePropertyW")
	procSetupDiGetDeviceRegistryPropertyW                    = modsetupapi.NewProc("SetupDiGetDeviceRegistryPropertyW")
	procSetupDiGetDriverInfoDetailW                          = modsetupapi.NewProc("SetupDiGetDriverInfoDetailW")
	procSetupDiGetSelectedDevice                             = modsetupapi.NewProc("SetupDiGetSelectedDevice")
	procSetupDiGetSelectedDriverW                            = modsetupapi.NewProc("SetupDiGetSelectedDriverW")
	procSetupDiOpenDevRegKey                                 = modsetupapi.NewProc("SetupDiOpenDevRegKey")
	procSetupDiSetClassInstallParamsW                        = modsetupapi.NewProc("SetupDiSetClassInstallParamsW")
	procSetupDiSetDeviceInstallParamsW                       = modsetupapi.NewProc("SetupDiSetDeviceInstallParamsW")
	procSetupDiSetDeviceRegistryPropertyW                    = modsetupapi.NewProc("SetupDiSetDeviceRegistryPropertyW")
	procSetupDiSetSelectedDevice                             = modsetupapi.NewProc("SetupDiSetSelectedDevice")
	procSetupDiSetSelectedDriverW                            = modsetupapi.NewProc("SetupDiSetSelectedDriverW")
	procSetupUninstallOEMInfW                                = modsetupapi.NewProc("SetupUninstallOEMInfW")
	procCommandLineToArgvW                                   = modshell32.NewProc("CommandLineToArgvW")
	procSHGetKnownFolderPath                                 = modshell32.NewProc("SHGetKnownFolderPath")
	procShellExecuteW                                        = modshell32.NewProc("ShellExecuteW")
	procEnumChildWindows                                     = moduser32.NewProc("EnumChildWindows")
	procEnumWindows                                          = moduser32.NewProc("EnumWindows")
	procExitWindowsEx                                        = moduser32.NewProc("ExitWindowsEx")
	procGetClassNameW                                        = moduser32.NewProc("GetClassNameW")
	procGetDesktopWindow                                     = moduser32.NewProc("GetDesktopWindow")
	procGetForegroundWindow                                  = moduser32.NewProc("GetForegroundWindow")
	procGetGUIThreadInfo                                     = moduser32.NewProc("GetGUIThreadInfo")
	procGetShellWindow                                       = moduser32.NewProc("GetShellWindow")
	procGetWindowThreadProcessId                             = moduser32.NewProc("GetWindowThreadProcessId")
	procIsWindow                                             = moduser32.NewProc("IsWindow")
	procIsWindowUnicode                                      = moduser32.NewProc("IsWindowUnicode")
	procIsWindowVisible                                      = moduser32.NewProc("IsWindowVisible")
	procMessageBoxW                                          = moduser32.NewProc("MessageBoxW")
	procCreateEnvironmentBlock                               = moduserenv.NewProc("CreateEnvironmentBlock")
	procDestroyEnvironmentBlock                              = moduserenv.NewProc("DestroyEnvironmentBlock")
	procGetUserProfileDirectoryW                             = moduserenv.NewProc("GetUserProfileDirectoryW")
	procGetFileVersionInfoSizeW                              = modversion.NewProc("GetFileVersionInfoSizeW")
	procGetFileVersionInfoW                                  = modversion.NewProc("GetFileVersionInfoW")
	procVerQueryValueW                                       = modversion.NewProc("VerQueryValueW")
	proctimeBeginPeriod                                      = modwinmm.NewProc("timeBeginPeriod")
	proctimeEndPeriod                                        = modwinmm.NewProc("timeEndPeriod")
	procWinVerifyTrustEx                                     = modwintrust.NewProc("WinVerifyTrustEx")
	procFreeAddrInfoW                                        = modws2_32.NewProc("FreeAddrInfoW")
	procGetAddrInfoW                                         = modws2_32.NewProc("GetAddrInfoW")
	procWSACleanup                                           = modws2_32.NewProc("WSACleanup")
	procWSAEnumProtocolsW                                    = modws2_32.NewProc("WSAEnumProtocolsW")
	procWSAGetOverlappedResult                               = modws2_32.NewProc("WSAGetOverlappedResult")
	procWSAIoctl                                             = modws2_32.NewProc("WSAIoctl")
	procWSALookupServiceBeginW                               = modws2_32.NewProc("WSALookupServiceBeginW")
	procWSALookupServiceEnd                                  = modws2_32.NewProc("WSALookupServiceEnd")
	procWSALookupServiceNextW                                = modws2_32.NewProc("WSALookupServiceNextW")
	procWSARecv                                              = modws2_32.NewProc("WSARecv")
	procWSARecvFrom                                          = modws2_32.NewProc("WSARecvFrom")
	procWSASend                                              = modws2_32.NewProc("WSASend")
	procWSASendTo                                            = modws2_32.NewProc("WSASendTo")
	procWSASocketW                                           = modws2_32.NewProc("WSASocketW")
	procWSAStartup                                           = modws2_32.NewProc("WSAStartup")
	procbind                                                 = modws2_32.NewProc("bind")
	procclosesocket                                          = modws2_32.NewProc("closesocket")
	procconnect                                              = modws2_32.NewProc("connect")
	procgethostbyname                                        = modws2_32.NewProc("gethostbyname")
	procgetpeername                                          = modws2_32.NewProc("getpeername")
	procgetprotobyname                                       = modws2_32.NewProc("getprotobyname")
	procgetservbyname                                        = modws2_32.NewProc("getservbyname")
	procgetsockname                                          = modws2_32.NewProc("getsockname")
	procgetsockopt                                           = modws2_32.NewProc("getsockopt")
	proclisten                                               = modws2_32.NewProc("listen")
	procntohs                                                = modws2_32.NewProc("ntohs")
	procrecvfrom                                             = modws2_32.NewProc("recvfrom")
	procsendto                                               = modws2_32.NewProc("sendto")
	procsetsockopt                                           = modws2_32.NewProc("setsockopt")
	procshutdown                                             = modws2_32.NewProc("shutdown")
	procsocket                                               = modws2_32.NewProc("socket")
	procWTSEnumerateSessionsW                                = modwtsapi32.NewProc("WTSEnumerateSessionsW")
	procWTSFreeMemory                                        = modwtsapi32.NewProc("WTSFreeMemory")
	procWTSQueryUserToken                                    = modwtsapi32.NewProc("WTSQueryUserToken")
)

func cm_Get_DevNode_Status(status *uint32, problemNumber *uint32, devInst DEVINST, flags uint32) (ret CONFIGRET) {
	r0, _, _ := syscall.Syscall6(procCM_Get_DevNode_Status.Addr(), 4, uintptr(unsafe.Pointer(status)), uintptr(unsafe.Pointer(problemNumber)), uintptr(devInst), uintptr(flags), 0, 0)
	ret = CONFIGRET(r0)
	return
}

func cm_Get_Device_Interface_List(interfaceClass *GUID, deviceID *uint16, buffer *uint16, bufferLen uint32, flags uint32) (ret CONFIGRET) {
	r0, _, _ := syscall.Syscall6(procCM_Get_Device_Interface_ListW.Addr(), 5, uintptr(unsafe.Pointer(interfaceClass)), uintptr(unsafe.Pointer(deviceID)), uintptr(unsafe.Pointer(buffer)), uintptr(bufferLen), uintptr(flags), 0)
	ret = CONFIGRET(r0)
	return
}

func cm_Get_Device_Interface_List_Size(len *uint32, interfaceClass *GUID, deviceID *uint16, flags uint32) (ret CONFIGRET) {
	r0, _, _ := syscall.Syscall6(procCM_Get_Device_Interface_List_SizeW.Addr(), 4, uintptr(unsafe.Pointer(len)), uintptr(unsafe.Pointer(interfaceClass)), uintptr(unsafe.Pointer(deviceID)), uintptr(flags), 0, 0)
	ret = CONFIGRET(r0)
	return
}

func cm_MapCrToWin32Err(configRet CONFIGRET, defaultWin32Error Errno) (ret Errno) {
	r0, _, _ := syscall.Syscall(procCM_MapCrToWin32Err.Addr(), 2, uintptr(configRet), uintptr(defaultWin32Error), 0)
	ret = Errno(r0)
	return
}

// 翻译提示:func  调整令牌组(token  令牌,  重置为默认值  bool,  新状态  *令牌组,  缓冲区长度  uint32,  前状态  *令牌组,  返回长度  *uint32)  (err  错误)  {}

// ff:
// token:
// resetToDefault:
// newstate:
// buflen:
// prevstate:
// returnlen:
// err:
func AdjustTokenGroups(token Token, resetToDefault bool, newstate *Tokengroups, buflen uint32, prevstate *Tokengroups, returnlen *uint32) (err error) {
	var _p0 uint32
	if resetToDefault {
		_p0 = 1
	}
	r1, _, e1 := syscall.Syscall6(procAdjustTokenGroups.Addr(), 6, uintptr(token), uintptr(_p0), uintptr(unsafe.Pointer(newstate)), uintptr(buflen), uintptr(unsafe.Pointer(prevstate)), uintptr(unsafe.Pointer(returnlen)))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  调整令牌权限(token  令牌,  禁用所有权限  bool,  新状态  *权限集,  缓冲区长度  uint32,  前状态  *权限集,  返回长度  *uint32)  (错误  error)  {}

// ff:
// token:
// disableAllPrivileges:
// newstate:
// buflen:
// prevstate:
// returnlen:
// err:
func AdjustTokenPrivileges(token Token, disableAllPrivileges bool, newstate *Tokenprivileges, buflen uint32, prevstate *Tokenprivileges, returnlen *uint32) (err error) {
	var _p0 uint32
	if disableAllPrivileges {
		_p0 = 1
	}
	r1, _, e1 := syscall.Syscall6(procAdjustTokenPrivileges.Addr(), 6, uintptr(token), uintptr(_p0), uintptr(unsafe.Pointer(newstate)), uintptr(buflen), uintptr(unsafe.Pointer(prevstate)), uintptr(unsafe.Pointer(returnlen)))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  分配并初始化安全标识符(identAuth  *安全权威标识符,  subAuth  字节,  subAuth0  uint32,  subAuth1  uint32,  subAuth2  uint32,  subAuth3  uint32,  subAuth4  uint32,  subAuth5  uint32,  subAuth6  uint32,  subAuth7  uint32,  sid  **安全标识符)  (错误  error)  {}

// ff:
// identAuth:
// subAuth:
// subAuth0:
// subAuth1:
// subAuth2:
// subAuth3:
// subAuth4:
// subAuth5:
// subAuth6:
// subAuth7:
// sid:
// err:
func AllocateAndInitializeSid(identAuth *SidIdentifierAuthority, subAuth byte, subAuth0 uint32, subAuth1 uint32, subAuth2 uint32, subAuth3 uint32, subAuth4 uint32, subAuth5 uint32, subAuth6 uint32, subAuth7 uint32, sid **SID) (err error) {
	r1, _, e1 := syscall.Syscall12(procAllocateAndInitializeSid.Addr(), 11, uintptr(unsafe.Pointer(identAuth)), uintptr(subAuth), uintptr(subAuth0), uintptr(subAuth1), uintptr(subAuth2), uintptr(subAuth3), uintptr(subAuth4), uintptr(subAuth5), uintptr(subAuth6), uintptr(subAuth7), uintptr(unsafe.Pointer(sid)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func buildSecurityDescriptor(owner *TRUSTEE, group *TRUSTEE, countAccessEntries uint32, accessEntries *EXPLICIT_ACCESS, countAuditEntries uint32, auditEntries *EXPLICIT_ACCESS, oldSecurityDescriptor *SECURITY_DESCRIPTOR, sizeNewSecurityDescriptor *uint32, newSecurityDescriptor **SECURITY_DESCRIPTOR) (ret error) {
	r0, _, _ := syscall.Syscall9(procBuildSecurityDescriptorW.Addr(), 9, uintptr(unsafe.Pointer(owner)), uintptr(unsafe.Pointer(group)), uintptr(countAccessEntries), uintptr(unsafe.Pointer(accessEntries)), uintptr(countAuditEntries), uintptr(unsafe.Pointer(auditEntries)), uintptr(unsafe.Pointer(oldSecurityDescriptor)), uintptr(unsafe.Pointer(sizeNewSecurityDescriptor)), uintptr(unsafe.Pointer(newSecurityDescriptor)))
	if r0 != 0 {
		ret = syscall.Errno(r0)
	}
	return
}

// 翻译提示:func  修改服务配置2(service  手柄,  信息级别  uint32,  配置信息  *字节)  (错误  错误)  {}

// ff:
// service:
// infoLevel:
// info:
// err:
func ChangeServiceConfig2(service Handle, infoLevel uint32, info *byte) (err error) {
	r1, _, e1 := syscall.Syscall(procChangeServiceConfig2W.Addr(), 3, uintptr(service), uintptr(infoLevel), uintptr(unsafe.Pointer(info)))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  修改服务配置(service  服务句柄,  服务类型  uint32,  启动类型  uint32,  错误控制  uint32,  可执行文件路径  *uint16,  负载顺序组  *uint16,  标签ID  *uint32,  依赖项  *uint16,  服务启动账户名  *uint16,  密码  *uint16,  显示名称  *uint16)  (错误  error)  {}

// ff:
// service:
// serviceType:
// startType:
// errorControl:
// binaryPathName:
// loadOrderGroup:
// tagId:
// dependencies:
// serviceStartName:
// password:
// displayName:
// err:
func ChangeServiceConfig(service Handle, serviceType uint32, startType uint32, errorControl uint32, binaryPathName *uint16, loadOrderGroup *uint16, tagId *uint32, dependencies *uint16, serviceStartName *uint16, password *uint16, displayName *uint16) (err error) {
	r1, _, e1 := syscall.Syscall12(procChangeServiceConfigW.Addr(), 11, uintptr(service), uintptr(serviceType), uintptr(startType), uintptr(errorControl), uintptr(unsafe.Pointer(binaryPathName)), uintptr(unsafe.Pointer(loadOrderGroup)), uintptr(unsafe.Pointer(tagId)), uintptr(unsafe.Pointer(dependencies)), uintptr(unsafe.Pointer(serviceStartName)), uintptr(unsafe.Pointer(password)), uintptr(unsafe.Pointer(displayName)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func checkTokenMembership(tokenHandle Token, sidToCheck *SID, isMember *int32) (err error) {
	r1, _, e1 := syscall.Syscall(procCheckTokenMembership.Addr(), 3, uintptr(tokenHandle), uintptr(unsafe.Pointer(sidToCheck)), uintptr(unsafe.Pointer(isMember)))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  关闭服务句柄(handle  句柄)  (错误  error)  {}

// ff:
// handle:
// err:
func CloseServiceHandle(handle Handle) (err error) {
	r1, _, e1 := syscall.Syscall(procCloseServiceHandle.Addr(), 1, uintptr(handle), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  控制服务(service句柄  Handle,  控制方式  uint32,  状态  *服务状态  SERVICE_STATUS)  (错误  error)  {}

// ff:
// service:
// control:
// status:
// err:
func ControlService(service Handle, control uint32, status *SERVICE_STATUS) (err error) {
	r1, _, e1 := syscall.Syscall(procControlService.Addr(), 3, uintptr(service), uintptr(control), uintptr(unsafe.Pointer(status)))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func convertSecurityDescriptorToStringSecurityDescriptor(sd *SECURITY_DESCRIPTOR, revision uint32, securityInformation SECURITY_INFORMATION, str **uint16, strLen *uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procConvertSecurityDescriptorToStringSecurityDescriptorW.Addr(), 5, uintptr(unsafe.Pointer(sd)), uintptr(revision), uintptr(securityInformation), uintptr(unsafe.Pointer(str)), uintptr(unsafe.Pointer(strLen)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  将Sid转换为字符串Sid(sid  *安全标识符,  字符串Sid  **宽字符指针)  (错误  error)  {}

// ff:
// sid:
// stringSid:
// err:
func ConvertSidToStringSid(sid *SID, stringSid **uint16) (err error) {
	r1, _, e1 := syscall.Syscall(procConvertSidToStringSidW.Addr(), 2, uintptr(unsafe.Pointer(sid)), uintptr(unsafe.Pointer(stringSid)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func convertStringSecurityDescriptorToSecurityDescriptor(str string, revision uint32, sd **SECURITY_DESCRIPTOR, size *uint32) (err error) {
	var _p0 *uint16
	_p0, err = syscall.UTF16PtrFromString(str)
	if err != nil {
		return
	}
	return _convertStringSecurityDescriptorToSecurityDescriptor(_p0, revision, sd, size)
}

func _convertStringSecurityDescriptorToSecurityDescriptor(str *uint16, revision uint32, sd **SECURITY_DESCRIPTOR, size *uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procConvertStringSecurityDescriptorToSecurityDescriptorW.Addr(), 4, uintptr(unsafe.Pointer(str)), uintptr(revision), uintptr(unsafe.Pointer(sd)), uintptr(unsafe.Pointer(size)), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  转换字符串SID到SID(stringSid  *字符串SID,  sid  **安全标识符)  (错误  error)  {}

// ff:
// stringSid:
// sid:
// err:
func ConvertStringSidToSid(stringSid *uint16, sid **SID) (err error) {
	r1, _, e1 := syscall.Syscall(procConvertStringSidToSidW.Addr(), 2, uintptr(unsafe.Pointer(stringSid)), uintptr(unsafe.Pointer(sid)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  复制安全标识符(destSidLen  uint32,  目标Sid  *SID,  源Sid  *SID)  (错误  error)  {}

// ff:
// destSidLen:
// destSid:
// srcSid:
// err:
func CopySid(destSidLen uint32, destSid *SID, srcSid *SID) (err error) {
	r1, _, e1 := syscall.Syscall(procCopySid.Addr(), 3, uintptr(destSidLen), uintptr(unsafe.Pointer(destSid)), uintptr(unsafe.Pointer(srcSid)))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  作为用户创建进程(token  用户令牌,  应用程序名称  *uint16,  命令行  *uint16,  进程安全属性  *SecurityAttributes,  线程安全属性  *SecurityAttributes,  继承句柄  bool,  创建标志  uint32,  环境变量  *uint16,  当前目录  *uint16,  启动信息  *StartupInfo,  进程信息  *ProcessInformation)  (错误  error)  {}

// ff:
// token:
// appName:
// commandLine:
// procSecurity:
// threadSecurity:
// inheritHandles:
// creationFlags:
// env:
// currentDir:
// startupInfo:
// outProcInfo:
// err:
func CreateProcessAsUser(token Token, appName *uint16, commandLine *uint16, procSecurity *SecurityAttributes, threadSecurity *SecurityAttributes, inheritHandles bool, creationFlags uint32, env *uint16, currentDir *uint16, startupInfo *StartupInfo, outProcInfo *ProcessInformation) (err error) {
	var _p0 uint32
	if inheritHandles {
		_p0 = 1
	}
	r1, _, e1 := syscall.Syscall12(procCreateProcessAsUserW.Addr(), 11, uintptr(token), uintptr(unsafe.Pointer(appName)), uintptr(unsafe.Pointer(commandLine)), uintptr(unsafe.Pointer(procSecurity)), uintptr(unsafe.Pointer(threadSecurity)), uintptr(_p0), uintptr(creationFlags), uintptr(unsafe.Pointer(env)), uintptr(unsafe.Pointer(currentDir)), uintptr(unsafe.Pointer(startupInfo)), uintptr(unsafe.Pointer(outProcInfo)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  创建服务(mgr  手柄,  服务名称  *uint16,  显示名称  *uint16,  访问权限  uint32,  服务类型  uint32,  启动类型  uint32,  错误控制  uint32,  路径名称  *uint16,  负载顺序组  *uint16,  标签ID  *uint32,  依赖项  *uint16,  服务启动名称  *uint16,  密码  *uint16)  (服务句柄  手柄,  错误  error)  {}

// ff:
// mgr:
// serviceName:
// displayName:
// access:
// srvType:
// startType:
// errCtl:
// pathName:
// loadOrderGroup:
// tagId:
// dependencies:
// serviceStartName:
// password:
// handle:
// err:
func CreateService(mgr Handle, serviceName *uint16, displayName *uint16, access uint32, srvType uint32, startType uint32, errCtl uint32, pathName *uint16, loadOrderGroup *uint16, tagId *uint32, dependencies *uint16, serviceStartName *uint16, password *uint16) (handle Handle, err error) {
	r0, _, e1 := syscall.Syscall15(procCreateServiceW.Addr(), 13, uintptr(mgr), uintptr(unsafe.Pointer(serviceName)), uintptr(unsafe.Pointer(displayName)), uintptr(access), uintptr(srvType), uintptr(startType), uintptr(errCtl), uintptr(unsafe.Pointer(pathName)), uintptr(unsafe.Pointer(loadOrderGroup)), uintptr(unsafe.Pointer(tagId)), uintptr(unsafe.Pointer(dependencies)), uintptr(unsafe.Pointer(serviceStartName)), uintptr(unsafe.Pointer(password)), 0, 0)
	handle = Handle(r0)
	if handle == 0 {
		err = errnoErr(e1)
	}
	return
}

func createWellKnownSid(sidType WELL_KNOWN_SID_TYPE, domainSid *SID, sid *SID, sizeSid *uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procCreateWellKnownSid.Addr(), 4, uintptr(sidType), uintptr(unsafe.Pointer(domainSid)), uintptr(unsafe.Pointer(sid)), uintptr(unsafe.Pointer(sizeSid)), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  加密获取上下文(provhandle  *句柄,  容器  *uint16,  提供者  *uint16,  类型  uint32,  标志  uint32)  (错误  error)  {}

// ff:
// provhandle:
// container:
// provider:
// provtype:
// flags:
// err:
func CryptAcquireContext(provhandle *Handle, container *uint16, provider *uint16, provtype uint32, flags uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procCryptAcquireContextW.Addr(), 5, uintptr(unsafe.Pointer(provhandle)), uintptr(unsafe.Pointer(container)), uintptr(unsafe.Pointer(provider)), uintptr(provtype), uintptr(flags), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  生成随机数(provhandle  模块句柄,  buflen  字节长度,  buf  字节缓冲区指针)  (err  错误)  {}

// ff:
// provhandle:
// buflen:
// buf:
// err:
func CryptGenRandom(provhandle Handle, buflen uint32, buf *byte) (err error) {
	r1, _, e1 := syscall.Syscall(procCryptGenRandom.Addr(), 3, uintptr(provhandle), uintptr(buflen), uintptr(unsafe.Pointer(buf)))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  加密释放上下文(provHandle  模块句柄,  标志  uint32)  (错误  error)  {}

// ff:
// provhandle:
// flags:
// err:
func CryptReleaseContext(provhandle Handle, flags uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procCryptReleaseContext.Addr(), 2, uintptr(provhandle), uintptr(flags), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  删除服务(service  服务句柄)  (错误  error)  {}

// ff:
// service:
// err:
func DeleteService(service Handle) (err error) {
	r1, _, e1 := syscall.Syscall(procDeleteService.Addr(), 1, uintptr(service), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  注销事件源(handle  手柄)  (错误  error)  {}

// ff:
// handle:
// err:
func DeregisterEventSource(handle Handle) (err error) {
	r1, _, e1 := syscall.Syscall(procDeregisterEventSource.Addr(), 1, uintptr(handle), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  复制令牌(existingToken  令牌,  所需访问权限  uint32,  令牌属性  *安全属性,  模仿级别  uint32,  令牌类型  uint32,  新令牌  *令牌)  (错误  error)  {}

// ff:
// existingToken:
// desiredAccess:
// tokenAttributes:
// impersonationLevel:
// tokenType:
// newToken:
// err:
func DuplicateTokenEx(existingToken Token, desiredAccess uint32, tokenAttributes *SecurityAttributes, impersonationLevel uint32, tokenType uint32, newToken *Token) (err error) {
	r1, _, e1 := syscall.Syscall6(procDuplicateTokenEx.Addr(), 6, uintptr(existingToken), uintptr(desiredAccess), uintptr(unsafe.Pointer(tokenAttributes)), uintptr(impersonationLevel), uintptr(tokenType), uintptr(unsafe.Pointer(newToken)))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  列举依赖服务(service  手柄,  活动状态  uint32,  服务状态  *ENUM_SERVICE_STATUS,  缓冲区大小  uint32,  需要字节数  *uint32,  返回服务数  *uint32)  (错误  error)  {}

// ff:
// service:
// activityState:
// services:
// buffSize:
// bytesNeeded:
// servicesReturned:
// err:
func EnumDependentServices(service Handle, activityState uint32, services *ENUM_SERVICE_STATUS, buffSize uint32, bytesNeeded *uint32, servicesReturned *uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procEnumDependentServicesW.Addr(), 6, uintptr(service), uintptr(activityState), uintptr(unsafe.Pointer(services)), uintptr(buffSize), uintptr(unsafe.Pointer(bytesNeeded)), uintptr(unsafe.Pointer(servicesReturned)))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  列举服务状态(mgr  手柄,  信息级别  uint32,  服务类型  uint32,  服务状态  uint32,  服务数据  *byte,  缓冲区大小  uint32,  需要字节数  *uint32,  返回服务数  *uint32,  恢复句柄  *uint32,  组名  *uint16)  (错误  error)  {}

// ff:
// mgr:
// infoLevel:
// serviceType:
// serviceState:
// services:
// bufSize:
// bytesNeeded:
// servicesReturned:
// resumeHandle:
// groupName:
// err:
func EnumServicesStatusEx(mgr Handle, infoLevel uint32, serviceType uint32, serviceState uint32, services *byte, bufSize uint32, bytesNeeded *uint32, servicesReturned *uint32, resumeHandle *uint32, groupName *uint16) (err error) {
	r1, _, e1 := syscall.Syscall12(procEnumServicesStatusExW.Addr(), 10, uintptr(mgr), uintptr(infoLevel), uintptr(serviceType), uintptr(serviceState), uintptr(unsafe.Pointer(services)), uintptr(bufSize), uintptr(unsafe.Pointer(bytesNeeded)), uintptr(unsafe.Pointer(servicesReturned)), uintptr(unsafe.Pointer(resumeHandle)), uintptr(unsafe.Pointer(groupName)), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  相等Sid(sid1  *安全标识符,  sid2  *安全标识符)  (是否相等  bool)  {}

// ff:
// sid1:
// sid2:
// isEqual:
func EqualSid(sid1 *SID, sid2 *SID) (isEqual bool) {
	r0, _, _ := syscall.Syscall(procEqualSid.Addr(), 2, uintptr(unsafe.Pointer(sid1)), uintptr(unsafe.Pointer(sid2)), 0)
	isEqual = r0 != 0
	return
}

// 翻译提示:func  释放Sid(sid  *安全标识符)  (错误  error)  {}

// ff:
// sid:
// err:
func FreeSid(sid *SID) (err error) {
	r1, _, e1 := syscall.Syscall(procFreeSid.Addr(), 1, uintptr(unsafe.Pointer(sid)), 0, 0)
	if r1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  获取SID长度(Sid  *安全标识符)  (长度  uint32)  {}

// ff:
// sid:
// len:
func GetLengthSid(sid *SID) (len uint32) {
	r0, _, _ := syscall.Syscall(procGetLengthSid.Addr(), 1, uintptr(unsafe.Pointer(sid)), 0, 0)
	len = uint32(r0)
	return
}

func getNamedSecurityInfo(objectName string, objectType SE_OBJECT_TYPE, securityInformation SECURITY_INFORMATION, owner **SID, group **SID, dacl **ACL, sacl **ACL, sd **SECURITY_DESCRIPTOR) (ret error) {
	var _p0 *uint16
	_p0, ret = syscall.UTF16PtrFromString(objectName)
	if ret != nil {
		return
	}
	return _getNamedSecurityInfo(_p0, objectType, securityInformation, owner, group, dacl, sacl, sd)
}

func _getNamedSecurityInfo(objectName *uint16, objectType SE_OBJECT_TYPE, securityInformation SECURITY_INFORMATION, owner **SID, group **SID, dacl **ACL, sacl **ACL, sd **SECURITY_DESCRIPTOR) (ret error) {
	r0, _, _ := syscall.Syscall9(procGetNamedSecurityInfoW.Addr(), 8, uintptr(unsafe.Pointer(objectName)), uintptr(objectType), uintptr(securityInformation), uintptr(unsafe.Pointer(owner)), uintptr(unsafe.Pointer(group)), uintptr(unsafe.Pointer(dacl)), uintptr(unsafe.Pointer(sacl)), uintptr(unsafe.Pointer(sd)), 0)
	if r0 != 0 {
		ret = syscall.Errno(r0)
	}
	return
}

func getSecurityDescriptorControl(sd *SECURITY_DESCRIPTOR, control *SECURITY_DESCRIPTOR_CONTROL, revision *uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procGetSecurityDescriptorControl.Addr(), 3, uintptr(unsafe.Pointer(sd)), uintptr(unsafe.Pointer(control)), uintptr(unsafe.Pointer(revision)))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func getSecurityDescriptorDacl(sd *SECURITY_DESCRIPTOR, daclPresent *bool, dacl **ACL, daclDefaulted *bool) (err error) {
	var _p0 uint32
	if *daclPresent {
		_p0 = 1
	}
	var _p1 uint32
	if *daclDefaulted {
		_p1 = 1
	}
	r1, _, e1 := syscall.Syscall6(procGetSecurityDescriptorDacl.Addr(), 4, uintptr(unsafe.Pointer(sd)), uintptr(unsafe.Pointer(&_p0)), uintptr(unsafe.Pointer(dacl)), uintptr(unsafe.Pointer(&_p1)), 0, 0)
	*daclPresent = _p0 != 0
	*daclDefaulted = _p1 != 0
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func getSecurityDescriptorGroup(sd *SECURITY_DESCRIPTOR, group **SID, groupDefaulted *bool) (err error) {
	var _p0 uint32
	if *groupDefaulted {
		_p0 = 1
	}
	r1, _, e1 := syscall.Syscall(procGetSecurityDescriptorGroup.Addr(), 3, uintptr(unsafe.Pointer(sd)), uintptr(unsafe.Pointer(group)), uintptr(unsafe.Pointer(&_p0)))
	*groupDefaulted = _p0 != 0
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func getSecurityDescriptorLength(sd *SECURITY_DESCRIPTOR) (len uint32) {
	r0, _, _ := syscall.Syscall(procGetSecurityDescriptorLength.Addr(), 1, uintptr(unsafe.Pointer(sd)), 0, 0)
	len = uint32(r0)
	return
}

func getSecurityDescriptorOwner(sd *SECURITY_DESCRIPTOR, owner **SID, ownerDefaulted *bool) (err error) {
	var _p0 uint32
	if *ownerDefaulted {
		_p0 = 1
	}
	r1, _, e1 := syscall.Syscall(procGetSecurityDescriptorOwner.Addr(), 3, uintptr(unsafe.Pointer(sd)), uintptr(unsafe.Pointer(owner)), uintptr(unsafe.Pointer(&_p0)))
	*ownerDefaulted = _p0 != 0
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func getSecurityDescriptorRMControl(sd *SECURITY_DESCRIPTOR, rmControl *uint8) (ret error) {
	r0, _, _ := syscall.Syscall(procGetSecurityDescriptorRMControl.Addr(), 2, uintptr(unsafe.Pointer(sd)), uintptr(unsafe.Pointer(rmControl)), 0)
	if r0 != 0 {
		ret = syscall.Errno(r0)
	}
	return
}

func getSecurityDescriptorSacl(sd *SECURITY_DESCRIPTOR, saclPresent *bool, sacl **ACL, saclDefaulted *bool) (err error) {
	var _p0 uint32
	if *saclPresent {
		_p0 = 1
	}
	var _p1 uint32
	if *saclDefaulted {
		_p1 = 1
	}
	r1, _, e1 := syscall.Syscall6(procGetSecurityDescriptorSacl.Addr(), 4, uintptr(unsafe.Pointer(sd)), uintptr(unsafe.Pointer(&_p0)), uintptr(unsafe.Pointer(sacl)), uintptr(unsafe.Pointer(&_p1)), 0, 0)
	*saclPresent = _p0 != 0
	*saclDefaulted = _p1 != 0
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func getSecurityInfo(handle Handle, objectType SE_OBJECT_TYPE, securityInformation SECURITY_INFORMATION, owner **SID, group **SID, dacl **ACL, sacl **ACL, sd **SECURITY_DESCRIPTOR) (ret error) {
	r0, _, _ := syscall.Syscall9(procGetSecurityInfo.Addr(), 8, uintptr(handle), uintptr(objectType), uintptr(securityInformation), uintptr(unsafe.Pointer(owner)), uintptr(unsafe.Pointer(group)), uintptr(unsafe.Pointer(dacl)), uintptr(unsafe.Pointer(sacl)), uintptr(unsafe.Pointer(sd)), 0)
	if r0 != 0 {
		ret = syscall.Errno(r0)
	}
	return
}

func getSidIdentifierAuthority(sid *SID) (authority *SidIdentifierAuthority) {
	r0, _, _ := syscall.Syscall(procGetSidIdentifierAuthority.Addr(), 1, uintptr(unsafe.Pointer(sid)), 0, 0)
	authority = (*SidIdentifierAuthority)(unsafe.Pointer(r0))
	return
}

func getSidSubAuthority(sid *SID, index uint32) (subAuthority *uint32) {
	r0, _, _ := syscall.Syscall(procGetSidSubAuthority.Addr(), 2, uintptr(unsafe.Pointer(sid)), uintptr(index), 0)
	subAuthority = (*uint32)(unsafe.Pointer(r0))
	return
}

func getSidSubAuthorityCount(sid *SID) (count *uint8) {
	r0, _, _ := syscall.Syscall(procGetSidSubAuthorityCount.Addr(), 1, uintptr(unsafe.Pointer(sid)), 0, 0)
	count = (*uint8)(unsafe.Pointer(r0))
	return
}

// 翻译提示:func  获取Token信息(token  令牌,  info类别  uint32,  info  *字节,  info长度  uint32,  返回的长度  *uint32)  (错误  error)  {}

// ff:
// token:
// infoClass:
// info:
// infoLen:
// returnedLen:
// err:
func GetTokenInformation(token Token, infoClass uint32, info *byte, infoLen uint32, returnedLen *uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procGetTokenInformation.Addr(), 5, uintptr(token), uintptr(infoClass), uintptr(unsafe.Pointer(info)), uintptr(infoLen), uintptr(unsafe.Pointer(returnedLen)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  自我模拟(模拟级别  uint32)  (错误  error)  {}

// ff:
// impersonationlevel:
// err:
func ImpersonateSelf(impersonationlevel uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procImpersonateSelf.Addr(), 1, uintptr(impersonationlevel), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func initializeSecurityDescriptor(absoluteSD *SECURITY_DESCRIPTOR, revision uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procInitializeSecurityDescriptor.Addr(), 2, uintptr(unsafe.Pointer(absoluteSD)), uintptr(revision), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  开始系统关机Ex(机器名  *uint16,  消息  *uint16,  超时时间  uint32,  强制关闭应用  bool,  重启后关机  bool,  原因  uint32)  (错误  error)  {}

// ff:
// machineName:
// message:
// timeout:
// forceAppsClosed:
// rebootAfterShutdown:
// reason:
// err:
func InitiateSystemShutdownEx(machineName *uint16, message *uint16, timeout uint32, forceAppsClosed bool, rebootAfterShutdown bool, reason uint32) (err error) {
	var _p0 uint32
	if forceAppsClosed {
		_p0 = 1
	}
	var _p1 uint32
	if rebootAfterShutdown {
		_p1 = 1
	}
	r1, _, e1 := syscall.Syscall6(procInitiateSystemShutdownExW.Addr(), 6, uintptr(unsafe.Pointer(machineName)), uintptr(unsafe.Pointer(message)), uintptr(timeout), uintptr(_p0), uintptr(_p1), uintptr(reason))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func isTokenRestricted(tokenHandle Token) (ret bool, err error) {
	r0, _, e1 := syscall.Syscall(procIsTokenRestricted.Addr(), 1, uintptr(tokenHandle), 0, 0)
	ret = r0 != 0
	if !ret {
		err = errnoErr(e1)
	}
	return
}

func isValidSecurityDescriptor(sd *SECURITY_DESCRIPTOR) (isValid bool) {
	r0, _, _ := syscall.Syscall(procIsValidSecurityDescriptor.Addr(), 1, uintptr(unsafe.Pointer(sd)), 0, 0)
	isValid = r0 != 0
	return
}

func isValidSid(sid *SID) (isValid bool) {
	r0, _, _ := syscall.Syscall(procIsValidSid.Addr(), 1, uintptr(unsafe.Pointer(sid)), 0, 0)
	isValid = r0 != 0
	return
}

func isWellKnownSid(sid *SID, sidType WELL_KNOWN_SID_TYPE) (isWellKnown bool) {
	r0, _, _ := syscall.Syscall(procIsWellKnownSid.Addr(), 2, uintptr(unsafe.Pointer(sid)), uintptr(sidType), 0)
	isWellKnown = r0 != 0
	return
}

// 翻译提示:func  查找账户名称(systemName  *uint16,  accountName  *uint16,  sid  *SID,  sidLength  *uint32,  引用域名  *uint16,  引用域名长度  *uint32,  账户类型  *uint32)  (错误  error)  {}

// ff:
// systemName:
// accountName:
// sid:
// sidLen:
// refdDomainName:
// refdDomainNameLen:
// use:
// err:
func LookupAccountName(systemName *uint16, accountName *uint16, sid *SID, sidLen *uint32, refdDomainName *uint16, refdDomainNameLen *uint32, use *uint32) (err error) {
	r1, _, e1 := syscall.Syscall9(procLookupAccountNameW.Addr(), 7, uintptr(unsafe.Pointer(systemName)), uintptr(unsafe.Pointer(accountName)), uintptr(unsafe.Pointer(sid)), uintptr(unsafe.Pointer(sidLen)), uintptr(unsafe.Pointer(refdDomainName)), uintptr(unsafe.Pointer(refdDomainNameLen)), uintptr(unsafe.Pointer(use)), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  查询账户Sid(systemName  *字符串,  sid  *安全标识符,  name  *字符串,  nameLen  *uint32,  refdDomainName  *字符串,  refdDomainNameLen  *uint32,  use  *uint32)  (错误  error)  {}

// ff:
// systemName:
// sid:
// name:
// nameLen:
// refdDomainName:
// refdDomainNameLen:
// use:
// err:
func LookupAccountSid(systemName *uint16, sid *SID, name *uint16, nameLen *uint32, refdDomainName *uint16, refdDomainNameLen *uint32, use *uint32) (err error) {
	r1, _, e1 := syscall.Syscall9(procLookupAccountSidW.Addr(), 7, uintptr(unsafe.Pointer(systemName)), uintptr(unsafe.Pointer(sid)), uintptr(unsafe.Pointer(name)), uintptr(unsafe.Pointer(nameLen)), uintptr(unsafe.Pointer(refdDomainName)), uintptr(unsafe.Pointer(refdDomainNameLen)), uintptr(unsafe.Pointer(use)), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  查询权限值(systemName  *uint16,  privilegeName  *uint16,  luid  *LUID)  (error)  {}

// ff:
// systemname:
// name:
// luid:
// err:
func LookupPrivilegeValue(systemname *uint16, name *uint16, luid *LUID) (err error) {
	r1, _, e1 := syscall.Syscall(procLookupPrivilegeValueW.Addr(), 3, uintptr(unsafe.Pointer(systemname)), uintptr(unsafe.Pointer(name)), uintptr(unsafe.Pointer(luid)))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func makeAbsoluteSD(selfRelativeSD *SECURITY_DESCRIPTOR, absoluteSD *SECURITY_DESCRIPTOR, absoluteSDSize *uint32, dacl *ACL, daclSize *uint32, sacl *ACL, saclSize *uint32, owner *SID, ownerSize *uint32, group *SID, groupSize *uint32) (err error) {
	r1, _, e1 := syscall.Syscall12(procMakeAbsoluteSD.Addr(), 11, uintptr(unsafe.Pointer(selfRelativeSD)), uintptr(unsafe.Pointer(absoluteSD)), uintptr(unsafe.Pointer(absoluteSDSize)), uintptr(unsafe.Pointer(dacl)), uintptr(unsafe.Pointer(daclSize)), uintptr(unsafe.Pointer(sacl)), uintptr(unsafe.Pointer(saclSize)), uintptr(unsafe.Pointer(owner)), uintptr(unsafe.Pointer(ownerSize)), uintptr(unsafe.Pointer(group)), uintptr(unsafe.Pointer(groupSize)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func makeSelfRelativeSD(absoluteSD *SECURITY_DESCRIPTOR, selfRelativeSD *SECURITY_DESCRIPTOR, selfRelativeSDSize *uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procMakeSelfRelativeSD.Addr(), 3, uintptr(unsafe.Pointer(absoluteSD)), uintptr(unsafe.Pointer(selfRelativeSD)), uintptr(unsafe.Pointer(selfRelativeSDSize)))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  监听服务状态变化(service  服务句柄,  状态变更掩码  uint32,  通知器  *服务通知)  (错误  error)  {}

// ff:
// service:
// notifyMask:
// notifier:
// ret:
func NotifyServiceStatusChange(service Handle, notifyMask uint32, notifier *SERVICE_NOTIFY) (ret error) {
	r0, _, _ := syscall.Syscall(procNotifyServiceStatusChangeW.Addr(), 3, uintptr(service), uintptr(notifyMask), uintptr(unsafe.Pointer(notifier)))
	if r0 != 0 {
		ret = syscall.Errno(r0)
	}
	return
}

// 翻译提示:func  打开进程令牌(process  进程句柄,  访问权限  uint32,  令牌  *安全令牌)  (错误  error)  {}

// ff:
// process:
// access:
// token:
// err:
func OpenProcessToken(process Handle, access uint32, token *Token) (err error) {
	r1, _, e1 := syscall.Syscall(procOpenProcessToken.Addr(), 3, uintptr(process), uintptr(access), uintptr(unsafe.Pointer(token)))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  打开服务控制管理器(machineName  *字节序列,  databaseName  *字节序列,  access  权限  uint32)  (句柄  Handle,  错误  error)  {}
// 
// //  注意：在实际编程中，通常不建议将函数名翻译成中文，因为这可能影响代码的可读性和标准性。以上翻译仅作为理解意义上的转换。

// ff:
// machineName:
// databaseName:
// access:
// handle:
// err:
func OpenSCManager(machineName *uint16, databaseName *uint16, access uint32) (handle Handle, err error) {
	r0, _, e1 := syscall.Syscall(procOpenSCManagerW.Addr(), 3, uintptr(unsafe.Pointer(machineName)), uintptr(unsafe.Pointer(databaseName)), uintptr(access))
	handle = Handle(r0)
	if handle == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  打开服务管理器(mgr  操作句柄,  服务名称  *uint16,  访问权限  uint32)  (服务句柄  Handle,  错误  error)  {}

// ff:
// mgr:
// serviceName:
// access:
// handle:
// err:
func OpenService(mgr Handle, serviceName *uint16, access uint32) (handle Handle, err error) {
	r0, _, e1 := syscall.Syscall(procOpenServiceW.Addr(), 3, uintptr(mgr), uintptr(unsafe.Pointer(serviceName)), uintptr(access))
	handle = Handle(r0)
	if handle == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  打开线程令牌(thread  进程句柄,  访问权限  uint32,  以自身身份打开  bool,  令牌  *安全令牌)  (错误  error)  {}

// ff:
// thread:
// access:
// openAsSelf:
// token:
// err:
func OpenThreadToken(thread Handle, access uint32, openAsSelf bool, token *Token) (err error) {
	var _p0 uint32
	if openAsSelf {
		_p0 = 1
	}
	r1, _, e1 := syscall.Syscall6(procOpenThreadToken.Addr(), 4, uintptr(thread), uintptr(access), uintptr(_p0), uintptr(unsafe.Pointer(token)), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  查询服务配置2(service  手柄,  信息级别  uint32,  缓冲区  *byte,  缓冲区大小  uint32,  需要字节数  *uint32)  (错误  error)  {}

// ff:
// service:
// infoLevel:
// buff:
// buffSize:
// bytesNeeded:
// err:
func QueryServiceConfig2(service Handle, infoLevel uint32, buff *byte, buffSize uint32, bytesNeeded *uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procQueryServiceConfig2W.Addr(), 5, uintptr(service), uintptr(infoLevel), uintptr(unsafe.Pointer(buff)), uintptr(buffSize), uintptr(unsafe.Pointer(bytesNeeded)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  查询服务配置(service  手柄,  serviceConfig  *服务配置结构体,  缓冲区大小  uint32,  需要字节数  *uint32)  (错误  error)  {}

// ff:
// service:
// serviceConfig:
// bufSize:
// bytesNeeded:
// err:
func QueryServiceConfig(service Handle, serviceConfig *QUERY_SERVICE_CONFIG, bufSize uint32, bytesNeeded *uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procQueryServiceConfigW.Addr(), 4, uintptr(service), uintptr(unsafe.Pointer(serviceConfig)), uintptr(bufSize), uintptr(unsafe.Pointer(bytesNeeded)), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  查询服务动态信息(service  服务句柄,  infoLevel  信息级别,  dynamicInfo  安全指针)  (err  错误)  {}

// ff:
// service:
// infoLevel:
// dynamicInfo:
// err:
func QueryServiceDynamicInformation(service Handle, infoLevel uint32, dynamicInfo unsafe.Pointer) (err error) {
	err = procQueryServiceDynamicInformation.Find()
	if err != nil {
		return
	}
	r1, _, e1 := syscall.Syscall(procQueryServiceDynamicInformation.Addr(), 3, uintptr(service), uintptr(infoLevel), uintptr(dynamicInfo))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  查询服务锁定状态(mgr  手柄,  锁定状态  *服务锁定状态结构体,  缓冲区大小  uint32,  需要字节数  *uint32)  (错误  error)  {}

// ff:
// mgr:
// lockStatus:
// bufSize:
// bytesNeeded:
// err:
func QueryServiceLockStatus(mgr Handle, lockStatus *QUERY_SERVICE_LOCK_STATUS, bufSize uint32, bytesNeeded *uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procQueryServiceLockStatusW.Addr(), 4, uintptr(mgr), uintptr(unsafe.Pointer(lockStatus)), uintptr(bufSize), uintptr(unsafe.Pointer(bytesNeeded)), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  查询服务状态(service  服务句柄,  status  *服务状态)  (错误  error)  {}

// ff:
// service:
// status:
// err:
func QueryServiceStatus(service Handle, status *SERVICE_STATUS) (err error) {
	r1, _, e1 := syscall.Syscall(procQueryServiceStatus.Addr(), 2, uintptr(service), uintptr(unsafe.Pointer(status)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  查询服务状态扩展(service句柄  Handle,  信息级别  uint32,  缓冲区  *byte,  缓冲区大小  uint32,  需要字节数  *uint32)  (错误  error)  {}

// ff:
// service:
// infoLevel:
// buff:
// buffSize:
// bytesNeeded:
// err:
func QueryServiceStatusEx(service Handle, infoLevel uint32, buff *byte, buffSize uint32, bytesNeeded *uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procQueryServiceStatusEx.Addr(), 5, uintptr(service), uintptr(infoLevel), uintptr(unsafe.Pointer(buff)), uintptr(buffSize), uintptr(unsafe.Pointer(bytesNeeded)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  关闭注册表键(key  手柄)  (注册表错误  error)  {}

// ff:
// key:
// regerrno:
func RegCloseKey(key Handle) (regerrno error) {
	r0, _, _ := syscall.Syscall(procRegCloseKey.Addr(), 1, uintptr(key), 0, 0)
	if r0 != 0 {
		regerrno = syscall.Errno(r0)
	}
	return
}

// 翻译提示:func  注册表枚举键(key  手柄,  index  uint32,  名称  **uint16,  名称长度  *uint32,  保留  *uint32,  类别  **uint16,  类别长度  *uint32,  最后写入时间  *Filetime)  (注册表错误  error)  {}

// ff:
// key:
// index:
// name:
// nameLen:
// reserved:
// class:
// classLen:
// lastWriteTime:
// regerrno:
func RegEnumKeyEx(key Handle, index uint32, name *uint16, nameLen *uint32, reserved *uint32, class *uint16, classLen *uint32, lastWriteTime *Filetime) (regerrno error) {
	r0, _, _ := syscall.Syscall9(procRegEnumKeyExW.Addr(), 8, uintptr(key), uintptr(index), uintptr(unsafe.Pointer(name)), uintptr(unsafe.Pointer(nameLen)), uintptr(unsafe.Pointer(reserved)), uintptr(unsafe.Pointer(class)), uintptr(unsafe.Pointer(classLen)), uintptr(unsafe.Pointer(lastWriteTime)), 0)
	if r0 != 0 {
		regerrno = syscall.Errno(r0)
	}
	return
}

// 翻译提示:func  注册监控键值更改(key  手柄,  监控子树  是否,  通知过滤器  uint32,  事件  手柄,  异步  是否)  (注册错误  错误)  {}

// ff:
// key:
// watchSubtree:
// notifyFilter:
// event:
// asynchronous:
// regerrno:
func RegNotifyChangeKeyValue(key Handle, watchSubtree bool, notifyFilter uint32, event Handle, asynchronous bool) (regerrno error) {
	var _p0 uint32
	if watchSubtree {
		_p0 = 1
	}
	var _p1 uint32
	if asynchronous {
		_p1 = 1
	}
	r0, _, _ := syscall.Syscall6(procRegNotifyChangeKeyValue.Addr(), 5, uintptr(key), uintptr(_p0), uintptr(notifyFilter), uintptr(event), uintptr(_p1), 0)
	if r0 != 0 {
		regerrno = syscall.Errno(r0)
	}
	return
}

// 翻译提示:func  注册表打开键扩展(key  手柄,  子键  *uint16,  选项  uint32,  所需访问权限  uint32,  结果  *手柄)  (注册表错误码  error)  {}

// ff:
// key:
// subkey:
// options:
// desiredAccess:
// result:
// regerrno:
func RegOpenKeyEx(key Handle, subkey *uint16, options uint32, desiredAccess uint32, result *Handle) (regerrno error) {
	r0, _, _ := syscall.Syscall6(procRegOpenKeyExW.Addr(), 5, uintptr(key), uintptr(unsafe.Pointer(subkey)), uintptr(options), uintptr(desiredAccess), uintptr(unsafe.Pointer(result)), 0)
	if r0 != 0 {
		regerrno = syscall.Errno(r0)
	}
	return
}

// 翻译提示:func  注册表查询信息键(key  手柄,  类别  *uint16,  类别长度  *uint32,  保留  *uint32,  子键数量  *uint32,  最大子键长度  *uint32,  最大类别长度  *uint32,  值数量  *uint32,  最大值名称长度  *uint32,  最大值数据长度  *uint32,  安全描述符长度  *uint32,  最后写入时间  *Filetime)  (注册表错误码  error)  {}

// ff:
// key:
// class:
// classLen:
// reserved:
// subkeysLen:
// maxSubkeyLen:
// maxClassLen:
// valuesLen:
// maxValueNameLen:
// maxValueLen:
// saLen:
// lastWriteTime:
// regerrno:
func RegQueryInfoKey(key Handle, class *uint16, classLen *uint32, reserved *uint32, subkeysLen *uint32, maxSubkeyLen *uint32, maxClassLen *uint32, valuesLen *uint32, maxValueNameLen *uint32, maxValueLen *uint32, saLen *uint32, lastWriteTime *Filetime) (regerrno error) {
	r0, _, _ := syscall.Syscall12(procRegQueryInfoKeyW.Addr(), 12, uintptr(key), uintptr(unsafe.Pointer(class)), uintptr(unsafe.Pointer(classLen)), uintptr(unsafe.Pointer(reserved)), uintptr(unsafe.Pointer(subkeysLen)), uintptr(unsafe.Pointer(maxSubkeyLen)), uintptr(unsafe.Pointer(maxClassLen)), uintptr(unsafe.Pointer(valuesLen)), uintptr(unsafe.Pointer(maxValueNameLen)), uintptr(unsafe.Pointer(maxValueLen)), uintptr(unsafe.Pointer(saLen)), uintptr(unsafe.Pointer(lastWriteTime)))
	if r0 != 0 {
		regerrno = syscall.Errno(r0)
	}
	return
}

// 翻译提示:func  注册表查询值(key  手柄,  名称  *uint16,  保留  *uint32,  值类型  *uint32,  缓冲区  *byte,  缓冲区长度  *uint32)  (注册表错误错误)  {}

// ff:
// key:
// name:
// reserved:
// valtype:
// buf:
// buflen:
// regerrno:
func RegQueryValueEx(key Handle, name *uint16, reserved *uint32, valtype *uint32, buf *byte, buflen *uint32) (regerrno error) {
	r0, _, _ := syscall.Syscall6(procRegQueryValueExW.Addr(), 6, uintptr(key), uintptr(unsafe.Pointer(name)), uintptr(unsafe.Pointer(reserved)), uintptr(unsafe.Pointer(valtype)), uintptr(unsafe.Pointer(buf)), uintptr(unsafe.Pointer(buflen)))
	if r0 != 0 {
		regerrno = syscall.Errno(r0)
	}
	return
}

// 翻译提示:func  注册事件源(服务器名称  *uint16,  源名称  *uint16)  (句柄  Handle,  错误  error)  {}

// ff:
// uncServerName:
// sourceName:
// handle:
// err:
func RegisterEventSource(uncServerName *uint16, sourceName *uint16) (handle Handle, err error) {
	r0, _, e1 := syscall.Syscall(procRegisterEventSourceW.Addr(), 2, uintptr(unsafe.Pointer(uncServerName)), uintptr(unsafe.Pointer(sourceName)), 0)
	handle = Handle(r0)
	if handle == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  注册服务控制处理器扩展(serviceName  *uint16,  处理器地址  uintptr,  上下文信息  uintptr)  (句柄  Handle,  错误  error)  {}

// ff:
// serviceName:
// handlerProc:
// context:
// handle:
// err:
func RegisterServiceCtrlHandlerEx(serviceName *uint16, handlerProc uintptr, context uintptr) (handle Handle, err error) {
	r0, _, e1 := syscall.Syscall(procRegisterServiceCtrlHandlerExW.Addr(), 3, uintptr(unsafe.Pointer(serviceName)), uintptr(handlerProc), uintptr(context))
	handle = Handle(r0)
	if handle == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  报告事件(log  手柄,  事件类型  uint16,  分类  uint16,  事件ID  uint32,  用户SID  uintptr,  字符串数量  uint16,  数据大小  uint32,  字符串指针  **uint16,  原始数据  *byte)  (错误  error)  {}

// ff:
// log:
// etype:
// category:
// eventId:
// usrSId:
// numStrings:
// dataSize:
// strings:
// rawData:
// err:
func ReportEvent(log Handle, etype uint16, category uint16, eventId uint32, usrSId uintptr, numStrings uint16, dataSize uint32, strings **uint16, rawData *byte) (err error) {
	r1, _, e1 := syscall.Syscall9(procReportEventW.Addr(), 9, uintptr(log), uintptr(etype), uintptr(category), uintptr(eventId), uintptr(usrSId), uintptr(numStrings), uintptr(dataSize), uintptr(unsafe.Pointer(strings)), uintptr(unsafe.Pointer(rawData)))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  自我还原()  (错误  error)  {}

// ff:
// err:
func RevertToSelf() (err error) {
	r1, _, e1 := syscall.Syscall(procRevertToSelf.Addr(), 0, 0, 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func setEntriesInAcl(countExplicitEntries uint32, explicitEntries *EXPLICIT_ACCESS, oldACL *ACL, newACL **ACL) (ret error) {
	r0, _, _ := syscall.Syscall6(procSetEntriesInAclW.Addr(), 4, uintptr(countExplicitEntries), uintptr(unsafe.Pointer(explicitEntries)), uintptr(unsafe.Pointer(oldACL)), uintptr(unsafe.Pointer(newACL)), 0, 0)
	if r0 != 0 {
		ret = syscall.Errno(r0)
	}
	return
}

// 翻译提示:func  设置内核对象安全属性(handle  文件句柄,  安全信息  SECURITY_INFORMATION,  安全描述符  *SECURITY_DESCRIPTOR)  (错误  error)  {}

// ff:
// handle:
// securityInformation:
// securityDescriptor:
// err:
func SetKernelObjectSecurity(handle Handle, securityInformation SECURITY_INFORMATION, securityDescriptor *SECURITY_DESCRIPTOR) (err error) {
	r1, _, e1 := syscall.Syscall(procSetKernelObjectSecurity.Addr(), 3, uintptr(handle), uintptr(securityInformation), uintptr(unsafe.Pointer(securityDescriptor)))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  设置命名安全信息(objectName  字符串,  objectType  对象类型,  securityInformation  安全信息类型,  owner  所有者SID,  group  组SID,  dacl  访问控制列表,  sacl  系统访问控制列表)  (返回值  错误)  {}

// ff:
// objectName:
// objectType:
// securityInformation:
// owner:
// group:
// dacl:
// sacl:
// ret:
func SetNamedSecurityInfo(objectName string, objectType SE_OBJECT_TYPE, securityInformation SECURITY_INFORMATION, owner *SID, group *SID, dacl *ACL, sacl *ACL) (ret error) {
	var _p0 *uint16
	_p0, ret = syscall.UTF16PtrFromString(objectName)
	if ret != nil {
		return
	}
	return _SetNamedSecurityInfo(_p0, objectType, securityInformation, owner, group, dacl, sacl)
}

func _SetNamedSecurityInfo(objectName *uint16, objectType SE_OBJECT_TYPE, securityInformation SECURITY_INFORMATION, owner *SID, group *SID, dacl *ACL, sacl *ACL) (ret error) {
	r0, _, _ := syscall.Syscall9(procSetNamedSecurityInfoW.Addr(), 7, uintptr(unsafe.Pointer(objectName)), uintptr(objectType), uintptr(securityInformation), uintptr(unsafe.Pointer(owner)), uintptr(unsafe.Pointer(group)), uintptr(unsafe.Pointer(dacl)), uintptr(unsafe.Pointer(sacl)), 0, 0)
	if r0 != 0 {
		ret = syscall.Errno(r0)
	}
	return
}

func setSecurityDescriptorControl(sd *SECURITY_DESCRIPTOR, controlBitsOfInterest SECURITY_DESCRIPTOR_CONTROL, controlBitsToSet SECURITY_DESCRIPTOR_CONTROL) (err error) {
	r1, _, e1 := syscall.Syscall(procSetSecurityDescriptorControl.Addr(), 3, uintptr(unsafe.Pointer(sd)), uintptr(controlBitsOfInterest), uintptr(controlBitsToSet))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func setSecurityDescriptorDacl(sd *SECURITY_DESCRIPTOR, daclPresent bool, dacl *ACL, daclDefaulted bool) (err error) {
	var _p0 uint32
	if daclPresent {
		_p0 = 1
	}
	var _p1 uint32
	if daclDefaulted {
		_p1 = 1
	}
	r1, _, e1 := syscall.Syscall6(procSetSecurityDescriptorDacl.Addr(), 4, uintptr(unsafe.Pointer(sd)), uintptr(_p0), uintptr(unsafe.Pointer(dacl)), uintptr(_p1), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func setSecurityDescriptorGroup(sd *SECURITY_DESCRIPTOR, group *SID, groupDefaulted bool) (err error) {
	var _p0 uint32
	if groupDefaulted {
		_p0 = 1
	}
	r1, _, e1 := syscall.Syscall(procSetSecurityDescriptorGroup.Addr(), 3, uintptr(unsafe.Pointer(sd)), uintptr(unsafe.Pointer(group)), uintptr(_p0))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func setSecurityDescriptorOwner(sd *SECURITY_DESCRIPTOR, owner *SID, ownerDefaulted bool) (err error) {
	var _p0 uint32
	if ownerDefaulted {
		_p0 = 1
	}
	r1, _, e1 := syscall.Syscall(procSetSecurityDescriptorOwner.Addr(), 3, uintptr(unsafe.Pointer(sd)), uintptr(unsafe.Pointer(owner)), uintptr(_p0))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func setSecurityDescriptorRMControl(sd *SECURITY_DESCRIPTOR, rmControl *uint8) {
	syscall.Syscall(procSetSecurityDescriptorRMControl.Addr(), 2, uintptr(unsafe.Pointer(sd)), uintptr(unsafe.Pointer(rmControl)), 0)
	return
}

func setSecurityDescriptorSacl(sd *SECURITY_DESCRIPTOR, saclPresent bool, sacl *ACL, saclDefaulted bool) (err error) {
	var _p0 uint32
	if saclPresent {
		_p0 = 1
	}
	var _p1 uint32
	if saclDefaulted {
		_p1 = 1
	}
	r1, _, e1 := syscall.Syscall6(procSetSecurityDescriptorSacl.Addr(), 4, uintptr(unsafe.Pointer(sd)), uintptr(_p0), uintptr(unsafe.Pointer(sacl)), uintptr(_p1), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  设置安全信息(handle  手柄,  objectType  对象类型,  securityInformation  安全信息类型,  owner  所有者  SID,  group  组  SID,  dacl  访问控制列表,  sacl  系统访问控制列表)  (错误  error)  {}

// ff:
// handle:
// objectType:
// securityInformation:
// owner:
// group:
// dacl:
// sacl:
// ret:
func SetSecurityInfo(handle Handle, objectType SE_OBJECT_TYPE, securityInformation SECURITY_INFORMATION, owner *SID, group *SID, dacl *ACL, sacl *ACL) (ret error) {
	r0, _, _ := syscall.Syscall9(procSetSecurityInfo.Addr(), 7, uintptr(handle), uintptr(objectType), uintptr(securityInformation), uintptr(unsafe.Pointer(owner)), uintptr(unsafe.Pointer(group)), uintptr(unsafe.Pointer(dacl)), uintptr(unsafe.Pointer(sacl)), 0, 0)
	if r0 != 0 {
		ret = syscall.Errno(r0)
	}
	return
}

// 翻译提示:func  设置服务状态(service  服务句柄,  serviceStatus  服务状态指针)  (错误  error)  {}

// ff:
// service:
// serviceStatus:
// err:
func SetServiceStatus(service Handle, serviceStatus *SERVICE_STATUS) (err error) {
	r1, _, e1 := syscall.Syscall(procSetServiceStatus.Addr(), 2, uintptr(service), uintptr(unsafe.Pointer(serviceStatus)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  设置线程令牌(thread  *进程句柄,  token  令牌)  (错误  error)  {}

// ff:
// thread:
// token:
// err:
func SetThreadToken(thread *Handle, token Token) (err error) {
	r1, _, e1 := syscall.Syscall(procSetThreadToken.Addr(), 2, uintptr(unsafe.Pointer(thread)), uintptr(token), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  设置令牌信息(token  令牌,  信息类别  uint32,  信息  *字节,  信息长度  uint32)  (错误  错误)  {}

// ff:
// token:
// infoClass:
// info:
// infoLen:
// err:
func SetTokenInformation(token Token, infoClass uint32, info *byte, infoLen uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procSetTokenInformation.Addr(), 4, uintptr(token), uintptr(infoClass), uintptr(unsafe.Pointer(info)), uintptr(infoLen), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  启动服务控制调度器(serviceTable  *服务表条目)  (错误  error)  {}

// ff:
// serviceTable:
// err:
func StartServiceCtrlDispatcher(serviceTable *SERVICE_TABLE_ENTRY) (err error) {
	r1, _, e1 := syscall.Syscall(procStartServiceCtrlDispatcherW.Addr(), 1, uintptr(unsafe.Pointer(serviceTable)), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  开始服务(service  接口句柄,  参数数量  uint32,  参数向量  **uint16)  (错误  error)  {}

// ff:
// service:
// numArgs:
// argVectors:
// err:
func StartService(service Handle, numArgs uint32, argVectors **uint16) (err error) {
	r1, _, e1 := syscall.Syscall(procStartServiceW.Addr(), 3, uintptr(service), uintptr(numArgs), uintptr(unsafe.Pointer(argVectors)))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  添加证书上下文到存储(store  手柄,  证书上下文  *证书上下文,  添加处置  uint32,  存储上下文  **证书上下文)  (错误  error)  {}

// ff:
// store:
// certContext:
// addDisposition:
// storeContext:
// err:
func CertAddCertificateContextToStore(store Handle, certContext *CertContext, addDisposition uint32, storeContext **CertContext) (err error) {
	r1, _, e1 := syscall.Syscall6(procCertAddCertificateContextToStore.Addr(), 4, uintptr(store), uintptr(unsafe.Pointer(certContext)), uintptr(addDisposition), uintptr(unsafe.Pointer(storeContext)), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  关闭证书存储(store  手柄,  标志  uint32)  (错误  error)  {}

// ff:
// store:
// flags:
// err:
func CertCloseStore(store Handle, flags uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procCertCloseStore.Addr(), 2, uintptr(store), uintptr(flags), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  创建证书上下文(cert编码类型  uint32,  证书编码  *byte,  编码长度  uint32)  (证书上下文  *CertContext,  错误  error)  {}

// ff:
// certEncodingType:
// certEncoded:
// encodedLen:
// context:
// err:
func CertCreateCertificateContext(certEncodingType uint32, certEncoded *byte, encodedLen uint32) (context *CertContext, err error) {
	r0, _, e1 := syscall.Syscall(procCertCreateCertificateContext.Addr(), 3, uintptr(certEncodingType), uintptr(unsafe.Pointer(certEncoded)), uintptr(encodedLen))
	context = (*CertContext)(unsafe.Pointer(r0))
	if context == nil {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  删除证书从存储区(certContext  *证书上下文)  (错误  error)  {}

// ff:
// certContext:
// err:
func CertDeleteCertificateFromStore(certContext *CertContext) (err error) {
	r1, _, e1 := syscall.Syscall(procCertDeleteCertificateFromStore.Addr(), 1, uintptr(unsafe.Pointer(certContext)), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  复制证书上下文(certContext  *证书上下文)  (dupContext  *证书上下文)  {}

// ff:
// certContext:
// dupContext:
func CertDuplicateCertificateContext(certContext *CertContext) (dupContext *CertContext) {
	r0, _, _ := syscall.Syscall(procCertDuplicateCertificateContext.Addr(), 1, uintptr(unsafe.Pointer(certContext)), 0, 0)
	dupContext = (*CertContext)(unsafe.Pointer(r0))
	return
}

// 翻译提示:func  证书遍历存储区(store  手柄,  前一个上下文  *证书上下文)  (*证书上下文,  错误)  {
//         var  context  *证书上下文
//         err  :=  CertEnumCertificatesInStore(store,  prevContext)
//         return  context,  err
// }

// ff:
// store:
// prevContext:
// context:
// err:
func CertEnumCertificatesInStore(store Handle, prevContext *CertContext) (context *CertContext, err error) {
	r0, _, e1 := syscall.Syscall(procCertEnumCertificatesInStore.Addr(), 2, uintptr(store), uintptr(unsafe.Pointer(prevContext)), 0)
	context = (*CertContext)(unsafe.Pointer(r0))
	if context == nil {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  证书在存储中查找(store  手柄,  编码类型  uint32,  查找标志  uint32,  查找类型  uint32,  查找参数  unsafe.Pointer,  前一个证书上下文  *CertContext)  (*CertContext,  error)  {
//         var  cert  *CertContext
//         _,  err  :=  certFindCertificateInStore(store,  编码类型,  查找标志,  查找类型,  查找参数,  prevCertContext,  &cert)
//         return  cert,  err
// }
// 
// //  原型声明，用于内部调用
// func  certFindCertificateInStore(store  手柄,  encodingType  uint32,  findFlags  uint32,  findType  uint32,  findPara  unsafe.Pointer,  prevCertContext  *CertContext,  certContext  **CertContext)  (bool,  error)  {}

// ff:
// store:
// certEncodingType:
// findFlags:
// findType:
// findPara:
// prevCertContext:
// cert:
// err:
func CertFindCertificateInStore(store Handle, certEncodingType uint32, findFlags uint32, findType uint32, findPara unsafe.Pointer, prevCertContext *CertContext) (cert *CertContext, err error) {
	r0, _, e1 := syscall.Syscall6(procCertFindCertificateInStore.Addr(), 6, uintptr(store), uintptr(certEncodingType), uintptr(findFlags), uintptr(findType), uintptr(findPara), uintptr(unsafe.Pointer(prevCertContext)))
	cert = (*CertContext)(unsafe.Pointer(r0))
	if cert == nil {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  查找证书链在存储区(ContextHandle  store,  证书编码类型  CertificateEncodingType,  查找标志  FindFlags,  查找类型  FindType,  查找参数  FindParameters,  上一个证书链上下文  *CertChainContext)  (*CertChainContext,  error)  {}

// ff:
// store:
// certEncodingType:
// findFlags:
// findType:
// findPara:
// prevChainContext:
// certchain:
// err:
func CertFindChainInStore(store Handle, certEncodingType uint32, findFlags uint32, findType uint32, findPara unsafe.Pointer, prevChainContext *CertChainContext) (certchain *CertChainContext, err error) {
	r0, _, e1 := syscall.Syscall6(procCertFindChainInStore.Addr(), 6, uintptr(store), uintptr(certEncodingType), uintptr(findFlags), uintptr(findType), uintptr(findPara), uintptr(unsafe.Pointer(prevChainContext)))
	certchain = (*CertChainContext)(unsafe.Pointer(r0))
	if certchain == nil {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  查找证书扩展(objId  *字节,  扩展数量  uint32,  扩展列表  *CertExtension)  (找到的扩展  *CertExtension)  {}

// ff:
// objId:
// countExtensions:
// extensions:
// ret:
func CertFindExtension(objId *byte, countExtensions uint32, extensions *CertExtension) (ret *CertExtension) {
	r0, _, _ := syscall.Syscall(procCertFindExtension.Addr(), 3, uintptr(unsafe.Pointer(objId)), uintptr(countExtensions), uintptr(unsafe.Pointer(extensions)))
	ret = (*CertExtension)(unsafe.Pointer(r0))
	return
}

// 翻译提示:func  释放证书链(ctx  *证书链上下文)  {}

// ff:
// ctx:
func CertFreeCertificateChain(ctx *CertChainContext) {
	syscall.Syscall(procCertFreeCertificateChain.Addr(), 1, uintptr(unsafe.Pointer(ctx)), 0, 0)
	return
}

// 翻译提示:func  释放证书上下文(ctx  *证书上下文)  (错误  error)  {}

// ff:
// ctx:
// err:
func CertFreeCertificateContext(ctx *CertContext) (err error) {
	r1, _, e1 := syscall.Syscall(procCertFreeCertificateContext.Addr(), 1, uintptr(unsafe.Pointer(ctx)), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  获取证书链引擎(handle  手柄,  证书  *证书上下文,  时间  *文件时间,  补充存储  手柄,  链参数  *证书链参数,  标志  uint32,  保留  uintptr,  链上下文  **证书链上下文)  (错误  error)  {}

// ff:
// engine:
// leaf:
// time:
// additionalStore:
// para:
// flags:
// reserved:
// chainCtx:
// err:
func CertGetCertificateChain(engine Handle, leaf *CertContext, time *Filetime, additionalStore Handle, para *CertChainPara, flags uint32, reserved uintptr, chainCtx **CertChainContext) (err error) {
	r1, _, e1 := syscall.Syscall9(procCertGetCertificateChain.Addr(), 8, uintptr(engine), uintptr(unsafe.Pointer(leaf)), uintptr(unsafe.Pointer(time)), uintptr(additionalStore), uintptr(unsafe.Pointer(para)), uintptr(flags), uintptr(reserved), uintptr(unsafe.Pointer(chainCtx)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  证书获取名称字符串(certContext  *证书上下文,  名称类型  uint32,  标志  uint32,  类型参数  unsafe.Pointer,  名称  *uint16,  大小  uint32)  (字符数  uint32)  {}

// ff:
// certContext:
// nameType:
// flags:
// typePara:
// name:
// size:
// chars:
func CertGetNameString(certContext *CertContext, nameType uint32, flags uint32, typePara unsafe.Pointer, name *uint16, size uint32) (chars uint32) {
	r0, _, _ := syscall.Syscall6(procCertGetNameStringW.Addr(), 6, uintptr(unsafe.Pointer(certContext)), uintptr(nameType), uintptr(flags), uintptr(typePara), uintptr(unsafe.Pointer(name)), uintptr(size))
	chars = uint32(r0)
	return
}

// 翻译提示:func  打开证书存储区(证书存储提供者  uintptr,  消息和证书编码类型  uint32,  加密提供者  uintptr,  标志  uint32,  参数  uintptr)  (句柄  Handle,  错误  error)  {}

// ff:
// storeProvider:
// msgAndCertEncodingType:
// cryptProv:
// flags:
// para:
// handle:
// err:
func CertOpenStore(storeProvider uintptr, msgAndCertEncodingType uint32, cryptProv uintptr, flags uint32, para uintptr) (handle Handle, err error) {
	r0, _, e1 := syscall.Syscall6(procCertOpenStore.Addr(), 5, uintptr(storeProvider), uintptr(msgAndCertEncodingType), uintptr(cryptProv), uintptr(flags), uintptr(para), 0)
	handle = Handle(r0)
	if handle == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  打开系统证书存储(hprov  操作句柄,  名称  *uint16)  (证书存储句柄  Handle,  错误  error)  {}

// ff:
// hprov:
// name:
// store:
// err:
func CertOpenSystemStore(hprov Handle, name *uint16) (store Handle, err error) {
	r0, _, e1 := syscall.Syscall(procCertOpenSystemStoreW.Addr(), 2, uintptr(hprov), uintptr(unsafe.Pointer(name)), 0)
	store = Handle(r0)
	if store == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  验证证书链策略(policyOID  uintptr,  证书链  *证书链上下文,  参数  *证书链策略参数,  状态  *证书链策略状态)  (错误  error)  {}

// ff:
// policyOID:
// chain:
// para:
// status:
// err:
func CertVerifyCertificateChainPolicy(policyOID uintptr, chain *CertChainContext, para *CertChainPolicyPara, status *CertChainPolicyStatus) (err error) {
	r1, _, e1 := syscall.Syscall6(procCertVerifyCertificateChainPolicy.Addr(), 4, uintptr(policyOID), uintptr(unsafe.Pointer(chain)), uintptr(unsafe.Pointer(para)), uintptr(unsafe.Pointer(status)), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示://  加密获取证书私钥
// func  CryptAcquireCertificatePrivateKey(cert  *CertContext,  flags  uint32,  parameters  unsafe.Pointer,  cryptProvOrNCryptKey  *Handle,  keySpec  *uint32,  callerFreeProvOrNCryptKey  *bool)  (err  error)  {
//         //  ...函数实现...
// }

// ff:
// cert:
// flags:
// parameters:
// cryptProvOrNCryptKey:
// keySpec:
// callerFreeProvOrNCryptKey:
// err:
func CryptAcquireCertificatePrivateKey(cert *CertContext, flags uint32, parameters unsafe.Pointer, cryptProvOrNCryptKey *Handle, keySpec *uint32, callerFreeProvOrNCryptKey *bool) (err error) {
	var _p0 uint32
	if *callerFreeProvOrNCryptKey {
		_p0 = 1
	}
	r1, _, e1 := syscall.Syscall6(procCryptAcquireCertificatePrivateKey.Addr(), 6, uintptr(unsafe.Pointer(cert)), uintptr(flags), uintptr(parameters), uintptr(unsafe.Pointer(cryptProvOrNCryptKey)), uintptr(unsafe.Pointer(keySpec)), uintptr(unsafe.Pointer(&_p0)))
	*callerFreeProvOrNCryptKey = _p0 != 0
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  加密解码对象(encodingType  uint32,  结构类型  *byte,  编码字节  *byte,  编码字节长度  uint32,  标志  uint32,  解码后数据  unsafe.Pointer,  解码后数据长度  *uint32)  (错误  error)  {}

// ff:
// encodingType:
// structType:
// encodedBytes:
// lenEncodedBytes:
// flags:
// decoded:
// decodedLen:
// err:
func CryptDecodeObject(encodingType uint32, structType *byte, encodedBytes *byte, lenEncodedBytes uint32, flags uint32, decoded unsafe.Pointer, decodedLen *uint32) (err error) {
	r1, _, e1 := syscall.Syscall9(procCryptDecodeObject.Addr(), 7, uintptr(encodingType), uintptr(unsafe.Pointer(structType)), uintptr(unsafe.Pointer(encodedBytes)), uintptr(lenEncodedBytes), uintptr(flags), uintptr(decoded), uintptr(unsafe.Pointer(decodedLen)), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  加密保护数据(dataIn  *数据块,  名称  *uint16,  可选熵  *数据块,  保留  uintptr,  提示结构  *加密保护提示结构,  标志  uint32,  dataOut  *数据块)  (错误  error)  {}  
// 
// 这里的方法名称和参数翻译如下：
// 
// -  CryptProtectData  ->  加密保护数据
// -  dataIn  ->  输入数据块
// -  name  ->  名称
// -  optionalEntropy  ->  可选熵数据块
// -  reserved  ->  保留参数
// -  promptStruct  ->  加密保护提示结构
// -  flags  ->  标志位
// -  dataOut  ->  输出数据块
// -  err  ->  错误

// ff:
// dataIn:
// name:
// optionalEntropy:
// reserved:
// promptStruct:
// flags:
// dataOut:
// err:
func CryptProtectData(dataIn *DataBlob, name *uint16, optionalEntropy *DataBlob, reserved uintptr, promptStruct *CryptProtectPromptStruct, flags uint32, dataOut *DataBlob) (err error) {
	r1, _, e1 := syscall.Syscall9(procCryptProtectData.Addr(), 7, uintptr(unsafe.Pointer(dataIn)), uintptr(unsafe.Pointer(name)), uintptr(unsafe.Pointer(optionalEntropy)), uintptr(reserved), uintptr(unsafe.Pointer(promptStruct)), uintptr(flags), uintptr(unsafe.Pointer(dataOut)), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  加密查询对象(objectType  uint32,  对象  unsafe.Pointer,  预期内容类型标志  uint32,  预期格式类型标志  uint32,  标志  uint32,  消息和证书编码类型  *uint32,  内容类型  *uint32,  格式类型  *uint32,  证书存储  *Handle,  消息  *Handle,  上下文  *unsafe.Pointer)  (错误  error)  {}

// ff:
// objectType:
// object:
// expectedContentTypeFlags:
// expectedFormatTypeFlags:
// flags:
// msgAndCertEncodingType:
// contentType:
// formatType:
// certStore:
// msg:
// context:
// err:
func CryptQueryObject(objectType uint32, object unsafe.Pointer, expectedContentTypeFlags uint32, expectedFormatTypeFlags uint32, flags uint32, msgAndCertEncodingType *uint32, contentType *uint32, formatType *uint32, certStore *Handle, msg *Handle, context *unsafe.Pointer) (err error) {
	r1, _, e1 := syscall.Syscall12(procCryptQueryObject.Addr(), 11, uintptr(objectType), uintptr(object), uintptr(expectedContentTypeFlags), uintptr(expectedFormatTypeFlags), uintptr(flags), uintptr(unsafe.Pointer(msgAndCertEncodingType)), uintptr(unsafe.Pointer(contentType)), uintptr(unsafe.Pointer(formatType)), uintptr(unsafe.Pointer(certStore)), uintptr(unsafe.Pointer(msg)), uintptr(unsafe.Pointer(context)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  解密数据(dataIn  *数据块,  name  **uint16,  可选熵  *数据块,  保留  uintptr,  提示结构体  *加密保护提示结构,  标志  uint32,  dataOut  *数据块)  (err  error)  {}

// ff:
// dataIn:
// name:
// optionalEntropy:
// reserved:
// promptStruct:
// flags:
// dataOut:
// err:
func CryptUnprotectData(dataIn *DataBlob, name **uint16, optionalEntropy *DataBlob, reserved uintptr, promptStruct *CryptProtectPromptStruct, flags uint32, dataOut *DataBlob) (err error) {
	r1, _, e1 := syscall.Syscall9(procCryptUnprotectData.Addr(), 7, uintptr(unsafe.Pointer(dataIn)), uintptr(unsafe.Pointer(name)), uintptr(unsafe.Pointer(optionalEntropy)), uintptr(reserved), uintptr(unsafe.Pointer(promptStruct)), uintptr(flags), uintptr(unsafe.Pointer(dataOut)), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  导入PFX证书存储(pfx  *密码数据块,  密码  *uint16,  标志  uint32)  (存储句柄  Handle,  错误  error)  {}

// ff:
// pfx:
// password:
// flags:
// store:
// err:
func PFXImportCertStore(pfx *CryptDataBlob, password *uint16, flags uint32) (store Handle, err error) {
	r0, _, e1 := syscall.Syscall(procPFXImportCertStore.Addr(), 3, uintptr(unsafe.Pointer(pfx)), uintptr(unsafe.Pointer(password)), uintptr(flags))
	store = Handle(r0)
	if store == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  域名比较(name1  *字符串指针,  name2  *字符串指针)  (相同  bool)  {}

// ff:
// name1:
// name2:
// same:
func DnsNameCompare(name1 *uint16, name2 *uint16) (same bool) {
	r0, _, _ := syscall.Syscall(procDnsNameCompare_W.Addr(), 2, uintptr(unsafe.Pointer(name1)), uintptr(unsafe.Pointer(name2)), 0)
	same = r0 != 0
	return
}

// 翻译提示:func  DNS查询(name  string,  类型  uint16,  选项  uint32,  额外信息  *byte,  记录  **DNSRecord,  应答  *byte)  (状态  error)  {}

// ff:
// name:
// qtype:
// options:
// extra:
// qrs:
// pr:
// status:
func DnsQuery(name string, qtype uint16, options uint32, extra *byte, qrs **DNSRecord, pr *byte) (status error) {
	var _p0 *uint16
	_p0, status = syscall.UTF16PtrFromString(name)
	if status != nil {
		return
	}
	return _DnsQuery(_p0, qtype, options, extra, qrs, pr)
}

func _DnsQuery(name *uint16, qtype uint16, options uint32, extra *byte, qrs **DNSRecord, pr *byte) (status error) {
	r0, _, _ := syscall.Syscall6(procDnsQuery_W.Addr(), 6, uintptr(unsafe.Pointer(name)), uintptr(qtype), uintptr(options), uintptr(unsafe.Pointer(extra)), uintptr(unsafe.Pointer(qrs)), uintptr(unsafe.Pointer(pr)))
	if r0 != 0 {
		status = syscall.Errno(r0)
	}
	return
}

// 翻译提示:func  Dns记录列表释放(rl  *DnsRecord,  释放类型  uint32)  {}

// ff:
// rl:
// freetype:
func DnsRecordListFree(rl *DNSRecord, freetype uint32) {
	syscall.Syscall(procDnsRecordListFree.Addr(), 2, uintptr(unsafe.Pointer(rl)), uintptr(freetype), 0)
	return
}

// 翻译提示:func  Dwm获取窗口属性(hwnd  窗口句柄,  属性  uint32,  值  unsafe.Pointer,  大小  uint32)  (返回值  error)  {}

// ff:
// hwnd:
// attribute:
// value:
// size:
// ret:
func DwmGetWindowAttribute(hwnd HWND, attribute uint32, value unsafe.Pointer, size uint32) (ret error) {
	r0, _, _ := syscall.Syscall6(procDwmGetWindowAttribute.Addr(), 4, uintptr(hwnd), uintptr(attribute), uintptr(value), uintptr(size), 0, 0)
	if r0 != 0 {
		ret = syscall.Errno(r0)
	}
	return
}

// 翻译提示:func  设置Dwm窗口属性(hwnd  窗口句柄,  属性  uint32,  值  unsafe.Pointer,  大小  uint32)  (错误  error)  {}

// ff:
// hwnd:
// attribute:
// value:
// size:
// ret:
func DwmSetWindowAttribute(hwnd HWND, attribute uint32, value unsafe.Pointer, size uint32) (ret error) {
	r0, _, _ := syscall.Syscall6(procDwmSetWindowAttribute.Addr(), 4, uintptr(hwnd), uintptr(attribute), uintptr(value), uintptr(size), 0, 0)
	if r0 != 0 {
		ret = syscall.Errno(r0)
	}
	return
}

// 翻译提示:func  获取网络适配器地址(family  网络家族类型,  flags  标志,  reserved  保留参数,  adapterAddresses  网络适配器地址结构体指针,  sizePointer  字符串大小指针)  (错误码  error)  {}

// ff:
// family:
// flags:
// reserved:
// adapterAddresses:
// sizePointer:
// errcode:
func GetAdaptersAddresses(family uint32, flags uint32, reserved uintptr, adapterAddresses *IpAdapterAddresses, sizePointer *uint32) (errcode error) {
	r0, _, _ := syscall.Syscall6(procGetAdaptersAddresses.Addr(), 5, uintptr(family), uintptr(flags), uintptr(reserved), uintptr(unsafe.Pointer(adapterAddresses)), uintptr(unsafe.Pointer(sizePointer)), 0)
	if r0 != 0 {
		errcode = syscall.Errno(r0)
	}
	return
}

// 翻译提示:func  获取网络适配器信息(adapterInfo  *网络适配器信息,  outLength  *uint32)  (错误码  error)  {}

// ff:
// ai:
// ol:
// errcode:
func GetAdaptersInfo(ai *IpAdapterInfo, ol *uint32) (errcode error) {
	r0, _, _ := syscall.Syscall(procGetAdaptersInfo.Addr(), 2, uintptr(unsafe.Pointer(ai)), uintptr(unsafe.Pointer(ol)), 0)
	if r0 != 0 {
		errcode = syscall.Errno(r0)
	}
	return
}

func getBestInterfaceEx(sockaddr unsafe.Pointer, pdwBestIfIndex *uint32) (errcode error) {
	r0, _, _ := syscall.Syscall(procGetBestInterfaceEx.Addr(), 2, uintptr(sockaddr), uintptr(unsafe.Pointer(pdwBestIfIndex)), 0)
	if r0 != 0 {
		errcode = syscall.Errno(r0)
	}
	return
}

// 翻译提示:func  获取接口条目(pIfRow  *Mib接口行)  (错误码  error)  {}

// ff:
// pIfRow:
// errcode:
func GetIfEntry(pIfRow *MibIfRow) (errcode error) {
	r0, _, _ := syscall.Syscall(procGetIfEntry.Addr(), 1, uintptr(unsafe.Pointer(pIfRow)), 0, 0)
	if r0 != 0 {
		errcode = syscall.Errno(r0)
	}
	return
}

// 翻译提示:func  添加Dll目录(path  *uint16)  (cookie  uintptr,  错误  error)  {}

// ff:
// path:
// cookie:
// err:
func AddDllDirectory(path *uint16) (cookie uintptr, err error) {
	r0, _, e1 := syscall.Syscall(procAddDllDirectory.Addr(), 1, uintptr(unsafe.Pointer(path)), 0, 0)
	cookie = uintptr(r0)
	if cookie == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  将进程分配到作业对象(jobHandle  手柄,  processHandle  进程手柄)  (错误  error)  {}

// ff:
// job:
// process:
// err:
func AssignProcessToJobObject(job Handle, process Handle) (err error) {
	r1, _, e1 := syscall.Syscall(procAssignProcessToJobObject.Addr(), 2, uintptr(job), uintptr(process), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  取消Io(s  文件句柄)  (错误  error)  {}

// ff:
// s:
// err:
func CancelIo(s Handle) (err error) {
	r1, _, e1 := syscall.Syscall(procCancelIo.Addr(), 1, uintptr(s), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  取消Io扩展(s  文件句柄,  o  超越重叠)  (错误  错误)  {}

// ff:
// s:
// o:
// err:
func CancelIoEx(s Handle, o *Overlapped) (err error) {
	r1, _, e1 := syscall.Syscall(procCancelIoEx.Addr(), 2, uintptr(s), uintptr(unsafe.Pointer(o)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  清除串行通信中断(handle  文件句柄)  (错误  error)  {}

// ff:
// handle:
// err:
func ClearCommBreak(handle Handle) (err error) {
	r1, _, e1 := syscall.Syscall(procClearCommBreak.Addr(), 1, uintptr(handle), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  清除串行通信错误(handle  设备句柄,  lpErrors  错误指针,  lpStat  通信状态指针)  (err  错误)  {}

// ff:
// handle:
// lpErrors:
// lpStat:
// err:
func ClearCommError(handle Handle, lpErrors *uint32, lpStat *ComStat) (err error) {
	r1, _, e1 := syscall.Syscall(procClearCommError.Addr(), 3, uintptr(handle), uintptr(unsafe.Pointer(lpErrors)), uintptr(unsafe.Pointer(lpStat)))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  关闭句柄(handle  句柄)  (错误  error)  {}

// ff:
// handle:
// err:
func CloseHandle(handle Handle) (err error) {
	r1, _, e1 := syscall.Syscall(procCloseHandle.Addr(), 1, uintptr(handle), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  关闭伪控制台(console  控制台句柄)  {}

// ff:
// console:
func ClosePseudoConsole(console Handle) {
	syscall.Syscall(procClosePseudoConsole.Addr(), 1, uintptr(console), 0, 0)
	return
}

// 翻译提示:func  连接命名管道(pipe  文件句柄,  overlapped  超时覆盖结构体)  (错误  错误)  {}

// ff:
// pipe:
// overlapped:
// err:
func ConnectNamedPipe(pipe Handle, overlapped *Overlapped) (err error) {
	r1, _, e1 := syscall.Syscall(procConnectNamedPipe.Addr(), 2, uintptr(pipe), uintptr(unsafe.Pointer(overlapped)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  创建目录(path  *uint16,  安全属性  *SecurityAttributes)  (错误  error)  {}

// ff:
// path:
// sa:
// err:
func CreateDirectory(path *uint16, sa *SecurityAttributes) (err error) {
	r1, _, e1 := syscall.Syscall(procCreateDirectoryW.Addr(), 2, uintptr(unsafe.Pointer(path)), uintptr(unsafe.Pointer(sa)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  创建事件Ex(event安全属性*SecurityAttributes,  名称*uint16,  标志uint32,  所需访问权限uint32)  (句柄Handle,  错误error)  {}

// ff:
// eventAttrs:
// name:
// flags:
// desiredAccess:
// handle:
// err:
func CreateEventEx(eventAttrs *SecurityAttributes, name *uint16, flags uint32, desiredAccess uint32) (handle Handle, err error) {
	r0, _, e1 := syscall.Syscall6(procCreateEventExW.Addr(), 4, uintptr(unsafe.Pointer(eventAttrs)), uintptr(unsafe.Pointer(name)), uintptr(flags), uintptr(desiredAccess), 0, 0)
	handle = Handle(r0)
	if handle == 0 || e1 == ERROR_ALREADY_EXISTS {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  创建事件(securityAttrs  *安全属性,  manualReset  是否手动重置,  initialState  初始状态,  name  *宽字符字符串)  (句柄  Handle,  错误  error)  {}

// ff:
// eventAttrs:
// manualReset:
// initialState:
// name:
// handle:
// err:
func CreateEvent(eventAttrs *SecurityAttributes, manualReset uint32, initialState uint32, name *uint16) (handle Handle, err error) {
	r0, _, e1 := syscall.Syscall6(procCreateEventW.Addr(), 4, uintptr(unsafe.Pointer(eventAttrs)), uintptr(manualReset), uintptr(initialState), uintptr(unsafe.Pointer(name)), 0, 0)
	handle = Handle(r0)
	if handle == 0 || e1 == ERROR_ALREADY_EXISTS {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  创建文件映射(fhandle  文件句柄,  sa  安全属性指针,  prot  保护模式  uint32,  maxSizeHigh  最大尺寸高32位  uint32,  maxSizeLow  最大尺寸低32位  uint32,  name  名称指针  *uint16)  (handle  文件映射句柄,  err  错误)  {}

// ff:
// fhandle:
// sa:
// prot:
// maxSizeHigh:
// maxSizeLow:
// name:
// handle:
// err:
func CreateFileMapping(fhandle Handle, sa *SecurityAttributes, prot uint32, maxSizeHigh uint32, maxSizeLow uint32, name *uint16) (handle Handle, err error) {
	r0, _, e1 := syscall.Syscall6(procCreateFileMappingW.Addr(), 6, uintptr(fhandle), uintptr(unsafe.Pointer(sa)), uintptr(prot), uintptr(maxSizeHigh), uintptr(maxSizeLow), uintptr(unsafe.Pointer(name)))
	handle = Handle(r0)
	if handle == 0 || e1 == ERROR_ALREADY_EXISTS {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  创建文件(name  *uint16,  访问权限  uint32,  模式  uint32,  安全属性  *SecurityAttributes,  创建模式  uint32,  属性  uint32,  模板文件  Handle)  (句柄  Handle,  错误  error)  {}

// ff:
// name:
// access:
// mode:
// sa:
// createmode:
// attrs:
// templatefile:
// handle:
// err:
func CreateFile(name *uint16, access uint32, mode uint32, sa *SecurityAttributes, createmode uint32, attrs uint32, templatefile Handle) (handle Handle, err error) {
	r0, _, e1 := syscall.Syscall9(procCreateFileW.Addr(), 7, uintptr(unsafe.Pointer(name)), uintptr(access), uintptr(mode), uintptr(unsafe.Pointer(sa)), uintptr(createmode), uintptr(attrs), uintptr(templatefile), 0, 0)
	handle = Handle(r0)
	if handle == InvalidHandle {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  创建硬链接(newFilename  *uint16,  existingFilename  *uint16,  reserved  uintptr)  (错误  error)  {}

// ff:
// filename:
// existingfilename:
// reserved:
// err:
func CreateHardLink(filename *uint16, existingfilename *uint16, reserved uintptr) (err error) {
	r1, _, e1 := syscall.Syscall(procCreateHardLinkW.Addr(), 3, uintptr(unsafe.Pointer(filename)), uintptr(unsafe.Pointer(existingfilename)), uintptr(reserved))
	if r1&0xff == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  创建IO完成端口(filehandle  文件句柄,  cphandle  IO完成端口句柄,  key  关键字,  threadcnt  线程数)  (新IO完成端口句柄  文件句柄,  错误  error)  {}

// ff:
// filehandle:
// cphandle:
// key:
// threadcnt:
// handle:
// err:
func CreateIoCompletionPort(filehandle Handle, cphandle Handle, key uintptr, threadcnt uint32) (handle Handle, err error) {
	r0, _, e1 := syscall.Syscall6(procCreateIoCompletionPort.Addr(), 4, uintptr(filehandle), uintptr(cphandle), uintptr(key), uintptr(threadcnt), 0, 0)
	handle = Handle(r0)
	if handle == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  创建作业对象(jobAttr  安全属性,  name  *宽字符)  (句柄  手柄,  错误  错误)  {}

// ff:
// jobAttr:
// name:
// handle:
// err:
func CreateJobObject(jobAttr *SecurityAttributes, name *uint16) (handle Handle, err error) {
	r0, _, e1 := syscall.Syscall(procCreateJobObjectW.Addr(), 2, uintptr(unsafe.Pointer(jobAttr)), uintptr(unsafe.Pointer(name)), 0)
	handle = Handle(r0)
	if handle == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  创建互斥体Ex(mutexAttrs  *安全属性,  name  *uint16,  标志  uint32,  所需权限  uint32)  (句柄  Handle,  错误  error)  {}

// ff:
// mutexAttrs:
// name:
// flags:
// desiredAccess:
// handle:
// err:
func CreateMutexEx(mutexAttrs *SecurityAttributes, name *uint16, flags uint32, desiredAccess uint32) (handle Handle, err error) {
	r0, _, e1 := syscall.Syscall6(procCreateMutexExW.Addr(), 4, uintptr(unsafe.Pointer(mutexAttrs)), uintptr(unsafe.Pointer(name)), uintptr(flags), uintptr(desiredAccess), 0, 0)
	handle = Handle(r0)
	if handle == 0 || e1 == ERROR_ALREADY_EXISTS {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  创建互斥体(securityAttrs  *安全属性,  初始拥有者  bool,  名称  *uint16)  (句柄  Handle,  错误  error)  {}

// ff:
// mutexAttrs:
// initialOwner:
// name:
// handle:
// err:
func CreateMutex(mutexAttrs *SecurityAttributes, initialOwner bool, name *uint16) (handle Handle, err error) {
	var _p0 uint32
	if initialOwner {
		_p0 = 1
	}
	r0, _, e1 := syscall.Syscall(procCreateMutexW.Addr(), 3, uintptr(unsafe.Pointer(mutexAttrs)), uintptr(_p0), uintptr(unsafe.Pointer(name)))
	handle = Handle(r0)
	if handle == 0 || e1 == ERROR_ALREADY_EXISTS {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  创建命名管道(name  *字符串,  标志  uint32,  管道模式  uint32,  最大实例数  uint32,  输出缓冲区大小  uint32,  输入缓冲区大小  uint32,  默认超时时间  uint32,  安全属性  *安全属性)  (句柄  Handle,  错误  error)  {}

// ff:
// name:
// flags:
// pipeMode:
// maxInstances:
// outSize:
// inSize:
// defaultTimeout:
// sa:
// handle:
// err:
func CreateNamedPipe(name *uint16, flags uint32, pipeMode uint32, maxInstances uint32, outSize uint32, inSize uint32, defaultTimeout uint32, sa *SecurityAttributes) (handle Handle, err error) {
	r0, _, e1 := syscall.Syscall9(procCreateNamedPipeW.Addr(), 8, uintptr(unsafe.Pointer(name)), uintptr(flags), uintptr(pipeMode), uintptr(maxInstances), uintptr(outSize), uintptr(inSize), uintptr(defaultTimeout), uintptr(unsafe.Pointer(sa)), 0)
	handle = Handle(r0)
	if handle == InvalidHandle {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  创建管道(readHandle  *句柄,  writeHandle  *句柄,  安全属性  *安全属性,  管道大小  uint32)  (错误  error)  {}

// ff:
// readhandle:
// writehandle:
// sa:
// size:
// err:
func CreatePipe(readhandle *Handle, writehandle *Handle, sa *SecurityAttributes, size uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procCreatePipe.Addr(), 4, uintptr(unsafe.Pointer(readhandle)), uintptr(unsafe.Pointer(writehandle)), uintptr(unsafe.Pointer(sa)), uintptr(size), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  创建进程(appName  *uint16,  命令行  *uint16,  进程安全属性  *SecurityAttributes,  线程安全属性  *SecurityAttributes,  继承句柄  bool,  创建标志  uint32,  环境变量  *uint16,  当前目录  *uint16,  启动信息  *StartupInfo,  进程信息  *ProcessInformation)  (错误  error)  {}

// ff:
// appName:
// commandLine:
// procSecurity:
// threadSecurity:
// inheritHandles:
// creationFlags:
// env:
// currentDir:
// startupInfo:
// outProcInfo:
// err:
func CreateProcess(appName *uint16, commandLine *uint16, procSecurity *SecurityAttributes, threadSecurity *SecurityAttributes, inheritHandles bool, creationFlags uint32, env *uint16, currentDir *uint16, startupInfo *StartupInfo, outProcInfo *ProcessInformation) (err error) {
	var _p0 uint32
	if inheritHandles {
		_p0 = 1
	}
	r1, _, e1 := syscall.Syscall12(procCreateProcessW.Addr(), 10, uintptr(unsafe.Pointer(appName)), uintptr(unsafe.Pointer(commandLine)), uintptr(unsafe.Pointer(procSecurity)), uintptr(unsafe.Pointer(threadSecurity)), uintptr(_p0), uintptr(creationFlags), uintptr(unsafe.Pointer(env)), uintptr(unsafe.Pointer(currentDir)), uintptr(unsafe.Pointer(startupInfo)), uintptr(unsafe.Pointer(outProcInfo)), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func createPseudoConsole(size uint32, in Handle, out Handle, flags uint32, pconsole *Handle) (hr error) {
	r0, _, _ := syscall.Syscall6(procCreatePseudoConsole.Addr(), 5, uintptr(size), uintptr(in), uintptr(out), uintptr(flags), uintptr(unsafe.Pointer(pconsole)), 0)
	if r0 != 0 {
		hr = syscall.Errno(r0)
	}
	return
}

// 翻译提示:func  创建符号链接(symlinkFileName  *uint16,  targetFileName  *uint16,  flags  uint32)  (错误  error)  {}

// ff:
// symlinkfilename:
// targetfilename:
// flags:
// err:
func CreateSymbolicLink(symlinkfilename *uint16, targetfilename *uint16, flags uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procCreateSymbolicLinkW.Addr(), 3, uintptr(unsafe.Pointer(symlinkfilename)), uintptr(unsafe.Pointer(targetfilename)), uintptr(flags))
	if r1&0xff == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  创建Toolhelp32快照(flags  uint32,  进程Id  uint32)  (句柄  Handle,  错误  error)  {}

// ff:
// flags:
// processId:
// handle:
// err:
func CreateToolhelp32Snapshot(flags uint32, processId uint32) (handle Handle, err error) {
	r0, _, e1 := syscall.Syscall(procCreateToolhelp32Snapshot.Addr(), 2, uintptr(flags), uintptr(processId), 0)
	handle = Handle(r0)
	if handle == InvalidHandle {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  定义Dos设备(flags  uint32,  设备名称  *uint16,  目标路径  *uint16)  (错误  error)  {}

// ff:
// flags:
// deviceName:
// targetPath:
// err:
func DefineDosDevice(flags uint32, deviceName *uint16, targetPath *uint16) (err error) {
	r1, _, e1 := syscall.Syscall(procDefineDosDeviceW.Addr(), 3, uintptr(flags), uintptr(unsafe.Pointer(deviceName)), uintptr(unsafe.Pointer(targetPath)))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  删除文件(path  *字符串)  (错误  error)  {}

// ff:
// path:
// err:
func DeleteFile(path *uint16) (err error) {
	r1, _, e1 := syscall.Syscall(procDeleteFileW.Addr(), 1, uintptr(unsafe.Pointer(path)), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func deleteProcThreadAttributeList(attrlist *ProcThreadAttributeList) {
	syscall.Syscall(procDeleteProcThreadAttributeList.Addr(), 1, uintptr(unsafe.Pointer(attrlist)), 0, 0)
	return
}

// 翻译提示:func  删除卷挂载点(volumeMountPoint  *字符串指针)  (错误  error)  {}

// ff:
// volumeMountPoint:
// err:
func DeleteVolumeMountPoint(volumeMountPoint *uint16) (err error) {
	r1, _, e1 := syscall.Syscall(procDeleteVolumeMountPointW.Addr(), 1, uintptr(unsafe.Pointer(volumeMountPoint)), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  设备控制IO句柄(handle  文件句柄,  控制码  uint32,  输入缓冲区  *byte,  输入缓冲区大小  uint32,  输出缓冲区  *byte,  输出缓冲区大小  uint32,  实际返回字节数  *uint32,  重叠结构体  *重叠结构)  (错误  error)  {}

// ff:
// handle:
// ioControlCode:
// inBuffer:
// inBufferSize:
// outBuffer:
// outBufferSize:
// bytesReturned:
// overlapped:
// err:
func DeviceIoControl(handle Handle, ioControlCode uint32, inBuffer *byte, inBufferSize uint32, outBuffer *byte, outBufferSize uint32, bytesReturned *uint32, overlapped *Overlapped) (err error) {
	r1, _, e1 := syscall.Syscall9(procDeviceIoControl.Addr(), 8, uintptr(handle), uintptr(ioControlCode), uintptr(unsafe.Pointer(inBuffer)), uintptr(inBufferSize), uintptr(unsafe.Pointer(outBuffer)), uintptr(outBufferSize), uintptr(unsafe.Pointer(bytesReturned)), uintptr(unsafe.Pointer(overlapped)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  断开命名管道(pipe  文件句柄)  (错误  error)  {}

// ff:
// pipe:
// err:
func DisconnectNamedPipe(pipe Handle) (err error) {
	r1, _, e1 := syscall.Syscall(procDisconnectNamedPipe.Addr(), 1, uintptr(pipe), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  复制句柄(source进程句柄  Handle,  source句柄  Handle,  target进程句柄  Handle,  目标句柄  *Handle,  请求访问权限  uint32,  是否继承  bool,  选项  uint32)  (错误  error)  {}

// ff:
// hSourceProcessHandle:
// hSourceHandle:
// hTargetProcessHandle:
// lpTargetHandle:
// dwDesiredAccess:
// bInheritHandle:
// dwOptions:
// err:
func DuplicateHandle(hSourceProcessHandle Handle, hSourceHandle Handle, hTargetProcessHandle Handle, lpTargetHandle *Handle, dwDesiredAccess uint32, bInheritHandle bool, dwOptions uint32) (err error) {
	var _p0 uint32
	if bInheritHandle {
		_p0 = 1
	}
	r1, _, e1 := syscall.Syscall9(procDuplicateHandle.Addr(), 7, uintptr(hSourceProcessHandle), uintptr(hSourceHandle), uintptr(hTargetProcessHandle), uintptr(unsafe.Pointer(lpTargetHandle)), uintptr(dwDesiredAccess), uintptr(_p0), uintptr(dwOptions), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  串口控制功能(handle  手柄,  功能码  uint32)  (错误  error)  {}

// ff:
// handle:
// dwFunc:
// err:
func EscapeCommFunction(handle Handle, dwFunc uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procEscapeCommFunction.Addr(), 2, uintptr(handle), uintptr(dwFunc), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  结束进程(exitCode  uint32)  {}

// ff:
// exitcode:
func ExitProcess(exitcode uint32) {
	syscall.Syscall(procExitProcess.Addr(), 1, uintptr(exitcode), 0, 0)
	return
}

// 翻译提示:func  扩展环境字符串(src  *宽字符指针,  dst  *宽字符指针,  size  字节大小)  (长度  uint32,  错误  error)  {}

// ff:
// src:
// dst:
// size:
// n:
// err:
func ExpandEnvironmentStrings(src *uint16, dst *uint16, size uint32) (n uint32, err error) {
	r0, _, e1 := syscall.Syscall(procExpandEnvironmentStringsW.Addr(), 3, uintptr(unsafe.Pointer(src)), uintptr(unsafe.Pointer(dst)), uintptr(size))
	n = uint32(r0)
	if n == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  关闭查找句柄(handle  文件句柄)  (错误  error)  {}

// ff:
// handle:
// err:
func FindClose(handle Handle) (err error) {
	r1, _, e1 := syscall.Syscall(procFindClose.Addr(), 1, uintptr(handle), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  关闭更改通知句柄(handle  文件更改通知句柄)  (错误  error)  {}

// ff:
// handle:
// err:
func FindCloseChangeNotification(handle Handle) (err error) {
	r1, _, e1 := syscall.Syscall(procFindCloseChangeNotification.Addr(), 1, uintptr(handle), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  开始监控更改通知(path  string,  监控子树  bool,  通知过滤器  uint32)  (句柄  Handle,  错误  error)  {}

// ff:
// path:
// watchSubtree:
// notifyFilter:
// handle:
// err:
func FindFirstChangeNotification(path string, watchSubtree bool, notifyFilter uint32) (handle Handle, err error) {
	var _p0 *uint16
	_p0, err = syscall.UTF16PtrFromString(path)
	if err != nil {
		return
	}
	return _FindFirstChangeNotification(_p0, watchSubtree, notifyFilter)
}

func _FindFirstChangeNotification(path *uint16, watchSubtree bool, notifyFilter uint32) (handle Handle, err error) {
	var _p1 uint32
	if watchSubtree {
		_p1 = 1
	}
	r0, _, e1 := syscall.Syscall(procFindFirstChangeNotificationW.Addr(), 3, uintptr(unsafe.Pointer(path)), uintptr(_p1), uintptr(notifyFilter))
	handle = Handle(r0)
	if handle == InvalidHandle {
		err = errnoErr(e1)
	}
	return
}

func findFirstFile1(name *uint16, data *win32finddata1) (handle Handle, err error) {
	r0, _, e1 := syscall.Syscall(procFindFirstFileW.Addr(), 2, uintptr(unsafe.Pointer(name)), uintptr(unsafe.Pointer(data)), 0)
	handle = Handle(r0)
	if handle == InvalidHandle {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  查找第一个卷挂载点(rootPathName  *字符串指针,  volumeMountPoint  *字符串指针,  缓冲区长度  uint32)  (句柄  Handle,  错误  error)  {}

// ff:
// rootPathName:
// volumeMountPoint:
// bufferLength:
// handle:
// err:
func FindFirstVolumeMountPoint(rootPathName *uint16, volumeMountPoint *uint16, bufferLength uint32) (handle Handle, err error) {
	r0, _, e1 := syscall.Syscall(procFindFirstVolumeMountPointW.Addr(), 3, uintptr(unsafe.Pointer(rootPathName)), uintptr(unsafe.Pointer(volumeMountPoint)), uintptr(bufferLength))
	handle = Handle(r0)
	if handle == InvalidHandle {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  第一次找到卷(volumeName  *uint16,  缓冲区长度  uint32)  (句柄  Handle,  错误  error)  {}

// ff:
// volumeName:
// bufferLength:
// handle:
// err:
func FindFirstVolume(volumeName *uint16, bufferLength uint32) (handle Handle, err error) {
	r0, _, e1 := syscall.Syscall(procFindFirstVolumeW.Addr(), 2, uintptr(unsafe.Pointer(volumeName)), uintptr(bufferLength), 0)
	handle = Handle(r0)
	if handle == InvalidHandle {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  查找下一个更改通知(handle  文件句柄)  (错误  error)  {}

// ff:
// handle:
// err:
func FindNextChangeNotification(handle Handle) (err error) {
	r1, _, e1 := syscall.Syscall(procFindNextChangeNotification.Addr(), 1, uintptr(handle), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func findNextFile1(handle Handle, data *win32finddata1) (err error) {
	r1, _, e1 := syscall.Syscall(procFindNextFileW.Addr(), 2, uintptr(handle), uintptr(unsafe.Pointer(data)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  查找下一个卷挂载点(findVolumeMountPoint  手柄,  volumeMountPoint  *uint16,  缓冲区长度  uint32)  (错误  error)  {}

// ff:
// findVolumeMountPoint:
// volumeMountPoint:
// bufferLength:
// err:
func FindNextVolumeMountPoint(findVolumeMountPoint Handle, volumeMountPoint *uint16, bufferLength uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procFindNextVolumeMountPointW.Addr(), 3, uintptr(findVolumeMountPoint), uintptr(unsafe.Pointer(volumeMountPoint)), uintptr(bufferLength))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  查找下一个卷(findVolume  手柄,  卷名称  *uint16,  缓冲区长度  uint32)  (错误  error)  {}

// ff:
// findVolume:
// volumeName:
// bufferLength:
// err:
func FindNextVolume(findVolume Handle, volumeName *uint16, bufferLength uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procFindNextVolumeW.Addr(), 3, uintptr(findVolume), uintptr(unsafe.Pointer(volumeName)), uintptr(bufferLength))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func findResource(module Handle, name uintptr, resType uintptr) (resInfo Handle, err error) {
	r0, _, e1 := syscall.Syscall(procFindResourceW.Addr(), 3, uintptr(module), uintptr(name), uintptr(resType))
	resInfo = Handle(r0)
	if resInfo == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  关闭卷查找(findVolume  卷句柄)  (错误  error)  {}

// ff:
// findVolume:
// err:
func FindVolumeClose(findVolume Handle) (err error) {
	r1, _, e1 := syscall.Syscall(procFindVolumeClose.Addr(), 1, uintptr(findVolume), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  关闭卷挂载点查找器(findVolumeMountPoint  手柄)  (错误  错误)  {}

// ff:
// findVolumeMountPoint:
// err:
func FindVolumeMountPointClose(findVolumeMountPoint Handle) (err error) {
	r1, _, e1 := syscall.Syscall(procFindVolumeMountPointClose.Addr(), 1, uintptr(findVolumeMountPoint), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  刷新文件缓冲(handle  文件句柄)  (错误  error)  {}

// ff:
// handle:
// err:
func FlushFileBuffers(handle Handle) (err error) {
	r1, _, e1 := syscall.Syscall(procFlushFileBuffers.Addr(), 1, uintptr(handle), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  刷新文件视图(addr  uintptr,  长度  uintptr)  (错误  error)  {}

// ff:
// addr:
// length:
// err:
func FlushViewOfFile(addr uintptr, length uintptr) (err error) {
	r1, _, e1 := syscall.Syscall(procFlushViewOfFile.Addr(), 2, uintptr(addr), uintptr(length), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  格式化消息(flags  uint32,  消息源  uintptr,  消息ID  uint32,  语言ID  uint32,  缓冲区  []uint16,  参数  *byte)  (写入字节数  uint32,  错误  error)  {}

// ff:
// flags:
// msgsrc:
// msgid:
// langid:
// buf:
// args:
// n:
// err:
func FormatMessage(flags uint32, msgsrc uintptr, msgid uint32, langid uint32, buf []uint16, args *byte) (n uint32, err error) {
	var _p0 *uint16
	if len(buf) > 0 {
		_p0 = &buf[0]
	}
	r0, _, e1 := syscall.Syscall9(procFormatMessageW.Addr(), 7, uintptr(flags), uintptr(msgsrc), uintptr(msgid), uintptr(langid), uintptr(unsafe.Pointer(_p0)), uintptr(len(buf)), uintptr(unsafe.Pointer(args)), 0, 0)
	n = uint32(r0)
	if n == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  释放环境字符串(envs  *uint16)  (错误  error)  {}

// ff:
// envs:
// err:
func FreeEnvironmentStrings(envs *uint16) (err error) {
	r1, _, e1 := syscall.Syscall(procFreeEnvironmentStringsW.Addr(), 1, uintptr(unsafe.Pointer(envs)), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  释放库(handle  模块句柄)  (错误  error)  {}

// ff:
// handle:
// err:
func FreeLibrary(handle Handle) (err error) {
	r1, _, e1 := syscall.Syscall(procFreeLibrary.Addr(), 1, uintptr(handle), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  生成控制台Ctrl事件(ctrlEventType  控制事件类型,  进程组ID  进程组标识)  (错误  错误)  {}

// ff:
// ctrlEvent:
// processGroupID:
// err:
func GenerateConsoleCtrlEvent(ctrlEvent uint32, processGroupID uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procGenerateConsoleCtrlEvent.Addr(), 2, uintptr(ctrlEvent), uintptr(processGroupID), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  获取ACP()  (字符集  uint32)  {}

// ff:
// acp:
func GetACP() (acp uint32) {
	r0, _, _ := syscall.Syscall(procGetACP.Addr(), 0, 0, 0, 0)
	acp = uint32(r0)
	return
}

// 翻译提示:func  获取活跃处理器计数(处理器组号  uint16)  (活跃处理器数量  uint32)  {}

// ff:
// groupNumber:
// ret:
func GetActiveProcessorCount(groupNumber uint16) (ret uint32) {
	r0, _, _ := syscall.Syscall(procGetActiveProcessorCount.Addr(), 1, uintptr(groupNumber), 0, 0)
	ret = uint32(r0)
	return
}

// 翻译提示:func  获取调制解调器状态(handle  文件句柄,  lpModemStat  调制解调器状态指针)  (错误  error)  {}

// ff:
// handle:
// lpModemStat:
// err:
func GetCommModemStatus(handle Handle, lpModemStat *uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procGetCommModemStatus.Addr(), 2, uintptr(handle), uintptr(unsafe.Pointer(lpModemStat)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  获取通信状态(handle  文件句柄,  lpDCB  通信控制块指针)  (错误  error)  {}

// ff:
// handle:
// lpDCB:
// err:
func GetCommState(handle Handle, lpDCB *DCB) (err error) {
	r1, _, e1 := syscall.Syscall(procGetCommState.Addr(), 2, uintptr(handle), uintptr(unsafe.Pointer(lpDCB)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  获取串行通信超时设置(handle  文件句柄,  timeouts  通信超时结构体指针)  (错误  error)  {}

// ff:
// handle:
// timeouts:
// err:
func GetCommTimeouts(handle Handle, timeouts *CommTimeouts) (err error) {
	r1, _, e1 := syscall.Syscall(procGetCommTimeouts.Addr(), 2, uintptr(handle), uintptr(unsafe.Pointer(timeouts)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  获取命令行()  (命令行  *uint16)  {}

// ff:
// cmd:
func GetCommandLine() (cmd *uint16) {
	r0, _, _ := syscall.Syscall(procGetCommandLineW.Addr(), 0, 0, 0, 0)
	cmd = (*uint16)(unsafe.Pointer(r0))
	return
}

// 翻译提示:func  获取计算机名Ex(nameType  uint32,  buffer  *uint16,  length  *uint32)  (错误  error)  {}

// ff:
// nametype:
// buf:
// n:
// err:
func GetComputerNameEx(nametype uint32, buf *uint16, n *uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procGetComputerNameExW.Addr(), 3, uintptr(nametype), uintptr(unsafe.Pointer(buf)), uintptr(unsafe.Pointer(n)))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  获取计算机名称(buf  *uint16,  n  *uint32)  (错误  error)  {}

// ff:
// buf:
// n:
// err:
func GetComputerName(buf *uint16, n *uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procGetComputerNameW.Addr(), 2, uintptr(unsafe.Pointer(buf)), uintptr(unsafe.Pointer(n)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  获取控制台模式(console  控制台句柄,  模式  *uint32)  (错误  error)  {}

// ff:
// console:
// mode:
// err:
func GetConsoleMode(console Handle, mode *uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procGetConsoleMode.Addr(), 2, uintptr(console), uintptr(unsafe.Pointer(mode)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  获取控制台屏幕缓冲区信息(consoleHandle  手柄,  info  *控制台屏幕缓冲区信息)  (错误  error)  {}

// ff:
// console:
// info:
// err:
func GetConsoleScreenBufferInfo(console Handle, info *ConsoleScreenBufferInfo) (err error) {
	r1, _, e1 := syscall.Syscall(procGetConsoleScreenBufferInfo.Addr(), 2, uintptr(console), uintptr(unsafe.Pointer(info)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  获取当前目录(buflen  uint32,  缓冲区  *uint16)  (目录长度  uint32,  错误  error)  {}

// ff:
// buflen:
// buf:
// n:
// err:
func GetCurrentDirectory(buflen uint32, buf *uint16) (n uint32, err error) {
	r0, _, e1 := syscall.Syscall(procGetCurrentDirectoryW.Addr(), 2, uintptr(buflen), uintptr(unsafe.Pointer(buf)), 0)
	n = uint32(r0)
	if n == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  获取当前进程Id()  (进程标识符  uint32)  {}

// ff:
// pid:
func GetCurrentProcessId() (pid uint32) {
	r0, _, _ := syscall.Syscall(procGetCurrentProcessId.Addr(), 0, 0, 0, 0)
	pid = uint32(r0)
	return
}

// 翻译提示:func  获取当前线程Id()  (线程Id  uint32)  {}

// ff:
// id:
func GetCurrentThreadId() (id uint32) {
	r0, _, _ := syscall.Syscall(procGetCurrentThreadId.Addr(), 0, 0, 0, 0)
	id = uint32(r0)
	return
}

// 翻译提示:func  获取磁盘自由空间Ex(directoryName  string,  freeBytesAvailableToCaller  *uint64,  totalNumberOfBytes  *uint64,  totalNumberOfFreeBytes  *uint64)  (错误  error)  {}

// ff:
// directoryName:
// freeBytesAvailableToCaller:
// totalNumberOfBytes:
// totalNumberOfFreeBytes:
// err:
func GetDiskFreeSpaceEx(directoryName *uint16, freeBytesAvailableToCaller *uint64, totalNumberOfBytes *uint64, totalNumberOfFreeBytes *uint64) (err error) {
	r1, _, e1 := syscall.Syscall6(procGetDiskFreeSpaceExW.Addr(), 4, uintptr(unsafe.Pointer(directoryName)), uintptr(unsafe.Pointer(freeBytesAvailableToCaller)), uintptr(unsafe.Pointer(totalNumberOfBytes)), uintptr(unsafe.Pointer(totalNumberOfFreeBytes)), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  获取驱动器类型(rootPathName  *字节序列)  (驱动器类型  uint32)  {}

// ff:
// rootPathName:
// driveType:
func GetDriveType(rootPathName *uint16) (driveType uint32) {
	r0, _, _ := syscall.Syscall(procGetDriveTypeW.Addr(), 1, uintptr(unsafe.Pointer(rootPathName)), 0, 0)
	driveType = uint32(r0)
	return
}

// 翻译提示:func  获取环境字符串()  (环境变量指针  *uint16,  错误  error)  {}

// ff:
// envs:
// err:
func GetEnvironmentStrings() (envs *uint16, err error) {
	r0, _, e1 := syscall.Syscall(procGetEnvironmentStringsW.Addr(), 0, 0, 0, 0)
	envs = (*uint16)(unsafe.Pointer(r0))
	if envs == nil {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  获取环境变量(name  *字符串指针,  buffer  *字符串指针,  size  字节大小)  (长度  uint32,  错误  error)  {}

// ff:
// name:
// buffer:
// size:
// n:
// err:
func GetEnvironmentVariable(name *uint16, buffer *uint16, size uint32) (n uint32, err error) {
	r0, _, e1 := syscall.Syscall(procGetEnvironmentVariableW.Addr(), 3, uintptr(unsafe.Pointer(name)), uintptr(unsafe.Pointer(buffer)), uintptr(size))
	n = uint32(r0)
	if n == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  获取退出代码进程(handle  进程句柄,  exitcode  *退出代码)  (错误  error)  {}

// ff:
// handle:
// exitcode:
// err:
func GetExitCodeProcess(handle Handle, exitcode *uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procGetExitCodeProcess.Addr(), 2, uintptr(handle), uintptr(unsafe.Pointer(exitcode)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  获取文件属性扩展(name  *字符串字节,  级别  uint32,  信息  *字节)  (错误  error)  {}

// ff:
// name:
// level:
// info:
// err:
func GetFileAttributesEx(name *uint16, level uint32, info *byte) (err error) {
	r1, _, e1 := syscall.Syscall(procGetFileAttributesExW.Addr(), 3, uintptr(unsafe.Pointer(name)), uintptr(level), uintptr(unsafe.Pointer(info)))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  获取文件属性(filename  *uint16)  (属性  uint32,  错误  error)  {}

// ff:
// name:
// attrs:
// err:
func GetFileAttributes(name *uint16) (attrs uint32, err error) {
	r0, _, e1 := syscall.Syscall(procGetFileAttributesW.Addr(), 1, uintptr(unsafe.Pointer(name)), 0, 0)
	attrs = uint32(r0)
	if attrs == INVALID_FILE_ATTRIBUTES {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  获取文件信息通过句柄(handle  文件句柄,  data  *通过句柄的文件信息)  (错误  error)  {}

// ff:
// handle:
// data:
// err:
func GetFileInformationByHandle(handle Handle, data *ByHandleFileInformation) (err error) {
	r1, _, e1 := syscall.Syscall(procGetFileInformationByHandle.Addr(), 2, uintptr(handle), uintptr(unsafe.Pointer(data)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  获取文件信息通过句柄扩展(handle  文件句柄,  class  类型标识,  outBuffer  输出缓冲区指针,  outBufferLen  输出缓冲区长度)  (错误  error)  {}

// ff:
// handle:
// class:
// outBuffer:
// outBufferLen:
// err:
func GetFileInformationByHandleEx(handle Handle, class uint32, outBuffer *byte, outBufferLen uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procGetFileInformationByHandleEx.Addr(), 4, uintptr(handle), uintptr(class), uintptr(unsafe.Pointer(outBuffer)), uintptr(outBufferLen), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  获取文件时间(handle  文件句柄,  创建时间  *时间戳,  最后访问时间  *时间戳,  最后写入时间  *时间戳)  (错误  error)  {}

// ff:
// handle:
// ctime:
// atime:
// wtime:
// err:
func GetFileTime(handle Handle, ctime *Filetime, atime *Filetime, wtime *Filetime) (err error) {
	r1, _, e1 := syscall.Syscall6(procGetFileTime.Addr(), 4, uintptr(handle), uintptr(unsafe.Pointer(ctime)), uintptr(unsafe.Pointer(atime)), uintptr(unsafe.Pointer(wtime)), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  获取文件类型(filehandle  文件句柄)  (fileType  字节计数,  错误  error)  {}

// ff:
// filehandle:
// n:
// err:
func GetFileType(filehandle Handle) (n uint32, err error) {
	r0, _, e1 := syscall.Syscall(procGetFileType.Addr(), 1, uintptr(filehandle), 0, 0)
	n = uint32(r0)
	if n == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  获取最终路径名通过句柄(file  文件句柄,  filePath  路径缓冲区指针,  filePathSize  路径缓冲区大小,  flags  标志)  (实际路径长度  uint32,  错误  error)  {}

// ff:
// file:
// filePath:
// filePathSize:
// flags:
// n:
// err:
func GetFinalPathNameByHandle(file Handle, filePath *uint16, filePathSize uint32, flags uint32) (n uint32, err error) {
	r0, _, e1 := syscall.Syscall6(procGetFinalPathNameByHandleW.Addr(), 4, uintptr(file), uintptr(unsafe.Pointer(filePath)), uintptr(filePathSize), uintptr(flags), 0, 0)
	n = uint32(r0)
	if n == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  获取完整路径名(path  *字符串指针,  缓冲区长度  uint32,  缓冲区  *字符串指针,  文件名  **字符串指针)  (路径长度  uint32,  错误  error)  {}

// ff:
// path:
// buflen:
// buf:
// fname:
// n:
// err:
func GetFullPathName(path *uint16, buflen uint32, buf *uint16, fname **uint16) (n uint32, err error) {
	r0, _, e1 := syscall.Syscall6(procGetFullPathNameW.Addr(), 4, uintptr(unsafe.Pointer(path)), uintptr(buflen), uintptr(unsafe.Pointer(buf)), uintptr(unsafe.Pointer(fname)), 0, 0)
	n = uint32(r0)
	if n == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  获取大页最小尺寸()  (大小  uintptr)  {}

// ff:
// size:
func GetLargePageMinimum() (size uintptr) {
	r0, _, _ := syscall.Syscall(procGetLargePageMinimum.Addr(), 0, 0, 0, 0)
	size = uintptr(r0)
	return
}

// 翻译提示:func  获取最近错误()  (最后错误  error)  {}

// ff:
// lasterr:
func GetLastError() (lasterr error) {
	r0, _, _ := syscall.Syscall(procGetLastError.Addr(), 0, 0, 0, 0)
	if r0 != 0 {
		lasterr = syscall.Errno(r0)
	}
	return
}

// 翻译提示:func  获取逻辑驱动器字符串(bufferLength  uint32,  buffer  *uint16)  (驱动器数  uint32,  错误  error)  {}

// ff:
// bufferLength:
// buffer:
// n:
// err:
func GetLogicalDriveStrings(bufferLength uint32, buffer *uint16) (n uint32, err error) {
	r0, _, e1 := syscall.Syscall(procGetLogicalDriveStringsW.Addr(), 2, uintptr(bufferLength), uintptr(unsafe.Pointer(buffer)), 0)
	n = uint32(r0)
	if n == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  获取逻辑驱动器()  (驱动器位掩码  uint32,  错误  error)  {}

// ff:
// drivesBitMask:
// err:
func GetLogicalDrives() (drivesBitMask uint32, err error) {
	r0, _, e1 := syscall.Syscall(procGetLogicalDrives.Addr(), 0, 0, 0, 0)
	drivesBitMask = uint32(r0)
	if drivesBitMask == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  获取长路径名(path  *字节序列,  buf  *字节序列,  缓冲区长度  uint32)  (长度  uint32,  错误  error)  {}

// ff:
// path:
// buf:
// buflen:
// n:
// err:
func GetLongPathName(path *uint16, buf *uint16, buflen uint32) (n uint32, err error) {
	r0, _, e1 := syscall.Syscall(procGetLongPathNameW.Addr(), 3, uintptr(unsafe.Pointer(path)), uintptr(unsafe.Pointer(buf)), uintptr(buflen))
	n = uint32(r0)
	if n == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  获取最大处理器计数(组号  uint16)  (返回值  uint32)  {}

// ff:
// groupNumber:
// ret:
func GetMaximumProcessorCount(groupNumber uint16) (ret uint32) {
	r0, _, _ := syscall.Syscall(procGetMaximumProcessorCount.Addr(), 1, uintptr(groupNumber), 0, 0)
	ret = uint32(r0)
	return
}

// 翻译提示:func  获取模块文件名(module  模块句柄,  filename  文件名指针,  size  字节数)  (长度  uint32,  错误  error)  {}

// ff:
// module:
// filename:
// size:
// n:
// err:
func GetModuleFileName(module Handle, filename *uint16, size uint32) (n uint32, err error) {
	r0, _, e1 := syscall.Syscall(procGetModuleFileNameW.Addr(), 3, uintptr(module), uintptr(unsafe.Pointer(filename)), uintptr(size))
	n = uint32(r0)
	if n == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  获取模块句柄扩展(flags  uint32,  模块名称  *uint16,  模块指针  *Handle)  (错误  error)  {}

// ff:
// flags:
// moduleName:
// module:
// err:
func GetModuleHandleEx(flags uint32, moduleName *uint16, module *Handle) (err error) {
	r1, _, e1 := syscall.Syscall(procGetModuleHandleExW.Addr(), 3, uintptr(flags), uintptr(unsafe.Pointer(moduleName)), uintptr(unsafe.Pointer(module)))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  获取命名管道状态(pipe  手柄,  状态  *uint32,  当前实例  *uint32,  最大收集计数  *uint32,  数据收集超时  *uint32,  用户名  *uint16,  最大用户名长度  uint32)  (错误  error)  {}

// ff:
// pipe:
// state:
// curInstances:
// maxCollectionCount:
// collectDataTimeout:
// userName:
// maxUserNameSize:
// err:
func GetNamedPipeHandleState(pipe Handle, state *uint32, curInstances *uint32, maxCollectionCount *uint32, collectDataTimeout *uint32, userName *uint16, maxUserNameSize uint32) (err error) {
	r1, _, e1 := syscall.Syscall9(procGetNamedPipeHandleStateW.Addr(), 7, uintptr(pipe), uintptr(unsafe.Pointer(state)), uintptr(unsafe.Pointer(curInstances)), uintptr(unsafe.Pointer(maxCollectionCount)), uintptr(unsafe.Pointer(collectDataTimeout)), uintptr(unsafe.Pointer(userName)), uintptr(maxUserNameSize), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  获取命名管道信息(pipe  接口句柄,  标志  *uint32,  输出缓冲区大小  *uint32,  输入缓冲区大小  *uint32,  最大实例数  *uint32)  (错误  error)  {}

// ff:
// pipe:
// flags:
// outSize:
// inSize:
// maxInstances:
// err:
func GetNamedPipeInfo(pipe Handle, flags *uint32, outSize *uint32, inSize *uint32, maxInstances *uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procGetNamedPipeInfo.Addr(), 5, uintptr(pipe), uintptr(unsafe.Pointer(flags)), uintptr(unsafe.Pointer(outSize)), uintptr(unsafe.Pointer(inSize)), uintptr(unsafe.Pointer(maxInstances)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  获取重叠结果(handle  操作句柄,  overlapped  重叠结构体指针,  done  完成计数指针,  wait  是否等待  bool)  (错误  error)  {}

// ff:
// handle:
// overlapped:
// done:
// wait:
// err:
func GetOverlappedResult(handle Handle, overlapped *Overlapped, done *uint32, wait bool) (err error) {
	var _p0 uint32
	if wait {
		_p0 = 1
	}
	r1, _, e1 := syscall.Syscall6(procGetOverlappedResult.Addr(), 4, uintptr(handle), uintptr(unsafe.Pointer(overlapped)), uintptr(unsafe.Pointer(done)), uintptr(_p0), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  获取进程优先级类(process  进程句柄)  (优先级  uint32,  错误  error)  {}

// ff:
// process:
// ret:
// err:
func GetPriorityClass(process Handle) (ret uint32, err error) {
	r0, _, e1 := syscall.Syscall(procGetPriorityClass.Addr(), 1, uintptr(process), 0, 0)
	ret = uint32(r0)
	if ret == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  获取ProcAddress(模块句柄  Handle,  函数名  string)  (函数指针  uintptr,  错误  error)  {}

// ff:
// module:
// procname:
// proc:
// err:
func GetProcAddress(module Handle, procname string) (proc uintptr, err error) {
	var _p0 *byte
	_p0, err = syscall.BytePtrFromString(procname)
	if err != nil {
		return
	}
	return _GetProcAddress(module, _p0)
}

func _GetProcAddress(module Handle, procname *byte) (proc uintptr, err error) {
	r0, _, e1 := syscall.Syscall(procGetProcAddress.Addr(), 2, uintptr(module), uintptr(unsafe.Pointer(procname)), 0)
	proc = uintptr(r0)
	if proc == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  获取进程ID(process  进程句柄)  (id  进程ID,  err  错误)  {}

// ff:
// process:
// id:
// err:
func GetProcessId(process Handle) (id uint32, err error) {
	r0, _, e1 := syscall.Syscall(procGetProcessId.Addr(), 1, uintptr(process), 0, 0)
	id = uint32(r0)
	if id == 0 {
		err = errnoErr(e1)
	}
	return
}

func getProcessPreferredUILanguages(flags uint32, numLanguages *uint32, buf *uint16, bufSize *uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procGetProcessPreferredUILanguages.Addr(), 4, uintptr(flags), uintptr(unsafe.Pointer(numLanguages)), uintptr(unsafe.Pointer(buf)), uintptr(unsafe.Pointer(bufSize)), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  获取进程关闭参数(level  *uint32,  flags  *uint32)  (错误  error)  {}

// ff:
// level:
// flags:
// err:
func GetProcessShutdownParameters(level *uint32, flags *uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procGetProcessShutdownParameters.Addr(), 2, uintptr(unsafe.Pointer(level)), uintptr(unsafe.Pointer(flags)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  获取进程时间(handle  进程句柄,  creationTime  创建时间  *文件时间,  exitTime  退出时间  *文件时间,  kernelTime  内核时间  *文件时间,  userTime  用户时间  *文件时间)  (err  错误)  {}

// ff:
// handle:
// creationTime:
// exitTime:
// kernelTime:
// userTime:
// err:
func GetProcessTimes(handle Handle, creationTime *Filetime, exitTime *Filetime, kernelTime *Filetime, userTime *Filetime) (err error) {
	r1, _, e1 := syscall.Syscall6(procGetProcessTimes.Addr(), 5, uintptr(handle), uintptr(unsafe.Pointer(creationTime)), uintptr(unsafe.Pointer(exitTime)), uintptr(unsafe.Pointer(kernelTime)), uintptr(unsafe.Pointer(userTime)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  获取进程工作集大小扩展(hProcess  进程句柄,  lpMinimumWorkingSetSize  最小工作集大小指针,  lpMaximumWorkingSetSize  最大工作集大小指针,  flags  参数标志指针)  {}

// ff:
// hProcess:
// lpMinimumWorkingSetSize:
// lpMaximumWorkingSetSize:
// flags:
func GetProcessWorkingSetSizeEx(hProcess Handle, lpMinimumWorkingSetSize *uintptr, lpMaximumWorkingSetSize *uintptr, flags *uint32) {
	syscall.Syscall6(procGetProcessWorkingSetSizeEx.Addr(), 4, uintptr(hProcess), uintptr(unsafe.Pointer(lpMinimumWorkingSetSize)), uintptr(unsafe.Pointer(lpMaximumWorkingSetSize)), uintptr(unsafe.Pointer(flags)), 0, 0)
	return
}

// 翻译提示:func  获取已完成状态(cphandle  文件句柄,  数量  *uint32,  关键字  *uintptr,  重叠  *Overlapped,  超时时间  uint32)  (错误  error)  {}

// ff:
// cphandle:
// qty:
// key:
// overlapped:
// timeout:
// err:
func GetQueuedCompletionStatus(cphandle Handle, qty *uint32, key *uintptr, overlapped **Overlapped, timeout uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procGetQueuedCompletionStatus.Addr(), 5, uintptr(cphandle), uintptr(unsafe.Pointer(qty)), uintptr(unsafe.Pointer(key)), uintptr(unsafe.Pointer(overlapped)), uintptr(timeout), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  获取短路径名(longPath  *uint16,  shortPath  *uint16,  bufferLength  uint32)  (转换长度  uint32,  错误  error)  {}

// ff:
// longpath:
// shortpath:
// buflen:
// n:
// err:
func GetShortPathName(longpath *uint16, shortpath *uint16, buflen uint32) (n uint32, err error) {
	r0, _, e1 := syscall.Syscall(procGetShortPathNameW.Addr(), 3, uintptr(unsafe.Pointer(longpath)), uintptr(unsafe.Pointer(shortpath)), uintptr(buflen))
	n = uint32(r0)
	if n == 0 {
		err = errnoErr(e1)
	}
	return
}

func getStartupInfo(startupInfo *StartupInfo) {
	syscall.Syscall(procGetStartupInfoW.Addr(), 1, uintptr(unsafe.Pointer(startupInfo)), 0, 0)
	return
}

// 翻译提示:func  获取标准句柄(stdhandle  uint32)  (句柄  Handle,  错误  error)  {}

// ff:
// stdhandle:
// handle:
// err:
func GetStdHandle(stdhandle uint32) (handle Handle, err error) {
	r0, _, e1 := syscall.Syscall(procGetStdHandle.Addr(), 1, uintptr(stdhandle), 0, 0)
	handle = Handle(r0)
	if handle == InvalidHandle {
		err = errnoErr(e1)
	}
	return
}

func getSystemDirectory(dir *uint16, dirLen uint32) (len uint32, err error) {
	r0, _, e1 := syscall.Syscall(procGetSystemDirectoryW.Addr(), 2, uintptr(unsafe.Pointer(dir)), uintptr(dirLen), 0)
	len = uint32(r0)
	if len == 0 {
		err = errnoErr(e1)
	}
	return
}

func getSystemPreferredUILanguages(flags uint32, numLanguages *uint32, buf *uint16, bufSize *uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procGetSystemPreferredUILanguages.Addr(), 4, uintptr(flags), uintptr(unsafe.Pointer(numLanguages)), uintptr(unsafe.Pointer(buf)), uintptr(unsafe.Pointer(bufSize)), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  获取系统时间作为文件时间(time  *文件时间)  {}

// ff:
// time:
func GetSystemTimeAsFileTime(time *Filetime) {
	syscall.Syscall(procGetSystemTimeAsFileTime.Addr(), 1, uintptr(unsafe.Pointer(time)), 0, 0)
	return
}

// 翻译提示:func  获取系统时间精确到文件时间(time  *文件时间)  {}

// ff:
// time:
func GetSystemTimePreciseAsFileTime(time *Filetime) {
	syscall.Syscall(procGetSystemTimePreciseAsFileTime.Addr(), 1, uintptr(unsafe.Pointer(time)), 0, 0)
	return
}

func getSystemWindowsDirectory(dir *uint16, dirLen uint32) (len uint32, err error) {
	r0, _, e1 := syscall.Syscall(procGetSystemWindowsDirectoryW.Addr(), 2, uintptr(unsafe.Pointer(dir)), uintptr(dirLen), 0)
	len = uint32(r0)
	if len == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  获取临时路径(buflen  字节长度,  buf  字符串指针)  (长度  uint32,  错误  error)  {}

// ff:
// buflen:
// buf:
// n:
// err:
func GetTempPath(buflen uint32, buf *uint16) (n uint32, err error) {
	r0, _, e1 := syscall.Syscall(procGetTempPathW.Addr(), 2, uintptr(buflen), uintptr(unsafe.Pointer(buf)), 0)
	n = uint32(r0)
	if n == 0 {
		err = errnoErr(e1)
	}
	return
}

func getThreadPreferredUILanguages(flags uint32, numLanguages *uint32, buf *uint16, bufSize *uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procGetThreadPreferredUILanguages.Addr(), 4, uintptr(flags), uintptr(unsafe.Pointer(numLanguages)), uintptr(unsafe.Pointer(buf)), uintptr(unsafe.Pointer(bufSize)), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func getTickCount64() (ms uint64) {
	r0, _, _ := syscall.Syscall(procGetTickCount64.Addr(), 0, 0, 0, 0)
	ms = uint64(r0)
	return
}

// 翻译提示:func  获取时区信息(tzi  *时区信息)  (返回码  uint32,  错误  error)  {}

// ff:
// tzi:
// rc:
// err:
func GetTimeZoneInformation(tzi *Timezoneinformation) (rc uint32, err error) {
	r0, _, e1 := syscall.Syscall(procGetTimeZoneInformation.Addr(), 1, uintptr(unsafe.Pointer(tzi)), 0, 0)
	rc = uint32(r0)
	if rc == 0xffffffff {
		err = errnoErr(e1)
	}
	return
}

func getUserPreferredUILanguages(flags uint32, numLanguages *uint32, buf *uint16, bufSize *uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procGetUserPreferredUILanguages.Addr(), 4, uintptr(flags), uintptr(unsafe.Pointer(numLanguages)), uintptr(unsafe.Pointer(buf)), uintptr(unsafe.Pointer(bufSize)), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  获取版本()  (版本号  uint32,  错误  error)  {}

// ff:
// ver:
// err:
func GetVersion() (ver uint32, err error) {
	r0, _, e1 := syscall.Syscall(procGetVersion.Addr(), 0, 0, 0, 0)
	ver = uint32(r0)
	if ver == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  获取卷信息ByHandle(file  文件句柄,  volumeNameBuffer  *uint16,  volumeNameSize  uint32,  volumeNameSerialNumber  *uint32,  maximumComponentLength  *uint32,  fileSystemFlags  *uint32,  fileSystemNameBuffer  *uint16,  fileSystemNameSize  uint32)  (err  错误)  {}

// ff:
// file:
// volumeNameBuffer:
// volumeNameSize:
// volumeNameSerialNumber:
// maximumComponentLength:
// fileSystemFlags:
// fileSystemNameBuffer:
// fileSystemNameSize:
// err:
func GetVolumeInformationByHandle(file Handle, volumeNameBuffer *uint16, volumeNameSize uint32, volumeNameSerialNumber *uint32, maximumComponentLength *uint32, fileSystemFlags *uint32, fileSystemNameBuffer *uint16, fileSystemNameSize uint32) (err error) {
	r1, _, e1 := syscall.Syscall9(procGetVolumeInformationByHandleW.Addr(), 8, uintptr(file), uintptr(unsafe.Pointer(volumeNameBuffer)), uintptr(volumeNameSize), uintptr(unsafe.Pointer(volumeNameSerialNumber)), uintptr(unsafe.Pointer(maximumComponentLength)), uintptr(unsafe.Pointer(fileSystemFlags)), uintptr(unsafe.Pointer(fileSystemNameBuffer)), uintptr(fileSystemNameSize), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  获取卷信息(rootPathName  *uint16,  volumeNameBuffer  *uint16,  volumeNameSize  uint32,  volumeSerialNumber  *uint32,  最大组件长度  *uint32,  文件系统标志  *uint32,  文件系统名称Buffer  *uint16,  文件系统名称Size  uint32)  (错误  error)  {}

// ff:
// rootPathName:
// volumeNameBuffer:
// volumeNameSize:
// volumeNameSerialNumber:
// maximumComponentLength:
// fileSystemFlags:
// fileSystemNameBuffer:
// fileSystemNameSize:
// err:
func GetVolumeInformation(rootPathName *uint16, volumeNameBuffer *uint16, volumeNameSize uint32, volumeNameSerialNumber *uint32, maximumComponentLength *uint32, fileSystemFlags *uint32, fileSystemNameBuffer *uint16, fileSystemNameSize uint32) (err error) {
	r1, _, e1 := syscall.Syscall9(procGetVolumeInformationW.Addr(), 8, uintptr(unsafe.Pointer(rootPathName)), uintptr(unsafe.Pointer(volumeNameBuffer)), uintptr(volumeNameSize), uintptr(unsafe.Pointer(volumeNameSerialNumber)), uintptr(unsafe.Pointer(maximumComponentLength)), uintptr(unsafe.Pointer(fileSystemFlags)), uintptr(unsafe.Pointer(fileSystemNameBuffer)), uintptr(fileSystemNameSize), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  获取卷名For卷安装点(volume安装点  *uint16,  卷名  *uint16,  缓冲区长度  uint32)  (错误  error)  {}

// ff:
// volumeMountPoint:
// volumeName:
// bufferlength:
// err:
func GetVolumeNameForVolumeMountPoint(volumeMountPoint *uint16, volumeName *uint16, bufferlength uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procGetVolumeNameForVolumeMountPointW.Addr(), 3, uintptr(unsafe.Pointer(volumeMountPoint)), uintptr(unsafe.Pointer(volumeName)), uintptr(bufferlength))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  获取卷路径名(fileName  *uint16,  volumePathName  *uint16,  缓冲区长度  uint32)  (错误  error)  {}

// ff:
// fileName:
// volumePathName:
// bufferLength:
// err:
func GetVolumePathName(fileName *uint16, volumePathName *uint16, bufferLength uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procGetVolumePathNameW.Addr(), 3, uintptr(unsafe.Pointer(fileName)), uintptr(unsafe.Pointer(volumePathName)), uintptr(bufferLength))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  获取卷路径名称For卷名称(volumeName  *uint16,  volumePathNames  *uint16,  缓冲区长度  uint32,  返回长度  *uint32)  (错误  error)  {}

// ff:
// volumeName:
// volumePathNames:
// bufferLength:
// returnLength:
// err:
func GetVolumePathNamesForVolumeName(volumeName *uint16, volumePathNames *uint16, bufferLength uint32, returnLength *uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procGetVolumePathNamesForVolumeNameW.Addr(), 4, uintptr(unsafe.Pointer(volumeName)), uintptr(unsafe.Pointer(volumePathNames)), uintptr(bufferLength), uintptr(unsafe.Pointer(returnLength)), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func getWindowsDirectory(dir *uint16, dirLen uint32) (len uint32, err error) {
	r0, _, e1 := syscall.Syscall(procGetWindowsDirectoryW.Addr(), 2, uintptr(unsafe.Pointer(dir)), uintptr(dirLen), 0)
	len = uint32(r0)
	if len == 0 {
		err = errnoErr(e1)
	}
	return
}

func initializeProcThreadAttributeList(attrlist *ProcThreadAttributeList, attrcount uint32, flags uint32, size *uintptr) (err error) {
	r1, _, e1 := syscall.Syscall6(procInitializeProcThreadAttributeList.Addr(), 4, uintptr(unsafe.Pointer(attrlist)), uintptr(attrcount), uintptr(flags), uintptr(unsafe.Pointer(size)), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  是否为Wow64进程(handle  进程句柄,  isWow64  *bool)  (错误  error)  {}

// ff:
// handle:
// isWow64:
// err:
func IsWow64Process(handle Handle, isWow64 *bool) (err error) {
	var _p0 uint32
	if *isWow64 {
		_p0 = 1
	}
	r1, _, e1 := syscall.Syscall(procIsWow64Process.Addr(), 2, uintptr(handle), uintptr(unsafe.Pointer(&_p0)), 0)
	*isWow64 = _p0 != 0
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  是否为Wow64进程2(handle  进程句柄,  processMachine  *uint16,  nativeMachine  *uint16)  (错误  error)  {}

// ff:
// handle:
// processMachine:
// nativeMachine:
// err:
func IsWow64Process2(handle Handle, processMachine *uint16, nativeMachine *uint16) (err error) {
	err = procIsWow64Process2.Find()
	if err != nil {
		return
	}
	r1, _, e1 := syscall.Syscall(procIsWow64Process2.Addr(), 3, uintptr(handle), uintptr(unsafe.Pointer(processMachine)), uintptr(unsafe.Pointer(nativeMachine)))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  加载库Ex(库名称  string,  零Handle  Handle,  标志  uintptr)  (库句柄  Handle,  错误  error)  {}

// ff:
// libname:
// zero:
// flags:
// handle:
// err:
func LoadLibraryEx(libname string, zero Handle, flags uintptr) (handle Handle, err error) {
	var _p0 *uint16
	_p0, err = syscall.UTF16PtrFromString(libname)
	if err != nil {
		return
	}
	return _LoadLibraryEx(_p0, zero, flags)
}

func _LoadLibraryEx(libname *uint16, zero Handle, flags uintptr) (handle Handle, err error) {
	r0, _, e1 := syscall.Syscall(procLoadLibraryExW.Addr(), 3, uintptr(unsafe.Pointer(libname)), uintptr(zero), uintptr(flags))
	handle = Handle(r0)
	if handle == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  加载库(libname  字符串)  (句柄  Handle,  错误  error)  {}

// ff:
// libname:
// handle:
// err:
func LoadLibrary(libname string) (handle Handle, err error) {
	var _p0 *uint16
	_p0, err = syscall.UTF16PtrFromString(libname)
	if err != nil {
		return
	}
	return _LoadLibrary(_p0)
}

func _LoadLibrary(libname *uint16) (handle Handle, err error) {
	r0, _, e1 := syscall.Syscall(procLoadLibraryW.Addr(), 1, uintptr(unsafe.Pointer(libname)), 0, 0)
	handle = Handle(r0)
	if handle == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  加载资源模块(hModule  模块句柄,  hResInfo  资源信息句柄)  (hResData  资源数据句柄,  err  错误)  {}

// ff:
// module:
// resInfo:
// resData:
// err:
func LoadResource(module Handle, resInfo Handle) (resData Handle, err error) {
	r0, _, e1 := syscall.Syscall(procLoadResource.Addr(), 2, uintptr(module), uintptr(resInfo), 0)
	resData = Handle(r0)
	if resData == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  本地分配内存(标志  uint32,  长度  uint32)  (指针  uintptr,  错误  error)  {}

// ff:
// flags:
// length:
// ptr:
// err:
func LocalAlloc(flags uint32, length uint32) (ptr uintptr, err error) {
	r0, _, e1 := syscall.Syscall(procLocalAlloc.Addr(), 2, uintptr(flags), uintptr(length), 0)
	ptr = uintptr(r0)
	if ptr == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  本地释放内存(hmem  手柄)  (新手柄  Handle,  错误  error)  {}

// ff:
// hmem:
// handle:
// err:
func LocalFree(hmem Handle) (handle Handle, err error) {
	r0, _, e1 := syscall.Syscall(procLocalFree.Addr(), 1, uintptr(hmem), 0, 0)
	handle = Handle(r0)
	if handle != 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  锁定文件扩展(file  文件句柄,  标志  uint32,  保留  uint32,  低字节长度  uint32,  高字节长度  uint32,  重叠结构  *重叠结构)  (错误  error)  {}

// ff:
// file:
// flags:
// reserved:
// bytesLow:
// bytesHigh:
// overlapped:
// err:
func LockFileEx(file Handle, flags uint32, reserved uint32, bytesLow uint32, bytesHigh uint32, overlapped *Overlapped) (err error) {
	r1, _, e1 := syscall.Syscall6(procLockFileEx.Addr(), 6, uintptr(file), uintptr(flags), uintptr(reserved), uintptr(bytesLow), uintptr(bytesHigh), uintptr(unsafe.Pointer(overlapped)))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  加锁资源(resData  手柄)  (地址  uintptr,  错误  error)  {}

// ff:
// resData:
// addr:
// err:
func LockResource(resData Handle) (addr uintptr, err error) {
	r0, _, e1 := syscall.Syscall(procLockResource.Addr(), 1, uintptr(resData), 0, 0)
	addr = uintptr(r0)
	if addr == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  映射视图到文件句柄(handle  文件句柄,  访问权限  uint32,  高偏移量  uint32,  低偏移量  uint32,  长度  uintptr)  (内存地址  uintptr,  错误  error)  {}

// ff:
// handle:
// access:
// offsetHigh:
// offsetLow:
// length:
// addr:
// err:
func MapViewOfFile(handle Handle, access uint32, offsetHigh uint32, offsetLow uint32, length uintptr) (addr uintptr, err error) {
	r0, _, e1 := syscall.Syscall6(procMapViewOfFile.Addr(), 5, uintptr(handle), uintptr(access), uintptr(offsetHigh), uintptr(offsetLow), uintptr(length), 0)
	addr = uintptr(r0)
	if addr == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  模块32首次获取(snapshot  进程快照句柄,  moduleEntry  模块信息32结构体指针)  (错误  error)  {}

// ff:
// snapshot:
// moduleEntry:
// err:
func Module32First(snapshot Handle, moduleEntry *ModuleEntry32) (err error) {
	r1, _, e1 := syscall.Syscall(procModule32FirstW.Addr(), 2, uintptr(snapshot), uintptr(unsafe.Pointer(moduleEntry)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  下一个模块32(snapshot  处理器句柄,  模块条目  *模块条目32)  (错误  error)  {}

// ff:
// snapshot:
// moduleEntry:
// err:
func Module32Next(snapshot Handle, moduleEntry *ModuleEntry32) (err error) {
	r1, _, e1 := syscall.Syscall(procModule32NextW.Addr(), 2, uintptr(snapshot), uintptr(unsafe.Pointer(moduleEntry)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  移动文件(from  源路径  *uint16,  to  目标路径  *uint16,  标志  uint32)  (错误  error)  {}

// ff:
// from:
// to:
// flags:
// err:
func MoveFileEx(from *uint16, to *uint16, flags uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procMoveFileExW.Addr(), 3, uintptr(unsafe.Pointer(from)), uintptr(unsafe.Pointer(to)), uintptr(flags))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  移动文件(源  *uint16,  目标  *uint16)  (错误  error)  {}

// ff:
// from:
// to:
// err:
func MoveFile(from *uint16, to *uint16) (err error) {
	r1, _, e1 := syscall.Syscall(procMoveFileW.Addr(), 2, uintptr(unsafe.Pointer(from)), uintptr(unsafe.Pointer(to)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  多字节转宽字符(codePage  uint32,  标志  uint32,  字符串  *byte,  字符串长度  int32,  宽字符  *uint16,  宽字符长度  int32)  (实际转换数  int32,  错误  error)  {}

// ff:
// codePage:
// dwFlags:
// str:
// nstr:
// wchar:
// nwchar:
// nwrite:
// err:
func MultiByteToWideChar(codePage uint32, dwFlags uint32, str *byte, nstr int32, wchar *uint16, nwchar int32) (nwrite int32, err error) {
	r0, _, e1 := syscall.Syscall6(procMultiByteToWideChar.Addr(), 6, uintptr(codePage), uintptr(dwFlags), uintptr(unsafe.Pointer(str)), uintptr(nstr), uintptr(unsafe.Pointer(wchar)), uintptr(nwchar))
	nwrite = int32(r0)
	if nwrite == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  打开事件(所需访问权限  uint32,  继承句柄  bool,  事件名称  *uint16)  (事件句柄  Handle,  错误  error)  {}

// ff:
// desiredAccess:
// inheritHandle:
// name:
// handle:
// err:
func OpenEvent(desiredAccess uint32, inheritHandle bool, name *uint16) (handle Handle, err error) {
	var _p0 uint32
	if inheritHandle {
		_p0 = 1
	}
	r0, _, e1 := syscall.Syscall(procOpenEventW.Addr(), 3, uintptr(desiredAccess), uintptr(_p0), uintptr(unsafe.Pointer(name)))
	handle = Handle(r0)
	if handle == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  创建互斥量(desiredAccess  访问权限:  uint32,  inheritHandle  是否继承:  bool,  name  名称:  *uint16)  (句柄  Handle,  错误  error)  {}

// ff:
// desiredAccess:
// inheritHandle:
// name:
// handle:
// err:
func OpenMutex(desiredAccess uint32, inheritHandle bool, name *uint16) (handle Handle, err error) {
	var _p0 uint32
	if inheritHandle {
		_p0 = 1
	}
	r0, _, e1 := syscall.Syscall(procOpenMutexW.Addr(), 3, uintptr(desiredAccess), uintptr(_p0), uintptr(unsafe.Pointer(name)))
	handle = Handle(r0)
	if handle == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  打开进程(desiredAccess  权限访问,  inheritHandle  是否继承句柄,  processId  进程ID)  (handle  进程句柄,  err  错误)  {}

// ff:
// desiredAccess:
// inheritHandle:
// processId:
// handle:
// err:
func OpenProcess(desiredAccess uint32, inheritHandle bool, processId uint32) (handle Handle, err error) {
	var _p0 uint32
	if inheritHandle {
		_p0 = 1
	}
	r0, _, e1 := syscall.Syscall(procOpenProcess.Addr(), 3, uintptr(desiredAccess), uintptr(_p0), uintptr(processId))
	handle = Handle(r0)
	if handle == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  创建线程(所需访问权限  uint32,  手柄继承  bool,  线程ID  uint32)  (线程句柄  Handle,  错误  error)  {}

// ff:
// desiredAccess:
// inheritHandle:
// threadId:
// handle:
// err:
func OpenThread(desiredAccess uint32, inheritHandle bool, threadId uint32) (handle Handle, err error) {
	var _p0 uint32
	if inheritHandle {
		_p0 = 1
	}
	r0, _, e1 := syscall.Syscall(procOpenThread.Addr(), 3, uintptr(desiredAccess), uintptr(_p0), uintptr(threadId))
	handle = Handle(r0)
	if handle == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  发送完成队列状态(cphandle  文件句柄,  完成数量  uint32,  关键字  uintptr,  重叠结构  *重叠结构)  (错误  error)  {}

// ff:
// cphandle:
// qty:
// key:
// overlapped:
// err:
func PostQueuedCompletionStatus(cphandle Handle, qty uint32, key uintptr, overlapped *Overlapped) (err error) {
	r1, _, e1 := syscall.Syscall6(procPostQueuedCompletionStatus.Addr(), 4, uintptr(cphandle), uintptr(qty), uintptr(key), uintptr(unsafe.Pointer(overlapped)), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  进程32首个(snapshot  进程快照,  进程条目  *进程条目32)  (错误  error)  {}

// ff:
// snapshot:
// procEntry:
// err:
func Process32First(snapshot Handle, procEntry *ProcessEntry32) (err error) {
	r1, _, e1 := syscall.Syscall(procProcess32FirstW.Addr(), 2, uintptr(snapshot), uintptr(unsafe.Pointer(procEntry)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  进程32位下一个(snapshot  进程快照,  procEntry  进程条目32)  (错误  error)  {}

// ff:
// snapshot:
// procEntry:
// err:
func Process32Next(snapshot Handle, procEntry *ProcessEntry32) (err error) {
	r1, _, e1 := syscall.Syscall(procProcess32NextW.Addr(), 2, uintptr(snapshot), uintptr(unsafe.Pointer(procEntry)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  进程Id转SessionId(pid  uint32,  sessionid  *uint32)  (错误  error)  {}

// ff:
// pid:
// sessionid:
// err:
func ProcessIdToSessionId(pid uint32, sessionid *uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procProcessIdToSessionId.Addr(), 2, uintptr(pid), uintptr(unsafe.Pointer(sessionid)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  脉冲事件(event  事件句柄)  (错误  error)  {}

// ff:
// event:
// err:
func PulseEvent(event Handle) (err error) {
	r1, _, e1 := syscall.Syscall(procPulseEvent.Addr(), 1, uintptr(event), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  清除通讯(handle  设备句柄,  标志  uint32)  (错误  error)  {}

// ff:
// handle:
// dwFlags:
// err:
func PurgeComm(handle Handle, dwFlags uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procPurgeComm.Addr(), 2, uintptr(handle), uintptr(dwFlags), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  查询Dos设备(deviceName  *uint16,  targetPath  *uint16,  max  uint32)  (长度  uint32,  错误  error)  {}

// ff:
// deviceName:
// targetPath:
// max:
// n:
// err:
func QueryDosDevice(deviceName *uint16, targetPath *uint16, max uint32) (n uint32, err error) {
	r0, _, e1 := syscall.Syscall(procQueryDosDeviceW.Addr(), 3, uintptr(unsafe.Pointer(deviceName)), uintptr(unsafe.Pointer(targetPath)), uintptr(max))
	n = uint32(r0)
	if n == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  查询完整进程映像名称(proc  进程句柄,  标志  uint32,  exeName  模块名称指针,  size  名称缓冲区大小指针)  (错误  error)  {}

// ff:
// proc:
// flags:
// exeName:
// size:
// err:
func QueryFullProcessImageName(proc Handle, flags uint32, exeName *uint16, size *uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procQueryFullProcessImageNameW.Addr(), 4, uintptr(proc), uintptr(flags), uintptr(unsafe.Pointer(exeName)), uintptr(unsafe.Pointer(size)), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  查询作业信息(job  手柄,  作业对象信息类别  int32,  作业对象信息  uintptr,  作业对象信息长度  uint32,  返回长度  *uint32)  (错误  error)  {}

// ff:
// job:
// JobObjectInformationClass:
// JobObjectInformation:
// JobObjectInformationLength:
// retlen:
// err:
func QueryInformationJobObject(job Handle, JobObjectInformationClass int32, JobObjectInformation uintptr, JobObjectInformationLength uint32, retlen *uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procQueryInformationJobObject.Addr(), 5, uintptr(job), uintptr(JobObjectInformationClass), uintptr(JobObjectInformation), uintptr(JobObjectInformationLength), uintptr(unsafe.Pointer(retlen)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  读取控制台(console  手柄,  缓冲区  *uint16,  要读取  uint32,  已读取  *uint32,  输入控制  *byte)  (错误  error)  {}

// ff:
// console:
// buf:
// toread:
// read:
// inputControl:
// err:
func ReadConsole(console Handle, buf *uint16, toread uint32, read *uint32, inputControl *byte) (err error) {
	r1, _, e1 := syscall.Syscall6(procReadConsoleW.Addr(), 5, uintptr(console), uintptr(unsafe.Pointer(buf)), uintptr(toread), uintptr(unsafe.Pointer(read)), uintptr(unsafe.Pointer(inputControl)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  监控目录变化(handle  文件句柄,  buf  数据缓冲区,  缓冲区长度  uint32,  是否监视子树  bool,  事件掩码  uint32,  返回长度  *uint32,  重叠结构  *重叠结构体,  完成例程  uintptr)  (错误  error)  {}

// ff:
// handle:
// buf:
// buflen:
// watchSubTree:
// mask:
// retlen:
// overlapped:
// completionRoutine:
// err:
func ReadDirectoryChanges(handle Handle, buf *byte, buflen uint32, watchSubTree bool, mask uint32, retlen *uint32, overlapped *Overlapped, completionRoutine uintptr) (err error) {
	var _p0 uint32
	if watchSubTree {
		_p0 = 1
	}
	r1, _, e1 := syscall.Syscall9(procReadDirectoryChangesW.Addr(), 8, uintptr(handle), uintptr(unsafe.Pointer(buf)), uintptr(buflen), uintptr(_p0), uintptr(mask), uintptr(unsafe.Pointer(retlen)), uintptr(unsafe.Pointer(overlapped)), uintptr(completionRoutine), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func readFile(handle Handle, buf []byte, done *uint32, overlapped *Overlapped) (err error) {
	var _p0 *byte
	if len(buf) > 0 {
		_p0 = &buf[0]
	}
	r1, _, e1 := syscall.Syscall6(procReadFile.Addr(), 5, uintptr(handle), uintptr(unsafe.Pointer(_p0)), uintptr(len(buf)), uintptr(unsafe.Pointer(done)), uintptr(unsafe.Pointer(overlapped)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  读取进程内存(process  进程句柄,  baseAddress  基地址,  buffer  内存缓冲区指针,  size  内存大小,  numberOfBytesRead  已读取字节数指针)  (错误  error)  {}

// ff:
// process:
// baseAddress:
// buffer:
// size:
// numberOfBytesRead:
// err:
func ReadProcessMemory(process Handle, baseAddress uintptr, buffer *byte, size uintptr, numberOfBytesRead *uintptr) (err error) {
	r1, _, e1 := syscall.Syscall6(procReadProcessMemory.Addr(), 5, uintptr(process), uintptr(baseAddress), uintptr(unsafe.Pointer(buffer)), uintptr(size), uintptr(unsafe.Pointer(numberOfBytesRead)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  释放互斥锁(mutex  MutexHandle)  (错误  error)  {}

// ff:
// mutex:
// err:
func ReleaseMutex(mutex Handle) (err error) {
	r1, _, e1 := syscall.Syscall(procReleaseMutex.Addr(), 1, uintptr(mutex), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  删除目录(path  *uint16)  (错误  error)  {}

// ff:
// path:
// err:
func RemoveDirectory(path *uint16) (err error) {
	r1, _, e1 := syscall.Syscall(procRemoveDirectoryW.Addr(), 1, uintptr(unsafe.Pointer(path)), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  移除Dll目录(cookie  uint64)  (错误  error)  {}

// ff:
// cookie:
// err:
func RemoveDllDirectory(cookie uintptr) (err error) {
	r1, _, e1 := syscall.Syscall(procRemoveDllDirectory.Addr(), 1, uintptr(cookie), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  重置事件(event  事件句柄)  (错误  error)  {}

// ff:
// event:
// err:
func ResetEvent(event Handle) (err error) {
	r1, _, e1 := syscall.Syscall(procResetEvent.Addr(), 1, uintptr(event), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func resizePseudoConsole(pconsole Handle, size uint32) (hr error) {
	r0, _, _ := syscall.Syscall(procResizePseudoConsole.Addr(), 2, uintptr(pconsole), uintptr(size), 0)
	if r0 != 0 {
		hr = syscall.Errno(r0)
	}
	return
}

// 翻译提示:func  恢复线程(thread  进程句柄)  (返回值  uint32,  错误  error)  {}

// ff:
// thread:
// ret:
// err:
func ResumeThread(thread Handle) (ret uint32, err error) {
	r0, _, e1 := syscall.Syscall(procResumeThread.Addr(), 1, uintptr(thread), 0, 0)
	ret = uint32(r0)
	if ret == 0xffffffff {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  设置串行通信暂停(handle  设备句柄)  (错误  error)  {}

// ff:
// handle:
// err:
func SetCommBreak(handle Handle) (err error) {
	r1, _, e1 := syscall.Syscall(procSetCommBreak.Addr(), 1, uintptr(handle), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  设置通信掩码(handle  文件句柄,  事件掩码  uint32)  (错误  error)  {}

// ff:
// handle:
// dwEvtMask:
// err:
func SetCommMask(handle Handle, dwEvtMask uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procSetCommMask.Addr(), 2, uintptr(handle), uintptr(dwEvtMask), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  设置串行通信状态(handle  文件句柄,  lpDCB  通信配置结构指针)  (错误  error)  {}

// ff:
// handle:
// lpDCB:
// err:
func SetCommState(handle Handle, lpDCB *DCB) (err error) {
	r1, _, e1 := syscall.Syscall(procSetCommState.Addr(), 2, uintptr(handle), uintptr(unsafe.Pointer(lpDCB)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  设置串行通信超时设置(handle  文件句柄,  timeouts  串行通信超时结构体指针)  (错误  error)  {}

// ff:
// handle:
// timeouts:
// err:
func SetCommTimeouts(handle Handle, timeouts *CommTimeouts) (err error) {
	r1, _, e1 := syscall.Syscall(procSetCommTimeouts.Addr(), 2, uintptr(handle), uintptr(unsafe.Pointer(timeouts)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func setConsoleCursorPosition(console Handle, position uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procSetConsoleCursorPosition.Addr(), 2, uintptr(console), uintptr(position), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  设置控制台模式(console  控制台句柄,  模式  uint32)  (错误  error)  {}

// ff:
// console:
// mode:
// err:
func SetConsoleMode(console Handle, mode uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procSetConsoleMode.Addr(), 2, uintptr(console), uintptr(mode), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  设置当前目录(path  *uint16)  (错误  error)  {}

// ff:
// path:
// err:
func SetCurrentDirectory(path *uint16) (err error) {
	r1, _, e1 := syscall.Syscall(procSetCurrentDirectoryW.Addr(), 1, uintptr(unsafe.Pointer(path)), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  设置默认DLL目录(directoryFlags  uint32)  (错误  error)  {}

// ff:
// directoryFlags:
// err:
func SetDefaultDllDirectories(directoryFlags uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procSetDefaultDllDirectories.Addr(), 1, uintptr(directoryFlags), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  设置Dll目录(path  string)  (错误  error)  {}

// ff:
// path:
// err:
func SetDllDirectory(path string) (err error) {
	var _p0 *uint16
	_p0, err = syscall.UTF16PtrFromString(path)
	if err != nil {
		return
	}
	return _SetDllDirectory(_p0)
}

func _SetDllDirectory(path *uint16) (err error) {
	r1, _, e1 := syscall.Syscall(procSetDllDirectoryW.Addr(), 1, uintptr(unsafe.Pointer(path)), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  设置文件结束符(handle  文件句柄)  (错误  error)  {}

// ff:
// handle:
// err:
func SetEndOfFile(handle Handle) (err error) {
	r1, _, e1 := syscall.Syscall(procSetEndOfFile.Addr(), 1, uintptr(handle), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  设置环境变量(name  *字节序列,  value  *字节序列)  (错误  error)  {}

// ff:
// name:
// value:
// err:
func SetEnvironmentVariable(name *uint16, value *uint16) (err error) {
	r1, _, e1 := syscall.Syscall(procSetEnvironmentVariableW.Addr(), 2, uintptr(unsafe.Pointer(name)), uintptr(unsafe.Pointer(value)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  设置错误模式(mode  uint32)  (返回值  uint32)  {}

// ff:
// mode:
// ret:
func SetErrorMode(mode uint32) (ret uint32) {
	r0, _, _ := syscall.Syscall(procSetErrorMode.Addr(), 1, uintptr(mode), 0, 0)
	ret = uint32(r0)
	return
}

// 翻译提示:func  设置事件(event  事件句柄)  (错误  error)  {}

// ff:
// event:
// err:
func SetEvent(event Handle) (err error) {
	r1, _, e1 := syscall.Syscall(procSetEvent.Addr(), 1, uintptr(event), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  设置文件属性(filename  *字符串编码,  attributes  文件属性)  (错误  error)  {}

// ff:
// name:
// attrs:
// err:
func SetFileAttributes(name *uint16, attrs uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procSetFileAttributesW.Addr(), 2, uintptr(unsafe.Pointer(name)), uintptr(attrs), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  设置文件完成通知模式(handle  文件句柄,  标志  uint8)  (错误  error)  {}

// ff:
// handle:
// flags:
// err:
func SetFileCompletionNotificationModes(handle Handle, flags uint8) (err error) {
	r1, _, e1 := syscall.Syscall(procSetFileCompletionNotificationModes.Addr(), 2, uintptr(handle), uintptr(flags), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  设置文件信息通过句柄(handle  文件句柄,  类型  uint32,  输入缓冲区  *字节,  输入缓冲区长度  uint32)  (错误  error)  {}

// ff:
// handle:
// class:
// inBuffer:
// inBufferLen:
// err:
func SetFileInformationByHandle(handle Handle, class uint32, inBuffer *byte, inBufferLen uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procSetFileInformationByHandle.Addr(), 4, uintptr(handle), uintptr(class), uintptr(unsafe.Pointer(inBuffer)), uintptr(inBufferLen), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  设置文件指针(handle  文件句柄,  低偏移量  int32,  高偏移量指针  *int32,  偏移依据  uint32)  (新低偏移量  uint32,  错误  error)  {}

// ff:
// handle:
// lowoffset:
// highoffsetptr:
// whence:
// newlowoffset:
// err:
func SetFilePointer(handle Handle, lowoffset int32, highoffsetptr *int32, whence uint32) (newlowoffset uint32, err error) {
	r0, _, e1 := syscall.Syscall6(procSetFilePointer.Addr(), 4, uintptr(handle), uintptr(lowoffset), uintptr(unsafe.Pointer(highoffsetptr)), uintptr(whence), 0, 0)
	newlowoffset = uint32(r0)
	if newlowoffset == 0xffffffff {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  设置文件时间(handle  文件句柄,  创建时间  *时间戳,  访问时间  *时间戳,  写入时间  *时间戳)  (错误  error)  {}

// ff:
// handle:
// ctime:
// atime:
// wtime:
// err:
func SetFileTime(handle Handle, ctime *Filetime, atime *Filetime, wtime *Filetime) (err error) {
	r1, _, e1 := syscall.Syscall6(procSetFileTime.Addr(), 4, uintptr(handle), uintptr(unsafe.Pointer(ctime)), uintptr(unsafe.Pointer(atime)), uintptr(unsafe.Pointer(wtime)), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  设置文件有效数据(handle  文件句柄,  validDataLength  有效数据长度  int64)  (错误  error)  {}

// ff:
// handle:
// validDataLength:
// err:
func SetFileValidData(handle Handle, validDataLength int64) (err error) {
	r1, _, e1 := syscall.Syscall(procSetFileValidData.Addr(), 2, uintptr(handle), uintptr(validDataLength), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  设置句柄信息(handle  句柄,  mask  信息掩码,  flags  标志位)  (错误  error)  {}

// ff:
// handle:
// mask:
// flags:
// err:
func SetHandleInformation(handle Handle, mask uint32, flags uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procSetHandleInformation.Addr(), 3, uintptr(handle), uintptr(mask), uintptr(flags))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  设置作业对象信息(job  执行句柄,  信息类别  uint32,  作业对象信息  unsafe.Pointer,  信息长度  uint32)  (返回值  int,  错误  error)  {}

// ff:
// job:
// JobObjectInformationClass:
// JobObjectInformation:
// JobObjectInformationLength:
// ret:
// err:
func SetInformationJobObject(job Handle, JobObjectInformationClass uint32, JobObjectInformation uintptr, JobObjectInformationLength uint32) (ret int, err error) {
	r0, _, e1 := syscall.Syscall6(procSetInformationJobObject.Addr(), 4, uintptr(job), uintptr(JobObjectInformationClass), uintptr(JobObjectInformation), uintptr(JobObjectInformationLength), 0, 0)
	ret = int(r0)
	if ret == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  设置命名管道状态(pipe  手柄,  状态  *uint32,  最大收集计数  *uint32,  数据收集超时  *uint32)  (错误  error)  {}  
// 
// 在这个函数中：
// -  `SetNamedPipeHandleState`  翻译成  `设置命名管道状态`
// -  `pipe`  保持原名，因为它是编程术语，表示管道句柄
// -  `state`  翻译成  `状态`
// -  `uint32`  是无符号32位整数类型，通常用于表示数值
// -  `maxCollectionCount`  翻译成  `最大收集计数`
// -  `collectDataTimeout`  翻译成  `数据收集超时`
// -  `err`  代表返回的错误信息，通常在Go中用于检查函数执行是否出错

// ff:
// pipe:
// state:
// maxCollectionCount:
// collectDataTimeout:
// err:
func SetNamedPipeHandleState(pipe Handle, state *uint32, maxCollectionCount *uint32, collectDataTimeout *uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procSetNamedPipeHandleState.Addr(), 4, uintptr(pipe), uintptr(unsafe.Pointer(state)), uintptr(unsafe.Pointer(maxCollectionCount)), uintptr(unsafe.Pointer(collectDataTimeout)), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  设置进程优先级类(process  进程句柄,  priorityClass  优先级类)  (错误  error)  {}

// ff:
// process:
// priorityClass:
// err:
func SetPriorityClass(process Handle, priorityClass uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procSetPriorityClass.Addr(), 2, uintptr(process), uintptr(priorityClass), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  设置进程优先级提升(process  进程句柄,  禁用  bool)  (错误  error)  {}

// ff:
// process:
// disable:
// err:
func SetProcessPriorityBoost(process Handle, disable bool) (err error) {
	var _p0 uint32
	if disable {
		_p0 = 1
	}
	r1, _, e1 := syscall.Syscall(procSetProcessPriorityBoost.Addr(), 2, uintptr(process), uintptr(_p0), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  设置进程关闭参数(level  uint32,  标志  uint32)  (错误  error)  {}

// ff:
// level:
// flags:
// err:
func SetProcessShutdownParameters(level uint32, flags uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procSetProcessShutdownParameters.Addr(), 2, uintptr(level), uintptr(flags), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  设置进程工作集大小扩展(hProcess  进程句柄,  最小工作集大小  uintptr,  最大工作集大小  uintptr,  标志  uint32)  (错误  error)  {}

// ff:
// hProcess:
// dwMinimumWorkingSetSize:
// dwMaximumWorkingSetSize:
// flags:
// err:
func SetProcessWorkingSetSizeEx(hProcess Handle, dwMinimumWorkingSetSize uintptr, dwMaximumWorkingSetSize uintptr, flags uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procSetProcessWorkingSetSizeEx.Addr(), 4, uintptr(hProcess), uintptr(dwMinimumWorkingSetSize), uintptr(dwMaximumWorkingSetSize), uintptr(flags), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  设置标准句柄(stdhandle:  uint32,  handle:  句柄)  (错误:  错误)  {}

// ff:
// stdhandle:
// handle:
// err:
func SetStdHandle(stdhandle uint32, handle Handle) (err error) {
	r1, _, e1 := syscall.Syscall(procSetStdHandle.Addr(), 2, uintptr(stdhandle), uintptr(handle), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  设置卷标(rootPathName  *字符串,  volumeName  *字符串)  (错误  error)  {}

// ff:
// rootPathName:
// volumeName:
// err:
func SetVolumeLabel(rootPathName *uint16, volumeName *uint16) (err error) {
	r1, _, e1 := syscall.Syscall(procSetVolumeLabelW.Addr(), 2, uintptr(unsafe.Pointer(rootPathName)), uintptr(unsafe.Pointer(volumeName)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  设置卷挂载点(volumeMountPoint  *uint16,  volumeName  *uint16)  (错误  error)  {}

// ff:
// volumeMountPoint:
// volumeName:
// err:
func SetVolumeMountPoint(volumeMountPoint *uint16, volumeName *uint16) (err error) {
	r1, _, e1 := syscall.Syscall(procSetVolumeMountPointW.Addr(), 2, uintptr(unsafe.Pointer(volumeMountPoint)), uintptr(unsafe.Pointer(volumeName)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  初始化串行通信(handle  手柄,  输入队列长度  uint32,  输出队列长度  uint32)  (错误  error)  {}

// ff:
// handle:
// dwInQueue:
// dwOutQueue:
// err:
func SetupComm(handle Handle, dwInQueue uint32, dwOutQueue uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procSetupComm.Addr(), 3, uintptr(handle), uintptr(dwInQueue), uintptr(dwOutQueue))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  资源大小(module  模块句柄,  resInfo  资源信息句柄)  (大小  uint32,  错误  error)  {}

// ff:
// module:
// resInfo:
// size:
// err:
func SizeofResource(module Handle, resInfo Handle) (size uint32, err error) {
	r0, _, e1 := syscall.Syscall(procSizeofResource.Addr(), 2, uintptr(module), uintptr(resInfo), 0)
	size = uint32(r0)
	if size == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  阻塞毫秒(millisecond  时间毫秒,  可中断  bool)  (返回值  时间毫秒)  {}

// ff:
// milliseconds:
// alertable:
// ret:
func SleepEx(milliseconds uint32, alertable bool) (ret uint32) {
	var _p0 uint32
	if alertable {
		_p0 = 1
	}
	r0, _, _ := syscall.Syscall(procSleepEx.Addr(), 2, uintptr(milliseconds), uintptr(_p0), 0)
	ret = uint32(r0)
	return
}

// 翻译提示:func  终止作业对象(job  作业句柄,  退出码  uint32)  (错误  error)  {}

// ff:
// job:
// exitCode:
// err:
func TerminateJobObject(job Handle, exitCode uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procTerminateJobObject.Addr(), 2, uintptr(job), uintptr(exitCode), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  结束进程(handle  进程句柄,  exitCode  退出码)  (错误  error)  {}

// ff:
// handle:
// exitcode:
// err:
func TerminateProcess(handle Handle, exitcode uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procTerminateProcess.Addr(), 2, uintptr(handle), uintptr(exitcode), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  线程32首个(snapshot  进程快照,  threadEntry  *线程条目32)  (错误  error)  {}

// ff:
// snapshot:
// threadEntry:
// err:
func Thread32First(snapshot Handle, threadEntry *ThreadEntry32) (err error) {
	r1, _, e1 := syscall.Syscall(procThread32First.Addr(), 2, uintptr(snapshot), uintptr(unsafe.Pointer(threadEntry)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  下一个线程32(snapshot  进程快照句柄,  线程条目  *线程条目32)  (错误  error)  {}

// ff:
// snapshot:
// threadEntry:
// err:
func Thread32Next(snapshot Handle, threadEntry *ThreadEntry32) (err error) {
	r1, _, e1 := syscall.Syscall(procThread32Next.Addr(), 2, uintptr(snapshot), uintptr(unsafe.Pointer(threadEntry)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  解锁文件扩展(file  文件句柄,  保留值  uint32,  低字节  uint32,  高字节  uint32,  重叠结构  *重叠结构体)  (错误  error)  {}

// ff:
// file:
// reserved:
// bytesLow:
// bytesHigh:
// overlapped:
// err:
func UnlockFileEx(file Handle, reserved uint32, bytesLow uint32, bytesHigh uint32, overlapped *Overlapped) (err error) {
	r1, _, e1 := syscall.Syscall6(procUnlockFileEx.Addr(), 5, uintptr(file), uintptr(reserved), uintptr(bytesLow), uintptr(bytesHigh), uintptr(unsafe.Pointer(overlapped)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  取消映射视图文件(addr  内存地址)  (错误  error)  {}

// ff:
// addr:
// err:
func UnmapViewOfFile(addr uintptr) (err error) {
	r1, _, e1 := syscall.Syscall(procUnmapViewOfFile.Addr(), 1, uintptr(addr), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func updateProcThreadAttribute(attrlist *ProcThreadAttributeList, flags uint32, attr uintptr, value unsafe.Pointer, size uintptr, prevvalue unsafe.Pointer, returnedsize *uintptr) (err error) {
	r1, _, e1 := syscall.Syscall9(procUpdateProcThreadAttribute.Addr(), 7, uintptr(unsafe.Pointer(attrlist)), uintptr(flags), uintptr(attr), uintptr(value), uintptr(size), uintptr(prevvalue), uintptr(unsafe.Pointer(returnedsize)), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  虚拟分配内存(address  uintptr,  大小  uintptr,  分配类型  uint32,  保护模式  uint32)  (分配结果  uintptr,  错误  error)  {}

// ff:
// address:
// size:
// alloctype:
// protect:
// value:
// err:
func VirtualAlloc(address uintptr, size uintptr, alloctype uint32, protect uint32) (value uintptr, err error) {
	r0, _, e1 := syscall.Syscall6(procVirtualAlloc.Addr(), 4, uintptr(address), uintptr(size), uintptr(alloctype), uintptr(protect), 0, 0)
	value = uintptr(r0)
	if value == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  虚拟释放内存(address  uint64,  size  uint64,  释放类型  uint32)  (错误  error)  {}

// ff:
// address:
// size:
// freetype:
// err:
func VirtualFree(address uintptr, size uintptr, freetype uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procVirtualFree.Addr(), 3, uintptr(address), uintptr(size), uintptr(freetype))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  虚拟锁定内存(addr  uint64,  长度  uint64)  (错误  error)  {}

// ff:
// addr:
// length:
// err:
func VirtualLock(addr uintptr, length uintptr) (err error) {
	r1, _, e1 := syscall.Syscall(procVirtualLock.Addr(), 2, uintptr(addr), uintptr(length), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  虚拟保护(address  uintptr,  size  uintptr,  新保护  uint32,  旧保护  *uint32)  (错误  error)  {}

// ff:
// address:
// size:
// newprotect:
// oldprotect:
// err:
func VirtualProtect(address uintptr, size uintptr, newprotect uint32, oldprotect *uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procVirtualProtect.Addr(), 4, uintptr(address), uintptr(size), uintptr(newprotect), uintptr(unsafe.Pointer(oldprotect)), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  虚拟保护修改(process  进程句柄,  address  内存地址,  size  内存大小,  newProtect  新保护模式,  oldProtect  *原保护模式)  (错误  error)  {}

// ff:
// process:
// address:
// size:
// newProtect:
// oldProtect:
// err:
func VirtualProtectEx(process Handle, address uintptr, size uintptr, newProtect uint32, oldProtect *uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procVirtualProtectEx.Addr(), 5, uintptr(process), uintptr(address), uintptr(size), uintptr(newProtect), uintptr(unsafe.Pointer(oldProtect)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  虚拟查询(address  uint64,  buffer  *内存基础信息,  length  uint64)  (错误  error)  {}

// ff:
// address:
// buffer:
// length:
// err:
func VirtualQuery(address uintptr, buffer *MemoryBasicInformation, length uintptr) (err error) {
	r1, _, e1 := syscall.Syscall(procVirtualQuery.Addr(), 3, uintptr(address), uintptr(unsafe.Pointer(buffer)), uintptr(length))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  虚拟查询扩展(process  进程句柄,  address  uintptr,  buffer  *内存基础信息,  length  uintptr)  (错误  error)  {}

// ff:
// process:
// address:
// buffer:
// length:
// err:
func VirtualQueryEx(process Handle, address uintptr, buffer *MemoryBasicInformation, length uintptr) (err error) {
	r1, _, e1 := syscall.Syscall6(procVirtualQueryEx.Addr(), 4, uintptr(process), uintptr(address), uintptr(unsafe.Pointer(buffer)), uintptr(length), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  虚拟解锁(address  uintptr,  长度  uintptr)  (错误  error)  {}

// ff:
// addr:
// length:
// err:
func VirtualUnlock(addr uintptr, length uintptr) (err error) {
	r1, _, e1 := syscall.Syscall(procVirtualUnlock.Addr(), 2, uintptr(addr), uintptr(length), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  获取活跃控制台会话ID()  (会话ID  uint32)  {}

// ff:
// sessionID:
func WTSGetActiveConsoleSessionId() (sessionID uint32) {
	r0, _, _ := syscall.Syscall(procWTSGetActiveConsoleSessionId.Addr(), 0, 0, 0, 0)
	sessionID = uint32(r0)
	return
}

// 翻译提示:func  等待通讯事件(handle  文件句柄,  lpEvtMask  事件掩码指针,  lpOverlapped  覆盖式重叠结构体指针)  (错误  error)  {}

// ff:
// handle:
// lpEvtMask:
// lpOverlapped:
// err:
func WaitCommEvent(handle Handle, lpEvtMask *uint32, lpOverlapped *Overlapped) (err error) {
	r1, _, e1 := syscall.Syscall(procWaitCommEvent.Addr(), 3, uintptr(handle), uintptr(unsafe.Pointer(lpEvtMask)), uintptr(unsafe.Pointer(lpOverlapped)))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func waitForMultipleObjects(count uint32, handles uintptr, waitAll bool, waitMilliseconds uint32) (event uint32, err error) {
	var _p0 uint32
	if waitAll {
		_p0 = 1
	}
	r0, _, e1 := syscall.Syscall6(procWaitForMultipleObjects.Addr(), 4, uintptr(count), uintptr(handles), uintptr(_p0), uintptr(waitMilliseconds), 0, 0)
	event = uint32(r0)
	if event == 0xffffffff {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  等待单个对象(handle  文件句柄,  等待毫秒  uint32)  (事件状态  uint32,  错误  error)  {}

// ff:
// handle:
// waitMilliseconds:
// event:
// err:
func WaitForSingleObject(handle Handle, waitMilliseconds uint32) (event uint32, err error) {
	r0, _, e1 := syscall.Syscall(procWaitForSingleObject.Addr(), 2, uintptr(handle), uintptr(waitMilliseconds), 0)
	event = uint32(r0)
	if event == 0xffffffff {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  写控制台(console  控制台句柄,  buf  *uint16,  要写入的字节数  uint32,  已写入字节数  *uint32,  保留  *byte)  (错误  error)  {}

// ff:
// console:
// buf:
// towrite:
// written:
// reserved:
// err:
func WriteConsole(console Handle, buf *uint16, towrite uint32, written *uint32, reserved *byte) (err error) {
	r1, _, e1 := syscall.Syscall6(procWriteConsoleW.Addr(), 5, uintptr(console), uintptr(unsafe.Pointer(buf)), uintptr(towrite), uintptr(unsafe.Pointer(written)), uintptr(unsafe.Pointer(reserved)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func writeFile(handle Handle, buf []byte, done *uint32, overlapped *Overlapped) (err error) {
	var _p0 *byte
	if len(buf) > 0 {
		_p0 = &buf[0]
	}
	r1, _, e1 := syscall.Syscall6(procWriteFile.Addr(), 5, uintptr(handle), uintptr(unsafe.Pointer(_p0)), uintptr(len(buf)), uintptr(unsafe.Pointer(done)), uintptr(unsafe.Pointer(overlapped)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  写入进程内存(process  进程句柄,  基地址  uintptr,  缓冲区  *byte,  缓冲区大小  uintptr,  实际写入字节数  *uintptr)  (错误  error)  {}

// ff:
// process:
// baseAddress:
// buffer:
// size:
// numberOfBytesWritten:
// err:
func WriteProcessMemory(process Handle, baseAddress uintptr, buffer *byte, size uintptr, numberOfBytesWritten *uintptr) (err error) {
	r1, _, e1 := syscall.Syscall6(procWriteProcessMemory.Addr(), 5, uintptr(process), uintptr(baseAddress), uintptr(unsafe.Pointer(buffer)), uintptr(size), uintptr(unsafe.Pointer(numberOfBytesWritten)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  接受扩展(ls  手柄,  as  手柄,  缓冲区  *byte,  接收数据长度  uint32,  左侧地址长度  uint32,  右侧地址长度  uint32,  已接收  *uint32,  重叠  *重叠)  (错误  error)  {}

// ff:
// ls:
// as:
// buf:
// rxdatalen:
// laddrlen:
// raddrlen:
// recvd:
// overlapped:
// err:
func AcceptEx(ls Handle, as Handle, buf *byte, rxdatalen uint32, laddrlen uint32, raddrlen uint32, recvd *uint32, overlapped *Overlapped) (err error) {
	r1, _, e1 := syscall.Syscall9(procAcceptEx.Addr(), 8, uintptr(ls), uintptr(as), uintptr(unsafe.Pointer(buf)), uintptr(rxdatalen), uintptr(laddrlen), uintptr(raddrlen), uintptr(unsafe.Pointer(recvd)), uintptr(unsafe.Pointer(overlapped)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  获取接受扩展套接字地址(buf  *字节,  接收数据长度  uint32,  左地址长度  uint32,  右地址长度  uint32,  左原始套接字地址  **原始套接字地址任意,  左地址长度指针  *int32,  右原始套接字地址  **原始套接字地址任意,  右地址长度指针  *int32)  {}

// ff:
// buf:
// rxdatalen:
// laddrlen:
// raddrlen:
// lrsa:
// lrsalen:
// rrsa:
// rrsalen:
func GetAcceptExSockaddrs(buf *byte, rxdatalen uint32, laddrlen uint32, raddrlen uint32, lrsa **RawSockaddrAny, lrsalen *int32, rrsa **RawSockaddrAny, rrsalen *int32) {
	syscall.Syscall9(procGetAcceptExSockaddrs.Addr(), 8, uintptr(unsafe.Pointer(buf)), uintptr(rxdatalen), uintptr(laddrlen), uintptr(raddrlen), uintptr(unsafe.Pointer(lrsa)), uintptr(unsafe.Pointer(lrsalen)), uintptr(unsafe.Pointer(rrsa)), uintptr(unsafe.Pointer(rrsalen)), 0)
	return
}

// 翻译提示:func  传输文件(s  文件句柄,  处理句柄  Handle,  写入字节数  uint32,  每次发送字节数  uint32,  重叠结构  *重叠结构,  传输文件缓存  *传输文件缓冲区,  标志  uint32)  (错误  error)  {}

// ff:
// s:
// handle:
// bytesToWrite:
// bytsPerSend:
// overlapped:
// transmitFileBuf:
// flags:
// err:
func TransmitFile(s Handle, handle Handle, bytesToWrite uint32, bytsPerSend uint32, overlapped *Overlapped, transmitFileBuf *TransmitFileBuffers, flags uint32) (err error) {
	r1, _, e1 := syscall.Syscall9(procTransmitFile.Addr(), 7, uintptr(s), uintptr(handle), uintptr(bytesToWrite), uintptr(bytsPerSend), uintptr(unsafe.Pointer(overlapped)), uintptr(unsafe.Pointer(transmitFileBuf)), uintptr(flags), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  网络缓冲区释放(buf  *字节)  (网络错误  error)  {}

// ff:
// buf:
// neterr:
func NetApiBufferFree(buf *byte) (neterr error) {
	r0, _, _ := syscall.Syscall(procNetApiBufferFree.Addr(), 1, uintptr(unsafe.Pointer(buf)), 0, 0)
	if r0 != 0 {
		neterr = syscall.Errno(r0)
	}
	return
}

// 翻译提示:func  获取网络连接信息(server  *字节串,  name  **字节串,  bufferType  *uint32)  (网络错误  error)  {}

// ff:
// server:
// name:
// bufType:
// neterr:
func NetGetJoinInformation(server *uint16, name **uint16, bufType *uint32) (neterr error) {
	r0, _, _ := syscall.Syscall(procNetGetJoinInformation.Addr(), 3, uintptr(unsafe.Pointer(server)), uintptr(unsafe.Pointer(name)), uintptr(unsafe.Pointer(bufType)))
	if r0 != 0 {
		neterr = syscall.Errno(r0)
	}
	return
}

// 翻译提示:func  获取用户信息(serverName  *字符串,  userName  *字符串,  信息级别  uint32,  数据缓冲区  **字节)  (网络错误  error)  {}

// ff:
// serverName:
// userName:
// level:
// buf:
// neterr:
func NetUserGetInfo(serverName *uint16, userName *uint16, level uint32, buf **byte) (neterr error) {
	r0, _, _ := syscall.Syscall6(procNetUserGetInfo.Addr(), 4, uintptr(unsafe.Pointer(serverName)), uintptr(unsafe.Pointer(userName)), uintptr(level), uintptr(unsafe.Pointer(buf)), 0, 0)
	if r0 != 0 {
		neterr = syscall.Errno(r0)
	}
	return
}

// 翻译提示:func  创建文件(handle  *文件句柄,  访问权限  uint32,  对象属性  *对象属性,  io状态块  *IO_STATUS_BLOCK,  分配大小  *int64,  属性  uint32,  共享模式  uint32,  处理方式  uint32,  选项  uint32,  扩展属性缓冲区  uintptr,  扩展属性长度  uint32)  (NTSTATUS  错误)  {}

// ff:
// handle:
// access:
// oa:
// iosb:
// allocationSize:
// attributes:
// share:
// disposition:
// options:
// eabuffer:
// ealength:
// ntstatus:
func NtCreateFile(handle *Handle, access uint32, oa *OBJECT_ATTRIBUTES, iosb *IO_STATUS_BLOCK, allocationSize *int64, attributes uint32, share uint32, disposition uint32, options uint32, eabuffer uintptr, ealength uint32) (ntstatus error) {
	r0, _, _ := syscall.Syscall12(procNtCreateFile.Addr(), 11, uintptr(unsafe.Pointer(handle)), uintptr(access), uintptr(unsafe.Pointer(oa)), uintptr(unsafe.Pointer(iosb)), uintptr(unsafe.Pointer(allocationSize)), uintptr(attributes), uintptr(share), uintptr(disposition), uintptr(options), uintptr(eabuffer), uintptr(ealength), 0)
	if r0 != 0 {
		ntstatus = NTStatus(r0)
	}
	return
}

// 翻译提示:func  创建命名管道(pipe  *句柄,  访问权限  uint32,  对象属性  *OBJECT_ATTRIBUTES,  io状态块  *IO_STATUS_BLOCK,  共享模式  uint32,  处理方式  uint32,  选项  uint32,  管道类型  uint32,  读取模式  uint32,  完成模式  uint32,  最大实例数  uint32,  输入配额  uint32,  输出配额  uint32,  超时时间  *int64)  (nt状态错误)  {}

// ff:
// pipe:
// access:
// oa:
// iosb:
// share:
// disposition:
// options:
// typ:
// readMode:
// completionMode:
// maxInstances:
// inboundQuota:
// outputQuota:
// timeout:
// ntstatus:
func NtCreateNamedPipeFile(pipe *Handle, access uint32, oa *OBJECT_ATTRIBUTES, iosb *IO_STATUS_BLOCK, share uint32, disposition uint32, options uint32, typ uint32, readMode uint32, completionMode uint32, maxInstances uint32, inboundQuota uint32, outputQuota uint32, timeout *int64) (ntstatus error) {
	r0, _, _ := syscall.Syscall15(procNtCreateNamedPipeFile.Addr(), 14, uintptr(unsafe.Pointer(pipe)), uintptr(access), uintptr(unsafe.Pointer(oa)), uintptr(unsafe.Pointer(iosb)), uintptr(share), uintptr(disposition), uintptr(options), uintptr(typ), uintptr(readMode), uintptr(completionMode), uintptr(maxInstances), uintptr(inboundQuota), uintptr(outputQuota), uintptr(unsafe.Pointer(timeout)), 0)
	if r0 != 0 {
		ntstatus = NTStatus(r0)
	}
	return
}

// 翻译提示:func  查询进程信息(processHandle  手柄,  processInfoClass  进程信息类,  processInfo  进程信息指针,  processInfoLen  进程信息长度  uint32,  retLen  返回长度指针  *uint32)  (ntstatus  错误)  {}

// ff:
// proc:
// procInfoClass:
// procInfo:
// procInfoLen:
// retLen:
// ntstatus:
func NtQueryInformationProcess(proc Handle, procInfoClass int32, procInfo unsafe.Pointer, procInfoLen uint32, retLen *uint32) (ntstatus error) {
	r0, _, _ := syscall.Syscall6(procNtQueryInformationProcess.Addr(), 5, uintptr(proc), uintptr(procInfoClass), uintptr(procInfo), uintptr(procInfoLen), uintptr(unsafe.Pointer(retLen)), 0)
	if r0 != 0 {
		ntstatus = NTStatus(r0)
	}
	return
}

// 翻译提示:func  查询系统信息(sysInfo类别  int32,  系统信息  unsafe.Pointer,  系统信息长度  uint32,  返回长度  *uint32)  (NTSTATUS  错误)  {}

// ff:
// sysInfoClass:
// sysInfo:
// sysInfoLen:
// retLen:
// ntstatus:
func NtQuerySystemInformation(sysInfoClass int32, sysInfo unsafe.Pointer, sysInfoLen uint32, retLen *uint32) (ntstatus error) {
	r0, _, _ := syscall.Syscall6(procNtQuerySystemInformation.Addr(), 4, uintptr(sysInfoClass), uintptr(sysInfo), uintptr(sysInfoLen), uintptr(unsafe.Pointer(retLen)), 0, 0)
	if r0 != 0 {
		ntstatus = NTStatus(r0)
	}
	return
}

// 翻译提示:func  设置信息文件(handle  文件句柄,  iosb  输入输出状态块指针,  inBuffer  输入缓冲区指针,  inBufferLen  输入缓冲区长度,  class  文件信息类别)  (ntstatus  错误)  {}

// ff:
// handle:
// iosb:
// inBuffer:
// inBufferLen:
// class:
// ntstatus:
func NtSetInformationFile(handle Handle, iosb *IO_STATUS_BLOCK, inBuffer *byte, inBufferLen uint32, class uint32) (ntstatus error) {
	r0, _, _ := syscall.Syscall6(procNtSetInformationFile.Addr(), 5, uintptr(handle), uintptr(unsafe.Pointer(iosb)), uintptr(unsafe.Pointer(inBuffer)), uintptr(inBufferLen), uintptr(class), 0)
	if r0 != 0 {
		ntstatus = NTStatus(r0)
	}
	return
}

// 翻译提示:func  设置进程信息(process  手柄,  进程信息类别  int32,  进程信息  unsafe.Pointer,  进程信息长度  uint32)  (NTSTATUS  错误)  {}

// ff:
// proc:
// procInfoClass:
// procInfo:
// procInfoLen:
// ntstatus:
func NtSetInformationProcess(proc Handle, procInfoClass int32, procInfo unsafe.Pointer, procInfoLen uint32) (ntstatus error) {
	r0, _, _ := syscall.Syscall6(procNtSetInformationProcess.Addr(), 4, uintptr(proc), uintptr(procInfoClass), uintptr(procInfo), uintptr(procInfoLen), 0, 0)
	if r0 != 0 {
		ntstatus = NTStatus(r0)
	}
	return
}

// 翻译提示:func  设置系统信息(sysInfo类别  int32,  系统信息指针  unsafe.Pointer,  系统信息长度  uint32)  (NTSTATUS  错误)  {}

// ff:
// sysInfoClass:
// sysInfo:
// sysInfoLen:
// ntstatus:
func NtSetSystemInformation(sysInfoClass int32, sysInfo unsafe.Pointer, sysInfoLen uint32) (ntstatus error) {
	r0, _, _ := syscall.Syscall(procNtSetSystemInformation.Addr(), 3, uintptr(sysInfoClass), uintptr(sysInfo), uintptr(sysInfoLen))
	if r0 != 0 {
		ntstatus = NTStatus(r0)
	}
	return
}

// 翻译提示:func  添加函数表(functionTable  *运行时函数表,  entryCount  uint32,  baseAddress  uintptr)  (返回值  bool)  {}

// ff:
// functionTable:
// entryCount:
// baseAddress:
// ret:
func RtlAddFunctionTable(functionTable *RUNTIME_FUNCTION, entryCount uint32, baseAddress uintptr) (ret bool) {
	r0, _, _ := syscall.Syscall(procRtlAddFunctionTable.Addr(), 3, uintptr(unsafe.Pointer(functionTable)), uintptr(entryCount), uintptr(baseAddress))
	ret = r0 != 0
	return
}

// 翻译提示:func  默认NPACL(acl  **ACL)  (系统状态错误  error)  {}

// ff:
// acl:
// ntstatus:
func RtlDefaultNpAcl(acl **ACL) (ntstatus error) {
	r0, _, _ := syscall.Syscall(procRtlDefaultNpAcl.Addr(), 1, uintptr(unsafe.Pointer(acl)), 0, 0)
	if r0 != 0 {
		ntstatus = NTStatus(r0)
	}
	return
}

// 翻译提示:func  删除函数表(functionTable  *运行时函数表)  (成功  bool)  {}

// ff:
// functionTable:
// ret:
func RtlDeleteFunctionTable(functionTable *RUNTIME_FUNCTION) (ret bool) {
	r0, _, _ := syscall.Syscall(procRtlDeleteFunctionTable.Addr(), 1, uintptr(unsafe.Pointer(functionTable)), 0, 0)
	ret = r0 != 0
	return
}

// 翻译提示:func  Dos路径名转Nt路径名(dosName  *uint16,  ntName  *NTUnicodeString,  ntFileNamePart  *uint16,  relativeName  *RTL_RELATIVE_NAME)  (ntstatus  error)  {}

// ff:
// dosName:
// ntName:
// ntFileNamePart:
// relativeName:
// ntstatus:
func RtlDosPathNameToNtPathName(dosName *uint16, ntName *NTUnicodeString, ntFileNamePart *uint16, relativeName *RTL_RELATIVE_NAME) (ntstatus error) {
	r0, _, _ := syscall.Syscall6(procRtlDosPathNameToNtPathName_U_WithStatus.Addr(), 4, uintptr(unsafe.Pointer(dosName)), uintptr(unsafe.Pointer(ntName)), uintptr(unsafe.Pointer(ntFileNamePart)), uintptr(unsafe.Pointer(relativeName)), 0, 0)
	if r0 != 0 {
		ntstatus = NTStatus(r0)
	}
	return
}

// 翻译提示:func  RtlDOS路径名转相对NT路径名(dosName  *uint16,  ntName  *NTUnicodeString,  ntFileNamePart  *uint16,  relativeName  *RTL_RELATIVE_NAME)  (ntstatus  error)  {}

// ff:
// dosName:
// ntName:
// ntFileNamePart:
// relativeName:
// ntstatus:
func RtlDosPathNameToRelativeNtPathName(dosName *uint16, ntName *NTUnicodeString, ntFileNamePart *uint16, relativeName *RTL_RELATIVE_NAME) (ntstatus error) {
	r0, _, _ := syscall.Syscall6(procRtlDosPathNameToRelativeNtPathName_U_WithStatus.Addr(), 4, uintptr(unsafe.Pointer(dosName)), uintptr(unsafe.Pointer(ntName)), uintptr(unsafe.Pointer(ntFileNamePart)), uintptr(unsafe.Pointer(relativeName)), 0, 0)
	if r0 != 0 {
		ntstatus = NTStatus(r0)
	}
	return
}

// 翻译提示:func  获取当前进程环境块()  (进程环境块  *PEB)  {}

// ff:
// peb:
func RtlGetCurrentPeb() (peb *PEB) {
	r0, _, _ := syscall.Syscall(procRtlGetCurrentPeb.Addr(), 0, 0, 0, 0)
	peb = (*PEB)(unsafe.Pointer(r0))
	return
}

func rtlGetNtVersionNumbers(majorVersion *uint32, minorVersion *uint32, buildNumber *uint32) {
	syscall.Syscall(procRtlGetNtVersionNumbers.Addr(), 3, uintptr(unsafe.Pointer(majorVersion)), uintptr(unsafe.Pointer(minorVersion)), uintptr(unsafe.Pointer(buildNumber)))
	return
}

func rtlGetVersion(info *OsVersionInfoEx) (ntstatus error) {
	r0, _, _ := syscall.Syscall(procRtlGetVersion.Addr(), 1, uintptr(unsafe.Pointer(info)), 0, 0)
	if r0 != 0 {
		ntstatus = NTStatus(r0)
	}
	return
}

// 翻译提示:func  初始化字符串(destinationString  字符串目标指针,  sourceString  字符源指针)  {}

// ff:
// destinationString:
// sourceString:
func RtlInitString(destinationString *NTString, sourceString *byte) {
	syscall.Syscall(procRtlInitString.Addr(), 2, uintptr(unsafe.Pointer(destinationString)), uintptr(unsafe.Pointer(sourceString)), 0)
	return
}

// 翻译提示:func  初始化Unicode字符串(destinationString  *Unicode字符串目标,  sourceString  *宽字符指针)  {}

// ff:
// destinationString:
// sourceString:
func RtlInitUnicodeString(destinationString *NTUnicodeString, sourceString *uint16) {
	syscall.Syscall(procRtlInitUnicodeString.Addr(), 2, uintptr(unsafe.Pointer(destinationString)), uintptr(unsafe.Pointer(sourceString)), 0)
	return
}

func rtlNtStatusToDosErrorNoTeb(ntstatus NTStatus) (ret syscall.Errno) {
	r0, _, _ := syscall.Syscall(procRtlNtStatusToDosErrorNoTeb.Addr(), 1, uintptr(ntstatus), 0, 0)
	ret = syscall.Errno(r0)
	return
}

func clsidFromString(lpsz *uint16, pclsid *GUID) (ret error) {
	r0, _, _ := syscall.Syscall(procCLSIDFromString.Addr(), 2, uintptr(unsafe.Pointer(lpsz)), uintptr(unsafe.Pointer(pclsid)), 0)
	if r0 != 0 {
		ret = syscall.Errno(r0)
	}
	return
}

func coCreateGuid(pguid *GUID) (ret error) {
	r0, _, _ := syscall.Syscall(procCoCreateGuid.Addr(), 1, uintptr(unsafe.Pointer(pguid)), 0, 0)
	if r0 != 0 {
		ret = syscall.Errno(r0)
	}
	return
}

// 翻译提示:func  获取COM对象(name  *uint16,  绑定选项  *BIND_OPTS3,  接口标识符  *GUID,  函数表  **uintptr)  (错误  error)  {}

// ff:
// name:
// bindOpts:
// guid:
// functionTable:
// ret:
func CoGetObject(name *uint16, bindOpts *BIND_OPTS3, guid *GUID, functionTable **uintptr) (ret error) {
	r0, _, _ := syscall.Syscall6(procCoGetObject.Addr(), 4, uintptr(unsafe.Pointer(name)), uintptr(unsafe.Pointer(bindOpts)), uintptr(unsafe.Pointer(guid)), uintptr(unsafe.Pointer(functionTable)), 0, 0)
	if r0 != 0 {
		ret = syscall.Errno(r0)
	}
	return
}

// 翻译提示:func  初始化Com库(reserved  保留参数,  coInit  初始化模式)  (返回值  error)  {}

// ff:
// reserved:
// coInit:
// ret:
func CoInitializeEx(reserved uintptr, coInit uint32) (ret error) {
	r0, _, _ := syscall.Syscall(procCoInitializeEx.Addr(), 2, uintptr(reserved), uintptr(coInit), 0)
	if r0 != 0 {
		ret = syscall.Errno(r0)
	}
	return
}

// 翻译提示:func  垃圾回收内存(address  安全指针)  {}

// ff:
// address:
func CoTaskMemFree(address unsafe.Pointer) {
	syscall.Syscall(procCoTaskMemFree.Addr(), 1, uintptr(address), 0, 0)
	return
}

// 翻译提示:func  卸载COM()  {}

// ff:
func CoUninitialize() {
	syscall.Syscall(procCoUninitialize.Addr(), 0, 0, 0, 0)
	return
}

func stringFromGUID2(rguid *GUID, lpsz *uint16, cchMax int32) (chars int32) {
	r0, _, _ := syscall.Syscall(procStringFromGUID2.Addr(), 3, uintptr(unsafe.Pointer(rguid)), uintptr(unsafe.Pointer(lpsz)), uintptr(cchMax))
	chars = int32(r0)
	return
}

// 翻译提示:func  列举进程模块(process  进程句柄,  module  模块指针,  cb  数据大小,  cbNeeded  需要数据大小指针)  (错误  error)  {}

// ff:
// process:
// module:
// cb:
// cbNeeded:
// err:
func EnumProcessModules(process Handle, module *Handle, cb uint32, cbNeeded *uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procEnumProcessModules.Addr(), 4, uintptr(process), uintptr(unsafe.Pointer(module)), uintptr(cb), uintptr(unsafe.Pointer(cbNeeded)), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  列举进程模块扩展(process  进程句柄,  module  模块指针,  cb  数据大小,  cbNeeded  需要数据大小指针,  filterFlag  过滤标志)  (err  错误)  {}

// ff:
// process:
// module:
// cb:
// cbNeeded:
// filterFlag:
// err:
func EnumProcessModulesEx(process Handle, module *Handle, cb uint32, cbNeeded *uint32, filterFlag uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procEnumProcessModulesEx.Addr(), 5, uintptr(process), uintptr(unsafe.Pointer(module)), uintptr(cb), uintptr(unsafe.Pointer(cbNeeded)), uintptr(filterFlag), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func enumProcesses(processIds *uint32, nSize uint32, bytesReturned *uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procEnumProcesses.Addr(), 3, uintptr(unsafe.Pointer(processIds)), uintptr(nSize), uintptr(unsafe.Pointer(bytesReturned)))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  获取模块基名(process  进程句柄,  module  模块句柄,  baseName  模块名称缓冲区指针,  size  名称缓冲区大小)  (error  错误信息)  {}

// ff:
// process:
// module:
// baseName:
// size:
// err:
func GetModuleBaseName(process Handle, module Handle, baseName *uint16, size uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procGetModuleBaseNameW.Addr(), 4, uintptr(process), uintptr(module), uintptr(unsafe.Pointer(baseName)), uintptr(size), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  获取模块文件名Ex(process  进程句柄,  module  模块句柄,  filename  文件名指针,  size  字节数)  (错误  error)  {}

// ff:
// process:
// module:
// filename:
// size:
// err:
func GetModuleFileNameEx(process Handle, module Handle, filename *uint16, size uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procGetModuleFileNameExW.Addr(), 4, uintptr(process), uintptr(module), uintptr(unsafe.Pointer(filename)), uintptr(size), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  获取模块信息(process  进程句柄,  module  模块句柄,  modinfo  模块信息指针,  cb  字节数)  (错误  error)  {}

// ff:
// process:
// module:
// modinfo:
// cb:
// err:
func GetModuleInformation(process Handle, module Handle, modinfo *ModuleInfo, cb uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procGetModuleInformation.Addr(), 4, uintptr(process), uintptr(module), uintptr(unsafe.Pointer(modinfo)), uintptr(cb), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  查询工作集扩展(process  进程句柄,  pv  未指定类型指针,  cb  字节数)  (错误  error)  {}

// ff:
// process:
// pv:
// cb:
// err:
func QueryWorkingSetEx(process Handle, pv uintptr, cb uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procQueryWorkingSetEx.Addr(), 3, uintptr(process), uintptr(pv), uintptr(cb))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  订阅服务变更通知(service句柄,  事件类型  uint32,  回调函数指针  uintptr,  回调上下文指针  uintptr,  订阅指针  *uintptr)  (错误  error)  {}

// ff:
// service:
// eventType:
// callback:
// callbackCtx:
// subscription:
// ret:
func SubscribeServiceChangeNotifications(service Handle, eventType uint32, callback uintptr, callbackCtx uintptr, subscription *uintptr) (ret error) {
	ret = procSubscribeServiceChangeNotifications.Find()
	if ret != nil {
		return
	}
	r0, _, _ := syscall.Syscall6(procSubscribeServiceChangeNotifications.Addr(), 5, uintptr(service), uintptr(eventType), uintptr(callback), uintptr(callbackCtx), uintptr(unsafe.Pointer(subscription)), 0)
	if r0 != 0 {
		ret = syscall.Errno(r0)
	}
	return
}

// 翻译提示:func  取消订阅服务变更通知(subscription  uintptr)  (错误  error)  {}

// ff:
// subscription:
// err:
func UnsubscribeServiceChangeNotifications(subscription uintptr) (err error) {
	err = procUnsubscribeServiceChangeNotifications.Find()
	if err != nil {
		return
	}
	syscall.Syscall(procUnsubscribeServiceChangeNotifications.Addr(), 1, uintptr(subscription), 0, 0)
	return
}

// 翻译提示:func  获取用户名称扩展(name格式  uint32,  名称缓存  *uint16,  名称大小  *uint32)  (错误  error)  {}

// ff:
// nameFormat:
// nameBuffre:
// nSize:
// err:
func GetUserNameEx(nameFormat uint32, nameBuffre *uint16, nSize *uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procGetUserNameExW.Addr(), 3, uintptr(nameFormat), uintptr(unsafe.Pointer(nameBuffre)), uintptr(unsafe.Pointer(nSize)))
	if r1&0xff == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  翻译账号名(accName  *uint16,  原始格式  uint32,  目标格式  uint32,  转换后名字  *uint16,  名字长度  *uint32)  (错误  error)  {}

// ff:
// accName:
// accNameFormat:
// desiredNameFormat:
// translatedName:
// nSize:
// err:
func TranslateName(accName *uint16, accNameFormat uint32, desiredNameFormat uint32, translatedName *uint16, nSize *uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procTranslateNameW.Addr(), 5, uintptr(unsafe.Pointer(accName)), uintptr(accNameFormat), uintptr(desiredNameFormat), uintptr(unsafe.Pointer(translatedName)), uintptr(unsafe.Pointer(nSize)), 0)
	if r1&0xff == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  初始化驱动信息列表(deviceInfoSet  设备信息集,  deviceInfoData  设备信息数据指针,  driverType  驱动类型)  (错误  error)  {}

// ff:
// deviceInfoSet:
// deviceInfoData:
// driverType:
// err:
func SetupDiBuildDriverInfoList(deviceInfoSet DevInfo, deviceInfoData *DevInfoData, driverType SPDIT) (err error) {
	r1, _, e1 := syscall.Syscall(procSetupDiBuildDriverInfoList.Addr(), 3, uintptr(deviceInfoSet), uintptr(unsafe.Pointer(deviceInfoData)), uintptr(driverType))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  初始化类安装程序函数(installFunction  安装功能,  deviceInfoSet  设备信息集,  deviceInfoData  *设备信息数据)  (error  错误)  {}

// ff:
// installFunction:
// deviceInfoSet:
// deviceInfoData:
// err:
func SetupDiCallClassInstaller(installFunction DI_FUNCTION, deviceInfoSet DevInfo, deviceInfoData *DevInfoData) (err error) {
	r1, _, e1 := syscall.Syscall(procSetupDiCallClassInstaller.Addr(), 3, uintptr(installFunction), uintptr(deviceInfoSet), uintptr(unsafe.Pointer(deviceInfoData)))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  初始化取消驱动程序信息搜索(deviceInfoSet  设备信息集)  (错误  错误)  {}

// ff:
// deviceInfoSet:
// err:
func SetupDiCancelDriverInfoSearch(deviceInfoSet DevInfo) (err error) {
	r1, _, e1 := syscall.Syscall(procSetupDiCancelDriverInfoSearch.Addr(), 1, uintptr(deviceInfoSet), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func setupDiClassGuidsFromNameEx(className *uint16, classGuidList *GUID, classGuidListSize uint32, requiredSize *uint32, machineName *uint16, reserved uintptr) (err error) {
	r1, _, e1 := syscall.Syscall6(procSetupDiClassGuidsFromNameExW.Addr(), 6, uintptr(unsafe.Pointer(className)), uintptr(unsafe.Pointer(classGuidList)), uintptr(classGuidListSize), uintptr(unsafe.Pointer(requiredSize)), uintptr(unsafe.Pointer(machineName)), uintptr(reserved))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func setupDiClassNameFromGuidEx(classGUID *GUID, className *uint16, classNameSize uint32, requiredSize *uint32, machineName *uint16, reserved uintptr) (err error) {
	r1, _, e1 := syscall.Syscall6(procSetupDiClassNameFromGuidExW.Addr(), 6, uintptr(unsafe.Pointer(classGUID)), uintptr(unsafe.Pointer(className)), uintptr(classNameSize), uintptr(unsafe.Pointer(requiredSize)), uintptr(unsafe.Pointer(machineName)), uintptr(reserved))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func setupDiCreateDeviceInfoListEx(classGUID *GUID, hwndParent uintptr, machineName *uint16, reserved uintptr) (handle DevInfo, err error) {
	r0, _, e1 := syscall.Syscall6(procSetupDiCreateDeviceInfoListExW.Addr(), 4, uintptr(unsafe.Pointer(classGUID)), uintptr(hwndParent), uintptr(unsafe.Pointer(machineName)), uintptr(reserved), 0, 0)
	handle = DevInfo(r0)
	if handle == DevInfo(InvalidHandle) {
		err = errnoErr(e1)
	}
	return
}

func setupDiCreateDeviceInfo(deviceInfoSet DevInfo, DeviceName *uint16, classGUID *GUID, DeviceDescription *uint16, hwndParent uintptr, CreationFlags DICD, deviceInfoData *DevInfoData) (err error) {
	r1, _, e1 := syscall.Syscall9(procSetupDiCreateDeviceInfoW.Addr(), 7, uintptr(deviceInfoSet), uintptr(unsafe.Pointer(DeviceName)), uintptr(unsafe.Pointer(classGUID)), uintptr(unsafe.Pointer(DeviceDescription)), uintptr(hwndParent), uintptr(CreationFlags), uintptr(unsafe.Pointer(deviceInfoData)), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  销毁设备信息列表(deviceInfoSet  设备信息集)  (错误  error)  {}

// ff:
// deviceInfoSet:
// err:
func SetupDiDestroyDeviceInfoList(deviceInfoSet DevInfo) (err error) {
	r1, _, e1 := syscall.Syscall(procSetupDiDestroyDeviceInfoList.Addr(), 1, uintptr(deviceInfoSet), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  销毁驱动信息列表(deviceInfoSet  设备信息集,  deviceInfoData  设备信息数据指针,  driverType  驱动类型)  (错误  error)  {}

// ff:
// deviceInfoSet:
// deviceInfoData:
// driverType:
// err:
func SetupDiDestroyDriverInfoList(deviceInfoSet DevInfo, deviceInfoData *DevInfoData, driverType SPDIT) (err error) {
	r1, _, e1 := syscall.Syscall(procSetupDiDestroyDriverInfoList.Addr(), 3, uintptr(deviceInfoSet), uintptr(unsafe.Pointer(deviceInfoData)), uintptr(driverType))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func setupDiEnumDeviceInfo(deviceInfoSet DevInfo, memberIndex uint32, deviceInfoData *DevInfoData) (err error) {
	r1, _, e1 := syscall.Syscall(procSetupDiEnumDeviceInfo.Addr(), 3, uintptr(deviceInfoSet), uintptr(memberIndex), uintptr(unsafe.Pointer(deviceInfoData)))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func setupDiEnumDriverInfo(deviceInfoSet DevInfo, deviceInfoData *DevInfoData, driverType SPDIT, memberIndex uint32, driverInfoData *DrvInfoData) (err error) {
	r1, _, e1 := syscall.Syscall6(procSetupDiEnumDriverInfoW.Addr(), 5, uintptr(deviceInfoSet), uintptr(unsafe.Pointer(deviceInfoData)), uintptr(driverType), uintptr(memberIndex), uintptr(unsafe.Pointer(driverInfoData)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func setupDiGetClassDevsEx(classGUID *GUID, Enumerator *uint16, hwndParent uintptr, Flags DIGCF, deviceInfoSet DevInfo, machineName *uint16, reserved uintptr) (handle DevInfo, err error) {
	r0, _, e1 := syscall.Syscall9(procSetupDiGetClassDevsExW.Addr(), 7, uintptr(unsafe.Pointer(classGUID)), uintptr(unsafe.Pointer(Enumerator)), uintptr(hwndParent), uintptr(Flags), uintptr(deviceInfoSet), uintptr(unsafe.Pointer(machineName)), uintptr(reserved), 0, 0)
	handle = DevInfo(r0)
	if handle == DevInfo(InvalidHandle) {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  设置Di获取类安装参数(deviceInfoSet  设备信息集,  deviceInfoData  *设备信息数据,  classInstallParams  *类安装头部,  classInstallParamsSize  uint32,  requiredSize  *uint32)  (错误  error)  {}

// ff:
// deviceInfoSet:
// deviceInfoData:
// classInstallParams:
// classInstallParamsSize:
// requiredSize:
// err:
func SetupDiGetClassInstallParams(deviceInfoSet DevInfo, deviceInfoData *DevInfoData, classInstallParams *ClassInstallHeader, classInstallParamsSize uint32, requiredSize *uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procSetupDiGetClassInstallParamsW.Addr(), 5, uintptr(deviceInfoSet), uintptr(unsafe.Pointer(deviceInfoData)), uintptr(unsafe.Pointer(classInstallParams)), uintptr(classInstallParamsSize), uintptr(unsafe.Pointer(requiredSize)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func setupDiGetDeviceInfoListDetail(deviceInfoSet DevInfo, deviceInfoSetDetailData *DevInfoListDetailData) (err error) {
	r1, _, e1 := syscall.Syscall(procSetupDiGetDeviceInfoListDetailW.Addr(), 2, uintptr(deviceInfoSet), uintptr(unsafe.Pointer(deviceInfoSetDetailData)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func setupDiGetDeviceInstallParams(deviceInfoSet DevInfo, deviceInfoData *DevInfoData, deviceInstallParams *DevInstallParams) (err error) {
	r1, _, e1 := syscall.Syscall(procSetupDiGetDeviceInstallParamsW.Addr(), 3, uintptr(deviceInfoSet), uintptr(unsafe.Pointer(deviceInfoData)), uintptr(unsafe.Pointer(deviceInstallParams)))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func setupDiGetDeviceInstanceId(deviceInfoSet DevInfo, deviceInfoData *DevInfoData, instanceId *uint16, instanceIdSize uint32, instanceIdRequiredSize *uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procSetupDiGetDeviceInstanceIdW.Addr(), 5, uintptr(deviceInfoSet), uintptr(unsafe.Pointer(deviceInfoData)), uintptr(unsafe.Pointer(instanceId)), uintptr(instanceIdSize), uintptr(unsafe.Pointer(instanceIdRequiredSize)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func setupDiGetDeviceProperty(deviceInfoSet DevInfo, deviceInfoData *DevInfoData, propertyKey *DEVPROPKEY, propertyType *DEVPROPTYPE, propertyBuffer *byte, propertyBufferSize uint32, requiredSize *uint32, flags uint32) (err error) {
	r1, _, e1 := syscall.Syscall9(procSetupDiGetDevicePropertyW.Addr(), 8, uintptr(deviceInfoSet), uintptr(unsafe.Pointer(deviceInfoData)), uintptr(unsafe.Pointer(propertyKey)), uintptr(unsafe.Pointer(propertyType)), uintptr(unsafe.Pointer(propertyBuffer)), uintptr(propertyBufferSize), uintptr(unsafe.Pointer(requiredSize)), uintptr(flags), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func setupDiGetDeviceRegistryProperty(deviceInfoSet DevInfo, deviceInfoData *DevInfoData, property SPDRP, propertyRegDataType *uint32, propertyBuffer *byte, propertyBufferSize uint32, requiredSize *uint32) (err error) {
	r1, _, e1 := syscall.Syscall9(procSetupDiGetDeviceRegistryPropertyW.Addr(), 7, uintptr(deviceInfoSet), uintptr(unsafe.Pointer(deviceInfoData)), uintptr(property), uintptr(unsafe.Pointer(propertyRegDataType)), uintptr(unsafe.Pointer(propertyBuffer)), uintptr(propertyBufferSize), uintptr(unsafe.Pointer(requiredSize)), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func setupDiGetDriverInfoDetail(deviceInfoSet DevInfo, deviceInfoData *DevInfoData, driverInfoData *DrvInfoData, driverInfoDetailData *DrvInfoDetailData, driverInfoDetailDataSize uint32, requiredSize *uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procSetupDiGetDriverInfoDetailW.Addr(), 6, uintptr(deviceInfoSet), uintptr(unsafe.Pointer(deviceInfoData)), uintptr(unsafe.Pointer(driverInfoData)), uintptr(unsafe.Pointer(driverInfoDetailData)), uintptr(driverInfoDetailDataSize), uintptr(unsafe.Pointer(requiredSize)))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func setupDiGetSelectedDevice(deviceInfoSet DevInfo, deviceInfoData *DevInfoData) (err error) {
	r1, _, e1 := syscall.Syscall(procSetupDiGetSelectedDevice.Addr(), 2, uintptr(deviceInfoSet), uintptr(unsafe.Pointer(deviceInfoData)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func setupDiGetSelectedDriver(deviceInfoSet DevInfo, deviceInfoData *DevInfoData, driverInfoData *DrvInfoData) (err error) {
	r1, _, e1 := syscall.Syscall(procSetupDiGetSelectedDriverW.Addr(), 3, uintptr(deviceInfoSet), uintptr(unsafe.Pointer(deviceInfoData)), uintptr(unsafe.Pointer(driverInfoData)))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  初始化设备注册键(deviceInfoSet  设备信息集,  deviceInfoData  *设备信息数据,  范围  设备配置标志,  硬件配置文件  uint32,  键类型  设备注册类型,  访问权限  uint32)  (键句柄  手柄,  错误  error)  {}

// ff:
// deviceInfoSet:
// deviceInfoData:
// Scope:
// HwProfile:
// KeyType:
// samDesired:
// key:
// err:
func SetupDiOpenDevRegKey(deviceInfoSet DevInfo, deviceInfoData *DevInfoData, Scope DICS_FLAG, HwProfile uint32, KeyType DIREG, samDesired uint32) (key Handle, err error) {
	r0, _, e1 := syscall.Syscall6(procSetupDiOpenDevRegKey.Addr(), 6, uintptr(deviceInfoSet), uintptr(unsafe.Pointer(deviceInfoData)), uintptr(Scope), uintptr(HwProfile), uintptr(KeyType), uintptr(samDesired))
	key = Handle(r0)
	if key == InvalidHandle {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  设置类安装参数(deviceInfoSet  设备信息集,  deviceInfoData  设备信息数据指针,  classInstallParams  类安装头指针,  classInstallParamsSize  uint32)  (错误  error)  {}

// ff:
// deviceInfoSet:
// deviceInfoData:
// classInstallParams:
// classInstallParamsSize:
// err:
func SetupDiSetClassInstallParams(deviceInfoSet DevInfo, deviceInfoData *DevInfoData, classInstallParams *ClassInstallHeader, classInstallParamsSize uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procSetupDiSetClassInstallParamsW.Addr(), 4, uintptr(deviceInfoSet), uintptr(unsafe.Pointer(deviceInfoData)), uintptr(unsafe.Pointer(classInstallParams)), uintptr(classInstallParamsSize), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  设置设备安装参数(deviceInfoSet  设备信息集,  deviceInfoData  设备信息数据指针,  deviceInstallParams  设备安装参数指针)  (错误  error)  {}

// ff:
// deviceInfoSet:
// deviceInfoData:
// deviceInstallParams:
// err:
func SetupDiSetDeviceInstallParams(deviceInfoSet DevInfo, deviceInfoData *DevInfoData, deviceInstallParams *DevInstallParams) (err error) {
	r1, _, e1 := syscall.Syscall(procSetupDiSetDeviceInstallParamsW.Addr(), 3, uintptr(deviceInfoSet), uintptr(unsafe.Pointer(deviceInfoData)), uintptr(unsafe.Pointer(deviceInstallParams)))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func setupDiSetDeviceRegistryProperty(deviceInfoSet DevInfo, deviceInfoData *DevInfoData, property SPDRP, propertyBuffer *byte, propertyBufferSize uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procSetupDiSetDeviceRegistryPropertyW.Addr(), 5, uintptr(deviceInfoSet), uintptr(unsafe.Pointer(deviceInfoData)), uintptr(property), uintptr(unsafe.Pointer(propertyBuffer)), uintptr(propertyBufferSize), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  设置选定设备(deviceInfoSet  设备信息集,  deviceInfoData  设备信息数据指针)  (错误  error)  {}

// ff:
// deviceInfoSet:
// deviceInfoData:
// err:
func SetupDiSetSelectedDevice(deviceInfoSet DevInfo, deviceInfoData *DevInfoData) (err error) {
	r1, _, e1 := syscall.Syscall(procSetupDiSetSelectedDevice.Addr(), 2, uintptr(deviceInfoSet), uintptr(unsafe.Pointer(deviceInfoData)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  设置选定驱动程序(deviceInfoSet  设备信息集,  deviceInfoData  设备信息数据指针,  driverInfoData  驱动信息数据指针)  (error)  {}

// ff:
// deviceInfoSet:
// deviceInfoData:
// driverInfoData:
// err:
func SetupDiSetSelectedDriver(deviceInfoSet DevInfo, deviceInfoData *DevInfoData, driverInfoData *DrvInfoData) (err error) {
	r1, _, e1 := syscall.Syscall(procSetupDiSetSelectedDriverW.Addr(), 3, uintptr(deviceInfoSet), uintptr(unsafe.Pointer(deviceInfoData)), uintptr(unsafe.Pointer(driverInfoData)))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func setupUninstallOEMInf(infFileName *uint16, flags SUOI, reserved uintptr) (err error) {
	r1, _, e1 := syscall.Syscall(procSetupUninstallOEMInfW.Addr(), 3, uintptr(unsafe.Pointer(infFileName)), uintptr(flags), uintptr(reserved))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func commandLineToArgv(cmd *uint16, argc *int32) (argv **uint16, err error) {
	r0, _, e1 := syscall.Syscall(procCommandLineToArgvW.Addr(), 2, uintptr(unsafe.Pointer(cmd)), uintptr(unsafe.Pointer(argc)), 0)
	argv = (**uint16)(unsafe.Pointer(r0))
	if argv == nil {
		err = errnoErr(e1)
	}
	return
}

func shGetKnownFolderPath(id *KNOWNFOLDERID, flags uint32, token Token, path **uint16) (ret error) {
	r0, _, _ := syscall.Syscall6(procSHGetKnownFolderPath.Addr(), 4, uintptr(unsafe.Pointer(id)), uintptr(flags), uintptr(token), uintptr(unsafe.Pointer(path)), 0, 0)
	if r0 != 0 {
		ret = syscall.Errno(r0)
	}
	return
}

// 翻译提示:func  执行壳层(hwnd  控制句柄,  操作字符串  *uint16,  文件路径  *uint16,  参数  *uint16,  当前工作目录  *uint16,  显示模式  int32)  (错误  error)  {}

// ff:
// hwnd:
// verb:
// file:
// args:
// cwd:
// showCmd:
// err:
func ShellExecute(hwnd Handle, verb *uint16, file *uint16, args *uint16, cwd *uint16, showCmd int32) (err error) {
	r1, _, e1 := syscall.Syscall6(procShellExecuteW.Addr(), 6, uintptr(hwnd), uintptr(unsafe.Pointer(verb)), uintptr(unsafe.Pointer(file)), uintptr(unsafe.Pointer(args)), uintptr(unsafe.Pointer(cwd)), uintptr(showCmd))
	if r1 <= 32 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  列举子窗口(hwnd  窗口句柄,  枚举函数  uintptr,  参数  unsafe.Pointer)  {}

// ff:
// hwnd:
// enumFunc:
// param:
func EnumChildWindows(hwnd HWND, enumFunc uintptr, param unsafe.Pointer) {
	syscall.Syscall(procEnumChildWindows.Addr(), 3, uintptr(hwnd), uintptr(enumFunc), uintptr(param))
	return
}

// 翻译提示:func  列举窗口(enumFunc  uintptr,  参数  unsafe.Pointer)  (错误  error)  {}

// ff:
// enumFunc:
// param:
// err:
func EnumWindows(enumFunc uintptr, param unsafe.Pointer) (err error) {
	r1, _, e1 := syscall.Syscall(procEnumWindows.Addr(), 2, uintptr(enumFunc), uintptr(param), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  关闭WindowsEx(标志  uint32,  原因  uint32)  (错误  error)  {}

// ff:
// flags:
// reason:
// err:
func ExitWindowsEx(flags uint32, reason uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procExitWindowsEx.Addr(), 2, uintptr(flags), uintptr(reason), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  获取窗口类名(hwnd  窗口句柄,  类名  *uint16,  最大长度  int32)  (复制长度  int32,  错误  error)  {}

// ff:
// hwnd:
// className:
// maxCount:
// copied:
// err:
func GetClassName(hwnd HWND, className *uint16, maxCount int32) (copied int32, err error) {
	r0, _, e1 := syscall.Syscall(procGetClassNameW.Addr(), 3, uintptr(hwnd), uintptr(unsafe.Pointer(className)), uintptr(maxCount))
	copied = int32(r0)
	if copied == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  获取桌面窗口()  (桌面窗口句柄  HWND)  {}

// ff:
// hwnd:
func GetDesktopWindow() (hwnd HWND) {
	r0, _, _ := syscall.Syscall(procGetDesktopWindow.Addr(), 0, 0, 0, 0)
	hwnd = HWND(r0)
	return
}

// 翻译提示:func  获取活跃窗口()  (窗口句柄  HWND)  {}

// ff:
// hwnd:
func GetForegroundWindow() (hwnd HWND) {
	r0, _, _ := syscall.Syscall(procGetForegroundWindow.Addr(), 0, 0, 0, 0)
	hwnd = HWND(r0)
	return
}

// 翻译提示:func  获取GUI线程信息(thread  uint32,  info  *GUI线程信息)  (错误  error)  {}

// ff:
// thread:
// info:
// err:
func GetGUIThreadInfo(thread uint32, info *GUIThreadInfo) (err error) {
	r1, _, e1 := syscall.Syscall(procGetGUIThreadInfo.Addr(), 2, uintptr(thread), uintptr(unsafe.Pointer(info)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  获取Shell窗口()  (shell窗口  HWND)  {}

// ff:
// shellWindow:
func GetShellWindow() (shellWindow HWND) {
	r0, _, _ := syscall.Syscall(procGetShellWindow.Addr(), 0, 0, 0, 0)
	shellWindow = HWND(r0)
	return
}

// 翻译提示:func  获取窗口线程进程ID(hwnd  窗口句柄,  pid  进程ID指针)  (线程ID  uint32,  错误  error)  {}

// ff:
// hwnd:
// pid:
// tid:
// err:
func GetWindowThreadProcessId(hwnd HWND, pid *uint32) (tid uint32, err error) {
	r0, _, e1 := syscall.Syscall(procGetWindowThreadProcessId.Addr(), 2, uintptr(hwnd), uintptr(unsafe.Pointer(pid)), 0)
	tid = uint32(r0)
	if tid == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  是否是窗口(hwnd  窗口句柄)  (是窗口  bool)  {}

// ff:
// hwnd:
// isWindow:
func IsWindow(hwnd HWND) (isWindow bool) {
	r0, _, _ := syscall.Syscall(procIsWindow.Addr(), 1, uintptr(hwnd), 0, 0)
	isWindow = r0 != 0
	return
}

// 翻译提示:func  是否为Unicode窗口(hwnd  轴心点句柄)  (是Unicode  bool)  {}

// ff:
// hwnd:
// isUnicode:
func IsWindowUnicode(hwnd HWND) (isUnicode bool) {
	r0, _, _ := syscall.Syscall(procIsWindowUnicode.Addr(), 1, uintptr(hwnd), 0, 0)
	isUnicode = r0 != 0
	return
}

// 翻译提示:func  窗口是否可见(hwnd  窗口句柄)  (可见性  是否可见)  {}

// ff:
// hwnd:
// isVisible:
func IsWindowVisible(hwnd HWND) (isVisible bool) {
	r0, _, _ := syscall.Syscall(procIsWindowVisible.Addr(), 1, uintptr(hwnd), 0, 0)
	isVisible = r0 != 0
	return
}

// 翻译提示:func  弹出对话框(hwnd  窗口句柄,  文本  *uint16,  标题  *uint16,  对话框类型  uint32)  (返回值  int32,  错误  error)  {}

// ff:
// hwnd:
// text:
// caption:
// boxtype:
// ret:
// err:
func MessageBox(hwnd HWND, text *uint16, caption *uint16, boxtype uint32) (ret int32, err error) {
	r0, _, e1 := syscall.Syscall6(procMessageBoxW.Addr(), 4, uintptr(hwnd), uintptr(unsafe.Pointer(text)), uintptr(unsafe.Pointer(caption)), uintptr(boxtype), 0, 0)
	ret = int32(r0)
	if ret == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  创建环境块(block  **字符串,  token  令牌,  继承现有  bool)  (错误  error)  {}

// ff:
// block:
// token:
// inheritExisting:
// err:
func CreateEnvironmentBlock(block **uint16, token Token, inheritExisting bool) (err error) {
	var _p0 uint32
	if inheritExisting {
		_p0 = 1
	}
	r1, _, e1 := syscall.Syscall(procCreateEnvironmentBlock.Addr(), 3, uintptr(unsafe.Pointer(block)), uintptr(token), uintptr(_p0))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  销毁环境块(block  *字节序)  (错误  error)  {}

// ff:
// block:
// err:
func DestroyEnvironmentBlock(block *uint16) (err error) {
	r1, _, e1 := syscall.Syscall(procDestroyEnvironmentBlock.Addr(), 1, uintptr(unsafe.Pointer(block)), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  获取用户配置目录(token  Token,  目录  *uint16,  目录长度  *uint32)  (错误  error)  {}

// ff:
// t:
// dir:
// dirLen:
// err:
func GetUserProfileDirectory(t Token, dir *uint16, dirLen *uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procGetUserProfileDirectoryW.Addr(), 3, uintptr(t), uintptr(unsafe.Pointer(dir)), uintptr(unsafe.Pointer(dirLen)))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  获取文件版本信息大小(filename  string,  zeroHandle  *处理器句柄)  (bufSize  uint32,  err  错误)  {}

// ff:
// filename:
// zeroHandle:
// bufSize:
// err:
func GetFileVersionInfoSize(filename string, zeroHandle *Handle) (bufSize uint32, err error) {
	var _p0 *uint16
	_p0, err = syscall.UTF16PtrFromString(filename)
	if err != nil {
		return
	}
	return _GetFileVersionInfoSize(_p0, zeroHandle)
}

func _GetFileVersionInfoSize(filename *uint16, zeroHandle *Handle) (bufSize uint32, err error) {
	r0, _, e1 := syscall.Syscall(procGetFileVersionInfoSizeW.Addr(), 2, uintptr(unsafe.Pointer(filename)), uintptr(unsafe.Pointer(zeroHandle)), 0)
	bufSize = uint32(r0)
	if bufSize == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  获取文件版本信息(filename  string,  handle  uint32,  bufSize  uint32,  buffer  unsafe.Pointer)  (错误  error)  {}

// ff:
// filename:
// handle:
// bufSize:
// buffer:
// err:
func GetFileVersionInfo(filename string, handle uint32, bufSize uint32, buffer unsafe.Pointer) (err error) {
	var _p0 *uint16
	_p0, err = syscall.UTF16PtrFromString(filename)
	if err != nil {
		return
	}
	return _GetFileVersionInfo(_p0, handle, bufSize, buffer)
}

func _GetFileVersionInfo(filename *uint16, handle uint32, bufSize uint32, buffer unsafe.Pointer) (err error) {
	r1, _, e1 := syscall.Syscall6(procGetFileVersionInfoW.Addr(), 4, uintptr(unsafe.Pointer(filename)), uintptr(handle), uintptr(bufSize), uintptr(buffer), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  查询版本信息块(block  unsafe.Pointer,  子块字符串  string,  指针到缓冲区指针  unsafe.Pointer,  缓冲区大小  *uint32)  (错误  error)  {}

// ff:
// block:
// subBlock:
// pointerToBufferPointer:
// bufSize:
// err:
func VerQueryValue(block unsafe.Pointer, subBlock string, pointerToBufferPointer unsafe.Pointer, bufSize *uint32) (err error) {
	var _p0 *uint16
	_p0, err = syscall.UTF16PtrFromString(subBlock)
	if err != nil {
		return
	}
	return _VerQueryValue(block, _p0, pointerToBufferPointer, bufSize)
}

func _VerQueryValue(block unsafe.Pointer, subBlock *uint16, pointerToBufferPointer unsafe.Pointer, bufSize *uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procVerQueryValueW.Addr(), 4, uintptr(block), uintptr(unsafe.Pointer(subBlock)), uintptr(pointerToBufferPointer), uintptr(unsafe.Pointer(bufSize)), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  开始时间周期(period  uint32)  (错误  error)  {}

// ff:
// period:
// err:
func TimeBeginPeriod(period uint32) (err error) {
	r1, _, e1 := syscall.Syscall(proctimeBeginPeriod.Addr(), 1, uintptr(period), 0, 0)
	if r1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  结束时间周期(period  uint32)  (error  err)  {}

// ff:
// period:
// err:
func TimeEndPeriod(period uint32) (err error) {
	r1, _, e1 := syscall.Syscall(proctimeEndPeriod.Addr(), 1, uintptr(period), 0, 0)
	if r1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  确认Windows信任策略(hwnd  窗口句柄,  actionId  操作ID指针,  数据  *Windows信任数据)  (错误  error)  {}

// ff:
// hwnd:
// actionId:
// data:
// ret:
func WinVerifyTrustEx(hwnd HWND, actionId *GUID, data *WinTrustData) (ret error) {
	r0, _, _ := syscall.Syscall(procWinVerifyTrustEx.Addr(), 3, uintptr(hwnd), uintptr(unsafe.Pointer(actionId)), uintptr(unsafe.Pointer(data)))
	if r0 != 0 {
		ret = syscall.Errno(r0)
	}
	return
}

// 翻译提示:func  释放AddrinfoW资源(addrinfo  *AddrinfoW)  {}

// ff:
// addrinfo:
func FreeAddrInfoW(addrinfo *AddrinfoW) {
	syscall.Syscall(procFreeAddrInfoW.Addr(), 1, uintptr(unsafe.Pointer(addrinfo)), 0, 0)
	return
}

// 翻译提示:func  获取地址信息W(nodeName  *字符串指针,  服务名  *字符串指针,  提示  *AddrinfoW,  结果  **AddrinfoW)  (套接字错误  error)  {}

// ff:
// nodename:
// servicename:
// hints:
// result:
// sockerr:
func GetAddrInfoW(nodename *uint16, servicename *uint16, hints *AddrinfoW, result **AddrinfoW) (sockerr error) {
	r0, _, _ := syscall.Syscall6(procGetAddrInfoW.Addr(), 4, uintptr(unsafe.Pointer(nodename)), uintptr(unsafe.Pointer(servicename)), uintptr(unsafe.Pointer(hints)), uintptr(unsafe.Pointer(result)), 0, 0)
	if r0 != 0 {
		sockerr = syscall.Errno(r0)
	}
	return
}

// 翻译提示:func  WinsockCleanup()  (错误  error)  {}

// ff:
// err:
func WSACleanup() (err error) {
	r1, _, e1 := syscall.Syscall(procWSACleanup.Addr(), 0, 0, 0, 0)
	if r1 == socket_error {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  WSAGetProtocols(协议ID  *int32,  协议信息Buffer  *WSAProtocolInfo,  缓冲区长度  *uint32)  (成功数量  int32,  错误  error)  {}

// ff:
// protocols:
// protocolBuffer:
// bufferLength:
// n:
// err:
func WSAEnumProtocols(protocols *int32, protocolBuffer *WSAProtocolInfo, bufferLength *uint32) (n int32, err error) {
	r0, _, e1 := syscall.Syscall(procWSAEnumProtocolsW.Addr(), 3, uintptr(unsafe.Pointer(protocols)), uintptr(unsafe.Pointer(protocolBuffer)), uintptr(unsafe.Pointer(bufferLength)))
	n = int32(r0)
	if n == -1 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  获取重叠结果(h  文件句柄,  o  *重叠结构,  字节传输  *uint32,  是否等待  bool,  标志  *uint32)  (错误  error)  {}

// ff:
// h:
// o:
// bytes:
// wait:
// flags:
// err:
func WSAGetOverlappedResult(h Handle, o *Overlapped, bytes *uint32, wait bool, flags *uint32) (err error) {
	var _p0 uint32
	if wait {
		_p0 = 1
	}
	r1, _, e1 := syscall.Syscall6(procWSAGetOverlappedResult.Addr(), 5, uintptr(h), uintptr(unsafe.Pointer(o)), uintptr(unsafe.Pointer(bytes)), uintptr(_p0), uintptr(unsafe.Pointer(flags)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  WSAGetLastError()  (错误  error)  {}
// 
// func  WSAStartup(wVersionRequested  uint16,  wsadata  *WSAData)  (err  error)  {}
// 
// func  WSACleanup()  (err  error)  {}
// 
// func  Socket(domain  Domain,  typ  Type,  protocol  int)  (s  Handle,  err  error)  {}
// 
// func  Connect(s  Handle,  addr  SOCKADDR,  addrlen  _Socklen)  (err  error)  {}
// 
// func  Bind(s  Handle,  addr  SOCKADDR,  addrlen  _Socklen)  (err  error)  {}
// 
// func  Listen(s  Handle,  backlog  int)  (err  error)  {}
// 
// func  Accept(s  Handle,  rsa  *SOCKADDR,  addrlen  *_Socklen)  (newS  Handle,  err  error)  {}
// 
// func  Send(s  Handle,  buf  []byte,  flags  int)  (sent  int,  err  error)  {}
// 
// func  Recv(s  Handle,  buf  []byte,  flags  int)  (received  int,  err  error)  {}
// 
// func  WSAIoctl(s  Handle,  控制码  uint32,  输入缓冲区  *byte,  输入缓冲区大小  uint32,  输出缓冲区  *byte,  输出缓冲区大小  uint32,  实际读取字节数  *uint32,  超时处理  *Overlapped,  完成例程  uintptr)  (err  error)  {}

// ff:
// s:
// iocc:
// inbuf:
// cbif:
// outbuf:
// cbob:
// cbbr:
// overlapped:
// completionRoutine:
// err:
func WSAIoctl(s Handle, iocc uint32, inbuf *byte, cbif uint32, outbuf *byte, cbob uint32, cbbr *uint32, overlapped *Overlapped, completionRoutine uintptr) (err error) {
	r1, _, e1 := syscall.Syscall9(procWSAIoctl.Addr(), 9, uintptr(s), uintptr(iocc), uintptr(unsafe.Pointer(inbuf)), uintptr(cbif), uintptr(unsafe.Pointer(outbuf)), uintptr(cbob), uintptr(unsafe.Pointer(cbbr)), uintptr(unsafe.Pointer(overlapped)), uintptr(completionRoutine))
	if r1 == socket_error {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  WSA开始查询服务(querySet  *查询集,  标志  uint32,  处理句柄  *句柄)  (错误  error)  {}
// 
// func  WSALookupServiceNext(handle  Handle,  flags  uint32,  bufferLen  *uint32,  querySet  **WSAQUERYSET)  (err  error)  {}
// 
// func  WSACancelLookupService(handle  Handle)  (err  error)  {}
// 
// func  WSASetService(stateSet  *WSAESRVCMD,  querySet  *WSAQUERYSET,  flags  uint32)  (err  error)  {}
// 
// func  WSAGetServiceClassInfo(classId  *GUID,  providerId  *GUID,  buffer  *byte,  bufferLen  *uint32)  (err  error)  {}
// 
// func  WSAGetServiceClassNameById(classId  *GUID,  name  **uint16,  languageId  *uint32)  (err  error)  {}
// 
// func  WSAGetProviderInfo(providerId  *GUID,  buffer  *byte,  bufferLen  *uint32)  (err  error)  {}
// 
// func  WSAStringToAddress(addressString  *uint16,  addressFamily  int32,  socketType  int32,  protocol  int32,  address  *_sockaddr,  addressLen  *int32)  (err  error)  {}
// 
// func  WSAApiVersion()  (majorVersion  uint16,  minorVersion  uint16)  {}
// 
// func  WSAProtocolInfoCount()  (count  uint32)  {}
// 
// func  WSAGetProtocolInfo(protocolId  *GUID,  info  **_WSAPROTOCOL_INFO)  (err  error)  {}
// 
// func  WSAEnumNameSpaceProviders(bufferLen  uint32,  providers  **_WSANAMESPACE_INFO)  (bytesNeeded  uint32,  err  error)  {}
// 
// func  WSACreateNameSpace(providerId  *GUID,  namespaceName  *uint16,  version  *WSANSAPROVIDERID,  description  *uint16,  flags  uint32,  providerContext  *any)  (err  error)  {}
// 
// func  WSADeleteNameSpace(providerId  *GUID)  (err  error)  {}
// 
// func  WSAGetServiceProviderInfo(providerId  *GUID,  buffer  *byte,  bufferLen  *uint32)  (err  error)  {}

// ff:
// querySet:
// flags:
// handle:
// err:
func WSALookupServiceBegin(querySet *WSAQUERYSET, flags uint32, handle *Handle) (err error) {
	r1, _, e1 := syscall.Syscall(procWSALookupServiceBeginW.Addr(), 3, uintptr(unsafe.Pointer(querySet)), uintptr(flags), uintptr(unsafe.Pointer(handle)))
	if r1 == socket_error {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  停止WSA服务查询(handle  连接句柄)  (错误  error)  {}

// ff:
// handle:
// err:
func WSALookupServiceEnd(handle Handle) (err error) {
	r1, _, e1 := syscall.Syscall(procWSALookupServiceEnd.Addr(), 1, uintptr(handle), 0, 0)
	if r1 == socket_error {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  WSA查询服务下一步句柄(handle  手柄,  标志  uint32,  大小  *int32,  查询集  *WSAQUERYSET)  (错误  error)  {}

// ff:
// handle:
// flags:
// size:
// querySet:
// err:
func WSALookupServiceNext(handle Handle, flags uint32, size *int32, querySet *WSAQUERYSET) (err error) {
	r1, _, e1 := syscall.Syscall6(procWSALookupServiceNextW.Addr(), 4, uintptr(handle), uintptr(flags), uintptr(unsafe.Pointer(size)), uintptr(unsafe.Pointer(querySet)), 0, 0)
	if r1 == socket_error {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  接收WSA(s  手柄,  bufs  *缓冲区,  缓冲区计数  uint32,  已接收  *uint32,  标志  *uint32,  重叠  *重叠,  协程  *字节)  (错误  error)  {}

// ff:
// s:
// bufs:
// bufcnt:
// recvd:
// flags:
// overlapped:
// croutine:
// err:
func WSARecv(s Handle, bufs *WSABuf, bufcnt uint32, recvd *uint32, flags *uint32, overlapped *Overlapped, croutine *byte) (err error) {
	r1, _, e1 := syscall.Syscall9(procWSARecv.Addr(), 7, uintptr(s), uintptr(unsafe.Pointer(bufs)), uintptr(bufcnt), uintptr(unsafe.Pointer(recvd)), uintptr(unsafe.Pointer(flags)), uintptr(unsafe.Pointer(overlapped)), uintptr(unsafe.Pointer(croutine)), 0, 0)
	if r1 == socket_error {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  接收来自WSA(s  手柄,  缓冲区  *WSABuf,  缓冲区计数  uint32,  已接收  *uint32,  标志  *uint32,  来源地址  *RawSockaddrAny,  地址长度  *int32,  重叠  *Overlapped,  协程  *byte)  (错误  error)  {}

// ff:
// s:
// bufs:
// bufcnt:
// recvd:
// flags:
// from:
// fromlen:
// overlapped:
// croutine:
// err:
func WSARecvFrom(s Handle, bufs *WSABuf, bufcnt uint32, recvd *uint32, flags *uint32, from *RawSockaddrAny, fromlen *int32, overlapped *Overlapped, croutine *byte) (err error) {
	r1, _, e1 := syscall.Syscall9(procWSARecvFrom.Addr(), 9, uintptr(s), uintptr(unsafe.Pointer(bufs)), uintptr(bufcnt), uintptr(unsafe.Pointer(recvd)), uintptr(unsafe.Pointer(flags)), uintptr(unsafe.Pointer(from)), uintptr(unsafe.Pointer(fromlen)), uintptr(unsafe.Pointer(overlapped)), uintptr(unsafe.Pointer(croutine)))
	if r1 == socket_error {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  WinsockSend(socket  手柄,  buffers  *WinsockBuffer,  bufferCount  uint32,  sentBytes  *uint32,  flags  uint32,  overlapped  *重叠结构,  completionRoutine  *字节)  (错误  error)  {}  
// 
// 其中：
// -  Handle  ->  手柄：代表网络连接的句柄
// -  WSABuf  ->  WinsockBuffer：Windows  Socket缓冲区结构体
// -  bufcnt  ->  bufferCount：缓冲区数量
// -  sent  ->  sentBytes：已发送字节数指针
// -  flags  ->  标志：用于控制发送操作的标志
// -  Overlapped  ->  重叠结构：用于异步操作的数据结构
// -  byte  ->  字节：可能表示完成例程的字节指针

// ff:
// s:
// bufs:
// bufcnt:
// sent:
// flags:
// overlapped:
// croutine:
// err:
func WSASend(s Handle, bufs *WSABuf, bufcnt uint32, sent *uint32, flags uint32, overlapped *Overlapped, croutine *byte) (err error) {
	r1, _, e1 := syscall.Syscall9(procWSASend.Addr(), 7, uintptr(s), uintptr(unsafe.Pointer(bufs)), uintptr(bufcnt), uintptr(unsafe.Pointer(sent)), uintptr(flags), uintptr(unsafe.Pointer(overlapped)), uintptr(unsafe.Pointer(croutine)), 0, 0)
	if r1 == socket_error {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  WinsockSendTo(socket  手柄,  buffers  *Wsabuf,  bufferCount  uint32,  sent  *uint32,  flags  uint32,  destination  *RawSockaddrAny,  destinationLength  int32,  overlapped  *Overlapped,  completionRoutine  *byte)  (error  错误)  {}

// ff:
// s:
// bufs:
// bufcnt:
// sent:
// flags:
// to:
// tolen:
// overlapped:
// croutine:
// err:
func WSASendTo(s Handle, bufs *WSABuf, bufcnt uint32, sent *uint32, flags uint32, to *RawSockaddrAny, tolen int32, overlapped *Overlapped, croutine *byte) (err error) {
	r1, _, e1 := syscall.Syscall9(procWSASendTo.Addr(), 9, uintptr(s), uintptr(unsafe.Pointer(bufs)), uintptr(bufcnt), uintptr(unsafe.Pointer(sent)), uintptr(flags), uintptr(unsafe.Pointer(to)), uintptr(tolen), uintptr(unsafe.Pointer(overlapped)), uintptr(unsafe.Pointer(croutine)))
	if r1 == socket_error {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  创建网络套接字(地址族  af  int32,  类型  typ  int32,  协议  protocol  int32,  协议信息  protoInfo  *WSAProtocolInfo,  组  group  uint32,  标志  flags  uint32)  (句柄  handle  Handle,  错误  err  error)  {}

// ff:
// af:
// typ:
// protocol:
// protoInfo:
// group:
// flags:
// handle:
// err:
func WSASocket(af int32, typ int32, protocol int32, protoInfo *WSAProtocolInfo, group uint32, flags uint32) (handle Handle, err error) {
	r0, _, e1 := syscall.Syscall6(procWSASocketW.Addr(), 6, uintptr(af), uintptr(typ), uintptr(protocol), uintptr(unsafe.Pointer(protoInfo)), uintptr(group), uintptr(flags))
	handle = Handle(r0)
	if handle == InvalidHandle {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  套接字启动(versionRequest  uint32,  数据  *WSAData)  (错误  error)  {}

// ff:
// verreq:
// data:
// sockerr:
func WSAStartup(verreq uint32, data *WSAData) (sockerr error) {
	r0, _, _ := syscall.Syscall(procWSAStartup.Addr(), 2, uintptr(verreq), uintptr(unsafe.Pointer(data)), 0)
	if r0 != 0 {
		sockerr = syscall.Errno(r0)
	}
	return
}

func bind(s Handle, name unsafe.Pointer, namelen int32) (err error) {
	r1, _, e1 := syscall.Syscall(procbind.Addr(), 3, uintptr(s), uintptr(name), uintptr(namelen))
	if r1 == socket_error {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  关闭套接字(socket  手柄)  (错误  error)  {}

// ff:
// s:
// err:
func Closesocket(s Handle) (err error) {
	r1, _, e1 := syscall.Syscall(procclosesocket.Addr(), 1, uintptr(s), 0, 0)
	if r1 == socket_error {
		err = errnoErr(e1)
	}
	return
}

func connect(s Handle, name unsafe.Pointer, namelen int32) (err error) {
	r1, _, e1 := syscall.Syscall(procconnect.Addr(), 3, uintptr(s), uintptr(name), uintptr(namelen))
	if r1 == socket_error {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  获取主机名通过名称(name  string)  (主机信息*Hostent,  错误error)  {}

// ff:
// name:
// h:
// err:
func GetHostByName(name string) (h *Hostent, err error) {
	var _p0 *byte
	_p0, err = syscall.BytePtrFromString(name)
	if err != nil {
		return
	}
	return _GetHostByName(_p0)
}

func _GetHostByName(name *byte) (h *Hostent, err error) {
	r0, _, e1 := syscall.Syscall(procgethostbyname.Addr(), 1, uintptr(unsafe.Pointer(name)), 0, 0)
	h = (*Hostent)(unsafe.Pointer(r0))
	if h == nil {
		err = errnoErr(e1)
	}
	return
}

func getpeername(s Handle, rsa *RawSockaddrAny, addrlen *int32) (err error) {
	r1, _, e1 := syscall.Syscall(procgetpeername.Addr(), 3, uintptr(s), uintptr(unsafe.Pointer(rsa)), uintptr(unsafe.Pointer(addrlen)))
	if r1 == socket_error {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  获取协议ByName(名称  string)  (协议信息  *Protoent,  错误  error)  {}

// ff:
// name:
// p:
// err:
func GetProtoByName(name string) (p *Protoent, err error) {
	var _p0 *byte
	_p0, err = syscall.BytePtrFromString(name)
	if err != nil {
		return
	}
	return _GetProtoByName(_p0)
}

func _GetProtoByName(name *byte) (p *Protoent, err error) {
	r0, _, e1 := syscall.Syscall(procgetprotobyname.Addr(), 1, uintptr(unsafe.Pointer(name)), 0, 0)
	p = (*Protoent)(unsafe.Pointer(r0))
	if p == nil {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  获取服务通过名称(name  string,  协议  string)  (服务信息  *Servent,  错误  error)  {}

// ff:
// name:
// proto:
// s:
// err:
func GetServByName(name string, proto string) (s *Servent, err error) {
	var _p0 *byte
	_p0, err = syscall.BytePtrFromString(name)
	if err != nil {
		return
	}
	var _p1 *byte
	_p1, err = syscall.BytePtrFromString(proto)
	if err != nil {
		return
	}
	return _GetServByName(_p0, _p1)
}

func _GetServByName(name *byte, proto *byte) (s *Servent, err error) {
	r0, _, e1 := syscall.Syscall(procgetservbyname.Addr(), 2, uintptr(unsafe.Pointer(name)), uintptr(unsafe.Pointer(proto)), 0)
	s = (*Servent)(unsafe.Pointer(r0))
	if s == nil {
		err = errnoErr(e1)
	}
	return
}

func getsockname(s Handle, rsa *RawSockaddrAny, addrlen *int32) (err error) {
	r1, _, e1 := syscall.Syscall(procgetsockname.Addr(), 3, uintptr(s), uintptr(unsafe.Pointer(rsa)), uintptr(unsafe.Pointer(addrlen)))
	if r1 == socket_error {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  设置套接字选项(s  手柄,  层级  int32,  选项名  int32,  选项值  *byte,  选项长度  *int32)  (错误  error)  {}

// ff:
// s:
// level:
// optname:
// optval:
// optlen:
// err:
func Getsockopt(s Handle, level int32, optname int32, optval *byte, optlen *int32) (err error) {
	r1, _, e1 := syscall.Syscall6(procgetsockopt.Addr(), 5, uintptr(s), uintptr(level), uintptr(optname), uintptr(unsafe.Pointer(optval)), uintptr(unsafe.Pointer(optlen)), 0)
	if r1 == socket_error {
		err = errnoErr(e1)
	}
	return
}

func listen(s Handle, backlog int32) (err error) {
	r1, _, e1 := syscall.Syscall(proclisten.Addr(), 2, uintptr(s), uintptr(backlog), 0)
	if r1 == socket_error {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  网短整型转主机短整型(netshort  网络短整型)  (u  主机短整型)  {}

// ff:
// netshort:
// u:
func Ntohs(netshort uint16) (u uint16) {
	r0, _, _ := syscall.Syscall(procntohs.Addr(), 1, uintptr(netshort), 0, 0)
	u = uint16(r0)
	return
}

func recvfrom(s Handle, buf []byte, flags int32, from *RawSockaddrAny, fromlen *int32) (n int32, err error) {
	var _p0 *byte
	if len(buf) > 0 {
		_p0 = &buf[0]
	}
	r0, _, e1 := syscall.Syscall6(procrecvfrom.Addr(), 6, uintptr(s), uintptr(unsafe.Pointer(_p0)), uintptr(len(buf)), uintptr(flags), uintptr(unsafe.Pointer(from)), uintptr(unsafe.Pointer(fromlen)))
	n = int32(r0)
	if n == -1 {
		err = errnoErr(e1)
	}
	return
}

func sendto(s Handle, buf []byte, flags int32, to unsafe.Pointer, tolen int32) (err error) {
	var _p0 *byte
	if len(buf) > 0 {
		_p0 = &buf[0]
	}
	r1, _, e1 := syscall.Syscall6(procsendto.Addr(), 6, uintptr(s), uintptr(unsafe.Pointer(_p0)), uintptr(len(buf)), uintptr(flags), uintptr(to), uintptr(tolen))
	if r1 == socket_error {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  设置套接字选项(s  手柄,  级别  int32,  选项名  int32,  选项值  *byte,  选项长度  int32)  (错误  error)  {}

// ff:
// s:
// level:
// optname:
// optval:
// optlen:
// err:
func Setsockopt(s Handle, level int32, optname int32, optval *byte, optlen int32) (err error) {
	r1, _, e1 := syscall.Syscall6(procsetsockopt.Addr(), 5, uintptr(s), uintptr(level), uintptr(optname), uintptr(unsafe.Pointer(optval)), uintptr(optlen), 0)
	if r1 == socket_error {
		err = errnoErr(e1)
	}
	return
}

func shutdown(s Handle, how int32) (err error) {
	r1, _, e1 := syscall.Syscall(procshutdown.Addr(), 2, uintptr(s), uintptr(how), 0)
	if r1 == socket_error {
		err = errnoErr(e1)
	}
	return
}

func socket(af int32, typ int32, protocol int32) (handle Handle, err error) {
	r0, _, e1 := syscall.Syscall(procsocket.Addr(), 3, uintptr(af), uintptr(typ), uintptr(protocol))
	handle = Handle(r0)
	if handle == InvalidHandle {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  列举会话(handle  会话句柄,  保留  uint32,  版本  uint32,  会话信息  **会话信息结构体,  会话计数  *uint32)  (错误  error)  {}

// ff:
// handle:
// reserved:
// version:
// sessions:
// count:
// err:
func WTSEnumerateSessions(handle Handle, reserved uint32, version uint32, sessions **WTS_SESSION_INFO, count *uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procWTSEnumerateSessionsW.Addr(), 5, uintptr(handle), uintptr(reserved), uintptr(version), uintptr(unsafe.Pointer(sessions)), uintptr(unsafe.Pointer(count)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

// 翻译提示:func  释放内存(ptr  指针  uintptr)  {}

// ff:
// ptr:
func WTSFreeMemory(ptr uintptr) {
	syscall.Syscall(procWTSFreeMemory.Addr(), 1, uintptr(ptr), 0, 0)
	return
}

// 翻译提示:func  查询用户令牌(session  uint32,  token  *用户令牌指针)  (错误  error)  {}

// ff:
// session:
// token:
// err:
func WTSQueryUserToken(session uint32, token *Token) (err error) {
	r1, _, e1 := syscall.Syscall(procWTSQueryUserToken.Addr(), 2, uintptr(session), uintptr(unsafe.Pointer(token)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}
