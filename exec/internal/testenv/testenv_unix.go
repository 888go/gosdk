// 版权所有 ? 2021 Go作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

//go:build unix

package testenv

import (
	"errors"
	"io/fs"
	"syscall"
)

// Sigquit 是发送给挂起子进程以杀死它的信号。
// 发送 SIGQUIT 以获取堆栈跟踪。
var Sigquit = syscall.SIGQUIT

func syscallIsNotSupported(err error) bool {
	if err == nil {
		return false
	}

	var errno syscall.Errno
	if errors.As(err, &errno) {
		switch errno {
		case syscall.EPERM, syscall.EROFS:
// 用户缺少权限：要么该调用需要root权限而用户并非root，要么该调用被容器安全策略所拒绝。
			return true
		case syscall.EINVAL:
// 有些容器在系统调用因安全策略被拒绝时，返回EINVAL而非EPERM。
			return true
		}
	}

	if errors.Is(err, fs.ErrPermission) || errors.Is(err, errors.ErrUnsupported) {
		return true
	}

	return false
}
