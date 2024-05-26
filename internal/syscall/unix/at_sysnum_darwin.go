// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package unix

// 翻译提示:const (
//    删除目录标志 = 0x80
//)
const AT_REMOVEDIR = 0x80
// 翻译提示:const AT_SYMLINK_NOFOLLOW = 0x0020
//
//可以翻译为：
//
//常量 AT 符号链接不跟随 = 0x0020
//
//这里的参数和返回值只有一个常量，无需单独翻译。在Go语言中，常量通常用英文表示，以保持代码的通用性和国际化的可读性。如果非要翻译，可以考虑：
//
//常量 不跟随符号链接 = 0x0020
//
//但请注意，这样的翻译可能会降低代码的清晰度，对于编程而言，原英文常量名通常是最佳选择。
const AT_SYMLINK_NOFOLLOW = 0x0020
