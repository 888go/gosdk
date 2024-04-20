
<原文开始>
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
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
# <翻译结束>


<原文开始>
// import "golang.org/x/sys/windows"
<原文结束>

# <翻译开始>
// import "golang.org/x/sys/windows"
# <翻译结束>


<原文开始>
// ByteSliceFromString returns a NUL-terminated slice of bytes
// containing the text of s. If s contains a NUL byte at any
// location, it returns (nil, syscall.EINVAL).
<原文结束>

# <翻译开始>
// ByteSliceFromString returns a NUL-terminated slice of bytes
// containing the text of s. If s contains a NUL byte at any
// location, it returns (nil, syscall.EINVAL).
# <翻译结束>


<原文开始>
// BytePtrFromString returns a pointer to a NUL-terminated array of
// bytes containing the text of s. If s contains a NUL byte at any
// location, it returns (nil, syscall.EINVAL).
<原文结束>

# <翻译开始>
// BytePtrFromString returns a pointer to a NUL-terminated array of
// bytes containing the text of s. If s contains a NUL byte at any
// location, it returns (nil, syscall.EINVAL).
# <翻译结束>


<原文开始>
// ByteSliceToString returns a string form of the text represented by the slice s, with a terminating NUL and any
// bytes after the NUL removed.
<原文结束>

# <翻译开始>
// ByteSliceToString returns a string form of the text represented by the slice s, with a terminating NUL and any
// bytes after the NUL removed.
# <翻译结束>


<原文开始>
// BytePtrToString takes a pointer to a sequence of text and returns the corresponding string.
// If the pointer is nil, it returns the empty string. It assumes that the text sequence is terminated
// at a zero byte; if the zero byte is not present, the program may crash.
<原文结束>

# <翻译开始>
// BytePtrToString takes a pointer to a sequence of text and returns the corresponding string.
// If the pointer is nil, it returns the empty string. It assumes that the text sequence is terminated
// at a zero byte; if the zero byte is not present, the program may crash.
# <翻译结束>


<原文开始>
// Single-word zero for use when we need a valid pointer to 0 bytes.
// See mksyscall.pl.
<原文结束>

# <翻译开始>
// Single-word zero for use when we need a valid pointer to 0 bytes.
// See mksyscall.pl.
# <翻译结束>

