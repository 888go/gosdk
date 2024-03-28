
<原文开始>
// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
# <翻译结束>


<原文开始>
// UTF16PtrToString is like UTF16ToString, but takes *uint16
// as a parameter instead of []uint16.
<原文结束>

# <翻译开始>
// UTF16PtrToString is like UTF16ToString, but takes *uint16
// as a parameter instead of []uint16.
# <翻译结束>


<原文开始>
// Turn *uint16 into []uint16.
<原文结束>

# <翻译开始>
// Turn *uint16 into []uint16.
# <翻译结束>


<原文开始>
// Decode []uint16 into string.
<原文结束>

# <翻译开始>
// Decode []uint16 into string.
# <翻译结束>


<原文开始>
//sys	GetAdaptersAddresses(family uint32, flags uint32, reserved uintptr, adapterAddresses *IpAdapterAddresses, sizePointer *uint32) (errcode error) = iphlpapi.GetAdaptersAddresses
//sys	GetComputerNameEx(nameformat uint32, buf *uint16, n *uint32) (err error) = GetComputerNameExW
//sys	MoveFileEx(from *uint16, to *uint16, flags uint32) (err error) = MoveFileExW
//sys	GetModuleFileName(module syscall.Handle, fn *uint16, len uint32) (n uint32, err error) = kernel32.GetModuleFileNameW
//sys	SetFileInformationByHandle(handle syscall.Handle, fileInformationClass uint32, buf uintptr, bufsize uint32) (err error) = kernel32.SetFileInformationByHandle
//sys	VirtualQuery(address uintptr, buffer *MemoryBasicInformation, length uintptr) (err error) = kernel32.VirtualQuery
<原文结束>

# <翻译开始>
//sys	GetAdaptersAddresses(family uint32, flags uint32, reserved uintptr, adapterAddresses *IpAdapterAddresses, sizePointer *uint32) (errcode error) = iphlpapi.GetAdaptersAddresses
//sys	GetComputerNameEx(nameformat uint32, buf *uint16, n *uint32) (err error) = GetComputerNameExW
//sys	MoveFileEx(from *uint16, to *uint16, flags uint32) (err error) = MoveFileExW
//sys	GetModuleFileName(module syscall.Handle, fn *uint16, len uint32) (n uint32, err error) = kernel32.GetModuleFileNameW
//sys	SetFileInformationByHandle(handle syscall.Handle, fileInformationClass uint32, buf uintptr, bufsize uint32) (err error) = kernel32.SetFileInformationByHandle
//sys	VirtualQuery(address uintptr, buffer *MemoryBasicInformation, length uintptr) (err error) = kernel32.VirtualQuery
# <翻译结束>


<原文开始>
// flags for CreateToolhelp32Snapshot
<原文结束>

# <翻译开始>
// flags for CreateToolhelp32Snapshot
# <翻译结束>


<原文开始>
//sys	Module32First(snapshot syscall.Handle, moduleEntry *ModuleEntry32) (err error) = kernel32.Module32FirstW
//sys	Module32Next(snapshot syscall.Handle, moduleEntry *ModuleEntry32) (err error) = kernel32.Module32NextW
<原文结束>

# <翻译开始>
//sys	Module32First(snapshot syscall.Handle, moduleEntry *ModuleEntry32) (err error) = kernel32.Module32FirstW
//sys	Module32Next(snapshot syscall.Handle, moduleEntry *ModuleEntry32) (err error) = kernel32.Module32NextW
# <翻译结束>


<原文开始>
//sys	WSASocket(af int32, typ int32, protocol int32, protinfo *syscall.WSAProtocolInfo, group uint32, flags uint32) (handle syscall.Handle, err error) [failretval==syscall.InvalidHandle] = ws2_32.WSASocketW
<原文结束>

# <翻译开始>
//sys	WSASocket(af int32, typ int32, protocol int32, protinfo *syscall.WSAProtocolInfo, group uint32, flags uint32) (handle syscall.Handle, err error) [failretval==syscall.InvalidHandle] = ws2_32.WSASocketW
# <翻译结束>


<原文开始>
//sys LockFileEx(file syscall.Handle, flags uint32, reserved uint32, bytesLow uint32, bytesHigh uint32, overlapped *syscall.Overlapped) (err error) = kernel32.LockFileEx
//sys UnlockFileEx(file syscall.Handle, reserved uint32, bytesLow uint32, bytesHigh uint32, overlapped *syscall.Overlapped) (err error) = kernel32.UnlockFileEx
<原文结束>

# <翻译开始>
//sys LockFileEx(file syscall.Handle, flags uint32, reserved uint32, bytesLow uint32, bytesHigh uint32, overlapped *syscall.Overlapped) (err error) = kernel32.LockFileEx
//sys UnlockFileEx(file syscall.Handle, reserved uint32, bytesLow uint32, bytesHigh uint32, overlapped *syscall.Overlapped) (err error) = kernel32.UnlockFileEx
# <翻译结束>


<原文开始>
//sys	GetACP() (acp uint32) = kernel32.GetACP
//sys	GetConsoleCP() (ccp uint32) = kernel32.GetConsoleCP
//sys	MultiByteToWideChar(codePage uint32, dwFlags uint32, str *byte, nstr int32, wchar *uint16, nwchar int32) (nwrite int32, err error) = kernel32.MultiByteToWideChar
//sys	GetCurrentThread() (pseudoHandle syscall.Handle, err error) = kernel32.GetCurrentThread
<原文结束>

# <翻译开始>
//sys	GetACP() (acp uint32) = kernel32.GetACP
//sys	GetConsoleCP() (ccp uint32) = kernel32.GetConsoleCP
//sys	MultiByteToWideChar(codePage uint32, dwFlags uint32, str *byte, nstr int32, wchar *uint16, nwchar int32) (nwrite int32, err error) = kernel32.MultiByteToWideChar
//sys	GetCurrentThread() (pseudoHandle syscall.Handle, err error) = kernel32.GetCurrentThread
# <翻译结束>


<原文开始>
//sys  NetShareAdd(serverName *uint16, level uint32, buf *byte, parmErr *uint16) (neterr error) = netapi32.NetShareAdd
//sys  NetShareDel(serverName *uint16, netName *uint16, reserved uint32) (neterr error) = netapi32.NetShareDel
<原文结束>

# <翻译开始>
//sys  NetShareAdd(serverName *uint16, level uint32, buf *byte, parmErr *uint16) (neterr error) = netapi32.NetShareAdd
//sys  NetShareDel(serverName *uint16, netName *uint16, reserved uint32) (neterr error) = netapi32.NetShareDel
# <翻译结束>


<原文开始>
//sys	GetFinalPathNameByHandle(file syscall.Handle, filePath *uint16, filePathSize uint32, flags uint32) (n uint32, err error) = kernel32.GetFinalPathNameByHandleW
<原文结束>

# <翻译开始>
//sys	GetFinalPathNameByHandle(file syscall.Handle, filePath *uint16, filePathSize uint32, flags uint32) (n uint32, err error) = kernel32.GetFinalPathNameByHandleW
# <翻译结束>


<原文开始>
//sys	CreateEnvironmentBlock(block **uint16, token syscall.Token, inheritExisting bool) (err error) = userenv.CreateEnvironmentBlock
//sys	DestroyEnvironmentBlock(block *uint16) (err error) = userenv.DestroyEnvironmentBlock
<原文结束>

# <翻译开始>
//sys	CreateEnvironmentBlock(block **uint16, token syscall.Token, inheritExisting bool) (err error) = userenv.CreateEnvironmentBlock
//sys	DestroyEnvironmentBlock(block *uint16) (err error) = userenv.DestroyEnvironmentBlock
# <翻译结束>


<原文开始>
//sys	RtlGenRandom(buf []byte) (err error) = advapi32.SystemFunction036
<原文结束>

# <翻译开始>
//sys	RtlGenRandom(buf []byte) (err error) = advapi32.SystemFunction036
# <翻译结束>

