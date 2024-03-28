
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
// For now, a lone type parameter is not permitted as RHS in a type declaration (issue #45639).
// type T[P any] P
// type A = T  // ERROR cannot use generic type
// var x A[int]
// var _ A
//
// type B = T[int]
// var y B = x
// var _ B /* ERROR not a generic type */ [int]
<原文结束>

# <翻译开始>
// For now, a lone type parameter is not permitted as RHS in a type declaration (issue #45639).
// type T[P any] P
// type A = T  // ERROR cannot use generic type
// var x A[int]
// var _ A
//
// type B = T[int]
// var y B = x
// var _ B /* ERROR not a generic type */ [int]
# <翻译结束>

