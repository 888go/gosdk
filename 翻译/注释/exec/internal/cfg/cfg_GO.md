
<原文开始>
// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// 版权所有 ? 2019 The Go Authors。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。
# <翻译结束>


<原文开始>
// Package cfg holds configuration shared by the Go command and internal/testenv.
// Definitions that don't need to be exposed outside of cmd/go should be in
// cmd/go/internal/cfg instead of this package.
<原文结束>

# <翻译开始>
// Package cfg 存放了由 Go 命令和 internal/testenv 共享的配置。
// 不需要对外部的 cmd/go 暴露的定义应位于 cmd/go/internal/cfg 中，而非此包中。
# <翻译结束>


<原文开始>
// KnownEnv is a list of environment variables that affect the operation
// of the Go command.
<原文结束>

# <翻译开始>
// KnownEnv 是一个环境变量列表，这些变量会影响 Go 命令的运行。
# <翻译结束>

