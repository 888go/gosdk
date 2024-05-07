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

[func Count(s, substr string) int {]
substr=检索文本
s=原文本
ff=取统计数

[func Contains(s, substr string) bool {]
ff=是否包含
substr=检索文本
s=原文本

[func ContainsAny(s, chars string) bool {]
ff=是否包含任意字符
chars=检索字符s
s=原文本

[func ContainsRune(s string, r rune) bool {]
ff=是否包含字符
r=检索字符
s=原文本

[func ContainsFunc(s string, f func(rune) bool) bool {]
ff=是否包含任意字符FUNC
f=回调函数
s=原文本

[func LastIndex(s, substr string) int {]
ff=取文本最后索引
substr=检索文本
s=原文本

[func IndexByte(s string, c byte) int {]
ff=取字节索引
c=检索字节
s=原文本

[func IndexRune(s string, r rune) int {]
ff=取字符索引
r=检索字符
s=原文本

[func IndexAny(s, chars string) int {]
ff=取任意字符索引
chars=检索字符s
s=原文本

[func LastIndexAny(s, chars string) int {]
ff=取任意字符最后索引
chars=检索字符
s=原文本

[func LastIndexByte(s string, c byte) int {]
ff=取字节最后索引
c=字节
s=原文本

[func SplitN(s, sep string, n int) #左中括号##右中括号#string {]
ff=分割文本并按数量
n=返回数量
sep=分隔符
s=原文本

[func SplitAfterN(s, sep string, n int) #左中括号##右中括号#string {]
ff=分割文本并按数量2
n=返回数量
sep=分隔符
s=原文本

[func Split(s, sep string) #左中括号##右中括号#string {]
ff=分割文本
sep=分隔符
s=原文本

[func SplitAfter(s, sep string) #左中括号##右中括号#string {]
ff=分割文本2
sep=分隔符
s=原文本

[func Fields(s string) #左中括号##右中括号#string {]
ff=分割单词
s=原文本

[func FieldsFunc(s string, f func(rune) bool) #左中括号##右中括号#string {]
ff=分割单词FUNC
f=回调函数
s=原文本

[func Join(elems #左中括号##右中括号#string, sep string) string {]
ff=连接
sep=连接符
elems=切片

[func HasPrefix(s, prefix string) bool {]
ff=是否为前缀
prefix=前缀
s=原文本

[func HasSuffix(s, suffix string) bool {]
ff=是否为后缀
suffix=后缀
s=原文本

[func Map(mapping func(rune) rune, s string) string {]
ff=遍历修改FUNC
s=原文本
mapping=回调函数

[func Repeat(s string, count int) string {]
ff=生成重复文本
count=重复次数
s=原文本

[func Title(s string) string {]
ff=弃用Title
s=原文本

[func ToUpper(s string) string {]
ff=到大写
s=原文本

[func ToLower(s string) string {]
ff=到小写
s=原文本

[func ToTitle(s string) string {]
ff=到大写2
s=原文本

[func ToUpperSpecial(c unicode.SpecialCase, s string) string {]
ff=到大写并按规则
s=原文本
c=大小写规则

[func ToLowerSpecial(c unicode.SpecialCase, s string) string {]
ff=到小写并按规则
s=原文本
c=大小写规则

[func ToTitleSpecial(c unicode.SpecialCase, s string) string {]
ff=到大写2并按规则
s=原文本
c=大小写规则

[func ToValidUTF8(s, replacement string) string {]
ff=替换无效UTF8字节
replacement=替换符
s=原文本

[func TrimLeftFunc(s string, f func(rune) bool) string {]
ff=删起始字符FUNC
f=回调函数
s=原文本

[func TrimRightFunc(s string, f func(rune) bool) string {]
ff=删尾部字符FUNC
f=回调函数
s=原文本

[func TrimFunc(s string, f func(rune) bool) string {]
ff=删首尾字符FUNC
f=回调函数
s=原文本

[func IndexFunc(s string, f func(rune) bool) int {]
f=回调函数
s=原文本

[func LastIndexFunc(s string, f func(rune) bool) int {]
s=原文本

[func TrimRight(s, cutset string) string {]
s=原文本

[func TrimSpace(s string) string {]
ff=删首尾空
s=原文本

[func TrimPrefix(s, prefix string) string {]
ff=删前缀
prefix=前缀
s=原文本

[func TrimSuffix(s, suffix string) string {]
ff=删后缀
suffix=后缀
s=原文本
