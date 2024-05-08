// 版权所有 2018 The Go 作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

//go:generate go test . -run=^TestGenerated$ -fix

package platform

// OSArch 是一个 GOOS 和 GOARCH 值的组合，表示一个操作系统平台。
type OSArch struct {
	GOOS, GOARCH string
}

func (p OSArch) String() string {
	return p.GOOS + "/" + p.GOARCH
}

// RaceDetectorSupported报告goos/goarch是否支持竞态条件检测器。在cmd/dist/test.go中也有此函数的副本。
// 竞态检测器仅在arm64架构上支持48位虚拟内存地址（VMA）。但在编译时我们没有VMA大小信息，因此对于arm64，它始终返回true。
func RaceDetectorSupported(goos, goarch string) bool {
	switch goos {
	case "linux":
		return goarch == "amd64" || goarch == "ppc64le" || goarch == "arm64" || goarch == "s390x"
	case "darwin":
		return goarch == "amd64" || goarch == "arm64"
	case "freebsd", "netbsd", "openbsd", "windows":
		return goarch == "amd64"
	default:
		return false
	}
}

// MSanSupported报告goos/goarch是否支持内存sanitizer选项。
func MSanSupported(goos, goarch string) bool {
	switch goos {
	case "linux":
		return goarch == "amd64" || goarch == "arm64" || goarch == "loong64"
	case "freebsd":
		return goarch == "amd64"
	default:
		return false
	}
}

// ASanSupported 判断当前 goos/goarch 是否支持使用地址 sanitizer 选项
func ASanSupported(goos, goarch string) bool {
	switch goos {
	case "linux":
		return goarch == "arm64" || goarch == "amd64" || goarch == "loong64" || goarch == "riscv64" || goarch == "ppc64le"
	default:
		return false
	}
}

// FuzzSupported 报告 goos/goarch 是否支持模糊测试（'go test -fuzz=.'）。
func FuzzSupported(goos, goarch string) bool {
	switch goos {
	case "darwin", "freebsd", "linux", "windows":
		return true
	default:
		return false
	}
}

// FuzzInstrumented 判断在 goos/goarch 环境中，是否使用了覆盖率（instrumentation）进行模糊测试。（FuzzInstrumented 意味着 FuzzSupported。）
func FuzzInstrumented(goos, goarch string) bool {
	switch goarch {
	case "amd64", "arm64":
		// TODO(#14565)：支持更多架构。
		return FuzzSupported(goos, goarch)
	default:
		return false
	}
}

// MustLinkExternal 报告 goos/goarch 是否要求使用外部链接，无论是否包含 cgo 依赖。
func MustLinkExternal(goos, goarch string, withCgo bool) bool {
	if withCgo {
		switch goarch {
		case "loong64", "mips", "mipsle", "mips64", "mips64le":
// 在某些架构上，内部链接cgo是不完整的。
// 参考：https://go.dev/issue/14449
			return true
		case "arm64":
			if goos == "windows" {
				// Windows/arm64内部链接功能未实现。
				return true
			}
		case "ppc64":
// 对于AIX或Linux，BigEndian PPC64的cgo内部链接功能未实现。
// 参见：https://go.dev/issue/8912
			if goos == "aix" || goos == "linux" {
				return true
			}
		}

		switch goos {
		case "android":
			return true
		case "dragonfly":
// 在 Dragonfly 系统上，线程局部存储似乎是通过动态链接器设置的，因此内部 cgo 链接无法正常工作。测试用例为 "go test runtime/cgo"。
			return true
		}
	}

	switch goos {
	case "android":
		if goarch != "arm64" {
			return true
		}
	case "ios":
		if goarch == "arm64" {
			return true
		}
	}
	return false
}

// BuildModeSupported 报告给定的编译器是否支持 goos/goarch 构建模式。
// 此函数的副本在 cmd/dist/test.go 中存在。
func BuildModeSupported(compiler, buildmode, goos, goarch string) bool {
	if compiler == "gccgo" {
		return true
	}

	if _, ok := distInfo[OSArch{goos, goarch}]; !ok {
		return false // platform unrecognized
	}

	platform := goos + "/" + goarch
	switch buildmode {
	case "archive":
		return true

	case "c-archive":
		switch goos {
		case "aix", "darwin", "ios", "windows":
			return true
		case "linux":
			switch goarch {
			case "386", "amd64", "arm", "armbe", "arm64", "arm64be", "loong64", "ppc64le", "riscv64", "s390x":
// 因为linux/ppc64还不支持外部链接模式，所以不被支持。
				return true
			default:
// 其他目标不支持 `-shared`，这是根据 `cmd/compile/internal/base/flag.go` 中的 ParseFlags 执行的。
// 对于 c-archives，Go 工具会传递 `-shared`，以便生成的结果适合包含在 PIE（可执行格式）或共享库中。
				return false
			}
		case "freebsd":
			return goarch == "amd64"
		}
		return false

	case "c-shared":
		switch platform {
		case "linux/amd64", "linux/arm", "linux/arm64", "linux/loong64", "linux/386", "linux/ppc64le", "linux/riscv64", "linux/s390x",
			"android/amd64", "android/arm", "android/arm64", "android/386",
			"freebsd/amd64",
			"darwin/amd64", "darwin/arm64",
			"windows/amd64", "windows/386", "windows/arm64":
			return true
		}
		return false

	case "default":
		return true

	case "exe":
		return true

	case "pie":
		switch platform {
		case "linux/386", "linux/amd64", "linux/arm", "linux/arm64", "linux/loong64", "linux/ppc64le", "linux/riscv64", "linux/s390x",
			"android/amd64", "android/arm", "android/arm64", "android/386",
			"freebsd/amd64",
			"darwin/amd64", "darwin/arm64",
			"ios/amd64", "ios/arm64",
			"aix/ppc64",
			"windows/386", "windows/amd64", "windows/arm", "windows/arm64":
			return true
		}
		return false

	case "shared":
		switch platform {
		case "linux/386", "linux/amd64", "linux/arm", "linux/arm64", "linux/ppc64le", "linux/s390x":
			return true
		}
		return false

	case "plugin":
		switch platform {
		case "linux/amd64", "linux/arm", "linux/arm64", "linux/386", "linux/loong64", "linux/s390x", "linux/ppc64le",
			"android/amd64", "android/386",
			"darwin/amd64", "darwin/arm64",
			"freebsd/amd64":
			return true
		}
		return false

	default:
		return false
	}
}

func InternalLinkPIESupported(goos, goarch string) bool {
	switch goos + "/" + goarch {
	case "android/arm64",
		"darwin/amd64", "darwin/arm64",
		"linux/amd64", "linux/arm64", "linux/ppc64le",
		"windows/386", "windows/amd64", "windows/arm", "windows/arm64":
		return true
	}
	return false
}

// DefaultPIE 报告在使用“默认”构建模式时，goos/goarch 是否会产生一个 PIE（位置无关可执行文件）二进制文件。对于 Windows 平台，这会受到 -race 标志的影响，因此要求调用者传入该标志以集中处理该选项。
func DefaultPIE(goos, goarch string, isRace bool) bool {
	switch goos {
	case "android", "ios":
		return true
	case "windows":
		if isRace {
// PIE在Windows上不支持-race选项；详情见https://go.dev/cl/416174。
			return false
		}
		return true
	case "darwin":
		return true
	}
	return false
}

// ExecutableHasDWARF 判断在 goos/goarch 平台上，链接后的可执行文件是否包含 DWARF 符号。
func ExecutableHasDWARF(goos, goarch string) bool {
	switch goos {
	case "plan9", "ios":
		return false
	}
	return true
}

// osArchInfo描述了从cmd/dist中提取的关于OSArch的信息，并存储在生成的distInfo映射中。
type osArchInfo struct {
	CgoSupported bool
	FirstClass   bool
	Broken       bool
}

// CgoSupported 报告 goos/goarch 是否支持 cgo。
func CgoSupported(goos, goarch string) bool {
	return distInfo[OSArch{goos, goarch}].CgoSupported
}

// FirstClass reports whether goos/goarch is considered a “first class” port.
// (See https://go.dev/wiki/PortingPolicy#first-class-ports.)
func FirstClass(goos, goarch string) bool {
	return distInfo[OSArch{goos, goarch}].FirstClass
}

// Broken reports 是否认为goos/goarch是一个被认为存在问题的端口。
// （参考 https://go.dev/wiki/PortingPolicy#broken-ports.）
func Broken(goos, goarch string) bool {
	return distInfo[OSArch{goos, goarch}].Broken
}
