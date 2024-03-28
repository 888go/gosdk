
<原文开始>
// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
# <翻译结束>


<原文开始>
// ResetCoverage sets all of the counters for each edge of the instrumented
// source code to 0.
<原文结束>

# <翻译开始>
// ResetCoverage sets all of the counters for each edge of the instrumented
// source code to 0.
# <翻译结束>


<原文开始>
// SnapshotCoverage copies the current counter values into coverageSnapshot,
// preserving them for later inspection. SnapshotCoverage also rounds each
// counter down to the nearest power of two. This lets the coordinator store
// multiple values for each counter by OR'ing them together.
<原文结束>

# <翻译开始>
// SnapshotCoverage copies the current counter values into coverageSnapshot,
// preserving them for later inspection. SnapshotCoverage also rounds each
// counter down to the nearest power of two. This lets the coordinator store
// multiple values for each counter by OR'ing them together.
# <翻译结束>


<原文开始>
// diffCoverage returns a set of bits set in snapshot but not in base.
// If there are no new bits set, diffCoverage returns nil.
<原文结束>

# <翻译开始>
// diffCoverage returns a set of bits set in snapshot but not in base.
// If there are no new bits set, diffCoverage returns nil.
# <翻译结束>


<原文开始>
// countNewCoverageBits returns the number of bits set in snapshot that are not
// set in base.
<原文结束>

# <翻译开始>
// countNewCoverageBits returns the number of bits set in snapshot that are not
// set in base.
# <翻译结束>


<原文开始>
// isCoverageSubset returns true if all the base coverage bits are set in
// snapshot.
<原文结束>

# <翻译开始>
// isCoverageSubset returns true if all the base coverage bits are set in
// snapshot.
# <翻译结束>


<原文开始>
// hasCoverageBit returns true if snapshot has at least one bit set that is
// also set in base.
<原文结束>

# <翻译开始>
// hasCoverageBit returns true if snapshot has at least one bit set that is
// also set in base.
# <翻译结束>


<原文开始>
	// _counters and _ecounters mark the start and end, respectively, of where
	// the 8-bit coverage counters reside in memory. They're known to cmd/link,
	// which specially assigns their addresses for this purpose.
<原文结束>

# <翻译开始>
	// _counters and _ecounters mark the start and end, respectively, of where
	// the 8-bit coverage counters reside in memory. They're known to cmd/link,
	// which specially assigns their addresses for this purpose.
# <翻译结束>

