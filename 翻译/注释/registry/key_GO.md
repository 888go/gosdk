
<原文开始>
// Close closes open key k.
<原文结束>

# <翻译开始>
// Close 关闭已打开的密钥 k。
# <翻译结束>


<原文开始>
// OpenKey opens a new key with path name relative to key k.
// It accepts any open key, including CURRENT_USER and others,
// and returns the new key and an error.
// The access parameter specifies desired access rights to the
// key to be opened.
<原文结束>

# <翻译开始>
// OpenKey 以相对于键 k 的路径名称打开一个新键。
// 它接受任何已打开的键，包括 CURRENT_USER 等，
// 并返回新键和错误信息。
// 参数 access 指定了对将要打开的键所期望的访问权限。
# <翻译结束>


<原文开始>
// OpenRemoteKey opens a predefined registry key on another
// computer pcname. The key to be opened is specified by k, but
// can only be one of LOCAL_MACHINE, PERFORMANCE_DATA or USERS.
// If pcname is "", OpenRemoteKey returns local computer key.
<原文结束>

# <翻译开始>
// OpenRemoteKey 用于在另一台计算机（pcname）上打开一个预定义的注册表键。
// 待打开的键由参数 k 指定，但其值只能是 LOCAL_MACHINE、PERFORMANCE_DATA 或 USERS 中的一个。
// 若 pcname 为空字符串（""），则 OpenRemoteKey 将返回本地计算机的键。
# <翻译结束>


<原文开始>
// ReadSubKeyNames returns the names of subkeys of key k.
// The parameter n controls the number of returned names,
// analogous to the way os.File.Readdirnames works.
<原文结束>

# <翻译开始>
// ReadSubKeyNames 返回键 k 的子键名称。
// 参数 n 用于控制返回的名称数量，其作用方式与 os.File.Readdirnames 类似。
# <翻译结束>


<原文开始>
// CreateKey creates a key named path under open key k.
// CreateKey returns the new key and a boolean flag that reports
// whether the key already existed.
// The access parameter specifies the access rights for the key
// to be created.
<原文结束>

# <翻译开始>
// CreateKey 在已打开的键 k 下创建一个名为 path 的键。
// CreateKey 返回新键以及一个布尔标志，该标志报告
// 该键是否已存在。
// 参数 access 指定了要创建的键的访问权限。
# <翻译结束>


<原文开始>
// DeleteKey deletes the subkey path of key k and its values.
<原文结束>

# <翻译开始>
// DeleteKey 删除键k的子键路径及其值。
# <翻译结束>


<原文开始>
// ModTime returns the key's last write time.
<原文结束>

# <翻译开始>
// ModTime 返回键的最后写入时间。
# <翻译结束>


<原文开始>
// Stat retrieves information about the open key k.
<原文结束>

# <翻译开始>
// Stat 获取关于已打开键 k 的信息。
# <翻译结束>

