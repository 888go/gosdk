
<原文开始>
// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// 版权所有 ? 2015 The Go 作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。
# <翻译结束>


<原文开始>
	// This test cannot be run in parallel because of a race similar
	// to the one reported in https://go.dev/issue/22315.
	//
	// Even though the pipe is opened with O_CLOEXEC, if another test forks in
	// between the call to os.Pipe and the call to r.Close, that child process can
	// retain an open copy of r's file descriptor until it execs. If one of our
	// Write calls occurs during that interval it can spuriously succeed,
	// buffering the write to the child's copy of the pipe (even though the child
	// will not actually read the buffered bytes).
<原文结束>

# <翻译开始>
	// 由于存在与在 https:	//go.dev/issue/22315 中报告的类似竞争条件，该测试不能并行运行。
	// 
	// 尽管管道以 O_CLOEXEC 标志打开，但如果在调用 os.Pipe 和调用 r.Close 之间有其他测试执行了 fork 操作，那么子进程可能会在其 exec 前保留对 r 文件描述符的打开副本。若我们的某个 Write 调用恰好在此期间发生，它可能会意外成功，将写入操作缓冲到子进程持有的管道副本中（尽管子进程实际上并不会读取这些已缓冲的字节）。
# <翻译结束>


<原文开始>
// 232 is Windows error code ERROR_NO_DATA, "The pipe is being closed".
<原文结束>

# <翻译开始>
// 232 是Windows错误代码ERROR_NO_DATA，表示"管道正在关闭"。
# <翻译结束>


<原文开始>
// Every time we write to the pipe we should get an EPIPE.
<原文结束>

# <翻译开始>
// 每次向管道写入时，我们都应该收到一个 EPIPE。
# <翻译结束>


<原文开始>
		// For stdout/stderr, we should have crashed with a broken pipe error.
		// The caller will be looking for that exit status,
		// so just exit normally here to cause a failure in the caller.
		// For descriptor 3, a normal exit is expected.
<原文结束>

# <翻译开始>
		// 对于stdout/stderr，我们应该因管道破裂错误而崩溃。
		// 调用者会期待那个退出状态，
		// 所以在这里正常退出即可在调用者那里引发失败。
		// 对于描述符3，期望一个正常的退出。
# <翻译结束>


<原文开始>
	// This test cannot be run in parallel due to the same race as for TestEPIPE.
	// (We expect a write to a closed pipe can fail, but a concurrent fork of a
	// child process can cause the pipe to unexpectedly remain open.)
<原文结束>

# <翻译开始>
	// 由于与TestEPIPE相同的竞态条件，该测试无法并行运行。
	// （我们预期向已关闭的管道写入可能会失败，但并发创建子进程可能导致管道意外保持打开状态。）
# <翻译结束>


<原文开始>
	// Invoke the test program to run the test and write to a closed pipe.
	// If sig is false:
	// writing to stdout or stderr should cause an immediate SIGPIPE;
	// writing to descriptor 3 should fail with EPIPE and then exit 0.
	// If sig is true:
	// all writes should fail with EPIPE and then exit 0.
<原文结束>

# <翻译开始>
	// 调用测试程序来运行测试，并向一个已关闭的管道中写入内容。
	// 如果sig为false：
	// 向stdout或stderr写入应立即导致SIGPIPE信号；
	// 向描述符3写入应失败并返回EPIPE，然后退出状态为0。
	// 如果sig为true：
	// 所有写入操作都应因EPIPE失败，然后退出状态为0。
# <翻译结束>


<原文开始>
// Test redirecting stdout but not stderr.  Issue 40076.
<原文结束>

# <翻译开始>
// 测试重定向stdout但不stderr。问题40076。
# <翻译结束>


<原文开始>
		// Get the amount we have to write to overload a pipe
		// with no reader.
<原文结束>

# <翻译开始>
		// 获取需要写入以使无读取器的管道超载的数量。
# <翻译结束>


<原文开始>
	// Close the read end of the pipe in a goroutine while we are
	// writing to the write end, or vice-versa.
<原文结束>

# <翻译开始>
	// 在我们向管道的写入端写入数据的同时，或反之，在一个goroutine中关闭管道的读取端。
# <翻译结束>


<原文开始>
		// Give the main goroutine a chance to enter the Read or
		// Write call. This is sloppy but the test will pass even
		// if we close before the read/write.
<原文结束>

# <翻译开始>
		// 给主 Goroutine 一个机会进入 Read 或 Write 调用。
		// 这里写得比较粗糙，但即使在读/写之前关闭，测试仍然会通过。
# <翻译结束>


<原文开始>
// Issue 20915: Reading on nonblocking fd should not return "waiting
// for unsupported file type." Currently it returns EAGAIN; it is
// possible that in the future it will simply wait for data.
<原文结束>

# <翻译开始>
// 问题20915：在非阻塞fd上读取时，不应该返回“等待不支持的文件类型”。目前它返回EAGAIN；将来有可能会简单地等待数据。
# <翻译结束>


<原文开始>
// os.NewFile returns a blocking mode file.
<原文结束>

# <翻译开始>
// os.NewFile返回一个阻塞模式的文件。
# <翻译结束>


<原文开始>
// Calling Fd will put the file into blocking mode.
<原文结束>

# <翻译开始>
// 调用Fd将使文件进入阻塞模式。
# <翻译结束>


<原文开始>
// Test that we don't let a blocking read prevent a close.
<原文结束>

# <翻译开始>
// 测试确保不会让阻塞读取阻止关闭操作。
# <翻译结束>


<原文开始>
	// Give the goroutine a chance to enter the Read
	// or Write call. This is sloppy but the test will
	// pass even if we close before the read/write.
<原文结束>

# <翻译开始>
	// 让goroutine有机会进入Read或Write调用。这很粗糙，但即使我们在读写之前关闭它，测试也会通过。
# <翻译结束>


<原文开始>
	// r.Close has completed, but since we assume r is in blocking mode that
	// probably didn't unblock the call to r.Read. Close w to unblock it.
<原文结束>

# <翻译开始>
	// r.Close 已经完成，但由于我们假设 r 在阻塞模式下，这可能并没有解阻塞对 r.Read 的调用。关闭 w 来解阻塞它。
# <翻译结束>


<原文开始>
// testPipeEOF tests that when the write side of a pipe or FIFO is closed,
// a blocked Read call on the reader side returns io.EOF.
//
// This scenario previously failed to unblock the Read call on darwin.
// (See https://go.dev/issue/24164.)
<原文结束>

# <翻译开始>
// testPipeEOF 测试当管道或FIFO的写入端被关闭时，
// 读取端阻塞的Read调用会返回io.EOF。
//
// 此场景之前在darwin上未能使Read调用解除阻塞。
// （参见 https://go.dev/issue/24164。）
# <翻译结束>


<原文开始>
	// parkDelay is an arbitrary delay we wait for a pipe-reader goroutine to park
	// before issuing the corresponding write. The test should pass no matter what
	// delay we use, but with a longer delay is has a higher chance of detecting
	// poller bugs.
<原文结束>

# <翻译开始>
	// parkDelay 是一个任意的延迟，我们等待管道读取器goroutine暂停
	// 在执行相应的写入之前。测试应该无论使用什么延迟都能通过，
	// 但使用更长的延迟更有可能检测到轮询器的bug。
# <翻译结束>


<原文开始>
	// This test starts 100 simultaneous goroutines, which could bury a more
	// interesting stack if this or some other test happens to panic. It is also
	// nearly instantaneous, so any latency benefit from running it in parallel
	// would be minimal.
<原文结束>

# <翻译开始>
	// 这个测试启动了100个并发的goroutine，如果这个或其他测试发生恐慌，可能会隐藏更有趣的堆栈。同时，它几乎是瞬间完成的，所以并行运行时的任何延迟好处都会非常小。
# <翻译结束>


<原文开始>
		// Give the other goroutine a chance to enter the Read.
		// It doesn't matter if this occasionally fails, the test
		// will still pass, it just won't test anything.
<原文结束>

# <翻译开始>
		// 让其他goroutine有机会进入Read。
		// 如果这个操作偶尔失败也没关系，测试仍然会通过，只是不会检验任何内容。
# <翻译结束>


<原文开始>
		// The bug was that Fd would hang until Read timed out.
		// If the bug is fixed, then writing to w and closing r here
		// will cause the Read to exit before the timeout expires.
<原文结束>

# <翻译开始>
		// 该bug表现为Fd会一直阻塞直到Read超时。
		// 若该bug已被修复，那么在此处向w写入数据并关闭r
		// 将导致Read在超时之前退出。
# <翻译结束>

