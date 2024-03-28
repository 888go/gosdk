
<原文开始>
// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
# <翻译结束>


<原文开始>
// SendFile wraps the TransmitFile call.
<原文结束>

# <翻译开始>
// SendFile wraps the TransmitFile call.
# <翻译结束>


<原文开始>
// TransmitFile does not work with pipes
<原文结束>

# <翻译开始>
// TransmitFile does not work with pipes
# <翻译结束>


<原文开始>
// TODO(brainman): skip calling syscall.Seek if OS allows it
<原文结束>

# <翻译开始>
// TODO(brainman): skip calling syscall.Seek if OS allows it
# <翻译结束>


<原文开始>
// Find the number of bytes offset from curpos until the end of the file.
<原文结束>

# <翻译开始>
// Find the number of bytes offset from curpos until the end of the file.
# <翻译结束>


<原文开始>
// Now seek back to the original position.
<原文结束>

# <翻译开始>
// Now seek back to the original position.
# <翻译结束>


<原文开始>
	// TransmitFile can be invoked in one call with at most
	// 2,147,483,646 bytes: the maximum value for a 32-bit integer minus 1.
	// See https://docs.microsoft.com/en-us/windows/win32/api/mswsock/nf-mswsock-transmitfile
<原文结束>

# <翻译开始>
	// TransmitFile can be invoked in one call with at most
	// 2,147,483,646 bytes: the maximum value for a 32-bit integer minus 1.
	// See https://docs.microsoft.com/en-us/windows/win32/api/mswsock/nf-mswsock-transmitfile
# <翻译结束>


<原文开始>
		// Some versions of Windows (Windows 10 1803) do not set
		// file position after TransmitFile completes.
		// So just use Seek to set file position.
<原文结束>

# <翻译开始>
		// Some versions of Windows (Windows 10 1803) do not set
		// file position after TransmitFile completes.
		// So just use Seek to set file position.
# <翻译结束>

