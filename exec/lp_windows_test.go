// 版权所有 2013 The Go Authors。保留所有权利。
// 使用本源代码受 BSD 风格许可协议约束，
// 该协议可在 LICENSE 文件中找到。

// 使用外部测试以避免 os/exec -> internal/testenv -> os/exec 之间的循环依赖关系

package exec_test

import (
	"errors"
	"fmt"
	"github.com/888go/gosdk/exec"
	"github.com/888go/gosdk/exec/internal/testenv"
	exec2 "os/exec"

	"io"
	"io/fs"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"
)

func init() {
	registerHelperCommand("printpath", cmdPrintPath)
}

func cmdPrintPath(args ...string) {
	exe, err := os.Executable()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Executable: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(exe)
}

// makePATH 函数返回一个引用给定目录（相对于根目录）的 PATH 变量。
//
// 空字符串将导致生成一个空条目。
// 以 . 开头的路径将作为相对条目保留。
func makePATH(root string, dirs []string) string {
	paths := make([]string, 0, len(dirs))
	for _, d := range dirs {
		switch {
		case d == "":
			paths = append(paths, "")
		case d == "." || (len(d) >= 2 && d[0] == '.' && os.IsPathSeparator(d[1])):
			paths = append(paths, filepath.Clean(d))
		default:
			paths = append(paths, filepath.Join(root, d))
		}
	}
	return strings.Join(paths, string(os.PathListSeparator))
}

// installProgs 在多个目标路径下创建可执行文件（或指向可执行文件的符号链接）。它使用 root 作为所有目标文件的前缀。
func installProgs(t *testing.T, root string, files []string) {
	for _, f := range files {
		dstPath := filepath.Join(root, f)

		dir := filepath.Dir(dstPath)
		if err := os.MkdirAll(dir, 0755); err != nil {
			t.Fatal(err)
		}

		if os.IsPathSeparator(f[len(f)-1]) {
			continue // 仅目录和 PATH 条目
		}
		if strings.EqualFold(filepath.Ext(f), ".bat") {
			installBat(t, dstPath)
		} else {
			installExe(t, dstPath)
		}
	}
}

// installExe 在给定位置安装测试可执行程序的一个副本，必要时创建目录。
//
// （我们使用副本而非仅使用符号链接，以确保无论其具体实现如何，os.Executable 始终报告一个明确的路径。）
func installExe(t *testing.T, dstPath string) {
	src, err := os.Open(exePath(t))
	if err != nil {
		t.Fatal(err)
	}
	defer src.Close()

	dst, err := os.OpenFile(dstPath, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0o777)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := dst.Close(); err != nil {
			t.Fatal(err)
		}
	}()

	_, err = io.Copy(dst, src)
	if err != nil {
		t.Fatal(err)
	}
}

// installBat 在dst位置创建一个批处理文件，当运行时会打印其自身路径
func installBat(t *testing.T, dstPath string) {
	dst, err := os.OpenFile(dstPath, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0o777)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := dst.Close(); err != nil {
			t.Fatal(err)
		}
	}()

	if _, err := fmt.Fprintf(dst, "@echo %s\r\n", dstPath); err != nil {
		t.Fatal(err)
	}
}

type lookPathTest struct {
	name            string
	PATHEXT         string // empty to use default
	files           []string
	PATH            []string // 若为nil，则使用文件的所有上级目录
	searchFor       string
	want            string
	wantErr         error
	skipCmdExeCheck bool // 如果为真，不检查want与cmd.exe行为的一致性
}

var lookPathTests = []lookPathTest{
	{
		name:      "first match",
		files:     []string{`p1\a.exe`, `p2\a.exe`, `p2\a`},
		searchFor: `a`,
		want:      `p1\a.exe`,
	},
	{
		name:      "dirs with extensions",
		files:     []string{`p1.dir\a`, `p2.dir\a.exe`},
		searchFor: `a`,
		want:      `p2.dir\a.exe`,
	},
	{
		name:      "first with extension",
		files:     []string{`p1\a.exe`, `p2\a.exe`},
		searchFor: `a.exe`,
		want:      `p1\a.exe`,
	},
	{
		name:      "specific name",
		files:     []string{`p1\a.exe`, `p2\b.exe`},
		searchFor: `b`,
		want:      `p2\b.exe`,
	},
	{
		name:      "no extension",
		files:     []string{`p1\b`, `p2\a`},
		searchFor: `a`,
		wantErr:   exec2.ErrNotFound,
	},
	{
		name:      "directory, no extension",
		files:     []string{`p1\a.exe`, `p2\a.exe`},
		searchFor: `p2\a`,
		want:      `p2\a.exe`,
	},
	{
		name:      "no match",
		files:     []string{`p1\a.exe`, `p2\a.exe`},
		searchFor: `b`,
		wantErr:   exec2.ErrNotFound,
	},
	{
		name:      "no match with dir",
		files:     []string{`p1\b.exe`, `p2\a.exe`},
		searchFor: `p2\b`,
		wantErr:   exec2.ErrNotFound,
	},
	{
		name:      "extensionless file in CWD ignored",
		files:     []string{`a`, `p1\a.exe`, `p2\a.exe`},
		searchFor: `a`,
		want:      `p1\a.exe`,
	},
	{
		name:      "extensionless file in PATH ignored",
		files:     []string{`p1\a`, `p2\a.exe`},
		searchFor: `a`,
		want:      `p2\a.exe`,
	},
	{
		name:      "specific extension",
		files:     []string{`p1\a.exe`, `p2\a.bat`},
		searchFor: `a.bat`,
		want:      `p2\a.bat`,
	},
	{
		name:      "mismatched extension",
		files:     []string{`p1\a.exe`, `p2\a.exe`},
		searchFor: `a.com`,
		wantErr:   exec2.ErrNotFound,
	},
	{
		name:      "doubled extension",
		files:     []string{`p1\a.exe.exe`},
		searchFor: `a.exe`,
		want:      `p1\a.exe.exe`,
	},
	{
		name:      "extension not in PATHEXT",
		PATHEXT:   `.COM;.BAT`,
		files:     []string{`p1\a.exe`, `p2\a.exe`},
		searchFor: `a.exe`,
		want:      `p1\a.exe`,
	},
	{
		name:      "first allowed by PATHEXT",
		PATHEXT:   `.COM;.EXE`,
		files:     []string{`p1\a.bat`, `p2\a.exe`},
		searchFor: `a`,
		want:      `p2\a.exe`,
	},
	{
		name:      "first directory containing a PATHEXT match",
		PATHEXT:   `.COM;.EXE;.BAT`,
		files:     []string{`p1\a.bat`, `p2\a.exe`},
		searchFor: `a`,
		want:      `p1\a.bat`,
	},
	{
		name:      "first PATHEXT entry",
		PATHEXT:   `.COM;.EXE;.BAT`,
		files:     []string{`p1\a.bat`, `p1\a.exe`, `p2\a.bat`, `p2\a.exe`},
		searchFor: `a`,
		want:      `p1\a.exe`,
	},
	{
		name:      "ignore dir with PATHEXT extension",
		files:     []string{`a.exe\`},
		searchFor: `a`,
		wantErr:   exec2.ErrNotFound,
	},
	{
		name:      "ignore empty PATH entry",
		files:     []string{`a.bat`, `p\a.bat`},
		PATH:      []string{`p`},
		searchFor: `a`,
		want:      `p\a.bat`,
		// 如果cmd.exe版本过旧，可能不支持NoDefaultCurrentDirectoryInExePath选项，
		// 因此跳过该检查。
		skipCmdExeCheck: true,
	},
	{
		name:      "return ErrDot if found by a different absolute path",
		files:     []string{`p1\a.bat`, `p2\a.bat`},
		PATH:      []string{`.\p1`, `p2`},
		searchFor: `a`,
		want:      `p1\a.bat`,
		wantErr:   exec2.ErrDot,
	},
	{
		name:      "suppress ErrDot if also found in absolute path",
		files:     []string{`p1\a.bat`, `p2\a.bat`},
		PATH:      []string{`.\p1`, `p1`, `p2`},
		searchFor: `a`,
		want:      `p1\a.bat`,
	},
}

func TestLookPathWindows(t *testing.T) {
	// 非并行：使用了 Chdir 和 Setenv。

	// 我们在这里使用“printpath”命令模式来测试exec.Command，
	// 因此我们不会调用helperCommand来解析它。
	// 这可能导致它看似未被使用。
	maySkipHelperCommand("printpath")

	// 在开始之前，找出cmd.exe的绝对路径。
	// 在非短模式下，我们将用它来校验测试“want”字段的实际情况。
	cmdExe, err := exec.LookPath("cmd")
	if err != nil {
		t.Fatal(err)
	}

	for _, tt := range lookPathTests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.want == "" && tt.wantErr == nil {
				t.Fatalf("test must specify either want or wantErr")
			}

			root := t.TempDir()
			installProgs(t, root, tt.files)

			if tt.PATHEXT != "" {
				t.Setenv("PATHEXT", tt.PATHEXT)
				t.Logf("set PATHEXT=%s", tt.PATHEXT)
			}

			var pathVar string
			if tt.PATH == nil {
				paths := make([]string, 0, len(tt.files))
				for _, f := range tt.files {
					dir := filepath.Join(root, filepath.Dir(f))
					if !slices.Contains(paths, dir) {
						paths = append(paths, dir)
					}
				}
				pathVar = strings.Join(paths, string(os.PathListSeparator))
			} else {
				pathVar = makePATH(root, tt.PATH)
			}
			t.Setenv("PATH", pathVar)
			t.Logf("set PATH=%s", pathVar)

			chdir(t, root)

			if !testing.Short() && !(tt.skipCmdExeCheck || errors.Is(tt.wantErr, exec2.ErrDot)) {
				// 检查 cmd.exe（作为我们的基准真相来源）是否认同我们的测试用例是正确的。
				cmd := testenv.Command(t, cmdExe, "/c", tt.searchFor, "printpath")
				out, err := cmd.Output()
				if err == nil {
					gotAbs := strings.TrimSpace(string(out))
					wantAbs := ""
					if tt.want != "" {
						wantAbs = filepath.Join(root, tt.want)
					}
					if gotAbs != wantAbs {
						// cmd.exe 不同意。可能是测试用例有误？
						t.Fatalf("%v\n\tresolved to %s\n\twant %s", cmd, gotAbs, wantAbs)
					}
				} else if tt.wantErr == nil {
					if ee, ok := err.(*exec.ExitError); ok && len(ee.F.Stderr) > 0 {
						t.Fatalf("%v: %v\n%s", cmd, err, ee.F.Stderr)
					}
					t.Fatalf("%v: %v", cmd, err)
				}
			}

			got, err := exec.LookPath(tt.searchFor)
			if filepath.IsAbs(got) {
				got, err = filepath.Rel(root, got)
				if err != nil {
					t.Fatal(err)
				}
			}
			if got != tt.want {
				t.Errorf("LookPath(%#q) = %#q; want %#q", tt.searchFor, got, tt.want)
			}
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("LookPath(%#q): %v; want %v", tt.searchFor, err, tt.wantErr)
			}
		})
	}
}

type commandTest struct {
	name       string
	PATH       []string
	files      []string
	dir        string
	arg0       string
	want       string
	wantPath   string // 解析得到的c.Path，如果与want不同
	wantErrDot bool
	wantRunErr error
}

var commandTests = []commandTest{
	// 测试不带斜杠的命令，如 `a.exe`
	{
		name:       "current directory",
		files:      []string{`a.exe`},
		PATH:       []string{"."},
		arg0:       `a.exe`,
		want:       `a.exe`,
		wantErrDot: true,
	},
	{
		name:       "with extra PATH",
		files:      []string{`a.exe`, `p\a.exe`, `p2\a.exe`},
		PATH:       []string{".", "p2", "p"},
		arg0:       `a.exe`,
		want:       `a.exe`,
		wantErrDot: true,
	},
	{
		name:       "with extra PATH and no extension",
		files:      []string{`a.exe`, `p\a.exe`, `p2\a.exe`},
		PATH:       []string{".", "p2", "p"},
		arg0:       `a`,
		want:       `a.exe`,
		wantErrDot: true,
	},
	// 测试使用反斜杠的命令，如 `.\a.exe`
	{
		name:  "with dir",
		files: []string{`p\a.exe`},
		PATH:  []string{"."},
		arg0:  `p\a.exe`,
		want:  `p\a.exe`,
	},
	{
		name:  "with explicit dot",
		files: []string{`p\a.exe`},
		PATH:  []string{"."},
		arg0:  `.\p\a.exe`,
		want:  `p\a.exe`,
	},
	{
		name:  "with irrelevant PATH",
		files: []string{`p\a.exe`, `p2\a.exe`},
		PATH:  []string{".", "p2"},
		arg0:  `p\a.exe`,
		want:  `p\a.exe`,
	},
	{
		name:  "with slash and no extension",
		files: []string{`p\a.exe`, `p2\a.exe`},
		PATH:  []string{".", "p2"},
		arg0:  `p\a`,
		want:  `p\a.exe`,
	},
	// 测试命令，如`a.exe`，其中已设置c.Dir
	{
		// 不应在p中找到a.exe，因为当Command调用LookPath(`a.exe`)时（在设置Dir之前），该操作将失败，且该错误是持久的。
		name:       "not found before Dir",
		files:      []string{`p\a.exe`},
		PATH:       []string{"."},
		dir:        `p`,
		arg0:       `a.exe`,
		want:       `p\a.exe`,
		wantRunErr: exec2.ErrNotFound,
	},
	{
		// LookPath(`a.exe`) 会解析为 `.\a.exe`，但将该路径前缀为目录 `p` 即 `p\a.exe` 时，将指向一个不存在的文件
		name:       "resolved before Dir",
		files:      []string{`a.exe`, `p\not_important_file`},
		PATH:       []string{"."},
		dir:        `p`,
		arg0:       `a.exe`,
		want:       `a.exe`,
		wantErrDot: true,
		wantRunErr: fs.ErrNotExist,
	},
	{
		// 类似于上面的做法，但通过将文件安装到所引用的目标位置（这样LookPath(`a.exe`)仍会找到`.\a.exe`，但我们成功执行了`p\a.exe`），使测试成功
		name:       "relative to Dir",
		files:      []string{`a.exe`, `p\a.exe`},
		PATH:       []string{"."},
		dir:        `p`,
		arg0:       `a.exe`,
		want:       `p\a.exe`,
		wantErrDot: true,
	},
	{
		// 与上面类似，但添加PATH以尝试打破测试
		name:       "relative to Dir with extra PATH",
		files:      []string{`a.exe`, `p\a.exe`, `p2\a.exe`},
		PATH:       []string{".", "p2", "p"},
		dir:        `p`,
		arg0:       `a.exe`,
		want:       `p\a.exe`,
		wantErrDot: true,
	},
	{
		// 与上述相同，但使用 "a" 代替 "a.exe" 作为命令
		name:       "relative to Dir with extra PATH and no extension",
		files:      []string{`a.exe`, `p\a.exe`, `p2\a.exe`},
		PATH:       []string{".", "p2", "p"},
		dir:        `p`,
		arg0:       `a`,
		want:       `p\a.exe`,
		wantErrDot: true,
	},
	{
		// 查找 `a.exe`，忽略 Dir，因为 Command 在设置 Dir 之前通过 LookPath 解析完整路径。
		name:  "from PATH with no match in Dir",
		files: []string{`p\a.exe`, `p2\a.exe`},
		PATH:  []string{".", "p2", "p"},
		dir:   `p`,
		arg0:  `a.exe`,
		want:  `p2\a.exe`,
	},
	// 测试命令，如 `.\a.exe`，其中已设置 c.Dir
	{
		// 当命令为路径形式时，如 ".\a.exe"，应使用 dir
		name:  "relative to Dir with explicit dot",
		files: []string{`p\a.exe`},
		PATH:  []string{"."},
		dir:   `p`,
		arg0:  `.\a.exe`,
		want:  `p\a.exe`,
	},
	{
		// 与上面类似，但尝试添加 PATH 以使其失效
		name:  "relative to Dir with dot and extra PATH",
		files: []string{`p\a.exe`, `p2\a.exe`},
		PATH:  []string{".", "p2"},
		dir:   `p`,
		arg0:  `.\a.exe`,
		want:  `p\a.exe`,
	},
	{
		// 在设置 Dir 之前，LookPath(".\a") 将会失败，且该错误是持久的。
		name:  "relative to Dir with dot and extra PATH and no extension",
		files: []string{`p\a.exe`, `p2\a.exe`},
		PATH:  []string{".", "p2"},
		dir:   `p`,
		arg0:  `.\a`,
		want:  `p\a.exe`,
	},
	{
		// 在设置 Dir 之前，LookPath(".\a") 将会失败，且该错误是持久的。
		name:  "relative to Dir with different extension",
		files: []string{`a.exe`, `p\a.bat`},
		PATH:  []string{"."},
		dir:   `p`,
		arg0:  `.\a`,
		want:  `p\a.bat`,
	},
}

func TestCommand(t *testing.T) {
	// 非并行：使用了 Chdir 和 Setenv。

	// 我们在这里使用“printpath”命令模式来测试exec.Command，
	// 因此我们不会调用helperCommand来解析它。
	// 这可能导致它看似未被使用。
	maySkipHelperCommand("printpath")

	for _, tt := range commandTests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.PATH == nil {
				t.Fatalf("test must specify PATH")
			}

			root := t.TempDir()
			installProgs(t, root, tt.files)

			pathVar := makePATH(root, tt.PATH)
			t.Setenv("PATH", pathVar)
			t.Logf("set PATH=%s", pathVar)

			chdir(t, root)

			cmd := exec.Command(tt.arg0, "printpath")
			cmd.F.Dir = filepath.Join(root, tt.dir)
			if tt.wantErrDot {
				if errors.Is(cmd.F.Err, exec2.ErrDot) {
					cmd.F.Err = nil
				} else {
					t.Fatalf("cmd.Err = %v; want ErrDot", cmd.F.Err)
				}
			}

			out, err := cmd.Output()
			if err != nil {
				if ee, ok := err.(*exec.ExitError); ok && len(ee.F.Stderr) > 0 {
					t.Logf("%v: %v\n%s", cmd, err, ee.F.Stderr)
				} else {
					t.Logf("%v: %v", cmd, err)
				}
				if !errors.Is(err, tt.wantRunErr) {
					t.Errorf("want %v", tt.wantRunErr)
				}
				return
			}

			got := strings.TrimSpace(string(out))
			if filepath.IsAbs(got) {
				got, err = filepath.Rel(root, got)
				if err != nil {
					t.Fatal(err)
				}
			}
			if got != tt.want {
				t.Errorf("\nran  %#q\nwant %#q", got, tt.want)
			}

			gotPath := cmd.F.Path
			wantPath := tt.wantPath
			if wantPath == "" {
				if strings.Contains(tt.arg0, `\`) {
					wantPath = tt.arg0
				} else if tt.wantErrDot {
					wantPath = strings.TrimPrefix(tt.want, tt.dir+`\`)
				} else {
					wantPath = filepath.Join(root, tt.want)
				}
			}
			if gotPath != wantPath {
				t.Errorf("\ncmd.Path = %#q\nwant       %#q", gotPath, wantPath)
			}
		})
	}
}

func TestAbsCommandWithDoubledExtension(t *testing.T) {
	t.Parallel()

	// 我们期望 ".com" 总是包含在 PATHEXT 中，但它也可能出现在 Go 包的导入路径中。如果它位于导入路径的根部，则生成的可执行文件可能被命名为 "example.com.exe"。
	//
	// 由于 "example.com" 看起来像一个合理的可执行文件名，因此 exec.Command 尝试直接运行它而不重新解析可能是可以接受的。
	// 但是，exec.LookPath 应该更努力地尝试确定它。

	comPath := filepath.Join(t.TempDir(), "example.com")
	batPath := comPath + ".bat"
	installBat(t, batPath)

	cmd := exec.Command(comPath)
	out, err := cmd.CombinedOutput()
	t.Logf("%v: %v\n%s", cmd, err, out)
	if !errors.Is(err, fs.ErrNotExist) {
		t.Errorf("Command(%#q).Run: %v\nwant fs.ErrNotExist", comPath, err)
	}

	resolved, err := exec.LookPath(comPath)
	if err != nil || resolved != batPath {
		t.Fatalf("LookPath(%#q) = %v, %v; want %#q, <nil>", comPath, resolved, err, batPath)
	}
}
