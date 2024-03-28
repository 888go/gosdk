
<原文开始>
// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
# <翻译结束>


<原文开始>
// MutatorUtil is a change in mutator utilization at a particular
// time. Mutator utilization functions are represented as a
// time-ordered []MutatorUtil.
<原文结束>

# <翻译开始>
// MutatorUtil is a change in mutator utilization at a particular
// time. Mutator utilization functions are represented as a
// time-ordered []MutatorUtil.
# <翻译结束>


<原文开始>
	// Util is the mean mutator utilization starting at Time. This
	// is in the range [0, 1].
<原文结束>

# <翻译开始>
	// Util is the mean mutator utilization starting at Time. This
	// is in the range [0, 1].
# <翻译结束>


<原文开始>
// UtilFlags controls the behavior of MutatorUtilization.
<原文结束>

# <翻译开始>
// UtilFlags controls the behavior of MutatorUtilization.
# <翻译结束>


<原文开始>
// UtilSTW means utilization should account for STW events.
<原文结束>

# <翻译开始>
// UtilSTW means utilization should account for STW events.
# <翻译结束>


<原文开始>
	// UtilBackground means utilization should account for
	// background mark workers.
<原文结束>

# <翻译开始>
	// UtilBackground means utilization should account for
	// background mark workers.
# <翻译结束>


<原文开始>
	// UtilAssist means utilization should account for mark
	// assists.
<原文结束>

# <翻译开始>
	// UtilAssist means utilization should account for mark
	// assists.
# <翻译结束>


<原文开始>
// UtilSweep means utilization should account for sweeping.
<原文结束>

# <翻译开始>
// UtilSweep means utilization should account for sweeping.
# <翻译结束>


<原文开始>
	// UtilPerProc means each P should be given a separate
	// utilization function. Otherwise, there is a single function
	// and each P is given a fraction of the utilization.
<原文结束>

# <翻译开始>
	// UtilPerProc means each P should be given a separate
	// utilization function. Otherwise, there is a single function
	// and each P is given a fraction of the utilization.
# <翻译结束>


<原文开始>
// MutatorUtilization returns a set of mutator utilization functions
// for the given trace. Each function will always end with 0
// utilization. The bounds of each function are implicit in the first
// and last event; outside of these bounds each function is undefined.
//
// If the UtilPerProc flag is not given, this always returns a single
// utilization function. Otherwise, it returns one function per P.
<原文结束>

# <翻译开始>
// MutatorUtilization returns a set of mutator utilization functions
// for the given trace. Each function will always end with 0
// utilization. The bounds of each function are implicit in the first
// and last event; outside of these bounds each function is undefined.
//
// If the UtilPerProc flag is not given, this always returns a single
// utilization function. Otherwise, it returns one function per P.
# <翻译结束>


<原文开始>
// gc > 0 indicates that GC is active on this P.
<原文结束>

# <翻译开始>
// gc > 0 indicates that GC is active on this P.
# <翻译结束>


<原文开始>
		// series the logical series number for this P. This
		// is necessary because Ps may be removed and then
		// re-added, and then the new P needs a new series.
<原文结束>

# <翻译开始>
		// series the logical series number for this P. This
		// is necessary because Ps may be removed and then
		// re-added, and then the new P needs a new series.
# <翻译结束>


<原文开始>
				// Background mark worker.
				//
				// If we're in per-proc mode, we don't
				// count dedicated workers because
				// they kick all of the goroutines off
				// that P, so don't directly
				// contribute to goroutine latency.
<原文结束>

# <翻译开始>
				// Background mark worker.
				//
				// If we're in per-proc mode, we don't
				// count dedicated workers because
				// they kick all of the goroutines off
				// that P, so don't directly
				// contribute to goroutine latency.
# <翻译结束>


<原文开始>
// Unblocked during assist.
<原文结束>

# <翻译开始>
// Unblocked during assist.
# <翻译结束>


<原文开始>
// Background mark worker done.
<原文结束>

# <翻译开始>
// Background mark worker done.
# <翻译结束>


<原文开始>
// Compute the current average utilization.
<原文结束>

# <翻译开始>
// Compute the current average utilization.
# <翻译结束>


<原文开始>
			// Record the utilization change. (Since
			// len(ps) == len(out), we know len(out) > 0.)
<原文结束>

# <翻译开始>
			// Record the utilization change. (Since
			// len(ps) == len(out), we know len(out) > 0.)
# <翻译结束>


<原文开始>
// Check for per-P utilization changes.
<原文结束>

# <翻译开始>
// Check for per-P utilization changes.
# <翻译结束>


<原文开始>
	// Add final 0 utilization event to any remaining series. This
	// is important to mark the end of the trace. The exact value
	// shouldn't matter since no window should extend beyond this,
	// but using 0 is symmetric with the start of the trace.
<原文结束>

# <翻译开始>
	// Add final 0 utilization event to any remaining series. This
	// is important to mark the end of the trace. The exact value
	// shouldn't matter since no window should extend beyond this,
	// but using 0 is symmetric with the start of the trace.
# <翻译结束>


<原文开始>
// Take the lowest utilization at a time stamp.
<原文结束>

# <翻译开始>
// Take the lowest utilization at a time stamp.
# <翻译结束>


<原文开始>
// totalUtil is total utilization, measured in nanoseconds. This is a
// separate type primarily to distinguish it from mean utilization,
// which is also a float64.
<原文结束>

# <翻译开始>
// totalUtil is total utilization, measured in nanoseconds. This is a
// separate type primarily to distinguish it from mean utilization,
// which is also a float64.
# <翻译结束>


<原文开始>
// mean returns the mean utilization over dur.
<原文结束>

# <翻译开始>
// mean returns the mean utilization over dur.
# <翻译结束>


<原文开始>
// An MMUCurve is the minimum mutator utilization curve across
// multiple window sizes.
<原文结束>

# <翻译开始>
// An MMUCurve is the minimum mutator utilization curve across
// multiple window sizes.
# <翻译结束>


<原文开始>
// sums[j] is the cumulative sum of util[:j].
<原文结束>

# <翻译开始>
// sums[j] is the cumulative sum of util[:j].
# <翻译结束>


<原文开始>
	// bands summarizes util in non-overlapping bands of duration
	// bandDur.
<原文结束>

# <翻译开始>
	// bands summarizes util in non-overlapping bands of duration
	// bandDur.
# <翻译结束>


<原文开始>
// bandDur is the duration of each band.
<原文结束>

# <翻译开始>
// bandDur is the duration of each band.
# <翻译结束>


<原文开始>
	// minUtil is the minimum instantaneous mutator utilization in
	// this band.
<原文结束>

# <翻译开始>
	// minUtil is the minimum instantaneous mutator utilization in
	// this band.
# <翻译结束>


<原文开始>
	// cumUtil is the cumulative total mutator utilization between
	// time 0 and the left edge of this band.
<原文结束>

# <翻译开始>
	// cumUtil is the cumulative total mutator utilization between
	// time 0 and the left edge of this band.
# <翻译结束>


<原文开始>
	// integrator is the integrator for the left edge of this
	// band.
<原文结束>

# <翻译开始>
	// integrator is the integrator for the left edge of this
	// band.
# <翻译结束>


<原文开始>
// NewMMUCurve returns an MMU curve for the given mutator utilization
// function.
<原文结束>

# <翻译开始>
// NewMMUCurve returns an MMU curve for the given mutator utilization
// function.
# <翻译结束>


<原文开始>
// bandsPerSeries is the number of bands to divide each series into.
// This is only changed by tests.
<原文结束>

# <翻译开始>
// bandsPerSeries is the number of bands to divide each series into.
// This is only changed by tests.
# <翻译结束>


<原文开始>
// Compute cumulative sum.
<原文结束>

# <翻译开始>
// Compute cumulative sum.
# <翻译结束>


<原文开始>
	// Divide the utilization curve up into equal size
	// non-overlapping "bands" and compute a summary for each of
	// these bands.
	//
	// Compute the duration of each band.
<原文结束>

# <翻译开始>
	// Divide the utilization curve up into equal size
	// non-overlapping "bands" and compute a summary for each of
	// these bands.
	//
	// Compute the duration of each band.
# <翻译结束>


<原文开始>
		// There's no point in having lots of bands if there
		// aren't many events.
<原文结束>

# <翻译开始>
		// There's no point in having lots of bands if there
		// aren't many events.
# <翻译结束>


<原文开始>
	// Compute the bands. There are numBands+1 bands in order to
	// record the final cumulative sum.
<原文结束>

# <翻译开始>
	// Compute the bands. There are numBands+1 bands in order to
	// record the final cumulative sum.
# <翻译结束>


<原文开始>
// Utilization series index
<原文结束>

# <翻译开始>
// Utilization series index
# <翻译结束>


<原文开始>
	// Lower bound of mutator utilization for all windows
	// with a left edge in this band.
<原文结束>

# <翻译开始>
	// Lower bound of mutator utilization for all windows
	// with a left edge in this band.
# <翻译结束>


<原文开始>
// UtilWindow is a specific window at Time.
<原文结束>

# <翻译开始>
// UtilWindow is a specific window at Time.
# <翻译结束>


<原文开始>
// MutatorUtil is the mean mutator utilization in this window.
<原文结束>

# <翻译开始>
// MutatorUtil is the mean mutator utilization in this window.
# <翻译结束>


<原文开始>
// An accumulator takes a windowed mutator utilization function and
// tracks various statistics for that function.
<原文结束>

# <翻译开始>
// An accumulator takes a windowed mutator utilization function and
// tracks various statistics for that function.
# <翻译结束>


<原文开始>
	// bound is the mutator utilization bound where adding any
	// mutator utilization above this bound cannot affect the
	// accumulated statistics.
<原文结束>

# <翻译开始>
	// bound is the mutator utilization bound where adding any
	// mutator utilization above this bound cannot affect the
	// accumulated statistics.
# <翻译结束>


<原文开始>
// Worst N window tracking
<原文结束>

# <翻译开始>
// Worst N window tracking
# <翻译结束>


<原文开始>
// Mutator utilization distribution tracking
<原文结束>

# <翻译开始>
// Mutator utilization distribution tracking
# <翻译结束>


<原文开始>
	// preciseMass is the distribution mass that must be precise
	// before accumulation is stopped.
<原文结束>

# <翻译开始>
	// preciseMass is the distribution mass that must be precise
	// before accumulation is stopped.
# <翻译结束>


<原文开始>
	// lastTime and lastMU are the previous point added to the
	// windowed mutator utilization function.
<原文结束>

# <翻译开始>
	// lastTime and lastMU are the previous point added to the
	// windowed mutator utilization function.
# <翻译结束>


<原文开始>
// resetTime declares a discontinuity in the windowed mutator
// utilization function by resetting the current time.
<原文结束>

# <翻译开始>
// resetTime declares a discontinuity in the windowed mutator
// utilization function by resetting the current time.
# <翻译结束>


<原文开始>
	// This only matters for distribution collection, since that's
	// the only thing that depends on the progression of the
	// windowed mutator utilization function.
<原文结束>

# <翻译开始>
	// This only matters for distribution collection, since that's
	// the only thing that depends on the progression of the
	// windowed mutator utilization function.
# <翻译结束>


<原文开始>
// addMU adds a point to the windowed mutator utilization function at
// (time, mu). This must be called for monotonically increasing values
// of time.
//
// It returns true if further calls to addMU would be pointless.
<原文结束>

# <翻译开始>
// addMU adds a point to the windowed mutator utilization function at
// (time, mu). This must be called for monotonically increasing values
// of time.
//
// It returns true if further calls to addMU would be pointless.
# <翻译结束>


<原文开始>
		// If the minimum has reached zero, it can't go any
		// lower, so we can stop early.
<原文结束>

# <翻译开始>
		// If the minimum has reached zero, it can't go any
		// lower, so we can stop early.
# <翻译结束>


<原文开始>
// Consider adding this window to the n worst.
<原文结束>

# <翻译开始>
// Consider adding this window to the n worst.
# <翻译结束>


<原文开始>
		// This window is lower than the K'th worst window.
		//
		// Check if there's any overlapping window
		// already in the heap and keep whichever is
		// worse.
<原文结束>

# <翻译开始>
		// This window is lower than the K'th worst window.
		//
		// Check if there's any overlapping window
		// already in the heap and keep whichever is
		// worse.
# <翻译结束>


<原文开始>
// Replace it with this window.
<原文结束>

# <翻译开始>
// Replace it with this window.
# <翻译结束>


<原文开始>
// We don't have N windows yet, so keep accumulating.
<原文结束>

# <翻译开始>
// We don't have N windows yet, so keep accumulating.
# <翻译结束>


<原文开始>
// Anything above the least worst window has no effect.
<原文结束>

# <翻译开始>
// Anything above the least worst window has no effect.
# <翻译结束>


<原文开始>
			// We haven't accumulated enough total precise
			// mass yet to even reach our goal, so keep
			// accumulating.
<原文结束>

# <翻译开始>
			// We haven't accumulated enough total precise
			// mass yet to even reach our goal, so keep
			// accumulating.
# <翻译结束>


<原文开始>
		// It's not worth checking percentiles every time, so
		// just keep accumulating this band.
<原文结束>

# <翻译开始>
		// It's not worth checking percentiles every time, so
		// just keep accumulating this band.
# <翻译结束>


<原文开始>
// If we've found enough 0 utilizations, we can stop immediately.
<原文结束>

# <翻译开始>
// If we've found enough 0 utilizations, we can stop immediately.
# <翻译结束>


<原文开始>
// MMU returns the minimum mutator utilization for the given time
// window. This is the minimum utilization for all windows of this
// duration across the execution. The returned value is in the range
// [0, 1].
<原文结束>

# <翻译开始>
// MMU returns the minimum mutator utilization for the given time
// window. This is the minimum utilization for all windows of this
// duration across the execution. The returned value is in the range
// [0, 1].
# <翻译结束>


<原文开始>
// Examples returns n specific examples of the lowest mutator
// utilization for the given window size. The returned windows will be
// disjoint (otherwise there would be a huge number of
// mostly-overlapping windows at the single lowest point). There are
// no guarantees on which set of disjoint windows this returns.
<原文结束>

# <翻译开始>
// Examples returns n specific examples of the lowest mutator
// utilization for the given window size. The returned windows will be
// disjoint (otherwise there would be a huge number of
// mostly-overlapping windows at the single lowest point). There are
// no guarantees on which set of disjoint windows this returns.
# <翻译结束>


<原文开始>
// MUD returns mutator utilization distribution quantiles for the
// given window size.
//
// The mutator utilization distribution is the distribution of mean
// mutator utilization across all windows of the given window size in
// the trace.
//
// The minimum mutator utilization is the minimum (0th percentile) of
// this distribution. (However, if only the minimum is desired, it's
// more efficient to use the MMU method.)
<原文结束>

# <翻译开始>
// MUD returns mutator utilization distribution quantiles for the
// given window size.
//
// The mutator utilization distribution is the distribution of mean
// mutator utilization across all windows of the given window size in
// the trace.
//
// The minimum mutator utilization is the minimum (0th percentile) of
// this distribution. (However, if only the minimum is desired, it's
// more efficient to use the MMU method.)
# <翻译结束>


<原文开始>
	// Each unrefined band contributes a known total mass to the
	// distribution (bandDur except at the end), but in an unknown
	// way. However, we know that all the mass it contributes must
	// be at or above its worst-case mean mutator utilization.
	//
	// Hence, we refine bands until the highest desired
	// distribution quantile is less than the next worst-case mean
	// mutator utilization. At this point, all further
	// contributions to the distribution must be beyond the
	// desired quantile and hence cannot affect it.
	//
	// First, find the highest desired distribution quantile.
<原文结束>

# <翻译开始>
	// Each unrefined band contributes a known total mass to the
	// distribution (bandDur except at the end), but in an unknown
	// way. However, we know that all the mass it contributes must
	// be at or above its worst-case mean mutator utilization.
	//
	// Hence, we refine bands until the highest desired
	// distribution quantile is less than the next worst-case mean
	// mutator utilization. At this point, all further
	// contributions to the distribution must be beyond the
	// desired quantile and hence cannot affect it.
	//
	// First, find the highest desired distribution quantile.
# <翻译结束>


<原文开始>
	// The distribution's mass is in units of time (it's not
	// normalized because this would make it more annoying to
	// account for future contributions of unrefined bands). The
	// total final mass will be the duration of the trace itself
	// minus the window size. Using this, we can compute the mass
	// corresponding to quantile maxQ.
<原文结束>

# <翻译开始>
	// The distribution's mass is in units of time (it's not
	// normalized because this would make it more annoying to
	// account for future contributions of unrefined bands). The
	// total final mass will be the duration of the trace itself
	// minus the window size. Using this, we can compute the mass
	// corresponding to quantile maxQ.
# <翻译结束>


<原文开始>
	// Accumulate the MUD until we have precise information for
	// everything to the left of qMass.
<原文结束>

# <翻译开始>
	// Accumulate the MUD until we have precise information for
	// everything to the left of qMass.
# <翻译结束>


<原文开始>
// Evaluate the quantiles on the accumulated MUD.
<原文结束>

# <翻译开始>
// Evaluate the quantiles on the accumulated MUD.
# <翻译结束>


<原文开始>
			// There are a few legitimate ways this can
			// happen:
			//
			// 1. If the window is the full trace
			// duration, then the windowed MU function is
			// only defined at a single point, so the MU
			// distribution is not well-defined.
			//
			// 2. If there are no events, then the MU
			// distribution has no mass.
			//
			// Either way, all of the quantiles will have
			// converged toward the MMU at this point.
<原文结束>

# <翻译开始>
			// There are a few legitimate ways this can
			// happen:
			//
			// 1. If the window is the full trace
			// duration, then the windowed MU function is
			// only defined at a single point, so the MU
			// distribution is not well-defined.
			//
			// 2. If there are no events, then the MU
			// distribution has no mass.
			//
			// Either way, all of the quantiles will have
			// converged toward the MMU at this point.
# <翻译结束>


<原文开始>
// Process bands from lowest utilization bound to highest.
<原文结束>

# <翻译开始>
// Process bands from lowest utilization bound to highest.
# <翻译结束>


<原文开始>
	// Refine each band into a precise window and MMU until
	// refining the next lowest band can no longer affect the MMU
	// or windows.
<原文结束>

# <翻译开始>
	// Refine each band into a precise window and MMU until
	// refining the next lowest band can no longer affect the MMU
	// or windows.
# <翻译结束>


<原文开始>
	// For each band, compute the worst-possible total mutator
	// utilization for all windows that start in that band.
<原文结束>

# <翻译开始>
	// For each band, compute the worst-possible total mutator
	// utilization for all windows that start in that band.
# <翻译结束>


<原文开始>
	// minBands is the minimum number of bands a window can span
	// and maxBands is the maximum number of bands a window can
	// span in any alignment.
<原文结束>

# <翻译开始>
	// minBands is the minimum number of bands a window can span
	// and maxBands is the maximum number of bands a window can
	// span in any alignment.
# <翻译结束>


<原文开始>
		// To compute the worst-case MU, we assume the minimum
		// for any bands that are only partially overlapped by
		// some window and the mean for any bands that are
		// completely covered by all windows.
<原文结束>

# <翻译开始>
		// To compute the worst-case MU, we assume the minimum
		// for any bands that are only partially overlapped by
		// some window and the mean for any bands that are
		// completely covered by all windows.
# <翻译结束>


<原文开始>
		// Find the lowest and second lowest of the partial
		// bands.
<原文结束>

# <翻译开始>
		// Find the lowest and second lowest of the partial
		// bands.
# <翻译结束>


<原文开始>
		// Assume the worst window maximally overlaps the
		// worst minimum and then the rest overlaps the second
		// worst minimum.
<原文结束>

# <翻译开始>
		// Assume the worst window maximally overlaps the
		// worst minimum and then the rest overlaps the second
		// worst minimum.
# <翻译结束>


<原文开始>
		// Add the total mean MU of bands that are completely
		// overlapped by all windows.
<原文结束>

# <翻译开始>
		// Add the total mean MU of bands that are completely
		// overlapped by all windows.
# <翻译结束>


<原文开始>
// bandMMU computes the precise minimum mutator utilization for
// windows with a left edge in band bandIdx.
<原文结束>

# <翻译开始>
// bandMMU computes the precise minimum mutator utilization for
// windows with a left edge in band bandIdx.
# <翻译结束>


<原文开始>
	// We think of the mutator utilization over time as the
	// box-filtered utilization function, which we call the
	// "windowed mutator utilization function". The resulting
	// function is continuous and piecewise linear (unless
	// window==0, which we handle elsewhere), where the boundaries
	// between segments occur when either edge of the window
	// encounters a change in the instantaneous mutator
	// utilization function. Hence, the minimum of this function
	// will always occur when one of the edges of the window
	// aligns with a utilization change, so these are the only
	// points we need to consider.
	//
	// We compute the mutator utilization function incrementally
	// by tracking the integral from t=0 to the left edge of the
	// window and to the right edge of the window.
<原文结束>

# <翻译开始>
	// We think of the mutator utilization over time as the
	// box-filtered utilization function, which we call the
	// "windowed mutator utilization function". The resulting
	// function is continuous and piecewise linear (unless
	// window==0, which we handle elsewhere), where the boundaries
	// between segments occur when either edge of the window
	// encounters a change in the instantaneous mutator
	// utilization function. Hence, the minimum of this function
	// will always occur when one of the edges of the window
	// aligns with a utilization change, so these are the only
	// points we need to consider.
	//
	// We compute the mutator utilization function incrementally
	// by tracking the integral from t=0 to the left edge of the
	// window and to the right edge of the window.
# <翻译结束>


<原文开始>
// Advance edges to time and time+window.
<原文结束>

# <翻译开始>
// Advance edges to time and time+window.
# <翻译结束>


<原文开始>
		// The maximum slope of the windowed mutator
		// utilization function is 1/window, so we can always
		// advance the time by at least (mu - mmu) * window
		// without dropping below mmu.
<原文结束>

# <翻译开始>
		// The maximum slope of the windowed mutator
		// utilization function is 1/window, so we can always
		// advance the time by at least (mu - mmu) * window
		// without dropping below mmu.
# <翻译结束>


<原文开始>
		// Advance the window to the next time where either
		// the left or right edge of the window encounters a
		// change in the utilization curve.
<原文结束>

# <翻译开始>
		// Advance the window to the next time where either
		// the left or right edge of the window encounters a
		// change in the utilization curve.
# <翻译结束>


<原文开始>
			// For MMUs we could stop here, but for MUDs
			// it's important that we span the entire
			// band.
<原文结束>

# <翻译开始>
			// For MMUs we could stop here, but for MUDs
			// it's important that we span the entire
			// band.
# <翻译结束>


<原文开始>
// An integrator tracks a position in a utilization function and
// integrates it.
<原文结束>

# <翻译开始>
// An integrator tracks a position in a utilization function and
// integrates it.
# <翻译结束>


<原文开始>
	// pos is the index in u.util of the current time's non-strict
	// predecessor.
<原文结束>

# <翻译开始>
	// pos is the index in u.util of the current time's non-strict
	// predecessor.
# <翻译结束>


<原文开始>
// advance returns the integral of the utilization function from 0 to
// time. advance must be called on monotonically increasing values of
// times.
<原文结束>

# <翻译开始>
// advance returns the integral of the utilization function from 0 to
// time. advance must be called on monotonically increasing values of
// times.
# <翻译结束>


<原文开始>
	// Advance pos until pos+1 is time's strict successor (making
	// pos time's non-strict predecessor).
	//
	// Very often, this will be nearby, so we optimize that case,
	// but it may be arbitrarily far away, so we handled that
	// efficiently, too.
<原文结束>

# <翻译开始>
	// Advance pos until pos+1 is time's strict successor (making
	// pos time's non-strict predecessor).
	//
	// Very often, this will be nearby, so we optimize that case,
	// but it may be arbitrarily far away, so we handled that
	// efficiently, too.
# <翻译结束>


<原文开始>
// Nearby. Use a linear scan.
<原文结束>

# <翻译开始>
// Nearby. Use a linear scan.
# <翻译结束>


<原文开始>
// Far. Binary search for time's strict successor.
<原文结束>

# <翻译开始>
// Far. Binary search for time's strict successor.
# <翻译结束>


<原文开始>
// Non-strict predecessor.
<原文结束>

# <翻译开始>
// Non-strict predecessor.
# <翻译结束>


<原文开始>
// next returns the smallest time t' > time of a change in the
// utilization function.
<原文结束>

# <翻译开始>
// next returns the smallest time t' > time of a change in the
// utilization function.
# <翻译结束>

