
<原文开始>
// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
# <翻译结束>


<原文开始>
// PanicOnExit0 reports whether to panic on a call to os.Exit(0).
// This is in the testlog package because, like other definitions in
// package testlog, it is a hook between the testing package and the
// os package. This is used to ensure that an early call to os.Exit(0)
// does not cause a test to pass.
<原文结束>

# <翻译开始>
// PanicOnExit0 reports whether to panic on a call to os.Exit(0).
// This is in the testlog package because, like other definitions in
// package testlog, it is a hook between the testing package and the
// os package. This is used to ensure that an early call to os.Exit(0)
// does not cause a test to pass.
# <翻译结束>


<原文开始>
// panicOnExit0 is the flag used for PanicOnExit0. This uses a lock
// because the value can be cleared via a timer call that may race
// with calls to os.Exit
<原文结束>

# <翻译开始>
// panicOnExit0 is the flag used for PanicOnExit0. This uses a lock
// because the value can be cleared via a timer call that may race
// with calls to os.Exit
# <翻译结束>


<原文开始>
// SetPanicOnExit0 sets panicOnExit0 to v.
<原文结束>

# <翻译开始>
// SetPanicOnExit0 sets panicOnExit0 to v.
# <翻译结束>

