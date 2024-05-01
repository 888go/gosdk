
<原文开始>
// FormatComplex converts the complex number c to a string of the
// form (a+bi) where a and b are the real and imaginary parts,
// formatted according to the format fmt and precision prec.
//
// The format fmt and precision prec have the same meaning as in FormatFloat.
// It rounds the result assuming that the original was obtained from a complex
// value of bitSize bits, which must be 64 for complex64 and 128 for complex128.
<原文结束>

# <翻译开始>
// FormatComplex 将复数 c 转换为字符串，形式为 (a+bi)，其中 a 和 b 是实部和虚部，根据格式 fmt 和精度 prec 进行格式化。
// 
// 格式 fmt 和精度 prec 的含义与 FormatFloat 中的相同。它假设原始值是从具有 bitSize 位的复数中获取的，对于 complex64 必须是 64 位，对于 complex128 必须是 128 位，并对结果进行四舍五入。
// md5:fd8d5d38ed45847e
# <翻译结束>

