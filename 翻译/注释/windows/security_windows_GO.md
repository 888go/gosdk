
<原文开始>
// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// 版权所有 2012 The Go 作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。
# <翻译结束>


<原文开始>
// TranslateAccountName converts a directory service
// object name from one format to another.
<原文结束>

# <翻译开始>
// TranslateAccountName 将目录服务对象名称从一种格式转换为另一种格式。
# <翻译结束>


<原文开始>
// Predefined domain-relative RIDs for local groups.
// See https://msdn.microsoft.com/en-us/library/windows/desktop/aa379649(v=vs.85).aspx
<原文结束>

# <翻译开始>
// 预定义的与域相关的本地组RID。 
// 参见 https://msdn.microsoft.com/en-us/library/windows/desktop/aa379649(v=vs.85).aspx
# <翻译结束>


<原文开始>
// The security identifier (SID) structure is a variable-length
// structure used to uniquely identify users or groups.
<原文结束>

# <翻译开始>
// 安全标识符（SID）结构是一种可变长度的结构，用于唯一标识用户或组。
# <翻译结束>


<原文开始>
// StringToSid converts a string-format security identifier
// SID into a valid, functional SID.
<原文结束>

# <翻译开始>
// StringToSid将字符串格式的安全标识符（SID）转换为有效且可用的SID。
# <翻译结束>


<原文开始>
// LookupSID retrieves a security identifier SID for the account
// and the name of the domain on which the account was found.
// System specify target computer to search.
<原文结束>

# <翻译开始>
// LookupSID 获取指定账户的安全标识符（SID）以及该账户所在的域名名称。
// 参数System指定了进行搜索的目标计算机。
# <翻译结束>


<原文开始>
// String converts SID to a string format suitable for display, storage, or transmission.
<原文结束>

# <翻译开始>
// String 将SID转换为适合显示、存储或传输的字符串格式
# <翻译结束>


<原文开始>
// Len returns the length, in bytes, of a valid security identifier SID.
<原文结束>

# <翻译开始>
// Len 返回一个有效安全标识符（SID）的长度，以字节为单位。
# <翻译结束>


<原文开始>
// Copy creates a duplicate of security identifier SID.
<原文结束>

# <翻译开始>
// Copy 创建一个安全标识符 SID 的副本
# <翻译结束>


<原文开始>
// IdentifierAuthority returns the identifier authority of the SID.
<原文结束>

# <翻译开始>
// IdentifierAuthority 返回该SID的标识符权限。
# <翻译结束>


<原文开始>
// SubAuthorityCount returns the number of sub-authorities in the SID.
<原文结束>

# <翻译开始>
// SubAuthorityCount 返回SID中子授权机构的数量
# <翻译结束>


<原文开始>
// SubAuthority returns the sub-authority of the SID as specified by
// the index, which must be less than sid.SubAuthorityCount().
<原文结束>

# <翻译开始>
// SubAuthority 返回由指定索引（该索引必须小于 sid.SubAuthorityCount()）所确定的 SID 的子权限部分
# <翻译结束>


<原文开始>
// IsValid returns whether the SID has a valid revision and length.
<原文结束>

# <翻译开始>
// IsValid 判断SID是否具有有效的修订版本和长度
# <翻译结束>


<原文开始>
// Equals compares two SIDs for equality.
<原文结束>

# <翻译开始>
// Equals 用于比较两个 SID 是否相等。
# <翻译结束>


<原文开始>
// IsWellKnown determines whether the SID matches the well-known sidType.
<原文结束>

# <翻译开始>
// IsWellKnown 判断SID是否匹配已知的sidType
# <翻译结束>


<原文开始>
// LookupAccount retrieves the name of the account for this SID
// and the name of the first domain on which this SID is found.
// System specify target computer to search for.
<原文结束>

# <翻译开始>
// LookupAccount 通过此 SID 查找账户名以及该 SID 首次出现的域的名称。
// System 指定要在其中搜索的目标计算机。
# <翻译结束>


<原文开始>
// Various types of pre-specified SIDs that can be synthesized and compared at runtime.
<原文结束>

# <翻译开始>
// 各种类型的预指定SID，可以在运行时进行合成与比较
# <翻译结束>


<原文开始>
// Creates a SID for a well-known predefined alias, generally using the constants of the form
// Win*Sid, for the local machine.
<原文结束>

# <翻译开始>
// 为本地计算机上某个预定义的知名别名创建一个SID（安全标识符），通常使用形如Win*Sid的常量。
# <翻译结束>


<原文开始>
// Creates a SID for a well-known predefined alias, generally using the constants of the form
// Win*Sid, for the domain specified by the domainSid parameter.
<原文结束>

# <翻译开始>
// 为指定由 domainSid 参数标识的域中某个预定义的知名别名创建一个 SID（安全标识符），通常使用形如 Win*Sid 的常量。
# <翻译结束>


<原文开始>
// Group attributes inside of Tokengroups.Groups[i].Attributes
<原文结束>

# <翻译开始>
// 在Tokengroups.Groups[i].Attributes内部对属性进行分组
# <翻译结束>


<原文开始>
// Use AllGroups() for iterating.
<原文结束>

# <翻译开始>
// 使用 AllGroups() 进行迭代。
# <翻译结束>


<原文开始>
// AllGroups returns a slice that can be used to iterate over the groups in g.
<原文结束>

# <翻译开始>
// AllGroups 返回一个切片，可用于遍历 g 中的各组。
# <翻译结束>


<原文开始>
// Use AllPrivileges() for iterating.
<原文结束>

# <翻译开始>
// 使用 AllPrivileges() 进行遍历。
# <翻译结束>


<原文开始>
// AllPrivileges returns a slice that can be used to iterate over the privileges in p.
<原文结束>

# <翻译开始>
// AllPrivileges 返回一个切片，可用于遍历集合p中的所有权限。
# <翻译结束>


<原文开始>
// An access token contains the security information for a logon session.
// The system creates an access token when a user logs on, and every
// process executed on behalf of the user has a copy of the token.
// The token identifies the user, the user's groups, and the user's
// privileges. The system uses the token to control access to securable
// objects and to control the ability of the user to perform various
// system-related operations on the local computer.
<原文结束>

# <翻译开始>
// 访问令牌包含登录会话的安全信息。
// 系统在用户登录时创建一个访问令牌，且为该用户执行的每个进程都有一份该令牌的副本。
// 该令牌用于标识用户、用户的组以及用户的权限。系统使用该令牌来控制对可保护对象的访问，并控制用户在本地计算机上执行各种系统相关操作的能力。
# <翻译结束>


<原文开始>
// OpenCurrentProcessToken opens an access token associated with current
// process with TOKEN_QUERY access. It is a real token that needs to be closed.
//
// Deprecated: Explicitly call OpenProcessToken(CurrentProcess(), ...)
// with the desired access instead, or use GetCurrentProcessToken for a
// TOKEN_QUERY token.
<原文结束>

# <翻译开始>
// OpenCurrentProcessToken 以 TOKEN_QUERY 访问权限打开与当前进程关联的访问令牌。这是一个需要被关闭的真实令牌。
//
// 已弃用：请显式调用 OpenProcessToken(CurrentProcess(), ...) 并指定所需的访问权限，或者使用 GetCurrentProcessToken 获取 TOKEN_QUERY 令牌。
# <翻译结束>


<原文开始>
// GetCurrentProcessToken returns the access token associated with
// the current process. It is a pseudo token that does not need
// to be closed.
<原文结束>

# <翻译开始>
// GetCurrentProcessToken 获取与当前进程关联的访问令牌。这是一个无需关闭的伪令牌。
# <翻译结束>


<原文开始>
// GetCurrentThreadToken return the access token associated with
// the current thread. It is a pseudo token that does not need
// to be closed.
<原文结束>

# <翻译开始>
// GetCurrentThreadToken 返回当前线程关联的访问令牌。它是一个无需关闭的伪令牌。
# <翻译结束>


<原文开始>
// GetCurrentThreadEffectiveToken returns the effective access token
// associated with the current thread. It is a pseudo token that does
// not need to be closed.
<原文结束>

# <翻译开始>
// GetCurrentThreadEffectiveToken 返回与当前线程关联的有效访问令牌。
// 它是一个无需关闭的伪令牌。
# <翻译结束>


<原文开始>
// Close releases access to access token.
<原文结束>

# <翻译开始>
// Close 释放对访问令牌的访问权限。
# <翻译结束>


<原文开始>
// getInfo retrieves a specified type of information about an access token.
<原文结束>

# <翻译开始>
// getInfo 用于获取访问令牌指定类型的信息。
# <翻译结束>


<原文开始>
// GetTokenUser retrieves access token t user account information.
<原文结束>

# <翻译开始>
// GetTokenUser 用于获取访问令牌对应的用户账户信息
# <翻译结束>


<原文开始>
// GetTokenGroups retrieves group accounts associated with access token t.
<原文结束>

# <翻译开始>
// GetTokenGroups 获取与访问令牌 t 关联的组账户。
# <翻译结束>


<原文开始>
// GetTokenPrimaryGroup retrieves access token t primary group information.
// A pointer to a SID structure representing a group that will become
// the primary group of any objects created by a process using this access token.
<原文结束>

# <翻译开始>
// GetTokenPrimaryGroup 获取访问令牌 t 的主组信息。
// 返回一个指向 SID 结构的指针，该结构代表一个组，此组将成为使用此访问令牌的进程创建的任何对象的主组。
# <翻译结束>


<原文开始>
// GetUserProfileDirectory retrieves path to the
// root directory of the access token t user's profile.
<原文结束>

# <翻译开始>
// GetUserProfileDirectory 获取访问令牌t对应用户的个人资料根目录路径。
# <翻译结束>


<原文开始>
// IsElevated returns whether the current token is elevated from a UAC perspective.
<原文结束>

# <翻译开始>
// IsElevated 返回当前令牌是否从UAC角度来看具有提升权限。
# <翻译结束>


<原文开始>
// GetLinkedToken returns the linked token, which may be an elevated UAC token.
<原文结束>

# <翻译开始>
// 获取关联的令牌，该令牌可能是一个提升的UAC令牌。
# <翻译结束>


<原文开始>
// GetSystemDirectory retrieves the path to current location of the system
// directory, which is typically, though not always, `C:\Windows\System32`.
<原文结束>

# <翻译开始>
// GetSystemDirectory 获取当前系统目录的路径，该目录通常（但并不总是）为 `C:\Windows\System32`。
# <翻译结束>


<原文开始>
// GetWindowsDirectory retrieves the path to current location of the Windows
// directory, which is typically, though not always, `C:\Windows`. This may
// be a private user directory in the case that the application is running
// under a terminal server.
<原文结束>

# <翻译开始>
// GetWindowsDirectory 函数用于获取当前 Windows 目录的路径，该目录通常（但并非总是）为 `C:\Windows`。在应用程序于终端服务器环境下运行时，此目录可能为用户的个人专属目录。
# <翻译结束>


<原文开始>
// GetSystemWindowsDirectory retrieves the path to current location of the
// Windows directory, which is typically, though not always, `C:\Windows`.
<原文结束>

# <翻译开始>
// GetSystemWindowsDirectory 获取当前 Windows 目录的路径，该目录通常（但并不总是）位于 `C:\Windows`。
# <翻译结束>


<原文开始>
// IsMember reports whether the access token t is a member of the provided SID.
<原文结束>

# <翻译开始>
// IsMember 判断访问令牌 t 是否为给定 SID 的成员。
# <翻译结束>


<原文开始>
// IsRestricted reports whether the access token t is a restricted token.
<原文结束>

# <翻译开始>
// IsRestricted 判断访问令牌 t 是否为受限令牌。
# <翻译结束>


<原文开始>
// If err is EINVAL, this returned ERROR_SUCCESS indicating a non-restricted token.
<原文结束>

# <翻译开始>
// 如果err为EINVAL，返回ERROR_SUCCESS，表示该令牌不受限制
# <翻译结束>


<原文开始>
// Constants for the ContextTrackingMode field of SECURITY_QUALITY_OF_SERVICE.
<原文结束>

# <翻译开始>
// ContextTrackingMode字段的SECURITY_QUALITY_OF_SERVICE常量
# <翻译结束>


<原文开始>
// Constants for type SE_OBJECT_TYPE
<原文结束>

# <翻译开始>
// SE_OBJECT_TYPE类型的常量
# <翻译结束>


<原文开始>
// Constants for type SECURITY_INFORMATION
<原文结束>

# <翻译开始>
// SECURITY_INFORMATION 类型的常量
# <翻译结束>


<原文开始>
// Constants for type SECURITY_DESCRIPTOR_CONTROL
<原文结束>

# <翻译开始>
// 定义类型SECURITY_DESCRIPTOR_CONTROL的常量
# <翻译结束>


<原文开始>
// Constants for type ACCESS_MASK
<原文结束>

# <翻译开始>
// ACCESS_MASK类型的常量
# <翻译结束>


<原文开始>
// Constants for type ACCESS_MODE
<原文结束>

# <翻译开始>
// ACCESS_MODE类型的常量
# <翻译结束>


<原文开始>
// Constants for AceFlags and Inheritance fields
<原文结束>

# <翻译开始>
// 用于AceFlags和Inheritance字段的常量
# <翻译结束>


<原文开始>
// Constants for MULTIPLE_TRUSTEE_OPERATION
<原文结束>

# <翻译开始>
// MULTIPLE_TRUSTEE_OPERATION相关的常量
# <翻译结束>


<原文开始>
// Constants for TRUSTEE_FORM
<原文结束>

# <翻译开始>
// TRUSTEE_FORM的常量
# <翻译结束>


<原文开始>
// Constants for TRUSTEE_TYPE
<原文结束>

# <翻译开始>
// TRUSTEE_TYPE类型的常量
# <翻译结束>


<原文开始>
// Constants for ObjectsPresent field
<原文结束>

# <翻译开始>
// 对于ObjectsPresent字段的常量
# <翻译结束>


<原文开始>
// This type is the union inside of TRUSTEE and must be created using one of the TrusteeValueFrom* functions.
<原文结束>

# <翻译开始>
// 此类型为TRUSTEE内的联合体，必须通过使用TrusteeValueFrom*系列函数之一来创建。
# <翻译结束>


<原文开始>
// Control returns the security descriptor control bits.
<原文结束>

# <翻译开始>
// Control 返回安全描述符控制位。
# <翻译结束>


<原文开始>
// SetControl sets the security descriptor control bits.
<原文结束>

# <翻译开始>
// SetControl 设置安全描述符控制位
# <翻译结束>


<原文开始>
// RMControl returns the security descriptor resource manager control bits.
<原文结束>

# <翻译开始>
// RMControl 返回安全描述符资源管理器控制位。
# <翻译结束>


<原文开始>
// SetRMControl sets the security descriptor resource manager control bits.
<原文结束>

# <翻译开始>
// SetRMControl 设置安全描述符资源管理器控制位
# <翻译结束>


<原文开始>
// DACL returns the security descriptor DACL and whether it was defaulted. The dacl return value may be nil
// if a DACL exists but is an "empty DACL", meaning fully permissive. If the DACL does not exist, err returns
// ERROR_OBJECT_NOT_FOUND.
<原文结束>

# <翻译开始>
// DACL 返回安全描述符的 DACL 以及它是否为默认值。dacl 的返回值可能为 nil，如果存在但为“空 DACL”，即表示完全无限制访问。如果不存在 DACL，则 err 返回 ERROR_OBJECT_NOT_FOUND。
# <翻译结束>


<原文开始>
// SetDACL sets the absolute security descriptor DACL.
<原文结束>

# <翻译开始>
// SetDACL 设置绝对安全描述符的 DACL
# <翻译结束>


<原文开始>
// SACL returns the security descriptor SACL and whether it was defaulted. The sacl return value may be nil
// if a SACL exists but is an "empty SACL", meaning fully permissive. If the SACL does not exist, err returns
// ERROR_OBJECT_NOT_FOUND.
<原文结束>

# <翻译开始>
// SACL 返回安全描述符的 SACL（系统访问控制列表）及其是否为默认值。如果存在 SACL，但它是“空 SACL”，即完全无限制，则 sacl 返回值可能为 nil。如果不存在 SACL，则 err 返回 ERROR_OBJECT_NOT_FOUND。
# <翻译结束>


<原文开始>
// SetSACL sets the absolute security descriptor SACL.
<原文结束>

# <翻译开始>
// SetSACL 设置绝对安全描述符的SACL
# <翻译结束>


<原文开始>
// Owner returns the security descriptor owner and whether it was defaulted.
<原文结束>

# <翻译开始>
// Owner 返回安全描述符的所有者以及该所有者是否为默认值。
# <翻译结束>


<原文开始>
// SetOwner sets the absolute security descriptor owner.
<原文结束>

# <翻译开始>
// SetOwner 设置绝对安全描述符的所有者
# <翻译结束>


<原文开始>
// Group returns the security descriptor group and whether it was defaulted.
<原文结束>

# <翻译开始>
// Group 返回安全描述符的组信息以及该组信息是否为默认值。
# <翻译结束>


<原文开始>
// SetGroup sets the absolute security descriptor owner.
<原文结束>

# <翻译开始>
// SetGroup 设置绝对安全描述符的所有者。
# <翻译结束>


<原文开始>
// Length returns the length of the security descriptor.
<原文结束>

# <翻译开始>
// Length 返回安全描述符的长度。
# <翻译结束>


<原文开始>
// IsValid returns whether the security descriptor is valid.
<原文结束>

# <翻译开始>
// IsValid 返回该安全描述符是否有效
# <翻译结束>


<原文开始>
// String returns the SDDL form of the security descriptor, with a function signature that can be
// used with %v formatting directives.
<原文结束>

# <翻译开始>
// String 方法返回安全描述符的 SDDL 形式，其函数签名适用于使用 %v 格式化指令。
# <翻译结束>


<原文开始>
// ToAbsolute converts a self-relative security descriptor into an absolute one.
<原文结束>

# <翻译开始>
// ToAbsolute 将一个自相对的（self-relative）安全描述符转换为绝对的（absolute）安全描述符。
# <翻译结束>


<原文开始>
// makeAbsoluteSD is expected to fail, but it succeeds.
<原文结束>

# <翻译开始>
// makeAbsoluteSD 预期会失败，但它却成功了。
# <翻译结束>


<原文开始>
// ToSelfRelative converts an absolute security descriptor into a self-relative one.
<原文结束>

# <翻译开始>
// ToSelfRelative 将一个绝对安全描述符转换为自相对形式。
# <翻译结束>


<原文开始>
// makeSelfRelativeSD is expected to fail, but it succeeds.
<原文结束>

# <翻译开始>
// makeSelfRelativeSD 预期会失败，但实际上却成功了。
# <翻译结束>


<原文开始>
	// SECURITY_DESCRIPTOR has pointers in it, which means checkptr expects for it to
	// be aligned properly. When we're copying a Windows-allocated struct to a
	// Go-allocated one, make sure that the Go allocation is aligned to the
	// pointer size.
<原文结束>

# <翻译开始>
// SECURITY_DESCRIPTOR 结构体内含有指针，这意味着 checkptr 预期其应被正确对齐。当我们将一个由 Windows 分配的结构体复制到由 Go 分配的一个时，确保 Go 的内存分配按指针大小对齐。
# <翻译结束>


<原文开始>
// SecurityDescriptorFromString converts an SDDL string describing a security descriptor into a
// self-relative security descriptor object allocated on the Go heap.
<原文结束>

# <翻译开始>
// SecurityDescriptorFromString 将描述安全描述符的 SDDL 字符串转换为在 Go 堆上分配的自相对性安全描述符对象。
# <翻译结束>


<原文开始>
// GetSecurityInfo queries the security information for a given handle and returns the self-relative security
// descriptor result on the Go heap.
<原文结束>

# <翻译开始>
// GetSecurityInfo 为给定的句柄查询安全信息，并在Go堆上返回自相关的安全描述符结果。
# <翻译结束>


<原文开始>
// GetNamedSecurityInfo queries the security information for a given named object and returns the self-relative security
// descriptor result on the Go heap.
<原文结束>

# <翻译开始>
// GetNamedSecurityInfo 为指定命名对象查询安全信息，并在Go堆上返回自相对安全描述符结果。
# <翻译结束>


<原文开始>
// BuildSecurityDescriptor makes a new security descriptor using the input trustees, explicit access lists, and
// prior security descriptor to be merged, any of which can be nil, returning the self-relative security descriptor
// result on the Go heap.
<原文结束>

# <翻译开始>
// BuildSecurityDescriptor 通过使用输入的信任主体、显式访问列表以及待合并的先前安全描述符（这些参数均可为 nil）构建一个新的安全描述符，并将生成的自相关安全描述符结果置于 Go 堆上返回。
# <翻译结束>


<原文开始>
// NewSecurityDescriptor creates and initializes a new absolute security descriptor.
<原文结束>

# <翻译开始>
// NewSecurityDescriptor 创建并初始化一个新的绝对安全描述符。
# <翻译结束>


<原文开始>
// ACLFromEntries returns a new ACL on the Go heap containing a list of explicit entries as well as those of another ACL.
// Both explicitEntries and mergedACL are optional and can be nil.
<原文结束>

# <翻译开始>
// ACLFromEntries 返回一个位于Go堆上的新ACL，其中包含一个显式条目列表以及另一个ACL的条目。
// explicitEntries 和 mergedACL 两者都是可选的，可以为nil。
# <翻译结束>

