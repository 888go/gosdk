
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
// Package singleflight provides a duplicate function call suppression
// mechanism.
<原文结束>

# <翻译开始>
// Package singleflight provides a duplicate function call suppression
// mechanism.
# <翻译结束>


<原文开始>
// call is an in-flight or completed singleflight.Do call
<原文结束>

# <翻译开始>
// call is an in-flight or completed singleflight.Do call
# <翻译结束>


<原文开始>
	// These fields are written once before the WaitGroup is done
	// and are only read after the WaitGroup is done.
<原文结束>

# <翻译开始>
	// These fields are written once before the WaitGroup is done
	// and are only read after the WaitGroup is done.
# <翻译结束>


<原文开始>
	// These fields are read and written with the singleflight
	// mutex held before the WaitGroup is done, and are read but
	// not written after the WaitGroup is done.
<原文结束>

# <翻译开始>
	// These fields are read and written with the singleflight
	// mutex held before the WaitGroup is done, and are read but
	// not written after the WaitGroup is done.
# <翻译结束>


<原文开始>
// Group represents a class of work and forms a namespace in
// which units of work can be executed with duplicate suppression.
<原文结束>

# <翻译开始>
// Group represents a class of work and forms a namespace in
// which units of work can be executed with duplicate suppression.
# <翻译结束>


<原文开始>
// Result holds the results of Do, so they can be passed
// on a channel.
<原文结束>

# <翻译开始>
// Result holds the results of Do, so they can be passed
// on a channel.
# <翻译结束>


<原文开始>
// Do executes and returns the results of the given function, making
// sure that only one execution is in-flight for a given key at a
// time. If a duplicate comes in, the duplicate caller waits for the
// original to complete and receives the same results.
// The return value shared indicates whether v was given to multiple callers.
<原文结束>

# <翻译开始>
// Do executes and returns the results of the given function, making
// sure that only one execution is in-flight for a given key at a
// time. If a duplicate comes in, the duplicate caller waits for the
// original to complete and receives the same results.
// The return value shared indicates whether v was given to multiple callers.
# <翻译结束>


<原文开始>
// DoChan is like Do but returns a channel that will receive the
// results when they are ready.
<原文结束>

# <翻译开始>
// DoChan is like Do but returns a channel that will receive the
// results when they are ready.
# <翻译结束>


<原文开始>
// doCall handles the single call for a key.
<原文结束>

# <翻译开始>
// doCall handles the single call for a key.
# <翻译结束>


<原文开始>
// ForgetUnshared tells the singleflight to forget about a key if it is not
// shared with any other goroutines. Future calls to Do for a forgotten key
// will call the function rather than waiting for an earlier call to complete.
// Returns whether the key was forgotten or unknown--that is, whether no
// other goroutines are waiting for the result.
<原文结束>

# <翻译开始>
// ForgetUnshared tells the singleflight to forget about a key if it is not
// shared with any other goroutines. Future calls to Do for a forgotten key
// will call the function rather than waiting for an earlier call to complete.
// Returns whether the key was forgotten or unknown--that is, whether no
// other goroutines are waiting for the result.
# <翻译结束>

