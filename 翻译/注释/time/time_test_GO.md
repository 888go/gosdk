
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
// We should be in PST/PDT, but if the time zone files are missing we
// won't be. The purpose of this test is to at least explain why some of
// the subsequent tests fail.
<原文结束>

# <翻译开始>
// 我们应该在PST/PDT时区，但如果时区文件缺失，我们不会在。这个测试的目的是至少解释为什么后续的一些测试会失败。
// md5:a877153cf135d94d
# <翻译结束>


<原文开始>
// PST is 8 hours west, PDT is 7 hours west. We could use the name but it's not unique.
<原文结束>

# <翻译开始>
// PST位于西边8小时，PDT位于西边7小时。我们可以使用这个名字，但它并不唯一。. md5:cbe7a8e2be638df6
# <翻译结束>


<原文开始>
// parsedTime is the struct representing a parsed time value.
<原文结束>

# <翻译开始>
// parsedTime 是表示解析后时间值的结构体。. md5:a90be61d11060ff5
# <翻译结束>


<原文开始>
// seconds east of UTC, e.g. -7*60*60 for -0700
<原文结束>

# <翻译开始>
// UTC以东的秒数，例如-7*60*60表示-0700. md5:5eace65a71c2a861
# <翻译结束>


<原文开始>
// Check individual entries.
<原文结束>

# <翻译开始>
// 检查单个条目。. md5:c7203452782bccc9
# <翻译结束>


<原文开始>
// Try a reasonable date first, then the huge ones.
<原文结束>

# <翻译开始>
// 首先尝试一个合理的日期，然后尝试较大的日期。. md5:379431a5c1a862c0
# <翻译结束>


<原文开始>
	// Try a small date first, then the large ones. (The span is only a few hundred years
	// for nanoseconds in an int64.)
<原文结束>

# <翻译开始>
// 首先尝试一个小的日期，然后是大的日期。（对于int64中的纳秒，跨度只有几百年）
// md5:3360e6892f5fb9e4
# <翻译结束>


<原文开始>
// The time routines provide no way to get absolute time
// (seconds since zero), but we need it to compute the right
// answer for bizarre roundings like "to the nearest 3 ns".
// Compute as t - year1 = (t - 1970) + (1970 - 2001) + (2001 - 1).
// t - 1970 is returned by Unix and Nanosecond.
// 1970 - 2001 is -(31*365+8)*86400 = -978307200 seconds.
// 2001 - 1 is 2000*365.2425*86400 = 63113904000 seconds.
<原文结束>

# <翻译开始>
// 时间 routines（调度）无法获取绝对时间（自零点的秒数），但我们需要它来计算出像"四舍五入到最近的3 ns"这样的奇数结果。计算方法为：t - year1 = (t - 1970) + (1970 - 2001) + (2001 - 1)。
// t - 1970 由 Unix 和 Nanosecond 返回。
// 1970 - 2001 是 -(31 * 365 + 8) * 86400 = -978307200 秒。
// 2001 - 1 是 2000 * 365.2425 * 86400 = 63113904000 秒。
// md5:a5c6dba33aaa4bcd
# <翻译结束>


<原文开始>
// abs returns the absolute time stored in t, as seconds and nanoseconds.
<原文结束>

# <翻译开始>
// abs 返回存储在 t 中的绝对时间，以秒和纳秒为单位。. md5:9b3a61e1f35f15df
# <翻译结束>


<原文开始>
// absString returns abs as a decimal string.
<原文结束>

# <翻译开始>
// absString 将 abs 转换为十进制字符串。. md5:f4ecea5809009a31
# <翻译结束>


<原文开始>
// 5.8*d rounds to 6*d, but .8*d+.8*d < 0 < d
<原文结束>

# <翻译开始>
// 5.8 * d 向上舍入到 6 * d，但是 .8 * d + .8 * d < 0 < d. md5:17cdc3615444f015
# <翻译结束>


<原文开始>
// Compute bt = absolute nanoseconds.
<原文结束>

# <翻译开始>
// 计算bt为绝对纳秒。. md5:bb29a5c2c2e5fcb0
# <翻译结束>


<原文开始>
// Compute quotient and remainder mod d.
<原文结束>

# <翻译开始>
// 计算商和模d的余数。. md5:1ad7456726aa86c2
# <翻译结束>


<原文开始>
		// To truncate, subtract remainder.
		// br is < d, so it fits in an int64.
<原文结束>

# <翻译开始>
// 通过取余来截断。
// br 小于 d，所以它可以容纳在 int64 中。
// md5:44c8c5d861d5fb5b
# <翻译结束>


<原文开始>
// Check that time.Truncate works.
<原文结束>

# <翻译开始>
// 检查time.Truncate是否正常工作。. md5:51e423e287005315
# <翻译结束>


<原文开始>
		// To round, add d back if remainder r > d/2 or r == exactly d/2.
		// The commented out code would round half to even instead of up,
		// but that makes it time-zone dependent, which is a bit strange.
<原文结束>

# <翻译开始>
// 如果余数 r 大于 d/2 或者等于 d/2，为了四舍五入，我们将 d 添加回去。
// 注释掉的代码会将半数四舍五入到偶数，而不是向上取整。但这会使它依赖时区，这有点奇怪。
// md5:5621afa60a7d98b6
# <翻译结束>


<原文开始>
// Check that time.Round works.
<原文结束>

# <翻译开始>
// 检查time.Round函数是否正常工作。. md5:e5b2f97fd408128f
# <翻译结束>


<原文开始>
// randomly generated test cases
<原文结束>

# <翻译开始>
// 随机生成的测试用例. md5:e6955df765195da3
# <翻译结束>


<原文开始>
		// Make room for unix ? internal conversion.
		// We don't care about behavior too close to ± 2^63 Unix seconds.
		// It is full of wraparounds but will never happen in a reasonable program.
		// (Or maybe not? See go.dev/issue/20678. In any event, they're not handled today.)
<原文结束>

# <翻译开始>
// 为Unix？内部转换腾出空间。
// 我们不关心距离±2^63秒的Unix时间太近的行为。
// 它充满了循环，但在合理的程序中永远不会发生。
// （或者不会？参见go.dev/issue/20678。无论如何，目前它们没有被处理。）
// md5:993ba03fb528e69f
# <翻译结束>


<原文开始>
// Selected dates and corner cases
<原文结束>

# <翻译开始>
// 选择的日期和边界情况. md5:4202016edce8d9ff
# <翻译结束>


<原文开始>
// The only real invariant: Jan 04 is in week 1
<原文结束>

# <翻译开始>
// 唯一真正的不变性：1月4日是在第1周内。. md5:278b4045520e4ef0
# <翻译结束>


<原文开始>
// Test YearDay in several different scenarios
// and corner cases
<原文结束>

# <翻译开始>
// 在多种不同场景和边界情况下测试YearDay
// md5:467db368371fe934
# <翻译结束>


<原文开始>
// Looks like leap-year (but isn't) tests
<原文结束>

# <翻译开始>
// 看起来像是闰年（但实际上不是）的测试. md5:ce6d977814d12f28
# <翻译结束>


<原文开始>
// Year one tests (non-leap)
<原文结束>

# <翻译开始>
// 第一年的测试（非闰年）. md5:4d82de1d5378f886
# <翻译结束>


<原文开始>
// Year minus one tests (non-leap)
<原文结束>

# <翻译开始>
// 减去一年的测试（非闰年）. md5:9b77b30db967b744
# <翻译结束>


<原文开始>
// 400 BC tests (leap-year)
<原文结束>

# <翻译开始>
// 400年前的BC测试（闰年）. md5:e252a45b423cb8ca
# <翻译结束>


<原文开始>
// Gregorian calendar change (no effect)
<原文结束>

# <翻译开始>
// 公历变更（无影响）. md5:93bdb417c060cda9
# <翻译结束>


<原文开始>
// Check to see if YearDay is location sensitive
<原文结束>

# <翻译开始>
// 检查YearDay是否为位置敏感的. md5:5b8dab4efd11f82d
# <翻译结束>


<原文开始>
// Many names for Fri Nov 18 7:56:35 PST 2011
<原文结束>

# <翻译开始>
// 2011年11月18日星期五7:56:35 PST的多种名称. md5:8dd2ad304f033b90
# <翻译结束>


<原文开始>
// (Jan-2) 18 7:56:35 2012
<原文结束>

# <翻译开始>
// (一月-2) 18 7:56:35 2012. md5:6af96243b33568b3
# <翻译结束>


<原文开始>
// (Dec+11) 18 7:56:35 2010
<原文结束>

# <翻译开始>
//(Dec+11) 18 07:56:35 2010
// 
// 这段Go语言注释的中文翻译是："（Dec+11）2010年12月11日 18:07:56"。这是时间戳的格式，Dec+11表示月份（12月后加11，即1月），18:07:56是24小时制的时间，2010则是年份。. md5:edaae9041aa5b491
# <翻译结束>


<原文开始>
// Several ways of getting from
// Fri Nov 18 7:56:35 PST 2011
// to
// Thu Mar 19 7:56:35 PST 2016
<原文结束>

# <翻译开始>
// 从
// 2011年11月18日星期五，太平洋标准时间7:56:35
// 到
// 2016年3月19日星期四，太平洋标准时间7:56:35
// 的多种方式
// md5:85909e1e910bda87
# <翻译结束>


<原文开始>
// January, first month, 31 days
<原文结束>

# <翻译开始>
// 一月，第一个月，31天. md5:3158b6d74f31a42d
# <翻译结束>


<原文开始>
// February, non-leap year, 28 days
<原文结束>

# <翻译开始>
// 二月，非闰年，28天. md5:e49798a45ef3f35c
# <翻译结束>


<原文开始>
// February, leap year, 29 days
<原文结束>

# <翻译开始>
// 二月，闰年，29天. md5:5026573ed7f29c97
# <翻译结束>


<原文开始>
// December, last month, 31 days
<原文结束>

# <翻译开始>
// 上个月，12月，有31天. md5:f795504febd08823
# <翻译结束>


<原文开始>
	// The daysIn function is not exported.
	// Test the daysIn function via the `var DaysIn = daysIn`
	// statement in the internal_test.go file.
<原文结束>

# <翻译开始>
// daysIn 函数没有被导出。
// 通过在 internal_test.go 文件中的 `var DaysIn = daysIn` 语句来测试 daysIn 函数。
// md5:a45861afd1e9e8fa
# <翻译结束>


<原文开始>
	// Add an amount to the current time to round it up to the next exact second.
	// This test checks that the nsec field still lies within the range [0, 999999999].
<原文结束>

# <翻译开始>
// 向当前时间添加一个量，使其向上舍入到下一个精确的秒。
// 此测试检查nsec字段是否仍然在范围[0, 999999999]内。
// md5:c18836b1f3d78e03
# <翻译结束>


<原文开始>
// Time.sec: 0x0123456789ABCDEF
<原文结束>

# <翻译开始>
// Time.sec: 0x0123456789ABCDEF
// 时间.sec: 0x0123456789ABCDEF（十六进制表示的秒值）. md5:9b45ef185347f6e0
# <翻译结束>


<原文开始>
// more than 9 digits after decimal point, see https://golang.org/issue/6617
<原文结束>

# <翻译开始>
// 小数点后有超过9位数字，参见 https://golang.org/issue/6617. md5:d67cd94bd42d243f
# <翻译结束>


<原文开始>
// 9007199254740993 = 1<<53+1 cannot be stored precisely in a float64
<原文结束>

# <翻译开始>
// 9007199254740993 = 1 左移 53 位再加上 1，这个数值无法精确地存储在一个 float64 类型中。. md5:c17cfac2fa94c135
# <翻译结束>


<原文开始>
// largest duration that can be represented by int64 in nanoseconds
<原文结束>

# <翻译开始>
// 可以用int64表示的最大纳秒时间间隔. md5:0e44b710b3398879
# <翻译结束>


<原文开始>
// largest negative round trip value, see https://golang.org/issue/48629
<原文结束>

# <翻译开始>
// 最大的负回环值，参见：https://golang.org/issue/48629. md5:5b191ebe8da8d694
# <翻译结束>


<原文开始>
// huge string; issue 15011.
<原文结束>

# <翻译开始>
// 大字符串；问题15011。. md5:07e8867ae8029c1f
# <翻译结束>


<原文开始>
// This value tests the first overflow check in leadingFraction.
<原文结束>

# <翻译开始>
// 此值用于测试leadingFraction中第一个溢出检查。. md5:0e26e39dfc483cd4
# <翻译结束>


<原文开始>
// https://golang.org/issue/48629
<原文结束>

# <翻译开始>
// https://golang.org/issue/48629 该问题链接指向Go语言官方仓库的一个问题讨论或报告。. md5:0add436cff6ef0f2
# <翻译结束>


<原文开始>
		// Resolutions finer than milliseconds will result in
		// imprecise round-trips.
<原文结束>

# <翻译开始>
// 毫秒级以下的分辨率会导致不精确的往返。
// md5:77015c25f30f6270
# <翻译结束>


<原文开始>
// reset the Once to trigger the race
<原文结束>

# <翻译开始>
// 将Once重置以触发竞争条件. md5:f7940df75a562e17
# <翻译结束>


<原文开始>
// Back to Los Angeles for subsequent tests:
<原文结束>

# <翻译开始>
// 返回洛杉矶进行后续测试：. md5:97e94316d288fb0e
# <翻译结束>


<原文开始>
// Issue 4064: handle locations without any zone transitions.
<原文结束>

# <翻译开始>
// 问题4064：处理没有时区转换的地理位置。. md5:39d7da51684886db
# <翻译结束>


<原文开始>
	// The tzdata name Etc/GMT+1 uses "east is negative",
	// but Go and most other systems use "east is positive".
	// So GMT+1 corresponds to -3600 in the Go zone, not +3600.
<原文结束>

# <翻译开始>
// Etc/GMT+1在tzdata中使用"东为负"的约定，但Go和其他大多数系统使用"东为正"。因此，GMT+1在Go时区中对应的是-3600秒，而不是+3600秒。
// md5:ba8d98a3734e5d88
# <翻译结束>


<原文开始>
	// The zone abbreviation is "-01" since tzdata-2016g, and "GMT+1"
	// on earlier versions; we accept both. (Issue #17276).
<原文结束>

# <翻译开始>
// 自tzdata-2016g以来，时区缩写为"-01"，早期版本为"GMT+1"；我们接受两者。（问题#17276）。
// md5:6d7d3da8d78f3359
# <翻译结束>


<原文开始>
// Using Equal since Add don't modify loc using "==" will cause a fail
<原文结束>

# <翻译开始>
// 使用Equal是因为Add操作不会修改loc，如果使用"=="比较会导致错误. md5:3622609b0185b9c0
# <翻译结束>


<原文开始>
//Original caus for this test case bug 15852
<原文结束>

# <翻译开始>
//此测试用例的原始原因：bug 15852. md5:a83da26ab0c43f24
# <翻译结束>


<原文开始>
	// Verify that all of Time's methods behave identically if loc is set to
	// nil or UTC.
<原文结束>

# <翻译开始>
// 验证如果loc被设置为nil或UTC，Time的所有方法行为是否相同。
// md5:ffb3babe06d51eb3
# <翻译结束>


<原文开始>
	// Like BenchmarkFormat, but easier, because the time zone
	// lookup cache is optimized for the present.
<原文结束>

# <翻译开始>
// 类似 BenchmarkFormat，但是更简单，因为时区查找缓存是针对当前优化的。
// md5:f06dd1b142ed8305
# <翻译结束>


<原文开始>
// short enough to be stack allocated
<原文结束>

# <翻译开始>
// 足够短以在栈上分配. md5:195247dda0133cd7
# <翻译结束>


<原文开始>
// Issue 17720: Zero value of time.Month fails to print
<原文结束>

# <翻译开始>
// 问题17720：time.Month的零值无法正常打印. md5:db136c905142b832
# <翻译结束>


<原文开始>
// Issue 24692: Out of range weekday panics
<原文结束>

# <翻译开始>
// 问题 24692：超出范围的工作日恐慌. md5:20b30cd630c70a9f
# <翻译结束>


<原文开始>
// Issue 25686: hard crash on concurrent timer access.
// Issue 37400: panic with "racy use of timers"
// This test deliberately invokes a race condition.
// We are testing that we don't crash with "fatal error: panic holding locks",
// and that we also don't panic.
<原文结束>

# <翻译开始>
// Issue 25686：并发访问定时器时的硬崩溃。
// Issue 37400：使用定时器时出现“竞态条件”导致panic。
// 此测试故意触发竞态条件。
// 我们正在测试在持有锁的情况下不出现“致命错误：panic”，并且也不会panic。
// md5:62f83f45c1a23d84
# <翻译结束>


<原文开始>
// Issue 37400: panic with "racy use of timers".
<原文结束>

# <翻译开始>
// Issue 37400: 因“计时器的竞态使用”导致的 panic。. md5:256cc7c7f68e7bb5
# <翻译结束>


<原文开始>
// Test it with positive delta.
<原文结束>

# <翻译开始>
// 使用正向的delta进行测试。. md5:4de1a0f28de88248
# <翻译结束>


<原文开始>
// Test it with negative delta.
<原文结束>

# <翻译开始>
// 使用负的delta值进行测试。. md5:e045f2e98546ccfc
# <翻译结束>


<原文开始>
// Issue 49284: time: ParseInLocation incorrectly because of Daylight Saving Time
<原文结束>

# <翻译开始>
// 问题 49284：由于夏令时，time.ParseInLocation 函数解析错误. md5:485ad6976ad20904
# <翻译结束>


<原文开始>
		// 14 Apr 1991 - Daylight Saving Time Started
		// When time of "Asia/Shanghai" was about to reach
		// Sunday, 14 April 1991, 02:00:00 clocks were turned forward 1 hour to
		// Sunday, 14 April 1991, 03:00:00 local daylight time instead.
		// The UTC time was 13 April 1991, 18:00:00
<原文结束>

# <翻译开始>
// 1991年4月14日 - 夏令时开始
// 当“亚洲/上海”时间即将到达
// 1991年4月14日，星期日，凌晨02:00:00时，时钟向前调整了1小时，
// 变为1991年4月14日，星期日，凌晨03:00:00的当地夏令时。
// 这时的UTC时间是1991年4月13日，18:00:00。
// md5:154cc70a3ad2ed2f
# <翻译结束>


<原文开始>
		// 15 Sep 1991 - Daylight Saving Time Ended
		// When local daylight time of "Asia/Shanghai" was about to reach
		// Sunday, 15 September 1991, 02:00:00 clocks were turned backward 1 hour to
		// Sunday, 15 September 1991, 01:00:00 local standard time instead.
		// The UTC time was 14 September 1991, 17:00:00
<原文结束>

# <翻译开始>
// 1991年9月15日 - 夏令时结束
// 当"Asia/Shanghai"的本地夏令时即将达到
// 1991年9月15日，02:00:00 时，时钟回拨一小时至
// 1991年9月15日，01:00:00 本地标准时间。
// UTC时间为1991年9月14日，17:00:00。
// md5:200f61ea8b858f25
# <翻译结束>


<原文开始>
// The ZoneBounds of a UTC location would just return two zero Time.
<原文结束>

# <翻译开始>
// UTC位置的区划边界将只返回两个零Time。. md5:3223c57808e355f0
# <翻译结束>


<原文开始>
	// If the zone begins at the beginning of time, start will be returned as a zero Time.
	// Use math.MinInt32 to avoid overflow of int arguments on 32-bit systems.
<原文结束>

# <翻译开始>
// 如果时区从时间的开始处开始，start将作为零Time返回。
// 使用math.MinInt32可以避免在32位系统上int参数溢出的问题。
// md5:b75e6e21de334f47
# <翻译结束>


<原文开始>
	// If the zone goes on forever, end will be returned as a zero Time.
	// Use math.MaxInt32 to avoid overflow of int arguments on 32-bit systems.
<原文结束>

# <翻译开始>
// 如果时区无限延续，end将被返回为一个零时间值。
// 在32位系统上，请使用math.MaxInt32以避免int型参数溢出。
// md5:f77aba81d8f5afa3
# <翻译结束>


<原文开始>
// Check some real-world cases to make sure we're getting the right bounds.
<原文结束>

# <翻译开始>
// 检查一些现实世界的情况，确保我们得到正确的边界。. md5:3e7391ac63d3b374
# <翻译结束>


<原文开始>
// The ZoneBounds of "Asia/Shanghai" Daylight Saving Time
<原文结束>

# <翻译开始>
// "Asia/Shanghai"夏令时的区域边界. md5:3ec7dffa23c5942e
# <翻译结束>


<原文开始>
// The ZoneBounds of a "Asia/Shanghai" after the last transition (Standard Time)
<原文结束>

# <翻译开始>
// "Asia/Shanghai"时区在最后一次（标准时间）转换之后的边界范围. md5:27f997ff7dc24c8c
# <翻译结束>


<原文开始>
		// The ZoneBounds of a local time would return two local Time.
		// Note: We preloaded "America/Los_Angeles" as time.Local for testing
<原文结束>

# <翻译开始>
// 获取一个本地时间的ZoneBounds将返回两个本地时间。
// 注意：为了测试，我们预先加载了"America/Los_Angeles"作为time.Local。
// md5:8bcb336908b3115a
# <翻译结束>

