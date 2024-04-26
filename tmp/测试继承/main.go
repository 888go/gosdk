package main

import "fmt"

type 父类 struct {
	名称 string
}

func 创建父类() *父类 {
	return nil
	//return &父类{
	//	名称: "123",
	//}
}

func (a 父类) 输出父类() {
	fmt.Println("这是父类")
}

type 子类 struct {
	F *父类
}

func 创建子类() *子类 {
	//panic("创建父类nil")
	return &子类{
		F: 创建父类(),
	}
}

func (a *子类) 输出子类() {
	//	a.F.输出父类()
	fmt.Println("这是子类")
}

func main() {
	返回 := 创建子类()
	fmt.Println(返回)
	返回.输出子类()
}
