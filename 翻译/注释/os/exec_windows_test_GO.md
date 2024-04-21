
<原文开始>
// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// 版权所有 2023 The Go Authors。保留所有权利。
// 使用本源代码须遵循 BSD 风格的许可证，
// 该许可证可在 LICENSE 文件中找到。
# <翻译结束>


<原文开始>
// Regression test for golang.org/issue/25965.
<原文结束>

# <翻译开始>
// 为golang.org/issue/25965问题的回归测试。
# <翻译结束>


<原文开始>
// First create n executables.
<原文结束>

# <翻译开始>
// 首先创建n个可执行文件。
# <翻译结束>


<原文开始>
	// Then run each executable and remove its directory.
	// Run each executable in a separate goroutine to add some load
	// and increase the chance of triggering the bug.
<原文结束>

# <翻译开始>
// 然后运行每个可执行文件并移除其所在的目录。
// 为增加负载并提高触发问题的可能性，以单独的goroutine运行每个可执行文件。
# <翻译结束>


<原文开始>
// Run test.exe without executing any test, just to make it do something.
<原文结束>

# <翻译开始>
// 运行test.exe而不执行任何测试，只是为了使其执行一些操作。
# <翻译结束>


<原文开始>
// Remove dir and check that it doesn't return `ERROR_ACCESS_DENIED`.
<原文结束>

# <翻译开始>
// 删除目录，并检查它不返回“ERROR_ACCESS_DENIED”错误。
# <翻译结束>

