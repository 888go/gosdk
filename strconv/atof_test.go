//版权所有2009年Go作者。所有权利保留。
//使用此源代码受BSD风格
//可以在LICENSE文件中找到的许可证。
// md5:2e9dc81828a3be8a

package strconv_test

import (
	. "github.com/888go/gosdk/strconv"
	"math"
	"math/rand"
	"reflect"
	strconv2 "strconv"
	"strings"
	"sync"
	"testing"
)

type atofTest struct {
	in  string
	out string
	err error
}

var atoftests = []atofTest{
	{"", "0", ErrSyntax},
	{"1", "1", nil},
	{"+1", "1", nil},
	{"1x", "0", ErrSyntax},
	{"1.1.", "0", ErrSyntax},
	{"1e23", "1e+23", nil},
	{"1E23", "1e+23", nil},
	{"100000000000000000000000", "1e+23", nil},
	{"1e-100", "1e-100", nil},
	{"123456700", "1.234567e+08", nil},
	{"99999999999999974834176", "9.999999999999997e+22", nil},
	{"100000000000000000000001", "1.0000000000000001e+23", nil},
	{"100000000000000008388608", "1.0000000000000001e+23", nil},
	{"100000000000000016777215", "1.0000000000000001e+23", nil},
	{"100000000000000016777216", "1.0000000000000003e+23", nil},
	{"-1", "-1", nil},
	{"-0.1", "-0.1", nil},
	{"-0", "-0", nil},
	{"1e-20", "1e-20", nil},
	{"625e-3", "0.625", nil},

	// 十六进制浮点数。. md5:de2d821e70967e8f
	{"0x1p0", "1", nil},
	{"0x1p1", "2", nil},
	{"0x1p-1", "0.5", nil},
	{"0x1ep-1", "15", nil},
	{"-0x1ep-1", "-15", nil},
	{"-0x1_ep-1", "-15", nil},
	{"0x1p-200", "6.223015277861142e-61", nil},
	{"0x1p200", "1.6069380442589903e+60", nil},
	{"0x1fFe2.p0", "131042", nil},
	{"0x1fFe2.P0", "131042", nil},
	{"-0x2p3", "-16", nil},
	{"0x0.fp4", "15", nil},
	{"0x0.fp0", "0.9375", nil},
	{"0x1e2", "0", ErrSyntax},
	{"1p2", "0", ErrSyntax},

	// zeros
	{"0", "0", nil},
	{"0e0", "0", nil},
	{"-0e0", "-0", nil},
	{"+0e0", "0", nil},
	{"0e-0", "0", nil},
	{"-0e-0", "-0", nil},
	{"+0e-0", "0", nil},
	{"0e+0", "0", nil},
	{"-0e+0", "-0", nil},
	{"+0e+0", "0", nil},
	{"0e+01234567890123456789", "0", nil},
	{"0.00e-01234567890123456789", "0", nil},
	{"-0e+01234567890123456789", "-0", nil},
	{"-0.00e-01234567890123456789", "-0", nil},
	{"0x0p+01234567890123456789", "0", nil},
	{"0x0.00p-01234567890123456789", "0", nil},
	{"-0x0p+01234567890123456789", "-0", nil},
	{"-0x0.00p-01234567890123456789", "-0", nil},

	{"0e291", "0", nil}, // issue 15364
	{"0e292", "0", nil}, // issue 15364
	{"0e347", "0", nil}, // issue 15364
	{"0e348", "0", nil}, // issue 15364
	{"-0e291", "-0", nil},
	{"-0e292", "-0", nil},
	{"-0e347", "-0", nil},
	{"-0e348", "-0", nil},
	{"0x0p126", "0", nil},
	{"0x0p127", "0", nil},
	{"0x0p128", "0", nil},
	{"0x0p129", "0", nil},
	{"0x0p130", "0", nil},
	{"0x0p1022", "0", nil},
	{"0x0p1023", "0", nil},
	{"0x0p1024", "0", nil},
	{"0x0p1025", "0", nil},
	{"0x0p1026", "0", nil},
	{"-0x0p126", "-0", nil},
	{"-0x0p127", "-0", nil},
	{"-0x0p128", "-0", nil},
	{"-0x0p129", "-0", nil},
	{"-0x0p130", "-0", nil},
	{"-0x0p1022", "-0", nil},
	{"-0x0p1023", "-0", nil},
	{"-0x0p1024", "-0", nil},
	{"-0x0p1025", "-0", nil},
	{"-0x0p1026", "-0", nil},

	// NaNs
	{"nan", "NaN", nil},
	{"NaN", "NaN", nil},
	{"NAN", "NaN", nil},

	// Infs
	{"inf", "+Inf", nil},
	{"-Inf", "-Inf", nil},
	{"+INF", "+Inf", nil},
	{"-Infinity", "-Inf", nil},
	{"+INFINITY", "+Inf", nil},
	{"Infinity", "+Inf", nil},

	// largest float64
	{"1.7976931348623157e308", "1.7976931348623157e+308", nil},
	{"-1.7976931348623157e308", "-1.7976931348623157e+308", nil},
	{"0x1.fffffffffffffp1023", "1.7976931348623157e+308", nil},
	{"-0x1.fffffffffffffp1023", "-1.7976931348623157e+308", nil},
	{"0x1fffffffffffffp+971", "1.7976931348623157e+308", nil},
	{"-0x1fffffffffffffp+971", "-1.7976931348623157e+308", nil},
	{"0x.1fffffffffffffp1027", "1.7976931348623157e+308", nil},
	{"-0x.1fffffffffffffp1027", "-1.7976931348623157e+308", nil},

	// next float64 - 太大了. md5:e595d550db35fd95
	{"1.7976931348623159e308", "+Inf", ErrRange},
	{"-1.7976931348623159e308", "-Inf", ErrRange},
	{"0x1p1024", "+Inf", ErrRange},
	{"-0x1p1024", "-Inf", ErrRange},
	{"0x2p1023", "+Inf", ErrRange},
	{"-0x2p1023", "-Inf", ErrRange},
	{"0x.1p1028", "+Inf", ErrRange},
	{"-0x.1p1028", "-Inf", ErrRange},
	{"0x.2p1027", "+Inf", ErrRange},
	{"-0x.2p1027", "-Inf", ErrRange},

// 边界是...158079
// 边缘 - 可接受
// md5:283760bcac5d4566
	{"1.7976931348623158e308", "1.7976931348623157e+308", nil},
	{"-1.7976931348623158e308", "-1.7976931348623157e+308", nil},
	{"0x1.fffffffffffff7fffp1023", "1.7976931348623157e+308", nil},
	{"-0x1.fffffffffffff7fffp1023", "-1.7976931348623157e+308", nil},
	// borderline - too large
	{"1.797693134862315808e308", "+Inf", ErrRange},
	{"-1.797693134862315808e308", "-Inf", ErrRange},
	{"0x1.fffffffffffff8p1023", "+Inf", ErrRange},
	{"-0x1.fffffffffffff8p1023", "-Inf", ErrRange},
	{"0x1fffffffffffff.8p+971", "+Inf", ErrRange},
	{"-0x1fffffffffffff8p+967", "-Inf", ErrRange},
	{"0x.1fffffffffffff8p1027", "+Inf", ErrRange},
	{"-0x.1fffffffffffff9p1027", "-Inf", ErrRange},

	// a little too large
	{"1e308", "1e+308", nil},
	{"2e308", "+Inf", ErrRange},
	{"1e309", "+Inf", ErrRange},
	{"0x1p1025", "+Inf", ErrRange},

	// way too large
	{"1e310", "+Inf", ErrRange},
	{"-1e310", "-Inf", ErrRange},
	{"1e400", "+Inf", ErrRange},
	{"-1e400", "-Inf", ErrRange},
	{"1e400000", "+Inf", ErrRange},
	{"-1e400000", "-Inf", ErrRange},
	{"0x1p1030", "+Inf", ErrRange},
	{"0x1p2000", "+Inf", ErrRange},
	{"0x1p2000000000", "+Inf", ErrRange},
	{"-0x1p1030", "-Inf", ErrRange},
	{"-0x1p2000", "-Inf", ErrRange},
	{"-0x1p2000000000", "-Inf", ErrRange},

	// denormalized
	{"1e-305", "1e-305", nil},
	{"1e-306", "1e-306", nil},
	{"1e-307", "1e-307", nil},
	{"1e-308", "1e-308", nil},
	{"1e-309", "1e-309", nil},
	{"1e-310", "1e-310", nil},
	{"1e-322", "1e-322", nil},
	// smallest denormal
	{"5e-324", "5e-324", nil},
	{"4e-324", "5e-324", nil},
	{"3e-324", "5e-324", nil},
	// too small
	{"2e-324", "0", nil},
	// way too small
	{"1e-350", "0", nil},
	{"1e-400000", "0", nil},

	// 接近于denormals和denormals。. md5:5213e06b56063627
	{"0x2.00000000000000p-1010", "1.8227805048890994e-304", nil}, // 0x00e0000000000000
	{"0x1.fffffffffffff0p-1010", "1.8227805048890992e-304", nil}, // 0x00dfffffffffffff
	{"0x1.fffffffffffff7p-1010", "1.8227805048890992e-304", nil}, // rounded down
	{"0x1.fffffffffffff8p-1010", "1.8227805048890994e-304", nil}, // rounded up
	{"0x1.fffffffffffff9p-1010", "1.8227805048890994e-304", nil}, // rounded up

	{"0x2.00000000000000p-1022", "4.450147717014403e-308", nil},  // 0x0020000000000000
	{"0x1.fffffffffffff0p-1022", "4.4501477170144023e-308", nil}, // 0x001fffffffffffff
	{"0x1.fffffffffffff7p-1022", "4.4501477170144023e-308", nil}, // rounded down
	{"0x1.fffffffffffff8p-1022", "4.450147717014403e-308", nil},  // rounded up
	{"0x1.fffffffffffff9p-1022", "4.450147717014403e-308", nil},  // rounded up

	{"0x1.00000000000000p-1022", "2.2250738585072014e-308", nil}, // 0x0010000000000000
	{"0x0.fffffffffffff0p-1022", "2.225073858507201e-308", nil},  // 0x000fffffffffffff
	{"0x0.ffffffffffffe0p-1022", "2.2250738585072004e-308", nil}, // 0x000ffffffffffffe
	{"0x0.ffffffffffffe7p-1022", "2.2250738585072004e-308", nil}, // rounded down
	{"0x1.ffffffffffffe8p-1023", "2.225073858507201e-308", nil},  // rounded up
	{"0x1.ffffffffffffe9p-1023", "2.225073858507201e-308", nil},  // rounded up

	{"0x0.00000003fffff0p-1022", "2.072261e-317", nil},  // 0x00000000003fffff
	{"0x0.00000003456780p-1022", "1.694649e-317", nil},  // 0x0000000000345678
	{"0x0.00000003456787p-1022", "1.694649e-317", nil},  // rounded down
	{"0x0.00000003456788p-1022", "1.694649e-317", nil},  // 向下取整（偶数舍入）. md5:600817f8100c1677
	{"0x0.00000003456790p-1022", "1.6946496e-317", nil}, // 0x0000000000345679
	{"0x0.00000003456789p-1022", "1.6946496e-317", nil}, // rounded up

	{"0x0.0000000345678800000000000000000000000001p-1022", "1.6946496e-317", nil}, // rounded up

	{"0x0.000000000000f0p-1022", "7.4e-323", nil}, // 0x000000000000000f
	{"0x0.00000000000060p-1022", "3e-323", nil},   // 0x0000000000000006
	{"0x0.00000000000058p-1022", "3e-323", nil},   // rounded up
	{"0x0.00000000000057p-1022", "2.5e-323", nil}, // rounded down
	{"0x0.00000000000050p-1022", "2.5e-323", nil}, // 0x0000000000000005

	{"0x0.00000000000010p-1022", "5e-324", nil},  // 0x0000000000000001
	{"0x0.000000000000081p-1022", "5e-324", nil}, // rounded up
	{"0x0.00000000000008p-1022", "0", nil},       // rounded down
	{"0x0.00000000000007fp-1022", "0", nil},      // rounded down

	// 尝试使指数溢出. md5:9b2de22c589b27df
	{"1e-4294967296", "0", nil},
	{"1e+4294967296", "+Inf", ErrRange},
	{"1e-18446744073709551616", "0", nil},
	{"1e+18446744073709551616", "+Inf", ErrRange},
	{"0x1p-4294967296", "0", nil},
	{"0x1p+4294967296", "+Inf", ErrRange},
	{"0x1p-18446744073709551616", "0", nil},
	{"0x1p+18446744073709551616", "+Inf", ErrRange},

	// Parse errors
	{"1e", "0", ErrSyntax},
	{"1e-", "0", ErrSyntax},
	{".e-1", "0", ErrSyntax},
	{"1\x00.2", "0", ErrSyntax},
	{"0x", "0", ErrSyntax},
	{"0x.", "0", ErrSyntax},
	{"0x1", "0", ErrSyntax},
	{"0x.1", "0", ErrSyntax},
	{"0x1p", "0", ErrSyntax},
	{"0x.1p", "0", ErrSyntax},
	{"0x1p+", "0", ErrSyntax},
	{"0x.1p+", "0", ErrSyntax},
	{"0x1p-", "0", ErrSyntax},
	{"0x.1p-", "0", ErrSyntax},
	{"0x1p+2", "4", nil},
	{"0x.1p+2", "0.25", nil},
	{"0x1p-2", "0.25", nil},
	{"0x.1p-2", "0.015625", nil},

	// https://www.exploringbinary.com/java-hangs-when-converting-2-2250738585072012e-308/
// 这段Go注释的中文翻译是：
// 
// https://www.exploringbinary.com/java-hangs-when-converting-2-2250738585072012e-308/
// 这段Go注释的中文翻译是：
// 
// https://www.exploringbinary.com/java-hangs-when-converting-2-2250738585072012e-308/
// 这段Go注释的中文翻译是：
// 
// https://www.exploringbinary.com/java-hangs-when-converting-2-2250738585072012e-308/
// 这段Go注释的中文翻译是：
// 
// https://www.exploringbinary.com/java-hangs-when-converting-2-2250738585072012e-308/
// 这段Go注释的中文翻译是：
// 
// https://www.exploringbinary.com/java-hangs-when-converting-2-2250738585072012e-308/
// 这段Go注释的中文翻译是：
// 
// 当将2.2250738585072012e-308转换时，Java会挂起（或停止响应）：https://www.exploringbinary.com/java-hangs-when-converting-2-2250738585072012e-308/. md5:230c0c4f17858772
	{"2.2250738585072012e-308", "2.2250738585072014e-308", nil},
	// https://www.exploringbinary.com/php-hangs-on-numeric-value-2.2250738585072011e-308/ 
// 
// 这段Go注释翻译成中文是：
// 
// 这是一个链接，内容是关于PHP中数值2.2250738585072011e-308会导致程序挂起的问题（https://www.exploringbinary.com/php-hangs-on-numeric-value-2.2250738585072011e-308/）. md5:e6656740a836a448
	{"2.2250738585072011e-308", "2.225073858507201e-308", nil},

	// 一个非常大的数字（最初被快速算法错误解析）。. md5:aaec7675e0bff74a
	{"4.630813248087435e+307", "4.630813248087435e+307", nil},

	// 一种不同类型的非常大的数字。. md5:6522045cedd96e71
	{"22.222222222222222", "22.22222222222222", nil},
	{"2." + strings.Repeat("2", 4000) + "e+1", "22.22222222222222", nil},
	{"0x1.1111111111111p222", "7.18931911124017e+66", nil},
	{"0x2.2222222222222p221", "7.18931911124017e+66", nil},
	{"0x2." + strings.Repeat("2", 4000) + "p221", "7.18931911124017e+66", nil},

// 完整位于1和math.Nextafter(1, 2)之间的中间点。
// 四舍五入到偶数（向下取整）。
// md5:e79d24edcbd74c15
	{"1.00000000000000011102230246251565404236316680908203125", "1", nil},
	{"0x1.00000000000008p0", "1", nil},
	// 稍微低一点；仍然向下舍入。. md5:dca3ee1e60a729a5
	{"1.00000000000000011102230246251565404236316680908203124", "1", nil},
	{"0x1.00000000000007Fp0", "1", nil},
	// 略高一点；向上取整。. md5:afd45945f9caf31d
	{"1.00000000000000011102230246251565404236316680908203126", "1.0000000000000002", nil},
	{"0x1.000000000000081p0", "1.0000000000000002", nil},
	{"0x1.00000000000009p0", "1.0000000000000002", nil},
	// 稍微高一点，但你必须读到结尾才能明白。. md5:b8273dbe0862f4f3
	{"1.00000000000000011102230246251565404236316680908203125" + strings.Repeat("0", 10000) + "1", "1.0000000000000002", nil},
	{"0x1.00000000000008" + strings.Repeat("0", 10000) + "1p0", "1.0000000000000002", nil},

// 在 `x := math.Nextafter(1, 2)` 和 `math.Nextafter(x, 2)` 之间的中间点
// 向偶数方向四舍五入（向上）
// md5:3ba5d040c89bb711
	{"1.00000000000000033306690738754696212708950042724609375", "1.0000000000000004", nil},
	{"0x1.00000000000018p0", "1.0000000000000004", nil},

// 在1090544144181609278303144771584和1090544144181609419040633126912之间的中间值
// （15497564393479157p+46，应该四舍五入到最接近的偶数15497564393479156p+46，参见问题36657）
// md5:4b42731b30cfcc6c
	{"1090544144181609348671888949248", "1.0905441441816093e+30", nil},
	// 略高于，向上取整. md5:a7a383a63db55e7c
	{"1090544144181609348835077142190", "1.0905441441816094e+30", nil},

	// Underscores.
	{"1_23.50_0_0e+1_2", "1.235e+14", nil},
	{"-_123.5e+12", "0", ErrSyntax},
	{"+_123.5e+12", "0", ErrSyntax},
	{"_123.5e+12", "0", ErrSyntax},
	{"1__23.5e+12", "0", ErrSyntax},
	{"123_.5e+12", "0", ErrSyntax},
	{"123._5e+12", "0", ErrSyntax},
	{"123.5_e+12", "0", ErrSyntax},
	{"123.5__0e+12", "0", ErrSyntax},
	{"123.5e_+12", "0", ErrSyntax},
	{"123.5e+_12", "0", ErrSyntax},
	{"123.5e_-12", "0", ErrSyntax},
	{"123.5e-_12", "0", ErrSyntax},
	{"123.5e+1__2", "0", ErrSyntax},
	{"123.5e+12_", "0", ErrSyntax},

	{"0x_1_2.3_4_5p+1_2", "74565", nil},
	{"-_0x12.345p+12", "0", ErrSyntax},
	{"+_0x12.345p+12", "0", ErrSyntax},
	{"_0x12.345p+12", "0", ErrSyntax},
	{"0x__12.345p+12", "0", ErrSyntax},
	{"0x1__2.345p+12", "0", ErrSyntax},
	{"0x12_.345p+12", "0", ErrSyntax},
	{"0x12._345p+12", "0", ErrSyntax},
	{"0x12.3__45p+12", "0", ErrSyntax},
	{"0x12.345_p+12", "0", ErrSyntax},
	{"0x12.345p_+12", "0", ErrSyntax},
	{"0x12.345p+_12", "0", ErrSyntax},
	{"0x12.345p_-12", "0", ErrSyntax},
	{"0x12.345p-_12", "0", ErrSyntax},
	{"0x12.345p+1__2", "0", ErrSyntax},
	{"0x12.345p+12_", "0", ErrSyntax},

	{"1e100x", "0", ErrSyntax},
	{"1e1000x", "0", ErrSyntax},
}

var atof32tests = []atofTest{
	// Hex
	{"0x1p-100", "7.888609e-31", nil},
	{"0x1p100", "1.2676506e+30", nil},

// 在1和下一个float32之间精确的一半。
// 向下取整（偶数）。
// md5:f1baed8995860339
	{"1.000000059604644775390625", "1", nil},
	{"0x1.000001p0", "1", nil},
	// Slightly lower.
	{"1.000000059604644775390624", "1", nil},
	{"0x1.0000008p0", "1", nil},
	{"0x1.000000fp0", "1", nil},
	// Slightly higher.
	{"1.000000059604644775390626", "1.0000001", nil},
	{"0x1.000002p0", "1.0000001", nil},
	{"0x1.0000018p0", "1.0000001", nil},
	{"0x1.0000011p0", "1.0000001", nil},
	// 稍微高一点，但你必须读到结尾才能明白。. md5:b8273dbe0862f4f3
	{"1.000000059604644775390625" + strings.Repeat("0", 10000) + "1", "1.0000001", nil},
	{"0x1.000001" + strings.Repeat("0", 10000) + "1p0", "1.0000001", nil},

	// 最大 float32：(1<<128) * (1 - 2^-24) 
// 
// 这段Go语言注释的意思是：这是计算float32类型的最大值的一种方法。`(1<<128)` 表示将1左移128位，得到的结果是一个非常大的数值。`1 - 2^-24` 表示1减去2的负24次方，这个表达式的结果接近于2的-24次方，但由于浮点数精度问题，实际上会稍微大于这个值。两者相乘，就可以得到float32类型的上限，但由于浮点数表示有局限性，这个结果并不完全精确。. md5:9f39d19e096aa65d
	{"340282346638528859811704183484516925440", "3.4028235e+38", nil},
	{"-340282346638528859811704183484516925440", "-3.4028235e+38", nil},
	{"0x.ffffffp128", "3.4028235e+38", nil},
	{"-340282346638528859811704183484516925440", "-3.4028235e+38", nil},
	{"-0x.ffffffp128", "-3.4028235e+38", nil},
	// 下一个 float32 - 太大了. md5:f1b130c05a88712b
	{"3.4028236e38", "+Inf", ErrRange},
	{"-3.4028236e38", "-Inf", ErrRange},
	{"0x1.0p128", "+Inf", ErrRange},
	{"-0x1.0p128", "-Inf", ErrRange},
// 边界值为 3.40282356779...e+38
// 边界线 - 正常
// md5:bf670afe49bf180c
	{"3.402823567e38", "3.4028235e+38", nil},
	{"-3.402823567e38", "-3.4028235e+38", nil},
	{"0x.ffffff7fp128", "3.4028235e+38", nil},
	{"-0x.ffffff7fp128", "-3.4028235e+38", nil},
	// borderline - too large
	{"3.4028235678e38", "+Inf", ErrRange},
	{"-3.4028235678e38", "-Inf", ErrRange},
	{"0x.ffffff8p128", "+Inf", ErrRange},
	{"-0x.ffffff8p128", "-Inf", ErrRange},

	// Denormals: 小于2^-126的数值. md5:3dd4289972b09074
	{"1e-38", "1e-38", nil},
	{"1e-39", "1e-39", nil},
	{"1e-40", "1e-40", nil},
	{"1e-41", "1e-41", nil},
	{"1e-42", "1e-42", nil},
	{"1e-43", "1e-43", nil},
	{"1e-44", "1e-44", nil},
	{"6e-45", "6e-45", nil}, // 4p-149 = 5.6e-45
	{"5e-45", "6e-45", nil},

	// Smallest denormal
	{"1e-45", "1e-45", nil}, // 1p-149 = 1.4e-45
	{"2e-45", "1e-45", nil},
	{"3e-45", "3e-45", nil},

	// 接近于denormals和denormals。. md5:5213e06b56063627
	{"0x0.89aBcDp-125", "1.2643093e-38", nil},  // 0x0089abcd
	{"0x0.8000000p-125", "1.1754944e-38", nil}, // 0x00800000
	{"0x0.1234560p-125", "1.671814e-39", nil},  // 0x00123456
	{"0x0.1234567p-125", "1.671814e-39", nil},  // rounded down
	{"0x0.1234568p-125", "1.671814e-39", nil},  // rounded down
	{"0x0.1234569p-125", "1.671815e-39", nil},  // rounded up
	{"0x0.1234570p-125", "1.671815e-39", nil},  // 0x00123457
	{"0x0.0000010p-125", "1e-45", nil},         // 0x00000001
	{"0x0.00000081p-125", "1e-45", nil},        // rounded up
	{"0x0.0000008p-125", "0", nil},             // rounded down
	{"0x0.0000007p-125", "0", nil},             // rounded down

// 2^92 = 8388608p+69 = 4951760157141521099596496896 (4.9517602e27) 是一个精确的2的幂，需要8位小数才能正确解析回原来的值。
// 前面的float32为16777215p+68 = 4.95175986e+27
// 中间点是4.951760009。如果使用一个错误的算法，将前一个float32误认为是8388607p+69，会错误地简化为4.95176e+27。
// md5:2ecc7a844307a81e
	{"4951760157141521099596496896", "4.9517602e+27", nil},
}

type atofSimpleTest struct {
	x float64
	s string
}

var (
	atofOnce               sync.Once
	atofRandomTests        []atofSimpleTest
	benchmarksRandomBits   [1024]string
	benchmarksRandomNormal [1024]string
)

func initAtof() {
	atofOnce.Do(initAtofOnce)
}

func initAtofOnce() {
// atof 函数返回一个包含错误和字符串的 NumErrors 包装。请将上面的表格转换为中文。
// md5:ef4283f67521389e
	for i := range atoftests {
		test := &atoftests[i]
		if test.err != nil {
			test.err = &strconv2.NumError{"ParseFloat", test.in, test.err}
		}
	}
	for i := range atof32tests {
		test := &atof32tests[i]
		if test.err != nil {
			test.err = &strconv2.NumError{"ParseFloat", test.in, test.err}
		}
	}

	// 为测试和基准生成随机输入. md5:92a91f9ab7ee46d6
	if testing.Short() {
		atofRandomTests = make([]atofSimpleTest, 100)
	} else {
		atofRandomTests = make([]atofSimpleTest, 10000)
	}
	for i := range atofRandomTests {
		n := uint64(rand.Uint32())<<32 | uint64(rand.Uint32())
		x := math.Float64frombits(n)
		s := FormatFloat(x, 'g', -1, 64)
		atofRandomTests[i] = atofSimpleTest{x, s}
	}

	for i := range benchmarksRandomBits {
		bits := uint64(rand.Uint32())<<32 | uint64(rand.Uint32())
		x := math.Float64frombits(bits)
		benchmarksRandomBits[i] = FormatFloat(x, 'g', -1, 64)
	}

	for i := range benchmarksRandomNormal {
		x := rand.NormFloat64()
		benchmarksRandomNormal[i] = FormatFloat(x, 'g', -1, 64)
	}
}

//func TestParseFloatPrefix(t *testing.T) {
//	for i := range atoftests {
//		test := &atoftests[i]
//		if test.err != nil {
//			continue
//		}
//		// Adding characters that do not extend a number should not invalidate it.
//		// Test a few. The "i" and "init" cases test that we accept "infi", "infinit"
//		// correctly as "inf" with suffix.
//		for _, suffix := range []string{" ", "q", "+", "-", "<", "=", ">", "(", ")", "i", "init"} {
//			in := test.in + suffix
//			_, n, err := ParseFloatPrefix(in, 64)
//			if err != nil {
//				t.Errorf("ParseFloatPrefix(%q, 64): err = %v; want no error", in, err)
//			}
//			if n != len(test.in) {
//				t.Errorf("ParseFloatPrefix(%q, 64): n = %d; want %d", in, n, len(test.in))
//			}
//		}
//	}
//}

func testAtof(t *testing.T, opt bool) {
	initAtof()
	oldopt := SetOptimize(opt)
	for i := 0; i < len(atoftests); i++ {
		test := &atoftests[i]
		out, err := ParseFloat(test.in, 64)
		outs := FormatFloat(out, 'g', -1, 64)
		if outs != test.out || !reflect.DeepEqual(err, test.err) {
			t.Errorf("ParseFloat(%v, 64) = %v, %v want %v, %v",
				test.in, out, err, test.out, test.err)
		}

		if float64(float32(out)) == out {
			out, err := ParseFloat(test.in, 32)
			out32 := float32(out)
			if float64(out32) != out {
				t.Errorf("ParseFloat(%v, 32) = %v, not a float32 (closest is %v)", test.in, out, float64(out32))
				continue
			}
			outs := FormatFloat(float64(out32), 'g', -1, 32)
			if outs != test.out || !reflect.DeepEqual(err, test.err) {
				t.Errorf("ParseFloat(%v, 32) = %v, %v want %v, %v  # %v",
					test.in, out32, err, test.out, test.err, out)
			}
		}
	}
	for _, test := range atof32tests {
		out, err := ParseFloat(test.in, 32)
		out32 := float32(out)
		if float64(out32) != out {
			t.Errorf("ParseFloat(%v, 32) = %v, not a float32 (closest is %v)", test.in, out, float64(out32))
			continue
		}
		outs := FormatFloat(float64(out32), 'g', -1, 32)
		if outs != test.out || !reflect.DeepEqual(err, test.err) {
			t.Errorf("ParseFloat(%v, 32) = %v, %v want %v, %v  # %v",
				test.in, out32, err, test.out, test.err, out)
		}
	}
	SetOptimize(oldopt)
}

func TestAtof(t *testing.T) { testAtof(t, true) }

func TestAtofSlow(t *testing.T) { testAtof(t, false) }

func TestAtofRandom(t *testing.T) {
	initAtof()
	for _, test := range atofRandomTests {
		x, _ := ParseFloat(test.s, 64)
		switch {
		default:
			t.Errorf("number %s badly parsed as %b (expected %b)", test.s, x, test.x)
		case x == test.x:
		case math.IsNaN(test.x) && math.IsNaN(x):
		}
	}
	t.Logf("tested %d random numbers", len(atofRandomTests))
}

var roundTripCases = []struct {
	f float64
	s string
}{
// 问题2917。
// 如果FPU使用的是80位寄存器而不是64位寄存器，这个测试将会破坏优化的转换，通常是因为操作系统使用80位精度初始化了线程，而Go运行时没有修复浮点控制字。
// md5:4872af9151884505
	{8865794286000691 << 39, "4.87402195346389e+27"},
	{8865794286000692 << 39, "4.8740219534638903e+27"},
}

func TestRoundTrip(t *testing.T) {
	for _, tt := range roundTripCases {
		old := SetOptimize(false)
		s := FormatFloat(tt.f, 'g', -1, 64)
		if s != tt.s {
			t.Errorf("no-opt FormatFloat(%b) = %s, want %s", tt.f, s, tt.s)
		}
		f, err := ParseFloat(tt.s, 64)
		if f != tt.f || err != nil {
			t.Errorf("no-opt ParseFloat(%s) = %b, %v want %b, nil", tt.s, f, err, tt.f)
		}
		SetOptimize(true)
		s = FormatFloat(tt.f, 'g', -1, 64)
		if s != tt.s {
			t.Errorf("opt FormatFloat(%b) = %s, want %s", tt.f, s, tt.s)
		}
		f, err = ParseFloat(tt.s, 64)
		if f != tt.f || err != nil {
			t.Errorf("opt ParseFloat(%s) = %b, %v want %b, nil", tt.s, f, err, tt.f)
		}
		SetOptimize(old)
	}
}

// TestRoundTrip32尝试所有有限的正浮点32值的很小一部分。. md5:bcc3b233d2561c2a
func TestRoundTrip32(t *testing.T) {
	step := uint32(997)
	if testing.Short() {
		step = 99991
	}
	count := 0
	for i := uint32(0); i < 0xff<<23; i += step {
		f := math.Float32frombits(i)
		if i&1 == 1 {
			f = -f // negative
		}
		s := FormatFloat(float64(f), 'g', -1, 32)

		parsed, err := ParseFloat(s, 32)
		parsed32 := float32(parsed)
		switch {
		case err != nil:
			t.Errorf("ParseFloat(%q, 32) gave error %s", s, err)
		case float64(parsed32) != parsed:
			t.Errorf("ParseFloat(%q, 32) = %v, not a float32 (nearest is %v)", s, parsed, parsed32)
		case parsed32 != f:
			t.Errorf("ParseFloat(%q, 32) = %b (expected %b)", s, parsed32, f)
		}
		count++
	}
	t.Logf("tested %d float32's", count)
}

// Issue 42297：很多现有代码可能会误调用 ParseFloat(s, 10)
// 或 ParseFloat(s, 0)，因此允许 bitSize 的值不仅限于 32 和 64。
// md5:5a23d02fde2bdecb
func TestParseFloatIncorrectBitSize(t *testing.T) {
	const s = "1.5e308"
	const want = 1.5e308

	for _, bitSize := range []int{0, 10, 100, 128} {
		f, err := ParseFloat(s, bitSize)
		if err != nil {
			t.Fatalf("ParseFloat(%q, %d) gave error %s", s, bitSize, err)
		}
		if f != want {
			t.Fatalf("ParseFloat(%q, %d) = %g (expected %g)", s, bitSize, f, want)
		}
	}
}

func BenchmarkAtof64Decimal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ParseFloat("33909", 64)
	}
}

func BenchmarkAtof64Float(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ParseFloat("339.7784", 64)
	}
}

func BenchmarkAtof64FloatExp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ParseFloat("-5.09e75", 64)
	}
}

func BenchmarkAtof64Big(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ParseFloat("123456789123456789123456789", 64)
	}
}

func BenchmarkAtof64RandomBits(b *testing.B) {
	initAtof()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ParseFloat(benchmarksRandomBits[i%1024], 64)
	}
}

func BenchmarkAtof64RandomFloats(b *testing.B) {
	initAtof()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ParseFloat(benchmarksRandomNormal[i%1024], 64)
	}
}

func BenchmarkAtof64RandomLongFloats(b *testing.B) {
	initAtof()
	samples := make([]string, len(atofRandomTests))
	for i, t := range atofRandomTests {
		samples[i] = FormatFloat(t.x, 'g', 20, 64)
	}
	b.ResetTimer()
	idx := 0
	for i := 0; i < b.N; i++ {
		ParseFloat(samples[idx], 64)
		idx++
		if idx == len(samples) {
			idx = 0
		}
	}
}

func BenchmarkAtof32Decimal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ParseFloat("33909", 32)
	}
}

func BenchmarkAtof32Float(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ParseFloat("339.778", 32)
	}
}

func BenchmarkAtof32FloatExp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ParseFloat("12.3456e32", 32)
	}
}

func BenchmarkAtof32Random(b *testing.B) {
	n := uint32(997)
	var float32strings [4096]string
	for i := range float32strings {
		n = (99991*n + 42) % (0xff << 23)
		float32strings[i] = FormatFloat(float64(math.Float32frombits(n)), 'g', -1, 32)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ParseFloat(float32strings[i%4096], 32)
	}
}

func BenchmarkAtof32RandomLong(b *testing.B) {
	n := uint32(997)
	var float32strings [4096]string
	for i := range float32strings {
		n = (99991*n + 42) % (0xff << 23)
		float32strings[i] = FormatFloat(float64(math.Float32frombits(n)), 'g', 20, 32)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ParseFloat(float32strings[i%4096], 32)
	}
}
