// 版权所有 ? 2012 The Go Authors。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

//go:build windows

package main

import (
	"syscall"
)

// BUG(brainman): 在Windows 7系统上，MessageBeep Windows API 存在问题，
// 因此该示例在Windows 7上以服务方式运行时无法发出蜂鸣声。

var (
	beepFunc = syscall.MustLoadDLL("user32.dll").MustFindProc("MessageBeep")
)

func beep() {
	beepFunc.Call(0xffffffff)
}
