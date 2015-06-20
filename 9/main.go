package main

import (
	"fmt"
)

func main() {
	arrays()
	slices()
}

func slices() {
	scores := []int{1, 4, 293, 4, 9}
	// make = allocate memory + initialize
	// new = allocate memory
	for _, value := range scores {
		fmt.Printf("Scores Slice: %d\n", value)
	}
	scorez := make([]int, 10)
	for _, value := range scorez {
		fmt.Printf("Scorez Initialized Slice: %d\n", value)
	}
}

func arrays() {
	var scores [10]int
	scores[0] = 339

	powers := [4]int{9001, 9333, 212, 33}

	fmt.Printf("There are %d powers\n", len(powers))
	for _, value := range scores {
		fmt.Printf("Scores Array: %d\n", value)
	}
}
