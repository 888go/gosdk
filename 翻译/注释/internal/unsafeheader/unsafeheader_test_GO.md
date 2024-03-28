
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
// TestTypeMatchesReflectType ensures that the name and layout of the
// unsafeheader types matches the corresponding Header types in the reflect
// package.
<原文结束>

# <翻译开始>
// TestTypeMatchesReflectType ensures that the name and layout of the
// unsafeheader types matches the corresponding Header types in the reflect
// package.
# <翻译结束>


<原文开始>
// TestWriteThroughHeader ensures that the headers in the unsafeheader package
// can successfully mutate variables of the corresponding built-in types.
//
// This test is expected to fail under -race (which implicitly enables
// -d=checkptr) if the runtime views the header types as incompatible with the
// underlying built-in types.
<原文结束>

# <翻译开始>
// TestWriteThroughHeader ensures that the headers in the unsafeheader package
// can successfully mutate variables of the corresponding built-in types.
//
// This test is expected to fail under -race (which implicitly enables
// -d=checkptr) if the runtime views the header types as incompatible with the
// underlying built-in types.
# <翻译结束>

