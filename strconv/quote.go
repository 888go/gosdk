package strconv


const (
	lowerhex = "0123456789abcdef"
	upperhex = "0123456789ABCDEF"
)// md5:89575fd411bf282e



// Quote returns a double-quoted Go string literal representing s. The
// returned string uses Go escape sequences (\t, \n, \xFF, \u0100) for
// control characters and non-printable characters as defined by
// IsPrint.
func Quote(s string) string { //md5:720e1f755c129ec4

}

// AppendQuote appends a double-quoted Go string literal representing s,
// as generated by Quote, to dst and returns the extended buffer.
func AppendQuote(dst []byte, s string) []byte { //md5:cf37bac05af703ba

}

// QuoteToASCII returns a double-quoted Go string literal representing s.
// The returned string uses Go escape sequences (\t, \n, \xFF, \u0100) for
// non-ASCII characters and non-printable characters as defined by IsPrint.
func QuoteToASCII(s string) string { //md5:4b7b24b1728336e3

}

// AppendQuoteToASCII appends a double-quoted Go string literal representing s,
// as generated by QuoteToASCII, to dst and returns the extended buffer.
func AppendQuoteToASCII(dst []byte, s string) []byte { //md5:f8ada45efe15eb13

}

// QuoteToGraphic returns a double-quoted Go string literal representing s.
// The returned string leaves Unicode graphic characters, as defined by
// IsGraphic, unchanged and uses Go escape sequences (\t, \n, \xFF, \u0100)
// for non-graphic characters.
func QuoteToGraphic(s string) string { //md5:7f85e94d973176b3

}

// AppendQuoteToGraphic appends a double-quoted Go string literal representing s,
// as generated by QuoteToGraphic, to dst and returns the extended buffer.
func AppendQuoteToGraphic(dst []byte, s string) []byte { //md5:2deccb3bb6ccbfb3

}

// QuoteRune returns a single-quoted Go character literal representing the
// rune. The returned string uses Go escape sequences (\t, \n, \xFF, \u0100)
// for control characters and non-printable characters as defined by IsPrint.
// If r is not a valid Unicode code point, it is interpreted as the Unicode
// replacement character U+FFFD.
func QuoteRune(r rune) string { //md5:b226c44e497fd3e3

}

// AppendQuoteRune appends a single-quoted Go character literal representing the rune,
// as generated by QuoteRune, to dst and returns the extended buffer.
func AppendQuoteRune(dst []byte, r rune) []byte { //md5:602282ed6520be78

}

// QuoteRuneToASCII returns a single-quoted Go character literal representing
// the rune. The returned string uses Go escape sequences (\t, \n, \xFF,
// \u0100) for non-ASCII characters and non-printable characters as defined
// by IsPrint.
// If r is not a valid Unicode code point, it is interpreted as the Unicode
// replacement character U+FFFD.
func QuoteRuneToASCII(r rune) string { //md5:742e283e1b10790d

}

// AppendQuoteRuneToASCII appends a single-quoted Go character literal representing the rune,
// as generated by QuoteRuneToASCII, to dst and returns the extended buffer.
func AppendQuoteRuneToASCII(dst []byte, r rune) []byte { //md5:66508131668caa97

}

// QuoteRuneToGraphic returns a single-quoted Go character literal representing
// the rune. If the rune is not a Unicode graphic character,
// as defined by IsGraphic, the returned string will use a Go escape sequence
// (\t, \n, \xFF, \u0100).
// If r is not a valid Unicode code point, it is interpreted as the Unicode
// replacement character U+FFFD.
func QuoteRuneToGraphic(r rune) string { //md5:3e6e1e87514384b6

}

// AppendQuoteRuneToGraphic appends a single-quoted Go character literal representing the rune,
// as generated by QuoteRuneToGraphic, to dst and returns the extended buffer.
func AppendQuoteRuneToGraphic(dst []byte, r rune) []byte { //md5:9a05751c921c9257

}

// CanBackquote reports whether the string s can be represented
// unchanged as a single-line backquoted string without control
// characters other than tab.
func CanBackquote(s string) bool { //md5:67370d10584f6df1

}

// UnquoteChar decodes the first character or byte in the escaped string
// or character literal represented by the string s.
// It returns four values:
//
//  1. value, the decoded Unicode code point or byte value;
//  2. multibyte, a boolean indicating whether the decoded character requires a multibyte UTF-8 representation;
//  3. tail, the remainder of the string after the character; and
//  4. an error that will be nil if the character is syntactically valid.
//
// The second argument, quote, specifies the type of literal being parsed
// and therefore which escaped quote character is permitted.
// If set to a single quote, it permits the sequence \' and disallows unescaped '.
// If set to a double quote, it permits \" and disallows unescaped ".
// If set to zero, it does not permit either escape and allows both quote characters to appear unescaped.
func UnquoteChar(s string, quote byte) (value rune, multibyte bool, tail string, err error) { //md5:47059d7ab99db197

}

// QuotedPrefix returns the quoted string (as understood by Unquote) at the prefix of s.
// If s does not start with a valid quoted string, QuotedPrefix returns an error.
func QuotedPrefix(s string) (string, error) { //md5:ebea22512ef55876

}

// Unquote interprets s as a single-quoted, double-quoted,
// or backquoted Go string literal, returning the string value
// that s quotes.  (If s is single-quoted, it would be a Go
// character literal; Unquote returns the corresponding
// one-character string.)
func Unquote(s string) (string, error) { //md5:189a812004f3bee5

}

// IsPrint reports whether the rune is defined as printable by Go, with
// the same definition as unicode.IsPrint: letters, numbers, punctuation,
// symbols and ASCII space.
func IsPrint(r rune) bool { //md5:c1063b162fbbeab1

}

// IsGraphic reports whether the rune is defined as a Graphic by Unicode. Such
// characters include letters, marks, numbers, punctuation, symbols, and
// spaces, from categories L, M, N, P, S, and Zs.
func IsGraphic(r rune) bool { //md5:a5bcb73011ce25df

}