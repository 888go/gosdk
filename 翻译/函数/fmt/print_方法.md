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
#    re.F.String()
# }
# //zj:
# 备注结束

[func Fprintf(w io.Writer, format string, a ...any) (n int, err error) {]
ff=格式化到IO写入器
err=错误
n=长度
a=参数
format=格式化模板
w=IO写入器

[func Printf(format string, a ...any) (n int, err error) {]
ff=打印格式化文本
err=错误
n=长度
a=参数
format=格式化模板

[func Sprintf(format string, a ...any) string {]
ff=格式化文本
a=参数
format=格式化模板

[func Appendf(b #左中括号##右中括号#byte, format string, a ...any) #左中括号##右中括号#byte {]
ff=格式化并追加切片
a=参数
format=格式化模板
b=切片

[func Fprint(w io.Writer, a ...any) (n int, err error) {]
ff=打印并写IO写入器
err=错误
n=长度
a=参数
w=IO写入器

[func Print(a ...any) (n int, err error) {]
ff=打印
err=错误
n=长度
a=参数

[func Sprint(a ...any) string {]
ff=拼接文本
a=参数

[func Append(b #左中括号##右中括号#byte, a ...any) #左中括号##右中括号#byte {]
ff=追加切片
a=参数
b=切片

[func Fprintln(w io.Writer, a ...any) (n int, err error) {]
ff=打印ln并写IO写入器
err=错误
n=长度
a=参数
w=IO写入器

[func Println(a ...any) (n int, err error) {]
ff=打印ln
err=错误
n=长度
a=参数

[func Sprintln(a ...any) string {]
ff=拼接文本ln
a=参数

[func Appendln(b #左中括号##右中括号#byte, a ...any) #左中括号##右中括号#byte {]
ff=追加切片ln
a=参数
b=切片
