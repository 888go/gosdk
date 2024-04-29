
<原文开始>
// Getpid returns the process id of the caller.
<原文结束>

# <翻译开始>
// Getpid 返回调用者（caller）的进程ID。
# <翻译结束>


<原文开始>
// Getppid returns the process id of the caller's parent.
<原文结束>

# <翻译开始>
// Getppid 返回调用者父进程的进程ID。
# <翻译结束>


<原文开始>
// FindProcess looks for a running process by its pid.
//
// The Process it returns can be used to obtain information
// about the underlying operating system process.
//
// On Unix systems, FindProcess always succeeds and returns a Process
// for the given pid, regardless of whether the process exists. To test whether
// the process actually exists, see whether p.Signal(syscall.Signal(0)) reports
// an error.
<原文结束>

# <翻译开始>
// FindProcess 通过进程的 pid 查找正在运行的进程。
//
// 返回的 Process 可用于获取底层操作系统进程的信息。
//
// 在 Unix 系统上，FindProcess 总是成功并返回给定 pid 的 Process，无论该进程是否存在。要检查进程是否实际存在，可以查看 p.Signal(syscall.Signal(0)) 是否报告错误。
# <翻译结束>


<原文开始>
// StartProcess starts a new process with the program, arguments and attributes
// specified by name, argv and attr. The argv slice will become os.Args in the
// new process, so it normally starts with the program name.
//
// If the calling goroutine has locked the operating system thread
// with runtime.LockOSThread and modified any inheritable OS-level
// thread state (for example, Linux or Plan 9 name spaces), the new
// process will inherit the caller's thread state.
//
// StartProcess is a low-level interface. The os/exec package provides
// higher-level interfaces.
//
// If there is an error, it will be of type *PathError.
<原文结束>

# <翻译开始>
// StartProcess 根据指定的名称（name）、参数（argv）和属性（attr），启动一个新的进程。argv 切片将在新进程中成为 os.Args，通常它会以程序名称开始。
// 
// 如果调用goroutine使用 runtime.LockOSThread 锁定了操作系统线程并修改了任何可继承的OS级线程状态（例如 Linux 或 Plan 9 命名空间），新进程将继承调用者的线程状态。
// 
// StartProcess 是一个低级别的接口。os/exec 包提供了更高级别的接口。
// 
// 如果出现错误，错误类型将是 *PathError。
# <翻译结束>


<原文开始>
// Release releases any resources associated with the Process p,
// rendering it unusable in the future.
// Release only needs to be called if Wait is not.
<原文结束>

# <翻译开始>
// Release 释放与进程 p 相关的任何资源，使其未来无法使用。
// 只有在未调用 Wait 时才需要调用 Release。
# <翻译结束>


<原文开始>
// Kill causes the Process to exit immediately. Kill does not wait until
// the Process has actually exited. This only kills the Process itself,
// not any other processes it may have started.
<原文结束>

# <翻译开始>
// Kill 会立即导致 Process 终止。Kill 不会等待 Process 实际退出。此操作仅终止 Process 本身，而不涉及其可能启动的任何其他进程。
# <翻译结束>


<原文开始>
// Wait waits for the Process to exit, and then returns a
// ProcessState describing its status and an error, if any.
// Wait releases any resources associated with the Process.
// On most operating systems, the Process must be a child
// of the current process or an error will be returned.
<原文结束>

# <翻译开始>
// Wait等待Process退出，然后返回一个描述其状态的ProcessState和任何可能出现的错误。
// Wait释放与该Process相关的所有资源。
// 在大多数操作系统上，该Process必须是当前进程的子进程，否则将返回错误。
# <翻译结束>


<原文开始>
// Signal sends a signal to the Process.
// Sending Interrupt on Windows is not implemented.
<原文结束>

# <翻译开始>
// Signal 向 Process 发送信号。
// 在 Windows 上发送中断信号未实现。
# <翻译结束>


<原文开始>
// UserTime returns the user CPU time of the exited process and its children.
<原文结束>

# <翻译开始>
// UserTime返回退出进程及其子进程的用户CPU时间。
# <翻译结束>


<原文开始>
// SystemTime returns the system CPU time of the exited process and its children.
<原文结束>

# <翻译开始>
// SystemTime 返回已退出进程及其子进程的系统CPU时间。
# <翻译结束>


<原文开始>
// Exited reports whether the program has exited.
// On Unix systems this reports true if the program exited due to calling exit,
// but false if the program terminated due to a signal.
<原文结束>

# <翻译开始>
// Exited报告程序是否已退出。
// 在Unix系统上，如果程序通过调用exit退出，则返回true，
// 但如果程序因信号而终止，则返回false。
# <翻译结束>


<原文开始>
// Success reports whether the program exited successfully,
// such as with exit status 0 on Unix.
<原文结束>

# <翻译开始>
// 成功报告程序是否成功退出，例如在Unix上退出状态为0。
# <翻译结束>


<原文开始>
// SysUsage returns system-dependent resource usage information about
// the exited process. Convert it to the appropriate underlying
// type, such as *syscall.Rusage on Unix, to access its contents.
// (On Unix, *syscall.Rusage matches struct rusage as defined in the
// getrusage(2) manual page.)
<原文结束>

# <翻译开始>
// SysUsage 返回关于退出进程的系统依赖的资源使用信息。将其转换为适当的底层类型，如 Unix 中的 *syscall.Rusage，以访问其内容。（在 Unix 上，*syscall.Rusage 与 getrusage(2) 手册页中定义的 struct rusage 相匹配。）
# <翻译结束>


<原文开始>
// Pid returns the process id of the exited process.
<原文结束>

# <翻译开始>
// Pid 返回已退出进程的进程ID。
# <翻译结束>


<原文开始>
// ExitCode returns the exit code of the exited process, or -1
// if the process hasn't exited or was terminated by a signal.
<原文结束>

# <翻译开始>
// ExitCode 返回退出的进程的退出代码，如果进程尚未退出或被信号终止，则返回 -1。
# <翻译结束>


<原文开始>
// Process stores the information about a process created by StartProcess.
<原文结束>

# <翻译开始>
// Process 存储了通过 StartProcess 创建的进程的信息。
# <翻译结束>


<原文开始>
// ProcAttr holds the attributes that will be applied to a new process
// started by StartProcess.
<原文结束>

# <翻译开始>
// ProcAttr 保存了将应用于通过 StartProcess 启动的新进程的属性
# <翻译结束>


<原文开始>
// ProcessState stores information about a process, as reported by Wait.
<原文结束>

# <翻译开始>
// ProcessState 存储了由 Wait 方法报告的进程信息。
# <翻译结束>


<原文开始>
// ErrProcessDone indicates a Process has finished.
<原文结束>

# <翻译开始>
// ErrProcessDone 表示一个进程已经完成。. md5:1900bc3b9ff279fd
# <翻译结束>

