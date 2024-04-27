// 代码由'go generate'命令生成；请勿编辑。

package windows

import (
	"github.com/888go/gosdk/os/internal/syscall/windows/sysdll"
	"syscall"
	"unsafe"
)

var _ unsafe.Pointer

// 只对常见的Errno值进行一次接口分配。
const (
	errnoERROR_IO_PENDING = 997
)

var (
	errERROR_IO_PENDING error = syscall.Errno(errnoERROR_IO_PENDING)
	errERROR_EINVAL     error = syscall.EINVAL
)

// errnoErr 返回常见的boxed Errno 值，以防止在运行时进行分配。
func errnoErr(e syscall.Errno) error {
	switch e {
	case 0:
		return errERROR_EINVAL
	case errnoERROR_IO_PENDING:
		return errERROR_IO_PENDING
	}
// TODO：在这里添加更多，收集完Windows上常见的错误值后（可能在运行all.bat时）？
	return e
}

var (
	modadvapi32         = syscall.NewLazyDLL(sysdll.Add("advapi32.dll"))
	modbcryptprimitives = syscall.NewLazyDLL(sysdll.Add("bcryptprimitives.dll"))
	modiphlpapi         = syscall.NewLazyDLL(sysdll.Add("iphlpapi.dll"))
	modkernel32         = syscall.NewLazyDLL(sysdll.Add("kernel32.dll"))
	modnetapi32         = syscall.NewLazyDLL(sysdll.Add("netapi32.dll"))
	modpsapi            = syscall.NewLazyDLL(sysdll.Add("psapi.dll"))
	moduserenv          = syscall.NewLazyDLL(sysdll.Add("userenv.dll"))
	modws2_32           = syscall.NewLazyDLL(sysdll.Add("ws2_32.dll"))

	procAdjustTokenPrivileges             = modadvapi32.NewProc("AdjustTokenPrivileges")
	procDuplicateTokenEx                  = modadvapi32.NewProc("DuplicateTokenEx")
	procImpersonateSelf                   = modadvapi32.NewProc("ImpersonateSelf")
	procLookupPrivilegeValueW             = modadvapi32.NewProc("LookupPrivilegeValueW")
	procOpenSCManagerW                    = modadvapi32.NewProc("OpenSCManagerW")
	procOpenServiceW                      = modadvapi32.NewProc("OpenServiceW")
	procOpenThreadToken                   = modadvapi32.NewProc("OpenThreadToken")
	procQueryServiceStatus                = modadvapi32.NewProc("QueryServiceStatus")
	procRevertToSelf                      = modadvapi32.NewProc("RevertToSelf")
	procSetTokenInformation               = modadvapi32.NewProc("SetTokenInformation")
	procProcessPrng                       = modbcryptprimitives.NewProc("ProcessPrng")
	procGetAdaptersAddresses              = modiphlpapi.NewProc("GetAdaptersAddresses")
	procCreateEventW                      = modkernel32.NewProc("CreateEventW")
	procGetACP                            = modkernel32.NewProc("GetACP")
	procGetComputerNameExW                = modkernel32.NewProc("GetComputerNameExW")
	procGetConsoleCP                      = modkernel32.NewProc("GetConsoleCP")
	procGetCurrentThread                  = modkernel32.NewProc("GetCurrentThread")
	procGetFileInformationByHandleEx      = modkernel32.NewProc("GetFileInformationByHandleEx")
	procGetFinalPathNameByHandleW         = modkernel32.NewProc("GetFinalPathNameByHandleW")
	procGetModuleFileNameW                = modkernel32.NewProc("GetModuleFileNameW")
	procGetTempPath2W                     = modkernel32.NewProc("GetTempPath2W")
	procGetVolumeInformationByHandleW     = modkernel32.NewProc("GetVolumeInformationByHandleW")
	procGetVolumeNameForVolumeMountPointW = modkernel32.NewProc("GetVolumeNameForVolumeMountPointW")
	procLockFileEx                        = modkernel32.NewProc("LockFileEx")
	procModule32FirstW                    = modkernel32.NewProc("Module32FirstW")
	procModule32NextW                     = modkernel32.NewProc("Module32NextW")
	procMoveFileExW                       = modkernel32.NewProc("MoveFileExW")
	procMultiByteToWideChar               = modkernel32.NewProc("MultiByteToWideChar")
	procRtlLookupFunctionEntry            = modkernel32.NewProc("RtlLookupFunctionEntry")
	procRtlVirtualUnwind                  = modkernel32.NewProc("RtlVirtualUnwind")
	procSetFileInformationByHandle        = modkernel32.NewProc("SetFileInformationByHandle")
	procUnlockFileEx                      = modkernel32.NewProc("UnlockFileEx")
	procVirtualQuery                      = modkernel32.NewProc("VirtualQuery")
	procNetShareAdd                       = modnetapi32.NewProc("NetShareAdd")
	procNetShareDel                       = modnetapi32.NewProc("NetShareDel")
	procNetUserGetLocalGroups             = modnetapi32.NewProc("NetUserGetLocalGroups")
	procGetProcessMemoryInfo              = modpsapi.NewProc("GetProcessMemoryInfo")
	procCreateEnvironmentBlock            = moduserenv.NewProc("CreateEnvironmentBlock")
	procDestroyEnvironmentBlock           = moduserenv.NewProc("DestroyEnvironmentBlock")
	procGetProfilesDirectoryW             = moduserenv.NewProc("GetProfilesDirectoryW")
	procWSASocketW                        = modws2_32.NewProc("WSASocketW")
)

func adjustTokenPrivileges(token syscall.Token, disableAllPrivileges bool, newstate *TOKEN_PRIVILEGES, buflen uint32, prevstate *TOKEN_PRIVILEGES, returnlen *uint32) (ret uint32, err error) {
	var _p0 uint32
	if disableAllPrivileges {
		_p0 = 1
	}
	r0, _, e1 := syscall.Syscall6(procAdjustTokenPrivileges.Addr(), 6, uintptr(token), uintptr(_p0), uintptr(unsafe.Pointer(newstate)), uintptr(buflen), uintptr(unsafe.Pointer(prevstate)), uintptr(unsafe.Pointer(returnlen)))
	ret = uint32(r0)
	if true {
		err = errnoErr(e1)
	}
	return
}


// ff:
// err:
// phNewToken:
// tokenType:
// impersonationLevel:
// lpTokenAttributes:
// dwDesiredAccess:
// hExistingToken:
func DuplicateTokenEx(hExistingToken syscall.Token, dwDesiredAccess uint32, lpTokenAttributes *syscall.SecurityAttributes, impersonationLevel uint32, tokenType TokenType, phNewToken *syscall.Token) (err error) {
	r1, _, e1 := syscall.Syscall6(procDuplicateTokenEx.Addr(), 6, uintptr(hExistingToken), uintptr(dwDesiredAccess), uintptr(unsafe.Pointer(lpTokenAttributes)), uintptr(impersonationLevel), uintptr(tokenType), uintptr(unsafe.Pointer(phNewToken)))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}


// ff:
// err:
// impersonationlevel:
func ImpersonateSelf(impersonationlevel uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procImpersonateSelf.Addr(), 1, uintptr(impersonationlevel), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}


// ff:
// err:
// luid:
// name:
// systemname:
func LookupPrivilegeValue(systemname *uint16, name *uint16, luid *LUID) (err error) {
	r1, _, e1 := syscall.Syscall(procLookupPrivilegeValueW.Addr(), 3, uintptr(unsafe.Pointer(systemname)), uintptr(unsafe.Pointer(name)), uintptr(unsafe.Pointer(luid)))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}


// ff:
// err:
// handle:
// access:
// databaseName:
// machineName:
func OpenSCManager(machineName *uint16, databaseName *uint16, access uint32) (handle syscall.Handle, err error) {
	r0, _, e1 := syscall.Syscall(procOpenSCManagerW.Addr(), 3, uintptr(unsafe.Pointer(machineName)), uintptr(unsafe.Pointer(databaseName)), uintptr(access))
	handle = syscall.Handle(r0)
	if handle == 0 {
		err = errnoErr(e1)
	}
	return
}


// ff:
// err:
// handle:
// access:
// serviceName:
// mgr:
func OpenService(mgr syscall.Handle, serviceName *uint16, access uint32) (handle syscall.Handle, err error) {
	r0, _, e1 := syscall.Syscall(procOpenServiceW.Addr(), 3, uintptr(mgr), uintptr(unsafe.Pointer(serviceName)), uintptr(access))
	handle = syscall.Handle(r0)
	if handle == 0 {
		err = errnoErr(e1)
	}
	return
}


// ff:
// err:
// token:
// openasself:
// access:
// h:
func OpenThreadToken(h syscall.Handle, access uint32, openasself bool, token *syscall.Token) (err error) {
	var _p0 uint32
	if openasself {
		_p0 = 1
	}
	r1, _, e1 := syscall.Syscall6(procOpenThreadToken.Addr(), 4, uintptr(h), uintptr(access), uintptr(_p0), uintptr(unsafe.Pointer(token)), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}


// ff:
// err:
// lpServiceStatus:
// hService:
func QueryServiceStatus(hService syscall.Handle, lpServiceStatus *SERVICE_STATUS) (err error) {
	r1, _, e1 := syscall.Syscall(procQueryServiceStatus.Addr(), 2, uintptr(hService), uintptr(unsafe.Pointer(lpServiceStatus)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}


// ff:
// err:
func RevertToSelf() (err error) {
	r1, _, e1 := syscall.Syscall(procRevertToSelf.Addr(), 0, 0, 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}


// ff:
// err:
// tokenInformationLength:
// tokenInformation:
// tokenInformationClass:
// tokenHandle:
func SetTokenInformation(tokenHandle syscall.Token, tokenInformationClass uint32, tokenInformation uintptr, tokenInformationLength uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procSetTokenInformation.Addr(), 4, uintptr(tokenHandle), uintptr(tokenInformationClass), uintptr(tokenInformation), uintptr(tokenInformationLength), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}


// ff:
// err:
// buf:
func ProcessPrng(buf []byte) (err error) {
	var _p0 *byte
	if len(buf) > 0 {
		_p0 = &buf[0]
	}
	r1, _, e1 := syscall.Syscall(procProcessPrng.Addr(), 2, uintptr(unsafe.Pointer(_p0)), uintptr(len(buf)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}


// ff:
// errcode:
// sizePointer:
// adapterAddresses:
// reserved:
// flags:
// family:
func GetAdaptersAddresses(family uint32, flags uint32, reserved uintptr, adapterAddresses *IpAdapterAddresses, sizePointer *uint32) (errcode error) {
	r0, _, _ := syscall.Syscall6(procGetAdaptersAddresses.Addr(), 5, uintptr(family), uintptr(flags), uintptr(reserved), uintptr(unsafe.Pointer(adapterAddresses)), uintptr(unsafe.Pointer(sizePointer)), 0)
	if r0 != 0 {
		errcode = syscall.Errno(r0)
	}
	return
}


// ff:
// err:
// handle:
// name:
// initialState:
// manualReset:
// eventAttrs:
func CreateEvent(eventAttrs *SecurityAttributes, manualReset uint32, initialState uint32, name *uint16) (handle syscall.Handle, err error) {
	r0, _, e1 := syscall.Syscall6(procCreateEventW.Addr(), 4, uintptr(unsafe.Pointer(eventAttrs)), uintptr(manualReset), uintptr(initialState), uintptr(unsafe.Pointer(name)), 0, 0)
	handle = syscall.Handle(r0)
	if handle == 0 {
		err = errnoErr(e1)
	}
	return
}


// ff:
// acp:
func GetACP() (acp uint32) {
	r0, _, _ := syscall.Syscall(procGetACP.Addr(), 0, 0, 0, 0)
	acp = uint32(r0)
	return
}


// ff:
// err:
// n:
// buf:
// nameformat:
func GetComputerNameEx(nameformat uint32, buf *uint16, n *uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procGetComputerNameExW.Addr(), 3, uintptr(nameformat), uintptr(unsafe.Pointer(buf)), uintptr(unsafe.Pointer(n)))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}


// ff:
// ccp:
func GetConsoleCP() (ccp uint32) {
	r0, _, _ := syscall.Syscall(procGetConsoleCP.Addr(), 0, 0, 0, 0)
	ccp = uint32(r0)
	return
}


// ff:
// err:
// pseudoHandle:
func GetCurrentThread() (pseudoHandle syscall.Handle, err error) {
	r0, _, e1 := syscall.Syscall(procGetCurrentThread.Addr(), 0, 0, 0, 0)
	pseudoHandle = syscall.Handle(r0)
	if pseudoHandle == 0 {
		err = errnoErr(e1)
	}
	return
}


// ff:
// err:
// bufsize:
// info:
// class:
// handle:
func GetFileInformationByHandleEx(handle syscall.Handle, class uint32, info *byte, bufsize uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procGetFileInformationByHandleEx.Addr(), 4, uintptr(handle), uintptr(class), uintptr(unsafe.Pointer(info)), uintptr(bufsize), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}


// ff:
// err:
// n:
// flags:
// filePathSize:
// filePath:
// file:
func GetFinalPathNameByHandle(file syscall.Handle, filePath *uint16, filePathSize uint32, flags uint32) (n uint32, err error) {
	r0, _, e1 := syscall.Syscall6(procGetFinalPathNameByHandleW.Addr(), 4, uintptr(file), uintptr(unsafe.Pointer(filePath)), uintptr(filePathSize), uintptr(flags), 0, 0)
	n = uint32(r0)
	if n == 0 {
		err = errnoErr(e1)
	}
	return
}


// ff:
// err:
// n:
// len:
// fn:
// module:
func GetModuleFileName(module syscall.Handle, fn *uint16, len uint32) (n uint32, err error) {
	r0, _, e1 := syscall.Syscall(procGetModuleFileNameW.Addr(), 3, uintptr(module), uintptr(unsafe.Pointer(fn)), uintptr(len))
	n = uint32(r0)
	if n == 0 {
		err = errnoErr(e1)
	}
	return
}


// ff:
// err:
// n:
// buf:
// buflen:
func GetTempPath2(buflen uint32, buf *uint16) (n uint32, err error) {
	r0, _, e1 := syscall.Syscall(procGetTempPath2W.Addr(), 2, uintptr(buflen), uintptr(unsafe.Pointer(buf)), 0)
	n = uint32(r0)
	if n == 0 {
		err = errnoErr(e1)
	}
	return
}


// ff:
// err:
// fileSystemNameSize:
// fileSystemNameBuffer:
// fileSystemFlags:
// maximumComponentLength:
// volumeNameSerialNumber:
// volumeNameSize:
// volumeNameBuffer:
// file:
func GetVolumeInformationByHandle(file syscall.Handle, volumeNameBuffer *uint16, volumeNameSize uint32, volumeNameSerialNumber *uint32, maximumComponentLength *uint32, fileSystemFlags *uint32, fileSystemNameBuffer *uint16, fileSystemNameSize uint32) (err error) {
	r1, _, e1 := syscall.Syscall9(procGetVolumeInformationByHandleW.Addr(), 8, uintptr(file), uintptr(unsafe.Pointer(volumeNameBuffer)), uintptr(volumeNameSize), uintptr(unsafe.Pointer(volumeNameSerialNumber)), uintptr(unsafe.Pointer(maximumComponentLength)), uintptr(unsafe.Pointer(fileSystemFlags)), uintptr(unsafe.Pointer(fileSystemNameBuffer)), uintptr(fileSystemNameSize), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}


// ff:
// err:
// bufferlength:
// volumeName:
// volumeMountPoint:
func GetVolumeNameForVolumeMountPoint(volumeMountPoint *uint16, volumeName *uint16, bufferlength uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procGetVolumeNameForVolumeMountPointW.Addr(), 3, uintptr(unsafe.Pointer(volumeMountPoint)), uintptr(unsafe.Pointer(volumeName)), uintptr(bufferlength))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}


// ff:
// err:
// overlapped:
// bytesHigh:
// bytesLow:
// reserved:
// flags:
// file:
func LockFileEx(file syscall.Handle, flags uint32, reserved uint32, bytesLow uint32, bytesHigh uint32, overlapped *syscall.Overlapped) (err error) {
	r1, _, e1 := syscall.Syscall6(procLockFileEx.Addr(), 6, uintptr(file), uintptr(flags), uintptr(reserved), uintptr(bytesLow), uintptr(bytesHigh), uintptr(unsafe.Pointer(overlapped)))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}


// ff:
// err:
// moduleEntry:
// snapshot:
func Module32First(snapshot syscall.Handle, moduleEntry *ModuleEntry32) (err error) {
	r1, _, e1 := syscall.Syscall(procModule32FirstW.Addr(), 2, uintptr(snapshot), uintptr(unsafe.Pointer(moduleEntry)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}


// ff:
// err:
// moduleEntry:
// snapshot:
func Module32Next(snapshot syscall.Handle, moduleEntry *ModuleEntry32) (err error) {
	r1, _, e1 := syscall.Syscall(procModule32NextW.Addr(), 2, uintptr(snapshot), uintptr(unsafe.Pointer(moduleEntry)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}


// ff:
// err:
// flags:
// to:
// from:
func MoveFileEx(from *uint16, to *uint16, flags uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procMoveFileExW.Addr(), 3, uintptr(unsafe.Pointer(from)), uintptr(unsafe.Pointer(to)), uintptr(flags))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}


// ff:
// err:
// nwrite:
// nwchar:
// wchar:
// nstr:
// str:
// dwFlags:
// codePage:
func MultiByteToWideChar(codePage uint32, dwFlags uint32, str *byte, nstr int32, wchar *uint16, nwchar int32) (nwrite int32, err error) {
	r0, _, e1 := syscall.Syscall6(procMultiByteToWideChar.Addr(), 6, uintptr(codePage), uintptr(dwFlags), uintptr(unsafe.Pointer(str)), uintptr(nstr), uintptr(unsafe.Pointer(wchar)), uintptr(nwchar))
	nwrite = int32(r0)
	if nwrite == 0 {
		err = errnoErr(e1)
	}
	return
}


// ff:
// ret:
// table:
// baseAddress:
// pc:
func RtlLookupFunctionEntry(pc uintptr, baseAddress *uintptr, table *byte) (ret uintptr) {
	r0, _, _ := syscall.Syscall(procRtlLookupFunctionEntry.Addr(), 3, uintptr(pc), uintptr(unsafe.Pointer(baseAddress)), uintptr(unsafe.Pointer(table)))
	ret = uintptr(r0)
	return
}


// ff:
// ret:
// ctxptrs:
// frame:
// data:
// ctxt:
// entry:
// pc:
// baseAddress:
// handlerType:
func RtlVirtualUnwind(handlerType uint32, baseAddress uintptr, pc uintptr, entry uintptr, ctxt uintptr, data *uintptr, frame *uintptr, ctxptrs *byte) (ret uintptr) {
	r0, _, _ := syscall.Syscall9(procRtlVirtualUnwind.Addr(), 8, uintptr(handlerType), uintptr(baseAddress), uintptr(pc), uintptr(entry), uintptr(ctxt), uintptr(unsafe.Pointer(data)), uintptr(unsafe.Pointer(frame)), uintptr(unsafe.Pointer(ctxptrs)), 0)
	ret = uintptr(r0)
	return
}


// ff:
// err:
// bufsize:
// buf:
// fileInformationClass:
// handle:
func SetFileInformationByHandle(handle syscall.Handle, fileInformationClass uint32, buf unsafe.Pointer, bufsize uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procSetFileInformationByHandle.Addr(), 4, uintptr(handle), uintptr(fileInformationClass), uintptr(buf), uintptr(bufsize), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}


// ff:
// err:
// overlapped:
// bytesHigh:
// bytesLow:
// reserved:
// file:
func UnlockFileEx(file syscall.Handle, reserved uint32, bytesLow uint32, bytesHigh uint32, overlapped *syscall.Overlapped) (err error) {
	r1, _, e1 := syscall.Syscall6(procUnlockFileEx.Addr(), 5, uintptr(file), uintptr(reserved), uintptr(bytesLow), uintptr(bytesHigh), uintptr(unsafe.Pointer(overlapped)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}


// ff:
// err:
// length:
// buffer:
// address:
func VirtualQuery(address uintptr, buffer *MemoryBasicInformation, length uintptr) (err error) {
	r1, _, e1 := syscall.Syscall(procVirtualQuery.Addr(), 3, uintptr(address), uintptr(unsafe.Pointer(buffer)), uintptr(length))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}


// ff:
// neterr:
// parmErr:
// buf:
// level:
// serverName:
func NetShareAdd(serverName *uint16, level uint32, buf *byte, parmErr *uint16) (neterr error) {
	r0, _, _ := syscall.Syscall6(procNetShareAdd.Addr(), 4, uintptr(unsafe.Pointer(serverName)), uintptr(level), uintptr(unsafe.Pointer(buf)), uintptr(unsafe.Pointer(parmErr)), 0, 0)
	if r0 != 0 {
		neterr = syscall.Errno(r0)
	}
	return
}


// ff:
// neterr:
// reserved:
// netName:
// serverName:
func NetShareDel(serverName *uint16, netName *uint16, reserved uint32) (neterr error) {
	r0, _, _ := syscall.Syscall(procNetShareDel.Addr(), 3, uintptr(unsafe.Pointer(serverName)), uintptr(unsafe.Pointer(netName)), uintptr(reserved))
	if r0 != 0 {
		neterr = syscall.Errno(r0)
	}
	return
}


// ff:
// neterr:
// totalEntries:
// entriesRead:
// prefMaxLen:
// buf:
// flags:
// level:
// userName:
// serverName:
func NetUserGetLocalGroups(serverName *uint16, userName *uint16, level uint32, flags uint32, buf **byte, prefMaxLen uint32, entriesRead *uint32, totalEntries *uint32) (neterr error) {
	r0, _, _ := syscall.Syscall9(procNetUserGetLocalGroups.Addr(), 8, uintptr(unsafe.Pointer(serverName)), uintptr(unsafe.Pointer(userName)), uintptr(level), uintptr(flags), uintptr(unsafe.Pointer(buf)), uintptr(prefMaxLen), uintptr(unsafe.Pointer(entriesRead)), uintptr(unsafe.Pointer(totalEntries)), 0)
	if r0 != 0 {
		neterr = syscall.Errno(r0)
	}
	return
}


// ff:
// err:
// cb:
// memCounters:
// handle:
func GetProcessMemoryInfo(handle syscall.Handle, memCounters *PROCESS_MEMORY_COUNTERS, cb uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procGetProcessMemoryInfo.Addr(), 3, uintptr(handle), uintptr(unsafe.Pointer(memCounters)), uintptr(cb))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}


// ff:
// err:
// inheritExisting:
// token:
// block:
func CreateEnvironmentBlock(block **uint16, token syscall.Token, inheritExisting bool) (err error) {
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


// ff:
// err:
// block:
func DestroyEnvironmentBlock(block *uint16) (err error) {
	r1, _, e1 := syscall.Syscall(procDestroyEnvironmentBlock.Addr(), 1, uintptr(unsafe.Pointer(block)), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}


// ff:
// err:
// dirLen:
// dir:
func GetProfilesDirectory(dir *uint16, dirLen *uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procGetProfilesDirectoryW.Addr(), 2, uintptr(unsafe.Pointer(dir)), uintptr(unsafe.Pointer(dirLen)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}


// ff:
// err:
// handle:
// flags:
// group:
// protinfo:
// protocol:
// typ:
// af:
func WSASocket(af int32, typ int32, protocol int32, protinfo *syscall.WSAProtocolInfo, group uint32, flags uint32) (handle syscall.Handle, err error) {
	r0, _, e1 := syscall.Syscall6(procWSASocketW.Addr(), 6, uintptr(af), uintptr(typ), uintptr(protocol), uintptr(unsafe.Pointer(protinfo)), uintptr(group), uintptr(flags))
	handle = syscall.Handle(r0)
	if handle == syscall.InvalidHandle {
		err = errnoErr(e1)
	}
	return
}
