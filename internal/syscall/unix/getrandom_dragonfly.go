// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package unix

// DragonFlyBSD getrandom system call number.
const getrandomTrap uintptr = 550

// 翻译提示:const (
//	// 随机源标志，仅为兼容性设置，在DragonFlyBSD上无效。
//	随机源 格式随机数标志 = 0x0001
//
//	// 非阻塞模式，当无法立即返回时返回EAGAIN错误。
//	非阻塞 格式随机数标志 = 0x0002
//) 
//
//// GetRandomFlag 是一个常量类型，代表获取随机数时的标志。
const (
	// GRND_RANDOM is only set for portability purpose, no-op on DragonFlyBSD.
	GRND_RANDOM GetRandomFlag = 0x0001

	// GRND_NONBLOCK means return EAGAIN rather than blocking.
	GRND_NONBLOCK GetRandomFlag = 0x0002
)
