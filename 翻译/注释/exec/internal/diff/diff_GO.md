
<原文开始>
// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// 版权所有 ? 2022 Go作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。
# <翻译结束>


<原文开始>
// A pair is a pair of values tracked for both the x and y side of a diff.
// It is typically a pair of line indexes.
<原文结束>

# <翻译开始>
// 一对是diff中同时跟踪x和y两侧的两个值。通常是一对行索引。
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
// 遍历待考虑的匹配项，
// 将每个匹配项扩展至包含周围的行，
// 然后打印差异块。
// 为避免循环外的设置/清理情况，
// tgs 在匹配项序列中返回一个前导的 {0,0} 和尾部的 {len(x), len(y)} 对。
# <翻译结束>


<原文开始>
// printed up to x[:done.x] and y[:done.y]
<原文结束>

# <翻译开始>
// 已打印至 x[:done.x] 和 y[:done.y]
# <翻译结束>


<原文开始>
// start lines of current chunk
<原文结束>

# <翻译开始>
// 当前块的起始行
# <翻译结束>


<原文开始>
// number of lines from each side in current chunk
<原文结束>

# <翻译开始>
// 当前块中每侧的行数
# <翻译结束>


<原文开始>
// Already handled scanning forward from earlier match.
<原文结束>

# <翻译开始>
// 已经处理了从先前匹配处向前扫描的情况
# <翻译结束>


<原文开始>
		// Expand matching lines as far as possible,
		// establishing that x[start.x:end.x] == y[start.y:end.y].
		// Note that on the first (or last) iteration we may (or definitely do)
		// have an empty match: start.x==end.x and start.y==end.y.
<原文结束>

# <翻译开始>
// 尽可能地扩展匹配行，
// 确保 x[start.x:end.x] 与 y[start.y:end.y] 相等。
// 注意在第一次（或最后一次）迭代中，我们可能（或肯定）会遇到空匹配：start.x==end.x 且 start.y==end.y。
# <翻译结束>


<原文开始>
		// Emit the mismatched lines before start into this chunk.
		// (No effect on first sentinel iteration, when start = {0,0}.)
<原文结束>

# <翻译开始>
// 将起始点之前的不匹配行输出到此块中。
// （在第一个哨兵迭代时无影响，此时 start = {0,0}。）
# <翻译结束>


<原文开始>
		// If we're not at EOF and have too few common lines,
		// the chunk includes all the common lines and continues.
<原文结束>

# <翻译开始>
// 如果我们没有到达EOF（文件结束符）且共同行数过少，
// 则该块应包含所有共同行并继续。
# <翻译结束>


<原文开始>
// number of context lines
<原文结束>

# <翻译开始>
// 上下文行数
# <翻译结束>


<原文开始>
// End chunk with common lines for context.
<原文结束>

# <翻译开始>
// 以包含上下文的通用行结束数据块
# <翻译结束>


<原文开始>
			// Format and emit chunk.
			// Convert line numbers to 1-indexed.
			// Special case: empty file shows up as 0,0 not 1,0.
<原文结束>

# <翻译开始>
// 格式化并输出代码块。
// 将行号转换为1索引制。
// 特殊情况：空文件显示为0,0而非1,0。
# <翻译结束>


<原文开始>
// If we reached EOF, we're done.
<原文结束>

# <翻译开始>
// 如果我们到达了EOF，就完成了。
# <翻译结束>


<原文开始>
// Otherwise start a new chunk.
<原文结束>

# <翻译开始>
// 否则开始一个新的分块
# <翻译结束>


<原文开始>
		// Treat last line as having a message about the missing newline attached,
		// using the same text as BSD/GNU diff (including the leading backslash).
<原文结束>

# <翻译开始>
// 将最后一行视为附带有关缺失换行符的消息，
// 使用与 BSD/GNU diff 相同的文本（包括前导反斜杠）。
# <翻译结束>


<原文开始>
	// Count the number of times each string appears in a and b.
	// We only care about 0, 1, many, counted as 0, -1, -2
	// for the x side and 0, -4, -8 for the y side.
	// Using negative numbers now lets us distinguish positive line numbers later.
<原文结束>

# <翻译开始>
// 统计字符串a和b中各字符串出现的次数。
// 我们仅关注每个字符串在a、b中出现0次、1次或多于1次的情况，分别计为0、-1、-2（对于x侧）以及0、-4、-8（对于y侧）。
// 现在使用负数，以便后续区分正向行号。
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
// 现在，可以通过 m[s] = -1+-4 来识别唯一的字符串。
// 
// 收集 x 和 y 中这些字符串的索引，构建如下：
//  xi[i] = x 中唯一字符串递增的索引。
//  yi[i] = y 中唯一字符串递增的索引。
// inv[i] = 索引 j，满足 x[xi[i]] = y[yi[j]]。
# <翻译结束>


<原文开始>
	// Apply Algorithm A from Szymanski's paper.
	// In those terms, A = J = inv and B = [0, n).
	// We add sentinel pairs {0,0}, and {len(x),len(y)}
	// to the returned sequence, to help the processing loop.
<原文结束>

# <翻译开始>
// 应用Szymanski论文中的算法A。
// 用该文中的术语表示，A = J = inv，B = [0, n)。
// 我们在返回的序列中添加哨兵对{0,0}和{len(x),len(y)}，以辅助处理循环。
# <翻译结束>

