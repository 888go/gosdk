
<原文开始>
// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// 版权所有 ? 2022 Go作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。
# <翻译结束>


<原文开始>
// This package contains a simple "batch" allocator for allocating
// coverage counters (slices of uint32 basically), for working with
// coverage data files. Collections of counter arrays tend to all be
// live/dead over the same time period, so a good fit for batch
// allocation.
<原文结束>

# <翻译开始>
// 该包包含一个用于分配覆盖率计数器（基本为uint32切片）的简单“批量”分配器，适用于处理覆盖率数据文件。计数器数组集合往往在同一时间段内同时存活/死亡，因此非常适合批量分配。
# <翻译结束>

