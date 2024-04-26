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

[func (f *File) Readdir(n int) (#左中括号##右中括号#FileInfo, error) {]
ff=读取目录2
n=返回数量

[func (f *File) Readdirnames(n int) (names #左中括号##右中括号#string, err error) {]
names=文件名s
n=返回数量

[func (f *File) ReadDir(n int) (#左中括号##右中括号#DirEntry, error) {]
ff=读取目录
n=返回数量

[func ReadDir(name string) (#左中括号##右中括号#DirEntry, error) {]
ff=读取目录
name=目录路径
