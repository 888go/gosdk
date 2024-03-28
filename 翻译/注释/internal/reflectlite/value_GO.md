
<原文开始>
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
# <翻译结束>


<原文开始>
// Value is the reflection interface to a Go value.
//
// Not all methods apply to all kinds of values. Restrictions,
// if any, are noted in the documentation for each method.
// Use the Kind method to find out the kind of value before
// calling kind-specific methods. Calling a method
// inappropriate to the kind of type causes a run time panic.
//
// The zero Value represents no value.
// Its IsValid method returns false, its Kind method returns Invalid,
// its String method returns "<invalid Value>", and all other methods panic.
// Most functions and methods never return an invalid value.
// If one does, its documentation states the conditions explicitly.
//
// A Value can be used concurrently by multiple goroutines provided that
// the underlying Go value can be used concurrently for the equivalent
// direct operations.
//
// To compare two Values, compare the results of the Interface method.
// Using == on two Values does not compare the underlying values
// they represent.
<原文结束>

# <翻译开始>
// Value is the reflection interface to a Go value.
//
// Not all methods apply to all kinds of values. Restrictions,
// if any, are noted in the documentation for each method.
// Use the Kind method to find out the kind of value before
// calling kind-specific methods. Calling a method
// inappropriate to the kind of type causes a run time panic.
//
// The zero Value represents no value.
// Its IsValid method returns false, its Kind method returns Invalid,
// its String method returns "<invalid Value>", and all other methods panic.
// Most functions and methods never return an invalid value.
// If one does, its documentation states the conditions explicitly.
//
// A Value can be used concurrently by multiple goroutines provided that
// the underlying Go value can be used concurrently for the equivalent
// direct operations.
//
// To compare two Values, compare the results of the Interface method.
// Using == on two Values does not compare the underlying values
// they represent.
# <翻译结束>


<原文开始>
// typ holds the type of the value represented by a Value.
<原文结束>

# <翻译开始>
// typ holds the type of the value represented by a Value.
# <翻译结束>


<原文开始>
	// Pointer-valued data or, if flagIndir is set, pointer to data.
	// Valid when either flagIndir is set or typ.pointers() is true.
<原文结束>

# <翻译开始>
	// Pointer-valued data or, if flagIndir is set, pointer to data.
	// Valid when either flagIndir is set or typ.pointers() is true.
# <翻译结束>


<原文开始>
	// flag holds metadata about the value.
	// The lowest bits are flag bits:
	//	- flagStickyRO: obtained via unexported not embedded field, so read-only
	//	- flagEmbedRO: obtained via unexported embedded field, so read-only
	//	- flagIndir: val holds a pointer to the data
	//	- flagAddr: v.CanAddr is true (implies flagIndir)
	// Value cannot represent method values.
	// The next five bits give the Kind of the value.
	// This repeats typ.Kind() except for method values.
	// The remaining 23+ bits give a method number for method values.
	// If flag.kind() != Func, code can assume that flagMethod is unset.
	// If ifaceIndir(typ), code can assume that flagIndir is set.
<原文结束>

# <翻译开始>
	// flag holds metadata about the value.
	// The lowest bits are flag bits:
	//	- flagStickyRO: obtained via unexported not embedded field, so read-only
	//	- flagEmbedRO: obtained via unexported embedded field, so read-only
	//	- flagIndir: val holds a pointer to the data
	//	- flagAddr: v.CanAddr is true (implies flagIndir)
	// Value cannot represent method values.
	// The next five bits give the Kind of the value.
	// This repeats typ.Kind() except for method values.
	// The remaining 23+ bits give a method number for method values.
	// If flag.kind() != Func, code can assume that flagMethod is unset.
	// If ifaceIndir(typ), code can assume that flagIndir is set.
# <翻译结束>


<原文开始>
	// A method value represents a curried method invocation
	// like r.Read for some receiver r. The typ+val+flag bits describe
	// the receiver r, but the flag's Kind bits say Func (methods are
	// functions), and the top bits of the flag give the method number
	// in r's type's method table.
<原文结束>

# <翻译开始>
	// A method value represents a curried method invocation
	// like r.Read for some receiver r. The typ+val+flag bits describe
	// the receiver r, but the flag's Kind bits say Func (methods are
	// functions), and the top bits of the flag give the method number
	// in r's type's method table.
# <翻译结束>


<原文开始>
// pointer returns the underlying pointer represented by v.
// v.Kind() must be Pointer, Map, Chan, Func, or UnsafePointer
<原文结束>

# <翻译开始>
// pointer returns the underlying pointer represented by v.
// v.Kind() must be Pointer, Map, Chan, Func, or UnsafePointer
# <翻译结束>


<原文开始>
// packEface converts v to the empty interface.
<原文结束>

# <翻译开始>
// packEface converts v to the empty interface.
# <翻译结束>


<原文开始>
// First, fill in the data portion of the interface.
<原文结束>

# <翻译开始>
// First, fill in the data portion of the interface.
# <翻译结束>


<原文开始>
// Value is indirect, and so is the interface we're making.
<原文结束>

# <翻译开始>
// Value is indirect, and so is the interface we're making.
# <翻译结束>


<原文开始>
			// TODO: pass safe boolean from valueInterface so
			// we don't need to copy if safe==true?
<原文结束>

# <翻译开始>
			// TODO: pass safe boolean from valueInterface so
			// we don't need to copy if safe==true?
# <翻译结束>


<原文开始>
		// Value is indirect, but interface is direct. We need
		// to load the data at v.ptr into the interface data word.
<原文结束>

# <翻译开始>
		// Value is indirect, but interface is direct. We need
		// to load the data at v.ptr into the interface data word.
# <翻译结束>


<原文开始>
// Value is direct, and so is the interface.
<原文结束>

# <翻译开始>
// Value is direct, and so is the interface.
# <翻译结束>


<原文开始>
	// Now, fill in the type portion. We're very careful here not
	// to have any operation between the e.word and e.typ assignments
	// that would let the garbage collector observe the partially-built
	// interface value.
<原文结束>

# <翻译开始>
	// Now, fill in the type portion. We're very careful here not
	// to have any operation between the e.word and e.typ assignments
	// that would let the garbage collector observe the partially-built
	// interface value.
# <翻译结束>


<原文开始>
// unpackEface converts the empty interface i to a Value.
<原文结束>

# <翻译开始>
// unpackEface converts the empty interface i to a Value.
# <翻译结束>


<原文开始>
// NOTE: don't read e.word until we know whether it is really a pointer or not.
<原文结束>

# <翻译开始>
// NOTE: don't read e.word until we know whether it is really a pointer or not.
# <翻译结束>


<原文开始>
// A ValueError occurs when a Value method is invoked on
// a Value that does not support it. Such cases are documented
// in the description of each method.
<原文结束>

# <翻译开始>
// A ValueError occurs when a Value method is invoked on
// a Value that does not support it. Such cases are documented
// in the description of each method.
# <翻译结束>


<原文开始>
// methodName returns the name of the calling method,
// assumed to be two stack frames above.
<原文结束>

# <翻译开始>
// methodName returns the name of the calling method,
// assumed to be two stack frames above.
# <翻译结束>


<原文开始>
// emptyInterface is the header for an interface{} value.
<原文结束>

# <翻译开始>
// emptyInterface is the header for an interface{} value.
# <翻译结束>


<原文开始>
// mustBeExported panics if f records that the value was obtained using
// an unexported field.
<原文结束>

# <翻译开始>
// mustBeExported panics if f records that the value was obtained using
// an unexported field.
# <翻译结束>


<原文开始>
// mustBeAssignable panics if f records that the value is not assignable,
// which is to say that either it was obtained using an unexported field
// or it is not addressable.
<原文结束>

# <翻译开始>
// mustBeAssignable panics if f records that the value is not assignable,
// which is to say that either it was obtained using an unexported field
// or it is not addressable.
# <翻译结束>


<原文开始>
// Assignable if addressable and not read-only.
<原文结束>

# <翻译开始>
// Assignable if addressable and not read-only.
# <翻译结束>


<原文开始>
// CanSet reports whether the value of v can be changed.
// A Value can be changed only if it is addressable and was not
// obtained by the use of unexported struct fields.
// If CanSet returns false, calling Set or any type-specific
// setter (e.g., SetBool, SetInt) will panic.
<原文结束>

# <翻译开始>
// CanSet reports whether the value of v can be changed.
// A Value can be changed only if it is addressable and was not
// obtained by the use of unexported struct fields.
// If CanSet returns false, calling Set or any type-specific
// setter (e.g., SetBool, SetInt) will panic.
# <翻译结束>


<原文开始>
// Elem returns the value that the interface v contains
// or that the pointer v points to.
// It panics if v's Kind is not Interface or Pointer.
// It returns the zero Value if v is nil.
<原文结束>

# <翻译开始>
// Elem returns the value that the interface v contains
// or that the pointer v points to.
// It panics if v's Kind is not Interface or Pointer.
// It returns the zero Value if v is nil.
# <翻译结束>


<原文开始>
// The returned value's address is v's value.
<原文结束>

# <翻译开始>
// The returned value's address is v's value.
# <翻译结束>


<原文开始>
		// Special case: return the element inside the interface.
		// Empty interface has one layout, all interfaces with
		// methods have a second layout.
<原文结束>

# <翻译开始>
		// Special case: return the element inside the interface.
		// Empty interface has one layout, all interfaces with
		// methods have a second layout.
# <翻译结束>


<原文开始>
// TODO: pass safe to packEface so we don't need to copy if safe==true?
<原文结束>

# <翻译开始>
// TODO: pass safe to packEface so we don't need to copy if safe==true?
# <翻译结束>


<原文开始>
// IsNil reports whether its argument v is nil. The argument must be
// a chan, func, interface, map, pointer, or slice value; if it is
// not, IsNil panics. Note that IsNil is not always equivalent to a
// regular comparison with nil in Go. For example, if v was created
// by calling ValueOf with an uninitialized interface variable i,
// i==nil will be true but v.IsNil will panic as v will be the zero
// Value.
<原文结束>

# <翻译开始>
// IsNil reports whether its argument v is nil. The argument must be
// a chan, func, interface, map, pointer, or slice value; if it is
// not, IsNil panics. Note that IsNil is not always equivalent to a
// regular comparison with nil in Go. For example, if v was created
// by calling ValueOf with an uninitialized interface variable i,
// i==nil will be true but v.IsNil will panic as v will be the zero
// Value.
# <翻译结束>


<原文开始>
		// if v.flag&flagMethod != 0 {
		// 	return false
		// }
<原文结束>

# <翻译开始>
		// if v.flag&flagMethod != 0 {
		// 	return false
		// }
# <翻译结束>


<原文开始>
		// Both interface and slice are nil if first word is 0.
		// Both are always bigger than a word; assume flagIndir.
<原文结束>

# <翻译开始>
		// Both interface and slice are nil if first word is 0.
		// Both are always bigger than a word; assume flagIndir.
# <翻译结束>


<原文开始>
// IsValid reports whether v represents a value.
// It returns false if v is the zero Value.
// If IsValid returns false, all other methods except String panic.
// Most functions and methods never return an invalid Value.
// If one does, its documentation states the conditions explicitly.
<原文结束>

# <翻译开始>
// IsValid reports whether v represents a value.
// It returns false if v is the zero Value.
// If IsValid returns false, all other methods except String panic.
// Most functions and methods never return an invalid Value.
// If one does, its documentation states the conditions explicitly.
# <翻译结束>


<原文开始>
// Kind returns v's Kind.
// If v is the zero Value (IsValid returns false), Kind returns Invalid.
<原文结束>

# <翻译开始>
// Kind returns v's Kind.
// If v is the zero Value (IsValid returns false), Kind returns Invalid.
# <翻译结束>


<原文开始>
// implemented in runtime:
<原文结束>

# <翻译开始>
// implemented in runtime:
# <翻译结束>


<原文开始>
// Len returns v's length.
// It panics if v's Kind is not Array, Chan, Map, Slice, or String.
<原文结束>

# <翻译开始>
// Len returns v's length.
// It panics if v's Kind is not Array, Chan, Map, Slice, or String.
# <翻译结束>


<原文开始>
// Slice is bigger than a word; assume flagIndir.
<原文结束>

# <翻译开始>
// Slice is bigger than a word; assume flagIndir.
# <翻译结束>


<原文开始>
// String is bigger than a word; assume flagIndir.
<原文结束>

# <翻译开始>
// String is bigger than a word; assume flagIndir.
# <翻译结束>


<原文开始>
// NumMethod returns the number of exported methods in the value's method set.
<原文结束>

# <翻译开始>
// NumMethod returns the number of exported methods in the value's method set.
# <翻译结束>


<原文开始>
// Set assigns x to the value v.
// It panics if CanSet returns false.
// As in Go, x's value must be assignable to v's type.
<原文结束>

# <翻译开始>
// Set assigns x to the value v.
// It panics if CanSet returns false.
// As in Go, x's value must be assignable to v's type.
# <翻译结束>


<原文开始>
// do not let unexported x leak
<原文结束>

# <翻译开始>
// do not let unexported x leak
# <翻译结束>


<原文开始>
// Method values not supported.
<原文结束>

# <翻译开始>
// Method values not supported.
# <翻译结束>


<原文开始>
// implemented in package runtime
<原文结束>

# <翻译开始>
// implemented in package runtime
# <翻译结束>


<原文开始>
// ValueOf returns a new Value initialized to the concrete value
// stored in the interface i. ValueOf(nil) returns the zero Value.
<原文结束>

# <翻译开始>
// ValueOf returns a new Value initialized to the concrete value
// stored in the interface i. ValueOf(nil) returns the zero Value.
# <翻译结束>


<原文开始>
	// TODO: Maybe allow contents of a Value to live on the stack.
	// For now we make the contents always escape to the heap. It
	// makes life easier in a few places (see chanrecv/mapassign
	// comment below).
<原文结束>

# <翻译开始>
	// TODO: Maybe allow contents of a Value to live on the stack.
	// For now we make the contents always escape to the heap. It
	// makes life easier in a few places (see chanrecv/mapassign
	// comment below).
# <翻译结束>


<原文开始>
// assignTo returns a value v that can be assigned directly to typ.
// It panics if v is not assignable to typ.
// For a conversion to an interface type, target is a suggested scratch space to use.
<原文结束>

# <翻译开始>
// assignTo returns a value v that can be assigned directly to typ.
// It panics if v is not assignable to typ.
// For a conversion to an interface type, target is a suggested scratch space to use.
# <翻译结束>


<原文开始>
	// if v.flag&flagMethod != 0 {
	// 	v = makeMethodValue(context, v)
	// }
<原文结束>

# <翻译开始>
	// if v.flag&flagMethod != 0 {
	// 	v = makeMethodValue(context, v)
	// }
# <翻译结束>


<原文开始>
		// Overwrite type so that they match.
		// Same memory layout, so no harm done.
<原文结束>

# <翻译开始>
		// Overwrite type so that they match.
		// Same memory layout, so no harm done.
# <翻译结束>


<原文开始>
			// A nil ReadWriter passed to nil Reader is OK,
			// but using ifaceE2I below will panic.
			// Avoid the panic by returning a nil dst (e.g., Reader) explicitly.
<原文结束>

# <翻译开始>
			// A nil ReadWriter passed to nil Reader is OK,
			// but using ifaceE2I below will panic.
			// Avoid the panic by returning a nil dst (e.g., Reader) explicitly.
# <翻译结束>


<原文开始>
// arrayAt returns the i-th element of p,
// an array whose elements are eltSize bytes wide.
// The array pointed at by p must have at least i+1 elements:
// it is invalid (but impossible to check here) to pass i >= len,
// because then the result will point outside the array.
// whySafe must explain why i < len. (Passing "i < len" is fine;
// the benefit is to surface this assumption at the call site.)
<原文结束>

# <翻译开始>
// arrayAt returns the i-th element of p,
// an array whose elements are eltSize bytes wide.
// The array pointed at by p must have at least i+1 elements:
// it is invalid (but impossible to check here) to pass i >= len,
// because then the result will point outside the array.
// whySafe must explain why i < len. (Passing "i < len" is fine;
// the benefit is to surface this assumption at the call site.)
# <翻译结束>


<原文开始>
// Dummy annotation marking that the value x escapes,
// for use in cases where the reflect code is so clever that
// the compiler cannot follow.
<原文结束>

# <翻译开始>
// Dummy annotation marking that the value x escapes,
// for use in cases where the reflect code is so clever that
// the compiler cannot follow.
# <翻译结束>

