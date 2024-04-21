
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
// For TestRawConnReadWrite.
<原文结束>

# <翻译开始>
// 用于TestRawConnReadWrite测试。
# <翻译结束>


<原文开始>
	// Use TempDir() to make sure we're on a local file system,
	// so that the group ids returned by Getgroups will be allowed
	// on the file. On NFS, the Getgroups groups are
	// basically useless.
<原文结束>

# <翻译开始>
// 使用TempDir()来确保我们在本地文件系统上，这样Getgroups返回的组ID将被允许应用于文件。在NFS上，Getgroups获取的组基本上没有用。
# <翻译结束>


<原文开始>
	// Can't change uid unless root, but can try
	// changing the group id. First try our current group.
<原文结束>

# <翻译开始>
// 如果不是root用户，不能改变uid，但可以尝试改变gid。首先尝试使用当前组。
# <翻译结束>


<原文开始>
// Then try all the auxiliary groups.
<原文结束>

# <翻译开始>
// 接着尝试所有的辅助组
# <翻译结束>


<原文开始>
// Since the Chown call failed, the file should be unmodified.
<原文结束>

# <翻译开始>
// 由于Chown调用失败，文件应该保持未修改。
# <翻译结束>


<原文开始>
// change back to gid to test fd.Chown
<原文结束>

# <翻译开始>
// 改回gid以测试fd.Chown
# <翻译结束>


<原文开始>
// Since the Lchown call failed, the file should be unmodified.
<原文结束>

# <翻译开始>
// 由于Lchown调用失败，文件应该未被修改。
# <翻译结束>


<原文开始>
// Check that link target's gid is unchanged.
<原文结束>

# <翻译开始>
// 检查链接目标的gid未发生改变
# <翻译结束>


<原文开始>
// Issue 16919: Readdir must return a non-empty slice or an error.
<原文结束>

# <翻译开始>
// 问题16919：Readdir必须返回一个非空的切片或一个错误。
# <翻译结束>


<原文开始>
// Act like it's been deleted.
<原文结束>

# <翻译开始>
// 像被删除一样表现。
# <翻译结束>


<原文开始>
// notably, greater than zero
<原文结束>

# <翻译开始>
// 值得注意的是，大于零
# <翻译结束>


<原文开始>
// This is what used to happen (Issue 16919)
<原文结束>

# <翻译开始>
// 这是过去发生的情况（问题 16919）
# <翻译结束>


<原文开始>
// Issue 23120: respect umask when doing Mkdir with the sticky bit
<原文结束>

# <翻译开始>
// 问题23120：在使用Mkdir时尊重umask设置，同时处理粘滞位
# <翻译结束>


<原文开始>
// See also issues: 22939, 24331
<原文结束>

# <翻译开始>
// 参见问题：22939，24331
# <翻译结束>


<原文开始>
// Set the read-side to non-blocking.
<原文结束>

# <翻译开始>
// 将读取侧设置为非阻塞。
# <翻译结束>


<原文开始>
		// Use a longer time to avoid flakes.
		// We won't be waiting this long anyhow.
<原文结束>

# <翻译开始>
// 使用较长的时间以避免出现异常。
// 无论如何我们都不会等待这么长时间。
# <翻译结束>


<原文开始>
// Try to read with deadline (but don't block forever).
<原文结束>

# <翻译开始>
// 尝试在截止时间前读取（但不要永远阻塞）。
# <翻译结束>


<原文开始>
// We want it to fail with a timeout.
<原文结束>

# <翻译开始>
// 我们希望它因为超时而失败。
# <翻译结束>


<原文开始>
// We want it to succeed after 100ms
<原文结束>

# <翻译开始>
// 我们希望它在100毫秒后成功
# <翻译结束>


<原文开始>
// Test that copying to files opened with O_APPEND works and
// the copy_file_range syscall isn't used on Linux.
//
// Regression test for go.dev/issue/60181
<原文结束>

# <翻译开始>
// 测试将内容复制到以 O_APPEND 方式打开的文件中是否正常工作，以及在 Linux 系统上未使用 copy_file_range 系统调用。
//
// 作为对 go.dev/issue/60181 问题的回归测试
# <翻译结束>


<原文开始>
	// This would fail on Linux in case the copy_file_range syscall was used because it doesn't
	// support destination files opened with O_APPEND, see
	// https://man7.org/linux/man-pages/man2/copy_file_range.2.html#ERRORS
<原文结束>

# <翻译开始>
// 在Linux上，如果使用了copy_file_range系统调用，此操作将失败，因为它不支持以O_APPEND打开的目标文件，详情见
// https://man7.org/linux/man-pages/man2/copy_file_range.2.html#ERRORS
# <翻译结束>

