
<原文开始>
// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
# <翻译结束>


<原文开始>
// Package registry provides access to the Windows registry.
//
// Here is a simple example, opening a registry key and reading a string value from it.
//
//	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows NT\CurrentVersion`, registry.QUERY_VALUE)
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer k.Close()
//
//	s, _, err := k.GetStringValue("SystemRoot")
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Printf("Windows system root is %q\n", s)
//
// NOTE: This package is a copy of golang.org/x/sys/windows/registry
// with KeyInfo.ModTime removed to prevent dependency cycles.
<原文结束>

# <翻译开始>
// Package registry provides access to the Windows registry.
//
// Here is a simple example, opening a registry key and reading a string value from it.
//
//	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows NT\CurrentVersion`, registry.QUERY_VALUE)
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer k.Close()
//
//	s, _, err := k.GetStringValue("SystemRoot")
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Printf("Windows system root is %q\n", s)
//
// NOTE: This package is a copy of golang.org/x/sys/windows/registry
// with KeyInfo.ModTime removed to prevent dependency cycles.
# <翻译结束>


<原文开始>
	// Registry key security and access rights.
	// See https://msdn.microsoft.com/en-us/library/windows/desktop/ms724878.aspx
	// for details.
<原文结束>

# <翻译开始>
	// Registry key security and access rights.
	// See https://msdn.microsoft.com/en-us/library/windows/desktop/ms724878.aspx
	// for details.
# <翻译结束>


<原文开始>
// Key is a handle to an open Windows registry key.
// Keys can be obtained by calling OpenKey; there are
// also some predefined root keys such as CURRENT_USER.
// Keys can be used directly in the Windows API.
<原文结束>

# <翻译开始>
// Key is a handle to an open Windows registry key.
// Keys can be obtained by calling OpenKey; there are
// also some predefined root keys such as CURRENT_USER.
// Keys can be used directly in the Windows API.
# <翻译结束>


<原文开始>
	// Windows defines some predefined root keys that are always open.
	// An application can use these keys as entry points to the registry.
	// Normally these keys are used in OpenKey to open new keys,
	// but they can also be used anywhere a Key is required.
<原文结束>

# <翻译开始>
	// Windows defines some predefined root keys that are always open.
	// An application can use these keys as entry points to the registry.
	// Normally these keys are used in OpenKey to open new keys,
	// but they can also be used anywhere a Key is required.
# <翻译结束>


<原文开始>
// Close closes open key k.
<原文结束>

# <翻译开始>
// Close closes open key k.
# <翻译结束>


<原文开始>
// OpenKey opens a new key with path name relative to key k.
// It accepts any open key, including CURRENT_USER and others,
// and returns the new key and an error.
// The access parameter specifies desired access rights to the
// key to be opened.
<原文结束>

# <翻译开始>
// OpenKey opens a new key with path name relative to key k.
// It accepts any open key, including CURRENT_USER and others,
// and returns the new key and an error.
// The access parameter specifies desired access rights to the
// key to be opened.
# <翻译结束>


<原文开始>
// ReadSubKeyNames returns the names of subkeys of key k.
<原文结束>

# <翻译开始>
// ReadSubKeyNames returns the names of subkeys of key k.
# <翻译结束>


<原文开始>
	// RegEnumKeyEx must be called repeatedly and to completion.
	// During this time, this goroutine cannot migrate away from
	// its current thread. See #49320.
<原文结束>

# <翻译开始>
	// RegEnumKeyEx must be called repeatedly and to completion.
	// During this time, this goroutine cannot migrate away from
	// its current thread. See #49320.
# <翻译结束>


<原文开始>
	// Registry key size limit is 255 bytes and described there:
	// https://msdn.microsoft.com/library/windows/desktop/ms724872.aspx
<原文结束>

# <翻译开始>
	// Registry key size limit is 255 bytes and described there:
	// https://msdn.microsoft.com/library/windows/desktop/ms724872.aspx
# <翻译结束>


<原文开始>
//plus extra room for terminating zero byte
<原文结束>

# <翻译开始>
//plus extra room for terminating zero byte
# <翻译结束>


<原文开始>
// Double buffer size and try again.
<原文结束>

# <翻译开始>
// Double buffer size and try again.
# <翻译结束>


<原文开始>
// CreateKey creates a key named path under open key k.
// CreateKey returns the new key and a boolean flag that reports
// whether the key already existed.
// The access parameter specifies the access rights for the key
// to be created.
<原文结束>

# <翻译开始>
// CreateKey creates a key named path under open key k.
// CreateKey returns the new key and a boolean flag that reports
// whether the key already existed.
// The access parameter specifies the access rights for the key
// to be created.
# <翻译结束>


<原文开始>
// DeleteKey deletes the subkey path of key k and its values.
<原文结束>

# <翻译开始>
// DeleteKey deletes the subkey path of key k and its values.
# <翻译结束>


<原文开始>
// A KeyInfo describes the statistics of a key. It is returned by Stat.
<原文结束>

# <翻译开始>
// A KeyInfo describes the statistics of a key. It is returned by Stat.
# <翻译结束>


<原文开始>
// Stat retrieves information about the open key k.
<原文结束>

# <翻译开始>
// Stat retrieves information about the open key k.
# <翻译结束>

