// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package unix

import (
	"sync/atomic"
	"syscall"
	"unsafe"
)

//go:cgo_import_dynamic libc_getrandom getrandom "libc.so"

//go:linkname procGetrandom libc_getrandom

var procGetrandom uintptr

var getrandomUnsupported int32 // atomic

// GetRandomFlag is a flag supported by the getrandom system call.
type GetRandomFlag uintptr

// 翻译提示:const (
//	// 非阻塞模式意味着返回EAGAIN而不是阻塞。
//	非阻塞模式 GetRandomFlag = 0x0001
//
//	// 随机源标志仅出于兼容性目的设置，在FreeBSD上无效。
//	随机源标志 GetRandomFlag = 0x0002
//)
const (
	// GRND_NONBLOCK means return EAGAIN rather than blocking.
	GRND_NONBLOCK GetRandomFlag = 0x0001

	// GRND_RANDOM means use the /dev/random pool instead of /dev/urandom.
	GRND_RANDOM GetRandomFlag = 0x0002
)

// GetRandom calls the getrandom system call.
// 翻译提示:func 获取随机字节(p []字节, 标志 GetRandomFlag) (长度 int, 错误 error) {}

// ff:
// err:
// n:
// flags:
// p:
func GetRandom(p []byte, flags GetRandomFlag) (n int, err error) {
	if len(p) == 0 {
		return 0, nil
	}
	if atomic.LoadInt32(&getrandomUnsupported) != 0 {
		return 0, syscall.ENOSYS
	}
	r1, _, errno := syscall6(uintptr(unsafe.Pointer(&procGetrandom)),
		3,
		uintptr(unsafe.Pointer(&p[0])),
		uintptr(len(p)),
		uintptr(flags),
		0, 0, 0)
	if errno != 0 {
		if errno == syscall.ENOSYS {
			atomic.StoreInt32(&getrandomUnsupported, 1)
		}
		return 0, errno
	}
	return int(r1), nil
}
