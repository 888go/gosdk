
<原文开始>
// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// 版权所有 2011 The Go 作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可协议约束，
// 该协议可在 LICENSE 文件中找到。
# <翻译结束>


<原文开始>
// We need to use LoadLibrary and GetProcAddress from the Go runtime, because
// the these symbols are loaded by the system linker and are required to
// dynamically load additional symbols. Note that in the Go runtime, these
// return syscall.Handle and syscall.Errno, but these are the same, in fact,
// as windows.Handle and windows.Errno, and we intend to keep these the same.
<原文结束>

# <翻译开始>
// 我们需要使用Go运行时中的LoadLibrary和GetProcAddress，因为这些符号是由系统链接器加载的，且对于动态加载额外符号是必需的。请注意，在Go运行时中，这些函数返回syscall.Handle和syscall.Errno，但实际上它们与windows.Handle和windows.Errno相同，我们打算保持这些类型一致不变。
# <翻译结束>


<原文开始>
// DLLError describes reasons for DLL load failures.
<原文结束>

# <翻译开始>
// DLLError 用于描述 DLL 加载失败的原因。
# <翻译结束>


<原文开始>
// A DLL implements access to a single DLL.
<原文结束>

# <翻译开始>
// DLL 实现了对单一 DLL 的访问。
# <翻译结束>


<原文开始>
// LoadDLL loads DLL file into memory.
//
// Warning: using LoadDLL without an absolute path name is subject to
// DLL preloading attacks. To safely load a system DLL, use LazyDLL
// with System set to true, or use LoadLibraryEx directly.
<原文结束>

# <翻译开始>
// LoadDLL 将 DLL 文件加载到内存中。
//
// 警告：在没有使用绝对路径的情况下调用 LoadDLL 可能会遭受 DLL 预加载攻击。为了安全地加载系统 DLL，请使用将 System 设置为 true 的 LazyDLL，或者直接使用 LoadLibraryEx。
# <翻译结束>


<原文开始>
// MustLoadDLL is like LoadDLL but panics if load operation failes.
<原文结束>

# <翻译开始>
// MustLoadDLL类似于LoadDLL，但若加载操作失败则会触发panic。
# <翻译结束>


<原文开始>
// FindProc searches DLL d for procedure named name and returns *Proc
// if found. It returns an error if search fails.
<原文结束>

# <翻译开始>
// FindProc在DLL d中搜索名为name的程序，并在找到时返回*Proc。如果搜索失败，它将返回错误。
# <翻译结束>


<原文开始>
// MustFindProc is like FindProc but panics if search fails.
<原文结束>

# <翻译开始>
// MustFindProc 与 FindProc 类似，但在搜索失败时会引发 panic。
# <翻译结束>


<原文开始>
// FindProcByOrdinal searches DLL d for procedure by ordinal and returns *Proc
// if found. It returns an error if search fails.
<原文结束>

# <翻译开始>
// FindProcByOrdinal 在 DLL d 中通过序数查找过程，并在找到时返回 *Proc。如果搜索失败，则返回错误。
# <翻译结束>


<原文开始>
// MustFindProcByOrdinal is like FindProcByOrdinal but panics if search fails.
<原文结束>

# <翻译开始>
// MustFindProcByOrdinal 与 FindProcByOrdinal 类似，但在搜索失败时会引发恐慌。
# <翻译结束>


<原文开始>
// Release unloads DLL d from memory.
<原文结束>

# <翻译开始>
// Release 从内存中卸载 DLL d。
# <翻译结束>


<原文开始>
// A Proc implements access to a procedure inside a DLL.
<原文结束>

# <翻译开始>
// Proc实现了对DLL内部过程的访问
# <翻译结束>


<原文开始>
// Addr returns the address of the procedure represented by p.
// The return value can be passed to Syscall to run the procedure.
<原文结束>

# <翻译开始>
// Addr 返回由 p 表示的程序的地址。
// 返回值可传递给 Syscall 以运行该程序。
# <翻译结束>


<原文开始>
// Call executes procedure p with arguments a. It will panic, if more than 15 arguments
// are supplied.
//
// The returned error is always non-nil, constructed from the result of GetLastError.
// Callers must inspect the primary return value to decide whether an error occurred
// (according to the semantics of the specific function being called) before consulting
// the error. The error will be guaranteed to contain windows.Errno.
<原文结束>

# <翻译开始>
// Call 函数以参数 a 调用过程 p。如果提供的参数超过 15 个，将会引发 panic。
//
// 返回的错误始终是非 nil 值，由 GetLastError 的结果构建而成。调用者必须首先根据被调用特定函数的语义检查主返回值来判断是否发生错误，之后再参考该错误。该错误将确保包含 windows.Errno。
# <翻译结束>


<原文开始>
// A LazyDLL implements access to a single DLL.
// It will delay the load of the DLL until the first
// call to its Handle method or to one of its
// LazyProc's Addr method.
<原文结束>

# <翻译开始>
// LazyDLL 实现了对单一 DLL 的访问。它会将 DLL 的加载延迟到首次调用其 Handle 方法或其某个 LazyProc 的 Addr 方法时。
# <翻译结束>


<原文开始>
	// System determines whether the DLL must be loaded from the
	// Windows System directory, bypassing the normal DLL search
	// path.
<原文结束>

# <翻译开始>
// System 指定 DLL 是否必须从 Windows 系统目录加载，从而绕过常规的 DLL 搜索路径。
# <翻译结束>


<原文开始>
// non nil once DLL is loaded
<原文结束>

# <翻译开始>
// 一旦DLL加载，即非空
# <翻译结束>


<原文开始>
// Load loads DLL file d.Name into memory. It returns an error if fails.
// Load will not try to load DLL, if it is already loaded into memory.
<原文结束>

# <翻译开始>
// Load 将 DLL 文件 d.Name 加载到内存中。加载失败时返回错误。
// 若 DLL 已经被加载到内存中，Load 不会尝试再次加载。
# <翻译结束>


<原文开始>
	// Non-racy version of:
	// if d.dll != nil {
<原文结束>

# <翻译开始>
// 非竞态版本的：
// if d.dll != nil {
# <翻译结束>


<原文开始>
	// kernel32.dll is special, since it's where LoadLibraryEx comes from.
	// The kernel already special-cases its name, so it's always
	// loaded from system32.
<原文结束>

# <翻译开始>
// kernel32.dll 是特殊的，因为它正是 LoadLibraryEx 函数的来源。
// 内核已经对其名称做了特例处理，所以它总是
// 从 system32 目录加载。
# <翻译结束>


<原文开始>
	// Non-racy version of:
	// d.dll = dll
<原文结束>

# <翻译开始>
// 非竞态版本的：
// d.dll = dll
# <翻译结束>


<原文开始>
// mustLoad is like Load but panics if search fails.
<原文结束>

# <翻译开始>
// mustLoad 与 Load 类似，但当搜索失败时会引发 panic。
# <翻译结束>


<原文开始>
// Handle returns d's module handle.
<原文结束>

# <翻译开始>
// Handle 返回 d 的模块句柄。
# <翻译结束>


<原文开始>
// NewProc returns a LazyProc for accessing the named procedure in the DLL d.
<原文结束>

# <翻译开始>
// NewProc 返回一个用于访问 DLL d 中指定名称的程序的 LazyProc。
# <翻译结束>


<原文开始>
// NewLazyDLL creates new LazyDLL associated with DLL file.
<原文结束>

# <翻译开始>
// NewLazyDLL 创建与 DLL 文件关联的新 LazyDLL。
# <翻译结束>


<原文开始>
// NewLazySystemDLL is like NewLazyDLL, but will only
// search Windows System directory for the DLL if name is
// a base name (like "advapi32.dll").
<原文结束>

# <翻译开始>
// NewLazySystemDLL 类似于 NewLazyDLL，但仅当 name 为基名（如 "advapi32.dll"）时，
// 才会在 Windows 系统目录下搜索该 DLL。
# <翻译结束>


<原文开始>
// A LazyProc implements access to a procedure inside a LazyDLL.
// It delays the lookup until the Addr method is called.
<原文结束>

# <翻译开始>
// LazyProc 实现了对 LazyDLL 中某个过程的访问。
// 它将查找操作延迟到调用 Addr 方法时进行。
# <翻译结束>


<原文开始>
// Find searches DLL for procedure named p.Name. It returns
// an error if search fails. Find will not search procedure,
// if it is already found and loaded into memory.
<原文结束>

# <翻译开始>
// Find 在 DLL 中搜索名为 p.Name 的过程。若搜索失败，将返回错误。
// 若该过程已找到并加载到内存中，则 Find 不会进行搜索。
# <翻译结束>


<原文开始>
	// Non-racy version of:
	// if p.proc == nil {
<原文结束>

# <翻译开始>
// 非竞态版本的：
// if p.proc == nil {
# <翻译结束>


<原文开始>
			// Non-racy version of:
			// p.proc = proc
<原文结束>

# <翻译开始>
// 非竞态版本的：
// p.proc = proc
# <翻译结束>


<原文开始>
// mustFind is like Find but panics if search fails.
<原文结束>

# <翻译开始>
// mustFind 与 Find 类似，但当搜索失败时会触发 panic。
# <翻译结束>


<原文开始>
// Addr returns the address of the procedure represented by p.
// The return value can be passed to Syscall to run the procedure.
// It will panic if the procedure cannot be found.
<原文结束>

# <翻译开始>
// Addr 返回由 p 表示的程序的地址。
// 返回值可传递给 Syscall 以运行该程序。
// 若无法找到该程序，将引发恐慌。
# <翻译结束>


<原文开始>
// Call executes procedure p with arguments a. It will panic, if more than 15 arguments
// are supplied. It will also panic if the procedure cannot be found.
//
// The returned error is always non-nil, constructed from the result of GetLastError.
// Callers must inspect the primary return value to decide whether an error occurred
// (according to the semantics of the specific function being called) before consulting
// the error. The error will be guaranteed to contain windows.Errno.
<原文结束>

# <翻译开始>
// Call 以参数 a 执行过程 p。如果提供的参数超过 15 个，将会引发恐慌。如果找不到该过程，也会引发恐慌。
//
// 返回的错误始终是非 nil 值，由 GetLastError 的结果构建而成。调用者必须先根据被调用函数的具体语义检查主要返回值，以判断是否发生错误，之后再查看该错误。该错误将确保包含 windows.Errno。
# <翻译结束>


<原文开始>
	// https://msdn.microsoft.com/en-us/library/ms684179(v=vs.85).aspx says:
	// "Windows 7, Windows Server 2008 R2, Windows Vista, and Windows
	// Server 2008: The LOAD_LIBRARY_SEARCH_* flags are available on
	// systems that have KB2533623 installed. To determine whether the
	// flags are available, use GetProcAddress to get the address of the
	// AddDllDirectory, RemoveDllDirectory, or SetDefaultDllDirectories
	// function. If GetProcAddress succeeds, the LOAD_LIBRARY_SEARCH_*
	// flags can be used with LoadLibraryEx."
<原文结束>

# <翻译开始>
// 根据 https://msdn.microsoft.com/en-us/library/ms684179(v=vs.85).aspx 中所述：
// “Windows 7、Windows Server 2008 R2、Windows Vista 和 Windows Server 2008：LOAD_LIBRARY_SEARCH_* 标志仅在已安装 KB2533623 的系统上可用。要判断这些标志是否可用，可使用 GetProcAddress 函数获取 AddDllDirectory、RemoveDllDirectory 或 SetDefaultDllDirectories 函数的地址。如果 GetProcAddress 成功，即可在 LoadLibraryEx 中使用 LOAD_LIBRARY_SEARCH_* 标志。”
# <翻译结束>


<原文开始>
// loadLibraryEx wraps the Windows LoadLibraryEx function.
//
// See https://msdn.microsoft.com/en-us/library/windows/desktop/ms684179(v=vs.85).aspx
//
// If name is not an absolute path, LoadLibraryEx searches for the DLL
// in a variety of automatic locations unless constrained by flags.
// See: https://msdn.microsoft.com/en-us/library/ff919712%28VS.85%29.aspx
<原文结束>

# <翻译开始>
// loadLibraryEx 封装了 Windows 的 LoadLibraryEx 函数。
//
// 详细信息参见：https://msdn.microsoft.com/en-us/library/windows/desktop/ms684179(v=vs.85).aspx
//
// 若 name 不是绝对路径，除非受限于 flags，否则 LoadLibraryEx 会在多种自动搜索路径中查找 DLL。
// 参见：https://msdn.microsoft.com/en-us/library/ff919712%28VS.85%29.aspx
# <翻译结束>


<原文开始>
			// WindowsXP or unpatched Windows machine
			// trying to load "foo.dll" out of the system
			// folder, but LoadLibraryEx doesn't support
			// that yet on their system, so emulate it.
<原文结束>

# <翻译开始>
// Windows XP 或未经补丁更新的 Windows 系统
// 尝试从系统文件夹加载“foo.dll”
// 但其系统上的 LoadLibraryEx 还不支持该操作
// 因此需要对此进行模拟实现
# <翻译结束>

