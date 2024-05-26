package time

import "time"

// Location 将时间点映射到该时间点使用的时区。
// 通常，Location 表示地理区域内使用的所有时间偏移量。对于许多 Location，时间偏移量取决于所处的时间点是否使用夏令时。
//
// Location 用于在打印 Time 值时提供时区，并进行可能跨越夏令时边界的计算。
// md5:db05da442bb7e413
type Location struct {
	F time.Location
} //md5:b2741259c471a509

//const (
//	alpha = -1 << 63  // math.MinInt64
//	omega = 1<<63 - 1 // math.MaxInt64
//) // md5:7214a6052ed65a5f

//const (
//	ruleJulian ruleKind = iota
//	ruleDOY
//	ruleMonthWeekDay
//) // md5:cd28ad5922a22ce1

// UTC 表示协调世界时间（Universal Coordinated Time）。. md5:991e0d09aaa17cec
var UTC *Location = &Location{*time.UTC} //md5:cb0833e2130b743d

// Local 表示系统的本地时区。
// 在Unix系统上，Local会查询TZ环境变量来找到要使用的时区。
// 没有TZ意味着使用系统默认的/etc/localtime。
// TZ="" 意味着使用UTC时间。
// TZ="foo" 意味着在系统时区目录中使用名为foo的文件。
// md5:38f2b49de5174f53
var Local *Location = &Location{*time.Local} //md5:57aa689da8ee76b4

// String 返回与 LoadLocation 或 FixedZone 的 name 参数对应的时区信息的描述性名称。
// md5:120873679bee5884
// 翻译提示:func  (l  *位置)  字符串()  字符串  {}

// ff:
func (l *Location) String() string { //md5:c44612032555be76
	return l.F.String()
}

// FixedZone返回一个Location，它始终使用给定的时区名称和偏移量（UTC的东秒数）。
// md5:2175485bbe6a45d6
// 翻译提示:func  固定时区(name  字符串,  偏移量  int)  *时区信息  {}

// ff:
// name:
// offset:
func FixedZone(name string, offset int) *Location { //md5:839d95c9de0383bc
	l := time.FixedZone(name, offset)
	if l == nil {
		return nil
	}
	return &Location{F: *l}
}

// LoadLocation 返回给定名称的 Location。
//
// 如果名称是 "" 或 "UTC"，LoadLocation 返回 UTC。
// 如果名称是 "Local"，LoadLocation 返回 Local。
//
// 否则，该名称被视为与 IANA 时区数据库中文件相对应的时区名称，例如 "America/New_York"。
//
// LoadLocation 按照以下顺序查找 IANA 时区数据库：
//
//   - 由 ZONEINFO 环境变量命名的目录或未压缩的zip文件
//   - 在 Unix 系统上，系统标准安装位置
//   - $GOROOT/lib/time/zoneinfo.zip
//   - 如果导入了 time/tzdata 包，则使用该包
//
// md5:aec8c6a750b5813b
// 翻译提示:func  加载时区(name  string)  (*时区信息,  error)  {}

// ff:
// name:
func LoadLocation(name string) (*Location, error) { //md5:506ad5b64122238f
	l, err := time.LoadLocation(name)
	if err != nil {
		return nil, err
	}
	return &Location{F: *l}, err
}
