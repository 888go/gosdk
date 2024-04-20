// 版权所有 ? 2009 Go作者。保留所有权利。
// 本源代码的使用受BSD风格
// 许可证约束，该许可证可在LICENSE文件中找到。

// Fork, exec, wait, etc.

package windows

import (
	errorspkg "errors"
	"unsafe"
)

// EscapeArg 按照 http://msdn.microsoft.com/en-us/library/ms880421 中所规定的规则重写命令行参数 s。
// 若 s 为空，则此函数返回空字符串（即 ""）。
// 否则，将进行如下转换：
//   - 只有当反斜杠 (\) 立即跟随双引号 (") 时，才将每个反斜杠加倍；
//   - 用反斜杠对每个双引号 (") 进行转义；
//   - 最后，只有当 s 内包含空格或制表符时，才将 s 用双引号包裹（即将 arg 转为 "arg"）。
func EscapeArg(s string) string {
	if len(s) == 0 {
		return `""`
	}
	n := len(s)
	hasSpace := false
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '"', '\\':
			n++
		case ' ', '\t':
			hasSpace = true
		}
	}
	if hasSpace {
		n += 2 // 为引号预留空间
	}
	if n == len(s) {
		return s
	}

	qs := make([]byte, n)
	j := 0
	if hasSpace {
		qs[j] = '"'
		j++
	}
	slashes := 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		default:
			slashes = 0
			qs[j] = s[i]
		case '\\':
			slashes++
			qs[j] = s[i]
		case '"':
			for ; slashes > 0; slashes-- {
				qs[j] = '\\'
				j++
			}
			qs[j] = '\\'
			j++
			qs[j] = s[i]
		}
		j++
	}
	if hasSpace {
		for ; slashes > 0; slashes-- {
			qs[j] = '\\'
			j++
		}
		qs[j] = '"'
		j++
	}
	return string(qs[:j])
}

// ComposeCommandLine 对给定的参数进行转义并拼接，适用于作为 Windows 命令行使用，
// 用于 CreateProcess 的 CommandLine 参数、CreateService/ChangeServiceConfig 的 BinaryPathName 参数，
// 或任何使用 CommandLineToArgv 的程序。
func ComposeCommandLine(args []string) string {
	if len(args) == 0 {
		return ""
	}

	// Per https://learn.microsoft.com/en-us/windows/win32/api/shellapi/nf-shellapi-commandlinetoargvw:
	// “This function accepts command lines that contain a program name; the
	// program name can be enclosed in quotation marks or not.”
	//
	// Unfortunately, it provides no means of escaping interior quotation marks
	// within that program name, and we have no way to report them here.
	prog := args[0]
	mustQuote := len(prog) == 0
	for i := 0; i < len(prog); i++ {
		c := prog[i]
		if c <= ' ' || (c == '"' && i == 0) {
// 强制对非仅ASCII空格和制表符（如MSDN文章所述）进行引号处理，还包括ASCII控制字符。
// CommandLineToArgvW函数的文档并未说明当第一个参数不是有效的程序名时会发生什么，
// 但实证表明，它似乎会忽略未加引号的控制字符。
			mustQuote = true
			break
		}
	}
	var commandLine []byte
	if mustQuote {
		commandLine = make([]byte, 0, len(prog)+2)
		commandLine = append(commandLine, '"')
		for i := 0; i < len(prog); i++ {
			c := prog[i]
			if c == '"' {
// 该引号会与我们周围的引号产生冲突。
// 我们无法报告错误，因此只需移除
// 有问题的字符即可。
				continue
			}
			commandLine = append(commandLine, c)
		}
		commandLine = append(commandLine, '"')
	} else {
		if len(args) == 1 {
// args[0] 是一个表示其自身的有效命令行参数。
// 无需为其分配新的切片或字符串。
			return prog
		}
		commandLine = []byte(prog)
	}

	for _, arg := range args[1:] {
		commandLine = append(commandLine, ' ')
// TODO(bcmills): 既然我们已在向切片追加内容，最好能避免 EscapeArg 的中间分配过程。
// 或许我们可以提取出一个 appendEscapedArg 函数。
		commandLine = append(commandLine, EscapeArg(arg)...)
	}
	return string(commandLine)
}

// DecomposeCommandLine使用CommandLineToArgv函数将其参数命令行分解为未转义的部分，这些命令行来自GetCommandLine、QUERY_SERVICE_CONFIG的BinaryPathName参数，或其他任何传递命令行的地方。
// 若commandLine包含NUL字符，DecomposeCommandLine将返回错误。
func DecomposeCommandLine(commandLine string) ([]string, error) {
	if len(commandLine) == 0 {
		return []string{}, nil
	}
	utf16CommandLine, err := UTF16FromString(commandLine)
	if err != nil {
		return nil, errorspkg.New("string with NUL passed to DecomposeCommandLine")
	}
	var argc int32
	argv, err := commandLineToArgv(&utf16CommandLine[0], &argc)
	if err != nil {
		return nil, err
	}
	defer LocalFree(Handle(unsafe.Pointer(argv)))

	var args []string
	for _, p := range unsafe.Slice(argv, argc) {
		args = append(args, UTF16PtrToString(p))
	}
	return args, nil
}

// CommandLineToArgv 函数解析一个Unicode格式的命令行字符串，并将解析得到的参数个数设置到 argc 中。
//
// 返回的内存应当使用一次 LocalFree 调用进行释放。
//
// 需要注意的是，虽然 CommandLineToArgv 的返回类型暗示了可容纳8192个、每个最多8192字符的条目，
// 实际解析出的参数数量可能超过8192个，并且 CommandLineToArgvW 的文档并未提及单个参数字符串长度的任何限制。
// （参见 https://go.dev/issue/63236。）
func CommandLineToArgv(cmd *uint16, argc *int32) (argv *[8192]*[8192]uint16, err error) {
	argp, err := commandLineToArgv(cmd, argc)
	argv = (*[8192]*[8192]uint16)(unsafe.Pointer(argp))
	return argv, err
}

func CloseOnExec(fd Handle) {
	SetHandleInformation(Handle(fd), HANDLE_FLAG_INHERIT, 0)
}

// FullPath 获取指定文件的完整路径
func FullPath(name string) (path string, err error) {
	p, err := UTF16PtrFromString(name)
	if err != nil {
		return "", err
	}
	n := uint32(100)
	for {
		buf := make([]uint16, n)
		n, err = GetFullPathName(p, uint32(len(buf)), &buf[0], nil)
		if err != nil {
			return "", err
		}
		if n <= uint32(len(buf)) {
			return UTF16ToString(buf[:n]), nil
		}
	}
}

// NewProcThreadAttributeList 分配一个新的 ProcThreadAttributeListContainer，其包含请求的最大属性数量。
func NewProcThreadAttributeList(maxAttrCount uint32) (*ProcThreadAttributeListContainer, error) {
	var size uintptr
	err := initializeProcThreadAttributeList(nil, maxAttrCount, 0, &size)
	if err != ERROR_INSUFFICIENT_BUFFER {
		if err == nil {
			return nil, errorspkg.New("unable to query buffer size from InitializeProcThreadAttributeList")
		}
		return nil, err
	}
	alloc, err := LocalAlloc(LMEM_FIXED, uint32(size))
	if err != nil {
		return nil, err
	}
	// size is guaranteed to be ≥1 by InitializeProcThreadAttributeList.
	al := &ProcThreadAttributeListContainer{data: (*ProcThreadAttributeList)(unsafe.Pointer(alloc))}
	err = initializeProcThreadAttributeList(al.data, maxAttrCount, 0, &size)
	if err != nil {
		return nil, err
	}
	return al, err
}

// Update 通过调用 UpdateProcThreadAttribute 函数来修改 ProcThreadAttributeList。
func (al *ProcThreadAttributeListContainer) Update(attribute uintptr, value unsafe.Pointer, size uintptr) error {
	al.pointers = append(al.pointers, value)
	return updateProcThreadAttribute(al.data, 0, attribute, value, size, nil, nil)
}

// Delete 释放 ProcThreadAttributeList 的资源。
func (al *ProcThreadAttributeListContainer) Delete() {
	deleteProcThreadAttributeList(al.data)
	LocalFree(Handle(unsafe.Pointer(al.data)))
	al.data = nil
	al.pointers = nil
}

// List 返回将传递给 StartupInfoEx 的实际 ProcThreadAttributeList。
func (al *ProcThreadAttributeListContainer) List() *ProcThreadAttributeList {
	return al.data
}
