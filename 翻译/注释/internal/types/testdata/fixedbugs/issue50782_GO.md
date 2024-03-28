
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
// Field accesses through type parameters are disabled
// until we have a more thorough understanding of the
// implications on the spec. See issue #51576.
<原文结束>

# <翻译开始>
// Field accesses through type parameters are disabled
// until we have a more thorough understanding of the
// implications on the spec. See issue #51576.
# <翻译结束>


<原文开始>
// The first example from the issue.
<原文结束>

# <翻译开始>
// The first example from the issue.
# <翻译结束>


<原文开始>
// numericAbs matches numeric types with an Abs method.
<原文结束>

# <翻译开始>
// numericAbs matches numeric types with an Abs method.
# <翻译结束>


<原文开始>
// AbsDifference computes the absolute value of the difference of
// a and b, where the absolute value is determined by the Abs method.
<原文结束>

# <翻译开始>
// AbsDifference computes the absolute value of the difference of
// a and b, where the absolute value is determined by the Abs method.
# <翻译结束>


<原文开始>
	// Field accesses are not permitted for now. Keep an error so
	// we can find and fix this code once the situation changes.
<原文结束>

# <翻译开始>
	// Field accesses are not permitted for now. Keep an error so
	// we can find and fix this code once the situation changes.
# <翻译结束>


<原文开始>
// ERROR a\.Value undefined
<原文结束>

# <翻译开始>
// ERROR a\.Value undefined
# <翻译结束>


<原文开始>
	// TODO: The error below should probably be positioned on the '-'.
	// d := a /* ERROR "invalid operation: operator - not defined" */ .Value - b.Value
	// return d.Abs()
<原文结束>

# <翻译开始>
	// TODO: The error below should probably be positioned on the '-'.
	// d := a /* ERROR "invalid operation: operator - not defined" */ .Value - b.Value
	// return d.Abs()
# <翻译结束>


<原文开始>
// The second example from the issue.
<原文结束>

# <翻译开始>
// The second example from the issue.
# <翻译结束>

