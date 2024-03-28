
<原文开始>
// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// 版权所有 ? 2020 Go作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。
# <翻译结束>


<原文开始>
// RegArgs is a struct that has space for each argument
// and return value register on the current architecture.
//
// Assembly code knows the layout of the first two fields
// of RegArgs.
//
// RegArgs also contains additional space to hold pointers
// when it may not be safe to keep them only in the integer
// register space otherwise.
<原文结束>

# <翻译开始>
// RegArgs 是一个结构体，为当前架构下的每个参数和返回值寄存器预留了空间。
//
// 汇编代码了解 RegArgs 的前两个字段的布局。
//
// 当在整数寄存器空间中保存指针可能不安全时，RegArgs 还包含额外的空间来存储指针。
# <翻译结束>


<原文开始>
	// Values in these slots should be precisely the bit-by-bit
	// representation of how they would appear in a register.
	//
	// This means that on big endian arches, integer values should
	// be in the top bits of the slot. Floats are usually just
	// directly represented, but some architectures treat narrow
	// width floating point values specially (e.g. they're promoted
	// first, or they need to be NaN-boxed).
<原文结束>

# <翻译开始>
// 这些槽位中的值应精确地按位表示它们在寄存器中出现的样子。
//
// 这意味着在大端架构（big endian arches）上，整数值应该位于槽位的高比特部分。浮点数通常直接表示，但在某些架构中，窄宽度浮点值会有特殊处理（例如，它们首先被提升精度，或者需要进行NaN包装）。
# <翻译结束>


<原文开始>
// untyped integer registers
<原文结束>

# <翻译开始>
// 无类型整数寄存器
# <翻译结束>


<原文开始>
// untyped float registers
<原文结束>

# <翻译开始>
// 无类型的浮点寄存器
# <翻译结束>


<原文开始>
// Fields above this point are known to assembly.
<原文结束>

# <翻译开始>
// 该处以上字段已在汇编中明确定义。
# <翻译结束>


<原文开始>
	// Ptrs is a space that duplicates Ints but with pointer type,
	// used to make pointers passed or returned  in registers
	// visible to the GC by making the type unsafe.Pointer.
<原文结束>

# <翻译开始>
// Ptrs 是一个与 Ints 空间重复但具有指针类型的区域，
// 通过将其类型设置为 unsafe.Pointer，使得在寄存器中传递或返回的指针对垃圾回收器（GC）可见。
# <翻译结束>


<原文开始>
	// ReturnIsPtr is a bitmap that indicates which registers
	// contain or will contain pointers on the return path from
	// a reflectcall. The i'th bit indicates whether the i'th
	// register contains or will contain a valid Go pointer.
<原文结束>

# <翻译开始>
// ReturnIsPtr 是一个位图，用于指示从 reflectcall 返回路径中哪些寄存器包含或将会包含指针。
// 第i位表示第i个寄存器是否包含或将会包含一个有效的Go指针。
# <翻译结束>


<原文开始>
// IntRegArgAddr returns a pointer inside of r.Ints[reg] that is appropriately
// offset for an argument of size argSize.
//
// argSize must be non-zero, fit in a register, and a power-of-two.
//
// This method is a helper for dealing with the endianness of different CPU
// architectures, since sub-word-sized arguments in big endian architectures
// need to be "aligned" to the upper edge of the register to be interpreted
// by the CPU correctly.
<原文结束>

# <翻译开始>
// IntRegArgAddr 返回一个指向 r.Ints[reg] 内部的指针，该指针根据参数大小 argSize 进行了适当偏移。
//
// argSize 必须为非零值，且需满足能放入寄存器中，并且是2的幂次方。
//
// 此方法是一个辅助工具，用于处理不同CPU架构的字节序问题。在大端字节序架构中，小于字长的参数需要“对齐”到寄存器的高位边缘，以便CPU正确地进行解释。
# <翻译结束>


<原文开始>
// IntArgRegBitmap is a bitmap large enough to hold one bit per
// integer argument/return register.
<原文结束>

# <翻译开始>
// IntArgRegBitmap 是一个足够大的位图，用于为每个整数参数/返回寄存器存储一位。
# <翻译结束>


<原文开始>
// Set sets the i'th bit of the bitmap to 1.
<原文结束>

# <翻译开始>
// Set将位图中第i位置为1。
# <翻译结束>


<原文开始>
// FuncPC* intrinsics.
//
// CAREFUL: In programs with plugins, FuncPC* can return different values
// for the same function (because there are actually multiple copies of
// the same function in the address space). To be safe, don't use the
// results of this function in any == expression. It is only safe to
// use the result as an address at which to start executing code.
<原文结束>

# <翻译开始>
// FuncPC* 内置函数。
//
// 注意：在包含插件的程序中，FuncPC* 可能会对同一函数返回不同的值（因为在地址空间中实际上存在该函数的多个副本）。为了安全起见，不要在任何“==”表达式中使用此函数的结果。唯一安全的做法是将结果作为执行代码的起始地址来使用。
# <翻译结束>


<原文开始>
// FuncPCABI0 returns the entry PC of the function f, which must be a
// direct reference of a function defined as ABI0. Otherwise it is a
// compile-time error.
//
// Implemented as a compile intrinsic.
<原文结束>

# <翻译开始>
// FuncPCABI0 返回函数 f 的入口地址，该函数必须是定义为 ABI0 的直接引用。否则，这将是一个编译时错误。
//
// 以编译器内建方式实现。
# <翻译结束>


<原文开始>
// FuncPCABIInternal returns the entry PC of the function f. If f is a
// direct reference of a function, it must be defined as ABIInternal.
// Otherwise it is a compile-time error. If f is not a direct reference
// of a defined function, it assumes that f is a func value. Otherwise
// the behavior is undefined.
//
// Implemented as a compile intrinsic.
<原文结束>

# <翻译开始>
// FuncPCABIInternal 返回函数f的入口地址。如果f是对函数的直接引用，该函数必须定义为ABIInternal。否则，编译时会出错。若f不是对已定义函数的直接引用，则假定f是一个func值。否则其行为是未定义的。
//
// 作为编译器内建功能实现。
# <翻译结束>

