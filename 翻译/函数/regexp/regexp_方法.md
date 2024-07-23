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

[func (re *Regexp) Copy() *Regexp {]
ff=取副本

[func Compile(expr string) (*Regexp, error) {]
ff=创建
expr=表达式

[func CompilePOSIX(expr string) (*Regexp, error) {]
ff=创建并按POSIX
expr=表达式

[func (re *Regexp) Longest() {]
ff=最长匹配模式

[func MustCompile(str string) *Regexp {]
ff=创建PANI
str=表达式

[func MustCompilePOSIX(str string) *Regexp {]
ff=创建PANI并按POSIX
str=表达式

[func (re *Regexp) NumSubexp() int {]
ff=取子表达式数量

[func (re *Regexp) SubexpNames() #左中括号##右中括号#string {]
ff=取子表达式名称

[func (re *Regexp) SubexpIndex(name string) int {]
ff=取子表达式索引
name=名称

[func (re *Regexp) LiteralPrefix() (prefix string, complete bool) {]
ff=取表达式前缀
complete=成功
prefix=前缀

[func (re *Regexp) MatchReader(r io.RuneReader) bool {]
ff=是否匹配IO读取器
r=读取器

[func (re *Regexp) MatchString(s string) bool {]
ff=是否匹配文本
s=文本

[func (re *Regexp) Match(b #左中括号##右中括号#byte) bool {]
ff=是否匹配字节集
b=字节集

[func MatchReader(pattern string, r io.RuneReader) (matched bool, err error) {]
ff=是否匹配IO读取器
err=错误
matched=匹配成功
r=读取器
pattern=表达式

[func MatchString(pattern string, s string) (matched bool, err error) {]
ff=是否匹配文本
err=错误
matched=匹配成功
s=文本
pattern=表达式

[func Match(pattern string, b #左中括号##右中括号#byte) (matched bool, err error) {]
ff=是否匹配字节集
err=错误
matched=匹配成功
pattern=表达式

[func (re *Regexp) ReplaceAllString(src, repl string) string {]
ff=替换文本
repl=替换文本
src=原文本

[func (re *Regexp) ReplaceAllLiteralString(src, repl string) string {]
ff=替换字面文本
repl=替换文本
src=原文本

[func (re *Regexp) ReplaceAllStringFunc(src string, repl func(string) string) string {]
ff=替换字面文本FUN
repl=替换文本
src=原文本

[func (re *Regexp) ReplaceAll(src, repl #左中括号##右中括号#byte) #左中括号##右中括号#byte {]
ff=替换文本FUN
repl=替换文本
src=原文本

[func (re *Regexp) ReplaceAllLiteral(src, repl #左中括号##右中括号#byte) #左中括号##右中括号#byte {]
ff=替换字节集
repl=替换字节集
src=原字节集

[func (re *Regexp) ReplaceAllFunc(src #左中括号##右中括号#byte, repl func(#左中括号##右中括号#byte) #左中括号##右中括号#byte) #左中括号##右中括号#byte {]
ff=替换字面字节集
repl=替换字节集
src=原字节集

[func QuoteMeta(s string) string {]
ff=转义特殊字符
s=文本

[func (re *Regexp) Find(b #左中括号##右中括号#byte) #左中括号##右中括号#byte {]
ff=查找字节集
b=原字节集

[func (re *Regexp) FindIndex(b #左中括号##右中括号#byte) (loc #左中括号##右中括号#int) {]
ff=查找索引
loc=索引
b=原字节集

[func (re *Regexp) FindString(s string) string {]
ff=查找文本
s=原文本

[func (re *Regexp) FindStringIndex(s string) (loc #左中括号##右中括号#int) {]
ff=查找文本索引
loc=索引
s=原文本

[func (re *Regexp) FindReaderIndex(r io.RuneReader) (loc #左中括号##右中括号#int) {]
ff=查找IO读取器索引
loc=索引
r=IO读取器

[func (re *Regexp) FindSubmatch(b #左中括号##右中括号#byte) #左中括号##右中括号##左中括号##右中括号#byte {]
ff=查找字节集子匹配
b=原字节集

[func (re *Regexp) Expand(dst #左中括号##右中括号#byte, template #左中括号##右中括号#byte, src #左中括号##右中括号#byte, match #左中括号##右中括号#int) #左中括号##右中括号#byte {]
ff=模板替换字节集
template=模板

[func (re *Regexp) ExpandString(dst #左中括号##右中括号#byte, template string, src string, match #左中括号##右中括号#int) #左中括号##右中括号#byte {]
ff=模板替换文本
template=模板

[func (re *Regexp) FindSubmatchIndex(b #左中括号##右中括号#byte) #左中括号##右中括号#int {]
ff=查找字节集子匹配索引
b=原字节集

[func (re *Regexp) FindStringSubmatch(s string) #左中括号##右中括号#string {]
ff=查找文本子匹配
s=原文本

[func (re *Regexp) FindStringSubmatchIndex(s string) #左中括号##右中括号#int {]
ff=查找文本子匹配索引
s=原文本

[func (re *Regexp) FindReaderSubmatchIndex(r io.RuneReader) #左中括号##右中括号#int {]
ff=查找IO读取器子匹配索引
r=IO读取器

[func (re *Regexp) FindAll(b #左中括号##右中括号#byte, n int) #左中括号##右中括号##左中括号##右中括号#byte {]
ff=查找所有字节集
n=最大返回数
b=原字节集

[func (re *Regexp) FindAllIndex(b #左中括号##右中括号#byte, n int) #左中括号##右中括号##左中括号##右中括号#int {]
ff=查找所有字节集索引
n=最大返回数
b=原字节集

[func (re *Regexp) FindAllString(s string, n int) #左中括号##右中括号#string {]
ff=查找所有文本
n=最大返回数
s=原文本

[func (re *Regexp) FindAllStringIndex(s string, n int) #左中括号##右中括号##左中括号##右中括号#int {]
ff=查找所有文本索引
n=最大返回数
s=原文本

[func (re *Regexp) FindAllSubmatch(b #左中括号##右中括号#byte, n int) #左中括号##右中括号##左中括号##右中括号##左中括号##右中括号#byte {]
ff=查找所有字节集子匹配
n=最大返回数
b=原字节集

[func (re *Regexp) FindAllSubmatchIndex(b #左中括号##右中括号#byte, n int) #左中括号##右中括号##左中括号##右中括号#int {]
ff=查找所有字节集子匹配索引
n=最大返回数
b=原字节集

[func (re *Regexp) FindAllStringSubmatch(s string, n int) #左中括号##右中括号##左中括号##右中括号#string {]
ff=查找所有文本子匹配
n=最大返回数

[func (re *Regexp) FindAllStringSubmatchIndex(s string, n int) #左中括号##右中括号##左中括号##右中括号#int {]
ff=查找所有文本子匹配索引
n=最大返回数

[func (re *Regexp) Split(s string, n int) #左中括号##右中括号#string {]
ff=分割文本
n=最大返回数
s=原文本
