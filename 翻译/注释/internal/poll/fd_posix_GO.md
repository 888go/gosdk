
<原文开始>
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
# <翻译结束>


<原文开始>
// eofError returns io.EOF when fd is available for reading end of
// file.
<原文结束>

# <翻译开始>
// eofError returns io.EOF when fd is available for reading end of
// file.
# <翻译结束>


<原文开始>
// Shutdown wraps syscall.Shutdown.
<原文结束>

# <翻译开始>
// Shutdown wraps syscall.Shutdown.
# <翻译结束>


<原文开始>
// Fchown wraps syscall.Fchown.
<原文结束>

# <翻译开始>
// Fchown wraps syscall.Fchown.
# <翻译结束>


<原文开始>
// Ftruncate wraps syscall.Ftruncate.
<原文结束>

# <翻译开始>
// Ftruncate wraps syscall.Ftruncate.
# <翻译结束>


<原文开始>
// RawControl invokes the user-defined function f for a non-IO
// operation.
<原文结束>

# <翻译开始>
// RawControl invokes the user-defined function f for a non-IO
// operation.
# <翻译结束>


<原文开始>
// ignoringEINTR makes a function call and repeats it if it returns
// an EINTR error. This appears to be required even though we install all
// signal handlers with SA_RESTART: see #22838, #38033, #38836, #40846.
// Also #20400 and #36644 are issues in which a signal handler is
// installed without setting SA_RESTART. None of these are the common case,
// but there are enough of them that it seems that we can't avoid
// an EINTR loop.
<原文结束>

# <翻译开始>
// ignoringEINTR makes a function call and repeats it if it returns
// an EINTR error. This appears to be required even though we install all
// signal handlers with SA_RESTART: see #22838, #38033, #38836, #40846.
// Also #20400 and #36644 are issues in which a signal handler is
// installed without setting SA_RESTART. None of these are the common case,
// but there are enough of them that it seems that we can't avoid
// an EINTR loop.
# <翻译结束>

