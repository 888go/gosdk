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

var (
	errERROR_IO_PENDING error = syscall.Errno(errnoERROR_IO_PENDING)
	errERROR_EINVAL     error = syscall.EINVAL
)

// errnoErr 返回常见的装箱 Errno 值，以防止运行时的分配。
func errnoErr(e syscall.Errno) error {
	switch e {
	case 0:
		return errERROR_EINVAL
	case errnoERROR_IO_PENDING:
		return errERROR_IO_PENDING
	}
// TODO: 在收集到Windows上常见的错误值数据后，此处应添加更多内容。（或许在运行all.bat时？）
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

func ChangeServiceConfig2(service Handle, infoLevel uint32, info *byte) (err error) {
	r1, _, e1 := syscall.Syscall(procChangeServiceConfig2W.Addr(), 3, uintptr(service), uintptr(infoLevel), uintptr(unsafe.Pointer(info)))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

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

func CloseServiceHandle(handle Handle) (err error) {
	r1, _, e1 := syscall.Syscall(procCloseServiceHandle.Addr(), 1, uintptr(handle), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

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

func ConvertStringSidToSid(stringSid *uint16, sid **SID) (err error) {
	r1, _, e1 := syscall.Syscall(procConvertStringSidToSidW.Addr(), 2, uintptr(unsafe.Pointer(stringSid)), uintptr(unsafe.Pointer(sid)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func CopySid(destSidLen uint32, destSid *SID, srcSid *SID) (err error) {
	r1, _, e1 := syscall.Syscall(procCopySid.Addr(), 3, uintptr(destSidLen), uintptr(unsafe.Pointer(destSid)), uintptr(unsafe.Pointer(srcSid)))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

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

func CryptAcquireContext(provhandle *Handle, container *uint16, provider *uint16, provtype uint32, flags uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procCryptAcquireContextW.Addr(), 5, uintptr(unsafe.Pointer(provhandle)), uintptr(unsafe.Pointer(container)), uintptr(unsafe.Pointer(provider)), uintptr(provtype), uintptr(flags), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func CryptGenRandom(provhandle Handle, buflen uint32, buf *byte) (err error) {
	r1, _, e1 := syscall.Syscall(procCryptGenRandom.Addr(), 3, uintptr(provhandle), uintptr(buflen), uintptr(unsafe.Pointer(buf)))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func CryptReleaseContext(provhandle Handle, flags uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procCryptReleaseContext.Addr(), 2, uintptr(provhandle), uintptr(flags), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func DeleteService(service Handle) (err error) {
	r1, _, e1 := syscall.Syscall(procDeleteService.Addr(), 1, uintptr(service), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func DeregisterEventSource(handle Handle) (err error) {
	r1, _, e1 := syscall.Syscall(procDeregisterEventSource.Addr(), 1, uintptr(handle), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func DuplicateTokenEx(existingToken Token, desiredAccess uint32, tokenAttributes *SecurityAttributes, impersonationLevel uint32, tokenType uint32, newToken *Token) (err error) {
	r1, _, e1 := syscall.Syscall6(procDuplicateTokenEx.Addr(), 6, uintptr(existingToken), uintptr(desiredAccess), uintptr(unsafe.Pointer(tokenAttributes)), uintptr(impersonationLevel), uintptr(tokenType), uintptr(unsafe.Pointer(newToken)))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func EnumDependentServices(service Handle, activityState uint32, services *ENUM_SERVICE_STATUS, buffSize uint32, bytesNeeded *uint32, servicesReturned *uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procEnumDependentServicesW.Addr(), 6, uintptr(service), uintptr(activityState), uintptr(unsafe.Pointer(services)), uintptr(buffSize), uintptr(unsafe.Pointer(bytesNeeded)), uintptr(unsafe.Pointer(servicesReturned)))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func EnumServicesStatusEx(mgr Handle, infoLevel uint32, serviceType uint32, serviceState uint32, services *byte, bufSize uint32, bytesNeeded *uint32, servicesReturned *uint32, resumeHandle *uint32, groupName *uint16) (err error) {
	r1, _, e1 := syscall.Syscall12(procEnumServicesStatusExW.Addr(), 10, uintptr(mgr), uintptr(infoLevel), uintptr(serviceType), uintptr(serviceState), uintptr(unsafe.Pointer(services)), uintptr(bufSize), uintptr(unsafe.Pointer(bytesNeeded)), uintptr(unsafe.Pointer(servicesReturned)), uintptr(unsafe.Pointer(resumeHandle)), uintptr(unsafe.Pointer(groupName)), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func EqualSid(sid1 *SID, sid2 *SID) (isEqual bool) {
	r0, _, _ := syscall.Syscall(procEqualSid.Addr(), 2, uintptr(unsafe.Pointer(sid1)), uintptr(unsafe.Pointer(sid2)), 0)
	isEqual = r0 != 0
	return
}

func FreeSid(sid *SID) (err error) {
	r1, _, e1 := syscall.Syscall(procFreeSid.Addr(), 1, uintptr(unsafe.Pointer(sid)), 0, 0)
	if r1 != 0 {
		err = errnoErr(e1)
	}
	return
}

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

func GetTokenInformation(token Token, infoClass uint32, info *byte, infoLen uint32, returnedLen *uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procGetTokenInformation.Addr(), 5, uintptr(token), uintptr(infoClass), uintptr(unsafe.Pointer(info)), uintptr(infoLen), uintptr(unsafe.Pointer(returnedLen)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

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

func LookupAccountName(systemName *uint16, accountName *uint16, sid *SID, sidLen *uint32, refdDomainName *uint16, refdDomainNameLen *uint32, use *uint32) (err error) {
	r1, _, e1 := syscall.Syscall9(procLookupAccountNameW.Addr(), 7, uintptr(unsafe.Pointer(systemName)), uintptr(unsafe.Pointer(accountName)), uintptr(unsafe.Pointer(sid)), uintptr(unsafe.Pointer(sidLen)), uintptr(unsafe.Pointer(refdDomainName)), uintptr(unsafe.Pointer(refdDomainNameLen)), uintptr(unsafe.Pointer(use)), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func LookupAccountSid(systemName *uint16, sid *SID, name *uint16, nameLen *uint32, refdDomainName *uint16, refdDomainNameLen *uint32, use *uint32) (err error) {
	r1, _, e1 := syscall.Syscall9(procLookupAccountSidW.Addr(), 7, uintptr(unsafe.Pointer(systemName)), uintptr(unsafe.Pointer(sid)), uintptr(unsafe.Pointer(name)), uintptr(unsafe.Pointer(nameLen)), uintptr(unsafe.Pointer(refdDomainName)), uintptr(unsafe.Pointer(refdDomainNameLen)), uintptr(unsafe.Pointer(use)), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

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

func NotifyServiceStatusChange(service Handle, notifyMask uint32, notifier *SERVICE_NOTIFY) (ret error) {
	r0, _, _ := syscall.Syscall(procNotifyServiceStatusChangeW.Addr(), 3, uintptr(service), uintptr(notifyMask), uintptr(unsafe.Pointer(notifier)))
	if r0 != 0 {
		ret = syscall.Errno(r0)
	}
	return
}

func OpenProcessToken(process Handle, access uint32, token *Token) (err error) {
	r1, _, e1 := syscall.Syscall(procOpenProcessToken.Addr(), 3, uintptr(process), uintptr(access), uintptr(unsafe.Pointer(token)))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func OpenSCManager(machineName *uint16, databaseName *uint16, access uint32) (handle Handle, err error) {
	r0, _, e1 := syscall.Syscall(procOpenSCManagerW.Addr(), 3, uintptr(unsafe.Pointer(machineName)), uintptr(unsafe.Pointer(databaseName)), uintptr(access))
	handle = Handle(r0)
	if handle == 0 {
		err = errnoErr(e1)
	}
	return
}

func OpenService(mgr Handle, serviceName *uint16, access uint32) (handle Handle, err error) {
	r0, _, e1 := syscall.Syscall(procOpenServiceW.Addr(), 3, uintptr(mgr), uintptr(unsafe.Pointer(serviceName)), uintptr(access))
	handle = Handle(r0)
	if handle == 0 {
		err = errnoErr(e1)
	}
	return
}

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

func QueryServiceConfig2(service Handle, infoLevel uint32, buff *byte, buffSize uint32, bytesNeeded *uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procQueryServiceConfig2W.Addr(), 5, uintptr(service), uintptr(infoLevel), uintptr(unsafe.Pointer(buff)), uintptr(buffSize), uintptr(unsafe.Pointer(bytesNeeded)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func QueryServiceConfig(service Handle, serviceConfig *QUERY_SERVICE_CONFIG, bufSize uint32, bytesNeeded *uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procQueryServiceConfigW.Addr(), 4, uintptr(service), uintptr(unsafe.Pointer(serviceConfig)), uintptr(bufSize), uintptr(unsafe.Pointer(bytesNeeded)), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

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

func QueryServiceLockStatus(mgr Handle, lockStatus *QUERY_SERVICE_LOCK_STATUS, bufSize uint32, bytesNeeded *uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procQueryServiceLockStatusW.Addr(), 4, uintptr(mgr), uintptr(unsafe.Pointer(lockStatus)), uintptr(bufSize), uintptr(unsafe.Pointer(bytesNeeded)), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func QueryServiceStatus(service Handle, status *SERVICE_STATUS) (err error) {
	r1, _, e1 := syscall.Syscall(procQueryServiceStatus.Addr(), 2, uintptr(service), uintptr(unsafe.Pointer(status)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func QueryServiceStatusEx(service Handle, infoLevel uint32, buff *byte, buffSize uint32, bytesNeeded *uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procQueryServiceStatusEx.Addr(), 5, uintptr(service), uintptr(infoLevel), uintptr(unsafe.Pointer(buff)), uintptr(buffSize), uintptr(unsafe.Pointer(bytesNeeded)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func RegCloseKey(key Handle) (regerrno error) {
	r0, _, _ := syscall.Syscall(procRegCloseKey.Addr(), 1, uintptr(key), 0, 0)
	if r0 != 0 {
		regerrno = syscall.Errno(r0)
	}
	return
}

func RegEnumKeyEx(key Handle, index uint32, name *uint16, nameLen *uint32, reserved *uint32, class *uint16, classLen *uint32, lastWriteTime *Filetime) (regerrno error) {
	r0, _, _ := syscall.Syscall9(procRegEnumKeyExW.Addr(), 8, uintptr(key), uintptr(index), uintptr(unsafe.Pointer(name)), uintptr(unsafe.Pointer(nameLen)), uintptr(unsafe.Pointer(reserved)), uintptr(unsafe.Pointer(class)), uintptr(unsafe.Pointer(classLen)), uintptr(unsafe.Pointer(lastWriteTime)), 0)
	if r0 != 0 {
		regerrno = syscall.Errno(r0)
	}
	return
}

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

func RegOpenKeyEx(key Handle, subkey *uint16, options uint32, desiredAccess uint32, result *Handle) (regerrno error) {
	r0, _, _ := syscall.Syscall6(procRegOpenKeyExW.Addr(), 5, uintptr(key), uintptr(unsafe.Pointer(subkey)), uintptr(options), uintptr(desiredAccess), uintptr(unsafe.Pointer(result)), 0)
	if r0 != 0 {
		regerrno = syscall.Errno(r0)
	}
	return
}

func RegQueryInfoKey(key Handle, class *uint16, classLen *uint32, reserved *uint32, subkeysLen *uint32, maxSubkeyLen *uint32, maxClassLen *uint32, valuesLen *uint32, maxValueNameLen *uint32, maxValueLen *uint32, saLen *uint32, lastWriteTime *Filetime) (regerrno error) {
	r0, _, _ := syscall.Syscall12(procRegQueryInfoKeyW.Addr(), 12, uintptr(key), uintptr(unsafe.Pointer(class)), uintptr(unsafe.Pointer(classLen)), uintptr(unsafe.Pointer(reserved)), uintptr(unsafe.Pointer(subkeysLen)), uintptr(unsafe.Pointer(maxSubkeyLen)), uintptr(unsafe.Pointer(maxClassLen)), uintptr(unsafe.Pointer(valuesLen)), uintptr(unsafe.Pointer(maxValueNameLen)), uintptr(unsafe.Pointer(maxValueLen)), uintptr(unsafe.Pointer(saLen)), uintptr(unsafe.Pointer(lastWriteTime)))
	if r0 != 0 {
		regerrno = syscall.Errno(r0)
	}
	return
}

func RegQueryValueEx(key Handle, name *uint16, reserved *uint32, valtype *uint32, buf *byte, buflen *uint32) (regerrno error) {
	r0, _, _ := syscall.Syscall6(procRegQueryValueExW.Addr(), 6, uintptr(key), uintptr(unsafe.Pointer(name)), uintptr(unsafe.Pointer(reserved)), uintptr(unsafe.Pointer(valtype)), uintptr(unsafe.Pointer(buf)), uintptr(unsafe.Pointer(buflen)))
	if r0 != 0 {
		regerrno = syscall.Errno(r0)
	}
	return
}

func RegisterEventSource(uncServerName *uint16, sourceName *uint16) (handle Handle, err error) {
	r0, _, e1 := syscall.Syscall(procRegisterEventSourceW.Addr(), 2, uintptr(unsafe.Pointer(uncServerName)), uintptr(unsafe.Pointer(sourceName)), 0)
	handle = Handle(r0)
	if handle == 0 {
		err = errnoErr(e1)
	}
	return
}

func RegisterServiceCtrlHandlerEx(serviceName *uint16, handlerProc uintptr, context uintptr) (handle Handle, err error) {
	r0, _, e1 := syscall.Syscall(procRegisterServiceCtrlHandlerExW.Addr(), 3, uintptr(unsafe.Pointer(serviceName)), uintptr(handlerProc), uintptr(context))
	handle = Handle(r0)
	if handle == 0 {
		err = errnoErr(e1)
	}
	return
}

func ReportEvent(log Handle, etype uint16, category uint16, eventId uint32, usrSId uintptr, numStrings uint16, dataSize uint32, strings **uint16, rawData *byte) (err error) {
	r1, _, e1 := syscall.Syscall9(procReportEventW.Addr(), 9, uintptr(log), uintptr(etype), uintptr(category), uintptr(eventId), uintptr(usrSId), uintptr(numStrings), uintptr(dataSize), uintptr(unsafe.Pointer(strings)), uintptr(unsafe.Pointer(rawData)))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

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

func SetKernelObjectSecurity(handle Handle, securityInformation SECURITY_INFORMATION, securityDescriptor *SECURITY_DESCRIPTOR) (err error) {
	r1, _, e1 := syscall.Syscall(procSetKernelObjectSecurity.Addr(), 3, uintptr(handle), uintptr(securityInformation), uintptr(unsafe.Pointer(securityDescriptor)))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

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

func SetSecurityInfo(handle Handle, objectType SE_OBJECT_TYPE, securityInformation SECURITY_INFORMATION, owner *SID, group *SID, dacl *ACL, sacl *ACL) (ret error) {
	r0, _, _ := syscall.Syscall9(procSetSecurityInfo.Addr(), 7, uintptr(handle), uintptr(objectType), uintptr(securityInformation), uintptr(unsafe.Pointer(owner)), uintptr(unsafe.Pointer(group)), uintptr(unsafe.Pointer(dacl)), uintptr(unsafe.Pointer(sacl)), 0, 0)
	if r0 != 0 {
		ret = syscall.Errno(r0)
	}
	return
}

func SetServiceStatus(service Handle, serviceStatus *SERVICE_STATUS) (err error) {
	r1, _, e1 := syscall.Syscall(procSetServiceStatus.Addr(), 2, uintptr(service), uintptr(unsafe.Pointer(serviceStatus)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func SetThreadToken(thread *Handle, token Token) (err error) {
	r1, _, e1 := syscall.Syscall(procSetThreadToken.Addr(), 2, uintptr(unsafe.Pointer(thread)), uintptr(token), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func SetTokenInformation(token Token, infoClass uint32, info *byte, infoLen uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procSetTokenInformation.Addr(), 4, uintptr(token), uintptr(infoClass), uintptr(unsafe.Pointer(info)), uintptr(infoLen), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func StartServiceCtrlDispatcher(serviceTable *SERVICE_TABLE_ENTRY) (err error) {
	r1, _, e1 := syscall.Syscall(procStartServiceCtrlDispatcherW.Addr(), 1, uintptr(unsafe.Pointer(serviceTable)), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func StartService(service Handle, numArgs uint32, argVectors **uint16) (err error) {
	r1, _, e1 := syscall.Syscall(procStartServiceW.Addr(), 3, uintptr(service), uintptr(numArgs), uintptr(unsafe.Pointer(argVectors)))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func CertAddCertificateContextToStore(store Handle, certContext *CertContext, addDisposition uint32, storeContext **CertContext) (err error) {
	r1, _, e1 := syscall.Syscall6(procCertAddCertificateContextToStore.Addr(), 4, uintptr(store), uintptr(unsafe.Pointer(certContext)), uintptr(addDisposition), uintptr(unsafe.Pointer(storeContext)), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func CertCloseStore(store Handle, flags uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procCertCloseStore.Addr(), 2, uintptr(store), uintptr(flags), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func CertCreateCertificateContext(certEncodingType uint32, certEncoded *byte, encodedLen uint32) (context *CertContext, err error) {
	r0, _, e1 := syscall.Syscall(procCertCreateCertificateContext.Addr(), 3, uintptr(certEncodingType), uintptr(unsafe.Pointer(certEncoded)), uintptr(encodedLen))
	context = (*CertContext)(unsafe.Pointer(r0))
	if context == nil {
		err = errnoErr(e1)
	}
	return
}

func CertDeleteCertificateFromStore(certContext *CertContext) (err error) {
	r1, _, e1 := syscall.Syscall(procCertDeleteCertificateFromStore.Addr(), 1, uintptr(unsafe.Pointer(certContext)), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func CertDuplicateCertificateContext(certContext *CertContext) (dupContext *CertContext) {
	r0, _, _ := syscall.Syscall(procCertDuplicateCertificateContext.Addr(), 1, uintptr(unsafe.Pointer(certContext)), 0, 0)
	dupContext = (*CertContext)(unsafe.Pointer(r0))
	return
}

func CertEnumCertificatesInStore(store Handle, prevContext *CertContext) (context *CertContext, err error) {
	r0, _, e1 := syscall.Syscall(procCertEnumCertificatesInStore.Addr(), 2, uintptr(store), uintptr(unsafe.Pointer(prevContext)), 0)
	context = (*CertContext)(unsafe.Pointer(r0))
	if context == nil {
		err = errnoErr(e1)
	}
	return
}

func CertFindCertificateInStore(store Handle, certEncodingType uint32, findFlags uint32, findType uint32, findPara unsafe.Pointer, prevCertContext *CertContext) (cert *CertContext, err error) {
	r0, _, e1 := syscall.Syscall6(procCertFindCertificateInStore.Addr(), 6, uintptr(store), uintptr(certEncodingType), uintptr(findFlags), uintptr(findType), uintptr(findPara), uintptr(unsafe.Pointer(prevCertContext)))
	cert = (*CertContext)(unsafe.Pointer(r0))
	if cert == nil {
		err = errnoErr(e1)
	}
	return
}

func CertFindChainInStore(store Handle, certEncodingType uint32, findFlags uint32, findType uint32, findPara unsafe.Pointer, prevChainContext *CertChainContext) (certchain *CertChainContext, err error) {
	r0, _, e1 := syscall.Syscall6(procCertFindChainInStore.Addr(), 6, uintptr(store), uintptr(certEncodingType), uintptr(findFlags), uintptr(findType), uintptr(findPara), uintptr(unsafe.Pointer(prevChainContext)))
	certchain = (*CertChainContext)(unsafe.Pointer(r0))
	if certchain == nil {
		err = errnoErr(e1)
	}
	return
}

func CertFindExtension(objId *byte, countExtensions uint32, extensions *CertExtension) (ret *CertExtension) {
	r0, _, _ := syscall.Syscall(procCertFindExtension.Addr(), 3, uintptr(unsafe.Pointer(objId)), uintptr(countExtensions), uintptr(unsafe.Pointer(extensions)))
	ret = (*CertExtension)(unsafe.Pointer(r0))
	return
}

func CertFreeCertificateChain(ctx *CertChainContext) {
	syscall.Syscall(procCertFreeCertificateChain.Addr(), 1, uintptr(unsafe.Pointer(ctx)), 0, 0)
	return
}

func CertFreeCertificateContext(ctx *CertContext) (err error) {
	r1, _, e1 := syscall.Syscall(procCertFreeCertificateContext.Addr(), 1, uintptr(unsafe.Pointer(ctx)), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func CertGetCertificateChain(engine Handle, leaf *CertContext, time *Filetime, additionalStore Handle, para *CertChainPara, flags uint32, reserved uintptr, chainCtx **CertChainContext) (err error) {
	r1, _, e1 := syscall.Syscall9(procCertGetCertificateChain.Addr(), 8, uintptr(engine), uintptr(unsafe.Pointer(leaf)), uintptr(unsafe.Pointer(time)), uintptr(additionalStore), uintptr(unsafe.Pointer(para)), uintptr(flags), uintptr(reserved), uintptr(unsafe.Pointer(chainCtx)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func CertGetNameString(certContext *CertContext, nameType uint32, flags uint32, typePara unsafe.Pointer, name *uint16, size uint32) (chars uint32) {
	r0, _, _ := syscall.Syscall6(procCertGetNameStringW.Addr(), 6, uintptr(unsafe.Pointer(certContext)), uintptr(nameType), uintptr(flags), uintptr(typePara), uintptr(unsafe.Pointer(name)), uintptr(size))
	chars = uint32(r0)
	return
}

func CertOpenStore(storeProvider uintptr, msgAndCertEncodingType uint32, cryptProv uintptr, flags uint32, para uintptr) (handle Handle, err error) {
	r0, _, e1 := syscall.Syscall6(procCertOpenStore.Addr(), 5, uintptr(storeProvider), uintptr(msgAndCertEncodingType), uintptr(cryptProv), uintptr(flags), uintptr(para), 0)
	handle = Handle(r0)
	if handle == 0 {
		err = errnoErr(e1)
	}
	return
}

func CertOpenSystemStore(hprov Handle, name *uint16) (store Handle, err error) {
	r0, _, e1 := syscall.Syscall(procCertOpenSystemStoreW.Addr(), 2, uintptr(hprov), uintptr(unsafe.Pointer(name)), 0)
	store = Handle(r0)
	if store == 0 {
		err = errnoErr(e1)
	}
	return
}

func CertVerifyCertificateChainPolicy(policyOID uintptr, chain *CertChainContext, para *CertChainPolicyPara, status *CertChainPolicyStatus) (err error) {
	r1, _, e1 := syscall.Syscall6(procCertVerifyCertificateChainPolicy.Addr(), 4, uintptr(policyOID), uintptr(unsafe.Pointer(chain)), uintptr(unsafe.Pointer(para)), uintptr(unsafe.Pointer(status)), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

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

func CryptDecodeObject(encodingType uint32, structType *byte, encodedBytes *byte, lenEncodedBytes uint32, flags uint32, decoded unsafe.Pointer, decodedLen *uint32) (err error) {
	r1, _, e1 := syscall.Syscall9(procCryptDecodeObject.Addr(), 7, uintptr(encodingType), uintptr(unsafe.Pointer(structType)), uintptr(unsafe.Pointer(encodedBytes)), uintptr(lenEncodedBytes), uintptr(flags), uintptr(decoded), uintptr(unsafe.Pointer(decodedLen)), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func CryptProtectData(dataIn *DataBlob, name *uint16, optionalEntropy *DataBlob, reserved uintptr, promptStruct *CryptProtectPromptStruct, flags uint32, dataOut *DataBlob) (err error) {
	r1, _, e1 := syscall.Syscall9(procCryptProtectData.Addr(), 7, uintptr(unsafe.Pointer(dataIn)), uintptr(unsafe.Pointer(name)), uintptr(unsafe.Pointer(optionalEntropy)), uintptr(reserved), uintptr(unsafe.Pointer(promptStruct)), uintptr(flags), uintptr(unsafe.Pointer(dataOut)), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func CryptQueryObject(objectType uint32, object unsafe.Pointer, expectedContentTypeFlags uint32, expectedFormatTypeFlags uint32, flags uint32, msgAndCertEncodingType *uint32, contentType *uint32, formatType *uint32, certStore *Handle, msg *Handle, context *unsafe.Pointer) (err error) {
	r1, _, e1 := syscall.Syscall12(procCryptQueryObject.Addr(), 11, uintptr(objectType), uintptr(object), uintptr(expectedContentTypeFlags), uintptr(expectedFormatTypeFlags), uintptr(flags), uintptr(unsafe.Pointer(msgAndCertEncodingType)), uintptr(unsafe.Pointer(contentType)), uintptr(unsafe.Pointer(formatType)), uintptr(unsafe.Pointer(certStore)), uintptr(unsafe.Pointer(msg)), uintptr(unsafe.Pointer(context)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func CryptUnprotectData(dataIn *DataBlob, name **uint16, optionalEntropy *DataBlob, reserved uintptr, promptStruct *CryptProtectPromptStruct, flags uint32, dataOut *DataBlob) (err error) {
	r1, _, e1 := syscall.Syscall9(procCryptUnprotectData.Addr(), 7, uintptr(unsafe.Pointer(dataIn)), uintptr(unsafe.Pointer(name)), uintptr(unsafe.Pointer(optionalEntropy)), uintptr(reserved), uintptr(unsafe.Pointer(promptStruct)), uintptr(flags), uintptr(unsafe.Pointer(dataOut)), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func PFXImportCertStore(pfx *CryptDataBlob, password *uint16, flags uint32) (store Handle, err error) {
	r0, _, e1 := syscall.Syscall(procPFXImportCertStore.Addr(), 3, uintptr(unsafe.Pointer(pfx)), uintptr(unsafe.Pointer(password)), uintptr(flags))
	store = Handle(r0)
	if store == 0 {
		err = errnoErr(e1)
	}
	return
}

func DnsNameCompare(name1 *uint16, name2 *uint16) (same bool) {
	r0, _, _ := syscall.Syscall(procDnsNameCompare_W.Addr(), 2, uintptr(unsafe.Pointer(name1)), uintptr(unsafe.Pointer(name2)), 0)
	same = r0 != 0
	return
}

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

func DnsRecordListFree(rl *DNSRecord, freetype uint32) {
	syscall.Syscall(procDnsRecordListFree.Addr(), 2, uintptr(unsafe.Pointer(rl)), uintptr(freetype), 0)
	return
}

func DwmGetWindowAttribute(hwnd HWND, attribute uint32, value unsafe.Pointer, size uint32) (ret error) {
	r0, _, _ := syscall.Syscall6(procDwmGetWindowAttribute.Addr(), 4, uintptr(hwnd), uintptr(attribute), uintptr(value), uintptr(size), 0, 0)
	if r0 != 0 {
		ret = syscall.Errno(r0)
	}
	return
}

func DwmSetWindowAttribute(hwnd HWND, attribute uint32, value unsafe.Pointer, size uint32) (ret error) {
	r0, _, _ := syscall.Syscall6(procDwmSetWindowAttribute.Addr(), 4, uintptr(hwnd), uintptr(attribute), uintptr(value), uintptr(size), 0, 0)
	if r0 != 0 {
		ret = syscall.Errno(r0)
	}
	return
}

func GetAdaptersAddresses(family uint32, flags uint32, reserved uintptr, adapterAddresses *IpAdapterAddresses, sizePointer *uint32) (errcode error) {
	r0, _, _ := syscall.Syscall6(procGetAdaptersAddresses.Addr(), 5, uintptr(family), uintptr(flags), uintptr(reserved), uintptr(unsafe.Pointer(adapterAddresses)), uintptr(unsafe.Pointer(sizePointer)), 0)
	if r0 != 0 {
		errcode = syscall.Errno(r0)
	}
	return
}

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

func GetIfEntry(pIfRow *MibIfRow) (errcode error) {
	r0, _, _ := syscall.Syscall(procGetIfEntry.Addr(), 1, uintptr(unsafe.Pointer(pIfRow)), 0, 0)
	if r0 != 0 {
		errcode = syscall.Errno(r0)
	}
	return
}

func AddDllDirectory(path *uint16) (cookie uintptr, err error) {
	r0, _, e1 := syscall.Syscall(procAddDllDirectory.Addr(), 1, uintptr(unsafe.Pointer(path)), 0, 0)
	cookie = uintptr(r0)
	if cookie == 0 {
		err = errnoErr(e1)
	}
	return
}

func AssignProcessToJobObject(job Handle, process Handle) (err error) {
	r1, _, e1 := syscall.Syscall(procAssignProcessToJobObject.Addr(), 2, uintptr(job), uintptr(process), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func CancelIo(s Handle) (err error) {
	r1, _, e1 := syscall.Syscall(procCancelIo.Addr(), 1, uintptr(s), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func CancelIoEx(s Handle, o *Overlapped) (err error) {
	r1, _, e1 := syscall.Syscall(procCancelIoEx.Addr(), 2, uintptr(s), uintptr(unsafe.Pointer(o)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func ClearCommBreak(handle Handle) (err error) {
	r1, _, e1 := syscall.Syscall(procClearCommBreak.Addr(), 1, uintptr(handle), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func ClearCommError(handle Handle, lpErrors *uint32, lpStat *ComStat) (err error) {
	r1, _, e1 := syscall.Syscall(procClearCommError.Addr(), 3, uintptr(handle), uintptr(unsafe.Pointer(lpErrors)), uintptr(unsafe.Pointer(lpStat)))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func CloseHandle(handle Handle) (err error) {
	r1, _, e1 := syscall.Syscall(procCloseHandle.Addr(), 1, uintptr(handle), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func ClosePseudoConsole(console Handle) {
	syscall.Syscall(procClosePseudoConsole.Addr(), 1, uintptr(console), 0, 0)
	return
}

func ConnectNamedPipe(pipe Handle, overlapped *Overlapped) (err error) {
	r1, _, e1 := syscall.Syscall(procConnectNamedPipe.Addr(), 2, uintptr(pipe), uintptr(unsafe.Pointer(overlapped)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func CreateDirectory(path *uint16, sa *SecurityAttributes) (err error) {
	r1, _, e1 := syscall.Syscall(procCreateDirectoryW.Addr(), 2, uintptr(unsafe.Pointer(path)), uintptr(unsafe.Pointer(sa)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func CreateEventEx(eventAttrs *SecurityAttributes, name *uint16, flags uint32, desiredAccess uint32) (handle Handle, err error) {
	r0, _, e1 := syscall.Syscall6(procCreateEventExW.Addr(), 4, uintptr(unsafe.Pointer(eventAttrs)), uintptr(unsafe.Pointer(name)), uintptr(flags), uintptr(desiredAccess), 0, 0)
	handle = Handle(r0)
	if handle == 0 || e1 == ERROR_ALREADY_EXISTS {
		err = errnoErr(e1)
	}
	return
}

func CreateEvent(eventAttrs *SecurityAttributes, manualReset uint32, initialState uint32, name *uint16) (handle Handle, err error) {
	r0, _, e1 := syscall.Syscall6(procCreateEventW.Addr(), 4, uintptr(unsafe.Pointer(eventAttrs)), uintptr(manualReset), uintptr(initialState), uintptr(unsafe.Pointer(name)), 0, 0)
	handle = Handle(r0)
	if handle == 0 || e1 == ERROR_ALREADY_EXISTS {
		err = errnoErr(e1)
	}
	return
}

func CreateFileMapping(fhandle Handle, sa *SecurityAttributes, prot uint32, maxSizeHigh uint32, maxSizeLow uint32, name *uint16) (handle Handle, err error) {
	r0, _, e1 := syscall.Syscall6(procCreateFileMappingW.Addr(), 6, uintptr(fhandle), uintptr(unsafe.Pointer(sa)), uintptr(prot), uintptr(maxSizeHigh), uintptr(maxSizeLow), uintptr(unsafe.Pointer(name)))
	handle = Handle(r0)
	if handle == 0 || e1 == ERROR_ALREADY_EXISTS {
		err = errnoErr(e1)
	}
	return
}

func CreateFile(name *uint16, access uint32, mode uint32, sa *SecurityAttributes, createmode uint32, attrs uint32, templatefile Handle) (handle Handle, err error) {
	r0, _, e1 := syscall.Syscall9(procCreateFileW.Addr(), 7, uintptr(unsafe.Pointer(name)), uintptr(access), uintptr(mode), uintptr(unsafe.Pointer(sa)), uintptr(createmode), uintptr(attrs), uintptr(templatefile), 0, 0)
	handle = Handle(r0)
	if handle == InvalidHandle {
		err = errnoErr(e1)
	}
	return
}

func CreateHardLink(filename *uint16, existingfilename *uint16, reserved uintptr) (err error) {
	r1, _, e1 := syscall.Syscall(procCreateHardLinkW.Addr(), 3, uintptr(unsafe.Pointer(filename)), uintptr(unsafe.Pointer(existingfilename)), uintptr(reserved))
	if r1&0xff == 0 {
		err = errnoErr(e1)
	}
	return
}

func CreateIoCompletionPort(filehandle Handle, cphandle Handle, key uintptr, threadcnt uint32) (handle Handle, err error) {
	r0, _, e1 := syscall.Syscall6(procCreateIoCompletionPort.Addr(), 4, uintptr(filehandle), uintptr(cphandle), uintptr(key), uintptr(threadcnt), 0, 0)
	handle = Handle(r0)
	if handle == 0 {
		err = errnoErr(e1)
	}
	return
}

func CreateJobObject(jobAttr *SecurityAttributes, name *uint16) (handle Handle, err error) {
	r0, _, e1 := syscall.Syscall(procCreateJobObjectW.Addr(), 2, uintptr(unsafe.Pointer(jobAttr)), uintptr(unsafe.Pointer(name)), 0)
	handle = Handle(r0)
	if handle == 0 {
		err = errnoErr(e1)
	}
	return
}

func CreateMutexEx(mutexAttrs *SecurityAttributes, name *uint16, flags uint32, desiredAccess uint32) (handle Handle, err error) {
	r0, _, e1 := syscall.Syscall6(procCreateMutexExW.Addr(), 4, uintptr(unsafe.Pointer(mutexAttrs)), uintptr(unsafe.Pointer(name)), uintptr(flags), uintptr(desiredAccess), 0, 0)
	handle = Handle(r0)
	if handle == 0 || e1 == ERROR_ALREADY_EXISTS {
		err = errnoErr(e1)
	}
	return
}

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

func CreateNamedPipe(name *uint16, flags uint32, pipeMode uint32, maxInstances uint32, outSize uint32, inSize uint32, defaultTimeout uint32, sa *SecurityAttributes) (handle Handle, err error) {
	r0, _, e1 := syscall.Syscall9(procCreateNamedPipeW.Addr(), 8, uintptr(unsafe.Pointer(name)), uintptr(flags), uintptr(pipeMode), uintptr(maxInstances), uintptr(outSize), uintptr(inSize), uintptr(defaultTimeout), uintptr(unsafe.Pointer(sa)), 0)
	handle = Handle(r0)
	if handle == InvalidHandle {
		err = errnoErr(e1)
	}
	return
}

func CreatePipe(readhandle *Handle, writehandle *Handle, sa *SecurityAttributes, size uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procCreatePipe.Addr(), 4, uintptr(unsafe.Pointer(readhandle)), uintptr(unsafe.Pointer(writehandle)), uintptr(unsafe.Pointer(sa)), uintptr(size), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

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

func CreateSymbolicLink(symlinkfilename *uint16, targetfilename *uint16, flags uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procCreateSymbolicLinkW.Addr(), 3, uintptr(unsafe.Pointer(symlinkfilename)), uintptr(unsafe.Pointer(targetfilename)), uintptr(flags))
	if r1&0xff == 0 {
		err = errnoErr(e1)
	}
	return
}

func CreateToolhelp32Snapshot(flags uint32, processId uint32) (handle Handle, err error) {
	r0, _, e1 := syscall.Syscall(procCreateToolhelp32Snapshot.Addr(), 2, uintptr(flags), uintptr(processId), 0)
	handle = Handle(r0)
	if handle == InvalidHandle {
		err = errnoErr(e1)
	}
	return
}

func DefineDosDevice(flags uint32, deviceName *uint16, targetPath *uint16) (err error) {
	r1, _, e1 := syscall.Syscall(procDefineDosDeviceW.Addr(), 3, uintptr(flags), uintptr(unsafe.Pointer(deviceName)), uintptr(unsafe.Pointer(targetPath)))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

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

func DeleteVolumeMountPoint(volumeMountPoint *uint16) (err error) {
	r1, _, e1 := syscall.Syscall(procDeleteVolumeMountPointW.Addr(), 1, uintptr(unsafe.Pointer(volumeMountPoint)), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func DeviceIoControl(handle Handle, ioControlCode uint32, inBuffer *byte, inBufferSize uint32, outBuffer *byte, outBufferSize uint32, bytesReturned *uint32, overlapped *Overlapped) (err error) {
	r1, _, e1 := syscall.Syscall9(procDeviceIoControl.Addr(), 8, uintptr(handle), uintptr(ioControlCode), uintptr(unsafe.Pointer(inBuffer)), uintptr(inBufferSize), uintptr(unsafe.Pointer(outBuffer)), uintptr(outBufferSize), uintptr(unsafe.Pointer(bytesReturned)), uintptr(unsafe.Pointer(overlapped)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func DisconnectNamedPipe(pipe Handle) (err error) {
	r1, _, e1 := syscall.Syscall(procDisconnectNamedPipe.Addr(), 1, uintptr(pipe), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

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

func EscapeCommFunction(handle Handle, dwFunc uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procEscapeCommFunction.Addr(), 2, uintptr(handle), uintptr(dwFunc), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func ExitProcess(exitcode uint32) {
	syscall.Syscall(procExitProcess.Addr(), 1, uintptr(exitcode), 0, 0)
	return
}

func ExpandEnvironmentStrings(src *uint16, dst *uint16, size uint32) (n uint32, err error) {
	r0, _, e1 := syscall.Syscall(procExpandEnvironmentStringsW.Addr(), 3, uintptr(unsafe.Pointer(src)), uintptr(unsafe.Pointer(dst)), uintptr(size))
	n = uint32(r0)
	if n == 0 {
		err = errnoErr(e1)
	}
	return
}

func FindClose(handle Handle) (err error) {
	r1, _, e1 := syscall.Syscall(procFindClose.Addr(), 1, uintptr(handle), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func FindCloseChangeNotification(handle Handle) (err error) {
	r1, _, e1 := syscall.Syscall(procFindCloseChangeNotification.Addr(), 1, uintptr(handle), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

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

func FindFirstVolumeMountPoint(rootPathName *uint16, volumeMountPoint *uint16, bufferLength uint32) (handle Handle, err error) {
	r0, _, e1 := syscall.Syscall(procFindFirstVolumeMountPointW.Addr(), 3, uintptr(unsafe.Pointer(rootPathName)), uintptr(unsafe.Pointer(volumeMountPoint)), uintptr(bufferLength))
	handle = Handle(r0)
	if handle == InvalidHandle {
		err = errnoErr(e1)
	}
	return
}

func FindFirstVolume(volumeName *uint16, bufferLength uint32) (handle Handle, err error) {
	r0, _, e1 := syscall.Syscall(procFindFirstVolumeW.Addr(), 2, uintptr(unsafe.Pointer(volumeName)), uintptr(bufferLength), 0)
	handle = Handle(r0)
	if handle == InvalidHandle {
		err = errnoErr(e1)
	}
	return
}

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

func FindNextVolumeMountPoint(findVolumeMountPoint Handle, volumeMountPoint *uint16, bufferLength uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procFindNextVolumeMountPointW.Addr(), 3, uintptr(findVolumeMountPoint), uintptr(unsafe.Pointer(volumeMountPoint)), uintptr(bufferLength))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

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

func FindVolumeClose(findVolume Handle) (err error) {
	r1, _, e1 := syscall.Syscall(procFindVolumeClose.Addr(), 1, uintptr(findVolume), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func FindVolumeMountPointClose(findVolumeMountPoint Handle) (err error) {
	r1, _, e1 := syscall.Syscall(procFindVolumeMountPointClose.Addr(), 1, uintptr(findVolumeMountPoint), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func FlushFileBuffers(handle Handle) (err error) {
	r1, _, e1 := syscall.Syscall(procFlushFileBuffers.Addr(), 1, uintptr(handle), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func FlushViewOfFile(addr uintptr, length uintptr) (err error) {
	r1, _, e1 := syscall.Syscall(procFlushViewOfFile.Addr(), 2, uintptr(addr), uintptr(length), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

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

func FreeEnvironmentStrings(envs *uint16) (err error) {
	r1, _, e1 := syscall.Syscall(procFreeEnvironmentStringsW.Addr(), 1, uintptr(unsafe.Pointer(envs)), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func FreeLibrary(handle Handle) (err error) {
	r1, _, e1 := syscall.Syscall(procFreeLibrary.Addr(), 1, uintptr(handle), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func GenerateConsoleCtrlEvent(ctrlEvent uint32, processGroupID uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procGenerateConsoleCtrlEvent.Addr(), 2, uintptr(ctrlEvent), uintptr(processGroupID), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func GetACP() (acp uint32) {
	r0, _, _ := syscall.Syscall(procGetACP.Addr(), 0, 0, 0, 0)
	acp = uint32(r0)
	return
}

func GetActiveProcessorCount(groupNumber uint16) (ret uint32) {
	r0, _, _ := syscall.Syscall(procGetActiveProcessorCount.Addr(), 1, uintptr(groupNumber), 0, 0)
	ret = uint32(r0)
	return
}

func GetCommModemStatus(handle Handle, lpModemStat *uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procGetCommModemStatus.Addr(), 2, uintptr(handle), uintptr(unsafe.Pointer(lpModemStat)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func GetCommState(handle Handle, lpDCB *DCB) (err error) {
	r1, _, e1 := syscall.Syscall(procGetCommState.Addr(), 2, uintptr(handle), uintptr(unsafe.Pointer(lpDCB)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func GetCommTimeouts(handle Handle, timeouts *CommTimeouts) (err error) {
	r1, _, e1 := syscall.Syscall(procGetCommTimeouts.Addr(), 2, uintptr(handle), uintptr(unsafe.Pointer(timeouts)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func GetCommandLine() (cmd *uint16) {
	r0, _, _ := syscall.Syscall(procGetCommandLineW.Addr(), 0, 0, 0, 0)
	cmd = (*uint16)(unsafe.Pointer(r0))
	return
}

func GetComputerNameEx(nametype uint32, buf *uint16, n *uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procGetComputerNameExW.Addr(), 3, uintptr(nametype), uintptr(unsafe.Pointer(buf)), uintptr(unsafe.Pointer(n)))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func GetComputerName(buf *uint16, n *uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procGetComputerNameW.Addr(), 2, uintptr(unsafe.Pointer(buf)), uintptr(unsafe.Pointer(n)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func GetConsoleMode(console Handle, mode *uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procGetConsoleMode.Addr(), 2, uintptr(console), uintptr(unsafe.Pointer(mode)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func GetConsoleScreenBufferInfo(console Handle, info *ConsoleScreenBufferInfo) (err error) {
	r1, _, e1 := syscall.Syscall(procGetConsoleScreenBufferInfo.Addr(), 2, uintptr(console), uintptr(unsafe.Pointer(info)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func GetCurrentDirectory(buflen uint32, buf *uint16) (n uint32, err error) {
	r0, _, e1 := syscall.Syscall(procGetCurrentDirectoryW.Addr(), 2, uintptr(buflen), uintptr(unsafe.Pointer(buf)), 0)
	n = uint32(r0)
	if n == 0 {
		err = errnoErr(e1)
	}
	return
}

func GetCurrentProcessId() (pid uint32) {
	r0, _, _ := syscall.Syscall(procGetCurrentProcessId.Addr(), 0, 0, 0, 0)
	pid = uint32(r0)
	return
}

func GetCurrentThreadId() (id uint32) {
	r0, _, _ := syscall.Syscall(procGetCurrentThreadId.Addr(), 0, 0, 0, 0)
	id = uint32(r0)
	return
}

func GetDiskFreeSpaceEx(directoryName *uint16, freeBytesAvailableToCaller *uint64, totalNumberOfBytes *uint64, totalNumberOfFreeBytes *uint64) (err error) {
	r1, _, e1 := syscall.Syscall6(procGetDiskFreeSpaceExW.Addr(), 4, uintptr(unsafe.Pointer(directoryName)), uintptr(unsafe.Pointer(freeBytesAvailableToCaller)), uintptr(unsafe.Pointer(totalNumberOfBytes)), uintptr(unsafe.Pointer(totalNumberOfFreeBytes)), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func GetDriveType(rootPathName *uint16) (driveType uint32) {
	r0, _, _ := syscall.Syscall(procGetDriveTypeW.Addr(), 1, uintptr(unsafe.Pointer(rootPathName)), 0, 0)
	driveType = uint32(r0)
	return
}

func GetEnvironmentStrings() (envs *uint16, err error) {
	r0, _, e1 := syscall.Syscall(procGetEnvironmentStringsW.Addr(), 0, 0, 0, 0)
	envs = (*uint16)(unsafe.Pointer(r0))
	if envs == nil {
		err = errnoErr(e1)
	}
	return
}

func GetEnvironmentVariable(name *uint16, buffer *uint16, size uint32) (n uint32, err error) {
	r0, _, e1 := syscall.Syscall(procGetEnvironmentVariableW.Addr(), 3, uintptr(unsafe.Pointer(name)), uintptr(unsafe.Pointer(buffer)), uintptr(size))
	n = uint32(r0)
	if n == 0 {
		err = errnoErr(e1)
	}
	return
}

func GetExitCodeProcess(handle Handle, exitcode *uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procGetExitCodeProcess.Addr(), 2, uintptr(handle), uintptr(unsafe.Pointer(exitcode)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func GetFileAttributesEx(name *uint16, level uint32, info *byte) (err error) {
	r1, _, e1 := syscall.Syscall(procGetFileAttributesExW.Addr(), 3, uintptr(unsafe.Pointer(name)), uintptr(level), uintptr(unsafe.Pointer(info)))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func GetFileAttributes(name *uint16) (attrs uint32, err error) {
	r0, _, e1 := syscall.Syscall(procGetFileAttributesW.Addr(), 1, uintptr(unsafe.Pointer(name)), 0, 0)
	attrs = uint32(r0)
	if attrs == INVALID_FILE_ATTRIBUTES {
		err = errnoErr(e1)
	}
	return
}

func GetFileInformationByHandle(handle Handle, data *ByHandleFileInformation) (err error) {
	r1, _, e1 := syscall.Syscall(procGetFileInformationByHandle.Addr(), 2, uintptr(handle), uintptr(unsafe.Pointer(data)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func GetFileInformationByHandleEx(handle Handle, class uint32, outBuffer *byte, outBufferLen uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procGetFileInformationByHandleEx.Addr(), 4, uintptr(handle), uintptr(class), uintptr(unsafe.Pointer(outBuffer)), uintptr(outBufferLen), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func GetFileTime(handle Handle, ctime *Filetime, atime *Filetime, wtime *Filetime) (err error) {
	r1, _, e1 := syscall.Syscall6(procGetFileTime.Addr(), 4, uintptr(handle), uintptr(unsafe.Pointer(ctime)), uintptr(unsafe.Pointer(atime)), uintptr(unsafe.Pointer(wtime)), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func GetFileType(filehandle Handle) (n uint32, err error) {
	r0, _, e1 := syscall.Syscall(procGetFileType.Addr(), 1, uintptr(filehandle), 0, 0)
	n = uint32(r0)
	if n == 0 {
		err = errnoErr(e1)
	}
	return
}

func GetFinalPathNameByHandle(file Handle, filePath *uint16, filePathSize uint32, flags uint32) (n uint32, err error) {
	r0, _, e1 := syscall.Syscall6(procGetFinalPathNameByHandleW.Addr(), 4, uintptr(file), uintptr(unsafe.Pointer(filePath)), uintptr(filePathSize), uintptr(flags), 0, 0)
	n = uint32(r0)
	if n == 0 {
		err = errnoErr(e1)
	}
	return
}

func GetFullPathName(path *uint16, buflen uint32, buf *uint16, fname **uint16) (n uint32, err error) {
	r0, _, e1 := syscall.Syscall6(procGetFullPathNameW.Addr(), 4, uintptr(unsafe.Pointer(path)), uintptr(buflen), uintptr(unsafe.Pointer(buf)), uintptr(unsafe.Pointer(fname)), 0, 0)
	n = uint32(r0)
	if n == 0 {
		err = errnoErr(e1)
	}
	return
}

func GetLargePageMinimum() (size uintptr) {
	r0, _, _ := syscall.Syscall(procGetLargePageMinimum.Addr(), 0, 0, 0, 0)
	size = uintptr(r0)
	return
}

func GetLastError() (lasterr error) {
	r0, _, _ := syscall.Syscall(procGetLastError.Addr(), 0, 0, 0, 0)
	if r0 != 0 {
		lasterr = syscall.Errno(r0)
	}
	return
}

func GetLogicalDriveStrings(bufferLength uint32, buffer *uint16) (n uint32, err error) {
	r0, _, e1 := syscall.Syscall(procGetLogicalDriveStringsW.Addr(), 2, uintptr(bufferLength), uintptr(unsafe.Pointer(buffer)), 0)
	n = uint32(r0)
	if n == 0 {
		err = errnoErr(e1)
	}
	return
}

func GetLogicalDrives() (drivesBitMask uint32, err error) {
	r0, _, e1 := syscall.Syscall(procGetLogicalDrives.Addr(), 0, 0, 0, 0)
	drivesBitMask = uint32(r0)
	if drivesBitMask == 0 {
		err = errnoErr(e1)
	}
	return
}

func GetLongPathName(path *uint16, buf *uint16, buflen uint32) (n uint32, err error) {
	r0, _, e1 := syscall.Syscall(procGetLongPathNameW.Addr(), 3, uintptr(unsafe.Pointer(path)), uintptr(unsafe.Pointer(buf)), uintptr(buflen))
	n = uint32(r0)
	if n == 0 {
		err = errnoErr(e1)
	}
	return
}

func GetMaximumProcessorCount(groupNumber uint16) (ret uint32) {
	r0, _, _ := syscall.Syscall(procGetMaximumProcessorCount.Addr(), 1, uintptr(groupNumber), 0, 0)
	ret = uint32(r0)
	return
}

func GetModuleFileName(module Handle, filename *uint16, size uint32) (n uint32, err error) {
	r0, _, e1 := syscall.Syscall(procGetModuleFileNameW.Addr(), 3, uintptr(module), uintptr(unsafe.Pointer(filename)), uintptr(size))
	n = uint32(r0)
	if n == 0 {
		err = errnoErr(e1)
	}
	return
}

func GetModuleHandleEx(flags uint32, moduleName *uint16, module *Handle) (err error) {
	r1, _, e1 := syscall.Syscall(procGetModuleHandleExW.Addr(), 3, uintptr(flags), uintptr(unsafe.Pointer(moduleName)), uintptr(unsafe.Pointer(module)))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func GetNamedPipeHandleState(pipe Handle, state *uint32, curInstances *uint32, maxCollectionCount *uint32, collectDataTimeout *uint32, userName *uint16, maxUserNameSize uint32) (err error) {
	r1, _, e1 := syscall.Syscall9(procGetNamedPipeHandleStateW.Addr(), 7, uintptr(pipe), uintptr(unsafe.Pointer(state)), uintptr(unsafe.Pointer(curInstances)), uintptr(unsafe.Pointer(maxCollectionCount)), uintptr(unsafe.Pointer(collectDataTimeout)), uintptr(unsafe.Pointer(userName)), uintptr(maxUserNameSize), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func GetNamedPipeInfo(pipe Handle, flags *uint32, outSize *uint32, inSize *uint32, maxInstances *uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procGetNamedPipeInfo.Addr(), 5, uintptr(pipe), uintptr(unsafe.Pointer(flags)), uintptr(unsafe.Pointer(outSize)), uintptr(unsafe.Pointer(inSize)), uintptr(unsafe.Pointer(maxInstances)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

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

func GetPriorityClass(process Handle) (ret uint32, err error) {
	r0, _, e1 := syscall.Syscall(procGetPriorityClass.Addr(), 1, uintptr(process), 0, 0)
	ret = uint32(r0)
	if ret == 0 {
		err = errnoErr(e1)
	}
	return
}

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

func GetProcessShutdownParameters(level *uint32, flags *uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procGetProcessShutdownParameters.Addr(), 2, uintptr(unsafe.Pointer(level)), uintptr(unsafe.Pointer(flags)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func GetProcessTimes(handle Handle, creationTime *Filetime, exitTime *Filetime, kernelTime *Filetime, userTime *Filetime) (err error) {
	r1, _, e1 := syscall.Syscall6(procGetProcessTimes.Addr(), 5, uintptr(handle), uintptr(unsafe.Pointer(creationTime)), uintptr(unsafe.Pointer(exitTime)), uintptr(unsafe.Pointer(kernelTime)), uintptr(unsafe.Pointer(userTime)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func GetProcessWorkingSetSizeEx(hProcess Handle, lpMinimumWorkingSetSize *uintptr, lpMaximumWorkingSetSize *uintptr, flags *uint32) {
	syscall.Syscall6(procGetProcessWorkingSetSizeEx.Addr(), 4, uintptr(hProcess), uintptr(unsafe.Pointer(lpMinimumWorkingSetSize)), uintptr(unsafe.Pointer(lpMaximumWorkingSetSize)), uintptr(unsafe.Pointer(flags)), 0, 0)
	return
}

func GetQueuedCompletionStatus(cphandle Handle, qty *uint32, key *uintptr, overlapped **Overlapped, timeout uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procGetQueuedCompletionStatus.Addr(), 5, uintptr(cphandle), uintptr(unsafe.Pointer(qty)), uintptr(unsafe.Pointer(key)), uintptr(unsafe.Pointer(overlapped)), uintptr(timeout), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

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

func GetSystemTimeAsFileTime(time *Filetime) {
	syscall.Syscall(procGetSystemTimeAsFileTime.Addr(), 1, uintptr(unsafe.Pointer(time)), 0, 0)
	return
}

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

func GetVersion() (ver uint32, err error) {
	r0, _, e1 := syscall.Syscall(procGetVersion.Addr(), 0, 0, 0, 0)
	ver = uint32(r0)
	if ver == 0 {
		err = errnoErr(e1)
	}
	return
}

func GetVolumeInformationByHandle(file Handle, volumeNameBuffer *uint16, volumeNameSize uint32, volumeNameSerialNumber *uint32, maximumComponentLength *uint32, fileSystemFlags *uint32, fileSystemNameBuffer *uint16, fileSystemNameSize uint32) (err error) {
	r1, _, e1 := syscall.Syscall9(procGetVolumeInformationByHandleW.Addr(), 8, uintptr(file), uintptr(unsafe.Pointer(volumeNameBuffer)), uintptr(volumeNameSize), uintptr(unsafe.Pointer(volumeNameSerialNumber)), uintptr(unsafe.Pointer(maximumComponentLength)), uintptr(unsafe.Pointer(fileSystemFlags)), uintptr(unsafe.Pointer(fileSystemNameBuffer)), uintptr(fileSystemNameSize), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func GetVolumeInformation(rootPathName *uint16, volumeNameBuffer *uint16, volumeNameSize uint32, volumeNameSerialNumber *uint32, maximumComponentLength *uint32, fileSystemFlags *uint32, fileSystemNameBuffer *uint16, fileSystemNameSize uint32) (err error) {
	r1, _, e1 := syscall.Syscall9(procGetVolumeInformationW.Addr(), 8, uintptr(unsafe.Pointer(rootPathName)), uintptr(unsafe.Pointer(volumeNameBuffer)), uintptr(volumeNameSize), uintptr(unsafe.Pointer(volumeNameSerialNumber)), uintptr(unsafe.Pointer(maximumComponentLength)), uintptr(unsafe.Pointer(fileSystemFlags)), uintptr(unsafe.Pointer(fileSystemNameBuffer)), uintptr(fileSystemNameSize), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func GetVolumeNameForVolumeMountPoint(volumeMountPoint *uint16, volumeName *uint16, bufferlength uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procGetVolumeNameForVolumeMountPointW.Addr(), 3, uintptr(unsafe.Pointer(volumeMountPoint)), uintptr(unsafe.Pointer(volumeName)), uintptr(bufferlength))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func GetVolumePathName(fileName *uint16, volumePathName *uint16, bufferLength uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procGetVolumePathNameW.Addr(), 3, uintptr(unsafe.Pointer(fileName)), uintptr(unsafe.Pointer(volumePathName)), uintptr(bufferLength))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

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

func LoadResource(module Handle, resInfo Handle) (resData Handle, err error) {
	r0, _, e1 := syscall.Syscall(procLoadResource.Addr(), 2, uintptr(module), uintptr(resInfo), 0)
	resData = Handle(r0)
	if resData == 0 {
		err = errnoErr(e1)
	}
	return
}

func LocalAlloc(flags uint32, length uint32) (ptr uintptr, err error) {
	r0, _, e1 := syscall.Syscall(procLocalAlloc.Addr(), 2, uintptr(flags), uintptr(length), 0)
	ptr = uintptr(r0)
	if ptr == 0 {
		err = errnoErr(e1)
	}
	return
}

func LocalFree(hmem Handle) (handle Handle, err error) {
	r0, _, e1 := syscall.Syscall(procLocalFree.Addr(), 1, uintptr(hmem), 0, 0)
	handle = Handle(r0)
	if handle != 0 {
		err = errnoErr(e1)
	}
	return
}

func LockFileEx(file Handle, flags uint32, reserved uint32, bytesLow uint32, bytesHigh uint32, overlapped *Overlapped) (err error) {
	r1, _, e1 := syscall.Syscall6(procLockFileEx.Addr(), 6, uintptr(file), uintptr(flags), uintptr(reserved), uintptr(bytesLow), uintptr(bytesHigh), uintptr(unsafe.Pointer(overlapped)))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func LockResource(resData Handle) (addr uintptr, err error) {
	r0, _, e1 := syscall.Syscall(procLockResource.Addr(), 1, uintptr(resData), 0, 0)
	addr = uintptr(r0)
	if addr == 0 {
		err = errnoErr(e1)
	}
	return
}

func MapViewOfFile(handle Handle, access uint32, offsetHigh uint32, offsetLow uint32, length uintptr) (addr uintptr, err error) {
	r0, _, e1 := syscall.Syscall6(procMapViewOfFile.Addr(), 5, uintptr(handle), uintptr(access), uintptr(offsetHigh), uintptr(offsetLow), uintptr(length), 0)
	addr = uintptr(r0)
	if addr == 0 {
		err = errnoErr(e1)
	}
	return
}

func Module32First(snapshot Handle, moduleEntry *ModuleEntry32) (err error) {
	r1, _, e1 := syscall.Syscall(procModule32FirstW.Addr(), 2, uintptr(snapshot), uintptr(unsafe.Pointer(moduleEntry)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func Module32Next(snapshot Handle, moduleEntry *ModuleEntry32) (err error) {
	r1, _, e1 := syscall.Syscall(procModule32NextW.Addr(), 2, uintptr(snapshot), uintptr(unsafe.Pointer(moduleEntry)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func MoveFileEx(from *uint16, to *uint16, flags uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procMoveFileExW.Addr(), 3, uintptr(unsafe.Pointer(from)), uintptr(unsafe.Pointer(to)), uintptr(flags))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func MoveFile(from *uint16, to *uint16) (err error) {
	r1, _, e1 := syscall.Syscall(procMoveFileW.Addr(), 2, uintptr(unsafe.Pointer(from)), uintptr(unsafe.Pointer(to)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func MultiByteToWideChar(codePage uint32, dwFlags uint32, str *byte, nstr int32, wchar *uint16, nwchar int32) (nwrite int32, err error) {
	r0, _, e1 := syscall.Syscall6(procMultiByteToWideChar.Addr(), 6, uintptr(codePage), uintptr(dwFlags), uintptr(unsafe.Pointer(str)), uintptr(nstr), uintptr(unsafe.Pointer(wchar)), uintptr(nwchar))
	nwrite = int32(r0)
	if nwrite == 0 {
		err = errnoErr(e1)
	}
	return
}

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

func PostQueuedCompletionStatus(cphandle Handle, qty uint32, key uintptr, overlapped *Overlapped) (err error) {
	r1, _, e1 := syscall.Syscall6(procPostQueuedCompletionStatus.Addr(), 4, uintptr(cphandle), uintptr(qty), uintptr(key), uintptr(unsafe.Pointer(overlapped)), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func Process32First(snapshot Handle, procEntry *ProcessEntry32) (err error) {
	r1, _, e1 := syscall.Syscall(procProcess32FirstW.Addr(), 2, uintptr(snapshot), uintptr(unsafe.Pointer(procEntry)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func Process32Next(snapshot Handle, procEntry *ProcessEntry32) (err error) {
	r1, _, e1 := syscall.Syscall(procProcess32NextW.Addr(), 2, uintptr(snapshot), uintptr(unsafe.Pointer(procEntry)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func ProcessIdToSessionId(pid uint32, sessionid *uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procProcessIdToSessionId.Addr(), 2, uintptr(pid), uintptr(unsafe.Pointer(sessionid)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func PulseEvent(event Handle) (err error) {
	r1, _, e1 := syscall.Syscall(procPulseEvent.Addr(), 1, uintptr(event), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func PurgeComm(handle Handle, dwFlags uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procPurgeComm.Addr(), 2, uintptr(handle), uintptr(dwFlags), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func QueryDosDevice(deviceName *uint16, targetPath *uint16, max uint32) (n uint32, err error) {
	r0, _, e1 := syscall.Syscall(procQueryDosDeviceW.Addr(), 3, uintptr(unsafe.Pointer(deviceName)), uintptr(unsafe.Pointer(targetPath)), uintptr(max))
	n = uint32(r0)
	if n == 0 {
		err = errnoErr(e1)
	}
	return
}

func QueryFullProcessImageName(proc Handle, flags uint32, exeName *uint16, size *uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procQueryFullProcessImageNameW.Addr(), 4, uintptr(proc), uintptr(flags), uintptr(unsafe.Pointer(exeName)), uintptr(unsafe.Pointer(size)), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func QueryInformationJobObject(job Handle, JobObjectInformationClass int32, JobObjectInformation uintptr, JobObjectInformationLength uint32, retlen *uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procQueryInformationJobObject.Addr(), 5, uintptr(job), uintptr(JobObjectInformationClass), uintptr(JobObjectInformation), uintptr(JobObjectInformationLength), uintptr(unsafe.Pointer(retlen)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func ReadConsole(console Handle, buf *uint16, toread uint32, read *uint32, inputControl *byte) (err error) {
	r1, _, e1 := syscall.Syscall6(procReadConsoleW.Addr(), 5, uintptr(console), uintptr(unsafe.Pointer(buf)), uintptr(toread), uintptr(unsafe.Pointer(read)), uintptr(unsafe.Pointer(inputControl)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

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

func ReadProcessMemory(process Handle, baseAddress uintptr, buffer *byte, size uintptr, numberOfBytesRead *uintptr) (err error) {
	r1, _, e1 := syscall.Syscall6(procReadProcessMemory.Addr(), 5, uintptr(process), uintptr(baseAddress), uintptr(unsafe.Pointer(buffer)), uintptr(size), uintptr(unsafe.Pointer(numberOfBytesRead)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func ReleaseMutex(mutex Handle) (err error) {
	r1, _, e1 := syscall.Syscall(procReleaseMutex.Addr(), 1, uintptr(mutex), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func RemoveDirectory(path *uint16) (err error) {
	r1, _, e1 := syscall.Syscall(procRemoveDirectoryW.Addr(), 1, uintptr(unsafe.Pointer(path)), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func RemoveDllDirectory(cookie uintptr) (err error) {
	r1, _, e1 := syscall.Syscall(procRemoveDllDirectory.Addr(), 1, uintptr(cookie), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

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

func ResumeThread(thread Handle) (ret uint32, err error) {
	r0, _, e1 := syscall.Syscall(procResumeThread.Addr(), 1, uintptr(thread), 0, 0)
	ret = uint32(r0)
	if ret == 0xffffffff {
		err = errnoErr(e1)
	}
	return
}

func SetCommBreak(handle Handle) (err error) {
	r1, _, e1 := syscall.Syscall(procSetCommBreak.Addr(), 1, uintptr(handle), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func SetCommMask(handle Handle, dwEvtMask uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procSetCommMask.Addr(), 2, uintptr(handle), uintptr(dwEvtMask), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func SetCommState(handle Handle, lpDCB *DCB) (err error) {
	r1, _, e1 := syscall.Syscall(procSetCommState.Addr(), 2, uintptr(handle), uintptr(unsafe.Pointer(lpDCB)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

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

func SetConsoleMode(console Handle, mode uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procSetConsoleMode.Addr(), 2, uintptr(console), uintptr(mode), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func SetCurrentDirectory(path *uint16) (err error) {
	r1, _, e1 := syscall.Syscall(procSetCurrentDirectoryW.Addr(), 1, uintptr(unsafe.Pointer(path)), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func SetDefaultDllDirectories(directoryFlags uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procSetDefaultDllDirectories.Addr(), 1, uintptr(directoryFlags), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

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

func SetEndOfFile(handle Handle) (err error) {
	r1, _, e1 := syscall.Syscall(procSetEndOfFile.Addr(), 1, uintptr(handle), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func SetEnvironmentVariable(name *uint16, value *uint16) (err error) {
	r1, _, e1 := syscall.Syscall(procSetEnvironmentVariableW.Addr(), 2, uintptr(unsafe.Pointer(name)), uintptr(unsafe.Pointer(value)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func SetErrorMode(mode uint32) (ret uint32) {
	r0, _, _ := syscall.Syscall(procSetErrorMode.Addr(), 1, uintptr(mode), 0, 0)
	ret = uint32(r0)
	return
}

func SetEvent(event Handle) (err error) {
	r1, _, e1 := syscall.Syscall(procSetEvent.Addr(), 1, uintptr(event), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func SetFileAttributes(name *uint16, attrs uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procSetFileAttributesW.Addr(), 2, uintptr(unsafe.Pointer(name)), uintptr(attrs), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func SetFileCompletionNotificationModes(handle Handle, flags uint8) (err error) {
	r1, _, e1 := syscall.Syscall(procSetFileCompletionNotificationModes.Addr(), 2, uintptr(handle), uintptr(flags), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func SetFileInformationByHandle(handle Handle, class uint32, inBuffer *byte, inBufferLen uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procSetFileInformationByHandle.Addr(), 4, uintptr(handle), uintptr(class), uintptr(unsafe.Pointer(inBuffer)), uintptr(inBufferLen), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func SetFilePointer(handle Handle, lowoffset int32, highoffsetptr *int32, whence uint32) (newlowoffset uint32, err error) {
	r0, _, e1 := syscall.Syscall6(procSetFilePointer.Addr(), 4, uintptr(handle), uintptr(lowoffset), uintptr(unsafe.Pointer(highoffsetptr)), uintptr(whence), 0, 0)
	newlowoffset = uint32(r0)
	if newlowoffset == 0xffffffff {
		err = errnoErr(e1)
	}
	return
}

func SetFileTime(handle Handle, ctime *Filetime, atime *Filetime, wtime *Filetime) (err error) {
	r1, _, e1 := syscall.Syscall6(procSetFileTime.Addr(), 4, uintptr(handle), uintptr(unsafe.Pointer(ctime)), uintptr(unsafe.Pointer(atime)), uintptr(unsafe.Pointer(wtime)), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func SetFileValidData(handle Handle, validDataLength int64) (err error) {
	r1, _, e1 := syscall.Syscall(procSetFileValidData.Addr(), 2, uintptr(handle), uintptr(validDataLength), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func SetHandleInformation(handle Handle, mask uint32, flags uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procSetHandleInformation.Addr(), 3, uintptr(handle), uintptr(mask), uintptr(flags))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func SetInformationJobObject(job Handle, JobObjectInformationClass uint32, JobObjectInformation uintptr, JobObjectInformationLength uint32) (ret int, err error) {
	r0, _, e1 := syscall.Syscall6(procSetInformationJobObject.Addr(), 4, uintptr(job), uintptr(JobObjectInformationClass), uintptr(JobObjectInformation), uintptr(JobObjectInformationLength), 0, 0)
	ret = int(r0)
	if ret == 0 {
		err = errnoErr(e1)
	}
	return
}

func SetNamedPipeHandleState(pipe Handle, state *uint32, maxCollectionCount *uint32, collectDataTimeout *uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procSetNamedPipeHandleState.Addr(), 4, uintptr(pipe), uintptr(unsafe.Pointer(state)), uintptr(unsafe.Pointer(maxCollectionCount)), uintptr(unsafe.Pointer(collectDataTimeout)), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func SetPriorityClass(process Handle, priorityClass uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procSetPriorityClass.Addr(), 2, uintptr(process), uintptr(priorityClass), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

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

func SetProcessShutdownParameters(level uint32, flags uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procSetProcessShutdownParameters.Addr(), 2, uintptr(level), uintptr(flags), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func SetProcessWorkingSetSizeEx(hProcess Handle, dwMinimumWorkingSetSize uintptr, dwMaximumWorkingSetSize uintptr, flags uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procSetProcessWorkingSetSizeEx.Addr(), 4, uintptr(hProcess), uintptr(dwMinimumWorkingSetSize), uintptr(dwMaximumWorkingSetSize), uintptr(flags), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func SetStdHandle(stdhandle uint32, handle Handle) (err error) {
	r1, _, e1 := syscall.Syscall(procSetStdHandle.Addr(), 2, uintptr(stdhandle), uintptr(handle), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func SetVolumeLabel(rootPathName *uint16, volumeName *uint16) (err error) {
	r1, _, e1 := syscall.Syscall(procSetVolumeLabelW.Addr(), 2, uintptr(unsafe.Pointer(rootPathName)), uintptr(unsafe.Pointer(volumeName)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func SetVolumeMountPoint(volumeMountPoint *uint16, volumeName *uint16) (err error) {
	r1, _, e1 := syscall.Syscall(procSetVolumeMountPointW.Addr(), 2, uintptr(unsafe.Pointer(volumeMountPoint)), uintptr(unsafe.Pointer(volumeName)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func SetupComm(handle Handle, dwInQueue uint32, dwOutQueue uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procSetupComm.Addr(), 3, uintptr(handle), uintptr(dwInQueue), uintptr(dwOutQueue))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func SizeofResource(module Handle, resInfo Handle) (size uint32, err error) {
	r0, _, e1 := syscall.Syscall(procSizeofResource.Addr(), 2, uintptr(module), uintptr(resInfo), 0)
	size = uint32(r0)
	if size == 0 {
		err = errnoErr(e1)
	}
	return
}

func SleepEx(milliseconds uint32, alertable bool) (ret uint32) {
	var _p0 uint32
	if alertable {
		_p0 = 1
	}
	r0, _, _ := syscall.Syscall(procSleepEx.Addr(), 2, uintptr(milliseconds), uintptr(_p0), 0)
	ret = uint32(r0)
	return
}

func TerminateJobObject(job Handle, exitCode uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procTerminateJobObject.Addr(), 2, uintptr(job), uintptr(exitCode), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func TerminateProcess(handle Handle, exitcode uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procTerminateProcess.Addr(), 2, uintptr(handle), uintptr(exitcode), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func Thread32First(snapshot Handle, threadEntry *ThreadEntry32) (err error) {
	r1, _, e1 := syscall.Syscall(procThread32First.Addr(), 2, uintptr(snapshot), uintptr(unsafe.Pointer(threadEntry)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func Thread32Next(snapshot Handle, threadEntry *ThreadEntry32) (err error) {
	r1, _, e1 := syscall.Syscall(procThread32Next.Addr(), 2, uintptr(snapshot), uintptr(unsafe.Pointer(threadEntry)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func UnlockFileEx(file Handle, reserved uint32, bytesLow uint32, bytesHigh uint32, overlapped *Overlapped) (err error) {
	r1, _, e1 := syscall.Syscall6(procUnlockFileEx.Addr(), 5, uintptr(file), uintptr(reserved), uintptr(bytesLow), uintptr(bytesHigh), uintptr(unsafe.Pointer(overlapped)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

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

func VirtualAlloc(address uintptr, size uintptr, alloctype uint32, protect uint32) (value uintptr, err error) {
	r0, _, e1 := syscall.Syscall6(procVirtualAlloc.Addr(), 4, uintptr(address), uintptr(size), uintptr(alloctype), uintptr(protect), 0, 0)
	value = uintptr(r0)
	if value == 0 {
		err = errnoErr(e1)
	}
	return
}

func VirtualFree(address uintptr, size uintptr, freetype uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procVirtualFree.Addr(), 3, uintptr(address), uintptr(size), uintptr(freetype))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func VirtualLock(addr uintptr, length uintptr) (err error) {
	r1, _, e1 := syscall.Syscall(procVirtualLock.Addr(), 2, uintptr(addr), uintptr(length), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func VirtualProtect(address uintptr, size uintptr, newprotect uint32, oldprotect *uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procVirtualProtect.Addr(), 4, uintptr(address), uintptr(size), uintptr(newprotect), uintptr(unsafe.Pointer(oldprotect)), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func VirtualProtectEx(process Handle, address uintptr, size uintptr, newProtect uint32, oldProtect *uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procVirtualProtectEx.Addr(), 5, uintptr(process), uintptr(address), uintptr(size), uintptr(newProtect), uintptr(unsafe.Pointer(oldProtect)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func VirtualQuery(address uintptr, buffer *MemoryBasicInformation, length uintptr) (err error) {
	r1, _, e1 := syscall.Syscall(procVirtualQuery.Addr(), 3, uintptr(address), uintptr(unsafe.Pointer(buffer)), uintptr(length))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func VirtualQueryEx(process Handle, address uintptr, buffer *MemoryBasicInformation, length uintptr) (err error) {
	r1, _, e1 := syscall.Syscall6(procVirtualQueryEx.Addr(), 4, uintptr(process), uintptr(address), uintptr(unsafe.Pointer(buffer)), uintptr(length), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func VirtualUnlock(addr uintptr, length uintptr) (err error) {
	r1, _, e1 := syscall.Syscall(procVirtualUnlock.Addr(), 2, uintptr(addr), uintptr(length), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func WTSGetActiveConsoleSessionId() (sessionID uint32) {
	r0, _, _ := syscall.Syscall(procWTSGetActiveConsoleSessionId.Addr(), 0, 0, 0, 0)
	sessionID = uint32(r0)
	return
}

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

func WaitForSingleObject(handle Handle, waitMilliseconds uint32) (event uint32, err error) {
	r0, _, e1 := syscall.Syscall(procWaitForSingleObject.Addr(), 2, uintptr(handle), uintptr(waitMilliseconds), 0)
	event = uint32(r0)
	if event == 0xffffffff {
		err = errnoErr(e1)
	}
	return
}

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

func WriteProcessMemory(process Handle, baseAddress uintptr, buffer *byte, size uintptr, numberOfBytesWritten *uintptr) (err error) {
	r1, _, e1 := syscall.Syscall6(procWriteProcessMemory.Addr(), 5, uintptr(process), uintptr(baseAddress), uintptr(unsafe.Pointer(buffer)), uintptr(size), uintptr(unsafe.Pointer(numberOfBytesWritten)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func AcceptEx(ls Handle, as Handle, buf *byte, rxdatalen uint32, laddrlen uint32, raddrlen uint32, recvd *uint32, overlapped *Overlapped) (err error) {
	r1, _, e1 := syscall.Syscall9(procAcceptEx.Addr(), 8, uintptr(ls), uintptr(as), uintptr(unsafe.Pointer(buf)), uintptr(rxdatalen), uintptr(laddrlen), uintptr(raddrlen), uintptr(unsafe.Pointer(recvd)), uintptr(unsafe.Pointer(overlapped)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func GetAcceptExSockaddrs(buf *byte, rxdatalen uint32, laddrlen uint32, raddrlen uint32, lrsa **RawSockaddrAny, lrsalen *int32, rrsa **RawSockaddrAny, rrsalen *int32) {
	syscall.Syscall9(procGetAcceptExSockaddrs.Addr(), 8, uintptr(unsafe.Pointer(buf)), uintptr(rxdatalen), uintptr(laddrlen), uintptr(raddrlen), uintptr(unsafe.Pointer(lrsa)), uintptr(unsafe.Pointer(lrsalen)), uintptr(unsafe.Pointer(rrsa)), uintptr(unsafe.Pointer(rrsalen)), 0)
	return
}

func TransmitFile(s Handle, handle Handle, bytesToWrite uint32, bytsPerSend uint32, overlapped *Overlapped, transmitFileBuf *TransmitFileBuffers, flags uint32) (err error) {
	r1, _, e1 := syscall.Syscall9(procTransmitFile.Addr(), 7, uintptr(s), uintptr(handle), uintptr(bytesToWrite), uintptr(bytsPerSend), uintptr(unsafe.Pointer(overlapped)), uintptr(unsafe.Pointer(transmitFileBuf)), uintptr(flags), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func NetApiBufferFree(buf *byte) (neterr error) {
	r0, _, _ := syscall.Syscall(procNetApiBufferFree.Addr(), 1, uintptr(unsafe.Pointer(buf)), 0, 0)
	if r0 != 0 {
		neterr = syscall.Errno(r0)
	}
	return
}

func NetGetJoinInformation(server *uint16, name **uint16, bufType *uint32) (neterr error) {
	r0, _, _ := syscall.Syscall(procNetGetJoinInformation.Addr(), 3, uintptr(unsafe.Pointer(server)), uintptr(unsafe.Pointer(name)), uintptr(unsafe.Pointer(bufType)))
	if r0 != 0 {
		neterr = syscall.Errno(r0)
	}
	return
}

func NetUserGetInfo(serverName *uint16, userName *uint16, level uint32, buf **byte) (neterr error) {
	r0, _, _ := syscall.Syscall6(procNetUserGetInfo.Addr(), 4, uintptr(unsafe.Pointer(serverName)), uintptr(unsafe.Pointer(userName)), uintptr(level), uintptr(unsafe.Pointer(buf)), 0, 0)
	if r0 != 0 {
		neterr = syscall.Errno(r0)
	}
	return
}

func NtCreateFile(handle *Handle, access uint32, oa *OBJECT_ATTRIBUTES, iosb *IO_STATUS_BLOCK, allocationSize *int64, attributes uint32, share uint32, disposition uint32, options uint32, eabuffer uintptr, ealength uint32) (ntstatus error) {
	r0, _, _ := syscall.Syscall12(procNtCreateFile.Addr(), 11, uintptr(unsafe.Pointer(handle)), uintptr(access), uintptr(unsafe.Pointer(oa)), uintptr(unsafe.Pointer(iosb)), uintptr(unsafe.Pointer(allocationSize)), uintptr(attributes), uintptr(share), uintptr(disposition), uintptr(options), uintptr(eabuffer), uintptr(ealength), 0)
	if r0 != 0 {
		ntstatus = NTStatus(r0)
	}
	return
}

func NtCreateNamedPipeFile(pipe *Handle, access uint32, oa *OBJECT_ATTRIBUTES, iosb *IO_STATUS_BLOCK, share uint32, disposition uint32, options uint32, typ uint32, readMode uint32, completionMode uint32, maxInstances uint32, inboundQuota uint32, outputQuota uint32, timeout *int64) (ntstatus error) {
	r0, _, _ := syscall.Syscall15(procNtCreateNamedPipeFile.Addr(), 14, uintptr(unsafe.Pointer(pipe)), uintptr(access), uintptr(unsafe.Pointer(oa)), uintptr(unsafe.Pointer(iosb)), uintptr(share), uintptr(disposition), uintptr(options), uintptr(typ), uintptr(readMode), uintptr(completionMode), uintptr(maxInstances), uintptr(inboundQuota), uintptr(outputQuota), uintptr(unsafe.Pointer(timeout)), 0)
	if r0 != 0 {
		ntstatus = NTStatus(r0)
	}
	return
}

func NtQueryInformationProcess(proc Handle, procInfoClass int32, procInfo unsafe.Pointer, procInfoLen uint32, retLen *uint32) (ntstatus error) {
	r0, _, _ := syscall.Syscall6(procNtQueryInformationProcess.Addr(), 5, uintptr(proc), uintptr(procInfoClass), uintptr(procInfo), uintptr(procInfoLen), uintptr(unsafe.Pointer(retLen)), 0)
	if r0 != 0 {
		ntstatus = NTStatus(r0)
	}
	return
}

func NtQuerySystemInformation(sysInfoClass int32, sysInfo unsafe.Pointer, sysInfoLen uint32, retLen *uint32) (ntstatus error) {
	r0, _, _ := syscall.Syscall6(procNtQuerySystemInformation.Addr(), 4, uintptr(sysInfoClass), uintptr(sysInfo), uintptr(sysInfoLen), uintptr(unsafe.Pointer(retLen)), 0, 0)
	if r0 != 0 {
		ntstatus = NTStatus(r0)
	}
	return
}

func NtSetInformationFile(handle Handle, iosb *IO_STATUS_BLOCK, inBuffer *byte, inBufferLen uint32, class uint32) (ntstatus error) {
	r0, _, _ := syscall.Syscall6(procNtSetInformationFile.Addr(), 5, uintptr(handle), uintptr(unsafe.Pointer(iosb)), uintptr(unsafe.Pointer(inBuffer)), uintptr(inBufferLen), uintptr(class), 0)
	if r0 != 0 {
		ntstatus = NTStatus(r0)
	}
	return
}

func NtSetInformationProcess(proc Handle, procInfoClass int32, procInfo unsafe.Pointer, procInfoLen uint32) (ntstatus error) {
	r0, _, _ := syscall.Syscall6(procNtSetInformationProcess.Addr(), 4, uintptr(proc), uintptr(procInfoClass), uintptr(procInfo), uintptr(procInfoLen), 0, 0)
	if r0 != 0 {
		ntstatus = NTStatus(r0)
	}
	return
}

func NtSetSystemInformation(sysInfoClass int32, sysInfo unsafe.Pointer, sysInfoLen uint32) (ntstatus error) {
	r0, _, _ := syscall.Syscall(procNtSetSystemInformation.Addr(), 3, uintptr(sysInfoClass), uintptr(sysInfo), uintptr(sysInfoLen))
	if r0 != 0 {
		ntstatus = NTStatus(r0)
	}
	return
}

func RtlAddFunctionTable(functionTable *RUNTIME_FUNCTION, entryCount uint32, baseAddress uintptr) (ret bool) {
	r0, _, _ := syscall.Syscall(procRtlAddFunctionTable.Addr(), 3, uintptr(unsafe.Pointer(functionTable)), uintptr(entryCount), uintptr(baseAddress))
	ret = r0 != 0
	return
}

func RtlDefaultNpAcl(acl **ACL) (ntstatus error) {
	r0, _, _ := syscall.Syscall(procRtlDefaultNpAcl.Addr(), 1, uintptr(unsafe.Pointer(acl)), 0, 0)
	if r0 != 0 {
		ntstatus = NTStatus(r0)
	}
	return
}

func RtlDeleteFunctionTable(functionTable *RUNTIME_FUNCTION) (ret bool) {
	r0, _, _ := syscall.Syscall(procRtlDeleteFunctionTable.Addr(), 1, uintptr(unsafe.Pointer(functionTable)), 0, 0)
	ret = r0 != 0
	return
}

func RtlDosPathNameToNtPathName(dosName *uint16, ntName *NTUnicodeString, ntFileNamePart *uint16, relativeName *RTL_RELATIVE_NAME) (ntstatus error) {
	r0, _, _ := syscall.Syscall6(procRtlDosPathNameToNtPathName_U_WithStatus.Addr(), 4, uintptr(unsafe.Pointer(dosName)), uintptr(unsafe.Pointer(ntName)), uintptr(unsafe.Pointer(ntFileNamePart)), uintptr(unsafe.Pointer(relativeName)), 0, 0)
	if r0 != 0 {
		ntstatus = NTStatus(r0)
	}
	return
}

func RtlDosPathNameToRelativeNtPathName(dosName *uint16, ntName *NTUnicodeString, ntFileNamePart *uint16, relativeName *RTL_RELATIVE_NAME) (ntstatus error) {
	r0, _, _ := syscall.Syscall6(procRtlDosPathNameToRelativeNtPathName_U_WithStatus.Addr(), 4, uintptr(unsafe.Pointer(dosName)), uintptr(unsafe.Pointer(ntName)), uintptr(unsafe.Pointer(ntFileNamePart)), uintptr(unsafe.Pointer(relativeName)), 0, 0)
	if r0 != 0 {
		ntstatus = NTStatus(r0)
	}
	return
}

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

func RtlInitString(destinationString *NTString, sourceString *byte) {
	syscall.Syscall(procRtlInitString.Addr(), 2, uintptr(unsafe.Pointer(destinationString)), uintptr(unsafe.Pointer(sourceString)), 0)
	return
}

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

func CoGetObject(name *uint16, bindOpts *BIND_OPTS3, guid *GUID, functionTable **uintptr) (ret error) {
	r0, _, _ := syscall.Syscall6(procCoGetObject.Addr(), 4, uintptr(unsafe.Pointer(name)), uintptr(unsafe.Pointer(bindOpts)), uintptr(unsafe.Pointer(guid)), uintptr(unsafe.Pointer(functionTable)), 0, 0)
	if r0 != 0 {
		ret = syscall.Errno(r0)
	}
	return
}

func CoInitializeEx(reserved uintptr, coInit uint32) (ret error) {
	r0, _, _ := syscall.Syscall(procCoInitializeEx.Addr(), 2, uintptr(reserved), uintptr(coInit), 0)
	if r0 != 0 {
		ret = syscall.Errno(r0)
	}
	return
}

func CoTaskMemFree(address unsafe.Pointer) {
	syscall.Syscall(procCoTaskMemFree.Addr(), 1, uintptr(address), 0, 0)
	return
}

func CoUninitialize() {
	syscall.Syscall(procCoUninitialize.Addr(), 0, 0, 0, 0)
	return
}

func stringFromGUID2(rguid *GUID, lpsz *uint16, cchMax int32) (chars int32) {
	r0, _, _ := syscall.Syscall(procStringFromGUID2.Addr(), 3, uintptr(unsafe.Pointer(rguid)), uintptr(unsafe.Pointer(lpsz)), uintptr(cchMax))
	chars = int32(r0)
	return
}

func EnumProcessModules(process Handle, module *Handle, cb uint32, cbNeeded *uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procEnumProcessModules.Addr(), 4, uintptr(process), uintptr(unsafe.Pointer(module)), uintptr(cb), uintptr(unsafe.Pointer(cbNeeded)), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

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

func GetModuleBaseName(process Handle, module Handle, baseName *uint16, size uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procGetModuleBaseNameW.Addr(), 4, uintptr(process), uintptr(module), uintptr(unsafe.Pointer(baseName)), uintptr(size), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func GetModuleFileNameEx(process Handle, module Handle, filename *uint16, size uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procGetModuleFileNameExW.Addr(), 4, uintptr(process), uintptr(module), uintptr(unsafe.Pointer(filename)), uintptr(size), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func GetModuleInformation(process Handle, module Handle, modinfo *ModuleInfo, cb uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procGetModuleInformation.Addr(), 4, uintptr(process), uintptr(module), uintptr(unsafe.Pointer(modinfo)), uintptr(cb), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func QueryWorkingSetEx(process Handle, pv uintptr, cb uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procQueryWorkingSetEx.Addr(), 3, uintptr(process), uintptr(pv), uintptr(cb))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

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

func UnsubscribeServiceChangeNotifications(subscription uintptr) (err error) {
	err = procUnsubscribeServiceChangeNotifications.Find()
	if err != nil {
		return
	}
	syscall.Syscall(procUnsubscribeServiceChangeNotifications.Addr(), 1, uintptr(subscription), 0, 0)
	return
}

func GetUserNameEx(nameFormat uint32, nameBuffre *uint16, nSize *uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procGetUserNameExW.Addr(), 3, uintptr(nameFormat), uintptr(unsafe.Pointer(nameBuffre)), uintptr(unsafe.Pointer(nSize)))
	if r1&0xff == 0 {
		err = errnoErr(e1)
	}
	return
}

func TranslateName(accName *uint16, accNameFormat uint32, desiredNameFormat uint32, translatedName *uint16, nSize *uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procTranslateNameW.Addr(), 5, uintptr(unsafe.Pointer(accName)), uintptr(accNameFormat), uintptr(desiredNameFormat), uintptr(unsafe.Pointer(translatedName)), uintptr(unsafe.Pointer(nSize)), 0)
	if r1&0xff == 0 {
		err = errnoErr(e1)
	}
	return
}

func SetupDiBuildDriverInfoList(deviceInfoSet DevInfo, deviceInfoData *DevInfoData, driverType SPDIT) (err error) {
	r1, _, e1 := syscall.Syscall(procSetupDiBuildDriverInfoList.Addr(), 3, uintptr(deviceInfoSet), uintptr(unsafe.Pointer(deviceInfoData)), uintptr(driverType))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func SetupDiCallClassInstaller(installFunction DI_FUNCTION, deviceInfoSet DevInfo, deviceInfoData *DevInfoData) (err error) {
	r1, _, e1 := syscall.Syscall(procSetupDiCallClassInstaller.Addr(), 3, uintptr(installFunction), uintptr(deviceInfoSet), uintptr(unsafe.Pointer(deviceInfoData)))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

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

func SetupDiDestroyDeviceInfoList(deviceInfoSet DevInfo) (err error) {
	r1, _, e1 := syscall.Syscall(procSetupDiDestroyDeviceInfoList.Addr(), 1, uintptr(deviceInfoSet), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

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

func SetupDiOpenDevRegKey(deviceInfoSet DevInfo, deviceInfoData *DevInfoData, Scope DICS_FLAG, HwProfile uint32, KeyType DIREG, samDesired uint32) (key Handle, err error) {
	r0, _, e1 := syscall.Syscall6(procSetupDiOpenDevRegKey.Addr(), 6, uintptr(deviceInfoSet), uintptr(unsafe.Pointer(deviceInfoData)), uintptr(Scope), uintptr(HwProfile), uintptr(KeyType), uintptr(samDesired))
	key = Handle(r0)
	if key == InvalidHandle {
		err = errnoErr(e1)
	}
	return
}

func SetupDiSetClassInstallParams(deviceInfoSet DevInfo, deviceInfoData *DevInfoData, classInstallParams *ClassInstallHeader, classInstallParamsSize uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procSetupDiSetClassInstallParamsW.Addr(), 4, uintptr(deviceInfoSet), uintptr(unsafe.Pointer(deviceInfoData)), uintptr(unsafe.Pointer(classInstallParams)), uintptr(classInstallParamsSize), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

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

func SetupDiSetSelectedDevice(deviceInfoSet DevInfo, deviceInfoData *DevInfoData) (err error) {
	r1, _, e1 := syscall.Syscall(procSetupDiSetSelectedDevice.Addr(), 2, uintptr(deviceInfoSet), uintptr(unsafe.Pointer(deviceInfoData)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

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

func ShellExecute(hwnd Handle, verb *uint16, file *uint16, args *uint16, cwd *uint16, showCmd int32) (err error) {
	r1, _, e1 := syscall.Syscall6(procShellExecuteW.Addr(), 6, uintptr(hwnd), uintptr(unsafe.Pointer(verb)), uintptr(unsafe.Pointer(file)), uintptr(unsafe.Pointer(args)), uintptr(unsafe.Pointer(cwd)), uintptr(showCmd))
	if r1 <= 32 {
		err = errnoErr(e1)
	}
	return
}

func EnumChildWindows(hwnd HWND, enumFunc uintptr, param unsafe.Pointer) {
	syscall.Syscall(procEnumChildWindows.Addr(), 3, uintptr(hwnd), uintptr(enumFunc), uintptr(param))
	return
}

func EnumWindows(enumFunc uintptr, param unsafe.Pointer) (err error) {
	r1, _, e1 := syscall.Syscall(procEnumWindows.Addr(), 2, uintptr(enumFunc), uintptr(param), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func ExitWindowsEx(flags uint32, reason uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procExitWindowsEx.Addr(), 2, uintptr(flags), uintptr(reason), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func GetClassName(hwnd HWND, className *uint16, maxCount int32) (copied int32, err error) {
	r0, _, e1 := syscall.Syscall(procGetClassNameW.Addr(), 3, uintptr(hwnd), uintptr(unsafe.Pointer(className)), uintptr(maxCount))
	copied = int32(r0)
	if copied == 0 {
		err = errnoErr(e1)
	}
	return
}

func GetDesktopWindow() (hwnd HWND) {
	r0, _, _ := syscall.Syscall(procGetDesktopWindow.Addr(), 0, 0, 0, 0)
	hwnd = HWND(r0)
	return
}

func GetForegroundWindow() (hwnd HWND) {
	r0, _, _ := syscall.Syscall(procGetForegroundWindow.Addr(), 0, 0, 0, 0)
	hwnd = HWND(r0)
	return
}

func GetGUIThreadInfo(thread uint32, info *GUIThreadInfo) (err error) {
	r1, _, e1 := syscall.Syscall(procGetGUIThreadInfo.Addr(), 2, uintptr(thread), uintptr(unsafe.Pointer(info)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func GetShellWindow() (shellWindow HWND) {
	r0, _, _ := syscall.Syscall(procGetShellWindow.Addr(), 0, 0, 0, 0)
	shellWindow = HWND(r0)
	return
}

func GetWindowThreadProcessId(hwnd HWND, pid *uint32) (tid uint32, err error) {
	r0, _, e1 := syscall.Syscall(procGetWindowThreadProcessId.Addr(), 2, uintptr(hwnd), uintptr(unsafe.Pointer(pid)), 0)
	tid = uint32(r0)
	if tid == 0 {
		err = errnoErr(e1)
	}
	return
}

func IsWindow(hwnd HWND) (isWindow bool) {
	r0, _, _ := syscall.Syscall(procIsWindow.Addr(), 1, uintptr(hwnd), 0, 0)
	isWindow = r0 != 0
	return
}

func IsWindowUnicode(hwnd HWND) (isUnicode bool) {
	r0, _, _ := syscall.Syscall(procIsWindowUnicode.Addr(), 1, uintptr(hwnd), 0, 0)
	isUnicode = r0 != 0
	return
}

func IsWindowVisible(hwnd HWND) (isVisible bool) {
	r0, _, _ := syscall.Syscall(procIsWindowVisible.Addr(), 1, uintptr(hwnd), 0, 0)
	isVisible = r0 != 0
	return
}

func MessageBox(hwnd HWND, text *uint16, caption *uint16, boxtype uint32) (ret int32, err error) {
	r0, _, e1 := syscall.Syscall6(procMessageBoxW.Addr(), 4, uintptr(hwnd), uintptr(unsafe.Pointer(text)), uintptr(unsafe.Pointer(caption)), uintptr(boxtype), 0, 0)
	ret = int32(r0)
	if ret == 0 {
		err = errnoErr(e1)
	}
	return
}

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

func DestroyEnvironmentBlock(block *uint16) (err error) {
	r1, _, e1 := syscall.Syscall(procDestroyEnvironmentBlock.Addr(), 1, uintptr(unsafe.Pointer(block)), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func GetUserProfileDirectory(t Token, dir *uint16, dirLen *uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procGetUserProfileDirectoryW.Addr(), 3, uintptr(t), uintptr(unsafe.Pointer(dir)), uintptr(unsafe.Pointer(dirLen)))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

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

func TimeBeginPeriod(period uint32) (err error) {
	r1, _, e1 := syscall.Syscall(proctimeBeginPeriod.Addr(), 1, uintptr(period), 0, 0)
	if r1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func TimeEndPeriod(period uint32) (err error) {
	r1, _, e1 := syscall.Syscall(proctimeEndPeriod.Addr(), 1, uintptr(period), 0, 0)
	if r1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func WinVerifyTrustEx(hwnd HWND, actionId *GUID, data *WinTrustData) (ret error) {
	r0, _, _ := syscall.Syscall(procWinVerifyTrustEx.Addr(), 3, uintptr(hwnd), uintptr(unsafe.Pointer(actionId)), uintptr(unsafe.Pointer(data)))
	if r0 != 0 {
		ret = syscall.Errno(r0)
	}
	return
}

func FreeAddrInfoW(addrinfo *AddrinfoW) {
	syscall.Syscall(procFreeAddrInfoW.Addr(), 1, uintptr(unsafe.Pointer(addrinfo)), 0, 0)
	return
}

func GetAddrInfoW(nodename *uint16, servicename *uint16, hints *AddrinfoW, result **AddrinfoW) (sockerr error) {
	r0, _, _ := syscall.Syscall6(procGetAddrInfoW.Addr(), 4, uintptr(unsafe.Pointer(nodename)), uintptr(unsafe.Pointer(servicename)), uintptr(unsafe.Pointer(hints)), uintptr(unsafe.Pointer(result)), 0, 0)
	if r0 != 0 {
		sockerr = syscall.Errno(r0)
	}
	return
}

func WSACleanup() (err error) {
	r1, _, e1 := syscall.Syscall(procWSACleanup.Addr(), 0, 0, 0, 0)
	if r1 == socket_error {
		err = errnoErr(e1)
	}
	return
}

func WSAEnumProtocols(protocols *int32, protocolBuffer *WSAProtocolInfo, bufferLength *uint32) (n int32, err error) {
	r0, _, e1 := syscall.Syscall(procWSAEnumProtocolsW.Addr(), 3, uintptr(unsafe.Pointer(protocols)), uintptr(unsafe.Pointer(protocolBuffer)), uintptr(unsafe.Pointer(bufferLength)))
	n = int32(r0)
	if n == -1 {
		err = errnoErr(e1)
	}
	return
}

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

func WSAIoctl(s Handle, iocc uint32, inbuf *byte, cbif uint32, outbuf *byte, cbob uint32, cbbr *uint32, overlapped *Overlapped, completionRoutine uintptr) (err error) {
	r1, _, e1 := syscall.Syscall9(procWSAIoctl.Addr(), 9, uintptr(s), uintptr(iocc), uintptr(unsafe.Pointer(inbuf)), uintptr(cbif), uintptr(unsafe.Pointer(outbuf)), uintptr(cbob), uintptr(unsafe.Pointer(cbbr)), uintptr(unsafe.Pointer(overlapped)), uintptr(completionRoutine))
	if r1 == socket_error {
		err = errnoErr(e1)
	}
	return
}

func WSALookupServiceBegin(querySet *WSAQUERYSET, flags uint32, handle *Handle) (err error) {
	r1, _, e1 := syscall.Syscall(procWSALookupServiceBeginW.Addr(), 3, uintptr(unsafe.Pointer(querySet)), uintptr(flags), uintptr(unsafe.Pointer(handle)))
	if r1 == socket_error {
		err = errnoErr(e1)
	}
	return
}

func WSALookupServiceEnd(handle Handle) (err error) {
	r1, _, e1 := syscall.Syscall(procWSALookupServiceEnd.Addr(), 1, uintptr(handle), 0, 0)
	if r1 == socket_error {
		err = errnoErr(e1)
	}
	return
}

func WSALookupServiceNext(handle Handle, flags uint32, size *int32, querySet *WSAQUERYSET) (err error) {
	r1, _, e1 := syscall.Syscall6(procWSALookupServiceNextW.Addr(), 4, uintptr(handle), uintptr(flags), uintptr(unsafe.Pointer(size)), uintptr(unsafe.Pointer(querySet)), 0, 0)
	if r1 == socket_error {
		err = errnoErr(e1)
	}
	return
}

func WSARecv(s Handle, bufs *WSABuf, bufcnt uint32, recvd *uint32, flags *uint32, overlapped *Overlapped, croutine *byte) (err error) {
	r1, _, e1 := syscall.Syscall9(procWSARecv.Addr(), 7, uintptr(s), uintptr(unsafe.Pointer(bufs)), uintptr(bufcnt), uintptr(unsafe.Pointer(recvd)), uintptr(unsafe.Pointer(flags)), uintptr(unsafe.Pointer(overlapped)), uintptr(unsafe.Pointer(croutine)), 0, 0)
	if r1 == socket_error {
		err = errnoErr(e1)
	}
	return
}

func WSARecvFrom(s Handle, bufs *WSABuf, bufcnt uint32, recvd *uint32, flags *uint32, from *RawSockaddrAny, fromlen *int32, overlapped *Overlapped, croutine *byte) (err error) {
	r1, _, e1 := syscall.Syscall9(procWSARecvFrom.Addr(), 9, uintptr(s), uintptr(unsafe.Pointer(bufs)), uintptr(bufcnt), uintptr(unsafe.Pointer(recvd)), uintptr(unsafe.Pointer(flags)), uintptr(unsafe.Pointer(from)), uintptr(unsafe.Pointer(fromlen)), uintptr(unsafe.Pointer(overlapped)), uintptr(unsafe.Pointer(croutine)))
	if r1 == socket_error {
		err = errnoErr(e1)
	}
	return
}

func WSASend(s Handle, bufs *WSABuf, bufcnt uint32, sent *uint32, flags uint32, overlapped *Overlapped, croutine *byte) (err error) {
	r1, _, e1 := syscall.Syscall9(procWSASend.Addr(), 7, uintptr(s), uintptr(unsafe.Pointer(bufs)), uintptr(bufcnt), uintptr(unsafe.Pointer(sent)), uintptr(flags), uintptr(unsafe.Pointer(overlapped)), uintptr(unsafe.Pointer(croutine)), 0, 0)
	if r1 == socket_error {
		err = errnoErr(e1)
	}
	return
}

func WSASendTo(s Handle, bufs *WSABuf, bufcnt uint32, sent *uint32, flags uint32, to *RawSockaddrAny, tolen int32, overlapped *Overlapped, croutine *byte) (err error) {
	r1, _, e1 := syscall.Syscall9(procWSASendTo.Addr(), 9, uintptr(s), uintptr(unsafe.Pointer(bufs)), uintptr(bufcnt), uintptr(unsafe.Pointer(sent)), uintptr(flags), uintptr(unsafe.Pointer(to)), uintptr(tolen), uintptr(unsafe.Pointer(overlapped)), uintptr(unsafe.Pointer(croutine)))
	if r1 == socket_error {
		err = errnoErr(e1)
	}
	return
}

func WSASocket(af int32, typ int32, protocol int32, protoInfo *WSAProtocolInfo, group uint32, flags uint32) (handle Handle, err error) {
	r0, _, e1 := syscall.Syscall6(procWSASocketW.Addr(), 6, uintptr(af), uintptr(typ), uintptr(protocol), uintptr(unsafe.Pointer(protoInfo)), uintptr(group), uintptr(flags))
	handle = Handle(r0)
	if handle == InvalidHandle {
		err = errnoErr(e1)
	}
	return
}

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

func WTSEnumerateSessions(handle Handle, reserved uint32, version uint32, sessions **WTS_SESSION_INFO, count *uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procWTSEnumerateSessionsW.Addr(), 5, uintptr(handle), uintptr(reserved), uintptr(version), uintptr(unsafe.Pointer(sessions)), uintptr(unsafe.Pointer(count)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func WTSFreeMemory(ptr uintptr) {
	syscall.Syscall(procWTSFreeMemory.Addr(), 1, uintptr(ptr), 0, 0)
	return
}

func WTSQueryUserToken(session uint32, token *Token) (err error) {
	r1, _, e1 := syscall.Syscall(procWTSQueryUserToken.Addr(), 2, uintptr(session), uintptr(unsafe.Pointer(token)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}
