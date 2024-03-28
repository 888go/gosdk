
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
// fdMutex is a specialized synchronization primitive that manages
// lifetime of an fd and serializes access to Read, Write and Close
// methods on FD.
<原文结束>

# <翻译开始>
// fdMutex is a specialized synchronization primitive that manages
// lifetime of an fd and serializes access to Read, Write and Close
// methods on FD.
# <翻译结束>


<原文开始>
// fdMutex.state is organized as follows:
// 1 bit - whether FD is closed, if set all subsequent lock operations will fail.
// 1 bit - lock for read operations.
// 1 bit - lock for write operations.
// 20 bits - total number of references (read+write+misc).
// 20 bits - number of outstanding read waiters.
// 20 bits - number of outstanding write waiters.
<原文结束>

# <翻译开始>
// fdMutex.state is organized as follows:
// 1 bit - whether FD is closed, if set all subsequent lock operations will fail.
// 1 bit - lock for read operations.
// 1 bit - lock for write operations.
// 20 bits - total number of references (read+write+misc).
// 20 bits - number of outstanding read waiters.
// 20 bits - number of outstanding write waiters.
# <翻译结束>


<原文开始>
// Read operations must do rwlock(true)/rwunlock(true).
//
// Write operations must do rwlock(false)/rwunlock(false).
//
// Misc operations must do incref/decref.
// Misc operations include functions like setsockopt and setDeadline.
// They need to use incref/decref to ensure that they operate on the
// correct fd in presence of a concurrent close call (otherwise fd can
// be closed under their feet).
//
// Close operations must do increfAndClose/decref.
<原文结束>

# <翻译开始>
// Read operations must do rwlock(true)/rwunlock(true).
//
// Write operations must do rwlock(false)/rwunlock(false).
//
// Misc operations must do incref/decref.
// Misc operations include functions like setsockopt and setDeadline.
// They need to use incref/decref to ensure that they operate on the
// correct fd in presence of a concurrent close call (otherwise fd can
// be closed under their feet).
//
// Close operations must do increfAndClose/decref.
# <翻译结束>


<原文开始>
// incref adds a reference to mu.
// It reports whether mu is available for reading or writing.
<原文结束>

# <翻译开始>
// incref adds a reference to mu.
// It reports whether mu is available for reading or writing.
# <翻译结束>


<原文开始>
// increfAndClose sets the state of mu to closed.
// It returns false if the file was already closed.
<原文结束>

# <翻译开始>
// increfAndClose sets the state of mu to closed.
// It returns false if the file was already closed.
# <翻译结束>


<原文开始>
// Mark as closed and acquire a reference.
<原文结束>

# <翻译开始>
// Mark as closed and acquire a reference.
# <翻译结束>


<原文开始>
// Remove all read and write waiters.
<原文结束>

# <翻译开始>
// Remove all read and write waiters.
# <翻译结束>


<原文开始>
			// Wake all read and write waiters,
			// they will observe closed flag after wakeup.
<原文结束>

# <翻译开始>
			// Wake all read and write waiters,
			// they will observe closed flag after wakeup.
# <翻译结束>


<原文开始>
// decref removes a reference from mu.
// It reports whether there is no remaining reference.
<原文结束>

# <翻译开始>
// decref removes a reference from mu.
// It reports whether there is no remaining reference.
# <翻译结束>


<原文开始>
// lock adds a reference to mu and locks mu.
// It reports whether mu is available for reading or writing.
<原文结束>

# <翻译开始>
// lock adds a reference to mu and locks mu.
// It reports whether mu is available for reading or writing.
# <翻译结束>


<原文开始>
// Lock is free, acquire it.
<原文结束>

# <翻译开始>
// Lock is free, acquire it.
# <翻译结束>


<原文开始>
// The signaller has subtracted mutexWait.
<原文结束>

# <翻译开始>
// The signaller has subtracted mutexWait.
# <翻译结束>


<原文开始>
// unlock removes a reference from mu and unlocks mu.
// It reports whether there is no remaining reference.
<原文结束>

# <翻译开始>
// unlock removes a reference from mu and unlocks mu.
// It reports whether there is no remaining reference.
# <翻译结束>


<原文开始>
// Drop lock, drop reference and wake read waiter if present.
<原文结束>

# <翻译开始>
// Drop lock, drop reference and wake read waiter if present.
# <翻译结束>


<原文开始>
// Implemented in runtime package.
<原文结束>

# <翻译开始>
// Implemented in runtime package.
# <翻译结束>


<原文开始>
// incref adds a reference to fd.
// It returns an error when fd cannot be used.
<原文结束>

# <翻译开始>
// incref adds a reference to fd.
// It returns an error when fd cannot be used.
# <翻译结束>


<原文开始>
// decref removes a reference from fd.
// It also closes fd when the state of fd is set to closed and there
// is no remaining reference.
<原文结束>

# <翻译开始>
// decref removes a reference from fd.
// It also closes fd when the state of fd is set to closed and there
// is no remaining reference.
# <翻译结束>


<原文开始>
// readLock adds a reference to fd and locks fd for reading.
// It returns an error when fd cannot be used for reading.
<原文结束>

# <翻译开始>
// readLock adds a reference to fd and locks fd for reading.
// It returns an error when fd cannot be used for reading.
# <翻译结束>


<原文开始>
// readUnlock removes a reference from fd and unlocks fd for reading.
// It also closes fd when the state of fd is set to closed and there
// is no remaining reference.
<原文结束>

# <翻译开始>
// readUnlock removes a reference from fd and unlocks fd for reading.
// It also closes fd when the state of fd is set to closed and there
// is no remaining reference.
# <翻译结束>


<原文开始>
// writeLock adds a reference to fd and locks fd for writing.
// It returns an error when fd cannot be used for writing.
<原文结束>

# <翻译开始>
// writeLock adds a reference to fd and locks fd for writing.
// It returns an error when fd cannot be used for writing.
# <翻译结束>


<原文开始>
// writeUnlock removes a reference from fd and unlocks fd for writing.
// It also closes fd when the state of fd is set to closed and there
// is no remaining reference.
<原文结束>

# <翻译开始>
// writeUnlock removes a reference from fd and unlocks fd for writing.
// It also closes fd when the state of fd is set to closed and there
// is no remaining reference.
# <翻译结束>

