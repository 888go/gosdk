
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
// vary the input alignment of tt.b
<原文结束>

# <翻译开始>
// 改变输入对齐方式为 tt.b. md5:ce1433df219092cd
# <翻译结束>


<原文开始>
// lengths to test in ascending order
<原文结束>

# <翻译开始>
// 以递增顺序进行测试的长度. md5:d422e60df7501431
# <翻译结束>


<原文开始>
// randomish but deterministic data. No 0 or 255.
<原文结束>

# <翻译开始>
// 随机但确定的数据。不包含0或255。. md5:60ed4a09009d8428
# <翻译结束>


<原文开始>
// data past the end is different
<原文结束>

# <翻译开始>
// 数据末尾之后的内容是不同的. md5:8d87dd9025a7c19f
# <翻译结束>


<原文开始>
	// This test compares byte slices that are almost identical, except one
	// difference that for some j, a[j]>b[j] and a[j+1]<b[j+1]. If the implementation
	// compares large chunks with wrong endianness, it gets wrong result.
	// no vector register is larger than 512 bytes for now
<原文结束>

# <翻译开始>
// 该测试比较两个几乎相同的字节切片，唯一的区别在于某个索引j处，a[j]大于b[j]并且a[j+1]小于b[j+1]。
// 如果实现方式错误地以错误的字节序比较了大块数据，它将得到错误的结果。
// 目前没有任何向量寄存器大于512字节。
// md5:4bc7ed687468d65e
# <翻译结束>

