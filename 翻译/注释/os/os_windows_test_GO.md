
<原文开始>
// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// 版权所有 2014 Go 作者。保留所有权利。
// 本源代码的使用由bsd风格的许可证控制，该许可证可在LICENSE文件中找到。
# <翻译结束>


<原文开始>
// For TestRawConnReadWrite.
<原文结束>

# <翻译开始>
// 用于TestRawConnReadWrite测试。
# <翻译结束>


<原文开始>
// chdir changes the current working directory to the named directory,
// and then restore the original working directory at the end of the test.
<原文结束>

# <翻译开始>
// chdir 将当前工作目录更改为指定的目录，然后在测试结束时恢复原来的目录。
# <翻译结束>


<原文开始>
// correspondent issue number (for broken tests)
<原文结束>

# <翻译开始>
// 对应的问题编号（针对失效的测试）
# <翻译结束>


<原文开始>
// reparseData is used to build reparse buffer data required for tests.
<原文结束>

# <翻译开始>
// reparseData 用于构建测试所需的重新解析缓冲区数据。
# <翻译结束>


<原文开始>
// do not include terminating NUL in the length (as per PrintNameLength and SubstituteNameLength documentation)
<原文结束>

# <翻译开始>
// 不要在长度中包含终止 NULL（根据 PrintNameLength 和 SubstituteNameLength 文档）
# <翻译结束>


<原文开始>
// pathBuffeLen returns length of rd pathBuf in bytes.
<原文结束>

# <翻译开始>
// pathBuffeLen 返回rd路径缓冲区的字节长度。
# <翻译结束>


<原文开始>
// Windows REPARSE_DATA_BUFFER contains union member, and cannot be
// translated into Go directly. _REPARSE_DATA_BUFFER type is to help
// construct alternative versions of Windows REPARSE_DATA_BUFFER with
// union part of SymbolicLinkReparseBuffer or MountPointReparseBuffer type.
<原文结束>

# <翻译开始>
// Windows REPARSE_DATA_BUFFER 结构包含联合成员，无法直接转换为 Go 语言表示。_REPARSE_DATA_BUFFER 类型旨在帮助构建 Windows REPARSE_DATA_BUFFER 的替代版本，其中联合部分为 SymbolicLinkReparseBuffer 或 MountPointReparseBuffer 类型。
# <翻译结束>


<原文开始>
// see ReparseDataLength documentation
<原文结束>

# <翻译开始>
// 参考ReparseDataLength的文档说明
# <翻译结束>


<原文开始>
// Create link similar to what mklink does, by inserting \??\ at the front of absolute target.
<原文结束>

# <翻译开始>
// 创建类似于mklink所做的链接，通过在绝对目标的前面插入\??\。
# <翻译结束>


<原文开始>
// Do as junction utility https://learn.microsoft.com/en-us/sysinternals/downloads/junction does - set PrintNameLength to 0.
<原文结束>

# <翻译开始>
// 像junction实用程序（https://learn.microsoft.com/en-us/sysinternals/downloads/junction）一样操作 - 将PrintNameLength设置为0。
# <翻译结束>


<原文开始>
// The rest of these test requires SeCreateSymbolicLinkPrivilege to be held.
<原文结束>

# <翻译开始>
// 余下的这些测试要求拥有SeCreateSymbolicLinkPrivilege权限。
# <翻译结束>


<原文开始>
//LanmanWorkstation is the service name, and Workstation is the display name.
<原文结束>

# <翻译开始>
//LanmanWorkstation是服务名称，而Workstation是显示名称。
# <翻译结束>


<原文开始>
// This normally happens when WSL still doesn't have a distro installed to run on.
<原文结束>

# <翻译开始>
// 这通常发生在WSL尚未安装运行的发行版时。
# <翻译结束>


<原文开始>
// This can happen depending on newer WSL versions when running as admin or in developer mode.
<原文结束>

# <翻译开始>
// 如果以管理员身份运行或在开发模式下运行，根据较新的WSL版本，可能会发生这种情况。
# <翻译结束>


<原文开始>
	// Stat'ing a IO_REPARSE_TAG_LX_SYMLINK from outside WSL always return ERROR_CANT_ACCESS_FILE.
	// We check this condition to validate that os.Stat has tried to follow the link.
<原文结束>

# <翻译开始>
	// 对于一个IO_REPARSE_TAG_LX_SYMLINK类型的符号链接，若从WSL外部进行Stat操作，总会返回ERROR_CANT_ACCESS_FILE错误。
	// 我们检查这一条件，以验证os.Stat确实尝试了跟踪该链接。
# <翻译结束>


<原文开始>
// This sets FILE_ATTRIBUTE_READONLY.
<原文结束>

# <翻译开始>
// 这将设置 FILE_ATTRIBUTE_READONLY 属性。
# <翻译结束>


<原文开始>
// syscallCommandLineToArgv calls syscall.CommandLineToArgv
// and converts returned result into []string.
<原文结束>

# <翻译开始>
// syscallCommandLineToArgv 调用 syscall.CommandLineToArgv 函数，并将返回的结果转换为 []string 类型。
# <翻译结束>


<原文开始>
// compareCommandLineToArgvWithSyscall ensures that
// os.CommandLineToArgv(cmd) and syscall.CommandLineToArgv(cmd)
// return the same result.
<原文结束>

# <翻译开始>
// compareCommandLineToArgvWithSyscall 确保
// os.CommandLineToArgv(cmd) 和 syscall.CommandLineToArgv(cmd)
// 返回相同的结果。
# <翻译结束>


<原文开始>
// examples from https://learn.microsoft.com/en-us/cpp/cpp/main-function-command-line-args
<原文结束>

# <翻译开始>
// 示例来自 https://learn.microsoft.com/zh-cn/cpp/cpp/main-function-command-line-args
# <翻译结束>


<原文开始>
		// http://daviddeley.com/autohotkey/parameters/parameters.htm#WINARGV
		// from 5.4  Examples
<原文结束>

# <翻译开始>
		// 参考：http:		//daviddeley.com/autohotkey/parameters/parameters.htm#WINARGV
		// 从5.4节：示例
# <翻译结束>


<原文开始>
// from 5.5  Some Common Tasks
<原文结束>

# <翻译开始>
// 从5.5 开始 一些常见任务
# <翻译结束>


<原文开始>
// from 5.6  The Microsoft Examples Explained
<原文结束>

# <翻译开始>
// 从 5.6 起，微软示例解释
# <翻译结束>


<原文开始>
// from 5.7  Double Double Quote Examples (pre 2008)
<原文结束>

# <翻译开始>
// 自 5.7 双重双引号示例（2008年前）
# <翻译结束>


<原文开始>
// test both syscall.EscapeArg and os.commandLineToArgv
<原文结束>

# <翻译开始>
// 测试syscall.EscapeArg和os.commandLineToArgv两个函数
# <翻译结束>


<原文开始>
// as per https://stackoverflow.com/questions/42519624/how-to-determine-location-of-onedrive-on-windows-7-and-8-in-c
<原文结束>

# <翻译开始>
// 根据https://stackoverflow.com/questions/42519624/how-to-determine-location-of-onedrive-on-windows-7-and-8-in-c 的说明
# <翻译结束>


<原文开始>
// TestOneDrive verifies that OneDrive folder is a directory and not a symlink.
<原文结束>

# <翻译开始>
// TestOneDrive 验证OneDrive文件夹是一个目录，而不是符号链接。
# <翻译结束>


<原文开始>
// TestSymlinkCreation verifies that creating a symbolic link
// works on Windows when developer mode is active.
// This is supported starting Windows 10 (1703, v10.0.14972).
<原文结束>

# <翻译开始>
// TestSymlinkCreation 验证在开发者模式激活时，创建符号链接在Windows系统上能够正常工作。
// 此功能自Windows 10（版本1703，内部版本号10.0.14972）起得到支持。
# <翻译结束>


<原文开始>
// isWindowsDeveloperModeActive checks whether or not the developer mode is active on Windows 10.
// Returns false for prior Windows versions.
// see https://docs.microsoft.com/en-us/windows/uwp/get-started/enable-your-device-for-development
<原文结束>

# <翻译开始>
// isWindowsDeveloperModeActive 检查Windows 10上是否启用了开发者模式。
// 对于早期版本的Windows，返回false。
// 参考：https://docs.microsoft.com/en-us/windows/uwp/get-started/enable-your-device-for-development
# <翻译结束>


<原文开始>
// TestRootRelativeDirSymlink verifies that symlinks to paths relative to the
// drive root (beginning with "\" but no volume name) are created with the
// correct symlink type.
// (See https://golang.org/issue/39183#issuecomment-632175728.)
<原文结束>

# <翻译开始>
// TestRootRelativeDirSymlink 验证对于从驱动器根目录（以 "\" 开头但没有卷名）开始的路径，symlink 是否被创建为正确的类型。
// （参见 https://golang.org/issue/39183#issuecomment-632175728。）
# <翻译结束>


<原文开始>
// leaves leading backslash
<原文结束>

# <翻译开始>
// 保留开头的反斜杠
# <翻译结束>


<原文开始>
// TestWorkingDirectoryRelativeSymlink verifies that symlinks to paths relative
// to the current working directory for the drive, such as "C:File.txt", are
// correctly converted to absolute links of the correct symlink type (per
// https://docs.microsoft.com/en-us/windows/win32/fileio/creating-symbolic-links).
<原文结束>

# <翻译开始>
// TestWorkingDirectoryRelativeSymlink 验证对于驱动器（如 "C:File.txt"）中相对于当前工作目录的符号链接，能够被正确转换为相应符号链接类型的绝对链接（参照 https://docs.microsoft.com/en-us/windows/win32/fileio/creating-symbolic-links）。
# <翻译结束>


<原文开始>
// Construct a directory to be symlinked.
<原文结束>

# <翻译开始>
// 创建一个用于建立软链接的目录。
# <翻译结束>


<原文开始>
	// Change to the temporary directory and construct a
	// working-directory-relative symlink.
<原文结束>

# <翻译开始>
	// 改变到临时目录，并构造一个相对于工作目录的符号链接。
# <翻译结束>


<原文开始>
// no backslash after volume.
<原文结束>

# <翻译开始>
// 体积后不跟反斜杠。
# <翻译结束>


<原文开始>
	// Now change back to the original working directory and verify that the
	// symlink still refers to its original path and is correctly marked as a
	// directory.
<原文结束>

# <翻译开始>
	// 现在切换回原始的工作目录，并验证该符号链接仍指向其原始路径，且正确标记为目录。
# <翻译结束>


<原文开始>
// TestStatOfInvalidName is regression test for issue #24999.
<原文结束>

# <翻译开始>
// TestStatOfInvalidName 是针对问题 #24999 的回归测试。
# <翻译结束>


<原文开始>
// findUnusedDriveLetter searches mounted drive list on the system
// (starting from Z: and ending at D:) for unused drive letter.
// It returns path to the found drive root directory (like Z:\) or error.
<原文结束>

# <翻译开始>
// findUnusedDriveLetter 在系统上搜索已挂载的驱动器列表（从 Z: 开始，到 D: 结束），寻找未使用的驱动器字母。它返回找到的驱动器根目录的路径（如 Z:\）或错误。
# <翻译结束>


<原文开始>
	// Do not use A: and B:, because they are reserved for floppy drive.
	// Do not use C:, because it is normally used for main drive.
<原文结束>

# <翻译开始>
	// 不要使用A:和B:，因为它们被保留给了软盘驱动。
	// 不要使用C:，因为它通常用于主驱动。
# <翻译结束>


<原文开始>
// Make sure tmpdir is not a symlink, otherwise tests will fail.
<原文结束>

# <翻译开始>
// 确保 tmpdir 不是符号链接，否则测试将会失败。
# <翻译结束>


<原文开始>
// relative directory junction resolves to absolute path
<原文结束>

# <翻译开始>
// 相对目录连接解析为绝对路径
# <翻译结束>


<原文开始>
// Make sure we have sufficient privilege to run mklink command.
<原文结束>

# <翻译开始>
// 确保我们有足够的权限来运行mklink命令。
# <翻译结束>


<原文开始>
	// Check opened directories can't be renamed until the handle is closed.
	// See issue 52747.
<原文结束>

# <翻译开始>
	// 检查打开的目录在句柄关闭之前不能重命名。参见问题52747。
# <翻译结束>


<原文开始>
	// We expect executables installed to %LOCALAPPDATA%\Microsoft\WindowsApps to
	// be reparse points with tag IO_REPARSE_TAG_APPEXECLINK. Here we check that
	// such reparse points are treated as irregular (but executable) files, not
	// broken symlinks.
<原文结束>

# <翻译开始>
	// 我们期望安装在 %LOCALAPPDATA%\Microsoft\WindowsApps 中的可执行文件
	// 是带有 IO_REPARSE_TAG_APPEXECLINK 标签的重分析点。此处我们检查这些重分析点
	// 是否被视为非标准（但可执行）文件，而非损坏的符号链接。
# <翻译结束>


<原文开始>
	// An APPEXECLINK reparse point is not a symlink, so os.Readlink should return
	// a non-nil error for it, and Stat should return results identical to Lstat.
<原文结束>

# <翻译开始>
	// APPEXECLINK 重解析点不是符号链接，因此 os.Readlink 应该为其返回非空错误，而 Stat 应该返回与 Lstat 相同的结果。
# <翻译结束>


<原文开始>
		// A reparse point is not a regular file, but we don't have a more appropriate
		// ModeType bit for it, so it should be marked as irregular.
<原文结束>

# <翻译开始>
		// 重定向点不是常规文件，但我们没有更合适的ModeType位来标记它，所以应该将其标记为不规则的。
# <翻译结束>


<原文开始>
// This UTF-16 file name is ill-formed as it contains low surrogates that are not preceded by high surrogates ([1:5]).
<原文结束>

# <翻译开始>
// 这个UTF-16文件名格式不正确，因为它包含没有被高-surrogate（[1:5]）前导的低-surrogates。
# <翻译结束>


<原文开始>
	// Create a file whose name contains unpaired surrogates.
	// Use syscall.CreateFile instead of os.Create to simulate a file that is created by
	// a non-Go program so the file name hasn't gone through syscall.UTF16FromString.
<原文结束>

# <翻译开始>
	// 创建一个文件名包含非配对代理项的文件。
	// 使用syscall.CreateFile 而非 os.Create，以模拟由非Go程序创建的文件，
	// 使得文件名未经 syscall.UTF16FromString 处理。
# <翻译结束>


<原文开始>
// Verify that os.Lstat can query the file.
<原文结束>

# <翻译开始>
// 验证os.Lstat能够查询文件。
# <翻译结束>


<原文开始>
// Verify that File.Readdirnames lists the file.
<原文结束>

# <翻译开始>
// 验证File.Readdirnames列出的文件。
# <翻译结束>


<原文开始>
	// Verify that os.RemoveAll can remove the directory
	// and that it doesn't hang.
<原文结束>

# <翻译开始>
	// 验证os.RemoveAll是否可以删除目录，以及它是否会挂起。
# <翻译结束>


<原文开始>
// Check that os.SameFile works with files returned by os.ReadDir.
<原文结束>

# <翻译开始>
// 检查 os.SameFile 函数是否适用于由 os.ReadDir 返回的文件。
# <翻译结束>


<原文开始>
// Check that os.SameFile works with a mix of os.ReadDir and os.Stat files.
<原文结束>

# <翻译开始>
// 检查os.SameFile在处理os.ReadDir和os.Stat获取的文件时是否能正常工作。
# <翻译结束>

