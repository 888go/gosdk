// 版权所有 ? 2021 Go作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

//---build---//go:build !unix && !windows

package exec_test

import "os"

var (
	quitSignal os.Signal = nil
	pipeSignal os.Signal = nil
)
