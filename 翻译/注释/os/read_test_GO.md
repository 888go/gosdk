
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
// We don't want to use CreateTemp directly, since that opens a file for us as 0600.
<原文结束>

# <翻译开始>
// 我们不希望直接使用CreateTemp，因为它会为我们创建一个权限为0600的文件。
# <翻译结束>


<原文开始>
	//if Getuid() == 0 {
	//	t.Skipf("Root can write to read-only files anyway, so skip the read-only test.")
	//}
<原文结束>

# <翻译开始>
// 如果Getuid()返回0（表示当前用户是root）{
//   t.Skipf("由于root用户可以无论如何写入只读文件，因此跳过只读文件的测试。")
// }
// md5:e0e59d68eb648704
# <翻译结束>

