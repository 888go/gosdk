
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
//sys	RegisterEventSource(uncServerName *uint16, sourceName *uint16) (handle Handle, err error) [failretval==0] = advapi32.RegisterEventSourceW
//sys	DeregisterEventSource(handle Handle) (err error) = advapi32.DeregisterEventSource
//sys	ReportEvent(log Handle, etype uint16, category uint16, eventId uint32, usrSId uintptr, numStrings uint16, dataSize uint32, strings **uint16, rawData *byte) (err error) = advapi32.ReportEventW
<原文结束>

# <翻译开始>
//sys 注册事件源(uncServerName *uint16, sourceName *uint16) (handle Handle, err error) [失败返回值==0] = advapi32.RegisterEventSourceW
//sys 注销事件源(handle Handle) (err error) = advapi32.DeregisterEventSource
//sys 报告事件(log Handle, etype uint16, category uint16, eventId uint32, usrSId uintptr, numStrings uint16, dataSize uint32, strings **uint16, rawData *byte) (err error) = advapi32.ReportEventW
# <翻译结束>

