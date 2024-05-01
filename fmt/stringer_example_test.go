// 版权所有 2017 年 The Go 语言作者。保留所有权利。
// 本源代码的使用受 BSD 风格许可协议的约束，
// 该许可协议可在 LICENSE 文件中找到。
// md5:680e85412fa918bd

package fmt_test

import (
	"github.com/888go/gosdk/fmt"
)

// Animal 具有 Name 和 Age，用于表示一个动物。. md5:cefdbfe6593f5f74
type Animal struct {
	Name string
	Age  uint
}

// String实现了Animal接口的Stringer方法。. md5:365336f5175098f1
func (a Animal) String() string {
	return fmt.Sprintf("%v (%d)", a.Name, a.Age)
}

func ExampleStringer() {
	a := Animal{
		Name: "Gopher",
		Age:  2,
	}
	fmt.Println(a)
	// Output: Gopher (2)
}
