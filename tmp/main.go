package main

import (
	"fmt"
	"strconv"
)

func main() {
	sr, _ := strconv.Unquote("\"大\t家\t好！\"")
	fmt.Println(sr)
	sr, _ = strconv.Unquote(`"大家好！"`)
	fmt.Println(sr)

}
