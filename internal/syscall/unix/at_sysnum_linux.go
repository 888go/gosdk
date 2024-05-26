// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package unix

import "syscall"

const unlinkatTrap uintptr = syscall.SYS_UNLINKAT
const openatTrap uintptr = syscall.SYS_OPENAT

// 翻译提示:const (
//	访问权限错误 = 0x200
//	当前目录文件描述符 = -0x64
//	删除目录 = 0x200
//	不跟随符号链接 = 0x100
//)
const (
	AT_EACCESS          = 0x200
	AT_FDCWD            = -0x64
	AT_REMOVEDIR        = 0x200
	AT_SYMLINK_NOFOLLOW = 0x100
)
