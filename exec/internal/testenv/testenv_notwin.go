// 版权所有 2016 The Go 作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

//go:build !windows

package testenv

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

func hasSymlink() (ok bool, reason string) {
	switch runtime.GOOS {
	case "plan9":
		return false, ""
	case "android", "wasip1":
// 对于wasip1，某些运行时禁止使用绝对路径的符号链接，
// 或者是那些指向当前工作目录之外的符号链接。
// 我们进行一个简单的测试以判断该运行时是否支持符号链接。
// 如果遇到权限错误，说明运行时不支持符号链接。
		dir, err := os.MkdirTemp("", "")
		if err != nil {
			return false, ""
		}
		defer func() {
			_ = os.RemoveAll(dir)
		}()
		fpath := filepath.Join(dir, "testfile.txt")
		if err := os.WriteFile(fpath, nil, 0644); err != nil {
			return false, ""
		}
		if err := os.Symlink(fpath, filepath.Join(dir, "testlink")); err != nil {
			if SyscallIsNotSupported(err) {
				return false, fmt.Sprintf("symlinks unsupported: %s", err.Error())
			}
			return false, ""
		}
	}

	return true, ""
}
