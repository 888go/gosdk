
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
// Service is used to access Windows service.
<原文结束>

# <翻译开始>
// Service is used to access Windows service.
# <翻译结束>


<原文开始>
// Delete marks service s for deletion from the service control manager database.
<原文结束>

# <翻译开始>
// Delete marks service s for deletion from the service control manager database.
# <翻译结束>


<原文开始>
// Close relinquish access to the service s.
<原文结束>

# <翻译开始>
// Close relinquish access to the service s.
# <翻译结束>


<原文开始>
// Start starts service s.
// args will be passed to svc.Handler.Execute.
<原文结束>

# <翻译开始>
// Start starts service s.
// args will be passed to svc.Handler.Execute.
# <翻译结束>


<原文开始>
// Control sends state change request c to the service s. It returns the most
// recent status the service reported to the service control manager, and an
// error if the state change request was not accepted.
// Note that the returned service status is only set if the status change
// request succeeded, or if it failed with error ERROR_INVALID_SERVICE_CONTROL,
// ERROR_SERVICE_CANNOT_ACCEPT_CTRL, or ERROR_SERVICE_NOT_ACTIVE.
<原文结束>

# <翻译开始>
// Control sends state change request c to the service s. It returns the most
// recent status the service reported to the service control manager, and an
// error if the state change request was not accepted.
// Note that the returned service status is only set if the status change
// request succeeded, or if it failed with error ERROR_INVALID_SERVICE_CONTROL,
// ERROR_SERVICE_CANNOT_ACCEPT_CTRL, or ERROR_SERVICE_NOT_ACTIVE.
# <翻译结束>


<原文开始>
// Query returns current status of service s.
<原文结束>

# <翻译开始>
// Query returns current status of service s.
# <翻译结束>


<原文开始>
// ListDependentServices returns the names of the services dependent on service s, which match the given status.
<原文结束>

# <翻译开始>
// ListDependentServices returns the names of the services dependent on service s, which match the given status.
# <翻译结束>


<原文开始>
// ERROR_MORE_DATA indicates the provided buffer was too small, run the call again after resizing the buffer
<原文结束>

# <翻译开始>
// ERROR_MORE_DATA indicates the provided buffer was too small, run the call again after resizing the buffer
# <翻译结束>


<原文开始>
	// The slice mutated by EnumDependentServices may have a length greater than returnedServiceCount, any elements
	// past that should be ignored.
<原文结束>

# <翻译开始>
	// The slice mutated by EnumDependentServices may have a length greater than returnedServiceCount, any elements
	// past that should be ignored.
# <翻译结束>

