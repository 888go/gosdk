
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
	// Tests are defined to run within their package source directory,
	// and this package's source directory is $GOROOT/src/internal/testenv.
	// The 'go' command is installed at $GOROOT/bin/go, so if the environment
	// is correct then testenv.GoTool() should be identical to ../../../bin/go.
<原文结束>

# <翻译开始>
// 测试定义为在其包源目录中运行，
// 而此包的源目录是 $GOROOT/src/internal/testenv。
// 'go' 命令安装在 $GOROOT/bin/go，因此如果环境设置正确，
// 则 testenv.GoTool() 应该与 ../../../bin/go 完全相同。
# <翻译结束>


<原文开始>
// No exec syscall, so these shouldn't be able to 'go build'.
<原文结束>

# <翻译开始>
// 无 exec 系统调用，因此它们不应能通过 'go build' 构建。
# <翻译结束>


<原文开始>
			// We shouldn't make assumptions about what kind of sandbox or build
			// environment external Go users may be running in.
<原文结束>

# <翻译开始>
// 我们不应对外部 Go 用户可能运行的沙盒或构建环境类型做出任何假设。
# <翻译结束>


<原文开始>
		// Since we control the Go builders, we know which ones ought
		// to be able to run 'go build'. Check that they can.
		//
		// (Note that we don't verify that any builders *can't* run 'go build'.
		// If a builder starts running 'go build' tests when it shouldn't,
		// we will presumably find out about it when those tests fail.)
<原文结束>

# <翻译开始>
// 由于我们掌控着Go构建器，我们知道哪些构建器应当能够运行`go build`。检查它们是否确实可以。
// 
// （注意，我们并不验证任何构建器*不能*运行`go build`。如果某个构建器在不应运行`go build`测试的情况下开始运行这些测试，我们推测当那些测试失败时，我们会发现这个问题。）
# <翻译结束>


<原文开始>
				// The corellium environment is self-hosting, so it should be able
				// to build even though real "ios" devices can't exec.
<原文结束>

# <翻译开始>
// Corellium 环境是自承载的，因此即使真实的“ios”设备无法执行，也应该能够进行构建。
# <翻译结束>


<原文开始>
				// The usual iOS sandbox does not allow the app to start another
				// process. If we add builders on stock iOS devices, they presumably
				// will not be able to exec, so we may as well allow that now.
<原文结束>

# <翻译开始>
// 在通常的iOS沙盒环境中，应用程序不允许启动另一个进程。如果我们要在原生iOS设备上添加构建器，它们大概率将无法执行exec操作，因此我们现在不妨允许这种情况发生。
# <翻译结束>


<原文开始>
				// As of 2023-05-02, the test environment on the emulated builders is
				// missing a C linker.
<原文结束>

# <翻译开始>
// 截至2023年5月2日，模拟构建器上的测试环境中缺少C链接器。
# <翻译结束>


<原文开始>
			// The -noopt builder sets GO_GCFLAGS, which causes tests of 'go build' to
			// be skipped.
<原文结束>

# <翻译开始>
// -noopt构建器设置了GO_GCFLAGS，导致对'go build'的测试被跳过。
# <翻译结束>


<原文开始>
// Most ios environments can't exec, but the corellium builder can.
<原文结束>

# <翻译开始>
// 大多数iOS环境无法执行，但Corellium构建器可以。
# <翻译结束>


<原文开始>
// Test that CleanCmdEnv sets PWD if cmd.Dir is set.
<原文结束>

# <翻译开始>
// 测试当 cmd.Dir 被设置时，CleanCmdEnv 是否会正确设置 PWD。
# <翻译结束>

