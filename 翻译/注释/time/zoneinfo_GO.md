
<原文开始>
// A Location maps time instants to the zone in use at that time.
// Typically, the Location represents the collection of time offsets
// in use in a geographical area. For many Locations the time offset varies
// depending on whether daylight savings time is in use at the time instant.
//
// Location is used to provide a time zone in a printed Time value and for
// calculations involving intervals that may cross daylight savings time
// boundaries.
<原文结束>

# <翻译开始>
// Location 将时间点映射到该时间点使用的时区。
// 通常，Location 表示地理区域内使用的所有时间偏移量。对于许多 Location，时间偏移量取决于所处的时间点是否使用夏令时。
//
// Location 用于在打印 Time 值时提供时区，并进行可能跨越夏令时边界的计算。
// md5:db05da442bb7e413
# <翻译结束>


<原文开始>
// UTC represents Universal Coordinated Time (UTC).
<原文结束>

# <翻译开始>
// UTC 表示协调世界时间（Universal Coordinated Time）。. md5:991e0d09aaa17cec
# <翻译结束>


<原文开始>
// Local represents the system's local time zone.
// On Unix systems, Local consults the TZ environment
// variable to find the time zone to use. No TZ means
// use the system default /etc/localtime.
// TZ="" means use UTC.
// TZ="foo" means use file foo in the system timezone directory.
<原文结束>

# <翻译开始>
// Local 表示系统的本地时区。
// 在Unix系统上，Local会查询TZ环境变量来找到要使用的时区。
// 没有TZ意味着使用系统默认的/etc/localtime。
// TZ="" 意味着使用UTC时间。
// TZ="foo" 意味着在系统时区目录中使用名为foo的文件。
// md5:38f2b49de5174f53
# <翻译结束>


<原文开始>
// String returns a descriptive name for the time zone information,
// corresponding to the name argument to LoadLocation or FixedZone.
<原文结束>

# <翻译开始>
// String 返回与 LoadLocation 或 FixedZone 的 name 参数对应的时区信息的描述性名称。
// md5:120873679bee5884
# <翻译结束>


<原文开始>
// FixedZone returns a Location that always uses
// the given zone name and offset (seconds east of UTC).
<原文结束>

# <翻译开始>
// FixedZone返回一个Location，它始终使用给定的时区名称和偏移量（UTC的东秒数）。
// md5:2175485bbe6a45d6
# <翻译结束>


<原文开始>
// LoadLocation returns the Location with the given name.
//
// If the name is "" or "UTC", LoadLocation returns UTC.
// If the name is "Local", LoadLocation returns Local.
//
// Otherwise, the name is taken to be a location name corresponding to a file
// in the IANA Time Zone database, such as "America/New_York".
//
// LoadLocation looks for the IANA Time Zone database in the following
// locations in order:
//
//   - the directory or uncompressed zip file named by the ZONEINFO environment variable
//   - on a Unix system, the system standard installation location
//   - $GOROOT/lib/time/zoneinfo.zip
//   - the time/tzdata package, if it was imported
<原文结束>

# <翻译开始>
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
// md5:aec8c6a750b5813b
# <翻译结束>

