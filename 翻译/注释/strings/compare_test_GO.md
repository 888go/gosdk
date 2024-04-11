
<原文开始>
// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// 版权所有 2013 The Go Authors。保留所有权利。
// 使用本源代码受 BSD 风格许可协议约束，
// 该协议可在 LICENSE 文件中找到。
# <翻译结束>


<原文开始>
// Derived from bytes/compare_test.go.
// Benchmarks omitted since the underlying implementation is identical.
<原文结束>

# <翻译开始>
// 源自 bytes/compare_test.go。
// 由于底层实现完全相同，故省略基准测试。
# <翻译结束>


<原文开始>
	// unsafeString converts a []byte to a string with no allocation.
	// The caller must not modify b while the result string is in use.
<原文结束>

# <翻译开始>
// unsafeString将[]byte无分配地转换为string。
// 在结果字符串使用期间，调用者不得修改b。
# <翻译结束>


<原文开始>
// lengths to test in ascending order
<原文结束>

# <翻译开始>
// 以升序排列的长度值进行测试
# <翻译结束>


<原文开始>
// randomish but deterministic data. No 0 or 255.
<原文结束>

# <翻译开始>
// 随机但确定性的数据。不含0或255。
# <翻译结束>


<原文开始>
// data past the end is different
<原文结束>

# <翻译开始>
// 结尾后的数据不同
# <翻译结束>

