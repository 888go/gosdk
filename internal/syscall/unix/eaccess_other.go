// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build unix && !linux

package unix

import "syscall"

// 翻译提示:func 访问权限检查(path string, 模式 uint32) error {}

// ff:
// mode:
// path:
func Eaccess(path string, mode uint32) error {
	return syscall.ENOSYS
}
