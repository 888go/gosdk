
<原文开始>
// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// 版权所有 2010 Go 作者。保留所有权利。
// 本源代码的使用由可在此文件中找到的BSD风格许可证进行管理。
// md5:a7e52fd757090935
# <翻译结束>


<原文开始>
// For each pattern/text pair, what is the expected output of each function?
// We can derive the textual results from the indexed results, the non-submatch
// results from the submatched results, the single results from the 'all' results,
// and the byte results from the string results. Therefore the table includes
// only the FindAllStringSubmatchIndex result.
<原文结束>

# <翻译开始>
// 对于每个模式/文本对，每个函数的预期输出是什么？
// 我们可以从索引结果推导出文本结果，从子匹配结果中推导出非子匹配结果，从'all'结果中推导出单个结果，以及从字符串结果中推导出字节结果。因此，表格只包含FindAllStringSubmatchIndex的结果。
// md5:adeed5721261f7b3
# <翻译结束>


<原文开始>
// can backslash-escape any punctuation
<原文结束>

# <翻译开始>
// 可以使用反斜杠转义任何标点符号. md5:a59e7bef5123e431
# <翻译结束>


<原文开始>
// long set of matches (longer than startSize)
<原文结束>

# <翻译开始>
// 长序列的匹配（长度超过startSize）. md5:03706d60429e08ee
# <翻译结束>


<原文开始>
// build is a helper to construct a [][]int by extracting n sequences from x.
// This represents n matches with len(x)/n submatches each.
<原文结束>

# <翻译开始>
// build是一个辅助函数，用于通过从x中提取n个序列来构造一个[][]int。这表示有n个匹配，每个匹配包含len(x)/n个子匹配。
// md5:4ad39c42cdaad5dc
# <翻译结束>


<原文开始>
// First the simple cases.
<原文结束>

# <翻译开始>
// 首先处理简单的情况。. md5:85a86fcf805b25a2
# <翻译结束>


<原文开始>
// Tricky because an empty result has two meanings: no match or empty match.
<原文结束>

# <翻译开始>
// 这个问题有点棘手，因为空结果有两种含义：没有匹配到或者匹配到的是空内容。. md5:65756f80e04b83b9
# <翻译结束>


<原文开始>
// Now come the simple All cases.
<原文结束>

# <翻译开始>
// 接下来是简单的All情况。. md5:49427df898552f28
# <翻译结束>


<原文开始>
// Now come the Submatch cases.
<原文结束>

# <翻译开始>
// 现在来看子模式的情况。. md5:b2121920b4698e11
# <翻译结束>


<原文开始>
// Now come the monster AllSubmatch cases.
<原文结束>

# <翻译开始>
// 现在来看看AllSubmatch的怪兽案例。. md5:9481825a4aa3ba59
# <翻译结束>

