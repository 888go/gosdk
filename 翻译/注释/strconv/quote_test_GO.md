
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
// Verify that our IsPrint agrees with unicode.IsPrint.
<原文结束>

# <翻译开始>
// 验证我们的IsPrint方法是否与unicode.IsPrint一致。. md5:5e874793f8fc1f0e
# <翻译结束>


<原文开始>
// Verify that our IsGraphic agrees with unicode.IsGraphic.
<原文结束>

# <翻译开始>
// 验证我们的IsGraphic方法与unicode.IsGraphic是否一致。. md5:8e209dd71d1c6c1c
# <翻译结束>


<原文开始>
// Some non-printable but graphic runes. Final column is double-quoted.
<原文结束>

# <翻译开始>
// 一些非可打印但可视的字符。最后一列使用双引号括起。. md5:ecc063c1010c3543
# <翻译结束>


<原文开始>
// Some differences between graphic and printable. Note the last column is double-quoted.
<原文结束>

# <翻译开始>
// graphic和printable之间的一些差异。请注意最后一列是使用双引号的。. md5:2b885147fc3d3a86
# <翻译结束>


<原文开始>
// Issue 23685: invalid UTF-8 should not go through the fast path.
<原文结束>

# <翻译开始>
// 问题 23685：无效的 UTF-8 不应该通过快速路径处理。. md5:3e0c097fad6333f0
# <翻译结束>


<原文开始>
	// Test QuotedPrefix.
	// Adding an arbitrary suffix should not change the result of QuotedPrefix
	// assume that the suffix doesn't accidentally terminate a truncated input.
<原文结束>

# <翻译开始>
	// 测试QuotedPrefix。
	// 添加任意后缀不应该改变QuotedPrefix的结果
	// 假设后缀不会意外地终止截断的输入。
	// md5:95e27f90457e7585
# <翻译结束>


<原文开始>
// special characters for quoted strings
<原文结束>

# <翻译开始>
// 引用字符串中的特殊字符. md5:feb054bcd5874ca7
# <翻译结束>


<原文开始>
// original input had trailing junk, reparse with only valid prefix
<原文结束>

# <翻译开始>
// 原始输入包含多余的尾部内容，使用有效的前缀重新解析. md5:8621be0d2464073a
# <翻译结束>

