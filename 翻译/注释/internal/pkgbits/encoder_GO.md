
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
// currentVersion is the current version number.
//
//   - v0: initial prototype
//
//   - v1: adds the flags uint32 word
//
// TODO(mdempsky): For the next version bump:
//   - remove the legacy "has init" bool from the public root
//   - remove obj's "derived func instance" bool
<原文结束>

# <翻译开始>
// currentVersion is the current version number.
//
//   - v0: initial prototype
//
//   - v1: adds the flags uint32 word
//
// TODO(mdempsky): For the next version bump:
//   - remove the legacy "has init" bool from the public root
//   - remove obj's "derived func instance" bool
# <翻译结束>


<原文开始>
// A PkgEncoder provides methods for encoding a package's Unified IR
// export data.
<原文结束>

# <翻译开始>
// A PkgEncoder provides methods for encoding a package's Unified IR
// export data.
# <翻译结束>


<原文开始>
// elems holds the bitstream for previously encoded elements.
<原文结束>

# <翻译开始>
// elems holds the bitstream for previously encoded elements.
# <翻译结束>


<原文开始>
	// stringsIdx maps previously encoded strings to their index within
	// the RelocString section, to allow deduplication. That is,
	// elems[RelocString][stringsIdx[s]] == s (if present).
<原文结束>

# <翻译开始>
	// stringsIdx maps previously encoded strings to their index within
	// the RelocString section, to allow deduplication. That is,
	// elems[RelocString][stringsIdx[s]] == s (if present).
# <翻译结束>


<原文开始>
	// syncFrames is the number of frames to write at each sync
	// marker. A negative value means sync markers are omitted.
<原文结束>

# <翻译开始>
	// syncFrames is the number of frames to write at each sync
	// marker. A negative value means sync markers are omitted.
# <翻译结束>


<原文开始>
// SyncMarkers reports whether pw uses sync markers.
<原文结束>

# <翻译开始>
// SyncMarkers reports whether pw uses sync markers.
# <翻译结束>


<原文开始>
// NewPkgEncoder returns an initialized PkgEncoder.
//
// syncFrames is the number of caller frames that should be serialized
// at Sync points. Serializing additional frames results in larger
// export data files, but can help diagnosing desync errors in
// higher-level Unified IR reader/writer code. If syncFrames is
// negative, then sync markers are omitted entirely.
<原文结束>

# <翻译开始>
// NewPkgEncoder returns an initialized PkgEncoder.
//
// syncFrames is the number of caller frames that should be serialized
// at Sync points. Serializing additional frames results in larger
// export data files, but can help diagnosing desync errors in
// higher-level Unified IR reader/writer code. If syncFrames is
// negative, then sync markers are omitted entirely.
# <翻译结束>


<原文开始>
// DumpTo writes the package's encoded data to out0 and returns the
// package fingerprint.
<原文结束>

# <翻译开始>
// DumpTo writes the package's encoded data to out0 and returns the
// package fingerprint.
# <翻译结束>


<原文开始>
// StringIdx adds a string value to the strings section, if not
// already present, and returns its index.
<原文结束>

# <翻译开始>
// StringIdx adds a string value to the strings section, if not
// already present, and returns its index.
# <翻译结束>


<原文开始>
// NewEncoder returns an Encoder for a new element within the given
// section, and encodes the given SyncMarker as the start of the
// element bitstream.
<原文结束>

# <翻译开始>
// NewEncoder returns an Encoder for a new element within the given
// section, and encodes the given SyncMarker as the start of the
// element bitstream.
# <翻译结束>


<原文开始>
// NewEncoderRaw returns an Encoder for a new element within the given
// section.
//
// Most callers should use NewEncoder instead.
<原文结束>

# <翻译开始>
// NewEncoderRaw returns an Encoder for a new element within the given
// section.
//
// Most callers should use NewEncoder instead.
# <翻译结束>


<原文开始>
// An Encoder provides methods for encoding an individual element's
// bitstream data.
<原文结束>

# <翻译开始>
// An Encoder provides methods for encoding an individual element's
// bitstream data.
# <翻译结束>


<原文开始>
// accumulated element bitstream data
<原文结束>

# <翻译开始>
// accumulated element bitstream data
# <翻译结束>


<原文开始>
// index within relocation section
<原文结束>

# <翻译开始>
// index within relocation section
# <翻译结束>


<原文开始>
// Flush finalizes the element's bitstream and returns its Index.
<原文结束>

# <翻译开始>
// Flush finalizes the element's bitstream and returns its Index.
# <翻译结束>


<原文开始>
// Backup the data so we write the relocations at the front.
<原文结束>

# <翻译开始>
// Backup the data so we write the relocations at the front.
# <翻译结束>


<原文开始>
	// TODO(mdempsky): Consider writing these out separately so they're
	// easier to strip, along with function bodies, so that we can prune
	// down to just the data that's relevant to go/types.
<原文结束>

# <翻译开始>
	// TODO(mdempsky): Consider writing these out separately so they're
	// easier to strip, along with function bodies, so that we can prune
	// down to just the data that's relevant to go/types.
# <翻译结束>


<原文开始>
	// Writing out stack frame string references requires working
	// relocations, but writing out the relocations themselves involves
	// sync markers. To prevent infinite recursion, we simply trim the
	// stack frame for sync markers within the relocation header.
<原文结束>

# <翻译开始>
	// Writing out stack frame string references requires working
	// relocations, but writing out the relocations themselves involves
	// sync markers. To prevent infinite recursion, we simply trim the
	// stack frame for sync markers within the relocation header.
# <翻译结束>


<原文开始>
	// TODO(mdempsky): Save space by writing out stack frames as a
	// linked list so we can share common stack frames.
<原文结束>

# <翻译开始>
	// TODO(mdempsky): Save space by writing out stack frames as a
	// linked list so we can share common stack frames.
# <翻译结束>


<原文开始>
// Bool encodes and writes a bool value into the element bitstream,
// and then returns the bool value.
//
// For simple, 2-alternative encodings, the idiomatic way to call Bool
// is something like:
//
//	if w.Bool(x != 0) {
//		// alternative #1
//	} else {
//		// alternative #2
//	}
//
// For multi-alternative encodings, use Code instead.
<原文结束>

# <翻译开始>
// Bool encodes and writes a bool value into the element bitstream,
// and then returns the bool value.
//
// For simple, 2-alternative encodings, the idiomatic way to call Bool
// is something like:
//
//	if w.Bool(x != 0) {
//		// alternative #1
//	} else {
//		// alternative #2
//	}
//
// For multi-alternative encodings, use Code instead.
# <翻译结束>


<原文开始>
// Int64 encodes and writes an int64 value into the element bitstream.
<原文结束>

# <翻译开始>
// Int64 encodes and writes an int64 value into the element bitstream.
# <翻译结束>


<原文开始>
// Uint64 encodes and writes a uint64 value into the element bitstream.
<原文结束>

# <翻译开始>
// Uint64 encodes and writes a uint64 value into the element bitstream.
# <翻译结束>


<原文开始>
// Len encodes and writes a non-negative int value into the element bitstream.
<原文结束>

# <翻译开始>
// Len encodes and writes a non-negative int value into the element bitstream.
# <翻译结束>


<原文开始>
// Int encodes and writes an int value into the element bitstream.
<原文结束>

# <翻译开始>
// Int encodes and writes an int value into the element bitstream.
# <翻译结束>


<原文开始>
// Len encodes and writes a uint value into the element bitstream.
<原文结束>

# <翻译开始>
// Len encodes and writes a uint value into the element bitstream.
# <翻译结束>


<原文开始>
// Reloc encodes and writes a relocation for the given (section,
// index) pair into the element bitstream.
//
// Note: Only the index is formally written into the element
// bitstream, so bitstream decoders must know from context which
// section an encoded relocation refers to.
<原文结束>

# <翻译开始>
// Reloc encodes and writes a relocation for the given (section,
// index) pair into the element bitstream.
//
// Note: Only the index is formally written into the element
// bitstream, so bitstream decoders must know from context which
// section an encoded relocation refers to.
# <翻译结束>


<原文开始>
// Code encodes and writes a Code value into the element bitstream.
<原文结束>

# <翻译开始>
// Code encodes and writes a Code value into the element bitstream.
# <翻译结束>


<原文开始>
// String encodes and writes a string value into the element
// bitstream.
//
// Internally, strings are deduplicated by adding them to the strings
// section (if not already present), and then writing a relocation
// into the element bitstream.
<原文结束>

# <翻译开始>
// String encodes and writes a string value into the element
// bitstream.
//
// Internally, strings are deduplicated by adding them to the strings
// section (if not already present), and then writing a relocation
// into the element bitstream.
# <翻译结束>


<原文开始>
// StringRef writes a reference to the given index, which must be a
// previously encoded string value.
<原文结束>

# <翻译开始>
// StringRef writes a reference to the given index, which must be a
// previously encoded string value.
# <翻译结束>


<原文开始>
// Strings encodes and writes a variable-length slice of strings into
// the element bitstream.
<原文结束>

# <翻译开始>
// Strings encodes and writes a variable-length slice of strings into
// the element bitstream.
# <翻译结束>


<原文开始>
// Value encodes and writes a constant.Value into the element
// bitstream.
<原文结束>

# <翻译开始>
// Value encodes and writes a constant.Value into the element
// bitstream.
# <翻译结束>


<原文开始>
// TODO: More efficient encoding.
<原文结束>

# <翻译开始>
// TODO: More efficient encoding.
# <翻译结束>

