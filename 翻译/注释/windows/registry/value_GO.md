
<原文开始>
// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// 版权所有 ? 2015 The Go 作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。
# <翻译结束>


<原文开始>
// ErrShortBuffer is returned when the buffer was too short for the operation.
<原文结束>

# <翻译开始>
// ErrShortBuffer 表示当缓冲区对于操作而言过短时返回的错误。
# <翻译结束>


<原文开始>
// ErrNotExist is returned when a registry key or value does not exist.
<原文结束>

# <翻译开始>
// ErrNotExist 表示当注册表键或值不存在时返回的错误。
# <翻译结束>


<原文开始>
// ErrUnexpectedType is returned by Get*Value when the value's type was unexpected.
<原文结束>

# <翻译开始>
// ErrUnexpectedType 会在获取值时遇到类型不符预期的情况时返回。
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
// GetValue 用于从与已打开键 k 关联的指定值中检索类型和数据。
// 它将数据填充到缓冲区 buf 中，并返回获取的字节数 n。
// 如果 buf 太小无法容纳存储的值，它将返回 ErrShortBuffer 错误以及所需的缓冲区大小 n。
// 如果未提供缓冲区，它将返回 true 以及实际缓冲区大小 n。
// 如果未提供缓冲区，GetValue 仅返回值的类型。
// 若该值不存在，返回的错误为 ErrNotExist。
//
// GetValue 是一个低级函数。若已知值的类型，请使用相应的 Get*Value 函数代替。
# <翻译结束>


<原文开始>
// GetStringValue retrieves the string value for the specified
// value name associated with an open key k. It also returns the value's type.
// If value does not exist, GetStringValue returns ErrNotExist.
// If value is not SZ or EXPAND_SZ, it will return the correct value
// type and ErrUnexpectedType.
<原文结束>

# <翻译开始>
// GetStringValue 用于获取与已打开键 k 关联的指定值名的字符串值。同时返回该值的类型。
// 若该值不存在，GetStringValue 将返回 ErrNotExist 错误。
// 若该值非 SZ 或 EXPAND_SZ 类型，它将返回正确的值类型以及 ErrUnexpectedType 错误。
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
// GetMUIStringValue 用于获取与已打开键 k 关联的指定值名所对应的本地化字符串值。
// 若该值名不存在，或无法解析其本地化字符串值，则 GetMUIStringValue 返回 ErrNotExist。
// 若系统不支持 regLoadMUIString，GetMUIStringValue 将引发恐慌。因此在调用此函数前，请使用 LoadRegLoadMUIString 检查系统是否支持 regLoadMUIString。
# <翻译结束>


<原文开始>
		// Try to resolve the string value using the system directory as
		// a DLL search path; this assumes the string value is of the form
		// @[path]\dllname,-strID but with no path given, e.g. @tzres.dll,-320.
<原文结束>

# <翻译开始>
// 尝试使用系统目录作为 DLL 搜索路径解析字符串值；
// 这种方式假设字符串值形如 @[路径]\dllname,-strID，但未给出具体路径，
// 例如：@tzres.dll,-320
# <翻译结束>


<原文开始>
		// This approach works with tzres.dll but may have to be revised
		// in the future to allow callers to provide custom search paths.
<原文结束>

# <翻译开始>
// 此方法适用于tzres.dll，但未来可能需要进行修订，
// 以便调用者能够提供自定义搜索路径。
# <翻译结束>


<原文开始>
// Buffer not growing, assume race; break
<原文结束>

# <翻译开始>
// 缓冲区未增长，假设存在竞争；中断
# <翻译结束>


<原文开始>
// ExpandString expands environment-variable strings and replaces
// them with the values defined for the current user.
// Use ExpandString to expand EXPAND_SZ strings.
<原文结束>

# <翻译开始>
// ExpandString 将环境变量字符串展开，并用当前用户所定义的值进行替换。
// 对于 EXPAND_SZ 类型的字符串，应使用 ExpandString 进行展开。
# <翻译结束>


<原文开始>
// GetStringsValue retrieves the []string value for the specified
// value name associated with an open key k. It also returns the value's type.
// If value does not exist, GetStringsValue returns ErrNotExist.
// If value is not MULTI_SZ, it will return the correct value
// type and ErrUnexpectedType.
<原文结束>

# <翻译开始>
// GetStringsValue 用于从与已打开键 k 关联的指定值名中检索 []string 值。同时返回该值的类型。
// 若值不存在，GetStringsValue 返回 ErrNotExist 错误。
// 若值并非 MULTI_SZ 类型，它将返回正确的值类型及 ErrUnexpectedType 错误。
# <翻译结束>


<原文开始>
// remove terminating null
<原文结束>

# <翻译开始>
// 移除终止空字符
# <翻译结束>


<原文开始>
// GetIntegerValue retrieves the integer value for the specified
// value name associated with an open key k. It also returns the value's type.
// If value does not exist, GetIntegerValue returns ErrNotExist.
// If value is not DWORD or QWORD, it will return the correct value
// type and ErrUnexpectedType.
<原文结束>

# <翻译开始>
// GetIntegerValue 用于获取与已打开键 k 关联的指定值名的整数值。同时返回该值的类型。
// 若该值不存在，GetIntegerValue 将返回 ErrNotExist 错误。
// 若该值不是 DWORD 或 QWORD 类型，它将返回正确的值类型以及 ErrUnexpectedType 错误。
# <翻译结束>


<原文开始>
// GetBinaryValue retrieves the binary value for the specified
// value name associated with an open key k. It also returns the value's type.
// If value does not exist, GetBinaryValue returns ErrNotExist.
// If value is not BINARY, it will return the correct value
// type and ErrUnexpectedType.
<原文结束>

# <翻译开始>
// GetBinaryValue 用于获取与已打开键 k 关联的指定值名称的二进制值。同时返回该值的类型。
// 若该值不存在，GetBinaryValue 将返回 ErrNotExist 错误。
// 若该值并非 BINARY 类型，它将返回正确的值类型以及 ErrUnexpectedType 错误。
# <翻译结束>


<原文开始>
// SetDWordValue sets the data and type of a name value
// under key k to value and DWORD.
<原文结束>

# <翻译开始>
// SetDWordValue 将键 k 下的某个名称值的数据和类型设置为 value 和 DWORD。
# <翻译结束>


<原文开始>
// SetQWordValue sets the data and type of a name value
// under key k to value and QWORD.
<原文结束>

# <翻译开始>
// SetQWordValue 将键 k 下名为 value 的数据及其类型设置为 QWORD。
# <翻译结束>


<原文开始>
// SetStringValue sets the data and type of a name value
// under key k to value and SZ. The value must not contain a zero byte.
<原文结束>

# <翻译开始>
// SetStringValue 将键 k 下的 name 值的数据和类型设置为 value 和 SZ。该值中不得包含零字节。
# <翻译结束>


<原文开始>
// SetExpandStringValue sets the data and type of a name value
// under key k to value and EXPAND_SZ. The value must not contain a zero byte.
<原文结束>

# <翻译开始>
// SetExpandStringValue 用于设置键 k 下名称值的数据和类型为 value 和 EXPAND_SZ。value 中不得包含零字节。
# <翻译结束>


<原文开始>
// SetStringsValue sets the data and type of a name value
// under key k to value and MULTI_SZ. The value strings
// must not contain a zero byte.
<原文结束>

# <翻译开始>
// SetStringsValue 将键 k 下名为 value 的值的数据类型及内容设置为 MULTI_SZ。其中，value 字符串中不得包含零字节。
# <翻译结束>


<原文开始>
// SetBinaryValue sets the data and type of a name value
// under key k to value and BINARY.
<原文结束>

# <翻译开始>
// SetBinaryValue 将键 k 下名为 name 的值的数据和类型设置为 value 和 BINARY。
# <翻译结束>


<原文开始>
// DeleteValue removes a named value from the key k.
<原文结束>

# <翻译开始>
// DeleteValue 从键 k 中移除一个命名值。
# <翻译结束>


<原文开始>
// ReadValueNames returns the value names of key k.
// The parameter n controls the number of returned names,
// analogous to the way os.File.Readdirnames works.
<原文结束>

# <翻译开始>
// ReadValueNames 返回键 k 的值名称。
// 参数 n 用于控制返回的名称数量，其作用方式与 os.File.Readdirnames 类似。
# <翻译结束>


<原文开始>
// extra room for terminating null character
<原文结束>

# <翻译开始>
// 额外预留终止空字符的空间
# <翻译结束>


<原文开始>
// Double buffer size and try again.
<原文结束>

# <翻译开始>
// 双倍缓冲区大小并重新尝试
# <翻译结束>

