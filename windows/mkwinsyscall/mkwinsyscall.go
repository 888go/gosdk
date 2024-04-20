// 版权所有 2013 The Go Authors。保留所有权利。
// 使用本源代码受 BSD 风格许可协议约束，
// 该协议可在 LICENSE 文件中找到。

/*
mkwinsyscall generates windows system call bodies

It parses all files specified on command line containing function
prototypes (like syscall_windows.go) and prints system call bodies
to standard output.

The prototypes are marked by lines beginning with "//sys" and read
like func declarations if //sys 被替换为 func，但是：

  - The parameter lists must give a name for each argument. This
    includes return parameters.

  - The parameter lists must give a type for each argument:
    the (x, y, z int) shorthand is not allowed.

  - If the return parameter is an error number, it must be named err.

  - If go func name needs to be different from its winapi dll name,
    the winapi name could be specified at the end, after "=" sign, like
    //sys LoadLibrary(libname string) (handle uint32, err error) = LoadLibraryA
// 
// 系统调用 LoadLibrary 接收一个字符串参数 libname，返回一个无符号整数类型的句柄（handle）和一个错误对象（err）。该函数等同于 LoadLibraryA。

  - Each function that returns err needs to supply a condition, that
    return value of winapi will be tested against to detect failure.
    This would set err to windows "last-error", otherwise it will be nil.
    The value can be provided at end of //sys declaration, like
    //sys LoadLibrary(libname string) (handle uint32, err error) [failretval==-1] = LoadLibraryA
// 
// 系统调用 LoadLibrary 接收一个字符串参数 libname，返回一个无符号整数类型的句柄（handle）和一个错误对象（err）。若函数执行失败，返回的句柄值应为-1。该函数实际绑定到名为 LoadLibraryA 的系统调用。
    and is [failretval==0] by default.

  - If the function name ends in a "?", then the function not existing is non-
    fatal, and an error will be returned instead of panicking.

Usage:

	mkwinsyscall [flags] [path ...]

The flags are:

	-output
		Specify output file name (outputs to console if blank).
	-trace
		Generate print statement after every syscall.
*/
package main

import (
	"bufio"
	"bytes"
	"errors"
	"flag"
	"fmt"
	"go/format"
	"go/parser"
	"go/token"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"text/template"
)

var (
	filename       = flag.String("output", "", "output file name (standard output if omitted)")
	printTraceFlag = flag.Bool("trace", false, "generate print statement after every syscall")
	systemDLL      = flag.Bool("systemdll", true, "whether all DLLs should be loaded from the Windows system directory")
)

func trim(s string) string {
	return strings.Trim(s, " \t")
}

var packageName string

func packagename() string {
	return packageName
}

func windowsdot() string {
	if packageName == "windows" {
		return ""
	}
	return "windows."
}

func syscalldot() string {
	if packageName == "syscall" {
		return ""
	}
	return "syscall."
}

// Param 是函数参数
type Param struct {
	Name      string
	Type      string
	fn        *Fn
	tmpVarIdx int
}

// tmpVar 返回一个临时变量名，该变量将用于在系统调用期间表示 p。
func (p *Param) tmpVar() string {
	if p.tmpVarIdx < 0 {
		p.tmpVarIdx = p.fn.curTmpVarIdx
		p.fn.curTmpVarIdx++
	}
	return fmt.Sprintf("_p%d", p.tmpVarIdx)
}

// BoolTmpVarCode 返回用于布尔临时变量的源代码
func (p *Param) BoolTmpVarCode() string {
	const code = `var %[1]s uint32
	if %[2]s {
		%[1]s = 1
	}`
	return fmt.Sprintf(code, p.tmpVar(), p.Name)
}

// BoolPointerTmpVarCode 返回用于布尔临时变量的源代码
func (p *Param) BoolPointerTmpVarCode() string {
	const code = `var %[1]s uint32
	if *%[2]s {
		%[1]s = 1
	}`
	return fmt.Sprintf(code, p.tmpVar(), p.Name)
}

// SliceTmpVarCode 返回用于切片临时变量的源代码
func (p *Param) SliceTmpVarCode() string {
	const code = `var %s *%s
	if len(%s) > 0 {
		%s = &%s[0]
	}`
	tmp := p.tmpVar()
	return fmt.Sprintf(code, tmp, p.Type[2:], p.Name, tmp, p.Name)
}

// StringTmpVarCode 返回字符串临时变量的源代码
func (p *Param) StringTmpVarCode() string {
	errvar := p.fn.Rets.ErrorVarName()
	if errvar == "" {
		errvar = "_"
	}
	tmp := p.tmpVar()
	const code = `var %s %s
	%s, %s = %s(%s)`
	s := fmt.Sprintf(code, tmp, p.fn.StrconvType(), tmp, errvar, p.fn.StrconvFunc(), p.Name)
	if errvar == "-" {
		return s
	}
	const morecode = `
	if %s != nil {
		return
	}`
	return s + fmt.Sprintf(morecode, errvar)
}

// TmpVarCode 返回用于临时变量的源代码
func (p *Param) TmpVarCode() string {
	switch {
	case p.Type == "bool":
		return p.BoolTmpVarCode()
	case p.Type == "*bool":
		return p.BoolPointerTmpVarCode()
	case strings.HasPrefix(p.Type, "[]"):
		return p.SliceTmpVarCode()
	default:
		return ""
	}
}

// TmpVarReadbackCode 返回用于将临时变量回读到原始变量的源代码
func (p *Param) TmpVarReadbackCode() string {
	switch {
	case p.Type == "*bool":
		return fmt.Sprintf("*%s = %s != 0", p.Name, p.tmpVar())
	default:
		return ""
	}
}

// TmpVarHelperCode 返回助手临时变量的源代码
func (p *Param) TmpVarHelperCode() string {
	if p.Type != "string" {
		return ""
	}
	return p.StringTmpVarCode()
}

// SyscallArgList 返回表示 syscall 中 p 参数的源代码片段。
// 切片会被转化为两个系统调用参数：指向首个元素的指针和长度。
func (p *Param) SyscallArgList() []string {
	t := p.HelperType()
	var s string
	switch {
	case t == "*bool":
		s = fmt.Sprintf("unsafe.Pointer(&%s)", p.tmpVar())
	case t[0] == '*':
		s = fmt.Sprintf("unsafe.Pointer(%s)", p.Name)
	case t == "bool":
		s = p.tmpVar()
	case strings.HasPrefix(t, "[]"):
		return []string{
			fmt.Sprintf("uintptr(unsafe.Pointer(%s))", p.tmpVar()),
			fmt.Sprintf("uintptr(len(%s))", p.Name),
		}
	default:
		s = p.Name
	}
	return []string{fmt.Sprintf("uintptr(%s)", s)}
}

// IsError 判断参数p是否用于返回错误
func (p *Param) IsError() bool {
	return p.Name == "err" && p.Type == "error"
}

// HelperType 返回在辅助函数中使用的参数 p 的类型
func (p *Param) HelperType() string {
	if p.Type == "string" {
		return p.fn.StrconvType()
	}
	return p.Type
}

// join 函数通过 sep 分隔符将参数 ps 连接成一个字符串。
// 在进行拼接之前，每个参数都会通过应用 fn 函数转化为字符串形式。
func join(ps []*Param, fn func(*Param) string, sep string) string {
	if len(ps) == 0 {
		return ""
	}
	a := make([]string, 0)
	for _, p := range ps {
		a = append(a, fn(p))
	}
	return strings.Join(a, sep)
}

// Rets 描述函数返回参数。
type Rets struct {
	Name          string
	Type          string
	ReturnsError  bool
	FailCond      string
	fnMaybeAbsent bool
}

// ErrorVarName 为 r 返回错误变量名。
func (r *Rets) ErrorVarName() string {
	if r.ReturnsError {
		return "err"
	}
	if r.Type == "error" {
		return r.Name
	}
	return ""
}

// ToParams 将 r 转换为 []*Param 类型的切片。
func (r *Rets) ToParams() []*Param {
	ps := make([]*Param, 0)
	if len(r.Name) > 0 {
		ps = append(ps, &Param{Name: r.Name, Type: r.Type})
	}
	if r.ReturnsError {
		ps = append(ps, &Param{Name: "err", Type: "error"})
	}
	return ps
}

// List 返回 syscall 返回参数的源代码
func (r *Rets) List() string {
	s := join(r.ToParams(), func(p *Param) string { return p.Name + " " + p.Type }, ", ")
	if len(s) > 0 {
		s = "(" + s + ")"
	} else if r.fnMaybeAbsent {
		s = "(err error)"
	}
	return s
}

// PrintList 返回与系统调用返回值对应的跟踪打印部分的源代码
func (r *Rets) PrintList() string {
	return join(r.ToParams(), func(p *Param) string { return fmt.Sprintf(`"%s=", %s, `, p.Name, p.Name) }, `", ", `)
}

// SetReturnValuesCode 返回用于接收系统调用返回值的源代码
func (r *Rets) SetReturnValuesCode() string {
	if r.Name == "" && !r.ReturnsError {
		return ""
	}
	retvar := "r0"
	if r.Name == "" {
		retvar = "r1"
	}
	errvar := "_"
	if r.ReturnsError {
		errvar = "e1"
	}
	return fmt.Sprintf("%s, _, %s := ", retvar, errvar)
}

func (r *Rets) useLongHandleErrorCode(retvar string) string {
	const code = `if %s {
		err = errnoErr(e1)
	}`
	cond := retvar + " == 0"
	if r.FailCond != "" {
		cond = strings.Replace(r.FailCond, "failretval", retvar, 1)
	}
	return fmt.Sprintf(code, cond)
}

// SetErrorCode 返回设置返回参数的源代码
func (r *Rets) SetErrorCode() string {
	const code = `if r0 != 0 {
		%s = %sErrno(r0)
	}`
	const ntstatus = `if r0 != 0 {
		ntstatus = %sNTStatus(r0)
	}`
	if r.Name == "" && !r.ReturnsError {
		return ""
	}
	if r.Name == "" {
		return r.useLongHandleErrorCode("r1")
	}
	if r.Type == "error" && r.Name == "ntstatus" {
		return fmt.Sprintf(ntstatus, windowsdot())
	}
	if r.Type == "error" {
		return fmt.Sprintf(code, r.Name, syscalldot())
	}
	s := ""
	switch {
	case r.Type[0] == '*':
		s = fmt.Sprintf("%s = (%s)(unsafe.Pointer(r0))", r.Name, r.Type)
	case r.Type == "bool":
		s = fmt.Sprintf("%s = r0 != 0", r.Name)
	default:
		s = fmt.Sprintf("%s = %s(r0)", r.Name, r.Type)
	}
	if !r.ReturnsError {
		return s
	}
	return s + "\n\t" + r.useLongHandleErrorCode(r.Name)
}

// Fn 描述系统调用函数。
type Fn struct {
	Name        string
	Params      []*Param
	Rets        *Rets
	PrintTrace  bool
	dllname     string
	dllfuncname string
	src         string
	// TODO：移除该字段，直接使用参数索引代替
	curTmpVarIdx int // 确保临时变量具有唯一名称
}

// extractParams 用于解析 s 以提取函数参数。
func extractParams(s string, f *Fn) ([]*Param, error) {
	s = trim(s)
	if s == "" {
		return nil, nil
	}
	a := strings.Split(s, ",")
	ps := make([]*Param, len(a))
	for i := range ps {
		s2 := trim(a[i])
		b := strings.Split(s2, " ")
		if len(b) != 2 {
			b = strings.Split(s2, "\t")
			if len(b) != 2 {
				return nil, errors.New("Could not extract function parameter from \"" + s2 + "\"")
			}
		}
		ps[i] = &Param{
			Name:      trim(b[0]),
			Type:      trim(b[1]),
			fn:        f,
			tmpVarIdx: -1,
		}
	}
	return ps, nil
}

// extractSection 从字符串 s 中提取起始于 start 之后、终止于 end 之前的文本。返回值 found 用于指示操作是否成功，而 prefix、body 和 suffix 将分别包含 s 字符串中对应的各部分。
func extractSection(s string, start, end rune) (prefix, body, suffix string, found bool) {
	s = trim(s)
	if strings.HasPrefix(s, string(start)) {
		// no prefix
		body = s[1:]
	} else {
		a := strings.SplitN(s, string(start), 2)
		if len(a) != 2 {
			return "", "", s, false
		}
		prefix = a[0]
		body = a[1]
	}
	a := strings.SplitN(body, string(end), 2)
	if len(a) != 2 {
		return "", "", "", false
	}
	return prefix, a[0], a[1], true
}

// newFn 用于解析字符串 s，并返回已创建的函数 Fn。
func newFn(s string) (*Fn, error) {
	s = trim(s)
	f := &Fn{
		Rets:       &Rets{},
		src:        s,
		PrintTrace: *printTraceFlag,
	}
	// function name and args
	prefix, body, s, found := extractSection(s, '(', ')')
	if !found || prefix == "" {
		return nil, errors.New("Could not extract function name and parameters from \"" + f.src + "\"")
	}
	f.Name = prefix
	var err error
	f.Params, err = extractParams(body, f)
	if err != nil {
		return nil, err
	}
	// return values
	_, body, s, found = extractSection(s, '(', ')')
	if found {
		r, err := extractParams(body, f)
		if err != nil {
			return nil, err
		}
		switch len(r) {
		case 0:
		case 1:
			if r[0].IsError() {
				f.Rets.ReturnsError = true
			} else {
				f.Rets.Name = r[0].Name
				f.Rets.Type = r[0].Type
			}
		case 2:
			if !r[1].IsError() {
				return nil, errors.New("Only last windows error is allowed as second return value in \"" + f.src + "\"")
			}
			f.Rets.ReturnsError = true
			f.Rets.Name = r[0].Name
			f.Rets.Type = r[0].Type
		default:
			return nil, errors.New("Too many return values in \"" + f.src + "\"")
		}
	}
	// fail condition
	_, body, s, found = extractSection(s, '[', ']')
	if found {
		f.Rets.FailCond = body
	}
	// dll 和 dll 函数名称
	s = trim(s)
	if s == "" {
		return f, nil
	}
	if !strings.HasPrefix(s, "=") {
		return nil, errors.New("Could not extract dll name from \"" + f.src + "\"")
	}
	s = trim(s[1:])
	if i := strings.LastIndex(s, "."); i >= 0 {
		f.dllname = s[:i]
		f.dllfuncname = s[i+1:]
	} else {
		f.dllfuncname = s
	}
	if f.dllfuncname == "" {
		return nil, fmt.Errorf("function name is not specified in %q", s)
	}
	if n := f.dllfuncname; strings.HasSuffix(n, "?") {
		f.dllfuncname = n[:len(n)-1]
		f.Rets.fnMaybeAbsent = true
	}
	return f, nil
}

// DLLName 返回函数f对应的DLL名称。
func (f *Fn) DLLName() string {
	if f.dllname == "" {
		return "kernel32"
	}
	return f.dllname
}

// DLLVar 返回一个表示 DLLName 的有效 Go 标识符。
func (f *Fn) DLLVar() string {
	id := strings.Map(func(r rune) rune {
		switch r {
		case '.', '-':
			return '_'
		default:
			return r
		}
	}, f.DLLName())
	if !token.IsIdentifier(id) {
		panic(fmt.Errorf("could not create Go identifier for DLLName %q", f.DLLName()))
	}
	return id
}

// DLLFuncName 为函数 f 返回 DLL 函数名。
func (f *Fn) DLLFuncName() string {
	if f.dllfuncname == "" {
		return f.Name
	}
	return f.dllfuncname
}

// ParamList 返回函数f参数的源代码
func (f *Fn) ParamList() string {
	return join(f.Params, func(p *Param) string { return p.Name + " " + p.Type }, ", ")
}

// HelperParamList 返回辅助函数f参数的源代码
func (f *Fn) HelperParamList() string {
	return join(f.Params, func(p *Param) string { return p.Name + " " + p.HelperType() }, ", ")
}

// ParamPrintList 返回与系统调用输入参数对应的跟踪打印部分的源代码
func (f *Fn) ParamPrintList() string {
	return join(f.Params, func(p *Param) string { return fmt.Sprintf(`"%s=", %s, `, p.Name, p.Name) }, `", ", `)
}

// ParamCount 返回函数f的系统调用参数数量
func (f *Fn) ParamCount() int {
	n := 0
	for _, p := range f.Params {
		n += len(p.SyscallArgList())
	}
	return n
}

// SyscallParamCount 用于确定应使用 Syscall、Syscall6、Syscall9 等中的哪个版本。它返回相应 SyscallX 函数的参数个数。
func (f *Fn) SyscallParamCount() int {
	n := f.ParamCount()
	switch {
	case n <= 3:
		return 3
	case n <= 6:
		return 6
	case n <= 9:
		return 9
	case n <= 12:
		return 12
	case n <= 15:
		return 15
	case n <= 42: // current SyscallN limit
		return n
	default:
		panic("too many arguments to system call")
	}
}

// Syscall 确定应使用哪个 SyscallX 函数来处理函数 f。
func (f *Fn) Syscall() string {
	c := f.SyscallParamCount()
	if c == 3 {
		return syscalldot() + "Syscall"
	}
	if c > 15 {
		return syscalldot() + "SyscallN"
	}
	return syscalldot() + "Syscall" + strconv.Itoa(c)
}

// SyscallParamList 为函数 f 生成用于 SyscallX 的参数源代码
func (f *Fn) SyscallParamList() string {
	a := make([]string, 0)
	for _, p := range f.Params {
		a = append(a, p.SyscallArgList()...)
	}
	for len(a) < f.SyscallParamCount() {
		a = append(a, "0")
	}
	return strings.Join(a, ", ")
}

// HelperCallParamList 返回对函数f辅助方法的调用源代码
func (f *Fn) HelperCallParamList() string {
	a := make([]string, 0, len(f.Params))
	for _, p := range f.Params {
		s := p.Name
		if p.Type == "string" {
			s = p.tmpVar()
		}
		a = append(a, s)
	}
	return strings.Join(a, ", ")
}

// MaybeAbsent 返回用于处理可能不可用的函数的源代码
func (p *Fn) MaybeAbsent() string {
	if !p.Rets.fnMaybeAbsent {
		return ""
	}
	const code = `%[1]s = proc%[2]s.Find()
	if %[1]s != nil {
		return
	}`
	errorVar := p.Rets.ErrorVarName()
	if errorVar == "" {
		errorVar = "err"
	}
	return fmt.Sprintf(code, errorVar, p.DLLFuncName())
}

// IsUTF16 为真，当 f 为 W（utf16）函数时。对于所有 A（ascii）函数，其为假。
func (f *Fn) IsUTF16() bool {
	s := f.DLLFuncName()
	return s[len(s)-1] == 'W'
}

// StrconvFunc 返回 Go 字符串到操作系统字符串函数的名称，该函数对应于 f。
func (f *Fn) StrconvFunc() string {
	if f.IsUTF16() {
		return syscalldot() + "UTF16PtrFromString"
	}
	return syscalldot() + "BytePtrFromString"
}

// StrconvType 返回用于操作系统字符串表示 f 的 Go 类型名称
func (f *Fn) StrconvType() string {
	if f.IsUTF16() {
		return "*uint16"
	}
	return "*byte"
}

// HasStringParam为真，当f至少有一个字符串参数时。
// 否则为假。
func (f *Fn) HasStringParam() bool {
	for _, p := range f.Params {
		if p.Type == "string" {
			return true
		}
	}
	return false
}

// HelperName 返回函数f的辅助函数名
func (f *Fn) HelperName() string {
	if !f.HasStringParam() {
		return f.Name
	}
	return "_" + f.Name
}

// DLL 是一个 DLL 文件名，以及一个在 Go 标识符中有效的字符串，该字符串应被用于命名引用该 DLL 的变量。
type DLL struct {
	Name string
	Var  string
}

// 源文件和函数
type Source struct {
	Funcs           []*Fn
	DLLFuncNames    []*Fn
	Files           []string
	StdLibImports   []string
	ExternalImports []string
}

func (src *Source) Import(pkg string) {
	src.StdLibImports = append(src.StdLibImports, pkg)
	sort.Strings(src.StdLibImports)
}

func (src *Source) ExternalImport(pkg string) {
	src.ExternalImports = append(src.ExternalImports, pkg)
	sort.Strings(src.ExternalImports)
}

// ParseFiles 解析 fs 中列出的文件，并从 sys 注释中提取所有 syscall 函数。若解析成功，返回包含源文件和函数集合的 *Source 对象。
func ParseFiles(fs []string) (*Source, error) {
	src := &Source{
		Funcs: make([]*Fn, 0),
		Files: make([]string, 0),
		StdLibImports: []string{
			"unsafe",
		},
		ExternalImports: make([]string, 0),
	}
	for _, file := range fs {
		if err := src.ParseFile(file); err != nil {
			return nil, err
		}
	}
	src.DLLFuncNames = make([]*Fn, 0, len(src.Funcs))
	uniq := make(map[string]bool, len(src.Funcs))
	for _, fn := range src.Funcs {
		name := fn.DLLFuncName()
		if !uniq[name] {
			src.DLLFuncNames = append(src.DLLFuncNames, fn)
			uniq[name] = true
		}
	}
	return src, nil
}

// DLLs 返回给定源集合src的dll名称。
func (src *Source) DLLs() []DLL {
	uniq := make(map[string]bool)
	r := make([]DLL, 0)
	for _, f := range src.Funcs {
		id := f.DLLVar()
		if _, found := uniq[id]; !found {
			uniq[id] = true
			r = append(r, DLL{f.DLLName(), id})
		}
	}
	sort.Slice(r, func(i, j int) bool {
		return r[i].Var < r[j].Var
	})
	return r
}

// ParseFile 向源代码集合 src 中添加额外的文件路径。
func (src *Source) ParseFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	s := bufio.NewScanner(file)
	for s.Scan() {
		t := trim(s.Text())
		if len(t) < 7 {
			continue
		}
		if !strings.HasPrefix(t, "//sys") {
			continue
		}
		t = t[5:]
		if !(t[0] == ' ' || t[0] == '\t') {
			continue
		}
		f, err := newFn(t[1:])
		if err != nil {
			return err
		}
		src.Funcs = append(src.Funcs, f)
	}
	if err := s.Err(); err != nil {
		return err
	}
	src.Files = append(src.Files, path)
	sort.Slice(src.Funcs, func(i, j int) bool {
		fi, fj := src.Funcs[i], src.Funcs[j]
		if fi.DLLName() == fj.DLLName() {
			return fi.DLLFuncName() < fj.DLLFuncName()
		}
		return fi.DLLName() < fj.DLLName()
	})

	// get package name
	fset := token.NewFileSet()
	_, err = file.Seek(0, 0)
	if err != nil {
		return err
	}
	pkg, err := parser.ParseFile(fset, "", file, parser.PackageClauseOnly)
	if err != nil {
		return err
	}
	packageName = pkg.Name.Name

	return nil
}

// IsStdRepo 判断 src 是否属于标准库。
func (src *Source) IsStdRepo() (bool, error) {
	if len(src.Files) == 0 {
		return false, errors.New("no input files provided")
	}
	abspath, err := filepath.Abs(src.Files[0])
	if err != nil {
		return false, err
	}
	goroot := runtime.GOROOT()
	if runtime.GOOS == "windows" {
		abspath = strings.ToLower(abspath)
		goroot = strings.ToLower(goroot)
	}
	sep := string(os.PathSeparator)
	if !strings.HasSuffix(goroot, sep) {
		goroot += sep
	}
	return strings.HasPrefix(abspath, goroot), nil
}

// 从源集合src生成输出源文件
func (src *Source) Generate(w io.Writer) error {
	const (
		pkgStd         = iota // 标准库中的任何包
		pkgXSysWindows        // x/sys/windows package
		pkgOther
	)
	isStdRepo, err := src.IsStdRepo()
	if err != nil {
		return err
	}
	var pkgtype int
	switch {
	case isStdRepo:
		pkgtype = pkgStd
	case packageName == "windows":
		// TODO：这需要比仅仅使用包名更合理的逻辑
		pkgtype = pkgXSysWindows
	default:
		pkgtype = pkgOther
	}
	if *systemDLL {
		switch pkgtype {
		case pkgStd:
			src.Import("internal/syscall/windows/sysdll")
		case pkgXSysWindows:
		default:
			src.ExternalImport("golang.org/x/sys/windows")
		}
	}
	if packageName != "syscall" {
		src.Import("syscall")
	}
	funcMap := template.FuncMap{
		"packagename": packagename,
		"syscalldot":  syscalldot,
		"newlazydll": func(dll string) string {
			arg := "\"" + dll + ".dll\""
			if !*systemDLL {
				return syscalldot() + "NewLazyDLL(" + arg + ")"
			}
			switch pkgtype {
			case pkgStd:
				return syscalldot() + "NewLazyDLL(sysdll.Add(" + arg + "))"
			case pkgXSysWindows:
				return "NewLazySystemDLL(" + arg + ")"
			default:
				return "windows.NewLazySystemDLL(" + arg + ")"
			}
		},
	}
	t := template.Must(template.New("main").Funcs(funcMap).Parse(srcTemplate))
	err = t.Execute(w, src)
	if err != nil {
		return errors.New("Failed to execute template: " + err.Error())
	}
	return nil
}

func writeTempSourceFile(data []byte) (string, error) {
	f, err := os.CreateTemp("", "mkwinsyscall-generated-*.go")
	if err != nil {
		return "", err
	}
	_, err = f.Write(data)
	if closeErr := f.Close(); err == nil {
		err = closeErr
	}
	if err != nil {
		os.Remove(f.Name()) // best effort
		return "", err
	}
	return f.Name(), nil
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: mkwinsyscall [flags] [path ...]\n")
	flag.PrintDefaults()
	os.Exit(1)
}

func main() {
	flag.Usage = usage
	flag.Parse()
	if len(flag.Args()) <= 0 {
		fmt.Fprintf(os.Stderr, "no files to parse provided\n")
		usage()
	}

	src, err := ParseFiles(flag.Args())
	if err != nil {
		log.Fatal(err)
	}

	var buf bytes.Buffer
	if err := src.Generate(&buf); err != nil {
		log.Fatal(err)
	}

	data, err := format.Source(buf.Bytes())
	if err != nil {
		log.Printf("failed to format source: %v", err)
		f, err := writeTempSourceFile(buf.Bytes())
		if err != nil {
			log.Fatalf("failed to write unformatted source to file: %v", err)
		}
		log.Fatalf("for diagnosis, wrote unformatted source to %v", f)
	}
	if *filename == "" {
		_, err = os.Stdout.Write(data)
	} else {
		err = os.WriteFile(*filename, data, 0644)
	}
	if err != nil {
		log.Fatal(err)
	}
}

// 待办事项：在以下模板中使用println进行打印
const srcTemplate = `

{{define "main"}}// 代码由'go generate'命令生成；请勿编辑。

package {{packagename}}

import (
{{range .StdLibImports}}"{{.}}"
{{end}}

{{range .ExternalImports}}"{{.}}"
{{end}}
)

var _ unsafe.Pointer

// 对于常见的Errno值，仅一次性完成接口分配。
const (
	errnoERROR_IO_PENDING = 997
)

var (
	errERROR_IO_PENDING error = {{syscalldot}}Errno(errnoERROR_IO_PENDING)
	errERROR_EINVAL error     = {{syscalldot}}EINVAL
)

// errnoErr 返回常见的装箱 Errno 值，以防止运行时的分配。
func errnoErr(e {{syscalldot}}Errno) error {
	switch e {
	case 0:
		return errERROR_EINVAL
	case errnoERROR_IO_PENDING:
		return errERROR_IO_PENDING
	}
// TODO: 在收集到Windows上常见的错误值数据后，此处应添加更多内容。（或许在运行all.bat时？）
	return e
}

var (
{{template "dlls" .}}
{{template "funcnames" .}})
{{range .Funcs}}{{if .HasStringParam}}{{template "helperbody" .}}{{end}}{{template "funcbody" .}}{{end}}
{{end}}

{{/* help functions */}}

{{define "dlls"}}{{range .DLLs}}	mod{{.Var}} = {{newlazydll .Name}}
{{end}}{{end}}

{{define "funcnames"}}{{range .DLLFuncNames}}	proc{{.DLLFuncName}} = mod{{.DLLVar}}.NewProc("{{.DLLFuncName}}")
{{end}}{{end}}

{{define "helperbody"}}
func {{.Name}}({{.ParamList}}) {{template "results" .}}{
{{template "helpertmpvars" .}}	return {{.HelperName}}({{.HelperCallParamList}})
}
{{end}}

{{define "funcbody"}}
func {{.HelperName}}({{.HelperParamList}}) {{template "results" .}}{
{{template "maybeabsent" .}}	{{template "tmpvars" .}}	{{template "syscall" .}}	{{template "tmpvarsreadback" .}}
{{template "seterror" .}}{{template "printtrace" .}}	return
}
{{end}}

{{define "helpertmpvars"}}{{range .Params}}{{if .TmpVarHelperCode}}	{{.TmpVarHelperCode}}
{{end}}{{end}}{{end}}

{{define "maybeabsent"}}{{if .MaybeAbsent}}{{.MaybeAbsent}}
{{end}}{{end}}

{{define "tmpvars"}}{{range .Params}}{{if .TmpVarCode}}	{{.TmpVarCode}}
{{end}}{{end}}{{end}}

{{define "results"}}{{if .Rets.List}}{{.Rets.List}} {{end}}{{end}}

{{define "syscall"}}{{.Rets.SetReturnValuesCode}}{{.Syscall}}(proc{{.DLLFuncName}}.Addr(),{{if le .ParamCount 15}} {{.ParamCount}},{{end}} {{.SyscallParamList}}){{end}}

{{define "tmpvarsreadback"}}{{range .Params}}{{if .TmpVarReadbackCode}}
{{.TmpVarReadbackCode}}{{end}}{{end}}{{end}}

{{define "seterror"}}{{if .Rets.SetErrorCode}}	{{.Rets.SetErrorCode}}
{{end}}{{end}}

{{define "printtrace"}}{{if .PrintTrace}}	print("SYSCALL: {{.Name}}(", {{.ParamPrintList}}") (", {{.Rets.PrintList}}")\n")
{{end}}{{end}}

`
