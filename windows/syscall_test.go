// 版权所有 ? 2013 The Go 作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可协议约束，
// 该协议可在 LICENSE 文件中找到。

//go:build windows

package windows_test

import (
	"strings"
	"syscall"
	"testing"

	"github.com/888go/gosdk/windows"
)

func testSetGetenv(t *testing.T, key, value string) {
	err := windows.Setenv(key, value)
	if err != nil {
		t.Fatalf("Setenv failed to set %q: %v", value, err)
	}
	newvalue, found := windows.Getenv(key)
	if !found {
		t.Fatalf("Getenv failed to find %v variable (want value %q)", key, value)
	}
	if newvalue != value {
		t.Fatalf("Getenv(%v) = %q; want %q", key, newvalue, value)
	}
}

func TestEnv(t *testing.T) {
	testSetGetenv(t, "TESTENV", "AVALUE")
	// 确保 TESTENV 被设置为 ""，而非被删除
	testSetGetenv(t, "TESTENV", "")
}

func TestGetProcAddressByOrdinal(t *testing.T) {
// 尝试调用 shlwapi.dll 中的 IsOS 函数，通过序数进行解析，如
// 下文所示：
	h, err := windows.LoadLibrary("shlwapi.dll")
	if err != nil {
		t.Fatalf("Failed to load shlwapi.dll: %s", err)
	}
	procIsOS, err := windows.GetProcAddressByOrdinal(h, 437)
	if err != nil {
		t.Fatalf("Could not find shlwapi.dll:IsOS by ordinal: %s", err)
	}
	const OS_NT = 1
	r, _, _ := syscall.Syscall(procIsOS, 1, OS_NT, 0, 0)
	if r == 0 {
		t.Error("shlwapi.dll:IsOS(OS_NT) returned 0, expected non-zero value")
	}
}

func TestGetSystemDirectory(t *testing.T) {
	d, err := windows.GetSystemDirectory()
	if err != nil {
		t.Fatalf("Failed to get system directory: %s", err)
	}
	if !strings.HasSuffix(strings.ToLower(d), "\\system32") {
		t.Fatalf("System directory does not end in system32: %s", d)
	}
}

func TestGetWindowsDirectory(t *testing.T) {
	d1, err := windows.GetWindowsDirectory()
	if err != nil {
		t.Fatalf("Failed to get Windows directory: %s", err)
	}
	d2, err := windows.GetSystemWindowsDirectory()
	if err != nil {
		t.Fatalf("Failed to get system Windows directory: %s", err)
	}
	if !strings.HasSuffix(strings.ToLower(d1), `\windows`) {
		t.Fatalf("Windows directory does not end in windows: %s", d1)
	}
	if !strings.HasSuffix(strings.ToLower(d2), `\windows`) {
		t.Fatalf("System Windows directory does not end in windows: %s", d2)
	}
}
func TestFindProcByOrdinal(t *testing.T) {
// 尝试调用 shlwapi.dll 中的 IsOS 函数，通过序数进行解析，如
// 下文所示：
	dll, err := windows.LoadDLL("shlwapi.dll")
	if err != nil {
		t.Fatalf("Failed to load shlwapi.dll: %s", err)
	}
	procIsOS, err := dll.FindProcByOrdinal(437)
	if err != nil {
		t.Fatalf("Could not find shlwapi.dll:IsOS by ordinal: %s", err)
	}
	if procIsOS.Name != "#437" {
		t.Fatalf("Proc's name is incorrect: %s,expected #437", procIsOS.Name)
	}
	const OS_NT = 1
	r, _, _ := syscall.Syscall(procIsOS.Addr(), 1, OS_NT, 0, 0)
	if r == 0 {
		t.Error("shlwapi.dll:IsOS(OS_NT) returned 0, expected non-zero value")
	}
}
