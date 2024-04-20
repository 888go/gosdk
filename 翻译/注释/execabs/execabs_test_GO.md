
<原文开始>
// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// 版权所有 2020 Go作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。
# <翻译结束>


<原文开始>
// hasExec reports whether the current system can start new processes
// using os.StartProcess or (more commonly) exec.Command.
// Copied from internal/testenv.HasExec
<原文结束>

# <翻译开始>
// hasExec 报告当前系统是否能够使用 os.StartProcess 或（更常见地）exec.Command 来启动新进程。
// 从 internal/testenv.HasExec 复制而来
# <翻译结束>


<原文开始>
// mustHaveExec checks that the current system can start new processes
// using os.StartProcess or (more commonly) exec.Command.
// If not, mustHaveExec calls t.Skip with an explanation.
// Copied from internal/testenv.MustHaveExec
<原文结束>

# <翻译开始>
// mustHaveExec 检查当前系统是否能够使用 os.StartProcess 或（更常见地）exec.Command 启动新进程。
// 如果不能，mustHaveExec 会调用 t.Skip 并附带解释原因。
// 该代码从 internal/testenv.MustHaveExec 复制而来
# <翻译结束>


<原文开始>
			// add "." to PATH so that exec.LookPath looks in the current directory on
			// non-windows platforms as well
<原文结束>

# <翻译开始>
// 在非Windows平台下，向PATH中添加"."，以便exec.LookPath也能在当前目录中查找
# <翻译结束>

