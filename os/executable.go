package os

import "os"

// Executable 返回启动当前进程的可执行文件的路径名。无法保证该路径仍然指向正确的可执行文件。如果使用软链接启动了进程，根据操作系统不同，结果可能是软链接本身或它指向的路径。如果需要稳定的结果，可以使用 path/filepath.EvalSymlinks 进行帮助。
//
// 如果没有发生错误，Executable 返回一个绝对路径。
//
// 主要用途是找到与可执行文件相对应的资源。
// ff:取当前进程文件路径
func Executable() (string, error) { //md5:c38e25d3342843455e87e2cc6b19f453
	return os.Executable()
}
