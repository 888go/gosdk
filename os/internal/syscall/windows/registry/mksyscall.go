//版权所有2016 The Go Authors。所有权利保留。
//使用此源代码受BSD风格的
//可在LICENSE文件中找到的许可协议管辖。

//---build---//go:build generate

package registry

//go:generate go run ../../../../syscall/mksyscall_windows.go -output zsyscall_windows.go syscall.go
