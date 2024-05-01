// 版权所有 2011 Go 语言作者。保留所有权利。
// 使用此源代码须遵循 BSD 风格的许可协议，
// 许可协议具体内容可在 LICENSE 文件中找到。
// md5:b5bcbe8a534f6077

package time_test

import (
	"fmt"
	"github.com/888go/gosdk/time"
)

func expensiveCall() {}

func ExampleDuration() {
	t0 := time.Now()
	expensiveCall()
	t1 := time.Now()
	fmt.Printf("The call took %v to run.\n", t1.Sub(t0))
}

func ExampleDuration_Round() {
	d, err := time.ParseDuration("1h15m30.918273645s")
	if err != nil {
		panic(err)
	}

	round := []time.Duration{
		time.Nanosecond,
		time.Microsecond,
		time.Millisecond,
		time.Second,
		2 * time.Second,
		time.Minute,
		10 * time.Minute,
		time.Hour,
	}

	for _, r := range round {
		fmt.Printf("d.Round(%6s) = %s\n", r, d.Round(r).String())
	}
	// Output:
	// d.Round(   1ns) = 1h15m30.918273645s
	// d.Round(   1µs) = 1h15m30.918274s
	// d.Round(   1ms) = 1h15m30.918s
	// d.Round(    1s) = 1h15m31s
	// d.Round(    2s) = 1h15m30s
	// d.Round(  1m0s) = 1h16m0s
	// d.Round( 10m0s) = 1h20m0s
	// d.Round(1h0m0s) = 1h0m0s
}

func ExampleDuration_String() {
	fmt.Println(1*time.Hour + 2*time.Minute + 300*time.Millisecond)
	fmt.Println(300 * time.Millisecond)
	// Output:
	// 1h2m0.3s
	// 300ms
}

func ExampleDuration_Truncate() {
	d, err := time.ParseDuration("1h15m30.918273645s")
	if err != nil {
		panic(err)
	}

	trunc := []time.Duration{
		time.Nanosecond,
		time.Microsecond,
		time.Millisecond,
		time.Second,
		2 * time.Second,
		time.Minute,
		10 * time.Minute,
		time.Hour,
	}

	for _, t := range trunc {
		fmt.Printf("d.Truncate(%6s) = %s\n", t, d.Truncate(t).String())
	}
	// Output:
	// d.Truncate(   1ns) = 1h15m30.918273645s
	// d.Truncate(   1µs) = 1h15m30.918273s
	// d.Truncate(   1ms) = 1h15m30.918s
	// d.Truncate(    1s) = 1h15m30s
	// d.Truncate(    2s) = 1h15m30s
	// d.Truncate(  1m0s) = 1h15m0s
	// d.Truncate( 10m0s) = 1h10m0s
	// d.Truncate(1h0m0s) = 1h0m0s
}

func ExampleParseDuration() {
	hours, _ := time.ParseDuration("10h")
	complex, _ := time.ParseDuration("1h10m10s")
	micro, _ := time.ParseDuration("1µs")
	// 此包也接受常见的但不正确的前缀 u 代表微（micro）。. md5:3e19ea04de8a976c
	micro2, _ := time.ParseDuration("1us")

	fmt.Println(hours)
	fmt.Println(complex)
	fmt.Printf("There are %.0f seconds in %v.\n", complex.Seconds(), complex)
	fmt.Printf("There are %d nanoseconds in %v.\n", micro.Nanoseconds(), micro)
	fmt.Printf("There are %6.2e seconds in %v.\n", micro2.Seconds(), micro)
	// Output:
	// 10h0m0s
	// 1h10m10s
	// There are 4210 seconds in 1h10m10s.
	// There are 1000 nanoseconds in 1µs.
	// There are 1.00e-06 seconds in 1µs.
}

func ExampleDuration_Hours() {
	h, _ := time.ParseDuration("4h30m")
	fmt.Printf("I've got %.1f hours of work left.", h.Hours())
	// Output: I've got 4.5 hours of work left.
}

func ExampleDuration_Microseconds() {
	u, _ := time.ParseDuration("1s")
	fmt.Printf("One second is %d microseconds.\n", u.Microseconds())
	// Output:
	// One second is 1000000 microseconds.
}

func ExampleDuration_Milliseconds() {
	u, _ := time.ParseDuration("1s")
	fmt.Printf("One second is %d milliseconds.\n", u.Milliseconds())
	// Output:
	// One second is 1000 milliseconds.
}

func ExampleDuration_Minutes() {
	m, _ := time.ParseDuration("1h30m")
	fmt.Printf("The movie is %.0f minutes long.", m.Minutes())
	// Output: The movie is 90 minutes long.
}

func ExampleDuration_Nanoseconds() {
	u, _ := time.ParseDuration("1µs")
	fmt.Printf("One microsecond is %d nanoseconds.\n", u.Nanoseconds())
	// Output:
	// One microsecond is 1000 nanoseconds.
}

func ExampleDuration_Seconds() {
	m, _ := time.ParseDuration("1m30s")
	fmt.Printf("Take off in t-%.0f seconds.", m.Seconds())
	// Output: Take off in t-90 seconds.
}

var c chan int

func handle(int) {}

func ExampleAfter() {
	select {
	case m := <-c:
		handle(m)
	case <-time.After(10 * time.Second):
		fmt.Println("timed out")
	}
}

func ExampleSleep() {
	time.Sleep(100 * time.Millisecond)
}

func statusUpdate() string { return "" }

func ExampleTick() {
	c := time.Tick(5 * time.Second)
	for next := range c {
		fmt.Printf("%v %s\n", next, statusUpdate())
	}
}

func ExampleMonth() {
	_, month, day := time.Now().Date()
	if month == time.November && day == 10 {
		fmt.Println("Happy Go day!")
	}
}

func ExampleDate() {
	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Printf("Go launched at %s\n", t.Local())
	// Output: Go launched at 2009-11-10 15:00:00 -0800 PST
}

func ExampleNewTicker() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	done := make(chan bool)
	go func() {
		time.Sleep(10 * time.Second)
		done <- true
	}()
	for {
		select {
		case <-done:
			fmt.Println("Done!")
			return
		case t := <-ticker.F.C:
			fmt.Println("Current time: ", t)
		}
	}
}

func ExampleTime_Format() {
	// 从标准Unix格式的字符串中解析时间值。. md5:f3f8abffc56da1da
	t, err := time.Parse(time.UnixDate, "Wed Feb 25 11:06:39 PST 2015")
	if err != nil { // 即使不应该发生错误，也始终检查错误。. md5:1ee31877d1f35f48
		panic(err)
	}

	tz, err := time.LoadLocation("Asia/Shanghai")
	if err != nil { // 即使不应该发生错误，也始终检查错误。. md5:1ee31877d1f35f48
		panic(err)
	}

	// time.Time 类型的 Stringer 接口在没有指定格式的情况下也很有用。. md5:d26240ee01fcdb99
	fmt.Println("default format:", t)

	// 包中预定义的常量实现了常见的布局。. md5:efc3cd559fe4fd78
	fmt.Println("Unix format:", t.Format(time.UnixDate))

	// 与时间值关联的时区会影响其输出。. md5:e4befa4c6f328243
	fmt.Println("Same, in UTC:", t.UTC().Format(time.UnixDate))

	fmt.Println("in Shanghai with seconds:", t.In(tz).Format("2006-01-02T15:04:05 -070000"))

	fmt.Println("in Shanghai with colon seconds:", t.In(tz).Format("2006-01-02T15:04:05 -07:00:00"))

	// 函数的剩余部分演示了在格式中使用的布局字符串的属性。
	// md5:18b912fdc17d3847

	// Parse函数和Format方法使用的布局字符串示例展示了参考时间应该如何表示。
	// 我们强调必须展示参考时间的格式，而不是用户选择的时间。因此，每个布局字符串都是时间戳的表示，
	// 例如：Jan 2 15:04:05 2006 MST
	// 记住这个值的一个简单方法是，当按照这个顺序呈现时，它包含了（与上面的元素对齐）：
	//	1 2  3  4  5    6  -7
	// 下面有一些说明示例。
	// md5:5c1ab7903b161365

	// 大多数情况下，Format和Parse的使用都采用诸如本包中定义的常量布局字符串，但该接口是灵活的，如下例所示。
	// md5:f3dbf95fe362fed3

	// 定义一个辅助函数，使示例的输出看起来更整洁。. md5:744e830c88834a10
	do := func(name, layout, want string) {
		got := t.Format(layout)
		if want != got {
			fmt.Printf("error: for %q got %q; expected %q\n", layout, got, want)
			return
		}
		fmt.Printf("%-16s %q gives %q\n", name, layout, got)
	}

	// 在我们的输出中打印一个标题。. md5:8db341353a187ad8
	fmt.Printf("\nFormats:\n\n")

	// 简单的启动示例。. md5:877bb2f43f13947f
	do("Basic full date", "Mon Jan 2 15:04:05 MST 2006", "Wed Feb 25 11:06:39 PST 2015")
	do("Basic short date", "2006/01/02", "2015/02/25")

	// 参考时间的小时为15，即下午3点。布局可以两种方式表示它，由于我们的值是在上午，所以我们应该看到它作为一个上午的时间。我们在一个格式字符串中展示两种表示。同时使用小写字母。
	// md5:4eee1220f173861d
	do("AM/PM", "3PM==3pm==15h", "11AM==11am==11h")

	// 在解析时，如果秒值后面跟着小数点和一些数字，
	// 即使布局字符串不表示分数秒，也会将其视为秒的分数。
	// 这里我们向上面使用的的时间值添加了一个分数秒。
	// md5:f479515c456c93fc
	t, err = time.Parse(time.UnixDate, "Wed Feb 25 11:06:39.1234 PST 2015")
	if err != nil {
		panic(err)
	}
	// 如果布局字符串不包含小数秒的表示，则不会出现在输出中。
	// md5:498a488987974671
	do("No fraction", time.UnixDate, "Wed Feb 25 11:06:39 PST 2015")

	// 在格式字符串的秒值中，通过在小数点后添加一串0或9，可以打印出小数部分秒。如果格式字符串中的数字是0，那么小数秒的宽度就是指定的。需要注意的是，输出将以0结尾。
	// md5:3063cbb9352ae856
	do("0s for fraction", "15:04:05.00000", "11:06:39.12340")

	// 如果布局中的分数部分为9s，将删除尾随的零。. md5:ac61ce54f30db97d
	do("9s for fraction", "15:04:05.99999999", "11:06:39.1234")

	// Output:
	// default format: 2015-02-25 11:06:39 -0800 PST
	// Unix format: Wed Feb 25 11:06:39 PST 2015
	// Same, in UTC: Wed Feb 25 19:06:39 UTC 2015
	//in Shanghai with seconds: 2015-02-26T03:06:39 +080000
	//in Shanghai with colon seconds: 2015-02-26T03:06:39 +08:00:00
	//
	// Formats:
	//
	// Basic full date  "Mon Jan 2 15:04:05 MST 2006" gives "Wed Feb 25 11:06:39 PST 2015"
	// Basic short date "2006/01/02" gives "2015/02/25"
	// AM/PM            "3PM==3pm==15h" gives "11AM==11am==11h"
	// No fraction      "Mon Jan _2 15:04:05 MST 2006" gives "Wed Feb 25 11:06:39 PST 2015"
	// 0s for fraction  "15:04:05.00000" gives "11:06:39.12340"
	// 9s for fraction  "15:04:05.99999999" gives "11:06:39.1234"

}

func ExampleTime_Format_pad() {
	// 从标准Unix格式的字符串中解析时间值。. md5:f3f8abffc56da1da
	t, err := time.Parse(time.UnixDate, "Sat Mar 7 11:06:39 PST 2015")
	if err != nil { // 即使不应该发生错误，也始终检查错误。. md5:1ee31877d1f35f48
		panic(err)
	}

	// 定义一个辅助函数，使示例的输出看起来更整洁。. md5:744e830c88834a10
	do := func(name, layout, want string) {
		got := t.Format(layout)
		if want != got {
			fmt.Printf("error: for %q got %q; expected %q\n", layout, got, want)
			return
		}
		fmt.Printf("%-16s %q gives %q\n", name, layout, got)
	}

	// 预定义的常量Unix使用下划线来填充日期。. md5:d9f4dbdd3de4ebe7
	do("Unix", time.UnixDate, "Sat Mar  7 11:06:39 PST 2015")

	// 对于可能为一或两个字符（例如日期，7 vs. 07）的固定宽度打印值，使用下划线_代替空格在格式字符串中。这里我们只打印日期，格式字符串中的数字是2，而实际值是7。
	// md5:a1ebfb451e3503bb
	do("No pad", "<2>", "<7>")

	// 下划线代表一个空格垫，如果日期只有一位数字。. md5:49def7f2b857b6f0
	do("Spaces", "<_2>", "< 7>")

	// "0" 表示为单数字值填充零。. md5:014b4fd453e8e5b1
	do("Zeros", "<02>", "<07>")

	// 如果值已经是正确的宽度，不需要填充。
	// 例如，我们值中的第二部分（参考时间中的05）为39，因此不需要填充，
	// 但分钟（04, 06）则需要。
	// md5:f641faf4ba09711a
	do("Suppressed pad", "04:05", "06:39")

	// Output:
	// Unix             "Mon Jan _2 15:04:05 MST 2006" gives "Sat Mar  7 11:06:39 PST 2015"
	// No pad           "<2>" gives "<7>"
	// Spaces           "<_2>" gives "< 7>"
	// Zeros            "<02>" gives "<07>"
	// Suppressed pad   "04:05" gives "06:39"

}

func ExampleTime_GoString() {
	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println(t.GoString())
	t = t.Add(1 * time.Minute)
	fmt.Println(t.GoString())
	t = t.AddDate(0, 1, 0)
	fmt.Println(t.GoString())
	t, _ = time.Parse("Jan 2, 2006 at 3:04pm (MST)", "Feb 3, 2013 at 7:54pm (UTC)")
	fmt.Println(t.GoString())

	// Output:
	// time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	// time.Date(2009, time.November, 10, 23, 1, 0, 0, time.UTC)
	// time.Date(2009, time.December, 10, 23, 1, 0, 0, time.UTC)
	// time.Date(2013, time.February, 3, 19, 54, 0, 0, time.UTC)
}

func ExampleParse() {
	// 参见Time.Format示例，详细了解如何定义解析时间.Time值的布局字符串；Parse和Format使用相同的模型来描述其输入和输出。
	// md5:9d063b6b107bc3c8

	// longForm通过示例展示参考时间在期望的格式中是如何表示的。
	// md5:761675d50d7047bd
	const longForm = "Jan 2, 2006 at 3:04pm (MST)"
	t, _ := time.Parse(longForm, "Feb 3, 2013 at 7:54pm (PST)")
	fmt.Println(t)

	// shortForm 是引用时间在期望布局下的另一种表示方式，
	// 其中不包含时区信息。
	// 注意：如果没有明确指定时区，默认返回的是UTC时间。
	// md5:82da07f19c343810
	const shortForm = "2006-Jan-02"
	t, _ = time.Parse(shortForm, "2013-Feb-03")
	fmt.Println(t)

	// 一些有效的布局可能对应无效的时间值，因为格式规范如_用于空格填充，Z用于时区信息。
	// 例如，RFC3339布局"2006-01-02T15:04:05Z07:00"包含了Z和时区偏移，以处理两种有效的情况：
	// 2006-01-02T15:04:05Z
	// 2006-01-02T15:04:05+07:00
	// md5:a0f0fb87b351aede
	t, _ = time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
	fmt.Println(t)
	t, _ = time.Parse(time.RFC3339, "2006-01-02T15:04:05+07:00")
	fmt.Println(t)
	_, err := time.Parse(time.RFC3339, time.RFC3339)
	fmt.Println("error", err) // 如果布局不是有效的时间值，返回一个错误. md5:487c96021bec619d

	// Output:
	// 2013-02-03 19:54:00 -0800 PST
	// 2013-02-03 00:00:00 +0000 UTC
	// 2006-01-02 15:04:05 +0000 UTC
	// 2006-01-02 15:04:05 +0700 +0700
	// error parsing time "2006-01-02T15:04:05Z07:00": extra text: "07:00"
}

func ExampleParseInLocation() {
	loc, _ := time.LoadLocation("Europe/Berlin")

	// 这将查找位于Europe/Berlin时区的"CEST"。. md5:fa2a6f1d39b85d79
	const longForm = "Jan 2, 2006 at 3:04pm (MST)"
	t, _ := time.ParseInLocation(longForm, "Jul 9, 2012 at 5:02am (CEST)", loc)
	fmt.Println(t)

	// 注意：如果没有明确指定时区，则按照给定的地点返回时间。. md5:b453850c1e5f79e8
	const shortForm = "2006-Jan-02"
	t, _ = time.ParseInLocation(shortForm, "2012-Jul-09", loc)
	fmt.Println(t)

	// Output:
	// 2012-07-09 05:02:00 +0200 CEST
	// 2012-07-09 00:00:00 +0200 CEST
}

func ExampleUnix() {
	unixTime := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println(unixTime.Unix())
	t := time.Unix(unixTime.Unix(), 0).UTC()
	fmt.Println(t)

	// Output:
	// 1257894000
	// 2009-11-10 23:00:00 +0000 UTC
}

func ExampleUnixMicro() {
	umt := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println(umt.UnixMicro())
	t := time.UnixMicro(umt.UnixMicro()).UTC()
	fmt.Println(t)

	// Output:
	// 1257894000000000
	// 2009-11-10 23:00:00 +0000 UTC
}

func ExampleUnixMilli() {
	umt := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println(umt.UnixMilli())
	t := time.UnixMilli(umt.UnixMilli()).UTC()
	fmt.Println(t)

	// Output:
	// 1257894000000
	// 2009-11-10 23:00:00 +0000 UTC
}

func ExampleTime_Unix() {
	// 以三种方式表示10亿秒的Unix时间。. md5:f6f77989ce38890c
	fmt.Println(time.Unix(1e9, 0).UTC())     // 1e9 seconds
	fmt.Println(time.Unix(0, 1e18).UTC())    // 1e18 nanoseconds
	fmt.Println(time.Unix(2e9, -1e18).UTC()) // 2e9 秒 - 1e18 纳秒. md5:bf13df9731ad7475

	t := time.Date(2001, time.September, 9, 1, 46, 40, 0, time.UTC)
	fmt.Println(t.Unix())     // seconds since 1970
	fmt.Println(t.UnixNano()) // nanoseconds since 1970

	// Output:
	// 2001-09-09 01:46:40 +0000 UTC
	// 2001-09-09 01:46:40 +0000 UTC
	// 2001-09-09 01:46:40 +0000 UTC
	// 1000000000
	// 1000000000000000000
}

func ExampleTime_Round() {
	t := time.Date(0, 0, 0, 12, 15, 30, 918273645, time.UTC)
	round := []time.Duration{
		time.Nanosecond,
		time.Microsecond,
		time.Millisecond,
		time.Second,
		2 * time.Second,
		time.Minute,
		10 * time.Minute,
		time.Hour,
	}

	for _, d := range round {
		fmt.Printf("t.Round(%6s) = %s\n", d, t.Round(d).Format("15:04:05.999999999"))
	}
	// Output:
	// t.Round(   1ns) = 12:15:30.918273645
	// t.Round(   1µs) = 12:15:30.918274
	// t.Round(   1ms) = 12:15:30.918
	// t.Round(    1s) = 12:15:31
	// t.Round(    2s) = 12:15:30
	// t.Round(  1m0s) = 12:16:00
	// t.Round( 10m0s) = 12:20:00
	// t.Round(1h0m0s) = 12:00:00
}

func ExampleTime_Truncate() {
	t, _ := time.Parse("2006 Jan 02 15:04:05", "2012 Dec 07 12:15:30.918273645")
	trunc := []time.Duration{
		time.Nanosecond,
		time.Microsecond,
		time.Millisecond,
		time.Second,
		2 * time.Second,
		time.Minute,
		10 * time.Minute,
	}

	for _, d := range trunc {
		fmt.Printf("t.Truncate(%5s) = %s\n", d, t.Truncate(d).Format("15:04:05.999999999"))
	}
	// 要四舍五入到本地时区的前一天午夜，创建一个新日期。. md5:f23f2d701f87ef7e
	midnight := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
	_ = midnight

	// Output:
	// t.Truncate(  1ns) = 12:15:30.918273645
	// t.Truncate(  1µs) = 12:15:30.918273
	// t.Truncate(  1ms) = 12:15:30.918
	// t.Truncate(   1s) = 12:15:30
	// t.Truncate(   2s) = 12:15:30
	// t.Truncate( 1m0s) = 12:15:00
	// t.Truncate(10m0s) = 12:10:00
}

func ExampleLoadLocation() {
	location, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		panic(err)
	}

	timeInUTC := time.Date(2018, 8, 30, 12, 0, 0, 0, time.UTC)
	fmt.Println(timeInUTC.In(location))
	// Output: 2018-08-30 05:00:00 -0700 PDT
}

func ExampleLocation() {
	// 中国没有夏令时。它使用相对于UTC固定的8小时偏移。. md5:6aab1ae63f661255
	secondsEastOfUTC := int((8 * time.Hour).Seconds())
	beijing := time.FixedZone("Beijing Time", secondsEastOfUTC)

	// 如果系统中存在时区数据库，那么可以从中加载一个时区，例如：
	//    newYork, err := time.LoadLocation("America/New_York")
	// md5:771c6af761901cc5

	// 创建一个时间点需要一个时区。常见的时区有time.Local和time.UTC。. md5:8603190e4b9df7b8
	timeInUTC := time.Date(2009, 1, 1, 12, 0, 0, 0, time.UTC)
	sameTimeInBeijing := time.Date(2009, 1, 1, 20, 0, 0, 0, beijing)

	// 虽然 UTC 时间是 1200，而北京时间是 2000，但由于北京比 UTC 领先 8 个小时，所以这两个日期实际上代表的是同一时刻。
	// md5:adb00e89cda10858
	timesAreEqual := timeInUTC.Equal(sameTimeInBeijing)
	fmt.Println(timesAreEqual)

	// Output:
	// true
}

func ExampleTime_Add() {
	start := time.Date(2009, 1, 1, 12, 0, 0, 0, time.UTC)
	afterTenSeconds := start.Add(time.Second * 10)
	afterTenMinutes := start.Add(time.Minute * 10)
	afterTenHours := start.Add(time.Hour * 10)
	afterTenDays := start.Add(time.Hour * 24 * 10)

	fmt.Printf("start = %v\n", start)
	fmt.Printf("start.Add(time.Second * 10) = %v\n", afterTenSeconds)
	fmt.Printf("start.Add(time.Minute * 10) = %v\n", afterTenMinutes)
	fmt.Printf("start.Add(time.Hour * 10) = %v\n", afterTenHours)
	fmt.Printf("start.Add(time.Hour * 24 * 10) = %v\n", afterTenDays)

	// Output:
	// start = 2009-01-01 12:00:00 +0000 UTC
	// start.Add(time.Second * 10) = 2009-01-01 12:00:10 +0000 UTC
	// start.Add(time.Minute * 10) = 2009-01-01 12:10:00 +0000 UTC
	// start.Add(time.Hour * 10) = 2009-01-01 22:00:00 +0000 UTC
	// start.Add(time.Hour * 24 * 10) = 2009-01-11 12:00:00 +0000 UTC
}

func ExampleTime_AddDate() {
	start := time.Date(2023, 03, 25, 12, 0, 0, 0, time.UTC)
	oneDayLater := start.AddDate(0, 0, 1)
	dayDuration := oneDayLater.Sub(start)
	oneMonthLater := start.AddDate(0, 1, 0)
	oneYearLater := start.AddDate(1, 0, 0)

	zurich, err := time.LoadLocation("Europe/Zurich")
	if err != nil {
		panic(err)
	}
	// This was the day before a daylight saving time transition in Zürich.
	startZurich := time.Date(2023, 03, 25, 12, 0, 0, 0, zurich)
	oneDayLaterZurich := startZurich.AddDate(0, 0, 1)
	dayDurationZurich := oneDayLaterZurich.Sub(startZurich)

	fmt.Printf("oneDayLater: start.AddDate(0, 0, 1) = %v\n", oneDayLater)
	fmt.Printf("oneMonthLater: start.AddDate(0, 1, 0) = %v\n", oneMonthLater)
	fmt.Printf("oneYearLater: start.AddDate(1, 0, 0) = %v\n", oneYearLater)
	fmt.Printf("oneDayLaterZurich: startZurich.AddDate(0, 0, 1) = %v\n", oneDayLaterZurich)
	fmt.Printf("Day duration in UTC: %v | Day duration in Zürich: %v\n", dayDuration, dayDurationZurich)

	// Output:
	// oneDayLater: start.AddDate(0, 0, 1) = 2023-03-26 12:00:00 +0000 UTC
	// oneMonthLater: start.AddDate(0, 1, 0) = 2023-04-25 12:00:00 +0000 UTC
	// oneYearLater: start.AddDate(1, 0, 0) = 2024-03-25 12:00:00 +0000 UTC
	// oneDayLaterZurich: startZurich.AddDate(0, 0, 1) = 2023-03-26 12:00:00 +0200 CEST
	// Day duration in UTC: 24h0m0s | Day duration in Zürich: 23h0m0s
}

func ExampleTime_After() {
	year2000 := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	year3000 := time.Date(3000, 1, 1, 0, 0, 0, 0, time.UTC)

	isYear3000AfterYear2000 := year3000.After(year2000) // True
	isYear2000AfterYear3000 := year2000.After(year3000) // False

	fmt.Printf("year3000.After(year2000) = %v\n", isYear3000AfterYear2000)
	fmt.Printf("year2000.After(year3000) = %v\n", isYear2000AfterYear3000)

	// Output:
	// year3000.After(year2000) = true
	// year2000.After(year3000) = false
}

func ExampleTime_Before() {
	year2000 := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	year3000 := time.Date(3000, 1, 1, 0, 0, 0, 0, time.UTC)

	isYear2000BeforeYear3000 := year2000.Before(year3000) // True
	isYear3000BeforeYear2000 := year3000.Before(year2000) // False

	fmt.Printf("year2000.Before(year3000) = %v\n", isYear2000BeforeYear3000)
	fmt.Printf("year3000.Before(year2000) = %v\n", isYear3000BeforeYear2000)

	// Output:
	// year2000.Before(year3000) = true
	// year3000.Before(year2000) = false
}

func ExampleTime_Date() {
	d := time.Date(2000, 2, 1, 12, 30, 0, 0, time.UTC)
	year, month, day := d.Date()

	fmt.Printf("year = %v\n", year)
	fmt.Printf("month = %v\n", month)
	fmt.Printf("day = %v\n", day)

	// Output:
	// year = 2000
	// month = February
	// day = 1
}

func ExampleTime_Day() {
	d := time.Date(2000, 2, 1, 12, 30, 0, 0, time.UTC)
	day := d.Day()

	fmt.Printf("day = %v\n", day)

	// Output:
	// day = 1
}

func ExampleTime_Equal() {
	secondsEastOfUTC := int((8 * time.Hour).Seconds())
	beijing := time.FixedZone("Beijing Time", secondsEastOfUTC)

	// 与相等运算符不同，Equal 方法能够识别 d1 和 d2 是同一时间点，只是处于不同的时区。
	// md5:9b87f1e5ff088a68
	d1 := time.Date(2000, 2, 1, 12, 30, 0, 0, time.UTC)
	d2 := time.Date(2000, 2, 1, 20, 30, 0, 0, beijing)

	datesEqualUsingEqualOperator := d1 == d2
	datesEqualUsingFunction := d1.Equal(d2)

	fmt.Printf("datesEqualUsingEqualOperator = %v\n", datesEqualUsingEqualOperator)
	fmt.Printf("datesEqualUsingFunction = %v\n", datesEqualUsingFunction)

	// Output:
	// datesEqualUsingEqualOperator = false
	// datesEqualUsingFunction = true
}

func ExampleTime_String() {
	timeWithNanoseconds := time.Date(2000, 2, 1, 12, 13, 14, 15, time.UTC)
	withNanoseconds := timeWithNanoseconds.String()

	timeWithoutNanoseconds := time.Date(2000, 2, 1, 12, 13, 14, 0, time.UTC)
	withoutNanoseconds := timeWithoutNanoseconds.String()

	fmt.Printf("withNanoseconds = %v\n", string(withNanoseconds))
	fmt.Printf("withoutNanoseconds = %v\n", string(withoutNanoseconds))

	// Output:
	// withNanoseconds = 2000-02-01 12:13:14.000000015 +0000 UTC
	// withoutNanoseconds = 2000-02-01 12:13:14 +0000 UTC
}

func ExampleTime_Sub() {
	start := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2000, 1, 1, 12, 0, 0, 0, time.UTC)

	difference := end.Sub(start)
	fmt.Printf("difference = %v\n", difference)

	// Output:
	// difference = 12h0m0s
}

func ExampleTime_AppendFormat() {
	t := time.Date(2017, time.November, 4, 11, 0, 0, 0, time.UTC)
	text := []byte("Time: ")

	text = t.AppendFormat(text, time.Kitchen)
	fmt.Println(string(text))

	// Output:
	// Time: 11:00AM
}

func ExampleFixedZone() {
	loc := time.FixedZone("UTC-8", -8*60*60)
	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, loc)
	fmt.Println("The time is:", t.Format(time.RFC822))
	// Output: The time is: 10 Nov 09 23:00 UTC-8
}
