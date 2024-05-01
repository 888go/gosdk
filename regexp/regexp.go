package regexp

import (
	"io"
	"regexp"
)

// Regexp是编译后的正则表达式的表示。
// Regexp是安全的，可以由多个goroutine并发使用，除了配置方法，如[Regexp.Longest]。
// md5:395c4b654f50016e
type Regexp struct {
	F regexp.Regexp
} //md5:7eb1a097f4ab35df

//var (
//	matchSize = [...]int
//	matchPool [len(matchSize)]sync.Pool
//)// md5:77c917aa79b33031

// String 返回用于编译正则表达式的源文本。. md5:357ff5dae5501f98
func (re *Regexp) String() string { //md5:45608727399296bc
	return re.F.String()
}

// Copy 返回一个新的 [Regexp] 对象，它是 re 的副本。
// 对一个副本调用 [Regexp.Longest] 不会影响另一个副本。
//
// 注意：在早期版本中，当在多个 goroutine 中使用 [Regexp] 时，
// 为每个 goroutine 提供自己的副本有助于避免锁竞争。
// 自从 Go 1.12 开始，为了避免锁竞争，不再需要使用 Copy。
// 如果使用 Copy 的原因是创建具有不同 [Regexp.Longest] 设置的两个副本，那么它可能仍然适用。
// md5:cd49f88bd26658ca
func (re *Regexp) Copy() *Regexp { //md5:65840f82f6249997
	返回 := re.F.Copy()
	if 返回 == nil {
		return nil
	}
	return &Regexp{*返回}
}

// Compile解析一个正则表达式，如果成功，将返回一个[Regexp]对象，可以用来匹配文本。
// 
// 在匹配文本时，正则表达式会返回一个尽可能早开始（最左）的匹配，并在这些匹配中选择回溯搜索最早找到的那个。这种所谓的“最左优先”匹配与Perl、Python和其他实现使用的语义相同，尽管这个包在没有回溯开销的情况下实现了它。对于POSIX的“最左最长”匹配，请参阅[CompilePOSIX]。
// md5:56c7e946ce8c8782
func Compile(expr string) (*Regexp, error) { //md5:3c826d7b6981284e
	re, err := regexp.Compile(expr)
	if err != nil {
		return nil, err
	}
	return &Regexp{*re}, err
}

// CompilePOSIX 类似于 [Compile]，但它限制正则表达式的语法为 POSIX ERE（egrep）语法，并更改匹配语义为左最长模式。
// 
// 即，当在文本中进行匹配时，正则表达式会返回一个尽可能早开始于输入（最左边）的匹配，并在这些匹配中选择长度尽可能长的一个。这种所谓的左最长匹配是早期正则表达式实现和 POSIX 规范所指定的相同语义。
// 
// 然而，可能存在多个左最长匹配，有不同的子匹配选择。在这个包中，与 POSIX 规范有所不同。在可能的左最长匹配中，此包会选择回溯搜索首先找到的那个，而 POSIX 规范规定应选择第一个子表达式的最长匹配，然后是第二个，从左到右依次类推。POSIX 规则在计算上是禁止的，甚至没有明确定义。有关详细信息，请参阅 https://swtch.com/~rsc/regexp/regexp2.html#posix。
// md5:5229ea7fcaf45216
func CompilePOSIX(expr string) (*Regexp, error) { //md5:bdbccd70ec6febdf
	re, err := regexp.CompilePOSIX(expr)
	if err != nil {
		return nil, err
	}
	return &Regexp{*re}, err
}

// Longest 使得未来的搜索倾向于选择最长的左端匹配。
// 具体来说，当针对文本进行匹配时，正则表达式会返回一个在输入文本中尽可能早开始的匹配（最左端），并且在这些匹配中，
// 它会选择长度尽可能长的那个匹配项。
// 此方法会修改[Regexp]对象，并且不能与其他任何方法同时调用。
// md5:d03097bc3e482938
func (re *Regexp) Longest() { //md5:9ce0816bf705b7c4
	re.F.Longest()
}

// MustCompile 类似于 [Compile]，但如果表达式无法解析，则会引发 panic。
// 它简化了安全初始化包含编译后正则表达式的全局变量的过程。
// md5:c78888013a93d591
func MustCompile(str string) *Regexp { //md5:03f26daf41d7ea01
	re := regexp.MustCompile(str)
	if re == nil {
		return nil
	}
	return &Regexp{*re}
}

// MustCompilePOSIX 类似于 [CompilePOSIX]，但如果表达式无法解析，则会引发恐慌。它简化了持有编译正则表达式的全局变量的安全初始化。
// md5:c96838f4464d1b6b
func MustCompilePOSIX(str string) *Regexp { //md5:91f8f834a82fb8b8
	re := regexp.MustCompilePOSIX(str)
	if re == nil {
		return nil
	}
	return &Regexp{*re}
}

// NumSubexp 返回这个 [Regexp] 中的括号子表达式数量。. md5:442e4def06925f03
func (re *Regexp) NumSubexp() int { //md5:017364e4b734df3d
	return re.F.NumSubexp()
}

// SubexpNames 返回此[Regexp]中括号内子表达式的名称。
// 第一个子表达式的名称为names[1]，因此如果m是一个匹配切片，那么m[i]的名称就是SubexpNames()[i]。
// 由于正则表达式整体不能被命名，names[0]始终为空字符串。这个切片不应该被修改。
// md5:ddbb656dd72597e1
func (re *Regexp) SubexpNames() []string { //md5:7f377e5d157cd380
	return re.F.SubexpNames()
}

// SubexpIndex 返回给定名称的第一个子表达式的索引，如果找不到该名称的子表达式，则返回 -1。
//
// 请注意，可以使用相同的名称编写多个子表达式，如：(?P<bob>a+)(?P<bob>b+)，这声明了两个名为 "bob" 的子表达式。
// 在这种情况下，SubexpIndex 返回正则表达式中最左边的此类子表达式的索引。
// md5:4b4413493a8a87cc
func (re *Regexp) SubexpIndex(name string) int { //md5:474bfdf6795dcc5b
	return re.F.SubexpIndex(name)
}

// LiteralPrefix 返回一个必须以正则表达式 `re` 开头的字面字符串。如果这个字面字符串构成了整个正则表达式，它将返回布尔值 true。
// md5:b70b8246393ebc52
func (re *Regexp) LiteralPrefix() (prefix string, complete bool) { //md5:a605e6907f30fbb1
	return re.F.LiteralPrefix()
}

// MatchReader检查由[io.RuneReader]返回的文本中是否包含正则表达式re的任何匹配项。
// md5:7b14404a810353f5
func (re *Regexp) MatchReader(r io.RuneReader) bool { //md5:97957aaa7d7cd8c2
	return re.F.MatchReader(r)
}

// MatchString 报告字符串 s 是否包含与正则表达式 re 的任何匹配。
// md5:e184252b583086d1
func (re *Regexp) MatchString(s string) bool { //md5:95e39a4e814e8845
	return re.F.MatchString(s)
}

// Match reports whether the byte slice b
// 包含任何与正则表达式 re 匹配的字符串。
// md5:7675b942438cbaf7
func (re *Regexp) Match(b []byte) bool { //md5:85631a140de5124a
	return re.F.Match(b)
}

// MatchReader报告RuneReader返回的文本是否包含正则表达式模式的任何匹配项。更复杂的查询需要使用[Compile]和完整的[Regexp]接口。
// md5:8ce85e1bade17804
func MatchReader(pattern string, r io.RuneReader) (matched bool, err error) { //md5:244f6294a698c034
	return regexp.MatchReader(pattern, r)
}

// MatchString 检查字符串 s 是否包含正则表达式模式的任何匹配项。对于更复杂的查询，需要使用 [Compile] 和完整的 [Regexp] 接口。
// md5:a7bdb640a51211a6
func MatchString(pattern string, s string) (matched bool, err error) { //md5:bf08e1319dd4b3d4
	return regexp.MatchString(pattern, s)
}

// Match 报告字节切片 b 是否包含正则表达式模式的任何匹配。
// 更复杂的查询需要使用 [Compile] 并利用完整的 [Regexp] 接口。
// md5:0a22330a54acbcbf
func Match(pattern string, b []byte) (matched bool, err error) { //md5:a10eb5cfb2626d89
	return regexp.Match(pattern, b)
}

// ReplaceAllString 返回 src 的副本，将其中匹配 [Regexp] 的部分替换为替换字符串 repl。
// 在 repl 中，$ 符号的解释方式与 [Regexp.Expand] 中相同。
// md5:a74f9fac9979e0e8
func (re *Regexp) ReplaceAllString(src, repl string) string { //md5:f9e3b643868b80d1
	return re.F.ReplaceAllString(src, repl)
}

// ReplaceAllLiteralString 函数返回一个副本src，将正则表达式[Regexp]匹配的部分替换为替换字符串repl。替换字符串repl直接插入，不使用[Regexp.Expand]方法进行扩展。
// md5:b88f0714796f145b
func (re *Regexp) ReplaceAllLiteralString(src, repl string) string { //md5:a2819e606a7d5429
	return re.F.ReplaceAllLiteralString(src, repl)
}

// ReplaceAllStringFunc 函数返回一个副本，其中 src 中所有匹配正则表达式的部分都已被应用到匹配子串的 repl 函数返回值所替换。repl 函数的替换结果直接插入，不使用 [Regexp.Expand]。
// md5:e03dcdab858f9591
func (re *Regexp) ReplaceAllStringFunc(src string, repl func(string) string) string { //md5:18d8f8ee140dfdf6
	return re.F.ReplaceAllStringFunc(src, repl)
}

// ReplaceAll 返回src的副本，将[Regexp]匹配到的部分替换为替换文本repl。
// 在repl内部，$符号的解释方式与[Regexp.Expand]中的一致。
// md5:825b09dd0af8913d
func (re *Regexp) ReplaceAll(src, repl []byte) []byte { //md5:12b4e297a507e3c6
	return re.F.ReplaceAll(src, repl)
}

// ReplaceAllLiteral 返回 src 的副本，将其中的正则表达式匹配项替换为替换字节序列 repl。替换字符串 repl 直接插入，不使用 Regexp.Expand 进行扩展。
// md5:69edfa9b0a7f50e0
func (re *Regexp) ReplaceAllLiteral(src, repl []byte) []byte { //md5:4d3f9b55e70dcf01
	return re.F.ReplaceAllLiteral(src, repl)
}

// ReplaceAllFunc 返回一个副本，其中src中的所有 [正则表达式] 都已被应用到匹配的字节切片上的函数repl的返回值所替换。repl返回的替换直接插入，不使用 [Regexp.Expand]。
// md5:516c5932c8778484
func (re *Regexp) ReplaceAllFunc(src []byte, repl func([]byte) []byte) []byte { //md5:bb40f4c677ebc89c
	return re.F.ReplaceAllFunc(src, repl)
}

// QuoteMeta 返回一个字符串，该字符串转义了参数文本中的所有正则表达式元字符；返回的字符串是一个匹配实际文本的正则表达式。
// md5:26793bc2bd722708
func QuoteMeta(s string) string { //md5:c8123a1e462e206c
	return regexp.QuoteMeta(s)
}

// Find 返回一个切片，其中包含正则表达式在 b 中最左侧匹配项的文本。
// 如果返回值为 nil，则表示没有匹配项。
// md5:8611759571256d57
func (re *Regexp) Find(b []byte) []byte { //md5:7a2013ecf326ab53
	return re.F.Find(b)
}

// FindIndex 返回一个包含两个整数的切片，定义了在 b 中正则表达式的最左边匹配位置。
// 匹配的内容位于 b[loc[0]:loc[1]]。
// 如果没有找到匹配，则返回值为 nil。
// md5:5a412fd21a67619b
func (re *Regexp) FindIndex(b []byte) (loc []int) { //md5:64b07d0b7c7bf391
	return re.F.FindIndex(b)
}

// FindString 在字符串s中查找正则表达式的第一匹配文本。如果没有找到匹配项，返回值为空字符串。但如果正则表达式成功匹配空字符串，也会返回空字符串。如果需要区分这两种情况，请使用[Regexp.FindStringIndex]或[Regexp.FindStringSubmatch]方法。
// md5:72b4595ef9c7f90b
func (re *Regexp) FindString(s string) string { //md5:5b3ee1592bc499ba
	return re.F.FindString(s)
}

// FindStringIndex 函数返回一个包含两个整数的切片，定义 s 中正则表达式的第一匹配位置。匹配本身在 s[loc[0]:loc[1]]。如果未找到匹配，返回值为 nil。
// md5:52ee5c85c529f41f
func (re *Regexp) FindStringIndex(s string) (loc []int) { //md5:468730f50f8c7d7e
	return re.F.FindStringIndex(s)
}

// FindReaderIndex 返回一个包含两个整数的切片，定义了从 [io.RuneReader] 读取的文本中
// 最左边匹配正则表达式的起始和结束位置。匹配的文本在输入流中的字节偏移位置为
// loc[0] 到 loc[1]-1。
// 如果返回值为 nil，则表示没有找到匹配项。
// md5:7046ebc9094c6b29
func (re *Regexp) FindReaderIndex(r io.RuneReader) (loc []int) { //md5:94ed6f7d8e505d91
	return re.F.FindReaderIndex(r)
}

// FindSubmatch 返回一个切片，其中包含正则表达式在 b 中最左侧的匹配文本及其子表达式的匹配文本（如果有的话），如包注释中'Submatch'的描述所示。
// 如果没有匹配项，则返回值为 nil。
// md5:fae7fcc8c8b9ff44
func (re *Regexp) FindSubmatch(b []byte) [][]byte { //md5:a061db106859700f
	return re.F.FindSubmatch(b)
}

// Expand 将模板追加到 dst 中，并返回结果。在追加过程中，Expand 会将模板中的变量替换为从 src 中对应的匹配项。match 切片应该由 [Regexp.FindSubmatchIndex] 返回。
// 
// 在模板中，变量由形如 `$name` 或 `${name}` 的子字符串表示，其中 name 是由字母、数字和下划线组成的非空序列。纯数字的名称（如 `$1`）引用与相应索引相对应的子匹配；其他名称引用使用 `(?P<name>...)` 语法命名的捕获括号。引用超出范围或未匹配的索引，或者在正则表达式中不存在的名称将被替换为空切片。
// 
// 在 `$name` 形式中，name 被认为尽可能长：`$1x` 等同于 `${1x}`，而不是 `${1}x`；而 `$10` 等同于 `${10}`，而不是 `${1}0`。
// 
// 若要在输出中插入一个实际的 `$`，请在模板中使用 `$$`。
// md5:95b0e8b3e555f01d
func (re *Regexp) Expand(dst []byte, template []byte, src []byte, match []int) []byte { //md5:ec5be754a6f94e44
	return re.F.Expand(dst, template, src, match)
}

// ExpandString类似于[Regexp.Expand]，但是模板和源是字符串。它会追加到并返回一个字节切片，以便于调用代码控制内存分配。
// md5:8f82294f3864dc06
func (re *Regexp) ExpandString(dst []byte, template string, src string, match []int) []byte { //md5:70c650c8c278521a
	return re.F.ExpandString(dst, template, src, match)
}

// FindSubmatchIndex 返回一个切片，其中包含标识正则表达式在b中最左边匹配及其子表达式（如果有的话）的匹配索引对。
// 这些匹配索引对是根据包注释中“Submatch”和“Index”描述来定义的。
// 如果返回值为nil，则表示没有找到匹配项。
// md5:23a8d21b554faac0
func (re *Regexp) FindSubmatchIndex(b []byte) []int { //md5:e8817c1413c48d83
	return re.F.FindSubmatchIndex(b)
}

// FindStringSubmatch 返回一个字符串切片，其中包含正则表达式在s中的左most匹配文本以及（如果有的话）其子表达式的匹配项，如包注释中'Submatch'部分所定义。
// 如果没有匹配，则返回nil。
// md5:ef4a3f9dda34fb3c
func (re *Regexp) FindStringSubmatch(s string) []string { //md5:1d72c5d9710b0389
	return re.F.FindStringSubmatch(s)
}

// FindStringSubmatchIndex 函数返回一个切片，其中包含在字符串 s 中正则表达式的最左匹配的索引对，以及（如果有的话）其子表达式的匹配，这些定义了包注释中的 "Submatch" 和 "Index" 描述。
// 如果没有匹配，则返回值为 nil。
// md5:827dd5d33024bdd7
func (re *Regexp) FindStringSubmatchIndex(s string) []int { //md5:b64637981e1117e6
	return re.F.FindStringSubmatchIndex(s)
}

// FindReaderSubmatchIndex 函数返回一个切片，其中包含由 [io.RuneReader] 读取的文本与正则表达式的最左侧匹配的索引对。此外，它还包含了子表达式（如 'Submatch' 和 'Index' 注释中定义的）的任何匹配结果。如果没有匹配，函数将返回 nil。
// md5:2a96d023de151411
func (re *Regexp) FindReaderSubmatchIndex(r io.RuneReader) []int { //md5:5d06b081cbcb49ed
	return re.F.FindReaderSubmatchIndex(r)
}

// FindAll 是 Find 的“全量”版本；它返回表达式所有连续匹配项的切片，这些匹配项的定义参见包注释中的“全量”描述。
// 返回值为 nil 表示没有匹配项。
// md5:92ace0e31b493c01
func (re *Regexp) FindAll(b []byte, n int) [][]byte { //md5:9b102bc1c7db0f21
	return re.F.FindAll(b, n)
}

// FindAllIndex 是 Regexp.FindIndex 的“全量”版本，它返回一个切片，包含表达式的所有连续匹配结果，如包注释中
// 对“全量”描述的那样。如果未找到匹配，则返回值为nil。
// md5:85d767b529d7a932
func (re *Regexp) FindAllIndex(b []byte, n int) [][]int { //md5:8a23223fa7f7fde6
	return re.F.FindAllIndex(b, n)
}

// FindAllString是[Regexp.FindString]的'All'版本；它返回一个切片，包含由包注释中'description'部分定义的所有连续匹配项。
// 如果没有匹配，返回值为nil。
// md5:b772d35cf4130731
func (re *Regexp) FindAllString(s string, n int) []string { //md5:73c667fc476c6048
	return re.F.FindAllString(s, n)
}

// FindAllStringIndex 是 [Regexp.FindStringIndex] 的 "All" 版本；它返回一个切片，包含表达式定义的所有连续匹配项，如包注释中所述的 "All" 描述。
// 如果没有匹配，则返回值为 nil。
// md5:8a6c9067907c9fe3
func (re *Regexp) FindAllStringIndex(s string, n int) [][]int { //md5:d089e4011e9eb9d4
	return re.F.FindAllStringIndex(s, n)
}

// FindAllSubmatch 是 [Regexp.FindSubmatch] 的“全量”版本；它返回表达式所有连续匹配的切片，
// 匹配规则遵循包注释中'All'描述。
// 返回值为nil表示没有找到匹配项。
// md5:4b981f0bc28c5b31
func (re *Regexp) FindAllSubmatch(b []byte, n int) [][][]byte { //md5:f127bd4bd23f3b9d
	return re.F.FindAllSubmatch(b, n)
}

// FindAllSubmatchIndex 是 Regexp.FindSubmatchIndex 的“全量”版本；它返回表达式的所有连续匹配结果，如包注释中“全量”描述所示。
// 返回值为 nil 表示没有匹配到结果。
// md5:2c0f2bedad924158
func (re *Regexp) FindAllSubmatchIndex(b []byte, n int) [][]int { //md5:80d3879e67d43ea0
	return re.F.FindAllSubmatchIndex(b, n)
}

// FindAllStringSubmatch 是 [Regexp.FindStringSubmatch] 的 "All" 版本；它返回一个切片，其中包含表达式定义的所有连续匹配，如包注释中的 "All" 描述所示。
// 返回值为 nil 表示未找到匹配。
// md5:eea680ee72e47e40
func (re *Regexp) FindAllStringSubmatch(s string, n int) [][]string { //md5:157e59ed8a431048
	return re.F.FindAllStringSubmatch(s, n)
}

// FindAllStringSubmatchIndex是[Regexp.FindStringSubmatchIndex]的"所有"版本；它返回一个切片，其中包含表达式的所有连续匹配，如包注释中所述的"所有"描述。
// 返回值为nil表示未找到匹配。
// md5:c8a726eed55d8150
func (re *Regexp) FindAllStringSubmatchIndex(s string, n int) [][]int { //md5:906dc4d31409fb06
	return re.F.FindAllStringSubmatchIndex(s, n)
}

// 使用表达式将字符串s切分为子字符串，并返回这些表达式匹配之间的子字符串切片。
//
// 该方法返回的切片包含s中的所有子字符串，这些子字符串不由[Regexp.FindAllString]返回的切片包含。当在一个不包含元字符的表达式上调用时，它等同于[strings.SplitN]。
//
// 示例：
//
//	s := regexp.MustCompile("a*").Split("abaabaccadaaae", 5)
//	// s: ["", "b", "b", "c", "cadaaae"]
//
// count参数决定了要返回的子字符串数量：
//
//	n > 0: 最多n个子字符串；最后一个子字符串将是未分割的剩余部分。
//	n == 0: 结果为nil（零个子字符串）
//	n < 0: 所有子字符串
// md5:18189d3547a10a51
func (re *Regexp) Split(s string, n int) []string { //md5:638c9452844fd6a2
	return re.F.Split(s, n)
}

// MarshalText 实现了 encoding.TextMarshaler 接口。输出的结果与调用 [Regexp.String] 方法相同。
//
// 请注意，在某些情况下，输出可能会有信息丢失：该方法不会表示 POSIX 正则表达式（即，通过调用 [CompilePOSIX] 编译的那些），也不会表示已经调用了 [Regexp.Longest] 方法的正则表达式。
// md5:00df23a876c2729a
func (re *Regexp) MarshalText() ([]byte, error) { //md5:838b77790edfcb06
	return re.F.MarshalText()
}

// UnmarshalText实现了[encoding.TextUnmarshaler]接口，通过调用[Compile]方法处理编码后的值。
// md5:6e7b96fb96620957
func (re *Regexp) UnmarshalText(text []byte) error { //md5:5ba92bc5599b3a4c
	return re.F.UnmarshalText(text)
}
