// 版权所有 2016 The Go 作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

//go:build unix || (js && wasm) || wasip1

package os_test

import (
	"github.com/888go/gosdk/os"
	"io/fs"
	"syscall"
)

func init() {
	isExistTests = append(isExistTests,
		isExistTest{err: &fs.PathError{Err: syscall.EEXIST}, is: true, isnot: false},
		isExistTest{err: &fs.PathError{Err: syscall.ENOTEMPTY}, is: true, isnot: false},

		isExistTest{err: &os.LinkError{Err: syscall.EEXIST}, is: true, isnot: false},
		isExistTest{err: &os.LinkError{Err: syscall.ENOTEMPTY}, is: true, isnot: false},

		isExistTest{err: &os.SyscallError{Err: syscall.EEXIST}, is: true, isnot: false},
		isExistTest{err: &os.SyscallError{Err: syscall.ENOTEMPTY}, is: true, isnot: false},
	)
	isPermissionTests = append(isPermissionTests,
		isPermissionTest{err: &fs.PathError{Err: syscall.EACCES}, want: true},
		isPermissionTest{err: &fs.PathError{Err: syscall.EPERM}, want: true},
		isPermissionTest{err: &fs.PathError{Err: syscall.EEXIST}, want: false},

		isPermissionTest{err: &os.LinkError{Err: syscall.EACCES}, want: true},
		isPermissionTest{err: &os.LinkError{Err: syscall.EPERM}, want: true},
		isPermissionTest{err: &os.LinkError{Err: syscall.EEXIST}, want: false},

		isPermissionTest{err: &os.SyscallError{Err: syscall.EACCES}, want: true},
		isPermissionTest{err: &os.SyscallError{Err: syscall.EPERM}, want: true},
		isPermissionTest{err: &os.SyscallError{Err: syscall.EEXIST}, want: false},
	)

}
