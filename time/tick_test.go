//版权所有2009年Go作者。所有权利保留。
//使用此源代码受BSD风格
//可以在LICENSE文件中找到的许可证。
// md5:2e9dc81828a3be8a

package time_test

import (
	"fmt"
	. "github.com/888go/gosdk/time"
	"runtime"
	"testing"
)

func TestTicker(t *testing.T) {
	// 我们想要测试ticker是否按预期消耗时间。
	// 由于我们不希望测试运行时间过长，因此不想使用过长的时间间隔。
	// 这使得测试本质上可能不稳定。
	// 先从较短的时间间隔开始，但如果第一次测试失败，再尝试使用较长的时间间隔。
	// md5:ad7e20ba1cc8ab8f

	baseCount := 10
	baseDelta := 20 * Millisecond

	// 在Darwin ARM64上，tick频率似乎受到限制。问题35692。. md5:7a0c4fbdc27fe8a7
	if (runtime.GOOS == "darwin" || runtime.GOOS == "ios") && runtime.GOARCH == "arm64" {
		// 下面的测试会运行计时器 `count/2` 次，然后将计时器的周期重置为剩余 `count/2` 的两倍。由于 Darwin ARM64 系统的计时频率有限，我们使用偶数来给计时器更多的时间让测试通过。参考 CL（Change List）220638。
		// md5:9679bd76bf1f345b
		baseCount = 6
		baseDelta = 100 * Millisecond
	}

	var errs []string
	logErrs := func() {
		for _, e := range errs {
			t.Log(e)
		}
	}

	for _, test := range []struct {
		count int
		delta Duration
	}{{
		count: baseCount,
		delta: baseDelta,
	}, {
		count: 8,
		delta: 1 * Second,
	}} {
		count, delta := test.count, test.delta
		ticker := NewTicker(delta)
		t0 := Now()
		for i := 0; i < count/2; i++ {
			<-ticker.F.C
		}
		ticker.Reset(delta * 2)
		for i := count / 2; i < count; i++ {
			<-ticker.F.C
		}
		ticker.Stop()
		t1 := Now()
		dt := t1.Sub(t0)
		target := 3 * delta * Duration(count/2)
		slop := target * 3 / 10
		if dt < target-slop || dt > target+slop {
			errs = append(errs, fmt.Sprintf("%d %s ticks then %d %s ticks took %s, expected [%s,%s]", count/2, delta, count/2, delta*2, dt, target-slop, target+slop))
			if dt > target+slop {
				// 系统可能过载；稍微睡一会儿，希望它能恢复。
				// md5:58a2798075f6c4be
				Sleep(Second / 2)
			}
			continue
		}
		// 现在测试ticker是否已停止。. md5:ba5a168a74615cec
		Sleep(2 * delta)
		select {
		case <-ticker.F.C:
			errs = append(errs, "Ticker did not shut down")
			continue
		default:
			// ok
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

//// Issue 21874
//func TestTickerStopWithDirectInitialization(t *testing.T) {
//	c := make(chan Time)
//	tk := &Ticker{C: c}
//	tk.Stop()
//}

// 测试修复了定时器拆卸时的bug。这个程序不应该出现死锁。. md5:306140edccec42f2
func TestTeardown(t *testing.T) {
	Delta := 100 * Millisecond
	if testing.Short() {
		Delta = 20 * Millisecond
	}
	for i := 0; i < 3; i++ {
		ticker := NewTicker(Delta)
		<-ticker.F.C
		ticker.Stop()
	}
}

// 测试Tick便利包装器。. md5:901afab2f3b2cca2
func TestTick(t *testing.T) {
	// 测试给一个负的持续时间是否返回nil。. md5:a6ce42ca827ba1db
	if got := Tick(-1); got != nil {
		t.Errorf("Tick(-1) = %v; want nil", got)
	}
}

// 测试当给定的持续时间小于零时，NewTicker是否会发生恐慌。. md5:c56054c00884775b
func TestNewTickerLtZeroDuration(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Errorf("NewTicker(-1) should have panicked")
		}
	}()
	NewTicker(-1)
}

// 测试当向Ticker重置传递一个小于零的持续时间时，它是否会引发恐慌。. md5:609097bbc9af9981
func TestTickerResetLtZeroDuration(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Errorf("Ticker.Reset(0) should have panicked")
		}
	}()
	tk := NewTicker(Second)
	tk.Reset(0)
}

func BenchmarkTicker(b *testing.B) {
	benchmark(b, func(n int) {
		ticker := NewTicker(Nanosecond)
		for i := 0; i < n; i++ {
			<-ticker.F.C
		}
		ticker.Stop()
	})
}

func BenchmarkTickerReset(b *testing.B) {
	benchmark(b, func(n int) {
		ticker := NewTicker(Nanosecond)
		for i := 0; i < n; i++ {
			ticker.Reset(Nanosecond * 2)
		}
		ticker.Stop()
	})
}

func BenchmarkTickerResetNaive(b *testing.B) {
	benchmark(b, func(n int) {
		ticker := NewTicker(Nanosecond)
		for i := 0; i < n; i++ {
			ticker.Stop()
			ticker = NewTicker(Nanosecond * 2)
		}
		ticker.Stop()
	})
}
