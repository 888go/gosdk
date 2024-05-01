
<原文开始>
// Equal reports whether a and b
// are the same length and contain the same bytes.
// A nil argument is equivalent to an empty slice.
<原文结束>

# <翻译开始>
// Equal 判断 a 和 b 是否具有相同的长度并且包含相同的字节。一个空指针（nil）等同于一个空切片（empty slice）。
// md5:dbece9ef3f693465
# <翻译结束>


<原文开始>
// Compare returns an integer comparing two byte slices lexicographically.
// The result will be 0 if a == b, -1 if a < b, and +1 if a > b.
// A nil argument is equivalent to an empty slice.
<原文结束>

# <翻译开始>
// Compare 函数按字典顺序比较两个字节切片。如果 a == b，返回 0；如果 a < b，返回 -1；如果 a > b，返回 +1。
// 空指针（nil）等同于空切片。
// md5:23a14b305b5a0036
# <翻译结束>


<原文开始>
// Count counts the number of non-overlapping instances of sep in s.
// If sep is an empty slice, Count returns 1 + the number of UTF-8-encoded code points in s.
<原文结束>

# <翻译开始>
// Count 计算字符串 s 中非重叠实例 sep 的数量。
// 如果 sep 是一个空切片，Count 返回 1 加上 s 中 UTF-8 编码的代码点数量。
// md5:5d59a1b45cedfba7
# <翻译结束>


<原文开始>
// Contains reports whether subslice is within b.
<原文结束>

# <翻译开始>
// Contains 函数报告子切片是否在 b 中。. md5:3075713b6bb2e961
# <翻译结束>


<原文开始>
// ContainsAny reports whether any of the UTF-8-encoded code points in chars are within b.
<原文结束>

# <翻译开始>
// ContainsAny报告chars中的任何UTF-8编码的代码点是否都在b中。. md5:a7123747027d0a57
# <翻译结束>


<原文开始>
// ContainsRune reports whether the rune is contained in the UTF-8-encoded byte slice b.
<原文结束>

# <翻译开始>
// ContainsRune 检查 rune 是否存在于 UTF-8 编码的字节切片 b 中。. md5:22a878ae15a5922f
# <翻译结束>


<原文开始>
// ContainsFunc reports whether any of the UTF-8-encoded code points r within b satisfy f(r).
<原文结束>

# <翻译开始>
// ContainsFunc 报告UTF-8编码的码点r在b中是否存在，满足条件f(r)。. md5:e58efa31a9ff987a
# <翻译结束>


<原文开始>
// IndexByte returns the index of the first instance of c in b, or -1 if c is not present in b.
<原文结束>

# <翻译开始>
// IndexByte 返回字节切片b中字符c的第一个出现位置，如果c不在b中，则返回-1。. md5:7bf26378aad4cd0b
# <翻译结束>


<原文开始>
// LastIndex returns the index of the last instance of sep in s, or -1 if sep is not present in s.
<原文结束>

# <翻译开始>
// LastIndex 返回 s 中 sep 的最后一个实例的索引，如果 s 中不存在 sep，则返回 -1。. md5:6bd49252012e79ad
# <翻译结束>


<原文开始>
// LastIndexByte returns the index of the last instance of c in s, or -1 if c is not present in s.
<原文结束>

# <翻译开始>
// LastIndexByte 返回字符串s中字符c的最后一个出现位置的索引，如果c不在s中，则返回-1。. md5:8a4f0f6af8efe072
# <翻译结束>


<原文开始>
// IndexRune interprets s as a sequence of UTF-8-encoded code points.
// It returns the byte index of the first occurrence in s of the given rune.
// It returns -1 if rune is not present in s.
// If r is utf8.RuneError, it returns the first instance of any
// invalid UTF-8 byte sequence.
<原文结束>

# <翻译开始>
// IndexRune 将 s 解释为一系列 UTF-8 编码的代码点。
// 它返回给定字符在 s 中首次出现的字节索引。
// 如果字符在 s 中不存在，则返回 -1。
// 如果 r 是 utf8.RuneError，它将返回任何无效 UTF-8 字节序列的第一个实例。
// md5:6dc9763770576138
# <翻译结束>


<原文开始>
// IndexAny interprets s as a sequence of UTF-8-encoded Unicode code points.
// It returns the byte index of the first occurrence in s of any of the Unicode
// code points in chars. It returns -1 if chars is empty or if there is no code
// point in common.
<原文结束>

# <翻译开始>
// IndexAny将s视为一个用UTF-8编码的Unicode码点序列。
// 它返回s中第一个出现的、在chars中任一Unicode码点的字节索引。如果chars为空或者没有共同的码点，它返回-1。
// md5:03920b48034c033b
# <翻译结束>


<原文开始>
// LastIndexAny interprets s as a sequence of UTF-8-encoded Unicode code
// points. It returns the byte index of the last occurrence in s of any of
// the Unicode code points in chars. It returns -1 if chars is empty or if
// there is no code point in common.
<原文结束>

# <翻译开始>
// LastIndexAny将s解释为UTF-8编码的Unicode代码点序列。它返回s中与chars中的任何Unicode代码点的最后一个出现的字节索引。如果chars为空或没有公共代码点，它将返回-1。
// md5:7c8b4adaf5973148
# <翻译结束>


<原文开始>
// SplitN slices s into subslices separated by sep and returns a slice of
// the subslices between those separators.
// If sep is empty, SplitN splits after each UTF-8 sequence.
// The count determines the number of subslices to return:
//
//	n > 0: at most n subslices; the last subslice will be the unsplit remainder.
//	n == 0: the result is nil (zero subslices)
//	n < 0: all subslices
//
// To split around the first instance of a separator, see Cut.
<原文结束>

# <翻译开始>
// SplitN 将字符串 s 按照分隔符 sep 进行切片，并返回这些分隔符之间的子切片的切片。
// 如果 sep 为空，SplitN 将在每个 UTF-8 序列后进行切分。
// count 决定返回的子切片数量：
// 
//	> 0：最多返回 n 个子切片；最后一个子切片将是未被分割的部分。
//	n == 0：结果为 nil（没有子切片）
//	n < 0：返回所有子切片
// 
// 要在第一个分隔符周围进行切分，请参阅 Cut 函数。
// md5:0be3f8772a05e662
# <翻译结束>


<原文开始>
// SplitAfterN slices s into subslices after each instance of sep and
// returns a slice of those subslices.
// If sep is empty, SplitAfterN splits after each UTF-8 sequence.
// The count determines the number of subslices to return:
//
//	n > 0: at most n subslices; the last subslice will be the unsplit remainder.
//	n == 0: the result is nil (zero subslices)
//	n < 0: all subslices
<原文结束>

# <翻译开始>
// SplitAfterN 将字符串 s 根据 sep 的每个实例进行切分，之后返回这些子切片的切片。
// 如果 sep 为空，SplitAfterN 会根据每个 UTF-8 序列进行切分。
// count 参数决定了要返回的子切片数量：
//
//	n > 0: 最多 n 个子切片；最后一个子切片将是未被切分的剩余部分。
//	n == 0: 结果为 nil（零个子切片）
//	n < 0: 所有子切片
// md5:93042132420aa229
# <翻译结束>


<原文开始>
// Split slices s into all subslices separated by sep and returns a slice of
// the subslices between those separators.
// If sep is empty, Split splits after each UTF-8 sequence.
// It is equivalent to SplitN with a count of -1.
//
// To split around the first instance of a separator, see Cut.
<原文结束>

# <翻译开始>
// Split 将切片 s 按照分隔符 sep 切分成多个子切片，并返回一个包含这些子切片的切片。
// 如果 sep 为空，Split 将在每个 UTF-8 字符序列后进行分割。
// 它等价于 SplitN 函数，传入的计数参数为 -1。
//
// 若要围绕第一个分隔符进行分割，可以使用 Cut 函数。
// md5:d5d8f14cae52de3f
# <翻译结束>


<原文开始>
// SplitAfter slices s into all subslices after each instance of sep and
// returns a slice of those subslices.
// If sep is empty, SplitAfter splits after each UTF-8 sequence.
// It is equivalent to SplitAfterN with a count of -1.
<原文结束>

# <翻译开始>
// SplitAfter 在每个sep实例之后将s分割成所有子切片，并返回这些子切片的切片。
// 如果sep为空，SplitAfter会在每个UTF-8序列后进行分割。
// 它等同于SplitAfterN，其中计数为-1。
// md5:c78b01ae43df412f
# <翻译结束>


<原文开始>
// Fields interprets s as a sequence of UTF-8-encoded code points.
// It splits the slice s around each instance of one or more consecutive white space
// characters, as defined by unicode.IsSpace, returning a slice of subslices of s or an
// empty slice if s contains only white space.
<原文结束>

# <翻译开始>
// Fields 将 s 解析为一系列 UTF-8 编码的代码点。
// 它会在 s 中每个连续或非连续的空白字符（由 unicode.IsSpace 定义）周围分割，返回 s 的子切片的切片，如果 s 只包含空白字符，则返回空切片。
// md5:7c06fb47034bd649
# <翻译结束>


<原文开始>
// FieldsFunc interprets s as a sequence of UTF-8-encoded code points.
// It splits the slice s at each run of code points c satisfying f(c) and
// returns a slice of subslices of s. If all code points in s satisfy f(c), or
// len(s) == 0, an empty slice is returned.
//
// FieldsFunc makes no guarantees about the order in which it calls f(c)
// and assumes that f always returns the same value for a given c.
<原文结束>

# <翻译开始>
// FieldsFunc 将 s 解释为一系列 UTF-8 编码的代码点。
// 它会在满足条件 f(c) 的每个代码点序列处将切片 s 分割，
// 并返回 s 的子切片切片。如果 s 中的所有代码点都满足 f(c)，或者
// len(s) == 0，则返回一个空切片。
//
// FieldsFunc 不保证以特定顺序调用 f(c)，
// 并假设对于给定的 c，f 总是返回相同的结果值。
// md5:e0df308c2de6e89c
# <翻译结束>


<原文开始>
// Join concatenates the elements of s to create a new byte slice. The separator
// sep is placed between elements in the resulting slice.
<原文结束>

# <翻译开始>
// Join 将 s 中的元素连接起来创建一个新的字节切片。分隔符 sep 会被插入到结果切片中的元素之间。
// md5:c65602b4a3a28272
# <翻译结束>


<原文开始>
// HasPrefix reports whether the byte slice s begins with prefix.
<原文结束>

# <翻译开始>
// HasPrefix 函数报告切片 s 是否以 prefix 开头。. md5:aeb71c9e082f4e79
# <翻译结束>


<原文开始>
// HasSuffix reports whether the byte slice s ends with suffix.
<原文结束>

# <翻译开始>
// HasSuffix 函数报告 byte slice s 是否以 suffix 结尾。. md5:42621549774f0616
# <翻译结束>


<原文开始>
// Map returns a copy of the byte slice s with all its characters modified
// according to the mapping function. If mapping returns a negative value, the character is
// dropped from the byte slice with no replacement. The characters in s and the
// output are interpreted as UTF-8-encoded code points.
<原文结束>

# <翻译开始>
// Map 函数返回一个副本，该副本中的字节切片 s 的所有字符都根据映射函数进行了修改。
// 如果映射函数返回一个负值，则该字符将从字节切片中移除且不做替换。
// 字节切片 s 中的字符以及输出结果都按 UTF-8 编码的码点进行解释。
// md5:f29d9994e9e93cbe
# <翻译结束>


<原文开始>
// Repeat returns a new byte slice consisting of count copies of b.
//
// It panics if count is negative or if the result of (len(b) * count)
// overflows.
<原文结束>

# <翻译开始>
// Repeat 返回一个新的字节切片，由 count 个 b 的副本组成。
//
// 如果 count 为负数，或者 (len(b) * count) 的结果溢出，则该函数会 panic。
// md5:a30126d911ce9e3a
# <翻译结束>


<原文开始>
// ToUpper returns a copy of the byte slice s with all Unicode letters mapped to
// their upper case.
<原文结束>

# <翻译开始>
// ToUpper返回一个副本，其中s中的所有Unicode字母映射到大写。
// md5:701556ffdae09171
# <翻译结束>


<原文开始>
// ToLower returns a copy of the byte slice s with all Unicode letters mapped to
// their lower case.
<原文结束>

# <翻译开始>
// ToLower返回一个副本，其中s中的所有Unicode字母映射到小写。
// md5:96198f3611614a06
# <翻译结束>


<原文开始>
// ToTitle treats s as UTF-8-encoded bytes and returns a copy with all the Unicode letters mapped to their title case.
<原文结束>

# <翻译开始>
// ToTitle将s视为UTF-8编码的字节，并返回一个副本，其中所有Unicode字母都被映射为它们的标题大小写形式。. md5:3a15836e7e4e4976
# <翻译结束>


<原文开始>
// ToUpperSpecial treats s as UTF-8-encoded bytes and returns a copy with all the Unicode letters mapped to their
// upper case, giving priority to the special casing rules.
<原文结束>

# <翻译开始>
// ToUpperSpecial将s视为UTF-8编码的字节，并返回一个副本，其中所有Unicode字母都映射为其大写形式，优先考虑特殊大小写规则。
// md5:7a8fc8bbf32666bf
# <翻译结束>


<原文开始>
// ToLowerSpecial treats s as UTF-8-encoded bytes and returns a copy with all the Unicode letters mapped to their
// lower case, giving priority to the special casing rules.
<原文结束>

# <翻译开始>
// ToLowerSpecial将s视为UTF-8编码的字节，并返回一个副本，其中所有Unicode字母都被映射到小写，优先考虑特殊情况规则。
// md5:1c197945b5358611
# <翻译结束>


<原文开始>
// ToTitleSpecial treats s as UTF-8-encoded bytes and returns a copy with all the Unicode letters mapped to their
// title case, giving priority to the special casing rules.
<原文结束>

# <翻译开始>
// ToTitleSpecial 将 s 视为 UTF-8 编码的字节，并返回一个副本，其中所有 Unicode 字母都被映射为其大写形式，优先遵循特殊的大写规则。
// md5:05fe1b145cbd5bda
# <翻译结束>


<原文开始>
// ToValidUTF8 treats s as UTF-8-encoded bytes and returns a copy with each run of bytes
// representing invalid UTF-8 replaced with the bytes in replacement, which may be empty.
<原文结束>

# <翻译开始>
// ToValidUTF8 将s视为UTF-8编码的字节序列，并返回一个副本，其中每个代表无效UTF-8的字节序列都被replacement中的字节替换，replacement可以是空字符串。
// md5:a8577838395945fd
# <翻译结束>


<原文开始>
// Title treats s as UTF-8-encoded bytes and returns a copy with all Unicode letters that begin
// words mapped to their title case.
//
// Deprecated: The rule Title uses for word boundaries does not handle Unicode
// punctuation properly. Use golang.org/x/text/cases instead.
<原文结束>

# <翻译开始>
// Title 将 s 视为 UTF-8 编码的字节，并返回一个副本，其中所有单词起始的Unicode字母都被转换为标题大小写。
//
// 已废弃：Title 使用的单词边界规则无法正确处理Unicode标点符号。请改用 golang.org/x/text/cases。
// md5:489de93c97e62d72
# <翻译结束>


<原文开始>
// TrimLeftFunc treats s as UTF-8-encoded bytes and returns a subslice of s by slicing off
// all leading UTF-8-encoded code points c that satisfy f(c).
<原文结束>

# <翻译开始>
// TrimLeftFunc将s视为UTF-8编码的字节，并通过移除所有满足条件f(c)的UTF-8编码的代码点c，返回s的子切片。
// md5:21e9929d142ec743
# <翻译结束>


<原文开始>
// TrimRightFunc returns a subslice of s by slicing off all trailing
// UTF-8-encoded code points c that satisfy f(c).
<原文结束>

# <翻译开始>
// TrimRightFunc通过移除所有满足条件f(c)的UTF-8编码的码点c，返回s的子切片。
// md5:3016ff9a3bf64754
# <翻译结束>


<原文开始>
// TrimFunc returns a subslice of s by slicing off all leading and trailing
// UTF-8-encoded code points c that satisfy f(c).
<原文结束>

# <翻译开始>
// TrimFunc 返回s的子切片，通过去除所有满足f(c)的前导和尾随UTF-8编码的字符c。
// md5:a1e6af674d552e9a
# <翻译结束>


<原文开始>
// TrimPrefix returns s without the provided leading prefix string.
// If s doesn't start with prefix, s is returned unchanged.
<原文结束>

# <翻译开始>
// TrimPrefix 函数返回字符串 s 去除指定前缀后的结果。
// 如果 s 不以 prefix 开头，那么将原样返回 s。
// md5:97cb4f309ba2947f
# <翻译结束>


<原文开始>
// TrimSuffix returns s without the provided trailing suffix string.
// If s doesn't end with suffix, s is returned unchanged.
<原文结束>

# <翻译开始>
// TrimSuffix 函数返回 s，但不包含提供的尾随后缀字符串。
// 如果 s 不以 suffix 结尾，则返回原始的 s。
// md5:6d9d00533816bde2
# <翻译结束>


<原文开始>
// IndexFunc interprets s as a sequence of UTF-8-encoded code points.
// It returns the byte index in s of the first Unicode
// code point satisfying f(c), or -1 if none do.
<原文结束>

# <翻译开始>
// IndexFunc 将 s 解释为一系列 UTF-8 编码的代码点。
// 它返回满足函数 f(c) 的第一个 Unicode 代码点在 s 中的字节索引，如果没有满足条件的，则返回 -1。
// md5:3faa26ff439023a4
# <翻译结束>


<原文开始>
// LastIndexFunc interprets s as a sequence of UTF-8-encoded code points.
// It returns the byte index in s of the last Unicode
// code point satisfying f(c), or -1 if none do.
<原文结束>

# <翻译开始>
// LastIndexFunc 将 s 解释为一系列 UTF-8 编码的代码点。
// 它返回 s 中满足 f(c) 的最后一个 Unicode 代码点的字节索引，
// 如果没有满足条件的，则返回 -1。
// md5:d29cc3d1192d35e5
# <翻译结束>


<原文开始>
// Trim returns a subslice of s by slicing off all leading and
// trailing UTF-8-encoded code points contained in cutset.
<原文结束>

# <翻译开始>
// Trim返回一个s的子切片，切掉cutset中包含的所有领先和尾随的UTF-8编码的字符。
// md5:c8162b6bb6472995
# <翻译结束>


<原文开始>
// TrimLeft returns a subslice of s by slicing off all leading
// UTF-8-encoded code points contained in cutset.
<原文结束>

# <翻译开始>
// TrimLeft通过移除s中所有UTF-8编码的代码点开始部分，返回一个子切片。
// md5:0908d5b6be56ad64
# <翻译结束>


<原文开始>
// TrimRight returns a subslice of s by slicing off all trailing
// UTF-8-encoded code points that are contained in cutset.
<原文结束>

# <翻译开始>
// TrimRight 通过裁剪掉所有由 cutset 所包含的 UTF-8 编码字符点，返回 s 的子切片。
// md5:f6b6ebb338bed118
# <翻译结束>


<原文开始>
// TrimSpace returns a subslice of s by slicing off all leading and
// trailing white space, as defined by Unicode.
<原文结束>

# <翻译开始>
// TrimSpace 返回 s 的子切片，通过去除所有前导和尾随的空白字符，这些空白字符根据 Unicode 定义。
// md5:bd70164451a18ab6
# <翻译结束>


<原文开始>
// Runes interprets s as a sequence of UTF-8-encoded code points.
// It returns a slice of runes (Unicode code points) equivalent to s.
<原文结束>

# <翻译开始>
// Runes 将 s 解释为一个使用UTF-8编码的字符序列。
// 它返回一个等价于 s 的 rune（Unicode 字符码点）切片。
// md5:10939a65a6e24e32
# <翻译结束>


<原文开始>
// Replace returns a copy of the slice s with the first n
// non-overlapping instances of old replaced by new.
// If old is empty, it matches at the beginning of the slice
// and after each UTF-8 sequence, yielding up to k+1 replacements
// for a k-rune slice.
// If n < 0, there is no limit on the number of replacements.
<原文结束>

# <翻译开始>
// Replace 函数返回一个切片 s 的副本，其中旧值 old 的第一个非重叠实例被新值 new 替换。如果 old 为空，它会在切片的开头和每个 UTF-8 序列之后匹配，最多替换 k+1 次，对于长度为 k 的切片。
// 如果 n 小于 0，替换次数没有限制。
// md5:f7ca90ba486844bf
# <翻译结束>


<原文开始>
// ReplaceAll returns a copy of the slice s with all
// non-overlapping instances of old replaced by new.
// If old is empty, it matches at the beginning of the slice
// and after each UTF-8 sequence, yielding up to k+1 replacements
// for a k-rune slice.
<原文结束>

# <翻译开始>
// ReplaceAll 返回一个副本，其中s中的所有非重叠的old被new替换。
// 如果old为空，它将在切片的开头和每个UTF-8序列后匹配，对于一个长度为k的切片，最多可以进行k+1次替换。
// md5:458b06a900a614cb
# <翻译结束>


<原文开始>
// EqualFold reports whether s and t, interpreted as UTF-8 strings,
// are equal under simple Unicode case-folding, which is a more general
// form of case-insensitivity.
<原文结束>

# <翻译开始>
// EqualFold 判断作为 UTF-8 字符串的 s 和 t，在执行简单的 Unicode 大小写折叠后是否相等。这是一种更普遍的大小写不敏感比较方式。
// md5:bf180308fee8f0a4
# <翻译结束>


<原文开始>
// Index returns the index of the first instance of sep in s, or -1 if sep is not present in s.
<原文结束>

# <翻译开始>
// Index 返回字符串 s 中 sep 的第一个实例的索引，如果 sep 不在 s 中，则返回 -1。. md5:b9f2aeee6a0f3c43
# <翻译结束>


<原文开始>
// Cut slices s around the first instance of sep,
// returning the text before and after sep.
// The found result reports whether sep appears in s.
// If sep does not appear in s, cut returns s, nil, false.
//
// Cut returns slices of the original slice s, not copies.
<原文结束>

# <翻译开始>
// 在s中第一个出现sep的位置周围切割，返回sep前后的文本。
// 找到的结果报告sep是否出现在s中。如果sep不在s中，cut将返回s、nil和false。
// 
// cut返回的是原始切片s的切片，而不是副本。
// md5:d93766c814c08f4c
# <翻译结束>


<原文开始>
// Clone returns a copy of b[:len(b)].
// The result may have additional unused capacity.
// Clone(nil) returns nil.
<原文结束>

# <翻译开始>
// Clone 返回 b[:len(b)] 的副本。
// 结果可能具有额外的未使用容量。
// Clone(nil) 返回 nil。
// md5:104eb96c6cfc5cc9
# <翻译结束>


<原文开始>
// CutPrefix returns s without the provided leading prefix byte slice
// and reports whether it found the prefix.
// If s doesn't start with prefix, CutPrefix returns s, false.
// If prefix is the empty byte slice, CutPrefix returns s, true.
//
// CutPrefix returns slices of the original slice s, not copies.
<原文结束>

# <翻译开始>
// CutPrefix 函数返回一个字符串 s，去除了给定的前缀字节切片，并报告是否找到了该前缀。
// 如果 s 不以 prefix 开头，则返回 s, false。
// 如果 prefix 是一个空字节切片，则返回 s, true。
//
// CutPrefix 返回的是原切片 s 的子切片，而不是副本。
// md5:03cd59cf1b324d12
# <翻译结束>


<原文开始>
// CutSuffix returns s without the provided ending suffix byte slice
// and reports whether it found the suffix.
// If s doesn't end with suffix, CutSuffix returns s, false.
// If suffix is the empty byte slice, CutSuffix returns s, true.
//
// CutSuffix returns slices of the original slice s, not copies.
<原文结束>

# <翻译开始>
// CutSuffix 函数返回 s 去除给定的结束后缀字节切片后的结果，并报告是否找到了该后缀。
// 如果 s 不以 suffix 结束，CutSuffix 返回 s，false。
// 如果 suffix 是空的字节切片，CutSuffix 返回 s，true。
//
// CutSuffix 返回原始切片 s 的切片，而不是复制的新切片。
// md5:f882290c414b24fa
# <翻译结束>

