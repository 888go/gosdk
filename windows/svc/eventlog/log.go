// 版权所有 2012 The Go 作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

//go:build windows

// 包eventlog实现了对Windows事件日志的访问。
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

// ff:
// source:
func Open(source string) (*Log, error) {
	return OpenRemote("", source)
}

// OpenRemote与Open功能相同，但作用于另一台计算机主机上。

// ff:
// source:
// host:
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

// ff:
func (l *Log) Close() error {
	return windows.DeregisterEventSource(l.Handle)
}

func (l *Log) report(etype uint16, eid uint32, msg string) error {
	ss := []*uint16{syscall.StringToUTF16Ptr(msg)}
	return windows.ReportEvent(l.Handle, etype, 0, eid, 0, 1, 0, &ss[0], nil)
}

// Info 函数向事件日志 l 的末尾写入一个信息事件，该事件包含事件 ID（eid）及消息（msg）。当使用 EventCreate.exe 工具时，eid 的值必须在 1 到 1000 之间。

// ff:
// msg:
// eid:
func (l *Log) Info(eid uint32, msg string) error {
	return l.report(windows.EVENTLOG_INFORMATION_TYPE, eid, msg)
}

// Warning 将带有事件ID eid 的警告事件消息写入事件日志l的末尾。
// 使用EventCreate.exe时，eid必须在1到1000之间。

// ff:
// msg:
// eid:
func (l *Log) Warning(eid uint32, msg string) error {
	return l.report(windows.EVENTLOG_WARNING_TYPE, eid, msg)
}

// Error 函数向事件日志 l 的末尾写入一个错误事件消息，该消息带有事件 ID eid。
// 当使用 EventCreate.exe 时，eid 的值必须介于 1 和 1000 之间。

// ff:
// msg:
// eid:
func (l *Log) Error(eid uint32, msg string) error {
	return l.report(windows.EVENTLOG_ERROR_TYPE, eid, msg)
}
