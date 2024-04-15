
<原文开始>
// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// 版权所有 ? 2017 The Go Authors。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。
# <翻译结束>


<原文开始>
// #49886: preserve weird Windows keys with leading "=" signs.
<原文结束>

# <翻译开始>
// #49886：保留带有前导"="符号的Windows奇特键。
# <翻译结束>


<原文开始>
			// #52436: preserve invalid key-value entries (for now).
			// (Maybe filter them out or error out on them at some point.)
<原文结束>

# <翻译开始>
// #52436：（暂时）保留无效的键值对条目。
// （未来可能会在某个时刻过滤掉它们或因它们而报错。）
# <翻译结束>


<原文开始>
// Filter out entries containing NULs.
<原文结束>

# <翻译开始>
// 过滤掉包含 NUL 的条目。
# <翻译结束>


<原文开始>
// Plan 9 needs to preserve environment variables with NUL (#56544).
<原文结束>

# <翻译开始>
// Plan 9需要保留带有NUL（空字符）的环境变量（#56544）。
# <翻译结束>

