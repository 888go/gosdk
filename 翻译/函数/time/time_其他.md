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
# //zj:
# func (re *Regexp) X取文本() string { 
# re.F.String()
# }
# //zj:
# 备注结束

[type Month int]
hm=月份
cz=type Month

[type Duration int64]
hm=时长
cz=type Duration

[type Time struct {]
hm=时间结构
cz=type Time

[January Month = 1]
qm=常量_月份_一月
zz=January.*#等号# 1

[February Month = 2]
qm=常量_月份_二月
zz=February.*#等号# 2

[March Month = 3]
qm=常量_月份_三月
zz=March.*#等号# 3

[April Month = 4]
qm=常量_月份_四月
zz=April.*#等号# 4

[May Month = 5]
qm=常量_月份_五月
zz=May.*#等号# 5

[June Month = 6]
qm=常量_月份_六月
zz=June.*#等号# 6

[July Month = 7]
qm=常量_月份_七月
zz=July.*#等号# 7

[August Month = 8]
qm=常量_月份_八月
zz=August.*#等号# 8

[September Month = 9]
qm=常量_月份_九月
zz=September.*#等号# 9

[October Month = 10]
qm=常量_月份_十月
zz=October.*#等号# 10

[November Month = 11]
qm=常量_月份_十一月
zz=November.*#等号# 11

[December Month = 12]
qm=常量_月份_十二月
zz=December.*#等号# 12

[Sunday Weekday = 0]
qm=常量_星期_周日
cz=Sunday Weekday #等号# 0

[Monday Weekday = 1]
qm=常量_星期_周一
cz=Monday Weekday #等号# 1

[Tuesday Weekday = 2]
qm=常量_星期_周二
cz=Tuesday Weekday #等号# 2

[Wednesday Weekday = 3]
qm=常量_星期_周三
cz=Wednesday Weekday #等号# 3

[Thursday Weekday = 4]
qm=常量_星期_周四
cz=Thursday Weekday #等号# 4

[Friday Weekday = 5]
qm=常量_星期_周五
cz=Friday Weekday #等号# 5

[Saturday Weekday = 6]
qm=常量_星期_周六
cz=Saturday Weekday #等号# 6

[Nanosecond Duration = 1]
qm=常量_时长_纳秒
zz=Nanosecond.*#等号# 1

[Microsecond = 1000 * Nanosecond]
qm=常量_时长_微秒
cz=Microsecond #等号# 1000 *

[Millisecond = 1000 * Microsecond]
qm=常量_时长_毫秒
cz=Millisecond #等号# 1000 *

[Second = 1000 * Millisecond]
qm=常量_时长_秒
cz=Second #等号# 1000 *

[Minute = 60 * Second]
qm=常量_时长_分钟
cz=Minute #等号# 60 *

[Hour = 60 * Minute]
qm=常量_时长_小时
cz=Hour #等号# 60 *
