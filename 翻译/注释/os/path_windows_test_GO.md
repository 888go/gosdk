
<原文开始>
// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// 版权所有 2016 The Go 作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。
# <翻译结束>


<原文开始>
	// 248 is long enough to trigger the longer-than-248 checks in
	// fixLongPath, but short enough not to make a path component
	// longer than 255, which is illegal on Windows. (which
	// doesn't really matter anyway, since this is purely a string
	// function we're testing, and it's not actually being used to
	// do a system call)
<原文结束>

# <翻译开始>
	// 248足够长，可以触发fixLongPath中的超过248个字符的检查，
	// 但又不至于使路径组件超过Windows系统中限制的255个字符。
	// （这实际上并不重要，因为我们正在测试的是一个纯字符串函数，
	// 并不会真正用于执行系统调用）
# <翻译结束>


<原文开始>
		// The "long" substring is replaced by a looooooong
		// string which triggers the rewriting. Except in the
		// cases below where it doesn't.
<原文结束>

# <翻译开始>
		// 将“long”子串替换为一个非常长的字符串，从而触发重写。但在下面的情况下不会进行替换（除非特别注明）。
# <翻译结束>


<原文开始>
// Create a unique-enough directory name in root.
<原文结束>

# <翻译开始>
// 在根目录中创建一个足够唯一的目录名称。
# <翻译结束>

