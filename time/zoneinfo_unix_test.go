// 版权所有 2020 Go 作者。保留所有权利。
// 使用此源代码受BSD风格许可证管辖，该许可证可从LICENSE文件中找到。
// md5:97282fec9c34c48c

//---build---//go:build unix && !ios && !android

package time_test

import (
	"github.com/888go/gosdk/time"
	"os"
	"testing"
)

func TestEnvTZUsage(t *testing.T) {
	const env = "TZ"
	tz, ok := os.LookupEnv(env)
	if !ok {
		defer os.Unsetenv(env)
	} else {
		defer os.Setenv(env, tz)
	}
	defer time.ForceUSPacificForTesting()

	localZoneName := "Local"
	// 文件可能不存在。. md5:722f2ec072d9ec25
	if _, err := os.Stat("/etc/localtime"); os.IsNotExist(err) {
		localZoneName = "UTC"
	}

	cases := []struct {
		nilFlag bool
		tz      string
		local   string
	}{
		// 没有$TZ表示使用系统的默认设置（/etc/localtime）。. md5:4ee8a1d41d00a89d
		{true, "", localZoneName},
		// $TZ="" means use UTC.
		{false, "", "UTC"},
		{false, ":", "UTC"},
		{false, "Asia/Shanghai", "Asia/Shanghai"},
		{false, ":Asia/Shanghai", "Asia/Shanghai"},
		{false, "/etc/localtime", localZoneName},
		{false, ":/etc/localtime", localZoneName},
	}

	for _, c := range cases {
		time.ResetLocalOnceForTest()
		if c.nilFlag {
			os.Unsetenv(env)
		} else {
			os.Setenv(env, c.tz)
		}
		if time.Local.String() != c.local {
			t.Errorf("invalid Local location name for %q: got %q want %q", c.tz, time.Local, c.local)
		}
	}

	time.ResetLocalOnceForTest()
	// 在Solaris 2和IRIX 6上，该文件可能不存在。. md5:a0af2a56773e8a83
	path := "/usr/share/zoneinfo/Asia/Shanghai"
	os.Setenv(env, path)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if time.Local.String() != "UTC" {
			t.Errorf(`invalid path should fallback to UTC: got %q want "UTC"`, time.Local)
		}
		return
	}
	if time.Local.String() != path {
		t.Errorf(`custom path should lead to path itself: got %q want %q`, time.Local, path)
	}

	timeInUTC := time.Date(2009, 1, 1, 12, 0, 0, 0, time.UTC)
	sameTimeInShanghai := time.Date(2009, 1, 1, 20, 0, 0, 0, time.Local)
	if !timeInUTC.Equal(sameTimeInShanghai) {
		t.Errorf("invalid timezone: got %q want %q", timeInUTC, sameTimeInShanghai)
	}

	time.ResetLocalOnceForTest()
	os.Setenv(env, ":"+path)
	if time.Local.String() != path {
		t.Errorf(`custom path should lead to path itself: got %q want %q`, time.Local, path)
	}

	time.ResetLocalOnceForTest()
	os.Setenv(env, path[:len(path)-1])
	if time.Local.String() != "UTC" {
		t.Errorf(`invalid path should fallback to UTC: got %q want "UTC"`, time.Local)
	}
}
