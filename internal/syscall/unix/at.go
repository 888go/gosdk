// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build dragonfly || freebsd || linux || netbsd || (openbsd && mips64)

package unix

import (
	"syscall"
	"unsafe"
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

	_, _, errno := syscall.Syscall(unlinkatTrap, uintptr(dirfd), uintptr(unsafe.Pointer(p)), uintptr(flags))
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

	fd, _, errno := syscall.Syscall6(openatTrap, uintptr(dirfd), uintptr(unsafe.Pointer(p)), uintptr(flags), uintptr(perm), 0, 0)
	if errno != 0 {
		return 0, errno
	}

	return int(fd), nil
}
