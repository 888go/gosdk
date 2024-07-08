package fmt

import (
	"fmt"
	"io"
)

// ScanState 表示传递给自定义扫描器的状态。扫描器可以进行逐字符扫描，或者询问 ScanState 以发现下一个由空格分隔的标记。
// md5:6e60fce9dca4739d
//type ScanState struct {
//	F fmt.ScanState
//} //md5:548d40c7bd78ce51

// Scanner 接口由任何具有 Scan 方法的值实现，该方法扫描输入以获取值的表示形式，并将结果存储在接收器中，为了有用，接收器必须是一个指针。对于实现了 Scanner 接口的任何 Scan、Scanf 或 Scanln 的参数，都会调用 Scan 方法。
// md5:273a12bd045b3428
//type Scanner struct {
//	F fmt.Scanner
//} //md5:ebff7c39f9497cbd

//const (
//	binaryDigits      = "01"
//	octalDigits       = "01234567"
//	decimalDigits     = "0123456789"
//	hexadecimalDigits = "0123456789aAbBcCdDeEfF"
//	sign              = "+-"
//	period            = "."
//	exponent          = "eEpP"
//)// md5:fc6f1a449642fbfc

// Scan 从标准输入扫描文本，将连续的空格分隔值存储到连续的参数中。换行符被视为空格。它返回成功扫描的项目数。如果这个数字小于参数的数量，err将报告原因。
// md5:bd4dc5cc6ece5796
func Scan(a ...any) (n int, err error) { //md5:74dd83898b33025f
	return fmt.Scan(a...)
}

// Scanln 与 Scan 类似，但是在遇到换行符时会停止扫描，
// 并且在最后一个项目之后必须有换行符或 EOF。
// md5:372edec2e2e5d8af
func Scanln(a ...any) (n int, err error) { //md5:1db5e088af11db20
	return fmt.Scanln(a...)
}

// Scanf 从标准输入读取文本，按照指定的格式将连续的空格分隔值存储到相应的参数中。
// 它返回成功扫描的项目数。如果这个数量小于参数的数量，err 将报告原因。
// 输入中的换行符必须与格式中的换行符匹配。
// 有一个例外： verbs `%c` 总是扫描输入中的下一个字符，即使它是空格（或制表符等）或换行符。
// md5:4ba6b79df6c03793
func Scanf(format string, a ...any) (n int, err error) { //md5:068a4471effc676f
	return fmt.Scanf(format, a...)
}

// Sscan 从参数字符串中扫描连续的空格分隔的值，将它们存储到相继的参数中。换行符被视为空格。它返回成功扫描的项目数量。如果这个数量少于参数的数量，err将报告原因。
// md5:c0024cfc4a1ed5a9
func Sscan(str string, a ...any) (n int, err error) { //md5:7088329337523b63
	return fmt.Sscan(str, a...)
}

// Sscanln类似于Sscan，但会在遇到换行符时停止扫描，并且在最后一个项目之后必须有一个换行符或文件结束。
// md5:c499c724c1b91c48
func Sscanln(str string, a ...any) (n int, err error) { //md5:8f83ae0440512197
	return fmt.Sscanln(str, a...)
}

// Sscanf 函数根据格式解析参数字符串，将连续的空格分隔值依次存储到后续的参数中。它返回成功解析的项目数量。
// 输入中的换行符必须与格式中的换行符匹配。
// md5:761d1a51af0c9d81
func Sscanf(str string, format string, a ...any) (n int, err error) { //md5:1eca9e23b2ed8572
	return fmt.Sscanf(str, format, a...)
}

// Fscan 从 r 中读取文本，将连续的空格分隔的值依次存储到后续的参数中。换行符被视为空格。它返回成功扫描的项目数。如果这个数量小于参数的数量，err 将报告失败的原因。
// md5:d7a2e5b032c871d0
func Fscan(r io.Reader, a ...any) (n int, err error) { //md5:42967ce1be4193c0
	return fmt.Fscan(r, a...)
}

// Fscanln类似于Fscan，但在遇到换行符时停止扫描。在最后一个项目之后，必须有换行符或文件结束（EOF）。
// md5:e547e532977d7b36
func Fscanln(r io.Reader, a ...any) (n int, err error) { //md5:2f2632f681df7704
	return fmt.Fscanln(r, a...)
}

// Fscanf 从 r 读取文本，根据格式将连续的空格分隔的值存储到相继的参数中。它返回成功解析的项数。输入中的换行符必须与格式中的换行符匹配。
// md5:9779449977491090
func Fscanf(r io.Reader, format string, a ...any) (n int, err error) { //md5:3b3a67d5e3089bc5
	return fmt.Fscanf(r, format, a...)
}
