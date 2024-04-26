// 版权所有 2011 Go 作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可协议约束，
// 该协议可在 LICENSE 文件中找到。

package os

import (
	"errors"
	"syscall"
	"time"
)

//Export for testing.

var lstat = Lstat
var errWriteAtInAppendMode = errors.New("os: invalid use of WriteAt on file opened with O_APPEND")

// testingForceReadDirLstat 强制 ReadDir 调用 Lstat，用于测试该代码路径。在某些 Unix 系统上，这可能很难触发。
var testingForceReadDirLstat bool
var errPatternHasSeparator = errors.New("pattern contains path separator")

// checkWrapErr 是用于检测 poll.ErrFileClosing 不期望的包装错误的测试钩子。
// 在 export_test.go 中将其设置为 true，以供测试（包括模糊测试）使用。
var checkWrapErr = false

// For testing.
func atime(fi FileInfo) time.Time {
	return time.Unix(0, fi.Sys().(*syscall.Win32FileAttributeData).LastAccessTime.Nanoseconds())
}

var Atime = atime
var LstatP = &lstat
var ErrWriteAtInAppendMode = errWriteAtInAppendMode
var TestingForceReadDirLstat = &testingForceReadDirLstat
var ErrPatternHasSeparator = errPatternHasSeparator

func init() {
	checkWrapErr = true
}
