
<原文开始>
// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
# <翻译结束>


<原文开始>
// Close unmaps the shared memory and closes the temporary file. If this
// sharedMem was created with sharedMemTempFile, Close also removes the file.
<原文结束>

# <翻译开始>
// Close unmaps the shared memory and closes the temporary file. If this
// sharedMem was created with sharedMemTempFile, Close also removes the file.
# <翻译结束>


<原文开始>
	// Attempt all operations, even if we get an error for an earlier operation.
	// os.File.Close may fail due to I/O errors, but we still want to delete
	// the temporary file.
<原文结束>

# <翻译开始>
	// Attempt all operations, even if we get an error for an earlier operation.
	// os.File.Close may fail due to I/O errors, but we still want to delete
	// the temporary file.
# <翻译结束>


<原文开始>
// setWorkerComm configures communication channels on the cmd that will
// run a worker process.
<原文结束>

# <翻译开始>
// setWorkerComm configures communication channels on the cmd that will
// run a worker process.
# <翻译结束>


<原文开始>
// getWorkerComm returns communication channels in the worker process.
<原文结束>

# <翻译开始>
// getWorkerComm returns communication channels in the worker process.
# <翻译结束>


<原文开始>
// isInterruptError returns whether an error was returned by a process that
// was terminated by an interrupt signal (SIGINT).
<原文结束>

# <翻译开始>
// isInterruptError returns whether an error was returned by a process that
// was terminated by an interrupt signal (SIGINT).
# <翻译结束>


<原文开始>
// terminationSignal checks if err is an exec.ExitError with a signal status.
// If it is, terminationSignal returns the signal and true.
// If not, -1 and false.
<原文结束>

# <翻译开始>
// terminationSignal checks if err is an exec.ExitError with a signal status.
// If it is, terminationSignal returns the signal and true.
// If not, -1 and false.
# <翻译结束>


<原文开始>
// isCrashSignal returns whether a signal was likely to have been caused by an
// error in the program that received it, triggered by a fuzz input. For
// example, SIGSEGV would be received after a nil pointer dereference.
// Other signals like SIGKILL or SIGHUP are more likely to have been sent by
// another process, and we shouldn't record a crasher if the worker process
// receives one of these.
//
// Note that Go installs its own signal handlers on startup, so some of these
// signals may only be received if signal handlers are changed. For example,
// SIGSEGV is normally transformed into a panic that causes the process to exit
// with status 2 if not recovered, which we handle as a crash.
<原文结束>

# <翻译开始>
// isCrashSignal returns whether a signal was likely to have been caused by an
// error in the program that received it, triggered by a fuzz input. For
// example, SIGSEGV would be received after a nil pointer dereference.
// Other signals like SIGKILL or SIGHUP are more likely to have been sent by
// another process, and we shouldn't record a crasher if the worker process
// receives one of these.
//
// Note that Go installs its own signal handlers on startup, so some of these
// signals may only be received if signal handlers are changed. For example,
// SIGSEGV is normally transformed into a panic that causes the process to exit
// with status 2 if not recovered, which we handle as a crash.
# <翻译结束>


<原文开始>
// invalid memory access (e.g., misaligned address)
<原文结束>

# <翻译开始>
// invalid memory access (e.g., misaligned address)
# <翻译结束>


<原文开始>
// math error, e.g., integer divide by zero
<原文结束>

# <翻译开始>
// math error, e.g., integer divide by zero
# <翻译结束>


<原文开始>
// invalid memory access (e.g., write to read-only)
<原文结束>

# <翻译开始>
// invalid memory access (e.g., write to read-only)
# <翻译结束>


<原文开始>
// sent data to closed pipe or socket
<原文结束>

# <翻译开始>
// sent data to closed pipe or socket
# <翻译结束>

