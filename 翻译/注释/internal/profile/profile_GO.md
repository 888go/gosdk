
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
// Package profile provides a representation of
// github.com/google/pprof/proto/profile.proto and
// methods to encode/decode/merge profiles in this format.
<原文结束>

# <翻译开始>
// Package profile provides a representation of
// github.com/google/pprof/proto/profile.proto and
// methods to encode/decode/merge profiles in this format.
# <翻译结束>


<原文开始>
// Profile is an in-memory representation of profile.proto.
<原文结束>

# <翻译开始>
// Profile is an in-memory representation of profile.proto.
# <翻译结束>


<原文开始>
// ValueType corresponds to Profile.ValueType
<原文结束>

# <翻译开始>
// ValueType corresponds to Profile.ValueType
# <翻译结束>


<原文开始>
// cpu, wall, inuse_space, etc
<原文结束>

# <翻译开始>
// cpu, wall, inuse_space, etc
# <翻译结束>


<原文开始>
// seconds, nanoseconds, bytes, etc
<原文结束>

# <翻译开始>
// seconds, nanoseconds, bytes, etc
# <翻译结束>


<原文开始>
// Sample corresponds to Profile.Sample
<原文结束>

# <翻译开始>
// Sample corresponds to Profile.Sample
# <翻译结束>


<原文开始>
// Label corresponds to Profile.Label
<原文结束>

# <翻译开始>
// Label corresponds to Profile.Label
# <翻译结束>


<原文开始>
// Exactly one of the two following values must be set
<原文结束>

# <翻译开始>
// Exactly one of the two following values must be set
# <翻译结束>


<原文开始>
// Integer value for this label
<原文结束>

# <翻译开始>
// Integer value for this label
# <翻译结束>


<原文开始>
// Mapping corresponds to Profile.Mapping
<原文结束>

# <翻译开始>
// Mapping corresponds to Profile.Mapping
# <翻译结束>


<原文开始>
// Location corresponds to Profile.Location
<原文结束>

# <翻译开始>
// Location corresponds to Profile.Location
# <翻译结束>


<原文开始>
// Function corresponds to Profile.Function
<原文结束>

# <翻译开始>
// Function corresponds to Profile.Function
# <翻译结束>


<原文开始>
// Parse parses a profile and checks for its validity. The input
// may be a gzip-compressed encoded protobuf or one of many legacy
// profile formats which may be unsupported in the future.
<原文结束>

# <翻译开始>
// Parse parses a profile and checks for its validity. The input
// may be a gzip-compressed encoded protobuf or one of many legacy
// profile formats which may be unsupported in the future.
# <翻译结束>


<原文开始>
// goroutine, threadcreate
<原文结束>

# <翻译开始>
// goroutine, threadcreate
# <翻译结束>


<原文开始>
// setMain scans Mapping entries and guesses which entry is main
// because legacy profiles don't obey the convention of putting main
// first.
<原文结束>

# <翻译开始>
// setMain scans Mapping entries and guesses which entry is main
// because legacy profiles don't obey the convention of putting main
// first.
# <翻译结束>


<原文开始>
// Swap what we guess is main to position 0.
<原文结束>

# <翻译开始>
// Swap what we guess is main to position 0.
# <翻译结束>


<原文开始>
// Write writes the profile as a gzip-compressed marshaled protobuf.
<原文结束>

# <翻译开始>
// Write writes the profile as a gzip-compressed marshaled protobuf.
# <翻译结束>


<原文开始>
// CheckValid tests whether the profile is valid. Checks include, but are
// not limited to:
//   - len(Profile.Sample[n].value) == len(Profile.value_unit)
//   - Sample.id has a corresponding Profile.Location
<原文结束>

# <翻译开始>
// CheckValid tests whether the profile is valid. Checks include, but are
// not limited to:
//   - len(Profile.Sample[n].value) == len(Profile.value_unit)
//   - Sample.id has a corresponding Profile.Location
# <翻译结束>


<原文开始>
// Check that sample values are consistent
<原文结束>

# <翻译开始>
// Check that sample values are consistent
# <翻译结束>


<原文开始>
	// Check that all mappings/locations/functions are in the tables
	// Check that there are no duplicate ids
<原文结束>

# <翻译开始>
	// Check that all mappings/locations/functions are in the tables
	// Check that there are no duplicate ids
# <翻译结束>


<原文开始>
// Aggregate merges the locations in the profile into equivalence
// classes preserving the request attributes. It also updates the
// samples to point to the merged locations.
<原文结束>

# <翻译开始>
// Aggregate merges the locations in the profile into equivalence
// classes preserving the request attributes. It also updates the
// samples to point to the merged locations.
# <翻译结束>


<原文开始>
// Print dumps a text representation of a profile. Intended mainly
// for debugging purposes.
<原文结束>

# <翻译开始>
// Print dumps a text representation of a profile. Intended mainly
// for debugging purposes.
# <翻译结束>


<原文开始>
// Do not print location details past the first line
<原文结束>

# <翻译开始>
// Do not print location details past the first line
# <翻译结束>


<原文开始>
// Merge adds profile p adjusted by ratio r into profile p. Profiles
// must be compatible (same Type and SampleType).
// TODO(rsilvera): consider normalizing the profiles based on the
// total samples collected.
<原文结束>

# <翻译开始>
// Merge adds profile p adjusted by ratio r into profile p. Profiles
// must be compatible (same Type and SampleType).
// TODO(rsilvera): consider normalizing the profiles based on the
// total samples collected.
# <翻译结束>


<原文开始>
// Keep the largest of the two periods.
<原文结束>

# <翻译开始>
// Keep the largest of the two periods.
# <翻译结束>


<原文开始>
// Compatible determines if two profiles can be compared/merged.
// returns nil if the profiles are compatible; otherwise an error with
// details on the incompatibility.
<原文结束>

# <翻译开始>
// Compatible determines if two profiles can be compared/merged.
// returns nil if the profiles are compatible; otherwise an error with
// details on the incompatibility.
# <翻译结束>


<原文开始>
// HasFunctions determines if all locations in this profile have
// symbolized function information.
<原文结束>

# <翻译开始>
// HasFunctions determines if all locations in this profile have
// symbolized function information.
# <翻译结束>


<原文开始>
// HasFileLines determines if all locations in this profile have
// symbolized file and line number information.
<原文结束>

# <翻译开始>
// HasFileLines determines if all locations in this profile have
// symbolized file and line number information.
# <翻译结束>


<原文开始>
// No grounds to disqualify.
<原文结束>

# <翻译开始>
// No grounds to disqualify.
# <翻译结束>


<原文开始>
// Copy makes a fully independent copy of a profile.
<原文结束>

# <翻译开始>
// Copy makes a fully independent copy of a profile.
# <翻译结束>


<原文开始>
// Demangler maps symbol names to a human-readable form. This may
// include C++ demangling and additional simplification. Names that
// are not demangled may be missing from the resulting map.
<原文结束>

# <翻译开始>
// Demangler maps symbol names to a human-readable form. This may
// include C++ demangling and additional simplification. Names that
// are not demangled may be missing from the resulting map.
# <翻译结束>


<原文开始>
// Demangle attempts to demangle and optionally simplify any function
// names referenced in the profile. It works on a best-effort basis:
// it will silently preserve the original names in case of any errors.
<原文结束>

# <翻译开始>
// Demangle attempts to demangle and optionally simplify any function
// names referenced in the profile. It works on a best-effort basis:
// it will silently preserve the original names in case of any errors.
# <翻译结束>


<原文开始>
// Collect names to demangle.
<原文结束>

# <翻译开始>
// Collect names to demangle.
# <翻译结束>


<原文开始>
// Update profile with demangled names.
<原文结束>

# <翻译开始>
// Update profile with demangled names.
# <翻译结束>


<原文开始>
// Empty reports whether the profile contains no samples.
<原文结束>

# <翻译开始>
// Empty reports whether the profile contains no samples.
# <翻译结束>


<原文开始>
// Scale multiplies all sample values in a profile by a constant.
<原文结束>

# <翻译开始>
// Scale multiplies all sample values in a profile by a constant.
# <翻译结束>


<原文开始>
// ScaleN multiplies each sample values in a sample by a different amount.
<原文结束>

# <翻译开始>
// ScaleN multiplies each sample values in a sample by a different amount.
# <翻译结束>

