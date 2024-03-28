
<原文开始>
// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// 版权所有 ? 2018 The Go Authors。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。
# <翻译结束>


<原文开始>
// The declarations below generate ABI wrappers for functions
// implemented in assembly in this package but declared in another
// package.
<原文结束>

# <翻译开始>
// 下面的声明为函数生成ABI（应用程序二进制接口）包装器，这些函数在此包中以汇编实现，但在另一个包中声明。
# <翻译结束>


<原文开始>
// The compiler generates calls to runtime.memequal and runtime.memequal_varlen.
// In addition, the runtime calls runtime.memequal explicitly.
// Those functions are implemented in this package.
<原文结束>

# <翻译开始>
// 编译器生成对runtime.memequal和runtime.memequal_varlen的调用。
// 此外，运行时还会明确地调用runtime.memequal函数。
// 这些函数在此包中实现。
# <翻译结束>

