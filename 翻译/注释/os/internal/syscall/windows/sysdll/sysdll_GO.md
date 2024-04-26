
<原文开始>
// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
//版权所有2016 The Go Authors。所有权利保留。
//使用此源代码受BSD风格的
//可在LICENSE文件中找到的许可协议管辖。
# <翻译结束>


<原文开始>
// Package sysdll is an internal leaf package that records and reports
// which Windows DLL names are used by Go itself. These DLLs are then
// only loaded from the System32 directory. See Issue 14959.
<原文结束>

# <翻译开始>
// 包sysdll是一个内部的叶子包，用于记录和报告Go自身使用了哪些Windows DLL名称。这些DLL随后仅从System32目录加载。参见问题14959。
# <翻译结束>


<原文开始>
// IsSystemDLL reports whether the named dll key (a base name, like
// "foo.dll") is a system DLL which should only be loaded from the
// Windows SYSTEM32 directory.
//
// Filenames are case sensitive, but that doesn't matter because
// the case registered with Add is also the same case used with
// LoadDLL later.
//
// It has no associated mutex and should only be mutated serially
// (currently: during init), and not concurrent with DLL loading.
<原文结束>

# <翻译开始>
// IsSystemDLL 判断指定的 dll 键（一个基名称，如 "foo.dll"）是否是系统 DLL，它应该只从 Windows 系统目录加载。
// 
// 文件名是区分大小写的，但这没关系，因为使用 Add 注册时的大小写与稍后使用 LoadDLL 时的大小写相同。
// 
// 它没有关联的互斥锁，只能串行修改（目前：仅在初始化期间），并且不应与 DLL 加载操作并发进行。
# <翻译结束>


<原文开始>
// Add notes that dll is a system32 DLL which should only be loaded
// from the Windows SYSTEM32 directory. It returns its argument back,
// for ease of use in generated code.
<原文结束>

# <翻译开始>
// 添加注释，说明dll是一个系统32库，应该仅从Windows SYSTEM32目录加载。它返回其参数，以便于生成代码的使用。
# <翻译结束>

