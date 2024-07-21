
<原文开始>
// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// 版权所有 ? 2017 The Go Authors。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。
# <翻译结束>


<原文开始>
// Check that subsequent operations didn't change the returned strings.
<原文结束>

# <翻译开始>
// 检查后续操作是否改变了返回的字符串。
# <翻译结束>


<原文开始>
	// Ensure that writing after Reset doesn't alter
	// previously returned strings.
<原文结束>

# <翻译开始>
	// 确保在调用 Reset 后进行写入操作不会更改先前已返回的字符串。
# <翻译结束>


<原文开始>
// should be only alloc, when growLen > 0
<原文结束>

# <翻译开始>
// 应仅在 growLen > 0 时进行分配
# <翻译结束>


<原文开始>
// when growLen < 0, should panic
<原文结束>

# <翻译开始>
// 当 growLen < 0 时，应当引发 panic
# <翻译结束>


<原文开始>
	// Issue 23382; verify that copyCheck doesn't force the
	// Builder to escape and be heap allocated.
<原文结束>

# <翻译开始>
	// Issue 23382；验证copyCheck不会导致
	// Builder被迫逃逸并被堆分配。
# <翻译结束>


<原文开始>
	// Invalid runes, including negative ones, should be written as
	// utf8.RuneError.
<原文结束>

# <翻译开始>
	// 无效的 rune（包括负值），应写作 utf8.RuneError。
# <翻译结束>

