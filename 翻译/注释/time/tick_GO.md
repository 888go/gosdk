
<原文开始>
// NewTicker returns a new Ticker containing a channel that will send
// the current time on the channel after each tick. The period of the
// ticks is specified by the duration argument. The ticker will adjust
// the time interval or drop ticks to make up for slow receivers.
// The duration d must be greater than zero; if not, NewTicker will
// panic. Stop the ticker to release associated resources.
<原文结束>

# <翻译开始>
// NewTicker 返回一个新的 Ticker，其中包含一个通道，该通道将在每次滴答声后发送当前时间。
// 滴答的周期由 duration 参数指定。如果接收者接收较慢，ticker 会调整时间间隔或丢弃滴答信号来补偿。
// 参数 d 必须大于零；如果不满足此条件，NewTicker 将引发恐慌。停止 ticker 以释放相关资源。
// md5:8d80705ea00f018e
# <翻译结束>


<原文开始>
// Stop turns off a ticker. After Stop, no more ticks will be sent.
// Stop does not close the channel, to prevent a concurrent goroutine
// reading from the channel from seeing an erroneous "tick".
<原文结束>

# <翻译开始>
// Stop 关闭一个计时器。Stop 之后将不会再发送任何时钟滴答。
// Stop 不会关闭通道，以防止并发的goroutine从通道读取时看到错误的"滴答"。
// md5:b3b968de11e4fb57
# <翻译结束>


<原文开始>
// Reset stops a ticker and resets its period to the specified duration.
// The next tick will arrive after the new period elapses. The duration d
// must be greater than zero; if not, Reset will panic.
<原文结束>

# <翻译开始>
// Reset 停止计时器并将其周期重置为指定的持续时间。
// 下一个Tick将在新的周期结束后到达。duration d必须大于零；否则，Reset将引发恐慌。
// md5:cc38de911b220071
# <翻译结束>


<原文开始>
// A Ticker holds a channel that delivers “ticks” of a clock
// at intervals.
<原文结束>

# <翻译开始>
// Ticker 用于持有在固定间隔发送时钟“滴答”信号的通道。
// md5:b1cdec933a8caca4
# <翻译结束>

