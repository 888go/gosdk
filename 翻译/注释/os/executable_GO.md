
<原文开始>
// Executable returns the path name for the executable that started
// the current process. There is no guarantee that the path is still
// pointing to the correct executable. If a symlink was used to start
// the process, depending on the operating system, the result might
// be the symlink or the path it pointed to. If a stable result is
// needed, path/filepath.EvalSymlinks might help.
//
// Executable returns an absolute path unless an error occurred.
//
// The main use case is finding resources located relative to an
// executable.
<原文结束>

# <翻译开始>
// Executable 返回启动当前进程的可执行文件的路径名。无法保证该路径仍然指向正确的可执行文件。如果使用软链接启动了进程，根据操作系统不同，结果可能是软链接本身或它指向的路径。如果需要稳定的结果，可以使用 path/filepath.EvalSymlinks 进行帮助。
//
// 如果没有发生错误，Executable 返回一个绝对路径。
//
// 主要用途是找到与可执行文件相对应的资源。
# <翻译结束>

