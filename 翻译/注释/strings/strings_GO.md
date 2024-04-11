
<原文开始>
// Count counts the number of non-overlapping instances of substr in s.
// If substr is an empty string, Count returns 1 + the number of Unicode code points in s.
<原文结束>

# <翻译开始>
// Count 计算字符串 s 中非重叠子串 substr 出现的次数。
// 若 substr 为空字符串，Count 返回 1 加上 s 中的 Unicode 编码点数量。
# <翻译结束>


<原文开始>
// Contains reports whether substr is within s.
<原文结束>

# <翻译开始>
// Contains 判断 substr 是否包含于 s 中
# <翻译结束>


<原文开始>
// ContainsAny reports whether any Unicode code points in chars are within s.
<原文结束>

# <翻译开始>
// ContainsAny 判断字符串 s 中是否包含 chars 中的任意一个Unicode字符。
# <翻译结束>


<原文开始>
// ContainsRune reports whether the Unicode code point r is within s.
<原文结束>

# <翻译开始>
// ContainsRune报告Unicode码点r是否在s中。
# <翻译结束>


<原文开始>
// ContainsFunc reports whether any Unicode code points r within s satisfy f(r).
<原文结束>

# <翻译开始>
// ContainsFunc 判断字符串 s 中是否存在任意一个 Unicode 编码点 r 满足函数 f(r) 的条件。
# <翻译结束>


<原文开始>
// LastIndex returns the index of the last instance of substr in s, or -1 if substr is not present in s.
<原文结束>

# <翻译开始>
// LastIndex 返回子串 substr 在字符串 s 中最后一次出现的索引，若 substr 不在 s 中，则返回 -1。
# <翻译结束>


<原文开始>
// IndexByte returns the index of the first instance of c in s, or -1 if c is not present in s.
<原文结束>

# <翻译开始>
// IndexByte 返回 s 中首次出现字符 c 的索引，若 c 不在 s 中，则返回 -1。
# <翻译结束>


<原文开始>
// IndexRune returns the index of the first instance of the Unicode code point
// r, or -1 if rune is not present in s.
// If r is utf8.RuneError, it returns the first instance of any
// invalid UTF-8 byte sequence.
<原文结束>

# <翻译开始>
// IndexRune 返回Unicode码点r在s中首次出现的索引，若r不在s中，则返回-1。
// 若r为utf8.RuneError，它将返回s中任意无效UTF-8字节序列的首次出现位置。
# <翻译结束>


<原文开始>
// IndexAny returns the index of the first instance of any Unicode code point
// from chars in s, or -1 if no Unicode code point from chars is present in s.
<原文结束>

# <翻译开始>
// IndexAny 返回 s 中首次出现的任意 Unicode 码点在 chars 中的索引，若 s 中不存在 chars 中的任何 Unicode 码点，则返回 -1。
# <翻译结束>


<原文开始>
// LastIndexAny returns the index of the last instance of any Unicode code
// point from chars in s, or -1 if no Unicode code point from chars is
// present in s.
<原文结束>

# <翻译开始>
// LastIndexAny返回s中任意Unicode码点字符在chars中的最后一个实例的索引，若s中不存在chars中的任何Unicode码点，则返回-1。
# <翻译结束>


<原文开始>
// LastIndexByte returns the index of the last instance of c in s, or -1 if c is not present in s.
<原文结束>

# <翻译开始>
// LastIndexByte 返回 s 中 c 最后一次出现的索引，若 c 不在 s 中，则返回 -1。
# <翻译结束>


<原文开始>
// SplitN slices s into substrings separated by sep and returns a slice of
// the substrings between those separators.
//
// The count determines the number of substrings to return:
//
//	n > 0: at most n substrings; the last substring will be the unsplit remainder.
//	n == 0: the result is nil (zero substrings)
//	n < 0: all substrings
//
// Edge cases for s and sep (for example, empty strings) are handled
// as described in the documentation for [Split].
//
// To split around the first instance of a separator, see Cut.
<原文结束>

# <翻译开始>
// SplitN 将s按sep分隔符切片，返回由这些分隔符之间的子串组成的切片。
//
// count参数决定返回的子串数量：
//
//	 n > 0：最多返回n个子串；最后一个子串为未分割的剩余部分。
//	 n == 0：结果为nil（零个子串）
//	 n < 0：返回所有子串
//
// 对于s和sep（如空字符串）的边缘情况，处理方式如[Split]文档所述。
//
// 若要按第一个分隔符实例进行分割，请参见Cut。
# <翻译结束>


<原文开始>
// SplitAfterN slices s into substrings after each instance of sep and
// returns a slice of those substrings.
//
// The count determines the number of substrings to return:
//
//	n > 0: at most n substrings; the last substring will be the unsplit remainder.
//	n == 0: the result is nil (zero substrings)
//	n < 0: all substrings
//
// Edge cases for s and sep (for example, empty strings) are handled
// as described in the documentation for SplitAfter.
<原文结束>

# <翻译开始>
// SplitAfterN 将s按sep出现的位置进行切片，返回由这些子串组成的切片。
//
// 参数count决定了返回子串的数量：
//
//	n > 0：最多返回n个子串；最后一个子串将是未分割的剩余部分。
//	n == 0：结果为nil（零个子串）
//	n < 0：返回所有子串
//
// 对于s和sep的边缘情况（如空字符串），处理方式与SplitAfter函数文档中所述一致。
# <翻译结束>


<原文开始>
// Split slices s into all substrings separated by sep and returns a slice of
// the substrings between those separators.
//
// If s does not contain sep and sep is not empty, Split returns a
// slice of length 1 whose only element is s.
//
// If sep is empty, Split splits after each UTF-8 sequence. If both s
// and sep are empty, Split returns an empty slice.
//
// It is equivalent to [SplitN] with a count of -1.
//
// To split around the first instance of a separator, see Cut.
<原文结束>

# <翻译开始>
// Split 函数将字符串 s 按照分隔符 sep 切分成所有子串，并返回由这些分隔符之间的子串组成的切片。
//
// 若 s 中不包含 sep，且 sep 非空时，Split 返回长度为 1 的切片，其中唯一元素为 s。
//
// 若 sep 为空，Split 将在每个 UTF-8 序列后进行分割。若 s 和 sep 同时为空，Split 返回一个空切片。
//
// 它等同于调用 [SplitN] 时传入计数参数 -1。
//
// 若要按照第一个分隔符实例进行分割，请参见 Cut 函数。
# <翻译结束>


<原文开始>
// SplitAfter slices s into all substrings after each instance of sep and
// returns a slice of those substrings.
//
// If s does not contain sep and sep is not empty, SplitAfter returns
// a slice of length 1 whose only element is s.
//
// If sep is empty, SplitAfter splits after each UTF-8 sequence. If
// both s and sep are empty, SplitAfter returns an empty slice.
//
// It is equivalent to [SplitAfterN] with a count of -1.
<原文结束>

# <翻译开始>
// SplitAfter 将字符串 s 按照每个出现的 sep 实例进行切分，并返回所有子串组成的切片。
//
// 若 s 中不包含 sep，且 sep 非空时，SplitAfter 返回长度为 1 的切片，其中唯一元素即为 s。
//
// 若 sep 为空，SplitAfter 将在每个 UTF-8 序列后进行切分。若 s 和 sep 同时为空，SplitAfter 返回一个空切片。
//
// 它等同于使用计数为 -1 的 [SplitAfterN]。
# <翻译结束>


<原文开始>
// Fields splits the string s around each instance of one or more consecutive white space
// characters, as defined by unicode.IsSpace, returning a slice of substrings of s or an
// empty slice if s contains only white space.
<原文结束>

# <翻译开始>
// Fields 函数将字符串 s 以其中连续出现的一个或多个空格字符（由 unicode.IsSpace 定义）为分隔，返回 s 的子字符串切片。如果 s 只包含空格，则返回一个空切片。
# <翻译结束>


<原文开始>
// FieldsFunc splits the string s at each run of Unicode code points c satisfying f(c)
// and returns an array of slices of s. If all code points in s satisfy f(c) or the
// string is empty, an empty slice is returned.
//
// FieldsFunc makes no guarantees about the order in which it calls f(c)
// and assumes that f always returns the same value for a given c.
<原文结束>

# <翻译开始>
// FieldsFunc 函数以满足 f(c) 条件的每个连续的 Unicode 编码点序列将字符串 s 分割开来，并返回 s 的切片数组。如果 s 中的所有编码点均满足 f(c) 或 s 为空，则返回一个空切片。
//
// FieldsFunc 不保证调用 f(c) 的顺序，且假定对于给定的 c，f 总是返回相同的值。
# <翻译结束>


<原文开始>
// Join concatenates the elements of its first argument to create a single string. The separator
// string sep is placed between elements in the resulting string.
<原文结束>

# <翻译开始>
// Join 函数将第一个参数中的元素连接起来，生成一个单一的字符串。分隔符字符串 sep 会被放置在结果字符串中各元素之间。
# <翻译结束>


<原文开始>
// HasPrefix reports whether the string s begins with prefix.
<原文结束>

# <翻译开始>
// HasPrefix 判断字符串 s 是否以 prefix 作为前缀。
# <翻译结束>


<原文开始>
// HasSuffix reports whether the string s ends with suffix.
<原文结束>

# <翻译开始>
// HasSuffix 判断字符串 s 是否以 suffix 结尾。
# <翻译结束>


<原文开始>
// Map returns a copy of the string s with all its characters modified
// according to the mapping function. If mapping returns a negative value, the character is
// dropped from the string with no replacement.
<原文结束>

# <翻译开始>
// Map 函数返回一个字符串 s 的副本，其中所有字符均根据映射函数进行修改。
// 若映射函数返回负值，则该字符将从字符串中删除，不作替换。
# <翻译结束>


<原文开始>
// Repeat returns a new string consisting of count copies of the string s.
//
// It panics if count is negative or if the result of (len(s) * count)
// overflows.
<原文结束>

# <翻译开始>
// Repeat 返回一个新字符串，其中包含字符串 s 的 count 份副本。
//
// 如果 count 为负数，或者 (len(s) * count) 的结果溢出时，将会引发恐慌。
# <翻译结束>


<原文开始>
// ToUpper returns s with all Unicode letters mapped to their upper case.
<原文结束>

# <翻译开始>
// ToUpper返回s，其中所有Unicode字母均映射为大写。
# <翻译结束>


<原文开始>
// ToLower returns s with all Unicode letters mapped to their lower case.
<原文结束>

# <翻译开始>
// ToLower返回s，其中所有Unicode字母已映射至其小写形式。
# <翻译结束>


<原文开始>
// ToTitle returns a copy of the string s with all Unicode letters mapped to
// their Unicode title case.
<原文结束>

# <翻译开始>
// ToTitle返回一个字符串s的副本，其中所有Unicode字母均映射至其Unicode标题大小写。
# <翻译结束>


<原文开始>
// ToUpperSpecial returns a copy of the string s with all Unicode letters mapped to their
// upper case using the case mapping specified by c.
<原文结束>

# <翻译开始>
// ToUpperSpecial返回字符串s的副本，其中所有Unicode字母均根据c指定的大小写映射规则转换为大写。
# <翻译结束>


<原文开始>
// ToLowerSpecial returns a copy of the string s with all Unicode letters mapped to their
// lower case using the case mapping specified by c.
<原文结束>

# <翻译开始>
// ToLowerSpecial返回一个字符串s的副本，其中所有Unicode字母都根据c指定的大小写映射规则转换为小写。
# <翻译结束>


<原文开始>
// ToTitleSpecial returns a copy of the string s with all Unicode letters mapped to their
// Unicode title case, giving priority to the special casing rules.
<原文结束>

# <翻译开始>
// ToTitleSpecial 返回字符串 s 的副本，其中所有 Unicode 字母都映射到其
// Unicode 标题大小写形式，并优先考虑特殊大小写规则。
# <翻译结束>


<原文开始>
// ToValidUTF8 returns a copy of the string s with each run of invalid UTF-8 byte sequences
// replaced by the replacement string, which may be empty.
<原文结束>

# <翻译开始>
// ToValidUTF8 返回字符串 s 的副本，其中每个无效的 UTF-8 字节序列均被替换为替换字符串（可能为空）。
# <翻译结束>


<原文开始>
// Title returns a copy of the string s with all Unicode letters that begin words
// mapped to their Unicode title case.
//
// Deprecated: The rule Title uses for word boundaries does not handle Unicode
// punctuation properly. Use golang.org/x/text/cases instead.
<原文结束>

# <翻译开始>
// Title返回字符串s的一个副本，其中所有作为单词开头的Unicode字母均转换为其Unicode标题大小写。
// 
// 已弃用：Title用于识别单词边界的方式无法正确处理Unicode标点符号。请改用golang.org/x/text/cases。
# <翻译结束>


<原文开始>
// TrimLeftFunc returns a slice of the string s with all leading
// Unicode code points c satisfying f(c) removed.
<原文结束>

# <翻译开始>
// TrimLeftFunc返回字符串s的一个切片，其中所有满足f(c)的起始Unicode码点c已被移除。
# <翻译结束>


<原文开始>
// TrimRightFunc returns a slice of the string s with all trailing
// Unicode code points c satisfying f(c) removed.
<原文结束>

# <翻译开始>
// TrimRightFunc返回字符串s的一个切片，其中所有满足f(c)的尾部Unicode码点均被移除。
# <翻译结束>


<原文开始>
// TrimFunc returns a slice of the string s with all leading
// and trailing Unicode code points c satisfying f(c) removed.
<原文结束>

# <翻译开始>
// TrimFunc返回字符串s的一个切片，其中所有满足f(c)的首尾Unicode字符码被移除。
# <翻译结束>


<原文开始>
// IndexFunc returns the index into s of the first Unicode
// code point satisfying f(c), or -1 if none do.
<原文结束>

# <翻译开始>
// IndexFunc返回s中第一个满足f(c)的Unicode码点的索引，如果不存在满足条件的码点，则返回-1。
# <翻译结束>


<原文开始>
// LastIndexFunc returns the index into s of the last
// Unicode code point satisfying f(c), or -1 if none do.
<原文结束>

# <翻译开始>
// LastIndexFunc返回s中最后一个满足f(c)的Unicode码点的索引，若无满足条件的码点，则返回-1。
# <翻译结束>


<原文开始>
// Trim returns a slice of the string s with all leading and
// trailing Unicode code points contained in cutset removed.
<原文结束>

# <翻译开始>
// Trim返回字符串s的一个切片，其中包含的所有在cutset中的前导和尾部Unicode字符码点已被移除。
# <翻译结束>


<原文开始>
// TrimLeft returns a slice of the string s with all leading
// Unicode code points contained in cutset removed.
//
// To remove a prefix, use [TrimPrefix] instead.
<原文结束>

# <翻译开始>
// TrimLeft返回字符串s的一个切片，其中所有位于开头的
// Unicode编码点（如果存在于cutset中）均被移除。
//
// 要移除前缀，请使用[TrimPrefix]代替。
# <翻译结束>


<原文开始>
// TrimRight returns a slice of the string s, with all trailing
// Unicode code points contained in cutset removed.
//
// To remove a suffix, use [TrimSuffix] instead.
<原文结束>

# <翻译开始>
// TrimRight返回字符串s的一个切片，其中所有尾部包含在cutset中的Unicode字符点已被移除。
//
// 要移除后缀，应使用[TrimSuffix]代替。
# <翻译结束>


<原文开始>
// TrimSpace returns a slice of the string s, with all leading
// and trailing white space removed, as defined by Unicode.
<原文结束>

# <翻译开始>
// TrimSpace返回字符串s的一个切片，其中已移除了所有Unicode定义的开头和结尾空白字符。
# <翻译结束>


<原文开始>
// TrimPrefix returns s without the provided leading prefix string.
// If s doesn't start with prefix, s is returned unchanged.
<原文结束>

# <翻译开始>
// TrimPrefix返回字符串s，其中已移除了提供的前缀字符串。
// 若s以prefix开头，则返回去掉prefix后的s；否则返回s保持原样不变。
# <翻译结束>


<原文开始>
// TrimSuffix returns s without the provided trailing suffix string.
// If s doesn't end with suffix, s is returned unchanged.
<原文结束>

# <翻译开始>
// TrimSuffix 函数返回去掉指定尾部后缀字符串后的 s。
// 若 s 末尾不包含该后缀，则返回 s 保持原样不变。
# <翻译结束>


<原文开始>
// Replace returns a copy of the string s with the first n
// non-overlapping instances of old replaced by new.
// If old is empty, it matches at the beginning of the string
// and after each UTF-8 sequence, yielding up to k+1 replacements
// for a k-rune string.
// If n < 0, there is no limit on the number of replacements.
<原文结束>

# <翻译开始>
// Replace返回一个字符串s的副本，其中前n个非重叠的old子串被new替换。
// 若old为空，它将在字符串起始处以及每个UTF-8序列之后匹配，对于一个包含k个rune字符的字符串，最多可进行k+1次替换。
// 若n < 0，则替换次数无限制。
# <翻译结束>


<原文开始>
// ReplaceAll returns a copy of the string s with all
// non-overlapping instances of old replaced by new.
// If old is empty, it matches at the beginning of the string
// and after each UTF-8 sequence, yielding up to k+1 replacements
// for a k-rune string.
<原文结束>

# <翻译开始>
// ReplaceAll返回字符串s的副本，其中所有非重叠的old实例均被new替换。
// 如果old为空，它将在字符串起始处以及每个UTF-8序列之后匹配，对于一个包含k个符文的字符串，最多产生k+1次替换。
# <翻译结束>


<原文开始>
// EqualFold reports whether s and t, interpreted as UTF-8 strings,
// are equal under simple Unicode case-folding, which is a more general
// form of case-insensitivity.
<原文结束>

# <翻译开始>
// EqualFold报告当s和t以UTF-8字符串形式解释时，它们在简单Unicode大小写折叠规则下是否相等。这是一种更通用的大小写不敏感比较方式。
# <翻译结束>


<原文开始>
// Index returns the index of the first instance of substr in s, or -1 if substr is not present in s.
<原文结束>

# <翻译开始>
// Index 返回字符串 s 中首次出现子串 substr 的索引，若 substr 不在 s 中，则返回 -1。
# <翻译结束>


<原文开始>
// Cut slices s around the first instance of sep,
// returning the text before and after sep.
// The found result reports whether sep appears in s.
// If sep does not appear in s, cut returns s, "", false.
<原文结束>

# <翻译开始>
// Cut 函数在字符串切片 s 中首次出现 sep 位置处进行切割，
// 返回 sep 之前及之后的文本。
// found 结果用于指示 sep 是否在 s 中出现。
// 若 sep 未在 s 中出现，则 Cut 返回 s，空字符串，false。
# <翻译结束>


<原文开始>
// CutPrefix returns s without the provided leading prefix string
// and reports whether it found the prefix.
// If s doesn't start with prefix, CutPrefix returns s, false.
// If prefix is the empty string, CutPrefix returns s, true.
<原文结束>

# <翻译开始>
// CutPrefix 函数返回去掉指定前缀字符串后的 s，并报告是否找到了该前缀。
// 若 s 未以 prefix 开头，CutPrefix 返回 s, false。
// 若 prefix 为空字符串，CutPrefix 返回 s, true。
# <翻译结束>


<原文开始>
// CutSuffix returns s without the provided ending suffix string
// and reports whether it found the suffix.
// If s doesn't end with suffix, CutSuffix returns s, false.
// If suffix is the empty string, CutSuffix returns s, true.
<原文结束>

# <翻译开始>
// CutSuffix 函数返回去掉指定结尾后缀字符串的 s，并报告是否找到了该后缀。
// 若 s 未以 suffix 结尾，CutSuffix 返回 s, false。
// 若 suffix 为空字符串，CutSuffix 返回 s, true。
# <翻译结束>

