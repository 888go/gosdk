// 版权所有 ? 2021 Go作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

package windows

import (
	"encoding/binary"
	"errors"
	"fmt"
	"runtime"
	"strings"
	"syscall"
	"unsafe"
)

// 本文件包含对SetupAPI.dll和CfgMgr32.dll的封装函数，
// 这些是用于管理硬件设备、驱动程序及PnP树的核心系统功能。
// 可在以下链接获取关于这些API的信息：
//     https://docs.microsoft.com/zh-cn/windows-hardware/drivers/install/setupapi
//     https://docs.microsoft.com/zh-cn/windows/win32/devinst/cfgmgr32-

const (
	ERROR_EXPECTED_SECTION_NAME                  Errno = 0x20000000 | 0xC0000000 | 0
	ERROR_BAD_SECTION_NAME_LINE                  Errno = 0x20000000 | 0xC0000000 | 1
	ERROR_SECTION_NAME_TOO_LONG                  Errno = 0x20000000 | 0xC0000000 | 2
	ERROR_GENERAL_SYNTAX                         Errno = 0x20000000 | 0xC0000000 | 3
	ERROR_WRONG_INF_STYLE                        Errno = 0x20000000 | 0xC0000000 | 0x100
	ERROR_SECTION_NOT_FOUND                      Errno = 0x20000000 | 0xC0000000 | 0x101
	ERROR_LINE_NOT_FOUND                         Errno = 0x20000000 | 0xC0000000 | 0x102
	ERROR_NO_BACKUP                              Errno = 0x20000000 | 0xC0000000 | 0x103
	ERROR_NO_ASSOCIATED_CLASS                    Errno = 0x20000000 | 0xC0000000 | 0x200
	ERROR_CLASS_MISMATCH                         Errno = 0x20000000 | 0xC0000000 | 0x201
	ERROR_DUPLICATE_FOUND                        Errno = 0x20000000 | 0xC0000000 | 0x202
	ERROR_NO_DRIVER_SELECTED                     Errno = 0x20000000 | 0xC0000000 | 0x203
	ERROR_KEY_DOES_NOT_EXIST                     Errno = 0x20000000 | 0xC0000000 | 0x204
	ERROR_INVALID_DEVINST_NAME                   Errno = 0x20000000 | 0xC0000000 | 0x205
	ERROR_INVALID_CLASS                          Errno = 0x20000000 | 0xC0000000 | 0x206
	ERROR_DEVINST_ALREADY_EXISTS                 Errno = 0x20000000 | 0xC0000000 | 0x207
	ERROR_DEVINFO_NOT_REGISTERED                 Errno = 0x20000000 | 0xC0000000 | 0x208
	ERROR_INVALID_REG_PROPERTY                   Errno = 0x20000000 | 0xC0000000 | 0x209
	ERROR_NO_INF                                 Errno = 0x20000000 | 0xC0000000 | 0x20A
	ERROR_NO_SUCH_DEVINST                        Errno = 0x20000000 | 0xC0000000 | 0x20B
	ERROR_CANT_LOAD_CLASS_ICON                   Errno = 0x20000000 | 0xC0000000 | 0x20C
	ERROR_INVALID_CLASS_INSTALLER                Errno = 0x20000000 | 0xC0000000 | 0x20D
	ERROR_DI_DO_DEFAULT                          Errno = 0x20000000 | 0xC0000000 | 0x20E
	ERROR_DI_NOFILECOPY                          Errno = 0x20000000 | 0xC0000000 | 0x20F
	ERROR_INVALID_HWPROFILE                      Errno = 0x20000000 | 0xC0000000 | 0x210
	ERROR_NO_DEVICE_SELECTED                     Errno = 0x20000000 | 0xC0000000 | 0x211
	ERROR_DEVINFO_LIST_LOCKED                    Errno = 0x20000000 | 0xC0000000 | 0x212
	ERROR_DEVINFO_DATA_LOCKED                    Errno = 0x20000000 | 0xC0000000 | 0x213
	ERROR_DI_BAD_PATH                            Errno = 0x20000000 | 0xC0000000 | 0x214
	ERROR_NO_CLASSINSTALL_PARAMS                 Errno = 0x20000000 | 0xC0000000 | 0x215
	ERROR_FILEQUEUE_LOCKED                       Errno = 0x20000000 | 0xC0000000 | 0x216
	ERROR_BAD_SERVICE_INSTALLSECT                Errno = 0x20000000 | 0xC0000000 | 0x217
	ERROR_NO_CLASS_DRIVER_LIST                   Errno = 0x20000000 | 0xC0000000 | 0x218
	ERROR_NO_ASSOCIATED_SERVICE                  Errno = 0x20000000 | 0xC0000000 | 0x219
	ERROR_NO_DEFAULT_DEVICE_INTERFACE            Errno = 0x20000000 | 0xC0000000 | 0x21A
	ERROR_DEVICE_INTERFACE_ACTIVE                Errno = 0x20000000 | 0xC0000000 | 0x21B
	ERROR_DEVICE_INTERFACE_REMOVED               Errno = 0x20000000 | 0xC0000000 | 0x21C
	ERROR_BAD_INTERFACE_INSTALLSECT              Errno = 0x20000000 | 0xC0000000 | 0x21D
	ERROR_NO_SUCH_INTERFACE_CLASS                Errno = 0x20000000 | 0xC0000000 | 0x21E
	ERROR_INVALID_REFERENCE_STRING               Errno = 0x20000000 | 0xC0000000 | 0x21F
	ERROR_INVALID_MACHINENAME                    Errno = 0x20000000 | 0xC0000000 | 0x220
	ERROR_REMOTE_COMM_FAILURE                    Errno = 0x20000000 | 0xC0000000 | 0x221
	ERROR_MACHINE_UNAVAILABLE                    Errno = 0x20000000 | 0xC0000000 | 0x222
	ERROR_NO_CONFIGMGR_SERVICES                  Errno = 0x20000000 | 0xC0000000 | 0x223
	ERROR_INVALID_PROPPAGE_PROVIDER              Errno = 0x20000000 | 0xC0000000 | 0x224
	ERROR_NO_SUCH_DEVICE_INTERFACE               Errno = 0x20000000 | 0xC0000000 | 0x225
	ERROR_DI_POSTPROCESSING_REQUIRED             Errno = 0x20000000 | 0xC0000000 | 0x226
	ERROR_INVALID_COINSTALLER                    Errno = 0x20000000 | 0xC0000000 | 0x227
	ERROR_NO_COMPAT_DRIVERS                      Errno = 0x20000000 | 0xC0000000 | 0x228
	ERROR_NO_DEVICE_ICON                         Errno = 0x20000000 | 0xC0000000 | 0x229
	ERROR_INVALID_INF_LOGCONFIG                  Errno = 0x20000000 | 0xC0000000 | 0x22A
	ERROR_DI_DONT_INSTALL                        Errno = 0x20000000 | 0xC0000000 | 0x22B
	ERROR_INVALID_FILTER_DRIVER                  Errno = 0x20000000 | 0xC0000000 | 0x22C
	ERROR_NON_WINDOWS_NT_DRIVER                  Errno = 0x20000000 | 0xC0000000 | 0x22D
	ERROR_NON_WINDOWS_DRIVER                     Errno = 0x20000000 | 0xC0000000 | 0x22E
	ERROR_NO_CATALOG_FOR_OEM_INF                 Errno = 0x20000000 | 0xC0000000 | 0x22F
	ERROR_DEVINSTALL_QUEUE_NONNATIVE             Errno = 0x20000000 | 0xC0000000 | 0x230
	ERROR_NOT_DISABLEABLE                        Errno = 0x20000000 | 0xC0000000 | 0x231
	ERROR_CANT_REMOVE_DEVINST                    Errno = 0x20000000 | 0xC0000000 | 0x232
	ERROR_INVALID_TARGET                         Errno = 0x20000000 | 0xC0000000 | 0x233
	ERROR_DRIVER_NONNATIVE                       Errno = 0x20000000 | 0xC0000000 | 0x234
	ERROR_IN_WOW64                               Errno = 0x20000000 | 0xC0000000 | 0x235
	ERROR_SET_SYSTEM_RESTORE_POINT               Errno = 0x20000000 | 0xC0000000 | 0x236
	ERROR_SCE_DISABLED                           Errno = 0x20000000 | 0xC0000000 | 0x238
	ERROR_UNKNOWN_EXCEPTION                      Errno = 0x20000000 | 0xC0000000 | 0x239
	ERROR_PNP_REGISTRY_ERROR                     Errno = 0x20000000 | 0xC0000000 | 0x23A
	ERROR_REMOTE_REQUEST_UNSUPPORTED             Errno = 0x20000000 | 0xC0000000 | 0x23B
	ERROR_NOT_AN_INSTALLED_OEM_INF               Errno = 0x20000000 | 0xC0000000 | 0x23C
	ERROR_INF_IN_USE_BY_DEVICES                  Errno = 0x20000000 | 0xC0000000 | 0x23D
	ERROR_DI_FUNCTION_OBSOLETE                   Errno = 0x20000000 | 0xC0000000 | 0x23E
	ERROR_NO_AUTHENTICODE_CATALOG                Errno = 0x20000000 | 0xC0000000 | 0x23F
	ERROR_AUTHENTICODE_DISALLOWED                Errno = 0x20000000 | 0xC0000000 | 0x240
	ERROR_AUTHENTICODE_TRUSTED_PUBLISHER         Errno = 0x20000000 | 0xC0000000 | 0x241
	ERROR_AUTHENTICODE_TRUST_NOT_ESTABLISHED     Errno = 0x20000000 | 0xC0000000 | 0x242
	ERROR_AUTHENTICODE_PUBLISHER_NOT_TRUSTED     Errno = 0x20000000 | 0xC0000000 | 0x243
	ERROR_SIGNATURE_OSATTRIBUTE_MISMATCH         Errno = 0x20000000 | 0xC0000000 | 0x244
	ERROR_ONLY_VALIDATE_VIA_AUTHENTICODE         Errno = 0x20000000 | 0xC0000000 | 0x245
	ERROR_DEVICE_INSTALLER_NOT_READY             Errno = 0x20000000 | 0xC0000000 | 0x246
	ERROR_DRIVER_STORE_ADD_FAILED                Errno = 0x20000000 | 0xC0000000 | 0x247
	ERROR_DEVICE_INSTALL_BLOCKED                 Errno = 0x20000000 | 0xC0000000 | 0x248
	ERROR_DRIVER_INSTALL_BLOCKED                 Errno = 0x20000000 | 0xC0000000 | 0x249
	ERROR_WRONG_INF_TYPE                         Errno = 0x20000000 | 0xC0000000 | 0x24A
	ERROR_FILE_HASH_NOT_IN_CATALOG               Errno = 0x20000000 | 0xC0000000 | 0x24B
	ERROR_DRIVER_STORE_DELETE_FAILED             Errno = 0x20000000 | 0xC0000000 | 0x24C
	ERROR_UNRECOVERABLE_STACK_OVERFLOW           Errno = 0x20000000 | 0xC0000000 | 0x300
	EXCEPTION_SPAPI_UNRECOVERABLE_STACK_OVERFLOW Errno = ERROR_UNRECOVERABLE_STACK_OVERFLOW
	ERROR_NO_DEFAULT_INTERFACE_DEVICE            Errno = ERROR_NO_DEFAULT_DEVICE_INTERFACE
	ERROR_INTERFACE_DEVICE_ACTIVE                Errno = ERROR_DEVICE_INTERFACE_ACTIVE
	ERROR_INTERFACE_DEVICE_REMOVED               Errno = ERROR_DEVICE_INTERFACE_REMOVED
	ERROR_NO_SUCH_INTERFACE_DEVICE               Errno = ERROR_NO_SUCH_DEVICE_INTERFACE
)

const (
	MAX_DEVICE_ID_LEN   = 200
	MAX_DEVNODE_ID_LEN  = MAX_DEVICE_ID_LEN
	MAX_GUID_STRING_LEN = 39 // 38个字符 + 终止符空字符
	MAX_CLASS_NAME_LEN  = 32
	MAX_PROFILE_LEN     = 80
	MAX_CONFIG_VALUE    = 9999
	MAX_INSTANCE_VALUE  = 9999
	CONFIGMG_VERSION    = 0x0400
)

// 最大字符串长度常量
const (
	LINE_LEN                    = 256  // Windows 9x 兼容的最大可显示字符串长度，该字符串来自设备 INF。
	MAX_INF_STRING_LENGTH       = 4096 // 实际INF字符串（包括字符串替换）的最大尺寸
	MAX_INF_SECTION_NAME_LENGTH = 255  // 为了与Windows 9x兼容，INF节名称应限制为32个字符。
	MAX_TITLE_LEN               = 60
	MAX_INSTRUCTION_LEN         = 256
	MAX_LABEL_LEN               = 30
	MAX_SERVICE_NAME_LEN        = 256
	MAX_SUBTITLE_LEN            = 256
)

const (
	// SP_MAX_MACHINENAME_LENGTH 定义了以 ConfigMgr32 CM_Connect_Machine 接收格式表示的机器名的最大长度（即，“\\\\MachineName\\0”）。
	SP_MAX_MACHINENAME_LENGTH = MAX_PATH + 3
)

// HSPFILEQ 是用于设置文件队列的类型
type HSPFILEQ uintptr

// DevInfo 用于保存设备信息集的引用
type DevInfo Handle

// DEVINST 是一个句柄，通常被 cfgmgr32 API 接口所识别
type DEVINST uint32

// DevInfoData 是一个设备信息结构（引用了作为设备信息集合成员的设备实例）
type DevInfoData struct {
	size      uint32
	ClassGUID GUID
	DevInst   DEVINST
	_         uintptr
}

// DevInfoListDetailData 是一个用于设备信息集详细信息的结构体（用于 SetupDiGetDeviceInfoListDetail，该函数取代了 SetupDiGetDeviceInfoListClass 的功能）。
type DevInfoListDetailData struct {
	size                uint32 // 使用unsafeSizeOf方法
	ClassGUID           GUID
	RemoteMachineHandle Handle
	remoteMachineName   [SP_MAX_MACHINENAME_LENGTH]uint16
}

func (*DevInfoListDetailData) unsafeSizeOf() uint32 {
	if unsafe.Sizeof(uintptr(0)) == 4 {
		// Windows 在 pshpack1.h 中声明了这一点
		return uint32(unsafe.Offsetof(DevInfoListDetailData{}.remoteMachineName) + unsafe.Sizeof(DevInfoListDetailData{}.remoteMachineName))
	}
	return uint32(unsafe.Sizeof(DevInfoListDetailData{}))
}


// ff:
func (data *DevInfoListDetailData) RemoteMachineName() string {
	return UTF16ToString(data.remoteMachineName[:])
}


// ff:
// remoteMachineName:
func (data *DevInfoListDetailData) SetRemoteMachineName(remoteMachineName string) error {
	str, err := UTF16FromString(remoteMachineName)
	if err != nil {
		return err
	}
	copy(data.remoteMachineName[:], str)
	return nil
}

// DI_FUNCTION 是设备安装器的功能函数类型
type DI_FUNCTION uint32

const (
	DIF_SELECTDEVICE                   DI_FUNCTION = 0x00000001
	DIF_INSTALLDEVICE                  DI_FUNCTION = 0x00000002
	DIF_ASSIGNRESOURCES                DI_FUNCTION = 0x00000003
	DIF_PROPERTIES                     DI_FUNCTION = 0x00000004
	DIF_REMOVE                         DI_FUNCTION = 0x00000005
	DIF_FIRSTTIMESETUP                 DI_FUNCTION = 0x00000006
	DIF_FOUNDDEVICE                    DI_FUNCTION = 0x00000007
	DIF_SELECTCLASSDRIVERS             DI_FUNCTION = 0x00000008
	DIF_VALIDATECLASSDRIVERS           DI_FUNCTION = 0x00000009
	DIF_INSTALLCLASSDRIVERS            DI_FUNCTION = 0x0000000A
	DIF_CALCDISKSPACE                  DI_FUNCTION = 0x0000000B
	DIF_DESTROYPRIVATEDATA             DI_FUNCTION = 0x0000000C
	DIF_VALIDATEDRIVER                 DI_FUNCTION = 0x0000000D
	DIF_DETECT                         DI_FUNCTION = 0x0000000F
	DIF_INSTALLWIZARD                  DI_FUNCTION = 0x00000010
	DIF_DESTROYWIZARDDATA              DI_FUNCTION = 0x00000011
	DIF_PROPERTYCHANGE                 DI_FUNCTION = 0x00000012
	DIF_ENABLECLASS                    DI_FUNCTION = 0x00000013
	DIF_DETECTVERIFY                   DI_FUNCTION = 0x00000014
	DIF_INSTALLDEVICEFILES             DI_FUNCTION = 0x00000015
	DIF_UNREMOVE                       DI_FUNCTION = 0x00000016
	DIF_SELECTBESTCOMPATDRV            DI_FUNCTION = 0x00000017
	DIF_ALLOW_INSTALL                  DI_FUNCTION = 0x00000018
	DIF_REGISTERDEVICE                 DI_FUNCTION = 0x00000019
	DIF_NEWDEVICEWIZARD_PRESELECT      DI_FUNCTION = 0x0000001A
	DIF_NEWDEVICEWIZARD_SELECT         DI_FUNCTION = 0x0000001B
	DIF_NEWDEVICEWIZARD_PREANALYZE     DI_FUNCTION = 0x0000001C
	DIF_NEWDEVICEWIZARD_POSTANALYZE    DI_FUNCTION = 0x0000001D
	DIF_NEWDEVICEWIZARD_FINISHINSTALL  DI_FUNCTION = 0x0000001E
	DIF_INSTALLINTERFACES              DI_FUNCTION = 0x00000020
	DIF_DETECTCANCEL                   DI_FUNCTION = 0x00000021
	DIF_REGISTER_COINSTALLERS          DI_FUNCTION = 0x00000022
	DIF_ADDPROPERTYPAGE_ADVANCED       DI_FUNCTION = 0x00000023
	DIF_ADDPROPERTYPAGE_BASIC          DI_FUNCTION = 0x00000024
	DIF_TROUBLESHOOTER                 DI_FUNCTION = 0x00000026
	DIF_POWERMESSAGEWAKE               DI_FUNCTION = 0x00000027
	DIF_ADDREMOTEPROPERTYPAGE_ADVANCED DI_FUNCTION = 0x00000028
	DIF_UPDATEDRIVER_UI                DI_FUNCTION = 0x00000029
	DIF_FINISHINSTALL_ACTION           DI_FUNCTION = 0x0000002A
)

// DevInstallParams 是设备安装参数结构（与特定的设备信息项关联，或全局与设备信息集关联）
type DevInstallParams struct {
	size                     uint32
	Flags                    DI_FLAGS
	FlagsEx                  DI_FLAGSEX
	hwndParent               uintptr
	InstallMsgHandler        uintptr
	InstallMsgHandlerContext uintptr
	FileQueue                HSPFILEQ
	_                        uintptr
	_                        uint32
	driverPath               [MAX_PATH]uint16
}


// ff:
func (params *DevInstallParams) DriverPath() string {
	return UTF16ToString(params.driverPath[:])
}


// ff:
// driverPath:
func (params *DevInstallParams) SetDriverPath(driverPath string) error {
	str, err := UTF16FromString(driverPath)
	if err != nil {
		return err
	}
	copy(params.driverPath[:], str)
	return nil
}

// DI_FLAGS 是 SP_DEVINSTALL_PARAMS.Flags 的值
type DI_FLAGS uint32

const (
	// 用于选择设备的标志
	DI_SHOWOEM       DI_FLAGS = 0x00000001 // 支持“其它…”按钮
	DI_SHOWCOMPAT    DI_FLAGS = 0x00000002 // 显示兼容性列表
	DI_SHOWCLASS     DI_FLAGS = 0x00000004 // show class list
	DI_SHOWALL       DI_FLAGS = 0x00000007 // 同时显示类与兼容列表
	DI_NOVCP         DI_FLAGS = 0x00000008 // don't create a new copy queue--use caller-supplied FileQueue
	DI_DIDCOMPAT     DI_FLAGS = 0x00000010 // 搜索兼容设备
	DI_DIDCLASS      DI_FLAGS = 0x00000020 // 搜索类设备
	DI_AUTOASSIGNRES DI_FLAGS = 0x00000040 // 如果可能，资源不使用UI

	// DiInstallDevice返回的标志，用于指示需要重启/重新启动
	DI_NEEDRESTART DI_FLAGS = 0x00000080 // 重启后生效
	DI_NEEDREBOOT  DI_FLAGS = 0x00000100 // ""

	// 设备安装标志
	DI_NOBROWSE DI_FLAGS = 0x00000200 // InsertDisk中无“浏览…”

	// 由 DiBuildDriverInfoList 设置的标志
	DI_MULTMFGS DI_FLAGS = 0x00000400 // 如果类驱动列表中存在多个制造商，则设置此选项

	// 标志表明设备已禁用
	DI_DISABLED DI_FLAGS = 0x00000800 // Set if device disabled

	// 设备/类属性标志
	DI_GENERALPAGE_ADDED  DI_FLAGS = 0x00001000
	DI_RESOURCEPAGE_ADDED DI_FLAGS = 0x00002000

	// 标志，表示为此设备（或此类）设置属性时引发了更改，因此可能需要更新设备管理器UI。
	DI_PROPERTIES_CHANGE DI_FLAGS = 0x00004000

	// 标志，表示应使用INF文件中的排序。
	DI_INF_IS_SORTED DI_FLAGS = 0x00008000

	// 标志用于指示仅应搜索由 SP_DEVINSTALL_PARAMS.DriverPath 指定的 INF。
	DI_ENUMSINGLEINF DI_FLAGS = 0x00010000

// 防止ConfigMgr在设备注册、安装和删除期间移除或重新枚举设备的标志。
	DI_DONOTCALLCONFIGMG DI_FLAGS = 0x00020000

	// 以下标志可用于安装一个禁用的设备
	DI_INSTALLDISABLED DI_FLAGS = 0x00040000

// 该标志导致 SetupDiBuildDriverInfoList 从设备现有的类驱动程序列表中构建其兼容驱动程序列表，而非进行常规的 INF 搜索。
	DI_COMPAT_FROM_CLASS DI_FLAGS = 0x00080000

	// 如果应使用类安装参数，则设置此标志
	DI_CLASSINSTALLPARAMS DI_FLAGS = 0x00100000

	// 如果调用者不希望在类安装程序返回 ERROR_DI_DO_DEFAULT 时执行内部默认操作，则设置此标志。
	DI_NODI_DEFAULTACTION DI_FLAGS = 0x00200000

	// 设备安装标志
	DI_QUIETINSTALL        DI_FLAGS = 0x00800000 // don't confuse the user with questions or excess info
	DI_NOFILECOPY          DI_FLAGS = 0x01000000 // No file Copy necessary
	DI_FORCECOPY           DI_FLAGS = 0x02000000 // 强制将文件从安装路径复制
	DI_DRIVERPAGE_ADDED    DI_FLAGS = 0x04000000 // Prop提供商已添加Driver页面。
	DI_USECI_SELECTSTRINGS DI_FLAGS = 0x08000000 // 在“选择设备”对话框中使用由类安装程序提供的字符串
	DI_OVERRIDE_INFFLAGS   DI_FLAGS = 0x10000000 // Override INF flags
	DI_PROPS_NOCHANGEUSAGE DI_FLAGS = 0x20000000 // 通用属性中无启用/禁用

	DI_NOSELECTICONS DI_FLAGS = 0x40000000 // 选择设备对话框中不显示小图标

	DI_NOWRITE_IDS DI_FLAGS = 0x80000000 // Don't write HW & Compat IDs on install
)

// DI_FLAGSEX 是 SP_DEVINSTALL_PARAMS 结构体中 FlagsEx 字段的值
type DI_FLAGSEX uint32

const (
	DI_FLAGSEX_CI_FAILED                DI_FLAGSEX = 0x00000004 // 加载/调用类安装程序失败
	DI_FLAGSEX_FINISHINSTALL_ACTION     DI_FLAGSEX = 0x00000008 // 类/共同安装程序希望在客户端上下文中获取一个 DIF_FINISH_INSTALL 动作
	DI_FLAGSEX_DIDINFOLIST              DI_FLAGSEX = 0x00000010 // 是否已获取班级信息列表
	DI_FLAGSEX_DIDCOMPATINFO            DI_FLAGSEX = 0x00000020 // 是否已获取兼容信息列表
	DI_FLAGSEX_FILTERCLASSES            DI_FLAGSEX = 0x00000040
	DI_FLAGSEX_SETFAILEDINSTALL         DI_FLAGSEX = 0x00000080
	DI_FLAGSEX_DEVICECHANGE             DI_FLAGSEX = 0x00000100
	DI_FLAGSEX_ALWAYSWRITEIDS           DI_FLAGSEX = 0x00000200
	DI_FLAGSEX_PROPCHANGE_PENDING       DI_FLAGSEX = 0x00000400 // 一个或多个设备属性表已发生更改，需要触发 DIF_PROPERTYCHANGE 事件。
	DI_FLAGSEX_ALLOWEXCLUDEDDRVS        DI_FLAGSEX = 0x00000800
	DI_FLAGSEX_NOUIONQUERYREMOVE        DI_FLAGSEX = 0x00001000
	DI_FLAGSEX_USECLASSFORCOMPAT        DI_FLAGSEX = 0x00002000 // Use the device's class when building compat drv list. (Ignored if DI_COMPAT_FROM_CLASS flag is specified.)
	DI_FLAGSEX_NO_DRVREG_MODIFY         DI_FLAGSEX = 0x00008000 // 不要为设备的软件（驱动）键运行 AddReg 和 DelReg。
	DI_FLAGSEX_IN_SYSTEM_SETUP          DI_FLAGSEX = 0x00010000 // 正在进行初始系统设置时的安装
	DI_FLAGSEX_INET_DRIVER              DI_FLAGSEX = 0x00020000 // 驱动程序来自Windows更新
	DI_FLAGSEX_APPENDDRIVERLIST         DI_FLAGSEX = 0x00040000 // 令 SetupDiBuildDriverInfoList 将一个新驱动列表追加至现有列表。
	DI_FLAGSEX_PREINSTALLBACKUP         DI_FLAGSEX = 0x00080000 // not used
	DI_FLAGSEX_BACKUPONREPLACE          DI_FLAGSEX = 0x00100000 // not used
	DI_FLAGSEX_DRIVERLIST_FROM_URL      DI_FLAGSEX = 0x00200000 // 从 SP_DEVINSTALL_PARAMS.DriverPath 指定的 URL（空字符串表示 Windows Update 网站）获取的 INF 文件中构建驱动程序列表
	DI_FLAGSEX_EXCLUDE_OLD_INET_DRIVERS DI_FLAGSEX = 0x00800000 // Don't include old Internet drivers when building a driver list. Ignored on Windows Vista and later.
	DI_FLAGSEX_POWERPAGE_ADDED          DI_FLAGSEX = 0x01000000 // 类安装程序添加了它们自己的电源页面
	DI_FLAGSEX_FILTERSIMILARDRIVERS     DI_FLAGSEX = 0x02000000 // 只在类别列表中包含相似的驱动程序
	DI_FLAGSEX_INSTALLEDDRIVER          DI_FLAGSEX = 0x04000000 // 只将已安装的驱动添加到类或兼容驱动列表中。用于对SetupDiBuildDriverInfoList的调用
	DI_FLAGSEX_NO_CLASSLIST_NODE_MERGE  DI_FLAGSEX = 0x08000000 // Don't remove identical driver nodes from the class list
	DI_FLAGSEX_ALTPLATFORM_DRVSEARCH    DI_FLAGSEX = 0x10000000 // 根据关联文件队列中指定的备用平台信息构建驱动列表
	DI_FLAGSEX_RESTART_DEVICE_ONLY      DI_FLAGSEX = 0x20000000 // 仅重启正在安装设备驱动的设备，而非重启所有使用这些驱动的设备
	DI_FLAGSEX_RECURSIVESEARCH          DI_FLAGSEX = 0x40000000 // 告诉 SetupDiBuildDriverInfoList 进行递归搜索
	DI_FLAGSEX_SEARCH_PUBLISHED_INFS    DI_FLAGSEX = 0x80000000 // 告诉SetupDiBuildDriverInfoList执行“已发布INF”搜索
)

// ClassInstallHeader 是任何类安装参数结构的第一个成员。它包含了定义该安装参数结构其余部分格式的设备安装请求代码。
type ClassInstallHeader struct {
	size            uint32
	InstallFunction DI_FUNCTION
}


// ff:
// installFunction:
func MakeClassInstallHeader(installFunction DI_FUNCTION) *ClassInstallHeader {
	hdr := &ClassInstallHeader{InstallFunction: installFunction}
	hdr.size = uint32(unsafe.Sizeof(*hdr))
	return hdr
}

// DICS_STATE 定义了一系列值，用于指示设备状态的变更
type DICS_STATE uint32

const (
	DICS_ENABLE     DICS_STATE = 0x00000001 // 设备正在被启用
	DICS_DISABLE    DICS_STATE = 0x00000002 // 设备正在被禁用。
	DICS_PROPCHANGE DICS_STATE = 0x00000003 // 设备的属性已更改。
	DICS_START      DICS_STATE = 0x00000004 // 设备正在启动（如果该请求针对当前活跃的硬件配置文件）。
	DICS_STOP       DICS_STATE = 0x00000005 // 设备正在被停止。驱动程序堆栈将被卸载，并且为该设备设置CSCONFIGFLAG_DO_NOT_START标志。
)

// DICS_FLAG 指定设备属性更改的范围
type DICS_FLAG uint32

const (
	DICS_FLAG_GLOBAL         DICS_FLAG = 0x00000001 // 在所有硬件配置文件中进行更改
	DICS_FLAG_CONFIGSPECIFIC DICS_FLAG = 0x00000002 // 只在指定配置文件中做出更改
	DICS_FLAG_CONFIGGENERAL  DICS_FLAG = 0x00000004 // 后续跟随 1 个或多个硬件配置特定的变更（已废弃）
)

// PropChangeParams 是与 DIF_PROPERTYCHANGE 安装函数对应的一个结构体。
type PropChangeParams struct {
	ClassInstallHeader ClassInstallHeader
	StateChange        DICS_STATE
	Scope              DICS_FLAG
	HwProfile          uint32
}

// DI_REMOVEDEVICE 指定了设备移除的范围
type DI_REMOVEDEVICE uint32

const (
	DI_REMOVEDEVICE_GLOBAL         DI_REMOVEDEVICE = 0x00000001 // 在所有硬件配置中进行此更改。从注册表中删除该设备的相关信息。
	DI_REMOVEDEVICE_CONFIGSPECIFIC DI_REMOVEDEVICE = 0x00000002 // 对指定由 HwProfile 的硬件配置文件做出此更改。此标志仅适用于根枚举设备。
// 当 Windows 从设备最后被配置的硬件配置文件中移除该设备时，Windows 将执行全局移除操作。
)

// RemoveDeviceParams 是与 DIF_REMOVE 安装函数对应的一个结构体。
type RemoveDeviceParams struct {
	ClassInstallHeader ClassInstallHeader
	Scope              DI_REMOVEDEVICE
	HwProfile          uint32
}

// DrvInfoData 是驱动程序信息结构（属于可能与特定设备实例关联，或（全局）与设备信息集关联的驱动程序信息列表的成员）
type DrvInfoData struct {
	size          uint32
	DriverType    uint32
	_             uintptr
	description   [LINE_LEN]uint16
	mfgName       [LINE_LEN]uint16
	providerName  [LINE_LEN]uint16
	DriverDate    Filetime
	DriverVersion uint64
}


// ff:
func (data *DrvInfoData) Description() string {
	return UTF16ToString(data.description[:])
}


// ff:
// description:
func (data *DrvInfoData) SetDescription(description string) error {
	str, err := UTF16FromString(description)
	if err != nil {
		return err
	}
	copy(data.description[:], str)
	return nil
}


// ff:
func (data *DrvInfoData) MfgName() string {
	return UTF16ToString(data.mfgName[:])
}


// ff:
// mfgName:
func (data *DrvInfoData) SetMfgName(mfgName string) error {
	str, err := UTF16FromString(mfgName)
	if err != nil {
		return err
	}
	copy(data.mfgName[:], str)
	return nil
}


// ff:
func (data *DrvInfoData) ProviderName() string {
	return UTF16ToString(data.providerName[:])
}


// ff:
// providerName:
func (data *DrvInfoData) SetProviderName(providerName string) error {
	str, err := UTF16FromString(providerName)
	if err != nil {
		return err
	}
	copy(data.providerName[:], str)
	return nil
}

// IsNewer 方法返回真，当 DrvInfoData 的日期和版本比提供的参数更新时。

// ff:
// driverVersion:
// driverDate:
func (data *DrvInfoData) IsNewer(driverDate Filetime, driverVersion uint64) bool {
	if data.DriverDate.HighDateTime > driverDate.HighDateTime {
		return true
	}
	if data.DriverDate.HighDateTime < driverDate.HighDateTime {
		return false
	}

	if data.DriverDate.LowDateTime > driverDate.LowDateTime {
		return true
	}
	if data.DriverDate.LowDateTime < driverDate.LowDateTime {
		return false
	}

	if data.DriverVersion > driverVersion {
		return true
	}
	if data.DriverVersion < driverVersion {
		return false
	}

	return false
}

// DrvInfoDetailData 是驾驶员信息详情结构（用于提供关于特定驾驶员信息结构的详细信息）
type DrvInfoDetailData struct {
	size            uint32 // 使用unsafeSizeOf方法
	InfDate         Filetime
	compatIDsOffset uint32
	compatIDsLength uint32
	_               uintptr
	sectionName     [LINE_LEN]uint16
	infFileName     [MAX_PATH]uint16
	drvDescription  [LINE_LEN]uint16
	hardwareID      [1]uint16
}

func (*DrvInfoDetailData) unsafeSizeOf() uint32 {
	if unsafe.Sizeof(uintptr(0)) == 4 {
		// Windows 在 pshpack1.h 中声明了这一点
		return uint32(unsafe.Offsetof(DrvInfoDetailData{}.hardwareID) + unsafe.Sizeof(DrvInfoDetailData{}.hardwareID))
	}
	return uint32(unsafe.Sizeof(DrvInfoDetailData{}))
}


// ff:
func (data *DrvInfoDetailData) SectionName() string {
	return UTF16ToString(data.sectionName[:])
}


// ff:
func (data *DrvInfoDetailData) InfFileName() string {
	return UTF16ToString(data.infFileName[:])
}


// ff:
func (data *DrvInfoDetailData) DrvDescription() string {
	return UTF16ToString(data.drvDescription[:])
}


// ff:
func (data *DrvInfoDetailData) HardwareID() string {
	if data.compatIDsOffset > 1 {
		bufW := data.getBuf()
		return UTF16ToString(bufW[:wcslen(bufW)])
	}

	return ""
}


// ff:
func (data *DrvInfoDetailData) CompatIDs() []string {
	a := make([]string, 0)

	if data.compatIDsLength > 0 {
		bufW := data.getBuf()
		bufW = bufW[data.compatIDsOffset : data.compatIDsOffset+data.compatIDsLength]
		for i := 0; i < len(bufW); {
			j := i + wcslen(bufW[i:])
			if i < j {
				a = append(a, UTF16ToString(bufW[i:j]))
			}
			i = j + 1
		}
	}

	return a
}

func (data *DrvInfoDetailData) getBuf() []uint16 {
	len := (data.size - uint32(unsafe.Offsetof(data.hardwareID))) / 2
	sl := struct {
		addr *uint16
		len  int
		cap  int
	}{&data.hardwareID[0], int(len), int(len)}
	return *(*[]uint16)(unsafe.Pointer(&sl))
}

// IsCompatible 方法用于检测给定的硬件ID是否与驱动程序匹配，或者是否列在兼容ID列表中。

// ff:
// hwid:
func (data *DrvInfoDetailData) IsCompatible(hwid string) bool {
	hwidLC := strings.ToLower(hwid)
	if strings.ToLower(data.HardwareID()) == hwidLC {
		return true
	}
	a := data.CompatIDs()
	for i := range a {
		if strings.ToLower(a[i]) == hwidLC {
			return true
		}
	}

	return false
}

// DICD 标志用于控制 SetupDiCreateDeviceInfo 函数
type DICD uint32

const (
	DICD_GENERATE_ID       DICD = 0x00000001
	DICD_INHERIT_CLASSDRVS DICD = 0x00000002
)

// SUOI 标志用于控制 SetupUninstallOEMInf
type SUOI uint32

const (
	SUOI_FORCEDELETE SUOI = 0x0001
)

// SPDIT标志用于区分类驱动程序和设备驱动程序。
// （作为“DriverType”参数传递给驱动程序信息列表API）
type SPDIT uint32

const (
	SPDIT_NODRIVER     SPDIT = 0x00000000
	SPDIT_CLASSDRIVER  SPDIT = 0x00000001
	SPDIT_COMPATDRIVER SPDIT = 0x00000002
)

// DIGCF 标志用于控制由 SetupDiGetClassDevs 函数构建的设备信息集中所包含的内容
type DIGCF uint32

const (
	DIGCF_DEFAULT         DIGCF = 0x00000001 // 仅在 DIGCF_DEVICEINTERFACE 下有效
	DIGCF_PRESENT         DIGCF = 0x00000002
	DIGCF_ALLCLASSES      DIGCF = 0x00000004
	DIGCF_PROFILE         DIGCF = 0x00000008
	DIGCF_DEVICEINTERFACE DIGCF = 0x00000010
)

// DIREG 指定用于 SetupDiCreateDevRegKey、SetupDiOpenDevRegKey 和 SetupDiDeleteDevRegKey 的值。
type DIREG uint32

const (
	DIREG_DEV  DIREG = 0x00000001 // 打开/创建/删除设备密钥
	DIREG_DRV  DIREG = 0x00000002 // 打开/创建/删除驱动键
	DIREG_BOTH DIREG = 0x00000004 // 删除driver和Device键
)

// SPDRP 定义设备注册表属性代码
// （标注为只读（R）的代码仅可用于 SetupDiGetDeviceRegistryProperty）
//
// 这些值应覆盖由 cfgmgr32.h 中定义的 CM_DRP 代码所表示的相同注册表属性集合。
//
// 注意：SPDRP 代码从零开始编号，而 CM_DRP 代码从一开始编号！
type SPDRP uint32

const (
	SPDRP_DEVICEDESC                  SPDRP = 0x00000000 // DeviceDesc (R/W)
	SPDRP_HARDWAREID                  SPDRP = 0x00000001 // HardwareID (R/W)
	SPDRP_COMPATIBLEIDS               SPDRP = 0x00000002 // CompatibleIDs (R/W)
	SPDRP_SERVICE                     SPDRP = 0x00000004 // Service (R/W)
	SPDRP_CLASS                       SPDRP = 0x00000007 // 类（与ClassGUID关联的R-）
	SPDRP_CLASSGUID                   SPDRP = 0x00000008 // ClassGUID (R/W)
	SPDRP_DRIVER                      SPDRP = 0x00000009 // Driver (R/W)
	SPDRP_CONFIGFLAGS                 SPDRP = 0x0000000A // ConfigFlags (R/W)
	SPDRP_MFG                         SPDRP = 0x0000000B // Mfg (R/W)
	SPDRP_FRIENDLYNAME                SPDRP = 0x0000000C // FriendlyName (R/W)
	SPDRP_LOCATION_INFORMATION        SPDRP = 0x0000000D // LocationInformation（可读/可写）
	SPDRP_PHYSICAL_DEVICE_OBJECT_NAME SPDRP = 0x0000000E // PhysicalDeviceObjectName (R)
	SPDRP_CAPABILITIES                SPDRP = 0x0000000F // Capabilities (R)
	SPDRP_UI_NUMBER                   SPDRP = 0x00000010 // UiNumber (R)
	SPDRP_UPPERFILTERS                SPDRP = 0x00000011 // UpperFilters (R/W)
	SPDRP_LOWERFILTERS                SPDRP = 0x00000012 // LowerFilters (R/W)
	SPDRP_BUSTYPEGUID                 SPDRP = 0x00000013 // BusTypeGUID (R)
	SPDRP_LEGACYBUSTYPE               SPDRP = 0x00000014 // LegacyBusType (R)
	SPDRP_BUSNUMBER                   SPDRP = 0x00000015 // BusNumber (R)
	SPDRP_ENUMERATOR_NAME             SPDRP = 0x00000016 // Enumerator Name (R)
	SPDRP_SECURITY                    SPDRP = 0x00000017 // 安全性（读/写，二进制形式）
	SPDRP_SECURITY_SDS                SPDRP = 0x00000018 // Security (W, SDS form)
	SPDRP_DEVTYPE                     SPDRP = 0x00000019 // Device Type (R/W)
	SPDRP_EXCLUSIVE                   SPDRP = 0x0000001A // 设备为独占访问（读/写）
	SPDRP_CHARACTERISTICS             SPDRP = 0x0000001B // 设备特性（读/写）
	SPDRP_ADDRESS                     SPDRP = 0x0000001C // Device Address (R)
	SPDRP_UI_NUMBER_DESC_FORMAT       SPDRP = 0x0000001D // UiNumberDescFormat (可读/写)
	SPDRP_DEVICE_POWER_DATA           SPDRP = 0x0000001E // Device Power Data (R)
	SPDRP_REMOVAL_POLICY              SPDRP = 0x0000001F // Removal Policy (R)
	SPDRP_REMOVAL_POLICY_HW_DEFAULT   SPDRP = 0x00000020 // 硬件移除策略 (R)
	SPDRP_REMOVAL_POLICY_OVERRIDE     SPDRP = 0x00000021 // 移除策略覆盖（可读写）
	SPDRP_INSTALL_STATE               SPDRP = 0x00000022 // 设备安装状态 (只读)
	SPDRP_LOCATION_PATHS              SPDRP = 0x00000023 // 设备位置路径 (只读)
	SPDRP_BASE_CONTAINERID            SPDRP = 0x00000024 // Base ContainerID (R)

	SPDRP_MAXIMUM_PROPERTY SPDRP = 0x00000025 // 对序数的上限
)

// DEVPROPTYPE 表示统一设备属性模型中设备属性值的数据类型标识符。
type DEVPROPTYPE uint32

const (
	DEVPROP_TYPEMOD_ARRAY DEVPROPTYPE = 0x00001000
	DEVPROP_TYPEMOD_LIST  DEVPROPTYPE = 0x00002000

	DEVPROP_TYPE_EMPTY                      DEVPROPTYPE = 0x00000000
	DEVPROP_TYPE_NULL                       DEVPROPTYPE = 0x00000001
	DEVPROP_TYPE_SBYTE                      DEVPROPTYPE = 0x00000002
	DEVPROP_TYPE_BYTE                       DEVPROPTYPE = 0x00000003
	DEVPROP_TYPE_INT16                      DEVPROPTYPE = 0x00000004
	DEVPROP_TYPE_UINT16                     DEVPROPTYPE = 0x00000005
	DEVPROP_TYPE_INT32                      DEVPROPTYPE = 0x00000006
	DEVPROP_TYPE_UINT32                     DEVPROPTYPE = 0x00000007
	DEVPROP_TYPE_INT64                      DEVPROPTYPE = 0x00000008
	DEVPROP_TYPE_UINT64                     DEVPROPTYPE = 0x00000009
	DEVPROP_TYPE_FLOAT                      DEVPROPTYPE = 0x0000000A
	DEVPROP_TYPE_DOUBLE                     DEVPROPTYPE = 0x0000000B
	DEVPROP_TYPE_DECIMAL                    DEVPROPTYPE = 0x0000000C
	DEVPROP_TYPE_GUID                       DEVPROPTYPE = 0x0000000D
	DEVPROP_TYPE_CURRENCY                   DEVPROPTYPE = 0x0000000E
	DEVPROP_TYPE_DATE                       DEVPROPTYPE = 0x0000000F
	DEVPROP_TYPE_FILETIME                   DEVPROPTYPE = 0x00000010
	DEVPROP_TYPE_BOOLEAN                    DEVPROPTYPE = 0x00000011
	DEVPROP_TYPE_STRING                     DEVPROPTYPE = 0x00000012
	DEVPROP_TYPE_STRING_LIST                DEVPROPTYPE = DEVPROP_TYPE_STRING | DEVPROP_TYPEMOD_LIST
	DEVPROP_TYPE_SECURITY_DESCRIPTOR        DEVPROPTYPE = 0x00000013
	DEVPROP_TYPE_SECURITY_DESCRIPTOR_STRING DEVPROPTYPE = 0x00000014
	DEVPROP_TYPE_DEVPROPKEY                 DEVPROPTYPE = 0x00000015
	DEVPROP_TYPE_DEVPROPTYPE                DEVPROPTYPE = 0x00000016
	DEVPROP_TYPE_BINARY                     DEVPROPTYPE = DEVPROP_TYPE_BYTE | DEVPROP_TYPEMOD_ARRAY
	DEVPROP_TYPE_ERROR                      DEVPROPTYPE = 0x00000017
	DEVPROP_TYPE_NTSTATUS                   DEVPROPTYPE = 0x00000018
	DEVPROP_TYPE_STRING_INDIRECT            DEVPROPTYPE = 0x00000019

	MAX_DEVPROP_TYPE    DEVPROPTYPE = 0x00000019
	MAX_DEVPROP_TYPEMOD DEVPROPTYPE = 0x00002000

	DEVPROP_MASK_TYPE    DEVPROPTYPE = 0x00000FFF
	DEVPROP_MASK_TYPEMOD DEVPROPTYPE = 0x0000F000
)

// DEVPROPGUID 定义了一个属性类别。
type DEVPROPGUID GUID

// DEVPROPID 用于唯一标识属性类别内的属性。
type DEVPROPID uint32

const DEVPROPID_FIRST_USABLE DEVPROPID = 2

// DEVPROPKEY 代表统一设备属性模型中设备属性的设备属性键。
type DEVPROPKEY struct {
	FmtID DEVPROPGUID
	PID   DEVPROPID
}

// CONFIGRET 是 cfgmgr32 API 返回的值或错误代码
type CONFIGRET uint32


// ff:
func (ret CONFIGRET) Error() string {
	if win32Error, ok := ret.Unwrap().(Errno); ok {
		return fmt.Sprintf("%s (CfgMgr error: 0x%08x)", win32Error.Error(), uint32(ret))
	}
	return fmt.Sprintf("CfgMgr error: 0x%08x", uint32(ret))
}


// ff:
// defaultError:
func (ret CONFIGRET) Win32Error(defaultError Errno) Errno {
	return cm_MapCrToWin32Err(ret, defaultError)
}


// ff:
func (ret CONFIGRET) Unwrap() error {
	const noMatch = Errno(^uintptr(0))
	win32Error := ret.Win32Error(noMatch)
	if win32Error == noMatch {
		return nil
	}
	return win32Error
}

const (
	CR_SUCCESS                  CONFIGRET = 0x00000000
	CR_DEFAULT                  CONFIGRET = 0x00000001
	CR_OUT_OF_MEMORY            CONFIGRET = 0x00000002
	CR_INVALID_POINTER          CONFIGRET = 0x00000003
	CR_INVALID_FLAG             CONFIGRET = 0x00000004
	CR_INVALID_DEVNODE          CONFIGRET = 0x00000005
	CR_INVALID_DEVINST                    = CR_INVALID_DEVNODE
	CR_INVALID_RES_DES          CONFIGRET = 0x00000006
	CR_INVALID_LOG_CONF         CONFIGRET = 0x00000007
	CR_INVALID_ARBITRATOR       CONFIGRET = 0x00000008
	CR_INVALID_NODELIST         CONFIGRET = 0x00000009
	CR_DEVNODE_HAS_REQS         CONFIGRET = 0x0000000A
	CR_DEVINST_HAS_REQS                   = CR_DEVNODE_HAS_REQS
	CR_INVALID_RESOURCEID       CONFIGRET = 0x0000000B
	CR_DLVXD_NOT_FOUND          CONFIGRET = 0x0000000C
	CR_NO_SUCH_DEVNODE          CONFIGRET = 0x0000000D
	CR_NO_SUCH_DEVINST                    = CR_NO_SUCH_DEVNODE
	CR_NO_MORE_LOG_CONF         CONFIGRET = 0x0000000E
	CR_NO_MORE_RES_DES          CONFIGRET = 0x0000000F
	CR_ALREADY_SUCH_DEVNODE     CONFIGRET = 0x00000010
	CR_ALREADY_SUCH_DEVINST               = CR_ALREADY_SUCH_DEVNODE
	CR_INVALID_RANGE_LIST       CONFIGRET = 0x00000011
	CR_INVALID_RANGE            CONFIGRET = 0x00000012
	CR_FAILURE                  CONFIGRET = 0x00000013
	CR_NO_SUCH_LOGICAL_DEV      CONFIGRET = 0x00000014
	CR_CREATE_BLOCKED           CONFIGRET = 0x00000015
	CR_NOT_SYSTEM_VM            CONFIGRET = 0x00000016
	CR_REMOVE_VETOED            CONFIGRET = 0x00000017
	CR_APM_VETOED               CONFIGRET = 0x00000018
	CR_INVALID_LOAD_TYPE        CONFIGRET = 0x00000019
	CR_BUFFER_SMALL             CONFIGRET = 0x0000001A
	CR_NO_ARBITRATOR            CONFIGRET = 0x0000001B
	CR_NO_REGISTRY_HANDLE       CONFIGRET = 0x0000001C
	CR_REGISTRY_ERROR           CONFIGRET = 0x0000001D
	CR_INVALID_DEVICE_ID        CONFIGRET = 0x0000001E
	CR_INVALID_DATA             CONFIGRET = 0x0000001F
	CR_INVALID_API              CONFIGRET = 0x00000020
	CR_DEVLOADER_NOT_READY      CONFIGRET = 0x00000021
	CR_NEED_RESTART             CONFIGRET = 0x00000022
	CR_NO_MORE_HW_PROFILES      CONFIGRET = 0x00000023
	CR_DEVICE_NOT_THERE         CONFIGRET = 0x00000024
	CR_NO_SUCH_VALUE            CONFIGRET = 0x00000025
	CR_WRONG_TYPE               CONFIGRET = 0x00000026
	CR_INVALID_PRIORITY         CONFIGRET = 0x00000027
	CR_NOT_DISABLEABLE          CONFIGRET = 0x00000028
	CR_FREE_RESOURCES           CONFIGRET = 0x00000029
	CR_QUERY_VETOED             CONFIGRET = 0x0000002A
	CR_CANT_SHARE_IRQ           CONFIGRET = 0x0000002B
	CR_NO_DEPENDENT             CONFIGRET = 0x0000002C
	CR_SAME_RESOURCES           CONFIGRET = 0x0000002D
	CR_NO_SUCH_REGISTRY_KEY     CONFIGRET = 0x0000002E
	CR_INVALID_MACHINENAME      CONFIGRET = 0x0000002F
	CR_REMOTE_COMM_FAILURE      CONFIGRET = 0x00000030
	CR_MACHINE_UNAVAILABLE      CONFIGRET = 0x00000031
	CR_NO_CM_SERVICES           CONFIGRET = 0x00000032
	CR_ACCESS_DENIED            CONFIGRET = 0x00000033
	CR_CALL_NOT_IMPLEMENTED     CONFIGRET = 0x00000034
	CR_INVALID_PROPERTY         CONFIGRET = 0x00000035
	CR_DEVICE_INTERFACE_ACTIVE  CONFIGRET = 0x00000036
	CR_NO_SUCH_DEVICE_INTERFACE CONFIGRET = 0x00000037
	CR_INVALID_REFERENCE_STRING CONFIGRET = 0x00000038
	CR_INVALID_CONFLICT_LIST    CONFIGRET = 0x00000039
	CR_INVALID_INDEX            CONFIGRET = 0x0000003A
	CR_INVALID_STRUCTURE_SIZE   CONFIGRET = 0x0000003B
	NUM_CR_RESULTS              CONFIGRET = 0x0000003C
)

const (
	CM_GET_DEVICE_INTERFACE_LIST_PRESENT     = 0 // 仅包含当前“活跃”的设备接口
	CM_GET_DEVICE_INTERFACE_LIST_ALL_DEVICES = 1 // 所有已注册的设备接口，无论是否活跃
)

const (
	DN_ROOT_ENUMERATED       = 0x00000001        // Was enumerated by ROOT
	DN_DRIVER_LOADED         = 0x00000002        // 是否已注册设备驱动
	DN_ENUM_LOADED           = 0x00000004        // 具有 Register_Enumerator
	DN_STARTED               = 0x00000008        // 当前已配置
	DN_MANUAL                = 0x00000010        // Manually installed
	DN_NEED_TO_ENUM          = 0x00000020        // May need reenumeration
	DN_NOT_FIRST_TIME        = 0x00000040        // Has received a config
	DN_HARDWARE_ENUM         = 0x00000080        // Enum 生成硬件ID
	DN_LIAR                  = 0x00000100        // 伪造了可重新配置一次的事实
	DN_HAS_MARK              = 0x00000200        // 最近没有调用 CM_Create_DevInst
	DN_HAS_PROBLEM           = 0x00000400        // Need device installer
	DN_FILTERED              = 0x00000800        // Is filtered
	DN_MOVED                 = 0x00001000        // Has been moved
	DN_DISABLEABLE           = 0x00002000        // Can be disabled
	DN_REMOVABLE             = 0x00004000        // Can be removed
	DN_PRIVATE_PROBLEM       = 0x00008000        // Has a private problem
	DN_MF_PARENT             = 0x00010000        // Multi function parent
	DN_MF_CHILD              = 0x00020000        // Multi function child
	DN_WILL_BE_REMOVED       = 0x00040000        // DevInst 正在被移除
	DN_NOT_FIRST_TIMEE       = 0x00080000        // 已收到配置枚举
	DN_STOP_FREE_RES         = 0x00100000        // 当子进程停止时，释放资源
	DN_REBAL_CANDIDATE       = 0x00200000        // Don't skip during rebalance
	DN_BAD_PARTIAL           = 0x00400000        // This devnode's log_confs do not have same resources
	DN_NT_ENUMERATOR         = 0x00800000        // This devnode's is an NT enumerator
	DN_NT_DRIVER             = 0x01000000        // This devnode's is an NT driver
	DN_NEEDS_LOCKING         = 0x02000000        // 需要锁定以进行devnode的恢复处理
	DN_ARM_WAKEUP            = 0x04000000        // Devnode 可作为唤醒设备
	DN_APM_ENUMERATOR        = 0x08000000        // APM aware enumerator
	DN_APM_DRIVER            = 0x10000000        // APM aware driver
	DN_SILENT_INSTALL        = 0x20000000        // Silent install
	DN_NO_SHOW_IN_DM         = 0x40000000        // 在设备管理器中无显示
	DN_BOOT_LOG_PROB         = 0x80000000        // 在预分配启动日志配置时遇到问题
	DN_NEED_RESTART          = DN_LIAR           // 该Devnode正常工作需要系统重启
	DN_DRIVER_BLOCKED        = DN_NOT_FIRST_TIME // 一个或多个驱动程序被阻止为此 Devnode 加载
	DN_LEGACY_DRIVER         = DN_MOVED          // 该设备正在使用一个旧版驱动程序
	DN_CHILD_WITH_INVALID_ID = DN_HAS_MARK       // 一个或多个子元素具有无效ID
	DN_DEVICE_DISCONNECTED   = DN_NEEDS_LOCKING  // 设备的驱动函数报告该设备未连接。通常这意味着无线设备已超出范围。
	DN_QUERY_REMOVE_PENDING  = DN_MF_PARENT      // Device 是一组待查询移除的关联设备集合中的一部分
	DN_QUERY_REMOVE_ACTIVE   = DN_MF_CHILD       // 设备当前正忙于处理一个查询-移除IRP
	DN_CHANGEABLE_FLAGS      = DN_NOT_FIRST_TIME | DN_HARDWARE_ENUM | DN_HAS_MARK | DN_DISABLEABLE | DN_REMOVABLE | DN_MF_CHILD | DN_MF_PARENT | DN_NOT_FIRST_TIMEE | DN_STOP_FREE_RES | DN_REBAL_CANDIDATE | DN_NT_ENUMERATOR | DN_NT_DRIVER | DN_SILENT_INSTALL | DN_NO_SHOW_IN_DM
)

//sys	创建设备信息列表（扩展版），参数如下：
//		classGUID：指向设备类GUID的指针；
//		hwndParent：父窗口句柄（uintptr类型）；
//		machineName：目标计算机名称（以UTF-16编码的uint16类型指针）；
//		reserved：保留参数，目前未使用（uintptr类型）。
// 返回值：
//		handle：成功时返回一个有效的DevInfo句柄，表示新创建的设备信息列表；
//		err：若出错则返回相应的错误信息。
// 备注：
//		当函数失败时，其返回的DevInfo句柄应等于DevInfo(InvalidHandle)。
// 导入符号：
//		本函数直接映射到setupapi库中的SetupDiCreateDeviceInfoListExW函数。

// SetupDiCreateDeviceInfoListEx 函数在远程或本地计算机上创建一个空的设备信息集合，并可选择性地将该集合与设备安装类相关联。

// ff:
// err:
// deviceInfoSet:
// machineName:
// hwndParent:
// classGUID:
func SetupDiCreateDeviceInfoListEx(classGUID *GUID, hwndParent uintptr, machineName string) (deviceInfoSet DevInfo, err error) {
	var machineNameUTF16 *uint16
	if machineName != "" {
		machineNameUTF16, err = UTF16PtrFromString(machineName)
		if err != nil {
			return
		}
	}
	return setupDiCreateDeviceInfoListEx(classGUID, hwndParent, machineNameUTF16, 0)
}

//sys	获取设备信息列表详细信息(deviceInfoSet DevInfo, deviceInfoSetDetailData *DevInfoListDetailData) (错误 error) = setupapi.SetupDiGetDeviceInfoListDetailW

// SetupDiGetDeviceInfoListDetail 函数用于获取与设备信息集关联的信息，包括类 GUID、远程计算机句柄以及远程计算机名称。

// ff:
// err:
// deviceInfoSetDetailData:
// deviceInfoSet:
func SetupDiGetDeviceInfoListDetail(deviceInfoSet DevInfo) (deviceInfoSetDetailData *DevInfoListDetailData, err error) {
	data := &DevInfoListDetailData{}
	data.size = data.unsafeSizeOf()

	return data, setupDiGetDeviceInfoListDetail(deviceInfoSet, data)
}

// DeviceInfoListDetail 方法用于获取与设备信息集关联的信息，包括类 GUID、远程计算机句柄以及远程计算机名。

// ff:
func (deviceInfoSet DevInfo) DeviceInfoListDetail() (*DevInfoListDetailData, error) {
	return SetupDiGetDeviceInfoListDetail(deviceInfoSet)
}

//sys	创建设备信息（setupDiCreateDeviceInfo）用于在指定的deviceInfoSet中，根据给定的DeviceName、classGUID、DeviceDescription、hwndParent和CreationFlags参数，新建一个设备项。函数返回一个错误（err），如果操作成功则为nil。该函数为setupapi包中的SetupDiCreateDeviceInfoW函数的封装。
// 
// 参数说明：
// - deviceInfoSet DevInfo：表示已打开的设备信息集合句柄。
// - DeviceName *uint16：指向设备名称的指针，以UTF-16编码表示。
// - classGUID *GUID：指向设备类别的GUID指针。
// - DeviceDescription *uint16：指向设备描述的指针，以UTF-16编码表示。
// - hwndParent uintptr：父窗口的句柄，通常用于显示与设备安装相关的用户界面。
// - CreationFlags DICD：创建设备信息时使用的标志，如DICD_GENERATE_ID等。
// - deviceInfoData *DevInfoData：输出参数，用于接收新创建设备项的相关信息。
// 
// 该注释定义了一个名为setupDiCreateDeviceInfo的系统调用，它封装了Windows API函数setupapi.SetupDiCreateDeviceInfoW，并提供了Go语言友好的接口。通过调用此函数，开发者可以在指定的设备信息集中添加一个新的设备项，同时指定其名称、类别、描述以及创建选项。函数执行结果通过返回的错误值反映，若成功则返回nil；否则返回具体的错误信息。此外，新创建的设备项详细数据将填充到传入的deviceInfoData参数所指向的结构体中。

// SetupDiCreateDeviceInfo函数用于创建一个新的设备信息项，并将其作为新成员添加到指定的设备信息集中。

// ff:
// err:
// deviceInfoData:
// creationFlags:
// hwndParent:
// deviceDescription:
// classGUID:
// deviceName:
// deviceInfoSet:
func SetupDiCreateDeviceInfo(deviceInfoSet DevInfo, deviceName string, classGUID *GUID, deviceDescription string, hwndParent uintptr, creationFlags DICD) (deviceInfoData *DevInfoData, err error) {
	deviceNameUTF16, err := UTF16PtrFromString(deviceName)
	if err != nil {
		return
	}

	var deviceDescriptionUTF16 *uint16
	if deviceDescription != "" {
		deviceDescriptionUTF16, err = UTF16PtrFromString(deviceDescription)
		if err != nil {
			return
		}
	}

	data := &DevInfoData{}
	data.size = uint32(unsafe.Sizeof(*data))

	return data, setupDiCreateDeviceInfo(deviceInfoSet, deviceNameUTF16, classGUID, deviceDescriptionUTF16, hwndParent, creationFlags, data)
}

// CreateDeviceInfo 方法用于创建一个新的设备信息元素，并将其作为新成员添加到指定的设备信息集中。

// ff:
// creationFlags:
// hwndParent:
// deviceDescription:
// classGUID:
// deviceName:
func (deviceInfoSet DevInfo) CreateDeviceInfo(deviceName string, classGUID *GUID, deviceDescription string, hwndParent uintptr, creationFlags DICD) (*DevInfoData, error) {
	return SetupDiCreateDeviceInfo(deviceInfoSet, deviceName, classGUID, deviceDescription, hwndParent, creationFlags)
}

//sys	枚举设备信息（setupDiEnumDeviceInfo）函数，接受以下参数：
//  - deviceInfoSet：一个DevInfo类型的句柄，表示要从中进行设备枚举的信息集合。
//  - memberIndex：一个uint32类型的整数，指定要获取的设备信息在集合中的索引位置。
//  - deviceInfoData：指向DevInfoData结构体的指针，用于接收所枚举到的设备信息。
//
// 函数调用setupapi库中的SetupDiEnumDeviceInfo函数，并返回一个错误类型（err），表示执行结果。

// SetupDiEnumDeviceInfo 函数返回一个 DevInfoData 结构，该结构用于指定设备信息集中的一项设备信息元素。

// ff:
// memberIndex:
// deviceInfoSet:
func SetupDiEnumDeviceInfo(deviceInfoSet DevInfo, memberIndex int) (*DevInfoData, error) {
	data := &DevInfoData{}
	data.size = uint32(unsafe.Sizeof(*data))

	return data, setupDiEnumDeviceInfo(deviceInfoSet, uint32(memberIndex), data)
}

// EnumDeviceInfo 方法返回一个 DevInfoData 结构，该结构用于指定设备信息集中某个设备信息元素。

// ff:
// memberIndex:
func (deviceInfoSet DevInfo) EnumDeviceInfo(memberIndex int) (*DevInfoData, error) {
	return SetupDiEnumDeviceInfo(deviceInfoSet, memberIndex)
}

// SetupDiDestroyDeviceInfoList 函数用于删除一个设备信息集，并释放与其关联的所有内存。
//sys	SetupDiDestroyDeviceInfoList(deviceInfoSet DevInfo) (err error) = setupapi.SetupDiDestroyDeviceInfoList

// Close方法用于删除一个设备信息集并释放所有相关的内存。

// ff:
func (deviceInfoSet DevInfo) Close() error {
	return SetupDiDestroyDeviceInfoList(deviceInfoSet)
}

//sys	SetupDiBuildDriverInfoList(deviceInfoSet DevInfo, deviceInfoData *DevInfoData, driverType SPDIT) (err error) = setupapi.SetupDiBuildDriverInfoList
// 
// 系统调用 SetupDiBuildDriverInfoList，用于构建设备信息集（deviceInfoSet）中指定设备（deviceInfoData）的驱动程序信息列表。参数说明如下：
//   - deviceInfoSet：类型为 DevInfo 的设备信息集合句柄。
//   - deviceInfoData：指向 DevInfoData 结构体的指针，该结构体包含了要为其构建驱动程序信息列表的设备相关信息。
//   - driverType：枚举类型 SPDIT，表示要构建的驱动程序信息类型。
// 
// 函数返回值：
//   - err：表示函数执行过程中是否发生错误。如果没有错误，返回 nil；否则，返回具体的错误信息。
// 
// 本函数为对 Windows SDK 中 setupapi 库中的 SetupDiBuildDriverInfoList 函数的封装。

// BuildDriverInfoList 方法用于为特定设备或设备信息集中全局类驱动列表构建关联的驱动程序列表。

// ff:
// driverType:
// deviceInfoData:
func (deviceInfoSet DevInfo) BuildDriverInfoList(deviceInfoData *DevInfoData, driverType SPDIT) error {
	return SetupDiBuildDriverInfoList(deviceInfoSet, deviceInfoData, driverType)
}

//sys	SetupDiCancelDriverInfoSearch(deviceInfoSet DevInfo) (err error) = setupapi.SetupDiCancelDriverInfoSearch
// 
// 系统调用 SetupDiCancelDriverInfoSearch，接收一个 deviceInfoSet 类型的参数 DevInfo，并返回一个错误 err。该函数等同于使用 setupapi 库中的 SetupDiCancelDriverInfoSearch 函数。

// CancelDriverInfoSearch 方法用于取消在另一个线程中正在进行的驾驶员信息列表搜索。

// ff:
func (deviceInfoSet DevInfo) CancelDriverInfoSearch() error {
	return SetupDiCancelDriverInfoSearch(deviceInfoSet)
}

//sys	枚举设备驱动信息(设备信息集 deviceInfoSet, 设备信息数据指针 deviceInfoData, 驱动类型 driverType, 成员索引 memberIndex, 驱动信息数据指针 driverInfoData) (返回错误 err) = setupapi.SetupDiEnumDriverInfoW

// SetupDiEnumDriverInfo 函数用于枚举驱动列表中的成员。

// ff:
// memberIndex:
// driverType:
// deviceInfoData:
// deviceInfoSet:
func SetupDiEnumDriverInfo(deviceInfoSet DevInfo, deviceInfoData *DevInfoData, driverType SPDIT, memberIndex int) (*DrvInfoData, error) {
	data := &DrvInfoData{}
	data.size = uint32(unsafe.Sizeof(*data))

	return data, setupDiEnumDriverInfo(deviceInfoSet, deviceInfoData, driverType, uint32(memberIndex), data)
}

// EnumDriverInfo 方法枚举驱动程序列表中的成员。

// ff:
// memberIndex:
// driverType:
// deviceInfoData:
func (deviceInfoSet DevInfo) EnumDriverInfo(deviceInfoData *DevInfoData, driverType SPDIT, memberIndex int) (*DrvInfoData, error) {
	return SetupDiEnumDriverInfo(deviceInfoSet, deviceInfoData, driverType, memberIndex)
}

//sys	设置Di获取选中的驱动程序（deviceInfoSet DevInfo, deviceInfoData *DevInfoData, driverInfoData *DrvInfoData）(返回错误 err) = setupapi.SetupDiGetSelectedDriverW

// SetupDiGetSelectedDriver 函数用于为设备信息集或特定设备信息元素检索所选驱动程序。

// ff:
// deviceInfoData:
// deviceInfoSet:
func SetupDiGetSelectedDriver(deviceInfoSet DevInfo, deviceInfoData *DevInfoData) (*DrvInfoData, error) {
	data := &DrvInfoData{}
	data.size = uint32(unsafe.Sizeof(*data))

	return data, setupDiGetSelectedDriver(deviceInfoSet, deviceInfoData, data)
}

// SelectedDriver 方法用于获取设备信息集或特定设备信息项中已选择的驱动程序。

// ff:
// deviceInfoData:
func (deviceInfoSet DevInfo) SelectedDriver(deviceInfoData *DevInfoData) (*DrvInfoData, error) {
	return SetupDiGetSelectedDriver(deviceInfoSet, deviceInfoData)
}

//sys	SetupDiSetSelectedDriver(deviceInfoSet DevInfo, deviceInfoData *DevInfoData, driverInfoData *DrvInfoData) (err error) = setupapi.SetupDiSetSelectedDriverW
// 
// 系统调用 SetupDiSetSelectedDriver 用于为指定的设备信息集（deviceInfoSet）中的设备（由 deviceInfoData 指针指向其相关信息）设置选中的驱动程序。该驱动程序的相关信息由 driverInfoData 指针所指向。此函数为 Windows API 中 setupapi 库的一部分，具体对应于名为 SetupDiSetSelectedDriverW 的函数。函数执行结果返回一个错误值（err），用于表示操作是否成功。

// SetSelectedDriver 方法用于设置或重置设备信息元素的选定驱动程序，或为设备信息集设置选定的类驱动程序。

// ff:
// driverInfoData:
// deviceInfoData:
func (deviceInfoSet DevInfo) SetSelectedDriver(deviceInfoData *DevInfoData, driverInfoData *DrvInfoData) error {
	return SetupDiSetSelectedDriver(deviceInfoSet, deviceInfoData, driverInfoData)
}

//sys	设置Di获取驱动信息详细（deviceInfoSet DevInfo, deviceInfoData *DevInfoData, driverInfoData *DrvInfoData, driverInfoDetailData *DrvInfoDetailData, driverInfoDetailDataSize uint32, requiredSize *uint32）(err error) = setupapi.SetupDiGetDriverInfoDetailW

// SetupDiGetDriverInfoDetail 函数用于为设备信息集中或该集合中特定设备信息元素获取驱动程序详细信息。

// ff:
// driverInfoData:
// deviceInfoData:
// deviceInfoSet:
func SetupDiGetDriverInfoDetail(deviceInfoSet DevInfo, deviceInfoData *DevInfoData, driverInfoData *DrvInfoData) (*DrvInfoDetailData, error) {
	reqSize := uint32(2048)
	for {
		buf := make([]byte, reqSize)
		data := (*DrvInfoDetailData)(unsafe.Pointer(&buf[0]))
		data.size = data.unsafeSizeOf()
		err := setupDiGetDriverInfoDetail(deviceInfoSet, deviceInfoData, driverInfoData, data, uint32(len(buf)), &reqSize)
		if err == ERROR_INSUFFICIENT_BUFFER {
			continue
		}
		if err != nil {
			return nil, err
		}
		data.size = reqSize
		return data, nil
	}
}

// DriverInfoDetail 方法用于检索设备信息集中某个设备信息项的驱动程序详细信息，或者为设备信息集中的特定设备信息元素获取驱动程序信息。

// ff:
// driverInfoData:
// deviceInfoData:
func (deviceInfoSet DevInfo) DriverInfoDetail(deviceInfoData *DevInfoData, driverInfoData *DrvInfoData) (*DrvInfoDetailData, error) {
	return SetupDiGetDriverInfoDetail(deviceInfoSet, deviceInfoData, driverInfoData)
}

//sys	SetupDiDestroyDriverInfoList(deviceInfoSet DevInfo, deviceInfoData *DevInfoData, driverType SPDIT) (err error) = setupapi.SetupDiDestroyDriverInfoList
// 
// 系统调用 SetupDiDestroyDriverInfoList，用于销毁与指定设备信息集（deviceInfoSet）关联的驱动程序信息列表。参数说明如下：
//   - deviceInfoSet：类型为 DevInfo，表示需要销毁其驱动程序信息列表的设备信息集。
//   - deviceInfoData：类型为 *DevInfoData，指向包含设备特定配置信息的结构体指针。此参数用于确定要销毁的驱动程序信息列表。
//   - driverType：类型为 SPDIT，定义了要销毁的驱动程序信息类型的枚举值。
// 
// 此函数映射到名为 setupapi.SetupDiDestroyDriverInfoList 的 WinAPI 函数，执行后返回一个错误（err），用于指示操作是否成功完成。

// DestroyDriverInfoList 方法用于删除一个驱动列表。

// ff:
// driverType:
// deviceInfoData:
func (deviceInfoSet DevInfo) DestroyDriverInfoList(deviceInfoData *DevInfoData, driverType SPDIT) error {
	return SetupDiDestroyDriverInfoList(deviceInfoSet, deviceInfoData, driverType)
}

//sys	setupDiGetClassDevsEx(classGUID *GUID, Enumerator *uint16, hwndParent uintptr, Flags DIGCF, deviceInfoSet DevInfo, machineName *uint16, reserved uintptr) (handle DevInfo, err error) [failretval==DevInfo(InvalidHandle)] = setupapi.SetupDiGetClassDevsExW
// 
// 系统调用 setupDiGetClassDevsEx，参数如下：
//   - classGUID：指向 GUID 类型的指针，表示设备类或接口类的唯一标识。
//   - Enumerator：指向 uint16 类型的指针，包含枚举器的名称（可选）。
//   - hwndParent：窗口句柄（uintptr 类型），指定父窗口（可选）。
//   - Flags：DIGCF 类型的标志位，用于指定获取设备信息集的方式。
//   - deviceInfoSet：DevInfo 类型，表示设备信息集合的句柄（可选，通常为零以创建新的集合）。
//   - machineName：指向 uint16 类型的指针，包含目标计算机的名称（可选，本地系统默认）。
//   - reserved：保留参数（uintptr 类型），必须为零。
// 
// 函数返回：
//   - handle：DevInfo 类型，成功时返回设备信息集合的有效句柄。
//   - err：error 类型，若操作失败则返回相应的错误。
// 
// 当函数返回值（handle）等于 DevInfo(InvalidHandle) 时，表示操作失败。本系统调用对应于 Windows API 函数 `setupapi.SetupDiGetClassDevsExW`。

// SetupDiGetClassDevsEx函数返回一个设备信息集的句柄，该集合包含了针对本地或远程计算机请求的设备信息元素。

// ff:
// err:
// handle:
// machineName:
// deviceInfoSet:
// flags:
// hwndParent:
// enumerator:
// classGUID:
func SetupDiGetClassDevsEx(classGUID *GUID, enumerator string, hwndParent uintptr, flags DIGCF, deviceInfoSet DevInfo, machineName string) (handle DevInfo, err error) {
	var enumeratorUTF16 *uint16
	if enumerator != "" {
		enumeratorUTF16, err = UTF16PtrFromString(enumerator)
		if err != nil {
			return
		}
	}
	var machineNameUTF16 *uint16
	if machineName != "" {
		machineNameUTF16, err = UTF16PtrFromString(machineName)
		if err != nil {
			return
		}
	}
	return setupDiGetClassDevsEx(classGUID, enumeratorUTF16, hwndParent, flags, deviceInfoSet, machineNameUTF16, 0)
}

// SetupDiCallClassInstaller 函数以指定的安装请求（DIF 代码）调用相应类安装程序及所有已注册的协同安装程序。
//sys	SetupDiCallClassInstaller(installFunction DI_FUNCTION, deviceInfoSet DevInfo, deviceInfoData *DevInfoData) (err error) = setupapi.SetupDiCallClassInstaller

// CallClassInstaller 成员调用相应类安装程序及已注册的所有共同安装程序，同时指定了安装请求（DIF 代码）。

// ff:
// deviceInfoData:
// installFunction:
func (deviceInfoSet DevInfo) CallClassInstaller(installFunction DI_FUNCTION, deviceInfoData *DevInfoData) error {
	return SetupDiCallClassInstaller(installFunction, deviceInfoSet, deviceInfoData)
}

// SetupDiOpenDevRegKey 函数用于打开一个注册表键，以便访问设备特定的配置信息。
//sys	SetupDiOpenDevRegKey(deviceInfoSet DevInfo, deviceInfoData *DevInfoData, Scope DICS_FLAG, HwProfile uint32, KeyType DIREG, samDesired uint32) (key Handle, err error) [failretval==InvalidHandle] = setupapi.SetupDiOpenDevRegKey
// 
// 注释翻译如下：
// 
// **SetupDiOpenDevRegKey** 函数用于打开一个与设备相关的注册表键，以获取设备特定的配置信息。
// 
// **函数参数说明**：
// - **deviceInfoSet**：一个 DevInfo 类型的变量，表示设备信息集。
// - **deviceInfoData**：指向 DevInfoData 结构体的指针，包含特定设备的相关数据。
// - **Scope**：DICS_FLAG 类型的枚举值，指定设备配置信息的范围（如系统范围、当前硬件配置等）。
// - **HwProfile**：无符号整型值，表示硬件配置索引。
// - **KeyType**：DIREG 类型的枚举值，指定要打开的注册表键类型（如设备类键、设备实例键等）。
// - **samDesired**：无符号整型值，定义了对所请求注册表键的访问权限。
// 
// **函数返回值**：
// - **key**：返回一个 Handle 类型的变量，表示已成功打开的注册表键句柄。
// - **err**：返回一个错误变量。如果没有错误发生，返回 nil；否则，返回具体的错误信息。
// 
// **附加说明**：
// [failretval==InvalidHandle] 表示当函数执行失败时，返回的 key 值应为 InvalidHandle。此函数在内部调用了名为 `setupapi.SetupDiOpenDevRegKey` 的系统库函数。

// OpenDevRegKey 方法用于打开一个注册表键，以便访问设备特定的配置信息。

// ff:
// Handle:
// samDesired:
// KeyType:
// HwProfile:
// Scope:
// DeviceInfoData:
func (deviceInfoSet DevInfo) OpenDevRegKey(DeviceInfoData *DevInfoData, Scope DICS_FLAG, HwProfile uint32, KeyType DIREG, samDesired uint32) (Handle, error) {
	return SetupDiOpenDevRegKey(deviceInfoSet, DeviceInfoData, Scope, HwProfile, KeyType, samDesired)
}

//sys	setupDiGetDeviceProperty(deviceInfoSet DevInfo, deviceInfoData *DevInfoData, propertyKey *DEVPROPKEY, propertyType *DEVPROPTYPE, propertyBuffer *byte, propertyBufferSize uint32, requiredSize *uint32, flags uint32) (err error) = setupapi.SetupDiGetDevicePropertyW
// 
// 翻译成中文为：
// 
//sys	通过setupapi.SetupDiGetDevicePropertyW调用实现setupDiGetDeviceProperty函数，其参数如下：
//  deviceInfoSet：DevInfo类型的设备信息集；
//  deviceInfoData：指向DevInfoData结构体的指针，包含设备信息数据；
//  propertyKey：指向DEVPROPKEY结构体的指针，表示要获取的设备属性键；
//  propertyType：指向DEVPROPTYPE变量的指针，用于接收所获取属性的数据类型；
//  propertyBuffer：指向字节型缓冲区的指针，用于存储获取到的属性值；
//  propertyBufferSize：uint32类型的变量，表示propertyBuffer的大小（以字节为单位）；
//  requiredSize：指向uint32变量的指针，用于接收实际所需的缓冲区大小（若propertyBuffer太小，不足以容纳属性值，则返回所需大小）；
//  flags：uint32类型的标志位，指定获取属性时的附加选项；
//  返回值：返回一个错误对象(err error)，表示执行过程中的错误状态。

// SetupDiGetDeviceProperty 函数用于检索指定设备实例的属性。

// ff:
// err:
// value:
// propertyKey:
// deviceInfoData:
// deviceInfoSet:
func SetupDiGetDeviceProperty(deviceInfoSet DevInfo, deviceInfoData *DevInfoData, propertyKey *DEVPROPKEY) (value interface{}, err error) {
	reqSize := uint32(256)
	for {
		var dataType DEVPROPTYPE
		buf := make([]byte, reqSize)
		err = setupDiGetDeviceProperty(deviceInfoSet, deviceInfoData, propertyKey, &dataType, &buf[0], uint32(len(buf)), &reqSize, 0)
		if err == ERROR_INSUFFICIENT_BUFFER {
			continue
		}
		if err != nil {
			return
		}
		switch dataType {
		case DEVPROP_TYPE_STRING:
			ret := UTF16ToString(bufToUTF16(buf))
			runtime.KeepAlive(buf)
			return ret, nil
		}
		return nil, errors.New("unimplemented property type")
	}
}

//sys	设置设备注册表属性（setupDiGetDeviceRegistryProperty）函数，用于获取设备信息集（deviceInfoSet）中指定设备（deviceInfoData）的特定注册表属性（property）。参数包括：
//	1. deviceInfoSet：设备信息集句柄。
//	2. deviceInfoData：指向包含目标设备信息的DevInfoData结构指针。
//	3. property：指定要查询的设备注册表属性（SPDRP枚举类型值）。
//	4. propertyRegDataType：指向一个uint32变量的指针，用于接收所查询属性的数据类型。
//	5. propertyBuffer：指向缓冲区的指针，用于接收查询到的属性值。若为nil，则仅计算所需缓冲区大小。
//	6. propertyBufferSize：提供的属性值缓冲区大小，以字节为单位。
//	7. requiredSize：指向一个uint32变量的指针，用于接收查询属性值所需的最小缓冲区大小。
// 
// 函数返回值：
//	1. err：表示执行结果的错误信息。若无错误发生，返回nil；否则返回具体错误。
// 
// 本函数是封装对Windows系统API函数`setupapi.SetupDiGetDeviceRegistryPropertyW`的调用。

// SetupDiGetDeviceRegistryProperty 函数用于检索指定即插即用设备的属性。

// ff:
// err:
// value:
// property:
// deviceInfoData:
// deviceInfoSet:
func SetupDiGetDeviceRegistryProperty(deviceInfoSet DevInfo, deviceInfoData *DevInfoData, property SPDRP) (value interface{}, err error) {
	reqSize := uint32(256)
	for {
		var dataType uint32
		buf := make([]byte, reqSize)
		err = setupDiGetDeviceRegistryProperty(deviceInfoSet, deviceInfoData, property, &dataType, &buf[0], uint32(len(buf)), &reqSize)
		if err == ERROR_INSUFFICIENT_BUFFER {
			continue
		}
		if err != nil {
			return
		}
		return getRegistryValue(buf[:reqSize], dataType)
	}
}

func getRegistryValue(buf []byte, dataType uint32) (interface{}, error) {
	switch dataType {
	case REG_SZ:
		ret := UTF16ToString(bufToUTF16(buf))
		runtime.KeepAlive(buf)
		return ret, nil
	case REG_EXPAND_SZ:
		value := UTF16ToString(bufToUTF16(buf))
		if value == "" {
			return "", nil
		}
		p, err := syscall.UTF16PtrFromString(value)
		if err != nil {
			return "", err
		}
		ret := make([]uint16, 100)
		for {
			n, err := ExpandEnvironmentStrings(p, &ret[0], uint32(len(ret)))
			if err != nil {
				return "", err
			}
			if n <= uint32(len(ret)) {
				return UTF16ToString(ret[:n]), nil
			}
			ret = make([]uint16, n)
		}
	case REG_BINARY:
		return buf, nil
	case REG_DWORD_LITTLE_ENDIAN:
		return binary.LittleEndian.Uint32(buf), nil
	case REG_DWORD_BIG_ENDIAN:
		return binary.BigEndian.Uint32(buf), nil
	case REG_MULTI_SZ:
		bufW := bufToUTF16(buf)
		a := []string{}
		for i := 0; i < len(bufW); {
			j := i + wcslen(bufW[i:])
			if i < j {
				a = append(a, UTF16ToString(bufW[i:j]))
			}
			i = j + 1
		}
		runtime.KeepAlive(buf)
		return a, nil
	case REG_QWORD_LITTLE_ENDIAN:
		return binary.LittleEndian.Uint64(buf), nil
	default:
		return nil, fmt.Errorf("Unsupported registry value type: %v", dataType)
	}
}

// bufToUTF16 函数将 []byte 缓冲区重新解释为 []uint16
func bufToUTF16(buf []byte) []uint16 {
	sl := struct {
		addr *uint16
		len  int
		cap  int
	}{(*uint16)(unsafe.Pointer(&buf[0])), len(buf) / 2, cap(buf) / 2}
	return *(*[]uint16)(unsafe.Pointer(&sl))
}

// utf16ToBuf 函数将 []uint16 重新解释为 []byte
func utf16ToBuf(buf []uint16) []byte {
	sl := struct {
		addr *byte
		len  int
		cap  int
	}{(*byte)(unsafe.Pointer(&buf[0])), len(buf) * 2, cap(buf) * 2}
	return *(*[]byte)(unsafe.Pointer(&sl))
}

func wcslen(str []uint16) int {
	for i := 0; i < len(str); i++ {
		if str[i] == 0 {
			return i
		}
	}
	return len(str)
}

// DeviceRegistryProperty 方法用于检索指定的即插即用设备属性。

// ff:
// property:
// deviceInfoData:
func (deviceInfoSet DevInfo) DeviceRegistryProperty(deviceInfoData *DevInfoData, property SPDRP) (interface{}, error) {
	return SetupDiGetDeviceRegistryProperty(deviceInfoSet, deviceInfoData, property)
}

//sys	setupDiSetDeviceRegistryProperty(deviceInfoSet DevInfo, deviceInfoData *DevInfoData, property SPDRP, propertyBuffer *byte, propertyBufferSize uint32) (err error) = setupapi.SetupDiSetDeviceRegistryPropertyW
// 
// 系统调用setupDiSetDeviceRegistryProperty，用于设置设备注册表属性。参数说明如下：
//   deviceInfoSet：一个DevInfo类型的句柄，标识设备信息集。
//   deviceInfoData：指向DevInfoData结构体的指针，包含要设置其注册表属性的设备相关信息。
//   property：SPDRP类型的枚举值，指定要设置的设备注册表属性。
//   propertyBuffer：指向字节数据的指针，提供待写入注册表的属性值。
//   propertyBufferSize：uint32类型，表示propertyBuffer指向的数据大小（以字节为单位）。
// 函数返回一个错误接口(err error)，用于报告执行结果。此系统调用对应于Windows API函数setupapi.SetupDiSetDeviceRegistryPropertyW。

// SetupDiSetDeviceRegistryProperty 函数用于为设备设置即插即用设备属性。

// ff:
// propertyBuffers:
// property:
// deviceInfoData:
// deviceInfoSet:
func SetupDiSetDeviceRegistryProperty(deviceInfoSet DevInfo, deviceInfoData *DevInfoData, property SPDRP, propertyBuffers []byte) error {
	return setupDiSetDeviceRegistryProperty(deviceInfoSet, deviceInfoData, property, &propertyBuffers[0], uint32(len(propertyBuffers)))
}

// SetDeviceRegistryProperty 函数用于为设备设置即插即用设备属性。

// ff:
// propertyBuffers:
// property:
// deviceInfoData:
func (deviceInfoSet DevInfo) SetDeviceRegistryProperty(deviceInfoData *DevInfoData, property SPDRP, propertyBuffers []byte) error {
	return SetupDiSetDeviceRegistryProperty(deviceInfoSet, deviceInfoData, property, propertyBuffers)
}

// SetDeviceRegistryPropertyString 方法为设备设置即插即用设备属性字符串。

// ff:
// str:
// property:
// deviceInfoData:
func (deviceInfoSet DevInfo) SetDeviceRegistryPropertyString(deviceInfoData *DevInfoData, property SPDRP, str string) error {
	str16, err := UTF16FromString(str)
	if err != nil {
		return err
	}
	err = SetupDiSetDeviceRegistryProperty(deviceInfoSet, deviceInfoData, property, utf16ToBuf(append(str16, 0)))
	runtime.KeepAlive(str16)
	return err
}

//sys	获取设备安装参数（setupDiGetDeviceInstallParams）函数通过deviceInfoSet（设备信息集）、deviceInfoData（设备信息数据指针）和deviceInstallParams（设备安装参数指针）三个参数，从setupapi库中调用SetupDiGetDeviceInstallParamsW方法，返回操作结果的错误信息（err）。

// SetupDiGetDeviceInstallParams 函数用于为一个设备信息集或特定的设备信息元素检索设备安装参数。

// ff:
// deviceInfoData:
// deviceInfoSet:
func SetupDiGetDeviceInstallParams(deviceInfoSet DevInfo, deviceInfoData *DevInfoData) (*DevInstallParams, error) {
	params := &DevInstallParams{}
	params.size = uint32(unsafe.Sizeof(*params))

	return params, setupDiGetDeviceInstallParams(deviceInfoSet, deviceInfoData, params)
}

// DeviceInstallParams 方法用于获取设备信息集或特定设备信息元素的设备安装参数。

// ff:
// deviceInfoData:
func (deviceInfoSet DevInfo) DeviceInstallParams(deviceInfoData *DevInfoData) (*DevInstallParams, error) {
	return SetupDiGetDeviceInstallParams(deviceInfoSet, deviceInfoData)
}

//sys	获取设备实例ID（setupDiGetDeviceInstanceId）函数，接收如下参数：
//		deviceInfoSet：一个DevInfo类型的设备信息集合；
//		deviceInfoData：指向DevInfoData结构体的指针，用于存储设备相关信息；
//		instanceId：指向uint16类型的指针，用于接收设备实例ID字符串；
//		instanceIdSize：一个uint32类型的变量，表示instanceId缓冲区的大小（以字节为单位）；
//		instanceIdRequiredSize：指向uint32类型的指针，用于返回实际所需的instanceId缓冲区大小。
//此函数是封装了setupapi库中的SetupDiGetDeviceInstanceIdW函数，并可能返回一个错误（err）。

// SetupDiGetDeviceInstanceId 函数用于获取设备的实例ID。

// ff:
// deviceInfoData:
// deviceInfoSet:
func SetupDiGetDeviceInstanceId(deviceInfoSet DevInfo, deviceInfoData *DevInfoData) (string, error) {
	reqSize := uint32(1024)
	for {
		buf := make([]uint16, reqSize)
		err := setupDiGetDeviceInstanceId(deviceInfoSet, deviceInfoData, &buf[0], uint32(len(buf)), &reqSize)
		if err == ERROR_INSUFFICIENT_BUFFER {
			continue
		}
		if err != nil {
			return "", err
		}
		return UTF16ToString(buf), nil
	}
}

// DeviceInstanceID 方法用于获取设备的实例ID。

// ff:
// deviceInfoData:
func (deviceInfoSet DevInfo) DeviceInstanceID(deviceInfoData *DevInfoData) (string, error) {
	return SetupDiGetDeviceInstanceId(deviceInfoSet, deviceInfoData)
}

// SetupDiGetClassInstallParams 函数用于为设备信息集或特定设备信息元素检索类安装参数。
//sys	SetupDiGetClassInstallParams(deviceInfoSet DevInfo, deviceInfoData *DevInfoData, classInstallParams *ClassInstallHeader, classInstallParamsSize uint32, requiredSize *uint32) (err error) = setupapi.SetupDiGetClassInstallParamsW

// ClassInstallParams 方法用于为设备信息集或特定设备信息元素检索类安装参数。

// ff:
// requiredSize:
// classInstallParamsSize:
// classInstallParams:
// deviceInfoData:
func (deviceInfoSet DevInfo) ClassInstallParams(deviceInfoData *DevInfoData, classInstallParams *ClassInstallHeader, classInstallParamsSize uint32, requiredSize *uint32) error {
	return SetupDiGetClassInstallParams(deviceInfoSet, deviceInfoData, classInstallParams, classInstallParamsSize, requiredSize)
}

//sys	SetupDiSetDeviceInstallParams(deviceInfoSet DevInfo, deviceInfoData *DevInfoData, deviceInstallParams *DevInstallParams) (err error) = setupapi.SetupDiSetDeviceInstallParamsW
// 
// 系统调用 SetupDiSetDeviceInstallParams 用于设置设备安装参数。具体说明如下：
// 
// 参数说明：
//   - deviceInfoSet（DevInfo 类型）：表示设备信息集合，它包含了待操作的设备相关信息。
//   - deviceInfoData (*DevInfoData)：指向一个 DevInfoData 结构体的指针，该结构体包含了目标设备的具体信息。
//   - deviceInstallParams (*DevInstallParams)：指向一个 DevInstallParams 结构体的指针，该结构体中包含要为设备设置的安装参数。
// 
// 返回值：
//   - err (error)：如果函数执行成功，返回 nil；否则，返回描述错误情况的对象。
// 
// 本系统调用对应于 Windows API 中的 "setupapi.SetupDiSetDeviceInstallParamsW" 函数，用于在 Unicode 编码环境下设置设备安装参数。

// SetDeviceInstallParams 成员为设备信息集或特定设备信息元素设置设备安装参数。

// ff:
// deviceInstallParams:
// deviceInfoData:
func (deviceInfoSet DevInfo) SetDeviceInstallParams(deviceInfoData *DevInfoData, deviceInstallParams *DevInstallParams) error {
	return SetupDiSetDeviceInstallParams(deviceInfoSet, deviceInfoData, deviceInstallParams)
}

// SetupDiSetClassInstallParams 函数用于为设备信息集或特定设备信息元素设置或清除类安装参数。
//sys	SetupDiSetClassInstallParams(deviceInfoSet DevInfo, deviceInfoData *DevInfoData, classInstallParams *ClassInstallHeader, classInstallParamsSize uint32) (err error) = setupapi.SetupDiSetClassInstallParamsW

// SetClassInstallParams 方法用于为一个设备信息集或特定的设备信息元素设置或清除类安装参数。

// ff:
// classInstallParamsSize:
// classInstallParams:
// deviceInfoData:
func (deviceInfoSet DevInfo) SetClassInstallParams(deviceInfoData *DevInfoData, classInstallParams *ClassInstallHeader, classInstallParamsSize uint32) error {
	return SetupDiSetClassInstallParams(deviceInfoSet, deviceInfoData, classInstallParams, classInstallParamsSize)
}

//sys	根据给定的类GUID、目标类名缓冲区、类名缓冲区大小、返回所需缓冲区大小指针、目标计算机名和预留参数，使用setupapi.SetupDiClassNameFromGuidExW函数获取类名，并返回可能的错误。

// SetupDiClassNameFromGuidEx 函数用于获取与类GUID关联的类名。此类可以安装在本地或远程计算机上。

// ff:
// err:
// className:
// machineName:
// classGUID:
func SetupDiClassNameFromGuidEx(classGUID *GUID, machineName string) (className string, err error) {
	var classNameUTF16 [MAX_CLASS_NAME_LEN]uint16

	var machineNameUTF16 *uint16
	if machineName != "" {
		machineNameUTF16, err = UTF16PtrFromString(machineName)
		if err != nil {
			return
		}
	}

	err = setupDiClassNameFromGuidEx(classGUID, &classNameUTF16[0], MAX_CLASS_NAME_LEN, nil, machineNameUTF16, 0)
	if err != nil {
		return
	}

	className = UTF16ToString(classNameUTF16[:])
	return
}

//sys	根据给定的类名（className）在指定计算机（machineName）上获取类GUID列表，将结果存储于classGuidList中。classGuidListSize参数指示classGuidList的容量（以GUID为单位）。requiredSize用于返回实际需要的GUID数量。若machineName为空，表示在本地计算机查询。reserved参数保留供将来使用。此函数为setupapi库中的SetupDiClassGuidsFromNameExW函数的封装，并在出错时返回错误信息（err）。

// SetupDiClassGuidsFromNameEx 函数从指定的类名中检索关联的 GUID。返回的列表包含了当前安装在本地或远程计算机上的此类别。

// ff:
// machineName:
// className:
func SetupDiClassGuidsFromNameEx(className string, machineName string) ([]GUID, error) {
	classNameUTF16, err := UTF16PtrFromString(className)
	if err != nil {
		return nil, err
	}

	var machineNameUTF16 *uint16
	if machineName != "" {
		machineNameUTF16, err = UTF16PtrFromString(machineName)
		if err != nil {
			return nil, err
		}
	}

	reqSize := uint32(4)
	for {
		buf := make([]GUID, reqSize)
		err = setupDiClassGuidsFromNameEx(classNameUTF16, &buf[0], uint32(len(buf)), &reqSize, machineNameUTF16, 0)
		if err == ERROR_INSUFFICIENT_BUFFER {
			continue
		}
		if err != nil {
			return nil, err
		}
		return buf[:reqSize], nil
	}
}

//sys	setupDiGetSelectedDevice(deviceInfoSet DevInfo, deviceInfoData *DevInfoData) (err error) = setupapi.SetupDiGetSelectedDevice
// 
// 系统调用 setupDiGetSelectedDevice，用于获取指定设备信息集合（deviceInfoSet）中已被选中的设备信息。参数说明如下：
//   - deviceInfoSet：类型为 DevInfo 的变量，表示要从中获取选中设备的设备信息集合。
//   - deviceInfoData：指向 DevInfoData 结构体的指针，用于接收所选设备的相关信息。
// 
// 该函数返回一个错误值（err），若操作成功则返回 nil；否则返回具体错误。该函数与 Windows API 函数 `setupapi.SetupDiGetSelectedDevice` 对应。

// SetupDiGetSelectedDevice 函数用于从设备信息集中检索选定的设备信息元素。

// ff:
// deviceInfoSet:
func SetupDiGetSelectedDevice(deviceInfoSet DevInfo) (*DevInfoData, error) {
	data := &DevInfoData{}
	data.size = uint32(unsafe.Sizeof(*data))

	return data, setupDiGetSelectedDevice(deviceInfoSet, data)
}

// SelectedDevice 方法从设备信息集中检索选定的设备信息项。

// ff:
func (deviceInfoSet DevInfo) SelectedDevice() (*DevInfoData, error) {
	return SetupDiGetSelectedDevice(deviceInfoSet)
}

// SetupDiSetSelectedDevice 函数将一个设备信息元素设置为设备信息集中被选中的成员。此函数通常由安装向导使用。
//sys	SetupDiSetSelectedDevice(deviceInfoSet DevInfo, deviceInfoData *DevInfoData) (err error) = setupapi.SetupDiSetSelectedDevice

// SetSelectedDevice 方法将一个设备信息元素设置为设备信息集的选定成员。此函数通常由安装向导使用。

// ff:
// deviceInfoData:
func (deviceInfoSet DevInfo) SetSelectedDevice(deviceInfoData *DevInfoData) error {
	return SetupDiSetSelectedDevice(deviceInfoSet, deviceInfoData)
}

//sys	setupUninstallOEMInf(infFileName *uint16, flags SUOI, reserved uintptr) (err error) = setupapi.SetupUninstallOEMInfW
// 
// 系统调用 setupUninstallOEMInf，用于卸载OEM信息文件。参数说明如下：
// - infFileName: 指向包含OEM信息文件名的 uint16 类型指针。
// - flags: 类型为 SUOI 的标志位，用于指定卸载操作的相关选项或行为。
// - reserved: 类型为 uintptr 的保留参数，目前未使用，应传入零值。
// 该函数等同于调用 setupapi 库中的 SetupUninstallOEMInfW 函数。

// SetupUninstallOEMInf 卸载指定的驱动程序

// ff:
// flags:
// infFileName:
func SetupUninstallOEMInf(infFileName string, flags SUOI) error {
	infFileName16, err := UTF16PtrFromString(infFileName)
	if err != nil {
		return err
	}
	return setupUninstallOEMInf(infFileName16, flags, 0)
}

//sys cm_MapCrToWin32Err(configRet CONFIGRET, defaultWin32Error Errno) (ret Errno) = CfgMgr32.CM_MapCrToWin32Err
// 
// 系统调用 cm_MapCrToWin32Err，将 CONFIGRET 类型的 configRet 映射到对应的 Win32 错误码。如果映射失败，使用 Errno 类型的 defaultWin32Error 作为默认返回值。该函数等价于调用 CfgMgr32 库中的 CM_MapCrToWin32Err 函数，并返回 Errno 类型的结果 ret。

//sys cm_Get_Device_Interface_List_Size(len *uint32, interfaceClass *GUID, deviceID *uint16, flags uint32) (ret CONFIGRET) = CfgMgr32.CM_Get_Device_Interface_List_SizeW
//sys cm_Get_Device_Interface_List(interfaceClass *GUID, deviceID *uint16, buffer *uint16, bufferLen uint32, flags uint32) (ret CONFIGRET) = CfgMgr32.CM_Get_Device_Interface_ListW
// 
// 翻译成中文为：
// 
//sys cm_Get_Device_Interface_List_Size(len *uint32, interfaceClass *GUID, deviceID *uint16, flags uint32) (ret CONFIGRET) = CfgMgr32.CM_Get_Device_Interface_List_SizeW
//sys cm_Get_Device_Interface_List(interfaceClass *GUID, deviceID *uint16, buffer *uint16, bufferLen uint32, flags uint32) (ret CONFIGRET) = CfgMgr32.CM_Get_Device_Interface_ListW
// 
// （注：以上为原文，由于原文为Go语言的系统调用声明，其内容已为英文注释形式，无需再进行翻译。）


// ff:
// flags:
// interfaceClass:
// deviceID:
func CM_Get_Device_Interface_List(deviceID string, interfaceClass *GUID, flags uint32) ([]string, error) {
	deviceID16, err := UTF16PtrFromString(deviceID)
	if err != nil {
		return nil, err
	}
	var buf []uint16
	var buflen uint32
	for {
		if ret := cm_Get_Device_Interface_List_Size(&buflen, interfaceClass, deviceID16, flags); ret != CR_SUCCESS {
			return nil, ret
		}
		buf = make([]uint16, buflen)
		if ret := cm_Get_Device_Interface_List(interfaceClass, deviceID16, &buf[0], buflen, flags); ret == CR_SUCCESS {
			break
		} else if ret != CR_BUFFER_SMALL {
			return nil, ret
		}
	}
	var interfaces []string
	for i := 0; i < len(buf); {
		j := i + wcslen(buf[i:])
		if i < j {
			interfaces = append(interfaces, UTF16ToString(buf[i:j]))
		}
		i = j + 1
	}
	if interfaces == nil {
		return nil, ERROR_NO_SUCH_DEVICE_INTERFACE
	}
	return interfaces, nil
}

//sys cm_Get_DevNode_Status(status *uint32, problemNumber *uint32, devInst DEVINST, flags uint32) (ret CONFIGRET) = CfgMgr32.CM_Get_DevNode_Status
// 
// 系统调用 cm_Get_DevNode_Status，其参数如下：
//   - status：指向一个 uint32 类型的指针，用于接收设备节点的状态信息。
//   - problemNumber：指向一个 uint32 类型的指针，用于接收设备问题编号。
//   - devInst：DEVINST 类型的设备实例句柄。
//   - flags：uint32 类型的标志位。
// 
// 该函数调用 CfgMgr32 库中的 CM_Get_DevNode_Status 函数，并返回 CONFIGRET 类型的结果。
// 
// 注释翻译成中文：
// 
//sys cm_Get_DevNode_Status(status *uint32, problemNumber *uint32, devInst DEVINST, flags uint32) (ret CONFIGRET) = CfgMgr32.CM_Get_DevNode_Status
// 
// 系统调用 cm_Get_DevNode_Status，其参数说明如下：
//   - status：指向一个 uint32 类型变量的指针，用于存储获取到的设备节点状态信息。
//   - problemNumber：指向一个 uint32 类型变量的指针，用于存储设备的问题编号。
//   - devInst：表示设备实例的 DEVINST 类型句柄。
//   - flags：uint32 类型的标志位，用于指定额外的操作选项或模式。
// 
// 该函数实际调用了 CfgMgr32 库中的 CM_Get_DevNode_Status 函数，并返回一个 CONFIGRET 类型的值作为执行结果。


// ff:
// flags:
// devInst:
// problemNumber:
// status:
func CM_Get_DevNode_Status(status *uint32, problemNumber *uint32, devInst DEVINST, flags uint32) error {
	ret := cm_Get_DevNode_Status(status, problemNumber, devInst, flags)
	if ret == CR_SUCCESS {
		return nil
	}
	return ret
}
