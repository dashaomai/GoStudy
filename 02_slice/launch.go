package main

import (
	"fmt"
)

func main() {
	s0 := []int{0, 0}
	s1 := append(s0, 2)
	s2 := append(s1, 3, 5, 7)
	s3 := append(s2, s0...)
	s4 := append(s2, s0)

	fmt.Printf(s3)
	fmt.Printf(s4)
}
