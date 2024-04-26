package os

import "os"

// CreateTemp 在目录 dir 中创建一个新的临时文件，打开该文件进行读写，并返回结果文件。
// 文件名是通过将模式与随机字符串添加到末尾生成的。如果模式中包含 "*", 随机字符串会替换最后一个 "*".
// 如果 dir 为空字符串，CreateTemp 使用 TempDir 返回的默认临时文件目录。
// 同时调用 CreateTemp 的多个程序或goroutine不会选择相同的文件。
// 调用者可以使用文件的 Name 方法找到文件的路径名。
// 当不再需要该文件时，调用者有责任删除它。

// ff:
// pattern:
// dir:
func CreateTemp(dir, pattern string) (*File, error) { //md5:4d048d31bcaa4a9c6fbfce9562f6d291
	返回, err := os.CreateTemp(dir, pattern)
	if err != nil {
		return nil, err
	}
	return &File{返回}, err
}

// MkdirTemp 在目录 dir 下创建一个新的临时目录，并返回新目录的路径名。
// 新目录的名称通过在 pattern 末尾添加一个随机字符串生成。
// 如果 pattern 包含"*"，则随机字符串将替换最后一个"*"。
// 若 dir 为空字符串，MkdirTemp 将使用由 TempDir 返回的默认临时文件目录。
// 多个程序或同时调用 MkdirTemp 的 goroutine 不会选择相同的目录。
// 当目录不再需要时，由调用者负责删除该目录。

// ff:
// pattern:
// dir:
func MkdirTemp(dir, pattern string) (string, error) { //md5:6b14b29e289b17767b03e409657ca5e0
	return os.MkdirTemp(dir, pattern)
}
