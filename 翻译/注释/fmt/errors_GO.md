
<原文开始>
// Errorf formats according to a format specifier and returns the string as a
// value that satisfies error.
//
// If the format specifier includes a %w verb with an error operand,
// the returned error will implement an Unwrap method returning the operand.
// If there is more than one %w verb, the returned error will implement an
// Unwrap method returning a []error containing all the %w operands in the
// order they appear in the arguments.
// It is invalid to supply the %w verb with an operand that does not implement
// the error interface. The %w verb is otherwise a synonym for %v.
<原文结束>

# <翻译开始>
// Errorf 根据格式说明符进行格式化，并将字符串作为满足 error 接口的值返回。
//
// 如果格式说明符中包含带有错误操作数的 %w 动词，那么返回的错误将实现一个 Unwrap 方法，该方法返回该操作数。
// 如果有多个 %w 动词，返回的错误将实现一个 Unwrap 方法，该方法返回一个 []error，其中包含所有 %w 操作数，这些操作数按照它们在参数中出现的顺序排列。
// 使用 %w 动词提供一个不实现错误接口的操作数是无效的。除此外，%w 动词与 %v 是同义的。
// md5:32b0316eed56798a
# <翻译结束>

