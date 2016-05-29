package main

import (
	"fmt"
)

func main() {
	arr := [...]int{0, 1, 3, 5, 7, 9}

	for _, v := range arr {
		fmt.Printf("\t%d", v)
	}
}
