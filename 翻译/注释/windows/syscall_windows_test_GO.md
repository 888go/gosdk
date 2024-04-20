
<原文开始>
// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// 版权所有 ? 2012 The Go Authors。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。
# <翻译结束>


<原文开始>
// it is unlikely to have this character in the filename
<原文结束>

# <翻译开始>
// 在文件名中出现这个字符的可能性很小
# <翻译结束>


<原文开始>
	// Go is not explicitly added to the application compatibility database, so
	// these two functions should return the same thing.
<原文结束>

# <翻译开始>
// Go并未被明确添加至应用程序兼容性数据库中，因此这两个函数应返回相同的结果。
# <翻译结束>


<原文开始>
// Grab a handle to the current process
<原文结束>

# <翻译开始>
// 获取当前进程的句柄
# <翻译结束>


<原文开始>
// Allocate memory to store the result of the query
<原文结束>

# <翻译开始>
// 分配内存以存储查询结果
# <翻译结束>


<原文开始>
// Set the new limits to the current ones
<原文结束>

# <翻译开始>
// 将新限制设置为当前限制
# <翻译结束>


<原文开始>
				// The documentation for CommandLineToArgv takes for granted that
				// the first argument is a valid file path, and doesn't describe any
				// specific behavior for malformed arguments. Empirically it seems to
				// tolerate anything we throw at it, but if we discover cases where it
				// actually returns an error we might need to relax this check.
<原文结束>

# <翻译开始>
// CommandLineToArgv 的文档假设第一个参数为有效的文件路径，并未描述针对格式错误的参数的具体行为。经验表明，它似乎能容忍我们传给它的任何内容，但如果发现其在某些情况下确实返回错误，我们可能需要放宽此检查。
# <翻译结束>


<原文开始>
			// Since DecomposeCommandLine can't handle this string,
			// interpret it as the raw arguments to ComposeCommandLine.
<原文结束>

# <翻译开始>
// 由于DecomposeCommandLine无法处理此字符串，
// 将其解释为ComposeCommandLine的原始参数。
# <翻译结束>


<原文开始>
					// We need to encode the arguments as UTF-16 to pass them to
					// CommandLineToArgvW, so skip inputs that are not valid: they might
					// have one or more runes converted to replacement characters.
<原文结束>

# <翻译开始>
// 我们需要将参数编码为UTF-16，以便传递给CommandLineToArgvW。因此，跳过无效的输入：它们可能有一个或多个字符被转换为替换字符。
# <翻译结束>


<原文开始>
		// It's ok if we compose a different command line than what was read.
		// Just check that we are able to compose something that round-trips
		// to the same results as the original.
<原文结束>

# <翻译开始>
// 即使我们构造的命令行与读取的不同，也是可以接受的。
// 我们只需确保所构造的命令行能够往返处理并得到与原始命令相同的结果。
# <翻译结束>


<原文开始>
					// It is possible that args[0] cannot be encoded exactly, because
					// CommandLineToArgvW doesn't unescape that argument in the same way
					// as the others: since the first argument is assumed to be the name
					// of the program itself, we only have the option of quoted or not.
					//
					// If args[0] contains a space or control character, we must quote it
					// to avoid it being split into multiple arguments.
					// If args[0] already starts with a quote character, we have no way
					// to indicate that that character is part of the literal argument.
					// In either case, if the string already contains a quote character
					// we must avoid misinterpriting that character as the end of the
					// quoted argument string.
					//
					// Unfortunately, ComposeCommandLine does not return an error, so we
					// can't report existing quote characters as errors.
					// Instead, we strip out the problematic quote characters from the
					// argument, and quote the remainder.
					// For paths like C:\"Program Files"\Go\bin\go.exe that is arguably
					// what the caller intended anyway, and for other strings it seems
					// less harmful than corrupting the subsequent arguments.
<原文结束>

# <翻译开始>
// 存在可能args[0]无法被精确编码，因为
// CommandLineToArgvW 对该参数的解码方式与其它参数不同：
// 由于第一个参数假定为程序自身的名称，我们仅能选择是否加引号。
//
// 若 args[0] 包含空格或控制字符，我们必须对其加引号以避免其被解析为多个参数。
// 若 args[0] 已经以引号字符开头，我们无法表明该字符是该实参的一部分。
// 在这两种情况下，若字符串中已经包含引号字符，我们必须避免将其误解释为引号包围的参数字符串的结尾。
//
// 不幸的是，ComposeCommandLine 函数并未返回错误，因此我们无法将已存在的引号字符作为错误报告。
// 作为替代，我们将有问题的引号字符从参数中移除，并对剩余部分加引号。
// 对于类似 C:\"Program Files"\Go\bin\go.exe 的路径，这或许正是调用者所期望的，
// 而对于其他字符串，这样做相比破坏后续参数也显得危害较小。
# <翻译结束>


<原文开始>
// Now that we've verified the legitimate file verifies, let's corrupt it and see if it correctly fails.
<原文结束>

# <翻译开始>
// 现在我们已经验证了合法文件的确能通过验证，接下来让我们篡改该文件，看看它是否能正确地失败。
# <翻译结束>


<原文开始>
// Regression check for go.dev/issue/60223
<原文结束>

# <翻译开始>
// 对 go.dev/issue/60223 的回归检查
# <翻译结束>


<原文开始>
	// Most likely, this should be [0, 4].
	// 0 is the system idle pseudo-process. 4 is the initial system process ID.
	// This test expects that at least one of the PIDs is not 0.
<原文结束>

# <翻译开始>
// 最有可能的情况是，这应为 [0, 4]。
// 其中，0 表示系统空闲伪进程。4 表示初始系统进程 ID。
// 本测试期望至少有一个 PID 不为 0。
# <翻译结束>


<原文开始>
// NB: Assume that we're always the first module. This technically isn't documented anywhere (that I could find), but seems to always hold.
<原文结束>

# <翻译开始>
// 注意：假定我们始终是第一个模块。虽然在任何文档（我发现的范围内）中并未明确指出这一点，但似乎总是成立的。
# <翻译结束>


<原文开始>
// Open test directory with NtCreateFile.
<原文结束>

# <翻译开始>
// 使用NtCreateFile打开测试目录
# <翻译结束>


<原文开始>
// Create a file in test directory with NtCreateFile.
<原文结束>

# <翻译开始>
// 通过NtCreateFile在测试目录中创建一个文件
# <翻译结束>


<原文开始>
// Rename file with NtSetInformationFile.
<原文结束>

# <翻译开始>
// 使用NtSetInformationFile重命名文件
# <翻译结束>


<原文开始>
// We allocate handles in a closure to provoke a UaF in the case of attributeList.Update being buggy.
<原文结束>

# <翻译开始>
// 我们在一个闭包中分配句柄，以便在 attributeList.Update 函数存在漏洞的情况下触发 UaF（使用后释放）问题。
# <翻译结束>


<原文开始>
// Go 1.16's pipe handles aren't inheritable, so mark it explicitly as such here.
<原文结束>

# <翻译开始>
// Go 1.16 版本的管道句柄不可继承，因此在这里明确将其标记为不可继承。
# <翻译结束>


<原文开始>
// no more data available - break the loop
<原文结束>

# <翻译开始>
// 无更多可用数据 - 终止循环
# <翻译结束>


<原文开始>
// buffer is too small - reallocate and try again
<原文结束>

# <翻译开始>
// 缓冲区太小 —— 重新分配并尝试再次
# <翻译结束>


<原文开始>
// found a record - display the item and fetch next item
<原文结束>

# <翻译开始>
// 找到一条记录 —— 显示该条目并获取下一条
# <翻译结束>


<原文开始>
// see https://go.dev/issue/31316
<原文结束>

# <翻译开始>
// 参见 https://go.dev/issue/31316
# <翻译结束>

