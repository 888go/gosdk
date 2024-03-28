
<原文开始>
// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
# <翻译结束>


<原文开始>
// Implements methods to filter samples from profiles.
<原文结束>

# <翻译开始>
// Implements methods to filter samples from profiles.
# <翻译结束>


<原文开始>
// FilterSamplesByName filters the samples in a profile and only keeps
// samples where at least one frame matches focus but none match ignore.
// Returns true is the corresponding regexp matched at least one sample.
<原文结束>

# <翻译开始>
// FilterSamplesByName filters the samples in a profile and only keeps
// samples where at least one frame matches focus but none match ignore.
// Returns true is the corresponding regexp matched at least one sample.
# <翻译结束>


<原文开始>
// Remove sample with no locations (by not adding it to s).
<原文结束>

# <翻译开始>
// Remove sample with no locations (by not adding it to s).
# <翻译结束>


<原文开始>
// matchesName reports whether the function name or file in the
// location matches the regular expression.
<原文结束>

# <翻译开始>
// matchesName reports whether the function name or file in the
// location matches the regular expression.
# <翻译结束>


<原文开始>
// unmatchedLines returns the lines in the location that do not match
// the regular expression.
<原文结束>

# <翻译开始>
// unmatchedLines returns the lines in the location that do not match
// the regular expression.
# <翻译结束>


<原文开始>
// focusedAndNotIgnored looks up a slice of ids against a map of
// focused/ignored locations. The map only contains locations that are
// explicitly focused or ignored. Returns whether there is at least
// one focused location but no ignored locations.
<原文结束>

# <翻译开始>
// focusedAndNotIgnored looks up a slice of ids against a map of
// focused/ignored locations. The map only contains locations that are
// explicitly focused or ignored. Returns whether there is at least
// one focused location but no ignored locations.
# <翻译结束>


<原文开始>
				// Found focused location. Must keep searching in case there
				// is an ignored one as well.
<原文结束>

# <翻译开始>
				// Found focused location. Must keep searching in case there
				// is an ignored one as well.
# <翻译结束>


<原文开始>
// Found ignored location. Can return false right away.
<原文结束>

# <翻译开始>
// Found ignored location. Can return false right away.
# <翻译结束>


<原文开始>
// TagMatch selects tags for filtering
<原文结束>

# <翻译开始>
// TagMatch selects tags for filtering
# <翻译结束>


<原文开始>
// FilterSamplesByTag removes all samples from the profile, except
// those that match focus and do not match the ignore regular
// expression.
<原文结束>

# <翻译开始>
// FilterSamplesByTag removes all samples from the profile, except
// those that match focus and do not match the ignore regular
// expression.
# <翻译结束>


<原文开始>
// focusedSample checks a sample against focus and ignore regexps.
// Returns whether the focus/ignore regexps match any tags.
<原文结束>

# <翻译开始>
// focusedSample checks a sample against focus and ignore regexps.
// Returns whether the focus/ignore regexps match any tags.
# <翻译结束>

