// 版权所有 ? 2015 The Go Authors。保留所有权利。
// 本源代码的使用受 BSD 风格许可协议约束，
// 该协议可在 LICENSE 文件中找到。

//go:build windows
// +build windows

package registry

func (k Key) SetValue(name string, valtype uint32, data []byte) error {
	return k.setValue(name, valtype, data)
}
