package bytes

import (
	"bytes"
	"io"
)

// Buffer是一个具有[Buffer.Read]和[Buffer.Write]方法的可变大小字节缓冲区。
// Buffer的零值是一个空缓冲区，已经准备好使用。
// md5:6eeb6c62d2c690fe
type Buffer struct { //hm:缓冲区结构  cz:type Buffer
	F bytes.Buffer
} //md5:6199995c60d37c83

//const (
//	opRead      readOp = -1 // 任何其他读取操作。. md5:640135789ce1c247
//	opInvalid   readOp = 0  // Non-read operation.
//	opReadRune1 readOp = 1  // Read rune of size 1.
//	opReadRune2 readOp = 2  // Read rune of size 2.
//	opReadRune3 readOp = 3  // Read rune of size 3.
//	opReadRune4 readOp = 4  // Read rune of size 4.
//) // md5:786a865ce796138e

// MinRead 是由[Buffer.ReadFrom]传递给Read调用的最小切片大小。只要[Buffer]中至少有MinRead字节超出存储r内容所需的字节数，ReadFrom将不会扩大底层缓冲区的大小。
// md5:9a6fbe111a4fae8b
const MinRead = 512 //md5:e6440ccdc461dcd6//hm:常量_缓冲区_最小读取大小  cz:const MinRead

// ErrTooLarge 如果无法分配内存来存储缓冲区中的数据，该错误会被传递给panic。. md5:ac0db4ded02d1d8f
var ErrTooLarge = bytes.ErrTooLarge //md5:5fb25188c1ad1858//qm:错误_缓冲区_太大  cz:ErrTooLarge = bytes.ErrTooLarge

// Bytes 返回一个长度为 b.Len() 的切片，其中包含缓冲区中未读的部分。
// 这个切片只在下一次修改缓冲区（即调用诸如 [Buffer.Read]、[Buffer.Write]、[Buffer.Reset] 或 [Buffer.Truncate] 等方法之前）有效。
// 切片至少会引用缓冲区的内容，直到下一次修改，因此立即对切片进行更改将影响未来读取的结果。
// md5:1dfdc5bab519f911
// ff:取缓冲区底层切片
// b:
func (b *Buffer) Bytes() []byte { //md5:3f81aeed184ada90
	return b.F.Bytes()
}

// String 方法返回缓冲区未读部分内容作为字符串。
// 如果 [Buffer] 是一个空指针，它将返回 "<nil>"。
//
// 要更高效地构建字符串，请参见 strings.Builder 类型。
// md5:1c3184cdbc946e8a
// ff:
// b:
func (b *Buffer) String() string { //md5:2a07d671bd398c91
	return b.F.String()
}

// Len 返回缓冲区中未读部分的字节数量；
// b.Len() 等于 len(b.Bytes())。
// md5:435301354c6501b1
// ff:取未读长度
// b:
func (b *Buffer) Len() int { //md5:3808bf9c6d07bc82
	return b.F.Len()
}

// Cap返回缓冲区底层字节切片的容量，即分配给缓冲区数据的总空间。
// md5:7bf4166f47a43520
// ff:取容量
// b:
func (b *Buffer) Cap() int { //md5:9ccaf75b2ffee8c7
	return b.F.Cap()
}

// Truncate 方法丢弃缓冲区中除前n个未读取字节以外的所有字节，
// 但仍然使用相同的已分配存储空间。
// 如果n是负数或大于缓冲区的长度，它会引发恐慌。
// md5:98136bb1970ac81b
// ff:截断至长度
// b:
// n:长度
func (b *Buffer) Truncate(n int) { //md5:fed2f1e2c899cb47
	b.F.Truncate(n)
}

// Reset 将缓冲区重置为空，
// 但保留底层存储，以便将来写入使用。
// Reset 功能等同于 [Buffer.Truncate](0)。
// md5:5e5cecee92025b5d
// ff:清空
// b:
func (b *Buffer) Reset() { //md5:a00af8e06e8f397f
	b.F.Reset()
}

// Grow 增大缓冲区的容量，如果需要的话，以确保能容纳另外 n 个字节。在调用 Grow(n) 后，至少可以写入 n 个字节到缓冲区而无需再次分配空间。
// 如果 n 是负数，Grow 将会引发 panic。
// 如果缓冲区无法增长，它将使用 [ErrTooLarge] 错误引发 panic。
// md5:48df9b8648f9a546
// ff:扩展容量
// b:
// n:扩展长度
func (b *Buffer) Grow(n int) { //md5:e9e4897870e0d141
	b.F.Grow(n)
}

// Write 将 p 的内容追加到缓冲区，如果需要会扩大缓冲区。返回值 n 是 p 的长度；err 总是为 nil。如果缓冲区变得过大，Write 会引发一个 [ErrTooLarge] 的 panic。
// md5:1dde3924b7871834
// ff:
// b:
// p:写入字节集
// n:写入长度
// err:错误
func (b *Buffer) Write(p []byte) (n int, err error) { //md5:6f1eceb67804b19f
	return b.F.Write(p)
}

// WriteString 将 s 的内容追加到缓冲区中，根据需要扩大缓冲区的大小。返回值 n 是 s 的长度；err 始终为 nil。如果缓冲区变得过大，WriteString 将会因 [ErrTooLarge] 而发生恐慌。
// md5:90c00c4ec56819db
// ff:写入文本
// b:
// s:文本
// n:写入长度
// err:错误
func (b *Buffer) WriteString(s string) (n int, err error) { //md5:74161787fa18801b
	return b.F.WriteString(s)
}

// ReadFrom从r读取数据直到EOF，并将数据追加到缓冲区，根据需要增长缓冲区。返回值n是读取的字节数。在读取过程中遇到的任何错误（除了io.EOF）也会被返回。如果缓冲区变得太大，ReadFrom将会因[ErrTooLarge]引发恐慌。
// md5:eaa622ebea86d082
// ff:写入并从IO读取器
// b:
// r:IO读取器
// n:写入长度
// err:错误
func (b *Buffer) ReadFrom(r io.Reader) (n int64, err error) { //md5:8c6eeb83f74f22c8
	return b.F.ReadFrom(r)
}

// WriteTo 将数据写入到 w，直到缓冲区被耗尽或出现错误。
// 返回值 n 是已写入的字节数；它始终可以容纳到一个 int 类型，但使用 int64 类型以匹配 io.WriterTo 接口。在写入过程中遇到的任何错误也会一并返回。
// md5:c7017c8e4e76bc05
// ff:读数据并写到IO写入器
// b:
// w:IO写入器
// n:读取长度
// err:错误
func (b *Buffer) WriteTo(w io.Writer) (n int64, err error) { //md5:c9cb2afb71c7aea9
	return b.F.WriteTo(w)
}

// WriteByte 将字节 c 追加到缓冲区，如果需要会扩大缓冲区。返回的错误始终为 nil，但包含它以匹配 [bufio.Writer] 的 WriteByte 方法。如果缓冲区变得太大，WriteByte 会用 [ErrTooLarge] 引发 panic。
// md5:3807962fe1acfbac
// ff:写字节
// b:
// c:字节
func (b *Buffer) WriteByte(c byte) error { //md5:d3b248f48547dbf9
	return b.F.WriteByte(c)
}

// WriteRune 将Unicode代码点r的UTF-8编码附加到缓冲区中，
// 并返回其长度和一个错误，该错误总是为nil，但为了与[bufio.Writer]的WriteRune方法匹配而包含在内。
// 缓冲区会根据需要自动扩展；如果缓冲区变得过大，WriteRune将会因[ErrTooLarge]错误而引发恐慌。
// md5:fcd3a4d673494541
// ff:写字符
// b:
// r:字符
// n:写入长度
// err:错误
func (b *Buffer) WriteRune(r rune) (n int, err error) { //md5:488284b1d4dd43e2
	return b.F.WriteRune(r)
}

// Read 从缓冲区中读取接下来的 len(p) 个字节，或者直到缓冲区被读空。返回值 n 是读取的字节数。如果缓冲区没有数据可读，err 将是 io.EOF（除非 len(p) 为零）；否则 err 为 nil。
// md5:6dec72f7698c2809
// ff:
// b:
// p:字节集变量
// n:读取长度
// err:错误
func (b *Buffer) Read(p []byte) (n int, err error) { //md5:d77b54302f43331e
	return b.F.Read(p)
}

// Next 函数从缓冲区中返回包含接下来 n 个字节的切片，就像 [Buffer.Read] 返回的一样。如果缓冲区中的字节数少于 n，Next 将返回整个缓冲区。这个切片在下一次调用读写方法之前都是有效的。
// md5:4b6d65f5719e684a
// ff:读取指定长度字节集
// b:
// n:
func (b *Buffer) Next(n int) []byte { //md5:e9b6b3c188f56ffc
	return b.F.Next(n)
}

// ReadByte 从缓冲区读取并返回下一个字节。如果没有可用的字节，它将返回错误 io.EOF。
// md5:9a65c587bf66dcf3
// ff:读取字节
// b:
func (b *Buffer) ReadByte() (byte, error) { //md5:8e7cb56a75e9df65
	return b.F.ReadByte()
}

// ReadRune 读取并返回缓冲区中的下一个UTF-8编码的Unicode码点。
// 如果没有可用的字节，返回的错误是io.EOF。
// 如果字节是错误的UTF-8编码，它将消耗一个字节并返回U+FFFD，1。
// md5:e6188dfcfac12306
// ff:读取字符
// b:
// r:字符
// size:长度
// err:错误
func (b *Buffer) ReadRune() (r rune, size int, err error) { //md5:9e9f745d7cd0e7b2
	return b.F.ReadRune()
}

// UnreadRune 回退由 Buffer.ReadRune 返回的最后一个字符。
// 如果对缓冲区的最近一次读写操作不是成功的 ReadRune，UnreadRune 将返回一个错误。
// 在这方面，它比 Buffer.UnreadByte 更严格，因为 Buffer.UnreadByte 可以回退任何读取操作的最后一个字节。
// md5:239e0a258391add0
// ff:反向读取字符
// b:
func (b *Buffer) UnreadRune() error { //md5:3e89d3d1a29568ab
	return b.F.UnreadRune()
}

// UnreadByte 将最近一次成功读取（至少读取了一个字节）后返回的最后一个字节 unread。如果在上次读取之后有写操作发生，或者上次读取返回了错误，或者读取了0个字节，UnreadByte 将返回一个错误。
// md5:ebefc8dc07168356
// ff:反向读取字节
// b:
func (b *Buffer) UnreadByte() error { //md5:1635a851b5d4816b
	return b.F.UnreadByte()
}

// ReadBytes 读取到输入中的第一个分隔符，返回包含到和包括分隔符的数据切片。如果在找到分隔符之前遇到错误，它将返回错误发生前已读取的数据以及错误（通常是 io.EOF）。只有当返回的数据不以分隔符结束时，ReadBytes 才会返回 err != nil。
// md5:4d18e9ed50415556
// ff:读取至分隔符
// b:
// delim:分隔符
// line:字节集
// err:错误
func (b *Buffer) ReadBytes(delim byte) (line []byte, err error) { //md5:43a270cb44071660
	return b.F.ReadBytes(delim)
}

// ReadString 读取输入直到遇到第一个 delim 字符为止，
// 并返回一个字符串，其中包含直到分隔符（包括分隔符本身）的所有数据。
// 如果 ReadString 在找到分隔符之前遇到错误，
// 它会返回遇到错误之前读取的数据以及该错误本身（常为 io.EOF）。
// ReadString 在且仅在返回的数据没有以 delim 结尾时，返回 err != nil。
// md5:2dbe6334091e50c0
// ff:读取文本至分隔符
// b:
// delim:分隔符
// line:文本
// err:错误
func (b *Buffer) ReadString(delim byte) (line string, err error) { //md5:d0fb6e8518de22bf
	return b.F.ReadString(delim)
}

// NewBuffer 使用 buf 作为初始内容创建并初始化一个新的 [Buffer]。新的 [Buffer] 将接管 buf，调用者在调用后不应再使用 buf。NewBuffer 的目的是准备一个 [Buffer] 来读取已有数据。它也可以用来设置内部缓冲区的初始写入大小。为此，buf 应具有所需的容量，但长度应为零。
//
// 在大多数情况下，new([Buffer])（或者直接声明一个 [Buffer] 变量）就足以初始化一个 [Buffer]。
// md5:160089e06b42c9d0
// ff:创建缓冲区并按字节集
// buf:字节集
func NewBuffer(buf []byte) *Buffer { //md5:eb38ebfe0970b97b
	b := bytes.NewBuffer(buf)
	if b == nil {
		return nil
	}
	return &Buffer{*b}
}

// NewBufferString 通过字符串 s 创建并初始化一个新的 [Buffer]。它的目的是准备一个缓冲区来读取已存在的字符串。
//
// 在大多数情况下，使用 new([Buffer])（或者只是声明一个 [Buffer] 变量）就足以初始化一个 [Buffer]。
// md5:52db730cdb1471dd
// ff:创建缓冲区并按文本
// s:文本
func NewBufferString(s string) *Buffer { //md5:65590cc504d5130e
	b := bytes.NewBufferString(s)
	if b == nil {
		return nil
	}
	return &Buffer{*b}
}

// zj:
// String 方法返回缓冲区未读部分内容作为字符串。
// 如果 [Buffer] 是一个空指针，它将返回 "<nil>"。
//
// 要更高效地构建字符串，请参见 strings.Builder 类型。
func (b *Buffer) X取未读文本() string { //md5:2a07d671bd398c91
	return b.F.String()
}

// Read 从缓冲区中读取接下来的 len(p) 个字节，或者直到缓冲区被读空。返回值 n 是读取的字节数。如果缓冲区没有数据可读，err 将是 io.EOF（除非 len(p) 为零）；否则 err 为 nil。
func (b *Buffer) X读取字节集(字节集变量 []byte) (读取长度 int, 错误 error) { //md5:d77b54302f43331e
	return b.F.Read(字节集变量)
}

// Write 将 p 的内容追加到缓冲区，如果需要会扩大缓冲区。返回值 n 是 p 的长度；err 总是为 nil。如果缓冲区变得过大，Write 会引发一个 [ErrTooLarge] 的 panic。
func (b *Buffer) X写入字节集(写入字节集 []byte) (写入长度 int, 错误 error) { //md5:6f1eceb67804b19f
	return b.F.Write(写入字节集)
}

//zj:
