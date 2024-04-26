
<原文开始>
// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// 版权所有 2015 The Go 作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。
# <翻译结束>


<原文开始>
// ErrShortBuffer is returned when the buffer was too short for the operation.
<原文结束>

# <翻译开始>
// ErrShortBuffer在缓冲区太短，无法进行操作时返回。
# <翻译结束>


<原文开始>
// ErrNotExist is returned when a registry key or value does not exist.
<原文结束>

# <翻译开始>
// ErrNotExist 当注册表键或值不存在时返回。
# <翻译结束>


<原文开始>
// ErrUnexpectedType is returned by Get*Value when the value's type was unexpected.
<原文结束>

# <翻译开始>
// ErrUnexpectedType 是在 Get*Value 函数遇到值类型不符合预期时返回的。
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
// GetValue 用于获取与打开的键 k 关联的指定值的类型和数据。它将数据填充到缓冲区 buf 中，并返回检索到的字节数 n。如果 buf 太小无法容纳存储的值，它会返回一个 ErrShortBuffer 错误，同时返回所需的缓冲区大小 n。如果没有提供缓冲区，它将返回 true 和实际的缓冲区大小 n。如果没有提供缓冲区，GetValue 只会返回值的类型。如果该值不存在，返回的错误是 ErrNotExist。
//
// GetValue 是一个低级函数。如果已知值的类型，应使用适当的 Get*Value 函数代替。
# <翻译结束>


<原文开始>
// GetStringValue retrieves the string value for the specified
// value name associated with an open key k. It also returns the value's type.
// If value does not exist, GetStringValue returns ErrNotExist.
// If value is not SZ or EXPAND_SZ, it will return the correct value
// type and ErrUnexpectedType.
<原文结束>

# <翻译开始>
//GetStringValue 从已打开的键k中获取与指定名称关联的字符串值。它还返回值的类型。如果不存在该值，GetStringValue将返回ErrNotExist。如果值不是SZ或EXPAND_SZ类型，它将返回正确的值类型并返回ErrUnexpectedType。
# <翻译结束>


<原文开始>
// GetMUIStringValue retrieves the localized string value for
// the specified value name associated with an open key k.
// If the value name doesn't exist or the localized string value
// can't be resolved, GetMUIStringValue returns ErrNotExist.
<原文结束>

# <翻译开始>
// GetMUIStringValue 从已打开的密钥 k 获取与指定值名称关联的本地化字符串值。如果值名称不存在或无法解析本地化字符串值，GetMUIStringValue 返回 ErrNotExist。
# <翻译结束>


<原文开始>
		// Try to resolve the string value using the system directory as
		// a DLL search path; this assumes the string value is of the form
		// @[path]\dllname,-strID but with no path given, e.g. @tzres.dll,-320.
<原文结束>

# <翻译开始>
// 尝试使用系统目录作为 DLL 搜索路径解析字符串值；
// 这种方式假定字符串值的形式为 @[路径]\dllname,-strID，但未给出路径，例如 @tzres.dll,-320。
# <翻译结束>


<原文开始>
		// This approach works with tzres.dll but may have to be revised
		// in the future to allow callers to provide custom search paths.
<原文结束>

# <翻译开始>
// 此方法适用于tzres.dll，但未来可能需要进行修订，以允许调用者提供自定义搜索路径。
# <翻译结束>


<原文开始>
// Buffer not growing, assume race; break
<原文结束>

# <翻译开始>
// 缓冲区没有增长，假设存在竞争条件；中断
# <翻译结束>


<原文开始>
// ExpandString expands environment-variable strings and replaces
// them with the values defined for the current user.
// Use ExpandString to expand EXPAND_SZ strings.
<原文结束>

# <翻译开始>
// ExpandString 扩展环境变量字符串，并用当前用户定义的值替换它们。
// 使用 ExpandString 来展开 EXPAND_SZ 字符串。
# <翻译结束>


<原文开始>
// GetStringsValue retrieves the []string value for the specified
// value name associated with an open key k. It also returns the value's type.
// If value does not exist, GetStringsValue returns ErrNotExist.
// If value is not MULTI_SZ, it will return the correct value
// type and ErrUnexpectedType.
<原文结束>

# <翻译开始>
// GetStringsValue 用于从与已打开键 k 关联的指定值名中获取 []string 值。同时返回该值的类型。
// 若该值不存在，GetStringsValue 将返回 ErrNotExist 错误。
// 若该值并非 MULTI_SZ 类型，它将返回正确的值类型及 ErrUnexpectedType 错误。
# <翻译结束>


<原文开始>
// remove terminating null
<原文结束>

# <翻译开始>
// 移除结束时的空字符
# <翻译结束>


<原文开始>
// GetIntegerValue retrieves the integer value for the specified
// value name associated with an open key k. It also returns the value's type.
// If value does not exist, GetIntegerValue returns ErrNotExist.
// If value is not DWORD or QWORD, it will return the correct value
// type and ErrUnexpectedType.
<原文结束>

# <翻译开始>
// GetIntegerValue 从已打开的密钥 k 关联的指定值名称中检索整数值。它还返回值的类型。如果值不存在，GetIntegerValue 返回 ErrNotExist。如果值不是 DWORD 或 QWORD 类型，它将返回正确的值类型并返回 ErrUnexpectedType。
# <翻译结束>


<原文开始>
// GetBinaryValue retrieves the binary value for the specified
// value name associated with an open key k. It also returns the value's type.
// If value does not exist, GetBinaryValue returns ErrNotExist.
// If value is not BINARY, it will return the correct value
// type and ErrUnexpectedType.
<原文结束>

# <翻译开始>
// GetBinaryValue 从已打开的键k关联的指定值名称中检索二进制值。它还返回值的类型。如果值不存在，GetBinaryValue 返回 ErrNotExist。如果值不是 BINARY 类型，它将返回正确的值类型和 ErrUnexpectedType。
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
// SetQWordValue 设置键k下名称值的数据和类型为value和QWORD。
# <翻译结束>


<原文开始>
// SetStringValue sets the data and type of a name value
// under key k to value and SZ. The value must not contain a zero byte.
<原文结束>

# <翻译开始>
// SetStringValue 将键 k 下的名称值的数据和类型设置为 value 和 SZ。值中不能包含零字节。
# <翻译结束>


<原文开始>
// SetExpandStringValue sets the data and type of a name value
// under key k to value and EXPAND_SZ. The value must not contain a zero byte.
<原文结束>

# <翻译开始>
// SetExpandStringValue 将键 k 下的名称值的数据和类型设置为 value 和 EXPAND_SZ。值中不能包含零字节。
# <翻译结束>


<原文开始>
// SetStringsValue sets the data and type of a name value
// under key k to value and MULTI_SZ. The value strings
// must not contain a zero byte.
<原文结束>

# <翻译开始>
// SetStringsValue 将键 k 下名为 name 的值的数据类型和值设置为 value 和 MULTI_SZ。值字符串中不得包含零字节。
# <翻译结束>


<原文开始>
// SetBinaryValue sets the data and type of a name value
// under key k to value and BINARY.
<原文结束>

# <翻译开始>
// SetBinaryValue 将键 k 下的名称值的数据和类型设置为 value 和 BINARY。
# <翻译结束>


<原文开始>
// DeleteValue removes a named value from the key k.
<原文结束>

# <翻译开始>
// DeleteValue 从键 k 中删除指定名称的值。
# <翻译结束>


<原文开始>
// ReadValueNames returns the value names of key k.
<原文结束>

# <翻译开始>
// ReadValueNames 返回键 k 的值名称。
# <翻译结束>


<原文开始>
// extra room for terminating null character
<原文结束>

# <翻译开始>
// 额外空间用于终止空字符
# <翻译结束>


<原文开始>
// Double buffer size and try again.
<原文结束>

# <翻译开始>
// 将缓冲区大小翻倍，然后重试。
# <翻译结束>

