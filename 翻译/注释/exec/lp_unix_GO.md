
<原文开始>
// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// 版权所有 ? 2010 Go 作者。保留所有权利。
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
	// ENOSYS means Eaccess is not available or not implemented.
	// EPERM can be returned by Linux containers employing seccomp.
	// In both cases, fall back to checking the permission bits.
<原文结束>

# <翻译开始>
// ENOSYS 表示 Eaccess 功能不可用或尚未实现。
// 在 Linux 容器使用 seccomp 的情况下，可能返回 EPERM 错误。
// 在这两种情况下，我们都回退到检查权限位。
# <翻译结束>


<原文开始>
// LookPath searches for an executable named file in the
// directories named by the PATH environment variable.
// If file contains a slash, it is tried directly and the PATH is not consulted.
// Otherwise, on success, the result is an absolute path.
//
// In older versions of Go, LookPath could return a path relative to the current directory.
// As of Go 1.19, LookPath will instead return that path along with an error satisfying
// errors.Is(err, ErrDot). See the package documentation for more details.
<原文结束>

# <翻译开始>
// ```markdown
// LookPath在PATH环境变量指定的目录中搜索名为file的可执行文件。
// 如果file包含斜杠，会直接尝试该路径而不查询PATH。
// 否则，在成功时，结果为一个绝对路径。
// 
// 在Go的旧版本中，LookPath可能返回一个相对于当前目录的路径。
// 从Go 1.19版本开始，LookPath将在返回该路径的同时返回一个错误，该错误满足errors.Is(err, ErrDot)条件。有关更多详细信息，请参阅包文档。
// ```
# <翻译结束>


<原文开始>
	// NOTE(rsc): I wish we could use the Plan 9 behavior here
	// (only bypass the path if file begins with / or ./ or ../)
	// but that would not match all the Unix shells.
<原文结束>

# <翻译开始>
// 注释（rsc）：我希望我们能在这里使用Plan 9的行为
// （仅当文件以 /、./ 或 ../ 开头时绕过路径）
// 但那样将无法匹配所有的Unix shell。
# <翻译结束>


<原文开始>
// Unix shell semantics: path element "" means "."
<原文结束>

# <翻译开始>
// Unix Shell 语义：路径元素 "" 表示 "."
# <翻译结束>

