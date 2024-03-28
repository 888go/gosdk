
<原文开始>
// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
# <翻译结束>


<原文开始>
// asyncIO implements asynchronous cancelable I/O.
// An asyncIO represents a single asynchronous Read or Write
// operation. The result is returned on the result channel.
// The undergoing I/O system call can either complete or be
// interrupted by a note.
<原文结束>

# <翻译开始>
// asyncIO implements asynchronous cancelable I/O.
// An asyncIO represents a single asynchronous Read or Write
// operation. The result is returned on the result channel.
// The undergoing I/O system call can either complete or be
// interrupted by a note.
# <翻译结束>


<原文开始>
// mu guards the pid field.
<原文结束>

# <翻译开始>
// mu guards the pid field.
# <翻译结束>


<原文开始>
	// pid holds the process id of
	// the process running the IO operation.
<原文结束>

# <翻译开始>
	// pid holds the process id of
	// the process running the IO operation.
# <翻译结束>


<原文开始>
// result is the return value of a Read or Write operation.
<原文结束>

# <翻译开始>
// result is the return value of a Read or Write operation.
# <翻译结束>


<原文开始>
// newAsyncIO returns a new asyncIO that performs an I/O
// operation by calling fn, which must do one and only one
// interruptible system call.
<原文结束>

# <翻译开始>
// newAsyncIO returns a new asyncIO that performs an I/O
// operation by calling fn, which must do one and only one
// interruptible system call.
# <翻译结束>


<原文开始>
		// Lock the current goroutine to its process
		// and store the pid in io so that Cancel can
		// interrupt it. We ignore the "hangup" signal,
		// so the signal does not take down the entire
		// Go runtime.
<原文结束>

# <翻译开始>
		// Lock the current goroutine to its process
		// and store the pid in io so that Cancel can
		// interrupt it. We ignore the "hangup" signal,
		// so the signal does not take down the entire
		// Go runtime.
# <翻译结束>


<原文开始>
// Cancel interrupts the I/O operation, causing
// the Wait function to return.
<原文结束>

# <翻译开始>
// Cancel interrupts the I/O operation, causing
// the Wait function to return.
# <翻译结束>


<原文开始>
// Wait for the I/O operation to complete.
<原文结束>

# <翻译开始>
// Wait for the I/O operation to complete.
# <翻译结束>


<原文开始>
// The following functions, provided by the runtime, are used to
// ignore and unignore the "hangup" signal received by the process.
<原文结束>

# <翻译开始>
// The following functions, provided by the runtime, are used to
// ignore and unignore the "hangup" signal received by the process.
# <翻译结束>

