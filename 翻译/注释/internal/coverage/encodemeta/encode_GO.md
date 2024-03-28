
<原文开始>
// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// 版权所有 ? 2021 Go作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。
# <翻译结束>


<原文开始>
// This package contains APIs and helpers for encoding the meta-data
// "blob" for a single Go package, created when coverage
// instrumentation is turned on.
<原文结束>

# <翻译开始>
// 此包包含用于编码单个Go包的元数据“blob”的API和辅助函数，这些元数据在启用覆盖率检测时生成。
# <翻译结束>


<原文开始>
// AddFunc registers a new function with the meta data builder.
<原文结束>

# <翻译开始>
// AddFunc 使用元数据构建器注册一个新的函数。
# <翻译结束>


<原文开始>
// Emit writes the meta-data accumulated so far in this builder to 'w'.
// Returns a hash of the meta-data payload and an error.
<原文结束>

# <翻译开始>
// Emit 将此构建器中迄今为止积累的元数据写入到 'w'。
// 返回元数据负载的哈希值和错误信息。
# <翻译结束>


<原文开始>
	// Emit header.  Length will initially be zero, we'll
	// back-patch it later.
<原文结束>

# <翻译开始>
// 发送（写入）头部。长度最初为零，我们将在稍后进行回溯修补。
# <翻译结束>


<原文开始>
// hash and length initially zero, will be back-patched
<原文结束>

# <翻译开始>
// 散列值和长度初始为零，将在后续进行回填
# <翻译结束>


<原文开始>
// Write function offsets section
<原文结束>

# <翻译开始>
// 写入函数偏移量部分
# <翻译结束>


<原文开始>
// Check for any errors up to this point.
<原文结束>

# <翻译开始>
// 检查到目前为止出现的任何错误。
# <翻译结束>


<原文开始>
// HashFuncDesc computes an md5 sum of a coverage.FuncDesc and returns
// a digest for it.
<原文结束>

# <翻译开始>
// HashFuncDesc 计算 coverage.FuncDesc 的 MD5 值，并返回其摘要。
# <翻译结束>


<原文开始>
// hashFuncDesc incorporates a given function 'f' into the hash 'h'.
<原文结束>

# <翻译开始>
// hashFuncDesc 将给定的函数 'f' 结合到散列 'h' 中。
# <翻译结束>

