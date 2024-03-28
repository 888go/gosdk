
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
			// The two IEEE 754 bit patterns used for the math.Float{64,32}frombits
			// encodings are non-math.NAN quiet-NaN values. Since they are not equal
			// to math.NaN(), they should be re-encoded to their bit patterns. They
			// are, respectively:
			//   * math.Float64bits(math.NaN())+1
			//   * math.Float32bits(float32(math.NaN()))+1
<原文结束>

# <翻译开始>
			// The two IEEE 754 bit patterns used for the math.Float{64,32}frombits
			// encodings are non-math.NAN quiet-NaN values. Since they are not equal
			// to math.NaN(), they should be re-encoded to their bit patterns. They
			// are, respectively:
			//   * math.Float64bits(math.NaN())+1
			//   * math.Float32bits(float32(math.NaN()))+1
# <翻译结束>


<原文开始>
			// Although we arbitrarily choose default integer bases (0 or 16), we may
			// want to change those arbitrary choices in the future and should not
			// break the parser. Verify that integers in the opposite bases still
			// parse correctly.
<原文结束>

# <翻译开始>
			// Although we arbitrarily choose default integer bases (0 or 16), we may
			// want to change those arbitrary choices in the future and should not
			// break the parser. Verify that integers in the opposite bases still
			// parse correctly.
# <翻译结束>


<原文开始>
// BenchmarkMarshalCorpusFile measures the time it takes to serialize byte
// slices of various sizes to a corpus file. The slice contains a repeating
// sequence of bytes 0-255 to mix escaped and non-escaped characters.
<原文结束>

# <翻译开始>
// BenchmarkMarshalCorpusFile measures the time it takes to serialize byte
// slices of various sizes to a corpus file. The slice contains a repeating
// sequence of bytes 0-255 to mix escaped and non-escaped characters.
# <翻译结束>


<原文开始>
// BenchmarkUnmarshalCorpusfile measures the time it takes to deserialize
// files encoding byte slices of various sizes. The slice contains a repeating
// sequence of bytes 0-255 to mix escaped and non-escaped characters.
<原文结束>

# <翻译开始>
// BenchmarkUnmarshalCorpusfile measures the time it takes to deserialize
// files encoding byte slices of various sizes. The slice contains a repeating
// sequence of bytes 0-255 to mix escaped and non-escaped characters.
# <翻译结束>

