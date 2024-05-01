
<原文开始>
// Regexp is the representation of a compiled regular expression.
// A Regexp is safe for concurrent use by multiple goroutines,
// except for configuration methods, such as [Regexp.Longest].
<原文结束>

# <翻译开始>
// Regexp是编译后的正则表达式的表示。
// Regexp是安全的，可以由多个goroutine并发使用，除了配置方法，如[Regexp.Longest]。
// md5:395c4b654f50016e
# <翻译结束>


<原文开始>
// String returns the source text used to compile the regular expression.
<原文结束>

# <翻译开始>
// String 返回用于编译正则表达式的源文本。. md5:357ff5dae5501f98
# <翻译结束>


<原文开始>
// Copy returns a new [Regexp] object copied from re.
// Calling [Regexp.Longest] on one copy does not affect another.
//
// Deprecated: In earlier releases, when using a [Regexp] in multiple goroutines,
// giving each goroutine its own copy helped to avoid lock contention.
// As of Go 1.12, using Copy is no longer necessary to avoid lock contention.
// Copy may still be appropriate if the reason for its use is to make
// two copies with different [Regexp.Longest] settings.
<原文结束>

# <翻译开始>
// Copy 返回一个新的 [Regexp] 对象，它是 re 的副本。
// 对一个副本调用 [Regexp.Longest] 不会影响另一个副本。
//
// 注意：在早期版本中，当在多个 goroutine 中使用 [Regexp] 时，
// 为每个 goroutine 提供自己的副本有助于避免锁竞争。
// 自从 Go 1.12 开始，为了避免锁竞争，不再需要使用 Copy。
// 如果使用 Copy 的原因是创建具有不同 [Regexp.Longest] 设置的两个副本，那么它可能仍然适用。
// md5:cd49f88bd26658ca
# <翻译结束>


<原文开始>
// Compile parses a regular expression and returns, if successful,
// a [Regexp] object that can be used to match against text.
//
// When matching against text, the regexp returns a match that
// begins as early as possible in the input (leftmost), and among those
// it chooses the one that a backtracking search would have found first.
// This so-called leftmost-first matching is the same semantics
// that Perl, Python, and other implementations use, although this
// package implements it without the expense of backtracking.
// For POSIX leftmost-longest matching, see [CompilePOSIX].
<原文结束>

# <翻译开始>
// Compile解析一个正则表达式，如果成功，将返回一个[Regexp]对象，可以用来匹配文本。
// 
// 在匹配文本时，正则表达式会返回一个尽可能早开始（最左）的匹配，并在这些匹配中选择回溯搜索最早找到的那个。这种所谓的“最左优先”匹配与Perl、Python和其他实现使用的语义相同，尽管这个包在没有回溯开销的情况下实现了它。对于POSIX的“最左最长”匹配，请参阅[CompilePOSIX]。
// md5:56c7e946ce8c8782
# <翻译结束>


<原文开始>
// CompilePOSIX is like [Compile] but restricts the regular expression
// to POSIX ERE (egrep) syntax and changes the match semantics to
// leftmost-longest.
//
// That is, when matching against text, the regexp returns a match that
// begins as early as possible in the input (leftmost), and among those
// it chooses a match that is as long as possible.
// This so-called leftmost-longest matching is the same semantics
// that early regular expression implementations used and that POSIX
// specifies.
//
// However, there can be multiple leftmost-longest matches, with different
// submatch choices, and here this package diverges from POSIX.
// Among the possible leftmost-longest matches, this package chooses
// the one that a backtracking search would have found first, while POSIX
// specifies that the match be chosen to maximize the length of the first
// subexpression, then the second, and so on from left to right.
// The POSIX rule is computationally prohibitive and not even well-defined.
// See https://swtch.com/~rsc/regexp/regexp2.html#posix for details.
<原文结束>

# <翻译开始>
// CompilePOSIX 类似于 [Compile]，但它限制正则表达式的语法为 POSIX ERE（egrep）语法，并更改匹配语义为左最长模式。
// 
// 即，当在文本中进行匹配时，正则表达式会返回一个尽可能早开始于输入（最左边）的匹配，并在这些匹配中选择长度尽可能长的一个。这种所谓的左最长匹配是早期正则表达式实现和 POSIX 规范所指定的相同语义。
// 
// 然而，可能存在多个左最长匹配，有不同的子匹配选择。在这个包中，与 POSIX 规范有所不同。在可能的左最长匹配中，此包会选择回溯搜索首先找到的那个，而 POSIX 规范规定应选择第一个子表达式的最长匹配，然后是第二个，从左到右依次类推。POSIX 规则在计算上是禁止的，甚至没有明确定义。有关详细信息，请参阅 https://swtch.com/~rsc/regexp/regexp2.html#posix。
// md5:5229ea7fcaf45216
# <翻译结束>


<原文开始>
// Longest makes future searches prefer the leftmost-longest match.
// That is, when matching against text, the regexp returns a match that
// begins as early as possible in the input (leftmost), and among those
// it chooses a match that is as long as possible.
// This method modifies the [Regexp] and may not be called concurrently
// with any other methods.
<原文结束>

# <翻译开始>
// Longest 使得未来的搜索倾向于选择最长的左端匹配。
// 具体来说，当针对文本进行匹配时，正则表达式会返回一个在输入文本中尽可能早开始的匹配（最左端），并且在这些匹配中，
// 它会选择长度尽可能长的那个匹配项。
// 此方法会修改[Regexp]对象，并且不能与其他任何方法同时调用。
// md5:d03097bc3e482938
# <翻译结束>


<原文开始>
// MustCompile is like [Compile] but panics if the expression cannot be parsed.
// It simplifies safe initialization of global variables holding compiled regular
// expressions.
<原文结束>

# <翻译开始>
// MustCompile 类似于 [Compile]，但如果表达式无法解析，则会引发 panic。
// 它简化了安全初始化包含编译后正则表达式的全局变量的过程。
// md5:c78888013a93d591
# <翻译结束>


<原文开始>
// MustCompilePOSIX is like [CompilePOSIX] but panics if the expression cannot be parsed.
// It simplifies safe initialization of global variables holding compiled regular
// expressions.
<原文结束>

# <翻译开始>
// MustCompilePOSIX 类似于 [CompilePOSIX]，但如果表达式无法解析，则会引发恐慌。它简化了持有编译正则表达式的全局变量的安全初始化。
// md5:c96838f4464d1b6b
# <翻译结束>


<原文开始>
// NumSubexp returns the number of parenthesized subexpressions in this [Regexp].
<原文结束>

# <翻译开始>
// NumSubexp 返回这个 [Regexp] 中的括号子表达式数量。. md5:442e4def06925f03
# <翻译结束>


<原文开始>
// SubexpNames returns the names of the parenthesized subexpressions
// in this [Regexp]. The name for the first sub-expression is names[1],
// so that if m is a match slice, the name for m[i] is SubexpNames()[i].
// Since the Regexp as a whole cannot be named, names[0] is always
// the empty string. The slice should not be modified.
<原文结束>

# <翻译开始>
// SubexpNames 返回此[Regexp]中括号内子表达式的名称。
// 第一个子表达式的名称为names[1]，因此如果m是一个匹配切片，那么m[i]的名称就是SubexpNames()[i]。
// 由于正则表达式整体不能被命名，names[0]始终为空字符串。这个切片不应该被修改。
// md5:ddbb656dd72597e1
# <翻译结束>


<原文开始>
// SubexpIndex returns the index of the first subexpression with the given name,
// or -1 if there is no subexpression with that name.
//
// Note that multiple subexpressions can be written using the same name, as in
// (?P<bob>a+)(?P<bob>b+), which declares two subexpressions named "bob".
// In this case, SubexpIndex returns the index of the leftmost such subexpression
// in the regular expression.
<原文结束>

# <翻译开始>
// SubexpIndex 返回给定名称的第一个子表达式的索引，如果找不到该名称的子表达式，则返回 -1。
//
// 请注意，可以使用相同的名称编写多个子表达式，如：(?P<bob>a+)(?P<bob>b+)，这声明了两个名为 "bob" 的子表达式。
// 在这种情况下，SubexpIndex 返回正则表达式中最左边的此类子表达式的索引。
// md5:4b4413493a8a87cc
# <翻译结束>


<原文开始>
// LiteralPrefix returns a literal string that must begin any match
// of the regular expression re. It returns the boolean true if the
// literal string comprises the entire regular expression.
<原文结束>

# <翻译开始>
// LiteralPrefix 返回一个必须以正则表达式 `re` 开头的字面字符串。如果这个字面字符串构成了整个正则表达式，它将返回布尔值 true。
// md5:b70b8246393ebc52
# <翻译结束>


<原文开始>
// MatchReader reports whether the text returned by the [io.RuneReader]
// contains any match of the regular expression re.
<原文结束>

# <翻译开始>
// MatchReader检查由[io.RuneReader]返回的文本中是否包含正则表达式re的任何匹配项。
// md5:7b14404a810353f5
# <翻译结束>


<原文开始>
// MatchString reports whether the string s
// contains any match of the regular expression re.
<原文结束>

# <翻译开始>
// MatchString 报告字符串 s 是否包含与正则表达式 re 的任何匹配。
// md5:e184252b583086d1
# <翻译结束>


<原文开始>
// Match reports whether the byte slice b
// contains any match of the regular expression re.
<原文结束>

# <翻译开始>
// Match reports whether the byte slice b
// 包含任何与正则表达式 re 匹配的字符串。
// md5:7675b942438cbaf7
# <翻译结束>


<原文开始>
// MatchReader reports whether the text returned by the RuneReader
// contains any match of the regular expression pattern.
// More complicated queries need to use [Compile] and the full [Regexp] interface.
<原文结束>

# <翻译开始>
// MatchReader报告RuneReader返回的文本是否包含正则表达式模式的任何匹配项。更复杂的查询需要使用[Compile]和完整的[Regexp]接口。
// md5:8ce85e1bade17804
# <翻译结束>


<原文开始>
// MatchString reports whether the string s
// contains any match of the regular expression pattern.
// More complicated queries need to use [Compile] and the full [Regexp] interface.
<原文结束>

# <翻译开始>
// MatchString 检查字符串 s 是否包含正则表达式模式的任何匹配项。对于更复杂的查询，需要使用 [Compile] 和完整的 [Regexp] 接口。
// md5:a7bdb640a51211a6
# <翻译结束>


<原文开始>
// Match reports whether the byte slice b
// contains any match of the regular expression pattern.
// More complicated queries need to use [Compile] and the full [Regexp] interface.
<原文结束>

# <翻译开始>
// Match 报告字节切片 b 是否包含正则表达式模式的任何匹配。
// 更复杂的查询需要使用 [Compile] 并利用完整的 [Regexp] 接口。
// md5:0a22330a54acbcbf
# <翻译结束>


<原文开始>
// ReplaceAllString returns a copy of src, replacing matches of the [Regexp]
// with the replacement string repl.
// Inside repl, $ signs are interpreted as in [Regexp.Expand].
<原文结束>

# <翻译开始>
// ReplaceAllString 返回 src 的副本，将其中匹配 [Regexp] 的部分替换为替换字符串 repl。
// 在 repl 中，$ 符号的解释方式与 [Regexp.Expand] 中相同。
// md5:a74f9fac9979e0e8
# <翻译结束>


<原文开始>
// ReplaceAllLiteralString returns a copy of src, replacing matches of the [Regexp]
// with the replacement string repl. The replacement repl is substituted directly,
// without using [Regexp.Expand].
<原文结束>

# <翻译开始>
// ReplaceAllLiteralString 函数返回一个副本src，将正则表达式[Regexp]匹配的部分替换为替换字符串repl。替换字符串repl直接插入，不使用[Regexp.Expand]方法进行扩展。
// md5:b88f0714796f145b
# <翻译结束>


<原文开始>
// ReplaceAllStringFunc returns a copy of src in which all matches of the
// [Regexp] have been replaced by the return value of function repl applied
// to the matched substring. The replacement returned by repl is substituted
// directly, without using [Regexp.Expand].
<原文结束>

# <翻译开始>
// ReplaceAllStringFunc 函数返回一个副本，其中 src 中所有匹配正则表达式的部分都已被应用到匹配子串的 repl 函数返回值所替换。repl 函数的替换结果直接插入，不使用 [Regexp.Expand]。
// md5:e03dcdab858f9591
# <翻译结束>


<原文开始>
// ReplaceAll returns a copy of src, replacing matches of the [Regexp]
// with the replacement text repl.
// Inside repl, $ signs are interpreted as in [Regexp.Expand].
<原文结束>

# <翻译开始>
// ReplaceAll 返回src的副本，将[Regexp]匹配到的部分替换为替换文本repl。
// 在repl内部，$符号的解释方式与[Regexp.Expand]中的一致。
// md5:825b09dd0af8913d
# <翻译结束>


<原文开始>
// ReplaceAllLiteral returns a copy of src, replacing matches of the [Regexp]
// with the replacement bytes repl. The replacement repl is substituted directly,
// without using [Regexp.Expand].
<原文结束>

# <翻译开始>
// ReplaceAllLiteral 返回 src 的副本，将其中的正则表达式匹配项替换为替换字节序列 repl。替换字符串 repl 直接插入，不使用 Regexp.Expand 进行扩展。
// md5:69edfa9b0a7f50e0
# <翻译结束>


<原文开始>
// ReplaceAllFunc returns a copy of src in which all matches of the
// [Regexp] have been replaced by the return value of function repl applied
// to the matched byte slice. The replacement returned by repl is substituted
// directly, without using [Regexp.Expand].
<原文结束>

# <翻译开始>
// ReplaceAllFunc 返回一个副本，其中src中的所有 [正则表达式] 都已被应用到匹配的字节切片上的函数repl的返回值所替换。repl返回的替换直接插入，不使用 [Regexp.Expand]。
// md5:516c5932c8778484
# <翻译结束>


<原文开始>
// QuoteMeta returns a string that escapes all regular expression metacharacters
// inside the argument text; the returned string is a regular expression matching
// the literal text.
<原文结束>

# <翻译开始>
// QuoteMeta 返回一个字符串，该字符串转义了参数文本中的所有正则表达式元字符；返回的字符串是一个匹配实际文本的正则表达式。
// md5:26793bc2bd722708
# <翻译结束>


<原文开始>
// Find returns a slice holding the text of the leftmost match in b of the regular expression.
// A return value of nil indicates no match.
<原文结束>

# <翻译开始>
// Find 返回一个切片，其中包含正则表达式在 b 中最左侧匹配项的文本。
// 如果返回值为 nil，则表示没有匹配项。
// md5:8611759571256d57
# <翻译结束>


<原文开始>
// FindIndex returns a two-element slice of integers defining the location of
// the leftmost match in b of the regular expression. The match itself is at
// b[loc[0]:loc[1]].
// A return value of nil indicates no match.
<原文结束>

# <翻译开始>
// FindIndex 返回一个包含两个整数的切片，定义了在 b 中正则表达式的最左边匹配位置。
// 匹配的内容位于 b[loc[0]:loc[1]]。
// 如果没有找到匹配，则返回值为 nil。
// md5:5a412fd21a67619b
# <翻译结束>


<原文开始>
// FindString returns a string holding the text of the leftmost match in s of the regular
// expression. If there is no match, the return value is an empty string,
// but it will also be empty if the regular expression successfully matches
// an empty string. Use [Regexp.FindStringIndex] or [Regexp.FindStringSubmatch] if it is
// necessary to distinguish these cases.
<原文结束>

# <翻译开始>
// FindString 在字符串s中查找正则表达式的第一匹配文本。如果没有找到匹配项，返回值为空字符串。但如果正则表达式成功匹配空字符串，也会返回空字符串。如果需要区分这两种情况，请使用[Regexp.FindStringIndex]或[Regexp.FindStringSubmatch]方法。
// md5:72b4595ef9c7f90b
# <翻译结束>


<原文开始>
// FindStringIndex returns a two-element slice of integers defining the
// location of the leftmost match in s of the regular expression. The match
// itself is at s[loc[0]:loc[1]].
// A return value of nil indicates no match.
<原文结束>

# <翻译开始>
// FindStringIndex 函数返回一个包含两个整数的切片，定义 s 中正则表达式的第一匹配位置。匹配本身在 s[loc[0]:loc[1]]。如果未找到匹配，返回值为 nil。
// md5:52ee5c85c529f41f
# <翻译结束>


<原文开始>
// FindReaderIndex returns a two-element slice of integers defining the
// location of the leftmost match of the regular expression in text read from
// the [io.RuneReader]. The match text was found in the input stream at
// byte offset loc[0] through loc[1]-1.
// A return value of nil indicates no match.
<原文结束>

# <翻译开始>
// FindReaderIndex 返回一个包含两个整数的切片，定义了从 [io.RuneReader] 读取的文本中
// 最左边匹配正则表达式的起始和结束位置。匹配的文本在输入流中的字节偏移位置为
// loc[0] 到 loc[1]-1。
// 如果返回值为 nil，则表示没有找到匹配项。
// md5:7046ebc9094c6b29
# <翻译结束>


<原文开始>
// FindSubmatch returns a slice of slices holding the text of the leftmost
// match of the regular expression in b and the matches, if any, of its
// subexpressions, as defined by the 'Submatch' descriptions in the package
// comment.
// A return value of nil indicates no match.
<原文结束>

# <翻译开始>
// FindSubmatch 返回一个切片，其中包含正则表达式在 b 中最左侧的匹配文本及其子表达式的匹配文本（如果有的话），如包注释中'Submatch'的描述所示。
// 如果没有匹配项，则返回值为 nil。
// md5:fae7fcc8c8b9ff44
# <翻译结束>


<原文开始>
// Expand appends template to dst and returns the result; during the
// append, Expand replaces variables in the template with corresponding
// matches drawn from src. The match slice should have been returned by
// [Regexp.FindSubmatchIndex].
//
// In the template, a variable is denoted by a substring of the form
// $name or ${name}, where name is a non-empty sequence of letters,
// digits, and underscores. A purely numeric name like $1 refers to
// the submatch with the corresponding index; other names refer to
// capturing parentheses named with the (?P<name>...) syntax. A
// reference to an out of range or unmatched index or a name that is not
// present in the regular expression is replaced with an empty slice.
//
// In the $name form, name is taken to be as long as possible: $1x is
// equivalent to ${1x}, not ${1}x, and, $10 is equivalent to ${10}, not ${1}0.
//
// To insert a literal $ in the output, use $$ in the template.
<原文结束>

# <翻译开始>
// Expand 将模板追加到 dst 中，并返回结果。在追加过程中，Expand 会将模板中的变量替换为从 src 中对应的匹配项。match 切片应该由 [Regexp.FindSubmatchIndex] 返回。
// 
// 在模板中，变量由形如 `$name` 或 `${name}` 的子字符串表示，其中 name 是由字母、数字和下划线组成的非空序列。纯数字的名称（如 `$1`）引用与相应索引相对应的子匹配；其他名称引用使用 `(?P<name>...)` 语法命名的捕获括号。引用超出范围或未匹配的索引，或者在正则表达式中不存在的名称将被替换为空切片。
// 
// 在 `$name` 形式中，name 被认为尽可能长：`$1x` 等同于 `${1x}`，而不是 `${1}x`；而 `$10` 等同于 `${10}`，而不是 `${1}0`。
// 
// 若要在输出中插入一个实际的 `$`，请在模板中使用 `$$`。
// md5:95b0e8b3e555f01d
# <翻译结束>


<原文开始>
// ExpandString is like [Regexp.Expand] but the template and source are strings.
// It appends to and returns a byte slice in order to give the calling
// code control over allocation.
<原文结束>

# <翻译开始>
// ExpandString类似于[Regexp.Expand]，但是模板和源是字符串。它会追加到并返回一个字节切片，以便于调用代码控制内存分配。
// md5:8f82294f3864dc06
# <翻译结束>


<原文开始>
// FindSubmatchIndex returns a slice holding the index pairs identifying the
// leftmost match of the regular expression in b and the matches, if any, of
// its subexpressions, as defined by the 'Submatch' and 'Index' descriptions
// in the package comment.
// A return value of nil indicates no match.
<原文结束>

# <翻译开始>
// FindSubmatchIndex 返回一个切片，其中包含标识正则表达式在b中最左边匹配及其子表达式（如果有的话）的匹配索引对。
// 这些匹配索引对是根据包注释中“Submatch”和“Index”描述来定义的。
// 如果返回值为nil，则表示没有找到匹配项。
// md5:23a8d21b554faac0
# <翻译结束>


<原文开始>
// FindStringSubmatch returns a slice of strings holding the text of the
// leftmost match of the regular expression in s and the matches, if any, of
// its subexpressions, as defined by the 'Submatch' description in the
// package comment.
// A return value of nil indicates no match.
<原文结束>

# <翻译开始>
// FindStringSubmatch 返回一个字符串切片，其中包含正则表达式在s中的左most匹配文本以及（如果有的话）其子表达式的匹配项，如包注释中'Submatch'部分所定义。
// 如果没有匹配，则返回nil。
// md5:ef4a3f9dda34fb3c
# <翻译结束>


<原文开始>
// FindStringSubmatchIndex returns a slice holding the index pairs
// identifying the leftmost match of the regular expression in s and the
// matches, if any, of its subexpressions, as defined by the 'Submatch' and
// 'Index' descriptions in the package comment.
// A return value of nil indicates no match.
<原文结束>

# <翻译开始>
// FindStringSubmatchIndex 函数返回一个切片，其中包含在字符串 s 中正则表达式的最左匹配的索引对，以及（如果有的话）其子表达式的匹配，这些定义了包注释中的 "Submatch" 和 "Index" 描述。
// 如果没有匹配，则返回值为 nil。
// md5:827dd5d33024bdd7
# <翻译结束>


<原文开始>
// FindReaderSubmatchIndex returns a slice holding the index pairs
// identifying the leftmost match of the regular expression of text read by
// the [io.RuneReader], and the matches, if any, of its subexpressions, as defined
// by the 'Submatch' and 'Index' descriptions in the package comment. A
// return value of nil indicates no match.
<原文结束>

# <翻译开始>
// FindReaderSubmatchIndex 函数返回一个切片，其中包含由 [io.RuneReader] 读取的文本与正则表达式的最左侧匹配的索引对。此外，它还包含了子表达式（如 'Submatch' 和 'Index' 注释中定义的）的任何匹配结果。如果没有匹配，函数将返回 nil。
// md5:2a96d023de151411
# <翻译结束>


<原文开始>
// FindAll is the 'All' version of Find; it returns a slice of all successive
// matches of the expression, as defined by the 'All' description in the
// package comment.
// A return value of nil indicates no match.
<原文结束>

# <翻译开始>
// FindAll 是 Find 的“全量”版本；它返回表达式所有连续匹配项的切片，这些匹配项的定义参见包注释中的“全量”描述。
// 返回值为 nil 表示没有匹配项。
// md5:92ace0e31b493c01
# <翻译结束>


<原文开始>
// FindAllIndex is the 'All' version of [Regexp.FindIndex]; it returns a slice of all
// successive matches of the expression, as defined by the 'All' description
// in the package comment.
// A return value of nil indicates no match.
<原文结束>

# <翻译开始>
// FindAllIndex 是 Regexp.FindIndex 的“全量”版本，它返回一个切片，包含表达式的所有连续匹配结果，如包注释中
// 对“全量”描述的那样。如果未找到匹配，则返回值为nil。
// md5:85d767b529d7a932
# <翻译结束>


<原文开始>
// FindAllString is the 'All' version of [Regexp.FindString]; it returns a slice of all
// successive matches of the expression, as defined by the 'All' description
// in the package comment.
// A return value of nil indicates no match.
<原文结束>

# <翻译开始>
// FindAllString是[Regexp.FindString]的'All'版本；它返回一个切片，包含由包注释中'description'部分定义的所有连续匹配项。
// 如果没有匹配，返回值为nil。
// md5:b772d35cf4130731
# <翻译结束>


<原文开始>
// FindAllStringIndex is the 'All' version of [Regexp.FindStringIndex]; it returns a
// slice of all successive matches of the expression, as defined by the 'All'
// description in the package comment.
// A return value of nil indicates no match.
<原文结束>

# <翻译开始>
// FindAllStringIndex 是 [Regexp.FindStringIndex] 的 "All" 版本；它返回一个切片，包含表达式定义的所有连续匹配项，如包注释中所述的 "All" 描述。
// 如果没有匹配，则返回值为 nil。
// md5:8a6c9067907c9fe3
# <翻译结束>


<原文开始>
// FindAllSubmatch is the 'All' version of [Regexp.FindSubmatch]; it returns a slice
// of all successive matches of the expression, as defined by the 'All'
// description in the package comment.
// A return value of nil indicates no match.
<原文结束>

# <翻译开始>
// FindAllSubmatch 是 [Regexp.FindSubmatch] 的“全量”版本；它返回表达式所有连续匹配的切片，
// 匹配规则遵循包注释中'All'描述。
// 返回值为nil表示没有找到匹配项。
// md5:4b981f0bc28c5b31
# <翻译结束>


<原文开始>
// FindAllSubmatchIndex is the 'All' version of [Regexp.FindSubmatchIndex]; it returns
// a slice of all successive matches of the expression, as defined by the
// 'All' description in the package comment.
// A return value of nil indicates no match.
<原文结束>

# <翻译开始>
// FindAllSubmatchIndex 是 Regexp.FindSubmatchIndex 的“全量”版本；它返回表达式的所有连续匹配结果，如包注释中“全量”描述所示。
// 返回值为 nil 表示没有匹配到结果。
// md5:2c0f2bedad924158
# <翻译结束>


<原文开始>
// FindAllStringSubmatch is the 'All' version of [Regexp.FindStringSubmatch]; it
// returns a slice of all successive matches of the expression, as defined by
// the 'All' description in the package comment.
// A return value of nil indicates no match.
<原文结束>

# <翻译开始>
// FindAllStringSubmatch 是 [Regexp.FindStringSubmatch] 的 "All" 版本；它返回一个切片，其中包含表达式定义的所有连续匹配，如包注释中的 "All" 描述所示。
// 返回值为 nil 表示未找到匹配。
// md5:eea680ee72e47e40
# <翻译结束>


<原文开始>
// FindAllStringSubmatchIndex is the 'All' version of
// [Regexp.FindStringSubmatchIndex]; it returns a slice of all successive matches of
// the expression, as defined by the 'All' description in the package
// comment.
// A return value of nil indicates no match.
<原文结束>

# <翻译开始>
// FindAllStringSubmatchIndex是[Regexp.FindStringSubmatchIndex]的"所有"版本；它返回一个切片，其中包含表达式的所有连续匹配，如包注释中所述的"所有"描述。
// 返回值为nil表示未找到匹配。
// md5:c8a726eed55d8150
# <翻译结束>


<原文开始>
// Split slices s into substrings separated by the expression and returns a slice of
// the substrings between those expression matches.
//
// The slice returned by this method consists of all the substrings of s
// not contained in the slice returned by [Regexp.FindAllString]. When called on an expression
// that contains no metacharacters, it is equivalent to [strings.SplitN].
//
// Example:
//
//	s := regexp.MustCompile("a*").Split("abaabaccadaaae", 5)
//	// s: ["", "b", "b", "c", "cadaaae"]
//
// The count determines the number of substrings to return:
//
//	n > 0: at most n substrings; the last substring will be the unsplit remainder.
//	n == 0: the result is nil (zero substrings)
//	n < 0: all substrings
<原文结束>

# <翻译开始>
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
# <翻译结束>


<原文开始>
// MarshalText implements [encoding.TextMarshaler]. The output
// matches that of calling the [Regexp.String] method.
//
// Note that the output is lossy in some cases: This method does not indicate
// POSIX regular expressions (i.e. those compiled by calling [CompilePOSIX]), or
// those for which the [Regexp.Longest] method has been called.
<原文结束>

# <翻译开始>
// MarshalText 实现了 encoding.TextMarshaler 接口。输出的结果与调用 [Regexp.String] 方法相同。
//
// 请注意，在某些情况下，输出可能会有信息丢失：该方法不会表示 POSIX 正则表达式（即，通过调用 [CompilePOSIX] 编译的那些），也不会表示已经调用了 [Regexp.Longest] 方法的正则表达式。
// md5:00df23a876c2729a
# <翻译结束>


<原文开始>
// UnmarshalText implements [encoding.TextUnmarshaler] by calling
// [Compile] on the encoded value.
<原文结束>

# <翻译开始>
// UnmarshalText实现了[encoding.TextUnmarshaler]接口，通过调用[Compile]方法处理编码后的值。
// md5:6e7b96fb96620957
# <翻译结束>

