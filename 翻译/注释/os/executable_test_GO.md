
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
// we want fn to be of the form "dir/prog"
<原文结束>

# <翻译开始>
// 我们希望`fn`的形式为"dir/prog"
# <翻译结束>


<原文开始>
// make child start with a relative program path
<原文结束>

# <翻译开始>
// 让子进程以相对程序路径开始
# <翻译结束>


<原文开始>
// OpenBSD and AIX rely on argv[0]
<原文结束>

# <翻译开始>
// OpenBSD 和 AIX 依赖于 argv[0]
# <翻译结束>


<原文开始>
		// forge argv[0] for child, so that we can verify we could correctly
		// get real path of the executable without influenced by argv[0].
<原文结束>

# <翻译开始>
		// 为子进程伪造argv[0]，以便我们可以验证在不受argv[0]影响的情况下正确获取可执行文件的真正路径。
# <翻译结束>


<原文开始>
// first chdir to another path
<原文结束>

# <翻译开始>
// 首先切换到另一个路径
# <翻译结束>

