// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build freebsd || (linux && loong64)

package unix

import "syscall"

// 翻译提示:func 文件状态统计在目录(fileDescriptor int, 路径 string, 状态结构体 *系统调用.Stat_t, 标志 int) error {}

// ff:
// flags:
// stat:
// path:
// dirfd:
func Fstatat(dirfd int, path string, stat *syscall.Stat_t, flags int) error {
	return syscall.Fstatat(dirfd, path, stat, flags)
}
