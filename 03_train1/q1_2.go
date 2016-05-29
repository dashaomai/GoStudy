package main

import (
	"fmt"
)

func main() {
	var i int
	i = 0
Loop:
	fmt.Printf("\t%d", i)
	i++
	if i < 10 {
		goto Loop
	}
}
