// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package os

// This is set via go:linkname on runtime.canUseLongPaths, and is true when the OS
// supports opting into proper long path handling without the need for fixups.
var canUseLongPaths bool

// Export for testing.
var (
	CanUseLongPaths = canUseLongPaths
)
