
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
// Field returns the i'th field of the struct v.
// It panics if v's Kind is not Struct or i is out of range.
<原文结束>

# <翻译开始>
// Field returns the i'th field of the struct v.
// It panics if v's Kind is not Struct or i is out of range.
# <翻译结束>


<原文开始>
// Inherit permission bits from v, but clear flagEmbedRO.
<原文结束>

# <翻译开始>
// Inherit permission bits from v, but clear flagEmbedRO.
# <翻译结束>


<原文开始>
// Using an unexported field forces flagRO.
<原文结束>

# <翻译开始>
// Using an unexported field forces flagRO.
# <翻译结束>


<原文开始>
	// Either flagIndir is set and v.ptr points at struct,
	// or flagIndir is not set and v.ptr is the actual struct data.
	// In the former case, we want v.ptr + offset.
	// In the latter case, we must have field.offset = 0,
	// so v.ptr + field.offset is still the correct address.
<原文结束>

# <翻译开始>
	// Either flagIndir is set and v.ptr points at struct,
	// or flagIndir is not set and v.ptr is the actual struct data.
	// In the former case, we want v.ptr + offset.
	// In the latter case, we must have field.offset = 0,
	// so v.ptr + field.offset is still the correct address.
# <翻译结束>


<原文开始>
// Field returns the i'th struct field.
<原文结束>

# <翻译开始>
// Field returns the i'th struct field.
# <翻译结束>


<原文开始>
// Zero returns a Value representing the zero value for the specified type.
// The result is different from the zero value of the Value struct,
// which represents no value at all.
// For example, Zero(TypeOf(42)) returns a Value with Kind Int and value 0.
// The returned value is neither addressable nor settable.
<原文结束>

# <翻译开始>
// Zero returns a Value representing the zero value for the specified type.
// The result is different from the zero value of the Value struct,
// which represents no value at all.
// For example, Zero(TypeOf(42)) returns a Value with Kind Int and value 0.
// The returned value is neither addressable nor settable.
# <翻译结束>


<原文开始>
// ToInterface returns v's current value as an interface{}.
// It is equivalent to:
//
//	var i interface{} = (v's underlying value)
//
// It panics if the Value was obtained by accessing
// unexported struct fields.
<原文结束>

# <翻译开始>
// ToInterface returns v's current value as an interface{}.
// It is equivalent to:
//
//	var i interface{} = (v's underlying value)
//
// It panics if the Value was obtained by accessing
// unexported struct fields.
# <翻译结束>

