
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
// This package contains APIs and helpers for decoding a single package's
// meta data "blob" emitted by the compiler when coverage instrumentation
// is turned on.
<原文结束>

# <翻译开始>
// 此包包含API和辅助函数，用于解码在启用覆盖率检测时，编译器生成的单个包的元数据“块”。
# <翻译结束>


<原文开始>
// See comments in the encodecovmeta package for details on the format.
<原文结束>

# <翻译开始>
// 有关该格式的详细信息，请参阅encodecovmeta包中的注释。
# <翻译结束>


<原文开始>
// Seek to the correct location to read the string table.
<原文结束>

# <翻译开始>
// 寻找到正确的位置以读取字符串表。
# <翻译结束>


<原文开始>
// ReadFunc reads the coverage meta-data for the function with index
// 'findex', filling it into the FuncDesc pointed to by 'f'.
<原文结束>

# <翻译开始>
// ReadFunc 读取具有索引 'findex' 的函数的覆盖率元数据，并将其填充到指向的 FuncDesc 结构体中。
// 'f' 参数即指向该结构体。
# <翻译结束>


<原文开始>
// Seek to the correct location to read the function offset and read it.
<原文结束>

# <翻译开始>
// 寻找到正确的位置以读取函数偏移量并读取它。
# <翻译结束>


<原文开始>
// Seek to the correct location to read the function.
<原文结束>

# <翻译开始>
// 寻找并定位到读取函数的正确位置。
# <翻译结束>


<原文开始>
// Preamble containing number of units, file, and function.
<原文结束>

# <翻译开始>
// 包含单元数、文件名和函数的前导注释。
# <翻译结束>

