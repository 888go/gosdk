
<原文开始>
// CommandContext is like Command but includes a context.
//
// The provided context is used to interrupt the process
// (by calling cmd.Cancel or os.Process.Kill)
// if the context becomes done before the command completes on its own.
//
// CommandContext sets the command's Cancel function to invoke the Kill method
// on its Process, and leaves its WaitDelay unset. The caller may change the
// cancellation behavior by modifying those fields before starting the command.
<原文结束>

# <翻译开始>
// CommandContext 类似于 Command，但包含了上下文信息。
//
// 提供的上下文用于在进程（通过调用 cmd.Cancel 或 os.Process.Kill）完成之前，
// 若上下文已结束，则中断该进程。
//
// CommandContext 将命令的 Cancel 函数设置为调用其 Process 的 Kill 方法，
// 并将其 WaitDelay 保持未设置状态。在启动命令之前，调用者可以通过修改这些字段来改变取消行为。
# <翻译结束>


<原文开始>
// String returns a human-readable description of c.
// It is intended only for debugging.
// In particular, it is not suitable for use as input to a shell.
// The output of String may vary across Go releases.
<原文结束>

# <翻译开始>
// String 返回对 c 的人类可读的描述。
// 它仅用于调试目的。
// 特别地，它不适合用作 shell 的输入。
// String 的输出可能会在不同的 Go 版本之间有所变化。
# <翻译结束>


<原文开始>
// Run starts the specified command and waits for it to complete.
//
// The returned error is nil if the command runs, has no problems
// copying stdin, stdout, and stderr, and exits with a zero exit
// status.
//
// If the command starts but does not complete successfully, the error is of
// type *ExitError. Other error types may be returned for other situations.
//
// If the calling goroutine has locked the operating system thread
// with runtime.LockOSThread and modified any inheritable OS-level
// thread state (for example, Linux or Plan 9 name spaces), the new
// process will inherit the caller's thread state.
<原文结束>

# <翻译开始>
// Run 方法启动指定的命令并等待其完成。
//
// 若命令运行正常，无标准输入、输出和错误复制问题，并以零退出状态退出，则返回的错误为nil。
//
// 如果命令已启动但未成功完成，返回的错误将是*ExitError类型。对于其他情况，可能会返回其他类型的错误。
//
// 如果调用goroutine通过runtime.LockOSThread锁定操作系统线程并对任何可继承的操作系统级线程状态进行了修改（例如Linux或Plan 9命名空间），则新进程将继承调用者的线程状态。
# <翻译结束>


<原文开始>
// Start starts the specified command but does not wait for it to complete.
//
// If Start returns successfully, the c.Process field will be set.
//
// After a successful call to Start the Wait method must be called in
// order to release associated system resources.
<原文结束>

# <翻译开始>
// Start 方法启动指定的命令，但不会等待其完成执行。
//
// 若 Start 方法成功返回，则会设置 c.Process 字段。
//
// 在成功调用 Start 方法后，必须调用 Wait 方法来释放相关的系统资源。
# <翻译结束>


<原文开始>
// Wait waits for the command to exit and waits for any copying to
// stdin or copying from stdout or stderr to complete.
//
// The command must have been started by Start.
//
// The returned error is nil if the command runs, has no problems
// copying stdin, stdout, and stderr, and exits with a zero exit
// status.
//
// If the command fails to run or doesn't complete successfully, the
// error is of type *ExitError. Other error types may be
// returned for I/O problems.
//
// If any of c.Stdin, c.Stdout or c.Stderr are not an *os.File, Wait also waits
// for the respective I/O loop copying to or from the process to complete.
//
// Wait releases any resources associated with the Cmd.
<原文结束>

# <翻译开始>
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
# <翻译结束>


<原文开始>
// Output runs the command and returns its standard output.
// Any returned error will usually be of type *ExitError.
// If c.Stderr was nil, Output populates ExitError.Stderr.
<原文结束>

# <翻译开始>
// Output 运行命令并返回其标准输出。
// 返回的任何错误通常会是 *ExitError 类型。
// 若 c.Stderr 为 nil，Output 会填充 ExitError.Stderr。
# <翻译结束>


<原文开始>
// CombinedOutput runs the command and returns its combined standard
// output and standard error.
<原文结束>

# <翻译开始>
// CombinedOutput 执行命令并返回其标准输出和标准错误的合并结果。
# <翻译结束>


<原文开始>
// StdinPipe returns a pipe that will be connected to the command's
// standard input when the command starts.
// The pipe will be closed automatically after Wait sees the command exit.
// A caller need only call Close to force the pipe to close sooner.
// For example, if the command being run will not exit until standard input
// is closed, the caller must close the pipe.
<原文结束>

# <翻译开始>
// StdinPipe 返回一个管道，该管道将在命令启动时连接到命令的标准输入。
// 当Wait观察到命令退出后，该管道将自动关闭。
// 调用者只需调用Close即可强制管道提前关闭。
// 例如，如果运行的命令在标准输入关闭之前不会退出，则调用者必须关闭该管道。
# <翻译结束>


<原文开始>
// StdoutPipe returns a pipe that will be connected to the command's
// standard output when the command starts.
//
// Wait will close the pipe after seeing the command exit, so most callers
// need not close the pipe themselves. It is thus incorrect to call Wait
// before all reads from the pipe have completed.
// For the same reason, it is incorrect to call Run when using StdoutPipe.
// See the example for idiomatic usage.
<原文结束>

# <翻译开始>
// StdoutPipe 返回一个管道，当命令启动时，该管道将与命令的标准输出相连。
//
// 当Wait观察到命令退出后，会关闭该管道，因此大多数调用者无需自行关闭管道。因此，在完成从管道的所有读取操作之前调用Wait是不正确的。出于同样的原因，在使用StdoutPipe时调用Run也是不正确的。有关惯用用法，请参阅示例。
# <翻译结束>


<原文开始>
// StderrPipe returns a pipe that will be connected to the command's
// standard error when the command starts.
//
// Wait will close the pipe after seeing the command exit, so most callers
// need not close the pipe themselves. It is thus incorrect to call Wait
// before all reads from the pipe have completed.
// For the same reason, it is incorrect to use Run when using StderrPipe.
// See the StdoutPipe example for idiomatic usage.
<原文结束>

# <翻译开始>
// StderrPipe 返回一个管道，当命令启动时，该管道将与命令的标准错误相连。
//
// 当Wait检测到命令退出后，会关闭该管道，因此大多数调用者无需自行关闭管道。因此，在完成从管道的所有读取操作之前调用Wait是不正确的。出于同样的原因，当使用StderrPipe时，调用Run也是不正确的。有关惯用用法，请参阅StdoutPipe示例。
# <翻译结束>


<原文开始>
// Environ returns a copy of the environment in which the command would be run
// as it is currently configured.
<原文结束>

# <翻译开始>
// Environ 返回当前配置下，命令将要运行的环境副本。
# <翻译结束>


<原文开始>
// ErrWaitDelay is returned by (*Cmd).Wait if the process exits with a
// successful status code but its output pipes are not closed before the
// command's WaitDelay expires.
<原文结束>

# <翻译开始>
// ErrWaitDelay 是在以下情况下由 (*Cmd).Wait 返回的错误：如果进程以成功状态退出，但在命令的 WaitDelay 超时之前其输出管道尚未关闭。
# <翻译结束>

