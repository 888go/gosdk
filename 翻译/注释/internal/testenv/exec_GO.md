
<原文开始>
// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
# <翻译结束>


<原文开始>
// HasExec reports whether the current system can start new processes
// using os.StartProcess or (more commonly) exec.Command.
<原文结束>

# <翻译开始>
// HasExec reports whether the current system can start new processes
// using os.StartProcess or (more commonly) exec.Command.
# <翻译结束>


<原文开始>
// MustHaveExec checks that the current system can start new processes
// using os.StartProcess or (more commonly) exec.Command.
// If not, MustHaveExec calls t.Skip with an explanation.
<原文结束>

# <翻译开始>
// MustHaveExec checks that the current system can start new processes
// using os.StartProcess or (more commonly) exec.Command.
// If not, MustHaveExec calls t.Skip with an explanation.
# <翻译结束>


<原文开始>
// MustHaveExecPath checks that the current system can start the named executable
// using os.StartProcess or (more commonly) exec.Command.
// If not, MustHaveExecPath calls t.Skip with an explanation.
<原文结束>

# <翻译开始>
// MustHaveExecPath checks that the current system can start the named executable
// using os.StartProcess or (more commonly) exec.Command.
// If not, MustHaveExecPath calls t.Skip with an explanation.
# <翻译结束>


<原文开始>
// CleanCmdEnv will fill cmd.Env with the environment, excluding certain
// variables that could modify the behavior of the Go tools such as
// GODEBUG and GOTRACEBACK.
<原文结束>

# <翻译开始>
// CleanCmdEnv will fill cmd.Env with the environment, excluding certain
// variables that could modify the behavior of the Go tools such as
// GODEBUG and GOTRACEBACK.
# <翻译结束>


<原文开始>
		// Exclude GODEBUG from the environment to prevent its output
		// from breaking tests that are trying to parse other command output.
<原文结束>

# <翻译开始>
		// Exclude GODEBUG from the environment to prevent its output
		// from breaking tests that are trying to parse other command output.
# <翻译结束>


<原文开始>
// Exclude GOTRACEBACK for the same reason.
<原文结束>

# <翻译开始>
// Exclude GOTRACEBACK for the same reason.
# <翻译结束>


<原文开始>
// CommandContext is like exec.CommandContext, but:
//   - skips t if the platform does not support os/exec,
//   - sends SIGQUIT (if supported by the platform) instead of SIGKILL
//     in its Cancel function
//   - if the test has a deadline, adds a Context timeout and WaitDelay
//     for an arbitrary grace period before the test's deadline expires,
//   - fails the test if the command does not complete before the test's deadline, and
//   - sets a Cleanup function that verifies that the test did not leak a subprocess.
<原文结束>

# <翻译开始>
// CommandContext is like exec.CommandContext, but:
//   - skips t if the platform does not support os/exec,
//   - sends SIGQUIT (if supported by the platform) instead of SIGKILL
//     in its Cancel function
//   - if the test has a deadline, adds a Context timeout and WaitDelay
//     for an arbitrary grace period before the test's deadline expires,
//   - fails the test if the command does not complete before the test's deadline, and
//   - sets a Cleanup function that verifies that the test did not leak a subprocess.
# <翻译结束>


<原文开始>
// unlimited unless the test has a deadline (to allow for interactive debugging)
<原文结束>

# <翻译开始>
// unlimited unless the test has a deadline (to allow for interactive debugging)
# <翻译结束>


<原文开始>
			// Start with a minimum grace period, just long enough to consume the
			// output of a reasonable program after it terminates.
<原文结束>

# <翻译开始>
			// Start with a minimum grace period, just long enough to consume the
			// output of a reasonable program after it terminates.
# <翻译结束>


<原文开始>
			// If time allows, increase the termination grace period to 5% of the
			// test's remaining time.
<原文结束>

# <翻译开始>
			// If time allows, increase the termination grace period to 5% of the
			// test's remaining time.
# <翻译结束>


<原文开始>
			// When we run commands that execute subprocesses, we want to reserve two
			// grace periods to clean up: one for the delay between the first
			// termination signal being sent (via the Cancel callback when the Context
			// expires) and the process being forcibly terminated (via the WaitDelay
			// field), and a second one for the delay becween the process being
			// terminated and and the test logging its output for debugging.
			//
			// (We want to ensure that the test process itself has enough time to
			// log the output before it is also terminated.)
<原文结束>

# <翻译开始>
			// When we run commands that execute subprocesses, we want to reserve two
			// grace periods to clean up: one for the delay between the first
			// termination signal being sent (via the Cancel callback when the Context
			// expires) and the process being forcibly terminated (via the WaitDelay
			// field), and a second one for the delay becween the process being
			// terminated and and the test logging its output for debugging.
			//
			// (We want to ensure that the test process itself has enough time to
			// log the output before it is also terminated.)
# <翻译结束>


<原文开始>
				// Either ctx doesn't have a deadline, or its deadline would expire
				// after (or too close before) the test has already timed out.
				// Add a shorter timeout so that the test will produce useful output.
<原文结束>

# <翻译开始>
				// Either ctx doesn't have a deadline, or its deadline would expire
				// after (or too close before) the test has already timed out.
				// Add a shorter timeout so that the test will produce useful output.
# <翻译结束>


<原文开始>
			// The command is being terminated due to ctx being canceled, but
			// apparently not due to an explicit test deadline that we added.
			// Log that information in case it is useful for diagnosing a failure,
			// but don't actually fail the test because of it.
<原文结束>

# <翻译开始>
			// The command is being terminated due to ctx being canceled, but
			// apparently not due to an explicit test deadline that we added.
			// Log that information in case it is useful for diagnosing a failure,
			// but don't actually fail the test because of it.
# <翻译结束>


<原文开始>
// Command is like exec.Command, but applies the same changes as
// testenv.CommandContext (with a default Context).
<原文结束>

# <翻译开始>
// Command is like exec.Command, but applies the same changes as
// testenv.CommandContext (with a default Context).
# <翻译结束>

