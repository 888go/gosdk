// 版权所有 ? 2022 Go作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

//2024-04-17 备注,单元测试通不过, 保留单元测试文件为了方便查看使用方法.
package exec_test

import (
	"errors"
	. "github.com/888go/gosdk/exec"
	"github.com/888go/gosdk/exec/internal/testenv"
	exec2 "os/exec"

	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)
 

var pathVar string = func() string {
	if runtime.GOOS == "plan9" {
		return "path"
	}
	return "PATH"
}()

func TestLookPath(t *testing.T) {
	testenv.MustHaveExec(t)
	// 非并行：使用了 Chdir 和 Setenv。

	tmpDir := filepath.Join(t.TempDir(), "testdir")
	if err := os.Mkdir(tmpDir, 0777); err != nil {
		t.Fatal(err)
	}

	executable := "execabs-test"
	if runtime.GOOS == "windows" {
		executable += ".exe"
	}
	if err := os.WriteFile(filepath.Join(tmpDir, executable), []byte{1, 2, 3}, 0777); err != nil {
		t.Fatal(err)
	}
	chdir(t, tmpDir)
	t.Setenv("PWD", tmpDir)
	t.Logf(". is %#q", tmpDir)

	origPath := os.Getenv(pathVar)

	// 向PATH环境变量中添加"."，确保在所有系统上exec.LookPath函数都会在当前目录中查找可执行文件。
	// 同时也尝试用"../testdir"来测试迷惑它。
	for _, errdot := range []string{"1", "0"} {
		t.Run("GODEBUG=execerrdot="+errdot, func(t *testing.T) {
			t.Setenv("GODEBUG", "execerrdot="+errdot+",execwait=2")
			for _, dir := range []string{".", "../testdir"} {
				t.Run(pathVar+"="+dir, func(t *testing.T) {
					t.Setenv(pathVar, dir+string(filepath.ListSeparator)+origPath)
					good := dir + "/execabs-test"
					if found, err := LookPath(good); err != nil || !strings.HasPrefix(found, good) {
						t.Fatalf(`LookPath(%#q) = %#q, %v, want "%s...", nil`, good, found, err, good)
					}
					if runtime.GOOS == "windows" {
						good = dir + `\execabs-test`
						if found, err := LookPath(good); err != nil || !strings.HasPrefix(found, good) {
							t.Fatalf(`LookPath(%#q) = %#q, %v, want "%s...", nil`, good, found, err, good)
						}
					}

					_, err := LookPath("execabs-test")
					if errdot == "1" {
						if err == nil {
							t.Fatalf("LookPath didn't fail when finding a non-relative path")
						} else if !errors.Is(err, exec2.ErrDot) {
							t.Fatalf("LookPath returned unexpected error: want Is ErrDot, got %q", err)
						}
					} else {
						if err != nil {
							t.Fatalf("LookPath failed unexpectedly: %v", err)
						}
					}

					cmd := Command("execabs-test")
					if errdot == "1" {
						if cmd.F.Err == nil {
							t.Fatalf("Command didn't fail when finding a non-relative path")
						} else if !errors.Is(cmd.F.Err, exec2.ErrDot) {
							t.Fatalf("Command returned unexpected error: want Is ErrDot, got %q", cmd.F.Err)
						}
						cmd.F.Err = nil
					} else {
						if cmd.F.Err != nil {
							t.Fatalf("Command failed unexpectedly: %v", err)
						}
					}

					// 清除cmd.Err应允许执行继续进行，
					// 并且由于它不是一个有效的二进制文件，所以应该会失败。
					if err := cmd.Run(); err == nil {
						t.Fatalf("Run did not fail: expected exec error")
					} else if errors.Is(err, exec2.ErrDot) {
						t.Fatalf("Run returned unexpected error ErrDot: want error like ENOEXEC: %q", err)
					}
				})
			}
		})
	}

	// 测试当PATH环境变量中的第一个条目为当前目录的绝对路径时的行为。
	//
	// 在Windows系统上，"." 可能会被隐式地包含在显式的 %PATH% 之前，这取决于进程环境；
	// 详细信息请参阅：https://go.dev/issue/4394。
	//
	// 如果从"."解析出的相对路径与仅通过 %PATH% 中绝对路径解析到的可执行文件相同，LookPath 应返回该路径的绝对版本而非 ErrDot。
	// （参考：https://go.dev/issue/53536。）
	//
	// 若 PATH 环境变量未隐式包含 "."（例如，在Unix平台或在Windows系统中配置为NoDefaultCurrentDirectoryInExePath时），则此查找应当无论如何都能成功，因此即使在这些平台上，将其作为对照测试案例运行也是有用的。
	t.Run(pathVar+"=$PWD", func(t *testing.T) {
		t.Setenv(pathVar, tmpDir+string(filepath.ListSeparator)+origPath)
		good := filepath.Join(tmpDir, "execabs-test")
		if found, err := LookPath(good); err != nil || !strings.HasPrefix(found, good) {
			t.Fatalf(`LookPath(%#q) = %#q, %v, want \"%s...\", nil`, good, found, err, good)
		}

		if found, err := LookPath("execabs-test"); err != nil || !strings.HasPrefix(found, good) {
			t.Fatalf(`LookPath(%#q) = %#q, %v, want \"%s...\", nil`, "execabs-test", found, err, good)
		}

		cmd := Command("execabs-test")
		if cmd.F.Err != nil {
			t.Fatalf("Command(%#q).Err = %v; want nil", "execabs-test", cmd.F.Err)
		}
	})

	t.Run(pathVar+"=$OTHER", func(t *testing.T) {
		// 控制用例：如果当 PATH 为空时，查找返回 ErrDot 错误，那么我们知道 PATH 暗示性地包含了 "."。若不是这样，则在这个测试中我们完全不期望看到 ErrDot 错误（因为路径将被明确无误地视为绝对路径）。
		wantErrDot := false
		t.Setenv(pathVar, "")
		if found, err := LookPath("execabs-test"); errors.Is(err, exec2.ErrDot) {
			wantErrDot = true
		} else if err == nil {
			t.Fatalf(`with PATH='', LookPath(%#q) = %#q; want non-nil error`, "execabs-test", found)
		}

		// 设置PATH环境变量，包含一个明确的目录，该目录下包含一个恰好与"."目录下的可执行文件同名的独立可执行文件。如果"."被隐式包含在内，查找（未限定路径的）可执行文件名将返回ErrDot错误；否则，"."目录下的可执行文件不应产生影响，并且查找应当无歧义地解析到PATH环境变量中指定的目录下的可执行文件。

		dir := t.TempDir()
		executable := "execabs-test"
		if runtime.GOOS == "windows" {
			executable += ".exe"
		}
		if err := os.WriteFile(filepath.Join(dir, executable), []byte{1, 2, 3}, 0777); err != nil {
			t.Fatal(err)
		}
		t.Setenv(pathVar, dir+string(filepath.ListSeparator)+origPath)

		found, err := LookPath("execabs-test")
		if wantErrDot {
			wantFound := filepath.Join(".", executable)
			if found != wantFound || !errors.Is(err, exec2.ErrDot) {
				t.Fatalf(`LookPath(%#q) = %#q, %v, want %#q, Is ErrDot`, "execabs-test", found, err, wantFound)
			}
		} else {
			wantFound := filepath.Join(dir, executable)
			if found != wantFound || err != nil {
				t.Fatalf(`LookPath(%#q) = %#q, %v, want %#q, nil`, "execabs-test", found, err, wantFound)
			}
		}
	})
}
