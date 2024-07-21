
<原文开始>
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// 版权所有 ? 2009 Go作者。保留所有权利。
// 本源代码的使用受BSD风格
// 许可证约束，该许可证可在LICENSE文件中找到。
# <翻译结束>


<原文开始>
// Use an external test to avoid os/exec -> net/http -> crypto/x509 -> os/exec
// circular dependency on non-cgo darwin.
<原文结束>

# <翻译开始>
// 使用外部测试以避免在非cgo环境下的darwin系统中出现os/exec -> net/http -> crypto/x509 -> os/exec的循环依赖关系
# <翻译结束>


<原文开始>
// haveUnexpectedFDs is set at init time to report whether any file descriptors
// were open at program start.
<原文结束>

# <翻译开始>
// haveUnexpectedFDs在初始化时被设置，用于报告程序启动时是否有任何文件描述符处于打开状态。
# <翻译结束>


<原文开始>
// TestMain allows the test binary to impersonate many other binaries,
// some of which may manipulate os.Stdin, os.Stdout, and/or os.Stderr
// (and thus cannot run as an ordinary Test function, since the testing
// package monkey-patches those variables before running tests).
<原文结束>

# <翻译开始>
// TestMain 允许测试二进制程序模拟多个其他程序，
// 其中一些可能操作 os.Stdin、os.Stdout 和/或 os.Stderr
// （因此无法作为普通 Test 函数运行，因为 testing 包在运行测试前会对这些变量进行monkey patch）。
# <翻译结束>


<原文开始>
			// Run a couple of GC cycles to increase the odds of detecting
			// process leaks using the finalizers installed by GODEBUG=execwait=2.
<原文结束>

# <翻译开始>
			// 运行几个GC周期，以提高通过由GODEBUG=execwait=2安装的终结器检测进程泄漏的可能性
# <翻译结束>


<原文开始>
// registerHelperCommand registers a command that the test process can impersonate.
// A command should be registered in the same source file in which it is used.
// If all tests are run and pass, all registered commands must be used.
// (This prevents stale commands from accreting if tests are removed or
// refactored over time.)
<原文结束>

# <翻译开始>
// registerHelperCommand 注册一个测试进程可以模拟执行的命令。
// 命令应在使用它的同一源文件中进行注册。
// 如果所有测试都运行且通过，那么所有已注册的命令必须被使用。
// （这可以防止在随时间移除或重构测试时积聚过时的命令。）
# <翻译结束>


<原文开始>
// maySkipHelperCommand records that the test that uses the named helper command
// was invoked, but may call Skip on the test before actually calling
// helperCommand.
<原文结束>

# <翻译开始>
// maySkipHelperCommand 用于记录使用指定辅助命令的测试已调用，但在实际调用 helperCommand 之前可能调用 Skip。
# <翻译结束>


<原文开始>
// helperCommand returns an exec.Cmd that will run the named helper command.
<原文结束>

# <翻译开始>
// helperCommand 返回一个 exec.Cmd，该命令将运行指定名称的辅助命令。
# <翻译结束>


<原文开始>
// helperCommandContext is like helperCommand, but also accepts a Context under
// which to run the command.
<原文结束>

# <翻译开始>
// helperCommandContext 与 helperCommand 类似，但还接受一个在其下运行命令的 Context。
# <翻译结束>


<原文开始>
// exePath returns the path to the running executable.
<原文结束>

# <翻译开始>
// exePath 返回运行中可执行文件的路径
# <翻译结束>


<原文开始>
		// Use os.Executable instead of os.Args[0] in case the caller modifies
		// cmd.Dir: if the test binary is invoked like "./exec.test", it should
		// not fail spuriously.
<原文结束>

# <翻译开始>
		// 使用os.Executable替代os.Args[0]，以防调用者修改
		// cmd.Dir：如果以"./exec.test"形式调用测试二进制文件，
		// 则不应出现异常失败。
# <翻译结束>


<原文开始>
			// Couldn't chdir back to the original working directory.
			// panic instead of t.Fatal so that we don't run other tests
			// in an unexpected location.
<原文结束>

# <翻译开始>
			// 无法切换回原始工作目录。
			// 使用panic替代t.Fatal，以防止我们在未知位置运行其他测试。
# <翻译结束>


<原文开始>
	// Run our own binary as a relative path
	// (e.g. "_test/exec.test") our parent directory.
<原文结束>

# <翻译开始>
	// 以相对路径方式运行我们自己的二进制文件（例如 "_test/exec.test"），该路径位于父目录中。
# <翻译结束>


<原文开始>
// "/tmp/go-buildNNNN/os/exec/_test"
<原文结束>

# <翻译开始>
// "/tmp/go-buildNNNN/os/exec/_test"
# <翻译结束>


<原文开始>
// "/tmp/go-buildNNNN/os/exec"
<原文结束>

# <翻译开始>
// "/tmp/go-buildNNNN/os/exec"
# <翻译结束>


<原文开始>
// Cat, testing stdin and stdout.
<原文结束>

# <翻译开始>
// Cat，用于测试标准输入和标准输出
# <翻译结束>


<原文开始>
// Testing combined output and error values.
<原文结束>

# <翻译开始>
// 测试输出与错误值的组合情况
# <翻译结束>


<原文开始>
// Can't run a non-existent executable
<原文结束>

# <翻译开始>
// 无法运行不存在的可执行文件
# <翻译结束>


<原文开始>
// Test that exit values are returned correctly
<原文结束>

# <翻译开始>
// 测试退出值是否正确返回
# <翻译结束>


<原文开始>
// Test that exit code are returned correctly
<原文结束>

# <翻译开始>
// 测试返回的退出码是否正确
# <翻译结束>


<原文开始>
// Test when command does not call Run().
<原文结束>

# <翻译开始>
// 测试命令不调用Run()的情况
# <翻译结束>


<原文开始>
// Check that we can access methods of the underlying os.File.`
<原文结束>

# <翻译开始>
// 检查我们是否可以访问底层os.File的方法
# <翻译结束>


<原文开始>
// Before the fix, this next line would race with cmd.Wait.
<原文结束>

# <翻译开始>
// 修复前，下一行代码将与 cmd.Wait 产生竞态条件。
# <翻译结束>


<原文开始>
// Issue 17647.
// It used to be the case that TestStdinClose, above, would fail when
// run under the race detector. This test is a variant of TestStdinClose
// that also used to fail when run under the race detector.
// This test is run by cmd/dist under the race detector to verify that
// the race detector no longer reports any problems.
<原文结束>

# <翻译开始>
// 问题 17647。
// 之前，当在竞态检测器下运行时，上述的 TestStdinClose 测试会失败。
// 本测试是 TestStdinClose 的一个变种，同样在竞态检测器下运行时也会失败。
// 此测试由 cmd/dist 在竞态检测器下运行，用于验证竞态检测器不再报告任何问题。
# <翻译结束>


<原文开始>
		// We don't check the error return of Kill. It is
		// possible that the process has already exited, in
		// which case Kill will return an error "process
		// already finished". The purpose of this test is to
		// see whether the race detector reports an error; it
		// doesn't matter whether this Kill succeeds or not.
<原文结束>

# <翻译开始>
		// 我们不对 Kill 的错误返回值进行检查。有可能该进程已经退出，此时 Kill 会返回一个“process already finished”的错误。本测试的目的是观察竞态检测器是否报告错误；Kill 是否成功并不重要。
# <翻译结束>


<原文开始>
		// Send the wrong string, so that the child fails even
		// if the other goroutine doesn't manage to kill it first.
		// This test is to check that the race detector does not
		// falsely report an error, so it doesn't matter how the
		// child process fails.
<原文结束>

# <翻译开始>
		// 发送错误的字符串，确保即使其他goroutine未能先将其终止，子进程也能失败。
		// 本测试旨在检查竞态检测器不会错误地报告错误，因此子进程如何失败并不重要。
# <翻译结束>


<原文开始>
// Not parallel: checks for leaked file descriptors
<原文结束>

# <翻译开始>
// 非并行：检查是否有泄露的文件描述符
# <翻译结束>


<原文开始>
	// Since this test is not running in parallel, we don't expect any new file
	// descriptors to be opened while it runs. However, if there are additional
	// FDs present at the start of the test (for example, opened by libc), those
	// may be closed due to a timeout of some sort. Allow those to go away, but
	// check that no new FDs are added.
<原文结束>

# <翻译开始>
	// 由于此测试并非并行运行，我们预期在它执行过程中不会打开任何新的文件描述符。然而，如果在测试开始时存在额外的FD（例如由libc打开），这些FD可能因某种超时而被关闭。允许它们消失，但要检查没有新增加的FD。
# <翻译结束>


<原文开始>
		// The point of this test is to make sure that any
		// descriptors we open are marked close-on-exec.
		// If haveUnexpectedFDs is true then there were other
		// descriptors open when we started the test,
		// so those descriptors are clearly not close-on-exec,
		// and they will confuse the test. We could modify
		// the test to expect those descriptors to remain open,
		// but since we don't know where they came from or what
		// they are doing, that seems fragile. For example,
		// perhaps they are from the startup code on this
		// system for some reason. Also, this test is not
		// system-specific; as long as most systems do not skip
		// the test, we will still be testing what we care about.
<原文结束>

# <翻译开始>
		// 本测试的目的是确保我们打开的所有描述符都被标记为“close-on-exec”。
		// 如果`haveUnexpectedFDs`为真，说明我们在开始测试时已有其他描述符处于打开状态，
		// 显然这些描述符没有设置“close-on-exec”，它们会干扰本次测试。虽然我们可以修改测试，
		// 使其预期这些描述符保持打开，但由于我们并不清楚它们来自何处、有何作用，
		// 这种做法显得很脆弱。例如，可能是因为某种原因，它们是由当前系统的启动代码所创建的。
		// 另外，本测试并非针对特定系统；只要大部分系统未跳过此测试，我们仍能对关注点进行有效验证。
# <翻译结束>


<原文开始>
	// This test runs with cgo disabled. External linking needs cgo, so
	// it doesn't work if external linking is required.
<原文结束>

# <翻译开始>
	// 该测试在禁用cgo的情况下运行。外部链接需要cgo，因此如果需要外部链接，则无法正常工作。
# <翻译结束>


<原文开始>
	// Force network usage, to verify the epoll (or whatever) fd
	// doesn't leak to the child,
<原文结束>

# <翻译开始>
	// 强制使用网络，以验证 epoll（或其他类似机制）的文件描述符
	// 不会泄漏给子进程
# <翻译结束>


<原文开始>
// Make sure duplicated fds don't leak to the child.
<原文结束>

# <翻译开始>
// 确保重复的文件描述符不会泄露给子进程。
# <翻译结束>


<原文开始>
	// Force TLS root certs to be loaded (which might involve
	// cgo), to make sure none of that potential C code leaks fds.
<原文结束>

# <翻译开始>
	// 强制加载TLS根证书（可能涉及cgo），以确保潜在的C代码不会泄漏文件描述符。
# <翻译结束>


<原文开始>
// quiet expected TLS handshake error "remote error: bad certificate"
<原文结束>

# <翻译开始>
// 忽略预期的TLS握手错误“remote error: bad certificate”
# <翻译结束>


<原文开始>
	// Build the test without cgo, so that C library functions don't
	// open descriptors unexpectedly. See issue 25628.
<原文结束>

# <翻译开始>
	// 在无cgo的情况下构建测试，以便C库函数不会意外打开描述符。参见问题25628。
# <翻译结束>


<原文开始>
// Use a deadline to try to get some output even if the program hangs.
<原文结束>

# <翻译开始>
// 使用截止时间来尝试获取输出，即使程序挂起也能有所收获
# <翻译结束>


<原文开始>
		// Leave a 20% grace period to flush output, which may be large on the
		// linux/386 builders because we're running the subprocess under strace.
<原文结束>

# <翻译开始>
		// 留出20%的宽限期以刷新输出，这在 linux/386 构建器上可能很大，因为我们正在使用 strace 运行子进程。
# <翻译结束>


<原文开始>
		// Some facilities in illumos are implemented via access
		// to /proc by libc; such accesses can briefly occupy a
		// low-numbered fd.  If this occurs concurrently with the
		// test that checks for leaked descriptors, the check can
		// become confused and report a spurious leaked descriptor.
		// (See issue #42431 for more detailed analysis.)
		//
		// Attempt to constrain the use of additional threads in the
		// child process to make this test less flaky:
<原文结束>

# <翻译开始>
		// illumos 中某些设施的实现依赖于 libc 对 /proc 的访问，这种访问可能会暂时占用一个低位文件描述符（fd）。如果这种情况恰好与检查泄漏描述符的测试并发发生，该检查可能会变得混乱，并错误地报告一个泄漏的描述符。（关于更详细的分析，请参见问题 #42431。）
		// 
		// 试图限制子进程中额外线程的使用，以减少此测试出现不稳定性：
# <翻译结束>


<原文开始>
// Issue 9173: ignore stdin pipe writes if the program completes successfully.
<原文结束>

# <翻译开始>
// Issue 9173：若程序成功完成，忽略对stdin管道的写入操作。
# <翻译结束>


<原文开始>
	// To reduce noise in the final goroutine dump,
	// let other parallel tests complete if possible.
<原文结束>

# <翻译开始>
	// 为了减少最终协程转储中的噪声，
	// 尽可能让其他并行测试完成。
# <翻译结束>


<原文开始>
// At this point the process is alive. Ensure it by sending data to stdin.
<原文结束>

# <翻译开始>
// 此时进程处于活跃状态。通过向其stdin发送数据来确保这一点。
# <翻译结束>


<原文开始>
	// Calling cancel should have killed the process, so writes
	// should now fail.  Give the process a little while to die.
<原文结束>

# <翻译开始>
	// 调用cancel应当已终止该进程，因此现在写入应会失败。给进程一点时间结束。
# <翻译结束>


<原文开始>
			// Panic instead of calling t.Fatal so that we get a goroutine dump.
			// We want to know exactly what the os/exec goroutines got stuck on.
<原文结束>

# <翻译开始>
			// 通过 panic 替代调用 t.Fatal，以便我们获取 goroutine 堆栈信息。
			// 我们希望确切了解 os/exec goroutines 阻塞在何处。
# <翻译结束>


<原文开始>
		// Back off exponentially (up to 1-second sleeps) to give the OS time to
		// terminate the process.
<原文结束>

# <翻译开始>
		// 以指数方式回退（最多休眠1秒），以便给操作系统时间来终止进程。
# <翻译结束>


<原文开始>
// test that environment variables are de-duped.
<原文结束>

# <翻译开始>
// 测试环境变量是否被去重
# <翻译结束>


<原文开始>
// TestDoubleStartLeavesPipesOpen checks for a regression in which calling
// Start twice, which returns an error on the second call, would spuriously
// close the pipes established in the first call.
<原文结束>

# <翻译开始>
// TestDoubleStartLeavesPipesOpen 测试一个回归问题，即当两次调用
// Start 方法时（第二次调用会返回错误），是否会导致第一次调用中建立的管道被错误地关闭。
# <翻译结束>


<原文开始>
// Release resources if cmd happens not to outlive this process.
<原文结束>

# <翻译开始>
// 如果cmd未能在此进程之后继续存在，则释放资源
# <翻译结束>


<原文开始>
// Signal that the process is set up by closing stdout.
<原文结束>

# <翻译开始>
// 通过关闭stdout来表明进程已设置完毕
# <翻译结束>


<原文开始>
// Ignore write errors: we want to keep reading even if stderr is closed.
<原文结束>

# <翻译开始>
// 忽略写入错误：即使stderr被关闭，我们也希望继续读取。
# <翻译结束>


<原文开始>
// A tickReader reads an unbounded sequence of timestamps at no more than a
// fixed interval.
<原文结束>

# <翻译开始>
// tickReader 读取一个无界的时间戳序列，但每个时间戳之间的间隔不会超过一个固定的值。
# <翻译结束>


<原文开始>
// Wait for cmd to close stdout to signal that its handlers are installed.
<原文结束>

# <翻译开始>
// 等待cmd关闭stdout以表明其处理器已安装完成
# <翻译结束>


<原文开始>
	// tooLong is an arbitrary duration that is expected to be much longer than
	// the test runs, but short enough that leaked processes will eventually exit
	// on their own.
<原文结束>

# <翻译开始>
	// tooLong 是一个任意的时长，预期远超过测试运行所需时间，但又足够短，以至于泄露的进程最终能够自行退出
# <翻译结束>


<原文开始>
	// Control case: with no cancellation and no WaitDelay, we should wait for the
	// process to exit.
<原文结束>

# <翻译开始>
	// 控制用例：在没有取消操作且未设置等待延迟的情况下，我们应该等待进程退出。
# <翻译结束>


<原文开始>
	// With a very long WaitDelay and no Cancel function, we should wait for the
	// process to exit even if the command's Context is cancelled.
<原文结束>

# <翻译开始>
	// 当WaitDelay设置得非常长且没有Cancel函数时，即使命令的Context被取消，我们也应该等待进程退出。
# <翻译结束>


<原文开始>
		// At this point cmd should still be running (because we passed nil to
		// startHang for the cancel signal). Sending it an explicit Interrupt signal
		// should succeed.
<原文结束>

# <翻译开始>
		// 此时，cmd 应该仍在运行（因为我们向 startHang 传递了 nil 作为取消信号）。向其发送一个明确的 Interrupt 信号应该能成功。
# <翻译结束>


<原文开始>
		// This program exits with status 0,
		// but pretty much always does so during the wait delay.
		// Since the Cmd itself didn't do anything to stop the process when the
		// context expired, a successful exit is valid (even if late) and does
		// not merit a non-nil error.
<原文结束>

# <翻译开始>
		// 该程序以状态0退出，
		// 但在等待延迟期间几乎总是如此。
		// 由于Cmd本身在上下文过期时没有做任何事情来停止进程，
		// 因此即使退出较晚，成功退出也是有效的，
		// 并不值得返回一个非空错误。
# <翻译结束>


<原文开始>
	// If the context is cancelled and the Cancel function sends os.Kill,
	// the process should be terminated immediately, and its output
	// pipes should be closed (causing Wait to return) after WaitDelay
	// even if a child process is still writing to them.
<原文结束>

# <翻译开始>
	// 如果上下文被取消且Cancel函数发送os.Kill信号，
	// 则进程应立即终止，且其输出管道应在WaitDelay后关闭（导致Wait返回），
	// 即使子进程仍在向这些管道写入数据。
# <翻译结束>


<原文开始>
		// This test should kill the child process after 10ms,
		// leaving a grandchild process writing probes in a loop.
		// The child process should be reported as failed,
		// and the grandchild will exit (or die by SIGPIPE) once the
		// stderr pipe is closed.
<原文结束>

# <翻译开始>
		// 该测试应在10毫秒后终止子进程，
		// 留下一个孙辈进程以循环方式写入探针。
		// 子进程应被报告为失败状态，
		// 当标准错误管道关闭时，孙辈进程将退出（或因SIGPIPE信号而终止）。
# <翻译结束>


<原文开始>
		// This child process should exit immediately,
		// leaving a grandchild process writing probes in a loop.
		// Since the child has no ExitError to report but we did not
		// read all of its output, Wait should return ErrWaitDelay.
<原文结束>

# <翻译开始>
		// 该子进程应立即退出，
		// 留下孙进程以循环方式写入探测数据。
		// 由于子进程没有 ExitError 需要报告，但我们尚未读取其所有输出，
		// 因此 Wait 函数应返回 ErrWaitDelay。
# <翻译结束>


<原文开始>
	// If the Cancel function sends a signal that the process can handle, and it
	// handles that signal without actually exiting, then it should be terminated
	// after the WaitDelay.
<原文结束>

# <翻译开始>
	// 如果Cancel函数发送了一个进程可以处理的信号，且该进程在接收到该信号后并未实际退出，而是进行了某种处理，则应在WaitDelay后将其强制终止。
# <翻译结束>


<原文开始>
		// This command ignores SIGINT, sleeping until it is killed.
		// Wait should return the usual error for a killed process.
<原文结束>

# <翻译开始>
		// 该命令忽略 SIGINT 信号，持续休眠直到被杀死。
		// Wait 方法应返回被杀死进程通常的错误信息。
# <翻译结束>


<原文开始>
	// If the process handles the cancellation signal and exits with status 0,
	// Wait should report a non-nil error (because the process had to be
	// interrupted), and it should be a context error (because there is no error
	// to report from the child process itself).
<原文结束>

# <翻译开始>
	// 如果进程捕获到了取消信号并以状态0退出，
	// 那么Wait应当返回一个非nil的错误（因为进程是被中断的），
	// 并且它应该是一个上下文错误（因为子进程中没有要报告的错误）。
# <翻译结束>


<原文开始>
	// If the Cancel function sends SIGQUIT, it should be handled in the usual
	// way: a Go program should dump its goroutines and exit with non-success
	// status. (We expect SIGQUIT to be a common pattern in real-world use.)
<原文结束>

# <翻译开始>
	// 如果Cancel函数发送SIGQUIT信号，它应以通常的方式进行处理：Go程序应输出其goroutine信息并以非成功状态退出。（我们预期SIGQUIT在实际使用中会成为一种常见模式。）
# <翻译结束>


<原文开始>
// The default os/signal handler exits with code 2.
<原文结束>

# <翻译开始>
// 默认的os/signal处理器以代码2退出
# <翻译结束>


<原文开始>
	// If Cancel returns a non-ErrProcessDone error and the process
	// exits successfully, Wait should wrap the error from Cancel.
<原文结束>

# <翻译开始>
	// 如果Cancel返回非ErrProcessDone的错误且进程成功退出，Wait应当封装Cancel的错误
# <翻译结束>


<原文开始>
		// We intentionally race Cancel against the process exiting,
		// but ensure that the process wins the race (and return ErrProcessDone
		// from Cancel to report that).
<原文结束>

# <翻译开始>
		// 我们特意让Cancel与进程退出进行竞争，但确保进程赢得此竞争（并在Cancel中返回ErrProcessDone以报告此情况）。
# <翻译结束>


<原文开始>
// reaches EOF when the process exits
<原文结束>

# <翻译开始>
// 当进程退出时到达EOF
# <翻译结束>


<原文开始>
	// If Cancel returns an error and the process is killed after
	// WaitDelay, Wait should report the usual SIGKILL ExitError, not the
	// error from Cancel.
<原文结束>

# <翻译开始>
	// 如果 Cancel 返回错误，并且在经过 WaitDelay 后进程被杀掉，那么 Wait 应该报告常规的 SIGKILL ExitError，而不是来自 Cancel 的错误。
# <翻译结束>


<原文开始>
		// Ensure that Cancel actually had the opportunity to
		// return the error.
<原文结束>

# <翻译开始>
		// 确保 Cancel 实际上有机会返回该错误
# <翻译结束>


<原文开始>
		// This test should kill the child process after 1ms,
		// To maximize compatibility with existing uses of exec.CommandContext, the
		// resulting error should be an exec.ExitError without additional wrapping.
<原文结束>

# <翻译开始>
		// 本测试应在1毫秒后终止子进程，
		// 为最大程度地与现有exec.CommandContext用法兼容，产生的错误应为未经额外封装的exec.ExitError。
# <翻译结束>


<原文开始>
	// If Cancel returns ErrProcessDone but the process is not actually done
	// (and has to be killed), Wait should report the usual SIGKILL ExitError,
	// not the error from Cancel.
<原文结束>

# <翻译开始>
	// 如果 Cancel 返回 ErrProcessDone，但实际上进程并未完成（仍需被杀死），则 Wait 应报告常规的 SIGKILL ExitError，而非来自 Cancel 的错误。
# <翻译结束>


<原文开始>
	// If Cancel returns an error and the process exits with an
	// unsuccessful exit code, the process error should take precedence over the
	// Cancel error.
<原文结束>

# <翻译开始>
	// 如果Cancel返回错误并且进程以非正常退出码结束，则应优先考虑进程错误而非Cancel错误。
# <翻译结束>


<原文开始>
// TestConcurrentExec is a regression test for https://go.dev/issue/61080.
//
// Forking multiple child processes concurrently would sometimes hang on darwin.
// (This test hung on a gomote with -count=100 after only a few iterations.)
<原文结束>

# <翻译开始>
// TestConcurrentExec 是针对 https://go.dev/issue/61080 的回归测试。
// 
// 在 Darwin 系统上，有时并发启动多个子进程会导致程序挂起。
// （在 gomote 上运行时，仅经过几次迭代后，在 -count=100 选项下该测试即出现挂起。）
# <翻译结束>


<原文开始>
	// This test will spawn nHangs subprocesses that hang reading from stdin,
	// and nExits subprocesses that exit immediately.
	//
	// When issue #61080 was present, a long-lived "hang" subprocess would
	// occasionally inherit the fork/exec status pipe from an "exit" subprocess,
	// causing the parent process (which expects to see an EOF on that pipe almost
	// immediately) to unexpectedly block on reading from the pipe.
<原文结束>

# <翻译开始>
	// 本测试将启动 nHangs 个挂起的子进程，它们从 stdin 读取数据时会挂起，
	// 同时启动 nExits 个立即退出的子进程。
	// 
	// 当存在 #61080 问题时，一个长期运行的“挂起”子进程偶尔会继承自“退出”子进程的 fork/exec 状态管道，
	// 导致父进程（期望几乎立刻在该管道上看到 EOF）意外地阻塞在从管道读取的操作上。
# <翻译结束>


<原文开始>
	// ready is done when the goroutines have done as much work as possible to
	// prepare to create subprocesses. It isn't strictly necessary for the test,
	// but helps to increase the repro rate by making it more likely that calls to
	// syscall.StartProcess for the "hang" and "exit" goroutines overlap.
<原文结束>

# <翻译开始>
	// ready表示当goroutine已经尽可能地完成了创建子进程的准备工作。虽然严格来说这不是测试所必需的，但通过增加“hang”和“exit”goroutine对syscall.StartProcess调用发生重叠的可能性，有助于提高问题复现率。
# <翻译结束>


<原文开始>
// TestPathRace tests that [Cmd.String] can be called concurrently
// with [Cmd.Start].
<原文结束>

# <翻译开始>
// TestPathRace 测试 [Cmd.String] 可以与 [Cmd.Start] 并发调用
# <翻译结束>


<原文开始>
	//for fd := uintptr(3); fd <= 100; fd++ {
	//	if poll.IsPollDescriptor(fd) {
	//		continue
	//	}
	//
	//	if fdtest.Exists(fd) {
	//		haveUnexpectedFDs = true
	//		return
	//	}
	//}
<原文结束>

# <翻译开始>
	// 对于fd从uintptr类型数值3开始，到数值100结束的循环：
	// 如果poll.IsPollDescriptor(fd)判断fd为poll描述符，
	// 则跳过本次循环继续下一轮。
	//
	// 若fdtest.Exists(fd)检测到fd存在，
	// 则将haveUnexpectedFDs设为true并立即返回。
# <翻译结束>

