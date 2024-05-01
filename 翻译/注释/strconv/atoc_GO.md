
<原文开始>
// ParseComplex converts the string s to a complex number
// with the precision specified by bitSize: 64 for complex64, or 128 for complex128.
// When bitSize=64, the result still has type complex128, but it will be
// convertible to complex64 without changing its value.
//
// The number represented by s must be of the form N, Ni, or N±Ni, where N stands
// for a floating-point number as recognized by ParseFloat, and i is the imaginary
// component. If the second N is unsigned, a + sign is required between the two components
// as indicated by the ±. If the second N is NaN, only a + sign is accepted.
// The form may be parenthesized and cannot contain any spaces.
// The resulting complex number consists of the two components converted by ParseFloat.
//
// The errors that ParseComplex returns have concrete type *NumError
// and include err.Num = s.
//
// If s is not syntactically well-formed, ParseComplex returns err.Err = ErrSyntax.
//
// If s is syntactically well-formed but either component is more than 1/2 ULP
// away from the largest floating point number of the given component's size,
// ParseComplex returns err.Err = ErrRange and c = ±Inf for the respective component.
<原文结束>

# <翻译开始>
// ParseComplex 函数将字符串 s 转换为复数，其精度由 bitSize 指定：64 表示 complex64，128 表示 complex128。
// 当 bitSize=64 时，结果的类型仍为 complex128，但它可以无损地转换为 complex64。
//
// s 表示的数字必须形如 N、Ni 或 N±Ni，其中 N 代表一个通过 ParseFloat 识别的浮点数，i 表示虚部。
// 如果第二个 N 是正数，则在两个部分之间需要有一个 + 号，如同 ± 所示。如果第二个 N 是 NaN，则只接受 + 号。
// 形式上可以使用括号包围且不能包含任何空格。
// 生成的复数由两个通过 ParseFloat 转换的组件组成。
//
// ParseComplex 返回的错误具有具体类型 *NumError，并包含 err.Num = s。
//
// 如果 s 语法不正确，ParseComplex 返回 err.Err = ErrSyntax。
//
// 如果 s 语法正确，但任一组成部分超出给定大小浮点数最大值的 1/2 ULP 范围，
// ParseComplex 返回 err.Err = ErrRange，并将相应部分的 c 设为 ±Inf。
// md5:2d5101902291657e
# <翻译结束>

