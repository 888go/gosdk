
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
// Types and constants related to the output files files written
// by code coverage tooling. When a coverage-instrumented binary
// is run, it emits two output files: a meta-data output file, and
// a counter data output file.
<原文结束>

# <翻译开始>
// 与代码覆盖率工具生成的输出文件相关的类型和常量。当一个经过覆盖率检测的二进制程序运行时，它会产生两个输出文件：元数据输出文件和计数器数据输出文件。
# <翻译结束>


<原文开始>
//.....................................................................
//
// Meta-data definitions:
//
// The meta-data file is composed of a file header, a series of
// meta-data blobs/sections (one per instrumented package), and an offsets
// area storing the offsets of each section. Format of the meta-data
// file looks like:
//
// --header----------
//  | magic: [4]byte magic string
//  | version
//  | total length of meta-data file in bytes
//  | numPkgs: number of package entries in file
//  | hash: [16]byte hash of entire meta-data payload
//  | offset to string table section
//  | length of string table
//  | number of entries in string table
//  | counter mode
//  | counter granularity
//  --package offsets table------
//  <offset to pkg 0>
//  <offset to pkg 1>
//  ...
//  --package lengths table------
//  <length of pkg 0>
//  <length of pkg 1>
//  ...
//  --string table------
//  <uleb128 len> 8
//  <data> "somestring"
//  ...
//  --package payloads------
//  <meta-symbol for pkg 0>
//  <meta-symbol for pkg 1>
//  ...
//
// Each package payload is a stand-alone blob emitted by the compiler,
// and does not depend on anything else in the meta-data file. In
// particular, each blob has it's own string table. Note that the
// file-level string table is expected to be very short (most strings
// will be in the meta-data blobs themselves).
<原文结束>

# <翻译开始>
// 以下是Go注释的中文翻译：
// 
// ```
//.....................................................................
//
// 元数据定义：
//
// 元数据文件由文件头、一系列元数据块/节（每个被监测包一个）以及存储每个节偏移量的区域组成。元数据文件的格式如下所示：
//
// --文件头-----------------
//  | 魔数: [4]byte 标识符字符串
//  | 版本号
//  | 元数据文件总长度（字节数）
//  | numPkgs: 文件中包条目的数量
//  | 哈希值: [16]byte 整个元数据负载的哈希值
//  | 字符串表部分的偏移量
//  | 字符串表长度
//  | 字符串表中的条目数量
//  | 计数器模式
//  | 计数器粒度
//  --包偏移量表-----------
//  <包0的偏移量>
//  <包1的偏移量>
//  ...
//  --包长度表-----------
//  <包0的长度>
//  <包1的长度>
//  ...
//  --字符串表-----------
//  <uleb128编码的长度> 8
//  <数据> "somestring"
//  ...
//  --包有效载荷--------
//  <包0的元符号>
//  <包1的元符号>
//  ...
//
// 每个包的有效载荷是由编译器独立生成的块，不依赖于元数据文件中的其他任何内容。特别是，每个块都有其自己的字符串表。请注意，文件级的字符串表预期会非常短（大多数字符串将在元数据块自身中）。
// ```
# <翻译结束>


<原文开始>
// CovMetaMagic holds the magic string for a meta-data file.
<原文结束>

# <翻译开始>
// CovMetaMagic 存储元数据文件的魔术字符串。
# <翻译结束>


<原文开始>
// MetaFilePref is a prefix used when emitting meta-data files; these
// files are of the form "covmeta.<hash>", where hash is a hash
// computed from the hashes of all the package meta-data symbols in
// the program.
<原文结束>

# <翻译开始>
// MetaFilePref 是用于生成元数据文件时使用的前缀；此类文件的格式为 "covmeta.<哈希值>"，其中“哈希值”是通过对程序中所有包元数据符号的哈希值进行计算得出的。
# <翻译结束>


<原文开始>
// MetaFileVersion contains the current (most recent) meta-data file version.
<原文结束>

# <翻译开始>
// MetaFileVersion包含当前（最新）的元数据文件版本。
# <翻译结束>


<原文开始>
// MetaFileHeader stores file header information for a meta-data file.
<原文结束>

# <翻译开始>
// MetaFileHeader 用于存储元数据文件的文件头信息。
# <翻译结束>


<原文开始>
// MetaSymbolHeader stores header information for a single
// meta-data blob, e.g. the coverage meta-data payload
// computed for a given Go package.
<原文结束>

# <翻译开始>
// MetaSymbolHeader 用于存储单个元数据 blob 的头部信息，
// 例如，为给定的 Go 包计算出的覆盖率元数据负载。
# <翻译结束>


<原文开始>
// size of meta-symbol payload in bytes
<原文结束>

# <翻译开始>
// 元符号负载的大小，以字节为单位
# <翻译结束>


<原文开始>
// keep in sync with above
<原文结束>

# <翻译开始>
// 与上面保持同步
# <翻译结束>


<原文开始>
// As an example, consider the following Go package:
//
// 01: package p
// 02:
// 03: var v, w, z int
// 04:
// 05: func small(x, y int) int {
// 06:   v++
// 07:   // comment
// 08:   if y == 0 {
// 09:     return x
// 10:   }
// 11:   return (x << 1) ^ (9 / y)
// 12: }
// 13:
// 14: func Medium(q, r int) int {
// 15:   s1 := small(q, r)
// 16:   z += s1
// 17:   s2 := small(r, q)
// 18:   w -= s2
// 19:   return w + z
// 20: }
//
// The meta-data blob for the single package above might look like the
// following:
//
// -- MetaSymbolHeader header----------
//  | size: size of this blob in bytes
//  | packagepath: <path to p>
//  | modulepath: <modpath for p>
//  | nfiles: 1
//  | nfunctions: 2
//  --func offsets table------
//  <offset to func 0>
//  <offset to func 1>
//  --string table (contains all files and functions)------
//  | <uleb128 len> 4
//  | <data> "p.go"
//  | <uleb128 len> 5
//  | <data> "small"
//  | <uleb128 len> 6
//  | <data> "Medium"
//  --func 0------
//  | <uleb128> num units: 3
//  | <uleb128> func name: S1 (index into string table)
//  | <uleb128> file: S0 (index into string table)
//  | <unit 0>:  S0   L6     L8    2
//  | <unit 1>:  S0   L9     L9    1
//  | <unit 2>:  S0   L11    L11   1
//  --func 1------
//  | <uleb128> num units: 1
//  | <uleb128> func name: S2 (index into string table)
//  | <uleb128> file: S0 (index into string table)
//  | <unit 0>:  S0   L15    L19   5
//  ---end-----------
<原文结束>

# <翻译开始>
// 作为示例，请考虑以下 Go 包：
// 
// ```go
// 01: package p
// 定义包名为 p
// 
// 03: var v, w, z int
// 定义整型变量 v、w 和 z
// 
// 05: func small(x, y int) int {
// 定义一个名为 small 的函数，接收两个 int 类型参数 x 和 y，并返回一个 int 类型结果
// 06:   v++
// 函数体内，将变量 v 自增 1
// 07:   // comment
// 注释部分
// 08:   if y == 0 {
// 判断条件：如果 y 等于 0
// 09:     return x
// 如果条件成立，则返回 x
// 11:   return (x << 1) ^ (9 / y)
// 否则，返回表达式 (x 左移 1 位) 按位异或 (9 除以 y 的结果)
// 
// 14: func Medium(q, r int) int {
// 定义一个名为 Medium 的函数，接收两个 int 类型参数 q 和 r，并返回一个 int 类型结果
// 15:   s1 := small(q, r)
// 在函数体内，调用 small 函数并将结果赋值给 s1
// 16:   z += s1
// 将 z 的值加上 s1
// 17:   s2 := small(r, q)
// 调用 small 函数并将参数顺序交换后赋值给 s2
// 18:   w -= s2
// 将 w 的值减去 s2
// 19:   return w + z
// 返回 w 与 z 的和
// 
// 下面的描述是一种可能的元数据结构，用于表示上述单个包的信息：
// ...
// ```
// 
// 这段注释接下来详细描述了如何为上面这个 Go 包构建元数据结构，包括函数偏移量表、字符串表以及各个函数的具体单元信息等。但这些内容并非 Go 代码本身，而是对 Go 包编译后的元数据组织形式的一种说明。
# <翻译结束>


<原文开始>
// The following types and constants used by the meta-data encoder/decoder.
<原文结束>

# <翻译开始>
// 以下类型和常量用于元数据编码/解码器。
# <翻译结束>


<原文开始>
// FuncDesc encapsulates the meta-data definitions for a single Go function.
// This version assumes that we're looking at a function before inlining;
// if we want to capture a post-inlining view of the world, the
// representations of source positions would need to be a good deal more
// complicated.
<原文结束>

# <翻译开始>
// FuncDesc 封装了单个 Go 函数的元数据定义。
// 此版本假设我们在内联之前查看函数；如果我们想要捕获内联后的视图，源位置的表示方式将会需要复杂得多。
# <翻译结束>


<原文开始>
// true if this is a function literal
<原文结束>

# <翻译开始>
// 如果这是一个函数字面量，则为true
# <翻译结束>


<原文开始>
// CounterMode tracks the "flavor" of the coverage counters being
// used in a given coverage-instrumented program.
<原文结束>

# <翻译开始>
// CounterMode 跟踪在给定的已进行覆盖率统计的程序中使用的覆盖率计数器的“类型”。
# <翻译结束>


<原文开始>
// registration-only pseudo-mode
<原文结束>

# <翻译开始>
// 注册-only伪模式
# <翻译结束>


<原文开始>
// CounterGranularity tracks the granularity of the coverage counters being
// used in a given coverage-instrumented program.
<原文结束>

# <翻译开始>
// CounterGranularity 跟踪在特定的覆盖率测试程序中使用的覆盖率计数器的精度或粒度。
# <翻译结束>


<原文开始>
//.....................................................................
//
// Counter data definitions:
//
<原文结束>

# <翻译开始>
//.....................................................................
//
// 计数器数据定义：
// 
// （注：由于上下文不全，该注释可能表示以下含义：此处是关于计数器相关的数据结构或变量的定义说明）
# <翻译结束>


<原文开始>
// A counter data file is composed of a file header followed by one or
// more "segments" (each segment representing a given run or partial
// run of a give binary) followed by a footer.
<原文结束>

# <翻译开始>
// 一个计数数据文件由文件头开始，随后包含一个或多个“段”（每个段表示给定二进制运行或部分运行的结果），最后是一个结束符。
# <翻译结束>


<原文开始>
// CovCounterMagic holds the magic string for a coverage counter-data file.
<原文结束>

# <翻译开始>
// CovCounterMagic 保存了覆盖率计数器数据文件的魔术字符串。
# <翻译结束>


<原文开始>
// CounterFileVersion stores the most recent counter data file version.
<原文结束>

# <翻译开始>
// CounterFileVersion 存储最新的计数器数据文件版本。
# <翻译结束>


<原文开始>
// CounterFileHeader stores files header information for a counter-data file.
<原文结束>

# <翻译开始>
// CounterFileHeader 用于存储计数器数据文件的文件头信息。
# <翻译结束>


<原文开始>
// CounterSegmentHeader encapsulates information about a specific
// segment in a counter data file, which at the moment contains
// counters data from a single execution of a coverage-instrumented
// program. Following the segment header will be the string table and
// args table, and then (possibly) padding bytes to bring the byte
// size of the preamble up to a multiple of 4. Immediately following
// that will be the counter payloads.
//
// The "args" section of a segment is used to store annotations
// describing where the counter data came from; this section is
// basically a series of key-value pairs (can be thought of as an
// encoded 'map[string]string'). At the moment we only write os.Args()
// data to this section, using pairs of the form "argc=<integer>",
// "argv0=<os.Args[0]>", "argv1=<os.Args[1]>", and so on. In the
// future the args table may also include things like GOOS/GOARCH
// values, and/or tags indicating which tests were run to generate the
// counter data.
<原文结束>

# <翻译开始>
// CounterSegmentHeader 封装了关于计数器数据文件中特定段的信息，当前该段包含来自单次执行的覆盖率测试程序的计数器数据。紧跟在段头后面的是字符串表和参数表，然后（可能）是填充字节，以使前置段的字节大小成为4的倍数。紧随其后的是计数器有效负载。
// 
// 段的“args”部分用于存储描述计数器数据来源的注解；这部分基本上是一系列键值对（可以视为编码形式的`map[string]string`）。目前，我们仅将os.Args()数据写入此部分，采用的形式为"argc=<整数>"、"argv0=<os.Args[0]>"、"argv1=<os.Args[1]>”等。在未来，参数表可能还会包括GOOS/GOARCH值，以及表示生成计数器数据时运行了哪些测试的标签。
# <翻译结束>


<原文开始>
// CounterFileFooter appears at the tail end of a counter data file,
// and stores the number of segments it contains.
<原文结束>

# <翻译开始>
// CounterFileFooter 会出现在计数器数据文件的尾部，
// 并存储它所包含的段落数量。
# <翻译结束>


<原文开始>
// CounterFilePref is the file prefix used when emitting coverage data
// output files. CounterFileTemplate describes the format of the file
// name: prefix followed by meta-file hash followed by process ID
// followed by emit UnixNanoTime.
<原文结束>

# <翻译开始>
// CounterFilePref 是在输出覆盖率数据时使用的文件前缀。
// CounterFileTemplate 描述了文件名的格式：前缀后面跟随元文件哈希，再跟随进程ID，最后是UnixNanoTime形式的生成时间。
# <翻译结束>


<原文开始>
// CounterFlavor describes how function and counters are
// stored/represented in the counter section of the file.
<原文结束>

# <翻译开始>
// CounterFlavor 描述了函数和计数器在文件的计数器部分如何存储/表示。
# <翻译结束>


<原文开始>
	// "Raw" representation: all values (pkg ID, func ID, num counters,
	// and counters themselves) are stored as uint32's.
<原文结束>

# <翻译开始>
// "原始"表示形式：所有值（包ID、函数ID、计数器数量以及计数器本身）都作为uint32存储。
# <翻译结束>


<原文开始>
	// "ULeb" representation: all values (pkg ID, func ID, num counters,
	// and counters themselves) are stored with ULEB128 encoding.
<原文结束>

# <翻译开始>
// "ULeb" 表示形式：所有值（包ID、函数ID、计数器数量以及计数器本身）均采用 ULEB128 编码方式进行存储。
# <翻译结束>


<原文开始>
//.....................................................................
//
// Runtime counter data definitions.
//
<原文结束>

# <翻译开始>
//...
//...
// 运行时计数器数据定义。
# <翻译结束>


<原文开始>
// At runtime within a coverage-instrumented program, the "counters"
// object we associated with instrumented function can be thought of
// as a struct of the following form:
//
// struct {
//     numCtrs uint32
//     pkgid uint32
//     funcid uint32
//     counterArray [numBlocks]uint32
// }
//
// where "numCtrs" is the number of blocks / coverable units within the
// function, "pkgid" is the unique index assigned to this package by
// the runtime, "funcid" is the index of this function within its containing
// package, and "counterArray" stores the actual counters.
//
// The counter variable itself is created not as a struct but as a flat
// array of uint32's; we then use the offsets below to index into it.
<原文结束>

# <翻译开始>
// 在运行时，在一个经过覆盖率统计的程序中，与已统计函数关联的“counters”对象可以被视为具有以下形式的结构体：
// 
// ```go
// 结构体定义如下：
// struct {
//     numCtrs  uint32 // 覆盖单元/块的数量
//     pkgid    uint32 // 运行时为此包分配的唯一索引
//     funcid   uint32 // 此函数在其所在包中的索引
//     counterArray [numBlocks]uint32 // 实际计数器数组
// }
// ```
// 
// 其中，“numCtrs”表示函数内部的覆盖单元或可统计块的数量，“pkgid”是运行时为此包分配的唯一索引，“funcid”是此函数在其所在包中的索引，“counterArray”用于存储实际的计数器。
// 
// 实际上，计数器变量本身并非以结构体的形式创建，而是作为一个扁平的uint32数组；然后我们使用下面的偏移量来索引它。
# <翻译结束>

