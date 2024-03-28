
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
// Package buildcfg provides access to the build configuration
// described by the current environment. It is for use by build tools
// such as cmd/go or cmd/compile and for setting up go/build's Default context.
//
// Note that it does NOT provide access to the build configuration used to
// build the currently-running binary. For that, use runtime.GOOS etc
// as well as internal/goexperiment.
<原文结束>

# <翻译开始>
// 包buildcfg提供了对当前环境所描述的构建配置的访问权限，适用于诸如cmd/go或cmd/compile等构建工具，并用于设置go/build的默认上下文。
//
// 注意，它并不提供用于构建当前运行二进制文件的构建配置的访问权限。为此，请使用runtime.GOOS等以及internal/goexperiment。
# <翻译结束>


<原文开始>
// Error is one of the errors found (if any) in the build configuration.
<原文结束>

# <翻译开始>
// Error 是在构建配置中找到的错误之一（如果存在的话）。
# <翻译结束>


<原文开始>
// Check exits the program with a fatal error if Error is non-nil.
<原文结束>

# <翻译开始>
// Check在Error非空时，将以致命错误退出程序。
# <翻译结束>


<原文开始>
// Android arm devices always support GOARM=7.
<原文结束>

# <翻译开始>
// Android ARM 设备始终支持 GOARM=7。
# <翻译结束>


<原文开始>
	// For each experiment that has been enabled in the toolchain, define a
	// build tag with the same name but prefixed by "goexperiment." which can be
	// used for compiling alternative files for the experiment. This allows
	// changes for the experiment, like extra struct fields in the runtime,
	// without affecting the base non-experiment code at all.
<原文结束>

# <翻译开始>
// 对于工具链中已启用的每个实验，定义一个以"goexperiment."开头、后接相同名称的构建标签。这个标签可用于为该实验编译替代文件。这样做可以在不影响基础非实验代码的情况下，使实验（如运行时中额外的结构体字段）能够进行变更。
# <翻译结束>


<原文开始>
// GOGOARCH returns the name and value of the GO$GOARCH setting.
// For example, if GOARCH is "amd64" it might return "GOAMD64", "v2".
<原文结束>

# <翻译开始>
// GOGOARCH 返回 GO$GOARCH 设置的名称和值。
// 例如，如果 GOARCH 为 "amd64"，它可能会返回 "GOAMD64", "v2"。
# <翻译结束>

