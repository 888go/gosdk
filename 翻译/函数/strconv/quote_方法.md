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
# re.F.String()
# }
# //zj:
# 备注结束

[func Quote(s string) string {]
ff=转义到go字面量
s=文本

[func AppendQuote(dst #左中括号##右中括号#byte, s string) #左中括号##右中括号#byte {]
ff=转义到go字面量并加到字节集
s=文本
dst=字节集

[func QuoteToASCII(s string) string {]
ff=转义非ASCII
s=文本

[func AppendQuoteToASCII(dst #左中括号##右中括号#byte, s string) #左中括号##右中括号#byte {]
ff=转义非ASCII并加到字节集
s=文本
dst=字节集

[func QuoteToGraphic(s string) string {]
ff=转义到图形字符
s=文本

[func AppendQuoteToGraphic(dst #左中括号##右中括号#byte, s string) #左中括号##右中括号#byte {]
ff=转义到图形字符并加到字节集
s=文本
dst=字节集

[func QuoteRune(r rune) string {]
ff=转义字符到go字面量
r=字符

[func AppendQuoteRune(dst #左中括号##右中括号#byte, r rune) #左中括号##右中括号#byte {]
ff=转义字符到go字面量并加到字节集
r=字符
dst=字节集

[func QuoteRuneToASCII(r rune) string {]
ff=转义非ASCII字符
r=字符

[func AppendQuoteRuneToASCII(dst #左中括号##右中括号#byte, r rune) #左中括号##右中括号#byte {]
ff=转义非ASCII字符并加到字节集
r=字符
dst=字节集

[func QuoteRuneToGraphic(r rune) string {]
ff=转义字符到图形字符
r=字符

[func AppendQuoteRuneToGraphic(dst #左中括号##右中括号#byte, r rune) #左中括号##右中括号#byte {]
ff=转义字符到图形字符并加到字节集
r=字符
dst=字节集

[func CanBackquote(s string) bool {]
ff=是否包含反引号或非ASCII
s=文本

[func UnquoteChar(s string, quote byte) (value rune, multibyte bool, tail string, err error) {]
ff=解析字符转义
err=错误
tail=剩余文本
multibyte=是否为多字节字符
value=返回字符
quote=引号字节
s=文本

[func QuotedPrefix(s string) (string, error) {]
ff=解析引号包裹前缀
s=文本

[func Unquote(s string) (string, error) {]
ff=解析go字面量转义
s=文本

[func IsPrint(r rune) bool {]
ff=是否为可打印字符
r=字符

[func IsGraphic(r rune) bool {]
ff=是否为图形字符
r=字符
