
<原文开始>
// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
# <翻译结束>


<原文开始>
		// For the test below we may decide to convert to int
		// rather than uint and then report a negative shift
		// count instead, which might be a better error. The
		// (minor) difference is that this would restrict the
		// shift count range by half (from all uint values to
		// the positive int values).
		// This depends on the exact spec wording which is not
		// done yet.
		// TODO(gri) revisit and adjust when spec change is done
<原文结束>

# <翻译开始>
		// For the test below we may decide to convert to int
		// rather than uint and then report a negative shift
		// count instead, which might be a better error. The
		// (minor) difference is that this would restrict the
		// shift count range by half (from all uint values to
		// the positive int values).
		// This depends on the exact spec wording which is not
		// done yet.
		// TODO(gri) revisit and adjust when spec change is done
# <翻译结束>


<原文开始>
// basic non-constant shifts
<原文结束>

# <翻译开始>
// basic non-constant shifts
# <翻译结束>


<原文开始>
// 1 has type int32; j == 0
<原文结束>

# <翻译开始>
// 1 has type int32; j == 0
# <翻译结束>


<原文开始>
// 1 has type uint64; k == 1<<33
<原文结束>

# <翻译开始>
// 1 has type uint64; k == 1<<33
# <翻译结束>


<原文开始>
// 1.0 has type int; n == false if ints are 32bits in size
<原文结束>

# <翻译开始>
// 1.0 has type int; n == false if ints are 32bits in size
# <翻译结束>


<原文开始>
// 1 and 2 have type int; o == true if ints are 32bits in size
<原文结束>

# <翻译开始>
// 1 and 2 have type int; o == true if ints are 32bits in size
# <翻译结束>


<原文开始>
// illegal if ints are 32bits in size: 1 has type int, but 1<<33 overflows int
<原文结束>

# <翻译开始>
// illegal if ints are 32bits in size: 1 has type int, but 1<<33 overflows int
# <翻译结束>


<原文开始>
// illegal: 1.0 has type float64, cannot shift
<原文结束>

# <翻译开始>
// illegal: 1.0 has type float64, cannot shift
# <翻译结束>


<原文开始>
// illegal: 1 has type float64, cannot shift
<原文结束>

# <翻译开始>
// illegal: 1 has type float64, cannot shift
# <翻译结束>


<原文开始>
// illegal: 1 has type float32, cannot shift
<原文结束>

# <翻译开始>
// illegal: 1 has type float32, cannot shift
# <翻译结束>


<原文开始>
// 1.0<<33 is a constant shift expression
<原文结束>

# <翻译开始>
// 1.0<<33 is a constant shift expression
# <翻译结束>


<原文开始>
// shifts in comparisons w/ untyped operands
<原文结束>

# <翻译开始>
// shifts in comparisons w/ untyped operands
# <翻译结束>


<原文开始>
// shifts in comparisons w/ typed operands
<原文结束>

# <翻译开始>
// shifts in comparisons w/ typed operands
# <翻译结束>


<原文开始>
// shifts as operands in non-arithmetic operations and as arguments
<原文结束>

# <翻译开始>
// shifts as operands in non-arithmetic operations and as arguments
# <翻译结束>


<原文开始>
	// TODO(gri) Re-enable these tests once types2 has the go/types fixes.
	//           Issue #52080.
	// _ = int32(0x80000000 /* ERROR "overflows int32" */ << s)
	// TODO(rfindley) Eliminate the redundant error here.
	// _ = int32(( /* ERROR "truncated to int32" */ 0x80000000 /* ERROR "truncated to int32" */ + 0i) << s)
<原文结束>

# <翻译开始>
	// TODO(gri) Re-enable these tests once types2 has the go/types fixes.
	//           Issue #52080.
	// _ = int32(0x80000000 /* ERROR "overflows int32" */ << s)
	// TODO(rfindley) Eliminate the redundant error here.
	// _ = int32(( /* ERROR "truncated to int32" */ 0x80000000 /* ERROR "truncated to int32" */ + 0i) << s)
# <翻译结束>


<原文开始>
	// _ = int((1+0i)<<s)
	// _ = int(1.0<<s)
	// _ = int(complex(1, 0)<<s)
<原文结束>

# <翻译开始>
	// _ = int((1+0i)<<s)
	// _ = int(1.0<<s)
	// _ = int(complex(1, 0)<<s)
# <翻译结束>


<原文开始>
	// TODO(gri) The delete below is not type-checked correctly yet.
	// var m1 map[int]string
	// delete(m1, 1<<s)
<原文结束>

# <翻译开始>
	// TODO(gri) The delete below is not type-checked correctly yet.
	// var m1 map[int]string
	// delete(m1, 1<<s)
# <翻译结束>


<原文开始>
// shift examples from shift discussion: better error messages
<原文结束>

# <翻译开始>
// shift examples from shift discussion: better error messages
# <翻译结束>


<原文开始>
	// TODO(gri) the error messages for these two are incorrect - disabled for now
	// _ = complex64(1<<s)
	// _ = complex64(1.<<s)
<原文结束>

# <翻译开始>
	// TODO(gri) the error messages for these two are incorrect - disabled for now
	// _ = complex64(1<<s)
	// _ = complex64(1.<<s)
# <翻译结束>


<原文开始>
	// various originally failing snippets of code from the std library
	// from src/compress/lzw/reader.go:90
<原文结束>

# <翻译开始>
	// various originally failing snippets of code from the std library
	// from src/compress/lzw/reader.go:90
# <翻译结束>


<原文开始>
// from src/debug/dwarf/buf.go:116
<原文结束>

# <翻译开始>
// from src/debug/dwarf/buf.go:116
# <翻译结束>


<原文开始>
// from src/encoding/asn1/asn1.go:160
<原文结束>

# <翻译开始>
// from src/encoding/asn1/asn1.go:160
# <翻译结束>


<原文开始>
// from src/math/big/rat.go:140
<原文结束>

# <翻译开始>
// from src/math/big/rat.go:140
# <翻译结束>


<原文开始>
// from src/net/interface.go:51
<原文结束>

# <翻译开始>
// from src/net/interface.go:51
# <翻译结束>


<原文开始>
// from src/runtime/softfloat64.go:234
<原文结束>

# <翻译开始>
// from src/runtime/softfloat64.go:234
# <翻译结束>


<原文开始>
// from src/strconv/atof.go:326
<原文结束>

# <翻译开始>
// from src/strconv/atof.go:326
# <翻译结束>


<原文开始>
// from src/route_bsd.go:82
<原文结束>

# <翻译开始>
// from src/route_bsd.go:82
# <翻译结束>


<原文开始>
// from src/text/scanner/scanner.go:540
<原文结束>

# <翻译开始>
// from src/text/scanner/scanner.go:540
# <翻译结束>


<原文开始>
// example from issue 11325
<原文结束>

# <翻译开始>
// example from issue 11325
# <翻译结束>


<原文开始>
// example from issue 11594
<原文结束>

# <翻译开始>
// example from issue 11594
# <翻译结束>


<原文开始>
// example from issue 22969
<原文结束>

# <翻译开始>
// example from issue 22969
# <翻译结束>

