
<原文开始>
// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// 版权所有 2019 Go作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可协议约束，
// 该协议可在 LICENSE 文件中找到。
# <翻译结束>


<原文开始>
// Equal reports whether a and b
// are the same length and contain the same bytes.
// A nil argument is equivalent to an empty slice.
//
// Equal is equivalent to bytes.Equal.
// It is provided here for convenience,
// because some packages cannot depend on bytes.
<原文结束>

# <翻译开始>
// Equal 函数报告 a 和 b 是否具有相同的长度，并且包含相同的字节。
// 如果参数为 nil，则视为空切片。
//
// Equal 等价于 bytes.Equal，这里为了方便提供该函数，
// 因为有些包不能依赖 bytes 包。
# <翻译结束>


<原文开始>
	// Neither cmd/compile nor gccgo allocates for these string conversions.
	// There is a test for this in package bytes.
<原文结束>

# <翻译开始>
// Go编译器（cmd/compile）和gccgo都不会为这些字符串转换分配内存。
// 在bytes包中有一个针对此行为的测试。
# <翻译结束>

