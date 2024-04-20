// 版权所有 ? 2012 The Go Authors。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

package windows

import (
	"syscall"
	"unsafe"
)

const (
	NameUnknown          = 0
	NameFullyQualifiedDN = 1
	NameSamCompatible    = 2
	NameDisplay          = 3
	NameUniqueId         = 6
	NameCanonical        = 7
	NameUserPrincipal    = 8
	NameCanonicalEx      = 9
	NameServicePrincipal = 10
	NameDnsDomain        = 12
)

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

// TranslateAccountName 将目录服务对象名从一种格式转换为另一种格式

// ff:
// initSize:
// to:
// from:
// username:
func TranslateAccountName(username string, from, to uint32, initSize int) (string, error) {
	u, e := UTF16PtrFromString(username)
	if e != nil {
		return "", e
	}
	n := uint32(50)
	for {
		b := make([]uint16, n)
		e = TranslateName(u, from, to, &b[0], &n)
		if e == nil {
			return UTF16ToString(b[:n]), nil
		}
		if e != ERROR_INSUFFICIENT_BUFFER {
			return "", e
		}
		if n <= uint32(len(b)) {
			return "", e
		}
	}
}

const (
	// do not reorder
	NetSetupUnknownStatus = iota
	NetSetupUnjoined
	NetSetupWorkgroupName
	NetSetupDomainName
)

type UserInfo10 struct {
	Name       *uint16
	Comment    *uint16
	UsrComment *uint16
	FullName   *uint16
}

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

const (
	// do not reorder
	SidTypeUser = 1 + iota
	SidTypeGroup
	SidTypeDomain
	SidTypeAlias
	SidTypeWellKnownGroup
	SidTypeDeletedAccount
	SidTypeInvalid
	SidTypeUnknown
	SidTypeComputer
	SidTypeLabel
)

type SidIdentifierAuthority struct {
	Value [6]byte
}

var (
	SECURITY_NULL_SID_AUTHORITY        = SidIdentifierAuthority{[6]byte{0, 0, 0, 0, 0, 0}}
	SECURITY_WORLD_SID_AUTHORITY       = SidIdentifierAuthority{[6]byte{0, 0, 0, 0, 0, 1}}
	SECURITY_LOCAL_SID_AUTHORITY       = SidIdentifierAuthority{[6]byte{0, 0, 0, 0, 0, 2}}
	SECURITY_CREATOR_SID_AUTHORITY     = SidIdentifierAuthority{[6]byte{0, 0, 0, 0, 0, 3}}
	SECURITY_NON_UNIQUE_AUTHORITY      = SidIdentifierAuthority{[6]byte{0, 0, 0, 0, 0, 4}}
	SECURITY_NT_AUTHORITY              = SidIdentifierAuthority{[6]byte{0, 0, 0, 0, 0, 5}}
	SECURITY_MANDATORY_LABEL_AUTHORITY = SidIdentifierAuthority{[6]byte{0, 0, 0, 0, 0, 16}}
)

const (
	SECURITY_NULL_RID                   = 0
	SECURITY_WORLD_RID                  = 0
	SECURITY_LOCAL_RID                  = 0
	SECURITY_CREATOR_OWNER_RID          = 0
	SECURITY_CREATOR_GROUP_RID          = 1
	SECURITY_DIALUP_RID                 = 1
	SECURITY_NETWORK_RID                = 2
	SECURITY_BATCH_RID                  = 3
	SECURITY_INTERACTIVE_RID            = 4
	SECURITY_LOGON_IDS_RID              = 5
	SECURITY_SERVICE_RID                = 6
	SECURITY_LOCAL_SYSTEM_RID           = 18
	SECURITY_BUILTIN_DOMAIN_RID         = 32
	SECURITY_PRINCIPAL_SELF_RID         = 10
	SECURITY_CREATOR_OWNER_SERVER_RID   = 0x2
	SECURITY_CREATOR_GROUP_SERVER_RID   = 0x3
	SECURITY_LOGON_IDS_RID_COUNT        = 0x3
	SECURITY_ANONYMOUS_LOGON_RID        = 0x7
	SECURITY_PROXY_RID                  = 0x8
	SECURITY_ENTERPRISE_CONTROLLERS_RID = 0x9
	SECURITY_SERVER_LOGON_RID           = SECURITY_ENTERPRISE_CONTROLLERS_RID
	SECURITY_AUTHENTICATED_USER_RID     = 0xb
	SECURITY_RESTRICTED_CODE_RID        = 0xc
	SECURITY_NT_NON_UNIQUE_RID          = 0x15
)

// 预定义的与域相关的本地组RID。  
// 参见 https://msdn.microsoft.com/en-us/library/windows/desktop/aa379649(v=vs.85).aspx
const (
	DOMAIN_ALIAS_RID_ADMINS                         = 0x220
	DOMAIN_ALIAS_RID_USERS                          = 0x221
	DOMAIN_ALIAS_RID_GUESTS                         = 0x222
	DOMAIN_ALIAS_RID_POWER_USERS                    = 0x223
	DOMAIN_ALIAS_RID_ACCOUNT_OPS                    = 0x224
	DOMAIN_ALIAS_RID_SYSTEM_OPS                     = 0x225
	DOMAIN_ALIAS_RID_PRINT_OPS                      = 0x226
	DOMAIN_ALIAS_RID_BACKUP_OPS                     = 0x227
	DOMAIN_ALIAS_RID_REPLICATOR                     = 0x228
	DOMAIN_ALIAS_RID_RAS_SERVERS                    = 0x229
	DOMAIN_ALIAS_RID_PREW2KCOMPACCESS               = 0x22a
	DOMAIN_ALIAS_RID_REMOTE_DESKTOP_USERS           = 0x22b
	DOMAIN_ALIAS_RID_NETWORK_CONFIGURATION_OPS      = 0x22c
	DOMAIN_ALIAS_RID_INCOMING_FOREST_TRUST_BUILDERS = 0x22d
	DOMAIN_ALIAS_RID_MONITORING_USERS               = 0x22e
	DOMAIN_ALIAS_RID_LOGGING_USERS                  = 0x22f
	DOMAIN_ALIAS_RID_AUTHORIZATIONACCESS            = 0x230
	DOMAIN_ALIAS_RID_TS_LICENSE_SERVERS             = 0x231
	DOMAIN_ALIAS_RID_DCOM_USERS                     = 0x232
	DOMAIN_ALIAS_RID_IUSERS                         = 0x238
	DOMAIN_ALIAS_RID_CRYPTO_OPERATORS               = 0x239
	DOMAIN_ALIAS_RID_CACHEABLE_PRINCIPALS_GROUP     = 0x23b
	DOMAIN_ALIAS_RID_NON_CACHEABLE_PRINCIPALS_GROUP = 0x23c
	DOMAIN_ALIAS_RID_EVENT_LOG_READERS_GROUP        = 0x23d
	DOMAIN_ALIAS_RID_CERTSVC_DCOM_ACCESS_GROUP      = 0x23e
)

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

// 安全标识符（SID）结构是一种可变长度的结构，用于唯一地标识用户或组。
type SID struct{}

// StringToSid 将字符串格式的安全标识符（SID）转换为有效且可用的 SID。

// ff:
// s:
func StringToSid(s string) (*SID, error) {
	var sid *SID
	p, e := UTF16PtrFromString(s)
	if e != nil {
		return nil, e
	}
	e = ConvertStringSidToSid(p, &sid)
	if e != nil {
		return nil, e
	}
	defer LocalFree((Handle)(unsafe.Pointer(sid)))
	return sid.Copy()
}

// LookupSID 通过账户名获取该账户的安全标识符（SID）以及账户所在的域名。
// 参数System用于指定要在其中进行搜索的目标计算机。

// ff:
// err:
// accType:
// domain:
// sid:
// account:
// system:
func LookupSID(system, account string) (sid *SID, domain string, accType uint32, err error) {
	if len(account) == 0 {
		return nil, "", 0, syscall.EINVAL
	}
	acc, e := UTF16PtrFromString(account)
	if e != nil {
		return nil, "", 0, e
	}
	var sys *uint16
	if len(system) > 0 {
		sys, e = UTF16PtrFromString(system)
		if e != nil {
			return nil, "", 0, e
		}
	}
	n := uint32(50)
	dn := uint32(50)
	for {
		b := make([]byte, n)
		db := make([]uint16, dn)
		sid = (*SID)(unsafe.Pointer(&b[0]))
		e = LookupAccountName(sys, acc, sid, &n, &db[0], &dn, &accType)
		if e == nil {
			return sid, UTF16ToString(db), accType, nil
		}
		if e != ERROR_INSUFFICIENT_BUFFER {
			return nil, "", 0, e
		}
		if n <= uint32(len(b)) {
			return nil, "", 0, e
		}
	}
}

// String 将SID转换为适合显示、存储或传输的字符串格式。

// ff:
func (sid *SID) String() string {
	var s *uint16
	e := ConvertSidToStringSid(sid, &s)
	if e != nil {
		return ""
	}
	defer LocalFree((Handle)(unsafe.Pointer(s)))
	return UTF16ToString((*[256]uint16)(unsafe.Pointer(s))[:])
}

// Len 返回一个有效安全标识符（SID）的字节长度。

// ff:
func (sid *SID) Len() int {
	return int(GetLengthSid(sid))
}

// Copy 创建安全标识符 SID 的副本。

// ff:
func (sid *SID) Copy() (*SID, error) {
	b := make([]byte, sid.Len())
	sid2 := (*SID)(unsafe.Pointer(&b[0]))
	e := CopySid(uint32(len(b)), sid2, sid)
	if e != nil {
		return nil, e
	}
	return sid2, nil
}

// IdentifierAuthority 返回 SID 的标识符权限。

// ff:
func (sid *SID) IdentifierAuthority() SidIdentifierAuthority {
	return *getSidIdentifierAuthority(sid)
}

// SubAuthorityCount 返回SID中子授权机构的数量

// ff:
func (sid *SID) SubAuthorityCount() uint8 {
	return *getSidSubAuthorityCount(sid)
}

// SubAuthority 返回由指定索引（该索引必须小于 sid.SubAuthorityCount()）确定的 SID 的子权限部分

// ff:
// idx:
func (sid *SID) SubAuthority(idx uint32) uint32 {
	if idx >= uint32(sid.SubAuthorityCount()) {
		panic("sub-authority index out of range")
	}
	return *getSidSubAuthority(sid, idx)
}

// IsValid 判断SID是否具有有效的修订版号和长度。

// ff:
func (sid *SID) IsValid() bool {
	return isValidSid(sid)
}

// Equals 用于比较两个 SID 是否相等。

// ff:
// sid2:
func (sid *SID) Equals(sid2 *SID) bool {
	return EqualSid(sid, sid2)
}

// IsWellKnown 判断SID是否匹配已知的sidType。

// ff:
// sidType:
func (sid *SID) IsWellKnown(sidType WELL_KNOWN_SID_TYPE) bool {
	return isWellKnownSid(sid, sidType)
}

// LookupAccount 通过此 SID 查找账户名称以及该 SID 所在的第一个域的名称。
// 参数 System 指定要在其中搜索的目标计算机。

// ff:
// err:
// accType:
// domain:
// account:
// system:
func (sid *SID) LookupAccount(system string) (account, domain string, accType uint32, err error) {
	var sys *uint16
	if len(system) > 0 {
		sys, err = UTF16PtrFromString(system)
		if err != nil {
			return "", "", 0, err
		}
	}
	n := uint32(50)
	dn := uint32(50)
	for {
		b := make([]uint16, n)
		db := make([]uint16, dn)
		e := LookupAccountSid(sys, sid, &b[0], &n, &db[0], &dn, &accType)
		if e == nil {
			return UTF16ToString(b), UTF16ToString(db), accType, nil
		}
		if e != ERROR_INSUFFICIENT_BUFFER {
			return "", "", 0, e
		}
		if n <= uint32(len(b)) {
			return "", "", 0, e
		}
	}
}

// 各种类型的预指定SID，可在运行时进行合成与比较
type WELL_KNOWN_SID_TYPE uint32

const (
	WinNullSid                                    = 0
	WinWorldSid                                   = 1
	WinLocalSid                                   = 2
	WinCreatorOwnerSid                            = 3
	WinCreatorGroupSid                            = 4
	WinCreatorOwnerServerSid                      = 5
	WinCreatorGroupServerSid                      = 6
	WinNtAuthoritySid                             = 7
	WinDialupSid                                  = 8
	WinNetworkSid                                 = 9
	WinBatchSid                                   = 10
	WinInteractiveSid                             = 11
	WinServiceSid                                 = 12
	WinAnonymousSid                               = 13
	WinProxySid                                   = 14
	WinEnterpriseControllersSid                   = 15
	WinSelfSid                                    = 16
	WinAuthenticatedUserSid                       = 17
	WinRestrictedCodeSid                          = 18
	WinTerminalServerSid                          = 19
	WinRemoteLogonIdSid                           = 20
	WinLogonIdsSid                                = 21
	WinLocalSystemSid                             = 22
	WinLocalServiceSid                            = 23
	WinNetworkServiceSid                          = 24
	WinBuiltinDomainSid                           = 25
	WinBuiltinAdministratorsSid                   = 26
	WinBuiltinUsersSid                            = 27
	WinBuiltinGuestsSid                           = 28
	WinBuiltinPowerUsersSid                       = 29
	WinBuiltinAccountOperatorsSid                 = 30
	WinBuiltinSystemOperatorsSid                  = 31
	WinBuiltinPrintOperatorsSid                   = 32
	WinBuiltinBackupOperatorsSid                  = 33
	WinBuiltinReplicatorSid                       = 34
	WinBuiltinPreWindows2000CompatibleAccessSid   = 35
	WinBuiltinRemoteDesktopUsersSid               = 36
	WinBuiltinNetworkConfigurationOperatorsSid    = 37
	WinAccountAdministratorSid                    = 38
	WinAccountGuestSid                            = 39
	WinAccountKrbtgtSid                           = 40
	WinAccountDomainAdminsSid                     = 41
	WinAccountDomainUsersSid                      = 42
	WinAccountDomainGuestsSid                     = 43
	WinAccountComputersSid                        = 44
	WinAccountControllersSid                      = 45
	WinAccountCertAdminsSid                       = 46
	WinAccountSchemaAdminsSid                     = 47
	WinAccountEnterpriseAdminsSid                 = 48
	WinAccountPolicyAdminsSid                     = 49
	WinAccountRasAndIasServersSid                 = 50
	WinNTLMAuthenticationSid                      = 51
	WinDigestAuthenticationSid                    = 52
	WinSChannelAuthenticationSid                  = 53
	WinThisOrganizationSid                        = 54
	WinOtherOrganizationSid                       = 55
	WinBuiltinIncomingForestTrustBuildersSid      = 56
	WinBuiltinPerfMonitoringUsersSid              = 57
	WinBuiltinPerfLoggingUsersSid                 = 58
	WinBuiltinAuthorizationAccessSid              = 59
	WinBuiltinTerminalServerLicenseServersSid     = 60
	WinBuiltinDCOMUsersSid                        = 61
	WinBuiltinIUsersSid                           = 62
	WinIUserSid                                   = 63
	WinBuiltinCryptoOperatorsSid                  = 64
	WinUntrustedLabelSid                          = 65
	WinLowLabelSid                                = 66
	WinMediumLabelSid                             = 67
	WinHighLabelSid                               = 68
	WinSystemLabelSid                             = 69
	WinWriteRestrictedCodeSid                     = 70
	WinCreatorOwnerRightsSid                      = 71
	WinCacheablePrincipalsGroupSid                = 72
	WinNonCacheablePrincipalsGroupSid             = 73
	WinEnterpriseReadonlyControllersSid           = 74
	WinAccountReadonlyControllersSid              = 75
	WinBuiltinEventLogReadersGroup                = 76
	WinNewEnterpriseReadonlyControllersSid        = 77
	WinBuiltinCertSvcDComAccessGroup              = 78
	WinMediumPlusLabelSid                         = 79
	WinLocalLogonSid                              = 80
	WinConsoleLogonSid                            = 81
	WinThisOrganizationCertificateSid             = 82
	WinApplicationPackageAuthoritySid             = 83
	WinBuiltinAnyPackageSid                       = 84
	WinCapabilityInternetClientSid                = 85
	WinCapabilityInternetClientServerSid          = 86
	WinCapabilityPrivateNetworkClientServerSid    = 87
	WinCapabilityPicturesLibrarySid               = 88
	WinCapabilityVideosLibrarySid                 = 89
	WinCapabilityMusicLibrarySid                  = 90
	WinCapabilityDocumentsLibrarySid              = 91
	WinCapabilitySharedUserCertificatesSid        = 92
	WinCapabilityEnterpriseAuthenticationSid      = 93
	WinCapabilityRemovableStorageSid              = 94
	WinBuiltinRDSRemoteAccessServersSid           = 95
	WinBuiltinRDSEndpointServersSid               = 96
	WinBuiltinRDSManagementServersSid             = 97
	WinUserModeDriversSid                         = 98
	WinBuiltinHyperVAdminsSid                     = 99
	WinAccountCloneableControllersSid             = 100
	WinBuiltinAccessControlAssistanceOperatorsSid = 101
	WinBuiltinRemoteManagementUsersSid            = 102
	WinAuthenticationAuthorityAssertedSid         = 103
	WinAuthenticationServiceAssertedSid           = 104
	WinLocalAccountSid                            = 105
	WinLocalAccountAndAdministratorSid            = 106
	WinAccountProtectedUsersSid                   = 107
	WinCapabilityAppointmentsSid                  = 108
	WinCapabilityContactsSid                      = 109
	WinAccountDefaultSystemManagedSid             = 110
	WinBuiltinDefaultSystemManagedGroupSid        = 111
	WinBuiltinStorageReplicaAdminsSid             = 112
	WinAccountKeyAdminsSid                        = 113
	WinAccountEnterpriseKeyAdminsSid              = 114
	WinAuthenticationKeyTrustSid                  = 115
	WinAuthenticationKeyPropertyMFASid            = 116
	WinAuthenticationKeyPropertyAttestationSid    = 117
	WinAuthenticationFreshKeyAuthSid              = 118
	WinBuiltinDeviceOwnersSid                     = 119
)

// 为本地计算机上一个预先定义的知名别名创建一个SID（安全标识符），通常使用形式为Win*Sid的常量。

// ff:
// sidType:
func CreateWellKnownSid(sidType WELL_KNOWN_SID_TYPE) (*SID, error) {
	return CreateWellKnownDomainSid(sidType, nil)
}

// 为指定由 domainSid 参数所确定域内的某个预定义的知名别名创建一个 SID（安全标识符），通常使用形如 Win*Sid 的常量。

// ff:
// domainSid:
// sidType:
func CreateWellKnownDomainSid(sidType WELL_KNOWN_SID_TYPE, domainSid *SID) (*SID, error) {
	n := uint32(50)
	for {
		b := make([]byte, n)
		sid := (*SID)(unsafe.Pointer(&b[0]))
		err := createWellKnownSid(sidType, domainSid, sid, &n)
		if err == nil {
			return sid, nil
		}
		if err != ERROR_INSUFFICIENT_BUFFER {
			return nil, err
		}
		if n <= uint32(len(b)) {
			return nil, err
		}
	}
}

const (
	// do not reorder
	TOKEN_ASSIGN_PRIMARY = 1 << iota
	TOKEN_DUPLICATE
	TOKEN_IMPERSONATE
	TOKEN_QUERY
	TOKEN_QUERY_SOURCE
	TOKEN_ADJUST_PRIVILEGES
	TOKEN_ADJUST_GROUPS
	TOKEN_ADJUST_DEFAULT
	TOKEN_ADJUST_SESSIONID

	TOKEN_ALL_ACCESS = STANDARD_RIGHTS_REQUIRED |
		TOKEN_ASSIGN_PRIMARY |
		TOKEN_DUPLICATE |
		TOKEN_IMPERSONATE |
		TOKEN_QUERY |
		TOKEN_QUERY_SOURCE |
		TOKEN_ADJUST_PRIVILEGES |
		TOKEN_ADJUST_GROUPS |
		TOKEN_ADJUST_DEFAULT |
		TOKEN_ADJUST_SESSIONID
	TOKEN_READ  = STANDARD_RIGHTS_READ | TOKEN_QUERY
	TOKEN_WRITE = STANDARD_RIGHTS_WRITE |
		TOKEN_ADJUST_PRIVILEGES |
		TOKEN_ADJUST_GROUPS |
		TOKEN_ADJUST_DEFAULT
	TOKEN_EXECUTE = STANDARD_RIGHTS_EXECUTE
)

const (
	// do not reorder
	TokenUser = 1 + iota
	TokenGroups
	TokenPrivileges
	TokenOwner
	TokenPrimaryGroup
	TokenDefaultDacl
	TokenSource
	TokenType
	TokenImpersonationLevel
	TokenStatistics
	TokenRestrictedSids
	TokenSessionId
	TokenGroupsAndPrivileges
	TokenSessionReference
	TokenSandBoxInert
	TokenAuditPolicy
	TokenOrigin
	TokenElevationType
	TokenLinkedToken
	TokenElevation
	TokenHasRestrictions
	TokenAccessInformation
	TokenVirtualizationAllowed
	TokenVirtualizationEnabled
	TokenIntegrityLevel
	TokenUIAccess
	TokenMandatoryPolicy
	TokenLogonSid
	MaxTokenInfoClass
)

// 在Tokengroups.Groups[i].Attributes内部对属性进行分组
const (
	SE_GROUP_MANDATORY          = 0x00000001
	SE_GROUP_ENABLED_BY_DEFAULT = 0x00000002
	SE_GROUP_ENABLED            = 0x00000004
	SE_GROUP_OWNER              = 0x00000008
	SE_GROUP_USE_FOR_DENY_ONLY  = 0x00000010
	SE_GROUP_INTEGRITY          = 0x00000020
	SE_GROUP_INTEGRITY_ENABLED  = 0x00000040
	SE_GROUP_LOGON_ID           = 0xC0000000
	SE_GROUP_RESOURCE           = 0x20000000
	SE_GROUP_VALID_ATTRIBUTES   = SE_GROUP_MANDATORY | SE_GROUP_ENABLED_BY_DEFAULT | SE_GROUP_ENABLED | SE_GROUP_OWNER | SE_GROUP_USE_FOR_DENY_ONLY | SE_GROUP_LOGON_ID | SE_GROUP_RESOURCE | SE_GROUP_INTEGRITY | SE_GROUP_INTEGRITY_ENABLED
)

// Privilege attributes
const (
	SE_PRIVILEGE_ENABLED_BY_DEFAULT = 0x00000001
	SE_PRIVILEGE_ENABLED            = 0x00000002
	SE_PRIVILEGE_REMOVED            = 0x00000004
	SE_PRIVILEGE_USED_FOR_ACCESS    = 0x80000000
	SE_PRIVILEGE_VALID_ATTRIBUTES   = SE_PRIVILEGE_ENABLED_BY_DEFAULT | SE_PRIVILEGE_ENABLED | SE_PRIVILEGE_REMOVED | SE_PRIVILEGE_USED_FOR_ACCESS
)

// Token types
const (
	TokenPrimary       = 1
	TokenImpersonation = 2
)

// Impersonation levels
const (
	SecurityAnonymous      = 0
	SecurityIdentification = 1
	SecurityImpersonation  = 2
	SecurityDelegation     = 3
)

type LUID struct {
	LowPart  uint32
	HighPart int32
}

type LUIDAndAttributes struct {
	Luid       LUID
	Attributes uint32
}

type SIDAndAttributes struct {
	Sid        *SID
	Attributes uint32
}

type Tokenuser struct {
	User SIDAndAttributes
}

type Tokenprimarygroup struct {
	PrimaryGroup *SID
}

type Tokengroups struct {
	GroupCount uint32
	Groups     [1]SIDAndAttributes // 使用 AllGroups() 进行迭代。
}

// AllGroups 返回一个切片，可用于遍历 g 中的所有组。

// ff:
func (g *Tokengroups) AllGroups() []SIDAndAttributes {
	return (*[(1 << 28) - 1]SIDAndAttributes)(unsafe.Pointer(&g.Groups[0]))[:g.GroupCount:g.GroupCount]
}

type Tokenprivileges struct {
	PrivilegeCount uint32
	Privileges     [1]LUIDAndAttributes // 使用 AllPrivileges() 进行遍历。
}

// AllPrivileges 返回一个切片，可用于遍历p中的所有权限。

// ff:
func (p *Tokenprivileges) AllPrivileges() []LUIDAndAttributes {
	return (*[(1 << 27) - 1]LUIDAndAttributes)(unsafe.Pointer(&p.Privileges[0]))[:p.PrivilegeCount:p.PrivilegeCount]
}

type Tokenmandatorylabel struct {
	Label SIDAndAttributes
}


// ff:
func (tml *Tokenmandatorylabel) Size() uint32 {
	return uint32(unsafe.Sizeof(Tokenmandatorylabel{})) + GetLengthSid(tml.Label.Sid)
}

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

// 访问令牌包含登录会话的安全信息。
// 当用户登录时，系统会创建一个访问令牌，且为该用户执行的每个进程都会拥有该令牌的一个副本。
// 令牌用于标识用户、用户所属的组以及用户的权限。系统使用令牌来控制对可保护对象的访问，并控制用户在本地计算机上执行各种系统相关操作的能力。
type Token Handle

// OpenCurrentProcessToken 以 TOKEN_QUERY 访问权限打开与当前进程关联的访问令牌。这是一个需要被关闭的真实令牌。
//
// 已弃用：请显式调用 OpenProcessToken(CurrentProcess(), ...) 并传入所需的访问权限，或者使用 GetCurrentProcessToken 获取一个用于 TOKEN_QUERY 的令牌。

// ff:
// Token:
func OpenCurrentProcessToken() (Token, error) {
	var token Token
	err := OpenProcessToken(CurrentProcess(), TOKEN_QUERY, &token)
	return token, err
}

// GetCurrentProcessToken 获取与当前进程关联的访问令牌。
// 它是一个无需关闭的伪令牌。

// ff:
func GetCurrentProcessToken() Token {
	return Token(^uintptr(4 - 1))
}

// GetCurrentThreadToken 返回与当前线程关联的访问令牌。它是一个无需关闭的伪令牌。

// ff:
func GetCurrentThreadToken() Token {
	return Token(^uintptr(5 - 1))
}

// GetCurrentThreadEffectiveToken 返回与当前线程关联的有效访问令牌。
// 它是一个无需关闭的伪令牌。

// ff:
func GetCurrentThreadEffectiveToken() Token {
	return Token(^uintptr(6 - 1))
}

// Close 释放对访问令牌的访问。

// ff:
func (t Token) Close() error {
	return CloseHandle(Handle(t))
}

// getInfo 用于获取关于访问令牌的特定类型信息。
func (t Token) getInfo(class uint32, initSize int) (unsafe.Pointer, error) {
	n := uint32(initSize)
	for {
		b := make([]byte, n)
		e := GetTokenInformation(t, class, &b[0], uint32(len(b)), &n)
		if e == nil {
			return unsafe.Pointer(&b[0]), nil
		}
		if e != ERROR_INSUFFICIENT_BUFFER {
			return nil, e
		}
		if n <= uint32(len(b)) {
			return nil, e
		}
	}
}

// GetTokenUser 用于获取访问令牌对应的用户账户信息

// ff:
func (t Token) GetTokenUser() (*Tokenuser, error) {
	i, e := t.getInfo(TokenUser, 50)
	if e != nil {
		return nil, e
	}
	return (*Tokenuser)(i), nil
}

// GetTokenGroups 获取与访问令牌 t 关联的组账户。

// ff:
func (t Token) GetTokenGroups() (*Tokengroups, error) {
	i, e := t.getInfo(TokenGroups, 50)
	if e != nil {
		return nil, e
	}
	return (*Tokengroups)(i), nil
}

// GetTokenPrimaryGroup 获取访问令牌 t 的主组信息。
// 返回一个指向 SID 结构的指针，该结构代表一个组，当使用此访问令牌的进程创建任何对象时，该组将成为这些对象的主组。

// ff:
func (t Token) GetTokenPrimaryGroup() (*Tokenprimarygroup, error) {
	i, e := t.getInfo(TokenPrimaryGroup, 50)
	if e != nil {
		return nil, e
	}
	return (*Tokenprimarygroup)(i), nil
}

// GetUserProfileDirectory 获取访问令牌所属用户个人资料的根目录路径。

// ff:
func (t Token) GetUserProfileDirectory() (string, error) {
	n := uint32(100)
	for {
		b := make([]uint16, n)
		e := GetUserProfileDirectory(t, &b[0], &n)
		if e == nil {
			return UTF16ToString(b), nil
		}
		if e != ERROR_INSUFFICIENT_BUFFER {
			return "", e
		}
		if n <= uint32(len(b)) {
			return "", e
		}
	}
}

// IsElevated 返回当前令牌是否从UAC角度来看已被提升。

// ff:
func (token Token) IsElevated() bool {
	var isElevated uint32
	var outLen uint32
	err := GetTokenInformation(token, TokenElevation, (*byte)(unsafe.Pointer(&isElevated)), uint32(unsafe.Sizeof(isElevated)), &outLen)
	if err != nil {
		return false
	}
	return outLen == uint32(unsafe.Sizeof(isElevated)) && isElevated != 0
}

// 获取关联的令牌，该令牌可能是一个提升的UAC令牌。

// ff:
// Token:
func (token Token) GetLinkedToken() (Token, error) {
	var linkedToken Token
	var outLen uint32
	err := GetTokenInformation(token, TokenLinkedToken, (*byte)(unsafe.Pointer(&linkedToken)), uint32(unsafe.Sizeof(linkedToken)), &outLen)
	if err != nil {
		return Token(0), err
	}
	return linkedToken, nil
}

// GetSystemDirectory 获取当前系统目录的路径，通常（但并不总是）为 `C:\Windows\System32`。

// ff:
func GetSystemDirectory() (string, error) {
	n := uint32(MAX_PATH)
	for {
		b := make([]uint16, n)
		l, e := getSystemDirectory(&b[0], n)
		if e != nil {
			return "", e
		}
		if l <= n {
			return UTF16ToString(b[:l]), nil
		}
		n = l
	}
}

// GetWindowsDirectory 函数用于获取当前 Windows 目录的路径，该目录通常（但并非总是）为 `C:\Windows`。在应用程序通过终端服务器运行的情况下，这可能是一个私有用户目录。

// ff:
func GetWindowsDirectory() (string, error) {
	n := uint32(MAX_PATH)
	for {
		b := make([]uint16, n)
		l, e := getWindowsDirectory(&b[0], n)
		if e != nil {
			return "", e
		}
		if l <= n {
			return UTF16ToString(b[:l]), nil
		}
		n = l
	}
}

// GetSystemWindowsDirectory 获取当前 Windows 目录的路径，该目录通常（但并不总是）位于 `C:\Windows`。

// ff:
func GetSystemWindowsDirectory() (string, error) {
	n := uint32(MAX_PATH)
	for {
		b := make([]uint16, n)
		l, e := getSystemWindowsDirectory(&b[0], n)
		if e != nil {
			return "", e
		}
		if l <= n {
			return UTF16ToString(b[:l]), nil
		}
		n = l
	}
}

// IsMember 判断访问令牌 t 是否为给定 SID 的成员。

// ff:
// sid:
func (t Token) IsMember(sid *SID) (bool, error) {
	var b int32
	if e := checkTokenMembership(t, sid, &b); e != nil {
		return false, e
	}
	return b != 0, nil
}

// IsRestricted 判断访问令牌 t 是否为受限令牌。

// ff:
// err:
// isRestricted:
func (t Token) IsRestricted() (isRestricted bool, err error) {
	isRestricted, err = isTokenRestricted(t)
	if !isRestricted && err == syscall.EINVAL {
		// 如果err为EINVAL，此返回ERROR_SUCCESS，表示这是一个非受限令牌。
		err = nil
	}
	return
}

const (
	WTS_CONSOLE_CONNECT        = 0x1
	WTS_CONSOLE_DISCONNECT     = 0x2
	WTS_REMOTE_CONNECT         = 0x3
	WTS_REMOTE_DISCONNECT      = 0x4
	WTS_SESSION_LOGON          = 0x5
	WTS_SESSION_LOGOFF         = 0x6
	WTS_SESSION_LOCK           = 0x7
	WTS_SESSION_UNLOCK         = 0x8
	WTS_SESSION_REMOTE_CONTROL = 0x9
	WTS_SESSION_CREATE         = 0xa
	WTS_SESSION_TERMINATE      = 0xb
)

const (
	WTSActive       = 0
	WTSConnected    = 1
	WTSConnectQuery = 2
	WTSShadow       = 3
	WTSDisconnected = 4
	WTSIdle         = 5
	WTSListen       = 6
	WTSReset        = 7
	WTSDown         = 8
	WTSInit         = 9
)

type WTSSESSION_NOTIFICATION struct {
	Size      uint32
	SessionID uint32
}

type WTS_SESSION_INFO struct {
	SessionID         uint32
	WindowStationName *uint16
	State             uint32
}

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

type ACL struct {
	aclRevision byte
	sbz1        byte
	aclSize     uint16
	aceCount    uint16
	sbz2        uint16
}

type SECURITY_DESCRIPTOR struct {
	revision byte
	sbz1     byte
	control  SECURITY_DESCRIPTOR_CONTROL
	owner    *SID
	group    *SID
	sacl     *ACL
	dacl     *ACL
}

type SECURITY_QUALITY_OF_SERVICE struct {
	Length              uint32
	ImpersonationLevel  uint32
	ContextTrackingMode byte
	EffectiveOnly       byte
}

// ContextTrackingMode字段的SECURITY_QUALITY_OF_SERVICE常量
const (
	SECURITY_STATIC_TRACKING  = 0
	SECURITY_DYNAMIC_TRACKING = 1
)

type SecurityAttributes struct {
	Length             uint32
	SecurityDescriptor *SECURITY_DESCRIPTOR
	InheritHandle      uint32
}

type SE_OBJECT_TYPE uint32

// SE_OBJECT_TYPE类型的常量
const (
	SE_UNKNOWN_OBJECT_TYPE     = 0
	SE_FILE_OBJECT             = 1
	SE_SERVICE                 = 2
	SE_PRINTER                 = 3
	SE_REGISTRY_KEY            = 4
	SE_LMSHARE                 = 5
	SE_KERNEL_OBJECT           = 6
	SE_WINDOW_OBJECT           = 7
	SE_DS_OBJECT               = 8
	SE_DS_OBJECT_ALL           = 9
	SE_PROVIDER_DEFINED_OBJECT = 10
	SE_WMIGUID_OBJECT          = 11
	SE_REGISTRY_WOW64_32KEY    = 12
	SE_REGISTRY_WOW64_64KEY    = 13
)

type SECURITY_INFORMATION uint32

// 安全信息类型常量
const (
	OWNER_SECURITY_INFORMATION            = 0x00000001
	GROUP_SECURITY_INFORMATION            = 0x00000002
	DACL_SECURITY_INFORMATION             = 0x00000004
	SACL_SECURITY_INFORMATION             = 0x00000008
	LABEL_SECURITY_INFORMATION            = 0x00000010
	ATTRIBUTE_SECURITY_INFORMATION        = 0x00000020
	SCOPE_SECURITY_INFORMATION            = 0x00000040
	BACKUP_SECURITY_INFORMATION           = 0x00010000
	PROTECTED_DACL_SECURITY_INFORMATION   = 0x80000000
	PROTECTED_SACL_SECURITY_INFORMATION   = 0x40000000
	UNPROTECTED_DACL_SECURITY_INFORMATION = 0x20000000
	UNPROTECTED_SACL_SECURITY_INFORMATION = 0x10000000
)

type SECURITY_DESCRIPTOR_CONTROL uint16

// 定义类型SECURITY_DESCRIPTOR_CONTROL的常量
const (
	SE_OWNER_DEFAULTED       = 0x0001
	SE_GROUP_DEFAULTED       = 0x0002
	SE_DACL_PRESENT          = 0x0004
	SE_DACL_DEFAULTED        = 0x0008
	SE_SACL_PRESENT          = 0x0010
	SE_SACL_DEFAULTED        = 0x0020
	SE_DACL_AUTO_INHERIT_REQ = 0x0100
	SE_SACL_AUTO_INHERIT_REQ = 0x0200
	SE_DACL_AUTO_INHERITED   = 0x0400
	SE_SACL_AUTO_INHERITED   = 0x0800
	SE_DACL_PROTECTED        = 0x1000
	SE_SACL_PROTECTED        = 0x2000
	SE_RM_CONTROL_VALID      = 0x4000
	SE_SELF_RELATIVE         = 0x8000
)

type ACCESS_MASK uint32

// ACCESS_MASK类型的常量
const (
	DELETE                   = 0x00010000
	READ_CONTROL             = 0x00020000
	WRITE_DAC                = 0x00040000
	WRITE_OWNER              = 0x00080000
	SYNCHRONIZE              = 0x00100000
	STANDARD_RIGHTS_REQUIRED = 0x000F0000
	STANDARD_RIGHTS_READ     = READ_CONTROL
	STANDARD_RIGHTS_WRITE    = READ_CONTROL
	STANDARD_RIGHTS_EXECUTE  = READ_CONTROL
	STANDARD_RIGHTS_ALL      = 0x001F0000
	SPECIFIC_RIGHTS_ALL      = 0x0000FFFF
	ACCESS_SYSTEM_SECURITY   = 0x01000000
	MAXIMUM_ALLOWED          = 0x02000000
	GENERIC_READ             = 0x80000000
	GENERIC_WRITE            = 0x40000000
	GENERIC_EXECUTE          = 0x20000000
	GENERIC_ALL              = 0x10000000
)

type ACCESS_MODE uint32

// ACCESS_MODE类型的常量
const (
	NOT_USED_ACCESS   = 0
	GRANT_ACCESS      = 1
	SET_ACCESS        = 2
	DENY_ACCESS       = 3
	REVOKE_ACCESS     = 4
	SET_AUDIT_SUCCESS = 5
	SET_AUDIT_FAILURE = 6
)

// 用于 AceFlags 和 Inheritance 字段的常量
const (
	NO_INHERITANCE                     = 0x0
	SUB_OBJECTS_ONLY_INHERIT           = 0x1
	SUB_CONTAINERS_ONLY_INHERIT        = 0x2
	SUB_CONTAINERS_AND_OBJECTS_INHERIT = 0x3
	INHERIT_NO_PROPAGATE               = 0x4
	INHERIT_ONLY                       = 0x8
	INHERITED_ACCESS_ENTRY             = 0x10
	INHERITED_PARENT                   = 0x10000000
	INHERITED_GRANDPARENT              = 0x20000000
	OBJECT_INHERIT_ACE                 = 0x1
	CONTAINER_INHERIT_ACE              = 0x2
	NO_PROPAGATE_INHERIT_ACE           = 0x4
	INHERIT_ONLY_ACE                   = 0x8
	INHERITED_ACE                      = 0x10
	VALID_INHERIT_FLAGS                = 0x1F
)

type MULTIPLE_TRUSTEE_OPERATION uint32

// MULTIPLE_TRUSTEE_OPERATION相关的常量
const (
	NO_MULTIPLE_TRUSTEE    = 0
	TRUSTEE_IS_IMPERSONATE = 1
)

type TRUSTEE_FORM uint32

// TRUSTEE_FORM的常量
const (
	TRUSTEE_IS_SID              = 0
	TRUSTEE_IS_NAME             = 1
	TRUSTEE_BAD_FORM            = 2
	TRUSTEE_IS_OBJECTS_AND_SID  = 3
	TRUSTEE_IS_OBJECTS_AND_NAME = 4
)

type TRUSTEE_TYPE uint32

// TRUSTEE_TYPE的常量
const (
	TRUSTEE_IS_UNKNOWN          = 0
	TRUSTEE_IS_USER             = 1
	TRUSTEE_IS_GROUP            = 2
	TRUSTEE_IS_DOMAIN           = 3
	TRUSTEE_IS_ALIAS            = 4
	TRUSTEE_IS_WELL_KNOWN_GROUP = 5
	TRUSTEE_IS_DELETED          = 6
	TRUSTEE_IS_INVALID          = 7
	TRUSTEE_IS_COMPUTER         = 8
)

// 对ObjectsPresent字段的常量定义
const (
	ACE_OBJECT_TYPE_PRESENT           = 0x1
	ACE_INHERITED_OBJECT_TYPE_PRESENT = 0x2
)

type EXPLICIT_ACCESS struct {
	AccessPermissions ACCESS_MASK
	AccessMode        ACCESS_MODE
	Inheritance       uint32
	Trustee           TRUSTEE
}

// 此类型是 TRUSTEE 内部的联合体，必须使用 TrusteeValueFrom* 系列函数之一进行创建。
type TrusteeValue uintptr


// ff:
// str:
func TrusteeValueFromString(str string) TrusteeValue {
	return TrusteeValue(unsafe.Pointer(StringToUTF16Ptr(str)))
}

// ff:
// sid:
func TrusteeValueFromSID(sid *SID) TrusteeValue {
	return TrusteeValue(unsafe.Pointer(sid))
}

// ff:
// objectsAndSid:
func TrusteeValueFromObjectsAndSid(objectsAndSid *OBJECTS_AND_SID) TrusteeValue {
	return TrusteeValue(unsafe.Pointer(objectsAndSid))
}

// ff:
// objectsAndName:
func TrusteeValueFromObjectsAndName(objectsAndName *OBJECTS_AND_NAME) TrusteeValue {
	return TrusteeValue(unsafe.Pointer(objectsAndName))
}

type TRUSTEE struct {
	MultipleTrustee          *TRUSTEE
	MultipleTrusteeOperation MULTIPLE_TRUSTEE_OPERATION
	TrusteeForm              TRUSTEE_FORM
	TrusteeType              TRUSTEE_TYPE
	TrusteeValue             TrusteeValue
}

type OBJECTS_AND_SID struct {
	ObjectsPresent          uint32
	ObjectTypeGuid          GUID
	InheritedObjectTypeGuid GUID
	Sid                     *SID
}

type OBJECTS_AND_NAME struct {
	ObjectsPresent          uint32
	ObjectType              SE_OBJECT_TYPE
	ObjectTypeName          *uint16
	InheritedObjectTypeName *uint16
	Name                    *uint16
}

//sys	获取安全信息(handle句柄, objectType SE_OBJECT_TYPE类型, securityInformation SECURITY_INFORMATION信息, owner **SID所有者指针, group **SID组指针, dacl **ACL自定义访问控制列表指针, sacl **ACL系统访问控制列表指针, sd **SECURITY_DESCRIPTOR安全描述符指针) (返回错误 ret) = advapi32.GetSecurityInfo
//sys	设置安全信息(handle句柄, objectType SE_OBJECT_TYPE类型, securityInformation SECURITY_INFORMATION信息, owner *SID所有者指针, group *SID组指针, dacl *ACL自定义访问控制列表指针, sacl *ACL系统访问控制列表指针) (返回错误 ret) = advapi32.SetSecurityInfo
//sys	获取命名安全信息(objectName字符串, objectType SE_OBJECT_TYPE类型, securityInformation SECURITY_INFORMATION信息, owner **SID所有者指针, group **SID组指针, dacl **ACL自定义访问控制列表指针, sacl **ACL系统访问控制列表指针, sd **SECURITY_DESCRIPTOR安全描述符指针) (返回错误 ret) = advapi32.GetNamedSecurityInfoW
//sys	设置命名安全信息(objectName字符串, objectType SE_OBJECT_TYPE类型, securityInformation SECURITY_INFORMATION信息, owner *SID所有者指针, group *SID组指针, dacl *ACL自定义访问控制列表指针, sacl *ACL系统访问控制列表指针) (返回错误 ret) = advapi32.SetNamedSecurityInfoW
//sys	设置内核对象安全(handle句柄, securityInformation SECURITY_INFORMATION信息, securityDescriptor *SECURITY_DESCRIPTOR安全描述符指针) (返回错误 err) = advapi32.SetKernelObjectSecurity

//sys	创建安全描述符(owner *TRUSTEE, group *TRUSTEE, 访问权限条目计数 uint32, 访问权限条目 *EXPLICIT_ACCESS, 审计条目计数 uint32, 审计条目 *EXPLICIT_ACCESS, 旧安全描述符 *SECURITY_DESCRIPTOR, 新安全描述符大小指针 *uint32, 新安全描述符指针 **SECURITY_DESCRIPTOR) (返回错误 ret error) = advapi32.BuildSecurityDescriptorW
//sys	初始化安全描述符(绝对安全描述符 *SECURITY_DESCRIPTOR, 版本号 uint32) (返回错误 err error) = advapi32.InitializeSecurityDescriptor

//sys	获取安全描述符控制信息（sd *SECURITY_DESCRIPTOR, 控制 *SECURITY_DESCRIPTOR_CONTROL, 修订版 *uint32）(错误 error) = advapi32.GetSecurityDescriptorControl
//sys	获取安全描述符的DACL（sd *SECURITY_DESCRIPTOR, daclPresent *bool, dacl **ACL, daclDefaulted *bool）(错误 error) = advapi32.GetSecurityDescriptorDacl
//sys	获取安全描述符的SACL（sd *SECURITY_DESCRIPTOR, saclPresent *bool, sacl **ACL, saclDefaulted *bool）(错误 error) = advapi32.GetSecurityDescriptorSacl
//sys	获取安全描述符的所有者（sd *SECURITY_DESCRIPTOR, 所有者 **SID, 所有者Defaulted *bool）(错误 error) = advapi32.GetSecurityDescriptorOwner
//sys	获取安全描述符的组（sd *SECURITY_DESCRIPTOR, 组 **SID, 组Defaulted *bool）(错误 error) = advapi32.GetSecurityDescriptorGroup
//sys	获取安全描述符的长度（sd *SECURITY_DESCRIPTOR）(长度 uint32) = advapi32.GetSecurityDescriptorLength
//sys	获取安全描述符的RM控制信息（sd *SECURITY_DESCRIPTOR, rmControl *uint8）(返回错误 error)[失败返回值!=0] = advapi32.GetSecurityDescriptorRMControl
//sys	验证安全描述符的有效性（sd *SECURITY_DESCRIPTOR）(有效 bool) = advapi32.IsValidSecurityDescriptor

//sys	设置安全描述符控制信息(sd *SECURITY_DESCRIPTOR, controlBitsOfInterest SECURITY_DESCRIPTOR_CONTROL, controlBitsToSet SECURITY_DESCRIPTOR_CONTROL) (err error) = advapi32.SetSecurityDescriptorControl
//sys	设置安全描述符的DACL(sd *SECURITY_DESCRIPTOR, daclPresent bool, dacl *ACL, daclDefaulted bool) (err error) = advapi32.SetSecurityDescriptorDacl
//sys	设置安全描述符的SACL(sd *SECURITY_DESCRIPTOR, saclPresent bool, sacl *ACL, saclDefaulted bool) (err error) = advapi32.SetSecurityDescriptorSacl
//sys	设置安全描述符的所有者(sd *SECURITY_DESCRIPTOR, owner *SID, ownerDefaulted bool) (err error) = advapi32.SetSecurityDescriptorOwner
//sys	设置安全描述符的组(sd *SECURITY_DESCRIPTOR, group *SID, groupDefaulted bool) (err error) = advapi32.SetSecurityDescriptorGroup
//sys	设置安全描述符的RM控制信息(sd *SECURITY_DESCRIPTOR, rmControl *uint8) = advapi32.SetSecurityDescriptorRMControl

//sys convertStringSecurityDescriptorToSecurityDescriptor(str string, revision uint32, sd **SECURITY_DESCRIPTOR, size *uint32) (err error) = advapi32.ConvertStringSecurityDescriptorToSecurityDescriptorW
// 系统调用：将字符串形式的安全描述符（str）转换为安全描述符结构（sd），指定修订版本（revision）。同时更新指向安全描述符的指针（**SECURITY_DESCRIPTOR）和其大小（*uint32）。该函数在advapi32库中名为ConvertStringSecurityDescriptorToSecurityDescriptorW，返回错误信息（err）。
// 
//sys convertSecurityDescriptorToStringSecurityDescriptor(sd *SECURITY_DESCRIPTOR, revision uint32, securityInformation SECURITY_INFORMATION, str **uint16, strLen *uint32) (err error) = advapi32.ConvertSecurityDescriptorToStringSecurityDescriptorW
// 系统调用：将安全描述符结构（sd）转换为字符串形式的安全描述符。输入参数包括安全描述符的修订版本（revision）、安全信息标志（securityInformation）。函数将返回指向转换后字符串的指针（**uint16）及其长度（*uint32）。该函数在advapi32库中名为ConvertSecurityDescriptorToStringSecurityDescriptorW，返回错误信息（err）。

//sys	makeAbsoluteSD(selfRelativeSD *SECURITY_DESCRIPTOR, absoluteSD *SECURITY_DESCRIPTOR, absoluteSDSize *uint32, dacl *ACL, daclSize *uint32, sacl *ACL, saclSize *uint32, owner *SID, ownerSize *uint32, group *SID, groupSize *uint32) (err error) = advapi32.MakeAbsoluteSD
// 
//sys	makeSelfRelativeSD(absoluteSD *SECURITY_DESCRIPTOR, selfRelativeSD *SECURITY_DESCRIPTOR, selfRelativeSDSize *uint32) (err error) = advapi32.MakeSelfRelativeSD
// 
// 注释翻译成中文如下：
// 
//sys	将自相关安全描述符（selfRelativeSD）转换为绝对安全描述符（absoluteSD），并分别计算并填充其大小（absoluteSDSize）、DACL（dacl）及其大小（daclSize）、SACL（sacl）及其大小（saclSize）、所有者（owner）及其大小（ownerSize）、组（group）及其大小（groupSize）。此操作可能返回错误，该函数对应于 advapi32.dll 中的 MakeAbsoluteSD 函数。
// 
//sys	将绝对安全描述符（absoluteSD）转换为自相关安全描述符（selfRelativeSD），并计算并填充其大小（selfRelativeSDSize）。此操作可能返回错误，该函数对应于 advapi32.dll 中的 MakeSelfRelativeSD 函数。

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

// Control 返回安全描述符控制位。

// ff:
// err:
// revision:
// control:
func (sd *SECURITY_DESCRIPTOR) Control() (control SECURITY_DESCRIPTOR_CONTROL, revision uint32, err error) {
	err = getSecurityDescriptorControl(sd, &control, &revision)
	return
}

// SetControl 设置安全描述符控制位

// ff:
// controlBitsToSet:
// controlBitsOfInterest:
func (sd *SECURITY_DESCRIPTOR) SetControl(controlBitsOfInterest SECURITY_DESCRIPTOR_CONTROL, controlBitsToSet SECURITY_DESCRIPTOR_CONTROL) error {
	return setSecurityDescriptorControl(sd, controlBitsOfInterest, controlBitsToSet)
}

// RMControl 返回安全描述符资源管理器控制位。

// ff:
// err:
// control:
func (sd *SECURITY_DESCRIPTOR) RMControl() (control uint8, err error) {
	err = getSecurityDescriptorRMControl(sd, &control)
	return
}

// SetRMControl 设置安全描述符的资源管理器控制位

// ff:
// rmControl:
func (sd *SECURITY_DESCRIPTOR) SetRMControl(rmControl uint8) {
	setSecurityDescriptorRMControl(sd, &rmControl)
}

// DACL 返回安全描述符的 DACL（Discretionary Access Control List）以及它是否为默认值。dacl 返回值可能为 nil，如果存在一个 DACL 但它是“空 DACL”，即表示完全无限制。如果不存在 DACL，则 err 返回 ERROR_OBJECT_NOT_FOUND。

// ff:
// err:
// defaulted:
// dacl:
func (sd *SECURITY_DESCRIPTOR) DACL() (dacl *ACL, defaulted bool, err error) {
	var present bool
	err = getSecurityDescriptorDacl(sd, &present, &dacl, &defaulted)
	if !present {
		err = ERROR_OBJECT_NOT_FOUND
	}
	return
}

// SetDACL 设置绝对安全描述符的DACL

// ff:
// defaulted:
// present:
// dacl:
func (absoluteSD *SECURITY_DESCRIPTOR) SetDACL(dacl *ACL, present, defaulted bool) error {
	return setSecurityDescriptorDacl(absoluteSD, present, dacl, defaulted)
}

// SACL 返回安全描述符的 SACL（系统访问控制列表）以及它是否为默认值。若存在 SACL 但为“空 SACL”，即完全无限制，sacl 返回值可能为 nil。如果 SACL 不存在，err 返回 ERROR_OBJECT_NOT_FOUND。

// ff:
// err:
// defaulted:
// sacl:
func (sd *SECURITY_DESCRIPTOR) SACL() (sacl *ACL, defaulted bool, err error) {
	var present bool
	err = getSecurityDescriptorSacl(sd, &present, &sacl, &defaulted)
	if !present {
		err = ERROR_OBJECT_NOT_FOUND
	}
	return
}

// SetSACL 设置安全描述符的绝对SACL。

// ff:
// defaulted:
// present:
// sacl:
func (absoluteSD *SECURITY_DESCRIPTOR) SetSACL(sacl *ACL, present, defaulted bool) error {
	return setSecurityDescriptorSacl(absoluteSD, present, sacl, defaulted)
}

// Owner 返回安全描述符的所有者以及它是否为默认值。

// ff:
// err:
// defaulted:
// owner:
func (sd *SECURITY_DESCRIPTOR) Owner() (owner *SID, defaulted bool, err error) {
	err = getSecurityDescriptorOwner(sd, &owner, &defaulted)
	return
}

// SetOwner 设置绝对安全描述符的所有者。

// ff:
// defaulted:
// owner:
func (absoluteSD *SECURITY_DESCRIPTOR) SetOwner(owner *SID, defaulted bool) error {
	return setSecurityDescriptorOwner(absoluteSD, owner, defaulted)
}

// Group 返回安全描述符的组信息及其是否为默认值。

// ff:
// err:
// defaulted:
// group:
func (sd *SECURITY_DESCRIPTOR) Group() (group *SID, defaulted bool, err error) {
	err = getSecurityDescriptorGroup(sd, &group, &defaulted)
	return
}

// SetGroup 设置绝对安全描述符的所有者。

// ff:
// defaulted:
// group:
func (absoluteSD *SECURITY_DESCRIPTOR) SetGroup(group *SID, defaulted bool) error {
	return setSecurityDescriptorGroup(absoluteSD, group, defaulted)
}

// Length 返回安全描述符的长度。

// ff:
func (sd *SECURITY_DESCRIPTOR) Length() uint32 {
	return getSecurityDescriptorLength(sd)
}

// IsValid 返回该安全描述符是否有效。

// ff:
func (sd *SECURITY_DESCRIPTOR) IsValid() bool {
	return isValidSecurityDescriptor(sd)
}

// String 方法返回安全描述符的 SDDL 形式，其函数签名适用于使用 %v 格式化指令。

// ff:
func (sd *SECURITY_DESCRIPTOR) String() string {
	var sddl *uint16
	err := convertSecurityDescriptorToStringSecurityDescriptor(sd, 1, 0xff, &sddl, nil)
	if err != nil {
		return ""
	}
	defer LocalFree(Handle(unsafe.Pointer(sddl)))
	return UTF16PtrToString(sddl)
}

// ToAbsolute 将一个自相对的（self-relative）安全描述符转换为绝对的（absolute）安全描述符。

// ff:
// err:
// absoluteSD:
func (selfRelativeSD *SECURITY_DESCRIPTOR) ToAbsolute() (absoluteSD *SECURITY_DESCRIPTOR, err error) {
	control, _, err := selfRelativeSD.Control()
	if err != nil {
		return
	}
	if control&SE_SELF_RELATIVE == 0 {
		err = ERROR_INVALID_PARAMETER
		return
	}
	var absoluteSDSize, daclSize, saclSize, ownerSize, groupSize uint32
	err = makeAbsoluteSD(selfRelativeSD, nil, &absoluteSDSize,
		nil, &daclSize, nil, &saclSize, nil, &ownerSize, nil, &groupSize)
	switch err {
	case ERROR_INSUFFICIENT_BUFFER:
	case nil:
		// makeAbsoluteSD 预期会失败，但它却成功了。
		return nil, ERROR_INTERNAL_ERROR
	default:
		return nil, err
	}
	if absoluteSDSize > 0 {
		absoluteSD = (*SECURITY_DESCRIPTOR)(unsafe.Pointer(&make([]byte, absoluteSDSize)[0]))
	}
	var (
		dacl  *ACL
		sacl  *ACL
		owner *SID
		group *SID
	)
	if daclSize > 0 {
		dacl = (*ACL)(unsafe.Pointer(&make([]byte, daclSize)[0]))
	}
	if saclSize > 0 {
		sacl = (*ACL)(unsafe.Pointer(&make([]byte, saclSize)[0]))
	}
	if ownerSize > 0 {
		owner = (*SID)(unsafe.Pointer(&make([]byte, ownerSize)[0]))
	}
	if groupSize > 0 {
		group = (*SID)(unsafe.Pointer(&make([]byte, groupSize)[0]))
	}
	err = makeAbsoluteSD(selfRelativeSD, absoluteSD, &absoluteSDSize,
		dacl, &daclSize, sacl, &saclSize, owner, &ownerSize, group, &groupSize)
	return
}

// ToSelfRelative 将一个绝对形式的安全描述符转换为自相对形式。

// ff:
// err:
// selfRelativeSD:
func (absoluteSD *SECURITY_DESCRIPTOR) ToSelfRelative() (selfRelativeSD *SECURITY_DESCRIPTOR, err error) {
	control, _, err := absoluteSD.Control()
	if err != nil {
		return
	}
	if control&SE_SELF_RELATIVE != 0 {
		err = ERROR_INVALID_PARAMETER
		return
	}
	var selfRelativeSDSize uint32
	err = makeSelfRelativeSD(absoluteSD, nil, &selfRelativeSDSize)
	switch err {
	case ERROR_INSUFFICIENT_BUFFER:
	case nil:
		// makeSelfRelativeSD 预计会失败，但它却成功了。
		return nil, ERROR_INTERNAL_ERROR
	default:
		return nil, err
	}
	if selfRelativeSDSize > 0 {
		selfRelativeSD = (*SECURITY_DESCRIPTOR)(unsafe.Pointer(&make([]byte, selfRelativeSDSize)[0]))
	}
	err = makeSelfRelativeSD(absoluteSD, selfRelativeSD, &selfRelativeSDSize)
	return
}

func (selfRelativeSD *SECURITY_DESCRIPTOR) copySelfRelativeSecurityDescriptor() *SECURITY_DESCRIPTOR {
	sdLen := int(selfRelativeSD.Length())
	const min = int(unsafe.Sizeof(SECURITY_DESCRIPTOR{}))
	if sdLen < min {
		sdLen = min
	}

	src := unsafe.Slice((*byte)(unsafe.Pointer(selfRelativeSD)), sdLen)
// SECURITY_DESCRIPTOR 结构中包含指针，这意味着 checkptr 预期其应正确对齐。当我们把一个由 Windows 分配的结构复制到一个由 Go 分配的结构时，确保 Go 的内存分配按指针大小对齐。
	const psize = int(unsafe.Sizeof(uintptr(0)))
	alloc := make([]uintptr, (sdLen+psize-1)/psize)
	dst := unsafe.Slice((*byte)(unsafe.Pointer(&alloc[0])), sdLen)
	copy(dst, src)
	return (*SECURITY_DESCRIPTOR)(unsafe.Pointer(&dst[0]))
}

// SecurityDescriptorFromString 将描述安全描述符的 SDDL 字符串转换为在 Go 堆上分配的自相对安全描述符对象。

// ff:
// err:
// sd:
// sddl:
func SecurityDescriptorFromString(sddl string) (sd *SECURITY_DESCRIPTOR, err error) {
	var winHeapSD *SECURITY_DESCRIPTOR
	err = convertStringSecurityDescriptorToSecurityDescriptor(sddl, 1, &winHeapSD, nil)
	if err != nil {
		return
	}
	defer LocalFree(Handle(unsafe.Pointer(winHeapSD)))
	return winHeapSD.copySelfRelativeSecurityDescriptor(), nil
}

// GetSecurityInfo 通过给定的句柄查询安全信息，并在Go堆上返回自相关的安全描述符结果。

// ff:
// err:
// sd:
// securityInformation:
// objectType:
// handle:
func GetSecurityInfo(handle Handle, objectType SE_OBJECT_TYPE, securityInformation SECURITY_INFORMATION) (sd *SECURITY_DESCRIPTOR, err error) {
	var winHeapSD *SECURITY_DESCRIPTOR
	err = getSecurityInfo(handle, objectType, securityInformation, nil, nil, nil, nil, &winHeapSD)
	if err != nil {
		return
	}
	defer LocalFree(Handle(unsafe.Pointer(winHeapSD)))
	return winHeapSD.copySelfRelativeSecurityDescriptor(), nil
}

// GetNamedSecurityInfo 为给定的命名对象查询安全信息，并在Go堆上返回自相关的安全描述符结果。

// ff:
// err:
// sd:
// securityInformation:
// objectType:
// objectName:
func GetNamedSecurityInfo(objectName string, objectType SE_OBJECT_TYPE, securityInformation SECURITY_INFORMATION) (sd *SECURITY_DESCRIPTOR, err error) {
	var winHeapSD *SECURITY_DESCRIPTOR
	err = getNamedSecurityInfo(objectName, objectType, securityInformation, nil, nil, nil, nil, &winHeapSD)
	if err != nil {
		return
	}
	defer LocalFree(Handle(unsafe.Pointer(winHeapSD)))
	return winHeapSD.copySelfRelativeSecurityDescriptor(), nil
}

// BuildSecurityDescriptor 通过使用输入的信任主体、显式访问列表以及待合并的先前安全描述符（这些参数均可为 nil）构建一个新的安全描述符，并在 Go 堆上返回生成的自相对安全描述符结果。

// ff:
// err:
// sd:
// mergedSecurityDescriptor:
// auditEntries:
// accessEntries:
// group:
// owner:
func BuildSecurityDescriptor(owner *TRUSTEE, group *TRUSTEE, accessEntries []EXPLICIT_ACCESS, auditEntries []EXPLICIT_ACCESS, mergedSecurityDescriptor *SECURITY_DESCRIPTOR) (sd *SECURITY_DESCRIPTOR, err error) {
	var winHeapSD *SECURITY_DESCRIPTOR
	var winHeapSDSize uint32
	var firstAccessEntry *EXPLICIT_ACCESS
	if len(accessEntries) > 0 {
		firstAccessEntry = &accessEntries[0]
	}
	var firstAuditEntry *EXPLICIT_ACCESS
	if len(auditEntries) > 0 {
		firstAuditEntry = &auditEntries[0]
	}
	err = buildSecurityDescriptor(owner, group, uint32(len(accessEntries)), firstAccessEntry, uint32(len(auditEntries)), firstAuditEntry, mergedSecurityDescriptor, &winHeapSDSize, &winHeapSD)
	if err != nil {
		return
	}
	defer LocalFree(Handle(unsafe.Pointer(winHeapSD)))
	return winHeapSD.copySelfRelativeSecurityDescriptor(), nil
}

// NewSecurityDescriptor 创建并初始化一个新的绝对安全描述符

// ff:
// err:
// absoluteSD:
func NewSecurityDescriptor() (absoluteSD *SECURITY_DESCRIPTOR, err error) {
	absoluteSD = &SECURITY_DESCRIPTOR{}
	err = initializeSecurityDescriptor(absoluteSD, 1)
	return
}

// ACLFromEntries 在Go堆上返回一个新的ACL，其中包含一个显式条目列表以及另一个ACL的条目。
// explicitEntries 和 mergedACL 两者都是可选的，可以为nil。

// ff:
// err:
// acl:
// mergedACL:
// explicitEntries:
func ACLFromEntries(explicitEntries []EXPLICIT_ACCESS, mergedACL *ACL) (acl *ACL, err error) {
	var firstExplicitEntry *EXPLICIT_ACCESS
	if len(explicitEntries) > 0 {
		firstExplicitEntry = &explicitEntries[0]
	}
	var winHeapACL *ACL
	err = setEntriesInAcl(uint32(len(explicitEntries)), firstExplicitEntry, mergedACL, &winHeapACL)
	if err != nil {
		return
	}
	defer LocalFree(Handle(unsafe.Pointer(winHeapACL)))
	aclBytes := make([]byte, winHeapACL.aclSize)
	copy(aclBytes, (*[(1 << 31) - 1]byte)(unsafe.Pointer(winHeapACL))[:len(aclBytes):len(aclBytes)])
	return (*ACL)(unsafe.Pointer(&aclBytes[0])), nil
}
