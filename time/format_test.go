//版权所有2009年Go作者。所有权利保留。
//使用此源代码受BSD风格
//可以在LICENSE文件中找到的许可证。
// md5:2e9dc81828a3be8a

package time_test

import (
	"fmt"
	. "github.com/888go/gosdk/time"
	"strconv"
	"strings"
	"testing"
	"testing/quick"
)

var nextStdChunkTests = []string{
	"(2006)-(01)-(02)T(15):(04):(05)(Z07:00)",
	"(2006)-(01)-(02) (002) (15):(04):(05)",
	"(2006)-(01) (002) (15):(04):(05)",
	"(2006)-(002) (15):(04):(05)",
	"(2006)(002)(01) (15):(04):(05)",
	"(2006)(002)(04) (15):(04):(05)",
}

//
//func TestNextStdChunk(t *testing.T) {
//// 大多数在Parse或Format中的错误归结于辅助函数nextStdChunk（这里称为NextStdChunk）中格式块边界的精确检测问题。这个测试直接检查nextStdChunk的行为，而不需要通过Parse/Format间接进行测试。
//// md5:dc0281b4ea39a780
//
//// markChunks 函数返回每个检测到的 "格式块" 都被括起来的格式。
//// 例如，showChunks("2006-01-02") == "(2006)-(01)-(02)"。
//// md5:b527b4cab40a3b2b
//	markChunks := func(format string) string {
//// 请注意，NextStdChunk和StdChunkNames
//// 并不属于time包的公共API。
//// 它们仅在此测试中通过export_test被导出。
//// md5:48587daf5dea2dc9
//		out := ""
//		for s := format; s != ""; {
//			prefix, std, suffix := NextStdChunk(s)
//			out += prefix
//			if std > 0 {
//				out += "(" + StdChunkNames[std] + ")"
//			}
//			s = suffix
//		}
//		return out
//	}
//
//	noParens := func(r rune) rune {
//		if r == '(' || r == ')' {
//			return -1
//		}
//		return r
//	}
//
//	for _, marked := range nextStdChunkTests {
//// marked 是 markChunks 函数的预期输出。
//// 如果我们删除括号并将它通过 markChunks 函数处理，
//// 我们应该能获得原始的输入返回。
//// md5:335b1d648fb01e33
//		format := strings.Map(noParens, marked)
//		out := markChunks(format)
//		if out != marked {
//			t.Errorf("nextStdChunk parses %q as %q, want %q", format, out, marked)
//		}
//	}
//}

type TimeFormatTest struct {
	time           Time
	formattedValue string
}

var rfc3339Formats = []TimeFormatTest{
	{Date(2008, 9, 17, 20, 4, 26, 0, UTC), "2008-09-17T20:04:26Z"},
	{Date(1994, 9, 17, 20, 4, 26, 0, FixedZone("EST", -18000)), "1994-09-17T20:04:26-05:00"},
	{Date(2000, 12, 26, 1, 15, 6, 0, FixedZone("OTO", 15600)), "2000-12-26T01:15:06+04:20"},
}

func TestRFC3339Conversion(t *testing.T) {
	for _, f := range rfc3339Formats {
		if f.time.Format(RFC3339) != f.formattedValue {
			t.Error("RFC3339:")
			t.Errorf("  want=%+v", f.formattedValue)
			t.Errorf("  have=%+v", f.time.Format(RFC3339))
		}
	}
}

//
//func TestAppendInt(t *testing.T) {
//	tests := []struct {
//		in    int
//		width int
//		want  string
//	}{
//		{0, 0, "0"},
//		{0, 1, "0"},
//		{0, 2, "00"},
//		{0, 3, "000"},
//		{1, 0, "1"},
//		{1, 1, "1"},
//		{1, 2, "01"},
//		{1, 3, "001"},
//		{-1, 0, "-1"},
//		{-1, 1, "-1"},
//		{-1, 2, "-01"},
//		{-1, 3, "-001"},
//		{99, 2, "99"},
//		{100, 2, "100"},
//		{1, 4, "0001"},
//		{12, 4, "0012"},
//		{123, 4, "0123"},
//		{1234, 4, "1234"},
//		{12345, 4, "12345"},
//		{1, 5, "00001"},
//		{12, 5, "00012"},
//		{123, 5, "00123"},
//		{1234, 5, "01234"},
//		{12345, 5, "12345"},
//		{123456, 5, "123456"},
//		{0, 9, "000000000"},
//		{123, 9, "000000123"},
//		{123456, 9, "000123456"},
//		{123456789, 9, "123456789"},
//	}
//	var got []byte
//	for _, tt := range tests {
//		got = AppendInt(got[:0], tt.in, tt.width)
//		if string(got) != tt.want {
//			t.Errorf("appendInt(%d, %d) = %s, want %s", tt.in, tt.width, got, tt.want)
//		}
//	}
//}

type FormatTest struct {
	name   string
	format string
	result string
}

var formatTests = []FormatTest{
	{"ANSIC", ANSIC, "Wed Feb  4 21:00:57 2009"},
	{"UnixDate", UnixDate, "Wed Feb  4 21:00:57 PST 2009"},
	{"RubyDate", RubyDate, "Wed Feb 04 21:00:57 -0800 2009"},
	{"RFC822", RFC822, "04 Feb 09 21:00 PST"},
	{"RFC850", RFC850, "Wednesday, 04-Feb-09 21:00:57 PST"},
	{"RFC1123", RFC1123, "Wed, 04 Feb 2009 21:00:57 PST"},
	{"RFC1123Z", RFC1123Z, "Wed, 04 Feb 2009 21:00:57 -0800"},
	{"RFC3339", RFC3339, "2009-02-04T21:00:57-08:00"},
	{"RFC3339Nano", RFC3339Nano, "2009-02-04T21:00:57.0123456-08:00"},
	{"Kitchen", Kitchen, "9:00PM"},
	{"am/pm", "3pm", "9pm"},
	{"AM/PM", "3PM", "9PM"},
	{"two-digit year", "06 01 02", "09 02 04"},
	// 三字母的月份和日子后面不能跟着小写字母。. md5:cc09856776348922
	{"Janet", "Hi Janet, the Month is January", "Hi Janet, the Month is February"},
	// 时间戳，小数秒。. md5:f94b2b512d5c574e
	{"Stamp", Stamp, "Feb  4 21:00:57"},
	{"StampMilli", StampMilli, "Feb  4 21:00:57.012"},
	{"StampMicro", StampMicro, "Feb  4 21:00:57.012345"},
	{"StampNano", StampNano, "Feb  4 21:00:57.012345600"},
	{"DateTime", DateTime, "2009-02-04 21:00:57"},
	{"DateOnly", DateOnly, "2009-02-04"},
	{"TimeOnly", TimeOnly, "21:00:57"},
	{"YearDay", "Jan  2 002 __2 2", "Feb  4 035  35 4"},
	{"Year", "2006 6 06 _6 __6 ___6", "2009 6 09 _6 __6 ___6"},
	{"Month", "Jan January 1 01 _1", "Feb February 2 02 _2"},
	{"DayOfMonth", "2 02 _2 __2", "4 04  4  35"},
	{"DayOfWeek", "Mon Monday", "Wed Wednesday"},
	{"Hour", "15 3 03 _3", "21 9 09 _9"},
	{"Minute", "4 04 _4", "0 00 _0"},
	{"Second", "5 05 _5", "57 57 _57"},
}

func TestFormat(t *testing.T) {
	// 数值时间表示的是2009年2月4日星期四，太平洋标准时间（PST）下的21:00:57.012345600。. md5:f8bd7a8cb5126dc5
	time := Unix(0, 1233810057012345600)
	//	time := time2.Unix(0, 1233810057012345600)
	for _, test := range formatTests {
		result := time.Format(test.format)
		if result != test.result {
			t.Errorf("%s expected %q got %q", test.name, test.result, result)
		}
	}
}

var goStringTests = []struct {
	in   Time
	want string
}{
	{Date(2009, February, 5, 5, 0, 57, 12345600, UTC),
		"time.Date(2009, time.February, 5, 5, 0, 57, 12345600, time.UTC)"},
	{Date(2009, February, 5, 5, 0, 57, 12345600, Local),
		"time.Date(2009, time.February, 5, 5, 0, 57, 12345600, time.Local)"},
	{Date(2009, February, 5, 5, 0, 57, 12345600, FixedZone("Europe/Berlin", 3*60*60)),
		`time.Date(2009, time.February, 5, 5, 0, 57, 12345600, time.Location("Europe/Berlin"))`,
	},
	{Date(2009, February, 5, 5, 0, 57, 12345600, FixedZone("Non-ASCII character ⏰", 3*60*60)),
		`time.Date(2009, time.February, 5, 5, 0, 57, 12345600, time.Location("Non-ASCII character \xe2\x8f\xb0"))`,
	},
}

func TestGoString(t *testing.T) {
	// 数值时间表示的是2009年2月4日星期四，太平洋标准时间（PST）下的21:00:57.012345600。. md5:f8bd7a8cb5126dc5
	for _, tt := range goStringTests {
		if tt.in.GoString() != tt.want {
			t.Errorf("GoString (%q): got %q want %q", tt.in, tt.in.GoString(), tt.want)
		}
	}
}

// issue 12440.
func TestFormatSingleDigits(t *testing.T) {
	time := Date(2001, 2, 3, 4, 5, 6, 700000000, UTC)
	test := FormatTest{"single digit format", "3:4:5", "4:5:6"}
	result := time.Format(test.format)
	if result != test.result {
		t.Errorf("%s expected %q got %q", test.name, test.result, result)
	}
}

func TestFormatShortYear(t *testing.T) {
	years := []int{
		-100001, -100000, -99999,
		-10001, -10000, -9999,
		-1001, -1000, -999,
		-101, -100, -99,
		-11, -10, -9,
		-1, 0, 1,
		9, 10, 11,
		99, 100, 101,
		999, 1000, 1001,
		9999, 10000, 10001,
		99999, 100000, 100001,
	}

	for _, y := range years {
		time := Date(y, January, 1, 0, 0, 0, 0, UTC)
		result := time.Format("2006.01.02")
		var want string
		if y < 0 {
			// %04d 中的 4 包含了负号，因此打印 -y 并自定义负号。
			// md5:68725574231a99d7
			want = fmt.Sprintf("-%04d.%02d.%02d", -y, 1, 1)
		} else {
			want = fmt.Sprintf("%04d.%02d.%02d", y, 1, 1)
		}
		if result != want {
			t.Errorf("(jan 1 %d).Format(\"2006.01.02\") = %q, want %q", y, result, want)
		}
	}
}

type ParseTest struct {
	name       string
	format     string
	value      string
	hasTZ      bool // contains a time zone
	hasWD      bool // contains a weekday
	yearSign   int  // 年份的符号，-1表示年份在该格式中不存在. md5:da1830c86f6953a1
	fracDigits int  // 小数部分的位数. md5:2cc296fbcae0d331
}

var parseTests = []ParseTest{
	{"ANSIC", ANSIC, "Thu Feb  4 21:00:57 2010", false, true, 1, 0},
	{"UnixDate", UnixDate, "Thu Feb  4 21:00:57 PST 2010", true, true, 1, 0},
	{"RubyDate", RubyDate, "Thu Feb 04 21:00:57 -0800 2010", true, true, 1, 0},
	{"RFC850", RFC850, "Thursday, 04-Feb-10 21:00:57 PST", true, true, 1, 0},
	{"RFC1123", RFC1123, "Thu, 04 Feb 2010 21:00:57 PST", true, true, 1, 0},
	{"RFC1123", RFC1123, "Thu, 04 Feb 2010 22:00:57 PDT", true, true, 1, 0},
	{"RFC1123Z", RFC1123Z, "Thu, 04 Feb 2010 21:00:57 -0800", true, true, 1, 0},
	{"RFC3339", RFC3339, "2010-02-04T21:00:57-08:00", true, false, 1, 0},
	{"custom: \"2006-01-02 15:04:05-07\"", "2006-01-02 15:04:05-07", "2010-02-04 21:00:57-08", true, false, 1, 0},
	// 可选的小数秒部分。. md5:00c9ac781457753e
	{"ANSIC", ANSIC, "Thu Feb  4 21:00:57.0 2010", false, true, 1, 1},
	{"UnixDate", UnixDate, "Thu Feb  4 21:00:57.01 PST 2010", true, true, 1, 2},
	{"RubyDate", RubyDate, "Thu Feb 04 21:00:57.012 -0800 2010", true, true, 1, 3},
	{"RFC850", RFC850, "Thursday, 04-Feb-10 21:00:57.0123 PST", true, true, 1, 4},
	{"RFC1123", RFC1123, "Thu, 04 Feb 2010 21:00:57.01234 PST", true, true, 1, 5},
	{"RFC1123Z", RFC1123Z, "Thu, 04 Feb 2010 21:00:57.01234 -0800", true, true, 1, 5},
	{"RFC3339", RFC3339, "2010-02-04T21:00:57.012345678-08:00", true, false, 1, 9},
	{"custom: \"2006-01-02 15:04:05\"", "2006-01-02 15:04:05", "2010-02-04 21:00:57.0", false, false, 1, 0},
	// 空格的数量不应有所影响。. md5:a18d704701640125
	{"ANSIC", ANSIC, "Thu Feb 4 21:00:57 2010", false, true, 1, 0},
	{"ANSIC", ANSIC, "Thu      Feb     4     21:00:57     2010", false, true, 1, 0},
	// Case should not matter
	{"ANSIC", ANSIC, "THU FEB 4 21:00:57 2010", false, true, 1, 0},
	{"ANSIC", ANSIC, "thu feb 4 21:00:57 2010", false, true, 1, 0},
	// Fractional seconds.
	{"millisecond:: dot separator", "Mon Jan _2 15:04:05.000 2006", "Thu Feb  4 21:00:57.012 2010", false, true, 1, 3},
	{"microsecond:: dot separator", "Mon Jan _2 15:04:05.000000 2006", "Thu Feb  4 21:00:57.012345 2010", false, true, 1, 6},
	{"nanosecond:: dot separator", "Mon Jan _2 15:04:05.000000000 2006", "Thu Feb  4 21:00:57.012345678 2010", false, true, 1, 9},
	{"millisecond:: comma separator", "Mon Jan _2 15:04:05,000 2006", "Thu Feb  4 21:00:57.012 2010", false, true, 1, 3},
	{"microsecond:: comma separator", "Mon Jan _2 15:04:05,000000 2006", "Thu Feb  4 21:00:57.012345 2010", false, true, 1, 6},
	{"nanosecond:: comma separator", "Mon Jan _2 15:04:05,000000000 2006", "Thu Feb  4 21:00:57.012345678 2010", false, true, 1, 9},

	// 其他地方的零不应被视为小数秒。. md5:55a24aa015ab5a17
	{"zero1", "2006.01.02.15.04.05.0", "2010.02.04.21.00.57.0", false, false, 1, 1},
	{"zero2", "2006.01.02.15.04.05.00", "2010.02.04.21.00.57.01", false, false, 1, 2},
	// 月份和星期的名称只在不跟一个小写字母的情况下匹配。. md5:7e2aa3dab305db43
	{"Janet", "Hi Janet, the Month is January: Jan _2 15:04:05 2006", "Hi Janet, the Month is February: Feb  4 21:00:57 2010", false, true, 1, 0},

	// GMT with offset.
	{"GMT-8", UnixDate, "Fri Feb  5 05:00:57 GMT-8 2010", true, true, 1, 0},

	// 接受任意数量的小数秒位数（包括没有小数位），如.999...
	// 在 Go 1 中，.999... 在格式中完全被忽略，这意味着前两个情况会成功，但接下来的四个情况不会。
	// Go 1.1 则接受所有这六种情况。
	// 小数点作为分隔符。
	// md5:2f78deb3265e8291
	{"", "2006-01-02 15:04:05.9999 -0700 MST", "2010-02-04 21:00:57 -0800 PST", true, false, 1, 0},
	{"", "2006-01-02 15:04:05.999999999 -0700 MST", "2010-02-04 21:00:57 -0800 PST", true, false, 1, 0},
	{"", "2006-01-02 15:04:05.9999 -0700 MST", "2010-02-04 21:00:57.0123 -0800 PST", true, false, 1, 4},
	{"", "2006-01-02 15:04:05.999999999 -0700 MST", "2010-02-04 21:00:57.0123 -0800 PST", true, false, 1, 4},
	{"", "2006-01-02 15:04:05.9999 -0700 MST", "2010-02-04 21:00:57.012345678 -0800 PST", true, false, 1, 9},
	{"", "2006-01-02 15:04:05.999999999 -0700 MST", "2010-02-04 21:00:57.012345678 -0800 PST", true, false, 1, 9},
	// comma "," separator.
	{"", "2006-01-02 15:04:05,9999 -0700 MST", "2010-02-04 21:00:57 -0800 PST", true, false, 1, 0},
	{"", "2006-01-02 15:04:05,999999999 -0700 MST", "2010-02-04 21:00:57 -0800 PST", true, false, 1, 0},
	{"", "2006-01-02 15:04:05,9999 -0700 MST", "2010-02-04 21:00:57.0123 -0800 PST", true, false, 1, 4},
	{"", "2006-01-02 15:04:05,999999999 -0700 MST", "2010-02-04 21:00:57.0123 -0800 PST", true, false, 1, 4},
	{"", "2006-01-02 15:04:05,9999 -0700 MST", "2010-02-04 21:00:57.012345678 -0800 PST", true, false, 1, 9},
	{"", "2006-01-02 15:04:05,999999999 -0700 MST", "2010-02-04 21:00:57.012345678 -0800 PST", true, false, 1, 9},

	// issue 4502.
	{"", StampNano, "Feb  4 21:00:57.012345678", false, false, -1, 9},
	{"", "Jan _2 15:04:05.999", "Feb  4 21:00:57.012300000", false, false, -1, 4},
	{"", "Jan _2 15:04:05.999", "Feb  4 21:00:57.012345678", false, false, -1, 9},
	{"", "Jan _2 15:04:05.999999999", "Feb  4 21:00:57.0123", false, false, -1, 4},
	{"", "Jan _2 15:04:05.999999999", "Feb  4 21:00:57.012345678", false, false, -1, 9},

	// Day of year.
	{"", "2006-01-02 002 15:04:05", "2010-02-04 035 21:00:57", false, false, 1, 0},
	{"", "2006-01 002 15:04:05", "2010-02 035 21:00:57", false, false, 1, 0},
	{"", "2006-002 15:04:05", "2010-035 21:00:57", false, false, 1, 0},
	{"", "200600201 15:04:05", "201003502 21:00:57", false, false, 1, 0},
	{"", "200600204 15:04:05", "201003504 21:00:57", false, false, 1, 0},
}

func TestParse(t *testing.T) {
	for _, test := range parseTests {
		time, err := Parse(test.format, test.value)
		if err != nil {
			t.Errorf("%s error: %v", test.name, err)
		} else {
			checkTime(time, &test, t)
		}
	}
}

// All parsed with ANSIC.
var dayOutOfRangeTests = []struct {
	date string
	ok   bool
}{
	{"Thu Jan 99 21:00:57 2010", false},
	{"Thu Jan 31 21:00:57 2010", true},
	{"Thu Jan 32 21:00:57 2010", false},
	{"Thu Feb 28 21:00:57 2012", true},
	{"Thu Feb 29 21:00:57 2012", true},
	{"Thu Feb 29 21:00:57 2010", false},
	{"Thu Mar 31 21:00:57 2010", true},
	{"Thu Mar 32 21:00:57 2010", false},
	{"Thu Apr 30 21:00:57 2010", true},
	{"Thu Apr 31 21:00:57 2010", false},
	{"Thu May 31 21:00:57 2010", true},
	{"Thu May 32 21:00:57 2010", false},
	{"Thu Jun 30 21:00:57 2010", true},
	{"Thu Jun 31 21:00:57 2010", false},
	{"Thu Jul 31 21:00:57 2010", true},
	{"Thu Jul 32 21:00:57 2010", false},
	{"Thu Aug 31 21:00:57 2010", true},
	{"Thu Aug 32 21:00:57 2010", false},
	{"Thu Sep 30 21:00:57 2010", true},
	{"Thu Sep 31 21:00:57 2010", false},
	{"Thu Oct 31 21:00:57 2010", true},
	{"Thu Oct 32 21:00:57 2010", false},
	{"Thu Nov 30 21:00:57 2010", true},
	{"Thu Nov 31 21:00:57 2010", false},
	{"Thu Dec 31 21:00:57 2010", true},
	{"Thu Dec 32 21:00:57 2010", false},
	{"Thu Dec 00 21:00:57 2010", false},
}

func TestParseDayOutOfRange(t *testing.T) {
	for _, test := range dayOutOfRangeTests {
		_, err := Parse(ANSIC, test.date)
		switch {
		case test.ok && err == nil:
			// OK
		case !test.ok && err != nil:
			if !strings.Contains(err.Error(), "day out of range") {
				t.Errorf("%q: expected 'day' error, got %v", test.date, err)
			}
		case test.ok && err != nil:
			t.Errorf("%q: unexpected error: %v", test.date, err)
		case !test.ok && err == nil:
			t.Errorf("%q: expected 'day' error, got none", test.date)
		}
	}
}

// TestParseInLocation 检查 Parse 和 ParseInLocation 函数不会因AST（阿拉伯标准时间）和AST（大西洋标准时间）是不同的时区，但具有相同的缩写而产生混淆。
//
// ICANN 已经逐渐淘汰了自创的缩写，转而使用数字时区（例如，在2017a版tzdata发布中，Asia/Baghdad时区的缩写从AST更改为+03）；但我们仍然要确保在使用稍旧版本tzdata包的系统上，time包不会产生混淆。
// md5:98c0bc82994ed15e
func TestParseInLocation(t *testing.T) {

	baghdad, err := LoadLocation("Asia/Baghdad")
	if err != nil {
		t.Fatal(err)
	}

	var t1, t2 Time

	t1, err = ParseInLocation("Jan 02 2006 MST", "Feb 01 2013 AST", baghdad)
	if err != nil {
		t.Fatal(err)
	}

	_, offset := t1.Zone()

	// 一个零偏移表示ParseInLocation没有识别'AST'缩写与当前位置（巴格达，我们期望+03小时偏移）匹配；可能是因为我们使用了较新的tzdata发布（2017a或更高版本）。如果发生这种情况，跳过巴格达测试。
	// md5:faba138ea75723dc
	if offset != 0 {
		t2 = Date(2013, February, 1, 00, 00, 00, 0, baghdad)
		if t1 != t2 {
			t.Fatalf("ParseInLocation(Feb 01 2013 AST, Baghdad) = %v, want %v", t1, t2)
		}
		if offset != 3*60*60 {
			t.Fatalf("ParseInLocation(Feb 01 2013 AST, Baghdad).Zone = _, %d, want _, %d", offset, 3*60*60)
		}
	}

	blancSablon, err := LoadLocation("America/Blanc-Sablon")
	if err != nil {
		t.Fatal(err)
	}

	// 在这种情况下，'AST'表示'大西洋标准时间'，我们期望缩写能够正确匹配美国的地理位置。
	// md5:d4fb808a8e9139d9
	t1, err = ParseInLocation("Jan 02 2006 MST", "Feb 01 2013 AST", blancSablon)
	if err != nil {
		t.Fatal(err)
	}
	t2 = Date(2013, February, 1, 00, 00, 00, 0, blancSablon)
	if t1 != t2 {
		t.Fatalf("ParseInLocation(Feb 01 2013 AST, Blanc-Sablon) = %v, want %v", t1, t2)
	}
	_, offset = t1.Zone()
	if offset != -4*60*60 {
		t.Fatalf("ParseInLocation(Feb 01 2013 AST, Blanc-Sablon).Zone = _, %d, want _, %d", offset, -4*60*60)
	}
}

//func TestLoadLocationZipFile(t *testing.T) {
//	undo := DisablePlatformSources()
//	defer undo()
//
//	_, err := LoadLocation("Australia/Sydney")
//	if err != nil {
//		t.Fatal(err)
//	}
//}

var rubyTests = []ParseTest{
	{"RubyDate", RubyDate, "Thu Feb 04 21:00:57 -0800 2010", true, true, 1, 0},
	// 在测试中忽略时区。如果能够解析，那就没问题。. md5:bf2d9ad29c726751
	{"RubyDate", RubyDate, "Thu Feb 04 21:00:57 -0000 2010", false, true, 1, 0},
	{"RubyDate", RubyDate, "Thu Feb 04 21:00:57 +0000 2010", false, true, 1, 0},
	{"RubyDate", RubyDate, "Thu Feb 04 21:00:57 +1130 2010", false, true, 1, 0},
}

// 有问题的时区格式需要特别的测试。. md5:c661a6b690158d00
func TestRubyParse(t *testing.T) {
	for _, test := range rubyTests {
		time, err := Parse(test.format, test.value)
		if err != nil {
			t.Errorf("%s error: %v", test.name, err)
		} else {
			checkTime(time, &test, t)
		}
	}
}

func checkTime(time Time, test *ParseTest, t *testing.T) {
	// 时间应为2010年2月4日星期四，太平洋标准时间21:00:57. md5:8b02281dad0ed335
	if test.yearSign >= 0 && test.yearSign*time.Year() != 2010 {
		t.Errorf("%s: bad year: %d not %d", test.name, time.Year(), 2010)
	}
	if time.Month() != February {
		t.Errorf("%s: bad month: %s not %s", test.name, time.Month(), February)
	}
	if time.Day() != 4 {
		t.Errorf("%s: bad day: %d not %d", test.name, time.Day(), 4)
	}
	if time.Hour() != 21 {
		t.Errorf("%s: bad hour: %d not %d", test.name, time.Hour(), 21)
	}
	if time.Minute() != 0 {
		t.Errorf("%s: bad minute: %d not %d", test.name, time.Minute(), 0)
	}
	if time.Second() != 57 {
		t.Errorf("%s: bad second: %d not %d", test.name, time.Second(), 57)
	}
	// 必须检查纳秒是否超出输入的精度。. md5:0ad2fdf7ae989ccf
	nanosec, err := strconv.ParseUint("012345678"[:test.fracDigits]+"000000000"[:9-test.fracDigits], 10, 0)
	if err != nil {
		panic(err)
	}
	if time.Nanosecond() != int(nanosec) {
		t.Errorf("%s: bad nanosecond: %d not %d", test.name, time.Nanosecond(), nanosec)
	}
	name, offset := time.Zone()
	if test.hasTZ && offset != -28800 {
		t.Errorf("%s: bad tz offset: %s %d not %d", test.name, name, offset, -28800)
	}
	if test.hasWD && time.Weekday() != Thursday {
		t.Errorf("%s: bad weekday: %s not %s", test.name, time.Weekday(), Thursday)
	}
}

func TestFormatAndParse(t *testing.T) {
	const fmt = "Mon MST " + RFC3339 // all fields
	f := func(sec int64) bool {
		t1 := Unix(sec/2, 0)
		if t1.Year() < 1000 || t1.Year() > 9999 || t1.Unix() != sec {
			// not required to work
			return true
		}
		t2, err := Parse(fmt, t1.Format(fmt))
		if err != nil {
			t.Errorf("error: %s", err)
			return false
		}
		if t1.Unix() != t2.Unix() || t1.Nanosecond() != t2.Nanosecond() {
			t.Errorf("FormatAndParse %d: %q(%d) %q(%d)", sec, t1, t1.Unix(), t2, t2.Unix())
			return false
		}
		return true
	}
	f32 := func(sec int32) bool { return f(int64(sec)) }
	cfg := &quick.Config{MaxCount: 10000}

	// 首先尝试一个合理的日期，然后尝试较大的日期。. md5:379431a5c1a862c0
	if err := quick.Check(f32, cfg); err != nil {
		t.Fatal(err)
	}
	if err := quick.Check(f, cfg); err != nil {
		t.Fatal(err)
	}
}

type ParseTimeZoneTest struct {
	value  string
	length int
	ok     bool
}

var parseTimeZoneTests = []ParseTimeZoneTest{
	{"gmt hi there", 0, false},
	{"GMT hi there", 3, true},
	{"GMT+12 hi there", 6, true},
	{"GMT+00 hi there", 6, true},
	{"GMT+", 3, true},
	{"GMT+3", 5, true},
	{"GMT+a", 3, true},
	{"GMT+3a", 5, true},
	{"GMT-5 hi there", 5, true},
	{"GMT-51 hi there", 3, true},
	{"ChST hi there", 4, true},
	{"MeST hi there", 4, true},
	{"MSDx", 3, true},
	{"MSDY", 0, false}, // 四个字母的单词必须以"T"结尾。. md5:cd05e32109e09796
	{"ESAST hi", 5, true},
	{"ESASTT hi", 0, false}, // 连续的大写字母过长。. md5:d02440bd8cae8bd3
	{"ESATY hi", 0, false},  // 五个字母必须以T结尾。. md5:46f84208446da6fc
	{"WITA hi", 4, true},    // Issue #18251
	// Issue #24071
	{"+03 hi", 3, true},
	{"-04 hi", 3, true},
	// Issue #26032
	{"+00", 3, true},
	{"-11", 3, true},
	{"-12", 3, true},
	{"-23", 3, true},
	{"-24", 0, false},
	{"+13", 3, true},
	{"+14", 3, true},
	{"+23", 3, true},
	{"+24", 0, false},
}

//func TestParseTimeZone(t *testing.T) {
//	for _, test := range parseTimeZoneTests {
//		length, ok := ParseTimeZone(test.value)
//		if ok != test.ok {
//			t.Errorf("expected %t for %q got %t", test.ok, test.value, ok)
//		} else if length != test.length {
//			t.Errorf("expected %d for %q got %d", test.length, test.value, length)
//		}
//	}
//}

type ParseErrorTest struct {
	format string
	value  string
	expect string // 必须出现在错误中. md5:aa43882c413d121f
}

var parseErrorTests = []ParseErrorTest{
	{ANSIC, "Feb  4 21:00:60 2010", `cannot parse "Feb  4 21:00:60 2010" as "Mon"`},
	{ANSIC, "Thu Feb  4 21:00:57 @2010", `cannot parse "@2010" as "2006"`},
	{ANSIC, "Thu Feb  4 21:00:60 2010", "second out of range"},
	{ANSIC, "Thu Feb  4 21:61:57 2010", "minute out of range"},
	{ANSIC, "Thu Feb  4 24:00:60 2010", "hour out of range"},
	{"Mon Jan _2 15:04:05.000 2006", "Thu Feb  4 23:00:59x01 2010", `cannot parse "x01 2010" as ".000"`},
	{"Mon Jan _2 15:04:05.000 2006", "Thu Feb  4 23:00:59.xxx 2010", `cannot parse ".xxx 2010" as ".000"`},
	{"Mon Jan _2 15:04:05.000 2006", "Thu Feb  4 23:00:59.-123 2010", "fractional second out of range"},
	// 问题4502。StampNano需要精确到9位数字。. md5:2dfabdb6247ef22f
	{StampNano, "Dec  7 11:22:01.000000", `cannot parse ".000000" as ".000000000"`},
	{StampNano, "Dec  7 11:22:01.0000000000", `extra text: "0"`},
	// issue 4493. 有用的错误信息。. md5:78fb3e19d3500b81
	{RFC3339, "2006-01-02T15:04:05Z07:00", `parsing time "2006-01-02T15:04:05Z07:00": extra text: "07:00"`},
	{RFC3339, "2006-01-02T15:04_abc", `parsing time "2006-01-02T15:04_abc" as "2006-01-02T15:04:05Z07:00": cannot parse "_abc" as ":"`},
	{RFC3339, "2006-01-02T15:04:05_abc", `parsing time "2006-01-02T15:04:05_abc" as "2006-01-02T15:04:05Z07:00": cannot parse "_abc" as "Z07:00"`},
	{RFC3339, "2006-01-02T15:04:05Z_abc", `parsing time "2006-01-02T15:04:05Z_abc": extra text: "_abc"`},
	// 非法的第二个部分，后面跟着可选的小数秒. md5:993f99cf63c16ad9
	{RFC3339, "2010-02-04T21:00:67.012345678-08:00", "second out of range"},
	// issue 54569
	{RFC3339, "0000-01-01T00:00:.0+00:00", `parsing time "0000-01-01T00:00:.0+00:00" as "2006-01-02T15:04:05Z07:00": cannot parse ".0+00:00" as "05"`},
	// issue 21113
	{"_2 Jan 06 15:04 MST", "4 --- 00 00:00 GMT", `cannot parse "--- 00 00:00 GMT" as "Jan"`},
	{"_2 January 06 15:04 MST", "4 --- 00 00:00 GMT", `cannot parse "--- 00 00:00 GMT" as "January"`},

	// 无效或不匹配的年中天数. md5:bae5796035fb1b20
	{"Jan _2 002 2006", "Feb  4 034 2006", "day-of-year does not match day"},
	{"Jan _2 002 2006", "Feb  4 004 2006", "day-of-year does not match month"},

	// issue 45391.
	{`"2006-01-02T15:04:05Z07:00"`, "0", `parsing time "0" as "\"2006-01-02T15:04:05Z07:00\"": cannot parse "0" as "\""`},
	{RFC3339, "\"", `parsing time "\"" as "2006-01-02T15:04:05Z07:00": cannot parse "\"" as "2006"`},

	// issue 54570
	{RFC3339, "0000-01-01T00:00:00+00:+0", `parsing time "0000-01-01T00:00:00+00:+0" as "2006-01-02T15:04:05Z07:00": cannot parse "+00:+0" as "Z07:00"`},
	{RFC3339, "0000-01-01T00:00:00+-0:00", `parsing time "0000-01-01T00:00:00+-0:00" as "2006-01-02T15:04:05Z07:00": cannot parse "+-0:00" as "Z07:00"`},

	// issue 56730
	{"2006-01-02", "22-10-25", `parsing time "22-10-25" as "2006-01-02": cannot parse "22-10-25" as "2006"`},
	{"06-01-02", "a2-10-25", `parsing time "a2-10-25" as "06-01-02": cannot parse "a2-10-25" as "06"`},
	{"03:04PM", "12:03pM", `parsing time "12:03pM" as "03:04PM": cannot parse "pM" as "PM"`},
	{"03:04pm", "12:03pM", `parsing time "12:03pM" as "03:04pm": cannot parse "pM" as "pm"`},
}

func TestParseErrors(t *testing.T) {
	for _, test := range parseErrorTests {
		_, err := Parse(test.format, test.value)
		if err == nil {
			t.Errorf("expected error for %q %q", test.format, test.value)
		} else if !strings.Contains(err.Error(), test.expect) {
			t.Errorf("expected error with %q for %q %q; got %s", test.expect, test.format, test.value, err)
		}
	}
}

func TestNoonIs12PM(t *testing.T) {
	noon := Date(0, January, 1, 12, 0, 0, 0, UTC)
	const expect = "12:00PM"
	got := noon.Format("3:04PM")
	if got != expect {
		t.Errorf("got %q; expect %q", got, expect)
	}
	got = noon.Format("03:04PM")
	if got != expect {
		t.Errorf("got %q; expect %q", got, expect)
	}
}

func TestMidnightIs12AM(t *testing.T) {
	midnight := Date(0, January, 1, 0, 0, 0, 0, UTC)
	expect := "12:00AM"
	got := midnight.Format("3:04PM")
	if got != expect {
		t.Errorf("got %q; expect %q", got, expect)
	}
	got = midnight.Format("03:04PM")
	if got != expect {
		t.Errorf("got %q; expect %q", got, expect)
	}
}

func Test12PMIsNoon(t *testing.T) {
	noon, err := Parse("3:04PM", "12:00PM")
	if err != nil {
		t.Fatal("error parsing date:", err)
	}
	if noon.Hour() != 12 {
		t.Errorf("got %d; expect 12", noon.Hour())
	}
	noon, err = Parse("03:04PM", "12:00PM")
	if err != nil {
		t.Fatal("error parsing date:", err)
	}
	if noon.Hour() != 12 {
		t.Errorf("got %d; expect 12", noon.Hour())
	}
}

func Test12AMIsMidnight(t *testing.T) {
	midnight, err := Parse("3:04PM", "12:00AM")
	if err != nil {
		t.Fatal("error parsing date:", err)
	}
	if midnight.Hour() != 0 {
		t.Errorf("got %d; expect 0", midnight.Hour())
	}
	midnight, err = Parse("03:04PM", "12:00AM")
	if err != nil {
		t.Fatal("error parsing date:", err)
	}
	if midnight.Hour() != 0 {
		t.Errorf("got %d; expect 0", midnight.Hour())
	}
}

// 检查一个没有时区的时间在使用MST作为请求时区格式化时，仍然能生成（数字表示的）时区。
// md5:9de8eb9b9dd546d6
func TestMissingZone(t *testing.T) {
	time, err := Parse(RubyDate, "Thu Feb 02 16:10:03 -0500 2006")
	if err != nil {
		t.Fatal("error parsing date:", err)
	}
	expect := "Thu Feb  2 16:10:03 -0500 2006" // -0500 not EST
	str := time.Format(UnixDate)               // 使用MST作为其时区. md5:8743d7a99d9ce6e6
	if str != expect {
		t.Errorf("got %s; expect %s", str, expect)
	}
}

func TestMinutesInTimeZone(t *testing.T) {
	time, err := Parse(RubyDate, "Mon Jan 02 15:04:05 +0123 2006")
	if err != nil {
		t.Fatal("error parsing date:", err)
	}
	expected := (1*60 + 23) * 60
	_, offset := time.Zone()
	if offset != expected {
		t.Errorf("ZoneOffset = %d, want %d", offset, expected)
	}
}

type SecondsTimeZoneOffsetTest struct {
	format         string
	value          string
	expectedoffset int
}

var secondsTimeZoneOffsetTests = []SecondsTimeZoneOffsetTest{
	{"2006-01-02T15:04:05-070000", "1871-01-01T05:33:02-003408", -(34*60 + 8)},
	{"2006-01-02T15:04:05-07:00:00", "1871-01-01T05:33:02-00:34:08", -(34*60 + 8)},
	{"2006-01-02T15:04:05-070000", "1871-01-01T05:33:02+003408", 34*60 + 8},
	{"2006-01-02T15:04:05-07:00:00", "1871-01-01T05:33:02+00:34:08", 34*60 + 8},
	{"2006-01-02T15:04:05Z070000", "1871-01-01T05:33:02-003408", -(34*60 + 8)},
	{"2006-01-02T15:04:05Z07:00:00", "1871-01-01T05:33:02+00:34:08", 34*60 + 8},
	{"2006-01-02T15:04:05-07", "1871-01-01T05:33:02+01", 1 * 60 * 60},
	{"2006-01-02T15:04:05-07", "1871-01-01T05:33:02-02", -2 * 60 * 60},
	{"2006-01-02T15:04:05Z07", "1871-01-01T05:33:02-02", -2 * 60 * 60},
}

func TestParseSecondsInTimeZone(t *testing.T) {
	// 应该接受秒级别的时区偏移，例如：时区 America/New_York -4:56:02 -    LMT    1883年11月18日12:03:58. md5:ed4076e1b8985c66
	for _, test := range secondsTimeZoneOffsetTests {
		time, err := Parse(test.format, test.value)
		if err != nil {
			t.Fatal("error parsing date:", err)
		}
		_, offset := time.Zone()
		if offset != test.expectedoffset {
			t.Errorf("ZoneOffset = %d, want %d", offset, test.expectedoffset)
		}
	}
}

func TestFormatSecondsInTimeZone(t *testing.T) {
	for _, test := range secondsTimeZoneOffsetTests {
		d := Date(1871, 1, 1, 5, 33, 2, 0, FixedZone("LMT", test.expectedoffset))
		timestr := d.Format(test.format)
		if timestr != test.value {
			t.Errorf("Format = %s, want %s", timestr, test.value)
		}
	}
}

// Issue 11334.
func TestUnderscoreTwoThousand(t *testing.T) {
	format := "15:04_20060102"
	input := "14:38_20150618"
	time, err := Parse(format, input)
	if err != nil {
		t.Error(err)
	}
	if y, m, d := time.Date(); y != 2015 || m != 6 || d != 18 {
		t.Errorf("Incorrect y/m/d, got %d/%d/%d", y, m, d)
	}
	if h := time.Hour(); h != 14 {
		t.Errorf("Incorrect hour, got %d", h)
	}
	if m := time.Minute(); m != 38 {
		t.Errorf("Incorrect minute, got %d", m)
	}
}

// Issue 29918, 29916
func TestStd0xParseError(t *testing.T) {
	tests := []struct {
		format, value, valueElemPrefix string
	}{
		{"01 MST", "0 MST", "0"},
		{"01 MST", "1 MST", "1"},
		{RFC850, "Thursday, 04-Feb-1 21:00:57 PST", "1"},
	}
	for _, tt := range tests {
		_, err := Parse(tt.format, tt.value)
		if err == nil {
			t.Errorf("Parse(%q, %q) did not fail as expected", tt.format, tt.value)
		} else if perr, ok := err.(*ParseError); !ok {
			t.Errorf("Parse(%q, %q) returned error type %T, expected ParseError", tt.format, tt.value, perr)
		} else if !strings.Contains(perr.Error(), "cannot parse") || !strings.HasPrefix(perr.F.ValueElem, tt.valueElemPrefix) {
			t.Errorf("Parse(%q, %q) returned wrong parsing error message: %v", tt.format, tt.value, perr)
		}
	}
}

var monthOutOfRangeTests = []struct {
	value string
	ok    bool
}{
	{"00-01", false},
	{"13-01", false},
	{"01-01", true},
}

func TestParseMonthOutOfRange(t *testing.T) {
	for _, test := range monthOutOfRangeTests {
		_, err := Parse("01-02", test.value)
		switch {
		case !test.ok && err != nil:
			if !strings.Contains(err.Error(), "month out of range") {
				t.Errorf("%q: expected 'month' error, got %v", test.value, err)
			}
		case test.ok && err != nil:
			t.Errorf("%q: unexpected error: %v", test.value, err)
		case !test.ok && err == nil:
			t.Errorf("%q: expected 'month' error, got none", test.value)
		}
	}
}

//// Issue 37387.
//func TestParseYday(t *testing.T) {
//	t.Parallel()
//	for i := 1; i <= 365; i++ {
//		d := fmt.Sprintf("2020-%03d", i)
//		tm, err := Parse("2006-002", d)
//		if err != nil {
//			t.Errorf("unexpected error for %s: %v", d, err)
//		} else if tm.Year() != 2020 || tm.YearDay() != i {
//			t.Errorf("got year %d yearday %d, want %d %d", tm.Year(), tm.YearDay(), 2020, i)
//		}
//	}
//}

//
//// Issue 45391.
//func TestQuote(t *testing.T) {
//	tests := []struct {
//		s, want string
//	}{
//		{`"`, `"\""`},
//		{`abc"xyz"`, `"abc\"xyz\""`},
//		{"", `""`},
//		{"abc", `"abc"`},
//		{`☺`, `"\xe2\x98\xba"`},
//		{`☺ hello ☺ hello`, `"\xe2\x98\xba hello \xe2\x98\xba hello"`},
//		{"\x04", `"\x04"`},
//	}
//	for _, tt := range tests {
//		if q := Quote(tt.s); q != tt.want {
//			t.Errorf("Quote(%q) = got %q, want %q", tt.s, q, tt.want)
//		}
//	}
//
//}

// Issue 48037
func TestFormatFractionalSecondSeparators(t *testing.T) {
	tests := []struct {
		s, want string
	}{
		{`15:04:05.000`, `21:00:57.012`},
		{`15:04:05.999`, `21:00:57.012`},
		{`15:04:05,000`, `21:00:57,012`},
		{`15:04:05,999`, `21:00:57,012`},
	}

	// 数值时间表示的是2009年2月4日星期四，太平洋标准时间（PST）下的21:00:57.012345600。. md5:f8bd7a8cb5126dc5
	time := Unix(0, 1233810057012345600)
	for _, tt := range tests {
		if q := time.Format(tt.s); q != tt.want {
			t.Errorf("Format(%q) = got %q, want %q", tt.s, q, tt.want)
		}
	}
}

var longFractionalDigitsTests = []struct {
	value string
	want  int
}{
	// 9 digits
	{"2021-09-29T16:04:33.000000000Z", 0},
	{"2021-09-29T16:04:33.000000001Z", 1},
	{"2021-09-29T16:04:33.100000000Z", 100_000_000},
	{"2021-09-29T16:04:33.100000001Z", 100_000_001},
	{"2021-09-29T16:04:33.999999999Z", 999_999_999},
	{"2021-09-29T16:04:33.012345678Z", 12_345_678},
	// 10 digits, truncates
	{"2021-09-29T16:04:33.0000000000Z", 0},
	{"2021-09-29T16:04:33.0000000001Z", 0},
	{"2021-09-29T16:04:33.1000000000Z", 100_000_000},
	{"2021-09-29T16:04:33.1000000009Z", 100_000_000},
	{"2021-09-29T16:04:33.9999999999Z", 999_999_999},
	{"2021-09-29T16:04:33.0123456789Z", 12_345_678},
	// 11 digits, truncates
	{"2021-09-29T16:04:33.10000000000Z", 100_000_000},
	{"2021-09-29T16:04:33.00123456789Z", 1_234_567},
	// 12 digits, truncates
	{"2021-09-29T16:04:33.000123456789Z", 123_456},
	// 15 digits, truncates
	{"2021-09-29T16:04:33.9999999999999999Z", 999_999_999},
}

// Issue 48685 and 54567.
func TestParseFractionalSecondsLongerThanNineDigits(t *testing.T) {
	for _, tt := range longFractionalDigitsTests {
		for _, format := range []string{RFC3339, RFC3339Nano} {
			tm, err := Parse(format, tt.value)
			if err != nil {
				t.Errorf("Parse(%q, %q) error: %v", format, tt.value, err)
				continue
			}
			if got := tm.Nanosecond(); got != tt.want {
				t.Errorf("Parse(%q, %q) = got %d, want %d", format, tt.value, got, tt.want)
			}
		}
	}
}

//
//func FuzzFormatRFC3339(f *testing.F) {
//	for _, ts := range [][2]int64{
//		{math.MinInt64, math.MinInt64}, // 292277026304年08月26日 15时42分51秒 Z（UTC时间）. md5:5b0e4967ddc39a23
//		{-62167219200, 0},              // 0000-01-01T00:00:00Z
//		{1661201140, 676836973},        // 2022年8月22日20时45分40.676836973秒Z. md5:1fbc93ed7295c9d9
//		{253402300799, 999999999},      // 1999年12月31日23时59分59秒999999999Z（这是ISO 8601日期时间格式，表示的时间是世界协调时间）. md5:81581273a6df41a8
//		{math.MaxInt64, math.MaxInt64}, // -292277022365-05-08T08:17:07Z
//		// 这是ISO 8601日期时间格式的注释，表示的是一个UTC（协调世界时）的时间戳。具体翻译如下：
//		//
//		// - -292277022365：这是一个大数，代表自1970年1月1日（Unix纪元）以来的秒数，大约是公元前4713年。
//		// - -05：时区偏移，这里是格林威治标准时间（GMT）减去5小时，通常用于表示美国东部时间（EST）。
//		// - 05-08：月份和日期，即5月8日。
//		// - T08:17:07Z：时间，其中T是时间与日期的分隔符，08:17:07是小时、分钟和秒，Z表示UTC（Zulu时区，即零时区）。
//		//
//		// 所以，这个注释是在说一个大约在公元前4713年5月8日早上8点17分07秒（美国东部时间）的UTC时间。. md5:570b58db9b49ec1f
//	} {
//		f.Add(ts[0], ts[1], true, false, 0)
//		f.Add(ts[0], ts[1], false, true, 0)
//		for _, offset := range []int{0, 60, 60 * 60, 99*60*60 + 99*60, 123456789} {
//			f.Add(ts[0], ts[1], false, false, -offset)
//			f.Add(ts[0], ts[1], false, false, +offset)
//		}
//	}
//
//	f.Fuzz(func(t *testing.T, sec, nsec int64, useUTC, useLocal bool, tzOffset int) {
//		var loc *Location
//		switch {
//		case useUTC:
//			loc = UTC
//		case useLocal:
//			loc = Local
//		default:
//			loc = FixedZone("", tzOffset)
//		}
//		ts := Unix(sec, nsec).In(loc)
//
//		got := AppendFormatRFC3339(ts, nil, false)
//		want := AppendFormatAny(ts, nil, RFC3339)
//		if !bytes.Equal(got, want) {
//			t.Errorf("Format(%s, RFC3339) mismatch:\n\tgot:  %s\n\twant: %s", ts, got, want)
//		}
//
//		gotNanos := AppendFormatRFC3339(ts, nil, true)
//		wantNanos := AppendFormatAny(ts, nil, RFC3339Nano)
//		if !bytes.Equal(got, want) {
//			t.Errorf("Format(%s, RFC3339Nano) mismatch:\n\tgot:  %s\n\twant: %s", ts, gotNanos, wantNanos)
//		}
//	})
//}
//
//func FuzzParseRFC3339(f *testing.F) {
//	for _, tt := range formatTests {
//		f.Add(tt.result)
//	}
//	for _, tt := range parseTests {
//		f.Add(tt.value)
//	}
//	for _, tt := range parseErrorTests {
//		f.Add(tt.value)
//	}
//	for _, tt := range longFractionalDigitsTests {
//		f.Add(tt.value)
//	}
//
//	f.Fuzz(func(t *testing.T, s string) {
//		// equalTime 类似于 time.Time.Equal，但还会比较时区。. md5:75caad6c5a63b0ff
//		equalTime := func(t1, t2 Time) bool {
//			name1, offset1 := t1.Zone()
//			name2, offset2 := t2.Zone()
//			return t1.Equal(t2) && name1 == name2 && offset1 == offset2
//		}
//
//		for _, tz := range []*Location{UTC, Local} {
//			// 将时间解析为RFC3339或RFC3339Nano应该是相同的。. md5:337424a02fb1612e
//			t1, err1 := ParseAny(RFC3339, s, UTC, tz)
//			t2, err2 := ParseAny(RFC3339Nano, s, UTC, tz)
//			switch {
//			case (err1 == nil) != (err2 == nil):
//				t.Fatalf("ParseAny(%q) error mismatch:\n\tgot:  %v\n\twant: %v", s, err1, err2)
//			case !equalTime(t1, t2):
//				t.Fatalf("ParseAny(%q) value mismatch:\n\tgot:  %v\n\twant: %v", s, t1, t2)
//			}
//
//			// TODO(https://go.dev/issue/54580): // 在ParseAny拒绝所有无效的RFC 3339之后，删除这些检查。
//			// md5:d137d8744b3d3843
//			if err1 == nil {
//				num2 := func(s string) byte { return 10*(s[0]-'0') + (s[1] - '0') }
//				switch {
//				case len(s) > 12 && s[12] == ':':
//					t.Skipf("ParseAny(%q) incorrectly allows single-digit hour fields", s)
//				case len(s) > 19 && s[19] == ',':
//					t.Skipf("ParseAny(%q) incorrectly allows comma as sub-second separator", s)
//				case !strings.HasSuffix(s, "Z") && len(s) > 4 && (num2(s[len(s)-5:]) >= 24 || num2(s[len(s)-2:]) >= 60):
//					t.Skipf("ParseAny(%q) incorrectly allows out-of-range zone offset", s)
//				}
//			}
//
//			// 定制解析器应与通用解析器完全相同。. md5:c589baa9a07544ae
//			switch got, ok := ParseRFC3339(s, tz); {
//			case ok != (err1 == nil):
//				t.Fatalf("ParseRFC3339(%q) error mismatch:\n\tgot:  %v\n\twant: %v", s, ok, err1 == nil)
//			case !equalTime(got, t1):
//				t.Fatalf("ParseRFC3339(%q) value mismatch:\n\tgot:  %v\n\twant: %v", s, got, t2)
//			}
//		}
//	})
//}
