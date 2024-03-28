
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
// LookPath searches for an executable named file in the
// directories named by the PATH environment variable.
// LookPath also uses PATHEXT environment variable to match
// a suitable candidate.
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
// LookPath还会使用PATHEXT环境变量来匹配合适的候选文件。
// 如果file包含斜杠（/），则会直接尝试该路径，而不查询PATH。
// 成功时，返回的结果是绝对路径。
//
// 在Go的旧版本中，LookPath可能返回一个相对于当前目录的路径。
// 从Go 1.19开始，LookPath将改为返回该路径以及一个满足errors.Is(err, ErrDot)的错误。有关更多详细信息，请参阅包文档。
// ```
# <翻译结束>


<原文开始>
	// On Windows, creating the NoDefaultCurrentDirectoryInExePath
	// environment variable (with any value or no value!) signals that
	// path lookups should skip the current directory.
	// In theory we are supposed to call NeedCurrentDirectoryForExePathW
	// "as the registry location of this environment variable can change"
	// but that seems exceedingly unlikely: it would break all users who
	// have configured their environment this way!
	// https://docs.microsoft.com/en-us/windows/win32/api/processenv/nf-processenv-needcurrentdirectoryforexepathw
	// See also go.dev/issue/43947.
<原文结束>

# <翻译开始>
// 在Windows系统中，创建名为NoDefaultCurrentDirectoryInExePath的环境变量（无论其值是什么或没有值）表示在查找路径时应跳过当前目录。理论上，我们应该调用NeedCurrentDirectoryForExePathW函数，因为“这个环境变量的注册表位置可能会改变”，但这种情况似乎极不可能发生：它会破坏所有以这种方式配置环境的用户！参考文档：https://docs.microsoft.com/en-us/windows/win32/api/processenv/nf-processenv-needcurrentdirectoryforexepathw 。同时，请参阅go.dev/issue/43947。
# <翻译结束>


<原文开始>
				// https://go.dev/issue/53536: if we resolved a relative path implicitly,
				// and it is the same executable that would be resolved from the explicit %PATH%,
				// prefer the explicit name for the executable (and, likely, no error) instead
				// of the equivalent implicit name with ErrDot.
				//
				// Otherwise, return the ErrDot for the implicit path as soon as we find
				// out that the explicit one doesn't match.
<原文结束>

# <翻译开始>
// 参考 go.dev/issue/53536：如果我们隐式解析了相对路径，并且该路径指向的可执行文件与通过显式%PATH%解析得到的是同一个，
// 那么我们应优先选择显式的可执行文件名（并且可能不会出现错误），而不是等效的带有 ErrDot 的隐式文件名。
//
// 否则，一旦发现显式路径不匹配，就立即返回隐式路径的 ErrDot 错误。
# <翻译结束>

