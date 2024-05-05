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

[Stdin  = NewFile(uintptr(syscall.Stdin), "/dev/stdin")]
qm=常量_标准输入

[Stdout = NewFile(uintptr(syscall.Stdout), "/dev/stdout")]
qm=常量_标准输出

[Stderr = NewFile(uintptr(syscall.Stderr), "/dev/stderr")]
qm=常量_标准错误

[O_RDONLY int = syscall.O_RDONLY]
qm=常量_文件标志_只读打开

[O_WRONLY int = syscall.O_WRONLY]
qm=常量_文件标志_只写模式

[O_RDWR   int = syscall.O_RDWR]
qm=常量_文件标志_读写模式

[O_APPEND int = syscall.O_APPEND]
qm=常量_文件标志_追加模式

[O_CREATE int = syscall.O_CREAT]
qm=常量_文件标志_不存在则创建

[O_EXCL   int = syscall.O_EXCL]
qm=常量_文件标志_独占创建

[O_SYNC   int = syscall.O_SYNC]
qm=常量_文件标志_同步写入

[O_TRUNC  int = syscall.O_TRUNC]
qm=常量_文件标志_清空文件

[SEEK_SET int = 0]
qm=常量_位置_从头开始

[SEEK_CUR int = 1]
qm=常量_位置_从当前偏移量

[SEEK_END int = 2]
qm=常量_位置_从末尾处
