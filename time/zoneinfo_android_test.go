//版权所有2016 The Go Authors. 所有权利保留。
//使用此源代码受BSD风格
//可以在LICENSE文件中找到的许可协议。
// md5:675a20b63ebe97ed

package time_test

import (
	. "github.com/888go/gosdk/time"
	"testing"
)

func TestAndroidTzdata(t *testing.T) {
	undo := ForceAndroidTzdataForTest()
	defer undo()
	if _, err := LoadLocation("America/Los_Angeles"); err != nil {
		t.Error(err)
	}
}
