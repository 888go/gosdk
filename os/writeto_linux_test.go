// 版权所有 2023 The Go Authors。保留所有权利。
// 使用本源代码须遵循 BSD 风格的许可证，
// 该许可证可在 LICENSE 文件中找到。

package os_test

import (
	"bytes"
	"internal/poll"
	"io"
	"math/rand"
	"net"
	. "os"
	"strconv"
	"syscall"
	"testing"
	"time"
)

func TestSendFile(t *testing.T) {
	sizes := []int{
		1,
		42,
		1025,
		syscall.Getpagesize() + 1,
		32769,
	}
	t.Run("sendfile-to-unix", func(t *testing.T) {
		for _, size := range sizes {
			t.Run(strconv.Itoa(size), func(t *testing.T) {
				testSendFile(t, "unix", int64(size))
			})
		}
	})
	t.Run("sendfile-to-tcp", func(t *testing.T) {
		for _, size := range sizes {
			t.Run(strconv.Itoa(size), func(t *testing.T) {
				testSendFile(t, "tcp", int64(size))
			})
		}
	})
}

func testSendFile(t *testing.T, proto string, size int64) {
	dst, src, recv, data, hook := newSendFileTest(t, proto, size)

	// 现在调用WriteTo（通过io.Copy），这将希望调用poll.SendFile
	n, err := io.Copy(dst, src)
	if err != nil {
		t.Fatalf("io.Copy error: %v", err)
	}

	// 我们应该使用正确的文件描述符参数调用poll.Splice。
	if n > 0 && !hook.called {
		t.Fatal("expected to called poll.SendFile")
	}
	if hook.called && hook.srcfd != int(src.Fd()) {
		t.Fatalf("wrong source file descriptor: got %d, want %d", hook.srcfd, src.Fd())
	}
	sc, ok := dst.(syscall.Conn)
	if !ok {
		t.Fatalf("destination is not a syscall.Conn")
	}
	rc, err := sc.SyscallConn()
	if err != nil {
		t.Fatalf("destination SyscallConn error: %v", err)
	}
	if err = rc.Control(func(fd uintptr) {
		if hook.called && hook.dstfd != int(fd) {
			t.Fatalf("wrong destination file descriptor: got %d, want %d", hook.dstfd, int(fd))
		}
	}); err != nil {
		t.Fatalf("destination Conn Control error: %v", err)
	}

	// 验证数据的大小和内容。
	dataSize := len(data)
	dstData := make([]byte, dataSize)
	m, err := io.ReadFull(recv, dstData)
	if err != nil {
		t.Fatalf("server Conn Read error: %v", err)
	}
	if n != int64(dataSize) {
		t.Fatalf("data length mismatch for io.Copy, got %d, want %d", n, dataSize)
	}
	if m != dataSize {
		t.Fatalf("data length mismatch for net.Conn.Read, got %d, want %d", m, dataSize)
	}
	if !bytes.Equal(dstData, data) {
		t.Errorf("data mismatch, got %s, want %s", dstData, data)
	}
}

// newSendFileTest 初始化一个用于sendfile的全新测试。
//
// 它创建源文件与目标套接字，并使用指定大小的随机数据填充源文件。同时，它对package os调用poll.Sendfile进行挂钩，并返回该挂钩以便进行检查。
func newSendFileTest(t *testing.T, proto string, size int64) (net.Conn, *File, net.Conn, []byte, *sendFileHook) {
	t.Helper()

	hook := hookSendFile(t)

	client, server := createSocketPair(t, proto)
	tempFile, data := createTempFile(t, size)

	return client, tempFile, server, data, hook
}

func hookSendFile(t *testing.T) *sendFileHook {
	h := new(sendFileHook)
	h.install()
	t.Cleanup(h.uninstall)
	return h
}

type sendFileHook struct {
	called bool
	dstfd  int
	srcfd  int
	remain int64

	written int64
	handled bool
	err     error

	original func(dst *poll.FD, src int, remain int64) (int64, error, bool)
}

func (h *sendFileHook) install() {
	h.original = *PollSendFile
	*PollSendFile = func(dst *poll.FD, src int, remain int64) (int64, error, bool) {
		h.called = true
		h.dstfd = dst.Sysfd
		h.srcfd = src
		h.remain = remain
		h.written, h.err, h.handled = h.original(dst, src, remain)
		return h.written, h.err, h.handled
	}
}

func (h *sendFileHook) uninstall() {
	*PollSendFile = h.original
}

func createTempFile(t *testing.T, size int64) (*File, []byte) {
	f, err := CreateTemp(t.TempDir(), "writeto-sendfile-to-socket")
	if err != nil {
		t.Fatalf("failed to create temporary file: %v", err)
	}
	t.Cleanup(func() {
		f.Close()
	})

	randSeed := time.Now().Unix()
	t.Logf("random data seed: %d\n", randSeed)
	prng := rand.New(rand.NewSource(randSeed))
	data := make([]byte, size)
	prng.Read(data)
	if _, err := f.Write(data); err != nil {
		t.Fatalf("failed to create and feed the file: %v", err)
	}
	if err := f.Sync(); err != nil {
		t.Fatalf("failed to save the file: %v", err)
	}
	if _, err := f.Seek(0, io.SeekStart); err != nil {
		t.Fatalf("failed to rewind the file: %v", err)
	}

	return f, data
}
