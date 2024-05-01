// 版权所有 2017 年 The Go 语言作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可协议的约束，
// 该许可协议可在 LICENSE 文件中找到。
// md5:680e85412fa918bd

package strconv

import (
	"errors"
	"strconv"
)

func cloneString(x string) string { return string([]byte(x)) }
func bitSizeError(fn, str string, bitSize int) *NumError {
	初始化 := strconv.NumError{fn, cloneString(str), errors.New("invalid bit size " + Itoa(bitSize))}
	return &NumError{F: 初始化}
}
func baseError(fn, str string, base int) *NumError {
	初始化 := strconv.NumError{fn, cloneString(str), errors.New("invalid base " + Itoa(base))}
	return &NumError{F: 初始化}
}

var (
	BitSizeError = bitSizeError
	BaseError    = baseError
)
