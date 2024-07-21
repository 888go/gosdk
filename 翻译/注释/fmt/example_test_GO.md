
<原文开始>
// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// 版权所有 2017 年 The Go 语言作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可协议的约束，
// 该许可协议可在 LICENSE 文件中找到。
// md5:680e85412fa918bd
# <翻译结束>


<原文开始>
// The Errorf function lets us use formatting features
// to create descriptive error messages.
<原文结束>

# <翻译开始>
// Errorf函数让我们可以使用格式化特性来创建描述性的错误消息。
// md5:621ce2ff5259ac95
# <翻译结束>


<原文开始>
	// It is conventional not to worry about any
	// error returned by Print.
<原文结束>

# <翻译开始>
	// 按照惯例，我们不关心 Print 返回的任何错误。
	// md5:f83a05df20f85fcf
# <翻译结束>


<原文开始>
	// It is conventional not to worry about any
	// error returned by Println.
<原文结束>

# <翻译开始>
	// 通常情况下，我们不需要担心Println函数返回的任何错误。
	// md5:c84777cd94958383
# <翻译结束>


<原文开始>
	// It is conventional not to worry about any
	// error returned by Printf.
<原文结束>

# <翻译开始>
	// 通常情况下，我们不会去担心 Printf 函数返回的任何错误。
	// md5:2e1e59f47661b1fd
# <翻译结束>


<原文开始>
// Ignoring error for simplicity.
<原文结束>

# <翻译开始>
// 为了简洁，忽略错误。. md5:97b92268ec0c3dd2
# <翻译结束>


<原文开始>
	// The n and err return values from Fprint are
	// those returned by the underlying io.Writer.
<原文结束>

# <翻译开始>
	// Fprint 函数返回的 n 和 err 值，
	// 是来自底层 io.Writer 的返回值。
	// md5:6127070feba91935
# <翻译结束>


<原文开始>
	// The n and err return values from Fprintln are
	// those returned by the underlying io.Writer.
<原文结束>

# <翻译开始>
	// Fprintln 函数返回的 n 和 err 值来自其底层的 io.Writer。
	// md5:ce13bb5d9c3fca1f
# <翻译结束>


<原文开始>
	// The n and err return values from Fprintf are
	// those returned by the underlying io.Writer.
<原文结束>

# <翻译开始>
	// Fprintf的n和err返回值与底层io.Writer的返回值相同。
	// md5:e7669699d5ff7f02
# <翻译结束>


<原文开始>
// Print, Println, and Printf lay out their arguments differently. In this example
// we can compare their behaviors. Println always adds blanks between the items it
// prints, while Print adds blanks only between non-string arguments and Printf
// does exactly what it is told.
// Sprint, Sprintln, Sprintf, Fprint, Fprintln, and Fprintf behave the same as
// their corresponding Print, Println, and Printf functions shown here.
<原文结束>

# <翻译开始>
// Print、Println和Printf以不同的方式布局它们的参数。在这个例子中，我们可以比较它们的行为。Println总是在打印的项之间添加空格，而Print只在非字符串参数之间添加空格，Printf则严格按照指示执行。
// 
// Sprint、Sprintln、Sprintf、Fprint、Fprintln和Fprintf的行为与这里所示的对应Print、Println和Printf函数相同。
// md5:ccfea35adbcc4073
# <翻译结束>


<原文开始>
	// Print inserts blanks between arguments when neither is a string.
	// It does not add a newline to the output, so we add one explicitly.
<原文结束>

# <翻译开始>
	// Print 在两个参数都不是字符串的情况下，在它们之间插入空格。
	// 它不会向输出中添加换行符，所以我们需要显式地添加一个。
	// md5:abaefe99f470b8eb
# <翻译结束>


<原文开始>
	// Println always inserts spaces between its arguments,
	// so it cannot be used to produce the same output as Print in this case;
	// its output has extra spaces.
	// Also, Println always adds a newline to the output.
<原文结束>

# <翻译开始>
	// Println 总是在其参数之间插入空格，
	// 因此在此情况下，它无法产生与 Print 相同的输出；
	// 它的输出会有额外的空格。
	// 另外，Println 总是会在输出后添加一个换行符。
	// md5:a9101b6a1223d2d4
# <翻译结束>


<原文开始>
	// Printf provides complete control but is more complex to use.
	// It does not add a newline to the output, so we add one explicitly
	// at the end of the format specifier string.
<原文结束>

# <翻译开始>
	// Printf 提供了完整的控制，但使用起来更复杂。它不会在输出中添加换行符，所以我们需要在格式化字符串的末尾显式添加一个。
	// 
	// 使用Printf可以实现详细的控制，但其用法相对复杂。因为它不会自动在输出后添加换行符，所以在格式化字符串的最后我们需要手动添加一个。
	// md5:021e04cbb39a7ef5
# <翻译结束>


<原文开始>
// These examples demonstrate the basics of printing using a format string. Printf,
// Sprintf, and Fprintf all take a format string that specifies how to format the
// subsequent arguments. For example, %d (we call that a 'verb') says to print the
// corresponding argument, which must be an integer (or something containing an
// integer, such as a slice of ints) in decimal. The verb %v ('v' for 'value')
// always formats the argument in its default form, just how Print or Println would
// show it. The special verb %T ('T' for 'Type') prints the type of the argument
// rather than its value. The examples are not exhaustive; see the package comment
// for all the details.
<原文结束>

# <翻译开始>
// 这些示例演示了使用格式字符串进行打印的基本方法。Printf、Sprintf 和 Fprintf 都接受一个格式字符串，该字符串指定如何格式化随后的参数。例如，%d（我们称之为“动词”）表示打印对应的参数，该参数必须是整数（或者是包含整数的类型，如整数切片）。动词 %v（'v' 代表 "值"）始终以默认形式格式化参数，就像 Print 或 Println 会显示的一样。特殊动词 %T（'T' 代表 "类型"）会打印参数的类型而不是其值。这些示例并非详尽无遗；请参阅包注释获取所有详细信息。
// md5:d850025ff01b1cdb
# <翻译结束>


<原文开始>
	// A basic set of examples showing that %v is the default format, in this
	// case decimal for integers, which can be explicitly requested with %d;
	// the output is just what Println generates.
<原文结束>

# <翻译开始>
	// 一组基础示例展示了%v是默认的格式，在这种情况下，对于整数而言，默认格式为十进制，可以通过%d明确指定十进制格式；输出内容与Println函数生成的内容相同。
	// md5:699c619e4dd1ca83
# <翻译结束>


<原文开始>
// Each of these prints "23" (without the quotes).
<原文结束>

# <翻译开始>
// 每个都会打印出 "23"（不包含引号）。. md5:7d9b8c8183abfbfd
# <翻译结束>


<原文开始>
// The special verb %T shows the type of an item rather than its value.
<原文结束>

# <翻译开始>
// 特殊动词%T显示的是一个项的类型，而不是它的值。. md5:4b4d50a03dfbb79d
# <翻译结束>


<原文开始>
	// Println(x) is the same as Printf("%v\n", x) so we will use only Printf
	// in the following examples. Each one demonstrates how to format values of
	// a particular type, such as integers or strings. We start each format
	// string with %v to show the default output and follow that with one or
	// more custom formats.
<原文结束>

# <翻译开始>
	// Println(x) 等同于 Printf("%v\n", x)，所以在接下来的示例中，我们将只使用 Printf。每个示例都演示了如何格式化特定类型（如整数或字符串）的值。我们以 %v 开始格式字符串，以显示默认输出，然后跟随着一个或多个自定义格式。
	// md5:f0011a23621f39da
# <翻译结束>


<原文开始>
// Booleans print as "true" or "false" with %v or %t.
<原文结束>

# <翻译开始>
// 布林值使用%v或%t打印时，会显示为"true"或"false"。. md5:e10bcfe59bef55be
# <翻译结束>


<原文开始>
	// Integers print as decimals with %v and %d,
	// or in hex with %x, octal with %o, or binary with %b.
<原文结束>

# <翻译开始>
	// 整数使用%v或%d以十进制形式打印， 
	// 或者使用%x以十六进制，%o以八进制，%b以二进制形式打印。
	// md5:f4b84013a3a907e4
# <翻译结束>


<原文开始>
// Result: 42 42 2a 52 101010
<原文结束>

# <翻译开始>
// 结果：42 42 2a 52 101010. md5:efef00f271c37da8
# <翻译结束>


<原文开始>
	// Floats have multiple formats: %v and %g print a compact representation,
	// while %f prints a decimal point and %e uses exponential notation. The
	// format %6.2f used here shows how to set the width and precision to
	// control the appearance of a floating-point value. In this instance, 6 is
	// the total width of the printed text for the value (note the extra spaces
	// in the output) and 2 is the number of decimal places to show.
<原文结束>

# <翻译开始>
	// 浮点数有多种格式：%v 和 %g 会打印紧凑的表示，而 %f 会打印小数点，%e 则使用指数表示法。这里使用的格式 %6.2f 展示了如何设置宽度和精度来控制浮点数值的显示样式。在这种情况下，6 是输出文本的总宽度（请注意输出中的额外空格），2 是要显示的小数位数。
	// md5:47636bc1b061b7d9
# <翻译结束>


<原文开始>
// Result: 3.141592653589793 3.141592653589793 3.14 (  3.14) 3.141593e+00
<原文结束>

# <翻译开始>
// 结果：3.141592653589793 3.141592653589793 3.14 （3.14）3.141593e+00. md5:1871de019883dd3a
# <翻译结束>


<原文开始>
	// Complex numbers format as parenthesized pairs of floats, with an 'i'
	// after the imaginary part.
<原文结束>

# <翻译开始>
	// 复数以带有'i'的浮点数括号对形式表示，其中'i'表示虚部。
	// md5:9fa9fb083ebdcc08
# <翻译结束>


<原文开始>
// Result: (110.7+22.5i) (110.7+22.5i) (110.70+22.50i) (1.11e+02+2.25e+01i)
<原文结束>

# <翻译开始>
// 结果：(110.7+22.5i) (110.7+22.5i) (110.70+22.50i) (1.11e+02+2.25e+01i)
// 
// 这段Go语言的注释是在描述一个计算或操作的结果，它包含了复数（complex numbers），如 (110.7+22.5i)，表示实部为110.7，虚部为22.5的复数。"e+02" 表示1.11乘以10的2次方，"e+01" 表示2.25乘以10的1次方。所以整个结果包括四个类似的复数。. md5:db70e561dd734964
# <翻译结束>


<原文开始>
	// Runes are integers but when printed with %c show the character with that
	// Unicode value. The %q verb shows them as quoted characters, %U as a
	// hex Unicode code point, and %#U as both a code point and a quoted
	// printable form if the rune is printable.
<原文结束>

# <翻译开始>
	// runes是整数，但当使用%c进行格式化输出时，它们会显示具有相应Unicode值的字符。%q格式会将它们显示为带引号的字符，%U以十六进制Unicode代码点形式显示，而%#U则同时显示代码点和可打印的带引号形式，如果rune是可打印的。
	// md5:1616f2219ffc912a
# <翻译结束>


<原文开始>
// Result: 128512 128512 ?? '??' U+1F600 U+1F600 '??'
<原文结束>

# <翻译开始>
// 结果：128512 128512 ?? '馃榾' U+1F600 U+1F600 '馃榾'. md5:9d02ad1ee15c865d
# <翻译结束>


<原文开始>
	// Strings are formatted with %v and %s as-is, with %q as quoted strings,
	// and %#q as backquoted strings.
<原文结束>

# <翻译开始>
	// 字符串使用%v和%s原样格式化，%q表示引号包围的字符串，
	// 而%#q表示反引号包围的字符串。
	// md5:b012715b8acb956e
# <翻译结束>


<原文开始>
// Result: foo "bar" foo "bar" "foo \"bar\"" `foo "bar"`
<原文结束>

# <翻译开始>
// 结果：foo "bar" foo "bar" "foo \"bar\"" `foo "bar"` 
// 
// 这段Go语言注释的中文翻译为： 
// 
// 结果：foo "bar" foo "bar" "foo \"bar\"" `foo "bar"` 
// 
// 这表示某处代码的输出或结果是四个字符串：`foo "bar"`，`foo "bar"`，`"foo \"bar\""` 和一个反斜杠`\`后面跟着`foo "bar"`。在Go语言中，反斜杠`\`通常用于转义字符串中的特殊字符。. md5:71f784de4ba6064e
# <翻译结束>


<原文开始>
	// Maps formatted with %v show keys and values in their default formats.
	// The %#v form (the # is called a "flag" in this context) shows the map in
	// the Go source format. Maps are printed in a consistent order, sorted
	// by the values of the keys.
<原文结束>

# <翻译开始>
	// 使用%v格式化的映射会显示键和值的默认格式。形式为%#v（在这个上下文中，#称为“标志”）会以Go源代码格式显示映射。映射以一致的顺序打印，按照键的值进行排序。
	// md5:c0d249062be1f15b
# <翻译结束>


<原文开始>
// Result: map[dachshund:false peanut:true] map[string]bool{"dachshund":false, "peanut":true}
<原文结束>

# <翻译开始>
// 结果：map[dachshund: false peanut: true] map[string]bool{"dachshund": false, "peanut": true}
// 
// 这个注释描述了两个映射(map)的内容，它们表示的是相同的数据结构，只是写法略有不同。映射中存储了两个键值对：一个是"dachshund"作为键，对应的值是false；另一个是"peanut"作为键，对应的值是true。第一种写法是简化的键值对表示，第二种写法是Go语言中详细的字面量表示，明确指出了映射的键为字符串类型(string)，值为布尔类型(bool)。. md5:2dd62d79d2dfc5c5
# <翻译结束>


<原文开始>
	// Structs formatted with %v show field values in their default formats.
	// The %+v form shows the fields by name, while %#v formats the struct in
	// Go source format.
<原文结束>

# <翻译开始>
	// 使用%v格式化的结构体显示其默认格式的字段值。
	// 使用%+v形式会按照字段名显示字段，而%#v则会以Go源代码格式展示结构体。
	// md5:b69a0b49cb302eb1
# <翻译结束>


<原文开始>
// Result: {Kim 22} {Name:Kim Age:22} struct { Name string; Age int }{Name:"Kim", Age:22}
<原文结束>

# <翻译开始>
// 结果：{Kim 22} {姓名：Kim，年龄：22} 
// struct { 
//     Name string; // 姓名
//     Age int    // 年龄
// } // 使用结构体表示：{Name:"Kim", Age:22}. md5:b4044e9790bbd63b
# <翻译结束>


<原文开始>
	// The default format for a pointer shows the underlying value preceded by
	// an ampersand. The %p verb prints the pointer value in hex. We use a
	// typed nil for the argument to %p here because the value of any non-nil
	// pointer would change from run to run; run the commented-out Printf
	// call yourself to see.
<原文结束>

# <翻译开始>
	// 指针的默认格式显示在“&”符号之前的底层值。%p动词以十六进制打印指针值。我们在这里使用类型化的nil作为%p的参数，因为任何非nil指针的值在每次运行时都会改变；你可以自行运行注释掉的Printf调用来查看。
	// md5:1cb5917e88f86397
# <翻译结束>


<原文开始>
	// Result: &{Kim 22} 0x0
	// fmt.Printf("%v %p\n", pointer, pointer)
	// Result: &{Kim 22} 0x010203 // See comment above.
<原文结束>

# <翻译开始>
	// 结果：&{Kim 22} 0x0
	// fmt.Printf("%v %p\n", pointer, pointer)
	// 结果：&{Kim 22} 0x010203 	// 参见上面的注释。
	// md5:c03b8440e69fdcf6
# <翻译结束>


<原文开始>
// Arrays and slices are formatted by applying the format to each element.
<原文结束>

# <翻译开始>
// 数组和切片通过将格式应用到每个元素来进行格式化。. md5:f3ff38e06bfa3570
# <翻译结束>


<原文开始>
// Result: [Kitano Kobayashi Kurosawa Miyazaki Ozu] ["Kitano" "Kobayashi" "Kurosawa" "Miyazaki" "Ozu"]
<原文结束>

# <翻译开始>
// 结果：[北野武 黑泽明 宫崎骏 小津安二郎] 
// ["Kitano" "Kobayashi" "Kurosawa" "Miyazaki" "Ozu"] 
// 
// 这段Go语言注释表示一个数组或变量的结果包含了五个导演的名字，分别是北野武（Kitano）、黑泽明（Kurosawa）、宫崎骏（Miyazaki）、小津安二郎（Ozu）。. md5:64bf8dcff4cffadf
# <翻译结束>


<原文开始>
// Result: [Kitano Kobayashi Kurosawa] ["Kitano" "Kobayashi" "Kurosawa"] []string{"Kitano", "Kobayashi", "Kurosawa"}
<原文结束>

# <翻译开始>
// 结果：[小林克己 黑泽明] ["Kitano" "Kobayashi" "Kurosawa"] []string{"Kitano", "Kobayashi", "Kurosawa"} 
// 
// 这段Go代码的注释表示：该部分的输出或结果是一个切片（slice），包含三个元素，分别是小林克己（Kitano）、黑泽明（Kurosawa）。这些元素都是字符串类型，存储在名为`["Kitano" "Kobayashi" "Kurosawa"]`的数组中，转换为Go语言的切片形式是`[]string{"Kitano", "Kobayashi", "Kurosawa"}`。. md5:c42fb31071436e43
# <翻译结束>


<原文开始>
	// Byte slices are special. Integer verbs like %d print the elements in
	// that format. The %s and %q forms treat the slice like a string. The %x
	// verb has a special form with the space flag that puts a space between
	// the bytes.
<原文结束>

# <翻译开始>
	// 字节切片是特殊的。像%d这样的整数 verbs 会以该格式打印元素。%s 和 %q 形式的 verbs 将切片视为字符串处理。%x verb 有一个带有空格标志的特殊形式，它在字节之间插入一个空格。
	// md5:45dc825baff70182
# <翻译结束>


<原文开始>
// Result: [97 226 140 152] [97 226 140 152] a? "a?" 61e28c98 61 e2 8c 98
<原文结束>

# <翻译开始>
// 结果: [97 226 140 152] [97 226 140 152] a? "a?" 61e28c98 61 e2 8c 98. md5:0ff6661bdb700fa0
# <翻译结束>


<原文开始>
	// Types that implement Stringer are printed the same as strings. Because
	// Stringers return a string, we can print them using a string-specific
	// verb such as %q.
<原文结束>

# <翻译开始>
	// 实现Stringer接口的类型会被打印为字符串。因为Stringer返回的是一个字符串，我们可以使用特定于字符串的动词（如%q）来打印它们。
	// md5:b3cf9c5eba768d17
# <翻译结束>


<原文开始>
// time.Time implements fmt.Stringer.
<原文结束>

# <翻译开始>
// time.Time实现了fmt.Stringer接口。. md5:649e2164a704ba4e
# <翻译结束>


<原文开始>
// Result: 1973-11-29 21:33:09 +0000 UTC "1973-11-29 21:33:09 +0000 UTC"
<原文结束>

# <翻译开始>
// 结果：1973年11月29日 21:33:09 +0000 UTC "1973-11-29 21:33:09 +0000 UTC". md5:0bd16b5fbfdddf8b
# <翻译结束>

