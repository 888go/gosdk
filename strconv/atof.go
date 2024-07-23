package strconv

import "strconv"

var optimize = true // set to false to force slow-path conversions for testing

// ParseFloat 将字符串 s 转换为浮点数，精度由 bitSize 指定：32 表示 float32，64 表示 float64。
// 当 bitSize=32 时，结果仍然为 float64 类型，但转换为 float32 不会改变其值。
//
// ParseFloat 接受 Go 语法中 [浮点数字面量] 定义的十进制和十六进制浮点数。
// 如果 s 形式正确且接近一个有效的浮点数，ParseFloat 返回最接近的浮点数，
// 使用 IEEE754 无偏舍入规则进行四舍五入。
// （解析十六进制浮点值时，只有当十六进制表示中的位数多于浮点数小数部分所能容纳的位数时才会进行四舍五入。）
//
// ParseFloat 返回的错误具有具体类型 *NumError，并且包括 err.Num = s。
//
// 如果 s 的语法不正确，ParseFloat 返回 err.Err = ErrSyntax。
//
// 如果 s 语法正确，但超出给定大小的最大浮点数的 1/2 机器精度（ULP），ParseFloat 返回 f = ±Inf，err.Err = ErrRange。
//
// ParseFloat 识别字符串 "NaN"，以及可能带符号的字符串 "Inf" 和 "Infinity" 作为它们各自的特殊浮点值。
// 匹配时忽略大小写。
//
// [浮点数字面量]: https://go.dev/ref/spec#Floating-point_literals
// md5:07a94f81ad4011bc
// ff:文本到小数64位
// s:文本
// bitSize:小数精度
func ParseFloat(s string, bitSize int) (float64, error) { //md5:0c9ee835e542b3ea
	return strconv.ParseFloat(s, bitSize)
}
