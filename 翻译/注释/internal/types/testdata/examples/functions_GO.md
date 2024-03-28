
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
// This file shows some examples of type-parameterized functions.
<原文结束>

# <翻译开始>
// This file shows some examples of type-parameterized functions.
# <翻译结束>


<原文开始>
// Reverse is a generic function that takes a []T argument and
// reverses that slice in place.
<原文结束>

# <翻译开始>
// Reverse is a generic function that takes a []T argument and
// reverses that slice in place.
# <翻译结束>


<原文开始>
// Reverse can be called with an explicit type argument.
<原文结束>

# <翻译开始>
// Reverse can be called with an explicit type argument.
# <翻译结束>


<原文开始>
	// Since the type parameter is used for an incoming argument,
	// it can be inferred from the provided argument's type.
<原文结束>

# <翻译开始>
	// Since the type parameter is used for an incoming argument,
	// it can be inferred from the provided argument's type.
# <翻译结束>


<原文开始>
	// But the incoming argument must have a type, even if it's a
	// default type. An untyped nil won't work.
	// Reverse(nil) // this won't type-check
<原文结束>

# <翻译开始>
	// But the incoming argument must have a type, even if it's a
	// default type. An untyped nil won't work.
	// Reverse(nil) // this won't type-check
# <翻译结束>


<原文开始>
// A typed nil will work, though.
<原文结束>

# <翻译开始>
// A typed nil will work, though.
# <翻译结束>


<原文开始>
// Certain functions, such as the built-in `new` could be written using
// type parameters.
<原文结束>

# <翻译开始>
// Certain functions, such as the built-in `new` could be written using
// type parameters.
# <翻译结束>


<原文开始>
// When calling our own `new`, we need to pass the type parameter
// explicitly since there is no (value) argument from which the
// result type could be inferred. We don't try to infer the
// result type from the assignment to keep things simple and
// easy to understand.
<原文结束>

# <翻译开始>
// When calling our own `new`, we need to pass the type parameter
// explicitly since there is no (value) argument from which the
// result type could be inferred. We don't try to infer the
// result type from the assignment to keep things simple and
// easy to understand.
# <翻译结束>


<原文开始>
// the result type is indeed *float64
<原文结束>

# <翻译开始>
// the result type is indeed *float64
# <翻译结束>


<原文开始>
// A function may have multiple type parameters, of course.
<原文结束>

# <翻译开始>
// A function may have multiple type parameters, of course.
# <翻译结束>


<原文开始>
// As before, we can pass type parameters explicitly.
<原文结束>

# <翻译开始>
// As before, we can pass type parameters explicitly.
# <翻译结束>


<原文开始>
// Or we can use type inference.
<原文结束>

# <翻译开始>
// Or we can use type inference.
# <翻译结束>


<原文开始>
// Type inference works in a straight-forward manner even
// for variadic functions.
<原文结束>

# <翻译开始>
// Type inference works in a straight-forward manner even
// for variadic functions.
# <翻译结束>


<原文开始>
// var _ = variadic(1)// ERROR not enough arguments
<原文结束>

# <翻译开始>
// var _ = variadic(1)// ERROR not enough arguments
# <翻译结束>


<原文开始>
// Type inference also works in recursive function calls where
// the inferred type is the type parameter of the caller.
<原文结束>

# <翻译开始>
// Type inference also works in recursive function calls where
// the inferred type is the type parameter of the caller.
# <翻译结束>


<原文开始>
// Here's an example of a recursive function call with variadic
// arguments and type inference inferring the type parameter of
// the caller (i.e., itself).
<原文结束>

# <翻译开始>
// Here's an example of a recursive function call with variadic
// arguments and type inference inferring the type parameter of
// the caller (i.e., itself).
# <翻译结束>


<原文开始>
// When inferring channel types, the channel direction is ignored
// for the purpose of type inference. Once the type has been in-
// fered, the usual parameter passing rules are applied.
// Thus even if a type can be inferred successfully, the function
// call may not be valid.
<原文结束>

# <翻译开始>
// When inferring channel types, the channel direction is ignored
// for the purpose of type inference. Once the type has been in-
// fered, the usual parameter passing rules are applied.
// Thus even if a type can be inferred successfully, the function
// call may not be valid.
# <翻译结束>


<原文开始>
// When inferring elements of unnamed composite parameter types,
// if the arguments are defined types, use their underlying types.
// Even though the matching types are not exactly structurally the
// same (one is a type literal, the other a named type), because
// assignment is permitted, parameter passing is permitted as well,
// so type inference should be able to handle these cases well.
<原文结束>

# <翻译开始>
// When inferring elements of unnamed composite parameter types,
// if the arguments are defined types, use their underlying types.
// Even though the matching types are not exactly structurally the
// same (one is a type literal, the other a named type), because
// assignment is permitted, parameter passing is permitted as well,
// so type inference should be able to handle these cases well.
# <翻译结束>


<原文开始>
// Here's a realistic example.
<原文结束>

# <翻译开始>
// Here's a realistic example.
# <翻译结束>


<原文开始>
// Generic type declarations cannot have empty type parameter lists
// (that would indicate a slice type). Thus, generic functions cannot
// have empty type parameter lists, either. This is a syntax error.
<原文结束>

# <翻译开始>
// Generic type declarations cannot have empty type parameter lists
// (that would indicate a slice type). Thus, generic functions cannot
// have empty type parameter lists, either. This is a syntax error.
# <翻译结束>


<原文开始>
// Generic functions must have a function body.
<原文结束>

# <翻译开始>
// Generic functions must have a function body.
# <翻译结束>

