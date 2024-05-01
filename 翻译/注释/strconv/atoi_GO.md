
<原文开始>
// A NumError records a failed conversion.
<原文结束>

# <翻译开始>
// NumError 记录了一次转换失败的情况。. md5:4bc87fcc15d03a84
# <翻译结束>


<原文开始>
// IntSize is the size in bits of an int or uint value.
<原文结束>

# <翻译开始>
// IntSize 是int或uint值的位数。. md5:c5f70ea31d120aca
# <翻译结束>


<原文开始>
// ErrRange indicates that a value is out of range for the target type.
<原文结束>

# <翻译开始>
// ErrRange表示值超出了目标类型的范围。. md5:f2114ade03b9c908
# <翻译结束>


<原文开始>
// ErrSyntax indicates that a value does not have the right syntax for the target type.
<原文结束>

# <翻译开始>
// ErrSyntax表示值的语法不符合目标类型的要求。. md5:a7b31dd1093b5f09
# <翻译结束>


<原文开始>
// ParseUint is like ParseInt but for unsigned numbers.
//
// A sign prefix is not permitted.
<原文结束>

# <翻译开始>
// ParseUint 与 ParseInt 类似，但用于无符号数字。
//
// 不允许有符号前缀。
// md5:46e1e2e7065dee31
# <翻译结束>


<原文开始>
// ParseInt interprets a string s in the given base (0, 2 to 36) and
// bit size (0 to 64) and returns the corresponding value i.
//
// The string may begin with a leading sign: "+" or "-".
//
// If the base argument is 0, the true base is implied by the string's
// prefix following the sign (if present): 2 for "0b", 8 for "0" or "0o",
// 16 for "0x", and 10 otherwise. Also, for argument base 0 only,
// underscore characters are permitted as defined by the Go syntax for
// [integer literals].
//
// The bitSize argument specifies the integer type
// that the result must fit into. Bit sizes 0, 8, 16, 32, and 64
// correspond to int, int8, int16, int32, and int64.
// If bitSize is below 0 or above 64, an error is returned.
//
// The errors that ParseInt returns have concrete type *NumError
// and include err.Num = s. If s is empty or contains invalid
// digits, err.Err = ErrSyntax and the returned value is 0;
// if the value corresponding to s cannot be represented by a
// signed integer of the given size, err.Err = ErrRange and the
// returned value is the maximum magnitude integer of the
// appropriate bitSize and sign.
//
// [integer literals]: https://go.dev/ref/spec#Integer_literals
<原文结束>

# <翻译开始>
// ParseInt 将给定基数（0、2 到 36）和位大小（0 到 64）的字符串 s 解析为相应的整数值 i。
//
// 字符串可能以正负号开头："+" 或 "-"。
//
// 如果基数参数为 0，真正的基数由符号后的字符串前缀决定：如果存在，"0b" 表示 2，"0" 或 "0o" 表示 8，"0x" 表示 16，否则表示 10。仅当基数为 0 时，允许使用下划线字符，遵循 Go 语言整数字面量的语法定义。
//
// 位大小参数指定结果必须适应的整数类型。位大小 0、8、16、32 和 64 分别对应于 int、int8、int16、int32 和 int64。
// 如果位大小低于 0 或超过 64，将返回一个错误。
//
// ParseInt 返回的错误具有具体类型 *NumError，并包含 err.Num = s。如果 s 空或包含无效的数字，err.Err = ErrSyntax，返回值为 0；如果 s 对应的值不能用给定大小的有符号整数表示，err.Err = ErrRange，返回值为相应位大小和符号的最大整数值。
//
// [整数字面量]：https://go.dev/ref/spec#Integer_literals
// md5:c7272ea65b20e967
# <翻译结束>


<原文开始>
// Atoi is equivalent to ParseInt(s, 10, 0), converted to type int.
<原文结束>

# <翻译开始>
// Atoi等同于ParseInt(s, 10, 0)，转换为int类型。. md5:53cc58b8acf08aaf
# <翻译结束>

