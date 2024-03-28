
<原文开始>
// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
# <翻译结束>


<原文开始>
// import "io"// for type assertion tests
<原文结束>

# <翻译开始>
// import "io"// for type assertion tests
# <翻译结束>


<原文开始>
// Constraints (incl. any) may be parenthesized.
<原文结束>

# <翻译开始>
// Constraints (incl. any) may be parenthesized.
# <翻译结束>


<原文开始>
// indexing with various combinations of map types in type sets (see issue #42616)
<原文结束>

# <翻译开始>
// indexing with various combinations of map types in type sets (see issue #42616)
# <翻译结束>


<原文开始>
// different map element types
<原文结束>

# <翻译开始>
// different map element types
# <翻译结束>


<原文开始>
// indexing with various combinations of array and other types in type sets
<原文结束>

# <翻译开始>
// indexing with various combinations of array and other types in type sets
# <翻译结束>


<原文开始>
// indexing with strings and non-variable arrays (assignment not permitted)
<原文结束>

# <翻译开始>
// indexing with strings and non-variable arrays (assignment not permitted)
# <翻译结束>


<原文开始>
// type inference with variadic functions
<原文结束>

# <翻译开始>
// type inference with variadic functions
# <翻译结束>


<原文开始>
// init functions cannot have type parameters
<原文结束>

# <翻译开始>
// init functions cannot have type parameters
# <翻译结束>


<原文开始>
// type inference across parameterized types
<原文结束>

# <翻译开始>
// type inference across parameterized types
# <翻译结束>


<原文开始>
// corner case for type inference
// (was bug: after instanting f11, the type-checker didn't mark f11 as non-generic)
<原文结束>

# <翻译开始>
// corner case for type inference
// (was bug: after instanting f11, the type-checker didn't mark f11 as non-generic)
# <翻译结束>


<原文开始>
// the previous example was extracted from
<原文结束>

# <翻译开始>
// the previous example was extracted from
# <翻译结束>


<原文开始>
// For now, a lone type parameter is not permitted as RHS in a type declaration (issue #45639).
// func f12[T interface{m() T}]() {}
// 
// type A[T any] T
// 
// func (a A[T]) m() A[T]
// 
// func _[T any]() {
// 	f12[A[T]]()
// }
<原文结束>

# <翻译开始>
// For now, a lone type parameter is not permitted as RHS in a type declaration (issue #45639).
// func f12[T interface{m() T}]() {}
// 
// type A[T any] T
// 
// func (a A[T]) m() A[T]
// 
// func _[T any]() {
// 	f12[A[T]]()
// }
# <翻译结束>


<原文开始>
// type assertions and type switches over generic types
// NOTE: These are currently disabled because it's unclear what the correct
// approach is, and one can always work around by assigning the variable to
// an interface first.
<原文结束>

# <翻译开始>
// type assertions and type switches over generic types
// NOTE: These are currently disabled because it's unclear what the correct
// approach is, and one can always work around by assigning the variable to
// an interface first.
# <翻译结束>


<原文开始>
// // ReadByte1 corresponds to the ReadByte example in the draft design.
// func ReadByte1[T io.Reader](r T) (byte, error) {
// 	if br, ok := r.(io.ByteReader); ok {
// 		return br.ReadByte()
// 	}
// 	var b [1]byte
// 	_, err := r.Read(b[:])
// 	return b[0], err
// }
//
// // ReadBytes2 is like ReadByte1 but uses a type switch instead.
// func ReadByte2[T io.Reader](r T) (byte, error) {
//         switch br := r.(type) {
//         case io.ByteReader:
//                 return br.ReadByte()
//         }
// 	var b [1]byte
// 	_, err := r.Read(b[:])
// 	return b[0], err
// }
//
// // type assertions and type switches over generic types are strict
// type I3 interface {
//         m(int)
// }
//
// type I4 interface {
//         m() int // different signature from I3.m
// }
//
// func _[T I3](x I3, p T) {
//         // type assertions and type switches over interfaces are not strict
//         _ = x.(I4)
//         switch x.(type) {
//         case I4:
//         }
//
//         // type assertions and type switches over generic types are strict
//         _ = p /* ERROR cannot have dynamic type I4 */.(I4)
//         switch p.(type) {
//         case I4 /* ERROR cannot have dynamic type I4 */ :
//         }
// }
<原文结束>

# <翻译开始>
// // ReadByte1 corresponds to the ReadByte example in the draft design.
// func ReadByte1[T io.Reader](r T) (byte, error) {
// 	if br, ok := r.(io.ByteReader); ok {
// 		return br.ReadByte()
// 	}
// 	var b [1]byte
// 	_, err := r.Read(b[:])
// 	return b[0], err
// }
//
// // ReadBytes2 is like ReadByte1 but uses a type switch instead.
// func ReadByte2[T io.Reader](r T) (byte, error) {
//         switch br := r.(type) {
//         case io.ByteReader:
//                 return br.ReadByte()
//         }
// 	var b [1]byte
// 	_, err := r.Read(b[:])
// 	return b[0], err
// }
//
// // type assertions and type switches over generic types are strict
// type I3 interface {
//         m(int)
// }
//
// type I4 interface {
//         m() int // different signature from I3.m
// }
//
// func _[T I3](x I3, p T) {
//         // type assertions and type switches over interfaces are not strict
//         _ = x.(I4)
//         switch x.(type) {
//         case I4:
//         }
//
//         // type assertions and type switches over generic types are strict
//         _ = p /* ERROR cannot have dynamic type I4 */.(I4)
//         switch p.(type) {
//         case I4 /* ERROR cannot have dynamic type I4 */ :
//         }
// }
# <翻译结束>


<原文开始>
// type assertions and type switches over generic types lead to errors for now
<原文结束>

# <翻译开始>
// type assertions and type switches over generic types lead to errors for now
# <翻译结束>


<原文开始>
// error messages related to type bounds mention those bounds
<原文结束>

# <翻译开始>
// error messages related to type bounds mention those bounds
# <翻译结束>

