package strconv//bm:转换类

import "strconv"

// ParseBool 将字符串解析为布尔值。
// 它接受 1、t、T、TRUE、true、True、0、f、F、FALSE、false、False。
// 对于其他任何值，它将返回一个错误。
// md5:03bedf3a8d17a25d
// ff:文本到布尔值
// str:布尔文本
func ParseBool(str string) (bool, error) { //md5:6f1b6c8e55a90306
	return strconv.ParseBool(str)
}

// FormatBool 根据b的值返回 "true" 或 "false"。. md5:fcfe59671d45b2c8
// ff:布尔值到文本
// b:布尔值
func FormatBool(b bool) string { //md5:b07d776a75e22602
	return strconv.FormatBool(b)
}

// AppendBool 将根据 b 的值（true 或 false）追加到 dst 后面，并返回扩展后的缓冲区。
// md5:84476359e5f68f99
// ff:布尔值加到字节集
// dst:字节集
// b:布尔值
func AppendBool(dst []byte, b bool) []byte { //md5:cfa4028967205e76
	return strconv.AppendBool(dst, b)
}
