package main

import (
	"fmt"
)

func main() {
	if cnt := getCount(); cnt > 5 {
		fmt.Println(cnt, "is greater than 5")
	} else if cnt < 3 {
		fmt.Println(cnt, "is less than 3")
	}

	sum := add(1, 4)
	switch sum.(type) {
	case int:
		fmt.Println("int")
	default:
		fmt.Println("unknown")
	}

	// Note strings are immutable
	stra := "the spice must flow"
	byts := []byte(str)
	strb := string(byts)
}

func getCount() int {
	return 2
}

func add(a interface{}, b interface{}) interface{} {
	return a.(int) + b.(int)
}
