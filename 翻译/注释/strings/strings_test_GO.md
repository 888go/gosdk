
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
// cases with one byte strings - test special case in Index()
<原文结束>

# <翻译开始>
// 对于包含单字节字符串的场景——测试Index()中的特殊情况
# <翻译结束>


<原文开始>
// test special cases in Index() for short strings
<原文结束>

# <翻译开始>
// 测试Index()函数在处理短字符串时的特殊情况
# <翻译结束>


<原文开始>
// test fallback to Rabin-Karp.
<原文结束>

# <翻译开始>
// 测试回退到 Rabin-Karp 算法
# <翻译结束>


<原文开始>
// Execute f on each test case.  funcName should be the name of f; it's used
// in failure reports.
<原文结束>

# <翻译开始>
// 对每个测试用例执行f。funcName应为f的名称，它用于失败报告中。
# <翻译结束>


<原文开始>
// something in the middle
<原文结束>

# <翻译开始>
// 中间的某物
# <翻译结束>


<原文开始>
// RuneError should match any invalid UTF-8 byte sequence.
<原文结束>

# <翻译开始>
// RuneError 应匹配任何无效的 UTF-8 字节序列。
# <翻译结束>


<原文开始>
// Invalid rune values should never match.
<原文结束>

# <翻译开始>
// 无效的rune值应当永远不会匹配。
# <翻译结束>


<原文开始>
// Test case for any function which accepts and returns a single string.
<原文结束>

# <翻译开始>
// 用于测试任何接受并返回单个字符串的函数的用例
# <翻译结束>


<原文开始>
// grows one byte per char
<原文结束>

# <翻译开始>
// 每个字符增长一个字节
# <翻译结束>


<原文开始>
// test utf8.RuneSelf and utf8.MaxRune
<原文结束>

# <翻译开始>
// 测试utf8.RuneSelf和utf8.MaxRune
# <翻译结束>


<原文开始>
// shrinks one byte per char
<原文结束>

# <翻译开始>
// 每个字符缩小一个字节
# <翻译结束>


<原文开始>
// User-defined self-inverse mapping function
<原文结束>

# <翻译开始>
// 用户自定义的自逆映射函数
# <翻译结束>


<原文开始>
// Run a couple of awful growth/shrinkage tests
<原文结束>

# <翻译开始>
// 运行几个糟糕的增长/缩小测试
# <翻译结束>


<原文开始>
// 1.  Grow. This triggers two reallocations in Map.
<原文结束>

# <翻译开始>
// 1. 增长。这将在Map中触发两次重新分配。
# <翻译结束>


<原文开始>
// 7. Handle invalid UTF-8 sequence
<原文结束>

# <翻译开始>
// 7. 处理无效的UTF-8序列
# <翻译结束>


<原文开始>
// 8. Check utf8.RuneSelf and utf8.MaxRune encoding
<原文结束>

# <翻译开始>
// 8. 检查utf8.RuneSelf和utf8.MaxRune的编码
# <翻译结束>


<原文开始>
// 9. Check mapping occurs in the front, middle and back
<原文结束>

# <翻译开始>
// 9. 检查映射发生在前部、中部和后部
# <翻译结束>


<原文开始>
// last rune in space is 3 bytes
<原文结束>

# <翻译开始>
// 空格中的最后一个rune占3个字节
# <翻译结束>


<原文开始>
// Make a string of all the runes.
<原文结束>

# <翻译开始>
// 生成包含所有字符的字符串
# <翻译结束>


<原文开始>
// Tests for results over the chunkLimit
<原文结束>

# <翻译开始>
// 测试超过chunkLimit的结果
# <翻译结束>


<原文开始>
// See Issue golang.org/issue/16237
<原文结束>

# <翻译开始>
// 参见问题 golang.org/issue/16237
# <翻译结束>


<原文开始>
// can only test reassembly if we didn't lose information
<原文结束>

# <翻译开始>
// 只有在未丢失信息的情况下，才能测试重组
# <翻译结束>


<原文开始>
	// cases to cover code in runtime/asm_amd64.s:indexShortStr
	// 2-byte needle
<原文结束>

# <翻译开始>
	// 用于覆盖runtime/asm_amd64.s中indexShortStr代码的测试用例
	// 2字节 needle
# <翻译结束>


<原文开始>
// Input is ~10% space, ~10% 2-byte UTF-8, rest ASCII non-space.
<原文结束>

# <翻译开始>
// 输入中约10%为空格，约10%为2字节UTF-8编码，其余为非空格ASCII字符
# <翻译结束>


<原文开始>
// Input is ~10% space, rest ASCII non-space.
<原文结束>

# <翻译开始>
// 输入中约有 10% 的空格，其余部分为非空格 ASCII 字符。
# <翻译结束>

