package strings

import "strings"

// Clone 返回一个 s 的全新副本。
// 它确保将 s 复制到新的内存分配中，
// 这对于仅保留一个较大字符串中的小部分子串而言非常重要。
// 使用 Clone 可帮助此类程序减少内存使用。当然，由于 Clone 会进行复制，
// 对 Clone 的过度使用可能会使程序占用更多内存。
// Clone 通常应仅在极少数情况下使用，并且仅当性能分析表明需要时才使用。
// 对于长度为零的字符串，将返回空字符串 ""，且不会进行内存分配。
// 翻译提示:func 克隆字符串(s 字符串) 字符串 {}

// ff:深拷贝
// s:文本
func Clone(s string) string { //md5:8c05310f5095d610d6ac0200dcd5ba45
	return strings.Clone(s)
}
