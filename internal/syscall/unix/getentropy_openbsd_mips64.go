// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package unix

import (
	"syscall"
	"unsafe"
)

// getentropy(2)'s syscall number, from /usr/src/sys/kern/syscalls.master
const entropyTrap uintptr = 7

// GetEntropy calls the OpenBSD getentropy system call.
// 翻译提示:func 获取熵(p []byte) error {}

// ff:
// p:
func GetEntropy(p []byte) error {
	_, _, errno := syscall.Syscall(entropyTrap,
		uintptr(unsafe.Pointer(&p[0])),
		uintptr(len(p)),
		0)
	if errno != 0 {
		return errno
	}
	return nil
}
