
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
// Numeric is type bound that matches any numeric type.
// It would likely be in a constraints package in the standard library.
<原文结束>

# <翻译开始>
// Numeric is type bound that matches any numeric type.
// It would likely be in a constraints package in the standard library.
# <翻译结束>


<原文开始>
// NumericAbs matches numeric types with an Abs method.
<原文结束>

# <翻译开始>
// NumericAbs matches numeric types with an Abs method.
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
// OrderedNumeric is a type bound that matches numeric types that support the < operator.
<原文结束>

# <翻译开始>
// OrderedNumeric is a type bound that matches numeric types that support the < operator.
# <翻译结束>


<原文开始>
// Complex is a type bound that matches the two complex types, which do not have a < operator.
<原文结束>

# <翻译开始>
// Complex is a type bound that matches the two complex types, which do not have a < operator.
# <翻译结束>


<原文开始>
// For now, a lone type parameter is not permitted as RHS in a type declaration (issue #45639).
// // OrderedAbs is a helper type that defines an Abs method for
// // ordered numeric types.
// type OrderedAbs[T OrderedNumeric] T
// 
// func (a OrderedAbs[T]) Abs() OrderedAbs[T] {
// 	if a < 0 {
// 		return -a
// 	}
// 	return a
// }
// 
// // ComplexAbs is a helper type that defines an Abs method for
// // complex types.
// type ComplexAbs[T Complex] T
// 
// func (a ComplexAbs[T]) Abs() ComplexAbs[T] {
// 	r := float64(real(a))
// 	i := float64(imag(a))
// 	d := math.Sqrt(r * r + i * i)
// 	return ComplexAbs[T](complex(d, 0))
// }
// 
// func OrderedAbsDifference[T OrderedNumeric](a, b T) T {
// 	return T(AbsDifference(OrderedAbs[T](a), OrderedAbs[T](b)))
// }
// 
// func ComplexAbsDifference[T Complex](a, b T) T {
// 	return T(AbsDifference(ComplexAbs[T](a), ComplexAbs[T](b)))
// }
<原文结束>

# <翻译开始>
// For now, a lone type parameter is not permitted as RHS in a type declaration (issue #45639).
// // OrderedAbs is a helper type that defines an Abs method for
// // ordered numeric types.
// type OrderedAbs[T OrderedNumeric] T
// 
// func (a OrderedAbs[T]) Abs() OrderedAbs[T] {
// 	if a < 0 {
// 		return -a
// 	}
// 	return a
// }
// 
// // ComplexAbs is a helper type that defines an Abs method for
// // complex types.
// type ComplexAbs[T Complex] T
// 
// func (a ComplexAbs[T]) Abs() ComplexAbs[T] {
// 	r := float64(real(a))
// 	i := float64(imag(a))
// 	d := math.Sqrt(r * r + i * i)
// 	return ComplexAbs[T](complex(d, 0))
// }
// 
// func OrderedAbsDifference[T OrderedNumeric](a, b T) T {
// 	return T(AbsDifference(OrderedAbs[T](a), OrderedAbs[T](b)))
// }
// 
// func ComplexAbsDifference[T Complex](a, b T) T {
// 	return T(AbsDifference(ComplexAbs[T](a), ComplexAbs[T](b)))
// }
# <翻译结束>

