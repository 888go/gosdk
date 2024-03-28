
<原文开始>
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// 版权所有 ? 2009 Go作者。保留所有权利。
// 本源代码的使用受BSD风格
// 许可证约束，该许可证可在LICENSE文件中找到。
# <翻译结束>


<原文开始>
// Package exec runs external commands. It wraps os.StartProcess to make it
// easier to remap stdin and stdout, connect I/O with pipes, and do other
// adjustments.
//
// Unlike the "system" library call from C and other languages, the
// os/exec package intentionally does not invoke the system shell and
// does not expand any glob patterns or handle other expansions,
// pipelines, or redirections typically done by shells. The package
// behaves more like C's "exec" family of functions. To expand glob
// patterns, either call the shell directly, taking care to escape any
// dangerous input, or use the path/filepath package's Glob function.
// To expand environment variables, use package os's ExpandEnv.
//
// Note that the examples in this package assume a Unix system.
// They may not run on Windows, and they do not run in the Go Playground
// used by golang.org and godoc.org.
//
// # Executables in the current directory
//
// The functions Command and LookPath look for a program
// in the directories listed in the current path, following the
// conventions of the host operating system.
// Operating systems have for decades included the current
// directory in this search, sometimes implicitly and sometimes
// configured explicitly that way by default.
// Modern practice is that including the current directory
// is usually unexpected and often leads to security problems.
//
// To avoid those security problems, as of Go 1.19, this package will not resolve a program
// using an implicit or explicit path entry relative to the current directory.
// That is, if you run exec.LookPath("go"), it will not successfully return
// ./go on Unix nor .\go.exe on Windows, no matter how the path is configured.
// Instead, if the usual path algorithms would result in that answer,
// these functions return an error err satisfying errors.Is(err, ErrDot).
//
// For example, consider these two program snippets:
//
//	path, err := exec.LookPath("prog")
//	if err != nil {
//		log.Fatal(err)
//	}
//	use(path)
//
// and
//
//	cmd := exec.Command("prog")
//	if err := cmd.Run(); err != nil {
//		log.Fatal(err)
//	}
//
// These will not find and run ./prog or .\prog.exe,
// no matter how the current path is configured.
//
// Code that always wants to run a program from the current directory
// can be rewritten to say "./prog" instead of "prog".
//
// Code that insists on including results from relative path entries
// can instead override the error using an errors.Is check:
//
//	path, err := exec.LookPath("prog")
//	if errors.Is(err, exec.ErrDot) {
//		err = nil
//	}
//	if err != nil {
//		log.Fatal(err)
//	}
//	use(path)
//
// and
//
//	cmd := exec.Command("prog")
//	if errors.Is(cmd.Err, exec.ErrDot) {
//		cmd.Err = nil
//	}
//	if err := cmd.Run(); err != nil {
//		log.Fatal(err)
//	}
//
// Setting the environment variable GODEBUG=execerrdot=0
// disables generation of ErrDot entirely, temporarily restoring the pre-Go 1.19
// behavior for programs that are unable to apply more targeted fixes.
// A future version of Go may remove support for this variable.
//
// Before adding such overrides, make sure you understand the
// security implications of doing so.
// See https://go.dev/blog/path-security for more information.
<原文结束>

# <翻译开始>
// ```markdown
// Package exec 执行外部命令。它包装了 os.StartProcess，以便更容易重定向标准输入和输出、通过管道连接 I/O，并进行其他调整。
// 
// 与 C 语言和其他语言中的 "system" 库调用不同，os/exec 包刻意不调用系统 shell，也不展开任何 glob 模式或其他通常由 shell 处理的扩展、管道或重定向。该包的行为更类似于 C 语言的 "exec" 家族函数。若要展开 glob 模式，请直接调用 shell（同时确保转义所有危险输入），或者使用 path/filepath 包的 Glob 函数。若要展开环境变量，请使用 os 包的 ExpandEnv 函数。
// 
// 注意，本包中的示例假设运行于 Unix 系统。它们可能无法在 Windows 上运行，也无法在 golang.org 和 godoc.org 使用的 Go Playground 中运行。
// 
// # 当前目录下的可执行文件
//
// Command 和 LookPath 函数按照宿主操作系统的约定，在当前路径列出的目录中查找程序。
// 数十年来，操作系统一直将当前目录包含在这个搜索路径中，有时隐式包含，有时默认配置为显式包含。
// 现代实践中，将当前目录包含在搜索路径中通常是意外的，并且往往会导致安全问题。
// 
// 为了避免这些安全问题，自 Go 1.19 版本起，本包不再通过相对于当前目录的隐式或显式路径条目解析程序。
// 即使路径被这样配置，当你运行 exec.LookPath("go") 时，它也不会成功返回 ./go（在 Unix 系统上）或 .\go.exe（在 Windows 上）。相反，如果按常规路径算法得出这样的结果，这些函数会返回一个满足 errors.Is(err, ErrDot) 的错误。
// 
// 例如，考虑以下两个程序片段：
// ...
// （此处省略了示例代码）
// 
// 这些代码将不会找到并运行 ./prog 或 .\prog.exe，无论当前路径如何配置。
// 
// 如果始终希望从当前目录运行程序，可以将代码改写为使用 "./prog" 而不是 "prog"。
// 
// 如果坚持包含来自相对路径条目的结果，可以通过 errors.Is 检查来覆盖错误：
// ...
// （此处省略了示例代码）
// 
// 设置环境变量 GODEBUG=execerrdot=0 可以完全禁用 ErrDot 错误的生成，暂时恢复到 Go 1.19 之前的预行为，适用于那些无法应用更具针对性修复方案的程序。未来版本的 Go 可能会移除对这个变量的支持。
// 
// 在添加此类覆盖之前，请确保您理解这样做带来的安全影响。有关更多信息，请参阅 https://go.dev/blog/path-security 。
// ```
# <翻译结束>


<原文开始>
// Error is returned by LookPath when it fails to classify a file as an
// executable.
<原文结束>

# <翻译开始>
// 当LookPath无法将文件识别为可执行文件时，会返回Error。
# <翻译结束>


<原文开始>
// Name is the file name for which the error occurred.
<原文结束>

# <翻译开始>
// Name 是发生错误时的文件名。
# <翻译结束>


<原文开始>
// Err is the underlying error.
<原文结束>

# <翻译开始>
// Err 是底层的错误。
# <翻译结束>


<原文开始>
// ErrWaitDelay is returned by (*Cmd).Wait if the process exits with a
// successful status code but its output pipes are not closed before the
// command's WaitDelay expires.
<原文结束>

# <翻译开始>
// ErrWaitDelay 是在以下情况下由 (*Cmd).Wait 返回的错误：如果进程以成功状态退出，但在命令的 WaitDelay 超时之前其输出管道尚未关闭。
# <翻译结束>


<原文开始>
// wrappedError wraps an error without relying on fmt.Errorf.
<原文结束>

# <翻译开始>
// wrappedError 通过不依赖 fmt.Errorf 函数来封装一个错误。
# <翻译结束>


<原文开始>
// Cmd represents an external command being prepared or run.
//
// A Cmd cannot be reused after calling its Run, Output or CombinedOutput
// methods.
<原文结束>

# <翻译开始>
// Cmd 表示一个正在准备或执行的外部命令。
//
// 在调用其 Run、Output 或 CombinedOutput 方法后，Cmd 不能再被重用。
# <翻译结束>


<原文开始>
	// Path is the path of the command to run.
	//
	// This is the only field that must be set to a non-zero
	// value. If Path is relative, it is evaluated relative
	// to Dir.
<原文结束>

# <翻译开始>
// Path 是要运行的命令的路径。
//
// 这是唯一必须设置为非零值的字段。如果 Path 是相对路径，则会根据 Dir 进行相对定位。
# <翻译结束>


<原文开始>
	// Args holds command line arguments, including the command as Args[0].
	// If the Args field is empty or nil, Run uses {Path}.
	//
	// In typical use, both Path and Args are set by calling Command.
<原文结束>

# <翻译开始>
// Args 字段持有命令行参数，包括作为 Args[0] 的命令。如果 Args 字段为空或为 nil，Run 方法将使用 {Path}。
//
// 在典型用法中，Path 和 Args 两个字段都是通过调用 Command 方法来设置的。
# <翻译结束>


<原文开始>
	// Env specifies the environment of the process.
	// Each entry is of the form "key=value".
	// If Env is nil, the new process uses the current process's
	// environment.
	// If Env contains duplicate environment keys, only the last
	// value in the slice for each duplicate key is used.
	// As a special case on Windows, SYSTEMROOT is always added if
	// missing and not explicitly set to the empty string.
<原文结束>

# <翻译开始>
// Env 指定进程的环境变量。
// 每一项的格式为 "key=value"。
// 若 Env 为空，新进程将使用当前进程的环境变量。
// 若 Env 中存在重复的环境变量键名，则对于每组重复键名仅使用切片中最后一个值。
// 特别地，在 Windows 系统上，若未显式设置（且非空字符串），总是会添加 SYSTEMROOT 环境变量。
# <翻译结束>


<原文开始>
	// Dir specifies the working directory of the command.
	// If Dir is the empty string, Run runs the command in the
	// calling process's current directory.
<原文结束>

# <翻译开始>
// Dir 指定命令的工作目录。
// 若 Dir 为空字符串，Run 将在调用进程的当前目录下运行命令。
# <翻译结束>


<原文开始>
	// Stdin specifies the process's standard input.
	//
	// If Stdin is nil, the process reads from the null device (os.DevNull).
	//
	// If Stdin is an *os.File, the process's standard input is connected
	// directly to that file.
	//
	// Otherwise, during the execution of the command a separate
	// goroutine reads from Stdin and delivers that data to the command
	// over a pipe. In this case, Wait does not complete until the goroutine
	// stops copying, either because it has reached the end of Stdin
	// (EOF or a read error), or because writing to the pipe returned an error,
	// or because a nonzero WaitDelay was set and expired.
<原文结束>

# <翻译开始>
// Stdin 指定了进程的标准输入。
//
// 如果 Stdin 为 nil，则进程从空设备（os.DevNull）读取数据。
//
// 如果 Stdin 是一个 *os.File 类型，进程的标准输入将直接连接到该文件。
//
// 否则，在执行命令的过程中，会有一个单独的 goroutine 从 Stdin 读取数据并通过管道传递给命令。在这种情况下，Wait 方法直到该 goroutine 停止复制操作才会完成，这可能是因为已到达 Stdin 的末尾（EOF 或读取错误），或者因为向管道写入数据时发生错误，或者是设置了非零的 WaitDelay 并且超时了。
# <翻译结束>


<原文开始>
	// Stdout and Stderr specify the process's standard output and error.
	//
	// If either is nil, Run connects the corresponding file descriptor
	// to the null device (os.DevNull).
	//
	// If either is an *os.File, the corresponding output from the process
	// is connected directly to that file.
	//
	// Otherwise, during the execution of the command a separate goroutine
	// reads from the process over a pipe and delivers that data to the
	// corresponding Writer. In this case, Wait does not complete until the
	// goroutine reaches EOF or encounters an error or a nonzero WaitDelay
	// expires.
	//
	// If Stdout and Stderr are the same writer, and have a type that can
	// be compared with ==, at most one goroutine at a time will call Write.
<原文结束>

# <翻译开始>
// Stdout 和 Stderr 指定了进程的标准输出和错误输出。
//
// 如果两者之一为 nil，则将相应的文件描述符连接到空设备（os.DevNull）。
//
// 如果两者之一是 *os.File 类型，那么进程相应的输出会直接连接到该文件。
//
// 否则，在执行命令期间，通过管道从进程中读取数据，并将这些数据传递给相应的 Writer。在这种情况下，Wait 方法直到该 goroutine 遇到 EOF、遇到错误或非零的 WaitDelay 超时之前不会完成。
//
// 如果 Stdout 和 Stderr 是相同的写入器，并且它们的类型可以用 == 进行比较，在同一时间最多只有一个 goroutine 会调用 Write。
# <翻译结束>


<原文开始>
	// ExtraFiles specifies additional open files to be inherited by the
	// new process. It does not include standard input, standard output, or
	// standard error. If non-nil, entry i becomes file descriptor 3+i.
	//
	// ExtraFiles is not supported on Windows.
<原文结束>

# <翻译开始>
// ExtraFiles 指定了新进程将继承的额外打开文件。这不包括标准输入、标准输出或标准错误。
// 若该参数非空，其中第 i 个条目会成为文件描述符 3+i。
//
// 注意：ExtraFiles 在 Windows 上不受支持。
# <翻译结束>


<原文开始>
	// SysProcAttr holds optional, operating system-specific attributes.
	// Run passes it to os.StartProcess as the os.ProcAttr's Sys field.
<原文结束>

# <翻译开始>
// SysProcAttr 存储可选的、与操作系统相关的属性。
// 在调用os.StartProcess时，Run会将其作为os.ProcAttr结构体的Sys字段传递。
# <翻译结束>


<原文开始>
// Process is the underlying process, once started.
<原文结束>

# <翻译开始>
// Process 是底层进程，一旦启动后。
# <翻译结束>


<原文开始>
	// ProcessState contains information about an exited process.
	// If the process was started successfully, Wait or Run will
	// populate its ProcessState when the command completes.
<原文结束>

# <翻译开始>
// ProcessState 包含关于已退出进程的信息。
// 如果进程成功启动，当命令完成时，Wait 或 Run 方法将填充其 ProcessState。
# <翻译结束>


<原文开始>
// ctx is the context passed to CommandContext, if any.
<原文结束>

# <翻译开始>
// ctx 是传递给 CommandContext（如果有的话）的上下文。
# <翻译结束>


<原文开始>
// LookPath error, if any.
<原文结束>

# <翻译开始>
// LookPath 返回的错误，如果有的话。
# <翻译结束>


<原文开始>
	// If Cancel is non-nil, the command must have been created with
	// CommandContext and Cancel will be called when the command's
	// Context is done. By default, CommandContext sets Cancel to
	// call the Kill method on the command's Process.
	//
	// Typically a custom Cancel will send a signal to the command's
	// Process, but it may instead take other actions to initiate cancellation,
	// such as closing a stdin or stdout pipe or sending a shutdown request on a
	// network socket.
	//
	// If the command exits with a success status after Cancel is
	// called, and Cancel does not return an error equivalent to
	// os.ErrProcessDone, then Wait and similar methods will return a non-nil
	// error: either an error wrapping the one returned by Cancel,
	// or the error from the Context.
	// (If the command exits with a non-success status, or Cancel
	// returns an error that wraps os.ErrProcessDone, Wait and similar methods
	// continue to return the command's usual exit status.)
	//
	// If Cancel is set to nil, nothing will happen immediately when the command's
	// Context is done, but a nonzero WaitDelay will still take effect. That may
	// be useful, for example, to work around deadlocks in commands that do not
	// support shutdown signals but are expected to always finish quickly.
	//
	// Cancel will not be called if Start returns a non-nil error.
<原文结束>

# <翻译开始>
// 如果Cancel非空，则该命令必须通过CommandContext创建，当命令的上下文结束时，将会调用Cancel。默认情况下，CommandContext将设置Cancel为调用命令进程的Kill方法。
// 
// 通常，自定义的Cancel会向命令的进程发送一个信号以进行取消操作，但它也可以采取其他行动来启动取消，比如关闭stdin或stdout管道，或者通过网络套接字发送关闭请求。
// 
// 若在调用Cancel之后命令以成功状态退出，并且Cancel返回的错误不等同于os.ErrProcessDone，那么Wait和类似的方法将返回一个非nil的错误：要么是包装了由Cancel返回的错误，要么是来自上下文的错误。（如果命令以非成功状态退出，或者Cancel返回的错误包装了os.ErrProcessDone，Wait和类似的方法将继续返回命令通常的退出状态。）
// 
// 如果将Cancel设置为nil，在命令上下文结束时不会立即发生任何事情，但非零的WaitDelay仍然生效。这可能在某些场景下有用，例如绕过那些不支持关闭信号但预期总会快速完成的命令中的死锁问题。
// 
// 如果Start返回非nil的错误，将不会调用Cancel。
# <翻译结束>


<原文开始>
	// childIOFiles holds closers for any of the child process's
	// stdin, stdout, and/or stderr files that were opened by the Cmd itself
	// (not supplied by the caller). These should be closed as soon as they
	// are inherited by the child process.
<原文结束>

# <翻译开始>
// childIOFiles 用于存储由Cmd自身（而非调用者提供）打开的子进程的任何stdin、stdout和/或stderr文件的关闭器。一旦这些文件被子进程继承，就应当立即关闭它们。
# <翻译结束>


<原文开始>
	// parentIOPipes holds closers for the parent's end of any pipes
	// connected to the child's stdin, stdout, and/or stderr streams
	// that were opened by the Cmd itself (not supplied by the caller).
	// These should be closed after Wait sees the command and copying
	// goroutines exit, or after WaitDelay has expired.
<原文结束>

# <翻译开始>
// parentIOPipes 用于持有指向与子进程的stdin、stdout和（或）stderr流相连管道的父进程端的关闭器。
// 这些管道是由Cmd自身打开的（不是由调用者提供的）。
// 在Wait观察到命令执行完毕且复制goroutine退出后，或者在WaitDelay超时后，应关闭这些管道。
# <翻译结束>


<原文开始>
	// goroutine holds a set of closures to execute to copy data
	// to and/or from the command's I/O pipes.
<原文结束>

# <翻译开始>
// goroutine 持有一组用于执行的闭包，用于从命令的 I/O 管道复制数据或将数据复制到管道中。
# <翻译结束>


<原文开始>
	// If goroutineErr is non-nil, it receives the first error from a copying
	// goroutine once all such goroutines have completed.
	// goroutineErr is set to nil once its error has been received.
<原文结束>

# <翻译开始>
// 如果goroutineErr不为nil，当所有复制协程完成时，它会接收到这些协程中的第一个错误。
// 一旦goroutineErr的错误已被接收，其值将被设置为nil。
# <翻译结束>


<原文开始>
// If ctxResult is non-nil, it receives the result of watchCtx exactly once.
<原文结束>

# <翻译开始>
// 如果ctxResult非空，它将恰好接收一次watchCtx的结果。
# <翻译结束>


<原文开始>
	// The stack saved when the Command was created, if GODEBUG contains
	// execwait=2. Used for debugging leaks.
<原文结束>

# <翻译开始>
// 如果GODEBUG环境变量包含execwait=2，那么在创建Command时保存的堆栈。用于调试内存泄漏问题。
# <翻译结束>


<原文开始>
	// For a security release long ago, we created x/sys/execabs,
	// which manipulated the unexported lookPathErr error field
	// in this struct. For Go 1.19 we exported the field as Err error,
	// above, but we have to keep lookPathErr around for use by
	// old programs building against new toolchains.
	// The String and Start methods look for an error in lookPathErr
	// in preference to Err, to preserve the errors that execabs sets.
	//
	// In general we don't guarantee misuse of reflect like this,
	// but the misuse of reflect was by us, the best of various bad
	// options to fix the security problem, and people depend on
	// those old copies of execabs continuing to work.
	// The result is that we have to leave this variable around for the
	// rest of time, a compatibility scar.
	//
	// See https://go.dev/blog/path-security
	// and https://go.dev/issue/43724 for more context.
<原文结束>

# <翻译开始>
// 在很久以前的一个安全版本发布中，我们创建了x/sys/execabs，它操作了该结构体中的未导出字段lookPathErr错误。对于Go 1.19，我们将该字段导出为上述的Err错误，但由于旧版程序可能在新工具链上构建，因此我们必须保留lookPathErr供其使用。
// String和Start方法优先查找lookPathErr中的错误，而不是Err，以保持execabs设置的错误状态。
//
// 通常，我们并不保证像这样误用reflect，但此次误用reflect是由我们进行的，在修复安全问题的各种不佳方案中，这是最佳选择。并且人们依赖那些旧版execabs继续工作。
// 结果是我们必须永久保留这个变量，作为一个兼容性遗留问题（兼容性伤疤）。
//
// 有关更多背景信息，请参阅https://go.dev/blog/path-security 和 https://go.dev/issue/43724。
# <翻译结束>


<原文开始>
// A ctxResult reports the result of watching the Context associated with a
// running command (and sending corresponding signals if needed).
<原文结束>

# <翻译开始>
// A ctxResult 用于报告与正在运行的命令关联的上下文（Context）的结果（并根据需要发送相应的信号）。
# <翻译结束>


<原文开始>
	// If timer is non-nil, it expires after WaitDelay has elapsed after
	// the Context is done.
	//
	// (If timer is nil, that means that the Context was not done before the
	// command completed, or no WaitDelay was set, or the WaitDelay already
	// expired and its effect was already applied.)
<原文结束>

# <翻译开始>
// 如果timer非空，当Context完成之后等待WaitDelay时间后，它将超时。
//
// （如果timer为空，这意味着在命令执行完成之前Context未完成，或者未设置WaitDelay，或者WaitDelay已经过期且其效果已应用。）
# <翻译结束>


<原文开始>
			// Obtain the caller stack. (This is equivalent to runtime/debug.Stack,
			// copied to avoid importing the whole package.)
<原文结束>

# <翻译开始>
// 获取调用者堆栈。（这等同于runtime/debug.Stack，复制此代码以避免引入整个包。）
# <翻译结束>


<原文开始>
			// Update cmd.Path even if err is non-nil.
			// If err is ErrDot (especially on Windows), lp may include a resolved
			// extension (like .exe or .bat) that should be preserved.
<原文结束>

# <翻译开始>
// 即使 err 不为 nil，也要更新 cmd.Path。
// 如果 err 为 ErrDot（特别是在 Windows 系统中），lp 可能包含需要保留的已解析扩展名（如 .exe 或 .bat）。
# <翻译结束>


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
// failed to resolve path; report the original requested path (plus args)
<原文结束>

# <翻译开始>
// 未能解析路径；报告原始请求的路径（加上参数）
# <翻译结束>


<原文开始>
// report the exact executable path (plus args)
<原文结束>

# <翻译开始>
// 报告精确的可执行文件路径（及参数）
# <翻译结束>


<原文开始>
// interfaceEqual protects against panics from doing equality tests on
// two interfaces with non-comparable underlying types.
<原文结束>

# <翻译开始>
// interfaceEqual 用于防止对两个具有不可比较底层类型的接口进行相等性测试时出现的 panic（异常）。
# <翻译结束>


<原文开始>
// writerDescriptor returns an os.File to which the child process
// can write to send data to w.
//
// If w is nil, writerDescriptor returns a File that writes to os.DevNull.
<原文结束>

# <翻译开始>
// writerDescriptor 返回一个os.File，子进程可以向该文件写入数据以发送给w。
//
// 如果w为nil，writerDescriptor将返回一个写入os.DevNull的File。
# <翻译结束>


<原文开始>
// in case io.Copy stopped due to write error
<原文结束>

# <翻译开始>
// 如果io.Copy因写入错误而停止
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
// lookExtensions finds windows executable by its dir and path.
// It uses LookPath to try appropriate extensions.
// lookExtensions does not search PATH, instead it converts `prog` into `.\prog`.
<原文结束>

# <翻译开始>
// lookExtensions 在指定的dir和path下查找Windows可执行文件。
// 它使用LookPath尝试适用的扩展名。
// lookExtensions 不搜索PATH，而是将`prog`转换为`.\prog`进行查找。
# <翻译结束>


<原文开始>
// We assume that LookPath will only add file extension.
<原文结束>

# <翻译开始>
// 我们假设LookPath只会添加文件扩展名。
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
	// Check for doubled Start calls before we defer failure cleanup. If the prior
	// call to Start succeeded, we don't want to spuriously close its pipes.
<原文结束>

# <翻译开始>
// 在进行失败清理的延迟调用前，检查是否连续调用了两次Start。如果先前对Start的成功调用，我们不希望错误地关闭其管道。
# <翻译结束>


<原文开始>
// Don't allocate the goroutineErr channel unless there are goroutines to start.
<原文结束>

# <翻译开始>
// 仅当有 goroutine 需要启动时，才分配 goroutineErr 通道。
# <翻译结束>


<原文开始>
// Allow the goroutines' closures to be GC'd when they complete.
<原文结束>

# <翻译开始>
// 当goroutine完成时，允许其闭包被垃圾回收（GC）。
# <翻译结束>


<原文开始>
	// If we have anything to do when the command's Context expires,
	// start a goroutine to watch for cancellation.
	//
	// (Even if the command was created by CommandContext, a helper library may
	// have explicitly set its Cancel field back to nil, indicating that it should
	// be allowed to continue running after cancellation after all.)
<原文结束>

# <翻译开始>
// 如果当命令的上下文过期时，我们有任何需要执行的操作，那么启动一个goroutine来监视取消情况。
// 
// （即使命令是通过CommandContext创建的，某个辅助库可能已经将它的Cancel字段显式设置为nil，这意味着在取消后它仍应被允许继续运行。）
# <翻译结束>


<原文开始>
// watchCtx watches c.ctx until it is able to send a result to resultc.
//
// If c.ctx is done before a result can be sent, watchCtx calls c.Cancel,
// and/or kills cmd.Process it after c.WaitDelay has elapsed.
//
// watchCtx manipulates c.goroutineErr, so its result must be received before
// c.awaitGoroutines is called.
<原文结束>

# <翻译开始>
// watchCtx 监视 c.ctx，直到能够向 resultc 发送结果。
//
// 如果在能够发送结果之前 c.ctx 已完成，watchCtx 将调用 c.Cancel，
// 并且在 c.WaitDelay 时间间隔过后（如果适用）杀死 cmd.Process。
//
// watchCtx 会操作 c.goroutineErr，因此在其结果被接收之前必须先调用 c.awaitGoroutines。
# <翻译结束>


<原文开始>
			// We appear to have successfully interrupted the command, so any
			// program behavior from this point may be due to ctx even if the
			// command exits with code 0.
<原文结束>

# <翻译开始>
// 我们似乎成功地中止了命令，因此，即使该命令以代码 0 退出，从这一点开始的任何程序行为都可能是由于 ctx 引起的。
# <翻译结束>


<原文开始>
			// The process already finished: we just didn't notice it yet.
			// (Perhaps c.Wait hadn't been called, or perhaps it happened to race with
			// c.ctx being cancelled.) Don't inject a needless error.
<原文结束>

# <翻译开始>
// 进程已经完成：我们只是尚未注意到这一点。
// （可能是因为还没有调用 c.Wait，或者恰好与取消 c.ctx 的操作发生了竞态条件。）
// 不要注入不必要的错误。
# <翻译结束>


<原文开始>
		// c.Process.Wait returned and we've handed the timer off to c.Wait.
		// It will take care of goroutine shutdown from here.
<原文结束>

# <翻译开始>
// c.Process.Wait 已返回，并且我们已将定时器交接给 c.Wait。
// 从这里开始，它将负责处理协程的关闭。
# <翻译结束>


<原文开始>
		// We appear to have killed the process. c.Process.Wait should return a
		// non-nil error to c.Wait unless the Kill signal races with a successful
		// exit, and if that does happen we shouldn't report a spurious error,
		// so don't set err to anything here.
<原文结束>

# <翻译开始>
// 我们似乎已经杀掉了该进程。在非特殊情况下，c.Process.Wait 应该通过返回一个非空错误给 c.Wait，除非发送 Kill 信号与进程成功退出这两个操作发生了竞态条件。如果这种情况真的发生，我们不希望报告一个错误信息，所以在这种情况下不需要在这里设置 err（错误信息）。
# <翻译结束>


<原文开始>
			// Forward goroutineErr only if we don't have reason to believe it was
			// caused by a call to Cancel or Kill above.
<原文结束>

# <翻译开始>
// 如果我们没有理由认为错误是由上面调用Cancel或Kill导致的，则将goroutineErr向前传递。
# <翻译结束>


<原文开始>
			// Close the child process's I/O pipes, in case it abandoned some
			// subprocess that inherited them and is still holding them open
			// (see https://go.dev/issue/23019).
			//
			// We close the goroutine pipes only after we have sent any signals we're
			// going to send to the process (via Signal or Kill above): if we send
			// SIGKILL to the process, we would prefer for it to die of SIGKILL, not
			// SIGPIPE. (However, this may still cause any orphaned subprocesses to
			// terminate with SIGPIPE.)
<原文结束>

# <翻译开始>
// 关闭子进程的I/O管道，以防它放弃了一些继承了这些管道并仍在保持它们打开状态的子进程
// （参考：https://go.dev/issue/23019）。
//
// 我们仅在已通过上述Signal或Kill方法向进程发送所有要发送的信号后，才关闭goroutine管道：
// 如果我们向进程发送SIGKILL信号，我们更希望它因SIGKILL而终止，而不是SIGPIPE。（然而，这仍可能导致任何孤儿子进程因SIGPIPE而终止。）
# <翻译结束>


<原文开始>
			// Wait for the copying goroutines to finish, but report ErrWaitDelay for
			// the error: any other error here could result from closing the pipes.
<原文结束>

# <翻译开始>
// 等待复制goroutine完成，但将错误报告为ErrWaitDelay：这里出现任何其他错误可能都是由于关闭管道导致的。
# <翻译结束>


<原文开始>
		// Since we have already received the only result from c.goroutineErr,
		// set it to nil to prevent awaitGoroutines from blocking on it.
<原文结束>

# <翻译开始>
// 由于我们已经从 c.goroutineErr 中接收到唯一的结果，
// 将其设置为 nil，以防止 awaitGoroutines 在其上阻塞。
# <翻译结束>


<原文开始>
// An ExitError reports an unsuccessful exit by a command.
<原文结束>

# <翻译开始>
// ExitError 报告了命令的非正常退出情况。
# <翻译结束>


<原文开始>
	// Stderr holds a subset of the standard error output from the
	// Cmd.Output method if standard error was not otherwise being
	// collected.
	//
	// If the error output is long, Stderr may contain only a prefix
	// and suffix of the output, with the middle replaced with
	// text about the number of omitted bytes.
	//
	// Stderr is provided for debugging, for inclusion in error messages.
	// Users with other needs should redirect Cmd.Stderr as needed.
<原文结束>

# <翻译开始>
// Stderr 保存了在未通过其他方式收集标准错误输出时，Cmd.Output 方法产生的标准错误输出子集。
//
// 如果错误输出内容较长，Stderr 可能仅包含输出的前缀和后缀部分，中间部分会被替换为关于省略字节数量的文本。
//
// Stderr 主要用于调试并包含在错误消息中。对于有其他需求的用户，应根据需要重定向 Cmd.Stderr。
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
		// If c.Process.Wait returned an error, prefer that.
		// Otherwise, report any error from the watchCtx goroutine,
		// such as a Context cancellation or a WaitDelay overrun.
<原文结束>

# <翻译开始>
// 如果c.Process.Wait返回错误，则优先使用该错误。
// 否则，报告来自watchCtx goroutine的任何错误，
// 例如Context取消或WaitDelay超时。
# <翻译结束>


<原文开始>
		// Report an error from the copying goroutines only if the program otherwise
		// exited normally on its own. Otherwise, the copying error may be due to the
		// abnormal termination.
<原文结束>

# <翻译开始>
// 如果程序在正常情况下自行退出，则仅从复制goroutine报告错误。否则，复制错误可能是由于非正常终止导致的。
# <翻译结束>


<原文开始>
// awaitGoroutines waits for the results of the goroutines copying data to or
// from the command's I/O pipes.
//
// If c.WaitDelay elapses before the goroutines complete, awaitGoroutines
// forcibly closes their pipes and returns ErrWaitDelay.
//
// If timer is non-nil, it must send to timer.C at the end of c.WaitDelay.
<原文结束>

# <翻译开始>
// awaitGoroutines 等待 goroutines 将数据复制到命令的 I/O 管道或从管道中复制数据的结果。
//
// 如果在 goroutines 完成之前已超过 c.WaitDelay，awaitGoroutines 将强制关闭它们的管道并返回 ErrWaitDelay 错误。
//
// 如果 timer 非空，则它必须在 c.WaitDelay 结束时向 timer.C 发送信号。
# <翻译结束>


<原文开始>
// No running goroutines to await.
<原文结束>

# <翻译开始>
// 无运行中的goroutine需要等待。
# <翻译结束>


<原文开始>
// Avoid the overhead of starting a timer.
<原文结束>

# <翻译开始>
// 避免启动计时器带来的开销。
# <翻译结束>


<原文开始>
		// No existing timer was started: either there is no Context associated with
		// the command, or c.Process.Wait completed before the Context was done.
<原文结束>

# <翻译开始>
// 不存在已启动的计时器：要么该命令没有关联的上下文，要么在上下文完成之前，c.Process.Wait 已经执行完毕。
# <翻译结束>


<原文开始>
		// Wait for the copying goroutines to finish, but ignore any error
		// (since it was probably caused by closing the pipes).
<原文结束>

# <翻译开始>
// 等待复制的goroutine完成，但忽略任何错误（因为这可能是因为关闭管道所导致的）。
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
// prefixSuffixSaver is an io.Writer which retains the first N bytes
// and the last N bytes written to it. The Bytes() methods reconstructs
// it with a pretty error message.
<原文结束>

# <翻译开始>
// prefixSuffixSaver 是一个 io.Writer，它会保留写入的前 N 个字节和最后 N 个字节。Bytes() 方法会使用一个漂亮的错误消息来重新构建这些数据。
# <翻译结束>


<原文开始>
// max size of prefix or suffix
<原文结束>

# <翻译开始>
// 最大前缀或后缀长度
# <翻译结束>


<原文开始>
// ring buffer once len(suffix) == N
<原文结束>

# <翻译开始>
// 当 suffix 的长度等于 N 时，环形缓冲区
# <翻译结束>


<原文开始>
// offset to write into suffix
<原文结束>

# <翻译开始>
// 待写入后缀的偏移量
# <翻译结束>


<原文开始>
	// TODO(bradfitz): we could keep one large []byte and use part of it for
	// the prefix, reserve space for the '... Omitting N bytes ...' message,
	// then the ring buffer suffix, and just rearrange the ring buffer
	// suffix when Bytes() is called, but it doesn't seem worth it for
	// now just for error messages. It's only ~64KB anyway.
<原文结束>

# <翻译开始>
// TODO(bradfitz): 我们可以保留一个大的[]byte切片，并使用其中一部分用于前缀，预留空间用于“... 省略N字节 ...”的消息，然后是环形缓冲区的后缀。当调用Bytes()方法时，只需重新排列环形缓冲区的后缀即可。但目前看来，仅为错误消息这样做似乎不值得。毕竟它只有大约64KB大小。
# <翻译结束>


<原文开始>
// Only keep the last w.N bytes of suffix data.
<原文结束>

# <翻译开始>
// 仅保留后缀数据的最后 w.N 个字节。
# <翻译结束>


<原文开始>
// w.suffix is full now if p is non-empty. Overwrite it in a circle.
<原文结束>

# <翻译开始>
// 如果p非空，w.suffix现在已满。以循环方式覆盖它。
# <翻译结束>


<原文开始>
// fill appends up to len(p) bytes of p to *dst, such that *dst does not
// grow larger than w.N. It returns the un-appended suffix of p.
<原文结束>

# <翻译开始>
// fill 将 p 的前 len(p) 个字节追加到 *dst，确保 *dst 不会超过 w.N 的大小。它返回未追加的 p 后缀部分。
# <翻译结束>


<原文开始>
// environ returns a best-effort copy of the environment in which the command
// would be run as it is currently configured. If an error occurs in computing
// the environment, it is returned alongside the best-effort copy.
<原文结束>

# <翻译开始>
// environ 返回当前配置下，命令运行环境的一个尽可能完整的副本。如果在计算环境过程中出现错误，将连同尽可能完整的副本一起返回该错误。
# <翻译结束>


<原文开始>
// Note that the non-nil err is preserved despite env being overridden.
<原文结束>

# <翻译开始>
// 注意，尽管 env 被覆盖了，但非 nil 的 err 仍被保留。
# <翻译结束>


<原文开始>
				// Windows and Plan 9 do not use the PWD variable, so we don't need to
				// keep it accurate.
<原文结束>

# <翻译开始>
// Windows系统和Plan 9系统不使用PWD变量，因此我们不需要保持其准确性。
# <翻译结束>


<原文开始>
// Environ returns a copy of the environment in which the command would be run
// as it is currently configured.
<原文结束>

# <翻译开始>
// Environ 返回当前配置下，命令将要运行的环境副本。
# <翻译结束>


<原文开始>
//  Intentionally ignore errors: environ returns a best-effort environment no matter what.
<原文结束>

# <翻译开始>
// 故意忽略错误：无论发生什么，environ 都会尽力返回一个尽可能好的环境。
# <翻译结束>


<原文开始>
// dedupEnv returns a copy of env with any duplicates removed, in favor of
// later values.
// Items not of the normal environment "key=value" form are preserved unchanged.
// Except on Plan 9, items containing NUL characters are removed, and
// an error is returned along with the remaining values.
<原文结束>

# <翻译开始>
// dedupEnv 返回env的一个副本，其中任何重复项已被移除，并倾向于保留较晚的值。
// 不是标准环境“key=value”形式的项将保持不变。
// 除了在Plan 9系统中，包含NUL字符的项会被移除，并返回一个错误以及剩余的值。
# <翻译结束>


<原文开始>
// dedupEnvCase is dedupEnv with a case option for testing.
// If caseInsensitive is true, the case of keys is ignored.
// If nulOK is false, items containing NUL characters are allowed.
<原文结束>

# <翻译开始>
// dedupEnvCase 是 dedupEnv 的一个用于测试的选项，带有大小写处理功能。
// 如果 caseInsensitive 为真，则忽略键的大小写。
// 如果 nulOK 为假，则允许包含 NUL 字符的项。
# <翻译结束>


<原文开始>
	// Construct the output in reverse order, to preserve the
	// last occurrence of each key.
<原文结束>

# <翻译开始>
// 为了保留每个键的最后一次出现，以逆序方式构建输出。
# <翻译结束>


<原文开始>
		// Reject NUL in environment variables to prevent security issues (#56284);
		// except on Plan 9, which uses NUL as os.PathListSeparator (#56544).
<原文结束>

# <翻译开始>
// 拒绝在环境变量中使用NUL字符以防止安全问题 (#56284)；
// 但在Plan 9系统中除外，因为它将NUL用作os.PathListSeparator (#56544)。
# <翻译结束>


<原文开始>
			// We observe in practice keys with a single leading "=" on Windows.
			// TODO(#49886): Should we consume only the first leading "=" as part
			// of the key, or parse through arbitrarily many of them until a non-"="?
<原文结束>

# <翻译开始>
// 在实际应用中，我们发现Windows系统中有单个前置"="的键。
// TODO(#49886)：我们应该只将第一个前置的"="视为键的一部分来处理，还是解析任意数量的前置"="，直到遇到非"="为止？
# <翻译结束>


<原文开始>
				// The entry is not of the form "key=value" (as it is required to be).
				// Leave it as-is for now.
				// TODO(#52436): should we strip or reject these bogus entries?
<原文结束>

# <翻译开始>
// 这个条目不是“key=value”的形式（按照要求应当是这种形式）。
// 目前先保持原样不变。
// TODO(#52436)：我们应该移除还是拒绝这些无效的条目？
# <翻译结束>


<原文开始>
// Now reverse the slice to restore the original order.
<原文结束>

# <翻译开始>
// 现在反转切片以恢复原始顺序。
# <翻译结束>


<原文开始>
// addCriticalEnv adds any critical environment variables that are required
// (or at least almost always required) on the operating system.
// Currently this is only used for Windows.
<原文结束>

# <翻译开始>
// addCriticalEnv 添加任何在操作系统上必需（或几乎总是必需）的关键环境变量。
// 目前，此功能仅在Windows系统上使用。
# <翻译结束>

