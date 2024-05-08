package main

import (
	"fmt"
)

// 定义动物接口
type 动物 interface {
	叫() string
}

// 猫的类型
type 猫 struct{}

// 猫叫的方法
func (c 猫) 叫() string {
	return "喵~"
}

// 狗的类型
type 狗 struct{}

// 狗叫的方法
func (d 狗) 叫() string {
	return "汪~"
}

// 叫声函数，接受一个动物接口作为参数
func 叫声(animal 动物) {
	fmt.Println(animal.叫())
}

func main() {
	// 创建猫和狗的实例
	miao := 猫{}
	wang := 狗{}

	// 调用叫声函数，传入不同类型的实例
	叫声(miao) // 输出: 喵~
	叫声(wang) // 输出: 汪~
}
