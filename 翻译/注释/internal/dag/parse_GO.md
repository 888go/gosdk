
<原文开始>
// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
# <翻译结束>


<原文开始>
// Package dag implements a language for expressing directed acyclic
// graphs.
//
// The general syntax of a rule is:
//
//	a, b < c, d;
//
// which means c and d come after a and b in the partial order
// (that is, there are edges from c and d to a and b),
// but doesn't provide a relative order between a vs b or c vs d.
//
// The rules can chain together, as in:
//
//	e < f, g < h;
//
// which is equivalent to
//
//	e < f, g;
//	f, g < h;
//
// Except for the special bottom element "NONE", each name
// must appear exactly once on the right-hand side of any rule.
// That rule serves as the definition of the allowed successor
// for that name. The definition must appear before any uses
// of the name on the left-hand side of a rule. (That is, the
// rules themselves must be ordered according to the partial
// order, for easier reading by people.)
//
// Negative assertions double-check the partial order:
//
//	i !< j
//
// means that it must NOT be the case that i < j.
// Negative assertions may appear anywhere in the rules,
// even before i and j have been defined.
//
// Comments begin with #.
<原文结束>

# <翻译开始>
// Package dag implements a language for expressing directed acyclic
// graphs.
//
// The general syntax of a rule is:
//
//	a, b < c, d;
//
// which means c and d come after a and b in the partial order
// (that is, there are edges from c and d to a and b),
// but doesn't provide a relative order between a vs b or c vs d.
//
// The rules can chain together, as in:
//
//	e < f, g < h;
//
// which is equivalent to
//
//	e < f, g;
//	f, g < h;
//
// Except for the special bottom element "NONE", each name
// must appear exactly once on the right-hand side of any rule.
// That rule serves as the definition of the allowed successor
// for that name. The definition must appear before any uses
// of the name on the left-hand side of a rule. (That is, the
// rules themselves must be ordered according to the partial
// order, for easier reading by people.)
//
// Negative assertions double-check the partial order:
//
//	i !< j
//
// means that it must NOT be the case that i < j.
// Negative assertions may appear anywhere in the rules,
// even before i and j have been defined.
//
// Comments begin with #.
# <翻译结束>


<原文开始>
// Parse parses the DAG language and returns the transitive closure of
// the described graph. In the returned graph, there is an edge from "b"
// to "a" if b < a (or a > b) in the partial order.
<原文结束>

# <翻译开始>
// Parse parses the DAG language and returns the transitive closure of
// the described graph. In the returned graph, there is an edge from "b"
// to "a" if b < a (or a > b) in the partial order.
# <翻译结束>


<原文开始>
// TODO: Add line numbers to errors.
<原文结束>

# <翻译开始>
// TODO: Add line numbers to errors.
# <翻译结束>


<原文开始>
// Check for missing definition.
<原文结束>

# <翻译开始>
// Check for missing definition.
# <翻译结束>


<原文开始>
// Complete transitive closure.
<原文结束>

# <翻译开始>
// Complete transitive closure.
# <翻译结束>


<原文开始>
						// Can only happen along with a "use of X before deps" error above,
						// but this error is more specific - it makes clear that reordering the
						// rules will not be enough to fix the problem.
<原文结束>

# <翻译开始>
						// Can only happen along with a "use of X before deps" error above,
						// but this error is more specific - it makes clear that reordering the
						// rules will not be enough to fix the problem.
# <翻译结束>


<原文开始>
// Check negative assertions against completed allowed graph.
<原文结束>

# <翻译开始>
// Check negative assertions against completed allowed graph.
# <翻译结束>


<原文开始>
// A rule is a line in the DAG language where "less < def" or "less !< def".
<原文结束>

# <翻译开始>
// A rule is a line in the DAG language where "less < def" or "less !< def".
# <翻译结束>


<原文开始>
// parseRules parses the rules of a DAG.
<原文结束>

# <翻译开始>
// parseRules parses the rules of a DAG.
# <翻译结束>


<原文开始>
// A rulesParser parses the depsRules syntax described above.
<原文结束>

# <翻译开始>
// A rulesParser parses the depsRules syntax described above.
# <翻译结束>


<原文开始>
// syntaxError reports a parsing error.
<原文结束>

# <翻译开始>
// syntaxError reports a parsing error.
# <翻译结束>


<原文开始>
// nextList parses and returns a comma-separated list of names.
<原文结束>

# <翻译开始>
// nextList parses and returns a comma-separated list of names.
# <翻译结束>


<原文开始>
// nextToken returns the next token in the deps rules,
// one of ";" "," "<" "!<" or a name.
<原文结束>

# <翻译开始>
// nextToken returns the next token in the deps rules,
// one of ";" "," "<" "!<" or a name.
# <翻译结束>

