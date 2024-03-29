package runtime_test

import (
	"github.com/888go/gosdk/runtime"
	"sync"
	"testing"
)

var wg sync.WaitGroup

func init() {
	runtime.NetpollGenericInit()
}

func BenchmarkNetpollBreak(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10; j++ {
			wg.Add(1)
			go func() {
				runtime.NetpollBreak()
				wg.Done()
			}()
		}
	}
	wg.Wait()
	b.StopTimer()
}
