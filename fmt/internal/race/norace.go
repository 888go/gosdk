// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !race

package race

import (
	"unsafe"
)

const Enabled = false


// ff:
// addr:
func Acquire(addr unsafe.Pointer) {
}


// ff:
// addr:
func Release(addr unsafe.Pointer) {
}


// ff:
// addr:
func ReleaseMerge(addr unsafe.Pointer) {
}


// ff:
func Disable() {
}


// ff:
func Enable() {
}


// ff:
// addr:
func Read(addr unsafe.Pointer) {
}


// ff:
// addr:
func Write(addr unsafe.Pointer) {
}


// ff:
// len:
// addr:
func ReadRange(addr unsafe.Pointer, len int) {
}


// ff:
// len:
// addr:
func WriteRange(addr unsafe.Pointer, len int) {
}


// ff:
func Errors() int { return 0 }
