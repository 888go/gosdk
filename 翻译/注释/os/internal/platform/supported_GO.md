
<原文开始>
// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// 版权所有 2018 The Go 作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。
# <翻译结束>


<原文开始>
// An OSArch is a pair of GOOS and GOARCH values indicating a platform.
<原文结束>

# <翻译开始>
// OSArch 是一个 GOOS 和 GOARCH 值的组合，表示一个操作系统平台。
# <翻译结束>


<原文开始>
// RaceDetectorSupported reports whether goos/goarch supports the race
// detector. There is a copy of this function in cmd/dist/test.go.
// Race detector only supports 48-bit VMA on arm64. But it will always
// return true for arm64, because we don't have VMA size information during
// the compile time.
<原文结束>

# <翻译开始>
// RaceDetectorSupported报告goos/goarch是否支持竞态条件检测器。在cmd/dist/test.go中也有此函数的副本。
// 竞态检测器仅在arm64架构上支持48位虚拟内存地址（VMA）。但在编译时我们没有VMA大小信息，因此对于arm64，它始终返回true。
# <翻译结束>


<原文开始>
// MSanSupported reports whether goos/goarch supports the memory
// sanitizer option.
<原文结束>

# <翻译开始>
// MSanSupported报告goos/goarch是否支持内存sanitizer选项。
# <翻译结束>


<原文开始>
// ASanSupported reports whether goos/goarch supports the address
// sanitizer option.
<原文结束>

# <翻译开始>
// ASanSupported 判断当前 goos/goarch 是否支持使用地址 sanitizer 选项
# <翻译结束>


<原文开始>
// FuzzSupported reports whether goos/goarch supports fuzzing
// ('go test -fuzz=.').
<原文结束>

# <翻译开始>
// FuzzSupported 报告 goos/goarch 是否支持模糊测试（'go test -fuzz=.'）。
# <翻译结束>


<原文开始>
// FuzzInstrumented reports whether fuzzing on goos/goarch uses coverage
// instrumentation. (FuzzInstrumented implies FuzzSupported.)
<原文结束>

# <翻译开始>
// FuzzInstrumented 判断在 goos/goarch 环境中，是否使用了覆盖率（instrumentation）进行模糊测试。（FuzzInstrumented 意味着 FuzzSupported。）
# <翻译结束>


<原文开始>
// TODO(#14565): support more architectures.
<原文结束>

# <翻译开始>
// TODO(#14565)：支持更多架构。
# <翻译结束>


<原文开始>
// MustLinkExternal reports whether goos/goarch requires external linking
// with or without cgo dependencies.
<原文结束>

# <翻译开始>
// MustLinkExternal 报告 goos/goarch 是否要求使用外部链接，无论是否包含 cgo 依赖。
# <翻译结束>


<原文开始>
			// Internally linking cgo is incomplete on some architectures.
			// https://go.dev/issue/14449
<原文结束>

# <翻译开始>
// 在某些架构上，内部链接cgo是不完整的。
// 参考：https://go.dev/issue/14449
# <翻译结束>


<原文开始>
// windows/arm64 internal linking is not implemented.
<原文结束>

# <翻译开始>
// Windows/arm64内部链接功能未实现。
# <翻译结束>


<原文开始>
			// Big Endian PPC64 cgo internal linking is not implemented for aix or linux.
			// https://go.dev/issue/8912
<原文结束>

# <翻译开始>
// 对于AIX或Linux，BigEndian PPC64的cgo内部链接功能未实现。
// 参见：https://go.dev/issue/8912
# <翻译结束>


<原文开始>
			// It seems that on Dragonfly thread local storage is
			// set up by the dynamic linker, so internal cgo linking
			// doesn't work. Test case is "go test runtime/cgo".
<原文结束>

# <翻译开始>
// 在 Dragonfly 系统上，线程局部存储似乎是通过动态链接器设置的，因此内部 cgo 链接无法正常工作。测试用例为 "go test runtime/cgo"。
# <翻译结束>


<原文开始>
// BuildModeSupported reports whether goos/goarch supports the given build mode
// using the given compiler.
// There is a copy of this function in cmd/dist/test.go.
<原文结束>

# <翻译开始>
// BuildModeSupported 报告给定的编译器是否支持 goos/goarch 构建模式。
// 此函数的副本在 cmd/dist/test.go 中存在。
# <翻译结束>


<原文开始>
				// linux/ppc64 not supported because it does
				// not support external linking mode yet.
<原文结束>

# <翻译开始>
// 因为linux/ppc64还不支持外部链接模式，所以不被支持。
# <翻译结束>


<原文开始>
				// Other targets do not support -shared,
				// per ParseFlags in
				// cmd/compile/internal/base/flag.go.
				// For c-archive the Go tool passes -shared,
				// so that the result is suitable for inclusion
				// in a PIE or shared library.
<原文结束>

# <翻译开始>
// 其他目标不支持 `-shared`，这是根据 `cmd/compile/internal/base/flag.go` 中的 ParseFlags 执行的。
// 对于 c-archives，Go 工具会传递 `-shared`，以便生成的结果适合包含在 PIE（可执行格式）或共享库中。
# <翻译结束>


<原文开始>
// DefaultPIE reports whether goos/goarch produces a PIE binary when using the
// "default" buildmode. On Windows this is affected by -race,
// so force the caller to pass that in to centralize that choice.
<原文结束>

# <翻译开始>
// DefaultPIE 报告在使用“默认”构建模式时，goos/goarch 是否会产生一个 PIE（位置无关可执行文件）二进制文件。对于 Windows 平台，这会受到 -race 标志的影响，因此要求调用者传入该标志以集中处理该选项。
# <翻译结束>


<原文开始>
			// PIE is not supported with -race on windows;
			// see https://go.dev/cl/416174.
<原文结束>

# <翻译开始>
// PIE在Windows上不支持-race选项；详情见https://go.dev/cl/416174。
# <翻译结束>


<原文开始>
// ExecutableHasDWARF reports whether the linked executable includes DWARF
// symbols on goos/goarch.
<原文结束>

# <翻译开始>
// ExecutableHasDWARF 判断在 goos/goarch 平台上，链接后的可执行文件是否包含 DWARF 符号。
# <翻译结束>


<原文开始>
// osArchInfo describes information about an OSArch extracted from cmd/dist and
// stored in the generated distInfo map.
<原文结束>

# <翻译开始>
// osArchInfo描述了从cmd/dist中提取的关于OSArch的信息，并存储在生成的distInfo映射中。
# <翻译结束>


<原文开始>
// CgoSupported reports whether goos/goarch supports cgo.
<原文结束>

# <翻译开始>
// CgoSupported 报告 goos/goarch 是否支持 cgo。
# <翻译结束>


<原文开始>
// Broken reportsr whether goos/goarch is considered a broken port.
// (See https://go.dev/wiki/PortingPolicy#broken-ports.)
<原文结束>

# <翻译开始>
// Broken reports 是否认为goos/goarch是一个被认为存在问题的端口。
// （参考 https://go.dev/wiki/PortingPolicy#broken-ports.）
# <翻译结束>

