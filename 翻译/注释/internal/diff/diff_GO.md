
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
// A pair is a pair of values tracked for both the x and y side of a diff.
// It is typically a pair of line indexes.
<原文结束>

# <翻译开始>
// A pair is a pair of values tracked for both the x and y side of a diff.
// It is typically a pair of line indexes.
# <翻译结束>


<原文开始>
	// Loop over matches to consider,
	// expanding each match to include surrounding lines,
	// and then printing diff chunks.
	// To avoid setup/teardown cases outside the loop,
	// tgs returns a leading {0,0} and trailing {len(x), len(y)} pair
	// in the sequence of matches.
<原文结束>

# <翻译开始>
	// Loop over matches to consider,
	// expanding each match to include surrounding lines,
	// and then printing diff chunks.
	// To avoid setup/teardown cases outside the loop,
	// tgs returns a leading {0,0} and trailing {len(x), len(y)} pair
	// in the sequence of matches.
# <翻译结束>


<原文开始>
// printed up to x[:done.x] and y[:done.y]
<原文结束>

# <翻译开始>
// printed up to x[:done.x] and y[:done.y]
# <翻译结束>


<原文开始>
// start lines of current chunk
<原文结束>

# <翻译开始>
// start lines of current chunk
# <翻译结束>


<原文开始>
// number of lines from each side in current chunk
<原文结束>

# <翻译开始>
// number of lines from each side in current chunk
# <翻译结束>


<原文开始>
// Already handled scanning forward from earlier match.
<原文结束>

# <翻译开始>
// Already handled scanning forward from earlier match.
# <翻译结束>


<原文开始>
		// Expand matching lines as far possible,
		// establishing that x[start.x:end.x] == y[start.y:end.y].
		// Note that on the first (or last) iteration we may (or definitey do)
		// have an empty match: start.x==end.x and start.y==end.y.
<原文结束>

# <翻译开始>
		// Expand matching lines as far possible,
		// establishing that x[start.x:end.x] == y[start.y:end.y].
		// Note that on the first (or last) iteration we may (or definitey do)
		// have an empty match: start.x==end.x and start.y==end.y.
# <翻译结束>


<原文开始>
		// Emit the mismatched lines before start into this chunk.
		// (No effect on first sentinel iteration, when start = {0,0}.)
<原文结束>

# <翻译开始>
		// Emit the mismatched lines before start into this chunk.
		// (No effect on first sentinel iteration, when start = {0,0}.)
# <翻译结束>


<原文开始>
		// If we're not at EOF and have too few common lines,
		// the chunk includes all the common lines and continues.
<原文结束>

# <翻译开始>
		// If we're not at EOF and have too few common lines,
		// the chunk includes all the common lines and continues.
# <翻译结束>


<原文开始>
// number of context lines
<原文结束>

# <翻译开始>
// number of context lines
# <翻译结束>


<原文开始>
// End chunk with common lines for context.
<原文结束>

# <翻译开始>
// End chunk with common lines for context.
# <翻译结束>


<原文开始>
			// Format and emit chunk.
			// Convert line numbers to 1-indexed.
			// Special case: empty file shows up as 0,0 not 1,0.
<原文结束>

# <翻译开始>
			// Format and emit chunk.
			// Convert line numbers to 1-indexed.
			// Special case: empty file shows up as 0,0 not 1,0.
# <翻译结束>


<原文开始>
// If we reached EOF, we're done.
<原文结束>

# <翻译开始>
// If we reached EOF, we're done.
# <翻译结束>


<原文开始>
// Otherwise start a new chunk.
<原文结束>

# <翻译开始>
// Otherwise start a new chunk.
# <翻译结束>


<原文开始>
		// Treat last line as having a message about the missing newline attached,
		// using the same text as BSD/GNU diff (including the leading backslash).
<原文结束>

# <翻译开始>
		// Treat last line as having a message about the missing newline attached,
		// using the same text as BSD/GNU diff (including the leading backslash).
# <翻译结束>


<原文开始>
	// Count the number of times each string appears in a and b.
	// We only care about 0, 1, many, counted as 0, -1, -2
	// for the x side and 0, -4, -8 for the y side.
	// Using negative numbers now lets us distinguish positive line numbers later.
<原文结束>

# <翻译开始>
	// Count the number of times each string appears in a and b.
	// We only care about 0, 1, many, counted as 0, -1, -2
	// for the x side and 0, -4, -8 for the y side.
	// Using negative numbers now lets us distinguish positive line numbers later.
# <翻译结束>


<原文开始>
	// Now unique strings can be identified by m[s] = -1+-4.
	//
	// Gather the indexes of those strings in x and y, building:
	//	xi[i] = increasing indexes of unique strings in x.
	//	yi[i] = increasing indexes of unique strings in y.
	//	inv[i] = index j such that x[xi[i]] = y[yi[j]].
<原文结束>

# <翻译开始>
	// Now unique strings can be identified by m[s] = -1+-4.
	//
	// Gather the indexes of those strings in x and y, building:
	//	xi[i] = increasing indexes of unique strings in x.
	//	yi[i] = increasing indexes of unique strings in y.
	//	inv[i] = index j such that x[xi[i]] = y[yi[j]].
# <翻译结束>


<原文开始>
	// Apply Algorithm A from Szymanski's paper.
	// In those terms, A = J = inv and B = [0, n).
	// We add sentinel pairs {0,0}, and {len(x),len(y)}
	// to the returned sequence, to help the processing loop.
<原文结束>

# <翻译开始>
	// Apply Algorithm A from Szymanski's paper.
	// In those terms, A = J = inv and B = [0, n).
	// We add sentinel pairs {0,0}, and {len(x),len(y)}
	// to the returned sequence, to help the processing loop.
# <翻译结束>

