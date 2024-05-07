package bytes

import (
	"bytes"
	"unicode"
)

// Equal 判断 a 和 b 是否具有相同的长度并且包含相同的字节。一个空指针（nil）等同于一个空切片（empty slice）。
// md5:dbece9ef3f693465

// ff:是否相等
// b:比较b
// a:比较a
func Equal(a, b []byte) bool { //md5:a460b4c618708bfc
	return bytes.Equal(a, b)
}

// Compare 函数按字典顺序比较两个字节切片。如果 a == b，返回 0；如果 a < b，返回 -1；如果 a > b，返回 +1。
// 空指针（nil）等同于空切片。
// md5:23a14b305b5a0036

// ff:顺序比较
// b:比较b
// a:比较a
func Compare(a, b []byte) int { //md5:c280bc577895ac56
	return bytes.Compare(a, b)
}

// Count 计算字符串 s 中非重叠实例 sep 的数量。
// 如果 sep 是一个空切片，Count 返回 1 加上 s 中 UTF-8 编码的代码点数量。
// md5:5d59a1b45cedfba7

// ff:统计次数
// sep:待统计
// s:原字节集
func Count(s, sep []byte) int { //md5:8dc4508cfbea56c6
	return bytes.Count(s, sep)
}

// Contains 函数报告子切片是否在 b 中。. md5:3075713b6bb2e961

// ff:是否包含
// subslice:待查找
// b:原字节集
func Contains(b, subslice []byte) bool { //md5:e5367376f35a456e
	return bytes.Contains(b, subslice)
}

// ContainsAny报告chars中的任何UTF-8编码的代码点是否都在b中。. md5:a7123747027d0a57

// ff:是否包含任意字符
// chars:检索字符s
// b:原字节集
func ContainsAny(b []byte, chars string) bool { //md5:3c6df9a7716addd6
	return bytes.ContainsAny(b, chars)
}

// ContainsRune 检查 rune 是否存在于 UTF-8 编码的字节切片 b 中。. md5:22a878ae15a5922f

// ff:是否包含Unicode字符
// r:检索字符
// b:原字节集
func ContainsRune(b []byte, r rune) bool { //md5:4238898d042dfb17
	return bytes.ContainsRune(b, r)
}

// ContainsFunc 报告UTF-8编码的码点r在b中是否存在，满足条件f(r)。. md5:e58efa31a9ff987a

// ff:是否包含任意字符FUNC
// f:回调函数
// b:原字节集
func ContainsFunc(b []byte, f func(rune) bool) bool { //md5:7f263561c791a82d
	return bytes.ContainsFunc(b, f)
}

// IndexByte 返回字节切片b中字符c的第一个出现位置，如果c不在b中，则返回-1。. md5:7bf26378aad4cd0b

// ff:取字节索引
// c:检索字节
// b:原字节集
func IndexByte(b []byte, c byte) int { //md5:cc27e81169256c2d
	return bytes.IndexByte(b, c)
}

// LastIndex 返回 s 中 sep 的最后一个实例的索引，如果 s 中不存在 sep，则返回 -1。. md5:6bd49252012e79ad

// ff:取字节集最后索引
// sep:待检索
// s:原字节集
func LastIndex(s, sep []byte) int { //md5:f71c5a442b965bb5
	return bytes.LastIndex(s, sep)
}

// LastIndexByte 返回字符串s中字符c的最后一个出现位置的索引，如果c不在s中，则返回-1。. md5:8a4f0f6af8efe072

// ff:取字节最后索引
// c:字节
// s:原字节集
func LastIndexByte(s []byte, c byte) int { //md5:3dc6c63525e2007e
	return bytes.LastIndexByte(s, c)
}

// IndexRune 将 s 解释为一系列 UTF-8 编码的代码点。
// 它返回给定字符在 s 中首次出现的字节索引。
// 如果字符在 s 中不存在，则返回 -1。
// 如果 r 是 utf8.RuneError，它将返回任何无效 UTF-8 字节序列的第一个实例。
// md5:6dc9763770576138

// ff:取Unicode字符索引
// r:检索字符
// s:原字节集
func IndexRune(s []byte, r rune) int { //md5:ee5f1aa9ff80778a
	return bytes.IndexRune(s, r)
}

// IndexAny将s视为一个用UTF-8编码的Unicode码点序列。
// 它返回s中第一个出现的、在chars中任一Unicode码点的字节索引。如果chars为空或者没有共同的码点，它返回-1。
// md5:03920b48034c033b

// ff:取任意字符索引
// chars:检索字符s
// s:原字节集
func IndexAny(s []byte, chars string) int { //md5:9082d9461cfcdcfe
	return bytes.IndexAny(s, chars)
}

// LastIndexAny将s解释为UTF-8编码的Unicode代码点序列。它返回s中与chars中的任何Unicode代码点的最后一个出现的字节索引。如果chars为空或没有公共代码点，它将返回-1。
// md5:7c8b4adaf5973148

// ff:取任意字符最后索引
// chars:检索字符
// s:原字节集
func LastIndexAny(s []byte, chars string) int { //md5:b7c6c3d4157d4ce8
	return bytes.LastIndexAny(s, chars)
}

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

// ff:分割并按数量
// n:返回数量
// sep:分隔符
// s:原字节集
func SplitN(s, sep []byte, n int) [][]byte { //md5:2ef42e1e1c8b9efa
	return bytes.SplitN(s, sep, n)
}

// SplitAfterN 将字符串 s 根据 sep 的每个实例进行切分，之后返回这些子切片的切片。
// 如果 sep 为空，SplitAfterN 会根据每个 UTF-8 序列进行切分。
// count 参数决定了要返回的子切片数量：
//
//	n > 0: 最多 n 个子切片；最后一个子切片将是未被切分的剩余部分。
//	n == 0: 结果为 nil（零个子切片）
//	n < 0: 所有子切片
//
// md5:93042132420aa229

// ff:分割并按数量2
// n:返回数量
// sep:分隔符
// s:原字节集
func SplitAfterN(s, sep []byte, n int) [][]byte { //md5:5c4f7c8941408be7
	return bytes.SplitAfterN(s, sep, n)
}

// Split 将切片 s 按照分隔符 sep 切分成多个子切片，并返回一个包含这些子切片的切片。
// 如果 sep 为空，Split 将在每个 UTF-8 字符序列后进行分割。
// 它等价于 SplitN 函数，传入的计数参数为 -1。
//
// 若要围绕第一个分隔符进行分割，可以使用 Cut 函数。
// md5:d5d8f14cae52de3f

// ff:分割
// sep:分隔符
// s:原字节集
func Split(s, sep []byte) [][]byte { //md5:9e39ccbb6ab10559
	return bytes.Split(s, sep)
}

// SplitAfter 在每个sep实例之后将s分割成所有子切片，并返回这些子切片的切片。
// 如果sep为空，SplitAfter会在每个UTF-8序列后进行分割。
// 它等同于SplitAfterN，其中计数为-1。
// md5:c78b01ae43df412f

// ff:分割2
// sep:分隔符
// s:原字节集
func SplitAfter(s, sep []byte) [][]byte { //md5:7feaa482c8e32826
	return bytes.SplitAfter(s, sep)
}

// Fields 将 s 解析为一系列 UTF-8 编码的代码点。
// 它会在 s 中每个连续或非连续的空白字符（由 unicode.IsSpace 定义）周围分割，返回 s 的子切片的切片，如果 s 只包含空白字符，则返回空切片。
// md5:7c06fb47034bd649

// ff:分割并按空白
// s:原字节集
func Fields(s []byte) [][]byte { //md5:1f8999413d2d4e00
	return bytes.Fields(s)
}

// FieldsFunc 将 s 解释为一系列 UTF-8 编码的代码点。
// 它会在满足条件 f(c) 的每个代码点序列处将切片 s 分割，
// 并返回 s 的子切片切片。如果 s 中的所有代码点都满足 f(c)，或者
// len(s) == 0，则返回一个空切片。
//
// FieldsFunc 不保证以特定顺序调用 f(c)，
// 并假设对于给定的 c，f 总是返回相同的结果值。
// md5:e0df308c2de6e89c

// ff:分割FUNC
// f:回调函数
// s:原字节集
func FieldsFunc(s []byte, f func(rune) bool) [][]byte { //md5:fb9dfb813e048fcf
	return bytes.FieldsFunc(s, f)
}

// Join 将 s 中的元素连接起来创建一个新的字节切片。分隔符 sep 会被插入到结果切片中的元素之间。
// md5:c65602b4a3a28272

// ff:连接
// sep:连接符
// s:字节集切片
func Join(s [][]byte, sep []byte) []byte { //md5:00ef0cc76f825778
	return bytes.Join(s, sep)
}

// HasPrefix 函数报告切片 s 是否以 prefix 开头。. md5:aeb71c9e082f4e79

// ff:是否为前缀
// prefix:前缀
// s:原字节集
func HasPrefix(s, prefix []byte) bool { //md5:e1f46e88f0e4ea58
	return bytes.HasPrefix(s, prefix)
}

// HasSuffix 函数报告 byte slice s 是否以 suffix 结尾。. md5:42621549774f0616

// ff:是否为后缀
// suffix:后缀
// s:原字节集
func HasSuffix(s, suffix []byte) bool { //md5:1ddeca0a9e56022d
	return bytes.HasSuffix(s, suffix)
}

// Map 函数返回一个副本，该副本中的字节切片 s 的所有字符都根据映射函数进行了修改。
// 如果映射函数返回一个负值，则该字符将从字节切片中移除且不做替换。
// 字节切片 s 中的字符以及输出结果都按 UTF-8 编码的码点进行解释。
// md5:f29d9994e9e93cbe

// ff:遍历修改FUNC
// s:原字节集
// mapping:回调函数
// r:
func Map(mapping func(r rune) rune, s []byte) []byte { //md5:63ed0c2423812652
	return bytes.Map(mapping, s)
}

// Repeat 返回一个新的字节切片，由 count 个 b 的副本组成。
//
// 如果 count 为负数，或者 (len(b) * count) 的结果溢出，则该函数会 panic。
// md5:a30126d911ce9e3a

// ff:生成重复字节集
// count:重复次数
// b:原字节集
func Repeat(b []byte, count int) []byte { //md5:d49e11180de57b1b
	return bytes.Repeat(b, count)
}

// ToUpper返回一个副本，其中s中的所有Unicode字母映射到大写。
// md5:701556ffdae09171

// ff:到大写
// s:原字节集
func ToUpper(s []byte) []byte { //md5:3debb184e946c25c
	return bytes.ToUpper(s)
}

// ToLower返回一个副本，其中s中的所有Unicode字母映射到小写。
// md5:96198f3611614a06

// ff:到小写
// s:原字节集
func ToLower(s []byte) []byte { //md5:474a927c2e62fa6c
	return bytes.ToLower(s)
}

// ToTitle将s视为UTF-8编码的字节，并返回一个副本，其中所有Unicode字母都被映射为它们的标题大小写形式。. md5:3a15836e7e4e4976

// ff:到大写2
// s:原字节集
func ToTitle(s []byte) []byte { //md5:7cb166b3ddebfd11
	return bytes.ToTitle(s)
}

// ToUpperSpecial将s视为UTF-8编码的字节，并返回一个副本，其中所有Unicode字母都映射为其大写形式，优先考虑特殊大小写规则。
// md5:7a8fc8bbf32666bf

// ff:到大写并按规则
// s:原字节集
// c:大小写规则
func ToUpperSpecial(c unicode.SpecialCase, s []byte) []byte { //md5:b4e286fb9418724e
	return bytes.ToUpperSpecial(c, s)
}

// ToLowerSpecial将s视为UTF-8编码的字节，并返回一个副本，其中所有Unicode字母都被映射到小写，优先考虑特殊情况规则。
// md5:1c197945b5358611

// ff:到小写并按规则
// s:原字节集
// c:大小写规则
func ToLowerSpecial(c unicode.SpecialCase, s []byte) []byte { //md5:dc9a6ec7663a80cd
	return bytes.ToLowerSpecial(c, s)
}

// ToTitleSpecial 将 s 视为 UTF-8 编码的字节，并返回一个副本，其中所有 Unicode 字母都被映射为其大写形式，优先遵循特殊的大写规则。
// md5:05fe1b145cbd5bda

// ff:到大写2并按规则
// s:原字节集
// c:大小写规则
func ToTitleSpecial(c unicode.SpecialCase, s []byte) []byte { //md5:6eebd87af27dac15
	return bytes.ToTitleSpecial(c, s)
}

// ToValidUTF8 将s视为UTF-8编码的字节序列，并返回一个副本，其中每个代表无效UTF-8的字节序列都被replacement中的字节替换，replacement可以是空字符串。
// md5:a8577838395945fd

// ff:替换无效UTF8字节
// replacement:替换符
// s:原字节集
func ToValidUTF8(s, replacement []byte) []byte { //md5:422f67249cb248a3
	return bytes.ToValidUTF8(s, replacement)
}

// Title 将 s 视为 UTF-8 编码的字节，并返回一个副本，其中所有单词起始的Unicode字母都被转换为标题大小写。
//
// 已废弃：Title 使用的单词边界规则无法正确处理Unicode标点符号。请改用 golang.org/x/text/cases。
// md5:489de93c97e62d72

// ff:弃用Title
// s:原字节集
func Title(s []byte) []byte { //md5:adaf62896a837fe0
	return bytes.Title(s)
}

// TrimLeftFunc将s视为UTF-8编码的字节，并通过移除所有满足条件f(c)的UTF-8编码的代码点c，返回s的子切片。
// md5:21e9929d142ec743

// ff:删起始字符FUNC
// f:回调函数
// r:
// s:原字节集
func TrimLeftFunc(s []byte, f func(r rune) bool) []byte { //md5:516609201171fcce
	return bytes.TrimLeftFunc(s, f)
}

// TrimRightFunc通过移除所有满足条件f(c)的UTF-8编码的码点c，返回s的子切片。
// md5:3016ff9a3bf64754

// ff:删尾部字符FUNC
// f:回调函数
// r:
// s:原字节集
func TrimRightFunc(s []byte, f func(r rune) bool) []byte { //md5:280f57b7a6e49458
	return bytes.TrimRightFunc(s, f)
}

// TrimFunc 返回s的子切片，通过去除所有满足f(c)的前导和尾随UTF-8编码的字符c。
// md5:a1e6af674d552e9a

// ff:删首尾字符FUNC
// f:回调函数
// r:
// s:原字节集
func TrimFunc(s []byte, f func(r rune) bool) []byte { //md5:f83f48e3eb47b62b
	return bytes.TrimFunc(s, f)
}

// TrimPrefix 函数返回字符串 s 去除指定前缀后的结果。
// 如果 s 不以 prefix 开头，那么将原样返回 s。
// md5:97cb4f309ba2947f

// ff:删前缀
// prefix:前缀
// s:原字节集
func TrimPrefix(s, prefix []byte) []byte { //md5:7b252f4ef2548b25
	return bytes.TrimPrefix(s, prefix)
}

// TrimSuffix 函数返回 s，但不包含提供的尾随后缀字符串。
// 如果 s 不以 suffix 结尾，则返回原始的 s。
// md5:6d9d00533816bde2

// ff:删后缀
// suffix:后缀
// s:原字节集
func TrimSuffix(s, suffix []byte) []byte { //md5:48e4d1998c003b0f
	return bytes.TrimSuffix(s, suffix)
}

// IndexFunc 将 s 解释为一系列 UTF-8 编码的代码点。
// 它返回满足函数 f(c) 的第一个 Unicode 代码点在 s 中的字节索引，如果没有满足条件的，则返回 -1。
// md5:3faa26ff439023a4

// ff:查找字符索引FUNC
// f:回调函数
// r:
// s:原字节集
func IndexFunc(s []byte, f func(r rune) bool) int { //md5:22c6b63d2713ff30
	return bytes.IndexFunc(s, f)
}

// LastIndexFunc 将 s 解释为一系列 UTF-8 编码的代码点。
// 它返回 s 中满足 f(c) 的最后一个 Unicode 代码点的字节索引，
// 如果没有满足条件的，则返回 -1。
// md5:d29cc3d1192d35e5

// ff:查找最后字符索引FUNC
// f:回调函数
// r:
// s:原字节集
func LastIndexFunc(s []byte, f func(r rune) bool) int { //md5:c29da2b69561585e
	return bytes.LastIndexFunc(s, f)
}

// Trim返回一个s的子切片，切掉cutset中包含的所有领先和尾随的UTF-8编码的字符。
// md5:c8162b6bb6472995

// ff:删首尾字符
// cutset:
// s:原字节集
func Trim(s []byte, cutset string) []byte { //md5:d4e63cddca34d674
	return bytes.Trim(s, cutset)
}

// TrimLeft通过移除s中所有UTF-8编码的代码点开始部分，返回一个子切片。
// md5:0908d5b6be56ad64

// ff:删起始字符
// cutset:
// s:原字节集
func TrimLeft(s []byte, cutset string) []byte { //md5:900b2617ce88019f
	return bytes.TrimLeft(s, cutset)
}

// TrimRight 通过裁剪掉所有由 cutset 所包含的 UTF-8 编码字符点，返回 s 的子切片。
// md5:f6b6ebb338bed118

// ff:删尾部字符
// cutset:
// s:原字节集
func TrimRight(s []byte, cutset string) []byte { //md5:5264a3ff78fb7d36
	return bytes.TrimRight(s, cutset)
}

// TrimSpace 返回 s 的子切片，通过去除所有前导和尾随的空白字符，这些空白字符根据 Unicode 定义。
// md5:bd70164451a18ab6

// ff:删首尾空
// s:原字节集
func TrimSpace(s []byte) []byte { //md5:3d7d6d251ac5a73c
	return bytes.TrimSpace(s)
}

// Runes 将 s 解释为一个使用UTF-8编码的字符序列。
// 它返回一个等价于 s 的 rune（Unicode 字符码点）切片。
// md5:10939a65a6e24e32

// ff:转换为Unicode字符切片
// s:字节集
func Runes(s []byte) []rune { //md5:bde235d2d7ce8207
	return bytes.Runes(s)
}

// Replace 函数返回一个切片 s 的副本，其中旧值 old 的第一个非重叠实例被新值 new 替换。如果 old 为空，它会在切片的开头和每个 UTF-8 序列之后匹配，最多替换 k+1 次，对于长度为 k 的切片。
// 如果 n 小于 0，替换次数没有限制。
// md5:f7ca90ba486844bf

// ff:替换
// n:次数
// new:新字节集
// old:旧字节集
// s:原字节集
func Replace(s, old, new []byte, n int) []byte { //md5:f43175906f034454
	return bytes.Replace(s, old, new, n)
}

// ReplaceAll 返回一个副本，其中s中的所有非重叠的old被new替换。
// 如果old为空，它将在切片的开头和每个UTF-8序列后匹配，对于一个长度为k的切片，最多可以进行k+1次替换。
// md5:458b06a900a614cb

// ff:替换全部
// new:新字节集
// old:旧字节集
// s:原字节集
func ReplaceAll(s, old, new []byte) []byte { //md5:9174341d7ee38be7
	return bytes.ReplaceAll(s, old, new)
}

// ff:
// EqualFold 判断作为 UTF-8 字符串的 s 和 t，在执行简单的 Unicode 大小写折叠后是否相等。这是一种更普遍的大小写不敏感比较方式。
// md5:bf180308fee8f0a4

// ff:是否相等并忽略大小写
// t:
// s:
func EqualFold(s, t []byte) bool { //md5:c0895dc9a6f5e3ac
	return bytes.EqualFold(s, t)
}

// Index 返回字符串 s 中 sep 的第一个实例的索引，如果 sep 不在 s 中，则返回 -1。. md5:b9f2aeee6a0f3c43

// ff:取字节集索引
// sep:待查找
// s:原字节集
func Index(s, sep []byte) int { //md5:7c8755646e2ccdea
	return bytes.Index(s, sep)
}

// 在s中第一个出现sep的位置周围切割，返回sep前后的文本。
// 找到的结果报告sep是否出现在s中。如果sep不在s中，cut将返回s、nil和false。
//
// cut返回的是原始切片s的切片，而不是副本。
// md5:d93766c814c08f4c

// ff:
// found:
// after:
// before:
// sep:
// s:
func Cut(s, sep []byte) (before, after []byte, found bool) { //md5:101ffdd375b539b6
	return bytes.Cut(s, sep)
}

// Clone 返回 b[:len(b)] 的副本。
// 结果可能具有额外的未使用容量。
// Clone(nil) 返回 nil。
// md5:104eb96c6cfc5cc9

// ff:深拷贝
// b:原字节集
func Clone(b []byte) []byte { //md5:8c05310f5095d610
	return bytes.Clone(b)
}

// CutPrefix 函数返回一个字符串 s，去除了给定的前缀字节切片，并报告是否找到了该前缀。
// 如果 s 不以 prefix 开头，则返回 s, false。
// 如果 prefix 是一个空字节切片，则返回 s, true。
//
// CutPrefix 返回的是原切片 s 的子切片，而不是副本。
// md5:03cd59cf1b324d12

// ff:删前缀2
// found:是否成功
// after:返回
// prefix:前缀
// s:原字节集
func CutPrefix(s, prefix []byte) (after []byte, found bool) { //md5:9ac4eac0ec5c4a09
	return bytes.CutPrefix(s, prefix)
}

// CutSuffix 函数返回 s 去除给定的结束后缀字节切片后的结果，并报告是否找到了该后缀。
// 如果 s 不以 suffix 结束，CutSuffix 返回 s，false。
// 如果 suffix 是空的字节切片，CutSuffix 返回 s，true。
//
// CutSuffix 返回原始切片 s 的切片，而不是复制的新切片。
// md5:f882290c414b24fa

// ff:删后缀2
// found:是否成功
// before:返回
// suffix:后缀
// s:原字节集
func CutSuffix(s, suffix []byte) (before []byte, found bool) { //md5:e0b07650a60a8f44
	return bytes.CutSuffix(s, suffix)
}
