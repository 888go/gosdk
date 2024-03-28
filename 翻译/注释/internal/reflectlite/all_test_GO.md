
<原文开始>
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
# <翻译结束>


<原文开始>
	// {struct {
	// 	x (interface {
	// 		a(func(func(int) int) func(func(int)) int)
	// 		b()
	// 	})
	// }{},
	// 	"interface { reflectlite_test.a(func(func(int) int) func(func(int)) int); reflectlite_test.b() }",
	// },
<原文结束>

# <翻译开始>
	// {struct {
	// 	x (interface {
	// 		a(func(func(int) int) func(func(int)) int)
	// 		b()
	// 	})
	// }{},
	// 	"interface { reflectlite_test.a(func(func(int) int) func(func(int)) int); reflectlite_test.b() }",
	// },
# <翻译结束>


<原文开始>
// assert(t, TypeString(v2.Type()), "interface {}")
<原文结束>

# <翻译开始>
// assert(t, TypeString(v2.Type()), "interface {}")
# <翻译结束>


<原文开始>
// Simple functions for DeepEqual tests.
<原文结束>

# <翻译开始>
// Simple functions for DeepEqual tests.
# <翻译结束>


<原文开始>
// Nil vs empty: not the same.
<原文结束>

# <翻译开始>
// Nil vs empty: not the same.
# <翻译结束>


<原文开始>
	// These implement IsNil.
	// Wrap in extra struct to hide interface type.
<原文结束>

# <翻译开始>
	// These implement IsNil.
	// Wrap in extra struct to hide interface type.
# <翻译结束>


<原文开始>
// panics if not okay to call
<原文结束>

# <翻译开始>
// panics if not okay to call
# <翻译结束>


<原文开始>
// Check the implementations
<原文结束>

# <翻译开始>
// Check the implementations
# <翻译结束>


<原文开始>
// Indirect returns the value that v points to.
// If v is a nil pointer, Indirect returns a zero Value.
// If v is not a pointer, Indirect returns v.
<原文结束>

# <翻译开始>
// Indirect returns the value that v points to.
// If v is a nil pointer, Indirect returns a zero Value.
// If v is not a pointer, Indirect returns v.
# <翻译结束>


<原文开始>
//println("Point.Dist", p.x, p.y, scale)
<原文结束>

# <翻译开始>
//println("Point.Dist", p.x, p.y, scale)
# <翻译结束>


<原文开始>
// Exercise no-argument/no-result paths.
<原文结束>

# <翻译开始>
// Exercise no-argument/no-result paths.
# <翻译结束>


<原文开始>
		// We can uncomment this when compiler escape analysis
		// is good enough to see that the integer assigned to i
		// does not escape and therefore need not be allocated.
		//
		// i = 42 + j
		// v = ValueOf(i)
		// if int(v.Int()) != 42+j {
		// 	panic("wrong int")
		// }
<原文结束>

# <翻译开始>
		// We can uncomment this when compiler escape analysis
		// is good enough to see that the integer assigned to i
		// does not escape and therefore need not be allocated.
		//
		// i = 42 + j
		// v = ValueOf(i)
		// if int(v.Int()) != 42+j {
		// 	panic("wrong int")
		// }
# <翻译结束>


<原文开始>
// Used to have inconsistency between IsValid() and Kind() != Invalid.
<原文结束>

# <翻译开始>
// Used to have inconsistency between IsValid() and Kind() != Invalid.
# <翻译结束>


<原文开始>
// TestUnaddressableField tests that the reflect package will not allow
// a type from another package to be used as a named type with an
// unexported field.
//
// This ensures that unexported fields cannot be modified by other packages.
<原文结束>

# <翻译开始>
// TestUnaddressableField tests that the reflect package will not allow
// a type from another package to be used as a named type with an
// unexported field.
//
// This ensures that unexported fields cannot be modified by other packages.
# <翻译结束>


<原文开始>
// type defined in reflect, a different package
<原文结束>

# <翻译开始>
// type defined in reflect, a different package
# <翻译结束>

