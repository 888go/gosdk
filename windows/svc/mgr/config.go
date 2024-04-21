// 版权所有 2012 The Go 作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

//go:build windows

package mgr

import (
	"syscall"
	"unsafe"

	"github.com/888go/gosdk/windows"
)

const (
	// Service start types.
	StartManual    = windows.SERVICE_DEMAND_START // 该服务必须手动启动
	StartAutomatic = windows.SERVICE_AUTO_START   // 该服务会在计算机重启时自行启动
	StartDisabled  = windows.SERVICE_DISABLED     // 服务无法启动

// 该服务启动失败时的错误严重性及已采取的行动
	ErrorCritical = windows.SERVICE_ERROR_CRITICAL
	ErrorIgnore   = windows.SERVICE_ERROR_IGNORE
	ErrorNormal   = windows.SERVICE_ERROR_NORMAL
	ErrorSevere   = windows.SERVICE_ERROR_SEVERE
)

// TODO(brainman)：windows.QueryServiceConfig并未返回密码，尚不清楚如何获取。

type Config struct {
	ServiceType      uint32
	StartType        uint32
	ErrorControl     uint32
	BinaryPathName   string // 完全限定的服务二进制文件路径，也可以包含用于自动启动服务的参数
	LoadOrderGroup   string
	TagId            uint32
	Dependencies     []string
	ServiceStartName string // 服务应以哪个账户的名义运行
	DisplayName      string
	Password         string
	Description      string
	SidType          uint32 // SERVICE_SID_TYPE中的一个，用于服务的sid类型
	DelayedAutoStart bool   // 该服务在其他自动启动的服务启动后，加上短暂延迟即开始
}

func toStringSlice(ps *uint16) []string {
	r := make([]string, 0)
	p := unsafe.Pointer(ps)

	for {
		s := windows.UTF16PtrToString((*uint16)(p))
		if len(s) == 0 {
			break
		}

		r = append(r, s)
		offset := unsafe.Sizeof(uint16(0)) * (uintptr)(len(s)+1)
		p = unsafe.Pointer(uintptr(p) + offset)
	}

	return r
}

// Config 获取服务 s 的配置参数。
func (s *Service) Config() (Config, error) {
	var p *windows.QUERY_SERVICE_CONFIG
	n := uint32(1024)
	for {
		b := make([]byte, n)
		p = (*windows.QUERY_SERVICE_CONFIG)(unsafe.Pointer(&b[0]))
		err := windows.QueryServiceConfig(s.Handle, p, n, &n)
		if err == nil {
			break
		}
		if err.(syscall.Errno) != syscall.ERROR_INSUFFICIENT_BUFFER {
			return Config{}, err
		}
		if n <= uint32(len(b)) {
			return Config{}, err
		}
	}

	b, err := s.queryServiceConfig2(windows.SERVICE_CONFIG_DESCRIPTION)
	if err != nil {
		return Config{}, err
	}
	p2 := (*windows.SERVICE_DESCRIPTION)(unsafe.Pointer(&b[0]))

	b, err = s.queryServiceConfig2(windows.SERVICE_CONFIG_DELAYED_AUTO_START_INFO)
	if err != nil {
		return Config{}, err
	}
	p3 := (*windows.SERVICE_DELAYED_AUTO_START_INFO)(unsafe.Pointer(&b[0]))
	delayedStart := false
	if p3.IsDelayedAutoStartUp != 0 {
		delayedStart = true
	}

	b, err = s.queryServiceConfig2(windows.SERVICE_CONFIG_SERVICE_SID_INFO)
	if err != nil {
		return Config{}, err
	}
	sidType := *(*uint32)(unsafe.Pointer(&b[0]))

	return Config{
		ServiceType:      p.ServiceType,
		StartType:        p.StartType,
		ErrorControl:     p.ErrorControl,
		BinaryPathName:   windows.UTF16PtrToString(p.BinaryPathName),
		LoadOrderGroup:   windows.UTF16PtrToString(p.LoadOrderGroup),
		TagId:            p.TagId,
		Dependencies:     toStringSlice(p.Dependencies),
		ServiceStartName: windows.UTF16PtrToString(p.ServiceStartName),
		DisplayName:      windows.UTF16PtrToString(p.DisplayName),
		Description:      windows.UTF16PtrToString(p2.Description),
		DelayedAutoStart: delayedStart,
		SidType:          sidType,
	}, nil
}

func updateDescription(handle windows.Handle, desc string) error {
	d := windows.SERVICE_DESCRIPTION{Description: toPtr(desc)}
	return windows.ChangeServiceConfig2(handle,
		windows.SERVICE_CONFIG_DESCRIPTION, (*byte)(unsafe.Pointer(&d)))
}

func updateSidType(handle windows.Handle, sidType uint32) error {
	return windows.ChangeServiceConfig2(handle, windows.SERVICE_CONFIG_SERVICE_SID_INFO, (*byte)(unsafe.Pointer(&sidType)))
}

func updateStartUp(handle windows.Handle, isDelayed bool) error {
	var d windows.SERVICE_DELAYED_AUTO_START_INFO
	if isDelayed {
		d.IsDelayedAutoStartUp = 1
	}
	return windows.ChangeServiceConfig2(handle,
		windows.SERVICE_CONFIG_DELAYED_AUTO_START_INFO, (*byte)(unsafe.Pointer(&d)))
}

// UpdateConfig 更新服务s的配置参数。
func (s *Service) UpdateConfig(c Config) error {
	err := windows.ChangeServiceConfig(s.Handle, c.ServiceType, c.StartType,
		c.ErrorControl, toPtr(c.BinaryPathName), toPtr(c.LoadOrderGroup),
		nil, toStringBlock(c.Dependencies), toPtr(c.ServiceStartName),
		toPtr(c.Password), toPtr(c.DisplayName))
	if err != nil {
		return err
	}
	err = updateSidType(s.Handle, c.SidType)
	if err != nil {
		return err
	}

	err = updateStartUp(s.Handle, c.DelayedAutoStart)
	if err != nil {
		return err
	}

	return updateDescription(s.Handle, c.Description)
}

// queryServiceConfig2 调用 Windows QueryServiceConfig2 函数，其中包含 infoLevel 参数，并返回所获取的服务配置信息。
func (s *Service) queryServiceConfig2(infoLevel uint32) ([]byte, error) {
	n := uint32(1024)
	for {
		b := make([]byte, n)
		err := windows.QueryServiceConfig2(s.Handle, infoLevel, &b[0], n, &n)
		if err == nil {
			return b, nil
		}
		if err.(syscall.Errno) != syscall.ERROR_INSUFFICIENT_BUFFER {
			return nil, err
		}
		if n <= uint32(len(b)) {
			return nil, err
		}
	}
}
