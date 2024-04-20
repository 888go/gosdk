
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
// Service is used to access Windows service.
<原文结束>

# <翻译开始>
// Service 用于访问 Windows 服务。
# <翻译结束>


<原文开始>
// Delete marks service s for deletion from the service control manager database.
<原文结束>

# <翻译开始>
// Delete 标记服务 s 以便从服务控制管理器数据库中删除。
# <翻译结束>


<原文开始>
// Close relinquish access to the service s.
<原文结束>

# <翻译开始>
// Close 放弃对服务 s 的访问。
# <翻译结束>


<原文开始>
// Start starts service s.
// args will be passed to svc.Handler.Execute.
<原文结束>

# <翻译开始>
// Start 启动服务 s。
// args 将被传递给 svc.Handler.Execute。
# <翻译结束>


<原文开始>
// Control sends state change request c to the service s. It returns the most
// recent status the service reported to the service control manager, and an
// error if the state change request was not accepted.
// Note that the returned service status is only set if the status change
// request succeeded, or if it failed with error ERROR_INVALID_SERVICE_CONTROL,
// ERROR_SERVICE_CANNOT_ACCEPT_CTRL, or ERROR_SERVICE_NOT_ACTIVE.
<原文结束>

# <翻译开始>
// Control 向服务 s 发送状态变更请求 c。它返回服务上报给服务控制管理器的最新状态，以及在状态变更请求未被接受时返回的错误。
// 注意，只有当状态变更请求成功，或者因错误 ERROR_INVALID_SERVICE_CONTROL、ERROR_SERVICE_CANNOT_ACCEPT_CTRL 或 ERROR_SERVICE_NOT_ACTIVE 而失败时，才会设置返回的服务状态。
# <翻译结束>


<原文开始>
// Query returns current status of service s.
<原文结束>

# <翻译开始>
// Query 返回服务 s 当前的状态
# <翻译结束>


<原文开始>
// ListDependentServices returns the names of the services dependent on service s, which match the given status.
<原文结束>

# <翻译开始>
// ListDependentServices 返回依赖于服务s且状态与给定状态匹配的服务名称。
# <翻译结束>


<原文开始>
// ERROR_MORE_DATA indicates the provided buffer was too small, run the call again after resizing the buffer
<原文结束>

# <翻译开始>
// ERROR_MORE_DATA 表示提供的缓冲区太小，应在调整缓冲区大小后再次调用
# <翻译结束>


<原文开始>
	// The slice mutated by EnumDependentServices may have a length greater than returnedServiceCount, any elements
	// past that should be ignored.
<原文结束>

# <翻译开始>
// EnumDependentServices 所修改的切片长度可能大于 returnedServiceCount，对于超出部分的任何元素应忽略。
# <翻译结束>

