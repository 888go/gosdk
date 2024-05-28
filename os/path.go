package os

import "os"

const (
	PathSeparator     = os.PathSeparator     // 操作系统特定的路径分隔符
	PathListSeparator = os.PathListSeparator // 操作系统特定的路径列表分隔符
) // md5:3b0736e915c07d11

//const (
//	PathSeparator     = '/'    // 操作系统特定的路径分隔符
//	PathListSeparator = '\000' // 操作系统特定的路径列表分隔符
//) // md5:1909ae9d3917a681
//
//const (
//	PathSeparator     = '/' // 操作系统特定的路径分隔符
//	PathListSeparator = ':' // 操作系统特定的路径列表分隔符
//) // md5:1e9afc997627521d

// MkdirAll 创建一个名为 path 的目录，以及所有必要的父目录，并返回 nil，否则返回错误。
// 使用权限位 perm（在 umask 之前）为 MkdirAll 创建的所有目录。
// 如果 path 已经是一个目录，MkdirAll 将不执行任何操作并返回 nil。
func MkdirAll(path string, perm FileMode) error { //md5:db52ef7497ef022cd3e461eb06d35cef
	return os.MkdirAll(path, perm)
}

// RemoveAll 移除路径及其包含的任何子项。
// 它尽可能多地删除，但返回遇到的第一个错误。如果路径不存在，RemoveAll 返回 nil（无错误）。
// 如果出现错误，错误类型为 *PathError。
func RemoveAll(path string) error { //md5:c12e43494b561070fc271f4708a20893
	return os.RemoveAll(path)
}

// IsPathSeparator 判断字符 c 是否为目录分隔符。
func IsPathSeparator(c uint8) bool { //md5:9aab2b529fc5134dc3bdf50343d15f81
	return os.IsPathSeparator(c)
}
