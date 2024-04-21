
<原文开始>
// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// 版权所有 2012 The Go 作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。
# <翻译结束>


<原文开始>
// Package mgr can be used to manage Windows service programs.
// It can be used to install and remove them. It can also start,
// stop and pause them. The package can query / change current
// service state and config parameters.
<原文结束>

# <翻译开始>
// Package mgr 用于管理 Windows 服务程序。
// 可以使用它来安装和卸载服务，以及启动、停止和暂停服务。该包可以查询/更改当前服务状态和配置参数。
# <翻译结束>


<原文开始>
// Mgr is used to manage Windows service.
<原文结束>

# <翻译开始>
// Mgr 用于管理 Windows 服务。
# <翻译结束>


<原文开始>
// Connect establishes a connection to the service control manager.
<原文结束>

# <翻译开始>
// Connect 建立与服务控制管理器的连接
# <翻译结束>


<原文开始>
// ConnectRemote establishes a connection to the
// service control manager on computer named host.
<原文结束>

# <翻译开始>
// ConnectRemote 与名为 host 的计算机上的
// 服务控制管理器建立连接。
# <翻译结束>


<原文开始>
// Disconnect closes connection to the service control manager m.
<原文结束>

# <翻译开始>
// Disconnect 关闭与服务控制管理器 m 的连接
# <翻译结束>


<原文开始>
// Whether the SCM has been locked.
<原文结束>

# <翻译开始>
// 是否已锁定SCM
# <翻译结束>


<原文开始>
// For how long the SCM has been locked.
<原文结束>

# <翻译开始>
// SCM被锁定的时长
# <翻译结束>


<原文开始>
// The name of the user who has locked the SCM.
<原文结束>

# <翻译开始>
// 锁定SCM的用户名称
# <翻译结束>


<原文开始>
// LockStatus returns whether the service control manager is locked by
// the system, for how long, and by whom. A locked SCM indicates that
// most service actions will block until the system unlocks the SCM.
<原文结束>

# <翻译开始>
// LockStatus 返回服务控制管理器是否被系统锁定，以及锁定时长和锁定者。SCM（服务控制管理器）被锁定表明大多数服务操作将阻塞，直至系统解锁SCM。
# <翻译结束>


<原文开始>
// toStringBlock terminates strings in ss with 0, and then
// concatenates them together. It also adds extra 0 at the end.
<原文结束>

# <翻译开始>
// toStringBlock 将 ss 中的字符串以 0 结尾，并将它们拼接在一起。同时在末尾额外添加一个 0。
# <翻译结束>


<原文开始>
// CreateService installs new service name on the system.
// The service will be executed by running exepath binary.
// Use config c to specify service parameters.
// Any args will be passed as command-line arguments when
// the service is started; these arguments are distinct from
// the arguments passed to Service.Start or via the "Start
// parameters" field in the service's Properties dialog box.
<原文结束>

# <翻译开始>
// 创建服务：在系统上安装新服务名称。
// 该服务将通过运行exepath指定的二进制文件来执行。
// 使用配置c来指定服务参数。
// 任何args参数将在服务启动时作为命令行参数传递；
// 这些参数与通过Service.Start或通过服务“属性”对话框中的“启动参数”字段传递的参数不同。
# <翻译结束>


<原文开始>
// OpenService retrieves access to service name, so it can
// be interrogated and controlled.
<原文结束>

# <翻译开始>
// OpenService 获取对服务名name的访问权限，以便对其进行查询和控制。
# <翻译结束>


<原文开始>
// ListServices enumerates services in the specified
// service control manager database m.
// If the caller does not have the SERVICE_QUERY_STATUS
// access right to a service, the service is silently
// omitted from the list of services returned.
<原文结束>

# <翻译开始>
// ListServices 在指定的服务控制管理器数据库 m 中枚举服务。
// 如果调用者对某个服务不具备 SERVICE_QUERY_STATUS 访问权限，
// 则该服务将被静默地从返回的服务列表中省略。
# <翻译结束>

