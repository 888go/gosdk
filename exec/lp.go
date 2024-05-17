package exec

import (
	"os/exec"
)

// LookPath 在由 path 环境变量指定的目录中搜索名为 file 的可执行文件。
// 如果 file 以 "/"、"#"、"./" 或 "../" 开头，会直接尝试查找，而不参考路径。
// 成功时，结果为一个绝对路径。
//
// 在 Go 的早期版本中，LookPath 可能会返回相对于当前目录的路径。
// 自 Go 1.19 起，LookPath 将改为返回该路径以及一个满足 errors.Is(err, ErrDot) 的错误。更多详情请参阅包文档。
func LookPath(file string) (string, error) { //md5:154a1e41bca1510f44308fb5c079e855
	return exec.LookPath(file)
}

// ErrNotFound is the error resulting if a path search failed to find an executable file.
var ErrNotFound = exec.ErrNotFound //md5:d9032907a99fd1c3

//
//// ErrNotFound is the error resulting if a path search failed to find an executable file.
//var ErrNotFound = errors.New("executable file not found in $PATH") //md5:590f4b3cb9d53300
//
//
//
//// ErrNotFound is the error resulting if a path search failed to find an executable file.
//var ErrNotFound = errors.New("executable file not found in %PATH%") //md5:ddc9857366a762b6
//
