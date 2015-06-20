package main

import (
	"fmt"
)

func main() {
	reslicing()
	arrAlg()
	// slice from literals
	names := []string{"leto", "jessica", "paul"}
	// when you want memory to be initialized already
	checks := make([]bool, 10)
	// when you want an arbitrary number
	var strs []string
	// when you want an initial capacity
	scores := make([]int, 0, 20)
	fmt.Printf("%d %d %d %d\n", len(names), len(checks), len(strs), len(scores))
}

func arrAlg() {
	scores := make([]int, 0, 5)
	c := cap(scores)
	fmt.Println(c)

	for i := 0; i < 25; i++ {
		scores = append(scores, i)

		// Prove that go will change capacity
		// capacity is doubled
		if cap(scores) != c {
			c = cap(scores)
			fmt.Println(c)
		}
	}
}

func reslicing() {
	// slice with capacity 10, default size 0
	scores := make([]int, 0, 10)
	// compile error: scores[5] = 9033
	scores = scores[0:6]
	scores[5] = 9033
	fmt.Println(scores)
}
