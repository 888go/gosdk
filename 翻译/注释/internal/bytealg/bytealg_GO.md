
<原文开始>
// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// 版权所有 ? 2018 The Go Authors。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。
# <翻译结束>


<原文开始>
// Offsets into internal/cpu records for use in assembly.
<原文结束>

# <翻译开始>
// Offsets偏移量进入内部/cpu记录，供汇编中使用。
# <翻译结束>


<原文开始>
// MaxLen is the maximum length of the string to be searched for (argument b) in Index.
// If MaxLen is not 0, make sure MaxLen >= 4.
<原文结束>

# <翻译开始>
// MaxLen 是在 Index 中被搜索的字符串（参数 b）的最大长度。
// 如果 MaxLen 不为 0，则确保 MaxLen >= 4。
# <翻译结束>


<原文开始>
// FIXME: the logic of HashStrBytes, HashStrRevBytes, IndexRabinKarpBytes and HashStr, HashStrRev,
// IndexRabinKarp are exactly the same, except that the types are different. Can we eliminate
// three of them without causing allocation?
<原文结束>

# <翻译开始>
// FIXME：HashStrBytes、HashStrRevBytes、IndexRabinKarpBytes 以及 HashStr、HashStrRev 的逻辑完全相同，只是类型不同。我们能否在不导致内存分配的情况下消除其中三个函数？
# <翻译结束>


<原文开始>
// PrimeRK is the prime base used in Rabin-Karp algorithm.
<原文结束>

# <翻译开始>
// PrimeRK 是在 Rabin-Karp 算法中使用的质数基数。
# <翻译结束>


<原文开始>
// HashStrBytes returns the hash and the appropriate multiplicative
// factor for use in Rabin-Karp algorithm.
<原文结束>

# <翻译开始>
// HashStrBytes 返回散列值以及在 Rabin-Karp 算法中使用的适当乘法因子。
# <翻译结束>


<原文开始>
// HashStr returns the hash and the appropriate multiplicative
// factor for use in Rabin-Karp algorithm.
<原文结束>

# <翻译开始>
// HashStr 函数返回一个哈希值及适用于 Rabin-Karp 算法的相应乘数因子。
# <翻译结束>


<原文开始>
// HashStrRevBytes returns the hash of the reverse of sep and the
// appropriate multiplicative factor for use in Rabin-Karp algorithm.
<原文结束>

# <翻译开始>
// HashStrRevBytes 返回sep反向字符串的哈希值以及在Rabin-Karp算法中使用的适当乘法因子。
# <翻译结束>


<原文开始>
// HashStrRev returns the hash of the reverse of sep and the
// appropriate multiplicative factor for use in Rabin-Karp algorithm.
<原文结束>

# <翻译开始>
// HashStrRev 计算 sep 反转字符串的哈希值以及在 Rabin-Karp 算法中使用的相应乘法因子。
# <翻译结束>


<原文开始>
// IndexRabinKarpBytes uses the Rabin-Karp search algorithm to return the index of the
// first occurrence of substr in s, or -1 if not present.
<原文结束>

# <翻译开始>
// IndexRabinKarpBytes 使用 Rabin-Karp 搜索算法在字符串 s 中查找子串 substr 的首次出现位置，并返回该索引，若未找到则返回 -1。
# <翻译结束>


<原文开始>
// IndexRabinKarp uses the Rabin-Karp search algorithm to return the index of the
// first occurrence of substr in s, or -1 if not present.
<原文结束>

# <翻译开始>
// IndexRabinKarp 使用 Rabin-Karp 搜索算法在字符串 s 中查找子串 substr 的首次出现位置，并返回该索引。如果子串不存在，则返回 -1。
# <翻译结束>

