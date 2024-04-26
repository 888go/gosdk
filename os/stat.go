package os

import "os"

// Stat返回一个描述指定名称文件的FileInfo。
// 如果出现错误，该错误将为*PathError类型。

// ff:
// FileInfo:
// name:
func Stat(name string) (FileInfo, error) { //md5:0fe76bf45aa469ca3743e871fe7053e9
	return os.Stat(name)
}

// Lstat 返回一个描述指定文件的FileInfo。如果文件是一个符号链接，返回的FileInfo将描述该符号链接。Lstat 不会尝试跟踪链接。
// 如果出现错误，错误类型将是 *PathError。
//
// 在Windows上，如果文件是一个代理重定向点（如符号链接或挂载的文件夹），返回的FileInfo将描述重定向点，并不会尝试解析它。

// ff:
// FileInfo:
// name:
func Lstat(name string) (FileInfo, error) { //md5:c74422e3b8e588a7cd1b053e845198d6
	return os.Lstat(name)
}

//// Stat返回描述文件的FileInfo结构体。如果出现错误，错误类型为*PathError。
//func (file *File) Stat() (FileInfo, error) { //md5:1d0732de4e88e21b512a96c0422d8dd7
//
//}
