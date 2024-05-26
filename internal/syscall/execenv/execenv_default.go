// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !windows

package execenv

import "syscall"

// Default will return the default environment
// variables based on the process attributes
// provided.
//
// Defaults to syscall.Environ() on all platforms
// other than Windows.
// 翻译提示:func 默认环境变量设置(sys *系统进程属性) ([]字符串, 错误) {}

// ff:
// sys:
func Default(sys *syscall.SysProcAttr) ([]string, error) {
	return syscall.Environ(), nil
}
