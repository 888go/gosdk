// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

import "github.com/888go/gosdk/internal/cpu"

func archauxv(tag, val uintptr) {
	switch tag {
	case _AT_HWCAP:
		cpu.HWCap = uint(val)
	}
}

func osArchInit() {}
