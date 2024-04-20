
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
// only add the installed driver to the class or compat driver list.  Used in calls to SetupDiBuildDriverInfoList
# <翻译结束>


<原文开始>
// Build driver list based on alternate platform information specified in associated file queue
<原文结束>

# <翻译开始>
// Build driver list based on alternate platform information specified in associated file queue
# <翻译结束>


<原文开始>
// only restart the device drivers are being installed on as opposed to restarting all devices using those drivers.
<原文结束>

# <翻译开始>
// only restart the device drivers are being installed on as opposed to restarting all devices using those drivers.
# <翻译结束>


<原文开始>
// Tell SetupDiBuildDriverInfoList to do a recursive search
<原文结束>

# <翻译开始>
// Tell SetupDiBuildDriverInfoList to do a recursive search
# <翻译结束>


<原文开始>
// Tell SetupDiBuildDriverInfoList to do a "published INF" search
<原文结束>

# <翻译开始>
// Tell SetupDiBuildDriverInfoList to do a "published INF" search
# <翻译结束>


<原文开始>
// ClassInstallHeader is the first member of any class install parameters structure. It contains the device installation request code that defines the format of the rest of the install parameters structure.
<原文结束>

# <翻译开始>
// ClassInstallHeader is the first member of any class install parameters structure. It contains the device installation request code that defines the format of the rest of the install parameters structure.
# <翻译结束>


<原文开始>
// DICS_STATE specifies values indicating a change in a device's state
<原文结束>

# <翻译开始>
// DICS_STATE specifies values indicating a change in a device's state
# <翻译结束>


<原文开始>
// The device is being enabled.
<原文结束>

# <翻译开始>
// The device is being enabled.
# <翻译结束>


<原文开始>
// The device is being disabled.
<原文结束>

# <翻译开始>
// The device is being disabled.
# <翻译结束>


<原文开始>
// The properties of the device have changed.
<原文结束>

# <翻译开始>
// The properties of the device have changed.
# <翻译结束>


<原文开始>
// The device is being started (if the request is for the currently active hardware profile).
<原文结束>

# <翻译开始>
// The device is being started (if the request is for the currently active hardware profile).
# <翻译结束>


<原文开始>
// The device is being stopped. The driver stack will be unloaded and the CSCONFIGFLAG_DO_NOT_START flag will be set for the device.
<原文结束>

# <翻译开始>
// The device is being stopped. The driver stack will be unloaded and the CSCONFIGFLAG_DO_NOT_START flag will be set for the device.
# <翻译结束>


<原文开始>
// DICS_FLAG specifies the scope of a device property change
<原文结束>

# <翻译开始>
// DICS_FLAG specifies the scope of a device property change
# <翻译结束>


<原文开始>
// make change in all hardware profiles
<原文结束>

# <翻译开始>
// make change in all hardware profiles
# <翻译结束>


<原文开始>
// make change in specified profile only
<原文结束>

# <翻译开始>
// make change in specified profile only
# <翻译结束>


<原文开始>
// 1 or more hardware profile-specific changes to follow (obsolete)
<原文结束>

# <翻译开始>
// 1 or more hardware profile-specific changes to follow (obsolete)
# <翻译结束>


<原文开始>
// PropChangeParams is a structure corresponding to a DIF_PROPERTYCHANGE install function.
<原文结束>

# <翻译开始>
// PropChangeParams is a structure corresponding to a DIF_PROPERTYCHANGE install function.
# <翻译结束>


<原文开始>
// DI_REMOVEDEVICE specifies the scope of the device removal
<原文结束>

# <翻译开始>
// DI_REMOVEDEVICE specifies the scope of the device removal
# <翻译结束>


<原文开始>
// Make this change in all hardware profiles. Remove information about the device from the registry.
<原文结束>

# <翻译开始>
// Make this change in all hardware profiles. Remove information about the device from the registry.
# <翻译结束>


<原文开始>
// Make this change to only the hardware profile specified by HwProfile. this flag only applies to root-enumerated devices. When Windows removes the device from the last hardware profile in which it was configured, Windows performs a global removal.
<原文结束>

# <翻译开始>
// Make this change to only the hardware profile specified by HwProfile. this flag only applies to root-enumerated devices. When Windows removes the device from the last hardware profile in which it was configured, Windows performs a global removal.
# <翻译结束>


<原文开始>
// RemoveDeviceParams is a structure corresponding to a DIF_REMOVE install function.
<原文结束>

# <翻译开始>
// RemoveDeviceParams is a structure corresponding to a DIF_REMOVE install function.
# <翻译结束>


<原文开始>
// DrvInfoData is driver information structure (member of a driver info list that may be associated with a particular device instance, or (globally) with a device information set)
<原文结束>

# <翻译开始>
// DrvInfoData is driver information structure (member of a driver info list that may be associated with a particular device instance, or (globally) with a device information set)
# <翻译结束>


<原文开始>
// IsNewer method returns true if DrvInfoData date and version is newer than supplied parameters.
<原文结束>

# <翻译开始>
// IsNewer method returns true if DrvInfoData date and version is newer than supplied parameters.
# <翻译结束>


<原文开始>
// DrvInfoDetailData is driver information details structure (provides detailed information about a particular driver information structure)
<原文结束>

# <翻译开始>
// DrvInfoDetailData is driver information details structure (provides detailed information about a particular driver information structure)
# <翻译结束>


<原文开始>
// IsCompatible method tests if given hardware ID matches the driver or is listed on the compatible ID list.
<原文结束>

# <翻译开始>
// IsCompatible method tests if given hardware ID matches the driver or is listed on the compatible ID list.
# <翻译结束>


<原文开始>
// DICD flags control SetupDiCreateDeviceInfo
<原文结束>

# <翻译开始>
// DICD flags control SetupDiCreateDeviceInfo
# <翻译结束>


<原文开始>
// SUOI flags control SetupUninstallOEMInf
<原文结束>

# <翻译开始>
// SUOI flags control SetupUninstallOEMInf
# <翻译结束>


<原文开始>
// SPDIT flags to distinguish between class drivers and
// device drivers. (Passed in 'DriverType' parameter of
// driver information list APIs)
<原文结束>

# <翻译开始>
// SPDIT flags to distinguish between class drivers and
// device drivers. (Passed in 'DriverType' parameter of
// driver information list APIs)
# <翻译结束>


<原文开始>
// DIGCF flags control what is included in the device information set built by SetupDiGetClassDevs
<原文结束>

# <翻译开始>
// DIGCF flags control what is included in the device information set built by SetupDiGetClassDevs
# <翻译结束>


<原文开始>
// only valid with DIGCF_DEVICEINTERFACE
<原文结束>

# <翻译开始>
// only valid with DIGCF_DEVICEINTERFACE
# <翻译结束>


<原文开始>
// DIREG specifies values for SetupDiCreateDevRegKey, SetupDiOpenDevRegKey, and SetupDiDeleteDevRegKey.
<原文结束>

# <翻译开始>
// DIREG specifies values for SetupDiCreateDevRegKey, SetupDiOpenDevRegKey, and SetupDiDeleteDevRegKey.
# <翻译结束>


<原文开始>
// Open/Create/Delete device key
<原文结束>

# <翻译开始>
// Open/Create/Delete device key
# <翻译结束>


<原文开始>
// Open/Create/Delete driver key
<原文结束>

# <翻译开始>
// Open/Create/Delete driver key
# <翻译结束>


<原文开始>
// Delete both driver and Device key
<原文结束>

# <翻译开始>
// Delete both driver and Device key
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
// SPDRP specifies device registry property codes
// (Codes marked as read-only (R) may only be used for
// SetupDiGetDeviceRegistryProperty)
//
// These values should cover the same set of registry properties
// as defined by the CM_DRP codes in cfgmgr32.h.
//
// Note that SPDRP codes are zero based while CM_DRP codes are one based!
# <翻译结束>


<原文开始>
// Class (R--tied to ClassGUID)
<原文结束>

# <翻译开始>
// Class (R--tied to ClassGUID)
# <翻译结束>


<原文开始>
// LocationInformation (R/W)
<原文结束>

# <翻译开始>
// LocationInformation (R/W)
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
// Security (R/W, binary form)
# <翻译结束>


<原文开始>
// Device is exclusive-access (R/W)
<原文结束>

# <翻译开始>
// Device is exclusive-access (R/W)
# <翻译结束>


<原文开始>
// Device Characteristics (R/W)
<原文结束>

# <翻译开始>
// Device Characteristics (R/W)
# <翻译结束>


<原文开始>
// UiNumberDescFormat (R/W)
<原文结束>

# <翻译开始>
// UiNumberDescFormat (R/W)
# <翻译结束>


<原文开始>
// Hardware Removal Policy (R)
<原文结束>

# <翻译开始>
// Hardware Removal Policy (R)
# <翻译结束>


<原文开始>
// Removal Policy Override (RW)
<原文结束>

# <翻译开始>
// Removal Policy Override (RW)
# <翻译结束>


<原文开始>
// Device Install State (R)
<原文结束>

# <翻译开始>
// Device Install State (R)
# <翻译结束>


<原文开始>
// Device Location Paths (R)
<原文结束>

# <翻译开始>
// Device Location Paths (R)
# <翻译结束>


<原文开始>
// Upper bound on ordinals
<原文结束>

# <翻译开始>
// Upper bound on ordinals
# <翻译结束>


<原文开始>
// DEVPROPTYPE represents the property-data-type identifier that specifies the
// data type of a device property value in the unified device property model.
<原文结束>

# <翻译开始>
// DEVPROPTYPE represents the property-data-type identifier that specifies the
// data type of a device property value in the unified device property model.
# <翻译结束>


<原文开始>
// DEVPROPGUID specifies a property category.
<原文结束>

# <翻译开始>
// DEVPROPGUID specifies a property category.
# <翻译结束>


<原文开始>
// DEVPROPID uniquely identifies the property within the property category.
<原文结束>

# <翻译开始>
// DEVPROPID uniquely identifies the property within the property category.
# <翻译结束>


<原文开始>
// DEVPROPKEY represents a device property key for a device property in the
// unified device property model.
<原文结束>

# <翻译开始>
// DEVPROPKEY represents a device property key for a device property in the
// unified device property model.
# <翻译结束>


<原文开始>
// CONFIGRET is a return value or error code from cfgmgr32 APIs
<原文结束>

# <翻译开始>
// CONFIGRET is a return value or error code from cfgmgr32 APIs
# <翻译结束>


<原文开始>
// only currently 'live' device interfaces
<原文结束>

# <翻译开始>
// only currently 'live' device interfaces
# <翻译结束>


<原文开始>
// all registered device interfaces, live or not
<原文结束>

# <翻译开始>
// all registered device interfaces, live or not
# <翻译结束>


<原文开始>
// Has Register_Device_Driver
<原文结束>

# <翻译开始>
// Has Register_Device_Driver
# <翻译结束>


<原文开始>
// Has Register_Enumerator
<原文结束>

# <翻译开始>
// Has Register_Enumerator
# <翻译结束>


<原文开始>
// Is currently configured
<原文结束>

# <翻译开始>
// Is currently configured
# <翻译结束>


<原文开始>
// Enum generates hardware ID
<原文结束>

# <翻译开始>
// Enum generates hardware ID
# <翻译结束>


<原文开始>
// Lied about can reconfig once
<原文结束>

# <翻译开始>
// Lied about can reconfig once
# <翻译结束>


<原文开始>
// Not CM_Create_DevInst lately
<原文结束>

# <翻译开始>
// Not CM_Create_DevInst lately
# <翻译结束>


<原文开始>
// DevInst is being removed
<原文结束>

# <翻译开始>
// DevInst is being removed
# <翻译结束>


<原文开始>
// Has received a config enumerate
<原文结束>

# <翻译开始>
// Has received a config enumerate
# <翻译结束>


<原文开始>
// When child is stopped, free resources
<原文结束>

# <翻译开始>
// When child is stopped, free resources
# <翻译结束>


<原文开始>
// Devnode need lock resume processing
<原文结束>

# <翻译开始>
// Devnode need lock resume processing
# <翻译结束>


<原文开始>
// Devnode can be the wakeup device
<原文结束>

# <翻译开始>
// Devnode can be the wakeup device
# <翻译结束>


<原文开始>
// No show in device manager
<原文结束>

# <翻译开始>
// No show in device manager
# <翻译结束>


<原文开始>
// Had a problem during preassignment of boot log conf
<原文结束>

# <翻译开始>
// Had a problem during preassignment of boot log conf
# <翻译结束>


<原文开始>
// System needs to be restarted for this Devnode to work properly
<原文结束>

# <翻译开始>
// System needs to be restarted for this Devnode to work properly
# <翻译结束>


<原文开始>
// One or more drivers are blocked from loading for this Devnode
<原文结束>

# <翻译开始>
// One or more drivers are blocked from loading for this Devnode
# <翻译结束>


<原文开始>
// This device is using a legacy driver
<原文结束>

# <翻译开始>
// This device is using a legacy driver
# <翻译结束>


<原文开始>
// One or more children have invalid IDs
<原文结束>

# <翻译开始>
// One or more children have invalid IDs
# <翻译结束>


<原文开始>
// The function driver for a device reported that the device is not connected.  Typically this means a wireless device is out of range.
<原文结束>

# <翻译开始>
// The function driver for a device reported that the device is not connected.  Typically this means a wireless device is out of range.
# <翻译结束>


<原文开始>
// Device is part of a set of related devices collectively pending query-removal
<原文结束>

# <翻译开始>
// Device is part of a set of related devices collectively pending query-removal
# <翻译结束>


<原文开始>
// Device is actively engaged in a query-remove IRP
<原文结束>

# <翻译开始>
// Device is actively engaged in a query-remove IRP
# <翻译结束>


<原文开始>
//sys	setupDiCreateDeviceInfoListEx(classGUID *GUID, hwndParent uintptr, machineName *uint16, reserved uintptr) (handle DevInfo, err error) [failretval==DevInfo(InvalidHandle)] = setupapi.SetupDiCreateDeviceInfoListExW
<原文结束>

# <翻译开始>
//sys	setupDiCreateDeviceInfoListEx(classGUID *GUID, hwndParent uintptr, machineName *uint16, reserved uintptr) (handle DevInfo, err error) [failretval==DevInfo(InvalidHandle)] = setupapi.SetupDiCreateDeviceInfoListExW
# <翻译结束>


<原文开始>
// SetupDiCreateDeviceInfoListEx function creates an empty device information set on a remote or a local computer and optionally associates the set with a device setup class.
<原文结束>

# <翻译开始>
// SetupDiCreateDeviceInfoListEx function creates an empty device information set on a remote or a local computer and optionally associates the set with a device setup class.
# <翻译结束>


<原文开始>
//sys	setupDiGetDeviceInfoListDetail(deviceInfoSet DevInfo, deviceInfoSetDetailData *DevInfoListDetailData) (err error) = setupapi.SetupDiGetDeviceInfoListDetailW
<原文结束>

# <翻译开始>
//sys	setupDiGetDeviceInfoListDetail(deviceInfoSet DevInfo, deviceInfoSetDetailData *DevInfoListDetailData) (err error) = setupapi.SetupDiGetDeviceInfoListDetailW
# <翻译结束>


<原文开始>
// SetupDiGetDeviceInfoListDetail function retrieves information associated with a device information set including the class GUID, remote computer handle, and remote computer name.
<原文结束>

# <翻译开始>
// SetupDiGetDeviceInfoListDetail function retrieves information associated with a device information set including the class GUID, remote computer handle, and remote computer name.
# <翻译结束>


<原文开始>
// DeviceInfoListDetail method retrieves information associated with a device information set including the class GUID, remote computer handle, and remote computer name.
<原文结束>

# <翻译开始>
// DeviceInfoListDetail method retrieves information associated with a device information set including the class GUID, remote computer handle, and remote computer name.
# <翻译结束>


<原文开始>
//sys	setupDiCreateDeviceInfo(deviceInfoSet DevInfo, DeviceName *uint16, classGUID *GUID, DeviceDescription *uint16, hwndParent uintptr, CreationFlags DICD, deviceInfoData *DevInfoData) (err error) = setupapi.SetupDiCreateDeviceInfoW
<原文结束>

# <翻译开始>
//sys	setupDiCreateDeviceInfo(deviceInfoSet DevInfo, DeviceName *uint16, classGUID *GUID, DeviceDescription *uint16, hwndParent uintptr, CreationFlags DICD, deviceInfoData *DevInfoData) (err error) = setupapi.SetupDiCreateDeviceInfoW
# <翻译结束>


<原文开始>
// SetupDiCreateDeviceInfo function creates a new device information element and adds it as a new member to the specified device information set.
<原文结束>

# <翻译开始>
// SetupDiCreateDeviceInfo function creates a new device information element and adds it as a new member to the specified device information set.
# <翻译结束>


<原文开始>
// CreateDeviceInfo method creates a new device information element and adds it as a new member to the specified device information set.
<原文结束>

# <翻译开始>
// CreateDeviceInfo method creates a new device information element and adds it as a new member to the specified device information set.
# <翻译结束>


<原文开始>
//sys	setupDiEnumDeviceInfo(deviceInfoSet DevInfo, memberIndex uint32, deviceInfoData *DevInfoData) (err error) = setupapi.SetupDiEnumDeviceInfo
<原文结束>

# <翻译开始>
//sys	setupDiEnumDeviceInfo(deviceInfoSet DevInfo, memberIndex uint32, deviceInfoData *DevInfoData) (err error) = setupapi.SetupDiEnumDeviceInfo
# <翻译结束>


<原文开始>
// SetupDiEnumDeviceInfo function returns a DevInfoData structure that specifies a device information element in a device information set.
<原文结束>

# <翻译开始>
// SetupDiEnumDeviceInfo function returns a DevInfoData structure that specifies a device information element in a device information set.
# <翻译结束>


<原文开始>
// EnumDeviceInfo method returns a DevInfoData structure that specifies a device information element in a device information set.
<原文结束>

# <翻译开始>
// EnumDeviceInfo method returns a DevInfoData structure that specifies a device information element in a device information set.
# <翻译结束>


<原文开始>
// SetupDiDestroyDeviceInfoList function deletes a device information set and frees all associated memory.
//sys	SetupDiDestroyDeviceInfoList(deviceInfoSet DevInfo) (err error) = setupapi.SetupDiDestroyDeviceInfoList
<原文结束>

# <翻译开始>
// SetupDiDestroyDeviceInfoList function deletes a device information set and frees all associated memory.
//sys	SetupDiDestroyDeviceInfoList(deviceInfoSet DevInfo) (err error) = setupapi.SetupDiDestroyDeviceInfoList
# <翻译结束>


<原文开始>
// Close method deletes a device information set and frees all associated memory.
<原文结束>

# <翻译开始>
// Close method deletes a device information set and frees all associated memory.
# <翻译结束>


<原文开始>
//sys	SetupDiBuildDriverInfoList(deviceInfoSet DevInfo, deviceInfoData *DevInfoData, driverType SPDIT) (err error) = setupapi.SetupDiBuildDriverInfoList
<原文结束>

# <翻译开始>
//sys	SetupDiBuildDriverInfoList(deviceInfoSet DevInfo, deviceInfoData *DevInfoData, driverType SPDIT) (err error) = setupapi.SetupDiBuildDriverInfoList
# <翻译结束>


<原文开始>
// BuildDriverInfoList method builds a list of drivers that is associated with a specific device or with the global class driver list for a device information set.
<原文结束>

# <翻译开始>
// BuildDriverInfoList method builds a list of drivers that is associated with a specific device or with the global class driver list for a device information set.
# <翻译结束>


<原文开始>
//sys	SetupDiCancelDriverInfoSearch(deviceInfoSet DevInfo) (err error) = setupapi.SetupDiCancelDriverInfoSearch
<原文结束>

# <翻译开始>
//sys	SetupDiCancelDriverInfoSearch(deviceInfoSet DevInfo) (err error) = setupapi.SetupDiCancelDriverInfoSearch
# <翻译结束>


<原文开始>
// CancelDriverInfoSearch method cancels a driver list search that is currently in progress in a different thread.
<原文结束>

# <翻译开始>
// CancelDriverInfoSearch method cancels a driver list search that is currently in progress in a different thread.
# <翻译结束>


<原文开始>
//sys	setupDiEnumDriverInfo(deviceInfoSet DevInfo, deviceInfoData *DevInfoData, driverType SPDIT, memberIndex uint32, driverInfoData *DrvInfoData) (err error) = setupapi.SetupDiEnumDriverInfoW
<原文结束>

# <翻译开始>
//sys	setupDiEnumDriverInfo(deviceInfoSet DevInfo, deviceInfoData *DevInfoData, driverType SPDIT, memberIndex uint32, driverInfoData *DrvInfoData) (err error) = setupapi.SetupDiEnumDriverInfoW
# <翻译结束>


<原文开始>
// SetupDiEnumDriverInfo function enumerates the members of a driver list.
<原文结束>

# <翻译开始>
// SetupDiEnumDriverInfo function enumerates the members of a driver list.
# <翻译结束>


<原文开始>
// EnumDriverInfo method enumerates the members of a driver list.
<原文结束>

# <翻译开始>
// EnumDriverInfo method enumerates the members of a driver list.
# <翻译结束>


<原文开始>
//sys	setupDiGetSelectedDriver(deviceInfoSet DevInfo, deviceInfoData *DevInfoData, driverInfoData *DrvInfoData) (err error) = setupapi.SetupDiGetSelectedDriverW
<原文结束>

# <翻译开始>
//sys	setupDiGetSelectedDriver(deviceInfoSet DevInfo, deviceInfoData *DevInfoData, driverInfoData *DrvInfoData) (err error) = setupapi.SetupDiGetSelectedDriverW
# <翻译结束>


<原文开始>
// SetupDiGetSelectedDriver function retrieves the selected driver for a device information set or a particular device information element.
<原文结束>

# <翻译开始>
// SetupDiGetSelectedDriver function retrieves the selected driver for a device information set or a particular device information element.
# <翻译结束>


<原文开始>
// SelectedDriver method retrieves the selected driver for a device information set or a particular device information element.
<原文结束>

# <翻译开始>
// SelectedDriver method retrieves the selected driver for a device information set or a particular device information element.
# <翻译结束>


<原文开始>
//sys	SetupDiSetSelectedDriver(deviceInfoSet DevInfo, deviceInfoData *DevInfoData, driverInfoData *DrvInfoData) (err error) = setupapi.SetupDiSetSelectedDriverW
<原文结束>

# <翻译开始>
//sys	SetupDiSetSelectedDriver(deviceInfoSet DevInfo, deviceInfoData *DevInfoData, driverInfoData *DrvInfoData) (err error) = setupapi.SetupDiSetSelectedDriverW
# <翻译结束>


<原文开始>
// SetSelectedDriver method sets, or resets, the selected driver for a device information element or the selected class driver for a device information set.
<原文结束>

# <翻译开始>
// SetSelectedDriver method sets, or resets, the selected driver for a device information element or the selected class driver for a device information set.
# <翻译结束>


<原文开始>
//sys	setupDiGetDriverInfoDetail(deviceInfoSet DevInfo, deviceInfoData *DevInfoData, driverInfoData *DrvInfoData, driverInfoDetailData *DrvInfoDetailData, driverInfoDetailDataSize uint32, requiredSize *uint32) (err error) = setupapi.SetupDiGetDriverInfoDetailW
<原文结束>

# <翻译开始>
//sys	setupDiGetDriverInfoDetail(deviceInfoSet DevInfo, deviceInfoData *DevInfoData, driverInfoData *DrvInfoData, driverInfoDetailData *DrvInfoDetailData, driverInfoDetailDataSize uint32, requiredSize *uint32) (err error) = setupapi.SetupDiGetDriverInfoDetailW
# <翻译结束>


<原文开始>
// SetupDiGetDriverInfoDetail function retrieves driver information detail for a device information set or a particular device information element in the device information set.
<原文结束>

# <翻译开始>
// SetupDiGetDriverInfoDetail function retrieves driver information detail for a device information set or a particular device information element in the device information set.
# <翻译结束>


<原文开始>
// DriverInfoDetail method retrieves driver information detail for a device information set or a particular device information element in the device information set.
<原文结束>

# <翻译开始>
// DriverInfoDetail method retrieves driver information detail for a device information set or a particular device information element in the device information set.
# <翻译结束>


<原文开始>
//sys	SetupDiDestroyDriverInfoList(deviceInfoSet DevInfo, deviceInfoData *DevInfoData, driverType SPDIT) (err error) = setupapi.SetupDiDestroyDriverInfoList
<原文结束>

# <翻译开始>
//sys	SetupDiDestroyDriverInfoList(deviceInfoSet DevInfo, deviceInfoData *DevInfoData, driverType SPDIT) (err error) = setupapi.SetupDiDestroyDriverInfoList
# <翻译结束>


<原文开始>
// DestroyDriverInfoList method deletes a driver list.
<原文结束>

# <翻译开始>
// DestroyDriverInfoList method deletes a driver list.
# <翻译结束>


<原文开始>
//sys	setupDiGetClassDevsEx(classGUID *GUID, Enumerator *uint16, hwndParent uintptr, Flags DIGCF, deviceInfoSet DevInfo, machineName *uint16, reserved uintptr) (handle DevInfo, err error) [failretval==DevInfo(InvalidHandle)] = setupapi.SetupDiGetClassDevsExW
<原文结束>

# <翻译开始>
//sys	setupDiGetClassDevsEx(classGUID *GUID, Enumerator *uint16, hwndParent uintptr, Flags DIGCF, deviceInfoSet DevInfo, machineName *uint16, reserved uintptr) (handle DevInfo, err error) [failretval==DevInfo(InvalidHandle)] = setupapi.SetupDiGetClassDevsExW
# <翻译结束>


<原文开始>
// SetupDiGetClassDevsEx function returns a handle to a device information set that contains requested device information elements for a local or a remote computer.
<原文结束>

# <翻译开始>
// SetupDiGetClassDevsEx function returns a handle to a device information set that contains requested device information elements for a local or a remote computer.
# <翻译结束>


<原文开始>
// SetupDiCallClassInstaller function calls the appropriate class installer, and any registered co-installers, with the specified installation request (DIF code).
//sys	SetupDiCallClassInstaller(installFunction DI_FUNCTION, deviceInfoSet DevInfo, deviceInfoData *DevInfoData) (err error) = setupapi.SetupDiCallClassInstaller
<原文结束>

# <翻译开始>
// SetupDiCallClassInstaller function calls the appropriate class installer, and any registered co-installers, with the specified installation request (DIF code).
//sys	SetupDiCallClassInstaller(installFunction DI_FUNCTION, deviceInfoSet DevInfo, deviceInfoData *DevInfoData) (err error) = setupapi.SetupDiCallClassInstaller
# <翻译结束>


<原文开始>
// CallClassInstaller member calls the appropriate class installer, and any registered co-installers, with the specified installation request (DIF code).
<原文结束>

# <翻译开始>
// CallClassInstaller member calls the appropriate class installer, and any registered co-installers, with the specified installation request (DIF code).
# <翻译结束>


<原文开始>
// SetupDiOpenDevRegKey function opens a registry key for device-specific configuration information.
//sys	SetupDiOpenDevRegKey(deviceInfoSet DevInfo, deviceInfoData *DevInfoData, Scope DICS_FLAG, HwProfile uint32, KeyType DIREG, samDesired uint32) (key Handle, err error) [failretval==InvalidHandle] = setupapi.SetupDiOpenDevRegKey
<原文结束>

# <翻译开始>
// SetupDiOpenDevRegKey function opens a registry key for device-specific configuration information.
//sys	SetupDiOpenDevRegKey(deviceInfoSet DevInfo, deviceInfoData *DevInfoData, Scope DICS_FLAG, HwProfile uint32, KeyType DIREG, samDesired uint32) (key Handle, err error) [failretval==InvalidHandle] = setupapi.SetupDiOpenDevRegKey
# <翻译结束>


<原文开始>
// OpenDevRegKey method opens a registry key for device-specific configuration information.
<原文结束>

# <翻译开始>
// OpenDevRegKey method opens a registry key for device-specific configuration information.
# <翻译结束>


<原文开始>
//sys	setupDiGetDeviceProperty(deviceInfoSet DevInfo, deviceInfoData *DevInfoData, propertyKey *DEVPROPKEY, propertyType *DEVPROPTYPE, propertyBuffer *byte, propertyBufferSize uint32, requiredSize *uint32, flags uint32) (err error) = setupapi.SetupDiGetDevicePropertyW
<原文结束>

# <翻译开始>
//sys	setupDiGetDeviceProperty(deviceInfoSet DevInfo, deviceInfoData *DevInfoData, propertyKey *DEVPROPKEY, propertyType *DEVPROPTYPE, propertyBuffer *byte, propertyBufferSize uint32, requiredSize *uint32, flags uint32) (err error) = setupapi.SetupDiGetDevicePropertyW
# <翻译结束>


<原文开始>
// SetupDiGetDeviceProperty function retrieves a specified device instance property.
<原文结束>

# <翻译开始>
// SetupDiGetDeviceProperty function retrieves a specified device instance property.
# <翻译结束>


<原文开始>
//sys	setupDiGetDeviceRegistryProperty(deviceInfoSet DevInfo, deviceInfoData *DevInfoData, property SPDRP, propertyRegDataType *uint32, propertyBuffer *byte, propertyBufferSize uint32, requiredSize *uint32) (err error) = setupapi.SetupDiGetDeviceRegistryPropertyW
<原文结束>

# <翻译开始>
//sys	setupDiGetDeviceRegistryProperty(deviceInfoSet DevInfo, deviceInfoData *DevInfoData, property SPDRP, propertyRegDataType *uint32, propertyBuffer *byte, propertyBufferSize uint32, requiredSize *uint32) (err error) = setupapi.SetupDiGetDeviceRegistryPropertyW
# <翻译结束>


<原文开始>
// SetupDiGetDeviceRegistryProperty function retrieves a specified Plug and Play device property.
<原文结束>

# <翻译开始>
// SetupDiGetDeviceRegistryProperty function retrieves a specified Plug and Play device property.
# <翻译结束>


<原文开始>
// bufToUTF16 function reinterprets []byte buffer as []uint16
<原文结束>

# <翻译开始>
// bufToUTF16 function reinterprets []byte buffer as []uint16
# <翻译结束>


<原文开始>
// utf16ToBuf function reinterprets []uint16 as []byte
<原文结束>

# <翻译开始>
// utf16ToBuf function reinterprets []uint16 as []byte
# <翻译结束>


<原文开始>
// DeviceRegistryProperty method retrieves a specified Plug and Play device property.
<原文结束>

# <翻译开始>
// DeviceRegistryProperty method retrieves a specified Plug and Play device property.
# <翻译结束>


<原文开始>
//sys	setupDiSetDeviceRegistryProperty(deviceInfoSet DevInfo, deviceInfoData *DevInfoData, property SPDRP, propertyBuffer *byte, propertyBufferSize uint32) (err error) = setupapi.SetupDiSetDeviceRegistryPropertyW
<原文结束>

# <翻译开始>
//sys	setupDiSetDeviceRegistryProperty(deviceInfoSet DevInfo, deviceInfoData *DevInfoData, property SPDRP, propertyBuffer *byte, propertyBufferSize uint32) (err error) = setupapi.SetupDiSetDeviceRegistryPropertyW
# <翻译结束>


<原文开始>
// SetupDiSetDeviceRegistryProperty function sets a Plug and Play device property for a device.
<原文结束>

# <翻译开始>
// SetupDiSetDeviceRegistryProperty function sets a Plug and Play device property for a device.
# <翻译结束>


<原文开始>
// SetDeviceRegistryProperty function sets a Plug and Play device property for a device.
<原文结束>

# <翻译开始>
// SetDeviceRegistryProperty function sets a Plug and Play device property for a device.
# <翻译结束>


<原文开始>
// SetDeviceRegistryPropertyString method sets a Plug and Play device property string for a device.
<原文结束>

# <翻译开始>
// SetDeviceRegistryPropertyString method sets a Plug and Play device property string for a device.
# <翻译结束>


<原文开始>
//sys	setupDiGetDeviceInstallParams(deviceInfoSet DevInfo, deviceInfoData *DevInfoData, deviceInstallParams *DevInstallParams) (err error) = setupapi.SetupDiGetDeviceInstallParamsW
<原文结束>

# <翻译开始>
//sys	setupDiGetDeviceInstallParams(deviceInfoSet DevInfo, deviceInfoData *DevInfoData, deviceInstallParams *DevInstallParams) (err error) = setupapi.SetupDiGetDeviceInstallParamsW
# <翻译结束>


<原文开始>
// SetupDiGetDeviceInstallParams function retrieves device installation parameters for a device information set or a particular device information element.
<原文结束>

# <翻译开始>
// SetupDiGetDeviceInstallParams function retrieves device installation parameters for a device information set or a particular device information element.
# <翻译结束>


<原文开始>
// DeviceInstallParams method retrieves device installation parameters for a device information set or a particular device information element.
<原文结束>

# <翻译开始>
// DeviceInstallParams method retrieves device installation parameters for a device information set or a particular device information element.
# <翻译结束>


<原文开始>
//sys	setupDiGetDeviceInstanceId(deviceInfoSet DevInfo, deviceInfoData *DevInfoData, instanceId *uint16, instanceIdSize uint32, instanceIdRequiredSize *uint32) (err error) = setupapi.SetupDiGetDeviceInstanceIdW
<原文结束>

# <翻译开始>
//sys	setupDiGetDeviceInstanceId(deviceInfoSet DevInfo, deviceInfoData *DevInfoData, instanceId *uint16, instanceIdSize uint32, instanceIdRequiredSize *uint32) (err error) = setupapi.SetupDiGetDeviceInstanceIdW
# <翻译结束>


<原文开始>
// SetupDiGetDeviceInstanceId function retrieves the instance ID of the device.
<原文结束>

# <翻译开始>
// SetupDiGetDeviceInstanceId function retrieves the instance ID of the device.
# <翻译结束>


<原文开始>
// DeviceInstanceID method retrieves the instance ID of the device.
<原文结束>

# <翻译开始>
// DeviceInstanceID method retrieves the instance ID of the device.
# <翻译结束>


<原文开始>
// SetupDiGetClassInstallParams function retrieves class installation parameters for a device information set or a particular device information element.
//sys	SetupDiGetClassInstallParams(deviceInfoSet DevInfo, deviceInfoData *DevInfoData, classInstallParams *ClassInstallHeader, classInstallParamsSize uint32, requiredSize *uint32) (err error) = setupapi.SetupDiGetClassInstallParamsW
<原文结束>

# <翻译开始>
// SetupDiGetClassInstallParams function retrieves class installation parameters for a device information set or a particular device information element.
//sys	SetupDiGetClassInstallParams(deviceInfoSet DevInfo, deviceInfoData *DevInfoData, classInstallParams *ClassInstallHeader, classInstallParamsSize uint32, requiredSize *uint32) (err error) = setupapi.SetupDiGetClassInstallParamsW
# <翻译结束>


<原文开始>
// ClassInstallParams method retrieves class installation parameters for a device information set or a particular device information element.
<原文结束>

# <翻译开始>
// ClassInstallParams method retrieves class installation parameters for a device information set or a particular device information element.
# <翻译结束>


<原文开始>
//sys	SetupDiSetDeviceInstallParams(deviceInfoSet DevInfo, deviceInfoData *DevInfoData, deviceInstallParams *DevInstallParams) (err error) = setupapi.SetupDiSetDeviceInstallParamsW
<原文结束>

# <翻译开始>
//sys	SetupDiSetDeviceInstallParams(deviceInfoSet DevInfo, deviceInfoData *DevInfoData, deviceInstallParams *DevInstallParams) (err error) = setupapi.SetupDiSetDeviceInstallParamsW
# <翻译结束>


<原文开始>
// SetDeviceInstallParams member sets device installation parameters for a device information set or a particular device information element.
<原文结束>

# <翻译开始>
// SetDeviceInstallParams member sets device installation parameters for a device information set or a particular device information element.
# <翻译结束>


<原文开始>
// SetupDiSetClassInstallParams function sets or clears class install parameters for a device information set or a particular device information element.
//sys	SetupDiSetClassInstallParams(deviceInfoSet DevInfo, deviceInfoData *DevInfoData, classInstallParams *ClassInstallHeader, classInstallParamsSize uint32) (err error) = setupapi.SetupDiSetClassInstallParamsW
<原文结束>

# <翻译开始>
// SetupDiSetClassInstallParams function sets or clears class install parameters for a device information set or a particular device information element.
//sys	SetupDiSetClassInstallParams(deviceInfoSet DevInfo, deviceInfoData *DevInfoData, classInstallParams *ClassInstallHeader, classInstallParamsSize uint32) (err error) = setupapi.SetupDiSetClassInstallParamsW
# <翻译结束>


<原文开始>
// SetClassInstallParams method sets or clears class install parameters for a device information set or a particular device information element.
<原文结束>

# <翻译开始>
// SetClassInstallParams method sets or clears class install parameters for a device information set or a particular device information element.
# <翻译结束>


<原文开始>
//sys	setupDiClassNameFromGuidEx(classGUID *GUID, className *uint16, classNameSize uint32, requiredSize *uint32, machineName *uint16, reserved uintptr) (err error) = setupapi.SetupDiClassNameFromGuidExW
<原文结束>

# <翻译开始>
//sys	setupDiClassNameFromGuidEx(classGUID *GUID, className *uint16, classNameSize uint32, requiredSize *uint32, machineName *uint16, reserved uintptr) (err error) = setupapi.SetupDiClassNameFromGuidExW
# <翻译结束>


<原文开始>
// SetupDiClassNameFromGuidEx function retrieves the class name associated with a class GUID. The class can be installed on a local or remote computer.
<原文结束>

# <翻译开始>
// SetupDiClassNameFromGuidEx function retrieves the class name associated with a class GUID. The class can be installed on a local or remote computer.
# <翻译结束>


<原文开始>
//sys	setupDiClassGuidsFromNameEx(className *uint16, classGuidList *GUID, classGuidListSize uint32, requiredSize *uint32, machineName *uint16, reserved uintptr) (err error) = setupapi.SetupDiClassGuidsFromNameExW
<原文结束>

# <翻译开始>
//sys	setupDiClassGuidsFromNameEx(className *uint16, classGuidList *GUID, classGuidListSize uint32, requiredSize *uint32, machineName *uint16, reserved uintptr) (err error) = setupapi.SetupDiClassGuidsFromNameExW
# <翻译结束>


<原文开始>
// SetupDiClassGuidsFromNameEx function retrieves the GUIDs associated with the specified class name. This resulting list contains the classes currently installed on a local or remote computer.
<原文结束>

# <翻译开始>
// SetupDiClassGuidsFromNameEx function retrieves the GUIDs associated with the specified class name. This resulting list contains the classes currently installed on a local or remote computer.
# <翻译结束>


<原文开始>
//sys	setupDiGetSelectedDevice(deviceInfoSet DevInfo, deviceInfoData *DevInfoData) (err error) = setupapi.SetupDiGetSelectedDevice
<原文结束>

# <翻译开始>
//sys	setupDiGetSelectedDevice(deviceInfoSet DevInfo, deviceInfoData *DevInfoData) (err error) = setupapi.SetupDiGetSelectedDevice
# <翻译结束>


<原文开始>
// SetupDiGetSelectedDevice function retrieves the selected device information element in a device information set.
<原文结束>

# <翻译开始>
// SetupDiGetSelectedDevice function retrieves the selected device information element in a device information set.
# <翻译结束>


<原文开始>
// SelectedDevice method retrieves the selected device information element in a device information set.
<原文结束>

# <翻译开始>
// SelectedDevice method retrieves the selected device information element in a device information set.
# <翻译结束>


<原文开始>
// SetupDiSetSelectedDevice function sets a device information element as the selected member of a device information set. This function is typically used by an installation wizard.
//sys	SetupDiSetSelectedDevice(deviceInfoSet DevInfo, deviceInfoData *DevInfoData) (err error) = setupapi.SetupDiSetSelectedDevice
<原文结束>

# <翻译开始>
// SetupDiSetSelectedDevice function sets a device information element as the selected member of a device information set. This function is typically used by an installation wizard.
//sys	SetupDiSetSelectedDevice(deviceInfoSet DevInfo, deviceInfoData *DevInfoData) (err error) = setupapi.SetupDiSetSelectedDevice
# <翻译结束>


<原文开始>
// SetSelectedDevice method sets a device information element as the selected member of a device information set. This function is typically used by an installation wizard.
<原文结束>

# <翻译开始>
// SetSelectedDevice method sets a device information element as the selected member of a device information set. This function is typically used by an installation wizard.
# <翻译结束>


<原文开始>
//sys	setupUninstallOEMInf(infFileName *uint16, flags SUOI, reserved uintptr) (err error) = setupapi.SetupUninstallOEMInfW
<原文结束>

# <翻译开始>
//sys	setupUninstallOEMInf(infFileName *uint16, flags SUOI, reserved uintptr) (err error) = setupapi.SetupUninstallOEMInfW
# <翻译结束>


<原文开始>
// SetupUninstallOEMInf uninstalls the specified driver.
<原文结束>

# <翻译开始>
// SetupUninstallOEMInf uninstalls the specified driver.
# <翻译结束>


<原文开始>
//sys cm_MapCrToWin32Err(configRet CONFIGRET, defaultWin32Error Errno) (ret Errno) = CfgMgr32.CM_MapCrToWin32Err
<原文结束>

# <翻译开始>
//sys cm_MapCrToWin32Err(configRet CONFIGRET, defaultWin32Error Errno) (ret Errno) = CfgMgr32.CM_MapCrToWin32Err
# <翻译结束>


<原文开始>
//sys cm_Get_Device_Interface_List_Size(len *uint32, interfaceClass *GUID, deviceID *uint16, flags uint32) (ret CONFIGRET) = CfgMgr32.CM_Get_Device_Interface_List_SizeW
//sys cm_Get_Device_Interface_List(interfaceClass *GUID, deviceID *uint16, buffer *uint16, bufferLen uint32, flags uint32) (ret CONFIGRET) = CfgMgr32.CM_Get_Device_Interface_ListW
<原文结束>

# <翻译开始>
//sys cm_Get_Device_Interface_List_Size(len *uint32, interfaceClass *GUID, deviceID *uint16, flags uint32) (ret CONFIGRET) = CfgMgr32.CM_Get_Device_Interface_List_SizeW
//sys cm_Get_Device_Interface_List(interfaceClass *GUID, deviceID *uint16, buffer *uint16, bufferLen uint32, flags uint32) (ret CONFIGRET) = CfgMgr32.CM_Get_Device_Interface_ListW
# <翻译结束>


<原文开始>
//sys cm_Get_DevNode_Status(status *uint32, problemNumber *uint32, devInst DEVINST, flags uint32) (ret CONFIGRET) = CfgMgr32.CM_Get_DevNode_Status
<原文结束>

# <翻译开始>
//sys cm_Get_DevNode_Status(status *uint32, problemNumber *uint32, devInst DEVINST, flags uint32) (ret CONFIGRET) = CfgMgr32.CM_Get_DevNode_Status
# <翻译结束>

