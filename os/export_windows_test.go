// 版权所有 2016 The Go 作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

package os

// Export for testing.

var (
	FixLongPath        = fixLongPath
	CanUseLongPaths    = canUseLongPaths
	NewConsoleFile     = newConsoleFile
	CommandLineToArgv  = commandLineToArgv
	AllowReadDirFileID = &allowReadDirFileID
)
