//版权所有2016 The Go Authors。所有权利保留。
//使用此源代码受BSD风格的
//可在LICENSE文件中找到的许可协议管辖。

package windows

import (
	"syscall"
	"unsafe"
)

const (
	FSCTL_SET_REPARSE_POINT    = 0x000900A4
	IO_REPARSE_TAG_MOUNT_POINT = 0xA0000003
	IO_REPARSE_TAG_DEDUP       = 0x80000013

	SYMLINK_FLAG_RELATIVE = 1
)

// 这些结构在以下文档中进行描述：
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-fscc/ca069dad-ed16-42aa-b057-b6b207f447cc
// 和
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-fscc/b41f1cbf-10df-4a47-98d4-1c52a833d913。

type REPARSE_DATA_BUFFER struct {
	ReparseTag        uint32
	ReparseDataLength uint16
	Reserved          uint16
	DUMMYUNIONNAME    byte
}

// REPARSE_DATA_BUFFER_HEADER 是 REPARSE_DATA_BUFFER 结构体的公共部分。
type REPARSE_DATA_BUFFER_HEADER struct {
	ReparseTag uint32
// 重新解析数据的大小，以字节为单位，紧跟在
// REPARSE_DATA_BUFFER 元素的公共部分后面。
// 这个值是从 SubstituteNameOffset 字段开始的数据长度。
	ReparseDataLength uint16
	Reserved          uint16
}

type SymbolicLinkReparseBuffer struct {
// 包含替换名称字符串在PathBuffer数组中的偏移量（以字节为单位），它是从PathBuffer的第一个字节开始计算的。需要注意的是，这个偏移量需要除以2来得到数组的索引。
	SubstituteNameOffset uint16
// 包含替换名称字符串的字节数的整数。如果此字符串以空字符终止，SubstituteNameLength 不包括 Unicode 空字符。
	SubstituteNameLength uint16
	// PrintNameOffset 与 SubstituteNameOffset 类似
	PrintNameOffset uint16
	// PrintNameLength 与 SubstituteNameLength 相似。
	PrintNameLength uint16
// Flags 指定替换名称是完整路径名还是相对于符号链接所在目录的路径名。
	Flags      uint32
	PathBuffer [1]uint16
}

// Path 返回 rb 中存储的路径。
func (rb *SymbolicLinkReparseBuffer) Path() string {
	n1 := rb.SubstituteNameOffset / 2
	n2 := (rb.SubstituteNameOffset + rb.SubstituteNameLength) / 2
	return syscall.UTF16ToString((*[0xffff]uint16)(unsafe.Pointer(&rb.PathBuffer[0]))[n1:n2:n2])
}

type MountPointReparseBuffer struct {
// 包含替换名称字符串在PathBuffer数组中的偏移量（以字节为单位），它是从PathBuffer的第一个字节开始计算的。需要注意的是，这个偏移量需要除以2来得到数组的索引。
	SubstituteNameOffset uint16
// 包含替换名称字符串的字节数的整数。如果此字符串以空字符终止，SubstituteNameLength 不包括 Unicode 空字符。
	SubstituteNameLength uint16
	// PrintNameOffset 与 SubstituteNameOffset 类似
	PrintNameOffset uint16
	// PrintNameLength 与 SubstituteNameLength 相似。
	PrintNameLength uint16
	PathBuffer      [1]uint16
}

// Path 返回 rb 中存储的路径。
func (rb *MountPointReparseBuffer) Path() string {
	n1 := rb.SubstituteNameOffset / 2
	n2 := (rb.SubstituteNameOffset + rb.SubstituteNameLength) / 2
	return syscall.UTF16ToString((*[0xffff]uint16)(unsafe.Pointer(&rb.PathBuffer[0]))[n1:n2:n2])
}
