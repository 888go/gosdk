
<原文开始>
// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
# <翻译结束>


<原文开始>
// Possible recovery actions that the service control manager can perform.
<原文结束>

# <翻译开始>
// Possible recovery actions that the service control manager can perform.
# <翻译结束>


<原文开始>
// RecoveryAction represents an action that the service control manager can perform when service fails.
// A service is considered failed when it terminates without reporting a status of SERVICE_STOPPED to the service controller.
<原文结束>

# <翻译开始>
// RecoveryAction represents an action that the service control manager can perform when service fails.
// A service is considered failed when it terminates without reporting a status of SERVICE_STOPPED to the service controller.
# <翻译结束>


<原文开始>
// one of NoAction, ComputerReboot, ServiceRestart or RunCommand
<原文结束>

# <翻译开始>
// one of NoAction, ComputerReboot, ServiceRestart or RunCommand
# <翻译结束>


<原文开始>
// the time to wait before performing the specified action
<原文结束>

# <翻译开始>
// the time to wait before performing the specified action
# <翻译结束>


<原文开始>
// SetRecoveryActions sets actions that service controller performs when service fails and
// the time after which to reset the service failure count to zero if there are no failures, in seconds.
// Specify INFINITE to indicate that service failure count should never be reset.
<原文结束>

# <翻译开始>
// SetRecoveryActions sets actions that service controller performs when service fails and
// the time after which to reset the service failure count to zero if there are no failures, in seconds.
// Specify INFINITE to indicate that service failure count should never be reset.
# <翻译结束>


<原文开始>
// RecoveryActions returns actions that service controller performs when service fails.
// The service control manager counts the number of times service s has failed since the system booted.
// The count is reset to 0 if the service has not failed for ResetPeriod seconds.
// When the service fails for the Nth time, the service controller performs the action specified in element [N-1] of returned slice.
// If N is greater than slice length, the service controller repeats the last action in the slice.
<原文结束>

# <翻译开始>
// RecoveryActions returns actions that service controller performs when service fails.
// The service control manager counts the number of times service s has failed since the system booted.
// The count is reset to 0 if the service has not failed for ResetPeriod seconds.
// When the service fails for the Nth time, the service controller performs the action specified in element [N-1] of returned slice.
// If N is greater than slice length, the service controller repeats the last action in the slice.
# <翻译结束>


<原文开始>
// ResetRecoveryActions deletes both reset period and array of failure actions.
<原文结束>

# <翻译开始>
// ResetRecoveryActions deletes both reset period and array of failure actions.
# <翻译结束>


<原文开始>
// ResetPeriod is the time after which to reset the service failure
// count to zero if there are no failures, in seconds.
<原文结束>

# <翻译开始>
// ResetPeriod is the time after which to reset the service failure
// count to zero if there are no failures, in seconds.
# <翻译结束>


<原文开始>
// SetRebootMessage sets service s reboot message.
// If msg is "", the reboot message is deleted and no message is broadcast.
<原文结束>

# <翻译开始>
// SetRebootMessage sets service s reboot message.
// If msg is "", the reboot message is deleted and no message is broadcast.
# <翻译结束>


<原文开始>
// RebootMessage is broadcast to server users before rebooting in response to the ComputerReboot service controller action.
<原文结束>

# <翻译开始>
// RebootMessage is broadcast to server users before rebooting in response to the ComputerReboot service controller action.
# <翻译结束>


<原文开始>
// SetRecoveryCommand sets the command line of the process to execute in response to the RunCommand service controller action.
// If cmd is "", the command is deleted and no program is run when the service fails.
<原文结束>

# <翻译开始>
// SetRecoveryCommand sets the command line of the process to execute in response to the RunCommand service controller action.
// If cmd is "", the command is deleted and no program is run when the service fails.
# <翻译结束>


<原文开始>
// RecoveryCommand is the command line of the process to execute in response to the RunCommand service controller action. This process runs under the same account as the service.
<原文结束>

# <翻译开始>
// RecoveryCommand is the command line of the process to execute in response to the RunCommand service controller action. This process runs under the same account as the service.
# <翻译结束>


<原文开始>
// SetRecoveryActionsOnNonCrashFailures sets the failure actions flag. If the
// flag is set to false, recovery actions will only be performed if the service
// terminates without reporting a status of SERVICE_STOPPED. If the flag is set
// to true, recovery actions are also perfomed if the service stops with a
// nonzero exit code.
<原文结束>

# <翻译开始>
// SetRecoveryActionsOnNonCrashFailures sets the failure actions flag. If the
// flag is set to false, recovery actions will only be performed if the service
// terminates without reporting a status of SERVICE_STOPPED. If the flag is set
// to true, recovery actions are also perfomed if the service stops with a
// nonzero exit code.
# <翻译结束>


<原文开始>
// RecoveryActionsOnNonCrashFailures returns the current value of the failure
// actions flag. If the flag is set to false, recovery actions will only be
// performed if the service terminates without reporting a status of
// SERVICE_STOPPED. If the flag is set to true, recovery actions are also
// perfomed if the service stops with a nonzero exit code.
<原文结束>

# <翻译开始>
// RecoveryActionsOnNonCrashFailures returns the current value of the failure
// actions flag. If the flag is set to false, recovery actions will only be
// performed if the service terminates without reporting a status of
// SERVICE_STOPPED. If the flag is set to true, recovery actions are also
// perfomed if the service stops with a nonzero exit code.
# <翻译结束>

