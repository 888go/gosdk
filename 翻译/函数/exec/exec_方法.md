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

[func Command(name string, arg ...string) *Cmd {]
arg=命令参数
name=进程名
ff=设置命令

[func CommandContext(ctx context.Context, name string, arg ...string) *Cmd {]
ff=设置命令并带上下文
arg=命令参数
name=进程名
ctx=上下文


[func (c *Cmd) Run() error {]
ff=运行

[func (c *Cmd) Start() error {]
ff=运行并按异步

[func (c *Cmd) Wait() error {]
ff=等待运行完成

[func (c *Cmd) Output() (#左中括号##右中括号#byte, error) {]
ff=运行并带返回值

[func (c *Cmd) CombinedOutput() (#左中括号##右中括号#byte, error) {]
ff=运行并带组合返回值

[func (c *Cmd) StdinPipe() (io.WriteCloser, error) {]
ff=取Stdin管道

[func (c *Cmd) StdoutPipe() (io.ReadCloser, error) {]
ff=取标准管道

[func (c *Cmd) StderrPipe() (io.ReadCloser, error) {]
ff=取Stderr管道

[func (c *Cmd) Environ() #左中括号##右中括号#string {]
ff=取环境变量切片
