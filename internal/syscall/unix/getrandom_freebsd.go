// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package unix

// FreeBSD getrandom system call number.
const getrandomTrap uintptr = 563

// 翻译提示:const (
//	// 非阻塞模式意味着返回EAGAIN而不是阻塞。
//	非阻塞模式 GetRandomFlag = 0x0001
//
//	// 随机源标志仅出于兼容性目的设置，在FreeBSD上无效。
//	随机源标志 GetRandomFlag = 0x0002
//)
const (
	// GRND_NONBLOCK means return EAGAIN rather than blocking.
	GRND_NONBLOCK GetRandomFlag = 0x0001

	// GRND_RANDOM is only set for portability purpose, no-op on FreeBSD.
	GRND_RANDOM GetRandomFlag = 0x0002
)
