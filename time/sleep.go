package time

import "time"

// Timer 类型表示单个事件。
// 当 Timer 到期时，除非该 Timer 是通过 AfterFunc 创建的，否则当前时间将被发送到 C 通道。
// Timer 必须使用 NewTimer 或 AfterFunc 创建。
// md5:8e661907d899fb2d
type Timer struct {//hm:延时器结构  cz:type Timer  
	F time.Timer
} //md5:c8aa453b91475f19

// After waits for the duration to elapse and then sends the current time
// on the returned channel.
// It is equivalent to NewTimer(d).C.
// The underlying Timer is not recovered by the garbage collector
// until the timer fires. If efficiency is a concern, use NewTimer
// instead and call Timer.Stop if the timer is no longer needed.
// ff:
// d:时长
func After(d Duration) <-chan time.Time {
	return NewTimer(d).F.C
}

// Sleep pauses the current goroutine for at least the duration d.
// A negative or zero duration causes Sleep to return immediately.
// ff:延时
// d:时长
func Sleep(d Duration) {
	time.Sleep(time.Duration(d))
}

// Stop 阻止定时器触发。如果调用成功停止了定时器，它会返回 true；如果定时器已经超时或已停止，则返回 false。
// Stop 不关闭通道，以防止通道读取错误地成功。
//
// 要确保在调用 Stop 后，通道为空，检查返回值并消耗通道中的所有数据。例如，假设程序尚未从 t.C 中接收：
//
//	if !t.Stop() {
//		<-t.C
//	}
//
// 这不能与从定时器通道的其他接收操作或对 Stop 方法的其他调用同时进行。
//
// 如果使用 AfterFunc(d, f) 创建定时器，如果 t.Stop 返回 false，那么定时器已经超时，函数 f 已经在一个单独的goroutine中启动；Stop 在返回之前不会等待 f 完成。
// 如果调用者需要知道 f 是否已完成，必须与 f 显式协调。
// md5:f5fa62452dbbe4f2
// ff:停止
// t:
func (t *Timer) Stop() bool { //md5:fc8cbe1af620cb13
	return t.F.Stop()
}

// NewTimer 创建一个新的 Timer，该 Timer 在至少经过 duration d 后会通过其通道发送当前时间。
// md5:2767c6954bd0c243
// ff:创建延时器
// d:时长
func NewTimer(d Duration) *Timer { //md5:11e89ccf3aaee20c
	t := time.NewTimer(time.Duration(d))
	if t == nil {
		return nil
	}
	return &Timer{*t}
}

// Reset 将计时器的过期时间重新设置为持续时间 d 后。
// 如果计时器之前是活动状态，则返回 true；如果计时器已过期或被停止，则返回 false。
//
// 对于使用 NewTimer 创建的计时器，Reset 应仅在已停止或已过期且通道已被清空的计时器上调用。
//
// 如果程序已经从 t.C 接收了一个值，那么可以确定计时器已过期且通道已清空，因此可以直接使用 t.Reset。
// 然而，如果程序还未从 t.C 接收到值，计时器必须先被停止，且—如果 Stop 报告计时器在被停止前已过期—则需要显式清空通道：
//
//	if !t.Stop() {
//		<-t.C
//	}
//	t.Reset(d)
//
// 这不应该与从计时器的通道进行其他接收操作同时进行。
//
// 注意，无法正确使用 Reset 的返回值，因为在清空通道和新计时器过期之间存在竞态条件。
// 如上所述，应始终在已停止或已过期的通道上调用 Reset。返回值的存在是为了保持与现有程序的兼容性。
//
// 对于使用 AfterFunc(d, f) 创建的计时器，Reset 要么重新安排 f 的执行时间（在这种情况下返回 true），要么安排 f 再次运行（返回 false）。
// 当 Reset 返回 false 时，它既不等待先前的 f 执行完成就返回，也不保证后续运行 f 的协程不会与前一个并发执行。
// 如果调用者需要了解前一次 f 的执行是否已完成，必须与 f 显式协调。
// md5:af314813b00709db
// ff:重置时长
// t:
// d:时长
func (t *Timer) Reset(d Duration) bool { //md5:ab76211fe40d3fa1
	return t.F.Reset(time.Duration(d))
}

// AfterFunc 等待duration时间过去后，在一个新的goroutine中调用函数f。它返回一个Timer，可以使用其Stop方法来取消调用。
// 返回的Timer的C字段不被使用，将为nil。
// md5:3f56daf6dab5cec5
// ff:延迟执行
// d:时长
// f:回调函数
func AfterFunc(d Duration, f func()) *Timer { //md5:c6cacd4aadc22dcb
	t := time.AfterFunc(time.Duration(d), f)
	if t == nil {
		return nil
	}
	return &Timer{*t}
}
