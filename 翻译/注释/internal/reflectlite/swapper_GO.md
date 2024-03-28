
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
// Swapper returns a function that swaps the elements in the provided
// slice.
//
// Swapper panics if the provided interface is not a slice.
<原文结束>

# <翻译开始>
// Swapper returns a function that swaps the elements in the provided
// slice.
//
// Swapper panics if the provided interface is not a slice.
# <翻译结束>


<原文开始>
// Fast path for slices of size 0 and 1. Nothing to swap.
<原文结束>

# <翻译开始>
// Fast path for slices of size 0 and 1. Nothing to swap.
# <翻译结束>


<原文开始>
// Some common & small cases, without using memmove:
<原文结束>

# <翻译开始>
// Some common & small cases, without using memmove:
# <翻译结束>

