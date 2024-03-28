
<原文开始>
// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
# <翻译结束>


<原文开始>
// Parameterized types with methods
<原文结束>

# <翻译开始>
// Parameterized types with methods
# <翻译结束>


<原文开始>
// A test case for instantiating types with other types (extracted from map.go2)
<原文结束>

# <翻译开始>
// A test case for instantiating types with other types (extracted from map.go2)
# <翻译结束>


<原文开始>
// A more complex test case testing type bounds (extracted from linalg.go2 and reduced to essence)
<原文结束>

# <翻译开始>
// A more complex test case testing type bounds (extracted from linalg.go2 and reduced to essence)
# <翻译结束>


<原文开始>
// For now, a lone type parameter is not permitted as RHS in a type declaration (issue #45639).
// type OrderedAbs[T any] T
// 
// func (a OrderedAbs[T]) Abs() OrderedAbs[T]
// 
// func OrderedAbsDifference[T any](x T) {
// 	AbsDifference(OrderedAbs[T](x))
// }
<原文结束>

# <翻译开始>
// For now, a lone type parameter is not permitted as RHS in a type declaration (issue #45639).
// type OrderedAbs[T any] T
// 
// func (a OrderedAbs[T]) Abs() OrderedAbs[T]
// 
// func OrderedAbsDifference[T any](x T) {
// 	AbsDifference(OrderedAbs[T](x))
// }
# <翻译结束>


<原文开始>
// same code, reduced to essence
<原文结束>

# <翻译开始>
// same code, reduced to essence
# <翻译结束>


<原文开始>
// For now, a lone type parameter is not permitted as RHS in a type declaration (issue #45639).
// type T4[P any] P
// 
// func (_ T4[P]) m() T4[P]
// 
// func _[Q any](x Q) {
// 	g(T4[Q](x))
// }
<原文结束>

# <翻译开始>
// For now, a lone type parameter is not permitted as RHS in a type declaration (issue #45639).
// type T4[P any] P
// 
// func (_ T4[P]) m() T4[P]
// 
// func _[Q any](x Q) {
// 	g(T4[Q](x))
// }
# <翻译结束>


<原文开始>
// Another test case that caused  problems in the past
<原文结束>

# <翻译开始>
// Another test case that caused  problems in the past
# <翻译结束>


<原文开始>
// Invoking methods with parameterized receiver types uses
// type inference to determine the actual type arguments matching
// the receiver type parameters from the actual receiver argument.
// Go does implicit address-taking and dereferenciation depending
// on the actual receiver and the method's receiver type. To make
// type inference work, the type-checker matches "pointer-ness"
// of the actual receiver and the method's receiver type.
// The following code tests this mechanism.
<原文结束>

# <翻译开始>
// Invoking methods with parameterized receiver types uses
// type inference to determine the actual type arguments matching
// the receiver type parameters from the actual receiver argument.
// Go does implicit address-taking and dereferenciation depending
// on the actual receiver and the method's receiver type. To make
// type inference work, the type-checker matches "pointer-ness"
// of the actual receiver and the method's receiver type.
// The following code tests this mechanism.
# <翻译结束>


<原文开始>
// It is ok to have multiple embedded unions.
<原文结束>

# <翻译开始>
// It is ok to have multiple embedded unions.
# <翻译结束>


<原文开始>
// Type sets may contain each type at most once.
<原文结束>

# <翻译开始>
// Type sets may contain each type at most once.
# <翻译结束>


<原文开始>
// Interface term lists can contain any type, incl. *Named types.
// Verify that we use the underlying type(s) of the type(s) in the
// term list when determining if an operation is permitted.
<原文结束>

# <翻译开始>
// Interface term lists can contain any type, incl. *Named types.
// Verify that we use the underlying type(s) of the type(s) in the
// term list when determining if an operation is permitted.
# <翻译结束>


<原文开始>
// Embedding of interfaces with term lists leads to interfaces
// with term lists that are the intersection of the embedded
// term lists.
<原文结束>

# <翻译开始>
// Embedding of interfaces with term lists leads to interfaces
// with term lists that are the intersection of the embedded
// term lists.
# <翻译结束>


<原文开始>
// Using a function instance as a type is an error.
<原文结束>

# <翻译开始>
// Using a function instance as a type is an error.
# <翻译结束>


<原文开始>
// Empty type sets can only be satisfied by empty type sets.
<原文结束>

# <翻译开始>
// Empty type sets can only be satisfied by empty type sets.
# <翻译结束>


<原文开始>
// force an empty type set
<原文结束>

# <翻译开始>
// force an empty type set
# <翻译结束>

