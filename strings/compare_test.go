// 版权所有 2013 The Go Authors。保留所有权利。
// 使用本源代码受 BSD 风格许可协议约束，
// 该协议可在 LICENSE 文件中找到。

package strings_test

// 源自 bytes/compare_test.go。
// 由于底层实现完全相同，故省略基准测试。

import (
	. "github.com/888go/gosdk/strings"

	"testing"
)

var compareTests = []struct {
	a, b string
	i    int
}{
	{"", "", 0},
	{"a", "", 1},
	{"", "a", -1},
	{"abc", "abc", 0},
	{"ab", "abc", -1},
	{"abc", "ab", 1},
	{"x", "ab", 1},
	{"ab", "x", -1},
	{"x", "a", 1},
	{"b", "x", -1},
	// test runtime·memeq's chunked implementation
	{"abcdefgh", "abcdefgh", 0},
	{"abcdefghi", "abcdefghi", 0},
	{"abcdefghi", "abcdefghj", -1},
}

func TestCompare(t *testing.T) {
	for _, tt := range compareTests {
		cmp := Compare(tt.a, tt.b)
		if cmp != tt.i {
			t.Errorf(`Compare(%q, %q) = %v`, tt.a, tt.b, cmp)
		}
	}
}

func TestCompareIdenticalString(t *testing.T) {
	var s = "Hello Gophers!"
	if Compare(s, s) != 0 {
		t.Error("s != s")
	}
	if Compare(s, s[:1]) != 1 {
		t.Error("s > s[:1] failed")
	}
}

//
//func TestCompareStrings(t *testing.T) {
//	// unsafeString converts a []byte to a string with no allocation.
//	// The caller must not modify b while the result string is in use.
//	unsafeString := func(b []byte) string {
//		return unsafe.String(unsafe.SliceData(b), len(b))
//	}
//
//	lengths := make([]int, 0) // 以升序排列的长度值进行测试
//	for i := 0; i <= 128; i++ {
//		lengths = append(lengths, i)
//	}
//	lengths = append(lengths, 256, 512, 1024, 1333, 4095, 4096, 4097)
//
//	if !testing.Short() || testenv.Builder() != "" {
//		lengths = append(lengths, 65535, 65536, 65537, 99999)
//	}
//
//	n := lengths[len(lengths)-1]
//	a := make([]byte, n+1)
//	b := make([]byte, n+1)
//	lastLen := 0
//	for _, len := range lengths {
//		// 随机但确定性的数据。不含0或255。
//		for i := 0; i < len; i++ {
//			a[i] = byte(1 + 31*i%254)
//			b[i] = byte(1 + 31*i%254)
//		}
//		// 结尾后的数据不同
//		for i := len; i <= n; i++ {
//			a[i] = 8
//			b[i] = 9
//		}
//
//		sa, sb := unsafeString(a), unsafeString(b)
//		cmp := Compare(sa[:len], sb[:len])
//		if cmp != 0 {
//			t.Errorf(`CompareIdentical(%d) = %d`, len, cmp)
//		}
//		if len > 0 {
//			cmp = Compare(sa[:len-1], sb[:len])
//			if cmp != -1 {
//				t.Errorf(`CompareAshorter(%d) = %d`, len, cmp)
//			}
//			cmp = Compare(sa[:len], sb[:len-1])
//			if cmp != 1 {
//				t.Errorf(`CompareBshorter(%d) = %d`, len, cmp)
//			}
//		}
//		for k := lastLen; k < len; k++ {
//			b[k] = a[k] - 1
//			cmp = Compare(unsafeString(a[:len]), unsafeString(b[:len]))
//			if cmp != 1 {
//				t.Errorf(`CompareAbigger(%d,%d) = %d`, len, k, cmp)
//			}
//			b[k] = a[k] + 1
//			cmp = Compare(unsafeString(a[:len]), unsafeString(b[:len]))
//			if cmp != -1 {
//				t.Errorf(`CompareBbigger(%d,%d) = %d`, len, k, cmp)
//			}
//			b[k] = a[k]
//		}
//		lastLen = len
//	}
//}
