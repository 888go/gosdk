
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
// This file defines the error codes that can be produced during type-checking.
// Collectively, these codes provide an identifier that may be used to
// implement special handling for certain types of errors.
//
// Error code values should not be changed: add new codes at the end.
//
// Error codes should be fine-grained enough that the exact nature of the error
// can be easily determined, but coarse enough that they are not an
// implementation detail of the type checking algorithm. As a rule-of-thumb,
// errors should be considered equivalent if there is a theoretical refactoring
// of the type checker in which they are emitted in exactly one place. For
// example, the type checker emits different error messages for "too many
// arguments" and "too few arguments", but one can imagine an alternative type
// checker where this check instead just emits a single "wrong number of
// arguments", so these errors should have the same code.
//
// Error code names should be as brief as possible while retaining accuracy and
// distinctiveness. In most cases names should start with an adjective
// describing the nature of the error (e.g. "invalid", "unused", "misplaced"),
// and end with a noun identifying the relevant language object. For example,
// "_DuplicateDecl" or "_InvalidSliceExpr". For brevity, naming follows the
// convention that "bad" implies a problem with syntax, and "invalid" implies a
// problem with types.
<原文结束>

# <翻译开始>
// This file defines the error codes that can be produced during type-checking.
// Collectively, these codes provide an identifier that may be used to
// implement special handling for certain types of errors.
//
// Error code values should not be changed: add new codes at the end.
//
// Error codes should be fine-grained enough that the exact nature of the error
// can be easily determined, but coarse enough that they are not an
// implementation detail of the type checking algorithm. As a rule-of-thumb,
// errors should be considered equivalent if there is a theoretical refactoring
// of the type checker in which they are emitted in exactly one place. For
// example, the type checker emits different error messages for "too many
// arguments" and "too few arguments", but one can imagine an alternative type
// checker where this check instead just emits a single "wrong number of
// arguments", so these errors should have the same code.
//
// Error code names should be as brief as possible while retaining accuracy and
// distinctiveness. In most cases names should start with an adjective
// describing the nature of the error (e.g. "invalid", "unused", "misplaced"),
// and end with a noun identifying the relevant language object. For example,
// "_DuplicateDecl" or "_InvalidSliceExpr". For brevity, naming follows the
// convention that "bad" implies a problem with syntax, and "invalid" implies a
// problem with types.
# <翻译结束>


<原文开始>
	// InvalidSyntaxTree occurs if an invalid syntax tree is provided
	// to the type checker. It should never happen.
<原文结束>

# <翻译开始>
	// InvalidSyntaxTree occurs if an invalid syntax tree is provided
	// to the type checker. It should never happen.
# <翻译结束>


<原文开始>
// The zero Code value indicates an unset (invalid) error code.
<原文结束>

# <翻译开始>
// The zero Code value indicates an unset (invalid) error code.
# <翻译结束>


<原文开始>
// Test is reserved for errors that only apply while in self-test mode.
<原文结束>

# <翻译开始>
// Test is reserved for errors that only apply while in self-test mode.
# <翻译结束>


<原文开始>
	// BlankPkgName occurs when a package name is the blank identifier "_".
	//
	// Per the spec:
	//  "The PackageName must not be the blank identifier."
	//
	// Example:
	//  package _
<原文结束>

# <翻译开始>
	// BlankPkgName occurs when a package name is the blank identifier "_".
	//
	// Per the spec:
	//  "The PackageName must not be the blank identifier."
	//
	// Example:
	//  package _
# <翻译结束>


<原文开始>
	// MismatchedPkgName occurs when a file's package name doesn't match the
	// package name already established by other files.
<原文结束>

# <翻译开始>
	// MismatchedPkgName occurs when a file's package name doesn't match the
	// package name already established by other files.
# <翻译结束>


<原文开始>
	// InvalidPkgUse occurs when a package identifier is used outside of a
	// selector expression.
	//
	// Example:
	//  import "fmt"
	//
	//  var _ = fmt
<原文结束>

# <翻译开始>
	// InvalidPkgUse occurs when a package identifier is used outside of a
	// selector expression.
	//
	// Example:
	//  import "fmt"
	//
	//  var _ = fmt
# <翻译结束>


<原文开始>
// BadImportPath occurs when an import path is not valid.
<原文结束>

# <翻译开始>
// BadImportPath occurs when an import path is not valid.
# <翻译结束>


<原文开始>
	// BrokenImport occurs when importing a package fails.
	//
	// Example:
	//  import "amissingpackage"
<原文结束>

# <翻译开始>
	// BrokenImport occurs when importing a package fails.
	//
	// Example:
	//  import "amissingpackage"
# <翻译结束>


<原文开始>
	// ImportCRenamed occurs when the special import "C" is renamed. "C" is a
	// pseudo-package, and must not be renamed.
	//
	// Example:
	//  import _ "C"
<原文结束>

# <翻译开始>
	// ImportCRenamed occurs when the special import "C" is renamed. "C" is a
	// pseudo-package, and must not be renamed.
	//
	// Example:
	//  import _ "C"
# <翻译结束>


<原文开始>
	// UnusedImport occurs when an import is unused.
	//
	// Example:
	//  import "fmt"
	//
	//  func main() {}
<原文结束>

# <翻译开始>
	// UnusedImport occurs when an import is unused.
	//
	// Example:
	//  import "fmt"
	//
	//  func main() {}
# <翻译结束>


<原文开始>
	// InvalidInitCycle occurs when an invalid cycle is detected within the
	// initialization graph.
	//
	// Example:
	//  var x int = f()
	//
	//  func f() int { return x }
<原文结束>

# <翻译开始>
	// InvalidInitCycle occurs when an invalid cycle is detected within the
	// initialization graph.
	//
	// Example:
	//  var x int = f()
	//
	//  func f() int { return x }
# <翻译结束>


<原文开始>
	// DuplicateDecl occurs when an identifier is declared multiple times.
	//
	// Example:
	//  var x = 1
	//  var x = 2
<原文结束>

# <翻译开始>
	// DuplicateDecl occurs when an identifier is declared multiple times.
	//
	// Example:
	//  var x = 1
	//  var x = 2
# <翻译结束>


<原文开始>
	// InvalidDeclCycle occurs when a declaration cycle is not valid.
	//
	// Example:
	//  type S struct {
	//  	S
	//  }
	//
<原文结束>

# <翻译开始>
	// InvalidDeclCycle occurs when a declaration cycle is not valid.
	//
	// Example:
	//  type S struct {
	//  	S
	//  }
	//
# <翻译结束>


<原文开始>
	// InvalidTypeCycle occurs when a cycle in type definitions results in a
	// type that is not well-defined.
	//
	// Example:
	//  import "unsafe"
	//
	//  type T [unsafe.Sizeof(T{})]int
<原文结束>

# <翻译开始>
	// InvalidTypeCycle occurs when a cycle in type definitions results in a
	// type that is not well-defined.
	//
	// Example:
	//  import "unsafe"
	//
	//  type T [unsafe.Sizeof(T{})]int
# <翻译结束>


<原文开始>
	// InvalidConstInit occurs when a const declaration has a non-constant
	// initializer.
	//
	// Example:
	//  var x int
	//  const _ = x
<原文结束>

# <翻译开始>
	// InvalidConstInit occurs when a const declaration has a non-constant
	// initializer.
	//
	// Example:
	//  var x int
	//  const _ = x
# <翻译结束>


<原文开始>
	// InvalidConstVal occurs when a const value cannot be converted to its
	// target type.
	//
	// TODO(findleyr): this error code and example are not very clear. Consider
	// removing it.
	//
	// Example:
	//  const _ = 1 << "hello"
<原文结束>

# <翻译开始>
	// InvalidConstVal occurs when a const value cannot be converted to its
	// target type.
	//
	// TODO(findleyr): this error code and example are not very clear. Consider
	// removing it.
	//
	// Example:
	//  const _ = 1 << "hello"
# <翻译结束>


<原文开始>
	// InvalidConstType occurs when the underlying type in a const declaration
	// is not a valid constant type.
	//
	// Example:
	//  const c *int = 4
<原文结束>

# <翻译开始>
	// InvalidConstType occurs when the underlying type in a const declaration
	// is not a valid constant type.
	//
	// Example:
	//  const c *int = 4
# <翻译结束>


<原文开始>
	// UntypedNilUse occurs when the predeclared (untyped) value nil is used to
	// initialize a variable declared without an explicit type.
	//
	// Example:
	//  var x = nil
<原文结束>

# <翻译开始>
	// UntypedNilUse occurs when the predeclared (untyped) value nil is used to
	// initialize a variable declared without an explicit type.
	//
	// Example:
	//  var x = nil
# <翻译结束>


<原文开始>
	// WrongAssignCount occurs when the number of values on the right-hand side
	// of an assignment or initialization expression does not match the number
	// of variables on the left-hand side.
	//
	// Example:
	//  var x = 1, 2
<原文结束>

# <翻译开始>
	// WrongAssignCount occurs when the number of values on the right-hand side
	// of an assignment or initialization expression does not match the number
	// of variables on the left-hand side.
	//
	// Example:
	//  var x = 1, 2
# <翻译结束>


<原文开始>
	// UnassignableOperand occurs when the left-hand side of an assignment is
	// not assignable.
	//
	// Example:
	//  func f() {
	//  	const c = 1
	//  	c = 2
	//  }
<原文结束>

# <翻译开始>
	// UnassignableOperand occurs when the left-hand side of an assignment is
	// not assignable.
	//
	// Example:
	//  func f() {
	//  	const c = 1
	//  	c = 2
	//  }
# <翻译结束>


<原文开始>
	// NoNewVar occurs when a short variable declaration (':=') does not declare
	// new variables.
	//
	// Example:
	//  func f() {
	//  	x := 1
	//  	x := 2
	//  }
<原文结束>

# <翻译开始>
	// NoNewVar occurs when a short variable declaration (':=') does not declare
	// new variables.
	//
	// Example:
	//  func f() {
	//  	x := 1
	//  	x := 2
	//  }
# <翻译结束>


<原文开始>
	// MultiValAssignOp occurs when an assignment operation (+=, *=, etc) does
	// not have single-valued left-hand or right-hand side.
	//
	// Per the spec:
	//  "In assignment operations, both the left- and right-hand expression lists
	//  must contain exactly one single-valued expression"
	//
	// Example:
	//  func f() int {
	//  	x, y := 1, 2
	//  	x, y += 1
	//  	return x + y
	//  }
<原文结束>

# <翻译开始>
	// MultiValAssignOp occurs when an assignment operation (+=, *=, etc) does
	// not have single-valued left-hand or right-hand side.
	//
	// Per the spec:
	//  "In assignment operations, both the left- and right-hand expression lists
	//  must contain exactly one single-valued expression"
	//
	// Example:
	//  func f() int {
	//  	x, y := 1, 2
	//  	x, y += 1
	//  	return x + y
	//  }
# <翻译结束>


<原文开始>
	// InvalidIfaceAssign occurs when a value of type T is used as an
	// interface, but T does not implement a method of the expected interface.
	//
	// Example:
	//  type I interface {
	//  	f()
	//  }
	//
	//  type T int
	//
	//  var x I = T(1)
<原文结束>

# <翻译开始>
	// InvalidIfaceAssign occurs when a value of type T is used as an
	// interface, but T does not implement a method of the expected interface.
	//
	// Example:
	//  type I interface {
	//  	f()
	//  }
	//
	//  type T int
	//
	//  var x I = T(1)
# <翻译结束>


<原文开始>
	// InvalidChanAssign occurs when a chan assignment is invalid.
	//
	// Per the spec, a value x is assignable to a channel type T if:
	//  "x is a bidirectional channel value, T is a channel type, x's type V and
	//  T have identical element types, and at least one of V or T is not a
	//  defined type."
	//
	// Example:
	//  type T1 chan int
	//  type T2 chan int
	//
	//  var x T1
	//  // Invalid assignment because both types are named
	//  var _ T2 = x
<原文结束>

# <翻译开始>
	// InvalidChanAssign occurs when a chan assignment is invalid.
	//
	// Per the spec, a value x is assignable to a channel type T if:
	//  "x is a bidirectional channel value, T is a channel type, x's type V and
	//  T have identical element types, and at least one of V or T is not a
	//  defined type."
	//
	// Example:
	//  type T1 chan int
	//  type T2 chan int
	//
	//  var x T1
	//  // Invalid assignment because both types are named
	//  var _ T2 = x
# <翻译结束>


<原文开始>
	// IncompatibleAssign occurs when the type of the right-hand side expression
	// in an assignment cannot be assigned to the type of the variable being
	// assigned.
	//
	// Example:
	//  var x []int
	//  var _ int = x
<原文结束>

# <翻译开始>
	// IncompatibleAssign occurs when the type of the right-hand side expression
	// in an assignment cannot be assigned to the type of the variable being
	// assigned.
	//
	// Example:
	//  var x []int
	//  var _ int = x
# <翻译结束>


<原文开始>
	// UnaddressableFieldAssign occurs when trying to assign to a struct field
	// in a map value.
	//
	// Example:
	//  func f() {
	//  	m := make(map[string]struct{i int})
	//  	m["foo"].i = 42
	//  }
<原文结束>

# <翻译开始>
	// UnaddressableFieldAssign occurs when trying to assign to a struct field
	// in a map value.
	//
	// Example:
	//  func f() {
	//  	m := make(map[string]struct{i int})
	//  	m["foo"].i = 42
	//  }
# <翻译结束>


<原文开始>
	// NotAType occurs when the identifier used as the underlying type in a type
	// declaration or the right-hand side of a type alias does not denote a type.
	//
	// Example:
	//  var S = 2
	//
	//  type T S
<原文结束>

# <翻译开始>
	// NotAType occurs when the identifier used as the underlying type in a type
	// declaration or the right-hand side of a type alias does not denote a type.
	//
	// Example:
	//  var S = 2
	//
	//  type T S
# <翻译结束>


<原文开始>
	// InvalidArrayLen occurs when an array length is not a constant value.
	//
	// Example:
	//  var n = 3
	//  var _ = [n]int{}
<原文结束>

# <翻译开始>
	// InvalidArrayLen occurs when an array length is not a constant value.
	//
	// Example:
	//  var n = 3
	//  var _ = [n]int{}
# <翻译结束>


<原文开始>
	// BlankIfaceMethod occurs when a method name is '_'.
	//
	// Per the spec:
	//  "The name of each explicitly specified method must be unique and not
	//  blank."
	//
	// Example:
	//  type T interface {
	//  	_(int)
	//  }
<原文结束>

# <翻译开始>
	// BlankIfaceMethod occurs when a method name is '_'.
	//
	// Per the spec:
	//  "The name of each explicitly specified method must be unique and not
	//  blank."
	//
	// Example:
	//  type T interface {
	//  	_(int)
	//  }
# <翻译结束>


<原文开始>
	// IncomparableMapKey occurs when a map key type does not support the == and
	// != operators.
	//
	// Per the spec:
	//  "The comparison operators == and != must be fully defined for operands of
	//  the key type; thus the key type must not be a function, map, or slice."
	//
	// Example:
	//  var x map[T]int
	//
	//  type T []int
<原文结束>

# <翻译开始>
	// IncomparableMapKey occurs when a map key type does not support the == and
	// != operators.
	//
	// Per the spec:
	//  "The comparison operators == and != must be fully defined for operands of
	//  the key type; thus the key type must not be a function, map, or slice."
	//
	// Example:
	//  var x map[T]int
	//
	//  type T []int
# <翻译结束>


<原文开始>
	// InvalidIfaceEmbed occurs when a non-interface type is embedded in an
	// interface (for go 1.17 or earlier).
<原文结束>

# <翻译开始>
	// InvalidIfaceEmbed occurs when a non-interface type is embedded in an
	// interface (for go 1.17 or earlier).
# <翻译结束>


<原文开始>
	// InvalidPtrEmbed occurs when an embedded field is of the pointer form *T,
	// and T itself is itself a pointer, an unsafe.Pointer, or an interface.
	//
	// Per the spec:
	//  "An embedded field must be specified as a type name T or as a pointer to
	//  a non-interface type name *T, and T itself may not be a pointer type."
	//
	// Example:
	//  type T *int
	//
	//  type S struct {
	//  	*T
	//  }
<原文结束>

# <翻译开始>
	// InvalidPtrEmbed occurs when an embedded field is of the pointer form *T,
	// and T itself is itself a pointer, an unsafe.Pointer, or an interface.
	//
	// Per the spec:
	//  "An embedded field must be specified as a type name T or as a pointer to
	//  a non-interface type name *T, and T itself may not be a pointer type."
	//
	// Example:
	//  type T *int
	//
	//  type S struct {
	//  	*T
	//  }
# <翻译结束>


<原文开始>
	// BadRecv occurs when a method declaration does not have exactly one
	// receiver parameter.
	//
	// Example:
	//  func () _() {}
<原文结束>

# <翻译开始>
	// BadRecv occurs when a method declaration does not have exactly one
	// receiver parameter.
	//
	// Example:
	//  func () _() {}
# <翻译结束>


<原文开始>
	// InvalidRecv occurs when a receiver type expression is not of the form T
	// or *T, or T is a pointer type.
	//
	// Example:
	//  type T struct {}
	//
	//  func (**T) m() {}
<原文结束>

# <翻译开始>
	// InvalidRecv occurs when a receiver type expression is not of the form T
	// or *T, or T is a pointer type.
	//
	// Example:
	//  type T struct {}
	//
	//  func (**T) m() {}
# <翻译结束>


<原文开始>
	// DuplicateFieldAndMethod occurs when an identifier appears as both a field
	// and method name.
	//
	// Example:
	//  type T struct {
	//  	m int
	//  }
	//
	//  func (T) m() {}
<原文结束>

# <翻译开始>
	// DuplicateFieldAndMethod occurs when an identifier appears as both a field
	// and method name.
	//
	// Example:
	//  type T struct {
	//  	m int
	//  }
	//
	//  func (T) m() {}
# <翻译结束>


<原文开始>
	// DuplicateMethod occurs when two methods on the same receiver type have
	// the same name.
	//
	// Example:
	//  type T struct {}
	//  func (T) m() {}
	//  func (T) m(i int) int { return i }
<原文结束>

# <翻译开始>
	// DuplicateMethod occurs when two methods on the same receiver type have
	// the same name.
	//
	// Example:
	//  type T struct {}
	//  func (T) m() {}
	//  func (T) m(i int) int { return i }
# <翻译结束>


<原文开始>
	// InvalidBlank occurs when a blank identifier is used as a value or type.
	//
	// Per the spec:
	//  "The blank identifier may appear as an operand only on the left-hand side
	//  of an assignment."
	//
	// Example:
	//  var x = _
<原文结束>

# <翻译开始>
	// InvalidBlank occurs when a blank identifier is used as a value or type.
	//
	// Per the spec:
	//  "The blank identifier may appear as an operand only on the left-hand side
	//  of an assignment."
	//
	// Example:
	//  var x = _
# <翻译结束>


<原文开始>
	// InvalidIota occurs when the predeclared identifier iota is used outside
	// of a constant declaration.
	//
	// Example:
	//  var x = iota
<原文结束>

# <翻译开始>
	// InvalidIota occurs when the predeclared identifier iota is used outside
	// of a constant declaration.
	//
	// Example:
	//  var x = iota
# <翻译结束>


<原文开始>
	// MissingInitBody occurs when an init function is missing its body.
	//
	// Example:
	//  func init()
<原文结束>

# <翻译开始>
	// MissingInitBody occurs when an init function is missing its body.
	//
	// Example:
	//  func init()
# <翻译结束>


<原文开始>
	// InvalidInitSig occurs when an init function declares parameters or
	// results.
	//
	// Deprecated: no longer emitted by the type checker. _InvalidInitDecl is
	// used instead.
<原文结束>

# <翻译开始>
	// InvalidInitSig occurs when an init function declares parameters or
	// results.
	//
	// Deprecated: no longer emitted by the type checker. _InvalidInitDecl is
	// used instead.
# <翻译结束>


<原文开始>
	// InvalidInitDecl occurs when init is declared as anything other than a
	// function.
	//
	// Example:
	//  var init = 1
	//
	// Example:
	//  func init() int { return 1 }
<原文结束>

# <翻译开始>
	// InvalidInitDecl occurs when init is declared as anything other than a
	// function.
	//
	// Example:
	//  var init = 1
	//
	// Example:
	//  func init() int { return 1 }
# <翻译结束>


<原文开始>
	// InvalidMainDecl occurs when main is declared as anything other than a
	// function, in a main package.
<原文结束>

# <翻译开始>
	// InvalidMainDecl occurs when main is declared as anything other than a
	// function, in a main package.
# <翻译结束>


<原文开始>
	// TooManyValues occurs when a function returns too many values for the
	// expression context in which it is used.
	//
	// Example:
	//  func ReturnTwo() (int, int) {
	//  	return 1, 2
	//  }
	//
	//  var x = ReturnTwo()
<原文结束>

# <翻译开始>
	// TooManyValues occurs when a function returns too many values for the
	// expression context in which it is used.
	//
	// Example:
	//  func ReturnTwo() (int, int) {
	//  	return 1, 2
	//  }
	//
	//  var x = ReturnTwo()
# <翻译结束>


<原文开始>
	// NotAnExpr occurs when a type expression is used where a value expression
	// is expected.
	//
	// Example:
	//  type T struct {}
	//
	//  func f() {
	//  	T
	//  }
<原文结束>

# <翻译开始>
	// NotAnExpr occurs when a type expression is used where a value expression
	// is expected.
	//
	// Example:
	//  type T struct {}
	//
	//  func f() {
	//  	T
	//  }
# <翻译结束>


<原文开始>
	// TruncatedFloat occurs when a float constant is truncated to an integer
	// value.
	//
	// Example:
	//  var _ int = 98.6
<原文结束>

# <翻译开始>
	// TruncatedFloat occurs when a float constant is truncated to an integer
	// value.
	//
	// Example:
	//  var _ int = 98.6
# <翻译结束>


<原文开始>
	// NumericOverflow occurs when a numeric constant overflows its target type.
	//
	// Example:
	//  var x int8 = 1000
<原文结束>

# <翻译开始>
	// NumericOverflow occurs when a numeric constant overflows its target type.
	//
	// Example:
	//  var x int8 = 1000
# <翻译结束>


<原文开始>
	// UndefinedOp occurs when an operator is not defined for the type(s) used
	// in an operation.
	//
	// Example:
	//  var c = "a" - "b"
<原文结束>

# <翻译开始>
	// UndefinedOp occurs when an operator is not defined for the type(s) used
	// in an operation.
	//
	// Example:
	//  var c = "a" - "b"
# <翻译结束>


<原文开始>
	// MismatchedTypes occurs when operand types are incompatible in a binary
	// operation.
	//
	// Example:
	//  var a = "hello"
	//  var b = 1
	//  var c = a - b
<原文结束>

# <翻译开始>
	// MismatchedTypes occurs when operand types are incompatible in a binary
	// operation.
	//
	// Example:
	//  var a = "hello"
	//  var b = 1
	//  var c = a - b
# <翻译结束>


<原文开始>
	// DivByZero occurs when a division operation is provable at compile
	// time to be a division by zero.
	//
	// Example:
	//  const divisor = 0
	//  var x int = 1/divisor
<原文结束>

# <翻译开始>
	// DivByZero occurs when a division operation is provable at compile
	// time to be a division by zero.
	//
	// Example:
	//  const divisor = 0
	//  var x int = 1/divisor
# <翻译结束>


<原文开始>
	// NonNumericIncDec occurs when an increment or decrement operator is
	// applied to a non-numeric value.
	//
	// Example:
	//  func f() {
	//  	var c = "c"
	//  	c++
	//  }
<原文结束>

# <翻译开始>
	// NonNumericIncDec occurs when an increment or decrement operator is
	// applied to a non-numeric value.
	//
	// Example:
	//  func f() {
	//  	var c = "c"
	//  	c++
	//  }
# <翻译结束>


<原文开始>
	// UnaddressableOperand occurs when the & operator is applied to an
	// unaddressable expression.
	//
	// Example:
	//  var x = &1
<原文结束>

# <翻译开始>
	// UnaddressableOperand occurs when the & operator is applied to an
	// unaddressable expression.
	//
	// Example:
	//  var x = &1
# <翻译结束>


<原文开始>
	// InvalidIndirection occurs when a non-pointer value is indirected via the
	// '*' operator.
	//
	// Example:
	//  var x int
	//  var y = *x
<原文结束>

# <翻译开始>
	// InvalidIndirection occurs when a non-pointer value is indirected via the
	// '*' operator.
	//
	// Example:
	//  var x int
	//  var y = *x
# <翻译结束>


<原文开始>
	// NonIndexableOperand occurs when an index operation is applied to a value
	// that cannot be indexed.
	//
	// Example:
	//  var x = 1
	//  var y = x[1]
<原文结束>

# <翻译开始>
	// NonIndexableOperand occurs when an index operation is applied to a value
	// that cannot be indexed.
	//
	// Example:
	//  var x = 1
	//  var y = x[1]
# <翻译结束>


<原文开始>
	// InvalidIndex occurs when an index argument is not of integer type,
	// negative, or out-of-bounds.
	//
	// Example:
	//  var s = [...]int{1,2,3}
	//  var x = s[5]
	//
	// Example:
	//  var s = []int{1,2,3}
	//  var _ = s[-1]
	//
	// Example:
	//  var s = []int{1,2,3}
	//  var i string
	//  var _ = s[i]
<原文结束>

# <翻译开始>
	// InvalidIndex occurs when an index argument is not of integer type,
	// negative, or out-of-bounds.
	//
	// Example:
	//  var s = [...]int{1,2,3}
	//  var x = s[5]
	//
	// Example:
	//  var s = []int{1,2,3}
	//  var _ = s[-1]
	//
	// Example:
	//  var s = []int{1,2,3}
	//  var i string
	//  var _ = s[i]
# <翻译结束>


<原文开始>
	// SwappedSliceIndices occurs when constant indices in a slice expression
	// are decreasing in value.
	//
	// Example:
	//  var _ = []int{1,2,3}[2:1]
<原文结束>

# <翻译开始>
	// SwappedSliceIndices occurs when constant indices in a slice expression
	// are decreasing in value.
	//
	// Example:
	//  var _ = []int{1,2,3}[2:1]
# <翻译结束>


<原文开始>
	// NonSliceableOperand occurs when a slice operation is applied to a value
	// whose type is not sliceable, or is unaddressable.
	//
	// Example:
	//  var x = [...]int{1, 2, 3}[:1]
	//
	// Example:
	//  var x = 1
	//  var y = 1[:1]
<原文结束>

# <翻译开始>
	// NonSliceableOperand occurs when a slice operation is applied to a value
	// whose type is not sliceable, or is unaddressable.
	//
	// Example:
	//  var x = [...]int{1, 2, 3}[:1]
	//
	// Example:
	//  var x = 1
	//  var y = 1[:1]
# <翻译结束>


<原文开始>
	// InvalidSliceExpr occurs when a three-index slice expression (a[x:y:z]) is
	// applied to a string.
	//
	// Example:
	//  var s = "hello"
	//  var x = s[1:2:3]
<原文结束>

# <翻译开始>
	// InvalidSliceExpr occurs when a three-index slice expression (a[x:y:z]) is
	// applied to a string.
	//
	// Example:
	//  var s = "hello"
	//  var x = s[1:2:3]
# <翻译结束>


<原文开始>
	// InvalidShiftCount occurs when the right-hand side of a shift operation is
	// either non-integer, negative, or too large.
	//
	// Example:
	//  var (
	//  	x string
	//  	y int = 1 << x
	//  )
<原文结束>

# <翻译开始>
	// InvalidShiftCount occurs when the right-hand side of a shift operation is
	// either non-integer, negative, or too large.
	//
	// Example:
	//  var (
	//  	x string
	//  	y int = 1 << x
	//  )
# <翻译结束>


<原文开始>
	// InvalidShiftOperand occurs when the shifted operand is not an integer.
	//
	// Example:
	//  var s = "hello"
	//  var x = s << 2
<原文结束>

# <翻译开始>
	// InvalidShiftOperand occurs when the shifted operand is not an integer.
	//
	// Example:
	//  var s = "hello"
	//  var x = s << 2
# <翻译结束>


<原文开始>
	// InvalidReceive occurs when there is a channel receive from a value that
	// is either not a channel, or is a send-only channel.
	//
	// Example:
	//  func f() {
	//  	var x = 1
	//  	<-x
	//  }
<原文结束>

# <翻译开始>
	// InvalidReceive occurs when there is a channel receive from a value that
	// is either not a channel, or is a send-only channel.
	//
	// Example:
	//  func f() {
	//  	var x = 1
	//  	<-x
	//  }
# <翻译结束>


<原文开始>
	// InvalidSend occurs when there is a channel send to a value that is not a
	// channel, or is a receive-only channel.
	//
	// Example:
	//  func f() {
	//  	var x = 1
	//  	x <- "hello!"
	//  }
<原文结束>

# <翻译开始>
	// InvalidSend occurs when there is a channel send to a value that is not a
	// channel, or is a receive-only channel.
	//
	// Example:
	//  func f() {
	//  	var x = 1
	//  	x <- "hello!"
	//  }
# <翻译结束>


<原文开始>
	// DuplicateLitKey occurs when an index is duplicated in a slice, array, or
	// map literal.
	//
	// Example:
	//  var _ = []int{0:1, 0:2}
	//
	// Example:
	//  var _ = map[string]int{"a": 1, "a": 2}
<原文结束>

# <翻译开始>
	// DuplicateLitKey occurs when an index is duplicated in a slice, array, or
	// map literal.
	//
	// Example:
	//  var _ = []int{0:1, 0:2}
	//
	// Example:
	//  var _ = map[string]int{"a": 1, "a": 2}
# <翻译结束>


<原文开始>
	// MissingLitKey occurs when a map literal is missing a key expression.
	//
	// Example:
	//  var _ = map[string]int{1}
<原文结束>

# <翻译开始>
	// MissingLitKey occurs when a map literal is missing a key expression.
	//
	// Example:
	//  var _ = map[string]int{1}
# <翻译结束>


<原文开始>
	// InvalidLitIndex occurs when the key in a key-value element of a slice or
	// array literal is not an integer constant.
	//
	// Example:
	//  var i = 0
	//  var x = []string{i: "world"}
<原文结束>

# <翻译开始>
	// InvalidLitIndex occurs when the key in a key-value element of a slice or
	// array literal is not an integer constant.
	//
	// Example:
	//  var i = 0
	//  var x = []string{i: "world"}
# <翻译结束>


<原文开始>
	// OversizeArrayLit occurs when an array literal exceeds its length.
	//
	// Example:
	//  var _ = [2]int{1,2,3}
<原文结束>

# <翻译开始>
	// OversizeArrayLit occurs when an array literal exceeds its length.
	//
	// Example:
	//  var _ = [2]int{1,2,3}
# <翻译结束>


<原文开始>
	// MixedStructLit occurs when a struct literal contains a mix of positional
	// and named elements.
	//
	// Example:
	//  var _ = struct{i, j int}{i: 1, 2}
<原文结束>

# <翻译开始>
	// MixedStructLit occurs when a struct literal contains a mix of positional
	// and named elements.
	//
	// Example:
	//  var _ = struct{i, j int}{i: 1, 2}
# <翻译结束>


<原文开始>
	// InvalidStructLit occurs when a positional struct literal has an incorrect
	// number of values.
	//
	// Example:
	//  var _ = struct{i, j int}{1,2,3}
<原文结束>

# <翻译开始>
	// InvalidStructLit occurs when a positional struct literal has an incorrect
	// number of values.
	//
	// Example:
	//  var _ = struct{i, j int}{1,2,3}
# <翻译结束>


<原文开始>
	// MissingLitField occurs when a struct literal refers to a field that does
	// not exist on the struct type.
	//
	// Example:
	//  var _ = struct{i int}{j: 2}
<原文结束>

# <翻译开始>
	// MissingLitField occurs when a struct literal refers to a field that does
	// not exist on the struct type.
	//
	// Example:
	//  var _ = struct{i int}{j: 2}
# <翻译结束>


<原文开始>
	// DuplicateLitField occurs when a struct literal contains duplicated
	// fields.
	//
	// Example:
	//  var _ = struct{i int}{i: 1, i: 2}
<原文结束>

# <翻译开始>
	// DuplicateLitField occurs when a struct literal contains duplicated
	// fields.
	//
	// Example:
	//  var _ = struct{i int}{i: 1, i: 2}
# <翻译结束>


<原文开始>
	// UnexportedLitField occurs when a positional struct literal implicitly
	// assigns an unexported field of an imported type.
<原文结束>

# <翻译开始>
	// UnexportedLitField occurs when a positional struct literal implicitly
	// assigns an unexported field of an imported type.
# <翻译结束>


<原文开始>
	// InvalidLitField occurs when a field name is not a valid identifier.
	//
	// Example:
	//  var _ = struct{i int}{1: 1}
<原文结束>

# <翻译开始>
	// InvalidLitField occurs when a field name is not a valid identifier.
	//
	// Example:
	//  var _ = struct{i int}{1: 1}
# <翻译结束>


<原文开始>
	// UntypedLit occurs when a composite literal omits a required type
	// identifier.
	//
	// Example:
	//  type outer struct{
	//  	inner struct { i int }
	//  }
	//
	//  var _ = outer{inner: {1}}
<原文结束>

# <翻译开始>
	// UntypedLit occurs when a composite literal omits a required type
	// identifier.
	//
	// Example:
	//  type outer struct{
	//  	inner struct { i int }
	//  }
	//
	//  var _ = outer{inner: {1}}
# <翻译结束>


<原文开始>
	// InvalidLit occurs when a composite literal expression does not match its
	// type.
	//
	// Example:
	//  type P *struct{
	//  	x int
	//  }
	//  var _ = P {}
<原文结束>

# <翻译开始>
	// InvalidLit occurs when a composite literal expression does not match its
	// type.
	//
	// Example:
	//  type P *struct{
	//  	x int
	//  }
	//  var _ = P {}
# <翻译结束>


<原文开始>
	// AmbiguousSelector occurs when a selector is ambiguous.
	//
	// Example:
	//  type E1 struct { i int }
	//  type E2 struct { i int }
	//  type T struct { E1; E2 }
	//
	//  var x T
	//  var _ = x.i
<原文结束>

# <翻译开始>
	// AmbiguousSelector occurs when a selector is ambiguous.
	//
	// Example:
	//  type E1 struct { i int }
	//  type E2 struct { i int }
	//  type T struct { E1; E2 }
	//
	//  var x T
	//  var _ = x.i
# <翻译结束>


<原文开始>
	// UndeclaredImportedName occurs when a package-qualified identifier is
	// undeclared by the imported package.
	//
	// Example:
	//  import "go/types"
	//
	//  var _ = types.NotAnActualIdentifier
<原文结束>

# <翻译开始>
	// UndeclaredImportedName occurs when a package-qualified identifier is
	// undeclared by the imported package.
	//
	// Example:
	//  import "go/types"
	//
	//  var _ = types.NotAnActualIdentifier
# <翻译结束>


<原文开始>
	// UnexportedName occurs when a selector refers to an unexported identifier
	// of an imported package.
	//
	// Example:
	//  import "reflect"
	//
	//  type _ reflect.flag
<原文结束>

# <翻译开始>
	// UnexportedName occurs when a selector refers to an unexported identifier
	// of an imported package.
	//
	// Example:
	//  import "reflect"
	//
	//  type _ reflect.flag
# <翻译结束>


<原文开始>
	// UndeclaredName occurs when an identifier is not declared in the current
	// scope.
	//
	// Example:
	//  var x T
<原文结束>

# <翻译开始>
	// UndeclaredName occurs when an identifier is not declared in the current
	// scope.
	//
	// Example:
	//  var x T
# <翻译结束>


<原文开始>
	// MissingFieldOrMethod occurs when a selector references a field or method
	// that does not exist.
	//
	// Example:
	//  type T struct {}
	//
	//  var x = T{}.f
<原文结束>

# <翻译开始>
	// MissingFieldOrMethod occurs when a selector references a field or method
	// that does not exist.
	//
	// Example:
	//  type T struct {}
	//
	//  var x = T{}.f
# <翻译结束>


<原文开始>
	// BadDotDotDotSyntax occurs when a "..." occurs in a context where it is
	// not valid.
	//
	// Example:
	//  var _ = map[int][...]int{0: {}}
<原文结束>

# <翻译开始>
	// BadDotDotDotSyntax occurs when a "..." occurs in a context where it is
	// not valid.
	//
	// Example:
	//  var _ = map[int][...]int{0: {}}
# <翻译结束>


<原文开始>
	// NonVariadicDotDotDot occurs when a "..." is used on the final argument to
	// a non-variadic function.
	//
	// Example:
	//  func printArgs(s []string) {
	//  	for _, a := range s {
	//  		println(a)
	//  	}
	//  }
	//
	//  func f() {
	//  	s := []string{"a", "b", "c"}
	//  	printArgs(s...)
	//  }
<原文结束>

# <翻译开始>
	// NonVariadicDotDotDot occurs when a "..." is used on the final argument to
	// a non-variadic function.
	//
	// Example:
	//  func printArgs(s []string) {
	//  	for _, a := range s {
	//  		println(a)
	//  	}
	//  }
	//
	//  func f() {
	//  	s := []string{"a", "b", "c"}
	//  	printArgs(s...)
	//  }
# <翻译结束>


<原文开始>
	// MisplacedDotDotDot occurs when a "..." is used somewhere other than the
	// final argument in a function declaration.
	//
	// Example:
	// 	func f(...int, int)
<原文结束>

# <翻译开始>
	// MisplacedDotDotDot occurs when a "..." is used somewhere other than the
	// final argument in a function declaration.
	//
	// Example:
	// 	func f(...int, int)
# <翻译结束>


<原文开始>
// InvalidDotDotDotOperand was removed.
<原文结束>

# <翻译开始>
// InvalidDotDotDotOperand was removed.
# <翻译结束>


<原文开始>
	// InvalidDotDotDot occurs when a "..." is used in a non-variadic built-in
	// function.
	//
	// Example:
	//  var s = []int{1, 2, 3}
	//  var l = len(s...)
<原文结束>

# <翻译开始>
	// InvalidDotDotDot occurs when a "..." is used in a non-variadic built-in
	// function.
	//
	// Example:
	//  var s = []int{1, 2, 3}
	//  var l = len(s...)
# <翻译结束>


<原文开始>
	// UncalledBuiltin occurs when a built-in function is used as a
	// function-valued expression, instead of being called.
	//
	// Per the spec:
	//  "The built-in functions do not have standard Go types, so they can only
	//  appear in call expressions; they cannot be used as function values."
	//
	// Example:
	//  var _ = copy
<原文结束>

# <翻译开始>
	// UncalledBuiltin occurs when a built-in function is used as a
	// function-valued expression, instead of being called.
	//
	// Per the spec:
	//  "The built-in functions do not have standard Go types, so they can only
	//  appear in call expressions; they cannot be used as function values."
	//
	// Example:
	//  var _ = copy
# <翻译结束>


<原文开始>
	// InvalidAppend occurs when append is called with a first argument that is
	// not a slice.
	//
	// Example:
	//  var _ = append(1, 2)
<原文结束>

# <翻译开始>
	// InvalidAppend occurs when append is called with a first argument that is
	// not a slice.
	//
	// Example:
	//  var _ = append(1, 2)
# <翻译结束>


<原文开始>
	// InvalidCap occurs when an argument to the cap built-in function is not of
	// supported type.
	//
	// See https://golang.org/ref/spec#Length_and_capacity for information on
	// which underlying types are supported as arguments to cap and len.
	//
	// Example:
	//  var s = 2
	//  var x = cap(s)
<原文结束>

# <翻译开始>
	// InvalidCap occurs when an argument to the cap built-in function is not of
	// supported type.
	//
	// See https://golang.org/ref/spec#Length_and_capacity for information on
	// which underlying types are supported as arguments to cap and len.
	//
	// Example:
	//  var s = 2
	//  var x = cap(s)
# <翻译结束>


<原文开始>
	// InvalidClose occurs when close(...) is called with an argument that is
	// not of channel type, or that is a receive-only channel.
	//
	// Example:
	//  func f() {
	//  	var x int
	//  	close(x)
	//  }
<原文结束>

# <翻译开始>
	// InvalidClose occurs when close(...) is called with an argument that is
	// not of channel type, or that is a receive-only channel.
	//
	// Example:
	//  func f() {
	//  	var x int
	//  	close(x)
	//  }
# <翻译结束>


<原文开始>
	// InvalidCopy occurs when the arguments are not of slice type or do not
	// have compatible type.
	//
	// See https://golang.org/ref/spec#Appending_and_copying_slices for more
	// information on the type requirements for the copy built-in.
	//
	// Example:
	//  func f() {
	//  	var x []int
	//  	y := []int64{1,2,3}
	//  	copy(x, y)
	//  }
<原文结束>

# <翻译开始>
	// InvalidCopy occurs when the arguments are not of slice type or do not
	// have compatible type.
	//
	// See https://golang.org/ref/spec#Appending_and_copying_slices for more
	// information on the type requirements for the copy built-in.
	//
	// Example:
	//  func f() {
	//  	var x []int
	//  	y := []int64{1,2,3}
	//  	copy(x, y)
	//  }
# <翻译结束>


<原文开始>
	// InvalidComplex occurs when the complex built-in function is called with
	// arguments with incompatible types.
	//
	// Example:
	//  var _ = complex(float32(1), float64(2))
<原文结束>

# <翻译开始>
	// InvalidComplex occurs when the complex built-in function is called with
	// arguments with incompatible types.
	//
	// Example:
	//  var _ = complex(float32(1), float64(2))
# <翻译结束>


<原文开始>
	// InvalidDelete occurs when the delete built-in function is called with a
	// first argument that is not a map.
	//
	// Example:
	//  func f() {
	//  	m := "hello"
	//  	delete(m, "e")
	//  }
<原文结束>

# <翻译开始>
	// InvalidDelete occurs when the delete built-in function is called with a
	// first argument that is not a map.
	//
	// Example:
	//  func f() {
	//  	m := "hello"
	//  	delete(m, "e")
	//  }
# <翻译结束>


<原文开始>
	// InvalidImag occurs when the imag built-in function is called with an
	// argument that does not have complex type.
	//
	// Example:
	//  var _ = imag(int(1))
<原文结束>

# <翻译开始>
	// InvalidImag occurs when the imag built-in function is called with an
	// argument that does not have complex type.
	//
	// Example:
	//  var _ = imag(int(1))
# <翻译结束>


<原文开始>
	// InvalidLen occurs when an argument to the len built-in function is not of
	// supported type.
	//
	// See https://golang.org/ref/spec#Length_and_capacity for information on
	// which underlying types are supported as arguments to cap and len.
	//
	// Example:
	//  var s = 2
	//  var x = len(s)
<原文结束>

# <翻译开始>
	// InvalidLen occurs when an argument to the len built-in function is not of
	// supported type.
	//
	// See https://golang.org/ref/spec#Length_and_capacity for information on
	// which underlying types are supported as arguments to cap and len.
	//
	// Example:
	//  var s = 2
	//  var x = len(s)
# <翻译结束>


<原文开始>
	// SwappedMakeArgs occurs when make is called with three arguments, and its
	// length argument is larger than its capacity argument.
	//
	// Example:
	//  var x = make([]int, 3, 2)
<原文结束>

# <翻译开始>
	// SwappedMakeArgs occurs when make is called with three arguments, and its
	// length argument is larger than its capacity argument.
	//
	// Example:
	//  var x = make([]int, 3, 2)
# <翻译结束>


<原文开始>
	// InvalidMake occurs when make is called with an unsupported type argument.
	//
	// See https://golang.org/ref/spec#Making_slices_maps_and_channels for
	// information on the types that may be created using make.
	//
	// Example:
	//  var x = make(int)
<原文结束>

# <翻译开始>
	// InvalidMake occurs when make is called with an unsupported type argument.
	//
	// See https://golang.org/ref/spec#Making_slices_maps_and_channels for
	// information on the types that may be created using make.
	//
	// Example:
	//  var x = make(int)
# <翻译结束>


<原文开始>
	// InvalidReal occurs when the real built-in function is called with an
	// argument that does not have complex type.
	//
	// Example:
	//  var _ = real(int(1))
<原文结束>

# <翻译开始>
	// InvalidReal occurs when the real built-in function is called with an
	// argument that does not have complex type.
	//
	// Example:
	//  var _ = real(int(1))
# <翻译结束>


<原文开始>
	// InvalidAssert occurs when a type assertion is applied to a
	// value that is not of interface type.
	//
	// Example:
	//  var x = 1
	//  var _ = x.(float64)
<原文结束>

# <翻译开始>
	// InvalidAssert occurs when a type assertion is applied to a
	// value that is not of interface type.
	//
	// Example:
	//  var x = 1
	//  var _ = x.(float64)
# <翻译结束>


<原文开始>
	// ImpossibleAssert occurs for a type assertion x.(T) when the value x of
	// interface cannot have dynamic type T, due to a missing or mismatching
	// method on T.
	//
	// Example:
	//  type T int
	//
	//  func (t *T) m() int { return int(*t) }
	//
	//  type I interface { m() int }
	//
	//  var x I
	//  var _ = x.(T)
<原文结束>

# <翻译开始>
	// ImpossibleAssert occurs for a type assertion x.(T) when the value x of
	// interface cannot have dynamic type T, due to a missing or mismatching
	// method on T.
	//
	// Example:
	//  type T int
	//
	//  func (t *T) m() int { return int(*t) }
	//
	//  type I interface { m() int }
	//
	//  var x I
	//  var _ = x.(T)
# <翻译结束>


<原文开始>
	// InvalidConversion occurs when the argument type cannot be converted to the
	// target.
	//
	// See https://golang.org/ref/spec#Conversions for the rules of
	// convertibility.
	//
	// Example:
	//  var x float64
	//  var _ = string(x)
<原文结束>

# <翻译开始>
	// InvalidConversion occurs when the argument type cannot be converted to the
	// target.
	//
	// See https://golang.org/ref/spec#Conversions for the rules of
	// convertibility.
	//
	// Example:
	//  var x float64
	//  var _ = string(x)
# <翻译结束>


<原文开始>
	// InvalidUntypedConversion occurs when there is no valid implicit
	// conversion from an untyped value satisfying the type constraints of the
	// context in which it is used.
	//
	// Example:
	//  var _ = 1 + []int{}
<原文结束>

# <翻译开始>
	// InvalidUntypedConversion occurs when there is no valid implicit
	// conversion from an untyped value satisfying the type constraints of the
	// context in which it is used.
	//
	// Example:
	//  var _ = 1 + []int{}
# <翻译结束>


<原文开始>
	// BadOffsetofSyntax occurs when unsafe.Offsetof is called with an argument
	// that is not a selector expression.
	//
	// Example:
	//  import "unsafe"
	//
	//  var x int
	//  var _ = unsafe.Offsetof(x)
<原文结束>

# <翻译开始>
	// BadOffsetofSyntax occurs when unsafe.Offsetof is called with an argument
	// that is not a selector expression.
	//
	// Example:
	//  import "unsafe"
	//
	//  var x int
	//  var _ = unsafe.Offsetof(x)
# <翻译结束>


<原文开始>
	// InvalidOffsetof occurs when unsafe.Offsetof is called with a method
	// selector, rather than a field selector, or when the field is embedded via
	// a pointer.
	//
	// Per the spec:
	//
	//  "If f is an embedded field, it must be reachable without pointer
	//  indirections through fields of the struct. "
	//
	// Example:
	//  import "unsafe"
	//
	//  type T struct { f int }
	//  type S struct { *T }
	//  var s S
	//  var _ = unsafe.Offsetof(s.f)
	//
	// Example:
	//  import "unsafe"
	//
	//  type S struct{}
	//
	//  func (S) m() {}
	//
	//  var s S
	//  var _ = unsafe.Offsetof(s.m)
<原文结束>

# <翻译开始>
	// InvalidOffsetof occurs when unsafe.Offsetof is called with a method
	// selector, rather than a field selector, or when the field is embedded via
	// a pointer.
	//
	// Per the spec:
	//
	//  "If f is an embedded field, it must be reachable without pointer
	//  indirections through fields of the struct. "
	//
	// Example:
	//  import "unsafe"
	//
	//  type T struct { f int }
	//  type S struct { *T }
	//  var s S
	//  var _ = unsafe.Offsetof(s.f)
	//
	// Example:
	//  import "unsafe"
	//
	//  type S struct{}
	//
	//  func (S) m() {}
	//
	//  var s S
	//  var _ = unsafe.Offsetof(s.m)
# <翻译结束>


<原文开始>
	// UnusedExpr occurs when a side-effect free expression is used as a
	// statement. Such a statement has no effect.
	//
	// Example:
	//  func f(i int) {
	//  	i*i
	//  }
<原文结束>

# <翻译开始>
	// UnusedExpr occurs when a side-effect free expression is used as a
	// statement. Such a statement has no effect.
	//
	// Example:
	//  func f(i int) {
	//  	i*i
	//  }
# <翻译结束>


<原文开始>
	// UnusedVar occurs when a variable is declared but unused.
	//
	// Example:
	//  func f() {
	//  	x := 1
	//  }
<原文结束>

# <翻译开始>
	// UnusedVar occurs when a variable is declared but unused.
	//
	// Example:
	//  func f() {
	//  	x := 1
	//  }
# <翻译结束>


<原文开始>
	// MissingReturn occurs when a function with results is missing a return
	// statement.
	//
	// Example:
	//  func f() int {}
<原文结束>

# <翻译开始>
	// MissingReturn occurs when a function with results is missing a return
	// statement.
	//
	// Example:
	//  func f() int {}
# <翻译结束>


<原文开始>
	// WrongResultCount occurs when a return statement returns an incorrect
	// number of values.
	//
	// Example:
	//  func ReturnOne() int {
	//  	return 1, 2
	//  }
<原文结束>

# <翻译开始>
	// WrongResultCount occurs when a return statement returns an incorrect
	// number of values.
	//
	// Example:
	//  func ReturnOne() int {
	//  	return 1, 2
	//  }
# <翻译结束>


<原文开始>
	// OutOfScopeResult occurs when the name of a value implicitly returned by
	// an empty return statement is shadowed in a nested scope.
	//
	// Example:
	//  func factor(n int) (i int) {
	//  	for i := 2; i < n; i++ {
	//  		if n%i == 0 {
	//  			return
	//  		}
	//  	}
	//  	return 0
	//  }
<原文结束>

# <翻译开始>
	// OutOfScopeResult occurs when the name of a value implicitly returned by
	// an empty return statement is shadowed in a nested scope.
	//
	// Example:
	//  func factor(n int) (i int) {
	//  	for i := 2; i < n; i++ {
	//  		if n%i == 0 {
	//  			return
	//  		}
	//  	}
	//  	return 0
	//  }
# <翻译结束>


<原文开始>
	// InvalidCond occurs when an if condition is not a boolean expression.
	//
	// Example:
	//  func checkReturn(i int) {
	//  	if i {
	//  		panic("non-zero return")
	//  	}
	//  }
<原文结束>

# <翻译开始>
	// InvalidCond occurs when an if condition is not a boolean expression.
	//
	// Example:
	//  func checkReturn(i int) {
	//  	if i {
	//  		panic("non-zero return")
	//  	}
	//  }
# <翻译结束>


<原文开始>
	// InvalidPostDecl occurs when there is a declaration in a for-loop post
	// statement.
	//
	// Example:
	//  func f() {
	//  	for i := 0; i < 10; j := 0 {}
	//  }
<原文结束>

# <翻译开始>
	// InvalidPostDecl occurs when there is a declaration in a for-loop post
	// statement.
	//
	// Example:
	//  func f() {
	//  	for i := 0; i < 10; j := 0 {}
	//  }
# <翻译结束>


<原文开始>
// InvalidChanRange was removed.
<原文结束>

# <翻译开始>
// InvalidChanRange was removed.
# <翻译结束>


<原文开始>
	// InvalidIterVar occurs when two iteration variables are used while ranging
	// over a channel.
	//
	// Example:
	//  func f(c chan int) {
	//  	for k, v := range c {
	//  		println(k, v)
	//  	}
	//  }
<原文结束>

# <翻译开始>
	// InvalidIterVar occurs when two iteration variables are used while ranging
	// over a channel.
	//
	// Example:
	//  func f(c chan int) {
	//  	for k, v := range c {
	//  		println(k, v)
	//  	}
	//  }
# <翻译结束>


<原文开始>
	// InvalidRangeExpr occurs when the type of a range expression is not array,
	// slice, string, map, or channel.
	//
	// Example:
	//  func f(i int) {
	//  	for j := range i {
	//  		println(j)
	//  	}
	//  }
<原文结束>

# <翻译开始>
	// InvalidRangeExpr occurs when the type of a range expression is not array,
	// slice, string, map, or channel.
	//
	// Example:
	//  func f(i int) {
	//  	for j := range i {
	//  		println(j)
	//  	}
	//  }
# <翻译结束>


<原文开始>
	// MisplacedBreak occurs when a break statement is not within a for, switch,
	// or select statement of the innermost function definition.
	//
	// Example:
	//  func f() {
	//  	break
	//  }
<原文结束>

# <翻译开始>
	// MisplacedBreak occurs when a break statement is not within a for, switch,
	// or select statement of the innermost function definition.
	//
	// Example:
	//  func f() {
	//  	break
	//  }
# <翻译结束>


<原文开始>
	// MisplacedContinue occurs when a continue statement is not within a for
	// loop of the innermost function definition.
	//
	// Example:
	//  func sumeven(n int) int {
	//  	proceed := func() {
	//  		continue
	//  	}
	//  	sum := 0
	//  	for i := 1; i <= n; i++ {
	//  		if i % 2 != 0 {
	//  			proceed()
	//  		}
	//  		sum += i
	//  	}
	//  	return sum
	//  }
<原文结束>

# <翻译开始>
	// MisplacedContinue occurs when a continue statement is not within a for
	// loop of the innermost function definition.
	//
	// Example:
	//  func sumeven(n int) int {
	//  	proceed := func() {
	//  		continue
	//  	}
	//  	sum := 0
	//  	for i := 1; i <= n; i++ {
	//  		if i % 2 != 0 {
	//  			proceed()
	//  		}
	//  		sum += i
	//  	}
	//  	return sum
	//  }
# <翻译结束>


<原文开始>
	// MisplacedFallthrough occurs when a fallthrough statement is not within an
	// expression switch.
	//
	// Example:
	//  func typename(i interface{}) string {
	//  	switch i.(type) {
	//  	case int64:
	//  		fallthrough
	//  	case int:
	//  		return "int"
	//  	}
	//  	return "unsupported"
	//  }
<原文结束>

# <翻译开始>
	// MisplacedFallthrough occurs when a fallthrough statement is not within an
	// expression switch.
	//
	// Example:
	//  func typename(i interface{}) string {
	//  	switch i.(type) {
	//  	case int64:
	//  		fallthrough
	//  	case int:
	//  		return "int"
	//  	}
	//  	return "unsupported"
	//  }
# <翻译结束>


<原文开始>
	// DuplicateCase occurs when a type or expression switch has duplicate
	// cases.
	//
	// Example:
	//  func printInt(i int) {
	//  	switch i {
	//  	case 1:
	//  		println("one")
	//  	case 1:
	//  		println("One")
	//  	}
	//  }
<原文结束>

# <翻译开始>
	// DuplicateCase occurs when a type or expression switch has duplicate
	// cases.
	//
	// Example:
	//  func printInt(i int) {
	//  	switch i {
	//  	case 1:
	//  		println("one")
	//  	case 1:
	//  		println("One")
	//  	}
	//  }
# <翻译结束>


<原文开始>
	// DuplicateDefault occurs when a type or expression switch has multiple
	// default clauses.
	//
	// Example:
	//  func printInt(i int) {
	//  	switch i {
	//  	case 1:
	//  		println("one")
	//  	default:
	//  		println("One")
	//  	default:
	//  		println("1")
	//  	}
	//  }
<原文结束>

# <翻译开始>
	// DuplicateDefault occurs when a type or expression switch has multiple
	// default clauses.
	//
	// Example:
	//  func printInt(i int) {
	//  	switch i {
	//  	case 1:
	//  		println("one")
	//  	default:
	//  		println("One")
	//  	default:
	//  		println("1")
	//  	}
	//  }
# <翻译结束>


<原文开始>
	// BadTypeKeyword occurs when a .(type) expression is used anywhere other
	// than a type switch.
	//
	// Example:
	//  type I interface {
	//  	m()
	//  }
	//  var t I
	//  var _ = t.(type)
<原文结束>

# <翻译开始>
	// BadTypeKeyword occurs when a .(type) expression is used anywhere other
	// than a type switch.
	//
	// Example:
	//  type I interface {
	//  	m()
	//  }
	//  var t I
	//  var _ = t.(type)
# <翻译结束>


<原文开始>
	// InvalidTypeSwitch occurs when .(type) is used on an expression that is
	// not of interface type.
	//
	// Example:
	//  func f(i int) {
	//  	switch x := i.(type) {}
	//  }
<原文结束>

# <翻译开始>
	// InvalidTypeSwitch occurs when .(type) is used on an expression that is
	// not of interface type.
	//
	// Example:
	//  func f(i int) {
	//  	switch x := i.(type) {}
	//  }
# <翻译结束>


<原文开始>
	// InvalidExprSwitch occurs when a switch expression is not comparable.
	//
	// Example:
	//  func _() {
	//  	var a struct{ _ func() }
	//  	switch a /* ERROR cannot switch on a */ {
	//  	}
	//  }
<原文结束>

# <翻译开始>
	// InvalidExprSwitch occurs when a switch expression is not comparable.
	//
	// Example:
	//  func _() {
	//  	var a struct{ _ func() }
	//  	switch a /* ERROR cannot switch on a */ {
	//  	}
	//  }
# <翻译结束>


<原文开始>
	// InvalidSelectCase occurs when a select case is not a channel send or
	// receive.
	//
	// Example:
	//  func checkChan(c <-chan int) bool {
	//  	select {
	//  	case c:
	//  		return true
	//  	default:
	//  		return false
	//  	}
	//  }
<原文结束>

# <翻译开始>
	// InvalidSelectCase occurs when a select case is not a channel send or
	// receive.
	//
	// Example:
	//  func checkChan(c <-chan int) bool {
	//  	select {
	//  	case c:
	//  		return true
	//  	default:
	//  		return false
	//  	}
	//  }
# <翻译结束>


<原文开始>
	// UndeclaredLabel occurs when an undeclared label is jumped to.
	//
	// Example:
	//  func f() {
	//  	goto L
	//  }
<原文结束>

# <翻译开始>
	// UndeclaredLabel occurs when an undeclared label is jumped to.
	//
	// Example:
	//  func f() {
	//  	goto L
	//  }
# <翻译结束>


<原文开始>
	// DuplicateLabel occurs when a label is declared more than once.
	//
	// Example:
	//  func f() int {
	//  L:
	//  L:
	//  	return 1
	//  }
<原文结束>

# <翻译开始>
	// DuplicateLabel occurs when a label is declared more than once.
	//
	// Example:
	//  func f() int {
	//  L:
	//  L:
	//  	return 1
	//  }
# <翻译结束>


<原文开始>
	// MisplacedLabel occurs when a break or continue label is not on a for,
	// switch, or select statement.
	//
	// Example:
	//  func f() {
	//  L:
	//  	a := []int{1,2,3}
	//  	for _, e := range a {
	//  		if e > 10 {
	//  			break L
	//  		}
	//  		println(a)
	//  	}
	//  }
<原文结束>

# <翻译开始>
	// MisplacedLabel occurs when a break or continue label is not on a for,
	// switch, or select statement.
	//
	// Example:
	//  func f() {
	//  L:
	//  	a := []int{1,2,3}
	//  	for _, e := range a {
	//  		if e > 10 {
	//  			break L
	//  		}
	//  		println(a)
	//  	}
	//  }
# <翻译结束>


<原文开始>
	// UnusedLabel occurs when a label is declared and not used.
	//
	// Example:
	//  func f() {
	//  L:
	//  }
<原文结束>

# <翻译开始>
	// UnusedLabel occurs when a label is declared and not used.
	//
	// Example:
	//  func f() {
	//  L:
	//  }
# <翻译结束>


<原文开始>
	// JumpOverDecl occurs when a label jumps over a variable declaration.
	//
	// Example:
	//  func f() int {
	//  	goto L
	//  	x := 2
	//  L:
	//  	x++
	//  	return x
	//  }
<原文结束>

# <翻译开始>
	// JumpOverDecl occurs when a label jumps over a variable declaration.
	//
	// Example:
	//  func f() int {
	//  	goto L
	//  	x := 2
	//  L:
	//  	x++
	//  	return x
	//  }
# <翻译结束>


<原文开始>
	// JumpIntoBlock occurs when a forward jump goes to a label inside a nested
	// block.
	//
	// Example:
	//  func f(x int) {
	//  	goto L
	//  	if x > 0 {
	//  	L:
	//  		print("inside block")
	//  	}
	// }
<原文结束>

# <翻译开始>
	// JumpIntoBlock occurs when a forward jump goes to a label inside a nested
	// block.
	//
	// Example:
	//  func f(x int) {
	//  	goto L
	//  	if x > 0 {
	//  	L:
	//  		print("inside block")
	//  	}
	// }
# <翻译结束>


<原文开始>
	// InvalidMethodExpr occurs when a pointer method is called but the argument
	// is not addressable.
	//
	// Example:
	//  type T struct {}
	//
	//  func (*T) m() int { return 1 }
	//
	//  var _ = T.m(T{})
<原文结束>

# <翻译开始>
	// InvalidMethodExpr occurs when a pointer method is called but the argument
	// is not addressable.
	//
	// Example:
	//  type T struct {}
	//
	//  func (*T) m() int { return 1 }
	//
	//  var _ = T.m(T{})
# <翻译结束>


<原文开始>
	// WrongArgCount occurs when too few or too many arguments are passed by a
	// function call.
	//
	// Example:
	//  func f(i int) {}
	//  var x = f()
<原文结束>

# <翻译开始>
	// WrongArgCount occurs when too few or too many arguments are passed by a
	// function call.
	//
	// Example:
	//  func f(i int) {}
	//  var x = f()
# <翻译结束>


<原文开始>
	// InvalidCall occurs when an expression is called that is not of function
	// type.
	//
	// Example:
	//  var x = "x"
	//  var y = x()
<原文结束>

# <翻译开始>
	// InvalidCall occurs when an expression is called that is not of function
	// type.
	//
	// Example:
	//  var x = "x"
	//  var y = x()
# <翻译结束>


<原文开始>
	// UnusedResults occurs when a restricted expression-only built-in function
	// is suspended via go or defer. Such a suspension discards the results of
	// these side-effect free built-in functions, and therefore is ineffectual.
	//
	// Example:
	//  func f(a []int) int {
	//  	defer len(a)
	//  	return i
	//  }
<原文结束>

# <翻译开始>
	// UnusedResults occurs when a restricted expression-only built-in function
	// is suspended via go or defer. Such a suspension discards the results of
	// these side-effect free built-in functions, and therefore is ineffectual.
	//
	// Example:
	//  func f(a []int) int {
	//  	defer len(a)
	//  	return i
	//  }
# <翻译结束>


<原文开始>
	// InvalidDefer occurs when a deferred expression is not a function call,
	// for example if the expression is a type conversion.
	//
	// Example:
	//  func f(i int) int {
	//  	defer int32(i)
	//  	return i
	//  }
<原文结束>

# <翻译开始>
	// InvalidDefer occurs when a deferred expression is not a function call,
	// for example if the expression is a type conversion.
	//
	// Example:
	//  func f(i int) int {
	//  	defer int32(i)
	//  	return i
	//  }
# <翻译结束>


<原文开始>
	// InvalidGo occurs when a go expression is not a function call, for example
	// if the expression is a type conversion.
	//
	// Example:
	//  func f(i int) int {
	//  	go int32(i)
	//  	return i
	//  }
<原文结束>

# <翻译开始>
	// InvalidGo occurs when a go expression is not a function call, for example
	// if the expression is a type conversion.
	//
	// Example:
	//  func f(i int) int {
	//  	go int32(i)
	//  	return i
	//  }
# <翻译结束>


<原文开始>
// All codes below were added in Go 1.17.
<原文结束>

# <翻译开始>
// All codes below were added in Go 1.17.
# <翻译结束>


<原文开始>
// BadDecl occurs when a declaration has invalid syntax.
<原文结束>

# <翻译开始>
// BadDecl occurs when a declaration has invalid syntax.
# <翻译结束>


<原文开始>
	// RepeatedDecl occurs when an identifier occurs more than once on the left
	// hand side of a short variable declaration.
	//
	// Example:
	//  func _() {
	//  	x, y, y := 1, 2, 3
	//  }
<原文结束>

# <翻译开始>
	// RepeatedDecl occurs when an identifier occurs more than once on the left
	// hand side of a short variable declaration.
	//
	// Example:
	//  func _() {
	//  	x, y, y := 1, 2, 3
	//  }
# <翻译结束>


<原文开始>
	// InvalidUnsafeAdd occurs when unsafe.Add is called with a
	// length argument that is not of integer type.
	// It also occurs if it is used in a package compiled for a
	// language version before go1.17.
	//
	// Example:
	//  import "unsafe"
	//
	//  var p unsafe.Pointer
	//  var _ = unsafe.Add(p, float64(1))
<原文结束>

# <翻译开始>
	// InvalidUnsafeAdd occurs when unsafe.Add is called with a
	// length argument that is not of integer type.
	// It also occurs if it is used in a package compiled for a
	// language version before go1.17.
	//
	// Example:
	//  import "unsafe"
	//
	//  var p unsafe.Pointer
	//  var _ = unsafe.Add(p, float64(1))
# <翻译结束>


<原文开始>
	// InvalidUnsafeSlice occurs when unsafe.Slice is called with a
	// pointer argument that is not of pointer type or a length argument
	// that is not of integer type, negative, or out of bounds.
	// It also occurs if it is used in a package compiled for a language
	// version before go1.17.
	//
	// Example:
	//  import "unsafe"
	//
	//  var x int
	//  var _ = unsafe.Slice(x, 1)
	//
	// Example:
	//  import "unsafe"
	//
	//  var x int
	//  var _ = unsafe.Slice(&x, float64(1))
	//
	// Example:
	//  import "unsafe"
	//
	//  var x int
	//  var _ = unsafe.Slice(&x, -1)
	//
	// Example:
	//  import "unsafe"
	//
	//  var x int
	//  var _ = unsafe.Slice(&x, uint64(1) << 63)
<原文结束>

# <翻译开始>
	// InvalidUnsafeSlice occurs when unsafe.Slice is called with a
	// pointer argument that is not of pointer type or a length argument
	// that is not of integer type, negative, or out of bounds.
	// It also occurs if it is used in a package compiled for a language
	// version before go1.17.
	//
	// Example:
	//  import "unsafe"
	//
	//  var x int
	//  var _ = unsafe.Slice(x, 1)
	//
	// Example:
	//  import "unsafe"
	//
	//  var x int
	//  var _ = unsafe.Slice(&x, float64(1))
	//
	// Example:
	//  import "unsafe"
	//
	//  var x int
	//  var _ = unsafe.Slice(&x, -1)
	//
	// Example:
	//  import "unsafe"
	//
	//  var x int
	//  var _ = unsafe.Slice(&x, uint64(1) << 63)
# <翻译结束>


<原文开始>
// All codes below were added in Go 1.18.
<原文结束>

# <翻译开始>
// All codes below were added in Go 1.18.
# <翻译结束>


<原文开始>
	// UnsupportedFeature occurs when a language feature is used that is not
	// supported at this Go version.
<原文结束>

# <翻译开始>
	// UnsupportedFeature occurs when a language feature is used that is not
	// supported at this Go version.
# <翻译结束>


<原文开始>
	// NotAGenericType occurs when a non-generic type is used where a generic
	// type is expected: in type or function instantiation.
	//
	// Example:
	//  type T int
	//
	//  var _ T[int]
<原文结束>

# <翻译开始>
	// NotAGenericType occurs when a non-generic type is used where a generic
	// type is expected: in type or function instantiation.
	//
	// Example:
	//  type T int
	//
	//  var _ T[int]
# <翻译结束>


<原文开始>
	// WrongTypeArgCount occurs when a type or function is instantiated with an
	// incorrent number of type arguments, including when a generic type or
	// function is used without instantiation.
	//
	// Errors inolving failed type inference are assigned other error codes.
	//
	// Example:
	//  type T[p any] int
	//
	//  var _ T[int, string]
	//
	// Example:
	//  func f[T any]() {}
	//
	//  var x = f
<原文结束>

# <翻译开始>
	// WrongTypeArgCount occurs when a type or function is instantiated with an
	// incorrent number of type arguments, including when a generic type or
	// function is used without instantiation.
	//
	// Errors inolving failed type inference are assigned other error codes.
	//
	// Example:
	//  type T[p any] int
	//
	//  var _ T[int, string]
	//
	// Example:
	//  func f[T any]() {}
	//
	//  var x = f
# <翻译结束>


<原文开始>
	// CannotInferTypeArgs occurs when type or function type argument inference
	// fails to infer all type arguments.
	//
	// Example:
	//  func f[T any]() {}
	//
	//  func _() {
	//  	f()
	//  }
<原文结束>

# <翻译开始>
	// CannotInferTypeArgs occurs when type or function type argument inference
	// fails to infer all type arguments.
	//
	// Example:
	//  func f[T any]() {}
	//
	//  func _() {
	//  	f()
	//  }
# <翻译结束>


<原文开始>
	// InvalidTypeArg occurs when a type argument does not satisfy its
	// corresponding type parameter constraints.
	//
	// Example:
	//  type T[P ~int] struct{}
	//
	//  var _ T[string]
<原文结束>

# <翻译开始>
	// InvalidTypeArg occurs when a type argument does not satisfy its
	// corresponding type parameter constraints.
	//
	// Example:
	//  type T[P ~int] struct{}
	//
	//  var _ T[string]
# <翻译结束>


<原文开始>
// arguments? InferenceFailed
<原文结束>

# <翻译开始>
// arguments? InferenceFailed
# <翻译结束>


<原文开始>
	// InvalidInstanceCycle occurs when an invalid cycle is detected
	// within the instantiation graph.
	//
	// Example:
	//  func f[T any]() { f[*T]() }
<原文结束>

# <翻译开始>
	// InvalidInstanceCycle occurs when an invalid cycle is detected
	// within the instantiation graph.
	//
	// Example:
	//  func f[T any]() { f[*T]() }
# <翻译结束>


<原文开始>
	// InvalidUnion occurs when an embedded union or approximation element is
	// not valid.
	//
	// Example:
	//  type _ interface {
	//   	~int | interface{ m() }
	//  }
<原文结束>

# <翻译开始>
	// InvalidUnion occurs when an embedded union or approximation element is
	// not valid.
	//
	// Example:
	//  type _ interface {
	//   	~int | interface{ m() }
	//  }
# <翻译结束>


<原文开始>
	// MisplacedConstraintIface occurs when a constraint-type interface is used
	// outside of constraint position.
	//
	// Example:
	//   type I interface { ~int }
	//
	//   var _ I
<原文结束>

# <翻译开始>
	// MisplacedConstraintIface occurs when a constraint-type interface is used
	// outside of constraint position.
	//
	// Example:
	//   type I interface { ~int }
	//
	//   var _ I
# <翻译结束>


<原文开始>
	// InvalidMethodTypeParams occurs when methods have type parameters.
	//
	// It cannot be encountered with an AST parsed using go/parser.
<原文结束>

# <翻译开始>
	// InvalidMethodTypeParams occurs when methods have type parameters.
	//
	// It cannot be encountered with an AST parsed using go/parser.
# <翻译结束>


<原文开始>
	// MisplacedTypeParam occurs when a type parameter is used in a place where
	// it is not permitted.
	//
	// Example:
	//  type T[P any] P
	//
	// Example:
	//  type T[P any] struct{ *P }
<原文结束>

# <翻译开始>
	// MisplacedTypeParam occurs when a type parameter is used in a place where
	// it is not permitted.
	//
	// Example:
	//  type T[P any] P
	//
	// Example:
	//  type T[P any] struct{ *P }
# <翻译结束>


<原文开始>
	// InvalidUnsafeSliceData occurs when unsafe.SliceData is called with
	// an argument that is not of slice type. It also occurs if it is used
	// in a package compiled for a language version before go1.20.
	//
	// Example:
	//  import "unsafe"
	//
	//  var x int
	//  var _ = unsafe.SliceData(x)
<原文结束>

# <翻译开始>
	// InvalidUnsafeSliceData occurs when unsafe.SliceData is called with
	// an argument that is not of slice type. It also occurs if it is used
	// in a package compiled for a language version before go1.20.
	//
	// Example:
	//  import "unsafe"
	//
	//  var x int
	//  var _ = unsafe.SliceData(x)
# <翻译结束>


<原文开始>
	// InvalidUnsafeString occurs when unsafe.String is called with
	// a length argument that is not of integer type, negative, or
	// out of bounds. It also occurs if it is used in a package
	// compiled for a language version before go1.20.
	//
	// Example:
	//  import "unsafe"
	//
	//  var b [10]byte
	//  var _ = unsafe.String(&b[0], -1)
<原文结束>

# <翻译开始>
	// InvalidUnsafeString occurs when unsafe.String is called with
	// a length argument that is not of integer type, negative, or
	// out of bounds. It also occurs if it is used in a package
	// compiled for a language version before go1.20.
	//
	// Example:
	//  import "unsafe"
	//
	//  var b [10]byte
	//  var _ = unsafe.String(&b[0], -1)
# <翻译结束>


<原文开始>
	// InvalidUnsafeStringData occurs if it is used in a package
	// compiled for a language version before go1.20.
<原文结束>

# <翻译开始>
	// InvalidUnsafeStringData occurs if it is used in a package
	// compiled for a language version before go1.20.
# <翻译结束>


<原文开始>
	// InvalidClear occurs when clear is called with an argument
	// that is not of map, slice, or pointer-to-array type.
	//
	// Example:
	//  func _(x int) {
	//  	clear(x)
	//  }
<原文结束>

# <翻译开始>
	// InvalidClear occurs when clear is called with an argument
	// that is not of map, slice, or pointer-to-array type.
	//
	// Example:
	//  func _(x int) {
	//  	clear(x)
	//  }
# <翻译结束>

