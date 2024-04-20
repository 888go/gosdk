
<原文开始>
// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// 版权所有 ? 2021 Go作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。
# <翻译结束>


<原文开始>
// This file contains functions that wrap SetupAPI.dll and CfgMgr32.dll,
// core system functions for managing hardware devices, drivers, and the PnP tree.
// Information about these APIs can be found at:
//     https://docs.microsoft.com/en-us/windows-hardware/drivers/install/setupapi
//     https://docs.microsoft.com/en-us/windows/win32/devinst/cfgmgr32-
<原文结束>

# <翻译开始>
// 本文件包含对SetupAPI.dll和CfgMgr32.dll的封装函数，
// 这些是用于管理硬件设备、驱动程序及PnP树的核心系统功能。
// 可在以下链接获取关于这些API的信息：
//     https://docs.microsoft.com/zh-cn/windows-hardware/drivers/install/setupapi
//     https://docs.microsoft.com/zh-cn/windows/win32/devinst/cfgmgr32-
# <翻译结束>


<原文开始>
// 38 chars + terminator null
<原文结束>

# <翻译开始>
// 38个字符 + 终止符空字符
# <翻译结束>


<原文开始>
// Maximum string length constants
<原文结束>

# <翻译开始>
// 最大字符串长度常量
# <翻译结束>


<原文开始>
// Windows 9x-compatible maximum for displayable strings coming from a device INF.
<原文结束>

# <翻译开始>
// Windows 9x 兼容的最大可显示字符串长度，该字符串来自设备 INF。
# <翻译结束>


<原文开始>
// Actual maximum size of an INF string (including string substitutions).
<原文结束>

# <翻译开始>
// 实际INF字符串（包括字符串替换）的最大尺寸
# <翻译结束>


<原文开始>
// For Windows 9x compatibility, INF section names should be constrained to 32 characters.
<原文结束>

# <翻译开始>
// 为了与Windows 9x兼容，INF节名称应限制为32个字符。
# <翻译结束>


<原文开始>
// SP_MAX_MACHINENAME_LENGTH defines maximum length of a machine name in the format expected by ConfigMgr32 CM_Connect_Machine (i.e., "\\\\MachineName\0").
<原文结束>

# <翻译开始>
// SP_MAX_MACHINENAME_LENGTH 定义了以 ConfigMgr32 CM_Connect_Machine 接收格式表示的机器名的最大长度（即，“\\\\MachineName\\0”）。
# <翻译结束>


<原文开始>
// HSPFILEQ is type for setup file queue
<原文结束>

# <翻译开始>
// HSPFILEQ 是用于设置文件队列的类型
# <翻译结束>


<原文开始>
// DevInfo holds reference to device information set
<原文结束>

# <翻译开始>
// DevInfo 用于保存设备信息集的引用
# <翻译结束>


<原文开始>
// DEVINST is a handle usually recognized by cfgmgr32 APIs
<原文结束>

# <翻译开始>
// DEVINST 是一个句柄，通常被 cfgmgr32 API 接口所识别
# <翻译结束>


<原文开始>
// DevInfoData is a device information structure (references a device instance that is a member of a device information set)
<原文结束>

# <翻译开始>
// DevInfoData 是一个设备信息结构（引用了作为设备信息集合成员的设备实例）
# <翻译结束>


<原文开始>
// DevInfoListDetailData is a structure for detailed information on a device information set (used for SetupDiGetDeviceInfoListDetail which supersedes the functionality of SetupDiGetDeviceInfoListClass).
<原文结束>

# <翻译开始>
// DevInfoListDetailData 是一个用于设备信息集详细信息的结构体（用于 SetupDiGetDeviceInfoListDetail，该函数取代了 SetupDiGetDeviceInfoListClass 的功能）。
# <翻译结束>


<原文开始>
// Use unsafeSizeOf method
<原文结束>

# <翻译开始>
// 使用unsafeSizeOf方法
# <翻译结束>


<原文开始>
// Windows declares this with pshpack1.h
<原文结束>

# <翻译开始>
// Windows 在 pshpack1.h 中声明了这一点
# <翻译结束>


<原文开始>
// DI_FUNCTION is function type for device installer
<原文结束>

# <翻译开始>
// DI_FUNCTION 是设备安装器的功能函数类型
# <翻译结束>


<原文开始>
// DevInstallParams is device installation parameters structure (associated with a particular device information element, or globally with a device information set)
<原文结束>

# <翻译开始>
// DevInstallParams 是设备安装参数结构（与特定的设备信息项关联，或全局与设备信息集关联）
# <翻译结束>


<原文开始>
// DI_FLAGS is SP_DEVINSTALL_PARAMS.Flags values
<原文结束>

# <翻译开始>
// DI_FLAGS 是 SP_DEVINSTALL_PARAMS.Flags 的值
# <翻译结束>


<原文开始>
// Flags for choosing a device
<原文结束>

# <翻译开始>
// 用于选择设备的标志
# <翻译结束>


<原文开始>
// support Other... button
<原文结束>

# <翻译开始>
// 支持“其它…”按钮
# <翻译结束>


<原文开始>
// show compatibility list
<原文结束>

# <翻译开始>
// 显示兼容性列表
# <翻译结束>


<原文开始>
// both class & compat list shown
<原文结束>

# <翻译开始>
// 同时显示类与兼容列表
# <翻译结束>


<原文开始>
// Searched for compatible devices
<原文结束>

# <翻译开始>
// 搜索兼容设备
# <翻译结束>


<原文开始>
// Searched for class devices
<原文结束>

# <翻译开始>
// 搜索类设备
# <翻译结束>


<原文开始>
// No UI for resources if possible
<原文结束>

# <翻译开始>
// 如果可能，资源不使用UI
# <翻译结束>


<原文开始>
// Flags returned by DiInstallDevice to indicate need to reboot/restart
<原文结束>

# <翻译开始>
// DiInstallDevice返回的标志，用于指示需要重启/重新启动
# <翻译结束>


<原文开始>
// Reboot required to take effect
<原文结束>

# <翻译开始>
// 重启后生效
# <翻译结束>


<原文开始>
// Flags for device installation
<原文结束>

# <翻译开始>
// 设备安装标志
# <翻译结束>


<原文开始>
// no Browse... in InsertDisk
<原文结束>

# <翻译开始>
// InsertDisk中无“浏览…”
# <翻译结束>


<原文开始>
// Flags set by DiBuildDriverInfoList
<原文结束>

# <翻译开始>
// 由 DiBuildDriverInfoList 设置的标志
# <翻译结束>


<原文开始>
// Set if multiple manufacturers in class driver list
<原文结束>

# <翻译开始>
// 如果类驱动列表中存在多个制造商，则设置此选项
# <翻译结束>


<原文开始>
// Flag indicates that device is disabled
<原文结束>

# <翻译开始>
// 标志表明设备已禁用
# <翻译结束>


<原文开始>
// Flags for Device/Class Properties
<原文结束>

# <翻译开始>
// 设备/类属性标志
# <翻译结束>


<原文开始>
// Flag to indicate the setting properties for this Device (or class) caused a change so the Dev Mgr UI probably needs to be updated.
<原文结束>

# <翻译开始>
// 标志，表示为此设备（或此类）设置属性时引发了更改，因此可能需要更新设备管理器UI。
# <翻译结束>


<原文开始>
// Flag to indicate that the sorting from the INF file should be used.
<原文结束>

# <翻译开始>
// 标志，表示应使用INF文件中的排序。
# <翻译结束>


<原文开始>
// Flag to indicate that only the INF specified by SP_DEVINSTALL_PARAMS.DriverPath should be searched.
<原文结束>

# <翻译开始>
// 标志用于指示仅应搜索由 SP_DEVINSTALL_PARAMS.DriverPath 指定的 INF。
# <翻译结束>


<原文开始>
	// Flag that prevents ConfigMgr from removing/re-enumerating devices during device
	// registration, installation, and deletion.
<原文结束>

# <翻译开始>
// 防止ConfigMgr在设备注册、安装和删除期间移除或重新枚举设备的标志。
# <翻译结束>


<原文开始>
// The following flag can be used to install a device disabled
<原文结束>

# <翻译开始>
// 以下标志可用于安装一个禁用的设备
# <翻译结束>


<原文开始>
	// Flag that causes SetupDiBuildDriverInfoList to build a device's compatible driver
	// list from its existing class driver list, instead of the normal INF search.
<原文结束>

# <翻译开始>
// 该标志导致 SetupDiBuildDriverInfoList 从设备现有的类驱动程序列表中构建其兼容驱动程序列表，而非进行常规的 INF 搜索。
# <翻译结束>


<原文开始>
// This flag is set if the Class Install params should be used.
<原文结束>

# <翻译开始>
// 如果应使用类安装参数，则设置此标志
# <翻译结束>


<原文开始>
// This flag is set if the caller of DiCallClassInstaller does NOT want the internal default action performed if the Class installer returns ERROR_DI_DO_DEFAULT.
<原文结束>

# <翻译开始>
// 如果调用者不希望在类安装程序返回 ERROR_DI_DO_DEFAULT 时执行内部默认操作，则设置此标志。
# <翻译结束>


<原文开始>
// Force files to be copied from install path
<原文结束>

# <翻译开始>
// 强制将文件从安装路径复制
# <翻译结束>


<原文开始>
// Prop provider added Driver page.
<原文结束>

# <翻译开始>
// Prop提供商已添加Driver页面。
# <翻译结束>


<原文开始>
// Use Class Installer Provided strings in the Select Device Dlg
<原文结束>

# <翻译开始>
// 在“选择设备”对话框中使用由类安装程序提供的字符串
# <翻译结束>


<原文开始>
// No Enable/Disable in General Props
<原文结束>

# <翻译开始>
// 通用属性中无启用/禁用
# <翻译结束>


<原文开始>
// No small icons in select device dialogs
<原文结束>

# <翻译开始>
// 选择设备对话框中不显示小图标
# <翻译结束>


<原文开始>
// DI_FLAGSEX is SP_DEVINSTALL_PARAMS.FlagsEx values
<原文结束>

# <翻译开始>
// DI_FLAGSEX 是 SP_DEVINSTALL_PARAMS 结构体中 FlagsEx 字段的值
# <翻译结束>


<原文开始>
// Failed to Load/Call class installer
<原文结束>

# <翻译开始>
// 加载/调用类安装程序失败
# <翻译结束>


<原文开始>
// Class/co-installer wants to get a DIF_FINISH_INSTALL action in client context.
<原文结束>

# <翻译开始>
// 类/共同安装程序希望在客户端上下文中获取一个 DIF_FINISH_INSTALL 动作
# <翻译结束>


<原文开始>
// Did the Class Info List
<原文结束>

# <翻译开始>
// 是否已获取班级信息列表
# <翻译结束>


<原文开始>
// Did the Compat Info List
<原文结束>

# <翻译开始>
// 是否已获取兼容信息列表
# <翻译结束>


<原文开始>
// One or more device property sheets have had changes made to them, and need to have a DIF_PROPERTYCHANGE occur.
<原文结束>

# <翻译开始>
// 一个或多个设备属性表已发生更改，需要触发 DIF_PROPERTYCHANGE 事件。
# <翻译结束>


<原文开始>
// Don't run AddReg and DelReg for device's software (driver) key.
<原文结束>

# <翻译开始>
// 不要为设备的软件（驱动）键运行 AddReg 和 DelReg。
# <翻译结束>


<原文开始>
// Installation is occurring during initial system setup.
<原文结束>

# <翻译开始>
// 正在进行初始系统设置时的安装
# <翻译结束>


<原文开始>
// Driver came from Windows Update
<原文结束>

# <翻译开始>
// 驱动程序来自Windows更新
# <翻译结束>


<原文开始>
// Cause SetupDiBuildDriverInfoList to append a new driver list to an existing list.
<原文结束>

# <翻译开始>
// 令 SetupDiBuildDriverInfoList 将一个新驱动列表追加至现有列表。
# <翻译结束>


<原文开始>
// build driver list from INF(s) retrieved from URL specified in SP_DEVINSTALL_PARAMS.DriverPath (empty string means Windows Update website)
<原文结束>

# <翻译开始>
// 从 SP_DEVINSTALL_PARAMS.DriverPath 指定的 URL（空字符串表示 Windows Update 网站）获取的 INF 文件中构建驱动程序列表
# <翻译结束>


<原文开始>
// class installer added their own power page
<原文结束>

# <翻译开始>
// 类安装程序添加了它们自己的电源页面
# <翻译结束>


<原文开始>
// only include similar drivers in class list
<原文结束>

# <翻译开始>
// 只在类别列表中包含相似的驱动程序
# <翻译结束>


<原文开始>
// only add the installed driver to the class or compat driver list.  Used in calls to SetupDiBuildDriverInfoList
<原文结束>

# <翻译开始>
// 只将已安装的驱动添加到类或兼容驱动列表中。用于对SetupDiBuildDriverInfoList的调用
# <翻译结束>


<原文开始>
// Build driver list based on alternate platform information specified in associated file queue
<原文结束>

# <翻译开始>
// 根据关联文件队列中指定的备用平台信息构建驱动列表
# <翻译结束>


<原文开始>
// only restart the device drivers are being installed on as opposed to restarting all devices using those drivers.
<原文结束>

# <翻译开始>
// 仅重启正在安装设备驱动的设备，而非重启所有使用这些驱动的设备
# <翻译结束>


<原文开始>
// Tell SetupDiBuildDriverInfoList to do a recursive search
<原文结束>

# <翻译开始>
// 告诉 SetupDiBuildDriverInfoList 进行递归搜索
# <翻译结束>


<原文开始>
// Tell SetupDiBuildDriverInfoList to do a "published INF" search
<原文结束>

# <翻译开始>
// 告诉SetupDiBuildDriverInfoList执行“已发布INF”搜索
# <翻译结束>


<原文开始>
// ClassInstallHeader is the first member of any class install parameters structure. It contains the device installation request code that defines the format of the rest of the install parameters structure.
<原文结束>

# <翻译开始>
// ClassInstallHeader 是任何类安装参数结构的第一个成员。它包含了定义该安装参数结构其余部分格式的设备安装请求代码。
# <翻译结束>


<原文开始>
// DICS_STATE specifies values indicating a change in a device's state
<原文结束>

# <翻译开始>
// DICS_STATE 定义了一系列值，用于指示设备状态的变更
# <翻译结束>


<原文开始>
// The device is being enabled.
<原文结束>

# <翻译开始>
// 设备正在被启用
# <翻译结束>


<原文开始>
// The device is being disabled.
<原文结束>

# <翻译开始>
// 设备正在被禁用。
# <翻译结束>


<原文开始>
// The properties of the device have changed.
<原文结束>

# <翻译开始>
// 设备的属性已更改。
# <翻译结束>


<原文开始>
// The device is being started (if the request is for the currently active hardware profile).
<原文结束>

# <翻译开始>
// 设备正在启动（如果该请求针对当前活跃的硬件配置文件）。
# <翻译结束>


<原文开始>
// The device is being stopped. The driver stack will be unloaded and the CSCONFIGFLAG_DO_NOT_START flag will be set for the device.
<原文结束>

# <翻译开始>
// 设备正在被停止。驱动程序堆栈将被卸载，并且为该设备设置CSCONFIGFLAG_DO_NOT_START标志。
# <翻译结束>


<原文开始>
// DICS_FLAG specifies the scope of a device property change
<原文结束>

# <翻译开始>
// DICS_FLAG 指定设备属性更改的范围
# <翻译结束>


<原文开始>
// make change in all hardware profiles
<原文结束>

# <翻译开始>
// 在所有硬件配置文件中进行更改
# <翻译结束>


<原文开始>
// make change in specified profile only
<原文结束>

# <翻译开始>
// 只在指定配置文件中做出更改
# <翻译结束>


<原文开始>
// 1 or more hardware profile-specific changes to follow (obsolete)
<原文结束>

# <翻译开始>
// 后续跟随 1 个或多个硬件配置特定的变更（已废弃）
# <翻译结束>


<原文开始>
// PropChangeParams is a structure corresponding to a DIF_PROPERTYCHANGE install function.
<原文结束>

# <翻译开始>
// PropChangeParams 是与 DIF_PROPERTYCHANGE 安装函数对应的一个结构体。
# <翻译结束>


<原文开始>
// DI_REMOVEDEVICE specifies the scope of the device removal
<原文结束>

# <翻译开始>
// DI_REMOVEDEVICE 指定了设备移除的范围
# <翻译结束>


<原文开始>
// Make this change in all hardware profiles. Remove information about the device from the registry.
<原文结束>

# <翻译开始>
// 在所有硬件配置中进行此更改。从注册表中删除该设备的相关信息。
# <翻译结束>


<原文开始>
// Make this change to only the hardware profile specified by HwProfile. this flag only applies to root-enumerated devices. When Windows removes the device from the last hardware profile in which it was configured, Windows performs a global removal.
<原文结束>

# <翻译开始>
// 对指定由 HwProfile 的硬件配置文件做出此更改。此标志仅适用于根枚举设备。
// 当 Windows 从设备最后被配置的硬件配置文件中移除该设备时，Windows 将执行全局移除操作。
# <翻译结束>


<原文开始>
// RemoveDeviceParams is a structure corresponding to a DIF_REMOVE install function.
<原文结束>

# <翻译开始>
// RemoveDeviceParams 是与 DIF_REMOVE 安装函数对应的一个结构体。
# <翻译结束>


<原文开始>
// DrvInfoData is driver information structure (member of a driver info list that may be associated with a particular device instance, or (globally) with a device information set)
<原文结束>

# <翻译开始>
// DrvInfoData 是驱动程序信息结构（属于可能与特定设备实例关联，或（全局）与设备信息集关联的驱动程序信息列表的成员）
# <翻译结束>


<原文开始>
// IsNewer method returns true if DrvInfoData date and version is newer than supplied parameters.
<原文结束>

# <翻译开始>
// IsNewer 方法返回真，当 DrvInfoData 的日期和版本比提供的参数更新时。
# <翻译结束>


<原文开始>
// DrvInfoDetailData is driver information details structure (provides detailed information about a particular driver information structure)
<原文结束>

# <翻译开始>
// DrvInfoDetailData 是驾驶员信息详情结构（用于提供关于特定驾驶员信息结构的详细信息）
# <翻译结束>


<原文开始>
// IsCompatible method tests if given hardware ID matches the driver or is listed on the compatible ID list.
<原文结束>

# <翻译开始>
// IsCompatible 方法用于检测给定的硬件ID是否与驱动程序匹配，或者是否列在兼容ID列表中。
# <翻译结束>


<原文开始>
// DICD flags control SetupDiCreateDeviceInfo
<原文结束>

# <翻译开始>
// DICD 标志用于控制 SetupDiCreateDeviceInfo 函数
# <翻译结束>


<原文开始>
// SUOI flags control SetupUninstallOEMInf
<原文结束>

# <翻译开始>
// SUOI 标志用于控制 SetupUninstallOEMInf
# <翻译结束>


<原文开始>
// SPDIT flags to distinguish between class drivers and
// device drivers. (Passed in 'DriverType' parameter of
// driver information list APIs)
<原文结束>

# <翻译开始>
// SPDIT标志用于区分类驱动程序和设备驱动程序。
// （作为“DriverType”参数传递给驱动程序信息列表API）
# <翻译结束>


<原文开始>
// DIGCF flags control what is included in the device information set built by SetupDiGetClassDevs
<原文结束>

# <翻译开始>
// DIGCF 标志用于控制由 SetupDiGetClassDevs 函数构建的设备信息集中所包含的内容
# <翻译结束>


<原文开始>
// only valid with DIGCF_DEVICEINTERFACE
<原文结束>

# <翻译开始>
// 仅在 DIGCF_DEVICEINTERFACE 下有效
# <翻译结束>


<原文开始>
// DIREG specifies values for SetupDiCreateDevRegKey, SetupDiOpenDevRegKey, and SetupDiDeleteDevRegKey.
<原文结束>

# <翻译开始>
// DIREG 指定用于 SetupDiCreateDevRegKey、SetupDiOpenDevRegKey 和 SetupDiDeleteDevRegKey 的值。
# <翻译结束>


<原文开始>
// Open/Create/Delete device key
<原文结束>

# <翻译开始>
// 打开/创建/删除设备密钥
# <翻译结束>


<原文开始>
// Open/Create/Delete driver key
<原文结束>

# <翻译开始>
// 打开/创建/删除驱动键
# <翻译结束>


<原文开始>
// Delete both driver and Device key
<原文结束>

# <翻译开始>
// 删除driver和Device键
# <翻译结束>


<原文开始>
// SPDRP specifies device registry property codes
// (Codes marked as read-only (R) may only be used for
// SetupDiGetDeviceRegistryProperty)
//
// These values should cover the same set of registry properties
// as defined by the CM_DRP codes in cfgmgr32.h.
//
// Note that SPDRP codes are zero based while CM_DRP codes are one based!
<原文结束>

# <翻译开始>
// SPDRP 定义设备注册表属性代码
// （标注为只读（R）的代码仅可用于 SetupDiGetDeviceRegistryProperty）
//
// 这些值应覆盖由 cfgmgr32.h 中定义的 CM_DRP 代码所表示的相同注册表属性集合。
//
// 注意：SPDRP 代码从零开始编号，而 CM_DRP 代码从一开始编号！
# <翻译结束>


<原文开始>
// Class (R--tied to ClassGUID)
<原文结束>

# <翻译开始>
// 类（与ClassGUID关联的R-）
# <翻译结束>


<原文开始>
// LocationInformation (R/W)
<原文结束>

# <翻译开始>
// LocationInformation（可读/可写）
# <翻译结束>


<原文开始>
// PhysicalDeviceObjectName (R)
<原文结束>

# <翻译开始>
// PhysicalDeviceObjectName (R)
# <翻译结束>


<原文开始>
// Security (R/W, binary form)
<原文结束>

# <翻译开始>
// 安全性（读/写，二进制形式）
# <翻译结束>


<原文开始>
// Device is exclusive-access (R/W)
<原文结束>

# <翻译开始>
// 设备为独占访问（读/写）
# <翻译结束>


<原文开始>
// Device Characteristics (R/W)
<原文结束>

# <翻译开始>
// 设备特性（读/写）
# <翻译结束>


<原文开始>
// UiNumberDescFormat (R/W)
<原文结束>

# <翻译开始>
// UiNumberDescFormat (可读/写)
# <翻译结束>


<原文开始>
// Hardware Removal Policy (R)
<原文结束>

# <翻译开始>
// 硬件移除策略 (R)
# <翻译结束>


<原文开始>
// Removal Policy Override (RW)
<原文结束>

# <翻译开始>
// 移除策略覆盖（可读写）
# <翻译结束>


<原文开始>
// Device Install State (R)
<原文结束>

# <翻译开始>
// 设备安装状态 (只读)
# <翻译结束>


<原文开始>
// Device Location Paths (R)
<原文结束>

# <翻译开始>
// 设备位置路径 (只读)
# <翻译结束>


<原文开始>
// Upper bound on ordinals
<原文结束>

# <翻译开始>
// 对序数的上限
# <翻译结束>


<原文开始>
// DEVPROPTYPE represents the property-data-type identifier that specifies the
// data type of a device property value in the unified device property model.
<原文结束>

# <翻译开始>
// DEVPROPTYPE 表示统一设备属性模型中设备属性值的数据类型标识符。
# <翻译结束>


<原文开始>
// DEVPROPGUID specifies a property category.
<原文结束>

# <翻译开始>
// DEVPROPGUID 定义了一个属性类别。
# <翻译结束>


<原文开始>
// DEVPROPID uniquely identifies the property within the property category.
<原文结束>

# <翻译开始>
// DEVPROPID 用于唯一标识属性类别内的属性。
# <翻译结束>


<原文开始>
// DEVPROPKEY represents a device property key for a device property in the
// unified device property model.
<原文结束>

# <翻译开始>
// DEVPROPKEY 代表统一设备属性模型中设备属性的设备属性键。
# <翻译结束>


<原文开始>
// CONFIGRET is a return value or error code from cfgmgr32 APIs
<原文结束>

# <翻译开始>
// CONFIGRET 是 cfgmgr32 API 返回的值或错误代码
# <翻译结束>


<原文开始>
// only currently 'live' device interfaces
<原文结束>

# <翻译开始>
// 仅包含当前“活跃”的设备接口
# <翻译结束>


<原文开始>
// all registered device interfaces, live or not
<原文结束>

# <翻译开始>
// 所有已注册的设备接口，无论是否活跃
# <翻译结束>


<原文开始>
// Has Register_Device_Driver
<原文结束>

# <翻译开始>
// 是否已注册设备驱动
# <翻译结束>


<原文开始>
// Has Register_Enumerator
<原文结束>

# <翻译开始>
// 具有 Register_Enumerator
# <翻译结束>


<原文开始>
// Is currently configured
<原文结束>

# <翻译开始>
// 当前已配置
# <翻译结束>


<原文开始>
// Enum generates hardware ID
<原文结束>

# <翻译开始>
// Enum 生成硬件ID
# <翻译结束>


<原文开始>
// Lied about can reconfig once
<原文结束>

# <翻译开始>
// 伪造了可重新配置一次的事实
# <翻译结束>


<原文开始>
// Not CM_Create_DevInst lately
<原文结束>

# <翻译开始>
// 最近没有调用 CM_Create_DevInst
# <翻译结束>


<原文开始>
// DevInst is being removed
<原文结束>

# <翻译开始>
// DevInst 正在被移除
# <翻译结束>


<原文开始>
// Has received a config enumerate
<原文结束>

# <翻译开始>
// 已收到配置枚举
# <翻译结束>


<原文开始>
// When child is stopped, free resources
<原文结束>

# <翻译开始>
// 当子进程停止时，释放资源
# <翻译结束>


<原文开始>
// Devnode need lock resume processing
<原文结束>

# <翻译开始>
// 需要锁定以进行devnode的恢复处理
# <翻译结束>


<原文开始>
// Devnode can be the wakeup device
<原文结束>

# <翻译开始>
// Devnode 可作为唤醒设备
# <翻译结束>


<原文开始>
// No show in device manager
<原文结束>

# <翻译开始>
// 在设备管理器中无显示
# <翻译结束>


<原文开始>
// Had a problem during preassignment of boot log conf
<原文结束>

# <翻译开始>
// 在预分配启动日志配置时遇到问题
# <翻译结束>


<原文开始>
// System needs to be restarted for this Devnode to work properly
<原文结束>

# <翻译开始>
// 该Devnode正常工作需要系统重启
# <翻译结束>


<原文开始>
// One or more drivers are blocked from loading for this Devnode
<原文结束>

# <翻译开始>
// 一个或多个驱动程序被阻止为此 Devnode 加载
# <翻译结束>


<原文开始>
// This device is using a legacy driver
<原文结束>

# <翻译开始>
// 该设备正在使用一个旧版驱动程序
# <翻译结束>


<原文开始>
// One or more children have invalid IDs
<原文结束>

# <翻译开始>
// 一个或多个子元素具有无效ID
# <翻译结束>


<原文开始>
// The function driver for a device reported that the device is not connected.  Typically this means a wireless device is out of range.
<原文结束>

# <翻译开始>
// 设备的驱动函数报告该设备未连接。通常这意味着无线设备已超出范围。
# <翻译结束>


<原文开始>
// Device is part of a set of related devices collectively pending query-removal
<原文结束>

# <翻译开始>
// Device 是一组待查询移除的关联设备集合中的一部分
# <翻译结束>


<原文开始>
// Device is actively engaged in a query-remove IRP
<原文结束>

# <翻译开始>
// 设备当前正忙于处理一个查询-移除IRP
# <翻译结束>


<原文开始>
//sys	setupDiCreateDeviceInfoListEx(classGUID *GUID, hwndParent uintptr, machineName *uint16, reserved uintptr) (handle DevInfo, err error) [failretval==DevInfo(InvalidHandle)] = setupapi.SetupDiCreateDeviceInfoListExW
<原文结束>

# <翻译开始>
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
# <翻译结束>


<原文开始>
// SetupDiCreateDeviceInfoListEx function creates an empty device information set on a remote or a local computer and optionally associates the set with a device setup class.
<原文结束>

# <翻译开始>
// SetupDiCreateDeviceInfoListEx 函数在远程或本地计算机上创建一个空的设备信息集合，并可选择性地将该集合与设备安装类相关联。
# <翻译结束>


<原文开始>
//sys	setupDiGetDeviceInfoListDetail(deviceInfoSet DevInfo, deviceInfoSetDetailData *DevInfoListDetailData) (err error) = setupapi.SetupDiGetDeviceInfoListDetailW
<原文结束>

# <翻译开始>
//sys	获取设备信息列表详细信息(deviceInfoSet DevInfo, deviceInfoSetDetailData *DevInfoListDetailData) (错误 error) = setupapi.SetupDiGetDeviceInfoListDetailW
# <翻译结束>


<原文开始>
// SetupDiGetDeviceInfoListDetail function retrieves information associated with a device information set including the class GUID, remote computer handle, and remote computer name.
<原文结束>

# <翻译开始>
// SetupDiGetDeviceInfoListDetail 函数用于获取与设备信息集关联的信息，包括类 GUID、远程计算机句柄以及远程计算机名称。
# <翻译结束>


<原文开始>
// DeviceInfoListDetail method retrieves information associated with a device information set including the class GUID, remote computer handle, and remote computer name.
<原文结束>

# <翻译开始>
// DeviceInfoListDetail 方法用于获取与设备信息集关联的信息，包括类 GUID、远程计算机句柄以及远程计算机名。
# <翻译结束>


<原文开始>
//sys	setupDiCreateDeviceInfo(deviceInfoSet DevInfo, DeviceName *uint16, classGUID *GUID, DeviceDescription *uint16, hwndParent uintptr, CreationFlags DICD, deviceInfoData *DevInfoData) (err error) = setupapi.SetupDiCreateDeviceInfoW
<原文结束>

# <翻译开始>
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
# <翻译结束>


<原文开始>
// SetupDiCreateDeviceInfo function creates a new device information element and adds it as a new member to the specified device information set.
<原文结束>

# <翻译开始>
// SetupDiCreateDeviceInfo函数用于创建一个新的设备信息项，并将其作为新成员添加到指定的设备信息集中。
# <翻译结束>


<原文开始>
// CreateDeviceInfo method creates a new device information element and adds it as a new member to the specified device information set.
<原文结束>

# <翻译开始>
// CreateDeviceInfo 方法用于创建一个新的设备信息元素，并将其作为新成员添加到指定的设备信息集中。
# <翻译结束>


<原文开始>
//sys	setupDiEnumDeviceInfo(deviceInfoSet DevInfo, memberIndex uint32, deviceInfoData *DevInfoData) (err error) = setupapi.SetupDiEnumDeviceInfo
<原文结束>

# <翻译开始>
//sys	枚举设备信息（setupDiEnumDeviceInfo）函数，接受以下参数：
//  - deviceInfoSet：一个DevInfo类型的句柄，表示要从中进行设备枚举的信息集合。
//  - memberIndex：一个uint32类型的整数，指定要获取的设备信息在集合中的索引位置。
//  - deviceInfoData：指向DevInfoData结构体的指针，用于接收所枚举到的设备信息。
//
// 函数调用setupapi库中的SetupDiEnumDeviceInfo函数，并返回一个错误类型（err），表示执行结果。
# <翻译结束>


<原文开始>
// SetupDiEnumDeviceInfo function returns a DevInfoData structure that specifies a device information element in a device information set.
<原文结束>

# <翻译开始>
// SetupDiEnumDeviceInfo 函数返回一个 DevInfoData 结构，该结构用于指定设备信息集中的一项设备信息元素。
# <翻译结束>


<原文开始>
// EnumDeviceInfo method returns a DevInfoData structure that specifies a device information element in a device information set.
<原文结束>

# <翻译开始>
// EnumDeviceInfo 方法返回一个 DevInfoData 结构，该结构用于指定设备信息集中某个设备信息元素。
# <翻译结束>


<原文开始>
// SetupDiDestroyDeviceInfoList function deletes a device information set and frees all associated memory.
//sys	SetupDiDestroyDeviceInfoList(deviceInfoSet DevInfo) (err error) = setupapi.SetupDiDestroyDeviceInfoList
<原文结束>

# <翻译开始>
// SetupDiDestroyDeviceInfoList 函数用于删除一个设备信息集，并释放与其关联的所有内存。
//sys	SetupDiDestroyDeviceInfoList(deviceInfoSet DevInfo) (err error) = setupapi.SetupDiDestroyDeviceInfoList
# <翻译结束>


<原文开始>
// Close method deletes a device information set and frees all associated memory.
<原文结束>

# <翻译开始>
// Close方法用于删除一个设备信息集并释放所有相关的内存。
# <翻译结束>


<原文开始>
//sys	SetupDiBuildDriverInfoList(deviceInfoSet DevInfo, deviceInfoData *DevInfoData, driverType SPDIT) (err error) = setupapi.SetupDiBuildDriverInfoList
<原文结束>

# <翻译开始>
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
# <翻译结束>


<原文开始>
// BuildDriverInfoList method builds a list of drivers that is associated with a specific device or with the global class driver list for a device information set.
<原文结束>

# <翻译开始>
// BuildDriverInfoList 方法用于为特定设备或设备信息集中全局类驱动列表构建关联的驱动程序列表。
# <翻译结束>


<原文开始>
//sys	SetupDiCancelDriverInfoSearch(deviceInfoSet DevInfo) (err error) = setupapi.SetupDiCancelDriverInfoSearch
<原文结束>

# <翻译开始>
//sys	SetupDiCancelDriverInfoSearch(deviceInfoSet DevInfo) (err error) = setupapi.SetupDiCancelDriverInfoSearch
// 
// 系统调用 SetupDiCancelDriverInfoSearch，接收一个 deviceInfoSet 类型的参数 DevInfo，并返回一个错误 err。该函数等同于使用 setupapi 库中的 SetupDiCancelDriverInfoSearch 函数。
# <翻译结束>


<原文开始>
// CancelDriverInfoSearch method cancels a driver list search that is currently in progress in a different thread.
<原文结束>

# <翻译开始>
// CancelDriverInfoSearch 方法用于取消在另一个线程中正在进行的驾驶员信息列表搜索。
# <翻译结束>


<原文开始>
//sys	setupDiEnumDriverInfo(deviceInfoSet DevInfo, deviceInfoData *DevInfoData, driverType SPDIT, memberIndex uint32, driverInfoData *DrvInfoData) (err error) = setupapi.SetupDiEnumDriverInfoW
<原文结束>

# <翻译开始>
//sys	枚举设备驱动信息(设备信息集 deviceInfoSet, 设备信息数据指针 deviceInfoData, 驱动类型 driverType, 成员索引 memberIndex, 驱动信息数据指针 driverInfoData) (返回错误 err) = setupapi.SetupDiEnumDriverInfoW
# <翻译结束>


<原文开始>
// SetupDiEnumDriverInfo function enumerates the members of a driver list.
<原文结束>

# <翻译开始>
// SetupDiEnumDriverInfo 函数用于枚举驱动列表中的成员。
# <翻译结束>


<原文开始>
// EnumDriverInfo method enumerates the members of a driver list.
<原文结束>

# <翻译开始>
// EnumDriverInfo 方法枚举驱动程序列表中的成员。
# <翻译结束>


<原文开始>
//sys	setupDiGetSelectedDriver(deviceInfoSet DevInfo, deviceInfoData *DevInfoData, driverInfoData *DrvInfoData) (err error) = setupapi.SetupDiGetSelectedDriverW
<原文结束>

# <翻译开始>
//sys	设置Di获取选中的驱动程序（deviceInfoSet DevInfo, deviceInfoData *DevInfoData, driverInfoData *DrvInfoData）(返回错误 err) = setupapi.SetupDiGetSelectedDriverW
# <翻译结束>


<原文开始>
// SetupDiGetSelectedDriver function retrieves the selected driver for a device information set or a particular device information element.
<原文结束>

# <翻译开始>
// SetupDiGetSelectedDriver 函数用于为设备信息集或特定设备信息元素检索所选驱动程序。
# <翻译结束>


<原文开始>
// SelectedDriver method retrieves the selected driver for a device information set or a particular device information element.
<原文结束>

# <翻译开始>
// SelectedDriver 方法用于获取设备信息集或特定设备信息项中已选择的驱动程序。
# <翻译结束>


<原文开始>
//sys	SetupDiSetSelectedDriver(deviceInfoSet DevInfo, deviceInfoData *DevInfoData, driverInfoData *DrvInfoData) (err error) = setupapi.SetupDiSetSelectedDriverW
<原文结束>

# <翻译开始>
//sys	SetupDiSetSelectedDriver(deviceInfoSet DevInfo, deviceInfoData *DevInfoData, driverInfoData *DrvInfoData) (err error) = setupapi.SetupDiSetSelectedDriverW
// 
// 系统调用 SetupDiSetSelectedDriver 用于为指定的设备信息集（deviceInfoSet）中的设备（由 deviceInfoData 指针指向其相关信息）设置选中的驱动程序。该驱动程序的相关信息由 driverInfoData 指针所指向。此函数为 Windows API 中 setupapi 库的一部分，具体对应于名为 SetupDiSetSelectedDriverW 的函数。函数执行结果返回一个错误值（err），用于表示操作是否成功。
# <翻译结束>


<原文开始>
// SetSelectedDriver method sets, or resets, the selected driver for a device information element or the selected class driver for a device information set.
<原文结束>

# <翻译开始>
// SetSelectedDriver 方法用于设置或重置设备信息元素的选定驱动程序，或为设备信息集设置选定的类驱动程序。
# <翻译结束>


<原文开始>
//sys	setupDiGetDriverInfoDetail(deviceInfoSet DevInfo, deviceInfoData *DevInfoData, driverInfoData *DrvInfoData, driverInfoDetailData *DrvInfoDetailData, driverInfoDetailDataSize uint32, requiredSize *uint32) (err error) = setupapi.SetupDiGetDriverInfoDetailW
<原文结束>

# <翻译开始>
//sys	设置Di获取驱动信息详细（deviceInfoSet DevInfo, deviceInfoData *DevInfoData, driverInfoData *DrvInfoData, driverInfoDetailData *DrvInfoDetailData, driverInfoDetailDataSize uint32, requiredSize *uint32）(err error) = setupapi.SetupDiGetDriverInfoDetailW
# <翻译结束>


<原文开始>
// SetupDiGetDriverInfoDetail function retrieves driver information detail for a device information set or a particular device information element in the device information set.
<原文结束>

# <翻译开始>
// SetupDiGetDriverInfoDetail 函数用于为设备信息集中或该集合中特定设备信息元素获取驱动程序详细信息。
# <翻译结束>


<原文开始>
// DriverInfoDetail method retrieves driver information detail for a device information set or a particular device information element in the device information set.
<原文结束>

# <翻译开始>
// DriverInfoDetail 方法用于检索设备信息集中某个设备信息项的驱动程序详细信息，或者为设备信息集中的特定设备信息元素获取驱动程序信息。
# <翻译结束>


<原文开始>
//sys	SetupDiDestroyDriverInfoList(deviceInfoSet DevInfo, deviceInfoData *DevInfoData, driverType SPDIT) (err error) = setupapi.SetupDiDestroyDriverInfoList
<原文结束>

# <翻译开始>
//sys	SetupDiDestroyDriverInfoList(deviceInfoSet DevInfo, deviceInfoData *DevInfoData, driverType SPDIT) (err error) = setupapi.SetupDiDestroyDriverInfoList
// 
// 系统调用 SetupDiDestroyDriverInfoList，用于销毁与指定设备信息集（deviceInfoSet）关联的驱动程序信息列表。参数说明如下：
//   - deviceInfoSet：类型为 DevInfo，表示需要销毁其驱动程序信息列表的设备信息集。
//   - deviceInfoData：类型为 *DevInfoData，指向包含设备特定配置信息的结构体指针。此参数用于确定要销毁的驱动程序信息列表。
//   - driverType：类型为 SPDIT，定义了要销毁的驱动程序信息类型的枚举值。
// 
// 此函数映射到名为 setupapi.SetupDiDestroyDriverInfoList 的 WinAPI 函数，执行后返回一个错误（err），用于指示操作是否成功完成。
# <翻译结束>


<原文开始>
// DestroyDriverInfoList method deletes a driver list.
<原文结束>

# <翻译开始>
// DestroyDriverInfoList 方法用于删除一个驱动列表。
# <翻译结束>


<原文开始>
//sys	setupDiGetClassDevsEx(classGUID *GUID, Enumerator *uint16, hwndParent uintptr, Flags DIGCF, deviceInfoSet DevInfo, machineName *uint16, reserved uintptr) (handle DevInfo, err error) [failretval==DevInfo(InvalidHandle)] = setupapi.SetupDiGetClassDevsExW
<原文结束>

# <翻译开始>
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
# <翻译结束>


<原文开始>
// SetupDiGetClassDevsEx function returns a handle to a device information set that contains requested device information elements for a local or a remote computer.
<原文结束>

# <翻译开始>
// SetupDiGetClassDevsEx函数返回一个设备信息集的句柄，该集合包含了针对本地或远程计算机请求的设备信息元素。
# <翻译结束>


<原文开始>
// SetupDiCallClassInstaller function calls the appropriate class installer, and any registered co-installers, with the specified installation request (DIF code).
//sys	SetupDiCallClassInstaller(installFunction DI_FUNCTION, deviceInfoSet DevInfo, deviceInfoData *DevInfoData) (err error) = setupapi.SetupDiCallClassInstaller
<原文结束>

# <翻译开始>
// SetupDiCallClassInstaller 函数以指定的安装请求（DIF 代码）调用相应类安装程序及所有已注册的协同安装程序。
//sys	SetupDiCallClassInstaller(installFunction DI_FUNCTION, deviceInfoSet DevInfo, deviceInfoData *DevInfoData) (err error) = setupapi.SetupDiCallClassInstaller
# <翻译结束>


<原文开始>
// CallClassInstaller member calls the appropriate class installer, and any registered co-installers, with the specified installation request (DIF code).
<原文结束>

# <翻译开始>
// CallClassInstaller 成员调用相应类安装程序及已注册的所有共同安装程序，同时指定了安装请求（DIF 代码）。
# <翻译结束>


<原文开始>
// SetupDiOpenDevRegKey function opens a registry key for device-specific configuration information.
//sys	SetupDiOpenDevRegKey(deviceInfoSet DevInfo, deviceInfoData *DevInfoData, Scope DICS_FLAG, HwProfile uint32, KeyType DIREG, samDesired uint32) (key Handle, err error) [failretval==InvalidHandle] = setupapi.SetupDiOpenDevRegKey
<原文结束>

# <翻译开始>
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
# <翻译结束>


<原文开始>
// OpenDevRegKey method opens a registry key for device-specific configuration information.
<原文结束>

# <翻译开始>
// OpenDevRegKey 方法用于打开一个注册表键，以便访问设备特定的配置信息。
# <翻译结束>


<原文开始>
//sys	setupDiGetDeviceProperty(deviceInfoSet DevInfo, deviceInfoData *DevInfoData, propertyKey *DEVPROPKEY, propertyType *DEVPROPTYPE, propertyBuffer *byte, propertyBufferSize uint32, requiredSize *uint32, flags uint32) (err error) = setupapi.SetupDiGetDevicePropertyW
<原文结束>

# <翻译开始>
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
# <翻译结束>


<原文开始>
// SetupDiGetDeviceProperty function retrieves a specified device instance property.
<原文结束>

# <翻译开始>
// SetupDiGetDeviceProperty 函数用于检索指定设备实例的属性。
# <翻译结束>


<原文开始>
//sys	setupDiGetDeviceRegistryProperty(deviceInfoSet DevInfo, deviceInfoData *DevInfoData, property SPDRP, propertyRegDataType *uint32, propertyBuffer *byte, propertyBufferSize uint32, requiredSize *uint32) (err error) = setupapi.SetupDiGetDeviceRegistryPropertyW
<原文结束>

# <翻译开始>
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
# <翻译结束>


<原文开始>
// SetupDiGetDeviceRegistryProperty function retrieves a specified Plug and Play device property.
<原文结束>

# <翻译开始>
// SetupDiGetDeviceRegistryProperty 函数用于检索指定即插即用设备的属性。
# <翻译结束>


<原文开始>
// bufToUTF16 function reinterprets []byte buffer as []uint16
<原文结束>

# <翻译开始>
// bufToUTF16 函数将 []byte 缓冲区重新解释为 []uint16
# <翻译结束>


<原文开始>
// utf16ToBuf function reinterprets []uint16 as []byte
<原文结束>

# <翻译开始>
// utf16ToBuf 函数将 []uint16 重新解释为 []byte
# <翻译结束>


<原文开始>
// DeviceRegistryProperty method retrieves a specified Plug and Play device property.
<原文结束>

# <翻译开始>
// DeviceRegistryProperty 方法用于检索指定的即插即用设备属性。
# <翻译结束>


<原文开始>
//sys	setupDiSetDeviceRegistryProperty(deviceInfoSet DevInfo, deviceInfoData *DevInfoData, property SPDRP, propertyBuffer *byte, propertyBufferSize uint32) (err error) = setupapi.SetupDiSetDeviceRegistryPropertyW
<原文结束>

# <翻译开始>
//sys	setupDiSetDeviceRegistryProperty(deviceInfoSet DevInfo, deviceInfoData *DevInfoData, property SPDRP, propertyBuffer *byte, propertyBufferSize uint32) (err error) = setupapi.SetupDiSetDeviceRegistryPropertyW
// 
// 系统调用setupDiSetDeviceRegistryProperty，用于设置设备注册表属性。参数说明如下：
//   deviceInfoSet：一个DevInfo类型的句柄，标识设备信息集。
//   deviceInfoData：指向DevInfoData结构体的指针，包含要设置其注册表属性的设备相关信息。
//   property：SPDRP类型的枚举值，指定要设置的设备注册表属性。
//   propertyBuffer：指向字节数据的指针，提供待写入注册表的属性值。
//   propertyBufferSize：uint32类型，表示propertyBuffer指向的数据大小（以字节为单位）。
// 函数返回一个错误接口(err error)，用于报告执行结果。此系统调用对应于Windows API函数setupapi.SetupDiSetDeviceRegistryPropertyW。
# <翻译结束>


<原文开始>
// SetupDiSetDeviceRegistryProperty function sets a Plug and Play device property for a device.
<原文结束>

# <翻译开始>
// SetupDiSetDeviceRegistryProperty 函数用于为设备设置即插即用设备属性。
# <翻译结束>


<原文开始>
// SetDeviceRegistryProperty function sets a Plug and Play device property for a device.
<原文结束>

# <翻译开始>
// SetDeviceRegistryProperty 函数用于为设备设置即插即用设备属性。
# <翻译结束>


<原文开始>
// SetDeviceRegistryPropertyString method sets a Plug and Play device property string for a device.
<原文结束>

# <翻译开始>
// SetDeviceRegistryPropertyString 方法为设备设置即插即用设备属性字符串。
# <翻译结束>


<原文开始>
//sys	setupDiGetDeviceInstallParams(deviceInfoSet DevInfo, deviceInfoData *DevInfoData, deviceInstallParams *DevInstallParams) (err error) = setupapi.SetupDiGetDeviceInstallParamsW
<原文结束>

# <翻译开始>
//sys	获取设备安装参数（setupDiGetDeviceInstallParams）函数通过deviceInfoSet（设备信息集）、deviceInfoData（设备信息数据指针）和deviceInstallParams（设备安装参数指针）三个参数，从setupapi库中调用SetupDiGetDeviceInstallParamsW方法，返回操作结果的错误信息（err）。
# <翻译结束>


<原文开始>
// SetupDiGetDeviceInstallParams function retrieves device installation parameters for a device information set or a particular device information element.
<原文结束>

# <翻译开始>
// SetupDiGetDeviceInstallParams 函数用于为一个设备信息集或特定的设备信息元素检索设备安装参数。
# <翻译结束>


<原文开始>
// DeviceInstallParams method retrieves device installation parameters for a device information set or a particular device information element.
<原文结束>

# <翻译开始>
// DeviceInstallParams 方法用于获取设备信息集或特定设备信息元素的设备安装参数。
# <翻译结束>


<原文开始>
//sys	setupDiGetDeviceInstanceId(deviceInfoSet DevInfo, deviceInfoData *DevInfoData, instanceId *uint16, instanceIdSize uint32, instanceIdRequiredSize *uint32) (err error) = setupapi.SetupDiGetDeviceInstanceIdW
<原文结束>

# <翻译开始>
//sys	获取设备实例ID（setupDiGetDeviceInstanceId）函数，接收如下参数：
//		deviceInfoSet：一个DevInfo类型的设备信息集合；
//		deviceInfoData：指向DevInfoData结构体的指针，用于存储设备相关信息；
//		instanceId：指向uint16类型的指针，用于接收设备实例ID字符串；
//		instanceIdSize：一个uint32类型的变量，表示instanceId缓冲区的大小（以字节为单位）；
//		instanceIdRequiredSize：指向uint32类型的指针，用于返回实际所需的instanceId缓冲区大小。
//此函数是封装了setupapi库中的SetupDiGetDeviceInstanceIdW函数，并可能返回一个错误（err）。
# <翻译结束>


<原文开始>
// SetupDiGetDeviceInstanceId function retrieves the instance ID of the device.
<原文结束>

# <翻译开始>
// SetupDiGetDeviceInstanceId 函数用于获取设备的实例ID。
# <翻译结束>


<原文开始>
// DeviceInstanceID method retrieves the instance ID of the device.
<原文结束>

# <翻译开始>
// DeviceInstanceID 方法用于获取设备的实例ID。
# <翻译结束>


<原文开始>
// SetupDiGetClassInstallParams function retrieves class installation parameters for a device information set or a particular device information element.
//sys	SetupDiGetClassInstallParams(deviceInfoSet DevInfo, deviceInfoData *DevInfoData, classInstallParams *ClassInstallHeader, classInstallParamsSize uint32, requiredSize *uint32) (err error) = setupapi.SetupDiGetClassInstallParamsW
<原文结束>

# <翻译开始>
// SetupDiGetClassInstallParams 函数用于为设备信息集或特定设备信息元素检索类安装参数。
//sys	SetupDiGetClassInstallParams(deviceInfoSet DevInfo, deviceInfoData *DevInfoData, classInstallParams *ClassInstallHeader, classInstallParamsSize uint32, requiredSize *uint32) (err error) = setupapi.SetupDiGetClassInstallParamsW
# <翻译结束>


<原文开始>
// ClassInstallParams method retrieves class installation parameters for a device information set or a particular device information element.
<原文结束>

# <翻译开始>
// ClassInstallParams 方法用于为设备信息集或特定设备信息元素检索类安装参数。
# <翻译结束>


<原文开始>
//sys	SetupDiSetDeviceInstallParams(deviceInfoSet DevInfo, deviceInfoData *DevInfoData, deviceInstallParams *DevInstallParams) (err error) = setupapi.SetupDiSetDeviceInstallParamsW
<原文结束>

# <翻译开始>
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
# <翻译结束>


<原文开始>
// SetDeviceInstallParams member sets device installation parameters for a device information set or a particular device information element.
<原文结束>

# <翻译开始>
// SetDeviceInstallParams 成员为设备信息集或特定设备信息元素设置设备安装参数。
# <翻译结束>


<原文开始>
// SetupDiSetClassInstallParams function sets or clears class install parameters for a device information set or a particular device information element.
//sys	SetupDiSetClassInstallParams(deviceInfoSet DevInfo, deviceInfoData *DevInfoData, classInstallParams *ClassInstallHeader, classInstallParamsSize uint32) (err error) = setupapi.SetupDiSetClassInstallParamsW
<原文结束>

# <翻译开始>
// SetupDiSetClassInstallParams 函数用于为设备信息集或特定设备信息元素设置或清除类安装参数。
//sys	SetupDiSetClassInstallParams(deviceInfoSet DevInfo, deviceInfoData *DevInfoData, classInstallParams *ClassInstallHeader, classInstallParamsSize uint32) (err error) = setupapi.SetupDiSetClassInstallParamsW
# <翻译结束>


<原文开始>
// SetClassInstallParams method sets or clears class install parameters for a device information set or a particular device information element.
<原文结束>

# <翻译开始>
// SetClassInstallParams 方法用于为一个设备信息集或特定的设备信息元素设置或清除类安装参数。
# <翻译结束>


<原文开始>
//sys	setupDiClassNameFromGuidEx(classGUID *GUID, className *uint16, classNameSize uint32, requiredSize *uint32, machineName *uint16, reserved uintptr) (err error) = setupapi.SetupDiClassNameFromGuidExW
<原文结束>

# <翻译开始>
//sys	根据给定的类GUID、目标类名缓冲区、类名缓冲区大小、返回所需缓冲区大小指针、目标计算机名和预留参数，使用setupapi.SetupDiClassNameFromGuidExW函数获取类名，并返回可能的错误。
# <翻译结束>


<原文开始>
// SetupDiClassNameFromGuidEx function retrieves the class name associated with a class GUID. The class can be installed on a local or remote computer.
<原文结束>

# <翻译开始>
// SetupDiClassNameFromGuidEx 函数用于获取与类GUID关联的类名。此类可以安装在本地或远程计算机上。
# <翻译结束>


<原文开始>
//sys	setupDiClassGuidsFromNameEx(className *uint16, classGuidList *GUID, classGuidListSize uint32, requiredSize *uint32, machineName *uint16, reserved uintptr) (err error) = setupapi.SetupDiClassGuidsFromNameExW
<原文结束>

# <翻译开始>
//sys	根据给定的类名（className）在指定计算机（machineName）上获取类GUID列表，将结果存储于classGuidList中。classGuidListSize参数指示classGuidList的容量（以GUID为单位）。requiredSize用于返回实际需要的GUID数量。若machineName为空，表示在本地计算机查询。reserved参数保留供将来使用。此函数为setupapi库中的SetupDiClassGuidsFromNameExW函数的封装，并在出错时返回错误信息（err）。
# <翻译结束>


<原文开始>
// SetupDiClassGuidsFromNameEx function retrieves the GUIDs associated with the specified class name. This resulting list contains the classes currently installed on a local or remote computer.
<原文结束>

# <翻译开始>
// SetupDiClassGuidsFromNameEx 函数从指定的类名中检索关联的 GUID。返回的列表包含了当前安装在本地或远程计算机上的此类别。
# <翻译结束>


<原文开始>
//sys	setupDiGetSelectedDevice(deviceInfoSet DevInfo, deviceInfoData *DevInfoData) (err error) = setupapi.SetupDiGetSelectedDevice
<原文结束>

# <翻译开始>
//sys	setupDiGetSelectedDevice(deviceInfoSet DevInfo, deviceInfoData *DevInfoData) (err error) = setupapi.SetupDiGetSelectedDevice
// 
// 系统调用 setupDiGetSelectedDevice，用于获取指定设备信息集合（deviceInfoSet）中已被选中的设备信息。参数说明如下：
//   - deviceInfoSet：类型为 DevInfo 的变量，表示要从中获取选中设备的设备信息集合。
//   - deviceInfoData：指向 DevInfoData 结构体的指针，用于接收所选设备的相关信息。
// 
// 该函数返回一个错误值（err），若操作成功则返回 nil；否则返回具体错误。该函数与 Windows API 函数 `setupapi.SetupDiGetSelectedDevice` 对应。
# <翻译结束>


<原文开始>
// SetupDiGetSelectedDevice function retrieves the selected device information element in a device information set.
<原文结束>

# <翻译开始>
// SetupDiGetSelectedDevice 函数用于从设备信息集中检索选定的设备信息元素。
# <翻译结束>


<原文开始>
// SelectedDevice method retrieves the selected device information element in a device information set.
<原文结束>

# <翻译开始>
// SelectedDevice 方法从设备信息集中检索选定的设备信息项。
# <翻译结束>


<原文开始>
// SetupDiSetSelectedDevice function sets a device information element as the selected member of a device information set. This function is typically used by an installation wizard.
//sys	SetupDiSetSelectedDevice(deviceInfoSet DevInfo, deviceInfoData *DevInfoData) (err error) = setupapi.SetupDiSetSelectedDevice
<原文结束>

# <翻译开始>
// SetupDiSetSelectedDevice 函数将一个设备信息元素设置为设备信息集中被选中的成员。此函数通常由安装向导使用。
//sys	SetupDiSetSelectedDevice(deviceInfoSet DevInfo, deviceInfoData *DevInfoData) (err error) = setupapi.SetupDiSetSelectedDevice
# <翻译结束>


<原文开始>
// SetSelectedDevice method sets a device information element as the selected member of a device information set. This function is typically used by an installation wizard.
<原文结束>

# <翻译开始>
// SetSelectedDevice 方法将一个设备信息元素设置为设备信息集的选定成员。此函数通常由安装向导使用。
# <翻译结束>


<原文开始>
//sys	setupUninstallOEMInf(infFileName *uint16, flags SUOI, reserved uintptr) (err error) = setupapi.SetupUninstallOEMInfW
<原文结束>

# <翻译开始>
//sys	setupUninstallOEMInf(infFileName *uint16, flags SUOI, reserved uintptr) (err error) = setupapi.SetupUninstallOEMInfW
// 
// 系统调用 setupUninstallOEMInf，用于卸载OEM信息文件。参数说明如下：
// - infFileName: 指向包含OEM信息文件名的 uint16 类型指针。
// - flags: 类型为 SUOI 的标志位，用于指定卸载操作的相关选项或行为。
// - reserved: 类型为 uintptr 的保留参数，目前未使用，应传入零值。
// 该函数等同于调用 setupapi 库中的 SetupUninstallOEMInfW 函数。
# <翻译结束>


<原文开始>
// SetupUninstallOEMInf uninstalls the specified driver.
<原文结束>

# <翻译开始>
// SetupUninstallOEMInf 卸载指定的驱动程序
# <翻译结束>


<原文开始>
//sys cm_MapCrToWin32Err(configRet CONFIGRET, defaultWin32Error Errno) (ret Errno) = CfgMgr32.CM_MapCrToWin32Err
<原文结束>

# <翻译开始>
//sys cm_MapCrToWin32Err(configRet CONFIGRET, defaultWin32Error Errno) (ret Errno) = CfgMgr32.CM_MapCrToWin32Err
// 
// 系统调用 cm_MapCrToWin32Err，将 CONFIGRET 类型的 configRet 映射到对应的 Win32 错误码。如果映射失败，使用 Errno 类型的 defaultWin32Error 作为默认返回值。该函数等价于调用 CfgMgr32 库中的 CM_MapCrToWin32Err 函数，并返回 Errno 类型的结果 ret。
# <翻译结束>


<原文开始>
//sys cm_Get_Device_Interface_List_Size(len *uint32, interfaceClass *GUID, deviceID *uint16, flags uint32) (ret CONFIGRET) = CfgMgr32.CM_Get_Device_Interface_List_SizeW
//sys cm_Get_Device_Interface_List(interfaceClass *GUID, deviceID *uint16, buffer *uint16, bufferLen uint32, flags uint32) (ret CONFIGRET) = CfgMgr32.CM_Get_Device_Interface_ListW
<原文结束>

# <翻译开始>
//sys cm_Get_Device_Interface_List_Size(len *uint32, interfaceClass *GUID, deviceID *uint16, flags uint32) (ret CONFIGRET) = CfgMgr32.CM_Get_Device_Interface_List_SizeW
//sys cm_Get_Device_Interface_List(interfaceClass *GUID, deviceID *uint16, buffer *uint16, bufferLen uint32, flags uint32) (ret CONFIGRET) = CfgMgr32.CM_Get_Device_Interface_ListW
// 
// 翻译成中文为：
// 
//sys cm_Get_Device_Interface_List_Size(len *uint32, interfaceClass *GUID, deviceID *uint16, flags uint32) (ret CONFIGRET) = CfgMgr32.CM_Get_Device_Interface_List_SizeW
//sys cm_Get_Device_Interface_List(interfaceClass *GUID, deviceID *uint16, buffer *uint16, bufferLen uint32, flags uint32) (ret CONFIGRET) = CfgMgr32.CM_Get_Device_Interface_ListW
// 
// （注：以上为原文，由于原文为Go语言的系统调用声明，其内容已为英文注释形式，无需再进行翻译。）
# <翻译结束>


<原文开始>
//sys cm_Get_DevNode_Status(status *uint32, problemNumber *uint32, devInst DEVINST, flags uint32) (ret CONFIGRET) = CfgMgr32.CM_Get_DevNode_Status
<原文结束>

# <翻译开始>
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
# <翻译结束>

