
<原文开始>
// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// 版权所有 ? 2015 Go作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。
# <翻译结束>


<原文开始>
// skipStdinCopyError optionally specifies a function which reports
// whether the provided stdin copy error should be ignored.
<原文结束>

# <翻译开始>
// skipStdinCopyError 可选地指定一个函数，该函数用于判断提供的从标准输入复制时产生的错误是否应被忽略。
# <翻译结束>


<原文开始>
	// Ignore EPIPE errors copying to stdin if the program
	// completed successfully otherwise.
	// See Issue 9173.
<原文结束>

# <翻译开始>
// 如果程序已经成功完成，忽略向stdin复制时出现的EPIPE错误。
// 参见Issue 9173。
# <翻译结束>

