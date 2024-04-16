
<原文开始>
// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// 版权所有 ? 2015 The Go Authors。保留所有权利。
// 本源代码的使用受 BSD 风格许可协议约束，
// 该协议可在 LICENSE 文件中找到。
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
// MustHaveExec 检查当前系统是否能够使用 os.StartProcess 或（更常见地）exec.Command 启动新的进程。
// 如果不能，MustHaveExec 将调用 t.Skip 并附带解释原因。
//
// 在某些平台上，MustHaveExec 通过重新执行当前可执行文件来检查 exec 支持，
// 这个可执行文件必须是由 'go test' 构建的二进制文件。我们故意不提供 HasExec 函数，
// 因为 TestMain 函数中存在不合适递归的风险。
//
// 要在测试之外检查 exec 支持，只需尝试执行命令即可。如果 exec 不被支持，
// 对于由此产生的错误，testenv.SyscallIsNotSupported 将返回 true。
# <翻译结束>


<原文开始>
// Assume that exec always works on non-mobile platforms and Android.
<原文结束>

# <翻译开始>
// 假设 exec 在非移动平台和 Android 上始终能正常工作
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
// 在 iOS 上存在 exec 系统调用，但在实际的 iOS 设备上，它可能会返回权限错误。在模拟环境中（如 Corellium 主机），exec 可能会成功执行。因此，如果我们需要使用 exec，就只能尝试一下并观察结果。
//
// 截至 2023-04-19，wasip1 和 js 完全没有 exec 系统调用，但我们也应采用相同的处理路径，以便在没有 iOS 环境的情况下对这一分支进行测试。
# <翻译结束>


<原文开始>
		// This isn't a standard 'go test' binary, so we don't know how to
		// self-exec in a way that should succeed without side effects.
		// Just forget it.
<原文结束>

# <翻译开始>
// 这并不是一个标准的“go test”二进制文件，因此我们不知道如何以一种没有副作用的方式使其自执行。
// 算了，放弃吧。
# <翻译结束>


<原文开始>
	// We know that this is a test executable. We should be able to run it with a
	// no-op flag to check for overall exec support.
<原文结束>

# <翻译开始>
// 我们知道这是一个测试可执行文件。我们应该能够使用一个无操作（no-op）标志来运行它，以检查整体执行支持情况。
# <翻译结束>


<原文开始>
// MustHaveExecPath checks that the current system can start the named executable
// using os.StartProcess or (more commonly) exec.Command.
// If not, MustHaveExecPath calls t.Skip with an explanation.
<原文结束>

# <翻译开始>
// MustHaveExecPath 检查当前系统是否能够使用 os.StartProcess 或（更常见地）exec.Command 启动指定的可执行文件。
// 若不能，MustHaveExecPath 将调用 t.Skip 并附带解释原因。
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
// CleanCmdEnv 会将环境变量填充到 cmd.Env 中，但会排除某些可能修改 Go 工具行为的特定变量，
// 如 GODEBUG 和 GOTRACEBACK。
// 
// 如果调用者想要设置 cmd.Dir，请在调用本函数前完成设置，以便环境中PWD能够正确设定。
# <翻译结束>


<原文开始>
		// Exclude GODEBUG from the environment to prevent its output
		// from breaking tests that are trying to parse other command output.
<原文结束>

# <翻译开始>
// 从环境中排除GODEBUG，以防止其输出干扰那些试图解析其他命令输出的测试。
# <翻译结束>


<原文开始>
// Exclude GOTRACEBACK for the same reason.
<原文结束>

# <翻译开始>
// 同理，排除GOTRACEBACK
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
// CommandContext 与 exec.CommandContext 类似，但具有以下特点：
//   - 若平台不支持 os/exec，则跳过 t；
//   - 在其 Cancel 函数中发送 SIGQUIT（若平台支持）而非 SIGKILL；
//   - 若测试具有截止时间，则在测试截止时间前设置一段任意的宽限期，添加 Context 超时及 WaitDelay；
//   - 若命令未能在测试截止时间前完成，导致测试失败；
//   - 设置一个 Cleanup 函数，用于验证测试未泄露子进程。
# <翻译结束>


<原文开始>
// unlimited unless the test has a deadline (to allow for interactive debugging)
<原文结束>

# <翻译开始>
// 除非测试具有截止时间（以允许进行交互式调试），否则不受限制
# <翻译结束>


<原文开始>
			// Start with a minimum grace period, just long enough to consume the
			// output of a reasonable program after it terminates.
<原文结束>

# <翻译开始>
// 以一个最小的宽限期开始，这个时间长度仅足以在程序终止后接收其合理的输出。
# <翻译结束>


<原文开始>
			// If time allows, increase the termination grace period to 5% of the
			// test's remaining time.
<原文结束>

# <翻译开始>
// 如果时间允许，将终止宽限期增加至测试剩余时间的5%。
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
// 当我们执行涉及子进程的命令时，我们希望预留两个清理阶段：
// 一是从发送第一个终止信号（通过Context过期时的Cancel回调）到强制终止进程（通过WaitDelay字段）之间的时间；
// 二是从进程被终止到测试为调试目的记录其输出之间的时间。
//
// （我们需要确保测试进程自身有足够时间在被终止前记录下输出。）
# <翻译结束>


<原文开始>
				// Either ctx doesn't have a deadline, or its deadline would expire
				// after (or too close before) the test has already timed out.
				// Add a shorter timeout so that the test will produce useful output.
<原文结束>

# <翻译开始>
// 或者ctx没有截止时间，或者其截止时间将在（或过早于）测试已超时时到期。
// 添加一个较短的超时时间，以便测试能产生有用输出。
# <翻译结束>


<原文开始>
			// The command is being terminated due to ctx being canceled, but
			// apparently not due to an explicit test deadline that we added.
			// Log that information in case it is useful for diagnosing a failure,
			// but don't actually fail the test because of it.
<原文结束>

# <翻译开始>
// 由于ctx被取消，命令正在被终止，但
// 显然并非由于我们添加的显式测试截止时间导致。
// 将此信息记录下来，以便于诊断失败原因，
// 但切勿因此而令测试失败。
# <翻译结束>


<原文开始>
// Command is like exec.Command, but applies the same changes as
// testenv.CommandContext (with a default Context).
<原文结束>

# <翻译开始>
// Command类似于exec.Command，但会应用与testenv.CommandContext相同的变化（使用默认的Context）。
# <翻译结束>

