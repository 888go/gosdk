// 版权所有 2021 Go 作者。保留所有权利。
// 本源代码的使用受BSD风格许可证管辖，该许可证可在LICENSE文件中找到。

//---build---//go:build unix

package testenv

import (
	"errors"
	"io/fs"
	"syscall"
)

// Sigquit 是用来杀死挂起子进程的信号。
// 发送 SIGQUIT 信号可以获得堆栈跟踪。
var Sigquit = syscall.SIGQUIT

func syscallIsNotSupported(err error) bool {
	if err == nil {
		return false
	}

	var errno syscall.Errno
	if errors.As(err, &errno) {
		switch errno {
		case syscall.EPERM, syscall.EROFS:
// 用户缺少权限：要么调用需要root权限，但用户不是root，要么被容器安全策略拒绝。
			return true
		case syscall.EINVAL:
// 如果系统调用被安全策略拒绝，一些容器可能会返回EINVAL而不是EPERM。
			return true
		}
	}

	if errors.Is(err, fs.ErrPermission) || errors.Is(err, errors.ErrUnsupported) {
		return true
	}

	return false
}
