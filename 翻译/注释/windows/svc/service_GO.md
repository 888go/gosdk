
<原文开始>
// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// 版权所有 ? 2012 The Go Authors。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。
# <翻译结束>


<原文开始>
// Package svc provides everything required to build Windows service.
<原文结束>

# <翻译开始>
// 包svc提供了构建Windows服务所需的一切。
# <翻译结束>


<原文开始>
// State describes service execution state (Stopped, Running and so on).
<原文结束>

# <翻译开始>
// State 描述服务执行状态（停止、运行等）。
# <翻译结束>


<原文开始>
// Cmd represents service state change request. It is sent to a service
// by the service manager, and should be actioned upon by the service.
<原文结束>

# <翻译开始>
// Cmd 代表服务状态变更请求。它由服务管理器发送至服务，并应由服务进行处理。
# <翻译结束>


<原文开始>
// Accepted is used to describe commands accepted by the service.
// Note that Interrogate is always accepted.
<原文结束>

# <翻译开始>
// Accepted 用于描述服务所接受的命令。
// 注意，Interrogate 命令始终被接受。
# <翻译结束>


<原文开始>
// ActivityStatus allows for services to be selected based on active and inactive categories of service state.
<原文结束>

# <翻译开始>
// ActivityStatus 允许根据服务状态的活跃和非活跃类别来选择服务。
# <翻译结束>


<原文开始>
// Status combines State and Accepted commands to fully describe running service.
<原文结束>

# <翻译开始>
// Status 结合了 State 和 Accepted 命令，以全面描述运行中的服务。
# <翻译结束>


<原文开始>
// used to report progress during a lengthy operation
<原文结束>

# <翻译开始>
// 用于在长时间操作期间报告进度
# <翻译结束>


<原文开始>
// estimated time required for a pending operation, in milliseconds
<原文结束>

# <翻译开始>
// 估计的待处理操作所需时间，以毫秒为单位
# <翻译结束>


<原文开始>
// if the service is running, the process identifier of it, and otherwise zero
<原文结束>

# <翻译开始>
// 如果服务正在运行，则返回其进程标识符，否则返回零
# <翻译结束>


<原文开始>
// set if the service has exited with a win32 exit code
<原文结束>

# <翻译开始>
// 设置服务是否已使用 Win32 退出代码退出
# <翻译结束>


<原文开始>
// set if the service has exited with a service-specific exit code
<原文结束>

# <翻译开始>
// 设置标志，表明服务已使用特定于服务的退出代码退出
# <翻译结束>


<原文开始>
// StartReason is the reason that the service was started.
<原文结束>

# <翻译开始>
// StartReason 是服务启动的原因。
# <翻译结束>


<原文开始>
// ChangeRequest is sent to the service Handler to request service status change.
<原文结束>

# <翻译开始>
// ChangeRequest 用于发送至服务Handler，以请求更改服务状态。
# <翻译结束>


<原文开始>
// Handler is the interface that must be implemented to build Windows service.
<原文结束>

# <翻译开始>
// Handler 是必须实现以构建 Windows 服务的接口。
# <翻译结束>


<原文开始>
	// Execute will be called by the package code at the start of
	// the service, and the service will exit once Execute completes.
	// Inside Execute you must read service change requests from r and
	// act accordingly. You must keep service control manager up to date
	// about state of your service by writing into s as required.
	// args contains service name followed by argument strings passed
	// to the service.
	// You can provide service exit code in exitCode return parameter,
	// with 0 being "no error". You can also indicate if exit code,
	// if any, is service specific or not by using svcSpecificEC
	// parameter.
<原文结束>

# <翻译开始>
// Execute 将在服务开始时由包代码调用，且一旦 Execute 执行完毕，服务将退出。
// 在 Execute 中，您必须从 r 读取服务变更请求并据此采取相应行动。您必须通过在必要时向 s 写入信息，保持服务控制管理器了解您服务的状态。
// args 包含服务名以及传递给服务的参数字符串。
// 您可以通过 exitCode 返回参数提供服务退出码，其中 0 表示“无错误”。您也可以使用 svcSpecificEC 参数指示（如果有）退出码是否特定于服务。
# <翻译结束>


<原文开始>
// service provides access to windows service api.
<原文结束>

# <翻译开始>
// service 提供对 Windows 服务 API 的访问。
# <翻译结束>


<原文开始>
// Set context to 123456 to test issue #25660.
<原文结束>

# <翻译开始>
// 将上下文设置为123456以测试问题#25660
# <翻译结束>


<原文开始>
// This is, unfortunately, a global, which means only one service per process.
<原文结束>

# <翻译开始>
// 很遗憾，这是一个全局变量，这意味着每个进程只能有一个服务。
# <翻译结束>


<原文开始>
// serviceMain is the entry point called by the service manager, registered earlier by
// the call to StartServiceCtrlDispatcher.
<原文结束>

# <翻译开始>
// serviceMain 是服务管理器调用的入口点，此前已通过调用 StartServiceCtrlDispatcher 注册。
# <翻译结束>


<原文开始>
// Run executes service name by calling appropriate handler function.
<原文结束>

# <翻译开始>
// Run 通过调用相应的处理函数来执行服务名。
# <翻译结束>


<原文开始>
// StatusHandle returns service status handle. It is safe to call this function
// from inside the Handler.Execute because then it is guaranteed to be set.
<原文结束>

# <翻译开始>
// StatusHandle 返回服务状态句柄。在 Handler.Execute 内部调用此函数是安全的，
// 因为此时可以确保其已被设置。
# <翻译结束>


<原文开始>
// DynamicStartReason returns the reason why the service was started. It is safe
// to call this function from inside the Handler.Execute because then it is
// guaranteed to be set.
<原文结束>

# <翻译开始>
// DynamicStartReason 返回服务启动的原因。在 Handler.Execute 内部调用此函数是安全的，
// 因为此时可以确保其已设置。
# <翻译结束>

