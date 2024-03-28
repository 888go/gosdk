
<原文开始>
// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// 版权所有 ? 2021 Go作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。
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
// 1 error per 8 characters, plus a few slop to start.
<原文结束>

# <翻译开始>
// 每8个字符出现1个错误，再加上一些初始误差。
# <翻译结束>

