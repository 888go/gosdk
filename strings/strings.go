package strings

import (
	"strings"
	"unicode"
)

// Count 计算字符串 s 中非重叠子串 substr 出现的次数。
// 若 substr 为空字符串，Count 返回 1 加上 s 中的 Unicode 编码点数量。
func Count(s, substr string) int { //md5:8dc4508cfbea56c6d406f8a4f9770e09
	return strings.Count(s, substr)
}

// Contains 判断 substr 是否包含于 s 中
func Contains(s, substr string) bool { //md5:e5367376f35a456e5c106746b475b47c
	return strings.Contains(s, substr)
}

// ContainsAny 判断字符串 s 中是否包含 chars 中的任意一个Unicode字符。
func ContainsAny(s, chars string) bool { //md5:3c6df9a7716addd6267b4d1f1e7d021c
	return strings.ContainsAny(s, chars)
}

// ContainsRune报告Unicode码点r是否在s中。
func ContainsRune(s string, r rune) bool { //md5:4238898d042dfb1776a3a6e6466358b2
	return strings.ContainsRune(s, r)
}

// LastIndex 返回子串 substr 在字符串 s 中最后一次出现的索引，若 substr 不在 s 中，则返回 -1。
func LastIndex(s, substr string) int { //md5:f71c5a442b965bb5bfaeed7e5c7c9576
	return strings.LastIndex(s, substr)
}

// IndexByte 返回 s 中首次出现字符 c 的索引，若 c 不在 s 中，则返回 -1。
func IndexByte(s string, c byte) int { //md5:cc27e81169256c2d002e7ba6cc205863
	return strings.IndexByte(s, c)
}

// IndexRune 返回Unicode码点r在s中首次出现的索引，若r不在s中，则返回-1。
// 若r为utf8.RuneError，它将返回s中任意无效UTF-8字节序列的首次出现位置。
func IndexRune(s string, r rune) int { //md5:ee5f1aa9ff80778af6ab99d3210cca6c
	return strings.IndexRune(s, r)
}

// IndexAny 返回 s 中首次出现的任意 Unicode 码点在 chars 中的索引，若 s 中不存在 chars 中的任何 Unicode 码点，则返回 -1。
func IndexAny(s, chars string) int { //md5:9082d9461cfcdcfe489a822547f285e9
	return strings.IndexAny(s, chars)
}

// LastIndexAny返回s中任意Unicode码点字符在chars中的最后一个实例的索引，若s中不存在chars中的任何Unicode码点，则返回-1。
func LastIndexAny(s, chars string) int { //md5:b7c6c3d4157d4ce87b044ac9be7aed6f
	return strings.LastIndexAny(s, chars)
}

// LastIndexByte 返回 s 中 c 最后一次出现的索引，若 c 不在 s 中，则返回 -1。
func LastIndexByte(s string, c byte) int { //md5:3dc6c63525e2007ee6cccdee606cb585
	return strings.LastIndexByte(s, c)
}

// SplitN 将s按sep分隔符切片，返回由这些分隔符之间的子串组成的切片。
//
// count参数决定返回的子串数量：
//
//	n > 0：最多返回n个子串；最后一个子串为未分割的剩余部分。
//	n == 0：结果为nil（零个子串）
//	n < 0：返回所有子串
//
// 对于s和sep（如空字符串）的边缘情况，处理方式如[Split]文档所述。
//
// 若要按第一个分隔符实例进行分割，请参见Cut。
func SplitN(s, sep string, n int) []string { //md5:2ef42e1e1c8b9efaf8efedea2b1b2b58
	return strings.SplitN(s, sep, n)
}

// SplitAfterN 将s按sep出现的位置进行切片，返回由这些子串组成的切片。
//
// 参数count决定了返回子串的数量：
//
//	n > 0：最多返回n个子串；最后一个子串将是未分割的剩余部分。
//	n == 0：结果为nil（零个子串）
//	n < 0：返回所有子串
//
// 对于s和sep的边缘情况（如空字符串），处理方式与SplitAfter函数文档中所述一致。
func SplitAfterN(s, sep string, n int) []string { //md5:5c4f7c8941408be7caa300dfd6629980
	return strings.SplitAfterN(s, sep, n)
}

// Split 函数将字符串 s 按照分隔符 sep 切分成所有子串，并返回由这些分隔符之间的子串组成的切片。
//
// 若 s 中不包含 sep，且 sep 非空时，Split 返回长度为 1 的切片，其中唯一元素为 s。
//
// 若 sep 为空，Split 将在每个 UTF-8 序列后进行分割。若 s 和 sep 同时为空，Split 返回一个空切片。
//
// 它等同于调用 [SplitN] 时传入计数参数 -1。
//
// 若要按照第一个分隔符实例进行分割，请参见 Cut 函数。
func Split(s, sep string) []string { //md5:9e39ccbb6ab10559dc6dfa7be187dc75
	return strings.Split(s, sep)
}

// SplitAfter 将字符串 s 按照每个出现的 sep 实例进行切分，并返回所有子串组成的切片。
//
// 若 s 中不包含 sep，且 sep 非空时，SplitAfter 返回长度为 1 的切片，其中唯一元素即为 s。
//
// 若 sep 为空，SplitAfter 将在每个 UTF-8 序列后进行切分。若 s 和 sep 同时为空，SplitAfter 返回一个空切片。
//
// 它等同于使用计数为 -1 的 [SplitAfterN]。
func SplitAfter(s, sep string) []string { //md5:7feaa482c8e32826bdbe24bf223ee6e7
	return strings.SplitAfter(s, sep)
}

// Fields 函数将字符串 s 以其中连续出现的一个或多个空格字符（由 unicode.IsSpace 定义）为分隔，返回 s 的子字符串切片。如果 s 只包含空格，则返回一个空切片。
func Fields(s string) []string { //md5:1f8999413d2d4e0022170f15e0baae55
	return strings.Fields(s)
}

// FieldsFunc 函数以满足 f(c) 条件的每个连续的 Unicode 编码点序列将字符串 s 分割开来，并返回 s 的切片数组。如果 s 中的所有编码点均满足 f(c) 或 s 为空，则返回一个空切片。
//
// FieldsFunc 不保证调用 f(c) 的顺序，且假定对于给定的 c，f 总是返回相同的值。
func FieldsFunc(s string, f func(rune) bool) []string { //md5:fb9dfb813e048fcf1d47059175b1f561
	return strings.FieldsFunc(s, f)
}

// Join 函数将第一个参数中的元素连接起来，生成一个单一的字符串。分隔符字符串 sep 会被放置在结果字符串中各元素之间。
func Join(elems []string, sep string) string { //md5:00ef0cc76f825778bdb1f1c6ab82a33d
	return strings.Join(elems, sep)
}

// HasPrefix 判断字符串 s 是否以 prefix 作为前缀。
func HasPrefix(s, prefix string) bool { //md5:e1f46e88f0e4ea58e45b8de9709c6e08
	return strings.HasPrefix(s, prefix)
}

// HasSuffix 判断字符串 s 是否以 suffix 结尾。
func HasSuffix(s, suffix string) bool { //md5:1ddeca0a9e56022d90d75810b20fb6da
	return strings.HasSuffix(s, suffix)
}

// Map 函数返回一个字符串 s 的副本，其中所有字符均根据映射函数进行修改。
// 若映射函数返回负值，则该字符将从字符串中删除，不作替换。
func Map(mapping func(rune) rune, s string) string { //md5:63ed0c242381265245671e75ecccb37f
	return strings.Map(mapping, s)
}

// Repeat 返回一个新字符串，其中包含字符串 s 的 count 份副本。
//
// 如果 count 为负数，或者 (len(s) * count) 的结果溢出时，将会引发恐慌。
func Repeat(s string, count int) string { //md5:d49e11180de57b1bad15b47471444a40
	return strings.Repeat(s, count)
}

// ToUpper返回s，其中所有Unicode字母均映射为大写。
func ToUpper(s string) string { //md5:3debb184e946c25c8eac8f6f13baedd6
	return strings.ToUpper(s)
}

// ToLower返回s，其中所有Unicode字母已映射至其小写形式。
func ToLower(s string) string { //md5:474a927c2e62fa6c5288c04fd40232ce
	return strings.ToLower(s)
}

// ToTitle返回一个字符串s的副本，其中所有Unicode字母均映射至其Unicode标题大小写。
func ToTitle(s string) string { //md5:7cb166b3ddebfd1125f25ad2a7a6bc0f
	return strings.ToTitle(s)
}

// ToUpperSpecial返回字符串s的副本，其中所有Unicode字母均根据c指定的大小写映射规则转换为大写。
func ToUpperSpecial(c unicode.SpecialCase, s string) string { //md5:b4e286fb9418724e1543c7af9b9e3af9
	return strings.ToUpperSpecial(c, s)
}

// ToLowerSpecial返回一个字符串s的副本，其中所有Unicode字母都根据c指定的大小写映射规则转换为小写。
func ToLowerSpecial(c unicode.SpecialCase, s string) string { //md5:dc9a6ec7663a80cdd586a2e9772618ac
	return strings.ToLowerSpecial(c, s)
}

// ToTitleSpecial 返回字符串 s 的副本，其中所有 Unicode 字母都映射到其
// Unicode 标题大小写形式，并优先考虑特殊大小写规则。
func ToTitleSpecial(c unicode.SpecialCase, s string) string { //md5:6eebd87af27dac15039ad70bc35e7318
	return strings.ToTitleSpecial(c, s)
}

// ToValidUTF8 返回字符串 s 的副本，其中每个无效的 UTF-8 字节序列均被替换为替换字符串（可能为空）。
func ToValidUTF8(s, replacement string) string { //md5:422f67249cb248a3bd9c0ab9b1d097e1
	return strings.ToValidUTF8(s, replacement)
}

// Title返回字符串s的一个副本，其中所有作为单词开头的Unicode字母均转换为其Unicode标题大小写。
//
// 已弃用：Title用于识别单词边界的方式无法正确处理Unicode标点符号。请改用golang.org/x/text/cases。
func Title(s string) string { //md5:adaf62896a837fe0a4b94cef489e74c6
	return strings.Title(s)
}

// TrimLeftFunc返回字符串s的一个切片，其中所有满足f(c)的起始Unicode码点c已被移除。
func TrimLeftFunc(s string, f func(rune) bool) string { //md5:516609201171fccebe5fce4ada48d1ad
	return strings.TrimLeftFunc(s, f)
}

// TrimRightFunc返回字符串s的一个切片，其中所有满足f(c)的尾部Unicode码点均被移除。
func TrimRightFunc(s string, f func(rune) bool) string { //md5:280f57b7a6e49458e7ab99785d88cfc4
	return strings.TrimRightFunc(s, f)
}

// TrimFunc返回字符串s的一个切片，其中所有满足f(c)的首尾Unicode字符码被移除。
func TrimFunc(s string, f func(rune) bool) string { //md5:f83f48e3eb47b62b7e1ecdc80aa8229b
	return strings.TrimFunc(s, f)
}

// IndexFunc返回s中第一个满足f(c)的Unicode码点的索引，如果不存在满足条件的码点，则返回-1。
func IndexFunc(s string, f func(rune) bool) int { //md5:22c6b63d2713ff3020a8ced746a3bcdb
	return strings.IndexFunc(s, f)
}

// LastIndexFunc返回s中最后一个满足f(c)的Unicode码点的索引，若无满足条件的码点，则返回-1。
func LastIndexFunc(s string, f func(rune) bool) int { //md5:c29da2b69561585e838689d903e72dac
	return strings.LastIndexFunc(s, f)
}

// Trim返回字符串s的一个切片，其中包含的所有在cutset中的前导和尾部Unicode字符码点已被移除。
func Trim(s, cutset string) string { //md5:d4e63cddca34d67410bbe36516a460d4
	return strings.Trim(s, cutset)
}

// TrimLeft返回字符串s的一个切片，其中所有位于开头的
// Unicode编码点（如果存在于cutset中）均被移除。
//
// 要移除前缀，请使用[TrimPrefix]代替。
func TrimLeft(s, cutset string) string { //md5:900b2617ce88019f221e18c5d48d88af
	return strings.TrimLeft(s, cutset)
}

// TrimRight返回字符串s的一个切片，其中所有尾部包含在cutset中的Unicode字符点已被移除。
//
// 要移除后缀，应使用[TrimSuffix]代替。
func TrimRight(s, cutset string) string { //md5:5264a3ff78fb7d36909186e6a8dbef64
	return strings.TrimRight(s, cutset)
}

// TrimSpace返回字符串s的一个切片，其中已移除了所有Unicode定义的开头和结尾空白字符。
func TrimSpace(s string) string { //md5:3d7d6d251ac5a73c7e895d7eab45c28c
	return strings.TrimSpace(s)
}

// TrimPrefix返回字符串s，其中已移除了提供的前缀字符串。
// 若s以prefix开头，则返回去掉prefix后的s；否则返回s保持原样不变。
func TrimPrefix(s, prefix string) string { //md5:7b252f4ef2548b2526b142c154ee698c
	return strings.TrimPrefix(s, prefix)
}

// TrimSuffix 函数返回去掉指定尾部后缀字符串后的 s。
// 若 s 末尾不包含该后缀，则返回 s 保持原样不变。
func TrimSuffix(s, suffix string) string { //md5:48e4d1998c003b0feefefd543d4a55dd
	return strings.TrimSuffix(s, suffix)
}

// Replace返回一个字符串s的副本，其中前n个非重叠的old子串被new替换。
// 若old为空，它将在字符串起始处以及每个UTF-8序列之后匹配，对于一个包含k个rune字符的字符串，最多可进行k+1次替换。
// 若n < 0，则替换次数无限制。
func Replace(s, old, new string, n int) string { //md5:f43175906f03445467e9423a396f9774
	return strings.Replace(s, old, new, n)
}

// ReplaceAll返回字符串s的副本，其中所有非重叠的old实例均被new替换。
// 如果old为空，它将在字符串起始处以及每个UTF-8序列之后匹配，对于一个包含k个符文的字符串，最多产生k+1次替换。
func ReplaceAll(s, old, new string) string { //md5:9174341d7ee38be79885b5c0fd72091e
	return strings.ReplaceAll(s, old, new)
}

// EqualFold报告当s和t以UTF-8字符串形式解释时，它们在简单Unicode大小写折叠规则下是否相等。这是一种更通用的大小写不敏感比较方式。
func EqualFold(s, t string) bool { //md5:c0895dc9a6f5e3ac2092811656c1c6aa
	return strings.EqualFold(s, t)
}

// Index 返回字符串 s 中首次出现子串 substr 的索引，若 substr 不在 s 中，则返回 -1。
func Index(s, substr string) int { //md5:7c8755646e2ccdea7c6a5360d8d1e685
	return strings.Index(s, substr)
}

// Cut 函数在字符串切片 s 中首次出现 sep 位置处进行切割，
// 返回 sep 之前及之后的文本。
// found 结果用于指示 sep 是否在 s 中出现。
// 若 sep 未在 s 中出现，则 Cut 返回 s，空字符串，false。
func Cut(s, sep string) (before, after string, found bool) { //md5:101ffdd375b539b69a47d8fc1f0cb22d
	return strings.Cut(s, sep)
}

// CutPrefix 函数返回去掉指定前缀字符串后的 s，并报告是否找到了该前缀。
// 若 s 未以 prefix 开头，CutPrefix 返回 s, false。
// 若 prefix 为空字符串，CutPrefix 返回 s, true。
func CutPrefix(s, prefix string) (after string, found bool) { //md5:9ac4eac0ec5c4a09c1e7ce3f8cfefe95
	return strings.CutPrefix(s, prefix)
}

// CutSuffix 函数返回去掉指定结尾后缀字符串的 s，并报告是否找到了该后缀。
// 若 s 未以 suffix 结尾，CutSuffix 返回 s, false。
// 若 suffix 为空字符串，CutSuffix 返回 s, true。
func CutSuffix(s, suffix string) (before string, found bool) { //md5:e0b07650a60a8f446976a06d09da4193
	return strings.CutSuffix(s, suffix)
}
