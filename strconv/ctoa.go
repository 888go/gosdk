package strconv

import "strconv"

// FormatComplex 将复数 c 转换为字符串，形式为 (a+bi)，其中 a 和 b 是实部和虚部，根据格式 fmt 和精度 prec 进行格式化。
// 
// 格式 fmt 和精度 prec 的含义与 FormatFloat 中的相同。它假设原始值是从具有 bitSize 位的复数中获取的，对于 complex64 必须是 64 位，对于 complex128 必须是 128 位，并对结果进行四舍五入。
// md5:fd8d5d38ed45847e

// ff:
// bitSize:
// prec:
// fmt:
// c:
func FormatComplex(c complex128, fmt byte, prec, bitSize int) string { //md5:a78e72e77ed11d14
	return strconv.FormatComplex(c, fmt, prec, bitSize)
}
