
<原文开始>
// String returns the accumulated string.
<原文结束>

# <翻译开始>
// String 返回累计生成的字符串。
# <翻译结束>


<原文开始>
// Len returns the number of accumulated bytes; b.Len() == len(b.String()).
<原文结束>

# <翻译开始>
// Len 返回已累积的字节数；b.Len() 等于 len(b.String())。
# <翻译结束>


<原文开始>
// Cap returns the capacity of the builder's underlying byte slice. It is the
// total space allocated for the string being built and includes any bytes
// already written.
<原文结束>

# <翻译开始>
// Cap 返回构建器底层字节切片的容量。它是指为正在构建的字符串所分配的总空间，包括已写入的所有字节。
# <翻译结束>


<原文开始>
// Reset resets the [Builder] to be empty.
<原文结束>

# <翻译开始>
// Reset 重置 [Builder] 为空。
# <翻译结束>


<原文开始>
// Grow grows b's capacity, if necessary, to guarantee space for
// another n bytes. After Grow(n), at least n bytes can be written to b
// without another allocation. If n is negative, Grow panics.
<原文结束>

# <翻译开始>
// Grow 会根据需要增加 b 的容量，以确保为接下来的 n 个字节提供足够的空间。在调用 Grow(n) 后，至少可以向 b 中写入 n 个字节而无需再进行一次分配。如果 n 为负值，Grow 将引发恐慌（panic）。
# <翻译结束>


<原文开始>
// Write appends the contents of p to b's buffer.
// Write always returns len(p), nil.
<原文结束>

# <翻译开始>
// Write 将 p 的内容追加至 b 的缓冲区。
// Write 总是返回 len(p), nil。
# <翻译结束>


<原文开始>
// WriteByte appends the byte c to b's buffer.
// The returned error is always nil.
<原文结束>

# <翻译开始>
// WriteByte 将字节 c 追加至 b 的缓冲区。
// 返回的错误始终为 nil。
# <翻译结束>


<原文开始>
// WriteRune appends the UTF-8 encoding of Unicode code point r to b's buffer.
// It returns the length of r and a nil error.
<原文结束>

# <翻译开始>
// WriteRune 将Unicode码点r的UTF-8编码追加到b的缓冲区中。
// 它返回r的长度（字节数）并返回一个nil错误。
# <翻译结束>


<原文开始>
// WriteString appends the contents of s to b's buffer.
// It returns the length of s and a nil error.
<原文结束>

# <翻译开始>
// WriteString 将 s 的内容追加到 b 的缓冲区中。
// 它返回 s 的长度以及一个空（nil）错误。
# <翻译结束>

