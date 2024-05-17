package exec

import (
	"context"
	"errors"
	"io"
	"os/exec"
)

// Error is returned by LookPath when it fails to classify a file as an
// executable.
//type Error struct { //md5:ab612c623fddc332
//
//}

// ErrWaitDelay is returned by (*Cmd).Wait if the process exits with a
// successful status code but its output pipes are not closed before the
// command's WaitDelay expires.
var ErrWaitDelay = errors.New("exec: WaitDelay expired before I/O complete") //md5:a9b7691a47b44f05

// ErrDot indicates that a path lookup resolved to an executable
// in the current directory due to ‘.’ being in the path, either
// implicitly or explicitly. See the package documentation for details.
//
// Note that functions in this package do not return ErrDot directly.
// Code should use errors.Is(err, ErrDot), not err == ErrDot,
// to test whether a returned error err is due to this condition.
var ErrDot = errors.New("cannot run executable found relative to current directory") //md5:bcceec82a404a66c

//func (e *Error) Error() string { //md5:47e007c6da65537c1ce793159301e0d2
//	return e.F.Error()
//}
//
//func (e *Error) Unwrap() error { //md5:fe0424a855c24b4387d85783d12d42d2
//	return e.F.Unwrap()
//}

// Cmd represents an external command being prepared or run.
//
// A Cmd cannot be reused after calling its Run, Output or CombinedOutput
// methods.
type Cmd struct { //md5:cdf91edf21fccb61
	F exec.Cmd
}

// Command 返回Cmd结构以使用给定参数执行命名程序。
//
// 它只设置返回结构中的Path和Args。
//
// 若名称不包含路径分隔符，则命令使用LookPath将名称解析为完整的路径（如果可能）。否则，它直接使用名称作为路径。
//
// 返回的Cmd的Args字段由命令名后跟arg元素构成， 因此arg不应包含命令名本身. 例如, I设置命令("echo", "hello").
// Args[0] 始终是名称，而不是可能解析的路径。
//
// 在Windows上，进程以单个字符串的形式接收整个命令行，并执行自己的解析.
// 命令使用与使用CommandLineToArgvW的应用程序兼容的算法（这是最常见的方法）将Args组合并引用到命令行字符串中.
// 值得注意的例外是msiexec.exe和cmd.exe（以及所有批处理文件）， 它们具有不同的去激励算法。
// 在这些或其他类似情况下, 您可以自己引用，并在SysProcAttr.CmdLine中提供完整的命令行，将Args留空。
func Command(name string, arg ...string) *Cmd { //md5:2912017d0548a6188a67825cb0633059
	return &Cmd{
		F: *exec.Command(name, arg...),
	}
}

// CommandContext 类似于 Command，但包含了上下文信息。
//
// 提供的上下文用于在进程（通过调用 cmd.Cancel 或 os.Process.Kill）完成之前，
// 若上下文已结束，则中断该进程。
//
// CommandContext 将命令的 Cancel 函数设置为调用其 Process 的 Kill 方法，
// 并将其 WaitDelay 保持未设置状态。在启动命令之前，调用者可以通过修改这些字段来改变取消行为。
func CommandContext(ctx context.Context, name string, arg ...string) *Cmd { //md5:6482c2ac4f2da9383d162ac6194f5a52
	return &Cmd{
		F: *exec.CommandContext(ctx, name, arg...),
	}
}

// String 返回对 c 的人类可读的描述。
// 它仅用于调试目的。
// 特别地，它不适合用作 shell 的输入。
// String 的输出可能会在不同的 Go 版本之间有所变化。
func (c *Cmd) String() string { //md5:81f6a84e6ffe77c0a4beb3a6ebcb002f
	return c.F.String()
}

// Run 方法启动指定的命令并等待其完成。
//
// 若命令运行正常，无标准输入、输出和错误复制问题，并以零退出状态退出，则返回的错误为nil。
//
// 如果命令已启动但未成功完成，返回的错误将是*ExitError类型。对于其他情况，可能会返回其他类型的错误。
//
// 如果调用goroutine通过runtime.LockOSThread锁定操作系统线程并对任何可继承的操作系统级线程状态进行了修改（例如Linux或Plan 9命名空间），则新进程将继承调用者的线程状态。
func (c *Cmd) Run() error { //md5:3c28fe2808b00288feef65a1a9bcc0df
	return c.F.Run()
}

// Start 方法启动指定的命令，但不会等待其完成执行。
//
// 若 Start 方法成功返回，则会设置 c.Process 字段。
//
// 在成功调用 Start 方法后，必须调用 Wait 方法来释放相关的系统资源。
func (c *Cmd) Start() error { //md5:e4e960bd99b0815c5f1d6601ec662dc4
	return c.F.Start()
}

// An ExitError reports an unsuccessful exit by a command.
type ExitError struct { //md5:efd049c987f07aab
	F exec.ExitError
}

func (e *ExitError) Error() string { //md5:e7d8165bbebf499178bc1e1a3a045796
	return e.F.Error()
}

// Wait 等待命令退出，并等待向标准输入复制或从标准输出、标准错误复制的任何操作完成。
//
// 命令必须已通过 Start 方法启动。
//
// 如果命令成功运行，在复制标准输入、标准输出和标准错误时没有问题，并以零退出状态退出，则返回的错误为 nil。
//
// 如果命令无法运行或未成功完成，返回的错误类型为 *ExitError。对于 I/O 问题，可能会返回其他类型的错误。
//
// 如果 c.Stdin、c.Stdout 或 c.Stderr 中任何一个不是 *os.File 类型，Wait 还会等待相应进程的 I/O 循环完成复制操作。
//
// Wait 释放与 Cmd 关联的所有资源。
func (c *Cmd) Wait() error { //md5:c1cc5cd3a68d4db76704a0a00058acc4
	return c.F.Wait()
}

// Output 运行命令并返回其标准输出。
// 返回的任何错误通常会是 *ExitError 类型。
// 若 c.Stderr 为 nil，Output 会填充 ExitError.Stderr。
func (c *Cmd) Output() ([]byte, error) { //md5:6c0c99752f0c5a95d947693df1a82db1
	return c.F.Output()
}

// CombinedOutput 执行命令并返回其标准输出和标准错误的合并结果。
func (c *Cmd) CombinedOutput() ([]byte, error) { //md5:fa60c475e01331ed5a06057e7774b761
	return c.F.CombinedOutput()
}

// StdinPipe 返回一个管道，该管道将在命令启动时连接到命令的标准输入。
// 当Wait观察到命令退出后，该管道将自动关闭。
// 调用者只需调用Close即可强制管道提前关闭。
// 例如，如果运行的命令在标准输入关闭之前不会退出，则调用者必须关闭该管道。
func (c *Cmd) StdinPipe() (io.WriteCloser, error) { //md5:a46d29cb606f6af4122775095f43f1d6
	return c.F.StdinPipe()
}

// StdoutPipe 返回一个管道，当命令启动时，该管道将与命令的标准输出相连。
//
// 当Wait观察到命令退出后，会关闭该管道，因此大多数调用者无需自行关闭管道。因此，在完成从管道的所有读取操作之前调用Wait是不正确的。出于同样的原因，在使用StdoutPipe时调用Run也是不正确的。有关惯用用法，请参阅示例。
func (c *Cmd) StdoutPipe() (io.ReadCloser, error) { //md5:ddeb49a7d48e93da480da5299d2eed2b
	return c.F.StdoutPipe()
}

// StderrPipe 返回一个管道，当命令启动时，该管道将与命令的标准错误相连。
//
// 当Wait检测到命令退出后，会关闭该管道，因此大多数调用者无需自行关闭管道。因此，在完成从管道的所有读取操作之前调用Wait是不正确的。出于同样的原因，当使用StderrPipe时，调用Run也是不正确的。有关惯用用法，请参阅StdoutPipe示例。
func (c *Cmd) StderrPipe() (io.ReadCloser, error) { //md5:9b10883945ef55930f9026118bde260b
	return c.F.StderrPipe()
}

// Environ 返回当前配置下，命令将要运行的环境副本。
func (c *Cmd) Environ() []string { //md5:891b61dc22bc81880c267878539c842b
	return c.F.Environ()
}
