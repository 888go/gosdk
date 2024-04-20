// 版权所有 ? 2018 The Go 作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可协议约束，
// 该协议可在 LICENSE 文件中找到。

//go:build windows && go1.9

package windows

import "syscall"

type Errno = syscall.Errno
type SysProcAttr = syscall.SysProcAttr
