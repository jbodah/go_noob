package main

import (
	"fmt"
)

func main() {
	lookup := make(map[string]int)
	lookup["goku"] = 9001
	power, exists := lookup["vegeta"]
	fmt.Println(power, exists)
	fmt.Println(len(lookup))
	delete(lookup, "goku")

	// initial size
	lookup2 := make(map[string]int, 100)
}
