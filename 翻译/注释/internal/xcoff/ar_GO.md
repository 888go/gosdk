
<原文开始>
// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
# <翻译结束>


<原文开始>
// 32-bits global symtab offset
<原文结束>

# <翻译开始>
// 32-bits global symtab offset
# <翻译结束>


<原文开始>
// 64-bits global symtab offset
<原文结束>

# <翻译开始>
// 64-bits global symtab offset
# <翻译结束>


<原文开始>
// First member on free list offset
<原文结束>

# <翻译开始>
// First member on free list offset
# <翻译结束>


<原文开始>
// Previous member pointer
<原文结束>

# <翻译开始>
// Previous member pointer
# <翻译结束>


<原文开始>
// File member mode (octal)
<原文结束>

# <翻译开始>
// File member mode (octal)
# <翻译结束>


<原文开始>
// File member name length
<原文结束>

# <翻译开始>
// File member name length
# <翻译结束>


<原文开始>
// _ar_nam is removed because it's easier to get name without it.
<原文结束>

# <翻译开始>
// _ar_nam is removed because it's easier to get name without it.
# <翻译结束>


<原文开始>
// Archive represents an open AIX big archive.
<原文结束>

# <翻译开始>
// Archive represents an open AIX big archive.
# <翻译结束>


<原文开始>
// MemberHeader holds information about a big archive file header
<原文结束>

# <翻译开始>
// MemberHeader holds information about a big archive file header
# <翻译结束>


<原文开始>
// Member represents a member of an AIX big archive.
<原文结束>

# <翻译开始>
// Member represents a member of an AIX big archive.
# <翻译结束>


<原文开始>
// MemberHeader holds information about a big archive member
<原文结束>

# <翻译开始>
// MemberHeader holds information about a big archive member
# <翻译结束>


<原文开始>
// OpenArchive opens the named archive using os.Open and prepares it for use
// as an AIX big archive.
<原文结束>

# <翻译开始>
// OpenArchive opens the named archive using os.Open and prepares it for use
// as an AIX big archive.
# <翻译结束>


<原文开始>
// Close closes the Archive.
// If the Archive was created using NewArchive directly instead of OpenArchive,
// Close has no effect.
<原文结束>

# <翻译开始>
// Close closes the Archive.
// If the Archive was created using NewArchive directly instead of OpenArchive,
// Close has no effect.
# <翻译结束>


<原文开始>
// NewArchive creates a new Archive for accessing an AIX big archive in an underlying reader.
<原文结束>

# <翻译开始>
// NewArchive creates a new Archive for accessing an AIX big archive in an underlying reader.
# <翻译结束>


<原文开始>
// Occurs if the archive is empty.
<原文结束>

# <翻译开始>
// Occurs if the archive is empty.
# <翻译结束>


<原文开始>
		// Read Member Header
		// The member header is normally 2 bytes larger. But it's easier
		// to read the name if the header is read without _ar_nam.
		// However, AIAFMAG must be read afterward.
<原文结束>

# <翻译开始>
		// Read Member Header
		// The member header is normally 2 bytes larger. But it's easier
		// to read the name if the header is read without _ar_nam.
		// However, AIAFMAG must be read afterward.
# <翻译结束>


<原文开始>
// Add the two bytes of AIAFMAG
<原文结束>

# <翻译开始>
// Add the two bytes of AIAFMAG
# <翻译结束>


<原文开始>
// GetFile returns the XCOFF file defined by member name.
// FIXME: This doesn't work if an archive has two members with the same
// name which can occur if a archive has both 32-bits and 64-bits files.
<原文结束>

# <翻译开始>
// GetFile returns the XCOFF file defined by member name.
// FIXME: This doesn't work if an archive has two members with the same
// name which can occur if a archive has both 32-bits and 64-bits files.
# <翻译结束>

