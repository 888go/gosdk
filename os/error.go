package os

import (
	"io/fs"
	"os"
)

// 可移植版本的某些常见系统调用错误。
//
// 从本包返回的错误可通过使用 errors.Is 与下列错误进行比较检测。
var (
// ErrInvalid 表示一个无效的参数。
// 当接收者为nil时，File类型的方法将返回此错误。
	ErrInvalid          = os.ErrInvalid //hs:常量_错误_无效参数        // "invalid argument"
	ErrPermission       = os.ErrPermission //hs:常量_错误_权限拒绝      // "permission denied"
	ErrExist            = os.ErrExist //hs:常量_错误_文件已存在      // "file already exists"
	ErrNotExist         = os.ErrNotExist //hs:常量_错误_文件不存在        // "file does not exist"
	ErrClosed           = os.ErrClosed //hs:常量_错误_文件已关闭        // "file already closed"
	ErrNoDeadline       = os.ErrNoDeadline //hs:常量_错误_文件类型不支持超时截止         // "文件类型不支持超时截止"
	ErrDeadlineExceeded = os.ErrDeadlineExceeded //hs:常量_错误_IO超时      // "I/O超时"
)

// PathError 记录了一个错误，以及导致该错误的操作和文件路径。
type PathError = fs.PathError //md5:e10fdcf70552a8f7904823850e701b25

// SyscallError 记录来自特定系统调用的错误。
type SyscallError struct { //md5:629fa3dd95978e22325ca4574ef483c6
	F *os.SyscallError
}


// ff:取错误文本
func (e *SyscallError) Error() string { //md5:12b20881dc68d5f1c98411209aca4a3a
	return e.F.Error()
}


// ff:取错误对象
func (e *SyscallError) Unwrap() error { //md5:383c868a76a65b16135dc3d1b98bdd75
	return e.Unwrap()
}

// Timeout 返回一个布尔值，表示这个错误是否代表了一个超时。

// ff:是否超时
func (e *SyscallError) Timeout() bool { //md5:d66cf10e3391c84ee831473ad6a2e5cb
	return e.F.Timeout()
}

// NewSyscallError 作为错误返回一个新的SyscallError，给定系统调用名称和错误详情。作为一个便利，如果err为nil，NewSyscallError将返回nil。

// ff:创建调用错误
// err:错误
// syscall:错误名称
func NewSyscallError(syscall string, err error) error { //md5:55203c659aec8ace18a96610a164b56b
	return os.NewSyscallError(syscall, err)
}

// IsExist 返回一个布尔值，表示错误是否报告文件或目录已存在。它对ErrExist以及其他一些syscall错误也成立。
//
// 这个函数早于errors.Is出现。它仅支持os包返回的错误。新代码应使用errors.Is(err, fs.ErrExist)。

// ff:是否文件已存在
// err:错误
func IsExist(err error) bool { //md5:635d346e3c48a1334f8f8a0d9813e4fd
	return os.IsExist(err)
}

// IsNotExist 返回一个布尔值，表示该错误是否已知为报告文件或目录不存在的错误。
// 它对 ErrNotExist 以及某些 syscall 错误均满足条件。
//
// 该函数早于 errors.Is 函数。它仅支持 os 包返回的错误。
// 新代码应使用 errors.Is(err, fs.ErrNotExist)。

// ff:是否文件不存在
// err:错误
func IsNotExist(err error) bool { //md5:3f35bc364c0b48c2bf65b8dae188cb64
	return os.IsNotExist(err)
}

// IsPermission 返回一个布尔值，表示错误是否已知表示权限被拒绝。这个函数对 ErrPermission 以及某些系统调用错误都成立。
//
// 这个函数早于 errors.Is 的存在。它只支持由 os 包返回的错误。新的代码应使用 errors.Is(err, fs.ErrPermission)。

// ff:是否权限拒绝
// err:错误
func IsPermission(err error) bool { //md5:d6c32cc7ada9abb36b6bdf32d5c80248
	return os.IsPermission(err)
}

// IsTimeout 返回一个布尔值，表示错误是否报告了超时。
//
// 该函数早于 errors.Is 出现，错误是否表示超时的概念可能有歧义。例如，Unix 错误 EWOULDBLOCK 有时表示超时，有时则不表示。新代码应使用 errors.Is，并传入与返回错误的调用相关的适当值，如 os.ErrDeadlineExceeded。

// ff:是否超时错误
// err:错误
func IsTimeout(err error) bool { //md5:a8da64488bb9e41dbb6285d2e2afece1
	return os.IsTimeout(err)
}
