//版权所有2009年Go作者。所有权利保留。
//使用此源代码受BSD风格
//可以在LICENSE文件中找到的许可证。
// md5:2e9dc81828a3be8a

package time_test

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	. "github.com/888go/gosdk/time"
	"math"
	"math/big"
	"math/rand"
	"runtime"
	"strings"
	"sync"
	"testing"
	"testing/quick"
)

// 我们应该在PST/PDT时区，但如果时区文件缺失，我们不会在。这个测试的目的是至少解释为什么后续的一些测试会失败。
// md5:a877153cf135d94d
func TestZoneData(t *testing.T) {
	lt := Now()
	// PST位于西边8小时，PDT位于西边7小时。我们可以使用这个名字，但它并不唯一。. md5:cbe7a8e2be638df6
	if name, off := lt.Zone(); off != -8*60*60 && off != -7*60*60 {
		t.Errorf("Unable to find US Pacific time zone data for testing; time zone is %q offset %d", name, off)
		t.Error("Likely problem: the time zone files have not been installed.")
	}
}

// parsedTime 是表示解析后时间值的结构体。. md5:a90be61d11060ff5
type parsedTime struct {
	Year                 int
	Month                Month
	Day                  int
	Hour, Minute, Second int // 15:04:05 is 15, 4, 5.
	Nanosecond           int // Fractional second.
	Weekday              Weekday
	ZoneOffset           int    // UTC以东的秒数，例如-7*60*60表示-0700. md5:5eace65a71c2a861
	Zone                 string // e.g., "MST"
}

type TimeTest struct {
	seconds int64
	golden  parsedTime
}

var utctests = []TimeTest{
	{0, parsedTime{1970, January, 1, 0, 0, 0, 0, Thursday, 0, "UTC"}},
	{1221681866, parsedTime{2008, September, 17, 20, 4, 26, 0, Wednesday, 0, "UTC"}},
	{-1221681866, parsedTime{1931, April, 16, 3, 55, 34, 0, Thursday, 0, "UTC"}},
	{-11644473600, parsedTime{1601, January, 1, 0, 0, 0, 0, Monday, 0, "UTC"}},
	{599529660, parsedTime{1988, December, 31, 0, 1, 0, 0, Saturday, 0, "UTC"}},
	{978220860, parsedTime{2000, December, 31, 0, 1, 0, 0, Sunday, 0, "UTC"}},
}

var nanoutctests = []TimeTest{
	{0, parsedTime{1970, January, 1, 0, 0, 0, 1e8, Thursday, 0, "UTC"}},
	{1221681866, parsedTime{2008, September, 17, 20, 4, 26, 2e8, Wednesday, 0, "UTC"}},
}

var localtests = []TimeTest{
	{0, parsedTime{1969, December, 31, 16, 0, 0, 0, Wednesday, -8 * 60 * 60, "PST"}},
	{1221681866, parsedTime{2008, September, 17, 13, 4, 26, 0, Wednesday, -7 * 60 * 60, "PDT"}},
	{2159200800, parsedTime{2038, June, 3, 11, 0, 0, 0, Thursday, -7 * 60 * 60, "PDT"}},
	{2152173599, parsedTime{2038, March, 14, 1, 59, 59, 0, Sunday, -8 * 60 * 60, "PST"}},
	{2152173600, parsedTime{2038, March, 14, 3, 0, 0, 0, Sunday, -7 * 60 * 60, "PDT"}},
	{2152173601, parsedTime{2038, March, 14, 3, 0, 1, 0, Sunday, -7 * 60 * 60, "PDT"}},
	{2172733199, parsedTime{2038, November, 7, 1, 59, 59, 0, Sunday, -7 * 60 * 60, "PDT"}},
	{2172733200, parsedTime{2038, November, 7, 1, 0, 0, 0, Sunday, -8 * 60 * 60, "PST"}},
	{2172733201, parsedTime{2038, November, 7, 1, 0, 1, 0, Sunday, -8 * 60 * 60, "PST"}},
}

var nanolocaltests = []TimeTest{
	{0, parsedTime{1969, December, 31, 16, 0, 0, 1e8, Wednesday, -8 * 60 * 60, "PST"}},
	{1221681866, parsedTime{2008, September, 17, 13, 4, 26, 3e8, Wednesday, -7 * 60 * 60, "PDT"}},
}

func same(t Time, u *parsedTime) bool {
	// Check aggregates.
	year, month, day := t.Date()
	hour, min, sec := t.Clock()
	name, offset := t.Zone()
	if year != u.Year || month != u.Month || day != u.Day ||
		hour != u.Hour || min != u.Minute || sec != u.Second ||
		name != u.Zone || offset != u.ZoneOffset {
		return false
	}
	// 检查单个条目。. md5:c7203452782bccc9
	return t.Year() == u.Year &&
		t.Month() == u.Month &&
		t.Day() == u.Day &&
		t.Hour() == u.Hour &&
		t.Minute() == u.Minute &&
		t.Second() == u.Second &&
		t.Nanosecond() == u.Nanosecond &&
		t.Weekday() == u.Weekday
}

func TestSecondsToUTC(t *testing.T) {
	for _, test := range utctests {
		sec := test.seconds
		golden := &test.golden
		tm := Unix(sec, 0).UTC()
		newsec := tm.Unix()
		if newsec != sec {
			t.Errorf("SecondsToUTC(%d).Seconds() = %d", sec, newsec)
		}
		if !same(tm, golden) {
			t.Errorf("SecondsToUTC(%d):  // %#v", sec, tm)
			t.Errorf("  want=%+v", *golden)
			t.Errorf("  have=%v", tm.Format(RFC3339+" MST"))
		}
	}
}

func TestNanosecondsToUTC(t *testing.T) {
	for _, test := range nanoutctests {
		golden := &test.golden
		nsec := test.seconds*1e9 + int64(golden.Nanosecond)
		tm := Unix(0, nsec).UTC()
		newnsec := tm.Unix()*1e9 + int64(tm.Nanosecond())
		if newnsec != nsec {
			t.Errorf("NanosecondsToUTC(%d).Nanoseconds() = %d", nsec, newnsec)
		}
		if !same(tm, golden) {
			t.Errorf("NanosecondsToUTC(%d):", nsec)
			t.Errorf("  want=%+v", *golden)
			t.Errorf("  have=%+v", tm.Format(RFC3339+" MST"))
		}
	}
}

func TestSecondsToLocalTime(t *testing.T) {
	for _, test := range localtests {
		sec := test.seconds
		golden := &test.golden
		tm := Unix(sec, 0)
		newsec := tm.Unix()
		if newsec != sec {
			t.Errorf("SecondsToLocalTime(%d).Seconds() = %d", sec, newsec)
		}
		if !same(tm, golden) {
			t.Errorf("SecondsToLocalTime(%d):", sec)
			t.Errorf("  want=%+v", *golden)
			t.Errorf("  have=%+v", tm.Format(RFC3339+" MST"))
		}
	}
}

func TestNanosecondsToLocalTime(t *testing.T) {
	for _, test := range nanolocaltests {
		golden := &test.golden
		nsec := test.seconds*1e9 + int64(golden.Nanosecond)
		tm := Unix(0, nsec)
		newnsec := tm.Unix()*1e9 + int64(tm.Nanosecond())
		if newnsec != nsec {
			t.Errorf("NanosecondsToLocalTime(%d).Seconds() = %d", nsec, newnsec)
		}
		if !same(tm, golden) {
			t.Errorf("NanosecondsToLocalTime(%d):", nsec)
			t.Errorf("  want=%+v", *golden)
			t.Errorf("  have=%+v", tm.Format(RFC3339+" MST"))
		}
	}
}

func TestSecondsToUTCAndBack(t *testing.T) {
	f := func(sec int64) bool { return Unix(sec, 0).UTC().Unix() == sec }
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

func TestNanosecondsToUTCAndBack(t *testing.T) {
	f := func(nsec int64) bool {
		t := Unix(0, nsec).UTC()
		ns := t.Unix()*1e9 + int64(t.Nanosecond())
		return ns == nsec
	}
	f32 := func(nsec int32) bool { return f(int64(nsec)) }
	cfg := &quick.Config{MaxCount: 10000}

	// 首先尝试一个小的日期，然后是大的日期。（对于int64中的纳秒，跨度只有几百年）
	// md5:3360e6892f5fb9e4
	if err := quick.Check(f32, cfg); err != nil {
		t.Fatal(err)
	}
	if err := quick.Check(f, cfg); err != nil {
		t.Fatal(err)
	}
}

func TestUnixMilli(t *testing.T) {
	f := func(msec int64) bool {
		t := UnixMilli(msec)
		return t.UnixMilli() == msec
	}
	cfg := &quick.Config{MaxCount: 10000}
	if err := quick.Check(f, cfg); err != nil {
		t.Fatal(err)
	}
}

func TestUnixMicro(t *testing.T) {
	f := func(usec int64) bool {
		t := UnixMicro(usec)
		return t.UnixMicro() == usec
	}
	cfg := &quick.Config{MaxCount: 10000}
	if err := quick.Check(f, cfg); err != nil {
		t.Fatal(err)
	}
}

// 时间 routines（调度）无法获取绝对时间（自零点的秒数），但我们需要它来计算出像"四舍五入到最近的3 ns"这样的奇数结果。计算方法为：t - year1 = (t - 1970) + (1970 - 2001) + (2001 - 1)。
// t - 1970 由 Unix 和 Nanosecond 返回。
// 1970 - 2001 是 -(31 * 365 + 8) * 86400 = -978307200 秒。
// 2001 - 1 是 2000 * 365.2425 * 86400 = 63113904000 秒。
// md5:a5c6dba33aaa4bcd
const unixToZero = -978307200 + 63113904000

// abs 返回存储在 t 中的绝对时间，以秒和纳秒为单位。. md5:9b3a61e1f35f15df
func abs(t Time) (sec, nsec int64) {
	unix := t.Unix()
	nano := t.Nanosecond()
	return unix + unixToZero, int64(nano)
}

// absString 将 abs 转换为十进制字符串。. md5:f4ecea5809009a31
func absString(t Time) string {
	sec, nsec := abs(t)
	if sec < 0 {
		sec = -sec
		nsec = -nsec
		if nsec < 0 {
			nsec += 1e9
			sec--
		}
		return fmt.Sprintf("-%d%09d", sec, nsec)
	}
	return fmt.Sprintf("%d%09d", sec, nsec)
}

var truncateRoundTests = []struct {
	t Time
	d Duration
}{
	{Date(-1, January, 1, 12, 15, 30, 5e8, UTC), 3},
	{Date(-1, January, 1, 12, 15, 31, 5e8, UTC), 3},
	{Date(2012, January, 1, 12, 15, 30, 5e8, UTC), Second},
	{Date(2012, January, 1, 12, 15, 31, 5e8, UTC), Second},
	{Unix(-19012425939, 649146258), 7435029458905025217}, // 5.8 * d 向上舍入到 6 * d，但是 .8 * d + .8 * d < 0 < d. md5:17cdc3615444f015
}

func TestTruncateRound(t *testing.T) {
	var (
		bsec  = new(big.Int)
		bnsec = new(big.Int)
		bd    = new(big.Int)
		bt    = new(big.Int)
		br    = new(big.Int)
		bq    = new(big.Int)
		b1e9  = new(big.Int)
	)

	b1e9.SetInt64(1e9)

	testOne := func(ti, tns, di int64) bool {
		t.Helper()

		t0 := Unix(ti, tns).UTC()
		d := Duration(di)
		if d < 0 {
			d = -d
		}
		if d <= 0 {
			d = 1
		}

		// 计算bt为绝对纳秒。. md5:bb29a5c2c2e5fcb0
		sec, nsec := abs(t0)
		bsec.SetInt64(sec)
		bnsec.SetInt64(nsec)
		bt.Mul(bsec, b1e9)
		bt.Add(bt, bnsec)

		// 计算商和模d的余数。. md5:1ad7456726aa86c2
		bd.SetInt64(int64(d))
		bq.DivMod(bt, bd, br)

		// 通过取余来截断。
		// br 小于 d，所以它可以容纳在 int64 中。
		// md5:44c8c5d861d5fb5b
		r := br.Int64()
		t1 := t0.Add(-Duration(r))

		// 检查time.Truncate是否正常工作。. md5:51e423e287005315
		if trunc := t0.Truncate(d); trunc != t1 {
			t.Errorf("Time.Truncate(%s, %s) = %s, want %s\n"+
				"%v trunc %v =\n%v want\n%v",
				t0.Format(RFC3339Nano), d, trunc, t1.Format(RFC3339Nano),
				absString(t0), int64(d), absString(trunc), absString(t1))
			return false
		}

		// 如果余数 r 大于 d/2 或者等于 d/2，为了四舍五入，我们将 d 添加回去。
		// 注释掉的代码会将半数四舍五入到偶数，而不是向上取整。但这会使它依赖时区，这有点奇怪。
		// md5:5621afa60a7d98b6
		if r > int64(d)/2 || r+r == int64(d) /*&& bq.Bit(0) == 1*/ {
			t1 = t1.Add(d)
		}

		// 检查time.Round函数是否正常工作。. md5:e5b2f97fd408128f
		if rnd := t0.Round(d); rnd != t1 {
			t.Errorf("Time.Round(%s, %s) = %s, want %s\n"+
				"%v round %v =\n%v want\n%v",
				t0.Format(RFC3339Nano), d, rnd, t1.Format(RFC3339Nano),
				absString(t0), int64(d), absString(rnd), absString(t1))
			return false
		}
		return true
	}

	// manual test cases
	for _, tt := range truncateRoundTests {
		testOne(tt.t.Unix(), int64(tt.t.Nanosecond()), int64(tt.d))
	}

	// exhaustive near 0
	for i := 0; i < 100; i++ {
		for j := 1; j < 100; j++ {
			testOne(unixToZero, int64(i), int64(j))
			testOne(unixToZero, -int64(i), int64(j))
			if t.Failed() {
				return
			}
		}
	}

	if t.Failed() {
		return
	}

	// 随机生成的测试用例. md5:e6955df765195da3
	cfg := &quick.Config{MaxCount: 100000}
	if testing.Short() {
		cfg.MaxCount = 1000
	}

	// divisors of Second
	f1 := func(ti int64, tns int32, logdi int32) bool {
		d := Duration(1)
		a, b := uint(logdi%9), (logdi>>16)%9
		d <<= a
		for i := 0; i < int(b); i++ {
			d *= 5
		}

		// Make room for unix ↔ internal conversion.
		// We don't care about behavior too close to ± 2^63 Unix seconds.
		// It is full of wraparounds but will never happen in a reasonable program.
		// (Or maybe not? See go.dev/issue/20678. In any event, they're not handled today.)
		ti >>= 1

		return testOne(ti, int64(tns), int64(d))
	}
	quick.Check(f1, cfg)

	// multiples of Second
	f2 := func(ti int64, tns int32, di int32) bool {
		d := Duration(di) * Second
		if d < 0 {
			d = -d
		}
		ti >>= 1 // see comment in f1
		return testOne(ti, int64(tns), int64(d))
	}
	quick.Check(f2, cfg)

	// halfway cases
	f3 := func(tns, di int64) bool {
		di &= 0xfffffffe
		if di == 0 {
			di = 2
		}
		tns -= tns % di
		if tns < 0 {
			tns += di / 2
		} else {
			tns -= di / 2
		}
		return testOne(0, tns, di)
	}
	quick.Check(f3, cfg)

	// full generality
	f4 := func(ti int64, tns int32, di int64) bool {
		ti >>= 1 // see comment in f1
		return testOne(ti, int64(tns), di)
	}
	quick.Check(f4, cfg)
}

type ISOWeekTest struct {
	year       int // year
	month, day int // month and day
	yex        int // expected year
	wex        int // expected week
}

var isoWeekTests = []ISOWeekTest{
	{1981, 1, 1, 1981, 1}, {1982, 1, 1, 1981, 53}, {1983, 1, 1, 1982, 52},
	{1984, 1, 1, 1983, 52}, {1985, 1, 1, 1985, 1}, {1986, 1, 1, 1986, 1},
	{1987, 1, 1, 1987, 1}, {1988, 1, 1, 1987, 53}, {1989, 1, 1, 1988, 52},
	{1990, 1, 1, 1990, 1}, {1991, 1, 1, 1991, 1}, {1992, 1, 1, 1992, 1},
	{1993, 1, 1, 1992, 53}, {1994, 1, 1, 1993, 52}, {1995, 1, 2, 1995, 1},
	{1996, 1, 1, 1996, 1}, {1996, 1, 7, 1996, 1}, {1996, 1, 8, 1996, 2},
	{1997, 1, 1, 1997, 1}, {1998, 1, 1, 1998, 1}, {1999, 1, 1, 1998, 53},
	{2000, 1, 1, 1999, 52}, {2001, 1, 1, 2001, 1}, {2002, 1, 1, 2002, 1},
	{2003, 1, 1, 2003, 1}, {2004, 1, 1, 2004, 1}, {2005, 1, 1, 2004, 53},
	{2006, 1, 1, 2005, 52}, {2007, 1, 1, 2007, 1}, {2008, 1, 1, 2008, 1},
	{2009, 1, 1, 2009, 1}, {2010, 1, 1, 2009, 53}, {2010, 1, 1, 2009, 53},
	{2011, 1, 1, 2010, 52}, {2011, 1, 2, 2010, 52}, {2011, 1, 3, 2011, 1},
	{2011, 1, 4, 2011, 1}, {2011, 1, 5, 2011, 1}, {2011, 1, 6, 2011, 1},
	{2011, 1, 7, 2011, 1}, {2011, 1, 8, 2011, 1}, {2011, 1, 9, 2011, 1},
	{2011, 1, 10, 2011, 2}, {2011, 1, 11, 2011, 2}, {2011, 6, 12, 2011, 23},
	{2011, 6, 13, 2011, 24}, {2011, 12, 25, 2011, 51}, {2011, 12, 26, 2011, 52},
	{2011, 12, 27, 2011, 52}, {2011, 12, 28, 2011, 52}, {2011, 12, 29, 2011, 52},
	{2011, 12, 30, 2011, 52}, {2011, 12, 31, 2011, 52}, {1995, 1, 1, 1994, 52},
	{2012, 1, 1, 2011, 52}, {2012, 1, 2, 2012, 1}, {2012, 1, 8, 2012, 1},
	{2012, 1, 9, 2012, 2}, {2012, 12, 23, 2012, 51}, {2012, 12, 24, 2012, 52},
	{2012, 12, 30, 2012, 52}, {2012, 12, 31, 2013, 1}, {2013, 1, 1, 2013, 1},
	{2013, 1, 6, 2013, 1}, {2013, 1, 7, 2013, 2}, {2013, 12, 22, 2013, 51},
	{2013, 12, 23, 2013, 52}, {2013, 12, 29, 2013, 52}, {2013, 12, 30, 2014, 1},
	{2014, 1, 1, 2014, 1}, {2014, 1, 5, 2014, 1}, {2014, 1, 6, 2014, 2},
	{2015, 1, 1, 2015, 1}, {2016, 1, 1, 2015, 53}, {2017, 1, 1, 2016, 52},
	{2018, 1, 1, 2018, 1}, {2019, 1, 1, 2019, 1}, {2020, 1, 1, 2020, 1},
	{2021, 1, 1, 2020, 53}, {2022, 1, 1, 2021, 52}, {2023, 1, 1, 2022, 52},
	{2024, 1, 1, 2024, 1}, {2025, 1, 1, 2025, 1}, {2026, 1, 1, 2026, 1},
	{2027, 1, 1, 2026, 53}, {2028, 1, 1, 2027, 52}, {2029, 1, 1, 2029, 1},
	{2030, 1, 1, 2030, 1}, {2031, 1, 1, 2031, 1}, {2032, 1, 1, 2032, 1},
	{2033, 1, 1, 2032, 53}, {2034, 1, 1, 2033, 52}, {2035, 1, 1, 2035, 1},
	{2036, 1, 1, 2036, 1}, {2037, 1, 1, 2037, 1}, {2038, 1, 1, 2037, 53},
	{2039, 1, 1, 2038, 52}, {2040, 1, 1, 2039, 52},
}

func TestISOWeek(t *testing.T) {
	// 选择的日期和边界情况. md5:4202016edce8d9ff
	for _, wt := range isoWeekTests {
		dt := Date(wt.year, Month(wt.month), wt.day, 0, 0, 0, 0, UTC)
		y, w := dt.ISOWeek()
		if w != wt.wex || y != wt.yex {
			t.Errorf("got %d/%d; expected %d/%d for %d-%02d-%02d",
				y, w, wt.yex, wt.wex, wt.year, wt.month, wt.day)
		}
	}

	// 唯一真正的不变性：1月4日是在第1周内。. md5:278b4045520e4ef0
	for year := 1950; year < 2100; year++ {
		if y, w := Date(year, January, 4, 0, 0, 0, 0, UTC).ISOWeek(); y != year || w != 1 {
			t.Errorf("got %d/%d; expected %d/1 for Jan 04", y, w, year)
		}
	}
}

type YearDayTest struct {
	year, month, day int
	yday             int
}

// 在多种不同场景和边界情况下测试YearDay
// md5:467db368371fe934
var yearDayTests = []YearDayTest{
	// Non-leap-year tests
	{2007, 1, 1, 1},
	{2007, 1, 15, 15},
	{2007, 2, 1, 32},
	{2007, 2, 15, 46},
	{2007, 3, 1, 60},
	{2007, 3, 15, 74},
	{2007, 4, 1, 91},
	{2007, 12, 31, 365},

	// Leap-year tests
	{2008, 1, 1, 1},
	{2008, 1, 15, 15},
	{2008, 2, 1, 32},
	{2008, 2, 15, 46},
	{2008, 3, 1, 61},
	{2008, 3, 15, 75},
	{2008, 4, 1, 92},
	{2008, 12, 31, 366},

	// 看起来像是闰年（但实际上不是）的测试. md5:ce6d977814d12f28
	{1900, 1, 1, 1},
	{1900, 1, 15, 15},
	{1900, 2, 1, 32},
	{1900, 2, 15, 46},
	{1900, 3, 1, 60},
	{1900, 3, 15, 74},
	{1900, 4, 1, 91},
	{1900, 12, 31, 365},

	// 第一年的测试（非闰年）. md5:4d82de1d5378f886
	{1, 1, 1, 1},
	{1, 1, 15, 15},
	{1, 2, 1, 32},
	{1, 2, 15, 46},
	{1, 3, 1, 60},
	{1, 3, 15, 74},
	{1, 4, 1, 91},
	{1, 12, 31, 365},

	// 减去一年的测试（非闰年）. md5:9b77b30db967b744
	{-1, 1, 1, 1},
	{-1, 1, 15, 15},
	{-1, 2, 1, 32},
	{-1, 2, 15, 46},
	{-1, 3, 1, 60},
	{-1, 3, 15, 74},
	{-1, 4, 1, 91},
	{-1, 12, 31, 365},

	// 400年前的BC测试（闰年）. md5:e252a45b423cb8ca
	{-400, 1, 1, 1},
	{-400, 1, 15, 15},
	{-400, 2, 1, 32},
	{-400, 2, 15, 46},
	{-400, 3, 1, 61},
	{-400, 3, 15, 75},
	{-400, 4, 1, 92},
	{-400, 12, 31, 366},

	// Special Cases

	// 公历变更（无影响）. md5:93bdb417c060cda9
	{1582, 10, 4, 277},
	{1582, 10, 15, 288},
}

// 检查YearDay是否为位置敏感的. md5:5b8dab4efd11f82d
var yearDayLocations = []*Location{
	FixedZone("UTC-8", -8*60*60),
	FixedZone("UTC-4", -4*60*60),
	UTC,
	FixedZone("UTC+4", 4*60*60),
	FixedZone("UTC+8", 8*60*60),
}

func TestYearDay(t *testing.T) {
	for i, loc := range yearDayLocations {
		for _, ydt := range yearDayTests {
			dt := Date(ydt.year, Month(ydt.month), ydt.day, 0, 0, 0, 0, loc)
			yday := dt.YearDay()
			if yday != ydt.yday {
				t.Errorf("Date(%d-%02d-%02d in %v).YearDay() = %d, want %d",
					ydt.year, ydt.month, ydt.day, loc, yday, ydt.yday)
				continue
			}

			if ydt.year < 0 || ydt.year > 9999 {
				continue
			}
			f := fmt.Sprintf("%04d-%02d-%02d %03d %+.2d00",
				ydt.year, ydt.month, ydt.day, ydt.yday, (i-2)*4)
			dt1, err := Parse("2006-01-02 002 -0700", f)
			if err != nil {
				t.Errorf(`Parse("2006-01-02 002 -0700", %q): %v`, f, err)
				continue
			}
			if !dt1.Equal(dt) {
				t.Errorf(`Parse("2006-01-02 002 -0700", %q) = %v, want %v`, f, dt1, dt)
			}
		}
	}
}

var durationTests = []struct {
	str string
	d   Duration
}{
	{"0s", 0},
	{"1ns", 1 * Nanosecond},
	{"1.1µs", 1100 * Nanosecond},
	{"2.2ms", 2200 * Microsecond},
	{"3.3s", 3300 * Millisecond},
	{"4m5s", 4*Minute + 5*Second},
	{"4m5.001s", 4*Minute + 5001*Millisecond},
	{"5h6m7.001s", 5*Hour + 6*Minute + 7001*Millisecond},
	{"8m0.000000001s", 8*Minute + 1*Nanosecond},
	{"2562047h47m16.854775807s", 1<<63 - 1},
	{"-2562047h47m16.854775808s", -1 << 63},
}

func TestDurationString(t *testing.T) {
	for _, tt := range durationTests {
		if str := tt.d.String(); str != tt.str {
			t.Errorf("Duration(%d).String() = %s, want %s", int64(tt.d), str, tt.str)
		}
		if tt.d > 0 {
			if str := (-tt.d).String(); str != "-"+tt.str {
				t.Errorf("Duration(%d).String() = %s, want %s", int64(-tt.d), str, "-"+tt.str)
			}
		}
	}
}

var dateTests = []struct {
	year, month, day, hour, min, sec, nsec int
	z                                      *Location
	unix                                   int64
}{
	{2011, 11, 6, 1, 0, 0, 0, Local, 1320566400},   // 1:00:00 PDT
	{2011, 11, 6, 1, 59, 59, 0, Local, 1320569999}, // 1:59:59 PDT
	{2011, 11, 6, 2, 0, 0, 0, Local, 1320573600},   // 2:00:00 PST

	{2011, 3, 13, 1, 0, 0, 0, Local, 1300006800},   // 1:00:00 PST
	{2011, 3, 13, 1, 59, 59, 0, Local, 1300010399}, // 1:59:59 PST
	{2011, 3, 13, 3, 0, 0, 0, Local, 1300010400},   // 3:00:00 PDT
	{2011, 3, 13, 2, 30, 0, 0, Local, 1300008600},  // 2:30:00 PDT ≡ 1:30 PST
	{2012, 12, 24, 0, 0, 0, 0, Local, 1356336000},  // Leap year

	// 2011年11月18日星期五7:56:35 PST的多种名称. md5:8dd2ad304f033b90
	{2011, 11, 18, 7, 56, 35, 0, Local, 1321631795},                 // Nov 18 7:56:35
	{2011, 11, 19, -17, 56, 35, 0, Local, 1321631795},               // Nov 19 -17:56:35
	{2011, 11, 17, 31, 56, 35, 0, Local, 1321631795},                // Nov 17 31:56:35
	{2011, 11, 18, 6, 116, 35, 0, Local, 1321631795},                // Nov 18 6:116:35
	{2011, 10, 49, 7, 56, 35, 0, Local, 1321631795},                 // Oct 49 7:56:35
	{2011, 11, 18, 7, 55, 95, 0, Local, 1321631795},                 // Nov 18 7:55:95
	{2011, 11, 18, 7, 56, 34, 1e9, Local, 1321631795},               // Nov 18 7:56:34 + 10⁹ns
	{2011, 12, -12, 7, 56, 35, 0, Local, 1321631795},                // Dec -21 7:56:35
	{2012, 1, -43, 7, 56, 35, 0, Local, 1321631795},                 // Jan -52 7:56:35 2012
	{2012, int(January - 2), 18, 7, 56, 35, 0, Local, 1321631795},   // (一月-2) 18 7:56:35 2012. md5:6af96243b33568b3
	{2010, int(December + 11), 18, 7, 56, 35, 0, Local, 1321631795}, //(Dec+11) 18 07:56:35 2010
	//
	// 这段Go语言注释的中文翻译是："（Dec+11）2010年12月11日 18:07:56"。这是时间戳的格式，Dec+11表示月份（12月后加11，即1月），18:07:56是24小时制的时间，2010则是年份。. md5:edaae9041aa5b491
}

func TestDate(t *testing.T) {
	for _, tt := range dateTests {
		time := Date(tt.year, Month(tt.month), tt.day, tt.hour, tt.min, tt.sec, tt.nsec, tt.z)
		want := Unix(tt.unix, 0)
		if !time.Equal(want) {
			t.Errorf("Date(%d, %d, %d, %d, %d, %d, %d, %s) = %v, want %v",
				tt.year, tt.month, tt.day, tt.hour, tt.min, tt.sec, tt.nsec, tt.z,
				time, want)
		}
	}
}

// 从
// 2011年11月18日星期五，太平洋标准时间7:56:35
// 到
// 2016年3月19日星期四，太平洋标准时间7:56:35
// 的多种方式
// md5:85909e1e910bda87
var addDateTests = []struct {
	years, months, days int
}{
	{4, 4, 1},
	{3, 16, 1},
	{3, 15, 30},
	{5, -6, -18 - 30 - 12},
}

func TestAddDate(t *testing.T) {
	t0 := Date(2011, 11, 18, 7, 56, 35, 0, UTC)
	t1 := Date(2016, 3, 19, 7, 56, 35, 0, UTC)
	for _, at := range addDateTests {
		time := t0.AddDate(at.years, at.months, at.days)
		if !time.Equal(t1) {
			t.Errorf("AddDate(%d, %d, %d) = %v, want %v",
				at.years, at.months, at.days,
				time, t1)
		}
	}
}

var daysInTests = []struct {
	year, month, di int
}{
	{2011, 1, 31},  // 一月，第一个月，31天. md5:3158b6d74f31a42d
	{2011, 2, 28},  // 二月，非闰年，28天. md5:e49798a45ef3f35c
	{2012, 2, 29},  // 二月，闰年，29天. md5:5026573ed7f29c97
	{2011, 6, 30},  // June, 30 days
	{2011, 12, 31}, // 上个月，12月，有31天. md5:f795504febd08823
}

//func TestDaysIn(t *testing.T) {
//// daysIn 函数没有被导出。
//// 通过在 internal_test.go 文件中的 `var DaysIn = daysIn` 语句来测试 daysIn 函数。
//// md5:a45861afd1e9e8fa
//	for _, tt := range daysInTests {
//		di := DaysIn(Month(tt.month), tt.year)
//		if di != tt.di {
//			t.Errorf("got %d; expected %d for %d-%02d",
//				di, tt.di, tt.year, tt.month)
//		}
//	}
//}

func TestAddToExactSecond(t *testing.T) {
	// 向当前时间添加一个量，使其向上舍入到下一个精确的秒。
	// 此测试检查nsec字段是否仍然在范围[0, 999999999]内。
	// md5:c18836b1f3d78e03
	t1 := Now()
	t2 := t1.Add(Second - Duration(t1.Nanosecond()))
	sec := (t1.Second() + 1) % 60
	if t2.Second() != sec || t2.Nanosecond() != 0 {
		t.Errorf("sec = %d, nsec = %d, want sec = %d, nsec = 0", t2.Second(), t2.Nanosecond(), sec)
	}
}

func equalTimeAndZone(a, b Time) bool {
	aname, aoffset := a.Zone()
	bname, boffset := b.Zone()
	return a.Equal(b) && aoffset == boffset && aname == bname
}

var gobTests = []Time{
	Date(0, 1, 2, 3, 4, 5, 6, UTC),
	Date(7, 8, 9, 10, 11, 12, 13, FixedZone("", 0)),
	Unix(81985467080890095, 0x76543210), // Time.sec: 0x0123456789ABCDEF
	// Time.sec: 0x0123456789ABCDEF
	// Time.sec: 0x0123456789ABCDEF
	// Time.sec: 0x0123456789ABCDEF
	// Time.sec: 0x0123456789ABCDEF
	// Time.sec: 0x0123456789ABCDEF
	// 时间.sec: 0x0123456789ABCDEF（十六进制表示的秒值）. md5:9b45ef185347f6e0
	{}, // nil location
	Date(1, 2, 3, 4, 5, 6, 7, FixedZone("", 32767*60)),
	Date(1, 2, 3, 4, 5, 6, 7, FixedZone("", -32768*60)),
}

func TestTimeGob(t *testing.T) {
	var b bytes.Buffer
	enc := gob.NewEncoder(&b)
	dec := gob.NewDecoder(&b)
	for _, tt := range gobTests {
		var gobtt Time
		if err := enc.Encode(&tt); err != nil {
			t.Errorf("%v gob Encode error = %q, want nil", tt, err)
		} else if err := dec.Decode(&gobtt); err != nil {
			t.Errorf("%v gob Decode error = %q, want nil", tt, err)
		} else if !equalTimeAndZone(gobtt, tt) {
			t.Errorf("Decoded time = %v, want %v", gobtt, tt)
		}
		b.Reset()
	}
}

var invalidEncodingTests = []struct {
	bytes []byte
	want  string
}{
	{[]byte{}, "Time.UnmarshalBinary: no data"},
	{[]byte{0, 2, 3}, "Time.UnmarshalBinary: unsupported version"},
	{[]byte{1, 2, 3}, "Time.UnmarshalBinary: invalid length"},
}

func TestInvalidTimeGob(t *testing.T) {
	for _, tt := range invalidEncodingTests {
		var ignored Time
		err := ignored.GobDecode(tt.bytes)
		if err == nil || err.Error() != tt.want {
			t.Errorf("time.GobDecode(%#v) error = %v, want %v", tt.bytes, err, tt.want)
		}
		err = ignored.UnmarshalBinary(tt.bytes)
		if err == nil || err.Error() != tt.want {
			t.Errorf("time.UnmarshalBinary(%#v) error = %v, want %v", tt.bytes, err, tt.want)
		}
	}
}

var notEncodableTimes = []struct {
	time Time
	want string
}{
	{Date(0, 1, 2, 3, 4, 5, 6, FixedZone("", -1*60)), "Time.MarshalBinary: unexpected zone offset"},
	{Date(0, 1, 2, 3, 4, 5, 6, FixedZone("", -32769*60)), "Time.MarshalBinary: unexpected zone offset"},
	{Date(0, 1, 2, 3, 4, 5, 6, FixedZone("", 32768*60)), "Time.MarshalBinary: unexpected zone offset"},
}

func TestNotGobEncodableTime(t *testing.T) {
	for _, tt := range notEncodableTimes {
		_, err := tt.time.GobEncode()
		if err == nil || err.Error() != tt.want {
			t.Errorf("%v GobEncode error = %v, want %v", tt.time, err, tt.want)
		}
		_, err = tt.time.MarshalBinary()
		if err == nil || err.Error() != tt.want {
			t.Errorf("%v MarshalBinary error = %v, want %v", tt.time, err, tt.want)
		}
	}
}

var jsonTests = []struct {
	time Time
	json string
}{
	{Date(9999, 4, 12, 23, 20, 50, 520*1e6, UTC), `"9999-04-12T23:20:50.52Z"`},
	{Date(1996, 12, 19, 16, 39, 57, 0, Local), `"1996-12-19T16:39:57-08:00"`},
	{Date(0, 1, 1, 0, 0, 0, 1, FixedZone("", 1*60)), `"0000-01-01T00:00:00.000000001+00:01"`},
	{Date(2020, 1, 1, 0, 0, 0, 0, FixedZone("", 23*60*60+59*60)), `"2020-01-01T00:00:00+23:59"`},
}

func TestTimeJSON(t *testing.T) {
	for _, tt := range jsonTests {
		var jsonTime Time

		if jsonBytes, err := json.Marshal(tt.time); err != nil {
			t.Errorf("%v json.Marshal error = %v, want nil", tt.time, err)
		} else if string(jsonBytes) != tt.json {
			t.Errorf("%v JSON = %#q, want %#q", tt.time, string(jsonBytes), tt.json)
		} else if err = json.Unmarshal(jsonBytes, &jsonTime); err != nil {
			t.Errorf("%v json.Unmarshal error = %v, want nil", tt.time, err)
		} else if !equalTimeAndZone(jsonTime, tt.time) {
			t.Errorf("Unmarshaled time = %v, want %v", jsonTime, tt.time)
		}
	}
}

func TestUnmarshalInvalidTimes(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{`{}`, "Time.UnmarshalJSON: input is not a JSON string"},
		{`[]`, "Time.UnmarshalJSON: input is not a JSON string"},
		{`"2000-01-01T1:12:34Z"`, `<nil>`},
		{`"2000-01-01T00:00:00,000Z"`, `<nil>`},
		{`"2000-01-01T00:00:00+24:00"`, `<nil>`},
		{`"2000-01-01T00:00:00+00:60"`, `<nil>`},
		{`"2000-01-01T00:00:00+123:45"`, `parsing time "2000-01-01T00:00:00+123:45" as "2006-01-02T15:04:05Z07:00": cannot parse "+123:45" as "Z07:00"`},
	}

	for _, tt := range tests {
		var ts Time

		want := tt.want
		err := json.Unmarshal([]byte(tt.in), &ts)
		if fmt.Sprint(err) != want {
			t.Errorf("Time.UnmarshalJSON(%s) = %v, want %v", tt.in, err, want)
		}

		if strings.HasPrefix(tt.in, `"`) && strings.HasSuffix(tt.in, `"`) {
			err = ts.UnmarshalText([]byte(strings.Trim(tt.in, `"`)))
			if fmt.Sprint(err) != want {
				t.Errorf("Time.UnmarshalText(%s) = %v, want %v", tt.in, err, want)
			}
		}
	}
}

func TestMarshalInvalidTimes(t *testing.T) {
	tests := []struct {
		time Time
		want string
	}{
		{Date(10000, 1, 1, 0, 0, 0, 0, UTC), "Time.MarshalJSON: year outside of range [0,9999]"},
		{Date(-998, 1, 1, 0, 0, 0, 0, UTC).Add(-Second), "Time.MarshalJSON: year outside of range [0,9999]"},
		{Date(0, 1, 1, 0, 0, 0, 0, UTC).Add(-Nanosecond), "Time.MarshalJSON: year outside of range [0,9999]"},
		{Date(2020, 1, 1, 0, 0, 0, 0, FixedZone("", 24*60*60)), "Time.MarshalJSON: timezone hour outside of range [0,23]"},
		{Date(2020, 1, 1, 0, 0, 0, 0, FixedZone("", 123*60*60)), "Time.MarshalJSON: timezone hour outside of range [0,23]"},
	}

	for _, tt := range tests {
		want := tt.want
		b, err := tt.time.MarshalJSON()
		switch {
		case b != nil:
			t.Errorf("(%v).MarshalText() = %q, want nil", tt.time, b)
		case err == nil || err.Error() != want:
			t.Errorf("(%v).MarshalJSON() error = %v, want %v", tt.time, err, want)
		}

		want = strings.ReplaceAll(tt.want, "JSON", "Text")
		b, err = tt.time.MarshalText()
		switch {
		case b != nil:
			t.Errorf("(%v).MarshalText() = %q, want nil", tt.time, b)
		case err == nil || err.Error() != want:
			t.Errorf("(%v).MarshalText() error = %v, want %v", tt.time, err, want)
		}
	}
}

var parseDurationTests = []struct {
	in   string
	want Duration
}{
	// simple
	{"0", 0},
	{"5s", 5 * Second},
	{"30s", 30 * Second},
	{"1478s", 1478 * Second},
	// sign
	{"-5s", -5 * Second},
	{"+5s", 5 * Second},
	{"-0", 0},
	{"+0", 0},
	// decimal
	{"5.0s", 5 * Second},
	{"5.6s", 5*Second + 600*Millisecond},
	{"5.s", 5 * Second},
	{".5s", 500 * Millisecond},
	{"1.0s", 1 * Second},
	{"1.00s", 1 * Second},
	{"1.004s", 1*Second + 4*Millisecond},
	{"1.0040s", 1*Second + 4*Millisecond},
	{"100.00100s", 100*Second + 1*Millisecond},
	// different units
	{"10ns", 10 * Nanosecond},
	{"11us", 11 * Microsecond},
	{"12µs", 12 * Microsecond}, // U+00B5
	{"12μs", 12 * Microsecond}, // U+03BC
	{"13ms", 13 * Millisecond},
	{"14s", 14 * Second},
	{"15m", 15 * Minute},
	{"16h", 16 * Hour},
	// composite durations
	{"3h30m", 3*Hour + 30*Minute},
	{"10.5s4m", 4*Minute + 10*Second + 500*Millisecond},
	{"-2m3.4s", -(2*Minute + 3*Second + 400*Millisecond)},
	{"1h2m3s4ms5us6ns", 1*Hour + 2*Minute + 3*Second + 4*Millisecond + 5*Microsecond + 6*Nanosecond},
	{"39h9m14.425s", 39*Hour + 9*Minute + 14*Second + 425*Millisecond},
	// large value
	{"52763797000ns", 52763797000 * Nanosecond},
	// 小数点后有超过9位数字，参见 https://golang.org/issue/6617. md5:d67cd94bd42d243f
	{"0.3333333333333333333h", 20 * Minute},
	// 9007199254740993 = 1 左移 53 位再加上 1，这个数值无法精确地存储在一个 float64 类型中。. md5:c17cfac2fa94c135
	{"9007199254740993ns", (1<<53 + 1) * Nanosecond},
	// 可以用int64表示的最大纳秒时间间隔. md5:0e44b710b3398879
	{"9223372036854775807ns", (1<<63 - 1) * Nanosecond},
	{"9223372036854775.807us", (1<<63 - 1) * Nanosecond},
	{"9223372036s854ms775us807ns", (1<<63 - 1) * Nanosecond},
	{"-9223372036854775808ns", -1 << 63 * Nanosecond},
	{"-9223372036854775.808us", -1 << 63 * Nanosecond},
	{"-9223372036s854ms775us808ns", -1 << 63 * Nanosecond},
	// largest negative value
	{"-9223372036854775808ns", -1 << 63 * Nanosecond},
	// 最大的负回环值，参见：https://golang.org/issue/48629. md5:5b191ebe8da8d694
	{"-2562047h47m16.854775808s", -1 << 63 * Nanosecond},
	// 大字符串；问题15011。. md5:07e8867ae8029c1f
	{"0.100000000000000000000h", 6 * Minute},
	// 此值用于测试leadingFraction中第一个溢出检查。. md5:0e26e39dfc483cd4
	{"0.830103483285477580700h", 49*Minute + 48*Second + 372539827*Nanosecond},
}

func TestParseDuration(t *testing.T) {
	for _, tc := range parseDurationTests {
		d, err := ParseDuration(tc.in)
		if err != nil || d != tc.want {
			t.Errorf("ParseDuration(%q) = %v, %v, want %v, nil", tc.in, d, err, tc.want)
		}
	}
}

var parseDurationErrorTests = []struct {
	in     string
	expect string
}{
	// invalid
	{"", `""`},
	{"3", `"3"`},
	{"-", `"-"`},
	{"s", `"s"`},
	{".", `"."`},
	{"-.", `"-."`},
	{".s", `".s"`},
	{"+.s", `"+.s"`},
	{"1d", `"1d"`},
	{"\x85\x85", `"\x85\x85"`},
	{"\xffff", `"\xffff"`},
	{"hello \xffff world", `"hello \xffff world"`},
	{"\uFFFD", `"\xef\xbf\xbd"`},                                             // utf8.RuneError
	{"\uFFFD hello \uFFFD world", `"\xef\xbf\xbd hello \xef\xbf\xbd world"`}, // utf8.RuneError
	// overflow
	{"9223372036854775810ns", `"9223372036854775810ns"`},
	{"9223372036854775808ns", `"9223372036854775808ns"`},
	{"-9223372036854775809ns", `"-9223372036854775809ns"`},
	{"9223372036854776us", `"9223372036854776us"`},
	{"3000000h", `"3000000h"`},
	{"9223372036854775.808us", `"9223372036854775.808us"`},
	{"9223372036854ms775us808ns", `"9223372036854ms775us808ns"`},
}

func TestParseDurationErrors(t *testing.T) {
	for _, tc := range parseDurationErrorTests {
		_, err := ParseDuration(tc.in)
		if err == nil {
			t.Errorf("ParseDuration(%q) = _, nil, want _, non-nil", tc.in)
		} else if !strings.Contains(err.Error(), tc.expect) {
			t.Errorf("ParseDuration(%q) = _, %q, error does not contain %q", tc.in, err, tc.expect)
		}
	}
}

func TestParseDurationRoundTrip(t *testing.T) {
	// https://golang.org/issue/48629 该问题链接指向Go语言官方仓库的一个问题讨论或报告。. md5:0add436cff6ef0f2
	max0 := Duration(math.MaxInt64)
	max1, err := ParseDuration(max0.String())
	if err != nil || max0 != max1 {
		t.Errorf("round-trip failed: %d => %q => %d, %v", max0, max0.String(), max1, err)
	}

	min0 := Duration(math.MinInt64)
	min1, err := ParseDuration(min0.String())
	if err != nil || min0 != min1 {
		t.Errorf("round-trip failed: %d => %q => %d, %v", min0, min0.String(), min1, err)
	}

	for i := 0; i < 100; i++ {
		// 毫秒级以下的分辨率会导致不精确的往返。
		// md5:77015c25f30f6270
		d0 := Duration(rand.Int31()) * Millisecond
		s := d0.String()
		d1, err := ParseDuration(s)
		if err != nil || d0 != d1 {
			t.Errorf("round-trip failed: %d => %q => %d, %v", d0, s, d1, err)
		}
	}
}

//// golang.org/issue/4622
//func TestLocationRace(t *testing.T) {
//	ResetLocalOnceForTest() // 将Once重置以触发竞争条件. md5:f7940df75a562e17
//
//	c := make(chan string, 1)
//	go func() {
//		c <- Now().String()
//	}()
//	_ = Now().String()
//	<-c
//	Sleep(100 * Millisecond)
//
//	// 返回洛杉矶进行后续测试：. md5:97e94316d288fb0e
//	ForceUSPacificForTesting()
//}

var (
	t Time
	u int64
)

var mallocTest = []struct {
	count int
	desc  string
	fn    func()
}{
	{0, `time.Now()`, func() { t = Now() }},
	{0, `time.Now().UnixNano()`, func() { u = Now().UnixNano() }},
	{0, `time.Now().UnixMilli()`, func() { u = Now().UnixMilli() }},
	{0, `time.Now().UnixMicro()`, func() { u = Now().UnixMicro() }},
}

func TestCountMallocs(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping malloc count in short mode")
	}
	if runtime.GOMAXPROCS(0) > 1 {
		t.Skip("skipping; GOMAXPROCS>1")
	}
	for _, mt := range mallocTest {
		allocs := int(testing.AllocsPerRun(100, mt.fn))
		if allocs > mt.count {
			t.Errorf("%s: %d allocs, want %d", mt.desc, allocs, mt.count)
		}
	}
}

func TestLoadFixed(t *testing.T) {
	// 问题4064：处理没有时区转换的地理位置。. md5:39d7da51684886db
	loc, err := LoadLocation("Etc/GMT+1")
	if err != nil {
		t.Fatal(err)
	}

	// Etc/GMT+1在tzdata中使用"东为负"的约定，但Go和其他大多数系统使用"东为正"。因此，GMT+1在Go时区中对应的是-3600秒，而不是+3600秒。
	// md5:ba8d98a3734e5d88
	name, offset := Now().In(loc).Zone()
	// 自tzdata-2016g以来，时区缩写为"-01"，早期版本为"GMT+1"；我们接受两者。（问题#17276）。
	// md5:6d7d3da8d78f3359
	if !(name == "GMT+1" || name == "-01") || offset != -1*60*60 {
		t.Errorf("Now().In(loc).Zone() = %q, %d, want %q or %q, %d",
			name, offset, "GMT+1", "-01", -1*60*60)
	}
}

const (
	minDuration Duration = -1 << 63
	maxDuration Duration = 1<<63 - 1
)

//var subTests = []struct {
//	t Time
//	u Time
//	d Duration
//}{
//	{Time{}, Time{}, Duration(0)},
//	{Date(2009, 11, 23, 0, 0, 0, 1, UTC), Date(2009, 11, 23, 0, 0, 0, 0, UTC), Duration(1)},
//	{Date(2009, 11, 23, 0, 0, 0, 0, UTC), Date(2009, 11, 24, 0, 0, 0, 0, UTC), -24 * Hour},
//	{Date(2009, 11, 24, 0, 0, 0, 0, UTC), Date(2009, 11, 23, 0, 0, 0, 0, UTC), 24 * Hour},
//	{Date(-2009, 11, 24, 0, 0, 0, 0, UTC), Date(-2009, 11, 23, 0, 0, 0, 0, UTC), 24 * Hour},
//	{Time{}, Date(2109, 11, 23, 0, 0, 0, 0, UTC), minDuration},
//	{Date(2109, 11, 23, 0, 0, 0, 0, UTC), Time{}, maxDuration},
//	{Time{}, Date(-2109, 11, 23, 0, 0, 0, 0, UTC), maxDuration},
//	{Date(-2109, 11, 23, 0, 0, 0, 0, UTC), Time{}, minDuration},
//	{Date(2290, 1, 1, 0, 0, 0, 0, UTC), Date(2000, 1, 1, 0, 0, 0, 0, UTC), 290*365*24*Hour + 71*24*Hour},
//	{Date(2300, 1, 1, 0, 0, 0, 0, UTC), Date(2000, 1, 1, 0, 0, 0, 0, UTC), maxDuration},
//	{Date(2000, 1, 1, 0, 0, 0, 0, UTC), Date(2290, 1, 1, 0, 0, 0, 0, UTC), -290*365*24*Hour - 71*24*Hour},
//	{Date(2000, 1, 1, 0, 0, 0, 0, UTC), Date(2300, 1, 1, 0, 0, 0, 0, UTC), minDuration},
//	{Date(2311, 11, 26, 02, 16, 47, 63535996, UTC), Date(2019, 8, 16, 2, 29, 30, 268436582, UTC), 9223372036795099414},
//	{MinMonoTime, MaxMonoTime, minDuration},
//	{MaxMonoTime, MinMonoTime, maxDuration},
//}

//func TestSub(t *testing.T) {
//	for i, st := range subTests {
//		got := st.t.Sub(st.u)
//		if got != st.d {
//			t.Errorf("#%d: Sub(%v, %v): got %v; want %v", i, st.t, st.u, got, st.d)
//		}
//	}
//}

var nsDurationTests = []struct {
	d    Duration
	want int64
}{
	{Duration(-1000), -1000},
	{Duration(-1), -1},
	{Duration(1), 1},
	{Duration(1000), 1000},
}

func TestDurationNanoseconds(t *testing.T) {
	for _, tt := range nsDurationTests {
		if got := tt.d.Nanoseconds(); got != tt.want {
			t.Errorf("Duration(%s).Nanoseconds() = %d; want: %d", tt.d, got, tt.want)
		}
	}
}

var usDurationTests = []struct {
	d    Duration
	want int64
}{
	{Duration(-1000), -1},
	{Duration(1000), 1},
}

func TestDurationMicroseconds(t *testing.T) {
	for _, tt := range usDurationTests {
		if got := tt.d.Microseconds(); got != tt.want {
			t.Errorf("Duration(%s).Microseconds() = %d; want: %d", tt.d, got, tt.want)
		}
	}
}

var msDurationTests = []struct {
	d    Duration
	want int64
}{
	{Duration(-1000000), -1},
	{Duration(1000000), 1},
}

func TestDurationMilliseconds(t *testing.T) {
	for _, tt := range msDurationTests {
		if got := tt.d.Milliseconds(); got != tt.want {
			t.Errorf("Duration(%s).Milliseconds() = %d; want: %d", tt.d, got, tt.want)
		}
	}
}

var secDurationTests = []struct {
	d    Duration
	want float64
}{
	{Duration(300000000), 0.3},
}

func TestDurationSeconds(t *testing.T) {
	for _, tt := range secDurationTests {
		if got := tt.d.Seconds(); got != tt.want {
			t.Errorf("Duration(%s).Seconds() = %g; want: %g", tt.d, got, tt.want)
		}
	}
}

var minDurationTests = []struct {
	d    Duration
	want float64
}{
	{Duration(-60000000000), -1},
	{Duration(-1), -1 / 60e9},
	{Duration(1), 1 / 60e9},
	{Duration(60000000000), 1},
	{Duration(3000), 5e-8},
}

func TestDurationMinutes(t *testing.T) {
	for _, tt := range minDurationTests {
		if got := tt.d.Minutes(); got != tt.want {
			t.Errorf("Duration(%s).Minutes() = %g; want: %g", tt.d, got, tt.want)
		}
	}
}

var hourDurationTests = []struct {
	d    Duration
	want float64
}{
	{Duration(-3600000000000), -1},
	{Duration(-1), -1 / 3600e9},
	{Duration(1), 1 / 3600e9},
	{Duration(3600000000000), 1},
	{Duration(36), 1e-11},
}

func TestDurationHours(t *testing.T) {
	for _, tt := range hourDurationTests {
		if got := tt.d.Hours(); got != tt.want {
			t.Errorf("Duration(%s).Hours() = %g; want: %g", tt.d, got, tt.want)
		}
	}
}

var durationTruncateTests = []struct {
	d    Duration
	m    Duration
	want Duration
}{
	{0, Second, 0},
	{Minute, -7 * Second, Minute},
	{Minute, 0, Minute},
	{Minute, 1, Minute},
	{Minute + 10*Second, 10 * Second, Minute + 10*Second},
	{2*Minute + 10*Second, Minute, 2 * Minute},
	{10*Minute + 10*Second, 3 * Minute, 9 * Minute},
	{Minute + 10*Second, Minute + 10*Second + 1, 0},
	{Minute + 10*Second, Hour, 0},
	{-Minute, Second, -Minute},
	{-10 * Minute, 3 * Minute, -9 * Minute},
	{-10 * Minute, Hour, 0},
}

func TestDurationTruncate(t *testing.T) {
	for _, tt := range durationTruncateTests {
		if got := tt.d.Truncate(tt.m); got != tt.want {
			t.Errorf("Duration(%s).Truncate(%s) = %s; want: %s", tt.d, tt.m, got, tt.want)
		}
	}
}

var durationRoundTests = []struct {
	d    Duration
	m    Duration
	want Duration
}{
	{0, Second, 0},
	{Minute, -11 * Second, Minute},
	{Minute, 0, Minute},
	{Minute, 1, Minute},
	{2 * Minute, Minute, 2 * Minute},
	{2*Minute + 10*Second, Minute, 2 * Minute},
	{2*Minute + 30*Second, Minute, 3 * Minute},
	{2*Minute + 50*Second, Minute, 3 * Minute},
	{-Minute, 1, -Minute},
	{-2 * Minute, Minute, -2 * Minute},
	{-2*Minute - 10*Second, Minute, -2 * Minute},
	{-2*Minute - 30*Second, Minute, -3 * Minute},
	{-2*Minute - 50*Second, Minute, -3 * Minute},
	{8e18, 3e18, 9e18},
	{9e18, 5e18, 1<<63 - 1},
	{-8e18, 3e18, -9e18},
	{-9e18, 5e18, -1 << 63},
	{3<<61 - 1, 3 << 61, 3 << 61},
}

func TestDurationRound(t *testing.T) {
	for _, tt := range durationRoundTests {
		if got := tt.d.Round(tt.m); got != tt.want {
			t.Errorf("Duration(%s).Round(%s) = %s; want: %s", tt.d, tt.m, got, tt.want)
		}
	}
}

var durationAbsTests = []struct {
	d    Duration
	want Duration
}{
	{0, 0},
	{1, 1},
	{-1, 1},
	{1 * Minute, 1 * Minute},
	{-1 * Minute, 1 * Minute},
	{minDuration, maxDuration},
	{minDuration + 1, maxDuration},
	{minDuration + 2, maxDuration - 1},
	{maxDuration, maxDuration},
	{maxDuration - 1, maxDuration - 1},
}

func TestDurationAbs(t *testing.T) {
	for _, tt := range durationAbsTests {
		if got := tt.d.Abs(); got != tt.want {
			t.Errorf("Duration(%s).Abs() = %s; want: %s", tt.d, got, tt.want)
		}
	}
}

var defaultLocTests = []struct {
	name string
	f    func(t1, t2 Time) bool
}{
	{"After", func(t1, t2 Time) bool { return t1.After(t2) == t2.After(t1) }},
	{"Before", func(t1, t2 Time) bool { return t1.Before(t2) == t2.Before(t1) }},
	{"Equal", func(t1, t2 Time) bool { return t1.Equal(t2) == t2.Equal(t1) }},
	{"Compare", func(t1, t2 Time) bool { return t1.Compare(t2) == t2.Compare(t1) }},

	{"IsZero", func(t1, t2 Time) bool { return t1.IsZero() == t2.IsZero() }},
	{"Date", func(t1, t2 Time) bool {
		a1, b1, c1 := t1.Date()
		a2, b2, c2 := t2.Date()
		return a1 == a2 && b1 == b2 && c1 == c2
	}},
	{"Year", func(t1, t2 Time) bool { return t1.Year() == t2.Year() }},
	{"Month", func(t1, t2 Time) bool { return t1.Month() == t2.Month() }},
	{"Day", func(t1, t2 Time) bool { return t1.Day() == t2.Day() }},
	{"Weekday", func(t1, t2 Time) bool { return t1.Weekday() == t2.Weekday() }},
	{"ISOWeek", func(t1, t2 Time) bool {
		a1, b1 := t1.ISOWeek()
		a2, b2 := t2.ISOWeek()
		return a1 == a2 && b1 == b2
	}},
	{"Clock", func(t1, t2 Time) bool {
		a1, b1, c1 := t1.Clock()
		a2, b2, c2 := t2.Clock()
		return a1 == a2 && b1 == b2 && c1 == c2
	}},
	{"Hour", func(t1, t2 Time) bool { return t1.Hour() == t2.Hour() }},
	{"Minute", func(t1, t2 Time) bool { return t1.Minute() == t2.Minute() }},
	{"Second", func(t1, t2 Time) bool { return t1.Second() == t2.Second() }},
	{"Nanosecond", func(t1, t2 Time) bool { return t1.Hour() == t2.Hour() }},
	{"YearDay", func(t1, t2 Time) bool { return t1.YearDay() == t2.YearDay() }},

	// 使用Equal是因为Add操作不会修改loc，如果使用"=="比较会导致错误. md5:3622609b0185b9c0
	{"Add", func(t1, t2 Time) bool { return t1.Add(Hour).Equal(t2.Add(Hour)) }},
	{"Sub", func(t1, t2 Time) bool { return t1.Sub(t2) == t2.Sub(t1) }},

	//此测试用例的原始原因：bug 15852. md5:a83da26ab0c43f24
	{"AddDate", func(t1, t2 Time) bool { return t1.AddDate(1991, 9, 3) == t2.AddDate(1991, 9, 3) }},

	{"UTC", func(t1, t2 Time) bool { return t1.UTC() == t2.UTC() }},
	{"Local", func(t1, t2 Time) bool { return t1.Local() == t2.Local() }},
	{"In", func(t1, t2 Time) bool { return t1.In(UTC) == t2.In(UTC) }},

	{"Local", func(t1, t2 Time) bool { return t1.Local() == t2.Local() }},
	{"Zone", func(t1, t2 Time) bool {
		a1, b1 := t1.Zone()
		a2, b2 := t2.Zone()
		return a1 == a2 && b1 == b2
	}},

	{"Unix", func(t1, t2 Time) bool { return t1.Unix() == t2.Unix() }},
	{"UnixNano", func(t1, t2 Time) bool { return t1.UnixNano() == t2.UnixNano() }},
	{"UnixMilli", func(t1, t2 Time) bool { return t1.UnixMilli() == t2.UnixMilli() }},
	{"UnixMicro", func(t1, t2 Time) bool { return t1.UnixMicro() == t2.UnixMicro() }},

	{"MarshalBinary", func(t1, t2 Time) bool {
		a1, b1 := t1.MarshalBinary()
		a2, b2 := t2.MarshalBinary()
		return bytes.Equal(a1, a2) && b1 == b2
	}},
	{"GobEncode", func(t1, t2 Time) bool {
		a1, b1 := t1.GobEncode()
		a2, b2 := t2.GobEncode()
		return bytes.Equal(a1, a2) && b1 == b2
	}},
	{"MarshalJSON", func(t1, t2 Time) bool {
		a1, b1 := t1.MarshalJSON()
		a2, b2 := t2.MarshalJSON()
		return bytes.Equal(a1, a2) && b1 == b2
	}},
	{"MarshalText", func(t1, t2 Time) bool {
		a1, b1 := t1.MarshalText()
		a2, b2 := t2.MarshalText()
		return bytes.Equal(a1, a2) && b1 == b2
	}},

	{"Truncate", func(t1, t2 Time) bool { return t1.Truncate(Hour).Equal(t2.Truncate(Hour)) }},
	{"Round", func(t1, t2 Time) bool { return t1.Round(Hour).Equal(t2.Round(Hour)) }},

	{"== Time{}", func(t1, t2 Time) bool { return (t1 == Time{}) == (t2 == Time{}) }},
}

func TestDefaultLoc(t *testing.T) {
	// 验证如果loc被设置为nil或UTC，Time的所有方法行为是否相同。
	// md5:ffb3babe06d51eb3
	for _, tt := range defaultLocTests {
		t1 := Time{}
		t2 := Time{}.UTC()
		if !tt.f(t1, t2) {
			t.Errorf("Time{} and Time{}.UTC() behave differently for %s", tt.name)
		}
	}
}

func BenchmarkNow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		t = Now()
	}
}

func BenchmarkNowUnixNano(b *testing.B) {
	for i := 0; i < b.N; i++ {
		u = Now().UnixNano()
	}
}

func BenchmarkNowUnixMilli(b *testing.B) {
	for i := 0; i < b.N; i++ {
		u = Now().UnixMilli()
	}
}

func BenchmarkNowUnixMicro(b *testing.B) {
	for i := 0; i < b.N; i++ {
		u = Now().UnixMicro()
	}
}

func BenchmarkFormat(b *testing.B) {
	t := Unix(1265346057, 0)
	for i := 0; i < b.N; i++ {
		t.Format("Mon Jan  2 15:04:05 2006")
	}
}

func BenchmarkFormatRFC3339(b *testing.B) {
	t := Unix(1265346057, 0)
	for i := 0; i < b.N; i++ {
		t.Format("2006-01-02T15:04:05Z07:00")
	}
}

func BenchmarkFormatRFC3339Nano(b *testing.B) {
	t := Unix(1265346057, 0)
	for i := 0; i < b.N; i++ {
		t.Format("2006-01-02T15:04:05.999999999Z07:00")
	}
}

func BenchmarkFormatNow(b *testing.B) {
	// 类似 BenchmarkFormat，但是更简单，因为时区查找缓存是针对当前优化的。
	// md5:f06dd1b142ed8305
	t := Now()
	for i := 0; i < b.N; i++ {
		t.Format("Mon Jan  2 15:04:05 2006")
	}
}

func BenchmarkMarshalJSON(b *testing.B) {
	t := Now()
	for i := 0; i < b.N; i++ {
		t.MarshalJSON()
	}
}

func BenchmarkMarshalText(b *testing.B) {
	t := Now()
	for i := 0; i < b.N; i++ {
		t.MarshalText()
	}
}

func BenchmarkParse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Parse(ANSIC, "Mon Jan  2 15:04:05 2006")
	}
}

const testdataRFC3339UTC = "2020-08-22T11:27:43.123456789Z"

func BenchmarkParseRFC3339UTC(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Parse(RFC3339, testdataRFC3339UTC)
	}
}

var testdataRFC3339UTCBytes = []byte(testdataRFC3339UTC)

func BenchmarkParseRFC3339UTCBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Parse(RFC3339, string(testdataRFC3339UTCBytes))
	}
}

const testdataRFC3339TZ = "2020-08-22T11:27:43.123456789-02:00"

func BenchmarkParseRFC3339TZ(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Parse(RFC3339, testdataRFC3339TZ)
	}
}

var testdataRFC3339TZBytes = []byte(testdataRFC3339TZ)

func BenchmarkParseRFC3339TZBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Parse(RFC3339, string(testdataRFC3339TZBytes))
	}
}

func BenchmarkParseDuration(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ParseDuration("9007199254.740993ms")
		ParseDuration("9007199254740993ns")
	}
}

func BenchmarkHour(b *testing.B) {
	t := Now()
	for i := 0; i < b.N; i++ {
		_ = t.Hour()
	}
}

func BenchmarkSecond(b *testing.B) {
	t := Now()
	for i := 0; i < b.N; i++ {
		_ = t.Second()
	}
}

func BenchmarkYear(b *testing.B) {
	t := Now()
	for i := 0; i < b.N; i++ {
		_ = t.Year()
	}
}

func BenchmarkDay(b *testing.B) {
	t := Now()
	for i := 0; i < b.N; i++ {
		_ = t.Day()
	}
}

func BenchmarkISOWeek(b *testing.B) {
	t := Now()
	for i := 0; i < b.N; i++ {
		_, _ = t.ISOWeek()
	}
}

func BenchmarkGoString(b *testing.B) {
	t := Now()
	for i := 0; i < b.N; i++ {
		_ = t.GoString()
	}
}

func BenchmarkUnmarshalText(b *testing.B) {
	var t Time
	in := []byte("2020-08-22T11:27:43.123456789-02:00")
	for i := 0; i < b.N; i++ {
		t.UnmarshalText(in)
	}
}

func TestMarshalBinaryZeroTime(t *testing.T) {
	t0 := Time{}
	enc, err := t0.MarshalBinary()
	if err != nil {
		t.Fatal(err)
	}
	t1 := Now() // not zero
	if err := t1.UnmarshalBinary(enc); err != nil {
		t.Fatal(err)
	}
	if t1 != t0 {
		t.Errorf("t0=%#v\nt1=%#v\nwant identical structures", t0, t1)
	}
}

func TestMarshalBinaryVersion2(t *testing.T) {
	t0, err := Parse(RFC3339, "1880-01-01T00:00:00Z")
	if err != nil {
		t.Errorf("Failed to parse time, error = %v", err)
	}
	loc, err := LoadLocation("US/Eastern")
	if err != nil {
		t.Errorf("Failed to load location, error = %v", err)
	}
	t1 := t0.In(loc)
	b, err := t1.MarshalBinary()
	if err != nil {
		t.Errorf("Failed to Marshal, error = %v", err)
	}

	t2 := Time{}
	err = t2.UnmarshalBinary(b)
	if err != nil {
		t.Errorf("Failed to Unmarshal, error = %v", err)
	}

	if !(t0.Equal(t1) && t1.Equal(t2)) {
		if !t0.Equal(t1) {
			t.Errorf("The result t1: %+v after Marshal is not matched original t0: %+v", t1, t0)
		}
		if !t1.Equal(t2) {
			t.Errorf("The result t2: %+v after Unmarshal is not matched original t1: %+v", t2, t1)
		}
	}
}

func TestUnmarshalTextAllocations(t *testing.T) {
	in := []byte(testdataRFC3339UTC) // 足够短以在栈上分配. md5:195247dda0133cd7
	if allocs := testing.AllocsPerRun(100, func() {
		var t Time
		t.UnmarshalText(in)
	}); allocs != 0 {
		t.Errorf("got %v allocs, want 0 allocs", allocs)
	}
}

// 问题17720：time.Month的零值无法正常打印. md5:db136c905142b832
func TestZeroMonthString(t *testing.T) {
	if got, want := Month(0).String(), "%!Month(0)"; got != want {
		t.Errorf("zero month = %q; want %q", got, want)
	}
}

// 问题 24692：超出范围的工作日恐慌. md5:20b30cd630c70a9f
func TestWeekdayString(t *testing.T) {
	if got, want := Tuesday.String(), "Tuesday"; got != want {
		t.Errorf("Tuesday weekday = %q; want %q", got, want)
	}
	if got, want := Weekday(14).String(), "%!Weekday(14)"; got != want {
		t.Errorf("14th weekday = %q; want %q", got, want)
	}
}

//func TestReadFileLimit(t *testing.T) {
//	const zero = "/dev/zero"
//	if _, err := os.Stat(zero); err != nil {
//		t.Skip("skipping test without a /dev/zero")
//	}
//	_, err := ReadFile(zero)
//	if err == nil || !strings.Contains(err.Error(), "is too large") {
//		t.Errorf("readFile(%q) error = %v; want error containing 'is too large'", zero, err)
//	}
//}

// Issue 25686：并发访问定时器时的硬崩溃。
// Issue 37400：使用定时器时出现“竞态条件”导致panic。
// 此测试故意触发竞态条件。
// 我们正在测试在持有锁的情况下不出现“致命错误：panic”，并且也不会panic。
// md5:62f83f45c1a23d84
func TestConcurrentTimerReset(t *testing.T) {
	const goroutines = 8
	const tries = 1000
	var wg sync.WaitGroup
	wg.Add(goroutines)
	timer := NewTimer(Hour)
	for i := 0; i < goroutines; i++ {
		go func(i int) {
			defer wg.Done()
			for j := 0; j < tries; j++ {
				timer.Reset(Hour + Duration(i*j))
			}
		}(i)
	}
	wg.Wait()
}

// Issue 37400: 因“计时器的竞态使用”导致的 panic。. md5:256cc7c7f68e7bb5
func TestConcurrentTimerResetStop(t *testing.T) {
	const goroutines = 8
	const tries = 1000
	var wg sync.WaitGroup
	wg.Add(goroutines * 2)
	timer := NewTimer(Hour)
	for i := 0; i < goroutines; i++ {
		go func(i int) {
			defer wg.Done()
			for j := 0; j < tries; j++ {
				timer.Reset(Hour + Duration(i*j))
			}
		}(i)
		go func(i int) {
			defer wg.Done()
			timer.Stop()
		}(i)
	}
	wg.Wait()
}

//
//func TestTimeIsDST(t *testing.T) {
//	undo := DisablePlatformSources()
//	defer undo()
//
//	tzWithDST, err := LoadLocation("Australia/Sydney")
//	if err != nil {
//		t.Fatalf("could not load tz 'Australia/Sydney': %v", err)
//	}
//	tzWithoutDST, err := LoadLocation("Australia/Brisbane")
//	if err != nil {
//		t.Fatalf("could not load tz 'Australia/Brisbane': %v", err)
//	}
//	tzFixed := FixedZone("FIXED_TIME", 12345)
//
//	tests := [...]struct {
//		time Time
//		want bool
//	}{
//		0: {Date(2009, 1, 1, 12, 0, 0, 0, UTC), false},
//		1: {Date(2009, 6, 1, 12, 0, 0, 0, UTC), false},
//		2: {Date(2009, 1, 1, 12, 0, 0, 0, tzWithDST), true},
//		3: {Date(2009, 6, 1, 12, 0, 0, 0, tzWithDST), false},
//		4: {Date(2009, 1, 1, 12, 0, 0, 0, tzWithoutDST), false},
//		5: {Date(2009, 6, 1, 12, 0, 0, 0, tzWithoutDST), false},
//		6: {Date(2009, 1, 1, 12, 0, 0, 0, tzFixed), false},
//		7: {Date(2009, 6, 1, 12, 0, 0, 0, tzFixed), false},
//	}
//
//	for i, tt := range tests {
//		got := tt.time.IsDST()
//		if got != tt.want {
//			t.Errorf("#%d:: (%#v).IsDST()=%t, want %t", i, tt.time.Format(RFC3339), got, tt.want)
//		}
//	}
//}
//
//func TestTimeAddSecOverflow(t *testing.T) {
//	// 使用正向的delta进行测试。. md5:4de1a0f28de88248
//	var maxInt64 int64 = 1<<63 - 1
//	timeExt := maxInt64 - UnixToInternal - 50
//	notMonoTime := Unix(timeExt, 0)
//	for i := int64(0); i < 100; i++ {
//		sec := notMonoTime.Unix()
//		notMonoTime = notMonoTime.Add(Duration(i * 1e9))
//		if newSec := notMonoTime.Unix(); newSec != sec+i && newSec+UnixToInternal != maxInt64 {
//			t.Fatalf("time ext: %d overflows with positive delta, overflow threshold: %d", newSec, maxInt64)
//		}
//	}
//
//	// 使用负的delta值进行测试。. md5:e045f2e98546ccfc
//	maxInt64 = -maxInt64
//	notMonoTime = NotMonoNegativeTime
//	for i := int64(0); i > -100; i-- {
//		sec := notMonoTime.Unix()
//		notMonoTime = notMonoTime.Add(Duration(i * 1e9))
//		if newSec := notMonoTime.Unix(); newSec != sec+i && newSec+UnixToInternal != maxInt64 {
//			t.Fatalf("time ext: %d overflows with positive delta, overflow threshold: %d", newSec, maxInt64)
//		}
//	}
//}
//
//// 问题 49284：由于夏令时，time.ParseInLocation 函数解析错误. md5:485ad6976ad20904
//func TestTimeWithZoneTransition(t *testing.T) {
//	undo := DisablePlatformSources()
//	defer undo()
//
//	loc, err := LoadLocation("Asia/Shanghai")
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	tests := [...]struct {
//		give Time
//		want Time
//	}{
//		// 1991年4月14日 - 夏令时开始
//		// 当“亚洲/上海”时间即将到达
//		// 1991年4月14日，星期日，凌晨02:00:00时，时钟向前调整了1小时，
//		// 变为1991年4月14日，星期日，凌晨03:00:00的当地夏令时。
//		// 这时的UTC时间是1991年4月13日，18:00:00。
//		// md5:154cc70a3ad2ed2f
//		0: {Date(1991, April, 13, 17, 50, 0, 0, loc), Date(1991, April, 13, 9, 50, 0, 0, UTC)},
//		1: {Date(1991, April, 13, 18, 0, 0, 0, loc), Date(1991, April, 13, 10, 0, 0, 0, UTC)},
//		2: {Date(1991, April, 14, 1, 50, 0, 0, loc), Date(1991, April, 13, 17, 50, 0, 0, UTC)},
//		3: {Date(1991, April, 14, 3, 0, 0, 0, loc), Date(1991, April, 13, 18, 0, 0, 0, UTC)},
//
//		// 1991年9月15日 - 夏令时结束
//		// 当"Asia/Shanghai"的本地夏令时即将达到
//		// 1991年9月15日，02:00:00 时，时钟回拨一小时至
//		// 1991年9月15日，01:00:00 本地标准时间。
//		// UTC时间为1991年9月14日，17:00:00。
//		// md5:200f61ea8b858f25
//		4: {Date(1991, September, 14, 16, 50, 0, 0, loc), Date(1991, September, 14, 7, 50, 0, 0, UTC)},
//		5: {Date(1991, September, 14, 17, 0, 0, 0, loc), Date(1991, September, 14, 8, 0, 0, 0, UTC)},
//		6: {Date(1991, September, 15, 0, 50, 0, 0, loc), Date(1991, September, 14, 15, 50, 0, 0, UTC)},
//		7: {Date(1991, September, 15, 2, 00, 0, 0, loc), Date(1991, September, 14, 18, 00, 0, 0, UTC)},
//	}
//
//	for i, tt := range tests {
//		if !tt.give.Equal(tt.want) {
//			t.Errorf("#%d:: %#v is not equal to %#v", i, tt.give.Format(RFC3339), tt.want.Format(RFC3339))
//		}
//	}
//}
//
//func TestZoneBounds(t *testing.T) {
//	undo := DisablePlatformSources()
//	defer undo()
//	loc, err := LoadLocation("Asia/Shanghai")
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	// UTC位置的区划边界将只返回两个零Time。. md5:3223c57808e355f0
//	for _, test := range utctests {
//		sec := test.seconds
//		golden := &test.golden
//		tm := Unix(sec, 0).UTC()
//		start, end := tm.ZoneBounds()
//		if !(start.IsZero() && end.IsZero()) {
//			t.Errorf("ZoneBounds of %+v expects two zero Time, got:\n  start=%v\n  end=%v", *golden, start, end)
//		}
//	}
//
//	// 如果时区从时间的开始处开始，start将作为零Time返回。
//	// 使用math.MinInt32可以避免在32位系统上int参数溢出的问题。
//	// md5:b75e6e21de334f47
//	beginTime := Date(math.MinInt32, January, 1, 0, 0, 0, 0, loc)
//	start, end := beginTime.ZoneBounds()
//	if !start.IsZero() || end.IsZero() {
//		t.Errorf("ZoneBounds of %v expects start is zero Time, got:\n  start=%v\n  end=%v", beginTime, start, end)
//	}
//
//	// 如果时区无限延续，end将被返回为一个零时间值。
//	// 在32位系统上，请使用math.MaxInt32以避免int型参数溢出。
//	// md5:f77aba81d8f5afa3
//	foreverTime := Date(math.MaxInt32, January, 1, 0, 0, 0, 0, loc)
//	start, end = foreverTime.ZoneBounds()
//	if start.IsZero() || !end.IsZero() {
//		t.Errorf("ZoneBounds of %v expects end is zero Time, got:\n  start=%v\n  end=%v", foreverTime, start, end)
//	}
//
//	// 检查一些现实世界的情况，确保我们得到正确的边界。. md5:3e7391ac63d3b374
//	boundOne := Date(1990, September, 16, 1, 0, 0, 0, loc)
//	boundTwo := Date(1991, April, 14, 3, 0, 0, 0, loc)
//	boundThree := Date(1991, September, 15, 1, 0, 0, 0, loc)
//	makeLocalTime := func(sec int64) Time { return Unix(sec, 0) }
//	realTests := [...]struct {
//		giveTime  Time
//		wantStart Time
//		wantEnd   Time
//	}{
//		// "Asia/Shanghai"夏令时的区域边界. md5:3ec7dffa23c5942e
//		0: {Date(1991, April, 13, 17, 50, 0, 0, loc), boundOne, boundTwo},
//		1: {Date(1991, April, 13, 18, 0, 0, 0, loc), boundOne, boundTwo},
//		2: {Date(1991, April, 14, 1, 50, 0, 0, loc), boundOne, boundTwo},
//		3: {boundTwo, boundTwo, boundThree},
//		4: {Date(1991, September, 14, 16, 50, 0, 0, loc), boundTwo, boundThree},
//		5: {Date(1991, September, 14, 17, 0, 0, 0, loc), boundTwo, boundThree},
//		6: {Date(1991, September, 15, 0, 50, 0, 0, loc), boundTwo, boundThree},
//
//		// "Asia/Shanghai"时区在最后一次（标准时间）转换之后的边界范围. md5:27f997ff7dc24c8c
//		7:  {boundThree, boundThree, Time{}},
//		8:  {Date(1991, December, 15, 1, 50, 0, 0, loc), boundThree, Time{}},
//		9:  {Date(1992, April, 13, 17, 50, 0, 0, loc), boundThree, Time{}},
//		10: {Date(1992, April, 13, 18, 0, 0, 0, loc), boundThree, Time{}},
//		11: {Date(1992, April, 14, 1, 50, 0, 0, loc), boundThree, Time{}},
//		12: {Date(1992, September, 14, 16, 50, 0, 0, loc), boundThree, Time{}},
//		13: {Date(1992, September, 14, 17, 0, 0, 0, loc), boundThree, Time{}},
//		14: {Date(1992, September, 15, 0, 50, 0, 0, loc), boundThree, Time{}},
//
//		// 获取一个本地时间的ZoneBounds将返回两个本地时间。
//		// 注意：为了测试，我们预先加载了"America/Los_Angeles"作为time.Local。
//		// md5:8bcb336908b3115a
//		15: {makeLocalTime(0), makeLocalTime(-5756400), makeLocalTime(9972000)},
//		16: {makeLocalTime(1221681866), makeLocalTime(1205056800), makeLocalTime(1225616400)},
//		17: {makeLocalTime(2152173599), makeLocalTime(2145916800), makeLocalTime(2152173600)},
//		18: {makeLocalTime(2152173600), makeLocalTime(2152173600), makeLocalTime(2172733200)},
//		19: {makeLocalTime(2152173601), makeLocalTime(2152173600), makeLocalTime(2172733200)},
//		20: {makeLocalTime(2159200800), makeLocalTime(2152173600), makeLocalTime(2172733200)},
//		21: {makeLocalTime(2172733199), makeLocalTime(2152173600), makeLocalTime(2172733200)},
//		22: {makeLocalTime(2172733200), makeLocalTime(2172733200), makeLocalTime(2177452800)},
//	}
//	for i, tt := range realTests {
//		start, end := tt.giveTime.ZoneBounds()
//		if !start.Equal(tt.wantStart) || !end.Equal(tt.wantEnd) {
//			t.Errorf("#%d:: ZoneBounds of %v expects right bounds:\n  got start=%v\n  want start=%v\n  got end=%v\n  want end=%v",
//				i, tt.giveTime, start, tt.wantStart, end, tt.wantEnd)
//		}
//	}
//}
