
<原文开始>
// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
# <翻译结束>


<原文开始>
// ErrShortBuffer is returned when the buffer was too short for the operation.
<原文结束>

# <翻译开始>
// ErrShortBuffer is returned when the buffer was too short for the operation.
# <翻译结束>


<原文开始>
// ErrNotExist is returned when a registry key or value does not exist.
<原文结束>

# <翻译开始>
// ErrNotExist is returned when a registry key or value does not exist.
# <翻译结束>


<原文开始>
// ErrUnexpectedType is returned by Get*Value when the value's type was unexpected.
<原文结束>

# <翻译开始>
// ErrUnexpectedType is returned by Get*Value when the value's type was unexpected.
# <翻译结束>


<原文开始>
// GetValue retrieves the type and data for the specified value associated
// with an open key k. It fills up buffer buf and returns the retrieved
// byte count n. If buf is too small to fit the stored value it returns
// ErrShortBuffer error along with the required buffer size n.
// If no buffer is provided, it returns true and actual buffer size n.
// If no buffer is provided, GetValue returns the value's type only.
// If the value does not exist, the error returned is ErrNotExist.
//
// GetValue is a low level function. If value's type is known, use the appropriate
// Get*Value function instead.
<原文结束>

# <翻译开始>
// GetValue retrieves the type and data for the specified value associated
// with an open key k. It fills up buffer buf and returns the retrieved
// byte count n. If buf is too small to fit the stored value it returns
// ErrShortBuffer error along with the required buffer size n.
// If no buffer is provided, it returns true and actual buffer size n.
// If no buffer is provided, GetValue returns the value's type only.
// If the value does not exist, the error returned is ErrNotExist.
//
// GetValue is a low level function. If value's type is known, use the appropriate
// Get*Value function instead.
# <翻译结束>


<原文开始>
// GetStringValue retrieves the string value for the specified
// value name associated with an open key k. It also returns the value's type.
// If value does not exist, GetStringValue returns ErrNotExist.
// If value is not SZ or EXPAND_SZ, it will return the correct value
// type and ErrUnexpectedType.
<原文结束>

# <翻译开始>
// GetStringValue retrieves the string value for the specified
// value name associated with an open key k. It also returns the value's type.
// If value does not exist, GetStringValue returns ErrNotExist.
// If value is not SZ or EXPAND_SZ, it will return the correct value
// type and ErrUnexpectedType.
# <翻译结束>


<原文开始>
// GetMUIStringValue retrieves the localized string value for
// the specified value name associated with an open key k.
// If the value name doesn't exist or the localized string value
// can't be resolved, GetMUIStringValue returns ErrNotExist.
// GetMUIStringValue panics if the system doesn't support
// regLoadMUIString; use LoadRegLoadMUIString to check if
// regLoadMUIString is supported before calling this function.
<原文结束>

# <翻译开始>
// GetMUIStringValue retrieves the localized string value for
// the specified value name associated with an open key k.
// If the value name doesn't exist or the localized string value
// can't be resolved, GetMUIStringValue returns ErrNotExist.
// GetMUIStringValue panics if the system doesn't support
// regLoadMUIString; use LoadRegLoadMUIString to check if
// regLoadMUIString is supported before calling this function.
# <翻译结束>


<原文开始>
		// Try to resolve the string value using the system directory as
		// a DLL search path; this assumes the string value is of the form
		// @[path]\dllname,-strID but with no path given, e.g. @tzres.dll,-320.
<原文结束>

# <翻译开始>
		// Try to resolve the string value using the system directory as
		// a DLL search path; this assumes the string value is of the form
		// @[path]\dllname,-strID but with no path given, e.g. @tzres.dll,-320.
# <翻译结束>


<原文开始>
		// This approach works with tzres.dll but may have to be revised
		// in the future to allow callers to provide custom search paths.
<原文结束>

# <翻译开始>
		// This approach works with tzres.dll but may have to be revised
		// in the future to allow callers to provide custom search paths.
# <翻译结束>


<原文开始>
// Buffer not growing, assume race; break
<原文结束>

# <翻译开始>
// Buffer not growing, assume race; break
# <翻译结束>


<原文开始>
// ExpandString expands environment-variable strings and replaces
// them with the values defined for the current user.
// Use ExpandString to expand EXPAND_SZ strings.
<原文结束>

# <翻译开始>
// ExpandString expands environment-variable strings and replaces
// them with the values defined for the current user.
// Use ExpandString to expand EXPAND_SZ strings.
# <翻译结束>


<原文开始>
// GetStringsValue retrieves the []string value for the specified
// value name associated with an open key k. It also returns the value's type.
// If value does not exist, GetStringsValue returns ErrNotExist.
// If value is not MULTI_SZ, it will return the correct value
// type and ErrUnexpectedType.
<原文结束>

# <翻译开始>
// GetStringsValue retrieves the []string value for the specified
// value name associated with an open key k. It also returns the value's type.
// If value does not exist, GetStringsValue returns ErrNotExist.
// If value is not MULTI_SZ, it will return the correct value
// type and ErrUnexpectedType.
# <翻译结束>


<原文开始>
// remove terminating null
<原文结束>

# <翻译开始>
// remove terminating null
# <翻译结束>


<原文开始>
// GetIntegerValue retrieves the integer value for the specified
// value name associated with an open key k. It also returns the value's type.
// If value does not exist, GetIntegerValue returns ErrNotExist.
// If value is not DWORD or QWORD, it will return the correct value
// type and ErrUnexpectedType.
<原文结束>

# <翻译开始>
// GetIntegerValue retrieves the integer value for the specified
// value name associated with an open key k. It also returns the value's type.
// If value does not exist, GetIntegerValue returns ErrNotExist.
// If value is not DWORD or QWORD, it will return the correct value
// type and ErrUnexpectedType.
# <翻译结束>


<原文开始>
// GetBinaryValue retrieves the binary value for the specified
// value name associated with an open key k. It also returns the value's type.
// If value does not exist, GetBinaryValue returns ErrNotExist.
// If value is not BINARY, it will return the correct value
// type and ErrUnexpectedType.
<原文结束>

# <翻译开始>
// GetBinaryValue retrieves the binary value for the specified
// value name associated with an open key k. It also returns the value's type.
// If value does not exist, GetBinaryValue returns ErrNotExist.
// If value is not BINARY, it will return the correct value
// type and ErrUnexpectedType.
# <翻译结束>


<原文开始>
// SetDWordValue sets the data and type of a name value
// under key k to value and DWORD.
<原文结束>

# <翻译开始>
// SetDWordValue sets the data and type of a name value
// under key k to value and DWORD.
# <翻译结束>


<原文开始>
// SetQWordValue sets the data and type of a name value
// under key k to value and QWORD.
<原文结束>

# <翻译开始>
// SetQWordValue sets the data and type of a name value
// under key k to value and QWORD.
# <翻译结束>


<原文开始>
// SetStringValue sets the data and type of a name value
// under key k to value and SZ. The value must not contain a zero byte.
<原文结束>

# <翻译开始>
// SetStringValue sets the data and type of a name value
// under key k to value and SZ. The value must not contain a zero byte.
# <翻译结束>


<原文开始>
// SetExpandStringValue sets the data and type of a name value
// under key k to value and EXPAND_SZ. The value must not contain a zero byte.
<原文结束>

# <翻译开始>
// SetExpandStringValue sets the data and type of a name value
// under key k to value and EXPAND_SZ. The value must not contain a zero byte.
# <翻译结束>


<原文开始>
// SetStringsValue sets the data and type of a name value
// under key k to value and MULTI_SZ. The value strings
// must not contain a zero byte.
<原文结束>

# <翻译开始>
// SetStringsValue sets the data and type of a name value
// under key k to value and MULTI_SZ. The value strings
// must not contain a zero byte.
# <翻译结束>


<原文开始>
// SetBinaryValue sets the data and type of a name value
// under key k to value and BINARY.
<原文结束>

# <翻译开始>
// SetBinaryValue sets the data and type of a name value
// under key k to value and BINARY.
# <翻译结束>


<原文开始>
// DeleteValue removes a named value from the key k.
<原文结束>

# <翻译开始>
// DeleteValue removes a named value from the key k.
# <翻译结束>


<原文开始>
// ReadValueNames returns the value names of key k.
// The parameter n controls the number of returned names,
// analogous to the way os.File.Readdirnames works.
<原文结束>

# <翻译开始>
// ReadValueNames returns the value names of key k.
// The parameter n controls the number of returned names,
// analogous to the way os.File.Readdirnames works.
# <翻译结束>


<原文开始>
// extra room for terminating null character
<原文结束>

# <翻译开始>
// extra room for terminating null character
# <翻译结束>


<原文开始>
// Double buffer size and try again.
<原文结束>

# <翻译开始>
// Double buffer size and try again.
# <翻译结束>

