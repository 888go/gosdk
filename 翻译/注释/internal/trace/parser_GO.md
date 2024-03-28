
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
// Event describes one event in the trace.
<原文结束>

# <翻译开始>
// Event describes one event in the trace.
# <翻译结束>


<原文开始>
// offset in input file (for debugging and error reporting)
<原文结束>

# <翻译开始>
// offset in input file (for debugging and error reporting)
# <翻译结束>


<原文开始>
// timestamp in nanoseconds
<原文结束>

# <翻译开始>
// timestamp in nanoseconds
# <翻译结束>


<原文开始>
// P on which the event happened (can be one of TimerP, NetpollP, SyscallP)
<原文结束>

# <翻译开始>
// P on which the event happened (can be one of TimerP, NetpollP, SyscallP)
# <翻译结束>


<原文开始>
// G on which the event happened
<原文结束>

# <翻译开始>
// G on which the event happened
# <翻译结束>


<原文开始>
// stack trace (can be empty)
<原文结束>

# <翻译开始>
// stack trace (can be empty)
# <翻译结束>


<原文开始>
// event-type-specific arguments
<原文结束>

# <翻译开始>
// event-type-specific arguments
# <翻译结束>


<原文开始>
// event-type-specific string args
<原文结束>

# <翻译开始>
// event-type-specific string args
# <翻译结束>


<原文开始>
	// linked event (can be nil), depends on event type:
	// for GCStart: the GCStop
	// for GCSTWStart: the GCSTWDone
	// for GCSweepStart: the GCSweepDone
	// for GoCreate: first GoStart of the created goroutine
	// for GoStart/GoStartLabel: the associated GoEnd, GoBlock or other blocking event
	// for GoSched/GoPreempt: the next GoStart
	// for GoBlock and other blocking events: the unblock event
	// for GoUnblock: the associated GoStart
	// for blocking GoSysCall: the associated GoSysExit
	// for GoSysExit: the next GoStart
	// for GCMarkAssistStart: the associated GCMarkAssistDone
	// for UserTaskCreate: the UserTaskEnd
	// for UserRegion: if the start region, the corresponding UserRegion end event
<原文结束>

# <翻译开始>
	// linked event (can be nil), depends on event type:
	// for GCStart: the GCStop
	// for GCSTWStart: the GCSTWDone
	// for GCSweepStart: the GCSweepDone
	// for GoCreate: first GoStart of the created goroutine
	// for GoStart/GoStartLabel: the associated GoEnd, GoBlock or other blocking event
	// for GoSched/GoPreempt: the next GoStart
	// for GoBlock and other blocking events: the unblock event
	// for GoUnblock: the associated GoStart
	// for blocking GoSysCall: the associated GoSysExit
	// for GoSysExit: the next GoStart
	// for GCMarkAssistStart: the associated GCMarkAssistDone
	// for UserTaskCreate: the UserTaskEnd
	// for UserRegion: if the start region, the corresponding UserRegion end event
# <翻译结束>


<原文开始>
// Frame is a frame in stack traces.
<原文结束>

# <翻译开始>
// Frame is a frame in stack traces.
# <翻译结束>


<原文开始>
// depicts network unblocks
<原文结束>

# <翻译开始>
// depicts network unblocks
# <翻译结束>


<原文开始>
// depicts returns from syscalls
<原文结束>

# <翻译开始>
// depicts returns from syscalls
# <翻译结束>


<原文开始>
// depicts recording of CPU profile samples
<原文结束>

# <翻译开始>
// depicts recording of CPU profile samples
# <翻译结束>


<原文开始>
// ParseResult is the result of Parse.
<原文结束>

# <翻译开始>
// ParseResult is the result of Parse.
# <翻译结束>


<原文开始>
// Events is the sorted list of Events in the trace.
<原文结束>

# <翻译开始>
// Events is the sorted list of Events in the trace.
# <翻译结束>


<原文开始>
// Stacks is the stack traces keyed by stack IDs from the trace.
<原文结束>

# <翻译开始>
// Stacks is the stack traces keyed by stack IDs from the trace.
# <翻译结束>


<原文开始>
// Parse parses, post-processes and verifies the trace.
<原文结束>

# <翻译开始>
// Parse parses, post-processes and verifies the trace.
# <翻译结束>


<原文开始>
// parse parses, post-processes and verifies the trace. It returns the
// trace version and the list of events.
<原文结束>

# <翻译开始>
// parse parses, post-processes and verifies the trace. It returns the
// trace version and the list of events.
# <翻译结束>


<原文开始>
// rawEvent is a helper type used during parsing.
<原文结束>

# <翻译开始>
// rawEvent is a helper type used during parsing.
# <翻译结束>


<原文开始>
// readTrace does wire-format parsing and verification.
// It does not care about specific event types and argument meaning.
<原文结束>

# <翻译开始>
// readTrace does wire-format parsing and verification.
// It does not care about specific event types and argument meaning.
# <翻译结束>


<原文开始>
// Read and validate trace header.
<原文结束>

# <翻译开始>
// Read and validate trace header.
# <翻译结束>


<原文开始>
		// Note: When adding a new version, confirm that canned traces from the
		// old version are part of the test suite. Add them using mkcanned.bash.
<原文结束>

# <翻译开始>
		// Note: When adding a new version, confirm that canned traces from the
		// old version are part of the test suite. Add them using mkcanned.bash.
# <翻译结束>


<原文开始>
// Read event type and number of arguments (1 byte).
<原文结束>

# <翻译开始>
// Read event type and number of arguments (1 byte).
# <翻译结束>


<原文开始>
// String dictionary entry [ID, length, string].
<原文结束>

# <翻译开始>
// String dictionary entry [ID, length, string].
# <翻译结束>


<原文开始>
// More than inlineArgs args, the first value is length of the event in bytes.
<原文结束>

# <翻译开始>
// More than inlineArgs args, the first value is length of the event in bytes.
# <翻译结束>


<原文开始>
// EvUserLog records are followed by a value string of length ev.args[len(ev.args)-1]
<原文结束>

# <翻译开始>
// EvUserLog records are followed by a value string of length ev.args[len(ev.args)-1]
# <翻译结束>


<原文开始>
// parseHeader parses trace header of the form "go 1.7 trace\x00\x00\x00\x00"
// and returns parsed version as 1007.
<原文结束>

# <翻译开始>
// parseHeader parses trace header of the form "go 1.7 trace\x00\x00\x00\x00"
// and returns parsed version as 1007.
# <翻译结束>


<原文开始>
// Parse events transforms raw events into events.
// It does analyze and verify per-event-type arguments.
<原文结束>

# <翻译开始>
// Parse events transforms raw events into events.
// It does analyze and verify per-event-type arguments.
# <翻译结束>


<原文开始>
// last goroutine running on P
<原文结束>

# <翻译开始>
// last goroutine running on P
# <翻译结束>


<原文开始>
				// The most likely cause for this is tick skew on different CPUs.
				// For example, solaris/amd64 seems to have wildly different
				// ticks on different CPUs.
<原文结束>

# <翻译开始>
				// The most likely cause for this is tick skew on different CPUs.
				// For example, solaris/amd64 seems to have wildly different
				// ticks on different CPUs.
# <翻译结束>


<原文开始>
// e.Args 0: taskID, 1:parentID, 2:nameID
<原文结束>

# <翻译开始>
// e.Args 0: taskID, 1:parentID, 2:nameID
# <翻译结束>


<原文开始>
// e.Args 0: taskID, 1: mode, 2:nameID
<原文结束>

# <翻译开始>
// e.Args 0: taskID, 1: mode, 2:nameID
# <翻译结束>


<原文开始>
// e.Args 0: taskID, 1:keyID, 2: stackID
<原文结束>

# <翻译开始>
// e.Args 0: taskID, 1:keyID, 2: stackID
# <翻译结束>


<原文开始>
				// Most events are written out by the active P at the exact
				// moment they describe. CPU profile samples are different
				// because they're written to the tracing log after some delay,
				// by a separate worker goroutine, into a separate buffer.
				//
				// We keep these in their own batch until all of the batches are
				// merged in timestamp order. We also (right before the merge)
				// re-sort these events by the timestamp captured in the
				// profiling signal handler.
<原文结束>

# <翻译开始>
				// Most events are written out by the active P at the exact
				// moment they describe. CPU profile samples are different
				// because they're written to the tracing log after some delay,
				// by a separate worker goroutine, into a separate buffer.
				//
				// We keep these in their own batch until all of the batches are
				// merged in timestamp order. We also (right before the merge)
				// re-sort these events by the timestamp captured in the
				// profiling signal handler.
# <翻译结束>


<原文开始>
// Translate cpu ticks to real time.
<原文结束>

# <翻译开始>
// Translate cpu ticks to real time.
# <翻译结束>


<原文开始>
// Use floating point to avoid integer overflows.
<原文结束>

# <翻译开始>
// Use floating point to avoid integer overflows.
# <翻译结束>


<原文开始>
// Move timers and syscalls to separate fake Ps.
<原文结束>

# <翻译开始>
// Move timers and syscalls to separate fake Ps.
# <翻译结束>


<原文开始>
// removeFutile removes all constituents of futile wakeups (block, unblock, start).
// For example, a goroutine was unblocked on a mutex, but another goroutine got
// ahead and acquired the mutex before the first goroutine is scheduled,
// so the first goroutine has to block again. Such wakeups happen on buffered
// channels and sync.Mutex, but are generally not interesting for end user.
<原文结束>

# <翻译开始>
// removeFutile removes all constituents of futile wakeups (block, unblock, start).
// For example, a goroutine was unblocked on a mutex, but another goroutine got
// ahead and acquired the mutex before the first goroutine is scheduled,
// so the first goroutine has to block again. Such wakeups happen on buffered
// channels and sync.Mutex, but are generally not interesting for end user.
# <翻译结束>


<原文开始>
	// Two non-trivial aspects:
	// 1. A goroutine can be preempted during a futile wakeup and migrate to another P.
	//	We want to remove all of that.
	// 2. Tracing can start in the middle of a futile wakeup.
	//	That is, we can see a futile wakeup event w/o the actual wakeup before it.
	// postProcessTrace runs after us and ensures that we leave the trace in a consistent state.
<原文结束>

# <翻译开始>
	// Two non-trivial aspects:
	// 1. A goroutine can be preempted during a futile wakeup and migrate to another P.
	//	We want to remove all of that.
	// 2. Tracing can start in the middle of a futile wakeup.
	//	That is, we can see a futile wakeup event w/o the actual wakeup before it.
	// postProcessTrace runs after us and ensures that we leave the trace in a consistent state.
# <翻译结束>


<原文开始>
// Phase 1: determine futile wakeup sequences.
<原文结束>

# <翻译开始>
// Phase 1: determine futile wakeup sequences.
# <翻译结束>


<原文开始>
// wakeup sequence (subject for removal)
<原文结束>

# <翻译开始>
// wakeup sequence (subject for removal)
# <翻译结束>


<原文开始>
// Phase 2: remove futile wakeup sequences.
<原文结束>

# <翻译开始>
// Phase 2: remove futile wakeup sequences.
# <翻译结束>


<原文开始>
// overwrite the original slice
<原文结束>

# <翻译开始>
// overwrite the original slice
# <翻译结束>


<原文开始>
// ErrTimeOrder is returned by Parse when the trace contains
// time stamps that do not respect actual event ordering.
<原文结束>

# <翻译开始>
// ErrTimeOrder is returned by Parse when the trace contains
// time stamps that do not respect actual event ordering.
# <翻译结束>


<原文开始>
// postProcessTrace does inter-event verification and information restoration.
// The resulting trace is guaranteed to be consistent
// (for example, a P does not run two Gs at the same time, or a G is indeed
// blocked before an unblock event).
<原文结束>

# <翻译开始>
// postProcessTrace does inter-event verification and information restoration.
// The resulting trace is guaranteed to be consistent
// (for example, a P does not run two Gs at the same time, or a G is indeed
// blocked before an unblock event).
# <翻译结束>


<原文开始>
// task id to task creation events
<原文结束>

# <翻译开始>
// task id to task creation events
# <翻译结束>


<原文开始>
// goroutine id to stack of regions
<原文结束>

# <翻译开始>
// goroutine id to stack of regions
# <翻译结束>


<原文开始>
// Attribute this to the global GC state.
<原文结束>

# <翻译开始>
// Attribute this to the global GC state.
# <翻译结束>


<原文开始>
// Before 1.10, EvGCSTWStart was per-P.
<原文结束>

# <翻译开始>
// Before 1.10, EvGCSTWStart was per-P.
# <翻译结束>


<原文开始>
// Before 1.10, EvGCSTWDone was per-P.
<原文结束>

# <翻译开始>
// Before 1.10, EvGCSTWDone was per-P.
# <翻译结束>


<原文开始>
			// Unlike most events, mark assists can be in progress when a
			// goroutine starts tracing, so we can't report an error here.
<原文结束>

# <翻译开始>
			// Unlike most events, mark assists can be in progress when a
			// goroutine starts tracing, so we can't report an error here.
# <翻译结束>


<原文开始>
// +1 because symbolizer expects return pc.
<原文结束>

# <翻译开始>
// +1 because symbolizer expects return pc.
# <翻译结束>


<原文开始>
// flush all active regions
<原文结束>

# <翻译开始>
// flush all active regions
# <翻译结束>


<原文开始>
// matching region start event is in the trace.
<原文结束>

# <翻译开始>
// matching region start event is in the trace.
# <翻译结束>


<原文开始>
// task id, region name mismatch
<原文结束>

# <翻译开始>
// task id, region name mismatch
# <翻译结束>


<原文开始>
// Link region start event with span end event
<原文结束>

# <翻译开始>
// Link region start event with span end event
# <翻译结束>


<原文开始>
	// TODO(dvyukov): restore stacks for EvGoStart events.
	// TODO(dvyukov): test that all EvGoStart events has non-nil Link.
<原文结束>

# <翻译开始>
	// TODO(dvyukov): restore stacks for EvGoStart events.
	// TODO(dvyukov): test that all EvGoStart events has non-nil Link.
# <翻译结束>


<原文开始>
// symbolize attaches func/file/line info to stack traces.
<原文结束>

# <翻译开始>
// symbolize attaches func/file/line info to stack traces.
# <翻译结束>


<原文开始>
// First, collect and dedup all pcs.
<原文结束>

# <翻译开始>
// First, collect and dedup all pcs.
# <翻译结束>


<原文开始>
	// Write all pcs to addr2line.
	// Need to copy pcs to an array, because map iteration order is non-deterministic.
<原文结束>

# <翻译开始>
	// Write all pcs to addr2line.
	// Need to copy pcs to an array, because map iteration order is non-deterministic.
# <翻译结束>


<原文开始>
// Replace frames in events array.
<原文结束>

# <翻译开始>
// Replace frames in events array.
# <翻译结束>


<原文开始>
// readVal reads unsigned base-128 value from r.
<原文结束>

# <翻译开始>
// readVal reads unsigned base-128 value from r.
# <翻译结束>


<原文开始>
// Print dumps events to stdout. For debugging.
<原文结束>

# <翻译开始>
// Print dumps events to stdout. For debugging.
# <翻译结束>


<原文开始>
// PrintEvent dumps the event to stdout. For debugging.
<原文结束>

# <翻译开始>
// PrintEvent dumps the event to stdout. For debugging.
# <翻译结束>


<原文开始>
// argNum returns total number of args for the event accounting for timestamps,
// sequence numbers and differences between trace format versions.
<原文结束>

# <翻译开始>
// argNum returns total number of args for the event accounting for timestamps,
// sequence numbers and differences between trace format versions.
# <翻译结束>


<原文开始>
// there was an unused arg before 1.7
<原文结束>

# <翻译开始>
// there was an unused arg before 1.7
# <翻译结束>


<原文开始>
// 1.9 added two arguments
<原文结束>

# <翻译开始>
// 1.9 added two arguments
# <翻译结束>


<原文开始>
// 1.7 added an additional seq arg
<原文结束>

# <翻译开始>
// 1.7 added an additional seq arg
# <翻译结束>


<原文开始>
// BreakTimestampsForTesting causes the parser to randomly alter timestamps (for testing of broken cputicks).
<原文结束>

# <翻译开始>
// BreakTimestampsForTesting causes the parser to randomly alter timestamps (for testing of broken cputicks).
# <翻译结束>


<原文开始>
// Event types in the trace.
// Verbatim copy from src/runtime/trace.go with the "trace" prefix removed.
<原文结束>

# <翻译开始>
// Event types in the trace.
// Verbatim copy from src/runtime/trace.go with the "trace" prefix removed.
# <翻译结束>


<原文开始>
// start of per-P batch of events [pid, timestamp]
<原文结束>

# <翻译开始>
// start of per-P batch of events [pid, timestamp]
# <翻译结束>


<原文开始>
// contains tracer timer frequency [frequency (ticks per second)]
<原文结束>

# <翻译开始>
// contains tracer timer frequency [frequency (ticks per second)]
# <翻译结束>


<原文开始>
// stack [stack id, number of PCs, array of {PC, func string ID, file string ID, line}]
<原文结束>

# <翻译开始>
// stack [stack id, number of PCs, array of {PC, func string ID, file string ID, line}]
# <翻译结束>


<原文开始>
// current value of GOMAXPROCS [timestamp, GOMAXPROCS, stack id]
<原文结束>

# <翻译开始>
// current value of GOMAXPROCS [timestamp, GOMAXPROCS, stack id]
# <翻译结束>


<原文开始>
// start of P [timestamp, thread id]
<原文结束>

# <翻译开始>
// start of P [timestamp, thread id]
# <翻译结束>


<原文开始>
// GC start [timestamp, seq, stack id]
<原文结束>

# <翻译开始>
// GC start [timestamp, seq, stack id]
# <翻译结束>


<原文开始>
// GC mark termination start [timestamp, kind]
<原文结束>

# <翻译开始>
// GC mark termination start [timestamp, kind]
# <翻译结束>


<原文开始>
// GC mark termination done [timestamp]
<原文结束>

# <翻译开始>
// GC mark termination done [timestamp]
# <翻译结束>


<原文开始>
// GC sweep start [timestamp, stack id]
<原文结束>

# <翻译开始>
// GC sweep start [timestamp, stack id]
# <翻译结束>


<原文开始>
// GC sweep done [timestamp, swept, reclaimed]
<原文结束>

# <翻译开始>
// GC sweep done [timestamp, swept, reclaimed]
# <翻译结束>


<原文开始>
// goroutine creation [timestamp, new goroutine id, new stack id, stack id]
<原文结束>

# <翻译开始>
// goroutine creation [timestamp, new goroutine id, new stack id, stack id]
# <翻译结束>


<原文开始>
// goroutine starts running [timestamp, goroutine id, seq]
<原文结束>

# <翻译开始>
// goroutine starts running [timestamp, goroutine id, seq]
# <翻译结束>


<原文开始>
// goroutine ends [timestamp]
<原文结束>

# <翻译开始>
// goroutine ends [timestamp]
# <翻译结束>


<原文开始>
// goroutine stops (like in select{}) [timestamp, stack]
<原文结束>

# <翻译开始>
// goroutine stops (like in select{}) [timestamp, stack]
# <翻译结束>


<原文开始>
// goroutine calls Gosched [timestamp, stack]
<原文结束>

# <翻译开始>
// goroutine calls Gosched [timestamp, stack]
# <翻译结束>


<原文开始>
// goroutine is preempted [timestamp, stack]
<原文结束>

# <翻译开始>
// goroutine is preempted [timestamp, stack]
# <翻译结束>


<原文开始>
// goroutine calls Sleep [timestamp, stack]
<原文结束>

# <翻译开始>
// goroutine calls Sleep [timestamp, stack]
# <翻译结束>


<原文开始>
// goroutine blocks [timestamp, stack]
<原文结束>

# <翻译开始>
// goroutine blocks [timestamp, stack]
# <翻译结束>


<原文开始>
// goroutine is unblocked [timestamp, goroutine id, seq, stack]
<原文结束>

# <翻译开始>
// goroutine is unblocked [timestamp, goroutine id, seq, stack]
# <翻译结束>


<原文开始>
// goroutine blocks on chan send [timestamp, stack]
<原文结束>

# <翻译开始>
// goroutine blocks on chan send [timestamp, stack]
# <翻译结束>


<原文开始>
// goroutine blocks on chan recv [timestamp, stack]
<原文结束>

# <翻译开始>
// goroutine blocks on chan recv [timestamp, stack]
# <翻译结束>


<原文开始>
// goroutine blocks on select [timestamp, stack]
<原文结束>

# <翻译开始>
// goroutine blocks on select [timestamp, stack]
# <翻译结束>


<原文开始>
// goroutine blocks on Mutex/RWMutex [timestamp, stack]
<原文结束>

# <翻译开始>
// goroutine blocks on Mutex/RWMutex [timestamp, stack]
# <翻译结束>


<原文开始>
// goroutine blocks on Cond [timestamp, stack]
<原文结束>

# <翻译开始>
// goroutine blocks on Cond [timestamp, stack]
# <翻译结束>


<原文开始>
// goroutine blocks on network [timestamp, stack]
<原文结束>

# <翻译开始>
// goroutine blocks on network [timestamp, stack]
# <翻译结束>


<原文开始>
// syscall enter [timestamp, stack]
<原文结束>

# <翻译开始>
// syscall enter [timestamp, stack]
# <翻译结束>


<原文开始>
// syscall exit [timestamp, goroutine id, seq, real timestamp]
<原文结束>

# <翻译开始>
// syscall exit [timestamp, goroutine id, seq, real timestamp]
# <翻译结束>


<原文开始>
// syscall blocks [timestamp]
<原文结束>

# <翻译开始>
// syscall blocks [timestamp]
# <翻译结束>


<原文开始>
// denotes that goroutine is blocked when tracing starts [timestamp, goroutine id]
<原文结束>

# <翻译开始>
// denotes that goroutine is blocked when tracing starts [timestamp, goroutine id]
# <翻译结束>


<原文开始>
// denotes that goroutine is in syscall when tracing starts [timestamp, goroutine id]
<原文结束>

# <翻译开始>
// denotes that goroutine is in syscall when tracing starts [timestamp, goroutine id]
# <翻译结束>


<原文开始>
// gcController.heapLive change [timestamp, heap live bytes]
<原文结束>

# <翻译开始>
// gcController.heapLive change [timestamp, heap live bytes]
# <翻译结束>


<原文开始>
// gcController.heapGoal change [timestamp, heap goal bytes]
<原文结束>

# <翻译开始>
// gcController.heapGoal change [timestamp, heap goal bytes]
# <翻译结束>


<原文开始>
// denotes timer goroutine [timer goroutine id]
<原文结束>

# <翻译开始>
// denotes timer goroutine [timer goroutine id]
# <翻译结束>


<原文开始>
// denotes that the previous wakeup of this goroutine was futile [timestamp]
<原文结束>

# <翻译开始>
// denotes that the previous wakeup of this goroutine was futile [timestamp]
# <翻译结束>


<原文开始>
// string dictionary entry [ID, length, string]
<原文结束>

# <翻译开始>
// string dictionary entry [ID, length, string]
# <翻译结束>


<原文开始>
// goroutine starts running on the same P as the last event [timestamp, goroutine id]
<原文结束>

# <翻译开始>
// goroutine starts running on the same P as the last event [timestamp, goroutine id]
# <翻译结束>


<原文开始>
// goroutine is unblocked on the same P as the last event [timestamp, goroutine id, stack]
<原文结束>

# <翻译开始>
// goroutine is unblocked on the same P as the last event [timestamp, goroutine id, stack]
# <翻译结束>


<原文开始>
// syscall exit on the same P as the last event [timestamp, goroutine id, real timestamp]
<原文结束>

# <翻译开始>
// syscall exit on the same P as the last event [timestamp, goroutine id, real timestamp]
# <翻译结束>


<原文开始>
// goroutine starts running with label [timestamp, goroutine id, seq, label string id]
<原文结束>

# <翻译开始>
// goroutine starts running with label [timestamp, goroutine id, seq, label string id]
# <翻译结束>


<原文开始>
// goroutine blocks on GC assist [timestamp, stack]
<原文结束>

# <翻译开始>
// goroutine blocks on GC assist [timestamp, stack]
# <翻译结束>


<原文开始>
// GC mark assist start [timestamp, stack]
<原文结束>

# <翻译开始>
// GC mark assist start [timestamp, stack]
# <翻译结束>


<原文开始>
// GC mark assist done [timestamp]
<原文结束>

# <翻译开始>
// GC mark assist done [timestamp]
# <翻译结束>


<原文开始>
// trace.NewContext [timestamp, internal task id, internal parent id, stack, name string]
<原文结束>

# <翻译开始>
// trace.NewContext [timestamp, internal task id, internal parent id, stack, name string]
# <翻译结束>


<原文开始>
// end of task [timestamp, internal task id, stack]
<原文结束>

# <翻译开始>
// end of task [timestamp, internal task id, stack]
# <翻译结束>


<原文开始>
// trace.WithRegion [timestamp, internal task id, mode(0:start, 1:end), stack, name string]
<原文结束>

# <翻译开始>
// trace.WithRegion [timestamp, internal task id, mode(0:start, 1:end), stack, name string]
# <翻译结束>


<原文开始>
// trace.Log [timestamp, internal id, key string id, stack, value string]
<原文结束>

# <翻译开始>
// trace.Log [timestamp, internal id, key string id, stack, value string]
# <翻译结束>


<原文开始>
// CPU profiling sample [timestamp, stack, real timestamp, real P id (-1 when absent), goroutine id]
<原文结束>

# <翻译开始>
// CPU profiling sample [timestamp, stack, real timestamp, real P id (-1 when absent), goroutine id]
# <翻译结束>


<原文开始>
// in 1.5 format it was {"p", "seq", "ticks"}
<原文结束>

# <翻译开始>
// in 1.5 format it was {"p", "seq", "ticks"}
# <翻译结束>


<原文开始>
// in 1.5 format it was {"freq", "unused"}
<原文结束>

# <翻译开始>
// in 1.5 format it was {"freq", "unused"}
# <翻译结束>


<原文开始>
// in 1.5 format it was {}
<原文结束>

# <翻译开始>
// in 1.5 format it was {}
# <翻译结束>


<原文开始>
// <= 1.9, args was {} (implicitly {0})
<原文结束>

# <翻译开始>
// <= 1.9, args was {} (implicitly {0})
# <翻译结束>


<原文开始>
// before 1.9, format was {}
<原文结束>

# <翻译开始>
// before 1.9, format was {}
# <翻译结束>


<原文开始>
// in 1.5 format it was {"g"}
<原文结束>

# <翻译开始>
// in 1.5 format it was {"g"}
# <翻译结束>


<原文开始>
// in 1.5 format it was {"g", "unused"}
<原文结束>

# <翻译开始>
// in 1.5 format it was {"g", "unused"}
# <翻译结束>

