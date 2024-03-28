
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
// encVersion1 will be the first line of a file with version 1 encoding.
<原文结束>

# <翻译开始>
// encVersion1 will be the first line of a file with version 1 encoding.
# <翻译结束>


<原文开始>
// marshalCorpusFile encodes an arbitrary number of arguments into the file format for the
// corpus.
<原文结束>

# <翻译开始>
// marshalCorpusFile encodes an arbitrary number of arguments into the file format for the
// corpus.
# <翻译结束>


<原文开始>
	// TODO(katiehockman): keep uint8 and int32 encoding where applicable,
	// instead of changing to byte and rune respectively.
<原文结束>

# <翻译开始>
	// TODO(katiehockman): keep uint8 and int32 encoding where applicable,
	// instead of changing to byte and rune respectively.
# <翻译结束>


<原文开始>
				// We encode unusual NaNs as hex values, because that is how users are
				// likely to encounter them in literature about floating-point encoding.
				// This allows us to reproduce fuzz failures that depend on the specific
				// NaN representation (for float32 there are about 2^24 possibilities!),
				// not just the fact that the value is *a* NaN.
				//
				// Note that the specific value of float32(math.NaN()) can vary based on
				// whether the architecture represents signaling NaNs using a low bit
				// (as is common) or a high bit (as commonly implemented on MIPS
				// hardware before around 2012). We believe that the increase in clarity
				// from identifying "NaN" with math.NaN() is worth the slight ambiguity
				// from a platform-dependent value.
<原文结束>

# <翻译开始>
				// We encode unusual NaNs as hex values, because that is how users are
				// likely to encounter them in literature about floating-point encoding.
				// This allows us to reproduce fuzz failures that depend on the specific
				// NaN representation (for float32 there are about 2^24 possibilities!),
				// not just the fact that the value is *a* NaN.
				//
				// Note that the specific value of float32(math.NaN()) can vary based on
				// whether the architecture represents signaling NaNs using a low bit
				// (as is common) or a high bit (as commonly implemented on MIPS
				// hardware before around 2012). We believe that the increase in clarity
				// from identifying "NaN" with math.NaN() is worth the slight ambiguity
				// from a platform-dependent value.
# <翻译结束>


<原文开始>
			// Although rune and int32 are represented by the same type, only a subset
			// of valid int32 values can be expressed as rune literals. Notably,
			// negative numbers, surrogate halves, and values above unicode.MaxRune
			// have no quoted representation.
			//
			// fmt with "%q" (and the corresponding functions in the strconv package)
			// would quote out-of-range values to the Unicode replacement character
			// instead of the original value (see https://go.dev/issue/51526), so
			// they must be treated as int32 instead.
			//
			// We arbitrarily draw the line at UTF-8 validity, which biases toward the
			// "rune" interpretation. (However, we accept either format as input.)
<原文结束>

# <翻译开始>
			// Although rune and int32 are represented by the same type, only a subset
			// of valid int32 values can be expressed as rune literals. Notably,
			// negative numbers, surrogate halves, and values above unicode.MaxRune
			// have no quoted representation.
			//
			// fmt with "%q" (and the corresponding functions in the strconv package)
			// would quote out-of-range values to the Unicode replacement character
			// instead of the original value (see https://go.dev/issue/51526), so
			// they must be treated as int32 instead.
			//
			// We arbitrarily draw the line at UTF-8 validity, which biases toward the
			// "rune" interpretation. (However, we accept either format as input.)
# <翻译结束>


<原文开始>
			// For bytes, we arbitrarily prefer the character interpretation.
			// (Every byte has a valid character encoding.)
<原文结束>

# <翻译开始>
			// For bytes, we arbitrarily prefer the character interpretation.
			// (Every byte has a valid character encoding.)
# <翻译结束>


<原文开始>
// unmarshalCorpusFile decodes corpus bytes into their respective values.
<原文结束>

# <翻译开始>
// unmarshalCorpusFile decodes corpus bytes into their respective values.
# <翻译结束>


<原文开始>
// Special case for negative numbers.
<原文结束>

# <翻译开始>
// Special case for negative numbers.
# <翻译结束>


<原文开始>
// parseInt returns an integer of value val and type typ.
<原文结束>

# <翻译开始>
// parseInt returns an integer of value val and type typ.
# <翻译结束>


<原文开始>
// parseUint returns an unsigned integer of value val and type typ.
<原文结束>

# <翻译开始>
// parseUint returns an unsigned integer of value val and type typ.
# <翻译结束>

