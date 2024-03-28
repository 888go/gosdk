
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
// CopyFileRange copies at most remain bytes of data from src to dst, using
// the copy_file_range system call. dst and src must refer to regular files.
<原文结束>

# <翻译开始>
// CopyFileRange copies at most remain bytes of data from src to dst, using
// the copy_file_range system call. dst and src must refer to regular files.
# <翻译结束>


<原文开始>
		// copy_file_range(2) is broken in various ways on kernels older than 5.3,
		// see issue #42400 and
		// https://man7.org/linux/man-pages/man2/copy_file_range.2.html#VERSIONS
<原文结束>

# <翻译开始>
		// copy_file_range(2) is broken in various ways on kernels older than 5.3,
		// see issue #42400 and
		// https://man7.org/linux/man-pages/man2/copy_file_range.2.html#VERSIONS
# <翻译结束>


<原文开始>
			// copy_file_range(2) was introduced in Linux 4.5.
			// Go supports Linux >= 2.6.33, so the system call
			// may not be present.
			//
			// If we see ENOSYS, we have certainly not transferred
			// any data, so we can tell the caller that we
			// couldn't handle the transfer and let them fall
			// back to more generic code.
<原文结束>

# <翻译开始>
			// copy_file_range(2) was introduced in Linux 4.5.
			// Go supports Linux >= 2.6.33, so the system call
			// may not be present.
			//
			// If we see ENOSYS, we have certainly not transferred
			// any data, so we can tell the caller that we
			// couldn't handle the transfer and let them fall
			// back to more generic code.
# <翻译结束>


<原文开始>
			// Prior to Linux 5.3, it was not possible to
			// copy_file_range across file systems. Similarly to
			// the ENOSYS case above, if we see EXDEV, we have
			// not transferred any data, and we can let the caller
			// fall back to generic code.
			//
			// As for EINVAL, that is what we see if, for example,
			// dst or src refer to a pipe rather than a regular
			// file. This is another case where no data has been
			// transferred, so we consider it unhandled.
			//
			// If src and dst are on CIFS, we can see EIO.
			// See issue #42334.
			//
			// If the file is on NFS, we can see EOPNOTSUPP.
			// See issue #40731.
			//
			// If the process is running inside a Docker container,
			// we might see EPERM instead of ENOSYS. See issue
			// #40893. Since EPERM might also be a legitimate error,
			// don't mark copy_file_range(2) as unsupported.
<原文结束>

# <翻译开始>
			// Prior to Linux 5.3, it was not possible to
			// copy_file_range across file systems. Similarly to
			// the ENOSYS case above, if we see EXDEV, we have
			// not transferred any data, and we can let the caller
			// fall back to generic code.
			//
			// As for EINVAL, that is what we see if, for example,
			// dst or src refer to a pipe rather than a regular
			// file. This is another case where no data has been
			// transferred, so we consider it unhandled.
			//
			// If src and dst are on CIFS, we can see EIO.
			// See issue #42334.
			//
			// If the file is on NFS, we can see EOPNOTSUPP.
			// See issue #40731.
			//
			// If the process is running inside a Docker container,
			// we might see EPERM instead of ENOSYS. See issue
			// #40893. Since EPERM might also be a legitimate error,
			// don't mark copy_file_range(2) as unsupported.
# <翻译结束>


<原文开始>
				// If we did not read any bytes at all,
				// then this file may be in a file system
				// where copy_file_range silently fails.
				// https://lore.kernel.org/linux-fsdevel/20210126233840.GG4626@dread.disaster.area/T/#m05753578c7f7882f6e9ffe01f981bc223edef2b0
<原文结束>

# <翻译开始>
				// If we did not read any bytes at all,
				// then this file may be in a file system
				// where copy_file_range silently fails.
				// https://lore.kernel.org/linux-fsdevel/20210126233840.GG4626@dread.disaster.area/T/#m05753578c7f7882f6e9ffe01f981bc223edef2b0
# <翻译结束>


<原文开始>
				// Otherwise src is at EOF, which means
				// we are done.
<原文结束>

# <翻译开始>
				// Otherwise src is at EOF, which means
				// we are done.
# <翻译结束>


<原文开始>
// copyFileRange performs one round of copy_file_range(2).
<原文结束>

# <翻译开始>
// copyFileRange performs one round of copy_file_range(2).
# <翻译结束>


<原文开始>
	// The signature of copy_file_range(2) is:
	//
	// ssize_t copy_file_range(int fd_in, loff_t *off_in,
	//                         int fd_out, loff_t *off_out,
	//                         size_t len, unsigned int flags);
	//
	// Note that in the call to unix.CopyFileRange below, we use nil
	// values for off_in and off_out. For the system call, this means
	// "use and update the file offsets". That is why we must acquire
	// locks for both file descriptors (and why this whole machinery is
	// in the internal/poll package to begin with).
<原文结束>

# <翻译开始>
	// The signature of copy_file_range(2) is:
	//
	// ssize_t copy_file_range(int fd_in, loff_t *off_in,
	//                         int fd_out, loff_t *off_out,
	//                         size_t len, unsigned int flags);
	//
	// Note that in the call to unix.CopyFileRange below, we use nil
	// values for off_in and off_out. For the system call, this means
	// "use and update the file offsets". That is why we must acquire
	// locks for both file descriptors (and why this whole machinery is
	// in the internal/poll package to begin with).
# <翻译结束>

