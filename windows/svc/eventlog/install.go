// 版权所有 2012 The Go 作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

//go:build windows

package eventlog

import (
	"errors"

	"github.com/888go/gosdk/windows"
	"github.com/888go/gosdk/windows/registry"
)

const (
	// Log levels.
	Info    = windows.EVENTLOG_INFORMATION_TYPE
	Warning = windows.EVENTLOG_WARNING_TYPE
	Error   = windows.EVENTLOG_ERROR_TYPE
)

const addKeyName = `SYSTEM\CurrentControlSet\Services\EventLog\Application`

// Install 通过修改 PC 注册表，允许使用事件源 src 进行日志记录。
// 它向事件日志注册键添加所有必需的键和值。
// Install 使用 msgFile 作为事件消息文件。如果 useExpandKey 为真，
// 则事件消息文件以 REG_EXPAND_SZ 值的形式安装，
// 否则以 REG_SZ 形式安装。使用 log.Error、log.Warning 和
// log.Info 的按位组合指定新事件源支持的事件。
func Install(src, msgFile string, useExpandKey bool, eventsSupported uint32) error {
	appkey, err := registry.OpenKey(registry.LOCAL_MACHINE, addKeyName, registry.CREATE_SUB_KEY)
	if err != nil {
		return err
	}
	defer appkey.Close()

	sk, alreadyExist, err := registry.CreateKey(appkey, src, registry.SET_VALUE)
	if err != nil {
		return err
	}
	defer sk.Close()
	if alreadyExist {
		return errors.New(addKeyName + `\` + src + " registry key already exists")
	}

	err = sk.SetDWordValue("CustomSource", 1)
	if err != nil {
		return err
	}
	if useExpandKey {
		err = sk.SetExpandStringValue("EventMessageFile", msgFile)
	} else {
		err = sk.SetStringValue("EventMessageFile", msgFile)
	}
	if err != nil {
		return err
	}
	err = sk.SetDWordValue("TypesSupported", eventsSupported)
	if err != nil {
		return err
	}
	return nil
}

// InstallAsEventCreate 与 Install 相同，但使用
// %SystemRoot%\System32\EventCreate.exe 作为事件消息文件。
func InstallAsEventCreate(src string, eventsSupported uint32) error {
	return Install(src, "%SystemRoot%\\System32\\EventCreate.exe", true, eventsSupported)
}

// Remove 删除由相应 Install 安装的所有注册表元素。
func Remove(src string) error {
	appkey, err := registry.OpenKey(registry.LOCAL_MACHINE, addKeyName, registry.SET_VALUE)
	if err != nil {
		return err
	}
	defer appkey.Close()
	return registry.DeleteKey(appkey, src)
}
