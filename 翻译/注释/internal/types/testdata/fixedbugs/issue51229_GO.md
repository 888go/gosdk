
<原文开始>
// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
# <翻译结束>


<原文开始>
// Constraint type inference should be independent of the
// ordering of the type parameter declarations. Try all
// permutations in the test case below.
// Permutations produced by https://go.dev/play/p/PHcZNGJTEBZ.
<原文结束>

# <翻译开始>
// Constraint type inference should be independent of the
// ordering of the type parameter declarations. Try all
// permutations in the test case below.
// Permutations produced by https://go.dev/play/p/PHcZNGJTEBZ.
# <翻译结束>


<原文开始>
// Constraint type inference may have to iterate.
// Again, the order of the type parameters shouldn't matter.
<原文结束>

# <翻译开始>
// Constraint type inference may have to iterate.
// Again, the order of the type parameters shouldn't matter.
# <翻译结束>


<原文开始>
// Worst-case scenario.
// There are 10 unknown type parameters. In each iteration of
// constraint type inference we infer one more, from right to left.
// Each iteration looks repeatedly at all 11 type parameters,
// requiring a total of 10*11 = 110 iterations with the current
// implementation. Pathological case.
<原文结束>

# <翻译开始>
// Worst-case scenario.
// There are 10 unknown type parameters. In each iteration of
// constraint type inference we infer one more, from right to left.
// Each iteration looks repeatedly at all 11 type parameters,
// requiring a total of 10*11 = 110 iterations with the current
// implementation. Pathological case.
# <翻译结束>


<原文开始>
// Examples with channel constraints and tilde.
<原文结束>

# <翻译开始>
// Examples with channel constraints and tilde.
# <翻译结束>


<原文开始>
// core(P) == chan<- int   (single type, no tilde)
<原文结束>

# <翻译开始>
// core(P) == chan<- int   (single type, no tilde)
# <翻译结束>


<原文开始>
// core(P) == ~chan<- int  (tilde)
<原文结束>

# <翻译开始>
// core(P) == ~chan<- int  (tilde)
# <翻译结束>


<原文开始>
// core(P) == chan<- E     (single type, no tilde)
<原文结束>

# <翻译开始>
// core(P) == chan<- E     (single type, no tilde)
# <翻译结束>


<原文开始>
// core(P) == ~chan<- E    (tilde)
<原文结束>

# <翻译开始>
// core(P) == ~chan<- E    (tilde)
# <翻译结束>


<原文开始>
// core(P) == chan<- int   (not a single type)
<原文结束>

# <翻译开始>
// core(P) == chan<- int   (not a single type)
# <翻译结束>


<原文开始>
// P can be inferred as there's a single specific type and no tilde.
<原文结束>

# <翻译开始>
// P can be inferred as there's a single specific type and no tilde.
# <翻译结束>


<原文开始>
// P cannot be inferred as there's a tilde.
<原文结束>

# <翻译开始>
// P cannot be inferred as there's a tilde.
# <翻译结束>


<原文开始>
// P cannot be inferred as there's more than one specific type and a tilde.
<原文结束>

# <翻译开始>
// P cannot be inferred as there's more than one specific type and a tilde.
# <翻译结束>


<原文开始>
// P cannot be inferred as there's more than one specific type.
<原文结束>

# <翻译开始>
// P cannot be inferred as there's more than one specific type.
# <翻译结束>

