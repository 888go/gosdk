
<原文开始>
// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// 版权所有 2020 Go作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。
# <翻译结束>


<原文开始>
// through traditional means
<原文结束>

# <翻译开始>
// 通过传统方式
# <翻译结束>


<原文开始>
// Read data from the file itself.
<原文结束>

# <翻译开始>
// 从文件本身读取数据
# <翻译结束>


<原文开始>
// It should wind up a double of the original data.
<原文结束>

# <翻译开始>
// 它应该得到原始数据的两倍。
# <翻译结束>


<原文开始>
			// The pipe is empty, and PIPE_BUF is large enough
			// for this, by (POSIX) definition, so there is no
			// need for an additional goroutine.
<原文结束>

# <翻译开始>
			// 管道为空，并且根据(POSIX)定义，PIPE_BUF足够大，因此不需要额外的goroutine。
# <翻译结束>


<原文开始>
// If we have a limit, wrap the reader.
<原文结束>

# <翻译开始>
// 如果我们有配额，包装读取器。
# <翻译结束>


<原文开始>
// Now call ReadFrom (through io.Copy), which will hopefully call poll.Splice
<原文结束>

# <翻译开始>
// 现在调用 ReadFrom（通过 io.Copy 实现），期望这将触发对 poll.Splice 的调用
# <翻译结束>


<原文开始>
// We should have called poll.Splice with the right file descriptor arguments.
<原文结束>

# <翻译开始>
// 我们应该使用正确的文件描述符参数调用poll.Splice。
# <翻译结束>


<原文开始>
	// Check that the offsets after the transfer make sense, that the size
	// of the transfer was reported correctly, and that the destination
	// file contains exactly the bytes we expect it to contain.
<原文结束>

# <翻译开始>
	// 检查传输后的偏移量是否合理，确认传输的大小被正确报告，并且目标文件确实包含我们期望的字节。
# <翻译结束>


<原文开始>
// If we had a limit, check that it was updated.
<原文结束>

# <翻译开始>
// 如果我们有上限，检查它是否已更新。
# <翻译结束>


<原文开始>
	// Call wg.Wait as the final deferred function,
	// because the goroutines may block until some of
	// the deferred Close calls.
<原文结束>

# <翻译开始>
	// 将 wg.Wait 作为最终的延迟调用函数，
	// 因为 goroutine 可能会阻塞直到某些延迟的 Close 调用完成。
# <翻译结束>


<原文开始>
	// Open the tty directly, rather than via OpenFile.
	// This bypasses the non-blocking support and is required
	// to recreate the problem in the issue (#59041).
<原文结束>

# <翻译开始>
	// 直接打开tty，而不是通过OpenFile。
	// 这绕过了非阻塞支持，并且是重现问题所必需的
	// 在问题(#59041)中。
# <翻译结束>


<原文开始>
		// The problem (issue #59041) occurs when writing
		// a series of blocks of data. It does not occur
		// when all the data is written at once.
<原文结束>

# <翻译开始>
		// 当写入一系列数据块时（问题#59041），问题会出现。如果一次性写入所有数据，问题不会出现。
# <翻译结束>


<原文开始>
				// If we get here because the client was
				// closed, skip the error.
<原文结束>

# <翻译开始>
				// 如果我们因为客户端被关闭而到达这里，跳过错误。
# <翻译结束>


<原文开始>
					// An error here doesn't matter for
					// our test.
<原文结束>

# <翻译开始>
					// 对于我们的测试而言，此处的错误无关紧要。
# <翻译结束>


<原文开始>
// Close Client to wake up the writing goroutine if necessary.
<原文结束>

# <翻译开始>
// 关闭客户端，如果需要的话，唤醒写入的goroutine。
# <翻译结束>


<原文开始>
	// Now call ReadFrom (through io.Copy), which will hopefully call
	// poll.CopyFileRange.
<原文结束>

# <翻译开始>
	// 现在调用ReadFrom（通过io.Copy），它将希望调用poll.CopyFileRange。
# <翻译结束>


<原文开始>
	// If we didn't have a limit, we should have called poll.CopyFileRange
	// with the right file descriptor arguments.
<原文结束>

# <翻译开始>
	// 如果我们没有限制，我们应该使用正确的文件描述符参数调用poll.CopyFileRange。
# <翻译结束>


<原文开始>
// newCopyFileRangeTest initializes a new test for copy_file_range.
//
// It creates source and destination files, and populates the source file
// with random data of the specified size. It also hooks package os' call
// to poll.CopyFileRange and returns the hook so it can be inspected.
<原文结束>

# <翻译开始>
// newCopyFileRangeTest 初始化一个用于测试 copy_file_range 的新测试。
//
// 它创建源文件和目标文件，并用指定大小的随机数据填充源文件。此外，它还对 package os 调用 poll.CopyFileRange 进行挂钩，并返回该挂钩以便进行检查。
# <翻译结束>


<原文开始>
	// Populate the source file with data, then rewind it, so it can be
	// consumed by copy_file_range(2).
<原文结束>

# <翻译开始>
	// 向源文件填充数据，然后将其回溯，以便可以由copy_file_range(2)函数消费。
# <翻译结束>


<原文开始>
// newSpliceFileTest initializes a new test for splice.
//
// It creates source sockets and destination file, and populates the source sockets
// with random data of the specified size. It also hooks package os' call
// to poll.Splice and returns the hook so it can be inspected.
<原文结束>

# <翻译开始>
// newSpliceFileTest 初始化splice的新测试。
// 
// 它创建源套接字和目标文件，并在源套接字中填充指定大小的随机数据。它还挂钩包"os"对poll.Splice的调用，并返回钩子，以便可以检查它。
# <翻译结束>


<原文开始>
// mustContainData ensures that the specified file contains exactly the
// specified data.
<原文结束>

# <翻译开始>
// mustContainData 确保指定的文件包含确切指定的数据。
# <翻译结束>


<原文开始>
// On some kernels copy_file_range fails on files in /proc.
<原文结束>

# <翻译开始>
// 在某些内核中，copy_file_range 对于 /proc 下的文件会失败
# <翻译结束>

