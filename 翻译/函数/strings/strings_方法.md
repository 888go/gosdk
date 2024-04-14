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

[func Count(s, substr string) int { //md5:8dc4508cfbea56c6d406f8a4f9770e09]
substr=检索文本
s=原文本
ff=取统计数

[func Contains(s, substr string) bool { //md5:e5367376f35a456e5c106746b475b47c]
ff=是否包含
substr=检索文本
s=原文本

[func ContainsAny(s, chars string) bool { //md5:3c6df9a7716addd6267b4d1f1e7d021c]
ff=是否包含任意字符
chars=检索字符s
s=原文本

[func ContainsRune(s string, r rune) bool { //md5:4238898d042dfb1776a3a6e6466358b2]
ff=是否包含字符
r=检索字符
s=原文本

[func ContainsFunc(s string, f func(rune) bool) bool { //md5:7f263561c791a82d7f0eb074a863c978]
ff=是否包含任意字符FUNC
f=回调函数
s=原文本

[func LastIndex(s, substr string) int { //md5:f71c5a442b965bb5bfaeed7e5c7c9576]
ff=取文本最后索引
substr=检索文本
s=原文本

[func IndexByte(s string, c byte) int { //md5:cc27e81169256c2d002e7ba6cc205863]
ff=取字节索引
c=检索字节
s=原文本

[func IndexRune(s string, r rune) int { //md5:ee5f1aa9ff80778af6ab99d3210cca6c]
ff=取字符索引
r=检索字符
s=原文本

[func IndexAny(s, chars string) int { //md5:9082d9461cfcdcfe489a822547f285e9]
ff=取任意字符索引
chars=检索字符s
s=原文本

[func LastIndexAny(s, chars string) int { //md5:b7c6c3d4157d4ce87b044ac9be7aed6f]
ff=取任意字符最后索引
chars=检索字符
s=原文本

[func LastIndexByte(s string, c byte) int { //md5:3dc6c63525e2007ee6cccdee606cb585]
ff=取字节最后索引
c=字节
s=原文本

[func SplitN(s, sep string, n int) #左中括号##右中括号#string { //md5:2ef42e1e1c8b9efaf8efedea2b1b2b58]
ff=分割文本并按数量
n=返回数量
sep=分隔符
s=原文本

[func SplitAfterN(s, sep string, n int) #左中括号##右中括号#string { //md5:5c4f7c8941408be7caa300dfd6629980]
ff=分割文本并按数量2
n=返回数量
sep=分隔符
s=原文本

[func Split(s, sep string) #左中括号##右中括号#string { //md5:9e39ccbb6ab10559dc6dfa7be187dc75]
ff=分割文本
sep=分隔符
s=原文本

[func SplitAfter(s, sep string) #左中括号##右中括号#string { //md5:7feaa482c8e32826bdbe24bf223ee6e7]
ff=分割文本2
sep=分隔符
s=原文本

[func Fields(s string) #左中括号##右中括号#string { //md5:1f8999413d2d4e0022170f15e0baae55]
ff=分割单词
s=原文本

[func FieldsFunc(s string, f func(rune) bool) #左中括号##右中括号#string { //md5:fb9dfb813e048fcf1d47059175b1f561]
ff=分割单词FUNC
f=回调函数
s=原文本

[func Join(elems #左中括号##右中括号#string, sep string) string { //md5:00ef0cc76f825778bdb1f1c6ab82a33d]
ff=连接
sep=连接符
elems=数组

[func HasPrefix(s, prefix string) bool { //md5:e1f46e88f0e4ea58e45b8de9709c6e08]
ff=是否为前缀
prefix=前缀
s=原文本

[func HasSuffix(s, suffix string) bool { //md5:1ddeca0a9e56022d90d75810b20fb6da]
ff=是否为后缀
suffix=后缀
s=原文本

[func Map(mapping func(rune) rune, s string) string { //md5:63ed0c242381265245671e75ecccb37f]
ff=遍历修改FUNC
s=原文本
mapping=回调函数

[func Repeat(s string, count int) string { //md5:d49e11180de57b1bad15b47471444a40]
ff=生成重复文本
count=重复次数
s=原文本

[func Title(s string) string { //md5:adaf62896a837fe0a4b94cef489e74c6]
ff=弃用Title
s=原文本

[func ToUpper(s string) string { //md5:3debb184e946c25c8eac8f6f13baedd6]
ff=到大写
s=原文本

[func ToLower(s string) string { //md5:474a927c2e62fa6c5288c04fd40232ce]
ff=到小写
s=原文本

[func ToTitle(s string) string { //md5:7cb166b3ddebfd1125f25ad2a7a6bc0f]
ff=到大写2
s=原文本

[func ToUpperSpecial(c unicode.SpecialCase, s string) string { //md5:b4e286fb9418724e1543c7af9b9e3af9]
ff=到大写并按规则
s=原文本
c=大小写规则

[func ToLowerSpecial(c unicode.SpecialCase, s string) string { //md5:dc9a6ec7663a80cdd586a2e9772618ac]
ff=到小写并按规则
s=原文本
c=大小写规则

[func ToTitleSpecial(c unicode.SpecialCase, s string) string { //md5:6eebd87af27dac15039ad70bc35e7318]
ff=到大写2并按规则
s=原文本
c=大小写规则

[func ToValidUTF8(s, replacement string) string { //md5:422f67249cb248a3bd9c0ab9b1d097e1]
ff=替换无效UTF8字节
replacement=替换符
s=原文本

[func TrimLeftFunc(s string, f func(rune) bool) string { //md5:516609201171fccebe5fce4ada48d1ad]
ff=删起始字符FUNC
f=回调函数
s=原文本

[func TrimRightFunc(s string, f func(rune) bool) string { //md5:280f57b7a6e49458e7ab99785d88cfc4]
ff=删尾部字符FUNC
f=回调函数
s=原文本

[func TrimFunc(s string, f func(rune) bool) string { //md5:f83f48e3eb47b62b7e1ecdc80aa8229b]
ff=删首尾字符FUNC
f=回调函数
s=原文本

[func IndexFunc(s string, f func(rune) bool) int { //md5:22c6b63d2713ff3020a8ced746a3bcdb]
f=回调函数
s=原文本

[func LastIndexFunc(s string, f func(rune) bool) int { //md5:c29da2b69561585e838689d903e72dac]
s=原文本

[func TrimRight(s, cutset string) string { //md5:5264a3ff78fb7d36909186e6a8dbef64]
s=原文本

[func TrimSpace(s string) string { //md5:3d7d6d251ac5a73c7e895d7eab45c28c]
ff=删首尾空
s=原文本

[func TrimPrefix(s, prefix string) string { //md5:7b252f4ef2548b2526b142c154ee698c]
ff=删前缀
prefix=前缀
s=原文本

[func TrimSuffix(s, suffix string) string { //md5:48e4d1998c003b0feefefd543d4a55dd]
ff=删后缀
suffix=后缀
s=原文本
