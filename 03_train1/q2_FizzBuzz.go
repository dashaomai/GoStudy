package main

import (
	"fmt"
)

func main() {
	for i, m := 1, 100; i < m; i++ {
		switch {
		case (i%3) == 0 && (i%5) == 0:
			fmt.Printf("FizzBuzz\n")
		case (i % 3) == 0:
			fmt.Printf("Fizz\n")
		case (i % 5) == 0:
			fmt.Printf("Buzz\n")
		default:
			fmt.Printf("%d\n", i)
		}
	}
}
