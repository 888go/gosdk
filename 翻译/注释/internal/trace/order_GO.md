
<原文开始>
// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
# <翻译结束>


<原文开始>
// order1007 merges a set of per-P event batches into a single, consistent stream.
// The high level idea is as follows. Events within an individual batch are in
// correct order, because they are emitted by a single P. So we need to produce
// a correct interleaving of the batches. To do this we take first unmerged event
// from each batch (frontier). Then choose subset that is "ready" to be merged,
// that is, events for which all dependencies are already merged. Then we choose
// event with the lowest timestamp from the subset, merge it and repeat.
// This approach ensures that we form a consistent stream even if timestamps are
// incorrect (condition observed on some machines).
<原文结束>

# <翻译开始>
// order1007 merges a set of per-P event batches into a single, consistent stream.
// The high level idea is as follows. Events within an individual batch are in
// correct order, because they are emitted by a single P. So we need to produce
// a correct interleaving of the batches. To do this we take first unmerged event
// from each batch (frontier). Then choose subset that is "ready" to be merged,
// that is, events for which all dependencies are already merged. Then we choose
// event with the lowest timestamp from the subset, merge it and repeat.
// This approach ensures that we form a consistent stream even if timestamps are
// incorrect (condition observed on some machines).
# <翻译结束>


<原文开始>
	// The ordering of CPU profile sample events in the data stream is based on
	// when each run of the signal handler was able to acquire the spinlock,
	// with original timestamps corresponding to when ReadTrace pulled the data
	// off of the profBuf queue. Re-sort them by the timestamp we captured
	// inside the signal handler.
<原文结束>

# <翻译开始>
	// The ordering of CPU profile sample events in the data stream is based on
	// when each run of the signal handler was able to acquire the spinlock,
	// with original timestamps corresponding to when ReadTrace pulled the data
	// off of the profBuf queue. Re-sort them by the timestamp we captured
	// inside the signal handler.
# <翻译结束>


<原文开始>
// Get rid of "Local" events, they are intended merely for ordering.
<原文结束>

# <翻译开始>
// Get rid of "Local" events, they are intended merely for ordering.
# <翻译结束>


<原文开始>
	// At this point we have a consistent stream of events.
	// Make sure time stamps respect the ordering.
	// The tests will skip (not fail) the test case if they see this error.
<原文结束>

# <翻译开始>
	// At this point we have a consistent stream of events.
	// Make sure time stamps respect the ordering.
	// The tests will skip (not fail) the test case if they see this error.
# <翻译结束>


<原文开始>
	// The last part is giving correct timestamps to EvGoSysExit events.
	// The problem with EvGoSysExit is that actual syscall exit timestamp (ev.Args[2])
	// is potentially acquired long before event emission. So far we've used
	// timestamp of event emission (ev.Ts).
	// We could not set ev.Ts = ev.Args[2] earlier, because it would produce
	// seemingly broken timestamps (misplaced event).
	// We also can't simply update the timestamp and resort events, because
	// if timestamps are broken we will misplace the event and later report
	// logically broken trace (instead of reporting broken timestamps).
<原文结束>

# <翻译开始>
	// The last part is giving correct timestamps to EvGoSysExit events.
	// The problem with EvGoSysExit is that actual syscall exit timestamp (ev.Args[2])
	// is potentially acquired long before event emission. So far we've used
	// timestamp of event emission (ev.Ts).
	// We could not set ev.Ts = ev.Args[2] earlier, because it would produce
	// seemingly broken timestamps (misplaced event).
	// We also can't simply update the timestamp and resort events, because
	// if timestamps are broken we will misplace the event and later report
	// logically broken trace (instead of reporting broken timestamps).
# <翻译结束>


<原文开始>
// stateTransition returns goroutine state (sequence and status) when the event
// becomes ready for merging (init) and the goroutine state after the event (next).
<原文结束>

# <翻译开始>
// stateTransition returns goroutine state (sequence and status) when the event
// becomes ready for merging (init) and the goroutine state after the event (next).
# <翻译结束>


<原文开始>
		// noseq means that this event is ready for merging as soon as
		// frontier reaches it (EvGoStartLocal is emitted on the same P
		// as the corresponding EvGoCreate/EvGoUnblock, and thus the latter
		// is already merged).
		// seqinc is a stub for cases when event increments g sequence,
		// but since we don't know current seq we also don't know next seq.
<原文结束>

# <翻译开始>
		// noseq means that this event is ready for merging as soon as
		// frontier reaches it (EvGoStartLocal is emitted on the same P
		// as the corresponding EvGoCreate/EvGoUnblock, and thus the latter
		// is already merged).
		// seqinc is a stub for cases when event increments g sequence,
		// but since we don't know current seq we also don't know next seq.
# <翻译结束>


<原文开始>
// no ordering requirements
<原文结束>

# <翻译开始>
// no ordering requirements
# <翻译结束>


<原文开始>
// order1005 merges a set of per-P event batches into a single, consistent stream.
<原文结束>

# <翻译开始>
// order1005 merges a set of per-P event batches into a single, consistent stream.
# <翻译结束>


<原文开始>
			// EvGoSysExit emission is delayed until the thread has a P.
			// Give it the real sequence number and time stamp.
<原文结束>

# <翻译开始>
			// EvGoSysExit emission is delayed until the thread has a P.
			// Give it the real sequence number and time stamp.
# <翻译结束>

