
<原文开始>
// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
# <翻译结束>


<原文开始>
	// Use function type inference to infer type A2 for T.
	// Don't use constraint type inference before function
	// type inference for typed arguments, otherwise it would
	// infer type [2]uint64 for T which doesn't have method m
	// (was the bug).
<原文结束>

# <翻译开始>
	// Use function type inference to infer type A2 for T.
	// Don't use constraint type inference before function
	// type inference for typed arguments, otherwise it would
	// infer type [2]uint64 for T which doesn't have method m
	// (was the bug).
# <翻译结束>


<原文开始>
// Keep using constraint type inference before function type
// inference for untyped arguments so we infer type float64
// for E below, and not int (which would not work).
<原文结束>

# <翻译开始>
// Keep using constraint type inference before function type
// inference for untyped arguments so we infer type float64
// for E below, and not int (which would not work).
# <翻译结束>


<原文开始>
// Keep using constraint type inference after function
// type inference for untyped arguments so we infer
// missing type arguments for which we only have the
// untyped arguments as starting point.
<原文结束>

# <翻译开始>
// Keep using constraint type inference after function
// type inference for untyped arguments so we infer
// missing type arguments for which we only have the
// untyped arguments as starting point.
# <翻译结束>

