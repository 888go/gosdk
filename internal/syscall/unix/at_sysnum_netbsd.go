// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package unix

import "syscall"

const unlinkatTrap uintptr = syscall.SYS_UNLINKAT
const openatTrap uintptr = syscall.SYS_OPENAT
const fstatatTrap uintptr = syscall.SYS_FSTATAT

// 翻译提示:const (
//    删除目录标志 = 0x800
//)
const AT_REMOVEDIR = 0x800
// 翻译提示:const (
//    // AT_SYMLINK_NOFOLLOW 表示在调用如open，access等系统调用时不跟随符号链接。
//    AT_SYMLINK_NOFOLLOW = 0x200
//)
const AT_SYMLINK_NOFOLLOW = 0x200
