//版权所有2018 Go语言作者。所有权利保留。
//使用此源代码受BSD风格的
//可在LICENSE文件中找到的许可证管辖。
// md5:b62d507ecbeb2e98

package fmt_test

import (
	"github.com/888go/gosdk/fmt"
)

// Address 包含城市、州和国家信息。. md5:5d6891003d64916a
type Address struct {
	City    string
	State   string
	Country string
}

// Person有一个Name、Age和Address属性。. md5:d29468e58cc31e7c
type Person struct {
	Name string
	Age  uint
	Addr *Address
}

// GoString让Person实现了GoStringer接口。
// 返回的值是有效的Go代码，可以用来复制Person结构体。
// md5:240e6a6ef5248fc4
func (p Person) GoString() string {
	if p.Addr != nil {
		return fmt.Sprintf("Person{Name: %q, Age: %d, Addr: &Address{City: %q, State: %q, Country: %q}}", p.Name, int(p.Age), p.Addr.City, p.Addr.State, p.Addr.Country)
	}
	return fmt.Sprintf("Person{Name: %q, Age: %d}", p.Name, int(p.Age))
}

func ExampleGoStringer() {
	p1 := Person{
		Name: "Warren",
		Age:  31,
		Addr: &Address{
			City:    "Denver",
			State:   "CO",
			Country: "U.S.A.",
		},
	}
// 如果没有实现GoString()方法，`fmt.Printf("%#v", p1)` 的输出将会类似于
// Person{Name:"Warren", Age:31, Addr:(*main.Address)(0x10448240)}
// md5:06accd0b6f883adc
	fmt.Printf("%#v\n", p1)

	p2 := Person{
		Name: "Theia",
		Age:  4,
	}
// 如果没有实现GoString()方法，`fmt.Printf("%#v", p2)`的输出将会类似于
// Person{Name:"Theia", Age:0x4, Addr:(*main.Address)(nil)}
// md5:df5825dd4a82752c
	fmt.Printf("%#v\n", p2)

	// Output:
	// Person{Name: "Warren", Age: 31, Addr: &Address{City: "Denver", State: "CO", Country: "U.S.A."}}
	// Person{Name: "Theia", Age: 4}
}
