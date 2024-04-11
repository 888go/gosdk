package strings

// Len 返回字符串未读部分的字节数。
func (r *Reader) Len() int { //md5:1aa1efbca526400e02843fb1c8e27012

}

// Size 返回底层字符串的原始长度。
// Size 是通过 [Reader.ReadAt] 可读取的字节数。
// 返回的值始终相同，不受对任何其他方法调用的影响。
func (r *Reader) Size() int64 { //md5:267dfc9e54f5f5cacb1b8e7004597412

}

// Read 实现了 [io.Reader] 接口。
func (r *Reader) Read(b []byte) (n int, err error) { //md5:817ec22fd5a0598255ff61f03f08774f

}

// 实现 io.ReaderAt 接口的 ReadAt 方法
func (r *Reader) ReadAt(b []byte, off int64) (n int, err error) { //md5:398d6e4aca3bd8cf1674a648b664ad07

}

// ReadByte 实现了 [io.ByteReader] 接口。
func (r *Reader) ReadByte() (byte, error) { //md5:34273d2f33dc39641f493a3938eca137

}

// UnreadByte 实现了 [io.ByteScanner] 接口。
func (r *Reader) UnreadByte() error { //md5:ac1435a66b739b0088b37a2a3ced6acc

}

// 实现了 io.RuneReader 接口的 ReadRune 方法
func (r *Reader) ReadRune() (ch rune, size int, err error) { //md5:21ad75ec2d3037fd8e43b28f34f786ce

}

// UnreadRune 实现了 [io.RuneScanner] 接口。
func (r *Reader) UnreadRune() error { //md5:42e1193e492b8c32a40d243685d919df

}

// 实现了 [io.Seeker] 接口的 Seek 方法
func (r *Reader) Seek(offset int64, whence int) (int64, error) { //md5:3c5ed8431f7dff12e0782fb5c343b1dc

}

// 实现了 io.WriterTo 接口的 WriteTo 方法
func (r *Reader) WriteTo(w io.Writer) (n int64, err error) { //md5:6bec38e7a3c36070ba76e2f596818440

}

// Reset 将[Reader]重置为从s读取。
func (r *Reader) Reset(s string) { //md5:cd44cb38cce65aba66e071f854e0d882

}

// NewReader 返回一个从 s 读取的新的 [Reader]。
// 它类似于 [bytes.NewBufferString]，但更高效且不可写。
func NewReader(s string) *Reader { //md5:eb6092e988edffe96d2da2029f5274e8

}
