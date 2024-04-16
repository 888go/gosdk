// 版权所有 ? 2021 Go作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

//go:build plan9

package fdtest

import (
	"syscall"
)

const errBadFd = syscall.ErrorString("fd out of range or not open")

// Exists 返回 true，如果 fd 是一个有效的文件描述符。

// ff:
// fd:
func Exists(fd uintptr) bool {
	var buf [1]byte
	_, err := syscall.Fstat(int(fd), buf[:])
	return err != errBadFd
}
