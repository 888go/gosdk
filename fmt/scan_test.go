//版权所有2009年Go作者。所有权利保留。
//使用此源代码受BSD风格
//可以在LICENSE文件中找到的许可证。
// md5:2e9dc81828a3be8a

package fmt_test

import (
	"bufio"
	"bytes"
	"errors"
	fmt2 "fmt"
	. "github.com/888go/gosdk/fmt"
	"io"
	"math"
	"reflect"
	"regexp"
	"strings"
	"testing"
	"testing/iotest"
	"unicode/utf8"
)

type ScanTest struct {
	text string
	in   any
	out  any
}

type ScanfTest struct {
	format string
	text   string
	in     any
	out    any
}

type ScanfMultiTest struct {
	format string
	text   string
	in     []any
	out    []any
	err    string
}

var (
	boolVal              bool
	intVal               int
	int8Val              int8
	int16Val             int16
	int32Val             int32
	int64Val             int64
	uintVal              uint
	uint8Val             uint8
	uint16Val            uint16
	uint32Val            uint32
	uint64Val            uint64
	uintptrVal           uintptr
	float32Val           float32
	float64Val           float64
	stringVal            string
	bytesVal             []byte
	runeVal              rune
	complex64Val         complex64
	complex128Val        complex128
	renamedBoolVal       renamedBool
	renamedIntVal        renamedInt
	renamedInt8Val       renamedInt8
	renamedInt16Val      renamedInt16
	renamedInt32Val      renamedInt32
	renamedInt64Val      renamedInt64
	renamedUintVal       renamedUint
	renamedUint8Val      renamedUint8
	renamedUint16Val     renamedUint16
	renamedUint32Val     renamedUint32
	renamedUint64Val     renamedUint64
	renamedUintptrVal    renamedUintptr
	renamedStringVal     renamedString
	renamedBytesVal      renamedBytes
	renamedFloat32Val    renamedFloat32
	renamedFloat64Val    renamedFloat64
	renamedComplex64Val  renamedComplex64
	renamedComplex128Val renamedComplex128
)

// Xs 接受任何非空的动词字符序列. md5:e2af7ad9167fbb6a
type Xs string

func (x *Xs) Scan(state fmt2.ScanState, verb rune) error {
	tok, err := state.Token(true, func(r rune) bool { return r == verb })
	if err != nil {
		return err
	}
	s := string(tok)
	if !regexp.MustCompile("^" + string(verb) + "+$").MatchString(s) {
		return errors.New("syntax error for xs")
	}
	*x = Xs(s)
	return nil
}

var xVal Xs

// IntString 接受一个整数后面紧接着一个字符串。
// 它测试在一个扫描内部嵌套扫描的情况。
// md5:141a29b2922ced2b
type IntString struct {
	i int
	s string
}

func (s *IntString) Scan(state fmt2.ScanState, verb rune) error {
	if _, err := Fscan(state, &s.i); err != nil {
		return err
	}

	tok, err := state.Token(true, nil)
	if err != nil {
		return err
	}
	s.s = string(tok)
	return nil
}

var intStringVal IntString

var scanTests = []ScanTest{
	// Basic types
	{"T\n", &boolVal, true},  // 将布尔测试值切换，以确保它们已写入. md5:74cbf6ee72eb57a1
	{"F\n", &boolVal, false}, // restored to zero value
	{"21\n", &intVal, 21},
	{"2_1\n", &intVal, 21},
	{"0\n", &intVal, 0},
	{"000\n", &intVal, 0},
	{"0x10\n", &intVal, 0x10},
	{"0x_1_0\n", &intVal, 0x10},
	{"-0x10\n", &intVal, -0x10},
	{"0377\n", &intVal, 0377},
	{"0_3_7_7\n", &intVal, 0377},
	{"0o377\n", &intVal, 0377},
	{"0o_3_7_7\n", &intVal, 0377},
	{"-0377\n", &intVal, -0377},
	{"-0o377\n", &intVal, -0377},
	{"0\n", &uintVal, uint(0)},
	{"000\n", &uintVal, uint(0)},
	{"0x10\n", &uintVal, uint(0x10)},
	{"0377\n", &uintVal, uint(0377)},
	{"22\n", &int8Val, int8(22)},
	{"23\n", &int16Val, int16(23)},
	{"24\n", &int32Val, int32(24)},
	{"25\n", &int64Val, int64(25)},
	{"127\n", &int8Val, int8(127)},
	{"-21\n", &intVal, -21},
	{"-22\n", &int8Val, int8(-22)},
	{"-23\n", &int16Val, int16(-23)},
	{"-24\n", &int32Val, int32(-24)},
	{"-25\n", &int64Val, int64(-25)},
	{"-128\n", &int8Val, int8(-128)},
	{"+21\n", &intVal, +21},
	{"+22\n", &int8Val, int8(+22)},
	{"+23\n", &int16Val, int16(+23)},
	{"+24\n", &int32Val, int32(+24)},
	{"+25\n", &int64Val, int64(+25)},
	{"+127\n", &int8Val, int8(+127)},
	{"26\n", &uintVal, uint(26)},
	{"27\n", &uint8Val, uint8(27)},
	{"28\n", &uint16Val, uint16(28)},
	{"29\n", &uint32Val, uint32(29)},
	{"30\n", &uint64Val, uint64(30)},
	{"31\n", &uintptrVal, uintptr(31)},
	{"255\n", &uint8Val, uint8(255)},
	{"32767\n", &int16Val, int16(32767)},
	{"2.3\n", &float64Val, 2.3},
	{"2.3e1\n", &float32Val, float32(2.3e1)},
	{"2.3e2\n", &float64Val, 2.3e2},
	{"2.3p2\n", &float64Val, 2.3 * 4},
	{"2.3p+2\n", &float64Val, 2.3 * 4},
	{"2.3p+66\n", &float64Val, 2.3 * (1 << 66)},
	{"2.3p-66\n", &float64Val, 2.3 / (1 << 66)},
	{"0x2.3p-66\n", &float64Val, float64(0x23) / (1 << 70)},
	{"2_3.4_5\n", &float64Val, 23.45},
	{"2.35\n", &stringVal, "2.35"},
	{"2345678\n", &bytesVal, []byte("2345678")},
	{"(3.4e1-2i)\n", &complex128Val, 3.4e1 - 2i},
	{"-3.45e1-3i\n", &complex64Val, complex64(-3.45e1 - 3i)},
	{"-.45e1-1e2i\n", &complex128Val, complex128(-.45e1 - 100i)},
	{"-.4_5e1-1E2i\n", &complex128Val, complex128(-.45e1 - 100i)},
	{"0x1.0p1+0x1.0P2i\n", &complex128Val, complex128(2 + 4i)},
	{"-0x1p1-0x1p2i\n", &complex128Val, complex128(-2 - 4i)},
	{"-0x1ep-1-0x1p2i\n", &complex128Val, complex128(-15 - 4i)},
	{"-0x1_Ep-1-0x1p0_2i\n", &complex128Val, complex128(-15 - 4i)},
	{"hello\n", &stringVal, "hello"},

	// 换行符后跟一个新行。 (我们始终将\r\n视为\n). md5:6d9a989cf65156d2
	{"hello\r\n", &stringVal, "hello"},
	{"27\r\n", &uint8Val, uint8(27)},

	// Renamed types
	{"true\n", &renamedBoolVal, renamedBool(true)},
	{"F\n", &renamedBoolVal, renamedBool(false)},
	{"101\n", &renamedIntVal, renamedInt(101)},
	{"102\n", &renamedIntVal, renamedInt(102)},
	{"103\n", &renamedUintVal, renamedUint(103)},
	{"104\n", &renamedUintVal, renamedUint(104)},
	{"105\n", &renamedInt8Val, renamedInt8(105)},
	{"106\n", &renamedInt16Val, renamedInt16(106)},
	{"107\n", &renamedInt32Val, renamedInt32(107)},
	{"108\n", &renamedInt64Val, renamedInt64(108)},
	{"109\n", &renamedUint8Val, renamedUint8(109)},
	{"110\n", &renamedUint16Val, renamedUint16(110)},
	{"111\n", &renamedUint32Val, renamedUint32(111)},
	{"112\n", &renamedUint64Val, renamedUint64(112)},
	{"113\n", &renamedUintptrVal, renamedUintptr(113)},
	{"114\n", &renamedStringVal, renamedString("114")},
	{"115\n", &renamedBytesVal, renamedBytes([]byte("115"))},

	// Custom scanners.
	{"  vvv ", &xVal, Xs("vvv")},
	{" 1234hello", &intStringVal, IntString{1234, "hello"}},

	// Fixed bugs
	{"2147483648\n", &int64Val, int64(2147483648)}, // was: integer overflow
}

var scanfTests = []ScanfTest{
	{"%v", "TRUE\n", &boolVal, true},
	{"%t", "false\n", &boolVal, false},
	{"%v", "-71\n", &intVal, -71},
	{"%v", "-7_1\n", &intVal, -71},
	{"%v", "0b111\n", &intVal, 7},
	{"%v", "0b_1_1_1\n", &intVal, 7},
	{"%v", "0377\n", &intVal, 0377},
	{"%v", "0_3_7_7\n", &intVal, 0377},
	{"%v", "0o377\n", &intVal, 0377},
	{"%v", "0o_3_7_7\n", &intVal, 0377},
	{"%v", "0x44\n", &intVal, 0x44},
	{"%v", "0x_4_4\n", &intVal, 0x44},
	{"%d", "72\n", &intVal, 72},
	{"%c", "a\n", &runeVal, 'a'},
	{"%c", "\u5072\n", &runeVal, '\u5072'},
	{"%c", "\u1234\n", &runeVal, '\u1234'},
	{"%d", "73\n", &int8Val, int8(73)},
	{"%d", "+74\n", &int16Val, int16(74)},
	{"%d", "75\n", &int32Val, int32(75)},
	{"%d", "76\n", &int64Val, int64(76)},
	{"%b", "1001001\n", &intVal, 73},
	{"%o", "075\n", &intVal, 075},
	{"%x", "a75\n", &intVal, 0xa75},
	{"%v", "71\n", &uintVal, uint(71)},
	{"%d", "72\n", &uintVal, uint(72)},
	{"%d", "7_2\n", &uintVal, uint(7)}, // 只有 %v 支持使用下划线. md5:0d5e8d26bf807d19
	{"%d", "73\n", &uint8Val, uint8(73)},
	{"%d", "74\n", &uint16Val, uint16(74)},
	{"%d", "75\n", &uint32Val, uint32(75)},
	{"%d", "76\n", &uint64Val, uint64(76)},
	{"%d", "77\n", &uintptrVal, uintptr(77)},
	{"%b", "1001001\n", &uintVal, uint(73)},
	{"%b", "100_1001\n", &uintVal, uint(4)},
	{"%o", "075\n", &uintVal, uint(075)},
	{"%o", "07_5\n", &uintVal, uint(07)}, // 只有 %v 支持使用下划线. md5:0d5e8d26bf807d19
	{"%x", "a75\n", &uintVal, uint(0xa75)},
	{"%x", "A75\n", &uintVal, uint(0xa75)},
	{"%x", "A7_5\n", &uintVal, uint(0xa7)}, // 只有 %v 支持使用下划线. md5:0d5e8d26bf807d19
	{"%U", "U+1234\n", &intVal, int(0x1234)},
	{"%U", "U+4567\n", &uintVal, uint(0x4567)},

	{"%e", "2.3\n", &float64Val, 2.3},
	{"%E", "2.3e1\n", &float32Val, float32(2.3e1)},
	{"%f", "2.3e2\n", &float64Val, 2.3e2},
	{"%g", "2.3p2\n", &float64Val, 2.3 * 4},
	{"%G", "2.3p+2\n", &float64Val, 2.3 * 4},
	{"%v", "2.3p+66\n", &float64Val, 2.3 * (1 << 66)},
	{"%f", "2.3p-66\n", &float64Val, 2.3 / (1 << 66)},
	{"%G", "0x2.3p-66\n", &float64Val, float64(0x23) / (1 << 70)},
	{"%E", "2_3.4_5\n", &float64Val, 23.45},

	// Strings
	{"%s", "using-%s\n", &stringVal, "using-%s"},
	{"%x", "7573696e672d2578\n", &stringVal, "using-%x"},
	{"%X", "7573696E672D2558\n", &stringVal, "using-%X"},
	{"%q", `"quoted\twith\\do\u0075bl\x65s"` + "\n", &stringVal, "quoted\twith\\doubles"},
	{"%q", "`quoted with backs`\n", &stringVal, "quoted with backs"},

	// Byte slices
	{"%s", "bytes-%s\n", &bytesVal, []byte("bytes-%s")},
	{"%x", "62797465732d2578\n", &bytesVal, []byte("bytes-%x")},
	{"%X", "62797465732D2558\n", &bytesVal, []byte("bytes-%X")},
	{"%q", `"bytes\rwith\vdo\u0075bl\x65s"` + "\n", &bytesVal, []byte("bytes\rwith\vdoubles")},
	{"%q", "`bytes with backs`\n", &bytesVal, []byte("bytes with backs")},

	// Renamed types
	{"%v\n", "true\n", &renamedBoolVal, renamedBool(true)},
	{"%t\n", "F\n", &renamedBoolVal, renamedBool(false)},
	{"%v", "101\n", &renamedIntVal, renamedInt(101)},
	{"%c", "\u0101\n", &renamedIntVal, renamedInt('\u0101')},
	{"%o", "0146\n", &renamedIntVal, renamedInt(102)},
	{"%v", "103\n", &renamedUintVal, renamedUint(103)},
	{"%d", "104\n", &renamedUintVal, renamedUint(104)},
	{"%d", "105\n", &renamedInt8Val, renamedInt8(105)},
	{"%d", "106\n", &renamedInt16Val, renamedInt16(106)},
	{"%d", "107\n", &renamedInt32Val, renamedInt32(107)},
	{"%d", "108\n", &renamedInt64Val, renamedInt64(108)},
	{"%x", "6D\n", &renamedUint8Val, renamedUint8(109)},
	{"%o", "0156\n", &renamedUint16Val, renamedUint16(110)},
	{"%d", "111\n", &renamedUint32Val, renamedUint32(111)},
	{"%d", "112\n", &renamedUint64Val, renamedUint64(112)},
	{"%d", "113\n", &renamedUintptrVal, renamedUintptr(113)},
	{"%s", "114\n", &renamedStringVal, renamedString("114")},
	{"%q", "\"1155\"\n", &renamedBytesVal, renamedBytes([]byte("1155"))},
	{"%g", "116e1\n", &renamedFloat32Val, renamedFloat32(116e1)},
	{"%g", "-11.7e+1", &renamedFloat64Val, renamedFloat64(-11.7e+1)},
	{"%g", "11+6e1i\n", &renamedComplex64Val, renamedComplex64(11 + 6e1i)},
	{"%g", "-11.+7e+1i", &renamedComplex128Val, renamedComplex128(-11. + 7e+1i)},

	// Interesting formats
	{"here is\tthe value:%d", "here is   the\tvalue:118\n", &intVal, 118},
	{"%% %%:%d", "% %:119\n", &intVal, 119},
	{"%d%%", "42%", &intVal, 42}, // %% at end of string.

	// Corner cases
	{"%x", "FFFFFFFF\n", &uint32Val, uint32(0xFFFFFFFF)},

	// Custom scanner.
	{"%s", "  sss ", &xVal, Xs("sss")},
	{"%2s", "sssss", &xVal, Xs("ss")},

	// Fixed bugs
	{"%d\n", "27\n", &intVal, 27},         // ok
	{"%d\n", "28 \n", &intVal, 28},        // 原来是：“意外的换行符”. md5:64dec408d934bbe5
	{"%v", "0", &intVal, 0},               // 曾经是："EOF"；0被用作基础前缀，不计入计数。. md5:ddffebc21736e294
	{"%v", "0", &uintVal, uint(0)},        // 曾经是："EOF"；0被用作基础前缀，不计入计数。. md5:ddffebc21736e294
	{"%c", " ", &uintVal, uint(' ')},      // %c必须接受空格。. md5:6e35ac5146edbe5e
	{"%c", "\t", &uintVal, uint('\t')},    // %c 必须接受任何空格。. md5:7117f90647194eaf
	{"%c", "\n", &uintVal, uint('\n')},    // %c 必须接受任何空格。. md5:7117f90647194eaf
	{"%d%%", "23%\n", &uintVal, uint(23)}, // %% matches literal %.
	{"%%%d", "%23\n", &uintVal, uint(23)}, // %% matches literal %.

	// space handling
	{"%d", "27", &intVal, 27},
	{"%d", "27 ", &intVal, 27},
	{"%d", " 27", &intVal, 27},
	{"%d", " 27 ", &intVal, 27},

	{"X%d", "X27", &intVal, 27},
	{"X%d", "X27 ", &intVal, 27},
	{"X%d", "X 27", &intVal, 27},
	{"X%d", "X 27 ", &intVal, 27},

	{"X %d", "X27", &intVal, nil},  // 预期在输入中有一个空格以匹配格式. md5:cbe1e1f4db29e55f
	{"X %d", "X27 ", &intVal, nil}, // 预期在输入中有一个空格以匹配格式. md5:cbe1e1f4db29e55f
	{"X %d", "X 27", &intVal, 27},
	{"X %d", "X 27 ", &intVal, 27},

	{"%dX", "27X", &intVal, 27},
	{"%dX", "27 X", &intVal, nil}, // 输入不符合格式. md5:d88e7c247ff5f55c
	{"%dX", " 27X", &intVal, 27},
	{"%dX", " 27 X", &intVal, nil}, // 输入不符合格式. md5:d88e7c247ff5f55c

	{"%d X", "27X", &intVal, nil}, // 预期在输入中有一个空格以匹配格式. md5:cbe1e1f4db29e55f
	{"%d X", "27 X", &intVal, 27},
	{"%d X", " 27X", &intVal, nil}, // 预期在输入中有一个空格以匹配格式. md5:cbe1e1f4db29e55f
	{"%d X", " 27 X", &intVal, 27},

	{"X %d X", "X27X", &intVal, nil},  // 预期在输入中有一个空格以匹配格式. md5:cbe1e1f4db29e55f
	{"X %d X", "X27 X", &intVal, nil}, // 预期在输入中有一个空格以匹配格式. md5:cbe1e1f4db29e55f
	{"X %d X", "X 27X", &intVal, nil}, // 预期在输入中有一个空格以匹配格式. md5:cbe1e1f4db29e55f
	{"X %d X", "X 27 X", &intVal, 27},

	{"X %s X", "X27X", &stringVal, nil},  // 预期在输入中有一个空格以匹配格式. md5:cbe1e1f4db29e55f
	{"X %s X", "X27 X", &stringVal, nil}, // 预期在输入中有一个空格以匹配格式. md5:cbe1e1f4db29e55f
	{"X %s X", "X 27X", &stringVal, nil}, // unexpected EOF
	{"X %s X", "X 27 X", &stringVal, "27"},

	{"X%sX", "X27X", &stringVal, nil},   // unexpected EOF
	{"X%sX", "X27 X", &stringVal, nil},  // 输入不符合格式. md5:d88e7c247ff5f55c
	{"X%sX", "X 27X", &stringVal, nil},  // unexpected EOF
	{"X%sX", "X 27 X", &stringVal, nil}, // 输入不符合格式. md5:d88e7c247ff5f55c

	{"X%s", "X27", &stringVal, "27"},
	{"X%s", "X27 ", &stringVal, "27"},
	{"X%s", "X 27", &stringVal, "27"},
	{"X%s", "X 27 ", &stringVal, "27"},

	{"X%dX", "X27X", &intVal, 27},
	{"X%dX", "X27 X", &intVal, nil}, // 输入不符合格式. md5:d88e7c247ff5f55c
	{"X%dX", "X 27X", &intVal, 27},
	{"X%dX", "X 27 X", &intVal, nil}, // 输入不符合格式. md5:d88e7c247ff5f55c

	{"X%dX", "X27X", &intVal, 27},
	{"X%dX", "X27X ", &intVal, 27},
	{"X%dX", " X27X", &intVal, nil},  // 输入不符合格式. md5:d88e7c247ff5f55c
	{"X%dX", " X27X ", &intVal, nil}, // 输入不符合格式. md5:d88e7c247ff5f55c

	{"X%dX\n", "X27X", &intVal, 27},
	{"X%dX \n", "X27X ", &intVal, 27},
	{"X%dX\n", "X27X\n", &intVal, 27},
	{"X%dX\n", "X27X \n", &intVal, 27},

	{"X%dX \n", "X27X", &intVal, 27},
	{"X%dX \n", "X27X ", &intVal, 27},
	{"X%dX \n", "X27X\n", &intVal, 27},
	{"X%dX \n", "X27X \n", &intVal, 27},

	{"X%c", "X\n", &runeVal, '\n'},
	{"X%c", "X \n", &runeVal, ' '},
	{"X %c", "X!", &runeVal, nil},  // 预期在输入中有一个空格以匹配格式. md5:cbe1e1f4db29e55f
	{"X %c", "X\n", &runeVal, nil}, // 输入中的换行符不符合格式. md5:12dd69534dfbf0a3
	{"X %c", "X !", &runeVal, '!'},
	{"X %c", "X \n", &runeVal, '\n'},

	{" X%dX", "X27X", &intVal, nil},  // 预期在输入中有一个空格以匹配格式. md5:cbe1e1f4db29e55f
	{" X%dX", "X27X ", &intVal, nil}, // 预期在输入中有一个空格以匹配格式. md5:cbe1e1f4db29e55f
	{" X%dX", " X27X", &intVal, 27},
	{" X%dX", " X27X ", &intVal, 27},

	{"X%dX ", "X27X", &intVal, 27},
	{"X%dX ", "X27X ", &intVal, 27},
	{"X%dX ", " X27X", &intVal, nil},  // 输入不符合格式. md5:d88e7c247ff5f55c
	{"X%dX ", " X27X ", &intVal, nil}, // 输入不符合格式. md5:d88e7c247ff5f55c

	{" X%dX ", "X27X", &intVal, nil},  // 预期在输入中有一个空格以匹配格式. md5:cbe1e1f4db29e55f
	{" X%dX ", "X27X ", &intVal, nil}, // 预期在输入中有一个空格以匹配格式. md5:cbe1e1f4db29e55f
	{" X%dX ", " X27X", &intVal, 27},
	{" X%dX ", " X27X ", &intVal, 27},

	{"%d\nX", "27\nX", &intVal, 27},
	{"%dX\n X", "27X\n X", &intVal, 27},
}

var overflowTests = []ScanTest{
	{"128", &int8Val, 0},
	{"32768", &int16Val, 0},
	{"-129", &int8Val, 0},
	{"-32769", &int16Val, 0},
	{"256", &uint8Val, 0},
	{"65536", &uint16Val, 0},
	{"1e100", &float32Val, 0},
	{"1e500", &float64Val, 0},
	{"(1e100+0i)", &complex64Val, 0},
	{"(1+1e100i)", &complex64Val, 0},
	{"(1-1e500i)", &complex128Val, 0},
}

var truth bool
var i, j, k int
var f float64
var s, t string
var c complex128
var x, y Xs
var z IntString
var r1, r2, r3 rune

var multiTests = []ScanfMultiTest{
	{"", "", []any{}, []any{}, ""},
	{"%d", "23", args(&i), args(23), ""},
	{"%2s%3s", "22333", args(&s, &t), args("22", "333"), ""},
	{"%2d%3d", "44555", args(&i, &j), args(44, 555), ""},
	{"%2d.%3d", "66.777", args(&i, &j), args(66, 777), ""},
	{"%d, %d", "23, 18", args(&i, &j), args(23, 18), ""},
	{"%3d22%3d", "33322333", args(&i, &j), args(333, 333), ""},
	{"%6vX=%3fY", "3+2iX=2.5Y", args(&c, &f), args((3 + 2i), 2.5), ""},
	{"%d%s", "123abc", args(&i, &s), args(123, "abc"), ""},
	{"%c%c%c", "2\u50c2X", args(&r1, &r2, &r3), args('2', '\u50c2', 'X'), ""},
	{"%5s%d", " 1234567 ", args(&s, &i), args("12345", 67), ""},
	{"%5s%d", " 12 34 567 ", args(&s, &i), args("12", 34), ""},

	// Custom scanners.
	{"%e%f", "eefffff", args(&x, &y), args(Xs("ee"), Xs("fffff")), ""},
	{"%4v%s", "12abcd", args(&z, &s), args(IntString{12, "ab"}, "cd"), ""},

	// Errors
	{"%t", "23 18", args(&i), nil, "bad verb"},
	{"%d %d %d", "23 18", args(&i, &j), args(23, 18), "too few operands"},
	{"%d %d", "23 18 27", args(&i, &j, &k), args(23, 18), "too many operands"},
	{"%c", "\u0100", args(&int8Val), nil, "overflow"},
	{"X%d", "10X", args(&intVal), nil, "input does not match format"},
	{"%d%", "42%", args(&intVal), args(42), "missing verb: % at end of format string"},
	{"%d% ", "42%", args(&intVal), args(42), "too few operands for format '% '"}, // 略显奇怪的错误，但却是正确的。. md5:3fc465eba5431502
	{"%%%d", "xxx 42", args(&intVal), args(42), "missing literal %"},
	{"%%%d", "x42", args(&intVal), args(42), "missing literal %"},
	{"%%%d", "42", args(&intVal), args(42), "missing literal %"},

	// 错误的UTF-8编码：应该查看每一个字节。. md5:4dad33770bdc46c7
	{"%c%c%c", "\xc2X\xc2", args(&r1, &r2, &r3), args(utf8.RuneError, 'X', utf8.RuneError), ""},

	// Fixed bugs
	{"%v%v", "FALSE23", args(&truth, &i), args(false, 23), ""},
}

var readers = []struct {
	name string
	f    func(string) io.Reader
}{
	{"StringReader", func(s string) io.Reader {
		return strings.NewReader(s)
	}},
	{"ReaderOnly", func(s string) io.Reader {
		return struct{ io.Reader }{strings.NewReader(s)}
	}},
	{"OneByteReader", func(s string) io.Reader {
		return iotest.OneByteReader(strings.NewReader(s))
	}},
	{"DataErrReader", func(s string) io.Reader {
		return iotest.DataErrReader(strings.NewReader(s))
	}},
}

func testScan(t *testing.T, f func(string) io.Reader, scan func(r io.Reader, a ...any) (int, error)) {
	for _, test := range scanTests {
		r := f(test.text)
		n, err := scan(r, test.in)
		if err != nil {
			m := ""
			if n > 0 {
				m = Sprintf(" (%d fields ok)", n)
			}
			t.Errorf("got error scanning %q: %s%s", test.text, err, m)
			continue
		}
		if n != 1 {
			t.Errorf("count error on entry %q: got %d", test.text, n)
			continue
		}
		// 进来的值可能是指针. md5:4f1140090237710f
		v := reflect.ValueOf(test.in)
		if p := v; p.Kind() == reflect.Pointer {
			v = p.Elem()
		}
		val := v.Interface()
		if !reflect.DeepEqual(val, test.out) {
			t.Errorf("scanning %q: expected %#v got %#v, type %T", test.text, test.out, val, val)
		}
	}
}

func TestScan(t *testing.T) {
	for _, r := range readers {
		t.Run(r.name, func(t *testing.T) {
			testScan(t, r.f, Fscan)
		})
	}
}

func TestScanln(t *testing.T) {
	for _, r := range readers {
		t.Run(r.name, func(t *testing.T) {
			testScan(t, r.f, Fscanln)
		})
	}
}

func TestScanf(t *testing.T) {
	for _, test := range scanfTests {
		n, err := Sscanf(test.text, test.format, test.in)
		if err != nil {
			if test.out != nil {
				t.Errorf("Sscanf(%q, %q): unexpected error: %v", test.text, test.format, err)
			}
			continue
		}
		if test.out == nil {
			t.Errorf("Sscanf(%q, %q): unexpected success", test.text, test.format)
			continue
		}
		if n != 1 {
			t.Errorf("Sscanf(%q, %q): parsed %d field, want 1", test.text, test.format, n)
			continue
		}
		// 进来的值可能是指针. md5:4f1140090237710f
		v := reflect.ValueOf(test.in)
		if p := v; p.Kind() == reflect.Pointer {
			v = p.Elem()
		}
		val := v.Interface()
		if !reflect.DeepEqual(val, test.out) {
			t.Errorf("Sscanf(%q, %q): parsed value %T(%#v), want %T(%#v)", test.text, test.format, val, val, test.out, test.out)
		}
	}
}

func TestScanOverflow(t *testing.T) {
	// 不同的机器和不同的类型会以不同的字符串报告错误。. md5:5c4eb09a7c2a976b
	re := regexp.MustCompile("overflow|too large|out of range|not representable")
	for _, test := range overflowTests {
		_, err := Sscan(test.text, test.in)
		if err == nil {
			t.Errorf("expected overflow scanning %q", test.text)
			continue
		}
		if !re.MatchString(err.Error()) {
			t.Errorf("expected overflow error scanning %q: %s", test.text, err)
		}
	}
}

func verifyNaN(str string, t *testing.T) {
	var f float64
	var f32 float32
	var f64 float64
	text := str + " " + str + " " + str
	n, err := Fscan(strings.NewReader(text), &f, &f32, &f64)
	if err != nil {
		t.Errorf("got error scanning %q: %s", text, err)
	}
	if n != 3 {
		t.Errorf("count error scanning %q: got %d", text, n)
	}
	if !math.IsNaN(float64(f)) || !math.IsNaN(float64(f32)) || !math.IsNaN(f64) {
		t.Errorf("didn't get NaNs scanning %q: got %g %g %g", text, f, f32, f64)
	}
}

func TestNaN(t *testing.T) {
	for _, s := range []string{"nan", "NAN", "NaN"} {
		verifyNaN(s, t)
	}
}

func verifyInf(str string, t *testing.T) {
	var f float64
	var f32 float32
	var f64 float64
	text := str + " " + str + " " + str
	n, err := Fscan(strings.NewReader(text), &f, &f32, &f64)
	if err != nil {
		t.Errorf("got error scanning %q: %s", text, err)
	}
	if n != 3 {
		t.Errorf("count error scanning %q: got %d", text, n)
	}
	sign := 1
	if str[0] == '-' {
		sign = -1
	}
	if !math.IsInf(float64(f), sign) || !math.IsInf(float64(f32), sign) || !math.IsInf(f64, sign) {
		t.Errorf("didn't get right Infs scanning %q: got %g %g %g", text, f, f32, f64)
	}
}

func TestInf(t *testing.T) {
	for _, s := range []string{"inf", "+inf", "-inf", "INF", "-INF", "+INF", "Inf", "-Inf", "+Inf"} {
		verifyInf(s, t)
	}
}

func testScanfMulti(t *testing.T, f func(string) io.Reader) {
	sliceType := reflect.TypeOf(make([]any, 1))
	for _, test := range multiTests {
		r := f(test.text)
		n, err := Fscanf(r, test.format, test.in...)
		if err != nil {
			if test.err == "" {
				t.Errorf("got error scanning (%q, %q): %q", test.format, test.text, err)
			} else if !strings.Contains(err.Error(), test.err) {
				t.Errorf("got wrong error scanning (%q, %q): %q; expected %q", test.format, test.text, err, test.err)
			}
			continue
		}
		if test.err != "" {
			t.Errorf("expected error %q error scanning (%q, %q)", test.err, test.format, test.text)
		}
		if n != len(test.out) {
			t.Errorf("count error on entry (%q, %q): expected %d got %d", test.format, test.text, len(test.out), n)
			continue
		}
		// 将指针切片转换为值切片. md5:04390588bb10d1ee
		resultVal := reflect.MakeSlice(sliceType, n, n)
		for i := 0; i < n; i++ {
			v := reflect.ValueOf(test.in[i]).Elem()
			resultVal.Index(i).Set(v)
		}
		result := resultVal.Interface()
		if !reflect.DeepEqual(result, test.out) {
			t.Errorf("scanning (%q, %q): expected %#v got %#v", test.format, test.text, test.out, result)
		}
	}
}

func TestScanfMulti(t *testing.T) {
	for _, r := range readers {
		t.Run(r.name, func(t *testing.T) {
			testScanfMulti(t, r.f)
		})
	}
}

func TestScanMultiple(t *testing.T) {
	var a int
	var s string
	n, err := Sscan("123abc", &a, &s)
	if n != 2 {
		t.Errorf("Sscan count error: expected 2: got %d", n)
	}
	if err != nil {
		t.Errorf("Sscan expected no error; got %s", err)
	}
	if a != 123 || s != "abc" {
		t.Errorf("Sscan wrong values: got (%d %q) expected (123 \"abc\")", a, s)
	}
	n, err = Sscan("asdf", &s, &a)
	if n != 1 {
		t.Errorf("Sscan count error: expected 1: got %d", n)
	}
	if err == nil {
		t.Errorf("Sscan expected error; got none: %s", err)
	}
	if s != "asdf" {
		t.Errorf("Sscan wrong values: got %q expected \"asdf\"", s)
	}
}

// 空字符串在扫描字符串时不是有效的输入。. md5:4be6dc662c6f1d7e
func TestScanEmpty(t *testing.T) {
	var s1, s2 string
	n, err := Sscan("abc", &s1, &s2)
	if n != 1 {
		t.Errorf("Sscan count error: expected 1: got %d", n)
	}
	if err == nil {
		t.Error("Sscan <one item> expected error; got none")
	}
	if s1 != "abc" {
		t.Errorf("Sscan wrong values: got %q expected \"abc\"", s1)
	}
	n, err = Sscan("", &s1, &s2)
	if n != 0 {
		t.Errorf("Sscan count error: expected 0: got %d", n)
	}
	if err == nil {
		t.Error("Sscan <empty> expected error; got none")
	}
	// 引用的空字符串是允许的。. md5:4022effa42f66a5a
	n, err = Sscanf(`""`, "%q", &s1)
	if n != 1 {
		t.Errorf("Sscanf count error: expected 1: got %d", n)
	}
	if err != nil {
		t.Errorf("Sscanf <empty> expected no error with quoted string; got %s", err)
	}
}

func TestScanNotPointer(t *testing.T) {
	r := strings.NewReader("1")
	var a int
	_, err := Fscan(r, a)
	if err == nil {
		t.Error("expected error scanning non-pointer")
	} else if !strings.Contains(err.Error(), "pointer") {
		t.Errorf("expected pointer error scanning non-pointer, got: %s", err)
	}
}

func TestScanlnNoNewline(t *testing.T) {
	var a int
	_, err := Sscanln("1 x\n", &a)
	if err == nil {
		t.Error("expected error scanning string missing newline")
	} else if !strings.Contains(err.Error(), "newline") {
		t.Errorf("expected newline error scanning string missing newline, got: %s", err)
	}
}

func TestScanlnWithMiddleNewline(t *testing.T) {
	r := strings.NewReader("123\n456\n")
	var a, b int
	_, err := Fscanln(r, &a, &b)
	if err == nil {
		t.Error("expected error scanning string with extra newline")
	} else if !strings.Contains(err.Error(), "newline") {
		t.Errorf("expected newline error scanning string with extra newline, got: %s", err)
	}
}

// eofCounter是一个特殊的Reader，它会在文件结尾处计数读取操作。. md5:afde126ad4be1c96
type eofCounter struct {
	reader   *strings.Reader
	eofCount int
}

func (ec *eofCounter) Read(b []byte) (n int, err error) {
	n, err = ec.reader.Read(b)
	if n == 0 {
		ec.eofCount++
	}
	return
}

// TestEOF 用于验证在扫描时，我们最多会在每次调用 Scan 函数时遇到一次 EOF，并且仅在其确实表示文件结束时出现。
// md5:26e30ed888e1cee7
func TestEOF(t *testing.T) {
	ec := &eofCounter{strings.NewReader("123\n"), 0}
	var a int
	n, err := Fscanln(ec, &a)
	if err != nil {
		t.Error("unexpected error", err)
	}
	if n != 1 {
		t.Error("expected to scan one item, got", n)
	}
	if ec.eofCount != 0 {
		t.Error("expected zero EOFs", ec.eofCount)
		ec.eofCount = 0 // reset for next test
	}
	n, err = Fscanln(ec, &a)
	if err == nil {
		t.Error("expected error scanning empty string")
	}
	if n != 0 {
		t.Error("expected to scan zero items, got", n)
	}
	if ec.eofCount != 1 {
		t.Error("expected one EOF, got", ec.eofCount)
	}
}

// TestEOFAtEndOfInput 验证在输入结束时，我们是否会看到EOF错误。
// 这曾是一个小bug：我们以前会收到"预期的整数"错误。
// md5:d611ff7fcfef8d80
func TestEOFAtEndOfInput(t *testing.T) {
	var i, j int
	n, err := Sscanf("23", "%d %d", &i, &j)
	if n != 1 || i != 23 {
		t.Errorf("Sscanf expected one value of 23; got %d %d", n, i)
	}
	if err != io.EOF {
		t.Errorf("Sscanf expected EOF; got %q", err)
	}
	n, err = Sscan("234", &i, &j)
	if n != 1 || i != 234 {
		t.Errorf("Sscan expected one value of 234; got %d %d", n, i)
	}
	if err != io.EOF {
		t.Errorf("Sscan expected EOF; got %q", err)
	}
	// 但是尾随空格就难了。. md5:d86bbb863034984b
	n, err = Sscan("234 ", &i, &j)
	if n != 1 || i != 234 {
		t.Errorf("Sscan expected one value of 234; got %d %d", n, i)
	}
	if err != io.EOF {
		t.Errorf("Sscan expected EOF; got %q", err)
	}
}

var eofTests = []struct {
	format string
	v      any
}{
	{"%s", &stringVal},
	{"%q", &stringVal},
	{"%x", &stringVal},
	{"%v", &stringVal},
	{"%v", &bytesVal},
	{"%v", &intVal},
	{"%v", &uintVal},
	{"%v", &boolVal},
	{"%v", &float32Val},
	{"%v", &complex64Val},
	{"%v", &renamedStringVal},
	{"%v", &renamedBytesVal},
	{"%v", &renamedIntVal},
	{"%v", &renamedUintVal},
	{"%v", &renamedBoolVal},
	{"%v", &renamedFloat32Val},
	{"%v", &renamedComplex64Val},
}

func TestEOFAllTypes(t *testing.T) {
	for i, test := range eofTests {
		if _, err := Sscanf("", test.format, test.v); err != io.EOF {
			t.Errorf("#%d: %s %T not eof on empty string: %s", i, test.format, test.v, err)
		}
		if _, err := Sscanf("   ", test.format, test.v); err != io.EOF {
			t.Errorf("#%d: %s %T not eof on trailing blanks: %s", i, test.format, test.v, err)
		}
	}
}

// TestUnreadRuneWithBufio 验证，至少在使用 bufio 时，连续调用 Fscan 不会丢失 runes。
// md5:ed3d0e5212b96ec6
func TestUnreadRuneWithBufio(t *testing.T) {
	r := bufio.NewReader(strings.NewReader("123αb"))
	var i int
	var a string
	n, err := Fscanf(r, "%d", &i)
	if n != 1 || err != nil {
		t.Errorf("reading int expected one item, no errors; got %d %q", n, err)
	}
	if i != 123 {
		t.Errorf("expected 123; got %d", i)
	}
	n, err = Fscanf(r, "%s", &a)
	if n != 1 || err != nil {
		t.Errorf("reading string expected one item, no errors; got %d %q", n, err)
	}
	if a != "αb" {
		t.Errorf("expected αb; got %q", a)
	}
}

type TwoLines string

// Scan 尝试读取两行内容到对象中。Scanln 应该阻止这种情况发生，因为它在遇到换行符时就会停止；而 Scan 和 Scanf 应该可以正常处理。
// md5:25c3fd3aa3eb80ea
//func (t *TwoLines) Scan(state ScanState, verb rune) error {
//	chars := make([]rune, 0, 100)
//	for nlCount := 0; nlCount < 2; {
//		c, _, err := state.ReadRune()
//		if err != nil {
//			return err
//		}
//		chars = append(chars, c)
//		if c == '\n' {
//			nlCount++
//		}
//	}
//	*t = TwoLines(string(chars))
//	return nil
//}

func TestMultiLine(t *testing.T) {
	input := "abc\ndef\n"
	// Sscan should work
	var tscan TwoLines
	n, err := Sscan(input, &tscan)
	if n != 1 {
		t.Errorf("Sscan: expected 1 item; got %d", n)
	}
	if err != nil {
		t.Errorf("Sscan: expected no error; got %s", err)
	}
	if string(tscan) != input {
		t.Errorf("Sscan: expected %q; got %q", input, tscan)
	}
	// Sscanf should work
	var tscanf TwoLines
	n, err = Sscanf(input, "%s", &tscanf)
	if n != 1 {
		t.Errorf("Sscanf: expected 1 item; got %d", n)
	}
	if err != nil {
		t.Errorf("Sscanf: expected no error; got %s", err)
	}
	if string(tscanf) != input {
		t.Errorf("Sscanf: expected %q; got %q", input, tscanf)
	}
	// Sscanln 应该无法正常工作. md5:0d88ad5fc8e737c7
	var tscanln TwoLines
	n, err = Sscanln(input, &tscanln)
	if n != 0 {
		t.Errorf("Sscanln: expected 0 items; got %d: %q", n, tscanln)
	}
	if err == nil {
		t.Error("Sscanln: expected error; got none")
	} else if err != io.ErrUnexpectedEOF {
		t.Errorf("Sscanln: expected io.ErrUnexpectedEOF (ha!); got %s", err)
	}
}

// TestLineByLineFscanf 测试Fscanf是否不会读取到换行符之后。问题号为3481。
// md5:c40bd9c1ad005957
func TestLineByLineFscanf(t *testing.T) {
	r := struct{ io.Reader }{strings.NewReader("1\n2\n")}
	var i, j int
	n, err := Fscanf(r, "%v\n", &i)
	if n != 1 || err != nil {
		t.Fatalf("first read: %d %q", n, err)
	}
	n, err = Fscanf(r, "%v\n", &j)
	if n != 1 || err != nil {
		t.Fatalf("second read: %d %q", n, err)
	}
	if i != 1 || j != 2 {
		t.Errorf("wrong values; wanted 1 2 got %d %d", i, j)
	}
}

// TestScanStateCount 验证返回的字节计数是否正确。问题 8512。. md5:854ee5084d21f394

// runeScanner 为 TestScanStateCount 实现了 Scanner 接口。. md5:0a264cc31ecdae6d
type runeScanner struct {
	rune rune
	size int
}

func (rs *runeScanner) Scan(state ScanState, verb rune) error {
	r, size, err := state.F.ReadRune()
	rs.rune = r
	rs.size = size
	return err
}

func TestScanStateCount(t *testing.T) {
	var a, b, c runeScanner
	n, err := Sscanf("12➂", "%c%c%c", &a, &b, &c)
	if err != nil {
		t.Fatal(err)
	}
	if n != 3 {
		t.Fatalf("expected 3 items consumed, got %d", n)
	}
	if a.rune != '1' || b.rune != '2' || c.rune != '➂' {
		t.Errorf("bad scan rune: %q %q %q should be '1' '2' '➂'", a.rune, b.rune, c.rune)
	}
	if a.size != 1 || b.size != 1 || c.size != 3 {
		t.Errorf("bad scan size: %q %q %q should be 1 1 3", a.size, b.size, c.size)
	}
}

// RecursiveInt 接受一个匹配 %d.%d.%d.... 格式的字符串，
// 并将其解析为一个链表。
// 这使得我们能够基准测试递归下降类型的扫描器。
// md5:7b0a5bb167cf0a5d
type RecursiveInt struct {
	i    int
	next *RecursiveInt
}

func (r *RecursiveInt) Scan(state fmt2.ScanState, verb rune) (err error) {
	_, err = Fscan(state, &r.i)
	if err != nil {
		return
	}
	next := new(RecursiveInt)
	_, err = Fscanf(state, ".%v", next)
	if err != nil {
		if err == io.ErrUnexpectedEOF {
			err = nil
		}
		return
	}
	r.next = next
	return
}

// scanInts 执行与 RecursiveInt.Scan 相同的扫描任务，但是不通过递归调用 scanner，这样我们可以更直接地比较性能。
// md5:3d178b4ed9a67c5e
func scanInts(r *RecursiveInt, b *bytes.Buffer) (err error) {
	r.next = nil
	_, err = Fscan(b, &r.i)
	if err != nil {
		return
	}
	c, _, err := b.ReadRune()
	if err != nil {
		if err == io.EOF {
			err = nil
		}
		return
	}
	if c != '.' {
		return
	}
	next := new(RecursiveInt)
	err = scanInts(next, b)
	if err == nil {
		r.next = next
	}
	return
}

func makeInts(n int) []byte {
	var buf bytes.Buffer
	Fprintf(&buf, "1")
	for i := 1; i < n; i++ {
		Fprintf(&buf, ".%d", i+1)
	}
	return buf.Bytes()
}

func TestScanInts(t *testing.T) {
	testScanInts(t, scanInts)
	testScanInts(t, func(r *RecursiveInt, b *bytes.Buffer) (err error) {
		_, err = Fscan(b, r)
		return
	})
}

// 对于不支持分段栈的gccgo平台，800足够小，不会溢出栈。
// md5:e811e59fb696d96b
const intCount = 800

func testScanInts(t *testing.T, scan func(*RecursiveInt, *bytes.Buffer) error) {
	r := new(RecursiveInt)
	ints := makeInts(intCount)
	buf := bytes.NewBuffer(ints)
	err := scan(r, buf)
	if err != nil {
		t.Error("unexpected error", err)
	}
	i := 1
	for ; r != nil; r = r.next {
		if r.i != i {
			t.Fatalf("bad scan: expected %d got %d", i, r.i)
		}
		i++
	}
	if i-1 != intCount {
		t.Fatalf("bad scan count: expected %d got %d", intCount, i-1)
	}
}

func BenchmarkScanInts(b *testing.B) {
	b.StopTimer()
	ints := makeInts(intCount)
	var r RecursiveInt
	for i := 0; i < b.N; i++ {
		buf := bytes.NewBuffer(ints)
		b.StartTimer()
		scanInts(&r, buf)
		b.StopTimer()
	}
}

func BenchmarkScanRecursiveInt(b *testing.B) {
	b.StopTimer()
	ints := makeInts(intCount)
	var r RecursiveInt
	for i := 0; i < b.N; i++ {
		buf := bytes.NewBuffer(ints)
		b.StartTimer()
		Fscan(buf, &r)
		b.StopTimer()
	}
}

func BenchmarkScanRecursiveIntReaderWrapper(b *testing.B) {
	b.StopTimer()
	ints := makeInts(intCount)
	var r RecursiveInt
	for i := 0; i < b.N; i++ {
		buf := struct{ io.Reader }{strings.NewReader(string(ints))}
		b.StartTimer()
		Fscan(buf, &r)
		b.StopTimer()
	}
}

// Issue 9124.
// 对于字节切片，%x 在遇到非空格字节终止扫描时无法正确处理。
// md5:c5a7f6f71b59296a
func TestHexBytes(t *testing.T) {
	var a, b []byte
	n, err := Sscanf("00010203", "%x", &a)
	if n != 1 || err != nil {
		t.Errorf("simple: got count, err = %d, %v; expected 1, nil", n, err)
	}
	check := func(msg string, x []byte) {
		if len(x) != 4 {
			t.Errorf("%s: bad length %d", msg, len(x))
		}
		for i, b := range x {
			if int(b) != i {
				t.Errorf("%s: bad x[%d] = %x", msg, i, x[i])
			}
		}
	}
	check("simple", a)
	a = nil

	n, err = Sscanf("00010203 00010203", "%x %x", &a, &b)
	if n != 2 || err != nil {
		t.Errorf("simple pair: got count, err = %d, %v; expected 2, nil", n, err)
	}
	check("simple pair a", a)
	check("simple pair b", b)
	a = nil
	b = nil

	n, err = Sscanf("00010203:", "%x", &a)
	if n != 1 || err != nil {
		t.Errorf("colon: got count, err = %d, %v; expected 1, nil", n, err)
	}
	check("colon", a)
	a = nil

	n, err = Sscanf("00010203:00010203", "%x:%x", &a, &b)
	if n != 2 || err != nil {
		t.Errorf("colon pair: got count, err = %d, %v; expected 2, nil", n, err)
	}
	check("colon pair a", a)
	check("colon pair b", b)
	a = nil
	b = nil

	// 这个例子失败是因为数据后面有一个十六进制字节，
	// 即输入的十六进制字节数为奇数。
	// md5:6bf44a19b1349df6
	n, err = Sscanf("000102034:", "%x", &a)
	if n != 0 || err == nil {
		t.Errorf("odd count: got count, err = %d, %v; expected 0, error", n, err)
	}
}

func TestScanNewlinesAreSpaces(t *testing.T) {
	var a, b int
	var tests = []struct {
		name  string
		text  string
		count int
	}{
		{"newlines", "1\n2\n", 2},
		{"no final newline", "1\n2", 2},
		{"newlines with spaces ", "1  \n  2  \n", 2},
		{"no final newline with spaces", "1  \n  2", 2},
	}
	for _, test := range tests {
		n, err := Sscan(test.text, &a, &b)
		if n != test.count {
			t.Errorf("%s: expected to scan %d item(s), scanned %d", test.name, test.count, n)
		}
		if err != nil {
			t.Errorf("%s: unexpected error: %s", test.name, err)
		}
	}
}

func TestScanlnNewlinesTerminate(t *testing.T) {
	var a, b int
	var tests = []struct {
		name  string
		text  string
		count int
		ok    bool
	}{
		{"one line one item", "1\n", 1, false},
		{"one line two items with spaces ", "   1 2    \n", 2, true},
		{"one line two items no newline", "   1 2", 2, true},
		{"two lines two items", "1\n2\n", 1, false},
	}
	for _, test := range tests {
		n, err := Sscanln(test.text, &a, &b)
		if n != test.count {
			t.Errorf("%s: expected to scan %d item(s), scanned %d", test.name, test.count, n)
		}
		if test.ok && err != nil {
			t.Errorf("%s: unexpected error: %s", test.name, err)
		}
		if !test.ok && err == nil {
			t.Errorf("%s: expected error; got none", test.name)
		}
	}
}

func TestScanfNewlineMatchFormat(t *testing.T) {
	var a, b int
	var tests = []struct {
		name   string
		text   string
		format string
		count  int
		ok     bool
	}{
		{"newline in both", "1\n2", "%d\n%d\n", 2, true},
		{"newline in input", "1\n2", "%d %d", 1, false},
		{"space-newline in input", "1 \n2", "%d %d", 1, false},
		{"newline in format", "1 2", "%d\n%d", 1, false},
		{"space-newline in format", "1 2", "%d \n%d", 1, false},
		{"space-newline in both", "1 \n2", "%d \n%d", 2, true},
		{"extra space in format", "1\n2", "%d\n %d", 2, true},
		{"two extra spaces in format", "1\n2", "%d \n %d", 2, true},
		{"space vs newline 0000", "1\n2", "%d\n%d", 2, true},
		{"space vs newline 0001", "1\n2", "%d\n %d", 2, true},
		{"space vs newline 0010", "1\n2", "%d \n%d", 2, true},
		{"space vs newline 0011", "1\n2", "%d \n %d", 2, true},
		{"space vs newline 0100", "1\n 2", "%d\n%d", 2, true},
		{"space vs newline 0101", "1\n 2", "%d\n%d ", 2, true},
		{"space vs newline 0110", "1\n 2", "%d \n%d", 2, true},
		{"space vs newline 0111", "1\n 2", "%d \n %d", 2, true},
		{"space vs newline 1000", "1 \n2", "%d\n%d", 2, true},
		{"space vs newline 1001", "1 \n2", "%d\n %d", 2, true},
		{"space vs newline 1010", "1 \n2", "%d \n%d", 2, true},
		{"space vs newline 1011", "1 \n2", "%d \n %d", 2, true},
		{"space vs newline 1100", "1 \n 2", "%d\n%d", 2, true},
		{"space vs newline 1101", "1 \n 2", "%d\n %d", 2, true},
		{"space vs newline 1110", "1 \n 2", "%d \n%d", 2, true},
		{"space vs newline 1111", "1 \n 2", "%d \n %d", 2, true},
		{"space vs newline no-percent 0000", "1\n2", "1\n2", 0, true},
		{"space vs newline no-percent 0001", "1\n2", "1\n 2", 0, true},
		{"space vs newline no-percent 0010", "1\n2", "1 \n2", 0, true},
		{"space vs newline no-percent 0011", "1\n2", "1 \n 2", 0, true},
		{"space vs newline no-percent 0100", "1\n 2", "1\n2", 0, false},  // 失败：输入中有换行符后跟空格，但模式中没有. md5:4ede794ccbc805f6
		{"space vs newline no-percent 0101", "1\n 2", "1\n2 ", 0, false}, // 失败：输入中有换行符后跟空格，但模式中没有. md5:4ede794ccbc805f6
		{"space vs newline no-percent 0110", "1\n 2", "1 \n2", 0, false}, // 失败：输入中有换行符后跟空格，但模式中没有. md5:4ede794ccbc805f6
		{"space vs newline no-percent 0111", "1\n 2", "1 \n 2", 0, true},
		{"space vs newline no-percent 1000", "1 \n2", "1\n2", 0, true},
		{"space vs newline no-percent 1001", "1 \n2", "1\n 2", 0, true},
		{"space vs newline no-percent 1010", "1 \n2", "1 \n2", 0, true},
		{"space vs newline no-percent 1011", "1 \n2", "1 \n 2", 0, true},
		{"space vs newline no-percent 1100", "1 \n 2", "1\n2", 0, false}, // 失败：输入中有换行符后跟空格，但模式中没有. md5:4ede794ccbc805f6
		{"space vs newline no-percent 1101", "1 \n 2", "1\n 2", 0, true},
		{"space vs newline no-percent 1110", "1 \n 2", "1 \n2", 0, false}, // 失败：输入中有换行符后跟空格，但模式中没有. md5:4ede794ccbc805f6
		{"space vs newline no-percent 1111", "1 \n 2", "1 \n 2", 0, true},
	}
	for _, test := range tests {
		var n int
		var err error
		if strings.Contains(test.format, "%") {
			n, err = Sscanf(test.text, test.format, &a, &b)
		} else {
			n, err = Sscanf(test.text, test.format)
		}
		if n != test.count {
			t.Errorf("%s: expected to scan %d item(s), scanned %d", test.name, test.count, n)
		}
		if test.ok && err != nil {
			t.Errorf("%s: unexpected error: %s", test.name, err)
		}
		if !test.ok && err == nil {
			t.Errorf("%s: expected error; got none", test.name)
		}
	}
}

// 测试问题12090：在EOF时未读取，导致字节被扫描两次。. md5:01b05f5a0bdea342

type hexBytes [2]byte

func (h *hexBytes) Scan(ss fmt2.ScanState, verb rune) error {
	var b []byte
	_, err := Fscanf(ss, "%4x", &b)
	if err != nil {
		panic(err) // Really shouldn't happen.
	}
	copy((*h)[:], b)
	return err
}

func TestHexByte(t *testing.T) {
	var h hexBytes
	n, err := Sscanln("0123\n", &h)
	if err != nil {
		t.Fatal(err)
	}
	if n != 1 {
		t.Fatalf("expected 1 item; scanned %d", n)
	}
	if h[0] != 0x01 || h[1] != 0x23 {
		t.Fatalf("expected 0123 got %x", h)
	}
}
