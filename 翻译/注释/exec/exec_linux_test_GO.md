
<原文开始>
// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// 版权所有 2020 Go作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。
# <翻译结束>


<原文开始>
// On systems that use glibc, calling malloc can create a new arena,
// and creating a new arena can read /sys/devices/system/cpu/online.
// If we are using cgo, we will call malloc when creating a new thread.
// That can break TestExtraFiles if we create a new thread that creates
// a new arena and opens the /sys file while we are checking for open
// file descriptors. Work around the problem by creating threads up front.
// See issue 25628.
<原文结束>

# <翻译开始>
// 在使用glibc的系统中，调用malloc可能会创建一个新的内存区域（arena），
// 而创建新的内存区域可能会读取/sys/devices/system/cpu/online文件。
// 如果我们使用cgo，在创建新线程时会调用malloc。
// 这样在我们在检查打开的文件描述符时，如果创建的新线程恰好创建了新的内存区域并打开了/sys文件，就可能破坏TestExtraFiles测试。
// 为了解决这个问题，提前创建线程作为一种解决办法。
// 参见issue 25628。
# <翻译结束>


<原文开始>
	// Start some threads. 10 is arbitrary but intended to be enough
	// to ensure that the code won't have to create any threads itself.
	// In particular this should be more than the number of threads
	// the garbage collector might create.
<原文结束>

# <翻译开始>
	// 启动一些线程。10 是一个随意选择的数值，但目的是确保代码本身不需要创建任何线程。
	// 特别地，这个数值应该足够大，以确保超过垃圾回收器可能创建的线程数量。
# <翻译结束>

