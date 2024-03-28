
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
// This file shows some examples of generic constraint interfaces.
<原文结束>

# <翻译开始>
// This file shows some examples of generic constraint interfaces.
# <翻译结束>


<原文开始>
// Arbitrary types may be embedded like interfaces.
<原文结束>

# <翻译开始>
// Arbitrary types may be embedded like interfaces.
# <翻译结束>


<原文开始>
// Types may be combined into a union.
<原文结束>

# <翻译开始>
// Types may be combined into a union.
# <翻译结束>


<原文开始>
// Union terms must describe disjoint (non-overlapping) type sets.
<原文结束>

# <翻译开始>
// Union terms must describe disjoint (non-overlapping) type sets.
# <翻译结束>


<原文开始>
// interfaces (here: union) are ignored when checking for overlap
<原文结束>

# <翻译开始>
// interfaces (here: union) are ignored when checking for overlap
# <翻译结束>


<原文开始>
// For now we do not permit interfaces with methods in unions.
<原文结束>

# <翻译开始>
// For now we do not permit interfaces with methods in unions.
# <翻译结束>


<原文开始>
// Tilde is not permitted on defined types or interfaces.
<原文结束>

# <翻译开始>
// Tilde is not permitted on defined types or interfaces.
# <翻译结束>


<原文开始>
// Stand-alone type parameters are not permitted as elements or terms in unions.
<原文结束>

# <翻译开始>
// Stand-alone type parameters are not permitted as elements or terms in unions.
# <翻译结束>


<原文开始>
// Multiple embedded union elements are intersected. The order in which they
// appear in the interface doesn't matter since intersection is a symmetric
// operation.
<原文结束>

# <翻译开始>
// Multiple embedded union elements are intersected. The order in which they
// appear in the interface doesn't matter since intersection is a symmetric
// operation.
# <翻译结束>


<原文开始>
// Here the intersections are empty - there's no type that's in the type set of T.
<原文结束>

# <翻译开始>
// Here the intersections are empty - there's no type that's in the type set of T.
# <翻译结束>


<原文开始>
// Union elements may be interfaces as long as they don't define
// any methods or embed comparable.
<原文结束>

# <翻译开始>
// Union elements may be interfaces as long as they don't define
// any methods or embed comparable.
# <翻译结束>

