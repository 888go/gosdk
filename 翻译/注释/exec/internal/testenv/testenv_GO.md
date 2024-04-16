
<原文开始>
// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// 版权所有 ? 2015 The Go Authors。保留所有权利。
// 本源代码的使用受 BSD 风格许可协议约束，
// 该协议可在 LICENSE 文件中找到。
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
// Package testenv 提供关于不同测试环境中由 Go 团队运行的功能可用性的信息。
//
// 该包为内部包，因为这些细节特定于 Go 团队的测试设置（在 build.golang.org 上），而非针对一般测试的基础性内容。
# <翻译结束>


<原文开始>
// Save the original environment during init for use in checks. A test
// binary may modify its environment before calling HasExec to change its
// behavior (such as mimicking a command-line tool), and that modified
// environment might cause environment checks to behave erratically.
<原文结束>

# <翻译开始>
// 在初始化时保存原始环境，供检查时使用。测试二进制文件可能在调用HasExec之前修改其环境以改变其行为（例如模拟命令行工具），而该修改后的环境可能导致环境检查行为异常。
# <翻译结束>


<原文开始>
// Builder reports the name of the builder running this test
// (for example, "linux-amd64" or "windows-386-gce").
// If the test is not running on the build infrastructure,
// Builder returns the empty string.
<原文结束>

# <翻译开始>
// Builder 返回运行此测试的构建器名称（例如，“linux-amd64”或“windows-386-gce”）。
// 如果该测试未在构建基础设施上运行，则 Builder 返回空字符串。
# <翻译结束>


<原文开始>
		// It's too much work to require every caller of the go command
		// to pass along "-gcflags="+os.Getenv("GO_GCFLAGS").
		// For now, if $GO_GCFLAGS is set, report that we simply can't
		// run go build.
<原文结束>

# <翻译开始>
// 要求 go 命令的每个调用者都传递 "-gcflags="+os.Getenv("GO_GCFLAGS") 给我们带来了太多工作。
// 目前，如果设置了 $GO_GCFLAGS，就报告我们无法运行 go build。
# <翻译结束>


<原文开始>
		// To run 'go build', we need to be able to exec a 'go' command.
		// We somewhat arbitrarily choose to exec 'go tool -n compile' because that
		// also confirms that cmd/go can find the compiler. (Before CL 472096,
		// we sometimes ended up with cmd/go installed in the test environment
		// without a cmd/compile it could use to actually build things.)
<原文结束>

# <翻译开始>
// 要运行 'go build'，我们需要能够执行一个 'go' 命令。
// 我们选择执行 'go tool -n compile' 是出于某种考虑，因为这也将确认 cmd/go 能够找到编译器。（在 CL 472096 之前，
// 我们有时会在测试环境中遇到仅安装了 cmd/go，但没有可用的 cmd/compile 来实际构建项目的情况。）
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
// 我们可以假定我们始终拥有一个完整的 Go 工具链可用。
// 然而，该平台要求即使构建纯 Go 程序（包括测试）也需要一个 C 编译器。在测试环境中我们是否拥有一个？
// （例如，在 Android 上，运行测试的设备可能没有安装 C 工具链。）
//
// 如果已明确设置了 CC，则假定我们拥有。否则，使用 'go env CC' 来确定它将默认使用哪个工具链。
# <翻译结束>


<原文开始>
// For now, having go run and having go build are the same.
<原文结束>

# <翻译开始>
// 目前，执行`go run`和`go build`是等效的。
# <翻译结束>


<原文开始>
// HasParallelism reports whether the current system can execute multiple
// threads in parallel.
// There is a copy of this function in cmd/dist/test.go.
<原文结束>

# <翻译开始>
// HasParallelism 判断当前系统是否支持并行执行多个线程。
// 在 cmd/dist/test.go 中存在此函数的副本。
# <翻译结束>


<原文开始>
// MustHaveParallelism checks that the current system can execute multiple
// threads in parallel. If not, MustHaveParallelism calls t.Skip with an explanation.
<原文结束>

# <翻译开始>
// MustHaveParallelism 检查当前系统是否能够并行执行多个线程。如果不能，MustHaveParallelism 将调用 t.Skip 并附带解释原因。
# <翻译结束>


<原文开始>
// GoToolPath reports the path to the Go tool.
// It is a convenience wrapper around GoTool.
// If the tool is unavailable GoToolPath calls t.Skip.
// If the tool should be available and isn't, GoToolPath calls t.Fatal.
<原文结束>

# <翻译开始>
// GoToolPath 返回 Go 工具的路径。
// 它是 GoTool 的便捷包装器。
// 如果工具不可用，GoToolPath 将调用 t.Skip。
// 如果工具本应可用但实际不可用，GoToolPath 将调用 t.Fatal。
# <翻译结束>


<原文开始>
	// Add all environment variables that affect the Go command to test metadata.
	// Cached test results will be invalidate when these variables change.
	// See golang.org/issue/32285.
<原文结束>

# <翻译开始>
// 将影响 Go 命令的所有环境变量添加到测试元数据中。
// 当这些变量发生变化时，缓存的测试结果将被失效。
// 参见 golang.org/issue/32285。
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
// 如果 runtime.GOROOT() 非空，假定其值有效。
//
// （实际情况可能并非如此：例如，用户可能已将 GOROOT 明确设置为错误的目录，或者明确设置了 GOROOT_FINAL 但未设置 GOROOT，并且尚未将树移动到 GOROOT_FINAL。但这些情况较为罕见，若发生此类问题，用户可自行修复。）
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
// runtime.GOROOT 无法确定 GOROOT 的位置（可能是因为测试二进制文件是使用 -trimpath 参数构建的，或者是因为仅设置了 GOROOT_FINAL 而未设置 GOROOT，且树状结构尚未移动到该位置）。
//
// 由于这里是 internal/testenv，我们可以采取捷径并假设调用者是对 GOROOT/src 某子目录下某个包进行的测试（'go test' 在包含被测包的目录中运行测试）。这意味着如果我们开始沿目录树向上遍历，最终应能找到 GOROOT/src/go.mod，并可以报告该文件的父目录。
//
// 值得注意的是，即使我们不能以子进程形式运行 'go env GOROOT'，这种方法依然有效。
# <翻译结束>


<原文开始>
// dir is either "." or only a volume name.
<原文结束>

# <翻译开始>
// dir 是 "." 或仅为一个卷名。
# <翻译结束>


<原文开始>
// Found "module std", which is the module declaration in GOROOT/src!
<原文结束>

# <翻译开始>
// 发现了“module std”，即GOROOT/src中的模块声明！
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
// GOROOT 返回包含Go项目源代码树根目录的路径。通常情况下，这与runtime.GOROOT等价，但在测试二进制文件使用-trimpath选项构建且无法执行"go env GOROOT"命令时仍可正常工作。
// 
// 若无法找到GOROOT，当t非nil时，GOROOT将跳过对t的操作，否则将引发恐慌。
# <翻译结束>


<原文开始>
// GoTool reports the path to the Go tool.
<原文结束>

# <翻译开始>
// GoTool 报告 Go 工具的路径。
# <翻译结束>


<原文开始>
// HasSrc reports whether the entire source tree is available under GOROOT.
<原文结束>

# <翻译开始>
// HasSrc 判断整个源代码树是否均可在 GOROOT 下获取到
# <翻译结束>


<原文开始>
// HasExternalNetwork reports whether the current system can use
// external (non-localhost) networks.
<原文结束>

# <翻译开始>
// HasExternalNetwork 判断当前系统是否可以使用非本地主机（外部）网络
# <翻译结束>


<原文开始>
// MustHaveExternalNetwork checks that the current system can use
// external (non-localhost) networks.
// If not, MustHaveExternalNetwork calls t.Skip with an explanation.
<原文结束>

# <翻译开始>
// MustHaveExternalNetwork 检查当前系统是否可以使用
// 外部（非本地主机）网络。
// 如果不能，MustHaveExternalNetwork 将调用 t.Skip 并附带解释。
# <翻译结束>


<原文开始>
// HasCGO reports whether the current system can use cgo.
<原文结束>

# <翻译开始>
// HasCGO 报告当前系统是否可以使用 cgo。
# <翻译结束>


<原文开始>
// MustHaveCGO calls t.Skip if cgo is not available.
<原文结束>

# <翻译开始>
// MustHaveCGO 若未启用cgo，则调用t.Skip。
# <翻译结束>


<原文开始>
// CanInternalLink reports whether the current system can link programs with
// internal linking.
<原文结束>

# <翻译开始>
// CanInternalLink 判断当前系统是否支持以内部链接方式链接程序
# <翻译结束>


<原文开始>
// MustInternalLink checks that the current system can link programs with internal
// linking.
// If not, MustInternalLink calls t.Skip with an explanation.
<原文结束>

# <翻译开始>
// MustInternalLink 检查当前系统是否支持以内部链接方式链接程序。
// 若不支持，MustInternalLink 将调用 t.Skip 并附带解释原因。
# <翻译结束>


<原文开始>
// MustHaveBuildMode reports whether the current system can build programs in
// the given build mode.
// If not, MustHaveBuildMode calls t.Skip with an explanation.
<原文结束>

# <翻译开始>
// MustHaveBuildMode 检查当前系统是否能够在给定的构建模式下构建程序。
// 如果不能，MustHaveBuildMode 会调用 t.Skip 并附带解释原因。
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
// MustHaveSymlink 检查当前系统是否支持使用 os.Symlink。如果不支持，MustHaveSymlink 会调用 t.Skip 并附带解释原因。
# <翻译结束>


<原文开始>
// HasLink reports whether the current system can use os.Link.
<原文结束>

# <翻译开始>
// HasLink 报告当前系统是否可以使用 os.Link。
# <翻译结束>


<原文开始>
	// From Android release M (Marshmallow), hard linking files is blocked
	// and an attempt to call link() on a file will return EACCES.
	// - https://code.google.com/p/android-developer-preview/issues/detail?id=3150
<原文结束>

# <翻译开始>
// 从Android版本M（Marshmallow）开始，系统禁止对文件进行硬链接操作，
// 若尝试调用link()函数对文件进行硬链接，将会返回EACCES错误。
// - 参考：https://code.google.com/p/android-developer-preview/issues/detail?id=3150
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
// CPUIsSlow 报告当前运行测试的 CPU 是否疑似性能较低
# <翻译结束>


<原文开始>
// SkipIfShortAndSlow skips t if -short is set and the CPU running the test is
// suspected to be slow.
//
// (This is useful for CPU-intensive tests that otherwise complete quickly.)
<原文结束>

# <翻译开始>
// SkipIfShortAndSlow 在设置了 -short 参数且运行测试的 CPU 被认为较慢时，跳过 t。
//
// （这对于那些在 CPU 强劲时能快速完成，但实际消耗 CPU 资源较多的测试非常有用。）
# <翻译结束>


<原文开始>
// SkipIfOptimizationOff skips t if optimization is disabled.
<原文结束>

# <翻译开始>
// 如果优化已禁用，则SkipIfOptimizationOff会跳过t。
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
// WriteImportcfg编写一个importcfg文件，该文件由编译器或链接器用于
// dstPath，其中包含packageFiles中文件映射的条目，以及
// 由pkgs中的包（们）递归导入的包。
// 
// pkgs可以包含任何可有效传递给'go list'的包模式，
// 因此它也可以是同一目录下所有Go源文件的列表。
# <翻译结束>


<原文开始>
// Use 'go list' to resolve any missing packages and rewrite the import map.
<原文结束>

# <翻译开始>
// 使用'go list'命令解析所有缺失的包并重写导入映射。
# <翻译结束>


<原文开始>
// SyscallIsNotSupported reports whether err may indicate that a system call is
// not supported by the current platform or execution environment.
<原文结束>

# <翻译开始>
// SyscallIsNotSupported 判断 err 是否可能表示当前平台或执行环境中不支持某个系统调用。
# <翻译结束>

