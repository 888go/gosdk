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
	// PST位于西边8小时，PDT位于西边7小时。我们可以使用这个名字，但它并不唯一。. md5:cbe7a8e2be638df6
	if name, off := lt.Zone(); off != -8*60*60 && off != -7*60*60 {
		fmt.Printf("Unable to find US Pacific time zone data for testing; time zone is %q offset %d", name, off)
		fmt.Printf("Likely problem: the time zone files have not been installed.")
	}
}
