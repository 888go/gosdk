
<原文开始>
// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// 版权所有 ? 2021 Go作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。
# <翻译结束>


<原文开始>
// Test that FuncPC* can get correct function PC.
<原文结束>

# <翻译开始>
// 测试FuncPC*能够获取正确的函数程序计数器（PC）地址。
# <翻译结束>


<原文开始>
// Test FuncPC for locally defined function
<原文结束>

# <翻译开始>
// 测试对本地定义函数的FuncPC
# <翻译结束>


<原文开始>
// Test FuncPC for imported function
<原文结束>

# <翻译开始>
// 测试FuncPC对于导入函数的使用
# <翻译结束>


<原文开始>
// Test that FuncPC* on a function of a mismatched ABI is rejected.
<原文结束>

# <翻译开始>
// 测试FuncPC*对不匹配ABI（应用程序二进制接口）的函数是否会被拒绝。
# <翻译结束>


<原文开始>
	// We want to test internal package, which we cannot normally import.
	// Run the assembler and compiler manually.
<原文结束>

# <翻译开始>
// 我们想要测试内部包，但正常情况下无法直接导入。
// 因此，手动运行汇编器和编译器。
# <翻译结束>


<原文开始>
// parse assembly code for symabi.
<原文结束>

# <翻译开始>
// 为 symabi 解析汇编代码。
# <翻译结束>


<原文开始>
// Expect errors in line 17, 18, 20, no errors on other lines.
<原文结束>

# <翻译开始>
// 预期第17行、第18行、第20行存在错误，其他行无错误。
# <翻译结束>

