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

[func Equal(a, b #左中括号##右中括号#byte) bool {]
ff=是否相等
b=比较b
a=比较a

[func Compare(a, b #左中括号##右中括号#byte) int {]
ff=顺序比较
b=比较b
a=比较a

[func Count(s, sep #左中括号##右中括号#byte) int {]
ff=统计次数
sep=待统计
s=原字节集

[func Contains(b, subslice #左中括号##右中括号#byte) bool {]
ff=是否包含
subslice=待查找
b=原字节集

[func ContainsAny(b #左中括号##右中括号#byte, chars string) bool {]
ff=是否包含任意字符
chars=检索字符s
b=原字节集

[func ContainsRune(b #左中括号##右中括号#byte, r rune) bool {]
ff=是否包含Unicode字符
r=检索字符
b=原字节集

[func IndexByte(b #左中括号##右中括号#byte, c byte) int {]
ff=取字节索引
c=检索字节
b=原字节集

[func LastIndex(s, sep #左中括号##右中括号#byte) int {]
ff=取字节集最后索引
sep=待检索
s=原字节集

[func LastIndexByte(s #左中括号##右中括号#byte, c byte) int {]
ff=取字节最后索引
c=字节
s=原字节集

[func IndexRune(s #左中括号##右中括号#byte, r rune) int {]
ff=取Unicode字符索引
r=检索字符
s=原字节集

[func IndexAny(s #左中括号##右中括号#byte, chars string) int {]
ff=取任意字符索引
chars=检索字符s
s=原字节集

[func LastIndexAny(s #左中括号##右中括号#byte, chars string) int {]
ff=取任意字符最后索引
chars=检索字符
s=原字节集

[func SplitN(s, sep #左中括号##右中括号#byte, n int) #左中括号##右中括号##左中括号##右中括号#byte {]
ff=分割并按数量
n=返回数量
sep=分隔符
s=原字节集

[func SplitAfterN(s, sep #左中括号##右中括号#byte, n int) #左中括号##右中括号##左中括号##右中括号#byte {]
ff=分割并按数量2
n=返回数量
sep=分隔符
s=原字节集

[func Split(s, sep #左中括号##右中括号#byte) #左中括号##右中括号##左中括号##右中括号#byte {]
ff=分割
sep=分隔符
s=原字节集

[func SplitAfter(s, sep #左中括号##右中括号#byte) #左中括号##右中括号##左中括号##右中括号#byte {]
ff=分割2
sep=分隔符
s=原字节集

[func Fields(s #左中括号##右中括号#byte) #左中括号##右中括号##左中括号##右中括号#byte {]
ff=分割并按空白
s=原字节集

[func FieldsFunc(s #左中括号##右中括号#byte, f func(rune) bool) #左中括号##右中括号##左中括号##右中括号#byte {]
ff=分割FUNC
f=回调函数
s=原字节集

[func Join(s #左中括号##右中括号##左中括号##右中括号#byte, sep #左中括号##右中括号#byte) #左中括号##右中括号#byte {]
ff=连接
sep=连接符
s=字节集切片

[func HasPrefix(s, prefix #左中括号##右中括号#byte) bool {]
ff=是否为前缀
prefix=前缀
s=原字节集

[func HasSuffix(s, suffix #左中括号##右中括号#byte) bool {]
ff=是否为后缀
suffix=后缀
s=原字节集

[func Map(mapping func(r rune) rune, s #左中括号##右中括号#byte) #左中括号##右中括号#byte {]
ff=遍历修改FUNC
s=原字节集
mapping=回调函数

[func Repeat(b #左中括号##右中括号#byte, count int) #左中括号##右中括号#byte {]
ff=生成重复字节集
count=重复次数
b=原字节集

[func ToUpper(s #左中括号##右中括号#byte) #左中括号##右中括号#byte {]
ff=到大写
s=原字节集

[func ToLower(s #左中括号##右中括号#byte) #左中括号##右中括号#byte {]
ff=到小写
s=原字节集

[func ToTitle(s #左中括号##右中括号#byte) #左中括号##右中括号#byte {]
ff=到大写2
s=原字节集

[func ToUpperSpecial(c unicode.SpecialCase, s #左中括号##右中括号#byte) #左中括号##右中括号#byte {]
ff=到大写并按规则
s=原字节集
c=大小写规则

[func ToLowerSpecial(c unicode.SpecialCase, s #左中括号##右中括号#byte) #左中括号##右中括号#byte {]
ff=到小写并按规则
s=原字节集
c=大小写规则

[func ToTitleSpecial(c unicode.SpecialCase, s #左中括号##右中括号#byte) #左中括号##右中括号#byte {]
ff=到大写2并按规则
s=原字节集
c=大小写规则

[func ToValidUTF8(s, replacement #左中括号##右中括号#byte) #左中括号##右中括号#byte {]
ff=替换无效UTF8字节
replacement=替换符
s=原字节集

[func Title(s #左中括号##右中括号#byte) #左中括号##右中括号#byte {]
ff=弃用Title
s=原字节集

[func TrimLeftFunc(s #左中括号##右中括号#byte, f func(r rune) bool) #左中括号##右中括号#byte {]
ff=删起始字符FUNC
f=回调函数
s=原字节集

[func TrimRightFunc(s #左中括号##右中括号#byte, f func(r rune) bool) #左中括号##右中括号#byte {]
ff=删尾部字符FUNC
f=回调函数
s=原字节集

[func TrimFunc(s #左中括号##右中括号#byte, f func(r rune) bool) #左中括号##右中括号#byte {]
ff=删首尾字符FUNC
f=回调函数
s=原字节集

[func TrimPrefix(s, prefix #左中括号##右中括号#byte) #左中括号##右中括号#byte {]
ff=删前缀
prefix=前缀
s=原字节集

[func TrimSuffix(s, suffix #左中括号##右中括号#byte) #左中括号##右中括号#byte {]
ff=删后缀
suffix=后缀
s=原字节集

[func IndexFunc(s #左中括号##右中括号#byte, f func(r rune) bool) int {]
ff=查找字符索引FUNC
f=回调函数
s=原字节集

[func LastIndexFunc(s #左中括号##右中括号#byte, f func(r rune) bool) int {]
ff=查找最后字符索引FUNC
f=回调函数
s=原字节集

[func Trim(s #左中括号##右中括号#byte, cutset string) #左中括号##右中括号#byte {]
ff=删首尾字符
s=原字节集

[func TrimLeft(s #左中括号##右中括号#byte, cutset string) #左中括号##右中括号#byte {]
ff=删起始字符
s=原字节集

[func TrimRight(s #左中括号##右中括号#byte, cutset string) #左中括号##右中括号#byte {]
ff=删尾部字符
s=原字节集

[func TrimSpace(s #左中括号##右中括号#byte) #左中括号##右中括号#byte {]
ff=删首尾空
s=原字节集

[func Runes(s #左中括号##右中括号#byte) #左中括号##右中括号#rune {]
ff=转换为Unicode字符切片
s=字节集

[func Replace(s, old, new #左中括号##右中括号#byte, n int) #左中括号##右中括号#byte {]
ff=替换
n=次数
new=新字节集
old=旧字节集
s=原字节集

[func ReplaceAll(s, old, new #左中括号##右中括号#byte) #左中括号##右中括号#byte {]
ff=替换全部
new=新字节集
old=旧字节集
s=原字节集

[func EqualFold(s, t #左中括号##右中括号#byte) bool {]
ff=是否相等并忽略大小写

[func Index(s, sep #左中括号##右中括号#byte) int {]
ff=取字节集索引
sep=待查找
s=原字节集

[func Clone(b #左中括号##右中括号#byte) #左中括号##右中括号#byte {]
ff=深拷贝
b=原字节集

[func CutPrefix(s, prefix #左中括号##右中括号#byte) (after #左中括号##右中括号#byte, found bool) {]
ff=删前缀2
found=是否成功
after=返回
prefix=前缀
s=原字节集

[func CutSuffix(s, suffix #左中括号##右中括号#byte) (before #左中括号##右中括号#byte, found bool) {]
ff=删后缀2
found=是否成功
before=返回
suffix=后缀
s=原字节集
