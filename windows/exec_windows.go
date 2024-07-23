// 版权所有 ? 2009 Go 作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 此许可证可在 LICENSE 文件中找到。

// Fork, exec, wait, etc.

package windows

import (
	errorspkg "errors"
	"unsafe"
)

// EscapeArg 按照 http://msdn.microsoft.com/en-us/library/ms880421 中所规定的方式重写命令行参数 s。
// 如果 s 为空，此函数返回 ""（两个双引号）。
// 另外，进行如下转换：
//   - 若反斜杠 (\) 立即跟随双引号 (")，则将每个反斜杠加倍；
//   - 用反斜杠对每个双引号 (") 进行转义；
//   - 最后，仅当 s 内包含空格或制表符时，用双引号包裹 s（将 arg 转换为 "arg"）。
// ff:
// s:
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

// ComposeCommandLine 对给定的参数进行转义并拼接，生成适用于作为Windows命令行的结果，
// 可用于CreateProcess函数的CommandLine参数、CreateService/ChangeServiceConfig函数的BinaryPathName参数，
// 或任何使用CommandLineToArgv函数的程序。
// ff:
// args:
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
// 强制对以下字符使用引号：不仅包括MSDN文章中描述的ASCII空格和制表符，还包括ASCII控制字符。
// CommandLineToArgvW函数的文档并未说明当第一个参数不是有效的程序名时会发生什么情况，但经实测，
// 该函数似乎会舍弃未加引号的控制字符。
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
// 此引号会与我们周围的引号冲突。
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
// TODO(bcmills): 既然我们已在向切片追加内容，那么能避免 `EscapeArg` 的中间分配操作就很好了。
// 或许我们可以提取出一个 `appendEscapedArg` 函数。
		commandLine = append(commandLine, EscapeArg(arg)...)
	}
	return string(commandLine)
}

// DecomposeCommandLine 将其参数 command line 使用 CommandLineToArgv 拆分为未转义的部分，这些部分可从 GetCommandLine、QUERY_SERVICE_CONFIG 的 BinaryPathName 参数或其他传递命令行的途径获取。
// 若 commandLine 包含 NUL 字符，DecomposeCommandLine 将返回错误。
// ff:
// commandLine:
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

// CommandLineToArgv 函数解析一个 Unicode 格式的命令行字符串，并将解析得到的参数个数设置到 argc 中。
//
// 返回的内存应当使用一次 LocalFree 调用来释放。
//
// 需要注意的是，尽管 CommandLineToArgv 的返回类型表明其可以容纳 8192 个长度不超过 8192 字符的条目，
// 实际解析得到的参数数量可能超过 8192，且 CommandLineToArgvW 文档并未提及单个参数字符串长度有任何限制。
// （参见 https://go.dev/issue/63236。）
// ff:
// cmd:
// argc:
// argv:
// err:
func CommandLineToArgv(cmd *uint16, argc *int32) (argv *[8192]*[8192]uint16, err error) {
	argp, err := commandLineToArgv(cmd, argc)
	argv = (*[8192]*[8192]uint16)(unsafe.Pointer(argp))
	return argv, err
}

// ff:
// fd:
func CloseOnExec(fd Handle) {
	SetHandleInformation(Handle(fd), HANDLE_FLAG_INHERIT, 0)
}

// FullPath 获取指定文件的完整路径
// ff:
// name:
// path:
// err:
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

// NewProcThreadAttributeList 分配一个具有所请求最大属性数量的新 ProcThreadAttributeListContainer。
// ff:
// maxAttrCount:
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
// ff:
// al:
// attribute:
// value:
// size:
func (al *ProcThreadAttributeListContainer) Update(attribute uintptr, value unsafe.Pointer, size uintptr) error {
	al.pointers = append(al.pointers, value)
	return updateProcThreadAttribute(al.data, 0, attribute, value, size, nil, nil)
}

// Delete 释放 ProcThreadAttributeList 的资源。
// ff:
// al:
func (al *ProcThreadAttributeListContainer) Delete() {
	deleteProcThreadAttributeList(al.data)
	LocalFree(Handle(unsafe.Pointer(al.data)))
	al.data = nil
	al.pointers = nil
}

// List 返回将传递给 StartupInfoEx 的实际 ProcThreadAttributeList。
// ff:
// al:
func (al *ProcThreadAttributeListContainer) List() *ProcThreadAttributeList {
	return al.data
}
