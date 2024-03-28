
<原文开始>
// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// 版权所有 ? 2022 Go作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。
# <翻译结束>


<原文开始>
// This package provides apis for producing human-readable summaries
// of coverage data (e.g. a coverage percentage for a given package or
// set of packages) and for writing data in the legacy test format
// emitted by "go test -coverprofile=<outfile>".
//
// The model for using these apis is to create a Formatter object,
// then make a series of calls to SetPackage and AddUnit passing in
// data read from coverage meta-data and counter-data files. E.g.
//
//		myformatter := cformat.NewFormatter()
//		...
//		for each package P in meta-data file: {
//			myformatter.SetPackage(P)
//			for each function F in P: {
//				for each coverable unit U in F: {
//					myformatter.AddUnit(U)
//				}
//			}
//		}
//		myformatter.EmitPercent(os.Stdout, "")
//		myformatter.EmitTextual(somefile)
//
// These apis are linked into tests that are built with "-cover", and
// called at the end of test execution to produce text output or
// emit coverage percentages.
<原文结束>

# <翻译开始>
// 该包提供了生成覆盖率数据（例如，给定包或一组包的覆盖率百分比）的人类可读摘要以及将数据写入由“go test -coverprofile=<outfile>”产生的传统测试格式的API。
// 
// 使用这些API的模型是创建一个Formatter对象，然后按照从覆盖率元数据和计数器数据文件中读取的数据，进行一系列对SetPackage和AddUnit的调用。例如：
// 
// ```go
// myformatter := cformat.NewFormatter()
// ...
// 遍历元数据文件中的每个包P
// for each package P in meta-data file {
//     myformatter.SetPackage(P)
// 遍历包P中的每个函数F
//     for each function F in P {
// 遍历函数F中的每个可覆盖单元U
//         for each coverable unit U in F {
//             myformatter.AddUnit(U)
//         }
//     }
// }
// myformatter.EmitPercent(os.Stdout, "")  // 输出覆盖率百分比
// myformatter.EmitTextual(somefile)       // 输出文本形式的覆盖率数据
// 
// ```
// 
// 这些API与使用“-cover”构建的测试相关联，并在测试执行结束时调用，以生成文本输出或计算并输出覆盖率百分比。
# <翻译结束>


<原文开始>
// Maps import path to package state.
<原文结束>

# <翻译开始>
// 将导入路径映射到包状态。
# <翻译结束>


<原文开始>
// Records current package being visited.
<原文结束>

# <翻译开始>
// 记录当前正在访问的包。
# <翻译结束>


<原文开始>
// Pointer to current package state.
<原文结束>

# <翻译开始>
// 指向当前包状态的指针。
# <翻译结束>


<原文开始>
// pstate records package-level coverage data state:
// - a table of functions (file/fname/literal)
// - a map recording the index/ID of each func encountered so far
// - a table storing execution count for the coverable units in each func
<原文结束>

# <翻译开始>
// pstate 记录包级别的覆盖率数据状态：
// - 一个包含函数信息的表格（文件名/函数名/字面量）
// - 一个映射表，记录迄今为止遇到的每个函数的索引/ID
// - 一个表格，存储每个函数中可执行单元的执行计数
# <翻译结束>


<原文开始>
// slice of unique functions
<原文结束>

# <翻译开始>
// 唯一函数的切片
# <翻译结束>


<原文开始>
// maps function to index in slice above (index acts as function ID)
<原文结束>

# <翻译开始>
// 将函数映射到上面切片中的索引（索引充当函数ID）
# <翻译结束>


<原文开始>
// A table storing coverage counts for each coverable unit.
<原文结束>

# <翻译开始>
// 一个存储每个可覆盖单元覆盖率计数的表格。
# <翻译结束>


<原文开始>
// extcu encapsulates a coverable unit within some function.
<原文结束>

# <翻译开始>
// extcu在某个函数中封装了一个可覆盖的单元。
# <翻译结束>


<原文开始>
// index into p.funcs slice
<原文结束>

# <翻译开始>
// 索引p.funcs切片
# <翻译结束>


<原文开始>
// fnfile is a function-name/file-name tuple.
<原文结束>

# <翻译开始>
// fnfile 是一个函数名称/文件名称的元组。
# <翻译结束>


<原文开始>
// SetPackage tells the formatter that we're about to visit the
// coverage data for the package with the specified import path.
// Note that it's OK to call SetPackage more than once with the
// same import path; counter data values will be accumulated.
<原文结束>

# <翻译开始>
// SetPackage 告诉格式化器我们将要访问指定导入路径的包的覆盖率数据。
// 请注意，多次调用 SetPackage 并传入相同的导入路径是没问题的；计数器数据值将会被累加。
# <翻译结束>


<原文开始>
// AddUnit passes info on a single coverable unit (file, funcname,
// literal flag, range of lines, and counter value) to the formatter.
// Counter values will be accumulated where appropriate.
<原文结束>

# <翻译开始>
// AddUnit 将可覆盖单元（文件、函数名、字面标志、行范围和计数器值）的信息传递给格式化程序。在适当的地方，计数器值将被累计。
# <翻译结束>


<原文开始>
// Use saturating arithmetic.
<原文结束>

# <翻译开始>
// 使用饱和算术。
# <翻译结束>


<原文开始>
// sortUnits sorts a slice of extcu objects in a package according to
// source position information (e.g. file and line). Note that we don't
// include function name as part of the sorting criteria, the thinking
// being that is better to provide things in the original source order.
<原文结束>

# <翻译开始>
// sortUnits 按照源位置信息（如文件和行号）对包中的 extcu 对象切片进行排序。请注意，我们没有将函数名称作为排序标准的一部分，其思路是按照原始源代码顺序提供内容更好。
# <翻译结束>


<原文开始>
		// NB: not taking function literal flag into account here (no
		// need, since other fields are guaranteed to be distinct).
<原文结束>

# <翻译开始>
// 注意：此处并未考虑函数字面量标志（不需要，因为其他字段保证是不同的）。
# <翻译结束>


<原文开始>
// EmitTextual writes the accumulated coverage data in the legacy
// cmd/cover text format to the writer 'w'. We sort the data items by
// importpath, source file, and line number before emitting (this sorting
// is not explicitly mandated by the format, but seems like a good idea
// for repeatable/deterministic dumps).
<原文结束>

# <翻译开始>
// EmitTextual 将累积的覆盖率数据以旧版 cmd/cover 文本格式写入到 writer 'w'。在输出之前，我们按 importpath、源文件和行号对数据项进行排序（虽然这种排序在该格式中并未明确要求，但为了确保重复性和确定性输出，这样做似乎是个好主意）。
# <翻译结束>


<原文开始>
// EmitPercent writes out a "percentage covered" string to the writer 'w'.
<原文结束>

# <翻译开始>
// EmitPercent 向写入器 'w' 输出一个“覆盖率百分比”字符串。
# <翻译结束>


<原文开始>
// EmitFuncs writes out a function-level summary to the writer 'w'. A
// note on handling function literals: although we collect coverage
// data for unnamed literals, it probably does not make sense to
// include them in the function summary since there isn't any good way
// to name them (this is also consistent with the legacy cmd/cover
// implementation). We do want to include their counts in the overall
// summary however.
<原文结束>

# <翻译开始>
// EmitFuncs 将函数级别的摘要写入到 writer 'w'。关于处理函数字面量的说明：尽管我们为未命名字面量收集覆盖率数据，但在函数摘要中包含它们可能没有意义，因为没有好的方法来命名它们（这也与旧版 cmd/cover 实现保持一致）。然而，我们确实希望在总体摘要中包含它们的计数。
# <翻译结束>


<原文开始>
// Emit functions for each package, sorted by import path.
<原文结束>

# <翻译开始>
// 根据导入路径为每个包生成函数，并进行排序。
# <翻译结束>


<原文开始>
		// Within a package, sort the units, then walk through the
		// sorted array. Each time we hit a new function, emit the
		// summary entry for the previous function, then make one last
		// emit call at the end of the loop.
<原文结束>

# <翻译开始>
// 在包内，先对单元进行排序，然后遍历已排序的数组。每次遇到新的函数时，输出前一个函数的摘要条目，在循环结束时再做一次最终的输出调用。
# <翻译结束>


<原文开始>
			// Don't emit entries for function literals (see discussion
			// in function header comment above).
<原文结束>

# <翻译开始>
// 不要为函数字面量输出条目（请参见上方函数头部注释中的讨论）。
# <翻译结束>


<原文开始>
// New function; emit entry for previous one.
<原文结束>

# <翻译开始>
// 新函数；为上一个函数生成入口记录。
# <翻译结束>

