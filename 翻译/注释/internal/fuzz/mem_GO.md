
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
// sharedMem manages access to a region of virtual memory mapped from a file,
// shared between multiple processes. The region includes space for a header and
// a value of variable length.
//
// When fuzzing, the coordinator creates a sharedMem from a temporary file for
// each worker. This buffer is used to pass values to fuzz between processes.
// Care must be taken to manage access to shared memory across processes;
// sharedMem provides no synchronization on its own. See workerComm for an
// explanation.
<原文结束>

# <翻译开始>
// sharedMem manages access to a region of virtual memory mapped from a file,
// shared between multiple processes. The region includes space for a header and
// a value of variable length.
//
// When fuzzing, the coordinator creates a sharedMem from a temporary file for
// each worker. This buffer is used to pass values to fuzz between processes.
// Care must be taken to manage access to shared memory across processes;
// sharedMem provides no synchronization on its own. See workerComm for an
// explanation.
# <翻译结束>


<原文开始>
// f is the file mapped into memory.
<原文结束>

# <翻译开始>
// f is the file mapped into memory.
# <翻译结束>


<原文开始>
	// region is the mapped region of virtual memory for f. The content of f may
	// be read or written through this slice.
<原文结束>

# <翻译开始>
	// region is the mapped region of virtual memory for f. The content of f may
	// be read or written through this slice.
# <翻译结束>


<原文开始>
// removeOnClose is true if the file should be deleted by Close.
<原文结束>

# <翻译开始>
// removeOnClose is true if the file should be deleted by Close.
# <翻译结束>


<原文开始>
// sys contains OS-specific information.
<原文结束>

# <翻译开始>
// sys contains OS-specific information.
# <翻译结束>


<原文开始>
// sharedMemHeader stores metadata in shared memory.
<原文结束>

# <翻译开始>
// sharedMemHeader stores metadata in shared memory.
# <翻译结束>


<原文开始>
	// count is the number of times the worker has called the fuzz function.
	// May be reset by coordinator.
<原文结束>

# <翻译开始>
	// count is the number of times the worker has called the fuzz function.
	// May be reset by coordinator.
# <翻译结束>


<原文开始>
// valueLen is the number of bytes in region which should be read.
<原文结束>

# <翻译开始>
// valueLen is the number of bytes in region which should be read.
# <翻译结束>


<原文开始>
// randState and randInc hold the state of a pseudo-random number generator.
<原文结束>

# <翻译开始>
// randState and randInc hold the state of a pseudo-random number generator.
# <翻译结束>


<原文开始>
	// rawInMem is true if the region holds raw bytes, which occurs during
	// minimization. If true after the worker fails during minimization, this
	// indicates that an unrecoverable error occurred, and the region can be
	// used to retrieve the raw bytes that caused the error.
<原文结束>

# <翻译开始>
	// rawInMem is true if the region holds raw bytes, which occurs during
	// minimization. If true after the worker fails during minimization, this
	// indicates that an unrecoverable error occurred, and the region can be
	// used to retrieve the raw bytes that caused the error.
# <翻译结束>


<原文开始>
// sharedMemSize returns the size needed for a shared memory buffer that can
// contain values of the given size.
<原文结束>

# <翻译开始>
// sharedMemSize returns the size needed for a shared memory buffer that can
// contain values of the given size.
# <翻译结束>


<原文开始>
// TODO(jayconrod): set a reasonable maximum size per platform.
<原文结束>

# <翻译开始>
// TODO(jayconrod): set a reasonable maximum size per platform.
# <翻译结束>


<原文开始>
// sharedMemTempFile creates a new temporary file of the given size, then maps
// it into memory. The file will be removed when the Close method is called.
<原文结束>

# <翻译开始>
// sharedMemTempFile creates a new temporary file of the given size, then maps
// it into memory. The file will be removed when the Close method is called.
# <翻译结束>


<原文开始>
// Create a temporary file.
<原文结束>

# <翻译开始>
// Create a temporary file.
# <翻译结束>


<原文开始>
// Resize it to the correct size.
<原文结束>

# <翻译开始>
// Resize it to the correct size.
# <翻译结束>


<原文开始>
// Map the file into memory.
<原文结束>

# <翻译开始>
// Map the file into memory.
# <翻译结束>


<原文开始>
// header returns a pointer to metadata within the shared memory region.
<原文结束>

# <翻译开始>
// header returns a pointer to metadata within the shared memory region.
# <翻译结束>


<原文开始>
// valueRef returns the value currently stored in shared memory. The returned
// slice points to shared memory; it is not a copy.
<原文结束>

# <翻译开始>
// valueRef returns the value currently stored in shared memory. The returned
// slice points to shared memory; it is not a copy.
# <翻译结束>


<原文开始>
// valueCopy returns a copy of the value stored in shared memory.
<原文结束>

# <翻译开始>
// valueCopy returns a copy of the value stored in shared memory.
# <翻译结束>


<原文开始>
// setValue copies the data in b into the shared memory buffer and sets
// the length. len(b) must be less than or equal to the capacity of the buffer
// (as returned by cap(m.value())).
<原文结束>

# <翻译开始>
// setValue copies the data in b into the shared memory buffer and sets
// the length. len(b) must be less than or equal to the capacity of the buffer
// (as returned by cap(m.value())).
# <翻译结束>


<原文开始>
// setValueLen sets the length of the shared memory buffer returned by valueRef
// to n, which may be at most the cap of that slice.
//
// Note that we can only store the length in the shared memory header. The full
// slice header contains a pointer, which is likely only valid for one process,
// since each process can map shared memory at a different virtual address.
<原文结束>

# <翻译开始>
// setValueLen sets the length of the shared memory buffer returned by valueRef
// to n, which may be at most the cap of that slice.
//
// Note that we can only store the length in the shared memory header. The full
// slice header contains a pointer, which is likely only valid for one process,
// since each process can map shared memory at a different virtual address.
# <翻译结束>


<原文开始>
// TODO(jayconrod): add method to resize the buffer. We'll need that when the
// mutator can increase input length. Only the coordinator will be able to
// do it, since we'll need to send a message to the worker telling it to
// remap the file.
<原文结束>

# <翻译开始>
// TODO(jayconrod): add method to resize the buffer. We'll need that when the
// mutator can increase input length. Only the coordinator will be able to
// do it, since we'll need to send a message to the worker telling it to
// remap the file.
# <翻译结束>

