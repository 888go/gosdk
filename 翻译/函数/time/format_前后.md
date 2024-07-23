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

[Layout = "01/02 03:04:05PM '06 -0700"]
qm=常量_时间格式_Layout
cz=Layout #等号# "01/02 03:04:05PM '06 -0700"

[ANSIC = "Mon Jan _2 15:04:05 2006"]
qm=常量_时间格式_ANSIC
cz=ANSIC #等号# "Mon Jan _2 15:04:05 2006"

[UnixDate = "Mon Jan _2 15:04:05 MST 2006"]
qm=常量_时间格式_Unix
cz=UnixDate #等号# "Mon Jan _2 15:04:05 MST 2006"

[RubyDate = "Mon Jan 02 15:04:05 -0700 2006"]
qm=常量_时间格式_Ruby
cz=RubyDate #等号# "Mon Jan 02 15:04:05 -0700 2006"

[RFC822 = "02 Jan 06 15:04 MST"]
qm=常量_时间格式_RFC822
cz=RFC822 #等号# "02 Jan 06 15:04 MST"

[RFC822Z = "02 Jan 06 15:04 -0700"]
qm=常量_时间格式_RFC822Z
cz=RFC822Z #等号# "02 Jan 06 15:04 -0700"

[RFC850 = "Monday, 02-Jan-06 15:04:05 MST"]
qm=常量_时间格式_RFC850
cz=RFC850 #等号# "Monday, 02-Jan-06 15:04:05 MST"

[RFC1123 = "Mon, 02 Jan 2006 15:04:05 MST"]
qm=常量_时间格式_RFC1123
cz=RFC1123 #等号# "Mon, 02 Jan 2006 15:04:05 MST"

[RFC1123Z = "Mon, 02 Jan 2006 15:04:05 -0700"]
qm=常量_时间格式_RFC1123Z
cz=RFC1123Z #等号# "Mon, 02 Jan 2006 15:04:05 -0700"

[RFC3339 = "2006-01-02T15:04:05Z07:00"]
qm=常量_时间格式_RFC3339
cz=RFC3339 #等号# "2006-01-02T15:04:05Z07:00"

[RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"]
qm=常量_时间格式_RFC3339Nano
cz=RFC3339Nano #等号# "2006-01-02T15:04:05.999999999Z07:00"

[Kitchen = "3:04PM"]
qm=常量_时间格式_厨房时间
cz=Kitchen #等号# "3:04PM"

[Stamp = "Jan _2 15:04:05"]
qm=常量_时间格式_时间戳
cz=Stamp #等号# "Jan _2 15:04:05"

[StampMilli = "Jan _2 15:04:05.000"]
qm=常量_时间格式_时间戳毫秒
cz=StampMilli #等号# "Jan _2 15:04:05.000"

[StampMicro = "Jan _2 15:04:05.000000"]
qm=常量_时间格式_时间戳微秒
cz=StampMicro #等号# "Jan _2 15:04:05.000000"

[StampNano = "Jan _2 15:04:05.000000000"]
qm=常量_时间格式_时间戳纳秒
cz=StampNano #等号# "Jan _2 15:04:05.000000000"

[DateTime = "2006-01-02 15:04:05"]
qm=常量_时间格式_日期时间
cz=DateTime #等号# "2006-01-02 15:04:05"

[DateOnly = "2006-01-02"]
qm=常量_时间格式_日期
cz=DateOnly #等号# "2006-01-02"

[TimeOnly = "15:04:05"]
qm=常量_时间格式_时间
cz=TimeOnly #等号# "15:04:05"
