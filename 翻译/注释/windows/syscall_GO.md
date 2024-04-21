
<原文开始>
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// 版权所有 ? 2009 Go 作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 此许可证可在 LICENSE 文件中找到。
# <翻译结束>


<原文开始>
// Package windows contains an interface to the low-level operating system
// primitives. OS details vary depending on the underlying system, and
// by default, godoc will display the OS-specific documentation for the current
// system. If you want godoc to display syscall documentation for another
// system, set $GOOS and $GOARCH to the desired system. For example, if
// you want to view documentation for freebsd/arm on linux/amd64, set $GOOS
// to freebsd and $GOARCH to arm.
//
// The primary use of this package is inside other packages that provide a more
// portable interface to the system, such as "os", "time" and "net".  Use
// those packages rather than this one if you can.
//
// For details of the functions and data types in this package consult
// the manuals for the appropriate operating system.
//
// These calls return err == nil to indicate success; otherwise
// err represents an operating system error describing the failure and
// holds a value of type syscall.Errno.
<原文结束>

# <翻译开始>
// Package windows 提供了对底层操作系统原语的接口。具体的操作系统细节会因底层系统的不同而有所变化，且默认情况下，godoc 会显示当前系统的特定文档。如果你想让 godoc 显示其他系统的 syscall 文档，请将 $GOOS 和 $GOARCH 环境变量分别设置为所期望的系统。例如，如果你在 linux/amd64 平台上想查看 freebsd/arm 的文档，应将 $GOOS 设为 freebsd，$GOARCH 设为 arm。
// 
// 本包的主要用途是在提供更通用系统接口的其他包（如 "os"、"time" 和 "net"）内部使用。如果可以的话，建议使用这些包而非本包。
// 
// 有关本包中函数和数据类型的详细信息，请参阅相应操作系统的手册。
// 
// 这些调用通过返回 err == nil 表示成功；否则，err 代表描述失败情况的操作系统错误，并持有 syscall.Errno 类型的值。
# <翻译结束>


<原文开始>
// import "golang.org/x/sys/windows"
<原文结束>

# <翻译开始>
// 导入 "golang.org/x/sys/windows"
# <翻译结束>


<原文开始>
// ByteSliceFromString returns a NUL-terminated slice of bytes
// containing the text of s. If s contains a NUL byte at any
// location, it returns (nil, syscall.EINVAL).
<原文结束>

# <翻译开始>
// ByteSliceFromString 返回一个以NUL为终止符的字节切片，其中包含字符串s的文本。如果s在任何位置包含NUL字节，则返回(nil, syscall.EINVAL)。
# <翻译结束>


<原文开始>
// BytePtrFromString returns a pointer to a NUL-terminated array of
// bytes containing the text of s. If s contains a NUL byte at any
// location, it returns (nil, syscall.EINVAL).
<原文结束>

# <翻译开始>
// BytePtrFromString返回一个指向以NUL为终止符的字节数组的指针，该数组包含字符串s的文本内容。如果s在任何位置包含NUL字节，则返回(nil, syscall.EINVAL)。
# <翻译结束>


<原文开始>
// ByteSliceToString returns a string form of the text represented by the slice s, with a terminating NUL and any
// bytes after the NUL removed.
<原文结束>

# <翻译开始>
// ByteSliceToString 将切片 s 所表示的文本转换为字符串形式，其中去除终止空字符（NUL）及其后的所有字节。
# <翻译结束>


<原文开始>
// BytePtrToString takes a pointer to a sequence of text and returns the corresponding string.
// If the pointer is nil, it returns the empty string. It assumes that the text sequence is terminated
// at a zero byte; if the zero byte is not present, the program may crash.
<原文结束>

# <翻译开始>
// BytePtrToString 接收一个指向文本序列的指针，并返回相应的字符串。
// 若该指针为 nil，则返回空字符串。它假设文本序列以零字节结束；
// 若未出现零字节，程序可能会崩溃。
# <翻译结束>


<原文开始>
// Single-word zero for use when we need a valid pointer to 0 bytes.
// See mksyscall.pl.
<原文结束>

# <翻译开始>
// 单词形式的零，用于当我们需要一个指向0字节的有效指针时使用。
// 参见 mksyscall.pl。
# <翻译结束>

