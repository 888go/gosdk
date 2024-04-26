// 版权所有 2021 Go 作者。保留所有权利。
// 本源代码的使用受BSD风格许可证管辖，该许可证可在LICENSE文件中找到。

//go:build windows || plan9 || (js && wasm) || wasip1

package testenv

import (
	"errors"
	"io/fs"
	"os"
)

// Sigquit是用于杀死挂起子进程的信号。
// 在Unix系统中，我们发送SIGQUIT信号，但在非Unix系统中，我们只使用os.Kill。
var Sigquit = os.Kill

func syscallIsNotSupported(err error) bool {
	return errors.Is(err, fs.ErrPermission) || errors.Is(err, errors.ErrUnsupported)
}
