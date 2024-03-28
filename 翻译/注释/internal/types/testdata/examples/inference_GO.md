
<原文开始>
// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
# <翻译结束>


<原文开始>
// This file shows some examples of type inference.
<原文结束>

# <翻译开始>
// This file shows some examples of type inference.
# <翻译结束>


<原文开始>
// min can be called with explicit instantiation.
<原文结束>

# <翻译开始>
// min can be called with explicit instantiation.
# <翻译结束>


<原文开始>
	// Alternatively, the type argument can be inferred from
	// one of the arguments. Untyped arguments will be considered
	// last.
<原文结束>

# <翻译开始>
	// Alternatively, the type argument can be inferred from
	// one of the arguments. Untyped arguments will be considered
	// last.
# <翻译结束>


<原文开始>
// mixed can be called with explicit instantiation.
<原文结束>

# <翻译开始>
// mixed can be called with explicit instantiation.
# <翻译结束>


<原文开始>
	// Alternatively, partial type arguments may be provided
	// (from left to right), and the other may be inferred.
<原文结束>

# <翻译开始>
	// Alternatively, partial type arguments may be provided
	// (from left to right), and the other may be inferred.
# <翻译结束>


<原文开始>
	// Provided type arguments always take precedence over
	// inferred types.
<原文结束>

# <翻译开始>
	// Provided type arguments always take precedence over
	// inferred types.
# <翻译结束>


<原文开始>
// related1 can be called with explicit instantiation.
<原文结束>

# <翻译开始>
// related1 can be called with explicit instantiation.
# <翻译结束>


<原文开始>
	// Alternatively, the 2nd type argument can be inferred
	// from the first one through constraint type inference.
<原文结束>

# <翻译开始>
	// Alternatively, the 2nd type argument can be inferred
	// from the first one through constraint type inference.
# <翻译结束>


<原文开始>
	// A type argument inferred from another explicitly provided
	// type argument overrides whatever value argument type is given.
<原文结束>

# <翻译开始>
	// A type argument inferred from another explicitly provided
	// type argument overrides whatever value argument type is given.
# <翻译结束>


<原文开始>
	// A type argument may be inferred from a value argument
	// and then help infer another type argument via constraint
	// type inference.
<原文结束>

# <翻译开始>
	// A type argument may be inferred from a value argument
	// and then help infer another type argument via constraint
	// type inference.
# <翻译结束>


<原文开始>
// related2 can be called with explicit instantiation.
<原文结束>

# <翻译开始>
// related2 can be called with explicit instantiation.
# <翻译结束>


<原文开始>
	// A type argument may be inferred from a value argument
	// and then help infer another type argument via constraint
	// type inference. Untyped arguments are always considered
	// last.
<原文结束>

# <翻译开始>
	// A type argument may be inferred from a value argument
	// and then help infer another type argument via constraint
	// type inference. Untyped arguments are always considered
	// last.
# <翻译结束>


<原文开始>
// TODO(gri) fix error position
<原文结束>

# <翻译开始>
// TODO(gri) fix error position
# <翻译结束>


<原文开始>
// related3 can be instantiated explicitly
<原文结束>

# <翻译开始>
// related3 can be instantiated explicitly
# <翻译结束>


<原文开始>
	// The 2nd type argument cannot be inferred from the first
	// one because there's two possible choices: []Elem and
	// List[Elem].
<原文结束>

# <翻译开始>
	// The 2nd type argument cannot be inferred from the first
	// one because there's two possible choices: []Elem and
	// List[Elem].
# <翻译结束>

