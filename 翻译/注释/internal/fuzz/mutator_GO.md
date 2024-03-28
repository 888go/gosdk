
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
// scratch slice to avoid additional allocations
<原文结束>

# <翻译开始>
// scratch slice to avoid additional allocations
# <翻译结束>


<原文开始>
// chooseLen chooses length of range mutation in range [1,n]. It gives
// preference to shorter ranges.
<原文结束>

# <翻译开始>
// chooseLen chooses length of range mutation in range [1,n]. It gives
// preference to shorter ranges.
# <翻译结束>


<原文开始>
// mutate performs several mutations on the provided values.
<原文结束>

# <翻译开始>
// mutate performs several mutations on the provided values.
# <翻译结束>


<原文开始>
	// TODO(katiehockman): pull some of these functions into helper methods and
	// test that each case is working as expected.
	// TODO(katiehockman): perform more types of mutations for []byte.
<原文结束>

# <翻译开始>
	// TODO(katiehockman): pull some of these functions into helper methods and
	// test that each case is working as expected.
	// TODO(katiehockman): perform more types of mutations for []byte.
# <翻译结束>


<原文开始>
	// maxPerVal will represent the maximum number of bytes that each value be
	// allowed after mutating, giving an equal amount of capacity to each line.
	// Allow a little wiggle room for the encoding.
<原文结束>

# <翻译开始>
	// maxPerVal will represent the maximum number of bytes that each value be
	// allowed after mutating, giving an equal amount of capacity to each line.
	// Allow a little wiggle room for the encoding.
# <翻译结束>


<原文开始>
	// Pick a random value to mutate.
	// TODO: consider mutating more than one value at a time.
<原文结束>

# <翻译开始>
	// Pick a random value to mutate.
	// TODO: consider mutating more than one value at a time.
# <翻译结束>


<原文开始>
// 50% chance of flipping the bool
<原文结束>

# <翻译开始>
// 50% chance of flipping the bool
# <翻译结束>


<原文开始>
// Don't let v exceed maxValue
<原文结束>

# <翻译开始>
// Don't let v exceed maxValue
# <翻译结束>


<原文开始>
// Subtract a random number
<原文结束>

# <翻译开始>
// Subtract a random number
# <翻译结束>


<原文开始>
// Don't let v drop below -maxValue
<原文结束>

# <翻译开始>
// Don't let v drop below -maxValue
# <翻译结束>


<原文开始>
// Don't let v drop below 0
<原文结束>

# <翻译开始>
// Don't let v drop below 0
# <翻译结束>


<原文开始>
// Multiply by a random number
<原文结束>

# <翻译开始>
// Multiply by a random number
# <翻译结束>


<原文开始>
// Don't let v go beyond the minimum or maximum value
<原文结束>

# <翻译开始>
// Don't let v go beyond the minimum or maximum value
# <翻译结束>


<原文开始>
// Divide by a random number
<原文结束>

# <翻译开始>
// Divide by a random number
# <翻译结束>

