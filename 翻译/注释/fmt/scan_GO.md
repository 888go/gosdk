
<原文开始>
// Scan scans text read from standard input, storing successive
// space-separated values into successive arguments. Newlines count
// as space. It returns the number of items successfully scanned.
// If that is less than the number of arguments, err will report why.
<原文结束>

# <翻译开始>
// Scan 从标准输入扫描文本，将连续的空格分隔值存储到连续的参数中。换行符被视为空格。它返回成功扫描的项目数。如果这个数字小于参数的数量，err将报告原因。
// md5:bd4dc5cc6ece5796
# <翻译结束>


<原文开始>
// Scanln is similar to Scan, but stops scanning at a newline and
// after the final item there must be a newline or EOF.
<原文结束>

# <翻译开始>
// Scanln 与 Scan 类似，但是在遇到换行符时会停止扫描，
// 并且在最后一个项目之后必须有换行符或 EOF。
// md5:372edec2e2e5d8af
# <翻译结束>


<原文开始>
// Scanf scans text read from standard input, storing successive
// space-separated values into successive arguments as determined by
// the format. It returns the number of items successfully scanned.
// If that is less than the number of arguments, err will report why.
// Newlines in the input must match newlines in the format.
// The one exception: the verb %c always scans the next rune in the
// input, even if it is a space (or tab etc.) or newline.
<原文结束>

# <翻译开始>
// Scanf 从标准输入读取文本，按照指定的格式将连续的空格分隔值存储到相应的参数中。
// 它返回成功扫描的项目数。如果这个数量小于参数的数量，err 将报告原因。
// 输入中的换行符必须与格式中的换行符匹配。
// 有一个例外： verbs `%c` 总是扫描输入中的下一个字符，即使它是空格（或制表符等）或换行符。
// md5:4ba6b79df6c03793
# <翻译结束>


<原文开始>
// Sscan scans the argument string, storing successive space-separated
// values into successive arguments. Newlines count as space. It
// returns the number of items successfully scanned. If that is less
// than the number of arguments, err will report why.
<原文结束>

# <翻译开始>
// Sscan 从参数字符串中扫描连续的空格分隔的值，将它们存储到相继的参数中。换行符被视为空格。它返回成功扫描的项目数量。如果这个数量少于参数的数量，err将报告原因。
// md5:c0024cfc4a1ed5a9
# <翻译结束>


<原文开始>
// Sscanln is similar to Sscan, but stops scanning at a newline and
// after the final item there must be a newline or EOF.
<原文结束>

# <翻译开始>
// Sscanln类似于Sscan，但会在遇到换行符时停止扫描，并且在最后一个项目之后必须有一个换行符或文件结束。
// md5:c499c724c1b91c48
# <翻译结束>


<原文开始>
// Sscanf scans the argument string, storing successive space-separated
// values into successive arguments as determined by the format. It
// returns the number of items successfully parsed.
// Newlines in the input must match newlines in the format.
<原文结束>

# <翻译开始>
// Sscanf 函数根据格式解析参数字符串，将连续的空格分隔值依次存储到后续的参数中。它返回成功解析的项目数量。
// 输入中的换行符必须与格式中的换行符匹配。
// md5:761d1a51af0c9d81
# <翻译结束>


<原文开始>
// Fscan scans text read from r, storing successive space-separated
// values into successive arguments. Newlines count as space. It
// returns the number of items successfully scanned. If that is less
// than the number of arguments, err will report why.
<原文结束>

# <翻译开始>
// Fscan 从 r 中读取文本，将连续的空格分隔的值依次存储到后续的参数中。换行符被视为空格。它返回成功扫描的项目数。如果这个数量小于参数的数量，err 将报告失败的原因。
// md5:d7a2e5b032c871d0
# <翻译结束>


<原文开始>
// Fscanln is similar to Fscan, but stops scanning at a newline and
// after the final item there must be a newline or EOF.
<原文结束>

# <翻译开始>
// Fscanln类似于Fscan，但在遇到换行符时停止扫描。在最后一个项目之后，必须有换行符或文件结束（EOF）。
// md5:e547e532977d7b36
# <翻译结束>


<原文开始>
// Fscanf scans text read from r, storing successive space-separated
// values into successive arguments as determined by the format. It
// returns the number of items successfully parsed.
// Newlines in the input must match newlines in the format.
<原文结束>

# <翻译开始>
// Fscanf 从 r 读取文本，根据格式将连续的空格分隔的值存储到相继的参数中。它返回成功解析的项数。输入中的换行符必须与格式中的换行符匹配。
// md5:9779449977491090
# <翻译结束>


<原文开始>
// ScanState represents the scanner state passed to custom scanners.
// Scanners may do rune-at-a-time scanning or ask the ScanState
// to discover the next space-delimited token.
<原文结束>

# <翻译开始>
// ScanState 表示传递给自定义扫描器的状态。扫描器可以进行逐字符扫描，或者询问 ScanState 以发现下一个由空格分隔的标记。
// md5:6e60fce9dca4739d
# <翻译结束>


<原文开始>
// Scanner is implemented by any value that has a Scan method, which scans
// the input for the representation of a value and stores the result in the
// receiver, which must be a pointer to be useful. The Scan method is called
// for any argument to Scan, Scanf, or Scanln that implements it.
<原文结束>

# <翻译开始>
// Scanner 接口由任何具有 Scan 方法的值实现，该方法扫描输入以获取值的表示形式，并将结果存储在接收器中，为了有用，接收器必须是一个指针。对于实现了 Scanner 接口的任何 Scan、Scanf 或 Scanln 的参数，都会调用 Scan 方法。
// md5:273a12bd045b3428
# <翻译结束>

