
<原文开始>
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// 版权所有 ? 2009 Go 作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 此许可证可在 LICENSE 文件中找到。
# <翻译结束>


<原文开始>
// EscapeArg rewrites command line argument s as prescribed
// in http://msdn.microsoft.com/en-us/library/ms880421.
// This function returns "" (2 double quotes) if s is empty.
// Alternatively, these transformations are done:
//   - every back slash (\) is doubled, but only if immediately
//     followed by double quote (");
//   - every double quote (") is escaped by back slash (\);
//   - finally, s is wrapped with double quotes (arg -> "arg"),
//     but only if there is space or tab inside s.
<原文结束>

# <翻译开始>
// EscapeArg 按照 http://msdn.microsoft.com/en-us/library/ms880421 中所规定的方式重写命令行参数 s。
// 如果 s 为空，此函数返回 ""（两个双引号）。
// 另外，进行如下转换：
//   - 若反斜杠 (\) 立即跟随双引号 (")，则将每个反斜杠加倍；
//   - 用反斜杠对每个双引号 (") 进行转义；
//   - 最后，仅当 s 内包含空格或制表符时，用双引号包裹 s（将 arg 转换为 "arg"）。
# <翻译结束>


<原文开始>
// Reserve space for quotes.
<原文结束>

# <翻译开始>
// 为引号预留空间
# <翻译结束>


<原文开始>
// ComposeCommandLine escapes and joins the given arguments suitable for use as a Windows command line,
// in CreateProcess's CommandLine argument, CreateService/ChangeServiceConfig's BinaryPathName argument,
// or any program that uses CommandLineToArgv.
<原文结束>

# <翻译开始>
// ComposeCommandLine 对给定的参数进行转义并拼接，生成适用于作为Windows命令行的结果，
// 可用于CreateProcess函数的CommandLine参数、CreateService/ChangeServiceConfig函数的BinaryPathName参数，
// 或任何使用CommandLineToArgv函数的程序。
# <翻译结束>


<原文开始>
			// Force quotes for not only the ASCII space and tab as described in the
			// MSDN article, but also ASCII control characters.
			// The documentation for CommandLineToArgvW doesn't say what happens when
			// the first argument is not a valid program name, but it empirically
			// seems to drop unquoted control characters.
<原文结束>

# <翻译开始>
			// 强制对以下字符使用引号：不仅包括MSDN文章中描述的ASCII空格和制表符，还包括ASCII控制字符。
			// CommandLineToArgvW函数的文档并未说明当第一个参数不是有效的程序名时会发生什么情况，但经实测，
			// 该函数似乎会舍弃未加引号的控制字符。
# <翻译结束>


<原文开始>
				// This quote would interfere with our surrounding quotes.
				// We have no way to report an error, so just strip out
				// the offending character instead.
<原文结束>

# <翻译开始>
				// 此引号会与我们周围的引号冲突。
				// 我们无法报告错误，因此只需移除
				// 有问题的字符即可。
# <翻译结束>


<原文开始>
			// args[0] is a valid command line representing itself.
			// No need to allocate a new slice or string for it.
<原文结束>

# <翻译开始>
			// args[0] 是一个表示其自身的有效命令行参数。
			// 无需为其分配新的切片或字符串。
# <翻译结束>


<原文开始>
		// TODO(bcmills): since we're already appending to a slice, it would be nice
		// to avoid the intermediate allocations of EscapeArg.
		// Perhaps we can factor out an appendEscapedArg function.
<原文结束>

# <翻译开始>
		// TODO(bcmills): 既然我们已在向切片追加内容，那么能避免 `EscapeArg` 的中间分配操作就很好了。
		// 或许我们可以提取出一个 `appendEscapedArg` 函数。
# <翻译结束>


<原文开始>
// DecomposeCommandLine breaks apart its argument command line into unescaped parts using CommandLineToArgv,
// as gathered from GetCommandLine, QUERY_SERVICE_CONFIG's BinaryPathName argument, or elsewhere that
// command lines are passed around.
// DecomposeCommandLine returns an error if commandLine contains NUL.
<原文结束>

# <翻译开始>
// DecomposeCommandLine 将其参数 command line 使用 CommandLineToArgv 拆分为未转义的部分，这些部分可从 GetCommandLine、QUERY_SERVICE_CONFIG 的 BinaryPathName 参数或其他传递命令行的途径获取。
// 若 commandLine 包含 NUL 字符，DecomposeCommandLine 将返回错误。
# <翻译结束>


<原文开始>
// CommandLineToArgv parses a Unicode command line string and sets
// argc to the number of parsed arguments.
//
// The returned memory should be freed using a single call to LocalFree.
//
// Note that although the return type of CommandLineToArgv indicates 8192
// entries of up to 8192 characters each, the actual count of parsed arguments
// may exceed 8192, and the documentation for CommandLineToArgvW does not mention
// any bound on the lengths of the individual argument strings.
// (See https://go.dev/issue/63236.)
<原文结束>

# <翻译开始>
// CommandLineToArgv 函数解析一个 Unicode 格式的命令行字符串，并将解析得到的参数个数设置到 argc 中。
//
// 返回的内存应当使用一次 LocalFree 调用来释放。
//
// 需要注意的是，尽管 CommandLineToArgv 的返回类型表明其可以容纳 8192 个长度不超过 8192 字符的条目，
// 实际解析得到的参数数量可能超过 8192，且 CommandLineToArgvW 文档并未提及单个参数字符串长度有任何限制。
// （参见 https://go.dev/issue/63236。）
# <翻译结束>


<原文开始>
// FullPath retrieves the full path of the specified file.
<原文结束>

# <翻译开始>
// FullPath 获取指定文件的完整路径
# <翻译结束>


<原文开始>
// NewProcThreadAttributeList allocates a new ProcThreadAttributeListContainer, with the requested maximum number of attributes.
<原文结束>

# <翻译开始>
// NewProcThreadAttributeList 分配一个具有所请求最大属性数量的新 ProcThreadAttributeListContainer。
# <翻译结束>


<原文开始>
// Update modifies the ProcThreadAttributeList using UpdateProcThreadAttribute.
<原文结束>

# <翻译开始>
// Update 通过调用 UpdateProcThreadAttribute 函数来修改 ProcThreadAttributeList。
# <翻译结束>


<原文开始>
// Delete frees ProcThreadAttributeList's resources.
<原文结束>

# <翻译开始>
// Delete 释放 ProcThreadAttributeList 的资源。
# <翻译结束>


<原文开始>
// List returns the actual ProcThreadAttributeList to be passed to StartupInfoEx.
<原文结束>

# <翻译开始>
// List 返回将传递给 StartupInfoEx 的实际 ProcThreadAttributeList。
# <翻译结束>

