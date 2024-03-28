
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
// Var decls must have a type or an initializer.
<原文结束>

# <翻译开始>
// Var decls must have a type or an initializer.
# <翻译结束>


<原文开始>
// The initializer must be an expression.
<原文结束>

# <翻译开始>
// The initializer must be an expression.
# <翻译结束>


<原文开始>
// Identifier and expression arity must match.
<原文结束>

# <翻译开始>
// Identifier and expression arity must match.
# <翻译结束>


<原文开始>
// Variables declared in function bodies must be 'used'.
<原文结束>

# <翻译开始>
// Variables declared in function bodies must be 'used'.
# <翻译结束>


<原文开始>
// Unused variables in function literals must lead to only one error (issue #22524).
<原文结束>

# <翻译开始>
// Unused variables in function literals must lead to only one error (issue #22524).
# <翻译结束>


<原文开始>
// Invalid variable declarations must not lead to "declared and not used errors".
// TODO(gri) enable these tests once go/types follows types2 logic for declared and not used variables
// func _() {
//	var a x                        // DISABLED_ERROR undefined: x
//	var b = x                      // DISABLED_ERROR undefined: x
//	var c int = x                  // DISABLED_ERROR undefined: x
//	var d, e, f x                  /* DISABLED_ERROR x */ /* DISABLED_ERROR x */ /* DISABLED_ERROR x */
//	var g, h, i = x, x, x          /* DISABLED_ERROR x */ /* DISABLED_ERROR x */ /* DISABLED_ERROR x */
//	var j, k, l float32 = x, x, x  /* DISABLED_ERROR x */ /* DISABLED_ERROR x */ /* DISABLED_ERROR x */
//	// but no "declared and not used" errors
// }
<原文结束>

# <翻译开始>
// Invalid variable declarations must not lead to "declared and not used errors".
// TODO(gri) enable these tests once go/types follows types2 logic for declared and not used variables
// func _() {
//	var a x                        // DISABLED_ERROR undefined: x
//	var b = x                      // DISABLED_ERROR undefined: x
//	var c int = x                  // DISABLED_ERROR undefined: x
//	var d, e, f x                  /* DISABLED_ERROR x */ /* DISABLED_ERROR x */ /* DISABLED_ERROR x */
//	var g, h, i = x, x, x          /* DISABLED_ERROR x */ /* DISABLED_ERROR x */ /* DISABLED_ERROR x */
//	var j, k, l float32 = x, x, x  /* DISABLED_ERROR x */ /* DISABLED_ERROR x */ /* DISABLED_ERROR x */
//	// but no "declared and not used" errors
// }
# <翻译结束>


<原文开始>
// Invalid (unused) expressions must not lead to spurious "declared and not used errors".
<原文结束>

# <翻译开始>
// Invalid (unused) expressions must not lead to spurious "declared and not used errors".
# <翻译结束>


<原文开始>
// Short variable declarations must declare at least one new non-blank variable.
<原文结束>

# <翻译开始>
// Short variable declarations must declare at least one new non-blank variable.
# <翻译结束>


<原文开始>
// Test case for variables depending on function literals (see also #22992).
<原文结束>

# <翻译开始>
// Test case for variables depending on function literals (see also #22992).
# <翻译结束>


<原文开始>
// The function literal below must not see a.
<原文结束>

# <翻译开始>
// The function literal below must not see a.
# <翻译结束>


<原文开始>
// The function literal below must not see x, y, or z.
<原文结束>

# <翻译开始>
// The function literal below must not see x, y, or z.
# <翻译结束>


<原文开始>
// TODO(gri) consolidate other var decl checks in this file
<原文结束>

# <翻译开始>
// TODO(gri) consolidate other var decl checks in this file
# <翻译结束>

