
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
// This file shows some examples of generic types.
<原文结束>

# <翻译开始>
// This file shows some examples of generic types.
# <翻译结束>


<原文开始>
// List is just what it says - a slice of E elements.
<原文结束>

# <翻译开始>
// List is just what it says - a slice of E elements.
# <翻译结束>


<原文开始>
// A generic (parameterized) type must always be instantiated
// before it can be used to designate the type of a variable
// (including a struct field, or function parameter); though
// for the latter cases, the provided type may be another type
// parameter. So:
<原文结束>

# <翻译开始>
// A generic (parameterized) type must always be instantiated
// before it can be used to designate the type of a variable
// (including a struct field, or function parameter); though
// for the latter cases, the provided type may be another type
// parameter. So:
# <翻译结束>


<原文开始>
// A generic binary tree might be declared as follows.
<原文结束>

# <翻译开始>
// A generic binary tree might be declared as follows.
# <翻译结束>


<原文开始>
// A simple instantiation of Tree:
<原文结束>

# <翻译开始>
// A simple instantiation of Tree:
# <翻译结束>


<原文开始>
// The actual type parameter provided may be a generic type itself:
<原文结束>

# <翻译开始>
// The actual type parameter provided may be a generic type itself:
# <翻译结束>


<原文开始>
// A couple of more complex examples.
// We don't need extra parentheses around the element type of the slices on
// the right (unlike when we use ()'s rather than []'s for type parameters).
<原文结束>

# <翻译开始>
// A couple of more complex examples.
// We don't need extra parentheses around the element type of the slices on
// the right (unlike when we use ()'s rather than []'s for type parameters).
# <翻译结束>


<原文开始>
// Type parameters act like type aliases when used in generic types
// in the sense that we can "emulate" a specific type instantiation
// with type aliases.
<原文结束>

# <翻译开始>
// Type parameters act like type aliases when used in generic types
// in the sense that we can "emulate" a specific type instantiation
// with type aliases.
# <翻译结束>


<原文开始>
	// This assignment is invalid because the types of x1, x2 are T1(...)
	// and T2(...) respectively, which are two different defined types.
<原文结束>

# <翻译开始>
	// This assignment is invalid because the types of x1, x2 are T1(...)
	// and T2(...) respectively, which are two different defined types.
# <翻译结束>


<原文开始>
	// This assignment is valid because the types of x1.f and x2.f are
	// both struct { g int }; the type parameters act like type aliases
	// and their actual names don't come into play here.
<原文结束>

# <翻译开始>
	// This assignment is valid because the types of x1.f and x2.f are
	// both struct { g int }; the type parameters act like type aliases
	// and their actual names don't come into play here.
# <翻译结束>


<原文开始>
// We can verify this behavior using type aliases instead:
<原文结束>

# <翻译开始>
// We can verify this behavior using type aliases instead:
# <翻译结束>


<原文开始>
// Another interesting corner case are generic types that don't use
// their type arguments. For instance:
<原文结束>

# <翻译开始>
// Another interesting corner case are generic types that don't use
// their type arguments. For instance:
# <翻译结束>


<原文开始>
// Are these two variables of the same type? After all, their underlying
// types are identical. We consider them to be different because each type
// instantiation creates a new named type, in this case T<int> and T<bool>
// even if their underlying types are identical. This is sensible because
// we might still have methods that have different signatures or behave
// differently depending on the type arguments, and thus we can't possibly
// consider such types identical. Consequently:
<原文结束>

# <翻译开始>
// Are these two variables of the same type? After all, their underlying
// types are identical. We consider them to be different because each type
// instantiation creates a new named type, in this case T<int> and T<bool>
// even if their underlying types are identical. This is sensible because
// we might still have methods that have different signatures or behave
// differently depending on the type arguments, and thus we can't possibly
// consider such types identical. Consequently:
# <翻译结束>


<原文开始>
// Generic types cannot be used without instantiation.
<原文结束>

# <翻译开始>
// Generic types cannot be used without instantiation.
# <翻译结束>


<原文开始>
// ERROR cannot use generic type T
<原文结束>

# <翻译开始>
// ERROR cannot use generic type T
# <翻译结束>


<原文开始>
// In type context, generic (parameterized) types cannot be parenthesized before
// being instantiated. See also NOTES entry from 12/4/2019.
<原文结束>

# <翻译开始>
// In type context, generic (parameterized) types cannot be parenthesized before
// being instantiated. See also NOTES entry from 12/4/2019.
# <翻译结束>


<原文开始>
// All types may be parameterized, including interfaces.
<原文结束>

# <翻译开始>
// All types may be parameterized, including interfaces.
# <翻译结束>


<原文开始>
// There is no such thing as a variadic generic type.
<原文结束>

# <翻译开始>
// There is no such thing as a variadic generic type.
# <翻译结束>


<原文开始>
// Generic interfaces may be embedded as one would expect.
<原文结束>

# <翻译开始>
// Generic interfaces may be embedded as one would expect.
# <翻译结束>


<原文开始>
// Issue #45639: We don't allow this anymore. Keep this code
//               in case we decide to revisit this decision.
//
// It's possible to declare local types whose underlying types
// are type parameters. As with ordinary type definitions, the
// types underlying properties are "inherited" but the methods
// are not.
// func _[T interface{ m(); ~int }]() {
// 	type L T
// 	var x L
// 
// 	// m is not defined on L (it is not "inherited" from
// 	// its underlying type).
// 	x.m /* ERROR x.m undefined */ ()
// 
// 	// But the properties of T, such that as that it supports
// 	// the operations of the types given by its type bound,
// 	// are also the properties of L.
// 	x++
// 	_ = x - x
// 
// 	// On the other hand, if we define a local alias for T,
// 	// that alias stands for T as expected.
// 	type A = T
// 	var y A
// 	y.m()
// 	_ = y < 0
// }
<原文结束>

# <翻译开始>
// Issue #45639: We don't allow this anymore. Keep this code
//               in case we decide to revisit this decision.
//
// It's possible to declare local types whose underlying types
// are type parameters. As with ordinary type definitions, the
// types underlying properties are "inherited" but the methods
// are not.
// func _[T interface{ m(); ~int }]() {
// 	type L T
// 	var x L
// 
// 	// m is not defined on L (it is not "inherited" from
// 	// its underlying type).
// 	x.m /* ERROR x.m undefined */ ()
// 
// 	// But the properties of T, such that as that it supports
// 	// the operations of the types given by its type bound,
// 	// are also the properties of L.
// 	x++
// 	_ = x - x
// 
// 	// On the other hand, if we define a local alias for T,
// 	// that alias stands for T as expected.
// 	type A = T
// 	var y A
// 	y.m()
// 	_ = y < 0
// }
# <翻译结束>


<原文开始>
// For now, a lone type parameter is not permitted as RHS in a type declaration (issue #45639).
// // It is not permitted to declare a local type whose underlying
// // type is a type parameter not declared by that type declaration.
// func _[T any]() {
// 	type _ T         // ERROR cannot use function type parameter T as RHS in type declaration
// 	type _ [_ any] T // ERROR cannot use function type parameter T as RHS in type declaration
// }
<原文结束>

# <翻译开始>
// For now, a lone type parameter is not permitted as RHS in a type declaration (issue #45639).
// // It is not permitted to declare a local type whose underlying
// // type is a type parameter not declared by that type declaration.
// func _[T any]() {
// 	type _ T         // ERROR cannot use function type parameter T as RHS in type declaration
// 	type _ [_ any] T // ERROR cannot use function type parameter T as RHS in type declaration
// }
# <翻译结束>


<原文开始>
// As a special case, an explicit type argument may be omitted
// from a type parameter bound if the type bound expects exactly
// one type argument. In that case, the type argument is the
// respective type parameter to which the type bound applies.
// Note: We may not permit this syntactic sugar at first.
// Note: This is now disabled. All examples below are adjusted.
<原文结束>

# <翻译开始>
// As a special case, an explicit type argument may be omitted
// from a type parameter bound if the type bound expects exactly
// one type argument. In that case, the type argument is the
// respective type parameter to which the type bound applies.
// Note: We may not permit this syntactic sugar at first.
// Note: This is now disabled. All examples below are adjusted.
# <翻译结束>


<原文开始>
// We don't need to explicitly instantiate the Adder bound
// if we have exactly one type parameter.
<原文结束>

# <翻译开始>
// We don't need to explicitly instantiate the Adder bound
// if we have exactly one type parameter.
# <翻译结束>


<原文开始>
// Valid and invalid variations.
<原文结束>

# <翻译开始>
// Valid and invalid variations.
# <翻译结束>


<原文开始>
// When the type argument is left away, the type bound is
// instantiated for each type parameter with that type
// parameter.
// Note: We may not permit this syntactic sugar at first.
<原文结束>

# <翻译开始>
// When the type argument is left away, the type bound is
// instantiated for each type parameter with that type
// parameter.
// Note: We may not permit this syntactic sugar at first.
# <翻译结束>


<原文开始>
// The type of variables (incl. parameters and return values) cannot
// be an interface with type constraints or be/embed comparable.
<原文结束>

# <翻译开始>
// The type of variables (incl. parameters and return values) cannot
// be an interface with type constraints or be/embed comparable.
# <翻译结束>


<原文开始>
// Type parameters are never const types, i.e., it's
// not possible to declare a constant of type parameter type.
// (If a type set contains just a single const type, we could
// allow it, but such type sets don't make much sense in the
// first place.)
<原文结束>

# <翻译开始>
// Type parameters are never const types, i.e., it's
// not possible to declare a constant of type parameter type.
// (If a type set contains just a single const type, we could
// allow it, but such type sets don't make much sense in the
// first place.)
# <翻译结束>


<原文开始>
// It is possible to create composite literals of type parameter
// type as long as it's possible to create a composite literal
// of the core type of the type parameter's constraint.
<原文结束>

# <翻译开始>
// It is possible to create composite literals of type parameter
// type as long as it's possible to create a composite literal
// of the core type of the type parameter's constraint.
# <翻译结束>


<原文开始>
// This is a degenerate case with a singleton type set, but we can create
// composite literals even if the core type is a defined type.
<原文结束>

# <翻译开始>
// This is a degenerate case with a singleton type set, but we can create
// composite literals even if the core type is a defined type.
# <翻译结束>

