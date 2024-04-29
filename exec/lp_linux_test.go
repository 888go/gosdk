// 版权所有 ? 2022 Go作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

package exec_test

import (
	"errors"
	"internal/syscall/unix"
	"internal/testenv"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
	"testing"
)

func TestFindExecutableVsNoexec(t *testing.T) {
	t.Parallel()

	// 该测试用例依赖于Linux v5.8中出现的faccessat2(2)系统调用。
	if major, minor := unix.KernelVersion(); major < 5 || (major == 5 && minor < 8) {
		t.Skip("requires Linux kernel v5.8 with faccessat2(2) syscall")
	}

	tmp := t.TempDir()

	// Create a tmpfs mount.
	err := syscall.Mount("tmpfs", tmp, "tmpfs", 0, "")
	if testenv.SyscallIsNotSupported(err) {
// 通常这意味着缺少 CAP_SYS_ADMIN 权限，但在受限的测试环境中可能有其他原因。
		t.Skipf("requires ability to mount tmpfs (%v)", err)
	} else if err != nil {
		t.Fatalf("mount %s failed: %v", tmp, err)
	}
	t.Cleanup(func() {
		if err := syscall.Unmount(tmp, 0); err != nil {
			t.Error(err)
		}
	})

	// Create an executable.
	path := filepath.Join(tmp, "program")
	err = os.WriteFile(path, []byte("#!/bin/sh\necho 123\n"), 0o755)
	if err != nil {
		t.Fatal(err)
	}

	// 检查其是否按预期工作。
	_, err = exec.LookPath(path)
	if err != nil {
		t.Fatalf("LookPath: got %v, want nil", err)
	}

	for {
		err = exec.Command(path).Run()
		if err == nil {
			break
		}
		if errors.Is(err, syscall.ETXTBSY) {
// 另一个进程中的fork+exec操作可能正持有我们用于写入可执行文件的文件描述符（参见https://go.dev/issue/22315）。
// 由于该描述符应已设置CLOEXEC标志，当fork出的子进程到达其exec调用时，问题应得到解决。
// 在此之前持续重试。
		} else {
			t.Fatalf("exec: got %v, want nil", err)
		}
	}

	// 使用noexec标志重新挂载
	err = syscall.Mount("", tmp, "", syscall.MS_REMOUNT|syscall.MS_NOEXEC, "")
	if testenv.SyscallIsNotSupported(err) {
		t.Skipf("requires ability to re-mount tmpfs (%v)", err)
	} else if err != nil {
		t.Fatalf("remount %s with noexec failed: %v", tmp, err)
	}

	if err := exec.Command(path).Run(); err == nil {
		t.Fatal("exec on noexec filesystem: got nil, want error")
	}

	_, err = exec.LookPath(path)
	if err == nil {
		t.Fatalf("LookPath: got nil, want error")
	}
}
