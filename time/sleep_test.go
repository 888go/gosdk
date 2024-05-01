//版权所有2009年Go作者。所有权利保留。
//使用此源代码受BSD风格
//可以在LICENSE文件中找到的许可证。
// md5:2e9dc81828a3be8a

package time_test

import (
	"errors"
	"fmt"
	. "github.com/888go/gosdk/time"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
)

// Go运行时使用不同的Windows定时器来获取time.Now和睡眠时间。
// 这些定时器可能以不同的频率工作，并且可能会不同步。
// 这种影响可以通过观察到，例如，在Windows XP SP3（windows/386）上，调用time.Sleep(100ms)实际上比100ms短，
// 在time.Sleep调用前后的时间.Now之间的差异中可以看到。windowsInaccuracy是为了忽略此类错误。
// md5:413315af6e6e879d
const windowsInaccuracy = 17 * Millisecond

//
//func TestSleep(t *testing.T) {
//	const delay = 100 * Millisecond
//	go func() {
//		Sleep(delay / 2)
//		Interrupt()
//	}()
//	start := Now()
//	Sleep(delay)
//	delayadj := delay
//	if runtime.GOOS == "windows" {
//		delayadj -= windowsInaccuracy
//	}
//	duration := Now().Sub(start)
//	if duration < delayadj {
//		t.Fatalf("Sleep(%s) slept for only %s", delay, duration)
//	}
//}

// 测试基本的函数调用行为。由于`After`和`AfterFunc`共享相同的代码，所以在这里只测试队列行为的基本正确性。
// md5:ea0a035c0fadf390
func TestAfterFunc(t *testing.T) {
	i := 10
	c := make(chan bool)
	var f func()
	f = func() {
		i--
		if i >= 0 {
			AfterFunc(0, f)
			Sleep(1 * Second)
		} else {
			c <- true
		}
	}

	AfterFunc(0, f)
	<-c
}

func TestAfterStress(t *testing.T) {
	var stop atomic.Bool
	go func() {
		for !stop.Load() {
			runtime.GC()
			// 让出执行权以便操作系统唤醒计时器线程，
			// 这样计时器线程就可以为主goroutine生成通道发送操作，
			// 而主goroutine最终会为我们设置stop = 1。
			// md5:ccb50fd753147a7c
			Sleep(Nanosecond)
		}
	}()
	ticker := NewTicker(1)
	for i := 0; i < 100; i++ {
		<-ticker.F.C
	}
	ticker.Stop()
	stop.Store(true)
}

func benchmark(b *testing.B, bench func(n int)) {

	// 在开始基准测试之前，在每个P上创建数量相等的垃圾回收定时器。
	// md5:a87f583366d1d2f7
	var wg sync.WaitGroup
	garbageAll := make([][]*Timer, runtime.GOMAXPROCS(0))
	for i := range garbageAll {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			garbage := make([]*Timer, 1<<15)
			for j := range garbage {
				garbage[j] = AfterFunc(Hour, nil)
			}
			garbageAll[i] = garbage
		}(i)
	}
	wg.Wait()

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			bench(1000)
		}
	})
	b.StopTimer()

	for _, garbage := range garbageAll {
		for _, t := range garbage {
			t.Stop()
		}
	}
}

func BenchmarkAfterFunc(b *testing.B) {
	benchmark(b, func(n int) {
		c := make(chan bool)
		var f func()
		f = func() {
			n--
			if n >= 0 {
				AfterFunc(0, f)
			} else {
				c <- true
			}
		}

		AfterFunc(0, f)
		<-c
	})
}

func BenchmarkAfter(b *testing.B) {
	benchmark(b, func(n int) {
		for i := 0; i < n; i++ {
			<-After(1)
		}
	})
}

func BenchmarkStop(b *testing.B) {
	benchmark(b, func(n int) {
		for i := 0; i < n; i++ {
			NewTimer(1 * Second).Stop()
		}
	})
}

func BenchmarkSimultaneousAfterFunc(b *testing.B) {
	benchmark(b, func(n int) {
		var wg sync.WaitGroup
		wg.Add(n)
		for i := 0; i < n; i++ {
			AfterFunc(0, wg.Done)
		}
		wg.Wait()
	})
}

func BenchmarkStartStop(b *testing.B) {
	benchmark(b, func(n int) {
		timers := make([]*Timer, n)
		for i := 0; i < n; i++ {
			timers[i] = AfterFunc(Hour, nil)
		}

		for i := 0; i < n; i++ {
			timers[i].Stop()
		}
	})
}

func BenchmarkReset(b *testing.B) {
	benchmark(b, func(n int) {
		t := NewTimer(Hour)
		for i := 0; i < n; i++ {
			t.Reset(Hour)
		}
		t.Stop()
	})
}

func BenchmarkSleep(b *testing.B) {
	benchmark(b, func(n int) {
		var wg sync.WaitGroup
		wg.Add(n)
		for i := 0; i < n; i++ {
			go func() {
				Sleep(Nanosecond)
				wg.Done()
			}()
		}
		wg.Wait()
	})
}

func TestAfter(t *testing.T) {
	const delay = 100 * Millisecond
	start := Now()
	end := <-After(delay)
	delayadj := delay
	if runtime.GOOS == "windows" {
		delayadj -= windowsInaccuracy
	}
	if duration := Now().Sub(start); duration < delayadj {
		t.Fatalf("After(%s) slept for only %d ns", delay, duration)
	}
	if min := start.Add(delayadj); end.Before(min.F) {
		t.Fatalf("After(%s) expect >= %s, got %s", delay, min, end)
	}
}

func TestAfterTick(t *testing.T) {
	const Count = 10
	Delta := 100 * Millisecond
	if testing.Short() {
		Delta = 10 * Millisecond
	}
	t0 := Now()
	for i := 0; i < Count; i++ {
		<-After(Delta)
	}
	t1 := Now()
	d := t1.Sub(t0)
	target := Delta * Count
	if d < target*9/10 {
		t.Fatalf("%d ticks of %s too fast: took %s, expected %s", Count, Delta, d, target)
	}
	if !testing.Short() && d > target*30/10 {
		t.Fatalf("%d ticks of %s too slow: took %s, expected %s", Count, Delta, d, target)
	}
}

func TestAfterStop(t *testing.T) {
	// 我们想测试在计时器运行前是否能停止它。
	// 同时，我们也想测试在较长的计时器后它是否未运行。
	// 由于我们不想让测试运行太久，所以我们不会使用过长的时间。这使得测试本质上具有一定的不稳定性。
	// 因此，只有当连续五次失败时才报告错误。
	// md5:69b5ca4d4ffdeb2e

	var errs []string
	logErrs := func() {
		for _, e := range errs {
			t.Log(e)
		}
	}

	for i := 0; i < 5; i++ {
		AfterFunc(100*Millisecond, func() {})
		t0 := NewTimer(50 * Millisecond)
		c1 := make(chan bool, 1)
		t1 := AfterFunc(150*Millisecond, func() { c1 <- true })
		c2 := After(200 * Millisecond)
		if !t0.Stop() {
			errs = append(errs, "failed to stop event 0")
			continue
		}
		if !t1.Stop() {
			errs = append(errs, "failed to stop event 1")
			continue
		}
		<-c2
		select {
		case <-t0.F.C:
			errs = append(errs, "event 0 was not stopped")
			continue
		case <-c1:
			errs = append(errs, "event 1 was not stopped")
			continue
		default:
		}
		if t1.Stop() {
			errs = append(errs, "Stop returned true twice")
			continue
		}

		// 测试通过，所以大功告成。. md5:898d4f57bd6d9649
		if len(errs) > 0 {
			t.Logf("saw %d errors, ignoring to avoid flakiness", len(errs))
			logErrs()
		}

		return
	}

	t.Errorf("saw %d errors", len(errs))
	logErrs()
}

//func TestAfterQueuing(t *testing.T) {
//	// 此测试在某些系统上可能出现不稳定，
//	// 因此我们在认定失败之前会尝试运行几次。
//	// md5:1c0a70fc679d8303
//	const attempts = 5
//	err := errors.New("!=nil")
//	for i := 0; i < attempts && err != nil; i++ {
//		delta := Duration(20+i*50) * Millisecond
//		if err = testAfterQueuing(delta); err != nil {
//			t.Logf("attempt %v failed: %v", i, err)
//		}
//	}
//	if err != nil {
//		t.Fatal(err)
//	}
//}

var slots = []int{5, 3, 6, 6, 6, 1, 1, 2, 7, 9, 4, 8, 0}

type afterResult struct {
	slot int
	t    Time
}

func await(slot int, result chan<- afterResult, ac <-chan Time) {
	result <- afterResult{slot, <-ac}
}

//
//func testAfterQueuing(delta Duration) error {
//	// 将结果通道设置为缓冲通道，因为我们不希望依赖于
//	// 可能在未来发生改变的通道排队语义。
//	// md5:14c7b165b90f6101
//	result := make(chan afterResult, len(slots))
//
//	t0 := Now()
//	for _, slot := range slots {
//		go await(slot, result, After(Duration(slot)*delta))
//	}
//	var order []int
//	var times []Time
//	for range slots {
//		r := <-result
//		order = append(order, r.slot)
//		times = append(times, r.t)
//	}
//	for i := range order {
//		if i > 0 && order[i] < order[i-1] {
//			return fmt.Errorf("After calls returned out of order: %v", order)
//		}
//	}
//	for i, t := range times {
//		dt := t.Sub(t0)
//		target := Duration(order[i]) * delta
//		if dt < target-delta/2 || dt > target+delta*10 {
//			return fmt.Errorf("After(%s) arrived at %s, expected [%s,%s]", target, dt, target-delta/2, target+delta*10)
//		}
//	}
//	return nil
//}

func TestTimerStopStress(t *testing.T) {
	if testing.Short() {
		return
	}
	for i := 0; i < 100; i++ {
		go func(i int) {
			timer := AfterFunc(2*Second, func() {
				t.Errorf("timer %d was not stopped", i)
			})
			Sleep(1 * Second)
			timer.Stop()
		}(i)
	}
	Sleep(3 * Second)
}

func TestSleepZeroDeadlock(t *testing.T) {
	// 使用Sleep(0)时，它会导致挂起，事件顺序如下：
	// Sleep(0) 将 G 的状态设置为 Gwaiting，但随后立即返回并留下该状态。
	// 然后，goroutine 调用例如 new，并由于待处理的垃圾回收（GC）而陷入调度器。
	// 在 GC 之后，没有人会将处于 Gwaiting 状态的 goroutine 唤醒。
	// md5:9ea4995b95112ee8
	defer runtime.GOMAXPROCS(runtime.GOMAXPROCS(4))
	c := make(chan bool)
	go func() {
		for i := 0; i < 100; i++ {
			runtime.GC()
		}
		c <- true
	}()
	for i := 0; i < 100; i++ {
		Sleep(0)
		tmp := make(chan bool, 1)
		tmp <- true
		<-tmp
	}
	<-c
}

func testReset(d Duration) error {
	t0 := NewTimer(2 * d)
	Sleep(d)
	if !t0.Reset(3 * d) {
		return errors.New("resetting unfired timer returned false")
	}
	Sleep(2 * d)
	select {
	case <-t0.F.C:
		return errors.New("timer fired early")
	default:
	}
	Sleep(2 * d)
	select {
	case <-t0.F.C:
	default:
		return errors.New("reset timer did not fire")
	}

	if t0.Reset(50 * Millisecond) {
		return errors.New("resetting expired timer returned true")
	}
	return nil
}

//func TestReset(t *testing.T) {
//	// 我们尝试使用越来越大的倍数运行此测试，
//	// 直到其中一个运行得如此之慢，负载硬件就不那么不可预测了，
//	// 但又不会无谓地减慢快机器的速度。
//	//
//	// (maxDuration远远超过了我们期望在空载快速机器上实际运行此测试所需的时间。)
//	// md5:0d4cdfa6cf3c4e19
//	d := 1 * Millisecond
//	const maxDuration = 10 * Second
//	for {
//		err := testReset(d)
//		if err == nil {
//			break
//		}
//		d *= 2
//		if d > maxDuration {
//			t.Error(err)
//		}
//		t.Logf("%v; trying duration %v", err, d)
//	}
//}

// 测试通过Sleep或Timer方式睡眠一个极大以至于溢出的时间间隔，不会导致实际睡眠时间变短。
// 同时确保这不会干扰其他计时器的执行。如果发生干扰，本测试或后续测试中的计时器可能无法正常触发。
// md5:886fd5d71b9b6fc0
func TestOverflowSleep(t *testing.T) {
	const big = Duration(int64(1<<63 - 1))

	go func() {
		Sleep(big)
		// 在失败时，此函数可能会在测试完成后返回，
		// 因此我们需要引发恐慌。
		// md5:3782d52fc8354241
		panic("big sleep returned")
	}()

	select {
	case <-After(big):
		t.Fatalf("big timeout fired")
	case <-After(25 * Millisecond):
		// OK
	}

	const neg = Duration(-1 << 63)
	Sleep(neg) // Returns immediately.
	select {
	case <-After(neg):
		// OK
	case <-After(1 * Second):
		t.Fatalf("negative timeout didn't fire")
	}
}

// 测试在删除定时器时发生恐慌，不会导致持有定时器的互斥锁，从而在defer中死锁ticker.Stop。
// md5:4a5f0fbc22171d44
func TestIssue5745(t *testing.T) {
	ticker := NewTicker(Hour)
	defer func() {
		// 在修复之前，这里会因为锁在 segmentation fault 之前被获取而死锁。
		// md5:4095774a3c34b926
		ticker.Stop()

		if r := recover(); r == nil {
			t.Error("Expected panic, but none happened.")
		}
	}()

	// 导致由于段错误而引发恐慌. md5:9e2d522d6882179b
	var timer *Timer
	timer.Stop()
	t.Error("Should be unreachable.")
}

//func TestOverflowPeriodRuntimeTimer(t *testing.T) {
//	// 如果定时器出现问题，此操作可能会永久阻塞。请参阅内部_test.go文件中CheckRuntimeTimerOverflow函数结尾处的注释。
//	// md5:049f552673ac0729
//	CheckRuntimeTimerPeriodOverflow()
//}

func checkZeroPanicString(t *testing.T) {
	e := recover()
	s, _ := e.(string)
	if want := "called on uninitialized Timer"; !strings.Contains(s, want) {
		t.Errorf("panic = %v; want substring %q", e, want)
	}
}

func TestZeroTimerResetPanics(t *testing.T) {
	defer checkZeroPanicString(t)
	var tr Timer
	tr.Reset(1)
}

func TestZeroTimerStopPanics(t *testing.T) {
	defer checkZeroPanicString(t)
	var tr Timer
	tr.Stop()
}

// 测试零时长的定时器不会被调度程序错过。针对问题44868的回归测试。. md5:d08c608639513c32
func TestZeroTimer(t *testing.T) {
	if testing.Short() {
		t.Skip("-short")
	}

	for i := 0; i < 1000000; i++ {
		s := Now()
		ti := NewTimer(0)
		<-ti.F.C
		if diff := Since(s); diff > 2*Second {
			t.Errorf("Expected time to get value from Timer channel in less than 2 sec, took %v", diff)
		}
	}
}

//
//// 测试快速将定时器向前移动不会导致它被丢弃。问题47329。
//// md5:5b7b156696041586
//func TestTimerModifiedEarlier(t *testing.T) {
//	if runtime.GOOS == "plan9" && runtime.GOARCH == "arm" {
//		testenv.SkipFlaky(t, 50470)
//	}
//
//	past := Until(Unix(0, 0))
//	count := 1000
//	fail := 0
//	for i := 0; i < count; i++ {
//		timer := NewTimer(Hour)
//		for j := 0; j < 10; j++ {
//			if !timer.Stop() {
//				<-timer.F.C
//			}
//			timer.Reset(past)
//		}
//
//		deadline := NewTimer(10 * Second)
//		defer deadline.Stop()
//		now := Now()
//		select {
//		case <-timer.F.C:
//			if since := Since(now); since > 8*Second {
//				t.Errorf("timer took too long (%v)", since)
//				fail++
//			}
//		case <-deadline.F.C:
//			t.Error("deadline expired")
//		}
//	}
//
//	if fail > 0 {
//		t.Errorf("%d failures", fail)
//	}
//}
//
//// 测试快速地将计时器提前和延后不会导致
//// 某些休眠时间丢失。
//// 问题 47762
//// md5:e2d773f7ca59cb7e
//func TestAdjustTimers(t *testing.T) {
//	var rnd = rand.New(rand.NewSource(Now().UnixNano()))
//
//	timers := make([]*Timer, 100)
//	states := make([]int, len(timers))
//	indices := rnd.Perm(len(timers))
//
//	for len(indices) != 0 {
//		var ii = rnd.Intn(len(indices))
//		var i = indices[ii]
//
//		var timer = timers[i]
//		var state = states[i]
//		states[i]++
//
//		switch state {
//		case 0:
//			timers[i] = NewTimer(0)
//		case 1:
//			<-timer.F.C // Timer is now idle.
//
//		// 重置为多个长持续时间，我们将会取消这些操作。. md5:4def5d15f369f372
//		case 2:
//			if timer.Reset(1 * Minute) {
//				panic("shouldn't be active (1)")
//			}
//		case 4:
//			if timer.Reset(3 * Minute) {
//				panic("shouldn't be active (3)")
//			}
//		case 6:
//			if timer.Reset(2 * Minute) {
//				panic("shouldn't be active (2)")
//			}
//
//		// 停止并释放一个长时间运行的计时器。. md5:1a645fced10a9490
//		case 3, 5, 7:
//			if !timer.Stop() {
//				t.Logf("timer %d state %d Stop returned false", i, state)
//				<-timer.F.C
//			}
//
//		// 启动一个我们期望不会阻塞的短时定时器。. md5:07ecf94e05725a83
//		case 8:
//			if timer.Reset(0) {
//				t.Fatal("timer.Reset returned true")
//			}
//		case 9:
//			now := Now()
//			<-timer.F.C
//			dur := Since(now)
//			if dur > 750*Millisecond {
//				t.Errorf("timer %d took %v to complete", i, dur)
//			}
//
//		// 计时器已完成。与尾部交换并移除。. md5:d6ea2013f90a3935
//		case 10:
//			indices[ii] = indices[len(indices)-1]
//			indices = indices[:len(indices)-1]
//		}
//	}
//}

// 测试当创建定时器的线程忙于其他工作时，定时器由其他线程服务时的延迟时间。
// 参考：https://golang.org/issue/38860
// md5:ebfd934f369f8698
func BenchmarkParallelTimerLatency(b *testing.B) {
	gmp := runtime.GOMAXPROCS(0)
	if gmp < 2 || runtime.NumCPU() < gmp {
		b.Skip("skipping with GOMAXPROCS < 2 or NumCPU < GOMAXPROCS")
	}

	// 现在分配内存以避免稍后出现垃圾收集干扰。. md5:06108e8f341b6a59
	timerCount := gmp - 1
	stats := make([]struct {
		sum   float64
		max   Duration
		count int64
		_     [5]int64 // cache line padding
	}, timerCount)

	// 确保开始新线程服务定时器的时间不会影响结果。
	// md5:67235cb4b6190520
	warmupScheduler(gmp)

	// 请注意，除了正在测量的AfterFunc调用之外，这个基准测试避免了使用任何其他定时器。具体来说，主goroutine使用doWork来为某些持续时间进行自旋，因为在Go 1.15及之前版本中，如果所有线程都处于空闲状态，当唤醒时sysmon可能会离开深度睡眠状态。
	// md5:48be121202ffd149

	// 确保sysmon处于深度睡眠状态。. md5:dd1381290960838a
	doWork(30 * Millisecond)

	b.ResetTimer()

	const delay = Millisecond
	var wg sync.WaitGroup
	var count int32
	for i := 0; i < b.N; i++ {
		wg.Add(timerCount)
		atomic.StoreInt32(&count, 0)
		for j := 0; j < timerCount; j++ {
			j := j
			expectedWakeup := Now().Add(delay)
			AfterFunc(delay, func() {
				late := Since(expectedWakeup)
				if late < 0 {
					late = 0
				}
				stats[j].count++
				stats[j].sum += float64(late.Nanoseconds())
				if late > stats[j].max {
					stats[j].max = late
				}
				atomic.AddInt32(&count, 1)
				for atomic.LoadInt32(&count) < int32(timerCount) {
					// 一直等待，直到所有计时器触发. md5:633b8d9d10dc0605
				}
				wg.Done()
			})
		}

		for atomic.LoadInt32(&count) < int32(timerCount) {
			// 一直等待，直到所有计时器触发. md5:633b8d9d10dc0605
		}
		wg.Wait()

		// 稍微等待一下，让其他调度线程空闲，然后再进入下一轮。
		// md5:561da4b1365595cd
		doWork(Millisecond)
	}
	var total float64
	var samples float64
	max := Duration(0)
	for _, s := range stats {
		if s.max > max {
			max = s.max
		}
		total += s.sum
		samples += float64(s.count)
	}
	b.ReportMetric(0, "ns/op")
	b.ReportMetric(total/samples, "avg-late-ns")
	b.ReportMetric(float64(max.Nanoseconds()), "max-late-ns")
}

// 使用交错的唤醒时间和不同的CPU密集型工作负载来基准测试计时器延迟。
// https://golang.org/issue/38860
// md5:11caeca6cee3a8ca
func BenchmarkStaggeredTickerLatency(b *testing.B) {
	gmp := runtime.GOMAXPROCS(0)
	if gmp < 2 || runtime.NumCPU() < gmp {
		b.Skip("skipping with GOMAXPROCS < 2 or NumCPU < GOMAXPROCS")
	}

	const delay = 3 * Millisecond

	for _, dur := range []Duration{300 * Microsecond, 2 * Millisecond} {
		b.Run(fmt.Sprintf("work-dur=%s", dur), func(b *testing.B) {
			for tickersPerP := 1; tickersPerP < int(delay/dur)+1; tickersPerP++ {
				tickerCount := gmp * tickersPerP
				b.Run(fmt.Sprintf("tickers-per-P=%d", tickersPerP), func(b *testing.B) {
					// 现在分配内存以避免稍后出现垃圾收集干扰。. md5:06108e8f341b6a59
					stats := make([]struct {
						sum   float64
						max   Duration
						count int64
						_     [5]int64 // cache line padding
					}, tickerCount)

					// 确保启动新线程来处理计时器的时间不会污染结果。
					// md5:41d02fd2d4916706
					warmupScheduler(gmp)

					b.ResetTimer()

					var wg sync.WaitGroup
					wg.Add(tickerCount)
					for j := 0; j < tickerCount; j++ {
						j := j
						doWork(delay / Duration(gmp))
						expectedWakeup := Now().Add(delay)
						ticker := NewTicker(delay)
						go func(c int, ticker *Ticker, firstWake Time) {
							defer ticker.Stop()

							for ; c > 0; c-- {
								<-ticker.F.C
								late := Since(expectedWakeup)
								if late < 0 {
									late = 0
								}
								stats[j].count++
								stats[j].sum += float64(late.Nanoseconds())
								if late > stats[j].max {
									stats[j].max = late
								}
								expectedWakeup = expectedWakeup.Add(delay)
								doWork(dur)
							}
							wg.Done()
						}(b.N, ticker, expectedWakeup)
					}
					wg.Wait()

					var total float64
					var samples float64
					max := Duration(0)
					for _, s := range stats {
						if s.max > max {
							max = s.max
						}
						total += s.sum
						samples += float64(s.count)
					}
					b.ReportMetric(0, "ns/op")
					b.ReportMetric(total/samples, "avg-late-ns")
					b.ReportMetric(float64(max.Nanoseconds()), "max-late-ns")
				})
			}
		})
	}
}

// warmupScheduler 确保调度器的线程池至少有targetThreadCount个线程。
// md5:bc91fc80f58bb498
func warmupScheduler(targetThreadCount int) {
	var wg sync.WaitGroup
	var count int32
	for i := 0; i < targetThreadCount; i++ {
		wg.Add(1)
		go func() {
			atomic.AddInt32(&count, 1)
			for atomic.LoadInt32(&count) < int32(targetThreadCount) {
				// 一直等待，直到所有线程启动. md5:ecc162218f853d96
			}

			// 再等待一段时间，以确保它们都在独立的CPU上运行。. md5:bc7d4fee547b09a1
			doWork(Millisecond)
			wg.Done()
		}()
	}
	wg.Wait()
}

func doWork(dur Duration) {
	start := Now()
	for Since(start) < dur {
	}
}
