// 版权所有 2013 Go 作者。保留所有权利。
// 使用此源代码受BSD风格许可证约束，该许可证可从LICENSE文件中找到。
// md5:19d1a3ed91182ee4

package regexp_test

import (
	"fmt"
	"github.com/888go/gosdk/regexp"
	"strings"
)

func Example() {
// 通常在初始化时编译表达式。使用原始字符串以避免需要转义反斜杠。
// md5:271c082b9d8a6c71
	var validID = regexp.MustCompile(`^[a-z]+\[[0-9]+\]$`)

	fmt.Println(validID.MatchString("adam[23]"))
	fmt.Println(validID.MatchString("eve[7]"))
	fmt.Println(validID.MatchString("Job[48]"))
	fmt.Println(validID.MatchString("snakey"))
	// Output:
	// true
	// true
	// false
	// false
}

func ExampleMatch() {
	matched, err := regexp.Match(`foo.*`, []byte(`seafood`))
	fmt.Println(matched, err)
	matched, err = regexp.Match(`bar.*`, []byte(`seafood`))
	fmt.Println(matched, err)
	matched, err = regexp.Match(`a(b`, []byte(`seafood`))
	fmt.Println(matched, err)

	// Output:
	// true <nil>
	// false <nil>
	// false error parsing regexp: missing closing ): `a(b`
}

func ExampleMatchString() {
	matched, err := regexp.MatchString(`foo.*`, "seafood")
	fmt.Println(matched, err)
	matched, err = regexp.MatchString(`bar.*`, "seafood")
	fmt.Println(matched, err)
	matched, err = regexp.MatchString(`a(b`, "seafood")
	fmt.Println(matched, err)
	// Output:
	// true <nil>
	// false <nil>
	// false error parsing regexp: missing closing ): `a(b`
}

func ExampleQuoteMeta() {
	fmt.Println(regexp.QuoteMeta(`Escaping symbols like: .+*?()|[]{}^$`))
	// Output:
	// Escaping symbols like: \.\+\*\?\(\)\|\[\]\{\}\^\$
}

func ExampleRegexp_Find() {
	re := regexp.MustCompile(`foo.?`)
	fmt.Printf("%q\n", re.Find([]byte(`seafood fool`)))

	// Output:
	// "food"
}

func ExampleRegexp_FindAll() {
	re := regexp.MustCompile(`foo.?`)
	fmt.Printf("%q\n", re.FindAll([]byte(`seafood fool`), -1))

	// Output:
	// ["food" "fool"]
}

func ExampleRegexp_FindAllSubmatch() {
	re := regexp.MustCompile(`foo(.?)`)
	fmt.Printf("%q\n", re.FindAllSubmatch([]byte(`seafood fool`), -1))

	// Output:
	// [["food" "d"] ["fool" "l"]]
}

func ExampleRegexp_FindSubmatch() {
	re := regexp.MustCompile(`foo(.?)`)
	fmt.Printf("%q\n", re.FindSubmatch([]byte(`seafood fool`)))

	// Output:
	// ["food" "d"]
}

func ExampleRegexp_Match() {
	re := regexp.MustCompile(`foo.?`)
	fmt.Println(re.Match([]byte(`seafood fool`)))
	fmt.Println(re.Match([]byte(`something else`)))

	// Output:
	// true
	// false
}

func ExampleRegexp_FindString() {
	re := regexp.MustCompile(`foo.?`)
	fmt.Printf("%q\n", re.FindString("seafood fool"))
	fmt.Printf("%q\n", re.FindString("meat"))
	// Output:
	// "food"
	// ""
}

func ExampleRegexp_FindStringIndex() {
	re := regexp.MustCompile(`ab?`)
	fmt.Println(re.FindStringIndex("tablett"))
	fmt.Println(re.FindStringIndex("foo") == nil)
	// Output:
	// [1 3]
	// true
}

func ExampleRegexp_FindStringSubmatch() {
	re := regexp.MustCompile(`a(x*)b(y|z)c`)
	fmt.Printf("%q\n", re.FindStringSubmatch("-axxxbyc-"))
	fmt.Printf("%q\n", re.FindStringSubmatch("-abzc-"))
	// Output:
	// ["axxxbyc" "xxx" "y"]
	// ["abzc" "" "z"]
}

func ExampleRegexp_FindAllString() {
	re := regexp.MustCompile(`a.`)
	fmt.Println(re.FindAllString("paranormal", -1))
	fmt.Println(re.FindAllString("paranormal", 2))
	fmt.Println(re.FindAllString("graal", -1))
	fmt.Println(re.FindAllString("none", -1))
	// Output:
	// [ar an al]
	// [ar an]
	// [aa]
	// []
}

func ExampleRegexp_FindAllStringSubmatch() {
	re := regexp.MustCompile(`a(x*)b`)
	fmt.Printf("%q\n", re.FindAllStringSubmatch("-ab-", -1))
	fmt.Printf("%q\n", re.FindAllStringSubmatch("-axxb-", -1))
	fmt.Printf("%q\n", re.FindAllStringSubmatch("-ab-axb-", -1))
	fmt.Printf("%q\n", re.FindAllStringSubmatch("-axxb-ab-", -1))
	// Output:
	// [["ab" ""]]
	// [["axxb" "xx"]]
	// [["ab" ""] ["axb" "x"]]
	// [["axxb" "xx"] ["ab" ""]]
}

func ExampleRegexp_FindAllStringSubmatchIndex() {
	re := regexp.MustCompile(`a(x*)b`)
// 索引：
//    01234567   012345678
//    ---a-xb---  ---x-xb-a---
// 
// 这段注释说明了索引的分布情况。在第一行中，从第3个位置（索引为2）开始到第6个位置（索引为5）是空的（标记为"-"），然后是字母"a"，接着是"x"和"b"，再接着又是空的。在第二行中，同样的模式，但这次是从第7个位置开始（索引为6）有"x"，然后是两个"x"，接着是字母"b"，最后再次是空的。
// md5:283dca55f0118208
	fmt.Println(re.FindAllStringSubmatchIndex("-ab-", -1))
	fmt.Println(re.FindAllStringSubmatchIndex("-axxb-", -1))
	fmt.Println(re.FindAllStringSubmatchIndex("-ab-axb-", -1))
	fmt.Println(re.FindAllStringSubmatchIndex("-axxb-ab-", -1))
	fmt.Println(re.FindAllStringSubmatchIndex("-foo-", -1))
	// Output:
	// [[1 3 2 2]]
	// [[1 5 2 4]]
	// [[1 3 2 2] [4 7 5 6]]
	// [[1 5 2 4] [6 8 7 7]]
	// []
}

func ExampleRegexp_FindSubmatchIndex() {
	re := regexp.MustCompile(`a(x*)b`)
// 索引：
//    01234567   012345678
//    ---a-xb---  ---x-xb-a---
// 
// 这段注释说明了索引的分布情况。在第一行中，从第3个位置（索引为2）开始到第6个位置（索引为5）是空的（标记为"-"），然后是字母"a"，接着是"x"和"b"，再接着又是空的。在第二行中，同样的模式，但这次是从第7个位置开始（索引为6）有"x"，然后是两个"x"，接着是字母"b"，最后再次是空的。
// md5:283dca55f0118208
	fmt.Println(re.FindSubmatchIndex([]byte("-ab-")))
	fmt.Println(re.FindSubmatchIndex([]byte("-axxb-")))
	fmt.Println(re.FindSubmatchIndex([]byte("-ab-axb-")))
	fmt.Println(re.FindSubmatchIndex([]byte("-axxb-ab-")))
	fmt.Println(re.FindSubmatchIndex([]byte("-foo-")))
	// Output:
	// [1 3 2 2]
	// [1 5 2 4]
	// [1 3 2 2]
	// [1 5 2 4]
	// []
}

func ExampleRegexp_Longest() {
	re := regexp.MustCompile(`a(|b)`)
	fmt.Println(re.FindString("ab"))
	re.Longest()
	fmt.Println(re.FindString("ab"))
	// Output:
	// a
	// ab
}

func ExampleRegexp_MatchString() {
	re := regexp.MustCompile(`(gopher){2}`)
	fmt.Println(re.MatchString("gopher"))
	fmt.Println(re.MatchString("gophergopher"))
	fmt.Println(re.MatchString("gophergophergopher"))
	// Output:
	// false
	// true
	// true
}

func ExampleRegexp_NumSubexp() {
	re0 := regexp.MustCompile(`a.`)
	fmt.Printf("%d\n", re0.NumSubexp())

	re := regexp.MustCompile(`(.*)((a)b)(.*)a`)
	fmt.Println(re.NumSubexp())
	// Output:
	// 0
	// 4
}

func ExampleRegexp_ReplaceAll() {
	re := regexp.MustCompile(`a(x*)b`)
	fmt.Printf("%s\n", re.ReplaceAll([]byte("-ab-axxb-"), []byte("T")))
	fmt.Printf("%s\n", re.ReplaceAll([]byte("-ab-axxb-"), []byte("$1")))
	fmt.Printf("%s\n", re.ReplaceAll([]byte("-ab-axxb-"), []byte("$1W")))
	fmt.Printf("%s\n", re.ReplaceAll([]byte("-ab-axxb-"), []byte("${1}W")))

	re2 := regexp.MustCompile(`a(?P<1W>x*)b`)
	fmt.Printf("%s\n", re2.ReplaceAll([]byte("-ab-axxb-"), []byte("$1W")))
	fmt.Printf("%s\n", re2.ReplaceAll([]byte("-ab-axxb-"), []byte("${1}W")))

	// Output:
	// -T-T-
	// --xx-
	// ---
	// -W-xxW-
	// --xx-
	// -W-xxW-
}

func ExampleRegexp_ReplaceAllLiteralString() {
	re := regexp.MustCompile(`a(x*)b`)
	fmt.Println(re.ReplaceAllLiteralString("-ab-axxb-", "T"))
	fmt.Println(re.ReplaceAllLiteralString("-ab-axxb-", "$1"))
	fmt.Println(re.ReplaceAllLiteralString("-ab-axxb-", "${1}"))
	// Output:
	// -T-T-
	// -$1-$1-
	// -${1}-${1}-
}

func ExampleRegexp_ReplaceAllString() {
	re := regexp.MustCompile(`a(x*)b`)
	fmt.Println(re.ReplaceAllString("-ab-axxb-", "T"))
	fmt.Println(re.ReplaceAllString("-ab-axxb-", "$1"))
	fmt.Println(re.ReplaceAllString("-ab-axxb-", "$1W"))
	fmt.Println(re.ReplaceAllString("-ab-axxb-", "${1}W"))

	re2 := regexp.MustCompile(`a(?P<1W>x*)b`)
	fmt.Printf("%s\n", re2.ReplaceAllString("-ab-axxb-", "$1W"))
	fmt.Println(re.ReplaceAllString("-ab-axxb-", "${1}W"))

	// Output:
	// -T-T-
	// --xx-
	// ---
	// -W-xxW-
	// --xx-
	// -W-xxW-
}

func ExampleRegexp_ReplaceAllStringFunc() {
	re := regexp.MustCompile(`[^aeiou]`)
	fmt.Println(re.ReplaceAllStringFunc("seafood fool", strings.ToUpper))
	// Output:
	// SeaFooD FooL
}

func ExampleRegexp_SubexpNames() {
	re := regexp.MustCompile(`(?P<first>[a-zA-Z]+) (?P<last>[a-zA-Z]+)`)
	fmt.Println(re.MatchString("Alan Turing"))
	fmt.Printf("%q\n", re.SubexpNames())
	reversed := fmt.Sprintf("${%s} ${%s}", re.SubexpNames()[2], re.SubexpNames()[1])
	fmt.Println(reversed)
	fmt.Println(re.ReplaceAllString("Alan Turing", reversed))
	// Output:
	// true
	// ["" "first" "last"]
	// ${last} ${first}
	// Turing Alan
}

func ExampleRegexp_SubexpIndex() {
	re := regexp.MustCompile(`(?P<first>[a-zA-Z]+) (?P<last>[a-zA-Z]+)`)
	fmt.Println(re.MatchString("Alan Turing"))
	matches := re.FindStringSubmatch("Alan Turing")
	lastIndex := re.SubexpIndex("last")
	fmt.Printf("last => %d\n", lastIndex)
	fmt.Println(matches[lastIndex])
	// Output:
	// true
	// last => 2
	// Turing
}

func ExampleRegexp_Split() {
	a := regexp.MustCompile(`a`)
	fmt.Println(a.Split("banana", -1))
	fmt.Println(a.Split("banana", 0))
	fmt.Println(a.Split("banana", 1))
	fmt.Println(a.Split("banana", 2))
	zp := regexp.MustCompile(`z+`)
	fmt.Println(zp.Split("pizza", -1))
	fmt.Println(zp.Split("pizza", 0))
	fmt.Println(zp.Split("pizza", 1))
	fmt.Println(zp.Split("pizza", 2))
	// Output:
	// [b n n ]
	// []
	// [banana]
	// [b nana]
	// [pi a]
	// []
	// [pizza]
	// [pi a]
}

func ExampleRegexp_Expand() {
	content := []byte(`
	# comment line
	option1: value1
	option2: value2

	# another comment line
	option3: value3
`)

	// 正则表达式模式用于从内容中捕获 "键: 值" 对。. md5:7e673e91ccc76b97
	pattern := regexp.MustCompile(`(?m)(?P<key>\w+):\s+(?P<value>\w+)$`)

// 模板用于通过引用正则表达式模式捕获的值将 "key: value" 转换为 "key=value"。
// md5:5bc55b3b185255ad
	template := []byte("$key=$value\n")

	result := []byte{}

	// 对内容中的正则表达式匹配项进行遍历。. md5:51840a927103ba23
	for _, submatches := range pattern.FindAllSubmatchIndex(content, -1) {
// 将捕获的子模式应用到模板中，并将输出追加到结果中。
// md5:0c0b9de7b94ce6a9
		result = pattern.Expand(result, template, content, submatches)
	}
	fmt.Println(string(result))
	// Output:
	// option1=value1
	// option2=value2
	// option3=value3
}

func ExampleRegexp_ExpandString() {
	content := `
	# comment line
	option1: value1
	option2: value2

	# another comment line
	option3: value3
`

	// 正则表达式模式用于从内容中捕获 "键: 值" 对。. md5:7e673e91ccc76b97
	pattern := regexp.MustCompile(`(?m)(?P<key>\w+):\s+(?P<value>\w+)$`)

// 模板用于通过引用正则表达式模式捕获的值将 "key: value" 转换为 "key=value"。
// md5:5bc55b3b185255ad
	template := "$key=$value\n"

	result := []byte{}

	// 对内容中的正则表达式匹配项进行遍历。. md5:51840a927103ba23
	for _, submatches := range pattern.FindAllStringSubmatchIndex(content, -1) {
// 将捕获的子模式应用到模板中，并将输出追加到结果中。
// md5:0c0b9de7b94ce6a9
		result = pattern.ExpandString(result, template, content, submatches)
	}
	fmt.Println(string(result))
	// Output:
	// option1=value1
	// option2=value2
	// option3=value3
}

func ExampleRegexp_FindIndex() {
	content := []byte(`
	# comment line
	option1: value1
	option2: value2
`)
	// 正则表达式模式用于从内容中捕获 "键: 值" 对。. md5:7e673e91ccc76b97
	pattern := regexp.MustCompile(`(?m)(?P<key>\w+):\s+(?P<value>\w+)$`)

	loc := pattern.FindIndex(content)
	fmt.Println(loc)
	fmt.Println(string(content[loc[0]:loc[1]]))
	// Output:
	// [18 33]
	// option1: value1
}

func ExampleRegexp_FindAllSubmatchIndex() {
	content := []byte(`
	# comment line
	option1: value1
	option2: value2
`)
	// 正则表达式模式用于从内容中捕获 "键: 值" 对。. md5:7e673e91ccc76b97
	pattern := regexp.MustCompile(`(?m)(?P<key>\w+):\s+(?P<value>\w+)$`)
	allIndexes := pattern.FindAllSubmatchIndex(content, -1)
	for _, loc := range allIndexes {
		fmt.Println(loc)
		fmt.Println(string(content[loc[0]:loc[1]]))
		fmt.Println(string(content[loc[2]:loc[3]]))
		fmt.Println(string(content[loc[4]:loc[5]]))
	}
	// Output:
	// [18 33 18 25 27 33]
	// option1: value1
	// option1
	// value1
	// [35 50 35 42 44 50]
	// option2: value2
	// option2
	// value2
}

func ExampleRegexp_FindAllIndex() {
	content := []byte("London")
	re := regexp.MustCompile(`o.`)
	fmt.Println(re.FindAllIndex(content, 1))
	fmt.Println(re.FindAllIndex(content, -1))
	// Output:
	// [[1 3]]
	// [[1 3] [4 6]]
}