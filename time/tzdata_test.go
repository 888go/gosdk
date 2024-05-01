// 版权所有 2020 Go 作者。保留所有权利。
// 使用此源代码受BSD风格许可证管辖，该许可证可从LICENSE文件中找到。
// md5:97282fec9c34c48c

package time_test

import (
	"reflect"
	"testing"
	_ "time/tzdata"
)

var zones = []string{
	"Asia/Jerusalem",
	"America/Los_Angeles",
}

//
//func TestEmbeddedTZData(t *testing.T) {
//	undo := time.DisablePlatformSources()
//	defer undo()
//
//	for _, zone := range zones {
//		ref, err := time.LoadLocation(zone)
//		if err != nil {
//			t.Errorf("LoadLocation(%q): %v", zone, err)
//			continue
//		}
//
//		embedded, err := time.LoadFromEmbeddedTZData(zone)
//		if err != nil {
//			t.Errorf("LoadFromEmbeddedTZData(%q): %v", zone, err)
//			continue
//		}
//		sample, err := time.LoadLocationFromTZData(zone, []byte(embedded))
//		if err != nil {
//			t.Errorf("LoadLocationFromTZData failed for %q: %v", zone, err)
//			continue
//		}
//
//// 比较ref和sample的name和zone字段。
//// tx字段随着tzdata更新而更快变化。
//// 预期cache字段会有所不同。
//// md5:de893c4f4d242062
//		v1 := reflect.ValueOf(ref).Elem()
//		v2 := reflect.ValueOf(sample).Elem()
//		typ := v1.Type()
//		nf := typ.NumField()
//		found := 0
//		for i := 0; i < nf; i++ {
//			ft := typ.Field(i)
//			if ft.Name != "name" && ft.Name != "zone" {
//				continue
//			}
//			found++
//			if !equal(t, v1.Field(i), v2.Field(i)) {
//				t.Errorf("zone %s: system and embedded tzdata field %s differs", zone, ft.Name)
//			}
//		}
//		if found != 2 {
//			t.Errorf("test must be updated for change to time.Location struct")
//		}
//	}
//}

// equal 是一个小型的 reflect.DeepEqual，我们使用它来比较时区信息中未导出字段的值。
// md5:47fbb24254dc5be7
func equal(t *testing.T, f1, f2 reflect.Value) bool {
	switch f1.Type().Kind() {
	case reflect.Slice:
		if f1.Len() != f2.Len() {
			return false
		}
		for i := 0; i < f1.Len(); i++ {
			if !equal(t, f1.Index(i), f2.Index(i)) {
				return false
			}
		}
		return true
	case reflect.Struct:
		nf := f1.Type().NumField()
		for i := 0; i < nf; i++ {
			if !equal(t, f1.Field(i), f2.Field(i)) {
				return false
			}
		}
		return true
	case reflect.String:
		return f1.String() == f2.String()
	case reflect.Bool:
		return f1.Bool() == f2.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return f1.Int() == f2.Int()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return f1.Uint() == f2.Uint()
	default:
		t.Errorf("test internal error: unsupported kind %v", f1.Type().Kind())
		return true
	}
}
