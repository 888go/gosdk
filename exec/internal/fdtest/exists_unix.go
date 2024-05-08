// 版权所有 ? 2021 Go作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

//---build---//go:build unix || wasm

// 包 fdtest 提供了在 exec 中处理文件描述符的测试辅助工具。
package fdtest

import (
	"syscall"
)

// Exists 返回 true，如果 fd 是一个有效的文件描述符。
func Exists(fd uintptr) bool {
	var s syscall.Stat_t
	err := syscall.Fstat(int(fd), &s)
	return err != syscall.EBADF
}
