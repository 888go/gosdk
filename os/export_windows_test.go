// 版权所有 2016 The Go 作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

package os

// 用于测试导出。
// 此值通过go:linkname与runtime.canUseLongPaths关联设置，并在操作系统支持选择正确处理长路径而无需修复时为真。
var canUseLongPaths bool

// allowReadDirFileID 表示如果底层文件系统支持，File.readdir 是否应该尝试使用 FILE_ID_BOTH_DIR_INFO。
// 适用于测试目的。
var allowReadDirFileID = true
var (
	// FixLongPath：修复长路径
	CanUseLongPaths = canUseLongPaths

	// CommandLineToArgv 等于 commandLineToArgv
	AllowReadDirFileID = &allowReadDirFileID
)
