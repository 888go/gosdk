// 版权所有 2011 The Go 作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可协议约束，
// 该协议可在 LICENSE 文件中找到。

package windows

import (
	"sync"
	"sync/atomic"
	"syscall"
	"unsafe"
)

// 我们需要使用Go运行时中的LoadLibrary和GetProcAddress，因为这些符号是由系统链接器加载的，且对于动态加载额外符号是必需的。请注意，在Go运行时中，这些函数返回syscall.Handle和syscall.Errno，但实际上它们与windows.Handle和windows.Errno相同，我们打算保持这些类型一致不变。

//go:linkname syscall_loadlibrary syscall.loadlibrary
func syscall_loadlibrary(filename *uint16) (handle Handle, err Errno)

//go:linkname syscall_getprocaddress syscall.getprocaddress
func syscall_getprocaddress(handle Handle, procname *uint8) (proc uintptr, err Errno)

// DLLError 用于描述 DLL 加载失败的原因。
type DLLError struct {
	Err     error
	ObjName string
	Msg     string
}

func (e *DLLError) Error() string { return e.Msg }

func (e *DLLError) Unwrap() error { return e.Err }

// DLL 实现了对单一 DLL 的访问。
type DLL struct {
	Name   string
	Handle Handle
}

// LoadDLL 将 DLL 文件加载到内存中。
//
// 警告：在没有使用绝对路径的情况下调用 LoadDLL 可能会遭受 DLL 预加载攻击。为了安全地加载系统 DLL，请使用将 System 设置为 true 的 LazyDLL，或者直接使用 LoadLibraryEx。
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

// MustLoadDLL类似于LoadDLL，但若加载操作失败则会触发panic。
func MustLoadDLL(name string) *DLL {
	d, e := LoadDLL(name)
	if e != nil {
		panic(e)
	}
	return d
}

// FindProc在DLL d中搜索名为name的程序，并在找到时返回*Proc。如果搜索失败，它将返回错误。
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
func (d *DLL) MustFindProc(name string) *Proc {
	p, e := d.FindProc(name)
	if e != nil {
		panic(e)
	}
	return p
}

// FindProcByOrdinal 在 DLL d 中通过序数查找过程，并在找到时返回 *Proc。如果搜索失败，则返回错误。
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

// MustFindProcByOrdinal 与 FindProcByOrdinal 类似，但在搜索失败时会引发恐慌。
func (d *DLL) MustFindProcByOrdinal(ordinal uintptr) *Proc {
	p, e := d.FindProcByOrdinal(ordinal)
	if e != nil {
		panic(e)
	}
	return p
}

// Release 从内存中卸载 DLL d。
func (d *DLL) Release() (err error) {
	return FreeLibrary(d.Handle)
}

// Proc实现了对DLL内部过程的访问
type Proc struct {
	Dll  *DLL
	Name string
	addr uintptr
}

// Addr 返回由 p 表示的程序的地址。
// 返回值可传递给 Syscall 以运行该程序。
func (p *Proc) Addr() uintptr {
	return p.addr
}

//go:uintptrescapes

// Call 函数以参数 a 调用过程 p。如果提供的参数超过 15 个，将会引发 panic。
//
// 返回的错误始终是非 nil 值，由 GetLastError 的结果构建而成。调用者必须首先根据被调用特定函数的语义检查主返回值来判断是否发生错误，之后再参考该错误。该错误将确保包含 windows.Errno。
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

// LazyDLL 实现了对单一 DLL 的访问。它会将 DLL 的加载延迟到首次调用其 Handle 方法或其某个 LazyProc 的 Addr 方法时。
type LazyDLL struct {
	Name string

// System 指定 DLL 是否必须从 Windows 系统目录加载，从而绕过常规的 DLL 搜索路径。
	System bool

	mu  sync.Mutex
	dll *DLL // 一旦DLL加载，即非空
}

// Load 将 DLL 文件 d.Name 加载到内存中。加载失败时返回错误。
// 若 DLL 已经被加载到内存中，Load 不会尝试再次加载。
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
// 内核已经对其名称做了特例处理，所以它总是
// 从 system32 目录加载。
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

// mustLoad 与 Load 类似，但当搜索失败时会引发 panic。
func (d *LazyDLL) mustLoad() {
	e := d.Load()
	if e != nil {
		panic(e)
	}
}

// Handle 返回 d 的模块句柄。
func (d *LazyDLL) Handle() uintptr {
	d.mustLoad()
	return uintptr(d.dll.Handle)
}

// NewProc 返回一个用于访问 DLL d 中指定名称的程序的 LazyProc。
func (d *LazyDLL) NewProc(name string) *LazyProc {
	return &LazyProc{l: d, Name: name}
}

// NewLazyDLL 创建与 DLL 文件关联的新 LazyDLL。
func NewLazyDLL(name string) *LazyDLL {
	return &LazyDLL{Name: name}
}

// NewLazySystemDLL 类似于 NewLazyDLL，但仅当 name 为基名（如 "advapi32.dll"）时，
// 才会在 Windows 系统目录下搜索该 DLL。
func NewLazySystemDLL(name string) *LazyDLL {
	return &LazyDLL{Name: name, System: true}
}

// LazyProc 实现了对 LazyDLL 中某个过程的访问。
// 它将查找操作延迟到调用 Addr 方法时进行。
type LazyProc struct {
	Name string

	mu   sync.Mutex
	l    *LazyDLL
	proc *Proc
}

// Find 在 DLL 中搜索名为 p.Name 的过程。若搜索失败，将返回错误。
// 若该过程已找到并加载到内存中，则 Find 不会进行搜索。
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
// 返回值可传递给 Syscall 以运行该程序。
// It will panic if the procedure cannot be found.
func (p *LazyProc) Addr() uintptr {
	p.mustFind()
	return p.proc.Addr()
}

//go:uintptrescapes

// Call 以参数 a 执行过程 p。如果提供的参数超过 15 个，将会引发恐慌。如果找不到该过程，也会引发恐慌。
//
// 返回的错误始终是非 nil 值，由 GetLastError 的结果构建而成。调用者必须先根据被调用函数的具体语义检查主要返回值，以判断是否发生错误，之后再查看该错误。该错误将确保包含 windows.Errno。
func (p *LazyProc) Call(a ...uintptr) (r1, r2 uintptr, lastErr error) {
	p.mustFind()
	return p.proc.Call(a...)
}

var canDoSearchSystem32Once struct {
	sync.Once
	v bool
}

func initCanDoSearchSystem32() {
// 根据 https://msdn.microsoft.com/en-us/library/ms684179(v=vs.85).aspx 中所述：
// “Windows 7、Windows Server 2008 R2、Windows Vista 和 Windows Server 2008：LOAD_LIBRARY_SEARCH_* 标志仅在已安装 KB2533623 的系统上可用。要判断这些标志是否可用，可使用 GetProcAddress 函数获取 AddDllDirectory、RemoveDllDirectory 或 SetDefaultDllDirectories 函数的地址。如果 GetProcAddress 成功，即可在 LoadLibraryEx 中使用 LOAD_LIBRARY_SEARCH_* 标志。”
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
// 详细信息参见：https://msdn.microsoft.com/en-us/library/windows/desktop/ms684179(v=vs.85).aspx
//
// 若 name 不是绝对路径，除非受限于 flags，否则 LoadLibraryEx 会在多种自动搜索路径中查找 DLL。
// 参见：https://msdn.microsoft.com/en-us/library/ff919712%28VS.85%29.aspx
func loadLibraryEx(name string, system bool) (*DLL, error) {
	loadDLL := name
	var flags uintptr
	if system {
		if canDoSearchSystem32() {
			flags = LOAD_LIBRARY_SEARCH_SYSTEM32
		} else if isBaseName(name) {
// Windows XP 或未经补丁更新的 Windows 系统
// 尝试从系统文件夹加载“foo.dll”
// 但其系统上的 LoadLibraryEx 还不支持该操作
// 因此需要对此进行模拟实现
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

func (s errString) Error() string { return string(s) }
