
<原文开始>
// Ranger returns a Sender and a Receiver. The Receiver provides a
// Next method to retrieve values. The Sender provides a Send method
// to send values and a Close method to stop sending values. The Next
// method indicates when the Sender has been closed, and the Send
// method indicates when the Receiver has been freed.
//
// This is a convenient way to exit a goroutine sending values when
// the receiver stops reading them.
<原文结束>

# <翻译开始>
// Ranger returns a Sender and a Receiver. The Receiver provides a
// Next method to retrieve values. The Sender provides a Send method
// to send values and a Close method to stop sending values. The Next
// method indicates when the Sender has been closed, and the Send
// method indicates when the Receiver has been freed.
//
// This is a convenient way to exit a goroutine sending values when
// the receiver stops reading them.
# <翻译结束>


<原文开始>
// A sender is used to send values to a Receiver.
<原文结束>

# <翻译开始>
// A sender is used to send values to a Receiver.
# <翻译结束>


<原文开始>
// Send sends a value to the receiver. It returns whether any more
// values may be sent; if it returns false the value was not sent.
<原文结束>

# <翻译开始>
// Send sends a value to the receiver. It returns whether any more
// values may be sent; if it returns false the value was not sent.
# <翻译结束>


<原文开始>
// Close tells the receiver that no more values will arrive.
// After Close is called, the Sender may no longer be used.
<原文结束>

# <翻译开始>
// Close tells the receiver that no more values will arrive.
// After Close is called, the Sender may no longer be used.
# <翻译结束>


<原文开始>
// A Receiver receives values from a Sender.
<原文结束>

# <翻译开始>
// A Receiver receives values from a Sender.
# <翻译结束>


<原文开始>
// Next returns the next value from the channel. The bool result
// indicates whether the value is valid, or whether the Sender has
// been closed and no more values will be received.
<原文结束>

# <翻译开始>
// Next returns the next value from the channel. The bool result
// indicates whether the value is valid, or whether the Sender has
// been closed and no more values will be received.
# <翻译结束>


<原文开始>
// finalize is a finalizer for the receiver.
<原文结束>

# <翻译开始>
// finalize is a finalizer for the receiver.
# <翻译结束>

