
<原文开始>
// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// 版权归2012年的Go作者所有。保留所有权利。
// 使用此源代码受BSD风格的
// 许可证管理，可在LICENSE文件中找到。
// md5:a2b8441cca608eb8
# <翻译结束>


<原文开始>
	// Test for the race detector, to verify ReadAt doesn't mutate
	// any state.
<原文结束>

# <翻译开始>
// 为竞态条件检测器进行的测试，以验证 ReadAt 方法不改变任何状态。
// md5:9cd790971d166721
# <翻译结束>


<原文开始>
	// Test for the race detector, to verify a Read that doesn't yield any bytes
	// is okay to use from multiple goroutines. This was our historic behavior.
	// See golang.org/issue/7856
<原文结束>

# <翻译开始>
// 此测试用于检测竞态条件检测器，以确认多个goroutine使用一个没有读取到任何字节的Read操作是安全的。
// 这是我们历史上的行为。
// 参见golang.org/issue/7856
// md5:ec9e90217f9da7a3
# <翻译结束>


<原文开始>
// verify that copying from an empty reader always has the same results,
// regardless of the presence of a WriteTo method.
<原文结束>

# <翻译开始>
// 验证从空读取器复制始终具有相同的结果，无论是否存在WriteTo方法。
// md5:c49f3acaa8519966
# <翻译结束>


<原文开始>
// tests that Len is affected by reads, but Size is not.
<原文结束>

# <翻译开始>
// 测试Len是否受读取影响，但Size不受影响。. md5:1c7a92525f02f091
# <翻译结束>

