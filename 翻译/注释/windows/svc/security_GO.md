
<原文开始>
// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
# <翻译结束>


<原文开始>
// IsAnInteractiveSession determines if calling process is running interactively.
// It queries the process token for membership in the Interactive group.
// http://stackoverflow.com/questions/2668851/how-do-i-detect-that-my-application-is-running-as-service-or-in-an-interactive-s
//
// Deprecated: Use IsWindowsService instead.
<原文结束>

# <翻译开始>
// IsAnInteractiveSession determines if calling process is running interactively.
// It queries the process token for membership in the Interactive group.
// http://stackoverflow.com/questions/2668851/how-do-i-detect-that-my-application-is-running-as-service-or-in-an-interactive-s
//
// Deprecated: Use IsWindowsService instead.
# <翻译结束>


<原文开始>
// IsWindowsService reports whether the process is currently executing
// as a Windows service.
<原文结束>

# <翻译开始>
// IsWindowsService reports whether the process is currently executing
// as a Windows service.
# <翻译结束>


<原文开始>
	// The below technique looks a bit hairy, but it's actually
	// exactly what the .NET framework does for the similarly named function:
	// https://github.com/dotnet/extensions/blob/f4066026ca06984b07e90e61a6390ac38152ba93/src/Hosting/WindowsServices/src/WindowsServiceHelpers.cs#L26-L31
	// Specifically, it looks up whether the parent process has session ID zero
	// and is called "services".
<原文结束>

# <翻译开始>
	// The below technique looks a bit hairy, but it's actually
	// exactly what the .NET framework does for the similarly named function:
	// https://github.com/dotnet/extensions/blob/f4066026ca06984b07e90e61a6390ac38152ba93/src/Hosting/WindowsServices/src/WindowsServiceHelpers.cs#L26-L31
	// Specifically, it looks up whether the parent process has session ID zero
	// and is called "services".
# <翻译结束>

