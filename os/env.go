package os

import "os"

// Expand 函数根据映射函数将字符串中的 `${var}` 或 `$var` 进行替换。
// 例如，os.ExpandEnv(s) 与 os.Expand(s, os.Getenv) 等效。
// ff:环境变量按模板替换文本FUNC
// s:原文本
// mapping:模板函数
func Expand(s string, mapping func(string) string) string { //md5:f019dd053f61718f994fe11643caa0b0
	return os.Expand(s, mapping)
}

// ExpandEnv 根据当前环境变量的值，将字符串中的 ${var} 或 $var 进行替换。对于未定义变量的引用，将以空字符串进行替换。
// ff:环境变量按变量替换文本
// s:模板文本
func ExpandEnv(s string) string { //md5:341922ac714614c1f6a554b986385c8b
	return os.ExpandEnv(s)
}

// Getenv 获取名为key的环境变量的值。
// 如果变量不存在，它将返回空值。
// 要区分空值和未设置的值，可以使用LookupEnv。
// ff:环境变量读取
// key:名称
func Getenv(key string) string { //md5:4bb8493005d858cbb20555ae9ceb7796
	return os.Getenv(key)
}

// LookupEnv 从环境中检索由键命名的变量的值。如果该变量存在于环境中，将返回该值（可能为空），并且布尔值为 true。否则，返回的值将为空，布尔值将为 false。
// ff:环境变量查找
// key:名称
func LookupEnv(key string) (string, bool) { //md5:c98bf5dc7868bd794abce426838242b0
	return os.LookupEnv(key)
}

// Setenv 设置由键命名的环境变量的值。如果发生任何错误，它将返回一个错误。
// ff:环境变量设置
// key:名称
// value:值
func Setenv(key, value string) error { //md5:57ee46ffef752d4453e03bcee8468605
	return os.Setenv(key, value)
}

// Unsetenv用于取消设置单个环境变量。
// ff:环境变量删除
// key:名称
func Unsetenv(key string) error { //md5:6e8e565967e456c48e3b53f54d00fd95
	return os.Unsetenv(key)
}

// Clearenv 删除所有环境变量。
// ff:环境变量清除
func Clearenv() { //md5:af4b98a67dabe7ecabfecc017a8f0693
	os.Clearenv()
}

// Environ返回一个字符串切片的副本，表示环境变量，形式为"key=value"。
// ff:环境变量获取所有
func Environ() []string { //md5:f2092b95865f8368abd494520c9f6464
	return os.Environ()
}
