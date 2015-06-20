package main

import (
	"fmt"
)

func main() {
	_, exists := power("goku")
	if exists == false {
		fmt.Println("rut roh")
	}
}

func power(name string) (int, bool) {
	return 9000, false
}

// Both a and b are ints
func add(a, b int) int {
	return 0
}
