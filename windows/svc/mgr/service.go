// 版权所有 2012 The Go 作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

//go:build windows

package mgr

import (
	"syscall"
	"unsafe"

	"github.com/888go/gosdk/windows"
	"github.com/888go/gosdk/windows/svc"
)

// Service 用于访问 Windows 服务。
type Service struct {
	Name   string
	Handle windows.Handle
}

// Delete 标记服务 s 以便从服务控制管理器数据库中删除。
func (s *Service) Delete() error {
	return windows.DeleteService(s.Handle)
}

// Close 放弃对服务 s 的访问。
func (s *Service) Close() error {
	return windows.CloseServiceHandle(s.Handle)
}

// Start 启动服务 s。
// args 将被传递给 svc.Handler.Execute。
func (s *Service) Start(args ...string) error {
	var p **uint16
	if len(args) > 0 {
		vs := make([]*uint16, len(args))
		for i := range vs {
			vs[i] = syscall.StringToUTF16Ptr(args[i])
		}
		p = &vs[0]
	}
	return windows.StartService(s.Handle, uint32(len(args)), p)
}

// Control 向服务 s 发送状态变更请求 c。它返回服务上报给服务控制管理器的最新状态，以及在状态变更请求未被接受时返回的错误。
// 注意，只有当状态变更请求成功，或者因错误 ERROR_INVALID_SERVICE_CONTROL、ERROR_SERVICE_CANNOT_ACCEPT_CTRL 或 ERROR_SERVICE_NOT_ACTIVE 而失败时，才会设置返回的服务状态。
func (s *Service) Control(c svc.Cmd) (svc.Status, error) {
	var t windows.SERVICE_STATUS
	err := windows.ControlService(s.Handle, uint32(c), &t)
	if err != nil &&
		err != windows.ERROR_INVALID_SERVICE_CONTROL &&
		err != windows.ERROR_SERVICE_CANNOT_ACCEPT_CTRL &&
		err != windows.ERROR_SERVICE_NOT_ACTIVE {
		return svc.Status{}, err
	}
	return svc.Status{
		State:   svc.State(t.CurrentState),
		Accepts: svc.Accepted(t.ControlsAccepted),
	}, err
}

// Query 返回服务 s 当前的状态
func (s *Service) Query() (svc.Status, error) {
	var t windows.SERVICE_STATUS_PROCESS
	var needed uint32
	err := windows.QueryServiceStatusEx(s.Handle, windows.SC_STATUS_PROCESS_INFO, (*byte)(unsafe.Pointer(&t)), uint32(unsafe.Sizeof(t)), &needed)
	if err != nil {
		return svc.Status{}, err
	}
	return svc.Status{
		State:                   svc.State(t.CurrentState),
		Accepts:                 svc.Accepted(t.ControlsAccepted),
		ProcessId:               t.ProcessId,
		Win32ExitCode:           t.Win32ExitCode,
		ServiceSpecificExitCode: t.ServiceSpecificExitCode,
	}, nil
}

// ListDependentServices 返回依赖于服务s且状态与给定状态匹配的服务名称。
func (s *Service) ListDependentServices(status svc.ActivityStatus) ([]string, error) {
	var bytesNeeded, returnedServiceCount uint32
	var services []windows.ENUM_SERVICE_STATUS
	for {
		var servicesPtr *windows.ENUM_SERVICE_STATUS
		if len(services) > 0 {
			servicesPtr = &services[0]
		}
		allocatedBytes := uint32(len(services)) * uint32(unsafe.Sizeof(windows.ENUM_SERVICE_STATUS{}))
		err := windows.EnumDependentServices(s.Handle, uint32(status), servicesPtr, allocatedBytes, &bytesNeeded,
			&returnedServiceCount)
		if err == nil {
			break
		}
		if err != syscall.ERROR_MORE_DATA {
			return nil, err
		}
		if bytesNeeded <= allocatedBytes {
			return nil, err
		}
		// ERROR_MORE_DATA 表示提供的缓冲区太小，应在调整缓冲区大小后再次调用
		requiredSliceLen := bytesNeeded / uint32(unsafe.Sizeof(windows.ENUM_SERVICE_STATUS{}))
		if bytesNeeded%uint32(unsafe.Sizeof(windows.ENUM_SERVICE_STATUS{})) != 0 {
			requiredSliceLen += 1
		}
		services = make([]windows.ENUM_SERVICE_STATUS, requiredSliceLen)
	}
	if returnedServiceCount == 0 {
		return nil, nil
	}

// EnumDependentServices 所修改的切片长度可能大于 returnedServiceCount，对于超出部分的任何元素应忽略。
	var dependents []string
	for i := 0; i < int(returnedServiceCount); i++ {
		dependents = append(dependents, windows.UTF16PtrToString(services[i].ServiceName))
	}
	return dependents, nil
}
