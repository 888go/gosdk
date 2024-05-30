package main

import (
	"fmt"
	"time"
)

func main() {
	unixTime := time.Now()
	fmt.Println(unixTime.UnixMicro())
}
