
<原文开始>
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
//版权所有2009年Go作者。所有权利保留。
//使用此源代码受BSD风格
//可以在LICENSE文件中找到的许可证。
// md5:2e9dc81828a3be8a
# <翻译结束>


<原文开始>
// For ease of reading, the test cases use strings that are converted to byte
// slices before invoking the functions.
<原文结束>

# <翻译开始>
// 为了便于阅读，测试用例使用在调用函数之前转换为字节切片的字符串。
// md5:cc0119a5cedee1b2
# <翻译结束>


<原文开始>
// Run the tests and check for allocation at the same time.
<原文结束>

# <翻译开始>
// 同时运行测试并检查内存分配情况。. md5:7eeb37e33320936f
# <翻译结束>


<原文开始>
// randomish but deterministic data
<原文结束>

# <翻译开始>
// 随机但确定的数据. md5:74d5619b89612fd4
# <翻译结束>


<原文开始>
// make sure Equal returns false for minimally different strings. The data
// is all zeros except for a single one in one location.
<原文结束>

# <翻译开始>
// 确保Equal函数对于稍微不同的字符串返回false。数据全部为零，
// 除了在某个位置有一个一。
// md5:e67bbd551b9ae3b6
# <翻译结束>


<原文开始>
// cases with one byte strings - test IndexByte and special case in Index()
<原文结束>

# <翻译开始>
// 一个字节字符串的案例 - 测试IndexByte和Index()的特殊情况. md5:f42df8c179943cdc
# <翻译结束>


<原文开始>
// test fallback to Rabin-Karp.
<原文结束>

# <翻译开始>
// 测试使用Rabin-Karp算法的备份。. md5:44782e3f078ab24b
# <翻译结束>


<原文开始>
// Execute f on each test case.  funcName should be the name of f; it's used
// in failure reports.
<原文结束>

# <翻译开始>
// 对每个测试用例执行f。funcName应该是f的名称；它在失败报告中使用。
// md5:c1927b781dfa03e1
# <翻译结束>


<原文开始>
// case for function Index.
<原文结束>

# <翻译开始>
// 用于函数 Index 的情况。. md5:3908461a593ff220
# <翻译结束>


<原文开始>
// case for function LastIndex.
<原文结束>

# <翻译开始>
// 对LastIndex函数的案例。. md5:110e5bc369debd5a
# <翻译结束>


<原文开始>
// something in the middle
<原文结束>

# <翻译开始>
// 中间的一些东西. md5:33e6ce4fb316248e
# <翻译结束>


<原文开始>
// test a larger buffer with different sizes and alignments
<原文结束>

# <翻译开始>
// 测试具有不同大小和对齐方式的较大缓冲区. md5:bdd23843bfb9fd53
# <翻译结束>


<原文开始>
// different start alignments
<原文结束>

# <翻译开始>
// 不同的起始对齐方式. md5:3c1d2601274622bb
# <翻译结束>


<原文开始>
// different end alignments
<原文结束>

# <翻译开始>
// 不同的结束对齐方式. md5:03ebb91baeab54f1
# <翻译结束>


<原文开始>
// different start and end alignments
<原文结束>

# <翻译开始>
// 不同的开始和结束对齐方式. md5:9c3f4acdc0b00b1c
# <翻译结束>


<原文开始>
// test a small index across all page offsets
<原文结束>

# <翻译开始>
// 测试小索引在所有页面偏移上的情况. md5:b0600eaff691216a
# <翻译结束>


<原文开始>
// Make sure we find the correct byte even when straddling a page.
<原文结束>

# <翻译开始>
// 确保我们即使在跨越页面的情况下也能找到正确的字节。. md5:91c3ac76aa5ffb4b
# <翻译结束>


<原文开始>
// Make sure matches outside the slice never trigger.
<原文结束>

# <翻译开始>
// 确保切片外部的匹配项永远不会触发。. md5:77a9d5f02b1924b6
# <翻译结束>


<原文开始>
// RuneError should match any invalid UTF-8 byte sequence.
<原文结束>

# <翻译开始>
// RuneError应该匹配任何无效的UTF-8字节序列。. md5:435d81d1f5f4f637
# <翻译结束>


<原文开始>
// Invalid rune values should never match.
<原文结束>

# <翻译开始>
// 无效的字符值应当永不匹配。. md5:e2994b05bbd0f2e8
# <翻译结束>


<原文开始>
// test count of a single byte across page offsets
<原文结束>

# <翻译开始>
// 测试单个字节在页面偏移量上的计数. md5:e62bbb3f055c41a3
# <翻译结束>


<原文开始>
// Make sure we don't count bytes outside our window
<原文结束>

# <翻译开始>
// 确保我们不会计算窗口之外的字节. md5:eb48fa1ed06414ad
# <翻译结束>


<原文开始>
// Fill the window with non-match
<原文结束>

# <翻译开始>
// 用非匹配填充窗口. md5:ee183c337a4b466f
# <翻译结束>


<原文开始>
// Try to find something that doesn't exist
<原文结束>

# <翻译开始>
// 尝试查找不存在的项. md5:8e7643a6e64b7d13
# <翻译结束>


<原文开始>
// Appending to the results should not change future results.
<原文结束>

# <翻译开始>
// 向结果中追加内容不应该改变未来的结果。. md5:96afbc6aa8af6e48
# <翻译结束>


<原文开始>
// Test case for any function which accepts and returns a byte slice.
// For ease of creation, we write the input byte slice as a string.
<原文结束>

# <翻译开始>
// 用于测试任何接受并返回字节切片的函数。为了方便创建，我们将输入字节切片写为字符串。
// md5:0a5359f1b96a57fc
# <翻译结束>


<原文开始>
// grows one byte per char
<原文结束>

# <翻译开始>
// 每个字符增长1字节. md5:41ef86d4dc86aeae
# <翻译结束>


<原文开始>
// test utf8.RuneSelf and utf8.MaxRune
<原文结束>

# <翻译开始>
// 测试 utf8.RuneSelf 和 utf8.MaxRune. md5:9be232d2f4ce5636
# <翻译结束>


<原文开始>
// shrinks one byte per char
<原文结束>

# <翻译开始>
// 每个字符缩小一个字节. md5:61f8c9702d4cc969
# <翻译结束>


<原文开始>
// User-defined self-inverse mapping function
<原文结束>

# <翻译开始>
// 用户自定义的自逆映射函数. md5:3f01f9afc8855318
# <翻译结束>


<原文开始>
// Run a couple of awful growth/shrinkage tests
<原文结束>

# <翻译开始>
// 运行一些糟糕的生长/收缩测试. md5:92178dba5f26024e
# <翻译结束>


<原文开始>
// 1.  Grow. This triggers two reallocations in Map.
<原文结束>

# <翻译开始>
// 1. 扩容。这将触发映射(Map)中的两次重新分配。. md5:2251939b1fa5eb1b
# <翻译结束>


<原文开始>
// Tests for results over the chunkLimit
<原文结束>

# <翻译开始>
// 对超过chunkLimit限制的结果进行测试. md5:f1cb95f92cb7c08d
# <翻译结束>


<原文开始>
// See Issue golang.org/issue/16237
<原文结束>

# <翻译开始>
// 请参阅问题 golang.org/issue/16237. md5:cb3e34cc0fbca988
# <翻译结束>


<原文开始>
// can only test reassembly if we didn't lose information
<原文结束>

# <翻译开始>
// 只有当我们没有丢失信息时，才能测试重装. md5:2c1bd20acd2e952a
# <翻译结束>


<原文开始>
	// The nils returned by TrimLeftFunc are odd behavior, but we need
	// to preserve backwards compatibility.
<原文结束>

# <翻译开始>
// TrimLeftFunc 返回的nils表现出奇特的行为，但我们需要保持向后兼容性。
// md5:b1a69a537dcc0f5f
# <翻译结束>


<原文开始>
// last rune in space is 3 bytes
<原文结束>

# <翻译开始>
// space中的最后一个字符是3个字节. md5:750aad870675d5f5
# <翻译结束>


<原文开始>
// Input is ~10% space, ~10% 2-byte UTF-8, rest ASCII non-space.
<原文结束>

# <翻译开始>
// 输入大约是10%的空格，约10%的2字节UTF-8编码，其余是ASCII非空格字符。. md5:d6a20654cb2dd72f
# <翻译结束>


<原文开始>
// Input is ~10% space, rest ASCII non-space.
<原文结束>

# <翻译开始>
// 输入大约是10%的空格，其余是ASCII非空格字符。. md5:2595a4bcaddb9459
# <翻译结束>

