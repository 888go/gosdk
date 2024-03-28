
<原文开始>
// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
# <翻译结束>


<原文开始>
// ERROR not enough arguments
<原文结束>

# <翻译开始>
// ERROR not enough arguments
# <翻译结束>


<原文开始>
// append a single element     s1 == []int{0, 0, 2}
<原文结束>

# <翻译开始>
// append a single element     s1 == []int{0, 0, 2}
# <翻译结束>


<原文开始>
// append multiple elements    s2 == []int{0, 0, 2, 3, 5, 7}
<原文结束>

# <翻译开始>
// append multiple elements    s2 == []int{0, 0, 2, 3, 5, 7}
# <翻译结束>


<原文开始>
// append a slice              s3 == []int{0, 0, 2, 3, 5, 7, 0, 0}
<原文结束>

# <翻译开始>
// append a slice              s3 == []int{0, 0, 2, 3, 5, 7, 0, 0}
# <翻译结束>


<原文开始>
// append overlapping slice    s4 == []int{3, 5, 7, 2, 3, 5, 7, 0, 0}
<原文结束>

# <翻译开始>
// append overlapping slice    s4 == []int{3, 5, 7, 2, 3, 5, 7, 0, 0}
# <翻译结束>


<原文开始>
//                             t == []interface{}{42, 3.1415, "foo"}
<原文结束>

# <翻译开始>
//                             t == []interface{}{42, 3.1415, "foo"}
# <翻译结束>


<原文开始>
// append string contents      b == []byte{'b', 'a', 'r' }
<原文结束>

# <翻译开始>
// append string contents      b == []byte{'b', 'a', 'r' }
# <翻译结束>


<原文开始>
// TODO(gri) better error message
<原文结束>

# <翻译开始>
// TODO(gri) better error message
# <翻译结束>


<原文开始>
// ERROR too many arguments
<原文结束>

# <翻译开始>
// ERROR too many arguments
# <翻译结束>


<原文开始>
// test cases for issue 7387
<原文结束>

# <翻译开始>
// test cases for issue 7387
# <翻译结束>


<原文开始>
// floating-point argument types must be identical
<原文结束>

# <翻译开始>
// floating-point argument types must be identical
# <翻译结束>


<原文开始>
// n1 == 6, s == []int{0, 1, 2, 3, 4, 5}
<原文结束>

# <翻译开始>
// n1 == 6, s == []int{0, 1, 2, 3, 4, 5}
# <翻译结束>


<原文开始>
// n2 == 4, s == []int{2, 3, 4, 5, 4, 5}
<原文结束>

# <翻译开始>
// n2 == 4, s == []int{2, 3, 4, 5, 4, 5}
# <翻译结束>


<原文开始>
// n3 == 5, b == []byte("Hello")
<原文结束>

# <翻译开始>
// n3 == 5, b == []byte("Hello")
# <翻译结束>


<原文开始>
// complex type may not be predeclared
<原文结束>

# <翻译开始>
// complex type may not be predeclared
# <翻译结束>


<原文开始>
// if argument is untyped, result is untyped
<原文结束>

# <翻译开始>
// if argument is untyped, result is untyped
# <翻译结束>


<原文开始>
// lhs constant shift operands are typed as complex128
<原文结束>

# <翻译开始>
// lhs constant shift operands are typed as complex128
# <翻译结束>


<原文开始>
// ok because n has unknown value and no error is reported
<原文结束>

# <翻译开始>
// ok because n has unknown value and no error is reported
# <翻译结束>


<原文开始>
// assuming types.DefaultPtrSize == 8
<原文结束>

# <翻译开始>
// assuming types.DefaultPtrSize == 8
# <翻译结束>


<原文开始>
// basic types have size guarantees
<原文结束>

# <翻译开始>
// basic types have size guarantees
# <翻译结束>


<原文开始>
// test case for issue 5670
<原文结束>

# <翻译开始>
// test case for issue 5670
# <翻译结束>


<原文开始>
// here we allow nil as ptr argument (in contrast to unsafe.Slice)
<原文结束>

# <翻译开始>
// here we allow nil as ptr argument (in contrast to unsafe.Slice)
# <翻译结束>


<原文开始>
	// Uncomment the code below to test trace - will produce console output
	// _ = trace /* ERROR no value */ ()
	// _ = trace(1)
	// _ = trace(true, 1.2, '\'', "foo", 42i, "foo" <= "bar")
<原文结束>

# <翻译开始>
	// Uncomment the code below to test trace - will produce console output
	// _ = trace /* ERROR no value */ ()
	// _ = trace(1)
	// _ = trace(true, 1.2, '\'', "foo", 42i, "foo" <= "bar")
# <翻译结束>


<原文开始>
	// Uncomment the code below to test trace - will produce console output
	// trace(f0())
	// trace(f1())
	// trace(f2())
	// trace(f3())
	// trace(f0(), 1)
	// trace(f1(), 1, 2)
	// trace(f2(), 1, 2, 3)
	// trace(f3(), 1, 2, 3, 4)
<原文结束>

# <翻译开始>
	// Uncomment the code below to test trace - will produce console output
	// trace(f0())
	// trace(f1())
	// trace(f2())
	// trace(f3())
	// trace(f0(), 1)
	// trace(f1(), 1, 2)
	// trace(f2(), 1, 2, 3)
	// trace(f3(), 1, 2, 3, 4)
# <翻译结束>

