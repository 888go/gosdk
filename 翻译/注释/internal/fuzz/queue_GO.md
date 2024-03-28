
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
// queue holds a growable sequence of inputs for fuzzing and minimization.
//
// For now, this is a simple ring buffer
// (https://en.wikipedia.org/wiki/Circular_buffer).
//
// TODO(golang.org/issue/46224): use a prioritization algorithm based on input
// size, previous duration, coverage, and any other metrics that seem useful.
<原文结束>

# <翻译开始>
// queue holds a growable sequence of inputs for fuzzing and minimization.
//
// For now, this is a simple ring buffer
// (https://en.wikipedia.org/wiki/Circular_buffer).
//
// TODO(golang.org/issue/46224): use a prioritization algorithm based on input
// size, previous duration, coverage, and any other metrics that seem useful.
# <翻译结束>


<原文开始>
	// elems holds a ring buffer.
	// The queue is empty when begin = end.
	// The queue is full (until grow is called) when end = begin + N - 1 (mod N)
	// where N = cap(elems).
<原文结束>

# <翻译开始>
	// elems holds a ring buffer.
	// The queue is empty when begin = end.
	// The queue is full (until grow is called) when end = begin + N - 1 (mod N)
	// where N = cap(elems).
# <翻译结束>

