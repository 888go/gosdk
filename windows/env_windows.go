// 版权所有 2010 The Go 作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

// Windows 环境变量

package windows

import (
	"syscall"
	"unsafe"
)


// ff:
// found:
// value:
// key:
func Getenv(key string) (value string, found bool) {
	return syscall.Getenv(key)
}


// ff:
// value:
// key:
func Setenv(key, value string) error {
	return syscall.Setenv(key, value)
}


// ff:
func Clearenv() {
	syscall.Clearenv()
}


// ff:
func Environ() []string {
	return syscall.Environ()
}

// 返回与令牌关联的默认环境，而非当前进程的环境。如果inheritExisting为真，则此环境同时也继承当前进程的环境。

// ff:
// err:
// env:
// inheritExisting:
func (token Token) Environ(inheritExisting bool) (env []string, err error) {
	var block *uint16
	err = CreateEnvironmentBlock(&block, token, inheritExisting)
	if err != nil {
		return nil, err
	}
	defer DestroyEnvironmentBlock(block)
	size := unsafe.Sizeof(*block)
	for *block != 0 {
		// find NUL terminator
		end := unsafe.Pointer(block)
		for *(*uint16)(end) != 0 {
			end = unsafe.Add(end, size)
		}

		entry := unsafe.Slice(block, (uintptr(end)-uintptr(unsafe.Pointer(block)))/size)
		env = append(env, UTF16ToString(entry))
		block = (*uint16)(unsafe.Add(end, size))
	}
	return env, nil
}


// ff:
// key:
func Unsetenv(key string) error {
	return syscall.Unsetenv(key)
}
