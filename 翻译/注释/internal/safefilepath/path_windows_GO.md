
<原文开始>
// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
# <翻译结束>


<原文开始>
// Find the next path element.
<原文结束>

# <翻译开始>
// Find the next path element.
# <翻译结束>


<原文开始>
// We can't depend on strings, so substitute \ for / manually.
<原文结束>

# <翻译开始>
// We can't depend on strings, so substitute \ for / manually.
# <翻译结束>


<原文开始>
// IsReservedName reports if name is a Windows reserved device name.
// It does not detect names with an extension, which are also reserved on some Windows versions.
//
// For details, search for PRN in
// https://docs.microsoft.com/en-us/windows/desktop/fileio/naming-a-file.
<原文结束>

# <翻译开始>
// IsReservedName reports if name is a Windows reserved device name.
// It does not detect names with an extension, which are also reserved on some Windows versions.
//
// For details, search for PRN in
// https://docs.microsoft.com/en-us/windows/desktop/fileio/naming-a-file.
# <翻译结束>


<原文开始>
// Device names can have arbitrary trailing characters following a dot or colon.
<原文结束>

# <翻译开始>
// Device names can have arbitrary trailing characters following a dot or colon.
# <翻译结束>


<原文开始>
// Trailing spaces in the last path element are ignored.
<原文结束>

# <翻译开始>
// Trailing spaces in the last path element are ignored.
# <翻译结束>


<原文开始>
	// The path element is a reserved name with an extension.
	// Some Windows versions consider this a reserved name,
	// while others do not. Use FullPath to see if the name is
	// reserved.
<原文结束>

# <翻译开始>
	// The path element is a reserved name with an extension.
	// Some Windows versions consider this a reserved name,
	// while others do not. Use FullPath to see if the name is
	// reserved.
# <翻译结束>


<原文开始>
	// Passing CONIN$ or CONOUT$ to CreateFile opens a console handle.
	// https://learn.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-createfilea#consoles
	//
	// While CONIN$ and CONOUT$ aren't documented as being files,
	// they behave the same as CON. For example, ./CONIN$ also opens the console input.
<原文结束>

# <翻译开始>
	// Passing CONIN$ or CONOUT$ to CreateFile opens a console handle.
	// https://learn.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-createfilea#consoles
	//
	// While CONIN$ and CONOUT$ aren't documented as being files,
	// they behave the same as CON. For example, ./CONIN$ also opens the console input.
# <翻译结束>

