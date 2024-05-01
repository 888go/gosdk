
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
	// We want to test that a ticker takes as much time as expected.
	// Since we don't want the test to run for too long, we don't
	// want to use lengthy times. This makes the test inherently flaky.
	// Start with a short time, but try again with a long one if the
	// first test fails.
<原文结束>

# <翻译开始>
// 我们想要测试ticker是否按预期消耗时间。
// 由于我们不希望测试运行时间过长，因此不想使用过长的时间间隔。
// 这使得测试本质上可能不稳定。
// 先从较短的时间间隔开始，但如果第一次测试失败，再尝试使用较长的时间间隔。
// md5:ad7e20ba1cc8ab8f
# <翻译结束>


<原文开始>
// On Darwin ARM64 the tick frequency seems limited. Issue 35692.
<原文结束>

# <翻译开始>
// 在Darwin ARM64上，tick频率似乎受到限制。问题35692。. md5:7a0c4fbdc27fe8a7
# <翻译结束>


<原文开始>
		// The following test will run ticker count/2 times then reset
		// the ticker to double the duration for the rest of count/2.
		// Since tick frequency is limited on Darwin ARM64, use even
		// number to give the ticks more time to let the test pass.
		// See CL 220638.
<原文结束>

# <翻译开始>
// 下面的测试会运行计时器 `count/2` 次，然后将计时器的周期重置为剩余 `count/2` 的两倍。由于 Darwin ARM64 系统的计时频率有限，我们使用偶数来给计时器更多的时间让测试通过。参考 CL（Change List）220638。
// md5:9679bd76bf1f345b
# <翻译结束>


<原文开始>
				// System may be overloaded; sleep a bit
				// in the hopes it will recover.
<原文结束>

# <翻译开始>
// 系统可能过载；稍微睡一会儿，希望它能恢复。
// md5:58a2798075f6c4be
# <翻译结束>


<原文开始>
// Now test that the ticker stopped.
<原文结束>

# <翻译开始>
// 现在测试ticker是否已停止。. md5:ba5a168a74615cec
# <翻译结束>


<原文开始>
// Test passed, so all done.
<原文结束>

# <翻译开始>
// 测试通过，所以大功告成。. md5:898d4f57bd6d9649
# <翻译结束>


<原文开始>
// Test that a bug tearing down a ticker has been fixed. This routine should not deadlock.
<原文结束>

# <翻译开始>
// 测试修复了定时器拆卸时的bug。这个程序不应该出现死锁。. md5:306140edccec42f2
# <翻译结束>


<原文开始>
// Test the Tick convenience wrapper.
<原文结束>

# <翻译开始>
// 测试Tick便利包装器。. md5:901afab2f3b2cca2
# <翻译结束>


<原文开始>
// Test that giving a negative duration returns nil.
<原文结束>

# <翻译开始>
// 测试给一个负的持续时间是否返回nil。. md5:a6ce42ca827ba1db
# <翻译结束>


<原文开始>
// Test that NewTicker panics when given a duration less than zero.
<原文结束>

# <翻译开始>
// 测试当给定的持续时间小于零时，NewTicker是否会发生恐慌。. md5:c56054c00884775b
# <翻译结束>


<原文开始>
// Test that Ticker.Reset panics when given a duration less than zero.
<原文结束>

# <翻译开始>
// 测试当向Ticker重置传递一个小于零的持续时间时，它是否会引发恐慌。. md5:609097bbc9af9981
# <翻译结束>

