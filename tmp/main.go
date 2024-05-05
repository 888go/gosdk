package main

import (
	"fmt"
	"github.com/888go/gosdk/time"
)

type TwoLines string

var c chan int

func handle(int) {}
func main() {
	lt := time.Now()

	if name, off := lt.Zone(); off != -8*60*60 && off != -7*60*60 {
		fmt.Printf("Unable to find US Pacific time zone data for testing; time zone is %q offset %d", name, off)
		fmt.Printf("Likely problem: the time zone files have not been installed.")
	}
}

type X结构体 struct { //hm:结构体 cz:type Buffer     
	X父类 bytes.Buffer //qm:父类 cz:F bytes.Buffer     
}

const (
	X常量 readOp = -1 //qm:常量 cz:OpRead      readOp     
)
