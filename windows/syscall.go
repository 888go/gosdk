// 版权所有 ? 2009 Go作者。保留所有权利。
// 本源代码的使用受BSD风格
// 许可证约束，该许可证可在LICENSE文件中找到。

//go:build windows

// Package windows 提供对底层操作系统原语的接口。具体操作系统的细节因底层系统不同而异，且默认情况下，godoc 将显示当前系统的特定文档。如果你想让 godoc 显示其他系统的系统调用文档，请将 $GOOS 和 $GOARCH 环境变量分别设置为所期望的系统和架构。例如，如果你想在 linux/amd64 系统上查看 freebsd/arm 的文档，应将 $GOOS 设为 freebsd，$GOARCH 设为 arm。
// 
// 本包主要用途在于为其他提供更通用系统接口的包（如 "os"、"time" 和 "net"）内部使用。如果可以，请优先使用这些包而非本包。
// 
// 关于本包中函数和数据类型的详细信息，请参考对应操作系统的手册。
// 
// 这些函数通过返回 err == nil 表示成功；否则，err 代表描述失败情况的操作系统错误，并持有 syscall.Errno 类型的值。
package windows // 导入 "golang.org/x/sys/windows"

import (
	"bytes"
	"strings"
	"syscall"
	"unsafe"
)

// ByteSliceFromString 返回一个以NUL字节结尾的、包含字符串s文本内容的字节切片。
// 若s在任意位置包含NUL字节，将返回(nil, syscall.EINVAL)。
func ByteSliceFromString(s string) ([]byte, error) {
	if strings.IndexByte(s, 0) != -1 {
		return nil, syscall.EINVAL
	}
	a := make([]byte, len(s)+1)
	copy(a, s)
	return a, nil
}

// BytePtrFromString 返回一个指向以 NUL 结尾的字节数组的指针，该数组包含字符串 s 的文本。如果 s 在任何位置包含 NUL 字节，则返回 (nil, syscall.EINVAL)。
func BytePtrFromString(s string) (*byte, error) {
	a, err := ByteSliceFromString(s)
	if err != nil {
		return nil, err
	}
	return &a[0], nil
}

// ByteSliceToString 将切片 s 表示的文本转换为字符串形式，删除终止空字符（NUL）及其后的所有字节。
func ByteSliceToString(s []byte) string {
	if i := bytes.IndexByte(s, 0); i != -1 {
		s = s[:i]
	}
	return string(s)
}

// BytePtrToString 接收一个指向文本序列的指针，并返回相应的字符串。
// 若该指针为nil，则返回空字符串。其假设文本序列以零字节结束；
// 若未出现零字节，程序可能会崩溃。
func BytePtrToString(p *byte) string {
	if p == nil {
		return ""
	}
	if *p == 0 {
		return ""
	}

	// Find NUL terminator.
	n := 0
	for ptr := unsafe.Pointer(p); *(*byte)(ptr) != 0; n++ {
		ptr = unsafe.Pointer(uintptr(ptr) + 1)
	}

	return string(unsafe.Slice(p, n))
}

// 单词形式的零，用于当我们需要一个指向0字节的有效指针时。
// 参见 mksyscall.pl。
var _zero uintptr

func (ts *Timespec) Unix() (sec int64, nsec int64) {
	return int64(ts.Sec), int64(ts.Nsec)
}

func (tv *Timeval) Unix() (sec int64, nsec int64) {
	return int64(tv.Sec), int64(tv.Usec) * 1000
}

func (ts *Timespec) Nano() int64 {
	return int64(ts.Sec)*1e9 + int64(ts.Nsec)
}

func (tv *Timeval) Nano() int64 {
	return int64(tv.Sec)*1e9 + int64(tv.Usec)*1000
}
