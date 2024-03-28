
<原文开始>
// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
# <翻译结束>


<原文开始>
// spliceNonblock makes calls to splice(2) non-blocking.
<原文结束>

# <翻译开始>
// spliceNonblock makes calls to splice(2) non-blocking.
# <翻译结束>


<原文开始>
	// maxSpliceSize is the maximum amount of data Splice asks
	// the kernel to move in a single call to splice(2).
	// We use 1MB as Splice writes data through a pipe, and 1MB is the default maximum pipe buffer size,
	// which is determined by /proc/sys/fs/pipe-max-size.
<原文结束>

# <翻译开始>
	// maxSpliceSize is the maximum amount of data Splice asks
	// the kernel to move in a single call to splice(2).
	// We use 1MB as Splice writes data through a pipe, and 1MB is the default maximum pipe buffer size,
	// which is determined by /proc/sys/fs/pipe-max-size.
# <翻译结束>


<原文开始>
// Splice transfers at most remain bytes of data from src to dst, using the
// splice system call to minimize copies of data from and to userspace.
//
// Splice gets a pipe buffer from the pool or creates a new one if needed, to serve as a buffer for the data transfer.
// src and dst must both be stream-oriented sockets.
//
// If err != nil, sc is the system call which caused the error.
<原文结束>

# <翻译开始>
// Splice transfers at most remain bytes of data from src to dst, using the
// splice system call to minimize copies of data from and to userspace.
//
// Splice gets a pipe buffer from the pool or creates a new one if needed, to serve as a buffer for the data transfer.
// src and dst must both be stream-oriented sockets.
//
// If err != nil, sc is the system call which caused the error.
# <翻译结束>


<原文开始>
		// The operation is considered handled if splice returns no
		// error, or an error other than EINVAL. An EINVAL means the
		// kernel does not support splice for the socket type of src.
		// The failed syscall does not consume any data so it is safe
		// to fall back to a generic copy.
		//
		// spliceDrain should never return EAGAIN, so if err != nil,
		// Splice cannot continue.
		//
		// If inPipe == 0 && err == nil, src is at EOF, and the
		// transfer is complete.
<原文结束>

# <翻译开始>
		// The operation is considered handled if splice returns no
		// error, or an error other than EINVAL. An EINVAL means the
		// kernel does not support splice for the socket type of src.
		// The failed syscall does not consume any data so it is safe
		// to fall back to a generic copy.
		//
		// spliceDrain should never return EAGAIN, so if err != nil,
		// Splice cannot continue.
		//
		// If inPipe == 0 && err == nil, src is at EOF, and the
		// transfer is complete.
# <翻译结束>


<原文开始>
// spliceDrain moves data from a socket to a pipe.
//
// Invariant: when entering spliceDrain, the pipe is empty. It is either in its
// initial state, or splicePump has emptied it previously.
//
// Given this, spliceDrain can reasonably assume that the pipe is ready for
// writing, so if splice returns EAGAIN, it must be because the socket is not
// ready for reading.
//
// If spliceDrain returns (0, nil), src is at EOF.
<原文结束>

# <翻译开始>
// spliceDrain moves data from a socket to a pipe.
//
// Invariant: when entering spliceDrain, the pipe is empty. It is either in its
// initial state, or splicePump has emptied it previously.
//
// Given this, spliceDrain can reasonably assume that the pipe is ready for
// writing, so if splice returns EAGAIN, it must be because the socket is not
// ready for reading.
//
// If spliceDrain returns (0, nil), src is at EOF.
# <翻译结束>


<原文开始>
// splicePump moves all the buffered data from a pipe to a socket.
//
// Invariant: when entering splicePump, there are exactly inPipe
// bytes of data in the pipe, from a previous call to spliceDrain.
//
// By analogy to the condition from spliceDrain, splicePump
// only needs to poll the socket for readiness, if splice returns
// EAGAIN.
//
// If splicePump cannot move all the data in a single call to
// splice(2), it loops over the buffered data until it has written
// all of it to the socket. This behavior is similar to the Write
// step of an io.Copy in userspace.
<原文结束>

# <翻译开始>
// splicePump moves all the buffered data from a pipe to a socket.
//
// Invariant: when entering splicePump, there are exactly inPipe
// bytes of data in the pipe, from a previous call to spliceDrain.
//
// By analogy to the condition from spliceDrain, splicePump
// only needs to poll the socket for readiness, if splice returns
// EAGAIN.
//
// If splicePump cannot move all the data in a single call to
// splice(2), it loops over the buffered data until it has written
// all of it to the socket. This behavior is similar to the Write
// step of an io.Copy in userspace.
# <翻译结束>


<原文开始>
		// Here, the condition n == 0 && err == nil should never be
		// observed, since Splice controls the write side of the pipe.
<原文结束>

# <翻译开始>
		// Here, the condition n == 0 && err == nil should never be
		// observed, since Splice controls the write side of the pipe.
# <翻译结束>


<原文开始>
// splice wraps the splice system call. Since the current implementation
// only uses splice on sockets and pipes, the offset arguments are unused.
// splice returns int instead of int64, because callers never ask it to
// move more data in a single call than can fit in an int32.
<原文结束>

# <翻译开始>
// splice wraps the splice system call. Since the current implementation
// only uses splice on sockets and pipes, the offset arguments are unused.
// splice returns int instead of int64, because callers never ask it to
// move more data in a single call than can fit in an int32.
# <翻译结束>


<原文开始>
	// We want to use a finalizer, so ensure that the size is
	// large enough to not use the tiny allocator.
<原文结束>

# <翻译开始>
	// We want to use a finalizer, so ensure that the size is
	// large enough to not use the tiny allocator.
# <翻译结束>


<原文开始>
// splicePipePool caches pipes to avoid high-frequency construction and destruction of pipe buffers.
// The garbage collector will free all pipes in the sync.Pool periodically, thus we need to set up
// a finalizer for each pipe to close its file descriptors before the actual GC.
<原文结束>

# <翻译开始>
// splicePipePool caches pipes to avoid high-frequency construction and destruction of pipe buffers.
// The garbage collector will free all pipes in the sync.Pool periodically, thus we need to set up
// a finalizer for each pipe to close its file descriptors before the actual GC.
# <翻译结束>


<原文开始>
	// Discard the error which occurred during the creation of pipe buffer,
	// redirecting the data transmission to the conventional way utilizing read() + write() as a fallback.
<原文结束>

# <翻译开始>
	// Discard the error which occurred during the creation of pipe buffer,
	// redirecting the data transmission to the conventional way utilizing read() + write() as a fallback.
# <翻译结束>


<原文开始>
// getPipe tries to acquire a pipe buffer from the pool or create a new one with newPipe() if it gets nil from the cache.
//
// Note that it may fail to create a new pipe buffer by newPipe(), in which case getPipe() will return a generic error
// and system call name splice in a string as the indication.
<原文结束>

# <翻译开始>
// getPipe tries to acquire a pipe buffer from the pool or create a new one with newPipe() if it gets nil from the cache.
//
// Note that it may fail to create a new pipe buffer by newPipe(), in which case getPipe() will return a generic error
// and system call name splice in a string as the indication.
# <翻译结束>


<原文开始>
	// If there is still data left in the pipe,
	// then close and discard it instead of putting it back into the pool.
<原文结束>

# <翻译开始>
	// If there is still data left in the pipe,
	// then close and discard it instead of putting it back into the pool.
# <翻译结束>


<原文开始>
// newPipe sets up a pipe for a splice operation.
<原文结束>

# <翻译开始>
// newPipe sets up a pipe for a splice operation.
# <翻译结束>


<原文开始>
	// Splice will loop writing maxSpliceSize bytes from the source to the pipe,
	// and then write those bytes from the pipe to the destination.
	// Set the pipe buffer size to maxSpliceSize to optimize that.
	// Ignore errors here, as a smaller buffer size will work,
	// although it will require more system calls.
<原文结束>

# <翻译开始>
	// Splice will loop writing maxSpliceSize bytes from the source to the pipe,
	// and then write those bytes from the pipe to the destination.
	// Set the pipe buffer size to maxSpliceSize to optimize that.
	// Ignore errors here, as a smaller buffer size will work,
	// although it will require more system calls.
# <翻译结束>


<原文开始>
// destroyPipe destroys a pipe.
<原文结束>

# <翻译开始>
// destroyPipe destroys a pipe.
# <翻译结束>

