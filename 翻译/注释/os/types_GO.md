
<原文开始>
// Getpagesize returns the underlying system's memory page size.
<原文结束>

# <翻译开始>
// Getpagesize 返回底层系统的内存页大小。
# <翻译结束>


<原文开始>
// SameFile reports whether fi1 and fi2 describe the same file.
// For example, on Unix this means that the device and inode fields
// of the two underlying structures are identical; on other systems
// the decision may be based on the path names.
// SameFile only applies to results returned by this package's Stat.
// It returns false in other cases.
<原文结束>

# <翻译开始>
// SameFile 报告 fi1 和 fi2 是否描述了同一个文件。
// 例如，在Unix系统上，这意味着两个底层结构的设备和inode字段完全相同；
// 在其他系统上，判断可能基于路径名称。
// SameFile 只适用于本包的Stat方法返回的结果。
// 在其他情况下，它会返回false。
# <翻译结束>


<原文开始>
// File represents an open file descriptor.
<原文结束>

# <翻译开始>
// File 代表一个打开的文件描述符。
# <翻译结束>


<原文开始>
// A FileInfo describes a file and is returned by Stat and Lstat.
<原文结束>

# <翻译开始>
// FileInfo 用于描述一个文件，由 Stat 和 Lstat 函数返回。
# <翻译结束>


<原文开始>
// A FileMode represents a file's mode and permission bits.
// The bits have the same definition on all systems, so that
// information about files can be moved from one system
// to another portably. Not all bits apply to all systems.
// The only required bit is ModeDir for directories.
<原文结束>

# <翻译开始>
// FileMode 表示文件的模式和权限位。
// 这些位在所有系统上的定义相同，以便可以将文件的信息
// 从一个系统移植到另一个系统。并非所有位都适用于所有系统。
// 必须有的位是 ModeDir，表示目录。
# <翻译结束>

