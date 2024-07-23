# 备注开始
# **_方法.md 文件备注:
# ff= 方法,重命名方法名称
# 如://ff:取文本
#
# yx=true,此方法优先翻译
# 如: //yx=true

# **_package.md 文件备注:
# bm= 包名,更换新的包名称 
# 如: package gin //bm:gin类

# **_其他.md 文件备注:
# qm= 前面,跳转到前面进行重命名.文档内如果有多个相同的,会一起重命名.
# hm= 后面,跳转到后面进行重命名.文档内如果有多个相同的,会一起重命名.
# cz= 查找,配合前面/后面使用,
# zz= 正则查找,配合前面/后面使用, 有设置正则查找,就不用设置上面的查找
# 如: type Regexp struct {//qm:正则 cz:Regexp struct
#
# th= 替换,用于替换文本,文档内如果有多个相同的,会一起替换
# 如:
# type Regexp struct {//th:type Regexp222 struct
#
# cf= 重复,用于重命名多次,
# 如: 
# 一个文档内有2个"One(result interface{}) error"需要重命名.
# 但是要注意,多个新名称要保持一致. 如:"X取一条(result interface{})"

# **_追加.md 文件备注:
# 在代码内追加代码,如:
# //zj:前面一行的代码,如果为空,追加到末尾行
# func (re *Regexp) X取文本() string { 
# re.F.String()
# }
# //zj:
# 备注结束

[func (t Time) After(u Time) bool {]
ff=是否为之后
u=时间

[func (t Time) Before(u Time) bool {]
ff=是否为之前
u=时间

[func (t Time) Compare(u Time) int {]
ff=比较
u=时间

[func (t Time) Equal(u Time) bool {]
ff=是否相等
u=时间

[func (m Month) String() string {]
ff=取英文月份

[func (d Weekday) String() string {]
ff=取英文星期几

[func (t Time) IsZero() bool {]
ff=是否为零

[func (t Time) Date() (year int, month Month, day int) {]
ff=取日期
day=日
month=月份
year=年份

[func (t Time) Year() int {]
ff=取年份

[func (t Time) Month() Month {]
ff=取月份

[func (t Time) Day() int {]
ff=取日

[func (t Time) Weekday() Weekday {]
ff=取星期几

[func (t Time) ISOWeek() (year, week int) {]
ff=取ISO年周
week=周数
year=年

[func (t Time) Clock() (hour, min, sec int) {]
ff=取时分秒
sec=秒
min=分钟
hour=小时

[func (t Time) Hour() int {]
ff=取小时

[func (t Time) Minute() int {]
ff=取分钟

[func (t Time) Second() int {]
ff=取秒

[func (t Time) Nanosecond() int {]
ff=取纳秒

[func (t Time) YearDay() int {]
ff=取第几天

[func (d Duration) Nanoseconds() int64 {]
ff=取纳秒数

[func (d Duration) Microseconds() int64 {]
ff=取微秒数

[func (d Duration) Milliseconds() int64 {]
ff=取毫秒数

[func (d Duration) Seconds() float64 {]
ff=取秒数

[func (d Duration) Minutes() float64 {]
ff=取分钟数

[func (d Duration) Hours() float64 {]
ff=取小时数

[func (d Duration) Truncate(m Duration) Duration {]
ff=截断并按下取整
m=间隔时长

[func (d Duration) Round(m Duration) Duration {]
ff=截断并按四舍五入

[func (d Duration) Abs() Duration {]
ff=取绝对值

[func (t Time) Add(d Duration) Time {]
ff=增加时长
d=时长

[func (t Time) Sub(u Time) Duration {]
ff=减去
u=时间

[func Since(t Time) Duration {]
t=时间

[func Now() Time {]
ff=取当前时间

[func (t Time) UTC() Time {]
ff=取UTC时间

[func (t Time) Local() Time {]
ff=取系统时间

[func (t Time) Location() *Location {]
ff=取时区

[func (t Time) Zone() (name string, offset int) {]
ff=取时区名称
offset=秒偏移量
name=时区名称

[func (t Time) ZoneBounds() (start, end Time) {]
ff=取时区边界
end=结束时间
start=开始时间

[func (t Time) Unix() int64 {]
ff=取Unix秒时间戳

[func (t Time) UnixMilli() int64 {]
ff=取Unix毫秒时间戳

[func (t Time) UnixMicro() int64 {]
ff=取Unix微秒时间戳

[func (t Time) UnixNano() int64 {]
ff=取Unix纳秒时间戳

[func Unix(sec int64, nsec int64) Time {]
ff=创建并按时间戳
nsec=可选毫秒
sec=可选秒

[func UnixMilli(msec int64) Time {]
ff=创建并按毫秒时间戳
msec=毫秒时间戳

[func UnixMicro(usec int64) Time {]
ff=创建并按微秒时间戳
usec=微秒时间戳

[func (t Time) IsDST() bool {]
ff=是否为夏令时

[func Date(year int, month Month, day, hour, min, sec, nsec int, loc *Location) Time {]
ff=创建时间
loc=时区
nsec=毫秒
sec=秒
min=分
hour=时
day=日
month=月
year=年

[func (t Time) Truncate(d Duration) Time {]
ff=截断并按下取整
d=时长

[func (t Time) Round(d Duration) Time {]
ff=截断并按四舍五入
d=时长
