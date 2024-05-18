// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package os

import (
	"errors"
	"syscall"
	"time"
)

// lstat is overridden in tests.
var lstat = Lstat

// testingForceReadDirLstat forces ReadDir to call Lstat, for testing that code path.
// This can be difficult to provoke on some Unix systems otherwise.
var testingForceReadDirLstat bool

// For testing.
func atime(fi FileInfo) time.Time {
	return time.Unix(0, fi.Sys().(*syscall.Win32FileAttributeData).LastAccessTime.Nanoseconds())
}

var errWriteAtInAppendMode = errors.New("os: invalid use of WriteAt on file opened with O_APPEND")

var errPatternHasSeparator = errors.New("pattern contains path separator")

// Export for testing.
var Atime = atime
var LstatP = &lstat
var ErrWriteAtInAppendMode = errWriteAtInAppendMode
var TestingForceReadDirLstat = &testingForceReadDirLstat
var ErrPatternHasSeparator = errPatternHasSeparator
