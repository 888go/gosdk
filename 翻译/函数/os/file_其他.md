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

[Stdin  = NewFile(uintptr(syscall.Stdin), "/dev/stdin")]
qm=常量_标准输入
cz=Stdin  #等号# NewFile(uintptr(syscall.Stdin), "/dev/stdin")

[Stdout = NewFile(uintptr(syscall.Stdout), "/dev/stdout")]
qm=常量_标准输出
cz=Stdout #等号# NewFile(uintptr(syscall.Stdout), "/dev/stdout")

[Stderr = NewFile(uintptr(syscall.Stderr), "/dev/stderr")]
qm=常量_标准错误
cz=Stderr #等号# NewFile(uintptr(syscall.Stderr), "/dev/stderr")

[O_RDONLY int = syscall.O_RDONLY]
qm=常量_文件标志_只读打开
cz=O_RDONLY int #等号# syscall.O_RDONLY

[O_WRONLY int = syscall.O_WRONLY]
qm=常量_文件标志_只写模式
cz=O_WRONLY int #等号# syscall.O_WRONLY

[O_RDWR   int = syscall.O_RDWR]
qm=常量_文件标志_读写模式
cz=O_RDWR   int #等号# syscall.O_RDWR

[O_APPEND int = syscall.O_APPEND]
qm=常量_文件标志_追加模式
cz=O_APPEND int #等号# syscall.O_APPEND

[O_CREATE int = syscall.O_CREAT]
qm=常量_文件标志_不存在则创建
cz=O_CREATE int #等号# syscall.O_CREAT

[O_EXCL   int = syscall.O_EXCL]
qm=常量_文件标志_独占创建
cz=O_EXCL   int #等号# syscall.O_EXCL

[O_SYNC   int = syscall.O_SYNC]
qm=常量_文件标志_同步写入
cz=O_SYNC   int #等号# syscall.O_SYNC

[O_TRUNC  int = syscall.O_TRUNC]
qm=常量_文件标志_清空文件
cz=O_TRUNC  int #等号# syscall.O_TRUNC

[SEEK_SET int = 0]
qm=常量_位置_从头开始
cz=SEEK_SET int #等号# 0

[SEEK_CUR int = 1]
qm=常量_位置_从当前偏移量
cz=SEEK_CUR int #等号# 1

[SEEK_END int = 2]
qm=常量_位置_从末尾处
cz=SEEK_END int #等号# 2
