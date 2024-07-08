package fmt

import (
	"fmt"
	"io"
)

// State 表示传递给自定义格式化程序的打印机状态。
// 它提供了对 io.Writer 接口的访问，以及有关操作数格式说明符的标志和选项的信息。
// md5:d0bc7f1421fd6664
//type State struct {
//	F fmt.State
//} //md5:f0407adf626b88fd

// Formatter 是任何具有 Format 方法的值所实现的接口。
// 实现方式控制了如何解释 State 和 rune，
// 并且可以调用 Sprint() 或 Fprint(f) 等方法来生成输出。
// md5:fce744c8d7dbc7b2
//type Formatter struct {
//	F fmt.Formatter
//} //md5:4ab256b2048f01dc

// Stringer 由具有String方法的任何值实现，
// 它定义了该值的“原生”格式。
// String方法用于将作为操作数传递的值打印到接受字符串的任何格式或打印到未格式化的打印机（如print）。
//type Stringer struct {
//	F fmt.Stringer
//} //md5:be8b1dc868dbc649

// GoStringer 接口由任何实现了 GoString 方法的值实现，该方法定义了该值在 Go 语法中的表示。
// GoString 方法用于打印作为 %#v 格式参数传递的值。
// md5:4fe42f9ac826c032
//type GoStringer struct {
//	F fmt.GoStringer
//} //md5:85ff5d99d1384523

// FormatString 返回一个字符串，表示 State 所捕获的完整格式化指令，后面跟着参数动词。（State 自身不包含动词。）结果以百分号开头，后面跟着任何标志、宽度和精度。缺失的标志、宽度和精度将被省略。这个函数允许 Formatter 重新构建引发 Format 调用的原始指令。
// md5:8c85abb2482198f7
//func FormatString(state fmt.State, verb rune) string { //md5:0a7c4dff05129806
//	return fmt.FormatString(state, verb)
//}

// Fprintf 根据格式说明符进行格式化，并将结果写入到 w。它返回写入的字节数和遇到的任何写入错误。
// md5:11dd19b9f67ba7d6
func Fprintf(w io.Writer, format string, a ...any) (n int, err error) { //md5:8994acd4ed68277e
	return fmt.Fprintf(w, format, a...)
}

// Printf 根据格式说明符进行格式化，并将结果写入标准输出。
// 它返回写入的字节数以及遇到的任何写入错误。
// md5:54ac01abf5326255
func Printf(format string, a ...any) (n int, err error) { //md5:32ce169e70f3a5fa
	return fmt.Printf(format, a...)
}

// Sprintf 根据指定的格式规范格式化数据，并返回结果字符串。. md5:d61280ff301bb2b2
func Sprintf(format string, a ...any) string { //md5:24724198127875e1
	return fmt.Sprintf(format, a...)
}

// Appendf 根据格式说明符进行格式化，将结果追加到字节切片中，并返回更新后的切片。
// md5:393dbb1b605c68f1
func Appendf(b []byte, format string, a ...any) []byte { //md5:f220524ac7a1fa6c
	return fmt.Appendf(b, format, a...)
}

// Fprint使用其操作数的默认格式进行格式化，并将结果写入到w。当两个操作数都不是字符串时，会在它们之间添加空格。它返回写入的字节数和遇到的任何写入错误。
// md5:ddee1c503e0a8436
func Fprint(w io.Writer, a ...any) (n int, err error) { //md5:e5ec71cd9b761315
	return fmt.Fprint(w, a...)
}

// Print 使用其操作数的默认格式进行格式化，并写入标准输出。
// 当两个操作数都不是字符串时，在它们之间添加空格。
// 它返回写入的字节数以及遇到的任何写入错误。
// md5:54b0b31967437408
func Print(a ...any) (n int, err error) { //md5:a607cfafb122e0ba
	return fmt.Print(a...)
}

// Sprint 使用默认格式化方式处理其操作数，并返回结果字符串。
// 当没有操作数是字符串时，会在它们之间添加空格。
// md5:bb92a3c2cd3c2853
func Sprint(a ...any) string { //md5:92c8691e66c394b4
	return fmt.Sprint(a...)
}

// 使用其操作数的默认格式进行格式化，将结果追加到字节切片中，并返回更新后的切片。
// md5:1b14355a42fa90e0
func Append(b []byte, a ...any) []byte { //md5:e9cd1020fcbaa6de
	return fmt.Append(b, a...)
}

// Fprintln使用其操作数的默认格式进行格式化，并将结果写入到w。在操作数之间始终会添加空格，最后还会添加一个换行符。它返回写入的字节数和遇到的任何写入错误。
// md5:54533f8543d392a7
func Fprintln(w io.Writer, a ...any) (n int, err error) { //md5:3f3d6b62cc1c26bf
	return fmt.Fprintln(w, a...)
}

// Println 使用其操作数的默认格式进行格式化，并写入标准输出。
// 操作数之间始终会添加空格，并附加一个换行符。
// 它返回写入的字节数以及遇到的任何写入错误。
// md5:2efc584ded7df262
func Println(a ...any) (n int, err error) { //md5:eb17a1f50e6abf59
	return fmt.Println(a...)
}

// Sprintln 使用其操作数的默认格式进行格式化，并返回结果字符串。
// 操作数之间总是添加空格，并在末尾追加一个换行符。
// md5:36c0bd5a115ae1c5
func Sprintln(a ...any) string { //md5:eb45268a6a270840
	return fmt.Sprintln(a...)
}

// Appendln使用其操作数的默认格式进行格式化，将结果追加到字节切片中，并返回更新后的切片。始终在操作数之间添加空格，并追加一个换行符。
// md5:ca307513588203d4
func Appendln(b []byte, a ...any) []byte { //md5:307da584946d7ba4
	return fmt.Appendln(b, a...)
}
