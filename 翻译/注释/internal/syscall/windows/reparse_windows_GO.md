
<原文开始>
// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
# <翻译结束>


<原文开始>
// These structures are described
// in https://msdn.microsoft.com/en-us/library/cc232007.aspx
// and https://msdn.microsoft.com/en-us/library/cc232006.aspx.
<原文结束>

# <翻译开始>
// These structures are described
// in https://msdn.microsoft.com/en-us/library/cc232007.aspx
// and https://msdn.microsoft.com/en-us/library/cc232006.aspx.
# <翻译结束>


<原文开始>
// REPARSE_DATA_BUFFER_HEADER is a common part of REPARSE_DATA_BUFFER structure.
<原文结束>

# <翻译开始>
// REPARSE_DATA_BUFFER_HEADER is a common part of REPARSE_DATA_BUFFER structure.
# <翻译结束>


<原文开始>
	// The size, in bytes, of the reparse data that follows
	// the common portion of the REPARSE_DATA_BUFFER element.
	// This value is the length of the data starting at the
	// SubstituteNameOffset field.
<原文结束>

# <翻译开始>
	// The size, in bytes, of the reparse data that follows
	// the common portion of the REPARSE_DATA_BUFFER element.
	// This value is the length of the data starting at the
	// SubstituteNameOffset field.
# <翻译结束>


<原文开始>
	// The integer that contains the offset, in bytes,
	// of the substitute name string in the PathBuffer array,
	// computed as an offset from byte 0 of PathBuffer. Note that
	// this offset must be divided by 2 to get the array index.
<原文结束>

# <翻译开始>
	// The integer that contains the offset, in bytes,
	// of the substitute name string in the PathBuffer array,
	// computed as an offset from byte 0 of PathBuffer. Note that
	// this offset must be divided by 2 to get the array index.
# <翻译结束>


<原文开始>
	// The integer that contains the length, in bytes, of the
	// substitute name string. If this string is null-terminated,
	// SubstituteNameLength does not include the Unicode null character.
<原文结束>

# <翻译开始>
	// The integer that contains the length, in bytes, of the
	// substitute name string. If this string is null-terminated,
	// SubstituteNameLength does not include the Unicode null character.
# <翻译结束>


<原文开始>
// PrintNameOffset is similar to SubstituteNameOffset.
<原文结束>

# <翻译开始>
// PrintNameOffset is similar to SubstituteNameOffset.
# <翻译结束>


<原文开始>
// PrintNameLength is similar to SubstituteNameLength.
<原文结束>

# <翻译开始>
// PrintNameLength is similar to SubstituteNameLength.
# <翻译结束>


<原文开始>
	// Flags specifies whether the substitute name is a full path name or
	// a path name relative to the directory containing the symbolic link.
<原文结束>

# <翻译开始>
	// Flags specifies whether the substitute name is a full path name or
	// a path name relative to the directory containing the symbolic link.
# <翻译结束>


<原文开始>
// Path returns path stored in rb.
<原文结束>

# <翻译开始>
// Path returns path stored in rb.
# <翻译结束>

