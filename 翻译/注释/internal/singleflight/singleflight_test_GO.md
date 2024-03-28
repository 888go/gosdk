
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
// pump; make available for any future calls
<原文结束>

# <翻译开始>
// pump; make available for any future calls
# <翻译结束>


<原文开始>
// let more goroutines enter Do
<原文结束>

# <翻译开始>
// let more goroutines enter Do
# <翻译结束>


<原文开始>
	// At least one goroutine is in fn now and all of them have at
	// least reached the line before the Do.
<原文结束>

# <翻译开始>
	// At least one goroutine is in fn now and all of them have at
	// least reached the line before the Do.
# <翻译结束>


<原文开始>
// from this point no two function using same key should be executed concurrently
<原文结束>

# <翻译开始>
// from this point no two function using same key should be executed concurrently
# <翻译结束>


<原文开始>
			// The goroutines didn't park in g.Do in time,
			// so the key was re-added and may have been shared after the call.
			// Try again with more time to park.
<原文结束>

# <翻译开始>
			// The goroutines didn't park in g.Do in time,
			// so the key was re-added and may have been shared after the call.
			// Try again with more time to park.
# <翻译结束>


<原文开始>
		// All of the Do calls ended up sharing the first
		// invocation, so the key should have been unused
		// (and therefore unshared) when they returned.
<原文结束>

# <翻译开始>
		// All of the Do calls ended up sharing the first
		// invocation, so the key should have been unused
		// (and therefore unshared) when they returned.
# <翻译结束>

