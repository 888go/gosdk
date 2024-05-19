// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build race

package race

import (
	"runtime"
	"unsafe"
)

const Enabled = true


// ff:
// addr:
func Acquire(addr unsafe.Pointer) {
	runtime.RaceAcquire(addr)
}


// ff:
// addr:
func Release(addr unsafe.Pointer) {
	runtime.RaceRelease(addr)
}


// ff:
// addr:
func ReleaseMerge(addr unsafe.Pointer) {
	runtime.RaceReleaseMerge(addr)
}


// ff:
func Disable() {
	runtime.RaceDisable()
}


// ff:
func Enable() {
	runtime.RaceEnable()
}


// ff:
// addr:
func Read(addr unsafe.Pointer) {
	runtime.RaceRead(addr)
}


// ff:
// addr:
func Write(addr unsafe.Pointer) {
	runtime.RaceWrite(addr)
}


// ff:
// len:
// addr:
func ReadRange(addr unsafe.Pointer, len int) {
	runtime.RaceReadRange(addr, len)
}


// ff:
// len:
// addr:
func WriteRange(addr unsafe.Pointer, len int) {
	runtime.RaceWriteRange(addr, len)
}


// ff:
func Errors() int {
	return runtime.RaceErrors()
}
