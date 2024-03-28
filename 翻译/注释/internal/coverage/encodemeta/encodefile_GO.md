
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
// This package contains APIs and helpers for writing out a meta-data
// file (composed of a file header, offsets/lengths, and then a series of
// meta-data blobs emitted by the compiler, one per Go package).
<原文结束>

# <翻译开始>
// 此包包含编写元数据文件所需的API和辅助函数（由文件头、偏移量/长度以及一系列由编译器生成的元数据块组成，每个Go包对应一个元数据块）。
# <翻译结束>


<原文开始>
// Emit package offsets section followed by package lengths section.
<原文结束>

# <翻译开始>
// 发射（生成）包偏移量部分，随后是包长度部分。
# <翻译结束>


<原文开始>
// Now emit blobs themselves.
<原文结束>

# <翻译开始>
// 现在开始输出（或生成）blob本身。
# <翻译结束>


<原文开始>
// Flush writer, and we're done.
<原文结束>

# <翻译开始>
// 刷新写入器，完成操作。
# <翻译结束>

