package os

import "os"

// Args 保存了命令行参数，从程序名称开始。
var Args []string

// Getuid 返回调用者的数字用户ID。
//
// 在Windows上，它返回-1。

// ff:
func Getuid() int { //md5:ed548081987604fdedc17b7ebb657030
	return os.Getuid()
}

// Geteuid 返回调用者的有效用户 ID 号。
//
// 在 Windows 系统中，它返回 -1。

// ff:
func Geteuid() int { //md5:e7a8d216aa40bbcb978c55ad12d802f0
	return os.Geteuid()
}

// Getgid 返回调用者所属的组的数字标识符。
//
// 在 Windows 系统中，它返回 -1。

// ff:
func Getgid() int { //md5:4f2e724b4884c5cfef42abb44a79b112
	return os.Getgid()
}

// Getegid 返回调用者当前有效的数字组ID。
//
// 在Windows系统上，它返回-1。

// ff:
func Getegid() int { //md5:bc1a9fba5dd8fddd6171764f1a230ba7
	return os.Getegid()
}

// Getgroups 返回一个列表，包含调用者所属的组的数字ID。
//
// 在Windows系统上，它会返回syscall.EWINDOWS错误。可以查看os/user包以寻求可能的替代方法。

// ff:
func Getgroups() ([]int, error) { //md5:ba25b1a38d6807aad8a1f78051f621d0
	return os.Getgroups()
}

// Exit 会导致当前程序使用给定的状态码退出。
// 通常，状态码为零表示成功，非零表示错误。
// 程序会立即终止；延迟执行的函数不会运行。
//
// 为了兼容性，状态码应该在[0, 125]范围内。

// ff:
// code:
func Exit(code int) { //md5:0901ce78771fb26a6481d82aa618730c
	os.Exit(code)
}
