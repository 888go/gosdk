package os

// PathError 记录了一个错误，以及导致该错误的操作和文件路径。
type PathError = fs.PathError //md5:e10fdcf70552a8f7904823850e701b25



// SyscallError 记录来自特定系统调用的错误。
type SyscallError struct { //md5:629fa3dd95978e22325ca4574ef483c6

}


func (e *SyscallError) Error() string { //md5:12b20881dc68d5f1c98411209aca4a3a

}


func (e *SyscallError) Unwrap() error { //md5:383c868a76a65b16135dc3d1b98bdd75

}

// Timeout 返回一个布尔值，表示这个错误是否代表了一个超时。
func (e *SyscallError) Timeout() bool { //md5:d66cf10e3391c84ee831473ad6a2e5cb

}

// NewSyscallError 作为错误返回一个新的SyscallError，给定系统调用名称和错误详情。作为一个便利，如果err为nil，NewSyscallError将返回nil。
func NewSyscallError(syscall string, err error) error { //md5:55203c659aec8ace18a96610a164b56b

}

// IsExist 返回一个布尔值，表示错误是否报告文件或目录已存在。它对ErrExist以及其他一些syscall错误也成立。
// 
// 这个函数早于errors.Is出现。它仅支持os包返回的错误。新代码应使用errors.Is(err, fs.ErrExist)。
func IsExist(err error) bool { //md5:635d346e3c48a1334f8f8a0d9813e4fd

}

// IsNotExist 返回一个布尔值，表示该错误是否已知为报告文件或目录不存在的错误。
// 它对 ErrNotExist 以及某些 syscall 错误均满足条件。
// 
// 该函数早于 errors.Is 函数。它仅支持 os 包返回的错误。
// 新代码应使用 errors.Is(err, fs.ErrNotExist)。
func IsNotExist(err error) bool { //md5:3f35bc364c0b48c2bf65b8dae188cb64

}

// IsPermission 返回一个布尔值，表示错误是否已知表示权限被拒绝。这个函数对 ErrPermission 以及某些系统调用错误都成立。
//
// 这个函数早于 errors.Is 的存在。它只支持由 os 包返回的错误。新的代码应使用 errors.Is(err, fs.ErrPermission)。
func IsPermission(err error) bool { //md5:d6c32cc7ada9abb36b6bdf32d5c80248

}

// IsTimeout 返回一个布尔值，表示错误是否报告了超时。
// 
// 该函数早于 errors.Is 出现，错误是否表示超时的概念可能有歧义。例如，Unix 错误 EWOULDBLOCK 有时表示超时，有时则不表示。新代码应使用 errors.Is，并传入与返回错误的调用相关的适当值，如 os.ErrDeadlineExceeded。
func IsTimeout(err error) bool { //md5:a8da64488bb9e41dbb6285d2e2afece1

}