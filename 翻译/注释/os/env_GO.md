
<原文开始>
// Expand replaces ${var} or $var in the string based on the mapping function.
// For example, os.ExpandEnv(s) is equivalent to os.Expand(s, os.Getenv).
<原文结束>

# <翻译开始>
// Expand 函数根据映射函数将字符串中的 `${var}` 或 `$var` 进行替换。
// 例如，os.ExpandEnv(s) 与 os.Expand(s, os.Getenv) 等效。
# <翻译结束>


<原文开始>
// ExpandEnv replaces ${var} or $var in the string according to the values
// of the current environment variables. References to undefined
// variables are replaced by the empty string.
<原文结束>

# <翻译开始>
// ExpandEnv 根据当前环境变量的值，将字符串中的 ${var} 或 $var 进行替换。对于未定义变量的引用，将以空字符串进行替换。
# <翻译结束>


<原文开始>
// Getenv retrieves the value of the environment variable named by the key.
// It returns the value, which will be empty if the variable is not present.
// To distinguish between an empty value and an unset value, use LookupEnv.
<原文结束>

# <翻译开始>
// Getenv 获取名为key的环境变量的值。
// 如果变量不存在，它将返回空值。
// 要区分空值和未设置的值，可以使用LookupEnv。
# <翻译结束>


<原文开始>
// LookupEnv retrieves the value of the environment variable named
// by the key. If the variable is present in the environment the
// value (which may be empty) is returned and the boolean is true.
// Otherwise the returned value will be empty and the boolean will
// be false.
<原文结束>

# <翻译开始>
// LookupEnv 从环境中检索由键命名的变量的值。如果该变量存在于环境中，将返回该值（可能为空），并且布尔值为 true。否则，返回的值将为空，布尔值将为 false。
# <翻译结束>


<原文开始>
// Setenv sets the value of the environment variable named by the key.
// It returns an error, if any.
<原文结束>

# <翻译开始>
// Setenv 设置由键命名的环境变量的值。如果发生任何错误，它将返回一个错误。
# <翻译结束>


<原文开始>
// Unsetenv unsets a single environment variable.
<原文结束>

# <翻译开始>
// Unsetenv用于取消设置单个环境变量。
# <翻译结束>


<原文开始>
// Clearenv deletes all environment variables.
<原文结束>

# <翻译开始>
// Clearenv 删除所有环境变量。
# <翻译结束>


<原文开始>
// Environ returns a copy of strings representing the environment,
// in the form "key=value".
<原文结束>

# <翻译开始>
// Environ返回一个字符串切片的副本，表示环境变量，形式为"key=value"。
# <翻译结束>

