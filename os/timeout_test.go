// 版权所有 2017 The Go 作者。保留所有权利。
// 使用本源代码须遵循 BSD 风格
// 许可协议，该协议可在 LICENSE 文件中找到。

//go:build !js && !plan9 && !wasip1 && !windows

package os_test

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
	"testing"
	"time"
)

func TestNonpollableDeadline(t *testing.T) {
// 在BSD系统中，似乎普通文件是可被轮询的，因此只在Linux上运行这个测试。
	if runtime.GOOS != "linux" {
		t.Skipf("skipping on %s", runtime.GOOS)
	}
	t.Parallel()

	f, err := os.CreateTemp("", "ostest")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(f.Name())
	defer f.Close()
	deadline := time.Now().Add(10 * time.Second)
	if err := f.SetDeadline(deadline); err != os.ErrNoDeadline {
		t.Errorf("SetDeadline on file returned %v, wanted %v", err, os.ErrNoDeadline)
	}
	if err := f.SetReadDeadline(deadline); err != os.ErrNoDeadline {
		t.Errorf("SetReadDeadline on file returned %v, wanted %v", err, os.ErrNoDeadline)
	}
	if err := f.SetWriteDeadline(deadline); err != os.ErrNoDeadline {
		t.Errorf("SetWriteDeadline on file returned %v, wanted %v", err, os.ErrNoDeadline)
	}
}

// noDeadline是一个零值time.Time，用于取消截止时间。
var noDeadline time.Time

var readTimeoutTests = []struct {
	timeout time.Duration
	xerrs   [2]error // 预期在转换中出现的错误
}{
// 测试读取超时是否起作用，即使已经有数据可以读取。
	{-5 * time.Second, [2]error{os.ErrDeadlineExceeded, os.ErrDeadlineExceeded}},

	{50 * time.Millisecond, [2]error{nil, os.ErrDeadlineExceeded}},
}

// 这里有一个与 net/timeout_test.go 中非常相似的副本。
func TestReadTimeout(t *testing.T) {
	t.Parallel()

	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()
	defer w.Close()

	if _, err := w.Write([]byte("READ TIMEOUT TEST")); err != nil {
		t.Fatal(err)
	}

	for i, tt := range readTimeoutTests {
		if err := r.SetReadDeadline(time.Now().Add(tt.timeout)); err != nil {
			t.Fatalf("#%d: %v", i, err)
		}
		var b [1]byte
		for j, xerr := range tt.xerrs {
			for {
				n, err := r.Read(b[:])
				if xerr != nil {
					if !isDeadlineExceeded(err) {
						t.Fatalf("#%d/%d: %v", i, j, err)
					}
				}
				if err == nil {
					time.Sleep(tt.timeout / 3)
					continue
				}
				if n != 0 {
					t.Fatalf("#%d/%d: read %d; want 0", i, j, n)
				}
				break
			}
		}
	}
}

// 这里有一个与 net/timeout_test.go 中非常相似的副本。
func TestReadTimeoutMustNotReturn(t *testing.T) {
	t.Parallel()

	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()
	defer w.Close()

	max := time.NewTimer(100 * time.Millisecond)
	defer max.Stop()
	ch := make(chan error)
	go func() {
		if err := r.SetDeadline(time.Now().Add(-5 * time.Second)); err != nil {
			t.Error(err)
		}
		if err := r.SetWriteDeadline(time.Now().Add(-5 * time.Second)); err != nil {
			t.Error(err)
		}
		if err := r.SetReadDeadline(noDeadline); err != nil {
			t.Error(err)
		}
		var b [1]byte
		_, err := r.Read(b[:])
		ch <- err
	}()

	select {
	case err := <-ch:
		t.Fatalf("expected Read to not return, but it returned with %v", err)
	case <-max.C:
		w.Close()
		err := <-ch // 等待测试者goroutine停止
		if os.IsTimeout(err) {
			t.Fatal(err)
		}
	}
}

var writeTimeoutTests = []struct {
	timeout time.Duration
	xerrs   [2]error // 预期在转换中出现的错误
}{
// 测试写入截止日期有效，即使有可用的缓冲区空间进行写入。
	{-5 * time.Second, [2]error{os.ErrDeadlineExceeded, os.ErrDeadlineExceeded}},

	{10 * time.Millisecond, [2]error{nil, os.ErrDeadlineExceeded}},
}

// 这里有一个与 net/timeout_test.go 中非常相似的副本。
func TestWriteTimeout(t *testing.T) {
	t.Parallel()

	for i, tt := range writeTimeoutTests {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			r, w, err := os.Pipe()
			if err != nil {
				t.Fatal(err)
			}
			defer r.Close()
			defer w.Close()

			if err := w.SetWriteDeadline(time.Now().Add(tt.timeout)); err != nil {
				t.Fatalf("%v", err)
			}
			for j, xerr := range tt.xerrs {
				for {
					n, err := w.Write([]byte("WRITE TIMEOUT TEST"))
					if xerr != nil {
						if !isDeadlineExceeded(err) {
							t.Fatalf("%d: %v", j, err)
						}
					}
					if err == nil {
						time.Sleep(tt.timeout / 3)
						continue
					}
					if n != 0 {
						t.Fatalf("%d: wrote %d; want 0", j, n)
					}
					break
				}
			}
		})
	}
}

// 这里有一个与 net/timeout_test.go 中非常相似的副本。
func TestWriteTimeoutMustNotReturn(t *testing.T) {
	t.Parallel()

	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()
	defer w.Close()

	max := time.NewTimer(100 * time.Millisecond)
	defer max.Stop()
	ch := make(chan error)
	go func() {
		if err := w.SetDeadline(time.Now().Add(-5 * time.Second)); err != nil {
			t.Error(err)
		}
		if err := w.SetReadDeadline(time.Now().Add(-5 * time.Second)); err != nil {
			t.Error(err)
		}
		if err := w.SetWriteDeadline(noDeadline); err != nil {
			t.Error(err)
		}
		var b [1]byte
		for {
			if _, err := w.Write(b[:]); err != nil {
				ch <- err
				break
			}
		}
	}()

	select {
	case err := <-ch:
		t.Fatalf("expected Write to not return, but it returned with %v", err)
	case <-max.C:
		r.Close()
		err := <-ch // 等待测试者goroutine停止
		if os.IsTimeout(err) {
			t.Fatal(err)
		}
	}
}

const (
// minDynamicTimeout是自动增加超时时间直到成功的测试尝试的最小超时时间。
// 
// 如果值接近真实最小值，较小的值可能会让测试更快成功。但如果超时设置得太低，可能需要更多迭代（并且在失败的尝试中浪费更多的时间和CPU资源）。
	minDynamicTimeout = 1 * time.Millisecond

// maxDynamicTimeout 是自动增加超时时间直到成功的测试尝试的最大超时时间。
// 
// 这应该是一个严格的上限，即使在运行缓慢或负载过重的机器上，也需要达到准确触发超时所需的延迟。如果测试会将超时时间增加到此值之上，那么测试将失败。
	maxDynamicTimeout = 4 * time.Second
)

// timeoutUpperBound 函数返回一个预期的最大时长，即对于持续时间为 d 的超时而言，我们期望其能让调用者返回所花费的最长时间。
func timeoutUpperBound(d time.Duration) time.Duration {
	switch runtime.GOOS {
	case "openbsd", "netbsd":
// NetBSD和OpenBSD似乎无法在长时间的绝对期限内可靠地满足期限。
// 在https://build.golang.org/log/c34f8685d020b98377dd4988cd38f0c5bd72267e中，
// 我们观察到一个openbsd-amd64-68构建器在2.983020682s的超时时限中花费了4.090948779s（37.1%的开销）。
// （有关更多详细信息，请参阅https://go.dev/issue/50189。）
// 为了补偿，给他们留出大量的余地。
		return d * 3 / 2
	}
// 其他平台似乎更可靠地满足其截止日期，至少当它们足够长以覆盖调度抖动时。
	return d * 11 / 10
}

// nextTimeout 在操作实际耗时超过给定超时时，返回下一个尝试的超时时间。
func nextTimeout(actual time.Duration) (next time.Duration, ok bool) {
	if actual >= maxDynamicTimeout {
		return maxDynamicTimeout, false
	}
// 由于上一次尝试已耗时actual，我们无法期望在显著程度上超过该耗时。对下一次尝试使用该耗时之上的一个任意因子，以便我们的增长曲线至少为指数级。
	next = actual * 5 / 4
	if next > maxDynamicTimeout {
		return maxDynamicTimeout, true
	}
	return next, true
}

// 这里有一个与 net/timeout_test.go 中非常相似的副本。
func TestReadTimeoutFluctuation(t *testing.T) {
	t.Parallel()

	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()
	defer w.Close()

	d := minDynamicTimeout
	b := make([]byte, 256)
	for {
		t.Logf("SetReadDeadline(+%v)", d)
		t0 := time.Now()
		deadline := t0.Add(d)
		if err = r.SetReadDeadline(deadline); err != nil {
			t.Fatalf("SetReadDeadline(%v): %v", deadline, err)
		}
		var n int
		n, err = r.Read(b)
		t1 := time.Now()

		if n != 0 || err == nil || !isDeadlineExceeded(err) {
			t.Errorf("Read did not return (0, timeout): (%d, %v)", n, err)
		}

		actual := t1.Sub(t0)
		if t1.Before(deadline) {
			t.Errorf("Read took %s; expected at least %s", actual, d)
		}
		if t.Failed() {
			return
		}
		if want := timeoutUpperBound(d); actual > want {
			next, ok := nextTimeout(actual)
			if !ok {
				t.Fatalf("Read took %s; expected at most %v", actual, want)
			}
// 可能这台机器的性能太慢，无法在请求的时间内可靠地调度goroutines。增加超时时间并重试。
			t.Logf("Read took %s (expected %s); trying with longer timeout", actual, d)
			d = next
			continue
		}

		break
	}
}

// 这里有一个与 net/timeout_test.go 中非常相似的副本。
func TestWriteTimeoutFluctuation(t *testing.T) {
	t.Parallel()

	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()
	defer w.Close()

	d := minDynamicTimeout
	for {
		t.Logf("SetWriteDeadline(+%v)", d)
		t0 := time.Now()
		deadline := t0.Add(d)
		if err := w.SetWriteDeadline(deadline); err != nil {
			t.Fatalf("SetWriteDeadline(%v): %v", deadline, err)
		}
		var n int64
		var err error
		for {
			var dn int
			dn, err = w.Write([]byte("TIMEOUT TRANSMITTER"))
			n += int64(dn)
			if err != nil {
				break
			}
		}
		t1 := time.Now()
		// Inv: err != nil
		if !isDeadlineExceeded(err) {
			t.Fatalf("Write did not return (any, timeout): (%d, %v)", n, err)
		}

		actual := t1.Sub(t0)
		if t1.Before(deadline) {
			t.Errorf("Write took %s; expected at least %s", actual, d)
		}
		if t.Failed() {
			return
		}
		if want := timeoutUpperBound(d); actual > want {
			if n > 0 {
				// SetWriteDeadline specifies a time “after which I/O operations fail
				// instead of blocking”. However, the kernel's send buffer is not yet
				// full, we may be able to write some arbitrary (but finite) number of
				// bytes to it without blocking.
				t.Logf("Wrote %d bytes into send buffer; retrying until buffer is full", n)
				if d <= maxDynamicTimeout/2 {
// 如果缓冲区已满，我们不知道实际写入循环会花多长时间，所以猜测一下并将其持续时间翻倍，这样下一次尝试就可以向填充缓冲区进更多步。
					d *= 2
				}
			} else if next, ok := nextTimeout(actual); !ok {
				t.Fatalf("Write took %s; expected at most %s", actual, want)
			} else {
				// Maybe this machine is too slow to reliably schedule goroutines within
				// the requested duration. Increase the timeout and try again.
				t.Logf("Write took %s (expected %s); trying with longer timeout", actual, d)
				d = next
			}
			continue
		}

		break
	}
}

// 这里有一个与 net/timeout_test.go 中非常相似的副本。
func TestVariousDeadlines(t *testing.T) {
	t.Parallel()
	testVariousDeadlines(t)
}

// 这里有一个与 net/timeout_test.go 中非常相似的副本。
func TestVariousDeadlines1Proc(t *testing.T) {
	// 不能使用 t.Parallel - 会修改全局的 GOMAXPROCS。
	if testing.Short() {
		t.Skip("skipping in short mode")
	}
	defer runtime.GOMAXPROCS(runtime.GOMAXPROCS(1))
	testVariousDeadlines(t)
}

// 这里有一个与 net/timeout_test.go 中非常相似的副本。
func TestVariousDeadlines4Proc(t *testing.T) {
	// 不能使用 t.Parallel - 会修改全局的 GOMAXPROCS。
	if testing.Short() {
		t.Skip("skipping in short mode")
	}
	defer runtime.GOMAXPROCS(runtime.GOMAXPROCS(4))
	testVariousDeadlines(t)
}

type neverEnding byte

func (b neverEnding) Read(p []byte) (int, error) {
	for i := range p {
		p[i] = byte(b)
	}
	return len(p), nil
}

func testVariousDeadlines(t *testing.T) {
	type result struct {
		n   int64
		err error
		d   time.Duration
	}

	handler := func(w *os.File, pasvch chan result) {
// 该写入器自身无任何超时设置，
// 尽可能快速地向客户端发送字节。
		t0 := time.Now()
		n, err := io.Copy(w, neverEnding('a'))
		dt := time.Since(t0)
		pasvch <- result{n, err, dt}
	}

	for _, timeout := range []time.Duration{
		1 * time.Nanosecond,
		2 * time.Nanosecond,
		5 * time.Nanosecond,
		50 * time.Nanosecond,
		100 * time.Nanosecond,
		200 * time.Nanosecond,
		500 * time.Nanosecond,
		750 * time.Nanosecond,
		1 * time.Microsecond,
		5 * time.Microsecond,
		25 * time.Microsecond,
		250 * time.Microsecond,
		500 * time.Microsecond,
		1 * time.Millisecond,
		5 * time.Millisecond,
		100 * time.Millisecond,
		250 * time.Millisecond,
		500 * time.Millisecond,
		1 * time.Second,
	} {
		numRuns := 3
		if testing.Short() {
			numRuns = 1
			if timeout > 500*time.Microsecond {
				continue
			}
		}
		for run := 0; run < numRuns; run++ {
			t.Run(fmt.Sprintf("%v-%d", timeout, run+1), func(t *testing.T) {
				r, w, err := os.Pipe()
				if err != nil {
					t.Fatal(err)
				}
				defer r.Close()
				defer w.Close()

				pasvch := make(chan result)
				go handler(w, pasvch)

				tooLong := 5 * time.Second
				max := time.NewTimer(tooLong)
				defer max.Stop()
				actvch := make(chan result)
				go func() {
					t0 := time.Now()
					if err := r.SetDeadline(t0.Add(timeout)); err != nil {
						t.Error(err)
					}
					n, err := io.Copy(io.Discard, r)
					dt := time.Since(t0)
					r.Close()
					actvch <- result{n, err, dt}
				}()

				select {
				case res := <-actvch:
					if !isDeadlineExceeded(err) {
						t.Logf("good client timeout after %v, reading %d bytes", res.d, res.n)
					} else {
						t.Fatalf("client Copy = %d, %v; want timeout", res.n, res.err)
					}
				case <-max.C:
					t.Fatalf("timeout (%v) waiting for client to timeout (%v) reading", tooLong, timeout)
				}

				select {
				case res := <-pasvch:
					t.Logf("writer in %v wrote %d: %v", res.d, res.n, res.err)
				case <-max.C:
					t.Fatalf("timeout waiting for writer to finish writing")
				}
			})
		}
	}
}

// 这里有一个与 net/timeout_test.go 中非常相似的副本。
func TestReadWriteDeadlineRace(t *testing.T) {
	t.Parallel()

	N := 1000
	if testing.Short() {
		N = 50
	}

	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()
	defer w.Close()

	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		defer wg.Done()
		tic := time.NewTicker(2 * time.Microsecond)
		defer tic.Stop()
		for i := 0; i < N; i++ {
			if err := r.SetReadDeadline(time.Now().Add(2 * time.Microsecond)); err != nil {
				break
			}
			if err := w.SetWriteDeadline(time.Now().Add(2 * time.Microsecond)); err != nil {
				break
			}
			<-tic.C
		}
	}()
	go func() {
		defer wg.Done()
		var b [1]byte
		for i := 0; i < N; i++ {
			_, err := r.Read(b[:])
			if err != nil && !isDeadlineExceeded(err) {
				t.Error("Read returned non-timeout error", err)
			}
		}
	}()
	go func() {
		defer wg.Done()
		var b [1]byte
		for i := 0; i < N; i++ {
			_, err := w.Write(b[:])
			if err != nil && !isDeadlineExceeded(err) {
				t.Error("Write returned non-timeout error", err)
			}
		}
	}()
	wg.Wait() // 等待测试者goroutine停止
}

// TestRacyRead 测试在取消操作发生后立即修改输入读取缓冲区是安全的。
func TestRacyRead(t *testing.T) {
	t.Parallel()

	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()
	defer w.Close()

	var wg sync.WaitGroup
	defer wg.Wait()

	go io.Copy(w, rand.New(rand.NewSource(0)))

	r.SetReadDeadline(time.Now().Add(time.Millisecond))
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			b1 := make([]byte, 1024)
			b2 := make([]byte, 1024)
			for j := 0; j < 100; j++ {
				_, err := r.Read(b1)
				copy(b1, b2) // 修改b1以触发潜在的竞争条件
				if err != nil {
					if !isDeadlineExceeded(err) {
						t.Error(err)
					}
					r.SetReadDeadline(time.Now().Add(time.Millisecond))
				}
			}
		}()
	}
}

// TestRacyWrite 测试在取消操作发生后，立即修改输入写缓冲区是安全的。
func TestRacyWrite(t *testing.T) {
	t.Parallel()

	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()
	defer w.Close()

	var wg sync.WaitGroup
	defer wg.Wait()

	go io.Copy(io.Discard, r)

	w.SetWriteDeadline(time.Now().Add(time.Millisecond))
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			b1 := make([]byte, 1024)
			b2 := make([]byte, 1024)
			for j := 0; j < 100; j++ {
				_, err := w.Write(b1)
				copy(b1, b2) // 修改b1以触发潜在的竞争条件
				if err != nil {
					if !isDeadlineExceeded(err) {
						t.Error(err)
					}
					w.SetWriteDeadline(time.Now().Add(time.Millisecond))
				}
			}
		}()
	}
}

// 在从TTY读取数据时关闭它，不应导致程序挂起。问题 23943。
func TestTTYClose(t *testing.T) {
	// 如果我们在后台运行，忽略SIGTTIN信号。
	signal.Ignore(syscall.SIGTTIN)
	defer signal.Reset(syscall.SIGTTIN)

	f, err := os.Open("/dev/tty")
	if err != nil {
		t.Skipf("skipping because opening /dev/tty failed: %v", err)
	}

	go func() {
		var buf [1]byte
		f.Read(buf[:])
	}()

// 让goroutine有机会进入读取阶段。
// 如果偶尔未能做到这一点没有太大关系，
// 我们不会测试到我们想要的，但测试会通过。
	time.Sleep(time.Millisecond)

	c := make(chan bool)
	go func() {
		defer close(c)
		f.Close()
	}()

	select {
	case <-c:
	case <-time.After(time.Second):
		t.Error("timed out waiting for close")
	}

// 在某些系统中，goroutine 可能会挂起。
// 对此我们无能为力。
}
