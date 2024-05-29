package main

import (
	"fmt"
	"strings"
)

func main() {
	// 创建一个Replacer实例，指定要替换的键值对
	replacer := strings.NewReplacer(
		"apple", "orange",
		"banana", "grape",
		"cherry", "blueberry",
	)

	// 原始字符串
	input := "I love apples and bananas, but not cherries."

	// 使用Replacer替换字符串中的子串
	changed := replacer.Replace(input)

	// 输出结果
	fmt.Println(changed)
}
