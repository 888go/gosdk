
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
// testGetenv gives us a controlled set of variables for testing Expand.
<原文结束>

# <翻译开始>
// testGetenv 为我们提供一组可控的变量，用于测试 Expand。
# <翻译结束>


<原文开始>
// invalid syntax; eat up the characters
<原文结束>

# <翻译开始>
// 无效的语法；消耗掉这些字符
# <翻译结束>


<原文开始>
			// Environment variables on Windows can begin with =
			// https://devblogs.microsoft.com/oldnewthing/20100506-00/?p=14133
<原文结束>

# <翻译开始>
			// 在Windows中，环境变量可以以=开头
			// https:			//devblogs.microsoft.com/oldnewthing/20100506-00/?p=14133
# <翻译结束>


<原文开始>
// On Windows, Environ was observed to report keys with a single leading "=".
// Check that they are properly reported by LookupEnv and can be set by SetEnv.
// See https://golang.org/issue/49886.
<原文结束>

# <翻译开始>
// 在 Windows 系统中，Environ 会报告以单个 "=" 开头的键。我们需要检查 LookupEnv 和 SetEnv 是否能正确处理这种情况。这与 Go 语言的问题 #49886 有关。
# <翻译结束>


<原文开始>
			// We observe in practice keys with a single leading "=" on Windows.
			// TODO(#49886): Should we consume only the first leading "=" as part
			// of the key, or parse through arbitrarily many of them until a non-=,
			// or try each possible key/value boundary until LookupEnv succeeds?
<原文结束>

# <翻译开始>
			// 在实践中，我们在 Windows 系统中观察到键前面只有一个 "="。
			// TODO(#49886): 我们应该只将第一个 "=" 作为键的一部分进行处理，还是解析任意数量的 "="，直到遇到非 "="，或者尝试每个可能的键/值边界，直到 LookupEnv 成功？
# <翻译结束>


<原文开始>
		// Since k=v is already present in the environment,
		// setting it should be a no-op.
<原文结束>

# <翻译开始>
		// 由于 k=v 已经存在于环境中，
		// 对其进行设置应为一个空操作（no-op）。
# <翻译结束>

