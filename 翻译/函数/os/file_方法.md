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

[func (f *File) Name() string {]
ff=取文件名

[func (e *LinkError) Error() string {]
ff=取错误文本

[func (e *LinkError) Unwrap() error {]
ff=取错误对象

[func (f *File) Read(b #左中括号##右中括号#byte) (n int, err error) {]
ff=读取
err=错误
n=读取长度
b=缓冲区变量

[func (f *File) ReadAt(b #左中括号##右中括号#byte, off int64) (n int, err error) {]
ff=读取并按偏移量
err=错误
n=读取长度
off=偏移量
b=缓冲区变量

[func (f *File) ReadFrom(r io.Reader) (n int64, err error) {]
ff=写数据并从ioReader
err=错误
n=写入长度

[func (f *File) Write(b #左中括号##右中括号#byte) (n int, err error) {]
ff=写字节集
err=错误
n=写入长度
b=字节集

[func (f *File) WriteAt(b #左中括号##右中括号#byte, off int64) (n int, err error) {]
ff=写字节集并按偏移量
err=错误
n=写入长度
off=偏移量
b=字节集

[func (f *File) WriteTo(w io.Writer) (n int64, err error) {]
ff=写数据并从IoWriter
err=错误
n=写入长度

[func (f *File) Seek(offset int64, whence int) (ret int64, err error) {]
ff=设置偏移量
err=错误
whence=常量_位置_
offset=偏移量

[func (f *File) WriteString(s string) (n int, err error) {]
ff=写文本
err=错误
n=写入长度
s=文本

[func Mkdir(name string, perm FileMode) error {]
ff=目录创建
perm=权限
name=路径

[func Chdir(dir string) error {]
ff=修改当前工作目录
dir=路径

[func Open(name string) (*File, error) {]
ff=打开
name=文件路径

[func Create(name string) (*File, error) {]
ff=文件创建
name=文件路径

[func OpenFile(name string, flag int, perm FileMode) (*File, error) {]
ff=打开文件
perm=权限
flag=常量_文件标志
name=路径

[func Rename(oldpath, newpath string) error {]
ff=重命名
newpath=新路径
oldpath=旧路径

[func Readlink(name string) (string, error) {]
ff=取符号链接地址
name=路径

[func TempDir() string {]
ff=取临时目录

[func UserCacheDir() (string, error) {]
ff=取用户缓存目录

[func UserConfigDir() (string, error) {]
ff=取用户配置目录

[func UserHomeDir() (string, error) {]
ff=取用户主目录

[func Chmod(name string, mode FileMode) error {]
ff=修改权限
mode=权限
name=文件路径

[func (f *File) Chmod(mode FileMode) error {]
ff=修改权限
mode=权限

[func (f *File) SetDeadline(t time.Time) error {]
ff=设置超时

[func (f *File) SetReadDeadline(t time.Time) error {]
ff=设置读取超时

[func (f *File) SetWriteDeadline(t time.Time) error {]
ff=设置写入超时

[func ReadFile(name string) (#左中括号##右中括号#byte, error) {]
ff=读文件
name=路径

[func WriteFile(name string, data #左中括号##右中括号#byte, perm FileMode) error {]
ff=写文件
perm=权限
data=字节集
name=路径

[func (f *File) Fd() uintptr {]
ff=取文件描述符

[func NewFile(fd uintptr, name string) *File {]
ff=创建文件对象
name=路径
fd=文件描述符

[func (f *File) Close() error {]
ff=关闭

[func (f *File) Stat() (FileInfo, error) {]
ff=取文件信息

[func (f *File) Truncate(size int64) error {]
ff=修改文件大小
size=长度

[func (f *File) Sync() error {]
ff=同步数据

[func Truncate(name string, size int64) error {]
ff=修改文件大小
size=长度
name=文件路径

[func Remove(name string) error {]
ff=删除文件
name=路径

[func Chtimes(name string, atime time.Time, mtime time.Time) error {]
ff=修改文件时间
mtime=修改时间
atime=访问时间
name=文件路径

[func Pipe() (r *File, w *File, err error) {]
ff=创建管道
err=错误
w=写入端
r=读取端

[func Link(oldname, newname string) error {]
ff=创建硬链接
newname=新路径
oldname=旧路径

[func Symlink(oldname, newname string) error {]
ff=创建符号链接
newname=新路径
oldname=旧路径

[func Chown(name string, uid, gid int) error {]
ff=修改文件所有者
gid=组id
uid=用户id
name=文件路径

[func (f *File) Chown(uid, gid int) error {]
ff=修改文件所有者
gid=组id
uid=用户id

[func Lchown(name string, uid, gid int) error {]
ff=修改文件所有者2
gid=组id
uid=用户id
name=文件路径

[func (f *File) Chdir() error {]
ff=修改当前工作目录到文件对象
