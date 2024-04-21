// 版权所有 2012 The Go 作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

//go:build windows

package main

import (
	"syscall"
)

// BUG(brainman)：MessageBeep Windows API 在 Windows 7 上存在故障，
// 导致此示例在以服务形式运行于 Windows 7 时无法发出蜂鸣声。

var (
	beepFunc = syscall.MustLoadDLL("user32.dll").MustFindProc("MessageBeep")
)

func beep() {
	beepFunc.Call(0xffffffff)
}
