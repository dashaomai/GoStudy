package main

import (
	"fmt"
)

func main() {
	for i, m := 1, 100; i <= m; i++ {
		s1 := make([]rune, i)
		for j := 0; j < i; j++ {
			s1[j] = 'A'
		}

		fmt.Printf("%s\n", string(s1))
	}
}
