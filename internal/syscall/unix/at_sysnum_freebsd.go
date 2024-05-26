// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package unix

import "syscall"

// 翻译提示:const (
//	删除目录标志     = 0x800
//	不跟随符号链接 = 0x200
//
//	删除at陷阱    uintptr = syscall.SYS_UNLINKAT
//	打开at陷阱     uintptr = syscall.SYS_OPENAT
//) 
//
//// 中文注释版
//// AT_REMOVEDIR 表示在AT指令中删除目录的标志
//// AT_SYMLINK_NOFOLLOW 表示在AT指令中不跟随符号链接的标志
//// unlinkatTrap 是删除at操作的系统调用陷阱，类型为uintptr
//// openatTrap 是打开at操作的系统调用陷阱，类型为uintptr
const (
	AT_REMOVEDIR        = 0x800
	AT_SYMLINK_NOFOLLOW = 0x200

	unlinkatTrap uintptr = syscall.SYS_UNLINKAT
	openatTrap   uintptr = syscall.SYS_OPENAT
)
