// 版权所有 2015 The Go 作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

// 包testenv提供了关于Go团队在不同测试环境中可用的功能信息。
// 
// 这是一个内部包，因为这些详细信息是针对Go团队的测试环境（在build.golang.org上）特定的，并非测试的一般原理。
package testenv

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"github.com/888go/gosdk/os/internal/cfg"
	"github.com/888go/gosdk/os/internal/platform"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"testing"
)

// 在init期间保存原始环境，以供检查时使用。测试二进制文件可能在调用HasExec之前修改其环境，以改变其行为（例如模拟命令行工具），而该修改后的环境可能导致环境检查行为异常。
var origEnv = os.Environ()

// Builder 返回运行此测试的构建器的名称（例如，“linux-amd64”或“windows-386-gce”）。
// 如果测试不在构建基础设施上运行，Builder 将返回空字符串。
func Builder() string {
	return os.Getenv("GO_BUILDER_NAME")
}

// HasGoBuild reports whether the current system can build programs with “go build”
// and then run them with os.StartProcess or exec.Command.
func HasGoBuild() bool {
	if os.Getenv("GO_GCFLAGS") != "" {
// 要求每次调用`go`命令的用户都传递`"-gcflags="+os.Getenv("GO_GCFLAGS")`太麻烦了。
// 当前，如果`$GO_GCFLAGS`已设置，我们报告无法运行`go build`。
		return false
	}

	goBuildOnce.Do(func() {
// 要运行`go build`，我们需要能够执行一个`go`命令。
// 我们有点随意地选择执行`go tool -n compile`，因为这也确认了cmd/go可以找到编译器。（在CL 472096之前，
// 我们有时会在测试环境中安装没有实际用于构建东西的cmd/go，而它缺少cmd/compile。）
		cmd := exec.Command("go", "tool", "-n", "compile")
		cmd.Env = origEnv
		out, err := cmd.Output()
		if err != nil {
			goBuildErr = fmt.Errorf("%v: %w", cmd, err)
			return
		}
		out = bytes.TrimSpace(out)
		if len(out) == 0 {
			goBuildErr = fmt.Errorf("%v: no tool reported", cmd)
			return
		}
		if _, err := exec.LookPath(string(out)); err != nil {
			goBuildErr = err
			return
		}

		if platform.MustLinkExternal(runtime.GOOS, runtime.GOARCH, false) {
// 我们可以认为我们总是拥有完整的 Go 工具链可用。
// 然而，此平台要求即使构建纯 Go 程序（包括测试）也需要 C 链接器。测试环境中是否具备？
// （例如在 Android 上，运行测试的设备可能没有安装 C 工具链。）
//
// 如果显式设置了 CC，则假定我们具备。否则，使用 'go env CC' 来确定它将默认使用哪个工具链。
			if os.Getenv("CC") == "" {
				cmd := exec.Command("go", "env", "CC")
				cmd.Env = origEnv
				out, err := cmd.Output()
				if err != nil {
					goBuildErr = fmt.Errorf("%v: %w", cmd, err)
					return
				}
				out = bytes.TrimSpace(out)
				if len(out) == 0 {
					goBuildErr = fmt.Errorf("%v: no CC reported", cmd)
					return
				}
				_, goBuildErr = exec.LookPath(string(out))
			}
		}
	})

	return goBuildErr == nil
}

var (
	goBuildOnce sync.Once
	goBuildErr  error
)

// MustHaveGoBuild checks that the current system can build programs with “go build”
// and then run them with os.StartProcess or exec.Command.
// If not, MustHaveGoBuild calls t.Skip with an explanation.
func MustHaveGoBuild(t testing.TB) {
	if os.Getenv("GO_GCFLAGS") != "" {
		t.Helper()
		t.Skipf("skipping test: 'go build' not compatible with setting $GO_GCFLAGS")
	}
	if !HasGoBuild() {
		t.Helper()
		t.Skipf("skipping test: 'go build' unavailable: %v", goBuildErr)
	}
}

// HasGoRun reports whether the current system can run programs with “go run”.
func HasGoRun() bool {
	// 目前，使用"go run"和使用"go build"是相同的。
	return HasGoBuild()
}

// MustHaveGoRun checks that the current system can run programs with “go run”.
// If not, MustHaveGoRun calls t.Skip with an explanation.
func MustHaveGoRun(t testing.TB) {
	if !HasGoRun() {
		t.Skipf("skipping test: 'go run' not available on %s/%s", runtime.GOOS, runtime.GOARCH)
	}
}

// HasParallelism报告当前系统是否可以并行执行多个线程。
// 在cmd/dist/test.go中有一个此函数的副本。
func HasParallelism() bool {
	switch runtime.GOOS {
	case "js", "wasip1":
		return false
	}
	return true
}

// MustHaveParallelism 检查当前系统是否可以同时执行多个线程。如果不可以，MustHaveParallelism 会调用 t.Skip 并附带一个解释。
func MustHaveParallelism(t testing.TB) {
	if !HasParallelism() {
		t.Skipf("skipping test: no parallelism available on %s/%s", runtime.GOOS, runtime.GOARCH)
	}
}

// GoToolPath 返回 Go 工具的路径。
// 它是对 GoTool 的便捷包装。
// 若工具不可用，GoToolPath 将调用 t.Skip。
// 若工具本应可用但实际不可用，GoToolPath 将调用 t.Fatal。
func GoToolPath(t testing.TB) string {
	MustHaveGoBuild(t)
	path, err := GoTool()
	if err != nil {
		t.Fatal(err)
	}
// 将所有影响Go命令的环境变量添加到测试元数据中。
// 当这些变量改变时，缓存的测试结果将被失效。
// 参见golang.org/issue/32285。
	for _, envVar := range strings.Fields(cfg.KnownEnv) {
		os.Getenv(envVar)
	}
	return path
}

var (
	gorootOnce sync.Once
	gorootPath string
	gorootErr  error
)

func findGOROOT() (string, error) {
	gorootOnce.Do(func() {
		gorootPath = runtime.GOROOT()
		if gorootPath != "" {
// 如果runtime.GOROOT()非空，假设它是有效的。
// 
// (它可能不是：例如，用户可能已明确将GOROOT设置为错误的目录，或者明确设置了GOROOT_FINAL但没有设置GOROOT，并且还没有将树移动到GOROOT_FINAL。但这些情况很少见，如果发生这种情况，用户可以修复他们破坏的东西。)
			return
		}

// runtime.GOROOT 不知道 GOROOT 在哪里（可能是因为测试二进制文件是使用 -trimpath 构建的，或者是因为设置了 GOROOT_FINAL 但没有将树移动到该位置）。
// 
// 由于这是 internal/testenv，我们可以作弊并假设调用者是 GOROOT/src 子目录中某个包的测试。`go test`会在包含被测试包的目录中运行测试。这意味着如果我们开始向上遍历树，最终应该找到 GOROOT/src/go.mod，我们可以报告出那个目录的父目录。
// 
// 值得注意的是，即使我们不能作为子进程运行 `go env GOROOT`，这仍然有效。

		cwd, err := os.Getwd()
		if err != nil {
			gorootErr = fmt.Errorf("finding GOROOT: %w", err)
			return
		}

		dir := cwd
		for {
			parent := filepath.Dir(dir)
			if parent == dir {
				// dir 是 "." 或仅为卷名。
				gorootErr = fmt.Errorf("failed to locate GOROOT/src in any parent directory")
				return
			}

			if base := filepath.Base(dir); base != "src" {
				dir = parent
				continue // dir cannot be GOROOT/src if it doesn't end in "src".
			}

			b, err := os.ReadFile(filepath.Join(dir, "go.mod"))
			if err != nil {
				if os.IsNotExist(err) {
					dir = parent
					continue
				}
				gorootErr = fmt.Errorf("finding GOROOT: %w", err)
				return
			}
			goMod := string(b)

			for goMod != "" {
				var line string
				line, goMod, _ = strings.Cut(goMod, "\n")
				fields := strings.Fields(line)
				if len(fields) >= 2 && fields[0] == "module" && fields[1] == "std" {
					// 发现了"module std"，这是在GOROOT/src中的模块声明！
					gorootPath = parent
					return
				}
			}
		}
	})

	return gorootPath, gorootErr
}

// GOROOT 会报告包含 Go 项目源代码树根目录的路径。这通常等同于 runtime.GOROOT，但在使用 -trimpath 构建测试二进制文件且无法执行 'go env GOROOT' 的情况下仍然有效。
// 
// 如果找不到 GOROOT，如果 t 不为 nil，则跳过 t，否则会 panic。
func GOROOT(t testing.TB) string {
	path, err := findGOROOT()
	if err != nil {
		if t == nil {
			panic(err)
		}
		t.Helper()
		t.Skip(err)
	}
	return path
}

// GoTool报告Go工具的路径。
func GoTool() (string, error) {
	if !HasGoBuild() {
		return "", errors.New("platform cannot run go tool")
	}
	goToolOnce.Do(func() {
		goToolPath, goToolErr = exec.LookPath("go")
	})
	return goToolPath, goToolErr
}

var (
	goToolOnce sync.Once
	goToolPath string
	goToolErr  error
)

// HasSrc 报告整个源代码树是否可在 GOROOT 下获取。
func HasSrc() bool {
	switch runtime.GOOS {
	case "ios":
		return false
	}
	return true
}

// HasExternalNetwork 报告当前系统是否可以使用外部（非本地回环）网络。
func HasExternalNetwork() bool {
	return !testing.Short() && runtime.GOOS != "js" && runtime.GOOS != "wasip1"
}

// MustHaveExternalNetwork 检查当前系统是否可以使用外部（非localhost）网络。
// 如果不能，MustHaveExternalNetwork 会调用 t.Skip 并附带一个解释。
func MustHaveExternalNetwork(t testing.TB) {
	if runtime.GOOS == "js" || runtime.GOOS == "wasip1" {
		t.Helper()
		t.Skipf("skipping test: no external network on %s", runtime.GOOS)
	}
	if testing.Short() {
		t.Helper()
		t.Skipf("skipping test: no external network in -short mode")
	}
}

// HasCGO报告当前系统是否可以使用cgo。
func HasCGO() bool {
	hasCgoOnce.Do(func() {
		goTool, err := GoTool()
		if err != nil {
			return
		}
		cmd := exec.Command(goTool, "env", "CGO_ENABLED")
		cmd.Env = origEnv
		out, err := cmd.Output()
		if err != nil {
			panic(fmt.Sprintf("%v: %v", cmd, out))
		}
		hasCgo, err = strconv.ParseBool(string(bytes.TrimSpace(out)))
		if err != nil {
			panic(fmt.Sprintf("%v: non-boolean output %q", cmd, out))
		}
	})
	return hasCgo
}

var (
	hasCgoOnce sync.Once
	hasCgo     bool
)

// MustHaveCGO 若无 cgo 可用则调用 t.Skip
func MustHaveCGO(t testing.TB) {
	if !HasCGO() {
		t.Skipf("skipping test: no cgo")
	}
}

// CanInternalLink 报告当前系统是否支持使用内部链接来链接程序。
func CanInternalLink(withCgo bool) bool {
	return !platform.MustLinkExternal(runtime.GOOS, runtime.GOARCH, withCgo)
}

// MustInternalLink 检查当前系统是否支持内部链接程序。
// 如果不支持，MustInternalLink 会调用 t.Skip 并附带一个解释。
func MustInternalLink(t testing.TB, withCgo bool) {
	if !CanInternalLink(withCgo) {
		if withCgo && CanInternalLink(false) {
			t.Skipf("skipping test: internal linking on %s/%s is not supported with cgo", runtime.GOOS, runtime.GOARCH)
		}
		t.Skipf("skipping test: internal linking on %s/%s is not supported", runtime.GOOS, runtime.GOARCH)
	}
}

// MustHaveBuildMode 检查当前系统是否能够在给定的构建模式下构建程序。
// 如果不能，MustHaveBuildMode 将调用 t.Skip 并附带一个解释。
func MustHaveBuildMode(t testing.TB, buildmode string) {
	if !platform.BuildModeSupported(runtime.Compiler, buildmode, runtime.GOOS, runtime.GOARCH) {
		t.Skipf("skipping test: build mode %s on %s/%s is not supported by the %s compiler", buildmode, runtime.GOOS, runtime.GOARCH, runtime.Compiler)
	}
}

// HasSymlink 报告当前系统是否可以使用 os.Symlink。
func HasSymlink() bool {
	ok, _ := hasSymlink()
	return ok
}

// MustHaveSymlink 检查当前系统是否支持使用 os.Symlink。
// 如果不支持，MustHaveSymlink 将调用 t.Skip 并附带解释。
func MustHaveSymlink(t testing.TB) {
	ok, reason := hasSymlink()
	if !ok {
		t.Skipf("skipping test: cannot make symlinks on %s/%s: %s", runtime.GOOS, runtime.GOARCH, reason)
	}
}

// HasLink报告当前系统是否可以使用os.Link。
func HasLink() bool {
// 从Android M（棉花糖）版本开始，硬链接文件被阻止，并尝试对文件调用link()会返回EACCES错误。
// - https://code.google.com/p/android-developer-preview/issues/detail?id=3150
	return runtime.GOOS != "plan9" && runtime.GOOS != "android"
}

// MustHaveLink 判断当前系统是否支持使用 os.Link。
// 若不支持，MustHaveLink 会调用 t.Skip 并附带解释原因。
func MustHaveLink(t testing.TB) {
	if !HasLink() {
		t.Skipf("skipping test: hardlinks are not supported on %s/%s", runtime.GOOS, runtime.GOARCH)
	}
}

var flaky = flag.Bool("flaky", false, "run known-flaky tests too")

func SkipFlaky(t testing.TB, issue int) {
	t.Helper()
	if !*flaky {
		t.Skipf("skipping known flaky test without the -flaky flag; see golang.org/issue/%d", issue)
	}
}

func SkipFlakyNet(t testing.TB) {
	t.Helper()
	if v, _ := strconv.ParseBool(os.Getenv("GO_BUILDER_FLAKY_NET")); v {
		t.Skip("skipping test on builder known to have frequent network failures")
	}
}

// CPUIsSlow 报告测试运行的CPU是否疑似性能较慢。
func CPUIsSlow() bool {
	switch runtime.GOARCH {
	case "arm", "mips", "mipsle", "mips64", "mips64le", "wasm":
		return true
	}
	return false
}

// SkipIfShortAndSlow 如果设置了 `-short` 并且怀疑运行测试的 CPU 性能较慢，则跳过 `t`。
//
// （这对于那些通常很快完成但需要大量 CPU 资源的测试很有用。）
func SkipIfShortAndSlow(t testing.TB) {
	if testing.Short() && CPUIsSlow() {
		t.Helper()
		t.Skipf("skipping test in -short mode on %s", runtime.GOARCH)
	}
}

// SkipIfOptimizationOff 如果优化被关闭，则跳过t。
func SkipIfOptimizationOff(t testing.TB) {
	if OptimizationOff() {
		t.Helper()
		t.Skip("skipping test with optimization disabled")
	}
}

// WriteImportcfg编写一个由编译器或链接器使用的importcfg文件，
// 并将其写入dstPath。该文件包含packageFiles中文件映射的条目，
// 同时也包含由pkgs中包（们）递归导入的所有包的信息。
// 
// pkgs可以包含任何可合法传递给'go list'命令的包模式，
// 因此它也可以是一组位于同一目录下的Go源文件列表。
func WriteImportcfg(t testing.TB, dstPath string, packageFiles map[string]string, pkgs ...string) {
	t.Helper()

	icfg := new(bytes.Buffer)
	icfg.WriteString("# import config\n")
	for k, v := range packageFiles {
		fmt.Fprintf(icfg, "packagefile %s=%s\n", k, v)
	}

	if len(pkgs) > 0 {
		// 使用 'go list' 命令来解析任何缺失的包，并重写导入映射。
		cmd := Command(t, GoToolPath(t), "list", "-export", "-deps", "-f", `{{if ne .ImportPath "command-line-arguments"}}{{if .Export}}{{.ImportPath}}={{.Export}}{{end}}{{end}}`)
		cmd.Args = append(cmd.Args, pkgs...)
		cmd.Stderr = new(strings.Builder)
		out, err := cmd.Output()
		if err != nil {
			t.Fatalf("%v: %v\n%s", cmd, err, cmd.Stderr)
		}

		for _, line := range strings.Split(string(out), "\n") {
			if line == "" {
				continue
			}
			importPath, export, ok := strings.Cut(line, "=")
			if !ok {
				t.Fatalf("invalid line in output from %v:\n%s", cmd, line)
			}
			if packageFiles[importPath] == "" {
				fmt.Fprintf(icfg, "packagefile %s=%s\n", importPath, export)
			}
		}
	}

	if err := os.WriteFile(dstPath, icfg.Bytes(), 0666); err != nil {
		t.Fatal(err)
	}
}

// SyscallIsNotSupported 检查错误是否可能表示当前平台或执行环境不支持系统调用。
func SyscallIsNotSupported(err error) bool {
	return syscallIsNotSupported(err)
}
