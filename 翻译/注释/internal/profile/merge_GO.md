
<原文开始>
// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
# <翻译结束>


<原文开始>
// Merge merges all the profiles in profs into a single Profile.
// Returns a new profile independent of the input profiles. The merged
// profile is compacted to eliminate unused samples, locations,
// functions and mappings. Profiles must have identical profile sample
// and period types or the merge will fail. profile.Period of the
// resulting profile will be the maximum of all profiles, and
// profile.TimeNanos will be the earliest nonzero one.
<原文结束>

# <翻译开始>
// Merge merges all the profiles in profs into a single Profile.
// Returns a new profile independent of the input profiles. The merged
// profile is compacted to eliminate unused samples, locations,
// functions and mappings. Profiles must have identical profile sample
// and period types or the merge will fail. profile.Period of the
// resulting profile will be the maximum of all profiles, and
// profile.TimeNanos will be the earliest nonzero one.
# <翻译结束>


<原文开始>
// Clear the profile-specific hash tables
<原文结束>

# <翻译开始>
// Clear the profile-specific hash tables
# <翻译结束>


<原文开始>
			// The Mapping list has the property that the first mapping
			// represents the main binary. Take the first Mapping we see,
			// otherwise the operations below will add mappings in an
			// arbitrary order.
<原文结束>

# <翻译开始>
			// The Mapping list has the property that the first mapping
			// represents the main binary. Take the first Mapping we see,
			// otherwise the operations below will add mappings in an
			// arbitrary order.
# <翻译结束>


<原文开始>
			// If there are any zero samples, re-merge the profile to GC
			// them.
<原文结束>

# <翻译开始>
			// If there are any zero samples, re-merge the profile to GC
			// them.
# <翻译结束>


<原文开始>
// Normalize normalizes the source profile by multiplying each value in profile by the
// ratio of the sum of the base profile's values of that sample type to the sum of the
// source profile's value of that sample type.
<原文结束>

# <翻译开始>
// Normalize normalizes the source profile by multiplying each value in profile by the
// ratio of the sum of the base profile's values of that sample type to the sum of the
// source profile's value of that sample type.
# <翻译结束>


<原文开始>
// Memoization tables within a profile.
<原文结束>

# <翻译开始>
// Memoization tables within a profile.
# <翻译结束>


<原文开始>
// Memoization tables for profile entities.
<原文结束>

# <翻译开始>
// Memoization tables for profile entities.
# <翻译结束>


<原文开始>
	// Check memoization table. Must be done on the remapped location to
	// account for the remapped mapping. Add current values to the
	// existing sample.
<原文结束>

# <翻译开始>
	// Check memoization table. Must be done on the remapped location to
	// account for the remapped mapping. Add current values to the
	// existing sample.
# <翻译结束>


<原文开始>
// key generates sampleKey to be used as a key for maps.
<原文结束>

# <翻译开始>
// key generates sampleKey to be used as a key for maps.
# <翻译结束>


<原文开始>
	// Check memoization table. Must be done on the remapped location to
	// account for the remapped mapping ID.
<原文结束>

# <翻译开始>
	// Check memoization table. Must be done on the remapped location to
	// account for the remapped mapping ID.
# <翻译结束>


<原文开始>
// key generates locationKey to be used as a key for maps.
<原文结束>

# <翻译开始>
// key generates locationKey to be used as a key for maps.
# <翻译结束>


<原文开始>
// Normalizes address to handle address space randomization.
<原文结束>

# <翻译开始>
// Normalizes address to handle address space randomization.
# <翻译结束>


<原文开始>
// Check memoization tables.
<原文结束>

# <翻译开始>
// Check memoization tables.
# <翻译结束>


<原文开始>
// Update memoization tables.
<原文结束>

# <翻译开始>
// Update memoization tables.
# <翻译结束>


<原文开始>
// key generates encoded strings of Mapping to be used as a key for
// maps.
<原文结束>

# <翻译开始>
// key generates encoded strings of Mapping to be used as a key for
// maps.
# <翻译结束>


<原文开始>
	// Normalize addresses to handle address space randomization.
	// Round up to next 4K boundary to avoid minor discrepancies.
<原文结束>

# <翻译开始>
	// Normalize addresses to handle address space randomization.
	// Round up to next 4K boundary to avoid minor discrepancies.
# <翻译结束>


<原文开始>
		// A mapping containing neither build ID nor file name is a fake mapping. A
		// key with empty buildIDOrFile is used for fake mappings so that they are
		// treated as the same mapping during merging.
<原文结束>

# <翻译开始>
		// A mapping containing neither build ID nor file name is a fake mapping. A
		// key with empty buildIDOrFile is used for fake mappings so that they are
		// treated as the same mapping during merging.
# <翻译结束>


<原文开始>
// key generates a struct to be used as a key for maps.
<原文结束>

# <翻译开始>
// key generates a struct to be used as a key for maps.
# <翻译结束>


<原文开始>
// combineHeaders checks that all profiles can be merged and returns
// their combined profile.
<原文结束>

# <翻译开始>
// combineHeaders checks that all profiles can be merged and returns
// their combined profile.
# <翻译结束>


<原文开始>
// compatible determines if two profiles can be compared/merged.
// returns nil if the profiles are compatible; otherwise an error with
// details on the incompatibility.
<原文结束>

# <翻译开始>
// compatible determines if two profiles can be compared/merged.
// returns nil if the profiles are compatible; otherwise an error with
// details on the incompatibility.
# <翻译结束>


<原文开始>
// equalValueType returns true if the two value types are semantically
// equal. It ignores the internal fields used during encode/decode.
<原文结束>

# <翻译开始>
// equalValueType returns true if the two value types are semantically
// equal. It ignores the internal fields used during encode/decode.
# <翻译结束>

