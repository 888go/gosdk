package os

// Process 存储了通过 StartProcess 创建的进程的信息。
type Process struct { //md5:1853bd128399ed63025af6e19df5e5e9

}

// ProcAttr 保存了将应用于通过 StartProcess 启动的新进程的属性
type ProcAttr struct { //md5:8961e05ca2979cd28c54469a8c4cd56e

}

// Getpid 返回调用者（caller）的进程ID。
func Getpid() int { //md5:d89d2f571caee3ac85abcebdcb43f584

}

// Getppid 返回调用者父进程的进程ID。
func Getppid() int { //md5:892d9411173a6ac655ce98c9656eb0b1

}

// FindProcess 通过进程的 pid 查找正在运行的进程。
//
// 返回的 Process 可用于获取底层操作系统进程的信息。
//
// 在 Unix 系统上，FindProcess 总是成功并返回给定 pid 的 Process，无论该进程是否存在。要检查进程是否实际存在，可以查看 p.Signal(syscall.Signal(0)) 是否报告错误。
func FindProcess(pid int) (*Process, error) { //md5:551f3e9284260b7c7de683688f3efc84

}

// StartProcess 根据指定的名称（name）、参数（argv）和属性（attr），启动一个新的进程。argv 切片将在新进程中成为 os.Args，通常它会以程序名称开始。
// 
// 如果调用goroutine使用 runtime.LockOSThread 锁定了操作系统线程并修改了任何可继承的OS级线程状态（例如 Linux 或 Plan 9 命名空间），新进程将继承调用者的线程状态。
// 
// StartProcess 是一个低级别的接口。os/exec 包提供了更高级别的接口。
// 
// 如果出现错误，错误类型将是 *PathError。
func StartProcess(name string, argv []string, attr *ProcAttr) (*Process, error) { //md5:2da46d468b16dd888e099b0858527a73

}

// Release 释放与进程 p 相关的任何资源，使其未来无法使用。
// 只有在未调用 Wait 时才需要调用 Release。
func (p *Process) Release() error { //md5:5f7092ca971957477e615d1698c703f0

}

// Kill 会立即导致 Process 终止。Kill 不会等待 Process 实际退出。此操作仅终止 Process 本身，而不涉及其可能启动的任何其他进程。
func (p *Process) Kill() error { //md5:656c329c26c55e1d9285fe1b63cb0b85

}

// Wait等待Process退出，然后返回一个描述其状态的ProcessState和任何可能出现的错误。
// Wait释放与该Process相关的所有资源。
// 在大多数操作系统上，该Process必须是当前进程的子进程，否则将返回错误。
func (p *Process) Wait() (*ProcessState, error) { //md5:9d22de8522c835ac20c209841782f64b

}

// Signal 向 Process 发送信号。
// 在 Windows 上发送中断信号未实现。
func (p *Process) Signal(sig Signal) error { //md5:80ebd83f0fb68d0715e1bcafb55295c6

}

// UserTime返回退出进程及其子进程的用户CPU时间。
func (p *ProcessState) UserTime() time.Duration { //md5:f0649e88df15652ecd81676a0e4b4761

}

// SystemTime 返回已退出进程及其子进程的系统CPU时间。
func (p *ProcessState) SystemTime() time.Duration { //md5:f8cc6564058bcf74322e6ebd6a2e8a1f

}

// Exited报告程序是否已退出。
// 在Unix系统上，如果程序通过调用exit退出，则返回true，
// 但如果程序因信号而终止，则返回false。
func (p *ProcessState) Exited() bool { //md5:4fa12b1e90de311f6d8aaaf51ab6831a

}

// 成功报告程序是否成功退出，例如在Unix上退出状态为0。
func (p *ProcessState) Success() bool { //md5:e030f8dde5d9f14556a74e1f09908e85

}

// Sys returns system-dependent exit information about
// the process. Convert it to the appropriate underlying
// type, such as syscall.WaitStatus on Unix, to access its contents.
func (p *ProcessState) Sys() any { //md5:29c81811e323269d0e7953fc4fe5598b

}

// SysUsage 返回关于退出进程的系统依赖的资源使用信息。将其转换为适当的底层类型，如 Unix 中的 *syscall.Rusage，以访问其内容。（在 Unix 上，*syscall.Rusage 与 getrusage(2) 手册页中定义的 struct rusage 相匹配。）
func (p *ProcessState) SysUsage() any { //md5:34d570db1d021dcc42e0d59d9b97763d

}

// ProcessState 存储了由 Wait 方法报告的进程信息。
type ProcessState struct { //md5:47f04d3d87bb0ee5c5f2cc6218fd1203

}

// Pid 返回已退出进程的进程ID。
func (p *ProcessState) Pid() int { //md5:80d92fb5bb66295f960b60befb831a2d

}


func (p *ProcessState) String() string { //md5:fda18bc2a5938fb5acecd410279c9ca3

}

// ExitCode 返回退出的进程的退出代码，如果进程尚未退出或被信号终止，则返回 -1。
func (p *ProcessState) ExitCode() int { //md5:43b73d914bfbeeb7e4a194a8421f0a12

}