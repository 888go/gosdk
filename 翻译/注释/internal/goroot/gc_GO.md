
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
// IsStandardPackage reports whether path is a standard package,
// given goroot and compiler.
<原文结束>

# <翻译开始>
// IsStandardPackage reports whether path is a standard package,
// given goroot and compiler.
# <翻译结束>


<原文开始>
// gccgoSearch holds the gccgo search directories.
<原文结束>

# <翻译开始>
// gccgoSearch holds the gccgo search directories.
# <翻译结束>


<原文开始>
// gccgoSearch is used to check whether a gccgo package exists in the
// standard library.
<原文结束>

# <翻译开始>
// gccgoSearch is used to check whether a gccgo package exists in the
// standard library.
# <翻译结束>


<原文开始>
// init finds the gccgo search directories. If this fails it leaves dirs == nil.
<原文结束>

# <翻译开始>
// init finds the gccgo search directories. If this fails it leaves dirs == nil.
# <翻译结束>


<原文开始>
// isStandard reports whether path is a standard library for gccgo.
<原文结束>

# <翻译开始>
// isStandard reports whether path is a standard library for gccgo.
# <翻译结束>


<原文开始>
	// Quick check: if the first path component has a '.', it's not
	// in the standard library. This skips most GOPATH directories.
<原文结束>

# <翻译开始>
	// Quick check: if the first path component has a '.', it's not
	// in the standard library. This skips most GOPATH directories.
# <翻译结束>


<原文开始>
		// We couldn't find the gccgo search directories.
		// Best guess, since the first component did not contain
		// '.', is that this is a standard library package.
<原文结束>

# <翻译开始>
		// We couldn't find the gccgo search directories.
		// Best guess, since the first component did not contain
		// '.', is that this is a standard library package.
# <翻译结束>

