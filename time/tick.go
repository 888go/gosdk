package time

import "time"

// Ticker 用于持有在固定间隔发送时钟“滴答”信号的通道。
// md5:b1cdec933a8caca4
type Ticker struct {
	F time.Ticker
} //md5:cf16cd3fe215ef83

// Tick is a convenience wrapper for NewTicker providing access to the ticking
// channel only. While Tick is useful for clients that have no need to shut down
// the Ticker, be aware that without a way to shut it down the underlying
// Ticker cannot be recovered by the garbage collector; it "leaks".
// Unlike NewTicker, Tick will return nil if d <= 0.

// ff:
// d:
func Tick(d Duration) <-chan time.Time {
	if d <= 0 {
		return nil
	}
	return NewTicker(d).F.C
}

// NewTicker 返回一个新的 Ticker，其中包含一个通道，该通道将在每次滴答声后发送当前时间。
// 滴答的周期由 duration 参数指定。如果接收者接收较慢，ticker 会调整时间间隔或丢弃滴答信号来补偿。
// 参数 d 必须大于零；如果不满足此条件，NewTicker 将引发恐慌。停止 ticker 以释放相关资源。
// md5:8d80705ea00f018e
// 翻译提示:func  创建计时器(d  时长)  *计时器  {}

// ff:
// d:
func NewTicker(d Duration) *Ticker { //md5:75190d9bc3baea99
	t := time.NewTicker(time.Duration(d))
	if t == nil {
		return nil
	}
	return &Ticker{*t}
}

// Stop 关闭一个计时器。Stop 之后将不会再发送任何时钟滴答。
// Stop 不会关闭通道，以防止并发的goroutine从通道读取时看到错误的"滴答"。
// md5:b3b968de11e4fb57
// 翻译提示:func  (t  *计时器)  停止()  {}

// ff:
func (t *Ticker) Stop() { //md5:15f11c7e2d81aa41
	t.F.Stop()
}

// Reset 停止计时器并将其周期重置为指定的持续时间。
// 下一个Tick将在新的周期结束后到达。duration d必须大于零；否则，Reset将引发恐慌。
// md5:cc38de911b220071
// 翻译提示:func  (t  *计时器)  重置(间隔  Duration)  {}

// ff:
// d:
func (t *Ticker) Reset(d Duration) { //md5:abfc23e6f86bacec
	t.F.Reset(time.Duration(d))
}
