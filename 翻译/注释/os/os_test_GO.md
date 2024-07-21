
<原文开始>
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// 版权所有 ? 2009 Go 作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 此许可证可在 LICENSE 文件中找到。
# <翻译结束>


<原文开始>
		// In a self-hosted iOS build the above files might
		// not exist. Look for system files instead below.
<原文结束>

# <翻译开始>
		// 在自托管的iOS构建中，上述文件可能不存在。请在下面查找系统文件。
# <翻译结束>


<原文开始>
		// wasmtime has issues resolving symbolic links that are often present
		// in directories like /etc/group below (e.g. private/etc/group on OSX).
		// For this reason we use files in the Go source tree instead.
<原文结束>

# <翻译开始>
		// wasmtime在解析像`/etc/group`这样的目录中经常存在的符号链接时存在问题（例如，在OSX上的`private/etc/group`）。
		// 因此，我们使用Go源代码树中的文件代替。
# <翻译结束>


<原文开始>
// localTmp returns a local temporary directory not on NFS.
<原文结束>

# <翻译开始>
// localTmp 返回一个不在NFS上的本地临时目录。
# <翻译结束>


<原文开始>
// Read with length 0 should not return EOF.
<原文结束>

# <翻译开始>
// 长度为0的读取操作不应返回EOF。
# <翻译结束>


<原文开始>
// Reading a closed file should return ErrClosed error
<原文结束>

# <翻译开始>
// 读取一个已关闭的文件应该返回ErrClosed错误
# <翻译结束>


<原文开始>
// Read the directory one entry at a time.
<原文结束>

# <翻译开始>
// 一次读取目录中的一个条目。
# <翻译结束>


<原文开始>
// Check that reading a directory one entry at a time gives the same result
// as reading it all at once.
<原文结束>

# <翻译开始>
// 检查一次读取目录一个条目是否与一次性读取整个目录的结果相同。
# <翻译结束>


<原文开始>
// big directory that doesn't change often.
<原文结束>

# <翻译开始>
// 大型目录，不常变动
# <翻译结束>


<原文开始>
// +100 in case we screw up
<原文结束>

# <翻译开始>
// +100 以防我们出错
# <翻译结束>


<原文开始>
// and tests buffer >100 case
<原文结束>

# <翻译开始>
// 并且测试缓冲区大小大于100的情况
# <翻译结束>


<原文开始>
		// Windows and Plan 9 already do this correctly,
		// but are structured with different syscalls such
		// that they don't use Lstat, so the hook below for
		// testing it wouldn't work.
<原文结束>

# <翻译开始>
		// Windows和Plan 9已经正确地完成了这个操作，
		// 但是它们的系统调用结构不同，
		// 所以它们不使用Lstat，因此下面用于测试的钩子将无法工作。
# <翻译结束>


<原文开始>
// will disappear or have an error
<原文结束>

# <翻译开始>
// 将消失或出现错误
# <翻译结束>


<原文开始>
// Readdir on a regular file should fail.
<原文结束>

# <翻译开始>
// 对普通文件执行Readdir应该失败。
# <翻译结束>


<原文开始>
// Check the returned error is well-formed.
<原文结束>

# <翻译开始>
// 检查返回的错误是否有效。
# <翻译结束>


<原文开始>
// We should not be able to perform the same Link() a second time
<原文结束>

# <翻译开始>
// 我们应该无法再次执行Link()操作
# <翻译结束>


<原文开始>
// chtmpdir changes the working directory to a new temporary directory and
// provides a cleanup function.
<原文结束>

# <翻译开始>
// chtmpdir 将工作目录更改为新的临时目录并提供一个清理函数。
# <翻译结束>


<原文开始>
// Long, but not too long: a common limit is 255.
<原文结束>

# <翻译开始>
// 长度适中，但不要过长：常见的限制是255个字符。
# <翻译结束>


<原文开始>
// Sanity check that the underlying filesystem is not case sensitive.
<原文结束>

# <翻译开始>
// 对底层文件系统是否区分大小写进行合理性检查。
# <翻译结束>


<原文开始>
			// Stat does not return the real case of the file (it returns what the called asked for)
			// So we have to use readdir to get the real name of the file.
<原文结束>

# <翻译开始>
			// Stat方法不会返回文件的实际大小（它返回的是调用者所请求的内容），所以我们必须使用readdir来获取文件的真正名称。
# <翻译结束>


<原文开始>
// Chmod is not supported on wasip1.
<原文结束>

# <翻译开始>
// Chmod 在 wasip1 上不受支持。
# <翻译结束>


<原文开始>
// Creation mode is read write
<原文结束>

# <翻译开始>
// 创建模式为读写
# <翻译结束>


<原文开始>
// wrote at offset past where hello, world was.
<原文结束>

# <翻译开始>
// 在“hello, world”之后的偏移处写入。
# <翻译结束>


<原文开始>
// Truncate shouldn't create any new file.
<原文结束>

# <翻译开始>
// Truncate不应该创建任何新的文件。
# <翻译结束>


<原文开始>
// Use TempDir (via newFile) to make sure we're on a local file system,
// so that timings are not distorted by latency and caching.
// On NFS, timings can be off due to caching of meta-data on
// NFS servers (Issue 848).
<原文结束>

# <翻译开始>
// 通过使用 TempDir（经由 newFile）确保我们位于本地文件系统上，
// 这样时间测量就不会因延迟和缓存而失真。
// 在 NFS（网络文件系统）上，由于 NFS 服务器上元数据的缓存问题（Issue 848），时间测量可能会出现偏差。
# <翻译结束>


<原文开始>
// Now change the times accordingly.
<原文结束>

# <翻译开始>
// 现在根据需要修改时间。
# <翻译结束>


<原文开始>
// Finally verify the expectations.
<原文结束>

# <翻译开始>
// 最后验证预期。
# <翻译结束>


<原文开始>
				// Mtime is the time of the last change of
				// content.  Similarly, atime is set whenever
				// the contents are accessed; also, it is set
				// whenever mtime is set.
<原文结束>

# <翻译开始>
				// Mtime是上次内容更改的时间。类似地，每当内容被访问时，atime也会被设置；同时，当mtime被设置时，它也会被设置。
# <翻译结束>


<原文开始>
							// On a 64-bit implementation, birth time is generally supported and cannot be changed.
							// When supported, atime update is restricted and depends on the file system and on the
							// OS configuration.
<原文结束>

# <翻译开始>
							// 在64位实现中，通常支持并不可更改出生时间（birth time）。
							// 若得到支持，访问时间（atime）的更新则受到限制，具体取决于文件系统以及操作系统配置。
# <翻译结束>


<原文开始>
// Use TempDir (via newDir) to make sure we're on a local file system,
// so that timings are not distorted by latency and caching.
// On NFS, timings can be off due to caching of meta-data on
// NFS servers (Issue 848).
<原文结束>

# <翻译开始>
// 使用TempDir（通过newDir）确保我们在本地文件系统上，
// 这样时间测量就不会因延迟和缓存而失真。
// 在NFS上，由于NFS服务器对元数据的缓存，时间可能会有偏差（问题848）。
# <翻译结束>


<原文开始>
// Move access and modification time back a second
<原文结束>

# <翻译开始>
// 将访问和修改时间向后回退一秒
# <翻译结束>


<原文开始>
	// These are chosen carefully not to be symlinks on a Mac
	// (unlike, say, /var, /etc), except /tmp, which we handle below.
<原文结束>

# <翻译开始>
	// 这些路径都是精心选择的，不会在Mac上成为符号链接（不像/var、/etc），除了/tmp，我们稍后会处理。
# <翻译结束>


<原文开始>
// /usr/bin does not usually exist on Plan 9 or Android.
<原文结束>

# <翻译开始>
// 在 Plan 9 或 Android 系统中，/usr/bin 通常不存在。
# <翻译结束>


<原文开始>
// Expand symlinks so path equality tests work.
<原文结束>

# <翻译开始>
// 展开符号链接，以便进行路径相等性测试。
# <翻译结束>


<原文开始>
				// We changed the current directory and cannot go back.
				// Don't let the tests continue; they'll scribble
				// all over some other directory.
<原文结束>

# <翻译开始>
				// 我们已经改变了当前目录，无法返回。
				// 不要让测试继续进行；它们会乱写
				// 到其他目录中。
# <翻译结束>


<原文开始>
// Test that Chdir+Getwd is program-wide.
<原文结束>

# <翻译开始>
// 测试Chdir+Getwd在整个程序中的功能。
# <翻译结束>


<原文开始>
			// It's not safe to continue with tests if we can't get back to
			// the original working directory.
<原文结束>

# <翻译开始>
			// 如果我们无法返回到原始工作目录，那么继续进行测试是不安全的。
# <翻译结束>


<原文开始>
	// Note the deferred Wait must be called after the deferred close(done),
	// to ensure the N goroutines have been released even if the main goroutine
	// calls Fatalf. It must be called before the Chdir back to the original
	// directory, and before the deferred deletion implied by TempDir,
	// so as not to interfere while the N goroutines are still running.
<原文结束>

# <翻译开始>
	// 注意，延迟调用的Wait必须在延迟关闭(done)之后调用，
	// 以确保即使主goroutine调用了Fatalf，N个goroutine也已经被释放。它必须在
	// 恢复到原始目录之前以及TempDir隐含的延迟删除之前调用，
	// 以免在N个goroutine仍在运行时产生干扰。
# <翻译结束>


<原文开始>
			// Lock half the goroutines in their own operating system
			// thread to exercise more scheduler possibilities.
<原文结束>

# <翻译开始>
			// 将一半的goroutine锁定在自己的操作系统线程中，以锻炼更多的调度可能性。
# <翻译结束>


<原文开始>
				// On Plan 9, after calling LockOSThread, the goroutines
				// run on different processes which don't share the working
				// directory. This used to be an issue because Go expects
				// the working directory to be program-wide.
				// See issue 9428.
<原文结束>

# <翻译开始>
				// 在Plan 9上，调用LockOSThread后，goroutine将在不同的进程中运行，这些进程不共享工作目录。这曾经是一个问题，因为Go期望工作目录是全局的。
				// 参见问题9428。
# <翻译结束>


<原文开始>
	// OS X sets TMPDIR to a symbolic link.
	// So we resolve our working directory again before the test.
<原文结束>

# <翻译开始>
	// 在OS X系统中，TMPDIR设置为一个符号链接。
	// 因此我们在测试之前再次解析我们的工作目录。
# <翻译结束>


<原文开始>
// Issue 21681, Windows 4G-1, etc:
<原文结束>

# <翻译开始>
// 问题21681，Windows 4G-1等：
# <翻译结束>


<原文开始>
// Reiserfs rejects the big seeks.
<原文结束>

# <翻译开始>
// Reiserfs拒绝大的 seek 操作。
# <翻译结束>


<原文开始>
					// Some Plan 9 file servers incorrectly return
					// EACCES rather than EISDIR when a directory is
					// opened for write.
<原文结束>

# <翻译开始>
					// 一些 Plan 9 文件服务器在打开目录进行写操作时，错误地返回 EACCES 而不是 EISDIR。
# <翻译结束>


<原文开始>
				// DragonFly incorrectly returns EACCES rather
				// EISDIR when a directory is opened for write.
<原文结束>

# <翻译开始>
				// DragonFly 系统在尝试以写入方式打开目录时，错误地返回 EACCES 而非 EISDIR。
# <翻译结束>


<原文开始>
// Run /bin/hostname and collect output.
<原文结束>

# <翻译开始>
// 运行/bin/hostname命令并收集输出。
# <翻译结束>


<原文开始>
	// There is no other way to fetch hostname on windows, but via winapi.
	// On Plan 9 it can be taken from #c/sysname as Hostname() does.
<原文结束>

# <翻译开始>
	// 在 Windows 上，没有其他方法可以获取主机名，但可以通过 WinAPI 实现。
	// 在 Plan 9 中，可以从 #c/sysname 获取，就像 Hostname() 函数所做的那样。
# <翻译结束>


<原文开始>
// No /bin/hostname to verify against.
<原文结束>

# <翻译开始>
// 没有 /bin/hostname 可以进行校验。
# <翻译结束>


<原文开始>
	// Check internal Hostname() against the output of /bin/hostname.
	// Allow that the internal Hostname returns a Fully Qualified Domain Name
	// and the /bin/hostname only returns the first component
<原文结束>

# <翻译开始>
	// 检查内部 Hostname() 函数的输出与 /bin/hostname 命令的输出是否一致。
	// 允许内部 Hostname 返回一个完整的域名（Fully Qualified Domain Name，FQDN），
	// 而 /bin/hostname 只返回第一个名称组件。
# <翻译结束>


<原文开始>
// Verify that ReadAt doesn't affect seek offset.
// In the Plan 9 kernel, there used to be a bug in the implementation of
// the pread syscall, where the channel offset was erroneously updated after
// calling pread on a file.
<原文结束>

# <翻译开始>
// 验证ReadAt方法不会影响 Seek 偏移量。
// 在Plan 9的内核中，pread系统调用的实现曾经存在一个错误，
// 在对文件调用pread后，会错误地更新通道偏移量。
# <翻译结束>


<原文开始>
// Verify that ReadAt doesn't allow negative offset.
<原文结束>

# <翻译开始>
// 验证ReadAt是否不允许负偏移量。
# <翻译结束>


<原文开始>
// Verify that WriteAt doesn't allow negative offset.
<原文结束>

# <翻译开始>
// 验证WriteAt是否不允许负偏移量。
# <翻译结束>


<原文开始>
// Verify that WriteAt doesn't work in append mode.
<原文结束>

# <翻译开始>
// 验证在追加模式下WriteAt方法不能正常工作
# <翻译结束>


<原文开始>
// Create new temporary directory and arrange to clean it up.
<原文结束>

# <翻译开始>
// 创建一个新的临时目录，并安排在完成后进行清理。
# <翻译结束>


<原文开始>
// Stat of path should succeed.
<原文结束>

# <翻译开始>
// 路径的状态应该成功。
# <翻译结束>


<原文开始>
// Stat of path+"/" should succeed too.
<原文结束>

# <翻译开始>
// 路径("/")的Stat也应该成功。
# <翻译结束>


<原文开始>
// This will make standard input a pipe.
<原文结束>

# <翻译开始>
// 这将会使标准输入变为管道。
# <翻译结束>


<原文开始>
// result will be like "prw-rw-rw"
<原文结束>

# <翻译开始>
// result 将类似于 "prw-rw-rw"
# <翻译结束>


<原文开始>
// Test the boundary of 247 and fewer bytes (normal) and 248 and more bytes (adjusted).
<原文结束>

# <翻译开始>
// 测试正常情况下（247字节及更少）和调整后情况下（248字节及以上）的边界。
# <翻译结束>


<原文开始>
// Ensure it does not end with a slash.
<原文结束>

# <翻译开始>
// 确保它不以斜杠结束。
# <翻译结束>


<原文开始>
			// The various sized runs are for this call to trigger the boundary
			// condition.
<原文结束>

# <翻译开始>
			// 各种大小的运行是为了触发此调用的边界条件。
# <翻译结束>


<原文开始>
// Re-exec the test binary to start a process that hangs until stdin is closed.
<原文结束>

# <翻译开始>
// 重新执行测试二进制文件，以启动一个进程，该进程将挂起，直到标准输入被关闭。
# <翻译结束>


<原文开始>
// Keep stdin alive until the process has finished dying.
<原文结束>

# <翻译开始>
// 保持标准输入流畅通，直到该进程完全退出。
# <翻译结束>


<原文开始>
	// Wait for the process to be started.
	// (It will close its stdout when it reaches TestMain.)
<原文结束>

# <翻译开始>
	// 等待进程启动。
	// （当它达到TestMain时，它将关闭其stdout。）
# <翻译结束>


<原文开始>
// TODO: golang.org/issue/8206
<原文结束>

# <翻译开始>
// 待办事项：golang.org/issue/8206
# <翻译结束>


<原文开始>
// verify that Getppid() from the forked process reports our process id
<原文结束>

# <翻译开始>
// 验证从forked进程调用的Getppid()是否报告我们的进程ID
# <翻译结束>


<原文开始>
// Test that all File methods give ErrInvalid if the receiver is nil.
<原文结束>

# <翻译开始>
// 测试如果接收者为nil，所有File方法都会返回ErrInvalid。
# <翻译结束>


<原文开始>
// Test that simultaneous RemoveAll do not report an error.
// As long as it gets removed, we should be happy.
<原文结束>

# <翻译开始>
// 测试同时移除所有是否报告错误。
// 只要它被移除了，我们就应该感到满意。
# <翻译结束>


<原文开始>
		// Windows has very strict rules about things like
		// removing directories while someone else has
		// them open. The racing doesn't work out nicely
		// like it does on Unix.
<原文结束>

# <翻译开始>
		// Windows 对诸如在其他程序打开目录时删除目录等操作有着非常严格的规定。这种竞态条件在 Unix 系统中可以很好地处理，但在 Windows 上并不理想。
# <翻译结束>


<原文开始>
// let workers race to remove root
<原文结束>

# <翻译开始>
// 让工作者竞争移除根节点
# <翻译结束>


<原文开始>
// Test that reading from a pipe doesn't use up a thread.
<原文结束>

# <翻译开始>
// 测试从管道读取不会占用线程。
# <翻译结束>


<原文开始>
// OpenBSD has a low default for max number of files.
<原文结束>

# <翻译开始>
// OpenBSD 对最大文件数量的默认值设置较低。
# <翻译结束>


<原文开始>
	// If we are still alive, it means that the 100 goroutines did
	// not require 100 threads.
<原文结束>

# <翻译开始>
	// 如果我们仍然存活，那就意味着这100个goroutine并未要求使用100个线程。
# <翻译结束>


<原文开始>
		// UserHomeDir may return a non-nil error if the environment variable
		// for the home directory is empty or unset in the environment.
<原文结束>

# <翻译开始>
		// UserHomeDir 如果环境变量指定的家目录为空或未设置，可能会返回一个非nil的错误。
# <翻译结束>


<原文开始>
			// The user's home directory has a well-defined location, but does not
			// exist. (Maybe nothing has written to it yet? That could happen, for
			// example, on minimal VM images used for CI testing.)
<原文结束>

# <翻译开始>
			// 用户的主目录有一个明确的位置，但可能不存在。（也许还没有写入内容？这在用于CI测试的最小VM映像中可能发生。）
# <翻译结束>


<原文开始>
	// See issue 37161. Read only one entry from a directory,
	// seek to the beginning, and read again. We should not see
	// duplicate entries.
<原文结束>

# <翻译开始>
	// 参见问题37161。从目录中只读取一个条目，然后移动到开头再读取一次。我们不应该看到重复的条目。
# <翻译结束>


<原文开始>
// isDeadlineExceeded reports whether err is or wraps ErrDeadlineExceeded.
// We also check that the error has a Timeout method that returns true.
<原文结束>

# <翻译开始>
// isDeadlineExceeded 判断 err 是否为 ErrDeadlineExceeded 错误或其封装错误。
// 同时，我们还会检查该错误是否具有返回 true 的 Timeout 方法。
# <翻译结束>


<原文开始>
// Test that opening a file does not change its permissions.  Issue 38225.
<原文结束>

# <翻译开始>
// 测试打开一个文件不会改变其权限。问题38225。
# <翻译结束>


<原文开始>
	// On Windows, we force the MFT to update by reading the actual metadata from GetFileInformationByHandle and then
	// explicitly setting that. Otherwise it might get out of sync with FindFirstFile. See golang.org/issues/42637.
<原文结束>

# <翻译开始>
	// 在Windows系统上，我们通过从GetFileInformationByHandle获取实际元数据并明确设置它来强制更新MFT。否则它可能会与FindFirstFile不一致。参见golang.org/issues/42637。
# <翻译结束>


<原文开始>
// This uses GetFileInformationByHandle internally.
<原文结束>

# <翻译开始>
// 这内部使用了GetFileInformationByHandle。
# <翻译结束>


<原文开始>
// We only log, not die, in case the test directory is not writable.
<原文结束>

# <翻译开始>
// 仅在测试目录不可写时进行日志记录，而不强制退出。
# <翻译结束>


<原文开始>
	// Test that the error message does not contain a backslash,
	// and does not contain the DirFS argument.
<原文结束>

# <翻译开始>
	// 测试错误消息中不包含反斜杠，
	// 且不包含DirFS参数。
# <翻译结束>


<原文开始>
// Test that Open does not accept backslash as separator.
<原文结束>

# <翻译开始>
// 测试Open方法是否接受反斜杠作为分隔符。
# <翻译结束>


<原文开始>
// Test that Open does not open Windows device files.
<原文结束>

# <翻译开始>
// 测试Open是否不会打开Windows设备文件。
# <翻译结束>


<原文开始>
// trim volume prefix (C:) on Windows
<原文结束>

# <翻译开始>
// 在Windows上移除卷标前缀（如C:)
# <翻译结束>


<原文开始>
// Test that Open can open a path starting at /.
<原文结束>

# <翻译开始>
// 测试Open函数是否可以打开以/开头的路径。
# <翻译结束>


<原文开始>
	// Linux files in /proc report 0 size,
	// but then if ReadFile reads just a single byte at offset 0,
	// the read at offset 1 returns EOF instead of more data.
	// ReadFile has a minimum read size of 512 to work around this,
	// but test explicitly that it's working.
<原文结束>

# <翻译开始>
	// 在Linux中，/proc目录下的文件报告大小为0，
	// 但如果使用ReadFile在偏移量0处读取单个字节后，
	// 偏移量1处的读取会返回EOF（文件结束）而不是更多的数据。
	// ReadFile有一个最小读取大小为512的处理机制来解决这个问题，
	// 但我们需要明确测试它是否有效。
# <翻译结束>


<原文开始>
// Test that it's OK to have parallel I/O and Close on a pipe.
<原文结束>

# <翻译开始>
// 测试并行I/O和管道Close操作是否正常。
# <翻译结束>


<原文开始>
// Skip on wasm, which doesn't have pipes.
<原文结束>

# <翻译开始>
// 在没有管道的wasm上跳过
# <翻译结束>


<原文开始>
				// We look at error strings as the
				// expected errors are OS-specific.
<原文结束>

# <翻译开始>
				// 我们将错误字符串视为依赖操作系统的
				// 预期错误。
# <翻译结束>


<原文开始>
// Ignore an expected error.
<原文结束>

# <翻译开始>
// 忽略预期的错误。
# <翻译结束>


<原文开始>
		// Let the other goroutines start. This is just to get
		// a better test, the test will still pass if they
		// don't start.
<原文结束>

# <翻译开始>
		// 让其他goroutine开始。这只是为了让测试更好，即使它们不开始，测试仍然会通过。
# <翻译结束>


<原文开始>
// Test that it's OK to call Close concurrently on a pipe.
<原文结束>

# <翻译开始>
// 测试对管道同时调用Close方法是正常的
# <翻译结束>


<原文开始>
//func testStartProcess(dir, cmd string, args []string, expect string) func(t *testing.T) {
//	return func(t *testing.T) {
//		t.Parallel()
//
//		r, w, err := Pipe()
//		if err != nil {
//			t.Fatalf("Pipe: %v", err)
//		}
//		defer r.Close()
//
//		attr := &ProcAttr{Dir: dir, Files: []*File{nil, w, Stderr}}
//		p, err := StartProcess(cmd, args, attr)
//		if err != nil {
//			t.Fatalf("StartProcess: %v", err)
//		}
//		w.Close()
//
//		var b strings.Builder
//		io.Copy(&b, r)
//		output := b.String()
//
//		fi1, _ := Stat(strings.TrimSpace(output))
//		fi2, _ := Stat(expect)
//		if !SameFile(fi1, fi2) {
//			t.Errorf("exec %q returned %q wanted %q",
//				strings.Join(append([]string{cmd}, args...), " "), output, expect)
//		}
//		p.Wait()
//	}
//}
<原文结束>

# <翻译开始>
//func testStartProcess(dir, cmd string, args []string, expect string) func(t *testing.T) {
//	return func(t *testing.T) {
//		t.Parallel()
//
//		r, w, err := Pipe()
//		if err != nil {
//			t.Fatalf("Pipe: %v", err)
//		}
//		defer r.Close()
//
//		attr := &ProcAttr{Dir: dir, Files: []*File{nil, w, Stderr}}
//		p, err := StartProcess(cmd, args, attr)
//		if err != nil {
//			t.Fatalf("StartProcess: %v", err)
//		}
//		w.Close()
//
//		var b strings.Builder
//		io.Copy(&b, r)
//		output := b.String()
//
//		fi1, _ := Stat(strings.TrimSpace(output))
//		fi2, _ := Stat(expect)
//		if !SameFile(fi1, fi2) {
//			t.Errorf("exec %q returned %q wanted %q",
//				strings.Join(append([]string{cmd}, args...), " "), output, expect)
//		}
//		p.Wait()
//	}
//}
# <翻译结束>


<原文开始>
//func TestStartProcess(t *testing.T) {
//	testenv.MustHaveExec(t)
//	t.Parallel()
//
//	var dir, cmd string
//	var args []string
//	switch runtime.GOOS {
//	case "android":
//		t.Skip("android doesn't have /bin/pwd")
//	case "windows":
//		cmd = Getenv("COMSPEC")
//		dir = Getenv("SystemRoot")
//		args = []string{"/c", "cd"}
//	default:
//		var err error
//		cmd, err = exec.LookPath("pwd")
//		if err != nil {
//			t.Fatalf("Can't find pwd: %v", err)
//		}
//		dir = "/"
//		args = []string{}
//		t.Logf("Testing with %v", cmd)
//	}
//	cmddir, cmdbase := filepath.Split(cmd)
//	args = append([]string{cmdbase}, args...)
//	t.Run("absolute", testStartProcess(dir, cmd, args, dir))
//	t.Run("relative", testStartProcess(cmddir, cmdbase, args, cmddir))
//}
<原文结束>

# <翻译开始>
// ```go
//func TestStartProcess(t *testing.T) {
// 测试环境必须支持执行命令
//testenv.MustHaveExec(t)
// 并行执行测试
//t.Parallel()
// 
// 根据操作系统设置变量
//var dir, cmd string
//var args []string
// 判断当前操作系统
//switch runtime.GOOS {
// 如果是Android，跳过测试（因为没有/bin/pwd）
//case "android":
//	t.Skip("android doesn't have /bin/pwd")
// 如果是Windows
//case "windows":
// 设置cmd为环境变量COMSPEC的值
//	cmd = Getenv("COMSPEC")
// 设置dir为环境变量SystemRoot的值
//	dir = Getenv("SystemRoot")
// 设置args为/c和cd
//	args = []string{"/c", "cd"}
// 其他情况（默认为非Windows系统）
//default:
// 使用LookPath查找"pwd"命令
//	var err error
//	cmd, err = exec.LookPath("pwd")
// 如果找不到"pwd"，终止测试
//	if err != nil {
//		t.Fatalf("Can't find pwd: %v", err)
//	}
// 设置dir为"/"
//	dir = "/"
// 设置args为空
//	args = []string{}
// 打印正在使用的命令进行测试
//	t.Logf("Testing with %v", cmd)
//}
// 
// 分离cmd的目录和基础命令
//cmddir, cmdbase := filepath.Split(cmd)
// 将基础命令添加到args列表的前面
//args = append([]string{cmdbase}, args...)
// 运行绝对路径测试
//t.Run("absolute", testStartProcess(dir, cmd, args, dir))
// 运行相对路径测试
//t.Run("relative", testStartProcess(cmddir, cmdbase, args, cmddir))
//}
// ```
// 
// 这段代码是一个Go语言的测试函数，用于测试`StartProcess`功能。它根据不同的操作系统设置命令和参数，并分别测试绝对路径和相对路径的情况。
# <翻译结束>


<原文开始>
//
//func TestRandomLen(t *testing.T) {
//	for range 5 {
//		dir, err := MkdirTemp(t.TempDir(), "*")
//		if err != nil {
//			t.Fatal(err)
//		}
//		base := filepath.Base(dir)
//		if len(base) > 10 {
//			t.Errorf("MkdirTemp returned len %d: %s", len(base), base)
//		}
//	}
//	for range 5 {
//		f, err := CreateTemp(t.TempDir(), "*")
//		if err != nil {
//			t.Fatal(err)
//		}
//		base := filepath.Base(f.Name())
//		f.Close()
//		if len(base) > 10 {
//			t.Errorf("CreateTemp returned len %d: %s", len(base), base)
//		}
//	}
//}
<原文结束>

# <翻译开始>
// 
// // 测试RandomLen函数
// func TestRandomLen(t *testing.T) {
//     // 重复5次循环
//     for i := 0; i < 5; i++ {
//         dir, err := MkdirTemp(t.TempDir(), "*") // 在临时目录下创建子目录
//         if err != nil {
//             t.Fatal(err) // 如果出错，打印错误并终止测试
//         }
//         base := filepath.Base(dir) // 获取子目录的基名
//         if len(base) > 10 { // 如果基名长度超过10
//             t.Errorf("MkdirTemp返回的基名长度为%d：%s", len(base), base) // 打印错误信息
//         }
//     }
//     // 再次重复5次循环
//     for i := 0; i < 5; i++ {
//         f, err := CreateTemp(t.TempDir(), "*") // 在临时目录下创建临时文件
//         if err != nil {
//             t.Fatal(err) // 如果出错，打印错误并终止测试
//         }
//         base := filepath.Base(f.Name()) // 获取文件的基名
//         f.Close() // 关闭文件
//         if len(base) > 10 { // 如果基名长度超过10
//             t.Errorf("CreateTemp返回的基名长度为%d：%s", len(base), base) // 打印错误信息
//         }
//     }
// }
# <翻译结束>

