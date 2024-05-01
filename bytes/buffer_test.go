//版权所有2009年Go作者。所有权利保留。
//使用此源代码受BSD风格
//可以在LICENSE文件中找到的许可证。
// md5:2e9dc81828a3be8a

package bytes_test

import (
	"fmt"
	. "github.com/888go/gosdk/bytes"
	"io"
	"math/rand"
	"strconv"
	"testing"
	"unicode/utf8"
)

const N = 10000       // 将其增大以便进行更大（但更慢）的测试. md5:bf0c2bf33a8d9434
var testString string // 写入测试的测试数据. md5:fb12fba62e7a297f
var testBytes []byte  // 测试数据；与testString相同，但以切片形式表示。. md5:84211fa4bf8f2e04

type negativeReader struct{}

func (r *negativeReader) Read([]byte) (int, error) { return -1, nil }

func init() {
	testBytes = make([]byte, N)
	for i := 0; i < N; i++ {
		testBytes[i] = 'a' + byte(i%26)
	}
	testString = string(testBytes)
}

// 验证buf中的内容与字符串s匹配。. md5:126cb1dbac61397b
func check(t *testing.T, testname string, buf *Buffer, s string) {
	bytes := buf.Bytes()
	str := buf.String()
	if buf.Len() != len(bytes) {
		t.Errorf("%s: buf.Len() == %d, len(buf.Bytes()) == %d", testname, buf.Len(), len(bytes))
	}

	if buf.Len() != len(str) {
		t.Errorf("%s: buf.Len() == %d, len(buf.String()) == %d", testname, buf.Len(), len(str))
	}

	if buf.Len() != len(s) {
		t.Errorf("%s: buf.Len() == %d, len(s) == %d", testname, buf.Len(), len(s))
	}

	if string(bytes) != s {
		t.Errorf("%s: string(buf.Bytes()) == %q, s == %q", testname, string(bytes), s)
	}
}

// 通过将字符串fus写入n次来填充buf。buf的初始内容对应于字符串s；返回的结果是buf的最终内容，以字符串形式返回。
// md5:4a2535b01b6cdaaa
func fillString(t *testing.T, testname string, buf *Buffer, s string, n int, fus string) string {
	check(t, testname+" (fill 1)", buf, s)
	for ; n > 0; n-- {
		m, err := buf.WriteString(fus)
		if m != len(fus) {
			t.Errorf(testname+" (fill 2): m == %d, expected %d", m, len(fus))
		}
		if err != nil {
			t.Errorf(testname+" (fill 3): err should always be nil, found err == %s", err)
		}
		s += fus
		check(t, testname+" (fill 4)", buf, s)
	}
	return s
}

// 通过n次写入字节片fub来填充buf。
// buf的初始内容对应于字符串s；
// 结果是以字符串形式返回buf的最终内容。
// md5:ffa83d80493fdd8d
func fillBytes(t *testing.T, testname string, buf *Buffer, s string, n int, fub []byte) string {
	check(t, testname+" (fill 1)", buf, s)
	for ; n > 0; n-- {
		m, err := buf.Write(fub)
		if m != len(fub) {
			t.Errorf(testname+" (fill 2): m == %d, expected %d", m, len(fub))
		}
		if err != nil {
			t.Errorf(testname+" (fill 3): err should always be nil, found err == %s", err)
		}
		s += string(fub)
		check(t, testname+" (fill 4)", buf, s)
	}
	return s
}

func TestNewBuffer(t *testing.T) {
	buf := NewBuffer(testBytes)
	check(t, "NewBuffer", buf, testString)
}

func TestNewBufferString(t *testing.T) {
	buf := NewBufferString(testString)
	check(t, "NewBufferString", buf, testString)
}

// 通过反复读取到fub中清空buf。
// buf的初始内容对应于字符串s。
// md5:930d750514b1c416
func empty(t *testing.T, testname string, buf *Buffer, s string, fub []byte) {
	check(t, testname+" (empty 1)", buf, s)

	for {
		n, err := buf.Read(fub)
		if n == 0 {
			break
		}
		if err != nil {
			t.Errorf(testname+" (empty 2): err should always be nil, found err == %s", err)
		}
		s = s[n:]
		check(t, testname+" (empty 3)", buf, s)
	}

	check(t, testname+" (empty 4)", buf, "")
}

func TestBasicOperations(t *testing.T) {
	var buf Buffer

	for i := 0; i < 5; i++ {
		check(t, "TestBasicOperations (1)", &buf, "")

		buf.Reset()
		check(t, "TestBasicOperations (2)", &buf, "")

		buf.Truncate(0)
		check(t, "TestBasicOperations (3)", &buf, "")

		n, err := buf.Write(testBytes[0:1])
		if want := 1; err != nil || n != want {
			t.Errorf("Write: got (%d, %v), want (%d, %v)", n, err, want, nil)
		}
		check(t, "TestBasicOperations (4)", &buf, "a")

		buf.WriteByte(testString[1])
		check(t, "TestBasicOperations (5)", &buf, "ab")

		n, err = buf.Write(testBytes[2:26])
		if want := 24; err != nil || n != want {
			t.Errorf("Write: got (%d, %v), want (%d, %v)", n, err, want, nil)
		}
		check(t, "TestBasicOperations (6)", &buf, testString[0:26])

		buf.Truncate(26)
		check(t, "TestBasicOperations (7)", &buf, testString[0:26])

		buf.Truncate(20)
		check(t, "TestBasicOperations (8)", &buf, testString[0:20])

		empty(t, "TestBasicOperations (9)", &buf, testString[0:20], make([]byte, 5))
		empty(t, "TestBasicOperations (10)", &buf, "", make([]byte, 100))

		buf.WriteByte(testString[1])
		c, err := buf.ReadByte()
		if want := testString[1]; err != nil || c != want {
			t.Errorf("ReadByte: got (%q, %v), want (%q, %v)", c, err, want, nil)
		}
		c, err = buf.ReadByte()
		if err != io.EOF {
			t.Errorf("ReadByte: got (%q, %v), want (%q, %v)", c, err, byte(0), io.EOF)
		}
	}
}

func TestLargeStringWrites(t *testing.T) {
	var buf Buffer
	limit := 30
	if testing.Short() {
		limit = 9
	}
	for i := 3; i < limit; i += 3 {
		s := fillString(t, "TestLargeWrites (1)", &buf, "", 5, testString)
		empty(t, "TestLargeStringWrites (2)", &buf, s, make([]byte, len(testString)/i))
	}
	check(t, "TestLargeStringWrites (3)", &buf, "")
}

func TestLargeByteWrites(t *testing.T) {
	var buf Buffer
	limit := 30
	if testing.Short() {
		limit = 9
	}
	for i := 3; i < limit; i += 3 {
		s := fillBytes(t, "TestLargeWrites (1)", &buf, "", 5, testBytes)
		empty(t, "TestLargeByteWrites (2)", &buf, s, make([]byte, len(testString)/i))
	}
	check(t, "TestLargeByteWrites (3)", &buf, "")
}

func TestLargeStringReads(t *testing.T) {
	var buf Buffer
	for i := 3; i < 30; i += 3 {
		s := fillString(t, "TestLargeReads (1)", &buf, "", 5, testString[0:len(testString)/i])
		empty(t, "TestLargeReads (2)", &buf, s, make([]byte, len(testString)))
	}
	check(t, "TestLargeStringReads (3)", &buf, "")
}

func TestLargeByteReads(t *testing.T) {
	var buf Buffer
	for i := 3; i < 30; i += 3 {
		s := fillBytes(t, "TestLargeReads (1)", &buf, "", 5, testBytes[0:len(testBytes)/i])
		empty(t, "TestLargeReads (2)", &buf, s, make([]byte, len(testString)))
	}
	check(t, "TestLargeByteReads (3)", &buf, "")
}

func TestMixedReadsAndWrites(t *testing.T) {
	var buf Buffer
	s := ""
	for i := 0; i < 50; i++ {
		wlen := rand.Intn(len(testString))
		if i%2 == 0 {
			s = fillString(t, "TestMixedReadsAndWrites (1)", &buf, s, 1, testString[0:wlen])
		} else {
			s = fillBytes(t, "TestMixedReadsAndWrites (1)", &buf, s, 1, testBytes[0:wlen])
		}

		rlen := rand.Intn(len(testString))
		fub := make([]byte, rlen)
		n, _ := buf.Read(fub)
		s = s[n:]
	}
	empty(t, "TestMixedReadsAndWrites (2)", &buf, s, make([]byte, buf.Len()))
}

func TestCapWithPreallocatedSlice(t *testing.T) {
	buf := NewBuffer(make([]byte, 10))
	n := buf.Cap()
	if n != 10 {
		t.Errorf("expected 10, got %d", n)
	}
}

func TestCapWithSliceAndWrittenData(t *testing.T) {
	buf := NewBuffer(make([]byte, 0, 10))
	buf.Write([]byte("test"))
	n := buf.Cap()
	if n != 10 {
		t.Errorf("expected 10, got %d", n)
	}
}

//func TestNil(t *testing.T) {
//	var b *Buffer
//	if b.String() != "<nil>" {
//		t.Errorf("expected <nil>; got %q", b.String())
//	}
//}

func TestReadFrom(t *testing.T) {
	var buf Buffer
	for i := 3; i < 30; i += 3 {
		s := fillBytes(t, "TestReadFrom (1)", &buf, "", 5, testBytes[0:len(testBytes)/i])
		var b Buffer
		b.ReadFrom(&buf)
		empty(t, "TestReadFrom (2)", &b, s, make([]byte, len(testString)))
	}
}

type panicReader struct{ panic bool }

func (r panicReader) Read(p []byte) (int, error) {
	if r.panic {
		panic("oops")
	}
	return 0, io.EOF
}

// 确保在发生恐慌的读取之前，一个空的Buffer在增长（"grown"）后仍然保持为空。
// md5:0e205b51231f8eb6
func TestReadFromPanicReader(t *testing.T) {

	// 首先验证非 panic 行为. md5:d0442d9c9eefa946
	var buf Buffer
	i, err := buf.ReadFrom(panicReader{})
	if err != nil {
		t.Fatal(err)
	}
	if i != 0 {
		t.Fatalf("unexpected return from bytes.ReadFrom (1): got: %d, want %d", i, 0)
	}
	check(t, "TestReadFromPanicReader (1)", &buf, "")

	// 确认当Reader发生恐慌时，空缓冲区仍然保持为空. md5:1e5afdd5554f8adb
	var buf2 Buffer
	defer func() {
		recover()
		check(t, "TestReadFromPanicReader (2)", &buf2, "")
	}()
	buf2.ReadFrom(panicReader{panic: true})
}

func TestReadFromNegativeReader(t *testing.T) {
	var b Buffer
	defer func() {
		switch err := recover().(type) {
		case nil:
			t.Fatal("bytes.Buffer.ReadFrom didn't panic")
		case error:
			// 这是errNegativeRead错误的字符串表示. md5:1f249ec61c28b754
			wantError := "bytes.Buffer: reader returned negative count from Read"
			if err.Error() != wantError {
				t.Fatalf("recovered panic: got %v, want %v", err.Error(), wantError)
			}
		default:
			t.Fatalf("unexpected panic value: %#v", err)
		}
	}()

	b.ReadFrom(new(negativeReader))
}

func TestWriteTo(t *testing.T) {
	var buf Buffer
	for i := 3; i < 30; i += 3 {
		s := fillBytes(t, "TestWriteTo (1)", &buf, "", 5, testBytes[0:len(testBytes)/i])
		var b Buffer
		buf.WriteTo(&b)
		empty(t, "TestWriteTo (2)", &b, s, make([]byte, len(testString)))
	}
}

func TestWriteAppend(t *testing.T) {
	var got Buffer
	var want []byte
	for i := 0; i < 1000; i++ {
		b := got.AvailableBuffer()
		b = strconv.AppendInt(b, int64(i), 10)
		want = strconv.AppendInt(want, int64(i), 10)
		got.Write(b)
	}
	if !Equal(got.Bytes(), want) {
		t.Fatalf("Bytes() = %q, want %q", got, want)
	}

	// 如果缓冲区足够大，就无需分配内存。. md5:91e6a0658a4d1773
	n := testing.AllocsPerRun(100, func() {
		got.Reset()
		for i := 0; i < 1000; i++ {
			b := got.AvailableBuffer()
			b = strconv.AppendInt(b, int64(i), 10)
			got.Write(b)
		}
	})
	if n > 0 {
		t.Errorf("allocations occurred while appending")
	}
}

func TestRuneIO(t *testing.T) {
	const NRune = 1000
	// 在写入数据的同时，构建测试切片. md5:297726db76b7b32f
	b := make([]byte, utf8.UTFMax*NRune)
	var buf Buffer
	n := 0
	for r := rune(0); r < NRune; r++ {
		size := utf8.EncodeRune(b[n:], r)
		nbytes, err := buf.WriteRune(r)
		if err != nil {
			t.Fatalf("WriteRune(%U) error: %s", r, err)
		}
		if nbytes != size {
			t.Fatalf("WriteRune(%U) expected %d, got %d", r, size, nbytes)
		}
		n += size
	}
	b = b[0:n]

	// 检查生成的字节. md5:d96610169d4659e0
	if !Equal(buf.Bytes(), b) {
		t.Fatalf("incorrect result from WriteRune: %q not %q", buf.Bytes(), b)
	}

	p := make([]byte, utf8.UTFMax)
	// 通过ReadRune读取回来. md5:7eff8ed0af3ba587
	for r := rune(0); r < NRune; r++ {
		size := utf8.EncodeRune(p, r)
		nr, nbytes, err := buf.ReadRune()
		if nr != r || nbytes != size || err != nil {
			t.Fatalf("ReadRune(%U) got %U,%d not %U,%d (err=%s)", r, nr, nbytes, r, size, err)
		}
	}

	// 检查UnreadRune是否正常工作. md5:4bee260d2e0cfec6
	buf.Reset()

	// check at EOF
	if err := buf.UnreadRune(); err == nil {
		t.Fatal("UnreadRune at EOF: got no error")
	}
	if _, _, err := buf.ReadRune(); err == nil {
		t.Fatal("ReadRune at EOF: got no error")
	}
	if err := buf.UnreadRune(); err == nil {
		t.Fatal("UnreadRune after ReadRune at EOF: got no error")
	}

	// check not at EOF
	buf.Write(b)
	for r := rune(0); r < NRune; r++ {
		r1, size, _ := buf.ReadRune()
		if err := buf.UnreadRune(); err != nil {
			t.Fatalf("UnreadRune(%U) got error %q", r, err)
		}
		r2, nbytes, err := buf.ReadRune()
		if r1 != r2 || r1 != r || nbytes != size || err != nil {
			t.Fatalf("ReadRune(%U) after UnreadRune got %U,%d not %U,%d (err=%s)", r, r2, nbytes, r, size, err)
		}
	}
}

func TestWriteInvalidRune(t *testing.T) {
	// 无效的 runes，包括负值，应写为 utf8.RuneError。
	// md5:681b779ca55ae8d0
	for _, r := range []rune{-1, utf8.MaxRune + 1} {
		var buf Buffer
		buf.WriteRune(r)
		check(t, fmt.Sprintf("TestWriteInvalidRune (%d)", r), &buf, "\uFFFD")
	}
}

func TestNext(t *testing.T) {
	b := []byte{0, 1, 2, 3, 4}
	tmp := make([]byte, 5)
	for i := 0; i <= 5; i++ {
		for j := i; j <= 5; j++ {
			for k := 0; k <= 6; k++ {
				// 0 <= i <= j <= 5; 0 <= k <= 6
				// 检查如果我们从偏移量i处的长度为j的缓冲区开始，并请求Next(k)，我们得到正确的字节。
				// md5:e2c406f6e1abd245
				buf := NewBuffer(b[0:j])
				n, _ := buf.Read(tmp[0:i])
				if n != i {
					t.Fatalf("Read %d returned %d", i, n)
				}
				bb := buf.Next(k)
				want := k
				if want > j-i {
					want = j - i
				}
				if len(bb) != want {
					t.Fatalf("in %d,%d: len(Next(%d)) == %d", i, j, k, len(bb))
				}
				for l, v := range bb {
					if v != byte(l+i) {
						t.Fatalf("in %d,%d: Next(%d)[%d] = %d, want %d", i, j, k, l, v, l+i)
					}
				}
			}
		}
	}
}

var readBytesTests = []struct {
	buffer   string
	delim    byte
	expected []string
	err      error
}{
	{"", 0, []string{""}, io.EOF},
	{"a\x00", 0, []string{"a\x00"}, nil},
	{"abbbaaaba", 'b', []string{"ab", "b", "b", "aaab"}, nil},
	{"hello\x01world", 1, []string{"hello\x01"}, nil},
	{"foo\nbar", 0, []string{"foo\nbar"}, io.EOF},
	{"alpha\nbeta\ngamma\n", '\n', []string{"alpha\n", "beta\n", "gamma\n"}, nil},
	{"alpha\nbeta\ngamma", '\n', []string{"alpha\n", "beta\n", "gamma"}, io.EOF},
}

func TestReadBytes(t *testing.T) {
	for _, test := range readBytesTests {
		buf := NewBufferString(test.buffer)
		var err error
		for _, expected := range test.expected {
			var bytes []byte
			bytes, err = buf.ReadBytes(test.delim)
			if string(bytes) != expected {
				t.Errorf("expected %q, got %q", expected, bytes)
			}
			if err != nil {
				break
			}
		}
		if err != test.err {
			t.Errorf("expected error %v, got %v", test.err, err)
		}
	}
}

func TestReadString(t *testing.T) {
	for _, test := range readBytesTests {
		buf := NewBufferString(test.buffer)
		var err error
		for _, expected := range test.expected {
			var s string
			s, err = buf.ReadString(test.delim)
			if s != expected {
				t.Errorf("expected %q, got %q", expected, s)
			}
			if err != nil {
				break
			}
		}
		if err != test.err {
			t.Errorf("expected error %v, got %v", test.err, err)
		}
	}
}

func BenchmarkReadString(b *testing.B) {
	const n = 32 << 10

	data := make([]byte, n)
	data[n-1] = 'x'
	b.SetBytes(int64(n))
	for i := 0; i < b.N; i++ {
		buf := NewBuffer(data)
		_, err := buf.ReadString('x')
		if err != nil {
			b.Fatal(err)
		}
	}
}

func TestGrow(t *testing.T) {
	x := []byte{'x'}
	y := []byte{'y'}
	tmp := make([]byte, 72)
	for _, growLen := range []int{0, 100, 1000, 10000, 100000} {
		for _, startLen := range []int{0, 100, 1000, 10000, 100000} {
			xBytes := Repeat(x, startLen)

			buf := NewBuffer(xBytes)
			// 如果我们进行读取操作，这会影响到buf.off的值，这是需要测试的一个重点。. md5:de673857faeb3086
			readBytes, _ := buf.Read(tmp)
			yBytes := Repeat(y, growLen)
			allocs := testing.AllocsPerRun(100, func() {
				buf.Grow(growLen)
				buf.Write(yBytes)
			})
			// 检查在单线程情况下，write操作中没有发生内存分配。. md5:345a0c0930748e59
			if allocs != 0 {
				t.Errorf("allocation occurred during write")
			}
			// 检查缓冲区是否有正确的数据。. md5:7fafa1ff34cff099
			if !Equal(buf.Bytes()[0:startLen-readBytes], xBytes[readBytes:]) {
				t.Errorf("bad initial data at %d %d", startLen, growLen)
			}
			if !Equal(buf.Bytes()[startLen-readBytes:startLen-readBytes+growLen], yBytes) {
				t.Errorf("bad written data at %d %d", startLen, growLen)
			}
		}
	}
}

func TestGrowOverflow(t *testing.T) {
	defer func() {
		if err := recover(); err != ErrTooLarge {
			t.Errorf("after too-large Grow, recover() = %v; want %v", err, ErrTooLarge)
		}
	}()

	buf := NewBuffer(make([]byte, 1))
	const maxInt = int(^uint(0) >> 1)
	buf.Grow(maxInt)
}

// 这是一个bug：在读取空切片时，它以前会返回EOF错误。. md5:b7f29e023c3d64c0
func TestReadEmptyAtEOF(t *testing.T) {
	b := new(Buffer)
	slice := make([]byte, 0)
	n, err := b.Read(slice)
	if err != nil {
		t.Errorf("read error: %v", err)
	}
	if n != 0 {
		t.Errorf("wrong count; got %d want 0", n)
	}
}

func TestUnreadByte(t *testing.T) {
	b := new(Buffer)

	// check at EOF
	if err := b.UnreadByte(); err == nil {
		t.Fatal("UnreadByte at EOF: got no error")
	}
	if _, err := b.ReadByte(); err == nil {
		t.Fatal("ReadByte at EOF: got no error")
	}
	if err := b.UnreadByte(); err == nil {
		t.Fatal("UnreadByte after ReadByte at EOF: got no error")
	}

	// check not at EOF
	b.WriteString("abcdefghijklmnopqrstuvwxyz")

	// 在读取失败后. md5:0ffe0e6d68de7072
	if n, err := b.Read(nil); n != 0 || err != nil {
		t.Fatalf("Read(nil) = %d,%v; want 0,nil", n, err)
	}
	if err := b.UnreadByte(); err == nil {
		t.Fatal("UnreadByte after Read(nil): got no error")
	}

	// after successful read
	if _, err := b.ReadBytes('m'); err != nil {
		t.Fatalf("ReadBytes: %v", err)
	}
	if err := b.UnreadByte(); err != nil {
		t.Fatalf("UnreadByte: %v", err)
	}
	c, err := b.ReadByte()
	if err != nil {
		t.Fatalf("ReadByte: %v", err)
	}
	if c != 'm' {
		t.Errorf("ReadByte = %q; want %q", c, 'm')
	}
}

// 测试我们偶尔进行的压缩。问题5154。. md5:4ab768dd436e6b2d
func TestBufferGrowth(t *testing.T) {
	var b Buffer
	buf := make([]byte, 1024)
	b.Write(buf[0:1])
	var cap0 int
	for i := 0; i < 5<<10; i++ {
		b.Write(buf)
		b.Read(buf)
		if i == 0 {
			cap0 = b.Cap()
		}
	}
	cap1 := b.Cap()
	//（*Buffer）.grow允许在滑动之前有2倍容量的冗余，所以我们将错误阈值设置为3倍。
	// md5:15593711c654c6c2
	if cap1 > cap0*3 {
		t.Errorf("buffer cap = %d; too big (grew from %d)", cap1, cap0)
	}
}

func BenchmarkWriteByte(b *testing.B) {
	const n = 4 << 10
	b.SetBytes(n)
	buf := NewBuffer(make([]byte, n))
	for i := 0; i < b.N; i++ {
		buf.Reset()
		for i := 0; i < n; i++ {
			buf.WriteByte('x')
		}
	}
}

func BenchmarkWriteRune(b *testing.B) {
	const n = 4 << 10
	const r = '☺'
	b.SetBytes(int64(n * utf8.RuneLen(r)))
	buf := NewBuffer(make([]byte, n*utf8.UTFMax))
	for i := 0; i < b.N; i++ {
		buf.Reset()
		for i := 0; i < n; i++ {
			buf.WriteRune(r)
		}
	}
}

// From Issue 5154.
func BenchmarkBufferNotEmptyWriteRead(b *testing.B) {
	buf := make([]byte, 1024)
	for i := 0; i < b.N; i++ {
		var b Buffer
		b.Write(buf[0:1])
		for i := 0; i < 5<<10; i++ {
			b.Write(buf)
			b.Read(buf)
		}
	}
}

// 检查我们不要过于频繁地进行紧凑。参考问题5154。. md5:094d39788b27155d
func BenchmarkBufferFullSmallReads(b *testing.B) {
	buf := make([]byte, 1024)
	for i := 0; i < b.N; i++ {
		var b Buffer
		b.Write(buf)
		for b.Len()+20 < b.Cap() {
			b.Write(buf[:10])
		}
		for i := 0; i < 5<<10; i++ {
			b.Read(buf[:1])
			b.Write(buf[:1])
		}
	}
}

func BenchmarkBufferWriteBlock(b *testing.B) {
	block := make([]byte, 1024)
	for _, n := range []int{1 << 12, 1 << 16, 1 << 20} {
		b.Run(fmt.Sprintf("N%d", n), func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				var bb Buffer
				for bb.Len() < n {
					bb.Write(block)
				}
			}
		})
	}
}

func BenchmarkBufferAppendNoCopy(b *testing.B) {
	var bb Buffer
	bb.Grow(16 << 20)
	b.SetBytes(int64(bb.Available()))
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		bb.Reset()
		b := bb.AvailableBuffer()
		b = b[:cap(b)] // 使用最大容量来模拟一个大的追加操作. md5:40dfb5f3c522b160
		bb.Write(b)    // 应该几乎是无限快速的. md5:1fd3483bfdabd715
	}
}
