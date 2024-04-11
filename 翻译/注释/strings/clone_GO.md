
<原文开始>
// Clone returns a fresh copy of s.
// It guarantees to make a copy of s into a new allocation,
// which can be important when retaining only a small substring
// of a much larger string. Using Clone can help such programs
// use less memory. Of course, since using Clone makes a copy,
// overuse of Clone can make programs use more memory.
// Clone should typically be used only rarely, and only when
// profiling indicates that it is needed.
// For strings of length zero the string "" will be returned
// and no allocation is made.
<原文结束>

# <翻译开始>
// Clone 返回一个 s 的全新副本。
// 它确保将 s 复制到新的内存分配中，
// 这对于仅保留一个较大字符串中的小部分子串而言非常重要。
// 使用 Clone 可帮助此类程序减少内存使用。当然，由于 Clone 会进行复制，
// 对 Clone 的过度使用可能会使程序占用更多内存。
// Clone 通常应仅在极少数情况下使用，并且仅当性能分析表明需要时才使用。
// 对于长度为零的字符串，将返回空字符串 ""，且不会进行内存分配。
# <翻译结束>

