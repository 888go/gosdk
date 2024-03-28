
<原文开始>
// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// 版权所有 ? 2011 Go 作者。保留所有权利。
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
// directories named by the path environment variable.
// If file begins with "/", "#", "./", or "../", it is tried
// directly and the path is not consulted.
// On success, the result is an absolute path.
//
// In older versions of Go, LookPath could return a path relative to the current directory.
// As of Go 1.19, LookPath will instead return that path along with an error satisfying
// errors.Is(err, ErrDot). See the package documentation for more details.
<原文结束>

# <翻译开始>
// LookPath 在由环境变量 `path` 指定的目录中搜索名为 file 的可执行文件。
// 如果 file 以 "/"、"#"、"./" 或 "../" 开头，则会直接尝试执行，而不查询路径。
// 成功时，返回的结果是绝对路径。
//
// 在 Go 的旧版本中，LookPath 可能会返回相对于当前目录的路径。
// 从 Go 1.19 版本开始，LookPath 将改为返回该路径以及一个满足 errors.Is(err, ErrDot) 的错误。有关更多详细信息，请参阅包文档。
# <翻译结束>


<原文开始>
// skip the path lookup for these prefixes
<原文结束>

# <翻译开始>
// 对于这些前缀，跳过路径查找
# <翻译结束>

