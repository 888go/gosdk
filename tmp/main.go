package main

import (
	"fmt"
	"github.com/888go/gosdk/os"
)

func main() {
	测试, _ := os.Readlink("e:\\Users\\admin\\Desktop\\1.lnk")

	fmt.Println(测试)
}
