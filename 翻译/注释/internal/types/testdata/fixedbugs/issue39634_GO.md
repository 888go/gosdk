
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
// Examples adjusted to match new [T any] syntax for type parameters.
// Also, previously permitted empty type parameter lists and instantiations
// are now syntax errors.
<原文结束>

# <翻译开始>
// Examples adjusted to match new [T any] syntax for type parameters.
// Also, previously permitted empty type parameter lists and instantiations
// are now syntax errors.
# <翻译结束>


<原文开始>
// crash 2
// Disabled: empty []'s are now syntax errors. This example leads to too many follow-on errors.
// type Numeric2 interface{t2 /* ERROR not a type */ }
// func t2[T Numeric2](s[]T){0 /* ERROR not a type */ []{s /* ERROR cannot index */ [0][0]}}
<原文结束>

# <翻译开始>
// crash 2
// Disabled: empty []'s are now syntax errors. This example leads to too many follow-on errors.
// type Numeric2 interface{t2 /* ERROR not a type */ }
// func t2[T Numeric2](s[]T){0 /* ERROR not a type */ []{s /* ERROR cannot index */ [0][0]}}
# <翻译结束>


<原文开始>
// ERROR cannot use a type parameter as RHS in type declaration
<原文结束>

# <翻译开始>
// ERROR cannot use a type parameter as RHS in type declaration
# <翻译结束>

