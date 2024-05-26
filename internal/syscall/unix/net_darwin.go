// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package unix

import (
	"github.com/888go/gosdk/internal/abi"
	"syscall"
	"unsafe"
)

const (
	AI_CANONNAME = 0x2
	AI_ALL       = 0x100
	AI_V4MAPPED  = 0x800
	AI_MASK      = 0x1407

	EAI_AGAIN    = 2
	EAI_NONAME   = 8
	EAI_SYSTEM   = 11
	EAI_OVERFLOW = 14

	NI_NAMEREQD = 4
)

type Addrinfo struct {
	Flags     int32
	Family    int32
	Socktype  int32
	Protocol  int32
	Addrlen   uint32
	Canonname *byte
	Addr      *syscall.RawSockaddr
	Next      *Addrinfo
}

//go:cgo_ldflag "-lresolv"

//go:cgo_import_dynamic libc_getaddrinfo getaddrinfo "/usr/lib/libSystem.B.dylib"
func libc_getaddrinfo_trampoline()

// 翻译提示:func 获取地址信息(hostName, 服务名 *byte, 提示 *Addrinfo, 结果 **Addrinfo) (错误码 int, 错误 error) {}

// ff:
// res:
// hints:
// servname:
// hostname:
func Getaddrinfo(hostname, servname *byte, hints *Addrinfo, res **Addrinfo) (int, error) {
	gerrno, _, errno := syscall_syscall6(abi.FuncPCABI0(libc_getaddrinfo_trampoline),
		uintptr(unsafe.Pointer(hostname)),
		uintptr(unsafe.Pointer(servname)),
		uintptr(unsafe.Pointer(hints)),
		uintptr(unsafe.Pointer(res)),
		0,
		0)
	var err error
	if errno != 0 {
		err = errno
	}
	return int(gerrno), err
}

//go:cgo_import_dynamic libc_freeaddrinfo freeaddrinfo "/usr/lib/libSystem.B.dylib"
func libc_freeaddrinfo_trampoline()

// 翻译提示:func 释放Addrinfo(ai *地址信息) {}

// ff:
// ai:
func Freeaddrinfo(ai *Addrinfo) {
	syscall_syscall6(abi.FuncPCABI0(libc_freeaddrinfo_trampoline),
		uintptr(unsafe.Pointer(ai)),
		0, 0, 0, 0, 0)
}

//go:cgo_import_dynamic libc_getnameinfo getnameinfo "/usr/lib/libSystem.B.dylib"
func libc_getnameinfo_trampoline()

// 翻译提示:func 获取名称信息(sockaddr *syscall.RawSockaddr, sockaddrLength int, 主机名 *byte, 主机名长度 int, 服务名 *byte, 服务名长度 int, 标志 int) (int, error) {}

// ff:
// flags:
// servlen:
// serv:
// hostlen:
// host:
// salen:
// sa:
func Getnameinfo(sa *syscall.RawSockaddr, salen int, host *byte, hostlen int, serv *byte, servlen int, flags int) (int, error) {
	gerrno, _, errno := syscall_syscall9(abi.FuncPCABI0(libc_getnameinfo_trampoline),
		uintptr(unsafe.Pointer(sa)),
		uintptr(salen),
		uintptr(unsafe.Pointer(host)),
		uintptr(hostlen),
		uintptr(unsafe.Pointer(serv)),
		uintptr(servlen),
		uintptr(flags),
		0,
		0)
	var err error
	if errno != 0 {
		err = errno
	}
	return int(gerrno), err
}

//go:cgo_import_dynamic libc_gai_strerror gai_strerror "/usr/lib/libSystem.B.dylib"
func libc_gai_strerror_trampoline()

// 翻译提示:func 错误码转字符串(ecode int) string {}

// ff:
// ecode:
func GaiStrerror(ecode int) string {
	r1, _, _ := syscall_syscall(abi.FuncPCABI0(libc_gai_strerror_trampoline),
		uintptr(ecode),
		0, 0)
	return GoString((*byte)(unsafe.Pointer(r1)))
}

// Implemented in the runtime package.
func gostring(*byte) string

// 翻译提示:函数：GoString
//参数：p *byte（指向字节的指针）
//返回值：string（字符串）
//
//中文重构后：
//
//func 字符串转Go格式(p *字节) 字符串 {}

// ff:
// p:
func GoString(p *byte) string {
	return gostring(p)
}

//go:linkname syscall_syscall syscall.syscall
func syscall_syscall(fn, a1, a2, a3 uintptr) (r1, r2 uintptr, err syscall.Errno)

//go:linkname syscall_syscallPtr syscall.syscallPtr
func syscall_syscallPtr(fn, a1, a2, a3 uintptr) (r1, r2 uintptr, err syscall.Errno)

//go:linkname syscall_syscall6 syscall.syscall6
func syscall_syscall6(fn, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2 uintptr, err syscall.Errno)

//go:linkname syscall_syscall6X syscall.syscall6X
func syscall_syscall6X(fn, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2 uintptr, err syscall.Errno)

//go:linkname syscall_syscall9 syscall.syscall9
func syscall_syscall9(fn, a1, a2, a3, a4, a5, a6, a7, a8, a9 uintptr) (r1, r2 uintptr, err syscall.Errno)

type ResState struct {
	unexported [70]uintptr
}

//go:cgo_import_dynamic libresolv_res_9_ninit res_9_ninit "/usr/lib/libresolv.9.dylib"
func libresolv_res_9_ninit_trampoline()

// 翻译提示:func 初始化ResState(state *解析状态) error {}

// ff:
// state:
func ResNinit(state *ResState) error {
	_, _, errno := syscall_syscall(abi.FuncPCABI0(libresolv_res_9_ninit_trampoline),
		uintptr(unsafe.Pointer(state)),
		0, 0)
	if errno != 0 {
		return errno
	}
	return nil
}

//go:cgo_import_dynamic libresolv_res_9_nclose res_9_nclose "/usr/lib/libresolv.9.dylib"
func libresolv_res_9_nclose_trampoline()

// 翻译提示:func 关闭资源(state *资源状态) {}

// ff:
// state:
func ResNclose(state *ResState) {
	syscall_syscall(abi.FuncPCABI0(libresolv_res_9_nclose_trampoline),
		uintptr(unsafe.Pointer(state)),
		0, 0)
}

//go:cgo_import_dynamic libresolv_res_9_nsearch res_9_nsearch "/usr/lib/libresolv.9.dylib"
func libresolv_res_9_nsearch_trampoline()


// ff:
// anslen:
// ans:
// typ:
// class:
// dname:
// state:
func ResNsearch(state *ResState, dname *byte, class, typ int, ans *byte, anslen int) (int, error) {
	r1, _, errno := syscall_syscall6(abi.FuncPCABI0(libresolv_res_9_nsearch_trampoline),
		uintptr(unsafe.Pointer(state)),
		uintptr(unsafe.Pointer(dname)),
		uintptr(class),
		uintptr(typ),
		uintptr(unsafe.Pointer(ans)),
		uintptr(anslen))
	if errno != 0 {
		return 0, errno
	}
	return int(int32(r1)), nil
}
