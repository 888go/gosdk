
<原文开始>
// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
# <翻译结束>


<原文开始>
// Package testlog provides a back-channel communication path
// between tests and package os, so that cmd/go can see which
// environment variables and files a test consults.
<原文结束>

# <翻译开始>
// Package testlog provides a back-channel communication path
// between tests and package os, so that cmd/go can see which
// environment variables and files a test consults.
# <翻译结束>


<原文开始>
// Interface is the interface required of test loggers.
// The os package will invoke the interface's methods to indicate that
// it is inspecting the given environment variables or files.
// Multiple goroutines may call these methods simultaneously.
<原文结束>

# <翻译开始>
// Interface is the interface required of test loggers.
// The os package will invoke the interface's methods to indicate that
// it is inspecting the given environment variables or files.
// Multiple goroutines may call these methods simultaneously.
# <翻译结束>


<原文开始>
// logger is the current logger Interface.
// We use an atomic.Value in case test startup
// is racing with goroutines started during init.
// That must not cause a race detector failure,
// although it will still result in limited visibility
// into exactly what those goroutines do.
<原文结束>

# <翻译开始>
// logger is the current logger Interface.
// We use an atomic.Value in case test startup
// is racing with goroutines started during init.
// That must not cause a race detector failure,
// although it will still result in limited visibility
// into exactly what those goroutines do.
# <翻译结束>


<原文开始>
// SetLogger sets the test logger implementation for the current process.
// It must be called only once, at process startup.
<原文结束>

# <翻译开始>
// SetLogger sets the test logger implementation for the current process.
// It must be called only once, at process startup.
# <翻译结束>


<原文开始>
// Logger returns the current test logger implementation.
// It returns nil if there is no logger.
<原文结束>

# <翻译开始>
// Logger returns the current test logger implementation.
// It returns nil if there is no logger.
# <翻译结束>


<原文开始>
// Getenv calls Logger().Getenv, if a logger has been set.
<原文结束>

# <翻译开始>
// Getenv calls Logger().Getenv, if a logger has been set.
# <翻译结束>


<原文开始>
// Open calls Logger().Open, if a logger has been set.
<原文结束>

# <翻译开始>
// Open calls Logger().Open, if a logger has been set.
# <翻译结束>


<原文开始>
// Stat calls Logger().Stat, if a logger has been set.
<原文结束>

# <翻译开始>
// Stat calls Logger().Stat, if a logger has been set.
# <翻译结束>

