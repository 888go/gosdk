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

1111