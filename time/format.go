package time

import "time"

// ParseError描述了解析时间字符串时遇到的问题。. md5:0ce67e9011eaba3c
// 翻译提示:type 解析错误 struct {
//     时间解析错误 time.ParseError
// }
// 
// func (e 解析错误) Error() string { 
//     // 返回错误信息，这里假设原始的Error方法保持不变，因为它是英文的标准库方法
//     return e.时间解析错误.Error()
// }
type ParseError struct {
	F time.ParseError
} //md5:6c5710d3aad84515

// 翻译提示:const (
// 	参考时间Layout      = "01/02 03:04:05下午 '06 -0700" // 参考时间格式，按数值顺序排列?? md5:8a23f5425fc797b4
// 	ANSIC时间       = "周一 一月 _2 15:04:05 2006"
// 	Unix时间    = "周一 一月 _2 15:04:05 MST 2006"
// 	Ruby时间    = "周一 一月 02 15:04:05 -0700 2006"
// 	RFC822时间      = "02 一月 06 15:04 MST"
// 	RFC822Z时间     = "02 一月 06 15:04 -0700" // 根据RFC822标准，包含数字时区的格式. md5:4f5e9fc6acf09697
// 	RFC850时间      = "星期一, 02-一月-06 15:04:05 MST"
// 	RFC1123时间     = "周一, 02 一月 2006 15:04:05 MST"
// 	RFC1123Z时间    = "周一, 02 一月 2006 15:04:05 -0700" // 符合RFC1123的数字时区格式. md5:648451dc72784dbf
// 	RFC3339时间     = "2006-01-02T15:04:05Z07:00"
// 	RFC3339Nano时间 = "2006-01-02T15:04:05.999999999Z07:00"
// 	厨房时间         = "下午3:04"
// 	// 实用的时间戳。
// 	时间戳      = "一月 _2 15:04:05"
// 	时间戳毫秒  = "一月 _2 15:04:05.000"
// 	时间戳微秒  = "一月 _2 15:04:05.000000"
// 	时间戳纳秒  = "一月 _2 15:04:05.000000000"
// 	日期时间     = "2006-01-02 15:04:05"
// 	日期         = "2006-01-02"
// 	时间         = "15:04:05"
// )
const (
	Layout      = "01/02 03:04:05PM '06 -0700" // 参考时间，按数值顺序排列。. md5:8a23f5425fc797b4
	ANSIC       = "Mon Jan _2 15:04:05 2006"
	UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
	RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
	RFC822      = "02 Jan 06 15:04 MST"
	RFC822Z     = "02 Jan 06 15:04 -0700" // 根据RFC822标准，带有数字时区的格式. md5:4f5e9fc6acf09697
	RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
	RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
	RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700" // 符合RFC1123的数字时区. md5:648451dc72784dbf
	RFC3339     = "2006-01-02T15:04:05Z07:00"
	RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
	Kitchen     = "3:04PM"
	// Handy time stamps.
	Stamp      = "Jan _2 15:04:05"
	StampMilli = "Jan _2 15:04:05.000"
	StampMicro = "Jan _2 15:04:05.000000"
	StampNano  = "Jan _2 15:04:05.000000000"
	DateTime   = "2006-01-02 15:04:05"
	DateOnly   = "2006-01-02"
	TimeOnly   = "15:04:05"
) // md5:0dd08e9a1e224ec7

// 翻译提示:const (
// 	_                                = iota
// 	标准完整月份名称                      = iota + 标准需要日期  // "January"
// 	月份缩写                                        // "Jan"
// 	数字月份                                        // "1"
// 	零填充月份                                      // "01"
// 	完整星期几名称                                 // "Monday"
// 	星期几缩写                                      // "Mon"
// 	日                                             // "2"
// 	下划线日                                       // "_2"
// 	零填充日                                       // "02"
// 	下划线年份日                                   // "__2"
// 	零填充年份日                                   // "002"
// 	小时                          = iota + 标准需要时钟 // "15"
// 	12小时制小时                                      // "3"
// 	零填充12小时制小时                               // "03"
// 	分钟                                             // "4"
// 	零填充分钟                                       // "04"
// 	秒                                               // "5"
// 	零填充秒                                         // "05"
// 	标准完整年份                      = iota + 标准需要日期  // "2006"
// 	年份缩写                                          // "06"
// 	下午AM/PM                         = iota + 标准需要时钟 // "PM"
// 	小写下午AM/PM                                       // "pm"
// 	时区缩写                                         // "MST"
// 	ISO8601时区缩写                                 // "Z0700"  // UTC打印为Z
// 	ISO8601带秒的时区缩写                           // "Z070000"
// 	ISO8601短时区缩写                               // "Z07"
// 	ISO8601带冒号的时区缩写                         // "Z07:00" // UTC打印为Z
// 	ISO8601带冒号和秒的时区缩写                    // "Z07:00:00"
// 	数值时区                                       // "-0700"  // 始终为数值
// 	数值带秒的时区                                 // "-070000"
// 	数值短时区                                    // "-07"    // 始终为数值
// 	数值带冒号的时区                               // "-07:00" // 始终为数值
// 	数值带冒号和秒的时区                           // "-07:00:00"
// 	包含小数点后零的秒                             // 如".0", ".00", .... md5:095a4186ac5b72d9
// 	小数点后不省略零的秒                           // ".9", ".99", ...，末尾的零不省略. md5:6c1cc814ca718990
// 
// 	标准需要日期         = 1 << 8             // 需要月份、日、年
// 	标准需要时钟         = 2 << 8             // 需要小时、分钟、秒. md5:c7247cac348c3f92
// 	额外参数位移         = 16                 // 高位的额外参数，高于低位的参数位移. md5:0d68305d152f8183
// 	分隔符位移           = 28                 // 高四位中的额外参数用于小数点分隔?? md5:67cb7129ebe3bc2c
// 	参数掩码             = 1<<标准参数位移 - 1 // 参数掩码
// )
const (
	_                        = iota
	stdLongMonth             = iota + stdNeedDate  // "January"
	stdMonth                                       // "Jan"
	stdNumMonth                                    // "1"
	stdZeroMonth                                   // "01"
	stdLongWeekDay                                 // "Monday"
	stdWeekDay                                     // "Mon"
	stdDay                                         // "2"
	stdUnderDay                                    // "_2"
	stdZeroDay                                     // "02"
	stdUnderYearDay                                // "__2"
	stdZeroYearDay                                 // "002"
	stdHour                  = iota + stdNeedClock // "15"
	stdHour12                                      // "3"
	stdZeroHour12                                  // "03"
	stdMinute                                      // "4"
	stdZeroMinute                                  // "04"
	stdSecond                                      // "5"
	stdZeroSecond                                  // "05"
	stdLongYear              = iota + stdNeedDate  // "2006"
	stdYear                                        // "06"
	stdPM                    = iota + stdNeedClock // "PM"
	stdpm                                          // "pm"
	stdTZ                    = iota                // "MST"
	stdISO8601TZ                                   // "Z0700"  // prints Z for UTC
	stdISO8601SecondsTZ                            // "Z070000"
	stdISO8601ShortTZ                              // "Z07"
	stdISO8601ColonTZ                              // "Z07:00" // prints Z for UTC
	stdISO8601ColonSecondsTZ                       // "Z07:00:00"
	stdNumTZ                                       // "-0700"  // always numeric
	stdNumSecondsTz                                // "-070000"
	stdNumShortTZ                                  // "-07"    // always numeric
	stdNumColonTZ                                  // "-07:00" // always numeric
	stdNumColonSecondsTZ                           // "-07:00:00"
	stdFracSecond0                                 // 包含小数点后的零，如".0", ".00", .... md5:095a4186ac5b72d9
	stdFracSecond9                                 // ".9", ".99", ...，末尾的零被省略. md5:6c1cc814ca718990

	stdNeedDate       = 1 << 8             // need month, day, year
	stdNeedClock      = 2 << 8             // 需要小时，分钟，秒. md5:c7247cac348c3f92
	stdArgShift       = 16                 // 高位的额外参数，高于低标准的ArgShift. md5:0d68305d152f8183
	stdSeparatorShift = 28                 // 高四位中的额外参数用于小数点分隔符. md5:67cb7129ebe3bc2c
	stdMask           = 1<<stdArgShift - 1 // mask out argument
) // md5:30d8e7706ca83b6b

//const (
//	lowerhex  = "0123456789abcdef"
//	runeSelf  = 0x80
//	runeError = '\uFFFD'
//)// md5:66ebe9c952b8908f

// String 方法返回使用格式字符串格式化的时间。
//
//	格式字符串为："2006-01-02 15:04:05.999999999 -0700 MST"
//
// 如果时间包含单调时钟读数，返回的字符串会附加一个最后的字段 "m=±<值>"，其中 <值> 是单调时钟读数以秒为单位的十进制数表示。
//
// 返回的字符串主要用于调试；若需要稳定的序列化表示，请使用 t.MarshalText、t.MarshalBinary 或者 t.Format 并提供明确的格式字符串。
// md5:534437fd6587d3d7
// 翻译提示:func (t 时间) 字符串() 字符串 {}
func (t Time) String() string { //md5:244060d240d2fba5
	return t.F.String()
}

// GoString 实现了 fmt.GoStringer 接口，用于将 t 格式化为可在 Go 源代码中打印的形式。
// md5:7a33ba38e316bbcb
// 翻译提示:func (t 时间) Go字符串() 字符串 {}
func (t Time) GoString() string { //md5:bbef06c0bba983b6
	return t.F.GoString()
}

// Format 根据提供的格式参数返回时间值的文本表示。
// 查看名为 Layout 的常量文档，了解如何表示布局格式。
//
// Time.Format 的可执行示例详细演示了布局字符串的工作原理，是一个很好的参考。
// md5:0e6869ec07aa8837
// 翻译提示:func (t 时间) 格式化(layout 字符串) 字符串 {}
func (t Time) Format(layout string) string { //md5:cd21a0becf8aba8a
	return t.F.Format(layout)
}

// AppendFormat 类似于 Format，但它将文本表示追加到 b 中，并返回扩展的缓冲区。
// md5:070a0e7c5343b175
// 翻译提示:func (t 时间) 添加格式化(b []字节, 格式字符串) []字节 {}
func (t Time) AppendFormat(b []byte, layout string) []byte { //md5:f1300a966a1e3949
	return t.F.AppendFormat(b, layout)
}

// Error 返回一个 ParseError 的字符串表示。. md5:c5730859166f42e5
// 翻译提示:func (e *解析错误) 错误() 字符串 {}
func (e *ParseError) Error() string { //md5:c2cf7f1bf647d37a
	return e.F.Error()
}

// Parse 解析格式化字符串并返回其表示的时间值。
// 有关如何表示格式的详细信息，请参阅名为 Layout 的常量的文档。
// 第二个参数必须能使用第一个参数提供的格式字符串（布局）进行解析。
//
// Time.Format 的示例详细展示了布局字符串的工作原理，是一个很好的参考。
//
// 在解析时（仅解析时），输入可能在秒字段之后立即包含一个分数秒字段，
// 即使布局字符串没有表明它的存在。在这种情况下，逗号或小数点后跟一串最大数量的数字
// 将被解析为分数秒。分数秒会被截断到纳秒精度。
//
// 布局中省略的元素默认为零，或者当零不可能时，默认为一，因此解析 "3:04pm"
// 返回对应于公元0年1月1日15:04:00 UTC的时间（请注意，因为年份是0，这个时间是在零时间之前）。
// 年份必须在0000..9999范围内。星期几会被检查语法但其他方面会被忽略。
//
// 对于指定两位数年份06的布局，值 NN >= 69 将被视为19NN，而 NN < 69 则被视为20NN。
//
// 以下剩余部分描述了时区的处理方式。
//
// 在缺少时区指示符的情况下，Parse 返回的是UTC时间。
//
// 当解析带有时区偏移（如-0700）的时间时，如果该偏移对应于当前位置（Local）所使用的时区，
// 那么Parse会使用该位置和时区在返回的时间中。否则，它会记录时间为一个具有给定时区偏移的虚构位置。
//
// 当解析带有时区缩写（如MST）的时间时，如果在当前位置该时区缩写的偏移已定义，则使用该偏移。
// 时区缩写"UTC"无论位置如何都会被识别为UTC。如果时区缩写未知，Parse会将时间记录为具有给定时区缩写和零偏移的虚构位置。
// 这种选择意味着这样的时间可以被解析并使用相同的布局无损地重新格式化，但在表示中使用的确切瞬间会根据实际时区偏移有所不同。
// 为了避免此类问题，建议优先使用包含数值时区偏移的时区布局，或使用ParseInLocation。
// md5:af59c84e115e04f1
// 翻译提示:func 解析布局(layout 字符串, value 字符串) (时间, 错误) {}
func Parse(layout, value string) (Time, error) { //md5:04f0e938ce9833d8
	t, err := time.Parse(layout, value)
	return Time{t}, err
}

// ParseInLocation 与 Parse 类似，但在两个重要方面有所不同。
// 首先，在没有时区信息的情况下，Parse 将时间解释为 UTC；而 ParseInLocation 则将时间解释为给定的时区。
// 其次，当给定时区偏移量或缩写时，Parse 会尝试将其与 Local 时区匹配；而 ParseInLocation 使用给定的时区。
// md5:23556e86126104cc
// 翻译提示:func 解析在时区(layout, 值 string, 时区 *地理位置) (时间 Time, 错误 error) {}
func ParseInLocation(layout, value string, loc *Location) (Time, error) { //md5:72c26b70b1a8980b
	t, err := time.ParseInLocation(layout, value, &loc.F)
	return Time{t}, err
}

// ParseDuration parses a duration string.
// A duration string is a possibly signed sequence of
// decimal numbers, each with optional fraction and a unit suffix,
// such as "300ms", "-1.5h" or "2h45m".
// Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
// 翻译提示:func 解析持续时间(s 字符串) (持续时间, 错误) {}
func ParseDuration(s string) (Duration, error) { //md5:2817ced6ba396e3a
	d, err := time.ParseDuration(s)
	return Duration(d), err
}
