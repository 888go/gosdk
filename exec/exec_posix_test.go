// 版权所有 ? 2017 The Go Authors。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

//---build---//go:build unix

package exec_test

import (
	"fmt"
	"internal/testenv"
	"io"
	"os"
	"os/user"
	"path/filepath"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"syscall"
	"testing"
	"time"
)

func init() {
	registerHelperCommand("pwd", cmdPwd)
}

func cmdPwd(...string) {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(pwd)
}

func TestCredentialNoSetGroups(t *testing.T) {
	if runtime.GOOS == "android" {
		maySkipHelperCommand("echo")
		t.Skip("unsupported on Android")
	}
	t.Parallel()

	u, err := user.Current()
	if err != nil {
		t.Fatalf("error getting current user: %v", err)
	}

	uid, err := strconv.Atoi(u.Uid)
	if err != nil {
		t.Fatalf("error converting Uid=%s to integer: %v", u.Uid, err)
	}

	gid, err := strconv.Atoi(u.Gid)
	if err != nil {
		t.Fatalf("error converting Gid=%s to integer: %v", u.Gid, err)
	}

	// 如果NoSetGroups 为真，则不会调用 setgroups，且 cmd.Run 应该能成功执行
	cmd := helperCommand(t, "echo", "foo")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Credential: &syscall.Credential{
			Uid:         uint32(uid),
			Gid:         uint32(gid),
			NoSetGroups: true,
		},
	}

	if err = cmd.Run(); err != nil {
		t.Errorf("Failed to run command: %v", err)
	}
}

// 为解决第19314号问题：确保SIGSTOP信号不会导致进程看似已完成。
func TestWaitid(t *testing.T) {
	t.Parallel()

	cmd := helperCommand(t, "pipetest")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		t.Fatal(err)
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		t.Fatal(err)
	}
	if err := cmd.Start(); err != nil {
		t.Fatal(err)
	}

	// 等待子进程启动，并注册任何信号处理器。
	const msg = "O:ping\n"
	if _, err := io.WriteString(stdin, msg); err != nil {
		t.Fatal(err)
	}
	buf := make([]byte, len(msg))
	if _, err := io.ReadFull(stdout, buf); err != nil {
		t.Fatal(err)
	}
	// 现在保持管道打开，以便进程会挂起直到我们关闭标准输入。

	if err := cmd.Process.Signal(syscall.SIGSTOP); err != nil {
		cmd.Process.Kill()
		t.Fatal(err)
	}

	ch := make(chan error)
	go func() {
		ch <- cmd.Wait()
	}()

// 给`Wait`一点时间来阻塞等待进程。
// （这只是为了触发bug提供一些时间；对于测试通过来说，这并不必要。）
	if testing.Short() {
		time.Sleep(1 * time.Millisecond)
	} else {
		time.Sleep(10 * time.Millisecond)
	}

// 此调用Signal应成功，因为进程仍然存在。
// （在修复#19314问题之前，这将因os.ErrProcessDone或等效错误而失败。）
	if err := cmd.Process.Signal(syscall.SIGCONT); err != nil {
		t.Error(err)
		syscall.Kill(cmd.Process.Pid, syscall.SIGCONT)
	}

// SIGCONT 信号应能使进程唤醒，注意到stdin已关闭，并成功退出。
	stdin.Close()
	err = <-ch
	if err != nil {
		t.Fatal(err)
	}
}

// https://go.dev/issue/50599：若未明确设置 Env，当设置 Dir 时应隐式更新 PWD 至正确路径，并且 Environ 应列出已更新的值。
func TestImplicitPWD(t *testing.T) {
	t.Parallel()

	cwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	cases := []struct {
		name string
		dir  string
		want string
	}{
		{"empty", "", cwd},
		{"dot", ".", cwd},
		{"dotdot", "..", filepath.Dir(cwd)},
		{"PWD", cwd, cwd},
		{"PWDdotdot", cwd + string(filepath.Separator) + "..", filepath.Dir(cwd)},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			cmd := helperCommand(t, "pwd")
			if cmd.Env != nil {
				t.Fatalf("test requires helperCommand not to set Env field")
			}
			cmd.Dir = tc.dir

			var pwds []string
			for _, kv := range cmd.Environ() {
				if strings.HasPrefix(kv, "PWD=") {
					pwds = append(pwds, strings.TrimPrefix(kv, "PWD="))
				}
			}

			wantPWDs := []string{tc.want}
			if tc.dir == "" {
				if _, ok := os.LookupEnv("PWD"); !ok {
					wantPWDs = nil
				}
			}
			if !reflect.DeepEqual(pwds, wantPWDs) {
				t.Errorf("PWD entries in cmd.Environ():\n\t%s\nwant:\n\t%s", strings.Join(pwds, "\n\t"), strings.Join(wantPWDs, "\n\t"))
			}

			cmd.Stderr = new(strings.Builder)
			out, err := cmd.Output()
			if err != nil {
				t.Fatalf("%v:\n%s", err, cmd.Stderr)
			}
			got := strings.Trim(string(out), "\r\n")
			t.Logf("in\n\t%s\n`pwd` reported\n\t%s", tc.dir, got)
			if got != tc.want {
				t.Errorf("want\n\t%s", tc.want)
			}
		})
	}
}

// 但是，如果已明确设置 cmd.Env，则设置 Dir 不应覆盖它。
// （这会检查为 https://go.dev/issue/50599 实现的代码是否不会破坏可能已显式错配 PWD 变量的现有用户。）
func TestExplicitPWD(t *testing.T) {
	t.Parallel()

	maySkipHelperCommand("pwd")
	testenv.MustHaveSymlink(t)

	cwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	link := filepath.Join(t.TempDir(), "link")
	if err := os.Symlink(cwd, link); err != nil {
		t.Fatal(err)
	}

// 现在link是cwd的另一个同样有效的名称。如果我们将Dir设置为其中一个，而PWD设置为另一个，则子进程应报告PWD版本。
	cases := []struct {
		name string
		dir  string
		pwd  string
	}{
		{name: "original PWD", pwd: cwd},
		{name: "link PWD", pwd: link},
		{name: "in link with original PWD", dir: link, pwd: cwd},
		{name: "in dir with link PWD", dir: cwd, pwd: link},
// 理想情况下，我们还希望测试将 PWD 设置为
// 完全无意义的值（或空字符串）时会发生什么情况，但那样的话，
// 我们就无法得知子进程实际应产生的输出是什么：cwd 本身可能包含
// 测试环境中的 PWD 值所保留的符号链接。
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			cmd := helperCommand(t, "pwd")
// 这里故意颠倒了通常设置 cmd.Dir 然后调用 cmd.Environ 的顺序。在这里，我们**希望** PWD 与 cmd.Dir 不匹配，因此我们并不关心 cmd.Dir 是否反映在 cmd.Environ 中。
			cmd.Env = append(cmd.Environ(), "PWD="+tc.pwd)
			cmd.Dir = tc.dir

			var pwds []string
			for _, kv := range cmd.Environ() {
				if strings.HasPrefix(kv, "PWD=") {
					pwds = append(pwds, strings.TrimPrefix(kv, "PWD="))
				}
			}

			wantPWDs := []string{tc.pwd}
			if !reflect.DeepEqual(pwds, wantPWDs) {
				t.Errorf("PWD entries in cmd.Environ():\n\t%s\nwant:\n\t%s", strings.Join(pwds, "\n\t"), strings.Join(wantPWDs, "\n\t"))
			}

			cmd.Stderr = new(strings.Builder)
			out, err := cmd.Output()
			if err != nil {
				t.Fatalf("%v:\n%s", err, cmd.Stderr)
			}
			got := strings.Trim(string(out), "\r\n")
			t.Logf("in\n\t%s\nwith PWD=%s\nsubprocess os.Getwd() reported\n\t%s", tc.dir, tc.pwd, got)
			if got != tc.pwd {
				t.Errorf("want\n\t%s", tc.pwd)
			}
		})
	}
}
