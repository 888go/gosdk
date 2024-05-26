package strings

import (
	"io"
	"strings"
)

// Reader结构体通过从字符串中读取数据来实现io.Reader、io.ReaderAt、io.ByteReader、io.ByteScanner、io.RuneReader、io.RuneScanner、io.Seeker和io.WriterTo接口。
// Reader的零值像一个空字符串的Reader。
type Reader struct {
	F strings.Reader
} //md5:e4cbd56cfee19521

// NewReader 返回一个从 s 读取的新的 [Reader]。
// 它类似于 [bytes.NewBufferString]，但更高效且不可写。

// ff:
// s:
func NewReader(s string) *Reader { //md5:eb6092e988edffe96d2da2029f5274e8
	return &Reader{
		F: *strings.NewReader(s),
	}
}

// Len 返回字符串未读部分的字节数。

// ff:
func (r *Reader) Len() int { //md5:1aa1efbca526400e02843fb1c8e27012
	return r.F.Len()
}

// Size 返回底层字符串的原始长度。
// Size 是通过 [Reader.ReadAt] 可读取的字节数。
// 返回的值始终相同，不受对任何其他方法调用的影响。

// ff:
func (r *Reader) Size() int64 { //md5:267dfc9e54f5f5cacb1b8e7004597412
	return r.F.Size()
}

// Read 实现了 [io.Reader] 接口。

// ff:
// b:
// n:
// err:
func (r *Reader) Read(b []byte) (n int, err error) { //md5:817ec22fd5a0598255ff61f03f08774f
	return r.F.Read(b)
}

// 实现 io.ReaderAt 接口的 ReadAt 方法
// 翻译提示:func  (r  *阅读器)  在位置读取(b  []字节,  偏移量  int64)  (读取长度  int,  错误  error)  {}

// ff:
// b:
// off:
// n:
// err:
func (r *Reader) ReadAt(b []byte, off int64) (n int, err error) { //md5:398d6e4aca3bd8cf1674a648b664ad07
	return r.F.ReadAt(b, off)
}

// ReadByte 实现了 [io.ByteReader] 接口。

// ff:
func (r *Reader) ReadByte() (byte, error) { //md5:34273d2f33dc39641f493a3938eca137
	return r.F.ReadByte()
}

// UnreadByte 实现了 [io.ByteScanner] 接口。

// ff:
func (r *Reader) UnreadByte() error { //md5:ac1435a66b739b0088b37a2a3ced6acc
	return r.F.UnreadByte()
}

// 实现了 io.RuneReader 接口的 ReadRune 方法
// 翻译提示:func  (r  *读取器)  读取字符()  (字符  ch  rune,  大小  size  int,  错误  err  error)  {}

// ff:
// ch:
// size:
// err:
func (r *Reader) ReadRune() (ch rune, size int, err error) { //md5:21ad75ec2d3037fd8e43b28f34f786ce
	return r.F.ReadRune()
}

// UnreadRune 实现了 [io.RuneScanner] 接口。

// ff:
func (r *Reader) UnreadRune() error { //md5:42e1193e492b8c32a40d243685d919df
	return r.F.UnreadRune()
}

// 实现了 [io.Seeker] 接口的 Seek 方法
// 翻译提示:func  (r  *阅读器)  寻找(offset  int64,  方向  int)  (位置  int64,  错误  error)  {}

// ff:
// offset:
// whence:
func (r *Reader) Seek(offset int64, whence int) (int64, error) { //md5:3c5ed8431f7dff12e0782fb5c343b1dc
	return r.F.Seek(offset, whence)
}

// 实现了 io.WriterTo 接口的 WriteTo 方法
// 翻译提示:func  (r  *读取器)  写入到(w  io.写入器)  (字节写入数  int64,  错误  error)  {}

// ff:
// w:
// n:
// err:
func (r *Reader) WriteTo(w io.Writer) (n int64, err error) { //md5:6bec38e7a3c36070ba76e2f596818440
	return r.F.WriteTo(w)
}

// Reset 将[Reader]重置为从s读取。

// ff:
// s:
func (r *Reader) Reset(s string) { //md5:cd44cb38cce65aba66e071f854e0d882
	r.F.Reset(s)
}
