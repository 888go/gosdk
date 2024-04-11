
<原文开始>
// Len returns the number of bytes of the unread portion of the
// string.
<原文结束>

# <翻译开始>
// Len 返回字符串未读部分的字节数。
# <翻译结束>


<原文开始>
// Size returns the original length of the underlying string.
// Size is the number of bytes available for reading via [Reader.ReadAt].
// The returned value is always the same and is not affected by calls
// to any other method.
<原文结束>

# <翻译开始>
// Size 返回底层字符串的原始长度。
// Size 是通过 [Reader.ReadAt] 可读取的字节数。
// 返回的值始终相同，不受对任何其他方法调用的影响。
# <翻译结束>


<原文开始>
// Read implements the [io.Reader] interface.
<原文结束>

# <翻译开始>
// Read 实现了 [io.Reader] 接口。
# <翻译结束>


<原文开始>
// ReadAt implements the [io.ReaderAt] interface.
<原文结束>

# <翻译开始>
// 实现 io.ReaderAt 接口的 ReadAt 方法
# <翻译结束>


<原文开始>
// ReadByte implements the [io.ByteReader] interface.
<原文结束>

# <翻译开始>
// ReadByte 实现了 [io.ByteReader] 接口。
# <翻译结束>


<原文开始>
// UnreadByte implements the [io.ByteScanner] interface.
<原文结束>

# <翻译开始>
// UnreadByte 实现了 [io.ByteScanner] 接口。
# <翻译结束>


<原文开始>
// ReadRune implements the [io.RuneReader] interface.
<原文结束>

# <翻译开始>
// 实现了 io.RuneReader 接口的 ReadRune 方法
# <翻译结束>


<原文开始>
// UnreadRune implements the [io.RuneScanner] interface.
<原文结束>

# <翻译开始>
// UnreadRune 实现了 [io.RuneScanner] 接口。
# <翻译结束>


<原文开始>
// Seek implements the [io.Seeker] interface.
<原文结束>

# <翻译开始>
// 实现了 [io.Seeker] 接口的 Seek 方法
# <翻译结束>


<原文开始>
// WriteTo implements the [io.WriterTo] interface.
<原文结束>

# <翻译开始>
// 实现了 io.WriterTo 接口的 WriteTo 方法
# <翻译结束>


<原文开始>
// Reset resets the [Reader] to be reading from s.
<原文结束>

# <翻译开始>
// Reset 将[Reader]重置为从s读取。
# <翻译结束>


<原文开始>
// NewReader returns a new [Reader] reading from s.
// It is similar to [bytes.NewBufferString] but more efficient and non-writable.
<原文结束>

# <翻译开始>
// NewReader 返回一个从 s 读取的新的 [Reader]。
// 它类似于 [bytes.NewBufferString]，但更高效且不可写。
# <翻译结束>

