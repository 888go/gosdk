
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
	// Most bugs in Parse or Format boil down to problems with
	// the exact detection of format chunk boundaries in the
	// helper function nextStdChunk (here called as NextStdChunk).
	// This test checks nextStdChunk's behavior directly,
	// instead of needing to test it only indirectly through Parse/Format.
<原文结束>

# <翻译开始>
	// 大多数在Parse或Format中的错误归结于辅助函数nextStdChunk（这里称为NextStdChunk）中格式块边界的精确检测问题。这个测试直接检查nextStdChunk的行为，而不需要通过Parse/Format间接进行测试。
	// md5:dc0281b4ea39a780
# <翻译结束>


<原文开始>
	// markChunks returns format with each detected
	// 'format chunk' parenthesized.
	// For example showChunks("2006-01-02") == "(2006)-(01)-(02)".
<原文结束>

# <翻译开始>
	// markChunks 函数返回每个检测到的 "格式块" 都被括起来的格式。
	// 例如，showChunks("2006-01-02") == "(2006)-(01)-(02)"。
	// md5:b527b4cab40a3b2b
# <翻译结束>


<原文开始>
		// Note that NextStdChunk and StdChunkNames
		// are not part of time's public API.
		// They are exported in export_test for this test.
<原文结束>

# <翻译开始>
		// 请注意，NextStdChunk和StdChunkNames
		// 并不属于time包的公共API。
		// 它们仅在此测试中通过export_test被导出。
		// md5:48587daf5dea2dc9
# <翻译结束>


<原文开始>
		// marked is an expected output from markChunks.
		// If we delete the parens and pass it through markChunks,
		// we should get the original back.
<原文结束>

# <翻译开始>
		// marked 是 markChunks 函数的预期输出。
		// 如果我们删除括号并将它通过 markChunks 函数处理，
		// 我们应该能获得原始的输入返回。
		// md5:335b1d648fb01e33
# <翻译结束>


<原文开始>
// Three-letter months and days must not be followed by lower-case letter.
<原文结束>

# <翻译开始>
// 三字母的月份和日子后面不能跟着小写字母。. md5:cc09856776348922
# <翻译结束>


<原文开始>
// Time stamps, Fractional seconds.
<原文结束>

# <翻译开始>
// 时间戳，小数秒。. md5:f94b2b512d5c574e
# <翻译结束>


<原文开始>
// The numeric time represents Thu Feb  4 21:00:57.012345600 PST 2009
<原文结束>

# <翻译开始>
// 数值时间表示的是2009年2月4日星期四，太平洋标准时间（PST）下的21:00:57.012345600。. md5:f8bd7a8cb5126dc5
# <翻译结束>


<原文开始>
			// The 4 in %04d counts the - sign, so print -y instead
			// and introduce our own - sign.
<原文结束>

# <翻译开始>
			// %04d 中的 4 包含了负号，因此打印 -y 并自定义负号。
			// md5:68725574231a99d7
# <翻译结束>


<原文开始>
// sign of year, -1 indicates the year is not present in the format
<原文结束>

# <翻译开始>
// 年份的符号，-1表示年份在该格式中不存在. md5:da1830c86f6953a1
# <翻译结束>


<原文开始>
// number of digits of fractional second
<原文结束>

# <翻译开始>
// 小数部分的位数. md5:2cc296fbcae0d331
# <翻译结束>


<原文开始>
// Optional fractional seconds.
<原文结束>

# <翻译开始>
// 可选的小数秒部分。. md5:00c9ac781457753e
# <翻译结束>


<原文开始>
// Amount of white space should not matter.
<原文结束>

# <翻译开始>
// 空格的数量不应有所影响。. md5:a18d704701640125
# <翻译结束>


<原文开始>
// Leading zeros in other places should not be taken as fractional seconds.
<原文结束>

# <翻译开始>
// 其他地方的零不应被视为小数秒。. md5:55a24aa015ab5a17
# <翻译结束>


<原文开始>
// Month and day names only match when not followed by a lower-case letter.
<原文结束>

# <翻译开始>
// 月份和星期的名称只在不跟一个小写字母的情况下匹配。. md5:7e2aa3dab305db43
# <翻译结束>


<原文开始>
	// Accept any number of fractional second digits (including none) for .999...
	// In Go 1, .999... was completely ignored in the format, meaning the first two
	// cases would succeed, but the next four would not. Go 1.1 accepts all six.
	// decimal "." separator.
<原文结束>

# <翻译开始>
	// 接受任意数量的小数秒位数（包括没有小数位），如.999...
	// 在 Go 1 中，.999... 在格式中完全被忽略，这意味着前两个情况会成功，但接下来的四个情况不会。
	// Go 1.1 则接受所有这六种情况。
	// 小数点作为分隔符。
	// md5:2f78deb3265e8291
# <翻译结束>


<原文开始>
// TestParseInLocation checks that the Parse and ParseInLocation
// functions do not get confused by the fact that AST (Arabia Standard
// Time) and AST (Atlantic Standard Time) are different time zones,
// even though they have the same abbreviation.
//
// ICANN has been slowly phasing out invented abbreviation in favor of
// numeric time zones (for example, the Asia/Baghdad time zone
// abbreviation got changed from AST to +03 in the 2017a tzdata
// release); but we still want to make sure that the time package does
// not get confused on systems with slightly older tzdata packages.
<原文结束>

# <翻译开始>
// TestParseInLocation 检查 Parse 和 ParseInLocation 函数不会因AST（阿拉伯标准时间）和AST（大西洋标准时间）是不同的时区，但具有相同的缩写而产生混淆。
//
// ICANN 已经逐渐淘汰了自创的缩写，转而使用数字时区（例如，在2017a版tzdata发布中，Asia/Baghdad时区的缩写从AST更改为+03）；但我们仍然要确保在使用稍旧版本tzdata包的系统上，time包不会产生混淆。
// md5:98c0bc82994ed15e
# <翻译结束>


<原文开始>
	// A zero offset means that ParseInLocation did not recognize the
	// 'AST' abbreviation as matching the current location (Baghdad,
	// where we'd expect a +03 hrs offset); likely because we're using
	// a recent tzdata release (2017a or newer).
	// If it happens, skip the Baghdad test.
<原文结束>

# <翻译开始>
	// 一个零偏移表示ParseInLocation没有识别'AST'缩写与当前位置（巴格达，我们期望+03小时偏移）匹配；可能是因为我们使用了较新的tzdata发布（2017a或更高版本）。如果发生这种情况，跳过巴格达测试。
	// md5:faba138ea75723dc
# <翻译结束>


<原文开始>
	// In this case 'AST' means 'Atlantic Standard Time', and we
	// expect the abbreviation to correctly match the american
	// location.
<原文结束>

# <翻译开始>
	// 在这种情况下，'AST'表示'大西洋标准时间'，我们期望缩写能够正确匹配美国的地理位置。
	// md5:d4fb808a8e9139d9
# <翻译结束>


<原文开始>
// Ignore the time zone in the test. If it parses, it'll be OK.
<原文结束>

# <翻译开始>
// 在测试中忽略时区。如果能够解析，那就没问题。. md5:bf2d9ad29c726751
# <翻译结束>


<原文开始>
// Problematic time zone format needs special tests.
<原文结束>

# <翻译开始>
// 有问题的时区格式需要特别的测试。. md5:c661a6b690158d00
# <翻译结束>


<原文开始>
// The time should be Thu Feb  4 21:00:57 PST 2010
<原文结束>

# <翻译开始>
// 时间应为2010年2月4日星期四，太平洋标准时间21:00:57. md5:8b02281dad0ed335
# <翻译结束>


<原文开始>
// Nanoseconds must be checked against the precision of the input.
<原文结束>

# <翻译开始>
// 必须检查纳秒是否超出输入的精度。. md5:0ad2fdf7ae989ccf
# <翻译结束>


<原文开始>
// Try a reasonable date first, then the huge ones.
<原文结束>

# <翻译开始>
// 首先尝试一个合理的日期，然后尝试较大的日期。. md5:379431a5c1a862c0
# <翻译结束>


<原文开始>
// four letters must end in T.
<原文结束>

# <翻译开始>
// 四个字母的单词必须以"T"结尾。. md5:cd05e32109e09796
# <翻译结束>


<原文开始>
// run of upper-case letters too long.
<原文结束>

# <翻译开始>
// 连续的大写字母过长。. md5:d02440bd8cae8bd3
# <翻译结束>


<原文开始>
// five letters must end in T.
<原文结束>

# <翻译开始>
// 五个字母必须以T结尾。. md5:46f84208446da6fc
# <翻译结束>


<原文开始>
// must appear within the error
<原文结束>

# <翻译开始>
// 必须出现在错误中. md5:aa43882c413d121f
# <翻译结束>


<原文开始>
// issue 4502. StampNano requires exactly 9 digits of precision.
<原文结束>

# <翻译开始>
// 问题4502。StampNano需要精确到9位数字。. md5:2dfabdb6247ef22f
# <翻译结束>


<原文开始>
// issue 4493. Helpful errors.
<原文结束>

# <翻译开始>
// issue 4493. 有用的错误信息。. md5:78fb3e19d3500b81
# <翻译结束>


<原文开始>
// invalid second followed by optional fractional seconds
<原文结束>

# <翻译开始>
// 非法的第二个部分，后面跟着可选的小数秒. md5:993f99cf63c16ad9
# <翻译结束>


<原文开始>
// invalid or mismatched day-of-year
<原文结束>

# <翻译开始>
// 无效或不匹配的年中天数. md5:bae5796035fb1b20
# <翻译结束>


<原文开始>
// Check that a time without a Zone still produces a (numeric) time zone
// when formatted with MST as a requested zone.
<原文结束>

# <翻译开始>
// 检查一个没有时区的时间在使用MST作为请求时区格式化时，仍然能生成（数字表示的）时区。
// md5:9de8eb9b9dd546d6
# <翻译结束>


<原文开始>
// uses MST as its time zone
<原文结束>

# <翻译开始>
// 使用MST作为其时区. md5:8743d7a99d9ce6e6
# <翻译结束>


<原文开始>
// should accept timezone offsets with seconds like: Zone America/New_York   -4:56:02 -      LMT     1883 Nov 18 12:03:58
<原文结束>

# <翻译开始>
// 应该接受秒级别的时区偏移，例如：时区 America/New_York -4:56:02 -    LMT    1883年11月18日12:03:58. md5:ed4076e1b8985c66
# <翻译结束>


<原文开始>
// 292277026304-08-26T15:42:51Z
<原文结束>

# <翻译开始>
// 292277026304年08月26日 15时42分51秒 Z（UTC时间）. md5:5b0e4967ddc39a23
# <翻译结束>


<原文开始>
// 2022-08-22T20:45:40.676836973Z
<原文结束>

# <翻译开始>
// 2022年8月22日20时45分40.676836973秒Z. md5:1fbc93ed7295c9d9
# <翻译结束>


<原文开始>
// 9999-12-31T23:59:59.999999999Z
<原文结束>

# <翻译开始>
// 1999年12月31日23时59分59秒999999999Z（这是ISO 8601日期时间格式，表示的时间是世界协调时间）. md5:81581273a6df41a8
# <翻译结束>


<原文开始>
// -292277022365-05-08T08:17:07Z
<原文结束>

# <翻译开始>
// -292277022365-05-08T08:17:07Z 
// 这是ISO 8601日期时间格式的注释，表示的是一个UTC（协调世界时）的时间戳。具体翻译如下：
// 
// - -292277022365：这是一个大数，代表自1970年1月1日（Unix纪元）以来的秒数，大约是公元前4713年。
// - -05：时区偏移，这里是格林威治标准时间（GMT）减去5小时，通常用于表示美国东部时间（EST）。
// - 05-08：月份和日期，即5月8日。
// - T08:17:07Z：时间，其中T是时间与日期的分隔符，08:17:07是小时、分钟和秒，Z表示UTC（Zulu时区，即零时区）。
// 
// 所以，这个注释是在说一个大约在公元前4713年5月8日早上8点17分07秒（美国东部时间）的UTC时间。. md5:570b58db9b49ec1f
# <翻译结束>


<原文开始>
// equalTime is like time.Time.Equal, but also compares the time zone.
<原文结束>

# <翻译开始>
// equalTime 类似于 time.Time.Equal，但还会比较时区。. md5:75caad6c5a63b0ff
# <翻译结束>


<原文开始>
// Parsing as RFC3339 or RFC3339Nano should be identical.
<原文结束>

# <翻译开始>
// 将时间解析为RFC3339或RFC3339Nano应该是相同的。. md5:337424a02fb1612e
# <翻译结束>


<原文开始>
			// TODO(https://go.dev/issue/54580):
			// Remove these checks after ParseAny rejects all invalid RFC 3339.
<原文结束>

# <翻译开始>
			// TODO(https:			//go.dev/issue/54580): 			// 在ParseAny拒绝所有无效的RFC 3339之后，删除这些检查。
			// md5:d137d8744b3d3843
# <翻译结束>


<原文开始>
// Customized parser should be identical to general parser.
<原文结束>

# <翻译开始>
// 定制解析器应与通用解析器完全相同。. md5:c589baa9a07544ae
# <翻译结束>

