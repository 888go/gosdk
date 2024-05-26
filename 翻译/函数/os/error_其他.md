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

[ErrInvalid          = os.ErrInvalid]
qm=常量_错误_无效参数
cz=ErrInvalid          #等号# os.ErrInvalid

[ErrPermission       = os.ErrPermission]
qm=常量_错误_权限拒绝
cz=ErrPermission       #等号# os.ErrPermission

[ErrExist            = os.ErrExist]
qm=常量_错误_文件已存在
cz=ErrExist            #等号# os.ErrExist

[ErrNotExist         = os.ErrNotExist]
qm=常量_错误_文件不存在
cz=ErrNotExist         #等号# os.ErrNotExist

[ErrClosed           = os.ErrClosed]
qm=常量_错误_文件已关闭
cz=ErrClosed           #等号# os.ErrClosed

[ErrNoDeadline       = os.ErrNoDeadline]
qm=常量_错误_文件类型不支持超时截止
cz=ErrNoDeadline       #等号# os.ErrNoDeadline

[ErrDeadlineExceeded = os.ErrDeadlineExceeded]
qm=常量_错误_IO超时
cz=ErrDeadlineExceeded #等号# os.ErrDeadlineExceeded
