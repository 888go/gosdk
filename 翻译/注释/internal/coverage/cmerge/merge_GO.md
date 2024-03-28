
<原文开始>
// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// 版权所有 ? 2022 Go作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。
# <翻译结束>


<原文开始>
// package cmerge provides a few small utility APIs for helping
// with merging of counter data for a given function.
<原文结束>

# <翻译开始>
// cmerge 包提供了一些小型辅助API，用于针对给定函数的计数器数据的合并操作。
# <翻译结束>


<原文开始>
// Merger provides state and methods to help manage the process of
// merging together coverage counter data for a given function, for
// tools that need to implicitly merge counter as they read multiple
// coverage counter data files.
<原文结束>

# <翻译开始>
// Merger 提供状态和方法，用于协助管理在特定工具读取多个覆盖率计数数据文件时，针对给定函数合并覆盖率计数数据的过程。
# <翻译结束>


<原文开始>
// MergeCounters takes the counter values in 'src' and merges them
// into 'dst' according to the correct counter mode.
<原文结束>

# <翻译开始>
// MergeCounters 将“src”中的计数器值根据正确的计数器模式合并到“dst”中。
# <翻译结束>


<原文开始>
// Saturating add does a saturating addition of 'dst' and 'src',
// returning added value or math.MaxUint32 if there is an overflow.
// Overflows are recorded in case the client needs to track them.
<原文结束>

# <翻译开始>
// 饱和加法对“dst”和“src”执行饱和加法，
// 如果发生溢出，则返回添加后的值或math.MaxUint32。
// 记录溢出情况，以便客户端在需要时跟踪它们。
# <翻译结束>


<原文开始>
// Saturating add does a saturing addition of 'dst' and 'src',
// returning added value or math.MaxUint32 plus an overflow flag.
<原文结束>

# <翻译开始>
// 饱和加法对'dst'和'src'执行饱和加法，
// 返回添加后的值或在溢出时返回math.MaxUint32及一个溢出标志。
# <翻译结束>


<原文开始>
// SetModeAndGranularity records the counter mode and granularity for
// the current merge. In the specific case of merging across coverage
// data files from different binaries, where we're combining data from
// more than one meta-data file, we need to check for mode/granularity
// clashes.
<原文结束>

# <翻译开始>
// SetModeAndGranularity 用于记录当前合并操作的计数器模式和粒度。在特定情况下，当我们在不同二进制文件的覆盖率数据文件之间进行合并，并且需要组合来自多个元数据文件的数据时，我们需要检查模式/粒度冲突。
# <翻译结束>


<原文开始>
// Collect counter mode and granularity so as to detect clashes.
<原文结束>

# <翻译开始>
// 收集计数器模式和粒度，以便检测冲突。
# <翻译结束>

