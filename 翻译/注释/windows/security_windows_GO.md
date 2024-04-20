
<原文开始>
// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// 版权所有 ? 2012 The Go Authors。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。
# <翻译结束>


<原文开始>
// This function returns 1 byte BOOLEAN rather than the 4 byte BOOL.
// http://blogs.msdn.com/b/drnick/archive/2007/12/19/windows-and-upn-format-credentials.aspx
//sys	TranslateName(accName *uint16, accNameFormat uint32, desiredNameFormat uint32, translatedName *uint16, nSize *uint32) (err error) [failretval&0xff==0] = secur32.TranslateNameW
//sys	GetUserNameEx(nameFormat uint32, nameBuffre *uint16, nSize *uint32) (err error) [failretval&0xff==0] = secur32.GetUserNameExW
<原文结束>

# <翻译开始>
// 此函数返回 1 字节的 BOOLEAN 类型，而非 4 字节的 BOOL 类型。
// 参考：http://blogs.msdn.com/b/drnick/archive/2007/12/19/windows-and-upn-format-credentials.aspx
//sys TranslateName(accName *uint16, accNameFormat uint32, desiredNameFormat uint32, translatedName *uint16, nSize *uint32) (err error) [failretval&0xff==0] = secur32.TranslateNameW
//sys GetUserNameEx(nameFormat uint32, nameBuffre *uint16, nSize *uint32) (err error) [failretval&0xff==0] = secur32.GetUserNameExW
// 
// 翻译：
// 
// 该函数返回 1 字节的 BOOLEAN 值，而非 4 字节的 BOOL 值。
// 参考文献：http://blogs.msdn.com/b/drnick/archive/2007/12/19/windows-and-upn-format-credentials.aspx
//sys TranslateName(accName *uint16, accNameFormat uint32, desiredNameFormat uint32, translatedName *uint16, nSize *uint32) (err error) [failretval&0xff==0] = secur32.TranslateNameW
//sys GetUserNameEx(nameFormat uint32, nameBuffre *uint16, nSize *uint32) (err error) [failretval&0xff==0] = secur32.GetUserNameExW
# <翻译结束>


<原文开始>
// TranslateAccountName converts a directory service
// object name from one format to another.
<原文结束>

# <翻译开始>
// TranslateAccountName 将目录服务对象名从一种格式转换为另一种格式
# <翻译结束>


<原文开始>
//sys	NetUserGetInfo(serverName *uint16, userName *uint16, level uint32, buf **byte) (neterr error) = netapi32.NetUserGetInfo
//sys	NetGetJoinInformation(server *uint16, name **uint16, bufType *uint32) (neterr error) = netapi32.NetGetJoinInformation
//sys	NetApiBufferFree(buf *byte) (neterr error) = netapi32.NetApiBufferFree
<原文结束>

# <翻译开始>
//sys	NetUserGetInfo(serverName *uint16, userName *uint16, level uint32, buf **byte) (neterr error) = netapi32.NetUserGetInfo
// 系统调用：NetUserGetInfo，用于获取用户信息。参数包括：
//  - serverName：指向包含服务器名的uint16类型指针
//  - userName：指向包含用户名的uint16类型指针
//  - level：指定要查询的信息级别（uint32类型）
//  - buf：输出参数，指向接收用户信息数据的字节指针指针
// 返回值：neterr，表示调用结果的错误信息
// 此系统调用对应于netapi32库中的NetUserGetInfo函数
// 
//sys	NetGetJoinInformation(server *uint16, name **uint16, bufType *uint32) (neterr error) = netapi32.NetGetJoinInformation
// 系统调用：NetGetJoinInformation，用于获取计算机加入域或工作组的相关信息。参数包括：
//  - server：指向包含服务器名的uint16类型指针
//  - name：输出参数，指向接收计算机名称的uint16类型指针指针
//  - bufType：输出参数，指向接收计算机加入类型（域或工作组）标识符的uint32类型指针
// 返回值：neterr，表示调用结果的错误信息
// 此系统调用对应于netapi32库中的NetGetJoinInformation函数
// 
//sys	NetApiBufferFree(buf *byte) (neterr error) = netapi32.NetApiBufferFree
// 系统调用：NetApiBufferFree，用于释放由NetUserGetInfo、NetGetJoinInformation等函数返回的缓冲区。参数：
//  - buf：需要释放的缓冲区地址，类型为字节指针
// 返回值：neterr，表示释放操作的错误信息
// 此系统调用对应于netapi32库中的NetApiBufferFree函数
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
//sys	LookupAccountSid(systemName *uint16, sid *SID, name *uint16, nameLen *uint32, refdDomainName *uint16, refdDomainNameLen *uint32, use *uint32) (err error) = advapi32.LookupAccountSidW
//sys	LookupAccountName(systemName *uint16, accountName *uint16, sid *SID, sidLen *uint32, refdDomainName *uint16, refdDomainNameLen *uint32, use *uint32) (err error) = advapi32.LookupAccountNameW
//sys	ConvertSidToStringSid(sid *SID, stringSid **uint16) (err error) = advapi32.ConvertSidToStringSidW
//sys	ConvertStringSidToSid(stringSid *uint16, sid **SID) (err error) = advapi32.ConvertStringSidToSidW
//sys	GetLengthSid(sid *SID) (len uint32) = advapi32.GetLengthSid
//sys	CopySid(destSidLen uint32, destSid *SID, srcSid *SID) (err error) = advapi32.CopySid
//sys	AllocateAndInitializeSid(identAuth *SidIdentifierAuthority, subAuth byte, subAuth0 uint32, subAuth1 uint32, subAuth2 uint32, subAuth3 uint32, subAuth4 uint32, subAuth5 uint32, subAuth6 uint32, subAuth7 uint32, sid **SID) (err error) = advapi32.AllocateAndInitializeSid
//sys	createWellKnownSid(sidType WELL_KNOWN_SID_TYPE, domainSid *SID, sid *SID, sizeSid *uint32) (err error) = advapi32.CreateWellKnownSid
//sys	isWellKnownSid(sid *SID, sidType WELL_KNOWN_SID_TYPE) (isWellKnown bool) = advapi32.IsWellKnownSid
//sys	FreeSid(sid *SID) (err error) [failretval!=0] = advapi32.FreeSid
//sys	EqualSid(sid1 *SID, sid2 *SID) (isEqual bool) = advapi32.EqualSid
//sys	getSidIdentifierAuthority(sid *SID) (authority *SidIdentifierAuthority) = advapi32.GetSidIdentifierAuthority
//sys	getSidSubAuthorityCount(sid *SID) (count *uint8) = advapi32.GetSidSubAuthorityCount
//sys	getSidSubAuthority(sid *SID, index uint32) (subAuthority *uint32) = advapi32.GetSidSubAuthority
//sys	isValidSid(sid *SID) (isValid bool) = advapi32.IsValidSid
<原文结束>

# <翻译开始>
//sys	LookupAccountSid 系统名称 *uint16, SID *SID, 账户名 *uint16, 账户名长度 *uint32, 引用的域名 *uint16, 引用的域名长度 *uint32, 使用类型 *uint32) (错误 error) = advapi32.LookupAccountSidW
//sys	LookupAccountName 系统名称 *uint16, 账户名 *uint16, SID *SID, SID长度 *uint32, 引用的域名 *uint16, 引用的域名长度 *uint32, 使用类型 *uint32) (错误 error) = advapi32.LookupAccountNameW
//sys	ConvertSidToStringSid SID *SID, 字符串SID **uint16) (错误 error) = advapi32.ConvertSidToStringSidW
//sys	ConvertStringSidToSid 字符串SID *uint16, SID **SID) (错误 error) = advapi32.ConvertStringSidToSidW
//sys	GetLengthSid SID *SID) (长度 uint32) = advapi32.GetLengthSid
//sys	CopySid 目标SID长度 uint32, 目标SID *SID, 源SID *SID) (错误 error) = advapi32.CopySid
//sys	AllocateAndInitializeSid 标识权限 *SidIdentifierAuthority, 子授权字节 byte, 子授权0 uint32, 子授权1 uint32, 子授权2 uint32, 子授权3 uint32, 子授权4 uint32, 子授权5 uint32, 子授权6 uint32, 子授权7 uint32, SID **SID) (错误 error) = advapi32.AllocateAndInitializeSid
//sys	createWellKnownSid SID类型 WELL_KNOWN_SID_TYPE, 域SID *SID, SID *SID, SID大小 *uint32) (错误 error) = advapi32.CreateWellKnownSid
//sys	isWellKnownSid SID *SID, SID类型 WELL_KNOWN_SID_TYPE) (是知名SID bool) = advapi32.IsWellKnownSid
//sys	FreeSid SID *SID) (错误 error) [失败返回值！= 0] = advapi32.FreeSid
//sys	EqualSid SID1 *SID, SID2 *SID) (相等 bool) = advapi32.EqualSid
//sys	getSidIdentifierAuthority SID *SID) (权威机构 *SidIdentifierAuthority) = advapi32.GetSidIdentifierAuthority
//sys	getSidSubAuthorityCount SID *SID) (计数器 *uint8) = advapi32.GetSidSubAuthorityCount
//sys	getSidSubAuthority SID *SID, 索引 uint32) (子授权 *uint32) = advapi32.GetSidSubAuthority
//sys	isValidSid SID *SID) (有效 bool) = advapi32.IsValidSid
# <翻译结束>


<原文开始>
// The security identifier (SID) structure is a variable-length
// structure used to uniquely identify users or groups.
<原文结束>

# <翻译开始>
// 安全标识符（SID）结构是一种可变长度的结构，用于唯一地标识用户或组。
# <翻译结束>


<原文开始>
// StringToSid converts a string-format security identifier
// SID into a valid, functional SID.
<原文结束>

# <翻译开始>
// StringToSid 将字符串格式的安全标识符（SID）转换为有效且可用的 SID。
# <翻译结束>


<原文开始>
// LookupSID retrieves a security identifier SID for the account
// and the name of the domain on which the account was found.
// System specify target computer to search.
<原文结束>

# <翻译开始>
// LookupSID 通过账户名获取该账户的安全标识符（SID）以及账户所在的域名。
// 参数System用于指定要在其中进行搜索的目标计算机。
# <翻译结束>


<原文开始>
// String converts SID to a string format suitable for display, storage, or transmission.
<原文结束>

# <翻译开始>
// String 将SID转换为适合显示、存储或传输的字符串格式。
# <翻译结束>


<原文开始>
// Len returns the length, in bytes, of a valid security identifier SID.
<原文结束>

# <翻译开始>
// Len 返回一个有效安全标识符（SID）的字节长度。
# <翻译结束>


<原文开始>
// Copy creates a duplicate of security identifier SID.
<原文结束>

# <翻译开始>
// Copy 创建安全标识符 SID 的副本。
# <翻译结束>


<原文开始>
// IdentifierAuthority returns the identifier authority of the SID.
<原文结束>

# <翻译开始>
// IdentifierAuthority 返回 SID 的标识符权限。
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
// SubAuthority 返回由指定索引（该索引必须小于 sid.SubAuthorityCount()）确定的 SID 的子权限部分
# <翻译结束>


<原文开始>
// IsValid returns whether the SID has a valid revision and length.
<原文结束>

# <翻译开始>
// IsValid 判断SID是否具有有效的修订版号和长度。
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
// IsWellKnown 判断SID是否匹配已知的sidType。
# <翻译结束>


<原文开始>
// LookupAccount retrieves the name of the account for this SID
// and the name of the first domain on which this SID is found.
// System specify target computer to search for.
<原文结束>

# <翻译开始>
// LookupAccount 通过此 SID 查找账户名称以及该 SID 所在的第一个域的名称。
// 参数 System 指定要在其中搜索的目标计算机。
# <翻译结束>


<原文开始>
// Various types of pre-specified SIDs that can be synthesized and compared at runtime.
<原文结束>

# <翻译开始>
// 各种类型的预指定SID，可在运行时进行合成与比较
# <翻译结束>


<原文开始>
// Creates a SID for a well-known predefined alias, generally using the constants of the form
// Win*Sid, for the local machine.
<原文结束>

# <翻译开始>
// 为本地计算机上一个预先定义的知名别名创建一个SID（安全标识符），通常使用形式为Win*Sid的常量。
# <翻译结束>


<原文开始>
// Creates a SID for a well-known predefined alias, generally using the constants of the form
// Win*Sid, for the domain specified by the domainSid parameter.
<原文结束>

# <翻译开始>
// 为指定由 domainSid 参数所确定域内的某个预定义的知名别名创建一个 SID（安全标识符），通常使用形如 Win*Sid 的常量。
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
// AllGroups 返回一个切片，可用于遍历 g 中的所有组。
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
// AllPrivileges 返回一个切片，可用于遍历p中的所有权限。
# <翻译结束>


<原文开始>
// Authorization Functions
//sys	checkTokenMembership(tokenHandle Token, sidToCheck *SID, isMember *int32) (err error) = advapi32.CheckTokenMembership
//sys	isTokenRestricted(tokenHandle Token) (ret bool, err error) [!failretval] = advapi32.IsTokenRestricted
//sys	OpenProcessToken(process Handle, access uint32, token *Token) (err error) = advapi32.OpenProcessToken
//sys	OpenThreadToken(thread Handle, access uint32, openAsSelf bool, token *Token) (err error) = advapi32.OpenThreadToken
//sys	ImpersonateSelf(impersonationlevel uint32) (err error) = advapi32.ImpersonateSelf
//sys	RevertToSelf() (err error) = advapi32.RevertToSelf
//sys	SetThreadToken(thread *Handle, token Token) (err error) = advapi32.SetThreadToken
//sys	LookupPrivilegeValue(systemname *uint16, name *uint16, luid *LUID) (err error) = advapi32.LookupPrivilegeValueW
//sys	AdjustTokenPrivileges(token Token, disableAllPrivileges bool, newstate *Tokenprivileges, buflen uint32, prevstate *Tokenprivileges, returnlen *uint32) (err error) = advapi32.AdjustTokenPrivileges
//sys	AdjustTokenGroups(token Token, resetToDefault bool, newstate *Tokengroups, buflen uint32, prevstate *Tokengroups, returnlen *uint32) (err error) = advapi32.AdjustTokenGroups
//sys	GetTokenInformation(token Token, infoClass uint32, info *byte, infoLen uint32, returnedLen *uint32) (err error) = advapi32.GetTokenInformation
//sys	SetTokenInformation(token Token, infoClass uint32, info *byte, infoLen uint32) (err error) = advapi32.SetTokenInformation
//sys	DuplicateTokenEx(existingToken Token, desiredAccess uint32, tokenAttributes *SecurityAttributes, impersonationLevel uint32, tokenType uint32, newToken *Token) (err error) = advapi32.DuplicateTokenEx
//sys	GetUserProfileDirectory(t Token, dir *uint16, dirLen *uint32) (err error) = userenv.GetUserProfileDirectoryW
//sys	getSystemDirectory(dir *uint16, dirLen uint32) (len uint32, err error) = kernel32.GetSystemDirectoryW
//sys	getWindowsDirectory(dir *uint16, dirLen uint32) (len uint32, err error) = kernel32.GetWindowsDirectoryW
//sys	getSystemWindowsDirectory(dir *uint16, dirLen uint32) (len uint32, err error) = kernel32.GetSystemWindowsDirectoryW
<原文结束>

# <翻译开始>
// Authorization Functions
//sys	checkTokenMembership(tokenHandle Token, sidToCheck *SID, isMember *int32) (err error) = advapi32.CheckTokenMembership
//sys	isTokenRestricted(tokenHandle Token) (ret bool, err error) [!failretval] = advapi32.IsTokenRestricted
//sys	OpenProcessToken(process Handle, access uint32, token *Token) (err error) = advapi32.OpenProcessToken
//sys	OpenThreadToken(thread Handle, access uint32, openAsSelf bool, token *Token) (err error) = advapi32.OpenThreadToken
//sys	ImpersonateSelf(impersonationlevel uint32) (err error) = advapi32.ImpersonateSelf
//sys	RevertToSelf() (err error) = advapi32.RevertToSelf
//sys	SetThreadToken(thread *Handle, token Token) (err error) = advapi32.SetThreadToken
//sys	LookupPrivilegeValue(systemname *uint16, name *uint16, luid *LUID) (err error) = advapi32.LookupPrivilegeValueW
//sys	AdjustTokenPrivileges(token Token, disableAllPrivileges bool, newstate *Tokenprivileges, buflen uint32, prevstate *Tokenprivileges, returnlen *uint32) (err error) = advapi32.AdjustTokenPrivileges
//sys	AdjustTokenGroups(token Token, resetToDefault bool, newstate *Tokengroups, buflen uint32, prevstate *Tokengroups, returnlen *uint32) (err error) = advapi32.AdjustTokenGroups
//sys	GetTokenInformation(token Token, infoClass uint32, info *byte, infoLen uint32, returnedLen *uint32) (err error) = advapi32.GetTokenInformation
//sys	SetTokenInformation(token Token, infoClass uint32, info *byte, infoLen uint32) (err error) = advapi32.SetTokenInformation
//sys	DuplicateTokenEx(existingToken Token, desiredAccess uint32, tokenAttributes *SecurityAttributes, impersonationLevel uint32, tokenType uint32, newToken *Token) (err error) = advapi32.DuplicateTokenEx
//sys	GetUserProfileDirectory(t Token, dir *uint16, dirLen *uint32) (err error) = userenv.GetUserProfileDirectoryW
//sys	getSystemDirectory(dir *uint16, dirLen uint32) (len uint32, err error) = kernel32.GetSystemDirectoryW
//sys	getWindowsDirectory(dir *uint16, dirLen uint32) (len uint32, err error) = kernel32.GetWindowsDirectoryW
//sys	getSystemWindowsDirectory(dir *uint16, dirLen uint32) (len uint32, err error) = kernel32.GetSystemWindowsDirectoryW
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
// 当用户登录时，系统会创建一个访问令牌，且为该用户执行的每个进程都会拥有该令牌的一个副本。
// 令牌用于标识用户、用户所属的组以及用户的权限。系统使用令牌来控制对可保护对象的访问，并控制用户在本地计算机上执行各种系统相关操作的能力。
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
// 已弃用：请显式调用 OpenProcessToken(CurrentProcess(), ...) 并传入所需的访问权限，或者使用 GetCurrentProcessToken 获取一个用于 TOKEN_QUERY 的令牌。
# <翻译结束>


<原文开始>
// GetCurrentProcessToken returns the access token associated with
// the current process. It is a pseudo token that does not need
// to be closed.
<原文结束>

# <翻译开始>
// GetCurrentProcessToken 获取与当前进程关联的访问令牌。
// 它是一个无需关闭的伪令牌。
# <翻译结束>


<原文开始>
// GetCurrentThreadToken return the access token associated with
// the current thread. It is a pseudo token that does not need
// to be closed.
<原文结束>

# <翻译开始>
// GetCurrentThreadToken 返回与当前线程关联的访问令牌。它是一个无需关闭的伪令牌。
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
// Close 释放对访问令牌的访问。
# <翻译结束>


<原文开始>
// getInfo retrieves a specified type of information about an access token.
<原文结束>

# <翻译开始>
// getInfo 用于获取关于访问令牌的特定类型信息。
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
// 返回一个指向 SID 结构的指针，该结构代表一个组，当使用此访问令牌的进程创建任何对象时，该组将成为这些对象的主组。
# <翻译结束>


<原文开始>
// GetUserProfileDirectory retrieves path to the
// root directory of the access token t user's profile.
<原文结束>

# <翻译开始>
// GetUserProfileDirectory 获取访问令牌所属用户个人资料的根目录路径。
# <翻译结束>


<原文开始>
// IsElevated returns whether the current token is elevated from a UAC perspective.
<原文结束>

# <翻译开始>
// IsElevated 返回当前令牌是否从UAC角度来看已被提升。
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
// GetSystemDirectory 获取当前系统目录的路径，通常（但并不总是）为 `C:\Windows\System32`。
# <翻译结束>


<原文开始>
// GetWindowsDirectory retrieves the path to current location of the Windows
// directory, which is typically, though not always, `C:\Windows`. This may
// be a private user directory in the case that the application is running
// under a terminal server.
<原文结束>

# <翻译开始>
// GetWindowsDirectory 函数用于获取当前 Windows 目录的路径，该目录通常（但并非总是）为 `C:\Windows`。在应用程序通过终端服务器运行的情况下，这可能是一个私有用户目录。
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
// 如果err为EINVAL，此返回ERROR_SUCCESS，表示这是一个非受限令牌。
# <翻译结束>


<原文开始>
//sys WTSQueryUserToken(session uint32, token *Token) (err error) = wtsapi32.WTSQueryUserToken
//sys WTSEnumerateSessions(handle Handle, reserved uint32, version uint32, sessions **WTS_SESSION_INFO, count *uint32) (err error) = wtsapi32.WTSEnumerateSessionsW
//sys WTSFreeMemory(ptr uintptr) = wtsapi32.WTSFreeMemory
//sys WTSGetActiveConsoleSessionId() (sessionID uint32)
<原文结束>

# <翻译开始>
//sys WTSQueryUserToken(session uint32, token *Token) (err error) = wtsapi32.WTSQueryUserToken
// 系统调用：根据给定的会话ID（session uint32）获取用户令牌（token *Token），并返回可能的错误（err error）。该函数等同于wtsapi32库中的WTSQueryUserToken。
// 
//sys WTSEnumerateSessions(handle Handle, reserved uint32, version uint32, sessions **WTS_SESSION_INFO, count *uint32) (err error) = wtsapi32.WTSEnumerateSessionsW
// 系统调用：使用指定句柄（handle Handle）、保留值（reserved uint32）、版本号（version uint32）枚举远程桌面服务会话。结果存储在sessions **WTS_SESSION_INFO指针所指向的结构体数组中，并通过count *uint32返回会话数量。返回可能的错误（err error）。该函数等同于wtsapi32库中的WTSEnumerateSessionsW。
// 
//sys WTSFreeMemory(ptr uintptr) = wtsapi32.WTSFreeMemory
// 系统调用：释放由wtsapi32库函数分配的内存，传入待释放内存的地址（ptr uintptr）。该函数等同于wtsapi32库中的WTSFreeMemory。
// 
//sys WTSGetActiveConsoleSessionId() (sessionID uint32)
// 系统调用：获取当前活动控制台会话的ID（sessionID uint32）。
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
// 安全信息类型常量
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
// 用于 AceFlags 和 Inheritance 字段的常量
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
// TRUSTEE_TYPE的常量
# <翻译结束>


<原文开始>
// Constants for ObjectsPresent field
<原文结束>

# <翻译开始>
// 对ObjectsPresent字段的常量定义
# <翻译结束>


<原文开始>
// This type is the union inside of TRUSTEE and must be created using one of the TrusteeValueFrom* functions.
<原文结束>

# <翻译开始>
// 此类型是 TRUSTEE 内部的联合体，必须使用 TrusteeValueFrom* 系列函数之一进行创建。
# <翻译结束>


<原文开始>
//sys	getSecurityInfo(handle Handle, objectType SE_OBJECT_TYPE, securityInformation SECURITY_INFORMATION, owner **SID, group **SID, dacl **ACL, sacl **ACL, sd **SECURITY_DESCRIPTOR) (ret error) = advapi32.GetSecurityInfo
//sys	SetSecurityInfo(handle Handle, objectType SE_OBJECT_TYPE, securityInformation SECURITY_INFORMATION, owner *SID, group *SID, dacl *ACL, sacl *ACL) (ret error) = advapi32.SetSecurityInfo
//sys	getNamedSecurityInfo(objectName string, objectType SE_OBJECT_TYPE, securityInformation SECURITY_INFORMATION, owner **SID, group **SID, dacl **ACL, sacl **ACL, sd **SECURITY_DESCRIPTOR) (ret error) = advapi32.GetNamedSecurityInfoW
//sys	SetNamedSecurityInfo(objectName string, objectType SE_OBJECT_TYPE, securityInformation SECURITY_INFORMATION, owner *SID, group *SID, dacl *ACL, sacl *ACL) (ret error) = advapi32.SetNamedSecurityInfoW
//sys	SetKernelObjectSecurity(handle Handle, securityInformation SECURITY_INFORMATION, securityDescriptor *SECURITY_DESCRIPTOR) (err error) = advapi32.SetKernelObjectSecurity
<原文结束>

# <翻译开始>
//sys	获取安全信息(handle句柄, objectType SE_OBJECT_TYPE类型, securityInformation SECURITY_INFORMATION信息, owner **SID所有者指针, group **SID组指针, dacl **ACL自定义访问控制列表指针, sacl **ACL系统访问控制列表指针, sd **SECURITY_DESCRIPTOR安全描述符指针) (返回错误 ret) = advapi32.GetSecurityInfo
//sys	设置安全信息(handle句柄, objectType SE_OBJECT_TYPE类型, securityInformation SECURITY_INFORMATION信息, owner *SID所有者指针, group *SID组指针, dacl *ACL自定义访问控制列表指针, sacl *ACL系统访问控制列表指针) (返回错误 ret) = advapi32.SetSecurityInfo
//sys	获取命名安全信息(objectName字符串, objectType SE_OBJECT_TYPE类型, securityInformation SECURITY_INFORMATION信息, owner **SID所有者指针, group **SID组指针, dacl **ACL自定义访问控制列表指针, sacl **ACL系统访问控制列表指针, sd **SECURITY_DESCRIPTOR安全描述符指针) (返回错误 ret) = advapi32.GetNamedSecurityInfoW
//sys	设置命名安全信息(objectName字符串, objectType SE_OBJECT_TYPE类型, securityInformation SECURITY_INFORMATION信息, owner *SID所有者指针, group *SID组指针, dacl *ACL自定义访问控制列表指针, sacl *ACL系统访问控制列表指针) (返回错误 ret) = advapi32.SetNamedSecurityInfoW
//sys	设置内核对象安全(handle句柄, securityInformation SECURITY_INFORMATION信息, securityDescriptor *SECURITY_DESCRIPTOR安全描述符指针) (返回错误 err) = advapi32.SetKernelObjectSecurity
# <翻译结束>


<原文开始>
//sys	buildSecurityDescriptor(owner *TRUSTEE, group *TRUSTEE, countAccessEntries uint32, accessEntries *EXPLICIT_ACCESS, countAuditEntries uint32, auditEntries *EXPLICIT_ACCESS, oldSecurityDescriptor *SECURITY_DESCRIPTOR, sizeNewSecurityDescriptor *uint32, newSecurityDescriptor **SECURITY_DESCRIPTOR) (ret error) = advapi32.BuildSecurityDescriptorW
//sys	initializeSecurityDescriptor(absoluteSD *SECURITY_DESCRIPTOR, revision uint32) (err error) = advapi32.InitializeSecurityDescriptor
<原文结束>

# <翻译开始>
//sys	创建安全描述符(owner *TRUSTEE, group *TRUSTEE, 访问权限条目计数 uint32, 访问权限条目 *EXPLICIT_ACCESS, 审计条目计数 uint32, 审计条目 *EXPLICIT_ACCESS, 旧安全描述符 *SECURITY_DESCRIPTOR, 新安全描述符大小指针 *uint32, 新安全描述符指针 **SECURITY_DESCRIPTOR) (返回错误 ret error) = advapi32.BuildSecurityDescriptorW
//sys	初始化安全描述符(绝对安全描述符 *SECURITY_DESCRIPTOR, 版本号 uint32) (返回错误 err error) = advapi32.InitializeSecurityDescriptor
# <翻译结束>


<原文开始>
//sys	getSecurityDescriptorControl(sd *SECURITY_DESCRIPTOR, control *SECURITY_DESCRIPTOR_CONTROL, revision *uint32) (err error) = advapi32.GetSecurityDescriptorControl
//sys	getSecurityDescriptorDacl(sd *SECURITY_DESCRIPTOR, daclPresent *bool, dacl **ACL, daclDefaulted *bool) (err error) = advapi32.GetSecurityDescriptorDacl
//sys	getSecurityDescriptorSacl(sd *SECURITY_DESCRIPTOR, saclPresent *bool, sacl **ACL, saclDefaulted *bool) (err error) = advapi32.GetSecurityDescriptorSacl
//sys	getSecurityDescriptorOwner(sd *SECURITY_DESCRIPTOR, owner **SID, ownerDefaulted *bool) (err error) = advapi32.GetSecurityDescriptorOwner
//sys	getSecurityDescriptorGroup(sd *SECURITY_DESCRIPTOR, group **SID, groupDefaulted *bool) (err error) = advapi32.GetSecurityDescriptorGroup
//sys	getSecurityDescriptorLength(sd *SECURITY_DESCRIPTOR) (len uint32) = advapi32.GetSecurityDescriptorLength
//sys	getSecurityDescriptorRMControl(sd *SECURITY_DESCRIPTOR, rmControl *uint8) (ret error) [failretval!=0] = advapi32.GetSecurityDescriptorRMControl
//sys	isValidSecurityDescriptor(sd *SECURITY_DESCRIPTOR) (isValid bool) = advapi32.IsValidSecurityDescriptor
<原文结束>

# <翻译开始>
//sys	获取安全描述符控制信息（sd *SECURITY_DESCRIPTOR, 控制 *SECURITY_DESCRIPTOR_CONTROL, 修订版 *uint32）(错误 error) = advapi32.GetSecurityDescriptorControl
//sys	获取安全描述符的DACL（sd *SECURITY_DESCRIPTOR, daclPresent *bool, dacl **ACL, daclDefaulted *bool）(错误 error) = advapi32.GetSecurityDescriptorDacl
//sys	获取安全描述符的SACL（sd *SECURITY_DESCRIPTOR, saclPresent *bool, sacl **ACL, saclDefaulted *bool）(错误 error) = advapi32.GetSecurityDescriptorSacl
//sys	获取安全描述符的所有者（sd *SECURITY_DESCRIPTOR, 所有者 **SID, 所有者Defaulted *bool）(错误 error) = advapi32.GetSecurityDescriptorOwner
//sys	获取安全描述符的组（sd *SECURITY_DESCRIPTOR, 组 **SID, 组Defaulted *bool）(错误 error) = advapi32.GetSecurityDescriptorGroup
//sys	获取安全描述符的长度（sd *SECURITY_DESCRIPTOR）(长度 uint32) = advapi32.GetSecurityDescriptorLength
//sys	获取安全描述符的RM控制信息（sd *SECURITY_DESCRIPTOR, rmControl *uint8）(返回错误 error)[失败返回值!=0] = advapi32.GetSecurityDescriptorRMControl
//sys	验证安全描述符的有效性（sd *SECURITY_DESCRIPTOR）(有效 bool) = advapi32.IsValidSecurityDescriptor
# <翻译结束>


<原文开始>
//sys	setSecurityDescriptorControl(sd *SECURITY_DESCRIPTOR, controlBitsOfInterest SECURITY_DESCRIPTOR_CONTROL, controlBitsToSet SECURITY_DESCRIPTOR_CONTROL) (err error) = advapi32.SetSecurityDescriptorControl
//sys	setSecurityDescriptorDacl(sd *SECURITY_DESCRIPTOR, daclPresent bool, dacl *ACL, daclDefaulted bool) (err error) = advapi32.SetSecurityDescriptorDacl
//sys	setSecurityDescriptorSacl(sd *SECURITY_DESCRIPTOR, saclPresent bool, sacl *ACL, saclDefaulted bool) (err error) = advapi32.SetSecurityDescriptorSacl
//sys	setSecurityDescriptorOwner(sd *SECURITY_DESCRIPTOR, owner *SID, ownerDefaulted bool) (err error) = advapi32.SetSecurityDescriptorOwner
//sys	setSecurityDescriptorGroup(sd *SECURITY_DESCRIPTOR, group *SID, groupDefaulted bool) (err error) = advapi32.SetSecurityDescriptorGroup
//sys	setSecurityDescriptorRMControl(sd *SECURITY_DESCRIPTOR, rmControl *uint8) = advapi32.SetSecurityDescriptorRMControl
<原文结束>

# <翻译开始>
//sys	设置安全描述符控制信息(sd *SECURITY_DESCRIPTOR, controlBitsOfInterest SECURITY_DESCRIPTOR_CONTROL, controlBitsToSet SECURITY_DESCRIPTOR_CONTROL) (err error) = advapi32.SetSecurityDescriptorControl
//sys	设置安全描述符的DACL(sd *SECURITY_DESCRIPTOR, daclPresent bool, dacl *ACL, daclDefaulted bool) (err error) = advapi32.SetSecurityDescriptorDacl
//sys	设置安全描述符的SACL(sd *SECURITY_DESCRIPTOR, saclPresent bool, sacl *ACL, saclDefaulted bool) (err error) = advapi32.SetSecurityDescriptorSacl
//sys	设置安全描述符的所有者(sd *SECURITY_DESCRIPTOR, owner *SID, ownerDefaulted bool) (err error) = advapi32.SetSecurityDescriptorOwner
//sys	设置安全描述符的组(sd *SECURITY_DESCRIPTOR, group *SID, groupDefaulted bool) (err error) = advapi32.SetSecurityDescriptorGroup
//sys	设置安全描述符的RM控制信息(sd *SECURITY_DESCRIPTOR, rmControl *uint8) = advapi32.SetSecurityDescriptorRMControl
# <翻译结束>


<原文开始>
//sys	convertStringSecurityDescriptorToSecurityDescriptor(str string, revision uint32, sd **SECURITY_DESCRIPTOR, size *uint32) (err error) = advapi32.ConvertStringSecurityDescriptorToSecurityDescriptorW
//sys	convertSecurityDescriptorToStringSecurityDescriptor(sd *SECURITY_DESCRIPTOR, revision uint32, securityInformation SECURITY_INFORMATION, str **uint16, strLen *uint32) (err error) = advapi32.ConvertSecurityDescriptorToStringSecurityDescriptorW
<原文结束>

# <翻译开始>
//sys convertStringSecurityDescriptorToSecurityDescriptor(str string, revision uint32, sd **SECURITY_DESCRIPTOR, size *uint32) (err error) = advapi32.ConvertStringSecurityDescriptorToSecurityDescriptorW
// 系统调用：将字符串形式的安全描述符（str）转换为安全描述符结构（sd），指定修订版本（revision）。同时更新指向安全描述符的指针（**SECURITY_DESCRIPTOR）和其大小（*uint32）。该函数在advapi32库中名为ConvertStringSecurityDescriptorToSecurityDescriptorW，返回错误信息（err）。
// 
//sys convertSecurityDescriptorToStringSecurityDescriptor(sd *SECURITY_DESCRIPTOR, revision uint32, securityInformation SECURITY_INFORMATION, str **uint16, strLen *uint32) (err error) = advapi32.ConvertSecurityDescriptorToStringSecurityDescriptorW
// 系统调用：将安全描述符结构（sd）转换为字符串形式的安全描述符。输入参数包括安全描述符的修订版本（revision）、安全信息标志（securityInformation）。函数将返回指向转换后字符串的指针（**uint16）及其长度（*uint32）。该函数在advapi32库中名为ConvertSecurityDescriptorToStringSecurityDescriptorW，返回错误信息（err）。
# <翻译结束>


<原文开始>
//sys	makeAbsoluteSD(selfRelativeSD *SECURITY_DESCRIPTOR, absoluteSD *SECURITY_DESCRIPTOR, absoluteSDSize *uint32, dacl *ACL, daclSize *uint32, sacl *ACL, saclSize *uint32, owner *SID, ownerSize *uint32, group *SID, groupSize *uint32) (err error) = advapi32.MakeAbsoluteSD
//sys	makeSelfRelativeSD(absoluteSD *SECURITY_DESCRIPTOR, selfRelativeSD *SECURITY_DESCRIPTOR, selfRelativeSDSize *uint32) (err error) = advapi32.MakeSelfRelativeSD
<原文结束>

# <翻译开始>
//sys	makeAbsoluteSD(selfRelativeSD *SECURITY_DESCRIPTOR, absoluteSD *SECURITY_DESCRIPTOR, absoluteSDSize *uint32, dacl *ACL, daclSize *uint32, sacl *ACL, saclSize *uint32, owner *SID, ownerSize *uint32, group *SID, groupSize *uint32) (err error) = advapi32.MakeAbsoluteSD
// 
//sys	makeSelfRelativeSD(absoluteSD *SECURITY_DESCRIPTOR, selfRelativeSD *SECURITY_DESCRIPTOR, selfRelativeSDSize *uint32) (err error) = advapi32.MakeSelfRelativeSD
// 
// 注释翻译成中文如下：
// 
//sys	将自相关安全描述符（selfRelativeSD）转换为绝对安全描述符（absoluteSD），并分别计算并填充其大小（absoluteSDSize）、DACL（dacl）及其大小（daclSize）、SACL（sacl）及其大小（saclSize）、所有者（owner）及其大小（ownerSize）、组（group）及其大小（groupSize）。此操作可能返回错误，该函数对应于 advapi32.dll 中的 MakeAbsoluteSD 函数。
// 
//sys	将绝对安全描述符（absoluteSD）转换为自相关安全描述符（selfRelativeSD），并计算并填充其大小（selfRelativeSDSize）。此操作可能返回错误，该函数对应于 advapi32.dll 中的 MakeSelfRelativeSD 函数。
# <翻译结束>


<原文开始>
//sys	setEntriesInAcl(countExplicitEntries uint32, explicitEntries *EXPLICIT_ACCESS, oldACL *ACL, newACL **ACL) (ret error) = advapi32.SetEntriesInAclW
<原文结束>

# <翻译开始>
//sys	setEntriesInAcl(countExplicitEntries uint32, explicitEntries *EXPLICIT_ACCESS, oldACL *ACL, newACL **ACL) (ret error) = advapi32.SetEntriesInAclW
// 
// 系统调用 setEntriesInAcl，其参数如下：
//   countExplicitEntries: 类型为 uint32，表示显式访问条目数量。
//   explicitEntries: 类型为 *EXPLICIT_ACCESS，指向显式访问条目的指针。
//   oldACL: 类型为 *ACL，指向旧的访问控制列表（ACL）。
//   newACL: 类型为 **ACL，输出参数，用于接收新生成的访问控制列表的指针。
//
// 该函数返回一个错误（ret error），在成功时返回 nil，失败时返回具体错误信息。
//
// 此系统调用绑定到 advapi32.dll 中的 SetEntriesInAclW 函数。
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
// SetRMControl 设置安全描述符的资源管理器控制位
# <翻译结束>


<原文开始>
// DACL returns the security descriptor DACL and whether it was defaulted. The dacl return value may be nil
// if a DACL exists but is an "empty DACL", meaning fully permissive. If the DACL does not exist, err returns
// ERROR_OBJECT_NOT_FOUND.
<原文结束>

# <翻译开始>
// DACL 返回安全描述符的 DACL（Discretionary Access Control List）以及它是否为默认值。dacl 返回值可能为 nil，如果存在一个 DACL 但它是“空 DACL”，即表示完全无限制。如果不存在 DACL，则 err 返回 ERROR_OBJECT_NOT_FOUND。
# <翻译结束>


<原文开始>
// SetDACL sets the absolute security descriptor DACL.
<原文结束>

# <翻译开始>
// SetDACL 设置绝对安全描述符的DACL
# <翻译结束>


<原文开始>
// SACL returns the security descriptor SACL and whether it was defaulted. The sacl return value may be nil
// if a SACL exists but is an "empty SACL", meaning fully permissive. If the SACL does not exist, err returns
// ERROR_OBJECT_NOT_FOUND.
<原文结束>

# <翻译开始>
// SACL 返回安全描述符的 SACL（系统访问控制列表）以及它是否为默认值。若存在 SACL 但为“空 SACL”，即完全无限制，sacl 返回值可能为 nil。如果 SACL 不存在，err 返回 ERROR_OBJECT_NOT_FOUND。
# <翻译结束>


<原文开始>
// SetSACL sets the absolute security descriptor SACL.
<原文结束>

# <翻译开始>
// SetSACL 设置安全描述符的绝对SACL。
# <翻译结束>


<原文开始>
// Owner returns the security descriptor owner and whether it was defaulted.
<原文结束>

# <翻译开始>
// Owner 返回安全描述符的所有者以及它是否为默认值。
# <翻译结束>


<原文开始>
// SetOwner sets the absolute security descriptor owner.
<原文结束>

# <翻译开始>
// SetOwner 设置绝对安全描述符的所有者。
# <翻译结束>


<原文开始>
// Group returns the security descriptor group and whether it was defaulted.
<原文结束>

# <翻译开始>
// Group 返回安全描述符的组信息及其是否为默认值。
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
// IsValid 返回该安全描述符是否有效。
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
// ToSelfRelative 将一个绝对形式的安全描述符转换为自相对形式。
# <翻译结束>


<原文开始>
// makeSelfRelativeSD is expected to fail, but it succeeds.
<原文结束>

# <翻译开始>
// makeSelfRelativeSD 预计会失败，但它却成功了。
# <翻译结束>


<原文开始>
	// SECURITY_DESCRIPTOR has pointers in it, which means checkptr expects for it to
	// be aligned properly. When we're copying a Windows-allocated struct to a
	// Go-allocated one, make sure that the Go allocation is aligned to the
	// pointer size.
<原文结束>

# <翻译开始>
// SECURITY_DESCRIPTOR 结构中包含指针，这意味着 checkptr 预期其应正确对齐。当我们把一个由 Windows 分配的结构复制到一个由 Go 分配的结构时，确保 Go 的内存分配按指针大小对齐。
# <翻译结束>


<原文开始>
// SecurityDescriptorFromString converts an SDDL string describing a security descriptor into a
// self-relative security descriptor object allocated on the Go heap.
<原文结束>

# <翻译开始>
// SecurityDescriptorFromString 将描述安全描述符的 SDDL 字符串转换为在 Go 堆上分配的自相对安全描述符对象。
# <翻译结束>


<原文开始>
// GetSecurityInfo queries the security information for a given handle and returns the self-relative security
// descriptor result on the Go heap.
<原文结束>

# <翻译开始>
// GetSecurityInfo 通过给定的句柄查询安全信息，并在Go堆上返回自相关的安全描述符结果。
# <翻译结束>


<原文开始>
// GetNamedSecurityInfo queries the security information for a given named object and returns the self-relative security
// descriptor result on the Go heap.
<原文结束>

# <翻译开始>
// GetNamedSecurityInfo 为给定的命名对象查询安全信息，并在Go堆上返回自相关的安全描述符结果。
# <翻译结束>


<原文开始>
// BuildSecurityDescriptor makes a new security descriptor using the input trustees, explicit access lists, and
// prior security descriptor to be merged, any of which can be nil, returning the self-relative security descriptor
// result on the Go heap.
<原文结束>

# <翻译开始>
// BuildSecurityDescriptor 通过使用输入的信任主体、显式访问列表以及待合并的先前安全描述符（这些参数均可为 nil）构建一个新的安全描述符，并在 Go 堆上返回生成的自相对安全描述符结果。
# <翻译结束>


<原文开始>
// NewSecurityDescriptor creates and initializes a new absolute security descriptor.
<原文结束>

# <翻译开始>
// NewSecurityDescriptor 创建并初始化一个新的绝对安全描述符
# <翻译结束>


<原文开始>
// ACLFromEntries returns a new ACL on the Go heap containing a list of explicit entries as well as those of another ACL.
// Both explicitEntries and mergedACL are optional and can be nil.
<原文结束>

# <翻译开始>
// ACLFromEntries 在Go堆上返回一个新的ACL，其中包含一个显式条目列表以及另一个ACL的条目。
// explicitEntries 和 mergedACL 两者都是可选的，可以为nil。
# <翻译结束>

