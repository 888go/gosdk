package strings

import "strings"

// Compare 函数按照字典顺序比较两个字符串并返回一个整数。
// 当 a == b 时返回 0，当 a < b 时返回 -1，当 a > b 时返回 +1。
//
// Compare 函数仅出于与 bytes 包的对称性而包含。通常使用内建的字符串比较操作符（如 ==、<、> 等）会更清晰且始终更快。

// ff:
// a:
// b:
func Compare(a, b string) int { //md5:c280bc577895ac56d0e305b1a6f856e5
	return strings.Compare(a, b)
}
