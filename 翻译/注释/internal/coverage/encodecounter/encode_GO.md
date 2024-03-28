
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
// This package contains APIs and helpers for encoding initial portions
// of the counter data files emitted at runtime when coverage instrumentation
// is enabled.  Counter data files may contain multiple segments; the file
// header and first segment are written via the "Write" method below, and
// additional segments can then be added using "AddSegment".
<原文结束>

# <翻译开始>
// 该包包含在启用覆盖率检测时，用于编码运行时生成的计数器数据文件初始部分的API和辅助函数。计数器数据文件可能包含多个段；通过下面的"Write"方法写入文件头和第一个段，然后可以使用"AddSegment"添加更多段。
# <翻译结束>


<原文开始>
// CounterVisitor describes a helper object used during counter file
// writing; when writing counter data files, clients pass a
// CounterVisitor to the write/emit routines. The writers will then
// first invoke the visitor's NumFuncs() method to find out how many
// function's worth of data to write, then it will invoke VisitFuncs.
// The expectation is that the VisitFuncs method will then invoke the
// callback "f" with data for each function to emit to the file.
<原文结束>

# <翻译开始>
// CounterVisitor 描述了一个在编写计数器文件过程中使用的辅助对象；当编写计数器数据文件时，客户端会将一个 CounterVisitor 传递给写入/发出例程。然后，写入器首先调用访问者的 NumFuncs() 方法来确定要写入多少个函数的数据，接着调用 VisitFuncs 方法。预期是 VisitFuncs 方法将使用每个函数要写入文件的数据调用回调函数 "f"。
# <翻译结束>


<原文开始>
// CounterVisitorFn describes a callback function invoked when writing
// coverage counter data.
<原文结束>

# <翻译开始>
// CounterVisitorFn 描述了一个回调函数，当写入覆盖率计数器数据时会调用该函数。
# <翻译结束>


<原文开始>
// Write writes the contents of the count-data file to the writer
// previously supplied to NewCoverageDataWriter. Returns an error
// if something went wrong somewhere with the write.
<原文结束>

# <翻译开始>
// Write 将 count-data 文件的内容写入先前提供给 NewCoverageDataWriter 的 writer（写入器）。如果在写入过程中发生错误，将会返回错误。
# <翻译结束>


<原文开始>
	// Write string table and args to a byte slice (since we need
	// to capture offsets at various points), then emit the slice
	// once we are done.
<原文结束>

# <翻译开始>
// 将字符串表和参数写入一个字节切片（因为我们需要在各个位置捕获偏移量），然后在完成之后输出该切片。
# <翻译结束>


<原文开始>
// At this point we can now do the actual write.
<原文结束>

# <翻译开始>
// 此时，我们现在可以执行实际的写入操作。
# <翻译结束>


<原文开始>
// AppendSegment appends a new segment to a counter data, with a new
// args section followed by a payload of counter data clauses.
<原文结束>

# <翻译开始>
// AppendSegment 向计数器数据追加一个新的段，该段包含一个新的参数部分，后面跟随一组计数器数据子句作为负载。
# <翻译结束>


<原文开始>
	// Notes:
	// - this version writes everything little-endian, which means
	//   a call is needed to encode every value (expensive)
	// - we may want to move to a model in which we just blast out
	//   all counters, or possibly mmap the file and do the write
	//   implicitly.
<原文结束>

# <翻译开始>
// 注意事项：
// - 该版本采用小端序写入所有数据，这意味着
//   需要对每个值调用编码（开销较大）
// - 我们可能希望转向一种模型，直接连续输出所有计数器，
//   或者可能是将文件映射到内存并通过隐式方式执行写入。
# <翻译结束>


<原文开始>
// Write out entries for each live function.
<原文结束>

# <翻译开始>
// 为每个活动函数写出条目。
# <翻译结束>

