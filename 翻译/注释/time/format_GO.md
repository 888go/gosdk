
<原文开始>
// ParseError describes a problem parsing a time string.
<原文结束>

# <翻译开始>
// ParseError描述了解析时间字符串时遇到的问题。. md5:0ce67e9011eaba3c
# <翻译结束>


<原文开始>
// The reference time, in numerical order.
<原文结束>

# <翻译开始>
// 参考时间，按数值顺序排列。. md5:8a23f5425fc797b4
# <翻译结束>


<原文开始>
// RFC822 with numeric zone
<原文结束>

# <翻译开始>
// 根据RFC822标准，带有数字时区的格式. md5:4f5e9fc6acf09697
# <翻译结束>


<原文开始>
// RFC1123 with numeric zone
<原文结束>

# <翻译开始>
// 符合RFC1123的数字时区. md5:648451dc72784dbf
# <翻译结束>


<原文开始>
// ".0", ".00", ... , trailing zeros included
<原文结束>

# <翻译开始>
// 包含小数点后的零，如".0", ".00", .... md5:095a4186ac5b72d9
# <翻译结束>


<原文开始>
// ".9", ".99", ..., trailing zeros omitted
<原文结束>

# <翻译开始>
// ".9", ".99", ...，末尾的零被省略. md5:6c1cc814ca718990
# <翻译结束>


<原文开始>
// need hour, minute, second
<原文结束>

# <翻译开始>
// 需要小时，分钟，秒. md5:c7247cac348c3f92
# <翻译结束>


<原文开始>
// extra argument in high bits, above low stdArgShift
<原文结束>

# <翻译开始>
// 高位的额外参数，高于低标准的ArgShift. md5:0d68305d152f8183
# <翻译结束>


<原文开始>
// extra argument in high 4 bits for fractional second separators
<原文结束>

# <翻译开始>
// 高四位中的额外参数用于小数点分隔符. md5:67cb7129ebe3bc2c
# <翻译结束>


<原文开始>
// GoString implements fmt.GoStringer and formats t to be printed in Go source
// code.
<原文结束>

# <翻译开始>
// GoString 实现了 fmt.GoStringer 接口，用于将 t 格式化为可在 Go 源代码中打印的形式。
// md5:7a33ba38e316bbcb
# <翻译结束>


<原文开始>
// Format returns a textual representation of the time value formatted according
// to the layout defined by the argument. See the documentation for the
// constant called Layout to see how to represent the layout format.
//
// The executable example for Time.Format demonstrates the working
// of the layout string in detail and is a good reference.
<原文结束>

# <翻译开始>
// Format 根据提供的格式参数返回时间值的文本表示。
// 查看名为 Layout 的常量文档，了解如何表示布局格式。
//
// Time.Format 的可执行示例详细演示了布局字符串的工作原理，是一个很好的参考。
// md5:0e6869ec07aa8837
# <翻译结束>


<原文开始>
// AppendFormat is like Format but appends the textual
// representation to b and returns the extended buffer.
<原文结束>

# <翻译开始>
// AppendFormat 类似于 Format，但它将文本表示追加到 b 中，并返回扩展的缓冲区。
// md5:070a0e7c5343b175
# <翻译结束>


<原文开始>
// Error returns the string representation of a ParseError.
<原文结束>

# <翻译开始>
// Error 返回一个 ParseError 的字符串表示。. md5:c5730859166f42e5
# <翻译结束>


<原文开始>
// Parse parses a formatted string and returns the time value it represents.
// See the documentation for the constant called Layout to see how to
// represent the format. The second argument must be parseable using
// the format string (layout) provided as the first argument.
//
// The example for Time.Format demonstrates the working of the layout string
// in detail and is a good reference.
//
// When parsing (only), the input may contain a fractional second
// field immediately after the seconds field, even if the layout does not
// signify its presence. In that case either a comma or a decimal point
// followed by a maximal series of digits is parsed as a fractional second.
// Fractional seconds are truncated to nanosecond precision.
//
// Elements omitted from the layout are assumed to be zero or, when
// zero is impossible, one, so parsing "3:04pm" returns the time
// corresponding to Jan 1, year 0, 15:04:00 UTC (note that because the year is
// 0, this time is before the zero Time).
// Years must be in the range 0000..9999. The day of the week is checked
// for syntax but it is otherwise ignored.
//
// For layouts specifying the two-digit year 06, a value NN >= 69 will be treated
// as 19NN and a value NN < 69 will be treated as 20NN.
//
// The remainder of this comment describes the handling of time zones.
//
// In the absence of a time zone indicator, Parse returns a time in UTC.
//
// When parsing a time with a zone offset like -0700, if the offset corresponds
// to a time zone used by the current location (Local), then Parse uses that
// location and zone in the returned time. Otherwise it records the time as
// being in a fabricated location with time fixed at the given zone offset.
//
// When parsing a time with a zone abbreviation like MST, if the zone abbreviation
// has a defined offset in the current location, then that offset is used.
// The zone abbreviation "UTC" is recognized as UTC regardless of location.
// If the zone abbreviation is unknown, Parse records the time as being
// in a fabricated location with the given zone abbreviation and a zero offset.
// This choice means that such a time can be parsed and reformatted with the
// same layout losslessly, but the exact instant used in the representation will
// differ by the actual zone offset. To avoid such problems, prefer time layouts
// that use a numeric zone offset, or use ParseInLocation.
<原文结束>

# <翻译开始>
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
# <翻译结束>


<原文开始>
// ParseInLocation is like Parse but differs in two important ways.
// First, in the absence of time zone information, Parse interprets a time as UTC;
// ParseInLocation interprets the time as in the given location.
// Second, when given a zone offset or abbreviation, Parse tries to match it
// against the Local location; ParseInLocation uses the given location.
<原文结束>

# <翻译开始>
// ParseInLocation 与 Parse 类似，但在两个重要方面有所不同。
// 首先，在没有时区信息的情况下，Parse 将时间解释为 UTC；而 ParseInLocation 则将时间解释为给定的时区。
// 其次，当给定时区偏移量或缩写时，Parse 会尝试将其与 Local 时区匹配；而 ParseInLocation 使用给定的时区。
// md5:23556e86126104cc
# <翻译结束>


<原文开始>
// String returns the time formatted using the format string
//
//	"2006-01-02 15:04:05.999999999 -0700 MST"
//
// If the time has a monotonic clock reading, the returned string
// includes a final field "m=±<value>", where value is the monotonic
// clock reading formatted as a decimal number of seconds.
//
// The returned string is meant for debugging; for a stable serialized
// representation, use t.MarshalText, t.MarshalBinary, or t.Format
// with an explicit format string.
<原文结束>

# <翻译开始>
// String 方法返回使用格式字符串格式化的时间。
//
//	格式字符串为："2006-01-02 15:04:05.999999999 -0700 MST"
//
// 如果时间包含单调时钟读数，返回的字符串会附加一个最后的字段 "m=±<值>"，其中 <值> 是单调时钟读数以秒为单位的十进制数表示。
//
// 返回的字符串主要用于调试；若需要稳定的序列化表示，请使用 t.MarshalText、t.MarshalBinary 或者 t.Format 并提供明确的格式字符串。
// md5:534437fd6587d3d7
# <翻译结束>

