// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package unix

//go:cgo_import_dynamic libc_fstatat fstatat "libc.a/shr_64.o"
//go:cgo_import_dynamic libc_openat openat "libc.a/shr_64.o"
//go:cgo_import_dynamic libc_unlinkat unlinkat "libc.a/shr_64.o"

// 翻译提示:const (
//	删除目录    = 0x1
//	不跟随符号链接 = 0x1
//)
//
//// 参数和返回值名称保持不变，因为这些常量通常用于系统调用或函数参数，而函数名和参数名通常是英文的，以保持与标准库的一致性。
const (
	AT_REMOVEDIR        = 0x1
	AT_SYMLINK_NOFOLLOW = 0x1
)
