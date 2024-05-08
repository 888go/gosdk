// 版权所有 2012 The Go 作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

//---build---//go:build windows

package debug

import (
	"os"
	"strconv"
)

// Log接口允许使用不同的日志实现。
type Log interface {
	Close() error
	Info(eid uint32, msg string) error
	Warning(eid uint32, msg string) error
	Error(eid uint32, msg string) error
}

// ConsoleLog 提供对控制台的访问。
type ConsoleLog struct {
	Name string
}

// New 创建新的 ConsoleLog。
func New(source string) *ConsoleLog {
	return &ConsoleLog{Name: source}
}

// Close 关闭控制台日志 l。
func (l *ConsoleLog) Close() error {
	return nil
}

func (l *ConsoleLog) report(kind string, eid uint32, msg string) error {
	s := l.Name + "." + kind + "(" + strconv.Itoa(int(eid)) + "): " + msg + "\n"
	_, err := os.Stdout.Write([]byte(s))
	return err
}

// Info 将带有事件ID eid 的信息事件msg写入到控制台l。
func (l *ConsoleLog) Info(eid uint32, msg string) error {
	return l.report("info", eid, msg)
}

// Warning 向控制台 l 输出一个带有事件ID eid 的警告事件消息。
func (l *ConsoleLog) Warning(eid uint32, msg string) error {
	return l.report("warn", eid, msg)
}

// 将带有事件ID eid 的错误事件消息 msg 写入到控制台 l。
func (l *ConsoleLog) Error(eid uint32, msg string) error {
	return l.report("error", eid, msg)
}
