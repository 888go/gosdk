// 版权所有 ? 2015 The Go 作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

//go:build windows

package registry

func (k Key) SetValue(name string, valtype uint32, data []byte) error {
	return k.setValue(name, valtype, data)
}
