// 版权所有 2011 Go 语言作者。保留所有权利。
// 使用此源代码须遵循 BSD 风格的许可协议，
// 许可协议具体内容可在 LICENSE 文件中找到。
// md5:b5bcbe8a534f6077

package time

//
//func init() {
//	// 强制使用美国/太平洋时区进行时区测试。. md5:46c4cbaf6eb4451c
//	ForceUSPacificForTesting()
//}
//
//func initTestingZone() {
//// 为了保证测试的封闭性，只使用测试的GOROOT目录下的tzinfo源，
//// 而不是系统的源代码，也不使用进程中可能设置的任何GOROOT环境变量。
//// 这个测试在GOROOT/src/time目录下运行，所以GOROOT是"../.."，
//// 但理论上可能存在其他情况。
//// md5:9df2aa7dd1c5efc8
//	sources := []string{"../../lib/time/zoneinfo.zip"}
//	z, err := loadLocation("America/Los_Angeles", sources)
//	if err != nil {
//		panic("cannot load America/Los_Angeles for testing: " + err.Error() + "; you may want to use -tags=timetzdata")
//	}
//	z.name = "Local"
//	localLoc = *z
//}
//
//var origPlatformZoneSources []string = platformZoneSources
//
//func disablePlatformSources() (undo func()) {
//	platformZoneSources = nil
//	return func() {
//		platformZoneSources = origPlatformZoneSources
//	}
//}
//
//var Interrupt = interrupt
//var DaysIn = daysIn
//
//func empty(arg any, seq uintptr) {}
//
//// 测试当运行时计时器的周期在过期时会溢出，但不会抛出异常或导致其他计时器挂起。
////
//// 这个测试必须放在internal_test.go中，因为它操作未公开的数据结构。
//// md5:6a4e803f3b22aef2
//func CheckRuntimeTimerPeriodOverflow() {
//// 我们手动创建一个具有巨大周期的runtimeTimer，但该定时器会立即到期。公共Timer接口需要在第一次更新之前等待整个周期。
//// md5:a4249b95fbeedfc0
//	r := &runtimeTimer{
//		when:   runtimeNano(),
//		period: 1<<63 - 1,
//		f:      empty,
//		arg:    nil,
//	}
//	startTimer(r)
//	defer stopTimer(r)
//
//// 如果此测试失败，我们将会抛出错误（当siftdownTimer在更新时检测到bad状态时），或者其它定时器将会挂起（如果堆中的定时器处于不良状态）。没有可靠的方法来测试这一点，但我们会在这里等待一个短时间的定时器作为基本功能测试（或者，后续测试中的定时器可能会挂起）。
//// md5:da77407212bfa69a
//	<-After(25 * Millisecond)
//}
//
//var (
//	MinMonoTime = Time{wall: 1 << 63, ext: -1 << 63, loc: UTC}
//	MaxMonoTime = Time{wall: 1 << 63, ext: 1<<63 - 1, loc: UTC}
//
//	NotMonoNegativeTime = Time{wall: 0, ext: -1<<63 + 50}
//)
