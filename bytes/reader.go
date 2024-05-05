package bytes

import (
	"bytes"
	"io"
)

// Reader实现了io.Reader、io.ReaderAt、io.WriterTo、io.Seeker、io.ByteScanner和io.RuneScanner接口，通过从字节切片中读取数据。
// 与[Buffer]不同，Reader是只读的，并且支持定位（seeking）操作。
// Reader的零值行为就像一个空切片的Reader。
// md5:c7615bfef4c8b2fe
type Reader struct {
	F bytes.Reader
} //md5:e4cbd56cfee19521

// Len 返回切片未读部分的字节长度。
// md5:60d512b53434757a

// ff:
func (r *Reader) Len() int { //md5:1aa1efbca526400e
	return r.F.Len()
}

// Size 返回原始字节切片的长度。
// Size 是通过 [Reader.ReadAt] 可读取的字节数量。
// 结果不受除 [Reader.Reset] 之外的任何方法调用影响。
// md5:7057c00761109102

// ff:
func (r *Reader) Size() int64 { //md5:267dfc9e54f5f5ca
	return r.F.Size()
}

// Read implements the [io.Reader] 接口。. md5:6b6e15958e2263c0

// ff:
// err:
// n:
// b:
func (r *Reader) Read(b []byte) (n int, err error) { //md5:817ec22fd5a05982
	return r.F.Read(b)
}

// ReadAt 实现了 [io.ReaderAt] 接口。. md5:aff23aac8adb605e

// ff:
// err:
// n:
// off:
// b:
func (r *Reader) ReadAt(b []byte, off int64) (n int, err error) { //md5:398d6e4aca3bd8cf
	return r.F.ReadAt(b, off)
}

// ReadByte 实现了 [io.ByteReader] 接口。. md5:9580ddd72af46f48

// ff:
func (r *Reader) ReadByte() (byte, error) { //md5:34273d2f33dc3964
	return r.F.ReadByte()
}

// UnreadByte 在实现 [io.ByteScanner] 接口时补充了 [Reader.ReadByte] 的功能。. md5:5aa8f5cc398bf852

// ff:
func (r *Reader) UnreadByte() error { //md5:ac1435a66b739b00
	return r.F.UnreadByte()
}

// ReadRune 实现了 [io.RuneReader] 接口。. md5:97f6bcc2e44946c6

// ff:
// err:
// size:
// ch:
func (r *Reader) ReadRune() (ch rune, size int, err error) { //md5:21ad75ec2d3037fd
	return r.F.ReadRune()
}

// UnreadRune 在实现 [io.RuneScanner] 接口时补充了 [Reader.ReadRune]。. md5:d0948bd1ce43eeef

// ff:
func (r *Reader) UnreadRune() error { //md5:42e1193e492b8c32
	return r.F.UnreadRune()
}

// Seek 实现了 [io.Seeker] 接口。. md5:407fc18ddaadacba

// ff:
// whence:
// offset:
func (r *Reader) Seek(offset int64, whence int) (int64, error) { //md5:3c5ed8431f7dff12
	return r.F.Seek(offset, whence)
}

// WriteTo 实现了 [io.WriterTo] 接口。. md5:93bec25fa975a33f

// ff:
// err:
// n:
// w:
func (r *Reader) WriteTo(w io.Writer) (n int64, err error) { //md5:6bec38e7a3c36070
	return r.F.WriteTo(w)
}

// Reset 将[Reader.Reader]重置为从b读取。. md5:3c05fdb3a579bd62

// ff:
// b:
func (r *Reader) Reset(b []byte) { //md5:cd44cb38cce65aba
	r.F.Reset(b)
}

// NewReader 返回一个新的从 b 读取的 [Reader.Reader]。. md5:807cf3d6e5b558e1

// ff:
// b:
func NewReader(b []byte) *Reader { //md5:eb6092e988edffe9
	r := bytes.NewReader(b)
	if r == nil {
		return nil
	}
	return &Reader{*r}
}
