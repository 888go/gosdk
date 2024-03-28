
<原文开始>
// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
# <翻译结束>


<原文开始>
	// If minimization was successful at any point during minimizeBytes,
	// then the vals slice in (*workerServer).minimizeInput will point to
	// tmp. Since tmp is altered while making new candidates, we need to
	// make sure that it is equal to the correct value, v, before exiting
	// this function.
<原文结束>

# <翻译开始>
	// If minimization was successful at any point during minimizeBytes,
	// then the vals slice in (*workerServer).minimizeInput will point to
	// tmp. Since tmp is altered while making new candidates, we need to
	// make sure that it is equal to the correct value, v, before exiting
	// this function.
# <翻译结束>


<原文开始>
// First, try to cut the tail.
<原文结束>

# <翻译开始>
// First, try to cut the tail.
# <翻译结束>


<原文开始>
// Set v to the new value to continue iterating.
<原文结束>

# <翻译开始>
// Set v to the new value to continue iterating.
# <翻译结束>


<原文开始>
// Then, try to remove each individual byte.
<原文结束>

# <翻译开始>
// Then, try to remove each individual byte.
# <翻译结束>


<原文开始>
// Update v to delete the value at index i.
<原文结束>

# <翻译开始>
// Update v to delete the value at index i.
# <翻译结束>


<原文开始>
		// v[i] is now different, so decrement i to redo this iteration
		// of the loop with the new value.
<原文结束>

# <翻译开始>
		// v[i] is now different, so decrement i to redo this iteration
		// of the loop with the new value.
# <翻译结束>


<原文开始>
// Then, try to remove each possible subset of bytes.
<原文结束>

# <翻译开始>
// Then, try to remove each possible subset of bytes.
# <翻译结束>


<原文开始>
// Update v and reset the loop with the new length.
<原文结束>

# <翻译开始>
// Update v and reset the loop with the new length.
# <翻译结束>


<原文开始>
	// Then, try to make it more simplified and human-readable by trying to replace each
	// byte with a printable character.
<原文结束>

# <翻译开始>
	// Then, try to make it more simplified and human-readable by trying to replace each
	// byte with a printable character.
# <翻译结束>


<原文开始>
// Successful. Move on to the next byte in v.
<原文结束>

# <翻译开始>
// Successful. Move on to the next byte in v.
# <翻译结束>


<原文开始>
// Unsuccessful. Revert v[i] back to original value.
<原文结束>

# <翻译开始>
// Unsuccessful. Revert v[i] back to original value.
# <翻译结束>

