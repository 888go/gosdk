package strconv

import "strconv"

// ParseComplex 函数将字符串 s 转换为复数，其精度由 bitSize 指定：64 表示 complex64，128 表示 complex128。
// 当 bitSize=64 时，结果的类型仍为 complex128，但它可以无损地转换为 complex64。
//
// s 表示的数字必须形如 N、Ni 或 N±Ni，其中 N 代表一个通过 ParseFloat 识别的浮点数，i 表示虚部。
// 如果第二个 N 是正数，则在两个部分之间需要有一个 + 号，如同 ± 所示。如果第二个 N 是 NaN，则只接受 + 号。
// 形式上可以使用括号包围且不能包含任何空格。
// 生成的复数由两个通过 ParseFloat 转换的组件组成。
//
// ParseComplex 返回的错误具有具体类型 *NumError，并包含 err.Num = s。
//
// 如果 s 语法不正确，ParseComplex 返回 err.Err = ErrSyntax。
//
// 如果 s 语法正确，但任一组成部分超出给定大小浮点数最大值的 1/2 ULP 范围，
// ParseComplex 返回 err.Err = ErrRange，并将相应部分的 c 设为 ±Inf。
// md5:2d5101902291657e
// 翻译提示:func 解析复数(s 字符串, 位大小 int) (复数128, 错误) {}

// ff:文本到复数
// s:文本
// bitSize:复数精度
func ParseComplex(s string, bitSize int) (complex128, error) { //md5:11a5d5e6c797768d
	return strconv.ParseComplex(s, bitSize)
}
