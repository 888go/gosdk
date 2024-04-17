//go:build windows

package registry

import (
	"golang.org/x/sys/windows/registry"
	"syscall"
)

const (
	_REG_OPTION_NON_VOLATILE = 0

	_REG_CREATED_NEW_KEY     = 1
	_REG_OPENED_EXISTING_KEY = 2

	_ERROR_NO_MORE_ITEMS syscall.Errno = 259
)


// ff:
func LoadRegLoadMUIString() error { //md5:d97a2024fe61c3e084a6983955f54642
	return registry.LoadRegLoadMUIString()
}
