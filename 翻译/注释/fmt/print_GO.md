
<原文开始>
// FormatString returns a string representing the fully qualified formatting
// directive captured by the State, followed by the argument verb. (State does not
// itself contain the verb.) The result has a leading percent sign followed by any
// flags, the width, and the precision. Missing flags, width, and precision are
// omitted. This function allows a Formatter to reconstruct the original
// directive triggering the call to Format.
<原文结束>

# <翻译开始>
// FormatString 返回一个字符串，表示 State 所捕获的完整格式化指令，后面跟着参数动词。（State 自身不包含动词。）结果以百分号开头，后面跟着任何标志、宽度和精度。缺失的标志、宽度和精度将被省略。这个函数允许 Formatter 重新构建引发 Format 调用的原始指令。
// md5:8c85abb2482198f7
# <翻译结束>


<原文开始>
// Fprintf formats according to a format specifier and writes to w.
// It returns the number of bytes written and any write error encountered.
<原文结束>

# <翻译开始>
// Fprintf 根据格式说明符进行格式化，并将结果写入到 w。它返回写入的字节数和遇到的任何写入错误。
// md5:11dd19b9f67ba7d6
# <翻译结束>


<原文开始>
// Printf formats according to a format specifier and writes to standard output.
// It returns the number of bytes written and any write error encountered.
<原文结束>

# <翻译开始>
// Printf 根据格式说明符进行格式化，并将结果写入标准输出。
// 它返回写入的字节数以及遇到的任何写入错误。
// md5:54ac01abf5326255
# <翻译结束>


<原文开始>
// Sprintf formats according to a format specifier and returns the resulting string.
<原文结束>

# <翻译开始>
// Sprintf 根据指定的格式规范格式化数据，并返回结果字符串。. md5:d61280ff301bb2b2
# <翻译结束>


<原文开始>
// Appendf formats according to a format specifier, appends the result to the byte
// slice, and returns the updated slice.
<原文结束>

# <翻译开始>
// Appendf 根据格式说明符进行格式化，将结果追加到字节切片中，并返回更新后的切片。
// md5:393dbb1b605c68f1
# <翻译结束>


<原文开始>
// Fprint formats using the default formats for its operands and writes to w.
// Spaces are added between operands when neither is a string.
// It returns the number of bytes written and any write error encountered.
<原文结束>

# <翻译开始>
// Fprint使用其操作数的默认格式进行格式化，并将结果写入到w。当两个操作数都不是字符串时，会在它们之间添加空格。它返回写入的字节数和遇到的任何写入错误。
// md5:ddee1c503e0a8436
# <翻译结束>


<原文开始>
// Print formats using the default formats for its operands and writes to standard output.
// Spaces are added between operands when neither is a string.
// It returns the number of bytes written and any write error encountered.
<原文结束>

# <翻译开始>
// Print 使用其操作数的默认格式进行格式化，并写入标准输出。
// 当两个操作数都不是字符串时，在它们之间添加空格。
// 它返回写入的字节数以及遇到的任何写入错误。
// md5:54b0b31967437408
# <翻译结束>


<原文开始>
// Sprint formats using the default formats for its operands and returns the resulting string.
// Spaces are added between operands when neither is a string.
<原文结束>

# <翻译开始>
// Sprint 使用默认格式化方式处理其操作数，并返回结果字符串。
// 当没有操作数是字符串时，会在它们之间添加空格。
// md5:bb92a3c2cd3c2853
# <翻译结束>


<原文开始>
// Append formats using the default formats for its operands, appends the result to
// the byte slice, and returns the updated slice.
<原文结束>

# <翻译开始>
// 使用其操作数的默认格式进行格式化，将结果追加到字节切片中，并返回更新后的切片。
// md5:1b14355a42fa90e0
# <翻译结束>


<原文开始>
// Fprintln formats using the default formats for its operands and writes to w.
// Spaces are always added between operands and a newline is appended.
// It returns the number of bytes written and any write error encountered.
<原文结束>

# <翻译开始>
// Fprintln使用其操作数的默认格式进行格式化，并将结果写入到w。在操作数之间始终会添加空格，最后还会添加一个换行符。它返回写入的字节数和遇到的任何写入错误。
// md5:54533f8543d392a7
# <翻译结束>


<原文开始>
// Println formats using the default formats for its operands and writes to standard output.
// Spaces are always added between operands and a newline is appended.
// It returns the number of bytes written and any write error encountered.
<原文结束>

# <翻译开始>
// Println 使用其操作数的默认格式进行格式化，并写入标准输出。
// 操作数之间始终会添加空格，并附加一个换行符。
// 它返回写入的字节数以及遇到的任何写入错误。
// md5:2efc584ded7df262
# <翻译结束>


<原文开始>
// Sprintln formats using the default formats for its operands and returns the resulting string.
// Spaces are always added between operands and a newline is appended.
<原文结束>

# <翻译开始>
// Sprintln 使用其操作数的默认格式进行格式化，并返回结果字符串。
// 操作数之间总是添加空格，并在末尾追加一个换行符。
// md5:36c0bd5a115ae1c5
# <翻译结束>


<原文开始>
// Appendln formats using the default formats for its operands, appends the result
// to the byte slice, and returns the updated slice. Spaces are always added
// between operands and a newline is appended.
<原文结束>

# <翻译开始>
// Appendln使用其操作数的默认格式进行格式化，将结果追加到字节切片中，并返回更新后的切片。始终在操作数之间添加空格，并追加一个换行符。
// md5:ca307513588203d4
# <翻译结束>


<原文开始>
// State represents the printer state passed to custom formatters.
// It provides access to the io.Writer interface plus information about
// the flags and options for the operand's format specifier.
<原文结束>

# <翻译开始>
// State 表示传递给自定义格式化程序的打印机状态。
// 它提供了对 io.Writer 接口的访问，以及有关操作数格式说明符的标志和选项的信息。
// md5:d0bc7f1421fd6664
# <翻译结束>


<原文开始>
// Formatter is implemented by any value that has a Format method.
// The implementation controls how State and rune are interpreted,
// and may call Sprint() or Fprint(f) etc. to generate its output.
<原文结束>

# <翻译开始>
// Formatter 是任何具有 Format 方法的值所实现的接口。
// 实现方式控制了如何解释 State 和 rune，
// 并且可以调用 Sprint() 或 Fprint(f) 等方法来生成输出。
// md5:fce744c8d7dbc7b2
# <翻译结束>


<原文开始>
// GoStringer is implemented by any value that has a GoString method,
// which defines the Go syntax for that value.
// The GoString method is used to print values passed as an operand
// to a %#v format.
<原文结束>

# <翻译开始>
// GoStringer 接口由任何实现了 GoString 方法的值实现，该方法定义了该值在 Go 语法中的表示。
// GoString 方法用于打印作为 %#v 格式参数传递的值。
// md5:4fe42f9ac826c032
# <翻译结束>

