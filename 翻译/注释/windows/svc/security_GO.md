
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
// IsAnInteractiveSession determines if calling process is running interactively.
// It queries the process token for membership in the Interactive group.
// http://stackoverflow.com/questions/2668851/how-do-i-detect-that-my-application-is-running-as-service-or-in-an-interactive-s
//
// Deprecated: Use IsWindowsService instead.
<原文结束>

# <翻译开始>
// IsAnInteractiveSession 用于判断调用进程是否以交互模式运行。
// 它通过查询进程令牌中是否存在“Interactive”组来实现这一判断。
// 参考：http://stackoverflow.com/questions/2668851/how-do-i-detect-that-my-application-is-running-as-service-or-in-an-interactive-s
// 
// 注：已弃用，请改用 IsWindowsService。
# <翻译结束>


<原文开始>
// IsWindowsService reports whether the process is currently executing
// as a Windows service.
<原文结束>

# <翻译开始>
// IsWindowsService 报告当前进程是否作为 Windows 服务执行。
# <翻译结束>


<原文开始>
	// The below technique looks a bit hairy, but it's actually
	// exactly what the .NET framework does for the similarly named function:
	// https://github.com/dotnet/extensions/blob/f4066026ca06984b07e90e61a6390ac38152ba93/src/Hosting/WindowsServices/src/WindowsServiceHelpers.cs#L26-L31
	// Specifically, it looks up whether the parent process has session ID zero
	// and is called "services".
<原文结束>

# <翻译开始>
// 下面的技巧看起来有些复杂，但实际上正是.NET框架对同名函数所采用的方法：
// https://github.com/dotnet/extensions/blob/f4066026ca06984b07e90e61a6390ac38152ba93/src/Hosting/WindowsServices/src/WindowsServiceHelpers.cs#L26-L31
// 具体来说，它会检查父进程是否具有会话ID为零，并且名为“services”。
# <翻译结束>

