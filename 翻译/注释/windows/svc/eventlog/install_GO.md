
<原文开始>
// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
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
// Install modifies PC registry to allow logging with an event source src.
// It adds all required keys and values to the event log registry key.
// Install uses msgFile as the event message file. If useExpandKey is true,
// the event message file is installed as REG_EXPAND_SZ value,
// otherwise as REG_SZ. Use bitwise of log.Error, log.Warning and
// log.Info to specify events supported by the new event source.
# <翻译结束>


<原文开始>
// InstallAsEventCreate is the same as Install, but uses
// %SystemRoot%\System32\EventCreate.exe as the event message file.
<原文结束>

# <翻译开始>
// InstallAsEventCreate is the same as Install, but uses
// %SystemRoot%\System32\EventCreate.exe as the event message file.
# <翻译结束>


<原文开始>
// Remove deletes all registry elements installed by the correspondent Install.
<原文结束>

# <翻译开始>
// Remove deletes all registry elements installed by the correspondent Install.
# <翻译结束>

