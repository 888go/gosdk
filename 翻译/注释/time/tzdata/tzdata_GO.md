
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
// UTC represents Universal Coordinated Time (UTC).
<原文结束>

# <翻译开始>
// UTC 表示协调世界时间（Universal Coordinated Time）。. md5:991e0d09aaa17cec
# <翻译结束>


<原文开始>
// Local represents the system's local time zone.
// On Unix systems, Local consults the TZ environment
// variable to find the time zone to use. No TZ means
// use the system default /etc/localtime.
// TZ="" means use UTC.
// TZ="foo" means use file foo in the system timezone directory.
<原文结束>

# <翻译开始>
// Local 表示系统的本地时区。
// 在Unix系统上，Local会查询TZ环境变量来找到要使用的时区。
// 没有TZ意味着使用系统默认的/etc/localtime。
// TZ="" 意味着使用UTC时间。
// TZ="foo" 意味着在系统时区目录中使用名为foo的文件。
// md5:38f2b49de5174f53
# <翻译结束>

