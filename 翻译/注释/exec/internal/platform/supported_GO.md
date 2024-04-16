
<原文开始>
// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// 版权所有 ? 2018 The Go 作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可协议约束，
// 该协议可在 LICENSE 文件中找到。
# <翻译结束>


<原文开始>
// An OSArch is a pair of GOOS and GOARCH values indicating a platform.
<原文结束>

# <翻译开始>
// OSArch 是一个包含 GOOS 和 GOARCH 值的组合，用于表示一个平台。
# <翻译结束>


<原文开始>
// RaceDetectorSupported reports whether goos/goarch supports the race
// detector. There is a copy of this function in cmd/dist/test.go.
// Race detector only supports 48-bit VMA on arm64. But it will always
// return true for arm64, because we don't have VMA size information during
// the compile time.
<原文结束>

# <翻译开始>
// RaceDetectorSupported 报告 goos/goarch 是否支持竞态检测器。cmd/dist/test.go 中存在此函数的一个副本。
// 竞态检测器仅支持 arm64 上的 48 位虚拟内存地址（VMA）。但对于 arm64，它始终会返回 true，因为我们编译时没有 VMA 大小的信息。
# <翻译结束>


<原文开始>
// MSanSupported reports whether goos/goarch supports the memory
// sanitizer option.
<原文结束>

# <翻译开始>
// MSanSupported 报告 goos/goarch 是否支持内存 sanitizer 选项。
# <翻译结束>


<原文开始>
// ASanSupported reports whether goos/goarch supports the address
// sanitizer option.
<原文结束>

# <翻译开始>
// ASanSupported 报告 goos/goarch 是否支持地址 sanitizer 选项。
# <翻译结束>


<原文开始>
// FuzzSupported reports whether goos/goarch supports fuzzing
// ('go test -fuzz=.').
<原文结束>

# <翻译开始>
// FuzzSupported 报告 goos/goarch 是否支持模糊测试（'go test -fuzz='）。
# <翻译结束>


<原文开始>
// FuzzInstrumented reports whether fuzzing on goos/goarch uses coverage
// instrumentation. (FuzzInstrumented implies FuzzSupported.)
<原文结束>

# <翻译开始>
// FuzzInstrumented 判断在 goos/goarch 平台上进行 fuzz 测试时是否使用了覆盖率检测工具。
// （若 FuzzInstrumented 为真，则意味着 FuzzSupported 亦为真。）
# <翻译结束>


<原文开始>
// TODO(#14565): support more architectures.
<原文结束>

# <翻译开始>
// 待办事项(#14565)：支持更多架构。
# <翻译结束>


<原文开始>
// MustLinkExternal reports whether goos/goarch requires external linking
// with or without cgo dependencies.
<原文结束>

# <翻译开始>
// MustLinkExternal 报告 goos/goarch 是否要求在有无 cgo 依赖的情况下进行外部链接
# <翻译结束>


<原文开始>
			// Internally linking cgo is incomplete on some architectures.
			// https://go.dev/issue/14449
<原文结束>

# <翻译开始>
// 在某些架构上，内部链接 cgo 仍不完整。
// 参见：https://go.dev/issue/14449
# <翻译结束>


<原文开始>
// windows/arm64 internal linking is not implemented.
<原文结束>

# <翻译开始>
// windows/arm64 不支持内部链接
# <翻译结束>


<原文开始>
			// Big Endian PPC64 cgo internal linking is not implemented for aix or linux.
			// https://go.dev/issue/8912
<原文结束>

# <翻译开始>
// PPC64大端模式下的cgo内部链接在aix或linux系统中尚未实现。
// 参见：https://go.dev/issue/8912
# <翻译结束>


<原文开始>
			// It seems that on Dragonfly thread local storage is
			// set up by the dynamic linker, so internal cgo linking
			// doesn't work. Test case is "go test runtime/cgo".
<原文结束>

# <翻译开始>
// 在 Dragonfly 系统上，线程局部存储似乎是由动态链接器设置的，因此内部 cgo 链接无法正常工作。测试用例为 "go test runtime/cgo"。
# <翻译结束>


<原文开始>
// BuildModeSupported reports whether goos/goarch supports the given build mode
// using the given compiler.
// There is a copy of this function in cmd/dist/test.go.
<原文结束>

# <翻译开始>
// BuildModeSupported 报告在使用给定编译器的情况下，goos/goarch 是否支持给定的构建模式。
// 在 cmd/dist/test.go 中存在此函数的一个副本。
# <翻译结束>


<原文开始>
				// linux/ppc64 not supported because it does
				// not support external linking mode yet.
<原文结束>

# <翻译开始>
// linux/ppc64 未予支持，因为它尚不支持外部链接模式。
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
// 其他目标不支持 `-shared` 选项，
// 这一点在 `cmd/compile/internal/base/flag.go` 的 `ParseFlags` 函数中有说明。
// 对于 `c-archive` 目标，Go 工具会传递 `-shared` 参数，
// 以确保生成结果适用于作为 PIE（位置无关可执行文件）或共享库的一部分。
# <翻译结束>


<原文开始>
// DefaultPIE reports whether goos/goarch produces a PIE binary when using the
// "default" buildmode. On Windows this is affected by -race,
// so force the caller to pass that in to centralize that choice.
<原文结束>

# <翻译开始>
// DefaultPIE 报告在使用“默认”构建模式时，goos/goarch 是否会产生一个 PIE（位置无关可执行文件）。在 Windows 上，这会受到 -race 参数的影响，因此要求调用者传入该参数以集中处理该选项。
# <翻译结束>


<原文开始>
			// PIE is not supported with -race on windows;
			// see https://go.dev/cl/416174.
<原文结束>

# <翻译开始>
// 在 Windows 系统上，PIE（位置无关可执行文件）与 `-race` 标志不兼容；
// 详情请参见 https://go.dev/cl/416174。
# <翻译结束>


<原文开始>
// ExecutableHasDWARF reports whether the linked executable includes DWARF
// symbols on goos/goarch.
<原文结束>

# <翻译开始>
// ExecutableHasDWARF 报告在 goos/goarch 平台上链接的可执行文件是否包含 DWARF 符号
# <翻译结束>


<原文开始>
// osArchInfo describes information about an OSArch extracted from cmd/dist and
// stored in the generated distInfo map.
<原文结束>

# <翻译开始>
// osArchInfo 描述了从 cmd/dist 中提取并存储在生成的 distInfo 映射中的 OSArch 信息。
# <翻译结束>


<原文开始>
// CgoSupported reports whether goos/goarch supports cgo.
<原文结束>

# <翻译开始>
// CgoSupported 判断 goos/goarch 是否支持 cgo
# <翻译结束>


<原文开始>
// Broken reportsr whether goos/goarch is considered a broken port.
// (See https://go.dev/wiki/PortingPolicy#broken-ports.)
<原文结束>

# <翻译开始>
// Broken 报告 goos/goarch 是否被视为已损坏的端口。
// （参见 https://go.dev/wiki/PortingPolicy#broken-ports。）
# <翻译结束>

