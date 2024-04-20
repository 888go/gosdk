// 版权所有 ? 2015 The Go Authors。保留所有权利。
// 本源代码的使用受 BSD 风格许可协议约束，
// 该协议可在 LICENSE 文件中找到。

//go:build windows

package registry

import "syscall"

const (
	_REG_OPTION_NON_VOLATILE = 0

	_REG_CREATED_NEW_KEY     = 1
	_REG_OPENED_EXISTING_KEY = 2

	_ERROR_NO_MORE_ITEMS syscall.Errno = 259
)

func LoadRegLoadMUIString() error {
	return procRegLoadMUIStringW.Find()
}

//sys	regCreateKeyEx(key syscall.Handle, subkey *uint16, reserved uint32, class *uint16, options uint32, desired uint32, sa *syscall.SecurityAttributes, result *syscall.Handle, disposition *uint32) (regerrno error) = advapi32.RegCreateKeyExW
// 系统调用：使用advapi32.RegCreateKeyExW函数创建一个注册表键。参数包括父键句柄（key）、子键名称指针（subkey）、保留值（reserved）、类名指针（class）、选项（options）、期望访问权限（desired）、安全属性指针（sa）、结果键句柄指针（result）以及操作结果指示器指针（disposition）。返回值为错误代码（regerrno）。
// 
//sys	regDeleteKey(key syscall.Handle, subkey *uint16) (regerrno error) = advapi32.RegDeleteKeyW
// 系统调用：使用advapi32.RegDeleteKeyW函数删除指定的注册表子键。参数包括父键句柄（key）和子键名称指针（subkey）。返回值为错误代码（regerrno）。
// 
//sys	regSetValueEx(key syscall.Handle, valueName *uint16, reserved uint32, vtype uint32, buf *byte, bufsize uint32) (regerrno error) = advapi32.RegSetValueExW
// 系统调用：使用advapi32.RegSetValueExW函数设置注册表键值。参数包括键句柄（key）、值名指针（valueName）、保留值（reserved）、值类型（vtype）、值数据缓冲区指针（buf）以及缓冲区大小（bufsize）。返回值为错误代码（regerrno）。
// 
//sys	regEnumValue(key syscall.Handle, index uint32, name *uint16, nameLen *uint32, reserved *uint32, valtype *uint32, buf *byte, buflen *uint32) (regerrno error) = advapi32.RegEnumValueW
// 系统调用：使用advapi32.RegEnumValueW函数枚举注册表键下的值。参数包括键句柄（key）、索引值（index）、值名缓冲区指针（name）、值名长度指针（nameLen）、保留值指针（reserved）、值类型指针（valtype）、值数据缓冲区指针（buf）以及缓冲区大小指针（buflen）。返回值为错误代码（regerrno）。
// 
//sys	regDeleteValue(key syscall.Handle, name *uint16) (regerrno error) = advapi32.RegDeleteValueW
// 系统调用：使用advapi32.RegDeleteValueW函数删除注册表键下的指定值。参数包括键句柄（key）和值名指针（name）。返回值为错误代码（regerrno）。
// 
//sys   regLoadMUIString(key syscall.Handle, name *uint16, buf *uint16, buflen uint32, buflenCopied *uint32, flags uint32, dir *uint16) (regerrno error) = advapi32.RegLoadMUIStringW
// 系统调用：使用advapi32.RegLoadMUIStringW函数加载注册表中的多语言用户界面字符串。参数包括键句柄（key）、值名指针（name）、缓冲区指针（buf）、缓冲区长度（buflen）、实际复制长度指针（buflenCopied）、标志（flags）以及目录路径指针（dir）。返回值为错误代码（regerrno）。
// 
//sys	regConnectRegistry(machinename *uint16, key syscall.Handle, result *syscall.Handle) (regerrno error) = advapi32.RegConnectRegistryW
// 系统调用：使用advapi32.RegConnectRegistryW函数连接到指定机器上的注册表键。参数包括目标机器名指针（machinename）、本地键句柄（key）以及结果远程键句柄指针（result）。返回值为错误代码（regerrno）。

//sys	expandEnvironmentStrings(src *uint16, dst *uint16, size uint32) (n uint32, err error) = kernel32.ExpandEnvironmentStringsW
// 
// 系统调用：expandEnvironmentStrings
// 输入参数：
//   src: 指向包含待扩展环境变量字符串的uint16指针
//   dst: 指向用于存储扩展后结果的uint16指针
//   size: 定义dst缓冲区大小的uint32值
// 返回值：
//   n: 实际写入dst缓冲区的字符数，以uint32表示
//   err: 扩展过程中可能遇到的错误信息（如果有的话）
// 该函数实际调用Windows API中的kernel32.ExpandEnvironmentStringsW方法
