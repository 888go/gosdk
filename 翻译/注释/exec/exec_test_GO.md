
<原文开始>
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// 版权所有 ? 2009 Go 作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。
# <翻译结束>


<原文开始>
// Use an external test to avoid os/exec -> net/http -> crypto/x509 -> os/exec
// circular dependency on non-cgo darwin.
<原文结束>

# <翻译开始>
// 使用外部测试以避免在非cgo darwin环境下出现os/exec -> net/http -> crypto/x509 -> os/exec循环依赖问题。
# <翻译结束>


<原文开始>
// haveUnexpectedFDs is set at init time to report whether any file descriptors
// were open at program start.
<原文结束>

# <翻译开始>
// haveUnexpectedFDs 在初始化时被设置，用于报告程序启动时是否有任何文件描述符处于打开状态。
# <翻译结束>


<原文开始>
// TestMain allows the test binary to impersonate many other binaries,
// some of which may manipulate os.Stdin, os.Stdout, and/or os.Stderr
// (and thus cannot run as an ordinary Test function, since the testing
// package monkey-patches those variables before running tests).
<原文结束>

# <翻译开始>
// TestMain 函数允许测试二进制程序模拟多种其他二进制程序，
// 其中一些可能操作 os.Stdin（标准输入）、os.Stdout（标准输出）和/或 os.Stderr（标准错误），
// 因此它们无法作为普通 Test 函数运行，因为 testing 包在运行测试之前会对此类变量进行猴子补丁式的修改。
# <翻译结束>


<原文开始>
			// Run a couple of GC cycles to increase the odds of detecting
			// process leaks using the finalizers installed by GODEBUG=execwait=2.
<原文结束>

# <翻译开始>
// 运行几个垃圾回收周期，以提高通过GODEBUG=execwait=2安装的终结器检测到进程泄漏的可能性。
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
// 命令应在被使用的同一源文件中进行注册。
// 如果所有测试都被运行并通过，那么所有已注册的命令都必须被使用。
// （这可以防止在随时间移除或重构测试时，残留未使用的命令。）
# <翻译结束>


<原文开始>
// maySkipHelperCommand records that the test that uses the named helper command
// was invoked, but may call Skip on the test before actually calling
// helperCommand.
<原文结束>

# <翻译开始>
// maySkipHelperCommand 记录了使用指定辅助命令的测试已调用，
// 但在实际调用 helperCommand 之前可能调用 Skip。
# <翻译结束>


<原文开始>
// helperCommand returns an exec.Cmd that will run the named helper command.
<原文结束>

# <翻译开始>
// helperCommand 返回一个exec.Cmd，该Cmd将会运行指定名称的辅助命令。
# <翻译结束>


<原文开始>
// helperCommandContext is like helperCommand, but also accepts a Context under
// which to run the command.
<原文结束>

# <翻译开始>
// helperCommandContext 类似于 helperCommand，但还接受一个在其中运行命令的上下文。
// （"under which" 在这里可以理解为“在该上下文之下”，即在指定的上下文中运行命令）
# <翻译结束>


<原文开始>
// exePath returns the path to the running executable.
<原文结束>

# <翻译开始>
// exePath 返回运行中可执行文件的路径。
# <翻译结束>


<原文开始>
		// Use os.Executable instead of os.Args[0] in case the caller modifies
		// cmd.Dir: if the test binary is invoked like "./exec.test", it should
		// not fail spuriously.
<原文结束>

# <翻译开始>
// 在调用方可能修改 os.Args[0] 的情况下，使用 os.Executable 替代，以防万一。
// 如果以 "./exec.test" 的形式调用测试二进制文件，应避免其无故失败。
# <翻译结束>


<原文开始>
	// Run our own binary as a relative path
	// (e.g. "_test/exec.test") our parent directory.
<原文结束>

# <翻译开始>
// 以相对路径运行我们自己的二进制文件（例如："_test/exec.test"）在父目录中。
# <翻译结束>


<原文开始>
// "/tmp/go-buildNNNN/os/exec/_test"
<原文结束>

# <翻译开始>
// "/tmp/go-buildNNNN/os/exec/_test" 
// 
// 这个Go语言的注释内容是对一个文件路径的描述，翻译为：
// 
// “/tmp/go-buildNNNN/os/exec/_test” 
// 
// 这里的"/tmp/go-buildNNNN/os/exec/_test" 是一个临时目录下的文件或目录路径，用于Go语言编译测试时生成的相关文件。其中，“NNNN”表示四位数字，可能是动态生成的编号。具体而言，这是Go在运行os/exec包下的测试时，用于存放相关编译输出或临时测试文件的目录路径。
# <翻译结束>


<原文开始>
// "/tmp/go-buildNNNN/os/exec"
<原文结束>

# <翻译开始>
// "/tmp/go-buildNNNN/os/exec" 的中文翻译可能为：
// “/tmp/go-buildNNNN/os/exec”目录（或文件） 
// 
// 这里"/tmp/go-buildNNNN/os/exec" 是一个路径，"NNNN"看起来是一个占位符，表示四位数字。在Go语言的上下文中，这很可能是编译时生成的临时文件或目录的路径，与os/exec包相关。
// 
// os/exec是Go标准库中的一个包，用于执行外部程序。而"/tmp/go-buildNNNN"这部分则通常指向Go在构建和运行时生成的临时文件的位置，具体NNNN会由实际的随机数或者进程ID等信息填充。
# <翻译结束>


<原文开始>
// Cat, testing stdin and stdout.
<原文结束>

# <翻译开始>
// Cat，用于测试标准输入和标准输出。
# <翻译结束>


<原文开始>
// Testing combined output and error values.
<原文结束>

# <翻译开始>
// 测试组合输出和错误值。
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
// 测试退出代码是否正确返回
# <翻译结束>


<原文开始>
// Test when command does not call Run().
<原文结束>

# <翻译开始>
// 测试命令未调用 Run() 的情况。
# <翻译结束>


<原文开始>
// Check that we can access methods of the underlying os.File.`
<原文结束>

# <翻译开始>
// 检查我们是否可以访问底层os.File的methods。
# <翻译结束>


<原文开始>
// Before the fix, this next line would race with cmd.Wait.
<原文结束>

# <翻译开始>
// 修复前，下一行代码将与 cmd.Wait 存在竞态条件。
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
// 之前的情况是，当在竞争检测器下运行时，上述的 TestStdinClose 测试会失败。这个测试是 TestStdinClose 的一个变体，同样在竞争检测器下运行时也会失败。
// 此测试由 cmd/dist 在竞争检测器下运行，用于验证竞争检测器不再报告任何问题。
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
// 我们不对Kill的错误返回值进行检查。有可能该进程已经退出，在这种情况下，Kill会返回一个错误“进程已结束”。这个测试的目的在于查看竞态检测器是否会报告错误；Kill是否成功并不重要。
# <翻译结束>


<原文开始>
		// Send the wrong string, so that the child fails even
		// if the other goroutine doesn't manage to kill it first.
		// This test is to check that the race detector does not
		// falsely report an error, so it doesn't matter how the
		// child process fails.
<原文结束>

# <翻译开始>
// 发送错误的字符串，即使其他goroutine未能首先终止子进程，也能使其失败。
// 该测试是为了检查竞态检测器不会错误地报告错误，因此子进程如何失败并不重要。
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
// 由于该测试未并行运行，我们预期在测试执行期间不会打开任何新的文件描述符。但是，如果在测试开始时存在额外的文件描述符（例如由libc打开的），则可能会因某种超时而关闭这些文件描述符。允许这些文件描述符消失，但需检查确保没有添加新的文件描述符。
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
// 这段Go语言注释的翻译如下：
// 
// 本测试的目的在于确保我们打开的所有描述符都被标记为“close-on-exec”。
// 如果haveUnexpectedFDs为真，说明在测试开始时存在其他已打开的描述符，
// 显然这些描述符并未设置为close-on-exec，它们将会使测试结果混乱。尽管我们可以修改测试以期望那些描述符保持打开状态，
// 但由于我们并不清楚它们来自何处或其具体作用，这样做似乎不太稳健。例如，可能是因为某种原因，它们是由当前系统启动代码产生的。
// 另外，这个测试并非针对特定系统；只要大部分系统未跳过该测试，我们仍将测试到我们关心的内容。
# <翻译结束>


<原文开始>
	// This test runs with cgo disabled. External linking needs cgo, so
	// it doesn't work if external linking is required.
<原文结束>

# <翻译开始>
// 这个测试在禁用cgo的情况下运行。外部链接需要cgo，因此如果需要外部链接，则无法正常工作。
# <翻译结束>


<原文开始>
	// Force network usage, to verify the epoll (or whatever) fd
	// doesn't leak to the child,
<原文结束>

# <翻译开始>
// 强制使用网络，以验证 epoll（或其他）文件描述符不会泄漏给子进程
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
// 强制加载TLS根证书（这可能涉及cgo），确保没有任何潜在的C代码泄露文件描述符。
# <翻译结束>


<原文开始>
// quiet expected TLS handshake error "remote error: bad certificate"
<原文结束>

# <翻译开始>
// 静默处理预期的TLS握手错误："远程错误：证书无效"
# <翻译结束>


<原文开始>
	// Build the test without cgo, so that C library functions don't
	// open descriptors unexpectedly. See issue 25628.
<原文结束>

# <翻译开始>
// 在不使用cgo的情况下构建测试，这样C库函数就不会意外地打开描述符。参见问题25628。
# <翻译结束>


<原文开始>
// Use a deadline to try to get some output even if the program hangs.
<原文结束>

# <翻译开始>
// 使用截止时间尝试在程序挂起时仍能获取一些输出。
# <翻译结束>


<原文开始>
		// Leave a 20% grace period to flush output, which may be large on the
		// linux/386 builders because we're running the subprocess under strace.
<原文结束>

# <翻译开始>
// 留出20％的缓冲时间来刷新输出，这在linux/386构建器上可能很大，
// 因为我们正在strace下运行子进程。
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
// 在illumos中，部分功能是通过libc对/proc的访问实现的；这种访问可能会短暂占用一个低位文件描述符(fd)。如果这与检查泄漏描述符的操作并发发生，该检查可能会变得混乱，并错误地报告一个泄漏的描述符。（关于更详细的分析，请参见问题#42431）
//
// 尝试约束子进程中额外线程的使用，以降低该测试的不稳定性：
# <翻译结束>


<原文开始>
// Issue 9173: ignore stdin pipe writes if the program completes successfully.
<原文结束>

# <翻译开始>
// Issue 9173：如果程序成功完成，则忽略标准输入管道的写入。
# <翻译结束>


<原文开始>
	// To reduce noise in the final goroutine dump,
	// let other parallel tests complete if possible.
<原文结束>

# <翻译开始>
// 为了在最终的协程转储中减少噪声，
// 尽可能让其他并行测试完成。
# <翻译结束>


<原文开始>
// At this point the process is alive. Ensure it by sending data to stdin.
<原文结束>

# <翻译开始>
// 此时进程处于运行状态。通过向标准输入发送数据确保这一点。
# <翻译结束>


<原文开始>
	// Calling cancel should have killed the process, so writes
	// should now fail.  Give the process a little while to die.
<原文结束>

# <翻译开始>
// 调用cancel本应结束该进程，因此现在写入应该会失败。给这个进程一点时间结束。
# <翻译结束>


<原文开始>
			// Panic instead of calling t.Fatal so that we get a goroutine dump.
			// We want to know exactly what the os/exec goroutines got stuck on.
<原文结束>

# <翻译开始>
// 通过调用 panic 而不是 t.Fatal，以便获取goroutine的堆栈转储。
// 我们想知道os/exec goroutines 正好卡在了哪里。
# <翻译结束>


<原文开始>
		// Back off exponentially (up to 1-second sleeps) to give the OS time to
		// terminate the process.
<原文结束>

# <翻译开始>
// 以指数级回退（最多1秒的休眠时间），给操作系统时间来终止进程。
# <翻译结束>


<原文开始>
// test that environment variables are de-duped.
<原文结束>

# <翻译开始>
// 测试环境变量是否去重。
# <翻译结束>


<原文开始>
// TestDoubleStartLeavesPipesOpen checks for a regression in which calling
// Start twice, which returns an error on the second call, would spuriously
// close the pipes established in the first call.
<原文结束>

# <翻译开始>
// TestDoubleStartLeavesPipesOpen 检查一个回归问题，即连续两次调用 Start（第二次调用会返回错误）时，会导致在第一次调用中建立的管道被错误地关闭。
# <翻译结束>


<原文开始>
// Release resources if cmd happens not to outlive this process.
<原文结束>

# <翻译开始>
// 如果cmd进程未能在此进程中存活，释放资源。
# <翻译结束>


<原文开始>
// Signal that the process is set up by closing stdout.
<原文结束>

# <翻译开始>
// 通过关闭stdout表示进程已经设置完成。
# <翻译结束>


<原文开始>
// Ignore write errors: we want to keep reading even if stderr is closed.
<原文结束>

# <翻译开始>
// 忽略写入错误：即使标准错误被关闭，我们也希望继续读取。
# <翻译结束>


<原文开始>
// A tickReader reads an unbounded sequence of timestamps at no more than a
// fixed interval.
<原文结束>

# <翻译开始>
// tickReader 以不超过固定间隔的时间读取无界的时间戳序列。
# <翻译结束>


<原文开始>
// Wait for cmd to close stdout to signal that its handlers are installed.
<原文结束>

# <翻译开始>
// 等待cmd关闭stdout，以此来表示它的处理器已安装完毕。
# <翻译结束>


<原文开始>
	// tooLong is an arbitrary duration that is expected to be much longer than
	// the test runs, but short enough that leaked processes will eventually exit
	// on their own.
<原文结束>

# <翻译开始>
// tooLong 是一个任意的持续时间，预计它会远比测试运行时间要长，但又足够短，以至于泄露出去的进程最终会在其自身作用下退出。
# <翻译结束>


<原文开始>
	// Control case: with no cancellation and no WaitDelay, we should wait for the
	// process to exit.
<原文结束>

# <翻译开始>
// 控制用例：如果没有取消操作且没有等待延迟，我们应该等待进程退出。
# <翻译结束>


<原文开始>
	// With a very long WaitDelay and no Cancel function, we should wait for the
	// process to exit even if the command's Context is cancelled.
<原文结束>

# <翻译开始>
// 当WaitDelay时间设置得非常长且没有Cancel函数时，即使命令的上下文被取消，我们也应该等待进程退出。
# <翻译结束>


<原文开始>
		// At this point cmd should still be running (because we passed nil to
		// startHang for the cancel signal). Sending it an explicit Interrupt signal
		// should succeed.
<原文结束>

# <翻译开始>
// 此时，cmd 应该仍在运行（因为我们向 startHang 传递了 nil 作为取消信号）。向其发送一个显式的中断信号应该能够成功。
# <翻译结束>


<原文开始>
		// This program exits with status 0,
		// but pretty much always does so during the wait delay.
		// Since the Cmd itself didn't do anything to stop the process when the
		// context expired, a successful exit is valid (even if late) and does
		// not merit a non-nil error.
<原文结束>

# <翻译开始>
// 此程序以状态0退出，
// 但在等待延迟期间几乎总是如此。
// 由于Cmd本身在上下文过期时没有做任何事情来停止进程，
// 成功退出是有效的（即使稍晚），并且不需要返回非空错误。
# <翻译结束>


<原文开始>
	// If the context is cancelled and the Cancel function sends os.Kill,
	// the process should be terminated immediately, and its output
	// pipes should be closed (causing Wait to return) after WaitDelay
	// even if a child process is still writing to them.
<原文结束>

# <翻译开始>
// 如果上下文被取消，并且Cancel函数发送了os.Kill信号，那么进程应立即终止。即使子进程仍在向其写入数据，其输出管道也应在WaitDelay时间后关闭（导致Wait函数返回）。
# <翻译结束>


<原文开始>
		// This test should kill the child process after 10ms,
		// leaving a grandchild process writing probes in a loop.
		// The child process should be reported as failed,
		// and the grandchild will exit (or die by SIGPIPE) once the
		// stderr pipe is closed.
<原文结束>

# <翻译开始>
// 此测试应在10毫秒后终止子进程，
// 留下孙子进程在一个循环中写入探测数据。
// 子进程应被报告为失败状态，
// 当标准错误管道关闭时，孙子进程将会退出（或因SIGPIPE信号而终止）。
# <翻译结束>


<原文开始>
		// This child process should exit immediately,
		// leaving a grandchild process writing probes in a loop.
		// Since the child has no ExitError to report but we did not
		// read all of its output, Wait should return ErrWaitDelay.
<原文结束>

# <翻译开始>
// 这个子进程应立即退出，
// 留下一个孙子进程在一个循环中写入探测信息。
// 由于子进程没有要报告的ExitError，但我们尚未读取其所有输出，
// 所以Wait函数应返回ErrWaitDelay错误。
# <翻译结束>


<原文开始>
	// If the Cancel function sends a signal that the process can handle, and it
	// handles that signal without actually exiting, then it should be terminated
	// after the WaitDelay.
<原文结束>

# <翻译开始>
// 如果Cancel函数发送了一个进程可以处理的信号，并且该进程在没有真正退出的情况下处理了这个信号，那么应该在WaitDelay之后终止该进程。
# <翻译结束>


<原文开始>
		// This command ignores SIGINT, sleeping until it is killed.
		// Wait should return the usual error for a killed process.
<原文结束>

# <翻译开始>
// 该命令忽略 SIGINT 信号，持续休眠直到被终止。
// 当进程被终止时，Wait 应返回通常的错误信息。
# <翻译结束>


<原文开始>
	// If the process handles the cancellation signal and exits with status 0,
	// Wait should report a non-nil error (because the process had to be
	// interrupted), and it should be a context error (because there is no error
	// to report from the child process itself).
<原文结束>

# <翻译开始>
// 如果进程处理了取消信号并以状态0退出，则Wait应当报告一个非nil错误（因为进程必须被中断），且它应该是一个上下文错误（因为子进程中没有错误需要报告）。
# <翻译结束>


<原文开始>
	// If the Cancel function sends SIGQUIT, it should be handled in the usual
	// way: a Go program should dump its goroutines and exit with non-success
	// status. (We expect SIGQUIT to be a common pattern in real-world use.)
<原文结束>

# <翻译开始>
// 如果Cancel函数发送SIGQUIT信号，那么应当按通常方式处理：Go程序应当dump其goroutine信息并以非成功状态退出。（我们预期SIGQUIT在实际使用中会是一个常见模式。）
# <翻译结束>


<原文开始>
// The default os/signal handler exits with code 2.
<原文结束>

# <翻译开始>
// 默认的 os/signal 处理器将以代码 2 退出。
# <翻译结束>


<原文开始>
	// If Cancel returns a non-ErrProcessDone error and the process
	// exits successfully, Wait should wrap the error from Cancel.
<原文结束>

# <翻译开始>
// 如果Cancel返回的错误不是ErrProcessDone，并且进程成功退出，则Wait应该包装来自Cancel的错误。
# <翻译结束>


<原文开始>
		// We intentionally race Cancel against the process exiting,
		// but ensure that the process wins the race (and return ErrProcessDone
		// from Cancel to report that).
<原文结束>

# <翻译开始>
// 我们特意让 Cancel 与进程退出进行竞争，但确保进程赢得竞争（并在 Cancel 中返回 ErrProcessDone 来报告这一情况）。
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
// 如果Cancel返回错误，并且在WaitDelay之后进程被终止，那么Wait应当报告通常的SIGKILL ExitError，而不是来自Cancel的错误。
# <翻译结束>


<原文开始>
		// Ensure that Cancel actually had the opportunity to
		// return the error.
<原文结束>

# <翻译开始>
// 确保 Cancel 实际上有机会返回错误。
# <翻译结束>


<原文开始>
		// This test should kill the child process after 1ms,
		// To maximize compatibility with existing uses of exec.CommandContext, the
		// resulting error should be an exec.ExitError without additional wrapping.
<原文结束>

# <翻译开始>
// 该测试应在1毫秒后结束子进程，
// 为了最大限度地与现有exec.CommandContext的使用兼容，返回的错误应为一个不带额外封装的exec.ExitError。
# <翻译结束>


<原文开始>
	// If Cancel returns ErrProcessDone but the process is not actually done
	// (and has to be killed), Wait should report the usual SIGKILL ExitError,
	// not the error from Cancel.
<原文结束>

# <翻译开始>
// 如果Cancel返回ErrProcessDone，但实际上进程并未完成（仍需要被杀死），则Wait应当报告通常的SIGKILL ExitError错误，而不是来自Cancel的错误。
# <翻译结束>


<原文开始>
	// If Cancel returns an error and the process exits with an
	// unsuccessful exit code, the process error should take precedence over the
	// Cancel error.
<原文结束>

# <翻译开始>
// 如果Cancel函数返回错误，并且该进程以非成功的退出码结束，则应优先考虑进程错误，而不是Cancel错误。
# <翻译结束>

