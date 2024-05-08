// 版权所有 2018 The Go Authors。保留所有权利。
// 使用本源代码须遵循 BSD 风格许可证，
// 该许可证可在 LICENSE 文件中找到。

//---build---//go:build windows && go1.9

package windows

import "syscall"

type Errno = syscall.Errno
type SysProcAttr = syscall.SysProcAttr
