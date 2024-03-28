
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
// Package intern lets you make smaller comparable values by boxing
// a larger comparable value (such as a 16 byte string header) down
// into a globally unique 8 byte pointer.
//
// The globally unique pointers are garbage collected with weak
// references and finalizers. This package hides that.
<原文结束>

# <翻译开始>
// Package intern lets you make smaller comparable values by boxing
// a larger comparable value (such as a 16 byte string header) down
// into a globally unique 8 byte pointer.
//
// The globally unique pointers are garbage collected with weak
// references and finalizers. This package hides that.
# <翻译结束>


<原文开始>
// A Value pointer is the handle to an underlying comparable value.
// See func Get for how Value pointers may be used.
<原文结束>

# <翻译开始>
// A Value pointer is the handle to an underlying comparable value.
// See func Get for how Value pointers may be used.
# <翻译结束>


<原文开始>
// prevent people from accidentally using value type as comparable
<原文结束>

# <翻译开始>
// prevent people from accidentally using value type as comparable
# <翻译结束>


<原文开始>
	// resurrected is guarded by mu (for all instances of Value).
	// It is set true whenever v is synthesized from a uintptr.
<原文结束>

# <翻译开始>
	// resurrected is guarded by mu (for all instances of Value).
	// It is set true whenever v is synthesized from a uintptr.
# <翻译结束>


<原文开始>
// Get returns the comparable value passed to the Get func
// that returned v.
<原文结束>

# <翻译开始>
// Get returns the comparable value passed to the Get func
// that returned v.
# <翻译结束>


<原文开始>
// key is a key in our global value map.
// It contains type-specialized fields to avoid allocations
// when converting common types to empty interfaces.
<原文结束>

# <翻译开始>
// key is a key in our global value map.
// It contains type-specialized fields to avoid allocations
// when converting common types to empty interfaces.
# <翻译结束>


<原文开始>
	// isString reports whether key contains a string.
	// Without it, the zero value of key is ambiguous.
<原文结束>

# <翻译开始>
	// isString reports whether key contains a string.
	// Without it, the zero value of key is ambiguous.
# <翻译结束>


<原文开始>
// keyFor returns a key to use with cmpVal.
<原文结束>

# <翻译开始>
// keyFor returns a key to use with cmpVal.
# <翻译结束>


<原文开始>
// Value returns a *Value built from k.
<原文结束>

# <翻译开始>
// Value returns a *Value built from k.
# <翻译结束>


<原文开始>
	// mu guards valMap, a weakref map of *Value by underlying value.
	// It also guards the resurrected field of all *Values.
<原文结束>

# <翻译开始>
	// mu guards valMap, a weakref map of *Value by underlying value.
	// It also guards the resurrected field of all *Values.
# <翻译结束>


<原文开始>
// non-nil in safe+leaky mode
<原文结束>

# <翻译开始>
// non-nil in safe+leaky mode
# <翻译结束>


<原文开始>
// safeMap returns a non-nil map if we're in safe-but-leaky mode,
// as controlled by GODEBUG=intern=leaky
<原文结束>

# <翻译开始>
// safeMap returns a non-nil map if we're in safe-but-leaky mode,
// as controlled by GODEBUG=intern=leaky
# <翻译结束>


<原文开始>
// Get returns a pointer representing the comparable value cmpVal.
//
// The returned pointer will be the same for Get(v) and Get(v2)
// if and only if v == v2, and can be used as a map key.
<原文结束>

# <翻译开始>
// Get returns a pointer representing the comparable value cmpVal.
//
// The returned pointer will be the same for Get(v) and Get(v2)
// if and only if v == v2, and can be used as a map key.
# <翻译结束>


<原文开始>
// GetByString is identical to Get, except that it is specialized for strings.
// This avoids an allocation from putting a string into an interface{}
// to pass as an argument to Get.
<原文结束>

# <翻译开始>
// GetByString is identical to Get, except that it is specialized for strings.
// This avoids an allocation from putting a string into an interface{}
// to pass as an argument to Get.
# <翻译结束>


<原文开始>
		// SetFinalizer before uintptr conversion (theoretical concern;
		// see https://github.com/go4org/intern/issues/13)
<原文结束>

# <翻译开始>
		// SetFinalizer before uintptr conversion (theoretical concern;
		// see https://github.com/go4org/intern/issues/13)
# <翻译结束>


<原文开始>
		// We lost the race. Somebody resurrected it while we
		// were about to finalize it. Try again next round.
<原文结束>

# <翻译开始>
		// We lost the race. Somebody resurrected it while we
		// were about to finalize it. Try again next round.
# <翻译结束>


<原文开始>
// Interning is simple if you don't require that unused values be
// garbage collectable. But we do require that; we don't want to be
// DOS vector. We do this by using a uintptr to hide the pointer from
// the garbage collector, and using a finalizer to eliminate the
// pointer when no other code is using it.
//
// The obvious implementation of this is to use a
// map[interface{}]uintptr-of-*interface{}, and set up a finalizer to
// delete from the map. Unfortunately, this is racy. Because pointers
// are being created in violation of Go's unsafety rules, it's
// possible to create a pointer to a value concurrently with the GC
// concluding that the value can be collected. There are other races
// that break the equality invariant as well, but the use-after-free
// will cause a runtime crash.
//
// To make this work, the finalizer needs to know that no references
// have been unsafely created since the finalizer was set up. To do
// this, values carry a "resurrected" sentinel, which gets set
// whenever a pointer is unsafely created. If the finalizer encounters
// the sentinel, it clears the sentinel and delays collection for one
// additional GC cycle, by re-installing itself as finalizer. This
// ensures that the unsafely created pointer is visible to the GC, and
// will correctly prevent collection.
//
// This technique does mean that interned values that get reused take
// at least 3 GC cycles to fully collect (1 to clear the sentinel, 1
// to clean up the unsafe map, 1 to be actually deleted).
//
// @ianlancetaylor commented in
// https://github.com/golang/go/issues/41303#issuecomment-717401656
// that it is possible to implement weak references in terms of
// finalizers without unsafe. Unfortunately, the approach he outlined
// does not work here, for two reasons. First, there is no way to
// construct a strong pointer out of a weak pointer; our map stores
// weak pointers, but we must return strong pointers to callers.
// Second, and more fundamentally, we must return not just _a_ strong
// pointer to callers, but _the same_ strong pointer to callers. In
// order to return _the same_ strong pointer to callers, we must track
// it, which is exactly what we cannot do with strong pointers.
//
// See https://github.com/inetaf/netaddr/issues/53 for more
// discussion, and https://github.com/go4org/intern/issues/2 for an
// illustration of the subtleties at play.
<原文结束>

# <翻译开始>
// Interning is simple if you don't require that unused values be
// garbage collectable. But we do require that; we don't want to be
// DOS vector. We do this by using a uintptr to hide the pointer from
// the garbage collector, and using a finalizer to eliminate the
// pointer when no other code is using it.
//
// The obvious implementation of this is to use a
// map[interface{}]uintptr-of-*interface{}, and set up a finalizer to
// delete from the map. Unfortunately, this is racy. Because pointers
// are being created in violation of Go's unsafety rules, it's
// possible to create a pointer to a value concurrently with the GC
// concluding that the value can be collected. There are other races
// that break the equality invariant as well, but the use-after-free
// will cause a runtime crash.
//
// To make this work, the finalizer needs to know that no references
// have been unsafely created since the finalizer was set up. To do
// this, values carry a "resurrected" sentinel, which gets set
// whenever a pointer is unsafely created. If the finalizer encounters
// the sentinel, it clears the sentinel and delays collection for one
// additional GC cycle, by re-installing itself as finalizer. This
// ensures that the unsafely created pointer is visible to the GC, and
// will correctly prevent collection.
//
// This technique does mean that interned values that get reused take
// at least 3 GC cycles to fully collect (1 to clear the sentinel, 1
// to clean up the unsafe map, 1 to be actually deleted).
//
// @ianlancetaylor commented in
// https://github.com/golang/go/issues/41303#issuecomment-717401656
// that it is possible to implement weak references in terms of
// finalizers without unsafe. Unfortunately, the approach he outlined
// does not work here, for two reasons. First, there is no way to
// construct a strong pointer out of a weak pointer; our map stores
// weak pointers, but we must return strong pointers to callers.
// Second, and more fundamentally, we must return not just _a_ strong
// pointer to callers, but _the same_ strong pointer to callers. In
// order to return _the same_ strong pointer to callers, we must track
// it, which is exactly what we cannot do with strong pointers.
//
// See https://github.com/inetaf/netaddr/issues/53 for more
// discussion, and https://github.com/go4org/intern/issues/2 for an
// illustration of the subtleties at play.
# <翻译结束>

