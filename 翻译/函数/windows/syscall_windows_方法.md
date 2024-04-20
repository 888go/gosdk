# **_方法.md 文件备注:
# ff= 方法,重命名方法名称
# 
# **_package.md 文件备注:
# bm= 包名,更换新的包名称, 如: package gin //bm:gin类
#
# **_其他.md 文件备注:
# hs= 行首,跳转到行首进行重命名.文档内如果有多个相同的,会一起重命名.
# th= 替换,用于替换文本,文档内如果有多个相同的,会一起替换
# cf= 重复,用于重命名多次,如: 一个文档内有2个"One(result interface{}) error"需要重命名.
#     但是要注意,多个新名称要保持一致. 如:"X取一条(result interface{})"
# zz= 正则表达式,用于结构名称替换或者复杂替换
#     如待替换: type authPair struct { //zz:^type *authPair

[func StringToUTF16(s string) #左中括号##右中括号#uint16 {]
ff=StringToUTF16弃用

[func UTF16FromString(s string) (#左中括号##右中括号#uint16, error) {]
ff=文本到UTF16
s=文本

[func UTF16ToString(s #左中括号##右中括号#uint16) string {]
ff=UTF16到文本

[func StringToUTF16Ptr(s string) *uint16 { return]
ff=StringToUTF16Ptr弃用

[func UTF16PtrFromString(s string) (*uint16, error) {]
ff=文本到UTF16指针
s=文本

[func UTF16PtrToString(p *uint16) string {]
ff=UTF16指针到文本
p=UTF16指针

[func GetCurrentProcess() (Handle, error) {]
ff=GetCurrentProcess弃用

[func CurrentProcess() Handle { return Handle(]
ff=取当前进程句柄

[func GetCurrentThread() (Handle, error) {]
ff=GetCurrentThread弃用

[func CurrentThread() Handle { return Handle(]
ff=取当前线程句柄
