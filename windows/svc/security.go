// 版权所有 ? 2012 The Go Authors。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

//go:build windows

package svc

import (
	"strings"
	"unsafe"

	"github.com/888go/gosdk/windows"
)

func allocSid(subAuth0 uint32) (*windows.SID, error) {
	var sid *windows.SID
	err := windows.AllocateAndInitializeSid(&windows.SECURITY_NT_AUTHORITY,
		1, subAuth0, 0, 0, 0, 0, 0, 0, 0, &sid)
	if err != nil {
		return nil, err
	}
	return sid, nil
}

// IsAnInteractiveSession 用于判断调用进程是否以交互模式运行。
// 它通过查询进程令牌中是否存在“Interactive”组成员身份来进行判断。
// 参考：http://stackoverflow.com/questions/2668851/how-do-i-detect-that-my-application-is-running-as-service-or-in-an-interactive-s
// 
// 已弃用：请改用 IsWindowsService。

// ff:
func IsAnInteractiveSession() (bool, error) {
	interSid, err := allocSid(windows.SECURITY_INTERACTIVE_RID)
	if err != nil {
		return false, err
	}
	defer windows.FreeSid(interSid)

	serviceSid, err := allocSid(windows.SECURITY_SERVICE_RID)
	if err != nil {
		return false, err
	}
	defer windows.FreeSid(serviceSid)

	t, err := windows.OpenCurrentProcessToken()
	if err != nil {
		return false, err
	}
	defer t.Close()

	gs, err := t.GetTokenGroups()
	if err != nil {
		return false, err
	}

	for _, g := range gs.AllGroups() {
		if windows.EqualSid(g.Sid, interSid) {
			return true, nil
		}
		if windows.EqualSid(g.Sid, serviceSid) {
			return false, nil
		}
	}
	return false, nil
}

// IsWindowsService 判断当前进程是否作为 Windows 服务执行

// ff:
func IsWindowsService() (bool, error) {
// 下面的技巧看起来有些复杂，但实际上正是.NET框架为同名函数所采用的方法：
// https://github.com/dotnet/extensions/blob/f4066026ca06984b07e90e61a6390ac38152ba93/src/Hosting/WindowsServices/src/WindowsServiceHelpers.cs#L26-L31
// 具体来说，它会检查父进程是否具有会话ID零，并且名为“services”。

	var currentProcess windows.PROCESS_BASIC_INFORMATION
	infoSize := uint32(unsafe.Sizeof(currentProcess))
	err := windows.NtQueryInformationProcess(windows.CurrentProcess(), windows.ProcessBasicInformation, unsafe.Pointer(&currentProcess), infoSize, &infoSize)
	if err != nil {
		return false, err
	}
	var parentProcess *windows.SYSTEM_PROCESS_INFORMATION
	for infoSize = uint32((unsafe.Sizeof(*parentProcess) + unsafe.Sizeof(uintptr(0))) * 1024); ; {
		parentProcess = (*windows.SYSTEM_PROCESS_INFORMATION)(unsafe.Pointer(&make([]byte, infoSize)[0]))
		err = windows.NtQuerySystemInformation(windows.SystemProcessInformation, unsafe.Pointer(parentProcess), infoSize, &infoSize)
		if err == nil {
			break
		} else if err != windows.STATUS_INFO_LENGTH_MISMATCH {
			return false, err
		}
	}
	for ; ; parentProcess = (*windows.SYSTEM_PROCESS_INFORMATION)(unsafe.Pointer(uintptr(unsafe.Pointer(parentProcess)) + uintptr(parentProcess.NextEntryOffset))) {
		if parentProcess.UniqueProcessID == currentProcess.InheritedFromUniqueProcessId {
			return parentProcess.SessionID == 0 && strings.EqualFold("services.exe", parentProcess.ImageName.String()), nil
		}
		if parentProcess.NextEntryOffset == 0 {
			break
		}
	}
	return false, nil
}
