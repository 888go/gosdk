// 版权所有 ? 2021 The Go 作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可协议约束，
// 该许可协议可在 LICENSE 文件中找到。

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

// 本文件包含对SetupAPI.dll与CfgMgr32.dll的封装函数，
// 这两个核心系统库负责硬件设备、驱动程序及PnP树的管理。
// 关于这些API的详细信息可参见：
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
	MAX_GUID_STRING_LEN = 39 // 38个字符 + 终止空字符
	MAX_CLASS_NAME_LEN  = 32
	MAX_PROFILE_LEN     = 80
	MAX_CONFIG_VALUE    = 9999
	MAX_INSTANCE_VALUE  = 9999
	CONFIGMG_VERSION    = 0x0400
)

// 最大字符串长度常量
const (
	LINE_LEN                    = 256  // 适用于从设备INF获取的可显示字符串的Windows 9x兼容最大值
	MAX_INF_STRING_LENGTH       = 4096 // 实际INF字符串（包括字符串替换）的最大尺寸。
	MAX_INF_SECTION_NAME_LENGTH = 255  // 为了兼容Windows 9x系统，INF章节名称应限制为32个字符。
	MAX_TITLE_LEN               = 60
	MAX_INSTRUCTION_LEN         = 256
	MAX_LABEL_LEN               = 30
	MAX_SERVICE_NAME_LEN        = 256
	MAX_SUBTITLE_LEN            = 256
)

const (
	// SP_MAX_MACHINENAME_LENGTH 定义了 ConfigMgr32 中 CM_Connect_Machine 期望的机器名格式（即，“\\\\MachineName\\0”）的最大长度。
	SP_MAX_MACHINENAME_LENGTH = MAX_PATH + 3
)

// HSPFILEQ 是用于设置文件队列的类型
type HSPFILEQ uintptr

// DevInfo 保存对设备信息集的引用
type DevInfo Handle

// DEVINST 是一个通常被 cfgmgr32 API 接口所识别的句柄
type DEVINST uint32

// DevInfoData 是一个设备信息结构（引用了作为设备信息集成员的设备实例）
type DevInfoData struct {
	size      uint32
	ClassGUID GUID
	DevInst   DEVINST
	_         uintptr
}

// DevInfoListDetailData 是一种用于设备信息集详细信息的结构（用于 SetupDiGetDeviceInfoListDetail，该函数取代了 SetupDiGetDeviceInfoListClass 的功能）。
type DevInfoListDetailData struct {
	size                uint32 // 使用 unsafeSizeOf 方法
	ClassGUID           GUID
	RemoteMachineHandle Handle
	remoteMachineName   [SP_MAX_MACHINENAME_LENGTH]uint16
}

func (*DevInfoListDetailData) unsafeSizeOf() uint32 {
	if unsafe.Sizeof(uintptr(0)) == 4 {
		// Windows 在 pshpack1.h 中声明了此内容
		return uint32(unsafe.Offsetof(DevInfoListDetailData{}.remoteMachineName) + unsafe.Sizeof(DevInfoListDetailData{}.remoteMachineName))
	}
	return uint32(unsafe.Sizeof(DevInfoListDetailData{}))
}

// 翻译提示:func  (data  *设备信息列表详细数据)  远程机器名称()  字符串  {}

// ff:
func (data *DevInfoListDetailData) RemoteMachineName() string {
	return UTF16ToString(data.remoteMachineName[:])
}

// 翻译提示:func  (data  *设备信息列表详细数据)  设置远程机器名(remoteMachineName  字符串)  错误  {}

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

// 翻译提示:const  (
// 	选择设备  DI_FUNCTION  =  0x00000001
// 	安装设备  DI_FUNCTION  =  0x00000002
// 	分配资源  DI_FUNCTION  =  0x00000003
// 	属性设置  DI_FUNCTION  =  0x00000004
// 	移除设备  DI_FUNCTION  =  0x00000005
// 	首次设置  DI_FUNCTION  =  0x00000006
// 	发现设备  DI_FUNCTION  =  0x00000007
// 	选择类驱动  DI_FUNCTION  =  0x00000008
// 	验证类驱动  DI_FUNCTION  =  0x00000009
// 	安装类驱动  DI_FUNCTION  =  0x0000000A
// 	计算磁盘空间  DI_FUNCTION  =  0x0000000B
// 	销毁私有数据  DI_FUNCTION  =  0x0000000C
// 	验证驱动  DI_FUNCTION  =  0x0000000D
// 	检测设备  DI_FUNCTION  =  0x0000000F
// 	安装向导  DI_FUNCTION  =  0x00000010
// 	销毁向导数据  DI_FUNCTION  =  0x00000011
// 	属性更改  DI_FUNCTION  =  0x00000012
// 	启用类  DI_FUNCTION  =  0x00000013
// 	检测验证  DI_FUNCTION  =  0x00000014
// 	安装设备文件  DI_FUNCTION  =  0x00000015
// 	取消移除  DI_FUNCTION  =  0x00000016
// 	选择最佳兼容驱动  DI_FUNCTION  =  0x00000017
// 	允许安装  DI_FUNCTION  =  0x00000018
// 	注册设备  DI_FUNCTION  =  0x00000019
// 	新设备向导预选  DI_FUNCTION  =  0x0000001A
// 	新设备向导选择  DI_FUNCTION  =  0x0000001B
// 	新设备向导预分析  DI_FUNCTION  =  0x0000001C
// 	新设备向导后分析  DI_FUNCTION  =  0x0000001D
// 	新设备向导完成安装  DI_FUNCTION  =  0x0000001E
// 	安装接口  DI_FUNCTION  =  0x00000020
// 	检测取消  DI_FUNCTION  =  0x00000021
// 	注册共同安装程序  DI_FUNCTION  =  0x00000022
// 	添加高级属性页  DI_FUNCTION  =  0x00000023
// 	添加基本属性页  DI_FUNCTION  =  0x00000024
// 	故障排除器  DI_FUNCTION  =  0x00000026
// 	电源唤醒消息  DI_FUNCTION  =  0x00000027
// 	添加远程高级属性页  DI_FUNCTION  =  0x00000028
// 	更新驱动程序UI  DI_FUNCTION  =  0x00000029
// 	完成安装操作  DI_FUNCTION  =  0x0000002A
// )
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

// DevInstallParams 是设备安装参数结构（与特定的设备信息元素关联，或全局与设备信息集合关联）
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

// 翻译提示:func  (params  *设备安装参数)  驱动程序路径()  string  {}

// ff:
func (params *DevInstallParams) DriverPath() string {
	return UTF16ToString(params.driverPath[:])
}

// 翻译提示:func  (params  *设备安装参数)  设置驱动程序路径(driverPath  字符串)  错误  {}

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
	DI_SHOWALL       DI_FLAGS = 0x00000007 // 显示类与兼容列表
	DI_NOVCP         DI_FLAGS = 0x00000008 // don't create a new copy queue--use caller-supplied FileQueue
	DI_DIDCOMPAT     DI_FLAGS = 0x00000010 // 搜索兼容设备
	DI_DIDCLASS      DI_FLAGS = 0x00000020 // 搜索类设备
	DI_AUTOASSIGNRES DI_FLAGS = 0x00000040 // 尽可能避免为资源提供UI

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

	// 标志用于指示应使用INF文件中的排序。
	DI_INF_IS_SORTED DI_FLAGS = 0x00008000

	// 标志用于指示仅应搜索由 SP_DEVINSTALL_PARAMS.DriverPath 指定的 INF
	DI_ENUMSINGLEINF DI_FLAGS = 0x00010000

// 阻止ConfigMgr在设备注册、安装和删除期间移除/重新枚举设备的标志。
	DI_DONOTCALLCONFIGMG DI_FLAGS = 0x00020000

	// 以下标志可用于安装一个禁用状态的设备
	DI_INSTALLDISABLED DI_FLAGS = 0x00040000

// 该标志导致 SetupDiBuildDriverInfoList 从设备现有的类驱动程序列表中构建其兼容驱动程序列表，而非进行常规的 INF 搜索。
	DI_COMPAT_FROM_CLASS DI_FLAGS = 0x00080000

	// 如果应使用类安装参数，则设置此标志。
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
	DI_FLAGSEX_FINISHINSTALL_ACTION     DI_FLAGSEX = 0x00000008 // 类/共同安装程序希望在客户端上下文中获取 DIF_FINISH_INSTALL 动作
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
	DI_FLAGSEX_NO_DRVREG_MODIFY         DI_FLAGSEX = 0x00008000 // 不为设备的软件（驱动）键运行 AddReg 和 DelReg。
	DI_FLAGSEX_IN_SYSTEM_SETUP          DI_FLAGSEX = 0x00010000 // 正在进行初始系统设置期间的安装
	DI_FLAGSEX_INET_DRIVER              DI_FLAGSEX = 0x00020000 // 驱动程序来自Windows Update
	DI_FLAGSEX_APPENDDRIVERLIST         DI_FLAGSEX = 0x00040000 // 令SetupDiBuildDriverInfoList将一个新驱动列表追加至已存在的列表。
	DI_FLAGSEX_PREINSTALLBACKUP         DI_FLAGSEX = 0x00080000 // not used
	DI_FLAGSEX_BACKUPONREPLACE          DI_FLAGSEX = 0x00100000 // not used
	DI_FLAGSEX_DRIVERLIST_FROM_URL      DI_FLAGSEX = 0x00200000 // 从 SP_DEVINSTALL_PARAMS.DriverPath 指定的 URL（空字符串表示 Windows Update 网站）获取的 INF 文件中构建驱动程序列表
	DI_FLAGSEX_EXCLUDE_OLD_INET_DRIVERS DI_FLAGSEX = 0x00800000 // Don't include old Internet drivers when building a driver list. Ignored on Windows Vista and later.
	DI_FLAGSEX_POWERPAGE_ADDED          DI_FLAGSEX = 0x01000000 // 类安装程序添加了他们自己的电源页面
	DI_FLAGSEX_FILTERSIMILARDRIVERS     DI_FLAGSEX = 0x02000000 // 只在类别列表中包含相似的驱动程序
	DI_FLAGSEX_INSTALLEDDRIVER          DI_FLAGSEX = 0x04000000 // 只将已安装的驱动添加到类或兼容驱动列表中。用于对 SetupDiBuildDriverInfoList 的调用
	DI_FLAGSEX_NO_CLASSLIST_NODE_MERGE  DI_FLAGSEX = 0x08000000 // Don't remove identical driver nodes from the class list
	DI_FLAGSEX_ALTPLATFORM_DRVSEARCH    DI_FLAGSEX = 0x10000000 // 根据关联文件队列中指定的备用平台信息构建驱动列表
	DI_FLAGSEX_RESTART_DEVICE_ONLY      DI_FLAGSEX = 0x20000000 // 仅重启正在安装设备驱动的设备，而非重启所有使用这些驱动的设备
	DI_FLAGSEX_RECURSIVESEARCH          DI_FLAGSEX = 0x40000000 // 告诉 SetupDiBuildDriverInfoList 执行递归搜索
	DI_FLAGSEX_SEARCH_PUBLISHED_INFS    DI_FLAGSEX = 0x80000000 // 告诉SetupDiBuildDriverInfoList执行“已发布INF”搜索
)

// ClassInstallHeader 是任何类安装参数结构体的第一个成员。它包含了定义该安装参数结构体剩余部分格式的设备安装请求代码。
type ClassInstallHeader struct {
	size            uint32
	InstallFunction DI_FUNCTION
}

// 翻译提示:func  创建类安装头安装函数(DI_FUNCTION  安装功能)  *类安装头  {}

// ff:
// installFunction:
func MakeClassInstallHeader(installFunction DI_FUNCTION) *ClassInstallHeader {
	hdr := &ClassInstallHeader{InstallFunction: installFunction}
	hdr.size = uint32(unsafe.Sizeof(*hdr))
	return hdr
}

// DICS_STATE 定义了表示设备状态变化的值
type DICS_STATE uint32

const (
	DICS_ENABLE     DICS_STATE = 0x00000001 // 设备正在被启用
	DICS_DISABLE    DICS_STATE = 0x00000002 // 设备正在被禁用
	DICS_PROPCHANGE DICS_STATE = 0x00000003 // 设备的属性已更改。
	DICS_START      DICS_STATE = 0x00000004 // 设备正在启动（如果该请求针对当前活跃的硬件配置文件）。
	DICS_STOP       DICS_STATE = 0x00000005 // 设备正在被停止。驱动程序堆栈将被卸载，并且为该设备设置CSCONFIGFLAG_DO_NOT_START标志。
)

// DICS_FLAG 指定设备属性更改的范围
type DICS_FLAG uint32

const (
	DICS_FLAG_GLOBAL         DICS_FLAG = 0x00000001 // 在所有硬件配置中进行更改
	DICS_FLAG_CONFIGSPECIFIC DICS_FLAG = 0x00000002 // 只在指定配置文件中进行更改
	DICS_FLAG_CONFIGGENERAL  DICS_FLAG = 0x00000004 // 后续跟随 1 个或多个硬件配置文件特定的更改（已废弃）
)

// PropChangeParams 是与 DIF_PROPERTYCHANGE 安装函数对应的结构体。
type PropChangeParams struct {
	ClassInstallHeader ClassInstallHeader
	StateChange        DICS_STATE
	Scope              DICS_FLAG
	HwProfile          uint32
}

// DI_REMOVEDEVICE 指定设备移除的范围
type DI_REMOVEDEVICE uint32

const (
	DI_REMOVEDEVICE_GLOBAL         DI_REMOVEDEVICE = 0x00000001 // 在所有硬件配置中进行此更改。从注册表中移除该设备的相关信息。
	DI_REMOVEDEVICE_CONFIGSPECIFIC DI_REMOVEDEVICE = 0x00000002 // 对指定由 HwProfile 的硬件配置文件执行此更改。此标志仅适用于根枚举设备。当 Windows 从设备所配置的最后一个硬件配置文件中移除该设备时，Windows 将执行全局移除。
)

// RemoveDeviceParams 是一种与 DIF_REMOVE 安装函数相对应的结构。
type RemoveDeviceParams struct {
	ClassInstallHeader ClassInstallHeader
	Scope              DI_REMOVEDEVICE
	HwProfile          uint32
}

// DrvInfoData 是驱动程序信息结构（可能与特定设备实例关联，或（全局）与设备信息集关联的驱动程序信息列表的成员）
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

// 翻译提示:func  (data  *DrvInfoData)  描述()  string  {}

// ff:
func (data *DrvInfoData) Description() string {
	return UTF16ToString(data.description[:])
}

// 翻译提示:func  (data  *DrvInfoData)  设置描述(description  string)  error  {}

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

// 翻译提示:func  (data  *DrvInfoData)  制造商名称()  string  {}

// ff:
func (data *DrvInfoData) MfgName() string {
	return UTF16ToString(data.mfgName[:])
}

// 翻译提示:func  (数据  *DrvInfoData)  设置制造商名称(mfgName  字符串)  error  {}

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

// 翻译提示:func  (data  *DrvInfoData)  驱动程序提供者名称()  string  {}

// ff:
func (data *DrvInfoData) ProviderName() string {
	return UTF16ToString(data.providerName[:])
}

// 翻译提示:func  (data  *DrvInfoData)  设置提供者名称(providerName  string)  error  {}

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
// driverDate:
// driverVersion:
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
	size            uint32 // 使用 unsafeSizeOf 方法
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
		// Windows 在 pshpack1.h 中声明了此内容
		return uint32(unsafe.Offsetof(DrvInfoDetailData{}.hardwareID) + unsafe.Sizeof(DrvInfoDetailData{}.hardwareID))
	}
	return uint32(unsafe.Sizeof(DrvInfoDetailData{}))
}

// 翻译提示:func  (data  *DrvInfoDetailData)  部分名称()  string  {}

// ff:
func (data *DrvInfoDetailData) SectionName() string {
	return UTF16ToString(data.sectionName[:])
}

// 翻译提示:func  (data  *DrvInfoDetailData)  驱动程序信息文件名()  string  {}

// ff:
func (data *DrvInfoDetailData) InfFileName() string {
	return UTF16ToString(data.infFileName[:])
}

// 翻译提示:func  (data  *DrvInfoDetailData)  驱动描述()  string  {}

// ff:
func (data *DrvInfoDetailData) DrvDescription() string {
	return UTF16ToString(data.drvDescription[:])
}

// 翻译提示:func  (data  *DrvInfoDetailData)  硬件ID()  string  {}

// ff:
func (data *DrvInfoDetailData) HardwareID() string {
	if data.compatIDsOffset > 1 {
		bufW := data.getBuf()
		return UTF16ToString(bufW[:wcslen(bufW)])
	}

	return ""
}

// 翻译提示:func  (data  *DrvInfoDetailData)  兼容ID()  []string  {}

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

// IsCompatible 方法用于检测给定的硬件ID是否与驱动程序匹配，或者是否出现在兼容ID列表中。

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

// 翻译提示:const  (
// 	生成ID                        DICD  =  0x00000001
// 	继承类驱动器            DICD  =  0x00000002
// )
const (
	DICD_GENERATE_ID       DICD = 0x00000001
	DICD_INHERIT_CLASSDRVS DICD = 0x00000002
)

// SUOI 标志控制 SetupUninstallOEMInf
type SUOI uint32

// 翻译提示:const  (
// 	强制删除SUOI  =  0x0001
// )
const (
	SUOI_FORCEDELETE SUOI = 0x0001
)

// SPDIT标志用于区分类驱动程序与设备驱动程序。
// （作为驱动程序信息列表API中“DriverType”参数传入）
type SPDIT uint32

// 翻译提示:const  (
// 	无驱动器更新  SPDIT  =  0x00000000
// 	类驱动器更新  SPDIT  =  0x00000001
// 	兼容驱动器更新  SPDIT  =  0x00000002
// )
const (
	SPDIT_NODRIVER     SPDIT = 0x00000000
	SPDIT_CLASSDRIVER  SPDIT = 0x00000001
	SPDIT_COMPATDRIVER SPDIT = 0x00000002
)

// DIGCF 标志控制由 SetupDiGetClassDevs 创建的设备信息集中所包含的内容
type DIGCF uint32

const (
	DIGCF_DEFAULT         DIGCF = 0x00000001 // 仅在 DIGCF_DEVICEINTERFACE 下有效
	DIGCF_PRESENT         DIGCF = 0x00000002
	DIGCF_ALLCLASSES      DIGCF = 0x00000004
	DIGCF_PROFILE         DIGCF = 0x00000008
	DIGCF_DEVICEINTERFACE DIGCF = 0x00000010
)

// DIREG 指定 SetupDiCreateDevRegKey、SetupDiOpenDevRegKey 和 SetupDiDeleteDevRegKey 函数的值。
type DIREG uint32

const (
	DIREG_DEV  DIREG = 0x00000001 // 打开/创建/删除设备密钥
	DIREG_DRV  DIREG = 0x00000002 // 打开/创建/删除驱动键
	DIREG_BOTH DIREG = 0x00000004 // 删除driver和Device键
)

// SPDRP 定义设备注册表属性代码
// （标记为只读（R）的代码只能用于
// SetupDiGetDeviceRegistryProperty）
//
// 这些值应涵盖由 cfgmgr32.h 中定义的 CM_DRP 代码所表示的相同注册表属性集合。
//
// 注意，SPDRP 代码基于零，而 CM_DRP 代码基于一！
type SPDRP uint32

const (
	SPDRP_DEVICEDESC                  SPDRP = 0x00000000 // DeviceDesc (R/W)
	SPDRP_HARDWAREID                  SPDRP = 0x00000001 // HardwareID (R/W)
	SPDRP_COMPATIBLEIDS               SPDRP = 0x00000002 // CompatibleIDs (R/W)
	SPDRP_SERVICE                     SPDRP = 0x00000004 // Service (R/W)
	SPDRP_CLASS                       SPDRP = 0x00000007 // 类（与ClassGUID关联的R--）
	SPDRP_CLASSGUID                   SPDRP = 0x00000008 // ClassGUID (R/W)
	SPDRP_DRIVER                      SPDRP = 0x00000009 // Driver (R/W)
	SPDRP_CONFIGFLAGS                 SPDRP = 0x0000000A // ConfigFlags (R/W)
	SPDRP_MFG                         SPDRP = 0x0000000B // Mfg (R/W)
	SPDRP_FRIENDLYNAME                SPDRP = 0x0000000C // FriendlyName (R/W)
	SPDRP_LOCATION_INFORMATION        SPDRP = 0x0000000D // LocationInformation (可读/可写)
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
	SPDRP_EXCLUSIVE                   SPDRP = 0x0000001A // Device为独占访问（读/写）
	SPDRP_CHARACTERISTICS             SPDRP = 0x0000001B // 设备特性（读/写）
	SPDRP_ADDRESS                     SPDRP = 0x0000001C // Device Address (R)
	SPDRP_UI_NUMBER_DESC_FORMAT       SPDRP = 0x0000001D // UiNumberDescFormat (R/W)
// 
// UiNumberDescFormat 是一个可读写属性
	SPDRP_DEVICE_POWER_DATA           SPDRP = 0x0000001E // Device Power Data (R)
	SPDRP_REMOVAL_POLICY              SPDRP = 0x0000001F // Removal Policy (R)
	SPDRP_REMOVAL_POLICY_HW_DEFAULT   SPDRP = 0x00000020 // 硬件移除策略 (R)
	SPDRP_REMOVAL_POLICY_OVERRIDE     SPDRP = 0x00000021 // 移除策略覆盖（可读写）
	SPDRP_INSTALL_STATE               SPDRP = 0x00000022 // 设备安装状态 (R)
	SPDRP_LOCATION_PATHS              SPDRP = 0x00000023 // 设备位置路径 (R)
	SPDRP_BASE_CONTAINERID            SPDRP = 0x00000024 // Base ContainerID (R)

	SPDRP_MAXIMUM_PROPERTY SPDRP = 0x00000025 // 对序数的上界
)

// DEVPROPTYPE 表示统一设备属性模型中设备属性值的数据类型标识符。
type DEVPROPTYPE uint32

// 翻译提示:const  (
// 	DEVPROP_TYPEMOD_ARRAY                        DEVPROPTYPE  =  0x00001000
// 	DEVPROP_TYPEMOD_LIST                          DEVPROPTYPE  =  0x00002000
// 
// 	DEVPROP_TYPE_EMPTY                              DEVPROPTYPE  =  0x00000000
// 	DEVPROP_TYPE_NULL                                DEVPROPTYPE  =  0x00000001
// 	DEVPROP_TYPE_SIGNED_BYTE                  DEVPROPTYPE  =  0x00000002
// 	DEVPROP_TYPE_UNSIGNED_BYTE              DEVPROPTYPE  =  0x00000003
// 	DEVPROP_TYPE_SHORT                              DEVPROPTYPE  =  0x00000004
// 	DEVPROP_TYPE_UNSIGNED_SHORT            DEVPROPTYPE  =  0x00000005
// 	DEVPROP_TYPE_INT                                  DEVPROPTYPE  =  0x00000006
// 	DEVPROP_TYPE_UNSIGNED_INT                DEVPROPTYPE  =  0x00000007
// 	DEVPROP_TYPE_LONG                                DEVPROPTYPE  =  0x00000008
// 	DEVPROP_TYPE_UNSIGNED_LONG              DEVPROPTYPE  =  0x00000009
// 	DEVPROP_TYPE_FLOATING_POINT              DEVPROPTYPE  =  0x0000000A
// 	DEVPROP_TYPE_DOUBLE_FLOAT                DEVPROPTYPE  =  0x0000000B
// 	DEVPROP_TYPE_DECIMAL_NUMBER              DEVPROPTYPE  =  0x0000000C
// 	DEVPROP_TYPE_UNIQUE_IDENTIFIER        DEVPROPTYPE  =  0x0000000D
// 	DEVPROP_TYPE_CURRENCY_VALUE              DEVPROPTYPE  =  0x0000000E
// 	DEVPROP_TYPE_DATE_TIME                      DEVPROPTYPE  =  0x0000000F
// 	DEVPROP_TYPE_FILETIME_STAMP              DEVPROPTYPE  =  0x00000010
// 	DEVPROP_TYPE_BOOLEAN_VALUE                DEVPROPTYPE  =  0x00000011
// 	DEVPROP_TYPE_TEXT_STRING                    DEVPROPTYPE  =  0x00000012
// 	DEVPROP_TYPE_STRING_LIST                    DEVPROPTYPE  =  DEVPROP_TYPE_TEXT_STRING  |  DEVPROP_TYPEMOD_LIST
// 	DEVPROP_TYPE_SECURITY_DESCRIPTOR    DEVPROPTYPE  =  0x00000013
// 	DEVPROP_TYPE_SECURITY_DESC_STRING  DEVPROPTYPE  =  0x00000014
// 	DEVPROP_TYPE_PROPERTY_KEY                  DEVPROPTYPE  =  0x00000015
// 	DEVPROP_TYPE_PROPERTY_TYPE                DEVPROPTYPE  =  0x00000016
// 	DEVPROP_TYPE_BINARY_DATA                    DEVPROPTYPE  =  DEVPROP_TYPE_UNSIGNED_BYTE  |  DEVPROP_TYPEMOD_ARRAY
// 	DEVPROP_TYPE_ERROR_CODE                      DEVPROPTYPE  =  0x00000017
// 	DEVPROP_TYPE_NTSTATUS_VALUE              DEVPROPTYPE  =  0x00000018
// 	DEVPROP_TYPE_STRING_INDIRECT_REF    DEVPROPTYPE  =  0x00000019
// 
// 	MAX_DEVPROP_TYPE                                DEVPROPTYPE  =  0x00000019
// 	MAX_DEVPROP_TYPEMOD                          DEVPROPTYPE  =  0x00002000
// 
// 	DEVPROP_MASK_TYPE                                DEVPROPTYPE  =  0x00000FFF
// 	DEVPROP_MASK_TYPEMOD                          DEVPROPTYPE  =  0x0000F000
// )
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

// DEVPROPGUID 指定一个属性类别。
type DEVPROPGUID GUID

// DEVPROPID 用于在属性类别中唯一标识该属性。
type DEVPROPID uint32

// 翻译提示:const  (
//         //  开始可用的设备属性ID
//         DEVPKEY_FIRST_USABLE  DEVPKEY  =  2
// )
const DEVPROPID_FIRST_USABLE DEVPROPID = 2

// DEVPROPKEY 代表统一设备属性模型中设备属性的设备属性键。
type DEVPROPKEY struct {
	FmtID DEVPROPGUID
	PID   DEVPROPID
}

// CONFIGRET 是 cfgmgr32 API 返回的值或错误代码
type CONFIGRET uint32

// 翻译提示:func  (ret  配置结果)  错误信息()  string  {}

// ff:
func (ret CONFIGRET) Error() string {
	if win32Error, ok := ret.Unwrap().(Errno); ok {
		return fmt.Sprintf("%s (CfgMgr error: 0x%08x)", win32Error.Error(), uint32(ret))
	}
	return fmt.Sprintf("CfgMgr error: 0x%08x", uint32(ret))
}

// 翻译提示:func  (ret  配置结果)  Windows32错误(默认错误  错误码)  错误码  {}

// ff:
// defaultError:
func (ret CONFIGRET) Win32Error(defaultError Errno) Errno {
	return cm_MapCrToWin32Err(ret, defaultError)
}

// 翻译提示:func  (ret  配置结果)  转换错误()  错误  {}

// ff:
func (ret CONFIGRET) Unwrap() error {
	const noMatch = Errno(^uintptr(0))
	win32Error := ret.Win32Error(noMatch)
	if win32Error == noMatch {
		return nil
	}
	return win32Error
}

// 翻译提示:const  (
// 	成功配置返回值                  CONFIGRET  =  0x00000000
// 	默认配置返回值                  CONFIGRET  =  0x00000001
// 	内存不足错误                      CONFIGRET  =  0x00000002
// 	无效指针错误                      CONFIGRET  =  0x00000003
// 	无效标志错误                      CONFIGRET  =  0x00000004
// 	无效设备节点错误              CONFIGRET  =  0x00000005
// 	无效设备实例错误                    =  CR_INVALID_DEVNODE
// 	无效资源描述符错误          CONFIGRET  =  0x00000006
// 	无效日志配置错误              CONFIGRET  =  0x00000007
// 	无效仲裁器错误                  CONFIGRET  =  0x00000008
// 	无效节点列表错误              CONFIGRET  =  0x00000009
// 	设备节点有请求                  CONFIGRET  =  0x0000000A
// 	设备实例有请求                        =  CR_DEVNODE_HAS_REQS
// 	无效资源ID错误                  CONFIGRET  =  0x0000000B
// 	DLVXD找不到错误                CONFIGRET  =  0x0000000C
// 	无此设备节点错误              CONFIGRET  =  0x0000000D
// 	无此设备实例错误                    =  CR_NO_SUCH_DEVNODE
// 	无更多日志配置错误            CONFIGRET  =  0x0000000E
// 	无更多资源描述符错误        CONFIGRET  =  0x0000000F
// 	已存在此设备节点              CONFIGRET  =  0x00000010
// 	已存在此设备实例                        =  CR_ALREADY_SUCH_DEVNODE
// 	无效范围列表错误              CONFIGRET  =  0x00000011
// 	无效范围错误                      CONFIGRET  =  0x00000012
// 	失败错误                                  CONFIGRET  =  0x00000013
// 	无此类逻辑设备                    CONFIGRET  =  0x00000014
// 	创建被阻止                            CONFIGRET  =  0x00000015
// 	非系统虚拟机错误              CONFIGRET  =  0x00000016
// 	移除被否决                              CONFIGRET  =  0x00000017
// 	APM被否决                                CONFIGRET  =  0x00000018
// 	无效加载类型错误              CONFIGRET  =  0x00000019
// 	缓冲区太小                              CONFIGRET  =  0x0000001A
// 	无仲裁器错误                        CONFIGRET  =  0x0000001B
// 	无注册表句柄错误              CONFIGRET  =  0x0000001C
// 	注册表错误                              CONFIGRET  =  0x0000001D
// 	无效设备ID错误                    CONFIGRET  =  0x0000001E
// 	无效数据错误                      CONFIGRET  =  0x0000001F
// 	无效API错误                          CONFIGRET  =  0x00000020
// 	设备加载器未准备好              CONFIGRET  =  0x00000021
// 	需要重新启动                          CONFIGRET  =  0x00000022
// 	无更多硬件配置                    CONFIGRET  =  0x00000023
// 	设备不存在                              CONFIGRET  =  0x00000024
// 	无此值错误                              CONFIGRET  =  0x00000025
// 	类型错误                                  CONFIGRET  =  0x00000026
// 	无效优先级错误                    CONFIGRET  =  0x00000027
// 	无法禁用                                  CONFIGRET  =  0x00000028
// 	释放资源                                  CONFIGRET  =  0x00000029
// 	查询被否决                              CONFIGRET  =  0x0000002A
// 	无法共享IRQ                            CONFIGRET  =  0x0000002B
// 	无依赖项                                  CONFIGRET  =  0x0000002C
// 	资源相同                                  CONFIGRET  =  0x0000002D
// 	无此注册表键错误              CONFIGRET  =  0x0000002E
// 	无效机器名错误                    CONFIGRET  =  0x0000002F
// 	远程通信故障                          CONFIGRET  =  0x00000030
// 	机器不可用                              CONFIGRET  =  0x00000031
// 	无CM服务错误                        CONFIGRET  =  0x00000032
// 	访问被拒绝                              CONFIGRET  =  0x00000033
// 	调用未实现错误                    CONFIGRET  =  0x00000034
// 	无效属性错误                        CONFIGRET  =  0x00000035
// 	设备接口活动                          CONFIGRET  =  0x00000036
// 	无此设备接口错误              CONFIGRET  =  0x00000037
// 	无效引用字符串错误            CONFIGRET  =  0x00000038
// 	无效冲突列表错误              CONFIGRET  =  0x00000039
// 	无效索引错误                        CONFIGRET  =  0x0000003A
// 	无效结构大小错误                CONFIGRET  =  0x0000003B
// 	配置返回值总数                    CONFIGRET  =  0x0000003C
// )
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

// 翻译提示:const  (
// 	获取活跃设备接口列表                =  0  //  只包含当前活动的设备接口
// 	获取所有已注册设备接口列表      =  1  //  包含所有已注册的设备接口，无论是否活动
// )
const (
	CM_GET_DEVICE_INTERFACE_LIST_PRESENT     = 0 // 仅当前“活跃”的设备接口
	CM_GET_DEVICE_INTERFACE_LIST_ALL_DEVICES = 1 // 所有已注册的设备接口，无论是否活跃
)

const (
	DN_ROOT_ENUMERATED       = 0x00000001        // Was enumerated by ROOT
	DN_DRIVER_LOADED         = 0x00000002        // 具有 Register_Device_Driver
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
	DN_NEED_RESTART          = DN_LIAR           // 为了使该Devnode正常工作，系统需要重新启动
	DN_DRIVER_BLOCKED        = DN_NOT_FIRST_TIME // 一个或多个驱动程序因该设备节点而被阻止加载
	DN_LEGACY_DRIVER         = DN_MOVED          // 该设备正在使用一个遗留驱动程序
	DN_CHILD_WITH_INVALID_ID = DN_HAS_MARK       // 一个或多个子元素具有无效的ID
	DN_DEVICE_DISCONNECTED   = DN_NEEDS_LOCKING  // 设备的驱动函数报告该设备未连接。通常这意味着无线设备已超出范围。
	DN_QUERY_REMOVE_PENDING  = DN_MF_PARENT      // Device 是一组待查询移除的关联设备集合中的一部分
	DN_QUERY_REMOVE_ACTIVE   = DN_MF_CHILD       // Device当前正忙于处理一个查询-移除IRP
	DN_CHANGEABLE_FLAGS      = DN_NOT_FIRST_TIME | DN_HARDWARE_ENUM | DN_HAS_MARK | DN_DISABLEABLE | DN_REMOVABLE | DN_MF_CHILD | DN_MF_PARENT | DN_NOT_FIRST_TIMEE | DN_STOP_FREE_RES | DN_REBAL_CANDIDATE | DN_NT_ENUMERATOR | DN_NT_DRIVER | DN_SILENT_INSTALL | DN_NO_SHOW_IN_DM
)

//sys	setupDiCreateDeviceInfoListEx(classGUID *GUID, hwndParent uintptr, machineName *uint16, reserved uintptr) (handle DevInfo, err error) [failretval==DevInfo(InvalidHandle)] = setupapi.SetupDiCreateDeviceInfoListExW

// SetupDiCreateDeviceInfoListEx 函数在远程或本地计算机上创建一个空的设备信息集，并可选择性地将此集与设备安装类关联。

// ff:
// classGUID:
// hwndParent:
// machineName:
// deviceInfoSet:
// err:
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

//sys	setupDiGetDeviceInfoListDetail(deviceInfoSet DevInfo, deviceInfoSetDetailData *DevInfoListDetailData) (err error) = setupapi.SetupDiGetDeviceInfoListDetailW

// SetupDiGetDeviceInfoListDetail 函数用于获取与设备信息集关联的信息，包括类 GUID、远程计算机句柄以及远程计算机名。

// ff:
// deviceInfoSet:
// deviceInfoSetDetailData:
// err:
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

//sys	setupDiCreateDeviceInfo(deviceInfoSet DevInfo, DeviceName *uint16, classGUID *GUID, DeviceDescription *uint16, hwndParent uintptr, CreationFlags DICD, deviceInfoData *DevInfoData) (err error) = setupapi.SetupDiCreateDeviceInfoW

// SetupDiCreateDeviceInfo 函数用于创建一个新的设备信息元素，并将其作为新成员添加到指定的设备信息集中。

// ff:
// deviceInfoSet:
// deviceName:
// classGUID:
// deviceDescription:
// hwndParent:
// creationFlags:
// deviceInfoData:
// err:
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

// CreateDeviceInfo 方法创建一个新的设备信息元素，并将其作为新成员添加到指定的设备信息集中。

// ff:
// deviceName:
// classGUID:
// deviceDescription:
// hwndParent:
// creationFlags:
func (deviceInfoSet DevInfo) CreateDeviceInfo(deviceName string, classGUID *GUID, deviceDescription string, hwndParent uintptr, creationFlags DICD) (*DevInfoData, error) {
	return SetupDiCreateDeviceInfo(deviceInfoSet, deviceName, classGUID, deviceDescription, hwndParent, creationFlags)
}

//sys	setupDiEnumDeviceInfo(deviceInfoSet DevInfo, memberIndex uint32, deviceInfoData *DevInfoData) (err error) = setupapi.SetupDiEnumDeviceInfo

// SetupDiEnumDeviceInfo 函数返回一个 DevInfoData 结构，该结构用于指定设备信息集中某个设备信息元素。

// ff:
// deviceInfoSet:
// memberIndex:
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

// SetupDiDestroyDeviceInfoList function deletes a device information set and frees all associated memory.
//sys	SetupDiDestroyDeviceInfoList(deviceInfoSet DevInfo) (err error) = setupapi.SetupDiDestroyDeviceInfoList

// Close 方法用于删除一个设备信息集并释放所有相关内存。

// ff:
func (deviceInfoSet DevInfo) Close() error {
	return SetupDiDestroyDeviceInfoList(deviceInfoSet)
}

//sys	SetupDiBuildDriverInfoList(deviceInfoSet DevInfo, deviceInfoData *DevInfoData, driverType SPDIT) (err error) = setupapi.SetupDiBuildDriverInfoList

// BuildDriverInfoList 方法用于为特定设备或设备信息集中与全局类驱动程序列表关联的驱动程序构建一个列表。

// ff:
// deviceInfoData:
// driverType:
func (deviceInfoSet DevInfo) BuildDriverInfoList(deviceInfoData *DevInfoData, driverType SPDIT) error {
	return SetupDiBuildDriverInfoList(deviceInfoSet, deviceInfoData, driverType)
}

//sys	SetupDiCancelDriverInfoSearch(deviceInfoSet DevInfo) (err error) = setupapi.SetupDiCancelDriverInfoSearch

// CancelDriverInfoSearch 方法取消在另一个线程中正在进行的驾驶员信息列表搜索。

// ff:
func (deviceInfoSet DevInfo) CancelDriverInfoSearch() error {
	return SetupDiCancelDriverInfoSearch(deviceInfoSet)
}

//sys	setupDiEnumDriverInfo(deviceInfoSet DevInfo, deviceInfoData *DevInfoData, driverType SPDIT, memberIndex uint32, driverInfoData *DrvInfoData) (err error) = setupapi.SetupDiEnumDriverInfoW

// SetupDiEnumDriverInfo 函数用于枚举驱动列表中的成员。

// ff:
// deviceInfoSet:
// deviceInfoData:
// driverType:
// memberIndex:
func SetupDiEnumDriverInfo(deviceInfoSet DevInfo, deviceInfoData *DevInfoData, driverType SPDIT, memberIndex int) (*DrvInfoData, error) {
	data := &DrvInfoData{}
	data.size = uint32(unsafe.Sizeof(*data))

	return data, setupDiEnumDriverInfo(deviceInfoSet, deviceInfoData, driverType, uint32(memberIndex), data)
}

// EnumDriverInfo 方法枚举驱动列表中的成员。

// ff:
// deviceInfoData:
// driverType:
// memberIndex:
func (deviceInfoSet DevInfo) EnumDriverInfo(deviceInfoData *DevInfoData, driverType SPDIT, memberIndex int) (*DrvInfoData, error) {
	return SetupDiEnumDriverInfo(deviceInfoSet, deviceInfoData, driverType, memberIndex)
}

//sys	setupDiGetSelectedDriver(deviceInfoSet DevInfo, deviceInfoData *DevInfoData, driverInfoData *DrvInfoData) (err error) = setupapi.SetupDiGetSelectedDriverW

// SetupDiGetSelectedDriver 函数用于为设备信息集或特定设备信息元素获取所选驱动程序。

// ff:
// deviceInfoSet:
// deviceInfoData:
func SetupDiGetSelectedDriver(deviceInfoSet DevInfo, deviceInfoData *DevInfoData) (*DrvInfoData, error) {
	data := &DrvInfoData{}
	data.size = uint32(unsafe.Sizeof(*data))

	return data, setupDiGetSelectedDriver(deviceInfoSet, deviceInfoData, data)
}

// SelectedDriver 方法用于检索设备信息集或特定设备信息元素所选中的驱动程序。

// ff:
// deviceInfoData:
func (deviceInfoSet DevInfo) SelectedDriver(deviceInfoData *DevInfoData) (*DrvInfoData, error) {
	return SetupDiGetSelectedDriver(deviceInfoSet, deviceInfoData)
}

//sys	SetupDiSetSelectedDriver(deviceInfoSet DevInfo, deviceInfoData *DevInfoData, driverInfoData *DrvInfoData) (err error) = setupapi.SetupDiSetSelectedDriverW

// SetSelectedDriver 方法用于设置或重置设备信息元素的已选驱动程序，或者为设备信息集合设置已选类驱动程序。

// ff:
// deviceInfoData:
// driverInfoData:
func (deviceInfoSet DevInfo) SetSelectedDriver(deviceInfoData *DevInfoData, driverInfoData *DrvInfoData) error {
	return SetupDiSetSelectedDriver(deviceInfoSet, deviceInfoData, driverInfoData)
}

//sys	setupDiGetDriverInfoDetail(deviceInfoSet DevInfo, deviceInfoData *DevInfoData, driverInfoData *DrvInfoData, driverInfoDetailData *DrvInfoDetailData, driverInfoDetailDataSize uint32, requiredSize *uint32) (err error) = setupapi.SetupDiGetDriverInfoDetailW

// SetupDiGetDriverInfoDetail 函数用于为设备信息集中或设备信息集中的特定设备信息元素获取驱动程序信息详细内容。

// ff:
// deviceInfoSet:
// deviceInfoData:
// driverInfoData:
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

// DriverInfoDetail 方法用于获取设备信息集中某个设备信息元素的驱动程序详细信息，或直接为整个设备信息集获取此类信息。

// ff:
// deviceInfoData:
// driverInfoData:
func (deviceInfoSet DevInfo) DriverInfoDetail(deviceInfoData *DevInfoData, driverInfoData *DrvInfoData) (*DrvInfoDetailData, error) {
	return SetupDiGetDriverInfoDetail(deviceInfoSet, deviceInfoData, driverInfoData)
}

//sys	SetupDiDestroyDriverInfoList(deviceInfoSet DevInfo, deviceInfoData *DevInfoData, driverType SPDIT) (err error) = setupapi.SetupDiDestroyDriverInfoList

// DestroyDriverInfoList 方法用于删除一个驱动列表。

// ff:
// deviceInfoData:
// driverType:
func (deviceInfoSet DevInfo) DestroyDriverInfoList(deviceInfoData *DevInfoData, driverType SPDIT) error {
	return SetupDiDestroyDriverInfoList(deviceInfoSet, deviceInfoData, driverType)
}

//sys	setupDiGetClassDevsEx(classGUID *GUID, Enumerator *uint16, hwndParent uintptr, Flags DIGCF, deviceInfoSet DevInfo, machineName *uint16, reserved uintptr) (handle DevInfo, err error) [failretval==DevInfo(InvalidHandle)] = setupapi.SetupDiGetClassDevsExW

// SetupDiGetClassDevsEx 函数返回一个设备信息集的句柄，该集合包含了为本地或远程计算机请求的设备信息元素。

// ff:
// classGUID:
// enumerator:
// hwndParent:
// flags:
// deviceInfoSet:
// machineName:
// handle:
// err:
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

// SetupDiCallClassInstaller function calls the appropriate class installer, and any registered co-installers, with the specified installation request (DIF code).
//sys	SetupDiCallClassInstaller(installFunction DI_FUNCTION, deviceInfoSet DevInfo, deviceInfoData *DevInfoData) (err error) = setupapi.SetupDiCallClassInstaller

// CallClassInstaller 成员调用相应类安装程序以及所有已注册的共同安装程序，同时传递指定的安装请求（DIF代码）。

// ff:
// installFunction:
// deviceInfoData:
func (deviceInfoSet DevInfo) CallClassInstaller(installFunction DI_FUNCTION, deviceInfoData *DevInfoData) error {
	return SetupDiCallClassInstaller(installFunction, deviceInfoSet, deviceInfoData)
}

// SetupDiOpenDevRegKey function opens a registry key for device-specific configuration information.
//sys	SetupDiOpenDevRegKey(deviceInfoSet DevInfo, deviceInfoData *DevInfoData, Scope DICS_FLAG, HwProfile uint32, KeyType DIREG, samDesired uint32) (key Handle, err error) [failretval==InvalidHandle] = setupapi.SetupDiOpenDevRegKey

// OpenDevRegKey 方法用于打开一个注册表键，以便访问设备特定的配置信息。

// ff:
// DeviceInfoData:
// Scope:
// HwProfile:
// KeyType:
// samDesired:
// Handle:
func (deviceInfoSet DevInfo) OpenDevRegKey(DeviceInfoData *DevInfoData, Scope DICS_FLAG, HwProfile uint32, KeyType DIREG, samDesired uint32) (Handle, error) {
	return SetupDiOpenDevRegKey(deviceInfoSet, DeviceInfoData, Scope, HwProfile, KeyType, samDesired)
}

//sys	setupDiGetDeviceProperty(deviceInfoSet DevInfo, deviceInfoData *DevInfoData, propertyKey *DEVPROPKEY, propertyType *DEVPROPTYPE, propertyBuffer *byte, propertyBufferSize uint32, requiredSize *uint32, flags uint32) (err error) = setupapi.SetupDiGetDevicePropertyW

// SetupDiGetDeviceProperty 函数用于检索指定的设备实例属性。

// ff:
// deviceInfoSet:
// deviceInfoData:
// propertyKey:
// value:
// err:
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

//sys	setupDiGetDeviceRegistryProperty(deviceInfoSet DevInfo, deviceInfoData *DevInfoData, property SPDRP, propertyRegDataType *uint32, propertyBuffer *byte, propertyBufferSize uint32, requiredSize *uint32) (err error) = setupapi.SetupDiGetDeviceRegistryPropertyW

// SetupDiGetDeviceRegistryProperty 函数用于检索指定即插即用设备的属性。

// ff:
// deviceInfoSet:
// deviceInfoData:
// property:
// value:
// err:
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
// deviceInfoData:
// property:
func (deviceInfoSet DevInfo) DeviceRegistryProperty(deviceInfoData *DevInfoData, property SPDRP) (interface{}, error) {
	return SetupDiGetDeviceRegistryProperty(deviceInfoSet, deviceInfoData, property)
}

//sys	setupDiSetDeviceRegistryProperty(deviceInfoSet DevInfo, deviceInfoData *DevInfoData, property SPDRP, propertyBuffer *byte, propertyBufferSize uint32) (err error) = setupapi.SetupDiSetDeviceRegistryPropertyW

// SetupDiSetDeviceRegistryProperty 函数为设备设置即插即用设备属性。

// ff:
// deviceInfoSet:
// deviceInfoData:
// property:
// propertyBuffers:
func SetupDiSetDeviceRegistryProperty(deviceInfoSet DevInfo, deviceInfoData *DevInfoData, property SPDRP, propertyBuffers []byte) error {
	return setupDiSetDeviceRegistryProperty(deviceInfoSet, deviceInfoData, property, &propertyBuffers[0], uint32(len(propertyBuffers)))
}

// SetDeviceRegistryProperty 函数用于为设备设置即插即用设备属性。

// ff:
// deviceInfoData:
// property:
// propertyBuffers:
func (deviceInfoSet DevInfo) SetDeviceRegistryProperty(deviceInfoData *DevInfoData, property SPDRP, propertyBuffers []byte) error {
	return SetupDiSetDeviceRegistryProperty(deviceInfoSet, deviceInfoData, property, propertyBuffers)
}

// SetDeviceRegistryPropertyString 方法为设备设置即插即用设备属性字符串。

// ff:
// deviceInfoData:
// property:
// str:
func (deviceInfoSet DevInfo) SetDeviceRegistryPropertyString(deviceInfoData *DevInfoData, property SPDRP, str string) error {
	str16, err := UTF16FromString(str)
	if err != nil {
		return err
	}
	err = SetupDiSetDeviceRegistryProperty(deviceInfoSet, deviceInfoData, property, utf16ToBuf(append(str16, 0)))
	runtime.KeepAlive(str16)
	return err
}

//sys	setupDiGetDeviceInstallParams(deviceInfoSet DevInfo, deviceInfoData *DevInfoData, deviceInstallParams *DevInstallParams) (err error) = setupapi.SetupDiGetDeviceInstallParamsW

// SetupDiGetDeviceInstallParams 函数用于为一个设备信息集或特定的设备信息元素检索设备安装参数。

// ff:
// deviceInfoSet:
// deviceInfoData:
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

//sys	setupDiGetDeviceInstanceId(deviceInfoSet DevInfo, deviceInfoData *DevInfoData, instanceId *uint16, instanceIdSize uint32, instanceIdRequiredSize *uint32) (err error) = setupapi.SetupDiGetDeviceInstanceIdW

// SetupDiGetDeviceInstanceId 函数用于获取设备的实例ID。

// ff:
// deviceInfoSet:
// deviceInfoData:
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

// SetupDiGetClassInstallParams function retrieves class installation parameters for a device information set or a particular device information element.
//sys	SetupDiGetClassInstallParams(deviceInfoSet DevInfo, deviceInfoData *DevInfoData, classInstallParams *ClassInstallHeader, classInstallParamsSize uint32, requiredSize *uint32) (err error) = setupapi.SetupDiGetClassInstallParamsW

// ClassInstallParams 方法用于为一个设备信息集或特定的设备信息元素检索类安装参数。

// ff:
// deviceInfoData:
// classInstallParams:
// classInstallParamsSize:
// requiredSize:
func (deviceInfoSet DevInfo) ClassInstallParams(deviceInfoData *DevInfoData, classInstallParams *ClassInstallHeader, classInstallParamsSize uint32, requiredSize *uint32) error {
	return SetupDiGetClassInstallParams(deviceInfoSet, deviceInfoData, classInstallParams, classInstallParamsSize, requiredSize)
}

//sys	SetupDiSetDeviceInstallParams(deviceInfoSet DevInfo, deviceInfoData *DevInfoData, deviceInstallParams *DevInstallParams) (err error) = setupapi.SetupDiSetDeviceInstallParamsW

// SetDeviceInstallParams 成员为一个设备信息集或特定的设备信息元素设置设备安装参数。

// ff:
// deviceInfoData:
// deviceInstallParams:
func (deviceInfoSet DevInfo) SetDeviceInstallParams(deviceInfoData *DevInfoData, deviceInstallParams *DevInstallParams) error {
	return SetupDiSetDeviceInstallParams(deviceInfoSet, deviceInfoData, deviceInstallParams)
}

// SetupDiSetClassInstallParams function sets or clears class install parameters for a device information set or a particular device information element.
//sys	SetupDiSetClassInstallParams(deviceInfoSet DevInfo, deviceInfoData *DevInfoData, classInstallParams *ClassInstallHeader, classInstallParamsSize uint32) (err error) = setupapi.SetupDiSetClassInstallParamsW

// SetClassInstallParams 方法用于为一个设备信息集或特定的设备信息元素设置或清除类安装参数。

// ff:
// deviceInfoData:
// classInstallParams:
// classInstallParamsSize:
func (deviceInfoSet DevInfo) SetClassInstallParams(deviceInfoData *DevInfoData, classInstallParams *ClassInstallHeader, classInstallParamsSize uint32) error {
	return SetupDiSetClassInstallParams(deviceInfoSet, deviceInfoData, classInstallParams, classInstallParamsSize)
}

//sys	setupDiClassNameFromGuidEx(classGUID *GUID, className *uint16, classNameSize uint32, requiredSize *uint32, machineName *uint16, reserved uintptr) (err error) = setupapi.SetupDiClassNameFromGuidExW

// SetupDiClassNameFromGuidEx 函数用于获取与类GUID关联的类名。该类可以安装在本地计算机或远程计算机上。

// ff:
// classGUID:
// machineName:
// className:
// err:
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

//sys	setupDiClassGuidsFromNameEx(className *uint16, classGuidList *GUID, classGuidListSize uint32, requiredSize *uint32, machineName *uint16, reserved uintptr) (err error) = setupapi.SetupDiClassGuidsFromNameExW

// SetupDiClassGuidsFromNameEx 函数用于获取与指定类名关联的 GUID。返回的列表包含本地或远程计算机上当前已安装的此类别。

// ff:
// className:
// machineName:
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

// SetupDiSetSelectedDevice function sets a device information element as the selected member of a device information set. This function is typically used by an installation wizard.
//sys	SetupDiSetSelectedDevice(deviceInfoSet DevInfo, deviceInfoData *DevInfoData) (err error) = setupapi.SetupDiSetSelectedDevice

// SetSelectedDevice 方法用于将一个设备信息元素设置为设备信息集的选定成员。此函数通常由安装向导使用。

// ff:
// deviceInfoData:
func (deviceInfoSet DevInfo) SetSelectedDevice(deviceInfoData *DevInfoData) error {
	return SetupDiSetSelectedDevice(deviceInfoSet, deviceInfoData)
}

//sys	setupUninstallOEMInf(infFileName *uint16, flags SUOI, reserved uintptr) (err error) = setupapi.SetupUninstallOEMInfW

// SetupUninstallOEMInf 卸载指定的驱动程序

// ff:
// infFileName:
// flags:
func SetupUninstallOEMInf(infFileName string, flags SUOI) error {
	infFileName16, err := UTF16PtrFromString(infFileName)
	if err != nil {
		return err
	}
	return setupUninstallOEMInf(infFileName16, flags, 0)
}

//sys cm_MapCrToWin32Err(configRet CONFIGRET, defaultWin32Error Errno) (ret Errno) = CfgMgr32.CM_MapCrToWin32Err

//sys cm_Get_Device_Interface_List_Size(len *uint32, interfaceClass *GUID, deviceID *uint16, flags uint32) (ret CONFIGRET) = CfgMgr32.CM_Get_Device_Interface_List_SizeW
//sys cm_Get_Device_Interface_List(interfaceClass *GUID, deviceID *uint16, buffer *uint16, bufferLen uint32, flags uint32) (ret CONFIGRET) = CfgMgr32.CM_Get_Device_Interface_ListW

// 翻译提示:func  获取设备接口列表(deviceID  字符串,  接口类别  *GUID,  标志  uint32)  ([]字符串,  错误)  {}

// ff:
// deviceID:
// interfaceClass:
// flags:
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

// 翻译提示:func  获取设备节点状态(status  *uint32,  problemNumber  *uint32,  设备实例  DEVINST,  标志  uint32)  error  {}

// ff:
// status:
// problemNumber:
// devInst:
// flags:
func CM_Get_DevNode_Status(status *uint32, problemNumber *uint32, devInst DEVINST, flags uint32) error {
	ret := cm_Get_DevNode_Status(status, problemNumber, devInst, flags)
	if ret == CR_SUCCESS {
		return nil
	}
	return ret
}
