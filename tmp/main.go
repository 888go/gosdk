package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := string([]byte{7, 8, 9, 10, 11, 12, 13}) // 包含不可打印的ASCII控制字符
	graphicQuoted := strconv.QuoteToGraphic(s)
	fmt.Println(graphicQuoted)
}
