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

[func OpenKey(k Key, path string, access uint32) (Key, error) {]
ff=打开表项
access=访问权限
path=路径

[func OpenRemoteKey(pcname string, k Key) (Key, error) {]
ff=打开远程表项
pcname=计算机名

[func (k Key) ReadSubKeyNames(n int) (#左中括号##右中括号#string, error) {]
ff=取所有子项名称

[func CreateKey(k Key, path string, access uint32) (newk Key, openedExisting bool, err error) {]
ff=创建表项
err=错误
openedExisting=是否已存在
access=访问权限
path=路径

[func DeleteKey(k Key, path string) error {]
ff=删除表项
path=路径

[func (ki *KeyInfo) ModTime() time.Time {]
ff=取写入时间

[func (k Key) Stat() (*KeyInfo, error) {]
ff=取对象信息
