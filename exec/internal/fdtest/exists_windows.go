// 版权所有 ? 2021 Go作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

//go:build windows

package fdtest

// Exists 在 Windows 上未实现，调用时会引发 panic。
func Exists(fd uintptr) bool {
	panic("unimplemented")
}
