package os

// File 代表一个打开的文件描述符。
type File struct { //md5:1b6a20b413818c28d384dba920ab157e

}

// FileInfo 用于描述一个文件，由 Stat 和 Lstat 函数返回。
type FileInfo = fs.FileInfo //md5:4c3f090d15e35247e13ba72e8da3ce80



// FileMode 表示文件的模式和权限位。
// 这些位在所有系统上的定义相同，以便可以将文件的信息
// 从一个系统移植到另一个系统。并非所有位都适用于所有系统。
// 必须有的位是 ModeDir，表示目录。
type FileMode = fs.FileMode //md5:e4b755cd6b7e001da355fd0b271a7f43



// Getpagesize 返回底层系统的内存页大小。
func Getpagesize() int { //md5:d379c113740c775f0f6aaac91ecbd1be

}

// SameFile 报告 fi1 和 fi2 是否描述了同一个文件。
// 例如，在Unix系统上，这意味着两个底层结构的设备和inode字段完全相同；
// 在其他系统上，判断可能基于路径名称。
// SameFile 只适用于本包的Stat方法返回的结果。
// 在其他情况下，它会返回false。
func SameFile(fi1, fi2 FileInfo) bool { //md5:b4e76f9d966d9daea55a383bca41dc5f

}