package strconv

import "strconv"

//const (
//	lowerhex = "0123456789abcdef"
//	upperhex = "0123456789ABCDEF"
//)// md5:89575fd411bf282e

// Quote 返回一个表示 s 的双引号包围的 Go 字符串字面量。返回的字符串使用 Go 的转义序列（\t, \n, \xFF, \u0100）来表示控制字符和非打印字符，这些字符根据 IsPrint 函数定义。
// md5:9c84d4d37e9e0c21

// ff:
// s:
func Quote(s string) string { //md5:720e1f755c129ec4
	return strconv.Quote(s)
}

// AppendQuote 将表示为 s 的双引号 Go 字符串字面量（由 Quote 生成）追加到 dst 并返回扩展后的缓冲区。
// md5:1e33d2c270caf40a

// ff:
// s:
// dst:
func AppendQuote(dst []byte, s string) []byte { //md5:cf37bac05af703ba
	return strconv.AppendQuote(dst, s)
}

// QuoteToASCII 返回一个双引号的 Go 字符串字面量，表示 s。返回的字符串使用 Go 转义序列（\t, \n, \xFF, \u0100）来表示非 ASCII 字符和根据 IsPrint 定义的非打印字符。
// md5:f464b38dde129de6

// ff:
// s:
func QuoteToASCII(s string) string { //md5:4b7b24b1728336e3
	return strconv.QuoteToASCII(s)
}

// AppendQuoteToASCII 将由 QuoteToASCII 生成的表示 s 的双引号 Go 字符串字面量追加到 dst，并返回扩展缓冲区。
// md5:332417ed389a29af

// ff:
// s:
// dst:
func AppendQuoteToASCII(dst []byte, s string) []byte { //md5:f8ada45efe15eb13
	return strconv.AppendQuoteToASCII(dst, s)
}

// QuoteToGraphic 返回一个表示 s 的双引号包围的 Go 字符串字面量。
// 返回的字符串中，Unicode 图形字符（根据 IsGraphic 定义）保持不变，
// 而非图形字符则使用 Go 的转义序列（如 \t, \n, \xFF, \u0100）进行表示。
// md5:797a44d52c17bfdf

// ff:
// s:
func QuoteToGraphic(s string) string { //md5:7f85e94d973176b3
	return strconv.QuoteToGraphic(s)
}

// AppendQuoteToGraphic 将使用QuoteToGraphic生成的表示s的双引号Go字符串字面量追加到dst中，并返回扩展后的缓冲区。
// md5:77b96f90cce2f943

// ff:
// s:
// dst:
func AppendQuoteToGraphic(dst []byte, s string) []byte { //md5:2deccb3bb6ccbfb3
	return strconv.AppendQuoteToGraphic(dst, s)
}

// QuoteRune 返回一个单引号的Go字符字面量，表示给定的rune。返回的字符串使用Go转义序列（\t, \n, \xFF, \u0100）来表示控制字符和根据IsPrint定义的不可打印字符。如果r不是一个有效的Unicode代码点，它将被解释为Unicode替换字符U+FFFD。
// md5:64a0cfc2843decf3

// ff:
// r:
func QuoteRune(r rune) string { //md5:b226c44e497fd3e3
	return strconv.QuoteRune(r)
}

// AppendQuoteRune在dst中追加一个由QuoteRune生成的单引号Go字符字面量，并返回扩展的缓冲区。
// md5:95ae8b0ba366d90a

// ff:
// r:
// dst:
func AppendQuoteRune(dst []byte, r rune) []byte { //md5:602282ed6520be78
	return strconv.AppendQuoteRune(dst, r)
}

// QuoteRuneToASCII 返回一个表示字符r的单引号包围的Go字符字面量。
// 返回的字符串使用Go的转义序列（\t, \n, \xFF, \u0100）来表示非ASCII字符和根据IsPrint定义的非打印字符。
// 如果r不是一个有效的Unicode代码点，则将其解释为Unicode替换字符U+FFFD。
// md5:4e79e3eef2dcf2ac

// ff:
// r:
func QuoteRuneToASCII(r rune) string { //md5:742e283e1b10790d
	return strconv.QuoteRuneToASCII(r)
}

// AppendQuoteRuneToASCII将表示字符的单引号Go字符字面量（由QuoteRuneToASCII生成）追加到dst，并返回扩展后的缓冲区。
// md5:4188657dadab9ed6

// ff:
// r:
// dst:
func AppendQuoteRuneToASCII(dst []byte, r rune) []byte { //md5:66508131668caa97
	return strconv.AppendQuoteRuneToASCII(dst, r)
}

// QuoteRuneToGraphic 返回一个单引号的 Go 字符字面量，表示该 rune。如果 rune 不是 IsGraphic 定义的 Unicode 图形字符，返回的字符串将使用 Go 转义序列（\t, \n, \xFF, \u0100）。
// 如果 r 不是一个有效的 Unicode 码点，它将被解释为 Unicode 替换字符 U+FFFD。
// md5:8348103cbd300c3c

// ff:
// r:
func QuoteRuneToGraphic(r rune) string { //md5:3e6e1e87514384b6
	return strconv.QuoteRuneToGraphic(r)
}

// AppendQuoteRuneToGraphic 将由 QuoteRuneToGraphic 生成的单引号 Go 字符字面量追加到 dst，并返回扩展的缓冲区。
// md5:7d58aaf5eca2b147

// ff:
// r:
// dst:
func AppendQuoteRuneToGraphic(dst []byte, r rune) []byte { //md5:9a05751c921c9257
	return strconv.AppendQuoteRuneToGraphic(dst, r)
}

// CanBackquote 报告字符串 s 是否可以保持不变地表示为单行反引号字符串，其中不包含除制表符以外的其他控制字符。
// md5:3e4f8ce026e7824c

// ff:
// s:
func CanBackquote(s string) bool { //md5:67370d10584f6df1
	return strconv.CanBackquote(s)
}

// UnquoteChar 解码由字符串 s 表示的转义字符串或字符字面量中的第一个字符或字节。
// 它返回四个值：
//
// 1. value，解码后的Unicode码点或字节值；
// 2. multibyte，一个布尔值，表示解码的字符是否需要多字节的UTF-8表示；
// 3. tail，字符之后的字符串剩余部分；以及
// 4. 一个错误，如果字符语法有效则为 nil。
//
// 第二个参数 quote 指定了正在解析的字面量类型，
// 因此指定了允许的转义引号字符。
// 如果设置为单引号，它允许序列 \' 并禁止未转义的 '。
// 如果设置为双引号，它允许 \" 并禁止未转义的 "。
// 如果设置为零，它不允许可转义的任何字符，并允许两个引号字符都未转义出现。
// md5:696a57be23951a5b

// ff:
// err:
// tail:
// multibyte:
// value:
// quote:
// s:
func UnquoteChar(s string, quote byte) (value rune, multibyte bool, tail string, err error) { //md5:47059d7ab99db197
	return strconv.UnquoteChar(s, quote)
}

// QuotedPrefix返回字符串`s`的前缀，该前缀是Unquote理解的引号字符串。
// 如果`s`不以一个有效的引号字符串开头，QuotedPrefix将返回一个错误。
// md5:4044f9096d29287c

// ff:
// s:
func QuotedPrefix(s string) (string, error) { //md5:ebea22512ef55876
	return strconv.QuotedPrefix(s)
}

// Unquote将`s`解释为单引号、双引号或反引号包裹的Go字符串字面量，返回`s`引用的字符串值。（如果`s`是单引号，它将是Go中的字符字面量；Unquote返回对应的单个字符字符串。）
// md5:4dc0a59aa91b7001

// ff:
// s:
func Unquote(s string) (string, error) { //md5:189a812004f3bee5
	return strconv.Unquote(s)
}

// IsPrint 报告一个字符是否被Go定义为可打印，其定义与 unicode.IsPrint 相同：包括字母、数字、标点符号、符号以及ASCII空格。
// md5:47d95c4564e1d7ad

// ff:
// r:
func IsPrint(r rune) bool { //md5:c1063b162fbbeab1
	return strconv.IsPrint(r)
}

// IsGraphic 报告该 rune 是否被 Unicode 定义为图形字符。这些字符包括字母、标记、数字、标点符号、符号和空格，它们来自类别 L, M, N, P, S 和 Zs。
// md5:a012ef5e5dfc3bfa

// ff:
// r:
func IsGraphic(r rune) bool { //md5:a5bcb73011ce25df
	return strconv.IsGraphic(r)
}
