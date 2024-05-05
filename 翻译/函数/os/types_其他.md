# **_方法.md 文件备注:
# ff= 方法,重命名方法名称
# 
# **_package.md 文件备注:
# bm= 包名,更换新的包名称, 如: package gin //bm:gin类
#
# **_其他.md 文件备注:
# qm= 行首,跳转到行首进行重命名.文档内如果有多个相同的,会一起重命名.
# th= 替换,用于替换文本,文档内如果有多个相同的,会一起替换
# cf= 重复,用于重命名多次,如: 一个文档内有2个"One(result interface{}) error"需要重命名.
#     但是要注意,多个新名称要保持一致. 如:"X取一条(result interface{})"
# zz= 正则表达式,用于结构名称替换或者复杂替换
#     如待替换: type authPair struct { //zz:^type *authPair

[ModeDir        = fs.ModeDir]
qm=常量_模式_目录

[ModeAppend     = fs.ModeAppend]
qm=常量_模式_追加

[ModeExclusive  = fs.ModeExclusive]
qm=常量_模式_独占使用

[ModeTemporary  = fs.ModeTemporary]
qm=常量_模式_临时文件

[ModeSymlink    = fs.ModeSymlink]
qm=常量_模式_符号链接

[ModeDevice     = fs.ModeDevice]
qm=常量_模式_设备

[ModeNamedPipe  = fs.ModeNamedPipe]
qm=常量_模式_命名管道

[ModeSocket     = fs.ModeSocket]
qm=常量_模式_套接字

[ModeSetuid     = fs.ModeSetuid]
qm=常量_模式_设置用户ID

[ModeSetgid     = fs.ModeSetgid]
qm=常量_模式_设置组ID

[ModeCharDevice = fs.ModeCharDevice]
qm=常量_模式_字符设备

[ModeSticky     = fs.ModeSticky]
qm=常量_模式_粘滞位

[ModeIrregular  = fs.ModeIrregular]
qm=常量_模式_非常规文件

[ModeType = fs.ModeType]
qm=常量_模式_类型

[ModePerm = fs.ModePerm]
qm=常量_模式_权限
