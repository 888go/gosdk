
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
// This package uses the SetFileCompletionNotificationModes Windows
// API to skip calling GetQueuedCompletionStatus if an IO operation
// completes synchronously. There is a known bug where
// SetFileCompletionNotificationModes crashes on some systems (see
// https://support.microsoft.com/kb/2568167 for details).
<原文结束>

# <翻译开始>
// This package uses the SetFileCompletionNotificationModes Windows
// API to skip calling GetQueuedCompletionStatus if an IO operation
// completes synchronously. There is a known bug where
// SetFileCompletionNotificationModes crashes on some systems (see
// https://support.microsoft.com/kb/2568167 for details).
# <翻译结束>


<原文开始>
// determines is SetFileCompletionNotificationModes is present and safe to use
<原文结束>

# <翻译开始>
// determines is SetFileCompletionNotificationModes is present and safe to use
# <翻译结束>


<原文开始>
// checkSetFileCompletionNotificationModes verifies that
// SetFileCompletionNotificationModes Windows API is present
// on the system and is safe to use.
// See https://support.microsoft.com/kb/2568167 for details.
<原文结束>

# <翻译开始>
// checkSetFileCompletionNotificationModes verifies that
// SetFileCompletionNotificationModes Windows API is present
// on the system and is safe to use.
// See https://support.microsoft.com/kb/2568167 for details.
# <翻译结束>


<原文开始>
// operation contains superset of data necessary to perform all async IO.
<原文结束>

# <翻译开始>
// operation contains superset of data necessary to perform all async IO.
# <翻译结束>


<原文开始>
	// Used by IOCP interface, it must be first field
	// of the struct, as our code rely on it.
<原文结束>

# <翻译开始>
	// Used by IOCP interface, it must be first field
	// of the struct, as our code rely on it.
# <翻译结束>


<原文开始>
// fields used by runtime.netpoll
<原文结束>

# <翻译开始>
// fields used by runtime.netpoll
# <翻译结束>


<原文开始>
// fields used only by net package
<原文结束>

# <翻译开始>
// fields used only by net package
# <翻译结束>


<原文开始>
// ClearBufs clears all pointers to Buffers parameter captured
// by InitBufs, so it can be released by garbage collector.
<原文结束>

# <翻译开始>
// ClearBufs clears all pointers to Buffers parameter captured
// by InitBufs, so it can be released by garbage collector.
# <翻译结束>


<原文开始>
// execIO executes a single IO operation o. It submits and cancels
// IO in the current thread for systems where Windows CancelIoEx API
// is available. Alternatively, it passes the request onto
// runtime netpoll and waits for completion or cancels request.
<原文结束>

# <翻译开始>
// execIO executes a single IO operation o. It submits and cancels
// IO in the current thread for systems where Windows CancelIoEx API
// is available. Alternatively, it passes the request onto
// runtime netpoll and waits for completion or cancels request.
# <翻译结束>


<原文开始>
// Notify runtime netpoll about starting IO.
<原文结束>

# <翻译开始>
// Notify runtime netpoll about starting IO.
# <翻译结束>


<原文开始>
// IO completed immediately
<原文结束>

# <翻译开始>
// IO completed immediately
# <翻译结束>


<原文开始>
// No completion message will follow, so return immediately.
<原文结束>

# <翻译开始>
// No completion message will follow, so return immediately.
# <翻译结束>


<原文开始>
// Need to get our completion message anyway.
<原文结束>

# <翻译开始>
// Need to get our completion message anyway.
# <翻译结束>


<原文开始>
// IO started, and we have to wait for its completion.
<原文结束>

# <翻译开始>
// IO started, and we have to wait for its completion.
# <翻译结束>


<原文开始>
// Wait for our request to complete.
<原文结束>

# <翻译开始>
// Wait for our request to complete.
# <翻译结束>


<原文开始>
// All is good. Extract our IO results and return.
<原文结束>

# <翻译开始>
// All is good. Extract our IO results and return.
# <翻译结束>


<原文开始>
// More data available. Return back the size of received data.
<原文结束>

# <翻译开始>
// More data available. Return back the size of received data.
# <翻译结束>


<原文开始>
// IO is interrupted by "close" or "timeout"
<原文结束>

# <翻译开始>
// IO is interrupted by "close" or "timeout"
# <翻译结束>


<原文开始>
// Assuming ERROR_NOT_FOUND is returned, if IO is completed.
<原文结束>

# <翻译开始>
// Assuming ERROR_NOT_FOUND is returned, if IO is completed.
# <翻译结束>


<原文开始>
// TODO(brainman): maybe do something else, but panic.
<原文结束>

# <翻译开始>
// TODO(brainman): maybe do something else, but panic.
# <翻译结束>


<原文开始>
// Wait for cancellation to complete.
<原文结束>

# <翻译开始>
// Wait for cancellation to complete.
# <翻译结束>


<原文开始>
	// We issued a cancellation request. But, it seems, IO operation succeeded
	// before the cancellation request run. We need to treat the IO operation as
	// succeeded (the bytes are actually sent/recv from network).
<原文结束>

# <翻译开始>
	// We issued a cancellation request. But, it seems, IO operation succeeded
	// before the cancellation request run. We need to treat the IO operation as
	// succeeded (the bytes are actually sent/recv from network).
# <翻译结束>


<原文开始>
// FD is a file descriptor. The net and os packages embed this type in
// a larger type representing a network connection or OS file.
<原文结束>

# <翻译开始>
// FD is a file descriptor. The net and os packages embed this type in
// a larger type representing a network connection or OS file.
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
// Used to implement pread/pwrite.
<原文结束>

# <翻译开始>
// Used to implement pread/pwrite.
# <翻译结束>


<原文开始>
// first few bytes of the last incomplete rune in last write
<原文结束>

# <翻译开始>
// first few bytes of the last incomplete rune in last write
# <翻译结束>


<原文开始>
// buffer to hold uint16s obtained with ReadConsole
<原文结束>

# <翻译开始>
// buffer to hold uint16s obtained with ReadConsole
# <翻译结束>


<原文开始>
// buffer to hold decoding of readuint16 from utf16 to utf8
<原文结束>

# <翻译开始>
// buffer to hold decoding of readuint16 from utf16 to utf8
# <翻译结束>


<原文开始>
// readbyte[readOffset:] is yet to be consumed with file.Read
<原文结束>

# <翻译开始>
// readbyte[readOffset:] is yet to be consumed with file.Read
# <翻译结束>


<原文开始>
// Semaphore signaled when file is closed.
<原文结束>

# <翻译开始>
// Semaphore signaled when file is closed.
# <翻译结束>


<原文开始>
	// Whether this is a streaming descriptor, as opposed to a
	// packet-based descriptor like a UDP socket.
<原文结束>

# <翻译开始>
	// Whether this is a streaming descriptor, as opposed to a
	// packet-based descriptor like a UDP socket.
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
// fileKind describes the kind of file.
<原文结束>

# <翻译开始>
// fileKind describes the kind of file.
# <翻译结束>


<原文开始>
// logInitFD is set by tests to enable file descriptor initialization logging.
<原文结束>

# <翻译开始>
// logInitFD is set by tests to enable file descriptor initialization logging.
# <翻译结束>


<原文开始>
// Init initializes the FD. The Sysfd field should already be set.
// This can be called multiple times on a single FD.
// The net argument is a network name from the net package (e.g., "tcp"),
// or "file" or "console" or "dir".
// Set pollable to true if fd should be managed by runtime netpoll.
<原文结束>

# <翻译开始>
// Init initializes the FD. The Sysfd field should already be set.
// This can be called multiple times on a single FD.
// The net argument is a network name from the net package (e.g., "tcp"),
// or "file" or "console" or "dir".
// Set pollable to true if fd should be managed by runtime netpoll.
# <翻译结束>


<原文开始>
		// Only call init for a network socket.
		// This means that we don't add files to the runtime poller.
		// Adding files to the runtime poller can confuse matters
		// if the user is doing their own overlapped I/O.
		// See issue #21172.
		//
		// In general the code below avoids calling the execIO
		// function for non-network sockets. If some method does
		// somehow call execIO, then execIO, and therefore the
		// calling method, will return an error, because
		// fd.pd.runtimeCtx will be 0.
<原文结束>

# <翻译开始>
		// Only call init for a network socket.
		// This means that we don't add files to the runtime poller.
		// Adding files to the runtime poller can confuse matters
		// if the user is doing their own overlapped I/O.
		// See issue #21172.
		//
		// In general the code below avoids calling the execIO
		// function for non-network sockets. If some method does
		// somehow call execIO, then execIO, and therefore the
		// calling method, will return an error, because
		// fd.pd.runtimeCtx will be 0.
# <翻译结束>


<原文开始>
// We do not use events, so we can skip them always.
<原文结束>

# <翻译开始>
// We do not use events, so we can skip them always.
# <翻译结束>


<原文开始>
		// It's not safe to skip completion notifications for UDP:
		// https://docs.microsoft.com/en-us/archive/blogs/winserverperformance/designing-applications-for-high-performance-part-iii
<原文结束>

# <翻译开始>
		// It's not safe to skip completion notifications for UDP:
		// https://docs.microsoft.com/en-us/archive/blogs/winserverperformance/designing-applications-for-high-performance-part-iii
# <翻译结束>


<原文开始>
	// Disable SIO_UDP_CONNRESET behavior.
	// http://support.microsoft.com/kb/263823
<原文结束>

# <翻译开始>
	// Disable SIO_UDP_CONNRESET behavior.
	// http://support.microsoft.com/kb/263823
# <翻译结束>


<原文开始>
	// Poller may want to unregister fd in readiness notification mechanism,
	// so this must be executed before fd.CloseFunc.
<原文结束>

# <翻译开始>
	// Poller may want to unregister fd in readiness notification mechanism,
	// so this must be executed before fd.CloseFunc.
# <翻译结束>


<原文开始>
// The net package uses the CloseFunc variable for testing.
<原文结束>

# <翻译开始>
// The net package uses the CloseFunc variable for testing.
# <翻译结束>


<原文开始>
// Close closes the FD. The underlying file descriptor is closed by
// the destroy method when there are no remaining references.
<原文结束>

# <翻译开始>
// Close closes the FD. The underlying file descriptor is closed by
// the destroy method when there are no remaining references.
# <翻译结束>


<原文开始>
// unblock pending reader and writer
<原文结束>

# <翻译开始>
// unblock pending reader and writer
# <翻译结束>


<原文开始>
	// Wait until the descriptor is closed. If this was the only
	// reference, it is already closed.
<原文结束>

# <翻译开始>
	// Wait until the descriptor is closed. If this was the only
	// reference, it is already closed.
# <翻译结束>


<原文开始>
// Windows ReadFile and WSARecv use DWORD (uint32) parameter to pass buffer length.
// This prevents us reading blocks larger than 4GB.
// See golang.org/issue/26923.
<原文结束>

# <翻译开始>
// Windows ReadFile and WSARecv use DWORD (uint32) parameter to pass buffer length.
// This prevents us reading blocks larger than 4GB.
// See golang.org/issue/26923.
# <翻译结束>


<原文开始>
// 1GB is large enough and keeps subsequent reads aligned
<原文结束>

# <翻译开始>
// 1GB is large enough and keeps subsequent reads aligned
# <翻译结束>


<原文开始>
// Read implements io.Reader.
<原文结束>

# <翻译开始>
// Read implements io.Reader.
# <翻译结束>


<原文开始>
				// Close uses CancelIoEx to interrupt concurrent I/O for pipes.
				// If the fd is a pipe and the Read was interrupted by CancelIoEx,
				// we assume it is interrupted by Close.
<原文结束>

# <翻译开始>
				// Close uses CancelIoEx to interrupt concurrent I/O for pipes.
				// If the fd is a pipe and the Read was interrupted by CancelIoEx,
				// we assume it is interrupted by Close.
# <翻译结束>


<原文开始>
// readConsole reads utf16 characters from console File,
// encodes them into utf8 and stores them in buffer b.
// It returns the number of utf8 bytes read and an error, if any.
<原文结束>

# <翻译开始>
// readConsole reads utf16 characters from console File,
// encodes them into utf8 and stores them in buffer b.
// It returns the number of utf8 bytes read and an error, if any.
# <翻译结束>


<原文开始>
		// Note: syscall.ReadConsole fails for very large buffers.
		// The limit is somewhere around (but not exactly) 16384.
		// Stay well below.
<原文结束>

# <翻译开始>
		// Note: syscall.ReadConsole fails for very large buffers.
		// The limit is somewhere around (but not exactly) 16384.
		// Stay well below.
# <翻译结束>


<原文开始>
// Save half surrogate pair for next time.
<原文结束>

# <翻译开始>
// Save half surrogate pair for next time.
# <翻译结束>


<原文开始>
// Pread emulates the Unix pread system call.
<原文结束>

# <翻译开始>
// Pread emulates the Unix pread system call.
# <翻译结束>


<原文开始>
	// Call incref, not readLock, because since pread specifies the
	// offset it is independent from other reads.
<原文结束>

# <翻译开始>
	// Call incref, not readLock, because since pread specifies the
	// offset it is independent from other reads.
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
// Write implements io.Writer.
<原文结束>

# <翻译开始>
// Write implements io.Writer.
# <翻译结束>


<原文开始>
					// Close uses CancelIoEx to interrupt concurrent I/O for pipes.
					// If the fd is a pipe and the Write was interrupted by CancelIoEx,
					// we assume it is interrupted by Close.
<原文结束>

# <翻译开始>
					// Close uses CancelIoEx to interrupt concurrent I/O for pipes.
					// If the fd is a pipe and the Write was interrupted by CancelIoEx,
					// we assume it is interrupted by Close.
# <翻译结束>


<原文开始>
// writeConsole writes len(b) bytes to the console File.
// It returns the number of bytes written and an error, if any.
<原文结束>

# <翻译开始>
// writeConsole writes len(b) bytes to the console File.
// It returns the number of bytes written and an error, if any.
# <翻译结束>


<原文开始>
	// syscall.WriteConsole seems to fail, if given large buffer.
	// So limit the buffer to 16000 characters. This number was
	// discovered by experimenting with syscall.WriteConsole.
<原文结束>

# <翻译开始>
	// syscall.WriteConsole seems to fail, if given large buffer.
	// So limit the buffer to 16000 characters. This number was
	// discovered by experimenting with syscall.WriteConsole.
# <翻译结束>


<原文开始>
// Pwrite emulates the Unix pwrite system call.
<原文结束>

# <翻译开始>
// Pwrite emulates the Unix pwrite system call.
# <翻译结束>


<原文开始>
	// Call incref, not writeLock, because since pwrite specifies the
	// offset it is independent from other writes.
<原文结束>

# <翻译开始>
	// Call incref, not writeLock, because since pwrite specifies the
	// offset it is independent from other writes.
# <翻译结束>


<原文开始>
// Writev emulates the Unix writev system call.
<原文结束>

# <翻译开始>
// Writev emulates the Unix writev system call.
# <翻译结束>


<原文开始>
// WriteTo wraps the sendto network call.
<原文结束>

# <翻译开始>
// WriteTo wraps the sendto network call.
# <翻译结束>


<原文开始>
// handle zero-byte payload
<原文结束>

# <翻译开始>
// handle zero-byte payload
# <翻译结束>


<原文开始>
// WriteToInet4 is WriteTo, specialized for syscall.SockaddrInet4.
<原文结束>

# <翻译开始>
// WriteToInet4 is WriteTo, specialized for syscall.SockaddrInet4.
# <翻译结束>


<原文开始>
// WriteToInet6 is WriteTo, specialized for syscall.SockaddrInet6.
<原文结束>

# <翻译开始>
// WriteToInet6 is WriteTo, specialized for syscall.SockaddrInet6.
# <翻译结束>


<原文开始>
// Call ConnectEx. This doesn't need any locking, since it is only
// called when the descriptor is first created. This is here rather
// than in the net package so that it can use fd.wop.
<原文结束>

# <翻译开始>
// Call ConnectEx. This doesn't need any locking, since it is only
// called when the descriptor is first created. This is here rather
// than in the net package so that it can use fd.wop.
# <翻译结束>


<原文开始>
// Inherit properties of the listening socket.
<原文结束>

# <翻译开始>
// Inherit properties of the listening socket.
# <翻译结束>


<原文开始>
// Accept handles accepting a socket. The sysSocket parameter is used
// to allocate the net socket.
<原文结束>

# <翻译开始>
// Accept handles accepting a socket. The sysSocket parameter is used
// to allocate the net socket.
# <翻译结束>


<原文开始>
		// Sometimes we see WSAECONNRESET and ERROR_NETNAME_DELETED is
		// returned here. These happen if connection reset is received
		// before AcceptEx could complete. These errors relate to new
		// connection, not to AcceptEx, so ignore broken connection and
		// try AcceptEx again for more connections.
<原文结束>

# <翻译开始>
		// Sometimes we see WSAECONNRESET and ERROR_NETNAME_DELETED is
		// returned here. These happen if connection reset is received
		// before AcceptEx could complete. These errors relate to new
		// connection, not to AcceptEx, so ignore broken connection and
		// try AcceptEx again for more connections.
# <翻译结束>


<原文开始>
// ignore these and try again
<原文结束>

# <翻译开始>
// ignore these and try again
# <翻译结束>


<原文开始>
// Seek wraps syscall.Seek.
<原文结束>

# <翻译开始>
// Seek wraps syscall.Seek.
# <翻译结束>


<原文开始>
// Fchmod updates syscall.ByHandleFileInformation.Fileattributes when needed.
<原文结束>

# <翻译开始>
// Fchmod updates syscall.ByHandleFileInformation.Fileattributes when needed.
# <翻译结束>


<原文开始>
// Fchdir wraps syscall.Fchdir.
<原文结束>

# <翻译开始>
// Fchdir wraps syscall.Fchdir.
# <翻译结束>


<原文开始>
// GetFileType wraps syscall.GetFileType.
<原文结束>

# <翻译开始>
// GetFileType wraps syscall.GetFileType.
# <翻译结束>


<原文开始>
// GetFileInformationByHandle wraps GetFileInformationByHandle.
<原文结束>

# <翻译开始>
// GetFileInformationByHandle wraps GetFileInformationByHandle.
# <翻译结束>


<原文开始>
// RawRead invokes the user-defined function f for a read operation.
<原文结束>

# <翻译开始>
// RawRead invokes the user-defined function f for a read operation.
# <翻译结束>


<原文开始>
		// Use a zero-byte read as a way to get notified when this
		// socket is readable. h/t https://stackoverflow.com/a/42019668/332798
<原文结束>

# <翻译开始>
		// Use a zero-byte read as a way to get notified when this
		// socket is readable. h/t https://stackoverflow.com/a/42019668/332798
# <翻译结束>


<原文开始>
// expected with a 0-byte peek, ignore.
<原文结束>

# <翻译开始>
// expected with a 0-byte peek, ignore.
# <翻译结束>


<原文开始>
// RawWrite invokes the user-defined function f for a write operation.
<原文结束>

# <翻译开始>
// RawWrite invokes the user-defined function f for a write operation.
# <翻译结束>


<原文开始>
// TODO(tmm1): find a way to detect socket writability
<原文结束>

# <翻译开始>
// TODO(tmm1): find a way to detect socket writability
# <翻译结束>


<原文开始>
// ReadMsg wraps the WSARecvMsg network call.
<原文结束>

# <翻译开始>
// ReadMsg wraps the WSARecvMsg network call.
# <翻译结束>


<原文开始>
// ReadMsgInet4 is ReadMsg, but specialized to return a syscall.SockaddrInet4.
<原文结束>

# <翻译开始>
// ReadMsgInet4 is ReadMsg, but specialized to return a syscall.SockaddrInet4.
# <翻译结束>


<原文开始>
// ReadMsgInet6 is ReadMsg, but specialized to return a syscall.SockaddrInet6.
<原文结束>

# <翻译开始>
// ReadMsgInet6 is ReadMsg, but specialized to return a syscall.SockaddrInet6.
# <翻译结束>


<原文开始>
// WriteMsg wraps the WSASendMsg network call.
<原文结束>

# <翻译开始>
// WriteMsg wraps the WSASendMsg network call.
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

