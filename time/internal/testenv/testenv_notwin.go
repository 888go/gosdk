//版权所有2016 The Go Authors。所有权利保留。
//使用此源代码受BSD风格的
//可在LICENSE文件中找到的许可协议管辖。

//---build---//go:build !windows

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
// 对于wasip1，某些运行时禁止使用绝对符号链接，
// 或者那些能够脱离当前工作目录的符号链接。
// 执行一个简单的测试以判断运行时是否支持符号链接。
// 如果我们遇到权限错误，则说明运行时不支持符号链接。
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
