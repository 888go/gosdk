
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
// Create a file mapping object. The object itself is not shared.
<原文结束>

# <翻译开始>
// Create a file mapping object. The object itself is not shared.
# <翻译结束>


<原文开始>
// Create a view from the file mapping object.
<原文结束>

# <翻译开始>
// Create a view from the file mapping object.
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
	// On Windows, we can't tell whether the process was interrupted by the error
	// returned by Wait. It looks like an ExitError with status 1.
<原文结束>

# <翻译开始>
	// On Windows, we can't tell whether the process was interrupted by the error
	// returned by Wait. It looks like an ExitError with status 1.
# <翻译结束>


<原文开始>
// terminationSignal returns -1 and false because Windows doesn't have signals.
<原文结束>

# <翻译开始>
// terminationSignal returns -1 and false because Windows doesn't have signals.
# <翻译结束>


<原文开始>
// isCrashSignal is not implemented because Windows doesn't have signals.
<原文结束>

# <翻译开始>
// isCrashSignal is not implemented because Windows doesn't have signals.
# <翻译结束>

