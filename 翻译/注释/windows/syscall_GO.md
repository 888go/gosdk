
<原文开始>
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// 版权所有 ? 2009 Go作者。保留所有权利。
// 本源代码的使用受BSD风格
// 许可证约束，该许可证可在LICENSE文件中找到。
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
// Package windows 提供对底层操作系统原语的接口。具体操作系统的细节因底层系统不同而异，且默认情况下，godoc 将显示当前系统的特定文档。如果你想让 godoc 显示其他系统的系统调用文档，请将 $GOOS 和 $GOARCH 环境变量分别设置为所期望的系统和架构。例如，如果你想在 linux/amd64 系统上查看 freebsd/arm 的文档，应将 $GOOS 设为 freebsd，$GOARCH 设为 arm。
// 
// 本包主要用途在于为其他提供更通用系统接口的包（如 "os"、"time" 和 "net"）内部使用。如果可以，请优先使用这些包而非本包。
// 
// 关于本包中函数和数据类型的详细信息，请参考对应操作系统的手册。
// 
// 这些函数通过返回 err == nil 表示成功；否则，err 代表描述失败情况的操作系统错误，并持有 syscall.Errno 类型的值。
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
// ByteSliceFromString 返回一个以NUL字节结尾的、包含字符串s文本内容的字节切片。
// 若s在任意位置包含NUL字节，将返回(nil, syscall.EINVAL)。
# <翻译结束>


<原文开始>
// BytePtrFromString returns a pointer to a NUL-terminated array of
// bytes containing the text of s. If s contains a NUL byte at any
// location, it returns (nil, syscall.EINVAL).
<原文结束>

# <翻译开始>
// BytePtrFromString 返回一个指向以 NUL 结尾的字节数组的指针，该数组包含字符串 s 的文本。如果 s 在任何位置包含 NUL 字节，则返回 (nil, syscall.EINVAL)。
# <翻译结束>


<原文开始>
// ByteSliceToString returns a string form of the text represented by the slice s, with a terminating NUL and any
// bytes after the NUL removed.
<原文结束>

# <翻译开始>
// ByteSliceToString 将切片 s 表示的文本转换为字符串形式，删除终止空字符（NUL）及其后的所有字节。
# <翻译结束>


<原文开始>
// BytePtrToString takes a pointer to a sequence of text and returns the corresponding string.
// If the pointer is nil, it returns the empty string. It assumes that the text sequence is terminated
// at a zero byte; if the zero byte is not present, the program may crash.
<原文结束>

# <翻译开始>
// BytePtrToString 接收一个指向文本序列的指针，并返回相应的字符串。
// 若该指针为nil，则返回空字符串。其假设文本序列以零字节结束；
// 若未出现零字节，程序可能会崩溃。
# <翻译结束>


<原文开始>
// Single-word zero for use when we need a valid pointer to 0 bytes.
// See mksyscall.pl.
<原文结束>

# <翻译开始>
// 单词形式的零，用于当我们需要一个指向0字节的有效指针时。
// 参见 mksyscall.pl。
# <翻译结束>

