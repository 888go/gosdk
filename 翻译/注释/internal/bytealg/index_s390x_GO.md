
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
	// Note: we're kind of lucky that this flag is available at this point.
	// The runtime sets HasVX when processing auxv records, and that happens
	// to happen *before* running the init functions of packages that
	// the runtime depends on.
	// TODO: it would really be nicer for internal/cpu to figure out this
	// flag by itself. Then we wouldn't need to depend on quirks of
	// early startup initialization order.
<原文结束>

# <翻译开始>
// 注意：我们相当幸运在这个点上能够访问到这个标志。在处理auxv记录时，runtime会设置HasVX标志，而这一过程恰好发生在运行runtime所依赖的包的初始化函数之前。
// 待办事项（TODO）：最好是让internal/cpu自行确定这个标志。这样我们就无需依赖早期启动初始化顺序的特性了。
# <翻译结束>


<原文开始>
// Cutover reports the number of failures of IndexByte we should tolerate
// before switching over to Index.
// n is the number of bytes processed so far.
// See the bytes.Index implementation for details.
<原文结束>

# <翻译开始>
// Cutover 报告我们在切换到 Index 之前应容忍的 IndexByte 失败次数。
// n 是迄今为止已处理的字节数。
// 有关详细信息，请参阅 bytes.Index 的实现。
# <翻译结束>


<原文开始>
// 1 error per 8 characters, plus a few slop to start.
<原文结束>

# <翻译开始>
// 每8个字符出现1个错误，再加上一些初始误差。
# <翻译结束>

