// 版权所有 ? 2009 Go作者。保留所有权利。
// 本源代码的使用受BSD风格
// 许可证约束，该许可证可在LICENSE文件中找到。

// 使用外部测试以避免在非cgo环境下的darwin系统中出现os/exec -> net/http -> crypto/x509 -> os/exec的循环依赖关系

// 2024-04-17 备注,单元测试通不过, 保留单元测试文件为了方便查看使用方法.
package exec_test

import (
	"bufio"
	"bytes"
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/888go/gosdk/exec"
	"github.com/888go/gosdk/exec/internal/testenv"
	exec2 "os/exec"

	"github.com/888go/gosdk/exec/internal/fdtest"
	"io"
	"net"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

// haveUnexpectedFDs在初始化时被设置，用于报告程序启动时是否有任何文件描述符处于打开状态。
var haveUnexpectedFDs bool

func init() {
	godebug := os.Getenv("GODEBUG")
	if godebug != "" {
		godebug += ","
	}
	godebug += "execwait=2"
	os.Setenv("GODEBUG", godebug)

	if os.Getenv("GO_EXEC_TEST_PID") != "" {
		return
	}
	if runtime.GOOS == "windows" {
		return
	}
	// 对于fd从uintptr类型数值3开始，到数值100结束的循环：
	// 如果poll.IsPollDescriptor(fd)判断fd为poll描述符，
	// 则跳过本次循环继续下一轮。
	//
	// 若fdtest.Exists(fd)检测到fd存在，
	// 则将haveUnexpectedFDs设为true并立即返回。
}

// TestMain 允许测试二进制程序模拟多个其他程序，
// 其中一些可能操作 os.Stdin、os.Stdout 和/或 os.Stderr
// （因此无法作为普通 Test 函数运行，因为 testing 包在运行测试前会对这些变量进行monkey patch）。
func TestMain(m *testing.M) {
	flag.Parse()

	pid := os.Getpid()
	if os.Getenv("GO_EXEC_TEST_PID") == "" {
		os.Setenv("GO_EXEC_TEST_PID", strconv.Itoa(pid))

		if runtime.GOOS == "windows" {
			// Normalize environment so that test behavior is consistent.
			// (The behavior of LookPath varies depending on this variable.)
			//
			// Ideally we would test both with the variable set and with it cleared,
			// but I (bcmills) am not sure that that's feasible: it may already be set
			// in the Windows registry, and I'm not sure if it is possible to remove
			// a registry variable in a program's environment.
			//
			// Per https://learn.microsoft.com/en-us/windows/win32/api/processenv/nf-processenv-needcurrentdirectoryforexepathw#remarks,
			// “the existence of the NoDefaultCurrentDirectoryInExePath environment
			// variable is checked, and not its value.”
			os.Setenv("NoDefaultCurrentDirectoryInExePath", "TRUE")
		}

		code := m.Run()
		if code == 0 && flag.Lookup("test.run").Value.String() == "" && flag.Lookup("test.list").Value.String() == "" {
			for cmd := range helperCommands {
				if _, ok := helperCommandUsed.Load(cmd); !ok {
					fmt.Fprintf(os.Stderr, "helper command unused: %q\n", cmd)
					code = 1
				}
			}
		}

		if !testing.Short() {
			// 运行几个GC周期，以提高通过由GODEBUG=execwait=2安装的终结器检测进程泄漏的可能性
			runtime.GC()
			runtime.GC()
		}

		os.Exit(code)
	}

	args := flag.Args()
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "No command\n")
		os.Exit(2)
	}

	cmd, args := args[0], args[1:]
	f, ok := helperCommands[cmd]
	if !ok {
		fmt.Fprintf(os.Stderr, "Unknown command %q\n", cmd)
		os.Exit(2)
	}
	f(args...)
	os.Exit(0)
}

// registerHelperCommand 注册一个测试进程可以模拟执行的命令。
// 命令应在使用它的同一源文件中进行注册。
// 如果所有测试都运行且通过，那么所有已注册的命令必须被使用。
// （这可以防止在随时间移除或重构测试时积聚过时的命令。）
func registerHelperCommand(name string, f func(...string)) {
	if helperCommands[name] != nil {
		panic("duplicate command registered: " + name)
	}
	helperCommands[name] = f
}

// maySkipHelperCommand 用于记录使用指定辅助命令的测试已调用，但在实际调用 helperCommand 之前可能调用 Skip。
func maySkipHelperCommand(name string) {
	helperCommandUsed.Store(name, true)
}

// helperCommand 返回一个 exec.Cmd，该命令将运行指定名称的辅助命令。
func helperCommand(t *testing.T, name string, args ...string) *exec.Cmd {
	t.Helper()
	return helperCommandContext(t, nil, name, args...)
}

// helperCommandContext 与 helperCommand 类似，但还接受一个在其下运行命令的 Context。
func helperCommandContext(t *testing.T, ctx context.Context, name string, args ...string) (cmd *exec.Cmd) {
	helperCommandUsed.LoadOrStore(name, true)

	t.Helper()
	testenv.MustHaveExec(t)

	cs := append([]string{name}, args...)
	if ctx != nil {
		cmd = exec.CommandContext(ctx, exePath(t), cs...)
	} else {
		cmd = exec.Command(exePath(t), cs...)
	}
	return cmd
}

// exePath 返回运行中可执行文件的路径
func exePath(t testing.TB) string {
	exeOnce.Do(func() {
		// 使用os.Executable替代os.Args[0]，以防调用者修改
		// cmd.Dir：如果以"./exec.test"形式调用测试二进制文件，
		// 则不应出现异常失败。
		exeOnce.path, exeOnce.err = os.Executable()
	})

	if exeOnce.err != nil {
		if t == nil {
			panic(exeOnce.err)
		}
		t.Fatal(exeOnce.err)
	}

	return exeOnce.path
}

var exeOnce struct {
	path string
	err  error
	sync.Once
}

func chdir(t *testing.T, dir string) {
	t.Helper()

	prev, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	if err := os.Chdir(dir); err != nil {
		t.Fatal(err)
	}
	t.Logf("Chdir(%#q)", dir)

	t.Cleanup(func() {
		if err := os.Chdir(prev); err != nil {
			// 无法切换回原始工作目录。
			// 使用panic替代t.Fatal，以防止我们在未知位置运行其他测试。
			panic("couldn't restore working directory: " + err.Error())
		}
	})
}

var helperCommandUsed sync.Map

var helperCommands = map[string]func(...string){
	"echo":          cmdEcho,
	"echoenv":       cmdEchoEnv,
	"cat":           cmdCat,
	"pipetest":      cmdPipeTest,
	"stdinClose":    cmdStdinClose,
	"exit":          cmdExit,
	"describefiles": cmdDescribeFiles,
	"stderrfail":    cmdStderrFail,
	"yes":           cmdYes,
	"hang":          cmdHang,
}

func cmdEcho(args ...string) {
	iargs := []any{}
	for _, s := range args {
		iargs = append(iargs, s)
	}
	fmt.Println(iargs...)
}

func cmdEchoEnv(args ...string) {
	for _, s := range args {
		fmt.Println(os.Getenv(s))
	}
}

func cmdCat(args ...string) {
	if len(args) == 0 {
		io.Copy(os.Stdout, os.Stdin)
		return
	}
	exit := 0
	for _, fn := range args {
		f, err := os.Open(fn)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			exit = 2
		} else {
			defer f.Close()
			io.Copy(os.Stdout, f)
		}
	}
	os.Exit(exit)
}

func cmdPipeTest(...string) {
	bufr := bufio.NewReader(os.Stdin)
	for {
		line, _, err := bufr.ReadLine()
		if err == io.EOF {
			break
		} else if err != nil {
			os.Exit(1)
		}
		if bytes.HasPrefix(line, []byte("O:")) {
			os.Stdout.Write(line)
			os.Stdout.Write([]byte{'\n'})
		} else if bytes.HasPrefix(line, []byte("E:")) {
			os.Stderr.Write(line)
			os.Stderr.Write([]byte{'\n'})
		} else {
			os.Exit(1)
		}
	}
}

func cmdStdinClose(...string) {
	b, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	if s := string(b); s != stdinCloseTestString {
		fmt.Fprintf(os.Stderr, "Error: Read %q, want %q", s, stdinCloseTestString)
		os.Exit(1)
	}
}

func cmdExit(args ...string) {
	n, _ := strconv.Atoi(args[0])
	os.Exit(n)
}

func cmdDescribeFiles(args ...string) {
	f := os.NewFile(3, fmt.Sprintf("fd3"))
	ln, err := net.FileListener(f)
	if err == nil {
		fmt.Printf("fd3: listener %s\n", ln.Addr())
		ln.Close()
	}
}

func cmdStderrFail(...string) {
	fmt.Fprintf(os.Stderr, "some stderr text\n")
	os.Exit(1)
}

func cmdYes(args ...string) {
	if len(args) == 0 {
		args = []string{"y"}
	}
	s := strings.Join(args, " ") + "\n"
	for {
		_, err := os.Stdout.WriteString(s)
		if err != nil {
			os.Exit(1)
		}
	}
}

func TestEcho(t *testing.T) {
	t.Parallel()

	bs, err := helperCommand(t, "echo", "foo bar", "baz").Output()
	if err != nil {
		t.Errorf("echo: %v", err)
	}
	if g, e := string(bs), "foo bar baz\n"; g != e {
		t.Errorf("echo: want %q, got %q", e, g)
	}
}

func TestCommandRelativeName(t *testing.T) {
	t.Parallel()

	cmd := helperCommand(t, "echo", "foo")

	// 以相对路径方式运行我们自己的二进制文件（例如 "_test/exec.test"），该路径位于父目录中。
	base := filepath.Base(os.Args[0]) // "exec.test"
	dir := filepath.Dir(os.Args[0])   // "/tmp/go-buildNNNN/os/exec/_test"
	if dir == "." {
		t.Skip("skipping; running test at root somehow")
	}
	parentDir := filepath.Dir(dir) // "/tmp/go-buildNNNN/os/exec"
	dirBase := filepath.Base(dir)  // "_test"
	if dirBase == "." {
		t.Skipf("skipping; unexpected shallow dir of %q", dir)
	}

	cmd.F.Path = filepath.Join(dirBase, base)
	cmd.F.Dir = parentDir

	out, err := cmd.Output()
	if err != nil {
		t.Errorf("echo: %v", err)
	}
	if g, e := string(out), "foo\n"; g != e {
		t.Errorf("echo: want %q, got %q", e, g)
	}
}

func TestCatStdin(t *testing.T) {
	t.Parallel()

	// Cat，用于测试标准输入和标准输出
	input := "Input string\nLine 2"
	p := helperCommand(t, "cat")
	p.F.Stdin = strings.NewReader(input)
	bs, err := p.Output()
	if err != nil {
		t.Errorf("cat: %v", err)
	}
	s := string(bs)
	if s != input {
		t.Errorf("cat: want %q, got %q", input, s)
	}
}

func TestEchoFileRace(t *testing.T) {
	t.Parallel()

	cmd := helperCommand(t, "echo")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		t.Fatalf("StdinPipe: %v", err)
	}
	if err := cmd.Start(); err != nil {
		t.Fatalf("Start: %v", err)
	}
	wrote := make(chan bool)
	go func() {
		defer close(wrote)
		fmt.Fprint(stdin, "echo\n")
	}()
	if err := cmd.Wait(); err != nil {
		t.Fatalf("Wait: %v", err)
	}
	<-wrote
}

func TestCatGoodAndBadFile(t *testing.T) {
	t.Parallel()

	// 测试输出与错误值的组合情况
	bs, err := helperCommand(t, "cat", "/bogus/file.foo", "exec_test.go").CombinedOutput()
	if _, ok := err.(*exec2.ExitError); !ok {
		t.Errorf("expected *exec.ExitError from cat combined; got %T: %v", err, err)
	}
	errLine, body, ok := strings.Cut(string(bs), "\n")
	if !ok {
		t.Fatalf("expected two lines from cat; got %q", bs)
	}
	if !strings.HasPrefix(errLine, "Error: open /bogus/file.foo") {
		t.Errorf("expected stderr to complain about file; got %q", errLine)
	}
	if !strings.Contains(body, "func TestCatGoodAndBadFile(t *testing.T)") {
		t.Errorf("expected test code; got %q (len %d)", body, len(body))
	}
}

func TestNoExistExecutable(t *testing.T) {
	t.Parallel()

	// 无法运行不存在的可执行文件
	err := exec.Command("/no-exist-executable").Run()
	if err == nil {
		t.Error("expected error from /no-exist-executable")
	}
}

func TestExitStatus(t *testing.T) {
	t.Parallel()

	// 测试退出值是否正确返回
	cmd := helperCommand(t, "exit", "42")
	err := cmd.Run()
	want := "exit status 42"
	switch runtime.GOOS {
	case "plan9":
		want = fmt.Sprintf("exit status: '%s %d: 42'", filepath.Base(cmd.F.Path), cmd.F.ProcessState.Pid())
	}
	if werr, ok := err.(*exec2.ExitError); ok {
		if s := werr.Error(); s != want {
			t.Errorf("from exit 42 got exit %q, want %q", s, want)
		}
	} else {
		t.Fatalf("expected *exec.ExitError from exit 42; got %T: %v", err, err)
	}
}

func TestExitCode(t *testing.T) {
	t.Parallel()

	// 测试返回的退出码是否正确
	cmd := helperCommand(t, "exit", "42")
	cmd.Run()
	want := 42
	if runtime.GOOS == "plan9" {
		want = 1
	}
	got := cmd.F.ProcessState.ExitCode()
	if want != got {
		t.Errorf("ExitCode got %d, want %d", got, want)
	}

	cmd = helperCommand(t, "/no-exist-executable")
	cmd.Run()
	want = 2
	if runtime.GOOS == "plan9" {
		want = 1
	}
	got = cmd.F.ProcessState.ExitCode()
	if want != got {
		t.Errorf("ExitCode got %d, want %d", got, want)
	}

	cmd = helperCommand(t, "exit", "255")
	cmd.Run()
	want = 255
	if runtime.GOOS == "plan9" {
		want = 1
	}
	got = cmd.F.ProcessState.ExitCode()
	if want != got {
		t.Errorf("ExitCode got %d, want %d", got, want)
	}

	cmd = helperCommand(t, "cat")
	cmd.Run()
	want = 0
	got = cmd.F.ProcessState.ExitCode()
	if want != got {
		t.Errorf("ExitCode got %d, want %d", got, want)
	}

	// 测试命令不调用Run()的情况
	cmd = helperCommand(t, "cat")
	want = -1
	got = cmd.F.ProcessState.ExitCode()
	if want != got {
		t.Errorf("ExitCode got %d, want %d", got, want)
	}
}

func TestPipes(t *testing.T) {
	t.Parallel()

	check := func(what string, err error) {
		if err != nil {
			t.Fatalf("%s: %v", what, err)
		}
	}
	// Cat，用于测试标准输入和标准输出
	c := helperCommand(t, "pipetest")
	stdin, err := c.StdinPipe()
	check("StdinPipe", err)
	stdout, err := c.StdoutPipe()
	check("StdoutPipe", err)
	stderr, err := c.StderrPipe()
	check("StderrPipe", err)

	outbr := bufio.NewReader(stdout)
	errbr := bufio.NewReader(stderr)
	line := func(what string, br *bufio.Reader) string {
		line, _, err := br.ReadLine()
		if err != nil {
			t.Fatalf("%s: %v", what, err)
		}
		return string(line)
	}

	err = c.Start()
	check("Start", err)

	_, err = stdin.Write([]byte("O:I am output\n"))
	check("first stdin Write", err)
	if g, e := line("first output line", outbr), "O:I am output"; g != e {
		t.Errorf("got %q, want %q", g, e)
	}

	_, err = stdin.Write([]byte("E:I am error\n"))
	check("second stdin Write", err)
	if g, e := line("first error line", errbr), "E:I am error"; g != e {
		t.Errorf("got %q, want %q", g, e)
	}

	_, err = stdin.Write([]byte("O:I am output2\n"))
	check("third stdin Write 3", err)
	if g, e := line("second output line", outbr), "O:I am output2"; g != e {
		t.Errorf("got %q, want %q", g, e)
	}

	stdin.Close()
	err = c.Wait()
	check("Wait", err)
}

const stdinCloseTestString = "Some test string."

// Issue 6270.
func TestStdinClose(t *testing.T) {
	t.Parallel()

	check := func(what string, err error) {
		if err != nil {
			t.Fatalf("%s: %v", what, err)
		}
	}
	cmd := helperCommand(t, "stdinClose")
	stdin, err := cmd.StdinPipe()
	check("StdinPipe", err)
	// 检查我们是否可以访问底层os.File的方法
	if _, ok := stdin.(interface {
		Fd() uintptr
	}); !ok {
		t.Error("can't access methods of underlying *os.File")
	}
	check("Start", cmd.Start())

	var wg sync.WaitGroup
	wg.Add(1)
	defer wg.Wait()
	go func() {
		defer wg.Done()

		_, err := io.Copy(stdin, strings.NewReader(stdinCloseTestString))
		check("Copy", err)

		// 修复前，下一行代码将与 cmd.Wait 产生竞态条件。
		if err := stdin.Close(); err != nil && !errors.Is(err, os.ErrClosed) {
			t.Errorf("Close: %v", err)
		}
	}()

	check("Wait", cmd.Wait())
}

// 问题 17647。
// 之前，当在竞态检测器下运行时，上述的 TestStdinClose 测试会失败。
// 本测试是 TestStdinClose 的一个变种，同样在竞态检测器下运行时也会失败。
// 此测试由 cmd/dist 在竞态检测器下运行，用于验证竞态检测器不再报告任何问题。
func TestStdinCloseRace(t *testing.T) {
	t.Parallel()

	cmd := helperCommand(t, "stdinClose")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		t.Fatalf("StdinPipe: %v", err)
	}
	if err := cmd.Start(); err != nil {
		t.Fatalf("Start: %v", err)

	}

	var wg sync.WaitGroup
	wg.Add(2)
	defer wg.Wait()

	go func() {
		defer wg.Done()
		// 我们不对 Kill 的错误返回值进行检查。有可能该进程已经退出，此时 Kill 会返回一个“process already finished”的错误。本测试的目的是观察竞态检测器是否报告错误；Kill 是否成功并不重要。
		cmd.F.Process.Kill()
	}()

	go func() {
		defer wg.Done()
		// 发送错误的字符串，确保即使其他goroutine未能先将其终止，子进程也能失败。
		// 本测试旨在检查竞态检测器不会错误地报告错误，因此子进程如何失败并不重要。
		io.Copy(stdin, strings.NewReader("unexpected string"))
		if err := stdin.Close(); err != nil && !errors.Is(err, os.ErrClosed) {
			t.Errorf("stdin.Close: %v", err)
		}
	}()

	if err := cmd.Wait(); err == nil {
		t.Fatalf("Wait: succeeded unexpectedly")
	}
}

// Issue 5071
func TestPipeLookPathLeak(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("we don't currently suppore counting open handles on windows")
	}
	// 非并行：检查是否有泄露的文件描述符

	openFDs := func() []uintptr {
		var fds []uintptr
		for i := uintptr(0); i < 100; i++ {
			if fdtest.Exists(i) {
				fds = append(fds, i)
			}
		}
		return fds
	}

	old := map[uintptr]bool{}
	for _, fd := range openFDs() {
		old[fd] = true
	}

	for i := 0; i < 6; i++ {
		cmd := exec.Command("something-that-does-not-exist-executable")
		cmd.StdoutPipe()
		cmd.StderrPipe()
		cmd.StdinPipe()
		if err := cmd.Run(); err == nil {
			t.Fatal("unexpected success")
		}
	}

	// 由于此测试并非并行运行，我们预期在它执行过程中不会打开任何新的文件描述符。然而，如果在测试开始时存在额外的FD（例如由libc打开），这些FD可能因某种超时而被关闭。允许它们消失，但要检查没有新增加的FD。
	for _, fd := range openFDs() {
		if !old[fd] {
			t.Errorf("leaked file descriptor %v", fd)
		}
	}
}

//func TestExtraFiles(t *testing.T) {
//	if testing.Short() {
//		t.Skipf("skipping test in short mode that would build a helper binary")
//	}
//
//	if haveUnexpectedFDs {
//		// 本测试的目的是确保我们打开的所有描述符都被标记为“close-on-exec”。
//		// 如果`haveUnexpectedFDs`为真，说明我们在开始测试时已有其他描述符处于打开状态，
//		// 显然这些描述符没有设置“close-on-exec”，它们会干扰本次测试。虽然我们可以修改测试，
//		// 使其预期这些描述符保持打开，但由于我们并不清楚它们来自何处、有何作用，
//		// 这种做法显得很脆弱。例如，可能是因为某种原因，它们是由当前系统的启动代码所创建的。
//		// 另外，本测试并非针对特定系统；只要大部分系统未跳过此测试，我们仍能对关注点进行有效验证。
//		t.Skip("skipping test because test was run with FDs open")
//	}
//
//	testenv.MustHaveExec(t)
//	testenv.MustHaveGoBuild(t)
//
//	// 该测试在禁用cgo的情况下运行。外部链接需要cgo，因此如果需要外部链接，则无法正常工作。
//	testenv.MustInternalLink(t, false)
//
//	if runtime.GOOS == "windows" {
//		t.Skipf("skipping test on %q", runtime.GOOS)
//	}
//
//	// 强制使用网络，以验证 epoll（或其他类似机制）的文件描述符
//	// 不会泄漏给子进程
//	ln, err := net.Listen("tcp", "127.0.0.1:0")
//	if err != nil {
//		t.Fatal(err)
//	}
//	defer ln.Close()
//
//	// 确保重复的文件描述符不会泄露给子进程。
//	f, err := ln.(*net.TCPListener).File()
//	if err != nil {
//		t.Fatal(err)
//	}
//	defer f.Close()
//	ln2, err := net.FileListener(f)
//	if err != nil {
//		t.Fatal(err)
//	}
//	defer ln2.Close()
//
//	// 强制加载TLS根证书（可能涉及cgo），以确保潜在的C代码不会泄漏文件描述符。
//	ts := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
//	// 忽略预期的TLS握手错误“remote error: bad certificate”
//	ts.Config.ErrorLog = log.New(io.Discard, "", 0)
//	ts.StartTLS()
//	defer ts.Close()
//	_, err = http.Get(ts.URL)
//	if err == nil {
//		t.Errorf("success trying to fetch %s; want an error", ts.URL)
//	}
//
//	tf, err := os.CreateTemp("", "")
//	if err != nil {
//		t.Fatalf("TempFile: %v", err)
//	}
//	defer os.Remove(tf.Name())
//	defer tf.Close()
//
//	const text = "Hello, fd 3!"
//	_, err = tf.Write([]byte(text))
//	if err != nil {
//		t.Fatalf("Write: %v", err)
//	}
//	_, err = tf.Seek(0, io.SeekStart)
//	if err != nil {
//		t.Fatalf("Seek: %v", err)
//	}
//
//	tempdir := t.TempDir()
//	exe := filepath.Join(tempdir, "read3.exe")
//
//	c := testenv.Command(t, testenv.GoToolPath(t), "build", "-o", exe, "read3.go")
//	// 在无cgo的情况下构建测试，以便C库函数不会意外打开描述符。参见问题25628。
//	c.Env = append(os.Environ(), "CGO_ENABLED=0")
//	if output, err := c.CombinedOutput(); err != nil {
//		t.Logf("go build -o %s read3.go\n%s", exe, output)
//		t.Fatalf("go build failed: %v", err)
//	}
//
//	// 使用截止时间来尝试获取输出，即使程序挂起也能有所收获
//	ctx := context.Background()
//	if deadline, ok := t.Deadline(); ok {
//		// 留出20%的宽限期以刷新输出，这在 linux/386 构建器上可能很大，因为我们正在使用 strace 运行子进程。
//		deadline = deadline.Add(-time.Until(deadline) / 5)
//
//		var cancel context.CancelFunc
//		ctx, cancel = context.WithDeadline(ctx, deadline)
//		defer cancel()
//	}
//
//	c = exec.CommandContext(ctx, exe)
//	var stdout, stderr strings.Builder
//
//	c.Stdout = &stdout
//	c.Stderr = &stderr
//	c.ExtraFiles = []*os.File{tf}
//	if runtime.GOOS == "illumos" {
//		// illumos 中某些设施的实现依赖于 libc 对 /proc 的访问，这种访问可能会暂时占用一个低位文件描述符（fd）。如果这种情况恰好与检查泄漏描述符的测试并发发生，该检查可能会变得混乱，并错误地报告一个泄漏的描述符。（关于更详细的分析，请参见问题 #42431。）
//		//
//		// 试图限制子进程中额外线程的使用，以减少此测试出现不稳定性：
//		c.Env = append(os.Environ(), "GOMAXPROCS=1")
//	}
//	err = c.Run()
//	if err != nil {
//		t.Fatalf("Run: %v\n--- stdout:\n%s--- stderr:\n%s", err, stdout.String(), stderr.String())
//	}
//	if stdout.String() != text {
//		t.Errorf("got stdout %q, stderr %q; want %q on stdout", stdout.String(), stderr.String(), text)
//	}
//}

func TestExtraFilesRace(t *testing.T) {
	if runtime.GOOS == "windows" {
		maySkipHelperCommand("describefiles")
		t.Skip("no operating system support; skipping")
	}
	t.Parallel()

	listen := func() net.Listener {
		ln, err := net.Listen("tcp", "127.0.0.1:0")
		if err != nil {
			t.Fatal(err)
		}
		return ln
	}
	listenerFile := func(ln net.Listener) *os.File {
		f, err := ln.(*net.TCPListener).File()
		if err != nil {
			t.Fatal(err)
		}
		return f
	}
	runCommand := func(c *exec.Cmd, out chan<- string) {
		bout, err := c.CombinedOutput()
		if err != nil {
			out <- "ERROR:" + err.Error()
		} else {
			out <- string(bout)
		}
	}

	for i := 0; i < 10; i++ {
		if testing.Short() && i >= 3 {
			break
		}
		la := listen()
		ca := helperCommand(t, "describefiles")
		ca.F.ExtraFiles = []*os.File{listenerFile(la)}
		lb := listen()
		cb := helperCommand(t, "describefiles")
		cb.F.ExtraFiles = []*os.File{listenerFile(lb)}
		ares := make(chan string)
		bres := make(chan string)
		go runCommand(ca, ares)
		go runCommand(cb, bres)
		if got, want := <-ares, fmt.Sprintf("fd3: listener %s\n", la.Addr()); got != want {
			t.Errorf("iteration %d, process A got:\n%s\nwant:\n%s\n", i, got, want)
		}
		if got, want := <-bres, fmt.Sprintf("fd3: listener %s\n", lb.Addr()); got != want {
			t.Errorf("iteration %d, process B got:\n%s\nwant:\n%s\n", i, got, want)
		}
		la.Close()
		lb.Close()
		for _, f := range ca.F.ExtraFiles {
			f.Close()
		}
		for _, f := range cb.F.ExtraFiles {
			f.Close()
		}
	}
}

type delayedInfiniteReader struct{}

func (delayedInfiniteReader) Read(b []byte) (int, error) {
	time.Sleep(100 * time.Millisecond)
	for i := range b {
		b[i] = 'x'
	}
	return len(b), nil
}

// Issue 9173：若程序成功完成，忽略对stdin管道的写入操作。
func TestIgnorePipeErrorOnSuccess(t *testing.T) {
	t.Parallel()

	testWith := func(r io.Reader) func(*testing.T) {
		return func(t *testing.T) {
			t.Parallel()

			cmd := helperCommand(t, "echo", "foo")
			var out strings.Builder
			cmd.F.Stdin = r
			cmd.F.Stdout = &out
			if err := cmd.Run(); err != nil {
				t.Fatal(err)
			}
			if got, want := out.String(), "foo\n"; got != want {
				t.Errorf("output = %q; want %q", got, want)
			}
		}
	}
	t.Run("10MB", testWith(strings.NewReader(strings.Repeat("x", 10<<20))))
	t.Run("Infinite", testWith(delayedInfiniteReader{}))
}

type badWriter struct{}

func (w *badWriter) Write(data []byte) (int, error) {
	return 0, io.ErrUnexpectedEOF
}

func TestClosePipeOnCopyError(t *testing.T) {
	t.Parallel()

	cmd := helperCommand(t, "yes")
	cmd.F.Stdout = new(badWriter)
	err := cmd.Run()
	if err == nil {
		t.Errorf("yes unexpectedly completed successfully")
	}
}

func TestOutputStderrCapture(t *testing.T) {
	t.Parallel()

	cmd := helperCommand(t, "stderrfail")
	_, err := cmd.Output()
	ee, ok := err.(*exec.ExitError)
	if !ok {
		t.Fatalf("Output error type = %T; want ExitError", err)
	}
	got := string(ee.F.Stderr)
	want := "some stderr text\n"
	if got != want {
		t.Errorf("ExitError.Stderr = %q; want %q", got, want)
	}
}

func TestContext(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	c := helperCommandContext(t, ctx, "pipetest")
	stdin, err := c.StdinPipe()
	if err != nil {
		t.Fatal(err)
	}
	stdout, err := c.StdoutPipe()
	if err != nil {
		t.Fatal(err)
	}
	if err := c.Start(); err != nil {
		t.Fatal(err)
	}

	if _, err := stdin.Write([]byte("O:hi\n")); err != nil {
		t.Fatal(err)
	}
	buf := make([]byte, 5)
	n, err := io.ReadFull(stdout, buf)
	if n != len(buf) || err != nil || string(buf) != "O:hi\n" {
		t.Fatalf("ReadFull = %d, %v, %q", n, err, buf[:n])
	}
	go cancel()

	if err := c.Wait(); err == nil {
		t.Fatal("expected Wait failure")
	}
}

//func TestContextCancel(t *testing.T) {
//	if runtime.GOOS == "netbsd" && runtime.GOARCH == "arm64" {
//		maySkipHelperCommand("cat")
//		testenv.SkipFlaky(t, 42061)
//	}
//
//	// 为了减少最终协程转储中的噪声，
//	// 尽可能让其他并行测试完成。
//	t.Parallel()
//
//	ctx, cancel := context.WithCancel(context.Background())
//	defer cancel()
//	c := helperCommandContext(t, ctx, "cat")
//
//	stdin, err := c.StdinPipe()
//	if err != nil {
//		t.Fatal(err)
//	}
//	defer stdin.Close()
//
//	if err := c.Start(); err != nil {
//		t.Fatal(err)
//	}
//
//	// 此时进程处于活跃状态。通过向其stdin发送数据来确保这一点。
//	if _, err := io.WriteString(stdin, "echo"); err != nil {
//		t.Fatal(err)
//	}
//
//	cancel()
//
//	// 调用cancel应当已终止该进程，因此现在写入应会失败。给进程一点时间结束。
//	start := time.Now()
//	delay := 1 * time.Millisecond
//	for {
//		if _, err := io.WriteString(stdin, "echo"); err != nil {
//			break
//		}
//
//		if time.Since(start) > time.Minute {
//			// 通过 panic 替代调用 t.Fatal，以便我们获取 goroutine 堆栈信息。
//			// 我们希望确切了解 os/exec goroutines 阻塞在何处。
//			debug.SetTraceback("system")
//			panic("canceling context did not stop program")
//		}
//
//		// 以指数方式回退（最多休眠1秒），以便给操作系统时间来终止进程。
//		delay *= 2
//		if delay > 1*time.Second {
//			delay = 1 * time.Second
//		}
//		time.Sleep(delay)
//	}
//
//	if err := c.Wait(); err == nil {
//		t.Error("program unexpectedly exited successfully")
//	} else {
//		t.Logf("exit status: %v", err)
//	}
//}

// 测试环境变量是否被去重
func TestDedupEnvEcho(t *testing.T) {
	t.Parallel()

	cmd := helperCommand(t, "echoenv", "FOO")
	cmd.F.Env = append(cmd.Environ(), "FOO=bad", "FOO=good")
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatal(err)
	}
	if got, want := strings.TrimSpace(string(out)), "good"; got != want {
		t.Errorf("output = %q; want %q", got, want)
	}
}

func TestEnvNULCharacter(t *testing.T) {
	if runtime.GOOS == "plan9" {
		t.Skip("plan9 explicitly allows NUL in the environment")
	}
	cmd := helperCommand(t, "echoenv", "FOO", "BAR")
	cmd.F.Env = append(cmd.Environ(), "FOO=foo\x00BAR=bar")
	out, err := cmd.CombinedOutput()
	if err == nil {
		t.Errorf("output = %q; want error", string(out))
	}
}

func TestString(t *testing.T) {
	t.Parallel()

	echoPath, err := exec.LookPath("echo")
	if err != nil {
		t.Skip(err)
	}
	tests := [...]struct {
		path string
		args []string
		want string
	}{
		{"echo", nil, echoPath},
		{"echo", []string{"a"}, echoPath + " a"},
		{"echo", []string{"a", "b"}, echoPath + " a b"},
	}
	for _, test := range tests {
		cmd := exec.Command(test.path, test.args...)
		if got := cmd.String(); got != test.want {
			t.Errorf("String(%q, %q) = %q, want %q", test.path, test.args, got, test.want)
		}
	}
}

func TestStringPathNotResolved(t *testing.T) {
	t.Parallel()

	_, err := exec.LookPath("makemeasandwich")
	if err == nil {
		t.Skip("wow, thanks")
	}

	cmd := exec.Command("makemeasandwich", "-lettuce")
	want := "makemeasandwich -lettuce"
	if got := cmd.String(); got != want {
		t.Errorf("String(%q, %q) = %q, want %q", "makemeasandwich", "-lettuce", got, want)
	}
}

func TestNoPath(t *testing.T) {
	err := new(exec2.Cmd).Start()
	want := "exec: no command"
	if err == nil || err.Error() != want {
		t.Errorf("new(Cmd).Start() = %v, want %q", err, want)
	}
}

// TestDoubleStartLeavesPipesOpen 测试一个回归问题，即当两次调用
// Start 方法时（第二次调用会返回错误），是否会导致第一次调用中建立的管道被错误地关闭。
func TestDoubleStartLeavesPipesOpen(t *testing.T) {
	t.Parallel()

	cmd := helperCommand(t, "pipetest")
	in, err := cmd.StdinPipe()
	if err != nil {
		t.Fatal(err)
	}
	out, err := cmd.StdoutPipe()
	if err != nil {
		t.Fatal(err)
	}

	if err := cmd.Start(); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := cmd.Wait(); err != nil {
			t.Error(err)
		}
	})

	if err := cmd.Start(); err == nil || !strings.HasSuffix(err.Error(), "already started") {
		t.Fatalf("second call to Start returned a nil; want an 'already started' error")
	}

	outc := make(chan []byte, 1)
	go func() {
		b, err := io.ReadAll(out)
		if err != nil {
			t.Error(err)
		}
		outc <- b
	}()

	const msg = "O:Hello, pipe!\n"

	_, err = io.WriteString(in, msg)
	if err != nil {
		t.Fatal(err)
	}
	in.Close()

	b := <-outc
	if !bytes.Equal(b, []byte(msg)) {
		t.Fatalf("read %q from stdout pipe; want %q", b, msg)
	}
}

func cmdHang(args ...string) {
	sleep, err := time.ParseDuration(args[0])
	if err != nil {
		panic(err)
	}

	fs := flag.NewFlagSet("hang", flag.ExitOnError)
	exitOnInterrupt := fs.Bool("interrupt", false, "if true, commands should exit 0 on os.Interrupt")
	subsleep := fs.Duration("subsleep", 0, "amount of time for the 'hang' helper to leave an orphaned subprocess sleeping with stderr open")
	probe := fs.Duration("probe", 0, "if nonzero, the 'hang' helper should write to stderr at this interval, and exit nonzero if a write fails")
	read := fs.Bool("read", false, "if true, the 'hang' helper should read stdin to completion before sleeping")
	fs.Parse(args[1:])

	pid := os.Getpid()

	if *subsleep != 0 {
		cmd := exec.Command(exePath(nil), "hang", subsleep.String(), "-read=true", "-probe="+probe.String())
		cmd.F.Stdin = os.Stdin
		cmd.F.Stderr = os.Stderr
		out, err := cmd.StdoutPipe()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		cmd.Start()

		buf := new(strings.Builder)
		if _, err := io.Copy(buf, out); err != nil {
			fmt.Fprintln(os.Stderr, err)
			cmd.F.Process.Kill()
			cmd.Wait()
			os.Exit(1)
		}
		fmt.Fprintf(os.Stderr, "%d: started %d: %v\n", pid, cmd.F.Process.Pid, cmd)
		go cmd.Wait() // 如果cmd未能在此进程之后继续存在，则释放资源
	}

	if *exitOnInterrupt {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		go func() {
			sig := <-c
			fmt.Fprintf(os.Stderr, "%d: received %v\n", pid, sig)
			os.Exit(0)
		}()
	} else {
		signal.Ignore(os.Interrupt)
	}

	// 通过关闭stdout来表明进程已设置完毕
	os.Stdout.Close()

	if *read {
		if pipeSignal != nil {
			signal.Ignore(pipeSignal)
		}
		r := bufio.NewReader(os.Stdin)
		for {
			line, err := r.ReadBytes('\n')
			if len(line) > 0 {
				// 忽略写入错误：即使stderr被关闭，我们也希望继续读取。
				fmt.Fprintf(os.Stderr, "%d: read %s", pid, line)
			}
			if err != nil {
				fmt.Fprintf(os.Stderr, "%d: finished read: %v", pid, err)
				break
			}
		}
	}

	if *probe != 0 {
		ticker := time.NewTicker(*probe)
		go func() {
			for range ticker.C {
				if _, err := fmt.Fprintf(os.Stderr, "%d: ok\n", pid); err != nil {
					os.Exit(1)
				}
			}
		}()
	}

	if sleep != 0 {
		time.Sleep(sleep)
		fmt.Fprintf(os.Stderr, "%d: slept %v\n", pid, sleep)
	}
}

// tickReader 读取一个无界的时间戳序列，但每个时间戳之间的间隔不会超过一个固定的值。
type tickReader struct {
	interval time.Duration
	lastTick time.Time
	s        string
}

func newTickReader(interval time.Duration) *tickReader {
	return &tickReader{interval: interval}
}

func (r *tickReader) Read(p []byte) (n int, err error) {
	if len(r.s) == 0 {
		if d := r.interval - time.Since(r.lastTick); d > 0 {
			time.Sleep(d)
		}
		r.lastTick = time.Now()
		r.s = r.lastTick.Format(time.RFC3339Nano + "\n")
	}

	n = copy(p, r.s)
	r.s = r.s[n:]
	return n, nil
}

func startHang(t *testing.T, ctx context.Context, hangTime time.Duration, interrupt os.Signal, waitDelay time.Duration, flags ...string) *exec.Cmd {
	t.Helper()

	args := append([]string{hangTime.String()}, flags...)
	cmd := helperCommandContext(t, ctx, "hang", args...)
	cmd.F.Stdin = newTickReader(1 * time.Millisecond)
	cmd.F.Stderr = new(strings.Builder)
	if interrupt == nil {
		cmd.F.Cancel = nil
	} else {
		cmd.F.Cancel = func() error {
			return cmd.F.Process.Signal(interrupt)
		}
	}
	cmd.F.WaitDelay = waitDelay
	out, err := cmd.StdoutPipe()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(cmd)
	if err := cmd.Start(); err != nil {
		t.Fatal(err)
	}

	// 等待cmd关闭stdout以表明其处理器已安装完成
	buf := new(strings.Builder)
	if _, err := io.Copy(buf, out); err != nil {
		t.Error(err)
		cmd.F.Process.Kill()
		cmd.Wait()
		t.FailNow()
	}
	if buf.Len() > 0 {
		t.Logf("stdout %v:\n%s", cmd.F.Args, buf)
	}

	return cmd
}

func TestWaitInterrupt(t *testing.T) {
	t.Parallel()

	// tooLong 是一个任意的时长，预期远超过测试运行所需时间，但又足够短，以至于泄露的进程最终能够自行退出
	const tooLong = 10 * time.Minute

	// 控制用例：在没有取消操作且未设置等待延迟的情况下，我们应该等待进程退出。
	t.Run("Wait", func(t *testing.T) {
		t.Parallel()
		cmd := startHang(t, context.Background(), 1*time.Millisecond, os.Kill, 0)
		err := cmd.Wait()
		t.Logf("stderr:\n%s", cmd.F.Stderr)
		t.Logf("[%d] %v", cmd.F.Process.Pid, err)

		if err != nil {
			t.Errorf("Wait: %v; want <nil>", err)
		}
		if ps := cmd.F.ProcessState; !ps.Exited() {
			t.Errorf("cmd did not exit: %v", ps)
		} else if code := ps.ExitCode(); code != 0 {
			t.Errorf("cmd.ProcessState.ExitCode() = %v; want 0", code)
		}
	})

	// 当WaitDelay设置得非常长且没有Cancel函数时，即使命令的Context被取消，我们也应该等待进程退出。
	t.Run("WaitDelay", func(t *testing.T) {
		if runtime.GOOS == "windows" {
			t.Skipf("skipping: os.Interrupt is not implemented on Windows")
		}
		t.Parallel()

		ctx, cancel := context.WithCancel(context.Background())
		cmd := startHang(t, ctx, tooLong, nil, tooLong, "-interrupt=true")
		cancel()

		time.Sleep(1 * time.Millisecond)
		// 此时，cmd 应该仍在运行（因为我们向 startHang 传递了 nil 作为取消信号）。向其发送一个明确的 Interrupt 信号应该能成功。
		if err := cmd.F.Process.Signal(os.Interrupt); err != nil {
			t.Error(err)
		}

		err := cmd.Wait()
		t.Logf("stderr:\n%s", cmd.F.Stderr)
		t.Logf("[%d] %v", cmd.F.Process.Pid, err)

		// 该程序以状态0退出，
		// 但在等待延迟期间几乎总是如此。
		// 由于Cmd本身在上下文过期时没有做任何事情来停止进程，
		// 因此即使退出较晚，成功退出也是有效的，
		// 并不值得返回一个非空错误。
		if err != nil {
			t.Errorf("Wait: %v; want nil", err)
		}
		if ps := cmd.F.ProcessState; !ps.Exited() {
			t.Errorf("cmd did not exit: %v", ps)
		} else if code := ps.ExitCode(); code != 0 {
			t.Errorf("cmd.ProcessState.ExitCode() = %v; want 0", code)
		}
	})

	// 如果上下文被取消且Cancel函数发送os.Kill信号，
	// 则进程应立即终止，且其输出管道应在WaitDelay后关闭（导致Wait返回），
	// 即使子进程仍在向这些管道写入数据。
	t.Run("SIGKILL-hang", func(t *testing.T) {
		t.Parallel()

		ctx, cancel := context.WithCancel(context.Background())
		cmd := startHang(t, ctx, tooLong, os.Kill, 10*time.Millisecond, "-subsleep=10m", "-probe=1ms")
		cancel()
		err := cmd.Wait()
		t.Logf("stderr:\n%s", cmd.F.Stderr)
		t.Logf("[%d] %v", cmd.F.Process.Pid, err)

		// 该测试应在10毫秒后终止子进程，
		// 留下一个孙辈进程以循环方式写入探针。
		// 子进程应被报告为失败状态，
		// 当标准错误管道关闭时，孙辈进程将退出（或因SIGPIPE信号而终止）。
		if ee := new(*exec.ExitError); !errors.As(err, ee) {
			t.Errorf("Wait error = %v; want %T", err, *ee)
		}
	})

	// If the process exits with status 0 but leaves a child behind writing
	// to its output pipes, Wait should only wait for WaitDelay before
	// closing the pipes and returning.  Wait should return ErrWaitDelay
	// to indicate that the piped output may be incomplete even though the
	// command returned a “success” code.
	t.Run("Exit-hang", func(t *testing.T) {
		t.Parallel()

		cmd := startHang(t, context.Background(), 1*time.Millisecond, nil, 10*time.Millisecond, "-subsleep=10m", "-probe=1ms")
		err := cmd.Wait()
		t.Logf("stderr:\n%s", cmd.F.Stderr)
		t.Logf("[%d] %v", cmd.F.Process.Pid, err)

		// 该子进程应立即退出，
		// 留下孙进程以循环方式写入探测数据。
		// 由于子进程没有 ExitError 需要报告，但我们尚未读取其所有输出，
		// 因此 Wait 函数应返回 ErrWaitDelay。
		if !errors.Is(err, exec2.ErrWaitDelay) {
			t.Errorf("Wait error = %v; want %T", err, exec2.ErrWaitDelay)
		}
	})

	// 如果Cancel函数发送了一个进程可以处理的信号，且该进程在接收到该信号后并未实际退出，而是进行了某种处理，则应在WaitDelay后将其强制终止。
	t.Run("SIGINT-ignored", func(t *testing.T) {
		if runtime.GOOS == "windows" {
			t.Skipf("skipping: os.Interrupt is not implemented on Windows")
		}
		t.Parallel()

		ctx, cancel := context.WithCancel(context.Background())
		cmd := startHang(t, ctx, tooLong, os.Interrupt, 10*time.Millisecond, "-interrupt=false")
		cancel()
		err := cmd.Wait()
		t.Logf("stderr:\n%s", cmd.F.Stderr)
		t.Logf("[%d] %v", cmd.F.Process.Pid, err)

		// 该命令忽略 SIGINT 信号，持续休眠直到被杀死。
		// Wait 方法应返回被杀死进程通常的错误信息。
		if ee := new(*exec.ExitError); !errors.As(err, ee) {
			t.Errorf("Wait error = %v; want %T", err, *ee)
		}
	})

	// 如果进程捕获到了取消信号并以状态0退出，
	// 那么Wait应当返回一个非nil的错误（因为进程是被中断的），
	// 并且它应该是一个上下文错误（因为子进程中没有要报告的错误）。
	t.Run("SIGINT-handled", func(t *testing.T) {
		if runtime.GOOS == "windows" {
			t.Skipf("skipping: os.Interrupt is not implemented on Windows")
		}
		t.Parallel()

		ctx, cancel := context.WithCancel(context.Background())
		cmd := startHang(t, ctx, tooLong, os.Interrupt, 0, "-interrupt=true")
		cancel()
		err := cmd.Wait()
		t.Logf("stderr:\n%s", cmd.F.Stderr)
		t.Logf("[%d] %v", cmd.F.Process.Pid, err)

		if !errors.Is(err, ctx.Err()) {
			t.Errorf("Wait error = %v; want %v", err, ctx.Err())
		}
		if ps := cmd.F.ProcessState; !ps.Exited() {
			t.Errorf("cmd did not exit: %v", ps)
		} else if code := ps.ExitCode(); code != 0 {
			t.Errorf("cmd.ProcessState.ExitCode() = %v; want 0", code)
		}
	})

	// 如果Cancel函数发送SIGQUIT信号，它应以通常的方式进行处理：Go程序应输出其goroutine信息并以非成功状态退出。（我们预期SIGQUIT在实际使用中会成为一种常见模式。）
	t.Run("SIGQUIT", func(t *testing.T) {
		if quitSignal == nil {
			t.Skipf("skipping: SIGQUIT is not supported on %v", runtime.GOOS)
		}
		t.Parallel()

		ctx, cancel := context.WithCancel(context.Background())
		cmd := startHang(t, ctx, tooLong, quitSignal, 0)
		cancel()
		err := cmd.Wait()
		t.Logf("stderr:\n%s", cmd.F.Stderr)
		t.Logf("[%d] %v", cmd.F.Process.Pid, err)

		if ee := new(*exec.ExitError); !errors.As(err, ee) {
			t.Errorf("Wait error = %v; want %v", err, ctx.Err())
		}

		if ps := cmd.F.ProcessState; !ps.Exited() {
			t.Errorf("cmd did not exit: %v", ps)
		} else if code := ps.ExitCode(); code != 2 {
			// 默认的os/signal处理器以代码2退出
			t.Errorf("cmd.ProcessState.ExitCode() = %v; want 2", code)
		}

		if !strings.Contains(fmt.Sprint(cmd.F.Stderr), "\n\ngoroutine ") {
			t.Errorf("cmd.Stderr does not contain a goroutine dump")
		}
	})
}

func TestCancelErrors(t *testing.T) {
	t.Parallel()

	// 如果Cancel返回非ErrProcessDone的错误且进程成功退出，Wait应当封装Cancel的错误
	t.Run("success after error", func(t *testing.T) {
		t.Parallel()

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		cmd := helperCommandContext(t, ctx, "pipetest")
		stdin, err := cmd.StdinPipe()
		if err != nil {
			t.Fatal(err)
		}

		errArbitrary := errors.New("arbitrary error")
		cmd.F.Cancel = func() error {
			stdin.Close()
			t.Logf("Cancel returning %v", errArbitrary)
			return errArbitrary
		}
		if err := cmd.Start(); err != nil {
			t.Fatal(err)
		}
		cancel()

		err = cmd.Wait()
		t.Logf("[%d] %v", cmd.F.Process.Pid, err)
		if !errors.Is(err, errArbitrary) || err == errArbitrary {
			t.Errorf("Wait error = %v; want an error wrapping %v", err, errArbitrary)
		}
	})

	// If Cancel returns an error equivalent to ErrProcessDone,
	// Wait should ignore that error. (ErrProcessDone indicates that the
	// process was already done before we tried to interrupt it — maybe we
	// just didn't notice because Wait hadn't been called yet.)
	t.Run("success after ErrProcessDone", func(t *testing.T) {
		t.Parallel()

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		cmd := helperCommandContext(t, ctx, "pipetest")
		stdin, err := cmd.StdinPipe()
		if err != nil {
			t.Fatal(err)
		}

		stdout, err := cmd.StdoutPipe()
		if err != nil {
			t.Fatal(err)
		}

		// 我们特意让Cancel与进程退出进行竞争，但确保进程赢得此竞争（并在Cancel中返回ErrProcessDone以报告此情况）。
		interruptCalled := make(chan struct{})
		done := make(chan struct{})
		cmd.F.Cancel = func() error {
			close(interruptCalled)
			<-done
			t.Logf("Cancel returning an error wrapping ErrProcessDone")
			return fmt.Errorf("%w: stdout closed", os.ErrProcessDone)
		}

		if err := cmd.Start(); err != nil {
			t.Fatal(err)
		}

		cancel()
		<-interruptCalled
		stdin.Close()
		io.Copy(io.Discard, stdout) // 当进程退出时到达EOF
		close(done)

		err = cmd.Wait()
		t.Logf("[%d] %v", cmd.F.Process.Pid, err)
		if err != nil {
			t.Errorf("Wait error = %v; want nil", err)
		}
	})

	// 如果 Cancel 返回错误，并且在经过 WaitDelay 后进程被杀掉，那么 Wait 应该报告常规的 SIGKILL ExitError，而不是来自 Cancel 的错误。
	t.Run("killed after error", func(t *testing.T) {
		t.Parallel()

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		cmd := helperCommandContext(t, ctx, "pipetest")
		stdin, err := cmd.StdinPipe()
		if err != nil {
			t.Fatal(err)
		}
		defer stdin.Close()

		errArbitrary := errors.New("arbitrary error")
		var interruptCalled atomic.Bool
		cmd.F.Cancel = func() error {
			t.Logf("Cancel called")
			interruptCalled.Store(true)
			return errArbitrary
		}
		cmd.F.WaitDelay = 1 * time.Millisecond
		if err := cmd.Start(); err != nil {
			t.Fatal(err)
		}
		cancel()

		err = cmd.Wait()
		t.Logf("[%d] %v", cmd.F.Process.Pid, err)

		// 确保 Cancel 实际上有机会返回该错误
		if !interruptCalled.Load() {
			t.Errorf("Cancel was not called when the context was canceled")
		}

		// 本测试应在1毫秒后终止子进程，
		// 为最大程度地与现有exec.CommandContext用法兼容，产生的错误应为未经额外封装的exec.ExitError。
		if ee, ok := err.(*exec.ExitError); !ok {
			t.Errorf("Wait error = %v; want %T", err, *ee)
		}
	})

	// 如果 Cancel 返回 ErrProcessDone，但实际上进程并未完成（仍需被杀死），则 Wait 应报告常规的 SIGKILL ExitError，而非来自 Cancel 的错误。
	t.Run("killed after spurious ErrProcessDone", func(t *testing.T) {
		t.Parallel()

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		cmd := helperCommandContext(t, ctx, "pipetest")
		stdin, err := cmd.StdinPipe()
		if err != nil {
			t.Fatal(err)
		}
		defer stdin.Close()

		var interruptCalled atomic.Bool
		cmd.F.Cancel = func() error {
			t.Logf("Cancel returning an error wrapping ErrProcessDone")
			interruptCalled.Store(true)
			return fmt.Errorf("%w: stdout closed", os.ErrProcessDone)
		}
		cmd.F.WaitDelay = 1 * time.Millisecond
		if err := cmd.Start(); err != nil {
			t.Fatal(err)
		}
		cancel()

		err = cmd.Wait()
		t.Logf("[%d] %v", cmd.F.Process.Pid, err)

		// 确保 Cancel 实际上有机会返回该错误
		if !interruptCalled.Load() {
			t.Errorf("Cancel was not called when the context was canceled")
		}

		// 本测试应在1毫秒后终止子进程，
		// 为最大程度地与现有exec.CommandContext用法兼容，产生的错误应为未经额外封装的exec.ExitError。
		if ee, ok := err.(*exec.ExitError); !ok {
			t.Errorf("Wait error of type %T; want %T", err, ee)
		}
	})

	// 如果Cancel返回错误并且进程以非正常退出码结束，则应优先考虑进程错误而非Cancel错误。
	t.Run("nonzero exit after error", func(t *testing.T) {
		t.Parallel()

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		cmd := helperCommandContext(t, ctx, "stderrfail")
		stderr, err := cmd.StderrPipe()
		if err != nil {
			t.Fatal(err)
		}

		errArbitrary := errors.New("arbitrary error")
		interrupted := make(chan struct{})
		cmd.F.Cancel = func() error {
			close(interrupted)
			return errArbitrary
		}
		if err := cmd.Start(); err != nil {
			t.Fatal(err)
		}
		cancel()
		<-interrupted
		io.Copy(io.Discard, stderr)

		err = cmd.Wait()
		t.Logf("[%d] %v", cmd.F.Process.Pid, err)

		if ee, ok := err.(*exec.ExitError); !ok || ee.F.ProcessState.ExitCode() != 1 {
			t.Errorf("Wait error = %v; want exit status 1", err)
		}
	})
}

// TestConcurrentExec 是针对 https://go.dev/issue/61080 的回归测试。
//
// 在 Darwin 系统上，有时并发启动多个子进程会导致程序挂起。
// （在 gomote 上运行时，仅经过几次迭代后，在 -count=100 选项下该测试即出现挂起。）
func TestConcurrentExec(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	// 本测试将启动 nHangs 个挂起的子进程，它们从 stdin 读取数据时会挂起，
	// 同时启动 nExits 个立即退出的子进程。
	//
	// 当存在 #61080 问题时，一个长期运行的“挂起”子进程偶尔会继承自“退出”子进程的 fork/exec 状态管道，
	// 导致父进程（期望几乎立刻在该管道上看到 EOF）意外地阻塞在从管道读取的操作上。
	var (
		nHangs       = runtime.GOMAXPROCS(0)
		nExits       = runtime.GOMAXPROCS(0)
		hangs, exits sync.WaitGroup
	)
	hangs.Add(nHangs)
	exits.Add(nExits)

	// ready表示当goroutine已经尽可能地完成了创建子进程的准备工作。虽然严格来说这不是测试所必需的，但通过增加“hang”和“exit”goroutine对syscall.StartProcess调用发生重叠的可能性，有助于提高问题复现率。
	var ready sync.WaitGroup
	ready.Add(nHangs + nExits)

	for i := 0; i < nHangs; i++ {
		go func() {
			defer hangs.Done()

			cmd := helperCommandContext(t, ctx, "pipetest")
			stdin, err := cmd.StdinPipe()
			if err != nil {
				ready.Done()
				t.Error(err)
				return
			}
			cmd.F.Cancel = stdin.Close
			ready.Done()

			ready.Wait()
			if err := cmd.Start(); err != nil {
				if !errors.Is(err, context.Canceled) {
					t.Error(err)
				}
				return
			}

			cmd.Wait()
		}()
	}

	for i := 0; i < nExits; i++ {
		go func() {
			defer exits.Done()

			cmd := helperCommandContext(t, ctx, "exit", "0")
			ready.Done()

			ready.Wait()
			if err := cmd.Run(); err != nil {
				t.Error(err)
			}
		}()
	}

	exits.Wait()
	cancel()
	hangs.Wait()
}

// TestPathRace 测试 [Cmd.String] 可以与 [Cmd.Start] 并发调用
func TestPathRace(t *testing.T) {
	cmd := helperCommand(t, "exit", "0")

	done := make(chan struct{})
	go func() {
		out, err := cmd.CombinedOutput()
		t.Logf("%v: %v\n%s", cmd, err, out)
		close(done)
	}()

	t.Logf("running in background: %v", cmd)
	<-done
}
