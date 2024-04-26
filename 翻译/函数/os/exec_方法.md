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

[func Getpid() int {]
ff=取当前进程ID

[func Getppid() int {]
ff=取父进程ID

[func FindProcess(pid int) (*Process, error) {]
ff=查找进程
pid=进程id

[func StartProcess(name string, argv #左中括号##右中括号#string, attr *ProcAttr) (*Process, error) {]
ff=启动进程
attr=属性
argv=参数
name=执行路径或名称

[func (p *Process) Release() error {]
ff=释放进程资源

[func (p *Process) Kill() error {]
ff=终止进程

[func (p *Process) Wait() (*ProcessState, error) {]
ff=等待进程

[func (p *ProcessState) UserTime() time.Duration {]
ff=取用户CPU时间

[func (p *ProcessState) SystemTime() time.Duration {]
ff=取系统CPU时间

[func (p *ProcessState) Exited() bool {]
ff=是否已退出

[func (p *ProcessState) Success() bool {]
ff=是否成功退出

[func (p *ProcessState) Sys() any {]
ff=取退出信息

[func (p *ProcessState) SysUsage() any {]
ff=取资源使用信息

[func (p *ProcessState) Pid() int {]
ff=取已退出进程ID

[func (p *ProcessState) ExitCode() int {]
ff=取已退出进程代码
