package time //bm:时间类

import "time"

// Month 指定一年中的月份（一月 = 1，...） 翻译备注:这两个不翻译,常用引用了这个,会出错. md5:c1896873d737e487
type Month int //hm:月份 cz:type Month   //md5:2e86b21d4a1befae

// Weekday 指定了一周中的某一天（周日=0，...） 翻译备注:这两个不翻译,常用引用了这个,会出错. md5:e7c6df3685894251
type Weekday int //md5:d67f7c5c595dac46

// Duration 表示两个瞬间之间的时间差，用一个 int64 秒级别的纳秒计数表示。这个表示方式限制了可表示的最长持续时间大约为 290 年。
// md5:4c8441d90787436e
type Duration int64 //hm:时长 cz:type Duration   //md5:f81176da35f473b9

// Time 表示时间，具有纳秒精度。
//
// 使用时间的程序通常会将它们作为值存储和传递，而不是指针。也就是说，时间变量和结构体字段应为 time.Time 类型，而不是 *time.Time 类型。
//
// 时间值可以在多个goroutine中同时使用，除了 GobDecode、UnmarshalBinary、UnmarshalJSON 和 UnmarshalText 方法不是并发安全的。
//
// 可以使用 Before、After 和 Equal 方法比较时间点。Sub 方法可以减去两个时间点，得到一个 Duration。Add 方法可以添加一个 Time 和一个 Duration，得到一个新的 Time。
//
// 时间类型的零值表示 UTC 的公元 1 年 1 月 1 日 00:00:00.000000000。由于这个时间在实际应用中很少出现，IsZero 方法提供了一种简单的方法来检测未明确初始化的时间。
//
// 每个时间都有一个关联的 Location（时区）。Local、UTC 和 In 方法返回具有特定 Location 的 Time。通过这些方法更改 Time 值的 Location 不会改变它所代表的实际时刻，只会改变其解析时的时区。
//
// GobEncode、MarshalBinary、MarshalJSON 和 MarshalText 方法保存的时间值表示形式存储了 Time.Location 的偏移量，但不包含时区名称，因此丢失了夏令时的信息。
//
// 除了基本的“墙钟”读取外，Time 还可能包含当前进程的单调计时器读取，以提供比较或减法操作的额外精度。有关详细信息，请参阅包文档中的“单调计时器”部分。
//
// 注意 Go 中的 == 运算符不仅比较时间点，还比较 Location 和单调计时器读取。因此，在将 Time 值用作 map 或数据库键之前，需要确保所有值都设置了相同的 Location（通过使用 UTC 或 Local 方法），并移除单调计时器读取，这可以通过设置 t = t.Round(0) 来实现。一般来说，推荐使用 t.Equal(u) 而不是 t == u，因为 t.Equal 使用最准确的比较方法，并正确处理只有一个参数具有单调计时器读取的情况。
// md5:98b88205e2fd3c1e
// [提示]
//
//	type 时间 struct {
//	    时间点 time.Time
//	}
//
// [结束]
type Time struct { //hm:时间结构 cz:type Time
	F time.Time
} //md5:a8f1e158c36dbc97

//const (
//	hasMonotonic = 1 << 63
//	maxWall      = wallToInternal + (1<<33 - 1) // year 2157
//	minWall      = wallToInternal               // year 1885
//	nsecMask     = 1<<30 - 1
//	nsecShift    = 30
//) // md5:84fc15a2be736955

// [提示]
// const (
//
//	一月   Month = 1
//	二月   Month = 2
//	三月   Month = 3
//	四月   Month = 4
//	五月   Month = 5
//	六月   Month = 6
//	七月   Month = 7
//	八月   Month = 8
//	九月   Month = 9
//	十月   Month = 10
//	十一月 Month = 11
//	十二月 Month = 12
//
// )
//
// 注意：在Go语言中，常量名称一般使用英文，以保持代码的国际通用性。这里为了符合你的需求，我将其翻译成了中文，但在实际编程中并不推荐这样做。
// [结束]
const (
	January   Month = 1  //qm:常量_月份_一月 cz:January Month = 1
	February  Month = 2  //qm:常量_月份_二月 cz:February Month = 2
	March     Month = 3  //qm:常量_月份_三月 cz:March Month = 3
	April     Month = 4  //qm:常量_月份_四月 cz:April Month = 4
	May       Month = 5  //qm:常量_月份_五月 cz:May Month = 5
	June      Month = 6  //qm:常量_月份_六月 cz:June Month = 6
	July      Month = 7  //qm:常量_月份_七月 cz:July Month = 7
	August    Month = 8  //qm:常量_月份_八月 cz:August Month = 8
	September Month = 9  //qm:常量_月份_九月 cz:September Month = 9
	October   Month = 10 //qm:常量_月份_十月 cz:October Month = 10
	November  Month = 11 //qm:常量_月份_十一月 cz:November Month = 11
	December  Month = 12 //qm:常量_月份_十二月 cz:December Month = 12
) // md5:147c9abc849d3405

// [提示]
// const (
//
//	周日    Weekday = 0
//	周一    Weekday = 1
//	周二   Weekday = 2
//	周三   Weekday = 3
//	周四   Weekday = 4
//	周五    Weekday = 5
//	周六    Weekday = 6
//
// )
// [结束]
const (
	Sunday    Weekday = 0 //qm:常量_星期_周日 cz:Sunday Weekday = 0
	Monday    Weekday = 1 //qm:常量_星期_周一 cz:Monday Weekday = 1
	Tuesday   Weekday = 2 //qm:常量_星期_周二 cz:Tuesday Weekday = 2
	Wednesday Weekday = 3 //qm:常量_星期_周三 cz:Wednesday Weekday = 3
	Thursday  Weekday = 4 //qm:常量_星期_周四 cz:Thursday Weekday = 4
	Friday    Weekday = 5 //qm:常量_星期_周五 cz:Friday Weekday = 5
	Saturday  Weekday = 6 //qm:常量_星期_周六 cz:Saturday Weekday = 6
) // md5:9efbd1538cfcbf96

// [提示]
// const (
//
//	绝对零年        = -292277022399
//
//	内部年份        = 1
//
//	绝对到内部秒差   = (绝对零年 - 内部年份) * 365.2425 * 秒_per_天
//	内部到绝对秒差   = -绝对到内部秒差
//
//	Unix到内部秒差   = (1969*365 + 1969/4 - 1969/100 + 1969/400) * 秒_per_天
//	内部到Unix秒差   = -Unix到内部秒差
//
//	墙钟到内部秒差   = (1884*365 + 1884/4 - 1884/100 + 1884/400) * 秒_per_天
//
// )
// [结束]
const (
	// 对于内部计算的无符号零年。
	// 必须是400的模1，但在此之前的时间将无法正确计算，
	// 但可以根据需要随意更改。
	// md5:9e3ce821b0e5bb93
	absoluteZeroYear = -292277022399

	// Time零值的年份。
	// 由下面的unixToInternal计算中假设。
	// md5:dc0def73b210ce3b
	internalYear = 1

	// 用于在内部时间和绝对时间或Unix时间之间转换的偏移量。. md5:29a0da9f73dbc84d
	absoluteToInternal int64 = (absoluteZeroYear - internalYear) * 365.2425 * secondsPerDay
	internalToAbsolute       = -absoluteToInternal

	unixToInternal int64 = (1969*365 + 1969/4 - 1969/100 + 1969/400) * secondsPerDay
	internalToUnix int64 = -unixToInternal

	wallToInternal int64 = (1884*365 + 1884/4 - 1884/100 + 1884/400) * secondsPerDay
) // md5:f6e0b667fa556e68

// [提示]
// const (
//
//	最小时间间隔 Duration = -1 << 63
//	最大时间间隔 Duration = 1<<63 - 1
//
// )
// [结束]
const (
	minDuration Duration = -1 << 63
	maxDuration Duration = 1<<63 - 1
) // md5:7fa4af1b9114d2d1

// [提示]
// const (
//
//	纳秒      Duration = 1
//	微秒              = 1000 * 纳秒
//	毫秒              = 1000 * 微秒
//	秒                   = 1000 * 毫秒
//	分钟                = 60 * 秒
//	小时                 = 60 * 分钟
//
// )
// [结束]
const (
	Nanosecond  Duration = 1                  //qm:常量_时长_纳秒 cz:Nanosecond Duration = 1
	Microsecond          = 1000 * Nanosecond  //qm:常量_时长_微秒 cz:Microsecond = 1000 *
	Millisecond          = 1000 * Microsecond //qm:常量_时长_毫秒 cz:Millisecond = 1000 *
	Second               = 1000 * Millisecond //qm:常量_时长_秒 cz:Second = 1000 *
	Minute               = 60 * Second        //qm:常量_时长_分钟 cz:Minute = 60 *
	Hour                 = 60 * Minute        //qm:常量_时长_小时 cz:Hour = 60 *
) // md5:a064d6044cbb28fe

// [提示]
// const (
//
//	分钟秒数 = 60
//	小时秒数 = 60 * 分钟秒数
//	天秒数   = 24 * 小时秒数
//	周秒数   = 7 * 天秒数
//	四百年天数 = 365*400 + 97
//	百年天数   = 365*100 + 24
//	四年天数   = 365*4 + 1
//
// )
// [结束]
const (
	secondsPerMinute = 60
	secondsPerHour   = 60 * secondsPerMinute
	secondsPerDay    = 24 * secondsPerHour
	secondsPerWeek   = 7 * secondsPerDay
	daysPer400Years  = 365*400 + 97
	daysPer100Years  = 365*100 + 24
	daysPer4Years    = 365*4 + 1
) // md5:c9382aa097574f5d

//const (
//	timeBinaryVersionV1 byte = iota + 1 // For general situation
//	timeBinaryVersionV2                 // For LMT only
//) // md5:7df46a836c46218d

// After 函数返回一个布尔值，表示时间点t是否在u之后。. md5:750eca8bb04e1a25
// [提示:] func (t 时间) 在(u 时间) 后(bool) {}
// ff:是否为之后
// u:时间
func (t Time) After(u Time) bool { //md5:5dfb6db123b2d436
	return t.F.After(u.F)
}

// Before 返回时间点t是否早于u。. md5:36690a50c1e8d9d4
// [提示:] func (t 时间) 在(u 时间) 前 bool {}
// ff:是否为之前
// u:时间
func (t Time) Before(u Time) bool { //md5:d098b27691e377d5
	return t.F.Before(u.F)
}

// Compare 比较时间点 t 与 u。如果 t 在 u 之前，返回 -1；如果 t 在 u 之后，返回 +1；如果它们相同，返回 0。
// md5:aad24dff20f54b6b
// [提示:] func (t 时间) 比较(u 时间) 整型 {}
// ff:比较
// u:时间
func (t Time) Compare(u Time) int { //md5:35c931a4ac1e3dea
	return t.F.Compare(u.F)
}

// Equal 报告 t 和 u 是否表示相同的时间瞬间。
// 即使两个时间处于不同的时区，它们也可能相等。
// 例如，6:00 +0200 和 4:00 UTC 是相等的。
// 请参阅 Time 类型的文档以了解使用 == 比较 Time 值时的陷阱；
// 大多数代码应改使用 Equal 方法。
// md5:e1f4d776dce4bc1a
// [提示:] func (t 时间) 是否相等(u 时间) bool {}
// ff:是否相等
// u:时间
func (t Time) Equal(u Time) bool { //md5:02203c2b950ebe2d
	return t.F.Equal(u.F)
}

// String 返回月份的英文名称（"January"，"February"，...）。. md5:b9366b746afcb481
// [提示:] func (m 月份) 字符串() 字符串 {}
// ff:取英文月份
func (m Month) String() string { //md5:014a3380bef17d93
	return time.Month(m).String()
}

// String 返回该天的英文名称（"Sunday"，"Monday"，...）。. md5:129fa202c9139f04
// [提示:] func (d 周几) 文本表示() 字符串 {}
// ff:取英文星期几
func (d Weekday) String() string { //md5:3aa0ba35fd23d3a0
	return time.Weekday(d).String()
}

// IsZero 判断 t 是否表示时间的零点，
// 即公元1年1月1日，UTC时间00:00:00。
// md5:4e2b46d4fa63a878
// [提示:] func (t 时间) 是否为零() bool {}
// ff:是否为零
func (t Time) IsZero() bool { //md5:fd95a0aa1783c81e
	return t.F.IsZero()
}

// Date 返回 t 发生的年、月、日。. md5:47962c441720015d
// [提示:] func (t 时间) 日期() (年份 int, 月份 Month, 日 int) {}
// ff:取日期
// year:年份
// month:月份
// day:日
func (t Time) Date() (year int, month Month, day int) { //md5:a5a774f5fec8bb67
	y, m, d := t.F.Date()
	return y, Month(m), d
}

// Year 返回t所发生的年份。. md5:d47b3752d238b9f1
// [提示:] func (t 时间) 年份() int {}
// ff:取年份
func (t Time) Year() int { //md5:01bea20ebc40f05e
	return t.F.Year()
}

// Month 返回由 t 指定的年份中的月份。. md5:84f113a801a5eb29
// [提示:] func (t 时间) 月份() 月份 {}
// ff:取月份
func (t Time) Month() Month { //md5:12e0562e3bfdd479
	return Month(t.F.Month())
}

// Day返回由t指定的月份中的某一天。. md5:1ac7526d83edc95a
// [提示:] func (t 时间) 获取日() int {}
// ff:取日
func (t Time) Day() int { //md5:0b0c2cb83d99f883
	return t.F.Day()
}

// Weekday 返回由 t 指定的星期几。. md5:68508f1017d5a5a0
// [提示:] func (t 时间) 星期() 星期 {}
// ff:取星期几
func (t Time) Weekday() Weekday { //md5:0eb77d8496ff6452
	return Weekday(t.F.Weekday())
}

// ISOWeek 返回t所在的一周的ISO 8601年和周数。
// 周范围从1到53。一年中的1月1日至1月3日可能属于前一年的第52或53周，
// 而12月29日至12月31日可能属于下一年的第1周。
// md5:86fba2ca5bd03a59
// [提示:] func (t 时间) ISO周() (年 int, 周数 int) {}
// ff:取ISO年周
// year:年
// week:周数
func (t Time) ISOWeek() (year, week int) { //md5:720dde5fce39726b
	return t.F.ISOWeek()
}

// Clock返回t指定的一天中的小时、分钟和秒。. md5:19d5f7c98fc60342
// [提示:] func (t 时间) 时钟() (小时, 分钟, 秒 int) {}
// ff:取时分秒
// hour:小时
// min:分钟
// sec:秒
func (t Time) Clock() (hour, min, sec int) { //md5:b31ce25ccbb0b1b8
	return t.F.Clock()
}

// Hour返回由t指定的日期中的小时，范围为[0, 23]。. md5:02a015f0f53b5ff6
// [提示:] func (t 时间) 小时() int {}
// ff:取小时
func (t Time) Hour() int { //md5:efa7745ae50c29d5
	return t.F.Hour()
}

// Minute返回t指定的小时内的分钟偏移量，范围为[0, 59]。. md5:d369275980351e92
// [提示:] func (t 时间) 分钟() int {}
// ff:取分钟
func (t Time) Minute() int { //md5:c0dd1485a43bcd50
	return t.F.Minute()
}

// Second 返回 t 所指定分钟内的秒偏移量，范围为 [0, 59]。. md5:17d201e831d4d083
// [提示:] func (t 时间) 秒() int {}
// ff:取秒
func (t Time) Second() int { //md5:e2014d911e6cd34e
	return t.F.Second()
}

// Nanosecond 返回给定时间 t 所在秒内的纳秒偏移，范围为 [0, 999999999]。
// md5:c1dcd3dd99062cf7
// [提示:] func (t 时间) 纳秒() int {}
// ff:取纳秒
func (t Time) Nanosecond() int { //md5:89292d986d7e3147
	return t.F.Nanosecond()
}

// YearDay 返回由 t 指定的年中的第几天，非闰年为 [1, 365] 范围，闰年为 [1, 366] 范围。
// md5:ae8ebdaab4474241
// [提示:] func (t 时间) 年份天数() int {}
// ff:取第几天
func (t Time) YearDay() int { //md5:0b29732a13a1b75a
	return t.F.YearDay()
}

// String 返回一个表示持续时间的字符串，格式为 "72h3m0.5s"。
// 前导零的单位会被省略。作为一个特殊的情况，小于一秒的持续时间会使用更小的单位（毫秒、微秒或纳秒）来确保
// 首位数字非零。零持续时间格式化为 0s。
// md5:339f14595ff43024
// [提示:] func (d 时长) 字符串() 字符串 {}
// ff:
func (d Duration) String() string { //md5:7e1f3f8f34886861
	return time.Duration(d).String()
}

// Nanoseconds 返回持续时间作为整数纳秒计数。. md5:c2f1f3ae945976b2
// [提示:] func (d 时长) 纳秒() int64 {}
// ff:取纳秒数
func (d Duration) Nanoseconds() int64 { //md5:31b7144fa28ceb5b
	return time.Duration(d).Nanoseconds()
}

// Microseconds 返回Duration作为以微秒为单位的整数计数。. md5:cd88141d82127466
// [提示:] func (d 时长) 微秒() int64 {}
// ff:取微秒数
func (d Duration) Microseconds() int64 { //md5:d4942111d046902a
	return time.Duration(d).Microseconds()
}

// Milliseconds 返回Duration作为整数毫秒计数。. md5:dfe2b5e9f874e024
// [提示:] func (d 时长) 毫秒() int64 {}
// ff:取毫秒数
func (d Duration) Milliseconds() int64 { //md5:6cc34c02dbb54bc2
	return time.Duration(d).Milliseconds()
}

// Seconds 将持续时间以秒为单位表示为浮点数。. md5:0f83173e0d4294bc
// [提示:] func (d 时长) 秒数() 浮点数 {}
// ff:取秒数
func (d Duration) Seconds() float64 { //md5:75def545f7232495
	return time.Duration(d).Seconds()
}

// Minutes 将持续时间转换为以分钟为单位的浮点数。. md5:49e3673b4db2654b
// [提示:] func (d 时长) 分钟数() 浮点型 {}
// ff:取分钟数
func (d Duration) Minutes() float64 { //md5:900bce301a35981f
	return time.Duration(d).Minutes()
}

// Hours 返回Duration作为小时的浮点数表示。. md5:8ed8d3411cf7a7e8
// [提示:] func (d 时长) 小时() 浮点数 {}
// ff:取小时数
func (d Duration) Hours() float64 { //md5:d5ebe85a747f47bc
	return time.Duration(d).Hours()
}

// Truncate 函数将浮点数 d 向零进行四舍五入，结果为 m 的倍数。如果 m 小于等于 0，则 Truncate 函数直接返回 d，不做任何改变。
// md5:1f339b99307b7c19
// [提示:] func (d 时长) 截断(间隔时长) 返回时长 {}
// ff:截断并按下取整
// m:
func (d Duration) Truncate(m Duration) Duration { //md5:fe65fdff50fff8d6
	return Duration(time.Duration(d).Truncate(time.Duration(m)))
}

// Round 返回将 d 四舍五入到最接近 m 的整数倍的结果。
// 对于中间值的舍入行为是远离零进行舍入。
// 如果结果超出了可以存储在 Duration 类型中的最大（或最小）
// 值，Round 将返回最大（或最小）的持续时间。
// 如果 m <= 0，则 Round 不做改变，直接返回 d。
// md5:f3481afd1effbd95
// [提示:] func (d 时长) 轮整到(时长单位) 返回时长 {}
// ff:截断并按四舍五入
// m:
func (d Duration) Round(m Duration) Duration { //md5:2d3f614e4b10efeb
	return Duration(time.Duration(d).Round(time.Duration(m)))
}

// Abs 返回 d 的绝对值。
// 特殊情况下，math.MinInt64 转换为 math.MaxInt64。
// md5:8a84b8cfa75a6b01
// [提示:] func (d 时长) 绝对值() 时长 {}
// ff:取绝对值
func (d Duration) Abs() Duration { //md5:ce1eb65bdaa80e6d
	return Duration(time.Duration(d).Abs())
}

// Add 返回时间 t 加上 d。. md5:35c7d797bb6a1d97
// [提示:] func (t 时间) 添加(d 时长) 时间 {}
// ff:增加时长
// d:时长
func (t Time) Add(d Duration) Time { //md5:8720c0d0de353e9c
	return Time{t.F.Add(time.Duration(d))}
}

// Sub 返回两个时间差 t-u。如果结果超过了Duration类型可以存储的最大（或最小）值，
// 则返回最大（或最小）的时间差。
// 要计算时间差 t-d，可以使用 t.Add(-d)。
// md5:c975e5087c03d3b9
// [提示:] func (t 时间) 减去(u 时间) 时长 {}
// ff:减去
// u:时间
func (t Time) Sub(u Time) Duration { //md5:df77f9b3e6bf5ada
	return Duration(t.F.Sub(u.F))
}

// Since 计算并返回从 t 开始经过的时间。
// 这是 time.Now().Sub(t) 的简写形式。
// md5:f8c050ab8ed2afc9
// [提示:] func 自从(t 时间) 时长 {}
// ff:
// t:时间
func Since(t Time) Duration { //md5:810dd8f30988f658
	return Duration(time.Since(t.F))
}

// Until 返回从现在到t的持续时间。
// 它是t.Sub(time.Now())的简写。
// md5:b08a1c65063fce30
// [提示:] func 直到(t 时间) 间隔 {}
// ff:
// t:
func Until(t Time) Duration { //md5:63b7023bdca00f52
	return Duration(time.Until(t.F))
}

// AddDate函数返回给定年份、月份和天数后对应的时间。例如，将-1、2和3应用到2011年1月1日，会返回2010年3月4日。
//
// 注意日期本质上与时区相关联，像天这样的历法周期并没有固定的持续时间。AddDate函数使用Time值的Location来确定这些持续时间。这意味着相同的AddDate参数可能会根据基础Time值及其Location产生不同的绝对时间偏移。例如，将0、0和1应用到3月27日12:00上，始终会返回3月28日12:00。在某些地点和某些年份中，这是一次24小时的偏移。而在其他地方，由于夏令时转换，它可能只有23小时的偏移。
//
// AddDate的结果规范化方式与Date相同，例如，向10月31日添加一个月，结果将是12月1日，这是对11月31日的规范化形式。
// md5:2eb1b728ab1baa75
// [提示:] func (t 时间) 添加日期(years int, months int, days int) 时间 {}
// ff:
// years:
// months:
// days:
func (t Time) AddDate(years int, months int, days int) Time { //md5:f604377cb448bd9a
	return Time{t.F.AddDate(years, months, days)}
}

// Now返回当前的本地时间。. md5:449d3f2d7b2dcfde
// [提示:] func 现在() 时间 {}
// ff:取当前时间
func Now() Time { //md5:b2c3ce5c3a0dddd8
	return Time{time.Now()}
}

// UTC 返回将位置设置为UTC的t。. md5:cf62de83ef34698a
// [提示:] func (t 时间) 通用协调时间() 时间 {}
// ff:取UTC时间
func (t Time) UTC() Time { //md5:dfc0ea55e77bdd85
	return Time{t.F.UTC()}
}

// Local 将 t 的时区设置为本地时间并返回。. md5:9497f35c6f715db4
// [提示:] func (t 时间) 本地化() 时间 {}
// ff:取系统时间
func (t Time) Local() Time { //md5:773d045186c2d5e6
	return Time{t.F.Local()}
}

// In 返回一个表示相同时间点的t的副本，但将副本的地理位置信息设置为loc，用于显示目的。
// 如果loc为nil，In会引发panic。
// md5:e57190592f953463
// [提示:] func (t 时间) 在(时区 *时区) 时间 {}
// ff:
// loc:
func (t Time) In(loc *Location) Time { //md5:cdf90b2adf199e29
	if loc == nil {
		panic("time.In方法,loc参数为nil,触发panic")
	}
	return Time{t.F.In(&loc.F)}
}

// Location 返回与 t 关联的时间区域信息。. md5:fd562b3c47810a07
// [提示:] func (t 时间) 时区() *地理位置 {}
// ff:取时区
func (t Time) Location() *Location { //md5:320904942ebf3a70
	返回 := t.F.Location()
	if 返回 == nil {
		panic("time.Location方法,返回为nil,触发panic")
	}
	return &Location{*返回}
}

// Zone 计算在时间t生效的时区，返回该时区的缩写名称（如"CET"）及其相对于UTC向东的秒偏移量。
// md5:6bc1ff0aad02056b
// [提示:] func (t 时间) 时区() (名称 string, 偏移量 int) {}
// ff:取时区名称
// name:时区名称
// offset:秒偏移量
func (t Time) Zone() (name string, offset int) { //md5:f6bf762564bdb7bd
	return t.F.Zone()
}

// ZoneBounds 返回在时间 t 有效的时间区域边界。
// 区域从 start 开始，下一个区域从 end 开始。
// 如果区域从时间的开始时刻开始，start 将被返回为零时间。
// 如果区域持续到永远，end 将被返回为零时间。
// 返回时间的地点与 t 相同。
// md5:15d397b4b14f6445
// [提示:] func (t 时间) 时区边界() (开始时间, 结束时间 时间) {}
// ff:取时区边界
// start:开始时间
// end:结束时间
func (t Time) ZoneBounds() (start, end Time) { //md5:764bd854389bfd7e
	t1, t2 := t.F.ZoneBounds()
	return Time{t1}, Time{t2}
}

// Unix 返回 t 作为 Unix 时间，即自 1970 年 1 月 1 日 UTC 起经过的秒数。结果不依赖于与 t 关联的位置。
// 类 Unix 操作系统通常记录时间为以秒为单位的 32 位计数，但由于这里的方法返回的是 64 位值，因此在未来数十亿年或过去也是有效的。
// md5:6eee6141cce29453
// [提示:] func (t 时间) Unix() 整型六十四位数 {}
// ff:取Unix秒时间戳
func (t Time) Unix() int64 { //md5:708a59c5afafb133
	return t.F.Unix()
}

// UnixMilli 将 t 转换为自 1970 年 1 月 1 日 UTC 起经过的毫秒数。如果超过 int64 类型可以表示的最大值（即 1970 年前或后超过 292 百万年的时间），结果是未定义的。转换结果不受 t 关联的位置影响。
// md5:fac7c190fc89c5a2
// [提示:] func (t 时间) Unix毫秒() int64 {}
// ff:取Unix毫秒时间戳
func (t Time) UnixMilli() int64 { //md5:8ad355b089164e81
	return t.F.UnixMilli()
}

// UnixMicro 返回 t 作为 Unix 时间，即从 1970 年 1 月 1 日 UTC 开始流逝的微秒数。
// 如果微秒级的 Unix 时间不能被 int64 表示（日期早于公元前 290307 年或晚于公元 294246 年），结果是未定义的。
// 结果与 t 关联的时区无关。
// md5:f78d98a12030333c
// [提示:] func (t 时间) Unix微秒() int64 {}
// ff:取Unix微秒时间戳
func (t Time) UnixMicro() int64 { //md5:df74c29f1d20ba8a
	return t.F.UnixMicro()
}

// UnixNano 返回 t 的 Unix 时间，即自 1970 年 1 月 1 日 UTC 以来经过的纳秒数。如果以纳秒为单位的 Unix 时间不能用 int64 表示（日期早于 1678 年或晚于 2262 年），结果是未定义的。注意，这意味着对零值 Time 调用 UnixNano 的结果是未定义的。结果与 t 关联的位置无关。
// md5:ca4946b1370d3715
// [提示:] func (t 时间) UnixNano() int64 {} // 返回时间t的Unix纪元（1970年1月1日00:00:00 UTC）以来的纳秒数。
// ff:取Unix纳秒时间戳
func (t Time) UnixNano() int64 { //md5:e37c334610d19b98
	return t.F.UnixNano()
}

// MarshalBinary 实现了 encoding.BinaryMarshaler 接口。. md5:330de8d0c2918ee7
// [提示:] func (t 时间) 序列化二进制() ([]字节, 错误) {}
// ff:
func (t Time) MarshalBinary() ([]byte, error) { //md5:da3dea82d6dc9326
	return t.F.MarshalBinary()
}

// UnmarshalBinary实现了encoding.BinaryUnmarshaler接口。. md5:943a77fbd3310dbb
// [提示:] func (t *时间) 解析二进制数据(data []字节) 错误 {}
// ff:
// data:
func (t *Time) UnmarshalBinary(data []byte) error { //md5:a4a2ca872058057e
	return t.F.UnmarshalBinary(data)
}

// GobEncode 实现了 gob.GobEncoder 接口。. md5:2123f9f64f9e5072
// [提示:] func (t 时间) Gob编码() ([]字节, 错误) {}
// ff:
func (t Time) GobEncode() ([]byte, error) { //md5:bd28fa55c2574953
	return t.F.GobEncode()
}

// GobDecode 实现了 gob.GobDecoder 接口。. md5:cab7a156d7ee388c
// [提示:] func (t *时间) gob解码(data []字节) 错误 {}
// ff:
// data:
func (t *Time) GobDecode(data []byte) error { //md5:f1c31f760b2eb0ef
	return t.F.GobDecode(data)
}

// MarshalJSON 实现了 json.Marshaler 接口。
// 时间以 RFC 3339 格式（精确到毫秒）作为引号字符串表示。
// 如果时间戳无法表示为有效的 RFC 3339 格式（例如，年份超出范围），则会报告错误。
// md5:4c8ae3f9584f964f
// [提示:] func (t 时间) 序列化为JSON() ([]byte, 错误) {}
// ff:
func (t Time) MarshalJSON() ([]byte, error) { //md5:45b1e08736011f81
	return t.F.MarshalJSON()
}

// UnmarshalJSON实现了json.Unmarshaler接口。
// 时间必须是RFC 3339格式的引用字符串。
// md5:abd2b9bf29f6f8c2
// [提示:] func (t *时间) 解码JSON(data []字节) 错误 {}
// ff:
// data:
func (t *Time) UnmarshalJSON(data []byte) error { //md5:471a9f0fcd628553
	return t.F.UnmarshalJSON(data)
}

// 实现encoding.TextMarshaler接口的MarshalText方法。
// 时间以RFC 3339格式表示，包含秒以下的精度。
// 如果时间戳无法表示为有效的RFC 3339格式（例如，年份超出范围），则会报告错误。
// md5:a394cd3ee2adc414
// [提示:] func (t 时间) 序列化文本() ([]字节, 错误) {}
// ff:
func (t Time) MarshalText() ([]byte, error) { //md5:53833c0a151d9869
	return t.F.MarshalText()
}

// UnmarshalText 实现了 encoding.TextUnmarshaler 接口。
// 时间必须遵循 RFC 3339 格式。
// md5:0d2034e46c63ce93
// [提示:] func (t *时间) 解析文本(data []字节) error {}
// ff:
// data:
func (t *Time) UnmarshalText(data []byte) error { //md5:ba75915cc097f6f0
	return t.F.UnmarshalText(data)
}

// Unix根据给定的Unix时间（自1970年UTC 1月1日以来的秒数和纳秒数）返回本地时间。
// 可以传递超出[0, 999999999]范围的nsec值。不是所有sec值都对应一个时间值。其中一个例子是1<<63-1（最大的int64值）。
// md5:c015127582e82f24
// [提示:] func Unix(秒 int64, 毫秒 int64) 时间 {}
// ff:创建并按时间戳
// sec:可选秒
// nsec:可选毫秒
func Unix(sec int64, nsec int64) Time { //md5:8f372b5d16d3f416
	return Time{time.Unix(sec, nsec)}
}

// UnixMilli 返回给定 Unix 时间（从 UTC 的 1970 年 1 月 1 日开始计算的毫秒数）对应的本地时间。
// md5:b3c1c9e51febfd0d
// [提示:] func Unix毫秒(msec int64) 时间 {}
// ff:创建并按毫秒时间戳
// msec:毫秒时间戳
func UnixMilli(msec int64) Time { //md5:94c00e0069987ce8
	return Time{time.UnixMilli(msec)}
}

// UnixMicro 返回对应于给定Unix时间的本地时间，该Unix时间是从1970年1月1日UTC开始的usec微秒。
// md5:4f4e706fb188cc2e
// [提示:] func UnixMicro(微秒 int64) 时间 {}
// ff:创建并按微秒时间戳
// usec:微秒时间戳
func UnixMicro(usec int64) Time { //md5:a519cb1431da9d9d
	return Time{time.UnixMicro(usec)}
}

// IsDST 报告配置的时区中给定时间是否处于夏令时。. md5:48b3bf0c1d4aba01
// [提示:] func (t 时间) 是否夏令时() bool {}
// ff:是否为夏令时
func (t Time) IsDST() bool { //md5:6a98901b623dee23
	return t.F.IsDST()
}

// Date返回与给定位置中相应时间的适当时区对应的Time。
//
// 月份、日期、小时、分钟、秒和纳秒值可能超出常规范围，在转换期间将被规范化。例如，10月32日会转换为11月1日。
//
// 夏令时转换会跳过或重复时间。例如，在美国，2011年3月13日2:15am并未发生，而2011年11月6日1:15am发生了两次。在这种情况下，时区的选择（以及因此对应的时间）并不明确。Date返回的是两个相关时区中的一个正确时间，但并不能保证是哪一个。
//
// 如果loc为nil，Date函数会引发panic。
// md5:db3e840a23378f8a
// [提示:] func 创建日期(year int, 月份 Month, 日 int, 小时 int, 分钟 int, 秒 int, 毫秒 int, 时区 *时区) 时间 {}
// ff:创建时间
// year:年
// month:月
// day:日
// hour:时
// min:分
// sec:秒
// nsec:毫秒
// loc:时区
func Date(year int, month Month, day, hour, min, sec, nsec int, loc *Location) Time { //md5:fc73fa423236be91
	if loc == nil {
		panic("time.Date方法,loc参数为nil,触发panic")
	}
	return Time{time.Date(year, time.Month(month), day, hour, min, sec, nsec, &loc.F)}
}

// Truncate 函数将时间 t 向下舍入到 d 的倍数（从零时间开始）。如果 d 小于等于 0，Truncate 将去掉时间的单调时钟读数，但保持不变。
//
// Truncate 以绝对自零时间以来的持续时间对时间进行操作；它不处理时间的呈现形式。因此，Truncate(Hour) 可能返回一个非零分钟的时间，具体取决于时间的时区。
// md5:f72e0e00b245e691
// [提示:] func (t 时间) 截断(d 持续时间) 时间 {}
// ff:截断并按下取整
// d:时长
func (t Time) Truncate(d Duration) Time { //md5:6c9853dfe8f0ce4e
	return Time{t.F.Truncate(time.Duration(d))}
}

// Round 返回将时间 t 四舍五入到最接近 d 的整数倍的结果（从零时间开始计算）。
// 对于中间值的舍入规则是向上舍入。
// 如果 d <= 0，则 Round 返回 t，但会剥离任何单调时钟读数，其他不变。
//
// Round 是基于时间作为从零时间起的绝对持续时间来进行操作的；
// 它并不作用于时间的展示形式。因此，Round(Hour) 可能会返回一个带有非零分钟的时间，
// 具体取决于该时间的时区位置。
// md5:b2557220790fc058
// [提示:] func (t 时间) 轮廓(d 持续时间) 时间 {}
// ff:截断并按四舍五入
// d:时长
func (t Time) Round(d Duration) Time { //md5:514f2ad77ae8743a
	return Time{t.F.Round(time.Duration(d))}
}
