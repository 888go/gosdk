// 版权所有 2012 The Go 作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

//go:build windows

// Package mgr 用于管理 Windows 服务程序。
// 可以使用它来安装和卸载服务，以及启动、停止和暂停服务。该包可以查询/更改当前服务状态和配置参数。
package mgr

import (
	"syscall"
	"time"
	"unicode/utf16"
	"unsafe"

	"github.com/888go/gosdk/windows"
)

// Mgr 用于管理 Windows 服务。
type Mgr struct {
	Handle windows.Handle
}

// Connect 建立与服务控制管理器的连接

// ff:
func Connect() (*Mgr, error) {
	return ConnectRemote("")
}

// ConnectRemote 与名为 host 的计算机上的
// 服务控制管理器建立连接。

// ff:
// host:
func ConnectRemote(host string) (*Mgr, error) {
	var s *uint16
	if host != "" {
		s = syscall.StringToUTF16Ptr(host)
	}
	h, err := windows.OpenSCManager(s, nil, windows.SC_MANAGER_ALL_ACCESS)
	if err != nil {
		return nil, err
	}
	return &Mgr{Handle: h}, nil
}

// Disconnect 关闭与服务控制管理器 m 的连接

// ff:
func (m *Mgr) Disconnect() error {
	return windows.CloseServiceHandle(m.Handle)
}

type LockStatus struct {
	IsLocked bool          // 是否已锁定SCM
	Age      time.Duration // SCM被锁定的时长
	Owner    string        // 锁定SCM的用户名称
}

// LockStatus 返回服务控制管理器是否被系统锁定，以及锁定时长和锁定者。SCM（服务控制管理器）被锁定表明大多数服务操作将阻塞，直至系统解锁SCM。

// ff:
func (m *Mgr) LockStatus() (*LockStatus, error) {
	bytesNeeded := uint32(unsafe.Sizeof(windows.QUERY_SERVICE_LOCK_STATUS{}) + 1024)
	for {
		bytes := make([]byte, bytesNeeded)
		lockStatus := (*windows.QUERY_SERVICE_LOCK_STATUS)(unsafe.Pointer(&bytes[0]))
		err := windows.QueryServiceLockStatus(m.Handle, lockStatus, uint32(len(bytes)), &bytesNeeded)
		if err == windows.ERROR_INSUFFICIENT_BUFFER && bytesNeeded >= uint32(unsafe.Sizeof(windows.QUERY_SERVICE_LOCK_STATUS{})) {
			continue
		}
		if err != nil {
			return nil, err
		}
		status := &LockStatus{
			IsLocked: lockStatus.IsLocked != 0,
			Age:      time.Duration(lockStatus.LockDuration) * time.Second,
			Owner:    windows.UTF16PtrToString(lockStatus.LockOwner),
		}
		return status, nil
	}
}

func toPtr(s string) *uint16 {
	if len(s) == 0 {
		return nil
	}
	return syscall.StringToUTF16Ptr(s)
}

// toStringBlock 将 ss 中的字符串以 0 结尾，并将它们拼接在一起。同时在末尾额外添加一个 0。
func toStringBlock(ss []string) *uint16 {
	if len(ss) == 0 {
		return nil
	}
	t := ""
	for _, s := range ss {
		if s != "" {
			t += s + "\x00"
		}
	}
	if t == "" {
		return nil
	}
	t += "\x00"
	return &utf16.Encode([]rune(t))[0]
}

// 创建服务：在系统上安装新服务名称。
// 该服务将通过运行exepath指定的二进制文件来执行。
// 使用配置c来指定服务参数。
// 任何args参数将在服务启动时作为命令行参数传递；
// 这些参数与通过Service.Start或通过服务“属性”对话框中的“启动参数”字段传递的参数不同。

// ff:
// args:
// c:
// exepath:
// name:
func (m *Mgr) CreateService(name, exepath string, c Config, args ...string) (*Service, error) {
	if c.StartType == 0 {
		c.StartType = StartManual
	}
	if c.ServiceType == 0 {
		c.ServiceType = windows.SERVICE_WIN32_OWN_PROCESS
	}
	s := syscall.EscapeArg(exepath)
	for _, v := range args {
		s += " " + syscall.EscapeArg(v)
	}
	h, err := windows.CreateService(m.Handle, toPtr(name), toPtr(c.DisplayName),
		windows.SERVICE_ALL_ACCESS, c.ServiceType,
		c.StartType, c.ErrorControl, toPtr(s), toPtr(c.LoadOrderGroup),
		nil, toStringBlock(c.Dependencies), toPtr(c.ServiceStartName), toPtr(c.Password))
	if err != nil {
		return nil, err
	}
	if c.SidType != windows.SERVICE_SID_TYPE_NONE {
		err = updateSidType(h, c.SidType)
		if err != nil {
			windows.DeleteService(h)
			windows.CloseServiceHandle(h)
			return nil, err
		}
	}
	if c.Description != "" {
		err = updateDescription(h, c.Description)
		if err != nil {
			windows.DeleteService(h)
			windows.CloseServiceHandle(h)
			return nil, err
		}
	}
	if c.DelayedAutoStart {
		err = updateStartUp(h, c.DelayedAutoStart)
		if err != nil {
			windows.DeleteService(h)
			windows.CloseServiceHandle(h)
			return nil, err
		}
	}
	return &Service{Name: name, Handle: h}, nil
}

// OpenService 获取对服务名name的访问权限，以便对其进行查询和控制。

// ff:
// name:
func (m *Mgr) OpenService(name string) (*Service, error) {
	h, err := windows.OpenService(m.Handle, syscall.StringToUTF16Ptr(name), windows.SERVICE_ALL_ACCESS)
	if err != nil {
		return nil, err
	}
	return &Service{Name: name, Handle: h}, nil
}

// ListServices 在指定的服务控制管理器数据库 m 中枚举服务。
// 如果调用者对某个服务不具备 SERVICE_QUERY_STATUS 访问权限，
// 则该服务将被静默地从返回的服务列表中省略。

// ff:
func (m *Mgr) ListServices() ([]string, error) {
	var err error
	var bytesNeeded, servicesReturned uint32
	var buf []byte
	for {
		var p *byte
		if len(buf) > 0 {
			p = &buf[0]
		}
		err = windows.EnumServicesStatusEx(m.Handle, windows.SC_ENUM_PROCESS_INFO,
			windows.SERVICE_WIN32, windows.SERVICE_STATE_ALL,
			p, uint32(len(buf)), &bytesNeeded, &servicesReturned, nil, nil)
		if err == nil {
			break
		}
		if err != syscall.ERROR_MORE_DATA {
			return nil, err
		}
		if bytesNeeded <= uint32(len(buf)) {
			return nil, err
		}
		buf = make([]byte, bytesNeeded)
	}
	if servicesReturned == 0 {
		return nil, nil
	}
	services := unsafe.Slice((*windows.ENUM_SERVICE_STATUS_PROCESS)(unsafe.Pointer(&buf[0])), int(servicesReturned))

	var names []string
	for _, s := range services {
		name := windows.UTF16PtrToString(s.ServiceName)
		names = append(names, name)
	}
	return names, nil
}
