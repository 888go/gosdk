
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
// Package orderedmap provides an ordered map, implemented as a binary tree.
<原文结束>

# <翻译开始>
// Package orderedmap provides an ordered map, implemented as a binary tree.
# <翻译结束>


<原文开始>
// TODO(gri) fix imports for tests
<原文结束>

# <翻译开始>
// TODO(gri) fix imports for tests
# <翻译结束>


<原文开始>
// node is the type of a node in the binary tree.
<原文结束>

# <翻译开始>
// node is the type of a node in the binary tree.
# <翻译结束>


<原文开始>
// find looks up key in the map, and returns either a pointer
// to the node holding key, or a pointer to the location where
// such a node would go.
<原文结束>

# <翻译开始>
// find looks up key in the map, and returns either a pointer
// to the node holding key, or a pointer to the location where
// such a node would go.
# <翻译结束>


<原文开始>
// Insert inserts a new key/value into the map.
// If the key is already present, the value is replaced.
// Returns true if this is a new key, false if already present.
<原文结束>

# <翻译开始>
// Insert inserts a new key/value into the map.
// If the key is already present, the value is replaced.
// Returns true if this is a new key, false if already present.
# <翻译结束>


<原文开始>
// Find returns the value associated with a key, or zero if not present.
// The found result reports whether the key was found.
<原文结束>

# <翻译开始>
// Find returns the value associated with a key, or zero if not present.
// The found result reports whether the key was found.
# <翻译结束>


<原文开始>
// see the discussion of zero values, above
<原文结束>

# <翻译开始>
// see the discussion of zero values, above
# <翻译结束>


<原文开始>
// keyValue is a pair of key and value used when iterating.
<原文结束>

# <翻译开始>
// keyValue is a pair of key and value used when iterating.
# <翻译结束>


<原文开始>
// InOrder returns an iterator that does an in-order traversal of the map.
<原文结束>

# <翻译开始>
// InOrder returns an iterator that does an in-order traversal of the map.
# <翻译结束>


<原文开始>
		// Stop sending values if sender.Send returns false,
		// meaning that nothing is listening at the receiver end.
<原文结束>

# <翻译开始>
		// Stop sending values if sender.Send returns false,
		// meaning that nothing is listening at the receiver end.
# <翻译结束>


<原文开始>
// Iterator is used to iterate over the map.
<原文结束>

# <翻译开始>
// Iterator is used to iterate over the map.
# <翻译结束>


<原文开始>
// Next returns the next key and value pair, and a boolean indicating
// whether they are valid or whether we have reached the end.
<原文结束>

# <翻译开始>
// Next returns the next key and value pair, and a boolean indicating
// whether they are valid or whether we have reached the end.
# <翻译结束>

