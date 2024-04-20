
<原文开始>
// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
# <翻译结束>


<原文开始>
// it is unlikely to have this character in the filename
<原文结束>

# <翻译开始>
// it is unlikely to have this character in the filename
# <翻译结束>


<原文开始>
	// Go is not explicitly added to the application compatibility database, so
	// these two functions should return the same thing.
<原文结束>

# <翻译开始>
	// Go is not explicitly added to the application compatibility database, so
	// these two functions should return the same thing.
# <翻译结束>


<原文开始>
// Grab a handle to the current process
<原文结束>

# <翻译开始>
// Grab a handle to the current process
# <翻译结束>


<原文开始>
// Allocate memory to store the result of the query
<原文结束>

# <翻译开始>
// Allocate memory to store the result of the query
# <翻译结束>


<原文开始>
// Set the new limits to the current ones
<原文结束>

# <翻译开始>
// Set the new limits to the current ones
# <翻译结束>


<原文开始>
				// The documentation for CommandLineToArgv takes for granted that
				// the first argument is a valid file path, and doesn't describe any
				// specific behavior for malformed arguments. Empirically it seems to
				// tolerate anything we throw at it, but if we discover cases where it
				// actually returns an error we might need to relax this check.
<原文结束>

# <翻译开始>
				// The documentation for CommandLineToArgv takes for granted that
				// the first argument is a valid file path, and doesn't describe any
				// specific behavior for malformed arguments. Empirically it seems to
				// tolerate anything we throw at it, but if we discover cases where it
				// actually returns an error we might need to relax this check.
# <翻译结束>


<原文开始>
			// Since DecomposeCommandLine can't handle this string,
			// interpret it as the raw arguments to ComposeCommandLine.
<原文结束>

# <翻译开始>
			// Since DecomposeCommandLine can't handle this string,
			// interpret it as the raw arguments to ComposeCommandLine.
# <翻译结束>


<原文开始>
					// We need to encode the arguments as UTF-16 to pass them to
					// CommandLineToArgvW, so skip inputs that are not valid: they might
					// have one or more runes converted to replacement characters.
<原文结束>

# <翻译开始>
					// We need to encode the arguments as UTF-16 to pass them to
					// CommandLineToArgvW, so skip inputs that are not valid: they might
					// have one or more runes converted to replacement characters.
# <翻译结束>


<原文开始>
		// It's ok if we compose a different command line than what was read.
		// Just check that we are able to compose something that round-trips
		// to the same results as the original.
<原文结束>

# <翻译开始>
		// It's ok if we compose a different command line than what was read.
		// Just check that we are able to compose something that round-trips
		// to the same results as the original.
# <翻译结束>


<原文开始>
					// It is possible that args[0] cannot be encoded exactly, because
					// CommandLineToArgvW doesn't unescape that argument in the same way
					// as the others: since the first argument is assumed to be the name
					// of the program itself, we only have the option of quoted or not.
					//
					// If args[0] contains a space or control character, we must quote it
					// to avoid it being split into multiple arguments.
					// If args[0] already starts with a quote character, we have no way
					// to indicate that that character is part of the literal argument.
					// In either case, if the string already contains a quote character
					// we must avoid misinterpriting that character as the end of the
					// quoted argument string.
					//
					// Unfortunately, ComposeCommandLine does not return an error, so we
					// can't report existing quote characters as errors.
					// Instead, we strip out the problematic quote characters from the
					// argument, and quote the remainder.
					// For paths like C:\"Program Files"\Go\bin\go.exe that is arguably
					// what the caller intended anyway, and for other strings it seems
					// less harmful than corrupting the subsequent arguments.
<原文结束>

# <翻译开始>
					// It is possible that args[0] cannot be encoded exactly, because
					// CommandLineToArgvW doesn't unescape that argument in the same way
					// as the others: since the first argument is assumed to be the name
					// of the program itself, we only have the option of quoted or not.
					//
					// If args[0] contains a space or control character, we must quote it
					// to avoid it being split into multiple arguments.
					// If args[0] already starts with a quote character, we have no way
					// to indicate that that character is part of the literal argument.
					// In either case, if the string already contains a quote character
					// we must avoid misinterpriting that character as the end of the
					// quoted argument string.
					//
					// Unfortunately, ComposeCommandLine does not return an error, so we
					// can't report existing quote characters as errors.
					// Instead, we strip out the problematic quote characters from the
					// argument, and quote the remainder.
					// For paths like C:\"Program Files"\Go\bin\go.exe that is arguably
					// what the caller intended anyway, and for other strings it seems
					// less harmful than corrupting the subsequent arguments.
# <翻译结束>


<原文开始>
// Now that we've verified the legitimate file verifies, let's corrupt it and see if it correctly fails.
<原文结束>

# <翻译开始>
// Now that we've verified the legitimate file verifies, let's corrupt it and see if it correctly fails.
# <翻译结束>


<原文开始>
// Regression check for go.dev/issue/60223
<原文结束>

# <翻译开始>
// Regression check for go.dev/issue/60223
# <翻译结束>


<原文开始>
	// Most likely, this should be [0, 4].
	// 0 is the system idle pseudo-process. 4 is the initial system process ID.
	// This test expects that at least one of the PIDs is not 0.
<原文结束>

# <翻译开始>
	// Most likely, this should be [0, 4].
	// 0 is the system idle pseudo-process. 4 is the initial system process ID.
	// This test expects that at least one of the PIDs is not 0.
# <翻译结束>


<原文开始>
// NB: Assume that we're always the first module. This technically isn't documented anywhere (that I could find), but seems to always hold.
<原文结束>

# <翻译开始>
// NB: Assume that we're always the first module. This technically isn't documented anywhere (that I could find), but seems to always hold.
# <翻译结束>


<原文开始>
// Open test directory with NtCreateFile.
<原文结束>

# <翻译开始>
// Open test directory with NtCreateFile.
# <翻译结束>


<原文开始>
// Create a file in test directory with NtCreateFile.
<原文结束>

# <翻译开始>
// Create a file in test directory with NtCreateFile.
# <翻译结束>


<原文开始>
// Rename file with NtSetInformationFile.
<原文结束>

# <翻译开始>
// Rename file with NtSetInformationFile.
# <翻译结束>


<原文开始>
// We allocate handles in a closure to provoke a UaF in the case of attributeList.Update being buggy.
<原文结束>

# <翻译开始>
// We allocate handles in a closure to provoke a UaF in the case of attributeList.Update being buggy.
# <翻译结束>


<原文开始>
// Go 1.16's pipe handles aren't inheritable, so mark it explicitly as such here.
<原文结束>

# <翻译开始>
// Go 1.16's pipe handles aren't inheritable, so mark it explicitly as such here.
# <翻译结束>


<原文开始>
// no more data available - break the loop
<原文结束>

# <翻译开始>
// no more data available - break the loop
# <翻译结束>


<原文开始>
// buffer is too small - reallocate and try again
<原文结束>

# <翻译开始>
// buffer is too small - reallocate and try again
# <翻译结束>


<原文开始>
// found a record - display the item and fetch next item
<原文结束>

# <翻译开始>
// found a record - display the item and fetch next item
# <翻译结束>


<原文开始>
// see https://go.dev/issue/31316
<原文结束>

# <翻译开始>
// see https://go.dev/issue/31316
# <翻译结束>

