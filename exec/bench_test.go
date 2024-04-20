// 版权所有 ? 2019 The Go Authors。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

//2024-04-17 备注,单元测试通不过, 保留单元测试文件为了方便查看使用方法.
package exec

import (
	"testing"
)

 
func BenchmarkExecHostname(b *testing.B) {
	b.ReportAllocs()
	path, err := LookPath("hostname")
	if err != nil {
		b.Fatalf("could not find hostname: %v", err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := Command(path).Run(); err != nil {
			b.Fatalf("hostname: %v", err)
		}
	}
}
