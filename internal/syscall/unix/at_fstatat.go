// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build dragonfly || (linux && !loong64) || netbsd || (openbsd && mips64)

package unix

import (
	"syscall"
	"unsafe"
)

// 翻译提示:func 文件状态统计在目录(fileDescriptor int, 路径 string, 状态结构体 *系统调用.Stat_t, 标志 int) error {}

// ff:
// flags:
// stat:
// path:
// dirfd:
func Fstatat(dirfd int, path string, stat *syscall.Stat_t, flags int) error {
	var p *byte
	p, err := syscall.BytePtrFromString(path)
	if err != nil {
		return err
	}

	_, _, errno := syscall.Syscall6(fstatatTrap, uintptr(dirfd), uintptr(unsafe.Pointer(p)), uintptr(unsafe.Pointer(stat)), uintptr(flags), 0, 0)
	if errno != 0 {
		return errno
	}

	return nil

}
