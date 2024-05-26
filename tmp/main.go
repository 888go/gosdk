package main

import (
	"fmt"
	"os"
)

func main() {
	dirPath := "E:\\SVN\\GO\\code\\WanCheng\\FanYiGo" // 替换为你要读取的目录路径
	file, err := os.Open(dirPath)
	if err != nil {
		fmt.Println("Error opening directory:", err)
		return
	}
	defer file.Close()

	names, err := file.Readdirnames(-1) // -1 表示读取所有文件
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	for _, name := range names {
		fmt.Println(name)
	}
}
