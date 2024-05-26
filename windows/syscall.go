// 版权所有 ? 2009 Go 作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 此许可证可在 LICENSE 文件中找到。

//go:build windows

// Package windows 提供了对底层操作系统原语的接口。具体的操作系统细节会因底层系统的不同而有所变化，且默认情况下，godoc 会显示当前系统的特定文档。如果你想让 godoc 显示其他系统的 syscall 文档，请将 $GOOS 和 $GOARCH 环境变量分别设置为所期望的系统。例如，如果你在 linux/amd64 平台上想查看 freebsd/arm 的文档，应将 $GOOS 设为 freebsd，$GOARCH 设为 arm。
// 
// 本包的主要用途是在提供更通用系统接口的其他包（如 "os"、"time" 和 "net"）内部使用。如果可以的话，建议使用这些包而非本包。
// 
// 有关本包中函数和数据类型的详细信息，请参阅相应操作系统的手册。
// 
// 这些调用通过返回 err == nil 表示成功；否则，err 代表描述失败情况的操作系统错误，并持有 syscall.Errno 类型的值。
package windows // 导入 "golang.org/x/sys/windows"

import (
	"bytes"
	"strings"
	"syscall"
	"unsafe"
)

// ByteSliceFromString 返回一个以NUL为终止符的字节切片，其中包含字符串s的文本。如果s在任何位置包含NUL字节，则返回(nil, syscall.EINVAL)。

// ff:
// s:
func ByteSliceFromString(s string) ([]byte, error) {
	if strings.IndexByte(s, 0) != -1 {
		return nil, syscall.EINVAL
	}
	a := make([]byte, len(s)+1)
	copy(a, s)
	return a, nil
}

// BytePtrFromString返回一个指向以NUL为终止符的字节数组的指针，该数组包含字符串s的文本内容。如果s在任何位置包含NUL字节，则返回(nil, syscall.EINVAL)。

// ff:
// s:
func BytePtrFromString(s string) (*byte, error) {
	a, err := ByteSliceFromString(s)
	if err != nil {
		return nil, err
	}
	return &a[0], nil
}

// ByteSliceToString 将切片 s 所表示的文本转换为字符串形式，其中去除终止空字符（NUL）及其后的所有字节。

// ff:
// s:
func ByteSliceToString(s []byte) string {
	if i := bytes.IndexByte(s, 0); i != -1 {
		s = s[:i]
	}
	return string(s)
}

// BytePtrToString 接收一个指向文本序列的指针，并返回相应的字符串。
// 若该指针为 nil，则返回空字符串。它假设文本序列以零字节结束；
// 若未出现零字节，程序可能会崩溃。

// ff:
// p:
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

// 单词形式的零，用于当我们需要一个指向0字节的有效指针时使用。
// 参见 mksyscall.pl。
var _zero uintptr

// 翻译提示:func  (ts  *时间戳)  Unix()  (秒  int64,  纳秒  int64)  {}

// ff:
// sec:
// nsec:
func (ts *Timespec) Unix() (sec int64, nsec int64) {
	return int64(ts.Sec), int64(ts.Nsec)
}

// 翻译提示:func  (tv  *时间戳)  Unix()  (秒  int64,  纳秒  int64)  {}

// ff:
// sec:
// nsec:
func (tv *Timeval) Unix() (sec int64, nsec int64) {
	return int64(tv.Sec), int64(tv.Usec) * 1000
}

// 翻译提示:func  (ts  *时间戳)  纳秒()  int64  {}

// ff:
func (ts *Timespec) Nano() int64 {
	return int64(ts.Sec)*1e9 + int64(ts.Nsec)
}

// 翻译提示:func  (tv  *时间戳)  获取纳秒()  int64  {}

// ff:
func (tv *Timeval) Nano() int64 {
	return int64(tv.Sec)*1e9 + int64(tv.Usec)*1000
}
