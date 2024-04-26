
<原文开始>
// Getuid returns the numeric user id of the caller.
//
// On Windows, it returns -1.
<原文结束>

# <翻译开始>
// Getuid 返回调用者的数字用户ID。
//
// 在Windows上，它返回-1。
# <翻译结束>


<原文开始>
// Geteuid returns the numeric effective user id of the caller.
//
// On Windows, it returns -1.
<原文结束>

# <翻译开始>
// Geteuid 返回调用者的有效用户 ID 号。
// 
// 在 Windows 系统中，它返回 -1。
# <翻译结束>


<原文开始>
// Getgid returns the numeric group id of the caller.
//
// On Windows, it returns -1.
<原文结束>

# <翻译开始>
// Getgid 返回调用者所属的组的数字标识符。
//
// 在 Windows 系统中，它返回 -1。
# <翻译结束>


<原文开始>
// Getegid returns the numeric effective group id of the caller.
//
// On Windows, it returns -1.
<原文结束>

# <翻译开始>
// Getegid 返回调用者当前有效的数字组ID。
//
// 在Windows系统上，它返回-1。
# <翻译结束>


<原文开始>
// Getgroups returns a list of the numeric ids of groups that the caller belongs to.
//
// On Windows, it returns syscall.EWINDOWS. See the os/user package
// for a possible alternative.
<原文结束>

# <翻译开始>
// Getgroups 返回一个列表，包含调用者所属的组的数字ID。
//
// 在Windows系统上，它会返回syscall.EWINDOWS错误。可以查看os/user包以寻求可能的替代方法。
# <翻译结束>


<原文开始>
// Exit causes the current program to exit with the given status code.
// Conventionally, code zero indicates success, non-zero an error.
// The program terminates immediately; deferred functions are not run.
//
// For portability, the status code should be in the range [0, 125].
<原文结束>

# <翻译开始>
// Exit 会导致当前程序使用给定的状态码退出。
// 通常，状态码为零表示成功，非零表示错误。
// 程序会立即终止；延迟执行的函数不会运行。
// 
// 为了兼容性，状态码应该在[0, 125]范围内。
# <翻译结束>


<原文开始>
// Args hold the command-line arguments, starting with the program name.
<原文结束>

# <翻译开始>
// Args 保存了命令行参数，从程序名称开始。
# <翻译结束>

