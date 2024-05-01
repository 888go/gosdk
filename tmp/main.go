package main

import (
	"fmt"
)

type TwoLines string

func main() {
	input := "abc\ndef\n"
	// Sscan should work
	var tscan TwoLines
	n, err := fmt.Sscan(input, &tscan)

	fmt.Println(n, err)
	fmt.Println(tscan)
}
