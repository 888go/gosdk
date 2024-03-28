
<原文开始>
// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// 版权所有 ? 2020 Go作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。
# <翻译结束>


<原文开始>
// This is a test program that verifies that it can read from
// descriptor 3 and that no other descriptors are open.
// This is not done via TestHelperProcess and GO_EXEC_TEST_PID
// because we want to ensure that this program does not use cgo,
// because C libraries can open file descriptors behind our backs
// and confuse the test. See issue 25628.
<原文结束>

# <翻译开始>
// 这是一个测试程序，用于验证可以从描述符3读取数据，并且没有其他描述符打开。
// 我们不通过TestHelperProcess和GO_EXEC_TEST_PID来完成这个操作，因为我们希望确保该程序不使用cgo。
// 因为C库可以在后台打开文件描述符，从而混淆我们的测试。参见问题25628。
# <翻译结束>


<原文开始>
	// Now verify that there are no other open fds.
	// stdin == 0
	// stdout == 1
	// stderr == 2
	// descriptor from parent == 3
	// All descriptors 4 and up should be available,
	// except for any used by the network poller.
<原文结束>

# <翻译开始>
// 现在验证没有其他打开的文件描述符。
// 标准输入 == 0
// 标准输出 == 1
// 标准错误输出 == 2
// 从父进程继承的描述符 == 3
// 所有 4 及以上的描述符都应该可用，
// 除了网络轮询器正在使用的那些。
# <翻译结束>


<原文开始>
// Determine which command to use to display open files.
<原文结束>

# <翻译开始>
// 确定用于显示打开文件的命令。
# <翻译结束>

