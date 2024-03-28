
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
// Package reflectlite implements lightweight version of reflect, not using
// any package except for "runtime" and "unsafe".
<原文结束>

# <翻译开始>
// Package reflectlite implements lightweight version of reflect, not using
// any package except for "runtime" and "unsafe".
# <翻译结束>


<原文开始>
// Type is the representation of a Go type.
//
// Not all methods apply to all kinds of types. Restrictions,
// if any, are noted in the documentation for each method.
// Use the Kind method to find out the kind of type before
// calling kind-specific methods. Calling a method
// inappropriate to the kind of type causes a run-time panic.
//
// Type values are comparable, such as with the == operator,
// so they can be used as map keys.
// Two Type values are equal if they represent identical types.
<原文结束>

# <翻译开始>
// Type is the representation of a Go type.
//
// Not all methods apply to all kinds of types. Restrictions,
// if any, are noted in the documentation for each method.
// Use the Kind method to find out the kind of type before
// calling kind-specific methods. Calling a method
// inappropriate to the kind of type causes a run-time panic.
//
// Type values are comparable, such as with the == operator,
// so they can be used as map keys.
// Two Type values are equal if they represent identical types.
# <翻译结束>


<原文开始>
// Methods applicable to all types.
<原文结束>

# <翻译开始>
// Methods applicable to all types.
# <翻译结束>


<原文开始>
	// Name returns the type's name within its package for a defined type.
	// For other (non-defined) types it returns the empty string.
<原文结束>

# <翻译开始>
	// Name returns the type's name within its package for a defined type.
	// For other (non-defined) types it returns the empty string.
# <翻译结束>


<原文开始>
	// PkgPath returns a defined type's package path, that is, the import path
	// that uniquely identifies the package, such as "encoding/base64".
	// If the type was predeclared (string, error) or not defined (*T, struct{},
	// []int, or A where A is an alias for a non-defined type), the package path
	// will be the empty string.
<原文结束>

# <翻译开始>
	// PkgPath returns a defined type's package path, that is, the import path
	// that uniquely identifies the package, such as "encoding/base64".
	// If the type was predeclared (string, error) or not defined (*T, struct{},
	// []int, or A where A is an alias for a non-defined type), the package path
	// will be the empty string.
# <翻译结束>


<原文开始>
	// Size returns the number of bytes needed to store
	// a value of the given type; it is analogous to unsafe.Sizeof.
<原文结束>

# <翻译开始>
	// Size returns the number of bytes needed to store
	// a value of the given type; it is analogous to unsafe.Sizeof.
# <翻译结束>


<原文开始>
// Kind returns the specific kind of this type.
<原文结束>

# <翻译开始>
// Kind returns the specific kind of this type.
# <翻译结束>


<原文开始>
// Implements reports whether the type implements the interface type u.
<原文结束>

# <翻译开始>
// Implements reports whether the type implements the interface type u.
# <翻译结束>


<原文开始>
// AssignableTo reports whether a value of the type is assignable to type u.
<原文结束>

# <翻译开始>
// AssignableTo reports whether a value of the type is assignable to type u.
# <翻译结束>


<原文开始>
// Comparable reports whether values of this type are comparable.
<原文结束>

# <翻译开始>
// Comparable reports whether values of this type are comparable.
# <翻译结束>


<原文开始>
	// String returns a string representation of the type.
	// The string representation may use shortened package names
	// (e.g., base64 instead of "encoding/base64") and is not
	// guaranteed to be unique among types. To test for type identity,
	// compare the Types directly.
<原文结束>

# <翻译开始>
	// String returns a string representation of the type.
	// The string representation may use shortened package names
	// (e.g., base64 instead of "encoding/base64") and is not
	// guaranteed to be unique among types. To test for type identity,
	// compare the Types directly.
# <翻译结束>


<原文开始>
	// Elem returns a type's element type.
	// It panics if the type's Kind is not Ptr.
<原文结束>

# <翻译开始>
	// Elem returns a type's element type.
	// It panics if the type's Kind is not Ptr.
# <翻译结束>


<原文开始>
// A Kind represents the specific kind of type that a Type represents.
// The zero Kind is not a valid kind.
<原文结束>

# <翻译开始>
// A Kind represents the specific kind of type that a Type represents.
// The zero Kind is not a valid kind.
# <翻译结束>


<原文开始>
// tflag is used by an rtype to signal what extra type information is
// available in the memory directly following the rtype value.
//
// tflag values must be kept in sync with copies in:
//
//	cmd/compile/internal/reflectdata/reflect.go
//	cmd/link/internal/ld/decodesym.go
//	runtime/type.go
<原文结束>

# <翻译开始>
// tflag is used by an rtype to signal what extra type information is
// available in the memory directly following the rtype value.
//
// tflag values must be kept in sync with copies in:
//
//	cmd/compile/internal/reflectdata/reflect.go
//	cmd/link/internal/ld/decodesym.go
//	runtime/type.go
# <翻译结束>


<原文开始>
	// tflagUncommon means that there is a pointer, *uncommonType,
	// just beyond the outer type structure.
	//
	// For example, if t.Kind() == Struct and t.tflag&tflagUncommon != 0,
	// then t has uncommonType data and it can be accessed as:
	//
	//	type tUncommon struct {
	//		structType
	//		u uncommonType
	//	}
	//	u := &(*tUncommon)(unsafe.Pointer(t)).u
<原文结束>

# <翻译开始>
	// tflagUncommon means that there is a pointer, *uncommonType,
	// just beyond the outer type structure.
	//
	// For example, if t.Kind() == Struct and t.tflag&tflagUncommon != 0,
	// then t has uncommonType data and it can be accessed as:
	//
	//	type tUncommon struct {
	//		structType
	//		u uncommonType
	//	}
	//	u := &(*tUncommon)(unsafe.Pointer(t)).u
# <翻译结束>


<原文开始>
	// tflagExtraStar means the name in the str field has an
	// extraneous '*' prefix. This is because for most types T in
	// a program, the type *T also exists and reusing the str data
	// saves binary size.
<原文结束>

# <翻译开始>
	// tflagExtraStar means the name in the str field has an
	// extraneous '*' prefix. This is because for most types T in
	// a program, the type *T also exists and reusing the str data
	// saves binary size.
# <翻译结束>


<原文开始>
// tflagNamed means the type has a name.
<原文结束>

# <翻译开始>
// tflagNamed means the type has a name.
# <翻译结束>


<原文开始>
	// tflagRegularMemory means that equal and hash functions can treat
	// this type as a single region of t.size bytes.
<原文结束>

# <翻译开始>
	// tflagRegularMemory means that equal and hash functions can treat
	// this type as a single region of t.size bytes.
# <翻译结束>


<原文开始>
// rtype is the common implementation of most values.
// It is embedded in other struct types.
//
// rtype must be kept in sync with ../runtime/type.go:/^type._type.
<原文结束>

# <翻译开始>
// rtype is the common implementation of most values.
// It is embedded in other struct types.
//
// rtype must be kept in sync with ../runtime/type.go:/^type._type.
# <翻译结束>


<原文开始>
// number of bytes in the type that can contain pointers
<原文结束>

# <翻译开始>
// number of bytes in the type that can contain pointers
# <翻译结束>


<原文开始>
// hash of type; avoids computation in hash tables
<原文结束>

# <翻译开始>
// hash of type; avoids computation in hash tables
# <翻译结束>


<原文开始>
// extra type information flags
<原文结束>

# <翻译开始>
// extra type information flags
# <翻译结束>


<原文开始>
// alignment of variable with this type
<原文结束>

# <翻译开始>
// alignment of variable with this type
# <翻译结束>


<原文开始>
// alignment of struct field with this type
<原文结束>

# <翻译开始>
// alignment of struct field with this type
# <翻译结束>


<原文开始>
	// function for comparing objects of this type
	// (ptr to object A, ptr to object B) -> ==?
<原文结束>

# <翻译开始>
	// function for comparing objects of this type
	// (ptr to object A, ptr to object B) -> ==?
# <翻译结束>


<原文开始>
// garbage collection data
<原文结束>

# <翻译开始>
// garbage collection data
# <翻译结束>


<原文开始>
// type for pointer to this type, may be zero
<原文结束>

# <翻译开始>
// type for pointer to this type, may be zero
# <翻译结束>


<原文开始>
// Method on non-interface type
<原文结束>

# <翻译开始>
// Method on non-interface type
# <翻译结束>


<原文开始>
// method type (without receiver)
<原文结束>

# <翻译开始>
// method type (without receiver)
# <翻译结束>


<原文开始>
// fn used in interface call (one-word receiver)
<原文结束>

# <翻译开始>
// fn used in interface call (one-word receiver)
# <翻译结束>


<原文开始>
// fn used for normal method call
<原文结束>

# <翻译开始>
// fn used for normal method call
# <翻译结束>


<原文开始>
// uncommonType is present only for defined types or types with methods
// (if T is a defined type, the uncommonTypes for T and *T have methods).
// Using a pointer to this struct reduces the overall size required
// to describe a non-defined type with no methods.
<原文结束>

# <翻译开始>
// uncommonType is present only for defined types or types with methods
// (if T is a defined type, the uncommonTypes for T and *T have methods).
// Using a pointer to this struct reduces the overall size required
// to describe a non-defined type with no methods.
# <翻译结束>


<原文开始>
// import path; empty for built-in types like int, string
<原文结束>

# <翻译开始>
// import path; empty for built-in types like int, string
# <翻译结束>


<原文开始>
// number of exported methods
<原文结束>

# <翻译开始>
// number of exported methods
# <翻译结束>


<原文开始>
// offset from this uncommontype to [mcount]method
<原文结束>

# <翻译开始>
// offset from this uncommontype to [mcount]method
# <翻译结束>


<原文开始>
// chanDir represents a channel type's direction.
<原文结束>

# <翻译开始>
// chanDir represents a channel type's direction.
# <翻译结束>


<原文开始>
// arrayType represents a fixed array type.
<原文结束>

# <翻译开始>
// arrayType represents a fixed array type.
# <翻译结束>


<原文开始>
// chanType represents a channel type.
<原文结束>

# <翻译开始>
// chanType represents a channel type.
# <翻译结束>


<原文开始>
// channel direction (chanDir)
<原文结束>

# <翻译开始>
// channel direction (chanDir)
# <翻译结束>


<原文开始>
// funcType represents a function type.
//
// A *rtype for each in and out parameter is stored in an array that
// directly follows the funcType (and possibly its uncommonType). So
// a function type with one method, one input, and one output is:
//
//	struct {
//		funcType
//		uncommonType
//		[2]*rtype    // [0] is in, [1] is out
//	}
<原文结束>

# <翻译开始>
// funcType represents a function type.
//
// A *rtype for each in and out parameter is stored in an array that
// directly follows the funcType (and possibly its uncommonType). So
// a function type with one method, one input, and one output is:
//
//	struct {
//		funcType
//		uncommonType
//		[2]*rtype    // [0] is in, [1] is out
//	}
# <翻译结束>


<原文开始>
// top bit is set if last input parameter is ...
<原文结束>

# <翻译开始>
// top bit is set if last input parameter is ...
# <翻译结束>


<原文开始>
// imethod represents a method on an interface type
<原文结束>

# <翻译开始>
// imethod represents a method on an interface type
# <翻译结束>


<原文开始>
// .(*FuncType) underneath
<原文结束>

# <翻译开始>
// .(*FuncType) underneath
# <翻译结束>


<原文开始>
// interfaceType represents an interface type.
<原文结束>

# <翻译开始>
// interfaceType represents an interface type.
# <翻译结束>


<原文开始>
// mapType represents a map type.
<原文结束>

# <翻译开始>
// mapType represents a map type.
# <翻译结束>


<原文开始>
// map element (value) type
<原文结束>

# <翻译开始>
// map element (value) type
# <翻译结束>


<原文开始>
// internal bucket structure
<原文结束>

# <翻译开始>
// internal bucket structure
# <翻译结束>


<原文开始>
// function for hashing keys (ptr to key, seed) -> hash
<原文结束>

# <翻译开始>
// function for hashing keys (ptr to key, seed) -> hash
# <翻译结束>


<原文开始>
// ptrType represents a pointer type.
<原文结束>

# <翻译开始>
// ptrType represents a pointer type.
# <翻译结束>


<原文开始>
// pointer element (pointed at) type
<原文结束>

# <翻译开始>
// pointer element (pointed at) type
# <翻译结束>


<原文开始>
// sliceType represents a slice type.
<原文结束>

# <翻译开始>
// sliceType represents a slice type.
# <翻译结束>


<原文开始>
// name is always non-empty
<原文结束>

# <翻译开始>
// name is always non-empty
# <翻译结束>


<原文开始>
// structType represents a struct type.
<原文结束>

# <翻译开始>
// structType represents a struct type.
# <翻译结束>


<原文开始>
// name is an encoded type name with optional extra data.
//
// The first byte is a bit field containing:
//
//	1<<0 the name is exported
//	1<<1 tag data follows the name
//	1<<2 pkgPath nameOff follows the name and tag
//
// The next two bytes are the data length:
//
//	l := uint16(data[1])<<8 | uint16(data[2])
//
// Bytes [3:3+l] are the string data.
//
// If tag data follows then bytes 3+l and 3+l+1 are the tag length,
// with the data following.
//
// If the import path follows, then 4 bytes at the end of
// the data form a nameOff. The import path is only set for concrete
// methods that are defined in a different package than their type.
//
// If a name starts with "*", then the exported bit represents
// whether the pointed to type is exported.
<原文结束>

# <翻译开始>
// name is an encoded type name with optional extra data.
//
// The first byte is a bit field containing:
//
//	1<<0 the name is exported
//	1<<1 tag data follows the name
//	1<<2 pkgPath nameOff follows the name and tag
//
// The next two bytes are the data length:
//
//	l := uint16(data[1])<<8 | uint16(data[2])
//
// Bytes [3:3+l] are the string data.
//
// If tag data follows then bytes 3+l and 3+l+1 are the tag length,
// with the data following.
//
// If the import path follows, then 4 bytes at the end of
// the data form a nameOff. The import path is only set for concrete
// methods that are defined in a different package than their type.
//
// If a name starts with "*", then the exported bit represents
// whether the pointed to type is exported.
# <翻译结束>


<原文开始>
// readVarint parses a varint as encoded by encoding/binary.
// It returns the number of encoded bytes and the encoded value.
<原文结束>

# <翻译开始>
// readVarint parses a varint as encoded by encoding/binary.
// It returns the number of encoded bytes and the encoded value.
# <翻译结束>


<原文开始>
	// Note that this field may not be aligned in memory,
	// so we cannot use a direct int32 assignment here.
<原文结束>

# <翻译开始>
	// Note that this field may not be aligned in memory,
	// so we cannot use a direct int32 assignment here.
# <翻译结束>


<原文开始>
// Type.gc points to GC program
<原文结束>

# <翻译开始>
// Type.gc points to GC program
# <翻译结束>


<原文开始>
// String returns the name of k.
<原文结束>

# <翻译开始>
// String returns the name of k.
# <翻译结束>


<原文开始>
// resolveNameOff resolves a name offset from a base pointer.
// The (*rtype).nameOff method is a convenience wrapper for this function.
// Implemented in the runtime package.
<原文结束>

# <翻译开始>
// resolveNameOff resolves a name offset from a base pointer.
// The (*rtype).nameOff method is a convenience wrapper for this function.
// Implemented in the runtime package.
# <翻译结束>


<原文开始>
// resolveTypeOff resolves an *rtype offset from a base type.
// The (*rtype).typeOff method is a convenience wrapper for this function.
// Implemented in the runtime package.
<原文结束>

# <翻译开始>
// resolveTypeOff resolves an *rtype offset from a base type.
// The (*rtype).typeOff method is a convenience wrapper for this function.
// Implemented in the runtime package.
# <翻译结束>


<原文开始>
// offset from top of text section
<原文结束>

# <翻译开始>
// offset from top of text section
# <翻译结束>


<原文开始>
// add returns p+x.
//
// The whySafe string is ignored, so that the function still inlines
// as efficiently as p+x, but all call sites should use the string to
// record why the addition is safe, which is to say why the addition
// does not cause x to advance to the very end of p's allocation
// and therefore point incorrectly at the next block in memory.
<原文结束>

# <翻译开始>
// add returns p+x.
//
// The whySafe string is ignored, so that the function still inlines
// as efficiently as p+x, but all call sites should use the string to
// record why the addition is safe, which is to say why the addition
// does not cause x to advance to the very end of p's allocation
// and therefore point incorrectly at the next block in memory.
# <翻译结束>


<原文开始>
// NumMethod returns the number of interface methods in the type's method set.
<原文结束>

# <翻译开始>
// NumMethod returns the number of interface methods in the type's method set.
# <翻译结束>


<原文开始>
// TypeOf returns the reflection Type that represents the dynamic type of i.
// If i is a nil interface value, TypeOf returns nil.
<原文结束>

# <翻译开始>
// TypeOf returns the reflection Type that represents the dynamic type of i.
// If i is a nil interface value, TypeOf returns nil.
# <翻译结束>


<原文开始>
// implements reports whether the type V implements the interface type T.
<原文结束>

# <翻译开始>
// implements reports whether the type V implements the interface type T.
# <翻译结束>


<原文开始>
	// The same algorithm applies in both cases, but the
	// method tables for an interface type and a concrete type
	// are different, so the code is duplicated.
	// In both cases the algorithm is a linear scan over the two
	// lists - T's methods and V's methods - simultaneously.
	// Since method tables are stored in a unique sorted order
	// (alphabetical, with no duplicate method names), the scan
	// through V's methods must hit a match for each of T's
	// methods along the way, or else V does not implement T.
	// This lets us run the scan in overall linear time instead of
	// the quadratic time  a naive search would require.
	// See also ../runtime/iface.go.
<原文结束>

# <翻译开始>
	// The same algorithm applies in both cases, but the
	// method tables for an interface type and a concrete type
	// are different, so the code is duplicated.
	// In both cases the algorithm is a linear scan over the two
	// lists - T's methods and V's methods - simultaneously.
	// Since method tables are stored in a unique sorted order
	// (alphabetical, with no duplicate method names), the scan
	// through V's methods must hit a match for each of T's
	// methods along the way, or else V does not implement T.
	// This lets us run the scan in overall linear time instead of
	// the quadratic time  a naive search would require.
	// See also ../runtime/iface.go.
# <翻译结束>


<原文开始>
// directlyAssignable reports whether a value x of type V can be directly
// assigned (using memmove) to a value of type T.
// https://golang.org/doc/go_spec.html#Assignability
// Ignoring the interface rules (implemented elsewhere)
// and the ideal constant rules (no ideal constants at run time).
<原文结束>

# <翻译开始>
// directlyAssignable reports whether a value x of type V can be directly
// assigned (using memmove) to a value of type T.
// https://golang.org/doc/go_spec.html#Assignability
// Ignoring the interface rules (implemented elsewhere)
// and the ideal constant rules (no ideal constants at run time).
# <翻译结束>


<原文开始>
// x's type V is identical to T?
<原文结束>

# <翻译开始>
// x's type V is identical to T?
# <翻译结束>


<原文开始>
	// Otherwise at least one of T and V must not be defined
	// and they must have the same kind.
<原文结束>

# <翻译开始>
	// Otherwise at least one of T and V must not be defined
	// and they must have the same kind.
# <翻译结束>


<原文开始>
// x's type T and V must  have identical underlying types.
<原文结束>

# <翻译开始>
// x's type T and V must  have identical underlying types.
# <翻译结束>


<原文开始>
	// Non-composite types of equal kind have same underlying type
	// (the predefined instance of the type).
<原文结束>

# <翻译开始>
	// Non-composite types of equal kind have same underlying type
	// (the predefined instance of the type).
# <翻译结束>


<原文开始>
		// Special case:
		// x is a bidirectional channel value, T is a channel type,
		// and x's type V and T have identical element types.
<原文结束>

# <翻译开始>
		// Special case:
		// x is a bidirectional channel value, T is a channel type,
		// and x's type V and T have identical element types.
# <翻译结束>


<原文开始>
// Otherwise continue test for identical underlying type.
<原文结束>

# <翻译开始>
// Otherwise continue test for identical underlying type.
# <翻译结束>


<原文开始>
		// Might have the same methods but still
		// need a run time conversion.
<原文结束>

# <翻译开始>
		// Might have the same methods but still
		// need a run time conversion.
# <翻译结束>


<原文开始>
// toType converts from a *rtype to a Type that can be returned
// to the client of package reflect. In gc, the only concern is that
// a nil *rtype must be replaced by a nil Type, but in gccgo this
// function takes care of ensuring that multiple *rtype for the same
// type are coalesced into a single Type.
<原文结束>

# <翻译开始>
// toType converts from a *rtype to a Type that can be returned
// to the client of package reflect. In gc, the only concern is that
// a nil *rtype must be replaced by a nil Type, but in gccgo this
// function takes care of ensuring that multiple *rtype for the same
// type are coalesced into a single Type.
# <翻译结束>


<原文开始>
// ifaceIndir reports whether t is stored indirectly in an interface value.
<原文结束>

# <翻译开始>
// ifaceIndir reports whether t is stored indirectly in an interface value.
# <翻译结束>

