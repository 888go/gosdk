
<原文开始>
// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// 版权所有 ? 2017 The Go Authors。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。
# <翻译结束>


<原文开始>
// If NoSetGroups is true, setgroups isn't called and cmd.Run should succeed
<原文结束>

# <翻译开始>
// 如果NoSetGroups 为真，则不会调用 setgroups，且 cmd.Run 应该能成功执行
# <翻译结束>


<原文开始>
// For issue #19314: make sure that SIGSTOP does not cause the process
// to appear done.
<原文结束>

# <翻译开始>
// 为解决第19314号问题：确保SIGSTOP信号不会导致进程看似已完成。
# <翻译结束>


<原文开始>
// Wait for the child process to come up and register any signal handlers.
<原文结束>

# <翻译开始>
// 等待子进程启动，并注册任何信号处理器。
# <翻译结束>


<原文开始>
// Now leave the pipes open so that the process will hang until we close stdin.
<原文结束>

# <翻译开始>
// 现在保持管道打开，以便进程会挂起直到我们关闭标准输入。
# <翻译结束>


<原文开始>
	// Give a little time for Wait to block on waiting for the process.
	// (This is just to give some time to trigger the bug; it should not be
	// necessary for the test to pass.)
<原文结束>

# <翻译开始>
	// 给`Wait`一点时间来阻塞等待进程。
	// （这只是为了触发bug提供一些时间；对于测试通过来说，这并不必要。）
# <翻译结束>


<原文开始>
	// This call to Signal should succeed because the process still exists.
	// (Prior to the fix for #19314, this would fail with os.ErrProcessDone
	// or an equivalent error.)
<原文结束>

# <翻译开始>
	// 此调用Signal应成功，因为进程仍然存在。
	// （在修复#19314问题之前，这将因os.ErrProcessDone或等效错误而失败。）
# <翻译结束>


<原文开始>
	// The SIGCONT should allow the process to wake up, notice that stdin
	// is closed, and exit successfully.
<原文结束>

# <翻译开始>
	// SIGCONT 信号应能使进程唤醒，注意到stdin已关闭，并成功退出。
# <翻译结束>


<原文开始>
// https://go.dev/issue/50599: if Env is not set explicitly, setting Dir should
// implicitly update PWD to the correct path, and Environ should list the
// updated value.
<原文结束>

# <翻译开始>
// https://go.dev/issue/50599：若未明确设置 Env，当设置 Dir 时应隐式更新 PWD 至正确路径，并且 Environ 应列出已更新的值。
# <翻译结束>


<原文开始>
// However, if cmd.Env is set explicitly, setting Dir should not override it.
// (This checks that the implementation for https://go.dev/issue/50599 doesn't
// break existing users who may have explicitly mismatched the PWD variable.)
<原文结束>

# <翻译开始>
// 但是，如果已明确设置 cmd.Env，则设置 Dir 不应覆盖它。
// （这会检查为 https://go.dev/issue/50599 实现的代码是否不会破坏可能已显式错配 PWD 变量的现有用户。）
# <翻译结束>


<原文开始>
	// Now link is another equally-valid name for cwd. If we set Dir to one and
	// PWD to the other, the subprocess should report the PWD version.
<原文结束>

# <翻译开始>
	// 现在link是cwd的另一个同样有效的名称。如果我们将Dir设置为其中一个，而PWD设置为另一个，则子进程应报告PWD版本。
# <翻译结束>


<原文开始>
		// Ideally we would also like to test what happens if we set PWD to
		// something totally bogus (or the empty string), but then we would have no
		// idea what output the subprocess should actually produce: cwd itself may
		// contain symlinks preserved from the PWD value in the test's environment.
<原文结束>

# <翻译开始>
		// 理想情况下，我们还希望测试将 PWD 设置为
		// 完全无意义的值（或空字符串）时会发生什么情况，但那样的话，
		// 我们就无法得知子进程实际应产生的输出是什么：cwd 本身可能包含
		// 测试环境中的 PWD 值所保留的符号链接。
# <翻译结束>


<原文开始>
			// This is intentionally opposite to the usual order of setting cmd.Dir
			// and then calling cmd.Environ. Here, we *want* PWD not to match cmd.Dir,
			// so we don't care whether cmd.Dir is reflected in cmd.Environ.
<原文结束>

# <翻译开始>
			// 这里故意颠倒了通常设置 cmd.Dir 然后调用 cmd.Environ 的顺序。在这里，我们**希望** PWD 与 cmd.Dir 不匹配，因此我们并不关心 cmd.Dir 是否反映在 cmd.Environ 中。
# <翻译结束>

