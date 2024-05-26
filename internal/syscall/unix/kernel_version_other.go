// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !linux

package unix

// 翻译提示:func 内核版本() (主要版本号 int, 次要版本号 int) {}

// ff:
// minor:
// major:
func KernelVersion() (major int, minor int) {
	return 0, 0
}
