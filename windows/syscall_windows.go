// 版权所有 ? 2009 Go 作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 此许可证可在 LICENSE 文件中找到。

// Windows system calls.

package windows

import (
	errorspkg "errors"
	"fmt"
	"runtime"
	"sync"
	"syscall"
	"time"
	"unicode/utf16"
	"unsafe"
)

type Handle uintptr
type HWND uintptr

const (
	InvalidHandle = ^Handle(0)
	InvalidHWND   = ^HWND(0)

	// DefineDosDevice 的标志
	DDD_EXACT_MATCH_ON_REMOVE = 0x00000004
	DDD_NO_BROADCAST_SYSTEM   = 0x00000008
	DDD_RAW_TARGET_PATH       = 0x00000001
	DDD_REMOVE_DEFINITION     = 0x00000002

	// GetDriveType函数的返回值
	DRIVE_UNKNOWN     = 0
	DRIVE_NO_ROOT_DIR = 1
	DRIVE_REMOVABLE   = 2
	DRIVE_FIXED       = 3
	DRIVE_REMOTE      = 4
	DRIVE_CDROM       = 5
	DRIVE_RAMDISK     = 6

	// 文件系统标志，来源于 GetVolumeInformation 和 GetVolumeInformationByHandle 函数。
	FILE_CASE_SENSITIVE_SEARCH        = 0x00000001
	FILE_CASE_PRESERVED_NAMES         = 0x00000002
	FILE_FILE_COMPRESSION             = 0x00000010
	FILE_DAX_VOLUME                   = 0x20000000
	FILE_NAMED_STREAMS                = 0x00040000
	FILE_PERSISTENT_ACLS              = 0x00000008
	FILE_READ_ONLY_VOLUME             = 0x00080000
	FILE_SEQUENTIAL_WRITE_ONCE        = 0x00100000
	FILE_SUPPORTS_ENCRYPTION          = 0x00020000
	FILE_SUPPORTS_EXTENDED_ATTRIBUTES = 0x00800000
	FILE_SUPPORTS_HARD_LINKS          = 0x00400000
	FILE_SUPPORTS_OBJECT_IDS          = 0x00010000
	FILE_SUPPORTS_OPEN_BY_FILE_ID     = 0x01000000
	FILE_SUPPORTS_REPARSE_POINTS      = 0x00000080
	FILE_SUPPORTS_SPARSE_FILES        = 0x00000040
	FILE_SUPPORTS_TRANSACTIONS        = 0x00200000
	FILE_SUPPORTS_USN_JOURNAL         = 0x02000000
	FILE_UNICODE_ON_DISK              = 0x00000004
	FILE_VOLUME_IS_COMPRESSED         = 0x00008000
	FILE_VOLUME_QUOTAS                = 0x00000020

	// Flags for LockFileEx.
	LOCKFILE_FAIL_IMMEDIATELY = 0x00000001
	LOCKFILE_EXCLUSIVE_LOCK   = 0x00000002

	// SleepEx及其他APC函数的返回值
	WAIT_IO_COMPLETION = 0x000000C0
)

// StringToUTF16 已被弃用。请改用 UTF16FromString。
// 若 s 中包含 NUL 字节，此函数将触发 panic，而非返回错误。
func StringToUTF16(s string) []uint16 {
	a, err := UTF16FromString(s)
	if err != nil {
		panic("windows: string with NUL passed to StringToUTF16")
	}
	return a
}

// UTF16FromString返回UTF-8字符串s的UTF-16编码，并在其末尾添加终止空字符（NUL）。如果s在任意位置包含空字节（NUL），则返回(nil, syscall.EINVAL)。
func UTF16FromString(s string) ([]uint16, error) {
	return syscall.UTF16FromString(s)
}

// UTF16ToString 将UTF-16序列s转换为UTF-8编码形式，
// 并移除终止的空字符（NUL）及该字符之后的所有字节。
func UTF16ToString(s []uint16) string {
	return syscall.UTF16ToString(s)
}

// StringToUTF16Ptr 已被弃用。请改用 UTF16PtrFromString。
// 若 s 中包含 NUL 字节，此函数将引发 panic，而非返回错误。
func StringToUTF16Ptr(s string) *uint16 { return &StringToUTF16(s)[0] }

// UTF16PtrFromString返回对UTF-8字符串s的UTF-16编码的指针，同时在其末尾添加一个终止空字符（NUL）。如果s在任何位置包含空字节（NUL），则返回(nil, syscall.EINVAL)。
func UTF16PtrFromString(s string) (*uint16, error) {
	a, err := UTF16FromString(s)
	if err != nil {
		return nil, err
	}
	return &a[0], nil
}

// UTF16PtrToString 接收一个指向UTF-16序列的指针，并返回对应的UTF-8编码字符串。
// 若该指针为nil，则返回空字符串。此函数假设UTF-16序列以零字词（zero word）结尾；
// 若未出现零字词，程序可能会崩溃。
func UTF16PtrToString(p *uint16) string {
	if p == nil {
		return ""
	}
	if *p == 0 {
		return ""
	}

	// Find NUL terminator.
	n := 0
	for ptr := unsafe.Pointer(p); *(*uint16)(ptr) != 0; n++ {
		ptr = unsafe.Pointer(uintptr(ptr) + unsafe.Sizeof(*p))
	}
	return UTF16ToString(unsafe.Slice(p, n))
}

func Getpagesize() int { return 4096 }

// NewCallback 将一个 Go 函数转换为遵循 stdcall 调用约定的函数指针。
// 这在与需要回调的 Windows 代码进行互操作时非常有用。
// 该参数应为具有一个 uintptr 大小结果的函数。该函数不得包含大小大于 uintptr 的参数。
func NewCallback(fn interface{}) uintptr {
	return syscall.NewCallback(fn)
}

// NewCallbackCDecl 将一个Go函数转换为遵循cdecl调用约定的函数指针。
// 这在与需要回调的Windows代码进行互操作时非常有用。
// 该参数应为具有一个uintptr大小结果的函数。该函数不得包含大小大于uintptr的参数。
func NewCallbackCDecl(fn interface{}) uintptr {
	return syscall.NewCallbackCDecl(fn)
}

// windows api calls

//sys	GetLastError() (lasterr error)
//sys	LoadLibrary(libname string) (handle Handle, err error) = LoadLibraryW
//sys	LoadLibraryEx(libname string, zero Handle, flags uintptr) (handle Handle, err error) = LoadLibraryExW
//sys	FreeLibrary(handle Handle) (err error)
//sys	GetProcAddress(module Handle, procname string) (proc uintptr, err error)
//sys	GetModuleFileName(module Handle, filename *uint16, size uint32) (n uint32, err error) = kernel32.GetModuleFileNameW
//sys	GetModuleHandleEx(flags uint32, moduleName *uint16, module *Handle) (err error) = kernel32.GetModuleHandleExW
//sys	SetDefaultDllDirectories(directoryFlags uint32) (err error)
//sys	AddDllDirectory(path *uint16) (cookie uintptr, err error) = kernel32.AddDllDirectory
//sys	RemoveDllDirectory(cookie uintptr) (err error) = kernel32.RemoveDllDirectory
//sys	SetDllDirectory(path string) (err error) = kernel32.SetDllDirectoryW
//sys	GetVersion() (ver uint32, err error)
//sys	FormatMessage(flags uint32, msgsrc uintptr, msgid uint32, langid uint32, buf []uint16, args *byte) (n uint32, err error) = FormatMessageW
//sys	ExitProcess(exitcode uint32)
//sys	IsWow64Process(handle Handle, isWow64 *bool) (err error) = IsWow64Process
//sys	IsWow64Process2(handle Handle, processMachine *uint16, nativeMachine *uint16) (err error) = IsWow64Process2?
//sys	CreateFile(name *uint16, access uint32, mode uint32, sa *SecurityAttributes, createmode uint32, attrs uint32, templatefile Handle) (handle Handle, err error) [failretval==InvalidHandle] = CreateFileW
//sys	CreateNamedPipe(name *uint16, flags uint32, pipeMode uint32, maxInstances uint32, outSize uint32, inSize uint32, defaultTimeout uint32, sa *SecurityAttributes) (handle Handle, err error)  [failretval==InvalidHandle] = CreateNamedPipeW
//sys	ConnectNamedPipe(pipe Handle, overlapped *Overlapped) (err error)
//sys	DisconnectNamedPipe(pipe Handle) (err error)
//sys	GetNamedPipeInfo(pipe Handle, flags *uint32, outSize *uint32, inSize *uint32, maxInstances *uint32) (err error)
//sys	GetNamedPipeHandleState(pipe Handle, state *uint32, curInstances *uint32, maxCollectionCount *uint32, collectDataTimeout *uint32, userName *uint16, maxUserNameSize uint32) (err error) = GetNamedPipeHandleStateW
//sys	SetNamedPipeHandleState(pipe Handle, state *uint32, maxCollectionCount *uint32, collectDataTimeout *uint32) (err error) = SetNamedPipeHandleState
//sys	readFile(handle Handle, buf []byte, done *uint32, overlapped *Overlapped) (err error) = ReadFile
//sys	writeFile(handle Handle, buf []byte, done *uint32, overlapped *Overlapped) (err error) = WriteFile
//sys	GetOverlappedResult(handle Handle, overlapped *Overlapped, done *uint32, wait bool) (err error)
//sys	SetFilePointer(handle Handle, lowoffset int32, highoffsetptr *int32, whence uint32) (newlowoffset uint32, err error) [failretval==0xffffffff]
//sys	CloseHandle(handle Handle) (err error)
//sys	GetStdHandle(stdhandle uint32) (handle Handle, err error) [failretval==InvalidHandle]
//sys	SetStdHandle(stdhandle uint32, handle Handle) (err error)
//sys	findFirstFile1(name *uint16, data *win32finddata1) (handle Handle, err error) [failretval==InvalidHandle] = FindFirstFileW
//sys	findNextFile1(handle Handle, data *win32finddata1) (err error) = FindNextFileW
//sys	FindClose(handle Handle) (err error)
//sys	GetFileInformationByHandle(handle Handle, data *ByHandleFileInformation) (err error)
//sys	GetFileInformationByHandleEx(handle Handle, class uint32, outBuffer *byte, outBufferLen uint32) (err error)
//sys	SetFileInformationByHandle(handle Handle, class uint32, inBuffer *byte, inBufferLen uint32) (err error)
//sys	GetCurrentDirectory(buflen uint32, buf *uint16) (n uint32, err error) = GetCurrentDirectoryW
//sys	SetCurrentDirectory(path *uint16) (err error) = SetCurrentDirectoryW
//sys	CreateDirectory(path *uint16, sa *SecurityAttributes) (err error) = CreateDirectoryW
//sys	RemoveDirectory(path *uint16) (err error) = RemoveDirectoryW
//sys	DeleteFile(path *uint16) (err error) = DeleteFileW
//sys	MoveFile(from *uint16, to *uint16) (err error) = MoveFileW
//sys	MoveFileEx(from *uint16, to *uint16, flags uint32) (err error) = MoveFileExW
//sys	LockFileEx(file Handle, flags uint32, reserved uint32, bytesLow uint32, bytesHigh uint32, overlapped *Overlapped) (err error)
//sys	UnlockFileEx(file Handle, reserved uint32, bytesLow uint32, bytesHigh uint32, overlapped *Overlapped) (err error)
//sys	GetComputerName(buf *uint16, n *uint32) (err error) = GetComputerNameW
//sys	GetComputerNameEx(nametype uint32, buf *uint16, n *uint32) (err error) = GetComputerNameExW
//sys	SetEndOfFile(handle Handle) (err error)
//sys	SetFileValidData(handle Handle, validDataLength int64) (err error)
//sys	GetSystemTimeAsFileTime(time *Filetime)
//sys	GetSystemTimePreciseAsFileTime(time *Filetime)
//sys	GetTimeZoneInformation(tzi *Timezoneinformation) (rc uint32, err error) [failretval==0xffffffff]
//sys	CreateIoCompletionPort(filehandle Handle, cphandle Handle, key uintptr, threadcnt uint32) (handle Handle, err error)
//sys	GetQueuedCompletionStatus(cphandle Handle, qty *uint32, key *uintptr, overlapped **Overlapped, timeout uint32) (err error)
//sys	PostQueuedCompletionStatus(cphandle Handle, qty uint32, key uintptr, overlapped *Overlapped) (err error)
//sys	CancelIo(s Handle) (err error)
//sys	CancelIoEx(s Handle, o *Overlapped) (err error)
//sys	CreateProcess(appName *uint16, commandLine *uint16, procSecurity *SecurityAttributes, threadSecurity *SecurityAttributes, inheritHandles bool, creationFlags uint32, env *uint16, currentDir *uint16, startupInfo *StartupInfo, outProcInfo *ProcessInformation) (err error) = CreateProcessW
//sys	CreateProcessAsUser(token Token, appName *uint16, commandLine *uint16, procSecurity *SecurityAttributes, threadSecurity *SecurityAttributes, inheritHandles bool, creationFlags uint32, env *uint16, currentDir *uint16, startupInfo *StartupInfo, outProcInfo *ProcessInformation) (err error) = advapi32.CreateProcessAsUserW
//sys   initializeProcThreadAttributeList(attrlist *ProcThreadAttributeList, attrcount uint32, flags uint32, size *uintptr) (err error) = InitializeProcThreadAttributeList
//sys   deleteProcThreadAttributeList(attrlist *ProcThreadAttributeList) = DeleteProcThreadAttributeList
//sys   updateProcThreadAttribute(attrlist *ProcThreadAttributeList, flags uint32, attr uintptr, value unsafe.Pointer, size uintptr, prevvalue unsafe.Pointer, returnedsize *uintptr) (err error) = UpdateProcThreadAttribute
//sys	OpenProcess(desiredAccess uint32, inheritHandle bool, processId uint32) (handle Handle, err error)
//sys	ShellExecute(hwnd Handle, verb *uint16, file *uint16, args *uint16, cwd *uint16, showCmd int32) (err error) [failretval<=32] = shell32.ShellExecuteW
//sys	GetWindowThreadProcessId(hwnd HWND, pid *uint32) (tid uint32, err error) = user32.GetWindowThreadProcessId
//sys	GetShellWindow() (shellWindow HWND) = user32.GetShellWindow
//sys	MessageBox(hwnd HWND, text *uint16, caption *uint16, boxtype uint32) (ret int32, err error) [failretval==0] = user32.MessageBoxW
//sys	ExitWindowsEx(flags uint32, reason uint32) (err error) = user32.ExitWindowsEx
//sys	shGetKnownFolderPath(id *KNOWNFOLDERID, flags uint32, token Token, path **uint16) (ret error) = shell32.SHGetKnownFolderPath
//sys	TerminateProcess(handle Handle, exitcode uint32) (err error)
//sys	GetExitCodeProcess(handle Handle, exitcode *uint32) (err error)
//sys	getStartupInfo(startupInfo *StartupInfo) = GetStartupInfoW
//sys	GetProcessTimes(handle Handle, creationTime *Filetime, exitTime *Filetime, kernelTime *Filetime, userTime *Filetime) (err error)
//sys	DuplicateHandle(hSourceProcessHandle Handle, hSourceHandle Handle, hTargetProcessHandle Handle, lpTargetHandle *Handle, dwDesiredAccess uint32, bInheritHandle bool, dwOptions uint32) (err error)
//sys	WaitForSingleObject(handle Handle, waitMilliseconds uint32) (event uint32, err error) [failretval==0xffffffff]
//sys	waitForMultipleObjects(count uint32, handles uintptr, waitAll bool, waitMilliseconds uint32) (event uint32, err error) [failretval==0xffffffff] = WaitForMultipleObjects
//sys	GetTempPath(buflen uint32, buf *uint16) (n uint32, err error) = GetTempPathW
//sys	CreatePipe(readhandle *Handle, writehandle *Handle, sa *SecurityAttributes, size uint32) (err error)
//sys	GetFileType(filehandle Handle) (n uint32, err error)
//sys	CryptAcquireContext(provhandle *Handle, container *uint16, provider *uint16, provtype uint32, flags uint32) (err error) = advapi32.CryptAcquireContextW
//sys	CryptReleaseContext(provhandle Handle, flags uint32) (err error) = advapi32.CryptReleaseContext
//sys	CryptGenRandom(provhandle Handle, buflen uint32, buf *byte) (err error) = advapi32.CryptGenRandom
//sys	GetEnvironmentStrings() (envs *uint16, err error) [failretval==nil] = kernel32.GetEnvironmentStringsW
//sys	FreeEnvironmentStrings(envs *uint16) (err error) = kernel32.FreeEnvironmentStringsW
//sys	GetEnvironmentVariable(name *uint16, buffer *uint16, size uint32) (n uint32, err error) = kernel32.GetEnvironmentVariableW
//sys	SetEnvironmentVariable(name *uint16, value *uint16) (err error) = kernel32.SetEnvironmentVariableW
//sys	ExpandEnvironmentStrings(src *uint16, dst *uint16, size uint32) (n uint32, err error) = kernel32.ExpandEnvironmentStringsW
//sys	CreateEnvironmentBlock(block **uint16, token Token, inheritExisting bool) (err error) = userenv.CreateEnvironmentBlock
//sys	DestroyEnvironmentBlock(block *uint16) (err error) = userenv.DestroyEnvironmentBlock
//sys	getTickCount64() (ms uint64) = kernel32.GetTickCount64
//sys   GetFileTime(handle Handle, ctime *Filetime, atime *Filetime, wtime *Filetime) (err error)
//sys	SetFileTime(handle Handle, ctime *Filetime, atime *Filetime, wtime *Filetime) (err error)
//sys	GetFileAttributes(name *uint16) (attrs uint32, err error) [failretval==INVALID_FILE_ATTRIBUTES] = kernel32.GetFileAttributesW
//sys	SetFileAttributes(name *uint16, attrs uint32) (err error) = kernel32.SetFileAttributesW
//sys	GetFileAttributesEx(name *uint16, level uint32, info *byte) (err error) = kernel32.GetFileAttributesExW
//sys	GetCommandLine() (cmd *uint16) = kernel32.GetCommandLineW
//sys	commandLineToArgv(cmd *uint16, argc *int32) (argv **uint16, err error) [failretval==nil] = shell32.CommandLineToArgvW
//sys	LocalFree(hmem Handle) (handle Handle, err error) [failretval!=0]
//sys	LocalAlloc(flags uint32, length uint32) (ptr uintptr, err error)
//sys	SetHandleInformation(handle Handle, mask uint32, flags uint32) (err error)
//sys	FlushFileBuffers(handle Handle) (err error)
//sys	GetFullPathName(path *uint16, buflen uint32, buf *uint16, fname **uint16) (n uint32, err error) = kernel32.GetFullPathNameW
//sys	GetLongPathName(path *uint16, buf *uint16, buflen uint32) (n uint32, err error) = kernel32.GetLongPathNameW
//sys	GetShortPathName(longpath *uint16, shortpath *uint16, buflen uint32) (n uint32, err error) = kernel32.GetShortPathNameW
//sys	GetFinalPathNameByHandle(file Handle, filePath *uint16, filePathSize uint32, flags uint32) (n uint32, err error) = kernel32.GetFinalPathNameByHandleW
//sys	CreateFileMapping(fhandle Handle, sa *SecurityAttributes, prot uint32, maxSizeHigh uint32, maxSizeLow uint32, name *uint16) (handle Handle, err error) [failretval == 0 || e1 == ERROR_ALREADY_EXISTS] = kernel32.CreateFileMappingW
//sys	MapViewOfFile(handle Handle, access uint32, offsetHigh uint32, offsetLow uint32, length uintptr) (addr uintptr, err error)
//sys	UnmapViewOfFile(addr uintptr) (err error)
//sys	FlushViewOfFile(addr uintptr, length uintptr) (err error)
//sys	VirtualLock(addr uintptr, length uintptr) (err error)
//sys	VirtualUnlock(addr uintptr, length uintptr) (err error)
//sys	VirtualAlloc(address uintptr, size uintptr, alloctype uint32, protect uint32) (value uintptr, err error) = kernel32.VirtualAlloc
//sys	VirtualFree(address uintptr, size uintptr, freetype uint32) (err error) = kernel32.VirtualFree
//sys	VirtualProtect(address uintptr, size uintptr, newprotect uint32, oldprotect *uint32) (err error) = kernel32.VirtualProtect
//sys	VirtualProtectEx(process Handle, address uintptr, size uintptr, newProtect uint32, oldProtect *uint32) (err error) = kernel32.VirtualProtectEx
//sys	VirtualQuery(address uintptr, buffer *MemoryBasicInformation, length uintptr) (err error) = kernel32.VirtualQuery
//sys	VirtualQueryEx(process Handle, address uintptr, buffer *MemoryBasicInformation, length uintptr) (err error) = kernel32.VirtualQueryEx
//sys	ReadProcessMemory(process Handle, baseAddress uintptr, buffer *byte, size uintptr, numberOfBytesRead *uintptr) (err error) = kernel32.ReadProcessMemory
//sys	WriteProcessMemory(process Handle, baseAddress uintptr, buffer *byte, size uintptr, numberOfBytesWritten *uintptr) (err error) = kernel32.WriteProcessMemory
//sys	TransmitFile(s Handle, handle Handle, bytesToWrite uint32, bytsPerSend uint32, overlapped *Overlapped, transmitFileBuf *TransmitFileBuffers, flags uint32) (err error) = mswsock.TransmitFile
//sys	ReadDirectoryChanges(handle Handle, buf *byte, buflen uint32, watchSubTree bool, mask uint32, retlen *uint32, overlapped *Overlapped, completionRoutine uintptr) (err error) = kernel32.ReadDirectoryChangesW
//sys	FindFirstChangeNotification(path string, watchSubtree bool, notifyFilter uint32) (handle Handle, err error) [failretval==InvalidHandle] = kernel32.FindFirstChangeNotificationW
//sys	FindNextChangeNotification(handle Handle) (err error)
//sys	FindCloseChangeNotification(handle Handle) (err error)
//sys	CertOpenSystemStore(hprov Handle, name *uint16) (store Handle, err error) = crypt32.CertOpenSystemStoreW
//sys	CertOpenStore(storeProvider uintptr, msgAndCertEncodingType uint32, cryptProv uintptr, flags uint32, para uintptr) (handle Handle, err error) = crypt32.CertOpenStore
//sys	CertEnumCertificatesInStore(store Handle, prevContext *CertContext) (context *CertContext, err error) [failretval==nil] = crypt32.CertEnumCertificatesInStore
//sys	CertAddCertificateContextToStore(store Handle, certContext *CertContext, addDisposition uint32, storeContext **CertContext) (err error) = crypt32.CertAddCertificateContextToStore
//sys	CertCloseStore(store Handle, flags uint32) (err error) = crypt32.CertCloseStore
//sys	CertDeleteCertificateFromStore(certContext *CertContext) (err error) = crypt32.CertDeleteCertificateFromStore
//sys	CertDuplicateCertificateContext(certContext *CertContext) (dupContext *CertContext) = crypt32.CertDuplicateCertificateContext
//sys	PFXImportCertStore(pfx *CryptDataBlob, password *uint16, flags uint32) (store Handle, err error) = crypt32.PFXImportCertStore
//sys	CertGetCertificateChain(engine Handle, leaf *CertContext, time *Filetime, additionalStore Handle, para *CertChainPara, flags uint32, reserved uintptr, chainCtx **CertChainContext) (err error) = crypt32.CertGetCertificateChain
//sys	CertFreeCertificateChain(ctx *CertChainContext) = crypt32.CertFreeCertificateChain
//sys	CertCreateCertificateContext(certEncodingType uint32, certEncoded *byte, encodedLen uint32) (context *CertContext, err error) [failretval==nil] = crypt32.CertCreateCertificateContext
//sys	CertFreeCertificateContext(ctx *CertContext) (err error) = crypt32.CertFreeCertificateContext
//sys	CertVerifyCertificateChainPolicy(policyOID uintptr, chain *CertChainContext, para *CertChainPolicyPara, status *CertChainPolicyStatus) (err error) = crypt32.CertVerifyCertificateChainPolicy
//sys	CertGetNameString(certContext *CertContext, nameType uint32, flags uint32, typePara unsafe.Pointer, name *uint16, size uint32) (chars uint32) = crypt32.CertGetNameStringW
//sys	CertFindExtension(objId *byte, countExtensions uint32, extensions *CertExtension) (ret *CertExtension) = crypt32.CertFindExtension
//sys   CertFindCertificateInStore(store Handle, certEncodingType uint32, findFlags uint32, findType uint32, findPara unsafe.Pointer, prevCertContext *CertContext) (cert *CertContext, err error) [failretval==nil] = crypt32.CertFindCertificateInStore
//sys   CertFindChainInStore(store Handle, certEncodingType uint32, findFlags uint32, findType uint32, findPara unsafe.Pointer, prevChainContext *CertChainContext) (certchain *CertChainContext, err error) [failretval==nil] = crypt32.CertFindChainInStore
//sys   CryptAcquireCertificatePrivateKey(cert *CertContext, flags uint32, parameters unsafe.Pointer, cryptProvOrNCryptKey *Handle, keySpec *uint32, callerFreeProvOrNCryptKey *bool) (err error) = crypt32.CryptAcquireCertificatePrivateKey
//sys	CryptQueryObject(objectType uint32, object unsafe.Pointer, expectedContentTypeFlags uint32, expectedFormatTypeFlags uint32, flags uint32, msgAndCertEncodingType *uint32, contentType *uint32, formatType *uint32, certStore *Handle, msg *Handle, context *unsafe.Pointer) (err error) = crypt32.CryptQueryObject
//sys	CryptDecodeObject(encodingType uint32, structType *byte, encodedBytes *byte, lenEncodedBytes uint32, flags uint32, decoded unsafe.Pointer, decodedLen *uint32) (err error) = crypt32.CryptDecodeObject
//sys	CryptProtectData(dataIn *DataBlob, name *uint16, optionalEntropy *DataBlob, reserved uintptr, promptStruct *CryptProtectPromptStruct, flags uint32, dataOut *DataBlob) (err error) = crypt32.CryptProtectData
//sys	CryptUnprotectData(dataIn *DataBlob, name **uint16, optionalEntropy *DataBlob, reserved uintptr, promptStruct *CryptProtectPromptStruct, flags uint32, dataOut *DataBlob) (err error) = crypt32.CryptUnprotectData
//sys	WinVerifyTrustEx(hwnd HWND, actionId *GUID, data *WinTrustData) (ret error) = wintrust.WinVerifyTrustEx
//sys	RegOpenKeyEx(key Handle, subkey *uint16, options uint32, desiredAccess uint32, result *Handle) (regerrno error) = advapi32.RegOpenKeyExW
//sys	RegCloseKey(key Handle) (regerrno error) = advapi32.RegCloseKey
//sys	RegQueryInfoKey(key Handle, class *uint16, classLen *uint32, reserved *uint32, subkeysLen *uint32, maxSubkeyLen *uint32, maxClassLen *uint32, valuesLen *uint32, maxValueNameLen *uint32, maxValueLen *uint32, saLen *uint32, lastWriteTime *Filetime) (regerrno error) = advapi32.RegQueryInfoKeyW
//sys	RegEnumKeyEx(key Handle, index uint32, name *uint16, nameLen *uint32, reserved *uint32, class *uint16, classLen *uint32, lastWriteTime *Filetime) (regerrno error) = advapi32.RegEnumKeyExW
//sys	RegQueryValueEx(key Handle, name *uint16, reserved *uint32, valtype *uint32, buf *byte, buflen *uint32) (regerrno error) = advapi32.RegQueryValueExW
//sys	RegNotifyChangeKeyValue(key Handle, watchSubtree bool, notifyFilter uint32, event Handle, asynchronous bool) (regerrno error) = advapi32.RegNotifyChangeKeyValue
//sys	GetCurrentProcessId() (pid uint32) = kernel32.GetCurrentProcessId
//sys	ProcessIdToSessionId(pid uint32, sessionid *uint32) (err error) = kernel32.ProcessIdToSessionId
//sys	ClosePseudoConsole(console Handle) = kernel32.ClosePseudoConsole
//sys	createPseudoConsole(size uint32, in Handle, out Handle, flags uint32, pconsole *Handle) (hr error) = kernel32.CreatePseudoConsole
//sys	GetConsoleMode(console Handle, mode *uint32) (err error) = kernel32.GetConsoleMode
//sys	SetConsoleMode(console Handle, mode uint32) (err error) = kernel32.SetConsoleMode
//sys	GetConsoleScreenBufferInfo(console Handle, info *ConsoleScreenBufferInfo) (err error) = kernel32.GetConsoleScreenBufferInfo
//sys	setConsoleCursorPosition(console Handle, position uint32) (err error) = kernel32.SetConsoleCursorPosition
//sys	WriteConsole(console Handle, buf *uint16, towrite uint32, written *uint32, reserved *byte) (err error) = kernel32.WriteConsoleW
//sys	ReadConsole(console Handle, buf *uint16, toread uint32, read *uint32, inputControl *byte) (err error) = kernel32.ReadConsoleW
//sys	resizePseudoConsole(pconsole Handle, size uint32) (hr error) = kernel32.ResizePseudoConsole
//sys	CreateToolhelp32Snapshot(flags uint32, processId uint32) (handle Handle, err error) [failretval==InvalidHandle] = kernel32.CreateToolhelp32Snapshot
//sys	Module32First(snapshot Handle, moduleEntry *ModuleEntry32) (err error) = kernel32.Module32FirstW
//sys	Module32Next(snapshot Handle, moduleEntry *ModuleEntry32) (err error) = kernel32.Module32NextW
//sys	Process32First(snapshot Handle, procEntry *ProcessEntry32) (err error) = kernel32.Process32FirstW
//sys	Process32Next(snapshot Handle, procEntry *ProcessEntry32) (err error) = kernel32.Process32NextW
//sys	Thread32First(snapshot Handle, threadEntry *ThreadEntry32) (err error)
//sys	Thread32Next(snapshot Handle, threadEntry *ThreadEntry32) (err error)
//sys	DeviceIoControl(handle Handle, ioControlCode uint32, inBuffer *byte, inBufferSize uint32, outBuffer *byte, outBufferSize uint32, bytesReturned *uint32, overlapped *Overlapped) (err error)
// This function returns 1 byte BOOLEAN rather than the 4 byte BOOL.
//sys	CreateSymbolicLink(symlinkfilename *uint16, targetfilename *uint16, flags uint32) (err error) [failretval&0xff==0] = CreateSymbolicLinkW
//sys	CreateHardLink(filename *uint16, existingfilename *uint16, reserved uintptr) (err error) [failretval&0xff==0] = CreateHardLinkW
//sys	GetCurrentThreadId() (id uint32)
//sys	CreateEvent(eventAttrs *SecurityAttributes, manualReset uint32, initialState uint32, name *uint16) (handle Handle, err error) [failretval == 0 || e1 == ERROR_ALREADY_EXISTS] = kernel32.CreateEventW
//sys	CreateEventEx(eventAttrs *SecurityAttributes, name *uint16, flags uint32, desiredAccess uint32) (handle Handle, err error) [failretval == 0 || e1 == ERROR_ALREADY_EXISTS] = kernel32.CreateEventExW
//sys	OpenEvent(desiredAccess uint32, inheritHandle bool, name *uint16) (handle Handle, err error) = kernel32.OpenEventW
//sys	SetEvent(event Handle) (err error) = kernel32.SetEvent
//sys	ResetEvent(event Handle) (err error) = kernel32.ResetEvent
//sys	PulseEvent(event Handle) (err error) = kernel32.PulseEvent
//sys	CreateMutex(mutexAttrs *SecurityAttributes, initialOwner bool, name *uint16) (handle Handle, err error) [failretval == 0 || e1 == ERROR_ALREADY_EXISTS] = kernel32.CreateMutexW
//sys	CreateMutexEx(mutexAttrs *SecurityAttributes, name *uint16, flags uint32, desiredAccess uint32) (handle Handle, err error) [failretval == 0 || e1 == ERROR_ALREADY_EXISTS] = kernel32.CreateMutexExW
//sys	OpenMutex(desiredAccess uint32, inheritHandle bool, name *uint16) (handle Handle, err error) = kernel32.OpenMutexW
//sys	ReleaseMutex(mutex Handle) (err error) = kernel32.ReleaseMutex
//sys	SleepEx(milliseconds uint32, alertable bool) (ret uint32) = kernel32.SleepEx
//sys	CreateJobObject(jobAttr *SecurityAttributes, name *uint16) (handle Handle, err error) = kernel32.CreateJobObjectW
//sys	AssignProcessToJobObject(job Handle, process Handle) (err error) = kernel32.AssignProcessToJobObject
//sys	TerminateJobObject(job Handle, exitCode uint32) (err error) = kernel32.TerminateJobObject
//sys	SetErrorMode(mode uint32) (ret uint32) = kernel32.SetErrorMode
//sys	ResumeThread(thread Handle) (ret uint32, err error) [failretval==0xffffffff] = kernel32.ResumeThread
//sys	SetPriorityClass(process Handle, priorityClass uint32) (err error) = kernel32.SetPriorityClass
//sys	GetPriorityClass(process Handle) (ret uint32, err error) = kernel32.GetPriorityClass
//sys	QueryInformationJobObject(job Handle, JobObjectInformationClass int32, JobObjectInformation uintptr, JobObjectInformationLength uint32, retlen *uint32) (err error) = kernel32.QueryInformationJobObject
//sys	SetInformationJobObject(job Handle, JobObjectInformationClass uint32, JobObjectInformation uintptr, JobObjectInformationLength uint32) (ret int, err error)
//sys	GenerateConsoleCtrlEvent(ctrlEvent uint32, processGroupID uint32) (err error)
//sys	GetProcessId(process Handle) (id uint32, err error)
//sys	QueryFullProcessImageName(proc Handle, flags uint32, exeName *uint16, size *uint32) (err error) = kernel32.QueryFullProcessImageNameW
//sys	OpenThread(desiredAccess uint32, inheritHandle bool, threadId uint32) (handle Handle, err error)
//sys	SetProcessPriorityBoost(process Handle, disable bool) (err error) = kernel32.SetProcessPriorityBoost
//sys	GetProcessWorkingSetSizeEx(hProcess Handle, lpMinimumWorkingSetSize *uintptr, lpMaximumWorkingSetSize *uintptr, flags *uint32)
//sys	SetProcessWorkingSetSizeEx(hProcess Handle, dwMinimumWorkingSetSize uintptr, dwMaximumWorkingSetSize uintptr, flags uint32) (err error)
//sys	ClearCommBreak(handle Handle) (err error)
//sys	ClearCommError(handle Handle, lpErrors *uint32, lpStat *ComStat) (err error)
//sys	EscapeCommFunction(handle Handle, dwFunc uint32) (err error)
//sys	GetCommState(handle Handle, lpDCB *DCB) (err error)
//sys	GetCommModemStatus(handle Handle, lpModemStat *uint32) (err error)
//sys	GetCommTimeouts(handle Handle, timeouts *CommTimeouts) (err error)
//sys	PurgeComm(handle Handle, dwFlags uint32) (err error)
//sys	SetCommBreak(handle Handle) (err error)
//sys	SetCommMask(handle Handle, dwEvtMask uint32) (err error)
//sys	SetCommState(handle Handle, lpDCB *DCB) (err error)
//sys	SetCommTimeouts(handle Handle, timeouts *CommTimeouts) (err error)
//sys	SetupComm(handle Handle, dwInQueue uint32, dwOutQueue uint32) (err error)
//sys	WaitCommEvent(handle Handle, lpEvtMask *uint32, lpOverlapped *Overlapped) (err error)
//sys	GetActiveProcessorCount(groupNumber uint16) (ret uint32)
//sys	GetMaximumProcessorCount(groupNumber uint16) (ret uint32)
//sys	EnumWindows(enumFunc uintptr, param unsafe.Pointer) (err error) = user32.EnumWindows
//sys	EnumChildWindows(hwnd HWND, enumFunc uintptr, param unsafe.Pointer) = user32.EnumChildWindows
//sys	GetClassName(hwnd HWND, className *uint16, maxCount int32) (copied int32, err error) = user32.GetClassNameW
//sys	GetDesktopWindow() (hwnd HWND) = user32.GetDesktopWindow
//sys	GetForegroundWindow() (hwnd HWND) = user32.GetForegroundWindow
//sys	IsWindow(hwnd HWND) (isWindow bool) = user32.IsWindow
//sys	IsWindowUnicode(hwnd HWND) (isUnicode bool) = user32.IsWindowUnicode
//sys	IsWindowVisible(hwnd HWND) (isVisible bool) = user32.IsWindowVisible
//sys	GetGUIThreadInfo(thread uint32, info *GUIThreadInfo) (err error) = user32.GetGUIThreadInfo
//sys	GetLargePageMinimum() (size uintptr)

// Volume Management Functions
//sys	DefineDosDevice(flags uint32, deviceName *uint16, targetPath *uint16) (err error) = DefineDosDeviceW
//sys	DeleteVolumeMountPoint(volumeMountPoint *uint16) (err error) = DeleteVolumeMountPointW
//sys	FindFirstVolume(volumeName *uint16, bufferLength uint32) (handle Handle, err error) [failretval==InvalidHandle] = FindFirstVolumeW
//sys	FindFirstVolumeMountPoint(rootPathName *uint16, volumeMountPoint *uint16, bufferLength uint32) (handle Handle, err error) [failretval==InvalidHandle] = FindFirstVolumeMountPointW
//sys	FindNextVolume(findVolume Handle, volumeName *uint16, bufferLength uint32) (err error) = FindNextVolumeW
//sys	FindNextVolumeMountPoint(findVolumeMountPoint Handle, volumeMountPoint *uint16, bufferLength uint32) (err error) = FindNextVolumeMountPointW
//sys	FindVolumeClose(findVolume Handle) (err error)
//sys	FindVolumeMountPointClose(findVolumeMountPoint Handle) (err error)
//sys	GetDiskFreeSpaceEx(directoryName *uint16, freeBytesAvailableToCaller *uint64, totalNumberOfBytes *uint64, totalNumberOfFreeBytes *uint64) (err error) = GetDiskFreeSpaceExW
//sys	GetDriveType(rootPathName *uint16) (driveType uint32) = GetDriveTypeW
//sys	GetLogicalDrives() (drivesBitMask uint32, err error) [failretval==0]
//sys	GetLogicalDriveStrings(bufferLength uint32, buffer *uint16) (n uint32, err error) [failretval==0] = GetLogicalDriveStringsW
//sys	GetVolumeInformation(rootPathName *uint16, volumeNameBuffer *uint16, volumeNameSize uint32, volumeNameSerialNumber *uint32, maximumComponentLength *uint32, fileSystemFlags *uint32, fileSystemNameBuffer *uint16, fileSystemNameSize uint32) (err error) = GetVolumeInformationW
//sys	GetVolumeInformationByHandle(file Handle, volumeNameBuffer *uint16, volumeNameSize uint32, volumeNameSerialNumber *uint32, maximumComponentLength *uint32, fileSystemFlags *uint32, fileSystemNameBuffer *uint16, fileSystemNameSize uint32) (err error) = GetVolumeInformationByHandleW
//sys	GetVolumeNameForVolumeMountPoint(volumeMountPoint *uint16, volumeName *uint16, bufferlength uint32) (err error) = GetVolumeNameForVolumeMountPointW
//sys	GetVolumePathName(fileName *uint16, volumePathName *uint16, bufferLength uint32) (err error) = GetVolumePathNameW
//sys	GetVolumePathNamesForVolumeName(volumeName *uint16, volumePathNames *uint16, bufferLength uint32, returnLength *uint32) (err error) = GetVolumePathNamesForVolumeNameW
//sys	QueryDosDevice(deviceName *uint16, targetPath *uint16, max uint32) (n uint32, err error) [failretval==0] = QueryDosDeviceW
//sys	SetVolumeLabel(rootPathName *uint16, volumeName *uint16) (err error) = SetVolumeLabelW
//sys	SetVolumeMountPoint(volumeMountPoint *uint16, volumeName *uint16) (err error) = SetVolumeMountPointW
//sys	InitiateSystemShutdownEx(machineName *uint16, message *uint16, timeout uint32, forceAppsClosed bool, rebootAfterShutdown bool, reason uint32) (err error) = advapi32.InitiateSystemShutdownExW
//sys	SetProcessShutdownParameters(level uint32, flags uint32) (err error) = kernel32.SetProcessShutdownParameters
//sys	GetProcessShutdownParameters(level *uint32, flags *uint32) (err error) = kernel32.GetProcessShutdownParameters
//sys	clsidFromString(lpsz *uint16, pclsid *GUID) (ret error) = ole32.CLSIDFromString
//sys	stringFromGUID2(rguid *GUID, lpsz *uint16, cchMax int32) (chars int32) = ole32.StringFromGUID2
//sys	coCreateGuid(pguid *GUID) (ret error) = ole32.CoCreateGuid
//sys	CoTaskMemFree(address unsafe.Pointer) = ole32.CoTaskMemFree
//sys	CoInitializeEx(reserved uintptr, coInit uint32) (ret error) = ole32.CoInitializeEx
//sys	CoUninitialize() = ole32.CoUninitialize
//sys	CoGetObject(name *uint16, bindOpts *BIND_OPTS3, guid *GUID, functionTable **uintptr) (ret error) = ole32.CoGetObject
//sys	getProcessPreferredUILanguages(flags uint32, numLanguages *uint32, buf *uint16, bufSize *uint32) (err error) = kernel32.GetProcessPreferredUILanguages
//sys	getThreadPreferredUILanguages(flags uint32, numLanguages *uint32, buf *uint16, bufSize *uint32) (err error) = kernel32.GetThreadPreferredUILanguages
//sys	getUserPreferredUILanguages(flags uint32, numLanguages *uint32, buf *uint16, bufSize *uint32) (err error) = kernel32.GetUserPreferredUILanguages
//sys	getSystemPreferredUILanguages(flags uint32, numLanguages *uint32, buf *uint16, bufSize *uint32) (err error) = kernel32.GetSystemPreferredUILanguages
//sys	findResource(module Handle, name uintptr, resType uintptr) (resInfo Handle, err error) = kernel32.FindResourceW
//sys	SizeofResource(module Handle, resInfo Handle) (size uint32, err error) = kernel32.SizeofResource
//sys	LoadResource(module Handle, resInfo Handle) (resData Handle, err error) = kernel32.LoadResource
//sys	LockResource(resData Handle) (addr uintptr, err error) = kernel32.LockResource

// Version APIs
//sys	GetFileVersionInfoSize(filename string, zeroHandle *Handle) (bufSize uint32, err error) = version.GetFileVersionInfoSizeW
//sys	GetFileVersionInfo(filename string, handle uint32, bufSize uint32, buffer unsafe.Pointer) (err error) = version.GetFileVersionInfoW
//sys	VerQueryValue(block unsafe.Pointer, subBlock string, pointerToBufferPointer unsafe.Pointer, bufSize *uint32) (err error) = version.VerQueryValueW

// Process Status API (PSAPI)
//sys	enumProcesses(processIds *uint32, nSize uint32, bytesReturned *uint32) (err error) = psapi.EnumProcesses
//sys	EnumProcessModules(process Handle, module *Handle, cb uint32, cbNeeded *uint32) (err error) = psapi.EnumProcessModules
//sys	EnumProcessModulesEx(process Handle, module *Handle, cb uint32, cbNeeded *uint32, filterFlag uint32) (err error) = psapi.EnumProcessModulesEx
//sys	GetModuleInformation(process Handle, module Handle, modinfo *ModuleInfo, cb uint32) (err error) = psapi.GetModuleInformation
//sys	GetModuleFileNameEx(process Handle, module Handle, filename *uint16, size uint32) (err error) = psapi.GetModuleFileNameExW
//sys	GetModuleBaseName(process Handle, module Handle, baseName *uint16, size uint32) (err error) = psapi.GetModuleBaseNameW
//sys   QueryWorkingSetEx(process Handle, pv uintptr, cb uint32) (err error) = psapi.QueryWorkingSetEx

// NT Native APIs
//sys	rtlNtStatusToDosErrorNoTeb(ntstatus NTStatus) (ret syscall.Errno) = ntdll.RtlNtStatusToDosErrorNoTeb
//sys	rtlGetVersion(info *OsVersionInfoEx) (ntstatus error) = ntdll.RtlGetVersion
//sys	rtlGetNtVersionNumbers(majorVersion *uint32, minorVersion *uint32, buildNumber *uint32) = ntdll.RtlGetNtVersionNumbers
//sys	RtlGetCurrentPeb() (peb *PEB) = ntdll.RtlGetCurrentPeb
//sys	RtlInitUnicodeString(destinationString *NTUnicodeString, sourceString *uint16) = ntdll.RtlInitUnicodeString
//sys	RtlInitString(destinationString *NTString, sourceString *byte) = ntdll.RtlInitString
//sys	NtCreateFile(handle *Handle, access uint32, oa *OBJECT_ATTRIBUTES, iosb *IO_STATUS_BLOCK, allocationSize *int64, attributes uint32, share uint32, disposition uint32, options uint32, eabuffer uintptr, ealength uint32) (ntstatus error) = ntdll.NtCreateFile
//sys	NtCreateNamedPipeFile(pipe *Handle, access uint32, oa *OBJECT_ATTRIBUTES, iosb *IO_STATUS_BLOCK, share uint32, disposition uint32, options uint32, typ uint32, readMode uint32, completionMode uint32, maxInstances uint32, inboundQuota uint32, outputQuota uint32, timeout *int64) (ntstatus error) = ntdll.NtCreateNamedPipeFile
//sys	NtSetInformationFile(handle Handle, iosb *IO_STATUS_BLOCK, inBuffer *byte, inBufferLen uint32, class uint32) (ntstatus error) = ntdll.NtSetInformationFile
//sys	RtlDosPathNameToNtPathName(dosName *uint16, ntName *NTUnicodeString, ntFileNamePart *uint16, relativeName *RTL_RELATIVE_NAME) (ntstatus error) = ntdll.RtlDosPathNameToNtPathName_U_WithStatus
//sys	RtlDosPathNameToRelativeNtPathName(dosName *uint16, ntName *NTUnicodeString, ntFileNamePart *uint16, relativeName *RTL_RELATIVE_NAME) (ntstatus error) = ntdll.RtlDosPathNameToRelativeNtPathName_U_WithStatus
//sys	RtlDefaultNpAcl(acl **ACL) (ntstatus error) = ntdll.RtlDefaultNpAcl
//sys	NtQueryInformationProcess(proc Handle, procInfoClass int32, procInfo unsafe.Pointer, procInfoLen uint32, retLen *uint32) (ntstatus error) = ntdll.NtQueryInformationProcess
//sys	NtSetInformationProcess(proc Handle, procInfoClass int32, procInfo unsafe.Pointer, procInfoLen uint32) (ntstatus error) = ntdll.NtSetInformationProcess
//sys	NtQuerySystemInformation(sysInfoClass int32, sysInfo unsafe.Pointer, sysInfoLen uint32, retLen *uint32) (ntstatus error) = ntdll.NtQuerySystemInformation
//sys	NtSetSystemInformation(sysInfoClass int32, sysInfo unsafe.Pointer, sysInfoLen uint32) (ntstatus error) = ntdll.NtSetSystemInformation
//sys	RtlAddFunctionTable(functionTable *RUNTIME_FUNCTION, entryCount uint32, baseAddress uintptr) (ret bool) = ntdll.RtlAddFunctionTable
//sys	RtlDeleteFunctionTable(functionTable *RUNTIME_FUNCTION) (ret bool) = ntdll.RtlDeleteFunctionTable

// Desktop Window Manager API (Dwmapi)
//sys	DwmGetWindowAttribute(hwnd HWND, attribute uint32, value unsafe.Pointer, size uint32) (ret error) = dwmapi.DwmGetWindowAttribute
//sys	DwmSetWindowAttribute(hwnd HWND, attribute uint32, value unsafe.Pointer, size uint32) (ret error) = dwmapi.DwmSetWindowAttribute

// Windows Multimedia API
//sys TimeBeginPeriod (period uint32) (err error) [failretval != 0] = winmm.timeBeginPeriod
//sys TimeEndPeriod (period uint32) (err error) [failretval != 0] = winmm.timeEndPeriod

// 为其他包实现的系统调用接口

// GetCurrentProcess 返回当前进程的句柄。
// 它是一个伪句柄，无需关闭。
// 返回的错误始终为 nil。
// 
// 已弃用：使用 CurrentProcess 直接获取相同的 Handle，且无 nil 错误。
func GetCurrentProcess() (Handle, error) {
	return CurrentProcess(), nil
}

// CurrentProcess 返回当前进程的句柄。
// 它是一个无需关闭的伪句柄。
func CurrentProcess() Handle { return Handle(^uintptr(1 - 1)) }

// GetCurrentThread 返回当前线程的句柄。
// 它是一个无需关闭的伪句柄。
// 返回的错误始终为 nil。
//
// 已废弃: 请使用 CurrentThread 直接获取相同的 Handle，其不会带有 nil 错误。
func GetCurrentThread() (Handle, error) {
	return CurrentThread(), nil
}

// CurrentThread 返回当前线程的句柄。
// 它是一个伪句柄，无需关闭。
func CurrentThread() Handle { return Handle(^uintptr(2 - 1)) }

// GetProcAddressByOrdinal 通过序数从模块中获取导出函数的地址
func GetProcAddressByOrdinal(module Handle, ordinal uintptr) (proc uintptr, err error) {
	r0, _, e1 := syscall.Syscall(procGetProcAddress.Addr(), 2, uintptr(module), ordinal, 0)
	proc = uintptr(r0)
	if proc == 0 {
		err = errnoErr(e1)
	}
	return
}

func Exit(code int) { ExitProcess(uint32(code)) }

func makeInheritSa() *SecurityAttributes {
	var sa SecurityAttributes
	sa.Length = uint32(unsafe.Sizeof(sa))
	sa.InheritHandle = 1
	return &sa
}

func Open(path string, mode int, perm uint32) (fd Handle, err error) {
	if len(path) == 0 {
		return InvalidHandle, ERROR_FILE_NOT_FOUND
	}
	pathp, err := UTF16PtrFromString(path)
	if err != nil {
		return InvalidHandle, err
	}
	var access uint32
	switch mode & (O_RDONLY | O_WRONLY | O_RDWR) {
	case O_RDONLY:
		access = GENERIC_READ
	case O_WRONLY:
		access = GENERIC_WRITE
	case O_RDWR:
		access = GENERIC_READ | GENERIC_WRITE
	}
	if mode&O_CREAT != 0 {
		access |= GENERIC_WRITE
	}
	if mode&O_APPEND != 0 {
		access &^= GENERIC_WRITE
		access |= FILE_APPEND_DATA
	}
	sharemode := uint32(FILE_SHARE_READ | FILE_SHARE_WRITE)
	var sa *SecurityAttributes
	if mode&O_CLOEXEC == 0 {
		sa = makeInheritSa()
	}
	var createmode uint32
	switch {
	case mode&(O_CREAT|O_EXCL) == (O_CREAT | O_EXCL):
		createmode = CREATE_NEW
	case mode&(O_CREAT|O_TRUNC) == (O_CREAT | O_TRUNC):
		createmode = CREATE_ALWAYS
	case mode&O_CREAT == O_CREAT:
		createmode = OPEN_ALWAYS
	case mode&O_TRUNC == O_TRUNC:
		createmode = TRUNCATE_EXISTING
	default:
		createmode = OPEN_EXISTING
	}
	var attrs uint32 = FILE_ATTRIBUTE_NORMAL
	if perm&S_IWRITE == 0 {
		attrs = FILE_ATTRIBUTE_READONLY
	}
	h, e := CreateFile(pathp, access, sharemode, sa, createmode, attrs, 0)
	return h, e
}

func Read(fd Handle, p []byte) (n int, err error) {
	var done uint32
	e := ReadFile(fd, p, &done, nil)
	if e != nil {
		if e == ERROR_BROKEN_PIPE {
			// 注意(brainman): 当从stdin读取EOF时，为解决返回ERROR_BROKEN_PIPE的问题
			return 0, nil
		}
		return 0, e
	}
	return int(done), nil
}

func Write(fd Handle, p []byte) (n int, err error) {
	if raceenabled {
		raceReleaseMerge(unsafe.Pointer(&ioSync))
	}
	var done uint32
	e := WriteFile(fd, p, &done, nil)
	if e != nil {
		return 0, e
	}
	return int(done), nil
}

func ReadFile(fd Handle, p []byte, done *uint32, overlapped *Overlapped) error {
	err := readFile(fd, p, done, overlapped)
	if raceenabled {
		if *done > 0 {
			raceWriteRange(unsafe.Pointer(&p[0]), int(*done))
		}
		raceAcquire(unsafe.Pointer(&ioSync))
	}
	return err
}

func WriteFile(fd Handle, p []byte, done *uint32, overlapped *Overlapped) error {
	if raceenabled {
		raceReleaseMerge(unsafe.Pointer(&ioSync))
	}
	err := writeFile(fd, p, done, overlapped)
	if raceenabled && *done > 0 {
		raceReadRange(unsafe.Pointer(&p[0]), int(*done))
	}
	return err
}

var ioSync int64

func Seek(fd Handle, offset int64, whence int) (newoffset int64, err error) {
	var w uint32
	switch whence {
	case 0:
		w = FILE_BEGIN
	case 1:
		w = FILE_CURRENT
	case 2:
		w = FILE_END
	}
	hi := int32(offset >> 32)
	lo := int32(offset)
	// 使用GetFileType检查管道，管道不能进行seek操作
	ft, _ := GetFileType(fd)
	if ft == FILE_TYPE_PIPE {
		return 0, syscall.EPIPE
	}
	rlo, e := SetFilePointer(fd, lo, &hi, w)
	if e != nil {
		return 0, e
	}
	return int64(hi)<<32 + int64(rlo), nil
}

func Close(fd Handle) (err error) {
	return CloseHandle(fd)
}

var (
	Stdin  = getStdHandle(STD_INPUT_HANDLE)
	Stdout = getStdHandle(STD_OUTPUT_HANDLE)
	Stderr = getStdHandle(STD_ERROR_HANDLE)
)

func getStdHandle(stdhandle uint32) (fd Handle) {
	r, _ := GetStdHandle(stdhandle)
	return r
}

const ImplementsGetwd = true

func Getwd() (wd string, err error) {
	b := make([]uint16, 300)
	n, e := GetCurrentDirectory(uint32(len(b)), &b[0])
	if e != nil {
		return "", e
	}
	return string(utf16.Decode(b[0:n])), nil
}

func Chdir(path string) (err error) {
	pathp, err := UTF16PtrFromString(path)
	if err != nil {
		return err
	}
	return SetCurrentDirectory(pathp)
}

func Mkdir(path string, mode uint32) (err error) {
	pathp, err := UTF16PtrFromString(path)
	if err != nil {
		return err
	}
	return CreateDirectory(pathp, nil)
}

func Rmdir(path string) (err error) {
	pathp, err := UTF16PtrFromString(path)
	if err != nil {
		return err
	}
	return RemoveDirectory(pathp)
}

func Unlink(path string) (err error) {
	pathp, err := UTF16PtrFromString(path)
	if err != nil {
		return err
	}
	return DeleteFile(pathp)
}

func Rename(oldpath, newpath string) (err error) {
	from, err := UTF16PtrFromString(oldpath)
	if err != nil {
		return err
	}
	to, err := UTF16PtrFromString(newpath)
	if err != nil {
		return err
	}
	return MoveFileEx(from, to, MOVEFILE_REPLACE_EXISTING)
}

func ComputerName() (name string, err error) {
	var n uint32 = MAX_COMPUTERNAME_LENGTH + 1
	b := make([]uint16, n)
	e := GetComputerName(&b[0], &n)
	if e != nil {
		return "", e
	}
	return string(utf16.Decode(b[0:n])), nil
}

func DurationSinceBoot() time.Duration {
	return time.Duration(getTickCount64()) * time.Millisecond
}

func Ftruncate(fd Handle, length int64) (err error) {
	curoffset, e := Seek(fd, 0, 1)
	if e != nil {
		return e
	}
	defer Seek(fd, curoffset, 0)
	_, e = Seek(fd, length, 0)
	if e != nil {
		return e
	}
	e = SetEndOfFile(fd)
	if e != nil {
		return e
	}
	return nil
}

func Gettimeofday(tv *Timeval) (err error) {
	var ft Filetime
	GetSystemTimeAsFileTime(&ft)
	*tv = NsecToTimeval(ft.Nanoseconds())
	return nil
}

func Pipe(p []Handle) (err error) {
	if len(p) != 2 {
		return syscall.EINVAL
	}
	var r, w Handle
	e := CreatePipe(&r, &w, makeInheritSa(), 0)
	if e != nil {
		return e
	}
	p[0] = r
	p[1] = w
	return nil
}

func Utimes(path string, tv []Timeval) (err error) {
	if len(tv) != 2 {
		return syscall.EINVAL
	}
	pathp, e := UTF16PtrFromString(path)
	if e != nil {
		return e
	}
	h, e := CreateFile(pathp,
		FILE_WRITE_ATTRIBUTES, FILE_SHARE_WRITE, nil,
		OPEN_EXISTING, FILE_FLAG_BACKUP_SEMANTICS, 0)
	if e != nil {
		return e
	}
	defer CloseHandle(h)
	a := NsecToFiletime(tv[0].Nanoseconds())
	w := NsecToFiletime(tv[1].Nanoseconds())
	return SetFileTime(h, nil, &a, &w)
}

func UtimesNano(path string, ts []Timespec) (err error) {
	if len(ts) != 2 {
		return syscall.EINVAL
	}
	pathp, e := UTF16PtrFromString(path)
	if e != nil {
		return e
	}
	h, e := CreateFile(pathp,
		FILE_WRITE_ATTRIBUTES, FILE_SHARE_WRITE, nil,
		OPEN_EXISTING, FILE_FLAG_BACKUP_SEMANTICS, 0)
	if e != nil {
		return e
	}
	defer CloseHandle(h)
	a := NsecToFiletime(TimespecToNsec(ts[0]))
	w := NsecToFiletime(TimespecToNsec(ts[1]))
	return SetFileTime(h, nil, &a, &w)
}

func Fsync(fd Handle) (err error) {
	return FlushFileBuffers(fd)
}

func Chmod(path string, mode uint32) (err error) {
	p, e := UTF16PtrFromString(path)
	if e != nil {
		return e
	}
	attrs, e := GetFileAttributes(p)
	if e != nil {
		return e
	}
	if mode&S_IWRITE != 0 {
		attrs &^= FILE_ATTRIBUTE_READONLY
	} else {
		attrs |= FILE_ATTRIBUTE_READONLY
	}
	return SetFileAttributes(p, attrs)
}

func LoadGetSystemTimePreciseAsFileTime() error {
	return procGetSystemTimePreciseAsFileTime.Find()
}

func LoadCancelIoEx() error {
	return procCancelIoEx.Find()
}

func LoadSetFileCompletionNotificationModes() error {
	return procSetFileCompletionNotificationModes.Find()
}

func WaitForMultipleObjects(handles []Handle, waitAll bool, waitMilliseconds uint32) (event uint32, err error) {
// 除本函数外，所有其他 Win32 数组 API 都采用“指针，计数”作为参数。因此，我们无法将其声明为常规的 [] 类型，因为 mksyscall 会使用相反的顺序。因此，我们自己为此简单地创建了一个存根（stub）。

	var handlePtr *Handle
	if len(handles) > 0 {
		handlePtr = &handles[0]
	}
	return waitForMultipleObjects(uint32(len(handles)), uintptr(unsafe.Pointer(handlePtr)), waitAll, waitMilliseconds)
}

// net api calls

const socket_error = uintptr(^uint32(0))

//sys	WSAStartup(verreq uint32, data *WSAData) (sockerr error) = ws2_32.WSAStartup
//sys	WSACleanup() (err error) [failretval==socket_error] = ws2_32.WSACleanup
//sys	WSAIoctl(s Handle, iocc uint32, inbuf *byte, cbif uint32, outbuf *byte, cbob uint32, cbbr *uint32, overlapped *Overlapped, completionRoutine uintptr) (err error) [failretval==socket_error] = ws2_32.WSAIoctl
//sys	WSALookupServiceBegin(querySet *WSAQUERYSET, flags uint32, handle *Handle) (err error) [failretval==socket_error] = ws2_32.WSALookupServiceBeginW
//sys	WSALookupServiceNext(handle Handle, flags uint32, size *int32, querySet *WSAQUERYSET) (err error) [failretval==socket_error] = ws2_32.WSALookupServiceNextW
//sys	WSALookupServiceEnd(handle Handle) (err error) [failretval==socket_error] = ws2_32.WSALookupServiceEnd
//sys	socket(af int32, typ int32, protocol int32) (handle Handle, err error) [failretval==InvalidHandle] = ws2_32.socket
//sys	sendto(s Handle, buf []byte, flags int32, to unsafe.Pointer, tolen int32) (err error) [failretval==socket_error] = ws2_32.sendto
//sys	recvfrom(s Handle, buf []byte, flags int32, from *RawSockaddrAny, fromlen *int32) (n int32, err error) [failretval==-1] = ws2_32.recvfrom
//sys	Setsockopt(s Handle, level int32, optname int32, optval *byte, optlen int32) (err error) [failretval==socket_error] = ws2_32.setsockopt
//sys	Getsockopt(s Handle, level int32, optname int32, optval *byte, optlen *int32) (err error) [failretval==socket_error] = ws2_32.getsockopt
//sys	bind(s Handle, name unsafe.Pointer, namelen int32) (err error) [failretval==socket_error] = ws2_32.bind
//sys	connect(s Handle, name unsafe.Pointer, namelen int32) (err error) [failretval==socket_error] = ws2_32.connect
//sys	getsockname(s Handle, rsa *RawSockaddrAny, addrlen *int32) (err error) [failretval==socket_error] = ws2_32.getsockname
//sys	getpeername(s Handle, rsa *RawSockaddrAny, addrlen *int32) (err error) [failretval==socket_error] = ws2_32.getpeername
//sys	listen(s Handle, backlog int32) (err error) [failretval==socket_error] = ws2_32.listen
//sys	shutdown(s Handle, how int32) (err error) [failretval==socket_error] = ws2_32.shutdown
//sys	Closesocket(s Handle) (err error) [failretval==socket_error] = ws2_32.closesocket
//sys	AcceptEx(ls Handle, as Handle, buf *byte, rxdatalen uint32, laddrlen uint32, raddrlen uint32, recvd *uint32, overlapped *Overlapped) (err error) = mswsock.AcceptEx
//sys	GetAcceptExSockaddrs(buf *byte, rxdatalen uint32, laddrlen uint32, raddrlen uint32, lrsa **RawSockaddrAny, lrsalen *int32, rrsa **RawSockaddrAny, rrsalen *int32) = mswsock.GetAcceptExSockaddrs
//sys	WSARecv(s Handle, bufs *WSABuf, bufcnt uint32, recvd *uint32, flags *uint32, overlapped *Overlapped, croutine *byte) (err error) [failretval==socket_error] = ws2_32.WSARecv
//sys	WSASend(s Handle, bufs *WSABuf, bufcnt uint32, sent *uint32, flags uint32, overlapped *Overlapped, croutine *byte) (err error) [failretval==socket_error] = ws2_32.WSASend
//sys	WSARecvFrom(s Handle, bufs *WSABuf, bufcnt uint32, recvd *uint32, flags *uint32,  from *RawSockaddrAny, fromlen *int32, overlapped *Overlapped, croutine *byte) (err error) [failretval==socket_error] = ws2_32.WSARecvFrom
//sys	WSASendTo(s Handle, bufs *WSABuf, bufcnt uint32, sent *uint32, flags uint32, to *RawSockaddrAny, tolen int32,  overlapped *Overlapped, croutine *byte) (err error) [failretval==socket_error] = ws2_32.WSASendTo
//sys	WSASocket(af int32, typ int32, protocol int32, protoInfo *WSAProtocolInfo, group uint32, flags uint32) (handle Handle, err error) [failretval==InvalidHandle] = ws2_32.WSASocketW
//sys	GetHostByName(name string) (h *Hostent, err error) [failretval==nil] = ws2_32.gethostbyname
//sys	GetServByName(name string, proto string) (s *Servent, err error) [failretval==nil] = ws2_32.getservbyname
//sys	Ntohs(netshort uint16) (u uint16) = ws2_32.ntohs
//sys	GetProtoByName(name string) (p *Protoent, err error) [failretval==nil] = ws2_32.getprotobyname
//sys	DnsQuery(name string, qtype uint16, options uint32, extra *byte, qrs **DNSRecord, pr *byte) (status error) = dnsapi.DnsQuery_W
//sys	DnsRecordListFree(rl *DNSRecord, freetype uint32) = dnsapi.DnsRecordListFree
//sys	DnsNameCompare(name1 *uint16, name2 *uint16) (same bool) = dnsapi.DnsNameCompare_W
//sys	GetAddrInfoW(nodename *uint16, servicename *uint16, hints *AddrinfoW, result **AddrinfoW) (sockerr error) = ws2_32.GetAddrInfoW
//sys	FreeAddrInfoW(addrinfo *AddrinfoW) = ws2_32.FreeAddrInfoW
//sys	GetIfEntry(pIfRow *MibIfRow) (errcode error) = iphlpapi.GetIfEntry
//sys	GetAdaptersInfo(ai *IpAdapterInfo, ol *uint32) (errcode error) = iphlpapi.GetAdaptersInfo
//sys	SetFileCompletionNotificationModes(handle Handle, flags uint8) (err error) = kernel32.SetFileCompletionNotificationModes
//sys	WSAEnumProtocols(protocols *int32, protocolBuffer *WSAProtocolInfo, bufferLength *uint32) (n int32, err error) [failretval==-1] = ws2_32.WSAEnumProtocolsW
//sys	WSAGetOverlappedResult(h Handle, o *Overlapped, bytes *uint32, wait bool, flags *uint32) (err error) = ws2_32.WSAGetOverlappedResult
//sys	GetAdaptersAddresses(family uint32, flags uint32, reserved uintptr, adapterAddresses *IpAdapterAddresses, sizePointer *uint32) (errcode error) = iphlpapi.GetAdaptersAddresses
//sys	GetACP() (acp uint32) = kernel32.GetACP
//sys	MultiByteToWideChar(codePage uint32, dwFlags uint32, str *byte, nstr int32, wchar *uint16, nwchar int32) (nwrite int32, err error) = kernel32.MultiByteToWideChar
//sys	getBestInterfaceEx(sockaddr unsafe.Pointer, pdwBestIfIndex *uint32) (errcode error) = iphlpapi.GetBestInterfaceEx

// 用于测试：客户端可以设置此标志，以强制IPv6套接字创建返回EAFNOSUPPORT错误。
var SocketDisableIPv6 bool

type RawSockaddrInet4 struct {
	Family uint16
	Port   uint16
	Addr   [4]byte /* in_addr */
	Zero   [8]uint8
}

type RawSockaddrInet6 struct {
	Family   uint16
	Port     uint16
	Flowinfo uint32
	Addr     [16]byte /* in6_addr */
	Scope_id uint32
}

type RawSockaddr struct {
	Family uint16
	Data   [14]int8
}

type RawSockaddrAny struct {
	Addr RawSockaddr
	Pad  [100]int8
}

type Sockaddr interface {
	sockaddr() (ptr unsafe.Pointer, len int32, err error) // 小写形式；只有我们能定义Sockaddrs
}

type SockaddrInet4 struct {
	Port int
	Addr [4]byte
	raw  RawSockaddrInet4
}

func (sa *SockaddrInet4) sockaddr() (unsafe.Pointer, int32, error) {
	if sa.Port < 0 || sa.Port > 0xFFFF {
		return nil, 0, syscall.EINVAL
	}
	sa.raw.Family = AF_INET
	p := (*[2]byte)(unsafe.Pointer(&sa.raw.Port))
	p[0] = byte(sa.Port >> 8)
	p[1] = byte(sa.Port)
	sa.raw.Addr = sa.Addr
	return unsafe.Pointer(&sa.raw), int32(unsafe.Sizeof(sa.raw)), nil
}

type SockaddrInet6 struct {
	Port   int
	ZoneId uint32
	Addr   [16]byte
	raw    RawSockaddrInet6
}

func (sa *SockaddrInet6) sockaddr() (unsafe.Pointer, int32, error) {
	if sa.Port < 0 || sa.Port > 0xFFFF {
		return nil, 0, syscall.EINVAL
	}
	sa.raw.Family = AF_INET6
	p := (*[2]byte)(unsafe.Pointer(&sa.raw.Port))
	p[0] = byte(sa.Port >> 8)
	p[1] = byte(sa.Port)
	sa.raw.Scope_id = sa.ZoneId
	sa.raw.Addr = sa.Addr
	return unsafe.Pointer(&sa.raw), int32(unsafe.Sizeof(sa.raw)), nil
}

type RawSockaddrUnix struct {
	Family uint16
	Path   [UNIX_PATH_MAX]int8
}

type SockaddrUnix struct {
	Name string
	raw  RawSockaddrUnix
}

func (sa *SockaddrUnix) sockaddr() (unsafe.Pointer, int32, error) {
	name := sa.Name
	n := len(name)
	if n > len(sa.raw.Path) {
		return nil, 0, syscall.EINVAL
	}
	if n == len(sa.raw.Path) && name[0] != '@' {
		return nil, 0, syscall.EINVAL
	}
	sa.raw.Family = AF_UNIX
	for i := 0; i < n; i++ {
		sa.raw.Path[i] = int8(name[i])
	}
	// length 包含家族（uint16 类型）、名称以及 NUL 字符
	sl := int32(2)
	if n > 0 {
		sl += int32(n) + 1
	}
	if sa.raw.Path[0] == '@' || (sa.raw.Path[0] == 0 && sl > 3) {
		// 检查 sl 是否大于 3，以避免改变无名套接字的行为
		sa.raw.Path[0] = 0
		// 对于抽象地址，不计算尾部的空字符（NUL）
		sl--
	}

	return unsafe.Pointer(&sa.raw), sl, nil
}

type RawSockaddrBth struct {
	AddressFamily  [2]byte
	BtAddr         [8]byte
	ServiceClassId [16]byte
	Port           [4]byte
}

type SockaddrBth struct {
	BtAddr         uint64
	ServiceClassId GUID
	Port           uint32

	raw RawSockaddrBth
}

func (sa *SockaddrBth) sockaddr() (unsafe.Pointer, int32, error) {
	family := AF_BTH
	sa.raw = RawSockaddrBth{
		AddressFamily:  *(*[2]byte)(unsafe.Pointer(&family)),
		BtAddr:         *(*[8]byte)(unsafe.Pointer(&sa.BtAddr)),
		Port:           *(*[4]byte)(unsafe.Pointer(&sa.Port)),
		ServiceClassId: *(*[16]byte)(unsafe.Pointer(&sa.ServiceClassId)),
	}
	return unsafe.Pointer(&sa.raw), int32(unsafe.Sizeof(sa.raw)), nil
}

func (rsa *RawSockaddrAny) Sockaddr() (Sockaddr, error) {
	switch rsa.Addr.Family {
	case AF_UNIX:
		pp := (*RawSockaddrUnix)(unsafe.Pointer(rsa))
		sa := new(SockaddrUnix)
		if pp.Path[0] == 0 {
// “抽象”Unix域套接字。
// 将起始空字符（NUL）重写为@以供文本显示。
// （这是标准约定。）
// 不适合在原地进行覆盖，
// 但下面的调用者并不关心这一点。
			pp.Path[0] = '@'
		}

// 假设路径以 NUL 结尾。
// 这并非严格遵循 Linux 对于抽象 Unix 域套接字的语义——它们本应被视为未经解释的固定大小二进制块——但大家普遍采用这种约定。
		n := 0
		for n < len(pp.Path) && pp.Path[n] != 0 {
			n++
		}
		sa.Name = string(unsafe.Slice((*byte)(unsafe.Pointer(&pp.Path[0])), n))
		return sa, nil

	case AF_INET:
		pp := (*RawSockaddrInet4)(unsafe.Pointer(rsa))
		sa := new(SockaddrInet4)
		p := (*[2]byte)(unsafe.Pointer(&pp.Port))
		sa.Port = int(p[0])<<8 + int(p[1])
		sa.Addr = pp.Addr
		return sa, nil

	case AF_INET6:
		pp := (*RawSockaddrInet6)(unsafe.Pointer(rsa))
		sa := new(SockaddrInet6)
		p := (*[2]byte)(unsafe.Pointer(&pp.Port))
		sa.Port = int(p[0])<<8 + int(p[1])
		sa.ZoneId = pp.Scope_id
		sa.Addr = pp.Addr
		return sa, nil
	}
	return nil, syscall.EAFNOSUPPORT
}

func Socket(domain, typ, proto int) (fd Handle, err error) {
	if domain == AF_INET6 && SocketDisableIPv6 {
		return InvalidHandle, syscall.EAFNOSUPPORT
	}
	return socket(int32(domain), int32(typ), int32(proto))
}

func SetsockoptInt(fd Handle, level, opt int, value int) (err error) {
	v := int32(value)
	return Setsockopt(fd, int32(level), int32(opt), (*byte)(unsafe.Pointer(&v)), int32(unsafe.Sizeof(v)))
}

func Bind(fd Handle, sa Sockaddr) (err error) {
	ptr, n, err := sa.sockaddr()
	if err != nil {
		return err
	}
	return bind(fd, ptr, n)
}

func Connect(fd Handle, sa Sockaddr) (err error) {
	ptr, n, err := sa.sockaddr()
	if err != nil {
		return err
	}
	return connect(fd, ptr, n)
}

func GetBestInterfaceEx(sa Sockaddr, pdwBestIfIndex *uint32) (err error) {
	ptr, _, err := sa.sockaddr()
	if err != nil {
		return err
	}
	return getBestInterfaceEx(ptr, pdwBestIfIndex)
}

func Getsockname(fd Handle) (sa Sockaddr, err error) {
	var rsa RawSockaddrAny
	l := int32(unsafe.Sizeof(rsa))
	if err = getsockname(fd, &rsa, &l); err != nil {
		return
	}
	return rsa.Sockaddr()
}

func Getpeername(fd Handle) (sa Sockaddr, err error) {
	var rsa RawSockaddrAny
	l := int32(unsafe.Sizeof(rsa))
	if err = getpeername(fd, &rsa, &l); err != nil {
		return
	}
	return rsa.Sockaddr()
}

func Listen(s Handle, n int) (err error) {
	return listen(s, int32(n))
}

func Shutdown(fd Handle, how int) (err error) {
	return shutdown(fd, int32(how))
}

func WSASendto(s Handle, bufs *WSABuf, bufcnt uint32, sent *uint32, flags uint32, to Sockaddr, overlapped *Overlapped, croutine *byte) (err error) {
	var rsa unsafe.Pointer
	var l int32
	if to != nil {
		rsa, l, err = to.sockaddr()
		if err != nil {
			return err
		}
	}
	return WSASendTo(s, bufs, bufcnt, sent, flags, (*RawSockaddrAny)(unsafe.Pointer(rsa)), l, overlapped, croutine)
}

func LoadGetAddrInfo() error {
	return procGetAddrInfoW.Find()
}

var connectExFunc struct {
	once sync.Once
	addr uintptr
	err  error
}

func LoadConnectEx() error {
	connectExFunc.once.Do(func() {
		var s Handle
		s, connectExFunc.err = Socket(AF_INET, SOCK_STREAM, IPPROTO_TCP)
		if connectExFunc.err != nil {
			return
		}
		defer CloseHandle(s)
		var n uint32
		connectExFunc.err = WSAIoctl(s,
			SIO_GET_EXTENSION_FUNCTION_POINTER,
			(*byte)(unsafe.Pointer(&WSAID_CONNECTEX)),
			uint32(unsafe.Sizeof(WSAID_CONNECTEX)),
			(*byte)(unsafe.Pointer(&connectExFunc.addr)),
			uint32(unsafe.Sizeof(connectExFunc.addr)),
			&n, nil, 0)
	})
	return connectExFunc.err
}

func connectEx(s Handle, name unsafe.Pointer, namelen int32, sendBuf *byte, sendDataLen uint32, bytesSent *uint32, overlapped *Overlapped) (err error) {
	r1, _, e1 := syscall.Syscall9(connectExFunc.addr, 7, uintptr(s), uintptr(name), uintptr(namelen), uintptr(unsafe.Pointer(sendBuf)), uintptr(sendDataLen), uintptr(unsafe.Pointer(bytesSent)), uintptr(unsafe.Pointer(overlapped)), 0, 0)
	if r1 == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func ConnectEx(fd Handle, sa Sockaddr, sendBuf *byte, sendDataLen uint32, bytesSent *uint32, overlapped *Overlapped) error {
	err := LoadConnectEx()
	if err != nil {
		return errorspkg.New("failed to find ConnectEx: " + err.Error())
	}
	ptr, n, err := sa.sockaddr()
	if err != nil {
		return err
	}
	return connectEx(fd, ptr, n, sendBuf, sendDataLen, bytesSent, overlapped)
}

var sendRecvMsgFunc struct {
	once     sync.Once
	sendAddr uintptr
	recvAddr uintptr
	err      error
}

func loadWSASendRecvMsg() error {
	sendRecvMsgFunc.once.Do(func() {
		var s Handle
		s, sendRecvMsgFunc.err = Socket(AF_INET, SOCK_DGRAM, IPPROTO_UDP)
		if sendRecvMsgFunc.err != nil {
			return
		}
		defer CloseHandle(s)
		var n uint32
		sendRecvMsgFunc.err = WSAIoctl(s,
			SIO_GET_EXTENSION_FUNCTION_POINTER,
			(*byte)(unsafe.Pointer(&WSAID_WSARECVMSG)),
			uint32(unsafe.Sizeof(WSAID_WSARECVMSG)),
			(*byte)(unsafe.Pointer(&sendRecvMsgFunc.recvAddr)),
			uint32(unsafe.Sizeof(sendRecvMsgFunc.recvAddr)),
			&n, nil, 0)
		if sendRecvMsgFunc.err != nil {
			return
		}
		sendRecvMsgFunc.err = WSAIoctl(s,
			SIO_GET_EXTENSION_FUNCTION_POINTER,
			(*byte)(unsafe.Pointer(&WSAID_WSASENDMSG)),
			uint32(unsafe.Sizeof(WSAID_WSASENDMSG)),
			(*byte)(unsafe.Pointer(&sendRecvMsgFunc.sendAddr)),
			uint32(unsafe.Sizeof(sendRecvMsgFunc.sendAddr)),
			&n, nil, 0)
	})
	return sendRecvMsgFunc.err
}

func WSASendMsg(fd Handle, msg *WSAMsg, flags uint32, bytesSent *uint32, overlapped *Overlapped, croutine *byte) error {
	err := loadWSASendRecvMsg()
	if err != nil {
		return err
	}
	r1, _, e1 := syscall.Syscall6(sendRecvMsgFunc.sendAddr, 6, uintptr(fd), uintptr(unsafe.Pointer(msg)), uintptr(flags), uintptr(unsafe.Pointer(bytesSent)), uintptr(unsafe.Pointer(overlapped)), uintptr(unsafe.Pointer(croutine)))
	if r1 == socket_error {
		err = errnoErr(e1)
	}
	return err
}

func WSARecvMsg(fd Handle, msg *WSAMsg, bytesReceived *uint32, overlapped *Overlapped, croutine *byte) error {
	err := loadWSASendRecvMsg()
	if err != nil {
		return err
	}
	r1, _, e1 := syscall.Syscall6(sendRecvMsgFunc.recvAddr, 5, uintptr(fd), uintptr(unsafe.Pointer(msg)), uintptr(unsafe.Pointer(bytesReceived)), uintptr(unsafe.Pointer(overlapped)), uintptr(unsafe.Pointer(croutine)), 0)
	if r1 == socket_error {
		err = errnoErr(e1)
	}
	return err
}

// 为满足package os所期望的内容，设计了如下结构体
type Rusage struct {
	CreationTime Filetime
	ExitTime     Filetime
	KernelTime   Filetime
	UserTime     Filetime
}

type WaitStatus struct {
	ExitCode uint32
}

func (w WaitStatus) Exited() bool { return true }

func (w WaitStatus) ExitStatus() int { return int(w.ExitCode) }

func (w WaitStatus) Signal() Signal { return -1 }

func (w WaitStatus) CoreDump() bool { return false }

func (w WaitStatus) Stopped() bool { return false }

func (w WaitStatus) Continued() bool { return false }

func (w WaitStatus) StopSignal() Signal { return -1 }

func (w WaitStatus) Signaled() bool { return false }

func (w WaitStatus) TrapCause() int { return -1 }

// Timespec 在 Windows 上是一个虚构的结构，但此处为了与其他操作系统对应的包保持一致性而存在。
type Timespec struct {
	Sec  int64
	Nsec int64
}

func TimespecToNsec(ts Timespec) int64 { return int64(ts.Sec)*1e9 + int64(ts.Nsec) }

func NsecToTimespec(nsec int64) (ts Timespec) {
	ts.Sec = nsec / 1e9
	ts.Nsec = nsec % 1e9
	return
}

// TODO(brainman): 修复net所需的全部内容

func Accept(fd Handle) (nfd Handle, sa Sockaddr, err error) { return 0, nil, syscall.EWINDOWS }

func Recvfrom(fd Handle, p []byte, flags int) (n int, from Sockaddr, err error) {
	var rsa RawSockaddrAny
	l := int32(unsafe.Sizeof(rsa))
	n32, err := recvfrom(fd, p, int32(flags), &rsa, &l)
	n = int(n32)
	if err != nil {
		return
	}
	from, err = rsa.Sockaddr()
	return
}

func Sendto(fd Handle, p []byte, flags int, to Sockaddr) (err error) {
	ptr, l, err := to.sockaddr()
	if err != nil {
		return err
	}
	return sendto(fd, p, int32(flags), ptr, l)
}

func SetsockoptTimeval(fd Handle, level, opt int, tv *Timeval) (err error) { return syscall.EWINDOWS }

// Linger 结构体是错误的，但我们直到 Go 1 发布后才注意到。
// sysLinger 是实际的系统调用结构体。

// BUG(brainman)：Linger 的定义并不适合直接与 Setsockopt 和 Getsockopt 一起使用。
// 请改用 SetsockoptLinger。

type Linger struct {
	Onoff  int32
	Linger int32
}

type sysLinger struct {
	Onoff  uint16
	Linger uint16
}

type IPMreq struct {
	Multiaddr [4]byte /* in_addr */
	Interface [4]byte /* in_addr */
}

type IPv6Mreq struct {
	Multiaddr [16]byte /* in6_addr */
	Interface uint32
}

func GetsockoptInt(fd Handle, level, opt int) (int, error) {
	v := int32(0)
	l := int32(unsafe.Sizeof(v))
	err := Getsockopt(fd, int32(level), int32(opt), (*byte)(unsafe.Pointer(&v)), &l)
	return int(v), err
}

func SetsockoptLinger(fd Handle, level, opt int, l *Linger) (err error) {
	sys := sysLinger{Onoff: uint16(l.Onoff), Linger: uint16(l.Linger)}
	return Setsockopt(fd, int32(level), int32(opt), (*byte)(unsafe.Pointer(&sys)), int32(unsafe.Sizeof(sys)))
}

func SetsockoptInet4Addr(fd Handle, level, opt int, value [4]byte) (err error) {
	return Setsockopt(fd, int32(level), int32(opt), (*byte)(unsafe.Pointer(&value[0])), 4)
}
func SetsockoptIPMreq(fd Handle, level, opt int, mreq *IPMreq) (err error) {
	return Setsockopt(fd, int32(level), int32(opt), (*byte)(unsafe.Pointer(mreq)), int32(unsafe.Sizeof(*mreq)))
}
func SetsockoptIPv6Mreq(fd Handle, level, opt int, mreq *IPv6Mreq) (err error) {
	return syscall.EWINDOWS
}

func EnumProcesses(processIds []uint32, bytesReturned *uint32) error {
// 函数EnumProcesses期望其size参数以字节为单位，但使用mksyscall生成的代码中，
// 实际使用的是processIds切片的长度。因此，添加此封装函数以修正这一差异。
	var p *uint32
	if len(processIds) > 0 {
		p = &processIds[0]
	}
	size := uint32(len(processIds) * 4)
	return enumProcesses(p, size, bytesReturned)
}

func Getpid() (pid int) { return int(GetCurrentProcessId()) }

func FindFirstFile(name *uint16, data *Win32finddata) (handle Handle, err error) {
// 注意(rsc): 对于系统调用，Win32finddata 结构体是错误的：
// 两个路径各自缺少一个 uint16。使用正确的结构体 win32finddata1，
// 然后将结果复制出来。这里没有表达力的损失，因为如果使用了最后一个 uint16，
// 它应该是空字符（NUL），而 Go 并不需要这个。对于 Go 1.1，
// 我们可能通过在该结构体中添加一个最后的 Bug [2]uint16 字段来避免在此处分配 win32finddata1，
// 然后直接调整结果中的字段。
	var data1 win32finddata1
	handle, err = findFirstFile1(name, &data1)
	if err == nil {
		copyFindData(data, &data1)
	}
	return
}

func FindNextFile(handle Handle, data *Win32finddata) (err error) {
	var data1 win32finddata1
	err = findNextFile1(handle, &data1)
	if err == nil {
		copyFindData(data, &data1)
	}
	return
}

func getProcessEntry(pid int) (*ProcessEntry32, error) {
	snapshot, err := CreateToolhelp32Snapshot(TH32CS_SNAPPROCESS, 0)
	if err != nil {
		return nil, err
	}
	defer CloseHandle(snapshot)
	var procEntry ProcessEntry32
	procEntry.Size = uint32(unsafe.Sizeof(procEntry))
	if err = Process32First(snapshot, &procEntry); err != nil {
		return nil, err
	}
	for {
		if procEntry.ProcessID == uint32(pid) {
			return &procEntry, nil
		}
		err = Process32Next(snapshot, &procEntry)
		if err != nil {
			return nil, err
		}
	}
}

func Getppid() (ppid int) {
	pe, err := getProcessEntry(Getpid())
	if err != nil {
		return -1
	}
	return int(pe.ParentProcessID)
}

// TODO(brainman): 修复os所需的全部内容
func Fchdir(fd Handle) (err error)             { return syscall.EWINDOWS }
func Link(oldpath, newpath string) (err error) { return syscall.EWINDOWS }
func Symlink(path, link string) (err error)    { return syscall.EWINDOWS }

func Fchmod(fd Handle, mode uint32) (err error)        { return syscall.EWINDOWS }
func Chown(path string, uid int, gid int) (err error)  { return syscall.EWINDOWS }
func Lchown(path string, uid int, gid int) (err error) { return syscall.EWINDOWS }
func Fchown(fd Handle, uid int, gid int) (err error)   { return syscall.EWINDOWS }

func Getuid() (uid int)                  { return -1 }
func Geteuid() (euid int)                { return -1 }
func Getgid() (gid int)                  { return -1 }
func Getegid() (egid int)                { return -1 }
func Getgroups() (gids []int, err error) { return nil, syscall.EWINDOWS }

type Signal int

func (s Signal) Signal() {}

func (s Signal) String() string {
	if 0 <= s && int(s) < len(signals) {
		str := signals[s]
		if str != "" {
			return str
		}
	}
	return "signal " + itoa(int(s))
}

func LoadCreateSymbolicLink() error {
	return procCreateSymbolicLinkW.Find()
}

// Readlink返回指定符号链接的目标。
func Readlink(path string, buf []byte) (n int, err error) {
	fd, err := CreateFile(StringToUTF16Ptr(path), GENERIC_READ, 0, nil, OPEN_EXISTING,
		FILE_FLAG_OPEN_REPARSE_POINT|FILE_FLAG_BACKUP_SEMANTICS, 0)
	if err != nil {
		return -1, err
	}
	defer CloseHandle(fd)

	rdbbuf := make([]byte, MAXIMUM_REPARSE_DATA_BUFFER_SIZE)
	var bytesReturned uint32
	err = DeviceIoControl(fd, FSCTL_GET_REPARSE_POINT, nil, 0, &rdbbuf[0], uint32(len(rdbbuf)), &bytesReturned, nil)
	if err != nil {
		return -1, err
	}

	rdb := (*reparseDataBuffer)(unsafe.Pointer(&rdbbuf[0]))
	var s string
	switch rdb.ReparseTag {
	case IO_REPARSE_TAG_SYMLINK:
		data := (*symbolicLinkReparseBuffer)(unsafe.Pointer(&rdb.reparseBuffer))
		p := (*[0xffff]uint16)(unsafe.Pointer(&data.PathBuffer[0]))
		s = UTF16ToString(p[data.PrintNameOffset/2 : (data.PrintNameLength-data.PrintNameOffset)/2])
	case IO_REPARSE_TAG_MOUNT_POINT:
		data := (*mountPointReparseBuffer)(unsafe.Pointer(&rdb.reparseBuffer))
		p := (*[0xffff]uint16)(unsafe.Pointer(&data.PathBuffer[0]))
		s = UTF16ToString(p[data.PrintNameOffset/2 : (data.PrintNameLength-data.PrintNameOffset)/2])
	default:
// 路径并非符号链接或junction（硬链接），而是一种其他类型的重解析点
		return -1, syscall.ENOENT
	}
	n = copy(buf, []byte(s))

	return n, nil
}

// GUIDFromString 将形如
// "{XXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX}" 的字符串解析为一个 GUID。
func GUIDFromString(str string) (GUID, error) {
	guid := GUID{}
	str16, err := syscall.UTF16PtrFromString(str)
	if err != nil {
		return guid, err
	}
	err = clsidFromString(str16, &guid)
	if err != nil {
		return guid, err
	}
	return guid, nil
}

// GenerateGUID 生成一个新的随机 GUID。
func GenerateGUID() (GUID, error) {
	guid := GUID{}
	err := coCreateGuid(&guid)
	if err != nil {
		return guid, err
	}
	return guid, nil
}

// String 方法返回 GUID 的规范字符串形式，
// 其格式为 "{XXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX}"。
func (guid GUID) String() string {
	var str [100]uint16
	chars := stringFromGUID2(&guid, &str[0], int32(len(str)))
	if chars <= 1 {
		return ""
	}
	return string(utf16.Decode(str[:chars-1]))
}

// KnownFolderPath 返回当前用户指定 FOLDERID_ 常量所表示的已知文件夹路径，该路径基于 KF_ 标志进行选择，并可选地根据该标志创建。
func KnownFolderPath(folderID *KNOWNFOLDERID, flags uint32) (string, error) {
	return Token(0).KnownFolderPath(folderID, flags)
}

// KnownFolderPath 返回用户令牌中由 FOLDERID_ 常量指定的一个已知文件夹路径，该路径基于 KF_ 标志的选择和可选创建。
func (t Token) KnownFolderPath(folderID *KNOWNFOLDERID, flags uint32) (string, error) {
	var p *uint16
	err := shGetKnownFolderPath(folderID, flags, t, &p)
	if err != nil {
		return "", err
	}
	defer CoTaskMemFree(unsafe.Pointer(p))
	return UTF16PtrToString(p), nil
}

// RtlGetVersion 返回底层操作系统版本，忽略清单语义但受应用程序兼容性层影响。
func RtlGetVersion() *OsVersionInfoEx {
	info := &OsVersionInfoEx{}
	info.osVersionInfoSize = uint32(unsafe.Sizeof(*info))
// 根据文档，此函数总是能成功。
// 该函数甚至不检查osVersionInfoSize成员的有效性。
// 对ntdll.dll进行反汇编表明，
// 文档在这一点上的描述确实正确。
	_ = rtlGetVersion(info)
	return info
}

// RtlGetNtVersionNumbers 返回底层操作系统版本，
// 忽略清单语义及应用程序兼容性层。
func RtlGetNtVersionNumbers() (majorVersion, minorVersion, buildNumber uint32) {
	rtlGetNtVersionNumbers(&majorVersion, &minorVersion, &buildNumber)
	buildNumber &= 0xffff
	return
}

// 获取进程首选的UI语言
func GetProcessPreferredUILanguages(flags uint32) ([]string, error) {
	return getUILanguages(flags, getProcessPreferredUILanguages)
}

// GetThreadPreferredUILanguages 获取当前线程的线程首选 UI 语言。
func GetThreadPreferredUILanguages(flags uint32) ([]string, error) {
	return getUILanguages(flags, getThreadPreferredUILanguages)
}

// GetUserPreferredUILanguages 获取用户首选的UI语言信息
func GetUserPreferredUILanguages(flags uint32) ([]string, error) {
	return getUILanguages(flags, getUserPreferredUILanguages)
}

// 获取系统首选的UI语言
func GetSystemPreferredUILanguages(flags uint32) ([]string, error) {
	return getUILanguages(flags, getSystemPreferredUILanguages)
}

func getUILanguages(flags uint32, f func(flags uint32, numLanguages *uint32, buf *uint16, bufSize *uint32) error) ([]string, error) {
	size := uint32(128)
	for {
		var numLanguages uint32
		buf := make([]uint16, size)
		err := f(flags, &numLanguages, &buf[0], &size)
		if err == ERROR_INSUFFICIENT_BUFFER {
			continue
		}
		if err != nil {
			return nil, err
		}
		buf = buf[:size]
		if numLanguages == 0 || len(buf) == 0 { // GetProcessPreferredUILanguages 可能返回 numLanguages==0 且内容为 "\0\0" 的情况
			return []string{}, nil
		}
		if buf[len(buf)-1] == 0 {
			buf = buf[:len(buf)-1] // 移除终止空字符
		}
		languages := make([]string, 0, numLanguages)
		from := 0
		for i, c := range buf {
			if c == 0 {
				languages = append(languages, string(utf16.Decode(buf[from:i])))
				from = i + 1
			}
		}
		return languages, nil
	}
}

func SetConsoleCursorPosition(console Handle, position Coord) error {
	return setConsoleCursorPosition(console, *((*uint32)(unsafe.Pointer(&position))))
}

func GetStartupInfo(startupInfo *StartupInfo) error {
	getStartupInfo(startupInfo)
	return nil
}

func (s NTStatus) Errno() syscall.Errno {
	return rtlNtStatusToDosErrorNoTeb(s)
}

func langID(pri, sub uint16) uint32 { return uint32(sub)<<10 | uint32(pri) }

func (s NTStatus) Error() string {
	b := make([]uint16, 300)
	n, err := FormatMessage(FORMAT_MESSAGE_FROM_SYSTEM|FORMAT_MESSAGE_FROM_HMODULE|FORMAT_MESSAGE_ARGUMENT_ARRAY, modntdll.Handle(), uint32(s), langID(LANG_ENGLISH, SUBLANG_ENGLISH_US), b, nil)
	if err != nil {
		return fmt.Sprintf("NTSTATUS 0x%08x", uint32(s))
	}
	// 去除结尾的\r和\n
	for ; n > 0 && (b[n-1] == '\n' || b[n-1] == '\r'); n-- {
	}
	return string(utf16.Decode(b[:n]))
}

// NewNTUnicodeString 用于返回一个新创建的 NTUnicodeString 结构，该结构适用于直接操作 NTUnicodeString 类型的原生 NT API。需要注意的是，大多数 Windows API 并不使用 NTUnicodeString 类型；对于更为常见的 *uint16 字符串类型，应使用 UTF16PtrFromString 函数。
func NewNTUnicodeString(s string) (*NTUnicodeString, error) {
	var u NTUnicodeString
	s16, err := UTF16PtrFromString(s)
	if err != nil {
		return nil, err
	}
	RtlInitUnicodeString(&u, s16)
	return &u, nil
}

// Slice 返回一个uint16切片，该切片别名引用NTUnicodeString中的数据。
func (s *NTUnicodeString) Slice() []uint16 {
	slice := unsafe.Slice(s.Buffer, s.MaximumLength)
	return slice[:s.Length]
}

func (s *NTUnicodeString) String() string {
	return UTF16ToString(s.Slice())
}

// NewNTString 函数用于创建一个新的 NTString 结构，以便与使用 NTString 类型的原生 NT API 一起工作。请注意，大多数 Windows API 并不使用 NTString，对于更为常见的 *uint16 字符串类型，应使用 UTF16PtrFromString 函数。
func NewNTString(s string) (*NTString, error) {
	var nts NTString
	s8, err := BytePtrFromString(s)
	if err != nil {
		return nil, err
	}
	RtlInitString(&nts, s8)
	return &nts, nil
}

// Slice 返回一个字节切片，该切片对NTString中的数据进行别名引用。
func (s *NTString) Slice() []byte {
	slice := unsafe.Slice(s.Buffer, s.MaximumLength)
	return slice[:s.Length]
}

func (s *NTString) String() string {
	return ByteSliceToString(s.Slice())
}

// FindResource 解析给定名称和资源类型的资源。
func FindResource(module Handle, name, resType ResourceIDOrString) (Handle, error) {
	var namePtr, resTypePtr uintptr
	var name16, resType16 *uint16
	var err error
	resolvePtr := func(i interface{}, keep **uint16) (uintptr, error) {
		switch v := i.(type) {
		case string:
			*keep, err = UTF16PtrFromString(v)
			if err != nil {
				return 0, err
			}
			return uintptr(unsafe.Pointer(*keep)), nil
		case ResourceID:
			return uintptr(v), nil
		}
		return 0, errorspkg.New("parameter must be a ResourceID or a string")
	}
	namePtr, err = resolvePtr(name, &name16)
	if err != nil {
		return 0, err
	}
	resTypePtr, err = resolvePtr(resType, &resType16)
	if err != nil {
		return 0, err
	}
	resInfo, err := findResource(module, namePtr, resTypePtr)
	runtime.KeepAlive(name16)
	runtime.KeepAlive(resType16)
	return resInfo, err
}

func LoadResourceData(module, resInfo Handle) (data []byte, err error) {
	size, err := SizeofResource(module, resInfo)
	if err != nil {
		return
	}
	resData, err := LoadResource(module, resInfo)
	if err != nil {
		return
	}
	ptr, err := LockResource(resData)
	if err != nil {
		return
	}
	data = unsafe.Slice((*byte)(unsafe.Pointer(ptr)), size)
	return
}

// PSAPI_WORKING_SET_EX_BLOCK 结构包含一页的扩展工作集信息。
type PSAPI_WORKING_SET_EX_BLOCK uint64

// Valid 返回此页面的有效性。
// 若该位为1，则后续成员有效；否则应忽略它们。
func (b PSAPI_WORKING_SET_EX_BLOCK) Valid() bool {
	return (b & 1) == 1
}

// ShareCount 表示共享此页面的进程数量。该成员的最大值为 7。
func (b PSAPI_WORKING_SET_EX_BLOCK) ShareCount() uint64 {
	return b.intField(1, 3)
}

// Win32Protection 表示页面的内存保护属性。有关值列表，请参阅
// https://docs.microsoft.com/zh-cn/windows/win32/memory/memory-protection-constants
func (b PSAPI_WORKING_SET_EX_BLOCK) Win32Protection() uint64 {
	return b.intField(4, 11)
}

// Shared 返回此页面的共享状态。
// 若该位为1，则表示该页面可被共享。
func (b PSAPI_WORKING_SET_EX_BLOCK) Shared() bool {
	return (b & (1 << 15)) == 1
}

// Node 代表 NUMA（非均匀内存访问）节点。该成员的最大值为 63。
func (b PSAPI_WORKING_SET_EX_BLOCK) Node() uint64 {
	return b.intField(16, 6)
}

// Locked 返回此页面的锁定状态。
// 若该位为1，则表示虚拟页面被锁定在物理内存中。
func (b PSAPI_WORKING_SET_EX_BLOCK) Locked() bool {
	return (b & (1 << 22)) == 1
}

// LargePage 返回该页面的大页状态。
// 若此位为1，则表示该页面为大页。
func (b PSAPI_WORKING_SET_EX_BLOCK) LargePage() bool {
	return (b & (1 << 23)) == 1
}

// Bad 返回此页面的错误状态。
// 若该位为1，则表示该页面已被报告为错误状态。
func (b PSAPI_WORKING_SET_EX_BLOCK) Bad() bool {
	return (b & (1 << 31)) == 1
}

// intField 从 PSAPI_WORKING_SET_EX_BLOCK 联合体中提取一个整型字段
func (b PSAPI_WORKING_SET_EX_BLOCK) intField(start, length int) uint64 {
	var mask PSAPI_WORKING_SET_EX_BLOCK
	for pos := start; pos < start+length; pos++ {
		mask |= (1 << pos)
	}

	masked := b & mask
	return uint64(masked >> start)
}

// PSAPI_WORKING_SET_EX_INFORMATION 结构包含进程的扩展工作集信息。
type PSAPI_WORKING_SET_EX_INFORMATION struct {
	// The virtual address.
	VirtualAddress Pointer
	// 一个PSAPI_WORKING_SET_EX_BLOCK联合体，用于指示VirtualAddress处页面的属性。
	VirtualAttributes PSAPI_WORKING_SET_EX_BLOCK
}

// 创建Windows伪控制台
func CreatePseudoConsole(size Coord, in Handle, out Handle, flags uint32, pconsole *Handle) error {
// 我们需要这个包装器来手动将 Coord 转换为 uint32。自动生成的包装器仅接受可以转换为 uintptr 的参数，而 Coord 无法做到这一点。
	return createPseudoConsole(*((*uint32)(unsafe.Pointer(&size))), in, out, flags, pconsole)
}

// ResizePseudoConsole 将伪控制台的内部缓冲区调整为`size`中指定的宽度和高度。
func ResizePseudoConsole(pconsole Handle, size Coord) error {
// 我们需要这个包装器来手动将 Coord 转换为 uint32。自动生成的包装器仅接受可以转换为 uintptr 的参数，而 Coord 无法做到这一点。
	return resizePseudoConsole(pconsole, *((*uint32)(unsafe.Pointer(&size))))
}

// DCB 常量。参见 https://learn.microsoft.com/en-us/windows/win32/api/winbase/ns-winbase-dcb.
const (
	CBR_110    = 110
	CBR_300    = 300
	CBR_600    = 600
	CBR_1200   = 1200
	CBR_2400   = 2400
	CBR_4800   = 4800
	CBR_9600   = 9600
	CBR_14400  = 14400
	CBR_19200  = 19200
	CBR_38400  = 38400
	CBR_57600  = 57600
	CBR_115200 = 115200
	CBR_128000 = 128000
	CBR_256000 = 256000

	DTR_CONTROL_DISABLE   = 0x00000000
	DTR_CONTROL_ENABLE    = 0x00000010
	DTR_CONTROL_HANDSHAKE = 0x00000020

	RTS_CONTROL_DISABLE   = 0x00000000
	RTS_CONTROL_ENABLE    = 0x00001000
	RTS_CONTROL_HANDSHAKE = 0x00002000
	RTS_CONTROL_TOGGLE    = 0x00003000

	NOPARITY    = 0
	ODDPARITY   = 1
	EVENPARITY  = 2
	MARKPARITY  = 3
	SPACEPARITY = 4

	ONESTOPBIT   = 0
	ONE5STOPBITS = 1
	TWOSTOPBITS  = 2
)

// EscapeCommFunction 常量。参见 https://learn.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-escapecommfunction.
const (
	SETXOFF  = 1
	SETXON   = 2
	SETRTS   = 3
	CLRRTS   = 4
	SETDTR   = 5
	CLRDTR   = 6
	SETBREAK = 8
	CLRBREAK = 9
)

// PurgeComm 常量。参见 https://learn.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-purgecomm。
const (
	PURGE_TXABORT = 0x0001
	PURGE_RXABORT = 0x0002
	PURGE_TXCLEAR = 0x0004
	PURGE_RXCLEAR = 0x0008
)

// SetCommMask 常量。参见 https://learn.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-setcommmask.
const (
	EV_RXCHAR  = 0x0001
	EV_RXFLAG  = 0x0002
	EV_TXEMPTY = 0x0004
	EV_CTS     = 0x0008
	EV_DSR     = 0x0010
	EV_RLSD    = 0x0020
	EV_BREAK   = 0x0040
	EV_ERR     = 0x0080
	EV_RING    = 0x0100
)
