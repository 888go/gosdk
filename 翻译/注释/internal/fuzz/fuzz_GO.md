
<原文开始>
// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
# <翻译结束>


<原文开始>
// Package fuzz provides common fuzzing functionality for tests built with
// "go test" and for programs that use fuzzing functionality in the testing
// package.
<原文结束>

# <翻译开始>
// Package fuzz provides common fuzzing functionality for tests built with
// "go test" and for programs that use fuzzing functionality in the testing
// package.
# <翻译结束>


<原文开始>
// CoordinateFuzzingOpts is a set of arguments for CoordinateFuzzing.
// The zero value is valid for each field unless specified otherwise.
<原文结束>

# <翻译开始>
// CoordinateFuzzingOpts is a set of arguments for CoordinateFuzzing.
// The zero value is valid for each field unless specified otherwise.
# <翻译结束>


<原文开始>
	// Log is a writer for logging progress messages and warnings.
	// If nil, io.Discard will be used instead.
<原文结束>

# <翻译开始>
	// Log is a writer for logging progress messages and warnings.
	// If nil, io.Discard will be used instead.
# <翻译结束>


<原文开始>
	// Timeout is the amount of wall clock time to spend fuzzing after the corpus
	// has loaded. If zero, there will be no time limit.
<原文结束>

# <翻译开始>
	// Timeout is the amount of wall clock time to spend fuzzing after the corpus
	// has loaded. If zero, there will be no time limit.
# <翻译结束>


<原文开始>
	// Limit is the number of random values to generate and test. If zero,
	// there will be no limit on the number of generated values.
<原文结束>

# <翻译开始>
	// Limit is the number of random values to generate and test. If zero,
	// there will be no limit on the number of generated values.
# <翻译结束>


<原文开始>
	// MinimizeTimeout is the amount of wall clock time to spend minimizing
	// after discovering a crasher. If zero, there will be no time limit. If
	// MinimizeTimeout and MinimizeLimit are both zero, then minimization will
	// be disabled.
<原文结束>

# <翻译开始>
	// MinimizeTimeout is the amount of wall clock time to spend minimizing
	// after discovering a crasher. If zero, there will be no time limit. If
	// MinimizeTimeout and MinimizeLimit are both zero, then minimization will
	// be disabled.
# <翻译结束>


<原文开始>
	// MinimizeLimit is the maximum number of calls to the fuzz function to be
	// made while minimizing after finding a crash. If zero, there will be no
	// limit. Calls to the fuzz function made when minimizing also count toward
	// Limit. If MinimizeTimeout and MinimizeLimit are both zero, then
	// minimization will be disabled.
<原文结束>

# <翻译开始>
	// MinimizeLimit is the maximum number of calls to the fuzz function to be
	// made while minimizing after finding a crash. If zero, there will be no
	// limit. Calls to the fuzz function made when minimizing also count toward
	// Limit. If MinimizeTimeout and MinimizeLimit are both zero, then
	// minimization will be disabled.
# <翻译结束>


<原文开始>
	// parallel is the number of worker processes to run in parallel. If zero,
	// CoordinateFuzzing will run GOMAXPROCS workers.
<原文结束>

# <翻译开始>
	// parallel is the number of worker processes to run in parallel. If zero,
	// CoordinateFuzzing will run GOMAXPROCS workers.
# <翻译结束>


<原文开始>
	// Seed is a list of seed values added by the fuzz target with testing.F.Add
	// and in testdata.
<原文结束>

# <翻译开始>
	// Seed is a list of seed values added by the fuzz target with testing.F.Add
	// and in testdata.
# <翻译结束>


<原文开始>
	// Types is the list of types which make up a corpus entry.
	// Types must be set and must match values in Seed.
<原文结束>

# <翻译开始>
	// Types is the list of types which make up a corpus entry.
	// Types must be set and must match values in Seed.
# <翻译结束>


<原文开始>
	// CorpusDir is a directory where files containing values that crash the
	// code being tested may be written. CorpusDir must be set.
<原文结束>

# <翻译开始>
	// CorpusDir is a directory where files containing values that crash the
	// code being tested may be written. CorpusDir must be set.
# <翻译结束>


<原文开始>
	// CacheDir is a directory containing additional "interesting" values.
	// The fuzzer may derive new values from these, and may write new values here.
<原文结束>

# <翻译开始>
	// CacheDir is a directory containing additional "interesting" values.
	// The fuzzer may derive new values from these, and may write new values here.
# <翻译结束>


<原文开始>
// CoordinateFuzzing creates several worker processes and communicates with
// them to test random inputs that could trigger crashes and expose bugs.
// The worker processes run the same binary in the same directory with the
// same environment variables as the coordinator process. Workers also run
// with the same arguments as the coordinator, except with the -test.fuzzworker
// flag prepended to the argument list.
//
// If a crash occurs, the function will return an error containing information
// about the crash, which can be reported to the user.
<原文结束>

# <翻译开始>
// CoordinateFuzzing creates several worker processes and communicates with
// them to test random inputs that could trigger crashes and expose bugs.
// The worker processes run the same binary in the same directory with the
// same environment variables as the coordinator process. Workers also run
// with the same arguments as the coordinator, except with the -test.fuzzworker
// flag prepended to the argument list.
//
// If a crash occurs, the function will return an error containing information
// about the crash, which can be reported to the user.
# <翻译结束>


<原文开始>
// Don't start more workers than we need.
<原文结束>

# <翻译开始>
// Don't start more workers than we need.
# <翻译结束>


<原文开始>
// fuzzCtx is used to stop workers, for example, after finding a crasher.
<原文结束>

# <翻译开始>
// fuzzCtx is used to stop workers, for example, after finding a crasher.
# <翻译结束>


<原文开始>
// stop is called when a worker encounters a fatal error.
<原文结束>

# <翻译开始>
// stop is called when a worker encounters a fatal error.
# <翻译结束>


<原文开始>
			// Suppress cancellation errors and terminations due to SIGINT.
			// The messages are not helpful since either the user triggered the error
			// (with ^C) or another more helpful message will be printed (a crasher).
<原文结束>

# <翻译开始>
			// Suppress cancellation errors and terminations due to SIGINT.
			// The messages are not helpful since either the user triggered the error
			// (with ^C) or another more helpful message will be printed (a crasher).
# <翻译结束>


<原文开始>
	// Ensure that any crash we find is written to the corpus, even if an error
	// or interruption occurs while minimizing it.
<原文结束>

# <翻译开始>
	// Ensure that any crash we find is written to the corpus, even if an error
	// or interruption occurs while minimizing it.
# <翻译结束>


<原文开始>
	// Start workers.
	// TODO(jayconrod): do we want to support fuzzing different binaries?
<原文结束>

# <翻译开始>
	// Start workers.
	// TODO(jayconrod): do we want to support fuzzing different binaries?
# <翻译结束>


<原文开始>
	// Main event loop.
	// Do not return until all workers have terminated. We avoid a deadlock by
	// receiving messages from workers even after ctx is cancelled.
<原文结束>

# <翻译开始>
	// Main event loop.
	// Do not return until all workers have terminated. We avoid a deadlock by
	// receiving messages from workers even after ctx is cancelled.
# <翻译结束>


<原文开始>
			// Interrupted, cancelled, or timed out.
			// stop sets doneC to nil so we don't busy wait here.
<原文结束>

# <翻译开始>
			// Interrupted, cancelled, or timed out.
			// stop sets doneC to nil so we don't busy wait here.
# <翻译结束>


<原文开始>
// A worker terminated, possibly after encountering a fatal error.
<原文结束>

# <翻译开始>
// A worker terminated, possibly after encountering a fatal error.
# <翻译结束>


<原文开始>
// Received response from worker.
<原文结束>

# <翻译开始>
// Received response from worker.
# <翻译结束>


<原文开始>
						// This crash is not minimized, and another crash is being minimized.
						// Ignore this one and wait for the other one to finish.
<原文结束>

# <翻译开始>
						// This crash is not minimized, and another crash is being minimized.
						// Ignore this one and wait for the other one to finish.
# <翻译结束>


<原文开始>
					// Found a crasher but haven't yet attempted to minimize it.
					// Send it back to a worker for minimization. Disable inputC so
					// other workers don't continue fuzzing.
<原文结束>

# <翻译开始>
					// Found a crasher but haven't yet attempted to minimize it.
					// Send it back to a worker for minimization. Disable inputC so
					// other workers don't continue fuzzing.
# <翻译结束>


<原文开始>
					// Found a crasher that's either minimized or not minimizable.
					// Write to corpus and stop.
<原文结束>

# <翻译开始>
					// Found a crasher that's either minimized or not minimizable.
					// Write to corpus and stop.
# <翻译结束>


<原文开始>
					// Found a value that expanded coverage.
					// It's not a crasher, but we may want to add it to the on-disk
					// corpus and prioritize it for future fuzzing.
					// TODO(jayconrod, katiehockman): Prioritize fuzzing these
					// values which expanded coverage, perhaps based on the
					// number of new edges that this result expanded.
					// TODO(jayconrod, katiehockman): Don't write a value that's already
					// in the corpus.
<原文结束>

# <翻译开始>
					// Found a value that expanded coverage.
					// It's not a crasher, but we may want to add it to the on-disk
					// corpus and prioritize it for future fuzzing.
					// TODO(jayconrod, katiehockman): Prioritize fuzzing these
					// values which expanded coverage, perhaps based on the
					// number of new edges that this result expanded.
					// TODO(jayconrod, katiehockman): Don't write a value that's already
					// in the corpus.
# <翻译结束>


<原文开始>
						// Send back to workers to find a smaller value that preserves
						// at least one new coverage bit.
<原文结束>

# <翻译开始>
						// Send back to workers to find a smaller value that preserves
						// at least one new coverage bit.
# <翻译结束>


<原文开始>
// Update the coordinator's coverage mask and save the value.
<原文结束>

# <翻译开始>
// Update the coordinator's coverage mask and save the value.
# <翻译结束>


<原文开始>
				// No error or coverage data was reported for this input during
				// warmup, so continue processing results.
<原文结束>

# <翻译开始>
				// No error or coverage data was reported for this input during
				// warmup, so continue processing results.
# <翻译结束>


<原文开始>
			// Once the result has been processed, stop the worker if we
			// have reached the fuzzing limit.
<原文结束>

# <翻译开始>
			// Once the result has been processed, stop the worker if we
			// have reached the fuzzing limit.
# <翻译结束>


<原文开始>
// Sent the next input to a worker.
<原文结束>

# <翻译开始>
// Sent the next input to a worker.
# <翻译结束>


<原文开始>
// Sent the next input for minimization to a worker.
<原文结束>

# <翻译开始>
// Sent the next input for minimization to a worker.
# <翻译结束>


<原文开始>
	// TODO(jayconrod,katiehockman): if a crasher can't be written to the corpus,
	// write to the cache instead.
<原文结束>

# <翻译开始>
	// TODO(jayconrod,katiehockman): if a crasher can't be written to the corpus,
	// write to the cache instead.
# <翻译结束>


<原文开始>
// crashError wraps a crasher written to the seed corpus. It saves the name
// of the file where the input causing the crasher was saved. The testing
// framework uses this to report a command to re-run that specific input.
<原文结束>

# <翻译开始>
// crashError wraps a crasher written to the seed corpus. It saves the name
// of the file where the input causing the crasher was saved. The testing
// framework uses this to report a command to re-run that specific input.
# <翻译结束>


<原文开始>
// addCorpusEntries adds entries to the corpus, and optionally writes the entries
// to the cache directory. If an entry is already in the corpus it is skipped. If
// all of the entries are unique, addCorpusEntries returns true and a nil error,
// if at least one of the entries was a duplicate, it returns false and a nil error.
<原文结束>

# <翻译开始>
// addCorpusEntries adds entries to the corpus, and optionally writes the entries
// to the cache directory. If an entry is already in the corpus it is skipped. If
// all of the entries are unique, addCorpusEntries returns true and a nil error,
// if at least one of the entries was a duplicate, it returns false and a nil error.
# <翻译结束>


<原文开始>
			// For entries written to disk, we don't hold onto the bytes,
			// since the corpus would consume a significant amount of
			// memory.
<原文结束>

# <翻译开始>
			// For entries written to disk, we don't hold onto the bytes,
			// since the corpus would consume a significant amount of
			// memory.
# <翻译结束>


<原文开始>
// CorpusEntry represents an individual input for fuzzing.
//
// We must use an equivalent type in the testing and testing/internal/testdeps
// packages, but testing can't import this package directly, and we don't want
// to export this type from testing. Instead, we use the same struct type and
// use a type alias (not a defined type) for convenience.
<原文结束>

# <翻译开始>
// CorpusEntry represents an individual input for fuzzing.
//
// We must use an equivalent type in the testing and testing/internal/testdeps
// packages, but testing can't import this package directly, and we don't want
// to export this type from testing. Instead, we use the same struct type and
// use a type alias (not a defined type) for convenience.
# <翻译结束>


<原文开始>
	// Path is the path of the corpus file, if the entry was loaded from disk.
	// For other entries, including seed values provided by f.Add, Path is the
	// name of the test, e.g. seed#0 or its hash.
<原文结束>

# <翻译开始>
	// Path is the path of the corpus file, if the entry was loaded from disk.
	// For other entries, including seed values provided by f.Add, Path is the
	// name of the test, e.g. seed#0 or its hash.
# <翻译结束>


<原文开始>
	// Data is the raw input data. Data should only be populated for seed
	// values. For on-disk corpus files, Data will be nil, as it will be loaded
	// from disk using Path.
<原文结束>

# <翻译开始>
	// Data is the raw input data. Data should only be populated for seed
	// values. For on-disk corpus files, Data will be nil, as it will be loaded
	// from disk using Path.
# <翻译结束>


<原文开始>
// Values is the unmarshaled values from a corpus file.
<原文结束>

# <翻译开始>
// Values is the unmarshaled values from a corpus file.
# <翻译结束>


<原文开始>
// IsSeed indicates whether this entry is part of the seed corpus.
<原文结束>

# <翻译开始>
// IsSeed indicates whether this entry is part of the seed corpus.
# <翻译结束>


<原文开始>
// corpusEntryData returns the raw input bytes, either from the data struct
// field, or from disk.
<原文结束>

# <翻译开始>
// corpusEntryData returns the raw input bytes, either from the data struct
// field, or from disk.
# <翻译结束>


<原文开始>
	// entry is the value to test initially. The worker will randomly mutate
	// values from this starting point.
<原文结束>

# <翻译开始>
	// entry is the value to test initially. The worker will randomly mutate
	// values from this starting point.
# <翻译结束>


<原文开始>
	// timeout is the time to spend fuzzing variations of this input,
	// not including starting or cleaning up.
<原文结束>

# <翻译开始>
	// timeout is the time to spend fuzzing variations of this input,
	// not including starting or cleaning up.
# <翻译结束>


<原文开始>
	// limit is the maximum number of calls to the fuzz function the worker may
	// make. The worker may make fewer calls, for example, if it finds an
	// error early. If limit is zero, there is no limit on calls to the
	// fuzz function.
<原文结束>

# <翻译开始>
	// limit is the maximum number of calls to the fuzz function the worker may
	// make. The worker may make fewer calls, for example, if it finds an
	// error early. If limit is zero, there is no limit on calls to the
	// fuzz function.
# <翻译结束>


<原文开始>
	// warmup indicates whether this is a warmup input before fuzzing begins. If
	// true, the input should not be fuzzed.
<原文结束>

# <翻译开始>
	// warmup indicates whether this is a warmup input before fuzzing begins. If
	// true, the input should not be fuzzed.
# <翻译结束>


<原文开始>
// coverageData reflects the coordinator's current coverageMask.
<原文结束>

# <翻译开始>
// coverageData reflects the coordinator's current coverageMask.
# <翻译结束>


<原文开始>
// entry is an interesting value or a crasher.
<原文结束>

# <翻译开始>
// entry is an interesting value or a crasher.
# <翻译结束>


<原文开始>
// crasherMsg is an error message from a crash. It's "" if no crash was found.
<原文结束>

# <翻译开始>
// crasherMsg is an error message from a crash. It's "" if no crash was found.
# <翻译结束>


<原文开始>
	// canMinimize is true if the worker should attempt to minimize this result.
	// It may be false because an attempt has already been made.
<原文结束>

# <翻译开始>
	// canMinimize is true if the worker should attempt to minimize this result.
	// It may be false because an attempt has already been made.
# <翻译结束>


<原文开始>
// coverageData is set if the worker found new coverage.
<原文结束>

# <翻译开始>
// coverageData is set if the worker found new coverage.
# <翻译结束>


<原文开始>
	// limit is the number of values the coordinator asked the worker
	// to test. 0 if there was no limit.
<原文结束>

# <翻译开始>
	// limit is the number of values the coordinator asked the worker
	// to test. 0 if there was no limit.
# <翻译结束>


<原文开始>
// count is the number of values the worker actually tested.
<原文结束>

# <翻译开始>
// count is the number of values the worker actually tested.
# <翻译结束>


<原文开始>
// totalDuration is the time the worker spent testing inputs.
<原文结束>

# <翻译开始>
// totalDuration is the time the worker spent testing inputs.
# <翻译结束>


<原文开始>
// entryDuration is the time the worker spent execution an interesting result
<原文结束>

# <翻译开始>
// entryDuration is the time the worker spent execution an interesting result
# <翻译结束>


<原文开始>
// entry is an interesting value or crasher to minimize.
<原文结束>

# <翻译开始>
// entry is an interesting value or crasher to minimize.
# <翻译结束>


<原文开始>
	// crasherMsg is an error message from a crash. It's "" if no crash was found.
	// If set, the worker will attempt to find a smaller input that also produces
	// an error, though not necessarily the same error.
<原文结束>

# <翻译开始>
	// crasherMsg is an error message from a crash. It's "" if no crash was found.
	// If set, the worker will attempt to find a smaller input that also produces
	// an error, though not necessarily the same error.
# <翻译结束>


<原文开始>
	// limit is the maximum number of calls to the fuzz function the worker may
	// make. The worker may make fewer calls, for example, if it can't reproduce
	// an error. If limit is zero, there is no limit on calls to the fuzz function.
<原文结束>

# <翻译开始>
	// limit is the maximum number of calls to the fuzz function the worker may
	// make. The worker may make fewer calls, for example, if it can't reproduce
	// an error. If limit is zero, there is no limit on calls to the fuzz function.
# <翻译结束>


<原文开始>
	// timeout is the time to spend minimizing this input.
	// A zero timeout means no limit.
<原文结束>

# <翻译开始>
	// timeout is the time to spend minimizing this input.
	// A zero timeout means no limit.
# <翻译结束>


<原文开始>
	// keepCoverage is a set of coverage bits that entry found that were not in
	// the coordinator's combined set. When minimizing, the worker should find an
	// input that preserves at least one of these bits. keepCoverage is nil for
	// crashing inputs.
<原文结束>

# <翻译开始>
	// keepCoverage is a set of coverage bits that entry found that were not in
	// the coordinator's combined set. When minimizing, the worker should find an
	// input that preserves at least one of these bits. keepCoverage is nil for
	// crashing inputs.
# <翻译结束>


<原文开始>
// coordinator holds channels that workers can use to communicate with
// the coordinator.
<原文结束>

# <翻译开始>
// coordinator holds channels that workers can use to communicate with
// the coordinator.
# <翻译结束>


<原文开始>
	// startTime is the time we started the workers after loading the corpus.
	// Used for logging.
<原文结束>

# <翻译开始>
	// startTime is the time we started the workers after loading the corpus.
	// Used for logging.
# <翻译结束>


<原文开始>
	// inputC is sent values to fuzz by the coordinator. Any worker may receive
	// values from this channel. Workers send results to resultC.
<原文结束>

# <翻译开始>
	// inputC is sent values to fuzz by the coordinator. Any worker may receive
	// values from this channel. Workers send results to resultC.
# <翻译结束>


<原文开始>
	// minimizeC is sent values to minimize by the coordinator. Any worker may
	// receive values from this channel. Workers send results to resultC.
<原文结束>

# <翻译开始>
	// minimizeC is sent values to minimize by the coordinator. Any worker may
	// receive values from this channel. Workers send results to resultC.
# <翻译结束>


<原文开始>
	// resultC is sent results of fuzzing by workers. The coordinator
	// receives these. Multiple types of messages are allowed.
<原文结束>

# <翻译开始>
	// resultC is sent results of fuzzing by workers. The coordinator
	// receives these. Multiple types of messages are allowed.
# <翻译结束>


<原文开始>
// count is the number of values fuzzed so far.
<原文结束>

# <翻译开始>
// count is the number of values fuzzed so far.
# <翻译结束>


<原文开始>
	// countLastLog is the number of values fuzzed when the output was last
	// logged.
<原文结束>

# <翻译开始>
	// countLastLog is the number of values fuzzed when the output was last
	// logged.
# <翻译结束>


<原文开始>
// timeLastLog is the time at which the output was last logged.
<原文结束>

# <翻译开始>
// timeLastLog is the time at which the output was last logged.
# <翻译结束>


<原文开始>
	// interestingCount is the number of unique interesting values which have
	// been found this execution.
<原文结束>

# <翻译开始>
	// interestingCount is the number of unique interesting values which have
	// been found this execution.
# <翻译结束>


<原文开始>
	// warmupInputCount is the count of all entries in the corpus which will
	// need to be received from workers to run once during warmup, but not fuzz.
	// This could be for coverage data, or only for the purposes of verifying
	// that the seed corpus doesn't have any crashers. See warmupRun.
<原文结束>

# <翻译开始>
	// warmupInputCount is the count of all entries in the corpus which will
	// need to be received from workers to run once during warmup, but not fuzz.
	// This could be for coverage data, or only for the purposes of verifying
	// that the seed corpus doesn't have any crashers. See warmupRun.
# <翻译结束>


<原文开始>
	// warmupInputLeft is the number of entries in the corpus which still need
	// to be received from workers to run once during warmup, but not fuzz.
	// See warmupInputLeft.
<原文结束>

# <翻译开始>
	// warmupInputLeft is the number of entries in the corpus which still need
	// to be received from workers to run once during warmup, but not fuzz.
	// See warmupInputLeft.
# <翻译结束>


<原文开始>
	// duration is the time spent fuzzing inside workers, not counting time
	// starting up or tearing down.
<原文结束>

# <翻译开始>
	// duration is the time spent fuzzing inside workers, not counting time
	// starting up or tearing down.
# <翻译结束>


<原文开始>
	// countWaiting is the number of fuzzing executions the coordinator is
	// waiting on workers to complete.
<原文结束>

# <翻译开始>
	// countWaiting is the number of fuzzing executions the coordinator is
	// waiting on workers to complete.
# <翻译结束>


<原文开始>
	// corpus is a set of interesting values, including the seed corpus and
	// generated values that workers reported as interesting.
<原文结束>

# <翻译开始>
	// corpus is a set of interesting values, including the seed corpus and
	// generated values that workers reported as interesting.
# <翻译结束>


<原文开始>
	// minimizationAllowed is true if one or more of the types of fuzz
	// function's parameters can be minimized.
<原文结束>

# <翻译开始>
	// minimizationAllowed is true if one or more of the types of fuzz
	// function's parameters can be minimized.
# <翻译结束>


<原文开始>
	// inputQueue is a queue of inputs that workers should try fuzzing. This is
	// initially populated from the seed corpus and cached inputs. More inputs
	// may be added as new coverage is discovered.
<原文结束>

# <翻译开始>
	// inputQueue is a queue of inputs that workers should try fuzzing. This is
	// initially populated from the seed corpus and cached inputs. More inputs
	// may be added as new coverage is discovered.
# <翻译结束>


<原文开始>
	// minimizeQueue is a queue of inputs that caused errors or exposed new
	// coverage. Workers should attempt to find smaller inputs that do the
	// same thing.
<原文结束>

# <翻译开始>
	// minimizeQueue is a queue of inputs that caused errors or exposed new
	// coverage. Workers should attempt to find smaller inputs that do the
	// same thing.
# <翻译结束>


<原文开始>
// crashMinimizing is the crash that is currently being minimized.
<原文结束>

# <翻译开始>
// crashMinimizing is the crash that is currently being minimized.
# <翻译结束>


<原文开始>
	// coverageMask aggregates coverage that was found for all inputs in the
	// corpus. Each byte represents a single basic execution block. Each set bit
	// within the byte indicates that an input has triggered that block at least
	// 1 << n times, where n is the position of the bit in the byte. For example, a
	// value of 12 indicates that separate inputs have triggered this block
	// between 4-7 times and 8-15 times.
<原文结束>

# <翻译开始>
	// coverageMask aggregates coverage that was found for all inputs in the
	// corpus. Each byte represents a single basic execution block. Each set bit
	// within the byte indicates that an input has triggered that block at least
	// 1 << n times, where n is the position of the bit in the byte. For example, a
	// value of 12 indicates that separate inputs have triggered this block
	// between 4-7 times and 8-15 times.
# <翻译结束>


<原文开始>
// Make sure all of the seed corpus has marshalled data.
<原文结束>

# <翻译开始>
// Make sure all of the seed corpus has marshalled data.
# <翻译结束>


<原文开始>
		// Even though a coverage-only run won't occur, we should still run all
		// of the seed corpus to make sure there are no existing failures before
		// we start fuzzing.
<原文结束>

# <翻译开始>
		// Even though a coverage-only run won't occur, we should still run all
		// of the seed corpus to make sure there are no existing failures before
		// we start fuzzing.
# <翻译结束>


<原文开始>
// Set c.coverageMask to a clean []byte full of zeros.
<原文结束>

# <翻译开始>
// Set c.coverageMask to a clean []byte full of zeros.
# <翻译结束>


<原文开始>
// peekInput returns the next value that should be sent to workers.
// If the number of executions is limited, the returned value includes
// a limit for one worker. If there are no executions left, peekInput returns
// a zero value and false.
//
// peekInput doesn't actually remove the input from the queue. The caller
// must call sentInput after sending the input.
//
// If the input queue is empty and the coverage/testing-only run has completed,
// queue refills it from the corpus.
<原文结束>

# <翻译开始>
// peekInput returns the next value that should be sent to workers.
// If the number of executions is limited, the returned value includes
// a limit for one worker. If there are no executions left, peekInput returns
// a zero value and false.
//
// peekInput doesn't actually remove the input from the queue. The caller
// must call sentInput after sending the input.
//
// If the input queue is empty and the coverage/testing-only run has completed,
// queue refills it from the corpus.
# <翻译结束>


<原文开始>
		// Already making the maximum number of calls to the fuzz function.
		// Don't send more inputs right now.
<原文结束>

# <翻译开始>
		// Already making the maximum number of calls to the fuzz function.
		// Don't send more inputs right now.
# <翻译结束>


<原文开始>
			// Wait for coverage/testing-only run to finish before sending more
			// inputs.
<原文结束>

# <翻译开始>
			// Wait for coverage/testing-only run to finish before sending more
			// inputs.
# <翻译结束>


<原文开始>
		// No fuzzing will occur, but it should count toward the limit set by
		// -fuzztime.
<原文结束>

# <翻译开始>
		// No fuzzing will occur, but it should count toward the limit set by
		// -fuzztime.
# <翻译结束>


<原文开始>
// sentInput updates internal counters after an input is sent to c.inputC.
<原文结束>

# <翻译开始>
// sentInput updates internal counters after an input is sent to c.inputC.
# <翻译结束>


<原文开始>
// refillInputQueue refills the input queue from the corpus after it becomes
// empty.
<原文结束>

# <翻译开始>
// refillInputQueue refills the input queue from the corpus after it becomes
// empty.
# <翻译结束>


<原文开始>
// queueForMinimization creates a fuzzMinimizeInput from result and adds it
// to the minimization queue to be sent to workers.
<原文结束>

# <翻译开始>
// queueForMinimization creates a fuzzMinimizeInput from result and adds it
// to the minimization queue to be sent to workers.
# <翻译结束>


<原文开始>
// peekMinimizeInput returns the next input that should be sent to workers for
// minimization.
<原文结束>

# <翻译开始>
// peekMinimizeInput returns the next input that should be sent to workers for
// minimization.
# <翻译结束>


<原文开始>
// sentMinimizeInput removes an input from the minimization queue after it's
// sent to minimizeC.
<原文结束>

# <翻译开始>
// sentMinimizeInput removes an input from the minimization queue after it's
// sent to minimizeC.
# <翻译结束>


<原文开始>
// warmupRun returns true while the coordinator is running inputs without
// mutating them as a warmup before fuzzing. This could be to gather baseline
// coverage data for entries in the corpus, or to test all of the seed corpus
// for errors before fuzzing begins.
//
// The coordinator doesn't store coverage data in the cache with each input
// because that data would be invalid when counter offsets in the test binary
// change.
//
// When gathering coverage, the coordinator sends each entry to a worker to
// gather coverage for that entry only, without fuzzing or minimizing. This
// phase ends when all workers have finished, and the coordinator has a combined
// coverage map.
<原文结束>

# <翻译开始>
// warmupRun returns true while the coordinator is running inputs without
// mutating them as a warmup before fuzzing. This could be to gather baseline
// coverage data for entries in the corpus, or to test all of the seed corpus
// for errors before fuzzing begins.
//
// The coordinator doesn't store coverage data in the cache with each input
// because that data would be invalid when counter offsets in the test binary
// change.
//
// When gathering coverage, the coordinator sends each entry to a worker to
// gather coverage for that entry only, without fuzzing or minimizing. This
// phase ends when all workers have finished, and the coordinator has a combined
// coverage map.
# <翻译结束>


<原文开始>
// updateCoverage sets bits in c.coverageMask that are set in newCoverage.
// updateCoverage returns the number of newly set bits. See the comment on
// coverageMask for the format.
<原文结束>

# <翻译开始>
// updateCoverage sets bits in c.coverageMask that are set in newCoverage.
// updateCoverage returns the number of newly set bits. See the comment on
// coverageMask for the format.
# <翻译结束>


<原文开始>
// canMinimize returns whether the coordinator should attempt to find smaller
// inputs that reproduce a crash or new coverage.
<原文结束>

# <翻译开始>
// canMinimize returns whether the coordinator should attempt to find smaller
// inputs that reproduce a crash or new coverage.
# <翻译结束>


<原文开始>
// readCache creates a combined corpus from seed values and values in the cache
// (in GOCACHE/fuzz).
//
// TODO(fuzzing): need a mechanism that can remove values that
// aren't useful anymore, for example, because they have the wrong type.
<原文结束>

# <翻译开始>
// readCache creates a combined corpus from seed values and values in the cache
// (in GOCACHE/fuzz).
//
// TODO(fuzzing): need a mechanism that can remove values that
// aren't useful anymore, for example, because they have the wrong type.
# <翻译结束>


<原文开始>
			// It's okay if some files in the cache directory are malformed and
			// are not included in the corpus, but fail if it's an I/O error.
<原文结束>

# <翻译开始>
			// It's okay if some files in the cache directory are malformed and
			// are not included in the corpus, but fail if it's an I/O error.
# <翻译结束>


<原文开始>
		// TODO(jayconrod,katiehockman): consider printing some kind of warning
		// indicating the number of files which were skipped because they are
		// malformed.
<原文结束>

# <翻译开始>
		// TODO(jayconrod,katiehockman): consider printing some kind of warning
		// indicating the number of files which were skipped because they are
		// malformed.
# <翻译结束>


<原文开始>
// MalformedCorpusError is an error found while reading the corpus from the
// filesystem. All of the errors are stored in the errs list. The testing
// framework uses this to report malformed files in testdata.
<原文结束>

# <翻译开始>
// MalformedCorpusError is an error found while reading the corpus from the
// filesystem. All of the errors are stored in the errs list. The testing
// framework uses this to report malformed files in testdata.
# <翻译结束>


<原文开始>
// ReadCorpus reads the corpus from the provided dir. The returned corpus
// entries are guaranteed to match the given types. Any malformed files will
// be saved in a MalformedCorpusError and returned, along with the most recent
// error.
<原文结束>

# <翻译开始>
// ReadCorpus reads the corpus from the provided dir. The returned corpus
// entries are guaranteed to match the given types. Any malformed files will
// be saved in a MalformedCorpusError and returned, along with the most recent
// error.
# <翻译结束>


<原文开始>
		// TODO(jayconrod,katiehockman): determine when a file is a fuzzing input
		// based on its name. We should only read files created by writeToCorpus.
		// If we read ALL files, we won't be able to change the file format by
		// changing the extension. We also won't be able to add files like
		// README.txt explaining why the directory exists.
<原文结束>

# <翻译开始>
		// TODO(jayconrod,katiehockman): determine when a file is a fuzzing input
		// based on its name. We should only read files created by writeToCorpus.
		// If we read ALL files, we won't be able to change the file format by
		// changing the extension. We also won't be able to add files like
		// README.txt explaining why the directory exists.
# <翻译结束>


<原文开始>
// CheckCorpus verifies that the types in vals match the expected types
// provided.
<原文结束>

# <翻译开始>
// CheckCorpus verifies that the types in vals match the expected types
// provided.
# <翻译结束>


<原文开始>
// writeToCorpus atomically writes the given bytes to a new file in testdata. If
// the directory does not exist, it will create one. If the file already exists,
// writeToCorpus will not rewrite it. writeToCorpus sets entry.Path to the new
// file that was just written or an error if it failed.
<原文结束>

# <翻译开始>
// writeToCorpus atomically writes the given bytes to a new file in testdata. If
// the directory does not exist, it will create one. If the file already exists,
// writeToCorpus will not rewrite it. writeToCorpus sets entry.Path to the new
// file that was just written or an error if it failed.
# <翻译结束>


<原文开始>
// remove partially written file
<原文结束>

# <翻译开始>
// remove partially written file
# <翻译结束>

