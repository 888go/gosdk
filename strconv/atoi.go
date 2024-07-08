package strconv

import (
	"errors"
	"strconv"
)

// NumError 记录了一次转换失败的情况。. md5:4bc87fcc15d03a84
type NumError struct {
	F strconv.NumError
} //md5:0cd5b5c6c9f477b9

// IntSize 是int或uint值的位数。. md5:c5f70ea31d120aca
const IntSize = strconv.IntSize //md5:19007de8c1f03835

// ErrRange表示值超出了目标类型的范围。. md5:f2114ade03b9c908
var ErrRange = errors.New("value out of range") //md5:879a72936976b39e

// ErrSyntax表示值的语法不符合目标类型的要求。. md5:a7b31dd1093b5f09
var ErrSyntax = errors.New("invalid syntax") //md5:5371aebaa9b3d94e

func (e *NumError) Error() string { //md5:d92759861dd6caa4
	return e.F.Error()
}

func (e *NumError) Unwrap() error { //md5:35635a0deb05e0a6
	return e.F.Unwrap()
}

// ParseUint 与 ParseInt 类似，但用于无符号数字。
//
// 不允许有符号前缀。
// md5:46e1e2e7065dee31
func ParseUint(s string, base int, bitSize int) (uint64, error) { //md5:f4c2829256142922
	return strconv.ParseUint(s, base, bitSize)
}

// ParseInt 将给定基数（0、2 到 36）和位大小（0 到 64）的字符串 s 解析为相应的整数值 i。
//
// 字符串可能以正负号开头："+" 或 "-"。
//
// 如果基数参数为 0，真正的基数由符号后的字符串前缀决定：如果存在，"0b" 表示 2，"0" 或 "0o" 表示 8，"0x" 表示 16，否则表示 10。仅当基数为 0 时，允许使用下划线字符，遵循 Go 语言整数字面量的语法定义。
//
// 位大小参数指定结果必须适应的整数类型。位大小 0、8、16、32 和 64 分别对应于 int、int8、int16、int32 和 int64。
// 如果位大小低于 0 或超过 64，将返回一个错误。
//
// ParseInt 返回的错误具有具体类型 *NumError，并包含 err.Num = s。如果 s 空或包含无效的数字，err.Err = ErrSyntax，返回值为 0；如果 s 对应的值不能用给定大小的有符号整数表示，err.Err = ErrRange，返回值为相应位大小和符号的最大整数值。
//
// [整数字面量]：https://go.dev/ref/spec#Integer_literals
// md5:c7272ea65b20e967
func ParseInt(s string, base int, bitSize int) (i int64, err error) { //md5:8bc39378cce23cd2
	return strconv.ParseInt(s, base, bitSize)
}

// Atoi等同于ParseInt(s, 10, 0)，转换为int类型。. md5:53cc58b8acf08aaf
func Atoi(s string) (int, error) { //md5:576d8377bb1d7532
	return strconv.Atoi(s)
}
