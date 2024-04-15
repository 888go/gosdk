
<原文开始>
// LookPath searches for an executable named file in the
// directories named by the path environment variable.
// If file begins with "/", "#", "./", or "../", it is tried
// directly and the path is not consulted.
// On success, the result is an absolute path.
//
// In older versions of Go, LookPath could return a path relative to the current directory.
// As of Go 1.19, LookPath will instead return that path along with an error satisfying
// errors.Is(err, ErrDot). See the package documentation for more details.
<原文结束>

# <翻译开始>
// LookPath 在由 path 环境变量指定的目录中搜索名为 file 的可执行文件。
// 如果 file 以 "/"、"#"、"./" 或 "../" 开头，会直接尝试查找，而不参考路径。
// 成功时，结果为一个绝对路径。
//
// 在 Go 的早期版本中，LookPath 可能会返回相对于当前目录的路径。
// 自 Go 1.19 起，LookPath 将改为返回该路径以及一个满足 errors.Is(err, ErrDot) 的错误。更多详情请参阅包文档。
# <翻译结束>

