
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
// If the file doesn't exist, create it, or append to the file
<原文结束>

# <翻译开始>
// 如果文件不存在，创建它，或者追加到文件末尾
# <翻译结束>


<原文开始>
// ignore error; Write error takes precedence
<原文结束>

# <翻译开始>
// 忽略错误；优先处理Write错误
# <翻译结束>


<原文开始>
	// Logs can be cleaned out earlier if needed by searching
	// for all directories whose suffix ends in *-logs.
<原文结束>

# <翻译开始>
	// 如果需要，可以通过搜索以 *-logs 结尾的目录来提前清理日志。
# <翻译结束>


<原文开始>
// First, we create a relative symlink to a file.
<原文结束>

# <翻译开始>
// 首先，我们创建一个指向文件的相对符号链接。
# <翻译结束>


<原文开始>
// Allow the example to run on platforms that do not support symbolic links.
<原文结束>

# <翻译开始>
// 允许示例在不支持符号链接的平台上运行。
# <翻译结束>


<原文开始>
// Readlink returns the relative path as passed to os.Symlink.
<原文结束>

# <翻译开始>
// Readlink 返回传递给 os.Symlink 的相对路径。
# <翻译结束>


<原文开始>
// Symlink targets are relative to the directory containing the link.
<原文结束>

# <翻译开始>
// 符号链接的目标是相对于包含该链接的目录的。
# <翻译结束>


<原文开始>
	// Check that the target is correct by comparing it with os.Stat
	// on the original target path.
<原文结束>

# <翻译开始>
	// 通过比较目标与原始目标路径上的os.Stat来检查目标是否正确。
# <翻译结束>


<原文开始>
			// The user has a config file but we couldn't read it.
			// Report the error instead of ignoring their configuration.
<原文结束>

# <翻译开始>
			// 用户有一个配置文件，但我们无法读取它。 			// 相反于忽略他们的配置，我们应该报告这个错误。
# <翻译结束>


<原文开始>
// Use and perhaps make changes to the config.
<原文结束>

# <翻译开始>
// 使用并可能对配置进行更改
# <翻译结束>

