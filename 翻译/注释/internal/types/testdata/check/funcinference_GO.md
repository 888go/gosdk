
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
// Embedding stand-alone type parameters is not permitted for now. Disabled.
// func f3[A any, B interface{~C}, C interface{~*A}](A, B, C)
// func _() {
// 	f := f3[int]
// 	var x int
// 	f(x, &x, &x)
// 	f3(x, &x, &x)
// }
<原文结束>

# <翻译开始>
// Embedding stand-alone type parameters is not permitted for now. Disabled.
// func f3[A any, B interface{~C}, C interface{~*A}](A, B, C)
// func _() {
// 	f := f3[int]
// 	var x int
// 	f(x, &x, &x)
// 	f3(x, &x, &x)
// }
# <翻译结束>


<原文开始>
// More realistic examples
<原文结束>

# <翻译开始>
// More realistic examples
# <翻译结束>


<原文开始>
		// The type of &result[i] is *T which is in the type set
		// of Setter, so we can convert it to PT.
<原文结束>

# <翻译开始>
		// The type of &result[i] is *T which is in the type set
		// of Setter, so we can convert it to PT.
# <翻译结束>


<原文开始>
// real code should not ignore the error
<原文结束>

# <翻译开始>
// real code should not ignore the error
# <翻译结束>

