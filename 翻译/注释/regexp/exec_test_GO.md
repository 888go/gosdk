
<原文开始>
// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// 版权所有 2010 Go 作者。保留所有权利。
// 本源代码的使用由可在此文件中找到的BSD风格许可证进行管理。
// md5:a7e52fd757090935
# <翻译结束>


<原文开始>
// Fatal because we'll get out of sync.
<原文结束>

# <翻译开始>
// 致命的，因为我们将会失去同步。. md5:c080815db9579e0b
# <翻译结束>


<原文开始>
// We don't and likely never will support \C; keep going.
<原文结束>

# <翻译开始>
// 我们不支持（且很可能永远不会支持）\C；继续进行。. md5:05cf365f5a0d4601
# <翻译结束>


<原文开始>
// Fatal because q worked, so this should always work.
<原文结束>

# <翻译开始>
// 致命的，因为q工作了，所以这应该始终有效。. md5:5a80e42177f75c6a
# <翻译结束>


<原文开始>
// A sequence of match results.
<原文结束>

# <翻译开始>
// 一系列比赛结果。. md5:ceb29a1a049a9fec
# <翻译结束>


<原文开始>
// Failed to compile: skip results.
<原文结束>

# <翻译开始>
// 编译失败：跳过结果。. md5:a47695f3060073b4
# <翻译结束>


<原文开始>
				// RE2's \B considers every byte position,
				// so it sees 'not word boundary' in the
				// middle of UTF-8 sequences. This package
				// only considers the positions between runes,
				// so it disagrees. Skip those cases.
<原文结束>

# <翻译开始>
				// RE2的\B会考虑每个字节位置，
				// 因此它会在UTF-8序列的中间看到"非单词边界"。这个包只考虑字符（runes）之间的位置，所以与RE2的看法不同。忽略这些情况。
				// md5:b8dc5d0311460a2a
# <翻译结束>


<原文开始>
// Protect against panic during Compile.
<原文结束>

# <翻译开始>
// 防止在编译期间发生恐慌。. md5:3b814a46143b0cca
# <翻译结束>


<原文开始>
// A single - indicates no match.
<原文结束>

# <翻译开始>
// 单个"-"表示不匹配。. md5:e55106445e602dc4
# <翻译结束>


<原文开始>
// Otherwise, a space-separated list of pairs.
<原文结束>

# <翻译开始>
// 否则，是一个由空格分隔的键值对列表。. md5:bc89e5ef73bf66f0
# <翻译结束>


<原文开始>
// Process a single pair.  - means no submatch.
<原文结束>

# <翻译开始>
// 处理单个对。- 表示没有子匹配。. md5:42972ce9608c61f8
# <翻译结束>


<原文开始>
// TestFowler runs this package's regexp API against the
// POSIX regular expression tests collected by Glenn Fowler
// at http://www2.research.att.com/~astopen/testregex/testregex.html.
<原文结束>

# <翻译开始>
// TestFowler运行此包的正则表达式API，针对Glenn Fowler在http://www2.research.att.com/~astopen/testregex/testregex.html上收集的POSIX正则表达式测试。
// md5:4526cbba9fe80f62
# <翻译结束>


<原文开始>
		// http://www2.research.att.com/~astopen/man/man1/testregex.html
		//
		// INPUT FORMAT
		//   Input lines may be blank, a comment beginning with #, or a test
		//   specification. A specification is five fields separated by one
		//   or more tabs. NULL denotes the empty string and NIL denotes the
		//   0 pointer.
<原文结束>

# <翻译开始>
		// http:		//www2.research.att.com/~astopen/man/man1/testregex.html
		//
		// 输入格式
		//   输入行可以是空白行，以#开头的注释，或是测试规格说明。一个规格说明由五个字段组成，这些字段通过一个或多个制表符分隔。NULL表示空字符串，而NIL表示0指针。
		// md5:0f45fd7ae460a781
# <翻译结束>


<原文开始>
		//   Field 1: the regex(3) flags to apply, one character per REG_feature
		//   flag. The test is skipped if REG_feature is not supported by the
		//   implementation. If the first character is not [BEASKLP] then the
		//   specification is a global control line. One or more of [BEASKLP] may be
		//   specified; the test will be repeated for each mode.
		//
		//     B 	basic			BRE	(grep, ed, sed)
		//     E 	REG_EXTENDED		ERE	(egrep)
		//     A	REG_AUGMENTED		ARE	(egrep with negation)
		//     S	REG_SHELL		SRE	(sh glob)
		//     K	REG_SHELL|REG_AUGMENTED	KRE	(ksh glob)
		//     L	REG_LITERAL		LRE	(fgrep)
		//
		//     a	REG_LEFT|REG_RIGHT	implicit ^...$
		//     b	REG_NOTBOL		lhs does not match ^
		//     c	REG_COMMENT		ignore space and #...\n
		//     d	REG_SHELL_DOT		explicit leading . match
		//     e	REG_NOTEOL		rhs does not match $
		//     f	REG_MULTIPLE		multiple \n separated patterns
		//     g	FNM_LEADING_DIR		testfnmatch only -- match until /
		//     h	REG_MULTIREF		multiple digit backref
		//     i	REG_ICASE		ignore case
		//     j	REG_SPAN		. matches \n
		//     k	REG_ESCAPE		\ to escape [...] delimiter
		//     l	REG_LEFT		implicit ^...
		//     m	REG_MINIMAL		minimal match
		//     n	REG_NEWLINE		explicit \n match
		//     o	REG_ENCLOSED		(|&) magic inside [@|&](...)
		//     p	REG_SHELL_PATH		explicit / match
		//     q	REG_DELIMITED		delimited pattern
		//     r	REG_RIGHT		implicit ...$
		//     s	REG_SHELL_ESCAPED	\ not special
		//     t	REG_MUSTDELIM		all delimiters must be specified
		//     u	standard unspecified behavior -- errors not counted
		//     v	REG_CLASS_ESCAPE	\ special inside [...]
		//     w	REG_NOSUB		no subexpression match array
		//     x	REG_LENIENT		let some errors slide
		//     y	REG_LEFT		regexec() implicit ^...
		//     z	REG_NULL		NULL subexpressions ok
		//     $	                        expand C \c escapes in fields 2 and 3
		//     /	                        field 2 is a regsubcomp() expression
		//     =	                        field 3 is a regdecomp() expression
		//
		//   Field 1 control lines:
		//
		//     C		set LC_COLLATE and LC_CTYPE to locale in field 2
		//
		//     ?test ...	output field 5 if passed and != EXPECTED, silent otherwise
		//     &test ...	output field 5 if current and previous passed
		//     |test ...	output field 5 if current passed and previous failed
		//     ; ...	output field 2 if previous failed
		//     {test ...	skip if failed until }
		//     }		end of skip
		//
		//     : comment		comment copied as output NOTE
		//     :comment:test	:comment: ignored
		//     N[OTE] comment	comment copied as output NOTE
		//     T[EST] comment	comment
		//
		//     number		use number for nmatch (20 by default)
<原文结束>

# <翻译开始>
		// 字段1：要应用的正则表达式标志，每个REG_feature标志对应一个字符。如果实现不支持REG_feature，则跳过测试。如果第一个字符不是[BELSKLP]，则说明这是一个全局控制行。可以指定一个或多个[BELSKLP]，测试将针对每种模式重复进行。
		//
		//     B 	basic			BRE	（grep, ed, sed）
		//     E 	REG_EXTENDED		ERE	（egrep）
		//     A	REG_AUGMENTED		ARE	（egrep带否定）
		//     S	REG_SHELL		SRE	（sh glob）
		//     K	REG_SHELL|REG_AUGMENTED	KRE	（ksh glob）
		//     L	REG_LITERAL		LRE	（fgrep）
		//
		//     a	REG_LEFT|REG_RIGHT	隐含^...$
		//     b	REG_NOTBOL		左侧不匹配^
		//     c	REG_COMMENT		忽略空格和#...\n
		//     d	REG_SHELL_DOT		显式匹配开头的 .
		//     e	REG_NOTEOL		右侧不匹配$
		//     f	REG_MULTIPLE		多行分隔的模式
		//     g	FNM_LEADING_DIR		testfnmatch仅匹配直到/
		//     h	REG_MULTIREF		多数字后引用
		//     i	REG_ICASE		不区分大小写
		//     j	REG_SPAN		. 匹配 \n
		//     k	REG_ESCAPE		使用\转义 [...] 分隔符
		//     l	REG_LEFT		隐含^...
		//     m	REG_MINIMAL		最小匹配
		//     n	REG_NEWLINE		显式匹配 \n
		//     o	REG_ENCLOSED		(|&) 在[@|&](...)内部有特殊含义
		//     p	REG_SHELL_PATH		显式匹配 /
		//     q	REG_DELIMITED		分隔符模式
		//     r	REG_RIGHT		隐含...$
		//     s	REG_SHELL_ESCAPED	\ 不特殊
		//     t	REG_MUSTDELIM		所有分隔符必须指定
		//     u	标准未指定行为 — 错误不计数
		//     v	REG_CLASS_ESCAPE	在[...]内\有特殊含义
		//     w	REG_NOSUB		无子表达式匹配数组
		//     x	REG_LENIENT		允许某些错误
		//     y	REG_LEFT		regexec() 隐含^...
		//     z	REG_NULL		允许空子表达式
		//     $	                        在字段2和3中展开C \c转义
		//     /	                        字段2是regsubcomp()表达式
		//     =	                        字段3是regdecomp()表达式
		//
		// 字段1控制行：
		//
		//     C		将LC_COLLATE和LC_CTYPE设置为字段2中的locale
		//
		//     ?test ...	如果通过且不等于EXPECTED，则输出字段5，否则保持沉默
		//     &test ...	如果当前和上一次都通过，则输出字段5
		//     |test ...	如果当前通过而上一次失败，则输出字段5
		//     ; ...		如果上一次失败，则输出字段2
		//     {test ...	如果失败则跳过，直到}
		//     }		结束跳过
		//
		//     : comment		复制注释作为输出的NOTE
		//     :comment:test	:comment: 忽略
		//     N[OTE] comment	复制注释作为输出的NOTE
		//     T[EST] comment	注释
		//
		//     number		使用number作为nmatch的值（默认为20）
		// md5:566579b7293133ee
# <翻译结束>


<原文开始>
			// Ignore all the control operators.
			// Just run everything.
<原文结束>

# <翻译开始>
			// 忽略所有控制运算符。
			// 只运行所有内容。
			// md5:15f2b9c91f5b5bcf
# <翻译结束>


<原文开始>
// Can check field count now that we've handled the myriad comment formats.
<原文结束>

# <翻译开始>
// 现在可以检查字段计数，因为我们已经处理了各种各样的注释格式。. md5:3e9c76d6403bdcf3
# <翻译结束>


<原文开始>
// Expand C escapes (a.k.a. Go escapes).
<原文结束>

# <翻译开始>
// 展开C风格的转义字符（也称为Go转义字符）。. md5:9418cfbebab46576
# <翻译结束>


<原文开始>
		//   Field 2: the regular expression pattern; SAME uses the pattern from
		//     the previous specification.
		//
<原文结束>

# <翻译开始>
		// 字段2：正则表达式模式；SAME 使用上一个规范中的模式。
		// md5:fad45707bad425d8
# <翻译结束>


<原文开始>
//   Field 3: the string to match.
<原文结束>

# <翻译开始>
// 第3个字段：要匹配的字符串。. md5:80ffd420a87af112
# <翻译结束>


<原文开始>
//   Field 4: the test outcome...
<原文结束>

# <翻译开始>
// 第4个字段：测试结果.... md5:740fa8a4eb45356c
# <翻译结束>


<原文开始>
//   Field 5: optional comment appended to the report.
<原文结束>

# <翻译开始>
//   字段 5：可选的注释，附加到报告中。. md5:011292d80bab76e5
# <翻译结束>


<原文开始>
// Run test once for each specified capital letter mode that we support.
<原文结束>

# <翻译开始>
// 对我们支持的每个指定的大写字母模式运行测试一次。. md5:bcd54f46649bef5f
# <翻译结束>


<原文开始>
// extended regexp (what we support)
<原文结束>

# <翻译开始>
// 扩展正则表达式（我们支持的）. md5:ce87a458a21b3b22
# <翻译结束>


<原文开始>
	//   Field 4: the test outcome. This is either one of the posix error
	//     codes (with REG_ omitted) or the match array, a list of (m,n)
	//     entries with m and n being first and last+1 positions in the
	//     field 3 string, or NULL if REG_NOSUB is in effect and success
	//     is expected. BADPAT is acceptable in place of any regcomp(3)
	//     error code. The match[] array is initialized to (-2,-2) before
	//     each test. All array elements from 0 to nmatch-1 must be specified
	//     in the outcome. Unspecified endpoints (offset -1) are denoted by ?.
	//     Unset endpoints (offset -2) are denoted by X. {x}(o:n) denotes a
	//     matched (?{...}) expression, where x is the text enclosed by {...},
	//     o is the expression ordinal counting from 1, and n is the length of
	//     the unmatched portion of the subject string. If x starts with a
	//     number then that is the return value of re_execf(), otherwise 0 is
	//     returned.
<原文结束>

# <翻译开始>
	//   字段4：测试结果。这可以是POSIX错误代码（REG_被省略）之一，或者是匹配数组，其中包含(m,n)条目，m和n是字段3字符串中的第一个和最后一个+1位置，如果使用REG_NOSUB并且期望成功，则为NULL。BADPAT可以替代任何regcomp(3)错误代码。在每次测试之前，匹配数组初始化为(-2,-2)。所有从0到nmatch-1的数组元素都必须在结果中指定。未指定的端点（偏移量-1）用?表示。未设置的端点（偏移量-2）用X表示。{x}(o:n)表示一个匹配(?{...})表达式，其中x是{...}包围的文本，o是从1开始的表达式序号，n是主体字符串中未匹配部分的长度。如果x以数字开头，则返回re_execf()的返回值，否则返回0。
	// md5:db6bc64391ced64a
# <翻译结束>


<原文开始>
// Match with no position information.
<原文结束>

# <翻译开始>
// 无位置信息的匹配。. md5:99ba8dc8b1dc1f78
# <翻译结束>


<原文开始>
// All the other error codes are compile errors.
<原文结束>

# <翻译开始>
// 所有其他错误代码均为编译错误。. md5:715e916162d60451
# <翻译结束>


<原文开始>
// TestProgramTooLongForBacktrack tests that a regex which is too long
// for the backtracker still executes properly.
<原文结束>

# <翻译开始>
// TestProgramTooLongForBacktrack 测试一个对于回溯器来说过长的正则表达式是否仍然能够正确执行。
// md5:ab0e6798be040d1f
# <翻译结束>

