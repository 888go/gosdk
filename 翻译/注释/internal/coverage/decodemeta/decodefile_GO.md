
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
// This package contains APIs and helpers for reading and decoding
// meta-data output files emitted by the runtime when a
// coverage-instrumented binary executes. A meta-data file contains
// top-level info (counter mode, number of packages) and then a
// separate self-contained meta-data section for each Go package.
<原文结束>

# <翻译开始>
// 该包包含用于读取和解码由运行时生成的元数据输出文件的API和辅助函数，
// 当一个带有覆盖率检测的二进制程序执行时，会生成这些文件。元数据文件包含顶层信息（计数器模式、包数量），
// 然后为每个Go包包含一个独立的、自包含的元数据部分。
# <翻译结束>


<原文开始>
// CoverageMetaFileReader provides state and methods for reading
// a meta-data file from a code coverage run.
<原文结束>

# <翻译开始>
// CoverageMetaFileReader 提供状态和方法，用于从代码覆盖率运行中读取元数据文件。
# <翻译结束>


<原文开始>
// NewCoverageMetaFileReader returns a new helper object for reading
// the coverage meta-data output file 'f'. The param 'fileView' is a
// read-only slice containing the contents of 'f' obtained by mmap'ing
// the file read-only; 'fileView' may be nil, in which case the helper
// will read the contents of the file using regular file Read
// operations.
<原文结束>

# <翻译开始>
// NewCoverageMetaFileReader 返回一个用于读取覆盖率元数据输出文件 'f' 的新辅助对象。
// 参数 'fileView' 是通过将文件以只读方式mmap映射得到的、包含文件 'f' 内容的只读切片；
// 'fileView' 可能为nil，在这种情况下，该辅助工具将使用常规文件Read操作来读取文件内容。
# <翻译结束>


<原文开始>
	// Vet the version. If this is a meta-data file from the future,
	// we won't be able to read it.
<原文结束>

# <翻译开始>
// 验证版本。如果这是一个来自未来的元数据文件，
// 我们将无法读取它。
# <翻译结束>


<原文开始>
// Read package offsets for good measure
<原文结束>

# <翻译开始>
// 为了稳妥起见，读取包的偏移量
# <翻译结束>


<原文开始>
// NumPackages returns the number of packages for which this file
// contains meta-data.
<原文结束>

# <翻译开始>
// NumPackages 返回此文件中包含元数据的包的数量。
# <翻译结束>


<原文开始>
// CounterMode returns the counter mode (set, count, atomic) used
// when building for coverage for the program that produce this
// meta-data file.
<原文结束>

# <翻译开始>
// CounterMode 返回在为生成此元数据文件的程序构建覆盖率时所使用的计数器模式（设置、计数、原子）
# <翻译结束>


<原文开始>
// CounterMode returns the counter granularity (single counter per
// function, or counter per block) selected when building for coverage
// for the program that produce this meta-data file.
<原文结束>

# <翻译开始>
// CounterMode 返回在为生成此元数据文件的程序构建覆盖率时选择的计数器粒度（每个函数单个计数器或每个块一个计数器）。
# <翻译结束>


<原文开始>
// FileHash returns the hash computed for all of the package meta-data
// blobs. Coverage counter data files refer to this hash, and the
// hash will be encoded into the meta-data file name.
<原文结束>

# <翻译开始>
// FileHash 返回对所有包元数据blob计算出的哈希值。
// 覆盖率计数器数据文件会引用这个哈希值，并且该哈希值会被编码到元数据文件名中。
# <翻译结束>


<原文开始>
// GetPackageDecoder requests a decoder object for the package within
// the meta-data file whose index is 'pkIdx'. If the
// CoverageMetaFileReader was set up with a read-only file view, a
// pointer into that file view will be returned, otherwise the buffer
// 'payloadbuf' will be written to (or if it is not of sufficient
// size, a new buffer will be allocated). Return value is the decoder,
// a byte slice with the encoded meta-data, and an error.
<原文结束>

# <翻译开始>
// GetPackageDecoder 通过索引值 'pkIdx' 请求元数据文件中包的解码器对象。如果 CoverageMetaFileReader 是使用只读文件视图设置的，则会返回该文件视图中的指针，否则将向缓冲区 'payloadbuf' 写入数据（或者如果其大小不足，则会分配新的缓冲区）。返回值包括解码器、编码后的元数据字节切片以及错误信息。
# <翻译结束>


<原文开始>
// GetPackagePayload returns the raw (encoded) meta-data payload for the
// package with index 'pkIdx'. As with GetPackageDecoder, if the
// CoverageMetaFileReader was set up with a read-only file view, a
// pointer into that file view will be returned, otherwise the buffer
// 'payloadbuf' will be written to (or if it is not of sufficient
// size, a new buffer will be allocated). Return value is the decoder,
// a byte slice with the encoded meta-data, and an error.
<原文结束>

# <翻译开始>
// GetPackagePayload 返回索引为'pkIdx'的包的原始（已编码）元数据有效负载。与GetPackageDecoder类似，如果CoverageMetaFileReader使用只读文件视图设置，则会返回该文件视图的指针；否则将写入缓冲区'payloadbuf'（如果其大小不足，则会分配新的缓冲区）。返回值包括解码器、包含已编码元数据的字节切片以及错误信息。
# <翻译结束>


<原文开始>
// Determine correct offset/length.
<原文结束>

# <翻译开始>
// 确定正确的偏移量/长度。
# <翻译结束>

