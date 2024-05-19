# 备注开始
# **_方法.md 文件备注:
# ff= 方法,重命名方法名称
# 如:
# //ff:取文本

# **_package.md 文件备注:
# bm= 包名,更换新的包名称 
# 如: 
# package gin //bm:gin类

# **_其他.md 文件备注:
# qm= 前面,跳转到前面进行重命名.文档内如果有多个相同的,会一起重命名.
# hm= 后面,跳转到后面进行重命名.文档内如果有多个相同的,会一起重命名.
# cz= 查找,配合前面/后面使用,
# 如:
# type Regexp struct {//qm:正则 cz:Regexp struct
#
# th= 替换,用于替换文本,文档内如果有多个相同的,会一起替换
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





[ModeDir        = fs.ModeDir]
qm=常量_模式_目录
cz=ModeDir        #等号# fs.ModeDir

[ModeAppend     = fs.ModeAppend]
qm=常量_模式_追加
cz=ModeAppend     #等号# fs.ModeAppend

[ModeExclusive  = fs.ModeExclusive]
qm=常量_模式_独占使用
cz=ModeExclusive  #等号# fs.ModeExclusive

[ModeTemporary  = fs.ModeTemporary]
qm=常量_模式_临时文件
cz=ModeTemporary  #等号# fs.ModeTemporary

[ModeSymlink    = fs.ModeSymlink]
qm=常量_模式_符号链接
cz=ModeSymlink    #等号# fs.ModeSymlink

[ModeDevice     = fs.ModeDevice]
qm=常量_模式_设备
cz=ModeDevice     #等号# fs.ModeDevice

[ModeNamedPipe  = fs.ModeNamedPipe]
qm=常量_模式_命名管道
cz=ModeNamedPipe  #等号# fs.ModeNamedPipe

[ModeSocket     = fs.ModeSocket]
qm=常量_模式_套接字
cz=ModeSocket     #等号# fs.ModeSocket

[ModeSetuid     = fs.ModeSetuid]
qm=常量_模式_设置用户ID
cz=ModeSetuid     #等号# fs.ModeSetuid

[ModeSetgid     = fs.ModeSetgid]
qm=常量_模式_设置组ID
cz=ModeSetgid     #等号# fs.ModeSetgid

[ModeCharDevice = fs.ModeCharDevice]
qm=常量_模式_字符设备
cz=ModeCharDevice #等号# fs.ModeCharDevice

[ModeSticky     = fs.ModeSticky]
qm=常量_模式_粘滞位
cz=ModeSticky     #等号# fs.ModeSticky

[ModeIrregular  = fs.ModeIrregular]
qm=常量_模式_非常规文件
cz=ModeIrregular  #等号# fs.ModeIrregular

[ModeType = fs.ModeType]
qm=常量_模式_类型
cz=ModeType #等号# fs.ModeType

[ModePerm = fs.ModePerm]
qm=常量_模式_权限
cz=ModePerm #等号# fs.ModePerm
