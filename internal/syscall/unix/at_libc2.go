// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build darwin || (openbsd && !mips64)

package unix

import (
	"syscall"
	_ "unsafe" // for linkname
)

// 翻译提示:func 删除链接(dirfd 文件描述符, path 路径字符串, flags 标志整数) error {}

// ff:
// flags:
// path:
// dirfd:
func Unlinkat(dirfd int, path string, flags int) error {
	return unlinkat(dirfd, path, flags)
}

// 翻译提示:func 在目录打开文件(dirfd int, 路径 string, 打开标志 int, 权限 uint32) (文件描述符 int, 错误 error) {}

// ff:
// perm:
// flags:
// path:
// dirfd:
func Openat(dirfd int, path string, flags int, perm uint32) (int, error) {
	return openat(dirfd, path, flags, perm)
}

// 翻译提示:func 文件状态统计在目录(fileDescriptor int, 路径 string, 状态结构体 *系统调用.Stat_t, 标志 int) error {}

// ff:
// flags:
// stat:
// path:
// dirfd:
func Fstatat(dirfd int, path string, stat *syscall.Stat_t, flags int) error {
	return fstatat(dirfd, path, stat, flags)
}

//go:linkname unlinkat syscall.unlinkat
func unlinkat(dirfd int, path string, flags int) error

//go:linkname openat syscall.openat
func openat(dirfd int, path string, flags int, perm uint32) (int, error)

//go:linkname fstatat syscall.fstatat
func fstatat(dirfd int, path string, stat *syscall.Stat_t, flags int) error
