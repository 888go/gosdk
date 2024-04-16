
<原文开始>
// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// 版权所有 2016 The Go 作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。
# <翻译结束>


<原文开始>
		// For wasip1, some runtimes forbid absolute symlinks,
		// or symlinks that escape the current working directory.
		// Perform a simple test to see whether the runtime
		// supports symlinks or not. If we get a permission
		// error, the runtime does not support symlinks.
<原文结束>

# <翻译开始>
// 对于wasip1，某些运行时禁止使用绝对路径的符号链接，
// 或者是那些指向当前工作目录之外的符号链接。
// 我们进行一个简单的测试以判断该运行时是否支持符号链接。
// 如果遇到权限错误，说明运行时不支持符号链接。
# <翻译结束>

