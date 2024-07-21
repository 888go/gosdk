
<原文开始>
	// The unsigned zero year for internal calculations.
	// Must be 1 mod 400, and times before it will not compute correctly,
	// but otherwise can be changed at will.
<原文结束>

# <翻译开始>
	// 对于内部计算的无符号零年。
	// 必须是400的模1，但在此之前的时间将无法正确计算，
	// 但可以根据需要随意更改。
	// md5:9e3ce821b0e5bb93
# <翻译结束>


<原文开始>
	// The year of the zero Time.
	// Assumed by the unixToInternal computation below.
<原文结束>

# <翻译开始>
	// Time零值的年份。
	// 由下面的unixToInternal计算中假设。
	// md5:dc0def73b210ce3b
# <翻译结束>


<原文开始>
// Offsets to convert between internal and absolute or Unix times.
<原文结束>

# <翻译开始>
// 用于在内部时间和绝对时间或Unix时间之间转换的偏移量。. md5:29a0da9f73dbc84d
# <翻译结束>


<原文开始>
// After reports whether the time instant t is after u.
<原文结束>

# <翻译开始>
// After 函数返回一个布尔值，表示时间点t是否在u之后。. md5:750eca8bb04e1a25
# <翻译结束>


<原文开始>
// Before reports whether the time instant t is before u.
<原文结束>

# <翻译开始>
// Before 返回时间点t是否早于u。. md5:36690a50c1e8d9d4
# <翻译结束>


<原文开始>
// Compare compares the time instant t with u. If t is before u, it returns -1;
// if t is after u, it returns +1; if they're the same, it returns 0.
<原文结束>

# <翻译开始>
// Compare 比较时间点 t 与 u。如果 t 在 u 之前，返回 -1；如果 t 在 u 之后，返回 +1；如果它们相同，返回 0。
// md5:aad24dff20f54b6b
# <翻译结束>


<原文开始>
// Equal reports whether t and u represent the same time instant.
// Two times can be equal even if they are in different locations.
// For example, 6:00 +0200 and 4:00 UTC are Equal.
// See the documentation on the Time type for the pitfalls of using == with
// Time values; most code should use Equal instead.
<原文结束>

# <翻译开始>
// Equal 报告 t 和 u 是否表示相同的时间瞬间。
// 即使两个时间处于不同的时区，它们也可能相等。
// 例如，6:00 +0200 和 4:00 UTC 是相等的。
// 请参阅 Time 类型的文档以了解使用 == 比较 Time 值时的陷阱；
// 大多数代码应改使用 Equal 方法。
// md5:e1f4d776dce4bc1a
# <翻译结束>


<原文开始>
// String returns the English name of the month ("January", "February", ...).
<原文结束>

# <翻译开始>
// String 返回月份的英文名称（"January"，"February"，...）。. md5:b9366b746afcb481
# <翻译结束>


<原文开始>
// String returns the English name of the day ("Sunday", "Monday", ...).
<原文结束>

# <翻译开始>
// String 返回该天的英文名称（"Sunday"，"Monday"，...）。. md5:129fa202c9139f04
# <翻译结束>


<原文开始>
// IsZero reports whether t represents the zero time instant,
// January 1, year 1, 00:00:00 UTC.
<原文结束>

# <翻译开始>
// IsZero 判断 t 是否表示时间的零点，
// 即公元1年1月1日，UTC时间00:00:00。
// md5:4e2b46d4fa63a878
# <翻译结束>


<原文开始>
// Date returns the year, month, and day in which t occurs.
<原文结束>

# <翻译开始>
// Date 返回 t 发生的年、月、日。. md5:47962c441720015d
# <翻译结束>


<原文开始>
// Year returns the year in which t occurs.
<原文结束>

# <翻译开始>
// Year 返回t所发生的年份。. md5:d47b3752d238b9f1
# <翻译结束>


<原文开始>
// Month returns the month of the year specified by t.
<原文结束>

# <翻译开始>
// Month 返回由 t 指定的年份中的月份。. md5:84f113a801a5eb29
# <翻译结束>


<原文开始>
// Day returns the day of the month specified by t.
<原文结束>

# <翻译开始>
// Day返回由t指定的月份中的某一天。. md5:1ac7526d83edc95a
# <翻译结束>


<原文开始>
// Weekday returns the day of the week specified by t.
<原文结束>

# <翻译开始>
// Weekday 返回由 t 指定的星期几。. md5:68508f1017d5a5a0
# <翻译结束>


<原文开始>
// ISOWeek returns the ISO 8601 year and week number in which t occurs.
// Week ranges from 1 to 53. Jan 01 to Jan 03 of year n might belong to
// week 52 or 53 of year n-1, and Dec 29 to Dec 31 might belong to week 1
// of year n+1.
<原文结束>

# <翻译开始>
// ISOWeek 返回t所在的一周的ISO 8601年和周数。
// 周范围从1到53。一年中的1月1日至1月3日可能属于前一年的第52或53周，
// 而12月29日至12月31日可能属于下一年的第1周。
// md5:86fba2ca5bd03a59
# <翻译结束>


<原文开始>
// Clock returns the hour, minute, and second within the day specified by t.
<原文结束>

# <翻译开始>
// Clock返回t指定的一天中的小时、分钟和秒。. md5:19d5f7c98fc60342
# <翻译结束>


<原文开始>
// Hour returns the hour within the day specified by t, in the range [0, 23].
<原文结束>

# <翻译开始>
// Hour返回由t指定的日期中的小时，范围为[0, 23]。. md5:02a015f0f53b5ff6
# <翻译结束>


<原文开始>
// Minute returns the minute offset within the hour specified by t, in the range [0, 59].
<原文结束>

# <翻译开始>
// Minute返回t指定的小时内的分钟偏移量，范围为[0, 59]。. md5:d369275980351e92
# <翻译结束>


<原文开始>
// Second returns the second offset within the minute specified by t, in the range [0, 59].
<原文结束>

# <翻译开始>
// Second 返回 t 所指定分钟内的秒偏移量，范围为 [0, 59]。. md5:17d201e831d4d083
# <翻译结束>


<原文开始>
// Nanosecond returns the nanosecond offset within the second specified by t,
// in the range [0, 999999999].
<原文结束>

# <翻译开始>
// Nanosecond 返回给定时间 t 所在秒内的纳秒偏移，范围为 [0, 999999999]。
// md5:c1dcd3dd99062cf7
# <翻译结束>


<原文开始>
// YearDay returns the day of the year specified by t, in the range [1,365] for non-leap years,
// and [1,366] in leap years.
<原文结束>

# <翻译开始>
// YearDay 返回由 t 指定的年中的第几天，非闰年为 [1, 365] 范围，闰年为 [1, 366] 范围。
// md5:ae8ebdaab4474241
# <翻译结束>


<原文开始>
// String returns a string representing the duration in the form "72h3m0.5s".
// Leading zero units are omitted. As a special case, durations less than one
// second format use a smaller unit (milli-, micro-, or nanoseconds) to ensure
// that the leading digit is non-zero. The zero duration formats as 0s.
<原文结束>

# <翻译开始>
// String 返回一个表示持续时间的字符串，格式为 "72h3m0.5s"。
// 前导零的单位会被省略。作为一个特殊的情况，小于一秒的持续时间会使用更小的单位（毫秒、微秒或纳秒）来确保
// 首位数字非零。零持续时间格式化为 0s。
// md5:339f14595ff43024
# <翻译结束>


<原文开始>
// Nanoseconds returns the duration as an integer nanosecond count.
<原文结束>

# <翻译开始>
// Nanoseconds 返回持续时间作为整数纳秒计数。. md5:c2f1f3ae945976b2
# <翻译结束>


<原文开始>
// Microseconds returns the duration as an integer microsecond count.
<原文结束>

# <翻译开始>
// Microseconds 返回Duration作为以微秒为单位的整数计数。. md5:cd88141d82127466
# <翻译结束>


<原文开始>
// Milliseconds returns the duration as an integer millisecond count.
<原文结束>

# <翻译开始>
// Milliseconds 返回Duration作为整数毫秒计数。. md5:dfe2b5e9f874e024
# <翻译结束>


<原文开始>
// Seconds returns the duration as a floating point number of seconds.
<原文结束>

# <翻译开始>
// Seconds 将持续时间以秒为单位表示为浮点数。. md5:0f83173e0d4294bc
# <翻译结束>


<原文开始>
// Minutes returns the duration as a floating point number of minutes.
<原文结束>

# <翻译开始>
// Minutes 将持续时间转换为以分钟为单位的浮点数。. md5:49e3673b4db2654b
# <翻译结束>


<原文开始>
// Hours returns the duration as a floating point number of hours.
<原文结束>

# <翻译开始>
// Hours 返回Duration作为小时的浮点数表示。. md5:8ed8d3411cf7a7e8
# <翻译结束>


<原文开始>
// Truncate returns the result of rounding d toward zero to a multiple of m.
// If m <= 0, Truncate returns d unchanged.
<原文结束>

# <翻译开始>
// Truncate 函数将浮点数 d 向零进行四舍五入，结果为 m 的倍数。如果 m 小于等于 0，则 Truncate 函数直接返回 d，不做任何改变。
// md5:1f339b99307b7c19
# <翻译结束>


<原文开始>
// Round returns the result of rounding d to the nearest multiple of m.
// The rounding behavior for halfway values is to round away from zero.
// If the result exceeds the maximum (or minimum)
// value that can be stored in a Duration,
// Round returns the maximum (or minimum) duration.
// If m <= 0, Round returns d unchanged.
<原文结束>

# <翻译开始>
// Round 返回将 d 四舍五入到最接近 m 的整数倍的结果。
// 对于中间值的舍入行为是远离零进行舍入。
// 如果结果超出了可以存储在 Duration 类型中的最大（或最小）
// 值，Round 将返回最大（或最小）的持续时间。
// 如果 m <= 0，则 Round 不做改变，直接返回 d。
// md5:f3481afd1effbd95
# <翻译结束>


<原文开始>
// Abs returns the absolute value of d.
// As a special case, math.MinInt64 is converted to math.MaxInt64.
<原文结束>

# <翻译开始>
// Abs 返回 d 的绝对值。
// 特殊情况下，math.MinInt64 转换为 math.MaxInt64。
// md5:8a84b8cfa75a6b01
# <翻译结束>


<原文开始>
// Add returns the time t+d.
<原文结束>

# <翻译开始>
// Add 返回时间 t 加上 d。. md5:35c7d797bb6a1d97
# <翻译结束>


<原文开始>
// Sub returns the duration t-u. If the result exceeds the maximum (or minimum)
// value that can be stored in a Duration, the maximum (or minimum) duration
// will be returned.
// To compute t-d for a duration d, use t.Add(-d).
<原文结束>

# <翻译开始>
// Sub 返回两个时间差 t-u。如果结果超过了Duration类型可以存储的最大（或最小）值，
// 则返回最大（或最小）的时间差。
// 要计算时间差 t-d，可以使用 t.Add(-d)。
// md5:c975e5087c03d3b9
# <翻译结束>


<原文开始>
// Since returns the time elapsed since t.
// It is shorthand for time.Now().Sub(t).
<原文结束>

# <翻译开始>
// Since 计算并返回从 t 开始经过的时间。
// 这是 time.Now().Sub(t) 的简写形式。
// md5:f8c050ab8ed2afc9
# <翻译结束>


<原文开始>
// Until returns the duration until t.
// It is shorthand for t.Sub(time.Now()).
<原文结束>

# <翻译开始>
// Until 返回从现在到t的持续时间。
// 它是t.Sub(time.Now())的简写。
// md5:b08a1c65063fce30
# <翻译结束>


<原文开始>
// AddDate returns the time corresponding to adding the
// given number of years, months, and days to t.
// For example, AddDate(-1, 2, 3) applied to January 1, 2011
// returns March 4, 2010.
//
// Note that dates are fundamentally coupled to timezones, and calendrical
// periods like days don't have fixed durations. AddDate uses the Location of
// the Time value to determine these durations. That means that the same
// AddDate arguments can produce a different shift in absolute time depending on
// the base Time value and its Location. For example, AddDate(0, 0, 1) applied
// to 12:00 on March 27 always returns 12:00 on March 28. At some locations and
// in some years this is a 24 hour shift. In others it's a 23 hour shift due to
// daylight savings time transitions.
//
// AddDate normalizes its result in the same way that Date does,
// so, for example, adding one month to October 31 yields
// December 1, the normalized form for November 31.
<原文结束>

# <翻译开始>
// AddDate函数返回给定年份、月份和天数后对应的时间。例如，将-1、2和3应用到2011年1月1日，会返回2010年3月4日。
//
// 注意日期本质上与时区相关联，像天这样的历法周期并没有固定的持续时间。AddDate函数使用Time值的Location来确定这些持续时间。这意味着相同的AddDate参数可能会根据基础Time值及其Location产生不同的绝对时间偏移。例如，将0、0和1应用到3月27日12:00上，始终会返回3月28日12:00。在某些地点和某些年份中，这是一次24小时的偏移。而在其他地方，由于夏令时转换，它可能只有23小时的偏移。
//
// AddDate的结果规范化方式与Date相同，例如，向10月31日添加一个月，结果将是12月1日，这是对11月31日的规范化形式。
// md5:2eb1b728ab1baa75
# <翻译结束>


<原文开始>
// Now returns the current local time.
<原文结束>

# <翻译开始>
// Now返回当前的本地时间。. md5:449d3f2d7b2dcfde
# <翻译结束>


<原文开始>
// UTC returns t with the location set to UTC.
<原文结束>

# <翻译开始>
// UTC 返回将位置设置为UTC的t。. md5:cf62de83ef34698a
# <翻译结束>


<原文开始>
// Local returns t with the location set to local time.
<原文结束>

# <翻译开始>
// Local 将 t 的时区设置为本地时间并返回。. md5:9497f35c6f715db4
# <翻译结束>


<原文开始>
// In returns a copy of t representing the same time instant, but
// with the copy's location information set to loc for display
// purposes.
//
// In panics if loc is nil.
<原文结束>

# <翻译开始>
// In 返回一个表示相同时间点的t的副本，但将副本的地理位置信息设置为loc，用于显示目的。
// 如果loc为nil，In会引发panic。
// md5:e57190592f953463
# <翻译结束>


<原文开始>
// Location returns the time zone information associated with t.
<原文结束>

# <翻译开始>
// Location 返回与 t 关联的时间区域信息。. md5:fd562b3c47810a07
# <翻译结束>


<原文开始>
// Zone computes the time zone in effect at time t, returning the abbreviated
// name of the zone (such as "CET") and its offset in seconds east of UTC.
<原文结束>

# <翻译开始>
// Zone 计算在时间t生效的时区，返回该时区的缩写名称（如"CET"）及其相对于UTC向东的秒偏移量。
// md5:6bc1ff0aad02056b
# <翻译结束>


<原文开始>
// ZoneBounds returns the bounds of the time zone in effect at time t.
// The zone begins at start and the next zone begins at end.
// If the zone begins at the beginning of time, start will be returned as a zero Time.
// If the zone goes on forever, end will be returned as a zero Time.
// The Location of the returned times will be the same as t.
<原文结束>

# <翻译开始>
// ZoneBounds 返回在时间 t 有效的时间区域边界。
// 区域从 start 开始，下一个区域从 end 开始。
// 如果区域从时间的开始时刻开始，start 将被返回为零时间。
// 如果区域持续到永远，end 将被返回为零时间。
// 返回时间的地点与 t 相同。
// md5:15d397b4b14f6445
# <翻译结束>


<原文开始>
// Unix returns t as a Unix time, the number of seconds elapsed
// since January 1, 1970 UTC. The result does not depend on the
// location associated with t.
// Unix-like operating systems often record time as a 32-bit
// count of seconds, but since the method here returns a 64-bit
// value it is valid for billions of years into the past or future.
<原文结束>

# <翻译开始>
// Unix 返回 t 作为 Unix 时间，即自 1970 年 1 月 1 日 UTC 起经过的秒数。结果不依赖于与 t 关联的位置。
// 类 Unix 操作系统通常记录时间为以秒为单位的 32 位计数，但由于这里的方法返回的是 64 位值，因此在未来数十亿年或过去也是有效的。
// md5:6eee6141cce29453
# <翻译结束>


<原文开始>
// UnixMilli returns t as a Unix time, the number of milliseconds elapsed since
// January 1, 1970 UTC. The result is undefined if the Unix time in
// milliseconds cannot be represented by an int64 (a date more than 292 million
// years before or after 1970). The result does not depend on the
// location associated with t.
<原文结束>

# <翻译开始>
// UnixMilli 将 t 转换为自 1970 年 1 月 1 日 UTC 起经过的毫秒数。如果超过 int64 类型可以表示的最大值（即 1970 年前或后超过 292 百万年的时间），结果是未定义的。转换结果不受 t 关联的位置影响。
// md5:fac7c190fc89c5a2
# <翻译结束>


<原文开始>
// UnixMicro returns t as a Unix time, the number of microseconds elapsed since
// January 1, 1970 UTC. The result is undefined if the Unix time in
// microseconds cannot be represented by an int64 (a date before year -290307 or
// after year 294246). The result does not depend on the location associated
// with t.
<原文结束>

# <翻译开始>
// UnixMicro 返回 t 作为 Unix 时间，即从 1970 年 1 月 1 日 UTC 开始流逝的微秒数。
// 如果微秒级的 Unix 时间不能被 int64 表示（日期早于公元前 290307 年或晚于公元 294246 年），结果是未定义的。
// 结果与 t 关联的时区无关。
// md5:f78d98a12030333c
# <翻译结束>


<原文开始>
// UnixNano returns t as a Unix time, the number of nanoseconds elapsed
// since January 1, 1970 UTC. The result is undefined if the Unix time
// in nanoseconds cannot be represented by an int64 (a date before the year
// 1678 or after 2262). Note that this means the result of calling UnixNano
// on the zero Time is undefined. The result does not depend on the
// location associated with t.
<原文结束>

# <翻译开始>
// UnixNano 返回 t 的 Unix 时间，即自 1970 年 1 月 1 日 UTC 以来经过的纳秒数。如果以纳秒为单位的 Unix 时间不能用 int64 表示（日期早于 1678 年或晚于 2262 年），结果是未定义的。注意，这意味着对零值 Time 调用 UnixNano 的结果是未定义的。结果与 t 关联的位置无关。
// md5:ca4946b1370d3715
# <翻译结束>


<原文开始>
// MarshalBinary implements the encoding.BinaryMarshaler interface.
<原文结束>

# <翻译开始>
// MarshalBinary 实现了 encoding.BinaryMarshaler 接口。. md5:330de8d0c2918ee7
# <翻译结束>


<原文开始>
// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface.
<原文结束>

# <翻译开始>
// UnmarshalBinary实现了encoding.BinaryUnmarshaler接口。. md5:943a77fbd3310dbb
# <翻译结束>


<原文开始>
// GobEncode implements the gob.GobEncoder interface.
<原文结束>

# <翻译开始>
// GobEncode 实现了 gob.GobEncoder 接口。. md5:2123f9f64f9e5072
# <翻译结束>


<原文开始>
// GobDecode implements the gob.GobDecoder interface.
<原文结束>

# <翻译开始>
// GobDecode 实现了 gob.GobDecoder 接口。. md5:cab7a156d7ee388c
# <翻译结束>


<原文开始>
// MarshalJSON implements the json.Marshaler interface.
// The time is a quoted string in the RFC 3339 format with sub-second precision.
// If the timestamp cannot be represented as valid RFC 3339
// (e.g., the year is out of range), then an error is reported.
<原文结束>

# <翻译开始>
// MarshalJSON 实现了 json.Marshaler 接口。
// 时间以 RFC 3339 格式（精确到毫秒）作为引号字符串表示。
// 如果时间戳无法表示为有效的 RFC 3339 格式（例如，年份超出范围），则会报告错误。
// md5:4c8ae3f9584f964f
# <翻译结束>


<原文开始>
// UnmarshalJSON implements the json.Unmarshaler interface.
// The time must be a quoted string in the RFC 3339 format.
<原文结束>

# <翻译开始>
// UnmarshalJSON实现了json.Unmarshaler接口。
// 时间必须是RFC 3339格式的引用字符串。
// md5:abd2b9bf29f6f8c2
# <翻译结束>


<原文开始>
// MarshalText implements the encoding.TextMarshaler interface.
// The time is formatted in RFC 3339 format with sub-second precision.
// If the timestamp cannot be represented as valid RFC 3339
// (e.g., the year is out of range), then an error is reported.
<原文结束>

# <翻译开始>
// 实现encoding.TextMarshaler接口的MarshalText方法。
// 时间以RFC 3339格式表示，包含秒以下的精度。
// 如果时间戳无法表示为有效的RFC 3339格式（例如，年份超出范围），则会报告错误。
// md5:a394cd3ee2adc414
# <翻译结束>


<原文开始>
// UnmarshalText implements the encoding.TextUnmarshaler interface.
// The time must be in the RFC 3339 format.
<原文结束>

# <翻译开始>
// UnmarshalText 实现了 encoding.TextUnmarshaler 接口。
// 时间必须遵循 RFC 3339 格式。
// md5:0d2034e46c63ce93
# <翻译结束>


<原文开始>
// Unix returns the local Time corresponding to the given Unix time,
// sec seconds and nsec nanoseconds since January 1, 1970 UTC.
// It is valid to pass nsec outside the range [0, 999999999].
// Not all sec values have a corresponding time value. One such
// value is 1<<63-1 (the largest int64 value).
<原文结束>

# <翻译开始>
// Unix根据给定的Unix时间（自1970年UTC 1月1日以来的秒数和纳秒数）返回本地时间。
// 可以传递超出[0, 999999999]范围的nsec值。不是所有sec值都对应一个时间值。其中一个例子是1<<63-1（最大的int64值）。
// md5:c015127582e82f24
# <翻译结束>


<原文开始>
// UnixMilli returns the local Time corresponding to the given Unix time,
// msec milliseconds since January 1, 1970 UTC.
<原文结束>

# <翻译开始>
// UnixMilli 返回给定 Unix 时间（从 UTC 的 1970 年 1 月 1 日开始计算的毫秒数）对应的本地时间。
// md5:b3c1c9e51febfd0d
# <翻译结束>


<原文开始>
// UnixMicro returns the local Time corresponding to the given Unix time,
// usec microseconds since January 1, 1970 UTC.
<原文结束>

# <翻译开始>
// UnixMicro 返回对应于给定Unix时间的本地时间，该Unix时间是从1970年1月1日UTC开始的usec微秒。
// md5:4f4e706fb188cc2e
# <翻译结束>


<原文开始>
// IsDST reports whether the time in the configured location is in Daylight Savings Time.
<原文结束>

# <翻译开始>
// IsDST 报告配置的时区中给定时间是否处于夏令时。. md5:48b3bf0c1d4aba01
# <翻译结束>


<原文开始>
// Date returns the Time corresponding to
//
//	yyyy-mm-dd hh:mm:ss + nsec nanoseconds
//
// in the appropriate zone for that time in the given location.
//
// The month, day, hour, min, sec, and nsec values may be outside
// their usual ranges and will be normalized during the conversion.
// For example, October 32 converts to November 1.
//
// A daylight savings time transition skips or repeats times.
// For example, in the United States, March 13, 2011 2:15am never occurred,
// while November 6, 2011 1:15am occurred twice. In such cases, the
// choice of time zone, and therefore the time, is not well-defined.
// Date returns a time that is correct in one of the two zones involved
// in the transition, but it does not guarantee which.
//
// Date panics if loc is nil.
<原文结束>

# <翻译开始>
// Date返回与给定位置中相应时间的适当时区对应的Time。
//
// 月份、日期、小时、分钟、秒和纳秒值可能超出常规范围，在转换期间将被规范化。例如，10月32日会转换为11月1日。
//
// 夏令时转换会跳过或重复时间。例如，在美国，2011年3月13日2:15am并未发生，而2011年11月6日1:15am发生了两次。在这种情况下，时区的选择（以及因此对应的时间）并不明确。Date返回的是两个相关时区中的一个正确时间，但并不能保证是哪一个。
//
// 如果loc为nil，Date函数会引发panic。
// md5:db3e840a23378f8a
# <翻译结束>


<原文开始>
// Truncate returns the result of rounding t down to a multiple of d (since the zero time).
// If d <= 0, Truncate returns t stripped of any monotonic clock reading but otherwise unchanged.
//
// Truncate operates on the time as an absolute duration since the
// zero time; it does not operate on the presentation form of the
// time. Thus, Truncate(Hour) may return a time with a non-zero
// minute, depending on the time's Location.
<原文结束>

# <翻译开始>
// Truncate 函数将时间 t 向下舍入到 d 的倍数（从零时间开始）。如果 d 小于等于 0，Truncate 将去掉时间的单调时钟读数，但保持不变。
//
// Truncate 以绝对自零时间以来的持续时间对时间进行操作；它不处理时间的呈现形式。因此，Truncate(Hour) 可能返回一个非零分钟的时间，具体取决于时间的时区。
// md5:f72e0e00b245e691
# <翻译结束>


<原文开始>
// Round returns the result of rounding t to the nearest multiple of d (since the zero time).
// The rounding behavior for halfway values is to round up.
// If d <= 0, Round returns t stripped of any monotonic clock reading but otherwise unchanged.
//
// Round operates on the time as an absolute duration since the
// zero time; it does not operate on the presentation form of the
// time. Thus, Round(Hour) may return a time with a non-zero
// minute, depending on the time's Location.
<原文结束>

# <翻译开始>
// Round 返回将时间 t 四舍五入到最接近 d 的整数倍的结果（从零时间开始计算）。
// 对于中间值的舍入规则是向上舍入。
// 如果 d <= 0，则 Round 返回 t，但会剥离任何单调时钟读数，其他不变。
//
// Round 是基于时间作为从零时间起的绝对持续时间来进行操作的；
// 它并不作用于时间的展示形式。因此，Round(Hour) 可能会返回一个带有非零分钟的时间，
// 具体取决于该时间的时区位置。
// md5:b2557220790fc058
# <翻译结束>


<原文开始>
// A Time represents an instant in time with nanosecond precision.
//
// Programs using times should typically store and pass them as values,
// not pointers. That is, time variables and struct fields should be of
// type time.Time, not *time.Time.
//
// A Time value can be used by multiple goroutines simultaneously except
// that the methods GobDecode, UnmarshalBinary, UnmarshalJSON and
// UnmarshalText are not concurrency-safe.
//
// Time instants can be compared using the Before, After, and Equal methods.
// The Sub method subtracts two instants, producing a Duration.
// The Add method adds a Time and a Duration, producing a Time.
//
// The zero value of type Time is January 1, year 1, 00:00:00.000000000 UTC.
// As this time is unlikely to come up in practice, the IsZero method gives
// a simple way of detecting a time that has not been initialized explicitly.
//
// Each time has an associated Location. The methods Local, UTC, and In return a
// Time with a specific Location. Changing the Location of a Time value with
// these methods does not change the actual instant it represents, only the time
// zone in which to interpret it.
//
// Representations of a Time value saved by the GobEncode, MarshalBinary,
// MarshalJSON, and MarshalText methods store the Time.Location's offset, but not
// the location name. They therefore lose information about Daylight Saving Time.
//
// In addition to the required “wall clock” reading, a Time may contain an optional
// reading of the current process's monotonic clock, to provide additional precision
// for comparison or subtraction.
// See the “Monotonic Clocks” section in the package documentation for details.
//
// Note that the Go == operator compares not just the time instant but also the
// Location and the monotonic clock reading. Therefore, Time values should not
// be used as map or database keys without first guaranteeing that the
// identical Location has been set for all values, which can be achieved
// through use of the UTC or Local method, and that the monotonic clock reading
// has been stripped by setting t = t.Round(0). In general, prefer t.Equal(u)
// to t == u, since t.Equal uses the most accurate comparison available and
// correctly handles the case when only one of its arguments has a monotonic
// clock reading.
<原文结束>

# <翻译开始>
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
# <翻译结束>


<原文开始>
// A Duration represents the elapsed time between two instants
// as an int64 nanosecond count. The representation limits the
// largest representable duration to approximately 290 years.
<原文结束>

# <翻译开始>
// Duration 表示两个瞬间之间的时间差，用一个 int64 秒级别的纳秒计数表示。这个表示方式限制了可表示的最长持续时间大约为 290 年。
// md5:4c8441d90787436e
# <翻译结束>


<原文开始>
// A Month specifies a month of the year (January = 1, ...).
<原文结束>

# <翻译开始>
// Month 指定一年中的月份（一月 = 1，...） 翻译备注:这两个不翻译,常用引用了这个,会出错. md5:c1896873d737e487
# <翻译结束>


<原文开始>
// A Weekday specifies a day of the week (Sunday = 0, ...).
<原文结束>

# <翻译开始>
// Weekday 指定了一周中的某一天（周日=0，...） 翻译备注:这两个不翻译,常用引用了这个,会出错. md5:e7c6df3685894251
# <翻译结束>

