
<原文开始>
// NewReplacer returns a new [Replacer] from a list of old, new string
// pairs. Replacements are performed in the order they appear in the
// target string, without overlapping matches. The old string
// comparisons are done in argument order.
//
// NewReplacer panics if given an odd number of arguments.
<原文结束>

# <翻译开始>
// NewReplacer 从一组旧字符串、新字符串对中返回一个新的[Replacer]。替换操作按照它们在目标字符串中出现的顺序进行，且不涉及重叠匹配。旧字符串的比较按参数顺序进行。
// 
// 如果提供的参数数目为奇数，NewReplacer 将引发 panic。
# <翻译结束>


<原文开始>
// Replace returns a copy of s with all replacements performed.
<原文结束>

# <翻译开始>
// Replace返回对s进行所有替换操作后的副本。
# <翻译结束>


<原文开始>
// WriteString writes s to w with all replacements performed.
<原文结束>

# <翻译开始>
// WriteString 将s写入w，同时执行所有替换操作。
# <翻译结束>


<原文开始>
// Write writes to the buffer to satisfy io.Writer.
<原文结束>

# <翻译开始>
// Write 向缓冲区写入数据以实现 io.Writer 接口
# <翻译结束>


<原文开始>
// WriteString writes to the buffer without string->[]byte->string allocations.
<原文结束>

# <翻译开始>
// WriteString 向缓冲区写入字符串，无需进行 string->[]byte->string 的内存分配。
# <翻译结束>

