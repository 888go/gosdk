// 版权所有 2011 Go 作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可协议约束，
// 该协议可在 LICENSE 文件中找到。

package windows//bm:win类

import (
	"sync"
	"sync/atomic"
	"syscall"
	"unsafe"
)

// 我们需要使用Go运行时中的LoadLibrary和GetProcAddress，因为这些符号由系统链接器加载，对于动态加载额外符号是必需的。请注意，在Go运行时中，这些函数返回syscall.Handle和syscall.Errno，但事实上它们与windows.Handle和windows.Errno相同，我们打算保持它们的一致性。

//go:linkname syscall_loadlibrary syscall.loadlibrary
func syscall_loadlibrary(filename *uint16) (handle Handle, err Errno)

//go:linkname syscall_getprocaddress syscall.getprocaddress
func syscall_getprocaddress(handle Handle, procname *uint8) (proc uintptr, err Errno)

// DLLError 描述了 DLL 加载失败的原因。
type DLLError struct {
	Err     error
	ObjName string
	Msg     string
}

// 翻译提示:func  (e  *DLL错误)  错误信息()  字符串  {}

// ff:取错误文本
func (e *DLLError) Error() string { return e.Msg }

// 翻译提示:func  (e  *DLL错误)  获取底层错误()  错误  {}

// ff:取错误
func (e *DLLError) Unwrap() error { return e.Err }

// DLL 结构体实现了对单一 DLL 的访问功能。
type DLL struct {
	Name   string
	Handle Handle
}

// LoadDLL 将 DLL 文件加载到内存中。
//
// 警告：在没有使用绝对路径的情况下使用 LoadDLL，可能导致 DLL 预加载攻击。为了安全地加载系统 DLL，请使用 System 参数设为 true 的 LazyDLL，或直接使用 LoadLibraryEx。

// ff:DLL加载
// name:dll名称
// dll:
// err:错误
func LoadDLL(name string) (dll *DLL, err error) {
	namep, err := UTF16PtrFromString(name)
	if err != nil {
		return nil, err
	}
	h, e := syscall_loadlibrary(namep)
	if e != 0 {
		return nil, &DLLError{
			Err:     e,
			ObjName: name,
			Msg:     "Failed to load " + name + ": " + e.Error(),
		}
	}
	d := &DLL{
		Name:   name,
		Handle: h,
	}
	return d, nil
}

// MustLoadDLL与LoadDLL类似，但若加载操作失败，则会引发恐慌。

// ff:DLL加载PANI
// name:dll名称
func MustLoadDLL(name string) *DLL {
	d, e := LoadDLL(name)
	if e != nil {
		panic(e)
	}
	return d
}

// FindProc在DLL d中搜索名为name的程序，并在找到时返回*Proc。如果搜索失败，它将返回一个错误。

// ff:查找命令
// name:命令名
// proc:
// err:错误
func (d *DLL) FindProc(name string) (proc *Proc, err error) {
	namep, err := BytePtrFromString(name)
	if err != nil {
		return nil, err
	}
	a, e := syscall_getprocaddress(d.Handle, namep)
	if e != 0 {
		return nil, &DLLError{
			Err:     e,
			ObjName: name,
			Msg:     "Failed to find " + name + " procedure in " + d.Name + ": " + e.Error(),
		}
	}
	p := &Proc{
		Dll:  d,
		Name: name,
		addr: a,
	}
	return p, nil
}

// MustFindProc 与 FindProc 类似，但在搜索失败时会引发 panic。

// ff:查找命令PANI
// name:命令名
func (d *DLL) MustFindProc(name string) *Proc {
	p, e := d.FindProc(name)
	if e != nil {
		panic(e)
	}
	return p
}

// FindProcByOrdinal 在 DLL d 中按序号搜索过程，并在找到时返回 *Proc。如果搜索失败，则返回错误。

// ff:查找命令并按序数
// ordinal:序数
// proc:
// err:错误
func (d *DLL) FindProcByOrdinal(ordinal uintptr) (proc *Proc, err error) {
	a, e := GetProcAddressByOrdinal(d.Handle, ordinal)
	name := "#" + itoa(int(ordinal))
	if e != nil {
		return nil, &DLLError{
			Err:     e,
			ObjName: name,
			Msg:     "Failed to find " + name + " procedure in " + d.Name + ": " + e.Error(),
		}
	}
	p := &Proc{
		Dll:  d,
		Name: name,
		addr: a,
	}
	return p, nil
}

// MustFindProcByOrdinal 与 FindProcByOrdinal 类似，但若搜索失败则触发 panic。

// ff:查找命令并按序数PANI
// ordinal:序数
func (d *DLL) MustFindProcByOrdinal(ordinal uintptr) *Proc {
	p, e := d.FindProcByOrdinal(ordinal)
	if e != nil {
		panic(e)
	}
	return p
}

// Release 从内存中卸载 DLL d。

// ff:卸载
// err:错误
func (d *DLL) Release() (err error) {
	return FreeLibrary(d.Handle)
}

// Proc 实现了对 DLL 中某个过程的访问。
type Proc struct {
	Dll  *DLL
	Name string
	addr uintptr
}

// Addr 返回由 p 表示的程序的地址。
// 返回值可以传递给 Syscall 以运行该程序。

// ff:取命令地址
func (p *Proc) Addr() uintptr {
	return p.addr
}

//go:uintptrescapes

// Call 函数使用参数 a 执行过程 p。如果提供的参数超过 15 个，将会引发 panic。
//
// 返回的错误始终为非 nil，由 GetLastError 的结果构建而成。调用者必须首先根据被调用函数的具体语义检查主返回值以判断是否发生错误，然后才能参考该错误。该错误将确保包含 windows.Errno。

// ff:调用
// a:命令参数s
// r1:
// r2:
// lastErr:最后一个错误
func (p *Proc) Call(a ...uintptr) (r1, r2 uintptr, lastErr error) {
	switch len(a) {
	case 0:
		return syscall.Syscall(p.Addr(), uintptr(len(a)), 0, 0, 0)
	case 1:
		return syscall.Syscall(p.Addr(), uintptr(len(a)), a[0], 0, 0)
	case 2:
		return syscall.Syscall(p.Addr(), uintptr(len(a)), a[0], a[1], 0)
	case 3:
		return syscall.Syscall(p.Addr(), uintptr(len(a)), a[0], a[1], a[2])
	case 4:
		return syscall.Syscall6(p.Addr(), uintptr(len(a)), a[0], a[1], a[2], a[3], 0, 0)
	case 5:
		return syscall.Syscall6(p.Addr(), uintptr(len(a)), a[0], a[1], a[2], a[3], a[4], 0)
	case 6:
		return syscall.Syscall6(p.Addr(), uintptr(len(a)), a[0], a[1], a[2], a[3], a[4], a[5])
	case 7:
		return syscall.Syscall9(p.Addr(), uintptr(len(a)), a[0], a[1], a[2], a[3], a[4], a[5], a[6], 0, 0)
	case 8:
		return syscall.Syscall9(p.Addr(), uintptr(len(a)), a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], 0)
	case 9:
		return syscall.Syscall9(p.Addr(), uintptr(len(a)), a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], a[8])
	case 10:
		return syscall.Syscall12(p.Addr(), uintptr(len(a)), a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], a[8], a[9], 0, 0)
	case 11:
		return syscall.Syscall12(p.Addr(), uintptr(len(a)), a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], a[8], a[9], a[10], 0)
	case 12:
		return syscall.Syscall12(p.Addr(), uintptr(len(a)), a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], a[8], a[9], a[10], a[11])
	case 13:
		return syscall.Syscall15(p.Addr(), uintptr(len(a)), a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], a[8], a[9], a[10], a[11], a[12], 0, 0)
	case 14:
		return syscall.Syscall15(p.Addr(), uintptr(len(a)), a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], a[8], a[9], a[10], a[11], a[12], a[13], 0)
	case 15:
		return syscall.Syscall15(p.Addr(), uintptr(len(a)), a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], a[8], a[9], a[10], a[11], a[12], a[13], a[14])
	default:
		panic("Call " + p.Name + " with too many arguments " + itoa(len(a)) + ".")
	}
}

// LazyDLL 实现了对单一 DLL 的访问。它将延迟 DLL 的加载，直到第一次调用其 Handle 方法或其某个 LazyProc 的 Addr 方法时为止。
type LazyDLL struct {
	Name string

// System 指定是否必须从 Windows 系统目录加载 DLL，从而绕过常规的 DLL 搜索路径。
	System bool

	mu  sync.Mutex
	dll *DLL // 一旦DLL加载，即为非空
}

// Load 将 DLL 文件 d.Name 载入内存。如果加载失败，将返回错误。
// 若 DLL 已经被载入内存，Load 不会尝试再次加载。

// ff:加载
func (d *LazyDLL) Load() error {
// 非竞态版本的：
// if d.dll != nil {
	if atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&d.dll))) != nil {
		return nil
	}
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.dll != nil {
		return nil
	}

// kernel32.dll 是特殊的，因为它正是 LoadLibraryEx 函数的来源。
// 内核已经对其名称进行了特例处理，因此它始终
// 从 system32 加载。
	var dll *DLL
	var err error
	if d.Name == "kernel32.dll" {
		dll, err = LoadDLL(d.Name)
	} else {
		dll, err = loadLibraryEx(d.Name, d.System)
	}
	if err != nil {
		return err
	}

// 非竞态版本的：
// d.dll = dll
	atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(&d.dll)), unsafe.Pointer(dll))
	return nil
}

// mustLoad 与 Load 类似，但当搜索失败时会触发 panic。
func (d *LazyDLL) mustLoad() {
	e := d.Load()
	if e != nil {
		panic(e)
	}
}

// Handle 返回 d 的模块句柄。

// ff:取模块地址
func (d *LazyDLL) Handle() uintptr {
	d.mustLoad()
	return uintptr(d.dll.Handle)
}

// NewProc 返回一个用于访问 DLL d 中指定名称过程的 LazyProc。

// ff:创建命令对象
// name:命令名
func (d *LazyDLL) NewProc(name string) *LazyProc {
	return &LazyProc{l: d, Name: name}
}

// NewLazyDLL 创建与 DLL 文件关联的新 LazyDLL。

// ff:DLL惰性加载
// name:dll名称
func NewLazyDLL(name string) *LazyDLL {
	return &LazyDLL{Name: name}
}

// NewLazySystemDLL类似于NewLazyDLL，但仅当name为基名（如"advapi32.dll"）时，
// 才会在Windows系统目录中搜索该DLL。

// ff:DLL惰性加载并从System
// name:dll名称
func NewLazySystemDLL(name string) *LazyDLL {
	return &LazyDLL{Name: name, System: true}
}

// LazyProc 实现了对 LazyDLL 内部某个过程的访问。
// 它将查找操作延迟到调用 Addr 方法时进行。
type LazyProc struct {
	Name string

	mu   sync.Mutex
	l    *LazyDLL
	proc *Proc
}

// Find在DLL中搜索名为p.Name的程序。如果搜索失败，它将返回一个错误。如果已找到并加载到内存中的程序，Find将不会进行搜索。

// ff:查找命令
func (p *LazyProc) Find() error {
// 非竞态版本的：
// if p.proc == nil {
	if atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&p.proc))) == nil {
		p.mu.Lock()
		defer p.mu.Unlock()
		if p.proc == nil {
			e := p.l.Load()
			if e != nil {
				return e
			}
			proc, e := p.l.dll.FindProc(p.Name)
			if e != nil {
				return e
			}
// 非竞态版本的：
// p.proc = proc
			atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(&p.proc)), unsafe.Pointer(proc))
		}
	}
	return nil
}

// mustFind 与 Find 类似，但当搜索失败时会触发 panic。
func (p *LazyProc) mustFind() {
	e := p.Find()
	if e != nil {
		panic(e)
	}
}

// Addr 返回由 p 表示的程序的地址。
// 返回值可以传递给 Syscall 以运行该程序。
// It will panic if the procedure cannot be found.
// 翻译提示:func  (p  *懒加载Proc)  地址()  uintptr  {}

// ff:取命令地址
func (p *LazyProc) Addr() uintptr {
	p.mustFind()
	return p.proc.Addr()
}

//go:uintptrescapes

// Call 以参数 a 执行过程 p。如果提供的参数超过 15 个，将会触发恐慌。如果找不到该过程，也会引发恐慌。
//
// 返回的错误始终是非空的，由 GetLastError 的结果构建而成。调用者必须先根据被调用函数的具体语义检查主返回值来判断是否发生错误，之后再查看错误。该错误将确保包含 windows.Errno。

// ff:调用
// a:命令参数s
// r1:
// r2:
// lastErr:最后一个错误
func (p *LazyProc) Call(a ...uintptr) (r1, r2 uintptr, lastErr error) {
	p.mustFind()
	return p.proc.Call(a...)
}

var canDoSearchSystem32Once struct {
	sync.Once
	v bool
}

func initCanDoSearchSystem32() {
// MSDN文档（https://msdn.microsoft.com/en-us/library/ms684179(v=vs.85).aspx）中指出：
// “在Windows 7、Windows Server 2008 R2、Windows Vista以及Windows Server 2008系统上，若已安装KB2533623更新补丁，
// 则可使用LOAD_LIBRARY_SEARCH_*标志。要检测这些标志是否可用，可通过调用GetProcAddress函数获取AddDllDirectory、
// RemoveDllDirectory或SetDefaultDllDirectories函数的地址。若GetProcAddress调用成功，则表明可以将LOAD_LIBRARY_SEARCH_*
// 标志与LoadLibraryEx函数一起使用。”
	canDoSearchSystem32Once.v = (modkernel32.NewProc("AddDllDirectory").Find() == nil)
}

func canDoSearchSystem32() bool {
	canDoSearchSystem32Once.Do(initCanDoSearchSystem32)
	return canDoSearchSystem32Once.v
}

func isBaseName(name string) bool {
	for _, c := range name {
		if c == ':' || c == '/' || c == '\\' {
			return false
		}
	}
	return true
}

// loadLibraryEx 封装了 Windows 的 LoadLibraryEx 函数。
//
// 详情参见：https://msdn.microsoft.com/en-us/library/windows/desktop/ms684179(v=vs.85).aspx
//
// 若 name 不是绝对路径，LoadLibraryEx 会在多种自动搜索位置查找 DLL 文件，
// 除非 flags 参数对此有所限制。具体见：
// https://msdn.microsoft.com/en-us/library/ff919712%28VS.85%29.aspx
func loadLibraryEx(name string, system bool) (*DLL, error) {
	loadDLL := name
	var flags uintptr
	if system {
		if canDoSearchSystem32() {
			flags = LOAD_LIBRARY_SEARCH_SYSTEM32
		} else if isBaseName(name) {
// Windows XP 或未打补丁的 Windows 系统
// 尝试从系统文件夹加载“foo.dll”，但其系统上的 LoadLibraryEx 函数尚不支持该操作，
// 故对其进行模拟实现。
			systemdir, err := GetSystemDirectory()
			if err != nil {
				return nil, err
			}
			loadDLL = systemdir + "\\" + name
		}
	}
	h, err := LoadLibraryEx(loadDLL, 0, flags)
	if err != nil {
		return nil, err
	}
	return &DLL{Name: name, Handle: h}, nil
}

type errString string

// 翻译提示:func  (s  错误字符串)  错误信息()  字符串  {}

// ff:取错误文本
// s:
func (s errString) Error() string { return string(s) }
