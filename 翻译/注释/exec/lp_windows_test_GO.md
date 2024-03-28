
<原文开始>
// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// 版权所有 ? 2013 Go 作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。
# <翻译结束>


<原文开始>
// Use an external test to avoid os/exec -> internal/testenv -> os/exec
// circular dependency.
<原文结束>

# <翻译开始>
// 使用外部测试以避免 os/exec -> internal/testenv -> os/exec 之间的循环依赖。
# <翻译结束>


<原文开始>
// test is expected to fail
<原文结束>

# <翻译开始>
// 预期测试将失败
# <翻译结束>


<原文开始>
// normalise program output
<原文结束>

# <翻译开始>
// 标准化程序输出
# <翻译结束>


<原文开始>
// trim terminating \r and \n that batch file outputs
<原文结束>

# <翻译开始>
// 去除批处理文件输出末尾的\r和\n
# <翻译结束>


<原文开始>
// Add dir in front of every directory in the PATH.
<原文结束>

# <翻译开始>
// 在PATH中每个目录的前面添加dir。
# <翻译结束>


<原文开始>
// createFiles copies srcPath file into multiply files.
// It uses dir as prefix for all destination files.
<原文结束>

# <翻译开始>
// createFiles 将srcPath文件复制到多个文件中。
// 它使用dir作为所有目标文件的前缀。
# <翻译结束>


<原文开始>
	// Run "cmd.exe /c test.searchFor" with new environment and
	// work directory set. All candidates are copies of printpath.exe.
	// These will output their program paths when run.
<原文结束>

# <翻译开始>
// 使用新的环境和工作目录运行 "cmd.exe /c test.searchFor"。所有候选项都是 printpath.exe 的副本。
// 当运行时，这些将会输出它们的程序路径。
# <翻译结束>


<原文开始>
// Run the lookpath program with new environment and work directory set.
<原文结束>

# <翻译开始>
// 在设置新的环境和工作目录下运行lookpath程序。
# <翻译结束>


<原文开始>
// both failed -> continue
<原文结束>

# <翻译开始>
// 两者都失败 -> 继续
# <翻译结束>


<原文开始>
// TODO(brainman): do not know why this fails
<原文结束>

# <翻译开始>
// TODO(脑曼): 不知道为什么这个会失败
# <翻译结束>


<原文开始>
	// If the command name specifies a path, the shell searches
	// the specified path for an executable file matching
	// the command name. If a match is found, the external
	// command (the executable file) executes.
<原文结束>

# <翻译开始>
// 如果命令名称指定了路径，shell将会在指定的路径中搜索与命令名称相匹配的可执行文件。
// 如果找到匹配项，则执行外部命令（即该可执行文件）。
# <翻译结束>


<原文开始>
	// If the command name specifies a path, the shell searches
	// the specified path for an executable file matching the command
	// name. ... If no match is found, the shell reports an error
	// and command processing completes.
<原文结束>

# <翻译开始>
// 如果命令名指定了路径，shell 会在指定的路径下搜索与命令名匹配的可执行文件。...如果没有找到匹配项，shell 会报告错误，并完成命令处理。
# <翻译结束>


<原文开始>
	// If the command name does not specify a path, the shell
	// searches the current directory for an executable file
	// matching the command name. If a match is found, the external
	// command (the executable file) executes.
<原文结束>

# <翻译开始>
// 如果命令名称未指定路径，shell将在当前目录下搜索与命令名称匹配的可执行文件。如果找到匹配项，则执行外部命令（即该可执行文件）。
# <翻译结束>


<原文开始>
	// The shell now searches each directory specified by the
	// PATH environment variable, in the order listed, for an
	// executable file matching the command name. If a match
	// is found, the external command (the executable file) executes.
<原文结束>

# <翻译开始>
// 现在，shell 按照列出的顺序搜索由 `PATH` 环境变量指定的每个目录，寻找与命令名匹配的可执行文件。如果找到匹配项，则执行外部命令（即该可执行文件）。
# <翻译结束>


<原文开始>
	// The shell now searches each directory specified by the
	// PATH environment variable, in the order listed, for an
	// executable file matching the command name. If no match
	// is found, the shell reports an error and command processing
	// completes.
<原文结束>

# <翻译开始>
// 现在，shell 按照列出的顺序搜索 PATH 环境变量中指定的每个目录，寻找与命令名匹配的可执行文件。如果没有找到匹配项，shell 将报告错误，并完成命令处理。
# <翻译结束>


<原文开始>
	// If the command name includes a file extension, the shell
	// searches each directory for the exact file name specified
	// by the command name.
<原文结束>

# <翻译开始>
// 如果命令名称中包含文件扩展名，shell将在每个目录下搜索与命令名称指定的精确文件名。
# <翻译结束>


<原文开始>
// includes extension and not exact file name match
<原文结束>

# <翻译开始>
// 包含扩展名且并非完全匹配文件名
# <翻译结束>


<原文开始>
	// If the command name does not include a file extension, the shell
	// adds the extensions listed in the PATHEXT environment variable,
	// one by one, and searches the directory for that file name. Note
	// that the shell tries all possible file extensions in a specific
	// directory before moving on to search the next directory
	// (if there is one).
<原文结束>

# <翻译开始>
// 如果命令名称不包含文件扩展名，shell 将会依次添加 PATHEXT 环境变量中列出的扩展名，并在该目录下搜索带有该扩展名的文件名。请注意，在尝试下一个目录（如果存在）之前，shell 会在特定目录下尝试所有可能的文件扩展名。
# <翻译结束>


<原文开始>
// tried all extensions in PATHEXT, but none matches
<原文结束>

# <翻译开始>
// 尝试了PATHEXT中的所有扩展名，但没有一个匹配成功
# <翻译结束>


<原文开始>
// testing commands with no slash, like `a.exe`
<原文结束>

# <翻译开始>
// 测试不带斜杠的命令，如 `a.exe`
# <翻译结束>


<原文开始>
// should find a.exe in current directory
<原文结束>

# <翻译开始>
// 应当在当前目录下找到 a.exe
# <翻译结束>


<原文开始>
// like above, but add PATH in attempt to break the test
<原文结束>

# <翻译开始>
// 类似上面的做法，但添加PATH以尝试打破测试
# <翻译结束>


<原文开始>
// like above, but use "a" instead of "a.exe" for command
<原文结束>

# <翻译开始>
// 与上面类似，但使用 "a" 而不是 "a.exe" 作为命令
# <翻译结束>


<原文开始>
// testing commands with slash, like `.\a.exe`
<原文结束>

# <翻译开始>
// 测试使用反斜杠的命令，如 `.\a.exe`
# <翻译结束>


<原文开始>
// like above, but adding `.` in front of executable should still be OK
<原文结束>

# <翻译开始>
// 如上所述，但在可执行文件前添加`.`也应该没问题
# <翻译结束>


<原文开始>
// like above, but with PATH added in attempt to break it
<原文结束>

# <翻译开始>
// 如上所示，但尝试通过添加 PATH 来破坏它
# <翻译结束>


<原文开始>
// like above, but make sure .exe is tried even for commands with slash
<原文结束>

# <翻译开始>
// 与上面类似，但确保即使命令带有斜杠也会尝试执行.exe
# <翻译结束>


<原文开始>
// tests commands, like `a.exe`, with c.Dir set
<原文结束>

# <翻译开始>
// 测试带有 c.Dir 设置的命令，如 `a.exe`
# <翻译结束>


<原文开始>
// should not find a.exe in p, because LookPath(`a.exe`) will fail
<原文结束>

# <翻译开始>
// 应该在 p 中找不到 a.exe，因为 LookPath(`a.exe`) 会失败
# <翻译结束>


<原文开始>
		// LookPath(`a.exe`) will find `.\a.exe`, but prefixing that with
		// dir `p\a.exe` will refer to a non-existent file
<原文结束>

# <翻译开始>
// LookPath("a.exe") 会找到 `.\a.exe`，但如果在该路径前添加目录前缀 `p\a.exe`，
// 则会指向一个不存在的文件。
# <翻译结束>


<原文开始>
		// like above, but making test succeed by installing file
		// in referred destination (so LookPath(`a.exe`) will still
		// find `.\a.exe`, but we successfully execute `p\a.exe`)
<原文结束>

# <翻译开始>
// 类似于上面的做法，但通过在引用的目标位置安装文件（这样LookPath(`a.exe`)仍然会找到`.\a.exe`，但我们成功执行的是`p\a.exe`），从而使得测试成功
# <翻译结束>


<原文开始>
		// finds `a.exe` in the PATH regardless of dir set
		// because LookPath returns full path in that case
<原文结束>

# <翻译开始>
// 此函数会无视当前目录设置，在PATH中查找并找到`a.exe`
// 因为在这种情况下，LookPath返回的是完整路径
# <翻译结束>


<原文开始>
// tests commands, like `.\a.exe`, with c.Dir set
<原文结束>

# <翻译开始>
// 测试命令，如 `.\a.exe`，其中设置了 c.Dir
# <翻译结束>


<原文开始>
// should use dir when command is path, like ".\a.exe"
<原文结束>

# <翻译开始>
// 当命令是路径时，应使用dir，例如“.\a.exe”
# <翻译结束>


<原文开始>
// buildPrintPathExe creates a Go program that prints its own path.
// dir is a temp directory where executable will be created.
// The function returns full path to the created program.
<原文结束>

# <翻译开始>
// buildPrintPathExe 创建一个Go程序，该程序会打印自身的路径。
// dir 是一个临时目录，用于在其中创建可执行文件。
// 该函数返回所创建程序的完整路径。
# <翻译结束>

