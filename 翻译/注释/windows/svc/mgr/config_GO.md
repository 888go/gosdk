
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
// the service must be started manually
<原文结束>

# <翻译开始>
// 该服务必须手动启动
# <翻译结束>


<原文开始>
// the service will start by itself whenever the computer reboots
<原文结束>

# <翻译开始>
// 该服务会在计算机重启时自动启动
# <翻译结束>


<原文开始>
// the service cannot be started
<原文结束>

# <翻译开始>
// 服务无法启动
# <翻译结束>


<原文开始>
	// The severity of the error, and action taken,
	// if this service fails to start.
<原文结束>

# <翻译开始>
// 该服务启动失败时的错误严重性及已采取的行动
# <翻译结束>


<原文开始>
// TODO(brainman): Password is not returned by windows.QueryServiceConfig, not sure how to get it.
<原文结束>

# <翻译开始>
// TODO(brainman)：windows.QueryServiceConfig并未返回密码，尚不清楚如何获取。
# <翻译结束>


<原文开始>
// fully qualified path to the service binary file, can also include arguments for an auto-start service
<原文结束>

# <翻译开始>
// 完全限定的服务二进制文件路径，也可包含用于自动启动服务的参数
# <翻译结束>


<原文开始>
// name of the account under which the service should run
<原文结束>

# <翻译开始>
// 服务应以哪个账户的名称运行
# <翻译结束>


<原文开始>
// one of SERVICE_SID_TYPE, the type of sid to use for the service
<原文结束>

# <翻译开始>
// SERVICE_SID_TYPE中的一个，用于指定服务的sid类型
# <翻译结束>


<原文开始>
// the service is started after other auto-start services are started plus a short delay
<原文结束>

# <翻译开始>
// 该服务在其他自动启动的服务启动后，加上短暂延迟再开始
# <翻译结束>


<原文开始>
// Config retrieves service s configuration paramteres.
<原文结束>

# <翻译开始>
// Config 获取服务 s 的配置参数。
# <翻译结束>


<原文开始>
// UpdateConfig updates service s configuration parameters.
<原文结束>

# <翻译开始>
// UpdateConfig 更新服务s的配置参数
# <翻译结束>


<原文开始>
// queryServiceConfig2 calls Windows QueryServiceConfig2 with infoLevel parameter and returns retrieved service configuration information.
<原文结束>

# <翻译开始>
// queryServiceConfig2 调用 Windows 的 QueryServiceConfig2 函数，并以 infoLevel 参数指定信息级别，返回所获取的服务配置信息。
# <翻译结束>

