
<原文开始>
// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// 版权所有 2017 The Go 作者。保留所有权利。
// 使用本源代码须遵循 BSD 风格
// 许可协议，该协议可在 LICENSE 文件中找到。
# <翻译结束>


<原文开始>
	// On BSD systems regular files seem to be pollable,
	// so just run this test on Linux.
<原文结束>

# <翻译开始>
// 在BSD系统中，似乎普通文件是可被轮询的，因此只在Linux上运行这个测试。
# <翻译结束>


<原文开始>
// noDeadline is a zero time.Time value, which cancels a deadline.
<原文结束>

# <翻译开始>
// noDeadline是一个零值time.Time，用于取消截止时间。
# <翻译结束>


<原文开始>
// expected errors in transition
<原文结束>

# <翻译开始>
// 预期在转换中出现的错误
# <翻译结束>


<原文开始>
	// Tests that read deadlines work, even if there's data ready
	// to be read.
<原文结束>

# <翻译开始>
// 测试读取超时是否起作用，即使已经有数据可以读取。
# <翻译结束>


<原文开始>
// There is a very similar copy of this in net/timeout_test.go.
<原文结束>

# <翻译开始>
// 这里有一个与 net/timeout_test.go 中非常相似的副本。
# <翻译结束>


<原文开始>
// wait for tester goroutine to stop
<原文结束>

# <翻译开始>
// 等待测试者goroutine停止
# <翻译结束>


<原文开始>
	// Tests that write deadlines work, even if there's buffer
	// space available to write.
<原文结束>

# <翻译开始>
// 测试写入截止日期有效，即使有可用的缓冲区空间进行写入。
# <翻译结束>


<原文开始>
	// minDynamicTimeout is the minimum timeout to attempt for
	// tests that automatically increase timeouts until success.
	//
	// Lower values may allow tests to succeed more quickly if the value is close
	// to the true minimum, but may require more iterations (and waste more time
	// and CPU power on failed attempts) if the timeout is too low.
<原文结束>

# <翻译开始>
// minDynamicTimeout是自动增加超时时间直到成功的测试尝试的最小超时时间。
// 
// 如果值接近真实最小值，较小的值可能会让测试更快成功。但如果超时设置得太低，可能需要更多迭代（并且在失败的尝试中浪费更多的时间和CPU资源）。
# <翻译结束>


<原文开始>
	// maxDynamicTimeout is the maximum timeout to attempt for
	// tests that automatically increase timeouts until success.
	//
	// This should be a strict upper bound on the latency required to hit a
	// timeout accurately, even on a slow or heavily-loaded machine. If a test
	// would increase the timeout beyond this value, the test fails.
<原文结束>

# <翻译开始>
// maxDynamicTimeout 是自动增加超时时间直到成功的测试尝试的最大超时时间。
// 
// 这应该是一个严格的上限，即使在运行缓慢或负载过重的机器上，也需要达到准确触发超时所需的延迟。如果测试会将超时时间增加到此值之上，那么测试将失败。
# <翻译结束>


<原文开始>
// timeoutUpperBound returns the maximum time that we expect a timeout of
// duration d to take to return the caller.
<原文结束>

# <翻译开始>
// timeoutUpperBound 函数返回一个预期的最大时长，即对于持续时间为 d 的超时而言，我们期望其能让调用者返回所花费的最长时间。
# <翻译结束>


<原文开始>
		// NetBSD and OpenBSD seem to be unable to reliably hit deadlines even when
		// the absolute durations are long.
		// In https://build.golang.org/log/c34f8685d020b98377dd4988cd38f0c5bd72267e,
		// we observed that an openbsd-amd64-68 builder took 4.090948779s for a
		// 2.983020682s timeout (37.1% overhead).
		// (See https://go.dev/issue/50189 for further detail.)
		// Give them lots of slop to compensate.
<原文结束>

# <翻译开始>
// NetBSD和OpenBSD似乎无法在长时间的绝对期限内可靠地满足期限。
// 在https://build.golang.org/log/c34f8685d020b98377dd4988cd38f0c5bd72267e中，
// 我们观察到一个openbsd-amd64-68构建器在2.983020682s的超时时限中花费了4.090948779s（37.1%的开销）。
// （有关更多详细信息，请参阅https://go.dev/issue/50189。）
// 为了补偿，给他们留出大量的余地。
# <翻译结束>


<原文开始>
	// Other platforms seem to hit their deadlines more reliably,
	// at least when they are long enough to cover scheduling jitter.
<原文结束>

# <翻译开始>
// 其他平台似乎更可靠地满足其截止日期，至少当它们足够长以覆盖调度抖动时。
# <翻译结束>


<原文开始>
// nextTimeout returns the next timeout to try after an operation took the given
// actual duration with a timeout shorter than that duration.
<原文结束>

# <翻译开始>
// nextTimeout 在操作实际耗时超过给定超时时，返回下一个尝试的超时时间。
# <翻译结束>


<原文开始>
	// Since the previous attempt took actual, we can't expect to beat that
	// duration by any significant margin. Try the next attempt with an arbitrary
	// factor above that, so that our growth curve is at least exponential.
<原文结束>

# <翻译开始>
// 由于上一次尝试已耗时actual，我们无法期望在显著程度上超过该耗时。对下一次尝试使用该耗时之上的一个任意因子，以便我们的增长曲线至少为指数级。
# <翻译结束>


<原文开始>
			// Maybe this machine is too slow to reliably schedule goroutines within
			// the requested duration. Increase the timeout and try again.
<原文结束>

# <翻译开始>
// 可能这台机器的性能太慢，无法在请求的时间内可靠地调度goroutines。增加超时时间并重试。
# <翻译结束>


<原文开始>
					// We don't know how long the actual write loop would have taken if
					// the buffer were full, so just guess and double the duration so that
					// the next attempt can make twice as much progress toward filling it.
<原文结束>

# <翻译开始>
// 如果缓冲区已满，我们不知道实际写入循环会花多长时间，所以猜测一下并将其持续时间翻倍，这样下一次尝试就可以向填充缓冲区进更多步。
# <翻译结束>


<原文开始>
// Cannot use t.Parallel - modifies global GOMAXPROCS.
<原文结束>

# <翻译开始>
// 不能使用 t.Parallel - 会修改全局的 GOMAXPROCS。
# <翻译结束>


<原文开始>
		// The writer, with no timeouts of its own,
		// sending bytes to clients as fast as it can.
<原文结束>

# <翻译开始>
// 该写入器自身无任何超时设置，
// 尽可能快速地向客户端发送字节。
# <翻译结束>


<原文开始>
// TestRacyRead tests that it is safe to mutate the input Read buffer
// immediately after cancellation has occurred.
<原文结束>

# <翻译开始>
// TestRacyRead 测试在取消操作发生后立即修改输入读取缓冲区是安全的。
# <翻译结束>


<原文开始>
// Mutate b1 to trigger potential race
<原文结束>

# <翻译开始>
// 修改b1以触发潜在的竞争条件
# <翻译结束>


<原文开始>
// TestRacyWrite tests that it is safe to mutate the input Write buffer
// immediately after cancellation has occurred.
<原文结束>

# <翻译开始>
// TestRacyWrite 测试在取消操作发生后，立即修改输入写缓冲区是安全的。
# <翻译结束>


<原文开始>
// Closing a TTY while reading from it should not hang.  Issue 23943.
<原文结束>

# <翻译开始>
// 在从TTY读取数据时关闭它，不应导致程序挂起。问题 23943。
# <翻译结束>


<原文开始>
// Ignore SIGTTIN in case we are running in the background.
<原文结束>

# <翻译开始>
// 如果我们在后台运行，忽略SIGTTIN信号。
# <翻译结束>


<原文开始>
	// Give the goroutine a chance to enter the read.
	// It doesn't matter much if it occasionally fails to do so,
	// we won't be testing what we want to test but the test will pass.
<原文结束>

# <翻译开始>
// 让goroutine有机会进入读取阶段。
// 如果偶尔未能做到这一点没有太大关系，
// 我们不会测试到我们想要的，但测试会通过。
# <翻译结束>


<原文开始>
	// On some systems the goroutines may now be hanging.
	// There's not much we can do about that.
<原文结束>

# <翻译开始>
// 在某些系统中，goroutine 可能会挂起。
// 对此我们无能为力。
# <翻译结束>

