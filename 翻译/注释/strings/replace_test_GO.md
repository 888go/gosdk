
<原文开始>
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// 版权所有 ? 2009 Go作者。保留所有权利。
// 本源代码的使用受BSD风格
// 许可证约束，该许可证可在LICENSE文件中找到。
# <翻译结束>


<原文开始>
// The http package's old HTML escaping function.
<原文结束>

# <翻译开始>
// http 包中的旧版 HTML 转义函数。
# <翻译结束>


<原文开始>
// TestReplacer tests the replacer implementations.
<原文结束>

# <翻译开始>
// TestReplacer 测试替换器实现。
# <翻译结束>


<原文开始>
// str converts 0xff to "\xff". This isn't just string(b) since that converts to UTF-8.
<原文结束>

# <翻译开始>
// str 将 0xff 转换为 "\xff"。这不仅仅是使用 string(b) 进行转换，因为后者会转为 UTF-8 格式。
# <翻译结束>


<原文开始>
// inc maps "\x00"->"\x01", ..., "a"->"b", "b"->"c", ..., "\xff"->"\x00".
<原文结束>

# <翻译开始>
// inc 将字符按以下规则进行映射：
//   "\x00" -> "\x01"
//   ...
//   "a" -> "b"
//   "b" -> "c"
//   ...
//   "\xff" -> "\x00"
# <翻译结束>


<原文开始>
// Test cases with 1-byte old strings, 1-byte new strings.
<原文结束>

# <翻译开始>
// 测试用例，包含1字节的旧字符串和1字节的新字符串
# <翻译结束>


<原文开始>
// repeat maps "a"->"a", "b"->"bb", "c"->"ccc", ...
<原文结束>

# <翻译开始>
// repeat 将 "a" 映射为 "a"，"b" 映射为 "bb"，"c" 映射为 "ccc"，以此类推
# <翻译结束>


<原文开始>
// Test cases with 1-byte old strings, variable length new strings.
<原文结束>

# <翻译开始>
// 测试用例，含1字节长度的旧字符串，及可变长度的新字符串。
# <翻译结束>


<原文开始>
// The remaining test cases have variable length old strings.
<原文结束>

# <翻译开始>
// 剩余的测试用例具有可变长度的旧字符串。
# <翻译结束>


<原文开始>
	// gen1 has multiple old strings of variable length. There is no
	// overall non-empty common prefix, but some pairwise common prefixes.
<原文结束>

# <翻译开始>
	// gen1 包含多个长度可变的旧字符串。它们之间不存在非空的整体公共前缀，但存在一些两两之间的公共前缀。
# <翻译结束>


<原文开始>
// gen2 has multiple old strings with no pairwise common prefix.
<原文结束>

# <翻译开始>
// gen2包含多个无两两共同前缀的旧字符串。
# <翻译结束>


<原文开始>
// gen3 has multiple old strings with an overall common prefix.
<原文结束>

# <翻译开始>
// gen3 含有多个具有共同前缀的旧字符串
# <翻译结束>


<原文开始>
	// foo{1,2,3,4} have multiple old strings with an overall common prefix
	// and 1- or 2- byte extensions from the common prefix.
<原文结束>

# <翻译开始>
	// foo{1,2,3,4}具有多个具有共同前缀的旧字符串，且从该公共前缀延伸出1个或2个字节。
# <翻译结束>


<原文开始>
// genAll maps "\x00\x01\x02...\xfe\xff" to "[all]", amongst other things.
<原文结束>

# <翻译开始>
// genAll 将 "\x00\x01\x02...\xfe\xff" 映射为 "[all]"，同时执行其他操作。
# <翻译结束>


<原文开始>
// Test cases with empty old strings.
<原文结束>

# <翻译开始>
// 测试用例，其中包含空的旧字符串。
# <翻译结束>


<原文开始>
// Issue 6659 cases (more single string replacer)
<原文结束>

# <翻译开始>
// Issue 6659 示例（更多单字符串替换器）
# <翻译结束>


<原文开始>
// TestPickAlgorithm tests that NewReplacer picks the correct algorithm.
<原文结束>

# <翻译开始>
// TestPickAlgorithm 测试 NewReplacer 选择正确算法的情况
# <翻译结束>


<原文开始>
// TestWriteStringError tests that WriteString returns an error
// received from the underlying io.Writer.
<原文结束>

# <翻译开始>
// TestWriteStringError 测试 WriteString 函数在接收到底层 io.Writer 传回的错误时能够正常返回该错误
# <翻译结束>


<原文开始>
// TestGenericTrieBuilding verifies the structure of the generated trie. There
// is one node per line, and the key ending with the current line is in the
// trie if it ends with a "+".
<原文结束>

# <翻译开始>
// TestGenericTrieBuilding 验证生成的字典树结构。每一行代表一个节点，若当前行以 "+" 结尾，则表示以该行结束的键存在于字典树中。
# <翻译结束>


<原文开始>
// Remove tabs from tc.out
<原文结束>

# <翻译开始>
// 从tc.out中移除制表符
# <翻译结束>


<原文开始>
// varying lengths forces generic
<原文结束>

# <翻译开始>
// 不同长度强制使用泛型
# <翻译结束>


<原文开始>
// BenchmarkByteByteReplaces compares byteByteImpl against multiple Replaces.
<原文结束>

# <翻译开始>
// BenchmarkByteByteReplaces对比byteByteImpl与多个Replaces的性能。
# <翻译结束>


<原文开始>
// BenchmarkByteByteMap compares byteByteImpl against Map.
<原文结束>

# <翻译开始>
// BenchmarkByteByteMap 对比 byteByteImpl 与 Map 的性能
# <翻译结束>


<原文开始>
// TestPickAlgorithm tests that NewReplacer picks the correct algorithm.
//func TestPickAlgorithm(t *testing.T) {
//	for i, tc := range algorithmTestCases {
//		got := fmt.Sprintf("%T", tc.r.Replacer())
//		if got != tc.want {
//			t.Errorf("%d. algorithm = %s, want %s", i, got, tc.want)
//		}
//	}
//}
<原文结束>

# <翻译开始>
// TestPickAlgorithm 测试 NewReplacer 函数是否选择了正确的算法。
//func TestPickAlgorithm(t *testing.T) {
//	for i, 测试用例tc := range 算法测试用例 {
//		获取的算法类型 := fmt.Sprintf("%T", tc.r.Replacer())
//		if 获取的算法类型 != tc.want {
//			t.Errorf("%d. 算法 = %s, 需要 %s", i, 获取的算法类型, tc.want)
//		}
//	}
//}
// md5:4cd37c0269abec05
# <翻译结束>


<原文开始>
// TestGenericTrieBuilding verifies the structure of the generated trie. There
// is one node per line, and the key ending with the current line is in the
// trie if it ends with a "+".
//func TestGenericTrieBuilding(t *testing.T) {
//	testCases := []struct{ in, out string }{
//		{"abc;abdef;abdefgh;xx;xy;z", `-
//			a-
//			.b-
//			..c+
//			..d-
//			...ef+
//			.....gh+
//			x-
//			.x+
//			.y+
//			z+
//			`},
//		{"abracadabra;abracadabrakazam;abraham;abrasion", `-
//			a-
//			.bra-
//			....c-
//			.....adabra+
//			...........kazam+
//			....h-
//			.....am+
//			....s-
//			.....ion+
//			`},
//		{"aaa;aa;a;i;longerst;longer;long;xx;x;X;Y", `-
//			X+
//			Y+
//			a+
//			.a+
//			..a+
//			i+
//			l-
//			.ong+
//			....er+
//			......st+
//			x+
//			.x+
//			`},
//		{"foo;;foo;foo1", `+
//			f-
//			.oo+
//			...1+
//			`},
//	}
//
//	for _, tc := range testCases {
//		keys := Split(tc.in, ";")
//		args := make([]string, len(keys)*2)
//		for i, key := range keys {
//			args[i*2] = key
//		}
//
//		got := NewReplacer(args...).PrintTrie()
//		// Remove tabs from tc.out
//		wantbuf := make([]byte, 0, len(tc.out))
//		for i := 0; i < len(tc.out); i++ {
//			if tc.out[i] != '\t' {
//				wantbuf = append(wantbuf, tc.out[i])
//			}
//		}
//		want := string(wantbuf)
//
//		if got != want {
//			t.Errorf("PrintTrie(%q)\ngot\n%swant\n%s", tc.in, got, want)
//		}
//	}
//}
<原文结束>

# <翻译开始>
// TestGenericTrieBuilding 验证生成的字典树结构。每一行代表一个节点，如果当前行末尾的键以"+"结尾，则表示该键存在于字典树中。
//func TestGenericTrieBuilding(t *testing.T) {
// 测试用例：
//tc := []struct{ in, out string }{
// 	{"abc;abdef;abdefgh;xx;xy;z", `-  
//			a-  
//			.b-  
//			..c+  
//			..d-  
//			...ef+  
//			.....gh+  
//			x-  
//			.x+  
//			.y+  
//			z+  
//			`},
// 	{"abracadabra;abracadabrakazam;abraham;abrasion", `-  
//			a-  
//			.bra-  
//			....c-  
//			.....adabra+  
//			...........kazam+  
//			....h-  
//			.....am+  
//			....s-  
//			.....ion+  
//			`},
// 	{"aaa;aa;a;i;longerst;longer;long;xx;x;X;Y", `-  
//			X+  
//			Y+  
//			a+  
//			.a+  
//			..a+  
//			i+  
//			l-  
//			.ong+  
//			....er+  
//			......st+  
//			x+  
//			.x+  
//			`},
// 	{"foo;;foo;foo1", `+  
//			f-  
//			.oo+  
//			...1+  
//			`},
// }
//
// 对于每个测试用例，我们首先将输入字符串按";"分割成键列表，然后创建相应的参数数组。接着，使用这些参数构建字典树，并获取打印结果。
// 
// 去掉输出字符串中的制表符，将其转换为字节切片。
// 然后比较实际得到的打印结果与期望结果是否一致，如果不一致则输出错误信息。
//}
// md5:cd6c8519103bddf4
# <翻译结束>

