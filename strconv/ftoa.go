package strconv

import "strconv"

// FormatFloat converts the floating-point number f to a string,
// according to the format fmt and precision prec. It rounds the
// result assuming that the original was obtained from a floating-point
// value of bitSize bits (32 for float32, 64 for float64).
//
// The format fmt is one of
// 'b' (-ddddp±ddd, a binary exponent),
// 'e' (-d.dddde±dd, a decimal exponent),
// 'E' (-d.ddddE±dd, a decimal exponent),
// 'f' (-ddd.dddd, no exponent),
// 'g' ('e' for large exponents, 'f' otherwise),
// 'G' ('E' for large exponents, 'f' otherwise),
// 'x' (-0xd.ddddp±ddd, a hexadecimal fraction and binary exponent), or
// 'X' (-0Xd.ddddP±ddd, a hexadecimal fraction and binary exponent).
//
// The precision prec controls the number of digits (excluding the exponent)
// printed by the 'e', 'E', 'f', 'g', 'G', 'x', and 'X' formats.
// For 'e', 'E', 'f', 'x', and 'X', it is the number of digits after the decimal point.
// For 'g' and 'G' it is the maximum number of significant digits (trailing
// zeros are removed).
// The special precision -1 uses the smallest number of digits
// necessary such that ParseFloat will return f exactly.
func FormatFloat(f float64, fmt byte, prec, bitSize int) string { //md5:627500a705d2a46e
	return strconv.FormatFloat(f, fmt, prec, bitSize)
}

// AppendFloat 将浮点数 f 的格式化字符串（由 FormatFloat 生成）追加到 dst 中，并返回扩展的缓冲区。
// md5:5f65576d0cc9f940
func AppendFloat(dst []byte, f float64, fmt byte, prec, bitSize int) []byte { //md5:6c1e86744a16fe25
	return strconv.AppendFloat(dst, f, fmt, prec, bitSize)
}
