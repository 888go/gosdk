package strings

// String 返回累计生成的字符串。
func (b *Builder) String() string { //md5:a46ceb48bbded6ebe25a9e8bc750e7ed

}

// Len 返回已累积的字节数；b.Len() 等于 len(b.String())。
func (b *Builder) Len() int { //md5:44cf05e8c04610821bfd318be49cee4f

}

// Cap 返回构建器底层字节切片的容量。它是指为正在构建的字符串所分配的总空间，包括已写入的所有字节。
func (b *Builder) Cap() int { //md5:4f3248513ca7b7f06fd67a832792520b

}

// Reset 重置 [Builder] 为空。
func (b *Builder) Reset() { //md5:2ae0012c2cfc3d0258ebe346a66c40a7

}

// Grow 会根据需要增加 b 的容量，以确保为接下来的 n 个字节提供足够的空间。在调用 Grow(n) 后，至少可以向 b 中写入 n 个字节而无需再进行一次分配。如果 n 为负值，Grow 将引发恐慌（panic）。
func (b *Builder) Grow(n int) { //md5:8f4946bccba8156a276c45af7e766afe

}

// Write 将 p 的内容追加至 b 的缓冲区。
// Write 总是返回 len(p), nil。
func (b *Builder) Write(p []byte) (int, error) { //md5:1b06c45a1dd29dd22b7e9639b5451bcf

}

// WriteByte 将字节 c 追加至 b 的缓冲区。
// 返回的错误始终为 nil。
func (b *Builder) WriteByte(c byte) error { //md5:00e290578c1ef840c83c950cee9da6c3

}

// WriteRune 将Unicode码点r的UTF-8编码追加到b的缓冲区中。
// 它返回r的长度（字节数）并返回一个nil错误。
func (b *Builder) WriteRune(r rune) (int, error) { //md5:515b9053f4031f5d53d1528d885d327c

}

// WriteString 将 s 的内容追加到 b 的缓冲区中。
// 它返回 s 的长度以及一个空（nil）错误。
func (b *Builder) WriteString(s string) (int, error) { //md5:cd0e8767fd6968e12a68b3980fb0c347

}
