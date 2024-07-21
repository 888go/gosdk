
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
// make this bigger for a larger (and slower) test
<原文结束>

# <翻译开始>
// 将其增大以便进行更大（但更慢）的测试. md5:bf0c2bf33a8d9434
# <翻译结束>


<原文开始>
// test data for write tests
<原文结束>

# <翻译开始>
// 写入测试的测试数据. md5:fb12fba62e7a297f
# <翻译结束>


<原文开始>
// test data; same as testString but as a slice.
<原文结束>

# <翻译开始>
// 测试数据；与testString相同，但以切片形式表示。. md5:84211fa4bf8f2e04
# <翻译结束>


<原文开始>
// Verify that contents of buf match the string s.
<原文结束>

# <翻译开始>
// 验证buf中的内容与字符串s匹配。. md5:126cb1dbac61397b
# <翻译结束>


<原文开始>
// Fill buf through n writes of string fus.
// The initial contents of buf corresponds to the string s;
// the result is the final contents of buf returned as a string.
<原文结束>

# <翻译开始>
// 通过将字符串fus写入n次来填充buf。buf的初始内容对应于字符串s；返回的结果是buf的最终内容，以字符串形式返回。
// md5:4a2535b01b6cdaaa
# <翻译结束>


<原文开始>
// Fill buf through n writes of byte slice fub.
// The initial contents of buf corresponds to the string s;
// the result is the final contents of buf returned as a string.
<原文结束>

# <翻译开始>
// 通过n次写入字节片fub来填充buf。
// buf的初始内容对应于字符串s；
// 结果是以字符串形式返回buf的最终内容。
// md5:ffa83d80493fdd8d
# <翻译结束>


<原文开始>
// Empty buf through repeated reads into fub.
// The initial contents of buf corresponds to the string s.
<原文结束>

# <翻译开始>
// 通过反复读取到fub中清空buf。
// buf的初始内容对应于字符串s。
// md5:930d750514b1c416
# <翻译结束>


<原文开始>
// Make sure that an empty Buffer remains empty when
// it is "grown" before a Read that panics
<原文结束>

# <翻译开始>
// 确保在发生恐慌的读取之前，一个空的Buffer在增长（"grown"）后仍然保持为空。
// md5:0e205b51231f8eb6
# <翻译结束>


<原文开始>
// First verify non-panic behaviour
<原文结束>

# <翻译开始>
// 首先验证非 panic 行为. md5:d0442d9c9eefa946
# <翻译结束>


<原文开始>
// Confirm that when Reader panics, the empty buffer remains empty
<原文结束>

# <翻译开始>
// 确认当Reader发生恐慌时，空缓冲区仍然保持为空. md5:1e5afdd5554f8adb
# <翻译结束>


<原文开始>
// this is the error string of errNegativeRead
<原文结束>

# <翻译开始>
// 这是errNegativeRead错误的字符串表示. md5:1f249ec61c28b754
# <翻译结束>


<原文开始>
// With a sufficiently sized buffer, there should be no allocations.
<原文结束>

# <翻译开始>
// 如果缓冲区足够大，就无需分配内存。. md5:91e6a0658a4d1773
# <翻译结束>


<原文开始>
// Built a test slice while we write the data
<原文结束>

# <翻译开始>
// 在写入数据的同时，构建测试切片. md5:297726db76b7b32f
# <翻译结束>


<原文开始>
// Check the resulting bytes
<原文结束>

# <翻译开始>
// 检查生成的字节. md5:d96610169d4659e0
# <翻译结束>


<原文开始>
// Read it back with ReadRune
<原文结束>

# <翻译开始>
// 通过ReadRune读取回来. md5:7eff8ed0af3ba587
# <翻译结束>


<原文开始>
// Check that UnreadRune works
<原文结束>

# <翻译开始>
// 检查UnreadRune是否正常工作. md5:4bee260d2e0cfec6
# <翻译结束>


<原文开始>
	// Invalid runes, including negative ones, should be written as
	// utf8.RuneError.
<原文结束>

# <翻译开始>
	// 无效的 runes，包括负值，应写为 utf8.RuneError。
	// md5:681b779ca55ae8d0
# <翻译结束>


<原文开始>
				// 0 <= i <= j <= 5; 0 <= k <= 6
				// Check that if we start with a buffer
				// of length j at offset i and ask for
				// Next(k), we get the right bytes.
<原文结束>

# <翻译开始>
				// 0 <= i <= j <= 5; 0 <= k <= 6
				// 检查如果我们从偏移量i处的长度为j的缓冲区开始，并请求Next(k)，我们得到正确的字节。
				// md5:e2c406f6e1abd245
# <翻译结束>


<原文开始>
// If we read, this affects buf.off, which is good to test.
<原文结束>

# <翻译开始>
// 如果我们进行读取操作，这会影响到buf.off的值，这是需要测试的一个重点。. md5:de673857faeb3086
# <翻译结束>


<原文开始>
// Check no allocation occurs in write, as long as we're single-threaded.
<原文结束>

# <翻译开始>
// 检查在单线程情况下，write操作中没有发生内存分配。. md5:345a0c0930748e59
# <翻译结束>


<原文开始>
// Check that buffer has correct data.
<原文结束>

# <翻译开始>
// 检查缓冲区是否有正确的数据。. md5:7fafa1ff34cff099
# <翻译结束>


<原文开始>
// Was a bug: used to give EOF reading empty slice at EOF.
<原文结束>

# <翻译开始>
// 这是一个bug：在读取空切片时，它以前会返回EOF错误。. md5:b7f29e023c3d64c0
# <翻译结束>


<原文开始>
// after unsuccessful read
<原文结束>

# <翻译开始>
// 在读取失败后. md5:0ffe0e6d68de7072
# <翻译结束>


<原文开始>
// Tests that we occasionally compact. Issue 5154.
<原文结束>

# <翻译开始>
// 测试我们偶尔进行的压缩。问题5154。. md5:4ab768dd436e6b2d
# <翻译结束>


<原文开始>
	// (*Buffer).grow allows for 2x capacity slop before sliding,
	// so set our error threshold at 3x.
<原文结束>

# <翻译开始>
	//（*Buffer）.grow允许在滑动之前有2倍容量的冗余，所以我们将错误阈值设置为3倍。
	// md5:15593711c654c6c2
# <翻译结束>


<原文开始>
// Check that we don't compact too often. From Issue 5154.
<原文结束>

# <翻译开始>
// 检查我们不要过于频繁地进行紧凑。参考问题5154。. md5:094d39788b27155d
# <翻译结束>


<原文开始>
// use max capacity to simulate a large append operation
<原文结束>

# <翻译开始>
// 使用最大容量来模拟一个大的追加操作. md5:40dfb5f3c522b160
# <翻译结束>


<原文开始>
// should be nearly infinitely fast
<原文结束>

# <翻译开始>
// 应该几乎是无限快速的. md5:1fd3483bfdabd715
# <翻译结束>

