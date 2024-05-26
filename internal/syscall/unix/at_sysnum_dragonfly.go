// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package unix

import "syscall"

const unlinkatTrap uintptr = syscall.SYS_UNLINKAT
const openatTrap uintptr = syscall.SYS_OPENAT
const fstatatTrap uintptr = syscall.SYS_FSTATAT

// 翻译提示:const 删除目录标志 = 0x2
const AT_REMOVEDIR = 0x2
// 翻译提示:const AT_SYMLINK_NOFOLLOW = 0x1
//
//// 改为中文版
//const AT 符号链接不跟随 = 0x1
const AT_SYMLINK_NOFOLLOW = 0x1
