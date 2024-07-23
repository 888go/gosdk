package fmt

import "fmt"

// Errorf 根据格式说明符进行格式化，并将字符串作为满足 error 接口的值返回。
//
// 如果格式说明符中包含带有错误操作数的 %w 动词，那么返回的错误将实现一个 Unwrap 方法，该方法返回该操作数。
// 如果有多个 %w 动词，返回的错误将实现一个 Unwrap 方法，该方法返回一个 []error，其中包含所有 %w 操作数，这些操作数按照它们在参数中出现的顺序排列。
// 使用 %w 动词提供一个不实现错误接口的操作数是无效的。除此外，%w 动词与 %v 是同义的。
// md5:32b0316eed56798a
// ff:错误格式化
// format:格式化模板
// a:参数
func Errorf(format string, a ...any) error { //md5:fbc8d826a3d24402
	return fmt.Errorf(format, a...)
}
