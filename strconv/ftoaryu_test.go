// 版权所有 2021 Go 作者。保留所有权利。
// 本源代码的使用受BSD风格许可证约束，该许可证可在LICENSE文件中找到。
// md5:3fcaafd8cf5d6a63

package strconv_test

import (
	. "github.com/888go/gosdk/strconv"
	"math"
	"testing"
)

func TestMulByLog2Log10(t *testing.T) {
	for x := -1600; x <= +1600; x++ {
		iMath := MulByLog2Log10(x)
		fMath := int(math.Floor(float64(x) * math.Ln2 / math.Ln10))
		if iMath != fMath {
			t.Errorf("mulByLog2Log10(%d) failed: %d vs %d\n", x, iMath, fMath)
		}
	}
}

func TestMulByLog10Log2(t *testing.T) {
	for x := -500; x <= +500; x++ {
		iMath := MulByLog10Log2(x)
		fMath := int(math.Floor(float64(x) * math.Ln10 / math.Ln2))
		if iMath != fMath {
			t.Errorf("mulByLog10Log2(%d) failed: %d vs %d\n", x, iMath, fMath)
		}
	}
}
