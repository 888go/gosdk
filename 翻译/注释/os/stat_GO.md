
<原文开始>
// Stat returns a FileInfo describing the named file.
// If there is an error, it will be of type *PathError.
<原文结束>

# <翻译开始>
// Stat返回一个描述指定名称文件的FileInfo。
// 如果出现错误，该错误将为*PathError类型。
# <翻译结束>


<原文开始>
// Lstat returns a FileInfo describing the named file.
// If the file is a symbolic link, the returned FileInfo
// describes the symbolic link. Lstat makes no attempt to follow the link.
// If there is an error, it will be of type *PathError.
//
// On Windows, if the file is a reparse point that is a surrogate for another
// named entity (such as a symbolic link or mounted folder), the returned
// FileInfo describes the reparse point, and makes no attempt to resolve it.
<原文结束>

# <翻译开始>
// Lstat 返回一个描述指定文件的FileInfo。如果文件是一个符号链接，返回的FileInfo将描述该符号链接。Lstat 不会尝试跟踪链接。
// 如果出现错误，错误类型将是 *PathError。
// 
// 在Windows上，如果文件是一个代理重定向点（如符号链接或挂载的文件夹），返回的FileInfo将描述重定向点，并不会尝试解析它。
# <翻译结束>


<原文开始>
// Stat returns the FileInfo structure describing file.
// If there is an error, it will be of type *PathError.
<原文结束>

# <翻译开始>
// Stat返回描述文件的FileInfo结构体。如果出现错误，错误类型为*PathError。
# <翻译结束>

