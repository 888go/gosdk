
<原文开始>
// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
# <翻译结束>


<原文开始>
// Check that all methods of T are collected before
// determining the result type of m (which embeds
// all methods of T).
<原文结束>

# <翻译开始>
// Check that all methods of T are collected before
// determining the result type of m (which embeds
// all methods of T).
# <翻译结束>


<原文开始>
// Check that unresolved forward chains are followed
// (see also comment in resolver.go, checker.typeDecl).
<原文结束>

# <翻译开始>
// Check that unresolved forward chains are followed
// (see also comment in resolver.go, checker.typeDecl).
# <翻译结束>


<原文开始>
// Check that interface type comparison for identity
// does not recur endlessly.
<原文结束>

# <翻译开始>
// Check that interface type comparison for identity
// does not recur endlessly.
# <翻译结束>


<原文开始>
	// Checking for assignability of interfaces must check
	// if all methods of x are present in y, and that they
	// have identical signatures. The signatures recur via
	// the result type, which is an interface that embeds
	// a single method m that refers to the very interface
	// that contains it. This requires cycle detection in
	// identity checks for interface types.
<原文结束>

# <翻译开始>
	// Checking for assignability of interfaces must check
	// if all methods of x are present in y, and that they
	// have identical signatures. The signatures recur via
	// the result type, which is an interface that embeds
	// a single method m that refers to the very interface
	// that contains it. This requires cycle detection in
	// identity checks for interface types.
# <翻译结束>


<原文开始>
// Check that interfaces are type-checked in order of
// (embedded interface) dependencies (was issue 7158).
<原文结束>

# <翻译开始>
// Check that interfaces are type-checked in order of
// (embedded interface) dependencies (was issue 7158).
# <翻译结束>


<原文开始>
// Actual test case from issue 7158.
<原文结束>

# <翻译开始>
// Actual test case from issue 7158.
# <翻译结束>


<原文开始>
// Check that accessing an interface method too early doesn't lead
// to follow-on errors due to an incorrectly computed type set.
<原文结束>

# <翻译开始>
// Check that accessing an interface method too early doesn't lead
// to follow-on errors due to an incorrectly computed type set.
# <翻译结束>

