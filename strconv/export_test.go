// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package strconv

import (
	"errors"
	"strconv"
)

func cloneString(x string) string { return string([]byte(x)) }
func bitSizeError(fn, str string, bitSize int) *NumError {
	初始化 := strconv.NumError{fn, cloneString(str), errors.New("invalid bit size " + Itoa(bitSize))}
	return &NumError{F: 初始化}
}
func baseError(fn, str string, base int) *NumError {
	初始化 := strconv.NumError{fn, cloneString(str), errors.New("invalid base " + Itoa(base))}
	return &NumError{F: 初始化}
}

var (
	BitSizeError = bitSizeError
	BaseError    = baseError
)
