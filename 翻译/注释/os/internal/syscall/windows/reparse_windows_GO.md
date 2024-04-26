
<原文开始>
// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
//版权所有2016 The Go Authors。所有权利保留。
//使用此源代码受BSD风格的
//可在LICENSE文件中找到的许可协议管辖。
# <翻译结束>


<原文开始>
// These structures are described
// in https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-fscc/ca069dad-ed16-42aa-b057-b6b207f447cc
// and https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-fscc/b41f1cbf-10df-4a47-98d4-1c52a833d913.
<原文结束>

# <翻译开始>
// 这些结构在以下文档中进行描述：
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-fscc/ca069dad-ed16-42aa-b057-b6b207f447cc
// 和
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-fscc/b41f1cbf-10df-4a47-98d4-1c52a833d913。
# <翻译结束>


<原文开始>
// REPARSE_DATA_BUFFER_HEADER is a common part of REPARSE_DATA_BUFFER structure.
<原文结束>

# <翻译开始>
// REPARSE_DATA_BUFFER_HEADER 是 REPARSE_DATA_BUFFER 结构体的公共部分。
# <翻译结束>


<原文开始>
	// The size, in bytes, of the reparse data that follows
	// the common portion of the REPARSE_DATA_BUFFER element.
	// This value is the length of the data starting at the
	// SubstituteNameOffset field.
<原文结束>

# <翻译开始>
// 重新解析数据的大小，以字节为单位，紧跟在
// REPARSE_DATA_BUFFER 元素的公共部分后面。
// 这个值是从 SubstituteNameOffset 字段开始的数据长度。
# <翻译结束>


<原文开始>
	// The integer that contains the offset, in bytes,
	// of the substitute name string in the PathBuffer array,
	// computed as an offset from byte 0 of PathBuffer. Note that
	// this offset must be divided by 2 to get the array index.
<原文结束>

# <翻译开始>
// 包含替换名称字符串在PathBuffer数组中的偏移量（以字节为单位），它是从PathBuffer的第一个字节开始计算的。需要注意的是，这个偏移量需要除以2来得到数组的索引。
# <翻译结束>


<原文开始>
	// The integer that contains the length, in bytes, of the
	// substitute name string. If this string is null-terminated,
	// SubstituteNameLength does not include the Unicode null character.
<原文结束>

# <翻译开始>
// 包含替换名称字符串的字节数的整数。如果此字符串以空字符终止，SubstituteNameLength 不包括 Unicode 空字符。
# <翻译结束>


<原文开始>
// PrintNameOffset is similar to SubstituteNameOffset.
<原文结束>

# <翻译开始>
// PrintNameOffset 与 SubstituteNameOffset 类似
# <翻译结束>


<原文开始>
// PrintNameLength is similar to SubstituteNameLength.
<原文结束>

# <翻译开始>
// PrintNameLength 与 SubstituteNameLength 相似。
# <翻译结束>


<原文开始>
	// Flags specifies whether the substitute name is a full path name or
	// a path name relative to the directory containing the symbolic link.
<原文结束>

# <翻译开始>
// Flags 指定替换名称是完整路径名还是相对于符号链接所在目录的路径名。
# <翻译结束>


<原文开始>
// Path returns path stored in rb.
<原文结束>

# <翻译开始>
// Path 返回 rb 中存储的路径。
# <翻译结束>

