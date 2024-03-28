
<原文开始>
// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
# <翻译结束>


<原文开始>
// T1 declared before its methods.
<原文结束>

# <翻译开始>
// T1 declared before its methods.
# <翻译结束>


<原文开始>
// Conflict between embedded field and method name,
// with the embedded field being a basic type.
<原文结束>

# <翻译开始>
// Conflict between embedded field and method name,
// with the embedded field being a basic type.
# <翻译结束>


<原文开始>
// Disabled for now: LookupFieldOrMethod will find Pointer even though
// it's double-declared (it would cost extra in the common case to verify
// this). But the MethodSet computation will not find it due to the name
// collision caused by the double-declaration, leading to an internal
// inconsistency while we are verifying one computation against the other.
// var _ = T1c{}.Pointer
<原文结束>

# <翻译开始>
// Disabled for now: LookupFieldOrMethod will find Pointer even though
// it's double-declared (it would cost extra in the common case to verify
// this). But the MethodSet computation will not find it due to the name
// collision caused by the double-declaration, leading to an internal
// inconsistency while we are verifying one computation against the other.
// var _ = T1c{}.Pointer
# <翻译结束>


<原文开始>
// T2's method declared before the type.
<原文结束>

# <翻译开始>
// T2's method declared before the type.
# <翻译结束>


<原文开始>
// Methods declared without a declared type.
<原文结束>

# <翻译开始>
// Methods declared without a declared type.
# <翻译结束>


<原文开始>
// Methods with receiver base type declared in another file.
<原文结束>

# <翻译开始>
// Methods with receiver base type declared in another file.
# <翻译结束>


<原文开始>
// Methods of non-struct type.
<原文结束>

# <翻译开始>
// Methods of non-struct type.
# <翻译结束>


<原文开始>
// Methods associated with an interface.
<原文结束>

# <翻译开始>
// Methods associated with an interface.
# <翻译结束>


<原文开始>
// Methods associated with a named pointer type.
<原文结束>

# <翻译开始>
// Methods associated with a named pointer type.
# <翻译结束>


<原文开始>
// Methods with zero or multiple receivers.
<原文结束>

# <翻译开始>
// Methods with zero or multiple receivers.
# <翻译结束>


<原文开始>
// Methods associated with non-local or unnamed types.
<原文结束>

# <翻译开始>
// Methods associated with non-local or unnamed types.
# <翻译结束>


<原文开始>
// Unsafe.Pointer is treated like a pointer when used as receiver type.
<原文结束>

# <翻译开始>
// Unsafe.Pointer is treated like a pointer when used as receiver type.
# <翻译结束>


<原文开始>
// Double declarations across package files
<原文结束>

# <翻译开始>
// Double declarations across package files
# <翻译结束>

