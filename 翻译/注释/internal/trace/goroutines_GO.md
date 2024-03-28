
<原文开始>
// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
# <翻译结束>


<原文开始>
// GDesc contains statistics and execution details of a single goroutine.
<原文结束>

# <翻译开始>
// GDesc contains statistics and execution details of a single goroutine.
# <翻译结束>


<原文开始>
// List of regions in the goroutine, sorted based on the start time.
<原文结束>

# <翻译开始>
// List of regions in the goroutine, sorted based on the start time.
# <翻译结束>


<原文开始>
// Statistics of execution time during the goroutine execution.
<原文结束>

# <翻译开始>
// Statistics of execution time during the goroutine execution.
# <翻译结束>


<原文开始>
// UserRegionDesc represents a region and goroutine execution stats
// while the region was active.
<原文结束>

# <翻译开始>
// UserRegionDesc represents a region and goroutine execution stats
// while the region was active.
# <翻译结束>


<原文开始>
	// Region start event. Normally EvUserRegion start event or nil,
	// but can be EvGoCreate event if the region is a synthetic
	// region representing task inheritance from the parent goroutine.
<原文结束>

# <翻译开始>
	// Region start event. Normally EvUserRegion start event or nil,
	// but can be EvGoCreate event if the region is a synthetic
	// region representing task inheritance from the parent goroutine.
# <翻译结束>


<原文开始>
	// Region end event. Normally EvUserRegion end event or nil,
	// but can be EvGoStop or EvGoEnd event if the goroutine
	// terminated without explicitly ending the region.
<原文结束>

# <翻译开始>
	// Region end event. Normally EvUserRegion end event or nil,
	// but can be EvGoStop or EvGoEnd event if the goroutine
	// terminated without explicitly ending the region.
# <翻译结束>


<原文开始>
// GExecutionStat contains statistics about a goroutine's execution
// during a period of time.
<原文结束>

# <翻译开始>
// GExecutionStat contains statistics about a goroutine's execution
// during a period of time.
# <翻译结束>


<原文开始>
// sub returns the stats v-s.
<原文结束>

# <翻译开始>
// sub returns the stats v-s.
# <翻译结束>


<原文开始>
// snapshotStat returns the snapshot of the goroutine execution statistics.
// This is called as we process the ordered trace event stream. lastTs and
// activeGCStartTime are used to process pending statistics if this is called
// before any goroutine end event.
<原文结束>

# <翻译开始>
// snapshotStat returns the snapshot of the goroutine execution statistics.
// This is called as we process the ordered trace event stream. lastTs and
// activeGCStartTime are used to process pending statistics if this is called
// before any goroutine end event.
# <翻译结束>


<原文开始>
// finalized GDesc. No pending state.
<原文结束>

# <翻译开始>
// finalized GDesc. No pending state.
# <翻译结束>


<原文开始>
// terminating while GC is active
<原文结束>

# <翻译开始>
// terminating while GC is active
# <翻译结束>


<原文开始>
			// The goroutine's lifetime completely overlaps
			// with a GC.
<原文结束>

# <翻译开始>
			// The goroutine's lifetime completely overlaps
			// with a GC.
# <翻译结束>


<原文开始>
// finalize is called when processing a goroutine end event or at
// the end of trace processing. This finalizes the execution stat
// and any active regions in the goroutine, in which case trigger is nil.
<原文结束>

# <翻译开始>
// finalize is called when processing a goroutine end event or at
// the end of trace processing. This finalizes the execution stat
// and any active regions in the goroutine, in which case trigger is nil.
# <翻译结束>


<原文开始>
	// System goroutines are never part of regions, even though they
	// "inherit" a task due to creation (EvGoCreate) from within a region.
	// This may happen e.g. if the first GC is triggered within a region,
	// starting the GC worker goroutines.
<原文结束>

# <翻译开始>
	// System goroutines are never part of regions, even though they
	// "inherit" a task due to creation (EvGoCreate) from within a region.
	// This may happen e.g. if the first GC is triggered within a region,
	// starting the GC worker goroutines.
# <翻译结束>


<原文开始>
// gdesc is a private part of GDesc that is required only during analysis.
<原文结束>

# <翻译开始>
// gdesc is a private part of GDesc that is required only during analysis.
# <翻译结束>


<原文开始>
// stack of active regions
<原文结束>

# <翻译开始>
// stack of active regions
# <翻译结束>


<原文开始>
// GoroutineStats generates statistics for all goroutines in the trace.
<原文结束>

# <翻译开始>
// GoroutineStats generates statistics for all goroutines in the trace.
# <翻译结束>


<原文开始>
// gcStartTime == 0 indicates gc is inactive.
<原文结束>

# <翻译开始>
// gcStartTime == 0 indicates gc is inactive.
# <翻译结束>


<原文开始>
			// When a goroutine is newly created, inherit the task
			// of the active region. For ease handling of this
			// case, we create a fake region description with the
			// task id. This isn't strictly necessary as this
			// goroutine may not be associated with the task, but
			// it can be convenient to see all children created
			// during a region.
<原文结束>

# <翻译开始>
			// When a goroutine is newly created, inherit the task
			// of the active region. For ease handling of this
			// case, we create a fake region description with the
			// task id. This isn't strictly necessary as this
			// goroutine may not be associated with the task, but
			// it can be convenient to see all children created
			// during a region.
# <翻译结束>


<原文开始>
// Sweep can happen during GC on system goroutine.
<原文结束>

# <翻译开始>
// Sweep can happen during GC on system goroutine.
# <翻译结束>


<原文开始>
// indicates gc is inactive.
<原文结束>

# <翻译开始>
// indicates gc is inactive.
# <翻译结束>


<原文开始>
// sort based on region start time
<原文结束>

# <翻译开始>
// sort based on region start time
# <翻译结束>


<原文开始>
// RelatedGoroutines finds a set of goroutines related to goroutine goid.
<原文结束>

# <翻译开始>
// RelatedGoroutines finds a set of goroutines related to goroutine goid.
# <翻译结束>


<原文开始>
	// BFS of depth 2 over "unblock" edges
	// (what goroutines unblock goroutine goid?).
<原文结束>

# <翻译开始>
	// BFS of depth 2 over "unblock" edges
	// (what goroutines unblock goroutine goid?).
# <翻译结束>


<原文开始>
	// This mimics runtime.isSystemGoroutine as closely as
	// possible.
	// Also, locked g in extra M (with empty entryFn) is system goroutine.
<原文结束>

# <翻译开始>
	// This mimics runtime.isSystemGoroutine as closely as
	// possible.
	// Also, locked g in extra M (with empty entryFn) is system goroutine.
# <翻译结束>

