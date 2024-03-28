
<原文开始>
// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
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
// Package testenv provides information about what functionality
// is available in different testing environments run by the Go team.
//
// It is an internal package because these details are specific
// to the Go team's test setup (on build.golang.org) and not
// fundamental to tests in general.
# <翻译结束>


<原文开始>
// Builder reports the name of the builder running this test
// (for example, "linux-amd64" or "windows-386-gce").
// If the test is not running on the build infrastructure,
// Builder returns the empty string.
<原文结束>

# <翻译开始>
// Builder reports the name of the builder running this test
// (for example, "linux-amd64" or "windows-386-gce").
// If the test is not running on the build infrastructure,
// Builder returns the empty string.
# <翻译结束>


<原文开始>
		// It's too much work to require every caller of the go command
		// to pass along "-gcflags="+os.Getenv("GO_GCFLAGS").
		// For now, if $GO_GCFLAGS is set, report that we simply can't
		// run go build.
<原文结束>

# <翻译开始>
		// It's too much work to require every caller of the go command
		// to pass along "-gcflags="+os.Getenv("GO_GCFLAGS").
		// For now, if $GO_GCFLAGS is set, report that we simply can't
		// run go build.
# <翻译结束>


<原文开始>
// For now, having go run and having go build are the same.
<原文结束>

# <翻译开始>
// For now, having go run and having go build are the same.
# <翻译结束>


<原文开始>
// GoToolPath reports the path to the Go tool.
// It is a convenience wrapper around GoTool.
// If the tool is unavailable GoToolPath calls t.Skip.
// If the tool should be available and isn't, GoToolPath calls t.Fatal.
<原文结束>

# <翻译开始>
// GoToolPath reports the path to the Go tool.
// It is a convenience wrapper around GoTool.
// If the tool is unavailable GoToolPath calls t.Skip.
// If the tool should be available and isn't, GoToolPath calls t.Fatal.
# <翻译结束>


<原文开始>
	// Add all environment variables that affect the Go command to test metadata.
	// Cached test results will be invalidate when these variables change.
	// See golang.org/issue/32285.
<原文结束>

# <翻译开始>
	// Add all environment variables that affect the Go command to test metadata.
	// Cached test results will be invalidate when these variables change.
	// See golang.org/issue/32285.
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
			// If runtime.GOROOT() is non-empty, assume that it is valid.
			//
			// (It might not be: for example, the user may have explicitly set GOROOT
			// to the wrong directory, or explicitly set GOROOT_FINAL but not GOROOT
			// and hasn't moved the tree to GOROOT_FINAL yet. But those cases are
			// rare, and if that happens the user can fix what they broke.)
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
<原文结束>

# <翻译开始>
		// runtime.GOROOT doesn't know where GOROOT is (perhaps because the test
		// binary was built with -trimpath, or perhaps because GOROOT_FINAL was set
		// without GOROOT and the tree hasn't been moved there yet).
		//
		// Since this is internal/testenv, we can cheat and assume that the caller
		// is a test of some package in a subdirectory of GOROOT/src. ('go test'
		// runs the test in the directory containing the packaged under test.) That
		// means that if we start walking up the tree, we should eventually find
		// GOROOT/src/go.mod, and we can report the parent directory of that.
# <翻译结束>


<原文开始>
// dir is either "." or only a volume name.
<原文结束>

# <翻译开始>
// dir is either "." or only a volume name.
# <翻译结束>


<原文开始>
// Found "module std", which is the module declaration in GOROOT/src!
<原文结束>

# <翻译开始>
// Found "module std", which is the module declaration in GOROOT/src!
# <翻译结束>


<原文开始>
// GOROOT reports the path to the directory containing the root of the Go
// project source tree. This is normally equivalent to runtime.GOROOT, but
// works even if the test binary was built with -trimpath.
//
// If GOROOT cannot be found, GOROOT skips t if t is non-nil,
// or panics otherwise.
<原文结束>

# <翻译开始>
// GOROOT reports the path to the directory containing the root of the Go
// project source tree. This is normally equivalent to runtime.GOROOT, but
// works even if the test binary was built with -trimpath.
//
// If GOROOT cannot be found, GOROOT skips t if t is non-nil,
// or panics otherwise.
# <翻译结束>


<原文开始>
// GoTool reports the path to the Go tool.
<原文结束>

# <翻译开始>
// GoTool reports the path to the Go tool.
# <翻译结束>


<原文开始>
// HasSrc reports whether the entire source tree is available under GOROOT.
<原文结束>

# <翻译开始>
// HasSrc reports whether the entire source tree is available under GOROOT.
# <翻译结束>


<原文开始>
// HasExternalNetwork reports whether the current system can use
// external (non-localhost) networks.
<原文结束>

# <翻译开始>
// HasExternalNetwork reports whether the current system can use
// external (non-localhost) networks.
# <翻译结束>


<原文开始>
// MustHaveExternalNetwork checks that the current system can use
// external (non-localhost) networks.
// If not, MustHaveExternalNetwork calls t.Skip with an explanation.
<原文结束>

# <翻译开始>
// MustHaveExternalNetwork checks that the current system can use
// external (non-localhost) networks.
// If not, MustHaveExternalNetwork calls t.Skip with an explanation.
# <翻译结束>


<原文开始>
// HasCGO reports whether the current system can use cgo.
<原文结束>

# <翻译开始>
// HasCGO reports whether the current system can use cgo.
# <翻译结束>


<原文开始>
// MustHaveCGO calls t.Skip if cgo is not available.
<原文结束>

# <翻译开始>
// MustHaveCGO calls t.Skip if cgo is not available.
# <翻译结束>


<原文开始>
// CanInternalLink reports whether the current system can link programs with
// internal linking.
<原文结束>

# <翻译开始>
// CanInternalLink reports whether the current system can link programs with
// internal linking.
# <翻译结束>


<原文开始>
// MustInternalLink checks that the current system can link programs with internal
// linking.
// If not, MustInternalLink calls t.Skip with an explanation.
<原文结束>

# <翻译开始>
// MustInternalLink checks that the current system can link programs with internal
// linking.
// If not, MustInternalLink calls t.Skip with an explanation.
# <翻译结束>


<原文开始>
// HasSymlink reports whether the current system can use os.Symlink.
<原文结束>

# <翻译开始>
// HasSymlink reports whether the current system can use os.Symlink.
# <翻译结束>


<原文开始>
// MustHaveSymlink reports whether the current system can use os.Symlink.
// If not, MustHaveSymlink calls t.Skip with an explanation.
<原文结束>

# <翻译开始>
// MustHaveSymlink reports whether the current system can use os.Symlink.
// If not, MustHaveSymlink calls t.Skip with an explanation.
# <翻译结束>


<原文开始>
// HasLink reports whether the current system can use os.Link.
<原文结束>

# <翻译开始>
// HasLink reports whether the current system can use os.Link.
# <翻译结束>


<原文开始>
	// From Android release M (Marshmallow), hard linking files is blocked
	// and an attempt to call link() on a file will return EACCES.
	// - https://code.google.com/p/android-developer-preview/issues/detail?id=3150
<原文结束>

# <翻译开始>
	// From Android release M (Marshmallow), hard linking files is blocked
	// and an attempt to call link() on a file will return EACCES.
	// - https://code.google.com/p/android-developer-preview/issues/detail?id=3150
# <翻译结束>


<原文开始>
// MustHaveLink reports whether the current system can use os.Link.
// If not, MustHaveLink calls t.Skip with an explanation.
<原文结束>

# <翻译开始>
// MustHaveLink reports whether the current system can use os.Link.
// If not, MustHaveLink calls t.Skip with an explanation.
# <翻译结束>


<原文开始>
// CPUIsSlow reports whether the CPU running the test is suspected to be slow.
<原文结束>

# <翻译开始>
// CPUIsSlow reports whether the CPU running the test is suspected to be slow.
# <翻译结束>


<原文开始>
// SkipIfShortAndSlow skips t if -short is set and the CPU running the test is
// suspected to be slow.
//
// (This is useful for CPU-intensive tests that otherwise complete quickly.)
<原文结束>

# <翻译开始>
// SkipIfShortAndSlow skips t if -short is set and the CPU running the test is
// suspected to be slow.
//
// (This is useful for CPU-intensive tests that otherwise complete quickly.)
# <翻译结束>


<原文开始>
// SkipIfOptimizationOff skips t if optimization is disabled.
<原文结束>

# <翻译开始>
// SkipIfOptimizationOff skips t if optimization is disabled.
# <翻译结束>


<原文开始>
// WriteImportcfg writes an importcfg file used by the compiler or linker to
// dstPath containing entries for the packages in std and cmd in addition
// to the package to package file mappings in additionalPackageFiles.
<原文结束>

# <翻译开始>
// WriteImportcfg writes an importcfg file used by the compiler or linker to
// dstPath containing entries for the packages in std and cmd in addition
// to the package to package file mappings in additionalPackageFiles.
# <翻译结束>

