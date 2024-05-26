// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package unix

import (
	"syscall"
	"unsafe"
)

// 翻译提示:func 文件范围复制(srcFD int, srcOffset *int64, dstFD int, dstOffset *int64, length int, flags int) (复制字节数 int, 错误 error) {}

// ff:
// err:
// n:
// flags:
// len:
// woff:
// wfd:
// roff:
// rfd:
func CopyFileRange(rfd int, roff *int64, wfd int, woff *int64, len int, flags int) (n int, err error) {
	r1, _, errno := syscall.Syscall6(copyFileRangeTrap,
		uintptr(rfd),
		uintptr(unsafe.Pointer(roff)),
		uintptr(wfd),
		uintptr(unsafe.Pointer(woff)),
		uintptr(len),
		uintptr(flags),
	)
	n = int(r1)
	if errno != 0 {
		err = errno
	}
	return
}
