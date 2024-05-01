//版权所有2009年Go作者。所有权利保留。
//使用此源代码受BSD风格
//可以在LICENSE文件中找到的许可证。
// md5:2e9dc81828a3be8a

// 为测试导出对strconv内部的访问. md5:a0ff8c6c86975e9b

package strconv

func NewDecimal(i uint64) *decimal {
	d := new(decimal)
	d.Assign(i)
	return d
}

func SetOptimize(b bool) bool {
	old := optimize
	optimize = b
	return old
}

//func ParseFloatPrefix(s string, bitSize int) (float64, int, error) {
//	return parseFloatPrefix(s, bitSize)
//}

func MulByLog2Log10(x int) int {
	return mulByLog2Log10(x)
}

func MulByLog10Log2(x int) int {
	return mulByLog10Log2(x)
}

//func parseFloatPrefix(s string, bitSize int) (float64, int, error) {
//	if bitSize == 32 {
//		f, n, err := atof32(s)
//		return float64(f), n, err
//	}
//	return atof64(s)
//}

// mulByLog2Log10 returns math.Floor(x * log(2)/log(10)) for an integer x in
// the range -1600 <= x && x <= +1600.
//
// The range restriction lets us work in faster integer arithmetic instead of
// slower floating point arithmetic. Correctness is verified by unit tests.
func mulByLog2Log10(x int) int {
	// log(2)/log(10) ≈ 0.30102999566 ≈ 78913 / 2^18
	return (x * 78913) >> 18
}

// mulByLog10Log2 returns math.Floor(x * log(10)/log(2)) for an integer x in
// the range -500 <= x && x <= +500.
//
// The range restriction lets us work in faster integer arithmetic instead of
// slower floating point arithmetic. Correctness is verified by unit tests.
func mulByLog10Log2(x int) int {
	// log(10)/log(2) ≈ 3.32192809489 ≈ 108853 / 2^15
	return (x * 108853) >> 15
}
