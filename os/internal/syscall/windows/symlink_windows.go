// 版权所有 2018 The Go 作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

package windows

import "syscall"

const (
	ERROR_INVALID_PARAMETER syscall.Errno = 87

	FILE_SUPPORTS_OPEN_BY_FILE_ID = 0x01000000

	// 对于 CreateSymbolicLink()，自 Windows 10（版本 1703，内部版本 10.0.14972）起支持符号链接
	SYMBOLIC_LINK_FLAG_ALLOW_UNPRIVILEGED_CREATE = 0x2

	// FileInformationClass 值
	FileBasicInfo                  = 0    // FILE_BASIC_INFO
	FileStandardInfo               = 1    // FILE_STANDARD_INFO
	FileNameInfo                   = 2    // FILE_NAME_INFO
	FileStreamInfo                 = 7    // FILE_STREAM_INFO
	FileCompressionInfo            = 8    // FILE_COMPRESSION_INFO
	FileAttributeTagInfo           = 9    // FILE_ATTRIBUTE_TAG_INFO：文件属性标签信息
	FileIdBothDirectoryInfo        = 0xa  // FILE_ID_BOTH_DIR_INFO
	FileIdBothDirectoryRestartInfo = 0xb  // FILE_ID_BOTH_DIR_INFO
	FileRemoteProtocolInfo         = 0xd  // FILE_REMOTE_PROTOCOL_INFO：文件远程协议信息. md5:296ef5f8110a634b
	FileFullDirectoryInfo          = 0xe  // FILE_FULL_DIR_INFO
	FileFullDirectoryRestartInfo   = 0xf  // FILE_FULL_DIR_INFO
	FileStorageInfo                = 0x10 // FILE_STORAGE_INFO
	FileAlignmentInfo              = 0x11 // FILE_ALIGNMENT_INFO
	FileIdInfo                     = 0x12 // FILE_ID_INFO
	FileIdExtdDirectoryInfo        = 0x13 // FILE_ID_EXTD_DIR_INFO
	FileIdExtdDirectoryRestartInfo = 0x14 // FILE_ID_EXTD_DIR_INFO
)

type FILE_ATTRIBUTE_TAG_INFO struct {
	FileAttributes uint32
	ReparseTag     uint32
}

//sys	GetFileInformationByHandleEx(handle syscall.Handle, class uint32, info *byte, bufsize uint32) (err error)
