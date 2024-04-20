
<原文开始>
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// 版权所有 ? 2009 Go作者。保留所有权利。
// 本源代码的使用受BSD风格
// 许可证约束，该许可证可在LICENSE文件中找到。
# <翻译结束>


<原文开始>
// Flags for DefineDosDevice.
<原文结束>

# <翻译开始>
// DefineDosDevice的标志
# <翻译结束>


<原文开始>
// Return values for GetDriveType.
<原文结束>

# <翻译开始>
// GetDriveType的返回值
# <翻译结束>


<原文开始>
// File system flags from GetVolumeInformation and GetVolumeInformationByHandle.
<原文结束>

# <翻译开始>
// 文件系统标志，来源于GetVolumeInformation和GetVolumeInformationByHandle函数。
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
// 若 s 中包含 NUL 字节，此函数将引发恐慌而非返回错误。
# <翻译结束>


<原文开始>
// UTF16FromString returns the UTF-16 encoding of the UTF-8 string
// s, with a terminating NUL added. If s contains a NUL byte at any
// location, it returns (nil, syscall.EINVAL).
<原文结束>

# <翻译开始>
// UTF16FromString返回UTF-8字符串s的UTF-16编码，并在其后添加一个终止空字符（NUL）。如果s在任何位置包含空字节（NUL），则返回(nil, syscall.EINVAL)。
# <翻译结束>


<原文开始>
// UTF16ToString returns the UTF-8 encoding of the UTF-16 sequence s,
// with a terminating NUL and any bytes after the NUL removed.
<原文结束>

# <翻译开始>
// UTF16ToString 将UTF-16序列s转换为其UTF-8编码，
// 并移除终止的NUL字符及该字符之后的所有字节。
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
// UTF16PtrFromString 返回UTF-8字符串s转换为的UTF-16编码的指针，并在末尾添加终止空字符（NUL）。如果s在任意位置包含空字节（NUL），则返回(nil, syscall.EINVAL)。
# <翻译结束>


<原文开始>
// UTF16PtrToString takes a pointer to a UTF-16 sequence and returns the corresponding UTF-8 encoded string.
// If the pointer is nil, it returns the empty string. It assumes that the UTF-16 sequence is terminated
// at a zero word; if the zero word is not present, the program may crash.
<原文结束>

# <翻译开始>
// UTF16PtrToString接收一个指向UTF-16序列的指针，并返回相应的UTF-8编码字符串。
// 若该指针为nil，则返回空字符串。该函数假定UTF-16序列以零字（word）作为终止符；若未出现零字，程序可能会崩溃。
# <翻译结束>


<原文开始>
// NewCallback converts a Go function to a function pointer conforming to the stdcall calling convention.
// This is useful when interoperating with Windows code requiring callbacks.
// The argument is expected to be a function with one uintptr-sized result. The function must not have arguments with size larger than the size of uintptr.
<原文结束>

# <翻译开始>
// NewCallback 将一个Go函数转换为符合stdcall调用约定的函数指针。
// 当与需要回调的Windows代码进行互操作时，这非常有用。
// 该参数应为具有一个uintptr大小结果的函数。该函数不得具有大于uintptr大小的参数。
# <翻译结束>


<原文开始>
// NewCallbackCDecl converts a Go function to a function pointer conforming to the cdecl calling convention.
// This is useful when interoperating with Windows code requiring callbacks.
// The argument is expected to be a function with one uintptr-sized result. The function must not have arguments with size larger than the size of uintptr.
<原文结束>

# <翻译开始>
// NewCallbackCDecl 将一个 Go 函数转换为遵循 cdecl 调用约定的函数指针。
// 这在与需要回调的 Windows 代码进行互操作时非常有用。
// 该参数应为具有一个 uintptr 大小结果的函数。此函数不得包含大小大于 uintptr 的参数。
# <翻译结束>


<原文开始>
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
<原文结束>

# <翻译开始>
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
# <翻译结束>


<原文开始>
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
<原文结束>

# <翻译开始>
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
# <翻译结束>


<原文开始>
// Version APIs
//sys	GetFileVersionInfoSize(filename string, zeroHandle *Handle) (bufSize uint32, err error) = version.GetFileVersionInfoSizeW
//sys	GetFileVersionInfo(filename string, handle uint32, bufSize uint32, buffer unsafe.Pointer) (err error) = version.GetFileVersionInfoW
//sys	VerQueryValue(block unsafe.Pointer, subBlock string, pointerToBufferPointer unsafe.Pointer, bufSize *uint32) (err error) = version.VerQueryValueW
<原文结束>

# <翻译开始>
// 版本相关API
//sys	GetFileVersionInfoSize(filename string, zeroHandle *Handle) (bufSize uint32, err error) = version.GetFileVersionInfoSizeW
//sys	GetFileVersionInfo(filename string, handle uint32, bufSize uint32, buffer unsafe.Pointer) (err error) = version.GetFileVersionInfoW
//sys	VerQueryValue(block unsafe.Pointer, subBlock string, pointerToBufferPointer unsafe.Pointer, bufSize *uint32) (err error) = version.VerQueryValueW
// 
// 翻译为中文：
// 
// 版本相关API
//sys	GetFileVersionInfoSize(filename string, zeroHandle *Handle) (bufSize uint32, err error) = version.GetFileVersionInfoSizeW
//sys	GetFileVersionInfo(filename string, handle uint32, bufSize uint32, buffer unsafe.Pointer) (err error) = version.GetFileVersionInfoW
//sys	VerQueryValue(block unsafe.Pointer, subBlock string, pointerToBufferPointer unsafe.Pointer, bufSize *uint32) (err error) = version.VerQueryValueW
# <翻译结束>


<原文开始>
// Process Status API (PSAPI)
//sys	enumProcesses(processIds *uint32, nSize uint32, bytesReturned *uint32) (err error) = psapi.EnumProcesses
//sys	EnumProcessModules(process Handle, module *Handle, cb uint32, cbNeeded *uint32) (err error) = psapi.EnumProcessModules
//sys	EnumProcessModulesEx(process Handle, module *Handle, cb uint32, cbNeeded *uint32, filterFlag uint32) (err error) = psapi.EnumProcessModulesEx
//sys	GetModuleInformation(process Handle, module Handle, modinfo *ModuleInfo, cb uint32) (err error) = psapi.GetModuleInformation
//sys	GetModuleFileNameEx(process Handle, module Handle, filename *uint16, size uint32) (err error) = psapi.GetModuleFileNameExW
//sys	GetModuleBaseName(process Handle, module Handle, baseName *uint16, size uint32) (err error) = psapi.GetModuleBaseNameW
//sys   QueryWorkingSetEx(process Handle, pv uintptr, cb uint32) (err error) = psapi.QueryWorkingSetEx
<原文结束>

# <翻译开始>
// 进程状态API（PSAPI）
//sys	enumProcesses(processIds *uint32, nSize uint32, bytesReturned *uint32) (err error) = psapi.EnumProcesses
//sys	EnumProcessModules(process Handle, module *Handle, cb uint32, cbNeeded *uint32) (err error) = psapi.EnumProcessModules
//sys	EnumProcessModulesEx(process Handle, module *Handle, cb uint32, cbNeeded *uint32, filterFlag uint32) (err error) = psapi.EnumProcessModulesEx
//sys	GetModuleInformation(process Handle, module Handle, modinfo *ModuleInfo, cb uint32) (err error) = psapi.GetModuleInformation
//sys	GetModuleFileNameEx(process Handle, module Handle, filename *uint16, size uint32) (err error) = psapi.GetModuleFileNameExW
//sys	GetModuleBaseName(process Handle, module Handle, baseName *uint16, size uint32) (err error) = psapi.GetModuleBaseNameW
//sys   QueryWorkingSetEx(process Handle, pv uintptr, cb uint32) (err error) = psapi.QueryWorkingSetEx
// 
// 进程状态API（PSAPI）
// 
//sys	enumProcesses：获取进程ID列表，将结果存储在processIds指向的缓冲区中，nSize为缓冲区大小。bytesReturned用于返回实际写入缓冲区的字节数。调用psapi.EnumProcesses函数实现。
//sys	EnumProcessModules：枚举指定进程的所有模块（DLL），将模块句柄存储在module指向的缓冲区中，cb表示缓冲区大小。cbNeeded用于返回实际需要的缓冲区大小。调用psapi.EnumProcessModules函数实现。
//sys	EnumProcessModulesEx：与EnumProcessModules类似，但额外接受一个filterFlag参数，用于指定模块枚举的筛选条件。调用psapi.EnumProcessModulesEx函数实现。
//sys	GetModuleInformation：获取指定进程和模块的相关信息，将其存入modinfo指向的ModuleInfo结构体中。cb表示ModuleInfo结构体大小。调用psapi.GetModuleInformation函数实现。
//sys	GetModuleFileNameEx：获取指定进程和模块的完整文件名，将其以Unicode字符形式存入filename指向的缓冲区中，size为缓冲区大小。调用psapi.GetModuleFileNameExW函数实现。
//sys	GetModuleBaseName：获取指定进程和模块的基本名称（不包含路径），将其以Unicode字符形式存入baseName指向的缓冲区中，size为缓冲区大小。调用psapi.GetModuleBaseNameW函数实现。
//sys	QueryWorkingSetEx：查询指定进程的工作集信息，pv参数为指向工作集数据的指针，cb为工作集数据大小。调用psapi.QueryWorkingSetEx函数实现。
# <翻译结束>


<原文开始>
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
<原文结束>

# <翻译开始>
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
# <翻译结束>


<原文开始>
// Desktop Window Manager API (Dwmapi)
//sys	DwmGetWindowAttribute(hwnd HWND, attribute uint32, value unsafe.Pointer, size uint32) (ret error) = dwmapi.DwmGetWindowAttribute
//sys	DwmSetWindowAttribute(hwnd HWND, attribute uint32, value unsafe.Pointer, size uint32) (ret error) = dwmapi.DwmSetWindowAttribute
<原文结束>

# <翻译开始>
// 桌面窗口管理器API（Dwmapi）
//sys	DwmGetWindowAttribute(hwnd HWND, attribute uint32, value unsafe.Pointer, size uint32) (ret error) = dwmapi.DwmGetWindowAttribute
//sys	DwmSetWindowAttribute(hwnd HWND, attribute uint32, value unsafe.Pointer, size uint32) (ret error) = dwmapi.DwmSetWindowAttribute
// 
// 桌面窗口管理器API（Dwmapi）
//sys	获取窗口属性（DwmGetWindowAttribute）：通过给定的hwnd（窗口句柄）、attribute（属性标识符）、value（指向接收属性值的指针）和size（属性值大小），获取指定窗口的相关属性，并返回错误信息。该函数对应于dwmapi库中的DwmGetWindowAttribute。
//sys	设置窗口属性（DwmSetWindowAttribute）：使用提供的hwnd（窗口句柄）、attribute（属性标识符）、value（指向属性值的指针）和size（属性值大小），为指定窗口设置相应属性，并返回错误信息。该函数实现调用了dwmapi库中的DwmSetWindowAttribute。
# <翻译结束>


<原文开始>
// Windows Multimedia API
//sys TimeBeginPeriod (period uint32) (err error) [failretval != 0] = winmm.timeBeginPeriod
//sys TimeEndPeriod (period uint32) (err error) [failretval != 0] = winmm.timeEndPeriod
<原文结束>

# <翻译开始>
// Windows 多媒体 API
//sys TimeBeginPeriod (period uint32) (err error) [failretval != 0] = winmm.timeBeginPeriod
//sys TimeEndPeriod (period uint32) (err error) [failretval != 0] = winmm.timeEndPeriod
// 
// 翻译成中文为：
// 
// Windows 多媒体 API
//sys TimeBeginPeriod(周期 uint32) (错误 error) [失败返回值 != 0] = winmm.timeBeginPeriod
//sys TimeEndPeriod(周期 uint32) (错误 error) [失败返回值 != 0] = winmm.timeEndPeriod
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
// 它是一个无需关闭的伪句柄。
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
// 已弃用：使用 CurrentThread 直接获取相同的句柄，且无 nil 错误。
# <翻译结束>


<原文开始>
// CurrentThread returns the handle for the current thread.
// It is a pseudo handle that does not need to be closed.
<原文结束>

# <翻译开始>
// CurrentThread 返回当前线程的句柄。
// 它是一个无需关闭的伪句柄。
# <翻译结束>


<原文开始>
// GetProcAddressByOrdinal retrieves the address of the exported
// function from module by ordinal.
<原文结束>

# <翻译开始>
// GetProcAddressByOrdinal 通过序数从模块中获取导出函数的地址。
# <翻译结束>


<原文开始>
// NOTE(brainman): work around ERROR_BROKEN_PIPE is returned on reading EOF from stdin
<原文结束>

# <翻译开始>
// 注意(brainman): 为解决从stdin读取EOF时返回ERROR_BROKEN_PIPE的问题
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
// 除本函数外，所有其他 win32 数组 API 都以“指针，计数”形式接收参数。因此，我们不能将其声明为常规的 [] 类型，因为 mksyscall 将使用相反的顺序。为此，我们自己简单地为此进行了存根处理。
# <翻译结束>


<原文开始>
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
<原文结束>

# <翻译开始>
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
// 全部小写；只有我们能定义Sockaddrs
# <翻译结束>


<原文开始>
// length is family (uint16), name, NUL.
<原文结束>

# <翻译开始>
// length 包含 family（uint16 类型）、name 以及 NUL。
# <翻译结束>


<原文开始>
// Check sl > 3 so we don't change unnamed socket behavior.
<原文结束>

# <翻译开始>
// 检查 sl 是否大于 3，以便我们不改变未命名套接字的行为。
# <翻译结束>


<原文开始>
// Don't count trailing NUL for abstract address.
<原文结束>

# <翻译开始>
// 对于抽象地址，不计算尾部的空字符（NUL）。
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
// 将前导空字符（NUL）重写为@，以供文本显示。
// （这是标准约定。）
// 不适合就地覆盖，
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
// 假设路径以NUL结束。
// 这并非严格遵循Linux对于抽象Unix域套接字的语义——它们本应被视为未经解释的固定大小二进制块——但大家普遍使用这一约定。
# <翻译结束>


<原文开始>
// Invented structures to support what package os expects.
<原文结束>

# <翻译开始>
// 为满足package os的预期而设计的结构体
# <翻译结束>


<原文开始>
// Timespec is an invented structure on Windows, but here for
// consistency with the corresponding package for other operating systems.
<原文结束>

# <翻译开始>
// Timespec 是一个在 Windows 上虚构的结构，但此处为了与其它操作系统上对应的包保持一致性而存在。
# <翻译结束>


<原文开始>
// TODO(brainman): fix all needed for net
<原文结束>

# <翻译开始>
// TODO(brainman): 修复net所需的所有内容
# <翻译结束>


<原文开始>
// The Linger struct is wrong but we only noticed after Go 1.
// sysLinger is the real system call structure.
<原文结束>

# <翻译开始>
// Linger 结构体是错误的，但我们在 Go 1 发布后才注意到这一点。
// sysLinger 是实际的系统调用结构体。
# <翻译结束>


<原文开始>
// BUG(brainman): The definition of Linger is not appropriate for direct use
// with Setsockopt and Getsockopt.
// Use SetsockoptLinger instead.
<原文结束>

# <翻译开始>
// BUG(brainman): Linger 的定义并不适合直接用于 Setsockopt 和 Getsockopt。
// 请改用 SetsockoptLinger。
# <翻译结束>


<原文开始>
	// EnumProcesses syscall expects the size parameter to be in bytes, but the code generated with mksyscall uses
	// the length of the processIds slice instead. Hence, this wrapper function is added to fix the discrepancy.
<原文结束>

# <翻译开始>
// syscall.EnumProcesses 系统调用期望 size 参数以字节为单位，但使用 mksyscall 生成的代码中，
// 使用的是 processIds 切片的长度。因此，添加此封装函数以修正这一差异。
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
// 注意(rsc)：对于系统调用，Win32finddata 结构体是错误的：
// 两个路径各缺少一个 uint16。使用正确的结构体 win32finddata1，
// 然后将结果复制出来。这样做不会损失表达力，因为如果使用到最后一个 uint16，
// 它应该是空字符（NUL），而 Go 并不需要这个。对于 Go 1.1 版本，
// 我们可以通过在结构体中添加一个最后的 Bug [2]uint16 字段来避免在此处分配 win32finddata1，
// 然后直接调整结果中的字段。
# <翻译结束>


<原文开始>
// TODO(brainman): fix all needed for os
<原文结束>

# <翻译开始>
// TODO(brainman): 修复os所需的所有内容
# <翻译结束>


<原文开始>
// Readlink returns the destination of the named symbolic link.
<原文结束>

# <翻译开始>
// Readlink返回指定符号链接的目标路径。
# <翻译结束>


<原文开始>
		// the path is not a symlink or junction but another type of reparse
		// point
<原文结束>

# <翻译开始>
// 路径并非符号链接或junction（链接点），而是另一种类型的重分析点
# <翻译结束>


<原文开始>
// GUIDFromString parses a string in the form of
// "{XXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX}" into a GUID.
<原文结束>

# <翻译开始>
// GUIDFromString 将形如
// "{XXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX}" 的字符串解析为 GUID。
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
// String 方法返回该GUID的标准字符串形式，
// 其格式为"{XXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX}"。
# <翻译结束>


<原文开始>
// KnownFolderPath returns a well-known folder path for the current user, specified by one of
// the FOLDERID_ constants, and chosen and optionally created based on a KF_ flag.
<原文结束>

# <翻译开始>
// KnownFolderPath 返回当前用户指定的某个已知文件夹路径，该路径由 FOLDERID_ 常量中的一个表示，并根据 KF_ 标志进行选择和可选创建。
# <翻译结束>


<原文开始>
// KnownFolderPath returns a well-known folder path for the user token, specified by one of
// the FOLDERID_ constants, and chosen and optionally created based on a KF_ flag.
<原文结束>

# <翻译开始>
// KnownFolderPath 返回一个与用户令牌关联的已知文件夹路径，该路径由某个 FOLDERID_ 常量指定，并根据一个 KF_ 标志进行选择和可选创建。
# <翻译结束>


<原文开始>
// RtlGetVersion returns the version of the underlying operating system, ignoring
// manifest semantics but is affected by the application compatibility layer.
<原文结束>

# <翻译开始>
// RtlGetVersion 返回底层操作系统的版本，忽略清单语义但受应用程序兼容性层影响。
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
// 文档在这一点上的描述确实是正确的。
# <翻译结束>


<原文开始>
// RtlGetNtVersionNumbers returns the version of the underlying operating system,
// ignoring manifest semantics and the application compatibility layer.
<原文结束>

# <翻译开始>
// RtlGetNtVersionNumbers 返回底层操作系统的版本信息，
// 在此过程中忽略清单语义及应用程序兼容性层。
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
// GetThreadPreferredUILanguages 获取当前线程的线程首选用户界面语言。
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
// GetProcessPreferredUILanguages 可能返回 numLanguages==0 以及 "\0\0"
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
// 去除结尾的 \r 和 \n
# <翻译结束>


<原文开始>
// NewNTUnicodeString returns a new NTUnicodeString structure for use with native
// NT APIs that work over the NTUnicodeString type. Note that most Windows APIs
// do not use NTUnicodeString, and instead UTF16PtrFromString should be used for
// the more common *uint16 string type.
<原文结束>

# <翻译开始>
// NewNTUnicodeString 返回一个用于与直接操作 NTUnicodeString 类型的原生 NT API 交互的新 NTUnicodeString 结构。请注意，大多数 Windows API 并不使用 NTUnicodeString，对于更为常见的 *uint16 字符串类型，应使用 UTF16PtrFromString。
# <翻译结束>


<原文开始>
// Slice returns a uint16 slice that aliases the data in the NTUnicodeString.
<原文结束>

# <翻译开始>
// Slice 返回一个 uint16 类型的切片，该切片与 NTUnicodeString 中的数据共用同一存储空间。
# <翻译结束>


<原文开始>
// NewNTString returns a new NTString structure for use with native
// NT APIs that work over the NTString type. Note that most Windows APIs
// do not use NTString, and instead UTF16PtrFromString should be used for
// the more common *uint16 string type.
<原文结束>

# <翻译开始>
// NewNTString 函数用于返回一个供与原生 NT API 交互的新的 NTString 结构体。
// 注意，大多数 Windows API 并不使用 NTString 类型，对于更为常见的 *uint16 字符串类型，
// 应该使用 UTF16PtrFromString 函数。
# <翻译结束>


<原文开始>
// Slice returns a byte slice that aliases the data in the NTString.
<原文结束>

# <翻译开始>
// Slice 返回一个字节切片，该切片别名引用NTString中的数据。
# <翻译结束>


<原文开始>
// FindResource resolves a resource of the given name and resource type.
<原文结束>

# <翻译开始>
// FindResource 用于解析指定名称和资源类型的资源。
# <翻译结束>


<原文开始>
// PSAPI_WORKING_SET_EX_BLOCK contains extended working set information for a page.
<原文结束>

# <翻译开始>
// PSAPI_WORKING_SET_EX_BLOCK 包含页面的扩展工作集信息。
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
// ShareCount 是共享此页面的进程数量。该成员的最大值为 7。
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
// Node 代表 NUMA（非统一内存访问）节点。该成员的最大值为 63。
# <翻译结束>


<原文开始>
// Locked returns the locked status of this page.
// If this bit is 1, the virtual page is locked in physical memory.
<原文结束>

# <翻译开始>
// Locked 返回此页面的锁定状态。
// 若该位为1，则表示虚拟页已被锁定在物理内存中。
# <翻译结束>


<原文开始>
// LargePage returns the large page status of this page.
// If this bit is 1, the page is a large page.
<原文结束>

# <翻译开始>
// LargePage 返回该页面的大页状态。
// 如果该位为1，则表示该页面为大页。
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
// 我们需要这个包装器来手动将 Coord 转换为 uint32。自动生成的包装器仅接受可转换为 uintptr 的参数，而 Coord 无法做到这一点。
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
// EscapeCommFunction 常量。参考 https://learn.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-escapecommfunction.
# <翻译结束>


<原文开始>
// PurgeComm constants. See https://learn.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-purgecomm.
<原文结束>

# <翻译开始>
// PurgeComm 常量。参见 https://learn.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-purgecomm.
# <翻译结束>


<原文开始>
// SetCommMask constants. See https://learn.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-setcommmask.
<原文结束>

# <翻译开始>
// SetCommMask 常量。参考 https://learn.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-setcommmask.
# <翻译结束>

