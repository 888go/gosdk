
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
// Building the runtime package with coverage instrumentation enabled
// is tricky.  For all other packages, you can be guaranteed that
// the package init function is run before any functions are executed,
// but this invariant is not maintained for packages such as "runtime",
// "internal/cpu", etc. To handle this, hard-code the package ID for
// the set of packages whose functions may be running before the
// init function of the package is complete.
//
// Hardcoding is unfortunate because it means that the tool that does
// coverage instrumentation has to keep a list of runtime packages,
// meaning that if someone makes changes to the pkg "runtime"
// dependencies, unexpected behavior will result for coverage builds.
// The coverage runtime will detect and report the unexpected
// behavior; look for an error of this form:
//
//    internal error in coverage meta-data tracking:
//    list of hard-coded runtime package IDs needs revising.
//    registered list:
//    slot: 0 path='internal/cpu'  hard-coded id: 1
//    slot: 1 path='internal/goarch'  hard-coded id: 2
//    slot: 2 path='runtime/internal/atomic'  hard-coded id: 3
//    slot: 3 path='internal/goos'
//    slot: 4 path='runtime/internal/sys'  hard-coded id: 5
//    slot: 5 path='internal/abi'  hard-coded id: 4
//    slot: 6 path='runtime/internal/math'  hard-coded id: 6
//    slot: 7 path='internal/bytealg'  hard-coded id: 7
//    slot: 8 path='internal/goexperiment'
//    slot: 9 path='runtime/internal/syscall'  hard-coded id: 8
//    slot: 10 path='runtime'  hard-coded id: 9
//    fatal error: runtime.addCovMeta
//
// For the error above, the hard-coded list is missing "internal/goos"
// and "internal/goexperiment" ; the developer in question will need
// to copy the list above into "rtPkgs" below.
//
// Note: this strategy assumes that the list of dependencies of
// package runtime is fixed, and doesn't vary depending on OS/arch. If
// this were to be the case, we would need a table of some sort below
// as opposed to a fixed list.
<原文结束>

# <翻译开始>
// 在启用覆盖率检测的情况下构建运行时包（runtime package）具有一定复杂性。对于所有其他包，可以确保在执行任何函数之前先运行包初始化函数，但对于诸如“runtime”、“internal/cpu”等包，这一不变式并不成立。为了解决这个问题，对于那些其函数可能在包初始化函数完成前就开始运行的包，我们硬编码了包ID。
// 
// 硬编码这种方法不幸的是，这意味着进行覆盖率检测的工具需要维护一组运行时包的列表。这意味着如果有人更改了“runtime”包的依赖关系，在覆盖率构建中将会导致意外的行为。覆盖率运行时会检测并报告这种意外行为，查找如下形式的错误：
// 
// ```plaintext
// 内部错误：覆盖率元数据跟踪：
// 硬编码的运行时包ID列表需要更新。
// 已注册列表：
// 槽位: 0 路径='internal/cpu'  硬编码ID: 1
// 槽位: 1 路径='internal/goarch'  硬编码ID: 2
// 槽位: 2 路径='runtime/internal/atomic'  硬编码ID: 3
// 槽位: 3 路径='internal/goos'
// 槽位: 4 路径='runtime/internal/sys'  硬编码ID: 5
// 槽位: 5 路径='internal/abi'  硬编码ID: 4
// 槽位: 6 路径='runtime/internal/math'  硬编码ID: 6
// 槽位: 7 路径='internal/bytealg'  硬编码ID: 7
// 槽位: 8 路径='internal/goexperiment'
// 槽位: 9 路径='runtime/internal/syscall'  硬编码ID: 8
// 槽位: 10 路径='runtime'  硬编码ID: 9
// 致命错误: runtime.addCovMeta
// ```
// 
// 对于上述错误，硬编码列表中遗漏了“internal/goos”和“internal/goexperiment”。相关开发者需要将上面的列表复制到下方的“rtPkgs”。
// 
// 注意：此策略假设package runtime的依赖关系列表是固定的，并且不随操作系统/架构变化而变化。如果情况并非如此，我们将需要一个表格或其他形式的动态列表来替代当前的固定列表。
# <翻译结束>


<原文开始>
// Scoping note: the constants and apis in this file are internal
// only, not expected to ever be exposed outside of the runtime (unlike
// other coverage file formats and APIs, which will likely be shared
// at some point).
<原文结束>

# <翻译开始>
// 范围说明：此文件中的常量和API仅限内部使用，预计永远不会在运行时外部公开（不像其他覆盖文件格式和API，它们可能在某个时刻会被共享）。
# <翻译结束>


<原文开始>
// NotHardCoded is a package pseudo-ID indicating that a given package
// is not part of the runtime and doesn't require a hard-coded ID.
<原文结束>

# <翻译开始>
// NotHardCoded 是一个包伪标识，表示给定的包不是运行时的一部分，
// 且不需要硬编码的ID。
# <翻译结束>


<原文开始>
// HardCodedPkgID returns the hard-coded ID for the specified package
// path, or -1 if we don't use a hard-coded ID. Hard-coded IDs start
// at -2 and decrease as we go down the list.
<原文结束>

# <翻译开始>
// HardCodedPkgID 返回指定包路径的硬编码ID，如果未使用硬编码ID，则返回-1。硬编码ID从-2开始，并随着我们在列表中向下移动而递减。
# <翻译结束>

