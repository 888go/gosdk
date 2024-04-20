package windows

import "golang.org/x/sys/windows"

// DLLError describes reasons for DLL load failures.
type DLLError struct {
	F windows.DLLError
}

func (e *DLLError) Error() string { //md5:a0c3c5d343d274524c4b59351d5bfed2
	return e.F.Error()
}

func (e *DLLError) Unwrap() error { //md5:499308a8661bb4bb32282dfd01bf2950
	return e.F.Unwrap()
}

// A DLL implements access to a single DLL.
type DLL struct {
	F windows.DLL
}

// LoadDLL 将 DLL 文件加载到内存中。
//
// 警告：在没有使用绝对路径的情况下调用 LoadDLL 可能会遭受 DLL 预加载攻击。为了安全地加载系统 DLL，请使用将 System 设置为 true 的 LazyDLL，或者直接使用 LoadLibraryEx。
func LoadDLL(name string) (dll *DLL, err error) { //md5:02840063d2c22d092d10ce1e9c44a9a0
	d, err := windows.LoadDLL(name)
	return &DLL{
		F: *d,
	}
}

// MustLoadDLL类似于LoadDLL，但若加载操作失败则会触发panic。
func MustLoadDLL(name string) *DLL { //md5:9e9ebb629c6997ead57fb219f428ce77

}

// FindProc在DLL d中搜索名为name的程序，并在找到时返回*Proc。如果搜索失败，它将返回错误。
func (d *DLL) FindProc(name string) (proc *Proc, err error) { //md5:13bf3ccbadb67e7ee4bb60152fa26544

}

// MustFindProc 与 FindProc 类似，但在搜索失败时会引发 panic。
func (d *DLL) MustFindProc(name string) *Proc { //md5:8bfda7eb3655cb44700316039f0c3e8c

}

// FindProcByOrdinal 在 DLL d 中通过序数查找过程，并在找到时返回 *Proc。如果搜索失败，则返回错误。
func (d *DLL) FindProcByOrdinal(ordinal uintptr) (proc *Proc, err error) { //md5:b4d388f9e6dfcfd7858f530c557c110d

}

// MustFindProcByOrdinal 与 FindProcByOrdinal 类似，但在搜索失败时会引发恐慌。
func (d *DLL) MustFindProcByOrdinal(ordinal uintptr) *Proc { //md5:87c82e0cbb531fefb2cfa0ad624f6a26

}

// Release 从内存中卸载 DLL d。
func (d *DLL) Release() (err error) { //md5:c97a424fd57132739dd8d7dec38fbb0c

}

// Addr 返回由 p 表示的程序的地址。
// 返回值可传递给 Syscall 以运行该程序。
func (p *Proc) Addr() uintptr { //md5:c82442aa6d12c576b0c44800736f7cd2

}

// Call 函数以参数 a 调用过程 p。如果提供的参数超过 15 个，将会引发 panic。
//
// 返回的错误始终是非 nil 值，由 GetLastError 的结果构建而成。调用者必须首先根据被调用特定函数的语义检查主返回值来判断是否发生错误，之后再参考该错误。该错误将确保包含 windows.Errno。
func (p *Proc) Call(a ...uintptr) (r1, r2 uintptr, lastErr error) { //md5:76cea572b0d7c7c570c203a529a1e066

}

// Load 将 DLL 文件 d.Name 加载到内存中。加载失败时返回错误。
// 若 DLL 已经被加载到内存中，Load 不会尝试再次加载。
func (d *LazyDLL) Load() error { //md5:d292739f1bd069b91ae1262ded7e256a

}

// Handle 返回 d 的模块句柄。
func (d *LazyDLL) Handle() uintptr { //md5:691b61190db964bc05dd26c6d4964b54

}

// NewProc 返回一个用于访问 DLL d 中指定名称的程序的 LazyProc。
func (d *LazyDLL) NewProc(name string) *LazyProc { //md5:0dd6fe96b1a6de8e28ccf5e4ed0b5673

}

// NewLazyDLL 创建与 DLL 文件关联的新 LazyDLL。
func NewLazyDLL(name string) *LazyDLL { //md5:02d78f635d635ad51bcc7066be088d70

}

// NewLazySystemDLL 类似于 NewLazyDLL，但仅当 name 为基名（如 "advapi32.dll"）时，
// 才会在 Windows 系统目录下搜索该 DLL。
func NewLazySystemDLL(name string) *LazyDLL { //md5:a86cd7930c4308904028b301cf473f58

}

// Find 在 DLL 中搜索名为 p.Name 的过程。若搜索失败，将返回错误。
// 若该过程已找到并加载到内存中，则 Find 不会进行搜索。
func (p *LazyProc) Find() error { //md5:4afb55ed5c7610792ae1c14ac14328c9

}

// Addr 返回由 p 表示的程序的地址。
// 返回值可传递给 Syscall 以运行该程序。
// It will panic if the procedure cannot be found.
func (p *LazyProc) Addr() uintptr { //md5:a2acf6939a88cd417a6c1c3c076b936f

}

// Call 以参数 a 执行过程 p。如果提供的参数超过 15 个，将会引发恐慌。如果找不到该过程，也会引发恐慌。
//
// 返回的错误始终是非 nil 值，由 GetLastError 的结果构建而成。调用者必须先根据被调用函数的具体语义检查主要返回值，以判断是否发生错误，之后再查看该错误。该错误将确保包含 windows.Errno。
func (p *LazyProc) Call(a ...uintptr) (r1, r2 uintptr, lastErr error) { //md5:fecf77616db1b45a85987c1b85b7526d

}
