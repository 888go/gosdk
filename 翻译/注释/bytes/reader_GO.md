
<原文开始>
// A Reader implements the io.Reader, io.ReaderAt, io.WriterTo, io.Seeker,
// io.ByteScanner, and io.RuneScanner interfaces by reading from
// a byte slice.
// Unlike a [Buffer], a Reader is read-only and supports seeking.
// The zero value for Reader operates like a Reader of an empty slice.
<原文结束>

# <翻译开始>
// Reader实现了io.Reader、io.ReaderAt、io.WriterTo、io.Seeker、io.ByteScanner和io.RuneScanner接口，通过从字节切片中读取数据。
// 与[Buffer]不同，Reader是只读的，并且支持定位（seeking）操作。
// Reader的零值行为就像一个空切片的Reader。
// md5:c7615bfef4c8b2fe
# <翻译结束>


<原文开始>
// Len returns the number of bytes of the unread portion of the
// slice.
<原文结束>

# <翻译开始>
// Len 返回切片未读部分的字节长度。
// md5:60d512b53434757a
# <翻译结束>


<原文开始>
// Size returns the original length of the underlying byte slice.
// Size is the number of bytes available for reading via [Reader.ReadAt].
// The result is unaffected by any method calls except [Reader.Reset].
<原文结束>

# <翻译开始>
// Size 返回原始字节切片的长度。
// Size 是通过 [Reader.ReadAt] 可读取的字节数量。
// 结果不受除 [Reader.Reset] 之外的任何方法调用影响。
// md5:7057c00761109102
# <翻译结束>


<原文开始>
// Read implements the [io.Reader] interface.
<原文结束>

# <翻译开始>
// Read implements the [io.Reader] 接口。. md5:6b6e15958e2263c0
# <翻译结束>


<原文开始>
// ReadAt implements the [io.ReaderAt] interface.
<原文结束>

# <翻译开始>
// ReadAt 实现了 [io.ReaderAt] 接口。. md5:aff23aac8adb605e
# <翻译结束>


<原文开始>
// ReadByte implements the [io.ByteReader] interface.
<原文结束>

# <翻译开始>
// ReadByte 实现了 [io.ByteReader] 接口。. md5:9580ddd72af46f48
# <翻译结束>


<原文开始>
// UnreadByte complements [Reader.ReadByte] in implementing the [io.ByteScanner] interface.
<原文结束>

# <翻译开始>
// UnreadByte 在实现 [io.ByteScanner] 接口时补充了 [Reader.ReadByte] 的功能。. md5:5aa8f5cc398bf852
# <翻译结束>


<原文开始>
// ReadRune implements the [io.RuneReader] interface.
<原文结束>

# <翻译开始>
// ReadRune 实现了 [io.RuneReader] 接口。. md5:97f6bcc2e44946c6
# <翻译结束>


<原文开始>
// UnreadRune complements [Reader.ReadRune] in implementing the [io.RuneScanner] interface.
<原文结束>

# <翻译开始>
// UnreadRune 在实现 [io.RuneScanner] 接口时补充了 [Reader.ReadRune]。. md5:d0948bd1ce43eeef
# <翻译结束>


<原文开始>
// Seek implements the [io.Seeker] interface.
<原文结束>

# <翻译开始>
// Seek 实现了 [io.Seeker] 接口。. md5:407fc18ddaadacba
# <翻译结束>


<原文开始>
// WriteTo implements the [io.WriterTo] interface.
<原文结束>

# <翻译开始>
// WriteTo 实现了 [io.WriterTo] 接口。. md5:93bec25fa975a33f
# <翻译结束>


<原文开始>
// Reset resets the [Reader.Reader] to be reading from b.
<原文结束>

# <翻译开始>
// Reset 将[Reader.Reader]重置为从b读取。. md5:3c05fdb3a579bd62
# <翻译结束>


<原文开始>
// NewReader returns a new [Reader.Reader] reading from b.
<原文结束>

# <翻译开始>
// NewReader 返回一个新的从 b 读取的 [Reader.Reader]。. md5:807cf3d6e5b558e1
# <翻译结束>

