
<原文开始>
// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// 版权所有 2011 Go 语言作者。保留所有权利。
// 使用此源代码须遵循 BSD 风格的许可协议，
// 许可协议具体内容可在 LICENSE 文件中找到。
// md5:b5bcbe8a534f6077
# <翻译结束>


<原文开始>
// The package also accepts the incorrect but common prefix u for micro.
<原文结束>

# <翻译开始>
// 此包也接受常见的但不正确的前缀 u 代表微（micro）。. md5:3e19ea04de8a976c
# <翻译结束>


<原文开始>
// Parse a time value from a string in the standard Unix format.
<原文结束>

# <翻译开始>
// 从标准Unix格式的字符串中解析时间值。. md5:f3f8abffc56da1da
# <翻译结束>


<原文开始>
// Always check errors even if they should not happen.
<原文结束>

# <翻译开始>
// 即使不应该发生错误，也始终检查错误。. md5:1ee31877d1f35f48
# <翻译结束>


<原文开始>
// time.Time's Stringer method is useful without any format.
<原文结束>

# <翻译开始>
// time.Time 类型的 Stringer 接口在没有指定格式的情况下也很有用。. md5:d26240ee01fcdb99
# <翻译结束>


<原文开始>
// Predefined constants in the package implement common layouts.
<原文结束>

# <翻译开始>
// 包中预定义的常量实现了常见的布局。. md5:efc3cd559fe4fd78
# <翻译结束>


<原文开始>
// The time zone attached to the time value affects its output.
<原文结束>

# <翻译开始>
// 与时间值关联的时区会影响其输出。. md5:e4befa4c6f328243
# <翻译结束>


<原文开始>
	// The rest of this function demonstrates the properties of the
	// layout string used in the format.
<原文结束>

# <翻译开始>
// 函数的剩余部分演示了在格式中使用的布局字符串的属性。
// md5:18b912fdc17d3847
# <翻译结束>


<原文开始>
	// The layout string used by the Parse function and Format method
	// shows by example how the reference time should be represented.
	// We stress that one must show how the reference time is formatted,
	// not a time of the user's choosing. Thus each layout string is a
	// representation of the time stamp,
	//	Jan 2 15:04:05 2006 MST
	// An easy way to remember this value is that it holds, when presented
	// in this order, the values (lined up with the elements above):
	//	  1 2  3  4  5    6  -7
	// There are some wrinkles illustrated below.
<原文结束>

# <翻译开始>
// Parse函数和Format方法使用的布局字符串示例展示了参考时间应该如何表示。
// 我们强调必须展示参考时间的格式，而不是用户选择的时间。因此，每个布局字符串都是时间戳的表示，
// 例如：Jan 2 15:04:05 2006 MST
// 记住这个值的一个简单方法是，当按照这个顺序呈现时，它包含了（与上面的元素对齐）：
//	1 2  3  4  5    6  -7
// 下面有一些说明示例。
// md5:5c1ab7903b161365
# <翻译结束>


<原文开始>
	// Most uses of Format and Parse use constant layout strings such as
	// the ones defined in this package, but the interface is flexible,
	// as these examples show.
<原文结束>

# <翻译开始>
// 大多数情况下，Format和Parse的使用都采用诸如本包中定义的常量布局字符串，但该接口是灵活的，如下例所示。
// md5:f3dbf95fe362fed3
# <翻译结束>


<原文开始>
// Define a helper function to make the examples' output look nice.
<原文结束>

# <翻译开始>
// 定义一个辅助函数，使示例的输出看起来更整洁。. md5:744e830c88834a10
# <翻译结束>


<原文开始>
// Print a header in our output.
<原文结束>

# <翻译开始>
// 在我们的输出中打印一个标题。. md5:8db341353a187ad8
# <翻译结束>


<原文开始>
// Simple starter examples.
<原文结束>

# <翻译开始>
// 简单的启动示例。. md5:877bb2f43f13947f
# <翻译结束>


<原文开始>
	// The hour of the reference time is 15, or 3PM. The layout can express
	// it either way, and since our value is the morning we should see it as
	// an AM time. We show both in one format string. Lower case too.
<原文结束>

# <翻译开始>
// 参考时间的小时为15，即下午3点。布局可以两种方式表示它，由于我们的值是在上午，所以我们应该看到它作为一个上午的时间。我们在一个格式字符串中展示两种表示。同时使用小写字母。
// md5:4eee1220f173861d
# <翻译结束>


<原文开始>
	// When parsing, if the seconds value is followed by a decimal point
	// and some digits, that is taken as a fraction of a second even if
	// the layout string does not represent the fractional second.
	// Here we add a fractional second to our time value used above.
<原文结束>

# <翻译开始>
// 在解析时，如果秒值后面跟着小数点和一些数字，
// 即使布局字符串不表示分数秒，也会将其视为秒的分数。
// 这里我们向上面使用的的时间值添加了一个分数秒。
// md5:f479515c456c93fc
# <翻译结束>


<原文开始>
	// It does not appear in the output if the layout string does not contain
	// a representation of the fractional second.
<原文结束>

# <翻译开始>
// 如果布局字符串不包含小数秒的表示，则不会出现在输出中。
// md5:498a488987974671
# <翻译结束>


<原文开始>
	// Fractional seconds can be printed by adding a run of 0s or 9s after
	// a decimal point in the seconds value in the layout string.
	// If the layout digits are 0s, the fractional second is of the specified
	// width. Note that the output has a trailing zero.
<原文结束>

# <翻译开始>
// 在格式字符串的秒值中，通过在小数点后添加一串0或9，可以打印出小数部分秒。如果格式字符串中的数字是0，那么小数秒的宽度就是指定的。需要注意的是，输出将以0结尾。
// md5:3063cbb9352ae856
# <翻译结束>


<原文开始>
// If the fraction in the layout is 9s, trailing zeros are dropped.
<原文结束>

# <翻译开始>
// 如果布局中的分数部分为9s，将删除尾随的零。. md5:ac61ce54f30db97d
# <翻译结束>


<原文开始>
// The predefined constant Unix uses an underscore to pad the day.
<原文结束>

# <翻译开始>
// 预定义的常量Unix使用下划线来填充日期。. md5:d9f4dbdd3de4ebe7
# <翻译结束>


<原文开始>
	// For fixed-width printing of values, such as the date, that may be one or
	// two characters (7 vs. 07), use an _ instead of a space in the layout string.
	// Here we print just the day, which is 2 in our layout string and 7 in our
	// value.
<原文结束>

# <翻译开始>
// 对于可能为一或两个字符（例如日期，7 vs. 07）的固定宽度打印值，使用下划线_代替空格在格式字符串中。这里我们只打印日期，格式字符串中的数字是2，而实际值是7。
// md5:a1ebfb451e3503bb
# <翻译结束>


<原文开始>
// An underscore represents a space pad, if the date only has one digit.
<原文结束>

# <翻译开始>
// 下划线代表一个空格垫，如果日期只有一位数字。. md5:49def7f2b857b6f0
# <翻译结束>


<原文开始>
// A "0" indicates zero padding for single-digit values.
<原文结束>

# <翻译开始>
// "0" 表示为单数字值填充零。. md5:014b4fd453e8e5b1
# <翻译结束>


<原文开始>
	// If the value is already the right width, padding is not used.
	// For instance, the second (05 in the reference time) in our value is 39,
	// so it doesn't need padding, but the minutes (04, 06) does.
<原文结束>

# <翻译开始>
// 如果值已经是正确的宽度，不需要填充。
// 例如，我们值中的第二部分（参考时间中的05）为39，因此不需要填充，
// 但分钟（04, 06）则需要。
// md5:f641faf4ba09711a
# <翻译结束>


<原文开始>
	// See the example for Time.Format for a thorough description of how
	// to define the layout string to parse a time.Time value; Parse and
	// Format use the same model to describe their input and output.
<原文结束>

# <翻译开始>
// 参见Time.Format示例，详细了解如何定义解析时间.Time值的布局字符串；Parse和Format使用相同的模型来描述其输入和输出。
// md5:9d063b6b107bc3c8
# <翻译结束>


<原文开始>
	// longForm shows by example how the reference time would be represented in
	// the desired layout.
<原文结束>

# <翻译开始>
// longForm通过示例展示参考时间在期望的格式中是如何表示的。
// md5:761675d50d7047bd
# <翻译结束>


<原文开始>
	// shortForm is another way the reference time would be represented
	// in the desired layout; it has no time zone present.
	// Note: without explicit zone, returns time in UTC.
<原文结束>

# <翻译开始>
// shortForm 是引用时间在期望布局下的另一种表示方式，
// 其中不包含时区信息。
// 注意：如果没有明确指定时区，默认返回的是UTC时间。
// md5:82da07f19c343810
# <翻译结束>


<原文开始>
	// Some valid layouts are invalid time values, due to format specifiers
	// such as _ for space padding and Z for zone information.
	// For example the RFC3339 layout 2006-01-02T15:04:05Z07:00
	// contains both Z and a time zone offset in order to handle both valid options:
	// 2006-01-02T15:04:05Z
	// 2006-01-02T15:04:05+07:00
<原文结束>

# <翻译开始>
// 一些有效的布局可能对应无效的时间值，因为格式规范如_用于空格填充，Z用于时区信息。
// 例如，RFC3339布局"2006-01-02T15:04:05Z07:00"包含了Z和时区偏移，以处理两种有效的情况：
// 2006-01-02T15:04:05Z
// 2006-01-02T15:04:05+07:00
// md5:a0f0fb87b351aede
# <翻译结束>


<原文开始>
// Returns an error as the layout is not a valid time value
<原文结束>

# <翻译开始>
// 如果布局不是有效的时间值，返回一个错误. md5:487c96021bec619d
# <翻译结束>


<原文开始>
// This will look for the name CEST in the Europe/Berlin time zone.
<原文结束>

# <翻译开始>
// 这将查找位于Europe/Berlin时区的"CEST"。. md5:fa2a6f1d39b85d79
# <翻译结束>


<原文开始>
// Note: without explicit zone, returns time in given location.
<原文结束>

# <翻译开始>
// 注意：如果没有明确指定时区，则按照给定的地点返回时间。. md5:b453850c1e5f79e8
# <翻译结束>


<原文开始>
// 1 billion seconds of Unix, three ways.
<原文结束>

# <翻译开始>
// 以三种方式表示10亿秒的Unix时间。. md5:f6f77989ce38890c
# <翻译结束>


<原文开始>
// 2e9 seconds - 1e18 nanoseconds
<原文结束>

# <翻译开始>
// 2e9 秒 - 1e18 纳秒. md5:bf13df9731ad7475
# <翻译结束>


<原文开始>
// To round to the last midnight in the local timezone, create a new Date.
<原文结束>

# <翻译开始>
// 要四舍五入到本地时区的前一天午夜，创建一个新日期。. md5:f23f2d701f87ef7e
# <翻译结束>


<原文开始>
// China doesn't have daylight saving. It uses a fixed 8 hour offset from UTC.
<原文结束>

# <翻译开始>
// 中国没有夏令时。它使用相对于UTC固定的8小时偏移。. md5:6aab1ae63f661255
# <翻译结束>


<原文开始>
	// If the system has a timezone database present, it's possible to load a location
	// from that, e.g.:
	//    newYork, err := time.LoadLocation("America/New_York")
<原文结束>

# <翻译开始>
// 如果系统中存在时区数据库，那么可以从中加载一个时区，例如：
//    newYork, err := time.LoadLocation("America/New_York")
// md5:771c6af761901cc5
# <翻译结束>


<原文开始>
// Creating a time requires a location. Common locations are time.Local and time.UTC.
<原文结束>

# <翻译开始>
// 创建一个时间点需要一个时区。常见的时区有time.Local和time.UTC。. md5:8603190e4b9df7b8
# <翻译结束>


<原文开始>
	// Although the UTC clock time is 1200 and the Beijing clock time is 2000, Beijing is
	// 8 hours ahead so the two dates actually represent the same instant.
<原文结束>

# <翻译开始>
// 虽然 UTC 时间是 1200，而北京时间是 2000，但由于北京比 UTC 领先 8 个小时，所以这两个日期实际上代表的是同一时刻。
// md5:adb00e89cda10858
# <翻译结束>


<原文开始>
	// Unlike the equal operator, Equal is aware that d1 and d2 are the
	// same instant but in different time zones.
<原文结束>

# <翻译开始>
// 与相等运算符不同，Equal 方法能够识别 d1 和 d2 是同一时间点，只是处于不同的时区。
// md5:9b87f1e5ff088a68
# <翻译结束>

