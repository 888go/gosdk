
<原文开始>
// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// 版权所有 ? 2013 The Go 作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可协议约束，
// 该协议可在 LICENSE 文件中找到。
# <翻译结束>


<原文开始>
// Param is function parameter
<原文结束>

# <翻译开始>
// Param 是函数参数
# <翻译结束>


<原文开始>
// tmpVar returns temp variable name that will be used to represent p during syscall.
<原文结束>

# <翻译开始>
// tmpVar 返回一个临时变量名，该变量将在系统调用期间用于表示 p。
# <翻译结束>


<原文开始>
// BoolTmpVarCode returns source code for bool temp variable.
<原文结束>

# <翻译开始>
// BoolTmpVarCode 返回用于布尔临时变量的源代码
# <翻译结束>


<原文开始>
// BoolPointerTmpVarCode returns source code for bool temp variable.
<原文结束>

# <翻译开始>
// BoolPointerTmpVarCode 返回用于布尔型临时变量的源代码
# <翻译结束>


<原文开始>
// SliceTmpVarCode returns source code for slice temp variable.
<原文结束>

# <翻译开始>
// SliceTmpVarCode 返回切片临时变量的源代码
# <翻译结束>


<原文开始>
// StringTmpVarCode returns source code for string temp variable.
<原文结束>

# <翻译开始>
// StringTmpVarCode 返回字符串临时变量的源代码
# <翻译结束>


<原文开始>
// TmpVarCode returns source code for temp variable.
<原文结束>

# <翻译开始>
// TmpVarCode 返回临时变量的源代码
# <翻译结束>


<原文开始>
// TmpVarReadbackCode returns source code for reading back the temp variable into the original variable.
<原文结束>

# <翻译开始>
// TmpVarReadbackCode 返回用于将临时变量回读到原变量的源代码
# <翻译结束>


<原文开始>
// TmpVarHelperCode returns source code for helper's temp variable.
<原文结束>

# <翻译开始>
// TmpVarHelperCode 返回助手临时变量的源代码
# <翻译结束>


<原文开始>
// SyscallArgList returns source code fragments representing p parameter
// in syscall. Slices are translated into 2 syscall parameters: pointer to
// the first element and length.
<原文结束>

# <翻译开始>
// SyscallArgList 返回表示在系统调用中p参数的源代码片段。
// 切片被转化为两个系统调用参数：指向第一个元素的指针和长度。
# <翻译结束>


<原文开始>
// IsError determines if p parameter is used to return error.
<原文结束>

# <翻译开始>
// IsError 判断参数p是否用于返回错误
# <翻译结束>


<原文开始>
// HelperType returns type of parameter p used in helper function.
<原文结束>

# <翻译开始>
// HelperType 返回在辅助函数中使用的参数 p 的类型
# <翻译结束>


<原文开始>
// join concatenates parameters ps into a string with sep separator.
// Each parameter is converted into string by applying fn to it
// before conversion.
<原文结束>

# <翻译开始>
// join 函数通过 sep 分隔符将参数 ps 连接成一个字符串。
// 在进行拼接前，每个参数都会通过应用 fn 函数转换为字符串形式。
# <翻译结束>


<原文开始>
// Rets describes function return parameters.
<原文结束>

# <翻译开始>
// Rets 描述函数返回参数。
# <翻译结束>


<原文开始>
// ErrorVarName returns error variable name for r.
<原文结束>

# <翻译开始>
// ErrorVarName 返回错误变量名给r
# <翻译结束>


<原文开始>
// ToParams converts r into slice of *Param.
<原文结束>

# <翻译开始>
// ToParams 将 r 转换为 []*Param 类型的切片。
# <翻译结束>


<原文开始>
// List returns source code of syscall return parameters.
<原文结束>

# <翻译开始>
// List 返回 syscall 返回参数的源代码
# <翻译结束>


<原文开始>
// PrintList returns source code of trace printing part correspondent
// to syscall return values.
<原文结束>

# <翻译开始>
// PrintList 返回与系统调用返回值相对应的跟踪打印部分的源代码
# <翻译结束>


<原文开始>
// SetReturnValuesCode returns source code that accepts syscall return values.
<原文结束>

# <翻译开始>
// SetReturnValuesCode 返回接受系统调用返回值的源代码
# <翻译结束>


<原文开始>
// SetErrorCode returns source code that sets return parameters.
<原文结束>

# <翻译开始>
// SetErrorCode 返回设置返回参数的源代码
# <翻译结束>


<原文开始>
// Fn describes syscall function.
<原文结束>

# <翻译开始>
// Fn 描述系统调用函数。
# <翻译结束>


<原文开始>
// TODO: get rid of this field and just use parameter index instead
<原文结束>

# <翻译开始>
// TODO：移除该字段，直接使用参数索引代替
# <翻译结束>


<原文开始>
// insure tmp variables have uniq names
<原文结束>

# <翻译开始>
// 确保临时变量具有唯一名称
# <翻译结束>


<原文开始>
// extractParams parses s to extract function parameters.
<原文结束>

# <翻译开始>
// extractParams 用于解析 s 以提取函数参数。
# <翻译结束>


<原文开始>
// extractSection extracts text out of string s starting after start
// and ending just before end. found return value will indicate success,
// and prefix, body and suffix will contain correspondent parts of string s.
<原文结束>

# <翻译开始>
// extractSection 从字符串 s 中提取始于 start 之后、止于 end 之前的文本。返回值 found 用于指示操作是否成功，prefix、body 和 suffix 将分别包含字符串 s 中对应的各部分。
# <翻译结束>


<原文开始>
// newFn parses string s and return created function Fn.
<原文结束>

# <翻译开始>
// newFn解析字符串s并返回创建的函数Fn
# <翻译结束>


<原文开始>
// dll and dll function names
<原文结束>

# <翻译开始>
// dll 和 dll 函数名称
# <翻译结束>


<原文开始>
// DLLName returns DLL name for function f.
<原文结束>

# <翻译开始>
// DLLName 返回函数f对应的DLL名称。
# <翻译结束>


<原文开始>
// DLLVar returns a valid Go identifier that represents DLLName.
<原文结束>

# <翻译开始>
// DLLVar 返回一个表示 DLLName 的有效 Go 标识符。
# <翻译结束>


<原文开始>
// DLLFuncName returns DLL function name for function f.
<原文结束>

# <翻译开始>
// DLLFuncName 返回函数f在DLL中的函数名
# <翻译结束>


<原文开始>
// ParamList returns source code for function f parameters.
<原文结束>

# <翻译开始>
// ParamList 返回函数 f 参数的源代码
# <翻译结束>


<原文开始>
// HelperParamList returns source code for helper function f parameters.
<原文结束>

# <翻译开始>
// HelperParamList 返回用于辅助函数f的参数源代码
# <翻译结束>


<原文开始>
// ParamPrintList returns source code of trace printing part correspondent
// to syscall input parameters.
<原文结束>

# <翻译开始>
// ParamPrintList 返回与系统调用输入参数对应的跟踪打印部分的源代码
# <翻译结束>


<原文开始>
// ParamCount return number of syscall parameters for function f.
<原文结束>

# <翻译开始>
// ParamCount 返回函数f的系统调用参数数量
# <翻译结束>


<原文开始>
// SyscallParamCount determines which version of Syscall/Syscall6/Syscall9/...
// to use. It returns parameter count for correspondent SyscallX function.
<原文结束>

# <翻译开始>
// SyscallParamCount 用于确定使用哪个版本的 Syscall、Syscall6、Syscall9 等函数。
// 它返回相应 SyscallX 函数的参数个数。
# <翻译结束>


<原文开始>
// Syscall determines which SyscallX function to use for function f.
<原文结束>

# <翻译开始>
// Syscall 确定应使用哪个 SyscallX 函数来处理函数 f。
# <翻译结束>


<原文开始>
// SyscallParamList returns source code for SyscallX parameters for function f.
<原文结束>

# <翻译开始>
// SyscallParamList 为函数 f 返回用于 SyscallX 的参数源代码
# <翻译结束>


<原文开始>
// HelperCallParamList returns source code of call into function f helper.
<原文结束>

# <翻译开始>
// HelperCallParamList 返回调用函数f辅助工具的源代码
# <翻译结束>


<原文开始>
// MaybeAbsent returns source code for handling functions that are possibly unavailable.
<原文结束>

# <翻译开始>
// MaybeAbsent 返回用于处理可能不可用的函数的源代码
# <翻译结束>


<原文开始>
// IsUTF16 is true, if f is W (utf16) function. It is false
// for all A (ascii) functions.
<原文结束>

# <翻译开始>
// IsUTF16 为真，当 f 为 W（utf16）函数时。对于所有 A（ascii）函数，其为假。
# <翻译结束>


<原文开始>
// StrconvFunc returns name of Go string to OS string function for f.
<原文结束>

# <翻译开始>
// StrconvFunc 为 f 返回 Go 字符串转操作系统字符串函数的名称。
# <翻译结束>


<原文开始>
// StrconvType returns Go type name used for OS string for f.
<原文结束>

# <翻译开始>
// StrconvType 返回用于操作系统字符串表示f的Go类型名
# <翻译结束>


<原文开始>
// HasStringParam is true, if f has at least one string parameter.
// Otherwise it is false.
<原文结束>

# <翻译开始>
// HasStringParam为true表示f至少有一个字符串参数。
// 否则，其为false。
# <翻译结束>


<原文开始>
// HelperName returns name of function f helper.
<原文结束>

# <翻译开始>
// HelperName 返回函数f的辅助函数名称
# <翻译结束>


<原文开始>
// DLL is a DLL's filename and a string that is valid in a Go identifier that should be used when
// naming a variable that refers to the DLL.
<原文结束>

# <翻译开始>
// DLL 是一个 DLL 文件名以及一个在 Go 标识符中有效的字符串，该字符串应在命名引用此 DLL 的变量时使用。
# <翻译结束>


<原文开始>
// Source files and functions.
<原文结束>

# <翻译开始>
// 源文件和函数
# <翻译结束>


<原文开始>
// ParseFiles parses files listed in fs and extracts all syscall
// functions listed in sys comments. It returns source files
// and functions collection *Source if successful.
<原文结束>

# <翻译开始>
// ParseFiles解析fs中列出的文件，并从sys注释中提取所有syscall函数。若成功，返回包含源文件和函数集合的*Source。
# <翻译结束>


<原文开始>
// DLLs return dll names for a source set src.
<原文结束>

# <翻译开始>
// DLLs 返回源集合src对应的dll名称。
# <翻译结束>


<原文开始>
// ParseFile adds additional file path to a source set src.
<原文结束>

# <翻译开始>
// ParseFile 向源码集合 src 中添加额外的文件路径。
# <翻译结束>


<原文开始>
// IsStdRepo reports whether src is part of standard library.
<原文结束>

# <翻译开始>
// IsStdRepo 判断 src 是否属于标准库。
# <翻译结束>


<原文开始>
// Generate output source file from a source set src.
<原文结束>

# <翻译开始>
// 从源集合src生成输出源文件
# <翻译结束>


<原文开始>
// any package in std library
<原文结束>

# <翻译开始>
// 标准库中的任何包
# <翻译结束>


<原文开始>
// TODO: this needs better logic than just using package name
<原文结束>

# <翻译开始>
// TODO：此处需要比单纯使用包名更合理的逻辑
# <翻译结束>


<原文开始>
// TODO: use println instead to print in the following template
<原文结束>

# <翻译开始>
// 待办：在以下模板中使用println进行打印
# <翻译结束>


<原文开始>
// Code generated by 'go generate'; DO NOT EDIT.
<原文结束>

# <翻译开始>
// 代码由'go generate'命令生成；请勿编辑。
# <翻译结束>


<原文开始>
// Do the interface allocations only once for common
// Errno values.
<原文结束>

# <翻译开始>
// 对于常见的Errno值，仅一次性完成接口分配。
# <翻译结束>


<原文开始>
// errnoErr returns common boxed Errno values, to prevent
// allocations at runtime.
<原文结束>

# <翻译开始>
// errnoErr 函数返回常见的已封装 Errno 值，旨在避免运行时的内存分配。
# <翻译结束>


<原文开始>
	// TODO: add more here, after collecting data on the common
	// error values see on Windows. (perhaps when running
	// all.bat?)
<原文结束>

# <翻译开始>
	// TODO: 在收集到Windows上常见的错误值数据后，此处应添加更多内容。（可能在运行all.bat时进行？）
# <翻译结束>

