// 版权所有 2011 Go 作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可协议约束，
// 该协议可在 LICENSE 文件中找到。

package os

// Export for testing.

var Atime = atime
var LstatP = &lstat
var ErrWriteAtInAppendMode = errWriteAtInAppendMode
var TestingForceReadDirLstat = &testingForceReadDirLstat
var ErrPatternHasSeparator = errPatternHasSeparator

func init() {
	checkWrapErr = true
}
