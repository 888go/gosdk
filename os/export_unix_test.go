// 版权所有 ? 2019 The Go Authors。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

//---build---//go:build unix || (js && wasm) || wasip1

package os

var SplitPath = splitPath
