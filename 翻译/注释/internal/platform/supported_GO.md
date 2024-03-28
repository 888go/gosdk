
<原文开始>
// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
# <翻译结束>


<原文开始>
// RaceDetectorSupported reports whether goos/goarch supports the race
// detector. There is a copy of this function in cmd/dist/test.go.
// Race detector only supports 48-bit VMA on arm64. But it will always
// return true for arm64, because we don't have VMA size information during
// the compile time.
<原文结束>

# <翻译开始>
// RaceDetectorSupported reports whether goos/goarch supports the race
// detector. There is a copy of this function in cmd/dist/test.go.
// Race detector only supports 48-bit VMA on arm64. But it will always
// return true for arm64, because we don't have VMA size information during
// the compile time.
# <翻译结束>


<原文开始>
// MSanSupported reports whether goos/goarch supports the memory
// sanitizer option.
// There is a copy of this function in misc/cgo/testsanitizers/cc_test.go.
<原文结束>

# <翻译开始>
// MSanSupported reports whether goos/goarch supports the memory
// sanitizer option.
// There is a copy of this function in misc/cgo/testsanitizers/cc_test.go.
# <翻译结束>


<原文开始>
// ASanSupported reports whether goos/goarch supports the address
// sanitizer option.
// There is a copy of this function in misc/cgo/testsanitizers/cc_test.go.
<原文结束>

# <翻译开始>
// ASanSupported reports whether goos/goarch supports the address
// sanitizer option.
// There is a copy of this function in misc/cgo/testsanitizers/cc_test.go.
# <翻译结束>


<原文开始>
// FuzzSupported reports whether goos/goarch supports fuzzing
// ('go test -fuzz=.').
<原文结束>

# <翻译开始>
// FuzzSupported reports whether goos/goarch supports fuzzing
// ('go test -fuzz=.').
# <翻译结束>


<原文开始>
// FuzzInstrumented reports whether fuzzing on goos/goarch uses coverage
// instrumentation. (FuzzInstrumented implies FuzzSupported.)
<原文结束>

# <翻译开始>
// FuzzInstrumented reports whether fuzzing on goos/goarch uses coverage
// instrumentation. (FuzzInstrumented implies FuzzSupported.)
# <翻译结束>


<原文开始>
// TODO(#14565): support more architectures.
<原文结束>

# <翻译开始>
// TODO(#14565): support more architectures.
# <翻译结束>


<原文开始>
// MustLinkExternal reports whether goos/goarch requires external linking.
<原文结束>

# <翻译开始>
// MustLinkExternal reports whether goos/goarch requires external linking.
# <翻译结束>


<原文开始>
// MustLinkExternalGo121 reports whether goos/goarch requires external linking,
// with or without cgo dependencies. [This version back-ported from
// Go 1.21 as part of a test].
<原文结束>

# <翻译开始>
// MustLinkExternalGo121 reports whether goos/goarch requires external linking,
// with or without cgo dependencies. [This version back-ported from
// Go 1.21 as part of a test].
# <翻译结束>


<原文开始>
			// Internally linking cgo is incomplete on some architectures.
			// https://go.dev/issue/14449
<原文结束>

# <翻译开始>
			// Internally linking cgo is incomplete on some architectures.
			// https://go.dev/issue/14449
# <翻译结束>


<原文开始>
// windows/arm64 internal linking is not implemented.
<原文结束>

# <翻译开始>
// windows/arm64 internal linking is not implemented.
# <翻译结束>


<原文开始>
			// Big Endian PPC64 cgo internal linking is not implemented for aix or linux.
			// https://go.dev/issue/8912
<原文结束>

# <翻译开始>
			// Big Endian PPC64 cgo internal linking is not implemented for aix or linux.
			// https://go.dev/issue/8912
# <翻译结束>


<原文开始>
			// It seems that on Dragonfly thread local storage is
			// set up by the dynamic linker, so internal cgo linking
			// doesn't work. Test case is "go test runtime/cgo".
<原文结束>

# <翻译开始>
			// It seems that on Dragonfly thread local storage is
			// set up by the dynamic linker, so internal cgo linking
			// doesn't work. Test case is "go test runtime/cgo".
# <翻译结束>


<原文开始>
// BuildModeSupported reports whether goos/goarch supports the given build mode
// using the given compiler.
<原文结束>

# <翻译开始>
// BuildModeSupported reports whether goos/goarch supports the given build mode
// using the given compiler.
# <翻译结束>


<原文开始>
		// TODO(bcmills): This seems dubious.
		// Do we really support c-archive mode on js/wasm?
<原文结束>

# <翻译开始>
		// TODO(bcmills): This seems dubious.
		// Do we really support c-archive mode on js/wasm?
# <翻译结束>

