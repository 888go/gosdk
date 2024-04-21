
<原文开始>
// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// 版权所有 ? 2015 The Go 作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。
# <翻译结束>


<原文开始>
	// We need to open a non-pollable file.
	// This is almost certainly Linux-specific,
	// but if other systems have non-pollable files,
	// we can add them here.
<原文结束>

# <翻译开始>
// 我们需要打开一个不可轮询的文件。
// 这几乎肯定是Linux特有的，
// 但如果其他系统也有不可轮询的文件，
// 我们可以在这里添加它们。
# <翻译结束>


<原文开始>
	// On a Linux laptop, before the problem was fixed,
	// this test failed about 50% of the time with this
	// number of iterations.
	// It takes about 1/2 second when it passes.
<原文结束>

# <翻译开始>
// 在一台Linux笔记本电脑上，在问题修复之前，
// 该测试在进行如此多轮迭代时大约有50%的几率失败。
// 当测试通过时，大约耗时半秒。
# <翻译结束>


<原文开始>
// The problem only occurs if we use O_NONBLOCK here.
<原文结束>

# <翻译开始>
// 问题只会在我们在这里使用O_NONBLOCK时出现。
# <翻译结束>

