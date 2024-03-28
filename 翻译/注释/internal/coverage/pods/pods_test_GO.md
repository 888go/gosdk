
<原文开始>
// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
# <翻译结束>


<原文开始>
//testenv.MustHaveGoBuild(t)
<原文结束>

# <翻译开始>
//testenv.MustHaveGoBuild(t)
# <翻译结束>


<原文开始>
// Create a couple of directories.
<原文结束>

# <翻译开始>
// Create a couple of directories.
# <翻译结束>


<原文开始>
// Add some random files (not coverage related)
<原文结束>

# <翻译开始>
// Add some random files (not coverage related)
# <翻译结束>


<原文开始>
// Add a meta-data file with two counter files to first dir.
<原文结束>

# <翻译开始>
// Add a meta-data file with two counter files to first dir.
# <翻译结束>


<原文开始>
// Add a counter file with no associated meta file.
<原文结束>

# <翻译开始>
// Add a counter file with no associated meta file.
# <翻译结束>


<原文开始>
// Add a meta-data file with three counter files to second dir.
<原文结束>

# <翻译开始>
// Add a meta-data file with three counter files to second dir.
# <翻译结束>


<原文开始>
	// Add a duplicate of the first meta-file and a corresponding
	// counter file to the second dir. This is intended to capture
	// the scenario where we have two different runs of the same
	// coverage-instrumented binary, but with the output files
	// sent to separate directories.
<原文结束>

# <翻译开始>
	// Add a duplicate of the first meta-file and a corresponding
	// counter file to the second dir. This is intended to capture
	// the scenario where we have two different runs of the same
	// coverage-instrumented binary, but with the output files
	// sent to separate directories.
# <翻译结束>


<原文开始>
// Check handling of bad/unreadable dir.
<原文结束>

# <翻译开始>
// Check handling of bad/unreadable dir.
# <翻译结束>

