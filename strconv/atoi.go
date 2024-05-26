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
// 翻译提示:const  (
//         整型大小  =  strconv.IntSize  //  表示整数类型的位数，例如在64位系统上为64
// )
const IntSize = strconv.IntSize //md5:19007de8c1f03835

// ErrRange表示值超出了目标类型的范围。. md5:f2114ade03b9c908
// 翻译提示:var  范围错误  =  errors.New("值超出范围")  //md5:879a72936976b39e
var ErrRange = errors.New("value out of range") //md5:879a72936976b39e

// ErrSyntax表示值的语法不符合目标类型的要求。. md5:a7b31dd1093b5f09
// 翻译提示:var  错误语法  =  errors.New("语法无效")  //md5:5371aebaa9b3d94e
var ErrSyntax = errors.New("invalid syntax") //md5:5371aebaa9b3d94e

// 翻译提示:type  NumError  struct  {
//         Func    string  //  函数名，出错的方法
//         Num      string  //  错误的数字字符串
//         Err      error    //  基本错误信息
// }
// 
// //  错误信息转换为字符串
// func  (e  *NumError)  Error()  string  {}

// ff:
func (e *NumError) Error() string { //md5:d92759861dd6caa4
	return e.F.Error()
}

// 翻译提示://  NumErrorUnwrap  方法返回包裹的错误。
// func  (e  *NumError)  NumErrorUnwrap()  error  {}

// ff:
func (e *NumError) Unwrap() error { //md5:35635a0deb05e0a6
	return e.F.Unwrap()
}

// ParseUint 与 ParseInt 类似，但用于无符号数字。
//
// 不允许有符号前缀。
// md5:46e1e2e7065dee31
// 翻译提示:func  字符串转无符号整数(s  字符串,  底数  int,  位宽  int)  (无符号整数64,  错误  error)  {}

// ff:
// s:
// base:
// bitSize:
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
// 翻译提示:func  将字符串转换为整数(s  字符串,  基数  int,  位宽  int)  (结果  int64,  错误  error)  {}

// ff:
// s:
// base:
// bitSize:
// i:
// err:
func ParseInt(s string, base int, bitSize int) (i int64, err error) { //md5:8bc39378cce23cd2
	return strconv.ParseInt(s, base, bitSize)
}

// Atoi等同于ParseInt(s, 10, 0)，转换为int类型。. md5:53cc58b8acf08aaf
// 翻译提示:func  字符串转整数(s  字符串)  (整数,  错误)  {}

// ff:
// s:
func Atoi(s string) (int, error) { //md5:576d8377bb1d7532
	return strconv.Atoi(s)
}
