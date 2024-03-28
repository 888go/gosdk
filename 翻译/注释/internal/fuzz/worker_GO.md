
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
	// workerFuzzDuration is the amount of time a worker can spend testing random
	// variations of an input given by the coordinator.
<原文结束>

# <翻译开始>
	// workerFuzzDuration is the amount of time a worker can spend testing random
	// variations of an input given by the coordinator.
# <翻译结束>


<原文开始>
	// workerTimeoutDuration is the amount of time a worker can go without
	// responding to the coordinator before being stopped.
<原文结束>

# <翻译开始>
	// workerTimeoutDuration is the amount of time a worker can go without
	// responding to the coordinator before being stopped.
# <翻译结束>


<原文开始>
	// workerExitCode is used as an exit code by fuzz worker processes after an internal error.
	// This distinguishes internal errors from uncontrolled panics and other crashes.
	// Keep in sync with internal/fuzz.workerExitCode.
<原文结束>

# <翻译开始>
	// workerExitCode is used as an exit code by fuzz worker processes after an internal error.
	// This distinguishes internal errors from uncontrolled panics and other crashes.
	// Keep in sync with internal/fuzz.workerExitCode.
# <翻译结束>


<原文开始>
	// workerSharedMemSize is the maximum size of the shared memory file used to
	// communicate with workers. This limits the size of fuzz inputs.
<原文结束>

# <翻译开始>
	// workerSharedMemSize is the maximum size of the shared memory file used to
	// communicate with workers. This limits the size of fuzz inputs.
# <翻译结束>


<原文开始>
// worker manages a worker process running a test binary. The worker object
// exists only in the coordinator (the process started by 'go test -fuzz').
// workerClient is used by the coordinator to send RPCs to the worker process,
// which handles them with workerServer.
<原文结束>

# <翻译开始>
// worker manages a worker process running a test binary. The worker object
// exists only in the coordinator (the process started by 'go test -fuzz').
// workerClient is used by the coordinator to send RPCs to the worker process,
// which handles them with workerServer.
# <翻译结束>


<原文开始>
// working directory, same as package directory
<原文结束>

# <翻译开始>
// working directory, same as package directory
# <翻译结束>


<原文开始>
// path to test executable
<原文结束>

# <翻译开始>
// path to test executable
# <翻译结束>


<原文开始>
// arguments for test executable
<原文结束>

# <翻译开始>
// arguments for test executable
# <翻译结束>


<原文开始>
// environment for test executable
<原文结束>

# <翻译开始>
// environment for test executable
# <翻译结束>


<原文开始>
// mutex guarding shared memory with worker; persists across processes.
<原文结束>

# <翻译开始>
// mutex guarding shared memory with worker; persists across processes.
# <翻译结束>


<原文开始>
// used to communicate with worker process
<原文结束>

# <翻译开始>
// used to communicate with worker process
# <翻译结束>


<原文开始>
// last error returned by wait, set before termC is closed.
<原文结束>

# <翻译开始>
// last error returned by wait, set before termC is closed.
# <翻译结束>


<原文开始>
// true after stop interrupts a running worker.
<原文结束>

# <翻译开始>
// true after stop interrupts a running worker.
# <翻译结束>


<原文开始>
// closed by wait when worker process terminates
<原文结束>

# <翻译开始>
// closed by wait when worker process terminates
# <翻译结束>


<原文开始>
// cleanup releases persistent resources associated with the worker.
<原文结束>

# <翻译开始>
// cleanup releases persistent resources associated with the worker.
# <翻译结束>


<原文开始>
// coordinate runs the test binary to perform fuzzing.
//
// coordinate loops until ctx is cancelled or a fatal error is encountered.
// If a test process terminates unexpectedly while fuzzing, coordinate will
// attempt to restart and continue unless the termination can be attributed
// to an interruption (from a timer or the user).
//
// While looping, coordinate receives inputs from the coordinator, passes
// those inputs to the worker process, then passes the results back to
// the coordinator.
<原文结束>

# <翻译开始>
// coordinate runs the test binary to perform fuzzing.
//
// coordinate loops until ctx is cancelled or a fatal error is encountered.
// If a test process terminates unexpectedly while fuzzing, coordinate will
// attempt to restart and continue unless the termination can be attributed
// to an interruption (from a timer or the user).
//
// While looping, coordinate receives inputs from the coordinator, passes
// those inputs to the worker process, then passes the results back to
// the coordinator.
# <翻译结束>


<原文开始>
// Start or restart the worker if it's not running.
<原文结束>

# <翻译开始>
// Start or restart the worker if it's not running.
# <翻译结束>


<原文开始>
// Worker was told to stop.
<原文结束>

# <翻译开始>
// Worker was told to stop.
# <翻译结束>


<原文开始>
// Worker process terminated unexpectedly while waiting for input.
<原文结束>

# <翻译开始>
// Worker process terminated unexpectedly while waiting for input.
# <翻译结束>


<原文开始>
				// Worker stopped, either by exiting with status 0 or after being
				// interrupted with a signal that was not sent by the coordinator.
				//
				// When the user presses ^C, on POSIX platforms, SIGINT is delivered to
				// all processes in the group concurrently, and the worker may see it
				// before the coordinator. The worker should exit 0 gracefully (in
				// theory).
				//
				// This condition is probably intended by the user, so suppress
				// the error.
<原文结束>

# <翻译开始>
				// Worker stopped, either by exiting with status 0 or after being
				// interrupted with a signal that was not sent by the coordinator.
				//
				// When the user presses ^C, on POSIX platforms, SIGINT is delivered to
				// all processes in the group concurrently, and the worker may see it
				// before the coordinator. The worker should exit 0 gracefully (in
				// theory).
				//
				// This condition is probably intended by the user, so suppress
				// the error.
# <翻译结束>


<原文开始>
				// Worker exited with a code indicating F.Fuzz was not called correctly,
				// for example, F.Fail was called first.
<原文结束>

# <翻译开始>
				// Worker exited with a code indicating F.Fuzz was not called correctly,
				// for example, F.Fail was called first.
# <翻译结束>


<原文开始>
			// Worker exited non-zero or was terminated by a non-interrupt
			// signal (for example, SIGSEGV) while fuzzing.
<原文结束>

# <翻译开始>
			// Worker exited non-zero or was terminated by a non-interrupt
			// signal (for example, SIGSEGV) while fuzzing.
# <翻译结束>


<原文开始>
// TODO(jayconrod,katiehockman): if -keepfuzzing, restart worker.
<原文结束>

# <翻译开始>
// TODO(jayconrod,katiehockman): if -keepfuzzing, restart worker.
# <翻译结束>


<原文开始>
// Received input from coordinator.
<原文结束>

# <翻译开始>
// Received input from coordinator.
# <翻译结束>


<原文开始>
// Error communicating with worker.
<原文结束>

# <翻译开始>
// Error communicating with worker.
# <翻译结束>


<原文开始>
// Timeout or interruption.
<原文结束>

# <翻译开始>
// Timeout or interruption.
# <翻译结束>


<原文开始>
					// Communication error before we stopped the worker.
					// Report an error, but don't record a crasher.
<原文结束>

# <翻译开始>
					// Communication error before we stopped the worker.
					// Report an error, but don't record a crasher.
# <翻译结束>


<原文开始>
					// Worker terminated by a signal that probably wasn't caused by a
					// specific input to the fuzz function. For example, on Linux,
					// the kernel (OOM killer) may send SIGKILL to a process using a lot
					// of memory. Or the shell might send SIGHUP when the terminal
					// is closed. Don't record a crasher.
<原文结束>

# <翻译开始>
					// Worker terminated by a signal that probably wasn't caused by a
					// specific input to the fuzz function. For example, on Linux,
					// the kernel (OOM killer) may send SIGKILL to a process using a lot
					// of memory. Or the shell might send SIGHUP when the terminal
					// is closed. Don't record a crasher.
# <翻译结束>


<原文开始>
					// An internal error occurred which shouldn't be considered
					// a crash.
<原文结束>

# <翻译开始>
					// An internal error occurred which shouldn't be considered
					// a crash.
# <翻译结束>


<原文开始>
				// Unexpected termination. Set error message and fall through.
				// We'll restart the worker on the next iteration.
				// Don't attempt to minimize this since it crashed the worker.
<原文结束>

# <翻译开始>
				// Unexpected termination. Set error message and fall through.
				// We'll restart the worker on the next iteration.
				// Don't attempt to minimize this since it crashed the worker.
# <翻译结束>


<原文开始>
// Received input to minimize from coordinator.
<原文结束>

# <翻译开始>
// Received input to minimize from coordinator.
# <翻译结束>


<原文开始>
				// Error minimizing. Send back the original input. If it didn't cause
				// an error before, report it as causing an error now.
				// TODO: double-check this is handled correctly when
				// implementing -keepfuzzing.
<原文结束>

# <翻译开始>
				// Error minimizing. Send back the original input. If it didn't cause
				// an error before, report it as causing an error now.
				// TODO: double-check this is handled correctly when
				// implementing -keepfuzzing.
# <翻译结束>


<原文开始>
// minimize tells a worker process to attempt to find a smaller value that
// either causes an error (if we started minimizing because we found an input
// that causes an error) or preserves new coverage (if we started minimizing
// because we found an input that expands coverage).
<原文结束>

# <翻译开始>
// minimize tells a worker process to attempt to find a smaller value that
// either causes an error (if we started minimizing because we found an input
// that causes an error) or preserves new coverage (if we started minimizing
// because we found an input that expands coverage).
# <翻译结束>


<原文开始>
			// Worker was interrupted, possibly by the user pressing ^C.
			// Normally, workers can handle interrupts and timeouts gracefully and
			// will return without error. An error here indicates the worker
			// may not have been in a good state, but the error won't be meaningful
			// to the user. Just return the original crasher without logging anything.
<原文结束>

# <翻译开始>
			// Worker was interrupted, possibly by the user pressing ^C.
			// Normally, workers can handle interrupts and timeouts gracefully and
			// will return without error. An error here indicates the worker
			// may not have been in a good state, but the error won't be meaningful
			// to the user. Just return the original crasher without logging anything.
# <翻译结束>


<原文开始>
// startAndPing starts the worker process and sends it a message to make sure it
// can communicate.
//
// startAndPing returns an error if any part of this didn't work, including if
// the context is expired or the worker process was interrupted before it
// responded. Errors that happen after start but before the ping response
// likely indicate that the worker did not call F.Fuzz or called F.Fail first.
// We don't record crashers for these errors.
<原文结束>

# <翻译开始>
// startAndPing starts the worker process and sends it a message to make sure it
// can communicate.
//
// startAndPing returns an error if any part of this didn't work, including if
// the context is expired or the worker process was interrupted before it
// responded. Errors that happen after start but before the ping response
// likely indicate that the worker did not call F.Fuzz or called F.Fail first.
// We don't record crashers for these errors.
# <翻译结束>


<原文开始>
// User may have pressed ^C before worker responded.
<原文结束>

# <翻译开始>
// User may have pressed ^C before worker responded.
# <翻译结束>


<原文开始>
// TODO: record and return stderr.
<原文结束>

# <翻译开始>
// TODO: record and return stderr.
# <翻译结束>


<原文开始>
// start runs a new worker process.
//
// If the process couldn't be started, start returns an error. Start won't
// return later termination errors from the process if they occur.
//
// If the process starts successfully, start returns nil. stop must be called
// once later to clean up, even if the process terminates on its own.
//
// When the process terminates, w.waitErr is set to the error (if any), and
// w.termC is closed.
<原文结束>

# <翻译开始>
// start runs a new worker process.
//
// If the process couldn't be started, start returns an error. Start won't
// return later termination errors from the process if they occur.
//
// If the process starts successfully, start returns nil. stop must be called
// once later to clean up, even if the process terminates on its own.
//
// When the process terminates, w.waitErr is set to the error (if any), and
// w.termC is closed.
# <翻译结束>


<原文开始>
	// Create the "fuzz_in" and "fuzz_out" pipes so we can communicate with
	// the worker. We don't use stdin and stdout, since the test binary may
	// do something else with those.
	//
	// Each pipe has a reader and a writer. The coordinator writes to fuzzInW
	// and reads from fuzzOutR. The worker inherits fuzzInR and fuzzOutW.
	// The coordinator closes fuzzInR and fuzzOutW after starting the worker,
	// since we have no further need of them.
<原文结束>

# <翻译开始>
	// Create the "fuzz_in" and "fuzz_out" pipes so we can communicate with
	// the worker. We don't use stdin and stdout, since the test binary may
	// do something else with those.
	//
	// Each pipe has a reader and a writer. The coordinator writes to fuzzInW
	// and reads from fuzzOutR. The worker inherits fuzzInR and fuzzOutW.
	// The coordinator closes fuzzInR and fuzzOutW after starting the worker,
	// since we have no further need of them.
# <翻译结束>


<原文开始>
// Start the worker process.
<原文结束>

# <翻译开始>
// Start the worker process.
# <翻译结束>


<原文开始>
	// Worker started successfully.
	// After this, w.client owns fuzzInW and fuzzOutR, so w.client.Close must be
	// called later by stop.
<原文结束>

# <翻译开始>
	// Worker started successfully.
	// After this, w.client owns fuzzInW and fuzzOutR, so w.client.Close must be
	// called later by stop.
# <翻译结束>


<原文开始>
// stop tells the worker process to exit by closing w.client, then blocks until
// it terminates. If the worker doesn't terminate after a short time, stop
// signals it with os.Interrupt (where supported), then os.Kill.
//
// stop returns the error the process terminated with, if any (same as
// w.waitErr).
//
// stop must be called at least once after start returns successfully, even if
// the worker process terminates unexpectedly.
<原文结束>

# <翻译开始>
// stop tells the worker process to exit by closing w.client, then blocks until
// it terminates. If the worker doesn't terminate after a short time, stop
// signals it with os.Interrupt (where supported), then os.Kill.
//
// stop returns the error the process terminated with, if any (same as
// w.waitErr).
//
// stop must be called at least once after start returns successfully, even if
// the worker process terminates unexpectedly.
# <翻译结束>


<原文开始>
// Worker already terminated.
<原文结束>

# <翻译开始>
// Worker already terminated.
# <翻译结束>


<原文开始>
// Possible unexpected termination.
<原文结束>

# <翻译开始>
// Possible unexpected termination.
# <翻译结束>


<原文开始>
	// Tell the worker to stop by closing fuzz_in. It won't actually stop until it
	// finishes with earlier calls.
<原文结束>

# <翻译开始>
	// Tell the worker to stop by closing fuzz_in. It won't actually stop until it
	// finishes with earlier calls.
# <翻译结束>


<原文开始>
// Timer fired before worker terminated.
<原文结束>

# <翻译开始>
// Timer fired before worker terminated.
# <翻译结束>


<原文开始>
// Try to stop the worker with SIGINT and wait a little longer.
<原文结束>

# <翻译开始>
// Try to stop the worker with SIGINT and wait a little longer.
# <翻译结束>


<原文开始>
// Try to stop the worker with SIGKILL and keep waiting.
<原文结束>

# <翻译开始>
// Try to stop the worker with SIGKILL and keep waiting.
# <翻译结束>


<原文开始>
// Still waiting. Print a message to let the user know why.
<原文结束>

# <翻译开始>
// Still waiting. Print a message to let the user know why.
# <翻译结束>


<原文开始>
// RunFuzzWorker is called in a worker process to communicate with the
// coordinator process in order to fuzz random inputs. RunFuzzWorker loops
// until the coordinator tells it to stop.
//
// fn is a wrapper on the fuzz function. It may return an error to indicate
// a given input "crashed". The coordinator will also record a crasher if
// the function times out or terminates the process.
//
// RunFuzzWorker returns an error if it could not communicate with the
// coordinator process.
<原文结束>

# <翻译开始>
// RunFuzzWorker is called in a worker process to communicate with the
// coordinator process in order to fuzz random inputs. RunFuzzWorker loops
// until the coordinator tells it to stop.
//
// fn is a wrapper on the fuzz function. It may return an error to indicate
// a given input "crashed". The coordinator will also record a crasher if
// the function times out or terminates the process.
//
// RunFuzzWorker returns an error if it could not communicate with the
// coordinator process.
# <翻译结束>


<原文开始>
// call is serialized and sent from the coordinator on fuzz_in. It acts as
// a minimalist RPC mechanism. Exactly one of its fields must be set to indicate
// which method to call.
<原文结束>

# <翻译开始>
// call is serialized and sent from the coordinator on fuzz_in. It acts as
// a minimalist RPC mechanism. Exactly one of its fields must be set to indicate
// which method to call.
# <翻译结束>


<原文开始>
// minimizeArgs contains arguments to workerServer.minimize. The value to
// minimize is already in shared memory.
<原文结束>

# <翻译开始>
// minimizeArgs contains arguments to workerServer.minimize. The value to
// minimize is already in shared memory.
# <翻译结束>


<原文开始>
	// Timeout is the time to spend minimizing. This may include time to start up,
	// especially if the input causes the worker process to terminated, requiring
	// repeated restarts.
<原文结束>

# <翻译开始>
	// Timeout is the time to spend minimizing. This may include time to start up,
	// especially if the input causes the worker process to terminated, requiring
	// repeated restarts.
# <翻译结束>


<原文开始>
	// Limit is the maximum number of values to test, without spending more time
	// than Duration. 0 indicates no limit.
<原文结束>

# <翻译开始>
	// Limit is the maximum number of values to test, without spending more time
	// than Duration. 0 indicates no limit.
# <翻译结束>


<原文开始>
	// KeepCoverage is a set of coverage counters the worker should attempt to
	// keep in minimized values. When provided, the worker will reject inputs that
	// don't cause at least one of these bits to be set.
<原文结束>

# <翻译开始>
	// KeepCoverage is a set of coverage counters the worker should attempt to
	// keep in minimized values. When provided, the worker will reject inputs that
	// don't cause at least one of these bits to be set.
# <翻译结束>


<原文开始>
// Index is the index of the fuzz target parameter to be minimized.
<原文结束>

# <翻译开始>
// Index is the index of the fuzz target parameter to be minimized.
# <翻译结束>


<原文开始>
// minimizeResponse contains results from workerServer.minimize.
<原文结束>

# <翻译开始>
// minimizeResponse contains results from workerServer.minimize.
# <翻译结束>


<原文开始>
	// WroteToMem is true if the worker found a smaller input and wrote it to
	// shared memory. If minimizeArgs.KeepCoverage was set, the minimized input
	// preserved at least one coverage bit and did not cause an error.
	// Otherwise, the minimized input caused some error, recorded in Err.
<原文结束>

# <翻译开始>
	// WroteToMem is true if the worker found a smaller input and wrote it to
	// shared memory. If minimizeArgs.KeepCoverage was set, the minimized input
	// preserved at least one coverage bit and did not cause an error.
	// Otherwise, the minimized input caused some error, recorded in Err.
# <翻译结束>


<原文开始>
// Err is the error string caused by the value in shared memory, if any.
<原文结束>

# <翻译开始>
// Err is the error string caused by the value in shared memory, if any.
# <翻译结束>


<原文开始>
	// CoverageData is the set of coverage bits activated by the minimized value
	// in shared memory. When set, it contains at least one bit from KeepCoverage.
	// CoverageData will be nil if Err is set or if minimization failed.
<原文结束>

# <翻译开始>
	// CoverageData is the set of coverage bits activated by the minimized value
	// in shared memory. When set, it contains at least one bit from KeepCoverage.
	// CoverageData will be nil if Err is set or if minimization failed.
# <翻译结束>


<原文开始>
// Duration is the time spent minimizing, not including starting or cleaning up.
<原文结束>

# <翻译开始>
// Duration is the time spent minimizing, not including starting or cleaning up.
# <翻译结束>


<原文开始>
// Count is the number of values tested.
<原文结束>

# <翻译开始>
// Count is the number of values tested.
# <翻译结束>


<原文开始>
// fuzzArgs contains arguments to workerServer.fuzz. The value to fuzz is
// passed in shared memory.
<原文结束>

# <翻译开始>
// fuzzArgs contains arguments to workerServer.fuzz. The value to fuzz is
// passed in shared memory.
# <翻译结束>


<原文开始>
	// Timeout is the time to spend fuzzing, not including starting or
	// cleaning up.
<原文结束>

# <翻译开始>
	// Timeout is the time to spend fuzzing, not including starting or
	// cleaning up.
# <翻译结束>


<原文开始>
	// Warmup indicates whether this is part of a warmup run, meaning that
	// fuzzing should not occur. If coverageEnabled is true, then coverage data
	// should be reported.
<原文结束>

# <翻译开始>
	// Warmup indicates whether this is part of a warmup run, meaning that
	// fuzzing should not occur. If coverageEnabled is true, then coverage data
	// should be reported.
# <翻译结束>


<原文开始>
	// CoverageData is the coverage data. If set, the worker should update its
	// local coverage data prior to fuzzing.
<原文结束>

# <翻译开始>
	// CoverageData is the coverage data. If set, the worker should update its
	// local coverage data prior to fuzzing.
# <翻译结束>


<原文开始>
// fuzzResponse contains results from workerServer.fuzz.
<原文结束>

# <翻译开始>
// fuzzResponse contains results from workerServer.fuzz.
# <翻译结束>


<原文开始>
// Duration is the time spent fuzzing, not including starting or cleaning up.
<原文结束>

# <翻译开始>
// Duration is the time spent fuzzing, not including starting or cleaning up.
# <翻译结束>


<原文开始>
	// CoverageData is set if the value in shared memory expands coverage
	// and therefore may be interesting to the coordinator.
<原文结束>

# <翻译开始>
	// CoverageData is set if the value in shared memory expands coverage
	// and therefore may be interesting to the coordinator.
# <翻译结束>


<原文开始>
	// Err is the error string caused by the value in shared memory, which is
	// non-empty if the value in shared memory caused a crash.
<原文结束>

# <翻译开始>
	// Err is the error string caused by the value in shared memory, which is
	// non-empty if the value in shared memory caused a crash.
# <翻译结束>


<原文开始>
	// InternalErr is the error string caused by an internal error in the
	// worker. This shouldn't be considered a crasher.
<原文结束>

# <翻译开始>
	// InternalErr is the error string caused by an internal error in the
	// worker. This shouldn't be considered a crasher.
# <翻译结束>


<原文开始>
// pingArgs contains arguments to workerServer.ping.
<原文结束>

# <翻译开始>
// pingArgs contains arguments to workerServer.ping.
# <翻译结束>


<原文开始>
// pingResponse contains results from workerServer.ping.
<原文结束>

# <翻译开始>
// pingResponse contains results from workerServer.ping.
# <翻译结束>


<原文开始>
// workerComm holds pipes and shared memory used for communication
// between the coordinator process (client) and a worker process (server).
// These values are unique to each worker; they are shared only with the
// coordinator, not with other workers.
//
// Access to shared memory is synchronized implicitly over the RPC protocol
// implemented in workerServer and workerClient. During a call, the client
// (worker) has exclusive access to shared memory; at other times, the server
// (coordinator) has exclusive access.
<原文结束>

# <翻译开始>
// workerComm holds pipes and shared memory used for communication
// between the coordinator process (client) and a worker process (server).
// These values are unique to each worker; they are shared only with the
// coordinator, not with other workers.
//
// Access to shared memory is synchronized implicitly over the RPC protocol
// implemented in workerServer and workerClient. During a call, the client
// (worker) has exclusive access to shared memory; at other times, the server
// (coordinator) has exclusive access.
# <翻译结束>


<原文开始>
// workerServer is a minimalist RPC server, run by fuzz worker processes.
// It allows the coordinator process (using workerClient) to call methods in a
// worker process. This system allows the coordinator to run multiple worker
// processes in parallel and to collect inputs that caused crashes from shared
// memory after a worker process terminates unexpectedly.
<原文结束>

# <翻译开始>
// workerServer is a minimalist RPC server, run by fuzz worker processes.
// It allows the coordinator process (using workerClient) to call methods in a
// worker process. This system allows the coordinator to run multiple worker
// processes in parallel and to collect inputs that caused crashes from shared
// memory after a worker process terminates unexpectedly.
# <翻译结束>


<原文开始>
	// coverageMask is the local coverage data for the worker. It is
	// periodically updated to reflect the data in the coordinator when new
	// coverage is found.
<原文结束>

# <翻译开始>
	// coverageMask is the local coverage data for the worker. It is
	// periodically updated to reflect the data in the coordinator when new
	// coverage is found.
# <翻译结束>


<原文开始>
	// fuzzFn runs the worker's fuzz target on the given input and returns an
	// error if it finds a crasher (the process may also exit or crash), and the
	// time it took to run the input. It sets a deadline of 10 seconds, at which
	// point it will panic with the assumption that the process is hanging or
	// deadlocked.
<原文结束>

# <翻译开始>
	// fuzzFn runs the worker's fuzz target on the given input and returns an
	// error if it finds a crasher (the process may also exit or crash), and the
	// time it took to run the input. It sets a deadline of 10 seconds, at which
	// point it will panic with the assumption that the process is hanging or
	// deadlocked.
# <翻译结束>


<原文开始>
// serve reads serialized RPC messages on fuzzIn. When serve receives a message,
// it calls the corresponding method, then sends the serialized result back
// on fuzzOut.
//
// serve handles RPC calls synchronously; it will not attempt to read a message
// until the previous call has finished.
//
// serve returns errors that occurred when communicating over pipes. serve
// does not return errors from method calls; those are passed through serialized
// responses.
<原文结束>

# <翻译开始>
// serve reads serialized RPC messages on fuzzIn. When serve receives a message,
// it calls the corresponding method, then sends the serialized result back
// on fuzzOut.
//
// serve handles RPC calls synchronously; it will not attempt to read a message
// until the previous call has finished.
//
// serve returns errors that occurred when communicating over pipes. serve
// does not return errors from method calls; those are passed through serialized
// responses.
# <翻译结束>


<原文开始>
// chainedMutations is how many mutations are applied before the worker
// resets the input to it's original state.
// NOTE: this number was picked without much thought. It is low enough that
// it seems to create a significant diversity in mutated inputs. We may want
// to consider looking into this more closely once we have a proper performance
// testing framework. Another option is to randomly pick the number of chained
// mutations on each invocation of the workerServer.fuzz method (this appears to
// be what libFuzzer does, although there seems to be no documentation which
// explains why this choice was made.)
<原文结束>

# <翻译开始>
// chainedMutations is how many mutations are applied before the worker
// resets the input to it's original state.
// NOTE: this number was picked without much thought. It is low enough that
// it seems to create a significant diversity in mutated inputs. We may want
// to consider looking into this more closely once we have a proper performance
// testing framework. Another option is to randomly pick the number of chained
// mutations on each invocation of the workerServer.fuzz method (this appears to
// be what libFuzzer does, although there seems to be no documentation which
// explains why this choice was made.)
# <翻译结束>


<原文开始>
// fuzz runs the test function on random variations of the input value in shared
// memory for a limited duration or number of iterations.
//
// fuzz returns early if it finds an input that crashes the fuzz function (with
// fuzzResponse.Err set) or an input that expands coverage (with
// fuzzResponse.InterestingDuration set).
//
// fuzz does not modify the input in shared memory. Instead, it saves the
// initial PRNG state in shared memory and increments a counter in shared
// memory before each call to the test function. The caller may reconstruct
// the crashing input with this information, since the PRNG is deterministic.
<原文结束>

# <翻译开始>
// fuzz runs the test function on random variations of the input value in shared
// memory for a limited duration or number of iterations.
//
// fuzz returns early if it finds an input that crashes the fuzz function (with
// fuzzResponse.Err set) or an input that expands coverage (with
// fuzzResponse.InterestingDuration set).
//
// fuzz does not modify the input in shared memory. Instead, it saves the
// initial PRNG state in shared memory and increments a counter in shared
// memory before each call to the test function. The caller may reconstruct
// the crashing input with this information, since the PRNG is deterministic.
# <翻译结束>


<原文开始>
	// Minimize the values in vals, then write to shared memory. We only write
	// to shared memory after completing minimization.
<原文结束>

# <翻译开始>
	// Minimize the values in vals, then write to shared memory. We only write
	// to shared memory after completing minimization.
# <翻译结束>


<原文开始>
			// If the values didn't change during minimization then coverageSnapshot is likely
			// a dirty snapshot which represents the very last step of minimization, not the
			// coverage for the initial input. In that case just return the coverage we were
			// given initially, since it more accurately represents the coverage map for the
			// input we are returning.
<原文结束>

# <翻译开始>
			// If the values didn't change during minimization then coverageSnapshot is likely
			// a dirty snapshot which represents the very last step of minimization, not the
			// coverage for the initial input. In that case just return the coverage we were
			// given initially, since it more accurately represents the coverage map for the
			// input we are returning.
# <翻译结束>


<原文开始>
// minimizeInput applies a series of minimizing transformations on the provided
// vals, ensuring that each minimization still causes an error, or keeps
// coverage, in fuzzFn. It uses the context to determine how long to run,
// stopping once closed. It returns a bool indicating whether minimization was
// successful and an error if one was found.
<原文结束>

# <翻译开始>
// minimizeInput applies a series of minimizing transformations on the provided
// vals, ensuring that each minimization still causes an error, or keeps
// coverage, in fuzzFn. It uses the context to determine how long to run,
// stopping once closed. It returns a bool indicating whether minimization was
// successful and an error if one was found.
# <翻译结束>


<原文开始>
	// Check that the original value preserves coverage or causes an error.
	// If not, then whatever caused us to think the value was interesting may
	// have been a flake, and we can't minimize it.
<原文结束>

# <翻译开始>
	// Check that the original value preserves coverage or causes an error.
	// If not, then whatever caused us to think the value was interesting may
	// have been a flake, and we can't minimize it.
# <翻译结束>


<原文开始>
	// tryMinimized runs the fuzz function with candidate replacing the value
	// at index valI. tryMinimized returns whether the input with candidate is
	// interesting for the same reason as the original input: it returns
	// an error if one was expected, or it preserves coverage.
<原文结束>

# <翻译开始>
	// tryMinimized runs the fuzz function with candidate replacing the value
	// at index valI. tryMinimized returns whether the input with candidate is
	// interesting for the same reason as the original input: it returns
	// an error if one was expected, or it preserves coverage.
# <翻译结束>


<原文开始>
				// Now that we've found a crash, that's more important than any
				// minimization of interesting inputs that was being done. Clear out
				// keepCoverage to only minimize the crash going forward.
<原文结束>

# <翻译开始>
				// Now that we've found a crash, that's more important than any
				// minimization of interesting inputs that was being done. Clear out
				// keepCoverage to only minimize the crash going forward.
# <翻译结束>


<原文开始>
// Minimization should preserve coverage bits.
<原文结束>

# <翻译开始>
// Minimization should preserve coverage bits.
# <翻译结束>


<原文开始>
// ping does nothing. The coordinator calls this method to ensure the worker
// has called F.Fuzz and can communicate.
<原文结束>

# <翻译开始>
// ping does nothing. The coordinator calls this method to ensure the worker
// has called F.Fuzz and can communicate.
# <翻译结束>


<原文开始>
// workerClient is a minimalist RPC client. The coordinator process uses a
// workerClient to call methods in each worker process (handled by
// workerServer).
<原文结束>

# <翻译开始>
// workerClient is a minimalist RPC client. The coordinator process uses a
// workerClient to call methods in each worker process (handled by
// workerServer).
# <翻译结束>


<原文开始>
	// mu is the mutex protecting the workerComm.fuzzIn pipe. This must be
	// locked before making calls to the workerServer. It prevents
	// workerClient.Close from closing fuzzIn while workerClient methods are
	// writing to it concurrently, and prevents multiple callers from writing to
	// fuzzIn concurrently.
<原文结束>

# <翻译开始>
	// mu is the mutex protecting the workerComm.fuzzIn pipe. This must be
	// locked before making calls to the workerServer. It prevents
	// workerClient.Close from closing fuzzIn while workerClient methods are
	// writing to it concurrently, and prevents multiple callers from writing to
	// fuzzIn concurrently.
# <翻译结束>


<原文开始>
// Close shuts down the connection to the RPC server (the worker process) by
// closing fuzz_in. Close drains fuzz_out (avoiding a SIGPIPE in the worker),
// and closes it after the worker process closes the other end.
<原文结束>

# <翻译开始>
// Close shuts down the connection to the RPC server (the worker process) by
// closing fuzz_in. Close drains fuzz_out (avoiding a SIGPIPE in the worker),
// and closes it after the worker process closes the other end.
# <翻译结束>


<原文开始>
	// Close fuzzIn. This signals to the server that there are no more calls,
	// and it should exit.
<原文结束>

# <翻译开始>
	// Close fuzzIn. This signals to the server that there are no more calls,
	// and it should exit.
# <翻译结束>


<原文开始>
	// Drain fuzzOut and close it. When the server exits, the kernel will close
	// its end of fuzzOut, and we'll get EOF.
<原文结束>

# <翻译开始>
	// Drain fuzzOut and close it. When the server exits, the kernel will close
	// its end of fuzzOut, and we'll get EOF.
# <翻译结束>


<原文开始>
// errSharedMemClosed is returned by workerClient methods that cannot access
// shared memory because it was closed and unmapped by another goroutine. That
// can happen when worker.cleanup is called in the worker goroutine while a
// workerClient.fuzz call runs concurrently.
//
// This error should not be reported. It indicates the operation was
// interrupted.
<原文结束>

# <翻译开始>
// errSharedMemClosed is returned by workerClient methods that cannot access
// shared memory because it was closed and unmapped by another goroutine. That
// can happen when worker.cleanup is called in the worker goroutine while a
// workerClient.fuzz call runs concurrently.
//
// This error should not be reported. It indicates the operation was
// interrupted.
# <翻译结束>


<原文开始>
// minimize tells the worker to call the minimize method. See
// workerServer.minimize.
<原文结束>

# <翻译开始>
// minimize tells the worker to call the minimize method. See
// workerServer.minimize.
# <翻译结束>


<原文开始>
// An unrecoverable error occurred before minimization began.
<原文结束>

# <翻译开始>
// An unrecoverable error occurred before minimization began.
# <翻译结束>


<原文开始>
			// An unrecoverable error occurred during minimization. mem now
			// holds the raw, unmarshalled bytes of entryIn.Values[i] that
			// caused the error.
<原文结束>

# <翻译开始>
			// An unrecoverable error occurred during minimization. mem now
			// holds the raw, unmarshalled bytes of entryIn.Values[i] that
			// caused the error.
# <翻译结束>


<原文开始>
// Stop minimizing; another unrecoverable error is likely to occur.
<原文结束>

# <翻译开始>
// Stop minimizing; another unrecoverable error is likely to occur.
# <翻译结束>


<原文开始>
// Minimization succeeded, and mem holds the marshaled data.
<原文结束>

# <翻译开始>
// Minimization succeeded, and mem holds the marshaled data.
# <翻译结束>


<原文开始>
// Prepare for next iteration of the loop.
<原文结束>

# <翻译开始>
// Prepare for next iteration of the loop.
# <翻译结束>


<原文开始>
// fuzz tells the worker to call the fuzz method. See workerServer.fuzz.
<原文结束>

# <翻译开始>
// fuzz tells the worker to call the fuzz method. See workerServer.fuzz.
# <翻译结束>


<原文开始>
// Only mutate the valuesOut if fuzzing actually occurred.
<原文结束>

# <翻译开始>
// Only mutate the valuesOut if fuzzing actually occurred.
# <翻译结束>


<原文开始>
			// The bytes weren't mutated, so if entryIn was a seed corpus value,
			// then entryOut is too.
<原文结束>

# <翻译开始>
			// The bytes weren't mutated, so if entryIn was a seed corpus value,
			// then entryOut is too.
# <翻译结束>


<原文开始>
// ping tells the worker to call the ping method. See workerServer.ping.
<原文结束>

# <翻译开始>
// ping tells the worker to call the ping method. See workerServer.ping.
# <翻译结束>


<原文开始>
// callLocked sends an RPC from the coordinator to the worker process and waits
// for the response. The callLocked may be cancelled with ctx.
<原文结束>

# <翻译开始>
// callLocked sends an RPC from the coordinator to the worker process and waits
// for the response. The callLocked may be cancelled with ctx.
# <翻译结束>


<原文开始>
// contextReader wraps a Reader with a Context. If the context is cancelled
// while the underlying reader is blocked, Read returns immediately.
//
// This is useful for reading from a pipe. Closing a pipe file descriptor does
// not unblock pending Reads on that file descriptor. All copies of the pipe's
// other file descriptor (the write end) must be closed in all processes that
// inherit it. This is difficult to do correctly in the situation we care about
// (process group termination).
<原文结束>

# <翻译开始>
// contextReader wraps a Reader with a Context. If the context is cancelled
// while the underlying reader is blocked, Read returns immediately.
//
// This is useful for reading from a pipe. Closing a pipe file descriptor does
// not unblock pending Reads on that file descriptor. All copies of the pipe's
// other file descriptor (the write end) must be closed in all processes that
// inherit it. This is difficult to do correctly in the situation we care about
// (process group termination).
# <翻译结束>


<原文开始>
	// This goroutine may stay blocked after Read returns because the underlying
	// read is blocked.
<原文结束>

# <翻译开始>
	// This goroutine may stay blocked after Read returns because the underlying
	// read is blocked.
# <翻译结束>

