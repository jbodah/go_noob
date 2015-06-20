package main

import (
	"fmt"
)

type Saiyan struct {
	Name  string
	Power int
}

func main() {
	goku := Saiyan{
		Name:  "Goku",
		Power: 9000,
	}
	josh := Saiyan{"Josh", 15000}
	Super(&goku)
	Super2(&josh)
	fmt.Println(goku.Power)
	fmt.Println(josh.Power)
}

func Super(s *Saiyan) {
	s.Power += 10000
}

func Super2(s *Saiyan) {
	// prove that s is a copy of the address of josh
	s = &Saiyan{"Gohan", 1000}
}
