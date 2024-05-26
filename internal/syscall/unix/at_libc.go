// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build aix || solaris

package unix

import (
	"syscall"
	"unsafe"
)

//go:linkname procFstatat libc_fstatat
//go:linkname procOpenat libc_openat
//go:linkname procUnlinkat libc_unlinkat

// 翻译提示:var (
//	进程Fstatat,
//	进程Openat,
//	进程Unlinkat uint64
//) 
//
//注意：在Go语言中，通常我们不鼓励使用中文作为变量名，因为这可能降低代码的可读性，特别是在国际化的开发环境中。以上的翻译只是为了理解变量含义，实际编码应保持原英文状态。
var (
	procFstatat,
	procOpenat,
	procUnlinkat uintptr
)

// 翻译提示:func 删除链接(dirfd 文件描述符, path 路径字符串, flags 标志整数) error {}

// ff:
// flags:
// path:
// dirfd:
func Unlinkat(dirfd int, path string, flags int) error {
	p, err := syscall.BytePtrFromString(path)
	if err != nil {
		return err
	}

	_, _, errno := syscall6(uintptr(unsafe.Pointer(&procUnlinkat)), 3, uintptr(dirfd), uintptr(unsafe.Pointer(p)), uintptr(flags), 0, 0, 0)
	if errno != 0 {
		return errno
	}

	return nil
}

// 翻译提示:func 在目录打开文件(dirfd int, 路径 string, 打开标志 int, 权限 uint32) (文件描述符 int, 错误 error) {}

// ff:
// perm:
// flags:
// path:
// dirfd:
func Openat(dirfd int, path string, flags int, perm uint32) (int, error) {
	p, err := syscall.BytePtrFromString(path)
	if err != nil {
		return 0, err
	}

	fd, _, errno := syscall6(uintptr(unsafe.Pointer(&procOpenat)), 4, uintptr(dirfd), uintptr(unsafe.Pointer(p)), uintptr(flags), uintptr(perm), 0, 0)
	if errno != 0 {
		return 0, errno
	}

	return int(fd), nil
}

// 翻译提示:func 文件状态统计在目录(fileDescriptor int, 路径 string, 状态结构体 *系统调用.Stat_t, 标志 int) error {}

// ff:
// flags:
// stat:
// path:
// dirfd:
func Fstatat(dirfd int, path string, stat *syscall.Stat_t, flags int) error {
	p, err := syscall.BytePtrFromString(path)
	if err != nil {
		return err
	}

	_, _, errno := syscall6(uintptr(unsafe.Pointer(&procFstatat)), 4, uintptr(dirfd), uintptr(unsafe.Pointer(p)), uintptr(unsafe.Pointer(stat)), uintptr(flags), 0, 0)
	if errno != 0 {
		return errno
	}

	return nil
}
