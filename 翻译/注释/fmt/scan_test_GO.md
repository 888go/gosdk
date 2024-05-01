
<原文开始>
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
//版权所有2009年Go作者。所有权利保留。
//使用此源代码受BSD风格
//可以在LICENSE文件中找到的许可证。
// md5:2e9dc81828a3be8a
# <翻译结束>


<原文开始>
// Xs accepts any non-empty run of the verb character
<原文结束>

# <翻译开始>
// Xs 接受任何非空的动词字符序列. md5:e2af7ad9167fbb6a
# <翻译结束>


<原文开始>
// IntString accepts an integer followed immediately by a string.
// It tests the embedding of a scan within a scan.
<原文结束>

# <翻译开始>
// IntString 接受一个整数后面紧接着一个字符串。
// 它测试在一个扫描内部嵌套扫描的情况。
// md5:141a29b2922ced2b
# <翻译结束>


<原文开始>
// boolean test vals toggle to be sure they are written
<原文结束>

# <翻译开始>
// 将布尔测试值切换，以确保它们已写入. md5:74cbf6ee72eb57a1
# <翻译结束>


<原文开始>
// Carriage-return followed by newline. (We treat \r\n as \n always.)
<原文结束>

# <翻译开始>
// 换行符后跟一个新行。 (我们始终将\r\n视为\n). md5:6d9a989cf65156d2
# <翻译结束>


<原文开始>
// only %v takes underscores
<原文结束>

# <翻译开始>
// 只有 %v 支持使用下划线. md5:0d5e8d26bf807d19
# <翻译结束>


<原文开始>
// was: "unexpected newline"
<原文结束>

# <翻译开始>
// 原来是：“意外的换行符”. md5:64dec408d934bbe5
# <翻译结束>


<原文开始>
// was: "EOF"; 0 was taken as base prefix and not counted.
<原文结束>

# <翻译开始>
// 曾经是："EOF"；0被用作基础前缀，不计入计数。. md5:ddffebc21736e294
# <翻译结束>


<原文开始>
// %c must accept a blank.
<原文结束>

# <翻译开始>
// %c必须接受空格。. md5:6e35ac5146edbe5e
# <翻译结束>


<原文开始>
// %c must accept any space.
<原文结束>

# <翻译开始>
// %c 必须接受任何空格。. md5:7117f90647194eaf
# <翻译结束>


<原文开始>
// expected space in input to match format
<原文结束>

# <翻译开始>
// 预期在输入中有一个空格以匹配格式. md5:cbe1e1f4db29e55f
# <翻译结束>


<原文开始>
// input does not match format
<原文结束>

# <翻译开始>
// 输入不符合格式. md5:d88e7c247ff5f55c
# <翻译结束>


<原文开始>
// newline in input does not match format
<原文结束>

# <翻译开始>
// 输入中的换行符不符合格式. md5:12dd69534dfbf0a3
# <翻译结束>


<原文开始>
// Slightly odd error, but correct.
<原文结束>

# <翻译开始>
// 略显奇怪的错误，但却是正确的。. md5:3fc465eba5431502
# <翻译结束>


<原文开始>
// Bad UTF-8: should see every byte.
<原文结束>

# <翻译开始>
// 错误的UTF-8编码：应该查看每一个字节。. md5:4dad33770bdc46c7
# <翻译结束>


<原文开始>
// The incoming value may be a pointer
<原文结束>

# <翻译开始>
// 进来的值可能是指针. md5:4f1140090237710f
# <翻译结束>


<原文开始>
// different machines and different types report errors with different strings.
<原文结束>

# <翻译开始>
// 不同的机器和不同的类型会以不同的字符串报告错误。. md5:5c4eb09a7c2a976b
# <翻译结束>


<原文开始>
// Convert the slice of pointers into a slice of values
<原文结束>

# <翻译开始>
// 将指针切片转换为值切片. md5:04390588bb10d1ee
# <翻译结束>


<原文开始>
// Empty strings are not valid input when scanning a string.
<原文结束>

# <翻译开始>
// 空字符串在扫描字符串时不是有效的输入。. md5:4be6dc662c6f1d7e
# <翻译结束>


<原文开始>
// Quoted empty string is OK.
<原文结束>

# <翻译开始>
// 引用的空字符串是允许的。. md5:4022effa42f66a5a
# <翻译结束>


<原文开始>
// eofCounter is a special Reader that counts reads at end of file.
<原文结束>

# <翻译开始>
// eofCounter是一个特殊的Reader，它会在文件结尾处计数读取操作。. md5:afde126ad4be1c96
# <翻译结束>


<原文开始>
// TestEOF verifies that when we scan, we see at most EOF once per call to a
// Scan function, and then only when it's really an EOF.
<原文结束>

# <翻译开始>
// TestEOF 用于验证在扫描时，我们最多会在每次调用 Scan 函数时遇到一次 EOF，并且仅在其确实表示文件结束时出现。
// md5:26e30ed888e1cee7
# <翻译结束>


<原文开始>
// TestEOFAtEndOfInput verifies that we see an EOF error if we run out of input.
// This was a buglet: we used to get "expected integer".
<原文结束>

# <翻译开始>
// TestEOFAtEndOfInput 验证在输入结束时，我们是否会看到EOF错误。
// 这曾是一个小bug：我们以前会收到"预期的整数"错误。
// md5:d611ff7fcfef8d80
# <翻译结束>


<原文开始>
// Trailing space is tougher.
<原文结束>

# <翻译开始>
// 但是尾随空格就难了。. md5:d86bbb863034984b
# <翻译结束>


<原文开始>
// TestUnreadRuneWithBufio verifies that, at least when using bufio, successive
// calls to Fscan do not lose runes.
<原文结束>

# <翻译开始>
// TestUnreadRuneWithBufio 验证，至少在使用 bufio 时，连续调用 Fscan 不会丢失 runes。
// md5:ed3d0e5212b96ec6
# <翻译结束>


<原文开始>
// Scan attempts to read two lines into the object. Scanln should prevent this
// because it stops at newline; Scan and Scanf should be fine.
<原文结束>

# <翻译开始>
// Scan 尝试读取两行内容到对象中。Scanln 应该阻止这种情况发生，因为它在遇到换行符时就会停止；而 Scan 和 Scanf 应该可以正常处理。
// md5:25c3fd3aa3eb80ea
# <翻译结束>


<原文开始>
// Sscanln should not work
<原文结束>

# <翻译开始>
// Sscanln 应该无法正常工作. md5:0d88ad5fc8e737c7
# <翻译结束>


<原文开始>
// TestLineByLineFscanf tests that Fscanf does not read past newline. Issue
// 3481.
<原文结束>

# <翻译开始>
// TestLineByLineFscanf 测试Fscanf是否不会读取到换行符之后。问题号为3481。
// md5:c40bd9c1ad005957
# <翻译结束>


<原文开始>
// TestScanStateCount verifies the correct byte count is returned. Issue 8512.
<原文结束>

# <翻译开始>
// TestScanStateCount 验证返回的字节计数是否正确。问题 8512。. md5:854ee5084d21f394
# <翻译结束>


<原文开始>
// runeScanner implements the Scanner interface for TestScanStateCount.
<原文结束>

# <翻译开始>
// runeScanner 为 TestScanStateCount 实现了 Scanner 接口。. md5:0a264cc31ecdae6d
# <翻译结束>


<原文开始>
// RecursiveInt accepts a string matching %d.%d.%d....
// and parses it into a linked list.
// It allows us to benchmark recursive descent style scanners.
<原文结束>

# <翻译开始>
// RecursiveInt 接受一个匹配 %d.%d.%d.... 格式的字符串，
// 并将其解析为一个链表。
// 这使得我们能够基准测试递归下降类型的扫描器。
// md5:7b0a5bb167cf0a5d
# <翻译结束>


<原文开始>
// scanInts performs the same scanning task as RecursiveInt.Scan
// but without recurring through scanner, so we can compare
// performance more directly.
<原文结束>

# <翻译开始>
// scanInts 执行与 RecursiveInt.Scan 相同的扫描任务，但是不通过递归调用 scanner，这样我们可以更直接地比较性能。
// md5:3d178b4ed9a67c5e
# <翻译结束>


<原文开始>
// 800 is small enough to not overflow the stack when using gccgo on a
// platform that does not support split stack.
<原文结束>

# <翻译开始>
// 对于不支持分段栈的gccgo平台，800足够小，不会溢出栈。
// md5:e811e59fb696d96b
# <翻译结束>


<原文开始>
// Issue 9124.
// %x on bytes couldn't handle non-space bytes terminating the scan.
<原文结束>

# <翻译开始>
// Issue 9124.
// 对于字节切片，%x 在遇到非空格字节终止扫描时无法正确处理。
// md5:c5a7f6f71b59296a
# <翻译结束>


<原文开始>
	// This one fails because there is a hex byte after the data,
	// that is, an odd number of hex input bytes.
<原文结束>

# <翻译开始>
// 这个例子失败是因为数据后面有一个十六进制字节， 
// 即输入的十六进制字节数为奇数。
// md5:6bf44a19b1349df6
# <翻译结束>


<原文开始>
// fails: space after nl in input but not pattern
<原文结束>

# <翻译开始>
// 失败：输入中有换行符后跟空格，但模式中没有. md5:4ede794ccbc805f6
# <翻译结束>


<原文开始>
// Test for issue 12090: Was unreading at EOF, double-scanning a byte.
<原文结束>

# <翻译开始>
// 测试问题12090：在EOF时未读取，导致字节被扫描两次。. md5:01b05f5a0bdea342
# <翻译结束>

