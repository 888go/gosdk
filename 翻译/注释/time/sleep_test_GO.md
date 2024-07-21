
<原文开始>
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
//版权所有2009年Go作者。所有权利保留。
//使用此源代码受BSD风格
//可以在LICENSE文件中找到的许可证。
// md5:2e9dc81828a3be8a
# <翻译结束>


<原文开始>
// Go runtime uses different Windows timers for time.Now and sleeping.
// These can tick at different frequencies and can arrive out of sync.
// The effect can be seen, for example, as time.Sleep(100ms) is actually
// shorter then 100ms when measured as difference between time.Now before and
// after time.Sleep call. This was observed on Windows XP SP3 (windows/386).
// windowsInaccuracy is to ignore such errors.
<原文结束>

# <翻译开始>
// Go运行时使用不同的Windows定时器来获取time.Now和睡眠时间。
// 这些定时器可能以不同的频率工作，并且可能会不同步。
// 这种影响可以通过观察到，例如，在Windows XP SP3（windows/386）上，调用time.Sleep(100ms)实际上比100ms短，
// 在time.Sleep调用前后的时间.Now之间的差异中可以看到。windowsInaccuracy是为了忽略此类错误。
// md5:413315af6e6e879d
# <翻译结束>


<原文开始>
// Test the basic function calling behavior. Correct queueing
// behavior is tested elsewhere, since After and AfterFunc share
// the same code.
<原文结束>

# <翻译开始>
// 测试基本的函数调用行为。由于`After`和`AfterFunc`共享相同的代码，所以在这里只测试队列行为的基本正确性。
// md5:ea0a035c0fadf390
# <翻译结束>


<原文开始>
			// Yield so that the OS can wake up the timer thread,
			// so that it can generate channel sends for the main goroutine,
			// which will eventually set stop = 1 for us.
<原文结束>

# <翻译开始>
			// 让出执行权以便操作系统唤醒计时器线程，
			// 这样计时器线程就可以为主goroutine生成通道发送操作，
			// 而主goroutine最终会为我们设置stop = 1。
			// md5:ccb50fd753147a7c
# <翻译结束>


<原文开始>
	// Create equal number of garbage timers on each P before starting
	// the benchmark.
<原文结束>

# <翻译开始>
	// 在开始基准测试之前，在每个P上创建数量相等的垃圾回收定时器。
	// md5:a87f583366d1d2f7
# <翻译结束>


<原文开始>
	// We want to test that we stop a timer before it runs.
	// We also want to test that it didn't run after a longer timer.
	// Since we don't want the test to run for too long, we don't
	// want to use lengthy times. That makes the test inherently flaky.
	// So only report an error if it fails five times in a row.
<原文结束>

# <翻译开始>
	// 我们想测试在计时器运行前是否能停止它。
	// 同时，我们也想测试在较长的计时器后它是否未运行。
	// 由于我们不想让测试运行太久，所以我们不会使用过长的时间。这使得测试本质上具有一定的不稳定性。
	// 因此，只有当连续五次失败时才报告错误。
	// md5:69b5ca4d4ffdeb2e
# <翻译结束>


<原文开始>
// Test passed, so all done.
<原文结束>

# <翻译开始>
// 测试通过，所以大功告成。. md5:898d4f57bd6d9649
# <翻译结束>


<原文开始>
	// This test flakes out on some systems,
	// so we'll try it a few times before declaring it a failure.
<原文结束>

# <翻译开始>
	// 此测试在某些系统上可能出现不稳定，
	// 因此我们在认定失败之前会尝试运行几次。
	// md5:1c0a70fc679d8303
# <翻译结束>


<原文开始>
	// make the result channel buffered because we don't want
	// to depend on channel queueing semantics that might
	// possibly change in the future.
<原文结束>

# <翻译开始>
	// 将结果通道设置为缓冲通道，因为我们不希望依赖于
	// 可能在未来发生改变的通道排队语义。
	// md5:14c7b165b90f6101
# <翻译结束>


<原文开始>
	// Sleep(0) used to hang, the sequence of events was as follows.
	// Sleep(0) sets G's status to Gwaiting, but then immediately returns leaving the status.
	// Then the goroutine calls e.g. new and falls down into the scheduler due to pending GC.
	// After the GC nobody wakes up the goroutine from Gwaiting status.
<原文结束>

# <翻译开始>
	// 使用Sleep(0)时，它会导致挂起，事件顺序如下：
	// Sleep(0) 将 G 的状态设置为 Gwaiting，但随后立即返回并留下该状态。
	// 然后，goroutine 调用例如 new，并由于待处理的垃圾回收（GC）而陷入调度器。
	// 在 GC 之后，没有人会将处于 Gwaiting 状态的 goroutine 唤醒。
	// md5:9ea4995b95112ee8
# <翻译结束>


<原文开始>
	// We try to run this test with increasingly larger multiples
	// until one works so slow, loaded hardware isn't as flaky,
	// but without slowing down fast machines unnecessarily.
	//
	// (maxDuration is several orders of magnitude longer than we
	// expect this test to actually take on a fast, unloaded machine.)
<原文结束>

# <翻译开始>
	// 我们尝试使用越来越大的倍数运行此测试，
	// 直到其中一个运行得如此之慢，负载硬件就不那么不可预测了，
	// 但又不会无谓地减慢快机器的速度。
	// 
	// (maxDuration远远超过了我们期望在空载快速机器上实际运行此测试所需的时间。)
	// md5:0d4cdfa6cf3c4e19
# <翻译结束>


<原文开始>
// Test that sleeping (via Sleep or Timer) for an interval so large it
// overflows does not result in a short sleep duration. Nor does it interfere
// with execution of other timers. If it does, timers in this or subsequent
// tests may not fire.
<原文结束>

# <翻译开始>
// 测试通过Sleep或Timer方式睡眠一个极大以至于溢出的时间间隔，不会导致实际睡眠时间变短。
// 同时确保这不会干扰其他计时器的执行。如果发生干扰，本测试或后续测试中的计时器可能无法正常触发。
// md5:886fd5d71b9b6fc0
# <翻译结束>


<原文开始>
		// On failure, this may return after the test has completed, so
		// we need to panic instead.
<原文结束>

# <翻译开始>
		// 在失败时，此函数可能会在测试完成后返回，
		// 因此我们需要引发恐慌。
		// md5:3782d52fc8354241
# <翻译结束>


<原文开始>
// Test that a panic while deleting a timer does not leave
// the timers mutex held, deadlocking a ticker.Stop in a defer.
<原文结束>

# <翻译开始>
// 测试在删除定时器时发生恐慌，不会导致持有定时器的互斥锁，从而在defer中死锁ticker.Stop。
// md5:4a5f0fbc22171d44
# <翻译结束>


<原文开始>
		// would deadlock here before the fix due to
		// lock taken before the segfault.
<原文结束>

# <翻译开始>
		// 在修复之前，这里会因为锁在 segmentation fault 之前被获取而死锁。
		// md5:4095774a3c34b926
# <翻译结束>


<原文开始>
// cause a panic due to a segfault
<原文结束>

# <翻译开始>
// 导致由于段错误而引发恐慌. md5:9e2d522d6882179b
# <翻译结束>


<原文开始>
	// This may hang forever if timers are broken. See comment near
	// the end of CheckRuntimeTimerOverflow in internal_test.go.
<原文结束>

# <翻译开始>
	// 如果定时器出现问题，此操作可能会永久阻塞。请参阅内部_test.go文件中CheckRuntimeTimerOverflow函数结尾处的注释。
	// md5:049f552673ac0729
# <翻译结束>


<原文开始>
// Test that zero duration timers aren't missed by the scheduler. Regression test for issue 44868.
<原文结束>

# <翻译开始>
// 测试零时长的定时器不会被调度程序错过。针对问题44868的回归测试。. md5:d08c608639513c32
# <翻译结束>


<原文开始>
// Test that rapidly moving a timer earlier doesn't cause it to get dropped.
// Issue 47329.
<原文结束>

# <翻译开始>
// 测试快速将定时器向前移动不会导致它被丢弃。问题47329。
// md5:5b7b156696041586
# <翻译结束>


<原文开始>
// Test that rapidly moving timers earlier and later doesn't cause
// some of the sleep times to be lost.
// Issue 47762
<原文结束>

# <翻译开始>
// 测试快速地将计时器提前和延后不会导致
// 某些休眠时间丢失。
// 问题 47762
// md5:e2d773f7ca59cb7e
# <翻译结束>


<原文开始>
// Reset to various long durations, which we'll cancel.
<原文结束>

# <翻译开始>
// 重置为多个长持续时间，我们将会取消这些操作。. md5:4def5d15f369f372
# <翻译结束>


<原文开始>
// Stop and drain a long-duration timer.
<原文结束>

# <翻译开始>
// 停止并释放一个长时间运行的计时器。. md5:1a645fced10a9490
# <翻译结束>


<原文开始>
// Start a short-duration timer we expect to select without blocking.
<原文结束>

# <翻译开始>
// 启动一个我们期望不会阻塞的短时定时器。. md5:07ecf94e05725a83
# <翻译结束>


<原文开始>
// Timer is done. Swap with tail and remove.
<原文结束>

# <翻译开始>
// 计时器已完成。与尾部交换并移除。. md5:d6ea2013f90a3935
# <翻译结束>


<原文开始>
// Benchmark timer latency when the thread that creates the timer is busy with
// other work and the timers must be serviced by other threads.
// https://golang.org/issue/38860
<原文结束>

# <翻译开始>
// 测试当创建定时器的线程忙于其他工作时，定时器由其他线程服务时的延迟时间。
// 参考：https://golang.org/issue/38860
// md5:ebfd934f369f8698
# <翻译结束>


<原文开始>
// allocate memory now to avoid GC interference later.
<原文结束>

# <翻译开始>
// 现在分配内存以避免稍后出现垃圾收集干扰。. md5:06108e8f341b6a59
# <翻译结束>


<原文开始>
	// Ensure the time to start new threads to service timers will not pollute
	// the results.
<原文结束>

# <翻译开始>
	// 确保开始新线程服务定时器的时间不会影响结果。
	// md5:67235cb4b6190520
# <翻译结束>


<原文开始>
	// Note that other than the AfterFunc calls this benchmark is measuring it
	// avoids using any other timers. In particular, the main goroutine uses
	// doWork to spin for some durations because up through Go 1.15 if all
	// threads are idle sysmon could leave deep sleep when we wake.
<原文结束>

# <翻译开始>
	// 请注意，除了正在测量的AfterFunc调用之外，这个基准测试避免了使用任何其他定时器。具体来说，主goroutine使用doWork来为某些持续时间进行自旋，因为在Go 1.15及之前版本中，如果所有线程都处于空闲状态，当唤醒时sysmon可能会离开深度睡眠状态。
	// md5:48be121202ffd149
# <翻译结束>


<原文开始>
// Ensure sysmon is in deep sleep.
<原文结束>

# <翻译开始>
// 确保sysmon处于深度睡眠状态。. md5:dd1381290960838a
# <翻译结束>


<原文开始>
// spin until all timers fired
<原文结束>

# <翻译开始>
// 一直等待，直到所有计时器触发. md5:633b8d9d10dc0605
# <翻译结束>


<原文开始>
		// Spin for a bit to let the other scheduler threads go idle before the
		// next round.
<原文结束>

# <翻译开始>
		// 稍微等待一下，让其他调度线程空闲，然后再进入下一轮。
		// md5:561da4b1365595cd
# <翻译结束>


<原文开始>
// Benchmark timer latency with staggered wakeup times and varying CPU bound
// workloads. https://golang.org/issue/38860
<原文结束>

# <翻译开始>
// 使用交错的唤醒时间和不同的CPU密集型工作负载来基准测试计时器延迟。
// https://golang.org/issue/38860
// md5:11caeca6cee3a8ca
# <翻译结束>


<原文开始>
					// Ensure the time to start new threads to service timers
					// will not pollute the results.
<原文结束>

# <翻译开始>
					// 确保启动新线程来处理计时器的时间不会污染结果。
					// md5:41d02fd2d4916706
# <翻译结束>


<原文开始>
// warmupScheduler ensures the scheduler has at least targetThreadCount threads
// in its thread pool.
<原文结束>

# <翻译开始>
// warmupScheduler 确保调度器的线程池至少有targetThreadCount个线程。
// md5:bc91fc80f58bb498
# <翻译结束>


<原文开始>
// spin until all threads started
<原文结束>

# <翻译开始>
// 一直等待，直到所有线程启动. md5:ecc162218f853d96
# <翻译结束>


<原文开始>
// spin a bit more to ensure they are all running on separate CPUs.
<原文结束>

# <翻译开始>
// 再等待一段时间，以确保它们都在独立的CPU上运行。. md5:bc7d4fee547b09a1
# <翻译结束>

