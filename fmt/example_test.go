// 版权所有 2017 年 The Go 语言作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可协议的约束，
// 该许可协议可在 LICENSE 文件中找到。
// md5:680e85412fa918bd

package fmt_test

import (
	"github.com/888go/gosdk/fmt"
	"io"
	"math"
	"os"
	"strings"
	"time"
)

// Errorf函数让我们可以使用格式化特性来创建描述性的错误消息。
// md5:621ce2ff5259ac95
func ExampleErrorf() {
	const name, id = "bueller", 17
	err := fmt.Errorf("user %q (id %d) not found", name, id)
	fmt.Println(err.Error())

	// Output: user "bueller" (id 17) not found
}

func ExampleFscanf() {
	var (
		i int
		b bool
		s string
	)
	r := strings.NewReader("5 true gophers")
	n, err := fmt.Fscanf(r, "%d %t %s", &i, &b, &s)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fscanf: %v\n", err)
	}
	fmt.Println(i, b, s)
	fmt.Println(n)
	// Output:
	// 5 true gophers
	// 3
}

func ExampleFscanln() {
	s := `dmr 1771 1.61803398875
	ken 271828 3.14159`
	r := strings.NewReader(s)
	var a string
	var b int
	var c float64
	for {
		n, err := fmt.Fscanln(r, &a, &b, &c)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		fmt.Printf("%d: %s, %d, %f\n", n, a, b, c)
	}
	// Output:
	// 3: dmr, 1771, 1.618034
	// 3: ken, 271828, 3.141590
}

func ExampleSscanf() {
	var name string
	var age int
	n, err := fmt.Sscanf("Kim is 22 years old", "%s is %d years old", &name, &age)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d: %s, %d\n", n, name, age)

	// Output:
	// 2: Kim, 22
}

func ExamplePrint() {
	const name, age = "Kim", 22
	fmt.Print(name, " is ", age, " years old.\n")

// 按照惯例，我们不关心 Print 返回的任何错误。
// md5:f83a05df20f85fcf

	// Output:
	// Kim is 22 years old.
}

func ExamplePrintln() {
	const name, age = "Kim", 22
	fmt.Println(name, "is", age, "years old.")

// 通常情况下，我们不需要担心Println函数返回的任何错误。
// md5:c84777cd94958383

	// Output:
	// Kim is 22 years old.
}

func ExamplePrintf() {
	const name, age = "Kim", 22
	fmt.Printf("%s is %d years old.\n", name, age)

// 通常情况下，我们不会去担心 Printf 函数返回的任何错误。
// md5:2e1e59f47661b1fd

	// Output:
	// Kim is 22 years old.
}

func ExampleSprint() {
	const name, age = "Kim", 22
	s := fmt.Sprint(name, " is ", age, " years old.\n")

	io.WriteString(os.Stdout, s) // 为了简洁，忽略错误。. md5:97b92268ec0c3dd2

	// Output:
	// Kim is 22 years old.
}

func ExampleSprintln() {
	const name, age = "Kim", 22
	s := fmt.Sprintln(name, "is", age, "years old.")

	io.WriteString(os.Stdout, s) // 为了简洁，忽略错误。. md5:97b92268ec0c3dd2

	// Output:
	// Kim is 22 years old.
}

func ExampleSprintf() {
	const name, age = "Kim", 22
	s := fmt.Sprintf("%s is %d years old.\n", name, age)

	io.WriteString(os.Stdout, s) // 为了简洁，忽略错误。. md5:97b92268ec0c3dd2

	// Output:
	// Kim is 22 years old.
}

func ExampleFprint() {
	const name, age = "Kim", 22
	n, err := fmt.Fprint(os.Stdout, name, " is ", age, " years old.\n")

// Fprint 函数返回的 n 和 err 值，
// 是来自底层 io.Writer 的返回值。
// md5:6127070feba91935
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fprint: %v\n", err)
	}
	fmt.Print(n, " bytes written.\n")

	// Output:
	// Kim is 22 years old.
	// 21 bytes written.
}

func ExampleFprintln() {
	const name, age = "Kim", 22
	n, err := fmt.Fprintln(os.Stdout, name, "is", age, "years old.")

// Fprintln 函数返回的 n 和 err 值来自其底层的 io.Writer。
// md5:ce13bb5d9c3fca1f
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fprintln: %v\n", err)
	}
	fmt.Println(n, "bytes written.")

	// Output:
	// Kim is 22 years old.
	// 21 bytes written.
}

func ExampleFprintf() {
	const name, age = "Kim", 22
	n, err := fmt.Fprintf(os.Stdout, "%s is %d years old.\n", name, age)

// Fprintf的n和err返回值与底层io.Writer的返回值相同。
// md5:e7669699d5ff7f02
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fprintf: %v\n", err)
	}
	fmt.Printf("%d bytes written.\n", n)

	// Output:
	// Kim is 22 years old.
	// 21 bytes written.
}

// Print、Println和Printf以不同的方式布局它们的参数。在这个例子中，我们可以比较它们的行为。Println总是在打印的项之间添加空格，而Print只在非字符串参数之间添加空格，Printf则严格按照指示执行。
// 
// Sprint、Sprintln、Sprintf、Fprint、Fprintln和Fprintf的行为与这里所示的对应Print、Println和Printf函数相同。
// md5:ccfea35adbcc4073
func Example_printers() {
	a, b := 3.0, 4.0
	h := math.Hypot(a, b)

// Print 在两个参数都不是字符串的情况下，在它们之间插入空格。
// 它不会向输出中添加换行符，所以我们需要显式地添加一个。
// md5:abaefe99f470b8eb
	fmt.Print("The vector (", a, b, ") has length ", h, ".\n")

// Println 总是在其参数之间插入空格，
// 因此在此情况下，它无法产生与 Print 相同的输出；
// 它的输出会有额外的空格。
// 另外，Println 总是会在输出后添加一个换行符。
// md5:a9101b6a1223d2d4
	fmt.Println("The vector (", a, b, ") has length", h, ".")

// Printf 提供了完整的控制，但使用起来更复杂。它不会在输出中添加换行符，所以我们需要在格式化字符串的末尾显式添加一个。
// 
// 使用Printf可以实现详细的控制，但其用法相对复杂。因为它不会自动在输出后添加换行符，所以在格式化字符串的最后我们需要手动添加一个。
// md5:021e04cbb39a7ef5
	fmt.Printf("The vector (%g %g) has length %g.\n", a, b, h)

	// Output:
	// The vector (3 4) has length 5.
	// The vector ( 3 4 ) has length 5 .
	// The vector (3 4) has length 5.
}

// 这些示例演示了使用格式字符串进行打印的基本方法。Printf、Sprintf 和 Fprintf 都接受一个格式字符串，该字符串指定如何格式化随后的参数。例如，%d（我们称之为“动词”）表示打印对应的参数，该参数必须是整数（或者是包含整数的类型，如整数切片）。动词 %v（'v' 代表 "值"）始终以默认形式格式化参数，就像 Print 或 Println 会显示的一样。特殊动词 %T（'T' 代表 "类型"）会打印参数的类型而不是其值。这些示例并非详尽无遗；请参阅包注释获取所有详细信息。
// md5:d850025ff01b1cdb
func Example_formats() {
// 一组基础示例展示了%v是默认的格式，在这种情况下，对于整数而言，默认格式为十进制，可以通过%d明确指定十进制格式；输出内容与Println函数生成的内容相同。
// md5:699c619e4dd1ca83
	integer := 23
	// 每个都会打印出 "23"（不包含引号）。. md5:7d9b8c8183abfbfd
	fmt.Println(integer)
	fmt.Printf("%v\n", integer)
	fmt.Printf("%d\n", integer)

	// 特殊动词%T显示的是一个项的类型，而不是它的值。. md5:4b4d50a03dfbb79d
	fmt.Printf("%T %T\n", integer, &integer)
	// Result: int *int

// Println(x) 等同于 Printf("%v\n", x)，所以在接下来的示例中，我们将只使用 Printf。每个示例都演示了如何格式化特定类型（如整数或字符串）的值。我们以 %v 开始格式字符串，以显示默认输出，然后跟随着一个或多个自定义格式。
// md5:f0011a23621f39da

	// 布林值使用%v或%t打印时，会显示为"true"或"false"。. md5:e10bcfe59bef55be
	truth := true
	fmt.Printf("%v %t\n", truth, truth)
	// Result: true true

// 整数使用%v或%d以十进制形式打印， 
// 或者使用%x以十六进制，%o以八进制，%b以二进制形式打印。
// md5:f4b84013a3a907e4
	answer := 42
	fmt.Printf("%v %d %x %o %b\n", answer, answer, answer, answer, answer)
	// 结果：42 42 2a 52 101010. md5:efef00f271c37da8

// 浮点数有多种格式：%v 和 %g 会打印紧凑的表示，而 %f 会打印小数点，%e 则使用指数表示法。这里使用的格式 %6.2f 展示了如何设置宽度和精度来控制浮点数值的显示样式。在这种情况下，6 是输出文本的总宽度（请注意输出中的额外空格），2 是要显示的小数位数。
// md5:47636bc1b061b7d9
	pi := math.Pi
	fmt.Printf("%v %g %.2f (%6.2f) %e\n", pi, pi, pi, pi, pi)
	// 结果：3.141592653589793 3.141592653589793 3.14 （3.14）3.141593e+00. md5:1871de019883dd3a

// 复数以带有'i'的浮点数括号对形式表示，其中'i'表示虚部。
// md5:9fa9fb083ebdcc08
	point := 110.7 + 22.5i
	fmt.Printf("%v %g %.2f %.2e\n", point, point, point, point)
	// 结果：(110.7+22.5i) (110.7+22.5i) (110.70+22.50i) (1.11e+02+2.25e+01i)
// 
// 这段Go语言的注释是在描述一个计算或操作的结果，它包含了复数（complex numbers），如 (110.7+22.5i)，表示实部为110.7，虚部为22.5的复数。"e+02" 表示1.11乘以10的2次方，"e+01" 表示2.25乘以10的1次方。所以整个结果包括四个类似的复数。. md5:db70e561dd734964

// runes是整数，但当使用%c进行格式化输出时，它们会显示具有相应Unicode值的字符。%q格式会将它们显示为带引号的字符，%U以十六进制Unicode代码点形式显示，而%#U则同时显示代码点和可打印的带引号形式，如果rune是可打印的。
// md5:1616f2219ffc912a
	smile := '😀'
	fmt.Printf("%v %d %c %q %U %#U\n", smile, smile, smile, smile, smile, smile)
	// Result: 128512 128512 😀 '😀' U+1F600 U+1F600 '😀'

// 字符串使用%v和%s原样格式化，%q表示引号包围的字符串，
// 而%#q表示反引号包围的字符串。
// md5:b012715b8acb956e
	placeholders := `foo "bar"`
	fmt.Printf("%v %s %q %#q\n", placeholders, placeholders, placeholders, placeholders)
	// 结果：foo "bar" foo "bar" "foo \"bar\"" `foo "bar"` 
// 
// 这段Go语言注释的中文翻译为： 
// 
// 结果：foo "bar" foo "bar" "foo \"bar\"" `foo "bar"` 
// 
// 这表示某处代码的输出或结果是四个字符串：`foo "bar"`，`foo "bar"`，`"foo \"bar\""` 和一个反斜杠`\`后面跟着`foo "bar"`。在Go语言中，反斜杠`\`通常用于转义字符串中的特殊字符。. md5:71f784de4ba6064e

// 使用%v格式化的映射会显示键和值的默认格式。形式为%#v（在这个上下文中，#称为“标志”）会以Go源代码格式显示映射。映射以一致的顺序打印，按照键的值进行排序。
// md5:c0d249062be1f15b
	isLegume := map[string]bool{
		"peanut":    true,
		"dachshund": false,
	}
	fmt.Printf("%v %#v\n", isLegume, isLegume)
	// 结果：map[dachshund: false peanut: true] map[string]bool{"dachshund": false, "peanut": true}
// 
// 这个注释描述了两个映射(map)的内容，它们表示的是相同的数据结构，只是写法略有不同。映射中存储了两个键值对：一个是"dachshund"作为键，对应的值是false；另一个是"peanut"作为键，对应的值是true。第一种写法是简化的键值对表示，第二种写法是Go语言中详细的字面量表示，明确指出了映射的键为字符串类型(string)，值为布尔类型(bool)。. md5:2dd62d79d2dfc5c5

// 使用%v格式化的结构体显示其默认格式的字段值。
// 使用%+v形式会按照字段名显示字段，而%#v则会以Go源代码格式展示结构体。
// md5:b69a0b49cb302eb1
	person := struct {
		Name string
		Age  int
	}{"Kim", 22}
	fmt.Printf("%v %+v %#v\n", person, person, person)
	// 结果：{Kim 22} {姓名：Kim，年龄：22} 
// struct { 
//     Name string; // 姓名
//     Age int    // 年龄
// 结果：{Kim 22} {姓名：Kim，年龄：22} 
// struct { 
//     Name string; // 姓名
//     Age int    // 年龄
// 结果：{Kim 22} {姓名：Kim，年龄：22} 
// struct { 
//     Name string; // 姓名
//     Age int    // 年龄
// 结果：{Kim 22} {姓名：Kim，年龄：22} 
// struct { 
//     Name string; // 姓名
//     Age int    // 年龄
// } // 使用结构体表示：{Name:"Kim", Age:22}. md5:b4044e9790bbd63b

// 指针的默认格式显示在“&”符号之前的底层值。%p动词以十六进制打印指针值。我们在这里使用类型化的nil作为%p的参数，因为任何非nil指针的值在每次运行时都会改变；你可以自行运行注释掉的Printf调用来查看。
// md5:1cb5917e88f86397
	pointer := &person
	fmt.Printf("%v %p\n", pointer, (*int)(nil))
// 结果：&{Kim 22} 0x0
// fmt.Printf("%v %p\n", pointer, pointer)
// 结果：&{Kim 22} 0x010203 // 参见上面的注释。
// md5:c03b8440e69fdcf6

	// 数组和切片通过将格式应用到每个元素来进行格式化。. md5:f3ff38e06bfa3570
	greats := [5]string{"Kitano", "Kobayashi", "Kurosawa", "Miyazaki", "Ozu"}
	fmt.Printf("%v %q\n", greats, greats)
	// 结果：[北野武 黑泽明 宫崎骏 小津安二郎] 
// ["Kitano" "Kobayashi" "Kurosawa" "Miyazaki" "Ozu"] 
// 
// 这段Go语言注释表示一个数组或变量的结果包含了五个导演的名字，分别是北野武（Kitano）、黑泽明（Kurosawa）、宫崎骏（Miyazaki）、小津安二郎（Ozu）。. md5:64bf8dcff4cffadf

	kGreats := greats[:3]
	fmt.Printf("%v %q %#v\n", kGreats, kGreats, kGreats)
	// 结果：[小林克己 黑泽明] ["Kitano" "Kobayashi" "Kurosawa"] []string{"Kitano", "Kobayashi", "Kurosawa"} 
// 
// 这段Go代码的注释表示：该部分的输出或结果是一个切片（slice），包含三个元素，分别是小林克己（Kitano）、黑泽明（Kurosawa）。这些元素都是字符串类型，存储在名为`["Kitano" "Kobayashi" "Kurosawa"]`的数组中，转换为Go语言的切片形式是`[]string{"Kitano", "Kobayashi", "Kurosawa"}`。. md5:c42fb31071436e43

// 字节切片是特殊的。像%d这样的整数 verbs 会以该格式打印元素。%s 和 %q 形式的 verbs 将切片视为字符串处理。%x verb 有一个带有空格标志的特殊形式，它在字节之间插入一个空格。
// md5:45dc825baff70182
	cmd := []byte("a⌘")
	fmt.Printf("%v %d %s %q %x % x\n", cmd, cmd, cmd, cmd, cmd, cmd)
	// Result: [97 226 140 152] [97 226 140 152] a⌘ "a⌘" 61e28c98 61 e2 8c 98

// 实现Stringer接口的类型会被打印为字符串。因为Stringer返回的是一个字符串，我们可以使用特定于字符串的动词（如%q）来打印它们。
// md5:b3cf9c5eba768d17
	now := time.Unix(123456789, 0).UTC() // time.Time实现了fmt.Stringer接口。. md5:649e2164a704ba4e
	fmt.Printf("%v %q\n", now, now)
	// 结果：1973年11月29日 21:33:09 +0000 UTC "1973-11-29 21:33:09 +0000 UTC". md5:0bd16b5fbfdddf8b

	// Output:
	// 23
	// 23
	// 23
	// int *int
	// true true
	// 42 42 2a 52 101010
	// 3.141592653589793 3.141592653589793 3.14 (  3.14) 3.141593e+00
	// (110.7+22.5i) (110.7+22.5i) (110.70+22.50i) (1.11e+02+2.25e+01i)
	// 128512 128512 😀 '😀' U+1F600 U+1F600 '😀'
	// foo "bar" foo "bar" "foo \"bar\"" `foo "bar"`
	// map[dachshund:false peanut:true] map[string]bool{"dachshund":false, "peanut":true}
	// {Kim 22} {Name:Kim Age:22} struct { Name string; Age int }{Name:"Kim", Age:22}
	// &{Kim 22} 0x0
	// [Kitano Kobayashi Kurosawa Miyazaki Ozu] ["Kitano" "Kobayashi" "Kurosawa" "Miyazaki" "Ozu"]
	// [Kitano Kobayashi Kurosawa] ["Kitano" "Kobayashi" "Kurosawa"] []string{"Kitano", "Kobayashi", "Kurosawa"}
	// [97 226 140 152] [97 226 140 152] a⌘ "a⌘" 61e28c98 61 e2 8c 98
	// 1973-11-29 21:33:09 +0000 UTC "1973-11-29 21:33:09 +0000 UTC"
}
