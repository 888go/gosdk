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

[func (b *Buffer) Bytes() #左中括号##右中括号#byte {]
ff=取缓冲区底层切片

[func (b *Buffer) Len() int {]
ff=取未读长度

[func (b *Buffer) Cap() int {]
ff=取容量

[func (b *Buffer) Truncate(n int) {]
ff=截断至长度
n=长度

[func (b *Buffer) Reset() {]
ff=清空

[func (b *Buffer) Grow(n int) {]
ff=扩展容量
n=扩展长度

[func (b *Buffer) Write(p #左中括号##右中括号#byte) (n int, err error) {]
err=错误
n=写入长度
p=写入字节集

[func (b *Buffer) WriteString(s string) (n int, err error) {]
ff=写入文本
err=错误
n=写入长度
s=文本

[func (b *Buffer) ReadFrom(r io.Reader) (n int64, err error) {]
ff=写入并从IO读取器
err=错误
n=写入长度
r=IO读取器

[func (b *Buffer) WriteTo(w io.Writer) (n int64, err error) {]
ff=读数据并写到IO写入器
err=错误
n=读取长度
w=IO写入器

[func (b *Buffer) WriteByte(c byte) error {]
ff=写字节
c=字节

[func (b *Buffer) WriteRune(r rune) (n int, err error) {]
ff=写字符
err=错误
n=写入长度
r=字符

[func (b *Buffer) Read(p #左中括号##右中括号#byte) (n int, err error) {]
err=错误
n=读取长度
p=字节集变量

[func (b *Buffer) Next(n int) #左中括号##右中括号#byte {]
ff=读取指定长度字节集

[func (b *Buffer) ReadByte() (byte, error) {]
ff=读取字节

[func (b *Buffer) ReadRune() (r rune, size int, err error) {]
ff=读取字符
err=错误
size=长度
r=字符

[func (b *Buffer) UnreadRune() error {]
ff=反向读取字符

[func (b *Buffer) UnreadByte() error {]
ff=反向读取字节

[func (b *Buffer) ReadBytes(delim byte) (line #左中括号##右中括号#byte, err error) {]
ff=读取至分隔符
err=错误
line=字节集
delim=分隔符

[func (b *Buffer) ReadString(delim byte) (line string, err error) {]
ff=读取文本至分隔符
err=错误
line=文本
delim=分隔符

[func NewBuffer(buf #左中括号##右中括号#byte) *Buffer {]
ff=创建缓冲区并按字节集
buf=字节集

[func NewBufferString(s string) *Buffer {]
ff=创建缓冲区并按文本
s=文本
