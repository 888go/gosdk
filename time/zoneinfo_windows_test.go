// 版权所有 2013 Go 作者。保留所有权利。
// 使用此源代码受BSD风格许可证约束，该许可证可从LICENSE文件中找到。
// md5:19d1a3ed91182ee4

package time_test

import (
	. "github.com/888go/gosdk/time"
	"testing"
)

func testZoneAbbr(t *testing.T) {
	t1 := Now()
	// discard nsec
	t1 = Date(t1.Year(), t1.Month(), t1.Day(), t1.Hour(), t1.Minute(), t1.Second(), 0, t1.Location())

	t2, err := Parse(RFC1123, t1.Format(RFC1123))
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}
	if t1 != t2 {
		t.Fatalf("t1 (%v) is not equal to t2 (%v)", t1, t2)
	}
}

//
//func TestUSPacificZoneAbbr(t *testing.T) {
//	ForceUSPacificFromTZIForTesting() // 将Once重置以触发竞争条件. md5:f7940df75a562e17
//	defer ForceUSPacificForTesting()
//	testZoneAbbr(t)
//}
//
//func TestAusZoneAbbr(t *testing.T) {
//	ForceAusFromTZIForTesting()
//	defer ForceUSPacificForTesting()
//	testZoneAbbr(t)
//}
//
//func TestToEnglishName(t *testing.T) {
//	const want = "Central Europe Standard Time"
//	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows NT\CurrentVersion\Time Zones\`+want, registry.READ)
//	if err != nil {
//		t.Fatalf("cannot open CEST time zone information from registry: %s", err)
//	}
//	defer k.Close()
//
//	var std, dlt string
//	// 首先尝试使用 MUI_Std 和 MUI_Dlt，如果遇到任何错误，则回退到 Std 和 Dlt。. md5:4839ee3ec81db384
//	std, err = k.GetMUIStringValue("MUI_Std")
//	if err == nil {
//		dlt, err = k.GetMUIStringValue("MUI_Dlt")
//	}
//	if err != nil { // 回退到Std和Dlt. md5:5980702591820792
//		if std, _, err = k.GetStringValue("Std"); err != nil {
//			t.Fatalf("cannot read CEST Std registry key: %s", err)
//		}
//		if dlt, _, err = k.GetStringValue("Dlt"); err != nil {
//			t.Fatalf("cannot read CEST Dlt registry key: %s", err)
//		}
//	}
//
//	name, err := ToEnglishName(std, dlt)
//	if err != nil {
//		t.Fatalf("toEnglishName failed: %s", err)
//	}
//	if name != want {
//		t.Fatalf("english name: %q, want: %q", name, want)
//	}
//}
