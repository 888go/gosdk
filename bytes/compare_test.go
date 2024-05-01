// 版权所有 2013 Go 作者。保留所有权利。
// 使用此源代码受BSD风格许可证约束，该许可证可从LICENSE文件中找到。
// md5:19d1a3ed91182ee4

package bytes_test

import (
	"fmt"
	. "github.com/888go/gosdk/bytes"
	"testing"
)

var compareTests = []struct {
	a, b []byte
	i    int
}{
	{[]byte(""), []byte(""), 0},
	{[]byte("a"), []byte(""), 1},
	{[]byte(""), []byte("a"), -1},
	{[]byte("abc"), []byte("abc"), 0},
	{[]byte("abd"), []byte("abc"), 1},
	{[]byte("abc"), []byte("abd"), -1},
	{[]byte("ab"), []byte("abc"), -1},
	{[]byte("abc"), []byte("ab"), 1},
	{[]byte("x"), []byte("ab"), 1},
	{[]byte("ab"), []byte("x"), -1},
	{[]byte("x"), []byte("a"), 1},
	{[]byte("b"), []byte("x"), -1},
	// test runtime·memeq's chunked implementation
	{[]byte("abcdefgh"), []byte("abcdefgh"), 0},
	{[]byte("abcdefghi"), []byte("abcdefghi"), 0},
	{[]byte("abcdefghi"), []byte("abcdefghj"), -1},
	{[]byte("abcdefghj"), []byte("abcdefghi"), 1},
	// nil tests
	{nil, nil, 0},
	{[]byte(""), nil, 0},
	{nil, []byte(""), 0},
	{[]byte("a"), nil, 1},
	{nil, []byte("a"), -1},
}

func TestCompare(t *testing.T) {
	for _, tt := range compareTests {
		numShifts := 16
		buffer := make([]byte, len(tt.b)+numShifts)
		// 改变输入对齐方式为 tt.b. md5:ce1433df219092cd
		for offset := 0; offset <= numShifts; offset++ {
			shiftedB := buffer[offset : len(tt.b)+offset]
			copy(shiftedB, tt.b)
			cmp := Compare(tt.a, shiftedB)
			if cmp != tt.i {
				t.Errorf(`Compare(%q, %q), offset %d = %v; want %v`, tt.a, tt.b, offset, cmp, tt.i)
			}
		}
	}
}

func TestCompareIdenticalSlice(t *testing.T) {
	var b = []byte("Hello Gophers!")
	if Compare(b, b) != 0 {
		t.Error("b != b")
	}
	if Compare(b, b[:1]) != 1 {
		t.Error("b > b[:1] failed")
	}
}

func TestCompareBytes(t *testing.T) {
	lengths := make([]int, 0) // 以递增顺序进行测试的长度. md5:d422e60df7501431
	for i := 0; i <= 128; i++ {
		lengths = append(lengths, i)
	}
	lengths = append(lengths, 256, 512, 1024, 1333, 4095, 4096, 4097)

	if !testing.Short() {
		lengths = append(lengths, 65535, 65536, 65537, 99999)
	}

	n := lengths[len(lengths)-1]
	a := make([]byte, n+1)
	b := make([]byte, n+1)
	for _, len := range lengths {
		// 随机但确定的数据。不包含0或255。. md5:60ed4a09009d8428
		for i := 0; i < len; i++ {
			a[i] = byte(1 + 31*i%254)
			b[i] = byte(1 + 31*i%254)
		}
		// 数据末尾之后的内容是不同的. md5:8d87dd9025a7c19f
		for i := len; i <= n; i++ {
			a[i] = 8
			b[i] = 9
		}
		cmp := Compare(a[:len], b[:len])
		if cmp != 0 {
			t.Errorf(`CompareIdentical(%d) = %d`, len, cmp)
		}
		if len > 0 {
			cmp = Compare(a[:len-1], b[:len])
			if cmp != -1 {
				t.Errorf(`CompareAshorter(%d) = %d`, len, cmp)
			}
			cmp = Compare(a[:len], b[:len-1])
			if cmp != 1 {
				t.Errorf(`CompareBshorter(%d) = %d`, len, cmp)
			}
		}
		for k := 0; k < len; k++ {
			b[k] = a[k] - 1
			cmp = Compare(a[:len], b[:len])
			if cmp != 1 {
				t.Errorf(`CompareAbigger(%d,%d) = %d`, len, k, cmp)
			}
			b[k] = a[k] + 1
			cmp = Compare(a[:len], b[:len])
			if cmp != -1 {
				t.Errorf(`CompareBbigger(%d,%d) = %d`, len, k, cmp)
			}
			b[k] = a[k]
		}
	}
}

func TestEndianBaseCompare(t *testing.T) {
// 该测试比较两个几乎相同的字节切片，唯一的区别在于某个索引j处，a[j]大于b[j]并且a[j+1]小于b[j+1]。
// 如果实现方式错误地以错误的字节序比较了大块数据，它将得到错误的结果。
// 目前没有任何向量寄存器大于512字节。
// md5:4bc7ed687468d65e
	const maxLength = 512
	a := make([]byte, maxLength)
	b := make([]byte, maxLength)
	// 随机但确定的数据。不包含0或255。. md5:60ed4a09009d8428
	for i := 0; i < maxLength; i++ {
		a[i] = byte(1 + 31*i%254)
		b[i] = byte(1 + 31*i%254)
	}
	for i := 2; i <= maxLength; i <<= 1 {
		for j := 0; j < i-1; j++ {
			a[j] = b[j] - 1
			a[j+1] = b[j+1] + 1
			cmp := Compare(a[:i], b[:i])
			if cmp != -1 {
				t.Errorf(`CompareBbigger(%d,%d) = %d`, i, j, cmp)
			}
			a[j] = b[j] + 1
			a[j+1] = b[j+1] - 1
			cmp = Compare(a[:i], b[:i])
			if cmp != 1 {
				t.Errorf(`CompareAbigger(%d,%d) = %d`, i, j, cmp)
			}
			a[j] = b[j]
			a[j+1] = b[j+1]
		}
	}
}

func BenchmarkCompareBytesEqual(b *testing.B) {
	b1 := []byte("Hello Gophers!")
	b2 := []byte("Hello Gophers!")
	for i := 0; i < b.N; i++ {
		if Compare(b1, b2) != 0 {
			b.Fatal("b1 != b2")
		}
	}
}

func BenchmarkCompareBytesToNil(b *testing.B) {
	b1 := []byte("Hello Gophers!")
	var b2 []byte
	for i := 0; i < b.N; i++ {
		if Compare(b1, b2) != 1 {
			b.Fatal("b1 > b2 failed")
		}
	}
}

func BenchmarkCompareBytesEmpty(b *testing.B) {
	b1 := []byte("")
	b2 := b1
	for i := 0; i < b.N; i++ {
		if Compare(b1, b2) != 0 {
			b.Fatal("b1 != b2")
		}
	}
}

func BenchmarkCompareBytesIdentical(b *testing.B) {
	b1 := []byte("Hello Gophers!")
	b2 := b1
	for i := 0; i < b.N; i++ {
		if Compare(b1, b2) != 0 {
			b.Fatal("b1 != b2")
		}
	}
}

func BenchmarkCompareBytesSameLength(b *testing.B) {
	b1 := []byte("Hello Gophers!")
	b2 := []byte("Hello, Gophers")
	for i := 0; i < b.N; i++ {
		if Compare(b1, b2) != -1 {
			b.Fatal("b1 < b2 failed")
		}
	}
}

func BenchmarkCompareBytesDifferentLength(b *testing.B) {
	b1 := []byte("Hello Gophers!")
	b2 := []byte("Hello, Gophers!")
	for i := 0; i < b.N; i++ {
		if Compare(b1, b2) != -1 {
			b.Fatal("b1 < b2 failed")
		}
	}
}

func benchmarkCompareBytesBigUnaligned(b *testing.B, offset int) {
	b.StopTimer()
	b1 := make([]byte, 0, 1<<20)
	for len(b1) < 1<<20 {
		b1 = append(b1, "Hello Gophers!"...)
	}
	b2 := append([]byte("12345678")[:offset], b1...)
	b.StartTimer()
	for j := 0; j < b.N; j++ {
		if Compare(b1, b2[offset:]) != 0 {
			b.Fatal("b1 != b2")
		}
	}
	b.SetBytes(int64(len(b1)))
}

func BenchmarkCompareBytesBigUnaligned(b *testing.B) {
	for i := 1; i < 8; i++ {
		b.Run(fmt.Sprintf("offset=%d", i), func(b *testing.B) {
			benchmarkCompareBytesBigUnaligned(b, i)
		})
	}
}

func benchmarkCompareBytesBigBothUnaligned(b *testing.B, offset int) {
	b.StopTimer()
	pattern := []byte("Hello Gophers!")
	b1 := make([]byte, 0, 1<<20+len(pattern))
	for len(b1) < 1<<20 {
		b1 = append(b1, pattern...)
	}
	b2 := make([]byte, len(b1))
	copy(b2, b1)
	b.StartTimer()
	for j := 0; j < b.N; j++ {
		if Compare(b1[offset:], b2[offset:]) != 0 {
			b.Fatal("b1 != b2")
		}
	}
	b.SetBytes(int64(len(b1[offset:])))
}

func BenchmarkCompareBytesBigBothUnaligned(b *testing.B) {
	for i := 0; i < 8; i++ {
		b.Run(fmt.Sprintf("offset=%d", i), func(b *testing.B) {
			benchmarkCompareBytesBigBothUnaligned(b, i)
		})
	}
}

func BenchmarkCompareBytesBig(b *testing.B) {
	b.StopTimer()
	b1 := make([]byte, 0, 1<<20)
	for len(b1) < 1<<20 {
		b1 = append(b1, "Hello Gophers!"...)
	}
	b2 := append([]byte{}, b1...)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		if Compare(b1, b2) != 0 {
			b.Fatal("b1 != b2")
		}
	}
	b.SetBytes(int64(len(b1)))
}

func BenchmarkCompareBytesBigIdentical(b *testing.B) {
	b.StopTimer()
	b1 := make([]byte, 0, 1<<20)
	for len(b1) < 1<<20 {
		b1 = append(b1, "Hello Gophers!"...)
	}
	b2 := b1
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		if Compare(b1, b2) != 0 {
			b.Fatal("b1 != b2")
		}
	}
	b.SetBytes(int64(len(b1)))
}
