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

[func Scan(a ...any) (n int, err error) {]
ff=标准输入到指针
err=错误
n=长度
a=变量指针

[func Scanln(a ...any) (n int, err error) {]
ff=标准输入到指针ln
err=错误
n=长度
a=变量指针

[func Scanf(format string, a ...any) (n int, err error) {]
ff=格式化标准输入到指针
err=错误
n=长度
a=变量指针
format=格式化模板

[func Sscan(str string, a ...any) (n int, err error) {]
ff=文本扫描到指针
err=错误
n=长度
a=变量指针
str=文本

[func Sscanln(str string, a ...any) (n int, err error) {]
ff=文本扫描到指针ln
err=错误
n=长度
a=变量指针
str=文本

[func Sscanf(str string, format string, a ...any) (n int, err error) {]
ff=格式化文本到指针
err=错误
n=长度
a=变量指针
format=格式化模板
str=文本

[func Fscan(r io.Reader, a ...any) (n int, err error) {]
ff=IO读取器到指针
err=错误
n=长度
a=变量指针
r=IO读取器

[func Fscanln(r io.Reader, a ...any) (n int, err error) {]
ff=IO读取器到指针ln
err=错误
n=长度
a=变量指针
r=IO读取器

[func Fscanf(r io.Reader, format string, a ...any) (n int, err error) {]
ff=格式化IO读取器到指针
err=错误
n=长度
a=变量指针
format=格式化模板
r=IO读取器
