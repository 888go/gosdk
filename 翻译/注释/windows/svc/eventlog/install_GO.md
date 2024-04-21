
<原文开始>
// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// 版权所有 2012 The Go 作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。
# <翻译结束>


<原文开始>
// Install modifies PC registry to allow logging with an event source src.
// It adds all required keys and values to the event log registry key.
// Install uses msgFile as the event message file. If useExpandKey is true,
// the event message file is installed as REG_EXPAND_SZ value,
// otherwise as REG_SZ. Use bitwise of log.Error, log.Warning and
// log.Info to specify events supported by the new event source.
<原文结束>

# <翻译开始>
// Install 通过修改 PC 注册表，允许使用事件源 src 进行日志记录。
// 它向事件日志注册键添加所有必需的键和值。
// Install 使用 msgFile 作为事件消息文件。如果 useExpandKey 为真，
// 则事件消息文件以 REG_EXPAND_SZ 值的形式安装，
// 否则以 REG_SZ 形式安装。使用 log.Error、log.Warning 和
// log.Info 的按位组合指定新事件源支持的事件。
# <翻译结束>


<原文开始>
// InstallAsEventCreate is the same as Install, but uses
// %SystemRoot%\System32\EventCreate.exe as the event message file.
<原文结束>

# <翻译开始>
// InstallAsEventCreate 与 Install 相同，但使用
// %SystemRoot%\System32\EventCreate.exe 作为事件消息文件。
# <翻译结束>


<原文开始>
// Remove deletes all registry elements installed by the correspondent Install.
<原文结束>

# <翻译开始>
// Remove 删除由相应 Install 安装的所有注册表元素。
# <翻译结束>

