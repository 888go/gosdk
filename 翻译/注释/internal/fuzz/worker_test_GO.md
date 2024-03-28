
<原文开始>
// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
# <翻译结束>


<原文开始>
// BenchmarkWorkerPing acts as the coordinator and measures the time it takes
// a worker to respond to N pings. This is a rough measure of our RPC latency.
<原文结束>

# <翻译开始>
// BenchmarkWorkerPing acts as the coordinator and measures the time it takes
// a worker to respond to N pings. This is a rough measure of our RPC latency.
# <翻译结束>


<原文开始>
// BenchmarkWorkerFuzz acts as the coordinator and measures the time it takes
// a worker to mutate a given input and call a trivial fuzz function N times.
<原文结束>

# <翻译开始>
// BenchmarkWorkerFuzz acts as the coordinator and measures the time it takes
// a worker to mutate a given input and call a trivial fuzz function N times.
# <翻译结束>


<原文开始>
// newWorkerForTest creates and starts a worker process for testing or
// benchmarking. The worker process calls RunFuzzWorker, which responds to
// RPC messages until it's stopped. The process is stopped and cleaned up
// automatically when the test is done.
<原文结束>

# <翻译开始>
// newWorkerForTest creates and starts a worker process for testing or
// benchmarking. The worker process calls RunFuzzWorker, which responds to
// RPC messages until it's stopped. The process is stopped and cleaned up
// automatically when the test is done.
# <翻译结束>

