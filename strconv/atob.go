package strconv

import "strconv"

// ParseBool 将字符串解析为布尔值。
// 它接受 1、t、T、TRUE、true、True、0、f、F、FALSE、false、False。
// 对于其他任何值，它将返回一个错误。
// md5:03bedf3a8d17a25d
func ParseBool(str string) (bool, error) { //md5:6f1b6c8e55a90306
	return strconv.ParseBool(str)
}

// FormatBool 根据b的值返回 "true" 或 "false"。. md5:fcfe59671d45b2c8
func FormatBool(b bool) string { //md5:b07d776a75e22602
	return strconv.FormatBool(b)
}

// AppendBool 将根据 b 的值（true 或 false）追加到 dst 后面，并返回扩展后的缓冲区。
// md5:84476359e5f68f99
func AppendBool(dst []byte, b bool) []byte { //md5:cfa4028967205e76
	return strconv.AppendBool(dst, b)
}
