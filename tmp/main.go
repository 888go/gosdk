package main

import (
	"fmt"
	"github.com/888go/gosdk/exec"
)

func main() {
	返回 := exec.Command("go.exe", "env")
	返回2, _ := 返回.Output()

	fmt.Println(string(返回2))
}
