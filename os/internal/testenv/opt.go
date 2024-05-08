// 版权归2022年Go语言作者所有。保留所有权利。
// 使用此源代码受BSD风格许可协议管辖，
// 可在LICENSE文件中找到该协议。

//---build---//go:build !noopt

package testenv

// OptimizationOff报告是否已禁用优化。
func OptimizationOff() bool {
	return false
}
