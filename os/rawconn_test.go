// 版权所有 2018 The Go Authors。保留所有权利。
// 使用本源代码须遵循 BSD 风格许可证，
// 该许可证可在 LICENSE 文件中找到。

// Test use of raw connections.
//
//go:build !plan9 && !js && !wasip1

package os_test

import (
	"github.com/888go/gosdk/os"
	"syscall"
	"testing"
)

func TestRawConnReadWrite(t *testing.T) {
	t.Parallel()

	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()
	defer w.Close()

	rconn, err := r.SyscallConn()
	if err != nil {
		t.Fatal(err)
	}
	wconn, err := w.SyscallConn()
	if err != nil {
		t.Fatal(err)
	}

	var operr error
	err = wconn.Write(func(s uintptr) bool {
		_, operr = syscall.Write(syscallDescriptor(s), []byte{'b'})
		return operr != syscall.EAGAIN
	})
	if err != nil {
		t.Fatal(err)
	}
	if operr != nil {
		t.Fatal(err)
	}

	var n int
	buf := make([]byte, 1)
	err = rconn.Read(func(s uintptr) bool {
		n, operr = syscall.Read(syscallDescriptor(s), buf)
		return operr != syscall.EAGAIN
	})
	if err != nil {
		t.Fatal(err)
	}
	if operr != nil {
		t.Fatal(operr)
	}
	if n != 1 {
		t.Errorf("read %d bytes, expected 1", n)
	}
	if buf[0] != 'b' {
		t.Errorf("read %q, expected %q", buf, "b")
	}
}
