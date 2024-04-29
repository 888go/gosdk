// 版权所有 ? 2022 Go作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

//go:build unix

package exec_test

import (
	"os"
	"syscall"
)

var (
	quitSignal os.Signal = syscall.SIGQUIT
	pipeSignal os.Signal = syscall.SIGPIPE
)
