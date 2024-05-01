// 版权所有 2020 Go 作者。保留所有权利。
// 使用此源代码受BSD风格许可证管辖，该许可证可从LICENSE文件中找到。
// md5:97282fec9c34c48c

package strconv_test

import (
	. "github.com/888go/gosdk/strconv"
	"testing"
)

func TestFormatComplex(t *testing.T) {
	tests := []struct {
		c       complex128
		fmt     byte
		prec    int
		bitSize int
		out     string
	}{
		// a variety of signs
		{1 + 2i, 'g', -1, 128, "(1+2i)"},
		{3 - 4i, 'g', -1, 128, "(3-4i)"},
		{-5 + 6i, 'g', -1, 128, "(-5+6i)"},
		{-7 - 8i, 'g', -1, 128, "(-7-8i)"},

		// 测试 fmt 和 prec 是否正常工作. md5:6a91165a522d9b22
		{3.14159 + 0.00123i, 'e', 3, 128, "(3.142e+00+1.230e-03i)"},
		{3.14159 + 0.00123i, 'f', 3, 128, "(3.142+0.001i)"},
		{3.14159 + 0.00123i, 'g', 3, 128, "(3.14+0.00123i)"},

		// 确保位大小的舍入工作正常. md5:f946a16c80ad0ba1
		{1.2345678901234567 + 9.876543210987654i, 'f', -1, 128, "(1.2345678901234567+9.876543210987654i)"},
		{1.2345678901234567 + 9.876543210987654i, 'f', -1, 64, "(1.2345679+9.876543i)"},

		// 其他情况由FormatFloat测试处理. md5:36416d5e6366147c
	}
	for _, test := range tests {
		out := FormatComplex(test.c, test.fmt, test.prec, test.bitSize)
		if out != test.out {
			t.Fatalf("FormatComplex(%v, %q, %d, %d) = %q; want %q",
				test.c, test.fmt, test.prec, test.bitSize, out, test.out)
		}
	}
}

func TestFormatComplexInvalidBitSize(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("expected panic due to invalid bitSize")
		}
	}()
	_ = FormatComplex(1+2i, 'g', -1, 100)
}
