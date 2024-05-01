
<原文开始>
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
//版权所有2009年Go作者。所有权利保留。
//使用此源代码受BSD风格
//可以在LICENSE文件中找到的许可证。
// md5:2e9dc81828a3be8a
# <翻译结束>


<原文开始>
// Wrapper around strconv.ParseFloat(x, 64).  Handles dddddp+ddd (binary exponent)
// itself, passes the rest on to strconv.ParseFloat.
<原文结束>

# <翻译开始>
// 这是一个 strconv.ParseFloat(x, 64) 的包装器。它自身处理 dddddp+ddd（二进制指数）部分，其余的传递给 strconv.ParseFloat。
// md5:da10d8aa57bdf05c
# <翻译结束>


<原文开始>
		// We expect that v*pow2(e) fits in a float64,
		// but pow2(e) by itself may not. Be careful.
<原文结束>

# <翻译开始>
// 我们期望 v*pow2(e) 能够适应一个 float64 的范围，
// 但 pow2(e) 单独计算时可能超出了这个范围。因此需要格外小心。
// md5:f5e925857d3ba2e6
# <翻译结束>


<原文开始>
// Wrapper around strconv.ParseFloat(x, 32).  Handles dddddp+ddd (binary exponent)
// itself, passes the rest on to strconv.ParseFloat.
<原文结束>

# <翻译开始>
// 包装了strconv.ParseFloat(x, 32)。处理dddddp+ddd（二进制指数）本身，其余的传递给strconv.ParseFloat。
// md5:a03fed9d2a4eb26f
# <翻译结束>

