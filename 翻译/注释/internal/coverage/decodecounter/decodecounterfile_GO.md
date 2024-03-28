
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
// This file contains helpers for reading counter data files created
// during the executions of a coverage-instrumented binary.
<原文结束>

# <翻译开始>
// 本文件包含一些辅助函数，用于读取在执行覆盖率统计的二进制文件过程中创建的计数器数据文件。
# <翻译结束>


<原文开始>
// GOARCH setting from run that produced counter data
<原文结束>

# <翻译开始>
// GOARCH 设置来自生成计数器数据的运行环境
# <翻译结束>


<原文开始>
// GOOS setting from run that produced counter data
<原文结束>

# <翻译开始>
// GOOS 设置来自生成计数器数据的运行环境
# <翻译结束>


<原文开始>
// Seek back to just past the file header.
<原文结束>

# <翻译开始>
// 将文件指针回溯到刚刚超过文件头的位置。
# <翻译结束>


<原文开始>
// Read preamble for first segment.
<原文结束>

# <翻译开始>
// 读取第一个段落的前导内容。
# <翻译结束>


<原文开始>
// readSegmentPreamble reads and consumes the segment header, segment string
// table, and segment args table.
<原文结束>

# <翻译开始>
// readSegmentPreamble 读取并消耗段头、段字符串表和段参数表。
# <翻译结束>


<原文开始>
// Read string table and args.
<原文结束>

# <翻译开始>
// 读取字符串表和参数。
# <翻译结束>


<原文开始>
// Seek past any padding to bring us up to a 4-byte boundary.
<原文结束>

# <翻译开始>
// 寻找并跳过任何填充内容，使我们移动到4字节边界处。
# <翻译结束>


<原文开始>
// OsArgs returns the program arguments (saved from os.Args during
// the run of the instrumented binary) read from the counter
// data file. Not all coverage data files will have os.Args values;
// for example, if a data file is produced by merging coverage
// data from two distinct runs, no os args will be available (an
// empty list is returned).
<原文结束>

# <翻译开始>
// OsArgs 从计数器数据文件中返回程序参数（在被监控的二进制文件运行期间，从os.Args保存的数据）。并非所有覆盖率数据文件都会包含os.Args值；例如，如果一个数据文件是由两个不同运行过程的覆盖率数据合并产生的，则不会提供os参数（将返回一个空列表）。
# <翻译结束>


<原文开始>
// Goos returns the GOOS setting in effect for the "-cover" binary
// that produced this counter data file. The GOOS value may be
// empty in the case where the counter data file was produced
// from a merge in which more than one GOOS value was present.
<原文结束>

# <翻译开始>
// Goos 函数返回生成此计数数据文件时用于“-cover”二进制文件的 GOOS 设置。当计数数据文件是由包含多个 GOOS 值的合并操作产生的时候，GOOS 的值可能会为空。
# <翻译结束>


<原文开始>
// Goarch returns the GOARCH setting in effect for the "-cover" binary
// that produced this counter data file. The GOARCH value may be
// empty in the case where the counter data file was produced
// from a merge in which more than one GOARCH value was present.
<原文结束>

# <翻译开始>
// Goarch 返回生成此计数数据文件的“-cover”二进制文件所采用的 GOARCH 设置。如果计数数据文件是由包含多个 GOARCH 值的合并操作产生的，则 GOARCH 值可能为空。
# <翻译结束>


<原文开始>
// FuncPayload encapsulates the counter data payload for a single
// function as read from a counter data file.
<原文结束>

# <翻译开始>
// FuncPayload 封装了从计数器数据文件中读取的单个函数的计数器数据负载。
# <翻译结束>


<原文开始>
// NumSegments returns the number of execution segments in the file.
<原文结束>

# <翻译开始>
// NumSegments 返回文件中的执行段落数量。
# <翻译结束>


<原文开始>
// BeginNextSegment sets up the the reader to read the next segment,
// returning TRUE if we do have another segment to read, or FALSE
// if we're done with all the segments (also an error if
// something went wrong).
<原文结束>

# <翻译开始>
// BeginNextSegment 函数用于设置读取器以读取下一个分段，
// 若存在下一个分段可供读取，则返回 TRUE，若所有分段已读完则返回 FALSE。
// 若在执行过程中出现错误，也将返回错误。
# <翻译结束>


<原文开始>
// Seek past footer from last segment.
<原文结束>

# <翻译开始>
// 从最后一个分段中跳过尾部。
# <翻译结束>


<原文开始>
// Read preamble for this segment.
<原文结束>

# <翻译开始>
// 读取本段的前导部分。
# <翻译结束>


<原文开始>
// NumFunctionsInSegment returns the number of live functions
// in the currently selected segment.
<原文结束>

# <翻译开始>
// NumFunctionsInSegment 返回当前选定段中活动函数的数量。
# <翻译结束>


<原文开始>
// NextFunc reads data for the next function in this current segment
// into "p", returning TRUE if the read was successful or FALSE
// if we've read all the functions already (also an error if
// something went wrong with the read or we hit a premature
// EOF).
<原文结束>

# <翻译开始>
// NextFunc 函数将当前段落中下一个函数的数据读取到 "p" 中，
// 如果读取成功则返回 TRUE，如果已经读取完所有函数则返回 FALSE。
// 如果在读取过程中出现错误或意外遇到文件结束符（EOF），也会返回 FALSE。
# <翻译结束>


<原文开始>
	// Alternative/experimental path: one way we could handling writing
	// out counter data would be to just memcpy the counter segment
	// out to a file, meaning that a region in the counter memory
	// corresponding to a dead (never-executed) function would just be
	// zeroes. The code path below handles this case.
<原文结束>

# <翻译开始>
// 另一种/实验性路径：处理写入计数器数据的一种方式是直接使用memcpy将计数器段复制到文件中，这意味着与从未执行过的“死亡”函数相对应的计数器内存区域将只是零。下面的代码路径处理了这种情况。
# <翻译结束>


<原文开始>
// Read package and func indices.
<原文结束>

# <翻译开始>
// 读取包和函数索引。
# <翻译结束>

