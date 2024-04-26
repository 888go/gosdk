
<原文开始>
// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
//版权所有2016 The Go Authors。所有权利保留。
//使用此源代码受BSD风格的
//可在LICENSE文件中找到的许可协议管辖。
# <翻译结束>


<原文开始>
		// For wasip1, some runtimes forbid absolute symlinks,
		// or symlinks that escape the current working directory.
		// Perform a simple test to see whether the runtime
		// supports symlinks or not. If we get a permission
		// error, the runtime does not support symlinks.
<原文结束>

# <翻译开始>
// 对于wasip1，某些运行时禁止使用绝对符号链接，
// 或者那些能够脱离当前工作目录的符号链接。
// 执行一个简单的测试以判断运行时是否支持符号链接。
// 如果我们遇到权限错误，则说明运行时不支持符号链接。
# <翻译结束>

