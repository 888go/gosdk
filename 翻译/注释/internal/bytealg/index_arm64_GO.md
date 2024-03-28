
<原文开始>
// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// 版权所有 ? 2018 The Go Authors。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。
# <翻译结束>


<原文开始>
// Empirical data shows that using Index can get better
// performance when len(s) <= 16.
<原文结束>

# <翻译开始>
// 经验数据表明，当 len(s) <= 16 时，使用 Index 可以获得更好的性能。
# <翻译结束>


<原文开始>
// Optimize cases where the length of the substring is less than 32 bytes
<原文结束>

# <翻译开始>
// 优化子字符串长度小于 32 字节的情况
# <翻译结束>


<原文开始>
// Cutover reports the number of failures of IndexByte we should tolerate
// before switching over to Index.
// n is the number of bytes processed so far.
// See the bytes.Index implementation for details.
<原文结束>

# <翻译开始>
// Cutover 报告我们在切换到 Index 之前应容忍的 IndexByte 失败次数。
// n 是迄今为止已处理的字节数。
// 有关详细信息，请参阅 bytes.Index 的实现。
# <翻译结束>


<原文开始>
// 1 error per 16 characters, plus a few slop to start.
<原文结束>

# <翻译开始>
// 每16个字符出现1个错误，另外还留有一些余量开始。
# <翻译结束>

