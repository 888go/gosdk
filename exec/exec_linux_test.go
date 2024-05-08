// 版权所有 2020 Go作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

//---build---//go:build linux && cgo

// 在使用glibc的系统中，调用malloc可能会创建一个新的内存区域（arena），
// 而创建新的内存区域可能会读取/sys/devices/system/cpu/online文件。
// 如果我们使用cgo，在创建新线程时会调用malloc。
// 这样在我们在检查打开的文件描述符时，如果创建的新线程恰好创建了新的内存区域并打开了/sys文件，就可能破坏TestExtraFiles测试。
// 为了解决这个问题，提前创建线程作为一种解决办法。
// 参见issue 25628。

package exec_test

import (
	"os"
	"sync"
	"syscall"
	"time"
)

func init() {
	if os.Getenv("GO_EXEC_TEST_PID") == "" {
		return
	}

// 启动一些线程。10 是一个随意选择的数值，但目的是确保代码本身不需要创建任何线程。
// 特别地，这个数值应该足够大，以确保超过垃圾回收器可能创建的线程数量。
	const threads = 10

	var wg sync.WaitGroup
	wg.Add(threads)
	ts := syscall.NsecToTimespec((100 * time.Microsecond).Nanoseconds())
	for i := 0; i < threads; i++ {
		go func() {
			defer wg.Done()
			syscall.Nanosleep(&ts, nil)
		}()
	}
	wg.Wait()
}
