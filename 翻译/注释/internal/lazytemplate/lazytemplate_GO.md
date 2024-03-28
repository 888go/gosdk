
<原文开始>
// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
# <翻译结束>


<原文开始>
// Package lazytemplate is a thin wrapper over text/template, allowing the use
// of global template variables without forcing them to be parsed at init.
<原文结束>

# <翻译开始>
// Package lazytemplate is a thin wrapper over text/template, allowing the use
// of global template variables without forcing them to be parsed at init.
# <翻译结束>


<原文开始>
// Template is a wrapper around text/template.Template, where the underlying
// template will be parsed the first time it is needed.
<原文结束>

# <翻译开始>
// Template is a wrapper around text/template.Template, where the underlying
// template will be parsed the first time it is needed.
# <翻译结束>


<原文开始>
// New creates a new lazy template, delaying the parsing work until it is first
// needed. If the code is being run as part of tests, the template parsing will
// happen immediately.
<原文结束>

# <翻译开始>
// New creates a new lazy template, delaying the parsing work until it is first
// needed. If the code is being run as part of tests, the template parsing will
// happen immediately.
# <翻译结束>


<原文开始>
// In tests, always parse the templates early.
<原文结束>

# <翻译开始>
// In tests, always parse the templates early.
# <翻译结束>

