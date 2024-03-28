
<原文开始>
// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
# <翻译结束>


<原文开始>
// HWCap may be initialized by archauxv and
// should not be changed after it was initialized.
<原文结束>

# <翻译开始>
// HWCap may be initialized by archauxv and
// should not be changed after it was initialized.
# <翻译结束>


<原文开始>
// HWCAP bits. These are exposed by Linux.
<原文结束>

# <翻译开始>
// HWCAP bits. These are exposed by Linux.
# <翻译结束>


<原文开始>
	// HWCap was populated by the runtime from the auxiliary vector.
	// Use HWCap information since reading aarch64 system registers
	// is not supported in user space on older linux kernels.
<原文结束>

# <翻译开始>
	// HWCap was populated by the runtime from the auxiliary vector.
	// Use HWCap information since reading aarch64 system registers
	// is not supported in user space on older linux kernels.
# <翻译结束>


<原文开始>
	// The Samsung S9+ kernel reports support for atomics, but not all cores
	// actually support them, resulting in SIGILL. See issue #28431.
	// TODO(elias.naur): Only disable the optimization on bad chipsets on android.
<原文结束>

# <翻译开始>
	// The Samsung S9+ kernel reports support for atomics, but not all cores
	// actually support them, resulting in SIGILL. See issue #28431.
	// TODO(elias.naur): Only disable the optimization on bad chipsets on android.
# <翻译结束>


<原文开始>
	// Check to see if executing on a NeoverseN1 and in order to do that,
	// check the AUXV for the CPUID bit. The getMIDR function executes an
	// instruction which would normally be an illegal instruction, but it's
	// trapped by the kernel, the value sanitized and then returned. Without
	// the CPUID bit the kernel will not trap the instruction and the process
	// will be terminated with SIGILL.
<原文结束>

# <翻译开始>
	// Check to see if executing on a NeoverseN1 and in order to do that,
	// check the AUXV for the CPUID bit. The getMIDR function executes an
	// instruction which would normally be an illegal instruction, but it's
	// trapped by the kernel, the value sanitized and then returned. Without
	// the CPUID bit the kernel will not trap the instruction and the process
	// will be terminated with SIGILL.
# <翻译结束>

