package exec


func (e *Error) Error() string { //md5:47e007c6da65537c1ce793159301e0d2

}


func (e *Error) Unwrap() error { //md5:fe0424a855c24b4387d85783d12d42d2

}


func (w wrappedError) Error() string { //md5:299c35975f0c83bd3edb8bb7999594a2

}


func (w wrappedError) Unwrap() error { //md5:688743fbe0cc4b9c4cd1c9455395791b

}

// Command returns the Cmd struct to execute the named program with
// the given arguments.
//
// It sets only the Path and Args in the returned structure.
//
// If name contains no path separators, Command uses LookPath to
// resolve name to a complete path if possible. Otherwise it uses name
// directly as Path.
//
// The returned Cmd's Args field is constructed from the command name
// followed by the elements of arg, so arg should not include the
// command name itself. For example, Command("echo", "hello").
// Args[0] is always name, not the possibly resolved Path.
//
// On Windows, processes receive the whole command line as a single string
// and do their own parsing. Command combines and quotes Args into a command
// line string with an algorithm compatible with applications using
// CommandLineToArgvW (which is the most common way). Notable exceptions are
// msiexec.exe and cmd.exe (and thus, all batch files), which have a different
// unquoting algorithm. In these or other similar cases, you can do the
// quoting yourself and provide the full command line in SysProcAttr.CmdLine,
// leaving Args empty.
func Command(name string, arg ...string) *Cmd { //md5:2912017d0548a6188a67825cb0633059

}

// CommandContext 类似于 Command，但包含了上下文信息。
//
// 提供的上下文用于在进程（通过调用 cmd.Cancel 或 os.Process.Kill）完成之前，
// 若上下文已结束，则中断该进程。
//
// CommandContext 将命令的 Cancel 函数设置为调用其 Process 的 Kill 方法，
// 并将其 WaitDelay 保持未设置状态。在启动命令之前，调用者可以通过修改这些字段来改变取消行为。
func CommandContext(ctx context.Context, name string, arg ...string) *Cmd { //md5:6482c2ac4f2da9383d162ac6194f5a52

}

// String 返回对 c 的人类可读的描述。
// 它仅用于调试目的。
// 特别地，它不适合用作 shell 的输入。
// String 的输出可能会在不同的 Go 版本之间有所变化。
func (c *Cmd) String() string { //md5:81f6a84e6ffe77c0a4beb3a6ebcb002f

}

// Run 方法启动指定的命令并等待其完成。
//
// 若命令运行正常，无标准输入、输出和错误复制问题，并以零退出状态退出，则返回的错误为nil。
//
// 如果命令已启动但未成功完成，返回的错误将是*ExitError类型。对于其他情况，可能会返回其他类型的错误。
//
// 如果调用goroutine通过runtime.LockOSThread锁定操作系统线程并对任何可继承的操作系统级线程状态进行了修改（例如Linux或Plan 9命名空间），则新进程将继承调用者的线程状态。
func (c *Cmd) Run() error { //md5:3c28fe2808b00288feef65a1a9bcc0df

}

// Start 方法启动指定的命令，但不会等待其完成执行。
//
// 若 Start 方法成功返回，则会设置 c.Process 字段。
//
// 在成功调用 Start 方法后，必须调用 Wait 方法来释放相关的系统资源。
func (c *Cmd) Start() error { //md5:e4e960bd99b0815c5f1d6601ec662dc4

}


func (e *ExitError) Error() string { //md5:e7d8165bbebf499178bc1e1a3a045796

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

}

// Output 运行命令并返回其标准输出。
// 返回的任何错误通常会是 *ExitError 类型。
// 若 c.Stderr 为 nil，Output 会填充 ExitError.Stderr。
func (c *Cmd) Output() ([]byte, error) { //md5:6c0c99752f0c5a95d947693df1a82db1

}

// CombinedOutput 执行命令并返回其标准输出和标准错误的合并结果。
func (c *Cmd) CombinedOutput() ([]byte, error) { //md5:fa60c475e01331ed5a06057e7774b761

}

// StdinPipe 返回一个管道，该管道将在命令启动时连接到命令的标准输入。
// 当Wait观察到命令退出后，该管道将自动关闭。
// 调用者只需调用Close即可强制管道提前关闭。
// 例如，如果运行的命令在标准输入关闭之前不会退出，则调用者必须关闭该管道。
func (c *Cmd) StdinPipe() (io.WriteCloser, error) { //md5:a46d29cb606f6af4122775095f43f1d6

}

// StdoutPipe 返回一个管道，当命令启动时，该管道将与命令的标准输出相连。
//
// 当Wait观察到命令退出后，会关闭该管道，因此大多数调用者无需自行关闭管道。因此，在完成从管道的所有读取操作之前调用Wait是不正确的。出于同样的原因，在使用StdoutPipe时调用Run也是不正确的。有关惯用用法，请参阅示例。
func (c *Cmd) StdoutPipe() (io.ReadCloser, error) { //md5:ddeb49a7d48e93da480da5299d2eed2b

}

// StderrPipe 返回一个管道，当命令启动时，该管道将与命令的标准错误相连。
//
// 当Wait检测到命令退出后，会关闭该管道，因此大多数调用者无需自行关闭管道。因此，在完成从管道的所有读取操作之前调用Wait是不正确的。出于同样的原因，当使用StderrPipe时，调用Run也是不正确的。有关惯用用法，请参阅StdoutPipe示例。
func (c *Cmd) StderrPipe() (io.ReadCloser, error) { //md5:9b10883945ef55930f9026118bde260b

}


func (w *prefixSuffixSaver) Write(p []byte) (n int, err error) { //md5:b96f3a5712d84fa3fbc26437e496280b

}


func (w *prefixSuffixSaver) Bytes() []byte { //md5:66c73a4889d14bec0dceec6067acac38

}

// Environ 返回当前配置下，命令将要运行的环境副本。
func (c *Cmd) Environ() []string { //md5:891b61dc22bc81880c267878539c842b

}