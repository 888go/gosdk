// 版权所有 (C) 2010 Go 作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可协议约束，
// 该协议可在 LICENSE 文件中找到。

// Windows 环境变量

package windows

import (
	"syscall"
	"unsafe"
)


// ff:取环境变量
// found:是否成功
// value:值
// key:名称
func Getenv(key string) (value string, found bool) {
	return syscall.Getenv(key)
}


// ff:设置环境变量
// value:值
// key:名称
func Setenv(key, value string) error {
	return syscall.Setenv(key, value)
}


// ff:删除所有环境变量
func Clearenv() {
	syscall.Clearenv()
}


// ff:取所有环境变量
func Environ() []string {
	return syscall.Environ()
}

// 返回与令牌关联的默认环境，而非当前进程的环境。如果inheritExisting为真，则此环境同时也继承当前进程的环境。

// ff:取所有环境变量
// err:错误
// env:所有环境变量
// inheritExisting:继承现有进程
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


// ff:删除环境变量
// key:名称
func Unsetenv(key string) error {
	return syscall.Unsetenv(key)
}
