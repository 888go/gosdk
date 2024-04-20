// 版权所有 ? 2012 The Go Authors。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

//go:build windows

// package eventlog 实现了对 Windows 事件日志的访问功能。
package eventlog

import (
	"errors"
	"syscall"

	"github.com/888go/gosdk/windows"
)

// Log 提供对系统日志的访问。
type Log struct {
	Handle windows.Handle
}

// Open 获取指定事件日志的句柄。
func Open(source string) (*Log, error) {
	return OpenRemote("", source)
}

// OpenRemote 与 Open 执行相同的操作，但针对的是另一台计算机主机。
func OpenRemote(host, source string) (*Log, error) {
	if source == "" {
		return nil, errors.New("Specify event log source")
	}
	var s *uint16
	if host != "" {
		s = syscall.StringToUTF16Ptr(host)
	}
	h, err := windows.RegisterEventSource(s, syscall.StringToUTF16Ptr(source))
	if err != nil {
		return nil, err
	}
	return &Log{Handle: h}, nil
}

// Close 关闭事件日志 l。
func (l *Log) Close() error {
	return windows.DeregisterEventSource(l.Handle)
}

func (l *Log) report(etype uint16, eid uint32, msg string) error {
	ss := []*uint16{syscall.StringToUTF16Ptr(msg)}
	return windows.ReportEvent(l.Handle, etype, 0, eid, 0, 1, 0, &ss[0], nil)
}

// Info 将带有事件ID eid 的信息事件msg写入事件日志l的末尾。
// 使用EventCreate.exe时，eid必须介于1和1000之间。
func (l *Log) Info(eid uint32, msg string) error {
	return l.report(windows.EVENTLOG_INFORMATION_TYPE, eid, msg)
}

// Warning 在事件日志 l 的末尾写入一个带有事件ID eid 的警告事件消息。
// 当使用 EventCreate.exe 时，eid 必须介于 1 和 1000 之间。
func (l *Log) Warning(eid uint32, msg string) error {
	return l.report(windows.EVENTLOG_WARNING_TYPE, eid, msg)
}

// Error 函数向事件日志 l 的末尾写入一个具有事件ID eid 的错误事件消息。
// 当使用 EventCreate.exe 时，eid 必须在 1 到 1000 之间。
func (l *Log) Error(eid uint32, msg string) error {
	return l.report(windows.EVENTLOG_ERROR_TYPE, eid, msg)
}
