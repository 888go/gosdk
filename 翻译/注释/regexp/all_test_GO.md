
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
// Test empty input and/or replacement, with pattern that matches the empty string.
<原文结束>

# <翻译开始>
// 测试空输入和/或替换，使用匹配空字符串的模式。. md5:d4aefe58464a2432
# <翻译结束>


<原文开始>
// Test empty input and/or replacement, with pattern that does not match the empty string.
<原文结束>

# <翻译开始>
// 测试空输入和/或替换，使用不匹配空字符串的模式。. md5:a3fcf7444f8e4a6c
# <翻译结束>


<原文开始>
	// Multibyte characters -- verify that we don't try to match in the middle
	// of a character.
<原文结束>

# <翻译开始>
	// 多字节字符 - 验证我们不会尝试在字符的中间进行匹配。
	// md5:27a01057cc18d1c9
# <翻译结束>


<原文开始>
// Start and end of a string.
<原文结束>

# <翻译开始>
// 字符串的开始和结束。. md5:7248835ea47be3dc
# <翻译结束>


<原文开始>
// Substitution when subexpression isn't found
<原文结束>

# <翻译开始>
// 当子表达式未找到时进行替换. md5:5af0b4284e3a98a5
# <翻译结束>


<原文开始>
// Substitutions involving a (x){0}
<原文结束>

# <翻译开始>
// 包含 (x){0} 的替换操作. md5:e00d86c7da06e396
# <翻译结束>


<原文开始>
// Run ReplaceAll tests that do not have $ expansions.
<原文结束>

# <翻译开始>
// Run ReplaceAll 测试，不包含美元符号（$）的替换操作。. md5:56041c90fbae1414
# <翻译结束>


<原文开始>
// Run literal-specific tests.
<原文结束>

# <翻译开始>
// 运行特定于字面量的测试。. md5:a13e411d2c87f256
# <翻译结束>


<原文开始>
// has meta but no operator
<原文结束>

# <翻译开始>
// 包含元数据但没有操作符. md5:bf7c380f434a98f4
# <翻译结束>


<原文开始>
// has escaped operators and real operators
<原文结束>

# <翻译开始>
// 包含转义的运算符和真实的运算符. md5:2d82d20a1cf951a1
# <翻译结束>


<原文开始>
	// See golang.org/issue/11175.
	// output is unused.
<原文结束>

# <翻译开始>
	// 参见 golang.org/issue/11175。
	// output未被使用。
	// md5:746fac53787b1349
# <翻译结束>


<原文开始>
// Verify that QuoteMeta returns the expected string.
<原文结束>

# <翻译开始>
// 验证QuoteMeta返回的字符串是否符合预期。. md5:a7e1d8380394f2e8
# <翻译结束>


<原文开始>
		// Verify that the quoted string is in fact treated as expected
		// by Compile -- i.e. that it matches the original, unquoted string.
<原文结束>

# <翻译开始>
		// 验证被引用的字符串实际上是否按预期被处理
		// 通过Compile函数——即它是否与原始的、未被引用的字符串匹配。
		// md5:1b5b6be407918186
# <翻译结束>


<原文开始>
// Literal method needs to scan the pattern.
<原文结束>

# <翻译开始>
// 字面量方法需要扫描模式。. md5:6a561a10a47905f1
# <翻译结束>


<原文开始>
// The following sequence of Match calls used to panic. See issue #12980.
<原文结束>

# <翻译开始>
// 以下的Match调用序列会导致panic。参见问题#12980。. md5:2a5b2b1eaff6b92e
# <翻译结束>


<原文开始>
// Check that one-pass cutoff does trigger.
<原文结束>

# <翻译开始>
// 检查单次通过截止是否触发。. md5:760ab679ffb961b2
# <翻译结束>


<原文开始>
// Check that the same machine can be used with the standard matcher
// and then the backtracker when there are no captures.
<原文结束>

# <翻译开始>
// 检查同一台机器是否可以首先使用标准匹配器，
// 然后在没有捕获的情况下使用回溯器。
// md5:79f19d160d436e55
# <翻译结束>


<原文开始>
// The following sequence of Match calls used to panic. See issue #10319.
<原文结束>

# <翻译开始>
// 下列的Match调用序列曾经会导致恐慌。参见问题#10319。. md5:eb11f2f6fd2dbc34
# <翻译结束>


<原文开始>
// triggers standard matcher
<原文结束>

# <翻译开始>
// 触发标准匹配器. md5:02a663df6450df58
# <翻译结束>


<原文开始>
	// 'b' is between 'a' and 'c', so the charclass
	// range checking is no help here.
<原文结束>

# <翻译开始>
	// 'b' 在 'a' 和 'c' 之间，所以字符类范围检查在这里没有帮助。
	// md5:0e51b2c00d67f989
# <翻译结束>


<原文开始>
// has always been true, since Go 1.
<原文结束>

# <翻译开始>
// 自Go 1以来一直为真。. md5:d865a8317f116dbb
# <翻译结束>

