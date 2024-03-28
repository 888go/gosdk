
<原文开始>
// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
# <翻译结束>


<原文开始>
// must refer to imported fmt rather than the fmt below
<原文结束>

# <翻译开始>
// must refer to imported fmt rather than the fmt below
# <翻译结束>


<原文开始>
// Check that a missing identifier doesn't lead to a spurious error cascade.
<原文结束>

# <翻译开始>
// Check that a missing identifier doesn't lead to a spurious error cascade.
# <翻译结束>


<原文开始>
// no error for composite literal based on unknown type
<原文结束>

# <翻译开始>
// no error for composite literal based on unknown type
# <翻译结束>


<原文开始>
// variadic builtin function
<原文结束>

# <翻译开始>
// variadic builtin function
# <翻译结束>


<原文开始>
// variadic user-defined function
<原文结束>

# <翻译开始>
// variadic user-defined function
# <翻译结束>


<原文开始>
// Check that embedding a non-interface type in an interface results in a good error message.
<原文结束>

# <翻译开始>
// Check that embedding a non-interface type in an interface results in a good error message.
# <翻译结束>


<原文开始>
// issue11347
// These should not crash.
<原文结束>

# <翻译开始>
// issue11347
// These should not crash.
# <翻译结束>


<原文开始>
// issue10260
// Check that error messages explain reason for interface assignment failures.
<原文结束>

# <翻译开始>
// issue10260
// Check that error messages explain reason for interface assignment failures.
# <翻译结束>


<原文开始>
// a few more - less exhaustive now
<原文结束>

# <翻译开始>
// a few more - less exhaustive now
# <翻译结束>


<原文开始>
// Check that constants representable as integers are in integer form
// before being used in operations that are only defined on integers.
<原文结束>

# <翻译开始>
// Check that constants representable as integers are in integer form
// before being used in operations that are only defined on integers.
# <翻译结束>


<原文开始>
// Check that in a n:1 variable declaration with type and initialization
// expression the type is distributed to all variables of the lhs before
// the initialization expression assignment is checked.
<原文结束>

# <翻译开始>
// Check that in a n:1 variable declaration with type and initialization
// expression the type is distributed to all variables of the lhs before
// the initialization expression assignment is checked.
# <翻译结束>


<原文开始>
// related: we should see an error since the result of f1 is ([]int, int)
<原文结束>

# <翻译开始>
// related: we should see an error since the result of f1 is ([]int, int)
# <翻译结束>


<原文开始>
// Test that we don't get "declared and not used"
// errors in the context of invalid/C objects.
<原文结束>

# <翻译开始>
// Test that we don't get "declared and not used"
// errors in the context of invalid/C objects.
# <翻译结束>


<原文开始>
	// these variables must be "used" even though
	// the LHS expressions/types below in which
	// context they are used are unknown/invalid
<原文结束>

# <翻译开始>
	// these variables must be "used" even though
	// the LHS expressions/types below in which
	// context they are used are unknown/invalid
# <翻译结束>


<原文开始>
// Test that we don't declare lhs variables in short variable
// declarations before we type-check function literals on the
// rhs.
<原文结束>

# <翻译开始>
// Test that we don't declare lhs variables in short variable
// declarations before we type-check function literals on the
// rhs.
# <翻译结束>


<原文开始>
// b and c must not be visible inside function literal
<原文结束>

# <翻译开始>
// b and c must not be visible inside function literal
# <翻译结束>


<原文开始>
// Test that we don't report a "missing return statement" error
// (due to incorrect context when type-checking interfaces).
<原文结束>

# <翻译开始>
// Test that we don't report a "missing return statement" error
// (due to incorrect context when type-checking interfaces).
# <翻译结束>


<原文开始>
// Test that we don't crash when the 'if' condition is missing.
<原文结束>

# <翻译开始>
// Test that we don't crash when the 'if' condition is missing.
# <翻译结束>


<原文开始>
// Test that we can embed alias type names in interfaces.
<原文结束>

# <翻译开始>
// Test that we can embed alias type names in interfaces.
# <翻译结束>


<原文开始>
// Test case from issue.
// cmd/compile reports a cycle as well.
<原文结束>

# <翻译开始>
// Test case from issue.
// cmd/compile reports a cycle as well.
# <翻译结束>


<原文开始>
// ERROR non-interface type struct\{\}
<原文结束>

# <翻译开始>
// ERROR non-interface type struct\{\}
# <翻译结束>


<原文开始>
// Test that method declarations don't introduce artificial cycles
// (issue #26124).
<原文结束>

# <翻译开始>
// Test that method declarations don't introduce artificial cycles
// (issue #26124).
# <翻译结束>


<原文开始>
// Reduced test case from issue #26124.
<原文结束>

# <翻译开始>
// Reduced test case from issue #26124.
# <翻译结束>


<原文开始>
// Test that we don't crash when type-checking composite literals
// containing errors in the type.
<原文结束>

# <翻译开始>
// Test that we don't crash when type-checking composite literals
// containing errors in the type.
# <翻译结束>


<原文开始>
// Test that invalid use of ... in parameter lists is recognized
// (issue #28281).
<原文结束>

# <翻译开始>
// Test that invalid use of ... in parameter lists is recognized
// (issue #28281).
# <翻译结束>


<原文开始>
// Issue #26234: Make various field/method lookup errors easier to read by matching cmd/compile's output
<原文结束>

# <翻译开始>
// Issue #26234: Make various field/method lookup errors easier to read by matching cmd/compile's output
# <翻译结束>


<原文开始>
	// The error message below should refer to the actual package name (syntax)
	// not the local package name (syn).
<原文结束>

# <翻译开始>
	// The error message below should refer to the actual package name (syntax)
	// not the local package name (syn).
# <翻译结束>


<原文开始>
// T is defined in this package, don't qualify its name with the package name.
<原文结束>

# <翻译开始>
// T is defined in this package, don't qualify its name with the package name.
# <翻译结束>


<原文开始>
// ERROR cannot use 0 \(untyped int constant\) as T
<原文结束>

# <翻译开始>
// ERROR cannot use 0 \(untyped int constant\) as T
# <翻译结束>


<原文开始>
// There is only one package with name syntax imported, only use the (global) package name in error messages.
<原文结束>

# <翻译开始>
// There is only one package with name syntax imported, only use the (global) package name in error messages.
# <翻译结束>


<原文开始>
// ERROR cannot use 0 \(untyped int constant\) as \*syntax.Prog
<原文结束>

# <翻译开始>
// ERROR cannot use 0 \(untyped int constant\) as \*syntax.Prog
# <翻译结束>


<原文开始>
	// Because both t1 and t2 have the same global package name (template),
	// qualify packages with full path name in this case.
<原文结束>

# <翻译开始>
	// Because both t1 and t2 have the same global package name (template),
	// qualify packages with full path name in this case.
# <翻译结束>

