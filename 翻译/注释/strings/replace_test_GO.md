
<原文开始>
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// 版权所有 ? 2009 Go作者。保留所有权利。
// 本源代码的使用受BSD风格
// 许可证约束，该许可证可在LICENSE文件中找到。
# <翻译结束>


<原文开始>
// The http package's old HTML escaping function.
<原文结束>

# <翻译开始>
// http 包中的旧版 HTML 转义函数。
# <翻译结束>


<原文开始>
// TestReplacer tests the replacer implementations.
<原文结束>

# <翻译开始>
// TestReplacer 测试替换器实现。
# <翻译结束>


<原文开始>
// str converts 0xff to "\xff". This isn't just string(b) since that converts to UTF-8.
<原文结束>

# <翻译开始>
// str 将 0xff 转换为 "\xff"。这不仅仅是使用 string(b) 进行转换，因为后者会转为 UTF-8 格式。
# <翻译结束>


<原文开始>
// inc maps "\x00"->"\x01", ..., "a"->"b", "b"->"c", ..., "\xff"->"\x00".
<原文结束>

# <翻译开始>
// inc 将字符按以下规则进行映射：
//   "\x00" -> "\x01"
//   ...
//   "a" -> "b"
//   "b" -> "c"
//   ...
//   "\xff" -> "\x00"
# <翻译结束>


<原文开始>
// Test cases with 1-byte old strings, 1-byte new strings.
<原文结束>

# <翻译开始>
// 测试用例，包含1字节的旧字符串和1字节的新字符串
# <翻译结束>


<原文开始>
// repeat maps "a"->"a", "b"->"bb", "c"->"ccc", ...
<原文结束>

# <翻译开始>
// repeat 将 "a" 映射为 "a"，"b" 映射为 "bb"，"c" 映射为 "ccc"，以此类推
# <翻译结束>


<原文开始>
// Test cases with 1-byte old strings, variable length new strings.
<原文结束>

# <翻译开始>
// 测试用例，含1字节长度的旧字符串，及可变长度的新字符串。
# <翻译结束>


<原文开始>
// The remaining test cases have variable length old strings.
<原文结束>

# <翻译开始>
// 剩余的测试用例具有可变长度的旧字符串。
# <翻译结束>


<原文开始>
	// gen1 has multiple old strings of variable length. There is no
	// overall non-empty common prefix, but some pairwise common prefixes.
<原文结束>

# <翻译开始>
// gen1 包含多个长度可变的旧字符串。它们之间不存在非空的整体公共前缀，但存在一些两两之间的公共前缀。
# <翻译结束>


<原文开始>
// gen2 has multiple old strings with no pairwise common prefix.
<原文结束>

# <翻译开始>
// gen2包含多个无两两共同前缀的旧字符串。
# <翻译结束>


<原文开始>
// gen3 has multiple old strings with an overall common prefix.
<原文结束>

# <翻译开始>
// gen3 含有多个具有共同前缀的旧字符串
# <翻译结束>


<原文开始>
	// foo{1,2,3,4} have multiple old strings with an overall common prefix
	// and 1- or 2- byte extensions from the common prefix.
<原文结束>

# <翻译开始>
// foo{1,2,3,4}具有多个具有共同前缀的旧字符串，且从该公共前缀延伸出1个或2个字节。
# <翻译结束>


<原文开始>
// genAll maps "\x00\x01\x02...\xfe\xff" to "[all]", amongst other things.
<原文结束>

# <翻译开始>
// genAll 将 "\x00\x01\x02...\xfe\xff" 映射为 "[all]"，同时执行其他操作。
# <翻译结束>


<原文开始>
// Test cases with empty old strings.
<原文结束>

# <翻译开始>
// 测试用例，其中包含空的旧字符串。
# <翻译结束>


<原文开始>
// Issue 6659 cases (more single string replacer)
<原文结束>

# <翻译开始>
// Issue 6659 示例（更多单字符串替换器）
# <翻译结束>


<原文开始>
// TestPickAlgorithm tests that NewReplacer picks the correct algorithm.
<原文结束>

# <翻译开始>
// TestPickAlgorithm 测试 NewReplacer 选择正确算法的情况
# <翻译结束>


<原文开始>
// TestWriteStringError tests that WriteString returns an error
// received from the underlying io.Writer.
<原文结束>

# <翻译开始>
// TestWriteStringError 测试 WriteString 函数在接收到底层 io.Writer 传回的错误时能够正常返回该错误
# <翻译结束>


<原文开始>
// TestGenericTrieBuilding verifies the structure of the generated trie. There
// is one node per line, and the key ending with the current line is in the
// trie if it ends with a "+".
<原文结束>

# <翻译开始>
// TestGenericTrieBuilding 验证生成的字典树结构。每一行代表一个节点，若当前行以 "+" 结尾，则表示以该行结束的键存在于字典树中。
# <翻译结束>


<原文开始>
// Remove tabs from tc.out
<原文结束>

# <翻译开始>
// 从tc.out中移除制表符
# <翻译结束>


<原文开始>
// varying lengths forces generic
<原文结束>

# <翻译开始>
// 不同长度强制使用泛型
# <翻译结束>


<原文开始>
// BenchmarkByteByteReplaces compares byteByteImpl against multiple Replaces.
<原文结束>

# <翻译开始>
// BenchmarkByteByteReplaces对比byteByteImpl与多个Replaces的性能。
# <翻译结束>


<原文开始>
// BenchmarkByteByteMap compares byteByteImpl against Map.
<原文结束>

# <翻译开始>
// BenchmarkByteByteMap 对比 byteByteImpl 与 Map 的性能
# <翻译结束>

