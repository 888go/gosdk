
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
// ErrNotFound is the error resulting if a path search failed to find an executable file.
<原文结束>

# <翻译开始>
// ErrNotFound 是当路径搜索未能找到可执行文件时产生的错误。
# <翻译结束>


<原文开始>
// LookPath searches for an executable named file in the
// directories named by the PATH environment variable.
// If file contains a slash, it is tried directly and the PATH is not consulted.
// The result may be an absolute path or a path relative to the current directory.
<原文结束>

# <翻译开始>
// LookPath 在由 PATH 环境变量指定的目录中搜索名为 file 的可执行文件。
// 如果 file 中包含斜杠，会直接尝试执行且不参考 PATH。
// 结果可能是绝对路径，也可能是相对于当前目录的路径。
# <翻译结束>


<原文开始>
// Wasm can not execute processes, so act as if there are no executables at all.
<原文结束>

# <翻译开始>
// Wasm无法执行进程，因此表现为完全没有可执行文件。
# <翻译结束>

