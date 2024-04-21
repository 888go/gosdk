
<原文开始>
// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// 版权所有 2018 The Go Authors。保留所有权利。
// 使用本源代码须遵循 BSD 风格许可证，
// 该许可证可在 LICENSE 文件中找到。
# <翻译结束>


<原文开始>
// Make a regular file and remove
<原文结束>

# <翻译开始>
// 创建一个普通文件并删除
# <翻译结束>


<原文开始>
// Make directory with 1 file and remove.
<原文结束>

# <翻译开始>
// 创建一个包含一个文件的目录并删除它。
# <翻译结束>


<原文开始>
// Make directory with file and subdirectory and remove.
<原文结束>

# <翻译开始>
// 创建带有文件和子目录的目录并将其删除
# <翻译结束>


<原文开始>
// Chmod is not supported under Windows or wasip1 and test fails as root.
<原文结束>

# <翻译开始>
// Chmod 在Windows或wasip1下不被支持，而且以root用户运行测试会导致失败。
# <翻译结束>


<原文开始>
// Make directory with file and subdirectory and trigger error.
<原文结束>

# <翻译开始>
// 创建包含文件和子目录的目录，并触发错误。
# <翻译结束>


<原文开始>
		// No error checking here: either RemoveAll
		// will or won't be able to remove dpath;
		// either way we want to see if it removes fpath
		// and path/zzz. Reasons why RemoveAll might
		// succeed in removing dpath as well include:
		//	* running as root
		//	* running on a file system without permissions (FAT)
<原文结束>

# <翻译开始>
// 在这里不进行错误检查：RemoveAll 或许能够移除 dpath，或许不能；无论如何，我们都想知道它是否能移除 fpath 和 path/zzz。RemoveAll 也可能会成功移除 dpath 的原因包括：
// * 以 root 身份运行
// * 在没有权限的文件系统（如 FAT）上运行
# <翻译结束>


<原文开始>
// Test RemoveAll on a large directory.
<原文结束>

# <翻译开始>
// 测试在大型目录上使用RemoveAll
# <翻译结束>


<原文开始>
// Make directory with 1000 files and remove.
<原文结束>

# <翻译开始>
// 创建一个包含1000个文件的目录，然后删除。
# <翻译结束>


<原文开始>
// Removing paths with over 4096 chars commonly fails
<原文结束>

# <翻译开始>
// 删除超过4096个字符的路径通常会失败
# <翻译结束>


<原文开始>
	// If an error occurs make it more likely that removing the
	// temporary directory will succeed.
<原文结束>

# <翻译开始>
// 如果发生错误，设法提高删除临时目录成功的可能性。
# <翻译结束>


<原文开始>
		// Defer changing the mode back so that the deferred
		// RemoveAll(tempDir) can succeed.
<原文结束>

# <翻译开始>
// 延迟恢复原模式，以便于后续执行的deferred RemoveAll(tempDir)能够成功。
# <翻译结束>


<原文开始>
	// The error should be of type *PathError.
	// see issue 30491 for details.
<原文结束>

# <翻译开始>
// 错误应为类型 *PathError。
// 有关详细信息，请参阅问题 30491。
# <翻译结束>


<原文开始>
// Make directory with 1025 read-only files.
<原文结束>

# <翻译开始>
// 创建一个包含1025个只读文件的目录。
# <翻译结束>


<原文开始>
	// Make the parent directory read-only. On some platforms, this is what
	// prevents Remove from removing the files within that directory.
<原文结束>

# <翻译开始>
// 将父目录设置为只读。在某些平台上，这是阻止Remove删除该目录内文件的原因。
# <翻译结束>


<原文开始>
	// This call should not hang, even on a platform that disallows file deletion
	// from read-only directories.
<原文结束>

# <翻译开始>
// 此调用不应阻塞，即使在不允许从只读目录中删除文件的平台上也是如此
# <翻译结束>


<原文开始>
// On many platforms, root can remove files from read-only directories.
<原文结束>

# <翻译开始>
// 在许多平台上，超级用户可以删除只读目录中的文件。
# <翻译结束>


<原文开始>
			// Marking a directory as read-only in Windows does not prevent the RemoveAll
			// from creating or removing files within it.
			//
			// For wasip1, there is no support for file permissions so we cannot prevent
			// RemoveAll from removing the files.
<原文结束>

# <翻译开始>
// 在Windows中，将目录标记为只读并不能阻止RemoveAll在其中创建或删除文件。
//
// 对于wasip1，没有文件权限支持，因此我们无法阻止RemoveAll删除文件。
# <翻译结束>


<原文开始>
	// Only test on Linux so that we can assume we have strace.
	// The code is OS-independent so if it passes on Linux
	// it should pass on other Unix systems.
<原文结束>

# <翻译开始>
// 只在Linux上进行测试，这样我们就可以假设我们有strace工具。
// 该代码是跨平台的，所以如果在Linux上通过了，那么在其他Unix系统上也应该通过。
# <翻译结束>


<原文开始>
	// Create 100 directories.
	// The test is that we can remove them without calling fcntl
	// on each one.
<原文结束>

# <翻译开始>
// 创建100个目录。
// 测试目标是我们可以在不针对每个目录调用fcntl的情况下将其删除。
# <翻译结束>

