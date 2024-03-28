
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
// Parameterized type declarations
<原文结束>

# <翻译开始>
// Parameterized type declarations
# <翻译结束>


<原文开始>
// For now, a lone type parameter is not permitted as RHS in a type declaration (issue #45639).
<原文结束>

# <翻译开始>
// For now, a lone type parameter is not permitted as RHS in a type declaration (issue #45639).
# <翻译结束>


<原文开始>
// ERROR cannot use a type parameter as RHS in type declaration
<原文结束>

# <翻译开始>
// ERROR cannot use a type parameter as RHS in type declaration
# <翻译结束>


<原文开始>
// int should still be in scope chain
<原文结束>

# <翻译开始>
// int should still be in scope chain
# <翻译结束>


<原文开始>
// Alias type declarations cannot have type parameters.
// Issue #46477 proposses to change that.
<原文结束>

# <翻译开始>
// Alias type declarations cannot have type parameters.
// Issue #46477 proposses to change that.
# <翻译结束>


<原文开始>
// Pending clarification of #46477 we disallow aliases
// of generic types.
<原文结束>

# <翻译开始>
// Pending clarification of #46477 we disallow aliases
// of generic types.
# <翻译结束>


<原文开始>
// ERROR cannot use generic type
<原文结束>

# <翻译开始>
// ERROR cannot use generic type
# <翻译结束>


<原文开始>
// Parameterized type instantiations
<原文结束>

# <翻译开始>
// Parameterized type instantiations
# <翻译结束>


<原文开始>
// ERROR expected type argument list
<原文结束>

# <翻译开始>
// ERROR expected type argument list
# <翻译结束>


<原文开始>
// TODO(gri) better error messages
<原文结束>

# <翻译开始>
// TODO(gri) better error messages
# <翻译结束>


<原文开始>
// Parameterized types containing parameterized types
<原文结束>

# <翻译开始>
// Parameterized types containing parameterized types
# <翻译结束>


<原文开始>
// Self-recursive generic types are not permitted
<原文结束>

# <翻译开始>
// Self-recursive generic types are not permitted
# <翻译结束>

