
<原文开始>
// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// 版权所有 ? 2021 Go作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。
# <翻译结束>


<原文开始>
// ExperimentFlags represents a set of GOEXPERIMENT flags relative to a baseline
// (platform-default) experiment configuration.
<原文结束>

# <翻译开始>
// ExperimentFlags 代表了一组相对于基准（平台默认）实验配置的 GOEXPERIMENT 标志。
# <翻译结束>


<原文开始>
// Experiment contains the toolchain experiments enabled for the
// current build.
//
// (This is not necessarily the set of experiments the compiler itself
// was built with.)
//
// experimentBaseline specifies the experiment flags that are enabled by
// default in the current toolchain. This is, in effect, the "control"
// configuration and any variation from this is an experiment.
<原文结束>

# <翻译开始>
// Experiment包含当前构建启用的工具链实验。
//
// （这不一定与编译器自身构建时所启用的实验集合相同。）
//
// experimentBaseline指定当前工具链中默认启用的实验标志。实际上，这就是“基准”配置，任何对此配置的变更都将被视为实验。
# <翻译结束>


<原文开始>
// DefaultGOEXPERIMENT is the embedded default GOEXPERIMENT string.
// It is not guaranteed to be canonical.
<原文结束>

# <翻译开始>
// DefaultGOEXPERIMENT 是嵌入的默认 GOEXPERIMENT 字符串。
// 但并不保证其为规范形式。
# <翻译结束>


<原文开始>
// FramePointerEnabled enables the use of platform conventions for
// saving frame pointers.
//
// This used to be an experiment, but now it's always enabled on
// platforms that support it.
//
// Note: must agree with runtime.framepointer_enabled.
<原文结束>

# <翻译开始>
// FramePointerEnabled 用于启用平台约定来保存帧指针。
//
// 这曾经是一项实验功能，但现在在支持该功能的平台上始终启用。
//
// 注意：必须与 runtime.framepointer_enabled 保持一致。
# <翻译结束>


<原文开始>
// ParseGOEXPERIMENT parses a (GOOS, GOARCH, GOEXPERIMENT)
// configuration tuple and returns the enabled and baseline experiment
// flag sets.
//
// TODO(mdempsky): Move to internal/goexperiment.
<原文结束>

# <翻译开始>
// ParseGOEXPERIMENT 解析一个 (GOOS, GOARCH, GOEXPERIMENT) 配置元组，并返回已启用的和基准实验标志集合。
//
// TODO(mdempsky): 移动到 internal/goexperiment。
# <翻译结束>


<原文开始>
	// regabiSupported is set to true on platforms where register ABI is
	// supported and enabled by default.
	// regabiAlwaysOn is set to true on platforms where register ABI is
	// always on.
<原文结束>

# <翻译开始>
// regabiSupported 在平台支持且默认启用寄存器 ABI 的情况下设置为 true。
// regabiAlwaysOn 在始终启用寄存器 ABI 的平台上设置为 true。
# <翻译结束>


<原文开始>
// Start with the statically enabled set of experiments.
<原文结束>

# <翻译开始>
// 从静态启用的一组实验开始。
# <翻译结束>


<原文开始>
	// Pick up any changes to the baseline configuration from the
	// GOEXPERIMENT environment. This can be set at make.bash time
	// and overridden at build time.
<原文结束>

# <翻译开始>
// 从 GOEXPERIMENT 环境变量中获取基准配置的任何更改。此变量可以在运行 make.bash 时设置，并在构建时覆盖。
# <翻译结束>


<原文开始>
// Create a map of known experiment names.
<原文结束>

# <翻译开始>
// 创建一个已知实验名称的映射（map）。
# <翻译结束>


<原文开始>
		// "regabi" is an alias for all working regabi
		// subexperiments, and not an experiment itself. Doing
		// this as an alias make both "regabi" and "noregabi"
		// do the right thing.
<原文结束>

# <翻译开始>
// "regabi" 是所有有效 regabi 子实验的别名，本身并不是一个实验。通过设置别名的方式，可以使 "regabi" 和 "noregabi" 都能正确执行相应操作。
# <翻译结束>


<原文开始>
				// GOEXPERIMENT=none disables all experiment flags.
				// This is used by cmd/dist, which doesn't know how
				// to build with any experiment flags.
<原文结束>

# <翻译开始>
// GOEXPERIMENT=none 禁用所有实验性标志。
// 这被 cmd/dist 使用，它不知道如何在任何实验性标志下构建。
# <翻译结束>


<原文开始>
// regabi is only supported on amd64, arm64, riscv64, ppc64 and ppc64le.
<原文结束>

# <翻译开始>
// regabi 仅支持 amd64、arm64、riscv64、ppc64 和 ppc64le 架构。
# <翻译结束>


<原文开始>
// Check regabi dependencies.
<原文结束>

# <翻译开始>
// 检查 regabi 依赖项。
# <翻译结束>


<原文开始>
// String returns the canonical GOEXPERIMENT string to enable this experiment
// configuration. (Experiments in the same state as in the baseline are elided.)
<原文结束>

# <翻译开始>
// String 方法返回用于启用该实验配置的规范 GOEXPERIMENT 字符串。
// （与基线相同状态的实验将被省略。）
# <翻译结束>


<原文开始>
// expList returns the list of lower-cased experiment names for
// experiments that differ from base. base may be nil to indicate no
// experiments. If all is true, then include all experiment flags,
// regardless of base.
<原文结束>

# <翻译开始>
// expList 返回与 base 不同的实验名称列表，所有名称均转换为小写。base 可以为 nil，表示没有实验。如果 all 为真，则无论 base 如何，都包含所有实验标志。
# <翻译结束>


<原文开始>
// Enabled returns a list of enabled experiments, as
// lower-cased experiment names.
<原文结束>

# <翻译开始>
// Enabled 返回一个已启用实验的列表，其格式为小写实验名称。
# <翻译结束>


<原文开始>
// All returns a list of all experiment settings.
// Disabled experiments appear in the list prefixed by "no".
<原文结束>

# <翻译开始>
// All 返回所有实验设置的列表。
// 已禁用的实验在列表中会以 "no" 前缀显示。
# <翻译结束>

