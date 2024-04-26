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

[func (e *SyscallError) Error() string {]
ff=取错误文本

[func (e *SyscallError) Unwrap() error {]
ff=取错误对象

[func (e *SyscallError) Timeout() bool {]
ff=是否超时

[func NewSyscallError(syscall string, err error) error {]
ff=创建调用错误
err=错误
syscall=错误名称

[func IsExist(err error) bool {]
ff=是否文件已存在
err=错误

[func IsNotExist(err error) bool {]
ff=是否文件不存在
err=错误

[func IsPermission(err error) bool {]
ff=是否权限拒绝
err=错误

[func IsTimeout(err error) bool {]
ff=是否超时错误
err=错误
