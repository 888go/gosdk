
<原文开始>
// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// 版权所有 ? 2012 The Go Authors。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。
# <翻译结束>


<原文开始>
	// Test for the race detector, to verify ReadAt doesn't mutate
	// any state.
<原文结束>

# <翻译开始>
// 用于竞态检测器的测试，以验证ReadAt不会改变任何状态
# <翻译结束>


<原文开始>
	// Test for the race detector, to verify a Read that doesn't yield any bytes
	// is okay to use from multiple goroutines. This was our historic behavior.
	// See golang.org/issue/7856
<原文结束>

# <翻译开始>
// 用于竞态检测器的测试，以验证从多个goroutine中使用未读取任何字节的Read操作是正常的。这是我们历史上的行为。
// 参见golang.org/issue/7856
# <翻译结束>


<原文开始>
// tests that Len is affected by reads, but Size is not.
<原文结束>

# <翻译开始>
// 测试Len会受到读取操作的影响，但Size不受影响
# <翻译结束>

