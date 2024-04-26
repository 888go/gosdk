// 版权所有 (C) 2010 Go 作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可协议约束，
// 该协议可在 LICENSE 文件中找到。

package os_test

import (
	. "github.com/888go/gosdk/os"
	"reflect"
	"strings"
	"testing"
)

// testGetenv 为我们提供一组可控的变量，用于测试 Expand。
func testGetenv(s string) string {
	switch s {
	case "*":
		return "all the args"
	case "#":
		return "NARGS"
	case "$":
		return "PID"
	case "1":
		return "ARGUMENT1"
	case "HOME":
		return "/usr/gopher"
	case "H":
		return "(Value of H)"
	case "home_1":
		return "/usr/foo"
	case "_":
		return "underscore"
	}
	return ""
}

var expandTests = []struct {
	in, out string
}{
	{"", ""},
	{"$*", "all the args"},
	{"$$", "PID"},
	{"${*}", "all the args"},
	{"$1", "ARGUMENT1"},
	{"${1}", "ARGUMENT1"},
	{"now is the time", "now is the time"},
	{"$HOME", "/usr/gopher"},
	{"$home_1", "/usr/foo"},
	{"${HOME}", "/usr/gopher"},
	{"${H}OME", "(Value of H)OME"},
	{"A$$$#$1$H$home_1*B", "APIDNARGSARGUMENT1(Value of H)/usr/foo*B"},
	{"start$+middle$^end$", "start$+middle$^end$"},
	{"mixed$|bag$$$", "mixed$|bagPID$"},
	{"$", "$"},
	{"$}", "$}"},
	{"${", ""},  // 无效的语法；消耗掉这些字符
	{"${}", ""}, // 无效的语法；消耗掉这些字符
}

func TestExpand(t *testing.T) {
	for _, test := range expandTests {
		result := Expand(test.in, testGetenv)
		if result != test.out {
			t.Errorf("Expand(%q)=%q; expected %q", test.in, result, test.out)
		}
	}
}

var global any

func BenchmarkExpand(b *testing.B) {
	b.Run("noop", func(b *testing.B) {
		var s string
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			s = Expand("tick tick tick tick", func(string) string { return "" })
		}
		global = s
	})
	b.Run("multiple", func(b *testing.B) {
		var s string
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			s = Expand("$a $a $a $a", func(string) string { return "boom" })
		}
		global = s
	})
}

func TestConsistentEnviron(t *testing.T) {
	e0 := Environ()
	for i := 0; i < 10; i++ {
		e1 := Environ()
		if !reflect.DeepEqual(e0, e1) {
			t.Fatalf("environment changed")
		}
	}
}

func TestUnsetenv(t *testing.T) {
	const testKey = "GO_TEST_UNSETENV"
	set := func() bool {
		prefix := testKey + "="
		for _, key := range Environ() {
			if strings.HasPrefix(key, prefix) {
				return true
			}
		}
		return false
	}
	if err := Setenv(testKey, "1"); err != nil {
		t.Fatalf("Setenv: %v", err)
	}
	if !set() {
		t.Error("Setenv didn't set TestUnsetenv")
	}
	if err := Unsetenv(testKey); err != nil {
		t.Fatalf("Unsetenv: %v", err)
	}
	if set() {
		t.Fatal("Unsetenv didn't clear TestUnsetenv")
	}
}

func TestClearenv(t *testing.T) {
	const testKey = "GO_TEST_CLEARENV"
	const testValue = "1"

	// reset env
	defer func(origEnv []string) {
		for _, pair := range origEnv {
			// 在Windows中，环境变量可以以=开头
			// https://devblogs.microsoft.com/oldnewthing/20100506-00/?p=14133
			i := strings.Index(pair[1:], "=") + 1
			if err := Setenv(pair[:i], pair[i+1:]); err != nil {
				t.Errorf("Setenv(%q, %q) failed during reset: %v", pair[:i], pair[i+1:], err)
			}
		}
	}(Environ())

	if err := Setenv(testKey, testValue); err != nil {
		t.Fatalf("Setenv(%q, %q) failed: %v", testKey, testValue, err)
	}
	if _, ok := LookupEnv(testKey); !ok {
		t.Errorf("Setenv(%q, %q) didn't set $%s", testKey, testValue, testKey)
	}
	Clearenv()
	if val, ok := LookupEnv(testKey); ok {
		t.Errorf("Clearenv() didn't clear $%s, remained with value %q", testKey, val)
	}
}

func TestLookupEnv(t *testing.T) {
	const smallpox = "SMALLPOX"      // No one has smallpox.
	value, ok := LookupEnv(smallpox) // Should not exist.
	if ok || value != "" {
		t.Fatalf("%s=%q", smallpox, value)
	}
	defer Unsetenv(smallpox)
	err := Setenv(smallpox, "virus")
	if err != nil {
		t.Fatalf("failed to release smallpox virus")
	}
	_, ok = LookupEnv(smallpox)
	if !ok {
		t.Errorf("smallpox release failed; world remains safe but LookupEnv is broken")
	}
}

// 在 Windows 系统中，Environ 会报告以单个 "=" 开头的键。我们需要检查 LookupEnv 和 SetEnv 是否能正确处理这种情况。这与 Go 语言的问题 #49886 有关。
func TestEnvironConsistency(t *testing.T) {
	t.Parallel()

	for _, kv := range Environ() {
		i := strings.Index(kv, "=")
		if i == 0 {
			// 在实践中，我们在 Windows 系统中观察到键前面只有一个 "="。
			// TODO(#49886): 我们应该只将第一个 "=" 作为键的一部分进行处理，还是解析任意数量的 "="，直到遇到非 "="，或者尝试每个可能的键/值边界，直到 LookupEnv 成功？
			i = strings.Index(kv[1:], "=") + 1
		}
		if i < 0 {
			t.Errorf("Environ entry missing '=': %q", kv)
		}

		k := kv[:i]
		v := kv[i+1:]
		v2, ok := LookupEnv(k)
		if ok && v == v2 {
			t.Logf("LookupEnv(%q) = %q, %t", k, v2, ok)
		} else {
			t.Errorf("Environ contains %q, but LookupEnv(%q) = %q, %t", kv, k, v2, ok)
		}

		// 由于 k=v 已经存在于环境中，
		// 对其进行设置应为一个空操作（no-op）。
		if err := Setenv(k, v); err == nil {
			t.Logf("Setenv(%q, %q)", k, v)
		} else {
			t.Errorf("Environ contains %q, but SetEnv(%q, %q) = %q", kv, k, v, err)
		}
	}
}
