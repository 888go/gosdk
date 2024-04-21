
<原文开始>
// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// 版权所有 2023 The Go Authors。保留所有权利。
// 使用本源代码须遵循 BSD 风格的许可证，
// 该许可证可在 LICENSE 文件中找到。
# <翻译结束>


<原文开始>
// Now call WriteTo (through io.Copy), which will hopefully call poll.SendFile
<原文结束>

# <翻译开始>
// 现在调用WriteTo（通过io.Copy），这将希望调用poll.SendFile
# <翻译结束>


<原文开始>
// We should have called poll.Splice with the right file descriptor arguments.
<原文结束>

# <翻译开始>
// 我们应该使用正确的文件描述符参数调用poll.Splice。
# <翻译结束>


<原文开始>
// Verify the data size and content.
<原文结束>

# <翻译开始>
// 验证数据的大小和内容。
# <翻译结束>


<原文开始>
// newSendFileTest initializes a new test for sendfile.
//
// It creates source file and destination sockets, and populates the source file
// with random data of the specified size. It also hooks package os' call
// to poll.Sendfile and returns the hook so it can be inspected.
<原文结束>

# <翻译开始>
// newSendFileTest 初始化一个用于sendfile的全新测试。
//
// 它创建源文件与目标套接字，并使用指定大小的随机数据填充源文件。同时，它对package os调用poll.Sendfile进行挂钩，并返回该挂钩以便进行检查。
# <翻译结束>

