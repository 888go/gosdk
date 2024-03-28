
<原文开始>
// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
# <翻译结束>


<原文开始>
// Package fmtsort provides a general stable ordering mechanism
// for maps, on behalf of the fmt and text/template packages.
// It is not guaranteed to be efficient and works only for types
// that are valid map keys.
<原文结束>

# <翻译开始>
// Package fmtsort provides a general stable ordering mechanism
// for maps, on behalf of the fmt and text/template packages.
// It is not guaranteed to be efficient and works only for types
// that are valid map keys.
# <翻译结束>


<原文开始>
// Note: Throughout this package we avoid calling reflect.Value.Interface as
// it is not always legal to do so and it's easier to avoid the issue than to face it.
<原文结束>

# <翻译开始>
// Note: Throughout this package we avoid calling reflect.Value.Interface as
// it is not always legal to do so and it's easier to avoid the issue than to face it.
# <翻译结束>


<原文开始>
// SortedMap represents a map's keys and values. The keys and values are
// aligned in index order: Value[i] is the value in the map corresponding to Key[i].
<原文结束>

# <翻译开始>
// SortedMap represents a map's keys and values. The keys and values are
// aligned in index order: Value[i] is the value in the map corresponding to Key[i].
# <翻译结束>


<原文开始>
// Sort accepts a map and returns a SortedMap that has the same keys and
// values but in a stable sorted order according to the keys, modulo issues
// raised by unorderable key values such as NaNs.
//
// The ordering rules are more general than with Go's < operator:
//
//   - when applicable, nil compares low
//   - ints, floats, and strings order by <
//   - NaN compares less than non-NaN floats
//   - bool compares false before true
//   - complex compares real, then imag
//   - pointers compare by machine address
//   - channel values compare by machine address
//   - structs compare each field in turn
//   - arrays compare each element in turn.
//     Otherwise identical arrays compare by length.
//   - interface values compare first by reflect.Type describing the concrete type
//     and then by concrete value as described in the previous rules.
<原文结束>

# <翻译开始>
// Sort accepts a map and returns a SortedMap that has the same keys and
// values but in a stable sorted order according to the keys, modulo issues
// raised by unorderable key values such as NaNs.
//
// The ordering rules are more general than with Go's < operator:
//
//   - when applicable, nil compares low
//   - ints, floats, and strings order by <
//   - NaN compares less than non-NaN floats
//   - bool compares false before true
//   - complex compares real, then imag
//   - pointers compare by machine address
//   - channel values compare by machine address
//   - structs compare each field in turn
//   - arrays compare each element in turn.
//     Otherwise identical arrays compare by length.
//   - interface values compare first by reflect.Type describing the concrete type
//     and then by concrete value as described in the previous rules.
# <翻译结束>


<原文开始>
	// Note: this code is arranged to not panic even in the presence
	// of a concurrent map update. The runtime is responsible for
	// yelling loudly if that happens. See issue 33275.
<原文结束>

# <翻译开始>
	// Note: this code is arranged to not panic even in the presence
	// of a concurrent map update. The runtime is responsible for
	// yelling loudly if that happens. See issue 33275.
# <翻译结束>


<原文开始>
// compare compares two values of the same type. It returns -1, 0, 1
// according to whether a > b (1), a == b (0), or a < b (-1).
// If the types differ, it returns -1.
// See the comment on Sort for the comparison rules.
<原文结束>

# <翻译开始>
// compare compares two values of the same type. It returns -1, 0, 1
// according to whether a > b (1), a == b (0), or a < b (-1).
// If the types differ, it returns -1.
// See the comment on Sort for the comparison rules.
# <翻译结束>


<原文开始>
// No good answer possible, but don't return 0: they're not equal.
<原文结束>

# <翻译开始>
// No good answer possible, but don't return 0: they're not equal.
# <翻译结束>


<原文开始>
// Certain types cannot appear as keys (maps, funcs, slices), but be explicit.
<原文结束>

# <翻译开始>
// Certain types cannot appear as keys (maps, funcs, slices), but be explicit.
# <翻译结束>


<原文开始>
// nilCompare checks whether either value is nil. If not, the boolean is false.
// If either value is nil, the boolean is true and the integer is the comparison
// value. The comparison is defined to be 0 if both are nil, otherwise the one
// nil value compares low. Both arguments must represent a chan, func,
// interface, map, pointer, or slice.
<原文结束>

# <翻译开始>
// nilCompare checks whether either value is nil. If not, the boolean is false.
// If either value is nil, the boolean is true and the integer is the comparison
// value. The comparison is defined to be 0 if both are nil, otherwise the one
// nil value compares low. Both arguments must represent a chan, func,
// interface, map, pointer, or slice.
# <翻译结束>


<原文开始>
// floatCompare compares two floating-point values. NaNs compare low.
<原文结束>

# <翻译开始>
// floatCompare compares two floating-point values. NaNs compare low.
# <翻译结束>

