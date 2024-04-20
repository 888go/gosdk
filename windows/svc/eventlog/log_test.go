// 版权所有 ? 2012 The Go Authors。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

//go:build windows

package eventlog_test

import (
	"testing"

	"github.com/888go/gosdk/windows/svc/eventlog"
)

func TestLog(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode - it modifies system logs")
	}

	const name = "mylog"
	const supports = eventlog.Error | eventlog.Warning | eventlog.Info
	err := eventlog.InstallAsEventCreate(name, supports)
	if err != nil {
		t.Fatalf("Install failed: %s", err)
	}
	defer func() {
		err = eventlog.Remove(name)
		if err != nil {
			t.Fatalf("Remove failed: %s", err)
		}
	}()

	l, err := eventlog.Open(name)
	if err != nil {
		t.Fatalf("Open failed: %s", err)
	}
	defer l.Close()

	err = l.Info(1, "info")
	if err != nil {
		t.Fatalf("Info failed: %s", err)
	}
	err = l.Warning(2, "warning")
	if err != nil {
		t.Fatalf("Warning failed: %s", err)
	}
	err = l.Error(3, "error")
	if err != nil {
		t.Fatalf("Error failed: %s", err)
	}
}
