package time

// The Timer type represents a single event.
// When the Timer expires, the current time will be sent on C,
// unless the Timer was created by AfterFunc.
// A Timer must be created with NewTimer or AfterFunc.
type Timer struct {

} //md5:c8aa453b91475f19

// Stop prevents the Timer from firing.
// It returns true if the call stops the timer, false if the timer has already
// expired or been stopped.
// Stop does not close the channel, to prevent a read from the channel succeeding
// incorrectly.
//
// To ensure the channel is empty after a call to Stop, check the
// return value and drain the channel.
// For example, assuming the program has not received from t.C already:
//
//	if !t.Stop() {
//		<-t.C
//	}
//
// This cannot be done concurrent to other receives from the Timer's
// channel or other calls to the Timer's Stop method.
//
// For a timer created with AfterFunc(d, f), if t.Stop returns false, then the timer
// has already expired and the function f has been started in its own goroutine;
// Stop does not wait for f to complete before returning.
// If the caller needs to know whether f is completed, it must coordinate
// with f explicitly.
func (t *Timer) Stop() bool { //md5:fc8cbe1af620cb13

}

// NewTimer creates a new Timer that will send
// the current time on its channel after at least duration d.
func NewTimer(d Duration) *Timer { //md5:11e89ccf3aaee20c

}

// Reset changes the timer to expire after duration d.
// It returns true if the timer had been active, false if the timer had
// expired or been stopped.
//
// For a Timer created with NewTimer, Reset should be invoked only on
// stopped or expired timers with drained channels.
//
// If a program has already received a value from t.C, the timer is known
// to have expired and the channel drained, so t.Reset can be used directly.
// If a program has not yet received a value from t.C, however,
// the timer must be stopped and—if Stop reports that the timer expired
// before being stopped—the channel explicitly drained:
//
//	if !t.Stop() {
//		<-t.C
//	}
//	t.Reset(d)
//
// This should not be done concurrent to other receives from the Timer's
// channel.
//
// Note that it is not possible to use Reset's return value correctly, as there
// is a race condition between draining the channel and the new timer expiring.
// Reset should always be invoked on stopped or expired channels, as described above.
// The return value exists to preserve compatibility with existing programs.
//
// For a Timer created with AfterFunc(d, f), Reset either reschedules
// when f will run, in which case Reset returns true, or schedules f
// to run again, in which case it returns false.
// When Reset returns false, Reset neither waits for the prior f to
// complete before returning nor does it guarantee that the subsequent
// goroutine running f does not run concurrently with the prior
// one. If the caller needs to know whether the prior execution of
// f is completed, it must coordinate with f explicitly.
func (t *Timer) Reset(d Duration) bool { //md5:ab76211fe40d3fa1

}

// AfterFunc waits for the duration to elapse and then calls f
// in its own goroutine. It returns a Timer that can
// be used to cancel the call using its Stop method.
// The returned Timer's C field is not used and will be nil.
func AfterFunc(d Duration, f func()) *Timer { //md5:c6cacd4aadc22dcb

}