
<原文开始>
// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// 版权所有 (C) 2010 Go 作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可协议约束，
// 该协议可在 LICENSE 文件中找到。
# <翻译结束>


<原文开始>
	// Separately testing "*xyz" (which has no prefix). That is when constructing the
	// pattern to assert on, as in the previous loop, using filepath.Join for an empty
	// prefix filepath.Join(dir, ""), produces the pattern:
	//     ^<DIR>[0-9]+xyz$
	// yet we just want to match
	//     "^<DIR>/[0-9]+xyz"
<原文结束>

# <翻译开始>
	// 单独测试 "*xyz"（它没有前缀）。也就是说，在构建断言模式时，
	// 如在前面的循环中，当使用filepath.Join创建一个空前缀的路径时，
	// 会产生如下模式：
	//     ^<DIR>[0-9]+xyz$
	// 然而我们只想匹配
	//     "^<DIR>/[0-9]+xyz"
# <翻译结束>


<原文开始>
// test that we return a nice error message if the dir argument to TempDir doesn't
// exist (or that it's empty and TempDir doesn't exist)
<原文结束>

# <翻译开始>
// 测试如果TempDir函数的dir参数不存在（或者为空且TempDir不存在）时，是否返回友好的错误消息
# <翻译结束>

