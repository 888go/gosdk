
<原文开始>
// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
# <翻译结束>


<原文开始>
//sys is replaced by func, but:
<原文结束>

# <翻译开始>
//sys is replaced by func, but:
# <翻译结束>


<原文开始>
//sys LoadLibrary(libname string) (handle uint32, err error) = LoadLibraryA
<原文结束>

# <翻译开始>
//sys LoadLibrary(libname string) (handle uint32, err error) = LoadLibraryA
# <翻译结束>


<原文开始>
//sys LoadLibrary(libname string) (handle uint32, err error) [failretval==-1] = LoadLibraryA
<原文结束>

# <翻译开始>
//sys LoadLibrary(libname string) (handle uint32, err error) [failretval==-1] = LoadLibraryA
# <翻译结束>


<原文开始>
// Param is function parameter
<原文结束>

# <翻译开始>
// Param is function parameter
# <翻译结束>


<原文开始>
// tmpVar returns temp variable name that will be used to represent p during syscall.
<原文结束>

# <翻译开始>
// tmpVar returns temp variable name that will be used to represent p during syscall.
# <翻译结束>


<原文开始>
// BoolTmpVarCode returns source code for bool temp variable.
<原文结束>

# <翻译开始>
// BoolTmpVarCode returns source code for bool temp variable.
# <翻译结束>


<原文开始>
// BoolPointerTmpVarCode returns source code for bool temp variable.
<原文结束>

# <翻译开始>
// BoolPointerTmpVarCode returns source code for bool temp variable.
# <翻译结束>


<原文开始>
// SliceTmpVarCode returns source code for slice temp variable.
<原文结束>

# <翻译开始>
// SliceTmpVarCode returns source code for slice temp variable.
# <翻译结束>


<原文开始>
// StringTmpVarCode returns source code for string temp variable.
<原文结束>

# <翻译开始>
// StringTmpVarCode returns source code for string temp variable.
# <翻译结束>


<原文开始>
// TmpVarCode returns source code for temp variable.
<原文结束>

# <翻译开始>
// TmpVarCode returns source code for temp variable.
# <翻译结束>


<原文开始>
// TmpVarReadbackCode returns source code for reading back the temp variable into the original variable.
<原文结束>

# <翻译开始>
// TmpVarReadbackCode returns source code for reading back the temp variable into the original variable.
# <翻译结束>


<原文开始>
// TmpVarHelperCode returns source code for helper's temp variable.
<原文结束>

# <翻译开始>
// TmpVarHelperCode returns source code for helper's temp variable.
# <翻译结束>


<原文开始>
// SyscallArgList returns source code fragments representing p parameter
// in syscall. Slices are translated into 2 syscall parameters: pointer to
// the first element and length.
<原文结束>

# <翻译开始>
// SyscallArgList returns source code fragments representing p parameter
// in syscall. Slices are translated into 2 syscall parameters: pointer to
// the first element and length.
# <翻译结束>


<原文开始>
// IsError determines if p parameter is used to return error.
<原文结束>

# <翻译开始>
// IsError determines if p parameter is used to return error.
# <翻译结束>


<原文开始>
// HelperType returns type of parameter p used in helper function.
<原文结束>

# <翻译开始>
// HelperType returns type of parameter p used in helper function.
# <翻译结束>


<原文开始>
// join concatenates parameters ps into a string with sep separator.
// Each parameter is converted into string by applying fn to it
// before conversion.
<原文结束>

# <翻译开始>
// join concatenates parameters ps into a string with sep separator.
// Each parameter is converted into string by applying fn to it
// before conversion.
# <翻译结束>


<原文开始>
// Rets describes function return parameters.
<原文结束>

# <翻译开始>
// Rets describes function return parameters.
# <翻译结束>


<原文开始>
// ErrorVarName returns error variable name for r.
<原文结束>

# <翻译开始>
// ErrorVarName returns error variable name for r.
# <翻译结束>


<原文开始>
// ToParams converts r into slice of *Param.
<原文结束>

# <翻译开始>
// ToParams converts r into slice of *Param.
# <翻译结束>


<原文开始>
// List returns source code of syscall return parameters.
<原文结束>

# <翻译开始>
// List returns source code of syscall return parameters.
# <翻译结束>


<原文开始>
// PrintList returns source code of trace printing part correspondent
// to syscall return values.
<原文结束>

# <翻译开始>
// PrintList returns source code of trace printing part correspondent
// to syscall return values.
# <翻译结束>


<原文开始>
// SetReturnValuesCode returns source code that accepts syscall return values.
<原文结束>

# <翻译开始>
// SetReturnValuesCode returns source code that accepts syscall return values.
# <翻译结束>


<原文开始>
// SetErrorCode returns source code that sets return parameters.
<原文结束>

# <翻译开始>
// SetErrorCode returns source code that sets return parameters.
# <翻译结束>


<原文开始>
// Fn describes syscall function.
<原文结束>

# <翻译开始>
// Fn describes syscall function.
# <翻译结束>


<原文开始>
// TODO: get rid of this field and just use parameter index instead
<原文结束>

# <翻译开始>
// TODO: get rid of this field and just use parameter index instead
# <翻译结束>


<原文开始>
// insure tmp variables have uniq names
<原文结束>

# <翻译开始>
// insure tmp variables have uniq names
# <翻译结束>


<原文开始>
// extractParams parses s to extract function parameters.
<原文结束>

# <翻译开始>
// extractParams parses s to extract function parameters.
# <翻译结束>


<原文开始>
// extractSection extracts text out of string s starting after start
// and ending just before end. found return value will indicate success,
// and prefix, body and suffix will contain correspondent parts of string s.
<原文结束>

# <翻译开始>
// extractSection extracts text out of string s starting after start
// and ending just before end. found return value will indicate success,
// and prefix, body and suffix will contain correspondent parts of string s.
# <翻译结束>


<原文开始>
// newFn parses string s and return created function Fn.
<原文结束>

# <翻译开始>
// newFn parses string s and return created function Fn.
# <翻译结束>


<原文开始>
// dll and dll function names
<原文结束>

# <翻译开始>
// dll and dll function names
# <翻译结束>


<原文开始>
// DLLName returns DLL name for function f.
<原文结束>

# <翻译开始>
// DLLName returns DLL name for function f.
# <翻译结束>


<原文开始>
// DLLVar returns a valid Go identifier that represents DLLName.
<原文结束>

# <翻译开始>
// DLLVar returns a valid Go identifier that represents DLLName.
# <翻译结束>


<原文开始>
// DLLFuncName returns DLL function name for function f.
<原文结束>

# <翻译开始>
// DLLFuncName returns DLL function name for function f.
# <翻译结束>


<原文开始>
// ParamList returns source code for function f parameters.
<原文结束>

# <翻译开始>
// ParamList returns source code for function f parameters.
# <翻译结束>


<原文开始>
// HelperParamList returns source code for helper function f parameters.
<原文结束>

# <翻译开始>
// HelperParamList returns source code for helper function f parameters.
# <翻译结束>


<原文开始>
// ParamPrintList returns source code of trace printing part correspondent
// to syscall input parameters.
<原文结束>

# <翻译开始>
// ParamPrintList returns source code of trace printing part correspondent
// to syscall input parameters.
# <翻译结束>


<原文开始>
// ParamCount return number of syscall parameters for function f.
<原文结束>

# <翻译开始>
// ParamCount return number of syscall parameters for function f.
# <翻译结束>


<原文开始>
// SyscallParamCount determines which version of Syscall/Syscall6/Syscall9/...
// to use. It returns parameter count for correspondent SyscallX function.
<原文结束>

# <翻译开始>
// SyscallParamCount determines which version of Syscall/Syscall6/Syscall9/...
// to use. It returns parameter count for correspondent SyscallX function.
# <翻译结束>


<原文开始>
// Syscall determines which SyscallX function to use for function f.
<原文结束>

# <翻译开始>
// Syscall determines which SyscallX function to use for function f.
# <翻译结束>


<原文开始>
// SyscallParamList returns source code for SyscallX parameters for function f.
<原文结束>

# <翻译开始>
// SyscallParamList returns source code for SyscallX parameters for function f.
# <翻译结束>


<原文开始>
// HelperCallParamList returns source code of call into function f helper.
<原文结束>

# <翻译开始>
// HelperCallParamList returns source code of call into function f helper.
# <翻译结束>


<原文开始>
// MaybeAbsent returns source code for handling functions that are possibly unavailable.
<原文结束>

# <翻译开始>
// MaybeAbsent returns source code for handling functions that are possibly unavailable.
# <翻译结束>


<原文开始>
// IsUTF16 is true, if f is W (utf16) function. It is false
// for all A (ascii) functions.
<原文结束>

# <翻译开始>
// IsUTF16 is true, if f is W (utf16) function. It is false
// for all A (ascii) functions.
# <翻译结束>


<原文开始>
// StrconvFunc returns name of Go string to OS string function for f.
<原文结束>

# <翻译开始>
// StrconvFunc returns name of Go string to OS string function for f.
# <翻译结束>


<原文开始>
// StrconvType returns Go type name used for OS string for f.
<原文结束>

# <翻译开始>
// StrconvType returns Go type name used for OS string for f.
# <翻译结束>


<原文开始>
// HasStringParam is true, if f has at least one string parameter.
// Otherwise it is false.
<原文结束>

# <翻译开始>
// HasStringParam is true, if f has at least one string parameter.
// Otherwise it is false.
# <翻译结束>


<原文开始>
// HelperName returns name of function f helper.
<原文结束>

# <翻译开始>
// HelperName returns name of function f helper.
# <翻译结束>


<原文开始>
// DLL is a DLL's filename and a string that is valid in a Go identifier that should be used when
// naming a variable that refers to the DLL.
<原文结束>

# <翻译开始>
// DLL is a DLL's filename and a string that is valid in a Go identifier that should be used when
// naming a variable that refers to the DLL.
# <翻译结束>


<原文开始>
// Source files and functions.
<原文结束>

# <翻译开始>
// Source files and functions.
# <翻译结束>


<原文开始>
// ParseFiles parses files listed in fs and extracts all syscall
// functions listed in sys comments. It returns source files
// and functions collection *Source if successful.
<原文结束>

# <翻译开始>
// ParseFiles parses files listed in fs and extracts all syscall
// functions listed in sys comments. It returns source files
// and functions collection *Source if successful.
# <翻译结束>


<原文开始>
// DLLs return dll names for a source set src.
<原文结束>

# <翻译开始>
// DLLs return dll names for a source set src.
# <翻译结束>


<原文开始>
// ParseFile adds additional file path to a source set src.
<原文结束>

# <翻译开始>
// ParseFile adds additional file path to a source set src.
# <翻译结束>


<原文开始>
// IsStdRepo reports whether src is part of standard library.
<原文结束>

# <翻译开始>
// IsStdRepo reports whether src is part of standard library.
# <翻译结束>


<原文开始>
// Generate output source file from a source set src.
<原文结束>

# <翻译开始>
// Generate output source file from a source set src.
# <翻译结束>


<原文开始>
// any package in std library
<原文结束>

# <翻译开始>
// any package in std library
# <翻译结束>


<原文开始>
// TODO: this needs better logic than just using package name
<原文结束>

# <翻译开始>
// TODO: this needs better logic than just using package name
# <翻译结束>


<原文开始>
// TODO: use println instead to print in the following template
<原文结束>

# <翻译开始>
// TODO: use println instead to print in the following template
# <翻译结束>


<原文开始>
// Code generated by 'go generate'; DO NOT EDIT.
<原文结束>

# <翻译开始>
// Code generated by 'go generate'; DO NOT EDIT.
# <翻译结束>


<原文开始>
// Do the interface allocations only once for common
// Errno values.
<原文结束>

# <翻译开始>
// Do the interface allocations only once for common
// Errno values.
# <翻译结束>


<原文开始>
// errnoErr returns common boxed Errno values, to prevent
// allocations at runtime.
<原文结束>

# <翻译开始>
// errnoErr returns common boxed Errno values, to prevent
// allocations at runtime.
# <翻译结束>


<原文开始>
	// TODO: add more here, after collecting data on the common
	// error values see on Windows. (perhaps when running
	// all.bat?)
<原文结束>

# <翻译开始>
	// TODO: add more here, after collecting data on the common
	// error values see on Windows. (perhaps when running
	// all.bat?)
# <翻译结束>

