
<原文开始>
// ParseFloat converts the string s to a floating-point number
// with the precision specified by bitSize: 32 for float32, or 64 for float64.
// When bitSize=32, the result still has type float64, but it will be
// convertible to float32 without changing its value.
//
// ParseFloat accepts decimal and hexadecimal floating-point numbers
// as defined by the Go syntax for [floating-point literals].
// If s is well-formed and near a valid floating-point number,
// ParseFloat returns the nearest floating-point number rounded
// using IEEE754 unbiased rounding.
// (Parsing a hexadecimal floating-point value only rounds when
// there are more bits in the hexadecimal representation than
// will fit in the mantissa.)
//
// The errors that ParseFloat returns have concrete type *NumError
// and include err.Num = s.
//
// If s is not syntactically well-formed, ParseFloat returns err.Err = ErrSyntax.
//
// If s is syntactically well-formed but is more than 1/2 ULP
// away from the largest floating point number of the given size,
// ParseFloat returns f = ±Inf, err.Err = ErrRange.
//
// ParseFloat recognizes the string "NaN", and the (possibly signed) strings "Inf" and "Infinity"
// as their respective special floating point values. It ignores case when matching.
//
// [floating-point literals]: https://go.dev/ref/spec#Floating-point_literals
<原文结束>

# <翻译开始>
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
# <翻译结束>

