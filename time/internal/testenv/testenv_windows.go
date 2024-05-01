//版权所有2016 The Go Authors。所有权利保留。
//使用此源代码受BSD风格的
//可在LICENSE文件中找到的许可协议管辖。

package testenv

import (
	"os"
	"path/filepath"
	"sync"
	"syscall"
)

var symlinkOnce sync.Once
var winSymlinkErr error

func initWinHasSymlink() {
	tmpdir, err := os.MkdirTemp("", "symtest")
	if err != nil {
		panic("failed to create temp directory: " + err.Error())
	}
	defer os.RemoveAll(tmpdir)

	err = os.Symlink("target", filepath.Join(tmpdir, "symlink"))
	if err != nil {
		err = err.(*os.LinkError).Err
		switch err {
		case syscall.EWINDOWS, syscall.ERROR_PRIVILEGE_NOT_HELD:
			winSymlinkErr = err
		}
	}
}

func hasSymlink() (ok bool, reason string) {
	symlinkOnce.Do(initWinHasSymlink)

	switch winSymlinkErr {
	case nil:
		return true, ""
	case syscall.EWINDOWS:
		return false, ": symlinks are not supported on your version of Windows"
	case syscall.ERROR_PRIVILEGE_NOT_HELD:
		return false, ": you don't have enough privileges to create symlinks"
	}

	return false, ""
}
