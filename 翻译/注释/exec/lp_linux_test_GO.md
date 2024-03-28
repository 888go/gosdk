
<原文开始>
// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// 版权所有 ? 2022 Go作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。
# <翻译结束>


<原文开始>
// This test case relies on faccessat2(2) syscall, which appeared in Linux v5.8.
<原文结束>

# <翻译开始>
// 该测试用例依赖于Linux v5.8中出现的faccessat2(2)系统调用。
# <翻译结束>


<原文开始>
		// Usually this means lack of CAP_SYS_ADMIN, but there might be
		// other reasons, especially in restricted test environments.
<原文结束>

# <翻译开始>
// 通常，这意味着缺少 CAP_SYS_ADMIN 权限，但在受限的测试环境中，可能还有其他原因。
# <翻译结束>


<原文开始>
// Check that it works as expected.
<原文结束>

# <翻译开始>
// 检查它是否按预期工作。
# <翻译结束>


<原文开始>
			// A fork+exec in another process may be holding open the FD that we used
			// to write the executable (see https://go.dev/issue/22315).
			// Since the descriptor should have CLOEXEC set, the problem should resolve
			// as soon as the forked child reaches its exec call.
			// Keep retrying until that happens.
<原文结束>

# <翻译开始>
// 另一个进程中的 fork+exec 可能正保持着我们用于写入可执行文件的文件描述符（参见 https://go.dev/issue/22315）。
// 由于该描述符应已设置 CLOEXEC 标志，当 fork 子进程执行到其 exec 调用时，问题应该会得到解决。
// 因此，我们需要一直重试直到这种情况发生。
# <翻译结束>


<原文开始>
// Remount with noexec flag.
<原文结束>

# <翻译开始>
// 使用noexec标志重新挂载。
# <翻译结束>

