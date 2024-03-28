
<原文开始>
// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
# <翻译结束>


<原文开始>
// Package cpu implements processor feature detection
// used by the Go standard library.
<原文结束>

# <翻译开始>
// Package cpu implements processor feature detection
// used by the Go standard library.
# <翻译结束>


<原文开始>
// DebugOptions is set to true by the runtime if the OS supports reading
// GODEBUG early in runtime startup.
// This should not be changed after it is initialized.
<原文结束>

# <翻译开始>
// DebugOptions is set to true by the runtime if the OS supports reading
// GODEBUG early in runtime startup.
// This should not be changed after it is initialized.
# <翻译结束>


<原文开始>
// CacheLinePad is used to pad structs to avoid false sharing.
<原文结束>

# <翻译开始>
// CacheLinePad is used to pad structs to avoid false sharing.
# <翻译结束>


<原文开始>
// CacheLineSize is the CPU's assumed cache line size.
// There is currently no runtime detection of the real cache line size
// so we use the constant per GOARCH CacheLinePadSize as an approximation.
<原文结束>

# <翻译开始>
// CacheLineSize is the CPU's assumed cache line size.
// There is currently no runtime detection of the real cache line size
// so we use the constant per GOARCH CacheLinePadSize as an approximation.
# <翻译结束>


<原文开始>
// The booleans in X86 contain the correspondingly named cpuid feature bit.
// HasAVX and HasAVX2 are only set if the OS does support XMM and YMM registers
// in addition to the cpuid feature bit being set.
// The struct is padded to avoid false sharing.
<原文结束>

# <翻译开始>
// The booleans in X86 contain the correspondingly named cpuid feature bit.
// HasAVX and HasAVX2 are only set if the OS does support XMM and YMM registers
// in addition to the cpuid feature bit being set.
// The struct is padded to avoid false sharing.
# <翻译结束>


<原文开始>
// The booleans in ARM contain the correspondingly named cpu feature bit.
// The struct is padded to avoid false sharing.
<原文结束>

# <翻译开始>
// The booleans in ARM contain the correspondingly named cpu feature bit.
// The struct is padded to avoid false sharing.
# <翻译结束>


<原文开始>
// The booleans in ARM64 contain the correspondingly named cpu feature bit.
// The struct is padded to avoid false sharing.
<原文结束>

# <翻译开始>
// The booleans in ARM64 contain the correspondingly named cpu feature bit.
// The struct is padded to avoid false sharing.
# <翻译结束>


<原文开始>
// For ppc64(le), it is safe to check only for ISA level starting on ISA v3.00,
// since there are no optional categories. There are some exceptions that also
// require kernel support to work (darn, scv), so there are feature bits for
// those as well. The minimum processor requirement is POWER8 (ISA 2.07).
// The struct is padded to avoid false sharing.
<原文结束>

# <翻译开始>
// For ppc64(le), it is safe to check only for ISA level starting on ISA v3.00,
// since there are no optional categories. There are some exceptions that also
// require kernel support to work (darn, scv), so there are feature bits for
// those as well. The minimum processor requirement is POWER8 (ISA 2.07).
// The struct is padded to avoid false sharing.
# <翻译结束>


<原文开始>
// Hardware random number generator (requires kernel enablement)
<原文结束>

# <翻译开始>
// Hardware random number generator (requires kernel enablement)
# <翻译结束>


<原文开始>
// Syscall vectored (requires kernel enablement)
<原文结束>

# <翻译开始>
// Syscall vectored (requires kernel enablement)
# <翻译结束>


<原文开始>
// z architecture mode is active [mandatory]
<原文结束>

# <翻译开始>
// z architecture mode is active [mandatory]
# <翻译结束>


<原文开始>
// store facility list extended [mandatory]
<原文结束>

# <翻译开始>
// store facility list extended [mandatory]
# <翻译结束>


<原文开始>
// long (20-bit) displacements [mandatory]
<原文结束>

# <翻译开始>
// long (20-bit) displacements [mandatory]
# <翻译结束>


<原文开始>
// 32-bit immediates [mandatory]
<原文结束>

# <翻译开始>
// 32-bit immediates [mandatory]
# <翻译结束>


<原文开始>
// message security assist (CPACF)
<原文结束>

# <翻译开始>
// message security assist (CPACF)
# <翻译结束>


<原文开始>
// KM-AES{128,192,256} functions
<原文结束>

# <翻译开始>
// KM-AES{128,192,256} functions
# <翻译结束>


<原文开始>
// KMC-AES{128,192,256} functions
<原文结束>

# <翻译开始>
// KMC-AES{128,192,256} functions
# <翻译结束>


<原文开始>
// KMCTR-AES{128,192,256} functions
<原文结束>

# <翻译开始>
// KMCTR-AES{128,192,256} functions
# <翻译结束>


<原文开始>
// KMA-GCM-AES{128,192,256} functions
<原文结束>

# <翻译开始>
// KMA-GCM-AES{128,192,256} functions
# <翻译结束>


<原文开始>
// K{I,L}MD-SHA-1 functions
<原文结束>

# <翻译开始>
// K{I,L}MD-SHA-1 functions
# <翻译结束>


<原文开始>
// K{I,L}MD-SHA-256 functions
<原文结束>

# <翻译开始>
// K{I,L}MD-SHA-256 functions
# <翻译结束>


<原文开始>
// K{I,L}MD-SHA-512 functions
<原文结束>

# <翻译开始>
// K{I,L}MD-SHA-512 functions
# <翻译结束>


<原文开始>
// K{I,L}MD-SHA3-{224,256,384,512} and K{I,L}MD-SHAKE-{128,256} functions
<原文结束>

# <翻译开始>
// K{I,L}MD-SHA3-{224,256,384,512} and K{I,L}MD-SHAKE-{128,256} functions
# <翻译结束>


<原文开始>
// vector facility. Note: the runtime sets this when it processes auxv records.
<原文结束>

# <翻译开始>
// vector facility. Note: the runtime sets this when it processes auxv records.
# <翻译结束>


<原文开始>
// vector-enhancements facility 1
<原文结束>

# <翻译开始>
// vector-enhancements facility 1
# <翻译结束>


<原文开始>
// elliptic curve functions
<原文结束>

# <翻译开始>
// elliptic curve functions
# <翻译结束>


<原文开始>
// Initialize examines the processor and sets the relevant variables above.
// This is called by the runtime package early in program initialization,
// before normal init functions are run. env is set by runtime if the OS supports
// cpu feature options in GODEBUG.
<原文结束>

# <翻译开始>
// Initialize examines the processor and sets the relevant variables above.
// This is called by the runtime package early in program initialization,
// before normal init functions are run. env is set by runtime if the OS supports
// cpu feature options in GODEBUG.
# <翻译结束>


<原文开始>
// options contains the cpu debug options that can be used in GODEBUG.
// Options are arch dependent and are added by the arch specific doinit functions.
// Features that are mandatory for the specific GOARCH should not be added to options
// (e.g. SSE2 on amd64).
<原文结束>

# <翻译开始>
// options contains the cpu debug options that can be used in GODEBUG.
// Options are arch dependent and are added by the arch specific doinit functions.
// Features that are mandatory for the specific GOARCH should not be added to options
// (e.g. SSE2 on amd64).
# <翻译结束>


<原文开始>
// Option names should be lower case. e.g. avx instead of AVX.
<原文结束>

# <翻译开始>
// Option names should be lower case. e.g. avx instead of AVX.
# <翻译结束>


<原文开始>
// whether feature value was specified in GODEBUG
<原文结束>

# <翻译开始>
// whether feature value was specified in GODEBUG
# <翻译结束>


<原文开始>
// whether feature should be enabled
<原文结束>

# <翻译开始>
// whether feature should be enabled
# <翻译结束>


<原文开始>
// processOptions enables or disables CPU feature values based on the parsed env string.
// The env string is expected to be of the form cpu.feature1=value1,cpu.feature2=value2...
// where feature names is one of the architecture specific list stored in the
// cpu packages options variable and values are either 'on' or 'off'.
// If env contains cpu.all=off then all cpu features referenced through the options
// variable are disabled. Other feature names and values result in warning messages.
<原文结束>

# <翻译开始>
// processOptions enables or disables CPU feature values based on the parsed env string.
// The env string is expected to be of the form cpu.feature1=value1,cpu.feature2=value2...
// where feature names is one of the architecture specific list stored in the
// cpu packages options variable and values are either 'on' or 'off'.
// If env contains cpu.all=off then all cpu features referenced through the options
// variable are disabled. Other feature names and values result in warning messages.
# <翻译结束>


<原文开始>
// indexByte returns the index of the first instance of c in s,
// or -1 if c is not present in s.
<原文结束>

# <翻译开始>
// indexByte returns the index of the first instance of c in s,
// or -1 if c is not present in s.
# <翻译结束>

