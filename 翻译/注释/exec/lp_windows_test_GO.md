
<原文开始>
// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// 版权所有 2013 The Go Authors。保留所有权利。
// 使用本源代码受 BSD 风格许可协议约束，
// 该协议可在 LICENSE 文件中找到。
# <翻译结束>


<原文开始>
// Use an external test to avoid os/exec -> internal/testenv -> os/exec
// circular dependency.
<原文结束>

# <翻译开始>
// 使用外部测试以避免 os/exec -> internal/testenv -> os/exec 之间的循环依赖关系
# <翻译结束>


<原文开始>
// makePATH returns a PATH variable referring to the
// given directories relative to a root directory.
//
// The empty string results in an empty entry.
// Paths beginning with . are kept as relative entries.
<原文结束>

# <翻译开始>
// makePATH 函数返回一个引用给定目录（相对于根目录）的 PATH 变量。
//
// 空字符串将导致生成一个空条目。
// 以 . 开头的路径将作为相对条目保留。
# <翻译结束>


<原文开始>
// installProgs creates executable files (or symlinks to executable files) at
// multiple destination paths. It uses root as prefix for all destination files.
<原文结束>

# <翻译开始>
// installProgs 在多个目标路径下创建可执行文件（或指向可执行文件的符号链接）。它使用 root 作为所有目标文件的前缀。
# <翻译结束>


<原文开始>
// directory and PATH entry only.
<原文结束>

# <翻译开始>
// 仅目录和 PATH 条目
# <翻译结束>


<原文开始>
// installExe installs a copy of the test executable
// at the given location, creating directories as needed.
//
// (We use a copy instead of just a symlink to ensure that os.Executable
// always reports an unambiguous path, regardless of how it is implemented.)
<原文结束>

# <翻译开始>
// installExe 在给定位置安装测试可执行程序的一个副本，必要时创建目录。
//
// （我们使用副本而非仅使用符号链接，以确保无论其具体实现如何，os.Executable 始终报告一个明确的路径。）
# <翻译结束>


<原文开始>
// installBat creates a batch file at dst that prints its own
// path when run.
<原文结束>

# <翻译开始>
// installBat 在dst位置创建一个批处理文件，当运行时会打印其自身路径
# <翻译结束>


<原文开始>
// if nil, use all parent directories from files
<原文结束>

# <翻译开始>
// 若为nil，则使用文件的所有上级目录
# <翻译结束>


<原文开始>
// if true, do not check want against the behavior of cmd.exe
<原文结束>

# <翻译开始>
// 如果为真，不检查want与cmd.exe行为的一致性
# <翻译结束>


<原文开始>
		// If cmd.exe is too old it might not respect NoDefaultCurrentDirectoryInExePath,
		// so skip that check.
<原文结束>

# <翻译开始>
		// 如果cmd.exe版本过旧，可能不支持NoDefaultCurrentDirectoryInExePath选项，
		// 因此跳过该检查。
# <翻译结束>


<原文开始>
// Not parallel: uses Chdir and Setenv.
<原文结束>

# <翻译开始>
// 非并行：使用了 Chdir 和 Setenv。
# <翻译结束>


<原文开始>
	// We are using the "printpath" command mode to test exec.Command here,
	// so we won't be calling helperCommand to resolve it.
	// That may cause it to appear to be unused.
<原文结束>

# <翻译开始>
	// 我们在这里使用“printpath”命令模式来测试exec.Command，
	// 因此我们不会调用helperCommand来解析它。
	// 这可能导致它看似未被使用。
# <翻译结束>


<原文开始>
	// Before we begin, find the absolute path to cmd.exe.
	// In non-short mode, we will use it to check the ground truth
	// of the test's "want" field.
<原文结束>

# <翻译开始>
	// 在开始之前，找出cmd.exe的绝对路径。
	// 在非短模式下，我们将用它来校验测试“want”字段的实际情况。
# <翻译结束>


<原文开始>
				// Check that cmd.exe, which is our source of ground truth,
				// agrees that our test case is correct.
<原文结束>

# <翻译开始>
				// 检查 cmd.exe（作为我们的基准真相来源）是否认同我们的测试用例是正确的。
# <翻译结束>


<原文开始>
// cmd.exe disagrees. Probably the test case is wrong?
<原文结束>

# <翻译开始>
// cmd.exe 不同意。可能是测试用例有误？
# <翻译结束>


<原文开始>
// the resolved c.Path, if different from want
<原文结束>

# <翻译开始>
// 解析得到的c.Path，如果与want不同
# <翻译结束>


<原文开始>
// testing commands with no slash, like `a.exe`
<原文结束>

# <翻译开始>
// 测试不带斜杠的命令，如 `a.exe`
# <翻译结束>


<原文开始>
// testing commands with slash, like `.\a.exe`
<原文结束>

# <翻译开始>
// 测试使用反斜杠的命令，如 `.\a.exe`
# <翻译结束>


<原文开始>
// tests commands, like `a.exe`, with c.Dir set
<原文结束>

# <翻译开始>
// 测试命令，如`a.exe`，其中已设置c.Dir
# <翻译结束>


<原文开始>
		// should not find a.exe in p, because LookPath(`a.exe`) will fail when
		// called by Command (before Dir is set), and that error is sticky.
<原文结束>

# <翻译开始>
		// 不应在p中找到a.exe，因为当Command调用LookPath(`a.exe`)时（在设置Dir之前），该操作将失败，且该错误是持久的。
# <翻译结束>


<原文开始>
		// LookPath(`a.exe`) will resolve to `.\a.exe`, but prefixing that with
		// dir `p\a.exe` will refer to a non-existent file
<原文结束>

# <翻译开始>
		// LookPath(`a.exe`) 会解析为 `.\a.exe`，但将该路径前缀为目录 `p` 即 `p\a.exe` 时，将指向一个不存在的文件
# <翻译结束>


<原文开始>
		// like above, but making test succeed by installing file
		// in referred destination (so LookPath(`a.exe`) will still
		// find `.\a.exe`, but we successfully execute `p\a.exe`)
<原文结束>

# <翻译开始>
		// 类似于上面的做法，但通过将文件安装到所引用的目标位置（这样LookPath(`a.exe`)仍会找到`.\a.exe`，但我们成功执行了`p\a.exe`），使测试成功
# <翻译结束>


<原文开始>
// like above, but add PATH in attempt to break the test
<原文结束>

# <翻译开始>
// 与上面类似，但添加PATH以尝试打破测试
# <翻译结束>


<原文开始>
// like above, but use "a" instead of "a.exe" for command
<原文结束>

# <翻译开始>
// 与上述相同，但使用 "a" 代替 "a.exe" 作为命令
# <翻译结束>


<原文开始>
		// finds `a.exe` in the PATH regardless of Dir because Command resolves the
		// full path (using LookPath) before Dir is set.
<原文结束>

# <翻译开始>
		// 查找 `a.exe`，忽略 Dir，因为 Command 在设置 Dir 之前通过 LookPath 解析完整路径。
# <翻译结束>


<原文开始>
// tests commands, like `.\a.exe`, with c.Dir set
<原文结束>

# <翻译开始>
// 测试命令，如 `.\a.exe`，其中已设置 c.Dir
# <翻译结束>


<原文开始>
// should use dir when command is path, like ".\a.exe"
<原文结束>

# <翻译开始>
// 当命令为路径形式时，如 ".\a.exe"，应使用 dir
# <翻译结束>


<原文开始>
// like above, but with PATH added in attempt to break it
<原文结束>

# <翻译开始>
// 与上面类似，但尝试添加 PATH 以使其失效
# <翻译结束>


<原文开始>
// LookPath(".\a") will fail before Dir is set, and that error is sticky.
<原文结束>

# <翻译开始>
// 在设置 Dir 之前，LookPath(".\a") 将会失败，且该错误是持久的。
# <翻译结束>


<原文开始>
	// We expect that ".com" is always included in PATHEXT, but it may also be
	// found in the import path of a Go package. If it is at the root of the
	// import path, the resulting executable may be named like "example.com.exe".
	//
	// Since "example.com" looks like a proper executable name, it is probably ok
	// for exec.Command to try to run it directly without re-resolving it.
	// However, exec.LookPath should try a little harder to figure it out.
<原文结束>

# <翻译开始>
	// 我们期望 ".com" 总是包含在 PATHEXT 中，但它也可能出现在 Go 包的导入路径中。如果它位于导入路径的根部，则生成的可执行文件可能被命名为 "example.com.exe"。
	// 
	// 由于 "example.com" 看起来像一个合理的可执行文件名，因此 exec.Command 尝试直接运行它而不重新解析可能是可以接受的。
	// 但是，exec.LookPath 应该更努力地尝试确定它。
# <翻译结束>

