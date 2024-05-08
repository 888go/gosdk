// 版权所有 2021 Go 作者。保留所有权利。
// 本源代码的使用受BSD风格许可证管辖，该许可证可在LICENSE文件中找到。

package windows

import (
	"sync"
	"syscall"
	_ "unsafe"
)

//go:linkname WSASendtoInet4 syscall.wsaSendtoInet4
//go:noescape
func WSASendtoInet4(s syscall.Handle, bufs *syscall.WSABuf, bufcnt uint32, sent *uint32, flags uint32, to *syscall.SockaddrInet4, overlapped *syscall.Overlapped, croutine *byte) (err error)

//go:linkname WSASendtoInet6 syscall.wsaSendtoInet6
//go:noescape
func WSASendtoInet6(s syscall.Handle, bufs *syscall.WSABuf, bufcnt uint32, sent *uint32, flags uint32, to *syscall.SockaddrInet6, overlapped *syscall.Overlapped, croutine *byte) (err error)

const (
	SIO_TCP_INITIAL_RTO                    = syscall.IOC_IN | syscall.IOC_VENDOR | 17
	TCP_INITIAL_RTO_UNSPECIFIED_RTT        = ^uint16(0)
	TCP_INITIAL_RTO_NO_SYN_RETRANSMISSIONS = ^uint8(1)
)

type TCP_INITIAL_RTO_PARAMETERS struct {
	Rtt                   uint16
	MaxSynRetransmissions uint8
}

var Support_TCP_INITIAL_RTO_NO_SYN_RETRANSMISSIONS = sync.OnceValue(func() bool {
	var maj, min, build uint32
	rtlGetNtVersionNumbers(&maj, &min, &build)
	return maj >= 10 && build&0xffff >= 16299
})

//go:linkname rtlGetNtVersionNumbers syscall.rtlGetNtVersionNumbers
//go:noescape
func rtlGetNtVersionNumbers(majorVersion *uint32, minorVersion *uint32, buildNumber *uint32)
