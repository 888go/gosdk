// 版权所有 ? 2021 Go作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

//---build---//go:build windows || plan9 || (js && wasm) || wasip1

package testenv

import (
	"errors"
	"io/fs"
	"os"
)

// Sigquit 用于向挂起的子进程发送信号以将其终止。
// 在 Unix 系统中，我们发送 SIGQUIT 信号；但在非 Unix 系统中，我们仅使用 os.Kill。
var Sigquit = os.Kill

func syscallIsNotSupported(err error) bool {
	return errors.Is(err, fs.ErrPermission) || errors.Is(err, errors.ErrUnsupported)
}
