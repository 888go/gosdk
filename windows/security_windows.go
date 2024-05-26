// 版权所有 2012 The Go 作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

package windows

import (
	"syscall"
	"unsafe"
)

// 翻译提示:const  (
// 	名称未知                    =  0
// 	完全限定DN名称          =  1
// 	SAM兼容名称              =  2
// 	显示名称                      =  3
// 	唯一标识名称                =  6
// 	规范名称                      =  7
// 	用户主体名称              =  8
// 	规范名称扩展                =  9
// 	服务主体名称              =  10
// 	DNS域名名称                =  12
// )
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

// This function returns 1 byte BOOLEAN rather than the 4 byte BOOL.
// http://blogs.msdn.com/b/drnick/archive/2007/12/19/windows-and-upn-format-credentials.aspx
//sys	TranslateName(accName *uint16, accNameFormat uint32, desiredNameFormat uint32, translatedName *uint16, nSize *uint32) (err error) [failretval&0xff==0] = secur32.TranslateNameW
//sys	GetUserNameEx(nameFormat uint32, nameBuffre *uint16, nSize *uint32) (err error) [failretval&0xff==0] = secur32.GetUserNameExW

// TranslateAccountName 将目录服务对象名称从一种格式转换为另一种格式。

// ff:
// username:
// from:
// to:
// initSize:
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

// 翻译提示:const  (
// 	//  不要重新排序
// 	网络配置未知状态  =  iota
// 	未加入网络
// 	工作组名称
// 	域名名称
// )
// 
// func  (status  NetworkConfigurationStatus)  String()  string  {
// 	switch  status  {
// 	case  NetSetupUnknownStatus:
// 		return  "未知状态"
// 	case  NetSetupUnjoined:
// 		return  "未加入"
// 	case  NetSetupWorkgroupName:
// 		return  "工作组名称"
// 	case  NetSetupDomainName:
// 		return  "域名名称"
// 	default:
// 		return  "未知状态("  +  strconv.Itoa(int(status))  +  ")"
// 	}
// }
// 
// type  NetworkConfigurationStatus  int
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
//sys	NetGetJoinInformation(server *uint16, name **uint16, bufType *uint32) (neterr error) = netapi32.NetGetJoinInformation
//sys	NetApiBufferFree(buf *byte) (neterr error) = netapi32.NetApiBufferFree

// 翻译提示:const  (
// 	//  不要重新排列
// 	用户Sid类型  =  1  +  iota
// 	组Sid类型
// 	域Sid类型
// 	别名Sid类型
// 	知名组Sid类型
// 	已删除账户Sid类型
// 	无效Sid类型
// 	未知Sid类型
// 	计算机Sid类型
// 	标签Sid类型
// )
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

// 翻译提示:const  (
// 	安全空标识符                                      =  0
// 	全局世界标识符                                    =  0
// 	本地标识符                                            =  0
// 	创建者所有者标识符                    =  0
// 	创建者组标识符                    =  1
// 	拨号标识符                                          =  1
// 	网络标识符                                            =  2
// 	批处理标识符                                        =  3
// 	交互式标识符                                        =  4
// 	登录标识符集合标识符                        =  5
// 	服务标识符                                            =  6
// 	本地系统标识符                                =  18
// 	内置域标识符                                      =  32
// 	主体自身标识符                              =  10
// 	创建者所有者服务器标识符      =  0x2
// 	创建者组服务器标识符      =  0x3
// 	登录标识符集合计数                =  0x3
// 	匿名登录标识符                          =  0x7
// 	代理标识符                                          =  0x8
// 	企业控制器标识符                      =  0x9
// 	服务器登录标识符                      =  企业控制器标识符
// 	验证用户标识符                          =  0xb
// 	受限代码标识符                                =  0xc
// 	NT非唯一标识符                            =  0x15
// )
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
// 翻译提示:const  (
// 	管理员组RID  =  0x220
// 	用户组RID      =  0x221
// 	访客组RID      =  0x222
// 	高级用户组RID  =  0x223
// 	账户操作员组RID  =  0x224
// 	系统操作员组RID  =  0x225
// 	打印操作员组RID  =  0x226
// 	备份操作员组RID  =  0x227
// 	复制器组RID  =  0x228
// 	远程访问服务器组RID  =  0x229
// 	旧版Windows兼容访问组RID  =  0x22a
// 	远程桌面用户组RID  =  0x22b
// 	网络配置操作员组RID  =  0x22c
// 	森林信任构建者组RID  =  0x22d
// 	监控用户组RID  =  0x22e
// 	日志记录用户组RID  =  0x22f
// 	授权访问组RID  =  0x230
// 	终端服务许可证服务器组RID  =  0x231
// 	DCOM用户组RID  =  0x232
// 	IIS用户组RID  =  0x238
// 	加密操作员组RID  =  0x239
// 	可缓存主体组RID  =  0x23b
// 	不可缓存主体组RID  =  0x23c
// 	事件日志读者组RID  =  0x23d
// 	CertSvc_DCOM访问组RID  =  0x23e
// )
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

// 安全标识符（SID）结构是一种可变长度的结构，用于唯一标识用户或组。
type SID struct{}

// StringToSid将字符串格式的安全标识符（SID）转换为有效且可用的SID。

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

// LookupSID 获取指定账户的安全标识符（SID）以及该账户所在的域名名称。
// 参数System指定了进行搜索的目标计算机。

// ff:
// system:
// account:
// sid:
// domain:
// accType:
// err:
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

// String 将SID转换为适合显示、存储或传输的字符串格式
// 翻译提示://  SID字符串表示
// func  (sid  *安全标识符)  String()  字符串  {}

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

// Len 返回一个有效安全标识符（SID）的长度，以字节为单位。

// ff:
func (sid *SID) Len() int {
	return int(GetLengthSid(sid))
}

// Copy 创建一个安全标识符 SID 的副本

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

// IdentifierAuthority 返回该SID的标识符权限。

// ff:
func (sid *SID) IdentifierAuthority() SidIdentifierAuthority {
	return *getSidIdentifierAuthority(sid)
}

// SubAuthorityCount 返回SID中子授权机构的数量

// ff:
func (sid *SID) SubAuthorityCount() uint8 {
	return *getSidSubAuthorityCount(sid)
}

// SubAuthority 返回由指定索引（该索引必须小于 sid.SubAuthorityCount()）所确定的 SID 的子权限部分
// 翻译提示:func  (sid  *安全标识符)  子权威索引(idx  uint32)  uint32  {}

// ff:
// idx:
func (sid *SID) SubAuthority(idx uint32) uint32 {
	if idx >= uint32(sid.SubAuthorityCount()) {
		panic("sub-authority index out of range")
	}
	return *getSidSubAuthority(sid, idx)
}

// IsValid 判断SID是否具有有效的修订版本和长度
// 翻译提示:func  (sid  *安全标识符)  是否有效()  bool  {}

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

// IsWellKnown 判断SID是否匹配已知的sidType
// 翻译提示:func  (sid  *SID)  是否为常见SID(sidType  常见SID类型)  bool  {}

// ff:
// sidType:
func (sid *SID) IsWellKnown(sidType WELL_KNOWN_SID_TYPE) bool {
	return isWellKnownSid(sid, sidType)
}

// LookupAccount 通过此 SID 查找账户名以及该 SID 首次出现的域的名称。
// System 指定要在其中搜索的目标计算机。

// ff:
// system:
// account:
// domain:
// accType:
// err:
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

// 各种类型的预指定SID，可以在运行时进行合成与比较
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

// 为本地计算机上某个预定义的知名别名创建一个SID（安全标识符），通常使用形如Win*Sid的常量。

// ff:
// sidType:
func CreateWellKnownSid(sidType WELL_KNOWN_SID_TYPE) (*SID, error) {
	return CreateWellKnownDomainSid(sidType, nil)
}

// 为指定由 domainSid 参数标识的域中某个预定义的知名别名创建一个 SID（安全标识符），通常使用形如 Win*Sid 的常量。

// ff:
// sidType:
// domainSid:
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

// 翻译提示:const  (
// 	//  不要重新排序
// 	令牌设置为主要  =  1  <<  iota
// 	令牌复制
// 	令牌模拟
// 	令牌查询
// 	令牌查询来源
// 	令牌调整权限
// 	令牌调整组
// 	令牌调整默认值
// 	令牌调整会话ID
// 
// 	令牌所有访问权限  =  标准权利要求  |
// 		令牌设置为主要  |
// 		令牌复制  |
// 		令牌模拟  |
// 		令牌查询  |
// 		令牌查询来源  |
// 		令牌调整权限  |
// 		令牌调整组  |
// 		令牌调整默认值  |
// 		令牌调整会话ID
// 	令牌读取  =  标准权利读取  |  令牌查询
// 	令牌写入  =  标准权利写入  |
// 		令牌调整权限  |
// 		令牌调整组  |
// 		令牌调整默认值
// 	令牌执行  =  标准权利执行
// )
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

// 翻译提示:const  (
// 	//  不要重新排序
// 	用户令牌  =  1  +  iota
// 	组令牌
// 	权限令牌
// 	拥有者令牌
// 	主要组令牌
// 	默认访问控制列表令牌
// 	令牌源
// 	令牌类型
// 	令牌模拟级别
// 	令牌统计信息
// 	受限安全标识符令牌
// 	会话标识符令牌
// 	组和权限令牌
// 	令牌会话引用
// 	令牌沙箱无效
// 	令牌审计策略
// 	令牌来源
// 	令牌提升类型
// 	链接令牌
// 	令牌提升
// 	令牌有限制
// 	令牌访问信息
// 	令牌虚拟化允许
// 	令牌虚拟化启用
// 	完整性级别令牌
// 	UI访问令牌
// 	强制策略令牌
// 	登录Sid令牌
// 	最大令牌信息类别
// )
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
// 翻译提示:const  (
// 	默认启用权限                =  0x00000001
// 	已启用权限                    =  0x00000002
// 	已移除权限                    =  0x00000004
// 	用于访问的权限            =  0x80000000
// 	有效权限属性              =  默认启用权限  |  已启用权限  |  已移除权限  |  用于访问的权限
// )
const (
	SE_PRIVILEGE_ENABLED_BY_DEFAULT = 0x00000001
	SE_PRIVILEGE_ENABLED            = 0x00000002
	SE_PRIVILEGE_REMOVED            = 0x00000004
	SE_PRIVILEGE_USED_FOR_ACCESS    = 0x80000000
	SE_PRIVILEGE_VALID_ATTRIBUTES   = SE_PRIVILEGE_ENABLED_BY_DEFAULT | SE_PRIVILEGE_ENABLED | SE_PRIVILEGE_REMOVED | SE_PRIVILEGE_USED_FOR_ACCESS
)

// Token types
// 翻译提示:const  (
// 	主要令牌  =  1
// 	模拟令牌  =  2
// )
const (
	TokenPrimary       = 1
	TokenImpersonation = 2
)

// Impersonation levels
// 翻译提示:const  (
// 	匿名安全级别      =  0
// 	身份验证安全级别  =  1
// 	模拟安全级别        =  2
// 	委派安全级别        =  3
// )
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

// AllGroups 返回一个切片，可用于遍历 g 中的各组。

// ff:
func (g *Tokengroups) AllGroups() []SIDAndAttributes {
	return (*[(1 << 28) - 1]SIDAndAttributes)(unsafe.Pointer(&g.Groups[0]))[:g.GroupCount:g.GroupCount]
}

type Tokenprivileges struct {
	PrivilegeCount uint32
	Privileges     [1]LUIDAndAttributes // 使用 AllPrivileges() 进行遍历。
}

// AllPrivileges 返回一个切片，可用于遍历集合p中的所有权限。

// ff:
func (p *Tokenprivileges) AllPrivileges() []LUIDAndAttributes {
	return (*[(1 << 27) - 1]LUIDAndAttributes)(unsafe.Pointer(&p.Privileges[0]))[:p.PrivilegeCount:p.PrivilegeCount]
}

type Tokenmandatorylabel struct {
	Label SIDAndAttributes
}

// 翻译提示:func  (tml  *安全标识符标签)  大小()  uint32  {}

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
// 系统在用户登录时创建一个访问令牌，且为该用户执行的每个进程都有一份该令牌的副本。
// 该令牌用于标识用户、用户的组以及用户的权限。系统使用该令牌来控制对可保护对象的访问，并控制用户在本地计算机上执行各种系统相关操作的能力。
type Token Handle

// OpenCurrentProcessToken 以 TOKEN_QUERY 访问权限打开与当前进程关联的访问令牌。这是一个需要被关闭的真实令牌。
//
// 已弃用：请显式调用 OpenProcessToken(CurrentProcess(), ...) 并指定所需的访问权限，或者使用 GetCurrentProcessToken 获取 TOKEN_QUERY 令牌。

// ff:
// Token:
func OpenCurrentProcessToken() (Token, error) {
	var token Token
	err := OpenProcessToken(CurrentProcess(), TOKEN_QUERY, &token)
	return token, err
}

// GetCurrentProcessToken 获取与当前进程关联的访问令牌。这是一个无需关闭的伪令牌。

// ff:
func GetCurrentProcessToken() Token {
	return Token(^uintptr(4 - 1))
}

// GetCurrentThreadToken 返回当前线程关联的访问令牌。它是一个无需关闭的伪令牌。

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

// Close 释放对访问令牌的访问权限。

// ff:
func (t Token) Close() error {
	return CloseHandle(Handle(t))
}

// getInfo 用于获取访问令牌指定类型的信息。
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
// 返回一个指向 SID 结构的指针，该结构代表一个组，此组将成为使用此访问令牌的进程创建的任何对象的主组。

// ff:
func (t Token) GetTokenPrimaryGroup() (*Tokenprimarygroup, error) {
	i, e := t.getInfo(TokenPrimaryGroup, 50)
	if e != nil {
		return nil, e
	}
	return (*Tokenprimarygroup)(i), nil
}

// GetUserProfileDirectory 获取访问令牌t对应用户的个人资料根目录路径。

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

// IsElevated 返回当前令牌是否从UAC角度来看具有提升权限。

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

// GetSystemDirectory 获取当前系统目录的路径，该目录通常（但并不总是）为 `C:\Windows\System32`。

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

// GetWindowsDirectory 函数用于获取当前 Windows 目录的路径，该目录通常（但并非总是）为 `C:\Windows`。在应用程序于终端服务器环境下运行时，此目录可能为用户的个人专属目录。

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
// isRestricted:
// err:
func (t Token) IsRestricted() (isRestricted bool, err error) {
	isRestricted, err = isTokenRestricted(t)
	if !isRestricted && err == syscall.EINVAL {
		// 如果err为EINVAL，返回ERROR_SUCCESS，表示该令牌不受限制
		err = nil
	}
	return
}

// 翻译提示:const  (
// 	控制台连接                =  0x1
// 	控制台断开连接          =  0x2
// 	远程连接                  =  0x3
// 	远程断开连接            =  0x4
// 	会话登录                    =  0x5
// 	会话注销                    =  0x6
// 	会话锁定                      =  0x7
// 	会话解锁                      =  0x8
// 	会话远程控制              =  0x9
// 	会话创建                    =  0xa
// 	会话终止                    =  0xb
// )
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

// 翻译提示:const  (
// 	活动状态              =  0
// 	已连接                =  1
// 	连接查询            =  2
// 	阴影连接            =  3
// 	已断开连接          =  4
// 	空闲状态            =  5
// 	监听状态            =  6
// 	重置状态            =  7
// 	系统关闭            =  8
// 	初始化状态        =  9
// )
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
//sys WTSEnumerateSessions(handle Handle, reserved uint32, version uint32, sessions **WTS_SESSION_INFO, count *uint32) (err error) = wtsapi32.WTSEnumerateSessionsW
//sys WTSFreeMemory(ptr uintptr) = wtsapi32.WTSFreeMemory
//sys WTSGetActiveConsoleSessionId() (sessionID uint32)

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
// 翻译提示:const  (
// 	安全静态追踪  =  0
// 	安全动态追踪  =  1
// )
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

// SECURITY_INFORMATION 类型的常量
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

// 用于AceFlags和Inheritance字段的常量
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

// TRUSTEE_TYPE类型的常量
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

// 对于ObjectsPresent字段的常量
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

// 此类型为TRUSTEE内的联合体，必须通过使用TrusteeValueFrom*系列函数之一来创建。
type TrusteeValue uintptr

// 翻译提示:func  从字符串转换TrusteeValue(str  字符串)  TrusteeValue  {}
// ```

// ff:
// str:
func TrusteeValueFromString(str string) TrusteeValue {
	return TrusteeValue(unsafe.Pointer(StringToUTF16Ptr(str)))
}
// 翻译提示:func  从SID获取受托者值(sid  *安全标识符)  受托者值  {}

// ff:
// sid:
func TrusteeValueFromSID(sid *SID) TrusteeValue {
	return TrusteeValue(unsafe.Pointer(sid))
}
// 翻译提示:func  从对象和Sid获取Trustee值(objectsAndSid  *对象和Sid)  TrusteeValue  {}

// ff:
// objectsAndSid:
func TrusteeValueFromObjectsAndSid(objectsAndSid *OBJECTS_AND_SID) TrusteeValue {
	return TrusteeValue(unsafe.Pointer(objectsAndSid))
}
// 翻译提示:func  从对象和名称获取受托人值(objectsAndName  *对象和名称)  受托人值  {}

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

//sys	getSecurityInfo(handle Handle, objectType SE_OBJECT_TYPE, securityInformation SECURITY_INFORMATION, owner **SID, group **SID, dacl **ACL, sacl **ACL, sd **SECURITY_DESCRIPTOR) (ret error) = advapi32.GetSecurityInfo
//sys	SetSecurityInfo(handle Handle, objectType SE_OBJECT_TYPE, securityInformation SECURITY_INFORMATION, owner *SID, group *SID, dacl *ACL, sacl *ACL) (ret error) = advapi32.SetSecurityInfo
//sys	getNamedSecurityInfo(objectName string, objectType SE_OBJECT_TYPE, securityInformation SECURITY_INFORMATION, owner **SID, group **SID, dacl **ACL, sacl **ACL, sd **SECURITY_DESCRIPTOR) (ret error) = advapi32.GetNamedSecurityInfoW
//sys	SetNamedSecurityInfo(objectName string, objectType SE_OBJECT_TYPE, securityInformation SECURITY_INFORMATION, owner *SID, group *SID, dacl *ACL, sacl *ACL) (ret error) = advapi32.SetNamedSecurityInfoW
//sys	SetKernelObjectSecurity(handle Handle, securityInformation SECURITY_INFORMATION, securityDescriptor *SECURITY_DESCRIPTOR) (err error) = advapi32.SetKernelObjectSecurity

//sys	buildSecurityDescriptor(owner *TRUSTEE, group *TRUSTEE, countAccessEntries uint32, accessEntries *EXPLICIT_ACCESS, countAuditEntries uint32, auditEntries *EXPLICIT_ACCESS, oldSecurityDescriptor *SECURITY_DESCRIPTOR, sizeNewSecurityDescriptor *uint32, newSecurityDescriptor **SECURITY_DESCRIPTOR) (ret error) = advapi32.BuildSecurityDescriptorW
//sys	initializeSecurityDescriptor(absoluteSD *SECURITY_DESCRIPTOR, revision uint32) (err error) = advapi32.InitializeSecurityDescriptor

//sys	getSecurityDescriptorControl(sd *SECURITY_DESCRIPTOR, control *SECURITY_DESCRIPTOR_CONTROL, revision *uint32) (err error) = advapi32.GetSecurityDescriptorControl
//sys	getSecurityDescriptorDacl(sd *SECURITY_DESCRIPTOR, daclPresent *bool, dacl **ACL, daclDefaulted *bool) (err error) = advapi32.GetSecurityDescriptorDacl
//sys	getSecurityDescriptorSacl(sd *SECURITY_DESCRIPTOR, saclPresent *bool, sacl **ACL, saclDefaulted *bool) (err error) = advapi32.GetSecurityDescriptorSacl
//sys	getSecurityDescriptorOwner(sd *SECURITY_DESCRIPTOR, owner **SID, ownerDefaulted *bool) (err error) = advapi32.GetSecurityDescriptorOwner
//sys	getSecurityDescriptorGroup(sd *SECURITY_DESCRIPTOR, group **SID, groupDefaulted *bool) (err error) = advapi32.GetSecurityDescriptorGroup
//sys	getSecurityDescriptorLength(sd *SECURITY_DESCRIPTOR) (len uint32) = advapi32.GetSecurityDescriptorLength
//sys	getSecurityDescriptorRMControl(sd *SECURITY_DESCRIPTOR, rmControl *uint8) (ret error) [failretval!=0] = advapi32.GetSecurityDescriptorRMControl
//sys	isValidSecurityDescriptor(sd *SECURITY_DESCRIPTOR) (isValid bool) = advapi32.IsValidSecurityDescriptor

//sys	setSecurityDescriptorControl(sd *SECURITY_DESCRIPTOR, controlBitsOfInterest SECURITY_DESCRIPTOR_CONTROL, controlBitsToSet SECURITY_DESCRIPTOR_CONTROL) (err error) = advapi32.SetSecurityDescriptorControl
//sys	setSecurityDescriptorDacl(sd *SECURITY_DESCRIPTOR, daclPresent bool, dacl *ACL, daclDefaulted bool) (err error) = advapi32.SetSecurityDescriptorDacl
//sys	setSecurityDescriptorSacl(sd *SECURITY_DESCRIPTOR, saclPresent bool, sacl *ACL, saclDefaulted bool) (err error) = advapi32.SetSecurityDescriptorSacl
//sys	setSecurityDescriptorOwner(sd *SECURITY_DESCRIPTOR, owner *SID, ownerDefaulted bool) (err error) = advapi32.SetSecurityDescriptorOwner
//sys	setSecurityDescriptorGroup(sd *SECURITY_DESCRIPTOR, group *SID, groupDefaulted bool) (err error) = advapi32.SetSecurityDescriptorGroup
//sys	setSecurityDescriptorRMControl(sd *SECURITY_DESCRIPTOR, rmControl *uint8) = advapi32.SetSecurityDescriptorRMControl

//sys	convertStringSecurityDescriptorToSecurityDescriptor(str string, revision uint32, sd **SECURITY_DESCRIPTOR, size *uint32) (err error) = advapi32.ConvertStringSecurityDescriptorToSecurityDescriptorW
//sys	convertSecurityDescriptorToStringSecurityDescriptor(sd *SECURITY_DESCRIPTOR, revision uint32, securityInformation SECURITY_INFORMATION, str **uint16, strLen *uint32) (err error) = advapi32.ConvertSecurityDescriptorToStringSecurityDescriptorW

//sys	makeAbsoluteSD(selfRelativeSD *SECURITY_DESCRIPTOR, absoluteSD *SECURITY_DESCRIPTOR, absoluteSDSize *uint32, dacl *ACL, daclSize *uint32, sacl *ACL, saclSize *uint32, owner *SID, ownerSize *uint32, group *SID, groupSize *uint32) (err error) = advapi32.MakeAbsoluteSD
//sys	makeSelfRelativeSD(absoluteSD *SECURITY_DESCRIPTOR, selfRelativeSD *SECURITY_DESCRIPTOR, selfRelativeSDSize *uint32) (err error) = advapi32.MakeSelfRelativeSD

//sys	setEntriesInAcl(countExplicitEntries uint32, explicitEntries *EXPLICIT_ACCESS, oldACL *ACL, newACL **ACL) (ret error) = advapi32.SetEntriesInAclW

// Control 返回安全描述符控制位。

// ff:
// control:
// revision:
// err:
func (sd *SECURITY_DESCRIPTOR) Control() (control SECURITY_DESCRIPTOR_CONTROL, revision uint32, err error) {
	err = getSecurityDescriptorControl(sd, &control, &revision)
	return
}

// SetControl 设置安全描述符控制位
// 翻译提示:func  (sd  *安全描述符)  设置控制位(controlBitsOfInterest  安全描述符控制位,  controlBitsToSet  安全描述符控制位)  error  {}

// ff:
// controlBitsOfInterest:
// controlBitsToSet:
func (sd *SECURITY_DESCRIPTOR) SetControl(controlBitsOfInterest SECURITY_DESCRIPTOR_CONTROL, controlBitsToSet SECURITY_DESCRIPTOR_CONTROL) error {
	return setSecurityDescriptorControl(sd, controlBitsOfInterest, controlBitsToSet)
}

// RMControl 返回安全描述符资源管理器控制位。

// ff:
// control:
// err:
func (sd *SECURITY_DESCRIPTOR) RMControl() (control uint8, err error) {
	err = getSecurityDescriptorRMControl(sd, &control)
	return
}

// SetRMControl 设置安全描述符资源管理器控制位

// ff:
// rmControl:
func (sd *SECURITY_DESCRIPTOR) SetRMControl(rmControl uint8) {
	setSecurityDescriptorRMControl(sd, &rmControl)
}

// DACL 返回安全描述符的 DACL 以及它是否为默认值。dacl 的返回值可能为 nil，如果存在但为“空 DACL”，即表示完全无限制访问。如果不存在 DACL，则 err 返回 ERROR_OBJECT_NOT_FOUND。

// ff:
// dacl:
// defaulted:
// err:
func (sd *SECURITY_DESCRIPTOR) DACL() (dacl *ACL, defaulted bool, err error) {
	var present bool
	err = getSecurityDescriptorDacl(sd, &present, &dacl, &defaulted)
	if !present {
		err = ERROR_OBJECT_NOT_FOUND
	}
	return
}

// SetDACL 设置绝对安全描述符的 DACL
// 翻译提示:func  (绝对安全描述符  *安全描述符)  设置DACL(访问控制列表  *ACL,  是否存在  bool,  是否默认  bool)  error  {}

// ff:
// dacl:
// present:
// defaulted:
func (absoluteSD *SECURITY_DESCRIPTOR) SetDACL(dacl *ACL, present, defaulted bool) error {
	return setSecurityDescriptorDacl(absoluteSD, present, dacl, defaulted)
}

// SACL 返回安全描述符的 SACL（系统访问控制列表）及其是否为默认值。如果存在 SACL，但它是“空 SACL”，即完全无限制，则 sacl 返回值可能为 nil。如果不存在 SACL，则 err 返回 ERROR_OBJECT_NOT_FOUND。

// ff:
// sacl:
// defaulted:
// err:
func (sd *SECURITY_DESCRIPTOR) SACL() (sacl *ACL, defaulted bool, err error) {
	var present bool
	err = getSecurityDescriptorSacl(sd, &present, &sacl, &defaulted)
	if !present {
		err = ERROR_OBJECT_NOT_FOUND
	}
	return
}

// SetSACL 设置绝对安全描述符的SACL
// 翻译提示:func  (绝对SD  *安全描述符)  设置系统访问控制列表(安全访问控制列表  *ACL,  是否存在  bool,  是否默认  bool)  error  {}

// ff:
// sacl:
// present:
// defaulted:
func (absoluteSD *SECURITY_DESCRIPTOR) SetSACL(sacl *ACL, present, defaulted bool) error {
	return setSecurityDescriptorSacl(absoluteSD, present, sacl, defaulted)
}

// Owner 返回安全描述符的所有者以及该所有者是否为默认值。

// ff:
// owner:
// defaulted:
// err:
func (sd *SECURITY_DESCRIPTOR) Owner() (owner *SID, defaulted bool, err error) {
	err = getSecurityDescriptorOwner(sd, &owner, &defaulted)
	return
}

// SetOwner 设置绝对安全描述符的所有者

// ff:
// owner:
// defaulted:
func (absoluteSD *SECURITY_DESCRIPTOR) SetOwner(owner *SID, defaulted bool) error {
	return setSecurityDescriptorOwner(absoluteSD, owner, defaulted)
}

// Group 返回安全描述符的组信息以及该组信息是否为默认值。

// ff:
// group:
// defaulted:
// err:
func (sd *SECURITY_DESCRIPTOR) Group() (group *SID, defaulted bool, err error) {
	err = getSecurityDescriptorGroup(sd, &group, &defaulted)
	return
}

// SetGroup 设置绝对安全描述符的所有者。

// ff:
// group:
// defaulted:
func (absoluteSD *SECURITY_DESCRIPTOR) SetGroup(group *SID, defaulted bool) error {
	return setSecurityDescriptorGroup(absoluteSD, group, defaulted)
}

// Length 返回安全描述符的长度。

// ff:
func (sd *SECURITY_DESCRIPTOR) Length() uint32 {
	return getSecurityDescriptorLength(sd)
}

// IsValid 返回该安全描述符是否有效
// 翻译提示:func  (sd  *安全描述符)  是否有效()  bool  {}

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
// absoluteSD:
// err:
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

// ToSelfRelative 将一个绝对安全描述符转换为自相对形式。

// ff:
// selfRelativeSD:
// err:
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
		// makeSelfRelativeSD 预期会失败，但实际上却成功了。
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
// SECURITY_DESCRIPTOR 结构体内含有指针，这意味着 checkptr 预期其应被正确对齐。当我们将一个由 Windows 分配的结构体复制到由 Go 分配的一个时，确保 Go 的内存分配按指针大小对齐。
	const psize = int(unsafe.Sizeof(uintptr(0)))
	alloc := make([]uintptr, (sdLen+psize-1)/psize)
	dst := unsafe.Slice((*byte)(unsafe.Pointer(&alloc[0])), sdLen)
	copy(dst, src)
	return (*SECURITY_DESCRIPTOR)(unsafe.Pointer(&dst[0]))
}

// SecurityDescriptorFromString 将描述安全描述符的 SDDL 字符串转换为在 Go 堆上分配的自相对性安全描述符对象。

// ff:
// sddl:
// sd:
// err:
func SecurityDescriptorFromString(sddl string) (sd *SECURITY_DESCRIPTOR, err error) {
	var winHeapSD *SECURITY_DESCRIPTOR
	err = convertStringSecurityDescriptorToSecurityDescriptor(sddl, 1, &winHeapSD, nil)
	if err != nil {
		return
	}
	defer LocalFree(Handle(unsafe.Pointer(winHeapSD)))
	return winHeapSD.copySelfRelativeSecurityDescriptor(), nil
}

// GetSecurityInfo 为给定的句柄查询安全信息，并在Go堆上返回自相关的安全描述符结果。

// ff:
// handle:
// objectType:
// securityInformation:
// sd:
// err:
func GetSecurityInfo(handle Handle, objectType SE_OBJECT_TYPE, securityInformation SECURITY_INFORMATION) (sd *SECURITY_DESCRIPTOR, err error) {
	var winHeapSD *SECURITY_DESCRIPTOR
	err = getSecurityInfo(handle, objectType, securityInformation, nil, nil, nil, nil, &winHeapSD)
	if err != nil {
		return
	}
	defer LocalFree(Handle(unsafe.Pointer(winHeapSD)))
	return winHeapSD.copySelfRelativeSecurityDescriptor(), nil
}

// GetNamedSecurityInfo 为指定命名对象查询安全信息，并在Go堆上返回自相对安全描述符结果。

// ff:
// objectName:
// objectType:
// securityInformation:
// sd:
// err:
func GetNamedSecurityInfo(objectName string, objectType SE_OBJECT_TYPE, securityInformation SECURITY_INFORMATION) (sd *SECURITY_DESCRIPTOR, err error) {
	var winHeapSD *SECURITY_DESCRIPTOR
	err = getNamedSecurityInfo(objectName, objectType, securityInformation, nil, nil, nil, nil, &winHeapSD)
	if err != nil {
		return
	}
	defer LocalFree(Handle(unsafe.Pointer(winHeapSD)))
	return winHeapSD.copySelfRelativeSecurityDescriptor(), nil
}

// BuildSecurityDescriptor 通过使用输入的信任主体、显式访问列表以及待合并的先前安全描述符（这些参数均可为 nil）构建一个新的安全描述符，并将生成的自相关安全描述符结果置于 Go 堆上返回。

// ff:
// owner:
// group:
// accessEntries:
// auditEntries:
// mergedSecurityDescriptor:
// sd:
// err:
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

// NewSecurityDescriptor 创建并初始化一个新的绝对安全描述符。

// ff:
// absoluteSD:
// err:
func NewSecurityDescriptor() (absoluteSD *SECURITY_DESCRIPTOR, err error) {
	absoluteSD = &SECURITY_DESCRIPTOR{}
	err = initializeSecurityDescriptor(absoluteSD, 1)
	return
}

// ACLFromEntries 返回一个位于Go堆上的新ACL，其中包含一个显式条目列表以及另一个ACL的条目。
// explicitEntries 和 mergedACL 两者都是可选的，可以为nil。

// ff:
// explicitEntries:
// mergedACL:
// acl:
// err:
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
