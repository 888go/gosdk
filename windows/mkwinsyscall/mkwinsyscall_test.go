// 版权所有 2023 The Go Authors。保留所有权利。
// 使用本源代码须遵循 BSD 风格的许可证，
// 该许可证可在 LICENSE 文件中找到。

package main

import (
	"bytes"
	"go/format"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestDLLFilenameEscaping(t *testing.T) {
	tests := []struct {
		name     string
		filename string
	}{
		{"no escaping necessary", "kernel32"},
		{"escape period", "windows.networking"},
		{"escape dash", "api-ms-win-wsl-api-l1-1-0"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 将一个虚构的系统调用写入临时文件以供测试
			const prefix = "package windows\n//sys Example() = "
			const suffix = ".Example"
			name := filepath.Join(t.TempDir(), "syscall.go")
			if err := os.WriteFile(name, []byte(prefix+tt.filename+suffix), 0666); err != nil {
				t.Fatal(err)
			}

// 确保解析、生成和格式化操作能无误运行。
// 这足以说明转义功能正常工作。
			src, err := ParseFiles([]string{name})
			if err != nil {
				t.Fatal(err)
			}
			var buf bytes.Buffer
			if err := src.Generate(&buf); err != nil {
				t.Fatal(err)
			}
			if _, err := format.Source(buf.Bytes()); err != nil {
				t.Log(buf.String())
				t.Fatal(err)
			}
		})
	}
}

func TestSyscallXGeneration(t *testing.T) {
	tests := []struct {
		name        string
		wantsysfunc string
		sig         string
	}{
		{
			name:        "syscall with 2 params",
			wantsysfunc: "syscall.Syscall",
			sig:         "Example(a1 *uint16, a2 *uint16) = ",
		},
		{
			name:        "syscall with 6 params",
			wantsysfunc: "syscall.Syscall6",
			sig:         "Example(a1 *uint, a2 *uint, a3 *uint, a4 *uint, a5 *uint, a6 *uint) = ",
		},
		{
			name:        "syscall with 15 params",
			wantsysfunc: "syscall.Syscall15",
			sig: strings.ReplaceAll(`Example(a1 *uint, a2 *uint, a3 *uint, a4 *uint, a5 *uint, a6 *uint,
						a7 *uint, a8 *uint, a9 *uint, a10 *uint, a11 *uint, a12 *uint,
						a13 *uint, a14 *uint, a15 *uint) = `, "\n", ""),
		},
		{
			name:        "syscall with 18 params",
			wantsysfunc: "syscall.SyscallN",
			sig: strings.ReplaceAll(`Example(a1 *uint, a2 *uint, a3 *uint, a4 *uint, a5 *uint, a6 *uint,
						a7 *uint, a8 *uint, a9 *uint, a10 *uint, a11 *uint, a12 *uint,
						a13 *uint, a14 *uint, a15 *uint, a16 *uint, a17 *uint, a18 *uint) = `, "\n", ""),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 将系统调用写入临时文件以供测试。
			prefix := "package windows\n//sys " + tt.sig
			suffix := ".Example"
			name := filepath.Join(t.TempDir(), "syscall.go")
			if err := os.WriteFile(name, []byte(prefix+"example"+suffix), 0666); err != nil {
				t.Fatal(err)
			}

// 确保解析、生成和格式化操作能无误运行。
// 这足以说明转义功能正常工作。
			src, err := ParseFiles([]string{name})
			if err != nil {
				t.Fatal(err)
			}
			var buf bytes.Buffer
			if err := src.Generate(&buf); err != nil {
				t.Fatal(err)
			}
			if _, err := format.Source(buf.Bytes()); err != nil {
				t.Fatal(err)
			}

			if !strings.Contains(buf.String(), tt.wantsysfunc+"(") {
				t.Fatalf("expected syscall func %q in buffer %s", tt.wantsysfunc, buf.String())
			}
		})
	}
}
