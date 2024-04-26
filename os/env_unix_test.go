// 版权所有 ? 2013 The Go 作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可协议约束，
// 该协议可在 LICENSE 文件中找到。

//go:build unix

package os_test

import (
	"fmt"
	. "github.com/888go/gosdk/os"
	"testing"
)

var setenvEinvalTests = []struct {
	k, v string
}{
	{"", ""},      // empty key
	{"k=v", ""},   // '=' in key
	{"\x00", ""},  // '\x00' in key
	{"k", "\x00"}, // '\x00' in value
}

func TestSetenvUnixEinval(t *testing.T) {
	for _, tt := range setenvEinvalTests {
		err := Setenv(tt.k, tt.v)
		if err == nil {
			t.Errorf(`Setenv(%q, %q) == nil, want error`, tt.k, tt.v)
		}
	}
}

var shellSpecialVarTests = []struct {
	k, v string
}{
	{"*", "asterisk"},
	{"#", "pound"},
	{"$", "dollar"},
	{"@", "at"},
	{"!", "exclamation mark"},
	{"?", "question mark"},
	{"-", "dash"},
}

func TestExpandEnvShellSpecialVar(t *testing.T) {
	for _, tt := range shellSpecialVarTests {
		Setenv(tt.k, tt.v)
		defer Unsetenv(tt.k)

		argRaw := fmt.Sprintf("$%s", tt.k)
		argWithBrace := fmt.Sprintf("${%s}", tt.k)
		if gotRaw, gotBrace := ExpandEnv(argRaw), ExpandEnv(argWithBrace); gotRaw != gotBrace {
			t.Errorf("ExpandEnv(%q) = %q, ExpandEnv(%q) = %q; expect them to be equal", argRaw, gotRaw, argWithBrace, gotBrace)
		}
	}
}
