
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
// FD is a file descriptor. The net and os packages use this type as a
// field of a larger type representing a network connection or OS file.
<原文结束>

# <翻译开始>
// FD is a file descriptor. The net and os packages use this type as a
// field of a larger type representing a network connection or OS file.
# <翻译结束>


<原文开始>
// Lock sysfd and serialize access to Read and Write methods.
<原文结束>

# <翻译开始>
// Lock sysfd and serialize access to Read and Write methods.
# <翻译结束>


<原文开始>
// System file descriptor. Immutable until Close.
<原文结束>

# <翻译开始>
// System file descriptor. Immutable until Close.
# <翻译结束>


<原文开始>
// Semaphore signaled when file is closed.
<原文结束>

# <翻译开始>
// Semaphore signaled when file is closed.
# <翻译结束>


<原文开始>
// Non-zero if this file has been set to blocking mode.
<原文结束>

# <翻译开始>
// Non-zero if this file has been set to blocking mode.
# <翻译结束>


<原文开始>
	// Whether this is a streaming descriptor, as opposed to a
	// packet-based descriptor like a UDP socket. Immutable.
<原文结束>

# <翻译开始>
	// Whether this is a streaming descriptor, as opposed to a
	// packet-based descriptor like a UDP socket. Immutable.
# <翻译结束>


<原文开始>
	// Whether a zero byte read indicates EOF. This is false for a
	// message based socket connection.
<原文结束>

# <翻译开始>
	// Whether a zero byte read indicates EOF. This is false for a
	// message based socket connection.
# <翻译结束>


<原文开始>
// Whether this is a file rather than a network socket.
<原文结束>

# <翻译开始>
// Whether this is a file rather than a network socket.
# <翻译结束>


<原文开始>
// Init initializes the FD. The Sysfd field should already be set.
// This can be called multiple times on a single FD.
// The net argument is a network name from the net package (e.g., "tcp"),
// or "file".
// Set pollable to true if fd should be managed by runtime netpoll.
<原文结束>

# <翻译开始>
// Init initializes the FD. The Sysfd field should already be set.
// This can be called multiple times on a single FD.
// The net argument is a network name from the net package (e.g., "tcp"),
// or "file".
// Set pollable to true if fd should be managed by runtime netpoll.
# <翻译结束>


<原文开始>
// We don't actually care about the various network types.
<原文结束>

# <翻译开始>
// We don't actually care about the various network types.
# <翻译结束>


<原文开始>
		// If we could not initialize the runtime poller,
		// assume we are using blocking mode.
<原文结束>

# <翻译开始>
		// If we could not initialize the runtime poller,
		// assume we are using blocking mode.
# <翻译结束>


<原文开始>
// Destroy closes the file descriptor. This is called when there are
// no remaining references.
<原文结束>

# <翻译开始>
// Destroy closes the file descriptor. This is called when there are
// no remaining references.
# <翻译结束>


<原文开始>
	// Poller may want to unregister fd in readiness notification mechanism,
	// so this must be executed before CloseFunc.
<原文结束>

# <翻译开始>
	// Poller may want to unregister fd in readiness notification mechanism,
	// so this must be executed before CloseFunc.
# <翻译结束>


<原文开始>
	// We don't use ignoringEINTR here because POSIX does not define
	// whether the descriptor is closed if close returns EINTR.
	// If the descriptor is indeed closed, using a loop would race
	// with some other goroutine opening a new descriptor.
	// (The Linux kernel guarantees that it is closed on an EINTR error.)
<原文结束>

# <翻译开始>
	// We don't use ignoringEINTR here because POSIX does not define
	// whether the descriptor is closed if close returns EINTR.
	// If the descriptor is indeed closed, using a loop would race
	// with some other goroutine opening a new descriptor.
	// (The Linux kernel guarantees that it is closed on an EINTR error.)
# <翻译结束>


<原文开始>
// Close closes the FD. The underlying file descriptor is closed by the
// destroy method when there are no remaining references.
<原文结束>

# <翻译开始>
// Close closes the FD. The underlying file descriptor is closed by the
// destroy method when there are no remaining references.
# <翻译结束>


<原文开始>
	// Unblock any I/O.  Once it all unblocks and returns,
	// so that it cannot be referring to fd.sysfd anymore,
	// the final decref will close fd.sysfd. This should happen
	// fairly quickly, since all the I/O is non-blocking, and any
	// attempts to block in the pollDesc will return errClosing(fd.isFile).
<原文结束>

# <翻译开始>
	// Unblock any I/O.  Once it all unblocks and returns,
	// so that it cannot be referring to fd.sysfd anymore,
	// the final decref will close fd.sysfd. This should happen
	// fairly quickly, since all the I/O is non-blocking, and any
	// attempts to block in the pollDesc will return errClosing(fd.isFile).
# <翻译结束>


<原文开始>
	// The call to decref will call destroy if there are no other
	// references.
<原文结束>

# <翻译开始>
	// The call to decref will call destroy if there are no other
	// references.
# <翻译结束>


<原文开始>
	// Wait until the descriptor is closed. If this was the only
	// reference, it is already closed. Only wait if the file has
	// not been set to blocking mode, as otherwise any current I/O
	// may be blocking, and that would block the Close.
	// No need for an atomic read of isBlocking, increfAndClose means
	// we have exclusive access to fd.
<原文结束>

# <翻译开始>
	// Wait until the descriptor is closed. If this was the only
	// reference, it is already closed. Only wait if the file has
	// not been set to blocking mode, as otherwise any current I/O
	// may be blocking, and that would block the Close.
	// No need for an atomic read of isBlocking, increfAndClose means
	// we have exclusive access to fd.
# <翻译结束>


<原文开始>
// SetBlocking puts the file into blocking mode.
<原文结束>

# <翻译开始>
// SetBlocking puts the file into blocking mode.
# <翻译结束>


<原文开始>
	// Atomic store so that concurrent calls to SetBlocking
	// do not cause a race condition. isBlocking only ever goes
	// from 0 to 1 so there is no real race here.
<原文结束>

# <翻译开始>
	// Atomic store so that concurrent calls to SetBlocking
	// do not cause a race condition. isBlocking only ever goes
	// from 0 to 1 so there is no real race here.
# <翻译结束>


<原文开始>
// Darwin and FreeBSD can't read or write 2GB+ files at a time,
// even on 64-bit systems.
// The same is true of socket implementations on many systems.
// See golang.org/issue/7812 and golang.org/issue/16266.
// Use 1GB instead of, say, 2GB-1, to keep subsequent reads aligned.
<原文结束>

# <翻译开始>
// Darwin and FreeBSD can't read or write 2GB+ files at a time,
// even on 64-bit systems.
// The same is true of socket implementations on many systems.
// See golang.org/issue/7812 and golang.org/issue/16266.
// Use 1GB instead of, say, 2GB-1, to keep subsequent reads aligned.
# <翻译结束>


<原文开始>
// Read implements io.Reader.
<原文结束>

# <翻译开始>
// Read implements io.Reader.
# <翻译结束>


<原文开始>
		// If the caller wanted a zero byte read, return immediately
		// without trying (but after acquiring the readLock).
		// Otherwise syscall.Read returns 0, nil which looks like
		// io.EOF.
		// TODO(bradfitz): make it wait for readability? (Issue 15735)
<原文结束>

# <翻译开始>
		// If the caller wanted a zero byte read, return immediately
		// without trying (but after acquiring the readLock).
		// Otherwise syscall.Read returns 0, nil which looks like
		// io.EOF.
		// TODO(bradfitz): make it wait for readability? (Issue 15735)
# <翻译结束>


<原文开始>
// Pread wraps the pread system call.
<原文结束>

# <翻译开始>
// Pread wraps the pread system call.
# <翻译结束>


<原文开始>
	// Call incref, not readLock, because since pread specifies the
	// offset it is independent from other reads.
	// Similarly, using the poller doesn't make sense for pread.
<原文结束>

# <翻译开始>
	// Call incref, not readLock, because since pread specifies the
	// offset it is independent from other reads.
	// Similarly, using the poller doesn't make sense for pread.
# <翻译结束>


<原文开始>
// ReadFrom wraps the recvfrom network call.
<原文结束>

# <翻译开始>
// ReadFrom wraps the recvfrom network call.
# <翻译结束>


<原文开始>
// ReadFromInet4 wraps the recvfrom network call for IPv4.
<原文结束>

# <翻译开始>
// ReadFromInet4 wraps the recvfrom network call for IPv4.
# <翻译结束>


<原文开始>
// ReadFromInet6 wraps the recvfrom network call for IPv6.
<原文结束>

# <翻译开始>
// ReadFromInet6 wraps the recvfrom network call for IPv6.
# <翻译结束>


<原文开始>
// ReadMsg wraps the recvmsg network call.
<原文结束>

# <翻译开始>
// ReadMsg wraps the recvmsg network call.
# <翻译结束>


<原文开始>
// TODO(dfc) should n and oobn be set to 0
<原文结束>

# <翻译开始>
// TODO(dfc) should n and oobn be set to 0
# <翻译结束>


<原文开始>
// ReadMsgInet4 is ReadMsg, but specialized for syscall.SockaddrInet4.
<原文结束>

# <翻译开始>
// ReadMsgInet4 is ReadMsg, but specialized for syscall.SockaddrInet4.
# <翻译结束>


<原文开始>
// ReadMsgInet6 is ReadMsg, but specialized for syscall.SockaddrInet6.
<原文结束>

# <翻译开始>
// ReadMsgInet6 is ReadMsg, but specialized for syscall.SockaddrInet6.
# <翻译结束>


<原文开始>
// Write implements io.Writer.
<原文结束>

# <翻译开始>
// Write implements io.Writer.
# <翻译结束>


<原文开始>
// Pwrite wraps the pwrite system call.
<原文结束>

# <翻译开始>
// Pwrite wraps the pwrite system call.
# <翻译结束>


<原文开始>
	// Call incref, not writeLock, because since pwrite specifies the
	// offset it is independent from other writes.
	// Similarly, using the poller doesn't make sense for pwrite.
<原文结束>

# <翻译开始>
	// Call incref, not writeLock, because since pwrite specifies the
	// offset it is independent from other writes.
	// Similarly, using the poller doesn't make sense for pwrite.
# <翻译结束>


<原文开始>
// WriteToInet4 wraps the sendto network call for IPv4 addresses.
<原文结束>

# <翻译开始>
// WriteToInet4 wraps the sendto network call for IPv4 addresses.
# <翻译结束>


<原文开始>
// WriteToInet6 wraps the sendto network call for IPv6 addresses.
<原文结束>

# <翻译开始>
// WriteToInet6 wraps the sendto network call for IPv6 addresses.
# <翻译结束>


<原文开始>
// WriteTo wraps the sendto network call.
<原文结束>

# <翻译开始>
// WriteTo wraps the sendto network call.
# <翻译结束>


<原文开始>
// WriteMsg wraps the sendmsg network call.
<原文结束>

# <翻译开始>
// WriteMsg wraps the sendmsg network call.
# <翻译结束>


<原文开始>
// WriteMsgInet4 is WriteMsg specialized for syscall.SockaddrInet4.
<原文结束>

# <翻译开始>
// WriteMsgInet4 is WriteMsg specialized for syscall.SockaddrInet4.
# <翻译结束>


<原文开始>
// WriteMsgInet6 is WriteMsg specialized for syscall.SockaddrInet6.
<原文结束>

# <翻译开始>
// WriteMsgInet6 is WriteMsg specialized for syscall.SockaddrInet6.
# <翻译结束>


<原文开始>
// Accept wraps the accept network call.
<原文结束>

# <翻译开始>
// Accept wraps the accept network call.
# <翻译结束>


<原文开始>
			// This means that a socket on the listen
			// queue was closed before we Accept()ed it;
			// it's a silly error, so try again.
<原文结束>

# <翻译开始>
			// This means that a socket on the listen
			// queue was closed before we Accept()ed it;
			// it's a silly error, so try again.
# <翻译结束>


<原文开始>
// Seek wraps syscall.Seek.
<原文结束>

# <翻译开始>
// Seek wraps syscall.Seek.
# <翻译结束>


<原文开始>
// ReadDirent wraps syscall.ReadDirent.
// We treat this like an ordinary system call rather than a call
// that tries to fill the buffer.
<原文结束>

# <翻译开始>
// ReadDirent wraps syscall.ReadDirent.
// We treat this like an ordinary system call rather than a call
// that tries to fill the buffer.
# <翻译结束>


<原文开始>
// Do not call eofError; caller does not expect to see io.EOF.
<原文结束>

# <翻译开始>
// Do not call eofError; caller does not expect to see io.EOF.
# <翻译结束>


<原文开始>
// Fchmod wraps syscall.Fchmod.
<原文结束>

# <翻译开始>
// Fchmod wraps syscall.Fchmod.
# <翻译结束>


<原文开始>
// Fchdir wraps syscall.Fchdir.
<原文结束>

# <翻译开始>
// Fchdir wraps syscall.Fchdir.
# <翻译结束>


<原文开始>
// Fstat wraps syscall.Fstat
<原文结束>

# <翻译开始>
// Fstat wraps syscall.Fstat
# <翻译结束>


<原文开始>
// tryDupCloexec indicates whether F_DUPFD_CLOEXEC should be used.
// If the kernel doesn't support it, this is set to 0.
<原文结束>

# <翻译开始>
// tryDupCloexec indicates whether F_DUPFD_CLOEXEC should be used.
// If the kernel doesn't support it, this is set to 0.
# <翻译结束>


<原文开始>
// DupCloseOnExec dups fd and marks it close-on-exec.
<原文结束>

# <翻译开始>
// DupCloseOnExec dups fd and marks it close-on-exec.
# <翻译结束>


<原文开始>
			// Old kernel, or js/wasm (which returns
			// ENOSYS). Fall back to the portable way from
			// now on.
<原文结束>

# <翻译开始>
			// Old kernel, or js/wasm (which returns
			// ENOSYS). Fall back to the portable way from
			// now on.
# <翻译结束>


<原文开始>
// dupCloseOnExecOld is the traditional way to dup an fd and
// set its O_CLOEXEC bit, using two system calls.
<原文结束>

# <翻译开始>
// dupCloseOnExecOld is the traditional way to dup an fd and
// set its O_CLOEXEC bit, using two system calls.
# <翻译结束>


<原文开始>
// Dup duplicates the file descriptor.
<原文结束>

# <翻译开始>
// Dup duplicates the file descriptor.
# <翻译结束>


<原文开始>
// On Unix variants only, expose the IO event for the net code.
<原文结束>

# <翻译开始>
// On Unix variants only, expose the IO event for the net code.
# <翻译结束>


<原文开始>
// WaitWrite waits until data can be read from fd.
<原文结束>

# <翻译开始>
// WaitWrite waits until data can be read from fd.
# <翻译结束>


<原文开始>
// WriteOnce is for testing only. It makes a single write call.
<原文结束>

# <翻译开始>
// WriteOnce is for testing only. It makes a single write call.
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


<原文开始>
// ignoringEINTRIO is like ignoringEINTR, but just for IO calls.
<原文结束>

# <翻译开始>
// ignoringEINTRIO is like ignoringEINTR, but just for IO calls.
# <翻译结束>

