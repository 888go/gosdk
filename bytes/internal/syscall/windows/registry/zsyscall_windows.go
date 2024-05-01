// 代码由'go generate'命令生成；请勿编辑。

package registry

import (
	"github.com/888go/gosdk/bytes/internal/syscall/windows/sysdll"
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
	modadvapi32 = syscall.NewLazyDLL(sysdll.Add("advapi32.dll"))
	modkernel32 = syscall.NewLazyDLL(sysdll.Add("kernel32.dll"))

	procRegCreateKeyExW           = modadvapi32.NewProc("RegCreateKeyExW")
	procRegDeleteKeyW             = modadvapi32.NewProc("RegDeleteKeyW")
	procRegDeleteValueW           = modadvapi32.NewProc("RegDeleteValueW")
	procRegEnumValueW             = modadvapi32.NewProc("RegEnumValueW")
	procRegLoadMUIStringW         = modadvapi32.NewProc("RegLoadMUIStringW")
	procRegSetValueExW            = modadvapi32.NewProc("RegSetValueExW")
	procExpandEnvironmentStringsW = modkernel32.NewProc("ExpandEnvironmentStringsW")
)

func regCreateKeyEx(key syscall.Handle, subkey *uint16, reserved uint32, class *uint16, options uint32, desired uint32, sa *syscall.SecurityAttributes, result *syscall.Handle, disposition *uint32) (regerrno error) {
	r0, _, _ := syscall.Syscall9(procRegCreateKeyExW.Addr(), 9, uintptr(key), uintptr(unsafe.Pointer(subkey)), uintptr(reserved), uintptr(unsafe.Pointer(class)), uintptr(options), uintptr(desired), uintptr(unsafe.Pointer(sa)), uintptr(unsafe.Pointer(result)), uintptr(unsafe.Pointer(disposition)))
	if r0 != 0 {
		regerrno = syscall.Errno(r0)
	}
	return
}

func regDeleteKey(key syscall.Handle, subkey *uint16) (regerrno error) {
	r0, _, _ := syscall.Syscall(procRegDeleteKeyW.Addr(), 2, uintptr(key), uintptr(unsafe.Pointer(subkey)), 0)
	if r0 != 0 {
		regerrno = syscall.Errno(r0)
	}
	return
}

func regDeleteValue(key syscall.Handle, name *uint16) (regerrno error) {
	r0, _, _ := syscall.Syscall(procRegDeleteValueW.Addr(), 2, uintptr(key), uintptr(unsafe.Pointer(name)), 0)
	if r0 != 0 {
		regerrno = syscall.Errno(r0)
	}
	return
}

func regEnumValue(key syscall.Handle, index uint32, name *uint16, nameLen *uint32, reserved *uint32, valtype *uint32, buf *byte, buflen *uint32) (regerrno error) {
	r0, _, _ := syscall.Syscall9(procRegEnumValueW.Addr(), 8, uintptr(key), uintptr(index), uintptr(unsafe.Pointer(name)), uintptr(unsafe.Pointer(nameLen)), uintptr(unsafe.Pointer(reserved)), uintptr(unsafe.Pointer(valtype)), uintptr(unsafe.Pointer(buf)), uintptr(unsafe.Pointer(buflen)), 0)
	if r0 != 0 {
		regerrno = syscall.Errno(r0)
	}
	return
}

func regLoadMUIString(key syscall.Handle, name *uint16, buf *uint16, buflen uint32, buflenCopied *uint32, flags uint32, dir *uint16) (regerrno error) {
	r0, _, _ := syscall.Syscall9(procRegLoadMUIStringW.Addr(), 7, uintptr(key), uintptr(unsafe.Pointer(name)), uintptr(unsafe.Pointer(buf)), uintptr(buflen), uintptr(unsafe.Pointer(buflenCopied)), uintptr(flags), uintptr(unsafe.Pointer(dir)), 0, 0)
	if r0 != 0 {
		regerrno = syscall.Errno(r0)
	}
	return
}

func regSetValueEx(key syscall.Handle, valueName *uint16, reserved uint32, vtype uint32, buf *byte, bufsize uint32) (regerrno error) {
	r0, _, _ := syscall.Syscall6(procRegSetValueExW.Addr(), 6, uintptr(key), uintptr(unsafe.Pointer(valueName)), uintptr(reserved), uintptr(vtype), uintptr(unsafe.Pointer(buf)), uintptr(bufsize))
	if r0 != 0 {
		regerrno = syscall.Errno(r0)
	}
	return
}

func expandEnvironmentStrings(src *uint16, dst *uint16, size uint32) (n uint32, err error) {
	r0, _, e1 := syscall.Syscall(procExpandEnvironmentStringsW.Addr(), 3, uintptr(unsafe.Pointer(src)), uintptr(unsafe.Pointer(dst)), uintptr(size))
	n = uint32(r0)
	if n == 0 {
		err = errnoErr(e1)
	}
	return
}
