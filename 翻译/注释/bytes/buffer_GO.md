
<原文开始>
// A Buffer is a variable-sized buffer of bytes with [Buffer.Read] and [Buffer.Write] methods.
// The zero value for Buffer is an empty buffer ready to use.
<原文结束>

# <翻译开始>
// Buffer是一个具有[Buffer.Read]和[Buffer.Write]方法的可变大小字节缓冲区。
// Buffer的零值是一个空缓冲区，已经准备好使用。
// md5:6eeb6c62d2c690fe
# <翻译结束>


<原文开始>
// Any other read operation.
<原文结束>

# <翻译开始>
// 任何其他读取操作。. md5:640135789ce1c247
# <翻译结束>


<原文开始>
// MinRead is the minimum slice size passed to a Read call by
// [Buffer.ReadFrom]. As long as the [Buffer] has at least MinRead bytes beyond
// what is required to hold the contents of r, ReadFrom will not grow the
// underlying buffer.
<原文结束>

# <翻译开始>
// MinRead 是由[Buffer.ReadFrom]传递给Read调用的最小切片大小。只要[Buffer]中至少有MinRead字节超出存储r内容所需的字节数，ReadFrom将不会扩大底层缓冲区的大小。
// md5:9a6fbe111a4fae8b
# <翻译结束>


<原文开始>
// ErrTooLarge is passed to panic if memory cannot be allocated to store data in a buffer.
<原文结束>

# <翻译开始>
// ErrTooLarge 如果无法分配内存来存储缓冲区中的数据，该错误会被传递给panic。. md5:ac0db4ded02d1d8f
# <翻译结束>


<原文开始>
// Bytes returns a slice of length b.Len() holding the unread portion of the buffer.
// The slice is valid for use only until the next buffer modification (that is,
// only until the next call to a method like [Buffer.Read], [Buffer.Write], [Buffer.Reset], or [Buffer.Truncate]).
// The slice aliases the buffer content at least until the next buffer modification,
// so immediate changes to the slice will affect the result of future reads.
<原文结束>

# <翻译开始>
// Bytes 返回一个长度为 b.Len() 的切片，其中包含缓冲区中未读的部分。
// 这个切片只在下一次修改缓冲区（即调用诸如 [Buffer.Read]、[Buffer.Write]、[Buffer.Reset] 或 [Buffer.Truncate] 等方法之前）有效。
// 切片至少会引用缓冲区的内容，直到下一次修改，因此立即对切片进行更改将影响未来读取的结果。
// md5:1dfdc5bab519f911
# <翻译结束>


<原文开始>
// AvailableBuffer returns an empty buffer with b.Available() capacity.
// This buffer is intended to be appended to and
// passed to an immediately succeeding [Buffer.Write] call.
// The buffer is only valid until the next write operation on b.
<原文结束>

# <翻译开始>
// AvailableBuffer 返回一个容量为 b.Available() 的空缓冲区。
// 此缓冲区旨在被追加到并传递给下一个立即接续的 [Buffer.Write] 调用。
// 直到对 b 进行下一次写操作，这个缓冲区才有效。
// md5:ffac6549b637e07e
# <翻译结束>


<原文开始>
// String returns the contents of the unread portion of the buffer
// as a string. If the [Buffer] is a nil pointer, it returns "<nil>".
//
// To build strings more efficiently, see the strings.Builder type.
<原文结束>

# <翻译开始>
// String 方法返回缓冲区未读部分内容作为字符串。
// 如果 [Buffer] 是一个空指针，它将返回 "<nil>"。
//
// 要更高效地构建字符串，请参见 strings.Builder 类型。
// md5:1c3184cdbc946e8a
# <翻译结束>


<原文开始>
// Len returns the number of bytes of the unread portion of the buffer;
// b.Len() == len(b.Bytes()).
<原文结束>

# <翻译开始>
// Len 返回缓冲区中未读部分的字节数量；
// b.Len() 等于 len(b.Bytes())。
// md5:435301354c6501b1
# <翻译结束>


<原文开始>
// Cap returns the capacity of the buffer's underlying byte slice, that is, the
// total space allocated for the buffer's data.
<原文结束>

# <翻译开始>
// Cap返回缓冲区底层字节切片的容量，即分配给缓冲区数据的总空间。
// md5:7bf4166f47a43520
# <翻译结束>


<原文开始>
// Available returns how many bytes are unused in the buffer.
<原文结束>

# <翻译开始>
// Available 返回缓冲区中未使用的字节数。. md5:fbfc156036bbe683
# <翻译结束>


<原文开始>
// Truncate discards all but the first n unread bytes from the buffer
// but continues to use the same allocated storage.
// It panics if n is negative or greater than the length of the buffer.
<原文结束>

# <翻译开始>
// Truncate 方法丢弃缓冲区中除前n个未读取字节以外的所有字节，
// 但仍然使用相同的已分配存储空间。
// 如果n是负数或大于缓冲区的长度，它会引发恐慌。
// md5:98136bb1970ac81b
# <翻译结束>


<原文开始>
// Reset resets the buffer to be empty,
// but it retains the underlying storage for use by future writes.
// Reset is the same as [Buffer.Truncate](0).
<原文结束>

# <翻译开始>
// Reset 将缓冲区重置为空，
// 但保留底层存储，以便将来写入使用。
// Reset 功能等同于 [Buffer.Truncate](0)。
// md5:5e5cecee92025b5d
# <翻译结束>


<原文开始>
// Grow grows the buffer's capacity, if necessary, to guarantee space for
// another n bytes. After Grow(n), at least n bytes can be written to the
// buffer without another allocation.
// If n is negative, Grow will panic.
// If the buffer can't grow it will panic with [ErrTooLarge].
<原文结束>

# <翻译开始>
// Grow 增大缓冲区的容量，如果需要的话，以确保能容纳另外 n 个字节。在调用 Grow(n) 后，至少可以写入 n 个字节到缓冲区而无需再次分配空间。
// 如果 n 是负数，Grow 将会引发 panic。
// 如果缓冲区无法增长，它将使用 [ErrTooLarge] 错误引发 panic。
// md5:48df9b8648f9a546
# <翻译结束>


<原文开始>
// Write appends the contents of p to the buffer, growing the buffer as
// needed. The return value n is the length of p; err is always nil. If the
// buffer becomes too large, Write will panic with [ErrTooLarge].
<原文结束>

# <翻译开始>
// Write 将 p 的内容追加到缓冲区，如果需要会扩大缓冲区。返回值 n 是 p 的长度；err 总是为 nil。如果缓冲区变得过大，Write 会引发一个 [ErrTooLarge] 的 panic。
// md5:1dde3924b7871834
# <翻译结束>


<原文开始>
// WriteString appends the contents of s to the buffer, growing the buffer as
// needed. The return value n is the length of s; err is always nil. If the
// buffer becomes too large, WriteString will panic with [ErrTooLarge].
<原文结束>

# <翻译开始>
// WriteString 将 s 的内容追加到缓冲区中，根据需要扩大缓冲区的大小。返回值 n 是 s 的长度；err 始终为 nil。如果缓冲区变得过大，WriteString 将会因 [ErrTooLarge] 而发生恐慌。
// md5:90c00c4ec56819db
# <翻译结束>


<原文开始>
// ReadFrom reads data from r until EOF and appends it to the buffer, growing
// the buffer as needed. The return value n is the number of bytes read. Any
// error except io.EOF encountered during the read is also returned. If the
// buffer becomes too large, ReadFrom will panic with [ErrTooLarge].
<原文结束>

# <翻译开始>
// ReadFrom从r读取数据直到EOF，并将数据追加到缓冲区，根据需要增长缓冲区。返回值n是读取的字节数。在读取过程中遇到的任何错误（除了io.EOF）也会被返回。如果缓冲区变得太大，ReadFrom将会因[ErrTooLarge]引发恐慌。
// md5:eaa622ebea86d082
# <翻译结束>


<原文开始>
// WriteTo writes data to w until the buffer is drained or an error occurs.
// The return value n is the number of bytes written; it always fits into an
// int, but it is int64 to match the io.WriterTo interface. Any error
// encountered during the write is also returned.
<原文结束>

# <翻译开始>
// WriteTo 将数据写入到 w，直到缓冲区被耗尽或出现错误。
// 返回值 n 是已写入的字节数；它始终可以容纳到一个 int 类型，但使用 int64 类型以匹配 io.WriterTo 接口。在写入过程中遇到的任何错误也会一并返回。
// md5:c7017c8e4e76bc05
# <翻译结束>


<原文开始>
// WriteByte appends the byte c to the buffer, growing the buffer as needed.
// The returned error is always nil, but is included to match [bufio.Writer]'s
// WriteByte. If the buffer becomes too large, WriteByte will panic with
// [ErrTooLarge].
<原文结束>

# <翻译开始>
// WriteByte 将字节 c 追加到缓冲区，如果需要会扩大缓冲区。返回的错误始终为 nil，但包含它以匹配 [bufio.Writer] 的 WriteByte 方法。如果缓冲区变得太大，WriteByte 会用 [ErrTooLarge] 引发 panic。
// md5:3807962fe1acfbac
# <翻译结束>


<原文开始>
// WriteRune appends the UTF-8 encoding of Unicode code point r to the
// buffer, returning its length and an error, which is always nil but is
// included to match [bufio.Writer]'s WriteRune. The buffer is grown as needed;
// if it becomes too large, WriteRune will panic with [ErrTooLarge].
<原文结束>

# <翻译开始>
// WriteRune 将Unicode代码点r的UTF-8编码附加到缓冲区中，
// 并返回其长度和一个错误，该错误总是为nil，但为了与[bufio.Writer]的WriteRune方法匹配而包含在内。
// 缓冲区会根据需要自动扩展；如果缓冲区变得过大，WriteRune将会因[ErrTooLarge]错误而引发恐慌。
// md5:fcd3a4d673494541
# <翻译结束>


<原文开始>
// Read reads the next len(p) bytes from the buffer or until the buffer
// is drained. The return value n is the number of bytes read. If the
// buffer has no data to return, err is io.EOF (unless len(p) is zero);
// otherwise it is nil.
<原文结束>

# <翻译开始>
// Read 从缓冲区中读取接下来的 len(p) 个字节，或者直到缓冲区被读空。返回值 n 是读取的字节数。如果缓冲区没有数据可读，err 将是 io.EOF（除非 len(p) 为零）；否则 err 为 nil。
// md5:6dec72f7698c2809
# <翻译结束>


<原文开始>
// Next returns a slice containing the next n bytes from the buffer,
// advancing the buffer as if the bytes had been returned by [Buffer.Read].
// If there are fewer than n bytes in the buffer, Next returns the entire buffer.
// The slice is only valid until the next call to a read or write method.
<原文结束>

# <翻译开始>
// Next 函数从缓冲区中返回包含接下来 n 个字节的切片，就像 [Buffer.Read] 返回的一样。如果缓冲区中的字节数少于 n，Next 将返回整个缓冲区。这个切片在下一次调用读写方法之前都是有效的。
// md5:4b6d65f5719e684a
# <翻译结束>


<原文开始>
// ReadByte reads and returns the next byte from the buffer.
// If no byte is available, it returns error io.EOF.
<原文结束>

# <翻译开始>
// ReadByte 从缓冲区读取并返回下一个字节。如果没有可用的字节，它将返回错误 io.EOF。
// md5:9a65c587bf66dcf3
# <翻译结束>


<原文开始>
// ReadRune reads and returns the next UTF-8-encoded
// Unicode code point from the buffer.
// If no bytes are available, the error returned is io.EOF.
// If the bytes are an erroneous UTF-8 encoding, it
// consumes one byte and returns U+FFFD, 1.
<原文结束>

# <翻译开始>
// ReadRune 读取并返回缓冲区中的下一个UTF-8编码的Unicode码点。
// 如果没有可用的字节，返回的错误是io.EOF。
// 如果字节是错误的UTF-8编码，它将消耗一个字节并返回U+FFFD，1。
// md5:e6188dfcfac12306
# <翻译结束>


<原文开始>
// UnreadRune unreads the last rune returned by [Buffer.ReadRune].
// If the most recent read or write operation on the buffer was
// not a successful [Buffer.ReadRune], UnreadRune returns an error.  (In this regard
// it is stricter than [Buffer.UnreadByte], which will unread the last byte
// from any read operation.)
<原文结束>

# <翻译开始>
// UnreadRune 回退由 Buffer.ReadRune 返回的最后一个字符。
// 如果对缓冲区的最近一次读写操作不是成功的 ReadRune，UnreadRune 将返回一个错误。
// 在这方面，它比 Buffer.UnreadByte 更严格，因为 Buffer.UnreadByte 可以回退任何读取操作的最后一个字节。
// md5:239e0a258391add0
# <翻译结束>


<原文开始>
// UnreadByte unreads the last byte returned by the most recent successful
// read operation that read at least one byte. If a write has happened since
// the last read, if the last read returned an error, or if the read read zero
// bytes, UnreadByte returns an error.
<原文结束>

# <翻译开始>
// UnreadByte 将最近一次成功读取（至少读取了一个字节）后返回的最后一个字节 unread。如果在上次读取之后有写操作发生，或者上次读取返回了错误，或者读取了0个字节，UnreadByte 将返回一个错误。
// md5:ebefc8dc07168356
# <翻译结束>


<原文开始>
// ReadBytes reads until the first occurrence of delim in the input,
// returning a slice containing the data up to and including the delimiter.
// If ReadBytes encounters an error before finding a delimiter,
// it returns the data read before the error and the error itself (often io.EOF).
// ReadBytes returns err != nil if and only if the returned data does not end in
// delim.
<原文结束>

# <翻译开始>
// ReadBytes 读取到输入中的第一个分隔符，返回包含到和包括分隔符的数据切片。如果在找到分隔符之前遇到错误，它将返回错误发生前已读取的数据以及错误（通常是 io.EOF）。只有当返回的数据不以分隔符结束时，ReadBytes 才会返回 err != nil。
// md5:4d18e9ed50415556
# <翻译结束>


<原文开始>
// ReadString reads until the first occurrence of delim in the input,
// returning a string containing the data up to and including the delimiter.
// If ReadString encounters an error before finding a delimiter,
// it returns the data read before the error and the error itself (often io.EOF).
// ReadString returns err != nil if and only if the returned data does not end
// in delim.
<原文结束>

# <翻译开始>
// ReadString 读取输入直到遇到第一个 delim 字符为止，
// 并返回一个字符串，其中包含直到分隔符（包括分隔符本身）的所有数据。
// 如果 ReadString 在找到分隔符之前遇到错误，
// 它会返回遇到错误之前读取的数据以及该错误本身（常为 io.EOF）。
// ReadString 在且仅在返回的数据没有以 delim 结尾时，返回 err != nil。
// md5:2dbe6334091e50c0
# <翻译结束>


<原文开始>
// NewBuffer creates and initializes a new [Buffer] using buf as its
// initial contents. The new [Buffer] takes ownership of buf, and the
// caller should not use buf after this call. NewBuffer is intended to
// prepare a [Buffer] to read existing data. It can also be used to set
// the initial size of the internal buffer for writing. To do that,
// buf should have the desired capacity but a length of zero.
//
// In most cases, new([Buffer]) (or just declaring a [Buffer] variable) is
// sufficient to initialize a [Buffer].
<原文结束>

# <翻译开始>
// NewBuffer 使用 buf 作为初始内容创建并初始化一个新的 [Buffer]。新的 [Buffer] 将接管 buf，调用者在调用后不应再使用 buf。NewBuffer 的目的是准备一个 [Buffer] 来读取已有数据。它也可以用来设置内部缓冲区的初始写入大小。为此，buf 应具有所需的容量，但长度应为零。
//
// 在大多数情况下，new([Buffer])（或者直接声明一个 [Buffer] 变量）就足以初始化一个 [Buffer]。
// md5:160089e06b42c9d0
# <翻译结束>


<原文开始>
// NewBufferString creates and initializes a new [Buffer] using string s as its
// initial contents. It is intended to prepare a buffer to read an existing
// string.
//
// In most cases, new([Buffer]) (or just declaring a [Buffer] variable) is
// sufficient to initialize a [Buffer].
<原文结束>

# <翻译开始>
// NewBufferString 通过字符串 s 创建并初始化一个新的 [Buffer]。它的目的是准备一个缓冲区来读取已存在的字符串。
// 
// 在大多数情况下，使用 new([Buffer])（或者只是声明一个 [Buffer] 变量）就足以初始化一个 [Buffer]。
// md5:52db730cdb1471dd
# <翻译结束>

