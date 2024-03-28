
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
// Lock sysfd and serialize access to Read and Write methods.
<原文结束>

# <翻译开始>
// Lock sysfd and serialize access to Read and Write methods.
# <翻译结束>


<原文开始>
// set true when read deadline has been reached
<原文结束>

# <翻译开始>
// set true when read deadline has been reached
# <翻译结束>


<原文开始>
// set true when write deadline has been reached
<原文结束>

# <翻译开始>
// set true when write deadline has been reached
# <翻译结束>


<原文开始>
	// Whether this is a normal file.
	// On Plan 9 we do not use this package for ordinary files,
	// so this is always false, but the field is present because
	// shared code in fd_mutex.go checks it.
<原文结束>

# <翻译开始>
	// Whether this is a normal file.
	// On Plan 9 we do not use this package for ordinary files,
	// so this is always false, but the field is present because
	// shared code in fd_mutex.go checks it.
# <翻译结束>


<原文开始>
// We need this to close out a file descriptor when it is unlocked,
// but the real implementation has to live in the net package because
// it uses os.File's.
<原文结束>

# <翻译开始>
// We need this to close out a file descriptor when it is unlocked,
// but the real implementation has to live in the net package because
// it uses os.File's.
# <翻译结束>


<原文开始>
// Close handles the locking for closing an FD. The real operation
// is in the net package.
<原文结束>

# <翻译开始>
// Close handles the locking for closing an FD. The real operation
// is in the net package.
# <翻译结束>


<原文开始>
// Read implements io.Reader.
<原文结束>

# <翻译开始>
// Read implements io.Reader.
# <翻译结束>


<原文开始>
// Write implements io.Writer.
<原文结束>

# <翻译开始>
// Write implements io.Writer.
# <翻译结束>


<原文开始>
// SetDeadline sets the read and write deadlines associated with fd.
<原文结束>

# <翻译开始>
// SetDeadline sets the read and write deadlines associated with fd.
# <翻译结束>


<原文开始>
// SetReadDeadline sets the read deadline associated with fd.
<原文结束>

# <翻译开始>
// SetReadDeadline sets the read deadline associated with fd.
# <翻译结束>


<原文开始>
// SetWriteDeadline sets the write deadline associated with fd.
<原文结束>

# <翻译开始>
// SetWriteDeadline sets the write deadline associated with fd.
# <翻译结束>


<原文开始>
// Interrupt I/O operation once timer has expired
<原文结束>

# <翻译开始>
// Interrupt I/O operation once timer has expired
# <翻译结束>


<原文开始>
// Interrupt current I/O operation
<原文结束>

# <翻译开始>
// Interrupt current I/O operation
# <翻译结束>


<原文开始>
// On Plan 9 only, expose the locking for the net code.
<原文结束>

# <翻译开始>
// On Plan 9 only, expose the locking for the net code.
# <翻译结束>


<原文开始>
// ReadLock wraps FD.readLock.
<原文结束>

# <翻译开始>
// ReadLock wraps FD.readLock.
# <翻译结束>


<原文开始>
// ReadUnlock wraps FD.readUnlock.
<原文结束>

# <翻译开始>
// ReadUnlock wraps FD.readUnlock.
# <翻译结束>


<原文开始>
// IsPollDescriptor reports whether fd is the descriptor being used by the poller.
// This is only used for testing.
<原文结束>

# <翻译开始>
// IsPollDescriptor reports whether fd is the descriptor being used by the poller.
// This is only used for testing.
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
// RawRead invokes the user-defined function f for a read operation.
<原文结束>

# <翻译开始>
// RawRead invokes the user-defined function f for a read operation.
# <翻译结束>


<原文开始>
// RawWrite invokes the user-defined function f for a write operation.
<原文结束>

# <翻译开始>
// RawWrite invokes the user-defined function f for a write operation.
# <翻译结束>

