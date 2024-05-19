// 版权所有 2018 The Go Authors。保留所有权利。
// 使用本源代码须遵循 BSD 风格许可证，
// 该许可证可在 LICENSE 文件中找到。

//go:build windows

package mgr

import (
	"errors"
	"syscall"
	"time"
	"unsafe"

	"github.com/888go/gosdk/windows"
)

const (
	// 服务控制管理器可能执行的恢复操作。
	NoAction       = windows.SC_ACTION_NONE        // no action
	ComputerReboot = windows.SC_ACTION_REBOOT      // reboot the computer
	ServiceRestart = windows.SC_ACTION_RESTART     // restart the service
	RunCommand     = windows.SC_ACTION_RUN_COMMAND // run a command
)

// RecoveryAction 表示服务控制管理器在服务失败时可执行的操作。
// 当服务在未向服务控制器报告 SERVICE_STOPPED 状态的情况下终止时，视为该服务已失败。
type RecoveryAction struct {
	Type  int           // 可为 NoAction、ComputerReboot、ServiceRestart 或 RunCommand 中的一种
	Delay time.Duration // 执行指定操作前的等待时间
}

// SetRecoveryActions 设置当服务失败时，服务控制器执行的操作以及在无故障情况下将服务失败计数器重置为零的时间间隔（以秒为单位）。
// 指定 INFINITE 表示服务失败计数器应永不重置。

// ff:
// resetPeriod:
// recoveryActions:
func (s *Service) SetRecoveryActions(recoveryActions []RecoveryAction, resetPeriod uint32) error {
	if recoveryActions == nil {
		return errors.New("recoveryActions cannot be nil")
	}
	actions := []windows.SC_ACTION{}
	for _, a := range recoveryActions {
		action := windows.SC_ACTION{
			Type:  uint32(a.Type),
			Delay: uint32(a.Delay.Nanoseconds() / 1000000),
		}
		actions = append(actions, action)
	}
	rActions := windows.SERVICE_FAILURE_ACTIONS{
		ActionsCount: uint32(len(actions)),
		Actions:      &actions[0],
		ResetPeriod:  resetPeriod,
	}
	return windows.ChangeServiceConfig2(s.Handle, windows.SERVICE_CONFIG_FAILURE_ACTIONS, (*byte)(unsafe.Pointer(&rActions)))
}

// RecoveryActions 返回服务控制器在服务失败时执行的操作。
// 服务控制管理器统计自系统启动以来服务 s 失败的次数。
// 如果该服务在 ResetPeriod 秒内未发生故障，计数将重置为 0。
// 当服务第 N 次失败时，服务控制器执行返回切片中元素 [N-1] 指定的动作。
// 若 N 大于切片长度，服务控制器将重复执行切片中的最后一个动作。

// ff:
func (s *Service) RecoveryActions() ([]RecoveryAction, error) {
	b, err := s.queryServiceConfig2(windows.SERVICE_CONFIG_FAILURE_ACTIONS)
	if err != nil {
		return nil, err
	}
	p := (*windows.SERVICE_FAILURE_ACTIONS)(unsafe.Pointer(&b[0]))
	if p.Actions == nil {
		return nil, err
	}

	actions := unsafe.Slice(p.Actions, int(p.ActionsCount))
	var recoveryActions []RecoveryAction
	for _, action := range actions {
		recoveryActions = append(recoveryActions, RecoveryAction{Type: int(action.Type), Delay: time.Duration(action.Delay) * time.Millisecond})
	}
	return recoveryActions, nil
}

// 重置恢复动作：删除复位周期与失败动作数组

// ff:
func (s *Service) ResetRecoveryActions() error {
	actions := make([]windows.SC_ACTION, 1)
	rActions := windows.SERVICE_FAILURE_ACTIONS{
		Actions: &actions[0],
	}
	return windows.ChangeServiceConfig2(s.Handle, windows.SERVICE_CONFIG_FAILURE_ACTIONS, (*byte)(unsafe.Pointer(&rActions)))
}

// ResetPeriod 是在无故障情况下将服务失败计数重置为零的时间间隔，以秒为单位。

// ff:
func (s *Service) ResetPeriod() (uint32, error) {
	b, err := s.queryServiceConfig2(windows.SERVICE_CONFIG_FAILURE_ACTIONS)
	if err != nil {
		return 0, err
	}
	p := (*windows.SERVICE_FAILURE_ACTIONS)(unsafe.Pointer(&b[0]))
	return p.ResetPeriod, nil
}

// SetRebootMessage 设置服务s的重启消息。
// 若msg为"", 则删除重启消息，且不广播任何消息。

// ff:
// msg:
func (s *Service) SetRebootMessage(msg string) error {
	rActions := windows.SERVICE_FAILURE_ACTIONS{
		RebootMsg: syscall.StringToUTF16Ptr(msg),
	}
	return windows.ChangeServiceConfig2(s.Handle, windows.SERVICE_CONFIG_FAILURE_ACTIONS, (*byte)(unsafe.Pointer(&rActions)))
}

// RebootMessage 是在响应 ComputerReboot 服务控制器操作重启前，向服务器用户广播的消息。

// ff:
func (s *Service) RebootMessage() (string, error) {
	b, err := s.queryServiceConfig2(windows.SERVICE_CONFIG_FAILURE_ACTIONS)
	if err != nil {
		return "", err
	}
	p := (*windows.SERVICE_FAILURE_ACTIONS)(unsafe.Pointer(&b[0]))
	return windows.UTF16PtrToString(p.RebootMsg), nil
}

// SetRecoveryCommand 用于设置在响应 RunCommand 服务控制器操作时执行的进程命令行。
// 若 cmd 为空字符串（""），则删除该命令，当服务失败时不再运行任何程序。

// ff:
// cmd:
func (s *Service) SetRecoveryCommand(cmd string) error {
	rActions := windows.SERVICE_FAILURE_ACTIONS{
		Command: syscall.StringToUTF16Ptr(cmd),
	}
	return windows.ChangeServiceConfig2(s.Handle, windows.SERVICE_CONFIG_FAILURE_ACTIONS, (*byte)(unsafe.Pointer(&rActions)))
}

// RecoveryCommand 是对服务控制器“RunCommand”操作做出响应时要执行的进程命令行。此进程以与服务相同的身份运行。

// ff:
func (s *Service) RecoveryCommand() (string, error) {
	b, err := s.queryServiceConfig2(windows.SERVICE_CONFIG_FAILURE_ACTIONS)
	if err != nil {
		return "", err
	}
	p := (*windows.SERVICE_FAILURE_ACTIONS)(unsafe.Pointer(&b[0]))
	return windows.UTF16PtrToString(p.Command), nil
}

// SetRecoveryActionsOnNonCrashFailures 设置失败动作标志。如果该标志设置为false，
// 则仅当服务终止时未报告SERVICE_STOPPED状态，才会执行恢复动作。
// 如果该标志设置为true，则当服务以非零退出代码停止时，也会执行恢复动作。

// ff:
// flag:
func (s *Service) SetRecoveryActionsOnNonCrashFailures(flag bool) error {
	var setting windows.SERVICE_FAILURE_ACTIONS_FLAG
	if flag {
		setting.FailureActionsOnNonCrashFailures = 1
	}
	return windows.ChangeServiceConfig2(s.Handle, windows.SERVICE_CONFIG_FAILURE_ACTIONS_FLAG, (*byte)(unsafe.Pointer(&setting)))
}

// RecoveryActionsOnNonCrashFailures 返回当前的失败操作标志值。如果该标志设为false，只有当服务终止且未报告SERVICE_STOPPED状态时，才会执行恢复动作。如果该标志设为true，只要服务停止时退出代码不为0，也会执行恢复动作。

// ff:
func (s *Service) RecoveryActionsOnNonCrashFailures() (bool, error) {
	b, err := s.queryServiceConfig2(windows.SERVICE_CONFIG_FAILURE_ACTIONS_FLAG)
	if err != nil {
		return false, err
	}
	p := (*windows.SERVICE_FAILURE_ACTIONS_FLAG)(unsafe.Pointer(&b[0]))
	return p.FailureActionsOnNonCrashFailures != 0, nil
}
