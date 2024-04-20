
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
// Package svc provides everything required to build Windows service.
<原文结束>

# <翻译开始>
// Package svc provides everything required to build Windows service.
# <翻译结束>


<原文开始>
// State describes service execution state (Stopped, Running and so on).
<原文结束>

# <翻译开始>
// State describes service execution state (Stopped, Running and so on).
# <翻译结束>


<原文开始>
// Cmd represents service state change request. It is sent to a service
// by the service manager, and should be actioned upon by the service.
<原文结束>

# <翻译开始>
// Cmd represents service state change request. It is sent to a service
// by the service manager, and should be actioned upon by the service.
# <翻译结束>


<原文开始>
// Accepted is used to describe commands accepted by the service.
// Note that Interrogate is always accepted.
<原文结束>

# <翻译开始>
// Accepted is used to describe commands accepted by the service.
// Note that Interrogate is always accepted.
# <翻译结束>


<原文开始>
// ActivityStatus allows for services to be selected based on active and inactive categories of service state.
<原文结束>

# <翻译开始>
// ActivityStatus allows for services to be selected based on active and inactive categories of service state.
# <翻译结束>


<原文开始>
// Status combines State and Accepted commands to fully describe running service.
<原文结束>

# <翻译开始>
// Status combines State and Accepted commands to fully describe running service.
# <翻译结束>


<原文开始>
// used to report progress during a lengthy operation
<原文结束>

# <翻译开始>
// used to report progress during a lengthy operation
# <翻译结束>


<原文开始>
// estimated time required for a pending operation, in milliseconds
<原文结束>

# <翻译开始>
// estimated time required for a pending operation, in milliseconds
# <翻译结束>


<原文开始>
// if the service is running, the process identifier of it, and otherwise zero
<原文结束>

# <翻译开始>
// if the service is running, the process identifier of it, and otherwise zero
# <翻译结束>


<原文开始>
// set if the service has exited with a win32 exit code
<原文结束>

# <翻译开始>
// set if the service has exited with a win32 exit code
# <翻译结束>


<原文开始>
// set if the service has exited with a service-specific exit code
<原文结束>

# <翻译开始>
// set if the service has exited with a service-specific exit code
# <翻译结束>


<原文开始>
// StartReason is the reason that the service was started.
<原文结束>

# <翻译开始>
// StartReason is the reason that the service was started.
# <翻译结束>


<原文开始>
// ChangeRequest is sent to the service Handler to request service status change.
<原文结束>

# <翻译开始>
// ChangeRequest is sent to the service Handler to request service status change.
# <翻译结束>


<原文开始>
// Handler is the interface that must be implemented to build Windows service.
<原文结束>

# <翻译开始>
// Handler is the interface that must be implemented to build Windows service.
# <翻译结束>


<原文开始>
	// Execute will be called by the package code at the start of
	// the service, and the service will exit once Execute completes.
	// Inside Execute you must read service change requests from r and
	// act accordingly. You must keep service control manager up to date
	// about state of your service by writing into s as required.
	// args contains service name followed by argument strings passed
	// to the service.
	// You can provide service exit code in exitCode return parameter,
	// with 0 being "no error". You can also indicate if exit code,
	// if any, is service specific or not by using svcSpecificEC
	// parameter.
<原文结束>

# <翻译开始>
	// Execute will be called by the package code at the start of
	// the service, and the service will exit once Execute completes.
	// Inside Execute you must read service change requests from r and
	// act accordingly. You must keep service control manager up to date
	// about state of your service by writing into s as required.
	// args contains service name followed by argument strings passed
	// to the service.
	// You can provide service exit code in exitCode return parameter,
	// with 0 being "no error". You can also indicate if exit code,
	// if any, is service specific or not by using svcSpecificEC
	// parameter.
# <翻译结束>


<原文开始>
// service provides access to windows service api.
<原文结束>

# <翻译开始>
// service provides access to windows service api.
# <翻译结束>


<原文开始>
// Set context to 123456 to test issue #25660.
<原文结束>

# <翻译开始>
// Set context to 123456 to test issue #25660.
# <翻译结束>


<原文开始>
// This is, unfortunately, a global, which means only one service per process.
<原文结束>

# <翻译开始>
// This is, unfortunately, a global, which means only one service per process.
# <翻译结束>


<原文开始>
// serviceMain is the entry point called by the service manager, registered earlier by
// the call to StartServiceCtrlDispatcher.
<原文结束>

# <翻译开始>
// serviceMain is the entry point called by the service manager, registered earlier by
// the call to StartServiceCtrlDispatcher.
# <翻译结束>


<原文开始>
// Run executes service name by calling appropriate handler function.
<原文结束>

# <翻译开始>
// Run executes service name by calling appropriate handler function.
# <翻译结束>


<原文开始>
// StatusHandle returns service status handle. It is safe to call this function
// from inside the Handler.Execute because then it is guaranteed to be set.
<原文结束>

# <翻译开始>
// StatusHandle returns service status handle. It is safe to call this function
// from inside the Handler.Execute because then it is guaranteed to be set.
# <翻译结束>


<原文开始>
// DynamicStartReason returns the reason why the service was started. It is safe
// to call this function from inside the Handler.Execute because then it is
// guaranteed to be set.
<原文结束>

# <翻译开始>
// DynamicStartReason returns the reason why the service was started. It is safe
// to call this function from inside the Handler.Execute because then it is
// guaranteed to be set.
# <翻译结束>

