package main

import (
	"fmt"
)

func Sum(addend, adder int) int {
	return addend + adder
}

func main() {
	var (
		a = 10
		b = 20
	)
	fmt.Printf("Sum of %d and %d is %d.", a, b, Sum(a, b))
}
