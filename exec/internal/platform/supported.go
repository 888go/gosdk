// 版权所有 ? 2018 The Go 作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可协议约束，
// 该协议可在 LICENSE 文件中找到。

//go:generate go test . -run=^TestGenerated$ -fix

package platform

// OSArch 是一个包含 GOOS 和 GOARCH 值的组合，用于表示一个平台。
type OSArch struct {
	GOOS, GOARCH string
}


// ff:
func (p OSArch) String() string {
	return p.GOOS + "/" + p.GOARCH
}

// RaceDetectorSupported 报告 goos/goarch 是否支持竞态检测器。cmd/dist/test.go 中存在此函数的一个副本。
// 竞态检测器仅支持 arm64 上的 48 位虚拟内存地址（VMA）。但对于 arm64，它始终会返回 true，因为我们编译时没有 VMA 大小的信息。

// ff:
// goarch:
// goos:
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

// MSanSupported 报告 goos/goarch 是否支持内存 sanitizer 选项。

// ff:
// goarch:
// goos:
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

// ASanSupported 报告 goos/goarch 是否支持地址 sanitizer 选项。

// ff:
// goarch:
// goos:
func ASanSupported(goos, goarch string) bool {
	switch goos {
	case "linux":
		return goarch == "arm64" || goarch == "amd64" || goarch == "loong64" || goarch == "riscv64" || goarch == "ppc64le"
	default:
		return false
	}
}

// FuzzSupported 报告 goos/goarch 是否支持模糊测试（'go test -fuzz='）。

// ff:
// goarch:
// goos:
func FuzzSupported(goos, goarch string) bool {
	switch goos {
	case "darwin", "freebsd", "linux", "windows":
		return true
	default:
		return false
	}
}

// FuzzInstrumented 判断在 goos/goarch 平台上进行 fuzz 测试时是否使用了覆盖率检测工具。
// （若 FuzzInstrumented 为真，则意味着 FuzzSupported 亦为真。）

// ff:
// goarch:
// goos:
func FuzzInstrumented(goos, goarch string) bool {
	switch goarch {
	case "amd64", "arm64":
		// 待办事项(#14565)：支持更多架构。
		return FuzzSupported(goos, goarch)
	default:
		return false
	}
}

// MustLinkExternal 报告 goos/goarch 是否要求在有无 cgo 依赖的情况下进行外部链接

// ff:
// withCgo:
// goarch:
// goos:
func MustLinkExternal(goos, goarch string, withCgo bool) bool {
	if withCgo {
		switch goarch {
		case "loong64", "mips", "mipsle", "mips64", "mips64le":
// 在某些架构上，内部链接 cgo 仍不完整。
// 参见：https://go.dev/issue/14449
			return true
		case "arm64":
			if goos == "windows" {
				// windows/arm64 不支持内部链接
				return true
			}
		case "ppc64":
// PPC64大端模式下的cgo内部链接在aix或linux系统中尚未实现。
// 参见：https://go.dev/issue/8912
			if goos == "aix" || goos == "linux" {
				return true
			}
		}

		switch goos {
		case "android":
			return true
		case "dragonfly":
// 在 Dragonfly 系统上，线程局部存储似乎是由动态链接器设置的，因此内部 cgo 链接无法正常工作。测试用例为 "go test runtime/cgo"。
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

// BuildModeSupported 报告在使用给定编译器的情况下，goos/goarch 是否支持给定的构建模式。
// 在 cmd/dist/test.go 中存在此函数的一个副本。

// ff:
// goarch:
// goos:
// buildmode:
// compiler:
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
// linux/ppc64 未予支持，因为它尚不支持外部链接模式。
				return true
			default:
// 其他目标不支持 `-shared` 选项，
// 这一点在 `cmd/compile/internal/base/flag.go` 的 `ParseFlags` 函数中有说明。
// 对于 `c-archive` 目标，Go 工具会传递 `-shared` 参数，
// 以确保生成结果适用于作为 PIE（位置无关可执行文件）或共享库的一部分。
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


// ff:
// goarch:
// goos:
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

// DefaultPIE 报告在使用“默认”构建模式时，goos/goarch 是否会产生一个 PIE（位置无关可执行文件）。在 Windows 上，这会受到 -race 参数的影响，因此要求调用者传入该参数以集中处理该选项。

// ff:
// isRace:
// goarch:
// goos:
func DefaultPIE(goos, goarch string, isRace bool) bool {
	switch goos {
	case "android", "ios":
		return true
	case "windows":
		if isRace {
// 在 Windows 系统上，PIE（位置无关可执行文件）与 `-race` 标志不兼容；
// 详情请参见 https://go.dev/cl/416174。
			return false
		}
		return true
	case "darwin":
		return true
	}
	return false
}

// ExecutableHasDWARF 报告在 goos/goarch 平台上链接的可执行文件是否包含 DWARF 符号

// ff:
// goarch:
// goos:
func ExecutableHasDWARF(goos, goarch string) bool {
	switch goos {
	case "plan9", "ios":
		return false
	}
	return true
}

// osArchInfo 描述了从 cmd/dist 中提取并存储在生成的 distInfo 映射中的 OSArch 信息。
type osArchInfo struct {
	CgoSupported bool
	FirstClass   bool
	Broken       bool
}

// CgoSupported 判断 goos/goarch 是否支持 cgo

// ff:
// goarch:
// goos:
func CgoSupported(goos, goarch string) bool {
	return distInfo[OSArch{goos, goarch}].CgoSupported
}

// FirstClass reports whether goos/goarch is considered a “first class” port.
// (See https://go.dev/wiki/PortingPolicy#first-class-ports.)

// ff:
// goarch:
// goos:
func FirstClass(goos, goarch string) bool {
	return distInfo[OSArch{goos, goarch}].FirstClass
}

// Broken 报告 goos/goarch 是否被视为已损坏的端口。
// （参见 https://go.dev/wiki/PortingPolicy#broken-ports。）

// ff:
// goarch:
// goos:
func Broken(goos, goarch string) bool {
	return distInfo[OSArch{goos, goarch}].Broken
}
