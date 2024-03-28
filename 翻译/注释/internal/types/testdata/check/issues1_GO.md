
<原文开始>
// -oldComparableSemantics
<原文结束>

# <翻译开始>
// -oldComparableSemantics
# <翻译结束>


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
// This file contains regression tests for bugs found.
<原文结束>

# <翻译开始>
// This file contains regression tests for bugs found.
# <翻译结束>


<原文开始>
// interfaces of different types
<原文结束>

# <翻译开始>
// interfaces of different types
# <翻译结束>


<原文开始>
// If we have a receiver of pointer to type parameter type (below: *T)
// we don't have any methods, like for interfaces.
<原文结束>

# <翻译开始>
// If we have a receiver of pointer to type parameter type (below: *T)
// we don't have any methods, like for interfaces.
# <翻译结束>


<原文开始>
// using an interface literal as bound
<原文结束>

# <翻译开始>
// using an interface literal as bound
# <翻译结束>


<原文开始>
// When a type parameter is used as an argument to instantiate a parameterized
// type, the type argument's type set must be a subset of the instantiated type
// parameter's type set.
<原文结束>

# <翻译开始>
// When a type parameter is used as an argument to instantiate a parameterized
// type, the type argument's type set must be a subset of the instantiated type
// parameter's type set.
# <翻译结束>


<原文开始>
// This is the original (simplified) program causing the same issue.
<原文结束>

# <翻译开始>
// This is the original (simplified) program causing the same issue.
# <翻译结束>


<原文开始>
// When we encounter an instantiated type such as Elem[T] we must
// not "expand" the instantiation when the type to be instantiated
// (Elem in this case) is not yet fully set up.
<原文结束>

# <翻译开始>
// When we encounter an instantiated type such as Elem[T] we must
// not "expand" the instantiation when the type to be instantiated
// (Elem in this case) is not yet fully set up.
# <翻译结束>


<原文开始>
// This is the original program causing the same issue.
<原文结束>

# <翻译开始>
// This is the original program causing the same issue.
# <翻译结束>


<原文开始>
// Self-recursive instantiations must work correctly.
<原文结束>

# <翻译开始>
// Self-recursive instantiations must work correctly.
# <翻译结束>


<原文开始>
// And a variation that also caused a problem with an
// unresolved underlying type.
<原文结束>

# <翻译开始>
// And a variation that also caused a problem with an
// unresolved underlying type.
# <翻译结束>


<原文开始>
// Infinite generic type declarations must lead to an error.
<原文结束>

# <翻译开始>
// Infinite generic type declarations must lead to an error.
# <翻译结束>


<原文开始>
// The implementation of conversions T(x) between integers and floating-point
// numbers checks that both T and x have either integer or floating-point
// type. When the type of T or x is a type parameter, the respective simple
// predicate disjunction in the implementation was wrong because if a type set
// contains both an integer and a floating-point type, the type parameter is
// neither an integer or a floating-point number.
<原文结束>

# <翻译开始>
// The implementation of conversions T(x) between integers and floating-point
// numbers checks that both T and x have either integer or floating-point
// type. When the type of T or x is a type parameter, the respective simple
// predicate disjunction in the implementation was wrong because if a type set
// contains both an integer and a floating-point type, the type parameter is
// neither an integer or a floating-point number.
# <翻译结束>


<原文开始>
// When testing binary operators, for +, the operand types must either be
// both numeric, or both strings. The implementation had the same problem
// with this check as the conversion issue above (issue #39623).
<原文结束>

# <翻译开始>
// When testing binary operators, for +, the operand types must either be
// both numeric, or both strings. The implementation had the same problem
// with this check as the conversion issue above (issue #39623).
# <翻译结束>


<原文开始>
// Simplified, from https://go2goplay.golang.org/p/efS6x6s-9NI:
<原文结束>

# <翻译开始>
// Simplified, from https://go2goplay.golang.org/p/efS6x6s-9NI:
# <翻译结束>


<原文开始>
// Assignability of an unnamed pointer type to a type parameter that
// has a matching underlying type.
<原文结束>

# <翻译开始>
// Assignability of an unnamed pointer type to a type parameter that
// has a matching underlying type.
# <翻译结束>


<原文开始>
// Indexing of type parameters containing type parameters in their constraint terms:
<原文结束>

# <翻译开始>
// Indexing of type parameters containing type parameters in their constraint terms:
# <翻译结束>


<原文开始>
// Conversion of a local type to a type parameter.
<原文结束>

# <翻译开始>
// Conversion of a local type to a type parameter.
# <翻译结束>


<原文开始>
// Indexing a type parameter with an array type bound checks length.
// (Example by mdempsky@.)
<原文结束>

# <翻译开始>
// Indexing a type parameter with an array type bound checks length.
// (Example by mdempsky@.)
# <翻译结束>


<原文开始>
// Pointer indirection of a type parameter.
<原文结束>

# <翻译开始>
// Pointer indirection of a type parameter.
# <翻译结束>


<原文开始>
// Channel sends and receives on type parameters.
<原文结束>

# <翻译开始>
// Channel sends and receives on type parameters.
# <翻译结束>


<原文开始>
// Calling of a generic variable.
<原文结束>

# <翻译开始>
// Calling of a generic variable.
# <翻译结束>


<原文开始>
// We must compare against the (possibly underlying) types of term list
// elements when checking if a constraint is satisfied by a type.
// The underlying type of each term must be computed after the
// interface has been instantiated as its constraint may contain
// a type parameter that was substituted with a defined type.
// Test case from an (originally) failing example.
<原文结束>

# <翻译开始>
// We must compare against the (possibly underlying) types of term list
// elements when checking if a constraint is satisfied by a type.
// The underlying type of each term must be computed after the
// interface has been instantiated as its constraint may contain
// a type parameter that was substituted with a defined type.
// Test case from an (originally) failing example.
# <翻译结束>


<原文开始>
// A generic function must be instantiated with a type, not a value.
<原文结束>

# <翻译开始>
// A generic function must be instantiated with a type, not a value.
# <翻译结束>

