package main

import (
	"fmt"
)

// Declare a function type
type Add func(a int, b int) int

func main() {
	// Pass an anonymous function
	fmt.Println(process(func(a int, b int) int {
		return a + b
	}))
}

// Accept the function type
func process(adder Add) int {
	return adder(1, 2)
}
