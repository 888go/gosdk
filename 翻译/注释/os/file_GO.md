
<原文开始>
// Name returns the name of the file as presented to Open.
<原文结束>

# <翻译开始>
// Name 返回传递给 Open 的文件名。
# <翻译结束>


<原文开始>
// Read reads up to len(b) bytes from the File and stores them in b.
// It returns the number of bytes read and any error encountered.
// At end of file, Read returns 0, io.EOF.
<原文结束>

# <翻译开始>
// Read从File读取最多len(b)字节，并将它们存储在b中。它返回已读取的字节数和遇到的任何错误。在文件结束时，Read返回0和io.EOF。
# <翻译结束>


<原文开始>
// ReadAt reads len(b) bytes from the File starting at byte offset off.
// It returns the number of bytes read and the error, if any.
// ReadAt always returns a non-nil error when n < len(b).
// At end of file, that error is io.EOF.
<原文结束>

# <翻译开始>
// ReadAt 从File的字节偏移量off处开始读取len(b)个字节。
// 它返回读取的字节数以及任何错误。当n < len(b)时，ReadAt总会返回一个非nil的错误。
// 在文件末尾，该错误为io.EOF。
# <翻译结束>


<原文开始>
// ReadFrom implements io.ReaderFrom.
<原文结束>

# <翻译开始>
// ReadFrom 实现了 io.ReaderFrom 接口。
# <翻译结束>


<原文开始>
// ReadFrom hides another ReadFrom method.
// It should never be called.
<原文结束>

# <翻译开始>
// ReadFrom隐藏了另一个ReadFrom方法。它不应该被调用。
# <翻译结束>


<原文开始>
// Write writes len(b) bytes from b to the File.
// It returns the number of bytes written and an error, if any.
// Write returns a non-nil error when n != len(b).
<原文结束>

# <翻译开始>
// Write 从b中写入len(b)字节到File。它返回写入的字节数和任何错误。当n不等于len(b)时，Write返回非nil错误。
# <翻译结束>


<原文开始>
// WriteAt writes len(b) bytes to the File starting at byte offset off.
// It returns the number of bytes written and an error, if any.
// WriteAt returns a non-nil error when n != len(b).
//
// If file was opened with the O_APPEND flag, WriteAt returns an error.
<原文结束>

# <翻译开始>
// WriteAt 从字节偏移量 off 开始向 File 写入 len(b) 个字节。
// 它返回已写入的字节数以及任何可能遇到的错误。
// 当 n != len(b) 时，WriteAt 返回一个非 nil 错误。
//
// 若文件以 O_APPEND 标志打开，WriteAt 将返回一个错误。
# <翻译结束>


<原文开始>
// WriteTo implements io.WriterTo.
<原文结束>

# <翻译开始>
// WriteTo 实现了 io.WriterTo 接口。
# <翻译结束>


<原文开始>
// WriteTo hides another WriteTo method.
// It should never be called.
<原文结束>

# <翻译开始>
// WriteTo 隐藏了另一个 WriteTo 方法。它不应该被调用。
# <翻译结束>


<原文开始>
// Seek sets the offset for the next Read or Write on file to offset, interpreted
// according to whence: 0 means relative to the origin of the file, 1 means
// relative to the current offset, and 2 means relative to the end.
// It returns the new offset and an error, if any.
// The behavior of Seek on a file opened with O_APPEND is not specified.
<原文结束>

# <翻译开始>
// Seek 设置文件的下一个 Read 或 Write 操作的偏移量，根据 whence 的值进行解释：0 表示相对于文件的起始位置，1 表示相对于当前偏移量，2 表示相对于文件末尾。它返回新的偏移量和任何错误。对于使用 O_APPEND 模式打开的文件，Seek 的行为未做规定。
# <翻译结束>


<原文开始>
// WriteString is like Write, but writes the contents of string s rather than
// a slice of bytes.
<原文结束>

# <翻译开始>
// WriteString 类似于 Write，但其写入的是字符串 s 的内容，而非字节片。
# <翻译结束>


<原文开始>
// Mkdir creates a new directory with the specified name and permission
// bits (before umask).
// If there is an error, it will be of type *PathError.
<原文结束>

# <翻译开始>
// Mkdir 创建一个具有指定名称和权限位（在umask之前）的新目录。
// 如果发生错误，该错误将为 *PathError 类型。
# <翻译结束>


<原文开始>
// Chdir changes the current working directory to the named directory.
// If there is an error, it will be of type *PathError.
<原文结束>

# <翻译开始>
// Chdir 将当前工作目录更改为指定的目录。如果出现错误，错误类型为 *PathError。
# <翻译结束>


<原文开始>
// Open opens the named file for reading. If successful, methods on
// the returned file can be used for reading; the associated file
// descriptor has mode O_RDONLY.
// If there is an error, it will be of type *PathError.
<原文结束>

# <翻译开始>
// Open 打开指定的文件进行读取。如果成功，可以使用返回文件的方法进行读取；关联的文件描述符的模式为 O_RDONLY。
// 如果出现错误，错误类型为 *PathError。
# <翻译结束>


<原文开始>
// Create creates or truncates the named file. If the file already exists,
// it is truncated. If the file does not exist, it is created with mode 0666
// (before umask). If successful, methods on the returned File can
// be used for I/O; the associated file descriptor has mode O_RDWR.
// If there is an error, it will be of type *PathError.
<原文结束>

# <翻译开始>
// Create 用于创建或截断指定名称的文件。若该文件已存在，将会被截断；若不存在，则会以0666（应用umask之前）的模式创建。操作成功后，可使用返回的File对象进行I/O操作，其关联的文件描述符具有O_RDWR模式。如果出现错误，其类型将为*PathError。
# <翻译结束>


<原文开始>
// OpenFile is the generalized open call; most users will use Open
// or Create instead. It opens the named file with specified flag
// (O_RDONLY etc.). If the file does not exist, and the O_CREATE flag
// is passed, it is created with mode perm (before umask). If successful,
// methods on the returned File can be used for I/O.
// If there is an error, it will be of type *PathError.
<原文结束>

# <翻译开始>
// OpenFile 是通用的打开文件调用；大多数用户应使用 Open 或 Create。它以指定的标志（如 O_RDONLY 等）
// 打开名为 name 的文件。如果文件不存在，且传递了 O_CREATE 标志，那么会使用权限 perm（在 umask 之前）
// 创建该文件。如果成功，返回的 File 对象上的方法可用于 I/O 操作。
// 如果出现错误，该错误将为 *PathError 类型。
# <翻译结束>


<原文开始>
// Rename renames (moves) oldpath to newpath.
// If newpath already exists and is not a directory, Rename replaces it.
// OS-specific restrictions may apply when oldpath and newpath are in different directories.
// Even within the same directory, on non-Unix platforms Rename is not an atomic operation.
// If there is an error, it will be of type *LinkError.
<原文结束>

# <翻译开始>
// Rename 将 oldpath 重命名（移动）到 newpath。如果 newpath 已经存在且不是目录，Rename 将替换它。当 oldpath 和 newpath 在不同的目录时，可能存在特定于操作系统的限制。即使在同一个目录内，在非 Unix 平台上，Rename 也不是一个原子操作。如果出现错误，错误类型将为 *LinkError。
# <翻译结束>


<原文开始>
// Readlink returns the destination of the named symbolic link.
// If there is an error, it will be of type *PathError.
//
// If the link destination is relative, Readlink returns the relative path
// without resolving it to an absolute one.
<原文结束>

# <翻译开始>
// Readlink 返回指定符号链接的目标位置。
// 如果出现错误，错误类型为 *PathError。
// 
// 如果链接目标是相对路径，Readlink 会返回未解析为绝对路径的相对路径。
# <翻译结束>


<原文开始>
// TempDir returns the default directory to use for temporary files.
//
// On Unix systems, it returns $TMPDIR if non-empty, else /tmp.
// On Windows, it uses GetTempPath, returning the first non-empty
// value from %TMP%, %TEMP%, %USERPROFILE%, or the Windows directory.
// On Plan 9, it returns /tmp.
//
// The directory is neither guaranteed to exist nor have accessible
// permissions.
<原文结束>

# <翻译开始>
// TempDir 返回用于临时文件的默认目录。
//
// 在 Unix 系统中，如果 $TMPDIR 非空，则返回其值，否则返回 /tmp。
// 在 Windows 上，它使用 GetTempPath 函数，按顺序返回以下各项中首个非空值：
// %TMP%、%TEMP%、%USERPROFILE% 或 Windows 目录。
// 在 Plan 9 系统上，返回 /tmp。
//
// 此目录并不保证一定存在或具有可访问权限。
# <翻译结束>


<原文开始>
// UserCacheDir returns the default root directory to use for user-specific
// cached data. Users should create their own application-specific subdirectory
// within this one and use that.
//
// On Unix systems, it returns $XDG_CACHE_HOME as specified by
// https://specifications.freedesktop.org/basedir-spec/basedir-spec-latest.html if
// non-empty, else $HOME/.cache.
// On Darwin, it returns $HOME/Library/Caches.
// On Windows, it returns %LocalAppData%.
// On Plan 9, it returns $home/lib/cache.
//
// If the location cannot be determined (for example, $HOME is not defined),
// then it will return an error.
<原文结束>

# <翻译开始>
// UserCacheDir 返回用于用户特定缓存数据的默认根目录。用户应在该目录下创建他们自己的应用程序特定子目录并使用它。
//
// 在Unix系统上，如果$XDG_CACHE_HOME非空（如https://specifications.freedesktop.org/basedir-spec/basedir-spec-latest.html中指定），则返回$XDG_CACHE_HOME，否则返回$HOME/.cache。
// 在Darwin系统上，返回$HOME/Library/Caches。
// 在Windows系统上，返回%LocalAppData%。
// 在Plan 9系统上，返回$home/lib/cache。
//
// 如果无法确定位置（例如，$HOME未定义），则会返回错误。
# <翻译结束>


<原文开始>
// UserConfigDir returns the default root directory to use for user-specific
// configuration data. Users should create their own application-specific
// subdirectory within this one and use that.
//
// On Unix systems, it returns $XDG_CONFIG_HOME as specified by
// https://specifications.freedesktop.org/basedir-spec/basedir-spec-latest.html if
// non-empty, else $HOME/.config.
// On Darwin, it returns $HOME/Library/Application Support.
// On Windows, it returns %AppData%.
// On Plan 9, it returns $home/lib.
//
// If the location cannot be determined (for example, $HOME is not defined),
// then it will return an error.
<原文结束>

# <翻译开始>
// UserConfigDir 返回默认的用户特定配置数据根目录。用户应在该目录下创建自己应用的子目录，并使用这个子目录。
// 
// 在 Unix 系统中，它会返回 $XDG_CONFIG_HOME，如 https://specifications.freedesktop.org/basedir-spec/basedir-spec-latest.html 中指定的，如果非空，则否则返回 $HOME/.config。
// 在 Darwin（macOS）上，它返回 $HOME/Library/Application Support。
// 在 Windows 上，它返回 %AppData%。
// 在 Plan 9 系统中，它返回 $home/lib。
// 
// 如果无法确定位置（例如，$HOME 未定义），则会返回错误。
# <翻译结束>


<原文开始>
// UserHomeDir returns the current user's home directory.
//
// On Unix, including macOS, it returns the $HOME environment variable.
// On Windows, it returns %USERPROFILE%.
// On Plan 9, it returns the $home environment variable.
//
// If the expected variable is not set in the environment, UserHomeDir
// returns either a platform-specific default value or a non-nil error.
<原文结束>

# <翻译开始>
// UserHomeDir 返回当前用户的主目录。
//
// 在 Unix（包括 macOS）中，它返回 $HOME 环境变量。在 Windows 中，它返回 %USERPROFILE%。
// 在 Plan 9 中，它返回 $home 环境变量。
//
// 如果预期的环境变量未设置，UserHomeDir 将返回一个平台特定的默认值或非空错误。
# <翻译结束>


<原文开始>
// Chmod changes the mode of the named file to mode.
// If the file is a symbolic link, it changes the mode of the link's target.
// If there is an error, it will be of type *PathError.
//
// A different subset of the mode bits are used, depending on the
// operating system.
//
// On Unix, the mode's permission bits, ModeSetuid, ModeSetgid, and
// ModeSticky are used.
//
// On Windows, only the 0200 bit (owner writable) of mode is used; it
// controls whether the file's read-only attribute is set or cleared.
// The other bits are currently unused. For compatibility with Go 1.12
// and earlier, use a non-zero mode. Use mode 0400 for a read-only
// file and 0600 for a readable+writable file.
//
// On Plan 9, the mode's permission bits, ModeAppend, ModeExclusive,
// and ModeTemporary are used.
<原文结束>

# <翻译开始>
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
# <翻译结束>


<原文开始>
// Chmod changes the mode of the file to mode.
// If there is an error, it will be of type *PathError.
<原文结束>

# <翻译开始>
// Chmod将文件的模式更改为mode。
// 如果出现错误，它将是一个类型为*PathError的错误。
# <翻译结束>


<原文开始>
// SetDeadline sets the read and write deadlines for a File.
// It is equivalent to calling both SetReadDeadline and SetWriteDeadline.
//
// Only some kinds of files support setting a deadline. Calls to SetDeadline
// for files that do not support deadlines will return ErrNoDeadline.
// On most systems ordinary files do not support deadlines, but pipes do.
//
// A deadline is an absolute time after which I/O operations fail with an
// error instead of blocking. The deadline applies to all future and pending
// I/O, not just the immediately following call to Read or Write.
// After a deadline has been exceeded, the connection can be refreshed
// by setting a deadline in the future.
//
// If the deadline is exceeded a call to Read or Write or to other I/O
// methods will return an error that wraps ErrDeadlineExceeded.
// This can be tested using errors.Is(err, os.ErrDeadlineExceeded).
// That error implements the Timeout method, and calling the Timeout
// method will return true, but there are other possible errors for which
// the Timeout will return true even if the deadline has not been exceeded.
//
// An idle timeout can be implemented by repeatedly extending
// the deadline after successful Read or Write calls.
//
// A zero value for t means I/O operations will not time out.
<原文结束>

# <翻译开始>
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
# <翻译结束>


<原文开始>
// SetReadDeadline sets the deadline for future Read calls and any
// currently-blocked Read call.
// A zero value for t means Read will not time out.
// Not all files support setting deadlines; see SetDeadline.
<原文结束>

# <翻译开始>
// SetReadDeadline 设置了未来Read调用和当前阻塞的Read调用的超时时间。
// 如果t的值为零，表示Read不会超时。
// 并非所有文件都支持设置超时；请参阅SetDeadline。
# <翻译结束>


<原文开始>
// SetWriteDeadline sets the deadline for any future Write calls and any
// currently-blocked Write call.
// Even if Write times out, it may return n > 0, indicating that
// some of the data was successfully written.
// A zero value for t means Write will not time out.
// Not all files support setting deadlines; see SetDeadline.
<原文结束>

# <翻译开始>
// SetWriteDeadline 设置任何未来 Write 调用以及当前阻塞的 Write 调用的截止时间。
// 即使 Write 超时，也可能返回 n > 0，表示已成功写入部分数据。
// 将 t 设置为零值表示 Write 不会超时。
// 并非所有文件都支持设置截止时间；请参阅 SetDeadline。
# <翻译结束>


<原文开始>
// SyscallConn returns a raw file.
// This implements the syscall.Conn interface.
<原文结束>

# <翻译开始>
// SyscallConn 返回一个原始文件描述符。
// 这实现了 syscall.Conn 接口。
# <翻译结束>


<原文开始>
// DirFS returns a file system (an fs.FS) for the tree of files rooted at the directory dir.
//
// Note that DirFS("/prefix") only guarantees that the Open calls it makes to the
// operating system will begin with "/prefix": DirFS("/prefix").Open("file") is the
// same as os.Open("/prefix/file"). So if /prefix/file is a symbolic link pointing outside
// the /prefix tree, then using DirFS does not stop the access any more than using
// os.Open does. Additionally, the root of the fs.FS returned for a relative path,
// DirFS("prefix"), will be affected by later calls to Chdir. DirFS is therefore not
// a general substitute for a chroot-style security mechanism when the directory tree
// contains arbitrary content.
//
// The directory dir must not be "".
//
// The result implements [io/fs.StatFS], [io/fs.ReadFileFS] and
// [io/fs.ReadDirFS].
<原文结束>

# <翻译开始>
// DirFS 返回一个文件系统（类型为 fs.FS），该文件系统表示以目录 dir 为根的文件树。
//
// 注意，DirFS("/prefix") 只保证其对操作系统的 Open 调用将以 "/prefix" 开头：即 DirFS("/prefix").Open("file") 等同于 os.Open("/prefix/file")。因此，如果 /prefix/file 是指向 /prefix 树外部的符号链接，那么使用 DirFS 并不会比直接使用 os.Open 更加限制对该链接的访问。此外，对于相对路径的根目录，如 DirFS("prefix")，其根目录会受到后续 Chdir 调用的影响。因此，当目录树包含任意内容时，DirFS 并不能作为 chroot 风格安全机制的一般替代。
//
// 目录 dir 必须非空字符串。
//
// 返回结果实现了 [io/fs.StatFS]、[io/fs.ReadFileFS] 和 [io/fs.ReadDirFS]。
# <翻译结束>


<原文开始>
// ReadFile reads the named file and returns the contents.
// A successful call returns err == nil, not err == EOF.
// Because ReadFile reads the whole file, it does not treat an EOF from Read
// as an error to be reported.
<原文结束>

# <翻译开始>
// ReadFile 读取指定名称的文件并返回其内容。
// 成功调用将返回 err == nil，而不是 err == EOF。
// 因为 ReadFile 会读取整个文件，所以它不会将 Read 返回的 EOF 视为需要报告的错误。
# <翻译结束>


<原文开始>
// WriteFile writes data to the named file, creating it if necessary.
// If the file does not exist, WriteFile creates it with permissions perm (before umask);
// otherwise WriteFile truncates it before writing, without changing permissions.
// Since WriteFile requires multiple system calls to complete, a failure mid-operation
// can leave the file in a partially written state.
<原文结束>

# <翻译开始>
// WriteFile 将数据写入指定的文件，如果必要则创建它。如果文件不存在，WriteFile 会使用权限 perm（在 umask 之后）创建它；否则，在写入之前会截断文件，但不改变权限。由于 WriteFile 需要多个系统调用来完成，操作过程中出现失败可能会使文件处于部分写入状态。
# <翻译结束>


<原文开始>
// Fd returns the integer Plan 9 file descriptor referencing the open file.
// If f is closed, the file descriptor becomes invalid.
// If f is garbage collected, a finalizer may close the file descriptor,
// making it invalid; see runtime.SetFinalizer for more information on when
// a finalizer might be run. On Unix systems this will cause the SetDeadline
// methods to stop working.
//
// As an alternative, see the f.SyscallConn method.
<原文结束>

# <翻译开始>
// Fd 返回引用打开文件的 Plan 9 文件描述符。如果 f 关闭，文件描述符将变为无效。如果 f 被垃圾回收，清理器可能会关闭文件描述符，使其无效；有关何时运行清理器的更多信息，请参阅 runtime.SetFinalizer。在 Unix 系统上，这将导致 SetDeadline 方法停止工作。
// 
// 作为替代方案，请参阅 f.SyscallConn 方法。
# <翻译结束>


<原文开始>
// NewFile returns a new File with the given file descriptor and
// name. The returned value will be nil if fd is not a valid file
// descriptor.
<原文结束>

# <翻译开始>
// NewFile 函数以给定的文件描述符和名称返回一个新的 File。若 fd 不是有效的文件描述符，则返回值将为 nil。
# <翻译结束>


<原文开始>
// Close closes the File, rendering it unusable for I/O.
// On files that support SetDeadline, any pending I/O operations will
// be canceled and return immediately with an ErrClosed error.
// Close will return an error if it has already been called.
<原文结束>

# <翻译开始>
// Close 关闭文件，使其无法进行输入/输出操作。
// 对于支持设置超时的文件，任何待处理的 I/O 操作将被取消并立即返回一个 ErrClosed 错误。
// 如果已经调用过 Close，则它将返回一个错误。
# <翻译结束>


<原文开始>
// Stat returns the FileInfo structure describing file.
// If there is an error, it will be of type *PathError.
<原文结束>

# <翻译开始>
// Stat返回描述文件的FileInfo结构体。如果出现错误，错误类型为*PathError。
# <翻译结束>


<原文开始>
// Truncate changes the size of the file.
// It does not change the I/O offset.
// If there is an error, it will be of type *PathError.
<原文结束>

# <翻译开始>
// Truncate 修改文件的大小。
// 它不会改变输入/输出偏移量。
// 如果出现错误，错误类型为 *PathError。
# <翻译结束>


<原文开始>
// Sync commits the current contents of the file to stable storage.
// Typically, this means flushing the file system's in-memory copy
// of recently written data to disk.
<原文结束>

# <翻译开始>
// Sync 将文件当前内容提交到稳定的存储中。
// 通常情况下，这意味着将文件系统内存中缓存的最近写入数据刷入磁盘。
# <翻译结束>


<原文开始>
// Truncate changes the size of the named file.
// If the file is a symbolic link, it changes the size of the link's target.
// If there is an error, it will be of type *PathError.
<原文结束>

# <翻译开始>
// Truncate 函数改变指定文件的大小。如果该文件是符号链接，它将改变链接目标的大小。
// 如果发生错误，该错误将为 *PathError 类型。
# <翻译结束>


<原文开始>
// Remove removes the named file or directory.
// If there is an error, it will be of type *PathError.
<原文结束>

# <翻译开始>
// Remove 删除指定的文件或目录。
// 如果出现错误，错误类型为 *PathError。
# <翻译结束>


<原文开始>
// Chtimes changes the access and modification times of the named
// file, similar to the Unix utime() or utimes() functions.
// A zero time.Time value will leave the corresponding file time unchanged.
//
// The underlying filesystem may truncate or round the values to a
// less precise time unit.
// If there is an error, it will be of type *PathError.
<原文结束>

# <翻译开始>
// Chtimes 类似于 Unix 的 utime() 或者 utimes() 函数，更改指定文件的访问和修改时间。
// 如果将 time.Time 值设置为零，则对应的文件时间将保持不变。
// 
// 在底层文件系统中，可能会将值截断或四舍五入到一个更不精确的时间单位。
// 如果出现错误，错误类型将是 *PathError。
# <翻译结束>


<原文开始>
// Pipe returns a connected pair of Files; reads from r return bytes
// written to w. It returns the files and an error, if any.
<原文结束>

# <翻译开始>
// Pipe 返回一对已连接的 Files；从 r 读取的数据即为写入 w 的字节。
// 它同时返回这两个文件以及可能存在的错误。
# <翻译结束>


<原文开始>
// Link creates newname as a hard link to the oldname file.
// If there is an error, it will be of type *LinkError.
<原文结束>

# <翻译开始>
// Link 创建一个名为newname的硬链接，指向oldname文件。
// 如果出现错误，它将是一个*LinkError类型。
# <翻译结束>


<原文开始>
// Symlink creates newname as a symbolic link to oldname.
// On Windows, a symlink to a non-existent oldname creates a file symlink;
// if oldname is later created as a directory the symlink will not work.
// If there is an error, it will be of type *LinkError.
<原文结束>

# <翻译开始>
// Symlink 创建新的符号链接newname，链接到oldname。
// 在Windows系统中，如果oldname不存在，创建的将是文件符号链接；如果oldname后来被创建为目录，该符号链接将无法使用。
// 如果出现错误，错误类型为*LinkError。
# <翻译结束>


<原文开始>
// Chown changes the numeric uid and gid of the named file.
// If the file is a symbolic link, it changes the uid and gid of the link's target.
// A uid or gid of -1 means to not change that value.
// If there is an error, it will be of type *PathError.
//
// On Windows or Plan 9, Chown always returns the syscall.EWINDOWS or
// EPLAN9 error, wrapped in *PathError.
<原文结束>

# <翻译开始>
// Chown 更改指定文件的数值用户ID（uid）和组ID（gid）。如果该文件是符号链接，它会更改链接目标的uid和gid。如果指定的uid或gid为-1，则表示不改变该值。如果出现错误，错误类型将是*PathError。
// 
// 在 Windows 或 Plan 9 上，Chown 总是返回 syscall.EWINDOWS 或 EPLAN9 错误，这些错误会被*PathError包装。
# <翻译结束>


<原文开始>
// Lchown changes the numeric uid and gid of the named file.
// If the file is a symbolic link, it changes the uid and gid of the link itself.
// If there is an error, it will be of type *PathError.
<原文结束>

# <翻译开始>
// Lchown 修改指定文件的数字用户ID和组ID。
// 若该文件为符号链接，将修改链接自身的用户ID和组ID。
// 若出现错误，其类型将为 *PathError。
# <翻译结束>


<原文开始>
// Chown changes the numeric uid and gid of the named file.
// If there is an error, it will be of type *PathError.
<原文结束>

# <翻译开始>
// Chown 更改指定文件的数字用户ID（uid）和组ID（gid）。
// 如果出现错误，该错误将为 *PathError 类型。
# <翻译结束>


<原文开始>
// Chdir changes the current working directory to the file,
// which must be a directory.
// If there is an error, it will be of type *PathError.
<原文结束>

# <翻译开始>
// Chdir 将当前工作目录更改为文件，该文件必须是一个目录。
// 如果出现错误，错误类型为 *PathError。
# <翻译结束>


<原文开始>
// Fd returns the Windows handle referencing the open file.
// If f is closed, the file descriptor becomes invalid.
// If f is garbage collected, a finalizer may close the file descriptor,
// making it invalid; see runtime.SetFinalizer for more information on when
// a finalizer might be run. On Unix systems this will cause the SetDeadline
// methods to stop working.
<原文结束>

# <翻译开始>
// Fd 返回引用打开文件的 Windows 处理。如果 f 关闭，文件描述符将变为无效。如果 f 被垃圾回收，清理器可能会关闭文件描述符，使其无效；有关何时运行清理器的更多信息，请参阅 runtime.SetFinalizer。在 Unix 系统上，这将导致 SetDeadline 方法停止工作。
# <翻译结束>


<原文开始>
// LinkError records an error during a link or symlink or rename
// system call and the paths that caused it.
<原文结束>

# <翻译开始>
// LinkError 记录链接、符号链接或重命名系统调用期间发生的错误，以及导致错误的路径。
# <翻译结束>

