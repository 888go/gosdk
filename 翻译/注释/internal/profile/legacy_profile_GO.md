
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
// This file implements parsers to convert legacy profiles into the
// profile.proto format.
<原文结束>

# <翻译开始>
// This file implements parsers to convert legacy profiles into the
// profile.proto format.
# <翻译结束>


<原文开始>
	// LegacyHeapAllocated instructs the heapz parsers to use the
	// allocated memory stats instead of the default in-use memory. Note
	// that tcmalloc doesn't provide all allocated memory, only in-use
	// stats.
<原文结束>

# <翻译开始>
	// LegacyHeapAllocated instructs the heapz parsers to use the
	// allocated memory stats instead of the default in-use memory. Note
	// that tcmalloc doesn't provide all allocated memory, only in-use
	// stats.
# <翻译结束>


<原文开始>
// parseGoCount parses a Go count profile (e.g., threadcreate or
// goroutine) and returns a new Profile.
<原文结束>

# <翻译开始>
// parseGoCount parses a Go count profile (e.g., threadcreate or
// goroutine) and returns a new Profile.
# <翻译结束>


<原文开始>
// Skip past comments and empty lines seeking a real header.
<原文结束>

# <翻译开始>
// Skip past comments and empty lines seeking a real header.
# <翻译结束>


<原文开始>
// Adjust all frames by -1 to land on the call instruction.
<原文结束>

# <翻译开始>
// Adjust all frames by -1 to land on the call instruction.
# <翻译结束>


<原文开始>
// remapLocationIDs ensures there is a location for each address
// referenced by a sample, and remaps the samples to point to the new
// location ids.
<原文结束>

# <翻译开始>
// remapLocationIDs ensures there is a location for each address
// referenced by a sample, and remaps the samples to point to the new
// location ids.
# <翻译结束>


<原文开始>
// remapMappingIDs matches location addresses with existing mappings
// and updates them appropriately. This is O(N*M), if this ever shows
// up as a bottleneck, evaluate sorting the mappings and doing a
// binary search, which would make it O(N*log(M)).
<原文结束>

# <翻译开始>
// remapMappingIDs matches location addresses with existing mappings
// and updates them appropriately. This is O(N*M), if this ever shows
// up as a bottleneck, evaluate sorting the mappings and doing a
// binary search, which would make it O(N*log(M)).
# <翻译结束>


<原文开始>
	// Some profile handlers will incorrectly set regions for the main
	// executable if its section is remapped. Fix them through heuristics.
<原文结束>

# <翻译开始>
	// Some profile handlers will incorrectly set regions for the main
	// executable if its section is remapped. Fix them through heuristics.
# <翻译结束>


<原文开始>
	// Remove the initial mapping if named '/anon_hugepage' and has a
	// consecutive adjacent mapping.
<原文结束>

# <翻译开始>
	// Remove the initial mapping if named '/anon_hugepage' and has a
	// consecutive adjacent mapping.
# <翻译结束>


<原文开始>
// ParseTracebacks parses a set of tracebacks and returns a newly
// populated profile. It will accept any text file and generate a
// Profile out of it with any hex addresses it can identify, including
// a process map if it can recognize one. Each sample will include a
// tag "source" with the addresses recognized in string format.
<原文结束>

# <翻译开始>
// ParseTracebacks parses a set of tracebacks and returns a newly
// populated profile. It will accept any text file and generate a
// Profile out of it with any hex addresses it can identify, including
// a process map if it can recognize one. Each sample will include a
// tag "source" with the addresses recognized in string format.
# <翻译结束>


<原文开始>
				// Addresses from stack traces point to the next instruction after
				// each call. Adjust by -1 to land somewhere on the actual call.
<原文结束>

# <翻译开始>
				// Addresses from stack traces point to the next instruction after
				// each call. Adjust by -1 to land somewhere on the actual call.
# <翻译结束>


<原文开始>
// Add final sample to save any leftover data.
<原文结束>

# <翻译开始>
// Add final sample to save any leftover data.
# <翻译结束>


<原文开始>
// parseCPU parses a profilez legacy profile and returns a newly
// populated Profile.
//
// The general format for profilez samples is a sequence of words in
// binary format. The first words are a header with the following data:
//
//	1st word -- 0
//	2nd word -- 3
//	3rd word -- 0 if a c++ application, 1 if a java application.
//	4th word -- Sampling period (in microseconds).
//	5th word -- Padding.
<原文结束>

# <翻译开始>
// parseCPU parses a profilez legacy profile and returns a newly
// populated Profile.
//
// The general format for profilez samples is a sequence of words in
// binary format. The first words are a header with the following data:
//
//	1st word -- 0
//	2nd word -- 3
//	3rd word -- 0 if a c++ application, 1 if a java application.
//	4th word -- Sampling period (in microseconds).
//	5th word -- Padding.
# <翻译结束>


<原文开始>
// cpuProfile returns a new Profile from C++ profilez data.
// b is the profile bytes after the header, period is the profiling
// period, and parse is a function to parse 8-byte chunks from the
// profile in its native endianness.
<原文结束>

# <翻译开始>
// cpuProfile returns a new Profile from C++ profilez data.
// b is the profile bytes after the header, period is the profiling
// period, and parse is a function to parse 8-byte chunks from the
// profile in its native endianness.
# <翻译结束>


<原文开始>
	// If all samples have the same second-to-the-bottom frame, it
	// strongly suggests that it is an uninteresting artifact of
	// measurement -- a stack frame pushed by the signal handler. The
	// bottom frame is always correct as it is picked up from the signal
	// structure, not the stack. Check if this is the case and if so,
	// remove.
<原文结束>

# <翻译开始>
	// If all samples have the same second-to-the-bottom frame, it
	// strongly suggests that it is an uninteresting artifact of
	// measurement -- a stack frame pushed by the signal handler. The
	// bottom frame is always correct as it is picked up from the signal
	// structure, not the stack. Check if this is the case and if so,
	// remove.
# <翻译结束>


<原文开始>
// parseCPUSamples parses a collection of profilez samples from a
// profile.
//
// profilez samples are a repeated sequence of stack frames of the
// form:
//
//	1st word -- The number of times this stack was encountered.
//	2nd word -- The size of the stack (StackSize).
//	3rd word -- The first address on the stack.
//	...
//	StackSize + 2 -- The last address on the stack
//
// The last stack trace is of the form:
//
//	1st word -- 0
//	2nd word -- 1
//	3rd word -- 0
//
// Addresses from stack traces may point to the next instruction after
// each call. Optionally adjust by -1 to land somewhere on the actual
// call (except for the leaf, which is not a call).
<原文结束>

# <翻译开始>
// parseCPUSamples parses a collection of profilez samples from a
// profile.
//
// profilez samples are a repeated sequence of stack frames of the
// form:
//
//	1st word -- The number of times this stack was encountered.
//	2nd word -- The size of the stack (StackSize).
//	3rd word -- The first address on the stack.
//	...
//	StackSize + 2 -- The last address on the stack
//
// The last stack trace is of the form:
//
//	1st word -- 0
//	2nd word -- 1
//	3rd word -- 0
//
// Addresses from stack traces may point to the next instruction after
// each call. Optionally adjust by -1 to land somewhere on the actual
// call (except for the leaf, which is not a call).
# <翻译结束>


<原文开始>
// Reached the end without finding the EOD marker.
<原文结束>

# <翻译开始>
// Reached the end without finding the EOD marker.
# <翻译结束>


<原文开始>
// parseHeap parses a heapz legacy or a growthz profile and
// returns a newly populated Profile.
<原文结束>

# <翻译开始>
// parseHeap parses a heapz legacy or a growthz profile and
// returns a newly populated Profile.
# <翻译结束>


<原文开始>
// parseHeapSample parses a single row from a heap profile into a new Sample.
<原文结束>

# <翻译开始>
// parseHeapSample parses a single row from a heap profile into a new Sample.
# <翻译结束>


<原文开始>
	// Use first two values by default; tcmalloc sampling generates the
	// same value for both, only the older heap-profile collect separate
	// stats for in-use and allocated objects.
<原文结束>

# <翻译开始>
	// Use first two values by default; tcmalloc sampling generates the
	// same value for both, only the older heap-profile collect separate
	// stats for in-use and allocated objects.
# <翻译结束>


<原文开始>
// extractHexAddresses extracts hex numbers from a string and returns
// them, together with their numeric value, in a slice.
<原文结束>

# <翻译开始>
// extractHexAddresses extracts hex numbers from a string and returns
// them, together with their numeric value, in a slice.
# <翻译结束>


<原文开始>
// Do not expect any parsing failures due to the regexp matching.
<原文结束>

# <翻译开始>
// Do not expect any parsing failures due to the regexp matching.
# <翻译结束>


<原文开始>
// parseHexAddresses parses hex numbers from a string and returns them
// in a slice.
<原文结束>

# <翻译开始>
// parseHexAddresses parses hex numbers from a string and returns them
// in a slice.
# <翻译结束>


<原文开始>
// scaleHeapSample adjusts the data from a heapz Sample to
// account for its probability of appearing in the collected
// data. heapz profiles are a sampling of the memory allocations
// requests in a program. We estimate the unsampled value by dividing
// each collected sample by its probability of appearing in the
// profile. heapz v2 profiles rely on a poisson process to determine
// which samples to collect, based on the desired average collection
// rate R. The probability of a sample of size S to appear in that
// profile is 1-exp(-S/R).
<原文结束>

# <翻译开始>
// scaleHeapSample adjusts the data from a heapz Sample to
// account for its probability of appearing in the collected
// data. heapz profiles are a sampling of the memory allocations
// requests in a program. We estimate the unsampled value by dividing
// each collected sample by its probability of appearing in the
// profile. heapz v2 profiles rely on a poisson process to determine
// which samples to collect, based on the desired average collection
// rate R. The probability of a sample of size S to appear in that
// profile is 1-exp(-S/R).
# <翻译结束>


<原文开始>
		// if rate==1 all samples were collected so no adjustment is needed.
		// if rate<1 treat as unknown and skip scaling.
<原文结束>

# <翻译开始>
		// if rate==1 all samples were collected so no adjustment is needed.
		// if rate<1 treat as unknown and skip scaling.
# <翻译结束>


<原文开始>
// parseContention parses a mutex or contention profile. There are 2 cases:
// "--- contentionz " for legacy C++ profiles (and backwards compatibility)
// "--- mutex:" or "--- contention:" for profiles generated by the Go runtime.
// This code converts the text output from runtime into a *Profile. (In the future
// the runtime might write a serialized Profile directly making this unnecessary.)
<原文结束>

# <翻译开始>
// parseContention parses a mutex or contention profile. There are 2 cases:
// "--- contentionz " for legacy C++ profiles (and backwards compatibility)
// "--- mutex:" or "--- contention:" for profiles generated by the Go runtime.
// This code converts the text output from runtime into a *Profile. (In the future
// the runtime might write a serialized Profile directly making this unnecessary.)
# <翻译结束>


<原文开始>
// parseCppContention parses the output from synchronization_profiling.cc
// for backward compatibility, and the compatible (non-debug) block profile
// output from the Go runtime.
<原文结束>

# <翻译开始>
// parseCppContention parses the output from synchronization_profiling.cc
// for backward compatibility, and the compatible (non-debug) block profile
// output from the Go runtime.
# <翻译结束>


<原文开始>
// Parse text of the form "attribute = value" before the samples.
<原文结束>

# <翻译开始>
// Parse text of the form "attribute = value" before the samples.
# <翻译结束>


<原文开始>
// CPP contentionz profiles don't have format.
<原文结束>

# <翻译开始>
// CPP contentionz profiles don't have format.
# <翻译结束>


<原文开始>
// CPP contentionz profiles don't have resolution.
<原文结束>

# <翻译开始>
// CPP contentionz profiles don't have resolution.
# <翻译结束>


<原文开始>
// parseContentionSample parses a single row from a contention profile
// into a new Sample.
<原文结束>

# <翻译开始>
// parseContentionSample parses a single row from a contention profile
// into a new Sample.
# <翻译结束>


<原文开始>
	// Unsample values if period and cpuHz are available.
	// - Delays are scaled to cycles and then to nanoseconds.
	// - Contentions are scaled to cycles.
<原文结束>

# <翻译开始>
	// Unsample values if period and cpuHz are available.
	// - Delays are scaled to cycles and then to nanoseconds.
	// - Contentions are scaled to cycles.
# <翻译结束>


<原文开始>
// parseThread parses a Threadz profile and returns a new Profile.
<原文结束>

# <翻译开始>
// parseThread parses a Threadz profile and returns a new Profile.
# <翻译结束>


<原文开始>
// Advance over initial comments until first stack trace.
<原文结束>

# <翻译开始>
// Advance over initial comments until first stack trace.
# <翻译结束>


<原文开始>
// Recognize each thread and populate profile samples.
<原文结束>

# <翻译开始>
// Recognize each thread and populate profile samples.
# <翻译结束>


<原文开始>
// We got a --same as previous threads--. Bump counters.
<原文结束>

# <翻译开始>
// We got a --same as previous threads--. Bump counters.
# <翻译结束>


<原文开始>
// parseThreadSample parses a symbolized or unsymbolized stack trace.
// Returns the first line after the traceback, the sample (or nil if
// it hits a 'same-as-previous' marker) and an error.
<原文结束>

# <翻译开始>
// parseThreadSample parses a symbolized or unsymbolized stack trace.
// Returns the first line after the traceback, the sample (or nil if
// it hits a 'same-as-previous' marker) and an error.
# <翻译结束>


<原文开始>
// parseAdditionalSections parses any additional sections in the
// profile, ignoring any unrecognized sections.
<原文结束>

# <翻译开始>
// parseAdditionalSections parses any additional sections in the
// profile, ignoring any unrecognized sections.
# <翻译结束>


<原文开始>
// Ignore any unrecognized sections.
<原文结束>

# <翻译开始>
// Ignore any unrecognized sections.
# <翻译结束>


<原文开始>
// ParseMemoryMap parses a memory map in the format of
// /proc/self/maps, and overrides the mappings in the current profile.
// It renumbers the samples and locations in the profile correspondingly.
<原文结束>

# <翻译开始>
// ParseMemoryMap parses a memory map in the format of
// /proc/self/maps, and overrides the mappings in the current profile.
// It renumbers the samples and locations in the profile correspondingly.
# <翻译结束>


<原文开始>
				// Recognize assignments of the form: attr=value, and replace
				// $attr with value on subsequent mappings.
<原文结束>

# <翻译开始>
				// Recognize assignments of the form: attr=value, and replace
				// $attr with value on subsequent mappings.
# <翻译结束>


<原文开始>
// Ignore any unrecognized entries
<原文结束>

# <翻译开始>
// Ignore any unrecognized entries
# <翻译结束>


<原文开始>
			// In some cases the first entry may include the address range
			// but not the name of the file. It should be followed by
			// another entry with the name.
<原文结束>

# <翻译开始>
			// In some cases the first entry may include the address range
			// but not the name of the file. It should be followed by
			// another entry with the name.
# <翻译结束>


<原文开始>
// Update the name if this is the entry following that empty one.
<原文结束>

# <翻译开始>
// Update the name if this is the entry following that empty one.
# <翻译结束>


<原文开始>
// Skip non-executable entries.
<原文结束>

# <翻译开始>
// Skip non-executable entries.
# <翻译结束>


<原文开始>
// early Go pprof profiles
<原文结束>

# <翻译开始>
// early Go pprof profiles
# <翻译结束>


<原文开始>
// Memory-allocation routines on OS X.
<原文结束>

# <翻译开始>
// Memory-allocation routines on OS X.
# <翻译结束>


<原文开始>
// Other misc. memory allocation routines
<原文结束>

# <翻译开始>
// Other misc. memory allocation routines
# <翻译结束>


<原文开始>
	// Preserve Go runtime frames that appear in the middle/bottom of
	// the stack.
<原文结束>

# <翻译开始>
	// Preserve Go runtime frames that appear in the middle/bottom of
	// the stack.
# <翻译结束>

