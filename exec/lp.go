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

// ff:查找路径
// file:文件
func LookPath(file string) (string, error) { //md5:154a1e41bca1510f44308fb5c079e855
	return exec.LookPath(file)
}
