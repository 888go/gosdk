
<原文开始>
// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
# <翻译结束>


<原文开始>
// This file implements accept for platforms that do not provide a fast path for
// setting SetNonblock and CloseOnExec.
<原文结束>

# <翻译开始>
// This file implements accept for platforms that do not provide a fast path for
// setting SetNonblock and CloseOnExec.
# <翻译结束>


<原文开始>
// Wrapper around the accept system call that marks the returned file
// descriptor as nonblocking and close-on-exec.
<原文结束>

# <翻译开始>
// Wrapper around the accept system call that marks the returned file
// descriptor as nonblocking and close-on-exec.
# <翻译结束>


<原文开始>
	// See ../syscall/exec_unix.go for description of ForkLock.
	// It is probably okay to hold the lock across syscall.Accept
	// because we have put fd.sysfd into non-blocking mode.
	// However, a call to the File method will put it back into
	// blocking mode. We can't take that risk, so no use of ForkLock here.
<原文结束>

# <翻译开始>
	// See ../syscall/exec_unix.go for description of ForkLock.
	// It is probably okay to hold the lock across syscall.Accept
	// because we have put fd.sysfd into non-blocking mode.
	// However, a call to the File method will put it back into
	// blocking mode. We can't take that risk, so no use of ForkLock here.
# <翻译结束>

