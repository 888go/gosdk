// 版权所有 ? 2015 The Go 作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

//go:build windows

package registry

import (
	"errors"
	"io"
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
	// ErrShortBuffer 表示当缓冲区对于操作而言过短时返回的错误。
	ErrShortBuffer = syscall.ERROR_MORE_DATA

	// ErrNotExist 表示当注册表键或值不存在时返回的错误。
	ErrNotExist = syscall.ERROR_FILE_NOT_FOUND

	// ErrUnexpectedType 会在获取值时遇到类型不符预期的情况时返回。
	ErrUnexpectedType = errors.New("unexpected key value type")
)

// GetValue 用于从与已打开键 k 关联的指定值中检索类型和数据。
// 它将数据填充到缓冲区 buf 中，并返回获取的字节数 n。
// 如果 buf 太小无法容纳存储的值，它将返回 ErrShortBuffer 错误以及所需的缓冲区大小 n。
// 如果未提供缓冲区，它将返回 true 以及实际缓冲区大小 n。
// 如果未提供缓冲区，GetValue 仅返回值的类型。
// 若该值不存在，返回的错误为 ErrNotExist。
//
// GetValue 是一个低级函数。若已知值的类型，请使用相应的 Get*Value 函数代替。

// ff:取值
// err:错误
// valtype:值类型
// n:
// buf:缓冲区
// name:名称
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

func (k Key) getValue(name string, buf []byte) (data []byte, valtype uint32, err error) {
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

// GetStringValue 用于获取与已打开键 k 关联的指定值名的字符串值。同时返回该值的类型。
// 若该值不存在，GetStringValue 将返回 ErrNotExist 错误。
// 若该值非 SZ 或 EXPAND_SZ 类型，它将返回正确的值类型以及 ErrUnexpectedType 错误。

// ff:取文本值
// err:错误
// valtype:值类型
// val:值
// name:名称
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

// GetMUIStringValue 用于获取与已打开键 k 关联的指定值名所对应的本地化字符串值。
// 若该值名不存在，或无法解析其本地化字符串值，则 GetMUIStringValue 返回 ErrNotExist。
// 若系统不支持 regLoadMUIString，GetMUIStringValue 将引发恐慌。因此在调用此函数前，请使用 LoadRegLoadMUIString 检查系统是否支持 regLoadMUIString。

// ff:取文本值PANI
// name:名称
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
// 这种方式假设字符串值形如 @[路径]\dllname,-strID，但未给出具体路径，
// 例如：@tzres.dll,-320

// 此方法适用于tzres.dll，但未来可能需要进行修订，
// 以便调用者能够提供自定义搜索路径。

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
			break // 缓冲区未增长，假设存在竞争；中断
		}
		buf = make([]uint16, buflen)
		err = regLoadMUIString(syscall.Handle(k), pname, &buf[0], uint32(len(buf)), &buflen, 0, pdir)
	}

	if err != nil {
		return "", err
	}

	return syscall.UTF16ToString(buf), nil
}

// ExpandString 将环境变量字符串展开，并用当前用户所定义的值进行替换。
// 对于 EXPAND_SZ 类型的字符串，应使用 ExpandString 进行展开。

// ff:解析环境变量
// value:值
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

// GetStringsValue 用于从与已打开键 k 关联的指定值名中检索 []string 值。同时返回该值的类型。
// 若值不存在，GetStringsValue 返回 ErrNotExist 错误。
// 若值并非 MULTI_SZ 类型，它将返回正确的值类型及 ErrUnexpectedType 错误。

// ff:取文本切片值
// err:错误
// valtype:值类型
// val:切片值
// name:名称
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
		p = p[:len(p)-1] // 移除终止空字符
	}
	val = make([]string, 0, 5)
	from := 0
	for i, c := range p {
		if c == 0 {
			val = append(val, string(utf16.Decode(p[from:i])))
			from = i + 1
		}
	}
	return val, typ, nil
}

// GetIntegerValue 用于获取与已打开键 k 关联的指定值名的整数值。同时返回该值的类型。
// 若该值不存在，GetIntegerValue 将返回 ErrNotExist 错误。
// 若该值不是 DWORD 或 QWORD 类型，它将返回正确的值类型以及 ErrUnexpectedType 错误。

// ff:取整数64位值
// err:错误
// valtype:值类型
// val:值
// name:名称
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
		var val32 uint32
		copy((*[4]byte)(unsafe.Pointer(&val32))[:], data)
		return uint64(val32), DWORD, nil
	case QWORD:
		if len(data) != 8 {
			return 0, typ, errors.New("QWORD value is not 8 bytes long")
		}
		copy((*[8]byte)(unsafe.Pointer(&val))[:], data)
		return val, QWORD, nil
	default:
		return 0, typ, ErrUnexpectedType
	}
}

// GetBinaryValue 用于获取与已打开键 k 关联的指定值名称的二进制值。同时返回该值的类型。
// 若该值不存在，GetBinaryValue 将返回 ErrNotExist 错误。
// 若该值并非 BINARY 类型，它将返回正确的值类型以及 ErrUnexpectedType 错误。

// ff:取字节集值
// err:错误
// valtype:值类型
// val:值
// name:名称
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

// ff:取整数32位值
// value:值
// name:名称
func (k Key) SetDWordValue(name string, value uint32) error {
	return k.setValue(name, DWORD, (*[4]byte)(unsafe.Pointer(&value))[:])
}

// SetQWordValue 将键 k 下名为 value 的数据及其类型设置为 QWORD。

// ff:设置整数64位值
// value:值
// name:名称
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

// SetStringValue 将键 k 下的 name 值的数据和类型设置为 value 和 SZ。该值中不得包含零字节。

// ff:设置文本值
// value:值
// name:名称
func (k Key) SetStringValue(name, value string) error {
	return k.setStringValue(name, SZ, value)
}

// SetExpandStringValue 用于设置键 k 下名称值的数据和类型为 value 和 EXPAND_SZ。value 中不得包含零字节。

// ff:设置文本值并按环境变量
// value:值
// name:名称
func (k Key) SetExpandStringValue(name, value string) error {
	return k.setStringValue(name, EXPAND_SZ, value)
}

// SetStringsValue 将键 k 下名为 value 的值的数据类型及内容设置为 MULTI_SZ。其中，value 字符串中不得包含零字节。

// ff:设置文本切片值
// value:切片值
// name:名称
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

// SetBinaryValue 将键 k 下名为 name 的值的数据和类型设置为 value 和 BINARY。

// ff:设置字节集值
// value:值
// name:名称
func (k Key) SetBinaryValue(name string, value []byte) error {
	return k.setValue(name, BINARY, value)
}

// DeleteValue 从键 k 中移除一个命名值。

// ff:删除值
// name:名称
func (k Key) DeleteValue(name string) error {
	return regDeleteValue(syscall.Handle(k), syscall.StringToUTF16Ptr(name))
}

// ReadValueNames 返回键 k 的值名称。
// 参数 n 用于控制返回的名称数量，其作用方式与 os.File.Readdirnames 类似。

// ff:取所有子项值
// n:返回数量
func (k Key) ReadValueNames(n int) ([]string, error) {
	ki, err := k.Stat()
	if err != nil {
		return nil, err
	}
	names := make([]string, 0, ki.ValueCount)
	buf := make([]uint16, ki.MaxValueNameLen+1) // 额外预留终止空字符的空间
loopItems:
	for i := uint32(0); ; i++ {
		if n > 0 {
			if len(names) == n {
				return names, nil
			}
		}
		l := uint32(len(buf))
		for {
			err := regEnumValue(syscall.Handle(k), i, &buf[0], &l, nil, nil, nil, nil)
			if err == nil {
				break
			}
			if err == syscall.ERROR_MORE_DATA {
				// 双倍缓冲区大小并重新尝试
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
	if n > len(names) {
		return names, io.EOF
	}
	return names, nil
}
