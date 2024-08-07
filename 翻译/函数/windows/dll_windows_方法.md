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
# zz= 正则查找,配合前面/后面使用, 有设置正则查找,就不用设置上面的查找
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
# //zj:前面一行的代码,如果为空,追加到末尾行
# func (re *Regexp) X取文本() string { 
# re.F.String()
# }
# //zj:
# 备注结束

[func (e *DLLError) Error() string { return e.Msg }]
ff=取错误文本

[func (e *DLLError) Unwrap() error { return e.Err }]
ff=取错误

[func LoadDLL(name string) (dll *DLL, err error) {]
ff=DLL加载
err=错误
name=dll名称

[func MustLoadDLL(name string) *DLL {]
ff=DLL加载PANI
name=dll名称

[func (d *DLL) FindProc(name string) (proc *Proc, err error) {]
ff=查找命令
err=错误
name=命令名

[func (d *DLL) MustFindProc(name string) *Proc {]
ff=查找命令PANI
name=命令名

[func (d *DLL) FindProcByOrdinal(ordinal uintptr) (proc *Proc, err error) {]
ff=查找命令并按序数
err=错误
ordinal=序数

[func (d *DLL) MustFindProcByOrdinal(ordinal uintptr) *Proc {]
ff=查找命令并按序数PANI
ordinal=序数

[func (d *DLL) Release() (err error) {]
ff=卸载
err=错误

[func (p *Proc) Addr() uintptr {]
ff=取命令地址

[func (p *Proc) Call(a ...uintptr) (r1, r2 uintptr, lastErr error) {]
ff=调用
lastErr=最后一个错误
a=命令参数s

[func (d *LazyDLL) Load() error {]
ff=加载

[func (d *LazyDLL) Handle() uintptr {]
ff=取模块地址

[func (d *LazyDLL) NewProc(name string) *LazyProc {]
ff=创建命令对象
name=命令名

[func NewLazyDLL(name string) *LazyDLL {]
ff=DLL惰性加载
name=dll名称

[func NewLazySystemDLL(name string) *LazyDLL {]
ff=DLL惰性加载并从System
name=dll名称

[func (p *LazyProc) Find() error {]
ff=查找命令

[func (p *LazyProc) Addr() uintptr {]
ff=取命令地址

[func (p *LazyProc) Call(a ...uintptr) (r1, r2 uintptr, lastErr error) {]
ff=调用
lastErr=最后一个错误
a=命令参数s

[func (s errString) Error() string { return string(s) }]
ff=取错误文本
