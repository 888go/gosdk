// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package unix

import "syscall"

const unlinkatTrap uintptr = syscall.SYS_UNLINKAT
const openatTrap uintptr = syscall.SYS_OPENAT
const fstatatTrap uintptr = syscall.SYS_FSTATAT

// 翻译提示:const (
//    删除目录标志 = 0x08
//)
const AT_REMOVEDIR = 0x08
// 翻译提示:const AT_SYMLINK_NOFOLLOW = 0x02
//
//可以翻译为：
//
//常量 AT 符号链接不跟随 = 0x02
//
//这里的参数和返回值只有一个常量，无需翻译。在Go语言中，常量名称通常保留英文，以保持代码的通用性和可读性，但在注释中可以提供中文解释。
const AT_SYMLINK_NOFOLLOW = 0x02
