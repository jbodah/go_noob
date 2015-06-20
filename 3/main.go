package main

import (
	"fmt"
)

func main() {
	name, power := "Goku", getPower()
	fmt.Printf("%s's power is over %d\n", name, power)
}

func getPower() int {
	return 9000
}
