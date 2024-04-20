
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
// Package eventlog implements access to Windows event log.
<原文结束>

# <翻译开始>
// package eventlog 实现了对 Windows 事件日志的访问功能。
# <翻译结束>


<原文开始>
// Log provides access to the system log.
<原文结束>

# <翻译开始>
// Log 提供对系统日志的访问。
# <翻译结束>


<原文开始>
// Open retrieves a handle to the specified event log.
<原文结束>

# <翻译开始>
// Open 获取指定事件日志的句柄。
# <翻译结束>


<原文开始>
// OpenRemote does the same as Open, but on different computer host.
<原文结束>

# <翻译开始>
// OpenRemote 与 Open 执行相同的操作，但针对的是另一台计算机主机。
# <翻译结束>


<原文开始>
// Close closes event log l.
<原文结束>

# <翻译开始>
// Close 关闭事件日志 l。
# <翻译结束>


<原文开始>
// Info writes an information event msg with event id eid to the end of event log l.
// When EventCreate.exe is used, eid must be between 1 and 1000.
<原文结束>

# <翻译开始>
// Info 将带有事件ID eid 的信息事件msg写入事件日志l的末尾。
// 使用EventCreate.exe时，eid必须介于1和1000之间。
# <翻译结束>


<原文开始>
// Warning writes an warning event msg with event id eid to the end of event log l.
// When EventCreate.exe is used, eid must be between 1 and 1000.
<原文结束>

# <翻译开始>
// Warning 在事件日志 l 的末尾写入一个带有事件ID eid 的警告事件消息。
// 当使用 EventCreate.exe 时，eid 必须介于 1 和 1000 之间。
# <翻译结束>


<原文开始>
// Error writes an error event msg with event id eid to the end of event log l.
// When EventCreate.exe is used, eid must be between 1 and 1000.
<原文结束>

# <翻译开始>
// Error 函数向事件日志 l 的末尾写入一个具有事件ID eid 的错误事件消息。
// 当使用 EventCreate.exe 时，eid 必须在 1 到 1000 之间。
# <翻译结束>

