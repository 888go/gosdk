// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build js

package unix

import (
	"syscall"
	_ "unsafe"
)


// ff:
// from:
// flags:
// p:
// fd:
func RecvfromInet4(fd int, p []byte, flags int, from *syscall.SockaddrInet4) (int, error) {
	return 0, syscall.ENOSYS
}


// ff:
// err:
// n:
// from:
// flags:
// p:
// fd:
func RecvfromInet6(fd int, p []byte, flags int, from *syscall.SockaddrInet6) (n int, err error) {
	return 0, syscall.ENOSYS
}


// ff:
// err:
// to:
// flags:
// p:
// fd:
func SendtoInet4(fd int, p []byte, flags int, to *syscall.SockaddrInet4) (err error) {
	return syscall.ENOSYS
}


// ff:
// err:
// to:
// flags:
// p:
// fd:
func SendtoInet6(fd int, p []byte, flags int, to *syscall.SockaddrInet6) (err error) {
	return syscall.ENOSYS
}


// ff:
// err:
// n:
// flags:
// to:
// oob:
// p:
// fd:
func SendmsgNInet4(fd int, p, oob []byte, to *syscall.SockaddrInet4, flags int) (n int, err error) {
	return 0, syscall.ENOSYS
}


// ff:
// err:
// n:
// flags:
// to:
// oob:
// p:
// fd:
func SendmsgNInet6(fd int, p, oob []byte, to *syscall.SockaddrInet6, flags int) (n int, err error) {
	return 0, syscall.ENOSYS
}


// ff:
// err:
// recvflags:
// oobn:
// n:
// from:
// flags:
// oob:
// p:
// fd:
func RecvmsgInet4(fd int, p, oob []byte, flags int, from *syscall.SockaddrInet4) (n, oobn int, recvflags int, err error) {
	return 0, 0, 0, syscall.ENOSYS
}


// ff:
// err:
// recvflags:
// oobn:
// n:
// from:
// flags:
// oob:
// p:
// fd:
func RecvmsgInet6(fd int, p, oob []byte, flags int, from *syscall.SockaddrInet6) (n, oobn int, recvflags int, err error) {
	return 0, 0, 0, syscall.ENOSYS
}
