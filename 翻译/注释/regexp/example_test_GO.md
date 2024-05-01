
<原文开始>
// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// 版权所有 2013 Go 作者。保留所有权利。
// 使用此源代码受BSD风格许可证约束，该许可证可从LICENSE文件中找到。
// md5:19d1a3ed91182ee4
# <翻译结束>


<原文开始>
	// Compile the expression once, usually at init time.
	// Use raw strings to avoid having to quote the backslashes.
<原文结束>

# <翻译开始>
// 通常在初始化时编译表达式。使用原始字符串以避免需要转义反斜杠。
// md5:271c082b9d8a6c71
# <翻译结束>


<原文开始>
	// Indices:
	//    01234567   012345678
	//    -ab-axb-   -axxb-ab-
<原文结束>

# <翻译开始>
// 索引：
//    01234567   012345678
//    ---a-xb---  ---x-xb-a---
// 
// 这段注释说明了索引的分布情况。在第一行中，从第3个位置（索引为2）开始到第6个位置（索引为5）是空的（标记为"-"），然后是字母"a"，接着是"x"和"b"，再接着又是空的。在第二行中，同样的模式，但这次是从第7个位置开始（索引为6）有"x"，然后是两个"x"，接着是字母"b"，最后再次是空的。
// md5:283dca55f0118208
# <翻译结束>


<原文开始>
// Regex pattern captures "key: value" pair from the content.
<原文结束>

# <翻译开始>
// 正则表达式模式用于从内容中捕获 "键: 值" 对。. md5:7e673e91ccc76b97
# <翻译结束>


<原文开始>
	// Template to convert "key: value" to "key=value" by
	// referencing the values captured by the regex pattern.
<原文结束>

# <翻译开始>
// 模板用于通过引用正则表达式模式捕获的值将 "key: value" 转换为 "key=value"。
// md5:5bc55b3b185255ad
# <翻译结束>


<原文开始>
// For each match of the regex in the content.
<原文结束>

# <翻译开始>
// 对内容中的正则表达式匹配项进行遍历。. md5:51840a927103ba23
# <翻译结束>


<原文开始>
		// Apply the captured submatches to the template and append the output
		// to the result.
<原文结束>

# <翻译开始>
// 将捕获的子模式应用到模板中，并将输出追加到结果中。
// md5:0c0b9de7b94ce6a9
# <翻译结束>

