// 版权所有 ? 2015 The Go 作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

// Test broken pipes on Unix systems.
//
//go:build !plan9 && !js && !wasip1

package os_test

import (
	"bufio"
	"bytes"
	"fmt"
	"internal/testenv"
	"io"
	"io/fs"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"testing"
	"time"
)

func TestEPIPE(t *testing.T) {
// 由于存在与在 https://go.dev/issue/22315 中报告的类似竞争条件，该测试不能并行运行。
// 
// 尽管管道以 O_CLOEXEC 标志打开，但如果在调用 os.Pipe 和调用 r.Close 之间有其他测试执行了 fork 操作，那么子进程可能会在其 exec 前保留对 r 文件描述符的打开副本。若我们的某个 Write 调用恰好在此期间发生，它可能会意外成功，将写入操作缓冲到子进程持有的管道副本中（尽管子进程实际上并不会读取这些已缓冲的字节）。

	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	if err := r.Close(); err != nil {
		t.Fatal(err)
	}

	expect := syscall.EPIPE
	if runtime.GOOS == "windows" {
		// 232 是Windows错误代码ERROR_NO_DATA，表示"管道正在关闭"。
		expect = syscall.Errno(232)
	}
	// 每次向管道写入时，我们都应该收到一个 EPIPE。
	for i := 0; i < 20; i++ {
		_, err = w.Write([]byte("hi"))
		if err == nil {
			t.Fatal("unexpected success of Write to broken pipe")
		}
		if pe, ok := err.(*fs.PathError); ok {
			err = pe.Err
		}
		if se, ok := err.(*os.SyscallError); ok {
			err = se.Err
		}
		if err != expect {
			t.Errorf("iteration %d: got %v, expected %v", i, err, expect)
		}
	}
}

func TestStdPipe(t *testing.T) {
	switch runtime.GOOS {
	case "windows":
		t.Skip("Windows doesn't support SIGPIPE")
	}

	if os.Getenv("GO_TEST_STD_PIPE_HELPER") != "" {
		if os.Getenv("GO_TEST_STD_PIPE_HELPER_SIGNAL") != "" {
			signal.Notify(make(chan os.Signal, 1), syscall.SIGPIPE)
		}
		switch os.Getenv("GO_TEST_STD_PIPE_HELPER") {
		case "1":
			os.Stdout.Write([]byte("stdout"))
		case "2":
			os.Stderr.Write([]byte("stderr"))
		case "3":
			if _, err := os.NewFile(3, "3").Write([]byte("3")); err == nil {
				os.Exit(3)
			}
		default:
			panic("unrecognized value for GO_TEST_STD_PIPE_HELPER")
		}
// 对于stdout/stderr，我们应该因管道破裂错误而崩溃。
// 调用者会期待那个退出状态，
// 所以在这里正常退出即可在调用者那里引发失败。
// 对于描述符3，期望一个正常的退出。
		os.Exit(0)
	}

	testenv.MustHaveExec(t)
// 由于与TestEPIPE相同的竞态条件，该测试无法并行运行。
// （我们预期向已关闭的管道写入可能会失败，但并发创建子进程可能导致管道意外保持打开状态。）

	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	if err := r.Close(); err != nil {
		t.Fatal(err)
	}
// 调用测试程序来运行测试，并向一个已关闭的管道中写入内容。
// 如果sig为false：
// 向stdout或stderr写入应立即导致SIGPIPE信号；
// 向描述符3写入应失败并返回EPIPE，然后退出状态为0。
// 如果sig为true：
// 所有写入操作都应因EPIPE失败，然后退出状态为0。
	for _, sig := range []bool{false, true} {
		for dest := 1; dest < 4; dest++ {
			cmd := testenv.Command(t, os.Args[0], "-test.run", "TestStdPipe")
			cmd.Stdout = w
			cmd.Stderr = w
			cmd.ExtraFiles = []*os.File{w}
			cmd.Env = append(os.Environ(), fmt.Sprintf("GO_TEST_STD_PIPE_HELPER=%d", dest))
			if sig {
				cmd.Env = append(cmd.Env, "GO_TEST_STD_PIPE_HELPER_SIGNAL=1")
			}
			if err := cmd.Run(); err == nil {
				if !sig && dest < 3 {
					t.Errorf("unexpected success of write to closed pipe %d sig %t in child", dest, sig)
				}
			} else if ee, ok := err.(*exec.ExitError); !ok {
				t.Errorf("unexpected exec error type %T: %v", err, err)
			} else if ws, ok := ee.Sys().(syscall.WaitStatus); !ok {
				t.Errorf("unexpected wait status type %T: %v", ee.Sys(), ee.Sys())
			} else if ws.Signaled() && ws.Signal() == syscall.SIGPIPE {
				if sig || dest > 2 {
					t.Errorf("unexpected SIGPIPE signal for descriptor %d sig %t", dest, sig)
				}
			} else {
				t.Errorf("unexpected exit status %v for descriptor %d sig %t", err, dest, sig)
			}
		}
	}

	// 测试重定向stdout但不stderr。问题40076。
	cmd := testenv.Command(t, os.Args[0], "-test.run", "TestStdPipe")
	cmd.Stdout = w
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	cmd.Env = append(cmd.Environ(), "GO_TEST_STD_PIPE_HELPER=1")
	if err := cmd.Run(); err == nil {
		t.Errorf("unexpected success of write to closed stdout")
	} else if ee, ok := err.(*exec.ExitError); !ok {
		t.Errorf("unexpected exec error type %T: %v", err, err)
	} else if ws, ok := ee.Sys().(syscall.WaitStatus); !ok {
		t.Errorf("unexpected wait status type %T: %v", ee.Sys(), ee.Sys())
	} else if !ws.Signaled() || ws.Signal() != syscall.SIGPIPE {
		t.Errorf("unexpected exit status %v for write to closed stdout", err)
	}
	if output := stderr.Bytes(); len(output) > 0 {
		t.Errorf("unexpected output on stderr: %s", output)
	}
}

func testClosedPipeRace(t *testing.T, read bool) {
// 由于与TestEPIPE相同的竞态条件，该测试无法并行运行。
// （我们预期向已关闭的管道写入可能会失败，但并发创建子进程可能导致管道意外保持打开状态。）

	limit := 1
	if !read {
// 获取需要写入以使无读取器的管道超载的数量。
		limit = 131073
		if b, err := os.ReadFile("/proc/sys/fs/pipe-max-size"); err == nil {
			if i, err := strconv.Atoi(strings.TrimSpace(string(b))); err == nil {
				limit = i + 1
			}
		}
		t.Logf("using pipe write limit of %d", limit)
	}

	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()
	defer w.Close()

// 在我们向管道的写入端写入数据的同时，或反之，在一个goroutine中关闭管道的读取端。
	go func() {
// 给主 Goroutine 一个机会进入 Read 或 Write 调用。
// 这里写得比较粗糙，但即使在读/写之前关闭，测试仍然会通过。
		time.Sleep(20 * time.Millisecond)

		var err error
		if read {
			err = r.Close()
		} else {
			err = w.Close()
		}
		if err != nil {
			t.Error(err)
		}
	}()

	b := make([]byte, limit)
	if read {
		_, err = r.Read(b[:])
	} else {
		_, err = w.Write(b[:])
	}
	if err == nil {
		t.Error("I/O on closed pipe unexpectedly succeeded")
	} else if pe, ok := err.(*fs.PathError); !ok {
		t.Errorf("I/O on closed pipe returned unexpected error type %T; expected fs.PathError", pe)
	} else if pe.Err != fs.ErrClosed {
		t.Errorf("got error %q but expected %q", pe.Err, fs.ErrClosed)
	} else {
		t.Logf("I/O returned expected error %q", err)
	}
}

func TestClosedPipeRaceRead(t *testing.T) {
	testClosedPipeRace(t, true)
}

func TestClosedPipeRaceWrite(t *testing.T) {
	testClosedPipeRace(t, false)
}

// 问题20915：在非阻塞fd上读取时，不应该返回“等待不支持的文件类型”。目前它返回EAGAIN；将来有可能会简单地等待数据。
func TestReadNonblockingFd(t *testing.T) {
	switch runtime.GOOS {
	case "windows":
		t.Skip("Windows doesn't support SetNonblock")
	}
	if os.Getenv("GO_WANT_READ_NONBLOCKING_FD") == "1" {
		fd := syscallDescriptor(os.Stdin.Fd())
		syscall.SetNonblock(fd, true)
		defer syscall.SetNonblock(fd, false)
		_, err := os.Stdin.Read(make([]byte, 1))
		if err != nil {
			if perr, ok := err.(*fs.PathError); !ok || perr.Err != syscall.EAGAIN {
				t.Fatalf("read on nonblocking stdin got %q, should have gotten EAGAIN", err)
			}
		}
		os.Exit(0)
	}

	testenv.MustHaveExec(t)
	t.Parallel()

	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()
	defer w.Close()
	cmd := testenv.Command(t, os.Args[0], "-test.run=^"+t.Name()+"$")
	cmd.Env = append(cmd.Environ(), "GO_WANT_READ_NONBLOCKING_FD=1")
	cmd.Stdin = r
	output, err := cmd.CombinedOutput()
	t.Logf("%s", output)
	if err != nil {
		t.Errorf("child process failed: %v", err)
	}
}

func TestCloseWithBlockingReadByNewFile(t *testing.T) {
	t.Parallel()

	var p [2]syscallDescriptor
	err := syscall.Pipe(p[:])
	if err != nil {
		t.Fatal(err)
	}
	// os.NewFile返回一个阻塞模式的文件。
	testCloseWithBlockingRead(t, os.NewFile(uintptr(p[0]), "reader"), os.NewFile(uintptr(p[1]), "writer"))
}

func TestCloseWithBlockingReadByFd(t *testing.T) {
	t.Parallel()

	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	// 调用Fd将使文件进入阻塞模式。
	_ = r.Fd()
	testCloseWithBlockingRead(t, r, w)
}

// 测试确保不会让阻塞读取阻止关闭操作。
func testCloseWithBlockingRead(t *testing.T, r, w *os.File) {
	var (
		enteringRead = make(chan struct{})
		done         = make(chan struct{})
	)
	go func() {
		var b [1]byte
		close(enteringRead)
		_, err := r.Read(b[:])
		if err == nil {
			t.Error("I/O on closed pipe unexpectedly succeeded")
		}

		if pe, ok := err.(*fs.PathError); ok {
			err = pe.Err
		}
		if err != io.EOF && err != fs.ErrClosed {
			t.Errorf("got %v, expected EOF or closed", err)
		}
		close(done)
	}()

// 让goroutine有机会进入Read或Write调用。这很粗糙，但即使我们在读写之前关闭它，测试也会通过。
	<-enteringRead
	time.Sleep(20 * time.Millisecond)

	if err := r.Close(); err != nil {
		t.Error(err)
	}
// r.Close 已经完成，但由于我们假设 r 在阻塞模式下，这可能并没有解阻塞对 r.Read 的调用。关闭 w 来解阻塞它。
	w.Close()
	<-done
}

func TestPipeEOF(t *testing.T) {
	t.Parallel()

	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}

	testPipeEOF(t, r, w)
}

// testPipeEOF 测试当管道或FIFO的写入端被关闭时，
// 读取端阻塞的Read调用会返回io.EOF。
//
// 此场景之前在darwin上未能使Read调用解除阻塞。
// （参见 https://go.dev/issue/24164。）
func testPipeEOF(t *testing.T, r io.ReadCloser, w io.WriteCloser) {
// parkDelay 是一个任意的延迟，我们等待管道读取器goroutine暂停
// 在执行相应的写入之前。测试应该无论使用什么延迟都能通过，
// 但使用更长的延迟更有可能检测到轮询器的bug。
	parkDelay := 10 * time.Millisecond
	if testing.Short() {
		parkDelay = 100 * time.Microsecond
	}
	writerDone := make(chan struct{})
	defer func() {
		if err := r.Close(); err != nil {
			t.Errorf("error closing reader: %v", err)
		}
		<-writerDone
	}()

	write := make(chan int, 1)
	go func() {
		defer close(writerDone)

		for i := range write {
			time.Sleep(parkDelay)
			_, err := fmt.Fprintf(w, "line %d\n", i)
			if err != nil {
				t.Errorf("error writing to fifo: %v", err)
				return
			}
		}

		time.Sleep(parkDelay)
		if err := w.Close(); err != nil {
			t.Errorf("error closing writer: %v", err)
		}
	}()

	rbuf := bufio.NewReader(r)
	for i := 0; i < 3; i++ {
		write <- i
		b, err := rbuf.ReadBytes('\n')
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%s\n", bytes.TrimSpace(b))
	}

	close(write)
	b, err := rbuf.ReadBytes('\n')
	if err != io.EOF || len(b) != 0 {
		t.Errorf(`ReadBytes: %q, %v; want "", io.EOF`, b, err)
	}
}

// Issue 24481.
func TestFdRace(t *testing.T) {
// 这个测试启动了100个并发的goroutine，如果这个或其他测试发生恐慌，可能会隐藏更有趣的堆栈。同时，它几乎是瞬间完成的，所以并行运行时的任何延迟好处都会非常小。

	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()
	defer w.Close()

	var wg sync.WaitGroup
	call := func() {
		defer wg.Done()
		w.Fd()
	}

	const tries = 100
	for i := 0; i < tries; i++ {
		wg.Add(1)
		go call()
	}
	wg.Wait()
}

func TestFdReadRace(t *testing.T) {
	t.Parallel()

	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()
	defer w.Close()

	const count = 10

	c := make(chan bool, 1)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		var buf [count]byte
		r.SetReadDeadline(time.Now().Add(time.Minute))
		c <- true
		if _, err := r.Read(buf[:]); os.IsTimeout(err) {
			t.Error("read timed out")
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		<-c
// 让其他goroutine有机会进入Read。
// 如果这个操作偶尔失败也没关系，测试仍然会通过，只是不会检验任何内容。
		time.Sleep(10 * time.Millisecond)
		r.Fd()

// 该bug表现为Fd会一直阻塞直到Read超时。
// 若该bug已被修复，那么在此处向w写入数据并关闭r
// 将导致Read在超时之前退出。
		w.Write(make([]byte, count))
		r.Close()
	}()

	wg.Wait()
}
