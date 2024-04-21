package os

// LinkError 记录链接、符号链接或重命名系统调用期间发生的错误，以及导致错误的路径。
type LinkError struct { //md5:e482108b96c3226d02e9e4db482a6af0

}

// Name 返回传递给 Open 的文件名。
func (f *File) Name() string { //md5:797d96c447cb1cd703eea2eba5e775ca

}


func (e *LinkError) Error() string { //md5:d31cca2335c1e2d94b14bb50d57d1422

}


func (e *LinkError) Unwrap() error { //md5:cd73098610689f7781bd0c37ecae1a6b

}

// Read从File读取最多len(b)字节，并将它们存储在b中。它返回已读取的字节数和遇到的任何错误。在文件结束时，Read返回0和io.EOF。
func (f *File) Read(b []byte) (n int, err error) { //md5:ad71ac702ff91fdcd93706e2f46dd3a2

}

// ReadAt 从File的字节偏移量off处开始读取len(b)个字节。
// 它返回读取的字节数以及任何错误。当n < len(b)时，ReadAt总会返回一个非nil的错误。
// 在文件末尾，该错误为io.EOF。
func (f *File) ReadAt(b []byte, off int64) (n int, err error) { //md5:d30542dbb7bb2327308ded03baf9a4e6

}

// ReadFrom 实现了 io.ReaderFrom 接口。
func (f *File) ReadFrom(r io.Reader) (n int64, err error) { //md5:e15d3a79d68be3fcd1fd179d8b774a60

}

// ReadFrom隐藏了另一个ReadFrom方法。它不应该被调用。
func (noReadFrom) ReadFrom(io.Reader) (int64, error) { //md5:c7bd7153fd3be890d8470fb88d35846f

}

// Write 从b中写入len(b)字节到File。它返回写入的字节数和任何错误。当n不等于len(b)时，Write返回非nil错误。
func (f *File) Write(b []byte) (n int, err error) { //md5:f10e7f2db4e8de222a935f06c34b614c

}

// WriteAt 从字节偏移量 off 开始向 File 写入 len(b) 个字节。
// 它返回已写入的字节数以及任何可能遇到的错误。
// 当 n != len(b) 时，WriteAt 返回一个非 nil 错误。
//
// 若文件以 O_APPEND 标志打开，WriteAt 将返回一个错误。
func (f *File) WriteAt(b []byte, off int64) (n int, err error) { //md5:80c9238c8125967bb904a4febd529e43

}

// WriteTo 实现了 io.WriterTo 接口。
func (f *File) WriteTo(w io.Writer) (n int64, err error) { //md5:708734e7024db013e0ed863a3d3737e3

}

// WriteTo 隐藏了另一个 WriteTo 方法。它不应该被调用。
func (noWriteTo) WriteTo(io.Writer) (int64, error) { //md5:d9bb541a0dcba5ad5c301c2af1f5334f

}

// Seek 设置文件的下一个 Read 或 Write 操作的偏移量，根据 whence 的值进行解释：0 表示相对于文件的起始位置，1 表示相对于当前偏移量，2 表示相对于文件末尾。它返回新的偏移量和任何错误。对于使用 O_APPEND 模式打开的文件，Seek 的行为未做规定。
func (f *File) Seek(offset int64, whence int) (ret int64, err error) { //md5:139f47ea5ff396841d16378f9b4f5757

}

// WriteString 类似于 Write，但其写入的是字符串 s 的内容，而非字节片。
func (f *File) WriteString(s string) (n int, err error) { //md5:98c1882151f4a7bf3e126fc8aa052a98

}

// Mkdir 创建一个具有指定名称和权限位（在umask之前）的新目录。
// 如果发生错误，该错误将为 *PathError 类型。
func Mkdir(name string, perm FileMode) error { //md5:c2f7f3e5c608f078329bca6aad4968a7

}

// Chdir 将当前工作目录更改为指定的目录。如果出现错误，错误类型为 *PathError。
func Chdir(dir string) error { //md5:1122edf88c5fabc8721b76f3f864f007

}

// Open 打开指定的文件进行读取。如果成功，可以使用返回文件的方法进行读取；关联的文件描述符的模式为 O_RDONLY。
// 如果出现错误，错误类型为 *PathError。
func Open(name string) (*File, error) { //md5:24216d7be3066e53419929b0c658593e

}

// Create 用于创建或截断指定名称的文件。若该文件已存在，将会被截断；若不存在，则会以0666（应用umask之前）的模式创建。操作成功后，可使用返回的File对象进行I/O操作，其关联的文件描述符具有O_RDWR模式。如果出现错误，其类型将为*PathError。
func Create(name string) (*File, error) { //md5:0901909cad4273888ab622dc826bfd50

}

// OpenFile 是通用的打开文件调用；大多数用户应使用 Open 或 Create。它以指定的标志（如 O_RDONLY 等）
// 打开名为 name 的文件。如果文件不存在，且传递了 O_CREATE 标志，那么会使用权限 perm（在 umask 之前）
// 创建该文件。如果成功，返回的 File 对象上的方法可用于 I/O 操作。
// 如果出现错误，该错误将为 *PathError 类型。
func OpenFile(name string, flag int, perm FileMode) (*File, error) { //md5:c46ec9700a28b8c79cd66cc5758fd324

}

// Rename 将 oldpath 重命名（移动）到 newpath。如果 newpath 已经存在且不是目录，Rename 将替换它。当 oldpath 和 newpath 在不同的目录时，可能存在特定于操作系统的限制。即使在同一个目录内，在非 Unix 平台上，Rename 也不是一个原子操作。如果出现错误，错误类型将为 *LinkError。
func Rename(oldpath, newpath string) error { //md5:68de958c41b824e8588a1376c5bd725a

}

// Readlink 返回指定符号链接的目标位置。
// 如果出现错误，错误类型为 *PathError。
// 
// 如果链接目标是相对路径，Readlink 会返回未解析为绝对路径的相对路径。
func Readlink(name string) (string, error) { //md5:a73610fda1bea460f5e497c22a7b1b03

}

// TempDir 返回用于临时文件的默认目录。
//
// 在 Unix 系统中，如果 $TMPDIR 非空，则返回其值，否则返回 /tmp。
// 在 Windows 上，它使用 GetTempPath 函数，按顺序返回以下各项中首个非空值：
// %TMP%、%TEMP%、%USERPROFILE% 或 Windows 目录。
// 在 Plan 9 系统上，返回 /tmp。
//
// 此目录并不保证一定存在或具有可访问权限。
func TempDir() string { //md5:1d1c94fcaa742437bb67f366b8c5326b

}

// UserCacheDir 返回用于用户特定缓存数据的默认根目录。用户应在该目录下创建他们自己的应用程序特定子目录并使用它。
//
// 在Unix系统上，如果$XDG_CACHE_HOME非空（如https://specifications.freedesktop.org/basedir-spec/basedir-spec-latest.html中指定），则返回$XDG_CACHE_HOME，否则返回$HOME/.cache。
// 在Darwin系统上，返回$HOME/Library/Caches。
// 在Windows系统上，返回%LocalAppData%。
// 在Plan 9系统上，返回$home/lib/cache。
//
// 如果无法确定位置（例如，$HOME未定义），则会返回错误。
func UserCacheDir() (string, error) { //md5:bf893e2b605145911aca52a346adafc0

}

// UserConfigDir 返回默认的用户特定配置数据根目录。用户应在该目录下创建自己应用的子目录，并使用这个子目录。
// 
// 在 Unix 系统中，它会返回 $XDG_CONFIG_HOME，如 https://specifications.freedesktop.org/basedir-spec/basedir-spec-latest.html 中指定的，如果非空，则否则返回 $HOME/.config。
// 在 Darwin（macOS）上，它返回 $HOME/Library/Application Support。
// 在 Windows 上，它返回 %AppData%。
// 在 Plan 9 系统中，它返回 $home/lib。
// 
// 如果无法确定位置（例如，$HOME 未定义），则会返回错误。
func UserConfigDir() (string, error) { //md5:322fb225f6945ecdd7703dca7bed8ed6

}

// UserHomeDir 返回当前用户的主目录。
//
// 在 Unix（包括 macOS）中，它返回 $HOME 环境变量。在 Windows 中，它返回 %USERPROFILE%。
// 在 Plan 9 中，它返回 $home 环境变量。
//
// 如果预期的环境变量未设置，UserHomeDir 将返回一个平台特定的默认值或非空错误。
func UserHomeDir() (string, error) { //md5:283aa4068141ca3262d0061495f7c9ba

}

// Chmod 修改指定文件名的文件模式为 mode。
// 若该文件为符号链接，则会更改其目标对象的模式。
// 若出现错误，其类型将为 *PathError。
//
// 操作系统不同，所使用的模式位子集也会不同。
//
// 在 Unix 系统中，使用 mode 的权限位（permission bits）、ModeSetuid、ModeSetgid 和 ModeSticky。
//
// 在 Windows 系统中，仅使用 mode 的 0200 位（所有者可写）；它控制文件的只读属性是否设置或清除。其余位当前未使用。为了与 Go 1.12 及更早版本兼容，使用非零 mode。对于只读文件，使用模式 0400；对于可读写文件，使用模式 0600。
//
// 在 Plan 9 系统中，使用 mode 的权限位、ModeAppend、ModeExclusive 和 ModeTemporary。
func Chmod(name string, mode FileMode) error { //md5:e6e4ce38e325608199ffcfdfb9718973

}

// Chmod将文件的模式更改为mode。
// 如果出现错误，它将是一个类型为*PathError的错误。
func (f *File) Chmod(mode FileMode) error { //md5:d2d0e79d92cc7b72a76de872be1ce3bd

}

// SetDeadline 为一个 File 设置读写截止时间。
// 相当于同时调用 SetReadDeadline 和 SetWriteDeadline。
//
// 只有某些类型的文件支持设置截止时间。对于不支持截止时间的文件，调用 SetDeadline 将返回 ErrNoDeadline。
// 在大多数系统中，普通文件不支持截止时间，但管道（pipes）支持。
//
// 截止时间是指在该时间点之后，I/O 操作将不再阻塞而是返回错误。这一截止时间适用于所有未来和待处理的 I/O，
// 而不仅仅是紧接着的 Read 或 Write 调用。一旦截止时间被超过，可以通过设置一个未来的时间来刷新连接。
//
// 如果截止时间被超过，对 Read、Write 或其他 I/O 方法的调用将返回一个封装了 ErrDeadlineExceeded 的错误。
// 可以使用 errors.Is(err, os.ErrDeadlineExceeded) 来检测这个错误。该错误实现了 Timeout 方法，
// 调用 Timeout 方法将返回 true，但存在其他可能的错误，即使截止时间未被超过，其 Timeout 也可能返回 true。
//
// 通过在成功读写后反复延长截止时间，可以实现空闲超时机制。
//
// 若 t 值为零，则 I/O 操作不会超时。
func (f *File) SetDeadline(t time.Time) error { //md5:17a25a847d7089ad07b254849d19f407

}

// SetReadDeadline 设置了未来Read调用和当前阻塞的Read调用的超时时间。
// 如果t的值为零，表示Read不会超时。
// 并非所有文件都支持设置超时；请参阅SetDeadline。
func (f *File) SetReadDeadline(t time.Time) error { //md5:b7a6a80ebdae680d8c81f9e2d59da1dc

}

// SetWriteDeadline 设置任何未来 Write 调用以及当前阻塞的 Write 调用的截止时间。
// 即使 Write 超时，也可能返回 n > 0，表示已成功写入部分数据。
// 将 t 设置为零值表示 Write 不会超时。
// 并非所有文件都支持设置截止时间；请参阅 SetDeadline。
func (f *File) SetWriteDeadline(t time.Time) error { //md5:46a85eb5330d3d52baa224ef62a4e652

}

// SyscallConn 返回一个原始文件描述符。
// 这实现了 syscall.Conn 接口。
func (f *File) SyscallConn() (syscall.RawConn, error) { //md5:3f153cc7042b961a3e13236a4c8b8d56

}

// DirFS 返回一个文件系统（类型为 fs.FS），该文件系统表示以目录 dir 为根的文件树。
//
// 注意，DirFS("/prefix") 只保证其对操作系统的 Open 调用将以 "/prefix" 开头：即 DirFS("/prefix").Open("file") 等同于 os.Open("/prefix/file")。因此，如果 /prefix/file 是指向 /prefix 树外部的符号链接，那么使用 DirFS 并不会比直接使用 os.Open 更加限制对该链接的访问。此外，对于相对路径的根目录，如 DirFS("prefix")，其根目录会受到后续 Chdir 调用的影响。因此，当目录树包含任意内容时，DirFS 并不能作为 chroot 风格安全机制的一般替代。
//
// 目录 dir 必须非空字符串。
//
// 返回结果实现了 [io/fs.StatFS]、[io/fs.ReadFileFS] 和 [io/fs.ReadDirFS]。
func DirFS(dir string) fs.FS { //md5:401dc46c2098e233c38a59f8017e976f

}

// ReadFile 读取指定名称的文件并返回其内容。
// 成功调用将返回 err == nil，而不是 err == EOF。
// 因为 ReadFile 会读取整个文件，所以它不会将 Read 返回的 EOF 视为需要报告的错误。
func ReadFile(name string) ([]byte, error) { //md5:068b3628ea225053c9f0c0c18fb04e45

}

// WriteFile 将数据写入指定的文件，如果必要则创建它。如果文件不存在，WriteFile 会使用权限 perm（在 umask 之后）创建它；否则，在写入之前会截断文件，但不改变权限。由于 WriteFile 需要多个系统调用来完成，操作过程中出现失败可能会使文件处于部分写入状态。
func WriteFile(name string, data []byte, perm FileMode) error { //md5:fd0677a73de1810629354b92eb29db6d

}

// Fd 返回引用打开文件的 Plan 9 文件描述符。如果 f 关闭，文件描述符将变为无效。如果 f 被垃圾回收，清理器可能会关闭文件描述符，使其无效；有关何时运行清理器的更多信息，请参阅 runtime.SetFinalizer。在 Unix 系统上，这将导致 SetDeadline 方法停止工作。
// 
// 作为替代方案，请参阅 f.SyscallConn 方法。
func (f *File) Fd() uintptr { //md5:9d077621526e325ee13f9351694db734

}

// NewFile 函数以给定的文件描述符和名称返回一个新的 File。若 fd 不是有效的文件描述符，则返回值将为 nil。
func NewFile(fd uintptr, name string) *File { //md5:4240320d589243e180f39489a70ff3d0

}

// Close 关闭文件，使其无法进行输入/输出操作。
// 对于支持设置超时的文件，任何待处理的 I/O 操作将被取消并立即返回一个 ErrClosed 错误。
// 如果已经调用过 Close，则它将返回一个错误。
func (f *File) Close() error { //md5:7059ce72dff80153c2c63c5bb4a7b269

}

// Stat返回描述文件的FileInfo结构体。如果出现错误，错误类型为*PathError。
func (f *File) Stat() (FileInfo, error) { //md5:6b78d9d3f0557c38bca468468636ac30

}

// Truncate 修改文件的大小。
// 它不会改变输入/输出偏移量。
// 如果出现错误，错误类型为 *PathError。
func (f *File) Truncate(size int64) error { //md5:255d9cfb34f337b46a64e8cbb77cf8f8

}

// Sync 将文件当前内容提交到稳定的存储中。
// 通常情况下，这意味着将文件系统内存中缓存的最近写入数据刷入磁盘。
func (f *File) Sync() error { //md5:71a4dd5841a579655b1e39b6db926e96

}

// Truncate 函数改变指定文件的大小。如果该文件是符号链接，它将改变链接目标的大小。
// 如果发生错误，该错误将为 *PathError 类型。
func Truncate(name string, size int64) error { //md5:f3181406c058e4a0f1b5114b051d247d

}

// Remove 删除指定的文件或目录。
// 如果出现错误，错误类型为 *PathError。
func Remove(name string) error { //md5:a57bcc3207ee49e385bed3a7f20c070d

}

// Chtimes 类似于 Unix 的 utime() 或者 utimes() 函数，更改指定文件的访问和修改时间。
// 如果将 time.Time 值设置为零，则对应的文件时间将保持不变。
// 
// 在底层文件系统中，可能会将值截断或四舍五入到一个更不精确的时间单位。
// 如果出现错误，错误类型将是 *PathError。
func Chtimes(name string, atime time.Time, mtime time.Time) error { //md5:0c33b414b3debf6e949ed0b68e9f55b1

}

// Pipe 返回一对已连接的 Files；从 r 读取的数据即为写入 w 的字节。
// 它同时返回这两个文件以及可能存在的错误。
func Pipe() (r *File, w *File, err error) { //md5:62203dc59d8a2aaf2d9c0f2693b287c0

}

// Link 创建一个名为newname的硬链接，指向oldname文件。
// 如果出现错误，它将是一个*LinkError类型。
func Link(oldname, newname string) error { //md5:73080a75f9eec32cccd19d6cc084cf0f

}

// Symlink 创建新的符号链接newname，链接到oldname。
// 在Windows系统中，如果oldname不存在，创建的将是文件符号链接；如果oldname后来被创建为目录，该符号链接将无法使用。
// 如果出现错误，错误类型为*LinkError。
func Symlink(oldname, newname string) error { //md5:e9f4790d388077176068fdc473bfd82a

}

// Chown 更改指定文件的数值用户ID（uid）和组ID（gid）。如果该文件是符号链接，它会更改链接目标的uid和gid。如果指定的uid或gid为-1，则表示不改变该值。如果出现错误，错误类型将是*PathError。
// 
// 在 Windows 或 Plan 9 上，Chown 总是返回 syscall.EWINDOWS 或 EPLAN9 错误，这些错误会被*PathError包装。
func Chown(name string, uid, gid int) error { //md5:bd2789bd099c5520f599a4f0f4171d2d

}

// Lchown 修改指定文件的数字用户ID和组ID。
// 若该文件为符号链接，将修改链接自身的用户ID和组ID。
// 若出现错误，其类型将为 *PathError。
func Lchown(name string, uid, gid int) error { //md5:7471c2ee870f4e17889c237d2ad82b6a

}

// Chown 更改指定文件的数字用户ID（uid）和组ID（gid）。
// 如果出现错误，该错误将为 *PathError 类型。
func (f *File) Chown(uid, gid int) error { //md5:56c6f8a8d1cd555a1d9fc4f26a6802d2

}

// Chdir 将当前工作目录更改为文件，该文件必须是一个目录。
// 如果出现错误，错误类型为 *PathError。
func (f *File) Chdir() error { //md5:9c746c718f7b00e1b62d7c81d3e4b24f

}

// Fd 返回引用打开文件的 Windows 处理。如果 f 关闭，文件描述符将变为无效。如果 f 被垃圾回收，清理器可能会关闭文件描述符，使其无效；有关何时运行清理器的更多信息，请参阅 runtime.SetFinalizer。在 Unix 系统上，这将导致 SetDeadline 方法停止工作。
func (file *File) Fd() uintptr { //md5:fa02f63c941e821df392b8ab5567228e

}