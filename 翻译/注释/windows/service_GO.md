
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
//sys	OpenSCManager(machineName *uint16, databaseName *uint16, access uint32) (handle Handle, err error) [failretval==0] = advapi32.OpenSCManagerW
//sys	CloseServiceHandle(handle Handle) (err error) = advapi32.CloseServiceHandle
//sys	CreateService(mgr Handle, serviceName *uint16, displayName *uint16, access uint32, srvType uint32, startType uint32, errCtl uint32, pathName *uint16, loadOrderGroup *uint16, tagId *uint32, dependencies *uint16, serviceStartName *uint16, password *uint16) (handle Handle, err error) [failretval==0] = advapi32.CreateServiceW
//sys	OpenService(mgr Handle, serviceName *uint16, access uint32) (handle Handle, err error) [failretval==0] = advapi32.OpenServiceW
//sys	DeleteService(service Handle) (err error) = advapi32.DeleteService
//sys	StartService(service Handle, numArgs uint32, argVectors **uint16) (err error) = advapi32.StartServiceW
//sys	QueryServiceStatus(service Handle, status *SERVICE_STATUS) (err error) = advapi32.QueryServiceStatus
//sys	QueryServiceLockStatus(mgr Handle, lockStatus *QUERY_SERVICE_LOCK_STATUS, bufSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceLockStatusW
//sys	ControlService(service Handle, control uint32, status *SERVICE_STATUS) (err error) = advapi32.ControlService
//sys	StartServiceCtrlDispatcher(serviceTable *SERVICE_TABLE_ENTRY) (err error) = advapi32.StartServiceCtrlDispatcherW
//sys	SetServiceStatus(service Handle, serviceStatus *SERVICE_STATUS) (err error) = advapi32.SetServiceStatus
//sys	ChangeServiceConfig(service Handle, serviceType uint32, startType uint32, errorControl uint32, binaryPathName *uint16, loadOrderGroup *uint16, tagId *uint32, dependencies *uint16, serviceStartName *uint16, password *uint16, displayName *uint16) (err error) = advapi32.ChangeServiceConfigW
//sys	QueryServiceConfig(service Handle, serviceConfig *QUERY_SERVICE_CONFIG, bufSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceConfigW
//sys	ChangeServiceConfig2(service Handle, infoLevel uint32, info *byte) (err error) = advapi32.ChangeServiceConfig2W
//sys	QueryServiceConfig2(service Handle, infoLevel uint32, buff *byte, buffSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceConfig2W
//sys	EnumServicesStatusEx(mgr Handle, infoLevel uint32, serviceType uint32, serviceState uint32, services *byte, bufSize uint32, bytesNeeded *uint32, servicesReturned *uint32, resumeHandle *uint32, groupName *uint16) (err error) = advapi32.EnumServicesStatusExW
//sys	QueryServiceStatusEx(service Handle, infoLevel uint32, buff *byte, buffSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceStatusEx
//sys	NotifyServiceStatusChange(service Handle, notifyMask uint32, notifier *SERVICE_NOTIFY) (ret error) = advapi32.NotifyServiceStatusChangeW
//sys	SubscribeServiceChangeNotifications(service Handle, eventType uint32, callback uintptr, callbackCtx uintptr, subscription *uintptr) (ret error) = sechost.SubscribeServiceChangeNotifications?
//sys	UnsubscribeServiceChangeNotifications(subscription uintptr) = sechost.UnsubscribeServiceChangeNotifications?
//sys	RegisterServiceCtrlHandlerEx(serviceName *uint16, handlerProc uintptr, context uintptr) (handle Handle, err error) = advapi32.RegisterServiceCtrlHandlerExW
//sys	QueryServiceDynamicInformation(service Handle, infoLevel uint32, dynamicInfo unsafe.Pointer) (err error) = advapi32.QueryServiceDynamicInformation?
//sys	EnumDependentServices(service Handle, activityState uint32, services *ENUM_SERVICE_STATUS, buffSize uint32, bytesNeeded *uint32, servicesReturned *uint32) (err error) = advapi32.EnumDependentServicesW
<原文结束>

# <翻译开始>
//sys	OpenSCManager(machineName *uint16, databaseName *uint16, access uint32) (handle Handle, err error) [failretval==0] = advapi32.OpenSCManagerW
//sys	CloseServiceHandle(handle Handle) (err error) = advapi32.CloseServiceHandle
//sys	CreateService(mgr Handle, serviceName *uint16, displayName *uint16, access uint32, srvType uint32, startType uint32, errCtl uint32, pathName *uint16, loadOrderGroup *uint16, tagId *uint32, dependencies *uint16, serviceStartName *uint16, password *uint16) (handle Handle, err error) [failretval==0] = advapi32.CreateServiceW
//sys	OpenService(mgr Handle, serviceName *uint16, access uint32) (handle Handle, err error) [failretval==0] = advapi32.OpenServiceW
//sys	DeleteService(service Handle) (err error) = advapi32.DeleteService
//sys	StartService(service Handle, numArgs uint32, argVectors **uint16) (err error) = advapi32.StartServiceW
//sys	QueryServiceStatus(service Handle, status *SERVICE_STATUS) (err error) = advapi32.QueryServiceStatus
//sys	QueryServiceLockStatus(mgr Handle, lockStatus *QUERY_SERVICE_LOCK_STATUS, bufSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceLockStatusW
//sys	ControlService(service Handle, control uint32, status *SERVICE_STATUS) (err error) = advapi32.ControlService
//sys	StartServiceCtrlDispatcher(serviceTable *SERVICE_TABLE_ENTRY) (err error) = advapi32.StartServiceCtrlDispatcherW
//sys	SetServiceStatus(service Handle, serviceStatus *SERVICE_STATUS) (err error) = advapi32.SetServiceStatus
//sys	ChangeServiceConfig(service Handle, serviceType uint32, startType uint32, errorControl uint32, binaryPathName *uint16, loadOrderGroup *uint16, tagId *uint32, dependencies *uint16, serviceStartName *uint16, password *uint16, displayName *uint16) (err error) = advapi32.ChangeServiceConfigW
//sys	QueryServiceConfig(service Handle, serviceConfig *QUERY_SERVICE_CONFIG, bufSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceConfigW
//sys	ChangeServiceConfig2(service Handle, infoLevel uint32, info *byte) (err error) = advapi32.ChangeServiceConfig2W
//sys	QueryServiceConfig2(service Handle, infoLevel uint32, buff *byte, buffSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceConfig2W
//sys	EnumServicesStatusEx(mgr Handle, infoLevel uint32, serviceType uint32, serviceState uint32, services *byte, bufSize uint32, bytesNeeded *uint32, servicesReturned *uint32, resumeHandle *uint32, groupName *uint16) (err error) = advapi32.EnumServicesStatusExW
//sys	QueryServiceStatusEx(service Handle, infoLevel uint32, buff *byte, buffSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceStatusEx
//sys	NotifyServiceStatusChange(service Handle, notifyMask uint32, notifier *SERVICE_NOTIFY) (ret error) = advapi32.NotifyServiceStatusChangeW
//sys	SubscribeServiceChangeNotifications(service Handle, eventType uint32, callback uintptr, callbackCtx uintptr, subscription *uintptr) (ret error) = sechost.SubscribeServiceChangeNotifications?
//sys	UnsubscribeServiceChangeNotifications(subscription uintptr) = sechost.UnsubscribeServiceChangeNotifications?
//sys	RegisterServiceCtrlHandlerEx(serviceName *uint16, handlerProc uintptr, context uintptr) (handle Handle, err error) = advapi32.RegisterServiceCtrlHandlerExW
//sys	QueryServiceDynamicInformation(service Handle, infoLevel uint32, dynamicInfo unsafe.Pointer) (err error) = advapi32.QueryServiceDynamicInformation?
//sys	EnumDependentServices(service Handle, activityState uint32, services *ENUM_SERVICE_STATUS, buffSize uint32, bytesNeeded *uint32, servicesReturned *uint32) (err error) = advapi32.EnumDependentServicesW
# <翻译结束>

