
<原文开始>
// LookPath searches for an executable named file in the directories
// named by the PATH environment variable. If file contains a slash,
// it is tried directly and the PATH is not consulted. The result will be
// an absolute path.
//
// LookPath differs from exec.LookPath in its handling of PATH lookups,
// which are used for file names without slashes. If exec.LookPath's
// PATH lookup would have returned an executable from the current directory,
// LookPath instead returns an error.
<原文结束>

# <翻译开始>
// LookPath 在由环境变量 PATH 指定的目录中搜索名为 file 的可执行文件。如果 file 中包含斜杠，
// 则会直接尝试该路径，不参考 PATH。返回结果将是一个绝对路径。
//
// LookPath 与 exec.LookPath 的区别在于处理 PATH 查找的方式，这用于查找不含斜杠的文件名。
// 如果 exec.LookPath 的 PATH 查找将从当前目录返回一个可执行文件，LookPath 则会返回错误。
# <翻译结束>


<原文开始>
// CommandContext is like Command but includes a context.
//
// The provided context is used to kill the process (by calling os.Process.Kill)
// if the context becomes done before the command completes on its own.
<原文结束>

# <翻译开始>
// CommandContext 与 Command 类似，但包含一个上下文。
//
// 如果在命令自行完成之前提供的上下文变为已结束状态，将使用该上下文来终止进程（通过调用 os.Process.Kill）。
# <翻译结束>


<原文开始>
// Command returns the Cmd struct to execute the named program with the given arguments.
// See exec.Command for most details.
//
// Command differs from exec.Command in its handling of PATH lookups,
// which are used when the program name contains no slashes.
// If exec.Command would have returned an exec.Cmd configured to run an
// executable from the current directory, Command instead
// returns an exec.Cmd that will return an error from Start or Run.
<原文结束>

# <翻译开始>
// Command 函数返回一个Cmd结构体，用于执行指定程序及给定参数。
// 具体细节可参阅exec.Command。
//
// Command 与 exec.Command 在处理PATH查找时有所不同，
// 当程序名中不含斜杠时，会用到此类查找。
// 若 exec.Command 将返回一个配置为运行当前目录下可执行文件的 exec.Cmd，
// 则 Command 将返回一个在调用 Start 或 Run 方法时将返回错误的 exec.Cmd。
# <翻译结束>

