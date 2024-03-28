
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
	// TODO(gri) The compiler limits integers to 512 bit and thus
	//           we cannot compute the value (1<<(1023 - 1 + 52))
	//           without overflow. For now we match the compiler.
	//           See also issue #44057.
	// smallestFloat64 = 1.0 / (1<<(1023 - 1 + 52))
<原文结束>

# <翻译开始>
	// TODO(gri) The compiler limits integers to 512 bit and thus
	//           we cannot compute the value (1<<(1023 - 1 + 52))
	//           without overflow. For now we match the compiler.
	//           See also issue #44057.
	// smallestFloat64 = 1.0 / (1<<(1023 - 1 + 52))
# <翻译结束>


<原文开始>
	// TODO(gri) The compiler limits integers to 512 bit and thus
	//           we cannot compute the value 1<<1023
	//           without overflow. For now we match the compiler.
	//           See also issue #44057.
	// maxFloat64 = 1<<1023 * (1<<53 - 1) / (1.0<<52)
<原文结束>

# <翻译开始>
	// TODO(gri) The compiler limits integers to 512 bit and thus
	//           we cannot compute the value 1<<1023
	//           without overflow. For now we match the compiler.
	//           See also issue #44057.
	// maxFloat64 = 1<<1023 * (1<<53 - 1) / (1.0<<52)
# <翻译结束>


<原文开始>
// TODO(gri) find smaller deltas below
<原文结束>

# <翻译开始>
// TODO(gri) find smaller deltas below
# <翻译结束>


<原文开始>
// Initialization of typed constant and conversion are the same:
<原文结束>

# <翻译开始>
// Initialization of typed constant and conversion are the same:
# <翻译结束>

