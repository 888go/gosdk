
<原文开始>
// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
# <翻译结束>


<原文开始>
// package goarch contains GOARCH-specific constants.
<原文结束>

# <翻译开始>
// package goarch contains GOARCH-specific constants.
# <翻译结束>


<原文开始>
// PtrSize is the size of a pointer in bytes - unsafe.Sizeof(uintptr(0)) but as an ideal constant.
// It is also the size of the machine's native word size (that is, 4 on 32-bit systems, 8 on 64-bit).
<原文结束>

# <翻译开始>
// PtrSize is the size of a pointer in bytes - unsafe.Sizeof(uintptr(0)) but as an ideal constant.
// It is also the size of the machine's native word size (that is, 4 on 32-bit systems, 8 on 64-bit).
# <翻译结束>


<原文开始>
// ArchFamily is the architecture family (AMD64, ARM, ...)
<原文结束>

# <翻译开始>
// ArchFamily is the architecture family (AMD64, ARM, ...)
# <翻译结束>


<原文开始>
// BigEndian reports whether the architecture is big-endian.
<原文结束>

# <翻译开始>
// BigEndian reports whether the architecture is big-endian.
# <翻译结束>


<原文开始>
// DefaultPhysPageSize is the default physical page size.
<原文结束>

# <翻译开始>
// DefaultPhysPageSize is the default physical page size.
# <翻译结束>


<原文开始>
// PCQuantum is the minimal unit for a program counter (1 on x86, 4 on most other systems).
// The various PC tables record PC deltas pre-divided by PCQuantum.
<原文结束>

# <翻译开始>
// PCQuantum is the minimal unit for a program counter (1 on x86, 4 on most other systems).
// The various PC tables record PC deltas pre-divided by PCQuantum.
# <翻译结束>


<原文开始>
// Int64Align is the required alignment for a 64-bit integer (4 on 32-bit systems, 8 on 64-bit).
<原文结束>

# <翻译开始>
// Int64Align is the required alignment for a 64-bit integer (4 on 32-bit systems, 8 on 64-bit).
# <翻译结束>


<原文开始>
// MinFrameSize is the size of the system-reserved words at the bottom
// of a frame (just above the architectural stack pointer).
// It is zero on x86 and PtrSize on most non-x86 (LR-based) systems.
// On PowerPC it is larger, to cover three more reserved words:
// the compiler word, the link editor word, and the TOC save word.
<原文结束>

# <翻译开始>
// MinFrameSize is the size of the system-reserved words at the bottom
// of a frame (just above the architectural stack pointer).
// It is zero on x86 and PtrSize on most non-x86 (LR-based) systems.
// On PowerPC it is larger, to cover three more reserved words:
// the compiler word, the link editor word, and the TOC save word.
# <翻译结束>


<原文开始>
// StackAlign is the required alignment of the SP register.
// The stack must be at least word aligned, but some architectures require more.
<原文结束>

# <翻译开始>
// StackAlign is the required alignment of the SP register.
// The stack must be at least word aligned, but some architectures require more.
# <翻译结束>

