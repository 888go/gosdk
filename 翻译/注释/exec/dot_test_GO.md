
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
// Not parallel: uses Chdir and Setenv.
<原文结束>

# <翻译开始>
// 非并行：使用了 Chdir 和 Setenv。
# <翻译结束>


<原文开始>
	// Add "." to PATH so that exec.LookPath looks in the current directory on all systems.
	// And try to trick it with "../testdir" too.
<原文结束>

# <翻译开始>
// 向PATH环境变量中添加"."，确保在所有系统上exec.LookPath函数都会在当前目录中查找可执行文件。
// 同时也尝试用"../testdir"来测试迷惑它。
# <翻译结束>


<原文开始>
					// Clearing cmd.Err should let the execution proceed,
					// and it should fail because it's not a valid binary.
<原文结束>

# <翻译开始>
// 清除cmd.Err应允许执行继续进行，
// 并且由于它不是一个有效的二进制文件，所以应该会失败。
# <翻译结束>


<原文开始>
	// Test the behavior when the first entry in PATH is an absolute name for the
	// current directory.
	//
	// On Windows, "." may or may not be implicitly included before the explicit
	// %PATH%, depending on the process environment;
	// see https://go.dev/issue/4394.
	//
	// If the relative entry from "." resolves to the same executable as what
	// would be resolved from an absolute entry in %PATH% alone, LookPath should
	// return the absolute version of the path instead of ErrDot.
	// (See https://go.dev/issue/53536.)
	//
	// If PATH does not implicitly include "." (such as on Unix platforms, or on
	// Windows configured with NoDefaultCurrentDirectoryInExePath), then this
	// lookup should succeed regardless of the behavior for ".", so it may be
	// useful to run as a control case even on those platforms.
<原文结束>

# <翻译开始>
// 测试当PATH环境变量中的第一个条目为当前目录的绝对路径时的行为。
//
// 在Windows系统上，"." 可能会被隐式地包含在显式的 %PATH% 之前，这取决于进程环境；
// 详细信息请参阅：https://go.dev/issue/4394。
//
// 如果从"."解析出的相对路径与仅通过 %PATH% 中绝对路径解析到的可执行文件相同，LookPath 应返回该路径的绝对版本而非 ErrDot。
// （参考：https://go.dev/issue/53536。）
//
// 若 PATH 环境变量未隐式包含 "."（例如，在Unix平台或在Windows系统中配置为NoDefaultCurrentDirectoryInExePath时），则此查找应当无论如何都能成功，因此即使在这些平台上，将其作为对照测试案例运行也是有用的。
# <翻译结束>


<原文开始>
		// Control case: if the lookup returns ErrDot when PATH is empty, then we
		// know that PATH implicitly includes ".". If it does not, then we don't
		// expect to see ErrDot at all in this test (because the path will be
		// unambiguously absolute).
<原文结束>

# <翻译开始>
// 控制用例：如果当 PATH 为空时，查找返回 ErrDot 错误，那么我们知道 PATH 暗示性地包含了 "."。若不是这样，则在这个测试中我们完全不期望看到 ErrDot 错误（因为路径将被明确无误地视为绝对路径）。
# <翻译结束>


<原文开始>
		// Set PATH to include an explicit directory that contains a completely
		// independent executable that happens to have the same name as an
		// executable in ".". If "." is included implicitly, looking up the
		// (unqualified) executable name will return ErrDot; otherwise, the
		// executable in "." should have no effect and the lookup should
		// unambiguously resolve to the directory in PATH.
<原文结束>

# <翻译开始>
// 设置PATH环境变量，包含一个明确的目录，该目录下包含一个恰好与"."目录下的可执行文件同名的独立可执行文件。如果"."被隐式包含在内，查找（未限定路径的）可执行文件名将返回ErrDot错误；否则，"."目录下的可执行文件不应产生影响，并且查找应当无歧义地解析到PATH环境变量中指定的目录下的可执行文件。
# <翻译结束>

