// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build js && wasm

package unix


// ff:
// err:
// nonblocking:
// fd:
func IsNonblock(fd int) (nonblocking bool, err error) {
	return false, nil
}


// ff:
// flag:
func HasNonblockFlag(flag int) bool {
	return false
}
