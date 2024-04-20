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

[func (k Key) GetValue(name string, buf #左中括号##右中括号#byte) (n int, valtype uint32, err error) {]
ff=取值
err=错误
valtype=值类型
buf=缓冲区
name=名称

[func (k Key) GetStringValue(name string) (val string, valtype uint32, err error) {]
ff=取文本值
err=错误
valtype=值类型
val=值
name=名称

[func (k Key) GetMUIStringValue(name string) (string, error) {]
ff=取文本值PANI
name=名称

[func ExpandString(value string) (string, error) {]
ff=解析环境变量
value=值

[func (k Key) GetStringsValue(name string) (val #左中括号##右中括号#string, valtype uint32, err error) {]
ff=取文本数组值
err=错误
valtype=值类型
val=数组值
name=名称

[func (k Key) GetIntegerValue(name string) (val uint64, valtype uint32, err error) {]
ff=取整数64位值
err=错误
valtype=值类型
val=值
name=名称

[func (k Key) GetBinaryValue(name string) (val #左中括号##右中括号#byte, valtype uint32, err error) {]
ff=取字节集值
err=错误
valtype=值类型
val=值
name=名称

[func (k Key) SetDWordValue(name string, value uint32) error {]
ff=取整数32位值
value=值
name=名称

[func (k Key) SetQWordValue(name string, value uint64) error {]
ff=设置整数64位值
value=值
name=名称

[func (k Key) SetStringValue(name, value string) error {]
ff=设置文本值
value=值
name=名称

[func (k Key) SetExpandStringValue(name, value string) error {]
ff=设置文本值并按环境变量
value=值
name=名称

[func (k Key) SetStringsValue(name string, value #左中括号##右中括号#string) error {]
ff=设置文本数组值
value=数组值
name=名称

[func (k Key) SetBinaryValue(name string, value #左中括号##右中括号#byte) error {]
ff=设置字节集值
value=值
name=名称

[func (k Key) DeleteValue(name string) error {]
ff=删除值
name=名称

[func (k Key) ReadValueNames(n int) (#左中括号##右中括号#string, error) {]
ff=取所有子项值
n=返回数量
