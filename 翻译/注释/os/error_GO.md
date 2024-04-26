
<原文开始>
// Timeout reports whether this error represents a timeout.
<原文结束>

# <翻译开始>
// Timeout 返回一个布尔值，表示这个错误是否代表了一个超时。
# <翻译结束>


<原文开始>
// NewSyscallError returns, as an error, a new SyscallError
// with the given system call name and error details.
// As a convenience, if err is nil, NewSyscallError returns nil.
<原文结束>

# <翻译开始>
// NewSyscallError 作为错误返回一个新的SyscallError，给定系统调用名称和错误详情。作为一个便利，如果err为nil，NewSyscallError将返回nil。
# <翻译结束>


<原文开始>
// IsExist returns a boolean indicating whether the error is known to report
// that a file or directory already exists. It is satisfied by ErrExist as
// well as some syscall errors.
//
// This function predates errors.Is. It only supports errors returned by
// the os package. New code should use errors.Is(err, fs.ErrExist).
<原文结束>

# <翻译开始>
// IsExist 返回一个布尔值，表示错误是否报告文件或目录已存在。它对ErrExist以及其他一些syscall错误也成立。
// 
// 这个函数早于errors.Is出现。它仅支持os包返回的错误。新代码应使用errors.Is(err, fs.ErrExist)。
# <翻译结束>


<原文开始>
// IsNotExist returns a boolean indicating whether the error is known to
// report that a file or directory does not exist. It is satisfied by
// ErrNotExist as well as some syscall errors.
//
// This function predates errors.Is. It only supports errors returned by
// the os package. New code should use errors.Is(err, fs.ErrNotExist).
<原文结束>

# <翻译开始>
// IsNotExist 返回一个布尔值，表示该错误是否已知为报告文件或目录不存在的错误。
// 它对 ErrNotExist 以及某些 syscall 错误均满足条件。
// 
// 该函数早于 errors.Is 函数。它仅支持 os 包返回的错误。
// 新代码应使用 errors.Is(err, fs.ErrNotExist)。
# <翻译结束>


<原文开始>
// IsPermission returns a boolean indicating whether the error is known to
// report that permission is denied. It is satisfied by ErrPermission as well
// as some syscall errors.
//
// This function predates errors.Is. It only supports errors returned by
// the os package. New code should use errors.Is(err, fs.ErrPermission).
<原文结束>

# <翻译开始>
// IsPermission 返回一个布尔值，表示错误是否已知表示权限被拒绝。这个函数对 ErrPermission 以及某些系统调用错误都成立。
//
// 这个函数早于 errors.Is 的存在。它只支持由 os 包返回的错误。新的代码应使用 errors.Is(err, fs.ErrPermission)。
# <翻译结束>


<原文开始>
// IsTimeout returns a boolean indicating whether the error is known
// to report that a timeout occurred.
//
// This function predates errors.Is, and the notion of whether an
// error indicates a timeout can be ambiguous. For example, the Unix
// error EWOULDBLOCK sometimes indicates a timeout and sometimes does not.
// New code should use errors.Is with a value appropriate to the call
// returning the error, such as os.ErrDeadlineExceeded.
<原文结束>

# <翻译开始>
// IsTimeout 返回一个布尔值，表示错误是否报告了超时。
// 
// 该函数早于 errors.Is 出现，错误是否表示超时的概念可能有歧义。例如，Unix 错误 EWOULDBLOCK 有时表示超时，有时则不表示。新代码应使用 errors.Is，并传入与返回错误的调用相关的适当值，如 os.ErrDeadlineExceeded。
# <翻译结束>


<原文开始>
// PathError records an error and the operation and file path that caused it.
<原文结束>

# <翻译开始>
// PathError 记录了一个错误，以及导致该错误的操作和文件路径。
# <翻译结束>


<原文开始>
// SyscallError records an error from a specific system call.
<原文结束>

# <翻译开始>
// SyscallError 记录来自特定系统调用的错误。
# <翻译结束>


<原文开始>
// Portable analogs of some common system call errors.
//
// Errors returned from this package may be tested against these errors
// with errors.Is.
<原文结束>

# <翻译开始>
// 可移植版本的某些常见系统调用错误。
//
// 从本包返回的错误可通过使用 errors.Is 与下列错误进行比较检测。
# <翻译结束>


<原文开始>
	// ErrInvalid indicates an invalid argument.
	// Methods on File will return this error when the receiver is nil.
<原文结束>

# <翻译开始>
// ErrInvalid 表示一个无效的参数。
// 当接收者为nil时，File类型的方法将返回此错误。
# <翻译结束>


<原文开始>
	//ErrNoDeadline       = errNoDeadline()       // "file type does not support deadline"
	//ErrDeadlineExceeded = errDeadlineExceeded() // "i/o timeout"
<原文结束>

# <翻译开始>
// ErrNoDeadline       = errNoDeadline()       // "文件类型不支持超时截止"
// ErrDeadlineExceeded = errDeadlineExceeded() // "I/O超时"
# <翻译结束>

