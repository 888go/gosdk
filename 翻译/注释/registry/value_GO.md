
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
// GetValue 用于从与已打开键 k 关联的指定值中获取类型和数据。它填充缓冲区 buf 并返回所检索到的字节计数 n。如果 buf 太小无法容纳存储的值，它将返回 ErrShortBuffer 错误以及所需的缓冲区大小 n。
// 如果未提供缓冲区，它将返回 true 和实际缓冲区大小 n。
// 如果未提供缓冲区，GetValue 仅返回值的类型。
// 若该值不存在，返回的错误为 ErrNotExist。
//
// GetValue 是一个低级函数。若已知值的类型，请改用相应的 Get*Value 函数。
# <翻译结束>


<原文开始>
// GetStringValue retrieves the string value for the specified
// value name associated with an open key k. It also returns the value's type.
// If value does not exist, GetStringValue returns ErrNotExist.
// If value is not SZ or EXPAND_SZ, it will return the correct value
// type and ErrUnexpectedType.
<原文结束>

# <翻译开始>
// GetStringValue 用于从已打开的键 k 中获取与指定值名称关联的字符串值。同时返回该值的类型。
// 若值不存在，GetStringValue 将返回 ErrNotExist 错误。
// 若值并非 SZ 或 EXPAND_SZ 类型，它将返回正确的值类型以及 ErrUnexpectedType 错误。
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
// 若该值名不存在，或者无法解析其本地化字符串值，则 GetMUIStringValue 返回 ErrNotExist。
// 当系统不支持 regLoadMUIString 时，GetMUIStringValue 将引发恐慌。因此在调用此函数前，请使用 LoadRegLoadMUIString 检查系统是否支持 regLoadMUIString。
# <翻译结束>


<原文开始>
// ExpandString expands environment-variable strings and replaces
// them with the values defined for the current user.
// Use ExpandString to expand EXPAND_SZ strings.
<原文结束>

# <翻译开始>
// ExpandString 将环境变量字符串展开，并用当前用户定义的值替换它们。
// 使用 ExpandString 来展开 EXPAND_SZ 类型的字符串。
# <翻译结束>


<原文开始>
// GetStringsValue retrieves the []string value for the specified
// value name associated with an open key k. It also returns the value's type.
// If value does not exist, GetStringsValue returns ErrNotExist.
// If value is not MULTI_SZ, it will return the correct value
// type and ErrUnexpectedType.
<原文结束>

# <翻译开始>
// GetStringsValue 用于获取与已打开键 k 关联的指定值名称所对应的 []string 值。同时返回该值的类型。
// 若该值不存在，则 GetStringsValue 返回 ErrNotExist 错误。
// 若该值并非 MULTI_SZ 类型，它将返回正确的值类型以及 ErrUnexpectedType 错误。
# <翻译结束>


<原文开始>
// GetIntegerValue retrieves the integer value for the specified
// value name associated with an open key k. It also returns the value's type.
// If value does not exist, GetIntegerValue returns ErrNotExist.
// If value is not DWORD or QWORD, it will return the correct value
// type and ErrUnexpectedType.
<原文结束>

# <翻译开始>
// GetIntegerValue 用于从已打开键 k 中获取与指定值名称关联的整数值。同时返回该值的类型。
// 若该值不存在，GetIntegerValue 将返回 ErrNotExist 错误。
// 若该值非 DWORD 或 QWORD 类型，它将返回正确的值类型及 ErrUnexpectedType 错误。
# <翻译结束>


<原文开始>
// GetBinaryValue retrieves the binary value for the specified
// value name associated with an open key k. It also returns the value's type.
// If value does not exist, GetBinaryValue returns ErrNotExist.
// If value is not BINARY, it will return the correct value
// type and ErrUnexpectedType.
<原文结束>

# <翻译开始>
// GetBinaryValue 用于获取与已打开键 k 关联的指定值名的二进制值。同时返回该值的类型。
// 若该值不存在，GetBinaryValue 将返回 ErrNotExist 错误。
// 若该值并非 BINARY 类型，它将返回正确的值类型及 ErrUnexpectedType 错误。
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
// SetQWordValue 将键 k 下的名称值的数据和类型设置为 value 和 QWORD。
# <翻译结束>


<原文开始>
// SetStringValue sets the data and type of a name value
// under key k to value and SZ. The value must not contain a zero byte.
<原文结束>

# <翻译开始>
// SetStringValue 将键 k 下的名称值的数据和类型设置为 value 和 SZ。value 中不得包含零字节。
# <翻译结束>


<原文开始>
// SetExpandStringValue sets the data and type of a name value
// under key k to value and EXPAND_SZ. The value must not contain a zero byte.
<原文结束>

# <翻译开始>
// SetExpandStringValue 用于设置键 k 下名为 value 的数据及其类型为 EXPAND_SZ。该值中不得包含零字节。
# <翻译结束>


<原文开始>
// SetStringsValue sets the data and type of a name value
// under key k to value and MULTI_SZ. The value strings
// must not contain a zero byte.
<原文结束>

# <翻译开始>
// SetStringsValue 将键 k 下名为 name 的值的数据类型及内容设置为 value 和 MULTI_SZ。value 字符串中不得包含零字节。
# <翻译结束>


<原文开始>
// SetBinaryValue sets the data and type of a name value
// under key k to value and BINARY.
<原文结束>

# <翻译开始>
// SetBinaryValue 将键 k 下名为 value 的数据及其类型设置为 BINARY。
# <翻译结束>


<原文开始>
// DeleteValue removes a named value from the key k.
<原文结束>

# <翻译开始>
// DeleteValue 从键 k 中删除指定名称的值。
# <翻译结束>


<原文开始>
// ReadValueNames returns the value names of key k.
// The parameter n controls the number of returned names,
// analogous to the way os.File.Readdirnames works.
<原文结束>

# <翻译开始>
// ReadValueNames 返回键 k 的值名列表。
// 参数 n 用于控制返回的名称数量，其作用方式类似于 os.File.Readdirnames。
# <翻译结束>

