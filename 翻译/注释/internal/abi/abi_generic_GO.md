
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
	// ABI-related constants.
	//
	// In the generic case, these are all zero
	// which lets them gracefully degrade to ABI0.
<原文结束>

# <翻译开始>
// ABI 相关常量。
//
// 在一般情况下，这些常量的值均为零，
// 这使得它们能够优雅地降级到 ABI0。
# <翻译结束>


<原文开始>
	// IntArgRegs is the number of registers dedicated
	// to passing integer argument values. Result registers are identical
	// to argument registers, so this number is used for those too.
<原文结束>

# <翻译开始>
// IntArgRegs 是用于传递整数参数值的专用寄存器数量。结果寄存器与参数寄存器相同，因此这个数字也用于表示结果寄存器的数量。
# <翻译结束>


<原文开始>
	// FloatArgRegs is the number of registers dedicated
	// to passing floating-point argument values. Result registers are
	// identical to argument registers, so this number is used for
	// those too.
<原文结束>

# <翻译开始>
// FloatArgRegs 表示专门用于传递浮点数参数值的寄存器数量。由于结果寄存器与参数寄存器相同，所以这个数字同时也表示结果寄存器的数量。
# <翻译结束>


<原文开始>
	// EffectiveFloatRegSize describes the width of floating point
	// registers on the current platform from the ABI's perspective.
	//
	// Since Go only supports 32-bit and 64-bit floating point primitives,
	// this number should be either 0, 4, or 8. 0 indicates no floating
	// point registers for the ABI or that floating point values will be
	// passed via the softfloat ABI.
	//
	// For platforms that support larger floating point register widths,
	// such as x87's 80-bit "registers" (not that we support x87 currently),
	// use 8.
<原文结束>

# <翻译开始>
// EffectiveFloatRegSize 描述了当前平台从ABI（应用程序二进制接口）视角下的浮点寄存器宽度。
//
// 由于Go仅支持32位和64位浮点数原生类型，因此这个数值应为0、4或8。0表示ABI中没有浮点寄存器，或者浮点值将通过softfloat ABI传递。
//
// 对于那些支持更大浮点寄存器宽度的平台，例如x87的80位“寄存器”（尽管目前我们并不支持x87），使用8作为该数值。
# <翻译结束>

