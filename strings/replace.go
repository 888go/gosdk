package strings

import (
	"io"
	"strings"
)

type Replacer struct {
	F strings.Replacer
} //md5:5a2f93739712f67f

// NewReplacer 从一组旧字符串、新字符串对中返回一个新的[Replacer]。替换操作按照它们在目标字符串中出现的顺序进行，且不涉及重叠匹配。旧字符串的比较按参数顺序进行。
//
// 如果提供的参数数目为奇数，NewReplacer 将引发 panic。

// ff:
// oldnew:
func NewReplacer(oldnew ...string) *Replacer { //md5:626ca1fb5d14db50e55f8da7a96f583c
	return &Replacer{
		F: *strings.NewReplacer(oldnew...),
	}
}

// Replace返回对s进行所有替换操作后的副本。

// ff:
// s:
func (r *Replacer) Replace(s string) string { //md5:c2a7d68dc010b9622ed8d85f2cc0bf97
	return r.F.Replace(s)
}

// WriteString 将s写入w，同时执行所有替换操作。

// ff:
// err:
// n:
// s:
// w:
func (r *Replacer) WriteString(w io.Writer, s string) (n int, err error) { //md5:e6848c200491de80abff300f82fb7068
	return r.F.WriteString(w, s)
}

//
//// Write 向缓冲区写入数据以实现 io.Writer 接口
//func (w *appendSliceWriter) Write(p []byte) (int, error) {  //md5:43016f2236bda21b54e776d513e4fbad
//
//}
//
//// WriteString 向缓冲区写入字符串，无需进行 string->[]byte->string 的内存分配。
//func (w *appendSliceWriter) WriteString(s string) (int, error) {  //md5:061548bbb2419e7c72837b713992e8e9
//
//}
//
//func (w stringWriter) WriteString(s string) (int, error) {  //md5:d562a4570958ed3f64d1889c2b8c464b
//
//}
//
//func (r *genericReplacer) Replace(s string) string {  //md5:22c233aa07f8d7685224736b9a8f4c4e
//
//}
//
//func (r *genericReplacer) WriteString(w io.Writer, s string) (n int, err error) {  //md5:cd0bf554bc800b43bffc1886fe3f0efe
//
//}
//
//func (r *singleStringReplacer) Replace(s string) string {  //md5:4c5fcc31c63e0af703eb3c8888621b97
//
//}
//
//func (r *singleStringReplacer) WriteString(w io.Writer, s string) (n int, err error) {  //md5:f2e60649ac25f427463fe5074a16e454
//
//}
//
//func (r *byteReplacer) Replace(s string) string {  //md5:209bdcdfa0cceb410c68d2d6d7c73e85
//
//}
//
//func (r *byteReplacer) WriteString(w io.Writer, s string) (n int, err error) {  //md5:aa33594b3168bb1630a56eefb4cb9622
//
//}
//
//func (r *byteStringReplacer) Replace(s string) string {  //md5:4d91d6eb4a77ac6f6608db7c97c068cd
//
//}
//
//func (r *byteStringReplacer) WriteString(w io.Writer, s string) (n int, err error) {  //md5:fe07adb1ac7dc9a9d057b72d1adc3a0c
//
//}
