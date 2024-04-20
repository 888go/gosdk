
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
// Package mgr can be used to manage Windows service programs.
// It can be used to install and remove them. It can also start,
// stop and pause them. The package can query / change current
// service state and config parameters.
<原文结束>

# <翻译开始>
// Package mgr can be used to manage Windows service programs.
// It can be used to install and remove them. It can also start,
// stop and pause them. The package can query / change current
// service state and config parameters.
# <翻译结束>


<原文开始>
// Mgr is used to manage Windows service.
<原文结束>

# <翻译开始>
// Mgr is used to manage Windows service.
# <翻译结束>


<原文开始>
// Connect establishes a connection to the service control manager.
<原文结束>

# <翻译开始>
// Connect establishes a connection to the service control manager.
# <翻译结束>


<原文开始>
// ConnectRemote establishes a connection to the
// service control manager on computer named host.
<原文结束>

# <翻译开始>
// ConnectRemote establishes a connection to the
// service control manager on computer named host.
# <翻译结束>


<原文开始>
// Disconnect closes connection to the service control manager m.
<原文结束>

# <翻译开始>
// Disconnect closes connection to the service control manager m.
# <翻译结束>


<原文开始>
// Whether the SCM has been locked.
<原文结束>

# <翻译开始>
// Whether the SCM has been locked.
# <翻译结束>


<原文开始>
// For how long the SCM has been locked.
<原文结束>

# <翻译开始>
// For how long the SCM has been locked.
# <翻译结束>


<原文开始>
// The name of the user who has locked the SCM.
<原文结束>

# <翻译开始>
// The name of the user who has locked the SCM.
# <翻译结束>


<原文开始>
// LockStatus returns whether the service control manager is locked by
// the system, for how long, and by whom. A locked SCM indicates that
// most service actions will block until the system unlocks the SCM.
<原文结束>

# <翻译开始>
// LockStatus returns whether the service control manager is locked by
// the system, for how long, and by whom. A locked SCM indicates that
// most service actions will block until the system unlocks the SCM.
# <翻译结束>


<原文开始>
// toStringBlock terminates strings in ss with 0, and then
// concatenates them together. It also adds extra 0 at the end.
<原文结束>

# <翻译开始>
// toStringBlock terminates strings in ss with 0, and then
// concatenates them together. It also adds extra 0 at the end.
# <翻译结束>


<原文开始>
// CreateService installs new service name on the system.
// The service will be executed by running exepath binary.
// Use config c to specify service parameters.
// Any args will be passed as command-line arguments when
// the service is started; these arguments are distinct from
// the arguments passed to Service.Start or via the "Start
// parameters" field in the service's Properties dialog box.
<原文结束>

# <翻译开始>
// CreateService installs new service name on the system.
// The service will be executed by running exepath binary.
// Use config c to specify service parameters.
// Any args will be passed as command-line arguments when
// the service is started; these arguments are distinct from
// the arguments passed to Service.Start or via the "Start
// parameters" field in the service's Properties dialog box.
# <翻译结束>


<原文开始>
// OpenService retrieves access to service name, so it can
// be interrogated and controlled.
<原文结束>

# <翻译开始>
// OpenService retrieves access to service name, so it can
// be interrogated and controlled.
# <翻译结束>


<原文开始>
// ListServices enumerates services in the specified
// service control manager database m.
// If the caller does not have the SERVICE_QUERY_STATUS
// access right to a service, the service is silently
// omitted from the list of services returned.
<原文结束>

# <翻译开始>
// ListServices enumerates services in the specified
// service control manager database m.
// If the caller does not have the SERVICE_QUERY_STATUS
// access right to a service, the service is silently
// omitted from the list of services returned.
# <翻译结束>

