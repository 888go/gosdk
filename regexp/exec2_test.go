// 版权所有 2013 Go 作者。保留所有权利。
// 使用此源代码受BSD风格许可证约束，该许可证可从LICENSE文件中找到。
// md5:19d1a3ed91182ee4

//go:build !race

package regexp

//2024-05-01  这个文件没有关于正则类的单元测试
//// This test is excluded when running under the race detector because
//// it is a very expensive test and takes too long.
//func TestRE2Exhaustive(t *testing.T) {
//	if testing.Short() {
//		t.Skip("skipping TestRE2Exhaustive during short test")
//	}
//	testRE2(t, "testdata/re2-exhaustive.txt.bz2")
//}
