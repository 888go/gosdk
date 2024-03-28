
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
// Package poll supports non-blocking I/O on file descriptors with polling.
// This supports I/O operations that block only a goroutine, not a thread.
// This is used by the net and os packages.
// It uses a poller built into the runtime, with support from the
// runtime scheduler.
<原文结束>

# <翻译开始>
// Package poll supports non-blocking I/O on file descriptors with polling.
// This supports I/O operations that block only a goroutine, not a thread.
// This is used by the net and os packages.
// It uses a poller built into the runtime, with support from the
// runtime scheduler.
# <翻译结束>


<原文开始>
// errNetClosing is the type of the variable ErrNetClosing.
// This is used to implement the net.Error interface.
<原文结束>

# <翻译开始>
// errNetClosing is the type of the variable ErrNetClosing.
// This is used to implement the net.Error interface.
# <翻译结束>


<原文开始>
// Error returns the error message for ErrNetClosing.
// Keep this string consistent because of issue #4373:
// since historically programs have not been able to detect
// this error, they look for the string.
<原文结束>

# <翻译开始>
// Error returns the error message for ErrNetClosing.
// Keep this string consistent because of issue #4373:
// since historically programs have not been able to detect
// this error, they look for the string.
# <翻译结束>


<原文开始>
// ErrNetClosing is returned when a network descriptor is used after
// it has been closed.
<原文结束>

# <翻译开始>
// ErrNetClosing is returned when a network descriptor is used after
// it has been closed.
# <翻译结束>


<原文开始>
// ErrFileClosing is returned when a file descriptor is used after it
// has been closed.
<原文结束>

# <翻译开始>
// ErrFileClosing is returned when a file descriptor is used after it
// has been closed.
# <翻译结束>


<原文开始>
// ErrNoDeadline is returned when a request is made to set a deadline
// on a file type that does not use the poller.
<原文结束>

# <翻译开始>
// ErrNoDeadline is returned when a request is made to set a deadline
// on a file type that does not use the poller.
# <翻译结束>


<原文开始>
// Return the appropriate closing error based on isFile.
<原文结束>

# <翻译开始>
// Return the appropriate closing error based on isFile.
# <翻译结束>


<原文开始>
// ErrDeadlineExceeded is returned for an expired deadline.
// This is exported by the os package as os.ErrDeadlineExceeded.
<原文结束>

# <翻译开始>
// ErrDeadlineExceeded is returned for an expired deadline.
// This is exported by the os package as os.ErrDeadlineExceeded.
# <翻译结束>


<原文开始>
// DeadlineExceededError is returned for an expired deadline.
<原文结束>

# <翻译开始>
// DeadlineExceededError is returned for an expired deadline.
# <翻译结束>


<原文开始>
// Implement the net.Error interface.
// The string is "i/o timeout" because that is what was returned
// by earlier Go versions. Changing it may break programs that
// match on error strings.
<原文结束>

# <翻译开始>
// Implement the net.Error interface.
// The string is "i/o timeout" because that is what was returned
// by earlier Go versions. Changing it may break programs that
// match on error strings.
# <翻译结束>


<原文开始>
// ErrNotPollable is returned when the file or socket is not suitable
// for event notification.
<原文结束>

# <翻译开始>
// ErrNotPollable is returned when the file or socket is not suitable
// for event notification.
# <翻译结束>


<原文开始>
// consume removes data from a slice of byte slices, for writev.
<原文结束>

# <翻译开始>
// consume removes data from a slice of byte slices, for writev.
# <翻译结束>


<原文开始>
// TestHookDidWritev is a hook for testing writev.
<原文结束>

# <翻译开始>
// TestHookDidWritev is a hook for testing writev.
# <翻译结束>

