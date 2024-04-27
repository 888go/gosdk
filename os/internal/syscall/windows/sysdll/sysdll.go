//版权所有2016 The Go Authors。所有权利保留。
//使用此源代码受BSD风格的
//可在LICENSE文件中找到的许可协议管辖。

//go:build windows

// 包sysdll是一个内部的叶子包，用于记录和报告Go自身使用了哪些Windows DLL名称。这些DLL随后仅从System32目录加载。参见问题14959。
package sysdll

// IsSystemDLL 判断指定的 dll 键（一个基名称，如 "foo.dll"）是否是系统 DLL，它应该只从 Windows 系统目录加载。
// 
// 文件名是区分大小写的，但这没关系，因为使用 Add 注册时的大小写与稍后使用 LoadDLL 时的大小写相同。
// 
// 它没有关联的互斥锁，只能串行修改（目前：仅在初始化期间），并且不应与 DLL 加载操作并发进行。
var IsSystemDLL = map[string]bool{}

// 添加注释，说明dll是一个系统32库，应该仅从Windows SYSTEM32目录加载。它返回其参数，以便于生成代码的使用。
func Add(dll string) string {
	IsSystemDLL[dll] = true
	return dll
}
