
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
// A PkgDecoder provides methods for decoding a package's Unified IR
// export data.
<原文结束>

# <翻译开始>
// A PkgDecoder provides methods for decoding a package's Unified IR
// export data.
# <翻译结束>


<原文开始>
// version is the file format version.
<原文结束>

# <翻译开始>
// version is the file format version.
# <翻译结束>


<原文开始>
// sync indicates whether the file uses sync markers.
<原文结束>

# <翻译开始>
// sync indicates whether the file uses sync markers.
# <翻译结束>


<原文开始>
	// pkgPath is the package path for the package to be decoded.
	//
	// TODO(mdempsky): Remove; unneeded since CL 391014.
<原文结束>

# <翻译开始>
	// pkgPath is the package path for the package to be decoded.
	//
	// TODO(mdempsky): Remove; unneeded since CL 391014.
# <翻译结束>


<原文开始>
	// elemData is the full data payload of the encoded package.
	// Elements are densely and contiguously packed together.
	//
	// The last 8 bytes of elemData are the package fingerprint.
<原文结束>

# <翻译开始>
	// elemData is the full data payload of the encoded package.
	// Elements are densely and contiguously packed together.
	//
	// The last 8 bytes of elemData are the package fingerprint.
# <翻译结束>


<原文开始>
	// elemEnds stores the byte-offset end positions of element
	// bitstreams within elemData.
	//
	// For example, element I's bitstream data starts at elemEnds[I-1]
	// (or 0, if I==0) and ends at elemEnds[I].
	//
	// Note: elemEnds is indexed by absolute indices, not
	// section-relative indices.
<原文结束>

# <翻译开始>
	// elemEnds stores the byte-offset end positions of element
	// bitstreams within elemData.
	//
	// For example, element I's bitstream data starts at elemEnds[I-1]
	// (or 0, if I==0) and ends at elemEnds[I].
	//
	// Note: elemEnds is indexed by absolute indices, not
	// section-relative indices.
# <翻译结束>


<原文开始>
	// elemEndsEnds stores the index-offset end positions of relocation
	// sections within elemEnds.
	//
	// For example, section K's end positions start at elemEndsEnds[K-1]
	// (or 0, if K==0) and end at elemEndsEnds[K].
<原文结束>

# <翻译开始>
	// elemEndsEnds stores the index-offset end positions of relocation
	// sections within elemEnds.
	//
	// For example, section K's end positions start at elemEndsEnds[K-1]
	// (or 0, if K==0) and end at elemEndsEnds[K].
# <翻译结束>


<原文开始>
// PkgPath returns the package path for the package
//
// TODO(mdempsky): Remove; unneeded since CL 391014.
<原文结束>

# <翻译开始>
// PkgPath returns the package path for the package
//
// TODO(mdempsky): Remove; unneeded since CL 391014.
# <翻译结束>


<原文开始>
// SyncMarkers reports whether pr uses sync markers.
<原文结束>

# <翻译开始>
// SyncMarkers reports whether pr uses sync markers.
# <翻译结束>


<原文开始>
// NewPkgDecoder returns a PkgDecoder initialized to read the Unified
// IR export data from input. pkgPath is the package path for the
// compilation unit that produced the export data.
//
// TODO(mdempsky): Remove pkgPath parameter; unneeded since CL 391014.
<原文结束>

# <翻译开始>
// NewPkgDecoder returns a PkgDecoder initialized to read the Unified
// IR export data from input. pkgPath is the package path for the
// compilation unit that produced the export data.
//
// TODO(mdempsky): Remove pkgPath parameter; unneeded since CL 391014.
# <翻译结束>


<原文开始>
	// TODO(mdempsky): Implement direct indexing of input string to
	// avoid copying the position information.
<原文结束>

# <翻译开始>
	// TODO(mdempsky): Implement direct indexing of input string to
	// avoid copying the position information.
# <翻译结束>


<原文开始>
// NumElems returns the number of elements in section k.
<原文结束>

# <翻译开始>
// NumElems returns the number of elements in section k.
# <翻译结束>


<原文开始>
// TotalElems returns the total number of elements across all sections.
<原文结束>

# <翻译开始>
// TotalElems returns the total number of elements across all sections.
# <翻译结束>


<原文开始>
// Fingerprint returns the package fingerprint.
<原文结束>

# <翻译开始>
// Fingerprint returns the package fingerprint.
# <翻译结束>


<原文开始>
// AbsIdx returns the absolute index for the given (section, index)
// pair.
<原文结束>

# <翻译开始>
// AbsIdx returns the absolute index for the given (section, index)
// pair.
# <翻译结束>


<原文开始>
// DataIdx returns the raw element bitstream for the given (section,
// index) pair.
<原文结束>

# <翻译开始>
// DataIdx returns the raw element bitstream for the given (section,
// index) pair.
# <翻译结束>


<原文开始>
// StringIdx returns the string value for the given string index.
<原文结束>

# <翻译开始>
// StringIdx returns the string value for the given string index.
# <翻译结束>


<原文开始>
// NewDecoder returns a Decoder for the given (section, index) pair,
// and decodes the given SyncMarker from the element bitstream.
<原文结束>

# <翻译开始>
// NewDecoder returns a Decoder for the given (section, index) pair,
// and decodes the given SyncMarker from the element bitstream.
# <翻译结束>


<原文开始>
// TempDecoder returns a Decoder for the given (section, index) pair,
// and decodes the given SyncMarker from the element bitstream.
// If possible the Decoder should be RetireDecoder'd when it is no longer
// needed, this will avoid heap allocations.
<原文结束>

# <翻译开始>
// TempDecoder returns a Decoder for the given (section, index) pair,
// and decodes the given SyncMarker from the element bitstream.
// If possible the Decoder should be RetireDecoder'd when it is no longer
// needed, this will avoid heap allocations.
# <翻译结束>


<原文开始>
// NewDecoderRaw returns a Decoder for the given (section, index) pair.
//
// Most callers should use NewDecoder instead.
<原文结束>

# <翻译开始>
// NewDecoderRaw returns a Decoder for the given (section, index) pair.
//
// Most callers should use NewDecoder instead.
# <翻译结束>


<原文开始>
// A Decoder provides methods for decoding an individual element's
// bitstream data.
<原文结束>

# <翻译开始>
// A Decoder provides methods for decoding an individual element's
// bitstream data.
# <翻译结束>


<原文开始>
// readUvarint is a type-specialized copy of encoding/binary.ReadUvarint.
// This avoids the interface conversion and thus has better escape properties,
// which flows up the stack.
<原文结束>

# <翻译开始>
// readUvarint is a type-specialized copy of encoding/binary.ReadUvarint.
// This avoids the interface conversion and thus has better escape properties,
// which flows up the stack.
# <翻译结束>


<原文开始>
// Sync decodes a sync marker from the element bitstream and asserts
// that it matches the expected marker.
//
// If EnableSync is false, then Sync is a no-op.
<原文结束>

# <翻译开始>
// Sync decodes a sync marker from the element bitstream and asserts
// that it matches the expected marker.
//
// If EnableSync is false, then Sync is a no-op.
# <翻译结束>


<原文开始>
	// There's some tension here between printing:
	//
	// (1) full file paths that tools can recognize (e.g., so emacs
	//     hyperlinks the "file:line" text for easy navigation), or
	//
	// (2) short file paths that are easier for humans to read (e.g., by
	//     omitting redundant or irrelevant details, so it's easier to
	//     focus on the useful bits that remain).
	//
	// The current formatting favors the former, as it seems more
	// helpful in practice. But perhaps the formatting could be improved
	// to better address both concerns. For example, use relative file
	// paths if they would be shorter, or rewrite file paths to contain
	// "$GOROOT" (like objabi.AbsFile does) if tools can be taught how
	// to reliably expand that again.
<原文结束>

# <翻译开始>
	// There's some tension here between printing:
	//
	// (1) full file paths that tools can recognize (e.g., so emacs
	//     hyperlinks the "file:line" text for easy navigation), or
	//
	// (2) short file paths that are easier for humans to read (e.g., by
	//     omitting redundant or irrelevant details, so it's easier to
	//     focus on the useful bits that remain).
	//
	// The current formatting favors the former, as it seems more
	// helpful in practice. But perhaps the formatting could be improved
	// to better address both concerns. For example, use relative file
	// paths if they would be shorter, or rewrite file paths to contain
	// "$GOROOT" (like objabi.AbsFile does) if tools can be taught how
	// to reliably expand that again.
# <翻译结束>


<原文开始>
// TODO(mdempsky): Dynamically size?
<原文结束>

# <翻译开始>
// TODO(mdempsky): Dynamically size?
# <翻译结束>


<原文开始>
	// We already printed a stack trace for the reader, so now we can
	// simply exit. Printing a second one with panic or base.Fatalf
	// would just be noise.
<原文结束>

# <翻译开始>
	// We already printed a stack trace for the reader, so now we can
	// simply exit. Printing a second one with panic or base.Fatalf
	// would just be noise.
# <翻译结束>


<原文开始>
// Bool decodes and returns a bool value from the element bitstream.
<原文结束>

# <翻译开始>
// Bool decodes and returns a bool value from the element bitstream.
# <翻译结束>


<原文开始>
// Int64 decodes and returns an int64 value from the element bitstream.
<原文结束>

# <翻译开始>
// Int64 decodes and returns an int64 value from the element bitstream.
# <翻译结束>


<原文开始>
// Int64 decodes and returns a uint64 value from the element bitstream.
<原文结束>

# <翻译开始>
// Int64 decodes and returns a uint64 value from the element bitstream.
# <翻译结束>


<原文开始>
// Len decodes and returns a non-negative int value from the element bitstream.
<原文结束>

# <翻译开始>
// Len decodes and returns a non-negative int value from the element bitstream.
# <翻译结束>


<原文开始>
// Int decodes and returns an int value from the element bitstream.
<原文结束>

# <翻译开始>
// Int decodes and returns an int value from the element bitstream.
# <翻译结束>


<原文开始>
// Uint decodes and returns a uint value from the element bitstream.
<原文结束>

# <翻译开始>
// Uint decodes and returns a uint value from the element bitstream.
# <翻译结束>


<原文开始>
// Code decodes a Code value from the element bitstream and returns
// its ordinal value. It's the caller's responsibility to convert the
// result to an appropriate Code type.
//
// TODO(mdempsky): Ideally this method would have signature "Code[T
// Code] T" instead, but we don't allow generic methods and the
// compiler can't depend on generics yet anyway.
<原文结束>

# <翻译开始>
// Code decodes a Code value from the element bitstream and returns
// its ordinal value. It's the caller's responsibility to convert the
// result to an appropriate Code type.
//
// TODO(mdempsky): Ideally this method would have signature "Code[T
// Code] T" instead, but we don't allow generic methods and the
// compiler can't depend on generics yet anyway.
# <翻译结束>


<原文开始>
// Reloc decodes a relocation of expected section k from the element
// bitstream and returns an index to the referenced element.
<原文结束>

# <翻译开始>
// Reloc decodes a relocation of expected section k from the element
// bitstream and returns an index to the referenced element.
# <翻译结束>


<原文开始>
// String decodes and returns a string value from the element
// bitstream.
<原文结束>

# <翻译开始>
// String decodes and returns a string value from the element
// bitstream.
# <翻译结束>


<原文开始>
// Strings decodes and returns a variable-length slice of strings from
// the element bitstream.
<原文结束>

# <翻译开始>
// Strings decodes and returns a variable-length slice of strings from
// the element bitstream.
# <翻译结束>


<原文开始>
// Value decodes and returns a constant.Value from the element
// bitstream.
<原文结束>

# <翻译开始>
// Value decodes and returns a constant.Value from the element
// bitstream.
# <翻译结束>


<原文开始>
// TODO(mdempsky): These should probably be removed. I think they're a
// smell that the export data format is not yet quite right.
<原文结束>

# <翻译开始>
// TODO(mdempsky): These should probably be removed. I think they're a
// smell that the export data format is not yet quite right.
# <翻译结束>


<原文开始>
// PeekPkgPath returns the package path for the specified package
// index.
<原文结束>

# <翻译开始>
// PeekPkgPath returns the package path for the specified package
// index.
# <翻译结束>


<原文开始>
// PeekObj returns the package path, object name, and CodeObj for the
// specified object index.
<原文结束>

# <翻译开始>
// PeekObj returns the package path, object name, and CodeObj for the
// specified object index.
# <翻译结束>

