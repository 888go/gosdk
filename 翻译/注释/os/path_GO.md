
<原文开始>
// MkdirAll creates a directory named path,
// along with any necessary parents, and returns nil,
// or else returns an error.
// The permission bits perm (before umask) are used for all
// directories that MkdirAll creates.
// If path is already a directory, MkdirAll does nothing
// and returns nil.
<原文结束>

# <翻译开始>
// MkdirAll 创建一个名为 path 的目录，以及所有必要的父目录，并返回 nil，否则返回错误。
// 使用权限位 perm（在 umask 之前）为 MkdirAll 创建的所有目录。
// 如果 path 已经是一个目录，MkdirAll 将不执行任何操作并返回 nil。
# <翻译结束>


<原文开始>
// RemoveAll removes path and any children it contains.
// It removes everything it can but returns the first error
// it encounters. If the path does not exist, RemoveAll
// returns nil (no error).
// If there is an error, it will be of type *PathError.
<原文结束>

# <翻译开始>
// RemoveAll 移除路径及其包含的任何子项。
// 它尽可能多地删除，但返回遇到的第一个错误。如果路径不存在，RemoveAll 返回 nil（无错误）。
// 如果出现错误，错误类型为 *PathError。
# <翻译结束>


<原文开始>
// IsPathSeparator reports whether c is a directory separator character.
<原文结束>

# <翻译开始>
// IsPathSeparator 判断字符 c 是否为目录分隔符。
# <翻译结束>

