// 版权所有 ? 2021 Go作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

package fdtest

import (
	"os"
	"runtime"
	"testing"
)

func TestExists(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Exists not implemented for windows")
	}

	if !Exists(os.Stdout.Fd()) {
		t.Errorf("Exists(%d) got false want true", os.Stdout.Fd())
	}
}
