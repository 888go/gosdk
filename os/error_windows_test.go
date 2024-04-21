// 版权所有 2016 The Go 作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

//go:build windows

package os_test

import (
	"io/fs"
	"os"
	"syscall"
)

func init() {
	const _ERROR_BAD_NETPATH = syscall.Errno(53)

	isExistTests = append(isExistTests,
		isExistTest{err: &fs.PathError{Err: syscall.ERROR_FILE_NOT_FOUND}, is: false, isnot: true},
		isExistTest{err: &os.LinkError{Err: syscall.ERROR_FILE_NOT_FOUND}, is: false, isnot: true},
		isExistTest{err: &os.SyscallError{Err: syscall.ERROR_FILE_NOT_FOUND}, is: false, isnot: true},

		isExistTest{err: &fs.PathError{Err: _ERROR_BAD_NETPATH}, is: false, isnot: true},
		isExistTest{err: &os.LinkError{Err: _ERROR_BAD_NETPATH}, is: false, isnot: true},
		isExistTest{err: &os.SyscallError{Err: _ERROR_BAD_NETPATH}, is: false, isnot: true},

		isExistTest{err: &fs.PathError{Err: syscall.ERROR_PATH_NOT_FOUND}, is: false, isnot: true},
		isExistTest{err: &os.LinkError{Err: syscall.ERROR_PATH_NOT_FOUND}, is: false, isnot: true},
		isExistTest{err: &os.SyscallError{Err: syscall.ERROR_PATH_NOT_FOUND}, is: false, isnot: true},

		isExistTest{err: &fs.PathError{Err: syscall.ERROR_DIR_NOT_EMPTY}, is: true, isnot: false},
		isExistTest{err: &os.LinkError{Err: syscall.ERROR_DIR_NOT_EMPTY}, is: true, isnot: false},
		isExistTest{err: &os.SyscallError{Err: syscall.ERROR_DIR_NOT_EMPTY}, is: true, isnot: false},
	)
	isPermissionTests = append(isPermissionTests,
		isPermissionTest{err: &fs.PathError{Err: syscall.ERROR_ACCESS_DENIED}, want: true},
		isPermissionTest{err: &os.LinkError{Err: syscall.ERROR_ACCESS_DENIED}, want: true},
		isPermissionTest{err: &os.SyscallError{Err: syscall.ERROR_ACCESS_DENIED}, want: true},
	)
}
