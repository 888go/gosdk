// 版权所有 2012 The Go 作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

//go:build windows

// Package svc 提供了构建 Windows 服务所需的一切。
package svc

import (
	"errors"
	"sync"
	"unsafe"

	"github.com/888go/gosdk/windows"
)

// State 描述服务执行状态（停止、运行等）。
type State uint32

const (
	Stopped         = State(windows.SERVICE_STOPPED)
	StartPending    = State(windows.SERVICE_START_PENDING)
	StopPending     = State(windows.SERVICE_STOP_PENDING)
	Running         = State(windows.SERVICE_RUNNING)
	ContinuePending = State(windows.SERVICE_CONTINUE_PENDING)
	PausePending    = State(windows.SERVICE_PAUSE_PENDING)
	Paused          = State(windows.SERVICE_PAUSED)
)

// Cmd 表示服务状态变更请求。它由服务管理器发送给服务，并应由服务执行相应的操作。
type Cmd uint32

const (
	Stop                  = Cmd(windows.SERVICE_CONTROL_STOP)
	Pause                 = Cmd(windows.SERVICE_CONTROL_PAUSE)
	Continue              = Cmd(windows.SERVICE_CONTROL_CONTINUE)
	Interrogate           = Cmd(windows.SERVICE_CONTROL_INTERROGATE)
	Shutdown              = Cmd(windows.SERVICE_CONTROL_SHUTDOWN)
	ParamChange           = Cmd(windows.SERVICE_CONTROL_PARAMCHANGE)
	NetBindAdd            = Cmd(windows.SERVICE_CONTROL_NETBINDADD)
	NetBindRemove         = Cmd(windows.SERVICE_CONTROL_NETBINDREMOVE)
	NetBindEnable         = Cmd(windows.SERVICE_CONTROL_NETBINDENABLE)
	NetBindDisable        = Cmd(windows.SERVICE_CONTROL_NETBINDDISABLE)
	DeviceEvent           = Cmd(windows.SERVICE_CONTROL_DEVICEEVENT)
	HardwareProfileChange = Cmd(windows.SERVICE_CONTROL_HARDWAREPROFILECHANGE)
	PowerEvent            = Cmd(windows.SERVICE_CONTROL_POWEREVENT)
	SessionChange         = Cmd(windows.SERVICE_CONTROL_SESSIONCHANGE)
	PreShutdown           = Cmd(windows.SERVICE_CONTROL_PRESHUTDOWN)
)

// Accepted 用于描述服务所接受的命令。
// 注意，Interrogate 命令总是被接受。
type Accepted uint32

const (
	AcceptStop                  = Accepted(windows.SERVICE_ACCEPT_STOP)
	AcceptShutdown              = Accepted(windows.SERVICE_ACCEPT_SHUTDOWN)
	AcceptPauseAndContinue      = Accepted(windows.SERVICE_ACCEPT_PAUSE_CONTINUE)
	AcceptParamChange           = Accepted(windows.SERVICE_ACCEPT_PARAMCHANGE)
	AcceptNetBindChange         = Accepted(windows.SERVICE_ACCEPT_NETBINDCHANGE)
	AcceptHardwareProfileChange = Accepted(windows.SERVICE_ACCEPT_HARDWAREPROFILECHANGE)
	AcceptPowerEvent            = Accepted(windows.SERVICE_ACCEPT_POWEREVENT)
	AcceptSessionChange         = Accepted(windows.SERVICE_ACCEPT_SESSIONCHANGE)
	AcceptPreShutdown           = Accepted(windows.SERVICE_ACCEPT_PRESHUTDOWN)
)

// ActivityStatus 允许根据服务的活跃和非活跃状态类别来选择服务。
type ActivityStatus uint32

const (
	Active      = ActivityStatus(windows.SERVICE_ACTIVE)
	Inactive    = ActivityStatus(windows.SERVICE_INACTIVE)
	AnyActivity = ActivityStatus(windows.SERVICE_STATE_ALL)
)

// Status 结合了 State 和 Accepted commands，以全面描述运行中的服务。
type Status struct {
	State                   State
	Accepts                 Accepted
	CheckPoint              uint32 // 用于在长时间操作过程中报告进度
	WaitHint                uint32 // 估计的待处理操作所需时间，以毫秒为单位
	ProcessId               uint32 // 如果服务正在运行，则为该服务的进程标识符，否则为零
	Win32ExitCode           uint32 // 设置该服务是否已使用Win32退出代码退出
	ServiceSpecificExitCode uint32 // 设置该标志，表示服务已使用特定于服务的退出代码终止
}

// StartReason 是服务启动的原因。
type StartReason uint32

const (
	StartReasonDemand           = StartReason(windows.SERVICE_START_REASON_DEMAND)
	StartReasonAuto             = StartReason(windows.SERVICE_START_REASON_AUTO)
	StartReasonTrigger          = StartReason(windows.SERVICE_START_REASON_TRIGGER)
	StartReasonRestartOnFailure = StartReason(windows.SERVICE_START_REASON_RESTART_ON_FAILURE)
	StartReasonDelayedAuto      = StartReason(windows.SERVICE_START_REASON_DELAYEDAUTO)
)

// ChangeRequest 用于发送至服务处理器，以请求服务状态变更。
type ChangeRequest struct {
	Cmd           Cmd
	EventType     uint32
	EventData     uintptr
	CurrentStatus Status
	Context       uintptr
}

// Handler 是必须实现以构建 Windows 服务的接口。
type Handler interface {
// Execute 将在服务开始时由包代码调用，且一旦 Execute 执行完毕，服务就会退出。
// 在 Execute 内部，您必须从 r 读取服务更改请求并据此采取相应行动。您必须通过在必要时向 s 中写入信息，
// 来保持服务控制管理器与您服务状态的同步。
// args 包含服务名以及传递给服务的参数字符串。
// 您可以通过 exitCode 返回参数提供服务退出码，其中 0 表示“无错误”。您还可以使用 svcSpecificEC 参数
// 指示（如果有）退出码是否特定于服务。
	Execute(args []string, r <-chan ChangeRequest, s chan<- Status) (svcSpecificEC bool, exitCode uint32)
}

type ctlEvent struct {
	cmd       Cmd
	eventType uint32
	eventData uintptr
	context   uintptr
	errno     uint32
}

// service 提供对 Windows 服务 API 的访问。
type service struct {
	name    string
	h       windows.Handle
	c       chan ctlEvent
	handler Handler
}

type exitCode struct {
	isSvcSpecific bool
	errno         uint32
}

func (s *service) updateStatus(status *Status, ec *exitCode) error {
	if s.h == 0 {
		return errors.New("updateStatus with no service status handle")
	}
	var t windows.SERVICE_STATUS
	t.ServiceType = windows.SERVICE_WIN32_OWN_PROCESS
	t.CurrentState = uint32(status.State)
	if status.Accepts&AcceptStop != 0 {
		t.ControlsAccepted |= windows.SERVICE_ACCEPT_STOP
	}
	if status.Accepts&AcceptShutdown != 0 {
		t.ControlsAccepted |= windows.SERVICE_ACCEPT_SHUTDOWN
	}
	if status.Accepts&AcceptPauseAndContinue != 0 {
		t.ControlsAccepted |= windows.SERVICE_ACCEPT_PAUSE_CONTINUE
	}
	if status.Accepts&AcceptParamChange != 0 {
		t.ControlsAccepted |= windows.SERVICE_ACCEPT_PARAMCHANGE
	}
	if status.Accepts&AcceptNetBindChange != 0 {
		t.ControlsAccepted |= windows.SERVICE_ACCEPT_NETBINDCHANGE
	}
	if status.Accepts&AcceptHardwareProfileChange != 0 {
		t.ControlsAccepted |= windows.SERVICE_ACCEPT_HARDWAREPROFILECHANGE
	}
	if status.Accepts&AcceptPowerEvent != 0 {
		t.ControlsAccepted |= windows.SERVICE_ACCEPT_POWEREVENT
	}
	if status.Accepts&AcceptSessionChange != 0 {
		t.ControlsAccepted |= windows.SERVICE_ACCEPT_SESSIONCHANGE
	}
	if status.Accepts&AcceptPreShutdown != 0 {
		t.ControlsAccepted |= windows.SERVICE_ACCEPT_PRESHUTDOWN
	}
	if ec.errno == 0 {
		t.Win32ExitCode = windows.NO_ERROR
		t.ServiceSpecificExitCode = windows.NO_ERROR
	} else if ec.isSvcSpecific {
		t.Win32ExitCode = uint32(windows.ERROR_SERVICE_SPECIFIC_ERROR)
		t.ServiceSpecificExitCode = ec.errno
	} else {
		t.Win32ExitCode = ec.errno
		t.ServiceSpecificExitCode = windows.NO_ERROR
	}
	t.CheckPoint = status.CheckPoint
	t.WaitHint = status.WaitHint
	return windows.SetServiceStatus(s.h, &t)
}

var (
	initCallbacks       sync.Once
	ctlHandlerCallback  uintptr
	serviceMainCallback uintptr
)

func ctlHandler(ctl, evtype, evdata, context uintptr) uintptr {
	s := (*service)(unsafe.Pointer(context))
	e := ctlEvent{cmd: Cmd(ctl), eventType: uint32(evtype), eventData: evdata, context: 123456} // 将上下文设置为 123456 以测试问题 #25660
	s.c <- e
	return 0
}

var theService service // 很遗憾，这是一个全局变量，这意味着每个进程只能有一个服务。

// serviceMain 是服务管理器调用的入口点，早些时候由对 StartServiceCtrlDispatcher 的调用进行注册。
func serviceMain(argc uint32, argv **uint16) uintptr {
	handle, err := windows.RegisterServiceCtrlHandlerEx(windows.StringToUTF16Ptr(theService.name), ctlHandlerCallback, uintptr(unsafe.Pointer(&theService)))
	if sysErr, ok := err.(windows.Errno); ok {
		return uintptr(sysErr)
	} else if err != nil {
		return uintptr(windows.ERROR_UNKNOWN_EXCEPTION)
	}
	theService.h = handle
	defer func() {
		theService.h = 0
	}()
	args16 := unsafe.Slice(argv, int(argc))

	args := make([]string, len(args16))
	for i, a := range args16 {
		args[i] = windows.UTF16PtrToString(a)
	}

	cmdsToHandler := make(chan ChangeRequest)
	changesFromHandler := make(chan Status)
	exitFromHandler := make(chan exitCode)

	go func() {
		ss, errno := theService.handler.Execute(args, cmdsToHandler, changesFromHandler)
		exitFromHandler <- exitCode{ss, errno}
	}()

	ec := exitCode{isSvcSpecific: true, errno: 0}
	outcr := ChangeRequest{
		CurrentStatus: Status{State: Stopped},
	}
	var outch chan ChangeRequest
	inch := theService.c
loop:
	for {
		select {
		case r := <-inch:
			if r.errno != 0 {
				ec.errno = r.errno
				break loop
			}
			inch = nil
			outch = cmdsToHandler
			outcr.Cmd = r.cmd
			outcr.EventType = r.eventType
			outcr.EventData = r.eventData
			outcr.Context = r.context
		case outch <- outcr:
			inch = theService.c
			outch = nil
		case c := <-changesFromHandler:
			err := theService.updateStatus(&c, &ec)
			if err != nil {
				ec.errno = uint32(windows.ERROR_EXCEPTION_IN_SERVICE)
				if err2, ok := err.(windows.Errno); ok {
					ec.errno = uint32(err2)
				}
				break loop
			}
			outcr.CurrentStatus = c
		case ec = <-exitFromHandler:
			break loop
		}
	}

	theService.updateStatus(&Status{State: Stopped}, &ec)

	return windows.NO_ERROR
}

// Run 通过调用相应的处理函数来执行服务名。

// ff:
// handler:
// name:
func Run(name string, handler Handler) error {
	initCallbacks.Do(func() {
		ctlHandlerCallback = windows.NewCallback(ctlHandler)
		serviceMainCallback = windows.NewCallback(serviceMain)
	})
	theService.name = name
	theService.handler = handler
	theService.c = make(chan ctlEvent)
	t := []windows.SERVICE_TABLE_ENTRY{
		{ServiceName: windows.StringToUTF16Ptr(theService.name), ServiceProc: serviceMainCallback},
		{ServiceName: nil, ServiceProc: 0},
	}
	return windows.StartServiceCtrlDispatcher(&t[0])
}

// StatusHandle 返回服务状态句柄。在 Handler.Execute 内部调用此函数是安全的，
// 因为此时可以确保其已被设置。

// ff:
func StatusHandle() windows.Handle {
	return theService.h
}

// DynamicStartReason 返回服务启动的原因。从 Handler.Execute 内部调用此函数是安全的，
// 因为此时可以确保其已被设置。

// ff:
// StartReason:
func DynamicStartReason() (StartReason, error) {
	var allocReason *uint32
	err := windows.QueryServiceDynamicInformation(theService.h, windows.SERVICE_DYNAMIC_INFORMATION_LEVEL_START_REASON, unsafe.Pointer(&allocReason))
	if err != nil {
		return 0, err
	}
	reason := StartReason(*allocReason)
	windows.LocalFree(windows.Handle(unsafe.Pointer(allocReason)))
	return reason, nil
}
