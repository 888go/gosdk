
<原文开始>
// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// 版权所有 2011 Go 语言作者。保留所有权利。
// 使用此源代码须遵循 BSD 风格的许可协议，
// 许可协议具体内容可在 LICENSE 文件中找到。
// md5:b5bcbe8a534f6077
# <翻译结束>


<原文开始>
// Force US/Pacific for time zone tests.
<原文结束>

# <翻译开始>
// 强制使用美国/太平洋时区进行时区测试。. md5:46c4cbaf6eb4451c
# <翻译结束>


<原文开始>
	// For hermeticity, use only tzinfo source from the test's GOROOT,
	// not the system sources and not whatever GOROOT may happen to be
	// set in the process's environment (if any).
	// This test runs in GOROOT/src/time, so GOROOT is "../..",
	// but it is theoretically possible
<原文结束>

# <翻译开始>
	// 为了保证测试的封闭性，只使用测试的GOROOT目录下的tzinfo源，
	// 而不是系统的源代码，也不使用进程中可能设置的任何GOROOT环境变量。
	// 这个测试在GOROOT/src/time目录下运行，所以GOROOT是"../.."，
	// 但理论上可能存在其他情况。
	// md5:9df2aa7dd1c5efc8
# <翻译结束>


<原文开始>
// Test that a runtimeTimer with a period that would overflow when on
// expiration does not throw or cause other timers to hang.
//
// This test has to be in internal_test.go since it fiddles with
// unexported data structures.
<原文结束>

# <翻译开始>
// 测试当运行时计时器的周期在过期时会溢出，但不会抛出异常或导致其他计时器挂起。
//
// 这个测试必须放在internal_test.go中，因为它操作未公开的数据结构。
// md5:6a4e803f3b22aef2
# <翻译结束>


<原文开始>
	// We manually create a runtimeTimer with huge period, but that expires
	// immediately. The public Timer interface would require waiting for
	// the entire period before the first update.
<原文结束>

# <翻译开始>
	// 我们手动创建一个具有巨大周期的runtimeTimer，但该定时器会立即到期。公共Timer接口需要在第一次更新之前等待整个周期。
	// md5:a4249b95fbeedfc0
# <翻译结束>


<原文开始>
	// If this test fails, we will either throw (when siftdownTimer detects
	// bad when on update), or other timers will hang (if the timer in a
	// heap is in a bad state). There is no reliable way to test this, but
	// we wait on a short timer here as a smoke test (alternatively, timers
	// in later tests may hang).
<原文结束>

# <翻译开始>
	// 如果此测试失败，我们将会抛出错误（当siftdownTimer在更新时检测到bad状态时），或者其它定时器将会挂起（如果堆中的定时器处于不良状态）。没有可靠的方法来测试这一点，但我们会在这里等待一个短时间的定时器作为基本功能测试（或者，后续测试中的定时器可能会挂起）。
	// md5:da77407212bfa69a
# <翻译结束>

