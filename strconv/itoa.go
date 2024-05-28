package strconv

import "strconv"

// FormatUint 将整数 i 转换为给定基数的字符串表示形式，其中 2 <= 基数 <= 36。当数字值大于等于 10 时，结果使用小写字母 'a' 到 'z' 表示数字。
// md5:5acf46fee8f8891d
// 翻译提示:func 格式化无符号整数(i 无符号长整型, base 进制) 字符串 {}

// ff:正整数64位到文本
// i:正整数64位
// base:进制
func FormatUint(i uint64, base int) string { //md5:25e76e61cf965490
	return strconv.FormatUint(i, base)
}

// FormatInt 返回 i 在给定基数（2 <= base <= 36）下的字符串表示形式。结果使用小写字母 'a' 到 'z' 来表示值大于等于 10 的数字。
// md5:b517dddedb91df08
// 翻译提示:func IntToText(数值 int64, 底数 int) 字符串 {}

// ff:整数64位到文本
// i:整数64位
// base:进制
func FormatInt(i int64, base int) string { //md5:c8ffdd9b6b31204e
	return strconv.FormatInt(i, base)
}

// Itoa 等同于 FormatInt(int64(i), 10)。. md5:77e0389bc2cc1fe3
// 翻译提示:func IntToAlpha(i int) string {}

// ff:整数到文本
// i:整数
func Itoa(i int) string { //md5:855a975d9bf4f4e5
	return strconv.Itoa(i)
}

// AppendInt 将整数 i 的字符串形式（由 FormatInt 生成）附加到 dst 后面，并返回扩展后的缓冲区。
// md5:8cb3415bb501072e
// 翻译提示:func 追加整数到字节切片(dst []byte, 值 int64, 底数 int) []byte {}

// ff:追加整数64位到字节集
// dst:字节集
// i:整数64位
// base:进制
func AppendInt(dst []byte, i int64, base int) []byte { //md5:f8f82dfdc15b8e82
	return strconv.AppendInt(dst, i, base)
}

// AppendUint 将无符号整数 i 的字符串形式（由 FormatUint 生成）追加到 dst 中，并返回扩展后的缓冲区。
// md5:707f7e65bd7aac9f
// 翻译提示:func 追加Uint(dst 字节切片, i 无符号64位整数, base 底数) 字节切片 {}

// ff:追加正整数64位到字节集
// dst:字节集
// i:正整数64位
// base:进制
func AppendUint(dst []byte, i uint64, base int) []byte { //md5:7faa1c64ec12eba6
	return strconv.AppendUint(dst, i, base)
}
