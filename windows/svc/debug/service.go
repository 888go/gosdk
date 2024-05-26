// 版权所有 2012 The Go 作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

//go:build windows

// Package debug 提供在控制台执行 svc.Handler 的功能。
package debug

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/888go/gosdk/windows/svc"
)

// Run通过调用相应的处理函数来执行服务名。
// 该过程在控制台运行，与实际服务不同。使用Ctrl+C向您的服务发送“停止”命令。

// ff:
// name:
// handler:
func Run(name string, handler svc.Handler) error {
	cmds := make(chan svc.ChangeRequest)
	changes := make(chan svc.Status)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig)

	go func() {
		status := svc.Status{State: svc.Stopped}
		for {
			select {
			case <-sig:
				cmds <- svc.ChangeRequest{Cmd: svc.Stop, CurrentStatus: status}
			case status = <-changes:
			}
		}
	}()

	_, errno := handler.Execute([]string{name}, cmds, changes)
	if errno != 0 {
		return syscall.Errno(errno)
	}
	return nil
}
