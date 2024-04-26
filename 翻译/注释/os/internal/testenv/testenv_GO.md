
<原文开始>
// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// 版权所有 2015 The Go 作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。
# <翻译结束>


<原文开始>
// Package testenv provides information about what functionality
// is available in different testing environments run by the Go team.
//
// It is an internal package because these details are specific
// to the Go team's test setup (on build.golang.org) and not
// fundamental to tests in general.
<原文结束>

# <翻译开始>
// 包testenv提供了关于Go团队在不同测试环境中可用的功能信息。
// 
// 这是一个内部包，因为这些详细信息是针对Go团队的测试环境（在build.golang.org上）特定的，并非测试的一般原理。
# <翻译结束>


<原文开始>
// Save the original environment during init for use in checks. A test
// binary may modify its environment before calling HasExec to change its
// behavior (such as mimicking a command-line tool), and that modified
// environment might cause environment checks to behave erratically.
<原文结束>

# <翻译开始>
// 在init期间保存原始环境，以供检查时使用。测试二进制文件可能在调用HasExec之前修改其环境，以改变其行为（例如模拟命令行工具），而该修改后的环境可能导致环境检查行为异常。
# <翻译结束>


<原文开始>
// Builder reports the name of the builder running this test
// (for example, "linux-amd64" or "windows-386-gce").
// If the test is not running on the build infrastructure,
// Builder returns the empty string.
<原文结束>

# <翻译开始>
// Builder 返回运行此测试的构建器的名称（例如，“linux-amd64”或“windows-386-gce”）。
// 如果测试不在构建基础设施上运行，Builder 将返回空字符串。
# <翻译结束>


<原文开始>
		// It's too much work to require every caller of the go command
		// to pass along "-gcflags="+os.Getenv("GO_GCFLAGS").
		// For now, if $GO_GCFLAGS is set, report that we simply can't
		// run go build.
<原文结束>

# <翻译开始>
// 要求每次调用`go`命令的用户都传递`"-gcflags="+os.Getenv("GO_GCFLAGS")`太麻烦了。
// 当前，如果`$GO_GCFLAGS`已设置，我们报告无法运行`go build`。
# <翻译结束>


<原文开始>
		// To run 'go build', we need to be able to exec a 'go' command.
		// We somewhat arbitrarily choose to exec 'go tool -n compile' because that
		// also confirms that cmd/go can find the compiler. (Before CL 472096,
		// we sometimes ended up with cmd/go installed in the test environment
		// without a cmd/compile it could use to actually build things.)
<原文结束>

# <翻译开始>
// 要运行`go build`，我们需要能够执行一个`go`命令。
// 我们有点随意地选择执行`go tool -n compile`，因为这也确认了cmd/go可以找到编译器。（在CL 472096之前，
// 我们有时会在测试环境中安装没有实际用于构建东西的cmd/go，而它缺少cmd/compile。）
# <翻译结束>


<原文开始>
			// We can assume that we always have a complete Go toolchain available.
			// However, this platform requires a C linker to build even pure Go
			// programs, including tests. Do we have one in the test environment?
			// (On Android, for example, the device running the test might not have a
			// C toolchain installed.)
			//
			// If CC is set explicitly, assume that we do. Otherwise, use 'go env CC'
			// to determine which toolchain it would use by default.
<原文结束>

# <翻译开始>
// 我们可以认为我们总是拥有完整的 Go 工具链可用。
// 然而，此平台要求即使构建纯 Go 程序（包括测试）也需要 C 链接器。测试环境中是否具备？
// （例如在 Android 上，运行测试的设备可能没有安装 C 工具链。）
//
// 如果显式设置了 CC，则假定我们具备。否则，使用 'go env CC' 来确定它将默认使用哪个工具链。
# <翻译结束>


<原文开始>
// For now, having go run and having go build are the same.
<原文结束>

# <翻译开始>
// 目前，使用"go run"和使用"go build"是相同的。
# <翻译结束>


<原文开始>
// HasParallelism reports whether the current system can execute multiple
// threads in parallel.
// There is a copy of this function in cmd/dist/test.go.
<原文结束>

# <翻译开始>
// HasParallelism报告当前系统是否可以并行执行多个线程。
// 在cmd/dist/test.go中有一个此函数的副本。
# <翻译结束>


<原文开始>
// MustHaveParallelism checks that the current system can execute multiple
// threads in parallel. If not, MustHaveParallelism calls t.Skip with an explanation.
<原文结束>

# <翻译开始>
// MustHaveParallelism 检查当前系统是否可以同时执行多个线程。如果不可以，MustHaveParallelism 会调用 t.Skip 并附带一个解释。
# <翻译结束>


<原文开始>
// GoToolPath reports the path to the Go tool.
// It is a convenience wrapper around GoTool.
// If the tool is unavailable GoToolPath calls t.Skip.
// If the tool should be available and isn't, GoToolPath calls t.Fatal.
<原文结束>

# <翻译开始>
// GoToolPath 返回 Go 工具的路径。
// 它是对 GoTool 的便捷包装。
// 若工具不可用，GoToolPath 将调用 t.Skip。
// 若工具本应可用但实际不可用，GoToolPath 将调用 t.Fatal。
# <翻译结束>


<原文开始>
	// Add all environment variables that affect the Go command to test metadata.
	// Cached test results will be invalidate when these variables change.
	// See golang.org/issue/32285.
<原文结束>

# <翻译开始>
// 将所有影响Go命令的环境变量添加到测试元数据中。
// 当这些变量改变时，缓存的测试结果将被失效。
// 参见golang.org/issue/32285。
# <翻译结束>


<原文开始>
			// If runtime.GOROOT() is non-empty, assume that it is valid.
			//
			// (It might not be: for example, the user may have explicitly set GOROOT
			// to the wrong directory, or explicitly set GOROOT_FINAL but not GOROOT
			// and hasn't moved the tree to GOROOT_FINAL yet. But those cases are
			// rare, and if that happens the user can fix what they broke.)
<原文结束>

# <翻译开始>
// 如果runtime.GOROOT()非空，假设它是有效的。
// 
// (它可能不是：例如，用户可能已明确将GOROOT设置为错误的目录，或者明确设置了GOROOT_FINAL但没有设置GOROOT，并且还没有将树移动到GOROOT_FINAL。但这些情况很少见，如果发生这种情况，用户可以修复他们破坏的东西。)
# <翻译结束>


<原文开始>
		// runtime.GOROOT doesn't know where GOROOT is (perhaps because the test
		// binary was built with -trimpath, or perhaps because GOROOT_FINAL was set
		// without GOROOT and the tree hasn't been moved there yet).
		//
		// Since this is internal/testenv, we can cheat and assume that the caller
		// is a test of some package in a subdirectory of GOROOT/src. ('go test'
		// runs the test in the directory containing the packaged under test.) That
		// means that if we start walking up the tree, we should eventually find
		// GOROOT/src/go.mod, and we can report the parent directory of that.
		//
		// Notably, this works even if we can't run 'go env GOROOT' as a
		// subprocess.
<原文结束>

# <翻译开始>
// runtime.GOROOT 不知道 GOROOT 在哪里（可能是因为测试二进制文件是使用 -trimpath 构建的，或者是因为设置了 GOROOT_FINAL 但没有将树移动到该位置）。
// 
// 由于这是 internal/testenv，我们可以作弊并假设调用者是 GOROOT/src 子目录中某个包的测试。`go test`会在包含被测试包的目录中运行测试。这意味着如果我们开始向上遍历树，最终应该找到 GOROOT/src/go.mod，我们可以报告出那个目录的父目录。
// 
// 值得注意的是，即使我们不能作为子进程运行 `go env GOROOT`，这仍然有效。
# <翻译结束>


<原文开始>
// dir is either "." or only a volume name.
<原文结束>

# <翻译开始>
// dir 是 "." 或仅为卷名。
# <翻译结束>


<原文开始>
// Found "module std", which is the module declaration in GOROOT/src!
<原文结束>

# <翻译开始>
// 发现了"module std"，这是在GOROOT/src中的模块声明！
# <翻译结束>


<原文开始>
// GOROOT reports the path to the directory containing the root of the Go
// project source tree. This is normally equivalent to runtime.GOROOT, but
// works even if the test binary was built with -trimpath and cannot exec
// 'go env GOROOT'.
//
// If GOROOT cannot be found, GOROOT skips t if t is non-nil,
// or panics otherwise.
<原文结束>

# <翻译开始>
// GOROOT 会报告包含 Go 项目源代码树根目录的路径。这通常等同于 runtime.GOROOT，但在使用 -trimpath 构建测试二进制文件且无法执行 'go env GOROOT' 的情况下仍然有效。
// 
// 如果找不到 GOROOT，如果 t 不为 nil，则跳过 t，否则会 panic。
# <翻译结束>


<原文开始>
// GoTool reports the path to the Go tool.
<原文结束>

# <翻译开始>
// GoTool报告Go工具的路径。
# <翻译结束>


<原文开始>
// HasSrc reports whether the entire source tree is available under GOROOT.
<原文结束>

# <翻译开始>
// HasSrc 报告整个源代码树是否可在 GOROOT 下获取。
# <翻译结束>


<原文开始>
// HasExternalNetwork reports whether the current system can use
// external (non-localhost) networks.
<原文结束>

# <翻译开始>
// HasExternalNetwork 报告当前系统是否可以使用外部（非本地回环）网络。
# <翻译结束>


<原文开始>
// MustHaveExternalNetwork checks that the current system can use
// external (non-localhost) networks.
// If not, MustHaveExternalNetwork calls t.Skip with an explanation.
<原文结束>

# <翻译开始>
// MustHaveExternalNetwork 检查当前系统是否可以使用外部（非localhost）网络。
// 如果不能，MustHaveExternalNetwork 会调用 t.Skip 并附带一个解释。
# <翻译结束>


<原文开始>
// HasCGO reports whether the current system can use cgo.
<原文结束>

# <翻译开始>
// HasCGO报告当前系统是否可以使用cgo。
# <翻译结束>


<原文开始>
// MustHaveCGO calls t.Skip if cgo is not available.
<原文结束>

# <翻译开始>
// MustHaveCGO 若无 cgo 可用则调用 t.Skip
# <翻译结束>


<原文开始>
// CanInternalLink reports whether the current system can link programs with
// internal linking.
<原文结束>

# <翻译开始>
// CanInternalLink 报告当前系统是否支持使用内部链接来链接程序。
# <翻译结束>


<原文开始>
// MustInternalLink checks that the current system can link programs with internal
// linking.
// If not, MustInternalLink calls t.Skip with an explanation.
<原文结束>

# <翻译开始>
// MustInternalLink 检查当前系统是否支持内部链接程序。
// 如果不支持，MustInternalLink 会调用 t.Skip 并附带一个解释。
# <翻译结束>


<原文开始>
// MustHaveBuildMode reports whether the current system can build programs in
// the given build mode.
// If not, MustHaveBuildMode calls t.Skip with an explanation.
<原文结束>

# <翻译开始>
// MustHaveBuildMode 检查当前系统是否能够在给定的构建模式下构建程序。
// 如果不能，MustHaveBuildMode 将调用 t.Skip 并附带一个解释。
# <翻译结束>


<原文开始>
// HasSymlink reports whether the current system can use os.Symlink.
<原文结束>

# <翻译开始>
// HasSymlink 报告当前系统是否可以使用 os.Symlink。
# <翻译结束>


<原文开始>
// MustHaveSymlink reports whether the current system can use os.Symlink.
// If not, MustHaveSymlink calls t.Skip with an explanation.
<原文结束>

# <翻译开始>
// MustHaveSymlink 检查当前系统是否支持使用 os.Symlink。
// 如果不支持，MustHaveSymlink 将调用 t.Skip 并附带解释。
# <翻译结束>


<原文开始>
// HasLink reports whether the current system can use os.Link.
<原文结束>

# <翻译开始>
// HasLink报告当前系统是否可以使用os.Link。
# <翻译结束>


<原文开始>
	// From Android release M (Marshmallow), hard linking files is blocked
	// and an attempt to call link() on a file will return EACCES.
	// - https://code.google.com/p/android-developer-preview/issues/detail?id=3150
<原文结束>

# <翻译开始>
// 从Android M（棉花糖）版本开始，硬链接文件被阻止，并尝试对文件调用link()会返回EACCES错误。
// - https://code.google.com/p/android-developer-preview/issues/detail?id=3150
# <翻译结束>


<原文开始>
// MustHaveLink reports whether the current system can use os.Link.
// If not, MustHaveLink calls t.Skip with an explanation.
<原文结束>

# <翻译开始>
// MustHaveLink 判断当前系统是否支持使用 os.Link。
// 若不支持，MustHaveLink 会调用 t.Skip 并附带解释原因。
# <翻译结束>


<原文开始>
// CPUIsSlow reports whether the CPU running the test is suspected to be slow.
<原文结束>

# <翻译开始>
// CPUIsSlow 报告测试运行的CPU是否疑似性能较慢。
# <翻译结束>


<原文开始>
// SkipIfShortAndSlow skips t if -short is set and the CPU running the test is
// suspected to be slow.
//
// (This is useful for CPU-intensive tests that otherwise complete quickly.)
<原文结束>

# <翻译开始>
// SkipIfShortAndSlow 如果设置了 `-short` 并且怀疑运行测试的 CPU 性能较慢，则跳过 `t`。
//
// （这对于那些通常很快完成但需要大量 CPU 资源的测试很有用。）
# <翻译结束>


<原文开始>
// SkipIfOptimizationOff skips t if optimization is disabled.
<原文结束>

# <翻译开始>
// SkipIfOptimizationOff 如果优化被关闭，则跳过t。
# <翻译结束>


<原文开始>
// WriteImportcfg writes an importcfg file used by the compiler or linker to
// dstPath containing entries for the file mappings in packageFiles, as well
// as for the packages transitively imported by the package(s) in pkgs.
//
// pkgs may include any package pattern that is valid to pass to 'go list',
// so it may also be a list of Go source files all in the same directory.
<原文结束>

# <翻译开始>
// WriteImportcfg编写一个由编译器或链接器使用的importcfg文件，
// 并将其写入dstPath。该文件包含packageFiles中文件映射的条目，
// 同时也包含由pkgs中包（们）递归导入的所有包的信息。
// 
// pkgs可以包含任何可合法传递给'go list'命令的包模式，
// 因此它也可以是一组位于同一目录下的Go源文件列表。
# <翻译结束>


<原文开始>
// Use 'go list' to resolve any missing packages and rewrite the import map.
<原文结束>

# <翻译开始>
// 使用 'go list' 命令来解析任何缺失的包，并重写导入映射。
# <翻译结束>


<原文开始>
// SyscallIsNotSupported reports whether err may indicate that a system call is
// not supported by the current platform or execution environment.
<原文结束>

# <翻译开始>
// SyscallIsNotSupported 检查错误是否可能表示当前平台或执行环境不支持系统调用。
# <翻译结束>

