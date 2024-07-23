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

[末尾]
zj=// zj:#换行#// String 方法返回缓冲区未读部分内容作为字符串。#换行#// 如果 [Buffer] 是一个空指针，它将返回 "<nil>"。#换行#//#换行#// 要更高效地构建字符串，请参见 strings.Builder 类型。#换行#func (b *Buffer) X取未读文本() string { //md5:2a07d671bd398c91#换行##TAB#return b.F.String()#换行#}#换行##换行#// Read 从缓冲区中读取接下来的 len(p) 个字节，或者直到缓冲区被读空。返回值 n 是读取的字节数。如果缓冲区没有数据可读，err 将是 io.EOF（除非 len(p) 为零）；否则 err 为 nil。#换行#func (b *Buffer) X读取字节集(字节集变量 []byte) (读取长度 int, 错误 error) { //md5:d77b54302f43331e#换行##TAB#return b.F.Read(字节集变量)#换行#}#换行##换行#// Write 将 p 的内容追加到缓冲区，如果需要会扩大缓冲区。返回值 n 是 p 的长度；err 总是为 nil。如果缓冲区变得过大，Write 会引发一个 [ErrTooLarge] 的 panic。#换行#func (b *Buffer) X写入字节集(写入字节集 []byte) (写入长度 int, 错误 error) { //md5:6f1eceb67804b19f#换行##TAB#return b.F.Write(写入字节集)#换行#}#换行##换行#//zj:
