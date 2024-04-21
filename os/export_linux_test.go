// 版权所有 2020 Go作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

package os

var (
	PollCopyFileRangeP  = &pollCopyFileRange
	PollSpliceFile      = &pollSplice
	PollSendFile        = &pollSendFile
	GetPollFDAndNetwork = getPollFDAndNetwork
)
