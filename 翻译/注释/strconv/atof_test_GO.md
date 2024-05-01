
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
// Hexadecimal floating-point.
<原文结束>

# <翻译开始>
// 十六进制浮点数。. md5:de2d821e70967e8f
# <翻译结束>


<原文开始>
// next float64 - too large
<原文结束>

# <翻译开始>
// next float64 - 太大了. md5:e595d550db35fd95
# <翻译结束>


<原文开始>
	// the border is ...158079
	// borderline - okay
<原文结束>

# <翻译开始>
// 边界是...158079
// 边缘 - 可接受
// md5:283760bcac5d4566
# <翻译结束>


<原文开始>
// Near denormals and denormals.
<原文结束>

# <翻译开始>
// 接近于denormals和denormals。. md5:5213e06b56063627
# <翻译结束>


<原文开始>
// rounded down (half to even)
<原文结束>

# <翻译开始>
// 向下取整（偶数舍入）. md5:600817f8100c1677
# <翻译结束>


<原文开始>
// try to overflow exponent
<原文结束>

# <翻译开始>
// 尝试使指数溢出. md5:9b2de22c589b27df
# <翻译结束>


<原文开始>
// https://www.exploringbinary.com/java-hangs-when-converting-2-2250738585072012e-308/
<原文结束>

# <翻译开始>
// https://www.exploringbinary.com/java-hangs-when-converting-2-2250738585072012e-308/
// 这段Go注释的中文翻译是：
// 
// 当将2.2250738585072012e-308转换时，Java会挂起（或停止响应）：https://www.exploringbinary.com/java-hangs-when-converting-2-2250738585072012e-308/. md5:230c0c4f17858772
# <翻译结束>


<原文开始>
// https://www.exploringbinary.com/php-hangs-on-numeric-value-2-2250738585072011e-308/
<原文结束>

# <翻译开始>
// https://www.exploringbinary.com/php-hangs-on-numeric-value-2.2250738585072011e-308/ 
// 
// 这段Go注释翻译成中文是：
// 
// 这是一个链接，内容是关于PHP中数值2.2250738585072011e-308会导致程序挂起的问题（https://www.exploringbinary.com/php-hangs-on-numeric-value-2.2250738585072011e-308/）. md5:e6656740a836a448
# <翻译结束>


<原文开始>
// A very large number (initially wrongly parsed by the fast algorithm).
<原文结束>

# <翻译开始>
// 一个非常大的数字（最初被快速算法错误解析）。. md5:aaec7675e0bff74a
# <翻译结束>


<原文开始>
// A different kind of very large number.
<原文结束>

# <翻译开始>
// 一种不同类型的非常大的数字。. md5:6522045cedd96e71
# <翻译结束>


<原文开始>
	// Exactly halfway between 1 and math.Nextafter(1, 2).
	// Round to even (down).
<原文结束>

# <翻译开始>
// 完整位于1和math.Nextafter(1, 2)之间的中间点。
// 四舍五入到偶数（向下取整）。
// md5:e79d24edcbd74c15
# <翻译结束>


<原文开始>
// Slightly lower; still round down.
<原文结束>

# <翻译开始>
// 稍微低一点；仍然向下舍入。. md5:dca3ee1e60a729a5
# <翻译结束>


<原文开始>
// Slightly higher; round up.
<原文结束>

# <翻译开始>
// 略高一点；向上取整。. md5:afd45945f9caf31d
# <翻译结束>


<原文开始>
// Slightly higher, but you have to read all the way to the end.
<原文结束>

# <翻译开始>
// 稍微高一点，但你必须读到结尾才能明白。. md5:b8273dbe0862f4f3
# <翻译结束>


<原文开始>
	// Halfway between x := math.Nextafter(1, 2) and math.Nextafter(x, 2)
	// Round to even (up).
<原文结束>

# <翻译开始>
// 在 `x := math.Nextafter(1, 2)` 和 `math.Nextafter(x, 2)` 之间的中间点
// 向偶数方向四舍五入（向上）
// md5:3ba5d040c89bb711
# <翻译结束>


<原文开始>
	// Halfway between 1090544144181609278303144771584 and 1090544144181609419040633126912
	// (15497564393479157p+46, should round to even 15497564393479156p+46, issue 36657)
<原文结束>

# <翻译开始>
// 在1090544144181609278303144771584和1090544144181609419040633126912之间的中间值
// （15497564393479157p+46，应该四舍五入到最接近的偶数15497564393479156p+46，参见问题36657）
// md5:4b42731b30cfcc6c
# <翻译结束>


<原文开始>
// slightly above, rounds up
<原文结束>

# <翻译开始>
// 略高于，向上取整. md5:a7a383a63db55e7c
# <翻译结束>


<原文开始>
	// Exactly halfway between 1 and the next float32.
	// Round to even (down).
<原文结束>

# <翻译开始>
// 在1和下一个float32之间精确的一半。
// 向下取整（偶数）。
// md5:f1baed8995860339
# <翻译结束>


<原文开始>
// largest float32: (1<<128) * (1 - 2^-24)
<原文结束>

# <翻译开始>
// 最大 float32：(1<<128) * (1 - 2^-24) 
// 
// 这段Go语言注释的意思是：这是计算float32类型的最大值的一种方法。`(1<<128)` 表示将1左移128位，得到的结果是一个非常大的数值。`1 - 2^-24` 表示1减去2的负24次方，这个表达式的结果接近于2的-24次方，但由于浮点数精度问题，实际上会稍微大于这个值。两者相乘，就可以得到float32类型的上限，但由于浮点数表示有局限性，这个结果并不完全精确。. md5:9f39d19e096aa65d
# <翻译结束>


<原文开始>
// next float32 - too large
<原文结束>

# <翻译开始>
// 下一个 float32 - 太大了. md5:f1b130c05a88712b
# <翻译结束>


<原文开始>
	// the border is 3.40282356779...e+38
	// borderline - okay
<原文结束>

# <翻译开始>
// 边界值为 3.40282356779...e+38
// 边界线 - 正常
// md5:bf670afe49bf180c
# <翻译结束>


<原文开始>
// Denormals: less than 2^-126
<原文结束>

# <翻译开始>
// Denormals: 小于2^-126的数值. md5:3dd4289972b09074
# <翻译结束>


<原文开始>
	// 2^92 = 8388608p+69 = 4951760157141521099596496896 (4.9517602e27)
	// is an exact power of two that needs 8 decimal digits to be correctly
	// parsed back.
	// The float32 before is 16777215p+68 = 4.95175986e+27
	// The halfway is 4.951760009. A bad algorithm that thinks the previous
	// float32 is 8388607p+69 will shorten incorrectly to 4.95176e+27.
<原文结束>

# <翻译开始>
// 2^92 = 8388608p+69 = 4951760157141521099596496896 (4.9517602e27) 是一个精确的2的幂，需要8位小数才能正确解析回原来的值。
// 前面的float32为16777215p+68 = 4.95175986e+27
// 中间点是4.951760009。如果使用一个错误的算法，将前一个float32误认为是8388607p+69，会错误地简化为4.95176e+27。
// md5:2ecc7a844307a81e
# <翻译结束>


<原文开始>
	// The atof routines return NumErrors wrapping
	// the error and the string. Convert the table above.
<原文结束>

# <翻译开始>
// atof 函数返回一个包含错误和字符串的 NumErrors 包装。请将上面的表格转换为中文。
// md5:ef4283f67521389e
# <翻译结束>


<原文开始>
// Generate random inputs for tests and benchmarks
<原文结束>

# <翻译开始>
// 为测试和基准生成随机输入. md5:92a91f9ab7ee46d6
# <翻译结束>


<原文开始>
		// Adding characters that do not extend a number should not invalidate it.
		// Test a few. The "i" and "init" cases test that we accept "infi", "infinit"
		// correctly as "inf" with suffix.
<原文结束>

# <翻译开始>
// 添加不会扩展数字的字符不应该使其无效。
// 测试几个例子。"i" 和 "init" 的情况测试我们是否正确接受 "infi"、"infinit" 作为带有后缀的 "inf"。
// md5:109092a575e26bbc
# <翻译结束>


<原文开始>
	// Issue 2917.
	// This test will break the optimized conversion if the
	// FPU is using 80-bit registers instead of 64-bit registers,
	// usually because the operating system initialized the
	// thread with 80-bit precision and the Go runtime didn't
	// fix the FP control word.
<原文结束>

# <翻译开始>
// 问题2917。
// 如果FPU使用的是80位寄存器而不是64位寄存器，这个测试将会破坏优化的转换，通常是因为操作系统使用80位精度初始化了线程，而Go运行时没有修复浮点控制字。
// md5:4872af9151884505
# <翻译结束>


<原文开始>
// TestRoundTrip32 tries a fraction of all finite positive float32 values.
<原文结束>

# <翻译开始>
// TestRoundTrip32尝试所有有限的正浮点32值的很小一部分。. md5:bcc3b233d2561c2a
# <翻译结束>


<原文开始>
// Issue 42297: a lot of code in the wild accidentally calls ParseFloat(s, 10)
// or ParseFloat(s, 0), so allow bitSize values other than 32 and 64.
<原文结束>

# <翻译开始>
// Issue 42297：很多现有代码可能会误调用 ParseFloat(s, 10)
// 或 ParseFloat(s, 0)，因此允许 bitSize 的值不仅限于 32 和 64。
// md5:5a23d02fde2bdecb
# <翻译结束>
 
