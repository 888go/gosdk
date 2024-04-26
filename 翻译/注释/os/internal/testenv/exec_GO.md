
<原文开始>
// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// 版权所有 2015 The Go 作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。
# <翻译结束>


<原文开始>
// MustHaveExec checks that the current system can start new processes
// using os.StartProcess or (more commonly) exec.Command.
// If not, MustHaveExec calls t.Skip with an explanation.
//
// On some platforms MustHaveExec checks for exec support by re-executing the
// current executable, which must be a binary built by 'go test'.
// We intentionally do not provide a HasExec function because of the risk of
// inappropriate recursion in TestMain functions.
//
// To check for exec support outside of a test, just try to exec the command.
// If exec is not supported, testenv.SyscallIsNotSupported will return true
// for the resulting error.
<原文结束>

# <翻译开始>
// MustHaveExec 检查当前系统是否能够使用 os.StartProcess 或（更常见地）exec.Command 启动新进程。
// 若不能，MustHaveExec 会调用 t.Skip 并附带解释原因。
//
// 在某些平台上，MustHaveExec 通过重新执行当前可执行文件来检查 exec 支持情况，
// 此可执行文件必须是由 'go test' 构建的二进制文件。我们有意不提供 HasExec 函数，
// 因为在 TestMain 函数中存在不合适递归的风险。
//
// 要在测试之外检查 exec 支持情况，只需尝试执行命令即可。若不支持 exec，
// 对于由此产生的错误，testenv.SyscallIsNotSupported 将返回 true。
# <翻译结束>


<原文开始>
// Assume that exec always works on non-mobile platforms and Android.
<原文结束>

# <翻译开始>
// 假设exec总是在非移动平台和Android上正常工作。
# <翻译结束>


<原文开始>
	// ios has an exec syscall but on real iOS devices it might return a
	// permission error. In an emulated environment (such as a Corellium host)
	// it might succeed, so if we need to exec we'll just have to try it and
	// find out.
	//
	// As of 2023-04-19 wasip1 and js don't have exec syscalls at all, but we
	// may as well use the same path so that this branch can be tested without
	// an ios environment.
<原文结束>

# <翻译开始>
// 在iOS中有一个exec系统调用，但在实际的iOS设备上可能会返回权限错误。在模拟环境中（如Corellium主机），它可能会成功。因此，如果我们需要执行，就只能尝试并找出结果。
// 
// 截至2023-04-19，wasip1和js根本没有exec系统调用，但我们最好使用相同的路径，这样即使没有iOS环境，这个分支也可以进行测试。
# <翻译结束>


<原文开始>
		// This isn't a standard 'go test' binary, so we don't know how to
		// self-exec in a way that should succeed without side effects.
		// Just forget it.
<原文结束>

# <翻译开始>
// 这不是一个标准的 "go test" 可执行文件，所以我们不知道如何以不会产生副作用的方式自我执行。就忽略它吧。
# <翻译结束>


<原文开始>
	// We know that this is a test executable. We should be able to run it with a
	// no-op flag to check for overall exec support.
<原文结束>

# <翻译开始>
// 我们知道这是一个测试可执行文件。我们应该能够使用一个无操作（no-op）标志来运行它，以检查总体上的执行支持情况。
# <翻译结束>


<原文开始>
// MustHaveExecPath checks that the current system can start the named executable
// using os.StartProcess or (more commonly) exec.Command.
// If not, MustHaveExecPath calls t.Skip with an explanation.
<原文结束>

# <翻译开始>
// MustHaveExecPath 检查当前系统是否能够使用 os.StartProcess 或（更常见地）exec.Command 启动指定的可执行文件。
// 如果不能，MustHaveExecPath 将调用 t.Skip 并给出解释。
# <翻译结束>


<原文开始>
// CleanCmdEnv will fill cmd.Env with the environment, excluding certain
// variables that could modify the behavior of the Go tools such as
// GODEBUG and GOTRACEBACK.
//
// If the caller wants to set cmd.Dir, set it before calling this function,
// so PWD will be set correctly in the environment.
<原文结束>

# <翻译开始>
// CleanCmdEnv 将向 cmd.Env 中填充环境，但排除可能修改 Go 工具行为的某些变量，如 GODEBUG 和 GOTRACEBACK。
// 
// 如果调用者想要设置 cmd.Dir，请在调用此函数之前设置，以便 PWD 在环境中被正确设置。
# <翻译结束>


<原文开始>
		// Exclude GODEBUG from the environment to prevent its output
		// from breaking tests that are trying to parse other command output.
<原文结束>

# <翻译开始>
// 从环境中排除GODEBUG，防止其输出破坏试图解析其他命令输出的测试。
# <翻译结束>


<原文开始>
// Exclude GOTRACEBACK for the same reason.
<原文结束>

# <翻译开始>
// Exclude GOTRACEBACK for the same reason.
# <翻译结束>


<原文开始>
// CommandContext is like exec.CommandContext, but:
//   - skips t if the platform does not support os/exec,
//   - sends SIGQUIT (if supported by the platform) instead of SIGKILL
//     in its Cancel function
//   - if the test has a deadline, adds a Context timeout and WaitDelay
//     for an arbitrary grace period before the test's deadline expires,
//   - fails the test if the command does not complete before the test's deadline, and
//   - sets a Cleanup function that verifies that the test did not leak a subprocess.
<原文结束>

# <翻译开始>
// CommandContext 类似于 exec.CommandContext，但有以下不同：
//   - 如果平台不支持 os/exec，则跳过 t
//   - 在 Cancel 函数中发送 SIGQUIT（如果平台支持）而不是 SIGKILL
//   - 如果测试有截止时间，会在测试截止时间前添加一个上下文超时和 WaitDelay 以提供一段任意的宽限期
//   - 如果命令在测试的截止时间之前未完成，会导致测试失败
//   - 设置一个 Cleanup 函数，用于检查测试是否泄漏了子进程
# <翻译结束>


<原文开始>
// unlimited unless the test has a deadline (to allow for interactive debugging)
<原文结束>

# <翻译开始>
// 除非测试有截止日期（为了允许交互式调试），否则是无限的
# <翻译结束>


<原文开始>
			// Start with a minimum grace period, just long enough to consume the
			// output of a reasonable program after it terminates.
<原文结束>

# <翻译开始>
// 从一个最小的超时时间开始，足够长的时间来消耗一个合理程序终止后的输出。
# <翻译结束>


<原文开始>
			// If time allows, increase the termination grace period to 5% of the
			// test's remaining time.
<原文结束>

# <翻译开始>
// 如果时间允许，将终止宽限期增加到测试剩余时间的5%。
# <翻译结束>


<原文开始>
			// When we run commands that execute subprocesses, we want to reserve two
			// grace periods to clean up: one for the delay between the first
			// termination signal being sent (via the Cancel callback when the Context
			// expires) and the process being forcibly terminated (via the WaitDelay
			// field), and a second one for the delay between the process being
			// terminated and the test logging its output for debugging.
			//
			// (We want to ensure that the test process itself has enough time to
			// log the output before it is also terminated.)
<原文结束>

# <翻译开始>
// 当我们运行产生子进程的命令时，我们需要预留两个宽限期来清理：
// 一个是从首次发送终止信号（通过Cancel回调在Context超时时）到
// 进程被强制终止（通过WaitDelay字段）之间的延迟，
// 另一个是进程被终止和测试记录其输出以供调试之间的延迟。
//
// （我们希望确保在测试进程本身被终止之前有足够的时间记录输出。）
# <翻译结束>


<原文开始>
				// Either ctx doesn't have a deadline, or its deadline would expire
				// after (or too close before) the test has already timed out.
				// Add a shorter timeout so that the test will produce useful output.
<原文结束>

# <翻译开始>
// 如果上下文没有截止日期，或者其截止日期在测试超时时（或接近）已经过期。
// 添加一个较短的超时时间，以便测试能够产生有用的输出。
# <翻译结束>


<原文开始>
			// The command is being terminated due to ctx being canceled, but
			// apparently not due to an explicit test deadline that we added.
			// Log that information in case it is useful for diagnosing a failure,
			// but don't actually fail the test because of it.
<原文结束>

# <翻译开始>
// 命令因 ctx 被取消而终止，但似乎不是由于我们添加的明确测试截止时间。如果这对诊断失败有帮助，请记录该信息，但不要因此实际失败测试。
# <翻译结束>


<原文开始>
// Command is like exec.Command, but applies the same changes as
// testenv.CommandContext (with a default Context).
<原文结束>

# <翻译开始>
// Command 类似于 exec.Command，但会应用与 testenv.CommandContext（使用默认的 Context）相同的变化。
# <翻译结束>

