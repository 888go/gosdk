
<原文开始>
// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// 版权所有 2018 The Go Authors。保留所有权利。
// 使用本源代码须遵循 BSD 风格许可证，
// 该许可证可在 LICENSE 文件中找到。
# <翻译结束>


<原文开始>
// Possible recovery actions that the service control manager can perform.
<原文结束>

# <翻译开始>
// 服务控制管理器可能执行的恢复操作。
# <翻译结束>


<原文开始>
// RecoveryAction represents an action that the service control manager can perform when service fails.
// A service is considered failed when it terminates without reporting a status of SERVICE_STOPPED to the service controller.
<原文结束>

# <翻译开始>
// RecoveryAction 表示服务控制管理器在服务失败时可执行的操作。
// 当服务在未向服务控制器报告 SERVICE_STOPPED 状态的情况下终止时，视为该服务已失败。
# <翻译结束>


<原文开始>
// one of NoAction, ComputerReboot, ServiceRestart or RunCommand
<原文结束>

# <翻译开始>
// 可为 NoAction、ComputerReboot、ServiceRestart 或 RunCommand 中的一种
# <翻译结束>


<原文开始>
// the time to wait before performing the specified action
<原文结束>

# <翻译开始>
// 执行指定操作前的等待时间
# <翻译结束>


<原文开始>
// SetRecoveryActions sets actions that service controller performs when service fails and
// the time after which to reset the service failure count to zero if there are no failures, in seconds.
// Specify INFINITE to indicate that service failure count should never be reset.
<原文结束>

# <翻译开始>
// SetRecoveryActions 设置当服务失败时，服务控制器执行的操作以及在无故障情况下将服务失败计数器重置为零的时间间隔（以秒为单位）。
// 指定 INFINITE 表示服务失败计数器应永不重置。
# <翻译结束>


<原文开始>
// RecoveryActions returns actions that service controller performs when service fails.
// The service control manager counts the number of times service s has failed since the system booted.
// The count is reset to 0 if the service has not failed for ResetPeriod seconds.
// When the service fails for the Nth time, the service controller performs the action specified in element [N-1] of returned slice.
// If N is greater than slice length, the service controller repeats the last action in the slice.
<原文结束>

# <翻译开始>
// RecoveryActions 返回服务控制器在服务失败时执行的操作。
// 服务控制管理器统计自系统启动以来服务 s 失败的次数。
// 如果该服务在 ResetPeriod 秒内未发生故障，计数将重置为 0。
// 当服务第 N 次失败时，服务控制器执行返回切片中元素 [N-1] 指定的动作。
// 若 N 大于切片长度，服务控制器将重复执行切片中的最后一个动作。
# <翻译结束>


<原文开始>
// ResetRecoveryActions deletes both reset period and array of failure actions.
<原文结束>

# <翻译开始>
// 重置恢复动作：删除复位周期与失败动作数组
# <翻译结束>


<原文开始>
// ResetPeriod is the time after which to reset the service failure
// count to zero if there are no failures, in seconds.
<原文结束>

# <翻译开始>
// ResetPeriod 是在无故障情况下将服务失败计数重置为零的时间间隔，以秒为单位。
# <翻译结束>


<原文开始>
// SetRebootMessage sets service s reboot message.
// If msg is "", the reboot message is deleted and no message is broadcast.
<原文结束>

# <翻译开始>
// SetRebootMessage 设置服务s的重启消息。
// 若msg为"", 则删除重启消息，且不广播任何消息。
# <翻译结束>


<原文开始>
// RebootMessage is broadcast to server users before rebooting in response to the ComputerReboot service controller action.
<原文结束>

# <翻译开始>
// RebootMessage 是在响应 ComputerReboot 服务控制器操作重启前，向服务器用户广播的消息。
# <翻译结束>


<原文开始>
// SetRecoveryCommand sets the command line of the process to execute in response to the RunCommand service controller action.
// If cmd is "", the command is deleted and no program is run when the service fails.
<原文结束>

# <翻译开始>
// SetRecoveryCommand 用于设置在响应 RunCommand 服务控制器操作时执行的进程命令行。
// 若 cmd 为空字符串（""），则删除该命令，当服务失败时不再运行任何程序。
# <翻译结束>


<原文开始>
// RecoveryCommand is the command line of the process to execute in response to the RunCommand service controller action. This process runs under the same account as the service.
<原文结束>

# <翻译开始>
// RecoveryCommand 是对服务控制器“RunCommand”操作做出响应时要执行的进程命令行。此进程以与服务相同的身份运行。
# <翻译结束>


<原文开始>
// SetRecoveryActionsOnNonCrashFailures sets the failure actions flag. If the
// flag is set to false, recovery actions will only be performed if the service
// terminates without reporting a status of SERVICE_STOPPED. If the flag is set
// to true, recovery actions are also perfomed if the service stops with a
// nonzero exit code.
<原文结束>

# <翻译开始>
// SetRecoveryActionsOnNonCrashFailures 设置失败动作标志。如果该标志设置为false，
// 则仅当服务终止时未报告SERVICE_STOPPED状态，才会执行恢复动作。
// 如果该标志设置为true，则当服务以非零退出代码停止时，也会执行恢复动作。
# <翻译结束>


<原文开始>
// RecoveryActionsOnNonCrashFailures returns the current value of the failure
// actions flag. If the flag is set to false, recovery actions will only be
// performed if the service terminates without reporting a status of
// SERVICE_STOPPED. If the flag is set to true, recovery actions are also
// perfomed if the service stops with a nonzero exit code.
<原文结束>

# <翻译开始>
// RecoveryActionsOnNonCrashFailures 返回当前的失败操作标志值。如果该标志设为false，只有当服务终止且未报告SERVICE_STOPPED状态时，才会执行恢复动作。如果该标志设为true，只要服务停止时退出代码不为0，也会执行恢复动作。
# <翻译结束>

