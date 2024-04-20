
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
//sys	regCreateKeyEx(key syscall.Handle, subkey *uint16, reserved uint32, class *uint16, options uint32, desired uint32, sa *syscall.SecurityAttributes, result *syscall.Handle, disposition *uint32) (regerrno error) = advapi32.RegCreateKeyExW
//sys	regDeleteKey(key syscall.Handle, subkey *uint16) (regerrno error) = advapi32.RegDeleteKeyW
//sys	regSetValueEx(key syscall.Handle, valueName *uint16, reserved uint32, vtype uint32, buf *byte, bufsize uint32) (regerrno error) = advapi32.RegSetValueExW
//sys	regEnumValue(key syscall.Handle, index uint32, name *uint16, nameLen *uint32, reserved *uint32, valtype *uint32, buf *byte, buflen *uint32) (regerrno error) = advapi32.RegEnumValueW
//sys	regDeleteValue(key syscall.Handle, name *uint16) (regerrno error) = advapi32.RegDeleteValueW
//sys   regLoadMUIString(key syscall.Handle, name *uint16, buf *uint16, buflen uint32, buflenCopied *uint32, flags uint32, dir *uint16) (regerrno error) = advapi32.RegLoadMUIStringW
//sys	regConnectRegistry(machinename *uint16, key syscall.Handle, result *syscall.Handle) (regerrno error) = advapi32.RegConnectRegistryW
<原文结束>

# <翻译开始>
//sys	regCreateKeyEx(key syscall.Handle, subkey *uint16, reserved uint32, class *uint16, options uint32, desired uint32, sa *syscall.SecurityAttributes, result *syscall.Handle, disposition *uint32) (regerrno error) = advapi32.RegCreateKeyExW
//sys	regDeleteKey(key syscall.Handle, subkey *uint16) (regerrno error) = advapi32.RegDeleteKeyW
//sys	regSetValueEx(key syscall.Handle, valueName *uint16, reserved uint32, vtype uint32, buf *byte, bufsize uint32) (regerrno error) = advapi32.RegSetValueExW
//sys	regEnumValue(key syscall.Handle, index uint32, name *uint16, nameLen *uint32, reserved *uint32, valtype *uint32, buf *byte, buflen *uint32) (regerrno error) = advapi32.RegEnumValueW
//sys	regDeleteValue(key syscall.Handle, name *uint16) (regerrno error) = advapi32.RegDeleteValueW
//sys   regLoadMUIString(key syscall.Handle, name *uint16, buf *uint16, buflen uint32, buflenCopied *uint32, flags uint32, dir *uint16) (regerrno error) = advapi32.RegLoadMUIStringW
//sys	regConnectRegistry(machinename *uint16, key syscall.Handle, result *syscall.Handle) (regerrno error) = advapi32.RegConnectRegistryW
# <翻译结束>


<原文开始>
//sys	expandEnvironmentStrings(src *uint16, dst *uint16, size uint32) (n uint32, err error) = kernel32.ExpandEnvironmentStringsW
<原文结束>

# <翻译开始>
//sys	expandEnvironmentStrings(src *uint16, dst *uint16, size uint32) (n uint32, err error) = kernel32.ExpandEnvironmentStringsW
# <翻译结束>

