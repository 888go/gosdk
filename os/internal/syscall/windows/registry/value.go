// 版权所有 2015 The Go 作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

//go:build windows

package registry

import (
	"errors"
	"syscall"
	"unicode/utf16"
	"unsafe"
)

const (
	// Registry value types.
	NONE                       = 0
	SZ                         = 1
	EXPAND_SZ                  = 2
	BINARY                     = 3
	DWORD                      = 4
	DWORD_BIG_ENDIAN           = 5
	LINK                       = 6
	MULTI_SZ                   = 7
	RESOURCE_LIST              = 8
	FULL_RESOURCE_DESCRIPTOR   = 9
	RESOURCE_REQUIREMENTS_LIST = 10
	QWORD                      = 11
)

var (
	// ErrShortBuffer在缓冲区太短，无法进行操作时返回。
	ErrShortBuffer = syscall.ERROR_MORE_DATA

	// ErrNotExist 当注册表键或值不存在时返回。
	ErrNotExist = syscall.ERROR_FILE_NOT_FOUND

	// ErrUnexpectedType 是在 Get*Value 函数遇到值类型不符合预期时返回的。
	ErrUnexpectedType = errors.New("unexpected key value type")
)

// GetValue 用于获取与打开的键 k 关联的指定值的类型和数据。它将数据填充到缓冲区 buf 中，并返回检索到的字节数 n。如果 buf 太小无法容纳存储的值，它会返回一个 ErrShortBuffer 错误，同时返回所需的缓冲区大小 n。如果没有提供缓冲区，它将返回 true 和实际的缓冲区大小 n。如果没有提供缓冲区，GetValue 只会返回值的类型。如果该值不存在，返回的错误是 ErrNotExist。
//
// GetValue 是一个低级函数。如果已知值的类型，应使用适当的 Get*Value 函数代替。
func (k Key) GetValue(name string, buf []byte) (n int, valtype uint32, err error) {
	pname, err := syscall.UTF16PtrFromString(name)
	if err != nil {
		return 0, 0, err
	}
	var pbuf *byte
	if len(buf) > 0 {
		pbuf = (*byte)(unsafe.Pointer(&buf[0]))
	}
	l := uint32(len(buf))
	err = syscall.RegQueryValueEx(syscall.Handle(k), pname, nil, &valtype, pbuf, &l)
	if err != nil {
		return int(l), valtype, err
	}
	return int(l), valtype, nil
}

func (k Key) getValue(name string, buf []byte) (date []byte, valtype uint32, err error) {
	p, err := syscall.UTF16PtrFromString(name)
	if err != nil {
		return nil, 0, err
	}
	var t uint32
	n := uint32(len(buf))
	for {
		err = syscall.RegQueryValueEx(syscall.Handle(k), p, nil, &t, (*byte)(unsafe.Pointer(&buf[0])), &n)
		if err == nil {
			return buf[:n], t, nil
		}
		if err != syscall.ERROR_MORE_DATA {
			return nil, 0, err
		}
		if n <= uint32(len(buf)) {
			return nil, 0, err
		}
		buf = make([]byte, n)
	}
}

//GetStringValue 从已打开的键k中获取与指定名称关联的字符串值。它还返回值的类型。如果不存在该值，GetStringValue将返回ErrNotExist。如果值不是SZ或EXPAND_SZ类型，它将返回正确的值类型并返回ErrUnexpectedType。
func (k Key) GetStringValue(name string) (val string, valtype uint32, err error) {
	data, typ, err2 := k.getValue(name, make([]byte, 64))
	if err2 != nil {
		return "", typ, err2
	}
	switch typ {
	case SZ, EXPAND_SZ:
	default:
		return "", typ, ErrUnexpectedType
	}
	if len(data) == 0 {
		return "", typ, nil
	}
	u := (*[1 << 29]uint16)(unsafe.Pointer(&data[0]))[: len(data)/2 : len(data)/2]
	return syscall.UTF16ToString(u), typ, nil
}

// GetMUIStringValue 从已打开的密钥 k 获取与指定值名称关联的本地化字符串值。如果值名称不存在或无法解析本地化字符串值，GetMUIStringValue 返回 ErrNotExist。
func (k Key) GetMUIStringValue(name string) (string, error) {
	pname, err := syscall.UTF16PtrFromString(name)
	if err != nil {
		return "", err
	}

	buf := make([]uint16, 1024)
	var buflen uint32
	var pdir *uint16

	err = regLoadMUIString(syscall.Handle(k), pname, &buf[0], uint32(len(buf)), &buflen, 0, pdir)
	if err == syscall.ERROR_FILE_NOT_FOUND { // Try fallback path

// 尝试使用系统目录作为 DLL 搜索路径解析字符串值；
// 这种方式假定字符串值的形式为 @[路径]\dllname,-strID，但未给出路径，例如 @tzres.dll,-320。

// 此方法适用于tzres.dll，但未来可能需要进行修订，以允许调用者提供自定义搜索路径。

		var s string
		s, err = ExpandString("%SystemRoot%\\system32\\")
		if err != nil {
			return "", err
		}
		pdir, err = syscall.UTF16PtrFromString(s)
		if err != nil {
			return "", err
		}

		err = regLoadMUIString(syscall.Handle(k), pname, &buf[0], uint32(len(buf)), &buflen, 0, pdir)
	}

	for err == syscall.ERROR_MORE_DATA { // Grow buffer if needed
		if buflen <= uint32(len(buf)) {
			break // 缓冲区没有增长，假设存在竞争条件；中断
		}
		buf = make([]uint16, buflen)
		err = regLoadMUIString(syscall.Handle(k), pname, &buf[0], uint32(len(buf)), &buflen, 0, pdir)
	}

	if err != nil {
		return "", err
	}

	return syscall.UTF16ToString(buf), nil
}

// ExpandString 扩展环境变量字符串，并用当前用户定义的值替换它们。
// 使用 ExpandString 来展开 EXPAND_SZ 字符串。
func ExpandString(value string) (string, error) {
	if value == "" {
		return "", nil
	}
	p, err := syscall.UTF16PtrFromString(value)
	if err != nil {
		return "", err
	}
	r := make([]uint16, 100)
	for {
		n, err := expandEnvironmentStrings(p, &r[0], uint32(len(r)))
		if err != nil {
			return "", err
		}
		if n <= uint32(len(r)) {
			return syscall.UTF16ToString(r[:n]), nil
		}
		r = make([]uint16, n)
	}
}

// GetStringsValue 用于从与已打开键 k 关联的指定值名中获取 []string 值。同时返回该值的类型。
// 若该值不存在，GetStringsValue 将返回 ErrNotExist 错误。
// 若该值并非 MULTI_SZ 类型，它将返回正确的值类型及 ErrUnexpectedType 错误。
func (k Key) GetStringsValue(name string) (val []string, valtype uint32, err error) {
	data, typ, err2 := k.getValue(name, make([]byte, 64))
	if err2 != nil {
		return nil, typ, err2
	}
	if typ != MULTI_SZ {
		return nil, typ, ErrUnexpectedType
	}
	if len(data) == 0 {
		return nil, typ, nil
	}
	p := (*[1 << 29]uint16)(unsafe.Pointer(&data[0]))[: len(data)/2 : len(data)/2]
	if len(p) == 0 {
		return nil, typ, nil
	}
	if p[len(p)-1] == 0 {
		p = p[:len(p)-1] // 移除结束时的空字符
	}
	val = make([]string, 0, 5)
	from := 0
	for i, c := range p {
		if c == 0 {
			val = append(val, syscall.UTF16ToString(p[from:i]))
			from = i + 1
		}
	}
	return val, typ, nil
}

// GetIntegerValue 从已打开的密钥 k 关联的指定值名称中检索整数值。它还返回值的类型。如果值不存在，GetIntegerValue 返回 ErrNotExist。如果值不是 DWORD 或 QWORD 类型，它将返回正确的值类型并返回 ErrUnexpectedType。
func (k Key) GetIntegerValue(name string) (val uint64, valtype uint32, err error) {
	data, typ, err2 := k.getValue(name, make([]byte, 8))
	if err2 != nil {
		return 0, typ, err2
	}
	switch typ {
	case DWORD:
		if len(data) != 4 {
			return 0, typ, errors.New("DWORD value is not 4 bytes long")
		}
		return uint64(*(*uint32)(unsafe.Pointer(&data[0]))), DWORD, nil
	case QWORD:
		if len(data) != 8 {
			return 0, typ, errors.New("QWORD value is not 8 bytes long")
		}
		return *(*uint64)(unsafe.Pointer(&data[0])), QWORD, nil
	default:
		return 0, typ, ErrUnexpectedType
	}
}

// GetBinaryValue 从已打开的键k关联的指定值名称中检索二进制值。它还返回值的类型。如果值不存在，GetBinaryValue 返回 ErrNotExist。如果值不是 BINARY 类型，它将返回正确的值类型和 ErrUnexpectedType。
func (k Key) GetBinaryValue(name string) (val []byte, valtype uint32, err error) {
	data, typ, err2 := k.getValue(name, make([]byte, 64))
	if err2 != nil {
		return nil, typ, err2
	}
	if typ != BINARY {
		return nil, typ, ErrUnexpectedType
	}
	return data, typ, nil
}

func (k Key) setValue(name string, valtype uint32, data []byte) error {
	p, err := syscall.UTF16PtrFromString(name)
	if err != nil {
		return err
	}
	if len(data) == 0 {
		return regSetValueEx(syscall.Handle(k), p, 0, valtype, nil, 0)
	}
	return regSetValueEx(syscall.Handle(k), p, 0, valtype, &data[0], uint32(len(data)))
}

// SetDWordValue 将键 k 下的某个名称值的数据和类型设置为 value 和 DWORD。
func (k Key) SetDWordValue(name string, value uint32) error {
	return k.setValue(name, DWORD, (*[4]byte)(unsafe.Pointer(&value))[:])
}

// SetQWordValue 设置键k下名称值的数据和类型为value和QWORD。
func (k Key) SetQWordValue(name string, value uint64) error {
	return k.setValue(name, QWORD, (*[8]byte)(unsafe.Pointer(&value))[:])
}

func (k Key) setStringValue(name string, valtype uint32, value string) error {
	v, err := syscall.UTF16FromString(value)
	if err != nil {
		return err
	}
	buf := (*[1 << 29]byte)(unsafe.Pointer(&v[0]))[: len(v)*2 : len(v)*2]
	return k.setValue(name, valtype, buf)
}

// SetStringValue 将键 k 下的名称值的数据和类型设置为 value 和 SZ。值中不能包含零字节。
func (k Key) SetStringValue(name, value string) error {
	return k.setStringValue(name, SZ, value)
}

// SetExpandStringValue 将键 k 下的名称值的数据和类型设置为 value 和 EXPAND_SZ。值中不能包含零字节。
func (k Key) SetExpandStringValue(name, value string) error {
	return k.setStringValue(name, EXPAND_SZ, value)
}

// SetStringsValue 将键 k 下名为 name 的值的数据类型和值设置为 value 和 MULTI_SZ。值字符串中不得包含零字节。
func (k Key) SetStringsValue(name string, value []string) error {
	ss := ""
	for _, s := range value {
		for i := 0; i < len(s); i++ {
			if s[i] == 0 {
				return errors.New("string cannot have 0 inside")
			}
		}
		ss += s + "\x00"
	}
	v := utf16.Encode([]rune(ss + "\x00"))
	buf := (*[1 << 29]byte)(unsafe.Pointer(&v[0]))[: len(v)*2 : len(v)*2]
	return k.setValue(name, MULTI_SZ, buf)
}

// SetBinaryValue 将键 k 下的名称值的数据和类型设置为 value 和 BINARY。
func (k Key) SetBinaryValue(name string, value []byte) error {
	return k.setValue(name, BINARY, value)
}

// DeleteValue 从键 k 中删除指定名称的值。
func (k Key) DeleteValue(name string) error {
	return regDeleteValue(syscall.Handle(k), syscall.StringToUTF16Ptr(name))
}

// ReadValueNames 返回键 k 的值名称。
func (k Key) ReadValueNames() ([]string, error) {
	ki, err := k.Stat()
	if err != nil {
		return nil, err
	}
	names := make([]string, 0, ki.ValueCount)
	buf := make([]uint16, ki.MaxValueNameLen+1) // 额外空间用于终止空字符
loopItems:
	for i := uint32(0); ; i++ {
		l := uint32(len(buf))
		for {
			err := regEnumValue(syscall.Handle(k), i, &buf[0], &l, nil, nil, nil, nil)
			if err == nil {
				break
			}
			if err == syscall.ERROR_MORE_DATA {
				// 将缓冲区大小翻倍，然后重试。
				l = uint32(2 * len(buf))
				buf = make([]uint16, l)
				continue
			}
			if err == _ERROR_NO_MORE_ITEMS {
				break loopItems
			}
			return names, err
		}
		names = append(names, syscall.UTF16ToString(buf[:l]))
	}
	return names, nil
}
