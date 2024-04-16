
<原文开始>
// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// 版权所有 ? 2018 The Go 作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可协议约束，
// 该协议可在 LICENSE 文件中找到。
# <翻译结束>


<原文开始>
// Package txtar implements a trivial text-based file archive format.
//
// The goals for the format are:
//
//   - be trivial enough to create and edit by hand.
//   - be able to store trees of text files describing go command test cases.
//   - diff nicely in git history and code reviews.
//
// Non-goals include being a completely general archive format,
// storing binary data, storing file modes, storing special files like
// symbolic links, and so on.
//
// # Txtar format
//
// A txtar archive is zero or more comment lines and then a sequence of file entries.
// Each file entry begins with a file marker line of the form "-- FILENAME --"
// and is followed by zero or more file content lines making up the file data.
// The comment or file content ends at the next file marker line.
// The file marker line must begin with the three-byte sequence "-- "
// and end with the three-byte sequence " --", but the enclosed
// file name can be surrounding by additional white space,
// all of which is stripped.
//
// If the txtar file is missing a trailing newline on the final line,
// parsers should consider a final newline to be present anyway.
//
// There are no possible syntax errors in a txtar archive.
<原文结束>

# <翻译开始>
// Package txtar 实现了一个简单的基于文本的文件归档格式。
//
// 该格式的目标如下：
//
//   - 简单到可以手动创建和编辑。
//   - 能够存储描述 go 命令测试用例的文本文件树。
//   - 在 git 历史记录和代码审查中易于进行 diff 比较。
//
// 非目标包括成为一个完全通用的归档格式、存储二进制数据、存储文件模式、存储符号链接等特殊文件等。
//
// # Txtar 格式
//
// 一个 txtar 归档由零个或多个注释行以及随后的一系列文件条目组成。每个文件条目以形如 "-- FILENAME --" 的文件标记行开始，后面跟随零个或多个构成文件数据的文件内容行。注释或文件内容在下一个文件标记行处结束。文件标记行必须以三字节序列 "-- " 开头，并以三字节序列 " --" 结尾，但其中包含的文件名两侧可以有额外的空白字符，这些空白字符都会被去除。
//
// 如果 txtar 文件在最后一行缺少尾部换行符，解析器应视其末尾存在换行符。
//
// txtar 归档中不存在任何可能的语法错误。
# <翻译结束>


<原文开始>
// An Archive is a collection of files.
<原文结束>

# <翻译开始>
// Archive 是一组文件的集合。
# <翻译结束>


<原文开始>
// A File is a single file in an archive.
<原文结束>

# <翻译开始>
// File 是档案中的单个文件。
# <翻译结束>


<原文开始>
// name of file ("foo/bar.txt")
<原文结束>

# <翻译开始>
// 文件名（"foo/bar.txt"）
# <翻译结束>


<原文开始>
// Format returns the serialized form of an Archive.
// It is assumed that the Archive data structure is well-formed:
// a.Comment and all a.File[i].Data contain no file marker lines,
// and all a.File[i].Name is non-empty.
<原文结束>

# <翻译开始>
// Format 函数返回 Archive 的序列化形式。
// 假定 Archive 数据结构是格式良好的：
// 其中 a.Comment 及所有 a.File[i].Data 不包含文件标记行，
// 并且所有 a.File[i].Name 非空。
# <翻译结束>


<原文开始>
// ParseFile parses the named file as an archive.
<原文结束>

# <翻译开始>
// ParseFile 以档案格式解析指定名称的文件。
# <翻译结束>


<原文开始>
// Parse parses the serialized form of an Archive.
// The returned Archive holds slices of data.
<原文结束>

# <翻译开始>
// Parse解析Archive的序列化形式。
// 返回的Archive持有数据切片。
# <翻译结束>


<原文开始>
// findFileMarker finds the next file marker in data,
// extracts the file name, and returns the data before the marker,
// the file name, and the data after the marker.
// If there is no next marker, findFileMarker returns before = fixNL(data), name = "", after = nil.
<原文结束>

# <翻译开始>
// findFileMarker 在 data 中查找下一个文件标记，
// 提取文件名，并返回标记前的数据、文件名以及标记后的数据。
// 若不存在下一个标记，findFileMarker 返回 before = fixNL(data)，name = ""，after = nil。
# <翻译结束>


<原文开始>
// positioned at start of new possible marker
<原文结束>

# <翻译开始>
// 位于新可能标记的起始处
# <翻译结束>


<原文开始>
// isMarker checks whether data begins with a file marker line.
// If so, it returns the name from the line and the data after the line.
// Otherwise it returns name == "" with an unspecified after.
<原文结束>

# <翻译开始>
// isMarker 检查 data 是否以文件标记行开头。
// 如果是，它将返回该行中的名称以及该行之后的数据。
// 否则返回 name == "" 以及未指定的 after。
# <翻译结束>


<原文开始>
// If data is empty or ends in \n, fixNL returns data.
// Otherwise fixNL returns a new slice consisting of data with a final \n added.
<原文结束>

# <翻译开始>
// 如果data为空或以\n结尾，fixNL函数返回data。
// 否则，fixNL函数返回一个新的切片，该切片由data加上末尾的\n组成。
# <翻译结束>

