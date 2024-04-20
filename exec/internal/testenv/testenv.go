// 版权所有 ? 2015 The Go Authors。保留所有权利。
// 本源代码的使用受 BSD 风格许可协议约束，
// 该协议可在 LICENSE 文件中找到。

// Package testenv 提供关于不同测试环境中由 Go 团队运行的功能可用性的信息。
//
// 该包为内部包，因为这些细节特定于 Go 团队的测试设置（在 build.golang.org 上），而非针对一般测试的基础性内容。
package testenv

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"github.com/888go/gosdk/exec/internal/cfg"
	"github.com/888go/gosdk/exec/internal/platform"

	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"testing"
)

// 在初始化时保存原始环境，供检查时使用。测试二进制文件可能在调用HasExec之前修改其环境以改变其行为（例如模拟命令行工具），而该修改后的环境可能导致环境检查行为异常。
var origEnv = os.Environ()

// Builder 返回运行此测试的构建器名称（例如，“linux-amd64”或“windows-386-gce”）。
// 如果该测试未在构建基础设施上运行，则 Builder 返回空字符串。
func Builder() string {
	return os.Getenv("GO_BUILDER_NAME")
}

// HasGoBuild reports whether the current system can build programs with “go build”
// and then run them with os.StartProcess or exec.Command.
func HasGoBuild() bool {
	if os.Getenv("GO_GCFLAGS") != "" {
// 要求 go 命令的每个调用者都传递 "-gcflags="+os.Getenv("GO_GCFLAGS") 给我们带来了太多工作。
// 目前，如果设置了 $GO_GCFLAGS，就报告我们无法运行 go build。
		return false
	}

	goBuildOnce.Do(func() {
// 要运行 'go build'，我们需要能够执行一个 'go' 命令。
// 我们选择执行 'go tool -n compile' 是出于某种考虑，因为这也将确认 cmd/go 能够找到编译器。（在 CL 472096 之前，
// 我们有时会在测试环境中遇到仅安装了 cmd/go，但没有可用的 cmd/compile 来实际构建项目的情况。）
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
// 我们可以假定我们始终拥有一个完整的 Go 工具链可用。
// 然而，该平台要求即使构建纯 Go 程序（包括测试）也需要一个 C 编译器。在测试环境中我们是否拥有一个？
// （例如，在 Android 上，运行测试的设备可能没有安装 C 工具链。）
//
// 如果已明确设置了 CC，则假定我们拥有。否则，使用 'go env CC' 来确定它将默认使用哪个工具链。
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
	// 目前，执行`go run`和`go build`是等效的。
	return HasGoBuild()
}

// MustHaveGoRun checks that the current system can run programs with “go run”.
// If not, MustHaveGoRun calls t.Skip with an explanation.
func MustHaveGoRun(t testing.TB) {
	if !HasGoRun() {
		t.Skipf("skipping test: 'go run' not available on %s/%s", runtime.GOOS, runtime.GOARCH)
	}
}

// HasParallelism 判断当前系统是否支持并行执行多个线程。
// 在 cmd/dist/test.go 中存在此函数的副本。
func HasParallelism() bool {
	switch runtime.GOOS {
	case "js", "wasip1":
		return false
	}
	return true
}

// MustHaveParallelism 检查当前系统是否能够并行执行多个线程。如果不能，MustHaveParallelism 将调用 t.Skip 并附带解释原因。
func MustHaveParallelism(t testing.TB) {
	if !HasParallelism() {
		t.Skipf("skipping test: no parallelism available on %s/%s", runtime.GOOS, runtime.GOARCH)
	}
}

// GoToolPath 返回 Go 工具的路径。
// 它是 GoTool 的便捷包装器。
// 如果工具不可用，GoToolPath 将调用 t.Skip。
// 如果工具本应可用但实际不可用，GoToolPath 将调用 t.Fatal。
func GoToolPath(t testing.TB) string {
	MustHaveGoBuild(t)
	path, err := GoTool()
	if err != nil {
		t.Fatal(err)
	}
// 将影响 Go 命令的所有环境变量添加到测试元数据中。
// 当这些变量发生变化时，缓存的测试结果将被失效。
// 参见 golang.org/issue/32285。
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
// 如果 runtime.GOROOT() 非空，假定其值有效。
//
// （实际情况可能并非如此：例如，用户可能已将 GOROOT 明确设置为错误的目录，或者明确设置了 GOROOT_FINAL 但未设置 GOROOT，并且尚未将树移动到 GOROOT_FINAL。但这些情况较为罕见，若发生此类问题，用户可自行修复。）
			return
		}

// runtime.GOROOT 无法确定 GOROOT 的位置（可能是因为测试二进制文件是使用 -trimpath 参数构建的，或者是因为仅设置了 GOROOT_FINAL 而未设置 GOROOT，且树状结构尚未移动到该位置）。
//
// 由于这里是 internal/testenv，我们可以采取捷径并假设调用者是对 GOROOT/src 某子目录下某个包进行的测试（'go test' 在包含被测包的目录中运行测试）。这意味着如果我们开始沿目录树向上遍历，最终应能找到 GOROOT/src/go.mod，并可以报告该文件的父目录。
//
// 值得注意的是，即使我们不能以子进程形式运行 'go env GOROOT'，这种方法依然有效。

		cwd, err := os.Getwd()
		if err != nil {
			gorootErr = fmt.Errorf("finding GOROOT: %w", err)
			return
		}

		dir := cwd
		for {
			parent := filepath.Dir(dir)
			if parent == dir {
				// dir 是 "." 或仅为一个卷名。
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
					// 发现了“module std”，即GOROOT/src中的模块声明！
					gorootPath = parent
					return
				}
			}
		}
	})

	return gorootPath, gorootErr
}

// GOROOT 返回包含Go项目源代码树根目录的路径。通常情况下，这与runtime.GOROOT等价，但在测试二进制文件使用-trimpath选项构建且无法执行"go env GOROOT"命令时仍可正常工作。
// 
// 若无法找到GOROOT，当t非nil时，GOROOT将跳过对t的操作，否则将引发恐慌。
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

// GoTool 报告 Go 工具的路径。
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

// HasSrc 判断整个源代码树是否均可在 GOROOT 下获取到
func HasSrc() bool {
	switch runtime.GOOS {
	case "ios":
		return false
	}
	return true
}

// HasExternalNetwork 判断当前系统是否可以使用非本地主机（外部）网络
func HasExternalNetwork() bool {
	return !testing.Short() && runtime.GOOS != "js" && runtime.GOOS != "wasip1"
}

// MustHaveExternalNetwork 检查当前系统是否可以使用
// 外部（非本地主机）网络。
// 如果不能，MustHaveExternalNetwork 将调用 t.Skip 并附带解释。
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

// HasCGO 报告当前系统是否可以使用 cgo。
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

// MustHaveCGO 若未启用cgo，则调用t.Skip。
func MustHaveCGO(t testing.TB) {
	if !HasCGO() {
		t.Skipf("skipping test: no cgo")
	}
}

// CanInternalLink 判断当前系统是否支持以内部链接方式链接程序
func CanInternalLink(withCgo bool) bool {
	return !platform.MustLinkExternal(runtime.GOOS, runtime.GOARCH, withCgo)
}

// MustInternalLink 检查当前系统是否支持以内部链接方式链接程序。
// 若不支持，MustInternalLink 将调用 t.Skip 并附带解释原因。
func MustInternalLink(t testing.TB, withCgo bool) {
	if !CanInternalLink(withCgo) {
		if withCgo && CanInternalLink(false) {
			t.Skipf("skipping test: internal linking on %s/%s is not supported with cgo", runtime.GOOS, runtime.GOARCH)
		}
		t.Skipf("skipping test: internal linking on %s/%s is not supported", runtime.GOOS, runtime.GOARCH)
	}
}

// MustHaveBuildMode 检查当前系统是否能够在给定的构建模式下构建程序。
// 如果不能，MustHaveBuildMode 会调用 t.Skip 并附带解释原因。
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

// MustHaveSymlink 检查当前系统是否支持使用 os.Symlink。如果不支持，MustHaveSymlink 会调用 t.Skip 并附带解释原因。
func MustHaveSymlink(t testing.TB) {
	ok, reason := hasSymlink()
	if !ok {
		t.Skipf("skipping test: cannot make symlinks on %s/%s: %s", runtime.GOOS, runtime.GOARCH, reason)
	}
}

// HasLink 报告当前系统是否可以使用 os.Link。
func HasLink() bool {
// 从Android版本M（Marshmallow）开始，系统禁止对文件进行硬链接操作，
// 若尝试调用link()函数对文件进行硬链接，将会返回EACCES错误。
// - 参考：https://code.google.com/p/android-developer-preview/issues/detail?id=3150
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

// CPUIsSlow 报告当前运行测试的 CPU 是否疑似性能较低
func CPUIsSlow() bool {
	switch runtime.GOARCH {
	case "arm", "mips", "mipsle", "mips64", "mips64le", "wasm":
		return true
	}
	return false
}

// SkipIfShortAndSlow 在设置了 -short 参数且运行测试的 CPU 被认为较慢时，跳过 t。
//
// （这对于那些在 CPU 强劲时能快速完成，但实际消耗 CPU 资源较多的测试非常有用。）
func SkipIfShortAndSlow(t testing.TB) {
	if testing.Short() && CPUIsSlow() {
		t.Helper()
		t.Skipf("skipping test in -short mode on %s", runtime.GOARCH)
	}
}

// 如果优化已禁用，则SkipIfOptimizationOff会跳过t。
func SkipIfOptimizationOff(t testing.TB) {
	if OptimizationOff() {
		t.Helper()
		t.Skip("skipping test with optimization disabled")
	}
}

// WriteImportcfg编写一个importcfg文件，该文件由编译器或链接器用于
// dstPath，其中包含packageFiles中文件映射的条目，以及
// 由pkgs中的包（们）递归导入的包。
// 
// pkgs可以包含任何可有效传递给'go list'的包模式，
// 因此它也可以是同一目录下所有Go源文件的列表。
func WriteImportcfg(t testing.TB, dstPath string, packageFiles map[string]string, pkgs ...string) {
	t.Helper()

	icfg := new(bytes.Buffer)
	icfg.WriteString("# import config\n")
	for k, v := range packageFiles {
		fmt.Fprintf(icfg, "packagefile %s=%s\n", k, v)
	}

	if len(pkgs) > 0 {
		// 使用'go list'命令解析所有缺失的包并重写导入映射。
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

// SyscallIsNotSupported 判断 err 是否可能表示当前平台或执行环境中不支持某个系统调用。
func SyscallIsNotSupported(err error) bool {
	return syscallIsNotSupported(err)
}
